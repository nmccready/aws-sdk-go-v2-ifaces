
package privatenetworks

import (
    "github.com/aws/aws-sdk-go-v2/service/privatenetworks"
)

// IClient defines the interface for privatenetworks
type IClient interface {
 Options() Options 
 AcknowledgeOrderReceipt(ctx context.Context, params *AcknowledgeOrderReceiptInput, optFns ...func(*Options)) (*AcknowledgeOrderReceiptOutput, error) 
 ActivateDeviceIdentifier(ctx context.Context, params *ActivateDeviceIdentifierInput, optFns ...func(*Options)) (*ActivateDeviceIdentifierOutput, error) 
 ActivateNetworkSite(ctx context.Context, params *ActivateNetworkSiteInput, optFns ...func(*Options)) (*ActivateNetworkSiteOutput, error) 
 ConfigureAccessPoint(ctx context.Context, params *ConfigureAccessPointInput, optFns ...func(*Options)) (*ConfigureAccessPointOutput, error) 
 CreateNetwork(ctx context.Context, params *CreateNetworkInput, optFns ...func(*Options)) (*CreateNetworkOutput, error) 
 CreateNetworkSite(ctx context.Context, params *CreateNetworkSiteInput, optFns ...func(*Options)) (*CreateNetworkSiteOutput, error) 
 DeactivateDeviceIdentifier(ctx context.Context, params *DeactivateDeviceIdentifierInput, optFns ...func(*Options)) (*DeactivateDeviceIdentifierOutput, error) 
 DeleteNetwork(ctx context.Context, params *DeleteNetworkInput, optFns ...func(*Options)) (*DeleteNetworkOutput, error) 
 DeleteNetworkSite(ctx context.Context, params *DeleteNetworkSiteInput, optFns ...func(*Options)) (*DeleteNetworkSiteOutput, error) 
 GetDeviceIdentifier(ctx context.Context, params *GetDeviceIdentifierInput, optFns ...func(*Options)) (*GetDeviceIdentifierOutput, error) 
 GetNetwork(ctx context.Context, params *GetNetworkInput, optFns ...func(*Options)) (*GetNetworkOutput, error) 
 GetNetworkResource(ctx context.Context, params *GetNetworkResourceInput, optFns ...func(*Options)) (*GetNetworkResourceOutput, error) 
 GetNetworkSite(ctx context.Context, params *GetNetworkSiteInput, optFns ...func(*Options)) (*GetNetworkSiteOutput, error) 
 GetOrder(ctx context.Context, params *GetOrderInput, optFns ...func(*Options)) (*GetOrderOutput, error) 
 ListDeviceIdentifiers(ctx context.Context, params *ListDeviceIdentifiersInput, optFns ...func(*Options)) (*ListDeviceIdentifiersOutput, error) 
 ListNetworkResources(ctx context.Context, params *ListNetworkResourcesInput, optFns ...func(*Options)) (*ListNetworkResourcesOutput, error) 
 ListNetworkSites(ctx context.Context, params *ListNetworkSitesInput, optFns ...func(*Options)) (*ListNetworkSitesOutput, error) 
 ListNetworks(ctx context.Context, params *ListNetworksInput, optFns ...func(*Options)) (*ListNetworksOutput, error) 
 ListOrders(ctx context.Context, params *ListOrdersInput, optFns ...func(*Options)) (*ListOrdersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 Ping(ctx context.Context, params *PingInput, optFns ...func(*Options)) (*PingOutput, error) 
 StartNetworkResourceUpdate(ctx context.Context, params *StartNetworkResourceUpdateInput, optFns ...func(*Options)) (*StartNetworkResourceUpdateOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateNetworkSite(ctx context.Context, params *UpdateNetworkSiteInput, optFns ...func(*Options)) (*UpdateNetworkSiteOutput, error) 
 UpdateNetworkSitePlan(ctx context.Context, params *UpdateNetworkSitePlanInput, optFns ...func(*Options)) (*UpdateNetworkSitePlanOutput, error) 
}
