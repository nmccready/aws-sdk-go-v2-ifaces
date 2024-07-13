
package appmesh_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/appmesh"
)

// IClient defines the interface for appmesh
type IClient interface {
 Options() Options 
 CreateGatewayRoute(ctx context.Context, params *CreateGatewayRouteInput, optFns ...func(*Options)) (*CreateGatewayRouteOutput, error) 
 CreateMesh(ctx context.Context, params *CreateMeshInput, optFns ...func(*Options)) (*CreateMeshOutput, error) 
 CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) 
 CreateVirtualGateway(ctx context.Context, params *CreateVirtualGatewayInput, optFns ...func(*Options)) (*CreateVirtualGatewayOutput, error) 
 CreateVirtualNode(ctx context.Context, params *CreateVirtualNodeInput, optFns ...func(*Options)) (*CreateVirtualNodeOutput, error) 
 CreateVirtualRouter(ctx context.Context, params *CreateVirtualRouterInput, optFns ...func(*Options)) (*CreateVirtualRouterOutput, error) 
 CreateVirtualService(ctx context.Context, params *CreateVirtualServiceInput, optFns ...func(*Options)) (*CreateVirtualServiceOutput, error) 
 DeleteGatewayRoute(ctx context.Context, params *DeleteGatewayRouteInput, optFns ...func(*Options)) (*DeleteGatewayRouteOutput, error) 
 DeleteMesh(ctx context.Context, params *DeleteMeshInput, optFns ...func(*Options)) (*DeleteMeshOutput, error) 
 DeleteRoute(ctx context.Context, params *DeleteRouteInput, optFns ...func(*Options)) (*DeleteRouteOutput, error) 
 DeleteVirtualGateway(ctx context.Context, params *DeleteVirtualGatewayInput, optFns ...func(*Options)) (*DeleteVirtualGatewayOutput, error) 
 DeleteVirtualNode(ctx context.Context, params *DeleteVirtualNodeInput, optFns ...func(*Options)) (*DeleteVirtualNodeOutput, error) 
 DeleteVirtualRouter(ctx context.Context, params *DeleteVirtualRouterInput, optFns ...func(*Options)) (*DeleteVirtualRouterOutput, error) 
 DeleteVirtualService(ctx context.Context, params *DeleteVirtualServiceInput, optFns ...func(*Options)) (*DeleteVirtualServiceOutput, error) 
 DescribeGatewayRoute(ctx context.Context, params *DescribeGatewayRouteInput, optFns ...func(*Options)) (*DescribeGatewayRouteOutput, error) 
 DescribeMesh(ctx context.Context, params *DescribeMeshInput, optFns ...func(*Options)) (*DescribeMeshOutput, error) 
 DescribeRoute(ctx context.Context, params *DescribeRouteInput, optFns ...func(*Options)) (*DescribeRouteOutput, error) 
 DescribeVirtualGateway(ctx context.Context, params *DescribeVirtualGatewayInput, optFns ...func(*Options)) (*DescribeVirtualGatewayOutput, error) 
 DescribeVirtualNode(ctx context.Context, params *DescribeVirtualNodeInput, optFns ...func(*Options)) (*DescribeVirtualNodeOutput, error) 
 DescribeVirtualRouter(ctx context.Context, params *DescribeVirtualRouterInput, optFns ...func(*Options)) (*DescribeVirtualRouterOutput, error) 
 DescribeVirtualService(ctx context.Context, params *DescribeVirtualServiceInput, optFns ...func(*Options)) (*DescribeVirtualServiceOutput, error) 
 ListGatewayRoutes(ctx context.Context, params *ListGatewayRoutesInput, optFns ...func(*Options)) (*ListGatewayRoutesOutput, error) 
 ListMeshes(ctx context.Context, params *ListMeshesInput, optFns ...func(*Options)) (*ListMeshesOutput, error) 
 ListRoutes(ctx context.Context, params *ListRoutesInput, optFns ...func(*Options)) (*ListRoutesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVirtualGateways(ctx context.Context, params *ListVirtualGatewaysInput, optFns ...func(*Options)) (*ListVirtualGatewaysOutput, error) 
 ListVirtualNodes(ctx context.Context, params *ListVirtualNodesInput, optFns ...func(*Options)) (*ListVirtualNodesOutput, error) 
 ListVirtualRouters(ctx context.Context, params *ListVirtualRoutersInput, optFns ...func(*Options)) (*ListVirtualRoutersOutput, error) 
 ListVirtualServices(ctx context.Context, params *ListVirtualServicesInput, optFns ...func(*Options)) (*ListVirtualServicesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGatewayRoute(ctx context.Context, params *UpdateGatewayRouteInput, optFns ...func(*Options)) (*UpdateGatewayRouteOutput, error) 
 UpdateMesh(ctx context.Context, params *UpdateMeshInput, optFns ...func(*Options)) (*UpdateMeshOutput, error) 
 UpdateRoute(ctx context.Context, params *UpdateRouteInput, optFns ...func(*Options)) (*UpdateRouteOutput, error) 
 UpdateVirtualGateway(ctx context.Context, params *UpdateVirtualGatewayInput, optFns ...func(*Options)) (*UpdateVirtualGatewayOutput, error) 
 UpdateVirtualNode(ctx context.Context, params *UpdateVirtualNodeInput, optFns ...func(*Options)) (*UpdateVirtualNodeOutput, error) 
 UpdateVirtualRouter(ctx context.Context, params *UpdateVirtualRouterInput, optFns ...func(*Options)) (*UpdateVirtualRouterOutput, error) 
 UpdateVirtualService(ctx context.Context, params *UpdateVirtualServiceInput, optFns ...func(*Options)) (*UpdateVirtualServiceOutput, error) 
}