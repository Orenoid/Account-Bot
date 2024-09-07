// Code generated by MockGen. DO NOT EDIT.
// Source: dal/user/repository.go

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/orenoid/telegram-account-bot/models"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CheckUserExists mocks base method.
func (m *MockRepository) CheckUserExists(userID uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExists indicates an expected call of CheckUserExists.
func (mr *MockRepositoryMockRecorder) CheckUserExists(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockRepository)(nil).CheckUserExists), userID)
}

// CreateToken mocks base method.
func (m *MockRepository) CreateToken(userID uint, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", userID, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockRepositoryMockRecorder) CreateToken(userID, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockRepository)(nil).CreateToken), userID, token)
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser() (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser")
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser))
}

// DisableAllTokens mocks base method.
func (m *MockRepository) DisableAllTokens(userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableAllTokens", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableAllTokens indicates an expected call of DisableAllTokens.
func (mr *MockRepositoryMockRecorder) DisableAllTokens(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAllTokens", reflect.TypeOf((*MockRepository)(nil).DisableAllTokens), userID)
}

// DisableToken mocks base method.
func (m *MockRepository) DisableToken(userID uint, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableToken", userID, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableToken indicates an expected call of DisableToken.
func (mr *MockRepositoryMockRecorder) DisableToken(userID, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableToken", reflect.TypeOf((*MockRepository)(nil).DisableToken), userID, token)
}

// GetUserBalance mocks base method.
func (m *MockRepository) GetUserBalance(userID uint) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalance", userID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalance indicates an expected call of GetUserBalance.
func (mr *MockRepositoryMockRecorder) GetUserBalance(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalance", reflect.TypeOf((*MockRepository)(nil).GetUserBalance), userID)
}

// MustGetToken mocks base method.
func (m *MockRepository) MustGetToken(token string) (*models.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetToken", token)
	ret0, _ := ret[0].(*models.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MustGetToken indicates an expected call of MustGetToken.
func (mr *MockRepositoryMockRecorder) MustGetToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetToken", reflect.TypeOf((*MockRepository)(nil).MustGetToken), token)
}

// SetUserBalance mocks base method.
func (m *MockRepository) SetUserBalance(userID uint, balance float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserBalance", userID, balance)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserBalance indicates an expected call of SetUserBalance.
func (mr *MockRepositoryMockRecorder) SetUserBalance(userID, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserBalance", reflect.TypeOf((*MockRepository)(nil).SetUserBalance), userID, balance)
}
