package codestar_test

// tests for the codestar service interface for this ../../../service/codestar/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestar/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestar/codestar_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodestarServiceCanBeMocked(t *testing.T) {
	var iface codestar_iface.IClient
	iface = &codestar.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codestar.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTeamMember", func(t *testing.T) {
        input := &codestar.AssociateTeamMemberInput{}
        output := &codestar.AssociateTeamMemberOutput{}

        mockClient.On("AssociateTeamMember", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTeamMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &codestar.CreateProjectInput{}
        output := &codestar.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserProfile", func(t *testing.T) {
        input := &codestar.CreateUserProfileInput{}
        output := &codestar.CreateUserProfileOutput{}

        mockClient.On("CreateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &codestar.DeleteProjectInput{}
        output := &codestar.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserProfile", func(t *testing.T) {
        input := &codestar.DeleteUserProfileInput{}
        output := &codestar.DeleteUserProfileOutput{}

        mockClient.On("DeleteUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &codestar.DescribeProjectInput{}
        output := &codestar.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserProfile", func(t *testing.T) {
        input := &codestar.DescribeUserProfileInput{}
        output := &codestar.DescribeUserProfileOutput{}

        mockClient.On("DescribeUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTeamMember", func(t *testing.T) {
        input := &codestar.DisassociateTeamMemberInput{}
        output := &codestar.DisassociateTeamMemberOutput{}

        mockClient.On("DisassociateTeamMember", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTeamMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &codestar.ListProjectsInput{}
        output := &codestar.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResources", func(t *testing.T) {
        input := &codestar.ListResourcesInput{}
        output := &codestar.ListResourcesOutput{}

        mockClient.On("ListResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForProject", func(t *testing.T) {
        input := &codestar.ListTagsForProjectInput{}
        output := &codestar.ListTagsForProjectOutput{}

        mockClient.On("ListTagsForProject", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTeamMembers", func(t *testing.T) {
        input := &codestar.ListTeamMembersInput{}
        output := &codestar.ListTeamMembersOutput{}

        mockClient.On("ListTeamMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTeamMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserProfiles", func(t *testing.T) {
        input := &codestar.ListUserProfilesInput{}
        output := &codestar.ListUserProfilesOutput{}

        mockClient.On("ListUserProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagProject", func(t *testing.T) {
        input := &codestar.TagProjectInput{}
        output := &codestar.TagProjectOutput{}

        mockClient.On("TagProject", ctx, input).Return(output, nil)

        result, err := mockClient.TagProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagProject", func(t *testing.T) {
        input := &codestar.UntagProjectInput{}
        output := &codestar.UntagProjectOutput{}

        mockClient.On("UntagProject", ctx, input).Return(output, nil)

        result, err := mockClient.UntagProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &codestar.UpdateProjectInput{}
        output := &codestar.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTeamMember", func(t *testing.T) {
        input := &codestar.UpdateTeamMemberInput{}
        output := &codestar.UpdateTeamMemberOutput{}

        mockClient.On("UpdateTeamMember", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTeamMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserProfile", func(t *testing.T) {
        input := &codestar.UpdateUserProfileInput{}
        output := &codestar.UpdateUserProfileOutput{}

        mockClient.On("UpdateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
