// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	costoptimizationhub "github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetPreferences provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPreferences(ctx context.Context, params *costoptimizationhub.GetPreferencesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetPreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPreferences")
	}

	var r0 *costoptimizationhub.GetPreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.GetPreferencesInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetPreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.GetPreferencesInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.GetPreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.GetPreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.GetPreferencesInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecommendation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRecommendation(ctx context.Context, params *costoptimizationhub.GetRecommendationInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetRecommendationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRecommendation")
	}

	var r0 *costoptimizationhub.GetRecommendationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.GetRecommendationInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetRecommendationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.GetRecommendationInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.GetRecommendationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.GetRecommendationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.GetRecommendationInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnrollmentStatuses provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListEnrollmentStatuses(ctx context.Context, params *costoptimizationhub.ListEnrollmentStatusesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListEnrollmentStatusesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEnrollmentStatuses")
	}

	var r0 *costoptimizationhub.ListEnrollmentStatusesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListEnrollmentStatusesInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListEnrollmentStatusesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListEnrollmentStatusesInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.ListEnrollmentStatusesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.ListEnrollmentStatusesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.ListEnrollmentStatusesInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecommendationSummaries provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRecommendationSummaries(ctx context.Context, params *costoptimizationhub.ListRecommendationSummariesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationSummariesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRecommendationSummaries")
	}

	var r0 *costoptimizationhub.ListRecommendationSummariesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListRecommendationSummariesInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationSummariesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListRecommendationSummariesInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.ListRecommendationSummariesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.ListRecommendationSummariesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.ListRecommendationSummariesInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecommendations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRecommendations(ctx context.Context, params *costoptimizationhub.ListRecommendationsInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRecommendations")
	}

	var r0 *costoptimizationhub.ListRecommendationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListRecommendationsInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.ListRecommendationsInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.ListRecommendationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.ListRecommendationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.ListRecommendationsInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() costoptimizationhub.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 costoptimizationhub.Options
	if rf, ok := ret.Get(0).(func() costoptimizationhub.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(costoptimizationhub.Options)
	}

	return r0
}

// UpdateEnrollmentStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateEnrollmentStatus(ctx context.Context, params *costoptimizationhub.UpdateEnrollmentStatusInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdateEnrollmentStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEnrollmentStatus")
	}

	var r0 *costoptimizationhub.UpdateEnrollmentStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.UpdateEnrollmentStatusInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdateEnrollmentStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.UpdateEnrollmentStatusInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.UpdateEnrollmentStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.UpdateEnrollmentStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.UpdateEnrollmentStatusInput, ...func(*costoptimizationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePreferences provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePreferences(ctx context.Context, params *costoptimizationhub.UpdatePreferencesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdatePreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePreferences")
	}

	var r0 *costoptimizationhub.UpdatePreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.UpdatePreferencesInput, ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdatePreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *costoptimizationhub.UpdatePreferencesInput, ...func(*costoptimizationhub.Options)) *costoptimizationhub.UpdatePreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*costoptimizationhub.UpdatePreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *costoptimizationhub.UpdatePreferencesInput, ...func(*costoptimizationhub.Options)) error); ok {
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
