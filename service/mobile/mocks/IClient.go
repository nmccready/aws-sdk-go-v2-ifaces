// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mobile "github.com/aws/aws-sdk-go-v2/service/mobile"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProject(ctx context.Context, params *mobile.CreateProjectInput, optFns ...func(*mobile.Options)) (*mobile.CreateProjectOutput, error) {
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

	var r0 *mobile.CreateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.CreateProjectInput, ...func(*mobile.Options)) (*mobile.CreateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.CreateProjectInput, ...func(*mobile.Options)) *mobile.CreateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.CreateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.CreateProjectInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProject(ctx context.Context, params *mobile.DeleteProjectInput, optFns ...func(*mobile.Options)) (*mobile.DeleteProjectOutput, error) {
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

	var r0 *mobile.DeleteProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DeleteProjectInput, ...func(*mobile.Options)) (*mobile.DeleteProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DeleteProjectInput, ...func(*mobile.Options)) *mobile.DeleteProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.DeleteProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.DeleteProjectInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeBundle provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeBundle(ctx context.Context, params *mobile.DescribeBundleInput, optFns ...func(*mobile.Options)) (*mobile.DescribeBundleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeBundle")
	}

	var r0 *mobile.DescribeBundleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DescribeBundleInput, ...func(*mobile.Options)) (*mobile.DescribeBundleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DescribeBundleInput, ...func(*mobile.Options)) *mobile.DescribeBundleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.DescribeBundleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.DescribeBundleInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeProject(ctx context.Context, params *mobile.DescribeProjectInput, optFns ...func(*mobile.Options)) (*mobile.DescribeProjectOutput, error) {
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

	var r0 *mobile.DescribeProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DescribeProjectInput, ...func(*mobile.Options)) (*mobile.DescribeProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.DescribeProjectInput, ...func(*mobile.Options)) *mobile.DescribeProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.DescribeProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.DescribeProjectInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportBundle provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExportBundle(ctx context.Context, params *mobile.ExportBundleInput, optFns ...func(*mobile.Options)) (*mobile.ExportBundleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportBundle")
	}

	var r0 *mobile.ExportBundleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ExportBundleInput, ...func(*mobile.Options)) (*mobile.ExportBundleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ExportBundleInput, ...func(*mobile.Options)) *mobile.ExportBundleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.ExportBundleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.ExportBundleInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ExportProject(ctx context.Context, params *mobile.ExportProjectInput, optFns ...func(*mobile.Options)) (*mobile.ExportProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExportProject")
	}

	var r0 *mobile.ExportProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ExportProjectInput, ...func(*mobile.Options)) (*mobile.ExportProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ExportProjectInput, ...func(*mobile.Options)) *mobile.ExportProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.ExportProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.ExportProjectInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBundles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListBundles(ctx context.Context, params *mobile.ListBundlesInput, optFns ...func(*mobile.Options)) (*mobile.ListBundlesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListBundles")
	}

	var r0 *mobile.ListBundlesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ListBundlesInput, ...func(*mobile.Options)) (*mobile.ListBundlesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ListBundlesInput, ...func(*mobile.Options)) *mobile.ListBundlesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.ListBundlesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.ListBundlesInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProjects(ctx context.Context, params *mobile.ListProjectsInput, optFns ...func(*mobile.Options)) (*mobile.ListProjectsOutput, error) {
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

	var r0 *mobile.ListProjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ListProjectsInput, ...func(*mobile.Options)) (*mobile.ListProjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.ListProjectsInput, ...func(*mobile.Options)) *mobile.ListProjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.ListProjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.ListProjectsInput, ...func(*mobile.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() mobile.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 mobile.Options
	if rf, ok := ret.Get(0).(func() mobile.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mobile.Options)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateProject(ctx context.Context, params *mobile.UpdateProjectInput, optFns ...func(*mobile.Options)) (*mobile.UpdateProjectOutput, error) {
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

	var r0 *mobile.UpdateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.UpdateProjectInput, ...func(*mobile.Options)) (*mobile.UpdateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mobile.UpdateProjectInput, ...func(*mobile.Options)) *mobile.UpdateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mobile.UpdateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mobile.UpdateProjectInput, ...func(*mobile.Options)) error); ok {
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
