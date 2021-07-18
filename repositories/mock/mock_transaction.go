// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nqtinh/go-gin-project/models"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// CreateUserTransaction mocks base method.
func (m *MockTransactionRepository) CreateUserTransaction(ctx context.Context, req *models.CreateUserTransactionReq) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTransaction", ctx, req)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTransaction indicates an expected call of CreateUserTransaction.
func (mr *MockTransactionRepositoryMockRecorder) CreateUserTransaction(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).CreateUserTransaction), ctx, req)
}

// DeleteTransactionByID mocks base method.
func (m *MockTransactionRepository) DeleteTransactionByID(ctx context.Context, transactionID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransactionByID", ctx, transactionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransactionByID indicates an expected call of DeleteTransactionByID.
func (mr *MockTransactionRepositoryMockRecorder) DeleteTransactionByID(ctx, transactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransactionByID", reflect.TypeOf((*MockTransactionRepository)(nil).DeleteTransactionByID), ctx, transactionID)
}

// DeleteUserAccountTransactions mocks base method.
func (m *MockTransactionRepository) DeleteUserAccountTransactions(ctx context.Context, userID, accountID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserAccountTransactions", ctx, userID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserAccountTransactions indicates an expected call of DeleteUserAccountTransactions.
func (mr *MockTransactionRepositoryMockRecorder) DeleteUserAccountTransactions(ctx, userID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserAccountTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).DeleteUserAccountTransactions), ctx, userID, accountID)
}

// DeleteUserTransactions mocks base method.
func (m *MockTransactionRepository) DeleteUserTransactions(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserTransactions", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserTransactions indicates an expected call of DeleteUserTransactions.
func (mr *MockTransactionRepositoryMockRecorder) DeleteUserTransactions(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).DeleteUserTransactions), ctx, userID)
}

// GetUserTransaction mocks base method.
func (m *MockTransactionRepository) GetUserTransaction(ctx context.Context, transactionID int) (*models.UserTransactionResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTransaction", ctx, transactionID)
	ret0, _ := ret[0].(*models.UserTransactionResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTransaction indicates an expected call of GetUserTransaction.
func (mr *MockTransactionRepositoryMockRecorder) GetUserTransaction(ctx, transactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).GetUserTransaction), ctx, transactionID)
}

// GetUserTransactions mocks base method.
func (m *MockTransactionRepository) GetUserTransactions(ctx context.Context, req *models.GetUserTransactionsReq) ([]models.UserTransactionResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTransactions", ctx, req)
	ret0, _ := ret[0].([]models.UserTransactionResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTransactions indicates an expected call of GetUserTransactions.
func (mr *MockTransactionRepositoryMockRecorder) GetUserTransactions(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).GetUserTransactions), ctx, req)
}

// UpdateUserAccountTransactions mocks base method.
func (m *MockTransactionRepository) UpdateUserAccountTransactions(ctx context.Context, req *models.UpdateUserAccountTransactionsReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAccountTransactions", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAccountTransactions indicates an expected call of UpdateUserAccountTransactions.
func (mr *MockTransactionRepositoryMockRecorder) UpdateUserAccountTransactions(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAccountTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateUserAccountTransactions), ctx, req)
}

// UpdateUserTransactions mocks base method.
func (m *MockTransactionRepository) UpdateUserTransactions(ctx context.Context, req *models.UpdateUserTransactionsReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserTransactions", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserTransactions indicates an expected call of UpdateUserTransactions.
func (mr *MockTransactionRepositoryMockRecorder) UpdateUserTransactions(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateUserTransactions), ctx, req)
}
