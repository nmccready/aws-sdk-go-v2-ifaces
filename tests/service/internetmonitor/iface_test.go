package internetmonitor_test

// tests for the internetmonitor service interface for this ../../../service/internetmonitor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/internetmonitor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/internetmonitor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/internetmonitor/internetmonitor_iface"
	"github.com/stretchr/testify/assert"
)

func TestInternetmonitorServiceCanBeMocked(t *testing.T) {
	var iface internetmonitor_iface.IClient
	iface = &internetmonitor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := internetmonitor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitor", func(t *testing.T) {
        input := &internetmonitor.CreateMonitorInput{}
        output := &internetmonitor.CreateMonitorOutput{}

        mockClient.On("CreateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitor", func(t *testing.T) {
        input := &internetmonitor.DeleteMonitorInput{}
        output := &internetmonitor.DeleteMonitorOutput{}

        mockClient.On("DeleteMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHealthEvent", func(t *testing.T) {
        input := &internetmonitor.GetHealthEventInput{}
        output := &internetmonitor.GetHealthEventOutput{}

        mockClient.On("GetHealthEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetHealthEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInternetEvent", func(t *testing.T) {
        input := &internetmonitor.GetInternetEventInput{}
        output := &internetmonitor.GetInternetEventOutput{}

        mockClient.On("GetInternetEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetInternetEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMonitor", func(t *testing.T) {
        input := &internetmonitor.GetMonitorInput{}
        output := &internetmonitor.GetMonitorOutput{}

        mockClient.On("GetMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.GetMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResults", func(t *testing.T) {
        input := &internetmonitor.GetQueryResultsInput{}
        output := &internetmonitor.GetQueryResultsOutput{}

        mockClient.On("GetQueryResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryStatus", func(t *testing.T) {
        input := &internetmonitor.GetQueryStatusInput{}
        output := &internetmonitor.GetQueryStatusOutput{}

        mockClient.On("GetQueryStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHealthEvents", func(t *testing.T) {
        input := &internetmonitor.ListHealthEventsInput{}
        output := &internetmonitor.ListHealthEventsOutput{}

        mockClient.On("ListHealthEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListHealthEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInternetEvents", func(t *testing.T) {
        input := &internetmonitor.ListInternetEventsInput{}
        output := &internetmonitor.ListInternetEventsOutput{}

        mockClient.On("ListInternetEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListInternetEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitors", func(t *testing.T) {
        input := &internetmonitor.ListMonitorsInput{}
        output := &internetmonitor.ListMonitorsOutput{}

        mockClient.On("ListMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &internetmonitor.ListTagsForResourceInput{}
        output := &internetmonitor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQuery", func(t *testing.T) {
        input := &internetmonitor.StartQueryInput{}
        output := &internetmonitor.StartQueryOutput{}

        mockClient.On("StartQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StartQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQuery", func(t *testing.T) {
        input := &internetmonitor.StopQueryInput{}
        output := &internetmonitor.StopQueryOutput{}

        mockClient.On("StopQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StopQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &internetmonitor.TagResourceInput{}
        output := &internetmonitor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &internetmonitor.UntagResourceInput{}
        output := &internetmonitor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitor", func(t *testing.T) {
        input := &internetmonitor.UpdateMonitorInput{}
        output := &internetmonitor.UpdateMonitorOutput{}

        mockClient.On("UpdateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
