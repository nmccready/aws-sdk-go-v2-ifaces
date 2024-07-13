package sagemakerruntime_test

// tests for the sagemakerruntime service interface for this ../../../service/sagemakerruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakerruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakerruntime/sagemakerruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakerruntimeServiceCanBeMocked(t *testing.T) {
	var iface sagemakerruntime_iface.IClient
	iface = &sagemakerruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakerruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeEndpoint", func(t *testing.T) {
        input := &sagemakerruntime.InvokeEndpointInput{}
        output := &sagemakerruntime.InvokeEndpointOutput{}

        mockClient.On("InvokeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeEndpointAsync", func(t *testing.T) {
        input := &sagemakerruntime.InvokeEndpointAsyncInput{}
        output := &sagemakerruntime.InvokeEndpointAsyncOutput{}

        mockClient.On("InvokeEndpointAsync", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeEndpointAsync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeEndpointWithResponseStream", func(t *testing.T) {
        input := &sagemakerruntime.InvokeEndpointWithResponseStreamInput{}
        output := &sagemakerruntime.InvokeEndpointWithResponseStreamOutput{}

        mockClient.On("InvokeEndpointWithResponseStream", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeEndpointWithResponseStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
