package eksauth_test

// tests for the eksauth service interface for this ../../../service/eksauth/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eksauth"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eksauth/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/eksauth/eksauth_iface"
	"github.com/stretchr/testify/assert"
)

func TestEksauthServiceCanBeMocked(t *testing.T) {
	var iface eksauth_iface.IClient
	iface = &eksauth.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := eksauth.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeRoleForPodIdentity", func(t *testing.T) {
        input := &eksauth.AssumeRoleForPodIdentityInput{}
        output := &eksauth.AssumeRoleForPodIdentityOutput{}

        mockClient.On("AssumeRoleForPodIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeRoleForPodIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
