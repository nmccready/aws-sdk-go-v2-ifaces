package fsx_test

// tests for the fsx service interface for this ../../../service/fsx/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fsx/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fsx/fsx_iface"
	"github.com/stretchr/testify/assert"
)

func TestFsxServiceCanBeMocked(t *testing.T) {
	var iface fsx_iface.IClient
	iface = &fsx.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := fsx.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFileSystemAliases", func(t *testing.T) {
        input := &fsx.AssociateFileSystemAliasesInput{}
        output := &fsx.AssociateFileSystemAliasesOutput{}

        mockClient.On("AssociateFileSystemAliases", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFileSystemAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDataRepositoryTask", func(t *testing.T) {
        input := &fsx.CancelDataRepositoryTaskInput{}
        output := &fsx.CancelDataRepositoryTaskOutput{}

        mockClient.On("CancelDataRepositoryTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDataRepositoryTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyBackup", func(t *testing.T) {
        input := &fsx.CopyBackupInput{}
        output := &fsx.CopyBackupOutput{}

        mockClient.On("CopyBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopySnapshotAndUpdateVolume", func(t *testing.T) {
        input := &fsx.CopySnapshotAndUpdateVolumeInput{}
        output := &fsx.CopySnapshotAndUpdateVolumeOutput{}

        mockClient.On("CopySnapshotAndUpdateVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CopySnapshotAndUpdateVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackup", func(t *testing.T) {
        input := &fsx.CreateBackupInput{}
        output := &fsx.CreateBackupOutput{}

        mockClient.On("CreateBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataRepositoryAssociation", func(t *testing.T) {
        input := &fsx.CreateDataRepositoryAssociationInput{}
        output := &fsx.CreateDataRepositoryAssociationOutput{}

        mockClient.On("CreateDataRepositoryAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataRepositoryAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataRepositoryTask", func(t *testing.T) {
        input := &fsx.CreateDataRepositoryTaskInput{}
        output := &fsx.CreateDataRepositoryTaskOutput{}

        mockClient.On("CreateDataRepositoryTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataRepositoryTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFileCache", func(t *testing.T) {
        input := &fsx.CreateFileCacheInput{}
        output := &fsx.CreateFileCacheOutput{}

        mockClient.On("CreateFileCache", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFileCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFileSystem", func(t *testing.T) {
        input := &fsx.CreateFileSystemInput{}
        output := &fsx.CreateFileSystemOutput{}

        mockClient.On("CreateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFileSystemFromBackup", func(t *testing.T) {
        input := &fsx.CreateFileSystemFromBackupInput{}
        output := &fsx.CreateFileSystemFromBackupOutput{}

        mockClient.On("CreateFileSystemFromBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFileSystemFromBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &fsx.CreateSnapshotInput{}
        output := &fsx.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorageVirtualMachine", func(t *testing.T) {
        input := &fsx.CreateStorageVirtualMachineInput{}
        output := &fsx.CreateStorageVirtualMachineOutput{}

        mockClient.On("CreateStorageVirtualMachine", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorageVirtualMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVolume", func(t *testing.T) {
        input := &fsx.CreateVolumeInput{}
        output := &fsx.CreateVolumeOutput{}

        mockClient.On("CreateVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVolumeFromBackup", func(t *testing.T) {
        input := &fsx.CreateVolumeFromBackupInput{}
        output := &fsx.CreateVolumeFromBackupOutput{}

        mockClient.On("CreateVolumeFromBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVolumeFromBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackup", func(t *testing.T) {
        input := &fsx.DeleteBackupInput{}
        output := &fsx.DeleteBackupOutput{}

        mockClient.On("DeleteBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataRepositoryAssociation", func(t *testing.T) {
        input := &fsx.DeleteDataRepositoryAssociationInput{}
        output := &fsx.DeleteDataRepositoryAssociationOutput{}

        mockClient.On("DeleteDataRepositoryAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataRepositoryAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFileCache", func(t *testing.T) {
        input := &fsx.DeleteFileCacheInput{}
        output := &fsx.DeleteFileCacheOutput{}

        mockClient.On("DeleteFileCache", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFileCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFileSystem", func(t *testing.T) {
        input := &fsx.DeleteFileSystemInput{}
        output := &fsx.DeleteFileSystemOutput{}

        mockClient.On("DeleteFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &fsx.DeleteSnapshotInput{}
        output := &fsx.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageVirtualMachine", func(t *testing.T) {
        input := &fsx.DeleteStorageVirtualMachineInput{}
        output := &fsx.DeleteStorageVirtualMachineOutput{}

        mockClient.On("DeleteStorageVirtualMachine", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageVirtualMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVolume", func(t *testing.T) {
        input := &fsx.DeleteVolumeInput{}
        output := &fsx.DeleteVolumeOutput{}

        mockClient.On("DeleteVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackups", func(t *testing.T) {
        input := &fsx.DescribeBackupsInput{}
        output := &fsx.DescribeBackupsOutput{}

        mockClient.On("DescribeBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataRepositoryAssociations", func(t *testing.T) {
        input := &fsx.DescribeDataRepositoryAssociationsInput{}
        output := &fsx.DescribeDataRepositoryAssociationsOutput{}

        mockClient.On("DescribeDataRepositoryAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataRepositoryAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataRepositoryTasks", func(t *testing.T) {
        input := &fsx.DescribeDataRepositoryTasksInput{}
        output := &fsx.DescribeDataRepositoryTasksOutput{}

        mockClient.On("DescribeDataRepositoryTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataRepositoryTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileCaches", func(t *testing.T) {
        input := &fsx.DescribeFileCachesInput{}
        output := &fsx.DescribeFileCachesOutput{}

        mockClient.On("DescribeFileCaches", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileCaches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileSystemAliases", func(t *testing.T) {
        input := &fsx.DescribeFileSystemAliasesInput{}
        output := &fsx.DescribeFileSystemAliasesOutput{}

        mockClient.On("DescribeFileSystemAliases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileSystemAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileSystems", func(t *testing.T) {
        input := &fsx.DescribeFileSystemsInput{}
        output := &fsx.DescribeFileSystemsOutput{}

        mockClient.On("DescribeFileSystems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileSystems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSharedVpcConfiguration", func(t *testing.T) {
        input := &fsx.DescribeSharedVpcConfigurationInput{}
        output := &fsx.DescribeSharedVpcConfigurationOutput{}

        mockClient.On("DescribeSharedVpcConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSharedVpcConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshots", func(t *testing.T) {
        input := &fsx.DescribeSnapshotsInput{}
        output := &fsx.DescribeSnapshotsOutput{}

        mockClient.On("DescribeSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorageVirtualMachines", func(t *testing.T) {
        input := &fsx.DescribeStorageVirtualMachinesInput{}
        output := &fsx.DescribeStorageVirtualMachinesOutput{}

        mockClient.On("DescribeStorageVirtualMachines", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorageVirtualMachines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumes", func(t *testing.T) {
        input := &fsx.DescribeVolumesInput{}
        output := &fsx.DescribeVolumesOutput{}

        mockClient.On("DescribeVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFileSystemAliases", func(t *testing.T) {
        input := &fsx.DisassociateFileSystemAliasesInput{}
        output := &fsx.DisassociateFileSystemAliasesOutput{}

        mockClient.On("DisassociateFileSystemAliases", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFileSystemAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &fsx.ListTagsForResourceInput{}
        output := &fsx.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseFileSystemNfsV3Locks", func(t *testing.T) {
        input := &fsx.ReleaseFileSystemNfsV3LocksInput{}
        output := &fsx.ReleaseFileSystemNfsV3LocksOutput{}

        mockClient.On("ReleaseFileSystemNfsV3Locks", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseFileSystemNfsV3Locks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreVolumeFromSnapshot", func(t *testing.T) {
        input := &fsx.RestoreVolumeFromSnapshotInput{}
        output := &fsx.RestoreVolumeFromSnapshotOutput{}

        mockClient.On("RestoreVolumeFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreVolumeFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMisconfiguredStateRecovery", func(t *testing.T) {
        input := &fsx.StartMisconfiguredStateRecoveryInput{}
        output := &fsx.StartMisconfiguredStateRecoveryOutput{}

        mockClient.On("StartMisconfiguredStateRecovery", ctx, input).Return(output, nil)

        result, err := mockClient.StartMisconfiguredStateRecovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &fsx.TagResourceInput{}
        output := &fsx.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &fsx.UntagResourceInput{}
        output := &fsx.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataRepositoryAssociation", func(t *testing.T) {
        input := &fsx.UpdateDataRepositoryAssociationInput{}
        output := &fsx.UpdateDataRepositoryAssociationOutput{}

        mockClient.On("UpdateDataRepositoryAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataRepositoryAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFileCache", func(t *testing.T) {
        input := &fsx.UpdateFileCacheInput{}
        output := &fsx.UpdateFileCacheOutput{}

        mockClient.On("UpdateFileCache", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFileCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFileSystem", func(t *testing.T) {
        input := &fsx.UpdateFileSystemInput{}
        output := &fsx.UpdateFileSystemOutput{}

        mockClient.On("UpdateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSharedVpcConfiguration", func(t *testing.T) {
        input := &fsx.UpdateSharedVpcConfigurationInput{}
        output := &fsx.UpdateSharedVpcConfigurationOutput{}

        mockClient.On("UpdateSharedVpcConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSharedVpcConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSnapshot", func(t *testing.T) {
        input := &fsx.UpdateSnapshotInput{}
        output := &fsx.UpdateSnapshotOutput{}

        mockClient.On("UpdateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStorageVirtualMachine", func(t *testing.T) {
        input := &fsx.UpdateStorageVirtualMachineInput{}
        output := &fsx.UpdateStorageVirtualMachineOutput{}

        mockClient.On("UpdateStorageVirtualMachine", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStorageVirtualMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVolume", func(t *testing.T) {
        input := &fsx.UpdateVolumeInput{}
        output := &fsx.UpdateVolumeOutput{}

        mockClient.On("UpdateVolume", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
