// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	sso "github.com/aws/aws-sdk-go-v2/service/sso"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetRoleCredentials provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRoleCredentials(ctx context.Context, params *sso.GetRoleCredentialsInput, optFns ...func(*sso.Options)) (*sso.GetRoleCredentialsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRoleCredentials")
	}

	var r0 *sso.GetRoleCredentialsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sso.GetRoleCredentialsInput, ...func(*sso.Options)) (*sso.GetRoleCredentialsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sso.GetRoleCredentialsInput, ...func(*sso.Options)) *sso.GetRoleCredentialsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sso.GetRoleCredentialsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sso.GetRoleCredentialsInput, ...func(*sso.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAccountRoles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAccountRoles(ctx context.Context, params *sso.ListAccountRolesInput, optFns ...func(*sso.Options)) (*sso.ListAccountRolesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAccountRoles")
	}

	var r0 *sso.ListAccountRolesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sso.ListAccountRolesInput, ...func(*sso.Options)) (*sso.ListAccountRolesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sso.ListAccountRolesInput, ...func(*sso.Options)) *sso.ListAccountRolesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sso.ListAccountRolesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sso.ListAccountRolesInput, ...func(*sso.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAccounts provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAccounts(ctx context.Context, params *sso.ListAccountsInput, optFns ...func(*sso.Options)) (*sso.ListAccountsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAccounts")
	}

	var r0 *sso.ListAccountsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sso.ListAccountsInput, ...func(*sso.Options)) (*sso.ListAccountsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sso.ListAccountsInput, ...func(*sso.Options)) *sso.ListAccountsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sso.ListAccountsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sso.ListAccountsInput, ...func(*sso.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Logout(ctx context.Context, params *sso.LogoutInput, optFns ...func(*sso.Options)) (*sso.LogoutOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 *sso.LogoutOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sso.LogoutInput, ...func(*sso.Options)) (*sso.LogoutOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sso.LogoutInput, ...func(*sso.Options)) *sso.LogoutOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sso.LogoutOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sso.LogoutInput, ...func(*sso.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() sso.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 sso.Options
	if rf, ok := ret.Get(0).(func() sso.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sso.Options)
	}

	return r0
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
