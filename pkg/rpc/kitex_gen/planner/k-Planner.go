// Code generated by Kitex v0.8.0. DO NOT EDIT.

package planner

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/datasinks"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/exprs"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/partitions"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/plannodes"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/querycache"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/types"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = datasinks.KitexUnusedProtection
	_ = exprs.KitexUnusedProtection
	_ = partitions.KitexUnusedProtection
	_ = plannodes.KitexUnusedProtection
	_ = querycache.KitexUnusedProtection
	_ = types.KitexUnusedProtection
)

func (p *TPlanFragment) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPartition bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPartition = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 9:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField9(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetPartition {
		fieldId = 6
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TPlanFragment[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TPlanFragment[fieldId]))
}

func (p *TPlanFragment) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := plannodes.NewTPlan()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Plan = tmp
	return offset, nil
}

func (p *TPlanFragment) FastReadField4(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.OutputExprs = make([]*exprs.TExpr, 0, size)
	for i := 0; i < size; i++ {
		_elem := exprs.NewTExpr()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.OutputExprs = append(p.OutputExprs, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *TPlanFragment) FastReadField5(buf []byte) (int, error) {
	offset := 0

	tmp := datasinks.NewTDataSink()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.OutputSink = tmp
	return offset, nil
}

func (p *TPlanFragment) FastReadField6(buf []byte) (int, error) {
	offset := 0

	tmp := partitions.NewTDataPartition()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Partition = tmp
	return offset, nil
}

func (p *TPlanFragment) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.MinReservationBytes = &v

	}
	return offset, nil
}

func (p *TPlanFragment) FastReadField8(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.InitialReservationTotalClaims = &v

	}
	return offset, nil
}

func (p *TPlanFragment) FastReadField9(buf []byte) (int, error) {
	offset := 0

	tmp := querycache.NewTQueryCacheParam()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.QueryCacheParam = tmp
	return offset, nil
}

// for compatibility
func (p *TPlanFragment) FastWrite(buf []byte) int {
	return 0
}

func (p *TPlanFragment) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TPlanFragment")
	if p != nil {
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField8(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField9(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TPlanFragment) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TPlanFragment")
	if p != nil {
		l += p.field2Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
		l += p.field9Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TPlanFragment) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPlan() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "plan", thrift.STRUCT, 2)
		offset += p.Plan.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetOutputExprs() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "output_exprs", thrift.LIST, 4)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
		var length int
		for _, v := range p.OutputExprs {
			length++
			offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetOutputSink() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "output_sink", thrift.STRUCT, 5)
		offset += p.OutputSink.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "partition", thrift.STRUCT, 6)
	offset += p.Partition.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPlanFragment) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetMinReservationBytes() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "min_reservation_bytes", thrift.I64, 7)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.MinReservationBytes)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) fastWriteField8(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetInitialReservationTotalClaims() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "initial_reservation_total_claims", thrift.I64, 8)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.InitialReservationTotalClaims)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) fastWriteField9(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetQueryCacheParam() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "query_cache_param", thrift.STRUCT, 9)
		offset += p.QueryCacheParam.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPlanFragment) field2Length() int {
	l := 0
	if p.IsSetPlan() {
		l += bthrift.Binary.FieldBeginLength("plan", thrift.STRUCT, 2)
		l += p.Plan.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPlanFragment) field4Length() int {
	l := 0
	if p.IsSetOutputExprs() {
		l += bthrift.Binary.FieldBeginLength("output_exprs", thrift.LIST, 4)
		l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.OutputExprs))
		for _, v := range p.OutputExprs {
			l += v.BLength()
		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPlanFragment) field5Length() int {
	l := 0
	if p.IsSetOutputSink() {
		l += bthrift.Binary.FieldBeginLength("output_sink", thrift.STRUCT, 5)
		l += p.OutputSink.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPlanFragment) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("partition", thrift.STRUCT, 6)
	l += p.Partition.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TPlanFragment) field7Length() int {
	l := 0
	if p.IsSetMinReservationBytes() {
		l += bthrift.Binary.FieldBeginLength("min_reservation_bytes", thrift.I64, 7)
		l += bthrift.Binary.I64Length(*p.MinReservationBytes)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPlanFragment) field8Length() int {
	l := 0
	if p.IsSetInitialReservationTotalClaims() {
		l += bthrift.Binary.FieldBeginLength("initial_reservation_total_claims", thrift.I64, 8)
		l += bthrift.Binary.I64Length(*p.InitialReservationTotalClaims)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPlanFragment) field9Length() int {
	l := 0
	if p.IsSetQueryCacheParam() {
		l += bthrift.Binary.FieldBeginLength("query_cache_param", thrift.STRUCT, 9)
		l += p.QueryCacheParam.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TScanRangeLocation) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetServer bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetServer = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetServer {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TScanRangeLocation[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TScanRangeLocation[fieldId]))
}

func (p *TScanRangeLocation) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := types.NewTNetworkAddress()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Server = tmp
	return offset, nil
}

func (p *TScanRangeLocation) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.VolumeId = v

	}
	return offset, nil
}

func (p *TScanRangeLocation) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.BackendId = &v

	}
	return offset, nil
}

// for compatibility
func (p *TScanRangeLocation) FastWrite(buf []byte) int {
	return 0
}

func (p *TScanRangeLocation) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TScanRangeLocation")
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TScanRangeLocation) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TScanRangeLocation")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TScanRangeLocation) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "server", thrift.STRUCT, 1)
	offset += p.Server.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TScanRangeLocation) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetVolumeId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "volume_id", thrift.I32, 2)
		offset += bthrift.Binary.WriteI32(buf[offset:], p.VolumeId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TScanRangeLocation) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBackendId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "backend_id", thrift.I64, 3)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.BackendId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TScanRangeLocation) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("server", thrift.STRUCT, 1)
	l += p.Server.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TScanRangeLocation) field2Length() int {
	l := 0
	if p.IsSetVolumeId() {
		l += bthrift.Binary.FieldBeginLength("volume_id", thrift.I32, 2)
		l += bthrift.Binary.I32Length(p.VolumeId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TScanRangeLocation) field3Length() int {
	l := 0
	if p.IsSetBackendId() {
		l += bthrift.Binary.FieldBeginLength("backend_id", thrift.I64, 3)
		l += bthrift.Binary.I64Length(*p.BackendId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TScanRangeLocations) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetScanRange bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetScanRange = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetScanRange {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TScanRangeLocations[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TScanRangeLocations[fieldId]))
}

func (p *TScanRangeLocations) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := plannodes.NewTScanRange()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.ScanRange = tmp
	return offset, nil
}

func (p *TScanRangeLocations) FastReadField2(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Locations = make([]*TScanRangeLocation, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewTScanRangeLocation()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.Locations = append(p.Locations, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *TScanRangeLocations) FastWrite(buf []byte) int {
	return 0
}

func (p *TScanRangeLocations) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TScanRangeLocations")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TScanRangeLocations) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TScanRangeLocations")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TScanRangeLocations) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "scan_range", thrift.STRUCT, 1)
	offset += p.ScanRange.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TScanRangeLocations) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "locations", thrift.LIST, 2)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.Locations {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TScanRangeLocations) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("scan_range", thrift.STRUCT, 1)
	l += p.ScanRange.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TScanRangeLocations) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("locations", thrift.LIST, 2)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.Locations))
	for _, v := range p.Locations {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
