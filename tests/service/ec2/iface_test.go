package ec2_test

// tests for the ec2 service interface for this ../../../service/ec2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ec2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ec2/ec2_iface"
	"github.com/stretchr/testify/assert"
)

func TestEc2ServiceCanBeMocked(t *testing.T) {
	var iface ec2_iface.IClient
	iface = &ec2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ec2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptAddressTransfer", func(t *testing.T) {
        input := &ec2.AcceptAddressTransferInput{}
        output := &ec2.AcceptAddressTransferOutput{}

        mockClient.On("AcceptAddressTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptAddressTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptReservedInstancesExchangeQuote", func(t *testing.T) {
        input := &ec2.AcceptReservedInstancesExchangeQuoteInput{}
        output := &ec2.AcceptReservedInstancesExchangeQuoteOutput{}

        mockClient.On("AcceptReservedInstancesExchangeQuote", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptReservedInstancesExchangeQuote(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptTransitGatewayMulticastDomainAssociations", func(t *testing.T) {
        input := &ec2.AcceptTransitGatewayMulticastDomainAssociationsInput{}
        output := &ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput{}

        mockClient.On("AcceptTransitGatewayMulticastDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptTransitGatewayMulticastDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptTransitGatewayPeeringAttachment", func(t *testing.T) {
        input := &ec2.AcceptTransitGatewayPeeringAttachmentInput{}
        output := &ec2.AcceptTransitGatewayPeeringAttachmentOutput{}

        mockClient.On("AcceptTransitGatewayPeeringAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptTransitGatewayPeeringAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptTransitGatewayVpcAttachment", func(t *testing.T) {
        input := &ec2.AcceptTransitGatewayVpcAttachmentInput{}
        output := &ec2.AcceptTransitGatewayVpcAttachmentOutput{}

        mockClient.On("AcceptTransitGatewayVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptTransitGatewayVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptVpcEndpointConnections", func(t *testing.T) {
        input := &ec2.AcceptVpcEndpointConnectionsInput{}
        output := &ec2.AcceptVpcEndpointConnectionsOutput{}

        mockClient.On("AcceptVpcEndpointConnections", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptVpcEndpointConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptVpcPeeringConnection", func(t *testing.T) {
        input := &ec2.AcceptVpcPeeringConnectionInput{}
        output := &ec2.AcceptVpcPeeringConnectionOutput{}

        mockClient.On("AcceptVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdvertiseByoipCidr", func(t *testing.T) {
        input := &ec2.AdvertiseByoipCidrInput{}
        output := &ec2.AdvertiseByoipCidrOutput{}

        mockClient.On("AdvertiseByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.AdvertiseByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateAddress", func(t *testing.T) {
        input := &ec2.AllocateAddressInput{}
        output := &ec2.AllocateAddressOutput{}

        mockClient.On("AllocateAddress", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateHosts", func(t *testing.T) {
        input := &ec2.AllocateHostsInput{}
        output := &ec2.AllocateHostsOutput{}

        mockClient.On("AllocateHosts", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateIpamPoolCidr", func(t *testing.T) {
        input := &ec2.AllocateIpamPoolCidrInput{}
        output := &ec2.AllocateIpamPoolCidrOutput{}

        mockClient.On("AllocateIpamPoolCidr", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateIpamPoolCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplySecurityGroupsToClientVpnTargetNetwork", func(t *testing.T) {
        input := &ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput{}
        output := &ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput{}

        mockClient.On("ApplySecurityGroupsToClientVpnTargetNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.ApplySecurityGroupsToClientVpnTargetNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignIpv6Addresses", func(t *testing.T) {
        input := &ec2.AssignIpv6AddressesInput{}
        output := &ec2.AssignIpv6AddressesOutput{}

        mockClient.On("AssignIpv6Addresses", ctx, input).Return(output, nil)

        result, err := mockClient.AssignIpv6Addresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignPrivateIpAddresses", func(t *testing.T) {
        input := &ec2.AssignPrivateIpAddressesInput{}
        output := &ec2.AssignPrivateIpAddressesOutput{}

        mockClient.On("AssignPrivateIpAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.AssignPrivateIpAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignPrivateNatGatewayAddress", func(t *testing.T) {
        input := &ec2.AssignPrivateNatGatewayAddressInput{}
        output := &ec2.AssignPrivateNatGatewayAddressOutput{}

        mockClient.On("AssignPrivateNatGatewayAddress", ctx, input).Return(output, nil)

        result, err := mockClient.AssignPrivateNatGatewayAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAddress", func(t *testing.T) {
        input := &ec2.AssociateAddressInput{}
        output := &ec2.AssociateAddressOutput{}

        mockClient.On("AssociateAddress", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateClientVpnTargetNetwork", func(t *testing.T) {
        input := &ec2.AssociateClientVpnTargetNetworkInput{}
        output := &ec2.AssociateClientVpnTargetNetworkOutput{}

        mockClient.On("AssociateClientVpnTargetNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateClientVpnTargetNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDhcpOptions", func(t *testing.T) {
        input := &ec2.AssociateDhcpOptionsInput{}
        output := &ec2.AssociateDhcpOptionsOutput{}

        mockClient.On("AssociateDhcpOptions", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDhcpOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEnclaveCertificateIamRole", func(t *testing.T) {
        input := &ec2.AssociateEnclaveCertificateIamRoleInput{}
        output := &ec2.AssociateEnclaveCertificateIamRoleOutput{}

        mockClient.On("AssociateEnclaveCertificateIamRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEnclaveCertificateIamRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIamInstanceProfile", func(t *testing.T) {
        input := &ec2.AssociateIamInstanceProfileInput{}
        output := &ec2.AssociateIamInstanceProfileOutput{}

        mockClient.On("AssociateIamInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIamInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateInstanceEventWindow", func(t *testing.T) {
        input := &ec2.AssociateInstanceEventWindowInput{}
        output := &ec2.AssociateInstanceEventWindowOutput{}

        mockClient.On("AssociateInstanceEventWindow", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateInstanceEventWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIpamByoasn", func(t *testing.T) {
        input := &ec2.AssociateIpamByoasnInput{}
        output := &ec2.AssociateIpamByoasnOutput{}

        mockClient.On("AssociateIpamByoasn", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIpamByoasn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateIpamResourceDiscovery", func(t *testing.T) {
        input := &ec2.AssociateIpamResourceDiscoveryInput{}
        output := &ec2.AssociateIpamResourceDiscoveryOutput{}

        mockClient.On("AssociateIpamResourceDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateIpamResourceDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateNatGatewayAddress", func(t *testing.T) {
        input := &ec2.AssociateNatGatewayAddressInput{}
        output := &ec2.AssociateNatGatewayAddressOutput{}

        mockClient.On("AssociateNatGatewayAddress", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateNatGatewayAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateRouteTable", func(t *testing.T) {
        input := &ec2.AssociateRouteTableInput{}
        output := &ec2.AssociateRouteTableOutput{}

        mockClient.On("AssociateRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSubnetCidrBlock", func(t *testing.T) {
        input := &ec2.AssociateSubnetCidrBlockInput{}
        output := &ec2.AssociateSubnetCidrBlockOutput{}

        mockClient.On("AssociateSubnetCidrBlock", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSubnetCidrBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTransitGatewayMulticastDomain", func(t *testing.T) {
        input := &ec2.AssociateTransitGatewayMulticastDomainInput{}
        output := &ec2.AssociateTransitGatewayMulticastDomainOutput{}

        mockClient.On("AssociateTransitGatewayMulticastDomain", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTransitGatewayMulticastDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTransitGatewayPolicyTable", func(t *testing.T) {
        input := &ec2.AssociateTransitGatewayPolicyTableInput{}
        output := &ec2.AssociateTransitGatewayPolicyTableOutput{}

        mockClient.On("AssociateTransitGatewayPolicyTable", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTransitGatewayPolicyTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTransitGatewayRouteTable", func(t *testing.T) {
        input := &ec2.AssociateTransitGatewayRouteTableInput{}
        output := &ec2.AssociateTransitGatewayRouteTableOutput{}

        mockClient.On("AssociateTransitGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTransitGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTrunkInterface", func(t *testing.T) {
        input := &ec2.AssociateTrunkInterfaceInput{}
        output := &ec2.AssociateTrunkInterfaceOutput{}

        mockClient.On("AssociateTrunkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTrunkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateVpcCidrBlock", func(t *testing.T) {
        input := &ec2.AssociateVpcCidrBlockInput{}
        output := &ec2.AssociateVpcCidrBlockOutput{}

        mockClient.On("AssociateVpcCidrBlock", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateVpcCidrBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachClassicLinkVpc", func(t *testing.T) {
        input := &ec2.AttachClassicLinkVpcInput{}
        output := &ec2.AttachClassicLinkVpcOutput{}

        mockClient.On("AttachClassicLinkVpc", ctx, input).Return(output, nil)

        result, err := mockClient.AttachClassicLinkVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachInternetGateway", func(t *testing.T) {
        input := &ec2.AttachInternetGatewayInput{}
        output := &ec2.AttachInternetGatewayOutput{}

        mockClient.On("AttachInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.AttachInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachNetworkInterface", func(t *testing.T) {
        input := &ec2.AttachNetworkInterfaceInput{}
        output := &ec2.AttachNetworkInterfaceOutput{}

        mockClient.On("AttachNetworkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.AttachNetworkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachVerifiedAccessTrustProvider", func(t *testing.T) {
        input := &ec2.AttachVerifiedAccessTrustProviderInput{}
        output := &ec2.AttachVerifiedAccessTrustProviderOutput{}

        mockClient.On("AttachVerifiedAccessTrustProvider", ctx, input).Return(output, nil)

        result, err := mockClient.AttachVerifiedAccessTrustProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachVolume", func(t *testing.T) {
        input := &ec2.AttachVolumeInput{}
        output := &ec2.AttachVolumeOutput{}

        mockClient.On("AttachVolume", ctx, input).Return(output, nil)

        result, err := mockClient.AttachVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachVpnGateway", func(t *testing.T) {
        input := &ec2.AttachVpnGatewayInput{}
        output := &ec2.AttachVpnGatewayOutput{}

        mockClient.On("AttachVpnGateway", ctx, input).Return(output, nil)

        result, err := mockClient.AttachVpnGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeClientVpnIngress", func(t *testing.T) {
        input := &ec2.AuthorizeClientVpnIngressInput{}
        output := &ec2.AuthorizeClientVpnIngressOutput{}

        mockClient.On("AuthorizeClientVpnIngress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeClientVpnIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeSecurityGroupEgress", func(t *testing.T) {
        input := &ec2.AuthorizeSecurityGroupEgressInput{}
        output := &ec2.AuthorizeSecurityGroupEgressOutput{}

        mockClient.On("AuthorizeSecurityGroupEgress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeSecurityGroupEgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeSecurityGroupIngress", func(t *testing.T) {
        input := &ec2.AuthorizeSecurityGroupIngressInput{}
        output := &ec2.AuthorizeSecurityGroupIngressOutput{}

        mockClient.On("AuthorizeSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBundleInstance", func(t *testing.T) {
        input := &ec2.BundleInstanceInput{}
        output := &ec2.BundleInstanceOutput{}

        mockClient.On("BundleInstance", ctx, input).Return(output, nil)

        result, err := mockClient.BundleInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelBundleTask", func(t *testing.T) {
        input := &ec2.CancelBundleTaskInput{}
        output := &ec2.CancelBundleTaskOutput{}

        mockClient.On("CancelBundleTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelBundleTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCapacityReservation", func(t *testing.T) {
        input := &ec2.CancelCapacityReservationInput{}
        output := &ec2.CancelCapacityReservationOutput{}

        mockClient.On("CancelCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCapacityReservationFleets", func(t *testing.T) {
        input := &ec2.CancelCapacityReservationFleetsInput{}
        output := &ec2.CancelCapacityReservationFleetsOutput{}

        mockClient.On("CancelCapacityReservationFleets", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCapacityReservationFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelConversionTask", func(t *testing.T) {
        input := &ec2.CancelConversionTaskInput{}
        output := &ec2.CancelConversionTaskOutput{}

        mockClient.On("CancelConversionTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelConversionTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelExportTask", func(t *testing.T) {
        input := &ec2.CancelExportTaskInput{}
        output := &ec2.CancelExportTaskOutput{}

        mockClient.On("CancelExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelImageLaunchPermission", func(t *testing.T) {
        input := &ec2.CancelImageLaunchPermissionInput{}
        output := &ec2.CancelImageLaunchPermissionOutput{}

        mockClient.On("CancelImageLaunchPermission", ctx, input).Return(output, nil)

        result, err := mockClient.CancelImageLaunchPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelImportTask", func(t *testing.T) {
        input := &ec2.CancelImportTaskInput{}
        output := &ec2.CancelImportTaskOutput{}

        mockClient.On("CancelImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelReservedInstancesListing", func(t *testing.T) {
        input := &ec2.CancelReservedInstancesListingInput{}
        output := &ec2.CancelReservedInstancesListingOutput{}

        mockClient.On("CancelReservedInstancesListing", ctx, input).Return(output, nil)

        result, err := mockClient.CancelReservedInstancesListing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSpotFleetRequests", func(t *testing.T) {
        input := &ec2.CancelSpotFleetRequestsInput{}
        output := &ec2.CancelSpotFleetRequestsOutput{}

        mockClient.On("CancelSpotFleetRequests", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSpotFleetRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSpotInstanceRequests", func(t *testing.T) {
        input := &ec2.CancelSpotInstanceRequestsInput{}
        output := &ec2.CancelSpotInstanceRequestsOutput{}

        mockClient.On("CancelSpotInstanceRequests", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSpotInstanceRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmProductInstance", func(t *testing.T) {
        input := &ec2.ConfirmProductInstanceInput{}
        output := &ec2.ConfirmProductInstanceOutput{}

        mockClient.On("ConfirmProductInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmProductInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyFpgaImage", func(t *testing.T) {
        input := &ec2.CopyFpgaImageInput{}
        output := &ec2.CopyFpgaImageOutput{}

        mockClient.On("CopyFpgaImage", ctx, input).Return(output, nil)

        result, err := mockClient.CopyFpgaImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyImage", func(t *testing.T) {
        input := &ec2.CopyImageInput{}
        output := &ec2.CopyImageOutput{}

        mockClient.On("CopyImage", ctx, input).Return(output, nil)

        result, err := mockClient.CopyImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopySnapshot", func(t *testing.T) {
        input := &ec2.CopySnapshotInput{}
        output := &ec2.CopySnapshotOutput{}

        mockClient.On("CopySnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopySnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCapacityReservation", func(t *testing.T) {
        input := &ec2.CreateCapacityReservationInput{}
        output := &ec2.CreateCapacityReservationOutput{}

        mockClient.On("CreateCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCapacityReservationFleet", func(t *testing.T) {
        input := &ec2.CreateCapacityReservationFleetInput{}
        output := &ec2.CreateCapacityReservationFleetOutput{}

        mockClient.On("CreateCapacityReservationFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCapacityReservationFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCarrierGateway", func(t *testing.T) {
        input := &ec2.CreateCarrierGatewayInput{}
        output := &ec2.CreateCarrierGatewayOutput{}

        mockClient.On("CreateCarrierGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCarrierGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClientVpnEndpoint", func(t *testing.T) {
        input := &ec2.CreateClientVpnEndpointInput{}
        output := &ec2.CreateClientVpnEndpointOutput{}

        mockClient.On("CreateClientVpnEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClientVpnEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClientVpnRoute", func(t *testing.T) {
        input := &ec2.CreateClientVpnRouteInput{}
        output := &ec2.CreateClientVpnRouteOutput{}

        mockClient.On("CreateClientVpnRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClientVpnRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCoipCidr", func(t *testing.T) {
        input := &ec2.CreateCoipCidrInput{}
        output := &ec2.CreateCoipCidrOutput{}

        mockClient.On("CreateCoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCoipPool", func(t *testing.T) {
        input := &ec2.CreateCoipPoolInput{}
        output := &ec2.CreateCoipPoolOutput{}

        mockClient.On("CreateCoipPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCoipPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomerGateway", func(t *testing.T) {
        input := &ec2.CreateCustomerGatewayInput{}
        output := &ec2.CreateCustomerGatewayOutput{}

        mockClient.On("CreateCustomerGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomerGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDefaultSubnet", func(t *testing.T) {
        input := &ec2.CreateDefaultSubnetInput{}
        output := &ec2.CreateDefaultSubnetOutput{}

        mockClient.On("CreateDefaultSubnet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDefaultSubnet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDefaultVpc", func(t *testing.T) {
        input := &ec2.CreateDefaultVpcInput{}
        output := &ec2.CreateDefaultVpcOutput{}

        mockClient.On("CreateDefaultVpc", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDefaultVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDhcpOptions", func(t *testing.T) {
        input := &ec2.CreateDhcpOptionsInput{}
        output := &ec2.CreateDhcpOptionsOutput{}

        mockClient.On("CreateDhcpOptions", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDhcpOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEgressOnlyInternetGateway", func(t *testing.T) {
        input := &ec2.CreateEgressOnlyInternetGatewayInput{}
        output := &ec2.CreateEgressOnlyInternetGatewayOutput{}

        mockClient.On("CreateEgressOnlyInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEgressOnlyInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &ec2.CreateFleetInput{}
        output := &ec2.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlowLogs", func(t *testing.T) {
        input := &ec2.CreateFlowLogsInput{}
        output := &ec2.CreateFlowLogsOutput{}

        mockClient.On("CreateFlowLogs", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlowLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFpgaImage", func(t *testing.T) {
        input := &ec2.CreateFpgaImageInput{}
        output := &ec2.CreateFpgaImageOutput{}

        mockClient.On("CreateFpgaImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFpgaImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImage", func(t *testing.T) {
        input := &ec2.CreateImageInput{}
        output := &ec2.CreateImageOutput{}

        mockClient.On("CreateImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceConnectEndpoint", func(t *testing.T) {
        input := &ec2.CreateInstanceConnectEndpointInput{}
        output := &ec2.CreateInstanceConnectEndpointOutput{}

        mockClient.On("CreateInstanceConnectEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceConnectEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceEventWindow", func(t *testing.T) {
        input := &ec2.CreateInstanceEventWindowInput{}
        output := &ec2.CreateInstanceEventWindowOutput{}

        mockClient.On("CreateInstanceEventWindow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceEventWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceExportTask", func(t *testing.T) {
        input := &ec2.CreateInstanceExportTaskInput{}
        output := &ec2.CreateInstanceExportTaskOutput{}

        mockClient.On("CreateInstanceExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInternetGateway", func(t *testing.T) {
        input := &ec2.CreateInternetGatewayInput{}
        output := &ec2.CreateInternetGatewayOutput{}

        mockClient.On("CreateInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpam", func(t *testing.T) {
        input := &ec2.CreateIpamInput{}
        output := &ec2.CreateIpamOutput{}

        mockClient.On("CreateIpam", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpamPool", func(t *testing.T) {
        input := &ec2.CreateIpamPoolInput{}
        output := &ec2.CreateIpamPoolOutput{}

        mockClient.On("CreateIpamPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpamResourceDiscovery", func(t *testing.T) {
        input := &ec2.CreateIpamResourceDiscoveryInput{}
        output := &ec2.CreateIpamResourceDiscoveryOutput{}

        mockClient.On("CreateIpamResourceDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpamResourceDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIpamScope", func(t *testing.T) {
        input := &ec2.CreateIpamScopeInput{}
        output := &ec2.CreateIpamScopeOutput{}

        mockClient.On("CreateIpamScope", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIpamScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeyPair", func(t *testing.T) {
        input := &ec2.CreateKeyPairInput{}
        output := &ec2.CreateKeyPairOutput{}

        mockClient.On("CreateKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchTemplate", func(t *testing.T) {
        input := &ec2.CreateLaunchTemplateInput{}
        output := &ec2.CreateLaunchTemplateOutput{}

        mockClient.On("CreateLaunchTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchTemplateVersion", func(t *testing.T) {
        input := &ec2.CreateLaunchTemplateVersionInput{}
        output := &ec2.CreateLaunchTemplateVersionOutput{}

        mockClient.On("CreateLaunchTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocalGatewayRoute", func(t *testing.T) {
        input := &ec2.CreateLocalGatewayRouteInput{}
        output := &ec2.CreateLocalGatewayRouteOutput{}

        mockClient.On("CreateLocalGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocalGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocalGatewayRouteTable", func(t *testing.T) {
        input := &ec2.CreateLocalGatewayRouteTableInput{}
        output := &ec2.CreateLocalGatewayRouteTableOutput{}

        mockClient.On("CreateLocalGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocalGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation", func(t *testing.T) {
        input := &ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
        output := &ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput{}

        mockClient.On("CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocalGatewayRouteTableVpcAssociation", func(t *testing.T) {
        input := &ec2.CreateLocalGatewayRouteTableVpcAssociationInput{}
        output := &ec2.CreateLocalGatewayRouteTableVpcAssociationOutput{}

        mockClient.On("CreateLocalGatewayRouteTableVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocalGatewayRouteTableVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateManagedPrefixList", func(t *testing.T) {
        input := &ec2.CreateManagedPrefixListInput{}
        output := &ec2.CreateManagedPrefixListOutput{}

        mockClient.On("CreateManagedPrefixList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateManagedPrefixList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNatGateway", func(t *testing.T) {
        input := &ec2.CreateNatGatewayInput{}
        output := &ec2.CreateNatGatewayOutput{}

        mockClient.On("CreateNatGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNatGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkAcl", func(t *testing.T) {
        input := &ec2.CreateNetworkAclInput{}
        output := &ec2.CreateNetworkAclOutput{}

        mockClient.On("CreateNetworkAcl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkAclEntry", func(t *testing.T) {
        input := &ec2.CreateNetworkAclEntryInput{}
        output := &ec2.CreateNetworkAclEntryOutput{}

        mockClient.On("CreateNetworkAclEntry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkAclEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkInsightsAccessScope", func(t *testing.T) {
        input := &ec2.CreateNetworkInsightsAccessScopeInput{}
        output := &ec2.CreateNetworkInsightsAccessScopeOutput{}

        mockClient.On("CreateNetworkInsightsAccessScope", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkInsightsAccessScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkInsightsPath", func(t *testing.T) {
        input := &ec2.CreateNetworkInsightsPathInput{}
        output := &ec2.CreateNetworkInsightsPathOutput{}

        mockClient.On("CreateNetworkInsightsPath", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkInsightsPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkInterface", func(t *testing.T) {
        input := &ec2.CreateNetworkInterfaceInput{}
        output := &ec2.CreateNetworkInterfaceOutput{}

        mockClient.On("CreateNetworkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkInterfacePermission", func(t *testing.T) {
        input := &ec2.CreateNetworkInterfacePermissionInput{}
        output := &ec2.CreateNetworkInterfacePermissionOutput{}

        mockClient.On("CreateNetworkInterfacePermission", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkInterfacePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlacementGroup", func(t *testing.T) {
        input := &ec2.CreatePlacementGroupInput{}
        output := &ec2.CreatePlacementGroupOutput{}

        mockClient.On("CreatePlacementGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlacementGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePublicIpv4Pool", func(t *testing.T) {
        input := &ec2.CreatePublicIpv4PoolInput{}
        output := &ec2.CreatePublicIpv4PoolOutput{}

        mockClient.On("CreatePublicIpv4Pool", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePublicIpv4Pool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplaceRootVolumeTask", func(t *testing.T) {
        input := &ec2.CreateReplaceRootVolumeTaskInput{}
        output := &ec2.CreateReplaceRootVolumeTaskOutput{}

        mockClient.On("CreateReplaceRootVolumeTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplaceRootVolumeTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReservedInstancesListing", func(t *testing.T) {
        input := &ec2.CreateReservedInstancesListingInput{}
        output := &ec2.CreateReservedInstancesListingOutput{}

        mockClient.On("CreateReservedInstancesListing", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReservedInstancesListing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRestoreImageTask", func(t *testing.T) {
        input := &ec2.CreateRestoreImageTaskInput{}
        output := &ec2.CreateRestoreImageTaskOutput{}

        mockClient.On("CreateRestoreImageTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRestoreImageTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoute", func(t *testing.T) {
        input := &ec2.CreateRouteInput{}
        output := &ec2.CreateRouteOutput{}

        mockClient.On("CreateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRouteTable", func(t *testing.T) {
        input := &ec2.CreateRouteTableInput{}
        output := &ec2.CreateRouteTableOutput{}

        mockClient.On("CreateRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityGroup", func(t *testing.T) {
        input := &ec2.CreateSecurityGroupInput{}
        output := &ec2.CreateSecurityGroupOutput{}

        mockClient.On("CreateSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &ec2.CreateSnapshotInput{}
        output := &ec2.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshots", func(t *testing.T) {
        input := &ec2.CreateSnapshotsInput{}
        output := &ec2.CreateSnapshotsOutput{}

        mockClient.On("CreateSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSpotDatafeedSubscription", func(t *testing.T) {
        input := &ec2.CreateSpotDatafeedSubscriptionInput{}
        output := &ec2.CreateSpotDatafeedSubscriptionOutput{}

        mockClient.On("CreateSpotDatafeedSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSpotDatafeedSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStoreImageTask", func(t *testing.T) {
        input := &ec2.CreateStoreImageTaskInput{}
        output := &ec2.CreateStoreImageTaskOutput{}

        mockClient.On("CreateStoreImageTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStoreImageTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubnet", func(t *testing.T) {
        input := &ec2.CreateSubnetInput{}
        output := &ec2.CreateSubnetOutput{}

        mockClient.On("CreateSubnet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubnet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubnetCidrReservation", func(t *testing.T) {
        input := &ec2.CreateSubnetCidrReservationInput{}
        output := &ec2.CreateSubnetCidrReservationOutput{}

        mockClient.On("CreateSubnetCidrReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubnetCidrReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &ec2.CreateTagsInput{}
        output := &ec2.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficMirrorFilter", func(t *testing.T) {
        input := &ec2.CreateTrafficMirrorFilterInput{}
        output := &ec2.CreateTrafficMirrorFilterOutput{}

        mockClient.On("CreateTrafficMirrorFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficMirrorFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficMirrorFilterRule", func(t *testing.T) {
        input := &ec2.CreateTrafficMirrorFilterRuleInput{}
        output := &ec2.CreateTrafficMirrorFilterRuleOutput{}

        mockClient.On("CreateTrafficMirrorFilterRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficMirrorFilterRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficMirrorSession", func(t *testing.T) {
        input := &ec2.CreateTrafficMirrorSessionInput{}
        output := &ec2.CreateTrafficMirrorSessionOutput{}

        mockClient.On("CreateTrafficMirrorSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficMirrorSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficMirrorTarget", func(t *testing.T) {
        input := &ec2.CreateTrafficMirrorTargetInput{}
        output := &ec2.CreateTrafficMirrorTargetOutput{}

        mockClient.On("CreateTrafficMirrorTarget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficMirrorTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGateway", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayInput{}
        output := &ec2.CreateTransitGatewayOutput{}

        mockClient.On("CreateTransitGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayConnect", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayConnectInput{}
        output := &ec2.CreateTransitGatewayConnectOutput{}

        mockClient.On("CreateTransitGatewayConnect", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayConnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayConnectPeer", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayConnectPeerInput{}
        output := &ec2.CreateTransitGatewayConnectPeerOutput{}

        mockClient.On("CreateTransitGatewayConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayMulticastDomain", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayMulticastDomainInput{}
        output := &ec2.CreateTransitGatewayMulticastDomainOutput{}

        mockClient.On("CreateTransitGatewayMulticastDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayMulticastDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayPeeringAttachment", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayPeeringAttachmentInput{}
        output := &ec2.CreateTransitGatewayPeeringAttachmentOutput{}

        mockClient.On("CreateTransitGatewayPeeringAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayPeeringAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayPolicyTable", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayPolicyTableInput{}
        output := &ec2.CreateTransitGatewayPolicyTableOutput{}

        mockClient.On("CreateTransitGatewayPolicyTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayPolicyTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayPrefixListReference", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayPrefixListReferenceInput{}
        output := &ec2.CreateTransitGatewayPrefixListReferenceOutput{}

        mockClient.On("CreateTransitGatewayPrefixListReference", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayPrefixListReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayRoute", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayRouteInput{}
        output := &ec2.CreateTransitGatewayRouteOutput{}

        mockClient.On("CreateTransitGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayRouteTable", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayRouteTableInput{}
        output := &ec2.CreateTransitGatewayRouteTableOutput{}

        mockClient.On("CreateTransitGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayRouteTableAnnouncement", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayRouteTableAnnouncementInput{}
        output := &ec2.CreateTransitGatewayRouteTableAnnouncementOutput{}

        mockClient.On("CreateTransitGatewayRouteTableAnnouncement", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayRouteTableAnnouncement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransitGatewayVpcAttachment", func(t *testing.T) {
        input := &ec2.CreateTransitGatewayVpcAttachmentInput{}
        output := &ec2.CreateTransitGatewayVpcAttachmentOutput{}

        mockClient.On("CreateTransitGatewayVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransitGatewayVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVerifiedAccessEndpoint", func(t *testing.T) {
        input := &ec2.CreateVerifiedAccessEndpointInput{}
        output := &ec2.CreateVerifiedAccessEndpointOutput{}

        mockClient.On("CreateVerifiedAccessEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVerifiedAccessEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVerifiedAccessGroup", func(t *testing.T) {
        input := &ec2.CreateVerifiedAccessGroupInput{}
        output := &ec2.CreateVerifiedAccessGroupOutput{}

        mockClient.On("CreateVerifiedAccessGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVerifiedAccessGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVerifiedAccessInstance", func(t *testing.T) {
        input := &ec2.CreateVerifiedAccessInstanceInput{}
        output := &ec2.CreateVerifiedAccessInstanceOutput{}

        mockClient.On("CreateVerifiedAccessInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVerifiedAccessInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVerifiedAccessTrustProvider", func(t *testing.T) {
        input := &ec2.CreateVerifiedAccessTrustProviderInput{}
        output := &ec2.CreateVerifiedAccessTrustProviderOutput{}

        mockClient.On("CreateVerifiedAccessTrustProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVerifiedAccessTrustProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVolume", func(t *testing.T) {
        input := &ec2.CreateVolumeInput{}
        output := &ec2.CreateVolumeOutput{}

        mockClient.On("CreateVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpc", func(t *testing.T) {
        input := &ec2.CreateVpcInput{}
        output := &ec2.CreateVpcOutput{}

        mockClient.On("CreateVpc", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpoint", func(t *testing.T) {
        input := &ec2.CreateVpcEndpointInput{}
        output := &ec2.CreateVpcEndpointOutput{}

        mockClient.On("CreateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpointConnectionNotification", func(t *testing.T) {
        input := &ec2.CreateVpcEndpointConnectionNotificationInput{}
        output := &ec2.CreateVpcEndpointConnectionNotificationOutput{}

        mockClient.On("CreateVpcEndpointConnectionNotification", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpointConnectionNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpointServiceConfiguration", func(t *testing.T) {
        input := &ec2.CreateVpcEndpointServiceConfigurationInput{}
        output := &ec2.CreateVpcEndpointServiceConfigurationOutput{}

        mockClient.On("CreateVpcEndpointServiceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpointServiceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcPeeringConnection", func(t *testing.T) {
        input := &ec2.CreateVpcPeeringConnectionInput{}
        output := &ec2.CreateVpcPeeringConnectionOutput{}

        mockClient.On("CreateVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpnConnection", func(t *testing.T) {
        input := &ec2.CreateVpnConnectionInput{}
        output := &ec2.CreateVpnConnectionOutput{}

        mockClient.On("CreateVpnConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpnConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpnConnectionRoute", func(t *testing.T) {
        input := &ec2.CreateVpnConnectionRouteInput{}
        output := &ec2.CreateVpnConnectionRouteOutput{}

        mockClient.On("CreateVpnConnectionRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpnConnectionRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpnGateway", func(t *testing.T) {
        input := &ec2.CreateVpnGatewayInput{}
        output := &ec2.CreateVpnGatewayOutput{}

        mockClient.On("CreateVpnGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpnGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCarrierGateway", func(t *testing.T) {
        input := &ec2.DeleteCarrierGatewayInput{}
        output := &ec2.DeleteCarrierGatewayOutput{}

        mockClient.On("DeleteCarrierGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCarrierGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClientVpnEndpoint", func(t *testing.T) {
        input := &ec2.DeleteClientVpnEndpointInput{}
        output := &ec2.DeleteClientVpnEndpointOutput{}

        mockClient.On("DeleteClientVpnEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClientVpnEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClientVpnRoute", func(t *testing.T) {
        input := &ec2.DeleteClientVpnRouteInput{}
        output := &ec2.DeleteClientVpnRouteOutput{}

        mockClient.On("DeleteClientVpnRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClientVpnRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoipCidr", func(t *testing.T) {
        input := &ec2.DeleteCoipCidrInput{}
        output := &ec2.DeleteCoipCidrOutput{}

        mockClient.On("DeleteCoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoipPool", func(t *testing.T) {
        input := &ec2.DeleteCoipPoolInput{}
        output := &ec2.DeleteCoipPoolOutput{}

        mockClient.On("DeleteCoipPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoipPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomerGateway", func(t *testing.T) {
        input := &ec2.DeleteCustomerGatewayInput{}
        output := &ec2.DeleteCustomerGatewayOutput{}

        mockClient.On("DeleteCustomerGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomerGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDhcpOptions", func(t *testing.T) {
        input := &ec2.DeleteDhcpOptionsInput{}
        output := &ec2.DeleteDhcpOptionsOutput{}

        mockClient.On("DeleteDhcpOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDhcpOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEgressOnlyInternetGateway", func(t *testing.T) {
        input := &ec2.DeleteEgressOnlyInternetGatewayInput{}
        output := &ec2.DeleteEgressOnlyInternetGatewayOutput{}

        mockClient.On("DeleteEgressOnlyInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEgressOnlyInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleets", func(t *testing.T) {
        input := &ec2.DeleteFleetsInput{}
        output := &ec2.DeleteFleetsOutput{}

        mockClient.On("DeleteFleets", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlowLogs", func(t *testing.T) {
        input := &ec2.DeleteFlowLogsInput{}
        output := &ec2.DeleteFlowLogsOutput{}

        mockClient.On("DeleteFlowLogs", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlowLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFpgaImage", func(t *testing.T) {
        input := &ec2.DeleteFpgaImageInput{}
        output := &ec2.DeleteFpgaImageOutput{}

        mockClient.On("DeleteFpgaImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFpgaImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceConnectEndpoint", func(t *testing.T) {
        input := &ec2.DeleteInstanceConnectEndpointInput{}
        output := &ec2.DeleteInstanceConnectEndpointOutput{}

        mockClient.On("DeleteInstanceConnectEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceConnectEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceEventWindow", func(t *testing.T) {
        input := &ec2.DeleteInstanceEventWindowInput{}
        output := &ec2.DeleteInstanceEventWindowOutput{}

        mockClient.On("DeleteInstanceEventWindow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceEventWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInternetGateway", func(t *testing.T) {
        input := &ec2.DeleteInternetGatewayInput{}
        output := &ec2.DeleteInternetGatewayOutput{}

        mockClient.On("DeleteInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpam", func(t *testing.T) {
        input := &ec2.DeleteIpamInput{}
        output := &ec2.DeleteIpamOutput{}

        mockClient.On("DeleteIpam", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpamPool", func(t *testing.T) {
        input := &ec2.DeleteIpamPoolInput{}
        output := &ec2.DeleteIpamPoolOutput{}

        mockClient.On("DeleteIpamPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpamResourceDiscovery", func(t *testing.T) {
        input := &ec2.DeleteIpamResourceDiscoveryInput{}
        output := &ec2.DeleteIpamResourceDiscoveryOutput{}

        mockClient.On("DeleteIpamResourceDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpamResourceDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIpamScope", func(t *testing.T) {
        input := &ec2.DeleteIpamScopeInput{}
        output := &ec2.DeleteIpamScopeOutput{}

        mockClient.On("DeleteIpamScope", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIpamScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyPair", func(t *testing.T) {
        input := &ec2.DeleteKeyPairInput{}
        output := &ec2.DeleteKeyPairOutput{}

        mockClient.On("DeleteKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchTemplate", func(t *testing.T) {
        input := &ec2.DeleteLaunchTemplateInput{}
        output := &ec2.DeleteLaunchTemplateOutput{}

        mockClient.On("DeleteLaunchTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchTemplateVersions", func(t *testing.T) {
        input := &ec2.DeleteLaunchTemplateVersionsInput{}
        output := &ec2.DeleteLaunchTemplateVersionsOutput{}

        mockClient.On("DeleteLaunchTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocalGatewayRoute", func(t *testing.T) {
        input := &ec2.DeleteLocalGatewayRouteInput{}
        output := &ec2.DeleteLocalGatewayRouteOutput{}

        mockClient.On("DeleteLocalGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocalGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocalGatewayRouteTable", func(t *testing.T) {
        input := &ec2.DeleteLocalGatewayRouteTableInput{}
        output := &ec2.DeleteLocalGatewayRouteTableOutput{}

        mockClient.On("DeleteLocalGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocalGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation", func(t *testing.T) {
        input := &ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
        output := &ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput{}

        mockClient.On("DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocalGatewayRouteTableVpcAssociation", func(t *testing.T) {
        input := &ec2.DeleteLocalGatewayRouteTableVpcAssociationInput{}
        output := &ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput{}

        mockClient.On("DeleteLocalGatewayRouteTableVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocalGatewayRouteTableVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteManagedPrefixList", func(t *testing.T) {
        input := &ec2.DeleteManagedPrefixListInput{}
        output := &ec2.DeleteManagedPrefixListOutput{}

        mockClient.On("DeleteManagedPrefixList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteManagedPrefixList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNatGateway", func(t *testing.T) {
        input := &ec2.DeleteNatGatewayInput{}
        output := &ec2.DeleteNatGatewayOutput{}

        mockClient.On("DeleteNatGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNatGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkAcl", func(t *testing.T) {
        input := &ec2.DeleteNetworkAclInput{}
        output := &ec2.DeleteNetworkAclOutput{}

        mockClient.On("DeleteNetworkAcl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkAclEntry", func(t *testing.T) {
        input := &ec2.DeleteNetworkAclEntryInput{}
        output := &ec2.DeleteNetworkAclEntryOutput{}

        mockClient.On("DeleteNetworkAclEntry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkAclEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInsightsAccessScope", func(t *testing.T) {
        input := &ec2.DeleteNetworkInsightsAccessScopeInput{}
        output := &ec2.DeleteNetworkInsightsAccessScopeOutput{}

        mockClient.On("DeleteNetworkInsightsAccessScope", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInsightsAccessScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInsightsAccessScopeAnalysis", func(t *testing.T) {
        input := &ec2.DeleteNetworkInsightsAccessScopeAnalysisInput{}
        output := &ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput{}

        mockClient.On("DeleteNetworkInsightsAccessScopeAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInsightsAccessScopeAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInsightsAnalysis", func(t *testing.T) {
        input := &ec2.DeleteNetworkInsightsAnalysisInput{}
        output := &ec2.DeleteNetworkInsightsAnalysisOutput{}

        mockClient.On("DeleteNetworkInsightsAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInsightsAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInsightsPath", func(t *testing.T) {
        input := &ec2.DeleteNetworkInsightsPathInput{}
        output := &ec2.DeleteNetworkInsightsPathOutput{}

        mockClient.On("DeleteNetworkInsightsPath", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInsightsPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInterface", func(t *testing.T) {
        input := &ec2.DeleteNetworkInterfaceInput{}
        output := &ec2.DeleteNetworkInterfaceOutput{}

        mockClient.On("DeleteNetworkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkInterfacePermission", func(t *testing.T) {
        input := &ec2.DeleteNetworkInterfacePermissionInput{}
        output := &ec2.DeleteNetworkInterfacePermissionOutput{}

        mockClient.On("DeleteNetworkInterfacePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkInterfacePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlacementGroup", func(t *testing.T) {
        input := &ec2.DeletePlacementGroupInput{}
        output := &ec2.DeletePlacementGroupOutput{}

        mockClient.On("DeletePlacementGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlacementGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublicIpv4Pool", func(t *testing.T) {
        input := &ec2.DeletePublicIpv4PoolInput{}
        output := &ec2.DeletePublicIpv4PoolOutput{}

        mockClient.On("DeletePublicIpv4Pool", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublicIpv4Pool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueuedReservedInstances", func(t *testing.T) {
        input := &ec2.DeleteQueuedReservedInstancesInput{}
        output := &ec2.DeleteQueuedReservedInstancesOutput{}

        mockClient.On("DeleteQueuedReservedInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueuedReservedInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoute", func(t *testing.T) {
        input := &ec2.DeleteRouteInput{}
        output := &ec2.DeleteRouteOutput{}

        mockClient.On("DeleteRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRouteTable", func(t *testing.T) {
        input := &ec2.DeleteRouteTableInput{}
        output := &ec2.DeleteRouteTableOutput{}

        mockClient.On("DeleteRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityGroup", func(t *testing.T) {
        input := &ec2.DeleteSecurityGroupInput{}
        output := &ec2.DeleteSecurityGroupOutput{}

        mockClient.On("DeleteSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &ec2.DeleteSnapshotInput{}
        output := &ec2.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSpotDatafeedSubscription", func(t *testing.T) {
        input := &ec2.DeleteSpotDatafeedSubscriptionInput{}
        output := &ec2.DeleteSpotDatafeedSubscriptionOutput{}

        mockClient.On("DeleteSpotDatafeedSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSpotDatafeedSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubnet", func(t *testing.T) {
        input := &ec2.DeleteSubnetInput{}
        output := &ec2.DeleteSubnetOutput{}

        mockClient.On("DeleteSubnet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubnet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubnetCidrReservation", func(t *testing.T) {
        input := &ec2.DeleteSubnetCidrReservationInput{}
        output := &ec2.DeleteSubnetCidrReservationOutput{}

        mockClient.On("DeleteSubnetCidrReservation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubnetCidrReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &ec2.DeleteTagsInput{}
        output := &ec2.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficMirrorFilter", func(t *testing.T) {
        input := &ec2.DeleteTrafficMirrorFilterInput{}
        output := &ec2.DeleteTrafficMirrorFilterOutput{}

        mockClient.On("DeleteTrafficMirrorFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficMirrorFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficMirrorFilterRule", func(t *testing.T) {
        input := &ec2.DeleteTrafficMirrorFilterRuleInput{}
        output := &ec2.DeleteTrafficMirrorFilterRuleOutput{}

        mockClient.On("DeleteTrafficMirrorFilterRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficMirrorFilterRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficMirrorSession", func(t *testing.T) {
        input := &ec2.DeleteTrafficMirrorSessionInput{}
        output := &ec2.DeleteTrafficMirrorSessionOutput{}

        mockClient.On("DeleteTrafficMirrorSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficMirrorSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficMirrorTarget", func(t *testing.T) {
        input := &ec2.DeleteTrafficMirrorTargetInput{}
        output := &ec2.DeleteTrafficMirrorTargetOutput{}

        mockClient.On("DeleteTrafficMirrorTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficMirrorTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGateway", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayInput{}
        output := &ec2.DeleteTransitGatewayOutput{}

        mockClient.On("DeleteTransitGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayConnect", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayConnectInput{}
        output := &ec2.DeleteTransitGatewayConnectOutput{}

        mockClient.On("DeleteTransitGatewayConnect", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayConnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayConnectPeer", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayConnectPeerInput{}
        output := &ec2.DeleteTransitGatewayConnectPeerOutput{}

        mockClient.On("DeleteTransitGatewayConnectPeer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayConnectPeer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayMulticastDomain", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayMulticastDomainInput{}
        output := &ec2.DeleteTransitGatewayMulticastDomainOutput{}

        mockClient.On("DeleteTransitGatewayMulticastDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayMulticastDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayPeeringAttachment", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayPeeringAttachmentInput{}
        output := &ec2.DeleteTransitGatewayPeeringAttachmentOutput{}

        mockClient.On("DeleteTransitGatewayPeeringAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayPeeringAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayPolicyTable", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayPolicyTableInput{}
        output := &ec2.DeleteTransitGatewayPolicyTableOutput{}

        mockClient.On("DeleteTransitGatewayPolicyTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayPolicyTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayPrefixListReference", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayPrefixListReferenceInput{}
        output := &ec2.DeleteTransitGatewayPrefixListReferenceOutput{}

        mockClient.On("DeleteTransitGatewayPrefixListReference", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayPrefixListReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayRoute", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayRouteInput{}
        output := &ec2.DeleteTransitGatewayRouteOutput{}

        mockClient.On("DeleteTransitGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayRouteTable", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayRouteTableInput{}
        output := &ec2.DeleteTransitGatewayRouteTableOutput{}

        mockClient.On("DeleteTransitGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayRouteTableAnnouncement", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayRouteTableAnnouncementInput{}
        output := &ec2.DeleteTransitGatewayRouteTableAnnouncementOutput{}

        mockClient.On("DeleteTransitGatewayRouteTableAnnouncement", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayRouteTableAnnouncement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransitGatewayVpcAttachment", func(t *testing.T) {
        input := &ec2.DeleteTransitGatewayVpcAttachmentInput{}
        output := &ec2.DeleteTransitGatewayVpcAttachmentOutput{}

        mockClient.On("DeleteTransitGatewayVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransitGatewayVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedAccessEndpoint", func(t *testing.T) {
        input := &ec2.DeleteVerifiedAccessEndpointInput{}
        output := &ec2.DeleteVerifiedAccessEndpointOutput{}

        mockClient.On("DeleteVerifiedAccessEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedAccessEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedAccessGroup", func(t *testing.T) {
        input := &ec2.DeleteVerifiedAccessGroupInput{}
        output := &ec2.DeleteVerifiedAccessGroupOutput{}

        mockClient.On("DeleteVerifiedAccessGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedAccessGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedAccessInstance", func(t *testing.T) {
        input := &ec2.DeleteVerifiedAccessInstanceInput{}
        output := &ec2.DeleteVerifiedAccessInstanceOutput{}

        mockClient.On("DeleteVerifiedAccessInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedAccessInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedAccessTrustProvider", func(t *testing.T) {
        input := &ec2.DeleteVerifiedAccessTrustProviderInput{}
        output := &ec2.DeleteVerifiedAccessTrustProviderOutput{}

        mockClient.On("DeleteVerifiedAccessTrustProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedAccessTrustProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVolume", func(t *testing.T) {
        input := &ec2.DeleteVolumeInput{}
        output := &ec2.DeleteVolumeOutput{}

        mockClient.On("DeleteVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpc", func(t *testing.T) {
        input := &ec2.DeleteVpcInput{}
        output := &ec2.DeleteVpcOutput{}

        mockClient.On("DeleteVpc", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpointConnectionNotifications", func(t *testing.T) {
        input := &ec2.DeleteVpcEndpointConnectionNotificationsInput{}
        output := &ec2.DeleteVpcEndpointConnectionNotificationsOutput{}

        mockClient.On("DeleteVpcEndpointConnectionNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpointConnectionNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpointServiceConfigurations", func(t *testing.T) {
        input := &ec2.DeleteVpcEndpointServiceConfigurationsInput{}
        output := &ec2.DeleteVpcEndpointServiceConfigurationsOutput{}

        mockClient.On("DeleteVpcEndpointServiceConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpointServiceConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpoints", func(t *testing.T) {
        input := &ec2.DeleteVpcEndpointsInput{}
        output := &ec2.DeleteVpcEndpointsOutput{}

        mockClient.On("DeleteVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcPeeringConnection", func(t *testing.T) {
        input := &ec2.DeleteVpcPeeringConnectionInput{}
        output := &ec2.DeleteVpcPeeringConnectionOutput{}

        mockClient.On("DeleteVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpnConnection", func(t *testing.T) {
        input := &ec2.DeleteVpnConnectionInput{}
        output := &ec2.DeleteVpnConnectionOutput{}

        mockClient.On("DeleteVpnConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpnConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpnConnectionRoute", func(t *testing.T) {
        input := &ec2.DeleteVpnConnectionRouteInput{}
        output := &ec2.DeleteVpnConnectionRouteOutput{}

        mockClient.On("DeleteVpnConnectionRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpnConnectionRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpnGateway", func(t *testing.T) {
        input := &ec2.DeleteVpnGatewayInput{}
        output := &ec2.DeleteVpnGatewayOutput{}

        mockClient.On("DeleteVpnGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpnGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprovisionByoipCidr", func(t *testing.T) {
        input := &ec2.DeprovisionByoipCidrInput{}
        output := &ec2.DeprovisionByoipCidrOutput{}

        mockClient.On("DeprovisionByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.DeprovisionByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprovisionIpamByoasn", func(t *testing.T) {
        input := &ec2.DeprovisionIpamByoasnInput{}
        output := &ec2.DeprovisionIpamByoasnOutput{}

        mockClient.On("DeprovisionIpamByoasn", ctx, input).Return(output, nil)

        result, err := mockClient.DeprovisionIpamByoasn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprovisionIpamPoolCidr", func(t *testing.T) {
        input := &ec2.DeprovisionIpamPoolCidrInput{}
        output := &ec2.DeprovisionIpamPoolCidrOutput{}

        mockClient.On("DeprovisionIpamPoolCidr", ctx, input).Return(output, nil)

        result, err := mockClient.DeprovisionIpamPoolCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprovisionPublicIpv4PoolCidr", func(t *testing.T) {
        input := &ec2.DeprovisionPublicIpv4PoolCidrInput{}
        output := &ec2.DeprovisionPublicIpv4PoolCidrOutput{}

        mockClient.On("DeprovisionPublicIpv4PoolCidr", ctx, input).Return(output, nil)

        result, err := mockClient.DeprovisionPublicIpv4PoolCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterImage", func(t *testing.T) {
        input := &ec2.DeregisterImageInput{}
        output := &ec2.DeregisterImageOutput{}

        mockClient.On("DeregisterImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterInstanceEventNotificationAttributes", func(t *testing.T) {
        input := &ec2.DeregisterInstanceEventNotificationAttributesInput{}
        output := &ec2.DeregisterInstanceEventNotificationAttributesOutput{}

        mockClient.On("DeregisterInstanceEventNotificationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterInstanceEventNotificationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTransitGatewayMulticastGroupMembers", func(t *testing.T) {
        input := &ec2.DeregisterTransitGatewayMulticastGroupMembersInput{}
        output := &ec2.DeregisterTransitGatewayMulticastGroupMembersOutput{}

        mockClient.On("DeregisterTransitGatewayMulticastGroupMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTransitGatewayMulticastGroupMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTransitGatewayMulticastGroupSources", func(t *testing.T) {
        input := &ec2.DeregisterTransitGatewayMulticastGroupSourcesInput{}
        output := &ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput{}

        mockClient.On("DeregisterTransitGatewayMulticastGroupSources", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTransitGatewayMulticastGroupSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &ec2.DescribeAccountAttributesInput{}
        output := &ec2.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddressTransfers", func(t *testing.T) {
        input := &ec2.DescribeAddressTransfersInput{}
        output := &ec2.DescribeAddressTransfersOutput{}

        mockClient.On("DescribeAddressTransfers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddressTransfers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddresses", func(t *testing.T) {
        input := &ec2.DescribeAddressesInput{}
        output := &ec2.DescribeAddressesOutput{}

        mockClient.On("DescribeAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddressesAttribute", func(t *testing.T) {
        input := &ec2.DescribeAddressesAttributeInput{}
        output := &ec2.DescribeAddressesAttributeOutput{}

        mockClient.On("DescribeAddressesAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddressesAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAggregateIdFormat", func(t *testing.T) {
        input := &ec2.DescribeAggregateIdFormatInput{}
        output := &ec2.DescribeAggregateIdFormatOutput{}

        mockClient.On("DescribeAggregateIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAggregateIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAvailabilityZones", func(t *testing.T) {
        input := &ec2.DescribeAvailabilityZonesInput{}
        output := &ec2.DescribeAvailabilityZonesOutput{}

        mockClient.On("DescribeAvailabilityZones", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAvailabilityZones(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAwsNetworkPerformanceMetricSubscriptions", func(t *testing.T) {
        input := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput{}
        output := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput{}

        mockClient.On("DescribeAwsNetworkPerformanceMetricSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAwsNetworkPerformanceMetricSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBundleTasks", func(t *testing.T) {
        input := &ec2.DescribeBundleTasksInput{}
        output := &ec2.DescribeBundleTasksOutput{}

        mockClient.On("DescribeBundleTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBundleTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeByoipCidrs", func(t *testing.T) {
        input := &ec2.DescribeByoipCidrsInput{}
        output := &ec2.DescribeByoipCidrsOutput{}

        mockClient.On("DescribeByoipCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeByoipCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCapacityBlockOfferings", func(t *testing.T) {
        input := &ec2.DescribeCapacityBlockOfferingsInput{}
        output := &ec2.DescribeCapacityBlockOfferingsOutput{}

        mockClient.On("DescribeCapacityBlockOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCapacityBlockOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCapacityReservationFleets", func(t *testing.T) {
        input := &ec2.DescribeCapacityReservationFleetsInput{}
        output := &ec2.DescribeCapacityReservationFleetsOutput{}

        mockClient.On("DescribeCapacityReservationFleets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCapacityReservationFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCapacityReservations", func(t *testing.T) {
        input := &ec2.DescribeCapacityReservationsInput{}
        output := &ec2.DescribeCapacityReservationsOutput{}

        mockClient.On("DescribeCapacityReservations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCapacityReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCarrierGateways", func(t *testing.T) {
        input := &ec2.DescribeCarrierGatewaysInput{}
        output := &ec2.DescribeCarrierGatewaysOutput{}

        mockClient.On("DescribeCarrierGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCarrierGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClassicLinkInstances", func(t *testing.T) {
        input := &ec2.DescribeClassicLinkInstancesInput{}
        output := &ec2.DescribeClassicLinkInstancesOutput{}

        mockClient.On("DescribeClassicLinkInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClassicLinkInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientVpnAuthorizationRules", func(t *testing.T) {
        input := &ec2.DescribeClientVpnAuthorizationRulesInput{}
        output := &ec2.DescribeClientVpnAuthorizationRulesOutput{}

        mockClient.On("DescribeClientVpnAuthorizationRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientVpnAuthorizationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientVpnConnections", func(t *testing.T) {
        input := &ec2.DescribeClientVpnConnectionsInput{}
        output := &ec2.DescribeClientVpnConnectionsOutput{}

        mockClient.On("DescribeClientVpnConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientVpnConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientVpnEndpoints", func(t *testing.T) {
        input := &ec2.DescribeClientVpnEndpointsInput{}
        output := &ec2.DescribeClientVpnEndpointsOutput{}

        mockClient.On("DescribeClientVpnEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientVpnEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientVpnRoutes", func(t *testing.T) {
        input := &ec2.DescribeClientVpnRoutesInput{}
        output := &ec2.DescribeClientVpnRoutesOutput{}

        mockClient.On("DescribeClientVpnRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientVpnRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientVpnTargetNetworks", func(t *testing.T) {
        input := &ec2.DescribeClientVpnTargetNetworksInput{}
        output := &ec2.DescribeClientVpnTargetNetworksOutput{}

        mockClient.On("DescribeClientVpnTargetNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientVpnTargetNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCoipPools", func(t *testing.T) {
        input := &ec2.DescribeCoipPoolsInput{}
        output := &ec2.DescribeCoipPoolsOutput{}

        mockClient.On("DescribeCoipPools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCoipPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConversionTasks", func(t *testing.T) {
        input := &ec2.DescribeConversionTasksInput{}
        output := &ec2.DescribeConversionTasksOutput{}

        mockClient.On("DescribeConversionTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConversionTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomerGateways", func(t *testing.T) {
        input := &ec2.DescribeCustomerGatewaysInput{}
        output := &ec2.DescribeCustomerGatewaysOutput{}

        mockClient.On("DescribeCustomerGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomerGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDhcpOptions", func(t *testing.T) {
        input := &ec2.DescribeDhcpOptionsInput{}
        output := &ec2.DescribeDhcpOptionsOutput{}

        mockClient.On("DescribeDhcpOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDhcpOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEgressOnlyInternetGateways", func(t *testing.T) {
        input := &ec2.DescribeEgressOnlyInternetGatewaysInput{}
        output := &ec2.DescribeEgressOnlyInternetGatewaysOutput{}

        mockClient.On("DescribeEgressOnlyInternetGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEgressOnlyInternetGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticGpus", func(t *testing.T) {
        input := &ec2.DescribeElasticGpusInput{}
        output := &ec2.DescribeElasticGpusOutput{}

        mockClient.On("DescribeElasticGpus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticGpus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportImageTasks", func(t *testing.T) {
        input := &ec2.DescribeExportImageTasksInput{}
        output := &ec2.DescribeExportImageTasksOutput{}

        mockClient.On("DescribeExportImageTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportImageTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportTasks", func(t *testing.T) {
        input := &ec2.DescribeExportTasksInput{}
        output := &ec2.DescribeExportTasksOutput{}

        mockClient.On("DescribeExportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFastLaunchImages", func(t *testing.T) {
        input := &ec2.DescribeFastLaunchImagesInput{}
        output := &ec2.DescribeFastLaunchImagesOutput{}

        mockClient.On("DescribeFastLaunchImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFastLaunchImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFastSnapshotRestores", func(t *testing.T) {
        input := &ec2.DescribeFastSnapshotRestoresInput{}
        output := &ec2.DescribeFastSnapshotRestoresOutput{}

        mockClient.On("DescribeFastSnapshotRestores", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFastSnapshotRestores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetHistory", func(t *testing.T) {
        input := &ec2.DescribeFleetHistoryInput{}
        output := &ec2.DescribeFleetHistoryOutput{}

        mockClient.On("DescribeFleetHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetInstances", func(t *testing.T) {
        input := &ec2.DescribeFleetInstancesInput{}
        output := &ec2.DescribeFleetInstancesOutput{}

        mockClient.On("DescribeFleetInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleets", func(t *testing.T) {
        input := &ec2.DescribeFleetsInput{}
        output := &ec2.DescribeFleetsOutput{}

        mockClient.On("DescribeFleets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlowLogs", func(t *testing.T) {
        input := &ec2.DescribeFlowLogsInput{}
        output := &ec2.DescribeFlowLogsOutput{}

        mockClient.On("DescribeFlowLogs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlowLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFpgaImageAttribute", func(t *testing.T) {
        input := &ec2.DescribeFpgaImageAttributeInput{}
        output := &ec2.DescribeFpgaImageAttributeOutput{}

        mockClient.On("DescribeFpgaImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFpgaImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFpgaImages", func(t *testing.T) {
        input := &ec2.DescribeFpgaImagesInput{}
        output := &ec2.DescribeFpgaImagesOutput{}

        mockClient.On("DescribeFpgaImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFpgaImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHostReservationOfferings", func(t *testing.T) {
        input := &ec2.DescribeHostReservationOfferingsInput{}
        output := &ec2.DescribeHostReservationOfferingsOutput{}

        mockClient.On("DescribeHostReservationOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHostReservationOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHostReservations", func(t *testing.T) {
        input := &ec2.DescribeHostReservationsInput{}
        output := &ec2.DescribeHostReservationsOutput{}

        mockClient.On("DescribeHostReservations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHostReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHosts", func(t *testing.T) {
        input := &ec2.DescribeHostsInput{}
        output := &ec2.DescribeHostsOutput{}

        mockClient.On("DescribeHosts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIamInstanceProfileAssociations", func(t *testing.T) {
        input := &ec2.DescribeIamInstanceProfileAssociationsInput{}
        output := &ec2.DescribeIamInstanceProfileAssociationsOutput{}

        mockClient.On("DescribeIamInstanceProfileAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIamInstanceProfileAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdFormat", func(t *testing.T) {
        input := &ec2.DescribeIdFormatInput{}
        output := &ec2.DescribeIdFormatOutput{}

        mockClient.On("DescribeIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityIdFormat", func(t *testing.T) {
        input := &ec2.DescribeIdentityIdFormatInput{}
        output := &ec2.DescribeIdentityIdFormatOutput{}

        mockClient.On("DescribeIdentityIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageAttribute", func(t *testing.T) {
        input := &ec2.DescribeImageAttributeInput{}
        output := &ec2.DescribeImageAttributeOutput{}

        mockClient.On("DescribeImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImages", func(t *testing.T) {
        input := &ec2.DescribeImagesInput{}
        output := &ec2.DescribeImagesOutput{}

        mockClient.On("DescribeImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImportImageTasks", func(t *testing.T) {
        input := &ec2.DescribeImportImageTasksInput{}
        output := &ec2.DescribeImportImageTasksOutput{}

        mockClient.On("DescribeImportImageTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImportImageTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImportSnapshotTasks", func(t *testing.T) {
        input := &ec2.DescribeImportSnapshotTasksInput{}
        output := &ec2.DescribeImportSnapshotTasksOutput{}

        mockClient.On("DescribeImportSnapshotTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImportSnapshotTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceAttribute", func(t *testing.T) {
        input := &ec2.DescribeInstanceAttributeInput{}
        output := &ec2.DescribeInstanceAttributeOutput{}

        mockClient.On("DescribeInstanceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceConnectEndpoints", func(t *testing.T) {
        input := &ec2.DescribeInstanceConnectEndpointsInput{}
        output := &ec2.DescribeInstanceConnectEndpointsOutput{}

        mockClient.On("DescribeInstanceConnectEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceConnectEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceCreditSpecifications", func(t *testing.T) {
        input := &ec2.DescribeInstanceCreditSpecificationsInput{}
        output := &ec2.DescribeInstanceCreditSpecificationsOutput{}

        mockClient.On("DescribeInstanceCreditSpecifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceCreditSpecifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceEventNotificationAttributes", func(t *testing.T) {
        input := &ec2.DescribeInstanceEventNotificationAttributesInput{}
        output := &ec2.DescribeInstanceEventNotificationAttributesOutput{}

        mockClient.On("DescribeInstanceEventNotificationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceEventNotificationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceEventWindows", func(t *testing.T) {
        input := &ec2.DescribeInstanceEventWindowsInput{}
        output := &ec2.DescribeInstanceEventWindowsOutput{}

        mockClient.On("DescribeInstanceEventWindows", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceEventWindows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceStatus", func(t *testing.T) {
        input := &ec2.DescribeInstanceStatusInput{}
        output := &ec2.DescribeInstanceStatusOutput{}

        mockClient.On("DescribeInstanceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceTopology", func(t *testing.T) {
        input := &ec2.DescribeInstanceTopologyInput{}
        output := &ec2.DescribeInstanceTopologyOutput{}

        mockClient.On("DescribeInstanceTopology", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceTopology(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceTypeOfferings", func(t *testing.T) {
        input := &ec2.DescribeInstanceTypeOfferingsInput{}
        output := &ec2.DescribeInstanceTypeOfferingsOutput{}

        mockClient.On("DescribeInstanceTypeOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceTypeOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceTypes", func(t *testing.T) {
        input := &ec2.DescribeInstanceTypesInput{}
        output := &ec2.DescribeInstanceTypesOutput{}

        mockClient.On("DescribeInstanceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstances", func(t *testing.T) {
        input := &ec2.DescribeInstancesInput{}
        output := &ec2.DescribeInstancesOutput{}

        mockClient.On("DescribeInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInternetGateways", func(t *testing.T) {
        input := &ec2.DescribeInternetGatewaysInput{}
        output := &ec2.DescribeInternetGatewaysOutput{}

        mockClient.On("DescribeInternetGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInternetGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpamByoasn", func(t *testing.T) {
        input := &ec2.DescribeIpamByoasnInput{}
        output := &ec2.DescribeIpamByoasnOutput{}

        mockClient.On("DescribeIpamByoasn", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpamByoasn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpamPools", func(t *testing.T) {
        input := &ec2.DescribeIpamPoolsInput{}
        output := &ec2.DescribeIpamPoolsOutput{}

        mockClient.On("DescribeIpamPools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpamPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpamResourceDiscoveries", func(t *testing.T) {
        input := &ec2.DescribeIpamResourceDiscoveriesInput{}
        output := &ec2.DescribeIpamResourceDiscoveriesOutput{}

        mockClient.On("DescribeIpamResourceDiscoveries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpamResourceDiscoveries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpamResourceDiscoveryAssociations", func(t *testing.T) {
        input := &ec2.DescribeIpamResourceDiscoveryAssociationsInput{}
        output := &ec2.DescribeIpamResourceDiscoveryAssociationsOutput{}

        mockClient.On("DescribeIpamResourceDiscoveryAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpamResourceDiscoveryAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpamScopes", func(t *testing.T) {
        input := &ec2.DescribeIpamScopesInput{}
        output := &ec2.DescribeIpamScopesOutput{}

        mockClient.On("DescribeIpamScopes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpamScopes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpams", func(t *testing.T) {
        input := &ec2.DescribeIpamsInput{}
        output := &ec2.DescribeIpamsOutput{}

        mockClient.On("DescribeIpams", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpv6Pools", func(t *testing.T) {
        input := &ec2.DescribeIpv6PoolsInput{}
        output := &ec2.DescribeIpv6PoolsOutput{}

        mockClient.On("DescribeIpv6Pools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpv6Pools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeyPairs", func(t *testing.T) {
        input := &ec2.DescribeKeyPairsInput{}
        output := &ec2.DescribeKeyPairsOutput{}

        mockClient.On("DescribeKeyPairs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeyPairs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLaunchTemplateVersions", func(t *testing.T) {
        input := &ec2.DescribeLaunchTemplateVersionsInput{}
        output := &ec2.DescribeLaunchTemplateVersionsOutput{}

        mockClient.On("DescribeLaunchTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLaunchTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLaunchTemplates", func(t *testing.T) {
        input := &ec2.DescribeLaunchTemplatesInput{}
        output := &ec2.DescribeLaunchTemplatesOutput{}

        mockClient.On("DescribeLaunchTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLaunchTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
        output := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput{}

        mockClient.On("DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGatewayRouteTableVpcAssociations", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput{}
        output := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput{}

        mockClient.On("DescribeLocalGatewayRouteTableVpcAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGatewayRouteTableVpcAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGatewayRouteTables", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewayRouteTablesInput{}
        output := &ec2.DescribeLocalGatewayRouteTablesOutput{}

        mockClient.On("DescribeLocalGatewayRouteTables", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGatewayRouteTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGatewayVirtualInterfaceGroups", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput{}
        output := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput{}

        mockClient.On("DescribeLocalGatewayVirtualInterfaceGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGatewayVirtualInterfaceGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGatewayVirtualInterfaces", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewayVirtualInterfacesInput{}
        output := &ec2.DescribeLocalGatewayVirtualInterfacesOutput{}

        mockClient.On("DescribeLocalGatewayVirtualInterfaces", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGatewayVirtualInterfaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocalGateways", func(t *testing.T) {
        input := &ec2.DescribeLocalGatewaysInput{}
        output := &ec2.DescribeLocalGatewaysOutput{}

        mockClient.On("DescribeLocalGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocalGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLockedSnapshots", func(t *testing.T) {
        input := &ec2.DescribeLockedSnapshotsInput{}
        output := &ec2.DescribeLockedSnapshotsOutput{}

        mockClient.On("DescribeLockedSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLockedSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMacHosts", func(t *testing.T) {
        input := &ec2.DescribeMacHostsInput{}
        output := &ec2.DescribeMacHostsOutput{}

        mockClient.On("DescribeMacHosts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMacHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeManagedPrefixLists", func(t *testing.T) {
        input := &ec2.DescribeManagedPrefixListsInput{}
        output := &ec2.DescribeManagedPrefixListsOutput{}

        mockClient.On("DescribeManagedPrefixLists", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeManagedPrefixLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMovingAddresses", func(t *testing.T) {
        input := &ec2.DescribeMovingAddressesInput{}
        output := &ec2.DescribeMovingAddressesOutput{}

        mockClient.On("DescribeMovingAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMovingAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNatGateways", func(t *testing.T) {
        input := &ec2.DescribeNatGatewaysInput{}
        output := &ec2.DescribeNatGatewaysOutput{}

        mockClient.On("DescribeNatGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNatGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkAcls", func(t *testing.T) {
        input := &ec2.DescribeNetworkAclsInput{}
        output := &ec2.DescribeNetworkAclsOutput{}

        mockClient.On("DescribeNetworkAcls", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkAcls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInsightsAccessScopeAnalyses", func(t *testing.T) {
        input := &ec2.DescribeNetworkInsightsAccessScopeAnalysesInput{}
        output := &ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput{}

        mockClient.On("DescribeNetworkInsightsAccessScopeAnalyses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInsightsAccessScopeAnalyses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInsightsAccessScopes", func(t *testing.T) {
        input := &ec2.DescribeNetworkInsightsAccessScopesInput{}
        output := &ec2.DescribeNetworkInsightsAccessScopesOutput{}

        mockClient.On("DescribeNetworkInsightsAccessScopes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInsightsAccessScopes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInsightsAnalyses", func(t *testing.T) {
        input := &ec2.DescribeNetworkInsightsAnalysesInput{}
        output := &ec2.DescribeNetworkInsightsAnalysesOutput{}

        mockClient.On("DescribeNetworkInsightsAnalyses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInsightsAnalyses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInsightsPaths", func(t *testing.T) {
        input := &ec2.DescribeNetworkInsightsPathsInput{}
        output := &ec2.DescribeNetworkInsightsPathsOutput{}

        mockClient.On("DescribeNetworkInsightsPaths", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInsightsPaths(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInterfaceAttribute", func(t *testing.T) {
        input := &ec2.DescribeNetworkInterfaceAttributeInput{}
        output := &ec2.DescribeNetworkInterfaceAttributeOutput{}

        mockClient.On("DescribeNetworkInterfaceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInterfaceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInterfacePermissions", func(t *testing.T) {
        input := &ec2.DescribeNetworkInterfacePermissionsInput{}
        output := &ec2.DescribeNetworkInterfacePermissionsOutput{}

        mockClient.On("DescribeNetworkInterfacePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInterfacePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNetworkInterfaces", func(t *testing.T) {
        input := &ec2.DescribeNetworkInterfacesInput{}
        output := &ec2.DescribeNetworkInterfacesOutput{}

        mockClient.On("DescribeNetworkInterfaces", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNetworkInterfaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePlacementGroups", func(t *testing.T) {
        input := &ec2.DescribePlacementGroupsInput{}
        output := &ec2.DescribePlacementGroupsOutput{}

        mockClient.On("DescribePlacementGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePlacementGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePrefixLists", func(t *testing.T) {
        input := &ec2.DescribePrefixListsInput{}
        output := &ec2.DescribePrefixListsOutput{}

        mockClient.On("DescribePrefixLists", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePrefixLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePrincipalIdFormat", func(t *testing.T) {
        input := &ec2.DescribePrincipalIdFormatInput{}
        output := &ec2.DescribePrincipalIdFormatOutput{}

        mockClient.On("DescribePrincipalIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePrincipalIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePublicIpv4Pools", func(t *testing.T) {
        input := &ec2.DescribePublicIpv4PoolsInput{}
        output := &ec2.DescribePublicIpv4PoolsOutput{}

        mockClient.On("DescribePublicIpv4Pools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePublicIpv4Pools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegions", func(t *testing.T) {
        input := &ec2.DescribeRegionsInput{}
        output := &ec2.DescribeRegionsOutput{}

        mockClient.On("DescribeRegions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplaceRootVolumeTasks", func(t *testing.T) {
        input := &ec2.DescribeReplaceRootVolumeTasksInput{}
        output := &ec2.DescribeReplaceRootVolumeTasksOutput{}

        mockClient.On("DescribeReplaceRootVolumeTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplaceRootVolumeTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstances", func(t *testing.T) {
        input := &ec2.DescribeReservedInstancesInput{}
        output := &ec2.DescribeReservedInstancesOutput{}

        mockClient.On("DescribeReservedInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstancesListings", func(t *testing.T) {
        input := &ec2.DescribeReservedInstancesListingsInput{}
        output := &ec2.DescribeReservedInstancesListingsOutput{}

        mockClient.On("DescribeReservedInstancesListings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstancesListings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstancesModifications", func(t *testing.T) {
        input := &ec2.DescribeReservedInstancesModificationsInput{}
        output := &ec2.DescribeReservedInstancesModificationsOutput{}

        mockClient.On("DescribeReservedInstancesModifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstancesModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstancesOfferings", func(t *testing.T) {
        input := &ec2.DescribeReservedInstancesOfferingsInput{}
        output := &ec2.DescribeReservedInstancesOfferingsOutput{}

        mockClient.On("DescribeReservedInstancesOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstancesOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRouteTables", func(t *testing.T) {
        input := &ec2.DescribeRouteTablesInput{}
        output := &ec2.DescribeRouteTablesOutput{}

        mockClient.On("DescribeRouteTables", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRouteTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledInstanceAvailability", func(t *testing.T) {
        input := &ec2.DescribeScheduledInstanceAvailabilityInput{}
        output := &ec2.DescribeScheduledInstanceAvailabilityOutput{}

        mockClient.On("DescribeScheduledInstanceAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledInstanceAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledInstances", func(t *testing.T) {
        input := &ec2.DescribeScheduledInstancesInput{}
        output := &ec2.DescribeScheduledInstancesOutput{}

        mockClient.On("DescribeScheduledInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityGroupReferences", func(t *testing.T) {
        input := &ec2.DescribeSecurityGroupReferencesInput{}
        output := &ec2.DescribeSecurityGroupReferencesOutput{}

        mockClient.On("DescribeSecurityGroupReferences", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityGroupReferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityGroupRules", func(t *testing.T) {
        input := &ec2.DescribeSecurityGroupRulesInput{}
        output := &ec2.DescribeSecurityGroupRulesOutput{}

        mockClient.On("DescribeSecurityGroupRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityGroupRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityGroups", func(t *testing.T) {
        input := &ec2.DescribeSecurityGroupsInput{}
        output := &ec2.DescribeSecurityGroupsOutput{}

        mockClient.On("DescribeSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshotAttribute", func(t *testing.T) {
        input := &ec2.DescribeSnapshotAttributeInput{}
        output := &ec2.DescribeSnapshotAttributeOutput{}

        mockClient.On("DescribeSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshotTierStatus", func(t *testing.T) {
        input := &ec2.DescribeSnapshotTierStatusInput{}
        output := &ec2.DescribeSnapshotTierStatusOutput{}

        mockClient.On("DescribeSnapshotTierStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshotTierStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshots", func(t *testing.T) {
        input := &ec2.DescribeSnapshotsInput{}
        output := &ec2.DescribeSnapshotsOutput{}

        mockClient.On("DescribeSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotDatafeedSubscription", func(t *testing.T) {
        input := &ec2.DescribeSpotDatafeedSubscriptionInput{}
        output := &ec2.DescribeSpotDatafeedSubscriptionOutput{}

        mockClient.On("DescribeSpotDatafeedSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotDatafeedSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotFleetInstances", func(t *testing.T) {
        input := &ec2.DescribeSpotFleetInstancesInput{}
        output := &ec2.DescribeSpotFleetInstancesOutput{}

        mockClient.On("DescribeSpotFleetInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotFleetInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotFleetRequestHistory", func(t *testing.T) {
        input := &ec2.DescribeSpotFleetRequestHistoryInput{}
        output := &ec2.DescribeSpotFleetRequestHistoryOutput{}

        mockClient.On("DescribeSpotFleetRequestHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotFleetRequestHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotFleetRequests", func(t *testing.T) {
        input := &ec2.DescribeSpotFleetRequestsInput{}
        output := &ec2.DescribeSpotFleetRequestsOutput{}

        mockClient.On("DescribeSpotFleetRequests", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotFleetRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotInstanceRequests", func(t *testing.T) {
        input := &ec2.DescribeSpotInstanceRequestsInput{}
        output := &ec2.DescribeSpotInstanceRequestsOutput{}

        mockClient.On("DescribeSpotInstanceRequests", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotInstanceRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpotPriceHistory", func(t *testing.T) {
        input := &ec2.DescribeSpotPriceHistoryInput{}
        output := &ec2.DescribeSpotPriceHistoryOutput{}

        mockClient.On("DescribeSpotPriceHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpotPriceHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStaleSecurityGroups", func(t *testing.T) {
        input := &ec2.DescribeStaleSecurityGroupsInput{}
        output := &ec2.DescribeStaleSecurityGroupsOutput{}

        mockClient.On("DescribeStaleSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStaleSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStoreImageTasks", func(t *testing.T) {
        input := &ec2.DescribeStoreImageTasksInput{}
        output := &ec2.DescribeStoreImageTasksOutput{}

        mockClient.On("DescribeStoreImageTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStoreImageTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubnets", func(t *testing.T) {
        input := &ec2.DescribeSubnetsInput{}
        output := &ec2.DescribeSubnetsOutput{}

        mockClient.On("DescribeSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &ec2.DescribeTagsInput{}
        output := &ec2.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficMirrorFilterRules", func(t *testing.T) {
        input := &ec2.DescribeTrafficMirrorFilterRulesInput{}
        output := &ec2.DescribeTrafficMirrorFilterRulesOutput{}

        mockClient.On("DescribeTrafficMirrorFilterRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficMirrorFilterRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficMirrorFilters", func(t *testing.T) {
        input := &ec2.DescribeTrafficMirrorFiltersInput{}
        output := &ec2.DescribeTrafficMirrorFiltersOutput{}

        mockClient.On("DescribeTrafficMirrorFilters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficMirrorFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficMirrorSessions", func(t *testing.T) {
        input := &ec2.DescribeTrafficMirrorSessionsInput{}
        output := &ec2.DescribeTrafficMirrorSessionsOutput{}

        mockClient.On("DescribeTrafficMirrorSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficMirrorSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficMirrorTargets", func(t *testing.T) {
        input := &ec2.DescribeTrafficMirrorTargetsInput{}
        output := &ec2.DescribeTrafficMirrorTargetsOutput{}

        mockClient.On("DescribeTrafficMirrorTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficMirrorTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayAttachments", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayAttachmentsInput{}
        output := &ec2.DescribeTransitGatewayAttachmentsOutput{}

        mockClient.On("DescribeTransitGatewayAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayConnectPeers", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayConnectPeersInput{}
        output := &ec2.DescribeTransitGatewayConnectPeersOutput{}

        mockClient.On("DescribeTransitGatewayConnectPeers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayConnectPeers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayConnects", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayConnectsInput{}
        output := &ec2.DescribeTransitGatewayConnectsOutput{}

        mockClient.On("DescribeTransitGatewayConnects", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayMulticastDomains", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayMulticastDomainsInput{}
        output := &ec2.DescribeTransitGatewayMulticastDomainsOutput{}

        mockClient.On("DescribeTransitGatewayMulticastDomains", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayMulticastDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayPeeringAttachments", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayPeeringAttachmentsInput{}
        output := &ec2.DescribeTransitGatewayPeeringAttachmentsOutput{}

        mockClient.On("DescribeTransitGatewayPeeringAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayPeeringAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayPolicyTables", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayPolicyTablesInput{}
        output := &ec2.DescribeTransitGatewayPolicyTablesOutput{}

        mockClient.On("DescribeTransitGatewayPolicyTables", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayPolicyTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayRouteTableAnnouncements", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayRouteTableAnnouncementsInput{}
        output := &ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput{}

        mockClient.On("DescribeTransitGatewayRouteTableAnnouncements", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayRouteTableAnnouncements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayRouteTables", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayRouteTablesInput{}
        output := &ec2.DescribeTransitGatewayRouteTablesOutput{}

        mockClient.On("DescribeTransitGatewayRouteTables", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayRouteTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGatewayVpcAttachments", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewayVpcAttachmentsInput{}
        output := &ec2.DescribeTransitGatewayVpcAttachmentsOutput{}

        mockClient.On("DescribeTransitGatewayVpcAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGatewayVpcAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransitGateways", func(t *testing.T) {
        input := &ec2.DescribeTransitGatewaysInput{}
        output := &ec2.DescribeTransitGatewaysOutput{}

        mockClient.On("DescribeTransitGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransitGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrunkInterfaceAssociations", func(t *testing.T) {
        input := &ec2.DescribeTrunkInterfaceAssociationsInput{}
        output := &ec2.DescribeTrunkInterfaceAssociationsOutput{}

        mockClient.On("DescribeTrunkInterfaceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrunkInterfaceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedAccessEndpoints", func(t *testing.T) {
        input := &ec2.DescribeVerifiedAccessEndpointsInput{}
        output := &ec2.DescribeVerifiedAccessEndpointsOutput{}

        mockClient.On("DescribeVerifiedAccessEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedAccessEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedAccessGroups", func(t *testing.T) {
        input := &ec2.DescribeVerifiedAccessGroupsInput{}
        output := &ec2.DescribeVerifiedAccessGroupsOutput{}

        mockClient.On("DescribeVerifiedAccessGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedAccessGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedAccessInstanceLoggingConfigurations", func(t *testing.T) {
        input := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput{}
        output := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput{}

        mockClient.On("DescribeVerifiedAccessInstanceLoggingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedAccessInstanceLoggingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedAccessInstances", func(t *testing.T) {
        input := &ec2.DescribeVerifiedAccessInstancesInput{}
        output := &ec2.DescribeVerifiedAccessInstancesOutput{}

        mockClient.On("DescribeVerifiedAccessInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedAccessInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedAccessTrustProviders", func(t *testing.T) {
        input := &ec2.DescribeVerifiedAccessTrustProvidersInput{}
        output := &ec2.DescribeVerifiedAccessTrustProvidersOutput{}

        mockClient.On("DescribeVerifiedAccessTrustProviders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedAccessTrustProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumeAttribute", func(t *testing.T) {
        input := &ec2.DescribeVolumeAttributeInput{}
        output := &ec2.DescribeVolumeAttributeOutput{}

        mockClient.On("DescribeVolumeAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumeAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumeStatus", func(t *testing.T) {
        input := &ec2.DescribeVolumeStatusInput{}
        output := &ec2.DescribeVolumeStatusOutput{}

        mockClient.On("DescribeVolumeStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumeStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumes", func(t *testing.T) {
        input := &ec2.DescribeVolumesInput{}
        output := &ec2.DescribeVolumesOutput{}

        mockClient.On("DescribeVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumesModifications", func(t *testing.T) {
        input := &ec2.DescribeVolumesModificationsInput{}
        output := &ec2.DescribeVolumesModificationsOutput{}

        mockClient.On("DescribeVolumesModifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumesModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcAttribute", func(t *testing.T) {
        input := &ec2.DescribeVpcAttributeInput{}
        output := &ec2.DescribeVpcAttributeOutput{}

        mockClient.On("DescribeVpcAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcClassicLink", func(t *testing.T) {
        input := &ec2.DescribeVpcClassicLinkInput{}
        output := &ec2.DescribeVpcClassicLinkOutput{}

        mockClient.On("DescribeVpcClassicLink", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcClassicLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcClassicLinkDnsSupport", func(t *testing.T) {
        input := &ec2.DescribeVpcClassicLinkDnsSupportInput{}
        output := &ec2.DescribeVpcClassicLinkDnsSupportOutput{}

        mockClient.On("DescribeVpcClassicLinkDnsSupport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcClassicLinkDnsSupport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpointConnectionNotifications", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointConnectionNotificationsInput{}
        output := &ec2.DescribeVpcEndpointConnectionNotificationsOutput{}

        mockClient.On("DescribeVpcEndpointConnectionNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpointConnectionNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpointConnections", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointConnectionsInput{}
        output := &ec2.DescribeVpcEndpointConnectionsOutput{}

        mockClient.On("DescribeVpcEndpointConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpointConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpointServiceConfigurations", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointServiceConfigurationsInput{}
        output := &ec2.DescribeVpcEndpointServiceConfigurationsOutput{}

        mockClient.On("DescribeVpcEndpointServiceConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpointServiceConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpointServicePermissions", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointServicePermissionsInput{}
        output := &ec2.DescribeVpcEndpointServicePermissionsOutput{}

        mockClient.On("DescribeVpcEndpointServicePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpointServicePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpointServices", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointServicesInput{}
        output := &ec2.DescribeVpcEndpointServicesOutput{}

        mockClient.On("DescribeVpcEndpointServices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpointServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpoints", func(t *testing.T) {
        input := &ec2.DescribeVpcEndpointsInput{}
        output := &ec2.DescribeVpcEndpointsOutput{}

        mockClient.On("DescribeVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcPeeringConnections", func(t *testing.T) {
        input := &ec2.DescribeVpcPeeringConnectionsInput{}
        output := &ec2.DescribeVpcPeeringConnectionsOutput{}

        mockClient.On("DescribeVpcPeeringConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcPeeringConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcs", func(t *testing.T) {
        input := &ec2.DescribeVpcsInput{}
        output := &ec2.DescribeVpcsOutput{}

        mockClient.On("DescribeVpcs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpnConnections", func(t *testing.T) {
        input := &ec2.DescribeVpnConnectionsInput{}
        output := &ec2.DescribeVpnConnectionsOutput{}

        mockClient.On("DescribeVpnConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpnConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpnGateways", func(t *testing.T) {
        input := &ec2.DescribeVpnGatewaysInput{}
        output := &ec2.DescribeVpnGatewaysOutput{}

        mockClient.On("DescribeVpnGateways", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpnGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachClassicLinkVpc", func(t *testing.T) {
        input := &ec2.DetachClassicLinkVpcInput{}
        output := &ec2.DetachClassicLinkVpcOutput{}

        mockClient.On("DetachClassicLinkVpc", ctx, input).Return(output, nil)

        result, err := mockClient.DetachClassicLinkVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachInternetGateway", func(t *testing.T) {
        input := &ec2.DetachInternetGatewayInput{}
        output := &ec2.DetachInternetGatewayOutput{}

        mockClient.On("DetachInternetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DetachInternetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachNetworkInterface", func(t *testing.T) {
        input := &ec2.DetachNetworkInterfaceInput{}
        output := &ec2.DetachNetworkInterfaceOutput{}

        mockClient.On("DetachNetworkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.DetachNetworkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachVerifiedAccessTrustProvider", func(t *testing.T) {
        input := &ec2.DetachVerifiedAccessTrustProviderInput{}
        output := &ec2.DetachVerifiedAccessTrustProviderOutput{}

        mockClient.On("DetachVerifiedAccessTrustProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DetachVerifiedAccessTrustProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachVolume", func(t *testing.T) {
        input := &ec2.DetachVolumeInput{}
        output := &ec2.DetachVolumeOutput{}

        mockClient.On("DetachVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DetachVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachVpnGateway", func(t *testing.T) {
        input := &ec2.DetachVpnGatewayInput{}
        output := &ec2.DetachVpnGatewayOutput{}

        mockClient.On("DetachVpnGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DetachVpnGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAddressTransfer", func(t *testing.T) {
        input := &ec2.DisableAddressTransferInput{}
        output := &ec2.DisableAddressTransferOutput{}

        mockClient.On("DisableAddressTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAddressTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAwsNetworkPerformanceMetricSubscription", func(t *testing.T) {
        input := &ec2.DisableAwsNetworkPerformanceMetricSubscriptionInput{}
        output := &ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput{}

        mockClient.On("DisableAwsNetworkPerformanceMetricSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAwsNetworkPerformanceMetricSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableEbsEncryptionByDefault", func(t *testing.T) {
        input := &ec2.DisableEbsEncryptionByDefaultInput{}
        output := &ec2.DisableEbsEncryptionByDefaultOutput{}

        mockClient.On("DisableEbsEncryptionByDefault", ctx, input).Return(output, nil)

        result, err := mockClient.DisableEbsEncryptionByDefault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableFastLaunch", func(t *testing.T) {
        input := &ec2.DisableFastLaunchInput{}
        output := &ec2.DisableFastLaunchOutput{}

        mockClient.On("DisableFastLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.DisableFastLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableFastSnapshotRestores", func(t *testing.T) {
        input := &ec2.DisableFastSnapshotRestoresInput{}
        output := &ec2.DisableFastSnapshotRestoresOutput{}

        mockClient.On("DisableFastSnapshotRestores", ctx, input).Return(output, nil)

        result, err := mockClient.DisableFastSnapshotRestores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableImage", func(t *testing.T) {
        input := &ec2.DisableImageInput{}
        output := &ec2.DisableImageOutput{}

        mockClient.On("DisableImage", ctx, input).Return(output, nil)

        result, err := mockClient.DisableImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableImageBlockPublicAccess", func(t *testing.T) {
        input := &ec2.DisableImageBlockPublicAccessInput{}
        output := &ec2.DisableImageBlockPublicAccessOutput{}

        mockClient.On("DisableImageBlockPublicAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DisableImageBlockPublicAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableImageDeprecation", func(t *testing.T) {
        input := &ec2.DisableImageDeprecationInput{}
        output := &ec2.DisableImageDeprecationOutput{}

        mockClient.On("DisableImageDeprecation", ctx, input).Return(output, nil)

        result, err := mockClient.DisableImageDeprecation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableImageDeregistrationProtection", func(t *testing.T) {
        input := &ec2.DisableImageDeregistrationProtectionInput{}
        output := &ec2.DisableImageDeregistrationProtectionOutput{}

        mockClient.On("DisableImageDeregistrationProtection", ctx, input).Return(output, nil)

        result, err := mockClient.DisableImageDeregistrationProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableIpamOrganizationAdminAccount", func(t *testing.T) {
        input := &ec2.DisableIpamOrganizationAdminAccountInput{}
        output := &ec2.DisableIpamOrganizationAdminAccountOutput{}

        mockClient.On("DisableIpamOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableIpamOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSerialConsoleAccess", func(t *testing.T) {
        input := &ec2.DisableSerialConsoleAccessInput{}
        output := &ec2.DisableSerialConsoleAccessOutput{}

        mockClient.On("DisableSerialConsoleAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSerialConsoleAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSnapshotBlockPublicAccess", func(t *testing.T) {
        input := &ec2.DisableSnapshotBlockPublicAccessInput{}
        output := &ec2.DisableSnapshotBlockPublicAccessOutput{}

        mockClient.On("DisableSnapshotBlockPublicAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSnapshotBlockPublicAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableTransitGatewayRouteTablePropagation", func(t *testing.T) {
        input := &ec2.DisableTransitGatewayRouteTablePropagationInput{}
        output := &ec2.DisableTransitGatewayRouteTablePropagationOutput{}

        mockClient.On("DisableTransitGatewayRouteTablePropagation", ctx, input).Return(output, nil)

        result, err := mockClient.DisableTransitGatewayRouteTablePropagation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableVgwRoutePropagation", func(t *testing.T) {
        input := &ec2.DisableVgwRoutePropagationInput{}
        output := &ec2.DisableVgwRoutePropagationOutput{}

        mockClient.On("DisableVgwRoutePropagation", ctx, input).Return(output, nil)

        result, err := mockClient.DisableVgwRoutePropagation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableVpcClassicLink", func(t *testing.T) {
        input := &ec2.DisableVpcClassicLinkInput{}
        output := &ec2.DisableVpcClassicLinkOutput{}

        mockClient.On("DisableVpcClassicLink", ctx, input).Return(output, nil)

        result, err := mockClient.DisableVpcClassicLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableVpcClassicLinkDnsSupport", func(t *testing.T) {
        input := &ec2.DisableVpcClassicLinkDnsSupportInput{}
        output := &ec2.DisableVpcClassicLinkDnsSupportOutput{}

        mockClient.On("DisableVpcClassicLinkDnsSupport", ctx, input).Return(output, nil)

        result, err := mockClient.DisableVpcClassicLinkDnsSupport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAddress", func(t *testing.T) {
        input := &ec2.DisassociateAddressInput{}
        output := &ec2.DisassociateAddressOutput{}

        mockClient.On("DisassociateAddress", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateClientVpnTargetNetwork", func(t *testing.T) {
        input := &ec2.DisassociateClientVpnTargetNetworkInput{}
        output := &ec2.DisassociateClientVpnTargetNetworkOutput{}

        mockClient.On("DisassociateClientVpnTargetNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateClientVpnTargetNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateEnclaveCertificateIamRole", func(t *testing.T) {
        input := &ec2.DisassociateEnclaveCertificateIamRoleInput{}
        output := &ec2.DisassociateEnclaveCertificateIamRoleOutput{}

        mockClient.On("DisassociateEnclaveCertificateIamRole", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateEnclaveCertificateIamRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIamInstanceProfile", func(t *testing.T) {
        input := &ec2.DisassociateIamInstanceProfileInput{}
        output := &ec2.DisassociateIamInstanceProfileOutput{}

        mockClient.On("DisassociateIamInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIamInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateInstanceEventWindow", func(t *testing.T) {
        input := &ec2.DisassociateInstanceEventWindowInput{}
        output := &ec2.DisassociateInstanceEventWindowOutput{}

        mockClient.On("DisassociateInstanceEventWindow", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateInstanceEventWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIpamByoasn", func(t *testing.T) {
        input := &ec2.DisassociateIpamByoasnInput{}
        output := &ec2.DisassociateIpamByoasnOutput{}

        mockClient.On("DisassociateIpamByoasn", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIpamByoasn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateIpamResourceDiscovery", func(t *testing.T) {
        input := &ec2.DisassociateIpamResourceDiscoveryInput{}
        output := &ec2.DisassociateIpamResourceDiscoveryOutput{}

        mockClient.On("DisassociateIpamResourceDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateIpamResourceDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateNatGatewayAddress", func(t *testing.T) {
        input := &ec2.DisassociateNatGatewayAddressInput{}
        output := &ec2.DisassociateNatGatewayAddressOutput{}

        mockClient.On("DisassociateNatGatewayAddress", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateNatGatewayAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRouteTable", func(t *testing.T) {
        input := &ec2.DisassociateRouteTableInput{}
        output := &ec2.DisassociateRouteTableOutput{}

        mockClient.On("DisassociateRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSubnetCidrBlock", func(t *testing.T) {
        input := &ec2.DisassociateSubnetCidrBlockInput{}
        output := &ec2.DisassociateSubnetCidrBlockOutput{}

        mockClient.On("DisassociateSubnetCidrBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSubnetCidrBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTransitGatewayMulticastDomain", func(t *testing.T) {
        input := &ec2.DisassociateTransitGatewayMulticastDomainInput{}
        output := &ec2.DisassociateTransitGatewayMulticastDomainOutput{}

        mockClient.On("DisassociateTransitGatewayMulticastDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTransitGatewayMulticastDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTransitGatewayPolicyTable", func(t *testing.T) {
        input := &ec2.DisassociateTransitGatewayPolicyTableInput{}
        output := &ec2.DisassociateTransitGatewayPolicyTableOutput{}

        mockClient.On("DisassociateTransitGatewayPolicyTable", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTransitGatewayPolicyTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTransitGatewayRouteTable", func(t *testing.T) {
        input := &ec2.DisassociateTransitGatewayRouteTableInput{}
        output := &ec2.DisassociateTransitGatewayRouteTableOutput{}

        mockClient.On("DisassociateTransitGatewayRouteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTransitGatewayRouteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTrunkInterface", func(t *testing.T) {
        input := &ec2.DisassociateTrunkInterfaceInput{}
        output := &ec2.DisassociateTrunkInterfaceOutput{}

        mockClient.On("DisassociateTrunkInterface", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTrunkInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateVpcCidrBlock", func(t *testing.T) {
        input := &ec2.DisassociateVpcCidrBlockInput{}
        output := &ec2.DisassociateVpcCidrBlockOutput{}

        mockClient.On("DisassociateVpcCidrBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateVpcCidrBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAddressTransfer", func(t *testing.T) {
        input := &ec2.EnableAddressTransferInput{}
        output := &ec2.EnableAddressTransferOutput{}

        mockClient.On("EnableAddressTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAddressTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAwsNetworkPerformanceMetricSubscription", func(t *testing.T) {
        input := &ec2.EnableAwsNetworkPerformanceMetricSubscriptionInput{}
        output := &ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput{}

        mockClient.On("EnableAwsNetworkPerformanceMetricSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAwsNetworkPerformanceMetricSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableEbsEncryptionByDefault", func(t *testing.T) {
        input := &ec2.EnableEbsEncryptionByDefaultInput{}
        output := &ec2.EnableEbsEncryptionByDefaultOutput{}

        mockClient.On("EnableEbsEncryptionByDefault", ctx, input).Return(output, nil)

        result, err := mockClient.EnableEbsEncryptionByDefault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableFastLaunch", func(t *testing.T) {
        input := &ec2.EnableFastLaunchInput{}
        output := &ec2.EnableFastLaunchOutput{}

        mockClient.On("EnableFastLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.EnableFastLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableFastSnapshotRestores", func(t *testing.T) {
        input := &ec2.EnableFastSnapshotRestoresInput{}
        output := &ec2.EnableFastSnapshotRestoresOutput{}

        mockClient.On("EnableFastSnapshotRestores", ctx, input).Return(output, nil)

        result, err := mockClient.EnableFastSnapshotRestores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableImage", func(t *testing.T) {
        input := &ec2.EnableImageInput{}
        output := &ec2.EnableImageOutput{}

        mockClient.On("EnableImage", ctx, input).Return(output, nil)

        result, err := mockClient.EnableImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableImageBlockPublicAccess", func(t *testing.T) {
        input := &ec2.EnableImageBlockPublicAccessInput{}
        output := &ec2.EnableImageBlockPublicAccessOutput{}

        mockClient.On("EnableImageBlockPublicAccess", ctx, input).Return(output, nil)

        result, err := mockClient.EnableImageBlockPublicAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableImageDeprecation", func(t *testing.T) {
        input := &ec2.EnableImageDeprecationInput{}
        output := &ec2.EnableImageDeprecationOutput{}

        mockClient.On("EnableImageDeprecation", ctx, input).Return(output, nil)

        result, err := mockClient.EnableImageDeprecation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableImageDeregistrationProtection", func(t *testing.T) {
        input := &ec2.EnableImageDeregistrationProtectionInput{}
        output := &ec2.EnableImageDeregistrationProtectionOutput{}

        mockClient.On("EnableImageDeregistrationProtection", ctx, input).Return(output, nil)

        result, err := mockClient.EnableImageDeregistrationProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableIpamOrganizationAdminAccount", func(t *testing.T) {
        input := &ec2.EnableIpamOrganizationAdminAccountInput{}
        output := &ec2.EnableIpamOrganizationAdminAccountOutput{}

        mockClient.On("EnableIpamOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableIpamOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableReachabilityAnalyzerOrganizationSharing", func(t *testing.T) {
        input := &ec2.EnableReachabilityAnalyzerOrganizationSharingInput{}
        output := &ec2.EnableReachabilityAnalyzerOrganizationSharingOutput{}

        mockClient.On("EnableReachabilityAnalyzerOrganizationSharing", ctx, input).Return(output, nil)

        result, err := mockClient.EnableReachabilityAnalyzerOrganizationSharing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSerialConsoleAccess", func(t *testing.T) {
        input := &ec2.EnableSerialConsoleAccessInput{}
        output := &ec2.EnableSerialConsoleAccessOutput{}

        mockClient.On("EnableSerialConsoleAccess", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSerialConsoleAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSnapshotBlockPublicAccess", func(t *testing.T) {
        input := &ec2.EnableSnapshotBlockPublicAccessInput{}
        output := &ec2.EnableSnapshotBlockPublicAccessOutput{}

        mockClient.On("EnableSnapshotBlockPublicAccess", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSnapshotBlockPublicAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableTransitGatewayRouteTablePropagation", func(t *testing.T) {
        input := &ec2.EnableTransitGatewayRouteTablePropagationInput{}
        output := &ec2.EnableTransitGatewayRouteTablePropagationOutput{}

        mockClient.On("EnableTransitGatewayRouteTablePropagation", ctx, input).Return(output, nil)

        result, err := mockClient.EnableTransitGatewayRouteTablePropagation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableVgwRoutePropagation", func(t *testing.T) {
        input := &ec2.EnableVgwRoutePropagationInput{}
        output := &ec2.EnableVgwRoutePropagationOutput{}

        mockClient.On("EnableVgwRoutePropagation", ctx, input).Return(output, nil)

        result, err := mockClient.EnableVgwRoutePropagation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableVolumeIO", func(t *testing.T) {
        input := &ec2.EnableVolumeIOInput{}
        output := &ec2.EnableVolumeIOOutput{}

        mockClient.On("EnableVolumeIO", ctx, input).Return(output, nil)

        result, err := mockClient.EnableVolumeIO(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableVpcClassicLink", func(t *testing.T) {
        input := &ec2.EnableVpcClassicLinkInput{}
        output := &ec2.EnableVpcClassicLinkOutput{}

        mockClient.On("EnableVpcClassicLink", ctx, input).Return(output, nil)

        result, err := mockClient.EnableVpcClassicLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableVpcClassicLinkDnsSupport", func(t *testing.T) {
        input := &ec2.EnableVpcClassicLinkDnsSupportInput{}
        output := &ec2.EnableVpcClassicLinkDnsSupportOutput{}

        mockClient.On("EnableVpcClassicLinkDnsSupport", ctx, input).Return(output, nil)

        result, err := mockClient.EnableVpcClassicLinkDnsSupport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportClientVpnClientCertificateRevocationList", func(t *testing.T) {
        input := &ec2.ExportClientVpnClientCertificateRevocationListInput{}
        output := &ec2.ExportClientVpnClientCertificateRevocationListOutput{}

        mockClient.On("ExportClientVpnClientCertificateRevocationList", ctx, input).Return(output, nil)

        result, err := mockClient.ExportClientVpnClientCertificateRevocationList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportClientVpnClientConfiguration", func(t *testing.T) {
        input := &ec2.ExportClientVpnClientConfigurationInput{}
        output := &ec2.ExportClientVpnClientConfigurationOutput{}

        mockClient.On("ExportClientVpnClientConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ExportClientVpnClientConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportImage", func(t *testing.T) {
        input := &ec2.ExportImageInput{}
        output := &ec2.ExportImageOutput{}

        mockClient.On("ExportImage", ctx, input).Return(output, nil)

        result, err := mockClient.ExportImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportTransitGatewayRoutes", func(t *testing.T) {
        input := &ec2.ExportTransitGatewayRoutesInput{}
        output := &ec2.ExportTransitGatewayRoutesOutput{}

        mockClient.On("ExportTransitGatewayRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.ExportTransitGatewayRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociatedEnclaveCertificateIamRoles", func(t *testing.T) {
        input := &ec2.GetAssociatedEnclaveCertificateIamRolesInput{}
        output := &ec2.GetAssociatedEnclaveCertificateIamRolesOutput{}

        mockClient.On("GetAssociatedEnclaveCertificateIamRoles", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociatedEnclaveCertificateIamRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociatedIpv6PoolCidrs", func(t *testing.T) {
        input := &ec2.GetAssociatedIpv6PoolCidrsInput{}
        output := &ec2.GetAssociatedIpv6PoolCidrsOutput{}

        mockClient.On("GetAssociatedIpv6PoolCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociatedIpv6PoolCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAwsNetworkPerformanceData", func(t *testing.T) {
        input := &ec2.GetAwsNetworkPerformanceDataInput{}
        output := &ec2.GetAwsNetworkPerformanceDataOutput{}

        mockClient.On("GetAwsNetworkPerformanceData", ctx, input).Return(output, nil)

        result, err := mockClient.GetAwsNetworkPerformanceData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCapacityReservationUsage", func(t *testing.T) {
        input := &ec2.GetCapacityReservationUsageInput{}
        output := &ec2.GetCapacityReservationUsageOutput{}

        mockClient.On("GetCapacityReservationUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetCapacityReservationUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoipPoolUsage", func(t *testing.T) {
        input := &ec2.GetCoipPoolUsageInput{}
        output := &ec2.GetCoipPoolUsageOutput{}

        mockClient.On("GetCoipPoolUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoipPoolUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConsoleOutput", func(t *testing.T) {
        input := &ec2.GetConsoleOutputInput{}
        output := &ec2.GetConsoleOutputOutput{}

        mockClient.On("GetConsoleOutput", ctx, input).Return(output, nil)

        result, err := mockClient.GetConsoleOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConsoleScreenshot", func(t *testing.T) {
        input := &ec2.GetConsoleScreenshotInput{}
        output := &ec2.GetConsoleScreenshotOutput{}

        mockClient.On("GetConsoleScreenshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetConsoleScreenshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDefaultCreditSpecification", func(t *testing.T) {
        input := &ec2.GetDefaultCreditSpecificationInput{}
        output := &ec2.GetDefaultCreditSpecificationOutput{}

        mockClient.On("GetDefaultCreditSpecification", ctx, input).Return(output, nil)

        result, err := mockClient.GetDefaultCreditSpecification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEbsDefaultKmsKeyId", func(t *testing.T) {
        input := &ec2.GetEbsDefaultKmsKeyIdInput{}
        output := &ec2.GetEbsDefaultKmsKeyIdOutput{}

        mockClient.On("GetEbsDefaultKmsKeyId", ctx, input).Return(output, nil)

        result, err := mockClient.GetEbsDefaultKmsKeyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEbsEncryptionByDefault", func(t *testing.T) {
        input := &ec2.GetEbsEncryptionByDefaultInput{}
        output := &ec2.GetEbsEncryptionByDefaultOutput{}

        mockClient.On("GetEbsEncryptionByDefault", ctx, input).Return(output, nil)

        result, err := mockClient.GetEbsEncryptionByDefault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowLogsIntegrationTemplate", func(t *testing.T) {
        input := &ec2.GetFlowLogsIntegrationTemplateInput{}
        output := &ec2.GetFlowLogsIntegrationTemplateOutput{}

        mockClient.On("GetFlowLogsIntegrationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowLogsIntegrationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupsForCapacityReservation", func(t *testing.T) {
        input := &ec2.GetGroupsForCapacityReservationInput{}
        output := &ec2.GetGroupsForCapacityReservationOutput{}

        mockClient.On("GetGroupsForCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupsForCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostReservationPurchasePreview", func(t *testing.T) {
        input := &ec2.GetHostReservationPurchasePreviewInput{}
        output := &ec2.GetHostReservationPurchasePreviewOutput{}

        mockClient.On("GetHostReservationPurchasePreview", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostReservationPurchasePreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageBlockPublicAccessState", func(t *testing.T) {
        input := &ec2.GetImageBlockPublicAccessStateInput{}
        output := &ec2.GetImageBlockPublicAccessStateOutput{}

        mockClient.On("GetImageBlockPublicAccessState", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageBlockPublicAccessState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceMetadataDefaults", func(t *testing.T) {
        input := &ec2.GetInstanceMetadataDefaultsInput{}
        output := &ec2.GetInstanceMetadataDefaultsOutput{}

        mockClient.On("GetInstanceMetadataDefaults", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceMetadataDefaults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceTpmEkPub", func(t *testing.T) {
        input := &ec2.GetInstanceTpmEkPubInput{}
        output := &ec2.GetInstanceTpmEkPubOutput{}

        mockClient.On("GetInstanceTpmEkPub", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceTpmEkPub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceTypesFromInstanceRequirements", func(t *testing.T) {
        input := &ec2.GetInstanceTypesFromInstanceRequirementsInput{}
        output := &ec2.GetInstanceTypesFromInstanceRequirementsOutput{}

        mockClient.On("GetInstanceTypesFromInstanceRequirements", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceTypesFromInstanceRequirements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceUefiData", func(t *testing.T) {
        input := &ec2.GetInstanceUefiDataInput{}
        output := &ec2.GetInstanceUefiDataOutput{}

        mockClient.On("GetInstanceUefiData", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceUefiData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamAddressHistory", func(t *testing.T) {
        input := &ec2.GetIpamAddressHistoryInput{}
        output := &ec2.GetIpamAddressHistoryOutput{}

        mockClient.On("GetIpamAddressHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamAddressHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamDiscoveredAccounts", func(t *testing.T) {
        input := &ec2.GetIpamDiscoveredAccountsInput{}
        output := &ec2.GetIpamDiscoveredAccountsOutput{}

        mockClient.On("GetIpamDiscoveredAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamDiscoveredAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamDiscoveredPublicAddresses", func(t *testing.T) {
        input := &ec2.GetIpamDiscoveredPublicAddressesInput{}
        output := &ec2.GetIpamDiscoveredPublicAddressesOutput{}

        mockClient.On("GetIpamDiscoveredPublicAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamDiscoveredPublicAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamDiscoveredResourceCidrs", func(t *testing.T) {
        input := &ec2.GetIpamDiscoveredResourceCidrsInput{}
        output := &ec2.GetIpamDiscoveredResourceCidrsOutput{}

        mockClient.On("GetIpamDiscoveredResourceCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamDiscoveredResourceCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamPoolAllocations", func(t *testing.T) {
        input := &ec2.GetIpamPoolAllocationsInput{}
        output := &ec2.GetIpamPoolAllocationsOutput{}

        mockClient.On("GetIpamPoolAllocations", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamPoolAllocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamPoolCidrs", func(t *testing.T) {
        input := &ec2.GetIpamPoolCidrsInput{}
        output := &ec2.GetIpamPoolCidrsOutput{}

        mockClient.On("GetIpamPoolCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamPoolCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIpamResourceCidrs", func(t *testing.T) {
        input := &ec2.GetIpamResourceCidrsInput{}
        output := &ec2.GetIpamResourceCidrsOutput{}

        mockClient.On("GetIpamResourceCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.GetIpamResourceCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchTemplateData", func(t *testing.T) {
        input := &ec2.GetLaunchTemplateDataInput{}
        output := &ec2.GetLaunchTemplateDataOutput{}

        mockClient.On("GetLaunchTemplateData", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchTemplateData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedPrefixListAssociations", func(t *testing.T) {
        input := &ec2.GetManagedPrefixListAssociationsInput{}
        output := &ec2.GetManagedPrefixListAssociationsOutput{}

        mockClient.On("GetManagedPrefixListAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedPrefixListAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedPrefixListEntries", func(t *testing.T) {
        input := &ec2.GetManagedPrefixListEntriesInput{}
        output := &ec2.GetManagedPrefixListEntriesOutput{}

        mockClient.On("GetManagedPrefixListEntries", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedPrefixListEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkInsightsAccessScopeAnalysisFindings", func(t *testing.T) {
        input := &ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput{}
        output := &ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput{}

        mockClient.On("GetNetworkInsightsAccessScopeAnalysisFindings", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkInsightsAccessScopeAnalysisFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkInsightsAccessScopeContent", func(t *testing.T) {
        input := &ec2.GetNetworkInsightsAccessScopeContentInput{}
        output := &ec2.GetNetworkInsightsAccessScopeContentOutput{}

        mockClient.On("GetNetworkInsightsAccessScopeContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkInsightsAccessScopeContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPasswordData", func(t *testing.T) {
        input := &ec2.GetPasswordDataInput{}
        output := &ec2.GetPasswordDataOutput{}

        mockClient.On("GetPasswordData", ctx, input).Return(output, nil)

        result, err := mockClient.GetPasswordData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservedInstancesExchangeQuote", func(t *testing.T) {
        input := &ec2.GetReservedInstancesExchangeQuoteInput{}
        output := &ec2.GetReservedInstancesExchangeQuoteOutput{}

        mockClient.On("GetReservedInstancesExchangeQuote", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservedInstancesExchangeQuote(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityGroupsForVpc", func(t *testing.T) {
        input := &ec2.GetSecurityGroupsForVpcInput{}
        output := &ec2.GetSecurityGroupsForVpcOutput{}

        mockClient.On("GetSecurityGroupsForVpc", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityGroupsForVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSerialConsoleAccessStatus", func(t *testing.T) {
        input := &ec2.GetSerialConsoleAccessStatusInput{}
        output := &ec2.GetSerialConsoleAccessStatusOutput{}

        mockClient.On("GetSerialConsoleAccessStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetSerialConsoleAccessStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnapshotBlockPublicAccessState", func(t *testing.T) {
        input := &ec2.GetSnapshotBlockPublicAccessStateInput{}
        output := &ec2.GetSnapshotBlockPublicAccessStateOutput{}

        mockClient.On("GetSnapshotBlockPublicAccessState", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnapshotBlockPublicAccessState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpotPlacementScores", func(t *testing.T) {
        input := &ec2.GetSpotPlacementScoresInput{}
        output := &ec2.GetSpotPlacementScoresOutput{}

        mockClient.On("GetSpotPlacementScores", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpotPlacementScores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubnetCidrReservations", func(t *testing.T) {
        input := &ec2.GetSubnetCidrReservationsInput{}
        output := &ec2.GetSubnetCidrReservationsOutput{}

        mockClient.On("GetSubnetCidrReservations", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubnetCidrReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayAttachmentPropagations", func(t *testing.T) {
        input := &ec2.GetTransitGatewayAttachmentPropagationsInput{}
        output := &ec2.GetTransitGatewayAttachmentPropagationsOutput{}

        mockClient.On("GetTransitGatewayAttachmentPropagations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayAttachmentPropagations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayMulticastDomainAssociations", func(t *testing.T) {
        input := &ec2.GetTransitGatewayMulticastDomainAssociationsInput{}
        output := &ec2.GetTransitGatewayMulticastDomainAssociationsOutput{}

        mockClient.On("GetTransitGatewayMulticastDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayMulticastDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayPolicyTableAssociations", func(t *testing.T) {
        input := &ec2.GetTransitGatewayPolicyTableAssociationsInput{}
        output := &ec2.GetTransitGatewayPolicyTableAssociationsOutput{}

        mockClient.On("GetTransitGatewayPolicyTableAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayPolicyTableAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayPolicyTableEntries", func(t *testing.T) {
        input := &ec2.GetTransitGatewayPolicyTableEntriesInput{}
        output := &ec2.GetTransitGatewayPolicyTableEntriesOutput{}

        mockClient.On("GetTransitGatewayPolicyTableEntries", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayPolicyTableEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayPrefixListReferences", func(t *testing.T) {
        input := &ec2.GetTransitGatewayPrefixListReferencesInput{}
        output := &ec2.GetTransitGatewayPrefixListReferencesOutput{}

        mockClient.On("GetTransitGatewayPrefixListReferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayPrefixListReferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayRouteTableAssociations", func(t *testing.T) {
        input := &ec2.GetTransitGatewayRouteTableAssociationsInput{}
        output := &ec2.GetTransitGatewayRouteTableAssociationsOutput{}

        mockClient.On("GetTransitGatewayRouteTableAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayRouteTableAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransitGatewayRouteTablePropagations", func(t *testing.T) {
        input := &ec2.GetTransitGatewayRouteTablePropagationsInput{}
        output := &ec2.GetTransitGatewayRouteTablePropagationsOutput{}

        mockClient.On("GetTransitGatewayRouteTablePropagations", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransitGatewayRouteTablePropagations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVerifiedAccessEndpointPolicy", func(t *testing.T) {
        input := &ec2.GetVerifiedAccessEndpointPolicyInput{}
        output := &ec2.GetVerifiedAccessEndpointPolicyOutput{}

        mockClient.On("GetVerifiedAccessEndpointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetVerifiedAccessEndpointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVerifiedAccessGroupPolicy", func(t *testing.T) {
        input := &ec2.GetVerifiedAccessGroupPolicyInput{}
        output := &ec2.GetVerifiedAccessGroupPolicyOutput{}

        mockClient.On("GetVerifiedAccessGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetVerifiedAccessGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpnConnectionDeviceSampleConfiguration", func(t *testing.T) {
        input := &ec2.GetVpnConnectionDeviceSampleConfigurationInput{}
        output := &ec2.GetVpnConnectionDeviceSampleConfigurationOutput{}

        mockClient.On("GetVpnConnectionDeviceSampleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpnConnectionDeviceSampleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpnConnectionDeviceTypes", func(t *testing.T) {
        input := &ec2.GetVpnConnectionDeviceTypesInput{}
        output := &ec2.GetVpnConnectionDeviceTypesOutput{}

        mockClient.On("GetVpnConnectionDeviceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpnConnectionDeviceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpnTunnelReplacementStatus", func(t *testing.T) {
        input := &ec2.GetVpnTunnelReplacementStatusInput{}
        output := &ec2.GetVpnTunnelReplacementStatusOutput{}

        mockClient.On("GetVpnTunnelReplacementStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpnTunnelReplacementStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportClientVpnClientCertificateRevocationList", func(t *testing.T) {
        input := &ec2.ImportClientVpnClientCertificateRevocationListInput{}
        output := &ec2.ImportClientVpnClientCertificateRevocationListOutput{}

        mockClient.On("ImportClientVpnClientCertificateRevocationList", ctx, input).Return(output, nil)

        result, err := mockClient.ImportClientVpnClientCertificateRevocationList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportImage", func(t *testing.T) {
        input := &ec2.ImportImageInput{}
        output := &ec2.ImportImageOutput{}

        mockClient.On("ImportImage", ctx, input).Return(output, nil)

        result, err := mockClient.ImportImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportInstance", func(t *testing.T) {
        input := &ec2.ImportInstanceInput{}
        output := &ec2.ImportInstanceOutput{}

        mockClient.On("ImportInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ImportInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportKeyPair", func(t *testing.T) {
        input := &ec2.ImportKeyPairInput{}
        output := &ec2.ImportKeyPairOutput{}

        mockClient.On("ImportKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.ImportKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportSnapshot", func(t *testing.T) {
        input := &ec2.ImportSnapshotInput{}
        output := &ec2.ImportSnapshotOutput{}

        mockClient.On("ImportSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ImportSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportVolume", func(t *testing.T) {
        input := &ec2.ImportVolumeInput{}
        output := &ec2.ImportVolumeOutput{}

        mockClient.On("ImportVolume", ctx, input).Return(output, nil)

        result, err := mockClient.ImportVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImagesInRecycleBin", func(t *testing.T) {
        input := &ec2.ListImagesInRecycleBinInput{}
        output := &ec2.ListImagesInRecycleBinOutput{}

        mockClient.On("ListImagesInRecycleBin", ctx, input).Return(output, nil)

        result, err := mockClient.ListImagesInRecycleBin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSnapshotsInRecycleBin", func(t *testing.T) {
        input := &ec2.ListSnapshotsInRecycleBinInput{}
        output := &ec2.ListSnapshotsInRecycleBinOutput{}

        mockClient.On("ListSnapshotsInRecycleBin", ctx, input).Return(output, nil)

        result, err := mockClient.ListSnapshotsInRecycleBin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLockSnapshot", func(t *testing.T) {
        input := &ec2.LockSnapshotInput{}
        output := &ec2.LockSnapshotOutput{}

        mockClient.On("LockSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.LockSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyAddressAttribute", func(t *testing.T) {
        input := &ec2.ModifyAddressAttributeInput{}
        output := &ec2.ModifyAddressAttributeOutput{}

        mockClient.On("ModifyAddressAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyAddressAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyAvailabilityZoneGroup", func(t *testing.T) {
        input := &ec2.ModifyAvailabilityZoneGroupInput{}
        output := &ec2.ModifyAvailabilityZoneGroupOutput{}

        mockClient.On("ModifyAvailabilityZoneGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyAvailabilityZoneGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCapacityReservation", func(t *testing.T) {
        input := &ec2.ModifyCapacityReservationInput{}
        output := &ec2.ModifyCapacityReservationOutput{}

        mockClient.On("ModifyCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCapacityReservationFleet", func(t *testing.T) {
        input := &ec2.ModifyCapacityReservationFleetInput{}
        output := &ec2.ModifyCapacityReservationFleetOutput{}

        mockClient.On("ModifyCapacityReservationFleet", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCapacityReservationFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClientVpnEndpoint", func(t *testing.T) {
        input := &ec2.ModifyClientVpnEndpointInput{}
        output := &ec2.ModifyClientVpnEndpointOutput{}

        mockClient.On("ModifyClientVpnEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClientVpnEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDefaultCreditSpecification", func(t *testing.T) {
        input := &ec2.ModifyDefaultCreditSpecificationInput{}
        output := &ec2.ModifyDefaultCreditSpecificationOutput{}

        mockClient.On("ModifyDefaultCreditSpecification", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDefaultCreditSpecification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEbsDefaultKmsKeyId", func(t *testing.T) {
        input := &ec2.ModifyEbsDefaultKmsKeyIdInput{}
        output := &ec2.ModifyEbsDefaultKmsKeyIdOutput{}

        mockClient.On("ModifyEbsDefaultKmsKeyId", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEbsDefaultKmsKeyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyFleet", func(t *testing.T) {
        input := &ec2.ModifyFleetInput{}
        output := &ec2.ModifyFleetOutput{}

        mockClient.On("ModifyFleet", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyFpgaImageAttribute", func(t *testing.T) {
        input := &ec2.ModifyFpgaImageAttributeInput{}
        output := &ec2.ModifyFpgaImageAttributeOutput{}

        mockClient.On("ModifyFpgaImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyFpgaImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyHosts", func(t *testing.T) {
        input := &ec2.ModifyHostsInput{}
        output := &ec2.ModifyHostsOutput{}

        mockClient.On("ModifyHosts", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIdFormat", func(t *testing.T) {
        input := &ec2.ModifyIdFormatInput{}
        output := &ec2.ModifyIdFormatOutput{}

        mockClient.On("ModifyIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIdentityIdFormat", func(t *testing.T) {
        input := &ec2.ModifyIdentityIdFormatInput{}
        output := &ec2.ModifyIdentityIdFormatOutput{}

        mockClient.On("ModifyIdentityIdFormat", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIdentityIdFormat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyImageAttribute", func(t *testing.T) {
        input := &ec2.ModifyImageAttributeInput{}
        output := &ec2.ModifyImageAttributeOutput{}

        mockClient.On("ModifyImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceAttribute", func(t *testing.T) {
        input := &ec2.ModifyInstanceAttributeInput{}
        output := &ec2.ModifyInstanceAttributeOutput{}

        mockClient.On("ModifyInstanceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceCapacityReservationAttributes", func(t *testing.T) {
        input := &ec2.ModifyInstanceCapacityReservationAttributesInput{}
        output := &ec2.ModifyInstanceCapacityReservationAttributesOutput{}

        mockClient.On("ModifyInstanceCapacityReservationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceCapacityReservationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceCreditSpecification", func(t *testing.T) {
        input := &ec2.ModifyInstanceCreditSpecificationInput{}
        output := &ec2.ModifyInstanceCreditSpecificationOutput{}

        mockClient.On("ModifyInstanceCreditSpecification", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceCreditSpecification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceEventStartTime", func(t *testing.T) {
        input := &ec2.ModifyInstanceEventStartTimeInput{}
        output := &ec2.ModifyInstanceEventStartTimeOutput{}

        mockClient.On("ModifyInstanceEventStartTime", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceEventStartTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceEventWindow", func(t *testing.T) {
        input := &ec2.ModifyInstanceEventWindowInput{}
        output := &ec2.ModifyInstanceEventWindowOutput{}

        mockClient.On("ModifyInstanceEventWindow", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceEventWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceMaintenanceOptions", func(t *testing.T) {
        input := &ec2.ModifyInstanceMaintenanceOptionsInput{}
        output := &ec2.ModifyInstanceMaintenanceOptionsOutput{}

        mockClient.On("ModifyInstanceMaintenanceOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceMaintenanceOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceMetadataDefaults", func(t *testing.T) {
        input := &ec2.ModifyInstanceMetadataDefaultsInput{}
        output := &ec2.ModifyInstanceMetadataDefaultsOutput{}

        mockClient.On("ModifyInstanceMetadataDefaults", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceMetadataDefaults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceMetadataOptions", func(t *testing.T) {
        input := &ec2.ModifyInstanceMetadataOptionsInput{}
        output := &ec2.ModifyInstanceMetadataOptionsOutput{}

        mockClient.On("ModifyInstanceMetadataOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceMetadataOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstancePlacement", func(t *testing.T) {
        input := &ec2.ModifyInstancePlacementInput{}
        output := &ec2.ModifyInstancePlacementOutput{}

        mockClient.On("ModifyInstancePlacement", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstancePlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIpam", func(t *testing.T) {
        input := &ec2.ModifyIpamInput{}
        output := &ec2.ModifyIpamOutput{}

        mockClient.On("ModifyIpam", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIpam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIpamPool", func(t *testing.T) {
        input := &ec2.ModifyIpamPoolInput{}
        output := &ec2.ModifyIpamPoolOutput{}

        mockClient.On("ModifyIpamPool", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIpamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIpamResourceCidr", func(t *testing.T) {
        input := &ec2.ModifyIpamResourceCidrInput{}
        output := &ec2.ModifyIpamResourceCidrOutput{}

        mockClient.On("ModifyIpamResourceCidr", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIpamResourceCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIpamResourceDiscovery", func(t *testing.T) {
        input := &ec2.ModifyIpamResourceDiscoveryInput{}
        output := &ec2.ModifyIpamResourceDiscoveryOutput{}

        mockClient.On("ModifyIpamResourceDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIpamResourceDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIpamScope", func(t *testing.T) {
        input := &ec2.ModifyIpamScopeInput{}
        output := &ec2.ModifyIpamScopeOutput{}

        mockClient.On("ModifyIpamScope", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIpamScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyLaunchTemplate", func(t *testing.T) {
        input := &ec2.ModifyLaunchTemplateInput{}
        output := &ec2.ModifyLaunchTemplateOutput{}

        mockClient.On("ModifyLaunchTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyLaunchTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyLocalGatewayRoute", func(t *testing.T) {
        input := &ec2.ModifyLocalGatewayRouteInput{}
        output := &ec2.ModifyLocalGatewayRouteOutput{}

        mockClient.On("ModifyLocalGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyLocalGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyManagedPrefixList", func(t *testing.T) {
        input := &ec2.ModifyManagedPrefixListInput{}
        output := &ec2.ModifyManagedPrefixListOutput{}

        mockClient.On("ModifyManagedPrefixList", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyManagedPrefixList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyNetworkInterfaceAttribute", func(t *testing.T) {
        input := &ec2.ModifyNetworkInterfaceAttributeInput{}
        output := &ec2.ModifyNetworkInterfaceAttributeOutput{}

        mockClient.On("ModifyNetworkInterfaceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyNetworkInterfaceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyPrivateDnsNameOptions", func(t *testing.T) {
        input := &ec2.ModifyPrivateDnsNameOptionsInput{}
        output := &ec2.ModifyPrivateDnsNameOptionsOutput{}

        mockClient.On("ModifyPrivateDnsNameOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyPrivateDnsNameOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReservedInstances", func(t *testing.T) {
        input := &ec2.ModifyReservedInstancesInput{}
        output := &ec2.ModifyReservedInstancesOutput{}

        mockClient.On("ModifyReservedInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReservedInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySecurityGroupRules", func(t *testing.T) {
        input := &ec2.ModifySecurityGroupRulesInput{}
        output := &ec2.ModifySecurityGroupRulesOutput{}

        mockClient.On("ModifySecurityGroupRules", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySecurityGroupRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySnapshotAttribute", func(t *testing.T) {
        input := &ec2.ModifySnapshotAttributeInput{}
        output := &ec2.ModifySnapshotAttributeOutput{}

        mockClient.On("ModifySnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySnapshotTier", func(t *testing.T) {
        input := &ec2.ModifySnapshotTierInput{}
        output := &ec2.ModifySnapshotTierOutput{}

        mockClient.On("ModifySnapshotTier", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySnapshotTier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySpotFleetRequest", func(t *testing.T) {
        input := &ec2.ModifySpotFleetRequestInput{}
        output := &ec2.ModifySpotFleetRequestOutput{}

        mockClient.On("ModifySpotFleetRequest", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySpotFleetRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySubnetAttribute", func(t *testing.T) {
        input := &ec2.ModifySubnetAttributeInput{}
        output := &ec2.ModifySubnetAttributeOutput{}

        mockClient.On("ModifySubnetAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySubnetAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTrafficMirrorFilterNetworkServices", func(t *testing.T) {
        input := &ec2.ModifyTrafficMirrorFilterNetworkServicesInput{}
        output := &ec2.ModifyTrafficMirrorFilterNetworkServicesOutput{}

        mockClient.On("ModifyTrafficMirrorFilterNetworkServices", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTrafficMirrorFilterNetworkServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTrafficMirrorFilterRule", func(t *testing.T) {
        input := &ec2.ModifyTrafficMirrorFilterRuleInput{}
        output := &ec2.ModifyTrafficMirrorFilterRuleOutput{}

        mockClient.On("ModifyTrafficMirrorFilterRule", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTrafficMirrorFilterRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTrafficMirrorSession", func(t *testing.T) {
        input := &ec2.ModifyTrafficMirrorSessionInput{}
        output := &ec2.ModifyTrafficMirrorSessionOutput{}

        mockClient.On("ModifyTrafficMirrorSession", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTrafficMirrorSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTransitGateway", func(t *testing.T) {
        input := &ec2.ModifyTransitGatewayInput{}
        output := &ec2.ModifyTransitGatewayOutput{}

        mockClient.On("ModifyTransitGateway", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTransitGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTransitGatewayPrefixListReference", func(t *testing.T) {
        input := &ec2.ModifyTransitGatewayPrefixListReferenceInput{}
        output := &ec2.ModifyTransitGatewayPrefixListReferenceOutput{}

        mockClient.On("ModifyTransitGatewayPrefixListReference", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTransitGatewayPrefixListReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTransitGatewayVpcAttachment", func(t *testing.T) {
        input := &ec2.ModifyTransitGatewayVpcAttachmentInput{}
        output := &ec2.ModifyTransitGatewayVpcAttachmentOutput{}

        mockClient.On("ModifyTransitGatewayVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTransitGatewayVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessEndpoint", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessEndpointInput{}
        output := &ec2.ModifyVerifiedAccessEndpointOutput{}

        mockClient.On("ModifyVerifiedAccessEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessEndpointPolicy", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessEndpointPolicyInput{}
        output := &ec2.ModifyVerifiedAccessEndpointPolicyOutput{}

        mockClient.On("ModifyVerifiedAccessEndpointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessEndpointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessGroup", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessGroupInput{}
        output := &ec2.ModifyVerifiedAccessGroupOutput{}

        mockClient.On("ModifyVerifiedAccessGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessGroupPolicy", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessGroupPolicyInput{}
        output := &ec2.ModifyVerifiedAccessGroupPolicyOutput{}

        mockClient.On("ModifyVerifiedAccessGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessInstance", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessInstanceInput{}
        output := &ec2.ModifyVerifiedAccessInstanceOutput{}

        mockClient.On("ModifyVerifiedAccessInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessInstanceLoggingConfiguration", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessInstanceLoggingConfigurationInput{}
        output := &ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput{}

        mockClient.On("ModifyVerifiedAccessInstanceLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessInstanceLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVerifiedAccessTrustProvider", func(t *testing.T) {
        input := &ec2.ModifyVerifiedAccessTrustProviderInput{}
        output := &ec2.ModifyVerifiedAccessTrustProviderOutput{}

        mockClient.On("ModifyVerifiedAccessTrustProvider", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVerifiedAccessTrustProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVolume", func(t *testing.T) {
        input := &ec2.ModifyVolumeInput{}
        output := &ec2.ModifyVolumeOutput{}

        mockClient.On("ModifyVolume", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVolumeAttribute", func(t *testing.T) {
        input := &ec2.ModifyVolumeAttributeInput{}
        output := &ec2.ModifyVolumeAttributeOutput{}

        mockClient.On("ModifyVolumeAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVolumeAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcAttribute", func(t *testing.T) {
        input := &ec2.ModifyVpcAttributeInput{}
        output := &ec2.ModifyVpcAttributeOutput{}

        mockClient.On("ModifyVpcAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcEndpoint", func(t *testing.T) {
        input := &ec2.ModifyVpcEndpointInput{}
        output := &ec2.ModifyVpcEndpointOutput{}

        mockClient.On("ModifyVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcEndpointConnectionNotification", func(t *testing.T) {
        input := &ec2.ModifyVpcEndpointConnectionNotificationInput{}
        output := &ec2.ModifyVpcEndpointConnectionNotificationOutput{}

        mockClient.On("ModifyVpcEndpointConnectionNotification", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcEndpointConnectionNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcEndpointServiceConfiguration", func(t *testing.T) {
        input := &ec2.ModifyVpcEndpointServiceConfigurationInput{}
        output := &ec2.ModifyVpcEndpointServiceConfigurationOutput{}

        mockClient.On("ModifyVpcEndpointServiceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcEndpointServiceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcEndpointServicePayerResponsibility", func(t *testing.T) {
        input := &ec2.ModifyVpcEndpointServicePayerResponsibilityInput{}
        output := &ec2.ModifyVpcEndpointServicePayerResponsibilityOutput{}

        mockClient.On("ModifyVpcEndpointServicePayerResponsibility", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcEndpointServicePayerResponsibility(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcEndpointServicePermissions", func(t *testing.T) {
        input := &ec2.ModifyVpcEndpointServicePermissionsInput{}
        output := &ec2.ModifyVpcEndpointServicePermissionsOutput{}

        mockClient.On("ModifyVpcEndpointServicePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcEndpointServicePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcPeeringConnectionOptions", func(t *testing.T) {
        input := &ec2.ModifyVpcPeeringConnectionOptionsInput{}
        output := &ec2.ModifyVpcPeeringConnectionOptionsOutput{}

        mockClient.On("ModifyVpcPeeringConnectionOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcPeeringConnectionOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpcTenancy", func(t *testing.T) {
        input := &ec2.ModifyVpcTenancyInput{}
        output := &ec2.ModifyVpcTenancyOutput{}

        mockClient.On("ModifyVpcTenancy", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpcTenancy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpnConnection", func(t *testing.T) {
        input := &ec2.ModifyVpnConnectionInput{}
        output := &ec2.ModifyVpnConnectionOutput{}

        mockClient.On("ModifyVpnConnection", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpnConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpnConnectionOptions", func(t *testing.T) {
        input := &ec2.ModifyVpnConnectionOptionsInput{}
        output := &ec2.ModifyVpnConnectionOptionsOutput{}

        mockClient.On("ModifyVpnConnectionOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpnConnectionOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpnTunnelCertificate", func(t *testing.T) {
        input := &ec2.ModifyVpnTunnelCertificateInput{}
        output := &ec2.ModifyVpnTunnelCertificateOutput{}

        mockClient.On("ModifyVpnTunnelCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpnTunnelCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyVpnTunnelOptions", func(t *testing.T) {
        input := &ec2.ModifyVpnTunnelOptionsInput{}
        output := &ec2.ModifyVpnTunnelOptionsOutput{}

        mockClient.On("ModifyVpnTunnelOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyVpnTunnelOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMonitorInstances", func(t *testing.T) {
        input := &ec2.MonitorInstancesInput{}
        output := &ec2.MonitorInstancesOutput{}

        mockClient.On("MonitorInstances", ctx, input).Return(output, nil)

        result, err := mockClient.MonitorInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMoveAddressToVpc", func(t *testing.T) {
        input := &ec2.MoveAddressToVpcInput{}
        output := &ec2.MoveAddressToVpcOutput{}

        mockClient.On("MoveAddressToVpc", ctx, input).Return(output, nil)

        result, err := mockClient.MoveAddressToVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMoveByoipCidrToIpam", func(t *testing.T) {
        input := &ec2.MoveByoipCidrToIpamInput{}
        output := &ec2.MoveByoipCidrToIpamOutput{}

        mockClient.On("MoveByoipCidrToIpam", ctx, input).Return(output, nil)

        result, err := mockClient.MoveByoipCidrToIpam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionByoipCidr", func(t *testing.T) {
        input := &ec2.ProvisionByoipCidrInput{}
        output := &ec2.ProvisionByoipCidrOutput{}

        mockClient.On("ProvisionByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionIpamByoasn", func(t *testing.T) {
        input := &ec2.ProvisionIpamByoasnInput{}
        output := &ec2.ProvisionIpamByoasnOutput{}

        mockClient.On("ProvisionIpamByoasn", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionIpamByoasn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionIpamPoolCidr", func(t *testing.T) {
        input := &ec2.ProvisionIpamPoolCidrInput{}
        output := &ec2.ProvisionIpamPoolCidrOutput{}

        mockClient.On("ProvisionIpamPoolCidr", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionIpamPoolCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionPublicIpv4PoolCidr", func(t *testing.T) {
        input := &ec2.ProvisionPublicIpv4PoolCidrInput{}
        output := &ec2.ProvisionPublicIpv4PoolCidrOutput{}

        mockClient.On("ProvisionPublicIpv4PoolCidr", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionPublicIpv4PoolCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseCapacityBlock", func(t *testing.T) {
        input := &ec2.PurchaseCapacityBlockInput{}
        output := &ec2.PurchaseCapacityBlockOutput{}

        mockClient.On("PurchaseCapacityBlock", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseCapacityBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseHostReservation", func(t *testing.T) {
        input := &ec2.PurchaseHostReservationInput{}
        output := &ec2.PurchaseHostReservationOutput{}

        mockClient.On("PurchaseHostReservation", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseHostReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedInstancesOffering", func(t *testing.T) {
        input := &ec2.PurchaseReservedInstancesOfferingInput{}
        output := &ec2.PurchaseReservedInstancesOfferingOutput{}

        mockClient.On("PurchaseReservedInstancesOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedInstancesOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseScheduledInstances", func(t *testing.T) {
        input := &ec2.PurchaseScheduledInstancesInput{}
        output := &ec2.PurchaseScheduledInstancesOutput{}

        mockClient.On("PurchaseScheduledInstances", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseScheduledInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootInstances", func(t *testing.T) {
        input := &ec2.RebootInstancesInput{}
        output := &ec2.RebootInstancesOutput{}

        mockClient.On("RebootInstances", ctx, input).Return(output, nil)

        result, err := mockClient.RebootInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterImage", func(t *testing.T) {
        input := &ec2.RegisterImageInput{}
        output := &ec2.RegisterImageOutput{}

        mockClient.On("RegisterImage", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterInstanceEventNotificationAttributes", func(t *testing.T) {
        input := &ec2.RegisterInstanceEventNotificationAttributesInput{}
        output := &ec2.RegisterInstanceEventNotificationAttributesOutput{}

        mockClient.On("RegisterInstanceEventNotificationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterInstanceEventNotificationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTransitGatewayMulticastGroupMembers", func(t *testing.T) {
        input := &ec2.RegisterTransitGatewayMulticastGroupMembersInput{}
        output := &ec2.RegisterTransitGatewayMulticastGroupMembersOutput{}

        mockClient.On("RegisterTransitGatewayMulticastGroupMembers", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTransitGatewayMulticastGroupMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTransitGatewayMulticastGroupSources", func(t *testing.T) {
        input := &ec2.RegisterTransitGatewayMulticastGroupSourcesInput{}
        output := &ec2.RegisterTransitGatewayMulticastGroupSourcesOutput{}

        mockClient.On("RegisterTransitGatewayMulticastGroupSources", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTransitGatewayMulticastGroupSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectTransitGatewayMulticastDomainAssociations", func(t *testing.T) {
        input := &ec2.RejectTransitGatewayMulticastDomainAssociationsInput{}
        output := &ec2.RejectTransitGatewayMulticastDomainAssociationsOutput{}

        mockClient.On("RejectTransitGatewayMulticastDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.RejectTransitGatewayMulticastDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectTransitGatewayPeeringAttachment", func(t *testing.T) {
        input := &ec2.RejectTransitGatewayPeeringAttachmentInput{}
        output := &ec2.RejectTransitGatewayPeeringAttachmentOutput{}

        mockClient.On("RejectTransitGatewayPeeringAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.RejectTransitGatewayPeeringAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectTransitGatewayVpcAttachment", func(t *testing.T) {
        input := &ec2.RejectTransitGatewayVpcAttachmentInput{}
        output := &ec2.RejectTransitGatewayVpcAttachmentOutput{}

        mockClient.On("RejectTransitGatewayVpcAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.RejectTransitGatewayVpcAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectVpcEndpointConnections", func(t *testing.T) {
        input := &ec2.RejectVpcEndpointConnectionsInput{}
        output := &ec2.RejectVpcEndpointConnectionsOutput{}

        mockClient.On("RejectVpcEndpointConnections", ctx, input).Return(output, nil)

        result, err := mockClient.RejectVpcEndpointConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectVpcPeeringConnection", func(t *testing.T) {
        input := &ec2.RejectVpcPeeringConnectionInput{}
        output := &ec2.RejectVpcPeeringConnectionOutput{}

        mockClient.On("RejectVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.RejectVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseAddress", func(t *testing.T) {
        input := &ec2.ReleaseAddressInput{}
        output := &ec2.ReleaseAddressOutput{}

        mockClient.On("ReleaseAddress", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseHosts", func(t *testing.T) {
        input := &ec2.ReleaseHostsInput{}
        output := &ec2.ReleaseHostsOutput{}

        mockClient.On("ReleaseHosts", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseIpamPoolAllocation", func(t *testing.T) {
        input := &ec2.ReleaseIpamPoolAllocationInput{}
        output := &ec2.ReleaseIpamPoolAllocationOutput{}

        mockClient.On("ReleaseIpamPoolAllocation", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseIpamPoolAllocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceIamInstanceProfileAssociation", func(t *testing.T) {
        input := &ec2.ReplaceIamInstanceProfileAssociationInput{}
        output := &ec2.ReplaceIamInstanceProfileAssociationOutput{}

        mockClient.On("ReplaceIamInstanceProfileAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceIamInstanceProfileAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceNetworkAclAssociation", func(t *testing.T) {
        input := &ec2.ReplaceNetworkAclAssociationInput{}
        output := &ec2.ReplaceNetworkAclAssociationOutput{}

        mockClient.On("ReplaceNetworkAclAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceNetworkAclAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceNetworkAclEntry", func(t *testing.T) {
        input := &ec2.ReplaceNetworkAclEntryInput{}
        output := &ec2.ReplaceNetworkAclEntryOutput{}

        mockClient.On("ReplaceNetworkAclEntry", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceNetworkAclEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceRoute", func(t *testing.T) {
        input := &ec2.ReplaceRouteInput{}
        output := &ec2.ReplaceRouteOutput{}

        mockClient.On("ReplaceRoute", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceRouteTableAssociation", func(t *testing.T) {
        input := &ec2.ReplaceRouteTableAssociationInput{}
        output := &ec2.ReplaceRouteTableAssociationOutput{}

        mockClient.On("ReplaceRouteTableAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceRouteTableAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceTransitGatewayRoute", func(t *testing.T) {
        input := &ec2.ReplaceTransitGatewayRouteInput{}
        output := &ec2.ReplaceTransitGatewayRouteOutput{}

        mockClient.On("ReplaceTransitGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceTransitGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceVpnTunnel", func(t *testing.T) {
        input := &ec2.ReplaceVpnTunnelInput{}
        output := &ec2.ReplaceVpnTunnelOutput{}

        mockClient.On("ReplaceVpnTunnel", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceVpnTunnel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReportInstanceStatus", func(t *testing.T) {
        input := &ec2.ReportInstanceStatusInput{}
        output := &ec2.ReportInstanceStatusOutput{}

        mockClient.On("ReportInstanceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ReportInstanceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestSpotFleet", func(t *testing.T) {
        input := &ec2.RequestSpotFleetInput{}
        output := &ec2.RequestSpotFleetOutput{}

        mockClient.On("RequestSpotFleet", ctx, input).Return(output, nil)

        result, err := mockClient.RequestSpotFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestSpotInstances", func(t *testing.T) {
        input := &ec2.RequestSpotInstancesInput{}
        output := &ec2.RequestSpotInstancesOutput{}

        mockClient.On("RequestSpotInstances", ctx, input).Return(output, nil)

        result, err := mockClient.RequestSpotInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetAddressAttribute", func(t *testing.T) {
        input := &ec2.ResetAddressAttributeInput{}
        output := &ec2.ResetAddressAttributeOutput{}

        mockClient.On("ResetAddressAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetAddressAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetEbsDefaultKmsKeyId", func(t *testing.T) {
        input := &ec2.ResetEbsDefaultKmsKeyIdInput{}
        output := &ec2.ResetEbsDefaultKmsKeyIdOutput{}

        mockClient.On("ResetEbsDefaultKmsKeyId", ctx, input).Return(output, nil)

        result, err := mockClient.ResetEbsDefaultKmsKeyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetFpgaImageAttribute", func(t *testing.T) {
        input := &ec2.ResetFpgaImageAttributeInput{}
        output := &ec2.ResetFpgaImageAttributeOutput{}

        mockClient.On("ResetFpgaImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetFpgaImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetImageAttribute", func(t *testing.T) {
        input := &ec2.ResetImageAttributeInput{}
        output := &ec2.ResetImageAttributeOutput{}

        mockClient.On("ResetImageAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetImageAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetInstanceAttribute", func(t *testing.T) {
        input := &ec2.ResetInstanceAttributeInput{}
        output := &ec2.ResetInstanceAttributeOutput{}

        mockClient.On("ResetInstanceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetInstanceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetNetworkInterfaceAttribute", func(t *testing.T) {
        input := &ec2.ResetNetworkInterfaceAttributeInput{}
        output := &ec2.ResetNetworkInterfaceAttributeOutput{}

        mockClient.On("ResetNetworkInterfaceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetNetworkInterfaceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetSnapshotAttribute", func(t *testing.T) {
        input := &ec2.ResetSnapshotAttributeInput{}
        output := &ec2.ResetSnapshotAttributeOutput{}

        mockClient.On("ResetSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ResetSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreAddressToClassic", func(t *testing.T) {
        input := &ec2.RestoreAddressToClassicInput{}
        output := &ec2.RestoreAddressToClassicOutput{}

        mockClient.On("RestoreAddressToClassic", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreAddressToClassic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreImageFromRecycleBin", func(t *testing.T) {
        input := &ec2.RestoreImageFromRecycleBinInput{}
        output := &ec2.RestoreImageFromRecycleBinOutput{}

        mockClient.On("RestoreImageFromRecycleBin", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreImageFromRecycleBin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreManagedPrefixListVersion", func(t *testing.T) {
        input := &ec2.RestoreManagedPrefixListVersionInput{}
        output := &ec2.RestoreManagedPrefixListVersionOutput{}

        mockClient.On("RestoreManagedPrefixListVersion", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreManagedPrefixListVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreSnapshotFromRecycleBin", func(t *testing.T) {
        input := &ec2.RestoreSnapshotFromRecycleBinInput{}
        output := &ec2.RestoreSnapshotFromRecycleBinOutput{}

        mockClient.On("RestoreSnapshotFromRecycleBin", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreSnapshotFromRecycleBin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreSnapshotTier", func(t *testing.T) {
        input := &ec2.RestoreSnapshotTierInput{}
        output := &ec2.RestoreSnapshotTierOutput{}

        mockClient.On("RestoreSnapshotTier", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreSnapshotTier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeClientVpnIngress", func(t *testing.T) {
        input := &ec2.RevokeClientVpnIngressInput{}
        output := &ec2.RevokeClientVpnIngressOutput{}

        mockClient.On("RevokeClientVpnIngress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeClientVpnIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSecurityGroupEgress", func(t *testing.T) {
        input := &ec2.RevokeSecurityGroupEgressInput{}
        output := &ec2.RevokeSecurityGroupEgressOutput{}

        mockClient.On("RevokeSecurityGroupEgress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSecurityGroupEgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSecurityGroupIngress", func(t *testing.T) {
        input := &ec2.RevokeSecurityGroupIngressInput{}
        output := &ec2.RevokeSecurityGroupIngressOutput{}

        mockClient.On("RevokeSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunInstances", func(t *testing.T) {
        input := &ec2.RunInstancesInput{}
        output := &ec2.RunInstancesOutput{}

        mockClient.On("RunInstances", ctx, input).Return(output, nil)

        result, err := mockClient.RunInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunScheduledInstances", func(t *testing.T) {
        input := &ec2.RunScheduledInstancesInput{}
        output := &ec2.RunScheduledInstancesOutput{}

        mockClient.On("RunScheduledInstances", ctx, input).Return(output, nil)

        result, err := mockClient.RunScheduledInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchLocalGatewayRoutes", func(t *testing.T) {
        input := &ec2.SearchLocalGatewayRoutesInput{}
        output := &ec2.SearchLocalGatewayRoutesOutput{}

        mockClient.On("SearchLocalGatewayRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.SearchLocalGatewayRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTransitGatewayMulticastGroups", func(t *testing.T) {
        input := &ec2.SearchTransitGatewayMulticastGroupsInput{}
        output := &ec2.SearchTransitGatewayMulticastGroupsOutput{}

        mockClient.On("SearchTransitGatewayMulticastGroups", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTransitGatewayMulticastGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTransitGatewayRoutes", func(t *testing.T) {
        input := &ec2.SearchTransitGatewayRoutesInput{}
        output := &ec2.SearchTransitGatewayRoutesOutput{}

        mockClient.On("SearchTransitGatewayRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTransitGatewayRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDiagnosticInterrupt", func(t *testing.T) {
        input := &ec2.SendDiagnosticInterruptInput{}
        output := &ec2.SendDiagnosticInterruptOutput{}

        mockClient.On("SendDiagnosticInterrupt", ctx, input).Return(output, nil)

        result, err := mockClient.SendDiagnosticInterrupt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstances", func(t *testing.T) {
        input := &ec2.StartInstancesInput{}
        output := &ec2.StartInstancesOutput{}

        mockClient.On("StartInstances", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNetworkInsightsAccessScopeAnalysis", func(t *testing.T) {
        input := &ec2.StartNetworkInsightsAccessScopeAnalysisInput{}
        output := &ec2.StartNetworkInsightsAccessScopeAnalysisOutput{}

        mockClient.On("StartNetworkInsightsAccessScopeAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartNetworkInsightsAccessScopeAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNetworkInsightsAnalysis", func(t *testing.T) {
        input := &ec2.StartNetworkInsightsAnalysisInput{}
        output := &ec2.StartNetworkInsightsAnalysisOutput{}

        mockClient.On("StartNetworkInsightsAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartNetworkInsightsAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVpcEndpointServicePrivateDnsVerification", func(t *testing.T) {
        input := &ec2.StartVpcEndpointServicePrivateDnsVerificationInput{}
        output := &ec2.StartVpcEndpointServicePrivateDnsVerificationOutput{}

        mockClient.On("StartVpcEndpointServicePrivateDnsVerification", ctx, input).Return(output, nil)

        result, err := mockClient.StartVpcEndpointServicePrivateDnsVerification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInstances", func(t *testing.T) {
        input := &ec2.StopInstancesInput{}
        output := &ec2.StopInstancesOutput{}

        mockClient.On("StopInstances", ctx, input).Return(output, nil)

        result, err := mockClient.StopInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateClientVpnConnections", func(t *testing.T) {
        input := &ec2.TerminateClientVpnConnectionsInput{}
        output := &ec2.TerminateClientVpnConnectionsOutput{}

        mockClient.On("TerminateClientVpnConnections", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateClientVpnConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateInstances", func(t *testing.T) {
        input := &ec2.TerminateInstancesInput{}
        output := &ec2.TerminateInstancesOutput{}

        mockClient.On("TerminateInstances", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnassignIpv6Addresses", func(t *testing.T) {
        input := &ec2.UnassignIpv6AddressesInput{}
        output := &ec2.UnassignIpv6AddressesOutput{}

        mockClient.On("UnassignIpv6Addresses", ctx, input).Return(output, nil)

        result, err := mockClient.UnassignIpv6Addresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnassignPrivateIpAddresses", func(t *testing.T) {
        input := &ec2.UnassignPrivateIpAddressesInput{}
        output := &ec2.UnassignPrivateIpAddressesOutput{}

        mockClient.On("UnassignPrivateIpAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.UnassignPrivateIpAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnassignPrivateNatGatewayAddress", func(t *testing.T) {
        input := &ec2.UnassignPrivateNatGatewayAddressInput{}
        output := &ec2.UnassignPrivateNatGatewayAddressOutput{}

        mockClient.On("UnassignPrivateNatGatewayAddress", ctx, input).Return(output, nil)

        result, err := mockClient.UnassignPrivateNatGatewayAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnlockSnapshot", func(t *testing.T) {
        input := &ec2.UnlockSnapshotInput{}
        output := &ec2.UnlockSnapshotOutput{}

        mockClient.On("UnlockSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.UnlockSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnmonitorInstances", func(t *testing.T) {
        input := &ec2.UnmonitorInstancesInput{}
        output := &ec2.UnmonitorInstancesOutput{}

        mockClient.On("UnmonitorInstances", ctx, input).Return(output, nil)

        result, err := mockClient.UnmonitorInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityGroupRuleDescriptionsEgress", func(t *testing.T) {
        input := &ec2.UpdateSecurityGroupRuleDescriptionsEgressInput{}
        output := &ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput{}

        mockClient.On("UpdateSecurityGroupRuleDescriptionsEgress", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityGroupRuleDescriptionsEgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityGroupRuleDescriptionsIngress", func(t *testing.T) {
        input := &ec2.UpdateSecurityGroupRuleDescriptionsIngressInput{}
        output := &ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput{}

        mockClient.On("UpdateSecurityGroupRuleDescriptionsIngress", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityGroupRuleDescriptionsIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestWithdrawByoipCidr", func(t *testing.T) {
        input := &ec2.WithdrawByoipCidrInput{}
        output := &ec2.WithdrawByoipCidrOutput{}

        mockClient.On("WithdrawByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.WithdrawByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
