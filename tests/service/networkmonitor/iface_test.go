package networkmonitor_test

// tests for the networkmonitor service interface for this ../../../service/networkmonitor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkmonitor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkmonitor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkmonitor/networkmonitor_iface"
	"github.com/stretchr/testify/assert"
)

func TestNetworkmonitorServiceCanBeMocked(t *testing.T) {
	var iface networkmonitor_iface.IClient
	iface = &networkmonitor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := networkmonitor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitor", func(t *testing.T) {
        input := &networkmonitor.CreateMonitorInput{}
        output := &networkmonitor.CreateMonitorOutput{}

        mockClient.On("CreateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProbe", func(t *testing.T) {
        input := &networkmonitor.CreateProbeInput{}
        output := &networkmonitor.CreateProbeOutput{}

        mockClient.On("CreateProbe", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProbe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitor", func(t *testing.T) {
        input := &networkmonitor.DeleteMonitorInput{}
        output := &networkmonitor.DeleteMonitorOutput{}

        mockClient.On("DeleteMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProbe", func(t *testing.T) {
        input := &networkmonitor.DeleteProbeInput{}
        output := &networkmonitor.DeleteProbeOutput{}

        mockClient.On("DeleteProbe", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProbe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMonitor", func(t *testing.T) {
        input := &networkmonitor.GetMonitorInput{}
        output := &networkmonitor.GetMonitorOutput{}

        mockClient.On("GetMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.GetMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProbe", func(t *testing.T) {
        input := &networkmonitor.GetProbeInput{}
        output := &networkmonitor.GetProbeOutput{}

        mockClient.On("GetProbe", ctx, input).Return(output, nil)

        result, err := mockClient.GetProbe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitors", func(t *testing.T) {
        input := &networkmonitor.ListMonitorsInput{}
        output := &networkmonitor.ListMonitorsOutput{}

        mockClient.On("ListMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &networkmonitor.ListTagsForResourceInput{}
        output := &networkmonitor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &networkmonitor.TagResourceInput{}
        output := &networkmonitor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &networkmonitor.UntagResourceInput{}
        output := &networkmonitor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitor", func(t *testing.T) {
        input := &networkmonitor.UpdateMonitorInput{}
        output := &networkmonitor.UpdateMonitorOutput{}

        mockClient.On("UpdateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProbe", func(t *testing.T) {
        input := &networkmonitor.UpdateProbeInput{}
        output := &networkmonitor.UpdateProbeOutput{}

        mockClient.On("UpdateProbe", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProbe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
