package sts_test

// tests for the sts service interface for this ../../../service/sts/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sts/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sts/sts_iface"
	"github.com/stretchr/testify/assert"
)

func TestStsServiceCanBeMocked(t *testing.T) {
	var iface sts_iface.IClient
	iface = &sts.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sts.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeRole", func(t *testing.T) {
        input := &sts.AssumeRoleInput{}
        output := &sts.AssumeRoleOutput{}

        mockClient.On("AssumeRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeRoleWithSAML", func(t *testing.T) {
        input := &sts.AssumeRoleWithSAMLInput{}
        output := &sts.AssumeRoleWithSAMLOutput{}

        mockClient.On("AssumeRoleWithSAML", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeRoleWithSAML(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeRoleWithWebIdentity", func(t *testing.T) {
        input := &sts.AssumeRoleWithWebIdentityInput{}
        output := &sts.AssumeRoleWithWebIdentityOutput{}

        mockClient.On("AssumeRoleWithWebIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeRoleWithWebIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecodeAuthorizationMessage", func(t *testing.T) {
        input := &sts.DecodeAuthorizationMessageInput{}
        output := &sts.DecodeAuthorizationMessageOutput{}

        mockClient.On("DecodeAuthorizationMessage", ctx, input).Return(output, nil)

        result, err := mockClient.DecodeAuthorizationMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessKeyInfo", func(t *testing.T) {
        input := &sts.GetAccessKeyInfoInput{}
        output := &sts.GetAccessKeyInfoOutput{}

        mockClient.On("GetAccessKeyInfo", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessKeyInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCallerIdentity", func(t *testing.T) {
        input := &sts.GetCallerIdentityInput{}
        output := &sts.GetCallerIdentityOutput{}

        mockClient.On("GetCallerIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetCallerIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFederationToken", func(t *testing.T) {
        input := &sts.GetFederationTokenInput{}
        output := &sts.GetFederationTokenOutput{}

        mockClient.On("GetFederationToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetFederationToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSessionToken", func(t *testing.T) {
        input := &sts.GetSessionTokenInput{}
        output := &sts.GetSessionTokenOutput{}

        mockClient.On("GetSessionToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetSessionToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
