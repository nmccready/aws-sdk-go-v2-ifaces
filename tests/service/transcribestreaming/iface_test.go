package transcribestreaming_test

// tests for the transcribestreaming service interface for this ../../../service/transcribestreaming/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transcribestreaming/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transcribestreaming/transcribestreaming_iface"
	"github.com/stretchr/testify/assert"
)

func TestTranscribestreamingServiceCanBeMocked(t *testing.T) {
	var iface transcribestreaming_iface.IClient
	iface = &transcribestreaming.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := transcribestreaming.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCallAnalyticsStreamTranscription", func(t *testing.T) {
        input := &transcribestreaming.StartCallAnalyticsStreamTranscriptionInput{}
        output := &transcribestreaming.StartCallAnalyticsStreamTranscriptionOutput{}

        mockClient.On("StartCallAnalyticsStreamTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartCallAnalyticsStreamTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMedicalStreamTranscription", func(t *testing.T) {
        input := &transcribestreaming.StartMedicalStreamTranscriptionInput{}
        output := &transcribestreaming.StartMedicalStreamTranscriptionOutput{}

        mockClient.On("StartMedicalStreamTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartMedicalStreamTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStreamTranscription", func(t *testing.T) {
        input := &transcribestreaming.StartStreamTranscriptionInput{}
        output := &transcribestreaming.StartStreamTranscriptionOutput{}

        mockClient.On("StartStreamTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartStreamTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
