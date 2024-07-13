// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mediapackagev2 "github.com/aws/aws-sdk-go-v2/service/mediapackagev2"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateChannel(ctx context.Context, params *mediapackagev2.CreateChannelInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateChannelOutput, error) {
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

	var r0 *mediapackagev2.CreateChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateChannelInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateChannelInput, ...func(*mediapackagev2.Options)) *mediapackagev2.CreateChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.CreateChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.CreateChannelInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateChannelGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateChannelGroup(ctx context.Context, params *mediapackagev2.CreateChannelGroupInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateChannelGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateChannelGroup")
	}

	var r0 *mediapackagev2.CreateChannelGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateChannelGroupInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateChannelGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateChannelGroupInput, ...func(*mediapackagev2.Options)) *mediapackagev2.CreateChannelGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.CreateChannelGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.CreateChannelGroupInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateOriginEndpoint(ctx context.Context, params *mediapackagev2.CreateOriginEndpointInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateOriginEndpointOutput, error) {
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

	var r0 *mediapackagev2.CreateOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateOriginEndpointInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.CreateOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.CreateOriginEndpointInput, ...func(*mediapackagev2.Options)) *mediapackagev2.CreateOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.CreateOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.CreateOriginEndpointInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteChannel(ctx context.Context, params *mediapackagev2.DeleteChannelInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelOutput, error) {
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

	var r0 *mediapackagev2.DeleteChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelInput, ...func(*mediapackagev2.Options)) *mediapackagev2.DeleteChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.DeleteChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.DeleteChannelInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChannelGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteChannelGroup(ctx context.Context, params *mediapackagev2.DeleteChannelGroupInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChannelGroup")
	}

	var r0 *mediapackagev2.DeleteChannelGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelGroupInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelGroupInput, ...func(*mediapackagev2.Options)) *mediapackagev2.DeleteChannelGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.DeleteChannelGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.DeleteChannelGroupInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChannelPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteChannelPolicy(ctx context.Context, params *mediapackagev2.DeleteChannelPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChannelPolicy")
	}

	var r0 *mediapackagev2.DeleteChannelPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteChannelPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteChannelPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.DeleteChannelPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.DeleteChannelPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.DeleteChannelPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteOriginEndpoint(ctx context.Context, params *mediapackagev2.DeleteOriginEndpointInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteOriginEndpointOutput, error) {
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

	var r0 *mediapackagev2.DeleteOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteOriginEndpointInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteOriginEndpointInput, ...func(*mediapackagev2.Options)) *mediapackagev2.DeleteOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.DeleteOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.DeleteOriginEndpointInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOriginEndpointPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteOriginEndpointPolicy(ctx context.Context, params *mediapackagev2.DeleteOriginEndpointPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteOriginEndpointPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOriginEndpointPolicy")
	}

	var r0 *mediapackagev2.DeleteOriginEndpointPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.DeleteOriginEndpointPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.DeleteOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.DeleteOriginEndpointPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.DeleteOriginEndpointPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.DeleteOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetChannel(ctx context.Context, params *mediapackagev2.GetChannelInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetChannel")
	}

	var r0 *mediapackagev2.GetChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelInput, ...func(*mediapackagev2.Options)) *mediapackagev2.GetChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.GetChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.GetChannelInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetChannelGroup(ctx context.Context, params *mediapackagev2.GetChannelGroupInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetChannelGroup")
	}

	var r0 *mediapackagev2.GetChannelGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelGroupInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelGroupInput, ...func(*mediapackagev2.Options)) *mediapackagev2.GetChannelGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.GetChannelGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.GetChannelGroupInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetChannelPolicy(ctx context.Context, params *mediapackagev2.GetChannelPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetChannelPolicy")
	}

	var r0 *mediapackagev2.GetChannelPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.GetChannelPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetChannelPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.GetChannelPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.GetChannelPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.GetChannelPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetOriginEndpoint(ctx context.Context, params *mediapackagev2.GetOriginEndpointInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.GetOriginEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOriginEndpoint")
	}

	var r0 *mediapackagev2.GetOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetOriginEndpointInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.GetOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetOriginEndpointInput, ...func(*mediapackagev2.Options)) *mediapackagev2.GetOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.GetOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.GetOriginEndpointInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOriginEndpointPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetOriginEndpointPolicy(ctx context.Context, params *mediapackagev2.GetOriginEndpointPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.GetOriginEndpointPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOriginEndpointPolicy")
	}

	var r0 *mediapackagev2.GetOriginEndpointPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.GetOriginEndpointPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.GetOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.GetOriginEndpointPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.GetOriginEndpointPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.GetOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListChannelGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListChannelGroups(ctx context.Context, params *mediapackagev2.ListChannelGroupsInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.ListChannelGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListChannelGroups")
	}

	var r0 *mediapackagev2.ListChannelGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListChannelGroupsInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.ListChannelGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListChannelGroupsInput, ...func(*mediapackagev2.Options)) *mediapackagev2.ListChannelGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.ListChannelGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.ListChannelGroupsInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListChannels provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListChannels(ctx context.Context, params *mediapackagev2.ListChannelsInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.ListChannelsOutput, error) {
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

	var r0 *mediapackagev2.ListChannelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListChannelsInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.ListChannelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListChannelsInput, ...func(*mediapackagev2.Options)) *mediapackagev2.ListChannelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.ListChannelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.ListChannelsInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOriginEndpoints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListOriginEndpoints(ctx context.Context, params *mediapackagev2.ListOriginEndpointsInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.ListOriginEndpointsOutput, error) {
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

	var r0 *mediapackagev2.ListOriginEndpointsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListOriginEndpointsInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.ListOriginEndpointsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListOriginEndpointsInput, ...func(*mediapackagev2.Options)) *mediapackagev2.ListOriginEndpointsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.ListOriginEndpointsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.ListOriginEndpointsInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *mediapackagev2.ListTagsForResourceInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.ListTagsForResourceOutput, error) {
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

	var r0 *mediapackagev2.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListTagsForResourceInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.ListTagsForResourceInput, ...func(*mediapackagev2.Options)) *mediapackagev2.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.ListTagsForResourceInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() mediapackagev2.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 mediapackagev2.Options
	if rf, ok := ret.Get(0).(func() mediapackagev2.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mediapackagev2.Options)
	}

	return r0
}

// PutChannelPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutChannelPolicy(ctx context.Context, params *mediapackagev2.PutChannelPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.PutChannelPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutChannelPolicy")
	}

	var r0 *mediapackagev2.PutChannelPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.PutChannelPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.PutChannelPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.PutChannelPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.PutChannelPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.PutChannelPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.PutChannelPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutOriginEndpointPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutOriginEndpointPolicy(ctx context.Context, params *mediapackagev2.PutOriginEndpointPolicyInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.PutOriginEndpointPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutOriginEndpointPolicy")
	}

	var r0 *mediapackagev2.PutOriginEndpointPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.PutOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.PutOriginEndpointPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.PutOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) *mediapackagev2.PutOriginEndpointPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.PutOriginEndpointPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.PutOriginEndpointPolicyInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *mediapackagev2.TagResourceInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.TagResourceOutput, error) {
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

	var r0 *mediapackagev2.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.TagResourceInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.TagResourceInput, ...func(*mediapackagev2.Options)) *mediapackagev2.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.TagResourceInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *mediapackagev2.UntagResourceInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.UntagResourceOutput, error) {
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

	var r0 *mediapackagev2.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UntagResourceInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UntagResourceInput, ...func(*mediapackagev2.Options)) *mediapackagev2.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.UntagResourceInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChannel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateChannel(ctx context.Context, params *mediapackagev2.UpdateChannelInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateChannelOutput, error) {
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

	var r0 *mediapackagev2.UpdateChannelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateChannelInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateChannelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateChannelInput, ...func(*mediapackagev2.Options)) *mediapackagev2.UpdateChannelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.UpdateChannelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.UpdateChannelInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChannelGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateChannelGroup(ctx context.Context, params *mediapackagev2.UpdateChannelGroupInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateChannelGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChannelGroup")
	}

	var r0 *mediapackagev2.UpdateChannelGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateChannelGroupInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateChannelGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateChannelGroupInput, ...func(*mediapackagev2.Options)) *mediapackagev2.UpdateChannelGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.UpdateChannelGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.UpdateChannelGroupInput, ...func(*mediapackagev2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOriginEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateOriginEndpoint(ctx context.Context, params *mediapackagev2.UpdateOriginEndpointInput, optFns ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateOriginEndpointOutput, error) {
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

	var r0 *mediapackagev2.UpdateOriginEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateOriginEndpointInput, ...func(*mediapackagev2.Options)) (*mediapackagev2.UpdateOriginEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediapackagev2.UpdateOriginEndpointInput, ...func(*mediapackagev2.Options)) *mediapackagev2.UpdateOriginEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediapackagev2.UpdateOriginEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediapackagev2.UpdateOriginEndpointInput, ...func(*mediapackagev2.Options)) error); ok {
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
