// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	timestreamquery "github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelQuery(ctx context.Context, params *timestreamquery.CancelQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.CancelQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelQuery")
	}

	var r0 *timestreamquery.CancelQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.CancelQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.CancelQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.CancelQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.CancelQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.CancelQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.CancelQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScheduledQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateScheduledQuery(ctx context.Context, params *timestreamquery.CreateScheduledQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.CreateScheduledQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateScheduledQuery")
	}

	var r0 *timestreamquery.CreateScheduledQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.CreateScheduledQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.CreateScheduledQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.CreateScheduledQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.CreateScheduledQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.CreateScheduledQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.CreateScheduledQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScheduledQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteScheduledQuery(ctx context.Context, params *timestreamquery.DeleteScheduledQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.DeleteScheduledQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduledQuery")
	}

	var r0 *timestreamquery.DeleteScheduledQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DeleteScheduledQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.DeleteScheduledQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DeleteScheduledQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.DeleteScheduledQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.DeleteScheduledQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.DeleteScheduledQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAccountSettings(ctx context.Context, params *timestreamquery.DescribeAccountSettingsInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.DescribeAccountSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAccountSettings")
	}

	var r0 *timestreamquery.DescribeAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeAccountSettingsInput, ...func(*timestreamquery.Options)) (*timestreamquery.DescribeAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeAccountSettingsInput, ...func(*timestreamquery.Options)) *timestreamquery.DescribeAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.DescribeAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.DescribeAccountSettingsInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeEndpoints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeEndpoints(ctx context.Context, params *timestreamquery.DescribeEndpointsInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.DescribeEndpointsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeEndpoints")
	}

	var r0 *timestreamquery.DescribeEndpointsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeEndpointsInput, ...func(*timestreamquery.Options)) (*timestreamquery.DescribeEndpointsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeEndpointsInput, ...func(*timestreamquery.Options)) *timestreamquery.DescribeEndpointsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.DescribeEndpointsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.DescribeEndpointsInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScheduledQuery(ctx context.Context, params *timestreamquery.DescribeScheduledQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.DescribeScheduledQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScheduledQuery")
	}

	var r0 *timestreamquery.DescribeScheduledQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeScheduledQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.DescribeScheduledQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.DescribeScheduledQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.DescribeScheduledQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.DescribeScheduledQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.DescribeScheduledQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScheduledQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExecuteScheduledQuery(ctx context.Context, params *timestreamquery.ExecuteScheduledQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.ExecuteScheduledQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteScheduledQuery")
	}

	var r0 *timestreamquery.ExecuteScheduledQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ExecuteScheduledQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.ExecuteScheduledQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ExecuteScheduledQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.ExecuteScheduledQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.ExecuteScheduledQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.ExecuteScheduledQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScheduledQueries provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListScheduledQueries(ctx context.Context, params *timestreamquery.ListScheduledQueriesInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.ListScheduledQueriesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListScheduledQueries")
	}

	var r0 *timestreamquery.ListScheduledQueriesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ListScheduledQueriesInput, ...func(*timestreamquery.Options)) (*timestreamquery.ListScheduledQueriesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ListScheduledQueriesInput, ...func(*timestreamquery.Options)) *timestreamquery.ListScheduledQueriesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.ListScheduledQueriesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.ListScheduledQueriesInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *timestreamquery.ListTagsForResourceInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.ListTagsForResourceOutput, error) {
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

	var r0 *timestreamquery.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ListTagsForResourceInput, ...func(*timestreamquery.Options)) (*timestreamquery.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.ListTagsForResourceInput, ...func(*timestreamquery.Options)) *timestreamquery.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.ListTagsForResourceInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() timestreamquery.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 timestreamquery.Options
	if rf, ok := ret.Get(0).(func() timestreamquery.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(timestreamquery.Options)
	}

	return r0
}

// PrepareQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PrepareQuery(ctx context.Context, params *timestreamquery.PrepareQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.PrepareQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PrepareQuery")
	}

	var r0 *timestreamquery.PrepareQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.PrepareQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.PrepareQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.PrepareQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.PrepareQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.PrepareQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.PrepareQueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Query(ctx context.Context, params *timestreamquery.QueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.QueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *timestreamquery.QueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.QueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.QueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.QueryInput, ...func(*timestreamquery.Options)) *timestreamquery.QueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.QueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.QueryInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *timestreamquery.TagResourceInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.TagResourceOutput, error) {
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

	var r0 *timestreamquery.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.TagResourceInput, ...func(*timestreamquery.Options)) (*timestreamquery.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.TagResourceInput, ...func(*timestreamquery.Options)) *timestreamquery.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.TagResourceInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *timestreamquery.UntagResourceInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.UntagResourceOutput, error) {
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

	var r0 *timestreamquery.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UntagResourceInput, ...func(*timestreamquery.Options)) (*timestreamquery.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UntagResourceInput, ...func(*timestreamquery.Options)) *timestreamquery.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.UntagResourceInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAccountSettings(ctx context.Context, params *timestreamquery.UpdateAccountSettingsInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.UpdateAccountSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountSettings")
	}

	var r0 *timestreamquery.UpdateAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UpdateAccountSettingsInput, ...func(*timestreamquery.Options)) (*timestreamquery.UpdateAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UpdateAccountSettingsInput, ...func(*timestreamquery.Options)) *timestreamquery.UpdateAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.UpdateAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.UpdateAccountSettingsInput, ...func(*timestreamquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScheduledQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateScheduledQuery(ctx context.Context, params *timestreamquery.UpdateScheduledQueryInput, optFns ...func(*timestreamquery.Options)) (*timestreamquery.UpdateScheduledQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateScheduledQuery")
	}

	var r0 *timestreamquery.UpdateScheduledQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UpdateScheduledQueryInput, ...func(*timestreamquery.Options)) (*timestreamquery.UpdateScheduledQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *timestreamquery.UpdateScheduledQueryInput, ...func(*timestreamquery.Options)) *timestreamquery.UpdateScheduledQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*timestreamquery.UpdateScheduledQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *timestreamquery.UpdateScheduledQueryInput, ...func(*timestreamquery.Options)) error); ok {
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
