// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	kafkaconnect "github.com/aws/aws-sdk-go-v2/service/kafkaconnect"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateConnector(ctx context.Context, params *kafkaconnect.CreateConnectorInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateConnector")
	}

	var r0 *kafkaconnect.CreateConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateConnectorInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateConnectorInput, ...func(*kafkaconnect.Options)) *kafkaconnect.CreateConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.CreateConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.CreateConnectorInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCustomPlugin provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCustomPlugin(ctx context.Context, params *kafkaconnect.CreateCustomPluginInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateCustomPluginOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomPlugin")
	}

	var r0 *kafkaconnect.CreateCustomPluginOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateCustomPluginInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateCustomPluginOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateCustomPluginInput, ...func(*kafkaconnect.Options)) *kafkaconnect.CreateCustomPluginOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.CreateCustomPluginOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.CreateCustomPluginInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkerConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateWorkerConfiguration(ctx context.Context, params *kafkaconnect.CreateWorkerConfigurationInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateWorkerConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateWorkerConfiguration")
	}

	var r0 *kafkaconnect.CreateWorkerConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateWorkerConfigurationInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.CreateWorkerConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.CreateWorkerConfigurationInput, ...func(*kafkaconnect.Options)) *kafkaconnect.CreateWorkerConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.CreateWorkerConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.CreateWorkerConfigurationInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteConnector(ctx context.Context, params *kafkaconnect.DeleteConnectorInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnector")
	}

	var r0 *kafkaconnect.DeleteConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteConnectorInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteConnectorInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DeleteConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DeleteConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DeleteConnectorInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomPlugin provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCustomPlugin(ctx context.Context, params *kafkaconnect.DeleteCustomPluginInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteCustomPluginOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomPlugin")
	}

	var r0 *kafkaconnect.DeleteCustomPluginOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteCustomPluginInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteCustomPluginOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteCustomPluginInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DeleteCustomPluginOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DeleteCustomPluginOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DeleteCustomPluginInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkerConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteWorkerConfiguration(ctx context.Context, params *kafkaconnect.DeleteWorkerConfigurationInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteWorkerConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWorkerConfiguration")
	}

	var r0 *kafkaconnect.DeleteWorkerConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteWorkerConfigurationInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DeleteWorkerConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DeleteWorkerConfigurationInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DeleteWorkerConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DeleteWorkerConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DeleteWorkerConfigurationInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeConnector(ctx context.Context, params *kafkaconnect.DescribeConnectorInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeConnector")
	}

	var r0 *kafkaconnect.DescribeConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeConnectorInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeConnectorInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DescribeConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DescribeConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DescribeConnectorInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeConnectorOperation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeConnectorOperation(ctx context.Context, params *kafkaconnect.DescribeConnectorOperationInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeConnectorOperationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeConnectorOperation")
	}

	var r0 *kafkaconnect.DescribeConnectorOperationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeConnectorOperationInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeConnectorOperationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeConnectorOperationInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DescribeConnectorOperationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DescribeConnectorOperationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DescribeConnectorOperationInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCustomPlugin provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeCustomPlugin(ctx context.Context, params *kafkaconnect.DescribeCustomPluginInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeCustomPluginOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCustomPlugin")
	}

	var r0 *kafkaconnect.DescribeCustomPluginOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeCustomPluginInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeCustomPluginOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeCustomPluginInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DescribeCustomPluginOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DescribeCustomPluginOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DescribeCustomPluginInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeWorkerConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeWorkerConfiguration(ctx context.Context, params *kafkaconnect.DescribeWorkerConfigurationInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeWorkerConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeWorkerConfiguration")
	}

	var r0 *kafkaconnect.DescribeWorkerConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeWorkerConfigurationInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.DescribeWorkerConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.DescribeWorkerConfigurationInput, ...func(*kafkaconnect.Options)) *kafkaconnect.DescribeWorkerConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.DescribeWorkerConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.DescribeWorkerConfigurationInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConnectorOperations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListConnectorOperations(ctx context.Context, params *kafkaconnect.ListConnectorOperationsInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.ListConnectorOperationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListConnectorOperations")
	}

	var r0 *kafkaconnect.ListConnectorOperationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListConnectorOperationsInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.ListConnectorOperationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListConnectorOperationsInput, ...func(*kafkaconnect.Options)) *kafkaconnect.ListConnectorOperationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.ListConnectorOperationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.ListConnectorOperationsInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConnectors provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListConnectors(ctx context.Context, params *kafkaconnect.ListConnectorsInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.ListConnectorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListConnectors")
	}

	var r0 *kafkaconnect.ListConnectorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListConnectorsInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.ListConnectorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListConnectorsInput, ...func(*kafkaconnect.Options)) *kafkaconnect.ListConnectorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.ListConnectorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.ListConnectorsInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCustomPlugins provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCustomPlugins(ctx context.Context, params *kafkaconnect.ListCustomPluginsInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.ListCustomPluginsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCustomPlugins")
	}

	var r0 *kafkaconnect.ListCustomPluginsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListCustomPluginsInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.ListCustomPluginsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListCustomPluginsInput, ...func(*kafkaconnect.Options)) *kafkaconnect.ListCustomPluginsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.ListCustomPluginsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.ListCustomPluginsInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *kafkaconnect.ListTagsForResourceInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.ListTagsForResourceOutput, error) {
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

	var r0 *kafkaconnect.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListTagsForResourceInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListTagsForResourceInput, ...func(*kafkaconnect.Options)) *kafkaconnect.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.ListTagsForResourceInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkerConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListWorkerConfigurations(ctx context.Context, params *kafkaconnect.ListWorkerConfigurationsInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.ListWorkerConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListWorkerConfigurations")
	}

	var r0 *kafkaconnect.ListWorkerConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListWorkerConfigurationsInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.ListWorkerConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.ListWorkerConfigurationsInput, ...func(*kafkaconnect.Options)) *kafkaconnect.ListWorkerConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.ListWorkerConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.ListWorkerConfigurationsInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() kafkaconnect.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 kafkaconnect.Options
	if rf, ok := ret.Get(0).(func() kafkaconnect.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(kafkaconnect.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *kafkaconnect.TagResourceInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.TagResourceOutput, error) {
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

	var r0 *kafkaconnect.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.TagResourceInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.TagResourceInput, ...func(*kafkaconnect.Options)) *kafkaconnect.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.TagResourceInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *kafkaconnect.UntagResourceInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.UntagResourceOutput, error) {
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

	var r0 *kafkaconnect.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.UntagResourceInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.UntagResourceInput, ...func(*kafkaconnect.Options)) *kafkaconnect.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.UntagResourceInput, ...func(*kafkaconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateConnector(ctx context.Context, params *kafkaconnect.UpdateConnectorInput, optFns ...func(*kafkaconnect.Options)) (*kafkaconnect.UpdateConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnector")
	}

	var r0 *kafkaconnect.UpdateConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.UpdateConnectorInput, ...func(*kafkaconnect.Options)) (*kafkaconnect.UpdateConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kafkaconnect.UpdateConnectorInput, ...func(*kafkaconnect.Options)) *kafkaconnect.UpdateConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafkaconnect.UpdateConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kafkaconnect.UpdateConnectorInput, ...func(*kafkaconnect.Options)) error); ok {
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
