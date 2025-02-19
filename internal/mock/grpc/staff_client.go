// Code generated by MockGen. DO NOT EDIT.
// Source: internal/genproto/staffpb/hub_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/genproto/staffpb/hub_grpc.pb.go -destination=internal/mock/grpc/staff_client.go -package=mockgrpc
//

// Package mockgrpc is a generated GoMock package.
package mockgrpc

import (
	context "context"
	reflect "reflect"

	staffpb "github.com/namhq1989/vocab-booster-server-admin/internal/genproto/staffpb"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStaffServiceClient is a mock of StaffServiceClient interface.
type MockStaffServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockStaffServiceClientMockRecorder
}

// MockStaffServiceClientMockRecorder is the mock recorder for MockStaffServiceClient.
type MockStaffServiceClientMockRecorder struct {
	mock *MockStaffServiceClient
}

// NewMockStaffServiceClient creates a new mock instance.
func NewMockStaffServiceClient(ctrl *gomock.Controller) *MockStaffServiceClient {
	mock := &MockStaffServiceClient{ctrl: ctrl}
	mock.recorder = &MockStaffServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaffServiceClient) EXPECT() *MockStaffServiceClientMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockStaffServiceClient) FindUserByEmail(ctx context.Context, in *staffpb.FindStaffByEmailRequest, opts ...grpc.CallOption) (*staffpb.FindStaffByEmailResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUserByEmail", varargs...)
	ret0, _ := ret[0].(*staffpb.FindStaffByEmailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockStaffServiceClientMockRecorder) FindUserByEmail(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockStaffServiceClient)(nil).FindUserByEmail), varargs...)
}

// FindUserByID mocks base method.
func (m *MockStaffServiceClient) FindUserByID(ctx context.Context, in *staffpb.FindStaffByIDRequest, opts ...grpc.CallOption) (*staffpb.FindStaffByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUserByID", varargs...)
	ret0, _ := ret[0].(*staffpb.FindStaffByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockStaffServiceClientMockRecorder) FindUserByID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockStaffServiceClient)(nil).FindUserByID), varargs...)
}

// MockStaffServiceServer is a mock of StaffServiceServer interface.
type MockStaffServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockStaffServiceServerMockRecorder
}

// MockStaffServiceServerMockRecorder is the mock recorder for MockStaffServiceServer.
type MockStaffServiceServerMockRecorder struct {
	mock *MockStaffServiceServer
}

// NewMockStaffServiceServer creates a new mock instance.
func NewMockStaffServiceServer(ctrl *gomock.Controller) *MockStaffServiceServer {
	mock := &MockStaffServiceServer{ctrl: ctrl}
	mock.recorder = &MockStaffServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaffServiceServer) EXPECT() *MockStaffServiceServerMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockStaffServiceServer) FindUserByEmail(arg0 context.Context, arg1 *staffpb.FindStaffByEmailRequest) (*staffpb.FindStaffByEmailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*staffpb.FindStaffByEmailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockStaffServiceServerMockRecorder) FindUserByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockStaffServiceServer)(nil).FindUserByEmail), arg0, arg1)
}

// FindUserByID mocks base method.
func (m *MockStaffServiceServer) FindUserByID(arg0 context.Context, arg1 *staffpb.FindStaffByIDRequest) (*staffpb.FindStaffByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0, arg1)
	ret0, _ := ret[0].(*staffpb.FindStaffByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockStaffServiceServerMockRecorder) FindUserByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockStaffServiceServer)(nil).FindUserByID), arg0, arg1)
}

// MockUnsafeStaffServiceServer is a mock of UnsafeStaffServiceServer interface.
type MockUnsafeStaffServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeStaffServiceServerMockRecorder
}

// MockUnsafeStaffServiceServerMockRecorder is the mock recorder for MockUnsafeStaffServiceServer.
type MockUnsafeStaffServiceServerMockRecorder struct {
	mock *MockUnsafeStaffServiceServer
}

// NewMockUnsafeStaffServiceServer creates a new mock instance.
func NewMockUnsafeStaffServiceServer(ctrl *gomock.Controller) *MockUnsafeStaffServiceServer {
	mock := &MockUnsafeStaffServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeStaffServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeStaffServiceServer) EXPECT() *MockUnsafeStaffServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedStaffServiceServer mocks base method.
func (m *MockUnsafeStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStaffServiceServer")
}

// mustEmbedUnimplementedStaffServiceServer indicates an expected call of mustEmbedUnimplementedStaffServiceServer.
func (mr *MockUnsafeStaffServiceServerMockRecorder) mustEmbedUnimplementedStaffServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStaffServiceServer", reflect.TypeOf((*MockUnsafeStaffServiceServer)(nil).mustEmbedUnimplementedStaffServiceServer))
}
