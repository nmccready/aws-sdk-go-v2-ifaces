// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	workmailmessageflow "github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetRawMessageContent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRawMessageContent(ctx context.Context, params *workmailmessageflow.GetRawMessageContentInput, optFns ...func(*workmailmessageflow.Options)) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRawMessageContent")
	}

	var r0 *workmailmessageflow.GetRawMessageContentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workmailmessageflow.GetRawMessageContentInput, ...func(*workmailmessageflow.Options)) (*workmailmessageflow.GetRawMessageContentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workmailmessageflow.GetRawMessageContentInput, ...func(*workmailmessageflow.Options)) *workmailmessageflow.GetRawMessageContentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workmailmessageflow.GetRawMessageContentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workmailmessageflow.GetRawMessageContentInput, ...func(*workmailmessageflow.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() workmailmessageflow.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 workmailmessageflow.Options
	if rf, ok := ret.Get(0).(func() workmailmessageflow.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(workmailmessageflow.Options)
	}

	return r0
}

// PutRawMessageContent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutRawMessageContent(ctx context.Context, params *workmailmessageflow.PutRawMessageContentInput, optFns ...func(*workmailmessageflow.Options)) (*workmailmessageflow.PutRawMessageContentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRawMessageContent")
	}

	var r0 *workmailmessageflow.PutRawMessageContentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workmailmessageflow.PutRawMessageContentInput, ...func(*workmailmessageflow.Options)) (*workmailmessageflow.PutRawMessageContentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workmailmessageflow.PutRawMessageContentInput, ...func(*workmailmessageflow.Options)) *workmailmessageflow.PutRawMessageContentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workmailmessageflow.PutRawMessageContentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workmailmessageflow.PutRawMessageContentInput, ...func(*workmailmessageflow.Options)) error); ok {
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