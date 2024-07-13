package workspaces_test

// tests for the workspaces service interface for this ../../../service/workspaces/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspaces/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspaces/workspaces_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkspacesServiceCanBeMocked(t *testing.T) {
	var iface workspaces_iface.IClient
	iface = &workspaces.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workspaces.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptAccountLinkInvitation", func(t *testing.T) {
        input := &workspaces.AcceptAccountLinkInvitationInput{}
        output := &workspaces.AcceptAccountLinkInvitationOutput{}

        mockClient.On("AcceptAccountLinkInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptAccountLinkInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateConnectionAlias", func(t *testing.T) {
        input := &workspaces.AssociateConnectionAliasInput{}
        output := &workspaces.AssociateConnectionAliasOutput{}

        mockClient.On("AssociateConnectionAlias", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateConnectionAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIpGroups", func(t *testing.T) {
        input := &workspaces.AssociateIpGroupsInput{}
        output := &workspaces.AssociateIpGroupsOutput{}

        mockClient.On("AssociateIpGroups", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIpGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWorkspaceApplication", func(t *testing.T) {
        input := &workspaces.AssociateWorkspaceApplicationInput{}
        output := &workspaces.AssociateWorkspaceApplicationOutput{}

        mockClient.On("AssociateWorkspaceApplication", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWorkspaceApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeIpRules", func(t *testing.T) {
        input := &workspaces.AuthorizeIpRulesInput{}
        output := &workspaces.AuthorizeIpRulesOutput{}

        mockClient.On("AuthorizeIpRules", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeIpRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyWorkspaceImage", func(t *testing.T) {
        input := &workspaces.CopyWorkspaceImageInput{}
        output := &workspaces.CopyWorkspaceImageOutput{}

        mockClient.On("CopyWorkspaceImage", ctx, input).Return(output, nil)

        result, err := mockClient.CopyWorkspaceImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccountLinkInvitation", func(t *testing.T) {
        input := &workspaces.CreateAccountLinkInvitationInput{}
        output := &workspaces.CreateAccountLinkInvitationOutput{}

        mockClient.On("CreateAccountLinkInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccountLinkInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectClientAddIn", func(t *testing.T) {
        input := &workspaces.CreateConnectClientAddInInput{}
        output := &workspaces.CreateConnectClientAddInOutput{}

        mockClient.On("CreateConnectClientAddIn", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectClientAddIn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectionAlias", func(t *testing.T) {
        input := &workspaces.CreateConnectionAliasInput{}
        output := &workspaces.CreateConnectionAliasOutput{}

        mockClient.On("CreateConnectionAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectionAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpGroup", func(t *testing.T) {
        input := &workspaces.CreateIpGroupInput{}
        output := &workspaces.CreateIpGroupOutput{}

        mockClient.On("CreateIpGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStandbyWorkspaces", func(t *testing.T) {
        input := &workspaces.CreateStandbyWorkspacesInput{}
        output := &workspaces.CreateStandbyWorkspacesOutput{}

        mockClient.On("CreateStandbyWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStandbyWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &workspaces.CreateTagsInput{}
        output := &workspaces.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUpdatedWorkspaceImage", func(t *testing.T) {
        input := &workspaces.CreateUpdatedWorkspaceImageInput{}
        output := &workspaces.CreateUpdatedWorkspaceImageOutput{}

        mockClient.On("CreateUpdatedWorkspaceImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUpdatedWorkspaceImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaceBundle", func(t *testing.T) {
        input := &workspaces.CreateWorkspaceBundleInput{}
        output := &workspaces.CreateWorkspaceBundleOutput{}

        mockClient.On("CreateWorkspaceBundle", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaceBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaceImage", func(t *testing.T) {
        input := &workspaces.CreateWorkspaceImageInput{}
        output := &workspaces.CreateWorkspaceImageOutput{}

        mockClient.On("CreateWorkspaceImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaceImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaces", func(t *testing.T) {
        input := &workspaces.CreateWorkspacesInput{}
        output := &workspaces.CreateWorkspacesOutput{}

        mockClient.On("CreateWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspacesPool", func(t *testing.T) {
        input := &workspaces.CreateWorkspacesPoolInput{}
        output := &workspaces.CreateWorkspacesPoolOutput{}

        mockClient.On("CreateWorkspacesPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspacesPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountLinkInvitation", func(t *testing.T) {
        input := &workspaces.DeleteAccountLinkInvitationInput{}
        output := &workspaces.DeleteAccountLinkInvitationOutput{}

        mockClient.On("DeleteAccountLinkInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountLinkInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClientBranding", func(t *testing.T) {
        input := &workspaces.DeleteClientBrandingInput{}
        output := &workspaces.DeleteClientBrandingOutput{}

        mockClient.On("DeleteClientBranding", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClientBranding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectClientAddIn", func(t *testing.T) {
        input := &workspaces.DeleteConnectClientAddInInput{}
        output := &workspaces.DeleteConnectClientAddInOutput{}

        mockClient.On("DeleteConnectClientAddIn", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectClientAddIn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectionAlias", func(t *testing.T) {
        input := &workspaces.DeleteConnectionAliasInput{}
        output := &workspaces.DeleteConnectionAliasOutput{}

        mockClient.On("DeleteConnectionAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectionAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpGroup", func(t *testing.T) {
        input := &workspaces.DeleteIpGroupInput{}
        output := &workspaces.DeleteIpGroupOutput{}

        mockClient.On("DeleteIpGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &workspaces.DeleteTagsInput{}
        output := &workspaces.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspaceBundle", func(t *testing.T) {
        input := &workspaces.DeleteWorkspaceBundleInput{}
        output := &workspaces.DeleteWorkspaceBundleOutput{}

        mockClient.On("DeleteWorkspaceBundle", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspaceBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspaceImage", func(t *testing.T) {
        input := &workspaces.DeleteWorkspaceImageInput{}
        output := &workspaces.DeleteWorkspaceImageOutput{}

        mockClient.On("DeleteWorkspaceImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspaceImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeployWorkspaceApplications", func(t *testing.T) {
        input := &workspaces.DeployWorkspaceApplicationsInput{}
        output := &workspaces.DeployWorkspaceApplicationsOutput{}

        mockClient.On("DeployWorkspaceApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DeployWorkspaceApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterWorkspaceDirectory", func(t *testing.T) {
        input := &workspaces.DeregisterWorkspaceDirectoryInput{}
        output := &workspaces.DeregisterWorkspaceDirectoryOutput{}

        mockClient.On("DeregisterWorkspaceDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterWorkspaceDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccount", func(t *testing.T) {
        input := &workspaces.DescribeAccountInput{}
        output := &workspaces.DescribeAccountOutput{}

        mockClient.On("DescribeAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountModifications", func(t *testing.T) {
        input := &workspaces.DescribeAccountModificationsInput{}
        output := &workspaces.DescribeAccountModificationsOutput{}

        mockClient.On("DescribeAccountModifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationAssociations", func(t *testing.T) {
        input := &workspaces.DescribeApplicationAssociationsInput{}
        output := &workspaces.DescribeApplicationAssociationsOutput{}

        mockClient.On("DescribeApplicationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplications", func(t *testing.T) {
        input := &workspaces.DescribeApplicationsInput{}
        output := &workspaces.DescribeApplicationsOutput{}

        mockClient.On("DescribeApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBundleAssociations", func(t *testing.T) {
        input := &workspaces.DescribeBundleAssociationsInput{}
        output := &workspaces.DescribeBundleAssociationsOutput{}

        mockClient.On("DescribeBundleAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBundleAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientBranding", func(t *testing.T) {
        input := &workspaces.DescribeClientBrandingInput{}
        output := &workspaces.DescribeClientBrandingOutput{}

        mockClient.On("DescribeClientBranding", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientBranding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientProperties", func(t *testing.T) {
        input := &workspaces.DescribeClientPropertiesInput{}
        output := &workspaces.DescribeClientPropertiesOutput{}

        mockClient.On("DescribeClientProperties", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectClientAddIns", func(t *testing.T) {
        input := &workspaces.DescribeConnectClientAddInsInput{}
        output := &workspaces.DescribeConnectClientAddInsOutput{}

        mockClient.On("DescribeConnectClientAddIns", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectClientAddIns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectionAliasPermissions", func(t *testing.T) {
        input := &workspaces.DescribeConnectionAliasPermissionsInput{}
        output := &workspaces.DescribeConnectionAliasPermissionsOutput{}

        mockClient.On("DescribeConnectionAliasPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectionAliasPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectionAliases", func(t *testing.T) {
        input := &workspaces.DescribeConnectionAliasesInput{}
        output := &workspaces.DescribeConnectionAliasesOutput{}

        mockClient.On("DescribeConnectionAliases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectionAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageAssociations", func(t *testing.T) {
        input := &workspaces.DescribeImageAssociationsInput{}
        output := &workspaces.DescribeImageAssociationsOutput{}

        mockClient.On("DescribeImageAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpGroups", func(t *testing.T) {
        input := &workspaces.DescribeIpGroupsInput{}
        output := &workspaces.DescribeIpGroupsOutput{}

        mockClient.On("DescribeIpGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &workspaces.DescribeTagsInput{}
        output := &workspaces.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceAssociations", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceAssociationsInput{}
        output := &workspaces.DescribeWorkspaceAssociationsOutput{}

        mockClient.On("DescribeWorkspaceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceBundles", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceBundlesInput{}
        output := &workspaces.DescribeWorkspaceBundlesOutput{}

        mockClient.On("DescribeWorkspaceBundles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceDirectories", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceDirectoriesInput{}
        output := &workspaces.DescribeWorkspaceDirectoriesOutput{}

        mockClient.On("DescribeWorkspaceDirectories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceDirectories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceImagePermissions", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceImagePermissionsInput{}
        output := &workspaces.DescribeWorkspaceImagePermissionsOutput{}

        mockClient.On("DescribeWorkspaceImagePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceImagePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceImages", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceImagesInput{}
        output := &workspaces.DescribeWorkspaceImagesOutput{}

        mockClient.On("DescribeWorkspaceImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceSnapshots", func(t *testing.T) {
        input := &workspaces.DescribeWorkspaceSnapshotsInput{}
        output := &workspaces.DescribeWorkspaceSnapshotsOutput{}

        mockClient.On("DescribeWorkspaceSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaces", func(t *testing.T) {
        input := &workspaces.DescribeWorkspacesInput{}
        output := &workspaces.DescribeWorkspacesOutput{}

        mockClient.On("DescribeWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspacesConnectionStatus", func(t *testing.T) {
        input := &workspaces.DescribeWorkspacesConnectionStatusInput{}
        output := &workspaces.DescribeWorkspacesConnectionStatusOutput{}

        mockClient.On("DescribeWorkspacesConnectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspacesConnectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspacesPoolSessions", func(t *testing.T) {
        input := &workspaces.DescribeWorkspacesPoolSessionsInput{}
        output := &workspaces.DescribeWorkspacesPoolSessionsOutput{}

        mockClient.On("DescribeWorkspacesPoolSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspacesPoolSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspacesPools", func(t *testing.T) {
        input := &workspaces.DescribeWorkspacesPoolsInput{}
        output := &workspaces.DescribeWorkspacesPoolsOutput{}

        mockClient.On("DescribeWorkspacesPools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspacesPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateConnectionAlias", func(t *testing.T) {
        input := &workspaces.DisassociateConnectionAliasInput{}
        output := &workspaces.DisassociateConnectionAliasOutput{}

        mockClient.On("DisassociateConnectionAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateConnectionAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIpGroups", func(t *testing.T) {
        input := &workspaces.DisassociateIpGroupsInput{}
        output := &workspaces.DisassociateIpGroupsOutput{}

        mockClient.On("DisassociateIpGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIpGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWorkspaceApplication", func(t *testing.T) {
        input := &workspaces.DisassociateWorkspaceApplicationInput{}
        output := &workspaces.DisassociateWorkspaceApplicationOutput{}

        mockClient.On("DisassociateWorkspaceApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWorkspaceApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountLink", func(t *testing.T) {
        input := &workspaces.GetAccountLinkInput{}
        output := &workspaces.GetAccountLinkOutput{}

        mockClient.On("GetAccountLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportClientBranding", func(t *testing.T) {
        input := &workspaces.ImportClientBrandingInput{}
        output := &workspaces.ImportClientBrandingOutput{}

        mockClient.On("ImportClientBranding", ctx, input).Return(output, nil)

        result, err := mockClient.ImportClientBranding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportWorkspaceImage", func(t *testing.T) {
        input := &workspaces.ImportWorkspaceImageInput{}
        output := &workspaces.ImportWorkspaceImageOutput{}

        mockClient.On("ImportWorkspaceImage", ctx, input).Return(output, nil)

        result, err := mockClient.ImportWorkspaceImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountLinks", func(t *testing.T) {
        input := &workspaces.ListAccountLinksInput{}
        output := &workspaces.ListAccountLinksOutput{}

        mockClient.On("ListAccountLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableManagementCidrRanges", func(t *testing.T) {
        input := &workspaces.ListAvailableManagementCidrRangesInput{}
        output := &workspaces.ListAvailableManagementCidrRangesOutput{}

        mockClient.On("ListAvailableManagementCidrRanges", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableManagementCidrRanges(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMigrateWorkspace", func(t *testing.T) {
        input := &workspaces.MigrateWorkspaceInput{}
        output := &workspaces.MigrateWorkspaceOutput{}

        mockClient.On("MigrateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.MigrateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyAccount", func(t *testing.T) {
        input := &workspaces.ModifyAccountInput{}
        output := &workspaces.ModifyAccountOutput{}

        mockClient.On("ModifyAccount", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCertificateBasedAuthProperties", func(t *testing.T) {
        input := &workspaces.ModifyCertificateBasedAuthPropertiesInput{}
        output := &workspaces.ModifyCertificateBasedAuthPropertiesOutput{}

        mockClient.On("ModifyCertificateBasedAuthProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCertificateBasedAuthProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClientProperties", func(t *testing.T) {
        input := &workspaces.ModifyClientPropertiesInput{}
        output := &workspaces.ModifyClientPropertiesOutput{}

        mockClient.On("ModifyClientProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClientProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySamlProperties", func(t *testing.T) {
        input := &workspaces.ModifySamlPropertiesInput{}
        output := &workspaces.ModifySamlPropertiesOutput{}

        mockClient.On("ModifySamlProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySamlProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySelfservicePermissions", func(t *testing.T) {
        input := &workspaces.ModifySelfservicePermissionsInput{}
        output := &workspaces.ModifySelfservicePermissionsOutput{}

        mockClient.On("ModifySelfservicePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySelfservicePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyStreamingProperties", func(t *testing.T) {
        input := &workspaces.ModifyStreamingPropertiesInput{}
        output := &workspaces.ModifyStreamingPropertiesOutput{}

        mockClient.On("ModifyStreamingProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyStreamingProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyWorkspaceAccessProperties", func(t *testing.T) {
        input := &workspaces.ModifyWorkspaceAccessPropertiesInput{}
        output := &workspaces.ModifyWorkspaceAccessPropertiesOutput{}

        mockClient.On("ModifyWorkspaceAccessProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyWorkspaceAccessProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyWorkspaceCreationProperties", func(t *testing.T) {
        input := &workspaces.ModifyWorkspaceCreationPropertiesInput{}
        output := &workspaces.ModifyWorkspaceCreationPropertiesOutput{}

        mockClient.On("ModifyWorkspaceCreationProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyWorkspaceCreationProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyWorkspaceProperties", func(t *testing.T) {
        input := &workspaces.ModifyWorkspacePropertiesInput{}
        output := &workspaces.ModifyWorkspacePropertiesOutput{}

        mockClient.On("ModifyWorkspaceProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyWorkspaceProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyWorkspaceState", func(t *testing.T) {
        input := &workspaces.ModifyWorkspaceStateInput{}
        output := &workspaces.ModifyWorkspaceStateOutput{}

        mockClient.On("ModifyWorkspaceState", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyWorkspaceState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootWorkspaces", func(t *testing.T) {
        input := &workspaces.RebootWorkspacesInput{}
        output := &workspaces.RebootWorkspacesOutput{}

        mockClient.On("RebootWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.RebootWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebuildWorkspaces", func(t *testing.T) {
        input := &workspaces.RebuildWorkspacesInput{}
        output := &workspaces.RebuildWorkspacesOutput{}

        mockClient.On("RebuildWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.RebuildWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterWorkspaceDirectory", func(t *testing.T) {
        input := &workspaces.RegisterWorkspaceDirectoryInput{}
        output := &workspaces.RegisterWorkspaceDirectoryOutput{}

        mockClient.On("RegisterWorkspaceDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterWorkspaceDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectAccountLinkInvitation", func(t *testing.T) {
        input := &workspaces.RejectAccountLinkInvitationInput{}
        output := &workspaces.RejectAccountLinkInvitationOutput{}

        mockClient.On("RejectAccountLinkInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.RejectAccountLinkInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreWorkspace", func(t *testing.T) {
        input := &workspaces.RestoreWorkspaceInput{}
        output := &workspaces.RestoreWorkspaceOutput{}

        mockClient.On("RestoreWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeIpRules", func(t *testing.T) {
        input := &workspaces.RevokeIpRulesInput{}
        output := &workspaces.RevokeIpRulesOutput{}

        mockClient.On("RevokeIpRules", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeIpRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkspaces", func(t *testing.T) {
        input := &workspaces.StartWorkspacesInput{}
        output := &workspaces.StartWorkspacesOutput{}

        mockClient.On("StartWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkspacesPool", func(t *testing.T) {
        input := &workspaces.StartWorkspacesPoolInput{}
        output := &workspaces.StartWorkspacesPoolOutput{}

        mockClient.On("StartWorkspacesPool", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkspacesPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopWorkspaces", func(t *testing.T) {
        input := &workspaces.StopWorkspacesInput{}
        output := &workspaces.StopWorkspacesOutput{}

        mockClient.On("StopWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.StopWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopWorkspacesPool", func(t *testing.T) {
        input := &workspaces.StopWorkspacesPoolInput{}
        output := &workspaces.StopWorkspacesPoolOutput{}

        mockClient.On("StopWorkspacesPool", ctx, input).Return(output, nil)

        result, err := mockClient.StopWorkspacesPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateWorkspaces", func(t *testing.T) {
        input := &workspaces.TerminateWorkspacesInput{}
        output := &workspaces.TerminateWorkspacesOutput{}

        mockClient.On("TerminateWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateWorkspacesPool", func(t *testing.T) {
        input := &workspaces.TerminateWorkspacesPoolInput{}
        output := &workspaces.TerminateWorkspacesPoolOutput{}

        mockClient.On("TerminateWorkspacesPool", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateWorkspacesPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateWorkspacesPoolSession", func(t *testing.T) {
        input := &workspaces.TerminateWorkspacesPoolSessionInput{}
        output := &workspaces.TerminateWorkspacesPoolSessionOutput{}

        mockClient.On("TerminateWorkspacesPoolSession", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateWorkspacesPoolSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectClientAddIn", func(t *testing.T) {
        input := &workspaces.UpdateConnectClientAddInInput{}
        output := &workspaces.UpdateConnectClientAddInOutput{}

        mockClient.On("UpdateConnectClientAddIn", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectClientAddIn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectionAliasPermission", func(t *testing.T) {
        input := &workspaces.UpdateConnectionAliasPermissionInput{}
        output := &workspaces.UpdateConnectionAliasPermissionOutput{}

        mockClient.On("UpdateConnectionAliasPermission", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectionAliasPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRulesOfIpGroup", func(t *testing.T) {
        input := &workspaces.UpdateRulesOfIpGroupInput{}
        output := &workspaces.UpdateRulesOfIpGroupOutput{}

        mockClient.On("UpdateRulesOfIpGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRulesOfIpGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspaceBundle", func(t *testing.T) {
        input := &workspaces.UpdateWorkspaceBundleInput{}
        output := &workspaces.UpdateWorkspaceBundleOutput{}

        mockClient.On("UpdateWorkspaceBundle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspaceBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspaceImagePermission", func(t *testing.T) {
        input := &workspaces.UpdateWorkspaceImagePermissionInput{}
        output := &workspaces.UpdateWorkspaceImagePermissionOutput{}

        mockClient.On("UpdateWorkspaceImagePermission", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspaceImagePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspacesPool", func(t *testing.T) {
        input := &workspaces.UpdateWorkspacesPoolInput{}
        output := &workspaces.UpdateWorkspacesPoolOutput{}

        mockClient.On("UpdateWorkspacesPool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspacesPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
