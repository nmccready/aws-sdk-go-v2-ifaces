package glacier_test

// tests for the glacier service interface for this ../../../service/glacier/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/glacier/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/glacier/glacier_iface"
	"github.com/stretchr/testify/assert"
)

func TestGlacierServiceCanBeMocked(t *testing.T) {
	var iface glacier_iface.IClient
	iface = &glacier.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := glacier.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortMultipartUpload", func(t *testing.T) {
        input := &glacier.AbortMultipartUploadInput{}
        output := &glacier.AbortMultipartUploadOutput{}

        mockClient.On("AbortMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.AbortMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortVaultLock", func(t *testing.T) {
        input := &glacier.AbortVaultLockInput{}
        output := &glacier.AbortVaultLockOutput{}

        mockClient.On("AbortVaultLock", ctx, input).Return(output, nil)

        result, err := mockClient.AbortVaultLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToVault", func(t *testing.T) {
        input := &glacier.AddTagsToVaultInput{}
        output := &glacier.AddTagsToVaultOutput{}

        mockClient.On("AddTagsToVault", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteMultipartUpload", func(t *testing.T) {
        input := &glacier.CompleteMultipartUploadInput{}
        output := &glacier.CompleteMultipartUploadOutput{}

        mockClient.On("CompleteMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteVaultLock", func(t *testing.T) {
        input := &glacier.CompleteVaultLockInput{}
        output := &glacier.CompleteVaultLockOutput{}

        mockClient.On("CompleteVaultLock", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteVaultLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVault", func(t *testing.T) {
        input := &glacier.CreateVaultInput{}
        output := &glacier.CreateVaultOutput{}

        mockClient.On("CreateVault", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArchive", func(t *testing.T) {
        input := &glacier.DeleteArchiveInput{}
        output := &glacier.DeleteArchiveOutput{}

        mockClient.On("DeleteArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVault", func(t *testing.T) {
        input := &glacier.DeleteVaultInput{}
        output := &glacier.DeleteVaultOutput{}

        mockClient.On("DeleteVault", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVaultAccessPolicy", func(t *testing.T) {
        input := &glacier.DeleteVaultAccessPolicyInput{}
        output := &glacier.DeleteVaultAccessPolicyOutput{}

        mockClient.On("DeleteVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVaultNotifications", func(t *testing.T) {
        input := &glacier.DeleteVaultNotificationsInput{}
        output := &glacier.DeleteVaultNotificationsOutput{}

        mockClient.On("DeleteVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJob", func(t *testing.T) {
        input := &glacier.DescribeJobInput{}
        output := &glacier.DescribeJobOutput{}

        mockClient.On("DescribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVault", func(t *testing.T) {
        input := &glacier.DescribeVaultInput{}
        output := &glacier.DescribeVaultOutput{}

        mockClient.On("DescribeVault", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataRetrievalPolicy", func(t *testing.T) {
        input := &glacier.GetDataRetrievalPolicyInput{}
        output := &glacier.GetDataRetrievalPolicyOutput{}

        mockClient.On("GetDataRetrievalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataRetrievalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobOutput", func(t *testing.T) {
        input := &glacier.GetJobOutputInput{}
        output := &glacier.GetJobOutputOutput{}

        mockClient.On("GetJobOutput", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVaultAccessPolicy", func(t *testing.T) {
        input := &glacier.GetVaultAccessPolicyInput{}
        output := &glacier.GetVaultAccessPolicyOutput{}

        mockClient.On("GetVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVaultLock", func(t *testing.T) {
        input := &glacier.GetVaultLockInput{}
        output := &glacier.GetVaultLockOutput{}

        mockClient.On("GetVaultLock", ctx, input).Return(output, nil)

        result, err := mockClient.GetVaultLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVaultNotifications", func(t *testing.T) {
        input := &glacier.GetVaultNotificationsInput{}
        output := &glacier.GetVaultNotificationsOutput{}

        mockClient.On("GetVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.GetVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateJob", func(t *testing.T) {
        input := &glacier.InitiateJobInput{}
        output := &glacier.InitiateJobOutput{}

        mockClient.On("InitiateJob", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateMultipartUpload", func(t *testing.T) {
        input := &glacier.InitiateMultipartUploadInput{}
        output := &glacier.InitiateMultipartUploadOutput{}

        mockClient.On("InitiateMultipartUpload", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateMultipartUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateVaultLock", func(t *testing.T) {
        input := &glacier.InitiateVaultLockInput{}
        output := &glacier.InitiateVaultLockOutput{}

        mockClient.On("InitiateVaultLock", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateVaultLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &glacier.ListJobsInput{}
        output := &glacier.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultipartUploads", func(t *testing.T) {
        input := &glacier.ListMultipartUploadsInput{}
        output := &glacier.ListMultipartUploadsOutput{}

        mockClient.On("ListMultipartUploads", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultipartUploads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParts", func(t *testing.T) {
        input := &glacier.ListPartsInput{}
        output := &glacier.ListPartsOutput{}

        mockClient.On("ListParts", ctx, input).Return(output, nil)

        result, err := mockClient.ListParts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisionedCapacity", func(t *testing.T) {
        input := &glacier.ListProvisionedCapacityInput{}
        output := &glacier.ListProvisionedCapacityOutput{}

        mockClient.On("ListProvisionedCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisionedCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForVault", func(t *testing.T) {
        input := &glacier.ListTagsForVaultInput{}
        output := &glacier.ListTagsForVaultOutput{}

        mockClient.On("ListTagsForVault", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVaults", func(t *testing.T) {
        input := &glacier.ListVaultsInput{}
        output := &glacier.ListVaultsOutput{}

        mockClient.On("ListVaults", ctx, input).Return(output, nil)

        result, err := mockClient.ListVaults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseProvisionedCapacity", func(t *testing.T) {
        input := &glacier.PurchaseProvisionedCapacityInput{}
        output := &glacier.PurchaseProvisionedCapacityOutput{}

        mockClient.On("PurchaseProvisionedCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseProvisionedCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromVault", func(t *testing.T) {
        input := &glacier.RemoveTagsFromVaultInput{}
        output := &glacier.RemoveTagsFromVaultOutput{}

        mockClient.On("RemoveTagsFromVault", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDataRetrievalPolicy", func(t *testing.T) {
        input := &glacier.SetDataRetrievalPolicyInput{}
        output := &glacier.SetDataRetrievalPolicyOutput{}

        mockClient.On("SetDataRetrievalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SetDataRetrievalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetVaultAccessPolicy", func(t *testing.T) {
        input := &glacier.SetVaultAccessPolicyInput{}
        output := &glacier.SetVaultAccessPolicyOutput{}

        mockClient.On("SetVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SetVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetVaultNotifications", func(t *testing.T) {
        input := &glacier.SetVaultNotificationsInput{}
        output := &glacier.SetVaultNotificationsOutput{}

        mockClient.On("SetVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.SetVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadArchive", func(t *testing.T) {
        input := &glacier.UploadArchiveInput{}
        output := &glacier.UploadArchiveOutput{}

        mockClient.On("UploadArchive", ctx, input).Return(output, nil)

        result, err := mockClient.UploadArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadMultipartPart", func(t *testing.T) {
        input := &glacier.UploadMultipartPartInput{}
        output := &glacier.UploadMultipartPartOutput{}

        mockClient.On("UploadMultipartPart", ctx, input).Return(output, nil)

        result, err := mockClient.UploadMultipartPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
