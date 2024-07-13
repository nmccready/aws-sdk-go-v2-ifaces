package ecrpublic_test

// tests for the ecrpublic service interface for this ../../../service/ecrpublic/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecrpublic/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecrpublic/ecrpublic_iface"
	"github.com/stretchr/testify/assert"
)

func TestEcrpublicServiceCanBeMocked(t *testing.T) {
	var iface ecrpublic_iface.IClient
	iface = &ecrpublic.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ecrpublic.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCheckLayerAvailability", func(t *testing.T) {
        input := &ecrpublic.BatchCheckLayerAvailabilityInput{}
        output := &ecrpublic.BatchCheckLayerAvailabilityOutput{}

        mockClient.On("BatchCheckLayerAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCheckLayerAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteImage", func(t *testing.T) {
        input := &ecrpublic.BatchDeleteImageInput{}
        output := &ecrpublic.BatchDeleteImageOutput{}

        mockClient.On("BatchDeleteImage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteLayerUpload", func(t *testing.T) {
        input := &ecrpublic.CompleteLayerUploadInput{}
        output := &ecrpublic.CompleteLayerUploadOutput{}

        mockClient.On("CompleteLayerUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteLayerUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepository", func(t *testing.T) {
        input := &ecrpublic.CreateRepositoryInput{}
        output := &ecrpublic.CreateRepositoryOutput{}

        mockClient.On("CreateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepository", func(t *testing.T) {
        input := &ecrpublic.DeleteRepositoryInput{}
        output := &ecrpublic.DeleteRepositoryOutput{}

        mockClient.On("DeleteRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepositoryPolicy", func(t *testing.T) {
        input := &ecrpublic.DeleteRepositoryPolicyInput{}
        output := &ecrpublic.DeleteRepositoryPolicyOutput{}

        mockClient.On("DeleteRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageTags", func(t *testing.T) {
        input := &ecrpublic.DescribeImageTagsInput{}
        output := &ecrpublic.DescribeImageTagsOutput{}

        mockClient.On("DescribeImageTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImages", func(t *testing.T) {
        input := &ecrpublic.DescribeImagesInput{}
        output := &ecrpublic.DescribeImagesOutput{}

        mockClient.On("DescribeImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistries", func(t *testing.T) {
        input := &ecrpublic.DescribeRegistriesInput{}
        output := &ecrpublic.DescribeRegistriesOutput{}

        mockClient.On("DescribeRegistries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRepositories", func(t *testing.T) {
        input := &ecrpublic.DescribeRepositoriesInput{}
        output := &ecrpublic.DescribeRepositoriesOutput{}

        mockClient.On("DescribeRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizationToken", func(t *testing.T) {
        input := &ecrpublic.GetAuthorizationTokenInput{}
        output := &ecrpublic.GetAuthorizationTokenOutput{}

        mockClient.On("GetAuthorizationToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizationToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegistryCatalogData", func(t *testing.T) {
        input := &ecrpublic.GetRegistryCatalogDataInput{}
        output := &ecrpublic.GetRegistryCatalogDataOutput{}

        mockClient.On("GetRegistryCatalogData", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegistryCatalogData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryCatalogData", func(t *testing.T) {
        input := &ecrpublic.GetRepositoryCatalogDataInput{}
        output := &ecrpublic.GetRepositoryCatalogDataOutput{}

        mockClient.On("GetRepositoryCatalogData", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryCatalogData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryPolicy", func(t *testing.T) {
        input := &ecrpublic.GetRepositoryPolicyInput{}
        output := &ecrpublic.GetRepositoryPolicyOutput{}

        mockClient.On("GetRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateLayerUpload", func(t *testing.T) {
        input := &ecrpublic.InitiateLayerUploadInput{}
        output := &ecrpublic.InitiateLayerUploadOutput{}

        mockClient.On("InitiateLayerUpload", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateLayerUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ecrpublic.ListTagsForResourceInput{}
        output := &ecrpublic.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImage", func(t *testing.T) {
        input := &ecrpublic.PutImageInput{}
        output := &ecrpublic.PutImageOutput{}

        mockClient.On("PutImage", ctx, input).Return(output, nil)

        result, err := mockClient.PutImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRegistryCatalogData", func(t *testing.T) {
        input := &ecrpublic.PutRegistryCatalogDataInput{}
        output := &ecrpublic.PutRegistryCatalogDataOutput{}

        mockClient.On("PutRegistryCatalogData", ctx, input).Return(output, nil)

        result, err := mockClient.PutRegistryCatalogData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRepositoryCatalogData", func(t *testing.T) {
        input := &ecrpublic.PutRepositoryCatalogDataInput{}
        output := &ecrpublic.PutRepositoryCatalogDataOutput{}

        mockClient.On("PutRepositoryCatalogData", ctx, input).Return(output, nil)

        result, err := mockClient.PutRepositoryCatalogData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetRepositoryPolicy", func(t *testing.T) {
        input := &ecrpublic.SetRepositoryPolicyInput{}
        output := &ecrpublic.SetRepositoryPolicyOutput{}

        mockClient.On("SetRepositoryPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SetRepositoryPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ecrpublic.TagResourceInput{}
        output := &ecrpublic.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ecrpublic.UntagResourceInput{}
        output := &ecrpublic.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadLayerPart", func(t *testing.T) {
        input := &ecrpublic.UploadLayerPartInput{}
        output := &ecrpublic.UploadLayerPartOutput{}

        mockClient.On("UploadLayerPart", ctx, input).Return(output, nil)

        result, err := mockClient.UploadLayerPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
