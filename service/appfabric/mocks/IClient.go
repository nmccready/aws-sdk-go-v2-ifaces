// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	appfabric "github.com/aws/aws-sdk-go-v2/service/appfabric"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchGetUserAccessTasks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetUserAccessTasks(ctx context.Context, params *appfabric.BatchGetUserAccessTasksInput, optFns ...func(*appfabric.Options)) (*appfabric.BatchGetUserAccessTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetUserAccessTasks")
	}

	var r0 *appfabric.BatchGetUserAccessTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.BatchGetUserAccessTasksInput, ...func(*appfabric.Options)) (*appfabric.BatchGetUserAccessTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.BatchGetUserAccessTasksInput, ...func(*appfabric.Options)) *appfabric.BatchGetUserAccessTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.BatchGetUserAccessTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.BatchGetUserAccessTasksInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectAppAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ConnectAppAuthorization(ctx context.Context, params *appfabric.ConnectAppAuthorizationInput, optFns ...func(*appfabric.Options)) (*appfabric.ConnectAppAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ConnectAppAuthorization")
	}

	var r0 *appfabric.ConnectAppAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ConnectAppAuthorizationInput, ...func(*appfabric.Options)) (*appfabric.ConnectAppAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ConnectAppAuthorizationInput, ...func(*appfabric.Options)) *appfabric.ConnectAppAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ConnectAppAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ConnectAppAuthorizationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAppAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAppAuthorization(ctx context.Context, params *appfabric.CreateAppAuthorizationInput, optFns ...func(*appfabric.Options)) (*appfabric.CreateAppAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAppAuthorization")
	}

	var r0 *appfabric.CreateAppAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateAppAuthorizationInput, ...func(*appfabric.Options)) (*appfabric.CreateAppAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateAppAuthorizationInput, ...func(*appfabric.Options)) *appfabric.CreateAppAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.CreateAppAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.CreateAppAuthorizationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAppBundle provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAppBundle(ctx context.Context, params *appfabric.CreateAppBundleInput, optFns ...func(*appfabric.Options)) (*appfabric.CreateAppBundleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAppBundle")
	}

	var r0 *appfabric.CreateAppBundleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateAppBundleInput, ...func(*appfabric.Options)) (*appfabric.CreateAppBundleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateAppBundleInput, ...func(*appfabric.Options)) *appfabric.CreateAppBundleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.CreateAppBundleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.CreateAppBundleInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIngestion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateIngestion(ctx context.Context, params *appfabric.CreateIngestionInput, optFns ...func(*appfabric.Options)) (*appfabric.CreateIngestionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateIngestion")
	}

	var r0 *appfabric.CreateIngestionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateIngestionInput, ...func(*appfabric.Options)) (*appfabric.CreateIngestionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateIngestionInput, ...func(*appfabric.Options)) *appfabric.CreateIngestionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.CreateIngestionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.CreateIngestionInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIngestionDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateIngestionDestination(ctx context.Context, params *appfabric.CreateIngestionDestinationInput, optFns ...func(*appfabric.Options)) (*appfabric.CreateIngestionDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateIngestionDestination")
	}

	var r0 *appfabric.CreateIngestionDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateIngestionDestinationInput, ...func(*appfabric.Options)) (*appfabric.CreateIngestionDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.CreateIngestionDestinationInput, ...func(*appfabric.Options)) *appfabric.CreateIngestionDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.CreateIngestionDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.CreateIngestionDestinationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAppAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAppAuthorization(ctx context.Context, params *appfabric.DeleteAppAuthorizationInput, optFns ...func(*appfabric.Options)) (*appfabric.DeleteAppAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAppAuthorization")
	}

	var r0 *appfabric.DeleteAppAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteAppAuthorizationInput, ...func(*appfabric.Options)) (*appfabric.DeleteAppAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteAppAuthorizationInput, ...func(*appfabric.Options)) *appfabric.DeleteAppAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.DeleteAppAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.DeleteAppAuthorizationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAppBundle provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAppBundle(ctx context.Context, params *appfabric.DeleteAppBundleInput, optFns ...func(*appfabric.Options)) (*appfabric.DeleteAppBundleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAppBundle")
	}

	var r0 *appfabric.DeleteAppBundleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteAppBundleInput, ...func(*appfabric.Options)) (*appfabric.DeleteAppBundleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteAppBundleInput, ...func(*appfabric.Options)) *appfabric.DeleteAppBundleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.DeleteAppBundleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.DeleteAppBundleInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIngestion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteIngestion(ctx context.Context, params *appfabric.DeleteIngestionInput, optFns ...func(*appfabric.Options)) (*appfabric.DeleteIngestionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIngestion")
	}

	var r0 *appfabric.DeleteIngestionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteIngestionInput, ...func(*appfabric.Options)) (*appfabric.DeleteIngestionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteIngestionInput, ...func(*appfabric.Options)) *appfabric.DeleteIngestionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.DeleteIngestionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.DeleteIngestionInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIngestionDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteIngestionDestination(ctx context.Context, params *appfabric.DeleteIngestionDestinationInput, optFns ...func(*appfabric.Options)) (*appfabric.DeleteIngestionDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIngestionDestination")
	}

	var r0 *appfabric.DeleteIngestionDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteIngestionDestinationInput, ...func(*appfabric.Options)) (*appfabric.DeleteIngestionDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.DeleteIngestionDestinationInput, ...func(*appfabric.Options)) *appfabric.DeleteIngestionDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.DeleteIngestionDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.DeleteIngestionDestinationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAppAuthorization(ctx context.Context, params *appfabric.GetAppAuthorizationInput, optFns ...func(*appfabric.Options)) (*appfabric.GetAppAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAppAuthorization")
	}

	var r0 *appfabric.GetAppAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetAppAuthorizationInput, ...func(*appfabric.Options)) (*appfabric.GetAppAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetAppAuthorizationInput, ...func(*appfabric.Options)) *appfabric.GetAppAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.GetAppAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.GetAppAuthorizationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppBundle provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAppBundle(ctx context.Context, params *appfabric.GetAppBundleInput, optFns ...func(*appfabric.Options)) (*appfabric.GetAppBundleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAppBundle")
	}

	var r0 *appfabric.GetAppBundleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetAppBundleInput, ...func(*appfabric.Options)) (*appfabric.GetAppBundleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetAppBundleInput, ...func(*appfabric.Options)) *appfabric.GetAppBundleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.GetAppBundleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.GetAppBundleInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIngestion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetIngestion(ctx context.Context, params *appfabric.GetIngestionInput, optFns ...func(*appfabric.Options)) (*appfabric.GetIngestionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIngestion")
	}

	var r0 *appfabric.GetIngestionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetIngestionInput, ...func(*appfabric.Options)) (*appfabric.GetIngestionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetIngestionInput, ...func(*appfabric.Options)) *appfabric.GetIngestionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.GetIngestionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.GetIngestionInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIngestionDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetIngestionDestination(ctx context.Context, params *appfabric.GetIngestionDestinationInput, optFns ...func(*appfabric.Options)) (*appfabric.GetIngestionDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIngestionDestination")
	}

	var r0 *appfabric.GetIngestionDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetIngestionDestinationInput, ...func(*appfabric.Options)) (*appfabric.GetIngestionDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.GetIngestionDestinationInput, ...func(*appfabric.Options)) *appfabric.GetIngestionDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.GetIngestionDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.GetIngestionDestinationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAppAuthorizations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAppAuthorizations(ctx context.Context, params *appfabric.ListAppAuthorizationsInput, optFns ...func(*appfabric.Options)) (*appfabric.ListAppAuthorizationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAppAuthorizations")
	}

	var r0 *appfabric.ListAppAuthorizationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListAppAuthorizationsInput, ...func(*appfabric.Options)) (*appfabric.ListAppAuthorizationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListAppAuthorizationsInput, ...func(*appfabric.Options)) *appfabric.ListAppAuthorizationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ListAppAuthorizationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ListAppAuthorizationsInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAppBundles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAppBundles(ctx context.Context, params *appfabric.ListAppBundlesInput, optFns ...func(*appfabric.Options)) (*appfabric.ListAppBundlesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAppBundles")
	}

	var r0 *appfabric.ListAppBundlesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListAppBundlesInput, ...func(*appfabric.Options)) (*appfabric.ListAppBundlesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListAppBundlesInput, ...func(*appfabric.Options)) *appfabric.ListAppBundlesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ListAppBundlesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ListAppBundlesInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIngestionDestinations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListIngestionDestinations(ctx context.Context, params *appfabric.ListIngestionDestinationsInput, optFns ...func(*appfabric.Options)) (*appfabric.ListIngestionDestinationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIngestionDestinations")
	}

	var r0 *appfabric.ListIngestionDestinationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListIngestionDestinationsInput, ...func(*appfabric.Options)) (*appfabric.ListIngestionDestinationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListIngestionDestinationsInput, ...func(*appfabric.Options)) *appfabric.ListIngestionDestinationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ListIngestionDestinationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ListIngestionDestinationsInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIngestions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListIngestions(ctx context.Context, params *appfabric.ListIngestionsInput, optFns ...func(*appfabric.Options)) (*appfabric.ListIngestionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIngestions")
	}

	var r0 *appfabric.ListIngestionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListIngestionsInput, ...func(*appfabric.Options)) (*appfabric.ListIngestionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListIngestionsInput, ...func(*appfabric.Options)) *appfabric.ListIngestionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ListIngestionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ListIngestionsInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *appfabric.ListTagsForResourceInput, optFns ...func(*appfabric.Options)) (*appfabric.ListTagsForResourceOutput, error) {
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

	var r0 *appfabric.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListTagsForResourceInput, ...func(*appfabric.Options)) (*appfabric.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.ListTagsForResourceInput, ...func(*appfabric.Options)) *appfabric.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.ListTagsForResourceInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() appfabric.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 appfabric.Options
	if rf, ok := ret.Get(0).(func() appfabric.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(appfabric.Options)
	}

	return r0
}

// StartIngestion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartIngestion(ctx context.Context, params *appfabric.StartIngestionInput, optFns ...func(*appfabric.Options)) (*appfabric.StartIngestionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartIngestion")
	}

	var r0 *appfabric.StartIngestionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StartIngestionInput, ...func(*appfabric.Options)) (*appfabric.StartIngestionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StartIngestionInput, ...func(*appfabric.Options)) *appfabric.StartIngestionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.StartIngestionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.StartIngestionInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartUserAccessTasks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartUserAccessTasks(ctx context.Context, params *appfabric.StartUserAccessTasksInput, optFns ...func(*appfabric.Options)) (*appfabric.StartUserAccessTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartUserAccessTasks")
	}

	var r0 *appfabric.StartUserAccessTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StartUserAccessTasksInput, ...func(*appfabric.Options)) (*appfabric.StartUserAccessTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StartUserAccessTasksInput, ...func(*appfabric.Options)) *appfabric.StartUserAccessTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.StartUserAccessTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.StartUserAccessTasksInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopIngestion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopIngestion(ctx context.Context, params *appfabric.StopIngestionInput, optFns ...func(*appfabric.Options)) (*appfabric.StopIngestionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopIngestion")
	}

	var r0 *appfabric.StopIngestionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StopIngestionInput, ...func(*appfabric.Options)) (*appfabric.StopIngestionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.StopIngestionInput, ...func(*appfabric.Options)) *appfabric.StopIngestionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.StopIngestionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.StopIngestionInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *appfabric.TagResourceInput, optFns ...func(*appfabric.Options)) (*appfabric.TagResourceOutput, error) {
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

	var r0 *appfabric.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.TagResourceInput, ...func(*appfabric.Options)) (*appfabric.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.TagResourceInput, ...func(*appfabric.Options)) *appfabric.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.TagResourceInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *appfabric.UntagResourceInput, optFns ...func(*appfabric.Options)) (*appfabric.UntagResourceOutput, error) {
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

	var r0 *appfabric.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UntagResourceInput, ...func(*appfabric.Options)) (*appfabric.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UntagResourceInput, ...func(*appfabric.Options)) *appfabric.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.UntagResourceInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAppAuthorization provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAppAuthorization(ctx context.Context, params *appfabric.UpdateAppAuthorizationInput, optFns ...func(*appfabric.Options)) (*appfabric.UpdateAppAuthorizationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAppAuthorization")
	}

	var r0 *appfabric.UpdateAppAuthorizationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UpdateAppAuthorizationInput, ...func(*appfabric.Options)) (*appfabric.UpdateAppAuthorizationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UpdateAppAuthorizationInput, ...func(*appfabric.Options)) *appfabric.UpdateAppAuthorizationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.UpdateAppAuthorizationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.UpdateAppAuthorizationInput, ...func(*appfabric.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIngestionDestination provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateIngestionDestination(ctx context.Context, params *appfabric.UpdateIngestionDestinationInput, optFns ...func(*appfabric.Options)) (*appfabric.UpdateIngestionDestinationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIngestionDestination")
	}

	var r0 *appfabric.UpdateIngestionDestinationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UpdateIngestionDestinationInput, ...func(*appfabric.Options)) (*appfabric.UpdateIngestionDestinationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appfabric.UpdateIngestionDestinationInput, ...func(*appfabric.Options)) *appfabric.UpdateIngestionDestinationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appfabric.UpdateIngestionDestinationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appfabric.UpdateIngestionDestinationInput, ...func(*appfabric.Options)) error); ok {
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
