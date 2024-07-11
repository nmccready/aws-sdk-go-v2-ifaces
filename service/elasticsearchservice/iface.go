
package elasticsearchservice

import (
    "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

// IElasticsearchservice defines the interface for elasticsearchservice
type IElasticsearchservice interface {
 Options() Options 
 AcceptInboundCrossClusterSearchConnection(ctx context.Context, params *AcceptInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*AcceptInboundCrossClusterSearchConnectionOutput, error) 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 AssociatePackage(ctx context.Context, params *AssociatePackageInput, optFns ...func(*Options)) (*AssociatePackageOutput, error) 
 AuthorizeVpcEndpointAccess(ctx context.Context, params *AuthorizeVpcEndpointAccessInput, optFns ...func(*Options)) (*AuthorizeVpcEndpointAccessOutput, error) 
 CancelDomainConfigChange(ctx context.Context, params *CancelDomainConfigChangeInput, optFns ...func(*Options)) (*CancelDomainConfigChangeOutput, error) 
 CancelElasticsearchServiceSoftwareUpdate(ctx context.Context, params *CancelElasticsearchServiceSoftwareUpdateInput, optFns ...func(*Options)) (*CancelElasticsearchServiceSoftwareUpdateOutput, error) 
 CreateElasticsearchDomain(ctx context.Context, params *CreateElasticsearchDomainInput, optFns ...func(*Options)) (*CreateElasticsearchDomainOutput, error) 
 CreateOutboundCrossClusterSearchConnection(ctx context.Context, params *CreateOutboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*CreateOutboundCrossClusterSearchConnectionOutput, error) 
 CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) 
 CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) 
 DeleteElasticsearchDomain(ctx context.Context, params *DeleteElasticsearchDomainInput, optFns ...func(*Options)) (*DeleteElasticsearchDomainOutput, error) 
 DeleteElasticsearchServiceRole(ctx context.Context, params *DeleteElasticsearchServiceRoleInput, optFns ...func(*Options)) (*DeleteElasticsearchServiceRoleOutput, error) 
 DeleteInboundCrossClusterSearchConnection(ctx context.Context, params *DeleteInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*DeleteInboundCrossClusterSearchConnectionOutput, error) 
 DeleteOutboundCrossClusterSearchConnection(ctx context.Context, params *DeleteOutboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*DeleteOutboundCrossClusterSearchConnectionOutput, error) 
 DeletePackage(ctx context.Context, params *DeletePackageInput, optFns ...func(*Options)) (*DeletePackageOutput, error) 
 DeleteVpcEndpoint(ctx context.Context, params *DeleteVpcEndpointInput, optFns ...func(*Options)) (*DeleteVpcEndpointOutput, error) 
 DescribeDomainAutoTunes(ctx context.Context, params *DescribeDomainAutoTunesInput, optFns ...func(*Options)) (*DescribeDomainAutoTunesOutput, error) 
 DescribeDomainChangeProgress(ctx context.Context, params *DescribeDomainChangeProgressInput, optFns ...func(*Options)) (*DescribeDomainChangeProgressOutput, error) 
 DescribeElasticsearchDomain(ctx context.Context, params *DescribeElasticsearchDomainInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainOutput, error) 
 DescribeElasticsearchDomainConfig(ctx context.Context, params *DescribeElasticsearchDomainConfigInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainConfigOutput, error) 
 DescribeElasticsearchDomains(ctx context.Context, params *DescribeElasticsearchDomainsInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainsOutput, error) 
 DescribeElasticsearchInstanceTypeLimits(ctx context.Context, params *DescribeElasticsearchInstanceTypeLimitsInput, optFns ...func(*Options)) (*DescribeElasticsearchInstanceTypeLimitsOutput, error) 
 DescribeInboundCrossClusterSearchConnections(ctx context.Context, params *DescribeInboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error) 
 DescribeOutboundCrossClusterSearchConnections(ctx context.Context, params *DescribeOutboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error) 
 DescribePackages(ctx context.Context, params *DescribePackagesInput, optFns ...func(*Options)) (*DescribePackagesOutput, error) 
 DescribeReservedElasticsearchInstanceOfferings(ctx context.Context, params *DescribeReservedElasticsearchInstanceOfferingsInput, optFns ...func(*Options)) (*DescribeReservedElasticsearchInstanceOfferingsOutput, error) 
 DescribeReservedElasticsearchInstances(ctx context.Context, params *DescribeReservedElasticsearchInstancesInput, optFns ...func(*Options)) (*DescribeReservedElasticsearchInstancesOutput, error) 
 DescribeVpcEndpoints(ctx context.Context, params *DescribeVpcEndpointsInput, optFns ...func(*Options)) (*DescribeVpcEndpointsOutput, error) 
 DissociatePackage(ctx context.Context, params *DissociatePackageInput, optFns ...func(*Options)) (*DissociatePackageOutput, error) 
 GetCompatibleElasticsearchVersions(ctx context.Context, params *GetCompatibleElasticsearchVersionsInput, optFns ...func(*Options)) (*GetCompatibleElasticsearchVersionsOutput, error) 
 GetPackageVersionHistory(ctx context.Context, params *GetPackageVersionHistoryInput, optFns ...func(*Options)) (*GetPackageVersionHistoryOutput, error) 
 GetUpgradeHistory(ctx context.Context, params *GetUpgradeHistoryInput, optFns ...func(*Options)) (*GetUpgradeHistoryOutput, error) 
 GetUpgradeStatus(ctx context.Context, params *GetUpgradeStatusInput, optFns ...func(*Options)) (*GetUpgradeStatusOutput, error) 
 ListDomainNames(ctx context.Context, params *ListDomainNamesInput, optFns ...func(*Options)) (*ListDomainNamesOutput, error) 
 ListDomainsForPackage(ctx context.Context, params *ListDomainsForPackageInput, optFns ...func(*Options)) (*ListDomainsForPackageOutput, error) 
 ListElasticsearchInstanceTypes(ctx context.Context, params *ListElasticsearchInstanceTypesInput, optFns ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error) 
 ListElasticsearchVersions(ctx context.Context, params *ListElasticsearchVersionsInput, optFns ...func(*Options)) (*ListElasticsearchVersionsOutput, error) 
 ListPackagesForDomain(ctx context.Context, params *ListPackagesForDomainInput, optFns ...func(*Options)) (*ListPackagesForDomainOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 ListVpcEndpointAccess(ctx context.Context, params *ListVpcEndpointAccessInput, optFns ...func(*Options)) (*ListVpcEndpointAccessOutput, error) 
 ListVpcEndpoints(ctx context.Context, params *ListVpcEndpointsInput, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) 
 ListVpcEndpointsForDomain(ctx context.Context, params *ListVpcEndpointsForDomainInput, optFns ...func(*Options)) (*ListVpcEndpointsForDomainOutput, error) 
 PurchaseReservedElasticsearchInstanceOffering(ctx context.Context, params *PurchaseReservedElasticsearchInstanceOfferingInput, optFns ...func(*Options)) (*PurchaseReservedElasticsearchInstanceOfferingOutput, error) 
 RejectInboundCrossClusterSearchConnection(ctx context.Context, params *RejectInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*RejectInboundCrossClusterSearchConnectionOutput, error) 
 RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) 
 RevokeVpcEndpointAccess(ctx context.Context, params *RevokeVpcEndpointAccessInput, optFns ...func(*Options)) (*RevokeVpcEndpointAccessOutput, error) 
 StartElasticsearchServiceSoftwareUpdate(ctx context.Context, params *StartElasticsearchServiceSoftwareUpdateInput, optFns ...func(*Options)) (*StartElasticsearchServiceSoftwareUpdateOutput, error) 
 UpdateElasticsearchDomainConfig(ctx context.Context, params *UpdateElasticsearchDomainConfigInput, optFns ...func(*Options)) (*UpdateElasticsearchDomainConfigOutput, error) 
 UpdatePackage(ctx context.Context, params *UpdatePackageInput, optFns ...func(*Options)) (*UpdatePackageOutput, error) 
 UpdateVpcEndpoint(ctx context.Context, params *UpdateVpcEndpointInput, optFns ...func(*Options)) (*UpdateVpcEndpointOutput, error) 
 UpgradeElasticsearchDomain(ctx context.Context, params *UpgradeElasticsearchDomainInput, optFns ...func(*Options)) (*UpgradeElasticsearchDomainOutput, error) 
}
