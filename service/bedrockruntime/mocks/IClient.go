// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bedrockruntime "github.com/aws/aws-sdk-go-v2/service/bedrockruntime"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// ApplyGuardrail provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ApplyGuardrail(ctx context.Context, params *bedrockruntime.ApplyGuardrailInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.ApplyGuardrailOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ApplyGuardrail")
	}

	var r0 *bedrockruntime.ApplyGuardrailOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ApplyGuardrailInput, ...func(*bedrockruntime.Options)) (*bedrockruntime.ApplyGuardrailOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ApplyGuardrailInput, ...func(*bedrockruntime.Options)) *bedrockruntime.ApplyGuardrailOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockruntime.ApplyGuardrailOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockruntime.ApplyGuardrailInput, ...func(*bedrockruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Converse provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Converse(ctx context.Context, params *bedrockruntime.ConverseInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.ConverseOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Converse")
	}

	var r0 *bedrockruntime.ConverseOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ConverseInput, ...func(*bedrockruntime.Options)) (*bedrockruntime.ConverseOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ConverseInput, ...func(*bedrockruntime.Options)) *bedrockruntime.ConverseOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockruntime.ConverseOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockruntime.ConverseInput, ...func(*bedrockruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConverseStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ConverseStream(ctx context.Context, params *bedrockruntime.ConverseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.ConverseStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ConverseStream")
	}

	var r0 *bedrockruntime.ConverseStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ConverseStreamInput, ...func(*bedrockruntime.Options)) (*bedrockruntime.ConverseStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.ConverseStreamInput, ...func(*bedrockruntime.Options)) *bedrockruntime.ConverseStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockruntime.ConverseStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockruntime.ConverseStreamInput, ...func(*bedrockruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) InvokeModel(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeModel")
	}

	var r0 *bedrockruntime.InvokeModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.InvokeModelInput, ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.InvokeModelInput, ...func(*bedrockruntime.Options)) *bedrockruntime.InvokeModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockruntime.InvokeModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockruntime.InvokeModelInput, ...func(*bedrockruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeModelWithResponseStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) InvokeModelWithResponseStream(ctx context.Context, params *bedrockruntime.InvokeModelWithResponseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeModelWithResponseStream")
	}

	var r0 *bedrockruntime.InvokeModelWithResponseStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.InvokeModelWithResponseStreamInput, ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockruntime.InvokeModelWithResponseStreamInput, ...func(*bedrockruntime.Options)) *bedrockruntime.InvokeModelWithResponseStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockruntime.InvokeModelWithResponseStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockruntime.InvokeModelWithResponseStreamInput, ...func(*bedrockruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() bedrockruntime.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bedrockruntime.Options
	if rf, ok := ret.Get(0).(func() bedrockruntime.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bedrockruntime.Options)
	}

	return r0
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
