// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mediastoredata "github.com/aws/aws-sdk-go-v2/service/mediastoredata"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DeleteObject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteObject(ctx context.Context, params *mediastoredata.DeleteObjectInput, optFns ...func(*mediastoredata.Options)) (*mediastoredata.DeleteObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteObject")
	}

	var r0 *mediastoredata.DeleteObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.DeleteObjectInput, ...func(*mediastoredata.Options)) (*mediastoredata.DeleteObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.DeleteObjectInput, ...func(*mediastoredata.Options)) *mediastoredata.DeleteObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastoredata.DeleteObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastoredata.DeleteObjectInput, ...func(*mediastoredata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeObject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeObject(ctx context.Context, params *mediastoredata.DescribeObjectInput, optFns ...func(*mediastoredata.Options)) (*mediastoredata.DescribeObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeObject")
	}

	var r0 *mediastoredata.DescribeObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.DescribeObjectInput, ...func(*mediastoredata.Options)) (*mediastoredata.DescribeObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.DescribeObjectInput, ...func(*mediastoredata.Options)) *mediastoredata.DescribeObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastoredata.DescribeObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastoredata.DescribeObjectInput, ...func(*mediastoredata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetObject(ctx context.Context, params *mediastoredata.GetObjectInput, optFns ...func(*mediastoredata.Options)) (*mediastoredata.GetObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 *mediastoredata.GetObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.GetObjectInput, ...func(*mediastoredata.Options)) (*mediastoredata.GetObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.GetObjectInput, ...func(*mediastoredata.Options)) *mediastoredata.GetObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastoredata.GetObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastoredata.GetObjectInput, ...func(*mediastoredata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListItems provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListItems(ctx context.Context, params *mediastoredata.ListItemsInput, optFns ...func(*mediastoredata.Options)) (*mediastoredata.ListItemsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListItems")
	}

	var r0 *mediastoredata.ListItemsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.ListItemsInput, ...func(*mediastoredata.Options)) (*mediastoredata.ListItemsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.ListItemsInput, ...func(*mediastoredata.Options)) *mediastoredata.ListItemsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastoredata.ListItemsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastoredata.ListItemsInput, ...func(*mediastoredata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() mediastoredata.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 mediastoredata.Options
	if rf, ok := ret.Get(0).(func() mediastoredata.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mediastoredata.Options)
	}

	return r0
}

// PutObject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutObject(ctx context.Context, params *mediastoredata.PutObjectInput, optFns ...func(*mediastoredata.Options)) (*mediastoredata.PutObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutObject")
	}

	var r0 *mediastoredata.PutObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.PutObjectInput, ...func(*mediastoredata.Options)) (*mediastoredata.PutObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastoredata.PutObjectInput, ...func(*mediastoredata.Options)) *mediastoredata.PutObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastoredata.PutObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastoredata.PutObjectInput, ...func(*mediastoredata.Options)) error); ok {
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
