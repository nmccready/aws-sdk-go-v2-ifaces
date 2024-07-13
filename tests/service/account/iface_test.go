package account_test

// tests for the account service interface for this ../../../service/account/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/account/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/account/account_iface"
	"github.com/stretchr/testify/assert"
)

func TestAccountServiceCanBeMocked(t *testing.T) {
	var iface account_iface.IClient
	iface = &account.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := account.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptPrimaryEmailUpdate", func(t *testing.T) {
        input := &account.AcceptPrimaryEmailUpdateInput{}
        output := &account.AcceptPrimaryEmailUpdateOutput{}

        mockClient.On("AcceptPrimaryEmailUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptPrimaryEmailUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlternateContact", func(t *testing.T) {
        input := &account.DeleteAlternateContactInput{}
        output := &account.DeleteAlternateContactOutput{}

        mockClient.On("DeleteAlternateContact", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlternateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableRegion", func(t *testing.T) {
        input := &account.DisableRegionInput{}
        output := &account.DisableRegionOutput{}

        mockClient.On("DisableRegion", ctx, input).Return(output, nil)

        result, err := mockClient.DisableRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableRegion", func(t *testing.T) {
        input := &account.EnableRegionInput{}
        output := &account.EnableRegionOutput{}

        mockClient.On("EnableRegion", ctx, input).Return(output, nil)

        result, err := mockClient.EnableRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAlternateContact", func(t *testing.T) {
        input := &account.GetAlternateContactInput{}
        output := &account.GetAlternateContactOutput{}

        mockClient.On("GetAlternateContact", ctx, input).Return(output, nil)

        result, err := mockClient.GetAlternateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactInformation", func(t *testing.T) {
        input := &account.GetContactInformationInput{}
        output := &account.GetContactInformationOutput{}

        mockClient.On("GetContactInformation", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrimaryEmail", func(t *testing.T) {
        input := &account.GetPrimaryEmailInput{}
        output := &account.GetPrimaryEmailOutput{}

        mockClient.On("GetPrimaryEmail", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrimaryEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegionOptStatus", func(t *testing.T) {
        input := &account.GetRegionOptStatusInput{}
        output := &account.GetRegionOptStatusOutput{}

        mockClient.On("GetRegionOptStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegionOptStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegions", func(t *testing.T) {
        input := &account.ListRegionsInput{}
        output := &account.ListRegionsOutput{}

        mockClient.On("ListRegions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAlternateContact", func(t *testing.T) {
        input := &account.PutAlternateContactInput{}
        output := &account.PutAlternateContactOutput{}

        mockClient.On("PutAlternateContact", ctx, input).Return(output, nil)

        result, err := mockClient.PutAlternateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutContactInformation", func(t *testing.T) {
        input := &account.PutContactInformationInput{}
        output := &account.PutContactInformationOutput{}

        mockClient.On("PutContactInformation", ctx, input).Return(output, nil)

        result, err := mockClient.PutContactInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPrimaryEmailUpdate", func(t *testing.T) {
        input := &account.StartPrimaryEmailUpdateInput{}
        output := &account.StartPrimaryEmailUpdateOutput{}

        mockClient.On("StartPrimaryEmailUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartPrimaryEmailUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
