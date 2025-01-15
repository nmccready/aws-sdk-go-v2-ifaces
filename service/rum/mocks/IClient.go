// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	rum "github.com/aws/aws-sdk-go-v2/service/rum"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchCreateRumMetricDefinitions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchCreateRumMetricDefinitions(ctx context.Context, params *rum.BatchCreateRumMetricDefinitionsInput, optFns ...func(*rum.Options)) (*rum.BatchCreateRumMetricDefinitionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchCreateRumMetricDefinitions")
	}

	var r0 *rum.BatchCreateRumMetricDefinitionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchCreateRumMetricDefinitionsInput, ...func(*rum.Options)) (*rum.BatchCreateRumMetricDefinitionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchCreateRumMetricDefinitionsInput, ...func(*rum.Options)) *rum.BatchCreateRumMetricDefinitionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.BatchCreateRumMetricDefinitionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.BatchCreateRumMetricDefinitionsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchDeleteRumMetricDefinitions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchDeleteRumMetricDefinitions(ctx context.Context, params *rum.BatchDeleteRumMetricDefinitionsInput, optFns ...func(*rum.Options)) (*rum.BatchDeleteRumMetricDefinitionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeleteRumMetricDefinitions")
	}

	var r0 *rum.BatchDeleteRumMetricDefinitionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchDeleteRumMetricDefinitionsInput, ...func(*rum.Options)) (*rum.BatchDeleteRumMetricDefinitionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchDeleteRumMetricDefinitionsInput, ...func(*rum.Options)) *rum.BatchDeleteRumMetricDefinitionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.BatchDeleteRumMetricDefinitionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.BatchDeleteRumMetricDefinitionsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchGetRumMetricDefinitions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetRumMetricDefinitions(ctx context.Context, params *rum.BatchGetRumMetricDefinitionsInput, optFns ...func(*rum.Options)) (*rum.BatchGetRumMetricDefinitionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetRumMetricDefinitions")
	}

	var r0 *rum.BatchGetRumMetricDefinitionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchGetRumMetricDefinitionsInput, ...func(*rum.Options)) (*rum.BatchGetRumMetricDefinitionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.BatchGetRumMetricDefinitionsInput, ...func(*rum.Options)) *rum.BatchGetRumMetricDefinitionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.BatchGetRumMetricDefinitionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.BatchGetRumMetricDefinitionsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAppMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAppMonitor(ctx context.Context, params *rum.CreateAppMonitorInput, optFns ...func(*rum.Options)) (*rum.CreateAppMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAppMonitor")
	}

	var r0 *rum.CreateAppMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.CreateAppMonitorInput, ...func(*rum.Options)) (*rum.CreateAppMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.CreateAppMonitorInput, ...func(*rum.Options)) *rum.CreateAppMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.CreateAppMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.CreateAppMonitorInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAppMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAppMonitor(ctx context.Context, params *rum.DeleteAppMonitorInput, optFns ...func(*rum.Options)) (*rum.DeleteAppMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAppMonitor")
	}

	var r0 *rum.DeleteAppMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.DeleteAppMonitorInput, ...func(*rum.Options)) (*rum.DeleteAppMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.DeleteAppMonitorInput, ...func(*rum.Options)) *rum.DeleteAppMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.DeleteAppMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.DeleteAppMonitorInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRumMetricsDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRumMetricsDestination(ctx context.Context, params *rum.DeleteRumMetricsDestinationInput, optFns ...func(*rum.Options)) (*rum.DeleteRumMetricsDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRumMetricsDestination")
	}

	var r0 *rum.DeleteRumMetricsDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.DeleteRumMetricsDestinationInput, ...func(*rum.Options)) (*rum.DeleteRumMetricsDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.DeleteRumMetricsDestinationInput, ...func(*rum.Options)) *rum.DeleteRumMetricsDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.DeleteRumMetricsDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.DeleteRumMetricsDestinationInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAppMonitor(ctx context.Context, params *rum.GetAppMonitorInput, optFns ...func(*rum.Options)) (*rum.GetAppMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAppMonitor")
	}

	var r0 *rum.GetAppMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.GetAppMonitorInput, ...func(*rum.Options)) (*rum.GetAppMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.GetAppMonitorInput, ...func(*rum.Options)) *rum.GetAppMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.GetAppMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.GetAppMonitorInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppMonitorData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAppMonitorData(ctx context.Context, params *rum.GetAppMonitorDataInput, optFns ...func(*rum.Options)) (*rum.GetAppMonitorDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAppMonitorData")
	}

	var r0 *rum.GetAppMonitorDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.GetAppMonitorDataInput, ...func(*rum.Options)) (*rum.GetAppMonitorDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.GetAppMonitorDataInput, ...func(*rum.Options)) *rum.GetAppMonitorDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.GetAppMonitorDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.GetAppMonitorDataInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAppMonitors provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAppMonitors(ctx context.Context, params *rum.ListAppMonitorsInput, optFns ...func(*rum.Options)) (*rum.ListAppMonitorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAppMonitors")
	}

	var r0 *rum.ListAppMonitorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListAppMonitorsInput, ...func(*rum.Options)) (*rum.ListAppMonitorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListAppMonitorsInput, ...func(*rum.Options)) *rum.ListAppMonitorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.ListAppMonitorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.ListAppMonitorsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRumMetricsDestinations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRumMetricsDestinations(ctx context.Context, params *rum.ListRumMetricsDestinationsInput, optFns ...func(*rum.Options)) (*rum.ListRumMetricsDestinationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRumMetricsDestinations")
	}

	var r0 *rum.ListRumMetricsDestinationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListRumMetricsDestinationsInput, ...func(*rum.Options)) (*rum.ListRumMetricsDestinationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListRumMetricsDestinationsInput, ...func(*rum.Options)) *rum.ListRumMetricsDestinationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.ListRumMetricsDestinationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.ListRumMetricsDestinationsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *rum.ListTagsForResourceInput, optFns ...func(*rum.Options)) (*rum.ListTagsForResourceOutput, error) {
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

	var r0 *rum.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListTagsForResourceInput, ...func(*rum.Options)) (*rum.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.ListTagsForResourceInput, ...func(*rum.Options)) *rum.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.ListTagsForResourceInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() rum.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 rum.Options
	if rf, ok := ret.Get(0).(func() rum.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rum.Options)
	}

	return r0
}

// PutRumEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutRumEvents(ctx context.Context, params *rum.PutRumEventsInput, optFns ...func(*rum.Options)) (*rum.PutRumEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRumEvents")
	}

	var r0 *rum.PutRumEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.PutRumEventsInput, ...func(*rum.Options)) (*rum.PutRumEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.PutRumEventsInput, ...func(*rum.Options)) *rum.PutRumEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.PutRumEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.PutRumEventsInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutRumMetricsDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutRumMetricsDestination(ctx context.Context, params *rum.PutRumMetricsDestinationInput, optFns ...func(*rum.Options)) (*rum.PutRumMetricsDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRumMetricsDestination")
	}

	var r0 *rum.PutRumMetricsDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.PutRumMetricsDestinationInput, ...func(*rum.Options)) (*rum.PutRumMetricsDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.PutRumMetricsDestinationInput, ...func(*rum.Options)) *rum.PutRumMetricsDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.PutRumMetricsDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.PutRumMetricsDestinationInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *rum.TagResourceInput, optFns ...func(*rum.Options)) (*rum.TagResourceOutput, error) {
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

	var r0 *rum.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.TagResourceInput, ...func(*rum.Options)) (*rum.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.TagResourceInput, ...func(*rum.Options)) *rum.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.TagResourceInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *rum.UntagResourceInput, optFns ...func(*rum.Options)) (*rum.UntagResourceOutput, error) {
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

	var r0 *rum.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UntagResourceInput, ...func(*rum.Options)) (*rum.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UntagResourceInput, ...func(*rum.Options)) *rum.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.UntagResourceInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAppMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAppMonitor(ctx context.Context, params *rum.UpdateAppMonitorInput, optFns ...func(*rum.Options)) (*rum.UpdateAppMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAppMonitor")
	}

	var r0 *rum.UpdateAppMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UpdateAppMonitorInput, ...func(*rum.Options)) (*rum.UpdateAppMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UpdateAppMonitorInput, ...func(*rum.Options)) *rum.UpdateAppMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.UpdateAppMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.UpdateAppMonitorInput, ...func(*rum.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRumMetricDefinition provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRumMetricDefinition(ctx context.Context, params *rum.UpdateRumMetricDefinitionInput, optFns ...func(*rum.Options)) (*rum.UpdateRumMetricDefinitionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRumMetricDefinition")
	}

	var r0 *rum.UpdateRumMetricDefinitionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UpdateRumMetricDefinitionInput, ...func(*rum.Options)) (*rum.UpdateRumMetricDefinitionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rum.UpdateRumMetricDefinitionInput, ...func(*rum.Options)) *rum.UpdateRumMetricDefinitionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rum.UpdateRumMetricDefinitionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rum.UpdateRumMetricDefinitionInput, ...func(*rum.Options)) error); ok {
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
