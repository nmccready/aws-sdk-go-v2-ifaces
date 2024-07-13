package eks_test

// tests for the eks service interface for this ../../../service/eks/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eks/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eks/eks_iface"
	"github.com/stretchr/testify/assert"
)

func TestEksServiceCanBeMocked(t *testing.T) {
	var iface eks_iface.IClient
	iface = &eks.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := eks.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAccessPolicy", func(t *testing.T) {
        input := &eks.AssociateAccessPolicyInput{}
        output := &eks.AssociateAccessPolicyOutput{}

        mockClient.On("AssociateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEncryptionConfig", func(t *testing.T) {
        input := &eks.AssociateEncryptionConfigInput{}
        output := &eks.AssociateEncryptionConfigOutput{}

        mockClient.On("AssociateEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIdentityProviderConfig", func(t *testing.T) {
        input := &eks.AssociateIdentityProviderConfigInput{}
        output := &eks.AssociateIdentityProviderConfigOutput{}

        mockClient.On("AssociateIdentityProviderConfig", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIdentityProviderConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessEntry", func(t *testing.T) {
        input := &eks.CreateAccessEntryInput{}
        output := &eks.CreateAccessEntryOutput{}

        mockClient.On("CreateAccessEntry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAddon", func(t *testing.T) {
        input := &eks.CreateAddonInput{}
        output := &eks.CreateAddonOutput{}

        mockClient.On("CreateAddon", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAddon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &eks.CreateClusterInput{}
        output := &eks.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEksAnywhereSubscription", func(t *testing.T) {
        input := &eks.CreateEksAnywhereSubscriptionInput{}
        output := &eks.CreateEksAnywhereSubscriptionOutput{}

        mockClient.On("CreateEksAnywhereSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEksAnywhereSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFargateProfile", func(t *testing.T) {
        input := &eks.CreateFargateProfileInput{}
        output := &eks.CreateFargateProfileOutput{}

        mockClient.On("CreateFargateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFargateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNodegroup", func(t *testing.T) {
        input := &eks.CreateNodegroupInput{}
        output := &eks.CreateNodegroupOutput{}

        mockClient.On("CreateNodegroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNodegroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePodIdentityAssociation", func(t *testing.T) {
        input := &eks.CreatePodIdentityAssociationInput{}
        output := &eks.CreatePodIdentityAssociationOutput{}

        mockClient.On("CreatePodIdentityAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePodIdentityAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessEntry", func(t *testing.T) {
        input := &eks.DeleteAccessEntryInput{}
        output := &eks.DeleteAccessEntryOutput{}

        mockClient.On("DeleteAccessEntry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAddon", func(t *testing.T) {
        input := &eks.DeleteAddonInput{}
        output := &eks.DeleteAddonOutput{}

        mockClient.On("DeleteAddon", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAddon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &eks.DeleteClusterInput{}
        output := &eks.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEksAnywhereSubscription", func(t *testing.T) {
        input := &eks.DeleteEksAnywhereSubscriptionInput{}
        output := &eks.DeleteEksAnywhereSubscriptionOutput{}

        mockClient.On("DeleteEksAnywhereSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEksAnywhereSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFargateProfile", func(t *testing.T) {
        input := &eks.DeleteFargateProfileInput{}
        output := &eks.DeleteFargateProfileOutput{}

        mockClient.On("DeleteFargateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFargateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNodegroup", func(t *testing.T) {
        input := &eks.DeleteNodegroupInput{}
        output := &eks.DeleteNodegroupOutput{}

        mockClient.On("DeleteNodegroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNodegroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePodIdentityAssociation", func(t *testing.T) {
        input := &eks.DeletePodIdentityAssociationInput{}
        output := &eks.DeletePodIdentityAssociationOutput{}

        mockClient.On("DeletePodIdentityAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePodIdentityAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterCluster", func(t *testing.T) {
        input := &eks.DeregisterClusterInput{}
        output := &eks.DeregisterClusterOutput{}

        mockClient.On("DeregisterCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccessEntry", func(t *testing.T) {
        input := &eks.DescribeAccessEntryInput{}
        output := &eks.DescribeAccessEntryOutput{}

        mockClient.On("DescribeAccessEntry", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccessEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddon", func(t *testing.T) {
        input := &eks.DescribeAddonInput{}
        output := &eks.DescribeAddonOutput{}

        mockClient.On("DescribeAddon", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddonConfiguration", func(t *testing.T) {
        input := &eks.DescribeAddonConfigurationInput{}
        output := &eks.DescribeAddonConfigurationOutput{}

        mockClient.On("DescribeAddonConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddonConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddonVersions", func(t *testing.T) {
        input := &eks.DescribeAddonVersionsInput{}
        output := &eks.DescribeAddonVersionsOutput{}

        mockClient.On("DescribeAddonVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddonVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &eks.DescribeClusterInput{}
        output := &eks.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEksAnywhereSubscription", func(t *testing.T) {
        input := &eks.DescribeEksAnywhereSubscriptionInput{}
        output := &eks.DescribeEksAnywhereSubscriptionOutput{}

        mockClient.On("DescribeEksAnywhereSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEksAnywhereSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFargateProfile", func(t *testing.T) {
        input := &eks.DescribeFargateProfileInput{}
        output := &eks.DescribeFargateProfileOutput{}

        mockClient.On("DescribeFargateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFargateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityProviderConfig", func(t *testing.T) {
        input := &eks.DescribeIdentityProviderConfigInput{}
        output := &eks.DescribeIdentityProviderConfigOutput{}

        mockClient.On("DescribeIdentityProviderConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityProviderConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInsight", func(t *testing.T) {
        input := &eks.DescribeInsightInput{}
        output := &eks.DescribeInsightOutput{}

        mockClient.On("DescribeInsight", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNodegroup", func(t *testing.T) {
        input := &eks.DescribeNodegroupInput{}
        output := &eks.DescribeNodegroupOutput{}

        mockClient.On("DescribeNodegroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNodegroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePodIdentityAssociation", func(t *testing.T) {
        input := &eks.DescribePodIdentityAssociationInput{}
        output := &eks.DescribePodIdentityAssociationOutput{}

        mockClient.On("DescribePodIdentityAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePodIdentityAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUpdate", func(t *testing.T) {
        input := &eks.DescribeUpdateInput{}
        output := &eks.DescribeUpdateOutput{}

        mockClient.On("DescribeUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAccessPolicy", func(t *testing.T) {
        input := &eks.DisassociateAccessPolicyInput{}
        output := &eks.DisassociateAccessPolicyOutput{}

        mockClient.On("DisassociateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIdentityProviderConfig", func(t *testing.T) {
        input := &eks.DisassociateIdentityProviderConfigInput{}
        output := &eks.DisassociateIdentityProviderConfigOutput{}

        mockClient.On("DisassociateIdentityProviderConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIdentityProviderConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessEntries", func(t *testing.T) {
        input := &eks.ListAccessEntriesInput{}
        output := &eks.ListAccessEntriesOutput{}

        mockClient.On("ListAccessEntries", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPolicies", func(t *testing.T) {
        input := &eks.ListAccessPoliciesInput{}
        output := &eks.ListAccessPoliciesOutput{}

        mockClient.On("ListAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAddons", func(t *testing.T) {
        input := &eks.ListAddonsInput{}
        output := &eks.ListAddonsOutput{}

        mockClient.On("ListAddons", ctx, input).Return(output, nil)

        result, err := mockClient.ListAddons(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedAccessPolicies", func(t *testing.T) {
        input := &eks.ListAssociatedAccessPoliciesInput{}
        output := &eks.ListAssociatedAccessPoliciesOutput{}

        mockClient.On("ListAssociatedAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &eks.ListClustersInput{}
        output := &eks.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEksAnywhereSubscriptions", func(t *testing.T) {
        input := &eks.ListEksAnywhereSubscriptionsInput{}
        output := &eks.ListEksAnywhereSubscriptionsOutput{}

        mockClient.On("ListEksAnywhereSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEksAnywhereSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFargateProfiles", func(t *testing.T) {
        input := &eks.ListFargateProfilesInput{}
        output := &eks.ListFargateProfilesOutput{}

        mockClient.On("ListFargateProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListFargateProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityProviderConfigs", func(t *testing.T) {
        input := &eks.ListIdentityProviderConfigsInput{}
        output := &eks.ListIdentityProviderConfigsOutput{}

        mockClient.On("ListIdentityProviderConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityProviderConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInsights", func(t *testing.T) {
        input := &eks.ListInsightsInput{}
        output := &eks.ListInsightsOutput{}

        mockClient.On("ListInsights", ctx, input).Return(output, nil)

        result, err := mockClient.ListInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNodegroups", func(t *testing.T) {
        input := &eks.ListNodegroupsInput{}
        output := &eks.ListNodegroupsOutput{}

        mockClient.On("ListNodegroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListNodegroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPodIdentityAssociations", func(t *testing.T) {
        input := &eks.ListPodIdentityAssociationsInput{}
        output := &eks.ListPodIdentityAssociationsOutput{}

        mockClient.On("ListPodIdentityAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPodIdentityAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &eks.ListTagsForResourceInput{}
        output := &eks.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUpdates", func(t *testing.T) {
        input := &eks.ListUpdatesInput{}
        output := &eks.ListUpdatesOutput{}

        mockClient.On("ListUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.ListUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCluster", func(t *testing.T) {
        input := &eks.RegisterClusterInput{}
        output := &eks.RegisterClusterOutput{}

        mockClient.On("RegisterCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &eks.TagResourceInput{}
        output := &eks.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &eks.UntagResourceInput{}
        output := &eks.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessEntry", func(t *testing.T) {
        input := &eks.UpdateAccessEntryInput{}
        output := &eks.UpdateAccessEntryOutput{}

        mockClient.On("UpdateAccessEntry", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAddon", func(t *testing.T) {
        input := &eks.UpdateAddonInput{}
        output := &eks.UpdateAddonOutput{}

        mockClient.On("UpdateAddon", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAddon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterConfig", func(t *testing.T) {
        input := &eks.UpdateClusterConfigInput{}
        output := &eks.UpdateClusterConfigOutput{}

        mockClient.On("UpdateClusterConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterVersion", func(t *testing.T) {
        input := &eks.UpdateClusterVersionInput{}
        output := &eks.UpdateClusterVersionOutput{}

        mockClient.On("UpdateClusterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEksAnywhereSubscription", func(t *testing.T) {
        input := &eks.UpdateEksAnywhereSubscriptionInput{}
        output := &eks.UpdateEksAnywhereSubscriptionOutput{}

        mockClient.On("UpdateEksAnywhereSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEksAnywhereSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNodegroupConfig", func(t *testing.T) {
        input := &eks.UpdateNodegroupConfigInput{}
        output := &eks.UpdateNodegroupConfigOutput{}

        mockClient.On("UpdateNodegroupConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNodegroupConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNodegroupVersion", func(t *testing.T) {
        input := &eks.UpdateNodegroupVersionInput{}
        output := &eks.UpdateNodegroupVersionOutput{}

        mockClient.On("UpdateNodegroupVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNodegroupVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePodIdentityAssociation", func(t *testing.T) {
        input := &eks.UpdatePodIdentityAssociationInput{}
        output := &eks.UpdatePodIdentityAssociationOutput{}

        mockClient.On("UpdatePodIdentityAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePodIdentityAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
