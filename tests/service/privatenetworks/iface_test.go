package privatenetworks_test

// tests for the privatenetworks service interface for this ../../../service/privatenetworks/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/privatenetworks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/privatenetworks/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/privatenetworks/privatenetworks_iface"
	"github.com/stretchr/testify/assert"
)

func TestPrivatenetworksServiceCanBeMocked(t *testing.T) {
	var iface privatenetworks_iface.IClient
	iface = &privatenetworks.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := privatenetworks.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcknowledgeOrderReceipt", func(t *testing.T) {
        input := &privatenetworks.AcknowledgeOrderReceiptInput{}
        output := &privatenetworks.AcknowledgeOrderReceiptOutput{}

        mockClient.On("AcknowledgeOrderReceipt", ctx, input).Return(output, nil)

        result, err := mockClient.AcknowledgeOrderReceipt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateDeviceIdentifier", func(t *testing.T) {
        input := &privatenetworks.ActivateDeviceIdentifierInput{}
        output := &privatenetworks.ActivateDeviceIdentifierOutput{}

        mockClient.On("ActivateDeviceIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateDeviceIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateNetworkSite", func(t *testing.T) {
        input := &privatenetworks.ActivateNetworkSiteInput{}
        output := &privatenetworks.ActivateNetworkSiteOutput{}

        mockClient.On("ActivateNetworkSite", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateNetworkSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureAccessPoint", func(t *testing.T) {
        input := &privatenetworks.ConfigureAccessPointInput{}
        output := &privatenetworks.ConfigureAccessPointOutput{}

        mockClient.On("ConfigureAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetwork", func(t *testing.T) {
        input := &privatenetworks.CreateNetworkInput{}
        output := &privatenetworks.CreateNetworkOutput{}

        mockClient.On("CreateNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkSite", func(t *testing.T) {
        input := &privatenetworks.CreateNetworkSiteInput{}
        output := &privatenetworks.CreateNetworkSiteOutput{}

        mockClient.On("CreateNetworkSite", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateDeviceIdentifier", func(t *testing.T) {
        input := &privatenetworks.DeactivateDeviceIdentifierInput{}
        output := &privatenetworks.DeactivateDeviceIdentifierOutput{}

        mockClient.On("DeactivateDeviceIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateDeviceIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetwork", func(t *testing.T) {
        input := &privatenetworks.DeleteNetworkInput{}
        output := &privatenetworks.DeleteNetworkOutput{}

        mockClient.On("DeleteNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkSite", func(t *testing.T) {
        input := &privatenetworks.DeleteNetworkSiteInput{}
        output := &privatenetworks.DeleteNetworkSiteOutput{}

        mockClient.On("DeleteNetworkSite", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceIdentifier", func(t *testing.T) {
        input := &privatenetworks.GetDeviceIdentifierInput{}
        output := &privatenetworks.GetDeviceIdentifierOutput{}

        mockClient.On("GetDeviceIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetwork", func(t *testing.T) {
        input := &privatenetworks.GetNetworkInput{}
        output := &privatenetworks.GetNetworkOutput{}

        mockClient.On("GetNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkResource", func(t *testing.T) {
        input := &privatenetworks.GetNetworkResourceInput{}
        output := &privatenetworks.GetNetworkResourceOutput{}

        mockClient.On("GetNetworkResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkSite", func(t *testing.T) {
        input := &privatenetworks.GetNetworkSiteInput{}
        output := &privatenetworks.GetNetworkSiteOutput{}

        mockClient.On("GetNetworkSite", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrder", func(t *testing.T) {
        input := &privatenetworks.GetOrderInput{}
        output := &privatenetworks.GetOrderOutput{}

        mockClient.On("GetOrder", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceIdentifiers", func(t *testing.T) {
        input := &privatenetworks.ListDeviceIdentifiersInput{}
        output := &privatenetworks.ListDeviceIdentifiersOutput{}

        mockClient.On("ListDeviceIdentifiers", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceIdentifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworkResources", func(t *testing.T) {
        input := &privatenetworks.ListNetworkResourcesInput{}
        output := &privatenetworks.ListNetworkResourcesOutput{}

        mockClient.On("ListNetworkResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworkResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworkSites", func(t *testing.T) {
        input := &privatenetworks.ListNetworkSitesInput{}
        output := &privatenetworks.ListNetworkSitesOutput{}

        mockClient.On("ListNetworkSites", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworkSites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworks", func(t *testing.T) {
        input := &privatenetworks.ListNetworksInput{}
        output := &privatenetworks.ListNetworksOutput{}

        mockClient.On("ListNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrders", func(t *testing.T) {
        input := &privatenetworks.ListOrdersInput{}
        output := &privatenetworks.ListOrdersOutput{}

        mockClient.On("ListOrders", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &privatenetworks.ListTagsForResourceInput{}
        output := &privatenetworks.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPing", func(t *testing.T) {
        input := &privatenetworks.PingInput{}
        output := &privatenetworks.PingOutput{}

        mockClient.On("Ping", ctx, input).Return(output, nil)

        result, err := mockClient.Ping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNetworkResourceUpdate", func(t *testing.T) {
        input := &privatenetworks.StartNetworkResourceUpdateInput{}
        output := &privatenetworks.StartNetworkResourceUpdateOutput{}

        mockClient.On("StartNetworkResourceUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartNetworkResourceUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &privatenetworks.TagResourceInput{}
        output := &privatenetworks.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &privatenetworks.UntagResourceInput{}
        output := &privatenetworks.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkSite", func(t *testing.T) {
        input := &privatenetworks.UpdateNetworkSiteInput{}
        output := &privatenetworks.UpdateNetworkSiteOutput{}

        mockClient.On("UpdateNetworkSite", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkSitePlan", func(t *testing.T) {
        input := &privatenetworks.UpdateNetworkSitePlanInput{}
        output := &privatenetworks.UpdateNetworkSitePlanOutput{}

        mockClient.On("UpdateNetworkSitePlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkSitePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
