package handle

import (
	"regexp"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/selectdb/ccr_syncer/pkg/ccr"
	"github.com/selectdb/ccr_syncer/pkg/ccr/base"
	"github.com/selectdb/ccr_syncer/pkg/ccr/record"
	festruct "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/frontendservice"
	tstatus "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/status"
	"github.com/selectdb/ccr_syncer/pkg/xerror"
	log "github.com/sirupsen/logrus"
)

func init() {
	ccr.RegisterJobHandle[*record.Upsert](festruct.TBinlogType_UPSERT, &UpsertHandle{})
}

type UpsertHandle struct {
}

func (h *UpsertHandle) IsIdempotent() bool {
	return true
}

func (h *UpsertHandle) IsBinlogCommitted(j *ccr.Job, upsert *record.Upsert) (bool, error) {
	return false, nil
}

func Committed(j *ccr.Job) {
	inMemoryData := j.GetInMemoryData().(*ccr.InMemoryData)
	log.Debugf("txn %d committed, commitSeq: %d, cleanup", inMemoryData.TxnId, j.GetJobProgress().CommitSeq)
	commitSeq := j.GetJobProgress().CommitSeq
	destTableIds := inMemoryData.DestTableIds
	if j.SyncType == ccr.DBSync && len(j.GetJobProgress().TableCommitSeqMap) > 0 {
		for _, tableId := range destTableIds {
			tableCommitSeq, ok := j.GetJobProgress().TableCommitSeqMap[tableId]
			if !ok {
				continue
			}

			if tableCommitSeq < commitSeq {
				j.GetJobProgress().TableCommitSeqMap[tableId] = commitSeq
			}
		}

		j.GetJobProgress().Persist()
	}
	j.GetJobProgress().Done()
}

func Rollback(j *ccr.Job, err error, inMemoryData *ccr.InMemoryData) {
	log.Debugf("txn %d rollback, commitSeq: %d, cleanup", inMemoryData.TxnId, j.GetJobProgress().CommitSeq)
	j.GetJobProgress().Done()
}

func UpdateInMemory(j *ccr.Job, upsert *record.Upsert) error {
	inMemoryData := j.GetInMemoryData().(*ccr.InMemoryData)
	inMemoryData.TxnId = upsert.TxnID
	inMemoryData.CommitSeq = upsert.CommitSeq
	inMemoryData.DestTableIds = make([]int64, 0, len(upsert.TableRecords))
	for tableId := range upsert.TableRecords {
		inMemoryData.DestTableIds = append(inMemoryData.DestTableIds, tableId)
	}
	return nil
}

func UpsertDone(j *ccr.Job, upsert *record.Upsert) error {
	log.Tracef("upsert: %v", upsert)

	// Step 1: get related tableRecords
	var isTxnInsert bool = false
	if len(upsert.Stids) > 0 {
		if !ccr.FeatureTxnInsert {
			log.Warnf("The txn insert is not supported yet")
			return xerror.Errorf(xerror.Normal, "The txn insert is not supported yet")
		}
		isTxnInsert = true
	}

	tableRecords, err := j.GetRelatedTableRecords(upsert)
	if err != nil {
		log.Errorf("get related table records failed, err: %+v", err)
		return err
	}
	if len(tableRecords) == 0 {
		log.Debug("no related table records")
		return nil
	}

	destTableIds := make([]int64, 0, len(tableRecords))
	if j.SyncType == ccr.DBSync {
		savedRecords := make([]*record.TableRecord, 0, len(tableRecords))
		for _, tableRecord := range tableRecords {
			if isAsyncMv, err := j.IsMaterializedViewTable(tableRecord.Id); err != nil {
				return err
			} else if isAsyncMv {
				// ignore the upsert of materialized view table.
				continue
			} else if destTableId, err := j.GetDestTableIdBySrc(tableRecord.Id); err != nil {
				return err
			} else {
				savedRecords = append(savedRecords, tableRecord)
				destTableIds = append(destTableIds, destTableId)
			}
		}
		tableRecords = savedRecords
	} else {
		destTableIds = append(destTableIds, j.Dest.TableId)
	}
	if len(tableRecords) == 0 {
		log.Debug("no related table records")
		return nil
	}

	log.Debugf("handle upsert, table records: %v", tableRecords)
	inMemoryData := &ccr.InMemoryData{
		CommitSeq:    upsert.CommitSeq,
		DestTableIds: destTableIds,
		TableRecords: tableRecords,
		IsTxnInsert:  isTxnInsert,
		SourceStids:  upsert.Stids,
		Label:        upsert.Label,
	}
	j.GetJobProgress().NextSubVolatile(ccr.BeginTransaction, inMemoryData)
	return nil
}

func UpsertBeginTransaction(j *ccr.Job, upsert *record.Upsert, dest *base.Spec) error {
	// Step 2: begin txn
	inMemoryData := j.GetJobProgress().InMemoryData.(*ccr.InMemoryData)
	commitSeq := inMemoryData.CommitSeq
	sourceStids := inMemoryData.SourceStids
	isTxnInsert := inMemoryData.IsTxnInsert

	destRpc, err := j.GetJobFactory().NewFeRpc(dest)
	if err != nil {
		return err
	}

	var label string
	if j.Extra.ReuseBinlogLabel {
		label = inMemoryData.Label
	} else {
		label = j.NewLabel(commitSeq)
	}
	log.Tracef("begin txn, label: %s, dest: %v, commitSeq: %d", label, dest, commitSeq)

	var beginTxnResp *festruct.TBeginTxnResult_
	if isTxnInsert {
		// when txn insert, give an array length in BeginTransaction, it will return a list of stid
		beginTxnResp, err = destRpc.BeginTransactionForTxnInsert(dest, label, inMemoryData.DestTableIds, int64(len(sourceStids)))
	} else {
		beginTxnResp, err = destRpc.BeginTransaction(dest, label, inMemoryData.DestTableIds)
	}
	if err != nil {
		return err
	}
	log.Tracef("begin txn resp: %v", beginTxnResp)

	if beginTxnResp.GetStatus().GetStatusCode() != tstatus.TStatusCode_OK {
		if isTableNotFound(beginTxnResp.GetStatus()) && j.SyncType == ccr.DBSync {
			// It might caused by the staled TableMapping entries.
			// In order to rebuild the dest table ids, this progress should be rollback.
			j.GetJobProgress().Rollback()
			for _, tableRecord := range inMemoryData.TableRecords {
				delete(j.GetJobProgress().TableMapping, tableRecord.Id)
			}
		}
		return xerror.Errorf(xerror.Normal, "begin txn failed, status: %v", beginTxnResp.GetStatus())
	}
	txnId := beginTxnResp.GetTxnId()
	if isTxnInsert {
		destStids := beginTxnResp.GetSubTxnIds()
		inMemoryData.DestStids = destStids
		log.Infof("begin txn %d, label: %s, db: %d, destStids: %v",
			txnId, label, beginTxnResp.GetDbId(), destStids)
	} else {
		log.Infof("begin txn %d, label: %s, db: %d", txnId, label, beginTxnResp.GetDbId())
	}

	inMemoryData.TxnId = txnId
	j.GetJobProgress().NextSubCheckpoint(ccr.IngestBinlog, inMemoryData)
	return nil
}

func UpsertIngestBinlog(j *ccr.Job, upsert *record.Upsert) error {
	log.Trace("ingest binlog")
	if err := UpdateInMemory(j, upsert); err != nil {
		return err
	}
	inMemoryData := j.GetJobProgress().InMemoryData.(*ccr.InMemoryData)
	tableRecords := inMemoryData.TableRecords
	txnId := inMemoryData.TxnId
	isTxnInsert := inMemoryData.IsTxnInsert
	commitSeq := inMemoryData.CommitSeq

	// make stidMap, source_stid to dest_stid
	stidMap := make(map[int64]int64)
	if isTxnInsert {
		sourceStids := inMemoryData.SourceStids
		destStids := inMemoryData.DestStids
		if len(sourceStids) == len(destStids) {
			for i := 0; i < len(sourceStids); i++ {
				stidMap[sourceStids[i]] = destStids[i]
			}
		}
	}

	// Step 3: ingest binlog
	if isTxnInsert {
		var allSubTxnInfos = make([]*festruct.TSubTxnInfo, 0, len(stidMap))
		for _, destTableId := range inMemoryData.DestTableIds {
			// When txn insert, use subTxnInfos to commit rather than commitInfos.
			subTxnInfos, err := j.IngestBinlogForTxnInsert(commitSeq, txnId, tableRecords, stidMap, destTableId)
			if err == ccr.ErrTriggerPartialSnapshot {
				if j.Extra.PartialSnapshotParams == nil {
					panic("partial snapshot params is nil when trigger partial snapshot")
				}
				j.GetJobProgress().NextSubCheckpoint(ccr.RollbackTransaction, inMemoryData)
			} else if err != nil {
				Rollback(j, err, inMemoryData)
				return err
			} else {
				subTxnInfos := subTxnInfos
				allSubTxnInfos = append(allSubTxnInfos, subTxnInfos...)
				j.GetJobProgress().NextSubCheckpoint(ccr.CommitTransaction, inMemoryData)
			}
		}
		inMemoryData.SubTxnInfos = allSubTxnInfos
	} else {
		commitInfos, err := j.IngestBinlog(commitSeq, txnId, tableRecords)
		if err == ccr.ErrTriggerPartialSnapshot {
			if j.Extra.PartialSnapshotParams == nil {
				panic("partial snapshot params is nil when trigger partial snapshot")
			}
			j.GetJobProgress().NextSubCheckpoint(ccr.RollbackTransaction, inMemoryData)
		} else if err != nil {
			Rollback(j, err, inMemoryData)
			return err
		} else {
			inMemoryData.CommitInfos = commitInfos
			j.GetJobProgress().NextSubCheckpoint(ccr.CommitTransaction, inMemoryData)
		}
	}
	return nil
}

func UpsertCommitTransaction(j *ccr.Job, upsert *record.Upsert, dest *base.Spec) error {
	// Step 4: commit txn
	log.Tracef("commit txn")
	if err := UpdateInMemory(j, upsert); err != nil {
		return err
	}
	inMemoryData := j.GetJobProgress().InMemoryData.(*ccr.InMemoryData)
	txnId := inMemoryData.TxnId
	commitInfos := inMemoryData.CommitInfos

	destRpc, err := j.GetJobFactory().NewFeRpc(dest)
	if err != nil {
		Rollback(j, err, inMemoryData)
		return err
	}

	isTxnInsert := inMemoryData.IsTxnInsert
	subTxnInfos := inMemoryData.SubTxnInfos
	var resp *festruct.TCommitTxnResult_
	if isTxnInsert {
		resp, err = destRpc.CommitTransactionForTxnInsert(dest, txnId, true, subTxnInfos)
	} else {
		onlyCommitTxn := ccr.FeatureSkipWaitingTxnPublish
		resp, err = destRpc.CommitTransaction(dest, txnId, commitInfos, onlyCommitTxn)
	}
	if err != nil {
		Rollback(j, err, inMemoryData)
		return err
	}
	log.Tracef("commit txn %d resp: %v", txnId, resp)

	if statusCode := resp.Status.GetStatusCode(); statusCode == tstatus.TStatusCode_PUBLISH_TIMEOUT {
		dest.WaitTransactionDone(txnId)
	} else if statusCode != tstatus.TStatusCode_OK {
		err := xerror.Errorf(xerror.Normal, "commit txn failed, status: %v", resp.Status)
		Rollback(j, err, inMemoryData)
		return err
	}

	log.Infof("commit txn %d success", txnId)
	Committed(j)
	return nil
}

func UpsertRollbackTransaction(j *ccr.Job, upsert *record.Upsert, dest *base.Spec) error {
	log.Tracef("Rollback txn")
	// Not Step 5: just rollback txn
	if err := UpdateInMemory(j, upsert); err != nil {
		return err
	}

	inMemoryData := j.GetJobProgress().InMemoryData.(*ccr.InMemoryData)
	txnId := inMemoryData.TxnId
	destRpc, err := j.GetJobFactory().NewFeRpc(dest)
	if err != nil {
		return err
	}

	resp, err := destRpc.RollbackTransaction(dest, txnId)
	if err != nil {
		return err
	}
	if resp.Status.GetStatusCode() != tstatus.TStatusCode_OK {
		if isTxnNotFound(resp.Status) {
			log.Warnf("txn not found, txnId: %d", txnId)
		} else if isTxnAborted(resp.Status) {
			log.Infof("txn already aborted, txnId: %d", txnId)
		} else if isTxnCommitted(resp.Status) {
			log.Infof("txn already committed, txnId: %d", txnId)
			Committed(j)
			return nil
		} else {
			return xerror.Errorf(xerror.Normal, "rollback txn failed, status: %v", resp.Status)
		}
	}

	log.Infof("rollback txn %d success", txnId)
	j.GetJobProgress().Rollback()
	return nil
}

func (h *UpsertHandle) Handle(j *ccr.Job, commitSeq int64, upsert *record.Upsert) error {
	dest := &j.Dest
	var err error
	switch j.GetJobProgress().SubSyncState {
	case ccr.Done:
		err = UpsertDone(j, upsert)
		if err != nil {
			logger.Errorf("upsert done failed, err: %+v", err)
			return err
		}
	case ccr.BeginTransaction:
		err = UpsertBeginTransaction(j, upsert, dest)
		if err != nil {
			logger.Errorf("upsert begin transaction failed, err: %+v", err)
			return err
		}
	case ccr.IngestBinlog:
		err = UpsertIngestBinlog(j, upsert)
		if err != nil {
			logger.Errorf("upsert ingest binlog failed, err: %+v", err)
			return err
		}
	case ccr.CommitTransaction:
		err = UpsertCommitTransaction(j, upsert, dest)
		if err != nil {
			logger.Errorf("upsert commit transaction failed, err: %+v", err)
			return err
		}
	case ccr.RollbackTransaction:
		err = UpsertRollbackTransaction(j, upsert, dest)
		if err != nil {
			logger.Errorf("upsert rollback transaction failed, err: %+v", err)
			return err
		}
	default:
		return xerror.Errorf(xerror.Normal, "invalid job sub sync state %d", j.GetJobProgress().SubSyncState)
	}

	return h.Handle(j, commitSeq, upsert)
}

func isTxnNotFound(status *tstatus.TStatus) bool {
	errMessages := status.GetErrorMsgs()
	for _, errMessage := range errMessages {
		// detailMessage = transaction not found
		// or detailMessage = transaction [12356] not found
		if strings.Contains(errMessage, "transaction not found") || regexp.MustCompile(`transaction \[\d+\] not found`).MatchString(errMessage) {
			return true
		}
	}
	return false
}

func isTxnCommitted(status *tstatus.TStatus) bool {
	return isStatusContainsAny(status, "is already COMMITTED")
}

func isTxnAborted(status *tstatus.TStatus) bool {
	return isStatusContainsAny(status, "is already aborted")
}

func isTableNotFound(status *tstatus.TStatus) bool {
	// 1. FE FrontendServiceImpl.beginTxnImpl
	// 2. FE FrontendServiceImpl.commitTxnImpl
	// 3. FE Table.tryWriteLockOrMetaException
	return isStatusContainsAny(status, "can't find table id:", "table not found", "unknown table")
}

func isStatusContainsAny(status *tstatus.TStatus, patterns ...string) bool {
	errMessages := status.GetErrorMsgs()
	for _, errMessage := range errMessages {
		for _, substr := range patterns {
			if strings.Contains(errMessage, substr) {
				return true
			}
		}
	}
	return false
}
