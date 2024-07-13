package licensemanagerusersubscriptions_test

// tests for the licensemanagerusersubscriptions service interface for this ../../../service/licensemanagerusersubscriptions/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanagerusersubscriptions/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanagerusersubscriptions/licensemanagerusersubscriptions_iface"
	"github.com/stretchr/testify/assert"
)

func TestLicensemanagerusersubscriptionsServiceCanBeMocked(t *testing.T) {
	var iface licensemanagerusersubscriptions_iface.IClient
	iface = &licensemanagerusersubscriptions.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := licensemanagerusersubscriptions.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateUser", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.AssociateUserInput{}
        output := &licensemanagerusersubscriptions.AssociateUserOutput{}

        mockClient.On("AssociateUser", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterIdentityProvider", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.DeregisterIdentityProviderInput{}
        output := &licensemanagerusersubscriptions.DeregisterIdentityProviderOutput{}

        mockClient.On("DeregisterIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateUser", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.DisassociateUserInput{}
        output := &licensemanagerusersubscriptions.DisassociateUserOutput{}

        mockClient.On("DisassociateUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityProviders", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.ListIdentityProvidersInput{}
        output := &licensemanagerusersubscriptions.ListIdentityProvidersOutput{}

        mockClient.On("ListIdentityProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstances", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.ListInstancesInput{}
        output := &licensemanagerusersubscriptions.ListInstancesOutput{}

        mockClient.On("ListInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProductSubscriptions", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.ListProductSubscriptionsInput{}
        output := &licensemanagerusersubscriptions.ListProductSubscriptionsOutput{}

        mockClient.On("ListProductSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListProductSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserAssociations", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.ListUserAssociationsInput{}
        output := &licensemanagerusersubscriptions.ListUserAssociationsOutput{}

        mockClient.On("ListUserAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterIdentityProvider", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.RegisterIdentityProviderInput{}
        output := &licensemanagerusersubscriptions.RegisterIdentityProviderOutput{}

        mockClient.On("RegisterIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartProductSubscription", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.StartProductSubscriptionInput{}
        output := &licensemanagerusersubscriptions.StartProductSubscriptionOutput{}

        mockClient.On("StartProductSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartProductSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopProductSubscription", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.StopProductSubscriptionInput{}
        output := &licensemanagerusersubscriptions.StopProductSubscriptionOutput{}

        mockClient.On("StopProductSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.StopProductSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityProviderSettings", func(t *testing.T) {
        input := &licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput{}
        output := &licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput{}

        mockClient.On("UpdateIdentityProviderSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityProviderSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
