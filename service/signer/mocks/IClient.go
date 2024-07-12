// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	signer "github.com/aws/aws-sdk-go-v2/service/signer"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AddProfilePermission provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AddProfilePermission(ctx context.Context, params *signer.AddProfilePermissionInput, optFns ...func(*signer.Options)) (*signer.AddProfilePermissionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddProfilePermission")
	}

	var r0 *signer.AddProfilePermissionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.AddProfilePermissionInput, ...func(*signer.Options)) (*signer.AddProfilePermissionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.AddProfilePermissionInput, ...func(*signer.Options)) *signer.AddProfilePermissionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.AddProfilePermissionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.AddProfilePermissionInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelSigningProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelSigningProfile(ctx context.Context, params *signer.CancelSigningProfileInput, optFns ...func(*signer.Options)) (*signer.CancelSigningProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelSigningProfile")
	}

	var r0 *signer.CancelSigningProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.CancelSigningProfileInput, ...func(*signer.Options)) (*signer.CancelSigningProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.CancelSigningProfileInput, ...func(*signer.Options)) *signer.CancelSigningProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.CancelSigningProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.CancelSigningProfileInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSigningJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSigningJob(ctx context.Context, params *signer.DescribeSigningJobInput, optFns ...func(*signer.Options)) (*signer.DescribeSigningJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSigningJob")
	}

	var r0 *signer.DescribeSigningJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.DescribeSigningJobInput, ...func(*signer.Options)) (*signer.DescribeSigningJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.DescribeSigningJobInput, ...func(*signer.Options)) *signer.DescribeSigningJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.DescribeSigningJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.DescribeSigningJobInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRevocationStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRevocationStatus(ctx context.Context, params *signer.GetRevocationStatusInput, optFns ...func(*signer.Options)) (*signer.GetRevocationStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRevocationStatus")
	}

	var r0 *signer.GetRevocationStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetRevocationStatusInput, ...func(*signer.Options)) (*signer.GetRevocationStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetRevocationStatusInput, ...func(*signer.Options)) *signer.GetRevocationStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.GetRevocationStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.GetRevocationStatusInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSigningPlatform provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSigningPlatform(ctx context.Context, params *signer.GetSigningPlatformInput, optFns ...func(*signer.Options)) (*signer.GetSigningPlatformOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSigningPlatform")
	}

	var r0 *signer.GetSigningPlatformOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetSigningPlatformInput, ...func(*signer.Options)) (*signer.GetSigningPlatformOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetSigningPlatformInput, ...func(*signer.Options)) *signer.GetSigningPlatformOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.GetSigningPlatformOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.GetSigningPlatformInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSigningProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSigningProfile(ctx context.Context, params *signer.GetSigningProfileInput, optFns ...func(*signer.Options)) (*signer.GetSigningProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSigningProfile")
	}

	var r0 *signer.GetSigningProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetSigningProfileInput, ...func(*signer.Options)) (*signer.GetSigningProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.GetSigningProfileInput, ...func(*signer.Options)) *signer.GetSigningProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.GetSigningProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.GetSigningProfileInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProfilePermissions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProfilePermissions(ctx context.Context, params *signer.ListProfilePermissionsInput, optFns ...func(*signer.Options)) (*signer.ListProfilePermissionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProfilePermissions")
	}

	var r0 *signer.ListProfilePermissionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListProfilePermissionsInput, ...func(*signer.Options)) (*signer.ListProfilePermissionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListProfilePermissionsInput, ...func(*signer.Options)) *signer.ListProfilePermissionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.ListProfilePermissionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.ListProfilePermissionsInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSigningJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSigningJobs(ctx context.Context, params *signer.ListSigningJobsInput, optFns ...func(*signer.Options)) (*signer.ListSigningJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSigningJobs")
	}

	var r0 *signer.ListSigningJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningJobsInput, ...func(*signer.Options)) (*signer.ListSigningJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningJobsInput, ...func(*signer.Options)) *signer.ListSigningJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.ListSigningJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.ListSigningJobsInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSigningPlatforms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSigningPlatforms(ctx context.Context, params *signer.ListSigningPlatformsInput, optFns ...func(*signer.Options)) (*signer.ListSigningPlatformsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSigningPlatforms")
	}

	var r0 *signer.ListSigningPlatformsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningPlatformsInput, ...func(*signer.Options)) (*signer.ListSigningPlatformsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningPlatformsInput, ...func(*signer.Options)) *signer.ListSigningPlatformsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.ListSigningPlatformsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.ListSigningPlatformsInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSigningProfiles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSigningProfiles(ctx context.Context, params *signer.ListSigningProfilesInput, optFns ...func(*signer.Options)) (*signer.ListSigningProfilesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSigningProfiles")
	}

	var r0 *signer.ListSigningProfilesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningProfilesInput, ...func(*signer.Options)) (*signer.ListSigningProfilesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListSigningProfilesInput, ...func(*signer.Options)) *signer.ListSigningProfilesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.ListSigningProfilesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.ListSigningProfilesInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *signer.ListTagsForResourceInput, optFns ...func(*signer.Options)) (*signer.ListTagsForResourceOutput, error) {
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

	var r0 *signer.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListTagsForResourceInput, ...func(*signer.Options)) (*signer.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.ListTagsForResourceInput, ...func(*signer.Options)) *signer.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.ListTagsForResourceInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() signer.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 signer.Options
	if rf, ok := ret.Get(0).(func() signer.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(signer.Options)
	}

	return r0
}

// PutSigningProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutSigningProfile(ctx context.Context, params *signer.PutSigningProfileInput, optFns ...func(*signer.Options)) (*signer.PutSigningProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutSigningProfile")
	}

	var r0 *signer.PutSigningProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.PutSigningProfileInput, ...func(*signer.Options)) (*signer.PutSigningProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.PutSigningProfileInput, ...func(*signer.Options)) *signer.PutSigningProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.PutSigningProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.PutSigningProfileInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveProfilePermission provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RemoveProfilePermission(ctx context.Context, params *signer.RemoveProfilePermissionInput, optFns ...func(*signer.Options)) (*signer.RemoveProfilePermissionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProfilePermission")
	}

	var r0 *signer.RemoveProfilePermissionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RemoveProfilePermissionInput, ...func(*signer.Options)) (*signer.RemoveProfilePermissionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RemoveProfilePermissionInput, ...func(*signer.Options)) *signer.RemoveProfilePermissionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.RemoveProfilePermissionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.RemoveProfilePermissionInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeSignature provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RevokeSignature(ctx context.Context, params *signer.RevokeSignatureInput, optFns ...func(*signer.Options)) (*signer.RevokeSignatureOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RevokeSignature")
	}

	var r0 *signer.RevokeSignatureOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RevokeSignatureInput, ...func(*signer.Options)) (*signer.RevokeSignatureOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RevokeSignatureInput, ...func(*signer.Options)) *signer.RevokeSignatureOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.RevokeSignatureOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.RevokeSignatureInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeSigningProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RevokeSigningProfile(ctx context.Context, params *signer.RevokeSigningProfileInput, optFns ...func(*signer.Options)) (*signer.RevokeSigningProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RevokeSigningProfile")
	}

	var r0 *signer.RevokeSigningProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RevokeSigningProfileInput, ...func(*signer.Options)) (*signer.RevokeSigningProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.RevokeSigningProfileInput, ...func(*signer.Options)) *signer.RevokeSigningProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.RevokeSigningProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.RevokeSigningProfileInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignPayload provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SignPayload(ctx context.Context, params *signer.SignPayloadInput, optFns ...func(*signer.Options)) (*signer.SignPayloadOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SignPayload")
	}

	var r0 *signer.SignPayloadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.SignPayloadInput, ...func(*signer.Options)) (*signer.SignPayloadOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.SignPayloadInput, ...func(*signer.Options)) *signer.SignPayloadOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.SignPayloadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.SignPayloadInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartSigningJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartSigningJob(ctx context.Context, params *signer.StartSigningJobInput, optFns ...func(*signer.Options)) (*signer.StartSigningJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartSigningJob")
	}

	var r0 *signer.StartSigningJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.StartSigningJobInput, ...func(*signer.Options)) (*signer.StartSigningJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.StartSigningJobInput, ...func(*signer.Options)) *signer.StartSigningJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.StartSigningJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.StartSigningJobInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *signer.TagResourceInput, optFns ...func(*signer.Options)) (*signer.TagResourceOutput, error) {
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

	var r0 *signer.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.TagResourceInput, ...func(*signer.Options)) (*signer.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.TagResourceInput, ...func(*signer.Options)) *signer.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.TagResourceInput, ...func(*signer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *signer.UntagResourceInput, optFns ...func(*signer.Options)) (*signer.UntagResourceOutput, error) {
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

	var r0 *signer.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *signer.UntagResourceInput, ...func(*signer.Options)) (*signer.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *signer.UntagResourceInput, ...func(*signer.Options)) *signer.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*signer.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *signer.UntagResourceInput, ...func(*signer.Options)) error); ok {
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