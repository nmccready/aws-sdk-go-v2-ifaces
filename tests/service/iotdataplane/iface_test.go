package iotdataplane_test

// tests for the iotdataplane service interface for this ../../../service/iotdataplane/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotdataplane/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotdataplane/iotdataplane_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotdataplaneServiceCanBeMocked(t *testing.T) {
	var iface iotdataplane_iface.IClient
	iface = &iotdataplane.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotdataplane.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThingShadow", func(t *testing.T) {
        input := &iotdataplane.DeleteThingShadowInput{}
        output := &iotdataplane.DeleteThingShadowOutput{}

        mockClient.On("DeleteThingShadow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThingShadow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRetainedMessage", func(t *testing.T) {
        input := &iotdataplane.GetRetainedMessageInput{}
        output := &iotdataplane.GetRetainedMessageOutput{}

        mockClient.On("GetRetainedMessage", ctx, input).Return(output, nil)

        result, err := mockClient.GetRetainedMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetThingShadow", func(t *testing.T) {
        input := &iotdataplane.GetThingShadowInput{}
        output := &iotdataplane.GetThingShadowOutput{}

        mockClient.On("GetThingShadow", ctx, input).Return(output, nil)

        result, err := mockClient.GetThingShadow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNamedShadowsForThing", func(t *testing.T) {
        input := &iotdataplane.ListNamedShadowsForThingInput{}
        output := &iotdataplane.ListNamedShadowsForThingOutput{}

        mockClient.On("ListNamedShadowsForThing", ctx, input).Return(output, nil)

        result, err := mockClient.ListNamedShadowsForThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRetainedMessages", func(t *testing.T) {
        input := &iotdataplane.ListRetainedMessagesInput{}
        output := &iotdataplane.ListRetainedMessagesOutput{}

        mockClient.On("ListRetainedMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListRetainedMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublish", func(t *testing.T) {
        input := &iotdataplane.PublishInput{}
        output := &iotdataplane.PublishOutput{}

        mockClient.On("Publish", ctx, input).Return(output, nil)

        result, err := mockClient.Publish(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThingShadow", func(t *testing.T) {
        input := &iotdataplane.UpdateThingShadowInput{}
        output := &iotdataplane.UpdateThingShadowOutput{}

        mockClient.On("UpdateThingShadow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThingShadow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
