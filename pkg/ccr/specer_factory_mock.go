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
// Source: ccr/base/specer_factory.go
//
// Generated by this command:
//
//	mockgen -source=ccr/base/specer_factory.go -destination=ccr/specer_factory_mock.go -package=ccr
//
// Package ccr is a generated GoMock package.
package ccr

import (
	reflect "reflect"

	base "github.com/selectdb/ccr_syncer/pkg/ccr/base"
	gomock "go.uber.org/mock/gomock"
)

// MockSpecerFactory is a mock of SpecerFactory interface.
type MockSpecerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSpecerFactoryMockRecorder
}

// MockSpecerFactoryMockRecorder is the mock recorder for MockSpecerFactory.
type MockSpecerFactoryMockRecorder struct {
	mock *MockSpecerFactory
}

// NewMockSpecerFactory creates a new mock instance.
func NewMockSpecerFactory(ctrl *gomock.Controller) *MockSpecerFactory {
	mock := &MockSpecerFactory{ctrl: ctrl}
	mock.recorder = &MockSpecerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpecerFactory) EXPECT() *MockSpecerFactoryMockRecorder {
	return m.recorder
}

// NewSpecer mocks base method.
func (m *MockSpecerFactory) NewSpecer(tableSpec *base.Spec) base.Specer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSpecer", tableSpec)
	ret0, _ := ret[0].(base.Specer)
	return ret0
}

// NewSpecer indicates an expected call of NewSpecer.
func (mr *MockSpecerFactoryMockRecorder) NewSpecer(tableSpec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSpecer", reflect.TypeOf((*MockSpecerFactory)(nil).NewSpecer), tableSpec)
}
