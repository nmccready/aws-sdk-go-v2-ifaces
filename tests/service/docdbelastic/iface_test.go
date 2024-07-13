package docdbelastic_test

// tests for the docdbelastic service interface for this ../../../service/docdbelastic/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdbelastic"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/docdbelastic/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/docdbelastic/docdbelastic_iface"
	"github.com/stretchr/testify/assert"
)

func TestDocdbelasticServiceCanBeMocked(t *testing.T) {
	var iface docdbelastic_iface.IClient
	iface = &docdbelastic.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := docdbelastic.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyClusterSnapshot", func(t *testing.T) {
        input := &docdbelastic.CopyClusterSnapshotInput{}
        output := &docdbelastic.CopyClusterSnapshotOutput{}

        mockClient.On("CopyClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &docdbelastic.CreateClusterInput{}
        output := &docdbelastic.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterSnapshot", func(t *testing.T) {
        input := &docdbelastic.CreateClusterSnapshotInput{}
        output := &docdbelastic.CreateClusterSnapshotOutput{}

        mockClient.On("CreateClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &docdbelastic.DeleteClusterInput{}
        output := &docdbelastic.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterSnapshot", func(t *testing.T) {
        input := &docdbelastic.DeleteClusterSnapshotInput{}
        output := &docdbelastic.DeleteClusterSnapshotOutput{}

        mockClient.On("DeleteClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCluster", func(t *testing.T) {
        input := &docdbelastic.GetClusterInput{}
        output := &docdbelastic.GetClusterOutput{}

        mockClient.On("GetCluster", ctx, input).Return(output, nil)

        result, err := mockClient.GetCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClusterSnapshot", func(t *testing.T) {
        input := &docdbelastic.GetClusterSnapshotInput{}
        output := &docdbelastic.GetClusterSnapshotOutput{}

        mockClient.On("GetClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusterSnapshots", func(t *testing.T) {
        input := &docdbelastic.ListClusterSnapshotsInput{}
        output := &docdbelastic.ListClusterSnapshotsOutput{}

        mockClient.On("ListClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &docdbelastic.ListClustersInput{}
        output := &docdbelastic.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &docdbelastic.ListTagsForResourceInput{}
        output := &docdbelastic.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreClusterFromSnapshot", func(t *testing.T) {
        input := &docdbelastic.RestoreClusterFromSnapshotInput{}
        output := &docdbelastic.RestoreClusterFromSnapshotOutput{}

        mockClient.On("RestoreClusterFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreClusterFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCluster", func(t *testing.T) {
        input := &docdbelastic.StartClusterInput{}
        output := &docdbelastic.StartClusterOutput{}

        mockClient.On("StartCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StartCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCluster", func(t *testing.T) {
        input := &docdbelastic.StopClusterInput{}
        output := &docdbelastic.StopClusterOutput{}

        mockClient.On("StopCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StopCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &docdbelastic.TagResourceInput{}
        output := &docdbelastic.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &docdbelastic.UntagResourceInput{}
        output := &docdbelastic.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &docdbelastic.UpdateClusterInput{}
        output := &docdbelastic.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
