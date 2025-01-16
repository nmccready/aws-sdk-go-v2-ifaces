// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package networkflowmonitor_test

// tests for the networkflowmonitor service interface for 
// this ../../../service/networkflowmonitor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkflowmonitor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkflowmonitor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkflowmonitor/networkflowmonitor_iface"
	"github.com/stretchr/testify/assert"
)

func TestNetworkflowmonitorServiceCanBeMocked(t *testing.T) {
	var iface networkflowmonitor_iface.IClient
	iface = &networkflowmonitor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := networkflowmonitor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitor", func(t *testing.T) {
        input := &networkflowmonitor.CreateMonitorInput{}
        output := &networkflowmonitor.CreateMonitorOutput{}

        mockClient.On("CreateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScope", func(t *testing.T) {
        input := &networkflowmonitor.CreateScopeInput{}
        output := &networkflowmonitor.CreateScopeOutput{}

        mockClient.On("CreateScope", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitor", func(t *testing.T) {
        input := &networkflowmonitor.DeleteMonitorInput{}
        output := &networkflowmonitor.DeleteMonitorOutput{}

        mockClient.On("DeleteMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScope", func(t *testing.T) {
        input := &networkflowmonitor.DeleteScopeInput{}
        output := &networkflowmonitor.DeleteScopeOutput{}

        mockClient.On("DeleteScope", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMonitor", func(t *testing.T) {
        input := &networkflowmonitor.GetMonitorInput{}
        output := &networkflowmonitor.GetMonitorOutput{}

        mockClient.On("GetMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.GetMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResultsMonitorTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryResultsMonitorTopContributorsInput{}
        output := &networkflowmonitor.GetQueryResultsMonitorTopContributorsOutput{}

        mockClient.On("GetQueryResultsMonitorTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResultsMonitorTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResultsWorkloadInsightsTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsInput{}
        output := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsOutput{}

        mockClient.On("GetQueryResultsWorkloadInsightsTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResultsWorkloadInsightsTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResultsWorkloadInsightsTopContributorsData", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsDataInput{}
        output := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsDataOutput{}

        mockClient.On("GetQueryResultsWorkloadInsightsTopContributorsData", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResultsWorkloadInsightsTopContributorsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryStatusMonitorTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryStatusMonitorTopContributorsInput{}
        output := &networkflowmonitor.GetQueryStatusMonitorTopContributorsOutput{}

        mockClient.On("GetQueryStatusMonitorTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryStatusMonitorTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryStatusWorkloadInsightsTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsInput{}
        output := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsOutput{}

        mockClient.On("GetQueryStatusWorkloadInsightsTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryStatusWorkloadInsightsTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryStatusWorkloadInsightsTopContributorsData", func(t *testing.T) {
        input := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsDataInput{}
        output := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsDataOutput{}

        mockClient.On("GetQueryStatusWorkloadInsightsTopContributorsData", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryStatusWorkloadInsightsTopContributorsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScope", func(t *testing.T) {
        input := &networkflowmonitor.GetScopeInput{}
        output := &networkflowmonitor.GetScopeOutput{}

        mockClient.On("GetScope", ctx, input).Return(output, nil)

        result, err := mockClient.GetScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitors", func(t *testing.T) {
        input := &networkflowmonitor.ListMonitorsInput{}
        output := &networkflowmonitor.ListMonitorsOutput{}

        mockClient.On("ListMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScopes", func(t *testing.T) {
        input := &networkflowmonitor.ListScopesInput{}
        output := &networkflowmonitor.ListScopesOutput{}

        mockClient.On("ListScopes", ctx, input).Return(output, nil)

        result, err := mockClient.ListScopes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &networkflowmonitor.ListTagsForResourceInput{}
        output := &networkflowmonitor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQueryMonitorTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.StartQueryMonitorTopContributorsInput{}
        output := &networkflowmonitor.StartQueryMonitorTopContributorsOutput{}

        mockClient.On("StartQueryMonitorTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.StartQueryMonitorTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQueryWorkloadInsightsTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsInput{}
        output := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsOutput{}

        mockClient.On("StartQueryWorkloadInsightsTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.StartQueryWorkloadInsightsTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQueryWorkloadInsightsTopContributorsData", func(t *testing.T) {
        input := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsDataInput{}
        output := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsDataOutput{}

        mockClient.On("StartQueryWorkloadInsightsTopContributorsData", ctx, input).Return(output, nil)

        result, err := mockClient.StartQueryWorkloadInsightsTopContributorsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQueryMonitorTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.StopQueryMonitorTopContributorsInput{}
        output := &networkflowmonitor.StopQueryMonitorTopContributorsOutput{}

        mockClient.On("StopQueryMonitorTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.StopQueryMonitorTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQueryWorkloadInsightsTopContributors", func(t *testing.T) {
        input := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsInput{}
        output := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsOutput{}

        mockClient.On("StopQueryWorkloadInsightsTopContributors", ctx, input).Return(output, nil)

        result, err := mockClient.StopQueryWorkloadInsightsTopContributors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQueryWorkloadInsightsTopContributorsData", func(t *testing.T) {
        input := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsDataInput{}
        output := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsDataOutput{}

        mockClient.On("StopQueryWorkloadInsightsTopContributorsData", ctx, input).Return(output, nil)

        result, err := mockClient.StopQueryWorkloadInsightsTopContributorsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &networkflowmonitor.TagResourceInput{}
        output := &networkflowmonitor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &networkflowmonitor.UntagResourceInput{}
        output := &networkflowmonitor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitor", func(t *testing.T) {
        input := &networkflowmonitor.UpdateMonitorInput{}
        output := &networkflowmonitor.UpdateMonitorOutput{}

        mockClient.On("UpdateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScope", func(t *testing.T) {
        input := &networkflowmonitor.UpdateScopeInput{}
        output := &networkflowmonitor.UpdateScopeOutput{}

        mockClient.On("UpdateScope", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
