package bedrockagentruntime_test

// tests for the bedrockagentruntime service interface for this ../../../service/bedrockagentruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockagentruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockagentruntime/bedrockagentruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestBedrockagentruntimeServiceCanBeMocked(t *testing.T) {
	var iface bedrockagentruntime_iface.IClient
	iface = &bedrockagentruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bedrockagentruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgentMemory", func(t *testing.T) {
        input := &bedrockagentruntime.DeleteAgentMemoryInput{}
        output := &bedrockagentruntime.DeleteAgentMemoryOutput{}

        mockClient.On("DeleteAgentMemory", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgentMemory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentMemory", func(t *testing.T) {
        input := &bedrockagentruntime.GetAgentMemoryInput{}
        output := &bedrockagentruntime.GetAgentMemoryOutput{}

        mockClient.On("GetAgentMemory", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentMemory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeAgent", func(t *testing.T) {
        input := &bedrockagentruntime.InvokeAgentInput{}
        output := &bedrockagentruntime.InvokeAgentOutput{}

        mockClient.On("InvokeAgent", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeFlow", func(t *testing.T) {
        input := &bedrockagentruntime.InvokeFlowInput{}
        output := &bedrockagentruntime.InvokeFlowOutput{}

        mockClient.On("InvokeFlow", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieve", func(t *testing.T) {
        input := &bedrockagentruntime.RetrieveInput{}
        output := &bedrockagentruntime.RetrieveOutput{}

        mockClient.On("Retrieve", ctx, input).Return(output, nil)

        result, err := mockClient.Retrieve(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieveAndGenerate", func(t *testing.T) {
        input := &bedrockagentruntime.RetrieveAndGenerateInput{}
        output := &bedrockagentruntime.RetrieveAndGenerateOutput{}

        mockClient.On("RetrieveAndGenerate", ctx, input).Return(output, nil)

        result, err := mockClient.RetrieveAndGenerate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
