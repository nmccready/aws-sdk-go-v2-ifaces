// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	rbin "github.com/aws/aws-sdk-go-v2/service/rbin"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRule(ctx context.Context, params *rbin.CreateRuleInput, optFns ...func(*rbin.Options)) (*rbin.CreateRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRule")
	}

	var r0 *rbin.CreateRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.CreateRuleInput, ...func(*rbin.Options)) (*rbin.CreateRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.CreateRuleInput, ...func(*rbin.Options)) *rbin.CreateRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.CreateRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.CreateRuleInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRule(ctx context.Context, params *rbin.DeleteRuleInput, optFns ...func(*rbin.Options)) (*rbin.DeleteRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRule")
	}

	var r0 *rbin.DeleteRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.DeleteRuleInput, ...func(*rbin.Options)) (*rbin.DeleteRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.DeleteRuleInput, ...func(*rbin.Options)) *rbin.DeleteRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.DeleteRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.DeleteRuleInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRule(ctx context.Context, params *rbin.GetRuleInput, optFns ...func(*rbin.Options)) (*rbin.GetRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRule")
	}

	var r0 *rbin.GetRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.GetRuleInput, ...func(*rbin.Options)) (*rbin.GetRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.GetRuleInput, ...func(*rbin.Options)) *rbin.GetRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.GetRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.GetRuleInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRules provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRules(ctx context.Context, params *rbin.ListRulesInput, optFns ...func(*rbin.Options)) (*rbin.ListRulesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRules")
	}

	var r0 *rbin.ListRulesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.ListRulesInput, ...func(*rbin.Options)) (*rbin.ListRulesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.ListRulesInput, ...func(*rbin.Options)) *rbin.ListRulesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.ListRulesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.ListRulesInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *rbin.ListTagsForResourceInput, optFns ...func(*rbin.Options)) (*rbin.ListTagsForResourceOutput, error) {
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

	var r0 *rbin.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.ListTagsForResourceInput, ...func(*rbin.Options)) (*rbin.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.ListTagsForResourceInput, ...func(*rbin.Options)) *rbin.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.ListTagsForResourceInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) LockRule(ctx context.Context, params *rbin.LockRuleInput, optFns ...func(*rbin.Options)) (*rbin.LockRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for LockRule")
	}

	var r0 *rbin.LockRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.LockRuleInput, ...func(*rbin.Options)) (*rbin.LockRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.LockRuleInput, ...func(*rbin.Options)) *rbin.LockRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.LockRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.LockRuleInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() rbin.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 rbin.Options
	if rf, ok := ret.Get(0).(func() rbin.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rbin.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *rbin.TagResourceInput, optFns ...func(*rbin.Options)) (*rbin.TagResourceOutput, error) {
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

	var r0 *rbin.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.TagResourceInput, ...func(*rbin.Options)) (*rbin.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.TagResourceInput, ...func(*rbin.Options)) *rbin.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.TagResourceInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlockRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UnlockRule(ctx context.Context, params *rbin.UnlockRuleInput, optFns ...func(*rbin.Options)) (*rbin.UnlockRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnlockRule")
	}

	var r0 *rbin.UnlockRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UnlockRuleInput, ...func(*rbin.Options)) (*rbin.UnlockRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UnlockRuleInput, ...func(*rbin.Options)) *rbin.UnlockRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.UnlockRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.UnlockRuleInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *rbin.UntagResourceInput, optFns ...func(*rbin.Options)) (*rbin.UntagResourceOutput, error) {
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

	var r0 *rbin.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UntagResourceInput, ...func(*rbin.Options)) (*rbin.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UntagResourceInput, ...func(*rbin.Options)) *rbin.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.UntagResourceInput, ...func(*rbin.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRule provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRule(ctx context.Context, params *rbin.UpdateRuleInput, optFns ...func(*rbin.Options)) (*rbin.UpdateRuleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRule")
	}

	var r0 *rbin.UpdateRuleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UpdateRuleInput, ...func(*rbin.Options)) (*rbin.UpdateRuleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbin.UpdateRuleInput, ...func(*rbin.Options)) *rbin.UpdateRuleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbin.UpdateRuleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbin.UpdateRuleInput, ...func(*rbin.Options)) error); ok {
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
