package wafv2_test

// tests for the wafv2 service interface for this ../../../service/wafv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wafv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wafv2/wafv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestWafv2ServiceCanBeMocked(t *testing.T) {
	var iface wafv2_iface.IClient
	iface = &wafv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := wafv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWebACL", func(t *testing.T) {
        input := &wafv2.AssociateWebACLInput{}
        output := &wafv2.AssociateWebACLOutput{}

        mockClient.On("AssociateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckCapacity", func(t *testing.T) {
        input := &wafv2.CheckCapacityInput{}
        output := &wafv2.CheckCapacityOutput{}

        mockClient.On("CheckCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.CheckCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAPIKey", func(t *testing.T) {
        input := &wafv2.CreateAPIKeyInput{}
        output := &wafv2.CreateAPIKeyOutput{}

        mockClient.On("CreateAPIKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAPIKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIPSet", func(t *testing.T) {
        input := &wafv2.CreateIPSetInput{}
        output := &wafv2.CreateIPSetOutput{}

        mockClient.On("CreateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegexPatternSet", func(t *testing.T) {
        input := &wafv2.CreateRegexPatternSetInput{}
        output := &wafv2.CreateRegexPatternSetOutput{}

        mockClient.On("CreateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleGroup", func(t *testing.T) {
        input := &wafv2.CreateRuleGroupInput{}
        output := &wafv2.CreateRuleGroupOutput{}

        mockClient.On("CreateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebACL", func(t *testing.T) {
        input := &wafv2.CreateWebACLInput{}
        output := &wafv2.CreateWebACLOutput{}

        mockClient.On("CreateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAPIKey", func(t *testing.T) {
        input := &wafv2.DeleteAPIKeyInput{}
        output := &wafv2.DeleteAPIKeyOutput{}

        mockClient.On("DeleteAPIKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAPIKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewallManagerRuleGroups", func(t *testing.T) {
        input := &wafv2.DeleteFirewallManagerRuleGroupsInput{}
        output := &wafv2.DeleteFirewallManagerRuleGroupsOutput{}

        mockClient.On("DeleteFirewallManagerRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewallManagerRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIPSet", func(t *testing.T) {
        input := &wafv2.DeleteIPSetInput{}
        output := &wafv2.DeleteIPSetOutput{}

        mockClient.On("DeleteIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggingConfiguration", func(t *testing.T) {
        input := &wafv2.DeleteLoggingConfigurationInput{}
        output := &wafv2.DeleteLoggingConfigurationOutput{}

        mockClient.On("DeleteLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionPolicy", func(t *testing.T) {
        input := &wafv2.DeletePermissionPolicyInput{}
        output := &wafv2.DeletePermissionPolicyOutput{}

        mockClient.On("DeletePermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegexPatternSet", func(t *testing.T) {
        input := &wafv2.DeleteRegexPatternSetInput{}
        output := &wafv2.DeleteRegexPatternSetOutput{}

        mockClient.On("DeleteRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleGroup", func(t *testing.T) {
        input := &wafv2.DeleteRuleGroupInput{}
        output := &wafv2.DeleteRuleGroupOutput{}

        mockClient.On("DeleteRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebACL", func(t *testing.T) {
        input := &wafv2.DeleteWebACLInput{}
        output := &wafv2.DeleteWebACLOutput{}

        mockClient.On("DeleteWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAllManagedProducts", func(t *testing.T) {
        input := &wafv2.DescribeAllManagedProductsInput{}
        output := &wafv2.DescribeAllManagedProductsOutput{}

        mockClient.On("DescribeAllManagedProducts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAllManagedProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeManagedProductsByVendor", func(t *testing.T) {
        input := &wafv2.DescribeManagedProductsByVendorInput{}
        output := &wafv2.DescribeManagedProductsByVendorOutput{}

        mockClient.On("DescribeManagedProductsByVendor", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeManagedProductsByVendor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeManagedRuleGroup", func(t *testing.T) {
        input := &wafv2.DescribeManagedRuleGroupInput{}
        output := &wafv2.DescribeManagedRuleGroupOutput{}

        mockClient.On("DescribeManagedRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeManagedRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWebACL", func(t *testing.T) {
        input := &wafv2.DisassociateWebACLInput{}
        output := &wafv2.DisassociateWebACLOutput{}

        mockClient.On("DisassociateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateMobileSdkReleaseUrl", func(t *testing.T) {
        input := &wafv2.GenerateMobileSdkReleaseUrlInput{}
        output := &wafv2.GenerateMobileSdkReleaseUrlOutput{}

        mockClient.On("GenerateMobileSdkReleaseUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateMobileSdkReleaseUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDecryptedAPIKey", func(t *testing.T) {
        input := &wafv2.GetDecryptedAPIKeyInput{}
        output := &wafv2.GetDecryptedAPIKeyOutput{}

        mockClient.On("GetDecryptedAPIKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetDecryptedAPIKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIPSet", func(t *testing.T) {
        input := &wafv2.GetIPSetInput{}
        output := &wafv2.GetIPSetOutput{}

        mockClient.On("GetIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingConfiguration", func(t *testing.T) {
        input := &wafv2.GetLoggingConfigurationInput{}
        output := &wafv2.GetLoggingConfigurationOutput{}

        mockClient.On("GetLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedRuleSet", func(t *testing.T) {
        input := &wafv2.GetManagedRuleSetInput{}
        output := &wafv2.GetManagedRuleSetOutput{}

        mockClient.On("GetManagedRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMobileSdkRelease", func(t *testing.T) {
        input := &wafv2.GetMobileSdkReleaseInput{}
        output := &wafv2.GetMobileSdkReleaseOutput{}

        mockClient.On("GetMobileSdkRelease", ctx, input).Return(output, nil)

        result, err := mockClient.GetMobileSdkRelease(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermissionPolicy", func(t *testing.T) {
        input := &wafv2.GetPermissionPolicyInput{}
        output := &wafv2.GetPermissionPolicyOutput{}

        mockClient.On("GetPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRateBasedStatementManagedKeys", func(t *testing.T) {
        input := &wafv2.GetRateBasedStatementManagedKeysInput{}
        output := &wafv2.GetRateBasedStatementManagedKeysOutput{}

        mockClient.On("GetRateBasedStatementManagedKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetRateBasedStatementManagedKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegexPatternSet", func(t *testing.T) {
        input := &wafv2.GetRegexPatternSetInput{}
        output := &wafv2.GetRegexPatternSetOutput{}

        mockClient.On("GetRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRuleGroup", func(t *testing.T) {
        input := &wafv2.GetRuleGroupInput{}
        output := &wafv2.GetRuleGroupOutput{}

        mockClient.On("GetRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSampledRequests", func(t *testing.T) {
        input := &wafv2.GetSampledRequestsInput{}
        output := &wafv2.GetSampledRequestsOutput{}

        mockClient.On("GetSampledRequests", ctx, input).Return(output, nil)

        result, err := mockClient.GetSampledRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebACL", func(t *testing.T) {
        input := &wafv2.GetWebACLInput{}
        output := &wafv2.GetWebACLOutput{}

        mockClient.On("GetWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebACLForResource", func(t *testing.T) {
        input := &wafv2.GetWebACLForResourceInput{}
        output := &wafv2.GetWebACLForResourceOutput{}

        mockClient.On("GetWebACLForResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebACLForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAPIKeys", func(t *testing.T) {
        input := &wafv2.ListAPIKeysInput{}
        output := &wafv2.ListAPIKeysOutput{}

        mockClient.On("ListAPIKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListAPIKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableManagedRuleGroupVersions", func(t *testing.T) {
        input := &wafv2.ListAvailableManagedRuleGroupVersionsInput{}
        output := &wafv2.ListAvailableManagedRuleGroupVersionsOutput{}

        mockClient.On("ListAvailableManagedRuleGroupVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableManagedRuleGroupVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableManagedRuleGroups", func(t *testing.T) {
        input := &wafv2.ListAvailableManagedRuleGroupsInput{}
        output := &wafv2.ListAvailableManagedRuleGroupsOutput{}

        mockClient.On("ListAvailableManagedRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableManagedRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIPSets", func(t *testing.T) {
        input := &wafv2.ListIPSetsInput{}
        output := &wafv2.ListIPSetsOutput{}

        mockClient.On("ListIPSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListIPSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggingConfigurations", func(t *testing.T) {
        input := &wafv2.ListLoggingConfigurationsInput{}
        output := &wafv2.ListLoggingConfigurationsOutput{}

        mockClient.On("ListLoggingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedRuleSets", func(t *testing.T) {
        input := &wafv2.ListManagedRuleSetsInput{}
        output := &wafv2.ListManagedRuleSetsOutput{}

        mockClient.On("ListManagedRuleSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedRuleSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMobileSdkReleases", func(t *testing.T) {
        input := &wafv2.ListMobileSdkReleasesInput{}
        output := &wafv2.ListMobileSdkReleasesOutput{}

        mockClient.On("ListMobileSdkReleases", ctx, input).Return(output, nil)

        result, err := mockClient.ListMobileSdkReleases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegexPatternSets", func(t *testing.T) {
        input := &wafv2.ListRegexPatternSetsInput{}
        output := &wafv2.ListRegexPatternSetsOutput{}

        mockClient.On("ListRegexPatternSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegexPatternSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcesForWebACL", func(t *testing.T) {
        input := &wafv2.ListResourcesForWebACLInput{}
        output := &wafv2.ListResourcesForWebACLOutput{}

        mockClient.On("ListResourcesForWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcesForWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleGroups", func(t *testing.T) {
        input := &wafv2.ListRuleGroupsInput{}
        output := &wafv2.ListRuleGroupsOutput{}

        mockClient.On("ListRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &wafv2.ListTagsForResourceInput{}
        output := &wafv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebACLs", func(t *testing.T) {
        input := &wafv2.ListWebACLsInput{}
        output := &wafv2.ListWebACLsOutput{}

        mockClient.On("ListWebACLs", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebACLs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingConfiguration", func(t *testing.T) {
        input := &wafv2.PutLoggingConfigurationInput{}
        output := &wafv2.PutLoggingConfigurationOutput{}

        mockClient.On("PutLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutManagedRuleSetVersions", func(t *testing.T) {
        input := &wafv2.PutManagedRuleSetVersionsInput{}
        output := &wafv2.PutManagedRuleSetVersionsOutput{}

        mockClient.On("PutManagedRuleSetVersions", ctx, input).Return(output, nil)

        result, err := mockClient.PutManagedRuleSetVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermissionPolicy", func(t *testing.T) {
        input := &wafv2.PutPermissionPolicyInput{}
        output := &wafv2.PutPermissionPolicyOutput{}

        mockClient.On("PutPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &wafv2.TagResourceInput{}
        output := &wafv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &wafv2.UntagResourceInput{}
        output := &wafv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIPSet", func(t *testing.T) {
        input := &wafv2.UpdateIPSetInput{}
        output := &wafv2.UpdateIPSetOutput{}

        mockClient.On("UpdateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateManagedRuleSetVersionExpiryDate", func(t *testing.T) {
        input := &wafv2.UpdateManagedRuleSetVersionExpiryDateInput{}
        output := &wafv2.UpdateManagedRuleSetVersionExpiryDateOutput{}

        mockClient.On("UpdateManagedRuleSetVersionExpiryDate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateManagedRuleSetVersionExpiryDate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegexPatternSet", func(t *testing.T) {
        input := &wafv2.UpdateRegexPatternSetInput{}
        output := &wafv2.UpdateRegexPatternSetOutput{}

        mockClient.On("UpdateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleGroup", func(t *testing.T) {
        input := &wafv2.UpdateRuleGroupInput{}
        output := &wafv2.UpdateRuleGroupOutput{}

        mockClient.On("UpdateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebACL", func(t *testing.T) {
        input := &wafv2.UpdateWebACLInput{}
        output := &wafv2.UpdateWebACLOutput{}

        mockClient.On("UpdateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
