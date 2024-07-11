
package ecr

import (
    "github.com/aws/aws-sdk-go-v2/service/ecr"
)

// IEcr defines the interface for ecr
type IEcr interface {
 Options() Options 
 BatchCheckLayerAvailability(ctx context.Context, params *BatchCheckLayerAvailabilityInput, optFns ...func(*Options)) (*BatchCheckLayerAvailabilityOutput, error) 
 BatchDeleteImage(ctx context.Context, params *BatchDeleteImageInput, optFns ...func(*Options)) (*BatchDeleteImageOutput, error) 
 BatchGetImage(ctx context.Context, params *BatchGetImageInput, optFns ...func(*Options)) (*BatchGetImageOutput, error) 
 BatchGetRepositoryScanningConfiguration(ctx context.Context, params *BatchGetRepositoryScanningConfigurationInput, optFns ...func(*Options)) (*BatchGetRepositoryScanningConfigurationOutput, error) 
 CompleteLayerUpload(ctx context.Context, params *CompleteLayerUploadInput, optFns ...func(*Options)) (*CompleteLayerUploadOutput, error) 
 CreatePullThroughCacheRule(ctx context.Context, params *CreatePullThroughCacheRuleInput, optFns ...func(*Options)) (*CreatePullThroughCacheRuleOutput, error) 
 CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) 
 DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) 
 DeletePullThroughCacheRule(ctx context.Context, params *DeletePullThroughCacheRuleInput, optFns ...func(*Options)) (*DeletePullThroughCacheRuleOutput, error) 
 DeleteRegistryPolicy(ctx context.Context, params *DeleteRegistryPolicyInput, optFns ...func(*Options)) (*DeleteRegistryPolicyOutput, error) 
 DeleteRepository(ctx context.Context, params *DeleteRepositoryInput, optFns ...func(*Options)) (*DeleteRepositoryOutput, error) 
 DeleteRepositoryPolicy(ctx context.Context, params *DeleteRepositoryPolicyInput, optFns ...func(*Options)) (*DeleteRepositoryPolicyOutput, error) 
 DescribeImageReplicationStatus(ctx context.Context, params *DescribeImageReplicationStatusInput, optFns ...func(*Options)) (*DescribeImageReplicationStatusOutput, error) 
 DescribeImageScanFindings(ctx context.Context, params *DescribeImageScanFindingsInput, optFns ...func(*Options)) (*DescribeImageScanFindingsOutput, error) 
 DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) 
 DescribePullThroughCacheRules(ctx context.Context, params *DescribePullThroughCacheRulesInput, optFns ...func(*Options)) (*DescribePullThroughCacheRulesOutput, error) 
 DescribeRegistry(ctx context.Context, params *DescribeRegistryInput, optFns ...func(*Options)) (*DescribeRegistryOutput, error) 
 DescribeRepositories(ctx context.Context, params *DescribeRepositoriesInput, optFns ...func(*Options)) (*DescribeRepositoriesOutput, error) 
 GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) 
 GetDownloadUrlForLayer(ctx context.Context, params *GetDownloadUrlForLayerInput, optFns ...func(*Options)) (*GetDownloadUrlForLayerOutput, error) 
 GetLifecyclePolicy(ctx context.Context, params *GetLifecyclePolicyInput, optFns ...func(*Options)) (*GetLifecyclePolicyOutput, error) 
 GetLifecyclePolicyPreview(ctx context.Context, params *GetLifecyclePolicyPreviewInput, optFns ...func(*Options)) (*GetLifecyclePolicyPreviewOutput, error) 
 GetRegistryPolicy(ctx context.Context, params *GetRegistryPolicyInput, optFns ...func(*Options)) (*GetRegistryPolicyOutput, error) 
 GetRegistryScanningConfiguration(ctx context.Context, params *GetRegistryScanningConfigurationInput, optFns ...func(*Options)) (*GetRegistryScanningConfigurationOutput, error) 
 GetRepositoryPolicy(ctx context.Context, params *GetRepositoryPolicyInput, optFns ...func(*Options)) (*GetRepositoryPolicyOutput, error) 
 InitiateLayerUpload(ctx context.Context, params *InitiateLayerUploadInput, optFns ...func(*Options)) (*InitiateLayerUploadOutput, error) 
 ListImages(ctx context.Context, params *ListImagesInput, optFns ...func(*Options)) (*ListImagesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutImage(ctx context.Context, params *PutImageInput, optFns ...func(*Options)) (*PutImageOutput, error) 
 PutImageScanningConfiguration(ctx context.Context, params *PutImageScanningConfigurationInput, optFns ...func(*Options)) (*PutImageScanningConfigurationOutput, error) 
 PutImageTagMutability(ctx context.Context, params *PutImageTagMutabilityInput, optFns ...func(*Options)) (*PutImageTagMutabilityOutput, error) 
 PutLifecyclePolicy(ctx context.Context, params *PutLifecyclePolicyInput, optFns ...func(*Options)) (*PutLifecyclePolicyOutput, error) 
 PutRegistryPolicy(ctx context.Context, params *PutRegistryPolicyInput, optFns ...func(*Options)) (*PutRegistryPolicyOutput, error) 
 PutRegistryScanningConfiguration(ctx context.Context, params *PutRegistryScanningConfigurationInput, optFns ...func(*Options)) (*PutRegistryScanningConfigurationOutput, error) 
 PutReplicationConfiguration(ctx context.Context, params *PutReplicationConfigurationInput, optFns ...func(*Options)) (*PutReplicationConfigurationOutput, error) 
 SetRepositoryPolicy(ctx context.Context, params *SetRepositoryPolicyInput, optFns ...func(*Options)) (*SetRepositoryPolicyOutput, error) 
 StartImageScan(ctx context.Context, params *StartImageScanInput, optFns ...func(*Options)) (*StartImageScanOutput, error) 
 StartLifecyclePolicyPreview(ctx context.Context, params *StartLifecyclePolicyPreviewInput, optFns ...func(*Options)) (*StartLifecyclePolicyPreviewOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePullThroughCacheRule(ctx context.Context, params *UpdatePullThroughCacheRuleInput, optFns ...func(*Options)) (*UpdatePullThroughCacheRuleOutput, error) 
 UploadLayerPart(ctx context.Context, params *UploadLayerPartInput, optFns ...func(*Options)) (*UploadLayerPartOutput, error) 
 ValidatePullThroughCacheRule(ctx context.Context, params *ValidatePullThroughCacheRuleInput, optFns ...func(*Options)) (*ValidatePullThroughCacheRuleOutput, error) 
}
