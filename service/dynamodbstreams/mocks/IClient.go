// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dynamodbstreams "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DescribeStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeStream(ctx context.Context, params *dynamodbstreams.DescribeStreamInput, optFns ...func(*dynamodbstreams.Options)) (*dynamodbstreams.DescribeStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeStream")
	}

	var r0 *dynamodbstreams.DescribeStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.DescribeStreamInput, ...func(*dynamodbstreams.Options)) (*dynamodbstreams.DescribeStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.DescribeStreamInput, ...func(*dynamodbstreams.Options)) *dynamodbstreams.DescribeStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodbstreams.DescribeStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodbstreams.DescribeStreamInput, ...func(*dynamodbstreams.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecords provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRecords(ctx context.Context, params *dynamodbstreams.GetRecordsInput, optFns ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetRecordsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRecords")
	}

	var r0 *dynamodbstreams.GetRecordsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.GetRecordsInput, ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetRecordsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.GetRecordsInput, ...func(*dynamodbstreams.Options)) *dynamodbstreams.GetRecordsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodbstreams.GetRecordsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodbstreams.GetRecordsInput, ...func(*dynamodbstreams.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShardIterator provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetShardIterator(ctx context.Context, params *dynamodbstreams.GetShardIteratorInput, optFns ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetShardIteratorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetShardIterator")
	}

	var r0 *dynamodbstreams.GetShardIteratorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.GetShardIteratorInput, ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetShardIteratorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.GetShardIteratorInput, ...func(*dynamodbstreams.Options)) *dynamodbstreams.GetShardIteratorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodbstreams.GetShardIteratorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodbstreams.GetShardIteratorInput, ...func(*dynamodbstreams.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStreams provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListStreams(ctx context.Context, params *dynamodbstreams.ListStreamsInput, optFns ...func(*dynamodbstreams.Options)) (*dynamodbstreams.ListStreamsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListStreams")
	}

	var r0 *dynamodbstreams.ListStreamsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.ListStreamsInput, ...func(*dynamodbstreams.Options)) (*dynamodbstreams.ListStreamsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodbstreams.ListStreamsInput, ...func(*dynamodbstreams.Options)) *dynamodbstreams.ListStreamsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodbstreams.ListStreamsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodbstreams.ListStreamsInput, ...func(*dynamodbstreams.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() dynamodbstreams.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 dynamodbstreams.Options
	if rf, ok := ret.Get(0).(func() dynamodbstreams.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dynamodbstreams.Options)
	}

	return r0
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