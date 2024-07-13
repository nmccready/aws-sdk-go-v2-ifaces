// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	connectparticipant "github.com/aws/aws-sdk-go-v2/service/connectparticipant"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CompleteAttachmentUpload provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CompleteAttachmentUpload(ctx context.Context, params *connectparticipant.CompleteAttachmentUploadInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.CompleteAttachmentUploadOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CompleteAttachmentUpload")
	}

	var r0 *connectparticipant.CompleteAttachmentUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.CompleteAttachmentUploadInput, ...func(*connectparticipant.Options)) (*connectparticipant.CompleteAttachmentUploadOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.CompleteAttachmentUploadInput, ...func(*connectparticipant.Options)) *connectparticipant.CompleteAttachmentUploadOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.CompleteAttachmentUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.CompleteAttachmentUploadInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateParticipantConnection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateParticipantConnection(ctx context.Context, params *connectparticipant.CreateParticipantConnectionInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.CreateParticipantConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateParticipantConnection")
	}

	var r0 *connectparticipant.CreateParticipantConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.CreateParticipantConnectionInput, ...func(*connectparticipant.Options)) (*connectparticipant.CreateParticipantConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.CreateParticipantConnectionInput, ...func(*connectparticipant.Options)) *connectparticipant.CreateParticipantConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.CreateParticipantConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.CreateParticipantConnectionInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeView provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeView(ctx context.Context, params *connectparticipant.DescribeViewInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.DescribeViewOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeView")
	}

	var r0 *connectparticipant.DescribeViewOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.DescribeViewInput, ...func(*connectparticipant.Options)) (*connectparticipant.DescribeViewOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.DescribeViewInput, ...func(*connectparticipant.Options)) *connectparticipant.DescribeViewOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.DescribeViewOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.DescribeViewInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisconnectParticipant provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisconnectParticipant(ctx context.Context, params *connectparticipant.DisconnectParticipantInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.DisconnectParticipantOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisconnectParticipant")
	}

	var r0 *connectparticipant.DisconnectParticipantOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.DisconnectParticipantInput, ...func(*connectparticipant.Options)) (*connectparticipant.DisconnectParticipantOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.DisconnectParticipantInput, ...func(*connectparticipant.Options)) *connectparticipant.DisconnectParticipantOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.DisconnectParticipantOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.DisconnectParticipantInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttachment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAttachment(ctx context.Context, params *connectparticipant.GetAttachmentInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.GetAttachmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAttachment")
	}

	var r0 *connectparticipant.GetAttachmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.GetAttachmentInput, ...func(*connectparticipant.Options)) (*connectparticipant.GetAttachmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.GetAttachmentInput, ...func(*connectparticipant.Options)) *connectparticipant.GetAttachmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.GetAttachmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.GetAttachmentInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTranscript provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTranscript(ctx context.Context, params *connectparticipant.GetTranscriptInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.GetTranscriptOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTranscript")
	}

	var r0 *connectparticipant.GetTranscriptOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.GetTranscriptInput, ...func(*connectparticipant.Options)) (*connectparticipant.GetTranscriptOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.GetTranscriptInput, ...func(*connectparticipant.Options)) *connectparticipant.GetTranscriptOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.GetTranscriptOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.GetTranscriptInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() connectparticipant.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 connectparticipant.Options
	if rf, ok := ret.Get(0).(func() connectparticipant.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(connectparticipant.Options)
	}

	return r0
}

// SendEvent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SendEvent(ctx context.Context, params *connectparticipant.SendEventInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.SendEventOutput, error) {
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

	var r0 *connectparticipant.SendEventOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.SendEventInput, ...func(*connectparticipant.Options)) (*connectparticipant.SendEventOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.SendEventInput, ...func(*connectparticipant.Options)) *connectparticipant.SendEventOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.SendEventOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.SendEventInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SendMessage(ctx context.Context, params *connectparticipant.SendMessageInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.SendMessageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 *connectparticipant.SendMessageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.SendMessageInput, ...func(*connectparticipant.Options)) (*connectparticipant.SendMessageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.SendMessageInput, ...func(*connectparticipant.Options)) *connectparticipant.SendMessageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.SendMessageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.SendMessageInput, ...func(*connectparticipant.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartAttachmentUpload provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartAttachmentUpload(ctx context.Context, params *connectparticipant.StartAttachmentUploadInput, optFns ...func(*connectparticipant.Options)) (*connectparticipant.StartAttachmentUploadOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartAttachmentUpload")
	}

	var r0 *connectparticipant.StartAttachmentUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.StartAttachmentUploadInput, ...func(*connectparticipant.Options)) (*connectparticipant.StartAttachmentUploadOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectparticipant.StartAttachmentUploadInput, ...func(*connectparticipant.Options)) *connectparticipant.StartAttachmentUploadOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectparticipant.StartAttachmentUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectparticipant.StartAttachmentUploadInput, ...func(*connectparticipant.Options)) error); ok {
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
