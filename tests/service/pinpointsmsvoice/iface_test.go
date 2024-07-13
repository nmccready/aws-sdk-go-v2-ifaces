package pinpointsmsvoice_test

// tests for the pinpointsmsvoice service interface for this ../../../service/pinpointsmsvoice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointsmsvoice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointsmsvoice/pinpointsmsvoice_iface"
	"github.com/stretchr/testify/assert"
)

func TestPinpointsmsvoiceServiceCanBeMocked(t *testing.T) {
	var iface pinpointsmsvoice_iface.IClient
	iface = &pinpointsmsvoice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pinpointsmsvoice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSet", func(t *testing.T) {
        input := &pinpointsmsvoice.CreateConfigurationSetInput{}
        output := &pinpointsmsvoice.CreateConfigurationSetOutput{}

        mockClient.On("CreateConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoice.CreateConfigurationSetEventDestinationInput{}
        output := &pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput{}

        mockClient.On("CreateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSet", func(t *testing.T) {
        input := &pinpointsmsvoice.DeleteConfigurationSetInput{}
        output := &pinpointsmsvoice.DeleteConfigurationSetOutput{}

        mockClient.On("DeleteConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput{}
        output := &pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput{}

        mockClient.On("DeleteConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationSetEventDestinations", func(t *testing.T) {
        input := &pinpointsmsvoice.GetConfigurationSetEventDestinationsInput{}
        output := &pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput{}

        mockClient.On("GetConfigurationSetEventDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationSetEventDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationSets", func(t *testing.T) {
        input := &pinpointsmsvoice.ListConfigurationSetsInput{}
        output := &pinpointsmsvoice.ListConfigurationSetsOutput{}

        mockClient.On("ListConfigurationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendVoiceMessage", func(t *testing.T) {
        input := &pinpointsmsvoice.SendVoiceMessageInput{}
        output := &pinpointsmsvoice.SendVoiceMessageOutput{}

        mockClient.On("SendVoiceMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendVoiceMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput{}
        output := &pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput{}

        mockClient.On("UpdateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
