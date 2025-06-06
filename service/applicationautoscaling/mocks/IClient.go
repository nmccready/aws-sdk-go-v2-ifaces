// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	applicationautoscaling "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DeleteScalingPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteScalingPolicy(ctx context.Context, params *applicationautoscaling.DeleteScalingPolicyInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScalingPolicy")
	}

	var r0 *applicationautoscaling.DeleteScalingPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeleteScalingPolicyInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeleteScalingPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeleteScalingPolicyInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DeleteScalingPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DeleteScalingPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DeleteScalingPolicyInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScheduledAction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteScheduledAction(ctx context.Context, params *applicationautoscaling.DeleteScheduledActionInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduledAction")
	}

	var r0 *applicationautoscaling.DeleteScheduledActionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeleteScheduledActionInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeleteScheduledActionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeleteScheduledActionInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DeleteScheduledActionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DeleteScheduledActionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DeleteScheduledActionInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterScalableTarget provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeregisterScalableTarget(ctx context.Context, params *applicationautoscaling.DeregisterScalableTargetInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterScalableTarget")
	}

	var r0 *applicationautoscaling.DeregisterScalableTargetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeregisterScalableTargetInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DeregisterScalableTargetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DeregisterScalableTargetInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DeregisterScalableTargetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DeregisterScalableTargetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DeregisterScalableTargetInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScalableTargets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScalableTargets(ctx context.Context, params *applicationautoscaling.DescribeScalableTargetsInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScalableTargets")
	}

	var r0 *applicationautoscaling.DescribeScalableTargetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalableTargetsInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalableTargetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalableTargetsInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DescribeScalableTargetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DescribeScalableTargetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DescribeScalableTargetsInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScalingActivities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScalingActivities(ctx context.Context, params *applicationautoscaling.DescribeScalingActivitiesInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScalingActivities")
	}

	var r0 *applicationautoscaling.DescribeScalingActivitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalingActivitiesInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalingActivitiesInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DescribeScalingActivitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DescribeScalingActivitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DescribeScalingActivitiesInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScalingPolicies provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScalingPolicies(ctx context.Context, params *applicationautoscaling.DescribeScalingPoliciesInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScalingPolicies")
	}

	var r0 *applicationautoscaling.DescribeScalingPoliciesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalingPoliciesInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScalingPoliciesInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DescribeScalingPoliciesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DescribeScalingPoliciesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DescribeScalingPoliciesInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledActions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeScheduledActions(ctx context.Context, params *applicationautoscaling.DescribeScheduledActionsInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScheduledActions")
	}

	var r0 *applicationautoscaling.DescribeScheduledActionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScheduledActionsInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScheduledActionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.DescribeScheduledActionsInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.DescribeScheduledActionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.DescribeScheduledActionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.DescribeScheduledActionsInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPredictiveScalingForecast provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPredictiveScalingForecast(ctx context.Context, params *applicationautoscaling.GetPredictiveScalingForecastInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.GetPredictiveScalingForecastOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPredictiveScalingForecast")
	}

	var r0 *applicationautoscaling.GetPredictiveScalingForecastOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.GetPredictiveScalingForecastInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.GetPredictiveScalingForecastOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.GetPredictiveScalingForecastInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.GetPredictiveScalingForecastOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.GetPredictiveScalingForecastOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.GetPredictiveScalingForecastInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *applicationautoscaling.ListTagsForResourceInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.ListTagsForResourceOutput, error) {
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

	var r0 *applicationautoscaling.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.ListTagsForResourceInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.ListTagsForResourceInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.ListTagsForResourceInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() applicationautoscaling.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 applicationautoscaling.Options
	if rf, ok := ret.Get(0).(func() applicationautoscaling.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(applicationautoscaling.Options)
	}

	return r0
}

// PutScalingPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutScalingPolicy(ctx context.Context, params *applicationautoscaling.PutScalingPolicyInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutScalingPolicy")
	}

	var r0 *applicationautoscaling.PutScalingPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.PutScalingPolicyInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.PutScalingPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.PutScalingPolicyInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.PutScalingPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.PutScalingPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.PutScalingPolicyInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutScheduledAction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutScheduledAction(ctx context.Context, params *applicationautoscaling.PutScheduledActionInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.PutScheduledActionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutScheduledAction")
	}

	var r0 *applicationautoscaling.PutScheduledActionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.PutScheduledActionInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.PutScheduledActionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.PutScheduledActionInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.PutScheduledActionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.PutScheduledActionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.PutScheduledActionInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterScalableTarget provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RegisterScalableTarget(ctx context.Context, params *applicationautoscaling.RegisterScalableTargetInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterScalableTarget")
	}

	var r0 *applicationautoscaling.RegisterScalableTargetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.RegisterScalableTargetInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.RegisterScalableTargetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.RegisterScalableTargetInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.RegisterScalableTargetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.RegisterScalableTargetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.RegisterScalableTargetInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *applicationautoscaling.TagResourceInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.TagResourceOutput, error) {
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

	var r0 *applicationautoscaling.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.TagResourceInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.TagResourceInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.TagResourceInput, ...func(*applicationautoscaling.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *applicationautoscaling.UntagResourceInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.UntagResourceOutput, error) {
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

	var r0 *applicationautoscaling.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.UntagResourceInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *applicationautoscaling.UntagResourceInput, ...func(*applicationautoscaling.Options)) *applicationautoscaling.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*applicationautoscaling.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *applicationautoscaling.UntagResourceInput, ...func(*applicationautoscaling.Options)) error); ok {
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
