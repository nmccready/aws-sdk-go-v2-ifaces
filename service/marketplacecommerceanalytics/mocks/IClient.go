// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplacecommerceanalytics "github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GenerateDataSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateDataSet(ctx context.Context, params *marketplacecommerceanalytics.GenerateDataSetInput, optFns ...func(*marketplacecommerceanalytics.Options)) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateDataSet")
	}

	var r0 *marketplacecommerceanalytics.GenerateDataSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecommerceanalytics.GenerateDataSetInput, ...func(*marketplacecommerceanalytics.Options)) (*marketplacecommerceanalytics.GenerateDataSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecommerceanalytics.GenerateDataSetInput, ...func(*marketplacecommerceanalytics.Options)) *marketplacecommerceanalytics.GenerateDataSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecommerceanalytics.GenerateDataSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecommerceanalytics.GenerateDataSetInput, ...func(*marketplacecommerceanalytics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() marketplacecommerceanalytics.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplacecommerceanalytics.Options
	if rf, ok := ret.Get(0).(func() marketplacecommerceanalytics.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplacecommerceanalytics.Options)
	}

	return r0
}

// StartSupportDataExport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartSupportDataExport(ctx context.Context, params *marketplacecommerceanalytics.StartSupportDataExportInput, optFns ...func(*marketplacecommerceanalytics.Options)) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartSupportDataExport")
	}

	var r0 *marketplacecommerceanalytics.StartSupportDataExportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecommerceanalytics.StartSupportDataExportInput, ...func(*marketplacecommerceanalytics.Options)) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplacecommerceanalytics.StartSupportDataExportInput, ...func(*marketplacecommerceanalytics.Options)) *marketplacecommerceanalytics.StartSupportDataExportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplacecommerceanalytics.StartSupportDataExportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplacecommerceanalytics.StartSupportDataExportInput, ...func(*marketplacecommerceanalytics.Options)) error); ok {
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