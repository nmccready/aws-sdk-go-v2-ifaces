// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	kinesisvideoarchivedmedia "github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// GetClip provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetClip(ctx context.Context, params *kinesisvideoarchivedmedia.GetClipInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetClip")
	}

	var r0 *kinesisvideoarchivedmedia.GetClipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetClipInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetClipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetClipInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.GetClipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.GetClipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.GetClipInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDASHStreamingSessionURL provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDASHStreamingSessionURL(ctx context.Context, params *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDASHStreamingSessionURL")
	}

	var r0 *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHLSStreamingSessionURL provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetHLSStreamingSessionURL(ctx context.Context, params *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetHLSStreamingSessionURL")
	}

	var r0 *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImages provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetImages(ctx context.Context, params *kinesisvideoarchivedmedia.GetImagesInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetImagesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetImages")
	}

	var r0 *kinesisvideoarchivedmedia.GetImagesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetImagesInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetImagesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetImagesInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.GetImagesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.GetImagesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.GetImagesInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMediaForFragmentList provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMediaForFragmentList(ctx context.Context, params *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMediaForFragmentList")
	}

	var r0 *kinesisvideoarchivedmedia.GetMediaForFragmentListOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.GetMediaForFragmentListOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFragments provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListFragments(ctx context.Context, params *kinesisvideoarchivedmedia.ListFragmentsInput, optFns ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFragments")
	}

	var r0 *kinesisvideoarchivedmedia.ListFragmentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.ListFragmentsInput, ...func(*kinesisvideoarchivedmedia.Options)) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinesisvideoarchivedmedia.ListFragmentsInput, ...func(*kinesisvideoarchivedmedia.Options)) *kinesisvideoarchivedmedia.ListFragmentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinesisvideoarchivedmedia.ListFragmentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinesisvideoarchivedmedia.ListFragmentsInput, ...func(*kinesisvideoarchivedmedia.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() kinesisvideoarchivedmedia.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 kinesisvideoarchivedmedia.Options
	if rf, ok := ret.Get(0).(func() kinesisvideoarchivedmedia.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(kinesisvideoarchivedmedia.Options)
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
