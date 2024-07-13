// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	migrationhubrefactorspaces "github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateApplication(ctx context.Context, params *migrationhubrefactorspaces.CreateApplicationInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateApplication")
	}

	var r0 *migrationhubrefactorspaces.CreateApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateApplicationInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateApplicationInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.CreateApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.CreateApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.CreateApplicationInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateEnvironment(ctx context.Context, params *migrationhubrefactorspaces.CreateEnvironmentInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateEnvironmentOutput, error) {
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

	var r0 *migrationhubrefactorspaces.CreateEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.CreateEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.CreateEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.CreateEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoute provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRoute(ctx context.Context, params *migrationhubrefactorspaces.CreateRouteInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateRouteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRoute")
	}

	var r0 *migrationhubrefactorspaces.CreateRouteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateRouteInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateRouteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateRouteInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.CreateRouteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.CreateRouteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.CreateRouteInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateService provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateService(ctx context.Context, params *migrationhubrefactorspaces.CreateServiceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateServiceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateService")
	}

	var r0 *migrationhubrefactorspaces.CreateServiceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateServiceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.CreateServiceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.CreateServiceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.CreateServiceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.CreateServiceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.CreateServiceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteApplication(ctx context.Context, params *migrationhubrefactorspaces.DeleteApplicationInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteApplication")
	}

	var r0 *migrationhubrefactorspaces.DeleteApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteApplicationInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteApplicationInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.DeleteApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.DeleteApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.DeleteApplicationInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteEnvironment(ctx context.Context, params *migrationhubrefactorspaces.DeleteEnvironmentInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteEnvironmentOutput, error) {
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

	var r0 *migrationhubrefactorspaces.DeleteEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.DeleteEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.DeleteEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.DeleteEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteResourcePolicy(ctx context.Context, params *migrationhubrefactorspaces.DeleteResourcePolicyInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResourcePolicy")
	}

	var r0 *migrationhubrefactorspaces.DeleteResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.DeleteResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.DeleteResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.DeleteResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoute provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRoute(ctx context.Context, params *migrationhubrefactorspaces.DeleteRouteInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteRouteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRoute")
	}

	var r0 *migrationhubrefactorspaces.DeleteRouteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteRouteInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteRouteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteRouteInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.DeleteRouteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.DeleteRouteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.DeleteRouteInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteService provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteService(ctx context.Context, params *migrationhubrefactorspaces.DeleteServiceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteServiceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteService")
	}

	var r0 *migrationhubrefactorspaces.DeleteServiceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteServiceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.DeleteServiceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.DeleteServiceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.DeleteServiceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.DeleteServiceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.DeleteServiceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetApplication(ctx context.Context, params *migrationhubrefactorspaces.GetApplicationInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetApplication")
	}

	var r0 *migrationhubrefactorspaces.GetApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetApplicationInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetApplicationInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.GetApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.GetApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.GetApplicationInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnvironment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetEnvironment(ctx context.Context, params *migrationhubrefactorspaces.GetEnvironmentInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetEnvironmentOutput, error) {
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

	var r0 *migrationhubrefactorspaces.GetEnvironmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetEnvironmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.GetEnvironmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.GetEnvironmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.GetEnvironmentInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourcePolicy(ctx context.Context, params *migrationhubrefactorspaces.GetResourcePolicyInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcePolicy")
	}

	var r0 *migrationhubrefactorspaces.GetResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.GetResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.GetResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.GetResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoute provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRoute(ctx context.Context, params *migrationhubrefactorspaces.GetRouteInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetRouteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRoute")
	}

	var r0 *migrationhubrefactorspaces.GetRouteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetRouteInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetRouteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetRouteInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.GetRouteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.GetRouteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.GetRouteInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetService provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetService(ctx context.Context, params *migrationhubrefactorspaces.GetServiceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetServiceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetService")
	}

	var r0 *migrationhubrefactorspaces.GetServiceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetServiceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.GetServiceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.GetServiceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.GetServiceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.GetServiceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.GetServiceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplications provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApplications(ctx context.Context, params *migrationhubrefactorspaces.ListApplicationsInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListApplicationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListApplications")
	}

	var r0 *migrationhubrefactorspaces.ListApplicationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListApplicationsInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListApplicationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListApplicationsInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListApplicationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListApplicationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListApplicationsInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnvironmentVpcs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListEnvironmentVpcs(ctx context.Context, params *migrationhubrefactorspaces.ListEnvironmentVpcsInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEnvironmentVpcs")
	}

	var r0 *migrationhubrefactorspaces.ListEnvironmentVpcsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentVpcsInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentVpcsInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListEnvironmentVpcsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentVpcsInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnvironments provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListEnvironments(ctx context.Context, params *migrationhubrefactorspaces.ListEnvironmentsInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListEnvironmentsOutput, error) {
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

	var r0 *migrationhubrefactorspaces.ListEnvironmentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentsInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListEnvironmentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentsInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListEnvironmentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListEnvironmentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListEnvironmentsInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRoutes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRoutes(ctx context.Context, params *migrationhubrefactorspaces.ListRoutesInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListRoutesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRoutes")
	}

	var r0 *migrationhubrefactorspaces.ListRoutesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListRoutesInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListRoutesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListRoutesInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListRoutesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListRoutesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListRoutesInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServices provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListServices(ctx context.Context, params *migrationhubrefactorspaces.ListServicesInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListServicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListServices")
	}

	var r0 *migrationhubrefactorspaces.ListServicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListServicesInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListServicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListServicesInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListServicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListServicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListServicesInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *migrationhubrefactorspaces.ListTagsForResourceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListTagsForResourceOutput, error) {
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

	var r0 *migrationhubrefactorspaces.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListTagsForResourceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.ListTagsForResourceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.ListTagsForResourceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() migrationhubrefactorspaces.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 migrationhubrefactorspaces.Options
	if rf, ok := ret.Get(0).(func() migrationhubrefactorspaces.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(migrationhubrefactorspaces.Options)
	}

	return r0
}

// PutResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutResourcePolicy(ctx context.Context, params *migrationhubrefactorspaces.PutResourcePolicyInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.PutResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutResourcePolicy")
	}

	var r0 *migrationhubrefactorspaces.PutResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.PutResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.PutResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.PutResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.PutResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.PutResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.PutResourcePolicyInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *migrationhubrefactorspaces.TagResourceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.TagResourceOutput, error) {
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

	var r0 *migrationhubrefactorspaces.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.TagResourceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.TagResourceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.TagResourceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *migrationhubrefactorspaces.UntagResourceInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.UntagResourceOutput, error) {
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

	var r0 *migrationhubrefactorspaces.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.UntagResourceInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.UntagResourceInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.UntagResourceInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRoute provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRoute(ctx context.Context, params *migrationhubrefactorspaces.UpdateRouteInput, optFns ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.UpdateRouteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRoute")
	}

	var r0 *migrationhubrefactorspaces.UpdateRouteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.UpdateRouteInput, ...func(*migrationhubrefactorspaces.Options)) (*migrationhubrefactorspaces.UpdateRouteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubrefactorspaces.UpdateRouteInput, ...func(*migrationhubrefactorspaces.Options)) *migrationhubrefactorspaces.UpdateRouteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubrefactorspaces.UpdateRouteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubrefactorspaces.UpdateRouteInput, ...func(*migrationhubrefactorspaces.Options)) error); ok {
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
