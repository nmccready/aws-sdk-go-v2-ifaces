package ecr_test

// tests for the ecr service interface for this ../../../service/ecr/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecr/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecr/ecr_iface"
	"github.com/stretchr/testify/assert"
)

func TestEcrServiceCanBeMocked(t *testing.T) {
	var iface ecr_iface.IClient
	iface = &ecr.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ecr.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCheckLayerAvailability", func(t *testing.T) {
        input := &ecr.BatchCheckLayerAvailabilityInput{}
        output := &ecr.BatchCheckLayerAvailabilityOutput{}

        mockClient.On("BatchCheckLayerAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCheckLayerAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteImage", func(t *testing.T) {
        input := &ecr.BatchDeleteImageInput{}
        output := &ecr.BatchDeleteImageOutput{}

        mockClient.On("BatchDeleteImage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetImage", func(t *testing.T) {
        input := &ecr.BatchGetImageInput{}
        output := &ecr.BatchGetImageOutput{}

        mockClient.On("BatchGetImage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetRepositoryScanningConfiguration", func(t *testing.T) {
        input := &ecr.BatchGetRepositoryScanningConfigurationInput{}
        output := &ecr.BatchGetRepositoryScanningConfigurationOutput{}

        mockClient.On("BatchGetRepositoryScanningConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetRepositoryScanningConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteLayerUpload", func(t *testing.T) {
        input := &ecr.CompleteLayerUploadInput{}
        output := &ecr.CompleteLayerUploadOutput{}

        mockClient.On("CompleteLayerUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteLayerUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePullThroughCacheRule", func(t *testing.T) {
        input := &ecr.CreatePullThroughCacheRuleInput{}
        output := &ecr.CreatePullThroughCacheRuleOutput{}

        mockClient.On("CreatePullThroughCacheRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePullThroughCacheRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepository", func(t *testing.T) {
        input := &ecr.CreateRepositoryInput{}
        output := &ecr.CreateRepositoryOutput{}

        mockClient.On("CreateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecyclePolicy", func(t *testing.T) {
        input := &ecr.DeleteLifecyclePolicyInput{}
        output := &ecr.DeleteLifecyclePolicyOutput{}

        mockClient.On("DeleteLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePullThroughCacheRule", func(t *testing.T) {
        input := &ecr.DeletePullThroughCacheRuleInput{}
        output := &ecr.DeletePullThroughCacheRuleOutput{}

        mockClient.On("DeletePullThroughCacheRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePullThroughCacheRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistryPolicy", func(t *testing.T) {
        input := &ecr.DeleteRegistryPolicyInput{}
        output := &ecr.DeleteRegistryPolicyOutput{}

        mockClient.On("DeleteRegistryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepository", func(t *testing.T) {
        input := &ecr.DeleteRepositoryInput{}
        output := &ecr.DeleteRepositoryOutput{}

        mockClient.On("DeleteRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepositoryPolicy", func(t *testing.T) {
        input := &ecr.DeleteRepositoryPolicyInput{}
        output := &ecr.DeleteRepositoryPolicyOutput{}

        mockClient.On("DeleteRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageReplicationStatus", func(t *testing.T) {
        input := &ecr.DescribeImageReplicationStatusInput{}
        output := &ecr.DescribeImageReplicationStatusOutput{}

        mockClient.On("DescribeImageReplicationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageReplicationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageScanFindings", func(t *testing.T) {
        input := &ecr.DescribeImageScanFindingsInput{}
        output := &ecr.DescribeImageScanFindingsOutput{}

        mockClient.On("DescribeImageScanFindings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageScanFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImages", func(t *testing.T) {
        input := &ecr.DescribeImagesInput{}
        output := &ecr.DescribeImagesOutput{}

        mockClient.On("DescribeImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePullThroughCacheRules", func(t *testing.T) {
        input := &ecr.DescribePullThroughCacheRulesInput{}
        output := &ecr.DescribePullThroughCacheRulesOutput{}

        mockClient.On("DescribePullThroughCacheRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePullThroughCacheRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistry", func(t *testing.T) {
        input := &ecr.DescribeRegistryInput{}
        output := &ecr.DescribeRegistryOutput{}

        mockClient.On("DescribeRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRepositories", func(t *testing.T) {
        input := &ecr.DescribeRepositoriesInput{}
        output := &ecr.DescribeRepositoriesOutput{}

        mockClient.On("DescribeRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizationToken", func(t *testing.T) {
        input := &ecr.GetAuthorizationTokenInput{}
        output := &ecr.GetAuthorizationTokenOutput{}

        mockClient.On("GetAuthorizationToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizationToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDownloadUrlForLayer", func(t *testing.T) {
        input := &ecr.GetDownloadUrlForLayerInput{}
        output := &ecr.GetDownloadUrlForLayerOutput{}

        mockClient.On("GetDownloadUrlForLayer", ctx, input).Return(output, nil)

        result, err := mockClient.GetDownloadUrlForLayer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicy", func(t *testing.T) {
        input := &ecr.GetLifecyclePolicyInput{}
        output := &ecr.GetLifecyclePolicyOutput{}

        mockClient.On("GetLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicyPreview", func(t *testing.T) {
        input := &ecr.GetLifecyclePolicyPreviewInput{}
        output := &ecr.GetLifecyclePolicyPreviewOutput{}

        mockClient.On("GetLifecyclePolicyPreview", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicyPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegistryPolicy", func(t *testing.T) {
        input := &ecr.GetRegistryPolicyInput{}
        output := &ecr.GetRegistryPolicyOutput{}

        mockClient.On("GetRegistryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegistryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegistryScanningConfiguration", func(t *testing.T) {
        input := &ecr.GetRegistryScanningConfigurationInput{}
        output := &ecr.GetRegistryScanningConfigurationOutput{}

        mockClient.On("GetRegistryScanningConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegistryScanningConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryPolicy", func(t *testing.T) {
        input := &ecr.GetRepositoryPolicyInput{}
        output := &ecr.GetRepositoryPolicyOutput{}

        mockClient.On("GetRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateLayerUpload", func(t *testing.T) {
        input := &ecr.InitiateLayerUploadInput{}
        output := &ecr.InitiateLayerUploadOutput{}

        mockClient.On("InitiateLayerUpload", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateLayerUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImages", func(t *testing.T) {
        input := &ecr.ListImagesInput{}
        output := &ecr.ListImagesOutput{}

        mockClient.On("ListImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ecr.ListTagsForResourceInput{}
        output := &ecr.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImage", func(t *testing.T) {
        input := &ecr.PutImageInput{}
        output := &ecr.PutImageOutput{}

        mockClient.On("PutImage", ctx, input).Return(output, nil)

        result, err := mockClient.PutImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImageScanningConfiguration", func(t *testing.T) {
        input := &ecr.PutImageScanningConfigurationInput{}
        output := &ecr.PutImageScanningConfigurationOutput{}

        mockClient.On("PutImageScanningConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutImageScanningConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImageTagMutability", func(t *testing.T) {
        input := &ecr.PutImageTagMutabilityInput{}
        output := &ecr.PutImageTagMutabilityOutput{}

        mockClient.On("PutImageTagMutability", ctx, input).Return(output, nil)

        result, err := mockClient.PutImageTagMutability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLifecyclePolicy", func(t *testing.T) {
        input := &ecr.PutLifecyclePolicyInput{}
        output := &ecr.PutLifecyclePolicyOutput{}

        mockClient.On("PutLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRegistryPolicy", func(t *testing.T) {
        input := &ecr.PutRegistryPolicyInput{}
        output := &ecr.PutRegistryPolicyOutput{}

        mockClient.On("PutRegistryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutRegistryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRegistryScanningConfiguration", func(t *testing.T) {
        input := &ecr.PutRegistryScanningConfigurationInput{}
        output := &ecr.PutRegistryScanningConfigurationOutput{}

        mockClient.On("PutRegistryScanningConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutRegistryScanningConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutReplicationConfiguration", func(t *testing.T) {
        input := &ecr.PutReplicationConfigurationInput{}
        output := &ecr.PutReplicationConfigurationOutput{}

        mockClient.On("PutReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetRepositoryPolicy", func(t *testing.T) {
        input := &ecr.SetRepositoryPolicyInput{}
        output := &ecr.SetRepositoryPolicyOutput{}

        mockClient.On("SetRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SetRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImageScan", func(t *testing.T) {
        input := &ecr.StartImageScanInput{}
        output := &ecr.StartImageScanOutput{}

        mockClient.On("StartImageScan", ctx, input).Return(output, nil)

        result, err := mockClient.StartImageScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLifecyclePolicyPreview", func(t *testing.T) {
        input := &ecr.StartLifecyclePolicyPreviewInput{}
        output := &ecr.StartLifecyclePolicyPreviewOutput{}

        mockClient.On("StartLifecyclePolicyPreview", ctx, input).Return(output, nil)

        result, err := mockClient.StartLifecyclePolicyPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ecr.TagResourceInput{}
        output := &ecr.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ecr.UntagResourceInput{}
        output := &ecr.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullThroughCacheRule", func(t *testing.T) {
        input := &ecr.UpdatePullThroughCacheRuleInput{}
        output := &ecr.UpdatePullThroughCacheRuleOutput{}

        mockClient.On("UpdatePullThroughCacheRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullThroughCacheRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadLayerPart", func(t *testing.T) {
        input := &ecr.UploadLayerPartInput{}
        output := &ecr.UploadLayerPartOutput{}

        mockClient.On("UploadLayerPart", ctx, input).Return(output, nil)

        result, err := mockClient.UploadLayerPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidatePullThroughCacheRule", func(t *testing.T) {
        input := &ecr.ValidatePullThroughCacheRuleInput{}
        output := &ecr.ValidatePullThroughCacheRuleOutput{}

        mockClient.On("ValidatePullThroughCacheRule", ctx, input).Return(output, nil)

        result, err := mockClient.ValidatePullThroughCacheRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
