// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplacereporting "github.com/aws/aws-sdk-go-v2/service/marketplacereporting"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetBuyerDashboard provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetBuyerDashboard(ctx context.Context, params *marketplacereporting.GetBuyerDashboardInput, optFns ...func(*marketplacereporting.Options)) (*marketplacereporting.GetBuyerDashboardOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBuyerDashboard")
	}

	var r0 *marketplacereporting.GetBuyerDashboardOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacereporting.GetBuyerDashboardInput, ...func(*marketplacereporting.Options)) (*marketplacereporting.GetBuyerDashboardOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacereporting.GetBuyerDashboardInput, ...func(*marketplacereporting.Options)) *marketplacereporting.GetBuyerDashboardOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacereporting.GetBuyerDashboardOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacereporting.GetBuyerDashboardInput, ...func(*marketplacereporting.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() marketplacereporting.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplacereporting.Options
	if rf, ok := ret.Get(0).(func() marketplacereporting.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplacereporting.Options)
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