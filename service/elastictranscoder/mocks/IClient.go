// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	elastictranscoder "github.com/aws/aws-sdk-go-v2/service/elastictranscoder"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelJob(ctx context.Context, params *elastictranscoder.CancelJobInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.CancelJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelJob")
	}

	var r0 *elastictranscoder.CancelJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CancelJobInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.CancelJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CancelJobInput, ...func(*elastictranscoder.Options)) *elastictranscoder.CancelJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.CancelJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.CancelJobInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateJob(ctx context.Context, params *elastictranscoder.CreateJobInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.CreateJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 *elastictranscoder.CreateJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreateJobInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.CreateJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreateJobInput, ...func(*elastictranscoder.Options)) *elastictranscoder.CreateJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.CreateJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.CreateJobInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePipeline provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePipeline(ctx context.Context, params *elastictranscoder.CreatePipelineInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.CreatePipelineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePipeline")
	}

	var r0 *elastictranscoder.CreatePipelineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreatePipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.CreatePipelineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreatePipelineInput, ...func(*elastictranscoder.Options)) *elastictranscoder.CreatePipelineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.CreatePipelineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.CreatePipelineInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePreset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePreset(ctx context.Context, params *elastictranscoder.CreatePresetInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.CreatePresetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePreset")
	}

	var r0 *elastictranscoder.CreatePresetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreatePresetInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.CreatePresetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.CreatePresetInput, ...func(*elastictranscoder.Options)) *elastictranscoder.CreatePresetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.CreatePresetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.CreatePresetInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePipeline provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePipeline(ctx context.Context, params *elastictranscoder.DeletePipelineInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.DeletePipelineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePipeline")
	}

	var r0 *elastictranscoder.DeletePipelineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.DeletePipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.DeletePipelineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.DeletePipelineInput, ...func(*elastictranscoder.Options)) *elastictranscoder.DeletePipelineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.DeletePipelineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.DeletePipelineInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePreset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePreset(ctx context.Context, params *elastictranscoder.DeletePresetInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.DeletePresetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePreset")
	}

	var r0 *elastictranscoder.DeletePresetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.DeletePresetInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.DeletePresetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.DeletePresetInput, ...func(*elastictranscoder.Options)) *elastictranscoder.DeletePresetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.DeletePresetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.DeletePresetInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListJobsByPipeline provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListJobsByPipeline(ctx context.Context, params *elastictranscoder.ListJobsByPipelineInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListJobsByPipeline")
	}

	var r0 *elastictranscoder.ListJobsByPipelineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListJobsByPipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListJobsByPipelineInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ListJobsByPipelineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ListJobsByPipelineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ListJobsByPipelineInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListJobsByStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListJobsByStatus(ctx context.Context, params *elastictranscoder.ListJobsByStatusInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListJobsByStatus")
	}

	var r0 *elastictranscoder.ListJobsByStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListJobsByStatusInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListJobsByStatusInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ListJobsByStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ListJobsByStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ListJobsByStatusInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPipelines provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPipelines(ctx context.Context, params *elastictranscoder.ListPipelinesInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPipelines")
	}

	var r0 *elastictranscoder.ListPipelinesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListPipelinesInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListPipelinesInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ListPipelinesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ListPipelinesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ListPipelinesInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPresets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPresets(ctx context.Context, params *elastictranscoder.ListPresetsInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPresets")
	}

	var r0 *elastictranscoder.ListPresetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListPresetsInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ListPresetsInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ListPresetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ListPresetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ListPresetsInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() elastictranscoder.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 elastictranscoder.Options
	if rf, ok := ret.Get(0).(func() elastictranscoder.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(elastictranscoder.Options)
	}

	return r0
}

// ReadJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ReadJob(ctx context.Context, params *elastictranscoder.ReadJobInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReadJob")
	}

	var r0 *elastictranscoder.ReadJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadJobInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadJobInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ReadJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ReadJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ReadJobInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadPipeline provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ReadPipeline(ctx context.Context, params *elastictranscoder.ReadPipelineInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadPipelineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReadPipeline")
	}

	var r0 *elastictranscoder.ReadPipelineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadPipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadPipelineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadPipelineInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ReadPipelineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ReadPipelineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ReadPipelineInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadPreset provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ReadPreset(ctx context.Context, params *elastictranscoder.ReadPresetInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadPresetOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReadPreset")
	}

	var r0 *elastictranscoder.ReadPresetOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadPresetInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ReadPresetOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.ReadPresetInput, ...func(*elastictranscoder.Options)) *elastictranscoder.ReadPresetOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.ReadPresetOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.ReadPresetInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestRole provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TestRole(ctx context.Context, params *elastictranscoder.TestRoleInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.TestRoleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TestRole")
	}

	var r0 *elastictranscoder.TestRoleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.TestRoleInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.TestRoleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.TestRoleInput, ...func(*elastictranscoder.Options)) *elastictranscoder.TestRoleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.TestRoleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.TestRoleInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePipeline provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePipeline(ctx context.Context, params *elastictranscoder.UpdatePipelineInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePipeline")
	}

	var r0 *elastictranscoder.UpdatePipelineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineInput, ...func(*elastictranscoder.Options)) *elastictranscoder.UpdatePipelineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.UpdatePipelineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.UpdatePipelineInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePipelineNotifications provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePipelineNotifications(ctx context.Context, params *elastictranscoder.UpdatePipelineNotificationsInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePipelineNotifications")
	}

	var r0 *elastictranscoder.UpdatePipelineNotificationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineNotificationsInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineNotificationsInput, ...func(*elastictranscoder.Options)) *elastictranscoder.UpdatePipelineNotificationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.UpdatePipelineNotificationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.UpdatePipelineNotificationsInput, ...func(*elastictranscoder.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePipelineStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePipelineStatus(ctx context.Context, params *elastictranscoder.UpdatePipelineStatusInput, optFns ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePipelineStatus")
	}

	var r0 *elastictranscoder.UpdatePipelineStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineStatusInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.UpdatePipelineStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *elastictranscoder.UpdatePipelineStatusInput, ...func(*elastictranscoder.Options)) *elastictranscoder.UpdatePipelineStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastictranscoder.UpdatePipelineStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *elastictranscoder.UpdatePipelineStatusInput, ...func(*elastictranscoder.Options)) error); ok {
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
