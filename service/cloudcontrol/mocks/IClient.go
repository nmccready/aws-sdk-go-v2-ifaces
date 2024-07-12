// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	cloudcontrol "github.com/aws/aws-sdk-go-v2/service/cloudcontrol"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelResourceRequest provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelResourceRequest(ctx context.Context, params *cloudcontrol.CancelResourceRequestInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.CancelResourceRequestOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelResourceRequest")
	}

	var r0 *cloudcontrol.CancelResourceRequestOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.CancelResourceRequestInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.CancelResourceRequestOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.CancelResourceRequestInput, ...func(*cloudcontrol.Options)) *cloudcontrol.CancelResourceRequestOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.CancelResourceRequestOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.CancelResourceRequestInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateResource(ctx context.Context, params *cloudcontrol.CreateResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.CreateResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateResource")
	}

	var r0 *cloudcontrol.CreateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.CreateResourceInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.CreateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.CreateResourceInput, ...func(*cloudcontrol.Options)) *cloudcontrol.CreateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.CreateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.CreateResourceInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteResource(ctx context.Context, params *cloudcontrol.DeleteResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.DeleteResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResource")
	}

	var r0 *cloudcontrol.DeleteResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.DeleteResourceInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.DeleteResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.DeleteResourceInput, ...func(*cloudcontrol.Options)) *cloudcontrol.DeleteResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.DeleteResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.DeleteResourceInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResource(ctx context.Context, params *cloudcontrol.GetResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResource")
	}

	var r0 *cloudcontrol.GetResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.GetResourceInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.GetResourceInput, ...func(*cloudcontrol.Options)) *cloudcontrol.GetResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.GetResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.GetResourceInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceRequestStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourceRequestStatus(ctx context.Context, params *cloudcontrol.GetResourceRequestStatusInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceRequestStatus")
	}

	var r0 *cloudcontrol.GetResourceRequestStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) *cloudcontrol.GetResourceRequestStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.GetResourceRequestStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResourceRequests provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListResourceRequests(ctx context.Context, params *cloudcontrol.ListResourceRequestsInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.ListResourceRequestsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListResourceRequests")
	}

	var r0 *cloudcontrol.ListResourceRequestsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.ListResourceRequestsInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.ListResourceRequestsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.ListResourceRequestsInput, ...func(*cloudcontrol.Options)) *cloudcontrol.ListResourceRequestsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.ListResourceRequestsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.ListResourceRequestsInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListResources(ctx context.Context, params *cloudcontrol.ListResourcesInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.ListResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListResources")
	}

	var r0 *cloudcontrol.ListResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.ListResourcesInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.ListResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.ListResourcesInput, ...func(*cloudcontrol.Options)) *cloudcontrol.ListResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.ListResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.ListResourcesInput, ...func(*cloudcontrol.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() cloudcontrol.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 cloudcontrol.Options
	if rf, ok := ret.Get(0).(func() cloudcontrol.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cloudcontrol.Options)
	}

	return r0
}

// UpdateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateResource(ctx context.Context, params *cloudcontrol.UpdateResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.UpdateResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateResource")
	}

	var r0 *cloudcontrol.UpdateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.UpdateResourceInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.UpdateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudcontrol.UpdateResourceInput, ...func(*cloudcontrol.Options)) *cloudcontrol.UpdateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudcontrol.UpdateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudcontrol.UpdateResourceInput, ...func(*cloudcontrol.Options)) error); ok {
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
