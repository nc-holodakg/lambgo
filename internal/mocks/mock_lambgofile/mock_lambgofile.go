// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JosiahWitt/lambgo/internal/lambgofile (interfaces: LoaderAPI)

// Package mock_lambgofile is a generated GoMock package.
package mock_lambgofile

import (
	reflect "reflect"

	lambgofile "github.com/JosiahWitt/lambgo/internal/lambgofile"
	gomock "github.com/golang/mock/gomock"
)

// MockLoaderAPI is a mock of LoaderAPI interface.
type MockLoaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderAPIMockRecorder
}

// MockLoaderAPIMockRecorder is the mock recorder for MockLoaderAPI.
type MockLoaderAPIMockRecorder struct {
	mock *MockLoaderAPI
}

// NewMockLoaderAPI creates a new mock instance.
func NewMockLoaderAPI(ctrl *gomock.Controller) *MockLoaderAPI {
	mock := &MockLoaderAPI{ctrl: ctrl}
	mock.recorder = &MockLoaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoaderAPI) EXPECT() *MockLoaderAPIMockRecorder {
	return m.recorder
}

// LoadConfig mocks base method.
func (m *MockLoaderAPI) LoadConfig(arg0 string) (*lambgofile.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadConfig", arg0)
	ret0, _ := ret[0].(*lambgofile.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadConfig indicates an expected call of LoadConfig.
func (mr *MockLoaderAPIMockRecorder) LoadConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadConfig", reflect.TypeOf((*MockLoaderAPI)(nil).LoadConfig), arg0)
}

// NEW creates a MockLoaderAPI.
func (*MockLoaderAPI) NEW(ctrl *gomock.Controller) *MockLoaderAPI {
	return NewMockLoaderAPI(ctrl)
}
