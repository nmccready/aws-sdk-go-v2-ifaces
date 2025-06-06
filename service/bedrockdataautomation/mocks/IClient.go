// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	bedrockdataautomation "github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateBlueprint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateBlueprint(ctx context.Context, params *bedrockdataautomation.CreateBlueprintInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateBlueprintOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlueprint")
	}

	var r0 *bedrockdataautomation.CreateBlueprintOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateBlueprintInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateBlueprintOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateBlueprintInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.CreateBlueprintOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.CreateBlueprintOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.CreateBlueprintInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBlueprintVersion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateBlueprintVersion(ctx context.Context, params *bedrockdataautomation.CreateBlueprintVersionInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateBlueprintVersionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlueprintVersion")
	}

	var r0 *bedrockdataautomation.CreateBlueprintVersionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateBlueprintVersionInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateBlueprintVersionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateBlueprintVersionInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.CreateBlueprintVersionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.CreateBlueprintVersionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.CreateBlueprintVersionInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataAutomationProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDataAutomationProject(ctx context.Context, params *bedrockdataautomation.CreateDataAutomationProjectInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateDataAutomationProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDataAutomationProject")
	}

	var r0 *bedrockdataautomation.CreateDataAutomationProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.CreateDataAutomationProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.CreateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.CreateDataAutomationProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.CreateDataAutomationProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.CreateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlueprint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteBlueprint(ctx context.Context, params *bedrockdataautomation.DeleteBlueprintInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.DeleteBlueprintOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlueprint")
	}

	var r0 *bedrockdataautomation.DeleteBlueprintOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.DeleteBlueprintInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.DeleteBlueprintOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.DeleteBlueprintInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.DeleteBlueprintOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.DeleteBlueprintOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.DeleteBlueprintInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDataAutomationProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDataAutomationProject(ctx context.Context, params *bedrockdataautomation.DeleteDataAutomationProjectInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.DeleteDataAutomationProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDataAutomationProject")
	}

	var r0 *bedrockdataautomation.DeleteDataAutomationProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.DeleteDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.DeleteDataAutomationProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.DeleteDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.DeleteDataAutomationProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.DeleteDataAutomationProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.DeleteDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlueprint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetBlueprint(ctx context.Context, params *bedrockdataautomation.GetBlueprintInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.GetBlueprintOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBlueprint")
	}

	var r0 *bedrockdataautomation.GetBlueprintOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.GetBlueprintInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.GetBlueprintOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.GetBlueprintInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.GetBlueprintOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.GetBlueprintOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.GetBlueprintInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataAutomationProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDataAutomationProject(ctx context.Context, params *bedrockdataautomation.GetDataAutomationProjectInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.GetDataAutomationProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDataAutomationProject")
	}

	var r0 *bedrockdataautomation.GetDataAutomationProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.GetDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.GetDataAutomationProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.GetDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.GetDataAutomationProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.GetDataAutomationProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.GetDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBlueprints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListBlueprints(ctx context.Context, params *bedrockdataautomation.ListBlueprintsInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListBlueprintsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListBlueprints")
	}

	var r0 *bedrockdataautomation.ListBlueprintsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListBlueprintsInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListBlueprintsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListBlueprintsInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.ListBlueprintsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.ListBlueprintsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.ListBlueprintsInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDataAutomationProjects provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDataAutomationProjects(ctx context.Context, params *bedrockdataautomation.ListDataAutomationProjectsInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListDataAutomationProjectsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDataAutomationProjects")
	}

	var r0 *bedrockdataautomation.ListDataAutomationProjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListDataAutomationProjectsInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListDataAutomationProjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListDataAutomationProjectsInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.ListDataAutomationProjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.ListDataAutomationProjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.ListDataAutomationProjectsInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *bedrockdataautomation.ListTagsForResourceInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListTagsForResourceOutput, error) {
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

	var r0 *bedrockdataautomation.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListTagsForResourceInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.ListTagsForResourceInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.ListTagsForResourceInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() bedrockdataautomation.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bedrockdataautomation.Options
	if rf, ok := ret.Get(0).(func() bedrockdataautomation.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bedrockdataautomation.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *bedrockdataautomation.TagResourceInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.TagResourceOutput, error) {
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

	var r0 *bedrockdataautomation.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.TagResourceInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.TagResourceInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.TagResourceInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *bedrockdataautomation.UntagResourceInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UntagResourceOutput, error) {
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

	var r0 *bedrockdataautomation.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UntagResourceInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UntagResourceInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.UntagResourceInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlueprint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateBlueprint(ctx context.Context, params *bedrockdataautomation.UpdateBlueprintInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UpdateBlueprintOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlueprint")
	}

	var r0 *bedrockdataautomation.UpdateBlueprintOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UpdateBlueprintInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UpdateBlueprintOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UpdateBlueprintInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.UpdateBlueprintOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.UpdateBlueprintOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.UpdateBlueprintInput, ...func(*bedrockdataautomation.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDataAutomationProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateDataAutomationProject(ctx context.Context, params *bedrockdataautomation.UpdateDataAutomationProjectInput, optFns ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UpdateDataAutomationProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataAutomationProject")
	}

	var r0 *bedrockdataautomation.UpdateDataAutomationProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UpdateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) (*bedrockdataautomation.UpdateDataAutomationProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockdataautomation.UpdateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) *bedrockdataautomation.UpdateDataAutomationProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockdataautomation.UpdateDataAutomationProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockdataautomation.UpdateDataAutomationProjectInput, ...func(*bedrockdataautomation.Options)) error); ok {
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
