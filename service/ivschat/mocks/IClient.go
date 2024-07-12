// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ivschat "github.com/aws/aws-sdk-go-v2/service/ivschat"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateChatToken provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateChatToken(ctx context.Context, params *ivschat.CreateChatTokenInput, optFns ...func(*ivschat.Options)) (*ivschat.CreateChatTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateChatToken")
	}

	var r0 *ivschat.CreateChatTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateChatTokenInput, ...func(*ivschat.Options)) (*ivschat.CreateChatTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateChatTokenInput, ...func(*ivschat.Options)) *ivschat.CreateChatTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.CreateChatTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.CreateChatTokenInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateLoggingConfiguration(ctx context.Context, params *ivschat.CreateLoggingConfigurationInput, optFns ...func(*ivschat.Options)) (*ivschat.CreateLoggingConfigurationOutput, error) {
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

	var r0 *ivschat.CreateLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateLoggingConfigurationInput, ...func(*ivschat.Options)) (*ivschat.CreateLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateLoggingConfigurationInput, ...func(*ivschat.Options)) *ivschat.CreateLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.CreateLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.CreateLoggingConfigurationInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoom provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRoom(ctx context.Context, params *ivschat.CreateRoomInput, optFns ...func(*ivschat.Options)) (*ivschat.CreateRoomOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRoom")
	}

	var r0 *ivschat.CreateRoomOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateRoomInput, ...func(*ivschat.Options)) (*ivschat.CreateRoomOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.CreateRoomInput, ...func(*ivschat.Options)) *ivschat.CreateRoomOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.CreateRoomOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.CreateRoomInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteLoggingConfiguration(ctx context.Context, params *ivschat.DeleteLoggingConfigurationInput, optFns ...func(*ivschat.Options)) (*ivschat.DeleteLoggingConfigurationOutput, error) {
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

	var r0 *ivschat.DeleteLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteLoggingConfigurationInput, ...func(*ivschat.Options)) (*ivschat.DeleteLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteLoggingConfigurationInput, ...func(*ivschat.Options)) *ivschat.DeleteLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.DeleteLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.DeleteLoggingConfigurationInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMessage(ctx context.Context, params *ivschat.DeleteMessageInput, optFns ...func(*ivschat.Options)) (*ivschat.DeleteMessageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMessage")
	}

	var r0 *ivschat.DeleteMessageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteMessageInput, ...func(*ivschat.Options)) (*ivschat.DeleteMessageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteMessageInput, ...func(*ivschat.Options)) *ivschat.DeleteMessageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.DeleteMessageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.DeleteMessageInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoom provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRoom(ctx context.Context, params *ivschat.DeleteRoomInput, optFns ...func(*ivschat.Options)) (*ivschat.DeleteRoomOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRoom")
	}

	var r0 *ivschat.DeleteRoomOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteRoomInput, ...func(*ivschat.Options)) (*ivschat.DeleteRoomOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DeleteRoomInput, ...func(*ivschat.Options)) *ivschat.DeleteRoomOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.DeleteRoomOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.DeleteRoomInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisconnectUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisconnectUser(ctx context.Context, params *ivschat.DisconnectUserInput, optFns ...func(*ivschat.Options)) (*ivschat.DisconnectUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisconnectUser")
	}

	var r0 *ivschat.DisconnectUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DisconnectUserInput, ...func(*ivschat.Options)) (*ivschat.DisconnectUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.DisconnectUserInput, ...func(*ivschat.Options)) *ivschat.DisconnectUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.DisconnectUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.DisconnectUserInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetLoggingConfiguration(ctx context.Context, params *ivschat.GetLoggingConfigurationInput, optFns ...func(*ivschat.Options)) (*ivschat.GetLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLoggingConfiguration")
	}

	var r0 *ivschat.GetLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.GetLoggingConfigurationInput, ...func(*ivschat.Options)) (*ivschat.GetLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.GetLoggingConfigurationInput, ...func(*ivschat.Options)) *ivschat.GetLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.GetLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.GetLoggingConfigurationInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoom provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRoom(ctx context.Context, params *ivschat.GetRoomInput, optFns ...func(*ivschat.Options)) (*ivschat.GetRoomOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRoom")
	}

	var r0 *ivschat.GetRoomOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.GetRoomInput, ...func(*ivschat.Options)) (*ivschat.GetRoomOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.GetRoomInput, ...func(*ivschat.Options)) *ivschat.GetRoomOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.GetRoomOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.GetRoomInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLoggingConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListLoggingConfigurations(ctx context.Context, params *ivschat.ListLoggingConfigurationsInput, optFns ...func(*ivschat.Options)) (*ivschat.ListLoggingConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLoggingConfigurations")
	}

	var r0 *ivschat.ListLoggingConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListLoggingConfigurationsInput, ...func(*ivschat.Options)) (*ivschat.ListLoggingConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListLoggingConfigurationsInput, ...func(*ivschat.Options)) *ivschat.ListLoggingConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.ListLoggingConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.ListLoggingConfigurationsInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRooms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRooms(ctx context.Context, params *ivschat.ListRoomsInput, optFns ...func(*ivschat.Options)) (*ivschat.ListRoomsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRooms")
	}

	var r0 *ivschat.ListRoomsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListRoomsInput, ...func(*ivschat.Options)) (*ivschat.ListRoomsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListRoomsInput, ...func(*ivschat.Options)) *ivschat.ListRoomsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.ListRoomsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.ListRoomsInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *ivschat.ListTagsForResourceInput, optFns ...func(*ivschat.Options)) (*ivschat.ListTagsForResourceOutput, error) {
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

	var r0 *ivschat.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListTagsForResourceInput, ...func(*ivschat.Options)) (*ivschat.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.ListTagsForResourceInput, ...func(*ivschat.Options)) *ivschat.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.ListTagsForResourceInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() ivschat.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 ivschat.Options
	if rf, ok := ret.Get(0).(func() ivschat.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ivschat.Options)
	}

	return r0
}

// SendEvent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SendEvent(ctx context.Context, params *ivschat.SendEventInput, optFns ...func(*ivschat.Options)) (*ivschat.SendEventOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendEvent")
	}

	var r0 *ivschat.SendEventOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.SendEventInput, ...func(*ivschat.Options)) (*ivschat.SendEventOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.SendEventInput, ...func(*ivschat.Options)) *ivschat.SendEventOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.SendEventOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.SendEventInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *ivschat.TagResourceInput, optFns ...func(*ivschat.Options)) (*ivschat.TagResourceOutput, error) {
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

	var r0 *ivschat.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.TagResourceInput, ...func(*ivschat.Options)) (*ivschat.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.TagResourceInput, ...func(*ivschat.Options)) *ivschat.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.TagResourceInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *ivschat.UntagResourceInput, optFns ...func(*ivschat.Options)) (*ivschat.UntagResourceOutput, error) {
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

	var r0 *ivschat.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UntagResourceInput, ...func(*ivschat.Options)) (*ivschat.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UntagResourceInput, ...func(*ivschat.Options)) *ivschat.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.UntagResourceInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateLoggingConfiguration(ctx context.Context, params *ivschat.UpdateLoggingConfigurationInput, optFns ...func(*ivschat.Options)) (*ivschat.UpdateLoggingConfigurationOutput, error) {
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

	var r0 *ivschat.UpdateLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UpdateLoggingConfigurationInput, ...func(*ivschat.Options)) (*ivschat.UpdateLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UpdateLoggingConfigurationInput, ...func(*ivschat.Options)) *ivschat.UpdateLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.UpdateLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.UpdateLoggingConfigurationInput, ...func(*ivschat.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRoom provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRoom(ctx context.Context, params *ivschat.UpdateRoomInput, optFns ...func(*ivschat.Options)) (*ivschat.UpdateRoomOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRoom")
	}

	var r0 *ivschat.UpdateRoomOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UpdateRoomInput, ...func(*ivschat.Options)) (*ivschat.UpdateRoomOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ivschat.UpdateRoomInput, ...func(*ivschat.Options)) *ivschat.UpdateRoomOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ivschat.UpdateRoomOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ivschat.UpdateRoomInput, ...func(*ivschat.Options)) error); ok {
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
