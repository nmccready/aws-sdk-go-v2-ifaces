// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	iotfleethub "github.com/aws/aws-sdk-go-v2/service/iotfleethub"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateApplication(ctx context.Context, params *iotfleethub.CreateApplicationInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.CreateApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateApplication")
	}

	var r0 *iotfleethub.CreateApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.CreateApplicationInput, ...func(*iotfleethub.Options)) (*iotfleethub.CreateApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.CreateApplicationInput, ...func(*iotfleethub.Options)) *iotfleethub.CreateApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.CreateApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.CreateApplicationInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteApplication(ctx context.Context, params *iotfleethub.DeleteApplicationInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.DeleteApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteApplication")
	}

	var r0 *iotfleethub.DeleteApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.DeleteApplicationInput, ...func(*iotfleethub.Options)) (*iotfleethub.DeleteApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.DeleteApplicationInput, ...func(*iotfleethub.Options)) *iotfleethub.DeleteApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.DeleteApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.DeleteApplicationInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeApplication(ctx context.Context, params *iotfleethub.DescribeApplicationInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.DescribeApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeApplication")
	}

	var r0 *iotfleethub.DescribeApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.DescribeApplicationInput, ...func(*iotfleethub.Options)) (*iotfleethub.DescribeApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.DescribeApplicationInput, ...func(*iotfleethub.Options)) *iotfleethub.DescribeApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.DescribeApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.DescribeApplicationInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplications provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApplications(ctx context.Context, params *iotfleethub.ListApplicationsInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.ListApplicationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListApplications")
	}

	var r0 *iotfleethub.ListApplicationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.ListApplicationsInput, ...func(*iotfleethub.Options)) (*iotfleethub.ListApplicationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.ListApplicationsInput, ...func(*iotfleethub.Options)) *iotfleethub.ListApplicationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.ListApplicationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.ListApplicationsInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *iotfleethub.ListTagsForResourceInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.ListTagsForResourceOutput, error) {
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

	var r0 *iotfleethub.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.ListTagsForResourceInput, ...func(*iotfleethub.Options)) (*iotfleethub.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.ListTagsForResourceInput, ...func(*iotfleethub.Options)) *iotfleethub.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.ListTagsForResourceInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() iotfleethub.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 iotfleethub.Options
	if rf, ok := ret.Get(0).(func() iotfleethub.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(iotfleethub.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *iotfleethub.TagResourceInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.TagResourceOutput, error) {
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

	var r0 *iotfleethub.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.TagResourceInput, ...func(*iotfleethub.Options)) (*iotfleethub.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.TagResourceInput, ...func(*iotfleethub.Options)) *iotfleethub.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.TagResourceInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *iotfleethub.UntagResourceInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.UntagResourceOutput, error) {
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

	var r0 *iotfleethub.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.UntagResourceInput, ...func(*iotfleethub.Options)) (*iotfleethub.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.UntagResourceInput, ...func(*iotfleethub.Options)) *iotfleethub.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.UntagResourceInput, ...func(*iotfleethub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateApplication(ctx context.Context, params *iotfleethub.UpdateApplicationInput, optFns ...func(*iotfleethub.Options)) (*iotfleethub.UpdateApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateApplication")
	}

	var r0 *iotfleethub.UpdateApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.UpdateApplicationInput, ...func(*iotfleethub.Options)) (*iotfleethub.UpdateApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotfleethub.UpdateApplicationInput, ...func(*iotfleethub.Options)) *iotfleethub.UpdateApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotfleethub.UpdateApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotfleethub.UpdateApplicationInput, ...func(*iotfleethub.Options)) error); ok {
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
