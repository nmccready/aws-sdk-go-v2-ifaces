package finspace_test

// tests for the finspace service interface for this ../../../service/finspace/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/finspace"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/finspace/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/finspace/finspace_iface"
	"github.com/stretchr/testify/assert"
)

func TestFinspaceServiceCanBeMocked(t *testing.T) {
	var iface finspace_iface.IClient
	iface = &finspace.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := finspace.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &finspace.CreateEnvironmentInput{}
        output := &finspace.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxChangeset", func(t *testing.T) {
        input := &finspace.CreateKxChangesetInput{}
        output := &finspace.CreateKxChangesetOutput{}

        mockClient.On("CreateKxChangeset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxChangeset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxCluster", func(t *testing.T) {
        input := &finspace.CreateKxClusterInput{}
        output := &finspace.CreateKxClusterOutput{}

        mockClient.On("CreateKxCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxDatabase", func(t *testing.T) {
        input := &finspace.CreateKxDatabaseInput{}
        output := &finspace.CreateKxDatabaseOutput{}

        mockClient.On("CreateKxDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxDataview", func(t *testing.T) {
        input := &finspace.CreateKxDataviewInput{}
        output := &finspace.CreateKxDataviewOutput{}

        mockClient.On("CreateKxDataview", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxDataview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxEnvironment", func(t *testing.T) {
        input := &finspace.CreateKxEnvironmentInput{}
        output := &finspace.CreateKxEnvironmentOutput{}

        mockClient.On("CreateKxEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxScalingGroup", func(t *testing.T) {
        input := &finspace.CreateKxScalingGroupInput{}
        output := &finspace.CreateKxScalingGroupOutput{}

        mockClient.On("CreateKxScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxUser", func(t *testing.T) {
        input := &finspace.CreateKxUserInput{}
        output := &finspace.CreateKxUserOutput{}

        mockClient.On("CreateKxUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKxVolume", func(t *testing.T) {
        input := &finspace.CreateKxVolumeInput{}
        output := &finspace.CreateKxVolumeOutput{}

        mockClient.On("CreateKxVolume", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKxVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &finspace.DeleteEnvironmentInput{}
        output := &finspace.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxCluster", func(t *testing.T) {
        input := &finspace.DeleteKxClusterInput{}
        output := &finspace.DeleteKxClusterOutput{}

        mockClient.On("DeleteKxCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxClusterNode", func(t *testing.T) {
        input := &finspace.DeleteKxClusterNodeInput{}
        output := &finspace.DeleteKxClusterNodeOutput{}

        mockClient.On("DeleteKxClusterNode", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxClusterNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxDatabase", func(t *testing.T) {
        input := &finspace.DeleteKxDatabaseInput{}
        output := &finspace.DeleteKxDatabaseOutput{}

        mockClient.On("DeleteKxDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxDataview", func(t *testing.T) {
        input := &finspace.DeleteKxDataviewInput{}
        output := &finspace.DeleteKxDataviewOutput{}

        mockClient.On("DeleteKxDataview", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxDataview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxEnvironment", func(t *testing.T) {
        input := &finspace.DeleteKxEnvironmentInput{}
        output := &finspace.DeleteKxEnvironmentOutput{}

        mockClient.On("DeleteKxEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxScalingGroup", func(t *testing.T) {
        input := &finspace.DeleteKxScalingGroupInput{}
        output := &finspace.DeleteKxScalingGroupOutput{}

        mockClient.On("DeleteKxScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxUser", func(t *testing.T) {
        input := &finspace.DeleteKxUserInput{}
        output := &finspace.DeleteKxUserOutput{}

        mockClient.On("DeleteKxUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKxVolume", func(t *testing.T) {
        input := &finspace.DeleteKxVolumeInput{}
        output := &finspace.DeleteKxVolumeOutput{}

        mockClient.On("DeleteKxVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKxVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &finspace.GetEnvironmentInput{}
        output := &finspace.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxChangeset", func(t *testing.T) {
        input := &finspace.GetKxChangesetInput{}
        output := &finspace.GetKxChangesetOutput{}

        mockClient.On("GetKxChangeset", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxChangeset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxCluster", func(t *testing.T) {
        input := &finspace.GetKxClusterInput{}
        output := &finspace.GetKxClusterOutput{}

        mockClient.On("GetKxCluster", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxConnectionString", func(t *testing.T) {
        input := &finspace.GetKxConnectionStringInput{}
        output := &finspace.GetKxConnectionStringOutput{}

        mockClient.On("GetKxConnectionString", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxConnectionString(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxDatabase", func(t *testing.T) {
        input := &finspace.GetKxDatabaseInput{}
        output := &finspace.GetKxDatabaseOutput{}

        mockClient.On("GetKxDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxDataview", func(t *testing.T) {
        input := &finspace.GetKxDataviewInput{}
        output := &finspace.GetKxDataviewOutput{}

        mockClient.On("GetKxDataview", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxDataview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxEnvironment", func(t *testing.T) {
        input := &finspace.GetKxEnvironmentInput{}
        output := &finspace.GetKxEnvironmentOutput{}

        mockClient.On("GetKxEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxScalingGroup", func(t *testing.T) {
        input := &finspace.GetKxScalingGroupInput{}
        output := &finspace.GetKxScalingGroupOutput{}

        mockClient.On("GetKxScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxUser", func(t *testing.T) {
        input := &finspace.GetKxUserInput{}
        output := &finspace.GetKxUserOutput{}

        mockClient.On("GetKxUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKxVolume", func(t *testing.T) {
        input := &finspace.GetKxVolumeInput{}
        output := &finspace.GetKxVolumeOutput{}

        mockClient.On("GetKxVolume", ctx, input).Return(output, nil)

        result, err := mockClient.GetKxVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &finspace.ListEnvironmentsInput{}
        output := &finspace.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxChangesets", func(t *testing.T) {
        input := &finspace.ListKxChangesetsInput{}
        output := &finspace.ListKxChangesetsOutput{}

        mockClient.On("ListKxChangesets", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxChangesets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxClusterNodes", func(t *testing.T) {
        input := &finspace.ListKxClusterNodesInput{}
        output := &finspace.ListKxClusterNodesOutput{}

        mockClient.On("ListKxClusterNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxClusterNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxClusters", func(t *testing.T) {
        input := &finspace.ListKxClustersInput{}
        output := &finspace.ListKxClustersOutput{}

        mockClient.On("ListKxClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxDatabases", func(t *testing.T) {
        input := &finspace.ListKxDatabasesInput{}
        output := &finspace.ListKxDatabasesOutput{}

        mockClient.On("ListKxDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxDataviews", func(t *testing.T) {
        input := &finspace.ListKxDataviewsInput{}
        output := &finspace.ListKxDataviewsOutput{}

        mockClient.On("ListKxDataviews", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxDataviews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxEnvironments", func(t *testing.T) {
        input := &finspace.ListKxEnvironmentsInput{}
        output := &finspace.ListKxEnvironmentsOutput{}

        mockClient.On("ListKxEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxScalingGroups", func(t *testing.T) {
        input := &finspace.ListKxScalingGroupsInput{}
        output := &finspace.ListKxScalingGroupsOutput{}

        mockClient.On("ListKxScalingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxScalingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxUsers", func(t *testing.T) {
        input := &finspace.ListKxUsersInput{}
        output := &finspace.ListKxUsersOutput{}

        mockClient.On("ListKxUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKxVolumes", func(t *testing.T) {
        input := &finspace.ListKxVolumesInput{}
        output := &finspace.ListKxVolumesOutput{}

        mockClient.On("ListKxVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.ListKxVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &finspace.ListTagsForResourceInput{}
        output := &finspace.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &finspace.TagResourceInput{}
        output := &finspace.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &finspace.UntagResourceInput{}
        output := &finspace.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &finspace.UpdateEnvironmentInput{}
        output := &finspace.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxClusterCodeConfiguration", func(t *testing.T) {
        input := &finspace.UpdateKxClusterCodeConfigurationInput{}
        output := &finspace.UpdateKxClusterCodeConfigurationOutput{}

        mockClient.On("UpdateKxClusterCodeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxClusterCodeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxClusterDatabases", func(t *testing.T) {
        input := &finspace.UpdateKxClusterDatabasesInput{}
        output := &finspace.UpdateKxClusterDatabasesOutput{}

        mockClient.On("UpdateKxClusterDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxClusterDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxDatabase", func(t *testing.T) {
        input := &finspace.UpdateKxDatabaseInput{}
        output := &finspace.UpdateKxDatabaseOutput{}

        mockClient.On("UpdateKxDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxDataview", func(t *testing.T) {
        input := &finspace.UpdateKxDataviewInput{}
        output := &finspace.UpdateKxDataviewOutput{}

        mockClient.On("UpdateKxDataview", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxDataview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxEnvironment", func(t *testing.T) {
        input := &finspace.UpdateKxEnvironmentInput{}
        output := &finspace.UpdateKxEnvironmentOutput{}

        mockClient.On("UpdateKxEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxEnvironmentNetwork", func(t *testing.T) {
        input := &finspace.UpdateKxEnvironmentNetworkInput{}
        output := &finspace.UpdateKxEnvironmentNetworkOutput{}

        mockClient.On("UpdateKxEnvironmentNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxEnvironmentNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxUser", func(t *testing.T) {
        input := &finspace.UpdateKxUserInput{}
        output := &finspace.UpdateKxUserOutput{}

        mockClient.On("UpdateKxUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKxVolume", func(t *testing.T) {
        input := &finspace.UpdateKxVolumeInput{}
        output := &finspace.UpdateKxVolumeOutput{}

        mockClient.On("UpdateKxVolume", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKxVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
