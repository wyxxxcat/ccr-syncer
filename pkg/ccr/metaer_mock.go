// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License
// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ccr/metaer.go
//
// Generated by this command:
//
//	mockgen -source=pkg/ccr/metaer.go -destination=pkg/ccr/metaer_mock.go -package=ccr
//
// Package ccr is a generated GoMock package.
package ccr

import (
	reflect "reflect"

	base "github.com/selectdb/ccr_syncer/pkg/ccr/base"
	rpc "github.com/selectdb/ccr_syncer/pkg/rpc"
	btree "github.com/tidwall/btree"
	gomock "go.uber.org/mock/gomock"
)

// MockMetaCleaner is a mock of MetaCleaner interface.
type MockMetaCleaner struct {
	ctrl     *gomock.Controller
	recorder *MockMetaCleanerMockRecorder
}

// MockMetaCleanerMockRecorder is the mock recorder for MockMetaCleaner.
type MockMetaCleanerMockRecorder struct {
	mock *MockMetaCleaner
}

// NewMockMetaCleaner creates a new mock instance.
func NewMockMetaCleaner(ctrl *gomock.Controller) *MockMetaCleaner {
	mock := &MockMetaCleaner{ctrl: ctrl}
	mock.recorder = &MockMetaCleanerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaCleaner) EXPECT() *MockMetaCleanerMockRecorder {
	return m.recorder
}

// ClearDB mocks base method.
func (m *MockMetaCleaner) ClearDB(dbName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearDB", dbName)
}

// ClearDB indicates an expected call of ClearDB.
func (mr *MockMetaCleanerMockRecorder) ClearDB(dbName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDB", reflect.TypeOf((*MockMetaCleaner)(nil).ClearDB), dbName)
}

// ClearTable mocks base method.
func (m *MockMetaCleaner) ClearTable(dbName, tableName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearTable", dbName, tableName)
}

// ClearTable indicates an expected call of ClearTable.
func (mr *MockMetaCleanerMockRecorder) ClearTable(dbName, tableName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearTable", reflect.TypeOf((*MockMetaCleaner)(nil).ClearTable), dbName, tableName)
}

// MockIngestBinlogMetaer is a mock of IngestBinlogMetaer interface.
type MockIngestBinlogMetaer struct {
	ctrl     *gomock.Controller
	recorder *MockIngestBinlogMetaerMockRecorder
}

// MockIngestBinlogMetaerMockRecorder is the mock recorder for MockIngestBinlogMetaer.
type MockIngestBinlogMetaerMockRecorder struct {
	mock *MockIngestBinlogMetaer
}

// NewMockIngestBinlogMetaer creates a new mock instance.
func NewMockIngestBinlogMetaer(ctrl *gomock.Controller) *MockIngestBinlogMetaer {
	mock := &MockIngestBinlogMetaer{ctrl: ctrl}
	mock.recorder = &MockIngestBinlogMetaerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestBinlogMetaer) EXPECT() *MockIngestBinlogMetaerMockRecorder {
	return m.recorder
}

// GetBackendMap mocks base method.
func (m *MockIngestBinlogMetaer) GetBackendMap() (map[int64]*base.Backend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendMap")
	ret0, _ := ret[0].(map[int64]*base.Backend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendMap indicates an expected call of GetBackendMap.
func (mr *MockIngestBinlogMetaerMockRecorder) GetBackendMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendMap", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetBackendMap))
}

// GetIndexIdMap mocks base method.
func (m *MockIngestBinlogMetaer) GetIndexIdMap(tableId, partitionId int64) (map[int64]*IndexMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexIdMap", tableId, partitionId)
	ret0, _ := ret[0].(map[int64]*IndexMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexIdMap indicates an expected call of GetIndexIdMap.
func (mr *MockIngestBinlogMetaerMockRecorder) GetIndexIdMap(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexIdMap", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetIndexIdMap), tableId, partitionId)
}

// GetIndexNameMap mocks base method.
func (m *MockIngestBinlogMetaer) GetIndexNameMap(tableId, partitionId int64) (map[string]*IndexMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexNameMap", tableId, partitionId)
	ret0, _ := ret[0].(map[string]*IndexMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexNameMap indicates an expected call of GetIndexNameMap.
func (mr *MockIngestBinlogMetaerMockRecorder) GetIndexNameMap(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexNameMap", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetIndexNameMap), tableId, partitionId)
}

// GetPartitionIdByRange mocks base method.
func (m *MockIngestBinlogMetaer) GetPartitionIdByRange(tableId int64, partitionRange string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionIdByRange", tableId, partitionRange)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionIdByRange indicates an expected call of GetPartitionIdByRange.
func (mr *MockIngestBinlogMetaerMockRecorder) GetPartitionIdByRange(tableId, partitionRange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIdByRange", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetPartitionIdByRange), tableId, partitionRange)
}

// GetPartitionRangeMap mocks base method.
func (m *MockIngestBinlogMetaer) GetPartitionRangeMap(tableId int64) (map[string]*PartitionMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionRangeMap", tableId)
	ret0, _ := ret[0].(map[string]*PartitionMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionRangeMap indicates an expected call of GetPartitionRangeMap.
func (mr *MockIngestBinlogMetaerMockRecorder) GetPartitionRangeMap(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionRangeMap", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetPartitionRangeMap), tableId)
}

// GetTablets mocks base method.
func (m *MockIngestBinlogMetaer) GetTablets(tableId, partitionId, indexId int64) (*btree.Map[int64, *TabletMeta], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTablets", tableId, partitionId, indexId)
	ret0, _ := ret[0].(*btree.Map[int64, *TabletMeta])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTablets indicates an expected call of GetTablets.
func (mr *MockIngestBinlogMetaerMockRecorder) GetTablets(tableId, partitionId, indexId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTablets", reflect.TypeOf((*MockIngestBinlogMetaer)(nil).GetTablets), tableId, partitionId, indexId)
}

// MockMetaer is a mock of Metaer interface.
type MockMetaer struct {
	ctrl     *gomock.Controller
	recorder *MockMetaerMockRecorder
}

// MockMetaerMockRecorder is the mock recorder for MockMetaer.
type MockMetaerMockRecorder struct {
	mock *MockMetaer
}

// NewMockMetaer creates a new mock instance.
func NewMockMetaer(ctrl *gomock.Controller) *MockMetaer {
	mock := &MockMetaer{ctrl: ctrl}
	mock.recorder = &MockMetaerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaer) EXPECT() *MockMetaerMockRecorder {
	return m.recorder
}

// CheckBinlogFeature mocks base method.
func (m *MockMetaer) CheckBinlogFeature() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBinlogFeature")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBinlogFeature indicates an expected call of CheckBinlogFeature.
func (mr *MockMetaerMockRecorder) CheckBinlogFeature() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBinlogFeature", reflect.TypeOf((*MockMetaer)(nil).CheckBinlogFeature))
}

// ClearDB mocks base method.
func (m *MockMetaer) ClearDB(dbName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearDB", dbName)
}

// ClearDB indicates an expected call of ClearDB.
func (mr *MockMetaerMockRecorder) ClearDB(dbName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDB", reflect.TypeOf((*MockMetaer)(nil).ClearDB), dbName)
}

// ClearTable mocks base method.
func (m *MockMetaer) ClearTable(dbName, tableName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearTable", dbName, tableName)
}

// ClearTable indicates an expected call of ClearTable.
func (mr *MockMetaerMockRecorder) ClearTable(dbName, tableName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearTable", reflect.TypeOf((*MockMetaer)(nil).ClearTable), dbName, tableName)
}

// DbExec mocks base method.
func (m *MockMetaer) DbExec(sql string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DbExec", sql)
	ret0, _ := ret[0].(error)
	return ret0
}

// DbExec indicates an expected call of DbExec.
func (mr *MockMetaerMockRecorder) DbExec(sql any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbExec", reflect.TypeOf((*MockMetaer)(nil).DbExec), sql)
}

// DirtyGetTables mocks base method.
func (m *MockMetaer) DirtyGetTables() map[int64]*TableMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirtyGetTables")
	ret0, _ := ret[0].(map[int64]*TableMeta)
	return ret0
}

// DirtyGetTables indicates an expected call of DirtyGetTables.
func (mr *MockMetaerMockRecorder) DirtyGetTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirtyGetTables", reflect.TypeOf((*MockMetaer)(nil).DirtyGetTables))
}

// GetBackendId mocks base method.
func (m *MockMetaer) GetBackendId(host, portStr string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendId", host, portStr)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendId indicates an expected call of GetBackendId.
func (mr *MockMetaerMockRecorder) GetBackendId(host, portStr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendId", reflect.TypeOf((*MockMetaer)(nil).GetBackendId), host, portStr)
}

// GetBackendMap mocks base method.
func (m *MockMetaer) GetBackendMap() (map[int64]*base.Backend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendMap")
	ret0, _ := ret[0].(map[int64]*base.Backend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendMap indicates an expected call of GetBackendMap.
func (mr *MockMetaerMockRecorder) GetBackendMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendMap", reflect.TypeOf((*MockMetaer)(nil).GetBackendMap))
}

// GetBackends mocks base method.
func (m *MockMetaer) GetBackends() ([]*base.Backend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackends")
	ret0, _ := ret[0].([]*base.Backend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackends indicates an expected call of GetBackends.
func (mr *MockMetaerMockRecorder) GetBackends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackends", reflect.TypeOf((*MockMetaer)(nil).GetBackends))
}

// GetDbId mocks base method.
func (m *MockMetaer) GetDbId() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDbId")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDbId indicates an expected call of GetDbId.
func (mr *MockMetaerMockRecorder) GetDbId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDbId", reflect.TypeOf((*MockMetaer)(nil).GetDbId))
}

// GetFrontends mocks base method.
func (m *MockMetaer) GetFrontends() ([]*base.Frontend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontends")
	ret0, _ := ret[0].([]*base.Frontend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFrontends indicates an expected call of GetFrontends.
func (mr *MockMetaerMockRecorder) GetFrontends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontends", reflect.TypeOf((*MockMetaer)(nil).GetFrontends))
}

// GetFullTableName mocks base method.
func (m *MockMetaer) GetFullTableName(tableName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullTableName", tableName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFullTableName indicates an expected call of GetFullTableName.
func (mr *MockMetaerMockRecorder) GetFullTableName(tableName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullTableName", reflect.TypeOf((*MockMetaer)(nil).GetFullTableName), tableName)
}

// GetIndexIdMap mocks base method.
func (m *MockMetaer) GetIndexIdMap(tableId, partitionId int64) (map[int64]*IndexMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexIdMap", tableId, partitionId)
	ret0, _ := ret[0].(map[int64]*IndexMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexIdMap indicates an expected call of GetIndexIdMap.
func (mr *MockMetaerMockRecorder) GetIndexIdMap(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexIdMap", reflect.TypeOf((*MockMetaer)(nil).GetIndexIdMap), tableId, partitionId)
}

// GetIndexNameMap mocks base method.
func (m *MockMetaer) GetIndexNameMap(tableId, partitionId int64) (map[string]*IndexMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexNameMap", tableId, partitionId)
	ret0, _ := ret[0].(map[string]*IndexMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexNameMap indicates an expected call of GetIndexNameMap.
func (mr *MockMetaerMockRecorder) GetIndexNameMap(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexNameMap", reflect.TypeOf((*MockMetaer)(nil).GetIndexNameMap), tableId, partitionId)
}

// GetMasterToken mocks base method.
func (m *MockMetaer) GetMasterToken(rpcFactory rpc.IRpcFactory) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterToken", rpcFactory)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterToken indicates an expected call of GetMasterToken.
func (mr *MockMetaerMockRecorder) GetMasterToken(rpcFactory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterToken", reflect.TypeOf((*MockMetaer)(nil).GetMasterToken), rpcFactory)
}

// GetPartitionIdByName mocks base method.
func (m *MockMetaer) GetPartitionIdByName(tableId int64, partitionName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionIdByName", tableId, partitionName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionIdByName indicates an expected call of GetPartitionIdByName.
func (mr *MockMetaerMockRecorder) GetPartitionIdByName(tableId, partitionName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIdByName", reflect.TypeOf((*MockMetaer)(nil).GetPartitionIdByName), tableId, partitionName)
}

// GetPartitionIdByRange mocks base method.
func (m *MockMetaer) GetPartitionIdByRange(tableId int64, partitionRange string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionIdByRange", tableId, partitionRange)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionIdByRange indicates an expected call of GetPartitionIdByRange.
func (mr *MockMetaerMockRecorder) GetPartitionIdByRange(tableId, partitionRange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIdByRange", reflect.TypeOf((*MockMetaer)(nil).GetPartitionIdByRange), tableId, partitionRange)
}

// GetPartitionIdMap mocks base method.
func (m *MockMetaer) GetPartitionIdMap(tableId int64) (map[int64]*PartitionMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionIdMap", tableId)
	ret0, _ := ret[0].(map[int64]*PartitionMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionIdMap indicates an expected call of GetPartitionIdMap.
func (mr *MockMetaerMockRecorder) GetPartitionIdMap(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIdMap", reflect.TypeOf((*MockMetaer)(nil).GetPartitionIdMap), tableId)
}

// GetPartitionIds mocks base method.
func (m *MockMetaer) GetPartitionIds(tableName string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionIds", tableName)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionIds indicates an expected call of GetPartitionIds.
func (mr *MockMetaerMockRecorder) GetPartitionIds(tableName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIds", reflect.TypeOf((*MockMetaer)(nil).GetPartitionIds), tableName)
}

// GetPartitionName mocks base method.
func (m *MockMetaer) GetPartitionName(tableId, partitionId int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionName", tableId, partitionId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionName indicates an expected call of GetPartitionName.
func (mr *MockMetaerMockRecorder) GetPartitionName(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionName", reflect.TypeOf((*MockMetaer)(nil).GetPartitionName), tableId, partitionId)
}

// GetPartitionRange mocks base method.
func (m *MockMetaer) GetPartitionRange(tableId, partitionId int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionRange", tableId, partitionId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionRange indicates an expected call of GetPartitionRange.
func (mr *MockMetaerMockRecorder) GetPartitionRange(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionRange", reflect.TypeOf((*MockMetaer)(nil).GetPartitionRange), tableId, partitionId)
}

// GetPartitionRangeMap mocks base method.
func (m *MockMetaer) GetPartitionRangeMap(tableId int64) (map[string]*PartitionMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionRangeMap", tableId)
	ret0, _ := ret[0].(map[string]*PartitionMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionRangeMap indicates an expected call of GetPartitionRangeMap.
func (mr *MockMetaerMockRecorder) GetPartitionRangeMap(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionRangeMap", reflect.TypeOf((*MockMetaer)(nil).GetPartitionRangeMap), tableId)
}

// GetReplicas mocks base method.
func (m *MockMetaer) GetReplicas(tableId, partitionId int64) (*btree.Map[int64, *ReplicaMeta], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicas", tableId, partitionId)
	ret0, _ := ret[0].(*btree.Map[int64, *ReplicaMeta])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicas indicates an expected call of GetReplicas.
func (mr *MockMetaerMockRecorder) GetReplicas(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicas", reflect.TypeOf((*MockMetaer)(nil).GetReplicas), tableId, partitionId)
}

// GetTable mocks base method.
func (m *MockMetaer) GetTable(tableId int64) (*TableMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTable", tableId)
	ret0, _ := ret[0].(*TableMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTable indicates an expected call of GetTable.
func (mr *MockMetaerMockRecorder) GetTable(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTable", reflect.TypeOf((*MockMetaer)(nil).GetTable), tableId)
}

// GetTableId mocks base method.
func (m *MockMetaer) GetTableId(tableName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTableId", tableName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTableId indicates an expected call of GetTableId.
func (mr *MockMetaerMockRecorder) GetTableId(tableName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTableId", reflect.TypeOf((*MockMetaer)(nil).GetTableId), tableName)
}

// GetTableNameById mocks base method.
func (m *MockMetaer) GetTableNameById(tableId int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTableNameById", tableId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTableNameById indicates an expected call of GetTableNameById.
func (mr *MockMetaerMockRecorder) GetTableNameById(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTableNameById", reflect.TypeOf((*MockMetaer)(nil).GetTableNameById), tableId)
}

// GetTables mocks base method.
func (m *MockMetaer) GetTables() (map[int64]*TableMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTables")
	ret0, _ := ret[0].(map[int64]*TableMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTables indicates an expected call of GetTables.
func (mr *MockMetaerMockRecorder) GetTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockMetaer)(nil).GetTables))
}

// GetTablets mocks base method.
func (m *MockMetaer) GetTablets(tableId, partitionId, indexId int64) (*btree.Map[int64, *TabletMeta], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTablets", tableId, partitionId, indexId)
	ret0, _ := ret[0].(*btree.Map[int64, *TabletMeta])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTablets indicates an expected call of GetTablets.
func (mr *MockMetaerMockRecorder) GetTablets(tableId, partitionId, indexId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTablets", reflect.TypeOf((*MockMetaer)(nil).GetTablets), tableId, partitionId, indexId)
}

// UpdateBackends mocks base method.
func (m *MockMetaer) UpdateBackends() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackends")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBackends indicates an expected call of UpdateBackends.
func (mr *MockMetaerMockRecorder) UpdateBackends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackends", reflect.TypeOf((*MockMetaer)(nil).UpdateBackends))
}

// UpdateIndexes mocks base method.
func (m *MockMetaer) UpdateIndexes(tableId, partitionId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIndexes", tableId, partitionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIndexes indicates an expected call of UpdateIndexes.
func (mr *MockMetaerMockRecorder) UpdateIndexes(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIndexes", reflect.TypeOf((*MockMetaer)(nil).UpdateIndexes), tableId, partitionId)
}

// UpdatePartitions mocks base method.
func (m *MockMetaer) UpdatePartitions(tableId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePartitions", tableId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePartitions indicates an expected call of UpdatePartitions.
func (mr *MockMetaerMockRecorder) UpdatePartitions(tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartitions", reflect.TypeOf((*MockMetaer)(nil).UpdatePartitions), tableId)
}

// UpdateReplicas mocks base method.
func (m *MockMetaer) UpdateReplicas(tableId, partitionId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReplicas", tableId, partitionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicas indicates an expected call of UpdateReplicas.
func (mr *MockMetaerMockRecorder) UpdateReplicas(tableId, partitionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicas", reflect.TypeOf((*MockMetaer)(nil).UpdateReplicas), tableId, partitionId)
}

// UpdateTable mocks base method.
func (m *MockMetaer) UpdateTable(tableName string, tableId int64) (*TableMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTable", tableName, tableId)
	ret0, _ := ret[0].(*TableMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTable indicates an expected call of UpdateTable.
func (mr *MockMetaerMockRecorder) UpdateTable(tableName, tableId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTable", reflect.TypeOf((*MockMetaer)(nil).UpdateTable), tableName, tableId)
}

// UpdateToken mocks base method.
func (m *MockMetaer) UpdateToken(rpcFactory rpc.IRpcFactory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", rpcFactory)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockMetaerMockRecorder) UpdateToken(rpcFactory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockMetaer)(nil).UpdateToken), rpcFactory)
}
