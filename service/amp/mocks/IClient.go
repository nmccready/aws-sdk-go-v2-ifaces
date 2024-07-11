// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	amp "github.com/aws/aws-sdk-go-v2/service/amp"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateAlertManagerDefinition provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAlertManagerDefinition(ctx context.Context, params *amp.CreateAlertManagerDefinitionInput, optFns ...func(*amp.Options)) (*amp.CreateAlertManagerDefinitionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAlertManagerDefinition")
	}

	var r0 *amp.CreateAlertManagerDefinitionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateAlertManagerDefinitionInput, ...func(*amp.Options)) (*amp.CreateAlertManagerDefinitionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateAlertManagerDefinitionInput, ...func(*amp.Options)) *amp.CreateAlertManagerDefinitionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.CreateAlertManagerDefinitionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.CreateAlertManagerDefinitionInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateLoggingConfiguration(ctx context.Context, params *amp.CreateLoggingConfigurationInput, optFns ...func(*amp.Options)) (*amp.CreateLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLoggingConfiguration")
	}

	var r0 *amp.CreateLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateLoggingConfigurationInput, ...func(*amp.Options)) (*amp.CreateLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateLoggingConfigurationInput, ...func(*amp.Options)) *amp.CreateLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.CreateLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.CreateLoggingConfigurationInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRuleGroupsNamespace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRuleGroupsNamespace(ctx context.Context, params *amp.CreateRuleGroupsNamespaceInput, optFns ...func(*amp.Options)) (*amp.CreateRuleGroupsNamespaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRuleGroupsNamespace")
	}

	var r0 *amp.CreateRuleGroupsNamespaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateRuleGroupsNamespaceInput, ...func(*amp.Options)) (*amp.CreateRuleGroupsNamespaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateRuleGroupsNamespaceInput, ...func(*amp.Options)) *amp.CreateRuleGroupsNamespaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.CreateRuleGroupsNamespaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.CreateRuleGroupsNamespaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScraper provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateScraper(ctx context.Context, params *amp.CreateScraperInput, optFns ...func(*amp.Options)) (*amp.CreateScraperOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateScraper")
	}

	var r0 *amp.CreateScraperOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateScraperInput, ...func(*amp.Options)) (*amp.CreateScraperOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateScraperInput, ...func(*amp.Options)) *amp.CreateScraperOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.CreateScraperOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.CreateScraperInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateWorkspace(ctx context.Context, params *amp.CreateWorkspaceInput, optFns ...func(*amp.Options)) (*amp.CreateWorkspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateWorkspace")
	}

	var r0 *amp.CreateWorkspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateWorkspaceInput, ...func(*amp.Options)) (*amp.CreateWorkspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.CreateWorkspaceInput, ...func(*amp.Options)) *amp.CreateWorkspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.CreateWorkspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.CreateWorkspaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAlertManagerDefinition provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAlertManagerDefinition(ctx context.Context, params *amp.DeleteAlertManagerDefinitionInput, optFns ...func(*amp.Options)) (*amp.DeleteAlertManagerDefinitionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAlertManagerDefinition")
	}

	var r0 *amp.DeleteAlertManagerDefinitionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteAlertManagerDefinitionInput, ...func(*amp.Options)) (*amp.DeleteAlertManagerDefinitionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteAlertManagerDefinitionInput, ...func(*amp.Options)) *amp.DeleteAlertManagerDefinitionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DeleteAlertManagerDefinitionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DeleteAlertManagerDefinitionInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteLoggingConfiguration(ctx context.Context, params *amp.DeleteLoggingConfigurationInput, optFns ...func(*amp.Options)) (*amp.DeleteLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLoggingConfiguration")
	}

	var r0 *amp.DeleteLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteLoggingConfigurationInput, ...func(*amp.Options)) (*amp.DeleteLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteLoggingConfigurationInput, ...func(*amp.Options)) *amp.DeleteLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DeleteLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DeleteLoggingConfigurationInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRuleGroupsNamespace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRuleGroupsNamespace(ctx context.Context, params *amp.DeleteRuleGroupsNamespaceInput, optFns ...func(*amp.Options)) (*amp.DeleteRuleGroupsNamespaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRuleGroupsNamespace")
	}

	var r0 *amp.DeleteRuleGroupsNamespaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteRuleGroupsNamespaceInput, ...func(*amp.Options)) (*amp.DeleteRuleGroupsNamespaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteRuleGroupsNamespaceInput, ...func(*amp.Options)) *amp.DeleteRuleGroupsNamespaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DeleteRuleGroupsNamespaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DeleteRuleGroupsNamespaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScraper provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteScraper(ctx context.Context, params *amp.DeleteScraperInput, optFns ...func(*amp.Options)) (*amp.DeleteScraperOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScraper")
	}

	var r0 *amp.DeleteScraperOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteScraperInput, ...func(*amp.Options)) (*amp.DeleteScraperOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteScraperInput, ...func(*amp.Options)) *amp.DeleteScraperOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DeleteScraperOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DeleteScraperInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteWorkspace(ctx context.Context, params *amp.DeleteWorkspaceInput, optFns ...func(*amp.Options)) (*amp.DeleteWorkspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWorkspace")
	}

	var r0 *amp.DeleteWorkspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteWorkspaceInput, ...func(*amp.Options)) (*amp.DeleteWorkspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DeleteWorkspaceInput, ...func(*amp.Options)) *amp.DeleteWorkspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DeleteWorkspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DeleteWorkspaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAlertManagerDefinition provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAlertManagerDefinition(ctx context.Context, params *amp.DescribeAlertManagerDefinitionInput, optFns ...func(*amp.Options)) (*amp.DescribeAlertManagerDefinitionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAlertManagerDefinition")
	}

	var r0 *amp.DescribeAlertManagerDefinitionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeAlertManagerDefinitionInput, ...func(*amp.Options)) (*amp.DescribeAlertManagerDefinitionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeAlertManagerDefinitionInput, ...func(*amp.Options)) *amp.DescribeAlertManagerDefinitionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DescribeAlertManagerDefinitionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DescribeAlertManagerDefinitionInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeLoggingConfiguration(ctx context.Context, params *amp.DescribeLoggingConfigurationInput, optFns ...func(*amp.Options)) (*amp.DescribeLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeLoggingConfiguration")
	}

	var r0 *amp.DescribeLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeLoggingConfigurationInput, ...func(*amp.Options)) (*amp.DescribeLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeLoggingConfigurationInput, ...func(*amp.Options)) *amp.DescribeLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DescribeLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DescribeLoggingConfigurationInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRuleGroupsNamespace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeRuleGroupsNamespace(ctx context.Context, params *amp.DescribeRuleGroupsNamespaceInput, optFns ...func(*amp.Options)) (*amp.DescribeRuleGroupsNamespaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRuleGroupsNamespace")
	}

	var r0 *amp.DescribeRuleGroupsNamespaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeRuleGroupsNamespaceInput, ...func(*amp.Options)) (*amp.DescribeRuleGroupsNamespaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeRuleGroupsNamespaceInput, ...func(*amp.Options)) *amp.DescribeRuleGroupsNamespaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DescribeRuleGroupsNamespaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DescribeRuleGroupsNamespaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScraper provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScraper(ctx context.Context, params *amp.DescribeScraperInput, optFns ...func(*amp.Options)) (*amp.DescribeScraperOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScraper")
	}

	var r0 *amp.DescribeScraperOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeScraperInput, ...func(*amp.Options)) (*amp.DescribeScraperOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeScraperInput, ...func(*amp.Options)) *amp.DescribeScraperOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DescribeScraperOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DescribeScraperInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeWorkspace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeWorkspace(ctx context.Context, params *amp.DescribeWorkspaceInput, optFns ...func(*amp.Options)) (*amp.DescribeWorkspaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeWorkspace")
	}

	var r0 *amp.DescribeWorkspaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeWorkspaceInput, ...func(*amp.Options)) (*amp.DescribeWorkspaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.DescribeWorkspaceInput, ...func(*amp.Options)) *amp.DescribeWorkspaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.DescribeWorkspaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.DescribeWorkspaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultScraperConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDefaultScraperConfiguration(ctx context.Context, params *amp.GetDefaultScraperConfigurationInput, optFns ...func(*amp.Options)) (*amp.GetDefaultScraperConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultScraperConfiguration")
	}

	var r0 *amp.GetDefaultScraperConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.GetDefaultScraperConfigurationInput, ...func(*amp.Options)) (*amp.GetDefaultScraperConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.GetDefaultScraperConfigurationInput, ...func(*amp.Options)) *amp.GetDefaultScraperConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.GetDefaultScraperConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.GetDefaultScraperConfigurationInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRuleGroupsNamespaces provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRuleGroupsNamespaces(ctx context.Context, params *amp.ListRuleGroupsNamespacesInput, optFns ...func(*amp.Options)) (*amp.ListRuleGroupsNamespacesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRuleGroupsNamespaces")
	}

	var r0 *amp.ListRuleGroupsNamespacesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListRuleGroupsNamespacesInput, ...func(*amp.Options)) (*amp.ListRuleGroupsNamespacesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListRuleGroupsNamespacesInput, ...func(*amp.Options)) *amp.ListRuleGroupsNamespacesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.ListRuleGroupsNamespacesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.ListRuleGroupsNamespacesInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScrapers provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListScrapers(ctx context.Context, params *amp.ListScrapersInput, optFns ...func(*amp.Options)) (*amp.ListScrapersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListScrapers")
	}

	var r0 *amp.ListScrapersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListScrapersInput, ...func(*amp.Options)) (*amp.ListScrapersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListScrapersInput, ...func(*amp.Options)) *amp.ListScrapersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.ListScrapersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.ListScrapersInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *amp.ListTagsForResourceInput, optFns ...func(*amp.Options)) (*amp.ListTagsForResourceOutput, error) {
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

	var r0 *amp.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListTagsForResourceInput, ...func(*amp.Options)) (*amp.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListTagsForResourceInput, ...func(*amp.Options)) *amp.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.ListTagsForResourceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaces provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListWorkspaces(ctx context.Context, params *amp.ListWorkspacesInput, optFns ...func(*amp.Options)) (*amp.ListWorkspacesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListWorkspaces")
	}

	var r0 *amp.ListWorkspacesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListWorkspacesInput, ...func(*amp.Options)) (*amp.ListWorkspacesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.ListWorkspacesInput, ...func(*amp.Options)) *amp.ListWorkspacesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.ListWorkspacesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.ListWorkspacesInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() amp.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 amp.Options
	if rf, ok := ret.Get(0).(func() amp.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(amp.Options)
	}

	return r0
}

// PutAlertManagerDefinition provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutAlertManagerDefinition(ctx context.Context, params *amp.PutAlertManagerDefinitionInput, optFns ...func(*amp.Options)) (*amp.PutAlertManagerDefinitionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutAlertManagerDefinition")
	}

	var r0 *amp.PutAlertManagerDefinitionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.PutAlertManagerDefinitionInput, ...func(*amp.Options)) (*amp.PutAlertManagerDefinitionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.PutAlertManagerDefinitionInput, ...func(*amp.Options)) *amp.PutAlertManagerDefinitionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.PutAlertManagerDefinitionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.PutAlertManagerDefinitionInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutRuleGroupsNamespace provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutRuleGroupsNamespace(ctx context.Context, params *amp.PutRuleGroupsNamespaceInput, optFns ...func(*amp.Options)) (*amp.PutRuleGroupsNamespaceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRuleGroupsNamespace")
	}

	var r0 *amp.PutRuleGroupsNamespaceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.PutRuleGroupsNamespaceInput, ...func(*amp.Options)) (*amp.PutRuleGroupsNamespaceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.PutRuleGroupsNamespaceInput, ...func(*amp.Options)) *amp.PutRuleGroupsNamespaceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.PutRuleGroupsNamespaceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.PutRuleGroupsNamespaceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *amp.TagResourceInput, optFns ...func(*amp.Options)) (*amp.TagResourceOutput, error) {
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

	var r0 *amp.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.TagResourceInput, ...func(*amp.Options)) (*amp.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.TagResourceInput, ...func(*amp.Options)) *amp.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.TagResourceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *amp.UntagResourceInput, optFns ...func(*amp.Options)) (*amp.UntagResourceOutput, error) {
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

	var r0 *amp.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UntagResourceInput, ...func(*amp.Options)) (*amp.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UntagResourceInput, ...func(*amp.Options)) *amp.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.UntagResourceInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateLoggingConfiguration(ctx context.Context, params *amp.UpdateLoggingConfigurationInput, optFns ...func(*amp.Options)) (*amp.UpdateLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLoggingConfiguration")
	}

	var r0 *amp.UpdateLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UpdateLoggingConfigurationInput, ...func(*amp.Options)) (*amp.UpdateLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UpdateLoggingConfigurationInput, ...func(*amp.Options)) *amp.UpdateLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.UpdateLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.UpdateLoggingConfigurationInput, ...func(*amp.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkspaceAlias provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateWorkspaceAlias(ctx context.Context, params *amp.UpdateWorkspaceAliasInput, optFns ...func(*amp.Options)) (*amp.UpdateWorkspaceAliasOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWorkspaceAlias")
	}

	var r0 *amp.UpdateWorkspaceAliasOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UpdateWorkspaceAliasInput, ...func(*amp.Options)) (*amp.UpdateWorkspaceAliasOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amp.UpdateWorkspaceAliasInput, ...func(*amp.Options)) *amp.UpdateWorkspaceAliasOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amp.UpdateWorkspaceAliasOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amp.UpdateWorkspaceAliasInput, ...func(*amp.Options)) error); ok {
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