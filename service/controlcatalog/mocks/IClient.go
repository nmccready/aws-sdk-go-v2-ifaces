// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	controlcatalog "github.com/aws/aws-sdk-go-v2/service/controlcatalog"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// ListCommonControls provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCommonControls(ctx context.Context, params *controlcatalog.ListCommonControlsInput, optFns ...func(*controlcatalog.Options)) (*controlcatalog.ListCommonControlsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCommonControls")
	}

	var r0 *controlcatalog.ListCommonControlsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListCommonControlsInput, ...func(*controlcatalog.Options)) (*controlcatalog.ListCommonControlsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListCommonControlsInput, ...func(*controlcatalog.Options)) *controlcatalog.ListCommonControlsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controlcatalog.ListCommonControlsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *controlcatalog.ListCommonControlsInput, ...func(*controlcatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDomains provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDomains(ctx context.Context, params *controlcatalog.ListDomainsInput, optFns ...func(*controlcatalog.Options)) (*controlcatalog.ListDomainsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDomains")
	}

	var r0 *controlcatalog.ListDomainsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListDomainsInput, ...func(*controlcatalog.Options)) (*controlcatalog.ListDomainsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListDomainsInput, ...func(*controlcatalog.Options)) *controlcatalog.ListDomainsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controlcatalog.ListDomainsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *controlcatalog.ListDomainsInput, ...func(*controlcatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectives provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListObjectives(ctx context.Context, params *controlcatalog.ListObjectivesInput, optFns ...func(*controlcatalog.Options)) (*controlcatalog.ListObjectivesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListObjectives")
	}

	var r0 *controlcatalog.ListObjectivesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListObjectivesInput, ...func(*controlcatalog.Options)) (*controlcatalog.ListObjectivesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *controlcatalog.ListObjectivesInput, ...func(*controlcatalog.Options)) *controlcatalog.ListObjectivesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controlcatalog.ListObjectivesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *controlcatalog.ListObjectivesInput, ...func(*controlcatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() controlcatalog.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 controlcatalog.Options
	if rf, ok := ret.Get(0).(func() controlcatalog.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(controlcatalog.Options)
	}

	return r0
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