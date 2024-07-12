// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	synthetics "github.com/aws/aws-sdk-go-v2/service/synthetics"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateResource(ctx context.Context, params *synthetics.AssociateResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.AssociateResourceOutput, error) {
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

	var r0 *synthetics.AssociateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.AssociateResourceInput, ...func(*synthetics.Options)) (*synthetics.AssociateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.AssociateResourceInput, ...func(*synthetics.Options)) *synthetics.AssociateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.AssociateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.AssociateResourceInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCanary(ctx context.Context, params *synthetics.CreateCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.CreateCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCanary")
	}

	var r0 *synthetics.CreateCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.CreateCanaryInput, ...func(*synthetics.Options)) (*synthetics.CreateCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.CreateCanaryInput, ...func(*synthetics.Options)) *synthetics.CreateCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.CreateCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.CreateCanaryInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateGroup(ctx context.Context, params *synthetics.CreateGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.CreateGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateGroup")
	}

	var r0 *synthetics.CreateGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.CreateGroupInput, ...func(*synthetics.Options)) (*synthetics.CreateGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.CreateGroupInput, ...func(*synthetics.Options)) *synthetics.CreateGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.CreateGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.CreateGroupInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCanary(ctx context.Context, params *synthetics.DeleteCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.DeleteCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCanary")
	}

	var r0 *synthetics.DeleteCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DeleteCanaryInput, ...func(*synthetics.Options)) (*synthetics.DeleteCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DeleteCanaryInput, ...func(*synthetics.Options)) *synthetics.DeleteCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DeleteCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DeleteCanaryInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteGroup(ctx context.Context, params *synthetics.DeleteGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.DeleteGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroup")
	}

	var r0 *synthetics.DeleteGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DeleteGroupInput, ...func(*synthetics.Options)) (*synthetics.DeleteGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DeleteGroupInput, ...func(*synthetics.Options)) *synthetics.DeleteGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DeleteGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DeleteGroupInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCanaries provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeCanaries(ctx context.Context, params *synthetics.DescribeCanariesInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeCanariesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCanaries")
	}

	var r0 *synthetics.DescribeCanariesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeCanariesInput, ...func(*synthetics.Options)) (*synthetics.DescribeCanariesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeCanariesInput, ...func(*synthetics.Options)) *synthetics.DescribeCanariesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DescribeCanariesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DescribeCanariesInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCanariesLastRun provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeCanariesLastRun(ctx context.Context, params *synthetics.DescribeCanariesLastRunInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeCanariesLastRunOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCanariesLastRun")
	}

	var r0 *synthetics.DescribeCanariesLastRunOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeCanariesLastRunInput, ...func(*synthetics.Options)) (*synthetics.DescribeCanariesLastRunOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeCanariesLastRunInput, ...func(*synthetics.Options)) *synthetics.DescribeCanariesLastRunOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DescribeCanariesLastRunOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DescribeCanariesLastRunInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRuntimeVersions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeRuntimeVersions(ctx context.Context, params *synthetics.DescribeRuntimeVersionsInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRuntimeVersions")
	}

	var r0 *synthetics.DescribeRuntimeVersionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeRuntimeVersionsInput, ...func(*synthetics.Options)) (*synthetics.DescribeRuntimeVersionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DescribeRuntimeVersionsInput, ...func(*synthetics.Options)) *synthetics.DescribeRuntimeVersionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DescribeRuntimeVersionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DescribeRuntimeVersionsInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateResource(ctx context.Context, params *synthetics.DisassociateResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.DisassociateResourceOutput, error) {
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

	var r0 *synthetics.DisassociateResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DisassociateResourceInput, ...func(*synthetics.Options)) (*synthetics.DisassociateResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.DisassociateResourceInput, ...func(*synthetics.Options)) *synthetics.DisassociateResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.DisassociateResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.DisassociateResourceInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCanary(ctx context.Context, params *synthetics.GetCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.GetCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCanary")
	}

	var r0 *synthetics.GetCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetCanaryInput, ...func(*synthetics.Options)) (*synthetics.GetCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetCanaryInput, ...func(*synthetics.Options)) *synthetics.GetCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.GetCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.GetCanaryInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCanaryRuns provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCanaryRuns(ctx context.Context, params *synthetics.GetCanaryRunsInput, optFns ...func(*synthetics.Options)) (*synthetics.GetCanaryRunsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCanaryRuns")
	}

	var r0 *synthetics.GetCanaryRunsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetCanaryRunsInput, ...func(*synthetics.Options)) (*synthetics.GetCanaryRunsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetCanaryRunsInput, ...func(*synthetics.Options)) *synthetics.GetCanaryRunsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.GetCanaryRunsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.GetCanaryRunsInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroup(ctx context.Context, params *synthetics.GetGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.GetGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroup")
	}

	var r0 *synthetics.GetGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetGroupInput, ...func(*synthetics.Options)) (*synthetics.GetGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.GetGroupInput, ...func(*synthetics.Options)) *synthetics.GetGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.GetGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.GetGroupInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAssociatedGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAssociatedGroups(ctx context.Context, params *synthetics.ListAssociatedGroupsInput, optFns ...func(*synthetics.Options)) (*synthetics.ListAssociatedGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAssociatedGroups")
	}

	var r0 *synthetics.ListAssociatedGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListAssociatedGroupsInput, ...func(*synthetics.Options)) (*synthetics.ListAssociatedGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListAssociatedGroupsInput, ...func(*synthetics.Options)) *synthetics.ListAssociatedGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.ListAssociatedGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.ListAssociatedGroupsInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroupResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroupResources(ctx context.Context, params *synthetics.ListGroupResourcesInput, optFns ...func(*synthetics.Options)) (*synthetics.ListGroupResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupResources")
	}

	var r0 *synthetics.ListGroupResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListGroupResourcesInput, ...func(*synthetics.Options)) (*synthetics.ListGroupResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListGroupResourcesInput, ...func(*synthetics.Options)) *synthetics.ListGroupResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.ListGroupResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.ListGroupResourcesInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroups(ctx context.Context, params *synthetics.ListGroupsInput, optFns ...func(*synthetics.Options)) (*synthetics.ListGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroups")
	}

	var r0 *synthetics.ListGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListGroupsInput, ...func(*synthetics.Options)) (*synthetics.ListGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListGroupsInput, ...func(*synthetics.Options)) *synthetics.ListGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.ListGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.ListGroupsInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *synthetics.ListTagsForResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.ListTagsForResourceOutput, error) {
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

	var r0 *synthetics.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListTagsForResourceInput, ...func(*synthetics.Options)) (*synthetics.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.ListTagsForResourceInput, ...func(*synthetics.Options)) *synthetics.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.ListTagsForResourceInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() synthetics.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 synthetics.Options
	if rf, ok := ret.Get(0).(func() synthetics.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(synthetics.Options)
	}

	return r0
}

// StartCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartCanary(ctx context.Context, params *synthetics.StartCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.StartCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartCanary")
	}

	var r0 *synthetics.StartCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.StartCanaryInput, ...func(*synthetics.Options)) (*synthetics.StartCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.StartCanaryInput, ...func(*synthetics.Options)) *synthetics.StartCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.StartCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.StartCanaryInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopCanary(ctx context.Context, params *synthetics.StopCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.StopCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopCanary")
	}

	var r0 *synthetics.StopCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.StopCanaryInput, ...func(*synthetics.Options)) (*synthetics.StopCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.StopCanaryInput, ...func(*synthetics.Options)) *synthetics.StopCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.StopCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.StopCanaryInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *synthetics.TagResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.TagResourceOutput, error) {
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

	var r0 *synthetics.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.TagResourceInput, ...func(*synthetics.Options)) (*synthetics.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.TagResourceInput, ...func(*synthetics.Options)) *synthetics.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.TagResourceInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *synthetics.UntagResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.UntagResourceOutput, error) {
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

	var r0 *synthetics.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.UntagResourceInput, ...func(*synthetics.Options)) (*synthetics.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.UntagResourceInput, ...func(*synthetics.Options)) *synthetics.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.UntagResourceInput, ...func(*synthetics.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCanary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCanary(ctx context.Context, params *synthetics.UpdateCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.UpdateCanaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCanary")
	}

	var r0 *synthetics.UpdateCanaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.UpdateCanaryInput, ...func(*synthetics.Options)) (*synthetics.UpdateCanaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *synthetics.UpdateCanaryInput, ...func(*synthetics.Options)) *synthetics.UpdateCanaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synthetics.UpdateCanaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *synthetics.UpdateCanaryInput, ...func(*synthetics.Options)) error); ok {
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