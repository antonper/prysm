// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: ValidatorServiceClient,ValidatorService_WaitForActivationClient,ValidatorService_WaitForChainStartClient)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockValidatorServiceClient is a mock of ValidatorServiceClient interface
type MockValidatorServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorServiceClientMockRecorder
}

// MockValidatorServiceClientMockRecorder is the mock recorder for MockValidatorServiceClient
type MockValidatorServiceClientMockRecorder struct {
	mock *MockValidatorServiceClient
}

// NewMockValidatorServiceClient creates a new mock instance
func NewMockValidatorServiceClient(ctrl *gomock.Controller) *MockValidatorServiceClient {
	mock := &MockValidatorServiceClient{ctrl: ctrl}
	mock.recorder = &MockValidatorServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorServiceClient) EXPECT() *MockValidatorServiceClientMockRecorder {
	return m.recorder
}

// CanonicalHead mocks base method
func (m *MockValidatorServiceClient) CanonicalHead(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*v1alpha1.BeaconBlock, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanonicalHead", varargs...)
	ret0, _ := ret[0].(*v1alpha1.BeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanonicalHead indicates an expected call of CanonicalHead
func (mr *MockValidatorServiceClientMockRecorder) CanonicalHead(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanonicalHead", reflect.TypeOf((*MockValidatorServiceClient)(nil).CanonicalHead), varargs...)
}

// CommitteeAssignment mocks base method
func (m *MockValidatorServiceClient) CommitteeAssignment(arg0 context.Context, arg1 *v1.AssignmentRequest, arg2 ...grpc.CallOption) (*v1.AssignmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitteeAssignment", varargs...)
	ret0, _ := ret[0].(*v1.AssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitteeAssignment indicates an expected call of CommitteeAssignment
func (mr *MockValidatorServiceClientMockRecorder) CommitteeAssignment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitteeAssignment", reflect.TypeOf((*MockValidatorServiceClient)(nil).CommitteeAssignment), varargs...)
}

// DomainData mocks base method
func (m *MockValidatorServiceClient) DomainData(arg0 context.Context, arg1 *v1.DomainRequest, arg2 ...grpc.CallOption) (*v1.DomainResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DomainData", varargs...)
	ret0, _ := ret[0].(*v1.DomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainData indicates an expected call of DomainData
func (mr *MockValidatorServiceClientMockRecorder) DomainData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainData", reflect.TypeOf((*MockValidatorServiceClient)(nil).DomainData), varargs...)
}

// ExitedValidators mocks base method
func (m *MockValidatorServiceClient) ExitedValidators(arg0 context.Context, arg1 *v1.ExitedValidatorsRequest, arg2 ...grpc.CallOption) (*v1.ExitedValidatorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExitedValidators", varargs...)
	ret0, _ := ret[0].(*v1.ExitedValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExitedValidators indicates an expected call of ExitedValidators
func (mr *MockValidatorServiceClientMockRecorder) ExitedValidators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitedValidators", reflect.TypeOf((*MockValidatorServiceClient)(nil).ExitedValidators), varargs...)
}

// ValidatorIndex mocks base method
func (m *MockValidatorServiceClient) ValidatorIndex(arg0 context.Context, arg1 *v1.ValidatorIndexRequest, arg2 ...grpc.CallOption) (*v1.ValidatorIndexResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidatorIndex", varargs...)
	ret0, _ := ret[0].(*v1.ValidatorIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorIndex indicates an expected call of ValidatorIndex
func (mr *MockValidatorServiceClientMockRecorder) ValidatorIndex(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorIndex", reflect.TypeOf((*MockValidatorServiceClient)(nil).ValidatorIndex), varargs...)
}

// ValidatorPerformance mocks base method
func (m *MockValidatorServiceClient) ValidatorPerformance(arg0 context.Context, arg1 *v1.ValidatorPerformanceRequest, arg2 ...grpc.CallOption) (*v1.ValidatorPerformanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidatorPerformance", varargs...)
	ret0, _ := ret[0].(*v1.ValidatorPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorPerformance indicates an expected call of ValidatorPerformance
func (mr *MockValidatorServiceClientMockRecorder) ValidatorPerformance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorPerformance", reflect.TypeOf((*MockValidatorServiceClient)(nil).ValidatorPerformance), varargs...)
}

// ValidatorStatus mocks base method
func (m *MockValidatorServiceClient) ValidatorStatus(arg0 context.Context, arg1 *v1.ValidatorIndexRequest, arg2 ...grpc.CallOption) (*v1.ValidatorStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidatorStatus", varargs...)
	ret0, _ := ret[0].(*v1.ValidatorStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorStatus indicates an expected call of ValidatorStatus
func (mr *MockValidatorServiceClientMockRecorder) ValidatorStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorStatus", reflect.TypeOf((*MockValidatorServiceClient)(nil).ValidatorStatus), varargs...)
}

// ProposeExit mocks base method
func (m *MockValidatorServiceClient) ProposeExit(arg0 context.Context, arg1 *v1alpha1.VoluntaryExit, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProposeExit", varargs)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeExit indicates an expected call of ProposeExit
func (mr *MockValidatorServiceClientMockRecorder) ProposeExit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeExit", reflect.TypeOf((*MockValidatorServiceClient)(nil).ProposeExit), varargs...)
}

// WaitForActivation mocks base method
func (m *MockValidatorServiceClient) WaitForActivation(arg0 context.Context, arg1 *v1.ValidatorActivationRequest, arg2 ...grpc.CallOption) (v1.ValidatorService_WaitForActivationClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForActivation", varargs...)
	ret0, _ := ret[0].(v1.ValidatorService_WaitForActivationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForActivation indicates an expected call of WaitForActivation
func (mr *MockValidatorServiceClientMockRecorder) WaitForActivation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForActivation", reflect.TypeOf((*MockValidatorServiceClient)(nil).WaitForActivation), varargs...)
}

// WaitForChainStart mocks base method
func (m *MockValidatorServiceClient) WaitForChainStart(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (v1.ValidatorService_WaitForChainStartClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForChainStart", varargs...)
	ret0, _ := ret[0].(v1.ValidatorService_WaitForChainStartClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForChainStart indicates an expected call of WaitForChainStart
func (mr *MockValidatorServiceClientMockRecorder) WaitForChainStart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForChainStart", reflect.TypeOf((*MockValidatorServiceClient)(nil).WaitForChainStart), varargs...)
}

// MockValidatorService_WaitForActivationClient is a mock of ValidatorService_WaitForActivationClient interface
type MockValidatorService_WaitForActivationClient struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorService_WaitForActivationClientMockRecorder
}

// MockValidatorService_WaitForActivationClientMockRecorder is the mock recorder for MockValidatorService_WaitForActivationClient
type MockValidatorService_WaitForActivationClientMockRecorder struct {
	mock *MockValidatorService_WaitForActivationClient
}

// NewMockValidatorService_WaitForActivationClient creates a new mock instance
func NewMockValidatorService_WaitForActivationClient(ctrl *gomock.Controller) *MockValidatorService_WaitForActivationClient {
	mock := &MockValidatorService_WaitForActivationClient{ctrl: ctrl}
	mock.recorder = &MockValidatorService_WaitForActivationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorService_WaitForActivationClient) EXPECT() *MockValidatorService_WaitForActivationClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockValidatorService_WaitForActivationClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockValidatorService_WaitForActivationClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).Context))
}

// Header mocks base method
func (m *MockValidatorService_WaitForActivationClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).Header))
}

// Recv mocks base method
func (m *MockValidatorService_WaitForActivationClient) Recv() (*v1.ValidatorActivationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.ValidatorActivationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockValidatorService_WaitForActivationClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockValidatorService_WaitForActivationClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockValidatorService_WaitForActivationClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockValidatorService_WaitForActivationClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockValidatorService_WaitForActivationClient)(nil).Trailer))
}

// MockValidatorService_WaitForChainStartClient is a mock of ValidatorService_WaitForChainStartClient interface
type MockValidatorService_WaitForChainStartClient struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorService_WaitForChainStartClientMockRecorder
}

// MockValidatorService_WaitForChainStartClientMockRecorder is the mock recorder for MockValidatorService_WaitForChainStartClient
type MockValidatorService_WaitForChainStartClientMockRecorder struct {
	mock *MockValidatorService_WaitForChainStartClient
}

// NewMockValidatorService_WaitForChainStartClient creates a new mock instance
func NewMockValidatorService_WaitForChainStartClient(ctrl *gomock.Controller) *MockValidatorService_WaitForChainStartClient {
	mock := &MockValidatorService_WaitForChainStartClient{ctrl: ctrl}
	mock.recorder = &MockValidatorService_WaitForChainStartClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorService_WaitForChainStartClient) EXPECT() *MockValidatorService_WaitForChainStartClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockValidatorService_WaitForChainStartClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockValidatorService_WaitForChainStartClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).Context))
}

// Header mocks base method
func (m *MockValidatorService_WaitForChainStartClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).Header))
}

// Recv mocks base method
func (m *MockValidatorService_WaitForChainStartClient) Recv() (*v1.ChainStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.ChainStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockValidatorService_WaitForChainStartClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockValidatorService_WaitForChainStartClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockValidatorService_WaitForChainStartClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockValidatorService_WaitForChainStartClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockValidatorService_WaitForChainStartClient)(nil).Trailer))
}