// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplaceentitlementservice "github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetEntitlements provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetEntitlements(ctx context.Context, params *marketplaceentitlementservice.GetEntitlementsInput, optFns ...func(*marketplaceentitlementservice.Options)) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEntitlements")
	}

	var r0 *marketplaceentitlementservice.GetEntitlementsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceentitlementservice.GetEntitlementsInput, ...func(*marketplaceentitlementservice.Options)) (*marketplaceentitlementservice.GetEntitlementsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceentitlementservice.GetEntitlementsInput, ...func(*marketplaceentitlementservice.Options)) *marketplaceentitlementservice.GetEntitlementsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplaceentitlementservice.GetEntitlementsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplaceentitlementservice.GetEntitlementsInput, ...func(*marketplaceentitlementservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() marketplaceentitlementservice.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplaceentitlementservice.Options
	if rf, ok := ret.Get(0).(func() marketplaceentitlementservice.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplaceentitlementservice.Options)
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
