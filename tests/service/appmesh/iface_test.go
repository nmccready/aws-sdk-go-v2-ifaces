package appmesh_test

// tests for the appmesh service interface for this ../../../service/appmesh/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appmesh/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appmesh/appmesh_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppmeshServiceCanBeMocked(t *testing.T) {
	var iface appmesh_iface.IClient
	iface = &appmesh.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appmesh.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGatewayRoute", func(t *testing.T) {
        input := &appmesh.CreateGatewayRouteInput{}
        output := &appmesh.CreateGatewayRouteOutput{}

        mockClient.On("CreateGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMesh", func(t *testing.T) {
        input := &appmesh.CreateMeshInput{}
        output := &appmesh.CreateMeshOutput{}

        mockClient.On("CreateMesh", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMesh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoute", func(t *testing.T) {
        input := &appmesh.CreateRouteInput{}
        output := &appmesh.CreateRouteOutput{}

        mockClient.On("CreateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualGateway", func(t *testing.T) {
        input := &appmesh.CreateVirtualGatewayInput{}
        output := &appmesh.CreateVirtualGatewayOutput{}

        mockClient.On("CreateVirtualGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualNode", func(t *testing.T) {
        input := &appmesh.CreateVirtualNodeInput{}
        output := &appmesh.CreateVirtualNodeOutput{}

        mockClient.On("CreateVirtualNode", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualRouter", func(t *testing.T) {
        input := &appmesh.CreateVirtualRouterInput{}
        output := &appmesh.CreateVirtualRouterOutput{}

        mockClient.On("CreateVirtualRouter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualRouter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualService", func(t *testing.T) {
        input := &appmesh.CreateVirtualServiceInput{}
        output := &appmesh.CreateVirtualServiceOutput{}

        mockClient.On("CreateVirtualService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGatewayRoute", func(t *testing.T) {
        input := &appmesh.DeleteGatewayRouteInput{}
        output := &appmesh.DeleteGatewayRouteOutput{}

        mockClient.On("DeleteGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMesh", func(t *testing.T) {
        input := &appmesh.DeleteMeshInput{}
        output := &appmesh.DeleteMeshOutput{}

        mockClient.On("DeleteMesh", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMesh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoute", func(t *testing.T) {
        input := &appmesh.DeleteRouteInput{}
        output := &appmesh.DeleteRouteOutput{}

        mockClient.On("DeleteRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualGateway", func(t *testing.T) {
        input := &appmesh.DeleteVirtualGatewayInput{}
        output := &appmesh.DeleteVirtualGatewayOutput{}

        mockClient.On("DeleteVirtualGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualNode", func(t *testing.T) {
        input := &appmesh.DeleteVirtualNodeInput{}
        output := &appmesh.DeleteVirtualNodeOutput{}

        mockClient.On("DeleteVirtualNode", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualRouter", func(t *testing.T) {
        input := &appmesh.DeleteVirtualRouterInput{}
        output := &appmesh.DeleteVirtualRouterOutput{}

        mockClient.On("DeleteVirtualRouter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualRouter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualService", func(t *testing.T) {
        input := &appmesh.DeleteVirtualServiceInput{}
        output := &appmesh.DeleteVirtualServiceOutput{}

        mockClient.On("DeleteVirtualService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGatewayRoute", func(t *testing.T) {
        input := &appmesh.DescribeGatewayRouteInput{}
        output := &appmesh.DescribeGatewayRouteOutput{}

        mockClient.On("DescribeGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMesh", func(t *testing.T) {
        input := &appmesh.DescribeMeshInput{}
        output := &appmesh.DescribeMeshOutput{}

        mockClient.On("DescribeMesh", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMesh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRoute", func(t *testing.T) {
        input := &appmesh.DescribeRouteInput{}
        output := &appmesh.DescribeRouteOutput{}

        mockClient.On("DescribeRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualGateway", func(t *testing.T) {
        input := &appmesh.DescribeVirtualGatewayInput{}
        output := &appmesh.DescribeVirtualGatewayOutput{}

        mockClient.On("DescribeVirtualGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualNode", func(t *testing.T) {
        input := &appmesh.DescribeVirtualNodeInput{}
        output := &appmesh.DescribeVirtualNodeOutput{}

        mockClient.On("DescribeVirtualNode", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualRouter", func(t *testing.T) {
        input := &appmesh.DescribeVirtualRouterInput{}
        output := &appmesh.DescribeVirtualRouterOutput{}

        mockClient.On("DescribeVirtualRouter", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualRouter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualService", func(t *testing.T) {
        input := &appmesh.DescribeVirtualServiceInput{}
        output := &appmesh.DescribeVirtualServiceOutput{}

        mockClient.On("DescribeVirtualService", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGatewayRoutes", func(t *testing.T) {
        input := &appmesh.ListGatewayRoutesInput{}
        output := &appmesh.ListGatewayRoutesOutput{}

        mockClient.On("ListGatewayRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.ListGatewayRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMeshes", func(t *testing.T) {
        input := &appmesh.ListMeshesInput{}
        output := &appmesh.ListMeshesOutput{}

        mockClient.On("ListMeshes", ctx, input).Return(output, nil)

        result, err := mockClient.ListMeshes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutes", func(t *testing.T) {
        input := &appmesh.ListRoutesInput{}
        output := &appmesh.ListRoutesOutput{}

        mockClient.On("ListRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appmesh.ListTagsForResourceInput{}
        output := &appmesh.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualGateways", func(t *testing.T) {
        input := &appmesh.ListVirtualGatewaysInput{}
        output := &appmesh.ListVirtualGatewaysOutput{}

        mockClient.On("ListVirtualGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualNodes", func(t *testing.T) {
        input := &appmesh.ListVirtualNodesInput{}
        output := &appmesh.ListVirtualNodesOutput{}

        mockClient.On("ListVirtualNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualRouters", func(t *testing.T) {
        input := &appmesh.ListVirtualRoutersInput{}
        output := &appmesh.ListVirtualRoutersOutput{}

        mockClient.On("ListVirtualRouters", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualRouters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualServices", func(t *testing.T) {
        input := &appmesh.ListVirtualServicesInput{}
        output := &appmesh.ListVirtualServicesOutput{}

        mockClient.On("ListVirtualServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appmesh.TagResourceInput{}
        output := &appmesh.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appmesh.UntagResourceInput{}
        output := &appmesh.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayRoute", func(t *testing.T) {
        input := &appmesh.UpdateGatewayRouteInput{}
        output := &appmesh.UpdateGatewayRouteOutput{}

        mockClient.On("UpdateGatewayRoute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMesh", func(t *testing.T) {
        input := &appmesh.UpdateMeshInput{}
        output := &appmesh.UpdateMeshOutput{}

        mockClient.On("UpdateMesh", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMesh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoute", func(t *testing.T) {
        input := &appmesh.UpdateRouteInput{}
        output := &appmesh.UpdateRouteOutput{}

        mockClient.On("UpdateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVirtualGateway", func(t *testing.T) {
        input := &appmesh.UpdateVirtualGatewayInput{}
        output := &appmesh.UpdateVirtualGatewayOutput{}

        mockClient.On("UpdateVirtualGateway", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVirtualGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVirtualNode", func(t *testing.T) {
        input := &appmesh.UpdateVirtualNodeInput{}
        output := &appmesh.UpdateVirtualNodeOutput{}

        mockClient.On("UpdateVirtualNode", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVirtualNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVirtualRouter", func(t *testing.T) {
        input := &appmesh.UpdateVirtualRouterInput{}
        output := &appmesh.UpdateVirtualRouterOutput{}

        mockClient.On("UpdateVirtualRouter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVirtualRouter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVirtualService", func(t *testing.T) {
        input := &appmesh.UpdateVirtualServiceInput{}
        output := &appmesh.UpdateVirtualServiceOutput{}

        mockClient.On("UpdateVirtualService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVirtualService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
