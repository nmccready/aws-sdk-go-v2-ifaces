// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	pipes "github.com/aws/aws-sdk-go-v2/service/pipes"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreatePipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePipe(ctx context.Context, params *pipes.CreatePipeInput, optFns ...func(*pipes.Options)) (*pipes.CreatePipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePipe")
	}

	var r0 *pipes.CreatePipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.CreatePipeInput, ...func(*pipes.Options)) (*pipes.CreatePipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.CreatePipeInput, ...func(*pipes.Options)) *pipes.CreatePipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.CreatePipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.CreatePipeInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePipe(ctx context.Context, params *pipes.DeletePipeInput, optFns ...func(*pipes.Options)) (*pipes.DeletePipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePipe")
	}

	var r0 *pipes.DeletePipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.DeletePipeInput, ...func(*pipes.Options)) (*pipes.DeletePipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.DeletePipeInput, ...func(*pipes.Options)) *pipes.DeletePipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.DeletePipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.DeletePipeInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribePipe(ctx context.Context, params *pipes.DescribePipeInput, optFns ...func(*pipes.Options)) (*pipes.DescribePipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribePipe")
	}

	var r0 *pipes.DescribePipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.DescribePipeInput, ...func(*pipes.Options)) (*pipes.DescribePipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.DescribePipeInput, ...func(*pipes.Options)) *pipes.DescribePipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.DescribePipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.DescribePipeInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPipes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPipes(ctx context.Context, params *pipes.ListPipesInput, optFns ...func(*pipes.Options)) (*pipes.ListPipesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPipes")
	}

	var r0 *pipes.ListPipesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.ListPipesInput, ...func(*pipes.Options)) (*pipes.ListPipesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.ListPipesInput, ...func(*pipes.Options)) *pipes.ListPipesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.ListPipesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.ListPipesInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *pipes.ListTagsForResourceInput, optFns ...func(*pipes.Options)) (*pipes.ListTagsForResourceOutput, error) {
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

	var r0 *pipes.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.ListTagsForResourceInput, ...func(*pipes.Options)) (*pipes.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.ListTagsForResourceInput, ...func(*pipes.Options)) *pipes.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.ListTagsForResourceInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() pipes.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 pipes.Options
	if rf, ok := ret.Get(0).(func() pipes.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pipes.Options)
	}

	return r0
}

// StartPipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartPipe(ctx context.Context, params *pipes.StartPipeInput, optFns ...func(*pipes.Options)) (*pipes.StartPipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartPipe")
	}

	var r0 *pipes.StartPipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.StartPipeInput, ...func(*pipes.Options)) (*pipes.StartPipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.StartPipeInput, ...func(*pipes.Options)) *pipes.StartPipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.StartPipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.StartPipeInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopPipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopPipe(ctx context.Context, params *pipes.StopPipeInput, optFns ...func(*pipes.Options)) (*pipes.StopPipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopPipe")
	}

	var r0 *pipes.StopPipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.StopPipeInput, ...func(*pipes.Options)) (*pipes.StopPipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.StopPipeInput, ...func(*pipes.Options)) *pipes.StopPipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.StopPipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.StopPipeInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *pipes.TagResourceInput, optFns ...func(*pipes.Options)) (*pipes.TagResourceOutput, error) {
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

	var r0 *pipes.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.TagResourceInput, ...func(*pipes.Options)) (*pipes.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.TagResourceInput, ...func(*pipes.Options)) *pipes.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.TagResourceInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *pipes.UntagResourceInput, optFns ...func(*pipes.Options)) (*pipes.UntagResourceOutput, error) {
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

	var r0 *pipes.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.UntagResourceInput, ...func(*pipes.Options)) (*pipes.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.UntagResourceInput, ...func(*pipes.Options)) *pipes.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.UntagResourceInput, ...func(*pipes.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePipe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePipe(ctx context.Context, params *pipes.UpdatePipeInput, optFns ...func(*pipes.Options)) (*pipes.UpdatePipeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePipe")
	}

	var r0 *pipes.UpdatePipeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.UpdatePipeInput, ...func(*pipes.Options)) (*pipes.UpdatePipeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipes.UpdatePipeInput, ...func(*pipes.Options)) *pipes.UpdatePipeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipes.UpdatePipeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipes.UpdatePipeInput, ...func(*pipes.Options)) error); ok {
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
