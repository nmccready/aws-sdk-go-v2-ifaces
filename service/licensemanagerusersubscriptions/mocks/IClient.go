// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	licensemanagerusersubscriptions "github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateUser(ctx context.Context, params *licensemanagerusersubscriptions.AssociateUserInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.AssociateUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateUser")
	}

	var r0 *licensemanagerusersubscriptions.AssociateUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.AssociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.AssociateUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.AssociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.AssociateUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.AssociateUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.AssociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLicenseServerEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateLicenseServerEndpoint(ctx context.Context, params *licensemanagerusersubscriptions.CreateLicenseServerEndpointInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.CreateLicenseServerEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLicenseServerEndpoint")
	}

	var r0 *licensemanagerusersubscriptions.CreateLicenseServerEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.CreateLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.CreateLicenseServerEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.CreateLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.CreateLicenseServerEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.CreateLicenseServerEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.CreateLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLicenseServerEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteLicenseServerEndpoint(ctx context.Context, params *licensemanagerusersubscriptions.DeleteLicenseServerEndpointInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DeleteLicenseServerEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLicenseServerEndpoint")
	}

	var r0 *licensemanagerusersubscriptions.DeleteLicenseServerEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DeleteLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DeleteLicenseServerEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DeleteLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.DeleteLicenseServerEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.DeleteLicenseServerEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.DeleteLicenseServerEndpointInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterIdentityProvider provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeregisterIdentityProvider(ctx context.Context, params *licensemanagerusersubscriptions.DeregisterIdentityProviderInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DeregisterIdentityProviderOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterIdentityProvider")
	}

	var r0 *licensemanagerusersubscriptions.DeregisterIdentityProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DeregisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DeregisterIdentityProviderOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DeregisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.DeregisterIdentityProviderOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.DeregisterIdentityProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.DeregisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateUser(ctx context.Context, params *licensemanagerusersubscriptions.DisassociateUserInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DisassociateUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateUser")
	}

	var r0 *licensemanagerusersubscriptions.DisassociateUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DisassociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.DisassociateUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.DisassociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.DisassociateUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.DisassociateUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.DisassociateUserInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentityProviders provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListIdentityProviders(ctx context.Context, params *licensemanagerusersubscriptions.ListIdentityProvidersInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdentityProviders")
	}

	var r0 *licensemanagerusersubscriptions.ListIdentityProvidersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListIdentityProvidersInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListIdentityProvidersInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListIdentityProvidersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListIdentityProvidersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListIdentityProvidersInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInstances provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListInstances(ctx context.Context, params *licensemanagerusersubscriptions.ListInstancesInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListInstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListInstances")
	}

	var r0 *licensemanagerusersubscriptions.ListInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListInstancesInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListInstancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListInstancesInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListInstancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListInstancesInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLicenseServerEndpoints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListLicenseServerEndpoints(ctx context.Context, params *licensemanagerusersubscriptions.ListLicenseServerEndpointsInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLicenseServerEndpoints")
	}

	var r0 *licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListLicenseServerEndpointsInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListLicenseServerEndpointsInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListLicenseServerEndpointsInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProductSubscriptions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProductSubscriptions(ctx context.Context, params *licensemanagerusersubscriptions.ListProductSubscriptionsInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProductSubscriptions")
	}

	var r0 *licensemanagerusersubscriptions.ListProductSubscriptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListProductSubscriptionsInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListProductSubscriptionsInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListProductSubscriptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListProductSubscriptionsInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *licensemanagerusersubscriptions.ListTagsForResourceInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListTagsForResourceOutput, error) {
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

	var r0 *licensemanagerusersubscriptions.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListTagsForResourceInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListTagsForResourceInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListTagsForResourceInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserAssociations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListUserAssociations(ctx context.Context, params *licensemanagerusersubscriptions.ListUserAssociationsInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListUserAssociations")
	}

	var r0 *licensemanagerusersubscriptions.ListUserAssociationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListUserAssociationsInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.ListUserAssociationsInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.ListUserAssociationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.ListUserAssociationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.ListUserAssociationsInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() licensemanagerusersubscriptions.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 licensemanagerusersubscriptions.Options
	if rf, ok := ret.Get(0).(func() licensemanagerusersubscriptions.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(licensemanagerusersubscriptions.Options)
	}

	return r0
}

// RegisterIdentityProvider provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RegisterIdentityProvider(ctx context.Context, params *licensemanagerusersubscriptions.RegisterIdentityProviderInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.RegisterIdentityProviderOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterIdentityProvider")
	}

	var r0 *licensemanagerusersubscriptions.RegisterIdentityProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.RegisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.RegisterIdentityProviderOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.RegisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.RegisterIdentityProviderOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.RegisterIdentityProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.RegisterIdentityProviderInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartProductSubscription provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartProductSubscription(ctx context.Context, params *licensemanagerusersubscriptions.StartProductSubscriptionInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.StartProductSubscriptionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartProductSubscription")
	}

	var r0 *licensemanagerusersubscriptions.StartProductSubscriptionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.StartProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.StartProductSubscriptionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.StartProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.StartProductSubscriptionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.StartProductSubscriptionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.StartProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopProductSubscription provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopProductSubscription(ctx context.Context, params *licensemanagerusersubscriptions.StopProductSubscriptionInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.StopProductSubscriptionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopProductSubscription")
	}

	var r0 *licensemanagerusersubscriptions.StopProductSubscriptionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.StopProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.StopProductSubscriptionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.StopProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.StopProductSubscriptionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.StopProductSubscriptionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.StopProductSubscriptionInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *licensemanagerusersubscriptions.TagResourceInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.TagResourceOutput, error) {
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

	var r0 *licensemanagerusersubscriptions.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.TagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.TagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.TagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *licensemanagerusersubscriptions.UntagResourceInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.UntagResourceOutput, error) {
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

	var r0 *licensemanagerusersubscriptions.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.UntagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.UntagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.UntagResourceInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentityProviderSettings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateIdentityProviderSettings(ctx context.Context, params *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput, optFns ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIdentityProviderSettings")
	}

	var r0 *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput, ...func(*licensemanagerusersubscriptions.Options)) (*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput, ...func(*licensemanagerusersubscriptions.Options)) *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput, ...func(*licensemanagerusersubscriptions.Options)) error); ok {
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
