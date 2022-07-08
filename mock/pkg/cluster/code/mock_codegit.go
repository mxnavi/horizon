// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cluster/code/codegit.go

// Package mock_code is a generated GoMock package.
package mock_code

import (
	context "context"
	reflect "reflect"

	code "g.hz.netease.com/horizon/pkg/cluster/code"
	gomock "github.com/golang/mock/gomock"
)

// MockGitGetter is a mock of GitGetter interface.
type MockGitGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGitGetterMockRecorder
}

// MockGitGetterMockRecorder is the mock recorder for MockGitGetter.
type MockGitGetterMockRecorder struct {
	mock *MockGitGetter
}

// NewMockGitGetter creates a new mock instance.
func NewMockGitGetter(ctrl *gomock.Controller) *MockGitGetter {
	mock := &MockGitGetter{ctrl: ctrl}
	mock.recorder = &MockGitGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitGetter) EXPECT() *MockGitGetterMockRecorder {
	return m.recorder
}

// GetCommit mocks base method.
func (m *MockGitGetter) GetCommit(ctx context.Context, gitURL, refType, ref string) (*code.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, gitURL, refType, ref)
	ret0, _ := ret[0].(*code.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockGitGetterMockRecorder) GetCommit(ctx, gitURL, refType, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockGitGetter)(nil).GetCommit), ctx, gitURL, refType, ref)
}

// ListBranch mocks base method.
func (m *MockGitGetter) ListBranch(ctx context.Context, gitURL string, params *code.SearchParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranch", ctx, gitURL, params)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranch indicates an expected call of ListBranch.
func (mr *MockGitGetterMockRecorder) ListBranch(ctx, gitURL, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranch", reflect.TypeOf((*MockGitGetter)(nil).ListBranch), ctx, gitURL, params)
}

// ListTag mocks base method.
func (m *MockGitGetter) ListTag(ctx context.Context, gitURL string, params *code.SearchParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTag", ctx, gitURL, params)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTag indicates an expected call of ListTag.
func (mr *MockGitGetterMockRecorder) ListTag(ctx, gitURL, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTag", reflect.TypeOf((*MockGitGetter)(nil).ListTag), ctx, gitURL, params)
}
