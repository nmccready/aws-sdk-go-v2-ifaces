// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	bcmdataexports "github.com/aws/aws-sdk-go-v2/service/bcmdataexports"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateExport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateExport(ctx context.Context, params *bcmdataexports.CreateExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.CreateExportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateExport")
	}

	var r0 *bcmdataexports.CreateExportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.CreateExportInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.CreateExportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.CreateExportInput, ...func(*bcmdataexports.Options)) *bcmdataexports.CreateExportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.CreateExportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.CreateExportInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteExport(ctx context.Context, params *bcmdataexports.DeleteExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.DeleteExportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExport")
	}

	var r0 *bcmdataexports.DeleteExportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.DeleteExportInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.DeleteExportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.DeleteExportInput, ...func(*bcmdataexports.Options)) *bcmdataexports.DeleteExportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.DeleteExportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.DeleteExportInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecution provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExecution(ctx context.Context, params *bcmdataexports.GetExecutionInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExecutionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExecution")
	}

	var r0 *bcmdataexports.GetExecutionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetExecutionInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExecutionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetExecutionInput, ...func(*bcmdataexports.Options)) *bcmdataexports.GetExecutionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.GetExecutionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.GetExecutionInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExport(ctx context.Context, params *bcmdataexports.GetExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExport")
	}

	var r0 *bcmdataexports.GetExportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetExportInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetExportInput, ...func(*bcmdataexports.Options)) *bcmdataexports.GetExportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.GetExportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.GetExportInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTable provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTable(ctx context.Context, params *bcmdataexports.GetTableInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTable")
	}

	var r0 *bcmdataexports.GetTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetTableInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.GetTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.GetTableInput, ...func(*bcmdataexports.Options)) *bcmdataexports.GetTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.GetTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.GetTableInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExecutions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExecutions(ctx context.Context, params *bcmdataexports.ListExecutionsInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExecutionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExecutions")
	}

	var r0 *bcmdataexports.ListExecutionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListExecutionsInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExecutionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListExecutionsInput, ...func(*bcmdataexports.Options)) *bcmdataexports.ListExecutionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.ListExecutionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.ListExecutionsInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExports provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListExports(ctx context.Context, params *bcmdataexports.ListExportsInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExportsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListExports")
	}

	var r0 *bcmdataexports.ListExportsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListExportsInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExportsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListExportsInput, ...func(*bcmdataexports.Options)) *bcmdataexports.ListExportsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.ListExportsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.ListExportsInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTables provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTables(ctx context.Context, params *bcmdataexports.ListTablesInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTablesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTables")
	}

	var r0 *bcmdataexports.ListTablesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListTablesInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTablesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListTablesInput, ...func(*bcmdataexports.Options)) *bcmdataexports.ListTablesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.ListTablesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.ListTablesInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *bcmdataexports.ListTagsForResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTagsForResourceOutput, error) {
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

	var r0 *bcmdataexports.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListTagsForResourceInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.ListTagsForResourceInput, ...func(*bcmdataexports.Options)) *bcmdataexports.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.ListTagsForResourceInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() bcmdataexports.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bcmdataexports.Options
	if rf, ok := ret.Get(0).(func() bcmdataexports.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bcmdataexports.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *bcmdataexports.TagResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.TagResourceOutput, error) {
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

	var r0 *bcmdataexports.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.TagResourceInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.TagResourceInput, ...func(*bcmdataexports.Options)) *bcmdataexports.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.TagResourceInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *bcmdataexports.UntagResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.UntagResourceOutput, error) {
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

	var r0 *bcmdataexports.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.UntagResourceInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.UntagResourceInput, ...func(*bcmdataexports.Options)) *bcmdataexports.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.UntagResourceInput, ...func(*bcmdataexports.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExport provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateExport(ctx context.Context, params *bcmdataexports.UpdateExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.UpdateExportOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExport")
	}

	var r0 *bcmdataexports.UpdateExportOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.UpdateExportInput, ...func(*bcmdataexports.Options)) (*bcmdataexports.UpdateExportOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bcmdataexports.UpdateExportInput, ...func(*bcmdataexports.Options)) *bcmdataexports.UpdateExportOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bcmdataexports.UpdateExportOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bcmdataexports.UpdateExportInput, ...func(*bcmdataexports.Options)) error); ok {
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
