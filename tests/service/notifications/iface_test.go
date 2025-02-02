// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package notifications_test

// tests for the notifications service interface for 
// this ../../../service/notifications/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/notifications"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/notifications/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/notifications/notifications_iface"
	"github.com/stretchr/testify/assert"
)

func TestNotificationsServiceCanBeMocked(t *testing.T) {
	var iface notifications_iface.IClient
	iface = &notifications.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := notifications.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateChannel", func(t *testing.T) {
        input := &notifications.AssociateChannelInput{}
        output := &notifications.AssociateChannelOutput{}

        mockClient.On("AssociateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateManagedNotificationAccountContact", func(t *testing.T) {
        input := &notifications.AssociateManagedNotificationAccountContactInput{}
        output := &notifications.AssociateManagedNotificationAccountContactOutput{}

        mockClient.On("AssociateManagedNotificationAccountContact", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateManagedNotificationAccountContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateManagedNotificationAdditionalChannel", func(t *testing.T) {
        input := &notifications.AssociateManagedNotificationAdditionalChannelInput{}
        output := &notifications.AssociateManagedNotificationAdditionalChannelOutput{}

        mockClient.On("AssociateManagedNotificationAdditionalChannel", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateManagedNotificationAdditionalChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventRule", func(t *testing.T) {
        input := &notifications.CreateEventRuleInput{}
        output := &notifications.CreateEventRuleOutput{}

        mockClient.On("CreateEventRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotificationConfiguration", func(t *testing.T) {
        input := &notifications.CreateNotificationConfigurationInput{}
        output := &notifications.CreateNotificationConfigurationOutput{}

        mockClient.On("CreateNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventRule", func(t *testing.T) {
        input := &notifications.DeleteEventRuleInput{}
        output := &notifications.DeleteEventRuleOutput{}

        mockClient.On("DeleteEventRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotificationConfiguration", func(t *testing.T) {
        input := &notifications.DeleteNotificationConfigurationInput{}
        output := &notifications.DeleteNotificationConfigurationOutput{}

        mockClient.On("DeleteNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterNotificationHub", func(t *testing.T) {
        input := &notifications.DeregisterNotificationHubInput{}
        output := &notifications.DeregisterNotificationHubOutput{}

        mockClient.On("DeregisterNotificationHub", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterNotificationHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableNotificationsAccessForOrganization", func(t *testing.T) {
        input := &notifications.DisableNotificationsAccessForOrganizationInput{}
        output := &notifications.DisableNotificationsAccessForOrganizationOutput{}

        mockClient.On("DisableNotificationsAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DisableNotificationsAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateChannel", func(t *testing.T) {
        input := &notifications.DisassociateChannelInput{}
        output := &notifications.DisassociateChannelOutput{}

        mockClient.On("DisassociateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateManagedNotificationAccountContact", func(t *testing.T) {
        input := &notifications.DisassociateManagedNotificationAccountContactInput{}
        output := &notifications.DisassociateManagedNotificationAccountContactOutput{}

        mockClient.On("DisassociateManagedNotificationAccountContact", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateManagedNotificationAccountContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateManagedNotificationAdditionalChannel", func(t *testing.T) {
        input := &notifications.DisassociateManagedNotificationAdditionalChannelInput{}
        output := &notifications.DisassociateManagedNotificationAdditionalChannelOutput{}

        mockClient.On("DisassociateManagedNotificationAdditionalChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateManagedNotificationAdditionalChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableNotificationsAccessForOrganization", func(t *testing.T) {
        input := &notifications.EnableNotificationsAccessForOrganizationInput{}
        output := &notifications.EnableNotificationsAccessForOrganizationOutput{}

        mockClient.On("EnableNotificationsAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.EnableNotificationsAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventRule", func(t *testing.T) {
        input := &notifications.GetEventRuleInput{}
        output := &notifications.GetEventRuleOutput{}

        mockClient.On("GetEventRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedNotificationChildEvent", func(t *testing.T) {
        input := &notifications.GetManagedNotificationChildEventInput{}
        output := &notifications.GetManagedNotificationChildEventOutput{}

        mockClient.On("GetManagedNotificationChildEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedNotificationChildEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedNotificationConfiguration", func(t *testing.T) {
        input := &notifications.GetManagedNotificationConfigurationInput{}
        output := &notifications.GetManagedNotificationConfigurationOutput{}

        mockClient.On("GetManagedNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedNotificationEvent", func(t *testing.T) {
        input := &notifications.GetManagedNotificationEventInput{}
        output := &notifications.GetManagedNotificationEventOutput{}

        mockClient.On("GetManagedNotificationEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedNotificationEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotificationConfiguration", func(t *testing.T) {
        input := &notifications.GetNotificationConfigurationInput{}
        output := &notifications.GetNotificationConfigurationOutput{}

        mockClient.On("GetNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotificationEvent", func(t *testing.T) {
        input := &notifications.GetNotificationEventInput{}
        output := &notifications.GetNotificationEventOutput{}

        mockClient.On("GetNotificationEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotificationEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotificationsAccessForOrganization", func(t *testing.T) {
        input := &notifications.GetNotificationsAccessForOrganizationInput{}
        output := &notifications.GetNotificationsAccessForOrganizationOutput{}

        mockClient.On("GetNotificationsAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotificationsAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &notifications.ListChannelsInput{}
        output := &notifications.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventRules", func(t *testing.T) {
        input := &notifications.ListEventRulesInput{}
        output := &notifications.ListEventRulesOutput{}

        mockClient.On("ListEventRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedNotificationChannelAssociations", func(t *testing.T) {
        input := &notifications.ListManagedNotificationChannelAssociationsInput{}
        output := &notifications.ListManagedNotificationChannelAssociationsOutput{}

        mockClient.On("ListManagedNotificationChannelAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedNotificationChannelAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedNotificationChildEvents", func(t *testing.T) {
        input := &notifications.ListManagedNotificationChildEventsInput{}
        output := &notifications.ListManagedNotificationChildEventsOutput{}

        mockClient.On("ListManagedNotificationChildEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedNotificationChildEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedNotificationConfigurations", func(t *testing.T) {
        input := &notifications.ListManagedNotificationConfigurationsInput{}
        output := &notifications.ListManagedNotificationConfigurationsOutput{}

        mockClient.On("ListManagedNotificationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedNotificationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedNotificationEvents", func(t *testing.T) {
        input := &notifications.ListManagedNotificationEventsInput{}
        output := &notifications.ListManagedNotificationEventsOutput{}

        mockClient.On("ListManagedNotificationEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedNotificationEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotificationConfigurations", func(t *testing.T) {
        input := &notifications.ListNotificationConfigurationsInput{}
        output := &notifications.ListNotificationConfigurationsOutput{}

        mockClient.On("ListNotificationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotificationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotificationEvents", func(t *testing.T) {
        input := &notifications.ListNotificationEventsInput{}
        output := &notifications.ListNotificationEventsOutput{}

        mockClient.On("ListNotificationEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotificationEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotificationHubs", func(t *testing.T) {
        input := &notifications.ListNotificationHubsInput{}
        output := &notifications.ListNotificationHubsOutput{}

        mockClient.On("ListNotificationHubs", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotificationHubs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &notifications.ListTagsForResourceInput{}
        output := &notifications.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterNotificationHub", func(t *testing.T) {
        input := &notifications.RegisterNotificationHubInput{}
        output := &notifications.RegisterNotificationHubOutput{}

        mockClient.On("RegisterNotificationHub", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterNotificationHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &notifications.TagResourceInput{}
        output := &notifications.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &notifications.UntagResourceInput{}
        output := &notifications.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventRule", func(t *testing.T) {
        input := &notifications.UpdateEventRuleInput{}
        output := &notifications.UpdateEventRuleOutput{}

        mockClient.On("UpdateEventRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotificationConfiguration", func(t *testing.T) {
        input := &notifications.UpdateNotificationConfigurationInput{}
        output := &notifications.UpdateNotificationConfigurationOutput{}

        mockClient.On("UpdateNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
