package apigatewaymanagementapi_test

// tests for the apigatewaymanagementapi service interface for this ../../../service/apigatewaymanagementapi/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigatewaymanagementapi/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigatewaymanagementapi/apigatewaymanagementapi_iface"
	"github.com/stretchr/testify/assert"
)

func TestApigatewaymanagementapiServiceCanBeMocked(t *testing.T) {
	var iface apigatewaymanagementapi_iface.IClient
	iface = &apigatewaymanagementapi.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := apigatewaymanagementapi.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &apigatewaymanagementapi.DeleteConnectionInput{}
        output := &apigatewaymanagementapi.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnection", func(t *testing.T) {
        input := &apigatewaymanagementapi.GetConnectionInput{}
        output := &apigatewaymanagementapi.GetConnectionOutput{}

        mockClient.On("GetConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostToConnection", func(t *testing.T) {
        input := &apigatewaymanagementapi.PostToConnectionInput{}
        output := &apigatewaymanagementapi.PostToConnectionOutput{}

        mockClient.On("PostToConnection", ctx, input).Return(output, nil)

        result, err := mockClient.PostToConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
