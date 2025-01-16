// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	verifiedpermissions "github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchGetPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetPolicy(ctx context.Context, params *verifiedpermissions.BatchGetPolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchGetPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetPolicy")
	}

	var r0 *verifiedpermissions.BatchGetPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchGetPolicyInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchGetPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchGetPolicyInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.BatchGetPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.BatchGetPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.BatchGetPolicyInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchIsAuthorized provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchIsAuthorized(ctx context.Context, params *verifiedpermissions.BatchIsAuthorizedInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchIsAuthorized")
	}

	var r0 *verifiedpermissions.BatchIsAuthorizedOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.BatchIsAuthorizedOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.BatchIsAuthorizedOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchIsAuthorizedWithToken provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchIsAuthorizedWithToken(ctx context.Context, params *verifiedpermissions.BatchIsAuthorizedWithTokenInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedWithTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchIsAuthorizedWithToken")
	}

	var r0 *verifiedpermissions.BatchIsAuthorizedWithTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedWithTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.BatchIsAuthorizedWithTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.BatchIsAuthorizedWithTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.BatchIsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIdentitySource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateIdentitySource(ctx context.Context, params *verifiedpermissions.CreateIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreateIdentitySourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateIdentitySource")
	}

	var r0 *verifiedpermissions.CreateIdentitySourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreateIdentitySourceInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreateIdentitySourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreateIdentitySourceInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.CreateIdentitySourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.CreateIdentitySourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.CreateIdentitySourceInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePolicy(ctx context.Context, params *verifiedpermissions.CreatePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePolicy")
	}

	var r0 *verifiedpermissions.CreatePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.CreatePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.CreatePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.CreatePolicyInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePolicyStore provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePolicyStore(ctx context.Context, params *verifiedpermissions.CreatePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyStoreOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePolicyStore")
	}

	var r0 *verifiedpermissions.CreatePolicyStoreOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyStoreInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyStoreOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyStoreInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.CreatePolicyStoreOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.CreatePolicyStoreOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.CreatePolicyStoreInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePolicyTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePolicyTemplate(ctx context.Context, params *verifiedpermissions.CreatePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePolicyTemplate")
	}

	var r0 *verifiedpermissions.CreatePolicyTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.CreatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.CreatePolicyTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.CreatePolicyTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.CreatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIdentitySource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteIdentitySource(ctx context.Context, params *verifiedpermissions.DeleteIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeleteIdentitySourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIdentitySource")
	}

	var r0 *verifiedpermissions.DeleteIdentitySourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeleteIdentitySourceInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeleteIdentitySourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeleteIdentitySourceInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.DeleteIdentitySourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.DeleteIdentitySourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.DeleteIdentitySourceInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePolicy(ctx context.Context, params *verifiedpermissions.DeletePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicy")
	}

	var r0 *verifiedpermissions.DeletePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.DeletePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.DeletePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.DeletePolicyInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicyStore provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePolicyStore(ctx context.Context, params *verifiedpermissions.DeletePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyStoreOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicyStore")
	}

	var r0 *verifiedpermissions.DeletePolicyStoreOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyStoreInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyStoreOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyStoreInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.DeletePolicyStoreOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.DeletePolicyStoreOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.DeletePolicyStoreInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicyTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePolicyTemplate(ctx context.Context, params *verifiedpermissions.DeletePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicyTemplate")
	}

	var r0 *verifiedpermissions.DeletePolicyTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyTemplateInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.DeletePolicyTemplateInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.DeletePolicyTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.DeletePolicyTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.DeletePolicyTemplateInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentitySource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetIdentitySource(ctx context.Context, params *verifiedpermissions.GetIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetIdentitySourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIdentitySource")
	}

	var r0 *verifiedpermissions.GetIdentitySourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetIdentitySourceInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetIdentitySourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetIdentitySourceInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.GetIdentitySourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.GetIdentitySourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.GetIdentitySourceInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPolicy(ctx context.Context, params *verifiedpermissions.GetPolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPolicy")
	}

	var r0 *verifiedpermissions.GetPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.GetPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.GetPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.GetPolicyInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPolicyStore provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPolicyStore(ctx context.Context, params *verifiedpermissions.GetPolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyStoreOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPolicyStore")
	}

	var r0 *verifiedpermissions.GetPolicyStoreOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyStoreInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyStoreOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyStoreInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.GetPolicyStoreOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.GetPolicyStoreOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.GetPolicyStoreInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPolicyTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPolicyTemplate(ctx context.Context, params *verifiedpermissions.GetPolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPolicyTemplate")
	}

	var r0 *verifiedpermissions.GetPolicyTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyTemplateInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetPolicyTemplateInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.GetPolicyTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.GetPolicyTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.GetPolicyTemplateInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchema provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSchema(ctx context.Context, params *verifiedpermissions.GetSchemaInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetSchemaOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSchema")
	}

	var r0 *verifiedpermissions.GetSchemaOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetSchemaInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetSchemaOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.GetSchemaInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.GetSchemaOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.GetSchemaOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.GetSchemaInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorized provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) IsAuthorized(ctx context.Context, params *verifiedpermissions.IsAuthorizedInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsAuthorized")
	}

	var r0 *verifiedpermissions.IsAuthorizedOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.IsAuthorizedInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.IsAuthorizedInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.IsAuthorizedOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.IsAuthorizedOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.IsAuthorizedInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorizedWithToken provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) IsAuthorizedWithToken(ctx context.Context, params *verifiedpermissions.IsAuthorizedWithTokenInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedWithTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsAuthorizedWithToken")
	}

	var r0 *verifiedpermissions.IsAuthorizedWithTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.IsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedWithTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.IsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.IsAuthorizedWithTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.IsAuthorizedWithTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.IsAuthorizedWithTokenInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentitySources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListIdentitySources(ctx context.Context, params *verifiedpermissions.ListIdentitySourcesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListIdentitySourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdentitySources")
	}

	var r0 *verifiedpermissions.ListIdentitySourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListIdentitySourcesInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListIdentitySourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListIdentitySourcesInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.ListIdentitySourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.ListIdentitySourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.ListIdentitySourcesInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicies provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPolicies(ctx context.Context, params *verifiedpermissions.ListPoliciesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPoliciesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPolicies")
	}

	var r0 *verifiedpermissions.ListPoliciesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPoliciesInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPoliciesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPoliciesInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.ListPoliciesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.ListPoliciesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.ListPoliciesInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicyStores provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPolicyStores(ctx context.Context, params *verifiedpermissions.ListPolicyStoresInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyStoresOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPolicyStores")
	}

	var r0 *verifiedpermissions.ListPolicyStoresOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPolicyStoresInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyStoresOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPolicyStoresInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.ListPolicyStoresOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.ListPolicyStoresOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.ListPolicyStoresInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicyTemplates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPolicyTemplates(ctx context.Context, params *verifiedpermissions.ListPolicyTemplatesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyTemplatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPolicyTemplates")
	}

	var r0 *verifiedpermissions.ListPolicyTemplatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPolicyTemplatesInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyTemplatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.ListPolicyTemplatesInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.ListPolicyTemplatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.ListPolicyTemplatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.ListPolicyTemplatesInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() verifiedpermissions.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 verifiedpermissions.Options
	if rf, ok := ret.Get(0).(func() verifiedpermissions.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(verifiedpermissions.Options)
	}

	return r0
}

// PutSchema provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutSchema(ctx context.Context, params *verifiedpermissions.PutSchemaInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.PutSchemaOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutSchema")
	}

	var r0 *verifiedpermissions.PutSchemaOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.PutSchemaInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.PutSchemaOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.PutSchemaInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.PutSchemaOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.PutSchemaOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.PutSchemaInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentitySource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateIdentitySource(ctx context.Context, params *verifiedpermissions.UpdateIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdateIdentitySourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIdentitySource")
	}

	var r0 *verifiedpermissions.UpdateIdentitySourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdateIdentitySourceInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdateIdentitySourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdateIdentitySourceInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.UpdateIdentitySourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.UpdateIdentitySourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.UpdateIdentitySourceInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePolicy(ctx context.Context, params *verifiedpermissions.UpdatePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePolicy")
	}

	var r0 *verifiedpermissions.UpdatePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.UpdatePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.UpdatePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.UpdatePolicyInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicyStore provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePolicyStore(ctx context.Context, params *verifiedpermissions.UpdatePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyStoreOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePolicyStore")
	}

	var r0 *verifiedpermissions.UpdatePolicyStoreOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyStoreInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyStoreOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyStoreInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.UpdatePolicyStoreOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.UpdatePolicyStoreOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.UpdatePolicyStoreInput, ...func(*verifiedpermissions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicyTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePolicyTemplate(ctx context.Context, params *verifiedpermissions.UpdatePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePolicyTemplate")
	}

	var r0 *verifiedpermissions.UpdatePolicyTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *verifiedpermissions.UpdatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) *verifiedpermissions.UpdatePolicyTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifiedpermissions.UpdatePolicyTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *verifiedpermissions.UpdatePolicyTemplateInput, ...func(*verifiedpermissions.Options)) error); ok {
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
