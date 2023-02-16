// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	"github.com/kyma-project/telemetry-manager/apis/telemetry/v1alpha1"
	"github.com/stretchr/testify/mock"
)

// InputValidator is an autogenerated mock type for the InputValidator type
type InputValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: logPipelineInput
func (_m *InputValidator) Validate(logPipelineInput *v1alpha1.Input) error {
	ret := _m.Called(logPipelineInput)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Input) error); ok {
		r0 = rf(logPipelineInput)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewInputValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewInputValidator creates a new instance of InputValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInputValidator(t mockConstructorTestingTNewInputValidator) *InputValidator {
	mock := &InputValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
