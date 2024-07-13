package cloudwatchevents_test

// tests for the cloudwatchevents service interface for this ../../../service/cloudwatchevents/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatchevents/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatchevents/cloudwatchevents_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudwatcheventsServiceCanBeMocked(t *testing.T) {
	var iface cloudwatchevents_iface.IClient
	iface = &cloudwatchevents.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudwatchevents.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateEventSource", func(t *testing.T) {
        input := &cloudwatchevents.ActivateEventSourceInput{}
        output := &cloudwatchevents.ActivateEventSourceOutput{}

        mockClient.On("ActivateEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelReplay", func(t *testing.T) {
        input := &cloudwatchevents.CancelReplayInput{}
        output := &cloudwatchevents.CancelReplayOutput{}

        mockClient.On("CancelReplay", ctx, input).Return(output, nil)

        result, err := mockClient.CancelReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiDestination", func(t *testing.T) {
        input := &cloudwatchevents.CreateApiDestinationInput{}
        output := &cloudwatchevents.CreateApiDestinationOutput{}

        mockClient.On("CreateApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateArchive", func(t *testing.T) {
        input := &cloudwatchevents.CreateArchiveInput{}
        output := &cloudwatchevents.CreateArchiveOutput{}

        mockClient.On("CreateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.CreateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &cloudwatchevents.CreateConnectionInput{}
        output := &cloudwatchevents.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventBus", func(t *testing.T) {
        input := &cloudwatchevents.CreateEventBusInput{}
        output := &cloudwatchevents.CreateEventBusOutput{}

        mockClient.On("CreateEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartnerEventSource", func(t *testing.T) {
        input := &cloudwatchevents.CreatePartnerEventSourceInput{}
        output := &cloudwatchevents.CreatePartnerEventSourceOutput{}

        mockClient.On("CreatePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateEventSource", func(t *testing.T) {
        input := &cloudwatchevents.DeactivateEventSourceInput{}
        output := &cloudwatchevents.DeactivateEventSourceOutput{}

        mockClient.On("DeactivateEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeauthorizeConnection", func(t *testing.T) {
        input := &cloudwatchevents.DeauthorizeConnectionInput{}
        output := &cloudwatchevents.DeauthorizeConnectionOutput{}

        mockClient.On("DeauthorizeConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeauthorizeConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiDestination", func(t *testing.T) {
        input := &cloudwatchevents.DeleteApiDestinationInput{}
        output := &cloudwatchevents.DeleteApiDestinationOutput{}

        mockClient.On("DeleteApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArchive", func(t *testing.T) {
        input := &cloudwatchevents.DeleteArchiveInput{}
        output := &cloudwatchevents.DeleteArchiveOutput{}

        mockClient.On("DeleteArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &cloudwatchevents.DeleteConnectionInput{}
        output := &cloudwatchevents.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventBus", func(t *testing.T) {
        input := &cloudwatchevents.DeleteEventBusInput{}
        output := &cloudwatchevents.DeleteEventBusOutput{}

        mockClient.On("DeleteEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartnerEventSource", func(t *testing.T) {
        input := &cloudwatchevents.DeletePartnerEventSourceInput{}
        output := &cloudwatchevents.DeletePartnerEventSourceOutput{}

        mockClient.On("DeletePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &cloudwatchevents.DeleteRuleInput{}
        output := &cloudwatchevents.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApiDestination", func(t *testing.T) {
        input := &cloudwatchevents.DescribeApiDestinationInput{}
        output := &cloudwatchevents.DescribeApiDestinationOutput{}

        mockClient.On("DescribeApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeArchive", func(t *testing.T) {
        input := &cloudwatchevents.DescribeArchiveInput{}
        output := &cloudwatchevents.DescribeArchiveOutput{}

        mockClient.On("DescribeArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnection", func(t *testing.T) {
        input := &cloudwatchevents.DescribeConnectionInput{}
        output := &cloudwatchevents.DescribeConnectionOutput{}

        mockClient.On("DescribeConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventBus", func(t *testing.T) {
        input := &cloudwatchevents.DescribeEventBusInput{}
        output := &cloudwatchevents.DescribeEventBusOutput{}

        mockClient.On("DescribeEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSource", func(t *testing.T) {
        input := &cloudwatchevents.DescribeEventSourceInput{}
        output := &cloudwatchevents.DescribeEventSourceOutput{}

        mockClient.On("DescribeEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePartnerEventSource", func(t *testing.T) {
        input := &cloudwatchevents.DescribePartnerEventSourceInput{}
        output := &cloudwatchevents.DescribePartnerEventSourceOutput{}

        mockClient.On("DescribePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplay", func(t *testing.T) {
        input := &cloudwatchevents.DescribeReplayInput{}
        output := &cloudwatchevents.DescribeReplayOutput{}

        mockClient.On("DescribeReplay", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRule", func(t *testing.T) {
        input := &cloudwatchevents.DescribeRuleInput{}
        output := &cloudwatchevents.DescribeRuleOutput{}

        mockClient.On("DescribeRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableRule", func(t *testing.T) {
        input := &cloudwatchevents.DisableRuleInput{}
        output := &cloudwatchevents.DisableRuleOutput{}

        mockClient.On("DisableRule", ctx, input).Return(output, nil)

        result, err := mockClient.DisableRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableRule", func(t *testing.T) {
        input := &cloudwatchevents.EnableRuleInput{}
        output := &cloudwatchevents.EnableRuleOutput{}

        mockClient.On("EnableRule", ctx, input).Return(output, nil)

        result, err := mockClient.EnableRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApiDestinations", func(t *testing.T) {
        input := &cloudwatchevents.ListApiDestinationsInput{}
        output := &cloudwatchevents.ListApiDestinationsOutput{}

        mockClient.On("ListApiDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListApiDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchives", func(t *testing.T) {
        input := &cloudwatchevents.ListArchivesInput{}
        output := &cloudwatchevents.ListArchivesOutput{}

        mockClient.On("ListArchives", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnections", func(t *testing.T) {
        input := &cloudwatchevents.ListConnectionsInput{}
        output := &cloudwatchevents.ListConnectionsOutput{}

        mockClient.On("ListConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventBuses", func(t *testing.T) {
        input := &cloudwatchevents.ListEventBusesInput{}
        output := &cloudwatchevents.ListEventBusesOutput{}

        mockClient.On("ListEventBuses", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventBuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventSources", func(t *testing.T) {
        input := &cloudwatchevents.ListEventSourcesInput{}
        output := &cloudwatchevents.ListEventSourcesOutput{}

        mockClient.On("ListEventSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerEventSourceAccounts", func(t *testing.T) {
        input := &cloudwatchevents.ListPartnerEventSourceAccountsInput{}
        output := &cloudwatchevents.ListPartnerEventSourceAccountsOutput{}

        mockClient.On("ListPartnerEventSourceAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerEventSourceAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerEventSources", func(t *testing.T) {
        input := &cloudwatchevents.ListPartnerEventSourcesInput{}
        output := &cloudwatchevents.ListPartnerEventSourcesOutput{}

        mockClient.On("ListPartnerEventSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerEventSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReplays", func(t *testing.T) {
        input := &cloudwatchevents.ListReplaysInput{}
        output := &cloudwatchevents.ListReplaysOutput{}

        mockClient.On("ListReplays", ctx, input).Return(output, nil)

        result, err := mockClient.ListReplays(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleNamesByTarget", func(t *testing.T) {
        input := &cloudwatchevents.ListRuleNamesByTargetInput{}
        output := &cloudwatchevents.ListRuleNamesByTargetOutput{}

        mockClient.On("ListRuleNamesByTarget", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleNamesByTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &cloudwatchevents.ListRulesInput{}
        output := &cloudwatchevents.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloudwatchevents.ListTagsForResourceInput{}
        output := &cloudwatchevents.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetsByRule", func(t *testing.T) {
        input := &cloudwatchevents.ListTargetsByRuleInput{}
        output := &cloudwatchevents.ListTargetsByRuleOutput{}

        mockClient.On("ListTargetsByRule", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetsByRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEvents", func(t *testing.T) {
        input := &cloudwatchevents.PutEventsInput{}
        output := &cloudwatchevents.PutEventsOutput{}

        mockClient.On("PutEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPartnerEvents", func(t *testing.T) {
        input := &cloudwatchevents.PutPartnerEventsInput{}
        output := &cloudwatchevents.PutPartnerEventsOutput{}

        mockClient.On("PutPartnerEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutPartnerEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermission", func(t *testing.T) {
        input := &cloudwatchevents.PutPermissionInput{}
        output := &cloudwatchevents.PutPermissionOutput{}

        mockClient.On("PutPermission", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRule", func(t *testing.T) {
        input := &cloudwatchevents.PutRuleInput{}
        output := &cloudwatchevents.PutRuleOutput{}

        mockClient.On("PutRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTargets", func(t *testing.T) {
        input := &cloudwatchevents.PutTargetsInput{}
        output := &cloudwatchevents.PutTargetsOutput{}

        mockClient.On("PutTargets", ctx, input).Return(output, nil)

        result, err := mockClient.PutTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &cloudwatchevents.RemovePermissionInput{}
        output := &cloudwatchevents.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTargets", func(t *testing.T) {
        input := &cloudwatchevents.RemoveTargetsInput{}
        output := &cloudwatchevents.RemoveTargetsOutput{}

        mockClient.On("RemoveTargets", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplay", func(t *testing.T) {
        input := &cloudwatchevents.StartReplayInput{}
        output := &cloudwatchevents.StartReplayOutput{}

        mockClient.On("StartReplay", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloudwatchevents.TagResourceInput{}
        output := &cloudwatchevents.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestEventPattern", func(t *testing.T) {
        input := &cloudwatchevents.TestEventPatternInput{}
        output := &cloudwatchevents.TestEventPatternOutput{}

        mockClient.On("TestEventPattern", ctx, input).Return(output, nil)

        result, err := mockClient.TestEventPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloudwatchevents.UntagResourceInput{}
        output := &cloudwatchevents.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiDestination", func(t *testing.T) {
        input := &cloudwatchevents.UpdateApiDestinationInput{}
        output := &cloudwatchevents.UpdateApiDestinationOutput{}

        mockClient.On("UpdateApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateArchive", func(t *testing.T) {
        input := &cloudwatchevents.UpdateArchiveInput{}
        output := &cloudwatchevents.UpdateArchiveOutput{}

        mockClient.On("UpdateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnection", func(t *testing.T) {
        input := &cloudwatchevents.UpdateConnectionInput{}
        output := &cloudwatchevents.UpdateConnectionOutput{}

        mockClient.On("UpdateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
