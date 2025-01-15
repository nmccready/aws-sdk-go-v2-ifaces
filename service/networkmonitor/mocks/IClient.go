// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	networkmonitor "github.com/aws/aws-sdk-go-v2/service/networkmonitor"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateMonitor(ctx context.Context, params *networkmonitor.CreateMonitorInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.CreateMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMonitor")
	}

	var r0 *networkmonitor.CreateMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.CreateMonitorInput, ...func(*networkmonitor.Options)) (*networkmonitor.CreateMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.CreateMonitorInput, ...func(*networkmonitor.Options)) *networkmonitor.CreateMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.CreateMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.CreateMonitorInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProbe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProbe(ctx context.Context, params *networkmonitor.CreateProbeInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.CreateProbeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProbe")
	}

	var r0 *networkmonitor.CreateProbeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.CreateProbeInput, ...func(*networkmonitor.Options)) (*networkmonitor.CreateProbeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.CreateProbeInput, ...func(*networkmonitor.Options)) *networkmonitor.CreateProbeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.CreateProbeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.CreateProbeInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMonitor(ctx context.Context, params *networkmonitor.DeleteMonitorInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.DeleteMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMonitor")
	}

	var r0 *networkmonitor.DeleteMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.DeleteMonitorInput, ...func(*networkmonitor.Options)) (*networkmonitor.DeleteMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.DeleteMonitorInput, ...func(*networkmonitor.Options)) *networkmonitor.DeleteMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.DeleteMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.DeleteMonitorInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProbe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProbe(ctx context.Context, params *networkmonitor.DeleteProbeInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.DeleteProbeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProbe")
	}

	var r0 *networkmonitor.DeleteProbeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.DeleteProbeInput, ...func(*networkmonitor.Options)) (*networkmonitor.DeleteProbeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.DeleteProbeInput, ...func(*networkmonitor.Options)) *networkmonitor.DeleteProbeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.DeleteProbeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.DeleteProbeInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMonitor(ctx context.Context, params *networkmonitor.GetMonitorInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.GetMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMonitor")
	}

	var r0 *networkmonitor.GetMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.GetMonitorInput, ...func(*networkmonitor.Options)) (*networkmonitor.GetMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.GetMonitorInput, ...func(*networkmonitor.Options)) *networkmonitor.GetMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.GetMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.GetMonitorInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProbe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetProbe(ctx context.Context, params *networkmonitor.GetProbeInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.GetProbeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProbe")
	}

	var r0 *networkmonitor.GetProbeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.GetProbeInput, ...func(*networkmonitor.Options)) (*networkmonitor.GetProbeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.GetProbeInput, ...func(*networkmonitor.Options)) *networkmonitor.GetProbeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.GetProbeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.GetProbeInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMonitors provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListMonitors(ctx context.Context, params *networkmonitor.ListMonitorsInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.ListMonitorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMonitors")
	}

	var r0 *networkmonitor.ListMonitorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.ListMonitorsInput, ...func(*networkmonitor.Options)) (*networkmonitor.ListMonitorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.ListMonitorsInput, ...func(*networkmonitor.Options)) *networkmonitor.ListMonitorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.ListMonitorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.ListMonitorsInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *networkmonitor.ListTagsForResourceInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.ListTagsForResourceOutput, error) {
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

	var r0 *networkmonitor.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.ListTagsForResourceInput, ...func(*networkmonitor.Options)) (*networkmonitor.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.ListTagsForResourceInput, ...func(*networkmonitor.Options)) *networkmonitor.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.ListTagsForResourceInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() networkmonitor.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 networkmonitor.Options
	if rf, ok := ret.Get(0).(func() networkmonitor.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(networkmonitor.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *networkmonitor.TagResourceInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.TagResourceOutput, error) {
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

	var r0 *networkmonitor.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.TagResourceInput, ...func(*networkmonitor.Options)) (*networkmonitor.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.TagResourceInput, ...func(*networkmonitor.Options)) *networkmonitor.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.TagResourceInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *networkmonitor.UntagResourceInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.UntagResourceOutput, error) {
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

	var r0 *networkmonitor.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UntagResourceInput, ...func(*networkmonitor.Options)) (*networkmonitor.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UntagResourceInput, ...func(*networkmonitor.Options)) *networkmonitor.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.UntagResourceInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMonitor provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateMonitor(ctx context.Context, params *networkmonitor.UpdateMonitorInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.UpdateMonitorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMonitor")
	}

	var r0 *networkmonitor.UpdateMonitorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UpdateMonitorInput, ...func(*networkmonitor.Options)) (*networkmonitor.UpdateMonitorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UpdateMonitorInput, ...func(*networkmonitor.Options)) *networkmonitor.UpdateMonitorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.UpdateMonitorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.UpdateMonitorInput, ...func(*networkmonitor.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProbe provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateProbe(ctx context.Context, params *networkmonitor.UpdateProbeInput, optFns ...func(*networkmonitor.Options)) (*networkmonitor.UpdateProbeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProbe")
	}

	var r0 *networkmonitor.UpdateProbeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UpdateProbeInput, ...func(*networkmonitor.Options)) (*networkmonitor.UpdateProbeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *networkmonitor.UpdateProbeInput, ...func(*networkmonitor.Options)) *networkmonitor.UpdateProbeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmonitor.UpdateProbeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *networkmonitor.UpdateProbeInput, ...func(*networkmonitor.Options)) error); ok {
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
