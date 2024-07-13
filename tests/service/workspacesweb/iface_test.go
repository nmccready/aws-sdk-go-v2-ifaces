package workspacesweb_test

// tests for the workspacesweb service interface for this ../../../service/workspacesweb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspacesweb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspacesweb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspacesweb/workspacesweb_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceswebServiceCanBeMocked(t *testing.T) {
	var iface workspacesweb_iface.IClient
	iface = &workspacesweb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workspacesweb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.AssociateBrowserSettingsInput{}
        output := &workspacesweb.AssociateBrowserSettingsOutput{}

        mockClient.On("AssociateBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.AssociateIpAccessSettingsInput{}
        output := &workspacesweb.AssociateIpAccessSettingsOutput{}

        mockClient.On("AssociateIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.AssociateNetworkSettingsInput{}
        output := &workspacesweb.AssociateNetworkSettingsOutput{}

        mockClient.On("AssociateNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTrustStore", func(t *testing.T) {
        input := &workspacesweb.AssociateTrustStoreInput{}
        output := &workspacesweb.AssociateTrustStoreOutput{}

        mockClient.On("AssociateTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.AssociateUserAccessLoggingSettingsInput{}
        output := &workspacesweb.AssociateUserAccessLoggingSettingsOutput{}

        mockClient.On("AssociateUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateUserSettings", func(t *testing.T) {
        input := &workspacesweb.AssociateUserSettingsInput{}
        output := &workspacesweb.AssociateUserSettingsOutput{}

        mockClient.On("AssociateUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.CreateBrowserSettingsInput{}
        output := &workspacesweb.CreateBrowserSettingsOutput{}

        mockClient.On("CreateBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdentityProvider", func(t *testing.T) {
        input := &workspacesweb.CreateIdentityProviderInput{}
        output := &workspacesweb.CreateIdentityProviderOutput{}

        mockClient.On("CreateIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.CreateIpAccessSettingsInput{}
        output := &workspacesweb.CreateIpAccessSettingsOutput{}

        mockClient.On("CreateIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.CreateNetworkSettingsInput{}
        output := &workspacesweb.CreateNetworkSettingsOutput{}

        mockClient.On("CreateNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePortal", func(t *testing.T) {
        input := &workspacesweb.CreatePortalInput{}
        output := &workspacesweb.CreatePortalOutput{}

        mockClient.On("CreatePortal", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrustStore", func(t *testing.T) {
        input := &workspacesweb.CreateTrustStoreInput{}
        output := &workspacesweb.CreateTrustStoreOutput{}

        mockClient.On("CreateTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.CreateUserAccessLoggingSettingsInput{}
        output := &workspacesweb.CreateUserAccessLoggingSettingsOutput{}

        mockClient.On("CreateUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserSettings", func(t *testing.T) {
        input := &workspacesweb.CreateUserSettingsInput{}
        output := &workspacesweb.CreateUserSettingsOutput{}

        mockClient.On("CreateUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.DeleteBrowserSettingsInput{}
        output := &workspacesweb.DeleteBrowserSettingsOutput{}

        mockClient.On("DeleteBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentityProvider", func(t *testing.T) {
        input := &workspacesweb.DeleteIdentityProviderInput{}
        output := &workspacesweb.DeleteIdentityProviderOutput{}

        mockClient.On("DeleteIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.DeleteIpAccessSettingsInput{}
        output := &workspacesweb.DeleteIpAccessSettingsOutput{}

        mockClient.On("DeleteIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.DeleteNetworkSettingsInput{}
        output := &workspacesweb.DeleteNetworkSettingsOutput{}

        mockClient.On("DeleteNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePortal", func(t *testing.T) {
        input := &workspacesweb.DeletePortalInput{}
        output := &workspacesweb.DeletePortalOutput{}

        mockClient.On("DeletePortal", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrustStore", func(t *testing.T) {
        input := &workspacesweb.DeleteTrustStoreInput{}
        output := &workspacesweb.DeleteTrustStoreOutput{}

        mockClient.On("DeleteTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.DeleteUserAccessLoggingSettingsInput{}
        output := &workspacesweb.DeleteUserAccessLoggingSettingsOutput{}

        mockClient.On("DeleteUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserSettings", func(t *testing.T) {
        input := &workspacesweb.DeleteUserSettingsInput{}
        output := &workspacesweb.DeleteUserSettingsOutput{}

        mockClient.On("DeleteUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.DisassociateBrowserSettingsInput{}
        output := &workspacesweb.DisassociateBrowserSettingsOutput{}

        mockClient.On("DisassociateBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.DisassociateIpAccessSettingsInput{}
        output := &workspacesweb.DisassociateIpAccessSettingsOutput{}

        mockClient.On("DisassociateIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.DisassociateNetworkSettingsInput{}
        output := &workspacesweb.DisassociateNetworkSettingsOutput{}

        mockClient.On("DisassociateNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTrustStore", func(t *testing.T) {
        input := &workspacesweb.DisassociateTrustStoreInput{}
        output := &workspacesweb.DisassociateTrustStoreOutput{}

        mockClient.On("DisassociateTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.DisassociateUserAccessLoggingSettingsInput{}
        output := &workspacesweb.DisassociateUserAccessLoggingSettingsOutput{}

        mockClient.On("DisassociateUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateUserSettings", func(t *testing.T) {
        input := &workspacesweb.DisassociateUserSettingsInput{}
        output := &workspacesweb.DisassociateUserSettingsOutput{}

        mockClient.On("DisassociateUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.GetBrowserSettingsInput{}
        output := &workspacesweb.GetBrowserSettingsOutput{}

        mockClient.On("GetBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityProvider", func(t *testing.T) {
        input := &workspacesweb.GetIdentityProviderInput{}
        output := &workspacesweb.GetIdentityProviderOutput{}

        mockClient.On("GetIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.GetIpAccessSettingsInput{}
        output := &workspacesweb.GetIpAccessSettingsOutput{}

        mockClient.On("GetIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.GetNetworkSettingsInput{}
        output := &workspacesweb.GetNetworkSettingsOutput{}

        mockClient.On("GetNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPortal", func(t *testing.T) {
        input := &workspacesweb.GetPortalInput{}
        output := &workspacesweb.GetPortalOutput{}

        mockClient.On("GetPortal", ctx, input).Return(output, nil)

        result, err := mockClient.GetPortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPortalServiceProviderMetadata", func(t *testing.T) {
        input := &workspacesweb.GetPortalServiceProviderMetadataInput{}
        output := &workspacesweb.GetPortalServiceProviderMetadataOutput{}

        mockClient.On("GetPortalServiceProviderMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetPortalServiceProviderMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrustStore", func(t *testing.T) {
        input := &workspacesweb.GetTrustStoreInput{}
        output := &workspacesweb.GetTrustStoreOutput{}

        mockClient.On("GetTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrustStoreCertificate", func(t *testing.T) {
        input := &workspacesweb.GetTrustStoreCertificateInput{}
        output := &workspacesweb.GetTrustStoreCertificateOutput{}

        mockClient.On("GetTrustStoreCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrustStoreCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.GetUserAccessLoggingSettingsInput{}
        output := &workspacesweb.GetUserAccessLoggingSettingsOutput{}

        mockClient.On("GetUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserSettings", func(t *testing.T) {
        input := &workspacesweb.GetUserSettingsInput{}
        output := &workspacesweb.GetUserSettingsOutput{}

        mockClient.On("GetUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.ListBrowserSettingsInput{}
        output := &workspacesweb.ListBrowserSettingsOutput{}

        mockClient.On("ListBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityProviders", func(t *testing.T) {
        input := &workspacesweb.ListIdentityProvidersInput{}
        output := &workspacesweb.ListIdentityProvidersOutput{}

        mockClient.On("ListIdentityProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.ListIpAccessSettingsInput{}
        output := &workspacesweb.ListIpAccessSettingsOutput{}

        mockClient.On("ListIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.ListNetworkSettingsInput{}
        output := &workspacesweb.ListNetworkSettingsOutput{}

        mockClient.On("ListNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPortals", func(t *testing.T) {
        input := &workspacesweb.ListPortalsInput{}
        output := &workspacesweb.ListPortalsOutput{}

        mockClient.On("ListPortals", ctx, input).Return(output, nil)

        result, err := mockClient.ListPortals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &workspacesweb.ListTagsForResourceInput{}
        output := &workspacesweb.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrustStoreCertificates", func(t *testing.T) {
        input := &workspacesweb.ListTrustStoreCertificatesInput{}
        output := &workspacesweb.ListTrustStoreCertificatesOutput{}

        mockClient.On("ListTrustStoreCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrustStoreCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrustStores", func(t *testing.T) {
        input := &workspacesweb.ListTrustStoresInput{}
        output := &workspacesweb.ListTrustStoresOutput{}

        mockClient.On("ListTrustStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrustStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.ListUserAccessLoggingSettingsInput{}
        output := &workspacesweb.ListUserAccessLoggingSettingsOutput{}

        mockClient.On("ListUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserSettings", func(t *testing.T) {
        input := &workspacesweb.ListUserSettingsInput{}
        output := &workspacesweb.ListUserSettingsOutput{}

        mockClient.On("ListUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &workspacesweb.TagResourceInput{}
        output := &workspacesweb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &workspacesweb.UntagResourceInput{}
        output := &workspacesweb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBrowserSettings", func(t *testing.T) {
        input := &workspacesweb.UpdateBrowserSettingsInput{}
        output := &workspacesweb.UpdateBrowserSettingsOutput{}

        mockClient.On("UpdateBrowserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBrowserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityProvider", func(t *testing.T) {
        input := &workspacesweb.UpdateIdentityProviderInput{}
        output := &workspacesweb.UpdateIdentityProviderOutput{}

        mockClient.On("UpdateIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIpAccessSettings", func(t *testing.T) {
        input := &workspacesweb.UpdateIpAccessSettingsInput{}
        output := &workspacesweb.UpdateIpAccessSettingsOutput{}

        mockClient.On("UpdateIpAccessSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIpAccessSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkSettings", func(t *testing.T) {
        input := &workspacesweb.UpdateNetworkSettingsInput{}
        output := &workspacesweb.UpdateNetworkSettingsOutput{}

        mockClient.On("UpdateNetworkSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePortal", func(t *testing.T) {
        input := &workspacesweb.UpdatePortalInput{}
        output := &workspacesweb.UpdatePortalOutput{}

        mockClient.On("UpdatePortal", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrustStore", func(t *testing.T) {
        input := &workspacesweb.UpdateTrustStoreInput{}
        output := &workspacesweb.UpdateTrustStoreOutput{}

        mockClient.On("UpdateTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserAccessLoggingSettings", func(t *testing.T) {
        input := &workspacesweb.UpdateUserAccessLoggingSettingsInput{}
        output := &workspacesweb.UpdateUserAccessLoggingSettingsOutput{}

        mockClient.On("UpdateUserAccessLoggingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserAccessLoggingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserSettings", func(t *testing.T) {
        input := &workspacesweb.UpdateUserSettingsInput{}
        output := &workspacesweb.UpdateUserSettingsOutput{}

        mockClient.On("UpdateUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
