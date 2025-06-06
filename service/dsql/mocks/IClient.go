// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dsql "github.com/aws/aws-sdk-go-v2/service/dsql"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCluster(ctx context.Context, params *dsql.CreateClusterInput, optFns ...func(*dsql.Options)) (*dsql.CreateClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCluster")
	}

	var r0 *dsql.CreateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.CreateClusterInput, ...func(*dsql.Options)) (*dsql.CreateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.CreateClusterInput, ...func(*dsql.Options)) *dsql.CreateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.CreateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.CreateClusterInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCluster(ctx context.Context, params *dsql.DeleteClusterInput, optFns ...func(*dsql.Options)) (*dsql.DeleteClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCluster")
	}

	var r0 *dsql.DeleteClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.DeleteClusterInput, ...func(*dsql.Options)) (*dsql.DeleteClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.DeleteClusterInput, ...func(*dsql.Options)) *dsql.DeleteClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.DeleteClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.DeleteClusterInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCluster(ctx context.Context, params *dsql.GetClusterInput, optFns ...func(*dsql.Options)) (*dsql.GetClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCluster")
	}

	var r0 *dsql.GetClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.GetClusterInput, ...func(*dsql.Options)) (*dsql.GetClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.GetClusterInput, ...func(*dsql.Options)) *dsql.GetClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.GetClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.GetClusterInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVpcEndpointServiceName provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetVpcEndpointServiceName(ctx context.Context, params *dsql.GetVpcEndpointServiceNameInput, optFns ...func(*dsql.Options)) (*dsql.GetVpcEndpointServiceNameOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetVpcEndpointServiceName")
	}

	var r0 *dsql.GetVpcEndpointServiceNameOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.GetVpcEndpointServiceNameInput, ...func(*dsql.Options)) (*dsql.GetVpcEndpointServiceNameOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.GetVpcEndpointServiceNameInput, ...func(*dsql.Options)) *dsql.GetVpcEndpointServiceNameOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.GetVpcEndpointServiceNameOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.GetVpcEndpointServiceNameInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListClusters(ctx context.Context, params *dsql.ListClustersInput, optFns ...func(*dsql.Options)) (*dsql.ListClustersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListClusters")
	}

	var r0 *dsql.ListClustersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.ListClustersInput, ...func(*dsql.Options)) (*dsql.ListClustersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.ListClustersInput, ...func(*dsql.Options)) *dsql.ListClustersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.ListClustersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.ListClustersInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *dsql.ListTagsForResourceInput, optFns ...func(*dsql.Options)) (*dsql.ListTagsForResourceOutput, error) {
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

	var r0 *dsql.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.ListTagsForResourceInput, ...func(*dsql.Options)) (*dsql.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.ListTagsForResourceInput, ...func(*dsql.Options)) *dsql.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.ListTagsForResourceInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() dsql.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 dsql.Options
	if rf, ok := ret.Get(0).(func() dsql.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dsql.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *dsql.TagResourceInput, optFns ...func(*dsql.Options)) (*dsql.TagResourceOutput, error) {
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

	var r0 *dsql.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.TagResourceInput, ...func(*dsql.Options)) (*dsql.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.TagResourceInput, ...func(*dsql.Options)) *dsql.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.TagResourceInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *dsql.UntagResourceInput, optFns ...func(*dsql.Options)) (*dsql.UntagResourceOutput, error) {
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

	var r0 *dsql.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.UntagResourceInput, ...func(*dsql.Options)) (*dsql.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.UntagResourceInput, ...func(*dsql.Options)) *dsql.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.UntagResourceInput, ...func(*dsql.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCluster(ctx context.Context, params *dsql.UpdateClusterInput, optFns ...func(*dsql.Options)) (*dsql.UpdateClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCluster")
	}

	var r0 *dsql.UpdateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.UpdateClusterInput, ...func(*dsql.Options)) (*dsql.UpdateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dsql.UpdateClusterInput, ...func(*dsql.Options)) *dsql.UpdateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dsql.UpdateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dsql.UpdateClusterInput, ...func(*dsql.Options)) error); ok {
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
