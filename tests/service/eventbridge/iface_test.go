package eventbridge_test

// tests for the eventbridge service interface for this ../../../service/eventbridge/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eventbridge/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eventbridge/eventbridge_iface"
	"github.com/stretchr/testify/assert"
)

func TestEventbridgeServiceCanBeMocked(t *testing.T) {
	var iface eventbridge_iface.IClient
	iface = &eventbridge.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := eventbridge.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateEventSource", func(t *testing.T) {
        input := &eventbridge.ActivateEventSourceInput{}
        output := &eventbridge.ActivateEventSourceOutput{}

        mockClient.On("ActivateEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelReplay", func(t *testing.T) {
        input := &eventbridge.CancelReplayInput{}
        output := &eventbridge.CancelReplayOutput{}

        mockClient.On("CancelReplay", ctx, input).Return(output, nil)

        result, err := mockClient.CancelReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiDestination", func(t *testing.T) {
        input := &eventbridge.CreateApiDestinationInput{}
        output := &eventbridge.CreateApiDestinationOutput{}

        mockClient.On("CreateApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateArchive", func(t *testing.T) {
        input := &eventbridge.CreateArchiveInput{}
        output := &eventbridge.CreateArchiveOutput{}

        mockClient.On("CreateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.CreateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &eventbridge.CreateConnectionInput{}
        output := &eventbridge.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpoint", func(t *testing.T) {
        input := &eventbridge.CreateEndpointInput{}
        output := &eventbridge.CreateEndpointOutput{}

        mockClient.On("CreateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventBus", func(t *testing.T) {
        input := &eventbridge.CreateEventBusInput{}
        output := &eventbridge.CreateEventBusOutput{}

        mockClient.On("CreateEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartnerEventSource", func(t *testing.T) {
        input := &eventbridge.CreatePartnerEventSourceInput{}
        output := &eventbridge.CreatePartnerEventSourceOutput{}

        mockClient.On("CreatePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateEventSource", func(t *testing.T) {
        input := &eventbridge.DeactivateEventSourceInput{}
        output := &eventbridge.DeactivateEventSourceOutput{}

        mockClient.On("DeactivateEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeauthorizeConnection", func(t *testing.T) {
        input := &eventbridge.DeauthorizeConnectionInput{}
        output := &eventbridge.DeauthorizeConnectionOutput{}

        mockClient.On("DeauthorizeConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeauthorizeConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiDestination", func(t *testing.T) {
        input := &eventbridge.DeleteApiDestinationInput{}
        output := &eventbridge.DeleteApiDestinationOutput{}

        mockClient.On("DeleteApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArchive", func(t *testing.T) {
        input := &eventbridge.DeleteArchiveInput{}
        output := &eventbridge.DeleteArchiveOutput{}

        mockClient.On("DeleteArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &eventbridge.DeleteConnectionInput{}
        output := &eventbridge.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &eventbridge.DeleteEndpointInput{}
        output := &eventbridge.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventBus", func(t *testing.T) {
        input := &eventbridge.DeleteEventBusInput{}
        output := &eventbridge.DeleteEventBusOutput{}

        mockClient.On("DeleteEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartnerEventSource", func(t *testing.T) {
        input := &eventbridge.DeletePartnerEventSourceInput{}
        output := &eventbridge.DeletePartnerEventSourceOutput{}

        mockClient.On("DeletePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &eventbridge.DeleteRuleInput{}
        output := &eventbridge.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApiDestination", func(t *testing.T) {
        input := &eventbridge.DescribeApiDestinationInput{}
        output := &eventbridge.DescribeApiDestinationOutput{}

        mockClient.On("DescribeApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeArchive", func(t *testing.T) {
        input := &eventbridge.DescribeArchiveInput{}
        output := &eventbridge.DescribeArchiveOutput{}

        mockClient.On("DescribeArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnection", func(t *testing.T) {
        input := &eventbridge.DescribeConnectionInput{}
        output := &eventbridge.DescribeConnectionOutput{}

        mockClient.On("DescribeConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoint", func(t *testing.T) {
        input := &eventbridge.DescribeEndpointInput{}
        output := &eventbridge.DescribeEndpointOutput{}

        mockClient.On("DescribeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventBus", func(t *testing.T) {
        input := &eventbridge.DescribeEventBusInput{}
        output := &eventbridge.DescribeEventBusOutput{}

        mockClient.On("DescribeEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSource", func(t *testing.T) {
        input := &eventbridge.DescribeEventSourceInput{}
        output := &eventbridge.DescribeEventSourceOutput{}

        mockClient.On("DescribeEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePartnerEventSource", func(t *testing.T) {
        input := &eventbridge.DescribePartnerEventSourceInput{}
        output := &eventbridge.DescribePartnerEventSourceOutput{}

        mockClient.On("DescribePartnerEventSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePartnerEventSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplay", func(t *testing.T) {
        input := &eventbridge.DescribeReplayInput{}
        output := &eventbridge.DescribeReplayOutput{}

        mockClient.On("DescribeReplay", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRule", func(t *testing.T) {
        input := &eventbridge.DescribeRuleInput{}
        output := &eventbridge.DescribeRuleOutput{}

        mockClient.On("DescribeRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableRule", func(t *testing.T) {
        input := &eventbridge.DisableRuleInput{}
        output := &eventbridge.DisableRuleOutput{}

        mockClient.On("DisableRule", ctx, input).Return(output, nil)

        result, err := mockClient.DisableRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableRule", func(t *testing.T) {
        input := &eventbridge.EnableRuleInput{}
        output := &eventbridge.EnableRuleOutput{}

        mockClient.On("EnableRule", ctx, input).Return(output, nil)

        result, err := mockClient.EnableRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApiDestinations", func(t *testing.T) {
        input := &eventbridge.ListApiDestinationsInput{}
        output := &eventbridge.ListApiDestinationsOutput{}

        mockClient.On("ListApiDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListApiDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchives", func(t *testing.T) {
        input := &eventbridge.ListArchivesInput{}
        output := &eventbridge.ListArchivesOutput{}

        mockClient.On("ListArchives", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnections", func(t *testing.T) {
        input := &eventbridge.ListConnectionsInput{}
        output := &eventbridge.ListConnectionsOutput{}

        mockClient.On("ListConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpoints", func(t *testing.T) {
        input := &eventbridge.ListEndpointsInput{}
        output := &eventbridge.ListEndpointsOutput{}

        mockClient.On("ListEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventBuses", func(t *testing.T) {
        input := &eventbridge.ListEventBusesInput{}
        output := &eventbridge.ListEventBusesOutput{}

        mockClient.On("ListEventBuses", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventBuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventSources", func(t *testing.T) {
        input := &eventbridge.ListEventSourcesInput{}
        output := &eventbridge.ListEventSourcesOutput{}

        mockClient.On("ListEventSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerEventSourceAccounts", func(t *testing.T) {
        input := &eventbridge.ListPartnerEventSourceAccountsInput{}
        output := &eventbridge.ListPartnerEventSourceAccountsOutput{}

        mockClient.On("ListPartnerEventSourceAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerEventSourceAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerEventSources", func(t *testing.T) {
        input := &eventbridge.ListPartnerEventSourcesInput{}
        output := &eventbridge.ListPartnerEventSourcesOutput{}

        mockClient.On("ListPartnerEventSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerEventSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReplays", func(t *testing.T) {
        input := &eventbridge.ListReplaysInput{}
        output := &eventbridge.ListReplaysOutput{}

        mockClient.On("ListReplays", ctx, input).Return(output, nil)

        result, err := mockClient.ListReplays(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleNamesByTarget", func(t *testing.T) {
        input := &eventbridge.ListRuleNamesByTargetInput{}
        output := &eventbridge.ListRuleNamesByTargetOutput{}

        mockClient.On("ListRuleNamesByTarget", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleNamesByTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &eventbridge.ListRulesInput{}
        output := &eventbridge.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &eventbridge.ListTagsForResourceInput{}
        output := &eventbridge.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetsByRule", func(t *testing.T) {
        input := &eventbridge.ListTargetsByRuleInput{}
        output := &eventbridge.ListTargetsByRuleOutput{}

        mockClient.On("ListTargetsByRule", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetsByRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEvents", func(t *testing.T) {
        input := &eventbridge.PutEventsInput{}
        output := &eventbridge.PutEventsOutput{}

        mockClient.On("PutEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPartnerEvents", func(t *testing.T) {
        input := &eventbridge.PutPartnerEventsInput{}
        output := &eventbridge.PutPartnerEventsOutput{}

        mockClient.On("PutPartnerEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutPartnerEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermission", func(t *testing.T) {
        input := &eventbridge.PutPermissionInput{}
        output := &eventbridge.PutPermissionOutput{}

        mockClient.On("PutPermission", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRule", func(t *testing.T) {
        input := &eventbridge.PutRuleInput{}
        output := &eventbridge.PutRuleOutput{}

        mockClient.On("PutRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTargets", func(t *testing.T) {
        input := &eventbridge.PutTargetsInput{}
        output := &eventbridge.PutTargetsOutput{}

        mockClient.On("PutTargets", ctx, input).Return(output, nil)

        result, err := mockClient.PutTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &eventbridge.RemovePermissionInput{}
        output := &eventbridge.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTargets", func(t *testing.T) {
        input := &eventbridge.RemoveTargetsInput{}
        output := &eventbridge.RemoveTargetsOutput{}

        mockClient.On("RemoveTargets", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplay", func(t *testing.T) {
        input := &eventbridge.StartReplayInput{}
        output := &eventbridge.StartReplayOutput{}

        mockClient.On("StartReplay", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &eventbridge.TagResourceInput{}
        output := &eventbridge.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestEventPattern", func(t *testing.T) {
        input := &eventbridge.TestEventPatternInput{}
        output := &eventbridge.TestEventPatternOutput{}

        mockClient.On("TestEventPattern", ctx, input).Return(output, nil)

        result, err := mockClient.TestEventPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &eventbridge.UntagResourceInput{}
        output := &eventbridge.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiDestination", func(t *testing.T) {
        input := &eventbridge.UpdateApiDestinationInput{}
        output := &eventbridge.UpdateApiDestinationOutput{}

        mockClient.On("UpdateApiDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateArchive", func(t *testing.T) {
        input := &eventbridge.UpdateArchiveInput{}
        output := &eventbridge.UpdateArchiveOutput{}

        mockClient.On("UpdateArchive", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnection", func(t *testing.T) {
        input := &eventbridge.UpdateConnectionInput{}
        output := &eventbridge.UpdateConnectionOutput{}

        mockClient.On("UpdateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpoint", func(t *testing.T) {
        input := &eventbridge.UpdateEndpointInput{}
        output := &eventbridge.UpdateEndpointOutput{}

        mockClient.On("UpdateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventBus", func(t *testing.T) {
        input := &eventbridge.UpdateEventBusInput{}
        output := &eventbridge.UpdateEventBusOutput{}

        mockClient.On("UpdateEventBus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventBus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
