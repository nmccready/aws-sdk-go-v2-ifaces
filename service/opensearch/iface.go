
package opensearch

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/opensearch"
)

// IClient defines the interface for opensearch
type IClient interface {
 Options() Options 
 AcceptInboundConnection(ctx context.Context, params *AcceptInboundConnectionInput, optFns ...func(*Options)) (*AcceptInboundConnectionOutput, error) 
 AddDataSource(ctx context.Context, params *AddDataSourceInput, optFns ...func(*Options)) (*AddDataSourceOutput, error) 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 AssociatePackage(ctx context.Context, params *AssociatePackageInput, optFns ...func(*Options)) (*AssociatePackageOutput, error) 
 AuthorizeVpcEndpointAccess(ctx context.Context, params *AuthorizeVpcEndpointAccessInput, optFns ...func(*Options)) (*AuthorizeVpcEndpointAccessOutput, error) 
 CancelDomainConfigChange(ctx context.Context, params *CancelDomainConfigChangeInput, optFns ...func(*Options)) (*CancelDomainConfigChangeOutput, error) 
 CancelServiceSoftwareUpdate(ctx context.Context, params *CancelServiceSoftwareUpdateInput, optFns ...func(*Options)) (*CancelServiceSoftwareUpdateOutput, error) 
 CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) 
 CreateOutboundConnection(ctx context.Context, params *CreateOutboundConnectionInput, optFns ...func(*Options)) (*CreateOutboundConnectionOutput, error) 
 CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) 
 CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) 
 DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) 
 DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) 
 DeleteInboundConnection(ctx context.Context, params *DeleteInboundConnectionInput, optFns ...func(*Options)) (*DeleteInboundConnectionOutput, error) 
 DeleteOutboundConnection(ctx context.Context, params *DeleteOutboundConnectionInput, optFns ...func(*Options)) (*DeleteOutboundConnectionOutput, error) 
 DeletePackage(ctx context.Context, params *DeletePackageInput, optFns ...func(*Options)) (*DeletePackageOutput, error) 
 DeleteVpcEndpoint(ctx context.Context, params *DeleteVpcEndpointInput, optFns ...func(*Options)) (*DeleteVpcEndpointOutput, error) 
 DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) 
 DescribeDomainAutoTunes(ctx context.Context, params *DescribeDomainAutoTunesInput, optFns ...func(*Options)) (*DescribeDomainAutoTunesOutput, error) 
 DescribeDomainChangeProgress(ctx context.Context, params *DescribeDomainChangeProgressInput, optFns ...func(*Options)) (*DescribeDomainChangeProgressOutput, error) 
 DescribeDomainConfig(ctx context.Context, params *DescribeDomainConfigInput, optFns ...func(*Options)) (*DescribeDomainConfigOutput, error) 
 DescribeDomainHealth(ctx context.Context, params *DescribeDomainHealthInput, optFns ...func(*Options)) (*DescribeDomainHealthOutput, error) 
 DescribeDomainNodes(ctx context.Context, params *DescribeDomainNodesInput, optFns ...func(*Options)) (*DescribeDomainNodesOutput, error) 
 DescribeDomains(ctx context.Context, params *DescribeDomainsInput, optFns ...func(*Options)) (*DescribeDomainsOutput, error) 
 DescribeDryRunProgress(ctx context.Context, params *DescribeDryRunProgressInput, optFns ...func(*Options)) (*DescribeDryRunProgressOutput, error) 
 DescribeInboundConnections(ctx context.Context, params *DescribeInboundConnectionsInput, optFns ...func(*Options)) (*DescribeInboundConnectionsOutput, error) 
 DescribeInstanceTypeLimits(ctx context.Context, params *DescribeInstanceTypeLimitsInput, optFns ...func(*Options)) (*DescribeInstanceTypeLimitsOutput, error) 
 DescribeOutboundConnections(ctx context.Context, params *DescribeOutboundConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundConnectionsOutput, error) 
 DescribePackages(ctx context.Context, params *DescribePackagesInput, optFns ...func(*Options)) (*DescribePackagesOutput, error) 
 DescribeReservedInstanceOfferings(ctx context.Context, params *DescribeReservedInstanceOfferingsInput, optFns ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error) 
 DescribeReservedInstances(ctx context.Context, params *DescribeReservedInstancesInput, optFns ...func(*Options)) (*DescribeReservedInstancesOutput, error) 
 DescribeVpcEndpoints(ctx context.Context, params *DescribeVpcEndpointsInput, optFns ...func(*Options)) (*DescribeVpcEndpointsOutput, error) 
 DissociatePackage(ctx context.Context, params *DissociatePackageInput, optFns ...func(*Options)) (*DissociatePackageOutput, error) 
 GetCompatibleVersions(ctx context.Context, params *GetCompatibleVersionsInput, optFns ...func(*Options)) (*GetCompatibleVersionsOutput, error) 
 GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) 
 GetDomainMaintenanceStatus(ctx context.Context, params *GetDomainMaintenanceStatusInput, optFns ...func(*Options)) (*GetDomainMaintenanceStatusOutput, error) 
 GetPackageVersionHistory(ctx context.Context, params *GetPackageVersionHistoryInput, optFns ...func(*Options)) (*GetPackageVersionHistoryOutput, error) 
 GetUpgradeHistory(ctx context.Context, params *GetUpgradeHistoryInput, optFns ...func(*Options)) (*GetUpgradeHistoryOutput, error) 
 GetUpgradeStatus(ctx context.Context, params *GetUpgradeStatusInput, optFns ...func(*Options)) (*GetUpgradeStatusOutput, error) 
 ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) 
 ListDomainMaintenances(ctx context.Context, params *ListDomainMaintenancesInput, optFns ...func(*Options)) (*ListDomainMaintenancesOutput, error) 
 ListDomainNames(ctx context.Context, params *ListDomainNamesInput, optFns ...func(*Options)) (*ListDomainNamesOutput, error) 
 ListDomainsForPackage(ctx context.Context, params *ListDomainsForPackageInput, optFns ...func(*Options)) (*ListDomainsForPackageOutput, error) 
 ListInstanceTypeDetails(ctx context.Context, params *ListInstanceTypeDetailsInput, optFns ...func(*Options)) (*ListInstanceTypeDetailsOutput, error) 
 ListPackagesForDomain(ctx context.Context, params *ListPackagesForDomainInput, optFns ...func(*Options)) (*ListPackagesForDomainOutput, error) 
 ListScheduledActions(ctx context.Context, params *ListScheduledActionsInput, optFns ...func(*Options)) (*ListScheduledActionsOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 ListVersions(ctx context.Context, params *ListVersionsInput, optFns ...func(*Options)) (*ListVersionsOutput, error) 
 ListVpcEndpointAccess(ctx context.Context, params *ListVpcEndpointAccessInput, optFns ...func(*Options)) (*ListVpcEndpointAccessOutput, error) 
 ListVpcEndpoints(ctx context.Context, params *ListVpcEndpointsInput, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) 
 ListVpcEndpointsForDomain(ctx context.Context, params *ListVpcEndpointsForDomainInput, optFns ...func(*Options)) (*ListVpcEndpointsForDomainOutput, error) 
 PurchaseReservedInstanceOffering(ctx context.Context, params *PurchaseReservedInstanceOfferingInput, optFns ...func(*Options)) (*PurchaseReservedInstanceOfferingOutput, error) 
 RejectInboundConnection(ctx context.Context, params *RejectInboundConnectionInput, optFns ...func(*Options)) (*RejectInboundConnectionOutput, error) 
 RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) 
 RevokeVpcEndpointAccess(ctx context.Context, params *RevokeVpcEndpointAccessInput, optFns ...func(*Options)) (*RevokeVpcEndpointAccessOutput, error) 
 StartDomainMaintenance(ctx context.Context, params *StartDomainMaintenanceInput, optFns ...func(*Options)) (*StartDomainMaintenanceOutput, error) 
 StartServiceSoftwareUpdate(ctx context.Context, params *StartServiceSoftwareUpdateInput, optFns ...func(*Options)) (*StartServiceSoftwareUpdateOutput, error) 
 UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) 
 UpdateDomainConfig(ctx context.Context, params *UpdateDomainConfigInput, optFns ...func(*Options)) (*UpdateDomainConfigOutput, error) 
 UpdatePackage(ctx context.Context, params *UpdatePackageInput, optFns ...func(*Options)) (*UpdatePackageOutput, error) 
 UpdateScheduledAction(ctx context.Context, params *UpdateScheduledActionInput, optFns ...func(*Options)) (*UpdateScheduledActionOutput, error) 
 UpdateVpcEndpoint(ctx context.Context, params *UpdateVpcEndpointInput, optFns ...func(*Options)) (*UpdateVpcEndpointOutput, error) 
 UpgradeDomain(ctx context.Context, params *UpgradeDomainInput, optFns ...func(*Options)) (*UpgradeDomainOutput, error) 
}
