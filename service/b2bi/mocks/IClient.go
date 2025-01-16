// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	b2bi "github.com/aws/aws-sdk-go-v2/service/b2bi"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateCapability provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCapability(ctx context.Context, params *b2bi.CreateCapabilityInput, optFns ...func(*b2bi.Options)) (*b2bi.CreateCapabilityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCapability")
	}

	var r0 *b2bi.CreateCapabilityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateCapabilityInput, ...func(*b2bi.Options)) (*b2bi.CreateCapabilityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateCapabilityInput, ...func(*b2bi.Options)) *b2bi.CreateCapabilityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.CreateCapabilityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.CreateCapabilityInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePartnership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePartnership(ctx context.Context, params *b2bi.CreatePartnershipInput, optFns ...func(*b2bi.Options)) (*b2bi.CreatePartnershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePartnership")
	}

	var r0 *b2bi.CreatePartnershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreatePartnershipInput, ...func(*b2bi.Options)) (*b2bi.CreatePartnershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreatePartnershipInput, ...func(*b2bi.Options)) *b2bi.CreatePartnershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.CreatePartnershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.CreatePartnershipInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProfile(ctx context.Context, params *b2bi.CreateProfileInput, optFns ...func(*b2bi.Options)) (*b2bi.CreateProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProfile")
	}

	var r0 *b2bi.CreateProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateProfileInput, ...func(*b2bi.Options)) (*b2bi.CreateProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateProfileInput, ...func(*b2bi.Options)) *b2bi.CreateProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.CreateProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.CreateProfileInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStarterMappingTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateStarterMappingTemplate(ctx context.Context, params *b2bi.CreateStarterMappingTemplateInput, optFns ...func(*b2bi.Options)) (*b2bi.CreateStarterMappingTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateStarterMappingTemplate")
	}

	var r0 *b2bi.CreateStarterMappingTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateStarterMappingTemplateInput, ...func(*b2bi.Options)) (*b2bi.CreateStarterMappingTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateStarterMappingTemplateInput, ...func(*b2bi.Options)) *b2bi.CreateStarterMappingTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.CreateStarterMappingTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.CreateStarterMappingTemplateInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransformer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTransformer(ctx context.Context, params *b2bi.CreateTransformerInput, optFns ...func(*b2bi.Options)) (*b2bi.CreateTransformerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransformer")
	}

	var r0 *b2bi.CreateTransformerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateTransformerInput, ...func(*b2bi.Options)) (*b2bi.CreateTransformerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.CreateTransformerInput, ...func(*b2bi.Options)) *b2bi.CreateTransformerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.CreateTransformerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.CreateTransformerInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCapability provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCapability(ctx context.Context, params *b2bi.DeleteCapabilityInput, optFns ...func(*b2bi.Options)) (*b2bi.DeleteCapabilityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCapability")
	}

	var r0 *b2bi.DeleteCapabilityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteCapabilityInput, ...func(*b2bi.Options)) (*b2bi.DeleteCapabilityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteCapabilityInput, ...func(*b2bi.Options)) *b2bi.DeleteCapabilityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.DeleteCapabilityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.DeleteCapabilityInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePartnership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePartnership(ctx context.Context, params *b2bi.DeletePartnershipInput, optFns ...func(*b2bi.Options)) (*b2bi.DeletePartnershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePartnership")
	}

	var r0 *b2bi.DeletePartnershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeletePartnershipInput, ...func(*b2bi.Options)) (*b2bi.DeletePartnershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeletePartnershipInput, ...func(*b2bi.Options)) *b2bi.DeletePartnershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.DeletePartnershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.DeletePartnershipInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProfile(ctx context.Context, params *b2bi.DeleteProfileInput, optFns ...func(*b2bi.Options)) (*b2bi.DeleteProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProfile")
	}

	var r0 *b2bi.DeleteProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteProfileInput, ...func(*b2bi.Options)) (*b2bi.DeleteProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteProfileInput, ...func(*b2bi.Options)) *b2bi.DeleteProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.DeleteProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.DeleteProfileInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransformer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTransformer(ctx context.Context, params *b2bi.DeleteTransformerInput, optFns ...func(*b2bi.Options)) (*b2bi.DeleteTransformerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransformer")
	}

	var r0 *b2bi.DeleteTransformerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteTransformerInput, ...func(*b2bi.Options)) (*b2bi.DeleteTransformerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.DeleteTransformerInput, ...func(*b2bi.Options)) *b2bi.DeleteTransformerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.DeleteTransformerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.DeleteTransformerInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateMapping provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateMapping(ctx context.Context, params *b2bi.GenerateMappingInput, optFns ...func(*b2bi.Options)) (*b2bi.GenerateMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateMapping")
	}

	var r0 *b2bi.GenerateMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GenerateMappingInput, ...func(*b2bi.Options)) (*b2bi.GenerateMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GenerateMappingInput, ...func(*b2bi.Options)) *b2bi.GenerateMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GenerateMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GenerateMappingInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCapability provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCapability(ctx context.Context, params *b2bi.GetCapabilityInput, optFns ...func(*b2bi.Options)) (*b2bi.GetCapabilityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCapability")
	}

	var r0 *b2bi.GetCapabilityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetCapabilityInput, ...func(*b2bi.Options)) (*b2bi.GetCapabilityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetCapabilityInput, ...func(*b2bi.Options)) *b2bi.GetCapabilityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GetCapabilityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GetCapabilityInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPartnership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPartnership(ctx context.Context, params *b2bi.GetPartnershipInput, optFns ...func(*b2bi.Options)) (*b2bi.GetPartnershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPartnership")
	}

	var r0 *b2bi.GetPartnershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetPartnershipInput, ...func(*b2bi.Options)) (*b2bi.GetPartnershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetPartnershipInput, ...func(*b2bi.Options)) *b2bi.GetPartnershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GetPartnershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GetPartnershipInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetProfile(ctx context.Context, params *b2bi.GetProfileInput, optFns ...func(*b2bi.Options)) (*b2bi.GetProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProfile")
	}

	var r0 *b2bi.GetProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetProfileInput, ...func(*b2bi.Options)) (*b2bi.GetProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetProfileInput, ...func(*b2bi.Options)) *b2bi.GetProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GetProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GetProfileInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransformer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTransformer(ctx context.Context, params *b2bi.GetTransformerInput, optFns ...func(*b2bi.Options)) (*b2bi.GetTransformerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTransformer")
	}

	var r0 *b2bi.GetTransformerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetTransformerInput, ...func(*b2bi.Options)) (*b2bi.GetTransformerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetTransformerInput, ...func(*b2bi.Options)) *b2bi.GetTransformerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GetTransformerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GetTransformerInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransformerJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTransformerJob(ctx context.Context, params *b2bi.GetTransformerJobInput, optFns ...func(*b2bi.Options)) (*b2bi.GetTransformerJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTransformerJob")
	}

	var r0 *b2bi.GetTransformerJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetTransformerJobInput, ...func(*b2bi.Options)) (*b2bi.GetTransformerJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.GetTransformerJobInput, ...func(*b2bi.Options)) *b2bi.GetTransformerJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.GetTransformerJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.GetTransformerJobInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCapabilities provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCapabilities(ctx context.Context, params *b2bi.ListCapabilitiesInput, optFns ...func(*b2bi.Options)) (*b2bi.ListCapabilitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCapabilities")
	}

	var r0 *b2bi.ListCapabilitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListCapabilitiesInput, ...func(*b2bi.Options)) (*b2bi.ListCapabilitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListCapabilitiesInput, ...func(*b2bi.Options)) *b2bi.ListCapabilitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.ListCapabilitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.ListCapabilitiesInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPartnerships provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPartnerships(ctx context.Context, params *b2bi.ListPartnershipsInput, optFns ...func(*b2bi.Options)) (*b2bi.ListPartnershipsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPartnerships")
	}

	var r0 *b2bi.ListPartnershipsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListPartnershipsInput, ...func(*b2bi.Options)) (*b2bi.ListPartnershipsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListPartnershipsInput, ...func(*b2bi.Options)) *b2bi.ListPartnershipsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.ListPartnershipsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.ListPartnershipsInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProfiles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProfiles(ctx context.Context, params *b2bi.ListProfilesInput, optFns ...func(*b2bi.Options)) (*b2bi.ListProfilesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProfiles")
	}

	var r0 *b2bi.ListProfilesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListProfilesInput, ...func(*b2bi.Options)) (*b2bi.ListProfilesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListProfilesInput, ...func(*b2bi.Options)) *b2bi.ListProfilesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.ListProfilesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.ListProfilesInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *b2bi.ListTagsForResourceInput, optFns ...func(*b2bi.Options)) (*b2bi.ListTagsForResourceOutput, error) {
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

	var r0 *b2bi.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListTagsForResourceInput, ...func(*b2bi.Options)) (*b2bi.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListTagsForResourceInput, ...func(*b2bi.Options)) *b2bi.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.ListTagsForResourceInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransformers provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTransformers(ctx context.Context, params *b2bi.ListTransformersInput, optFns ...func(*b2bi.Options)) (*b2bi.ListTransformersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTransformers")
	}

	var r0 *b2bi.ListTransformersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListTransformersInput, ...func(*b2bi.Options)) (*b2bi.ListTransformersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.ListTransformersInput, ...func(*b2bi.Options)) *b2bi.ListTransformersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.ListTransformersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.ListTransformersInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() b2bi.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 b2bi.Options
	if rf, ok := ret.Get(0).(func() b2bi.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(b2bi.Options)
	}

	return r0
}

// StartTransformerJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartTransformerJob(ctx context.Context, params *b2bi.StartTransformerJobInput, optFns ...func(*b2bi.Options)) (*b2bi.StartTransformerJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartTransformerJob")
	}

	var r0 *b2bi.StartTransformerJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.StartTransformerJobInput, ...func(*b2bi.Options)) (*b2bi.StartTransformerJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.StartTransformerJobInput, ...func(*b2bi.Options)) *b2bi.StartTransformerJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.StartTransformerJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.StartTransformerJobInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *b2bi.TagResourceInput, optFns ...func(*b2bi.Options)) (*b2bi.TagResourceOutput, error) {
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

	var r0 *b2bi.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TagResourceInput, ...func(*b2bi.Options)) (*b2bi.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TagResourceInput, ...func(*b2bi.Options)) *b2bi.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.TagResourceInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestConversion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TestConversion(ctx context.Context, params *b2bi.TestConversionInput, optFns ...func(*b2bi.Options)) (*b2bi.TestConversionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TestConversion")
	}

	var r0 *b2bi.TestConversionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestConversionInput, ...func(*b2bi.Options)) (*b2bi.TestConversionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestConversionInput, ...func(*b2bi.Options)) *b2bi.TestConversionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.TestConversionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.TestConversionInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestMapping provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TestMapping(ctx context.Context, params *b2bi.TestMappingInput, optFns ...func(*b2bi.Options)) (*b2bi.TestMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TestMapping")
	}

	var r0 *b2bi.TestMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestMappingInput, ...func(*b2bi.Options)) (*b2bi.TestMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestMappingInput, ...func(*b2bi.Options)) *b2bi.TestMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.TestMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.TestMappingInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestParsing provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TestParsing(ctx context.Context, params *b2bi.TestParsingInput, optFns ...func(*b2bi.Options)) (*b2bi.TestParsingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TestParsing")
	}

	var r0 *b2bi.TestParsingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestParsingInput, ...func(*b2bi.Options)) (*b2bi.TestParsingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.TestParsingInput, ...func(*b2bi.Options)) *b2bi.TestParsingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.TestParsingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.TestParsingInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *b2bi.UntagResourceInput, optFns ...func(*b2bi.Options)) (*b2bi.UntagResourceOutput, error) {
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

	var r0 *b2bi.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UntagResourceInput, ...func(*b2bi.Options)) (*b2bi.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UntagResourceInput, ...func(*b2bi.Options)) *b2bi.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.UntagResourceInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCapability provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCapability(ctx context.Context, params *b2bi.UpdateCapabilityInput, optFns ...func(*b2bi.Options)) (*b2bi.UpdateCapabilityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCapability")
	}

	var r0 *b2bi.UpdateCapabilityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateCapabilityInput, ...func(*b2bi.Options)) (*b2bi.UpdateCapabilityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateCapabilityInput, ...func(*b2bi.Options)) *b2bi.UpdateCapabilityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.UpdateCapabilityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.UpdateCapabilityInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePartnership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePartnership(ctx context.Context, params *b2bi.UpdatePartnershipInput, optFns ...func(*b2bi.Options)) (*b2bi.UpdatePartnershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePartnership")
	}

	var r0 *b2bi.UpdatePartnershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdatePartnershipInput, ...func(*b2bi.Options)) (*b2bi.UpdatePartnershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdatePartnershipInput, ...func(*b2bi.Options)) *b2bi.UpdatePartnershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.UpdatePartnershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.UpdatePartnershipInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateProfile(ctx context.Context, params *b2bi.UpdateProfileInput, optFns ...func(*b2bi.Options)) (*b2bi.UpdateProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfile")
	}

	var r0 *b2bi.UpdateProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateProfileInput, ...func(*b2bi.Options)) (*b2bi.UpdateProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateProfileInput, ...func(*b2bi.Options)) *b2bi.UpdateProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.UpdateProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.UpdateProfileInput, ...func(*b2bi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransformer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTransformer(ctx context.Context, params *b2bi.UpdateTransformerInput, optFns ...func(*b2bi.Options)) (*b2bi.UpdateTransformerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransformer")
	}

	var r0 *b2bi.UpdateTransformerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateTransformerInput, ...func(*b2bi.Options)) (*b2bi.UpdateTransformerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *b2bi.UpdateTransformerInput, ...func(*b2bi.Options)) *b2bi.UpdateTransformerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*b2bi.UpdateTransformerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *b2bi.UpdateTransformerInput, ...func(*b2bi.Options)) error); ok {
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
