// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	iotjobsdataplane "github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DescribeJobExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeJobExecution(ctx context.Context, params *iotjobsdataplane.DescribeJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeJobExecution")
	}

	var r0 *iotjobsdataplane.DescribeJobExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.DescribeJobExecutionInput, ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.DescribeJobExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.DescribeJobExecutionInput, ...func(*iotjobsdataplane.Options)) *iotjobsdataplane.DescribeJobExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotjobsdataplane.DescribeJobExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotjobsdataplane.DescribeJobExecutionInput, ...func(*iotjobsdataplane.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingJobExecutions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPendingJobExecutions(ctx context.Context, params *iotjobsdataplane.GetPendingJobExecutionsInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingJobExecutions")
	}

	var r0 *iotjobsdataplane.GetPendingJobExecutionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.GetPendingJobExecutionsInput, ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.GetPendingJobExecutionsInput, ...func(*iotjobsdataplane.Options)) *iotjobsdataplane.GetPendingJobExecutionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotjobsdataplane.GetPendingJobExecutionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotjobsdataplane.GetPendingJobExecutionsInput, ...func(*iotjobsdataplane.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() iotjobsdataplane.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 iotjobsdataplane.Options
	if rf, ok := ret.Get(0).(func() iotjobsdataplane.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(iotjobsdataplane.Options)
	}

	return r0
}

// StartCommandExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartCommandExecution(ctx context.Context, params *iotjobsdataplane.StartCommandExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.StartCommandExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartCommandExecution")
	}

	var r0 *iotjobsdataplane.StartCommandExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.StartCommandExecutionInput, ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.StartCommandExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.StartCommandExecutionInput, ...func(*iotjobsdataplane.Options)) *iotjobsdataplane.StartCommandExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotjobsdataplane.StartCommandExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotjobsdataplane.StartCommandExecutionInput, ...func(*iotjobsdataplane.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartNextPendingJobExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartNextPendingJobExecution(ctx context.Context, params *iotjobsdataplane.StartNextPendingJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartNextPendingJobExecution")
	}

	var r0 *iotjobsdataplane.StartNextPendingJobExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.StartNextPendingJobExecutionInput, ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.StartNextPendingJobExecutionInput, ...func(*iotjobsdataplane.Options)) *iotjobsdataplane.StartNextPendingJobExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotjobsdataplane.StartNextPendingJobExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotjobsdataplane.StartNextPendingJobExecutionInput, ...func(*iotjobsdataplane.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateJobExecution(ctx context.Context, params *iotjobsdataplane.UpdateJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJobExecution")
	}

	var r0 *iotjobsdataplane.UpdateJobExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.UpdateJobExecutionInput, ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.UpdateJobExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotjobsdataplane.UpdateJobExecutionInput, ...func(*iotjobsdataplane.Options)) *iotjobsdataplane.UpdateJobExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotjobsdataplane.UpdateJobExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotjobsdataplane.UpdateJobExecutionInput, ...func(*iotjobsdataplane.Options)) error); ok {
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
