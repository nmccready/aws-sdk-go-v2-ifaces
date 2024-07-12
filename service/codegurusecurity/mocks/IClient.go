// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	codegurusecurity "github.com/aws/aws-sdk-go-v2/service/codegurusecurity"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchGetFindings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetFindings(ctx context.Context, params *codegurusecurity.BatchGetFindingsInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.BatchGetFindingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetFindings")
	}

	var r0 *codegurusecurity.BatchGetFindingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.BatchGetFindingsInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.BatchGetFindingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.BatchGetFindingsInput, ...func(*codegurusecurity.Options)) *codegurusecurity.BatchGetFindingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.BatchGetFindingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.BatchGetFindingsInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScan provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateScan(ctx context.Context, params *codegurusecurity.CreateScanInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.CreateScanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateScan")
	}

	var r0 *codegurusecurity.CreateScanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.CreateScanInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.CreateScanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.CreateScanInput, ...func(*codegurusecurity.Options)) *codegurusecurity.CreateScanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.CreateScanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.CreateScanInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUploadUrl provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateUploadUrl(ctx context.Context, params *codegurusecurity.CreateUploadUrlInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.CreateUploadUrlOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateUploadUrl")
	}

	var r0 *codegurusecurity.CreateUploadUrlOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.CreateUploadUrlInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.CreateUploadUrlOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.CreateUploadUrlInput, ...func(*codegurusecurity.Options)) *codegurusecurity.CreateUploadUrlOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.CreateUploadUrlOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.CreateUploadUrlInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAccountConfiguration(ctx context.Context, params *codegurusecurity.GetAccountConfigurationInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.GetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountConfiguration")
	}

	var r0 *codegurusecurity.GetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetAccountConfigurationInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.GetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetAccountConfigurationInput, ...func(*codegurusecurity.Options)) *codegurusecurity.GetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.GetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.GetAccountConfigurationInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFindings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetFindings(ctx context.Context, params *codegurusecurity.GetFindingsInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.GetFindingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFindings")
	}

	var r0 *codegurusecurity.GetFindingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetFindingsInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.GetFindingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetFindingsInput, ...func(*codegurusecurity.Options)) *codegurusecurity.GetFindingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.GetFindingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.GetFindingsInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetricsSummary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMetricsSummary(ctx context.Context, params *codegurusecurity.GetMetricsSummaryInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.GetMetricsSummaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMetricsSummary")
	}

	var r0 *codegurusecurity.GetMetricsSummaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetMetricsSummaryInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.GetMetricsSummaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetMetricsSummaryInput, ...func(*codegurusecurity.Options)) *codegurusecurity.GetMetricsSummaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.GetMetricsSummaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.GetMetricsSummaryInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScan provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetScan(ctx context.Context, params *codegurusecurity.GetScanInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.GetScanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetScan")
	}

	var r0 *codegurusecurity.GetScanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetScanInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.GetScanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.GetScanInput, ...func(*codegurusecurity.Options)) *codegurusecurity.GetScanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.GetScanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.GetScanInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFindingsMetrics provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListFindingsMetrics(ctx context.Context, params *codegurusecurity.ListFindingsMetricsInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.ListFindingsMetricsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFindingsMetrics")
	}

	var r0 *codegurusecurity.ListFindingsMetricsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListFindingsMetricsInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.ListFindingsMetricsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListFindingsMetricsInput, ...func(*codegurusecurity.Options)) *codegurusecurity.ListFindingsMetricsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.ListFindingsMetricsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.ListFindingsMetricsInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScans provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListScans(ctx context.Context, params *codegurusecurity.ListScansInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.ListScansOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListScans")
	}

	var r0 *codegurusecurity.ListScansOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListScansInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.ListScansOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListScansInput, ...func(*codegurusecurity.Options)) *codegurusecurity.ListScansOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.ListScansOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.ListScansInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *codegurusecurity.ListTagsForResourceInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.ListTagsForResourceOutput, error) {
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

	var r0 *codegurusecurity.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListTagsForResourceInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.ListTagsForResourceInput, ...func(*codegurusecurity.Options)) *codegurusecurity.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.ListTagsForResourceInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() codegurusecurity.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 codegurusecurity.Options
	if rf, ok := ret.Get(0).(func() codegurusecurity.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(codegurusecurity.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *codegurusecurity.TagResourceInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.TagResourceOutput, error) {
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

	var r0 *codegurusecurity.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.TagResourceInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.TagResourceInput, ...func(*codegurusecurity.Options)) *codegurusecurity.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.TagResourceInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *codegurusecurity.UntagResourceInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.UntagResourceOutput, error) {
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

	var r0 *codegurusecurity.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.UntagResourceInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.UntagResourceInput, ...func(*codegurusecurity.Options)) *codegurusecurity.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.UntagResourceInput, ...func(*codegurusecurity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAccountConfiguration(ctx context.Context, params *codegurusecurity.UpdateAccountConfigurationInput, optFns ...func(*codegurusecurity.Options)) (*codegurusecurity.UpdateAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountConfiguration")
	}

	var r0 *codegurusecurity.UpdateAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.UpdateAccountConfigurationInput, ...func(*codegurusecurity.Options)) (*codegurusecurity.UpdateAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codegurusecurity.UpdateAccountConfigurationInput, ...func(*codegurusecurity.Options)) *codegurusecurity.UpdateAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codegurusecurity.UpdateAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codegurusecurity.UpdateAccountConfigurationInput, ...func(*codegurusecurity.Options)) error); ok {
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
