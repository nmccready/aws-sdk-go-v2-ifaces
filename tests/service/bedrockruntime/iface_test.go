package bedrockruntime_test

// tests for the bedrockruntime service interface for this ../../../service/bedrockruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockruntime/bedrockruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestBedrockruntimeServiceCanBeMocked(t *testing.T) {
	var iface bedrockruntime_iface.IClient
	iface = &bedrockruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bedrockruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyGuardrail", func(t *testing.T) {
        input := &bedrockruntime.ApplyGuardrailInput{}
        output := &bedrockruntime.ApplyGuardrailOutput{}

        mockClient.On("ApplyGuardrail", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyGuardrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConverse", func(t *testing.T) {
        input := &bedrockruntime.ConverseInput{}
        output := &bedrockruntime.ConverseOutput{}

        mockClient.On("Converse", ctx, input).Return(output, nil)

        result, err := mockClient.Converse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConverseStream", func(t *testing.T) {
        input := &bedrockruntime.ConverseStreamInput{}
        output := &bedrockruntime.ConverseStreamOutput{}

        mockClient.On("ConverseStream", ctx, input).Return(output, nil)

        result, err := mockClient.ConverseStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeModel", func(t *testing.T) {
        input := &bedrockruntime.InvokeModelInput{}
        output := &bedrockruntime.InvokeModelOutput{}

        mockClient.On("InvokeModel", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeModelWithResponseStream", func(t *testing.T) {
        input := &bedrockruntime.InvokeModelWithResponseStreamInput{}
        output := &bedrockruntime.InvokeModelWithResponseStreamOutput{}

        mockClient.On("InvokeModelWithResponseStream", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeModelWithResponseStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
