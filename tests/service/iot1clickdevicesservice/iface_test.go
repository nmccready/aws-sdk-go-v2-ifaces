package iot1clickdevicesservice_test

// tests for the iot1clickdevicesservice service interface for this ../../../service/iot1clickdevicesservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot1clickdevicesservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot1clickdevicesservice/iot1clickdevicesservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestIot1clickdevicesserviceServiceCanBeMocked(t *testing.T) {
	var iface iot1clickdevicesservice_iface.IClient
	iface = &iot1clickdevicesservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iot1clickdevicesservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClaimDevicesByClaimCode", func(t *testing.T) {
        input := &iot1clickdevicesservice.ClaimDevicesByClaimCodeInput{}
        output := &iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput{}

        mockClient.On("ClaimDevicesByClaimCode", ctx, input).Return(output, nil)

        result, err := mockClient.ClaimDevicesByClaimCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevice", func(t *testing.T) {
        input := &iot1clickdevicesservice.DescribeDeviceInput{}
        output := &iot1clickdevicesservice.DescribeDeviceOutput{}

        mockClient.On("DescribeDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFinalizeDeviceClaim", func(t *testing.T) {
        input := &iot1clickdevicesservice.FinalizeDeviceClaimInput{}
        output := &iot1clickdevicesservice.FinalizeDeviceClaimOutput{}

        mockClient.On("FinalizeDeviceClaim", ctx, input).Return(output, nil)

        result, err := mockClient.FinalizeDeviceClaim(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceMethods", func(t *testing.T) {
        input := &iot1clickdevicesservice.GetDeviceMethodsInput{}
        output := &iot1clickdevicesservice.GetDeviceMethodsOutput{}

        mockClient.On("GetDeviceMethods", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceMethods(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateDeviceClaim", func(t *testing.T) {
        input := &iot1clickdevicesservice.InitiateDeviceClaimInput{}
        output := &iot1clickdevicesservice.InitiateDeviceClaimOutput{}

        mockClient.On("InitiateDeviceClaim", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateDeviceClaim(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeDeviceMethod", func(t *testing.T) {
        input := &iot1clickdevicesservice.InvokeDeviceMethodInput{}
        output := &iot1clickdevicesservice.InvokeDeviceMethodOutput{}

        mockClient.On("InvokeDeviceMethod", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeDeviceMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceEvents", func(t *testing.T) {
        input := &iot1clickdevicesservice.ListDeviceEventsInput{}
        output := &iot1clickdevicesservice.ListDeviceEventsOutput{}

        mockClient.On("ListDeviceEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &iot1clickdevicesservice.ListDevicesInput{}
        output := &iot1clickdevicesservice.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iot1clickdevicesservice.ListTagsForResourceInput{}
        output := &iot1clickdevicesservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iot1clickdevicesservice.TagResourceInput{}
        output := &iot1clickdevicesservice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnclaimDevice", func(t *testing.T) {
        input := &iot1clickdevicesservice.UnclaimDeviceInput{}
        output := &iot1clickdevicesservice.UnclaimDeviceOutput{}

        mockClient.On("UnclaimDevice", ctx, input).Return(output, nil)

        result, err := mockClient.UnclaimDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iot1clickdevicesservice.UntagResourceInput{}
        output := &iot1clickdevicesservice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceState", func(t *testing.T) {
        input := &iot1clickdevicesservice.UpdateDeviceStateInput{}
        output := &iot1clickdevicesservice.UpdateDeviceStateOutput{}

        mockClient.On("UpdateDeviceState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
