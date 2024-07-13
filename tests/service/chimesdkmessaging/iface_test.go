package chimesdkmessaging_test

// tests for the chimesdkmessaging service interface for this ../../../service/chimesdkmessaging/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmessaging/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmessaging/chimesdkmessaging_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimesdkmessagingServiceCanBeMocked(t *testing.T) {
	var iface chimesdkmessaging_iface.IClient
	iface = &chimesdkmessaging.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chimesdkmessaging.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.AssociateChannelFlowInput{}
        output := &chimesdkmessaging.AssociateChannelFlowOutput{}

        mockClient.On("AssociateChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateChannelMembership", func(t *testing.T) {
        input := &chimesdkmessaging.BatchCreateChannelMembershipInput{}
        output := &chimesdkmessaging.BatchCreateChannelMembershipOutput{}

        mockClient.On("BatchCreateChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChannelFlowCallback", func(t *testing.T) {
        input := &chimesdkmessaging.ChannelFlowCallbackInput{}
        output := &chimesdkmessaging.ChannelFlowCallbackOutput{}

        mockClient.On("ChannelFlowCallback", ctx, input).Return(output, nil)

        result, err := mockClient.ChannelFlowCallback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &chimesdkmessaging.CreateChannelInput{}
        output := &chimesdkmessaging.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelBan", func(t *testing.T) {
        input := &chimesdkmessaging.CreateChannelBanInput{}
        output := &chimesdkmessaging.CreateChannelBanOutput{}

        mockClient.On("CreateChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.CreateChannelFlowInput{}
        output := &chimesdkmessaging.CreateChannelFlowOutput{}

        mockClient.On("CreateChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelMembership", func(t *testing.T) {
        input := &chimesdkmessaging.CreateChannelMembershipInput{}
        output := &chimesdkmessaging.CreateChannelMembershipOutput{}

        mockClient.On("CreateChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelModerator", func(t *testing.T) {
        input := &chimesdkmessaging.CreateChannelModeratorInput{}
        output := &chimesdkmessaging.CreateChannelModeratorOutput{}

        mockClient.On("CreateChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelInput{}
        output := &chimesdkmessaging.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelBan", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelBanInput{}
        output := &chimesdkmessaging.DeleteChannelBanOutput{}

        mockClient.On("DeleteChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelFlowInput{}
        output := &chimesdkmessaging.DeleteChannelFlowOutput{}

        mockClient.On("DeleteChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelMembership", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelMembershipInput{}
        output := &chimesdkmessaging.DeleteChannelMembershipOutput{}

        mockClient.On("DeleteChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelMessage", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelMessageInput{}
        output := &chimesdkmessaging.DeleteChannelMessageOutput{}

        mockClient.On("DeleteChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelModerator", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteChannelModeratorInput{}
        output := &chimesdkmessaging.DeleteChannelModeratorOutput{}

        mockClient.On("DeleteChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMessagingStreamingConfigurations", func(t *testing.T) {
        input := &chimesdkmessaging.DeleteMessagingStreamingConfigurationsInput{}
        output := &chimesdkmessaging.DeleteMessagingStreamingConfigurationsOutput{}

        mockClient.On("DeleteMessagingStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMessagingStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelInput{}
        output := &chimesdkmessaging.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelBan", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelBanInput{}
        output := &chimesdkmessaging.DescribeChannelBanOutput{}

        mockClient.On("DescribeChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelFlowInput{}
        output := &chimesdkmessaging.DescribeChannelFlowOutput{}

        mockClient.On("DescribeChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelMembership", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelMembershipInput{}
        output := &chimesdkmessaging.DescribeChannelMembershipOutput{}

        mockClient.On("DescribeChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelMembershipForAppInstanceUser", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput{}
        output := &chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput{}

        mockClient.On("DescribeChannelMembershipForAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelMembershipForAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelModeratedByAppInstanceUser", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput{}
        output := &chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput{}

        mockClient.On("DescribeChannelModeratedByAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelModeratedByAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelModerator", func(t *testing.T) {
        input := &chimesdkmessaging.DescribeChannelModeratorInput{}
        output := &chimesdkmessaging.DescribeChannelModeratorOutput{}

        mockClient.On("DescribeChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.DisassociateChannelFlowInput{}
        output := &chimesdkmessaging.DisassociateChannelFlowOutput{}

        mockClient.On("DisassociateChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelMembershipPreferences", func(t *testing.T) {
        input := &chimesdkmessaging.GetChannelMembershipPreferencesInput{}
        output := &chimesdkmessaging.GetChannelMembershipPreferencesOutput{}

        mockClient.On("GetChannelMembershipPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelMembershipPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelMessage", func(t *testing.T) {
        input := &chimesdkmessaging.GetChannelMessageInput{}
        output := &chimesdkmessaging.GetChannelMessageOutput{}

        mockClient.On("GetChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelMessageStatus", func(t *testing.T) {
        input := &chimesdkmessaging.GetChannelMessageStatusInput{}
        output := &chimesdkmessaging.GetChannelMessageStatusOutput{}

        mockClient.On("GetChannelMessageStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelMessageStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMessagingSessionEndpoint", func(t *testing.T) {
        input := &chimesdkmessaging.GetMessagingSessionEndpointInput{}
        output := &chimesdkmessaging.GetMessagingSessionEndpointOutput{}

        mockClient.On("GetMessagingSessionEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetMessagingSessionEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMessagingStreamingConfigurations", func(t *testing.T) {
        input := &chimesdkmessaging.GetMessagingStreamingConfigurationsInput{}
        output := &chimesdkmessaging.GetMessagingStreamingConfigurationsOutput{}

        mockClient.On("GetMessagingStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.GetMessagingStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelBans", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelBansInput{}
        output := &chimesdkmessaging.ListChannelBansOutput{}

        mockClient.On("ListChannelBans", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelBans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelFlows", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelFlowsInput{}
        output := &chimesdkmessaging.ListChannelFlowsOutput{}

        mockClient.On("ListChannelFlows", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMemberships", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelMembershipsInput{}
        output := &chimesdkmessaging.ListChannelMembershipsOutput{}

        mockClient.On("ListChannelMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMembershipsForAppInstanceUser", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput{}
        output := &chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput{}

        mockClient.On("ListChannelMembershipsForAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMembershipsForAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMessages", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelMessagesInput{}
        output := &chimesdkmessaging.ListChannelMessagesOutput{}

        mockClient.On("ListChannelMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelModerators", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelModeratorsInput{}
        output := &chimesdkmessaging.ListChannelModeratorsOutput{}

        mockClient.On("ListChannelModerators", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelModerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelsInput{}
        output := &chimesdkmessaging.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelsAssociatedWithChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput{}
        output := &chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput{}

        mockClient.On("ListChannelsAssociatedWithChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelsAssociatedWithChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelsModeratedByAppInstanceUser", func(t *testing.T) {
        input := &chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput{}
        output := &chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput{}

        mockClient.On("ListChannelsModeratedByAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelsModeratedByAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubChannels", func(t *testing.T) {
        input := &chimesdkmessaging.ListSubChannelsInput{}
        output := &chimesdkmessaging.ListSubChannelsOutput{}

        mockClient.On("ListSubChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chimesdkmessaging.ListTagsForResourceInput{}
        output := &chimesdkmessaging.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutChannelExpirationSettings", func(t *testing.T) {
        input := &chimesdkmessaging.PutChannelExpirationSettingsInput{}
        output := &chimesdkmessaging.PutChannelExpirationSettingsOutput{}

        mockClient.On("PutChannelExpirationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutChannelExpirationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutChannelMembershipPreferences", func(t *testing.T) {
        input := &chimesdkmessaging.PutChannelMembershipPreferencesInput{}
        output := &chimesdkmessaging.PutChannelMembershipPreferencesOutput{}

        mockClient.On("PutChannelMembershipPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.PutChannelMembershipPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMessagingStreamingConfigurations", func(t *testing.T) {
        input := &chimesdkmessaging.PutMessagingStreamingConfigurationsInput{}
        output := &chimesdkmessaging.PutMessagingStreamingConfigurationsOutput{}

        mockClient.On("PutMessagingStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.PutMessagingStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRedactChannelMessage", func(t *testing.T) {
        input := &chimesdkmessaging.RedactChannelMessageInput{}
        output := &chimesdkmessaging.RedactChannelMessageOutput{}

        mockClient.On("RedactChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.RedactChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchChannels", func(t *testing.T) {
        input := &chimesdkmessaging.SearchChannelsInput{}
        output := &chimesdkmessaging.SearchChannelsOutput{}

        mockClient.On("SearchChannels", ctx, input).Return(output, nil)

        result, err := mockClient.SearchChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendChannelMessage", func(t *testing.T) {
        input := &chimesdkmessaging.SendChannelMessageInput{}
        output := &chimesdkmessaging.SendChannelMessageOutput{}

        mockClient.On("SendChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chimesdkmessaging.TagResourceInput{}
        output := &chimesdkmessaging.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chimesdkmessaging.UntagResourceInput{}
        output := &chimesdkmessaging.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &chimesdkmessaging.UpdateChannelInput{}
        output := &chimesdkmessaging.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelFlow", func(t *testing.T) {
        input := &chimesdkmessaging.UpdateChannelFlowInput{}
        output := &chimesdkmessaging.UpdateChannelFlowOutput{}

        mockClient.On("UpdateChannelFlow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelMessage", func(t *testing.T) {
        input := &chimesdkmessaging.UpdateChannelMessageInput{}
        output := &chimesdkmessaging.UpdateChannelMessageOutput{}

        mockClient.On("UpdateChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelReadMarker", func(t *testing.T) {
        input := &chimesdkmessaging.UpdateChannelReadMarkerInput{}
        output := &chimesdkmessaging.UpdateChannelReadMarkerOutput{}

        mockClient.On("UpdateChannelReadMarker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelReadMarker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
