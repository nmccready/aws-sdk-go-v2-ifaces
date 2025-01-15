// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	keyspaces "github.com/aws/aws-sdk-go-v2/service/keyspaces"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateKeyspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateKeyspace(ctx context.Context, params *keyspaces.CreateKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.CreateKeyspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateKeyspace")
	}

	var r0 *keyspaces.CreateKeyspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.CreateKeyspaceInput, ...func(*keyspaces.Options)) (*keyspaces.CreateKeyspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.CreateKeyspaceInput, ...func(*keyspaces.Options)) *keyspaces.CreateKeyspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.CreateKeyspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.CreateKeyspaceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTable(ctx context.Context, params *keyspaces.CreateTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.CreateTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTable")
	}

	var r0 *keyspaces.CreateTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.CreateTableInput, ...func(*keyspaces.Options)) (*keyspaces.CreateTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.CreateTableInput, ...func(*keyspaces.Options)) *keyspaces.CreateTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.CreateTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.CreateTableInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteKeyspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteKeyspace(ctx context.Context, params *keyspaces.DeleteKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.DeleteKeyspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteKeyspace")
	}

	var r0 *keyspaces.DeleteKeyspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.DeleteKeyspaceInput, ...func(*keyspaces.Options)) (*keyspaces.DeleteKeyspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.DeleteKeyspaceInput, ...func(*keyspaces.Options)) *keyspaces.DeleteKeyspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.DeleteKeyspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.DeleteKeyspaceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTable(ctx context.Context, params *keyspaces.DeleteTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.DeleteTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTable")
	}

	var r0 *keyspaces.DeleteTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.DeleteTableInput, ...func(*keyspaces.Options)) (*keyspaces.DeleteTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.DeleteTableInput, ...func(*keyspaces.Options)) *keyspaces.DeleteTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.DeleteTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.DeleteTableInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeyspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetKeyspace(ctx context.Context, params *keyspaces.GetKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetKeyspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetKeyspace")
	}

	var r0 *keyspaces.GetKeyspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetKeyspaceInput, ...func(*keyspaces.Options)) (*keyspaces.GetKeyspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetKeyspaceInput, ...func(*keyspaces.Options)) *keyspaces.GetKeyspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.GetKeyspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.GetKeyspaceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTable(ctx context.Context, params *keyspaces.GetTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTable")
	}

	var r0 *keyspaces.GetTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetTableInput, ...func(*keyspaces.Options)) (*keyspaces.GetTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetTableInput, ...func(*keyspaces.Options)) *keyspaces.GetTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.GetTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.GetTableInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTableAutoScalingSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTableAutoScalingSettings(ctx context.Context, params *keyspaces.GetTableAutoScalingSettingsInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetTableAutoScalingSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTableAutoScalingSettings")
	}

	var r0 *keyspaces.GetTableAutoScalingSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetTableAutoScalingSettingsInput, ...func(*keyspaces.Options)) (*keyspaces.GetTableAutoScalingSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.GetTableAutoScalingSettingsInput, ...func(*keyspaces.Options)) *keyspaces.GetTableAutoScalingSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.GetTableAutoScalingSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.GetTableAutoScalingSettingsInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListKeyspaces provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListKeyspaces(ctx context.Context, params *keyspaces.ListKeyspacesInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListKeyspacesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListKeyspaces")
	}

	var r0 *keyspaces.ListKeyspacesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListKeyspacesInput, ...func(*keyspaces.Options)) (*keyspaces.ListKeyspacesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListKeyspacesInput, ...func(*keyspaces.Options)) *keyspaces.ListKeyspacesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.ListKeyspacesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.ListKeyspacesInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTables provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTables(ctx context.Context, params *keyspaces.ListTablesInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListTablesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTables")
	}

	var r0 *keyspaces.ListTablesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListTablesInput, ...func(*keyspaces.Options)) (*keyspaces.ListTablesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListTablesInput, ...func(*keyspaces.Options)) *keyspaces.ListTablesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.ListTablesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.ListTablesInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *keyspaces.ListTagsForResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListTagsForResourceOutput, error) {
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

	var r0 *keyspaces.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListTagsForResourceInput, ...func(*keyspaces.Options)) (*keyspaces.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.ListTagsForResourceInput, ...func(*keyspaces.Options)) *keyspaces.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.ListTagsForResourceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() keyspaces.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 keyspaces.Options
	if rf, ok := ret.Get(0).(func() keyspaces.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(keyspaces.Options)
	}

	return r0
}

// RestoreTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RestoreTable(ctx context.Context, params *keyspaces.RestoreTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.RestoreTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RestoreTable")
	}

	var r0 *keyspaces.RestoreTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.RestoreTableInput, ...func(*keyspaces.Options)) (*keyspaces.RestoreTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.RestoreTableInput, ...func(*keyspaces.Options)) *keyspaces.RestoreTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.RestoreTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.RestoreTableInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *keyspaces.TagResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.TagResourceOutput, error) {
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

	var r0 *keyspaces.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.TagResourceInput, ...func(*keyspaces.Options)) (*keyspaces.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.TagResourceInput, ...func(*keyspaces.Options)) *keyspaces.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.TagResourceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *keyspaces.UntagResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.UntagResourceOutput, error) {
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

	var r0 *keyspaces.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.UntagResourceInput, ...func(*keyspaces.Options)) (*keyspaces.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.UntagResourceInput, ...func(*keyspaces.Options)) *keyspaces.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.UntagResourceInput, ...func(*keyspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTable(ctx context.Context, params *keyspaces.UpdateTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.UpdateTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTable")
	}

	var r0 *keyspaces.UpdateTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.UpdateTableInput, ...func(*keyspaces.Options)) (*keyspaces.UpdateTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *keyspaces.UpdateTableInput, ...func(*keyspaces.Options)) *keyspaces.UpdateTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keyspaces.UpdateTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *keyspaces.UpdateTableInput, ...func(*keyspaces.Options)) error); ok {
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
