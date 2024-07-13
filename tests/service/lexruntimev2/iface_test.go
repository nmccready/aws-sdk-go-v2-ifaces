package lexruntimev2_test

// tests for the lexruntimev2 service interface for this ../../../service/lexruntimev2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexruntimev2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexruntimev2/lexruntimev2_iface"
	"github.com/stretchr/testify/assert"
)

func TestLexruntimev2ServiceCanBeMocked(t *testing.T) {
	var iface lexruntimev2_iface.IClient
	iface = &lexruntimev2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lexruntimev2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSession", func(t *testing.T) {
        input := &lexruntimev2.DeleteSessionInput{}
        output := &lexruntimev2.DeleteSessionOutput{}

        mockClient.On("DeleteSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &lexruntimev2.GetSessionInput{}
        output := &lexruntimev2.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSession", func(t *testing.T) {
        input := &lexruntimev2.PutSessionInput{}
        output := &lexruntimev2.PutSessionOutput{}

        mockClient.On("PutSession", ctx, input).Return(output, nil)

        result, err := mockClient.PutSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecognizeText", func(t *testing.T) {
        input := &lexruntimev2.RecognizeTextInput{}
        output := &lexruntimev2.RecognizeTextOutput{}

        mockClient.On("RecognizeText", ctx, input).Return(output, nil)

        result, err := mockClient.RecognizeText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecognizeUtterance", func(t *testing.T) {
        input := &lexruntimev2.RecognizeUtteranceInput{}
        output := &lexruntimev2.RecognizeUtteranceOutput{}

        mockClient.On("RecognizeUtterance", ctx, input).Return(output, nil)

        result, err := mockClient.RecognizeUtterance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConversation", func(t *testing.T) {
        input := &lexruntimev2.StartConversationInput{}
        output := &lexruntimev2.StartConversationOutput{}

        mockClient.On("StartConversation", ctx, input).Return(output, nil)

        result, err := mockClient.StartConversation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
