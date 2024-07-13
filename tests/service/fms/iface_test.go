package fms_test

// tests for the fms service interface for this ../../../service/fms/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fms/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fms/fms_iface"
	"github.com/stretchr/testify/assert"
)

func TestFmsServiceCanBeMocked(t *testing.T) {
	var iface fms_iface.IClient
	iface = &fms.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := fms.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAdminAccount", func(t *testing.T) {
        input := &fms.AssociateAdminAccountInput{}
        output := &fms.AssociateAdminAccountOutput{}

        mockClient.On("AssociateAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateThirdPartyFirewall", func(t *testing.T) {
        input := &fms.AssociateThirdPartyFirewallInput{}
        output := &fms.AssociateThirdPartyFirewallOutput{}

        mockClient.On("AssociateThirdPartyFirewall", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateThirdPartyFirewall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateResource", func(t *testing.T) {
        input := &fms.BatchAssociateResourceInput{}
        output := &fms.BatchAssociateResourceOutput{}

        mockClient.On("BatchAssociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateResource", func(t *testing.T) {
        input := &fms.BatchDisassociateResourceInput{}
        output := &fms.BatchDisassociateResourceOutput{}

        mockClient.On("BatchDisassociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppsList", func(t *testing.T) {
        input := &fms.DeleteAppsListInput{}
        output := &fms.DeleteAppsListOutput{}

        mockClient.On("DeleteAppsList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotificationChannel", func(t *testing.T) {
        input := &fms.DeleteNotificationChannelInput{}
        output := &fms.DeleteNotificationChannelOutput{}

        mockClient.On("DeleteNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &fms.DeletePolicyInput{}
        output := &fms.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProtocolsList", func(t *testing.T) {
        input := &fms.DeleteProtocolsListInput{}
        output := &fms.DeleteProtocolsListOutput{}

        mockClient.On("DeleteProtocolsList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProtocolsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceSet", func(t *testing.T) {
        input := &fms.DeleteResourceSetInput{}
        output := &fms.DeleteResourceSetOutput{}

        mockClient.On("DeleteResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAdminAccount", func(t *testing.T) {
        input := &fms.DisassociateAdminAccountInput{}
        output := &fms.DisassociateAdminAccountOutput{}

        mockClient.On("DisassociateAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateThirdPartyFirewall", func(t *testing.T) {
        input := &fms.DisassociateThirdPartyFirewallInput{}
        output := &fms.DisassociateThirdPartyFirewallOutput{}

        mockClient.On("DisassociateThirdPartyFirewall", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateThirdPartyFirewall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdminAccount", func(t *testing.T) {
        input := &fms.GetAdminAccountInput{}
        output := &fms.GetAdminAccountOutput{}

        mockClient.On("GetAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdminScope", func(t *testing.T) {
        input := &fms.GetAdminScopeInput{}
        output := &fms.GetAdminScopeOutput{}

        mockClient.On("GetAdminScope", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdminScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppsList", func(t *testing.T) {
        input := &fms.GetAppsListInput{}
        output := &fms.GetAppsListOutput{}

        mockClient.On("GetAppsList", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceDetail", func(t *testing.T) {
        input := &fms.GetComplianceDetailInput{}
        output := &fms.GetComplianceDetailOutput{}

        mockClient.On("GetComplianceDetail", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceDetail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotificationChannel", func(t *testing.T) {
        input := &fms.GetNotificationChannelInput{}
        output := &fms.GetNotificationChannelOutput{}

        mockClient.On("GetNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &fms.GetPolicyInput{}
        output := &fms.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProtectionStatus", func(t *testing.T) {
        input := &fms.GetProtectionStatusInput{}
        output := &fms.GetProtectionStatusOutput{}

        mockClient.On("GetProtectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetProtectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProtocolsList", func(t *testing.T) {
        input := &fms.GetProtocolsListInput{}
        output := &fms.GetProtocolsListOutput{}

        mockClient.On("GetProtocolsList", ctx, input).Return(output, nil)

        result, err := mockClient.GetProtocolsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceSet", func(t *testing.T) {
        input := &fms.GetResourceSetInput{}
        output := &fms.GetResourceSetOutput{}

        mockClient.On("GetResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetThirdPartyFirewallAssociationStatus", func(t *testing.T) {
        input := &fms.GetThirdPartyFirewallAssociationStatusInput{}
        output := &fms.GetThirdPartyFirewallAssociationStatusOutput{}

        mockClient.On("GetThirdPartyFirewallAssociationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetThirdPartyFirewallAssociationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetViolationDetails", func(t *testing.T) {
        input := &fms.GetViolationDetailsInput{}
        output := &fms.GetViolationDetailsOutput{}

        mockClient.On("GetViolationDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetViolationDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAdminAccountsForOrganization", func(t *testing.T) {
        input := &fms.ListAdminAccountsForOrganizationInput{}
        output := &fms.ListAdminAccountsForOrganizationOutput{}

        mockClient.On("ListAdminAccountsForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.ListAdminAccountsForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAdminsManagingAccount", func(t *testing.T) {
        input := &fms.ListAdminsManagingAccountInput{}
        output := &fms.ListAdminsManagingAccountOutput{}

        mockClient.On("ListAdminsManagingAccount", ctx, input).Return(output, nil)

        result, err := mockClient.ListAdminsManagingAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppsLists", func(t *testing.T) {
        input := &fms.ListAppsListsInput{}
        output := &fms.ListAppsListsOutput{}

        mockClient.On("ListAppsLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppsLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComplianceStatus", func(t *testing.T) {
        input := &fms.ListComplianceStatusInput{}
        output := &fms.ListComplianceStatusOutput{}

        mockClient.On("ListComplianceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListComplianceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDiscoveredResources", func(t *testing.T) {
        input := &fms.ListDiscoveredResourcesInput{}
        output := &fms.ListDiscoveredResourcesOutput{}

        mockClient.On("ListDiscoveredResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDiscoveredResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMemberAccounts", func(t *testing.T) {
        input := &fms.ListMemberAccountsInput{}
        output := &fms.ListMemberAccountsOutput{}

        mockClient.On("ListMemberAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListMemberAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicies", func(t *testing.T) {
        input := &fms.ListPoliciesInput{}
        output := &fms.ListPoliciesOutput{}

        mockClient.On("ListPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtocolsLists", func(t *testing.T) {
        input := &fms.ListProtocolsListsInput{}
        output := &fms.ListProtocolsListsOutput{}

        mockClient.On("ListProtocolsLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtocolsLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceSetResources", func(t *testing.T) {
        input := &fms.ListResourceSetResourcesInput{}
        output := &fms.ListResourceSetResourcesOutput{}

        mockClient.On("ListResourceSetResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceSetResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceSets", func(t *testing.T) {
        input := &fms.ListResourceSetsInput{}
        output := &fms.ListResourceSetsOutput{}

        mockClient.On("ListResourceSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &fms.ListTagsForResourceInput{}
        output := &fms.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThirdPartyFirewallFirewallPolicies", func(t *testing.T) {
        input := &fms.ListThirdPartyFirewallFirewallPoliciesInput{}
        output := &fms.ListThirdPartyFirewallFirewallPoliciesOutput{}

        mockClient.On("ListThirdPartyFirewallFirewallPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListThirdPartyFirewallFirewallPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAdminAccount", func(t *testing.T) {
        input := &fms.PutAdminAccountInput{}
        output := &fms.PutAdminAccountOutput{}

        mockClient.On("PutAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.PutAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppsList", func(t *testing.T) {
        input := &fms.PutAppsListInput{}
        output := &fms.PutAppsListOutput{}

        mockClient.On("PutAppsList", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutNotificationChannel", func(t *testing.T) {
        input := &fms.PutNotificationChannelInput{}
        output := &fms.PutNotificationChannelOutput{}

        mockClient.On("PutNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.PutNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPolicy", func(t *testing.T) {
        input := &fms.PutPolicyInput{}
        output := &fms.PutPolicyOutput{}

        mockClient.On("PutPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProtocolsList", func(t *testing.T) {
        input := &fms.PutProtocolsListInput{}
        output := &fms.PutProtocolsListOutput{}

        mockClient.On("PutProtocolsList", ctx, input).Return(output, nil)

        result, err := mockClient.PutProtocolsList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourceSet", func(t *testing.T) {
        input := &fms.PutResourceSetInput{}
        output := &fms.PutResourceSetOutput{}

        mockClient.On("PutResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &fms.TagResourceInput{}
        output := &fms.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &fms.UntagResourceInput{}
        output := &fms.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
