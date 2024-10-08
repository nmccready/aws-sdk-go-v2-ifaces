
package globalaccelerator_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

// IClient defines the interface for globalaccelerator
type IClient interface {
 Options() Options 
 AddCustomRoutingEndpoints(ctx context.Context, params *AddCustomRoutingEndpointsInput, optFns ...func(*Options)) (*AddCustomRoutingEndpointsOutput, error) 
 AddEndpoints(ctx context.Context, params *AddEndpointsInput, optFns ...func(*Options)) (*AddEndpointsOutput, error) 
 AdvertiseByoipCidr(ctx context.Context, params *AdvertiseByoipCidrInput, optFns ...func(*Options)) (*AdvertiseByoipCidrOutput, error) 
 AllowCustomRoutingTraffic(ctx context.Context, params *AllowCustomRoutingTrafficInput, optFns ...func(*Options)) (*AllowCustomRoutingTrafficOutput, error) 
 CreateAccelerator(ctx context.Context, params *CreateAcceleratorInput, optFns ...func(*Options)) (*CreateAcceleratorOutput, error) 
 CreateCrossAccountAttachment(ctx context.Context, params *CreateCrossAccountAttachmentInput, optFns ...func(*Options)) (*CreateCrossAccountAttachmentOutput, error) 
 CreateCustomRoutingAccelerator(ctx context.Context, params *CreateCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*CreateCustomRoutingAcceleratorOutput, error) 
 CreateCustomRoutingEndpointGroup(ctx context.Context, params *CreateCustomRoutingEndpointGroupInput, optFns ...func(*Options)) (*CreateCustomRoutingEndpointGroupOutput, error) 
 CreateCustomRoutingListener(ctx context.Context, params *CreateCustomRoutingListenerInput, optFns ...func(*Options)) (*CreateCustomRoutingListenerOutput, error) 
 CreateEndpointGroup(ctx context.Context, params *CreateEndpointGroupInput, optFns ...func(*Options)) (*CreateEndpointGroupOutput, error) 
 CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) 
 DeleteAccelerator(ctx context.Context, params *DeleteAcceleratorInput, optFns ...func(*Options)) (*DeleteAcceleratorOutput, error) 
 DeleteCrossAccountAttachment(ctx context.Context, params *DeleteCrossAccountAttachmentInput, optFns ...func(*Options)) (*DeleteCrossAccountAttachmentOutput, error) 
 DeleteCustomRoutingAccelerator(ctx context.Context, params *DeleteCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*DeleteCustomRoutingAcceleratorOutput, error) 
 DeleteCustomRoutingEndpointGroup(ctx context.Context, params *DeleteCustomRoutingEndpointGroupInput, optFns ...func(*Options)) (*DeleteCustomRoutingEndpointGroupOutput, error) 
 DeleteCustomRoutingListener(ctx context.Context, params *DeleteCustomRoutingListenerInput, optFns ...func(*Options)) (*DeleteCustomRoutingListenerOutput, error) 
 DeleteEndpointGroup(ctx context.Context, params *DeleteEndpointGroupInput, optFns ...func(*Options)) (*DeleteEndpointGroupOutput, error) 
 DeleteListener(ctx context.Context, params *DeleteListenerInput, optFns ...func(*Options)) (*DeleteListenerOutput, error) 
 DenyCustomRoutingTraffic(ctx context.Context, params *DenyCustomRoutingTrafficInput, optFns ...func(*Options)) (*DenyCustomRoutingTrafficOutput, error) 
 DeprovisionByoipCidr(ctx context.Context, params *DeprovisionByoipCidrInput, optFns ...func(*Options)) (*DeprovisionByoipCidrOutput, error) 
 DescribeAccelerator(ctx context.Context, params *DescribeAcceleratorInput, optFns ...func(*Options)) (*DescribeAcceleratorOutput, error) 
 DescribeAcceleratorAttributes(ctx context.Context, params *DescribeAcceleratorAttributesInput, optFns ...func(*Options)) (*DescribeAcceleratorAttributesOutput, error) 
 DescribeCrossAccountAttachment(ctx context.Context, params *DescribeCrossAccountAttachmentInput, optFns ...func(*Options)) (*DescribeCrossAccountAttachmentOutput, error) 
 DescribeCustomRoutingAccelerator(ctx context.Context, params *DescribeCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*DescribeCustomRoutingAcceleratorOutput, error) 
 DescribeCustomRoutingAcceleratorAttributes(ctx context.Context, params *DescribeCustomRoutingAcceleratorAttributesInput, optFns ...func(*Options)) (*DescribeCustomRoutingAcceleratorAttributesOutput, error) 
 DescribeCustomRoutingEndpointGroup(ctx context.Context, params *DescribeCustomRoutingEndpointGroupInput, optFns ...func(*Options)) (*DescribeCustomRoutingEndpointGroupOutput, error) 
 DescribeCustomRoutingListener(ctx context.Context, params *DescribeCustomRoutingListenerInput, optFns ...func(*Options)) (*DescribeCustomRoutingListenerOutput, error) 
 DescribeEndpointGroup(ctx context.Context, params *DescribeEndpointGroupInput, optFns ...func(*Options)) (*DescribeEndpointGroupOutput, error) 
 DescribeListener(ctx context.Context, params *DescribeListenerInput, optFns ...func(*Options)) (*DescribeListenerOutput, error) 
 ListAccelerators(ctx context.Context, params *ListAcceleratorsInput, optFns ...func(*Options)) (*ListAcceleratorsOutput, error) 
 ListByoipCidrs(ctx context.Context, params *ListByoipCidrsInput, optFns ...func(*Options)) (*ListByoipCidrsOutput, error) 
 ListCrossAccountAttachments(ctx context.Context, params *ListCrossAccountAttachmentsInput, optFns ...func(*Options)) (*ListCrossAccountAttachmentsOutput, error) 
 ListCrossAccountResourceAccounts(ctx context.Context, params *ListCrossAccountResourceAccountsInput, optFns ...func(*Options)) (*ListCrossAccountResourceAccountsOutput, error) 
 ListCrossAccountResources(ctx context.Context, params *ListCrossAccountResourcesInput, optFns ...func(*Options)) (*ListCrossAccountResourcesOutput, error) 
 ListCustomRoutingAccelerators(ctx context.Context, params *ListCustomRoutingAcceleratorsInput, optFns ...func(*Options)) (*ListCustomRoutingAcceleratorsOutput, error) 
 ListCustomRoutingEndpointGroups(ctx context.Context, params *ListCustomRoutingEndpointGroupsInput, optFns ...func(*Options)) (*ListCustomRoutingEndpointGroupsOutput, error) 
 ListCustomRoutingListeners(ctx context.Context, params *ListCustomRoutingListenersInput, optFns ...func(*Options)) (*ListCustomRoutingListenersOutput, error) 
 ListCustomRoutingPortMappings(ctx context.Context, params *ListCustomRoutingPortMappingsInput, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsOutput, error) 
 ListCustomRoutingPortMappingsByDestination(ctx context.Context, params *ListCustomRoutingPortMappingsByDestinationInput, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsByDestinationOutput, error) 
 ListEndpointGroups(ctx context.Context, params *ListEndpointGroupsInput, optFns ...func(*Options)) (*ListEndpointGroupsOutput, error) 
 ListListeners(ctx context.Context, params *ListListenersInput, optFns ...func(*Options)) (*ListListenersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ProvisionByoipCidr(ctx context.Context, params *ProvisionByoipCidrInput, optFns ...func(*Options)) (*ProvisionByoipCidrOutput, error) 
 RemoveCustomRoutingEndpoints(ctx context.Context, params *RemoveCustomRoutingEndpointsInput, optFns ...func(*Options)) (*RemoveCustomRoutingEndpointsOutput, error) 
 RemoveEndpoints(ctx context.Context, params *RemoveEndpointsInput, optFns ...func(*Options)) (*RemoveEndpointsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccelerator(ctx context.Context, params *UpdateAcceleratorInput, optFns ...func(*Options)) (*UpdateAcceleratorOutput, error) 
 UpdateAcceleratorAttributes(ctx context.Context, params *UpdateAcceleratorAttributesInput, optFns ...func(*Options)) (*UpdateAcceleratorAttributesOutput, error) 
 UpdateCrossAccountAttachment(ctx context.Context, params *UpdateCrossAccountAttachmentInput, optFns ...func(*Options)) (*UpdateCrossAccountAttachmentOutput, error) 
 UpdateCustomRoutingAccelerator(ctx context.Context, params *UpdateCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*UpdateCustomRoutingAcceleratorOutput, error) 
 UpdateCustomRoutingAcceleratorAttributes(ctx context.Context, params *UpdateCustomRoutingAcceleratorAttributesInput, optFns ...func(*Options)) (*UpdateCustomRoutingAcceleratorAttributesOutput, error) 
 UpdateCustomRoutingListener(ctx context.Context, params *UpdateCustomRoutingListenerInput, optFns ...func(*Options)) (*UpdateCustomRoutingListenerOutput, error) 
 UpdateEndpointGroup(ctx context.Context, params *UpdateEndpointGroupInput, optFns ...func(*Options)) (*UpdateEndpointGroupOutput, error) 
 UpdateListener(ctx context.Context, params *UpdateListenerInput, optFns ...func(*Options)) (*UpdateListenerOutput, error) 
 WithdrawByoipCidr(ctx context.Context, params *WithdrawByoipCidrInput, optFns ...func(*Options)) (*WithdrawByoipCidrOutput, error) 
}
