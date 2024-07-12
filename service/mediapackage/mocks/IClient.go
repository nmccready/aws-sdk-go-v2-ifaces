// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mediapackage "github.com/aws/aws-sdk-go-v2/service/mediapackage"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// ConfigureLogs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ConfigureLogs(ctx context.Context, params *mediapackage.ConfigureLogsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.ConfigureLogsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ConfigureLogs")
	}

	var r0 *mediapackage.ConfigureLogsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ConfigureLogsInput, ...func(*mediapackage.Options)) (*mediapackage.ConfigureLogsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ConfigureLogsInput, ...func(*mediapackage.Options)) *mediapackage.ConfigureLogsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.ConfigureLogsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.ConfigureLogsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateChannel(ctx context.Context, params *mediapackage.CreateChannelInput, optFns ...func(*mediapackage.Options)) (*mediapackage.CreateChannelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateChannel")
	}

	var r0 *mediapackage.CreateChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateChannelInput, ...func(*mediapackage.Options)) (*mediapackage.CreateChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateChannelInput, ...func(*mediapackage.Options)) *mediapackage.CreateChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.CreateChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.CreateChannelInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHarvestJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateHarvestJob(ctx context.Context, params *mediapackage.CreateHarvestJobInput, optFns ...func(*mediapackage.Options)) (*mediapackage.CreateHarvestJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateHarvestJob")
	}

	var r0 *mediapackage.CreateHarvestJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateHarvestJobInput, ...func(*mediapackage.Options)) (*mediapackage.CreateHarvestJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateHarvestJobInput, ...func(*mediapackage.Options)) *mediapackage.CreateHarvestJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.CreateHarvestJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.CreateHarvestJobInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateOriginEndpoint(ctx context.Context, params *mediapackage.CreateOriginEndpointInput, optFns ...func(*mediapackage.Options)) (*mediapackage.CreateOriginEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOriginEndpoint")
	}

	var r0 *mediapackage.CreateOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateOriginEndpointInput, ...func(*mediapackage.Options)) (*mediapackage.CreateOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.CreateOriginEndpointInput, ...func(*mediapackage.Options)) *mediapackage.CreateOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.CreateOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.CreateOriginEndpointInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteChannel(ctx context.Context, params *mediapackage.DeleteChannelInput, optFns ...func(*mediapackage.Options)) (*mediapackage.DeleteChannelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChannel")
	}

	var r0 *mediapackage.DeleteChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DeleteChannelInput, ...func(*mediapackage.Options)) (*mediapackage.DeleteChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DeleteChannelInput, ...func(*mediapackage.Options)) *mediapackage.DeleteChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.DeleteChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.DeleteChannelInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteOriginEndpoint(ctx context.Context, params *mediapackage.DeleteOriginEndpointInput, optFns ...func(*mediapackage.Options)) (*mediapackage.DeleteOriginEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOriginEndpoint")
	}

	var r0 *mediapackage.DeleteOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DeleteOriginEndpointInput, ...func(*mediapackage.Options)) (*mediapackage.DeleteOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DeleteOriginEndpointInput, ...func(*mediapackage.Options)) *mediapackage.DeleteOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.DeleteOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.DeleteOriginEndpointInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeChannel(ctx context.Context, params *mediapackage.DescribeChannelInput, optFns ...func(*mediapackage.Options)) (*mediapackage.DescribeChannelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeChannel")
	}

	var r0 *mediapackage.DescribeChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeChannelInput, ...func(*mediapackage.Options)) (*mediapackage.DescribeChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeChannelInput, ...func(*mediapackage.Options)) *mediapackage.DescribeChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.DescribeChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.DescribeChannelInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHarvestJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeHarvestJob(ctx context.Context, params *mediapackage.DescribeHarvestJobInput, optFns ...func(*mediapackage.Options)) (*mediapackage.DescribeHarvestJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeHarvestJob")
	}

	var r0 *mediapackage.DescribeHarvestJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeHarvestJobInput, ...func(*mediapackage.Options)) (*mediapackage.DescribeHarvestJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeHarvestJobInput, ...func(*mediapackage.Options)) *mediapackage.DescribeHarvestJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.DescribeHarvestJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.DescribeHarvestJobInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeOriginEndpoint(ctx context.Context, params *mediapackage.DescribeOriginEndpointInput, optFns ...func(*mediapackage.Options)) (*mediapackage.DescribeOriginEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeOriginEndpoint")
	}

	var r0 *mediapackage.DescribeOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeOriginEndpointInput, ...func(*mediapackage.Options)) (*mediapackage.DescribeOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.DescribeOriginEndpointInput, ...func(*mediapackage.Options)) *mediapackage.DescribeOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.DescribeOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.DescribeOriginEndpointInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListChannels provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListChannels(ctx context.Context, params *mediapackage.ListChannelsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.ListChannelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListChannels")
	}

	var r0 *mediapackage.ListChannelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListChannelsInput, ...func(*mediapackage.Options)) (*mediapackage.ListChannelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListChannelsInput, ...func(*mediapackage.Options)) *mediapackage.ListChannelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.ListChannelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.ListChannelsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHarvestJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListHarvestJobs(ctx context.Context, params *mediapackage.ListHarvestJobsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.ListHarvestJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListHarvestJobs")
	}

	var r0 *mediapackage.ListHarvestJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListHarvestJobsInput, ...func(*mediapackage.Options)) (*mediapackage.ListHarvestJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListHarvestJobsInput, ...func(*mediapackage.Options)) *mediapackage.ListHarvestJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.ListHarvestJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.ListHarvestJobsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOriginEndpoints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListOriginEndpoints(ctx context.Context, params *mediapackage.ListOriginEndpointsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.ListOriginEndpointsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListOriginEndpoints")
	}

	var r0 *mediapackage.ListOriginEndpointsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListOriginEndpointsInput, ...func(*mediapackage.Options)) (*mediapackage.ListOriginEndpointsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListOriginEndpointsInput, ...func(*mediapackage.Options)) *mediapackage.ListOriginEndpointsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.ListOriginEndpointsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.ListOriginEndpointsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *mediapackage.ListTagsForResourceInput, optFns ...func(*mediapackage.Options)) (*mediapackage.ListTagsForResourceOutput, error) {
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

	var r0 *mediapackage.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListTagsForResourceInput, ...func(*mediapackage.Options)) (*mediapackage.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.ListTagsForResourceInput, ...func(*mediapackage.Options)) *mediapackage.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.ListTagsForResourceInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() mediapackage.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 mediapackage.Options
	if rf, ok := ret.Get(0).(func() mediapackage.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mediapackage.Options)
	}

	return r0
}

// RotateChannelCredentials provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RotateChannelCredentials(ctx context.Context, params *mediapackage.RotateChannelCredentialsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.RotateChannelCredentialsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RotateChannelCredentials")
	}

	var r0 *mediapackage.RotateChannelCredentialsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.RotateChannelCredentialsInput, ...func(*mediapackage.Options)) (*mediapackage.RotateChannelCredentialsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.RotateChannelCredentialsInput, ...func(*mediapackage.Options)) *mediapackage.RotateChannelCredentialsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.RotateChannelCredentialsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.RotateChannelCredentialsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RotateIngestEndpointCredentials provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RotateIngestEndpointCredentials(ctx context.Context, params *mediapackage.RotateIngestEndpointCredentialsInput, optFns ...func(*mediapackage.Options)) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RotateIngestEndpointCredentials")
	}

	var r0 *mediapackage.RotateIngestEndpointCredentialsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.RotateIngestEndpointCredentialsInput, ...func(*mediapackage.Options)) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.RotateIngestEndpointCredentialsInput, ...func(*mediapackage.Options)) *mediapackage.RotateIngestEndpointCredentialsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.RotateIngestEndpointCredentialsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.RotateIngestEndpointCredentialsInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *mediapackage.TagResourceInput, optFns ...func(*mediapackage.Options)) (*mediapackage.TagResourceOutput, error) {
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

	var r0 *mediapackage.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.TagResourceInput, ...func(*mediapackage.Options)) (*mediapackage.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.TagResourceInput, ...func(*mediapackage.Options)) *mediapackage.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.TagResourceInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *mediapackage.UntagResourceInput, optFns ...func(*mediapackage.Options)) (*mediapackage.UntagResourceOutput, error) {
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

	var r0 *mediapackage.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UntagResourceInput, ...func(*mediapackage.Options)) (*mediapackage.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UntagResourceInput, ...func(*mediapackage.Options)) *mediapackage.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.UntagResourceInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateChannel(ctx context.Context, params *mediapackage.UpdateChannelInput, optFns ...func(*mediapackage.Options)) (*mediapackage.UpdateChannelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChannel")
	}

	var r0 *mediapackage.UpdateChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UpdateChannelInput, ...func(*mediapackage.Options)) (*mediapackage.UpdateChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UpdateChannelInput, ...func(*mediapackage.Options)) *mediapackage.UpdateChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.UpdateChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.UpdateChannelInput, ...func(*mediapackage.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateOriginEndpoint(ctx context.Context, params *mediapackage.UpdateOriginEndpointInput, optFns ...func(*mediapackage.Options)) (*mediapackage.UpdateOriginEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOriginEndpoint")
	}

	var r0 *mediapackage.UpdateOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UpdateOriginEndpointInput, ...func(*mediapackage.Options)) (*mediapackage.UpdateOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackage.UpdateOriginEndpointInput, ...func(*mediapackage.Options)) *mediapackage.UpdateOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackage.UpdateOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackage.UpdateOriginEndpointInput, ...func(*mediapackage.Options)) error); ok {
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
