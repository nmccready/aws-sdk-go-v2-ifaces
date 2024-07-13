package globalaccelerator_test

// tests for the globalaccelerator service interface for this ../../../service/globalaccelerator/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/globalaccelerator/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/globalaccelerator/globalaccelerator_iface"
	"github.com/stretchr/testify/assert"
)

func TestGlobalacceleratorServiceCanBeMocked(t *testing.T) {
	var iface globalaccelerator_iface.IClient
	iface = &globalaccelerator.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := globalaccelerator.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddCustomRoutingEndpoints", func(t *testing.T) {
        input := &globalaccelerator.AddCustomRoutingEndpointsInput{}
        output := &globalaccelerator.AddCustomRoutingEndpointsOutput{}

        mockClient.On("AddCustomRoutingEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.AddCustomRoutingEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddEndpoints", func(t *testing.T) {
        input := &globalaccelerator.AddEndpointsInput{}
        output := &globalaccelerator.AddEndpointsOutput{}

        mockClient.On("AddEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.AddEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdvertiseByoipCidr", func(t *testing.T) {
        input := &globalaccelerator.AdvertiseByoipCidrInput{}
        output := &globalaccelerator.AdvertiseByoipCidrOutput{}

        mockClient.On("AdvertiseByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.AdvertiseByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllowCustomRoutingTraffic", func(t *testing.T) {
        input := &globalaccelerator.AllowCustomRoutingTrafficInput{}
        output := &globalaccelerator.AllowCustomRoutingTrafficOutput{}

        mockClient.On("AllowCustomRoutingTraffic", ctx, input).Return(output, nil)

        result, err := mockClient.AllowCustomRoutingTraffic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccelerator", func(t *testing.T) {
        input := &globalaccelerator.CreateAcceleratorInput{}
        output := &globalaccelerator.CreateAcceleratorOutput{}

        mockClient.On("CreateAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCrossAccountAttachment", func(t *testing.T) {
        input := &globalaccelerator.CreateCrossAccountAttachmentInput{}
        output := &globalaccelerator.CreateCrossAccountAttachmentOutput{}

        mockClient.On("CreateCrossAccountAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCrossAccountAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomRoutingAccelerator", func(t *testing.T) {
        input := &globalaccelerator.CreateCustomRoutingAcceleratorInput{}
        output := &globalaccelerator.CreateCustomRoutingAcceleratorOutput{}

        mockClient.On("CreateCustomRoutingAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomRoutingAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomRoutingEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.CreateCustomRoutingEndpointGroupInput{}
        output := &globalaccelerator.CreateCustomRoutingEndpointGroupOutput{}

        mockClient.On("CreateCustomRoutingEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomRoutingEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomRoutingListener", func(t *testing.T) {
        input := &globalaccelerator.CreateCustomRoutingListenerInput{}
        output := &globalaccelerator.CreateCustomRoutingListenerOutput{}

        mockClient.On("CreateCustomRoutingListener", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomRoutingListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.CreateEndpointGroupInput{}
        output := &globalaccelerator.CreateEndpointGroupOutput{}

        mockClient.On("CreateEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateListener", func(t *testing.T) {
        input := &globalaccelerator.CreateListenerInput{}
        output := &globalaccelerator.CreateListenerOutput{}

        mockClient.On("CreateListener", ctx, input).Return(output, nil)

        result, err := mockClient.CreateListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccelerator", func(t *testing.T) {
        input := &globalaccelerator.DeleteAcceleratorInput{}
        output := &globalaccelerator.DeleteAcceleratorOutput{}

        mockClient.On("DeleteAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCrossAccountAttachment", func(t *testing.T) {
        input := &globalaccelerator.DeleteCrossAccountAttachmentInput{}
        output := &globalaccelerator.DeleteCrossAccountAttachmentOutput{}

        mockClient.On("DeleteCrossAccountAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCrossAccountAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomRoutingAccelerator", func(t *testing.T) {
        input := &globalaccelerator.DeleteCustomRoutingAcceleratorInput{}
        output := &globalaccelerator.DeleteCustomRoutingAcceleratorOutput{}

        mockClient.On("DeleteCustomRoutingAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomRoutingAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomRoutingEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.DeleteCustomRoutingEndpointGroupInput{}
        output := &globalaccelerator.DeleteCustomRoutingEndpointGroupOutput{}

        mockClient.On("DeleteCustomRoutingEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomRoutingEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomRoutingListener", func(t *testing.T) {
        input := &globalaccelerator.DeleteCustomRoutingListenerInput{}
        output := &globalaccelerator.DeleteCustomRoutingListenerOutput{}

        mockClient.On("DeleteCustomRoutingListener", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomRoutingListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.DeleteEndpointGroupInput{}
        output := &globalaccelerator.DeleteEndpointGroupOutput{}

        mockClient.On("DeleteEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteListener", func(t *testing.T) {
        input := &globalaccelerator.DeleteListenerInput{}
        output := &globalaccelerator.DeleteListenerOutput{}

        mockClient.On("DeleteListener", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDenyCustomRoutingTraffic", func(t *testing.T) {
        input := &globalaccelerator.DenyCustomRoutingTrafficInput{}
        output := &globalaccelerator.DenyCustomRoutingTrafficOutput{}

        mockClient.On("DenyCustomRoutingTraffic", ctx, input).Return(output, nil)

        result, err := mockClient.DenyCustomRoutingTraffic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprovisionByoipCidr", func(t *testing.T) {
        input := &globalaccelerator.DeprovisionByoipCidrInput{}
        output := &globalaccelerator.DeprovisionByoipCidrOutput{}

        mockClient.On("DeprovisionByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.DeprovisionByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccelerator", func(t *testing.T) {
        input := &globalaccelerator.DescribeAcceleratorInput{}
        output := &globalaccelerator.DescribeAcceleratorOutput{}

        mockClient.On("DescribeAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAcceleratorAttributes", func(t *testing.T) {
        input := &globalaccelerator.DescribeAcceleratorAttributesInput{}
        output := &globalaccelerator.DescribeAcceleratorAttributesOutput{}

        mockClient.On("DescribeAcceleratorAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAcceleratorAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCrossAccountAttachment", func(t *testing.T) {
        input := &globalaccelerator.DescribeCrossAccountAttachmentInput{}
        output := &globalaccelerator.DescribeCrossAccountAttachmentOutput{}

        mockClient.On("DescribeCrossAccountAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCrossAccountAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomRoutingAccelerator", func(t *testing.T) {
        input := &globalaccelerator.DescribeCustomRoutingAcceleratorInput{}
        output := &globalaccelerator.DescribeCustomRoutingAcceleratorOutput{}

        mockClient.On("DescribeCustomRoutingAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomRoutingAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomRoutingAcceleratorAttributes", func(t *testing.T) {
        input := &globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput{}
        output := &globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput{}

        mockClient.On("DescribeCustomRoutingAcceleratorAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomRoutingAcceleratorAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomRoutingEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.DescribeCustomRoutingEndpointGroupInput{}
        output := &globalaccelerator.DescribeCustomRoutingEndpointGroupOutput{}

        mockClient.On("DescribeCustomRoutingEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomRoutingEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomRoutingListener", func(t *testing.T) {
        input := &globalaccelerator.DescribeCustomRoutingListenerInput{}
        output := &globalaccelerator.DescribeCustomRoutingListenerOutput{}

        mockClient.On("DescribeCustomRoutingListener", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomRoutingListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.DescribeEndpointGroupInput{}
        output := &globalaccelerator.DescribeEndpointGroupOutput{}

        mockClient.On("DescribeEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeListener", func(t *testing.T) {
        input := &globalaccelerator.DescribeListenerInput{}
        output := &globalaccelerator.DescribeListenerOutput{}

        mockClient.On("DescribeListener", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccelerators", func(t *testing.T) {
        input := &globalaccelerator.ListAcceleratorsInput{}
        output := &globalaccelerator.ListAcceleratorsOutput{}

        mockClient.On("ListAccelerators", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccelerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListByoipCidrs", func(t *testing.T) {
        input := &globalaccelerator.ListByoipCidrsInput{}
        output := &globalaccelerator.ListByoipCidrsOutput{}

        mockClient.On("ListByoipCidrs", ctx, input).Return(output, nil)

        result, err := mockClient.ListByoipCidrs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrossAccountAttachments", func(t *testing.T) {
        input := &globalaccelerator.ListCrossAccountAttachmentsInput{}
        output := &globalaccelerator.ListCrossAccountAttachmentsOutput{}

        mockClient.On("ListCrossAccountAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrossAccountAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrossAccountResourceAccounts", func(t *testing.T) {
        input := &globalaccelerator.ListCrossAccountResourceAccountsInput{}
        output := &globalaccelerator.ListCrossAccountResourceAccountsOutput{}

        mockClient.On("ListCrossAccountResourceAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrossAccountResourceAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrossAccountResources", func(t *testing.T) {
        input := &globalaccelerator.ListCrossAccountResourcesInput{}
        output := &globalaccelerator.ListCrossAccountResourcesOutput{}

        mockClient.On("ListCrossAccountResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrossAccountResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomRoutingAccelerators", func(t *testing.T) {
        input := &globalaccelerator.ListCustomRoutingAcceleratorsInput{}
        output := &globalaccelerator.ListCustomRoutingAcceleratorsOutput{}

        mockClient.On("ListCustomRoutingAccelerators", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomRoutingAccelerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomRoutingEndpointGroups", func(t *testing.T) {
        input := &globalaccelerator.ListCustomRoutingEndpointGroupsInput{}
        output := &globalaccelerator.ListCustomRoutingEndpointGroupsOutput{}

        mockClient.On("ListCustomRoutingEndpointGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomRoutingEndpointGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomRoutingListeners", func(t *testing.T) {
        input := &globalaccelerator.ListCustomRoutingListenersInput{}
        output := &globalaccelerator.ListCustomRoutingListenersOutput{}

        mockClient.On("ListCustomRoutingListeners", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomRoutingListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomRoutingPortMappings", func(t *testing.T) {
        input := &globalaccelerator.ListCustomRoutingPortMappingsInput{}
        output := &globalaccelerator.ListCustomRoutingPortMappingsOutput{}

        mockClient.On("ListCustomRoutingPortMappings", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomRoutingPortMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomRoutingPortMappingsByDestination", func(t *testing.T) {
        input := &globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput{}
        output := &globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput{}

        mockClient.On("ListCustomRoutingPortMappingsByDestination", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomRoutingPortMappingsByDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpointGroups", func(t *testing.T) {
        input := &globalaccelerator.ListEndpointGroupsInput{}
        output := &globalaccelerator.ListEndpointGroupsOutput{}

        mockClient.On("ListEndpointGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpointGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListListeners", func(t *testing.T) {
        input := &globalaccelerator.ListListenersInput{}
        output := &globalaccelerator.ListListenersOutput{}

        mockClient.On("ListListeners", ctx, input).Return(output, nil)

        result, err := mockClient.ListListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &globalaccelerator.ListTagsForResourceInput{}
        output := &globalaccelerator.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionByoipCidr", func(t *testing.T) {
        input := &globalaccelerator.ProvisionByoipCidrInput{}
        output := &globalaccelerator.ProvisionByoipCidrOutput{}

        mockClient.On("ProvisionByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveCustomRoutingEndpoints", func(t *testing.T) {
        input := &globalaccelerator.RemoveCustomRoutingEndpointsInput{}
        output := &globalaccelerator.RemoveCustomRoutingEndpointsOutput{}

        mockClient.On("RemoveCustomRoutingEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveCustomRoutingEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveEndpoints", func(t *testing.T) {
        input := &globalaccelerator.RemoveEndpointsInput{}
        output := &globalaccelerator.RemoveEndpointsOutput{}

        mockClient.On("RemoveEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &globalaccelerator.TagResourceInput{}
        output := &globalaccelerator.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &globalaccelerator.UntagResourceInput{}
        output := &globalaccelerator.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccelerator", func(t *testing.T) {
        input := &globalaccelerator.UpdateAcceleratorInput{}
        output := &globalaccelerator.UpdateAcceleratorOutput{}

        mockClient.On("UpdateAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAcceleratorAttributes", func(t *testing.T) {
        input := &globalaccelerator.UpdateAcceleratorAttributesInput{}
        output := &globalaccelerator.UpdateAcceleratorAttributesOutput{}

        mockClient.On("UpdateAcceleratorAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAcceleratorAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCrossAccountAttachment", func(t *testing.T) {
        input := &globalaccelerator.UpdateCrossAccountAttachmentInput{}
        output := &globalaccelerator.UpdateCrossAccountAttachmentOutput{}

        mockClient.On("UpdateCrossAccountAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCrossAccountAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomRoutingAccelerator", func(t *testing.T) {
        input := &globalaccelerator.UpdateCustomRoutingAcceleratorInput{}
        output := &globalaccelerator.UpdateCustomRoutingAcceleratorOutput{}

        mockClient.On("UpdateCustomRoutingAccelerator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomRoutingAccelerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomRoutingAcceleratorAttributes", func(t *testing.T) {
        input := &globalaccelerator.UpdateCustomRoutingAcceleratorAttributesInput{}
        output := &globalaccelerator.UpdateCustomRoutingAcceleratorAttributesOutput{}

        mockClient.On("UpdateCustomRoutingAcceleratorAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomRoutingAcceleratorAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomRoutingListener", func(t *testing.T) {
        input := &globalaccelerator.UpdateCustomRoutingListenerInput{}
        output := &globalaccelerator.UpdateCustomRoutingListenerOutput{}

        mockClient.On("UpdateCustomRoutingListener", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomRoutingListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpointGroup", func(t *testing.T) {
        input := &globalaccelerator.UpdateEndpointGroupInput{}
        output := &globalaccelerator.UpdateEndpointGroupOutput{}

        mockClient.On("UpdateEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateListener", func(t *testing.T) {
        input := &globalaccelerator.UpdateListenerInput{}
        output := &globalaccelerator.UpdateListenerOutput{}

        mockClient.On("UpdateListener", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestWithdrawByoipCidr", func(t *testing.T) {
        input := &globalaccelerator.WithdrawByoipCidrInput{}
        output := &globalaccelerator.WithdrawByoipCidrOutput{}

        mockClient.On("WithdrawByoipCidr", ctx, input).Return(output, nil)

        result, err := mockClient.WithdrawByoipCidr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
