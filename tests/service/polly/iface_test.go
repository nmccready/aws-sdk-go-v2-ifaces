package polly_test

// tests for the polly service interface for this ../../../service/polly/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/polly/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/polly/polly_iface"
	"github.com/stretchr/testify/assert"
)

func TestPollyServiceCanBeMocked(t *testing.T) {
	var iface polly_iface.IClient
	iface = &polly.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := polly.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLexicon", func(t *testing.T) {
        input := &polly.DeleteLexiconInput{}
        output := &polly.DeleteLexiconOutput{}

        mockClient.On("DeleteLexicon", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLexicon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVoices", func(t *testing.T) {
        input := &polly.DescribeVoicesInput{}
        output := &polly.DescribeVoicesOutput{}

        mockClient.On("DescribeVoices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVoices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLexicon", func(t *testing.T) {
        input := &polly.GetLexiconInput{}
        output := &polly.GetLexiconOutput{}

        mockClient.On("GetLexicon", ctx, input).Return(output, nil)

        result, err := mockClient.GetLexicon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpeechSynthesisTask", func(t *testing.T) {
        input := &polly.GetSpeechSynthesisTaskInput{}
        output := &polly.GetSpeechSynthesisTaskOutput{}

        mockClient.On("GetSpeechSynthesisTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpeechSynthesisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLexicons", func(t *testing.T) {
        input := &polly.ListLexiconsInput{}
        output := &polly.ListLexiconsOutput{}

        mockClient.On("ListLexicons", ctx, input).Return(output, nil)

        result, err := mockClient.ListLexicons(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpeechSynthesisTasks", func(t *testing.T) {
        input := &polly.ListSpeechSynthesisTasksInput{}
        output := &polly.ListSpeechSynthesisTasksOutput{}

        mockClient.On("ListSpeechSynthesisTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpeechSynthesisTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLexicon", func(t *testing.T) {
        input := &polly.PutLexiconInput{}
        output := &polly.PutLexiconOutput{}

        mockClient.On("PutLexicon", ctx, input).Return(output, nil)

        result, err := mockClient.PutLexicon(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSpeechSynthesisTask", func(t *testing.T) {
        input := &polly.StartSpeechSynthesisTaskInput{}
        output := &polly.StartSpeechSynthesisTaskOutput{}

        mockClient.On("StartSpeechSynthesisTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartSpeechSynthesisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSynthesizeSpeech", func(t *testing.T) {
        input := &polly.SynthesizeSpeechInput{}
        output := &polly.SynthesizeSpeechOutput{}

        mockClient.On("SynthesizeSpeech", ctx, input).Return(output, nil)

        result, err := mockClient.SynthesizeSpeech(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
