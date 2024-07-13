package sso_test

// tests for the sso service interface for this ../../../service/sso/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sso/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sso/sso_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsoServiceCanBeMocked(t *testing.T) {
	var iface sso_iface.IClient
	iface = &sso.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sso.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoleCredentials", func(t *testing.T) {
        input := &sso.GetRoleCredentialsInput{}
        output := &sso.GetRoleCredentialsOutput{}

        mockClient.On("GetRoleCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoleCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountRoles", func(t *testing.T) {
        input := &sso.ListAccountRolesInput{}
        output := &sso.ListAccountRolesOutput{}

        mockClient.On("ListAccountRoles", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccounts", func(t *testing.T) {
        input := &sso.ListAccountsInput{}
        output := &sso.ListAccountsOutput{}

        mockClient.On("ListAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLogout", func(t *testing.T) {
        input := &sso.LogoutInput{}
        output := &sso.LogoutOutput{}

        mockClient.On("Logout", ctx, input).Return(output, nil)

        result, err := mockClient.Logout(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
