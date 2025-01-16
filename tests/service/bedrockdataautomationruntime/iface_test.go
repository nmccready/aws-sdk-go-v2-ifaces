// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package bedrockdataautomationruntime_test

// tests for the bedrockdataautomationruntime service interface for 
// this ../../../service/bedrockdataautomationruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockdataautomationruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockdataautomationruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockdataautomationruntime/bedrockdataautomationruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestBedrockdataautomationruntimeServiceCanBeMocked(t *testing.T) {
	var iface bedrockdataautomationruntime_iface.IClient
	iface = &bedrockdataautomationruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bedrockdataautomationruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataAutomationStatus", func(t *testing.T) {
        input := &bedrockdataautomationruntime.GetDataAutomationStatusInput{}
        output := &bedrockdataautomationruntime.GetDataAutomationStatusOutput{}

        mockClient.On("GetDataAutomationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataAutomationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeDataAutomationAsync", func(t *testing.T) {
        input := &bedrockdataautomationruntime.InvokeDataAutomationAsyncInput{}
        output := &bedrockdataautomationruntime.InvokeDataAutomationAsyncOutput{}

        mockClient.On("InvokeDataAutomationAsync", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeDataAutomationAsync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
