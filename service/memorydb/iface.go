
package memorydb

import (
    "github.com/aws/aws-sdk-go-v2/service/memorydb"
)

// IClient defines the interface for memorydb
type IClient interface {
 Options() Options 
 BatchUpdateCluster(ctx context.Context, params *BatchUpdateClusterInput, optFns ...func(*Options)) (*BatchUpdateClusterOutput, error) 
 CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) 
 CreateACL(ctx context.Context, params *CreateACLInput, optFns ...func(*Options)) (*CreateACLOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateParameterGroup(ctx context.Context, params *CreateParameterGroupInput, optFns ...func(*Options)) (*CreateParameterGroupOutput, error) 
 CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) 
 CreateSubnetGroup(ctx context.Context, params *CreateSubnetGroupInput, optFns ...func(*Options)) (*CreateSubnetGroupOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteACL(ctx context.Context, params *DeleteACLInput, optFns ...func(*Options)) (*DeleteACLOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteParameterGroup(ctx context.Context, params *DeleteParameterGroupInput, optFns ...func(*Options)) (*DeleteParameterGroupOutput, error) 
 DeleteSnapshot(ctx context.Context, params *DeleteSnapshotInput, optFns ...func(*Options)) (*DeleteSnapshotOutput, error) 
 DeleteSubnetGroup(ctx context.Context, params *DeleteSubnetGroupInput, optFns ...func(*Options)) (*DeleteSubnetGroupOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DescribeACLs(ctx context.Context, params *DescribeACLsInput, optFns ...func(*Options)) (*DescribeACLsOutput, error) 
 DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) 
 DescribeEngineVersions(ctx context.Context, params *DescribeEngineVersionsInput, optFns ...func(*Options)) (*DescribeEngineVersionsOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeParameterGroups(ctx context.Context, params *DescribeParameterGroupsInput, optFns ...func(*Options)) (*DescribeParameterGroupsOutput, error) 
 DescribeParameters(ctx context.Context, params *DescribeParametersInput, optFns ...func(*Options)) (*DescribeParametersOutput, error) 
 DescribeReservedNodes(ctx context.Context, params *DescribeReservedNodesInput, optFns ...func(*Options)) (*DescribeReservedNodesOutput, error) 
 DescribeReservedNodesOfferings(ctx context.Context, params *DescribeReservedNodesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedNodesOfferingsOutput, error) 
 DescribeServiceUpdates(ctx context.Context, params *DescribeServiceUpdatesInput, optFns ...func(*Options)) (*DescribeServiceUpdatesOutput, error) 
 DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) 
 DescribeSubnetGroups(ctx context.Context, params *DescribeSubnetGroupsInput, optFns ...func(*Options)) (*DescribeSubnetGroupsOutput, error) 
 DescribeUsers(ctx context.Context, params *DescribeUsersInput, optFns ...func(*Options)) (*DescribeUsersOutput, error) 
 FailoverShard(ctx context.Context, params *FailoverShardInput, optFns ...func(*Options)) (*FailoverShardOutput, error) 
 ListAllowedNodeTypeUpdates(ctx context.Context, params *ListAllowedNodeTypeUpdatesInput, optFns ...func(*Options)) (*ListAllowedNodeTypeUpdatesOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 PurchaseReservedNodesOffering(ctx context.Context, params *PurchaseReservedNodesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedNodesOfferingOutput, error) 
 ResetParameterGroup(ctx context.Context, params *ResetParameterGroupInput, optFns ...func(*Options)) (*ResetParameterGroupOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateACL(ctx context.Context, params *UpdateACLInput, optFns ...func(*Options)) (*UpdateACLOutput, error) 
 UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) 
 UpdateParameterGroup(ctx context.Context, params *UpdateParameterGroupInput, optFns ...func(*Options)) (*UpdateParameterGroupOutput, error) 
 UpdateSubnetGroup(ctx context.Context, params *UpdateSubnetGroupInput, optFns ...func(*Options)) (*UpdateSubnetGroupOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
}
