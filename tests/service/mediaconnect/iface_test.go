package mediaconnect_test

// tests for the mediaconnect service interface for this ../../../service/mediaconnect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediaconnect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediaconnect/mediaconnect_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediaconnectServiceCanBeMocked(t *testing.T) {
	var iface mediaconnect_iface.IClient
	iface = &mediaconnect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediaconnect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddBridgeOutputs", func(t *testing.T) {
        input := &mediaconnect.AddBridgeOutputsInput{}
        output := &mediaconnect.AddBridgeOutputsOutput{}

        mockClient.On("AddBridgeOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.AddBridgeOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddBridgeSources", func(t *testing.T) {
        input := &mediaconnect.AddBridgeSourcesInput{}
        output := &mediaconnect.AddBridgeSourcesOutput{}

        mockClient.On("AddBridgeSources", ctx, input).Return(output, nil)

        result, err := mockClient.AddBridgeSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddFlowMediaStreams", func(t *testing.T) {
        input := &mediaconnect.AddFlowMediaStreamsInput{}
        output := &mediaconnect.AddFlowMediaStreamsOutput{}

        mockClient.On("AddFlowMediaStreams", ctx, input).Return(output, nil)

        result, err := mockClient.AddFlowMediaStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddFlowOutputs", func(t *testing.T) {
        input := &mediaconnect.AddFlowOutputsInput{}
        output := &mediaconnect.AddFlowOutputsOutput{}

        mockClient.On("AddFlowOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.AddFlowOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddFlowSources", func(t *testing.T) {
        input := &mediaconnect.AddFlowSourcesInput{}
        output := &mediaconnect.AddFlowSourcesOutput{}

        mockClient.On("AddFlowSources", ctx, input).Return(output, nil)

        result, err := mockClient.AddFlowSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddFlowVpcInterfaces", func(t *testing.T) {
        input := &mediaconnect.AddFlowVpcInterfacesInput{}
        output := &mediaconnect.AddFlowVpcInterfacesOutput{}

        mockClient.On("AddFlowVpcInterfaces", ctx, input).Return(output, nil)

        result, err := mockClient.AddFlowVpcInterfaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBridge", func(t *testing.T) {
        input := &mediaconnect.CreateBridgeInput{}
        output := &mediaconnect.CreateBridgeOutput{}

        mockClient.On("CreateBridge", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBridge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlow", func(t *testing.T) {
        input := &mediaconnect.CreateFlowInput{}
        output := &mediaconnect.CreateFlowOutput{}

        mockClient.On("CreateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGateway", func(t *testing.T) {
        input := &mediaconnect.CreateGatewayInput{}
        output := &mediaconnect.CreateGatewayOutput{}

        mockClient.On("CreateGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBridge", func(t *testing.T) {
        input := &mediaconnect.DeleteBridgeInput{}
        output := &mediaconnect.DeleteBridgeOutput{}

        mockClient.On("DeleteBridge", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBridge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlow", func(t *testing.T) {
        input := &mediaconnect.DeleteFlowInput{}
        output := &mediaconnect.DeleteFlowOutput{}

        mockClient.On("DeleteFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGateway", func(t *testing.T) {
        input := &mediaconnect.DeleteGatewayInput{}
        output := &mediaconnect.DeleteGatewayOutput{}

        mockClient.On("DeleteGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterGatewayInstance", func(t *testing.T) {
        input := &mediaconnect.DeregisterGatewayInstanceInput{}
        output := &mediaconnect.DeregisterGatewayInstanceOutput{}

        mockClient.On("DeregisterGatewayInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterGatewayInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBridge", func(t *testing.T) {
        input := &mediaconnect.DescribeBridgeInput{}
        output := &mediaconnect.DescribeBridgeOutput{}

        mockClient.On("DescribeBridge", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBridge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlow", func(t *testing.T) {
        input := &mediaconnect.DescribeFlowInput{}
        output := &mediaconnect.DescribeFlowOutput{}

        mockClient.On("DescribeFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlowSourceMetadata", func(t *testing.T) {
        input := &mediaconnect.DescribeFlowSourceMetadataInput{}
        output := &mediaconnect.DescribeFlowSourceMetadataOutput{}

        mockClient.On("DescribeFlowSourceMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlowSourceMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGateway", func(t *testing.T) {
        input := &mediaconnect.DescribeGatewayInput{}
        output := &mediaconnect.DescribeGatewayOutput{}

        mockClient.On("DescribeGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGatewayInstance", func(t *testing.T) {
        input := &mediaconnect.DescribeGatewayInstanceInput{}
        output := &mediaconnect.DescribeGatewayInstanceOutput{}

        mockClient.On("DescribeGatewayInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGatewayInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOffering", func(t *testing.T) {
        input := &mediaconnect.DescribeOfferingInput{}
        output := &mediaconnect.DescribeOfferingOutput{}

        mockClient.On("DescribeOffering", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservation", func(t *testing.T) {
        input := &mediaconnect.DescribeReservationInput{}
        output := &mediaconnect.DescribeReservationOutput{}

        mockClient.On("DescribeReservation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGrantFlowEntitlements", func(t *testing.T) {
        input := &mediaconnect.GrantFlowEntitlementsInput{}
        output := &mediaconnect.GrantFlowEntitlementsOutput{}

        mockClient.On("GrantFlowEntitlements", ctx, input).Return(output, nil)

        result, err := mockClient.GrantFlowEntitlements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBridges", func(t *testing.T) {
        input := &mediaconnect.ListBridgesInput{}
        output := &mediaconnect.ListBridgesOutput{}

        mockClient.On("ListBridges", ctx, input).Return(output, nil)

        result, err := mockClient.ListBridges(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntitlements", func(t *testing.T) {
        input := &mediaconnect.ListEntitlementsInput{}
        output := &mediaconnect.ListEntitlementsOutput{}

        mockClient.On("ListEntitlements", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntitlements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlows", func(t *testing.T) {
        input := &mediaconnect.ListFlowsInput{}
        output := &mediaconnect.ListFlowsOutput{}

        mockClient.On("ListFlows", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGatewayInstances", func(t *testing.T) {
        input := &mediaconnect.ListGatewayInstancesInput{}
        output := &mediaconnect.ListGatewayInstancesOutput{}

        mockClient.On("ListGatewayInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListGatewayInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGateways", func(t *testing.T) {
        input := &mediaconnect.ListGatewaysInput{}
        output := &mediaconnect.ListGatewaysOutput{}

        mockClient.On("ListGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOfferings", func(t *testing.T) {
        input := &mediaconnect.ListOfferingsInput{}
        output := &mediaconnect.ListOfferingsOutput{}

        mockClient.On("ListOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.ListOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReservations", func(t *testing.T) {
        input := &mediaconnect.ListReservationsInput{}
        output := &mediaconnect.ListReservationsOutput{}

        mockClient.On("ListReservations", ctx, input).Return(output, nil)

        result, err := mockClient.ListReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediaconnect.ListTagsForResourceInput{}
        output := &mediaconnect.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseOffering", func(t *testing.T) {
        input := &mediaconnect.PurchaseOfferingInput{}
        output := &mediaconnect.PurchaseOfferingOutput{}

        mockClient.On("PurchaseOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveBridgeOutput", func(t *testing.T) {
        input := &mediaconnect.RemoveBridgeOutputInput{}
        output := &mediaconnect.RemoveBridgeOutputOutput{}

        mockClient.On("RemoveBridgeOutput", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveBridgeOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveBridgeSource", func(t *testing.T) {
        input := &mediaconnect.RemoveBridgeSourceInput{}
        output := &mediaconnect.RemoveBridgeSourceOutput{}

        mockClient.On("RemoveBridgeSource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveBridgeSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFlowMediaStream", func(t *testing.T) {
        input := &mediaconnect.RemoveFlowMediaStreamInput{}
        output := &mediaconnect.RemoveFlowMediaStreamOutput{}

        mockClient.On("RemoveFlowMediaStream", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFlowMediaStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFlowOutput", func(t *testing.T) {
        input := &mediaconnect.RemoveFlowOutputInput{}
        output := &mediaconnect.RemoveFlowOutputOutput{}

        mockClient.On("RemoveFlowOutput", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFlowOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFlowSource", func(t *testing.T) {
        input := &mediaconnect.RemoveFlowSourceInput{}
        output := &mediaconnect.RemoveFlowSourceOutput{}

        mockClient.On("RemoveFlowSource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFlowSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFlowVpcInterface", func(t *testing.T) {
        input := &mediaconnect.RemoveFlowVpcInterfaceInput{}
        output := &mediaconnect.RemoveFlowVpcInterfaceOutput{}

        mockClient.On("RemoveFlowVpcInterface", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFlowVpcInterface(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeFlowEntitlement", func(t *testing.T) {
        input := &mediaconnect.RevokeFlowEntitlementInput{}
        output := &mediaconnect.RevokeFlowEntitlementOutput{}

        mockClient.On("RevokeFlowEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeFlowEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFlow", func(t *testing.T) {
        input := &mediaconnect.StartFlowInput{}
        output := &mediaconnect.StartFlowOutput{}

        mockClient.On("StartFlow", ctx, input).Return(output, nil)

        result, err := mockClient.StartFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopFlow", func(t *testing.T) {
        input := &mediaconnect.StopFlowInput{}
        output := &mediaconnect.StopFlowOutput{}

        mockClient.On("StopFlow", ctx, input).Return(output, nil)

        result, err := mockClient.StopFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediaconnect.TagResourceInput{}
        output := &mediaconnect.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediaconnect.UntagResourceInput{}
        output := &mediaconnect.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBridge", func(t *testing.T) {
        input := &mediaconnect.UpdateBridgeInput{}
        output := &mediaconnect.UpdateBridgeOutput{}

        mockClient.On("UpdateBridge", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBridge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBridgeOutput", func(t *testing.T) {
        input := &mediaconnect.UpdateBridgeOutputInput{}
        output := &mediaconnect.UpdateBridgeOutputOutput{}

        mockClient.On("UpdateBridgeOutput", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBridgeOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBridgeSource", func(t *testing.T) {
        input := &mediaconnect.UpdateBridgeSourceInput{}
        output := &mediaconnect.UpdateBridgeSourceOutput{}

        mockClient.On("UpdateBridgeSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBridgeSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBridgeState", func(t *testing.T) {
        input := &mediaconnect.UpdateBridgeStateInput{}
        output := &mediaconnect.UpdateBridgeStateOutput{}

        mockClient.On("UpdateBridgeState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBridgeState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlow", func(t *testing.T) {
        input := &mediaconnect.UpdateFlowInput{}
        output := &mediaconnect.UpdateFlowOutput{}

        mockClient.On("UpdateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowEntitlement", func(t *testing.T) {
        input := &mediaconnect.UpdateFlowEntitlementInput{}
        output := &mediaconnect.UpdateFlowEntitlementOutput{}

        mockClient.On("UpdateFlowEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowMediaStream", func(t *testing.T) {
        input := &mediaconnect.UpdateFlowMediaStreamInput{}
        output := &mediaconnect.UpdateFlowMediaStreamOutput{}

        mockClient.On("UpdateFlowMediaStream", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowMediaStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowOutput", func(t *testing.T) {
        input := &mediaconnect.UpdateFlowOutputInput{}
        output := &mediaconnect.UpdateFlowOutputOutput{}

        mockClient.On("UpdateFlowOutput", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowSource", func(t *testing.T) {
        input := &mediaconnect.UpdateFlowSourceInput{}
        output := &mediaconnect.UpdateFlowSourceOutput{}

        mockClient.On("UpdateFlowSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayInstance", func(t *testing.T) {
        input := &mediaconnect.UpdateGatewayInstanceInput{}
        output := &mediaconnect.UpdateGatewayInstanceOutput{}

        mockClient.On("UpdateGatewayInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
