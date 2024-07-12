// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	artifact "github.com/aws/aws-sdk-go-v2/service/artifact"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAccountSettings(ctx context.Context, params *artifact.GetAccountSettingsInput, optFns ...func(*artifact.Options)) (*artifact.GetAccountSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountSettings")
	}

	var r0 *artifact.GetAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetAccountSettingsInput, ...func(*artifact.Options)) (*artifact.GetAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetAccountSettingsInput, ...func(*artifact.Options)) *artifact.GetAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.GetAccountSettingsInput, ...func(*artifact.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetReport(ctx context.Context, params *artifact.GetReportInput, optFns ...func(*artifact.Options)) (*artifact.GetReportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetReport")
	}

	var r0 *artifact.GetReportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetReportInput, ...func(*artifact.Options)) (*artifact.GetReportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetReportInput, ...func(*artifact.Options)) *artifact.GetReportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetReportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.GetReportInput, ...func(*artifact.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportMetadata provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetReportMetadata(ctx context.Context, params *artifact.GetReportMetadataInput, optFns ...func(*artifact.Options)) (*artifact.GetReportMetadataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetReportMetadata")
	}

	var r0 *artifact.GetReportMetadataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetReportMetadataInput, ...func(*artifact.Options)) (*artifact.GetReportMetadataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetReportMetadataInput, ...func(*artifact.Options)) *artifact.GetReportMetadataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetReportMetadataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.GetReportMetadataInput, ...func(*artifact.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTermForReport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTermForReport(ctx context.Context, params *artifact.GetTermForReportInput, optFns ...func(*artifact.Options)) (*artifact.GetTermForReportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTermForReport")
	}

	var r0 *artifact.GetTermForReportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetTermForReportInput, ...func(*artifact.Options)) (*artifact.GetTermForReportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.GetTermForReportInput, ...func(*artifact.Options)) *artifact.GetTermForReportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetTermForReportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.GetTermForReportInput, ...func(*artifact.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReports provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListReports(ctx context.Context, params *artifact.ListReportsInput, optFns ...func(*artifact.Options)) (*artifact.ListReportsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListReports")
	}

	var r0 *artifact.ListReportsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.ListReportsInput, ...func(*artifact.Options)) (*artifact.ListReportsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.ListReportsInput, ...func(*artifact.Options)) *artifact.ListReportsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.ListReportsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.ListReportsInput, ...func(*artifact.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() artifact.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 artifact.Options
	if rf, ok := ret.Get(0).(func() artifact.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(artifact.Options)
	}

	return r0
}

// PutAccountSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutAccountSettings(ctx context.Context, params *artifact.PutAccountSettingsInput, optFns ...func(*artifact.Options)) (*artifact.PutAccountSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutAccountSettings")
	}

	var r0 *artifact.PutAccountSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.PutAccountSettingsInput, ...func(*artifact.Options)) (*artifact.PutAccountSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.PutAccountSettingsInput, ...func(*artifact.Options)) *artifact.PutAccountSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.PutAccountSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.PutAccountSettingsInput, ...func(*artifact.Options)) error); ok {
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