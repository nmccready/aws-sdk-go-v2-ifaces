// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplacedeployment "github.com/aws/aws-sdk-go-v2/service/marketplacedeployment"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *marketplacedeployment.ListTagsForResourceInput, optFns ...func(*marketplacedeployment.Options)) (*marketplacedeployment.ListTagsForResourceOutput, error) {
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

	var r0 *marketplacedeployment.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.ListTagsForResourceInput, ...func(*marketplacedeployment.Options)) (*marketplacedeployment.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.ListTagsForResourceInput, ...func(*marketplacedeployment.Options)) *marketplacedeployment.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacedeployment.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacedeployment.ListTagsForResourceInput, ...func(*marketplacedeployment.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() marketplacedeployment.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplacedeployment.Options
	if rf, ok := ret.Get(0).(func() marketplacedeployment.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplacedeployment.Options)
	}

	return r0
}

// PutDeploymentParameter provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutDeploymentParameter(ctx context.Context, params *marketplacedeployment.PutDeploymentParameterInput, optFns ...func(*marketplacedeployment.Options)) (*marketplacedeployment.PutDeploymentParameterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutDeploymentParameter")
	}

	var r0 *marketplacedeployment.PutDeploymentParameterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.PutDeploymentParameterInput, ...func(*marketplacedeployment.Options)) (*marketplacedeployment.PutDeploymentParameterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.PutDeploymentParameterInput, ...func(*marketplacedeployment.Options)) *marketplacedeployment.PutDeploymentParameterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacedeployment.PutDeploymentParameterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacedeployment.PutDeploymentParameterInput, ...func(*marketplacedeployment.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *marketplacedeployment.TagResourceInput, optFns ...func(*marketplacedeployment.Options)) (*marketplacedeployment.TagResourceOutput, error) {
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

	var r0 *marketplacedeployment.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.TagResourceInput, ...func(*marketplacedeployment.Options)) (*marketplacedeployment.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.TagResourceInput, ...func(*marketplacedeployment.Options)) *marketplacedeployment.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacedeployment.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacedeployment.TagResourceInput, ...func(*marketplacedeployment.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *marketplacedeployment.UntagResourceInput, optFns ...func(*marketplacedeployment.Options)) (*marketplacedeployment.UntagResourceOutput, error) {
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

	var r0 *marketplacedeployment.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.UntagResourceInput, ...func(*marketplacedeployment.Options)) (*marketplacedeployment.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacedeployment.UntagResourceInput, ...func(*marketplacedeployment.Options)) *marketplacedeployment.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacedeployment.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacedeployment.UntagResourceInput, ...func(*marketplacedeployment.Options)) error); ok {
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
