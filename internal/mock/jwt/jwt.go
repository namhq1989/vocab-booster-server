// Code generated by MockGen. DO NOT EDIT.
// Source: internal/utils/jwt/jwt.go
//
// Generated by this command:
//
//	mockgen -source=internal/utils/jwt/jwt.go -destination=internal/mock/jwt/jwt.go -package=mockjwt
//

// Package mockjwt is a generated GoMock package.
package mockjwt

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	gomock "go.uber.org/mock/gomock"
)

// MockJWTInterface is a mock of JWTInterface interface.
type MockJWTInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJWTInterfaceMockRecorder
}

// MockJWTInterfaceMockRecorder is the mock recorder for MockJWTInterface.
type MockJWTInterfaceMockRecorder struct {
	mock *MockJWTInterface
}

// NewMockJWTInterface creates a new mock instance.
func NewMockJWTInterface(ctrl *gomock.Controller) *MockJWTInterface {
	mock := &MockJWTInterface{ctrl: ctrl}
	mock.recorder = &MockJWTInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTInterface) EXPECT() *MockJWTInterfaceMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockJWTInterface) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockJWTInterfaceMockRecorder) GenerateAccessToken(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockJWTInterface)(nil).GenerateAccessToken), ctx, userID)
}

// GenerateTokens mocks base method.
func (m *MockJWTInterface) GenerateTokens(ctx *appcontext.AppContext, userID string) (*appjwt.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokens", ctx, userID)
	ret0, _ := ret[0].(*appjwt.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokens indicates an expected call of GenerateTokens.
func (mr *MockJWTInterfaceMockRecorder) GenerateTokens(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokens", reflect.TypeOf((*MockJWTInterface)(nil).GenerateTokens), ctx, userID)
}

// ParseAccessToken mocks base method.
func (m *MockJWTInterface) ParseAccessToken(ctx *appcontext.AppContext, token string) (*appjwt.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseAccessToken", ctx, token)
	ret0, _ := ret[0].(*appjwt.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAccessToken indicates an expected call of ParseAccessToken.
func (mr *MockJWTInterfaceMockRecorder) ParseAccessToken(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAccessToken", reflect.TypeOf((*MockJWTInterface)(nil).ParseAccessToken), ctx, token)
}
