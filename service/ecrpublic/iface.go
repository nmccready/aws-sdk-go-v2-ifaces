
package ecrpublic

import (
    "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
)

// IEcrpublic defines the interface for ecrpublic
type IEcrpublic interface {
 Options() Options 
 BatchCheckLayerAvailability(ctx context.Context, params *BatchCheckLayerAvailabilityInput, optFns ...func(*Options)) (*BatchCheckLayerAvailabilityOutput, error) 
 BatchDeleteImage(ctx context.Context, params *BatchDeleteImageInput, optFns ...func(*Options)) (*BatchDeleteImageOutput, error) 
 CompleteLayerUpload(ctx context.Context, params *CompleteLayerUploadInput, optFns ...func(*Options)) (*CompleteLayerUploadOutput, error) 
 CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) 
 DeleteRepository(ctx context.Context, params *DeleteRepositoryInput, optFns ...func(*Options)) (*DeleteRepositoryOutput, error) 
 DeleteRepositoryPolicy(ctx context.Context, params *DeleteRepositoryPolicyInput, optFns ...func(*Options)) (*DeleteRepositoryPolicyOutput, error) 
 DescribeImageTags(ctx context.Context, params *DescribeImageTagsInput, optFns ...func(*Options)) (*DescribeImageTagsOutput, error) 
 DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) 
 DescribeRegistries(ctx context.Context, params *DescribeRegistriesInput, optFns ...func(*Options)) (*DescribeRegistriesOutput, error) 
 DescribeRepositories(ctx context.Context, params *DescribeRepositoriesInput, optFns ...func(*Options)) (*DescribeRepositoriesOutput, error) 
 GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) 
 GetRegistryCatalogData(ctx context.Context, params *GetRegistryCatalogDataInput, optFns ...func(*Options)) (*GetRegistryCatalogDataOutput, error) 
 GetRepositoryCatalogData(ctx context.Context, params *GetRepositoryCatalogDataInput, optFns ...func(*Options)) (*GetRepositoryCatalogDataOutput, error) 
 GetRepositoryPolicy(ctx context.Context, params *GetRepositoryPolicyInput, optFns ...func(*Options)) (*GetRepositoryPolicyOutput, error) 
 InitiateLayerUpload(ctx context.Context, params *InitiateLayerUploadInput, optFns ...func(*Options)) (*InitiateLayerUploadOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutImage(ctx context.Context, params *PutImageInput, optFns ...func(*Options)) (*PutImageOutput, error) 
 PutRegistryCatalogData(ctx context.Context, params *PutRegistryCatalogDataInput, optFns ...func(*Options)) (*PutRegistryCatalogDataOutput, error) 
 PutRepositoryCatalogData(ctx context.Context, params *PutRepositoryCatalogDataInput, optFns ...func(*Options)) (*PutRepositoryCatalogDataOutput, error) 
 SetRepositoryPolicy(ctx context.Context, params *SetRepositoryPolicyInput, optFns ...func(*Options)) (*SetRepositoryPolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UploadLayerPart(ctx context.Context, params *UploadLayerPartInput, optFns ...func(*Options)) (*UploadLayerPartOutput, error) 
}
