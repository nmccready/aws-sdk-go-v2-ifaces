// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bedrockagentruntime "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DeleteAgentMemory provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAgentMemory(ctx context.Context, params *bedrockagentruntime.DeleteAgentMemoryInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.DeleteAgentMemoryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAgentMemory")
	}

	var r0 *bedrockagentruntime.DeleteAgentMemoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.DeleteAgentMemoryInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.DeleteAgentMemoryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.DeleteAgentMemoryInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.DeleteAgentMemoryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.DeleteAgentMemoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.DeleteAgentMemoryInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateQuery provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateQuery(ctx context.Context, params *bedrockagentruntime.GenerateQueryInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.GenerateQueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateQuery")
	}

	var r0 *bedrockagentruntime.GenerateQueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.GenerateQueryInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.GenerateQueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.GenerateQueryInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.GenerateQueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.GenerateQueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.GenerateQueryInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentMemory provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAgentMemory(ctx context.Context, params *bedrockagentruntime.GetAgentMemoryInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.GetAgentMemoryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAgentMemory")
	}

	var r0 *bedrockagentruntime.GetAgentMemoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.GetAgentMemoryInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.GetAgentMemoryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.GetAgentMemoryInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.GetAgentMemoryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.GetAgentMemoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.GetAgentMemoryInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeAgent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) InvokeAgent(ctx context.Context, params *bedrockagentruntime.InvokeAgentInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeAgentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeAgent")
	}

	var r0 *bedrockagentruntime.InvokeAgentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeAgentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.InvokeAgentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.InvokeAgentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeFlow provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) InvokeFlow(ctx context.Context, params *bedrockagentruntime.InvokeFlowInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeFlowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeFlow")
	}

	var r0 *bedrockagentruntime.InvokeFlowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeFlowInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeFlowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeFlowInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.InvokeFlowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.InvokeFlowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.InvokeFlowInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeInlineAgent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) InvokeInlineAgent(ctx context.Context, params *bedrockagentruntime.InvokeInlineAgentInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeInlineAgentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeInlineAgent")
	}

	var r0 *bedrockagentruntime.InvokeInlineAgentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeInlineAgentInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeInlineAgentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeInlineAgentInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.InvokeInlineAgentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.InvokeInlineAgentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.InvokeInlineAgentInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptimizePrompt provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) OptimizePrompt(ctx context.Context, params *bedrockagentruntime.OptimizePromptInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.OptimizePromptOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for OptimizePrompt")
	}

	var r0 *bedrockagentruntime.OptimizePromptOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.OptimizePromptInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.OptimizePromptOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.OptimizePromptInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.OptimizePromptOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.OptimizePromptOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.OptimizePromptInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() bedrockagentruntime.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bedrockagentruntime.Options
	if rf, ok := ret.Get(0).(func() bedrockagentruntime.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bedrockagentruntime.Options)
	}

	return r0
}

// Rerank provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Rerank(ctx context.Context, params *bedrockagentruntime.RerankInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RerankOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Rerank")
	}

	var r0 *bedrockagentruntime.RerankOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RerankInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RerankOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RerankInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RerankOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RerankOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RerankInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Retrieve provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Retrieve(ctx context.Context, params *bedrockagentruntime.RetrieveInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Retrieve")
	}

	var r0 *bedrockagentruntime.RetrieveOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RetrieveOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RetrieveOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAndGenerate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RetrieveAndGenerate(ctx context.Context, params *bedrockagentruntime.RetrieveAndGenerateInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAndGenerate")
	}

	var r0 *bedrockagentruntime.RetrieveAndGenerateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RetrieveAndGenerateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RetrieveAndGenerateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAndGenerateStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RetrieveAndGenerateStream(ctx context.Context, params *bedrockagentruntime.RetrieveAndGenerateStreamInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAndGenerateStream")
	}

	var r0 *bedrockagentruntime.RetrieveAndGenerateStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateStreamInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateStreamInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RetrieveAndGenerateStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RetrieveAndGenerateStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateStreamInput, ...func(*bedrockagentruntime.Options)) error); ok {
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
