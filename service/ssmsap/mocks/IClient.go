// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	ssmsap "github.com/aws/aws-sdk-go-v2/service/ssmsap"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DeleteResourcePermission provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteResourcePermission(ctx context.Context, params *ssmsap.DeleteResourcePermissionInput, optFns ...func(*ssmsap.Options)) (*ssmsap.DeleteResourcePermissionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResourcePermission")
	}

	var r0 *ssmsap.DeleteResourcePermissionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.DeleteResourcePermissionInput, ...func(*ssmsap.Options)) (*ssmsap.DeleteResourcePermissionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.DeleteResourcePermissionInput, ...func(*ssmsap.Options)) *ssmsap.DeleteResourcePermissionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.DeleteResourcePermissionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.DeleteResourcePermissionInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeregisterApplication(ctx context.Context, params *ssmsap.DeregisterApplicationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.DeregisterApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterApplication")
	}

	var r0 *ssmsap.DeregisterApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.DeregisterApplicationInput, ...func(*ssmsap.Options)) (*ssmsap.DeregisterApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.DeregisterApplicationInput, ...func(*ssmsap.Options)) *ssmsap.DeregisterApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.DeregisterApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.DeregisterApplicationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetApplication(ctx context.Context, params *ssmsap.GetApplicationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.GetApplicationOutput, error) {
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

	var r0 *ssmsap.GetApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetApplicationInput, ...func(*ssmsap.Options)) (*ssmsap.GetApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetApplicationInput, ...func(*ssmsap.Options)) *ssmsap.GetApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.GetApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.GetApplicationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetComponent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetComponent(ctx context.Context, params *ssmsap.GetComponentInput, optFns ...func(*ssmsap.Options)) (*ssmsap.GetComponentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetComponent")
	}

	var r0 *ssmsap.GetComponentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetComponentInput, ...func(*ssmsap.Options)) (*ssmsap.GetComponentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetComponentInput, ...func(*ssmsap.Options)) *ssmsap.GetComponentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.GetComponentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.GetComponentInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabase provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDatabase(ctx context.Context, params *ssmsap.GetDatabaseInput, optFns ...func(*ssmsap.Options)) (*ssmsap.GetDatabaseOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabase")
	}

	var r0 *ssmsap.GetDatabaseOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetDatabaseInput, ...func(*ssmsap.Options)) (*ssmsap.GetDatabaseOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetDatabaseInput, ...func(*ssmsap.Options)) *ssmsap.GetDatabaseOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.GetDatabaseOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.GetDatabaseInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetOperation(ctx context.Context, params *ssmsap.GetOperationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.GetOperationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOperation")
	}

	var r0 *ssmsap.GetOperationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetOperationInput, ...func(*ssmsap.Options)) (*ssmsap.GetOperationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetOperationInput, ...func(*ssmsap.Options)) *ssmsap.GetOperationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.GetOperationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.GetOperationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePermission provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourcePermission(ctx context.Context, params *ssmsap.GetResourcePermissionInput, optFns ...func(*ssmsap.Options)) (*ssmsap.GetResourcePermissionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcePermission")
	}

	var r0 *ssmsap.GetResourcePermissionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetResourcePermissionInput, ...func(*ssmsap.Options)) (*ssmsap.GetResourcePermissionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.GetResourcePermissionInput, ...func(*ssmsap.Options)) *ssmsap.GetResourcePermissionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.GetResourcePermissionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.GetResourcePermissionInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplications provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApplications(ctx context.Context, params *ssmsap.ListApplicationsInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListApplicationsOutput, error) {
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

	var r0 *ssmsap.ListApplicationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListApplicationsInput, ...func(*ssmsap.Options)) (*ssmsap.ListApplicationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListApplicationsInput, ...func(*ssmsap.Options)) *ssmsap.ListApplicationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListApplicationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListApplicationsInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListComponents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListComponents(ctx context.Context, params *ssmsap.ListComponentsInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListComponentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListComponents")
	}

	var r0 *ssmsap.ListComponentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListComponentsInput, ...func(*ssmsap.Options)) (*ssmsap.ListComponentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListComponentsInput, ...func(*ssmsap.Options)) *ssmsap.ListComponentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListComponentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListComponentsInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDatabases provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDatabases(ctx context.Context, params *ssmsap.ListDatabasesInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListDatabasesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabases")
	}

	var r0 *ssmsap.ListDatabasesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListDatabasesInput, ...func(*ssmsap.Options)) (*ssmsap.ListDatabasesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListDatabasesInput, ...func(*ssmsap.Options)) *ssmsap.ListDatabasesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListDatabasesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListDatabasesInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOperationEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListOperationEvents(ctx context.Context, params *ssmsap.ListOperationEventsInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListOperationEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListOperationEvents")
	}

	var r0 *ssmsap.ListOperationEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListOperationEventsInput, ...func(*ssmsap.Options)) (*ssmsap.ListOperationEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListOperationEventsInput, ...func(*ssmsap.Options)) *ssmsap.ListOperationEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListOperationEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListOperationEventsInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOperations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListOperations(ctx context.Context, params *ssmsap.ListOperationsInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListOperationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListOperations")
	}

	var r0 *ssmsap.ListOperationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListOperationsInput, ...func(*ssmsap.Options)) (*ssmsap.ListOperationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListOperationsInput, ...func(*ssmsap.Options)) *ssmsap.ListOperationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListOperationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListOperationsInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *ssmsap.ListTagsForResourceInput, optFns ...func(*ssmsap.Options)) (*ssmsap.ListTagsForResourceOutput, error) {
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

	var r0 *ssmsap.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListTagsForResourceInput, ...func(*ssmsap.Options)) (*ssmsap.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.ListTagsForResourceInput, ...func(*ssmsap.Options)) *ssmsap.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.ListTagsForResourceInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() ssmsap.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 ssmsap.Options
	if rf, ok := ret.Get(0).(func() ssmsap.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ssmsap.Options)
	}

	return r0
}

// PutResourcePermission provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutResourcePermission(ctx context.Context, params *ssmsap.PutResourcePermissionInput, optFns ...func(*ssmsap.Options)) (*ssmsap.PutResourcePermissionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutResourcePermission")
	}

	var r0 *ssmsap.PutResourcePermissionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.PutResourcePermissionInput, ...func(*ssmsap.Options)) (*ssmsap.PutResourcePermissionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.PutResourcePermissionInput, ...func(*ssmsap.Options)) *ssmsap.PutResourcePermissionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.PutResourcePermissionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.PutResourcePermissionInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RegisterApplication(ctx context.Context, params *ssmsap.RegisterApplicationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.RegisterApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterApplication")
	}

	var r0 *ssmsap.RegisterApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.RegisterApplicationInput, ...func(*ssmsap.Options)) (*ssmsap.RegisterApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.RegisterApplicationInput, ...func(*ssmsap.Options)) *ssmsap.RegisterApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.RegisterApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.RegisterApplicationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartApplication(ctx context.Context, params *ssmsap.StartApplicationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.StartApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartApplication")
	}

	var r0 *ssmsap.StartApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StartApplicationInput, ...func(*ssmsap.Options)) (*ssmsap.StartApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StartApplicationInput, ...func(*ssmsap.Options)) *ssmsap.StartApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.StartApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.StartApplicationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartApplicationRefresh provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartApplicationRefresh(ctx context.Context, params *ssmsap.StartApplicationRefreshInput, optFns ...func(*ssmsap.Options)) (*ssmsap.StartApplicationRefreshOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartApplicationRefresh")
	}

	var r0 *ssmsap.StartApplicationRefreshOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StartApplicationRefreshInput, ...func(*ssmsap.Options)) (*ssmsap.StartApplicationRefreshOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StartApplicationRefreshInput, ...func(*ssmsap.Options)) *ssmsap.StartApplicationRefreshOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.StartApplicationRefreshOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.StartApplicationRefreshInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopApplication(ctx context.Context, params *ssmsap.StopApplicationInput, optFns ...func(*ssmsap.Options)) (*ssmsap.StopApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopApplication")
	}

	var r0 *ssmsap.StopApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StopApplicationInput, ...func(*ssmsap.Options)) (*ssmsap.StopApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.StopApplicationInput, ...func(*ssmsap.Options)) *ssmsap.StopApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.StopApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.StopApplicationInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *ssmsap.TagResourceInput, optFns ...func(*ssmsap.Options)) (*ssmsap.TagResourceOutput, error) {
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

	var r0 *ssmsap.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.TagResourceInput, ...func(*ssmsap.Options)) (*ssmsap.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.TagResourceInput, ...func(*ssmsap.Options)) *ssmsap.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.TagResourceInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *ssmsap.UntagResourceInput, optFns ...func(*ssmsap.Options)) (*ssmsap.UntagResourceOutput, error) {
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

	var r0 *ssmsap.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.UntagResourceInput, ...func(*ssmsap.Options)) (*ssmsap.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.UntagResourceInput, ...func(*ssmsap.Options)) *ssmsap.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.UntagResourceInput, ...func(*ssmsap.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplicationSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateApplicationSettings(ctx context.Context, params *ssmsap.UpdateApplicationSettingsInput, optFns ...func(*ssmsap.Options)) (*ssmsap.UpdateApplicationSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateApplicationSettings")
	}

	var r0 *ssmsap.UpdateApplicationSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.UpdateApplicationSettingsInput, ...func(*ssmsap.Options)) (*ssmsap.UpdateApplicationSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ssmsap.UpdateApplicationSettingsInput, ...func(*ssmsap.Options)) *ssmsap.UpdateApplicationSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssmsap.UpdateApplicationSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ssmsap.UpdateApplicationSettingsInput, ...func(*ssmsap.Options)) error); ok {
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
