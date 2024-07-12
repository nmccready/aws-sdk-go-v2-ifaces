// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	servicecatalogappregistry "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.AssociateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateAttributeGroup")
	}

	var r0 *servicecatalogappregistry.AssociateAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.AssociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.AssociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.AssociateAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.AssociateAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.AssociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateResource(ctx context.Context, params *servicecatalogappregistry.AssociateResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateResource")
	}

	var r0 *servicecatalogappregistry.AssociateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.AssociateResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.AssociateResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.AssociateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.AssociateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.AssociateResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateApplication(ctx context.Context, params *servicecatalogappregistry.CreateApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateApplicationOutput, error) {
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

	var r0 *servicecatalogappregistry.CreateApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.CreateApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.CreateApplicationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.CreateApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.CreateApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.CreateApplicationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.CreateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAttributeGroup")
	}

	var r0 *servicecatalogappregistry.CreateAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.CreateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.CreateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.CreateAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.CreateAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.CreateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteApplication(ctx context.Context, params *servicecatalogappregistry.DeleteApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteApplicationOutput, error) {
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

	var r0 *servicecatalogappregistry.DeleteApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DeleteApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DeleteApplicationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.DeleteApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.DeleteApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.DeleteApplicationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAttributeGroup(ctx context.Context, params *servicecatalogappregistry.DeleteAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAttributeGroup")
	}

	var r0 *servicecatalogappregistry.DeleteAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DeleteAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DeleteAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.DeleteAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.DeleteAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.DeleteAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.DisassociateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateAttributeGroup")
	}

	var r0 *servicecatalogappregistry.DisassociateAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DisassociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DisassociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.DisassociateAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.DisassociateAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.DisassociateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateResource(ctx context.Context, params *servicecatalogappregistry.DisassociateResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateResource")
	}

	var r0 *servicecatalogappregistry.DisassociateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DisassociateResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.DisassociateResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.DisassociateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.DisassociateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.DisassociateResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetApplication(ctx context.Context, params *servicecatalogappregistry.GetApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error) {
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

	var r0 *servicecatalogappregistry.GetApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetApplicationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.GetApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.GetApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.GetApplicationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssociatedResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAssociatedResource(ctx context.Context, params *servicecatalogappregistry.GetAssociatedResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAssociatedResource")
	}

	var r0 *servicecatalogappregistry.GetAssociatedResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetAssociatedResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetAssociatedResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.GetAssociatedResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.GetAssociatedResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.GetAssociatedResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAttributeGroup(ctx context.Context, params *servicecatalogappregistry.GetAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAttributeGroup")
	}

	var r0 *servicecatalogappregistry.GetAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.GetAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.GetAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.GetAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetConfiguration(ctx context.Context, params *servicecatalogappregistry.GetConfigurationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConfiguration")
	}

	var r0 *servicecatalogappregistry.GetConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetConfigurationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.GetConfigurationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.GetConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.GetConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.GetConfigurationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplications provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApplications(ctx context.Context, params *servicecatalogappregistry.ListApplicationsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error) {
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

	var r0 *servicecatalogappregistry.ListApplicationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListApplicationsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListApplicationsInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListApplicationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListApplicationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListApplicationsInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAssociatedAttributeGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAssociatedAttributeGroups(ctx context.Context, params *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAssociatedAttributeGroups")
	}

	var r0 *servicecatalogappregistry.ListAssociatedAttributeGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListAssociatedAttributeGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAssociatedResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAssociatedResources(ctx context.Context, params *servicecatalogappregistry.ListAssociatedResourcesInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAssociatedResources")
	}

	var r0 *servicecatalogappregistry.ListAssociatedResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAssociatedResourcesInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAssociatedResourcesInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListAssociatedResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListAssociatedResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListAssociatedResourcesInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAttributeGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAttributeGroups(ctx context.Context, params *servicecatalogappregistry.ListAttributeGroupsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAttributeGroups")
	}

	var r0 *servicecatalogappregistry.ListAttributeGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListAttributeGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListAttributeGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAttributeGroupsForApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAttributeGroupsForApplication(ctx context.Context, params *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAttributeGroupsForApplication")
	}

	var r0 *servicecatalogappregistry.ListAttributeGroupsForApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListAttributeGroupsForApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *servicecatalogappregistry.ListTagsForResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error) {
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

	var r0 *servicecatalogappregistry.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListTagsForResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.ListTagsForResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.ListTagsForResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() servicecatalogappregistry.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 servicecatalogappregistry.Options
	if rf, ok := ret.Get(0).(func() servicecatalogappregistry.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(servicecatalogappregistry.Options)
	}

	return r0
}

// PutConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutConfiguration(ctx context.Context, params *servicecatalogappregistry.PutConfigurationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.PutConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutConfiguration")
	}

	var r0 *servicecatalogappregistry.PutConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.PutConfigurationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.PutConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.PutConfigurationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.PutConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.PutConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.PutConfigurationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SyncResource(ctx context.Context, params *servicecatalogappregistry.SyncResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.SyncResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SyncResource")
	}

	var r0 *servicecatalogappregistry.SyncResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.SyncResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.SyncResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.SyncResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.SyncResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.SyncResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.SyncResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *servicecatalogappregistry.TagResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.TagResourceOutput, error) {
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

	var r0 *servicecatalogappregistry.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.TagResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.TagResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.TagResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *servicecatalogappregistry.UntagResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UntagResourceOutput, error) {
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

	var r0 *servicecatalogappregistry.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UntagResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UntagResourceInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.UntagResourceInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateApplication(ctx context.Context, params *servicecatalogappregistry.UpdateApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateApplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateApplication")
	}

	var r0 *servicecatalogappregistry.UpdateApplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UpdateApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateApplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UpdateApplicationInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.UpdateApplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.UpdateApplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.UpdateApplicationInput, ...func(*servicecatalogappregistry.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAttributeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.UpdateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateAttributeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAttributeGroup")
	}

	var r0 *servicecatalogappregistry.UpdateAttributeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UpdateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateAttributeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *servicecatalogappregistry.UpdateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) *servicecatalogappregistry.UpdateAttributeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecatalogappregistry.UpdateAttributeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *servicecatalogappregistry.UpdateAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) error); ok {
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