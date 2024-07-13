package storagegateway_test

// tests for the storagegateway service interface for this ../../../service/storagegateway/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/storagegateway/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/storagegateway/storagegateway_iface"
	"github.com/stretchr/testify/assert"
)

func TestStoragegatewayServiceCanBeMocked(t *testing.T) {
	var iface storagegateway_iface.IClient
	iface = &storagegateway.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := storagegateway.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateGateway", func(t *testing.T) {
        input := &storagegateway.ActivateGatewayInput{}
        output := &storagegateway.ActivateGatewayOutput{}

        mockClient.On("ActivateGateway", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddCache", func(t *testing.T) {
        input := &storagegateway.AddCacheInput{}
        output := &storagegateway.AddCacheOutput{}

        mockClient.On("AddCache", ctx, input).Return(output, nil)

        result, err := mockClient.AddCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &storagegateway.AddTagsToResourceInput{}
        output := &storagegateway.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddUploadBuffer", func(t *testing.T) {
        input := &storagegateway.AddUploadBufferInput{}
        output := &storagegateway.AddUploadBufferOutput{}

        mockClient.On("AddUploadBuffer", ctx, input).Return(output, nil)

        result, err := mockClient.AddUploadBuffer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddWorkingStorage", func(t *testing.T) {
        input := &storagegateway.AddWorkingStorageInput{}
        output := &storagegateway.AddWorkingStorageOutput{}

        mockClient.On("AddWorkingStorage", ctx, input).Return(output, nil)

        result, err := mockClient.AddWorkingStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignTapePool", func(t *testing.T) {
        input := &storagegateway.AssignTapePoolInput{}
        output := &storagegateway.AssignTapePoolOutput{}

        mockClient.On("AssignTapePool", ctx, input).Return(output, nil)

        result, err := mockClient.AssignTapePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFileSystem", func(t *testing.T) {
        input := &storagegateway.AssociateFileSystemInput{}
        output := &storagegateway.AssociateFileSystemOutput{}

        mockClient.On("AssociateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachVolume", func(t *testing.T) {
        input := &storagegateway.AttachVolumeInput{}
        output := &storagegateway.AttachVolumeOutput{}

        mockClient.On("AttachVolume", ctx, input).Return(output, nil)

        result, err := mockClient.AttachVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelArchival", func(t *testing.T) {
        input := &storagegateway.CancelArchivalInput{}
        output := &storagegateway.CancelArchivalOutput{}

        mockClient.On("CancelArchival", ctx, input).Return(output, nil)

        result, err := mockClient.CancelArchival(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelRetrieval", func(t *testing.T) {
        input := &storagegateway.CancelRetrievalInput{}
        output := &storagegateway.CancelRetrievalOutput{}

        mockClient.On("CancelRetrieval", ctx, input).Return(output, nil)

        result, err := mockClient.CancelRetrieval(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCachediSCSIVolume", func(t *testing.T) {
        input := &storagegateway.CreateCachediSCSIVolumeInput{}
        output := &storagegateway.CreateCachediSCSIVolumeOutput{}

        mockClient.On("CreateCachediSCSIVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCachediSCSIVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNFSFileShare", func(t *testing.T) {
        input := &storagegateway.CreateNFSFileShareInput{}
        output := &storagegateway.CreateNFSFileShareOutput{}

        mockClient.On("CreateNFSFileShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNFSFileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSMBFileShare", func(t *testing.T) {
        input := &storagegateway.CreateSMBFileShareInput{}
        output := &storagegateway.CreateSMBFileShareOutput{}

        mockClient.On("CreateSMBFileShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSMBFileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &storagegateway.CreateSnapshotInput{}
        output := &storagegateway.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshotFromVolumeRecoveryPoint", func(t *testing.T) {
        input := &storagegateway.CreateSnapshotFromVolumeRecoveryPointInput{}
        output := &storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput{}

        mockClient.On("CreateSnapshotFromVolumeRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshotFromVolumeRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorediSCSIVolume", func(t *testing.T) {
        input := &storagegateway.CreateStorediSCSIVolumeInput{}
        output := &storagegateway.CreateStorediSCSIVolumeOutput{}

        mockClient.On("CreateStorediSCSIVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorediSCSIVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTapePool", func(t *testing.T) {
        input := &storagegateway.CreateTapePoolInput{}
        output := &storagegateway.CreateTapePoolOutput{}

        mockClient.On("CreateTapePool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTapePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTapeWithBarcode", func(t *testing.T) {
        input := &storagegateway.CreateTapeWithBarcodeInput{}
        output := &storagegateway.CreateTapeWithBarcodeOutput{}

        mockClient.On("CreateTapeWithBarcode", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTapeWithBarcode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTapes", func(t *testing.T) {
        input := &storagegateway.CreateTapesInput{}
        output := &storagegateway.CreateTapesOutput{}

        mockClient.On("CreateTapes", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTapes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAutomaticTapeCreationPolicy", func(t *testing.T) {
        input := &storagegateway.DeleteAutomaticTapeCreationPolicyInput{}
        output := &storagegateway.DeleteAutomaticTapeCreationPolicyOutput{}

        mockClient.On("DeleteAutomaticTapeCreationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAutomaticTapeCreationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBandwidthRateLimit", func(t *testing.T) {
        input := &storagegateway.DeleteBandwidthRateLimitInput{}
        output := &storagegateway.DeleteBandwidthRateLimitOutput{}

        mockClient.On("DeleteBandwidthRateLimit", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBandwidthRateLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChapCredentials", func(t *testing.T) {
        input := &storagegateway.DeleteChapCredentialsInput{}
        output := &storagegateway.DeleteChapCredentialsOutput{}

        mockClient.On("DeleteChapCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChapCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFileShare", func(t *testing.T) {
        input := &storagegateway.DeleteFileShareInput{}
        output := &storagegateway.DeleteFileShareOutput{}

        mockClient.On("DeleteFileShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGateway", func(t *testing.T) {
        input := &storagegateway.DeleteGatewayInput{}
        output := &storagegateway.DeleteGatewayOutput{}

        mockClient.On("DeleteGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshotSchedule", func(t *testing.T) {
        input := &storagegateway.DeleteSnapshotScheduleInput{}
        output := &storagegateway.DeleteSnapshotScheduleOutput{}

        mockClient.On("DeleteSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTape", func(t *testing.T) {
        input := &storagegateway.DeleteTapeInput{}
        output := &storagegateway.DeleteTapeOutput{}

        mockClient.On("DeleteTape", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTape(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTapeArchive", func(t *testing.T) {
        input := &storagegateway.DeleteTapeArchiveInput{}
        output := &storagegateway.DeleteTapeArchiveOutput{}

        mockClient.On("DeleteTapeArchive", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTapeArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTapePool", func(t *testing.T) {
        input := &storagegateway.DeleteTapePoolInput{}
        output := &storagegateway.DeleteTapePoolOutput{}

        mockClient.On("DeleteTapePool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTapePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVolume", func(t *testing.T) {
        input := &storagegateway.DeleteVolumeInput{}
        output := &storagegateway.DeleteVolumeOutput{}

        mockClient.On("DeleteVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAvailabilityMonitorTest", func(t *testing.T) {
        input := &storagegateway.DescribeAvailabilityMonitorTestInput{}
        output := &storagegateway.DescribeAvailabilityMonitorTestOutput{}

        mockClient.On("DescribeAvailabilityMonitorTest", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAvailabilityMonitorTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBandwidthRateLimit", func(t *testing.T) {
        input := &storagegateway.DescribeBandwidthRateLimitInput{}
        output := &storagegateway.DescribeBandwidthRateLimitOutput{}

        mockClient.On("DescribeBandwidthRateLimit", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBandwidthRateLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBandwidthRateLimitSchedule", func(t *testing.T) {
        input := &storagegateway.DescribeBandwidthRateLimitScheduleInput{}
        output := &storagegateway.DescribeBandwidthRateLimitScheduleOutput{}

        mockClient.On("DescribeBandwidthRateLimitSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBandwidthRateLimitSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCache", func(t *testing.T) {
        input := &storagegateway.DescribeCacheInput{}
        output := &storagegateway.DescribeCacheOutput{}

        mockClient.On("DescribeCache", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCachediSCSIVolumes", func(t *testing.T) {
        input := &storagegateway.DescribeCachediSCSIVolumesInput{}
        output := &storagegateway.DescribeCachediSCSIVolumesOutput{}

        mockClient.On("DescribeCachediSCSIVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCachediSCSIVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChapCredentials", func(t *testing.T) {
        input := &storagegateway.DescribeChapCredentialsInput{}
        output := &storagegateway.DescribeChapCredentialsOutput{}

        mockClient.On("DescribeChapCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChapCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileSystemAssociations", func(t *testing.T) {
        input := &storagegateway.DescribeFileSystemAssociationsInput{}
        output := &storagegateway.DescribeFileSystemAssociationsOutput{}

        mockClient.On("DescribeFileSystemAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileSystemAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGatewayInformation", func(t *testing.T) {
        input := &storagegateway.DescribeGatewayInformationInput{}
        output := &storagegateway.DescribeGatewayInformationOutput{}

        mockClient.On("DescribeGatewayInformation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGatewayInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceStartTime", func(t *testing.T) {
        input := &storagegateway.DescribeMaintenanceStartTimeInput{}
        output := &storagegateway.DescribeMaintenanceStartTimeOutput{}

        mockClient.On("DescribeMaintenanceStartTime", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceStartTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNFSFileShares", func(t *testing.T) {
        input := &storagegateway.DescribeNFSFileSharesInput{}
        output := &storagegateway.DescribeNFSFileSharesOutput{}

        mockClient.On("DescribeNFSFileShares", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNFSFileShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSMBFileShares", func(t *testing.T) {
        input := &storagegateway.DescribeSMBFileSharesInput{}
        output := &storagegateway.DescribeSMBFileSharesOutput{}

        mockClient.On("DescribeSMBFileShares", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSMBFileShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSMBSettings", func(t *testing.T) {
        input := &storagegateway.DescribeSMBSettingsInput{}
        output := &storagegateway.DescribeSMBSettingsOutput{}

        mockClient.On("DescribeSMBSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSMBSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshotSchedule", func(t *testing.T) {
        input := &storagegateway.DescribeSnapshotScheduleInput{}
        output := &storagegateway.DescribeSnapshotScheduleOutput{}

        mockClient.On("DescribeSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorediSCSIVolumes", func(t *testing.T) {
        input := &storagegateway.DescribeStorediSCSIVolumesInput{}
        output := &storagegateway.DescribeStorediSCSIVolumesOutput{}

        mockClient.On("DescribeStorediSCSIVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorediSCSIVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTapeArchives", func(t *testing.T) {
        input := &storagegateway.DescribeTapeArchivesInput{}
        output := &storagegateway.DescribeTapeArchivesOutput{}

        mockClient.On("DescribeTapeArchives", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTapeArchives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTapeRecoveryPoints", func(t *testing.T) {
        input := &storagegateway.DescribeTapeRecoveryPointsInput{}
        output := &storagegateway.DescribeTapeRecoveryPointsOutput{}

        mockClient.On("DescribeTapeRecoveryPoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTapeRecoveryPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTapes", func(t *testing.T) {
        input := &storagegateway.DescribeTapesInput{}
        output := &storagegateway.DescribeTapesOutput{}

        mockClient.On("DescribeTapes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTapes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUploadBuffer", func(t *testing.T) {
        input := &storagegateway.DescribeUploadBufferInput{}
        output := &storagegateway.DescribeUploadBufferOutput{}

        mockClient.On("DescribeUploadBuffer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUploadBuffer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVTLDevices", func(t *testing.T) {
        input := &storagegateway.DescribeVTLDevicesInput{}
        output := &storagegateway.DescribeVTLDevicesOutput{}

        mockClient.On("DescribeVTLDevices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVTLDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkingStorage", func(t *testing.T) {
        input := &storagegateway.DescribeWorkingStorageInput{}
        output := &storagegateway.DescribeWorkingStorageOutput{}

        mockClient.On("DescribeWorkingStorage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkingStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachVolume", func(t *testing.T) {
        input := &storagegateway.DetachVolumeInput{}
        output := &storagegateway.DetachVolumeOutput{}

        mockClient.On("DetachVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DetachVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableGateway", func(t *testing.T) {
        input := &storagegateway.DisableGatewayInput{}
        output := &storagegateway.DisableGatewayOutput{}

        mockClient.On("DisableGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DisableGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFileSystem", func(t *testing.T) {
        input := &storagegateway.DisassociateFileSystemInput{}
        output := &storagegateway.DisassociateFileSystemOutput{}

        mockClient.On("DisassociateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestJoinDomain", func(t *testing.T) {
        input := &storagegateway.JoinDomainInput{}
        output := &storagegateway.JoinDomainOutput{}

        mockClient.On("JoinDomain", ctx, input).Return(output, nil)

        result, err := mockClient.JoinDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutomaticTapeCreationPolicies", func(t *testing.T) {
        input := &storagegateway.ListAutomaticTapeCreationPoliciesInput{}
        output := &storagegateway.ListAutomaticTapeCreationPoliciesOutput{}

        mockClient.On("ListAutomaticTapeCreationPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutomaticTapeCreationPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFileShares", func(t *testing.T) {
        input := &storagegateway.ListFileSharesInput{}
        output := &storagegateway.ListFileSharesOutput{}

        mockClient.On("ListFileShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListFileShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFileSystemAssociations", func(t *testing.T) {
        input := &storagegateway.ListFileSystemAssociationsInput{}
        output := &storagegateway.ListFileSystemAssociationsOutput{}

        mockClient.On("ListFileSystemAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFileSystemAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGateways", func(t *testing.T) {
        input := &storagegateway.ListGatewaysInput{}
        output := &storagegateway.ListGatewaysOutput{}

        mockClient.On("ListGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLocalDisks", func(t *testing.T) {
        input := &storagegateway.ListLocalDisksInput{}
        output := &storagegateway.ListLocalDisksOutput{}

        mockClient.On("ListLocalDisks", ctx, input).Return(output, nil)

        result, err := mockClient.ListLocalDisks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &storagegateway.ListTagsForResourceInput{}
        output := &storagegateway.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTapePools", func(t *testing.T) {
        input := &storagegateway.ListTapePoolsInput{}
        output := &storagegateway.ListTapePoolsOutput{}

        mockClient.On("ListTapePools", ctx, input).Return(output, nil)

        result, err := mockClient.ListTapePools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTapes", func(t *testing.T) {
        input := &storagegateway.ListTapesInput{}
        output := &storagegateway.ListTapesOutput{}

        mockClient.On("ListTapes", ctx, input).Return(output, nil)

        result, err := mockClient.ListTapes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVolumeInitiators", func(t *testing.T) {
        input := &storagegateway.ListVolumeInitiatorsInput{}
        output := &storagegateway.ListVolumeInitiatorsOutput{}

        mockClient.On("ListVolumeInitiators", ctx, input).Return(output, nil)

        result, err := mockClient.ListVolumeInitiators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVolumeRecoveryPoints", func(t *testing.T) {
        input := &storagegateway.ListVolumeRecoveryPointsInput{}
        output := &storagegateway.ListVolumeRecoveryPointsOutput{}

        mockClient.On("ListVolumeRecoveryPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListVolumeRecoveryPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVolumes", func(t *testing.T) {
        input := &storagegateway.ListVolumesInput{}
        output := &storagegateway.ListVolumesOutput{}

        mockClient.On("ListVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.ListVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyWhenUploaded", func(t *testing.T) {
        input := &storagegateway.NotifyWhenUploadedInput{}
        output := &storagegateway.NotifyWhenUploadedOutput{}

        mockClient.On("NotifyWhenUploaded", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyWhenUploaded(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRefreshCache", func(t *testing.T) {
        input := &storagegateway.RefreshCacheInput{}
        output := &storagegateway.RefreshCacheOutput{}

        mockClient.On("RefreshCache", ctx, input).Return(output, nil)

        result, err := mockClient.RefreshCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &storagegateway.RemoveTagsFromResourceInput{}
        output := &storagegateway.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetCache", func(t *testing.T) {
        input := &storagegateway.ResetCacheInput{}
        output := &storagegateway.ResetCacheOutput{}

        mockClient.On("ResetCache", ctx, input).Return(output, nil)

        result, err := mockClient.ResetCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieveTapeArchive", func(t *testing.T) {
        input := &storagegateway.RetrieveTapeArchiveInput{}
        output := &storagegateway.RetrieveTapeArchiveOutput{}

        mockClient.On("RetrieveTapeArchive", ctx, input).Return(output, nil)

        result, err := mockClient.RetrieveTapeArchive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieveTapeRecoveryPoint", func(t *testing.T) {
        input := &storagegateway.RetrieveTapeRecoveryPointInput{}
        output := &storagegateway.RetrieveTapeRecoveryPointOutput{}

        mockClient.On("RetrieveTapeRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.RetrieveTapeRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLocalConsolePassword", func(t *testing.T) {
        input := &storagegateway.SetLocalConsolePasswordInput{}
        output := &storagegateway.SetLocalConsolePasswordOutput{}

        mockClient.On("SetLocalConsolePassword", ctx, input).Return(output, nil)

        result, err := mockClient.SetLocalConsolePassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSMBGuestPassword", func(t *testing.T) {
        input := &storagegateway.SetSMBGuestPasswordInput{}
        output := &storagegateway.SetSMBGuestPasswordOutput{}

        mockClient.On("SetSMBGuestPassword", ctx, input).Return(output, nil)

        result, err := mockClient.SetSMBGuestPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestShutdownGateway", func(t *testing.T) {
        input := &storagegateway.ShutdownGatewayInput{}
        output := &storagegateway.ShutdownGatewayOutput{}

        mockClient.On("ShutdownGateway", ctx, input).Return(output, nil)

        result, err := mockClient.ShutdownGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAvailabilityMonitorTest", func(t *testing.T) {
        input := &storagegateway.StartAvailabilityMonitorTestInput{}
        output := &storagegateway.StartAvailabilityMonitorTestOutput{}

        mockClient.On("StartAvailabilityMonitorTest", ctx, input).Return(output, nil)

        result, err := mockClient.StartAvailabilityMonitorTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartGateway", func(t *testing.T) {
        input := &storagegateway.StartGatewayInput{}
        output := &storagegateway.StartGatewayOutput{}

        mockClient.On("StartGateway", ctx, input).Return(output, nil)

        result, err := mockClient.StartGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAutomaticTapeCreationPolicy", func(t *testing.T) {
        input := &storagegateway.UpdateAutomaticTapeCreationPolicyInput{}
        output := &storagegateway.UpdateAutomaticTapeCreationPolicyOutput{}

        mockClient.On("UpdateAutomaticTapeCreationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAutomaticTapeCreationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBandwidthRateLimit", func(t *testing.T) {
        input := &storagegateway.UpdateBandwidthRateLimitInput{}
        output := &storagegateway.UpdateBandwidthRateLimitOutput{}

        mockClient.On("UpdateBandwidthRateLimit", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBandwidthRateLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBandwidthRateLimitSchedule", func(t *testing.T) {
        input := &storagegateway.UpdateBandwidthRateLimitScheduleInput{}
        output := &storagegateway.UpdateBandwidthRateLimitScheduleOutput{}

        mockClient.On("UpdateBandwidthRateLimitSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBandwidthRateLimitSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChapCredentials", func(t *testing.T) {
        input := &storagegateway.UpdateChapCredentialsInput{}
        output := &storagegateway.UpdateChapCredentialsOutput{}

        mockClient.On("UpdateChapCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChapCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFileSystemAssociation", func(t *testing.T) {
        input := &storagegateway.UpdateFileSystemAssociationInput{}
        output := &storagegateway.UpdateFileSystemAssociationOutput{}

        mockClient.On("UpdateFileSystemAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFileSystemAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayInformation", func(t *testing.T) {
        input := &storagegateway.UpdateGatewayInformationInput{}
        output := &storagegateway.UpdateGatewayInformationOutput{}

        mockClient.On("UpdateGatewayInformation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewaySoftwareNow", func(t *testing.T) {
        input := &storagegateway.UpdateGatewaySoftwareNowInput{}
        output := &storagegateway.UpdateGatewaySoftwareNowOutput{}

        mockClient.On("UpdateGatewaySoftwareNow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewaySoftwareNow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMaintenanceStartTime", func(t *testing.T) {
        input := &storagegateway.UpdateMaintenanceStartTimeInput{}
        output := &storagegateway.UpdateMaintenanceStartTimeOutput{}

        mockClient.On("UpdateMaintenanceStartTime", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMaintenanceStartTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNFSFileShare", func(t *testing.T) {
        input := &storagegateway.UpdateNFSFileShareInput{}
        output := &storagegateway.UpdateNFSFileShareOutput{}

        mockClient.On("UpdateNFSFileShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNFSFileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSMBFileShare", func(t *testing.T) {
        input := &storagegateway.UpdateSMBFileShareInput{}
        output := &storagegateway.UpdateSMBFileShareOutput{}

        mockClient.On("UpdateSMBFileShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSMBFileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSMBFileShareVisibility", func(t *testing.T) {
        input := &storagegateway.UpdateSMBFileShareVisibilityInput{}
        output := &storagegateway.UpdateSMBFileShareVisibilityOutput{}

        mockClient.On("UpdateSMBFileShareVisibility", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSMBFileShareVisibility(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSMBLocalGroups", func(t *testing.T) {
        input := &storagegateway.UpdateSMBLocalGroupsInput{}
        output := &storagegateway.UpdateSMBLocalGroupsOutput{}

        mockClient.On("UpdateSMBLocalGroups", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSMBLocalGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSMBSecurityStrategy", func(t *testing.T) {
        input := &storagegateway.UpdateSMBSecurityStrategyInput{}
        output := &storagegateway.UpdateSMBSecurityStrategyOutput{}

        mockClient.On("UpdateSMBSecurityStrategy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSMBSecurityStrategy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSnapshotSchedule", func(t *testing.T) {
        input := &storagegateway.UpdateSnapshotScheduleInput{}
        output := &storagegateway.UpdateSnapshotScheduleOutput{}

        mockClient.On("UpdateSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVTLDeviceType", func(t *testing.T) {
        input := &storagegateway.UpdateVTLDeviceTypeInput{}
        output := &storagegateway.UpdateVTLDeviceTypeOutput{}

        mockClient.On("UpdateVTLDeviceType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVTLDeviceType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
