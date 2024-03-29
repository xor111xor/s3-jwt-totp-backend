// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/model.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/model.go -destination=internal/domain/mocks/model.go
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	domain "github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockRepoDB is a mock of RepoDB interface.
type MockRepoDB struct {
	ctrl     *gomock.Controller
	recorder *MockRepoDBMockRecorder
}

// MockRepoDBMockRecorder is the mock recorder for MockRepoDB.
type MockRepoDBMockRecorder struct {
	mock *MockRepoDB
}

// NewMockRepoDB creates a new mock instance.
func NewMockRepoDB(ctrl *gomock.Controller) *MockRepoDB {
	mock := &MockRepoDB{ctrl: ctrl}
	mock.recorder = &MockRepoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoDB) EXPECT() *MockRepoDBMockRecorder {
	return m.recorder
}

// FileDelete mocks base method.
func (m *MockRepoDB) FileDelete(file domain.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileDelete", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileDelete indicates an expected call of FileDelete.
func (mr *MockRepoDBMockRecorder) FileDelete(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileDelete", reflect.TypeOf((*MockRepoDB)(nil).FileDelete), file)
}

// FileGetByID mocks base method.
func (m *MockRepoDB) FileGetByID(arg0 string) (domain.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileGetByID", arg0)
	ret0, _ := ret[0].(domain.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileGetByID indicates an expected call of FileGetByID.
func (mr *MockRepoDBMockRecorder) FileGetByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileGetByID", reflect.TypeOf((*MockRepoDB)(nil).FileGetByID), arg0)
}

// FileInsert mocks base method.
func (m *MockRepoDB) FileInsert(file domain.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileInsert", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileInsert indicates an expected call of FileInsert.
func (mr *MockRepoDBMockRecorder) FileInsert(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileInsert", reflect.TypeOf((*MockRepoDB)(nil).FileInsert), file)
}

// FilesGetByUserID mocks base method.
func (m *MockRepoDB) FilesGetByUserID(arg0 string) (*[]domain.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesGetByUserID", arg0)
	ret0, _ := ret[0].(*[]domain.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesGetByUserID indicates an expected call of FilesGetByUserID.
func (mr *MockRepoDBMockRecorder) FilesGetByUserID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesGetByUserID", reflect.TypeOf((*MockRepoDB)(nil).FilesGetByUserID), arg0)
}

// UserAddOnRegistration mocks base method.
func (m *MockRepoDB) UserAddOnRegistration(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAddOnRegistration", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserAddOnRegistration indicates an expected call of UserAddOnRegistration.
func (mr *MockRepoDBMockRecorder) UserAddOnRegistration(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAddOnRegistration", reflect.TypeOf((*MockRepoDB)(nil).UserAddOnRegistration), user)
}

// UserCheckExistByMail mocks base method.
func (m *MockRepoDB) UserCheckExistByMail(mail string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCheckExistByMail", mail)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCheckExistByMail indicates an expected call of UserCheckExistByMail.
func (mr *MockRepoDBMockRecorder) UserCheckExistByMail(mail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCheckExistByMail", reflect.TypeOf((*MockRepoDB)(nil).UserCheckExistByMail), mail)
}

// UserGetByMail mocks base method.
func (m *MockRepoDB) UserGetByMail(mail string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGetByMail", mail)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGetByMail indicates an expected call of UserGetByMail.
func (mr *MockRepoDBMockRecorder) UserGetByMail(mail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGetByMail", reflect.TypeOf((*MockRepoDB)(nil).UserGetByMail), mail)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCache) Add(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCacheMockRecorder) Add(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCache)(nil).Add), user)
}

// CheckVerifyRegString mocks base method.
func (m *MockCache) CheckVerifyRegString(checkMail string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVerifyRegString", checkMail)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVerifyRegString indicates an expected call of CheckVerifyRegString.
func (mr *MockCacheMockRecorder) CheckVerifyRegString(checkMail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVerifyRegString", reflect.TypeOf((*MockCache)(nil).CheckVerifyRegString), checkMail)
}

// Get mocks base method.
func (m *MockCache) Get(mail string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", mail)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(mail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), mail)
}

// LengthCache mocks base method.
func (m *MockCache) LengthCache() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LengthCache")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LengthCache indicates an expected call of LengthCache.
func (mr *MockCacheMockRecorder) LengthCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LengthCache", reflect.TypeOf((*MockCache)(nil).LengthCache))
}

// LengthUnverifiedUsers mocks base method.
func (m *MockCache) LengthUnverifiedUsers() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LengthUnverifiedUsers")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LengthUnverifiedUsers indicates an expected call of LengthUnverifiedUsers.
func (mr *MockCacheMockRecorder) LengthUnverifiedUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LengthUnverifiedUsers", reflect.TypeOf((*MockCache)(nil).LengthUnverifiedUsers))
}

// Update mocks base method.
func (m *MockCache) Update(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCacheMockRecorder) Update(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCache)(nil).Update), user)
}

// MockStorageS3 is a mock of StorageS3 interface.
type MockStorageS3 struct {
	ctrl     *gomock.Controller
	recorder *MockStorageS3MockRecorder
}

// MockStorageS3MockRecorder is the mock recorder for MockStorageS3.
type MockStorageS3MockRecorder struct {
	mock *MockStorageS3
}

// NewMockStorageS3 creates a new mock instance.
func NewMockStorageS3(ctrl *gomock.Controller) *MockStorageS3 {
	mock := &MockStorageS3{ctrl: ctrl}
	mock.recorder = &MockStorageS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageS3) EXPECT() *MockStorageS3MockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockStorageS3) Download(arg0 domain.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockStorageS3MockRecorder) Download(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockStorageS3)(nil).Download), arg0)
}

// Remove mocks base method.
func (m *MockStorageS3) Remove(arg0 domain.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStorageS3MockRecorder) Remove(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStorageS3)(nil).Remove), arg0)
}

// Upload mocks base method.
func (m *MockStorageS3) Upload(arg0 domain.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockStorageS3MockRecorder) Upload(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockStorageS3)(nil).Upload), arg0)
}
