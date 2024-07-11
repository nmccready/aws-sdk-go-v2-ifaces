
package docdbelastic

import (
    "github.com/aws/aws-sdk-go-v2/service/docdbelastic"
)

// IClient defines the interface for docdbelastic
type IClient interface {
 Options() Options 
 CopyClusterSnapshot(ctx context.Context, params *CopyClusterSnapshotInput, optFns ...func(*Options)) (*CopyClusterSnapshotOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateClusterSnapshot(ctx context.Context, params *CreateClusterSnapshotInput, optFns ...func(*Options)) (*CreateClusterSnapshotOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteClusterSnapshot(ctx context.Context, params *DeleteClusterSnapshotInput, optFns ...func(*Options)) (*DeleteClusterSnapshotOutput, error) 
 GetCluster(ctx context.Context, params *GetClusterInput, optFns ...func(*Options)) (*GetClusterOutput, error) 
 GetClusterSnapshot(ctx context.Context, params *GetClusterSnapshotInput, optFns ...func(*Options)) (*GetClusterSnapshotOutput, error) 
 ListClusterSnapshots(ctx context.Context, params *ListClusterSnapshotsInput, optFns ...func(*Options)) (*ListClusterSnapshotsOutput, error) 
 ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RestoreClusterFromSnapshot(ctx context.Context, params *RestoreClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreClusterFromSnapshotOutput, error) 
 StartCluster(ctx context.Context, params *StartClusterInput, optFns ...func(*Options)) (*StartClusterOutput, error) 
 StopCluster(ctx context.Context, params *StopClusterInput, optFns ...func(*Options)) (*StopClusterOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) 
}
