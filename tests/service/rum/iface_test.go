package rum_test

// tests for the rum service interface for this ../../../service/rum/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rum"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rum/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rum/rum_iface"
	"github.com/stretchr/testify/assert"
)

func TestRumServiceCanBeMocked(t *testing.T) {
	var iface rum_iface.IClient
	iface = &rum.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rum.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateRumMetricDefinitions", func(t *testing.T) {
        input := &rum.BatchCreateRumMetricDefinitionsInput{}
        output := &rum.BatchCreateRumMetricDefinitionsOutput{}

        mockClient.On("BatchCreateRumMetricDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateRumMetricDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteRumMetricDefinitions", func(t *testing.T) {
        input := &rum.BatchDeleteRumMetricDefinitionsInput{}
        output := &rum.BatchDeleteRumMetricDefinitionsOutput{}

        mockClient.On("BatchDeleteRumMetricDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteRumMetricDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetRumMetricDefinitions", func(t *testing.T) {
        input := &rum.BatchGetRumMetricDefinitionsInput{}
        output := &rum.BatchGetRumMetricDefinitionsOutput{}

        mockClient.On("BatchGetRumMetricDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetRumMetricDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppMonitor", func(t *testing.T) {
        input := &rum.CreateAppMonitorInput{}
        output := &rum.CreateAppMonitorOutput{}

        mockClient.On("CreateAppMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppMonitor", func(t *testing.T) {
        input := &rum.DeleteAppMonitorInput{}
        output := &rum.DeleteAppMonitorOutput{}

        mockClient.On("DeleteAppMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRumMetricsDestination", func(t *testing.T) {
        input := &rum.DeleteRumMetricsDestinationInput{}
        output := &rum.DeleteRumMetricsDestinationOutput{}

        mockClient.On("DeleteRumMetricsDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRumMetricsDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppMonitor", func(t *testing.T) {
        input := &rum.GetAppMonitorInput{}
        output := &rum.GetAppMonitorOutput{}

        mockClient.On("GetAppMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppMonitorData", func(t *testing.T) {
        input := &rum.GetAppMonitorDataInput{}
        output := &rum.GetAppMonitorDataOutput{}

        mockClient.On("GetAppMonitorData", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppMonitorData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppMonitors", func(t *testing.T) {
        input := &rum.ListAppMonitorsInput{}
        output := &rum.ListAppMonitorsOutput{}

        mockClient.On("ListAppMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRumMetricsDestinations", func(t *testing.T) {
        input := &rum.ListRumMetricsDestinationsInput{}
        output := &rum.ListRumMetricsDestinationsOutput{}

        mockClient.On("ListRumMetricsDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRumMetricsDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &rum.ListTagsForResourceInput{}
        output := &rum.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRumEvents", func(t *testing.T) {
        input := &rum.PutRumEventsInput{}
        output := &rum.PutRumEventsOutput{}

        mockClient.On("PutRumEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutRumEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRumMetricsDestination", func(t *testing.T) {
        input := &rum.PutRumMetricsDestinationInput{}
        output := &rum.PutRumMetricsDestinationOutput{}

        mockClient.On("PutRumMetricsDestination", ctx, input).Return(output, nil)

        result, err := mockClient.PutRumMetricsDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &rum.TagResourceInput{}
        output := &rum.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &rum.UntagResourceInput{}
        output := &rum.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppMonitor", func(t *testing.T) {
        input := &rum.UpdateAppMonitorInput{}
        output := &rum.UpdateAppMonitorOutput{}

        mockClient.On("UpdateAppMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRumMetricDefinition", func(t *testing.T) {
        input := &rum.UpdateRumMetricDefinitionInput{}
        output := &rum.UpdateRumMetricDefinitionOutput{}

        mockClient.On("UpdateRumMetricDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRumMetricDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
