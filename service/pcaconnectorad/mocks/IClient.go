// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	pcaconnectorad "github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateConnector(ctx context.Context, params *pcaconnectorad.CreateConnectorInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateConnector")
	}

	var r0 *pcaconnectorad.CreateConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateConnectorInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateConnectorInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.CreateConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.CreateConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.CreateConnectorInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDirectoryRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDirectoryRegistration(ctx context.Context, params *pcaconnectorad.CreateDirectoryRegistrationInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateDirectoryRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDirectoryRegistration")
	}

	var r0 *pcaconnectorad.CreateDirectoryRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateDirectoryRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.CreateDirectoryRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.CreateDirectoryRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.CreateDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServicePrincipalName provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateServicePrincipalName(ctx context.Context, params *pcaconnectorad.CreateServicePrincipalNameInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateServicePrincipalNameOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateServicePrincipalName")
	}

	var r0 *pcaconnectorad.CreateServicePrincipalNameOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateServicePrincipalNameOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.CreateServicePrincipalNameOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.CreateServicePrincipalNameOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.CreateServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTemplate(ctx context.Context, params *pcaconnectorad.CreateTemplateInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTemplate")
	}

	var r0 *pcaconnectorad.CreateTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateTemplateInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateTemplateInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.CreateTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.CreateTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.CreateTemplateInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTemplateGroupAccessControlEntry provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTemplateGroupAccessControlEntry(ctx context.Context, params *pcaconnectorad.CreateTemplateGroupAccessControlEntryInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTemplateGroupAccessControlEntry")
	}

	var r0 *pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.CreateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.CreateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteConnector(ctx context.Context, params *pcaconnectorad.DeleteConnectorInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnector")
	}

	var r0 *pcaconnectorad.DeleteConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteConnectorInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteConnectorInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.DeleteConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.DeleteConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.DeleteConnectorInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDirectoryRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDirectoryRegistration(ctx context.Context, params *pcaconnectorad.DeleteDirectoryRegistrationInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteDirectoryRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDirectoryRegistration")
	}

	var r0 *pcaconnectorad.DeleteDirectoryRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteDirectoryRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.DeleteDirectoryRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.DeleteDirectoryRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.DeleteDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteServicePrincipalName provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteServicePrincipalName(ctx context.Context, params *pcaconnectorad.DeleteServicePrincipalNameInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteServicePrincipalNameOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteServicePrincipalName")
	}

	var r0 *pcaconnectorad.DeleteServicePrincipalNameOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteServicePrincipalNameOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.DeleteServicePrincipalNameOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.DeleteServicePrincipalNameOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.DeleteServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTemplate(ctx context.Context, params *pcaconnectorad.DeleteTemplateInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTemplate")
	}

	var r0 *pcaconnectorad.DeleteTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteTemplateInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteTemplateInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.DeleteTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.DeleteTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.DeleteTemplateInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTemplateGroupAccessControlEntry provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTemplateGroupAccessControlEntry(ctx context.Context, params *pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTemplateGroupAccessControlEntry")
	}

	var r0 *pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetConnector(ctx context.Context, params *pcaconnectorad.GetConnectorInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetConnectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConnector")
	}

	var r0 *pcaconnectorad.GetConnectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetConnectorInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetConnectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetConnectorInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.GetConnectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.GetConnectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.GetConnectorInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectoryRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDirectoryRegistration(ctx context.Context, params *pcaconnectorad.GetDirectoryRegistrationInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetDirectoryRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDirectoryRegistration")
	}

	var r0 *pcaconnectorad.GetDirectoryRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetDirectoryRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.GetDirectoryRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.GetDirectoryRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.GetDirectoryRegistrationInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServicePrincipalName provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetServicePrincipalName(ctx context.Context, params *pcaconnectorad.GetServicePrincipalNameInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetServicePrincipalNameOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetServicePrincipalName")
	}

	var r0 *pcaconnectorad.GetServicePrincipalNameOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetServicePrincipalNameOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.GetServicePrincipalNameOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.GetServicePrincipalNameOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.GetServicePrincipalNameInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTemplate(ctx context.Context, params *pcaconnectorad.GetTemplateInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplate")
	}

	var r0 *pcaconnectorad.GetTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetTemplateInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetTemplateInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.GetTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.GetTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.GetTemplateInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTemplateGroupAccessControlEntry provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTemplateGroupAccessControlEntry(ctx context.Context, params *pcaconnectorad.GetTemplateGroupAccessControlEntryInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetTemplateGroupAccessControlEntryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplateGroupAccessControlEntry")
	}

	var r0 *pcaconnectorad.GetTemplateGroupAccessControlEntryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.GetTemplateGroupAccessControlEntryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.GetTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.GetTemplateGroupAccessControlEntryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.GetTemplateGroupAccessControlEntryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.GetTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConnectors provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListConnectors(ctx context.Context, params *pcaconnectorad.ListConnectorsInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListConnectorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListConnectors")
	}

	var r0 *pcaconnectorad.ListConnectorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListConnectorsInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListConnectorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListConnectorsInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListConnectorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListConnectorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListConnectorsInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDirectoryRegistrations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDirectoryRegistrations(ctx context.Context, params *pcaconnectorad.ListDirectoryRegistrationsInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListDirectoryRegistrationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDirectoryRegistrations")
	}

	var r0 *pcaconnectorad.ListDirectoryRegistrationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListDirectoryRegistrationsInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListDirectoryRegistrationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListDirectoryRegistrationsInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListDirectoryRegistrationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListDirectoryRegistrationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListDirectoryRegistrationsInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServicePrincipalNames provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListServicePrincipalNames(ctx context.Context, params *pcaconnectorad.ListServicePrincipalNamesInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListServicePrincipalNamesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListServicePrincipalNames")
	}

	var r0 *pcaconnectorad.ListServicePrincipalNamesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListServicePrincipalNamesInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListServicePrincipalNamesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListServicePrincipalNamesInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListServicePrincipalNamesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListServicePrincipalNamesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListServicePrincipalNamesInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *pcaconnectorad.ListTagsForResourceInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTagsForResourceOutput, error) {
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

	var r0 *pcaconnectorad.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTagsForResourceInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTagsForResourceInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListTagsForResourceInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTemplateGroupAccessControlEntries provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTemplateGroupAccessControlEntries(ctx context.Context, params *pcaconnectorad.ListTemplateGroupAccessControlEntriesInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTemplateGroupAccessControlEntries")
	}

	var r0 *pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTemplateGroupAccessControlEntriesInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTemplateGroupAccessControlEntriesInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListTemplateGroupAccessControlEntriesInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTemplates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTemplates(ctx context.Context, params *pcaconnectorad.ListTemplatesInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTemplatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTemplates")
	}

	var r0 *pcaconnectorad.ListTemplatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTemplatesInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.ListTemplatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.ListTemplatesInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.ListTemplatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.ListTemplatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.ListTemplatesInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() pcaconnectorad.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 pcaconnectorad.Options
	if rf, ok := ret.Get(0).(func() pcaconnectorad.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pcaconnectorad.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *pcaconnectorad.TagResourceInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.TagResourceOutput, error) {
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

	var r0 *pcaconnectorad.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.TagResourceInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.TagResourceInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.TagResourceInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *pcaconnectorad.UntagResourceInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UntagResourceOutput, error) {
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

	var r0 *pcaconnectorad.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UntagResourceInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UntagResourceInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.UntagResourceInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTemplate(ctx context.Context, params *pcaconnectorad.UpdateTemplateInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UpdateTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTemplate")
	}

	var r0 *pcaconnectorad.UpdateTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UpdateTemplateInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UpdateTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UpdateTemplateInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.UpdateTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.UpdateTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.UpdateTemplateInput, ...func(*pcaconnectorad.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTemplateGroupAccessControlEntry provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTemplateGroupAccessControlEntry(ctx context.Context, params *pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput, optFns ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTemplateGroupAccessControlEntry")
	}

	var r0 *pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) (*pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) *pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput, ...func(*pcaconnectorad.Options)) error); ok {
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
