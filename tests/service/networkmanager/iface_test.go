package networkmanager_test

// tests for the networkmanager service interface for this ../../../service/networkmanager/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkmanager/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkmanager/networkmanager_iface"
	"github.com/stretchr/testify/assert"
)

func TestNetworkmanagerServiceCanBeMocked(t *testing.T) {
	var iface networkmanager_iface.IClient
	iface = &networkmanager.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := networkmanager.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptAttachment", func(t *testing.T) {
        input := &networkmanager.AcceptAttachmentInput{}
        output := &networkmanager.AcceptAttachmentOutput{}

        mockClient.On("AcceptAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateConnectPeer", func(t *testing.T) {
        input := &networkmanager.AssociateConnectPeerInput{}
        output := &networkmanager.AssociateConnectPeerOutput{}

        mockClient.On("AssociateConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateCustomerGateway", func(t *testing.T) {
        input := &networkmanager.AssociateCustomerGatewayInput{}
        output := &networkmanager.AssociateCustomerGatewayOutput{}

        mockClient.On("AssociateCustomerGateway", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateCustomerGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLink", func(t *testing.T) {
        input := &networkmanager.AssociateLinkInput{}
        output := &networkmanager.AssociateLinkOutput{}

        mockClient.On("AssociateLink", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTransitGatewayConnectPeer", func(t *testing.T) {
        input := &networkmanager.AssociateTransitGatewayConnectPeerInput{}
        output := &networkmanager.AssociateTransitGatewayConnectPeerOutput{}

        mockClient.On("AssociateTransitGatewayConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTransitGatewayConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectAttachment", func(t *testing.T) {
        input := &networkmanager.CreateConnectAttachmentInput{}
        output := &networkmanager.CreateConnectAttachmentOutput{}

        mockClient.On("CreateConnectAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectPeer", func(t *testing.T) {
        input := &networkmanager.CreateConnectPeerInput{}
        output := &networkmanager.CreateConnectPeerOutput{}

        mockClient.On("CreateConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &networkmanager.CreateConnectionInput{}
        output := &networkmanager.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCoreNetwork", func(t *testing.T) {
        input := &networkmanager.CreateCoreNetworkInput{}
        output := &networkmanager.CreateCoreNetworkOutput{}

        mockClient.On("CreateCoreNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCoreNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDevice", func(t *testing.T) {
        input := &networkmanager.CreateDeviceInput{}
        output := &networkmanager.CreateDeviceOutput{}

        mockClient.On("CreateDevice", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalNetwork", func(t *testing.T) {
        input := &networkmanager.CreateGlobalNetworkInput{}
        output := &networkmanager.CreateGlobalNetworkOutput{}

        mockClient.On("CreateGlobalNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLink", func(t *testing.T) {
        input := &networkmanager.CreateLinkInput{}
        output := &networkmanager.CreateLinkOutput{}

        mockClient.On("CreateLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSite", func(t *testing.T) {
        input := &networkmanager.CreateSiteInput{}
        output := &networkmanager.CreateSiteOutput{}

        mockClient.On("CreateSite", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSiteToSiteVpnAttachment", func(t *testing.T) {
        input := &networkmanager.CreateSiteToSiteVpnAttachmentInput{}
        output := &networkmanager.CreateSiteToSiteVpnAttachmentOutput{}

        mockClient.On("CreateSiteToSiteVpnAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSiteToSiteVpnAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayPeering", func(t *testing.T) {
        input := &networkmanager.CreateTransitGatewayPeeringInput{}
        output := &networkmanager.CreateTransitGatewayPeeringOutput{}

        mockClient.On("CreateTransitGatewayPeering", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayPeering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayRouteTableAttachment", func(t *testing.T) {
        input := &networkmanager.CreateTransitGatewayRouteTableAttachmentInput{}
        output := &networkmanager.CreateTransitGatewayRouteTableAttachmentOutput{}

        mockClient.On("CreateTransitGatewayRouteTableAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayRouteTableAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcAttachment", func(t *testing.T) {
        input := &networkmanager.CreateVpcAttachmentInput{}
        output := &networkmanager.CreateVpcAttachmentOutput{}

        mockClient.On("CreateVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttachment", func(t *testing.T) {
        input := &networkmanager.DeleteAttachmentInput{}
        output := &networkmanager.DeleteAttachmentOutput{}

        mockClient.On("DeleteAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectPeer", func(t *testing.T) {
        input := &networkmanager.DeleteConnectPeerInput{}
        output := &networkmanager.DeleteConnectPeerOutput{}

        mockClient.On("DeleteConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &networkmanager.DeleteConnectionInput{}
        output := &networkmanager.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoreNetwork", func(t *testing.T) {
        input := &networkmanager.DeleteCoreNetworkInput{}
        output := &networkmanager.DeleteCoreNetworkOutput{}

        mockClient.On("DeleteCoreNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoreNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoreNetworkPolicyVersion", func(t *testing.T) {
        input := &networkmanager.DeleteCoreNetworkPolicyVersionInput{}
        output := &networkmanager.DeleteCoreNetworkPolicyVersionOutput{}

        mockClient.On("DeleteCoreNetworkPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoreNetworkPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevice", func(t *testing.T) {
        input := &networkmanager.DeleteDeviceInput{}
        output := &networkmanager.DeleteDeviceOutput{}

        mockClient.On("DeleteDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlobalNetwork", func(t *testing.T) {
        input := &networkmanager.DeleteGlobalNetworkInput{}
        output := &networkmanager.DeleteGlobalNetworkOutput{}

        mockClient.On("DeleteGlobalNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlobalNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLink", func(t *testing.T) {
        input := &networkmanager.DeleteLinkInput{}
        output := &networkmanager.DeleteLinkOutput{}

        mockClient.On("DeleteLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePeering", func(t *testing.T) {
        input := &networkmanager.DeletePeeringInput{}
        output := &networkmanager.DeletePeeringOutput{}

        mockClient.On("DeletePeering", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePeering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &networkmanager.DeleteResourcePolicyInput{}
        output := &networkmanager.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSite", func(t *testing.T) {
        input := &networkmanager.DeleteSiteInput{}
        output := &networkmanager.DeleteSiteOutput{}

        mockClient.On("DeleteSite", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTransitGateway", func(t *testing.T) {
        input := &networkmanager.DeregisterTransitGatewayInput{}
        output := &networkmanager.DeregisterTransitGatewayOutput{}

        mockClient.On("DeregisterTransitGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTransitGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalNetworks", func(t *testing.T) {
        input := &networkmanager.DescribeGlobalNetworksInput{}
        output := &networkmanager.DescribeGlobalNetworksOutput{}

        mockClient.On("DescribeGlobalNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateConnectPeer", func(t *testing.T) {
        input := &networkmanager.DisassociateConnectPeerInput{}
        output := &networkmanager.DisassociateConnectPeerOutput{}

        mockClient.On("DisassociateConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateCustomerGateway", func(t *testing.T) {
        input := &networkmanager.DisassociateCustomerGatewayInput{}
        output := &networkmanager.DisassociateCustomerGatewayOutput{}

        mockClient.On("DisassociateCustomerGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateCustomerGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLink", func(t *testing.T) {
        input := &networkmanager.DisassociateLinkInput{}
        output := &networkmanager.DisassociateLinkOutput{}

        mockClient.On("DisassociateLink", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTransitGatewayConnectPeer", func(t *testing.T) {
        input := &networkmanager.DisassociateTransitGatewayConnectPeerInput{}
        output := &networkmanager.DisassociateTransitGatewayConnectPeerOutput{}

        mockClient.On("DisassociateTransitGatewayConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTransitGatewayConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteCoreNetworkChangeSet", func(t *testing.T) {
        input := &networkmanager.ExecuteCoreNetworkChangeSetInput{}
        output := &networkmanager.ExecuteCoreNetworkChangeSetOutput{}

        mockClient.On("ExecuteCoreNetworkChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteCoreNetworkChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectAttachment", func(t *testing.T) {
        input := &networkmanager.GetConnectAttachmentInput{}
        output := &networkmanager.GetConnectAttachmentOutput{}

        mockClient.On("GetConnectAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectPeer", func(t *testing.T) {
        input := &networkmanager.GetConnectPeerInput{}
        output := &networkmanager.GetConnectPeerOutput{}

        mockClient.On("GetConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectPeerAssociations", func(t *testing.T) {
        input := &networkmanager.GetConnectPeerAssociationsInput{}
        output := &networkmanager.GetConnectPeerAssociationsOutput{}

        mockClient.On("GetConnectPeerAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectPeerAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnections", func(t *testing.T) {
        input := &networkmanager.GetConnectionsInput{}
        output := &networkmanager.GetConnectionsOutput{}

        mockClient.On("GetConnections", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreNetwork", func(t *testing.T) {
        input := &networkmanager.GetCoreNetworkInput{}
        output := &networkmanager.GetCoreNetworkOutput{}

        mockClient.On("GetCoreNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreNetworkChangeEvents", func(t *testing.T) {
        input := &networkmanager.GetCoreNetworkChangeEventsInput{}
        output := &networkmanager.GetCoreNetworkChangeEventsOutput{}

        mockClient.On("GetCoreNetworkChangeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreNetworkChangeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreNetworkChangeSet", func(t *testing.T) {
        input := &networkmanager.GetCoreNetworkChangeSetInput{}
        output := &networkmanager.GetCoreNetworkChangeSetOutput{}

        mockClient.On("GetCoreNetworkChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreNetworkChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreNetworkPolicy", func(t *testing.T) {
        input := &networkmanager.GetCoreNetworkPolicyInput{}
        output := &networkmanager.GetCoreNetworkPolicyOutput{}

        mockClient.On("GetCoreNetworkPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreNetworkPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomerGatewayAssociations", func(t *testing.T) {
        input := &networkmanager.GetCustomerGatewayAssociationsInput{}
        output := &networkmanager.GetCustomerGatewayAssociationsOutput{}

        mockClient.On("GetCustomerGatewayAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomerGatewayAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevices", func(t *testing.T) {
        input := &networkmanager.GetDevicesInput{}
        output := &networkmanager.GetDevicesOutput{}

        mockClient.On("GetDevices", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLinkAssociations", func(t *testing.T) {
        input := &networkmanager.GetLinkAssociationsInput{}
        output := &networkmanager.GetLinkAssociationsOutput{}

        mockClient.On("GetLinkAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetLinkAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLinks", func(t *testing.T) {
        input := &networkmanager.GetLinksInput{}
        output := &networkmanager.GetLinksOutput{}

        mockClient.On("GetLinks", ctx, input).Return(output, nil)

        result, err := mockClient.GetLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkResourceCounts", func(t *testing.T) {
        input := &networkmanager.GetNetworkResourceCountsInput{}
        output := &networkmanager.GetNetworkResourceCountsOutput{}

        mockClient.On("GetNetworkResourceCounts", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkResourceCounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkResourceRelationships", func(t *testing.T) {
        input := &networkmanager.GetNetworkResourceRelationshipsInput{}
        output := &networkmanager.GetNetworkResourceRelationshipsOutput{}

        mockClient.On("GetNetworkResourceRelationships", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkResourceRelationships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkResources", func(t *testing.T) {
        input := &networkmanager.GetNetworkResourcesInput{}
        output := &networkmanager.GetNetworkResourcesOutput{}

        mockClient.On("GetNetworkResources", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkRoutes", func(t *testing.T) {
        input := &networkmanager.GetNetworkRoutesInput{}
        output := &networkmanager.GetNetworkRoutesOutput{}

        mockClient.On("GetNetworkRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkTelemetry", func(t *testing.T) {
        input := &networkmanager.GetNetworkTelemetryInput{}
        output := &networkmanager.GetNetworkTelemetryOutput{}

        mockClient.On("GetNetworkTelemetry", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkTelemetry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &networkmanager.GetResourcePolicyInput{}
        output := &networkmanager.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRouteAnalysis", func(t *testing.T) {
        input := &networkmanager.GetRouteAnalysisInput{}
        output := &networkmanager.GetRouteAnalysisOutput{}

        mockClient.On("GetRouteAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.GetRouteAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSiteToSiteVpnAttachment", func(t *testing.T) {
        input := &networkmanager.GetSiteToSiteVpnAttachmentInput{}
        output := &networkmanager.GetSiteToSiteVpnAttachmentOutput{}

        mockClient.On("GetSiteToSiteVpnAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.GetSiteToSiteVpnAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSites", func(t *testing.T) {
        input := &networkmanager.GetSitesInput{}
        output := &networkmanager.GetSitesOutput{}

        mockClient.On("GetSites", ctx, input).Return(output, nil)

        result, err := mockClient.GetSites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayConnectPeerAssociations", func(t *testing.T) {
        input := &networkmanager.GetTransitGatewayConnectPeerAssociationsInput{}
        output := &networkmanager.GetTransitGatewayConnectPeerAssociationsOutput{}

        mockClient.On("GetTransitGatewayConnectPeerAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayConnectPeerAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayPeering", func(t *testing.T) {
        input := &networkmanager.GetTransitGatewayPeeringInput{}
        output := &networkmanager.GetTransitGatewayPeeringOutput{}

        mockClient.On("GetTransitGatewayPeering", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayPeering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayRegistrations", func(t *testing.T) {
        input := &networkmanager.GetTransitGatewayRegistrationsInput{}
        output := &networkmanager.GetTransitGatewayRegistrationsOutput{}

        mockClient.On("GetTransitGatewayRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayRouteTableAttachment", func(t *testing.T) {
        input := &networkmanager.GetTransitGatewayRouteTableAttachmentInput{}
        output := &networkmanager.GetTransitGatewayRouteTableAttachmentOutput{}

        mockClient.On("GetTransitGatewayRouteTableAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayRouteTableAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpcAttachment", func(t *testing.T) {
        input := &networkmanager.GetVpcAttachmentInput{}
        output := &networkmanager.GetVpcAttachmentOutput{}

        mockClient.On("GetVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachments", func(t *testing.T) {
        input := &networkmanager.ListAttachmentsInput{}
        output := &networkmanager.ListAttachmentsOutput{}

        mockClient.On("ListAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectPeers", func(t *testing.T) {
        input := &networkmanager.ListConnectPeersInput{}
        output := &networkmanager.ListConnectPeersOutput{}

        mockClient.On("ListConnectPeers", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectPeers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoreNetworkPolicyVersions", func(t *testing.T) {
        input := &networkmanager.ListCoreNetworkPolicyVersionsInput{}
        output := &networkmanager.ListCoreNetworkPolicyVersionsOutput{}

        mockClient.On("ListCoreNetworkPolicyVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoreNetworkPolicyVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoreNetworks", func(t *testing.T) {
        input := &networkmanager.ListCoreNetworksInput{}
        output := &networkmanager.ListCoreNetworksOutput{}

        mockClient.On("ListCoreNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoreNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationServiceAccessStatus", func(t *testing.T) {
        input := &networkmanager.ListOrganizationServiceAccessStatusInput{}
        output := &networkmanager.ListOrganizationServiceAccessStatusOutput{}

        mockClient.On("ListOrganizationServiceAccessStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationServiceAccessStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPeerings", func(t *testing.T) {
        input := &networkmanager.ListPeeringsInput{}
        output := &networkmanager.ListPeeringsOutput{}

        mockClient.On("ListPeerings", ctx, input).Return(output, nil)

        result, err := mockClient.ListPeerings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &networkmanager.ListTagsForResourceInput{}
        output := &networkmanager.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCoreNetworkPolicy", func(t *testing.T) {
        input := &networkmanager.PutCoreNetworkPolicyInput{}
        output := &networkmanager.PutCoreNetworkPolicyOutput{}

        mockClient.On("PutCoreNetworkPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutCoreNetworkPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &networkmanager.PutResourcePolicyInput{}
        output := &networkmanager.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTransitGateway", func(t *testing.T) {
        input := &networkmanager.RegisterTransitGatewayInput{}
        output := &networkmanager.RegisterTransitGatewayOutput{}

        mockClient.On("RegisterTransitGateway", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTransitGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectAttachment", func(t *testing.T) {
        input := &networkmanager.RejectAttachmentInput{}
        output := &networkmanager.RejectAttachmentOutput{}

        mockClient.On("RejectAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.RejectAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreCoreNetworkPolicyVersion", func(t *testing.T) {
        input := &networkmanager.RestoreCoreNetworkPolicyVersionInput{}
        output := &networkmanager.RestoreCoreNetworkPolicyVersionOutput{}

        mockClient.On("RestoreCoreNetworkPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreCoreNetworkPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartOrganizationServiceAccessUpdate", func(t *testing.T) {
        input := &networkmanager.StartOrganizationServiceAccessUpdateInput{}
        output := &networkmanager.StartOrganizationServiceAccessUpdateOutput{}

        mockClient.On("StartOrganizationServiceAccessUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartOrganizationServiceAccessUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRouteAnalysis", func(t *testing.T) {
        input := &networkmanager.StartRouteAnalysisInput{}
        output := &networkmanager.StartRouteAnalysisOutput{}

        mockClient.On("StartRouteAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartRouteAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &networkmanager.TagResourceInput{}
        output := &networkmanager.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &networkmanager.UntagResourceInput{}
        output := &networkmanager.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnection", func(t *testing.T) {
        input := &networkmanager.UpdateConnectionInput{}
        output := &networkmanager.UpdateConnectionOutput{}

        mockClient.On("UpdateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCoreNetwork", func(t *testing.T) {
        input := &networkmanager.UpdateCoreNetworkInput{}
        output := &networkmanager.UpdateCoreNetworkOutput{}

        mockClient.On("UpdateCoreNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCoreNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevice", func(t *testing.T) {
        input := &networkmanager.UpdateDeviceInput{}
        output := &networkmanager.UpdateDeviceOutput{}

        mockClient.On("UpdateDevice", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalNetwork", func(t *testing.T) {
        input := &networkmanager.UpdateGlobalNetworkInput{}
        output := &networkmanager.UpdateGlobalNetworkOutput{}

        mockClient.On("UpdateGlobalNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLink", func(t *testing.T) {
        input := &networkmanager.UpdateLinkInput{}
        output := &networkmanager.UpdateLinkOutput{}

        mockClient.On("UpdateLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkResourceMetadata", func(t *testing.T) {
        input := &networkmanager.UpdateNetworkResourceMetadataInput{}
        output := &networkmanager.UpdateNetworkResourceMetadataOutput{}

        mockClient.On("UpdateNetworkResourceMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkResourceMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSite", func(t *testing.T) {
        input := &networkmanager.UpdateSiteInput{}
        output := &networkmanager.UpdateSiteOutput{}

        mockClient.On("UpdateSite", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcAttachment", func(t *testing.T) {
        input := &networkmanager.UpdateVpcAttachmentInput{}
        output := &networkmanager.UpdateVpcAttachmentOutput{}

        mockClient.On("UpdateVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
