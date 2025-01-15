// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	fis "github.com/aws/aws-sdk-go-v2/service/fis"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateExperimentTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateExperimentTemplate(ctx context.Context, params *fis.CreateExperimentTemplateInput, optFns ...func(*fis.Options)) (*fis.CreateExperimentTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateExperimentTemplate")
	}

	var r0 *fis.CreateExperimentTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.CreateExperimentTemplateInput, ...func(*fis.Options)) (*fis.CreateExperimentTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.CreateExperimentTemplateInput, ...func(*fis.Options)) *fis.CreateExperimentTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.CreateExperimentTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.CreateExperimentTemplateInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTargetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateTargetAccountConfiguration(ctx context.Context, params *fis.CreateTargetAccountConfigurationInput, optFns ...func(*fis.Options)) (*fis.CreateTargetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTargetAccountConfiguration")
	}

	var r0 *fis.CreateTargetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.CreateTargetAccountConfigurationInput, ...func(*fis.Options)) (*fis.CreateTargetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.CreateTargetAccountConfigurationInput, ...func(*fis.Options)) *fis.CreateTargetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.CreateTargetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.CreateTargetAccountConfigurationInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExperimentTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteExperimentTemplate(ctx context.Context, params *fis.DeleteExperimentTemplateInput, optFns ...func(*fis.Options)) (*fis.DeleteExperimentTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExperimentTemplate")
	}

	var r0 *fis.DeleteExperimentTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.DeleteExperimentTemplateInput, ...func(*fis.Options)) (*fis.DeleteExperimentTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.DeleteExperimentTemplateInput, ...func(*fis.Options)) *fis.DeleteExperimentTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.DeleteExperimentTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.DeleteExperimentTemplateInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTargetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTargetAccountConfiguration(ctx context.Context, params *fis.DeleteTargetAccountConfigurationInput, optFns ...func(*fis.Options)) (*fis.DeleteTargetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTargetAccountConfiguration")
	}

	var r0 *fis.DeleteTargetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.DeleteTargetAccountConfigurationInput, ...func(*fis.Options)) (*fis.DeleteTargetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.DeleteTargetAccountConfigurationInput, ...func(*fis.Options)) *fis.DeleteTargetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.DeleteTargetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.DeleteTargetAccountConfigurationInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAction(ctx context.Context, params *fis.GetActionInput, optFns ...func(*fis.Options)) (*fis.GetActionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAction")
	}

	var r0 *fis.GetActionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetActionInput, ...func(*fis.Options)) (*fis.GetActionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetActionInput, ...func(*fis.Options)) *fis.GetActionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetActionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetActionInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExperiment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExperiment(ctx context.Context, params *fis.GetExperimentInput, optFns ...func(*fis.Options)) (*fis.GetExperimentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExperiment")
	}

	var r0 *fis.GetExperimentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentInput, ...func(*fis.Options)) (*fis.GetExperimentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentInput, ...func(*fis.Options)) *fis.GetExperimentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetExperimentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetExperimentInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExperimentTargetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExperimentTargetAccountConfiguration(ctx context.Context, params *fis.GetExperimentTargetAccountConfigurationInput, optFns ...func(*fis.Options)) (*fis.GetExperimentTargetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExperimentTargetAccountConfiguration")
	}

	var r0 *fis.GetExperimentTargetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentTargetAccountConfigurationInput, ...func(*fis.Options)) (*fis.GetExperimentTargetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentTargetAccountConfigurationInput, ...func(*fis.Options)) *fis.GetExperimentTargetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetExperimentTargetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetExperimentTargetAccountConfigurationInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExperimentTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExperimentTemplate(ctx context.Context, params *fis.GetExperimentTemplateInput, optFns ...func(*fis.Options)) (*fis.GetExperimentTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExperimentTemplate")
	}

	var r0 *fis.GetExperimentTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentTemplateInput, ...func(*fis.Options)) (*fis.GetExperimentTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetExperimentTemplateInput, ...func(*fis.Options)) *fis.GetExperimentTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetExperimentTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetExperimentTemplateInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTargetAccountConfiguration(ctx context.Context, params *fis.GetTargetAccountConfigurationInput, optFns ...func(*fis.Options)) (*fis.GetTargetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTargetAccountConfiguration")
	}

	var r0 *fis.GetTargetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetTargetAccountConfigurationInput, ...func(*fis.Options)) (*fis.GetTargetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetTargetAccountConfigurationInput, ...func(*fis.Options)) *fis.GetTargetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetTargetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetTargetAccountConfigurationInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetResourceType provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTargetResourceType(ctx context.Context, params *fis.GetTargetResourceTypeInput, optFns ...func(*fis.Options)) (*fis.GetTargetResourceTypeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTargetResourceType")
	}

	var r0 *fis.GetTargetResourceTypeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetTargetResourceTypeInput, ...func(*fis.Options)) (*fis.GetTargetResourceTypeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.GetTargetResourceTypeInput, ...func(*fis.Options)) *fis.GetTargetResourceTypeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.GetTargetResourceTypeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.GetTargetResourceTypeInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListActions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListActions(ctx context.Context, params *fis.ListActionsInput, optFns ...func(*fis.Options)) (*fis.ListActionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListActions")
	}

	var r0 *fis.ListActionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListActionsInput, ...func(*fis.Options)) (*fis.ListActionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListActionsInput, ...func(*fis.Options)) *fis.ListActionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListActionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListActionsInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExperimentResolvedTargets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExperimentResolvedTargets(ctx context.Context, params *fis.ListExperimentResolvedTargetsInput, optFns ...func(*fis.Options)) (*fis.ListExperimentResolvedTargetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExperimentResolvedTargets")
	}

	var r0 *fis.ListExperimentResolvedTargetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentResolvedTargetsInput, ...func(*fis.Options)) (*fis.ListExperimentResolvedTargetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentResolvedTargetsInput, ...func(*fis.Options)) *fis.ListExperimentResolvedTargetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListExperimentResolvedTargetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListExperimentResolvedTargetsInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExperimentTargetAccountConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExperimentTargetAccountConfigurations(ctx context.Context, params *fis.ListExperimentTargetAccountConfigurationsInput, optFns ...func(*fis.Options)) (*fis.ListExperimentTargetAccountConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExperimentTargetAccountConfigurations")
	}

	var r0 *fis.ListExperimentTargetAccountConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentTargetAccountConfigurationsInput, ...func(*fis.Options)) (*fis.ListExperimentTargetAccountConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentTargetAccountConfigurationsInput, ...func(*fis.Options)) *fis.ListExperimentTargetAccountConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListExperimentTargetAccountConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListExperimentTargetAccountConfigurationsInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExperimentTemplates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExperimentTemplates(ctx context.Context, params *fis.ListExperimentTemplatesInput, optFns ...func(*fis.Options)) (*fis.ListExperimentTemplatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExperimentTemplates")
	}

	var r0 *fis.ListExperimentTemplatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentTemplatesInput, ...func(*fis.Options)) (*fis.ListExperimentTemplatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentTemplatesInput, ...func(*fis.Options)) *fis.ListExperimentTemplatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListExperimentTemplatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListExperimentTemplatesInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExperiments provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExperiments(ctx context.Context, params *fis.ListExperimentsInput, optFns ...func(*fis.Options)) (*fis.ListExperimentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExperiments")
	}

	var r0 *fis.ListExperimentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentsInput, ...func(*fis.Options)) (*fis.ListExperimentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListExperimentsInput, ...func(*fis.Options)) *fis.ListExperimentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListExperimentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListExperimentsInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *fis.ListTagsForResourceInput, optFns ...func(*fis.Options)) (*fis.ListTagsForResourceOutput, error) {
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

	var r0 *fis.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTagsForResourceInput, ...func(*fis.Options)) (*fis.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTagsForResourceInput, ...func(*fis.Options)) *fis.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListTagsForResourceInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTargetAccountConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTargetAccountConfigurations(ctx context.Context, params *fis.ListTargetAccountConfigurationsInput, optFns ...func(*fis.Options)) (*fis.ListTargetAccountConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTargetAccountConfigurations")
	}

	var r0 *fis.ListTargetAccountConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTargetAccountConfigurationsInput, ...func(*fis.Options)) (*fis.ListTargetAccountConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTargetAccountConfigurationsInput, ...func(*fis.Options)) *fis.ListTargetAccountConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListTargetAccountConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListTargetAccountConfigurationsInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTargetResourceTypes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTargetResourceTypes(ctx context.Context, params *fis.ListTargetResourceTypesInput, optFns ...func(*fis.Options)) (*fis.ListTargetResourceTypesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTargetResourceTypes")
	}

	var r0 *fis.ListTargetResourceTypesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTargetResourceTypesInput, ...func(*fis.Options)) (*fis.ListTargetResourceTypesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.ListTargetResourceTypesInput, ...func(*fis.Options)) *fis.ListTargetResourceTypesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.ListTargetResourceTypesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.ListTargetResourceTypesInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() fis.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 fis.Options
	if rf, ok := ret.Get(0).(func() fis.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(fis.Options)
	}

	return r0
}

// StartExperiment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartExperiment(ctx context.Context, params *fis.StartExperimentInput, optFns ...func(*fis.Options)) (*fis.StartExperimentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartExperiment")
	}

	var r0 *fis.StartExperimentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.StartExperimentInput, ...func(*fis.Options)) (*fis.StartExperimentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.StartExperimentInput, ...func(*fis.Options)) *fis.StartExperimentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.StartExperimentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.StartExperimentInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopExperiment provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopExperiment(ctx context.Context, params *fis.StopExperimentInput, optFns ...func(*fis.Options)) (*fis.StopExperimentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopExperiment")
	}

	var r0 *fis.StopExperimentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.StopExperimentInput, ...func(*fis.Options)) (*fis.StopExperimentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.StopExperimentInput, ...func(*fis.Options)) *fis.StopExperimentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.StopExperimentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.StopExperimentInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *fis.TagResourceInput, optFns ...func(*fis.Options)) (*fis.TagResourceOutput, error) {
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

	var r0 *fis.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.TagResourceInput, ...func(*fis.Options)) (*fis.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.TagResourceInput, ...func(*fis.Options)) *fis.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.TagResourceInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *fis.UntagResourceInput, optFns ...func(*fis.Options)) (*fis.UntagResourceOutput, error) {
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

	var r0 *fis.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UntagResourceInput, ...func(*fis.Options)) (*fis.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UntagResourceInput, ...func(*fis.Options)) *fis.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.UntagResourceInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExperimentTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateExperimentTemplate(ctx context.Context, params *fis.UpdateExperimentTemplateInput, optFns ...func(*fis.Options)) (*fis.UpdateExperimentTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExperimentTemplate")
	}

	var r0 *fis.UpdateExperimentTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UpdateExperimentTemplateInput, ...func(*fis.Options)) (*fis.UpdateExperimentTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UpdateExperimentTemplateInput, ...func(*fis.Options)) *fis.UpdateExperimentTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.UpdateExperimentTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.UpdateExperimentTemplateInput, ...func(*fis.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTargetAccountConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTargetAccountConfiguration(ctx context.Context, params *fis.UpdateTargetAccountConfigurationInput, optFns ...func(*fis.Options)) (*fis.UpdateTargetAccountConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTargetAccountConfiguration")
	}

	var r0 *fis.UpdateTargetAccountConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UpdateTargetAccountConfigurationInput, ...func(*fis.Options)) (*fis.UpdateTargetAccountConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fis.UpdateTargetAccountConfigurationInput, ...func(*fis.Options)) *fis.UpdateTargetAccountConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fis.UpdateTargetAccountConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fis.UpdateTargetAccountConfigurationInput, ...func(*fis.Options)) error); ok {
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
