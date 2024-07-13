package connectparticipant_test

// tests for the connectparticipant service interface for this ../../../service/connectparticipant/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connectparticipant"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectparticipant/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectparticipant/connectparticipant_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectparticipantServiceCanBeMocked(t *testing.T) {
	var iface connectparticipant_iface.IClient
	iface = &connectparticipant.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connectparticipant.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteAttachmentUpload", func(t *testing.T) {
        input := &connectparticipant.CompleteAttachmentUploadInput{}
        output := &connectparticipant.CompleteAttachmentUploadOutput{}

        mockClient.On("CompleteAttachmentUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteAttachmentUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParticipantConnection", func(t *testing.T) {
        input := &connectparticipant.CreateParticipantConnectionInput{}
        output := &connectparticipant.CreateParticipantConnectionOutput{}

        mockClient.On("CreateParticipantConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParticipantConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeView", func(t *testing.T) {
        input := &connectparticipant.DescribeViewInput{}
        output := &connectparticipant.DescribeViewOutput{}

        mockClient.On("DescribeView", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectParticipant", func(t *testing.T) {
        input := &connectparticipant.DisconnectParticipantInput{}
        output := &connectparticipant.DisconnectParticipantOutput{}

        mockClient.On("DisconnectParticipant", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectParticipant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttachment", func(t *testing.T) {
        input := &connectparticipant.GetAttachmentInput{}
        output := &connectparticipant.GetAttachmentOutput{}

        mockClient.On("GetAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTranscript", func(t *testing.T) {
        input := &connectparticipant.GetTranscriptInput{}
        output := &connectparticipant.GetTranscriptOutput{}

        mockClient.On("GetTranscript", ctx, input).Return(output, nil)

        result, err := mockClient.GetTranscript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEvent", func(t *testing.T) {
        input := &connectparticipant.SendEventInput{}
        output := &connectparticipant.SendEventOutput{}

        mockClient.On("SendEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SendEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendMessage", func(t *testing.T) {
        input := &connectparticipant.SendMessageInput{}
        output := &connectparticipant.SendMessageOutput{}

        mockClient.On("SendMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAttachmentUpload", func(t *testing.T) {
        input := &connectparticipant.StartAttachmentUploadInput{}
        output := &connectparticipant.StartAttachmentUploadOutput{}

        mockClient.On("StartAttachmentUpload", ctx, input).Return(output, nil)

        result, err := mockClient.StartAttachmentUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
