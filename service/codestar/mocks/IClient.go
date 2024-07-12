// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	codestar "github.com/aws/aws-sdk-go-v2/service/codestar"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateTeamMember provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateTeamMember(ctx context.Context, params *codestar.AssociateTeamMemberInput, optFns ...func(*codestar.Options)) (*codestar.AssociateTeamMemberOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateTeamMember")
	}

	var r0 *codestar.AssociateTeamMemberOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.AssociateTeamMemberInput, ...func(*codestar.Options)) (*codestar.AssociateTeamMemberOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.AssociateTeamMemberInput, ...func(*codestar.Options)) *codestar.AssociateTeamMemberOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.AssociateTeamMemberOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.AssociateTeamMemberInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProject(ctx context.Context, params *codestar.CreateProjectInput, optFns ...func(*codestar.Options)) (*codestar.CreateProjectOutput, error) {
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

	var r0 *codestar.CreateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.CreateProjectInput, ...func(*codestar.Options)) (*codestar.CreateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.CreateProjectInput, ...func(*codestar.Options)) *codestar.CreateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.CreateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.CreateProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateUserProfile(ctx context.Context, params *codestar.CreateUserProfileInput, optFns ...func(*codestar.Options)) (*codestar.CreateUserProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserProfile")
	}

	var r0 *codestar.CreateUserProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.CreateUserProfileInput, ...func(*codestar.Options)) (*codestar.CreateUserProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.CreateUserProfileInput, ...func(*codestar.Options)) *codestar.CreateUserProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.CreateUserProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.CreateUserProfileInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProject(ctx context.Context, params *codestar.DeleteProjectInput, optFns ...func(*codestar.Options)) (*codestar.DeleteProjectOutput, error) {
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

	var r0 *codestar.DeleteProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DeleteProjectInput, ...func(*codestar.Options)) (*codestar.DeleteProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DeleteProjectInput, ...func(*codestar.Options)) *codestar.DeleteProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.DeleteProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.DeleteProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteUserProfile(ctx context.Context, params *codestar.DeleteUserProfileInput, optFns ...func(*codestar.Options)) (*codestar.DeleteUserProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserProfile")
	}

	var r0 *codestar.DeleteUserProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DeleteUserProfileInput, ...func(*codestar.Options)) (*codestar.DeleteUserProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DeleteUserProfileInput, ...func(*codestar.Options)) *codestar.DeleteUserProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.DeleteUserProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.DeleteUserProfileInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeProject(ctx context.Context, params *codestar.DescribeProjectInput, optFns ...func(*codestar.Options)) (*codestar.DescribeProjectOutput, error) {
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

	var r0 *codestar.DescribeProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DescribeProjectInput, ...func(*codestar.Options)) (*codestar.DescribeProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DescribeProjectInput, ...func(*codestar.Options)) *codestar.DescribeProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.DescribeProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.DescribeProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeUserProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeUserProfile(ctx context.Context, params *codestar.DescribeUserProfileInput, optFns ...func(*codestar.Options)) (*codestar.DescribeUserProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeUserProfile")
	}

	var r0 *codestar.DescribeUserProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DescribeUserProfileInput, ...func(*codestar.Options)) (*codestar.DescribeUserProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DescribeUserProfileInput, ...func(*codestar.Options)) *codestar.DescribeUserProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.DescribeUserProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.DescribeUserProfileInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateTeamMember provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateTeamMember(ctx context.Context, params *codestar.DisassociateTeamMemberInput, optFns ...func(*codestar.Options)) (*codestar.DisassociateTeamMemberOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateTeamMember")
	}

	var r0 *codestar.DisassociateTeamMemberOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DisassociateTeamMemberInput, ...func(*codestar.Options)) (*codestar.DisassociateTeamMemberOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.DisassociateTeamMemberInput, ...func(*codestar.Options)) *codestar.DisassociateTeamMemberOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.DisassociateTeamMemberOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.DisassociateTeamMemberInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProjects(ctx context.Context, params *codestar.ListProjectsInput, optFns ...func(*codestar.Options)) (*codestar.ListProjectsOutput, error) {
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

	var r0 *codestar.ListProjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListProjectsInput, ...func(*codestar.Options)) (*codestar.ListProjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListProjectsInput, ...func(*codestar.Options)) *codestar.ListProjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.ListProjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.ListProjectsInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListResources(ctx context.Context, params *codestar.ListResourcesInput, optFns ...func(*codestar.Options)) (*codestar.ListResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListResources")
	}

	var r0 *codestar.ListResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListResourcesInput, ...func(*codestar.Options)) (*codestar.ListResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListResourcesInput, ...func(*codestar.Options)) *codestar.ListResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.ListResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.ListResourcesInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForProject(ctx context.Context, params *codestar.ListTagsForProjectInput, optFns ...func(*codestar.Options)) (*codestar.ListTagsForProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTagsForProject")
	}

	var r0 *codestar.ListTagsForProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListTagsForProjectInput, ...func(*codestar.Options)) (*codestar.ListTagsForProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListTagsForProjectInput, ...func(*codestar.Options)) *codestar.ListTagsForProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.ListTagsForProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.ListTagsForProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTeamMembers provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTeamMembers(ctx context.Context, params *codestar.ListTeamMembersInput, optFns ...func(*codestar.Options)) (*codestar.ListTeamMembersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTeamMembers")
	}

	var r0 *codestar.ListTeamMembersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListTeamMembersInput, ...func(*codestar.Options)) (*codestar.ListTeamMembersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListTeamMembersInput, ...func(*codestar.Options)) *codestar.ListTeamMembersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.ListTeamMembersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.ListTeamMembersInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserProfiles provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListUserProfiles(ctx context.Context, params *codestar.ListUserProfilesInput, optFns ...func(*codestar.Options)) (*codestar.ListUserProfilesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListUserProfiles")
	}

	var r0 *codestar.ListUserProfilesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListUserProfilesInput, ...func(*codestar.Options)) (*codestar.ListUserProfilesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.ListUserProfilesInput, ...func(*codestar.Options)) *codestar.ListUserProfilesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.ListUserProfilesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.ListUserProfilesInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() codestar.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 codestar.Options
	if rf, ok := ret.Get(0).(func() codestar.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(codestar.Options)
	}

	return r0
}

// TagProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagProject(ctx context.Context, params *codestar.TagProjectInput, optFns ...func(*codestar.Options)) (*codestar.TagProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TagProject")
	}

	var r0 *codestar.TagProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.TagProjectInput, ...func(*codestar.Options)) (*codestar.TagProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.TagProjectInput, ...func(*codestar.Options)) *codestar.TagProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.TagProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.TagProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagProject(ctx context.Context, params *codestar.UntagProjectInput, optFns ...func(*codestar.Options)) (*codestar.UntagProjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UntagProject")
	}

	var r0 *codestar.UntagProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UntagProjectInput, ...func(*codestar.Options)) (*codestar.UntagProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UntagProjectInput, ...func(*codestar.Options)) *codestar.UntagProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.UntagProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.UntagProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProject provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateProject(ctx context.Context, params *codestar.UpdateProjectInput, optFns ...func(*codestar.Options)) (*codestar.UpdateProjectOutput, error) {
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

	var r0 *codestar.UpdateProjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateProjectInput, ...func(*codestar.Options)) (*codestar.UpdateProjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateProjectInput, ...func(*codestar.Options)) *codestar.UpdateProjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.UpdateProjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.UpdateProjectInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTeamMember provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateTeamMember(ctx context.Context, params *codestar.UpdateTeamMemberInput, optFns ...func(*codestar.Options)) (*codestar.UpdateTeamMemberOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTeamMember")
	}

	var r0 *codestar.UpdateTeamMemberOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateTeamMemberInput, ...func(*codestar.Options)) (*codestar.UpdateTeamMemberOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateTeamMemberInput, ...func(*codestar.Options)) *codestar.UpdateTeamMemberOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.UpdateTeamMemberOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.UpdateTeamMemberInput, ...func(*codestar.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserProfile provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateUserProfile(ctx context.Context, params *codestar.UpdateUserProfileInput, optFns ...func(*codestar.Options)) (*codestar.UpdateUserProfileOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserProfile")
	}

	var r0 *codestar.UpdateUserProfileOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateUserProfileInput, ...func(*codestar.Options)) (*codestar.UpdateUserProfileOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codestar.UpdateUserProfileInput, ...func(*codestar.Options)) *codestar.UpdateUserProfileOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codestar.UpdateUserProfileOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codestar.UpdateUserProfileInput, ...func(*codestar.Options)) error); ok {
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
