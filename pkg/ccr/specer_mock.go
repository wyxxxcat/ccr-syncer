// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License
// Code generated by MockGen. DO NOT EDIT.
// Source: ccr/base/specer.go
//
// Generated by this command:
//
//	mockgen -source=ccr/base/specer.go -destination=ccr/specer_mock.go -package=ccr
//
// Package ccr is a generated GoMock package.
package ccr

import (
	sql "database/sql"
	reflect "reflect"

	base "github.com/selectdb/ccr_syncer/pkg/ccr/base"
	utils "github.com/selectdb/ccr_syncer/pkg/utils"
	gomock "go.uber.org/mock/gomock"
)

// MockSpecer is a mock of Specer interface.
type MockSpecer struct {
	ctrl     *gomock.Controller
	recorder *MockSpecerMockRecorder
}

// MockSpecerMockRecorder is the mock recorder for MockSpecer.
type MockSpecerMockRecorder struct {
	mock *MockSpecer
}

// NewMockSpecer creates a new mock instance.
func NewMockSpecer(ctrl *gomock.Controller) *MockSpecer {
	mock := &MockSpecer{ctrl: ctrl}
	mock.recorder = &MockSpecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpecer) EXPECT() *MockSpecerMockRecorder {
	return m.recorder
}

// CheckDatabaseExists mocks base method.
func (m *MockSpecer) CheckDatabaseExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDatabaseExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDatabaseExists indicates an expected call of CheckDatabaseExists.
func (mr *MockSpecerMockRecorder) CheckDatabaseExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDatabaseExists", reflect.TypeOf((*MockSpecer)(nil).CheckDatabaseExists))
}

// CheckRestoreFinished mocks base method.
func (m *MockSpecer) CheckRestoreFinished(snapshotName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRestoreFinished", snapshotName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRestoreFinished indicates an expected call of CheckRestoreFinished.
func (mr *MockSpecerMockRecorder) CheckRestoreFinished(snapshotName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRestoreFinished", reflect.TypeOf((*MockSpecer)(nil).CheckRestoreFinished), snapshotName)
}

// CheckTableExists mocks base method.
func (m *MockSpecer) CheckTableExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTableExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTableExists indicates an expected call of CheckTableExists.
func (mr *MockSpecerMockRecorder) CheckTableExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTableExists", reflect.TypeOf((*MockSpecer)(nil).CheckTableExists))
}

// ClearDB mocks base method.
func (m *MockSpecer) ClearDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearDB indicates an expected call of ClearDB.
func (mr *MockSpecerMockRecorder) ClearDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDB", reflect.TypeOf((*MockSpecer)(nil).ClearDB))
}

// Connect mocks base method.
func (m *MockSpecer) Connect() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockSpecerMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockSpecer)(nil).Connect))
}

// ConnectDB mocks base method.
func (m *MockSpecer) ConnectDB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectDB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectDB indicates an expected call of ConnectDB.
func (mr *MockSpecerMockRecorder) ConnectDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectDB", reflect.TypeOf((*MockSpecer)(nil).ConnectDB))
}

// CreateDatabase mocks base method.
func (m *MockSpecer) CreateDatabase() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabase")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDatabase indicates an expected call of CreateDatabase.
func (mr *MockSpecerMockRecorder) CreateDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabase", reflect.TypeOf((*MockSpecer)(nil).CreateDatabase))
}

// CreateSnapshotAndWaitForDone mocks base method.
func (m *MockSpecer) CreateSnapshotAndWaitForDone(tables []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshotAndWaitForDone", tables)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshotAndWaitForDone indicates an expected call of CreateSnapshotAndWaitForDone.
func (mr *MockSpecerMockRecorder) CreateSnapshotAndWaitForDone(tables any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshotAndWaitForDone", reflect.TypeOf((*MockSpecer)(nil).CreateSnapshotAndWaitForDone), tables)
}

// CreateTable mocks base method.
func (m *MockSpecer) CreateTable(stmt string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", stmt)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockSpecerMockRecorder) CreateTable(stmt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockSpecer)(nil).CreateTable), stmt)
}

// DbExec mocks base method.
func (m *MockSpecer) DbExec(sql string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DbExec", sql)
	ret0, _ := ret[0].(error)
	return ret0
}

// DbExec indicates an expected call of DbExec.
func (mr *MockSpecerMockRecorder) DbExec(sql any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbExec", reflect.TypeOf((*MockSpecer)(nil).DbExec), sql)
}

// Exec mocks base method.
func (m *MockSpecer) Exec(sql string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", sql)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockSpecerMockRecorder) Exec(sql any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockSpecer)(nil).Exec), sql)
}

// GetAllTables mocks base method.
func (m *MockSpecer) GetAllTables() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTables")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTables indicates an expected call of GetAllTables.
func (mr *MockSpecerMockRecorder) GetAllTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTables", reflect.TypeOf((*MockSpecer)(nil).GetAllTables))
}

// IsDatabaseEnableBinlog mocks base method.
func (m *MockSpecer) IsDatabaseEnableBinlog() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDatabaseEnableBinlog")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDatabaseEnableBinlog indicates an expected call of IsDatabaseEnableBinlog.
func (mr *MockSpecerMockRecorder) IsDatabaseEnableBinlog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDatabaseEnableBinlog", reflect.TypeOf((*MockSpecer)(nil).IsDatabaseEnableBinlog))
}

// IsTableEnableBinlog mocks base method.
func (m *MockSpecer) IsTableEnableBinlog() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTableEnableBinlog")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsTableEnableBinlog indicates an expected call of IsTableEnableBinlog.
func (mr *MockSpecerMockRecorder) IsTableEnableBinlog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTableEnableBinlog", reflect.TypeOf((*MockSpecer)(nil).IsTableEnableBinlog))
}

// Notify mocks base method.
func (m *MockSpecer) Notify(arg0 base.SpecEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", arg0)
}

// Notify indicates an expected call of Notify.
func (mr *MockSpecerMockRecorder) Notify(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockSpecer)(nil).Notify), arg0)
}

// Register mocks base method.
func (m *MockSpecer) Register(arg0 utils.Observer[base.SpecEvent]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register.
func (mr *MockSpecerMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSpecer)(nil).Register), arg0)
}

// Unregister mocks base method.
func (m *MockSpecer) Unregister(arg0 utils.Observer[base.SpecEvent]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unregister", arg0)
}

// Unregister indicates an expected call of Unregister.
func (mr *MockSpecerMockRecorder) Unregister(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockSpecer)(nil).Unregister), arg0)
}

// Valid mocks base method.
func (m *MockSpecer) Valid() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid")
	ret0, _ := ret[0].(error)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockSpecerMockRecorder) Valid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockSpecer)(nil).Valid))
}

// WaitTransactionDone mocks base method.
func (m *MockSpecer) WaitTransactionDone(txnId int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitTransactionDone", txnId)
}

// WaitTransactionDone indicates an expected call of WaitTransactionDone.
func (mr *MockSpecerMockRecorder) WaitTransactionDone(txnId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitTransactionDone", reflect.TypeOf((*MockSpecer)(nil).WaitTransactionDone), txnId)
}
