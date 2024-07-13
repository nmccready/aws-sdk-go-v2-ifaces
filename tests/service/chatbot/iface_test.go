package chatbot_test

// tests for the chatbot service interface for this ../../../service/chatbot/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chatbot"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chatbot/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chatbot/chatbot_iface"
	"github.com/stretchr/testify/assert"
)

func TestChatbotServiceCanBeMocked(t *testing.T) {
	var iface chatbot_iface.IClient
	iface = &chatbot.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chatbot.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChimeWebhookConfiguration", func(t *testing.T) {
        input := &chatbot.CreateChimeWebhookConfigurationInput{}
        output := &chatbot.CreateChimeWebhookConfigurationOutput{}

        mockClient.On("CreateChimeWebhookConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChimeWebhookConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMicrosoftTeamsChannelConfiguration", func(t *testing.T) {
        input := &chatbot.CreateMicrosoftTeamsChannelConfigurationInput{}
        output := &chatbot.CreateMicrosoftTeamsChannelConfigurationOutput{}

        mockClient.On("CreateMicrosoftTeamsChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMicrosoftTeamsChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSlackChannelConfiguration", func(t *testing.T) {
        input := &chatbot.CreateSlackChannelConfigurationInput{}
        output := &chatbot.CreateSlackChannelConfigurationOutput{}

        mockClient.On("CreateSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChimeWebhookConfiguration", func(t *testing.T) {
        input := &chatbot.DeleteChimeWebhookConfigurationInput{}
        output := &chatbot.DeleteChimeWebhookConfigurationOutput{}

        mockClient.On("DeleteChimeWebhookConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChimeWebhookConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMicrosoftTeamsChannelConfiguration", func(t *testing.T) {
        input := &chatbot.DeleteMicrosoftTeamsChannelConfigurationInput{}
        output := &chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput{}

        mockClient.On("DeleteMicrosoftTeamsChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMicrosoftTeamsChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMicrosoftTeamsConfiguredTeam", func(t *testing.T) {
        input := &chatbot.DeleteMicrosoftTeamsConfiguredTeamInput{}
        output := &chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput{}

        mockClient.On("DeleteMicrosoftTeamsConfiguredTeam", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMicrosoftTeamsConfiguredTeam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMicrosoftTeamsUserIdentity", func(t *testing.T) {
        input := &chatbot.DeleteMicrosoftTeamsUserIdentityInput{}
        output := &chatbot.DeleteMicrosoftTeamsUserIdentityOutput{}

        mockClient.On("DeleteMicrosoftTeamsUserIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMicrosoftTeamsUserIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlackChannelConfiguration", func(t *testing.T) {
        input := &chatbot.DeleteSlackChannelConfigurationInput{}
        output := &chatbot.DeleteSlackChannelConfigurationOutput{}

        mockClient.On("DeleteSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlackUserIdentity", func(t *testing.T) {
        input := &chatbot.DeleteSlackUserIdentityInput{}
        output := &chatbot.DeleteSlackUserIdentityOutput{}

        mockClient.On("DeleteSlackUserIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlackUserIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlackWorkspaceAuthorization", func(t *testing.T) {
        input := &chatbot.DeleteSlackWorkspaceAuthorizationInput{}
        output := &chatbot.DeleteSlackWorkspaceAuthorizationOutput{}

        mockClient.On("DeleteSlackWorkspaceAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlackWorkspaceAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChimeWebhookConfigurations", func(t *testing.T) {
        input := &chatbot.DescribeChimeWebhookConfigurationsInput{}
        output := &chatbot.DescribeChimeWebhookConfigurationsOutput{}

        mockClient.On("DescribeChimeWebhookConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChimeWebhookConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSlackChannelConfigurations", func(t *testing.T) {
        input := &chatbot.DescribeSlackChannelConfigurationsInput{}
        output := &chatbot.DescribeSlackChannelConfigurationsOutput{}

        mockClient.On("DescribeSlackChannelConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSlackChannelConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSlackUserIdentities", func(t *testing.T) {
        input := &chatbot.DescribeSlackUserIdentitiesInput{}
        output := &chatbot.DescribeSlackUserIdentitiesOutput{}

        mockClient.On("DescribeSlackUserIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSlackUserIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSlackWorkspaces", func(t *testing.T) {
        input := &chatbot.DescribeSlackWorkspacesInput{}
        output := &chatbot.DescribeSlackWorkspacesOutput{}

        mockClient.On("DescribeSlackWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSlackWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountPreferences", func(t *testing.T) {
        input := &chatbot.GetAccountPreferencesInput{}
        output := &chatbot.GetAccountPreferencesOutput{}

        mockClient.On("GetAccountPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMicrosoftTeamsChannelConfiguration", func(t *testing.T) {
        input := &chatbot.GetMicrosoftTeamsChannelConfigurationInput{}
        output := &chatbot.GetMicrosoftTeamsChannelConfigurationOutput{}

        mockClient.On("GetMicrosoftTeamsChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetMicrosoftTeamsChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMicrosoftTeamsChannelConfigurations", func(t *testing.T) {
        input := &chatbot.ListMicrosoftTeamsChannelConfigurationsInput{}
        output := &chatbot.ListMicrosoftTeamsChannelConfigurationsOutput{}

        mockClient.On("ListMicrosoftTeamsChannelConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListMicrosoftTeamsChannelConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMicrosoftTeamsConfiguredTeams", func(t *testing.T) {
        input := &chatbot.ListMicrosoftTeamsConfiguredTeamsInput{}
        output := &chatbot.ListMicrosoftTeamsConfiguredTeamsOutput{}

        mockClient.On("ListMicrosoftTeamsConfiguredTeams", ctx, input).Return(output, nil)

        result, err := mockClient.ListMicrosoftTeamsConfiguredTeams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMicrosoftTeamsUserIdentities", func(t *testing.T) {
        input := &chatbot.ListMicrosoftTeamsUserIdentitiesInput{}
        output := &chatbot.ListMicrosoftTeamsUserIdentitiesOutput{}

        mockClient.On("ListMicrosoftTeamsUserIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListMicrosoftTeamsUserIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chatbot.ListTagsForResourceInput{}
        output := &chatbot.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chatbot.TagResourceInput{}
        output := &chatbot.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chatbot.UntagResourceInput{}
        output := &chatbot.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountPreferences", func(t *testing.T) {
        input := &chatbot.UpdateAccountPreferencesInput{}
        output := &chatbot.UpdateAccountPreferencesOutput{}

        mockClient.On("UpdateAccountPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChimeWebhookConfiguration", func(t *testing.T) {
        input := &chatbot.UpdateChimeWebhookConfigurationInput{}
        output := &chatbot.UpdateChimeWebhookConfigurationOutput{}

        mockClient.On("UpdateChimeWebhookConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChimeWebhookConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMicrosoftTeamsChannelConfiguration", func(t *testing.T) {
        input := &chatbot.UpdateMicrosoftTeamsChannelConfigurationInput{}
        output := &chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput{}

        mockClient.On("UpdateMicrosoftTeamsChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMicrosoftTeamsChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSlackChannelConfiguration", func(t *testing.T) {
        input := &chatbot.UpdateSlackChannelConfigurationInput{}
        output := &chatbot.UpdateSlackChannelConfigurationOutput{}

        mockClient.On("UpdateSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
