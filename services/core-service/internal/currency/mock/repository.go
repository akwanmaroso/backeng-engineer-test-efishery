// Code generated by MockGen. DO NOT EDIT.
// Source: internal/currency/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	gomock "github.com/golang/mock/gomock"
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

// GetCurrencyUSDToIDR mocks base method.
func (m *MockRepository) GetCurrencyUSDToIDR(arg0 context.Context) (models.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencyUSDToIDR", arg0)
	ret0, _ := ret[0].(models.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencyUSDToIDR indicates an expected call of GetCurrencyUSDToIDR.
func (mr *MockRepositoryMockRecorder) GetCurrencyUSDToIDR(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencyUSDToIDR", reflect.TypeOf((*MockRepository)(nil).GetCurrencyUSDToIDR), arg0)
}
