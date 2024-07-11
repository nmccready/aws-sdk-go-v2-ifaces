
package efs

import (
    "github.com/aws/aws-sdk-go-v2/service/efs"
)

// IClient defines the interface for efs
type IClient interface {
 Options() Options 
 CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) 
 CreateFileSystem(ctx context.Context, params *CreateFileSystemInput, optFns ...func(*Options)) (*CreateFileSystemOutput, error) 
 CreateMountTarget(ctx context.Context, params *CreateMountTargetInput, optFns ...func(*Options)) (*CreateMountTargetOutput, error) 
 CreateReplicationConfiguration(ctx context.Context, params *CreateReplicationConfigurationInput, optFns ...func(*Options)) (*CreateReplicationConfigurationOutput, error) 
 CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) 
 DeleteAccessPoint(ctx context.Context, params *DeleteAccessPointInput, optFns ...func(*Options)) (*DeleteAccessPointOutput, error) 
 DeleteFileSystem(ctx context.Context, params *DeleteFileSystemInput, optFns ...func(*Options)) (*DeleteFileSystemOutput, error) 
 DeleteFileSystemPolicy(ctx context.Context, params *DeleteFileSystemPolicyInput, optFns ...func(*Options)) (*DeleteFileSystemPolicyOutput, error) 
 DeleteMountTarget(ctx context.Context, params *DeleteMountTargetInput, optFns ...func(*Options)) (*DeleteMountTargetOutput, error) 
 DeleteReplicationConfiguration(ctx context.Context, params *DeleteReplicationConfigurationInput, optFns ...func(*Options)) (*DeleteReplicationConfigurationOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DescribeAccessPoints(ctx context.Context, params *DescribeAccessPointsInput, optFns ...func(*Options)) (*DescribeAccessPointsOutput, error) 
 DescribeAccountPreferences(ctx context.Context, params *DescribeAccountPreferencesInput, optFns ...func(*Options)) (*DescribeAccountPreferencesOutput, error) 
 DescribeBackupPolicy(ctx context.Context, params *DescribeBackupPolicyInput, optFns ...func(*Options)) (*DescribeBackupPolicyOutput, error) 
 DescribeFileSystemPolicy(ctx context.Context, params *DescribeFileSystemPolicyInput, optFns ...func(*Options)) (*DescribeFileSystemPolicyOutput, error) 
 DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) 
 DescribeLifecycleConfiguration(ctx context.Context, params *DescribeLifecycleConfigurationInput, optFns ...func(*Options)) (*DescribeLifecycleConfigurationOutput, error) 
 DescribeMountTargetSecurityGroups(ctx context.Context, params *DescribeMountTargetSecurityGroupsInput, optFns ...func(*Options)) (*DescribeMountTargetSecurityGroupsOutput, error) 
 DescribeMountTargets(ctx context.Context, params *DescribeMountTargetsInput, optFns ...func(*Options)) (*DescribeMountTargetsOutput, error) 
 DescribeReplicationConfigurations(ctx context.Context, params *DescribeReplicationConfigurationsInput, optFns ...func(*Options)) (*DescribeReplicationConfigurationsOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyMountTargetSecurityGroups(ctx context.Context, params *ModifyMountTargetSecurityGroupsInput, optFns ...func(*Options)) (*ModifyMountTargetSecurityGroupsOutput, error) 
 PutAccountPreferences(ctx context.Context, params *PutAccountPreferencesInput, optFns ...func(*Options)) (*PutAccountPreferencesOutput, error) 
 PutBackupPolicy(ctx context.Context, params *PutBackupPolicyInput, optFns ...func(*Options)) (*PutBackupPolicyOutput, error) 
 PutFileSystemPolicy(ctx context.Context, params *PutFileSystemPolicyInput, optFns ...func(*Options)) (*PutFileSystemPolicyOutput, error) 
 PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) 
 UpdateFileSystemProtection(ctx context.Context, params *UpdateFileSystemProtectionInput, optFns ...func(*Options)) (*UpdateFileSystemProtectionOutput, error) 
}
