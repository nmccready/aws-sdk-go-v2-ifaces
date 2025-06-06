// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	scheduler "github.com/aws/aws-sdk-go-v2/service/scheduler"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSchedule")
	}

	var r0 *scheduler.CreateScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.CreateScheduleInput, ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.CreateScheduleInput, ...func(*scheduler.Options)) *scheduler.CreateScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.CreateScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.CreateScheduleInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScheduleGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateScheduleGroup(ctx context.Context, params *scheduler.CreateScheduleGroupInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateScheduleGroup")
	}

	var r0 *scheduler.CreateScheduleGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.CreateScheduleGroupInput, ...func(*scheduler.Options)) (*scheduler.CreateScheduleGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.CreateScheduleGroupInput, ...func(*scheduler.Options)) *scheduler.CreateScheduleGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.CreateScheduleGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.CreateScheduleGroupInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSchedule(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSchedule")
	}

	var r0 *scheduler.DeleteScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.DeleteScheduleInput, ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.DeleteScheduleInput, ...func(*scheduler.Options)) *scheduler.DeleteScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.DeleteScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.DeleteScheduleInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScheduleGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteScheduleGroup(ctx context.Context, params *scheduler.DeleteScheduleGroupInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduleGroup")
	}

	var r0 *scheduler.DeleteScheduleGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.DeleteScheduleGroupInput, ...func(*scheduler.Options)) (*scheduler.DeleteScheduleGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.DeleteScheduleGroupInput, ...func(*scheduler.Options)) *scheduler.DeleteScheduleGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.DeleteScheduleGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.DeleteScheduleGroupInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSchedule(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSchedule")
	}

	var r0 *scheduler.GetScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.GetScheduleInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.GetScheduleInput, ...func(*scheduler.Options)) *scheduler.GetScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.GetScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.GetScheduleInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScheduleGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetScheduleGroup(ctx context.Context, params *scheduler.GetScheduleGroupInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetScheduleGroup")
	}

	var r0 *scheduler.GetScheduleGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.GetScheduleGroupInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.GetScheduleGroupInput, ...func(*scheduler.Options)) *scheduler.GetScheduleGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.GetScheduleGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.GetScheduleGroupInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScheduleGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListScheduleGroups(ctx context.Context, params *scheduler.ListScheduleGroupsInput, optFns ...func(*scheduler.Options)) (*scheduler.ListScheduleGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListScheduleGroups")
	}

	var r0 *scheduler.ListScheduleGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListScheduleGroupsInput, ...func(*scheduler.Options)) (*scheduler.ListScheduleGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListScheduleGroupsInput, ...func(*scheduler.Options)) *scheduler.ListScheduleGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.ListScheduleGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.ListScheduleGroupsInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchedules provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSchedules(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSchedules")
	}

	var r0 *scheduler.ListSchedulesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListSchedulesInput, ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListSchedulesInput, ...func(*scheduler.Options)) *scheduler.ListSchedulesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.ListSchedulesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.ListSchedulesInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *scheduler.ListTagsForResourceInput, optFns ...func(*scheduler.Options)) (*scheduler.ListTagsForResourceOutput, error) {
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

	var r0 *scheduler.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListTagsForResourceInput, ...func(*scheduler.Options)) (*scheduler.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.ListTagsForResourceInput, ...func(*scheduler.Options)) *scheduler.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.ListTagsForResourceInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() scheduler.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 scheduler.Options
	if rf, ok := ret.Get(0).(func() scheduler.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(scheduler.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *scheduler.TagResourceInput, optFns ...func(*scheduler.Options)) (*scheduler.TagResourceOutput, error) {
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

	var r0 *scheduler.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.TagResourceInput, ...func(*scheduler.Options)) (*scheduler.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.TagResourceInput, ...func(*scheduler.Options)) *scheduler.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.TagResourceInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *scheduler.UntagResourceInput, optFns ...func(*scheduler.Options)) (*scheduler.UntagResourceOutput, error) {
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

	var r0 *scheduler.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.UntagResourceInput, ...func(*scheduler.Options)) (*scheduler.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.UntagResourceInput, ...func(*scheduler.Options)) *scheduler.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.UntagResourceInput, ...func(*scheduler.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSchedule(ctx context.Context, params *scheduler.UpdateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.UpdateScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSchedule")
	}

	var r0 *scheduler.UpdateScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.UpdateScheduleInput, ...func(*scheduler.Options)) (*scheduler.UpdateScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.UpdateScheduleInput, ...func(*scheduler.Options)) *scheduler.UpdateScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.UpdateScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.UpdateScheduleInput, ...func(*scheduler.Options)) error); ok {
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
