// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	observabilityadmin "github.com/aws/aws-sdk-go-v2/service/observabilityadmin"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetTelemetryEvaluationStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTelemetryEvaluationStatus(ctx context.Context, params *observabilityadmin.GetTelemetryEvaluationStatusInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.GetTelemetryEvaluationStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTelemetryEvaluationStatus")
	}

	var r0 *observabilityadmin.GetTelemetryEvaluationStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.GetTelemetryEvaluationStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusInput, ...func(*observabilityadmin.Options)) *observabilityadmin.GetTelemetryEvaluationStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.GetTelemetryEvaluationStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTelemetryEvaluationStatusForOrganization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTelemetryEvaluationStatusForOrganization(ctx context.Context, params *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.GetTelemetryEvaluationStatusForOrganizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTelemetryEvaluationStatusForOrganization")
	}

	var r0 *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.GetTelemetryEvaluationStatusForOrganizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.GetTelemetryEvaluationStatusForOrganizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.GetTelemetryEvaluationStatusForOrganizationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResourceTelemetry provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListResourceTelemetry(ctx context.Context, params *observabilityadmin.ListResourceTelemetryInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.ListResourceTelemetryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListResourceTelemetry")
	}

	var r0 *observabilityadmin.ListResourceTelemetryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.ListResourceTelemetryInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.ListResourceTelemetryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.ListResourceTelemetryInput, ...func(*observabilityadmin.Options)) *observabilityadmin.ListResourceTelemetryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.ListResourceTelemetryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.ListResourceTelemetryInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResourceTelemetryForOrganization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListResourceTelemetryForOrganization(ctx context.Context, params *observabilityadmin.ListResourceTelemetryForOrganizationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.ListResourceTelemetryForOrganizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListResourceTelemetryForOrganization")
	}

	var r0 *observabilityadmin.ListResourceTelemetryForOrganizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.ListResourceTelemetryForOrganizationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.ListResourceTelemetryForOrganizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.ListResourceTelemetryForOrganizationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.ListResourceTelemetryForOrganizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.ListResourceTelemetryForOrganizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.ListResourceTelemetryForOrganizationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() observabilityadmin.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 observabilityadmin.Options
	if rf, ok := ret.Get(0).(func() observabilityadmin.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(observabilityadmin.Options)
	}

	return r0
}

// StartTelemetryEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartTelemetryEvaluation(ctx context.Context, params *observabilityadmin.StartTelemetryEvaluationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.StartTelemetryEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartTelemetryEvaluation")
	}

	var r0 *observabilityadmin.StartTelemetryEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.StartTelemetryEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.StartTelemetryEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.StartTelemetryEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartTelemetryEvaluationForOrganization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartTelemetryEvaluationForOrganization(ctx context.Context, params *observabilityadmin.StartTelemetryEvaluationForOrganizationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.StartTelemetryEvaluationForOrganizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartTelemetryEvaluationForOrganization")
	}

	var r0 *observabilityadmin.StartTelemetryEvaluationForOrganizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.StartTelemetryEvaluationForOrganizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.StartTelemetryEvaluationForOrganizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.StartTelemetryEvaluationForOrganizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.StartTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopTelemetryEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopTelemetryEvaluation(ctx context.Context, params *observabilityadmin.StopTelemetryEvaluationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.StopTelemetryEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopTelemetryEvaluation")
	}

	var r0 *observabilityadmin.StopTelemetryEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.StopTelemetryEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.StopTelemetryEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.StopTelemetryEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopTelemetryEvaluationForOrganization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopTelemetryEvaluationForOrganization(ctx context.Context, params *observabilityadmin.StopTelemetryEvaluationForOrganizationInput, optFns ...func(*observabilityadmin.Options)) (*observabilityadmin.StopTelemetryEvaluationForOrganizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopTelemetryEvaluationForOrganization")
	}

	var r0 *observabilityadmin.StopTelemetryEvaluationForOrganizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) (*observabilityadmin.StopTelemetryEvaluationForOrganizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) *observabilityadmin.StopTelemetryEvaluationForOrganizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observabilityadmin.StopTelemetryEvaluationForOrganizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observabilityadmin.StopTelemetryEvaluationForOrganizationInput, ...func(*observabilityadmin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}