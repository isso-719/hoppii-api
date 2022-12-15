// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package announcement is a generated GoMock package.
package announcement

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFAnnouncementService is a mock of IFAnnouncementService interface.
type MockIFAnnouncementService struct {
	ctrl     *gomock.Controller
	recorder *MockIFAnnouncementServiceMockRecorder
}

// MockIFAnnouncementServiceMockRecorder is the mock recorder for MockIFAnnouncementService.
type MockIFAnnouncementServiceMockRecorder struct {
	mock *MockIFAnnouncementService
}

// NewMockIFAnnouncementService creates a new mock instance.
func NewMockIFAnnouncementService(ctrl *gomock.Controller) *MockIFAnnouncementService {
	mock := &MockIFAnnouncementService{ctrl: ctrl}
	mock.recorder = &MockIFAnnouncementServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFAnnouncementService) EXPECT() *MockIFAnnouncementServiceMockRecorder {
	return m.recorder
}

// User mocks base method.
func (m *MockIFAnnouncementService) User(input *AnnouncementUserInput) (*AnnouncementUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", input)
	ret0, _ := ret[0].(*AnnouncementUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User.
func (mr *MockIFAnnouncementServiceMockRecorder) User(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockIFAnnouncementService)(nil).User), input)
}

// MockifPrivateAnnouncementService is a mock of ifPrivateAnnouncementService interface.
type MockifPrivateAnnouncementService struct {
	ctrl     *gomock.Controller
	recorder *MockifPrivateAnnouncementServiceMockRecorder
}

// MockifPrivateAnnouncementServiceMockRecorder is the mock recorder for MockifPrivateAnnouncementService.
type MockifPrivateAnnouncementServiceMockRecorder struct {
	mock *MockifPrivateAnnouncementService
}

// NewMockifPrivateAnnouncementService creates a new mock instance.
func NewMockifPrivateAnnouncementService(ctrl *gomock.Controller) *MockifPrivateAnnouncementService {
	mock := &MockifPrivateAnnouncementService{ctrl: ctrl}
	mock.recorder = &MockifPrivateAnnouncementServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockifPrivateAnnouncementService) EXPECT() *MockifPrivateAnnouncementServiceMockRecorder {
	return m.recorder
}
