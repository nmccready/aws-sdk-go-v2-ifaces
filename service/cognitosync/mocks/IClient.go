// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	cognitosync "github.com/aws/aws-sdk-go-v2/service/cognitosync"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BulkPublish provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BulkPublish(ctx context.Context, params *cognitosync.BulkPublishInput, optFns ...func(*cognitosync.Options)) (*cognitosync.BulkPublishOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BulkPublish")
	}

	var r0 *cognitosync.BulkPublishOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.BulkPublishInput, ...func(*cognitosync.Options)) (*cognitosync.BulkPublishOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.BulkPublishInput, ...func(*cognitosync.Options)) *cognitosync.BulkPublishOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.BulkPublishOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.BulkPublishInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDataset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDataset(ctx context.Context, params *cognitosync.DeleteDatasetInput, optFns ...func(*cognitosync.Options)) (*cognitosync.DeleteDatasetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDataset")
	}

	var r0 *cognitosync.DeleteDatasetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DeleteDatasetInput, ...func(*cognitosync.Options)) (*cognitosync.DeleteDatasetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DeleteDatasetInput, ...func(*cognitosync.Options)) *cognitosync.DeleteDatasetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.DeleteDatasetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.DeleteDatasetInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDataset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDataset(ctx context.Context, params *cognitosync.DescribeDatasetInput, optFns ...func(*cognitosync.Options)) (*cognitosync.DescribeDatasetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDataset")
	}

	var r0 *cognitosync.DescribeDatasetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeDatasetInput, ...func(*cognitosync.Options)) (*cognitosync.DescribeDatasetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeDatasetInput, ...func(*cognitosync.Options)) *cognitosync.DescribeDatasetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.DescribeDatasetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.DescribeDatasetInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentityPoolUsage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeIdentityPoolUsage(ctx context.Context, params *cognitosync.DescribeIdentityPoolUsageInput, optFns ...func(*cognitosync.Options)) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeIdentityPoolUsage")
	}

	var r0 *cognitosync.DescribeIdentityPoolUsageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeIdentityPoolUsageInput, ...func(*cognitosync.Options)) (*cognitosync.DescribeIdentityPoolUsageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeIdentityPoolUsageInput, ...func(*cognitosync.Options)) *cognitosync.DescribeIdentityPoolUsageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.DescribeIdentityPoolUsageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.DescribeIdentityPoolUsageInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentityUsage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeIdentityUsage(ctx context.Context, params *cognitosync.DescribeIdentityUsageInput, optFns ...func(*cognitosync.Options)) (*cognitosync.DescribeIdentityUsageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeIdentityUsage")
	}

	var r0 *cognitosync.DescribeIdentityUsageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeIdentityUsageInput, ...func(*cognitosync.Options)) (*cognitosync.DescribeIdentityUsageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.DescribeIdentityUsageInput, ...func(*cognitosync.Options)) *cognitosync.DescribeIdentityUsageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.DescribeIdentityUsageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.DescribeIdentityUsageInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBulkPublishDetails provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetBulkPublishDetails(ctx context.Context, params *cognitosync.GetBulkPublishDetailsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBulkPublishDetails")
	}

	var r0 *cognitosync.GetBulkPublishDetailsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetBulkPublishDetailsInput, ...func(*cognitosync.Options)) (*cognitosync.GetBulkPublishDetailsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetBulkPublishDetailsInput, ...func(*cognitosync.Options)) *cognitosync.GetBulkPublishDetailsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.GetBulkPublishDetailsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.GetBulkPublishDetailsInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCognitoEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCognitoEvents(ctx context.Context, params *cognitosync.GetCognitoEventsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.GetCognitoEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCognitoEvents")
	}

	var r0 *cognitosync.GetCognitoEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetCognitoEventsInput, ...func(*cognitosync.Options)) (*cognitosync.GetCognitoEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetCognitoEventsInput, ...func(*cognitosync.Options)) *cognitosync.GetCognitoEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.GetCognitoEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.GetCognitoEventsInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityPoolConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetIdentityPoolConfiguration(ctx context.Context, params *cognitosync.GetIdentityPoolConfigurationInput, optFns ...func(*cognitosync.Options)) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIdentityPoolConfiguration")
	}

	var r0 *cognitosync.GetIdentityPoolConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) (*cognitosync.GetIdentityPoolConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.GetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) *cognitosync.GetIdentityPoolConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.GetIdentityPoolConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.GetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDatasets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDatasets(ctx context.Context, params *cognitosync.ListDatasetsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.ListDatasetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDatasets")
	}

	var r0 *cognitosync.ListDatasetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListDatasetsInput, ...func(*cognitosync.Options)) (*cognitosync.ListDatasetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListDatasetsInput, ...func(*cognitosync.Options)) *cognitosync.ListDatasetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.ListDatasetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.ListDatasetsInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentityPoolUsage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListIdentityPoolUsage(ctx context.Context, params *cognitosync.ListIdentityPoolUsageInput, optFns ...func(*cognitosync.Options)) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdentityPoolUsage")
	}

	var r0 *cognitosync.ListIdentityPoolUsageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListIdentityPoolUsageInput, ...func(*cognitosync.Options)) (*cognitosync.ListIdentityPoolUsageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListIdentityPoolUsageInput, ...func(*cognitosync.Options)) *cognitosync.ListIdentityPoolUsageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.ListIdentityPoolUsageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.ListIdentityPoolUsageInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecords provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRecords(ctx context.Context, params *cognitosync.ListRecordsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.ListRecordsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRecords")
	}

	var r0 *cognitosync.ListRecordsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListRecordsInput, ...func(*cognitosync.Options)) (*cognitosync.ListRecordsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.ListRecordsInput, ...func(*cognitosync.Options)) *cognitosync.ListRecordsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.ListRecordsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.ListRecordsInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() cognitosync.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 cognitosync.Options
	if rf, ok := ret.Get(0).(func() cognitosync.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cognitosync.Options)
	}

	return r0
}

// RegisterDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RegisterDevice(ctx context.Context, params *cognitosync.RegisterDeviceInput, optFns ...func(*cognitosync.Options)) (*cognitosync.RegisterDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterDevice")
	}

	var r0 *cognitosync.RegisterDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.RegisterDeviceInput, ...func(*cognitosync.Options)) (*cognitosync.RegisterDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.RegisterDeviceInput, ...func(*cognitosync.Options)) *cognitosync.RegisterDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.RegisterDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.RegisterDeviceInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCognitoEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SetCognitoEvents(ctx context.Context, params *cognitosync.SetCognitoEventsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.SetCognitoEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetCognitoEvents")
	}

	var r0 *cognitosync.SetCognitoEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SetCognitoEventsInput, ...func(*cognitosync.Options)) (*cognitosync.SetCognitoEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SetCognitoEventsInput, ...func(*cognitosync.Options)) *cognitosync.SetCognitoEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.SetCognitoEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.SetCognitoEventsInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetIdentityPoolConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SetIdentityPoolConfiguration(ctx context.Context, params *cognitosync.SetIdentityPoolConfigurationInput, optFns ...func(*cognitosync.Options)) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetIdentityPoolConfiguration")
	}

	var r0 *cognitosync.SetIdentityPoolConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) (*cognitosync.SetIdentityPoolConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) *cognitosync.SetIdentityPoolConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.SetIdentityPoolConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.SetIdentityPoolConfigurationInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToDataset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SubscribeToDataset(ctx context.Context, params *cognitosync.SubscribeToDatasetInput, optFns ...func(*cognitosync.Options)) (*cognitosync.SubscribeToDatasetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeToDataset")
	}

	var r0 *cognitosync.SubscribeToDatasetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SubscribeToDatasetInput, ...func(*cognitosync.Options)) (*cognitosync.SubscribeToDatasetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.SubscribeToDatasetInput, ...func(*cognitosync.Options)) *cognitosync.SubscribeToDatasetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.SubscribeToDatasetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.SubscribeToDatasetInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnsubscribeFromDataset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UnsubscribeFromDataset(ctx context.Context, params *cognitosync.UnsubscribeFromDatasetInput, optFns ...func(*cognitosync.Options)) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnsubscribeFromDataset")
	}

	var r0 *cognitosync.UnsubscribeFromDatasetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.UnsubscribeFromDatasetInput, ...func(*cognitosync.Options)) (*cognitosync.UnsubscribeFromDatasetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.UnsubscribeFromDatasetInput, ...func(*cognitosync.Options)) *cognitosync.UnsubscribeFromDatasetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.UnsubscribeFromDatasetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.UnsubscribeFromDatasetInput, ...func(*cognitosync.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRecords provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRecords(ctx context.Context, params *cognitosync.UpdateRecordsInput, optFns ...func(*cognitosync.Options)) (*cognitosync.UpdateRecordsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRecords")
	}

	var r0 *cognitosync.UpdateRecordsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.UpdateRecordsInput, ...func(*cognitosync.Options)) (*cognitosync.UpdateRecordsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitosync.UpdateRecordsInput, ...func(*cognitosync.Options)) *cognitosync.UpdateRecordsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitosync.UpdateRecordsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitosync.UpdateRecordsInput, ...func(*cognitosync.Options)) error); ok {
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
