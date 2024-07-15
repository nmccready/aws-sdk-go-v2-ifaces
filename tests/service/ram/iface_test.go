// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package ram_test

// tests for the ram service interface for 
// this ../../../service/ram/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ram/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ram/ram_iface"
	"github.com/stretchr/testify/assert"
)

func TestRamServiceCanBeMocked(t *testing.T) {
	var iface ram_iface.IClient
	iface = &ram.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ram.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptResourceShareInvitation", func(t *testing.T) {
        input := &ram.AcceptResourceShareInvitationInput{}
        output := &ram.AcceptResourceShareInvitationOutput{}

        mockClient.On("AcceptResourceShareInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptResourceShareInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResourceShare", func(t *testing.T) {
        input := &ram.AssociateResourceShareInput{}
        output := &ram.AssociateResourceShareOutput{}

        mockClient.On("AssociateResourceShare", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResourceShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResourceSharePermission", func(t *testing.T) {
        input := &ram.AssociateResourceSharePermissionInput{}
        output := &ram.AssociateResourceSharePermissionOutput{}

        mockClient.On("AssociateResourceSharePermission", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResourceSharePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePermission", func(t *testing.T) {
        input := &ram.CreatePermissionInput{}
        output := &ram.CreatePermissionOutput{}

        mockClient.On("CreatePermission", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePermissionVersion", func(t *testing.T) {
        input := &ram.CreatePermissionVersionInput{}
        output := &ram.CreatePermissionVersionOutput{}

        mockClient.On("CreatePermissionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePermissionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceShare", func(t *testing.T) {
        input := &ram.CreateResourceShareInput{}
        output := &ram.CreateResourceShareOutput{}

        mockClient.On("CreateResourceShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermission", func(t *testing.T) {
        input := &ram.DeletePermissionInput{}
        output := &ram.DeletePermissionOutput{}

        mockClient.On("DeletePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionVersion", func(t *testing.T) {
        input := &ram.DeletePermissionVersionInput{}
        output := &ram.DeletePermissionVersionOutput{}

        mockClient.On("DeletePermissionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceShare", func(t *testing.T) {
        input := &ram.DeleteResourceShareInput{}
        output := &ram.DeleteResourceShareOutput{}

        mockClient.On("DeleteResourceShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResourceShare", func(t *testing.T) {
        input := &ram.DisassociateResourceShareInput{}
        output := &ram.DisassociateResourceShareOutput{}

        mockClient.On("DisassociateResourceShare", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResourceShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResourceSharePermission", func(t *testing.T) {
        input := &ram.DisassociateResourceSharePermissionInput{}
        output := &ram.DisassociateResourceSharePermissionOutput{}

        mockClient.On("DisassociateResourceSharePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResourceSharePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSharingWithAwsOrganization", func(t *testing.T) {
        input := &ram.EnableSharingWithAwsOrganizationInput{}
        output := &ram.EnableSharingWithAwsOrganizationOutput{}

        mockClient.On("EnableSharingWithAwsOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSharingWithAwsOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermission", func(t *testing.T) {
        input := &ram.GetPermissionInput{}
        output := &ram.GetPermissionOutput{}

        mockClient.On("GetPermission", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicies", func(t *testing.T) {
        input := &ram.GetResourcePoliciesInput{}
        output := &ram.GetResourcePoliciesOutput{}

        mockClient.On("GetResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceShareAssociations", func(t *testing.T) {
        input := &ram.GetResourceShareAssociationsInput{}
        output := &ram.GetResourceShareAssociationsOutput{}

        mockClient.On("GetResourceShareAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceShareAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceShareInvitations", func(t *testing.T) {
        input := &ram.GetResourceShareInvitationsInput{}
        output := &ram.GetResourceShareInvitationsOutput{}

        mockClient.On("GetResourceShareInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceShareInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceShares", func(t *testing.T) {
        input := &ram.GetResourceSharesInput{}
        output := &ram.GetResourceSharesOutput{}

        mockClient.On("GetResourceShares", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPendingInvitationResources", func(t *testing.T) {
        input := &ram.ListPendingInvitationResourcesInput{}
        output := &ram.ListPendingInvitationResourcesOutput{}

        mockClient.On("ListPendingInvitationResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListPendingInvitationResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionAssociations", func(t *testing.T) {
        input := &ram.ListPermissionAssociationsInput{}
        output := &ram.ListPermissionAssociationsOutput{}

        mockClient.On("ListPermissionAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionVersions", func(t *testing.T) {
        input := &ram.ListPermissionVersionsInput{}
        output := &ram.ListPermissionVersionsOutput{}

        mockClient.On("ListPermissionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissions", func(t *testing.T) {
        input := &ram.ListPermissionsInput{}
        output := &ram.ListPermissionsOutput{}

        mockClient.On("ListPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrincipals", func(t *testing.T) {
        input := &ram.ListPrincipalsInput{}
        output := &ram.ListPrincipalsOutput{}

        mockClient.On("ListPrincipals", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrincipals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReplacePermissionAssociationsWork", func(t *testing.T) {
        input := &ram.ListReplacePermissionAssociationsWorkInput{}
        output := &ram.ListReplacePermissionAssociationsWorkOutput{}

        mockClient.On("ListReplacePermissionAssociationsWork", ctx, input).Return(output, nil)

        result, err := mockClient.ListReplacePermissionAssociationsWork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceSharePermissions", func(t *testing.T) {
        input := &ram.ListResourceSharePermissionsInput{}
        output := &ram.ListResourceSharePermissionsOutput{}

        mockClient.On("ListResourceSharePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceSharePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceTypes", func(t *testing.T) {
        input := &ram.ListResourceTypesInput{}
        output := &ram.ListResourceTypesOutput{}

        mockClient.On("ListResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResources", func(t *testing.T) {
        input := &ram.ListResourcesInput{}
        output := &ram.ListResourcesOutput{}

        mockClient.On("ListResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromotePermissionCreatedFromPolicy", func(t *testing.T) {
        input := &ram.PromotePermissionCreatedFromPolicyInput{}
        output := &ram.PromotePermissionCreatedFromPolicyOutput{}

        mockClient.On("PromotePermissionCreatedFromPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PromotePermissionCreatedFromPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromoteResourceShareCreatedFromPolicy", func(t *testing.T) {
        input := &ram.PromoteResourceShareCreatedFromPolicyInput{}
        output := &ram.PromoteResourceShareCreatedFromPolicyOutput{}

        mockClient.On("PromoteResourceShareCreatedFromPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PromoteResourceShareCreatedFromPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectResourceShareInvitation", func(t *testing.T) {
        input := &ram.RejectResourceShareInvitationInput{}
        output := &ram.RejectResourceShareInvitationOutput{}

        mockClient.On("RejectResourceShareInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.RejectResourceShareInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplacePermissionAssociations", func(t *testing.T) {
        input := &ram.ReplacePermissionAssociationsInput{}
        output := &ram.ReplacePermissionAssociationsOutput{}

        mockClient.On("ReplacePermissionAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ReplacePermissionAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultPermissionVersion", func(t *testing.T) {
        input := &ram.SetDefaultPermissionVersionInput{}
        output := &ram.SetDefaultPermissionVersionOutput{}

        mockClient.On("SetDefaultPermissionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultPermissionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ram.TagResourceInput{}
        output := &ram.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ram.UntagResourceInput{}
        output := &ram.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceShare", func(t *testing.T) {
        input := &ram.UpdateResourceShareInput{}
        output := &ram.UpdateResourceShareOutput{}

        mockClient.On("UpdateResourceShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
