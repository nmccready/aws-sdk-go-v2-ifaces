// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	amplifyuibuilder "github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateComponent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateComponent(ctx context.Context, params *amplifyuibuilder.CreateComponentInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateComponentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateComponent")
	}

	var r0 *amplifyuibuilder.CreateComponentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateComponentInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateComponentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateComponentInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.CreateComponentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.CreateComponentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.CreateComponentInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateForm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateForm(ctx context.Context, params *amplifyuibuilder.CreateFormInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateFormOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateForm")
	}

	var r0 *amplifyuibuilder.CreateFormOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateFormInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateFormOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateFormInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.CreateFormOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.CreateFormOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.CreateFormInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTheme provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTheme(ctx context.Context, params *amplifyuibuilder.CreateThemeInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateThemeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTheme")
	}

	var r0 *amplifyuibuilder.CreateThemeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateThemeInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.CreateThemeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.CreateThemeInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.CreateThemeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.CreateThemeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.CreateThemeInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteComponent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteComponent(ctx context.Context, params *amplifyuibuilder.DeleteComponentInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteComponentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComponent")
	}

	var r0 *amplifyuibuilder.DeleteComponentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteComponentInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteComponentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteComponentInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.DeleteComponentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.DeleteComponentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.DeleteComponentInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteForm(ctx context.Context, params *amplifyuibuilder.DeleteFormInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteFormOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteForm")
	}

	var r0 *amplifyuibuilder.DeleteFormOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteFormInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteFormOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteFormInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.DeleteFormOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.DeleteFormOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.DeleteFormInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTheme provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTheme(ctx context.Context, params *amplifyuibuilder.DeleteThemeInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteThemeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTheme")
	}

	var r0 *amplifyuibuilder.DeleteThemeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteThemeInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.DeleteThemeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.DeleteThemeInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.DeleteThemeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.DeleteThemeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.DeleteThemeInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCodeForToken provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExchangeCodeForToken(ctx context.Context, params *amplifyuibuilder.ExchangeCodeForTokenInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExchangeCodeForTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeCodeForToken")
	}

	var r0 *amplifyuibuilder.ExchangeCodeForTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExchangeCodeForTokenInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExchangeCodeForTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExchangeCodeForTokenInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ExchangeCodeForTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ExchangeCodeForTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ExchangeCodeForTokenInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportComponents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExportComponents(ctx context.Context, params *amplifyuibuilder.ExportComponentsInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportComponentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportComponents")
	}

	var r0 *amplifyuibuilder.ExportComponentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportComponentsInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportComponentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportComponentsInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ExportComponentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ExportComponentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ExportComponentsInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportForms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExportForms(ctx context.Context, params *amplifyuibuilder.ExportFormsInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportFormsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportForms")
	}

	var r0 *amplifyuibuilder.ExportFormsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportFormsInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportFormsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportFormsInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ExportFormsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ExportFormsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ExportFormsInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportThemes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExportThemes(ctx context.Context, params *amplifyuibuilder.ExportThemesInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportThemesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportThemes")
	}

	var r0 *amplifyuibuilder.ExportThemesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportThemesInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ExportThemesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ExportThemesInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ExportThemesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ExportThemesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ExportThemesInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCodegenJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCodegenJob(ctx context.Context, params *amplifyuibuilder.GetCodegenJobInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetCodegenJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCodegenJob")
	}

	var r0 *amplifyuibuilder.GetCodegenJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetCodegenJobInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetCodegenJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetCodegenJobInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.GetCodegenJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.GetCodegenJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.GetCodegenJobInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetComponent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetComponent(ctx context.Context, params *amplifyuibuilder.GetComponentInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetComponentOutput, error) {
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

	var r0 *amplifyuibuilder.GetComponentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetComponentInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetComponentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetComponentInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.GetComponentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.GetComponentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.GetComponentInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetForm(ctx context.Context, params *amplifyuibuilder.GetFormInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetFormOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetForm")
	}

	var r0 *amplifyuibuilder.GetFormOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetFormInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetFormOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetFormInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.GetFormOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.GetFormOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.GetFormInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMetadata(ctx context.Context, params *amplifyuibuilder.GetMetadataInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetMetadataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 *amplifyuibuilder.GetMetadataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetMetadataInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetMetadataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetMetadataInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.GetMetadataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.GetMetadataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.GetMetadataInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTheme provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTheme(ctx context.Context, params *amplifyuibuilder.GetThemeInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetThemeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTheme")
	}

	var r0 *amplifyuibuilder.GetThemeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetThemeInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.GetThemeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.GetThemeInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.GetThemeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.GetThemeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.GetThemeInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCodegenJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCodegenJobs(ctx context.Context, params *amplifyuibuilder.ListCodegenJobsInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListCodegenJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCodegenJobs")
	}

	var r0 *amplifyuibuilder.ListCodegenJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListCodegenJobsInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListCodegenJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListCodegenJobsInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ListCodegenJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ListCodegenJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ListCodegenJobsInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListComponents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListComponents(ctx context.Context, params *amplifyuibuilder.ListComponentsInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListComponentsOutput, error) {
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

	var r0 *amplifyuibuilder.ListComponentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListComponentsInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListComponentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListComponentsInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ListComponentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ListComponentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ListComponentsInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListForms(ctx context.Context, params *amplifyuibuilder.ListFormsInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListFormsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListForms")
	}

	var r0 *amplifyuibuilder.ListFormsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListFormsInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListFormsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListFormsInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ListFormsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ListFormsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ListFormsInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *amplifyuibuilder.ListTagsForResourceInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListTagsForResourceOutput, error) {
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

	var r0 *amplifyuibuilder.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListTagsForResourceInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListTagsForResourceInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ListTagsForResourceInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListThemes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListThemes(ctx context.Context, params *amplifyuibuilder.ListThemesInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListThemesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListThemes")
	}

	var r0 *amplifyuibuilder.ListThemesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListThemesInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.ListThemesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.ListThemesInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.ListThemesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.ListThemesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.ListThemesInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() amplifyuibuilder.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 amplifyuibuilder.Options
	if rf, ok := ret.Get(0).(func() amplifyuibuilder.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(amplifyuibuilder.Options)
	}

	return r0
}

// PutMetadataFlag provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutMetadataFlag(ctx context.Context, params *amplifyuibuilder.PutMetadataFlagInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.PutMetadataFlagOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutMetadataFlag")
	}

	var r0 *amplifyuibuilder.PutMetadataFlagOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.PutMetadataFlagInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.PutMetadataFlagOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.PutMetadataFlagInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.PutMetadataFlagOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.PutMetadataFlagOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.PutMetadataFlagInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RefreshToken(ctx context.Context, params *amplifyuibuilder.RefreshTokenInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.RefreshTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 *amplifyuibuilder.RefreshTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.RefreshTokenInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.RefreshTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.RefreshTokenInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.RefreshTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.RefreshTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.RefreshTokenInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartCodegenJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartCodegenJob(ctx context.Context, params *amplifyuibuilder.StartCodegenJobInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.StartCodegenJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartCodegenJob")
	}

	var r0 *amplifyuibuilder.StartCodegenJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.StartCodegenJobInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.StartCodegenJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.StartCodegenJobInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.StartCodegenJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.StartCodegenJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.StartCodegenJobInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *amplifyuibuilder.TagResourceInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.TagResourceOutput, error) {
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

	var r0 *amplifyuibuilder.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.TagResourceInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.TagResourceInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.TagResourceInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *amplifyuibuilder.UntagResourceInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UntagResourceOutput, error) {
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

	var r0 *amplifyuibuilder.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UntagResourceInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UntagResourceInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.UntagResourceInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComponent provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateComponent(ctx context.Context, params *amplifyuibuilder.UpdateComponentInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateComponentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateComponent")
	}

	var r0 *amplifyuibuilder.UpdateComponentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateComponentInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateComponentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateComponentInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.UpdateComponentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.UpdateComponentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.UpdateComponentInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateForm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateForm(ctx context.Context, params *amplifyuibuilder.UpdateFormInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateFormOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateForm")
	}

	var r0 *amplifyuibuilder.UpdateFormOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateFormInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateFormOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateFormInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.UpdateFormOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.UpdateFormOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.UpdateFormInput, ...func(*amplifyuibuilder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTheme provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTheme(ctx context.Context, params *amplifyuibuilder.UpdateThemeInput, optFns ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateThemeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTheme")
	}

	var r0 *amplifyuibuilder.UpdateThemeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateThemeInput, ...func(*amplifyuibuilder.Options)) (*amplifyuibuilder.UpdateThemeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *amplifyuibuilder.UpdateThemeInput, ...func(*amplifyuibuilder.Options)) *amplifyuibuilder.UpdateThemeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amplifyuibuilder.UpdateThemeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *amplifyuibuilder.UpdateThemeInput, ...func(*amplifyuibuilder.Options)) error); ok {
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
