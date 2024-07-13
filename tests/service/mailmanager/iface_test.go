package mailmanager_test

// tests for the mailmanager service interface for this ../../../service/mailmanager/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mailmanager"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mailmanager/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mailmanager/mailmanager_iface"
	"github.com/stretchr/testify/assert"
)

func TestMailmanagerServiceCanBeMocked(t *testing.T) {
	var iface mailmanager_iface.IClient
	iface = &mailmanager.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mailmanager.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAddonInstance", func(t *testing.T) {
        input := &mailmanager.CreateAddonInstanceInput{}
        output := &mailmanager.CreateAddonInstanceOutput{}

        mockClient.On("CreateAddonInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAddonInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAddonSubscription", func(t *testing.T) {
        input := &mailmanager.CreateAddonSubscriptionInput{}
        output := &mailmanager.CreateAddonSubscriptionOutput{}

        mockClient.On("CreateAddonSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAddonSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateArchive", func(t *testing.T) {
        input := &mailmanager.CreateArchiveInput{}
        output := &mailmanager.CreateArchiveOutput{}

        mockClient.On("CreateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.CreateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIngressPoint", func(t *testing.T) {
        input := &mailmanager.CreateIngressPointInput{}
        output := &mailmanager.CreateIngressPointOutput{}

        mockClient.On("CreateIngressPoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIngressPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRelay", func(t *testing.T) {
        input := &mailmanager.CreateRelayInput{}
        output := &mailmanager.CreateRelayOutput{}

        mockClient.On("CreateRelay", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRelay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleSet", func(t *testing.T) {
        input := &mailmanager.CreateRuleSetInput{}
        output := &mailmanager.CreateRuleSetOutput{}

        mockClient.On("CreateRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficPolicy", func(t *testing.T) {
        input := &mailmanager.CreateTrafficPolicyInput{}
        output := &mailmanager.CreateTrafficPolicyOutput{}

        mockClient.On("CreateTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAddonInstance", func(t *testing.T) {
        input := &mailmanager.DeleteAddonInstanceInput{}
        output := &mailmanager.DeleteAddonInstanceOutput{}

        mockClient.On("DeleteAddonInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAddonInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAddonSubscription", func(t *testing.T) {
        input := &mailmanager.DeleteAddonSubscriptionInput{}
        output := &mailmanager.DeleteAddonSubscriptionOutput{}

        mockClient.On("DeleteAddonSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAddonSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArchive", func(t *testing.T) {
        input := &mailmanager.DeleteArchiveInput{}
        output := &mailmanager.DeleteArchiveOutput{}

        mockClient.On("DeleteArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIngressPoint", func(t *testing.T) {
        input := &mailmanager.DeleteIngressPointInput{}
        output := &mailmanager.DeleteIngressPointOutput{}

        mockClient.On("DeleteIngressPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIngressPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRelay", func(t *testing.T) {
        input := &mailmanager.DeleteRelayInput{}
        output := &mailmanager.DeleteRelayOutput{}

        mockClient.On("DeleteRelay", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRelay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleSet", func(t *testing.T) {
        input := &mailmanager.DeleteRuleSetInput{}
        output := &mailmanager.DeleteRuleSetOutput{}

        mockClient.On("DeleteRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficPolicy", func(t *testing.T) {
        input := &mailmanager.DeleteTrafficPolicyInput{}
        output := &mailmanager.DeleteTrafficPolicyOutput{}

        mockClient.On("DeleteTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAddonInstance", func(t *testing.T) {
        input := &mailmanager.GetAddonInstanceInput{}
        output := &mailmanager.GetAddonInstanceOutput{}

        mockClient.On("GetAddonInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetAddonInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAddonSubscription", func(t *testing.T) {
        input := &mailmanager.GetAddonSubscriptionInput{}
        output := &mailmanager.GetAddonSubscriptionOutput{}

        mockClient.On("GetAddonSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetAddonSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchive", func(t *testing.T) {
        input := &mailmanager.GetArchiveInput{}
        output := &mailmanager.GetArchiveOutput{}

        mockClient.On("GetArchive", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveExport", func(t *testing.T) {
        input := &mailmanager.GetArchiveExportInput{}
        output := &mailmanager.GetArchiveExportOutput{}

        mockClient.On("GetArchiveExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveMessage", func(t *testing.T) {
        input := &mailmanager.GetArchiveMessageInput{}
        output := &mailmanager.GetArchiveMessageOutput{}

        mockClient.On("GetArchiveMessage", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveMessageContent", func(t *testing.T) {
        input := &mailmanager.GetArchiveMessageContentInput{}
        output := &mailmanager.GetArchiveMessageContentOutput{}

        mockClient.On("GetArchiveMessageContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveMessageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveSearch", func(t *testing.T) {
        input := &mailmanager.GetArchiveSearchInput{}
        output := &mailmanager.GetArchiveSearchOutput{}

        mockClient.On("GetArchiveSearch", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveSearch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveSearchResults", func(t *testing.T) {
        input := &mailmanager.GetArchiveSearchResultsInput{}
        output := &mailmanager.GetArchiveSearchResultsOutput{}

        mockClient.On("GetArchiveSearchResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveSearchResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIngressPoint", func(t *testing.T) {
        input := &mailmanager.GetIngressPointInput{}
        output := &mailmanager.GetIngressPointOutput{}

        mockClient.On("GetIngressPoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetIngressPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelay", func(t *testing.T) {
        input := &mailmanager.GetRelayInput{}
        output := &mailmanager.GetRelayOutput{}

        mockClient.On("GetRelay", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRuleSet", func(t *testing.T) {
        input := &mailmanager.GetRuleSetInput{}
        output := &mailmanager.GetRuleSetOutput{}

        mockClient.On("GetRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrafficPolicy", func(t *testing.T) {
        input := &mailmanager.GetTrafficPolicyInput{}
        output := &mailmanager.GetTrafficPolicyOutput{}

        mockClient.On("GetTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAddonInstances", func(t *testing.T) {
        input := &mailmanager.ListAddonInstancesInput{}
        output := &mailmanager.ListAddonInstancesOutput{}

        mockClient.On("ListAddonInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAddonInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAddonSubscriptions", func(t *testing.T) {
        input := &mailmanager.ListAddonSubscriptionsInput{}
        output := &mailmanager.ListAddonSubscriptionsOutput{}

        mockClient.On("ListAddonSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAddonSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchiveExports", func(t *testing.T) {
        input := &mailmanager.ListArchiveExportsInput{}
        output := &mailmanager.ListArchiveExportsOutput{}

        mockClient.On("ListArchiveExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchiveExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchiveSearches", func(t *testing.T) {
        input := &mailmanager.ListArchiveSearchesInput{}
        output := &mailmanager.ListArchiveSearchesOutput{}

        mockClient.On("ListArchiveSearches", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchiveSearches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchives", func(t *testing.T) {
        input := &mailmanager.ListArchivesInput{}
        output := &mailmanager.ListArchivesOutput{}

        mockClient.On("ListArchives", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIngressPoints", func(t *testing.T) {
        input := &mailmanager.ListIngressPointsInput{}
        output := &mailmanager.ListIngressPointsOutput{}

        mockClient.On("ListIngressPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListIngressPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRelays", func(t *testing.T) {
        input := &mailmanager.ListRelaysInput{}
        output := &mailmanager.ListRelaysOutput{}

        mockClient.On("ListRelays", ctx, input).Return(output, nil)

        result, err := mockClient.ListRelays(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleSets", func(t *testing.T) {
        input := &mailmanager.ListRuleSetsInput{}
        output := &mailmanager.ListRuleSetsOutput{}

        mockClient.On("ListRuleSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mailmanager.ListTagsForResourceInput{}
        output := &mailmanager.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicies", func(t *testing.T) {
        input := &mailmanager.ListTrafficPoliciesInput{}
        output := &mailmanager.ListTrafficPoliciesOutput{}

        mockClient.On("ListTrafficPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartArchiveExport", func(t *testing.T) {
        input := &mailmanager.StartArchiveExportInput{}
        output := &mailmanager.StartArchiveExportOutput{}

        mockClient.On("StartArchiveExport", ctx, input).Return(output, nil)

        result, err := mockClient.StartArchiveExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartArchiveSearch", func(t *testing.T) {
        input := &mailmanager.StartArchiveSearchInput{}
        output := &mailmanager.StartArchiveSearchOutput{}

        mockClient.On("StartArchiveSearch", ctx, input).Return(output, nil)

        result, err := mockClient.StartArchiveSearch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopArchiveExport", func(t *testing.T) {
        input := &mailmanager.StopArchiveExportInput{}
        output := &mailmanager.StopArchiveExportOutput{}

        mockClient.On("StopArchiveExport", ctx, input).Return(output, nil)

        result, err := mockClient.StopArchiveExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopArchiveSearch", func(t *testing.T) {
        input := &mailmanager.StopArchiveSearchInput{}
        output := &mailmanager.StopArchiveSearchOutput{}

        mockClient.On("StopArchiveSearch", ctx, input).Return(output, nil)

        result, err := mockClient.StopArchiveSearch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mailmanager.TagResourceInput{}
        output := &mailmanager.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mailmanager.UntagResourceInput{}
        output := &mailmanager.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateArchive", func(t *testing.T) {
        input := &mailmanager.UpdateArchiveInput{}
        output := &mailmanager.UpdateArchiveOutput{}

        mockClient.On("UpdateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIngressPoint", func(t *testing.T) {
        input := &mailmanager.UpdateIngressPointInput{}
        output := &mailmanager.UpdateIngressPointOutput{}

        mockClient.On("UpdateIngressPoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIngressPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRelay", func(t *testing.T) {
        input := &mailmanager.UpdateRelayInput{}
        output := &mailmanager.UpdateRelayOutput{}

        mockClient.On("UpdateRelay", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRelay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleSet", func(t *testing.T) {
        input := &mailmanager.UpdateRuleSetInput{}
        output := &mailmanager.UpdateRuleSetOutput{}

        mockClient.On("UpdateRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrafficPolicy", func(t *testing.T) {
        input := &mailmanager.UpdateTrafficPolicyInput{}
        output := &mailmanager.UpdateTrafficPolicyOutput{}

        mockClient.On("UpdateTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
