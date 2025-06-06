// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplacecatalog "github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchDescribeEntities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchDescribeEntities(ctx context.Context, params *marketplacecatalog.BatchDescribeEntitiesInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.BatchDescribeEntitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDescribeEntities")
	}

	var r0 *marketplacecatalog.BatchDescribeEntitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.BatchDescribeEntitiesInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.BatchDescribeEntitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.BatchDescribeEntitiesInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.BatchDescribeEntitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.BatchDescribeEntitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.BatchDescribeEntitiesInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelChangeSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelChangeSet(ctx context.Context, params *marketplacecatalog.CancelChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.CancelChangeSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelChangeSet")
	}

	var r0 *marketplacecatalog.CancelChangeSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.CancelChangeSetInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.CancelChangeSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.CancelChangeSetInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.CancelChangeSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.CancelChangeSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.CancelChangeSetInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteResourcePolicy(ctx context.Context, params *marketplacecatalog.DeleteResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DeleteResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResourcePolicy")
	}

	var r0 *marketplacecatalog.DeleteResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DeleteResourcePolicyInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DeleteResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DeleteResourcePolicyInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.DeleteResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.DeleteResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.DeleteResourcePolicyInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeChangeSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeChangeSet(ctx context.Context, params *marketplacecatalog.DescribeChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeChangeSet")
	}

	var r0 *marketplacecatalog.DescribeChangeSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DescribeChangeSetInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeChangeSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DescribeChangeSetInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.DescribeChangeSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.DescribeChangeSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.DescribeChangeSetInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeEntity provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeEntity(ctx context.Context, params *marketplacecatalog.DescribeEntityInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeEntityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeEntity")
	}

	var r0 *marketplacecatalog.DescribeEntityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DescribeEntityInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeEntityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.DescribeEntityInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.DescribeEntityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.DescribeEntityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.DescribeEntityInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourcePolicy(ctx context.Context, params *marketplacecatalog.GetResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.GetResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcePolicy")
	}

	var r0 *marketplacecatalog.GetResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.GetResourcePolicyInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.GetResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.GetResourcePolicyInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.GetResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.GetResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.GetResourcePolicyInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListChangeSets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListChangeSets(ctx context.Context, params *marketplacecatalog.ListChangeSetsInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListChangeSetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListChangeSets")
	}

	var r0 *marketplacecatalog.ListChangeSetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListChangeSetsInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListChangeSetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListChangeSetsInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.ListChangeSetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.ListChangeSetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.ListChangeSetsInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEntities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListEntities(ctx context.Context, params *marketplacecatalog.ListEntitiesInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListEntitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEntities")
	}

	var r0 *marketplacecatalog.ListEntitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListEntitiesInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListEntitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListEntitiesInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.ListEntitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.ListEntitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.ListEntitiesInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *marketplacecatalog.ListTagsForResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListTagsForResourceOutput, error) {
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

	var r0 *marketplacecatalog.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListTagsForResourceInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.ListTagsForResourceInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.ListTagsForResourceInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() marketplacecatalog.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplacecatalog.Options
	if rf, ok := ret.Get(0).(func() marketplacecatalog.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplacecatalog.Options)
	}

	return r0
}

// PutResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutResourcePolicy(ctx context.Context, params *marketplacecatalog.PutResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.PutResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutResourcePolicy")
	}

	var r0 *marketplacecatalog.PutResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.PutResourcePolicyInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.PutResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.PutResourcePolicyInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.PutResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.PutResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.PutResourcePolicyInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartChangeSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartChangeSet(ctx context.Context, params *marketplacecatalog.StartChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.StartChangeSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartChangeSet")
	}

	var r0 *marketplacecatalog.StartChangeSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.StartChangeSetInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.StartChangeSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.StartChangeSetInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.StartChangeSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.StartChangeSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.StartChangeSetInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *marketplacecatalog.TagResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.TagResourceOutput, error) {
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

	var r0 *marketplacecatalog.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.TagResourceInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.TagResourceInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.TagResourceInput, ...func(*marketplacecatalog.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *marketplacecatalog.UntagResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.UntagResourceOutput, error) {
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

	var r0 *marketplacecatalog.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.UntagResourceInput, ...func(*marketplacecatalog.Options)) (*marketplacecatalog.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecatalog.UntagResourceInput, ...func(*marketplacecatalog.Options)) *marketplacecatalog.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecatalog.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecatalog.UntagResourceInput, ...func(*marketplacecatalog.Options)) error); ok {
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
