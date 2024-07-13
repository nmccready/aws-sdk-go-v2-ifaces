package licensemanagerlinuxsubscriptions_test

// tests for the licensemanagerlinuxsubscriptions service interface for this ../../../service/licensemanagerlinuxsubscriptions/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanagerlinuxsubscriptions/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanagerlinuxsubscriptions/licensemanagerlinuxsubscriptions_iface"
	"github.com/stretchr/testify/assert"
)

func TestLicensemanagerlinuxsubscriptionsServiceCanBeMocked(t *testing.T) {
	var iface licensemanagerlinuxsubscriptions_iface.IClient
	iface = &licensemanagerlinuxsubscriptions.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := licensemanagerlinuxsubscriptions.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterSubscriptionProvider", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.DeregisterSubscriptionProviderInput{}
        output := &licensemanagerlinuxsubscriptions.DeregisterSubscriptionProviderOutput{}

        mockClient.On("DeregisterSubscriptionProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterSubscriptionProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegisteredSubscriptionProvider", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.GetRegisteredSubscriptionProviderInput{}
        output := &licensemanagerlinuxsubscriptions.GetRegisteredSubscriptionProviderOutput{}

        mockClient.On("GetRegisteredSubscriptionProvider", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegisteredSubscriptionProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceSettings", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.GetServiceSettingsInput{}
        output := &licensemanagerlinuxsubscriptions.GetServiceSettingsOutput{}

        mockClient.On("GetServiceSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLinuxSubscriptionInstances", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput{}
        output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput{}

        mockClient.On("ListLinuxSubscriptionInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListLinuxSubscriptionInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLinuxSubscriptions", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput{}
        output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput{}

        mockClient.On("ListLinuxSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLinuxSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegisteredSubscriptionProviders", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.ListRegisteredSubscriptionProvidersInput{}
        output := &licensemanagerlinuxsubscriptions.ListRegisteredSubscriptionProvidersOutput{}

        mockClient.On("ListRegisteredSubscriptionProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegisteredSubscriptionProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.ListTagsForResourceInput{}
        output := &licensemanagerlinuxsubscriptions.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterSubscriptionProvider", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.RegisterSubscriptionProviderInput{}
        output := &licensemanagerlinuxsubscriptions.RegisterSubscriptionProviderOutput{}

        mockClient.On("RegisterSubscriptionProvider", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterSubscriptionProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.TagResourceInput{}
        output := &licensemanagerlinuxsubscriptions.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.UntagResourceInput{}
        output := &licensemanagerlinuxsubscriptions.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSettings", func(t *testing.T) {
        input := &licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput{}
        output := &licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput{}

        mockClient.On("UpdateServiceSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
