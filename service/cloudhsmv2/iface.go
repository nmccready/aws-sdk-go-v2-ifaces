
package cloudhsmv2

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

// ICloudhsmv2 defines the interface for cloudhsmv2
type ICloudhsmv2 interface {
 Options() Options 
 CopyBackupToRegion(ctx context.Context, params *CopyBackupToRegionInput, optFns ...func(*Options)) (*CopyBackupToRegionOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) 
 DeleteBackup(ctx context.Context, params *DeleteBackupInput, optFns ...func(*Options)) (*DeleteBackupOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteHsm(ctx context.Context, params *DeleteHsmInput, optFns ...func(*Options)) (*DeleteHsmOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) 
 DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 InitializeCluster(ctx context.Context, params *InitializeClusterInput, optFns ...func(*Options)) (*InitializeClusterOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 ModifyBackupAttributes(ctx context.Context, params *ModifyBackupAttributesInput, optFns ...func(*Options)) (*ModifyBackupAttributesOutput, error) 
 ModifyCluster(ctx context.Context, params *ModifyClusterInput, optFns ...func(*Options)) (*ModifyClusterOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 RestoreBackup(ctx context.Context, params *RestoreBackupInput, optFns ...func(*Options)) (*RestoreBackupOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
