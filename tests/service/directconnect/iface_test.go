package directconnect_test

// tests for the directconnect service interface for this ../../../service/directconnect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/directconnect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/directconnect/directconnect_iface"
	"github.com/stretchr/testify/assert"
)

func TestDirectconnectServiceCanBeMocked(t *testing.T) {
	var iface directconnect_iface.IClient
	iface = &directconnect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := directconnect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptDirectConnectGatewayAssociationProposal", func(t *testing.T) {
        input := &directconnect.AcceptDirectConnectGatewayAssociationProposalInput{}
        output := &directconnect.AcceptDirectConnectGatewayAssociationProposalOutput{}

        mockClient.On("AcceptDirectConnectGatewayAssociationProposal", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptDirectConnectGatewayAssociationProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateConnectionOnInterconnect", func(t *testing.T) {
        input := &directconnect.AllocateConnectionOnInterconnectInput{}
        output := &directconnect.AllocateConnectionOnInterconnectOutput{}

        mockClient.On("AllocateConnectionOnInterconnect", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateConnectionOnInterconnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateHostedConnection", func(t *testing.T) {
        input := &directconnect.AllocateHostedConnectionInput{}
        output := &directconnect.AllocateHostedConnectionOutput{}

        mockClient.On("AllocateHostedConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateHostedConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocatePrivateVirtualInterface", func(t *testing.T) {
        input := &directconnect.AllocatePrivateVirtualInterfaceInput{}
        output := &directconnect.AllocatePrivateVirtualInterfaceOutput{}

        mockClient.On("AllocatePrivateVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AllocatePrivateVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocatePublicVirtualInterface", func(t *testing.T) {
        input := &directconnect.AllocatePublicVirtualInterfaceInput{}
        output := &directconnect.AllocatePublicVirtualInterfaceOutput{}

        mockClient.On("AllocatePublicVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AllocatePublicVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateTransitVirtualInterface", func(t *testing.T) {
        input := &directconnect.AllocateTransitVirtualInterfaceInput{}
        output := &directconnect.AllocateTransitVirtualInterfaceOutput{}

        mockClient.On("AllocateTransitVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateTransitVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateConnectionWithLag", func(t *testing.T) {
        input := &directconnect.AssociateConnectionWithLagInput{}
        output := &directconnect.AssociateConnectionWithLagOutput{}

        mockClient.On("AssociateConnectionWithLag", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateConnectionWithLag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateHostedConnection", func(t *testing.T) {
        input := &directconnect.AssociateHostedConnectionInput{}
        output := &directconnect.AssociateHostedConnectionOutput{}

        mockClient.On("AssociateHostedConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateHostedConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMacSecKey", func(t *testing.T) {
        input := &directconnect.AssociateMacSecKeyInput{}
        output := &directconnect.AssociateMacSecKeyOutput{}

        mockClient.On("AssociateMacSecKey", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMacSecKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateVirtualInterface", func(t *testing.T) {
        input := &directconnect.AssociateVirtualInterfaceInput{}
        output := &directconnect.AssociateVirtualInterfaceOutput{}

        mockClient.On("AssociateVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmConnection", func(t *testing.T) {
        input := &directconnect.ConfirmConnectionInput{}
        output := &directconnect.ConfirmConnectionOutput{}

        mockClient.On("ConfirmConnection", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmCustomerAgreement", func(t *testing.T) {
        input := &directconnect.ConfirmCustomerAgreementInput{}
        output := &directconnect.ConfirmCustomerAgreementOutput{}

        mockClient.On("ConfirmCustomerAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmCustomerAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmPrivateVirtualInterface", func(t *testing.T) {
        input := &directconnect.ConfirmPrivateVirtualInterfaceInput{}
        output := &directconnect.ConfirmPrivateVirtualInterfaceOutput{}

        mockClient.On("ConfirmPrivateVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmPrivateVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmPublicVirtualInterface", func(t *testing.T) {
        input := &directconnect.ConfirmPublicVirtualInterfaceInput{}
        output := &directconnect.ConfirmPublicVirtualInterfaceOutput{}

        mockClient.On("ConfirmPublicVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmPublicVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmTransitVirtualInterface", func(t *testing.T) {
        input := &directconnect.ConfirmTransitVirtualInterfaceInput{}
        output := &directconnect.ConfirmTransitVirtualInterfaceOutput{}

        mockClient.On("ConfirmTransitVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmTransitVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBGPPeer", func(t *testing.T) {
        input := &directconnect.CreateBGPPeerInput{}
        output := &directconnect.CreateBGPPeerOutput{}

        mockClient.On("CreateBGPPeer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBGPPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &directconnect.CreateConnectionInput{}
        output := &directconnect.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectConnectGateway", func(t *testing.T) {
        input := &directconnect.CreateDirectConnectGatewayInput{}
        output := &directconnect.CreateDirectConnectGatewayOutput{}

        mockClient.On("CreateDirectConnectGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectConnectGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectConnectGatewayAssociation", func(t *testing.T) {
        input := &directconnect.CreateDirectConnectGatewayAssociationInput{}
        output := &directconnect.CreateDirectConnectGatewayAssociationOutput{}

        mockClient.On("CreateDirectConnectGatewayAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectConnectGatewayAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectConnectGatewayAssociationProposal", func(t *testing.T) {
        input := &directconnect.CreateDirectConnectGatewayAssociationProposalInput{}
        output := &directconnect.CreateDirectConnectGatewayAssociationProposalOutput{}

        mockClient.On("CreateDirectConnectGatewayAssociationProposal", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectConnectGatewayAssociationProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInterconnect", func(t *testing.T) {
        input := &directconnect.CreateInterconnectInput{}
        output := &directconnect.CreateInterconnectOutput{}

        mockClient.On("CreateInterconnect", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInterconnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLag", func(t *testing.T) {
        input := &directconnect.CreateLagInput{}
        output := &directconnect.CreateLagOutput{}

        mockClient.On("CreateLag", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrivateVirtualInterface", func(t *testing.T) {
        input := &directconnect.CreatePrivateVirtualInterfaceInput{}
        output := &directconnect.CreatePrivateVirtualInterfaceOutput{}

        mockClient.On("CreatePrivateVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrivateVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePublicVirtualInterface", func(t *testing.T) {
        input := &directconnect.CreatePublicVirtualInterfaceInput{}
        output := &directconnect.CreatePublicVirtualInterfaceOutput{}

        mockClient.On("CreatePublicVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePublicVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitVirtualInterface", func(t *testing.T) {
        input := &directconnect.CreateTransitVirtualInterfaceInput{}
        output := &directconnect.CreateTransitVirtualInterfaceOutput{}

        mockClient.On("CreateTransitVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBGPPeer", func(t *testing.T) {
        input := &directconnect.DeleteBGPPeerInput{}
        output := &directconnect.DeleteBGPPeerOutput{}

        mockClient.On("DeleteBGPPeer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBGPPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &directconnect.DeleteConnectionInput{}
        output := &directconnect.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectConnectGateway", func(t *testing.T) {
        input := &directconnect.DeleteDirectConnectGatewayInput{}
        output := &directconnect.DeleteDirectConnectGatewayOutput{}

        mockClient.On("DeleteDirectConnectGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectConnectGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectConnectGatewayAssociation", func(t *testing.T) {
        input := &directconnect.DeleteDirectConnectGatewayAssociationInput{}
        output := &directconnect.DeleteDirectConnectGatewayAssociationOutput{}

        mockClient.On("DeleteDirectConnectGatewayAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectConnectGatewayAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectConnectGatewayAssociationProposal", func(t *testing.T) {
        input := &directconnect.DeleteDirectConnectGatewayAssociationProposalInput{}
        output := &directconnect.DeleteDirectConnectGatewayAssociationProposalOutput{}

        mockClient.On("DeleteDirectConnectGatewayAssociationProposal", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectConnectGatewayAssociationProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInterconnect", func(t *testing.T) {
        input := &directconnect.DeleteInterconnectInput{}
        output := &directconnect.DeleteInterconnectOutput{}

        mockClient.On("DeleteInterconnect", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInterconnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLag", func(t *testing.T) {
        input := &directconnect.DeleteLagInput{}
        output := &directconnect.DeleteLagOutput{}

        mockClient.On("DeleteLag", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualInterface", func(t *testing.T) {
        input := &directconnect.DeleteVirtualInterfaceInput{}
        output := &directconnect.DeleteVirtualInterfaceOutput{}

        mockClient.On("DeleteVirtualInterface", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectionLoa", func(t *testing.T) {
        input := &directconnect.DescribeConnectionLoaInput{}
        output := &directconnect.DescribeConnectionLoaOutput{}

        mockClient.On("DescribeConnectionLoa", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectionLoa(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnections", func(t *testing.T) {
        input := &directconnect.DescribeConnectionsInput{}
        output := &directconnect.DescribeConnectionsOutput{}

        mockClient.On("DescribeConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectionsOnInterconnect", func(t *testing.T) {
        input := &directconnect.DescribeConnectionsOnInterconnectInput{}
        output := &directconnect.DescribeConnectionsOnInterconnectOutput{}

        mockClient.On("DescribeConnectionsOnInterconnect", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectionsOnInterconnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomerMetadata", func(t *testing.T) {
        input := &directconnect.DescribeCustomerMetadataInput{}
        output := &directconnect.DescribeCustomerMetadataOutput{}

        mockClient.On("DescribeCustomerMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomerMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectConnectGatewayAssociationProposals", func(t *testing.T) {
        input := &directconnect.DescribeDirectConnectGatewayAssociationProposalsInput{}
        output := &directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput{}

        mockClient.On("DescribeDirectConnectGatewayAssociationProposals", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectConnectGatewayAssociationProposals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectConnectGatewayAssociations", func(t *testing.T) {
        input := &directconnect.DescribeDirectConnectGatewayAssociationsInput{}
        output := &directconnect.DescribeDirectConnectGatewayAssociationsOutput{}

        mockClient.On("DescribeDirectConnectGatewayAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectConnectGatewayAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectConnectGatewayAttachments", func(t *testing.T) {
        input := &directconnect.DescribeDirectConnectGatewayAttachmentsInput{}
        output := &directconnect.DescribeDirectConnectGatewayAttachmentsOutput{}

        mockClient.On("DescribeDirectConnectGatewayAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectConnectGatewayAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectConnectGateways", func(t *testing.T) {
        input := &directconnect.DescribeDirectConnectGatewaysInput{}
        output := &directconnect.DescribeDirectConnectGatewaysOutput{}

        mockClient.On("DescribeDirectConnectGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectConnectGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHostedConnections", func(t *testing.T) {
        input := &directconnect.DescribeHostedConnectionsInput{}
        output := &directconnect.DescribeHostedConnectionsOutput{}

        mockClient.On("DescribeHostedConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHostedConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInterconnectLoa", func(t *testing.T) {
        input := &directconnect.DescribeInterconnectLoaInput{}
        output := &directconnect.DescribeInterconnectLoaOutput{}

        mockClient.On("DescribeInterconnectLoa", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInterconnectLoa(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInterconnects", func(t *testing.T) {
        input := &directconnect.DescribeInterconnectsInput{}
        output := &directconnect.DescribeInterconnectsOutput{}

        mockClient.On("DescribeInterconnects", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInterconnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLags", func(t *testing.T) {
        input := &directconnect.DescribeLagsInput{}
        output := &directconnect.DescribeLagsOutput{}

        mockClient.On("DescribeLags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoa", func(t *testing.T) {
        input := &directconnect.DescribeLoaInput{}
        output := &directconnect.DescribeLoaOutput{}

        mockClient.On("DescribeLoa", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoa(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocations", func(t *testing.T) {
        input := &directconnect.DescribeLocationsInput{}
        output := &directconnect.DescribeLocationsOutput{}

        mockClient.On("DescribeLocations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRouterConfiguration", func(t *testing.T) {
        input := &directconnect.DescribeRouterConfigurationInput{}
        output := &directconnect.DescribeRouterConfigurationOutput{}

        mockClient.On("DescribeRouterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRouterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &directconnect.DescribeTagsInput{}
        output := &directconnect.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualGateways", func(t *testing.T) {
        input := &directconnect.DescribeVirtualGatewaysInput{}
        output := &directconnect.DescribeVirtualGatewaysOutput{}

        mockClient.On("DescribeVirtualGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualInterfaces", func(t *testing.T) {
        input := &directconnect.DescribeVirtualInterfacesInput{}
        output := &directconnect.DescribeVirtualInterfacesOutput{}

        mockClient.On("DescribeVirtualInterfaces", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualInterfaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateConnectionFromLag", func(t *testing.T) {
        input := &directconnect.DisassociateConnectionFromLagInput{}
        output := &directconnect.DisassociateConnectionFromLagOutput{}

        mockClient.On("DisassociateConnectionFromLag", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateConnectionFromLag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMacSecKey", func(t *testing.T) {
        input := &directconnect.DisassociateMacSecKeyInput{}
        output := &directconnect.DisassociateMacSecKeyOutput{}

        mockClient.On("DisassociateMacSecKey", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMacSecKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualInterfaceTestHistory", func(t *testing.T) {
        input := &directconnect.ListVirtualInterfaceTestHistoryInput{}
        output := &directconnect.ListVirtualInterfaceTestHistoryOutput{}

        mockClient.On("ListVirtualInterfaceTestHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualInterfaceTestHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBgpFailoverTest", func(t *testing.T) {
        input := &directconnect.StartBgpFailoverTestInput{}
        output := &directconnect.StartBgpFailoverTestOutput{}

        mockClient.On("StartBgpFailoverTest", ctx, input).Return(output, nil)

        result, err := mockClient.StartBgpFailoverTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBgpFailoverTest", func(t *testing.T) {
        input := &directconnect.StopBgpFailoverTestInput{}
        output := &directconnect.StopBgpFailoverTestOutput{}

        mockClient.On("StopBgpFailoverTest", ctx, input).Return(output, nil)

        result, err := mockClient.StopBgpFailoverTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &directconnect.TagResourceInput{}
        output := &directconnect.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &directconnect.UntagResourceInput{}
        output := &directconnect.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnection", func(t *testing.T) {
        input := &directconnect.UpdateConnectionInput{}
        output := &directconnect.UpdateConnectionOutput{}

        mockClient.On("UpdateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDirectConnectGateway", func(t *testing.T) {
        input := &directconnect.UpdateDirectConnectGatewayInput{}
        output := &directconnect.UpdateDirectConnectGatewayOutput{}

        mockClient.On("UpdateDirectConnectGateway", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDirectConnectGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDirectConnectGatewayAssociation", func(t *testing.T) {
        input := &directconnect.UpdateDirectConnectGatewayAssociationInput{}
        output := &directconnect.UpdateDirectConnectGatewayAssociationOutput{}

        mockClient.On("UpdateDirectConnectGatewayAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDirectConnectGatewayAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLag", func(t *testing.T) {
        input := &directconnect.UpdateLagInput{}
        output := &directconnect.UpdateLagOutput{}

        mockClient.On("UpdateLag", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVirtualInterfaceAttributes", func(t *testing.T) {
        input := &directconnect.UpdateVirtualInterfaceAttributesInput{}
        output := &directconnect.UpdateVirtualInterfaceAttributesOutput{}

        mockClient.On("UpdateVirtualInterfaceAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVirtualInterfaceAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
