package cloudhsmv2_test

// tests for the cloudhsmv2 service interface for this ../../../service/cloudhsmv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudhsmv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudhsmv2/cloudhsmv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudhsmv2ServiceCanBeMocked(t *testing.T) {
	var iface cloudhsmv2_iface.IClient
	iface = &cloudhsmv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudhsmv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyBackupToRegion", func(t *testing.T) {
        input := &cloudhsmv2.CopyBackupToRegionInput{}
        output := &cloudhsmv2.CopyBackupToRegionOutput{}

        mockClient.On("CopyBackupToRegion", ctx, input).Return(output, nil)

        result, err := mockClient.CopyBackupToRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &cloudhsmv2.CreateClusterInput{}
        output := &cloudhsmv2.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHsm", func(t *testing.T) {
        input := &cloudhsmv2.CreateHsmInput{}
        output := &cloudhsmv2.CreateHsmOutput{}

        mockClient.On("CreateHsm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackup", func(t *testing.T) {
        input := &cloudhsmv2.DeleteBackupInput{}
        output := &cloudhsmv2.DeleteBackupOutput{}

        mockClient.On("DeleteBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &cloudhsmv2.DeleteClusterInput{}
        output := &cloudhsmv2.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHsm", func(t *testing.T) {
        input := &cloudhsmv2.DeleteHsmInput{}
        output := &cloudhsmv2.DeleteHsmOutput{}

        mockClient.On("DeleteHsm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &cloudhsmv2.DeleteResourcePolicyInput{}
        output := &cloudhsmv2.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackups", func(t *testing.T) {
        input := &cloudhsmv2.DescribeBackupsInput{}
        output := &cloudhsmv2.DescribeBackupsOutput{}

        mockClient.On("DescribeBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusters", func(t *testing.T) {
        input := &cloudhsmv2.DescribeClustersInput{}
        output := &cloudhsmv2.DescribeClustersOutput{}

        mockClient.On("DescribeClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &cloudhsmv2.GetResourcePolicyInput{}
        output := &cloudhsmv2.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitializeCluster", func(t *testing.T) {
        input := &cloudhsmv2.InitializeClusterInput{}
        output := &cloudhsmv2.InitializeClusterOutput{}

        mockClient.On("InitializeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.InitializeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &cloudhsmv2.ListTagsInput{}
        output := &cloudhsmv2.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyBackupAttributes", func(t *testing.T) {
        input := &cloudhsmv2.ModifyBackupAttributesInput{}
        output := &cloudhsmv2.ModifyBackupAttributesOutput{}

        mockClient.On("ModifyBackupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyBackupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCluster", func(t *testing.T) {
        input := &cloudhsmv2.ModifyClusterInput{}
        output := &cloudhsmv2.ModifyClusterOutput{}

        mockClient.On("ModifyCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &cloudhsmv2.PutResourcePolicyInput{}
        output := &cloudhsmv2.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreBackup", func(t *testing.T) {
        input := &cloudhsmv2.RestoreBackupInput{}
        output := &cloudhsmv2.RestoreBackupOutput{}

        mockClient.On("RestoreBackup", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloudhsmv2.TagResourceInput{}
        output := &cloudhsmv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloudhsmv2.UntagResourceInput{}
        output := &cloudhsmv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
