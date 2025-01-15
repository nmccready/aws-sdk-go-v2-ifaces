// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	iot1clickprojects "github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateDeviceWithPlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateDeviceWithPlacement(ctx context.Context, params *iot1clickprojects.AssociateDeviceWithPlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateDeviceWithPlacement")
	}

	var r0 *iot1clickprojects.AssociateDeviceWithPlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.AssociateDeviceWithPlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.AssociateDeviceWithPlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.AssociateDeviceWithPlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.AssociateDeviceWithPlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.AssociateDeviceWithPlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreatePlacement(ctx context.Context, params *iot1clickprojects.CreatePlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.CreatePlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePlacement")
	}

	var r0 *iot1clickprojects.CreatePlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.CreatePlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.CreatePlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.CreatePlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.CreatePlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.CreatePlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.CreatePlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProject(ctx context.Context, params *iot1clickprojects.CreateProjectInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.CreateProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 *iot1clickprojects.CreateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.CreateProjectInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.CreateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.CreateProjectInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.CreateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.CreateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.CreateProjectInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeletePlacement(ctx context.Context, params *iot1clickprojects.DeletePlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DeletePlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePlacement")
	}

	var r0 *iot1clickprojects.DeletePlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DeletePlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DeletePlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DeletePlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.DeletePlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.DeletePlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.DeletePlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProject(ctx context.Context, params *iot1clickprojects.DeleteProjectInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DeleteProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProject")
	}

	var r0 *iot1clickprojects.DeleteProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DeleteProjectInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DeleteProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DeleteProjectInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.DeleteProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.DeleteProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.DeleteProjectInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribePlacement(ctx context.Context, params *iot1clickprojects.DescribePlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DescribePlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribePlacement")
	}

	var r0 *iot1clickprojects.DescribePlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DescribePlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DescribePlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DescribePlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.DescribePlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.DescribePlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.DescribePlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeProject(ctx context.Context, params *iot1clickprojects.DescribeProjectInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DescribeProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeProject")
	}

	var r0 *iot1clickprojects.DescribeProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DescribeProjectInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DescribeProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DescribeProjectInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.DescribeProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.DescribeProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.DescribeProjectInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateDeviceFromPlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateDeviceFromPlacement(ctx context.Context, params *iot1clickprojects.DisassociateDeviceFromPlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateDeviceFromPlacement")
	}

	var r0 *iot1clickprojects.DisassociateDeviceFromPlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DisassociateDeviceFromPlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.DisassociateDeviceFromPlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.DisassociateDeviceFromPlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.DisassociateDeviceFromPlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.DisassociateDeviceFromPlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesInPlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDevicesInPlacement(ctx context.Context, params *iot1clickprojects.GetDevicesInPlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDevicesInPlacement")
	}

	var r0 *iot1clickprojects.GetDevicesInPlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.GetDevicesInPlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.GetDevicesInPlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.GetDevicesInPlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.GetDevicesInPlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.GetDevicesInPlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.GetDevicesInPlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlacements provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPlacements(ctx context.Context, params *iot1clickprojects.ListPlacementsInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListPlacementsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPlacements")
	}

	var r0 *iot1clickprojects.ListPlacementsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListPlacementsInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListPlacementsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListPlacementsInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.ListPlacementsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.ListPlacementsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.ListPlacementsInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProjects(ctx context.Context, params *iot1clickprojects.ListProjectsInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListProjectsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjects")
	}

	var r0 *iot1clickprojects.ListProjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListProjectsInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListProjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListProjectsInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.ListProjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.ListProjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.ListProjectsInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *iot1clickprojects.ListTagsForResourceInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListTagsForResourceOutput, error) {
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

	var r0 *iot1clickprojects.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListTagsForResourceInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.ListTagsForResourceInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.ListTagsForResourceInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() iot1clickprojects.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 iot1clickprojects.Options
	if rf, ok := ret.Get(0).(func() iot1clickprojects.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(iot1clickprojects.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *iot1clickprojects.TagResourceInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.TagResourceOutput, error) {
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

	var r0 *iot1clickprojects.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.TagResourceInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.TagResourceInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.TagResourceInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *iot1clickprojects.UntagResourceInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UntagResourceOutput, error) {
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

	var r0 *iot1clickprojects.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UntagResourceInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UntagResourceInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.UntagResourceInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlacement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdatePlacement(ctx context.Context, params *iot1clickprojects.UpdatePlacementInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UpdatePlacementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePlacement")
	}

	var r0 *iot1clickprojects.UpdatePlacementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UpdatePlacementInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UpdatePlacementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UpdatePlacementInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.UpdatePlacementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.UpdatePlacementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.UpdatePlacementInput, ...func(*iot1clickprojects.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateProject(ctx context.Context, params *iot1clickprojects.UpdateProjectInput, optFns ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UpdateProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProject")
	}

	var r0 *iot1clickprojects.UpdateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UpdateProjectInput, ...func(*iot1clickprojects.Options)) (*iot1clickprojects.UpdateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickprojects.UpdateProjectInput, ...func(*iot1clickprojects.Options)) *iot1clickprojects.UpdateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickprojects.UpdateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickprojects.UpdateProjectInput, ...func(*iot1clickprojects.Options)) error); ok {
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
