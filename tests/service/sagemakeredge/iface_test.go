package sagemakeredge_test

// tests for the sagemakeredge service interface for this ../../../service/sagemakeredge/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakeredge"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakeredge/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakeredge/sagemakeredge_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakeredgeServiceCanBeMocked(t *testing.T) {
	var iface sagemakeredge_iface.IClient
	iface = &sagemakeredge.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakeredge.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployments", func(t *testing.T) {
        input := &sagemakeredge.GetDeploymentsInput{}
        output := &sagemakeredge.GetDeploymentsOutput{}

        mockClient.On("GetDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceRegistration", func(t *testing.T) {
        input := &sagemakeredge.GetDeviceRegistrationInput{}
        output := &sagemakeredge.GetDeviceRegistrationOutput{}

        mockClient.On("GetDeviceRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendHeartbeat", func(t *testing.T) {
        input := &sagemakeredge.SendHeartbeatInput{}
        output := &sagemakeredge.SendHeartbeatOutput{}

        mockClient.On("SendHeartbeat", ctx, input).Return(output, nil)

        result, err := mockClient.SendHeartbeat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
