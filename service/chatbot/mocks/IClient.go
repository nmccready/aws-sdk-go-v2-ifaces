// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	chatbot "github.com/aws/aws-sdk-go-v2/service/chatbot"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateChimeWebhookConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateChimeWebhookConfiguration(ctx context.Context, params *chatbot.CreateChimeWebhookConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.CreateChimeWebhookConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateChimeWebhookConfiguration")
	}

	var r0 *chatbot.CreateChimeWebhookConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) (*chatbot.CreateChimeWebhookConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) *chatbot.CreateChimeWebhookConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.CreateChimeWebhookConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.CreateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMicrosoftTeamsChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *chatbot.CreateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.CreateMicrosoftTeamsChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMicrosoftTeamsChannelConfiguration")
	}

	var r0 *chatbot.CreateMicrosoftTeamsChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.CreateMicrosoftTeamsChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.CreateMicrosoftTeamsChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.CreateMicrosoftTeamsChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.CreateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSlackChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSlackChannelConfiguration(ctx context.Context, params *chatbot.CreateSlackChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.CreateSlackChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSlackChannelConfiguration")
	}

	var r0 *chatbot.CreateSlackChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateSlackChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.CreateSlackChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.CreateSlackChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.CreateSlackChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.CreateSlackChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.CreateSlackChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChimeWebhookConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteChimeWebhookConfiguration(ctx context.Context, params *chatbot.DeleteChimeWebhookConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteChimeWebhookConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChimeWebhookConfiguration")
	}

	var r0 *chatbot.DeleteChimeWebhookConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteChimeWebhookConfigurationInput, ...func(*chatbot.Options)) (*chatbot.DeleteChimeWebhookConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteChimeWebhookConfigurationInput, ...func(*chatbot.Options)) *chatbot.DeleteChimeWebhookConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteChimeWebhookConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteChimeWebhookConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMicrosoftTeamsChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMicrosoftTeamsChannelConfiguration(ctx context.Context, params *chatbot.DeleteMicrosoftTeamsChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMicrosoftTeamsChannelConfiguration")
	}

	var r0 *chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMicrosoftTeamsConfiguredTeam provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMicrosoftTeamsConfiguredTeam(ctx context.Context, params *chatbot.DeleteMicrosoftTeamsConfiguredTeamInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMicrosoftTeamsConfiguredTeam")
	}

	var r0 *chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsConfiguredTeamInput, ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsConfiguredTeamInput, ...func(*chatbot.Options)) *chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteMicrosoftTeamsConfiguredTeamInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMicrosoftTeamsUserIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMicrosoftTeamsUserIdentity(ctx context.Context, params *chatbot.DeleteMicrosoftTeamsUserIdentityInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsUserIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMicrosoftTeamsUserIdentity")
	}

	var r0 *chatbot.DeleteMicrosoftTeamsUserIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsUserIdentityInput, ...func(*chatbot.Options)) (*chatbot.DeleteMicrosoftTeamsUserIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteMicrosoftTeamsUserIdentityInput, ...func(*chatbot.Options)) *chatbot.DeleteMicrosoftTeamsUserIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteMicrosoftTeamsUserIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteMicrosoftTeamsUserIdentityInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSlackChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSlackChannelConfiguration(ctx context.Context, params *chatbot.DeleteSlackChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteSlackChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSlackChannelConfiguration")
	}

	var r0 *chatbot.DeleteSlackChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.DeleteSlackChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.DeleteSlackChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteSlackChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteSlackChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSlackUserIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSlackUserIdentity(ctx context.Context, params *chatbot.DeleteSlackUserIdentityInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteSlackUserIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSlackUserIdentity")
	}

	var r0 *chatbot.DeleteSlackUserIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackUserIdentityInput, ...func(*chatbot.Options)) (*chatbot.DeleteSlackUserIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackUserIdentityInput, ...func(*chatbot.Options)) *chatbot.DeleteSlackUserIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteSlackUserIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteSlackUserIdentityInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSlackWorkspaceAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSlackWorkspaceAuthorization(ctx context.Context, params *chatbot.DeleteSlackWorkspaceAuthorizationInput, optFns ...func(*chatbot.Options)) (*chatbot.DeleteSlackWorkspaceAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSlackWorkspaceAuthorization")
	}

	var r0 *chatbot.DeleteSlackWorkspaceAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackWorkspaceAuthorizationInput, ...func(*chatbot.Options)) (*chatbot.DeleteSlackWorkspaceAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DeleteSlackWorkspaceAuthorizationInput, ...func(*chatbot.Options)) *chatbot.DeleteSlackWorkspaceAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DeleteSlackWorkspaceAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DeleteSlackWorkspaceAuthorizationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeChimeWebhookConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeChimeWebhookConfigurations(ctx context.Context, params *chatbot.DescribeChimeWebhookConfigurationsInput, optFns ...func(*chatbot.Options)) (*chatbot.DescribeChimeWebhookConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeChimeWebhookConfigurations")
	}

	var r0 *chatbot.DescribeChimeWebhookConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeChimeWebhookConfigurationsInput, ...func(*chatbot.Options)) (*chatbot.DescribeChimeWebhookConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeChimeWebhookConfigurationsInput, ...func(*chatbot.Options)) *chatbot.DescribeChimeWebhookConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DescribeChimeWebhookConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DescribeChimeWebhookConfigurationsInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSlackChannelConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSlackChannelConfigurations(ctx context.Context, params *chatbot.DescribeSlackChannelConfigurationsInput, optFns ...func(*chatbot.Options)) (*chatbot.DescribeSlackChannelConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSlackChannelConfigurations")
	}

	var r0 *chatbot.DescribeSlackChannelConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackChannelConfigurationsInput, ...func(*chatbot.Options)) (*chatbot.DescribeSlackChannelConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackChannelConfigurationsInput, ...func(*chatbot.Options)) *chatbot.DescribeSlackChannelConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DescribeSlackChannelConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DescribeSlackChannelConfigurationsInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSlackUserIdentities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSlackUserIdentities(ctx context.Context, params *chatbot.DescribeSlackUserIdentitiesInput, optFns ...func(*chatbot.Options)) (*chatbot.DescribeSlackUserIdentitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSlackUserIdentities")
	}

	var r0 *chatbot.DescribeSlackUserIdentitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackUserIdentitiesInput, ...func(*chatbot.Options)) (*chatbot.DescribeSlackUserIdentitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackUserIdentitiesInput, ...func(*chatbot.Options)) *chatbot.DescribeSlackUserIdentitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DescribeSlackUserIdentitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DescribeSlackUserIdentitiesInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSlackWorkspaces provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSlackWorkspaces(ctx context.Context, params *chatbot.DescribeSlackWorkspacesInput, optFns ...func(*chatbot.Options)) (*chatbot.DescribeSlackWorkspacesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSlackWorkspaces")
	}

	var r0 *chatbot.DescribeSlackWorkspacesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackWorkspacesInput, ...func(*chatbot.Options)) (*chatbot.DescribeSlackWorkspacesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.DescribeSlackWorkspacesInput, ...func(*chatbot.Options)) *chatbot.DescribeSlackWorkspacesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.DescribeSlackWorkspacesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.DescribeSlackWorkspacesInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountPreferences provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAccountPreferences(ctx context.Context, params *chatbot.GetAccountPreferencesInput, optFns ...func(*chatbot.Options)) (*chatbot.GetAccountPreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountPreferences")
	}

	var r0 *chatbot.GetAccountPreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.GetAccountPreferencesInput, ...func(*chatbot.Options)) (*chatbot.GetAccountPreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.GetAccountPreferencesInput, ...func(*chatbot.Options)) *chatbot.GetAccountPreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.GetAccountPreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.GetAccountPreferencesInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMicrosoftTeamsChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMicrosoftTeamsChannelConfiguration(ctx context.Context, params *chatbot.GetMicrosoftTeamsChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.GetMicrosoftTeamsChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMicrosoftTeamsChannelConfiguration")
	}

	var r0 *chatbot.GetMicrosoftTeamsChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.GetMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.GetMicrosoftTeamsChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.GetMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.GetMicrosoftTeamsChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.GetMicrosoftTeamsChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.GetMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMicrosoftTeamsChannelConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListMicrosoftTeamsChannelConfigurations(ctx context.Context, params *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, optFns ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMicrosoftTeamsChannelConfigurations")
	}

	var r0 *chatbot.ListMicrosoftTeamsChannelConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, ...func(*chatbot.Options)) *chatbot.ListMicrosoftTeamsChannelConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMicrosoftTeamsConfiguredTeams provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListMicrosoftTeamsConfiguredTeams(ctx context.Context, params *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, optFns ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMicrosoftTeamsConfiguredTeams")
	}

	var r0 *chatbot.ListMicrosoftTeamsConfiguredTeamsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, ...func(*chatbot.Options)) *chatbot.ListMicrosoftTeamsConfiguredTeamsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMicrosoftTeamsUserIdentities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListMicrosoftTeamsUserIdentities(ctx context.Context, params *chatbot.ListMicrosoftTeamsUserIdentitiesInput, optFns ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMicrosoftTeamsUserIdentities")
	}

	var r0 *chatbot.ListMicrosoftTeamsUserIdentitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsUserIdentitiesInput, ...func(*chatbot.Options)) (*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListMicrosoftTeamsUserIdentitiesInput, ...func(*chatbot.Options)) *chatbot.ListMicrosoftTeamsUserIdentitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.ListMicrosoftTeamsUserIdentitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.ListMicrosoftTeamsUserIdentitiesInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *chatbot.ListTagsForResourceInput, optFns ...func(*chatbot.Options)) (*chatbot.ListTagsForResourceOutput, error) {
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

	var r0 *chatbot.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListTagsForResourceInput, ...func(*chatbot.Options)) (*chatbot.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.ListTagsForResourceInput, ...func(*chatbot.Options)) *chatbot.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.ListTagsForResourceInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() chatbot.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 chatbot.Options
	if rf, ok := ret.Get(0).(func() chatbot.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(chatbot.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *chatbot.TagResourceInput, optFns ...func(*chatbot.Options)) (*chatbot.TagResourceOutput, error) {
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

	var r0 *chatbot.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.TagResourceInput, ...func(*chatbot.Options)) (*chatbot.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.TagResourceInput, ...func(*chatbot.Options)) *chatbot.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.TagResourceInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *chatbot.UntagResourceInput, optFns ...func(*chatbot.Options)) (*chatbot.UntagResourceOutput, error) {
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

	var r0 *chatbot.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UntagResourceInput, ...func(*chatbot.Options)) (*chatbot.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UntagResourceInput, ...func(*chatbot.Options)) *chatbot.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.UntagResourceInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountPreferences provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAccountPreferences(ctx context.Context, params *chatbot.UpdateAccountPreferencesInput, optFns ...func(*chatbot.Options)) (*chatbot.UpdateAccountPreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountPreferences")
	}

	var r0 *chatbot.UpdateAccountPreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateAccountPreferencesInput, ...func(*chatbot.Options)) (*chatbot.UpdateAccountPreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateAccountPreferencesInput, ...func(*chatbot.Options)) *chatbot.UpdateAccountPreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.UpdateAccountPreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.UpdateAccountPreferencesInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChimeWebhookConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateChimeWebhookConfiguration(ctx context.Context, params *chatbot.UpdateChimeWebhookConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.UpdateChimeWebhookConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChimeWebhookConfiguration")
	}

	var r0 *chatbot.UpdateChimeWebhookConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) (*chatbot.UpdateChimeWebhookConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) *chatbot.UpdateChimeWebhookConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.UpdateChimeWebhookConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.UpdateChimeWebhookConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMicrosoftTeamsChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *chatbot.UpdateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMicrosoftTeamsChannelConfiguration")
	}

	var r0 *chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.UpdateMicrosoftTeamsChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSlackChannelConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSlackChannelConfiguration(ctx context.Context, params *chatbot.UpdateSlackChannelConfigurationInput, optFns ...func(*chatbot.Options)) (*chatbot.UpdateSlackChannelConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSlackChannelConfiguration")
	}

	var r0 *chatbot.UpdateSlackChannelConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateSlackChannelConfigurationInput, ...func(*chatbot.Options)) (*chatbot.UpdateSlackChannelConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chatbot.UpdateSlackChannelConfigurationInput, ...func(*chatbot.Options)) *chatbot.UpdateSlackChannelConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chatbot.UpdateSlackChannelConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chatbot.UpdateSlackChannelConfigurationInput, ...func(*chatbot.Options)) error); ok {
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
