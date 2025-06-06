// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	workspacesthinclient "github.com/aws/aws-sdk-go-v2/service/workspacesthinclient"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateEnvironment(ctx context.Context, params *workspacesthinclient.CreateEnvironmentInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.CreateEnvironmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateEnvironment")
	}

	var r0 *workspacesthinclient.CreateEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.CreateEnvironmentInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.CreateEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.CreateEnvironmentInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.CreateEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.CreateEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.CreateEnvironmentInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDevice(ctx context.Context, params *workspacesthinclient.DeleteDeviceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeleteDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDevice")
	}

	var r0 *workspacesthinclient.DeleteDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeleteDeviceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeleteDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeleteDeviceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.DeleteDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.DeleteDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.DeleteDeviceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteEnvironment(ctx context.Context, params *workspacesthinclient.DeleteEnvironmentInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeleteEnvironmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEnvironment")
	}

	var r0 *workspacesthinclient.DeleteEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeleteEnvironmentInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeleteEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeleteEnvironmentInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.DeleteEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.DeleteEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.DeleteEnvironmentInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeregisterDevice(ctx context.Context, params *workspacesthinclient.DeregisterDeviceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeregisterDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterDevice")
	}

	var r0 *workspacesthinclient.DeregisterDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeregisterDeviceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.DeregisterDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.DeregisterDeviceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.DeregisterDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.DeregisterDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.DeregisterDeviceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDevice(ctx context.Context, params *workspacesthinclient.GetDeviceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDevice")
	}

	var r0 *workspacesthinclient.GetDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetDeviceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetDeviceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.GetDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.GetDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.GetDeviceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetEnvironment(ctx context.Context, params *workspacesthinclient.GetEnvironmentInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetEnvironmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEnvironment")
	}

	var r0 *workspacesthinclient.GetEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetEnvironmentInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetEnvironmentInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.GetEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.GetEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.GetEnvironmentInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSoftwareSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSoftwareSet(ctx context.Context, params *workspacesthinclient.GetSoftwareSetInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetSoftwareSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSoftwareSet")
	}

	var r0 *workspacesthinclient.GetSoftwareSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetSoftwareSetInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.GetSoftwareSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.GetSoftwareSetInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.GetSoftwareSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.GetSoftwareSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.GetSoftwareSetInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDevices provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDevices(ctx context.Context, params *workspacesthinclient.ListDevicesInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListDevicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 *workspacesthinclient.ListDevicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListDevicesInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListDevicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListDevicesInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.ListDevicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.ListDevicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.ListDevicesInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnvironments provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListEnvironments(ctx context.Context, params *workspacesthinclient.ListEnvironmentsInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListEnvironmentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEnvironments")
	}

	var r0 *workspacesthinclient.ListEnvironmentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListEnvironmentsInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListEnvironmentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListEnvironmentsInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.ListEnvironmentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.ListEnvironmentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.ListEnvironmentsInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSoftwareSets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSoftwareSets(ctx context.Context, params *workspacesthinclient.ListSoftwareSetsInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListSoftwareSetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSoftwareSets")
	}

	var r0 *workspacesthinclient.ListSoftwareSetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListSoftwareSetsInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListSoftwareSetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListSoftwareSetsInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.ListSoftwareSetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.ListSoftwareSetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.ListSoftwareSetsInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *workspacesthinclient.ListTagsForResourceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListTagsForResourceOutput, error) {
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

	var r0 *workspacesthinclient.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListTagsForResourceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.ListTagsForResourceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.ListTagsForResourceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() workspacesthinclient.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 workspacesthinclient.Options
	if rf, ok := ret.Get(0).(func() workspacesthinclient.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(workspacesthinclient.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *workspacesthinclient.TagResourceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.TagResourceOutput, error) {
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

	var r0 *workspacesthinclient.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.TagResourceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.TagResourceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.TagResourceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *workspacesthinclient.UntagResourceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UntagResourceOutput, error) {
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

	var r0 *workspacesthinclient.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UntagResourceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UntagResourceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.UntagResourceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDevice provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateDevice(ctx context.Context, params *workspacesthinclient.UpdateDeviceInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDevice")
	}

	var r0 *workspacesthinclient.UpdateDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateDeviceInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateDeviceInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.UpdateDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.UpdateDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.UpdateDeviceInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateEnvironment(ctx context.Context, params *workspacesthinclient.UpdateEnvironmentInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateEnvironmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEnvironment")
	}

	var r0 *workspacesthinclient.UpdateEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateEnvironmentInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateEnvironmentInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.UpdateEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.UpdateEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.UpdateEnvironmentInput, ...func(*workspacesthinclient.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSoftwareSet provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSoftwareSet(ctx context.Context, params *workspacesthinclient.UpdateSoftwareSetInput, optFns ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateSoftwareSetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSoftwareSet")
	}

	var r0 *workspacesthinclient.UpdateSoftwareSetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateSoftwareSetInput, ...func(*workspacesthinclient.Options)) (*workspacesthinclient.UpdateSoftwareSetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *workspacesthinclient.UpdateSoftwareSetInput, ...func(*workspacesthinclient.Options)) *workspacesthinclient.UpdateSoftwareSetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspacesthinclient.UpdateSoftwareSetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *workspacesthinclient.UpdateSoftwareSetInput, ...func(*workspacesthinclient.Options)) error); ok {
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
