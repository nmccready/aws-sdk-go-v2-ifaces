// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	forecastquery "github.com/aws/aws-sdk-go-v2/service/forecastquery"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() forecastquery.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 forecastquery.Options
	if rf, ok := ret.Get(0).(func() forecastquery.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(forecastquery.Options)
	}

	return r0
}

// QueryForecast provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) QueryForecast(ctx context.Context, params *forecastquery.QueryForecastInput, optFns ...func(*forecastquery.Options)) (*forecastquery.QueryForecastOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryForecast")
	}

	var r0 *forecastquery.QueryForecastOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *forecastquery.QueryForecastInput, ...func(*forecastquery.Options)) (*forecastquery.QueryForecastOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *forecastquery.QueryForecastInput, ...func(*forecastquery.Options)) *forecastquery.QueryForecastOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forecastquery.QueryForecastOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *forecastquery.QueryForecastInput, ...func(*forecastquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWhatIfForecast provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) QueryWhatIfForecast(ctx context.Context, params *forecastquery.QueryWhatIfForecastInput, optFns ...func(*forecastquery.Options)) (*forecastquery.QueryWhatIfForecastOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryWhatIfForecast")
	}

	var r0 *forecastquery.QueryWhatIfForecastOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *forecastquery.QueryWhatIfForecastInput, ...func(*forecastquery.Options)) (*forecastquery.QueryWhatIfForecastOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *forecastquery.QueryWhatIfForecastInput, ...func(*forecastquery.Options)) *forecastquery.QueryWhatIfForecastOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forecastquery.QueryWhatIfForecastOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *forecastquery.QueryWhatIfForecastInput, ...func(*forecastquery.Options)) error); ok {
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