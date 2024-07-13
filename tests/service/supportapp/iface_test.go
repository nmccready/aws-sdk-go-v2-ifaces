package supportapp_test

// tests for the supportapp service interface for this ../../../service/supportapp/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/supportapp"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/supportapp/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/supportapp/supportapp_iface"
	"github.com/stretchr/testify/assert"
)

func TestSupportappServiceCanBeMocked(t *testing.T) {
	var iface supportapp_iface.IClient
	iface = &supportapp.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := supportapp.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSlackChannelConfiguration", func(t *testing.T) {
        input := &supportapp.CreateSlackChannelConfigurationInput{}
        output := &supportapp.CreateSlackChannelConfigurationOutput{}

        mockClient.On("CreateSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountAlias", func(t *testing.T) {
        input := &supportapp.DeleteAccountAliasInput{}
        output := &supportapp.DeleteAccountAliasOutput{}

        mockClient.On("DeleteAccountAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlackChannelConfiguration", func(t *testing.T) {
        input := &supportapp.DeleteSlackChannelConfigurationInput{}
        output := &supportapp.DeleteSlackChannelConfigurationOutput{}

        mockClient.On("DeleteSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlackWorkspaceConfiguration", func(t *testing.T) {
        input := &supportapp.DeleteSlackWorkspaceConfigurationInput{}
        output := &supportapp.DeleteSlackWorkspaceConfigurationOutput{}

        mockClient.On("DeleteSlackWorkspaceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlackWorkspaceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountAlias", func(t *testing.T) {
        input := &supportapp.GetAccountAliasInput{}
        output := &supportapp.GetAccountAliasOutput{}

        mockClient.On("GetAccountAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSlackChannelConfigurations", func(t *testing.T) {
        input := &supportapp.ListSlackChannelConfigurationsInput{}
        output := &supportapp.ListSlackChannelConfigurationsOutput{}

        mockClient.On("ListSlackChannelConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSlackChannelConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSlackWorkspaceConfigurations", func(t *testing.T) {
        input := &supportapp.ListSlackWorkspaceConfigurationsInput{}
        output := &supportapp.ListSlackWorkspaceConfigurationsOutput{}

        mockClient.On("ListSlackWorkspaceConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSlackWorkspaceConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountAlias", func(t *testing.T) {
        input := &supportapp.PutAccountAliasInput{}
        output := &supportapp.PutAccountAliasOutput{}

        mockClient.On("PutAccountAlias", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterSlackWorkspaceForOrganization", func(t *testing.T) {
        input := &supportapp.RegisterSlackWorkspaceForOrganizationInput{}
        output := &supportapp.RegisterSlackWorkspaceForOrganizationOutput{}

        mockClient.On("RegisterSlackWorkspaceForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterSlackWorkspaceForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSlackChannelConfiguration", func(t *testing.T) {
        input := &supportapp.UpdateSlackChannelConfigurationInput{}
        output := &supportapp.UpdateSlackChannelConfigurationOutput{}

        mockClient.On("UpdateSlackChannelConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSlackChannelConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
