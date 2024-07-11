
package s3control

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/s3control"
)

// IClient defines the interface for s3control
type IClient interface {
 Options() Options 
 AssociateAccessGrantsIdentityCenter(ctx context.Context, params *AssociateAccessGrantsIdentityCenterInput, optFns ...func(*Options)) (*AssociateAccessGrantsIdentityCenterOutput, error) 
 CreateAccessGrant(ctx context.Context, params *CreateAccessGrantInput, optFns ...func(*Options)) (*CreateAccessGrantOutput, error) 
 CreateAccessGrantsInstance(ctx context.Context, params *CreateAccessGrantsInstanceInput, optFns ...func(*Options)) (*CreateAccessGrantsInstanceOutput, error) 
 CreateAccessGrantsLocation(ctx context.Context, params *CreateAccessGrantsLocationInput, optFns ...func(*Options)) (*CreateAccessGrantsLocationOutput, error) 
 CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) 
 CreateAccessPointForObjectLambda(ctx context.Context, params *CreateAccessPointForObjectLambdaInput, optFns ...func(*Options)) (*CreateAccessPointForObjectLambdaOutput, error) 
 CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateMultiRegionAccessPoint(ctx context.Context, params *CreateMultiRegionAccessPointInput, optFns ...func(*Options)) (*CreateMultiRegionAccessPointOutput, error) 
 CreateStorageLensGroup(ctx context.Context, params *CreateStorageLensGroupInput, optFns ...func(*Options)) (*CreateStorageLensGroupOutput, error) 
 DeleteAccessGrant(ctx context.Context, params *DeleteAccessGrantInput, optFns ...func(*Options)) (*DeleteAccessGrantOutput, error) 
 DeleteAccessGrantsInstance(ctx context.Context, params *DeleteAccessGrantsInstanceInput, optFns ...func(*Options)) (*DeleteAccessGrantsInstanceOutput, error) 
 DeleteAccessGrantsInstanceResourcePolicy(ctx context.Context, params *DeleteAccessGrantsInstanceResourcePolicyInput, optFns ...func(*Options)) (*DeleteAccessGrantsInstanceResourcePolicyOutput, error) 
 DeleteAccessGrantsLocation(ctx context.Context, params *DeleteAccessGrantsLocationInput, optFns ...func(*Options)) (*DeleteAccessGrantsLocationOutput, error) 
 DeleteAccessPoint(ctx context.Context, params *DeleteAccessPointInput, optFns ...func(*Options)) (*DeleteAccessPointOutput, error) 
 DeleteAccessPointForObjectLambda(ctx context.Context, params *DeleteAccessPointForObjectLambdaInput, optFns ...func(*Options)) (*DeleteAccessPointForObjectLambdaOutput, error) 
 DeleteAccessPointPolicy(ctx context.Context, params *DeleteAccessPointPolicyInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyOutput, error) 
 DeleteAccessPointPolicyForObjectLambda(ctx context.Context, params *DeleteAccessPointPolicyForObjectLambdaInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyForObjectLambdaOutput, error) 
 DeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*Options)) (*DeleteBucketOutput, error) 
 DeleteBucketLifecycleConfiguration(ctx context.Context, params *DeleteBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*DeleteBucketLifecycleConfigurationOutput, error) 
 DeleteBucketPolicy(ctx context.Context, params *DeleteBucketPolicyInput, optFns ...func(*Options)) (*DeleteBucketPolicyOutput, error) 
 DeleteBucketReplication(ctx context.Context, params *DeleteBucketReplicationInput, optFns ...func(*Options)) (*DeleteBucketReplicationOutput, error) 
 DeleteBucketTagging(ctx context.Context, params *DeleteBucketTaggingInput, optFns ...func(*Options)) (*DeleteBucketTaggingOutput, error) 
 DeleteJobTagging(ctx context.Context, params *DeleteJobTaggingInput, optFns ...func(*Options)) (*DeleteJobTaggingOutput, error) 
 DeleteMultiRegionAccessPoint(ctx context.Context, params *DeleteMultiRegionAccessPointInput, optFns ...func(*Options)) (*DeleteMultiRegionAccessPointOutput, error) 
 DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error) 
 DeleteStorageLensConfiguration(ctx context.Context, params *DeleteStorageLensConfigurationInput, optFns ...func(*Options)) (*DeleteStorageLensConfigurationOutput, error) 
 DeleteStorageLensConfigurationTagging(ctx context.Context, params *DeleteStorageLensConfigurationTaggingInput, optFns ...func(*Options)) (*DeleteStorageLensConfigurationTaggingOutput, error) 
 DeleteStorageLensGroup(ctx context.Context, params *DeleteStorageLensGroupInput, optFns ...func(*Options)) (*DeleteStorageLensGroupOutput, error) 
 DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) 
 DescribeMultiRegionAccessPointOperation(ctx context.Context, params *DescribeMultiRegionAccessPointOperationInput, optFns ...func(*Options)) (*DescribeMultiRegionAccessPointOperationOutput, error) 
 DissociateAccessGrantsIdentityCenter(ctx context.Context, params *DissociateAccessGrantsIdentityCenterInput, optFns ...func(*Options)) (*DissociateAccessGrantsIdentityCenterOutput, error) 
 GetAccessGrant(ctx context.Context, params *GetAccessGrantInput, optFns ...func(*Options)) (*GetAccessGrantOutput, error) 
 GetAccessGrantsInstance(ctx context.Context, params *GetAccessGrantsInstanceInput, optFns ...func(*Options)) (*GetAccessGrantsInstanceOutput, error) 
 GetAccessGrantsInstanceForPrefix(ctx context.Context, params *GetAccessGrantsInstanceForPrefixInput, optFns ...func(*Options)) (*GetAccessGrantsInstanceForPrefixOutput, error) 
 GetAccessGrantsInstanceResourcePolicy(ctx context.Context, params *GetAccessGrantsInstanceResourcePolicyInput, optFns ...func(*Options)) (*GetAccessGrantsInstanceResourcePolicyOutput, error) 
 GetAccessGrantsLocation(ctx context.Context, params *GetAccessGrantsLocationInput, optFns ...func(*Options)) (*GetAccessGrantsLocationOutput, error) 
 GetAccessPoint(ctx context.Context, params *GetAccessPointInput, optFns ...func(*Options)) (*GetAccessPointOutput, error) 
 GetAccessPointConfigurationForObjectLambda(ctx context.Context, params *GetAccessPointConfigurationForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointConfigurationForObjectLambdaOutput, error) 
 GetAccessPointForObjectLambda(ctx context.Context, params *GetAccessPointForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointForObjectLambdaOutput, error) 
 GetAccessPointPolicy(ctx context.Context, params *GetAccessPointPolicyInput, optFns ...func(*Options)) (*GetAccessPointPolicyOutput, error) 
 GetAccessPointPolicyForObjectLambda(ctx context.Context, params *GetAccessPointPolicyForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointPolicyForObjectLambdaOutput, error) 
 GetAccessPointPolicyStatus(ctx context.Context, params *GetAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusOutput, error) 
 GetAccessPointPolicyStatusForObjectLambda(ctx context.Context, params *GetAccessPointPolicyStatusForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusForObjectLambdaOutput, error) 
 GetBucket(ctx context.Context, params *GetBucketInput, optFns ...func(*Options)) (*GetBucketOutput, error) 
 GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) 
 GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error) 
 GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) 
 GetBucketTagging(ctx context.Context, params *GetBucketTaggingInput, optFns ...func(*Options)) (*GetBucketTaggingOutput, error) 
 GetBucketVersioning(ctx context.Context, params *GetBucketVersioningInput, optFns ...func(*Options)) (*GetBucketVersioningOutput, error) 
 GetDataAccess(ctx context.Context, params *GetDataAccessInput, optFns ...func(*Options)) (*GetDataAccessOutput, error) 
 GetJobTagging(ctx context.Context, params *GetJobTaggingInput, optFns ...func(*Options)) (*GetJobTaggingOutput, error) 
 GetMultiRegionAccessPoint(ctx context.Context, params *GetMultiRegionAccessPointInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointOutput, error) 
 GetMultiRegionAccessPointPolicy(ctx context.Context, params *GetMultiRegionAccessPointPolicyInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointPolicyOutput, error) 
 GetMultiRegionAccessPointPolicyStatus(ctx context.Context, params *GetMultiRegionAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointPolicyStatusOutput, error) 
 GetMultiRegionAccessPointRoutes(ctx context.Context, params *GetMultiRegionAccessPointRoutesInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointRoutesOutput, error) 
 GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error) 
 GetStorageLensConfiguration(ctx context.Context, params *GetStorageLensConfigurationInput, optFns ...func(*Options)) (*GetStorageLensConfigurationOutput, error) 
 GetStorageLensConfigurationTagging(ctx context.Context, params *GetStorageLensConfigurationTaggingInput, optFns ...func(*Options)) (*GetStorageLensConfigurationTaggingOutput, error) 
 GetStorageLensGroup(ctx context.Context, params *GetStorageLensGroupInput, optFns ...func(*Options)) (*GetStorageLensGroupOutput, error) 
 ListAccessGrants(ctx context.Context, params *ListAccessGrantsInput, optFns ...func(*Options)) (*ListAccessGrantsOutput, error) 
 ListAccessGrantsInstances(ctx context.Context, params *ListAccessGrantsInstancesInput, optFns ...func(*Options)) (*ListAccessGrantsInstancesOutput, error) 
 ListAccessGrantsLocations(ctx context.Context, params *ListAccessGrantsLocationsInput, optFns ...func(*Options)) (*ListAccessGrantsLocationsOutput, error) 
 ListAccessPoints(ctx context.Context, params *ListAccessPointsInput, optFns ...func(*Options)) (*ListAccessPointsOutput, error) 
 ListAccessPointsForObjectLambda(ctx context.Context, params *ListAccessPointsForObjectLambdaInput, optFns ...func(*Options)) (*ListAccessPointsForObjectLambdaOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListMultiRegionAccessPoints(ctx context.Context, params *ListMultiRegionAccessPointsInput, optFns ...func(*Options)) (*ListMultiRegionAccessPointsOutput, error) 
 ListRegionalBuckets(ctx context.Context, params *ListRegionalBucketsInput, optFns ...func(*Options)) (*ListRegionalBucketsOutput, error) 
 ListStorageLensConfigurations(ctx context.Context, params *ListStorageLensConfigurationsInput, optFns ...func(*Options)) (*ListStorageLensConfigurationsOutput, error) 
 ListStorageLensGroups(ctx context.Context, params *ListStorageLensGroupsInput, optFns ...func(*Options)) (*ListStorageLensGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAccessGrantsInstanceResourcePolicy(ctx context.Context, params *PutAccessGrantsInstanceResourcePolicyInput, optFns ...func(*Options)) (*PutAccessGrantsInstanceResourcePolicyOutput, error) 
 PutAccessPointConfigurationForObjectLambda(ctx context.Context, params *PutAccessPointConfigurationForObjectLambdaInput, optFns ...func(*Options)) (*PutAccessPointConfigurationForObjectLambdaOutput, error) 
 PutAccessPointPolicy(ctx context.Context, params *PutAccessPointPolicyInput, optFns ...func(*Options)) (*PutAccessPointPolicyOutput, error) 
 PutAccessPointPolicyForObjectLambda(ctx context.Context, params *PutAccessPointPolicyForObjectLambdaInput, optFns ...func(*Options)) (*PutAccessPointPolicyForObjectLambdaOutput, error) 
 PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) 
 PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error) 
 PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error) 
 PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) 
 PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) 
 PutJobTagging(ctx context.Context, params *PutJobTaggingInput, optFns ...func(*Options)) (*PutJobTaggingOutput, error) 
 PutMultiRegionAccessPointPolicy(ctx context.Context, params *PutMultiRegionAccessPointPolicyInput, optFns ...func(*Options)) (*PutMultiRegionAccessPointPolicyOutput, error) 
 PutPublicAccessBlock(ctx context.Context, params *PutPublicAccessBlockInput, optFns ...func(*Options)) (*PutPublicAccessBlockOutput, error) 
 PutStorageLensConfiguration(ctx context.Context, params *PutStorageLensConfigurationInput, optFns ...func(*Options)) (*PutStorageLensConfigurationOutput, error) 
 PutStorageLensConfigurationTagging(ctx context.Context, params *PutStorageLensConfigurationTaggingInput, optFns ...func(*Options)) (*PutStorageLensConfigurationTaggingOutput, error) 
 SubmitMultiRegionAccessPointRoutes(ctx context.Context, params *SubmitMultiRegionAccessPointRoutesInput, optFns ...func(*Options)) (*SubmitMultiRegionAccessPointRoutesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccessGrantsLocation(ctx context.Context, params *UpdateAccessGrantsLocationInput, optFns ...func(*Options)) (*UpdateAccessGrantsLocationOutput, error) 
 UpdateJobPriority(ctx context.Context, params *UpdateJobPriorityInput, optFns ...func(*Options)) (*UpdateJobPriorityOutput, error) 
 UpdateJobStatus(ctx context.Context, params *UpdateJobStatusInput, optFns ...func(*Options)) (*UpdateJobStatusOutput, error) 
 UpdateStorageLensGroup(ctx context.Context, params *UpdateStorageLensGroupInput, optFns ...func(*Options)) (*UpdateStorageLensGroupOutput, error) 
}
