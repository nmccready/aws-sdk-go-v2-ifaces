// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	braket "github.com/aws/aws-sdk-go-v2/service/braket"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelJob(ctx context.Context, params *braket.CancelJobInput, optFns ...func(*braket.Options)) (*braket.CancelJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelJob")
	}

	var r0 *braket.CancelJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CancelJobInput, ...func(*braket.Options)) (*braket.CancelJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CancelJobInput, ...func(*braket.Options)) *braket.CancelJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.CancelJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.CancelJobInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelQuantumTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelQuantumTask(ctx context.Context, params *braket.CancelQuantumTaskInput, optFns ...func(*braket.Options)) (*braket.CancelQuantumTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelQuantumTask")
	}

	var r0 *braket.CancelQuantumTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CancelQuantumTaskInput, ...func(*braket.Options)) (*braket.CancelQuantumTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CancelQuantumTaskInput, ...func(*braket.Options)) *braket.CancelQuantumTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.CancelQuantumTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.CancelQuantumTaskInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateJob(ctx context.Context, params *braket.CreateJobInput, optFns ...func(*braket.Options)) (*braket.CreateJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 *braket.CreateJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CreateJobInput, ...func(*braket.Options)) (*braket.CreateJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CreateJobInput, ...func(*braket.Options)) *braket.CreateJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.CreateJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.CreateJobInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQuantumTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateQuantumTask(ctx context.Context, params *braket.CreateQuantumTaskInput, optFns ...func(*braket.Options)) (*braket.CreateQuantumTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateQuantumTask")
	}

	var r0 *braket.CreateQuantumTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CreateQuantumTaskInput, ...func(*braket.Options)) (*braket.CreateQuantumTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.CreateQuantumTaskInput, ...func(*braket.Options)) *braket.CreateQuantumTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.CreateQuantumTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.CreateQuantumTaskInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDevice(ctx context.Context, params *braket.GetDeviceInput, optFns ...func(*braket.Options)) (*braket.GetDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDevice")
	}

	var r0 *braket.GetDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetDeviceInput, ...func(*braket.Options)) (*braket.GetDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetDeviceInput, ...func(*braket.Options)) *braket.GetDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.GetDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.GetDeviceInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetJob(ctx context.Context, params *braket.GetJobInput, optFns ...func(*braket.Options)) (*braket.GetJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetJob")
	}

	var r0 *braket.GetJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetJobInput, ...func(*braket.Options)) (*braket.GetJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetJobInput, ...func(*braket.Options)) *braket.GetJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.GetJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.GetJobInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuantumTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetQuantumTask(ctx context.Context, params *braket.GetQuantumTaskInput, optFns ...func(*braket.Options)) (*braket.GetQuantumTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetQuantumTask")
	}

	var r0 *braket.GetQuantumTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetQuantumTaskInput, ...func(*braket.Options)) (*braket.GetQuantumTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.GetQuantumTaskInput, ...func(*braket.Options)) *braket.GetQuantumTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.GetQuantumTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.GetQuantumTaskInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *braket.ListTagsForResourceInput, optFns ...func(*braket.Options)) (*braket.ListTagsForResourceOutput, error) {
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

	var r0 *braket.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.ListTagsForResourceInput, ...func(*braket.Options)) (*braket.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.ListTagsForResourceInput, ...func(*braket.Options)) *braket.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.ListTagsForResourceInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() braket.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 braket.Options
	if rf, ok := ret.Get(0).(func() braket.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(braket.Options)
	}

	return r0
}

// SearchDevices provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SearchDevices(ctx context.Context, params *braket.SearchDevicesInput, optFns ...func(*braket.Options)) (*braket.SearchDevicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SearchDevices")
	}

	var r0 *braket.SearchDevicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchDevicesInput, ...func(*braket.Options)) (*braket.SearchDevicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchDevicesInput, ...func(*braket.Options)) *braket.SearchDevicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.SearchDevicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.SearchDevicesInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SearchJobs(ctx context.Context, params *braket.SearchJobsInput, optFns ...func(*braket.Options)) (*braket.SearchJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SearchJobs")
	}

	var r0 *braket.SearchJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchJobsInput, ...func(*braket.Options)) (*braket.SearchJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchJobsInput, ...func(*braket.Options)) *braket.SearchJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.SearchJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.SearchJobsInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchQuantumTasks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SearchQuantumTasks(ctx context.Context, params *braket.SearchQuantumTasksInput, optFns ...func(*braket.Options)) (*braket.SearchQuantumTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SearchQuantumTasks")
	}

	var r0 *braket.SearchQuantumTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchQuantumTasksInput, ...func(*braket.Options)) (*braket.SearchQuantumTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.SearchQuantumTasksInput, ...func(*braket.Options)) *braket.SearchQuantumTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.SearchQuantumTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.SearchQuantumTasksInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *braket.TagResourceInput, optFns ...func(*braket.Options)) (*braket.TagResourceOutput, error) {
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

	var r0 *braket.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.TagResourceInput, ...func(*braket.Options)) (*braket.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.TagResourceInput, ...func(*braket.Options)) *braket.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.TagResourceInput, ...func(*braket.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *braket.UntagResourceInput, optFns ...func(*braket.Options)) (*braket.UntagResourceOutput, error) {
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

	var r0 *braket.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *braket.UntagResourceInput, ...func(*braket.Options)) (*braket.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *braket.UntagResourceInput, ...func(*braket.Options)) *braket.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*braket.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *braket.UntagResourceInput, ...func(*braket.Options)) error); ok {
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
