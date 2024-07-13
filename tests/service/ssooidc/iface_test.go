package ssooidc_test

// tests for the ssooidc service interface for this ../../../service/ssooidc/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssooidc/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssooidc/ssooidc_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsooidcServiceCanBeMocked(t *testing.T) {
	var iface ssooidc_iface.IClient
	iface = &ssooidc.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssooidc.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateToken", func(t *testing.T) {
        input := &ssooidc.CreateTokenInput{}
        output := &ssooidc.CreateTokenOutput{}

        mockClient.On("CreateToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTokenWithIAM", func(t *testing.T) {
        input := &ssooidc.CreateTokenWithIAMInput{}
        output := &ssooidc.CreateTokenWithIAMOutput{}

        mockClient.On("CreateTokenWithIAM", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTokenWithIAM(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterClient", func(t *testing.T) {
        input := &ssooidc.RegisterClientInput{}
        output := &ssooidc.RegisterClientOutput{}

        mockClient.On("RegisterClient", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeviceAuthorization", func(t *testing.T) {
        input := &ssooidc.StartDeviceAuthorizationInput{}
        output := &ssooidc.StartDeviceAuthorizationOutput{}

        mockClient.On("StartDeviceAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeviceAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
