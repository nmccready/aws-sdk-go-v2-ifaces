
package neptune_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/neptune"
)

// IClient defines the interface for neptune
type IClient interface {
 Options() Options 
 AddRoleToDBCluster(ctx context.Context, params *AddRoleToDBClusterInput, optFns ...func(*Options)) (*AddRoleToDBClusterOutput, error) 
 AddSourceIdentifierToSubscription(ctx context.Context, params *AddSourceIdentifierToSubscriptionInput, optFns ...func(*Options)) (*AddSourceIdentifierToSubscriptionOutput, error) 
 AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) 
 ApplyPendingMaintenanceAction(ctx context.Context, params *ApplyPendingMaintenanceActionInput, optFns ...func(*Options)) (*ApplyPendingMaintenanceActionOutput, error) 
 CopyDBClusterParameterGroup(ctx context.Context, params *CopyDBClusterParameterGroupInput, optFns ...func(*Options)) (*CopyDBClusterParameterGroupOutput, error) 
 CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) 
 CopyDBParameterGroup(ctx context.Context, params *CopyDBParameterGroupInput, optFns ...func(*Options)) (*CopyDBParameterGroupOutput, error) 
 CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) 
 CreateDBClusterEndpoint(ctx context.Context, params *CreateDBClusterEndpointInput, optFns ...func(*Options)) (*CreateDBClusterEndpointOutput, error) 
 CreateDBClusterParameterGroup(ctx context.Context, params *CreateDBClusterParameterGroupInput, optFns ...func(*Options)) (*CreateDBClusterParameterGroupOutput, error) 
 CreateDBClusterSnapshot(ctx context.Context, params *CreateDBClusterSnapshotInput, optFns ...func(*Options)) (*CreateDBClusterSnapshotOutput, error) 
 CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) 
 CreateDBParameterGroup(ctx context.Context, params *CreateDBParameterGroupInput, optFns ...func(*Options)) (*CreateDBParameterGroupOutput, error) 
 CreateDBSubnetGroup(ctx context.Context, params *CreateDBSubnetGroupInput, optFns ...func(*Options)) (*CreateDBSubnetGroupOutput, error) 
 CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) 
 CreateGlobalCluster(ctx context.Context, params *CreateGlobalClusterInput, optFns ...func(*Options)) (*CreateGlobalClusterOutput, error) 
 DeleteDBCluster(ctx context.Context, params *DeleteDBClusterInput, optFns ...func(*Options)) (*DeleteDBClusterOutput, error) 
 DeleteDBClusterEndpoint(ctx context.Context, params *DeleteDBClusterEndpointInput, optFns ...func(*Options)) (*DeleteDBClusterEndpointOutput, error) 
 DeleteDBClusterParameterGroup(ctx context.Context, params *DeleteDBClusterParameterGroupInput, optFns ...func(*Options)) (*DeleteDBClusterParameterGroupOutput, error) 
 DeleteDBClusterSnapshot(ctx context.Context, params *DeleteDBClusterSnapshotInput, optFns ...func(*Options)) (*DeleteDBClusterSnapshotOutput, error) 
 DeleteDBInstance(ctx context.Context, params *DeleteDBInstanceInput, optFns ...func(*Options)) (*DeleteDBInstanceOutput, error) 
 DeleteDBParameterGroup(ctx context.Context, params *DeleteDBParameterGroupInput, optFns ...func(*Options)) (*DeleteDBParameterGroupOutput, error) 
 DeleteDBSubnetGroup(ctx context.Context, params *DeleteDBSubnetGroupInput, optFns ...func(*Options)) (*DeleteDBSubnetGroupOutput, error) 
 DeleteEventSubscription(ctx context.Context, params *DeleteEventSubscriptionInput, optFns ...func(*Options)) (*DeleteEventSubscriptionOutput, error) 
 DeleteGlobalCluster(ctx context.Context, params *DeleteGlobalClusterInput, optFns ...func(*Options)) (*DeleteGlobalClusterOutput, error) 
 DescribeDBClusterEndpoints(ctx context.Context, params *DescribeDBClusterEndpointsInput, optFns ...func(*Options)) (*DescribeDBClusterEndpointsOutput, error) 
 DescribeDBClusterParameterGroups(ctx context.Context, params *DescribeDBClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBClusterParameterGroupsOutput, error) 
 DescribeDBClusterParameters(ctx context.Context, params *DescribeDBClusterParametersInput, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) 
 DescribeDBClusterSnapshotAttributes(ctx context.Context, params *DescribeDBClusterSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotAttributesOutput, error) 
 DescribeDBClusterSnapshots(ctx context.Context, params *DescribeDBClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) 
 DescribeDBClusters(ctx context.Context, params *DescribeDBClustersInput, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) 
 DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) 
 DescribeDBInstances(ctx context.Context, params *DescribeDBInstancesInput, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) 
 DescribeDBParameterGroups(ctx context.Context, params *DescribeDBParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBParameterGroupsOutput, error) 
 DescribeDBParameters(ctx context.Context, params *DescribeDBParametersInput, optFns ...func(*Options)) (*DescribeDBParametersOutput, error) 
 DescribeDBSubnetGroups(ctx context.Context, params *DescribeDBSubnetGroupsInput, optFns ...func(*Options)) (*DescribeDBSubnetGroupsOutput, error) 
 DescribeEngineDefaultClusterParameters(ctx context.Context, params *DescribeEngineDefaultClusterParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultClusterParametersOutput, error) 
 DescribeEngineDefaultParameters(ctx context.Context, params *DescribeEngineDefaultParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error) 
 DescribeEventCategories(ctx context.Context, params *DescribeEventCategoriesInput, optFns ...func(*Options)) (*DescribeEventCategoriesOutput, error) 
 DescribeEventSubscriptions(ctx context.Context, params *DescribeEventSubscriptionsInput, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeGlobalClusters(ctx context.Context, params *DescribeGlobalClustersInput, optFns ...func(*Options)) (*DescribeGlobalClustersOutput, error) 
 DescribeOrderableDBInstanceOptions(ctx context.Context, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) 
 DescribePendingMaintenanceActions(ctx context.Context, params *DescribePendingMaintenanceActionsInput, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) 
 DescribeValidDBInstanceModifications(ctx context.Context, params *DescribeValidDBInstanceModificationsInput, optFns ...func(*Options)) (*DescribeValidDBInstanceModificationsOutput, error) 
 FailoverDBCluster(ctx context.Context, params *FailoverDBClusterInput, optFns ...func(*Options)) (*FailoverDBClusterOutput, error) 
 FailoverGlobalCluster(ctx context.Context, params *FailoverGlobalClusterInput, optFns ...func(*Options)) (*FailoverGlobalClusterOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) 
 ModifyDBClusterEndpoint(ctx context.Context, params *ModifyDBClusterEndpointInput, optFns ...func(*Options)) (*ModifyDBClusterEndpointOutput, error) 
 ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) 
 ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) 
 ModifyDBInstance(ctx context.Context, params *ModifyDBInstanceInput, optFns ...func(*Options)) (*ModifyDBInstanceOutput, error) 
 ModifyDBParameterGroup(ctx context.Context, params *ModifyDBParameterGroupInput, optFns ...func(*Options)) (*ModifyDBParameterGroupOutput, error) 
 ModifyDBSubnetGroup(ctx context.Context, params *ModifyDBSubnetGroupInput, optFns ...func(*Options)) (*ModifyDBSubnetGroupOutput, error) 
 ModifyEventSubscription(ctx context.Context, params *ModifyEventSubscriptionInput, optFns ...func(*Options)) (*ModifyEventSubscriptionOutput, error) 
 ModifyGlobalCluster(ctx context.Context, params *ModifyGlobalClusterInput, optFns ...func(*Options)) (*ModifyGlobalClusterOutput, error) 
 PromoteReadReplicaDBCluster(ctx context.Context, params *PromoteReadReplicaDBClusterInput, optFns ...func(*Options)) (*PromoteReadReplicaDBClusterOutput, error) 
 RebootDBInstance(ctx context.Context, params *RebootDBInstanceInput, optFns ...func(*Options)) (*RebootDBInstanceOutput, error) 
 RemoveFromGlobalCluster(ctx context.Context, params *RemoveFromGlobalClusterInput, optFns ...func(*Options)) (*RemoveFromGlobalClusterOutput, error) 
 RemoveRoleFromDBCluster(ctx context.Context, params *RemoveRoleFromDBClusterInput, optFns ...func(*Options)) (*RemoveRoleFromDBClusterOutput, error) 
 RemoveSourceIdentifierFromSubscription(ctx context.Context, params *RemoveSourceIdentifierFromSubscriptionInput, optFns ...func(*Options)) (*RemoveSourceIdentifierFromSubscriptionOutput, error) 
 RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) 
 ResetDBClusterParameterGroup(ctx context.Context, params *ResetDBClusterParameterGroupInput, optFns ...func(*Options)) (*ResetDBClusterParameterGroupOutput, error) 
 ResetDBParameterGroup(ctx context.Context, params *ResetDBParameterGroupInput, optFns ...func(*Options)) (*ResetDBParameterGroupOutput, error) 
 RestoreDBClusterFromSnapshot(ctx context.Context, params *RestoreDBClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreDBClusterFromSnapshotOutput, error) 
 RestoreDBClusterToPointInTime(ctx context.Context, params *RestoreDBClusterToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBClusterToPointInTimeOutput, error) 
 StartDBCluster(ctx context.Context, params *StartDBClusterInput, optFns ...func(*Options)) (*StartDBClusterOutput, error) 
 StopDBCluster(ctx context.Context, params *StopDBClusterInput, optFns ...func(*Options)) (*StopDBClusterOutput, error) 
 SwitchoverGlobalCluster(ctx context.Context, params *SwitchoverGlobalClusterInput, optFns ...func(*Options)) (*SwitchoverGlobalClusterOutput, error) 
}
