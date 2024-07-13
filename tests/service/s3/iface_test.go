package s3_test

// tests for the s3 service interface for this ../../../service/s3/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3/s3_iface"
	"github.com/stretchr/testify/assert"
)

func TestS3ServiceCanBeMocked(t *testing.T) {
	var iface s3_iface.IClient
	iface = &s3.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := s3.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortMultipartUpload", func(t *testing.T) {
        input := &s3.AbortMultipartUploadInput{}
        output := &s3.AbortMultipartUploadOutput{}

        mockClient.On("AbortMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.AbortMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteMultipartUpload", func(t *testing.T) {
        input := &s3.CompleteMultipartUploadInput{}
        output := &s3.CompleteMultipartUploadOutput{}

        mockClient.On("CompleteMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyObject", func(t *testing.T) {
        input := &s3.CopyObjectInput{}
        output := &s3.CopyObjectOutput{}

        mockClient.On("CopyObject", ctx, input).Return(output, nil)

        result, err := mockClient.CopyObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBucket", func(t *testing.T) {
        input := &s3.CreateBucketInput{}
        output := &s3.CreateBucketOutput{}

        mockClient.On("CreateBucket", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMultipartUpload", func(t *testing.T) {
        input := &s3.CreateMultipartUploadInput{}
        output := &s3.CreateMultipartUploadOutput{}

        mockClient.On("CreateMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSession", func(t *testing.T) {
        input := &s3.CreateSessionInput{}
        output := &s3.CreateSessionOutput{}

        mockClient.On("CreateSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucket", func(t *testing.T) {
        input := &s3.DeleteBucketInput{}
        output := &s3.DeleteBucketOutput{}

        mockClient.On("DeleteBucket", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketAnalyticsConfiguration", func(t *testing.T) {
        input := &s3.DeleteBucketAnalyticsConfigurationInput{}
        output := &s3.DeleteBucketAnalyticsConfigurationOutput{}

        mockClient.On("DeleteBucketAnalyticsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketAnalyticsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketCors", func(t *testing.T) {
        input := &s3.DeleteBucketCorsInput{}
        output := &s3.DeleteBucketCorsOutput{}

        mockClient.On("DeleteBucketCors", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketCors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketEncryption", func(t *testing.T) {
        input := &s3.DeleteBucketEncryptionInput{}
        output := &s3.DeleteBucketEncryptionOutput{}

        mockClient.On("DeleteBucketEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketIntelligentTieringConfiguration", func(t *testing.T) {
        input := &s3.DeleteBucketIntelligentTieringConfigurationInput{}
        output := &s3.DeleteBucketIntelligentTieringConfigurationOutput{}

        mockClient.On("DeleteBucketIntelligentTieringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketIntelligentTieringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketInventoryConfiguration", func(t *testing.T) {
        input := &s3.DeleteBucketInventoryConfigurationInput{}
        output := &s3.DeleteBucketInventoryConfigurationOutput{}

        mockClient.On("DeleteBucketInventoryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketInventoryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketLifecycle", func(t *testing.T) {
        input := &s3.DeleteBucketLifecycleInput{}
        output := &s3.DeleteBucketLifecycleOutput{}

        mockClient.On("DeleteBucketLifecycle", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketLifecycle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketMetricsConfiguration", func(t *testing.T) {
        input := &s3.DeleteBucketMetricsConfigurationInput{}
        output := &s3.DeleteBucketMetricsConfigurationOutput{}

        mockClient.On("DeleteBucketMetricsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketMetricsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketOwnershipControls", func(t *testing.T) {
        input := &s3.DeleteBucketOwnershipControlsInput{}
        output := &s3.DeleteBucketOwnershipControlsOutput{}

        mockClient.On("DeleteBucketOwnershipControls", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketOwnershipControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketPolicy", func(t *testing.T) {
        input := &s3.DeleteBucketPolicyInput{}
        output := &s3.DeleteBucketPolicyOutput{}

        mockClient.On("DeleteBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketReplication", func(t *testing.T) {
        input := &s3.DeleteBucketReplicationInput{}
        output := &s3.DeleteBucketReplicationOutput{}

        mockClient.On("DeleteBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketTagging", func(t *testing.T) {
        input := &s3.DeleteBucketTaggingInput{}
        output := &s3.DeleteBucketTaggingOutput{}

        mockClient.On("DeleteBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketWebsite", func(t *testing.T) {
        input := &s3.DeleteBucketWebsiteInput{}
        output := &s3.DeleteBucketWebsiteOutput{}

        mockClient.On("DeleteBucketWebsite", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketWebsite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObject", func(t *testing.T) {
        input := &s3.DeleteObjectInput{}
        output := &s3.DeleteObjectOutput{}

        mockClient.On("DeleteObject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObjectTagging", func(t *testing.T) {
        input := &s3.DeleteObjectTaggingInput{}
        output := &s3.DeleteObjectTaggingOutput{}

        mockClient.On("DeleteObjectTagging", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObjectTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObjects", func(t *testing.T) {
        input := &s3.DeleteObjectsInput{}
        output := &s3.DeleteObjectsOutput{}

        mockClient.On("DeleteObjects", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublicAccessBlock", func(t *testing.T) {
        input := &s3.DeletePublicAccessBlockInput{}
        output := &s3.DeletePublicAccessBlockOutput{}

        mockClient.On("DeletePublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketAccelerateConfiguration", func(t *testing.T) {
        input := &s3.GetBucketAccelerateConfigurationInput{}
        output := &s3.GetBucketAccelerateConfigurationOutput{}

        mockClient.On("GetBucketAccelerateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketAccelerateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketAcl", func(t *testing.T) {
        input := &s3.GetBucketAclInput{}
        output := &s3.GetBucketAclOutput{}

        mockClient.On("GetBucketAcl", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketAnalyticsConfiguration", func(t *testing.T) {
        input := &s3.GetBucketAnalyticsConfigurationInput{}
        output := &s3.GetBucketAnalyticsConfigurationOutput{}

        mockClient.On("GetBucketAnalyticsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketAnalyticsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketCors", func(t *testing.T) {
        input := &s3.GetBucketCorsInput{}
        output := &s3.GetBucketCorsOutput{}

        mockClient.On("GetBucketCors", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketCors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketEncryption", func(t *testing.T) {
        input := &s3.GetBucketEncryptionInput{}
        output := &s3.GetBucketEncryptionOutput{}

        mockClient.On("GetBucketEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketIntelligentTieringConfiguration", func(t *testing.T) {
        input := &s3.GetBucketIntelligentTieringConfigurationInput{}
        output := &s3.GetBucketIntelligentTieringConfigurationOutput{}

        mockClient.On("GetBucketIntelligentTieringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketIntelligentTieringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketInventoryConfiguration", func(t *testing.T) {
        input := &s3.GetBucketInventoryConfigurationInput{}
        output := &s3.GetBucketInventoryConfigurationOutput{}

        mockClient.On("GetBucketInventoryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketInventoryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketLifecycleConfiguration", func(t *testing.T) {
        input := &s3.GetBucketLifecycleConfigurationInput{}
        output := &s3.GetBucketLifecycleConfigurationOutput{}

        mockClient.On("GetBucketLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketLocation", func(t *testing.T) {
        input := &s3.GetBucketLocationInput{}
        output := &s3.GetBucketLocationOutput{}

        mockClient.On("GetBucketLocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketLogging", func(t *testing.T) {
        input := &s3.GetBucketLoggingInput{}
        output := &s3.GetBucketLoggingOutput{}

        mockClient.On("GetBucketLogging", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketMetricsConfiguration", func(t *testing.T) {
        input := &s3.GetBucketMetricsConfigurationInput{}
        output := &s3.GetBucketMetricsConfigurationOutput{}

        mockClient.On("GetBucketMetricsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketMetricsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketNotificationConfiguration", func(t *testing.T) {
        input := &s3.GetBucketNotificationConfigurationInput{}
        output := &s3.GetBucketNotificationConfigurationOutput{}

        mockClient.On("GetBucketNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketOwnershipControls", func(t *testing.T) {
        input := &s3.GetBucketOwnershipControlsInput{}
        output := &s3.GetBucketOwnershipControlsOutput{}

        mockClient.On("GetBucketOwnershipControls", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketOwnershipControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketPolicy", func(t *testing.T) {
        input := &s3.GetBucketPolicyInput{}
        output := &s3.GetBucketPolicyOutput{}

        mockClient.On("GetBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketPolicyStatus", func(t *testing.T) {
        input := &s3.GetBucketPolicyStatusInput{}
        output := &s3.GetBucketPolicyStatusOutput{}

        mockClient.On("GetBucketPolicyStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketPolicyStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketReplication", func(t *testing.T) {
        input := &s3.GetBucketReplicationInput{}
        output := &s3.GetBucketReplicationOutput{}

        mockClient.On("GetBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketRequestPayment", func(t *testing.T) {
        input := &s3.GetBucketRequestPaymentInput{}
        output := &s3.GetBucketRequestPaymentOutput{}

        mockClient.On("GetBucketRequestPayment", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketRequestPayment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketTagging", func(t *testing.T) {
        input := &s3.GetBucketTaggingInput{}
        output := &s3.GetBucketTaggingOutput{}

        mockClient.On("GetBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketVersioning", func(t *testing.T) {
        input := &s3.GetBucketVersioningInput{}
        output := &s3.GetBucketVersioningOutput{}

        mockClient.On("GetBucketVersioning", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketVersioning(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketWebsite", func(t *testing.T) {
        input := &s3.GetBucketWebsiteInput{}
        output := &s3.GetBucketWebsiteOutput{}

        mockClient.On("GetBucketWebsite", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketWebsite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObject", func(t *testing.T) {
        input := &s3.GetObjectInput{}
        output := &s3.GetObjectOutput{}

        mockClient.On("GetObject", ctx, input).Return(output, nil)

        result, err := mockClient.GetObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectAcl", func(t *testing.T) {
        input := &s3.GetObjectAclInput{}
        output := &s3.GetObjectAclOutput{}

        mockClient.On("GetObjectAcl", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectAttributes", func(t *testing.T) {
        input := &s3.GetObjectAttributesInput{}
        output := &s3.GetObjectAttributesOutput{}

        mockClient.On("GetObjectAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectLegalHold", func(t *testing.T) {
        input := &s3.GetObjectLegalHoldInput{}
        output := &s3.GetObjectLegalHoldOutput{}

        mockClient.On("GetObjectLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectLockConfiguration", func(t *testing.T) {
        input := &s3.GetObjectLockConfigurationInput{}
        output := &s3.GetObjectLockConfigurationOutput{}

        mockClient.On("GetObjectLockConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectLockConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectRetention", func(t *testing.T) {
        input := &s3.GetObjectRetentionInput{}
        output := &s3.GetObjectRetentionOutput{}

        mockClient.On("GetObjectRetention", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectRetention(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectTagging", func(t *testing.T) {
        input := &s3.GetObjectTaggingInput{}
        output := &s3.GetObjectTaggingOutput{}

        mockClient.On("GetObjectTagging", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectTorrent", func(t *testing.T) {
        input := &s3.GetObjectTorrentInput{}
        output := &s3.GetObjectTorrentOutput{}

        mockClient.On("GetObjectTorrent", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectTorrent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicAccessBlock", func(t *testing.T) {
        input := &s3.GetPublicAccessBlockInput{}
        output := &s3.GetPublicAccessBlockOutput{}

        mockClient.On("GetPublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestHeadBucket", func(t *testing.T) {
        input := &s3.HeadBucketInput{}
        output := &s3.HeadBucketOutput{}

        mockClient.On("HeadBucket", ctx, input).Return(output, nil)

        result, err := mockClient.HeadBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestHeadObject", func(t *testing.T) {
        input := &s3.HeadObjectInput{}
        output := &s3.HeadObjectOutput{}

        mockClient.On("HeadObject", ctx, input).Return(output, nil)

        result, err := mockClient.HeadObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBucketAnalyticsConfigurations", func(t *testing.T) {
        input := &s3.ListBucketAnalyticsConfigurationsInput{}
        output := &s3.ListBucketAnalyticsConfigurationsOutput{}

        mockClient.On("ListBucketAnalyticsConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBucketAnalyticsConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBucketIntelligentTieringConfigurations", func(t *testing.T) {
        input := &s3.ListBucketIntelligentTieringConfigurationsInput{}
        output := &s3.ListBucketIntelligentTieringConfigurationsOutput{}

        mockClient.On("ListBucketIntelligentTieringConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBucketIntelligentTieringConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBucketInventoryConfigurations", func(t *testing.T) {
        input := &s3.ListBucketInventoryConfigurationsInput{}
        output := &s3.ListBucketInventoryConfigurationsOutput{}

        mockClient.On("ListBucketInventoryConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBucketInventoryConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBucketMetricsConfigurations", func(t *testing.T) {
        input := &s3.ListBucketMetricsConfigurationsInput{}
        output := &s3.ListBucketMetricsConfigurationsOutput{}

        mockClient.On("ListBucketMetricsConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBucketMetricsConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuckets", func(t *testing.T) {
        input := &s3.ListBucketsInput{}
        output := &s3.ListBucketsOutput{}

        mockClient.On("ListBuckets", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDirectoryBuckets", func(t *testing.T) {
        input := &s3.ListDirectoryBucketsInput{}
        output := &s3.ListDirectoryBucketsOutput{}

        mockClient.On("ListDirectoryBuckets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDirectoryBuckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultipartUploads", func(t *testing.T) {
        input := &s3.ListMultipartUploadsInput{}
        output := &s3.ListMultipartUploadsOutput{}

        mockClient.On("ListMultipartUploads", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultipartUploads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectVersions", func(t *testing.T) {
        input := &s3.ListObjectVersionsInput{}
        output := &s3.ListObjectVersionsOutput{}

        mockClient.On("ListObjectVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjects", func(t *testing.T) {
        input := &s3.ListObjectsInput{}
        output := &s3.ListObjectsOutput{}

        mockClient.On("ListObjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectsV2", func(t *testing.T) {
        input := &s3.ListObjectsV2Input{}
        output := &s3.ListObjectsV2Output{}

        mockClient.On("ListObjectsV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectsV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParts", func(t *testing.T) {
        input := &s3.ListPartsInput{}
        output := &s3.ListPartsOutput{}

        mockClient.On("ListParts", ctx, input).Return(output, nil)

        result, err := mockClient.ListParts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketAccelerateConfiguration", func(t *testing.T) {
        input := &s3.PutBucketAccelerateConfigurationInput{}
        output := &s3.PutBucketAccelerateConfigurationOutput{}

        mockClient.On("PutBucketAccelerateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketAccelerateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketAcl", func(t *testing.T) {
        input := &s3.PutBucketAclInput{}
        output := &s3.PutBucketAclOutput{}

        mockClient.On("PutBucketAcl", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketAnalyticsConfiguration", func(t *testing.T) {
        input := &s3.PutBucketAnalyticsConfigurationInput{}
        output := &s3.PutBucketAnalyticsConfigurationOutput{}

        mockClient.On("PutBucketAnalyticsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketAnalyticsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketCors", func(t *testing.T) {
        input := &s3.PutBucketCorsInput{}
        output := &s3.PutBucketCorsOutput{}

        mockClient.On("PutBucketCors", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketCors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketEncryption", func(t *testing.T) {
        input := &s3.PutBucketEncryptionInput{}
        output := &s3.PutBucketEncryptionOutput{}

        mockClient.On("PutBucketEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketIntelligentTieringConfiguration", func(t *testing.T) {
        input := &s3.PutBucketIntelligentTieringConfigurationInput{}
        output := &s3.PutBucketIntelligentTieringConfigurationOutput{}

        mockClient.On("PutBucketIntelligentTieringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketIntelligentTieringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketInventoryConfiguration", func(t *testing.T) {
        input := &s3.PutBucketInventoryConfigurationInput{}
        output := &s3.PutBucketInventoryConfigurationOutput{}

        mockClient.On("PutBucketInventoryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketInventoryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketLifecycleConfiguration", func(t *testing.T) {
        input := &s3.PutBucketLifecycleConfigurationInput{}
        output := &s3.PutBucketLifecycleConfigurationOutput{}

        mockClient.On("PutBucketLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketLogging", func(t *testing.T) {
        input := &s3.PutBucketLoggingInput{}
        output := &s3.PutBucketLoggingOutput{}

        mockClient.On("PutBucketLogging", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketMetricsConfiguration", func(t *testing.T) {
        input := &s3.PutBucketMetricsConfigurationInput{}
        output := &s3.PutBucketMetricsConfigurationOutput{}

        mockClient.On("PutBucketMetricsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketMetricsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketNotificationConfiguration", func(t *testing.T) {
        input := &s3.PutBucketNotificationConfigurationInput{}
        output := &s3.PutBucketNotificationConfigurationOutput{}

        mockClient.On("PutBucketNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketOwnershipControls", func(t *testing.T) {
        input := &s3.PutBucketOwnershipControlsInput{}
        output := &s3.PutBucketOwnershipControlsOutput{}

        mockClient.On("PutBucketOwnershipControls", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketOwnershipControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketPolicy", func(t *testing.T) {
        input := &s3.PutBucketPolicyInput{}
        output := &s3.PutBucketPolicyOutput{}

        mockClient.On("PutBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketReplication", func(t *testing.T) {
        input := &s3.PutBucketReplicationInput{}
        output := &s3.PutBucketReplicationOutput{}

        mockClient.On("PutBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketRequestPayment", func(t *testing.T) {
        input := &s3.PutBucketRequestPaymentInput{}
        output := &s3.PutBucketRequestPaymentOutput{}

        mockClient.On("PutBucketRequestPayment", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketRequestPayment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketTagging", func(t *testing.T) {
        input := &s3.PutBucketTaggingInput{}
        output := &s3.PutBucketTaggingOutput{}

        mockClient.On("PutBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketVersioning", func(t *testing.T) {
        input := &s3.PutBucketVersioningInput{}
        output := &s3.PutBucketVersioningOutput{}

        mockClient.On("PutBucketVersioning", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketVersioning(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketWebsite", func(t *testing.T) {
        input := &s3.PutBucketWebsiteInput{}
        output := &s3.PutBucketWebsiteOutput{}

        mockClient.On("PutBucketWebsite", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketWebsite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObject", func(t *testing.T) {
        input := &s3.PutObjectInput{}
        output := &s3.PutObjectOutput{}

        mockClient.On("PutObject", ctx, input).Return(output, nil)

        result, err := mockClient.PutObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObjectAcl", func(t *testing.T) {
        input := &s3.PutObjectAclInput{}
        output := &s3.PutObjectAclOutput{}

        mockClient.On("PutObjectAcl", ctx, input).Return(output, nil)

        result, err := mockClient.PutObjectAcl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObjectLegalHold", func(t *testing.T) {
        input := &s3.PutObjectLegalHoldInput{}
        output := &s3.PutObjectLegalHoldOutput{}

        mockClient.On("PutObjectLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.PutObjectLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObjectLockConfiguration", func(t *testing.T) {
        input := &s3.PutObjectLockConfigurationInput{}
        output := &s3.PutObjectLockConfigurationOutput{}

        mockClient.On("PutObjectLockConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutObjectLockConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObjectRetention", func(t *testing.T) {
        input := &s3.PutObjectRetentionInput{}
        output := &s3.PutObjectRetentionOutput{}

        mockClient.On("PutObjectRetention", ctx, input).Return(output, nil)

        result, err := mockClient.PutObjectRetention(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObjectTagging", func(t *testing.T) {
        input := &s3.PutObjectTaggingInput{}
        output := &s3.PutObjectTaggingOutput{}

        mockClient.On("PutObjectTagging", ctx, input).Return(output, nil)

        result, err := mockClient.PutObjectTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPublicAccessBlock", func(t *testing.T) {
        input := &s3.PutPublicAccessBlockInput{}
        output := &s3.PutPublicAccessBlockOutput{}

        mockClient.On("PutPublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.PutPublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreObject", func(t *testing.T) {
        input := &s3.RestoreObjectInput{}
        output := &s3.RestoreObjectOutput{}

        mockClient.On("RestoreObject", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSelectObjectContent", func(t *testing.T) {
        input := &s3.SelectObjectContentInput{}
        output := &s3.SelectObjectContentOutput{}

        mockClient.On("SelectObjectContent", ctx, input).Return(output, nil)

        result, err := mockClient.SelectObjectContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadPart", func(t *testing.T) {
        input := &s3.UploadPartInput{}
        output := &s3.UploadPartOutput{}

        mockClient.On("UploadPart", ctx, input).Return(output, nil)

        result, err := mockClient.UploadPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadPartCopy", func(t *testing.T) {
        input := &s3.UploadPartCopyInput{}
        output := &s3.UploadPartCopyOutput{}

        mockClient.On("UploadPartCopy", ctx, input).Return(output, nil)

        result, err := mockClient.UploadPartCopy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestWriteGetObjectResponse", func(t *testing.T) {
        input := &s3.WriteGetObjectResponseInput{}
        output := &s3.WriteGetObjectResponseOutput{}

        mockClient.On("WriteGetObjectResponse", ctx, input).Return(output, nil)

        result, err := mockClient.WriteGetObjectResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
