package identitystore_test

// tests for the identitystore service interface for this ../../../service/identitystore/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/identitystore/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/identitystore/identitystore_iface"
	"github.com/stretchr/testify/assert"
)

func TestIdentitystoreServiceCanBeMocked(t *testing.T) {
	var iface identitystore_iface.IClient
	iface = &identitystore.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := identitystore.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &identitystore.CreateGroupInput{}
        output := &identitystore.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroupMembership", func(t *testing.T) {
        input := &identitystore.CreateGroupMembershipInput{}
        output := &identitystore.CreateGroupMembershipOutput{}

        mockClient.On("CreateGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &identitystore.CreateUserInput{}
        output := &identitystore.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &identitystore.DeleteGroupInput{}
        output := &identitystore.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroupMembership", func(t *testing.T) {
        input := &identitystore.DeleteGroupMembershipInput{}
        output := &identitystore.DeleteGroupMembershipOutput{}

        mockClient.On("DeleteGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &identitystore.DeleteUserInput{}
        output := &identitystore.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroup", func(t *testing.T) {
        input := &identitystore.DescribeGroupInput{}
        output := &identitystore.DescribeGroupOutput{}

        mockClient.On("DescribeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroupMembership", func(t *testing.T) {
        input := &identitystore.DescribeGroupMembershipInput{}
        output := &identitystore.DescribeGroupMembershipOutput{}

        mockClient.On("DescribeGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &identitystore.DescribeUserInput{}
        output := &identitystore.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupId", func(t *testing.T) {
        input := &identitystore.GetGroupIdInput{}
        output := &identitystore.GetGroupIdOutput{}

        mockClient.On("GetGroupId", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupMembershipId", func(t *testing.T) {
        input := &identitystore.GetGroupMembershipIdInput{}
        output := &identitystore.GetGroupMembershipIdOutput{}

        mockClient.On("GetGroupMembershipId", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupMembershipId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserId", func(t *testing.T) {
        input := &identitystore.GetUserIdInput{}
        output := &identitystore.GetUserIdOutput{}

        mockClient.On("GetUserId", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIsMemberInGroups", func(t *testing.T) {
        input := &identitystore.IsMemberInGroupsInput{}
        output := &identitystore.IsMemberInGroupsOutput{}

        mockClient.On("IsMemberInGroups", ctx, input).Return(output, nil)

        result, err := mockClient.IsMemberInGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupMemberships", func(t *testing.T) {
        input := &identitystore.ListGroupMembershipsInput{}
        output := &identitystore.ListGroupMembershipsOutput{}

        mockClient.On("ListGroupMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupMembershipsForMember", func(t *testing.T) {
        input := &identitystore.ListGroupMembershipsForMemberInput{}
        output := &identitystore.ListGroupMembershipsForMemberOutput{}

        mockClient.On("ListGroupMembershipsForMember", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupMembershipsForMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &identitystore.ListGroupsInput{}
        output := &identitystore.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &identitystore.ListUsersInput{}
        output := &identitystore.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &identitystore.UpdateGroupInput{}
        output := &identitystore.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &identitystore.UpdateUserInput{}
        output := &identitystore.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
