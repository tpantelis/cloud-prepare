// SPDX-License-Identifier: Apache-2.0
//
// Copyright Contributors to the Submariner project.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./machinesets.go

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MockMachineSetDeployer is a mock of MachineSetDeployer interface.
type MockMachineSetDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockMachineSetDeployerMockRecorder
}

// MockMachineSetDeployerMockRecorder is the mock recorder for MockMachineSetDeployer.
type MockMachineSetDeployerMockRecorder struct {
	mock *MockMachineSetDeployer
}

// NewMockMachineSetDeployer creates a new mock instance.
func NewMockMachineSetDeployer(ctrl *gomock.Controller) *MockMachineSetDeployer {
	mock := &MockMachineSetDeployer{ctrl: ctrl}
	mock.recorder = &MockMachineSetDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineSetDeployer) EXPECT() *MockMachineSetDeployerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMachineSetDeployer) Delete(machineSet *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", machineSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMachineSetDeployerMockRecorder) Delete(machineSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMachineSetDeployer)(nil).Delete), machineSet)
}

// DeleteByName mocks base method.
func (m *MockMachineSetDeployer) DeleteByName(name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockMachineSetDeployerMockRecorder) DeleteByName(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockMachineSetDeployer)(nil).DeleteByName), name, namespace)
}

// Deploy mocks base method.
func (m *MockMachineSetDeployer) Deploy(machineSet *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", machineSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockMachineSetDeployerMockRecorder) Deploy(machineSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockMachineSetDeployer)(nil).Deploy), machineSet)
}

// GetWorkerNodeImage mocks base method.
func (m *MockMachineSetDeployer) GetWorkerNodeImage(machineSet *unstructured.Unstructured, infraID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkerNodeImage", machineSet, infraID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerNodeImage indicates an expected call of GetWorkerNodeImage.
func (mr *MockMachineSetDeployerMockRecorder) GetWorkerNodeImage(machineSet, infraID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerNodeImage", reflect.TypeOf((*MockMachineSetDeployer)(nil).GetWorkerNodeImage), machineSet, infraID)
}

// List mocks base method.
func (m *MockMachineSetDeployer) List() ([]unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMachineSetDeployerMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMachineSetDeployer)(nil).List))
}
