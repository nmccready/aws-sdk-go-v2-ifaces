// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	pi "github.com/aws/aws-sdk-go-v2/service/pi"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreatePerformanceAnalysisReport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePerformanceAnalysisReport(ctx context.Context, params *pi.CreatePerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.CreatePerformanceAnalysisReportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePerformanceAnalysisReport")
	}

	var r0 *pi.CreatePerformanceAnalysisReportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.CreatePerformanceAnalysisReportInput, ...func(*pi.Options)) (*pi.CreatePerformanceAnalysisReportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.CreatePerformanceAnalysisReportInput, ...func(*pi.Options)) *pi.CreatePerformanceAnalysisReportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.CreatePerformanceAnalysisReportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.CreatePerformanceAnalysisReportInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePerformanceAnalysisReport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePerformanceAnalysisReport(ctx context.Context, params *pi.DeletePerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.DeletePerformanceAnalysisReportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePerformanceAnalysisReport")
	}

	var r0 *pi.DeletePerformanceAnalysisReportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.DeletePerformanceAnalysisReportInput, ...func(*pi.Options)) (*pi.DeletePerformanceAnalysisReportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.DeletePerformanceAnalysisReportInput, ...func(*pi.Options)) *pi.DeletePerformanceAnalysisReportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.DeletePerformanceAnalysisReportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.DeletePerformanceAnalysisReportInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDimensionKeys provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDimensionKeys(ctx context.Context, params *pi.DescribeDimensionKeysInput, optFns ...func(*pi.Options)) (*pi.DescribeDimensionKeysOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDimensionKeys")
	}

	var r0 *pi.DescribeDimensionKeysOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.DescribeDimensionKeysInput, ...func(*pi.Options)) (*pi.DescribeDimensionKeysOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.DescribeDimensionKeysInput, ...func(*pi.Options)) *pi.DescribeDimensionKeysOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.DescribeDimensionKeysOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.DescribeDimensionKeysInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDimensionKeyDetails provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDimensionKeyDetails(ctx context.Context, params *pi.GetDimensionKeyDetailsInput, optFns ...func(*pi.Options)) (*pi.GetDimensionKeyDetailsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDimensionKeyDetails")
	}

	var r0 *pi.GetDimensionKeyDetailsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetDimensionKeyDetailsInput, ...func(*pi.Options)) (*pi.GetDimensionKeyDetailsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetDimensionKeyDetailsInput, ...func(*pi.Options)) *pi.GetDimensionKeyDetailsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.GetDimensionKeyDetailsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.GetDimensionKeyDetailsInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerformanceAnalysisReport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPerformanceAnalysisReport(ctx context.Context, params *pi.GetPerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.GetPerformanceAnalysisReportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPerformanceAnalysisReport")
	}

	var r0 *pi.GetPerformanceAnalysisReportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetPerformanceAnalysisReportInput, ...func(*pi.Options)) (*pi.GetPerformanceAnalysisReportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetPerformanceAnalysisReportInput, ...func(*pi.Options)) *pi.GetPerformanceAnalysisReportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.GetPerformanceAnalysisReportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.GetPerformanceAnalysisReportInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceMetadata provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourceMetadata(ctx context.Context, params *pi.GetResourceMetadataInput, optFns ...func(*pi.Options)) (*pi.GetResourceMetadataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceMetadata")
	}

	var r0 *pi.GetResourceMetadataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetResourceMetadataInput, ...func(*pi.Options)) (*pi.GetResourceMetadataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetResourceMetadataInput, ...func(*pi.Options)) *pi.GetResourceMetadataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.GetResourceMetadataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.GetResourceMetadataInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceMetrics provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourceMetrics(ctx context.Context, params *pi.GetResourceMetricsInput, optFns ...func(*pi.Options)) (*pi.GetResourceMetricsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceMetrics")
	}

	var r0 *pi.GetResourceMetricsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetResourceMetricsInput, ...func(*pi.Options)) (*pi.GetResourceMetricsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.GetResourceMetricsInput, ...func(*pi.Options)) *pi.GetResourceMetricsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.GetResourceMetricsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.GetResourceMetricsInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAvailableResourceDimensions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAvailableResourceDimensions(ctx context.Context, params *pi.ListAvailableResourceDimensionsInput, optFns ...func(*pi.Options)) (*pi.ListAvailableResourceDimensionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAvailableResourceDimensions")
	}

	var r0 *pi.ListAvailableResourceDimensionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListAvailableResourceDimensionsInput, ...func(*pi.Options)) (*pi.ListAvailableResourceDimensionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListAvailableResourceDimensionsInput, ...func(*pi.Options)) *pi.ListAvailableResourceDimensionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.ListAvailableResourceDimensionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.ListAvailableResourceDimensionsInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAvailableResourceMetrics provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAvailableResourceMetrics(ctx context.Context, params *pi.ListAvailableResourceMetricsInput, optFns ...func(*pi.Options)) (*pi.ListAvailableResourceMetricsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAvailableResourceMetrics")
	}

	var r0 *pi.ListAvailableResourceMetricsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListAvailableResourceMetricsInput, ...func(*pi.Options)) (*pi.ListAvailableResourceMetricsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListAvailableResourceMetricsInput, ...func(*pi.Options)) *pi.ListAvailableResourceMetricsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.ListAvailableResourceMetricsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.ListAvailableResourceMetricsInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPerformanceAnalysisReports provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPerformanceAnalysisReports(ctx context.Context, params *pi.ListPerformanceAnalysisReportsInput, optFns ...func(*pi.Options)) (*pi.ListPerformanceAnalysisReportsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPerformanceAnalysisReports")
	}

	var r0 *pi.ListPerformanceAnalysisReportsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListPerformanceAnalysisReportsInput, ...func(*pi.Options)) (*pi.ListPerformanceAnalysisReportsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListPerformanceAnalysisReportsInput, ...func(*pi.Options)) *pi.ListPerformanceAnalysisReportsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.ListPerformanceAnalysisReportsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.ListPerformanceAnalysisReportsInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *pi.ListTagsForResourceInput, optFns ...func(*pi.Options)) (*pi.ListTagsForResourceOutput, error) {
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

	var r0 *pi.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListTagsForResourceInput, ...func(*pi.Options)) (*pi.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.ListTagsForResourceInput, ...func(*pi.Options)) *pi.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.ListTagsForResourceInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() pi.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 pi.Options
	if rf, ok := ret.Get(0).(func() pi.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pi.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *pi.TagResourceInput, optFns ...func(*pi.Options)) (*pi.TagResourceOutput, error) {
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

	var r0 *pi.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.TagResourceInput, ...func(*pi.Options)) (*pi.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.TagResourceInput, ...func(*pi.Options)) *pi.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.TagResourceInput, ...func(*pi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *pi.UntagResourceInput, optFns ...func(*pi.Options)) (*pi.UntagResourceOutput, error) {
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

	var r0 *pi.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pi.UntagResourceInput, ...func(*pi.Options)) (*pi.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pi.UntagResourceInput, ...func(*pi.Options)) *pi.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pi.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pi.UntagResourceInput, ...func(*pi.Options)) error); ok {
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
