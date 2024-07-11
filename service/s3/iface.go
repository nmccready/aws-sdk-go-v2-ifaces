
package s3

import (
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// IClient defines the interface for s3
type IClient interface {
 Options() Options 
 AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error) 
 CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) 
 CopyObject(ctx context.Context, params *CopyObjectInput, optFns ...func(*Options)) (*CopyObjectOutput, error) 
 CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) 
 CreateMultipartUpload(ctx context.Context, params *CreateMultipartUploadInput, optFns ...func(*Options)) (*CreateMultipartUploadOutput, error) 
 CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) 
 DeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*Options)) (*DeleteBucketOutput, error) 
 PresignDeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 DeleteBucketAnalyticsConfiguration(ctx context.Context, params *DeleteBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketAnalyticsConfigurationOutput, error) 
 DeleteBucketCors(ctx context.Context, params *DeleteBucketCorsInput, optFns ...func(*Options)) (*DeleteBucketCorsOutput, error) 
 DeleteBucketEncryption(ctx context.Context, params *DeleteBucketEncryptionInput, optFns ...func(*Options)) (*DeleteBucketEncryptionOutput, error) 
 DeleteBucketIntelligentTieringConfiguration(ctx context.Context, params *DeleteBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*DeleteBucketIntelligentTieringConfigurationOutput, error) 
 DeleteBucketInventoryConfiguration(ctx context.Context, params *DeleteBucketInventoryConfigurationInput, optFns ...func(*Options)) (*DeleteBucketInventoryConfigurationOutput, error) 
 DeleteBucketLifecycle(ctx context.Context, params *DeleteBucketLifecycleInput, optFns ...func(*Options)) (*DeleteBucketLifecycleOutput, error) 
 DeleteBucketMetricsConfiguration(ctx context.Context, params *DeleteBucketMetricsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketMetricsConfigurationOutput, error) 
 DeleteBucketOwnershipControls(ctx context.Context, params *DeleteBucketOwnershipControlsInput, optFns ...func(*Options)) (*DeleteBucketOwnershipControlsOutput, error) 
 DeleteBucketPolicy(ctx context.Context, params *DeleteBucketPolicyInput, optFns ...func(*Options)) (*DeleteBucketPolicyOutput, error) 
 DeleteBucketReplication(ctx context.Context, params *DeleteBucketReplicationInput, optFns ...func(*Options)) (*DeleteBucketReplicationOutput, error) 
 DeleteBucketTagging(ctx context.Context, params *DeleteBucketTaggingInput, optFns ...func(*Options)) (*DeleteBucketTaggingOutput, error) 
 DeleteBucketWebsite(ctx context.Context, params *DeleteBucketWebsiteInput, optFns ...func(*Options)) (*DeleteBucketWebsiteOutput, error) 
 DeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*Options)) (*DeleteObjectOutput, error) 
 PresignDeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 DeleteObjectTagging(ctx context.Context, params *DeleteObjectTaggingInput, optFns ...func(*Options)) (*DeleteObjectTaggingOutput, error) 
 DeleteObjects(ctx context.Context, params *DeleteObjectsInput, optFns ...func(*Options)) (*DeleteObjectsOutput, error) 
 DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error) 
 GetBucketAccelerateConfiguration(ctx context.Context, params *GetBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*GetBucketAccelerateConfigurationOutput, error) 
 GetBucketAcl(ctx context.Context, params *GetBucketAclInput, optFns ...func(*Options)) (*GetBucketAclOutput, error) 
 GetBucketAnalyticsConfiguration(ctx context.Context, params *GetBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*GetBucketAnalyticsConfigurationOutput, error) 
 GetBucketCors(ctx context.Context, params *GetBucketCorsInput, optFns ...func(*Options)) (*GetBucketCorsOutput, error) 
 GetBucketEncryption(ctx context.Context, params *GetBucketEncryptionInput, optFns ...func(*Options)) (*GetBucketEncryptionOutput, error) 
 GetBucketIntelligentTieringConfiguration(ctx context.Context, params *GetBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*GetBucketIntelligentTieringConfigurationOutput, error) 
 GetBucketInventoryConfiguration(ctx context.Context, params *GetBucketInventoryConfigurationInput, optFns ...func(*Options)) (*GetBucketInventoryConfigurationOutput, error) 
 GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) 
 GetBucketLocation(ctx context.Context, params *GetBucketLocationInput, optFns ...func(*Options)) (*GetBucketLocationOutput, error) 
 GetBucketLogging(ctx context.Context, params *GetBucketLoggingInput, optFns ...func(*Options)) (*GetBucketLoggingOutput, error) 
 GetBucketMetricsConfiguration(ctx context.Context, params *GetBucketMetricsConfigurationInput, optFns ...func(*Options)) (*GetBucketMetricsConfigurationOutput, error) 
 GetBucketNotificationConfiguration(ctx context.Context, params *GetBucketNotificationConfigurationInput, optFns ...func(*Options)) (*GetBucketNotificationConfigurationOutput, error) 
 GetBucketOwnershipControls(ctx context.Context, params *GetBucketOwnershipControlsInput, optFns ...func(*Options)) (*GetBucketOwnershipControlsOutput, error) 
 GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error) 
 GetBucketPolicyStatus(ctx context.Context, params *GetBucketPolicyStatusInput, optFns ...func(*Options)) (*GetBucketPolicyStatusOutput, error) 
 GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) 
 GetBucketRequestPayment(ctx context.Context, params *GetBucketRequestPaymentInput, optFns ...func(*Options)) (*GetBucketRequestPaymentOutput, error) 
 GetBucketTagging(ctx context.Context, params *GetBucketTaggingInput, optFns ...func(*Options)) (*GetBucketTaggingOutput, error) 
 GetBucketVersioning(ctx context.Context, params *GetBucketVersioningInput, optFns ...func(*Options)) (*GetBucketVersioningOutput, error) 
 GetBucketWebsite(ctx context.Context, params *GetBucketWebsiteInput, optFns ...func(*Options)) (*GetBucketWebsiteOutput, error) 
 GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error) 
 PresignGetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 GetObjectAcl(ctx context.Context, params *GetObjectAclInput, optFns ...func(*Options)) (*GetObjectAclOutput, error) 
 GetObjectAttributes(ctx context.Context, params *GetObjectAttributesInput, optFns ...func(*Options)) (*GetObjectAttributesOutput, error) 
 GetObjectLegalHold(ctx context.Context, params *GetObjectLegalHoldInput, optFns ...func(*Options)) (*GetObjectLegalHoldOutput, error) 
 GetObjectLockConfiguration(ctx context.Context, params *GetObjectLockConfigurationInput, optFns ...func(*Options)) (*GetObjectLockConfigurationOutput, error) 
 GetObjectRetention(ctx context.Context, params *GetObjectRetentionInput, optFns ...func(*Options)) (*GetObjectRetentionOutput, error) 
 GetObjectTagging(ctx context.Context, params *GetObjectTaggingInput, optFns ...func(*Options)) (*GetObjectTaggingOutput, error) 
 GetObjectTorrent(ctx context.Context, params *GetObjectTorrentInput, optFns ...func(*Options)) (*GetObjectTorrentOutput, error) 
 GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error) 
 HeadBucket(ctx context.Context, params *HeadBucketInput, optFns ...func(*Options)) (*HeadBucketOutput, error) 
 PresignHeadBucket(ctx context.Context, params *HeadBucketInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 HeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*Options)) (*HeadObjectOutput, error) 
 PresignHeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 ListBucketAnalyticsConfigurations(ctx context.Context, params *ListBucketAnalyticsConfigurationsInput, optFns ...func(*Options)) (*ListBucketAnalyticsConfigurationsOutput, error) 
 ListBucketIntelligentTieringConfigurations(ctx context.Context, params *ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*Options)) (*ListBucketIntelligentTieringConfigurationsOutput, error) 
 ListBucketInventoryConfigurations(ctx context.Context, params *ListBucketInventoryConfigurationsInput, optFns ...func(*Options)) (*ListBucketInventoryConfigurationsOutput, error) 
 ListBucketMetricsConfigurations(ctx context.Context, params *ListBucketMetricsConfigurationsInput, optFns ...func(*Options)) (*ListBucketMetricsConfigurationsOutput, error) 
 ListBuckets(ctx context.Context, params *ListBucketsInput, optFns ...func(*Options)) (*ListBucketsOutput, error) 
 ListDirectoryBuckets(ctx context.Context, params *ListDirectoryBucketsInput, optFns ...func(*Options)) (*ListDirectoryBucketsOutput, error) 
 ListMultipartUploads(ctx context.Context, params *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error) 
 ListObjectVersions(ctx context.Context, params *ListObjectVersionsInput, optFns ...func(*Options)) (*ListObjectVersionsOutput, error) 
 ListObjects(ctx context.Context, params *ListObjectsInput, optFns ...func(*Options)) (*ListObjectsOutput, error) 
 ListObjectsV2(ctx context.Context, params *ListObjectsV2Input, optFns ...func(*Options)) (*ListObjectsV2Output, error) 
 ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) 
 PutBucketAccelerateConfiguration(ctx context.Context, params *PutBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*PutBucketAccelerateConfigurationOutput, error) 
 PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error) 
 PutBucketAnalyticsConfiguration(ctx context.Context, params *PutBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*PutBucketAnalyticsConfigurationOutput, error) 
 PutBucketCors(ctx context.Context, params *PutBucketCorsInput, optFns ...func(*Options)) (*PutBucketCorsOutput, error) 
 PutBucketEncryption(ctx context.Context, params *PutBucketEncryptionInput, optFns ...func(*Options)) (*PutBucketEncryptionOutput, error) 
 PutBucketIntelligentTieringConfiguration(ctx context.Context, params *PutBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*PutBucketIntelligentTieringConfigurationOutput, error) 
 PutBucketInventoryConfiguration(ctx context.Context, params *PutBucketInventoryConfigurationInput, optFns ...func(*Options)) (*PutBucketInventoryConfigurationOutput, error) 
 PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) 
 PutBucketLogging(ctx context.Context, params *PutBucketLoggingInput, optFns ...func(*Options)) (*PutBucketLoggingOutput, error) 
 PutBucketMetricsConfiguration(ctx context.Context, params *PutBucketMetricsConfigurationInput, optFns ...func(*Options)) (*PutBucketMetricsConfigurationOutput, error) 
 PutBucketNotificationConfiguration(ctx context.Context, params *PutBucketNotificationConfigurationInput, optFns ...func(*Options)) (*PutBucketNotificationConfigurationOutput, error) 
 PutBucketOwnershipControls(ctx context.Context, params *PutBucketOwnershipControlsInput, optFns ...func(*Options)) (*PutBucketOwnershipControlsOutput, error) 
 PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error) 
 PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error) 
 PutBucketRequestPayment(ctx context.Context, params *PutBucketRequestPaymentInput, optFns ...func(*Options)) (*PutBucketRequestPaymentOutput, error) 
 PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) 
 PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) 
 PutBucketWebsite(ctx context.Context, params *PutBucketWebsiteInput, optFns ...func(*Options)) (*PutBucketWebsiteOutput, error) 
 PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error) 
 PresignPutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 PutObjectAcl(ctx context.Context, params *PutObjectAclInput, optFns ...func(*Options)) (*PutObjectAclOutput, error) 
 PutObjectLegalHold(ctx context.Context, params *PutObjectLegalHoldInput, optFns ...func(*Options)) (*PutObjectLegalHoldOutput, error) 
 PutObjectLockConfiguration(ctx context.Context, params *PutObjectLockConfigurationInput, optFns ...func(*Options)) (*PutObjectLockConfigurationOutput, error) 
 PutObjectRetention(ctx context.Context, params *PutObjectRetentionInput, optFns ...func(*Options)) (*PutObjectRetentionOutput, error) 
 PutObjectTagging(ctx context.Context, params *PutObjectTaggingInput, optFns ...func(*Options)) (*PutObjectTaggingOutput, error) 
 PutPublicAccessBlock(ctx context.Context, params *PutPublicAccessBlockInput, optFns ...func(*Options)) (*PutPublicAccessBlockOutput, error) 
 RestoreObject(ctx context.Context, params *RestoreObjectInput, optFns ...func(*Options)) (*RestoreObjectOutput, error) 
 SelectObjectContent(ctx context.Context, params *SelectObjectContentInput, optFns ...func(*Options)) (*SelectObjectContentOutput, error) 
 UploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*Options)) (*UploadPartOutput, error) 
 PresignUploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 UploadPartCopy(ctx context.Context, params *UploadPartCopyInput, optFns ...func(*Options)) (*UploadPartCopyOutput, error) 
 WriteGetObjectResponse(ctx context.Context, params *WriteGetObjectResponseInput, optFns ...func(*Options)) (*WriteGetObjectResponseOutput, error) 
 ListObjectVersions(ctx context.Context, input *ListObjectVersionsInput, optFns ...func(*Options)) (*ListObjectVersionsOutput, error) 
 ListMultipartUploads(ctx context.Context, input *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error) 
}
