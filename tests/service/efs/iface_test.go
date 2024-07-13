package efs_test

// tests for the efs service interface for this ../../../service/efs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/efs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/efs/efs_iface"
	"github.com/stretchr/testify/assert"
)

func TestEfsServiceCanBeMocked(t *testing.T) {
	var iface efs_iface.IClient
	iface = &efs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := efs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPoint", func(t *testing.T) {
        input := &efs.CreateAccessPointInput{}
        output := &efs.CreateAccessPointOutput{}

        mockClient.On("CreateAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFileSystem", func(t *testing.T) {
        input := &efs.CreateFileSystemInput{}
        output := &efs.CreateFileSystemOutput{}

        mockClient.On("CreateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMountTarget", func(t *testing.T) {
        input := &efs.CreateMountTargetInput{}
        output := &efs.CreateMountTargetOutput{}

        mockClient.On("CreateMountTarget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMountTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationConfiguration", func(t *testing.T) {
        input := &efs.CreateReplicationConfigurationInput{}
        output := &efs.CreateReplicationConfigurationOutput{}

        mockClient.On("CreateReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &efs.CreateTagsInput{}
        output := &efs.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPoint", func(t *testing.T) {
        input := &efs.DeleteAccessPointInput{}
        output := &efs.DeleteAccessPointOutput{}

        mockClient.On("DeleteAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFileSystem", func(t *testing.T) {
        input := &efs.DeleteFileSystemInput{}
        output := &efs.DeleteFileSystemOutput{}

        mockClient.On("DeleteFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFileSystemPolicy", func(t *testing.T) {
        input := &efs.DeleteFileSystemPolicyInput{}
        output := &efs.DeleteFileSystemPolicyOutput{}

        mockClient.On("DeleteFileSystemPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFileSystemPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMountTarget", func(t *testing.T) {
        input := &efs.DeleteMountTargetInput{}
        output := &efs.DeleteMountTargetOutput{}

        mockClient.On("DeleteMountTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMountTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationConfiguration", func(t *testing.T) {
        input := &efs.DeleteReplicationConfigurationInput{}
        output := &efs.DeleteReplicationConfigurationOutput{}

        mockClient.On("DeleteReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &efs.DeleteTagsInput{}
        output := &efs.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccessPoints", func(t *testing.T) {
        input := &efs.DescribeAccessPointsInput{}
        output := &efs.DescribeAccessPointsOutput{}

        mockClient.On("DescribeAccessPoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccessPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountPreferences", func(t *testing.T) {
        input := &efs.DescribeAccountPreferencesInput{}
        output := &efs.DescribeAccountPreferencesOutput{}

        mockClient.On("DescribeAccountPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackupPolicy", func(t *testing.T) {
        input := &efs.DescribeBackupPolicyInput{}
        output := &efs.DescribeBackupPolicyOutput{}

        mockClient.On("DescribeBackupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileSystemPolicy", func(t *testing.T) {
        input := &efs.DescribeFileSystemPolicyInput{}
        output := &efs.DescribeFileSystemPolicyOutput{}

        mockClient.On("DescribeFileSystemPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileSystemPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFileSystems", func(t *testing.T) {
        input := &efs.DescribeFileSystemsInput{}
        output := &efs.DescribeFileSystemsOutput{}

        mockClient.On("DescribeFileSystems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFileSystems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLifecycleConfiguration", func(t *testing.T) {
        input := &efs.DescribeLifecycleConfigurationInput{}
        output := &efs.DescribeLifecycleConfigurationOutput{}

        mockClient.On("DescribeLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMountTargetSecurityGroups", func(t *testing.T) {
        input := &efs.DescribeMountTargetSecurityGroupsInput{}
        output := &efs.DescribeMountTargetSecurityGroupsOutput{}

        mockClient.On("DescribeMountTargetSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMountTargetSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMountTargets", func(t *testing.T) {
        input := &efs.DescribeMountTargetsInput{}
        output := &efs.DescribeMountTargetsOutput{}

        mockClient.On("DescribeMountTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMountTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationConfigurations", func(t *testing.T) {
        input := &efs.DescribeReplicationConfigurationsInput{}
        output := &efs.DescribeReplicationConfigurationsOutput{}

        mockClient.On("DescribeReplicationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &efs.DescribeTagsInput{}
        output := &efs.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &efs.ListTagsForResourceInput{}
        output := &efs.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyMountTargetSecurityGroups", func(t *testing.T) {
        input := &efs.ModifyMountTargetSecurityGroupsInput{}
        output := &efs.ModifyMountTargetSecurityGroupsOutput{}

        mockClient.On("ModifyMountTargetSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyMountTargetSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountPreferences", func(t *testing.T) {
        input := &efs.PutAccountPreferencesInput{}
        output := &efs.PutAccountPreferencesOutput{}

        mockClient.On("PutAccountPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBackupPolicy", func(t *testing.T) {
        input := &efs.PutBackupPolicyInput{}
        output := &efs.PutBackupPolicyOutput{}

        mockClient.On("PutBackupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutBackupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFileSystemPolicy", func(t *testing.T) {
        input := &efs.PutFileSystemPolicyInput{}
        output := &efs.PutFileSystemPolicyOutput{}

        mockClient.On("PutFileSystemPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutFileSystemPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLifecycleConfiguration", func(t *testing.T) {
        input := &efs.PutLifecycleConfigurationInput{}
        output := &efs.PutLifecycleConfigurationOutput{}

        mockClient.On("PutLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &efs.TagResourceInput{}
        output := &efs.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &efs.UntagResourceInput{}
        output := &efs.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFileSystem", func(t *testing.T) {
        input := &efs.UpdateFileSystemInput{}
        output := &efs.UpdateFileSystemOutput{}

        mockClient.On("UpdateFileSystem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFileSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFileSystemProtection", func(t *testing.T) {
        input := &efs.UpdateFileSystemProtectionInput{}
        output := &efs.UpdateFileSystemProtectionOutput{}

        mockClient.On("UpdateFileSystemProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFileSystemProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
