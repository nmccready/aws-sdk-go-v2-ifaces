// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	resourcegroups "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateGroup(ctx context.Context, params *resourcegroups.CreateGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.CreateGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateGroup")
	}

	var r0 *resourcegroups.CreateGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.CreateGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.CreateGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.CreateGroupInput, ...func(*resourcegroups.Options)) *resourcegroups.CreateGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.CreateGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.CreateGroupInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteGroup(ctx context.Context, params *resourcegroups.DeleteGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.DeleteGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroup")
	}

	var r0 *resourcegroups.DeleteGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.DeleteGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.DeleteGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.DeleteGroupInput, ...func(*resourcegroups.Options)) *resourcegroups.DeleteGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.DeleteGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.DeleteGroupInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAccountSettings(ctx context.Context, params *resourcegroups.GetAccountSettingsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetAccountSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountSettings")
	}

	var r0 *resourcegroups.GetAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetAccountSettingsInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetAccountSettingsInput, ...func(*resourcegroups.Options)) *resourcegroups.GetAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GetAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GetAccountSettingsInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroup(ctx context.Context, params *resourcegroups.GetGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroup")
	}

	var r0 *resourcegroups.GetGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupInput, ...func(*resourcegroups.Options)) *resourcegroups.GetGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GetGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GetGroupInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroupConfiguration(ctx context.Context, params *resourcegroups.GetGroupConfigurationInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupConfiguration")
	}

	var r0 *resourcegroups.GetGroupConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupConfigurationInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupConfigurationInput, ...func(*resourcegroups.Options)) *resourcegroups.GetGroupConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GetGroupConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GetGroupConfigurationInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroupQuery(ctx context.Context, params *resourcegroups.GetGroupQueryInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupQuery")
	}

	var r0 *resourcegroups.GetGroupQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupQueryInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetGroupQueryInput, ...func(*resourcegroups.Options)) *resourcegroups.GetGroupQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GetGroupQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GetGroupQueryInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTags(ctx context.Context, params *resourcegroups.GetTagsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

	var r0 *resourcegroups.GetTagsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetTagsInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GetTagsInput, ...func(*resourcegroups.Options)) *resourcegroups.GetTagsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GetTagsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GetTagsInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GroupResources(ctx context.Context, params *resourcegroups.GroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GroupResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GroupResources")
	}

	var r0 *resourcegroups.GroupResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GroupResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.GroupResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.GroupResourcesInput, ...func(*resourcegroups.Options)) *resourcegroups.GroupResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.GroupResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.GroupResourcesInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroupResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroupResources(ctx context.Context, params *resourcegroups.ListGroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupResources")
	}

	var r0 *resourcegroups.ListGroupResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.ListGroupResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.ListGroupResourcesInput, ...func(*resourcegroups.Options)) *resourcegroups.ListGroupResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.ListGroupResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.ListGroupResourcesInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroups(ctx context.Context, params *resourcegroups.ListGroupsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroups")
	}

	var r0 *resourcegroups.ListGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.ListGroupsInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.ListGroupsInput, ...func(*resourcegroups.Options)) *resourcegroups.ListGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.ListGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.ListGroupsInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() resourcegroups.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 resourcegroups.Options
	if rf, ok := ret.Get(0).(func() resourcegroups.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(resourcegroups.Options)
	}

	return r0
}

// PutGroupConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutGroupConfiguration(ctx context.Context, params *resourcegroups.PutGroupConfigurationInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.PutGroupConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutGroupConfiguration")
	}

	var r0 *resourcegroups.PutGroupConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.PutGroupConfigurationInput, ...func(*resourcegroups.Options)) (*resourcegroups.PutGroupConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.PutGroupConfigurationInput, ...func(*resourcegroups.Options)) *resourcegroups.PutGroupConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.PutGroupConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.PutGroupConfigurationInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SearchResources(ctx context.Context, params *resourcegroups.SearchResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SearchResources")
	}

	var r0 *resourcegroups.SearchResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.SearchResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.SearchResourcesInput, ...func(*resourcegroups.Options)) *resourcegroups.SearchResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.SearchResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.SearchResourcesInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tag provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Tag(ctx context.Context, params *resourcegroups.TagInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.TagOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Tag")
	}

	var r0 *resourcegroups.TagOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.TagInput, ...func(*resourcegroups.Options)) (*resourcegroups.TagOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.TagInput, ...func(*resourcegroups.Options)) *resourcegroups.TagOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.TagOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.TagInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UngroupResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UngroupResources(ctx context.Context, params *resourcegroups.UngroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UngroupResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UngroupResources")
	}

	var r0 *resourcegroups.UngroupResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UngroupResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.UngroupResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UngroupResourcesInput, ...func(*resourcegroups.Options)) *resourcegroups.UngroupResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.UngroupResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.UngroupResourcesInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Untag provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Untag(ctx context.Context, params *resourcegroups.UntagInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UntagOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Untag")
	}

	var r0 *resourcegroups.UntagOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UntagInput, ...func(*resourcegroups.Options)) (*resourcegroups.UntagOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UntagInput, ...func(*resourcegroups.Options)) *resourcegroups.UntagOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.UntagOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.UntagInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAccountSettings(ctx context.Context, params *resourcegroups.UpdateAccountSettingsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateAccountSettingsOutput, error) {
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

	var r0 *resourcegroups.UpdateAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateAccountSettingsInput, ...func(*resourcegroups.Options)) (*resourcegroups.UpdateAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateAccountSettingsInput, ...func(*resourcegroups.Options)) *resourcegroups.UpdateAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.UpdateAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.UpdateAccountSettingsInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateGroup(ctx context.Context, params *resourcegroups.UpdateGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGroup")
	}

	var r0 *resourcegroups.UpdateGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateGroupInput, ...func(*resourcegroups.Options)) *resourcegroups.UpdateGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.UpdateGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.UpdateGroupInput, ...func(*resourcegroups.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroupQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateGroupQuery(ctx context.Context, params *resourcegroups.UpdateGroupQueryInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGroupQuery")
	}

	var r0 *resourcegroups.UpdateGroupQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateGroupQueryInput, ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroups.UpdateGroupQueryInput, ...func(*resourcegroups.Options)) *resourcegroups.UpdateGroupQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroups.UpdateGroupQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroups.UpdateGroupQueryInput, ...func(*resourcegroups.Options)) error); ok {
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
