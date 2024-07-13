
package finspace_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/finspace"
)

// IClient defines the interface for finspace
type IClient interface {
 Options() Options 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 CreateKxChangeset(ctx context.Context, params *CreateKxChangesetInput, optFns ...func(*Options)) (*CreateKxChangesetOutput, error) 
 CreateKxCluster(ctx context.Context, params *CreateKxClusterInput, optFns ...func(*Options)) (*CreateKxClusterOutput, error) 
 CreateKxDatabase(ctx context.Context, params *CreateKxDatabaseInput, optFns ...func(*Options)) (*CreateKxDatabaseOutput, error) 
 CreateKxDataview(ctx context.Context, params *CreateKxDataviewInput, optFns ...func(*Options)) (*CreateKxDataviewOutput, error) 
 CreateKxEnvironment(ctx context.Context, params *CreateKxEnvironmentInput, optFns ...func(*Options)) (*CreateKxEnvironmentOutput, error) 
 CreateKxScalingGroup(ctx context.Context, params *CreateKxScalingGroupInput, optFns ...func(*Options)) (*CreateKxScalingGroupOutput, error) 
 CreateKxUser(ctx context.Context, params *CreateKxUserInput, optFns ...func(*Options)) (*CreateKxUserOutput, error) 
 CreateKxVolume(ctx context.Context, params *CreateKxVolumeInput, optFns ...func(*Options)) (*CreateKxVolumeOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 DeleteKxCluster(ctx context.Context, params *DeleteKxClusterInput, optFns ...func(*Options)) (*DeleteKxClusterOutput, error) 
 DeleteKxClusterNode(ctx context.Context, params *DeleteKxClusterNodeInput, optFns ...func(*Options)) (*DeleteKxClusterNodeOutput, error) 
 DeleteKxDatabase(ctx context.Context, params *DeleteKxDatabaseInput, optFns ...func(*Options)) (*DeleteKxDatabaseOutput, error) 
 DeleteKxDataview(ctx context.Context, params *DeleteKxDataviewInput, optFns ...func(*Options)) (*DeleteKxDataviewOutput, error) 
 DeleteKxEnvironment(ctx context.Context, params *DeleteKxEnvironmentInput, optFns ...func(*Options)) (*DeleteKxEnvironmentOutput, error) 
 DeleteKxScalingGroup(ctx context.Context, params *DeleteKxScalingGroupInput, optFns ...func(*Options)) (*DeleteKxScalingGroupOutput, error) 
 DeleteKxUser(ctx context.Context, params *DeleteKxUserInput, optFns ...func(*Options)) (*DeleteKxUserOutput, error) 
 DeleteKxVolume(ctx context.Context, params *DeleteKxVolumeInput, optFns ...func(*Options)) (*DeleteKxVolumeOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 GetKxChangeset(ctx context.Context, params *GetKxChangesetInput, optFns ...func(*Options)) (*GetKxChangesetOutput, error) 
 GetKxCluster(ctx context.Context, params *GetKxClusterInput, optFns ...func(*Options)) (*GetKxClusterOutput, error) 
 GetKxConnectionString(ctx context.Context, params *GetKxConnectionStringInput, optFns ...func(*Options)) (*GetKxConnectionStringOutput, error) 
 GetKxDatabase(ctx context.Context, params *GetKxDatabaseInput, optFns ...func(*Options)) (*GetKxDatabaseOutput, error) 
 GetKxDataview(ctx context.Context, params *GetKxDataviewInput, optFns ...func(*Options)) (*GetKxDataviewOutput, error) 
 GetKxEnvironment(ctx context.Context, params *GetKxEnvironmentInput, optFns ...func(*Options)) (*GetKxEnvironmentOutput, error) 
 GetKxScalingGroup(ctx context.Context, params *GetKxScalingGroupInput, optFns ...func(*Options)) (*GetKxScalingGroupOutput, error) 
 GetKxUser(ctx context.Context, params *GetKxUserInput, optFns ...func(*Options)) (*GetKxUserOutput, error) 
 GetKxVolume(ctx context.Context, params *GetKxVolumeInput, optFns ...func(*Options)) (*GetKxVolumeOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListKxChangesets(ctx context.Context, params *ListKxChangesetsInput, optFns ...func(*Options)) (*ListKxChangesetsOutput, error) 
 ListKxClusterNodes(ctx context.Context, params *ListKxClusterNodesInput, optFns ...func(*Options)) (*ListKxClusterNodesOutput, error) 
 ListKxClusters(ctx context.Context, params *ListKxClustersInput, optFns ...func(*Options)) (*ListKxClustersOutput, error) 
 ListKxDatabases(ctx context.Context, params *ListKxDatabasesInput, optFns ...func(*Options)) (*ListKxDatabasesOutput, error) 
 ListKxDataviews(ctx context.Context, params *ListKxDataviewsInput, optFns ...func(*Options)) (*ListKxDataviewsOutput, error) 
 ListKxEnvironments(ctx context.Context, params *ListKxEnvironmentsInput, optFns ...func(*Options)) (*ListKxEnvironmentsOutput, error) 
 ListKxScalingGroups(ctx context.Context, params *ListKxScalingGroupsInput, optFns ...func(*Options)) (*ListKxScalingGroupsOutput, error) 
 ListKxUsers(ctx context.Context, params *ListKxUsersInput, optFns ...func(*Options)) (*ListKxUsersOutput, error) 
 ListKxVolumes(ctx context.Context, params *ListKxVolumesInput, optFns ...func(*Options)) (*ListKxVolumesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
 UpdateKxClusterCodeConfiguration(ctx context.Context, params *UpdateKxClusterCodeConfigurationInput, optFns ...func(*Options)) (*UpdateKxClusterCodeConfigurationOutput, error) 
 UpdateKxClusterDatabases(ctx context.Context, params *UpdateKxClusterDatabasesInput, optFns ...func(*Options)) (*UpdateKxClusterDatabasesOutput, error) 
 UpdateKxDatabase(ctx context.Context, params *UpdateKxDatabaseInput, optFns ...func(*Options)) (*UpdateKxDatabaseOutput, error) 
 UpdateKxDataview(ctx context.Context, params *UpdateKxDataviewInput, optFns ...func(*Options)) (*UpdateKxDataviewOutput, error) 
 UpdateKxEnvironment(ctx context.Context, params *UpdateKxEnvironmentInput, optFns ...func(*Options)) (*UpdateKxEnvironmentOutput, error) 
 UpdateKxEnvironmentNetwork(ctx context.Context, params *UpdateKxEnvironmentNetworkInput, optFns ...func(*Options)) (*UpdateKxEnvironmentNetworkOutput, error) 
 UpdateKxUser(ctx context.Context, params *UpdateKxUserInput, optFns ...func(*Options)) (*UpdateKxUserOutput, error) 
 UpdateKxVolume(ctx context.Context, params *UpdateKxVolumeInput, optFns ...func(*Options)) (*UpdateKxVolumeOutput, error) 
}