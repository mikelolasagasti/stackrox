// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/pkg/license/validator (interfaces: Validator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	license "github.com/stackrox/rox/generated/shared/license"
	validator "github.com/stackrox/rox/pkg/license/validator"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// RegisterSigningKey mocks base method
func (m *MockValidator) RegisterSigningKey(arg0 string, arg1 []byte, arg2 validator.SigningKeyRestrictions) error {
	ret := m.ctrl.Call(m, "RegisterSigningKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterSigningKey indicates an expected call of RegisterSigningKey
func (mr *MockValidatorMockRecorder) RegisterSigningKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSigningKey", reflect.TypeOf((*MockValidator)(nil).RegisterSigningKey), arg0, arg1, arg2)
}

// ValidateLicenseKey mocks base method
func (m *MockValidator) ValidateLicenseKey(arg0 string) (*license.License, error) {
	ret := m.ctrl.Call(m, "ValidateLicenseKey", arg0)
	ret0, _ := ret[0].(*license.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateLicenseKey indicates an expected call of ValidateLicenseKey
func (mr *MockValidatorMockRecorder) ValidateLicenseKey(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateLicenseKey", reflect.TypeOf((*MockValidator)(nil).ValidateLicenseKey), arg0)
}
