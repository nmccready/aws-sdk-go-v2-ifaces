// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	identitystore "github.com/aws/aws-sdk-go-v2/service/identitystore"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateGroup(ctx context.Context, params *identitystore.CreateGroupInput, optFns ...func(*identitystore.Options)) (*identitystore.CreateGroupOutput, error) {
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

	var r0 *identitystore.CreateGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateGroupInput, ...func(*identitystore.Options)) (*identitystore.CreateGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateGroupInput, ...func(*identitystore.Options)) *identitystore.CreateGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.CreateGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.CreateGroupInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroupMembership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateGroupMembership(ctx context.Context, params *identitystore.CreateGroupMembershipInput, optFns ...func(*identitystore.Options)) (*identitystore.CreateGroupMembershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateGroupMembership")
	}

	var r0 *identitystore.CreateGroupMembershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateGroupMembershipInput, ...func(*identitystore.Options)) (*identitystore.CreateGroupMembershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateGroupMembershipInput, ...func(*identitystore.Options)) *identitystore.CreateGroupMembershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.CreateGroupMembershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.CreateGroupMembershipInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateUser(ctx context.Context, params *identitystore.CreateUserInput, optFns ...func(*identitystore.Options)) (*identitystore.CreateUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *identitystore.CreateUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateUserInput, ...func(*identitystore.Options)) (*identitystore.CreateUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.CreateUserInput, ...func(*identitystore.Options)) *identitystore.CreateUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.CreateUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.CreateUserInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteGroup(ctx context.Context, params *identitystore.DeleteGroupInput, optFns ...func(*identitystore.Options)) (*identitystore.DeleteGroupOutput, error) {
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

	var r0 *identitystore.DeleteGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteGroupInput, ...func(*identitystore.Options)) (*identitystore.DeleteGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteGroupInput, ...func(*identitystore.Options)) *identitystore.DeleteGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DeleteGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DeleteGroupInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroupMembership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteGroupMembership(ctx context.Context, params *identitystore.DeleteGroupMembershipInput, optFns ...func(*identitystore.Options)) (*identitystore.DeleteGroupMembershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroupMembership")
	}

	var r0 *identitystore.DeleteGroupMembershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteGroupMembershipInput, ...func(*identitystore.Options)) (*identitystore.DeleteGroupMembershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteGroupMembershipInput, ...func(*identitystore.Options)) *identitystore.DeleteGroupMembershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DeleteGroupMembershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DeleteGroupMembershipInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteUser(ctx context.Context, params *identitystore.DeleteUserInput, optFns ...func(*identitystore.Options)) (*identitystore.DeleteUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 *identitystore.DeleteUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteUserInput, ...func(*identitystore.Options)) (*identitystore.DeleteUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DeleteUserInput, ...func(*identitystore.Options)) *identitystore.DeleteUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DeleteUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DeleteUserInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeGroup(ctx context.Context, params *identitystore.DescribeGroupInput, optFns ...func(*identitystore.Options)) (*identitystore.DescribeGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeGroup")
	}

	var r0 *identitystore.DescribeGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeGroupInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeGroupInput, ...func(*identitystore.Options)) *identitystore.DescribeGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DescribeGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DescribeGroupInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeGroupMembership provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeGroupMembership(ctx context.Context, params *identitystore.DescribeGroupMembershipInput, optFns ...func(*identitystore.Options)) (*identitystore.DescribeGroupMembershipOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeGroupMembership")
	}

	var r0 *identitystore.DescribeGroupMembershipOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeGroupMembershipInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupMembershipOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeGroupMembershipInput, ...func(*identitystore.Options)) *identitystore.DescribeGroupMembershipOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DescribeGroupMembershipOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DescribeGroupMembershipInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeUser(ctx context.Context, params *identitystore.DescribeUserInput, optFns ...func(*identitystore.Options)) (*identitystore.DescribeUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeUser")
	}

	var r0 *identitystore.DescribeUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeUserInput, ...func(*identitystore.Options)) (*identitystore.DescribeUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.DescribeUserInput, ...func(*identitystore.Options)) *identitystore.DescribeUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.DescribeUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.DescribeUserInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupId provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroupId(ctx context.Context, params *identitystore.GetGroupIdInput, optFns ...func(*identitystore.Options)) (*identitystore.GetGroupIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupId")
	}

	var r0 *identitystore.GetGroupIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetGroupIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetGroupIdInput, ...func(*identitystore.Options)) *identitystore.GetGroupIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.GetGroupIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.GetGroupIdInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupMembershipId provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetGroupMembershipId(ctx context.Context, params *identitystore.GetGroupMembershipIdInput, optFns ...func(*identitystore.Options)) (*identitystore.GetGroupMembershipIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupMembershipId")
	}

	var r0 *identitystore.GetGroupMembershipIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetGroupMembershipIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupMembershipIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetGroupMembershipIdInput, ...func(*identitystore.Options)) *identitystore.GetGroupMembershipIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.GetGroupMembershipIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.GetGroupMembershipIdInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserId provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetUserId(ctx context.Context, params *identitystore.GetUserIdInput, optFns ...func(*identitystore.Options)) (*identitystore.GetUserIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetUserId")
	}

	var r0 *identitystore.GetUserIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetUserIdInput, ...func(*identitystore.Options)) (*identitystore.GetUserIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.GetUserIdInput, ...func(*identitystore.Options)) *identitystore.GetUserIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.GetUserIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.GetUserIdInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsMemberInGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) IsMemberInGroups(ctx context.Context, params *identitystore.IsMemberInGroupsInput, optFns ...func(*identitystore.Options)) (*identitystore.IsMemberInGroupsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsMemberInGroups")
	}

	var r0 *identitystore.IsMemberInGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.IsMemberInGroupsInput, ...func(*identitystore.Options)) (*identitystore.IsMemberInGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.IsMemberInGroupsInput, ...func(*identitystore.Options)) *identitystore.IsMemberInGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.IsMemberInGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.IsMemberInGroupsInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroupMemberships provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroupMemberships(ctx context.Context, params *identitystore.ListGroupMembershipsInput, optFns ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupMemberships")
	}

	var r0 *identitystore.ListGroupMembershipsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupMembershipsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupMembershipsInput, ...func(*identitystore.Options)) *identitystore.ListGroupMembershipsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.ListGroupMembershipsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.ListGroupMembershipsInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroupMembershipsForMember provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroupMembershipsForMember(ctx context.Context, params *identitystore.ListGroupMembershipsForMemberInput, optFns ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsForMemberOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupMembershipsForMember")
	}

	var r0 *identitystore.ListGroupMembershipsForMemberOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupMembershipsForMemberInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsForMemberOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupMembershipsForMemberInput, ...func(*identitystore.Options)) *identitystore.ListGroupMembershipsForMemberOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.ListGroupMembershipsForMemberOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.ListGroupMembershipsForMemberInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListGroups(ctx context.Context, params *identitystore.ListGroupsInput, optFns ...func(*identitystore.Options)) (*identitystore.ListGroupsOutput, error) {
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

	var r0 *identitystore.ListGroupsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListGroupsInput, ...func(*identitystore.Options)) *identitystore.ListGroupsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.ListGroupsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.ListGroupsInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListUsers(ctx context.Context, params *identitystore.ListUsersInput, optFns ...func(*identitystore.Options)) (*identitystore.ListUsersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 *identitystore.ListUsersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListUsersInput, ...func(*identitystore.Options)) (*identitystore.ListUsersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.ListUsersInput, ...func(*identitystore.Options)) *identitystore.ListUsersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.ListUsersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.ListUsersInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() identitystore.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 identitystore.Options
	if rf, ok := ret.Get(0).(func() identitystore.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(identitystore.Options)
	}

	return r0
}

// UpdateGroup provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateGroup(ctx context.Context, params *identitystore.UpdateGroupInput, optFns ...func(*identitystore.Options)) (*identitystore.UpdateGroupOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGroup")
	}

	var r0 *identitystore.UpdateGroupOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.UpdateGroupInput, ...func(*identitystore.Options)) (*identitystore.UpdateGroupOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.UpdateGroupInput, ...func(*identitystore.Options)) *identitystore.UpdateGroupOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.UpdateGroupOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.UpdateGroupInput, ...func(*identitystore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateUser(ctx context.Context, params *identitystore.UpdateUserInput, optFns ...func(*identitystore.Options)) (*identitystore.UpdateUserOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *identitystore.UpdateUserOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.UpdateUserInput, ...func(*identitystore.Options)) (*identitystore.UpdateUserOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *identitystore.UpdateUserInput, ...func(*identitystore.Options)) *identitystore.UpdateUserOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identitystore.UpdateUserOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *identitystore.UpdateUserInput, ...func(*identitystore.Options)) error); ok {
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
