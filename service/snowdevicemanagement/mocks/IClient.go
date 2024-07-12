// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	snowdevicemanagement "github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelTask(ctx context.Context, params *snowdevicemanagement.CancelTaskInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.CancelTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelTask")
	}

	var r0 *snowdevicemanagement.CancelTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.CancelTaskInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.CancelTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.CancelTaskInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.CancelTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.CancelTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.CancelTaskInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTask(ctx context.Context, params *snowdevicemanagement.CreateTaskInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.CreateTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 *snowdevicemanagement.CreateTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.CreateTaskInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.CreateTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.CreateTaskInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.CreateTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.CreateTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.CreateTaskInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDevice(ctx context.Context, params *snowdevicemanagement.DescribeDeviceInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDevice")
	}

	var r0 *snowdevicemanagement.DescribeDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeDeviceInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeDeviceInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.DescribeDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.DescribeDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.DescribeDeviceInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDeviceEc2Instances provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDeviceEc2Instances(ctx context.Context, params *snowdevicemanagement.DescribeDeviceEc2InstancesInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDeviceEc2Instances")
	}

	var r0 *snowdevicemanagement.DescribeDeviceEc2InstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeDeviceEc2InstancesInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeDeviceEc2InstancesInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.DescribeDeviceEc2InstancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.DescribeDeviceEc2InstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.DescribeDeviceEc2InstancesInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeExecution(ctx context.Context, params *snowdevicemanagement.DescribeExecutionInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeExecution")
	}

	var r0 *snowdevicemanagement.DescribeExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeExecutionInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeExecutionInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.DescribeExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.DescribeExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.DescribeExecutionInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeTask(ctx context.Context, params *snowdevicemanagement.DescribeTaskInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeTask")
	}

	var r0 *snowdevicemanagement.DescribeTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeTaskInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.DescribeTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.DescribeTaskInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.DescribeTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.DescribeTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.DescribeTaskInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeviceResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDeviceResources(ctx context.Context, params *snowdevicemanagement.ListDeviceResourcesInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListDeviceResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDeviceResources")
	}

	var r0 *snowdevicemanagement.ListDeviceResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListDeviceResourcesInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListDeviceResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListDeviceResourcesInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.ListDeviceResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.ListDeviceResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.ListDeviceResourcesInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDevices provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDevices(ctx context.Context, params *snowdevicemanagement.ListDevicesInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListDevicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 *snowdevicemanagement.ListDevicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListDevicesInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListDevicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListDevicesInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.ListDevicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.ListDevicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.ListDevicesInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExecutions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExecutions(ctx context.Context, params *snowdevicemanagement.ListExecutionsInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListExecutionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExecutions")
	}

	var r0 *snowdevicemanagement.ListExecutionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListExecutionsInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListExecutionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListExecutionsInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.ListExecutionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.ListExecutionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.ListExecutionsInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *snowdevicemanagement.ListTagsForResourceInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListTagsForResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTagsForResource")
	}

	var r0 *snowdevicemanagement.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListTagsForResourceInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListTagsForResourceInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.ListTagsForResourceInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTasks(ctx context.Context, params *snowdevicemanagement.ListTasksInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTasks")
	}

	var r0 *snowdevicemanagement.ListTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListTasksInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.ListTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.ListTasksInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.ListTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.ListTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.ListTasksInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() snowdevicemanagement.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 snowdevicemanagement.Options
	if rf, ok := ret.Get(0).(func() snowdevicemanagement.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(snowdevicemanagement.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *snowdevicemanagement.TagResourceInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.TagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TagResource")
	}

	var r0 *snowdevicemanagement.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.TagResourceInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.TagResourceInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.TagResourceInput, ...func(*snowdevicemanagement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *snowdevicemanagement.UntagResourceInput, optFns ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.UntagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UntagResource")
	}

	var r0 *snowdevicemanagement.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.UntagResourceInput, ...func(*snowdevicemanagement.Options)) (*snowdevicemanagement.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowdevicemanagement.UntagResourceInput, ...func(*snowdevicemanagement.Options)) *snowdevicemanagement.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowdevicemanagement.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowdevicemanagement.UntagResourceInput, ...func(*snowdevicemanagement.Options)) error); ok {
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
