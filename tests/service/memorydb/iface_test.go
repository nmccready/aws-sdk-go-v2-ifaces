package memorydb_test

// tests for the memorydb service interface for this ../../../service/memorydb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/memorydb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/memorydb/memorydb_iface"
	"github.com/stretchr/testify/assert"
)

func TestMemorydbServiceCanBeMocked(t *testing.T) {
	var iface memorydb_iface.IClient
	iface = &memorydb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := memorydb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateCluster", func(t *testing.T) {
        input := &memorydb.BatchUpdateClusterInput{}
        output := &memorydb.BatchUpdateClusterOutput{}

        mockClient.On("BatchUpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopySnapshot", func(t *testing.T) {
        input := &memorydb.CopySnapshotInput{}
        output := &memorydb.CopySnapshotOutput{}

        mockClient.On("CopySnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopySnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateACL", func(t *testing.T) {
        input := &memorydb.CreateACLInput{}
        output := &memorydb.CreateACLOutput{}

        mockClient.On("CreateACL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &memorydb.CreateClusterInput{}
        output := &memorydb.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParameterGroup", func(t *testing.T) {
        input := &memorydb.CreateParameterGroupInput{}
        output := &memorydb.CreateParameterGroupOutput{}

        mockClient.On("CreateParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &memorydb.CreateSnapshotInput{}
        output := &memorydb.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubnetGroup", func(t *testing.T) {
        input := &memorydb.CreateSubnetGroupInput{}
        output := &memorydb.CreateSubnetGroupOutput{}

        mockClient.On("CreateSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &memorydb.CreateUserInput{}
        output := &memorydb.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteACL", func(t *testing.T) {
        input := &memorydb.DeleteACLInput{}
        output := &memorydb.DeleteACLOutput{}

        mockClient.On("DeleteACL", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &memorydb.DeleteClusterInput{}
        output := &memorydb.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteParameterGroup", func(t *testing.T) {
        input := &memorydb.DeleteParameterGroupInput{}
        output := &memorydb.DeleteParameterGroupOutput{}

        mockClient.On("DeleteParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &memorydb.DeleteSnapshotInput{}
        output := &memorydb.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubnetGroup", func(t *testing.T) {
        input := &memorydb.DeleteSubnetGroupInput{}
        output := &memorydb.DeleteSubnetGroupOutput{}

        mockClient.On("DeleteSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &memorydb.DeleteUserInput{}
        output := &memorydb.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeACLs", func(t *testing.T) {
        input := &memorydb.DescribeACLsInput{}
        output := &memorydb.DescribeACLsOutput{}

        mockClient.On("DescribeACLs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeACLs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusters", func(t *testing.T) {
        input := &memorydb.DescribeClustersInput{}
        output := &memorydb.DescribeClustersOutput{}

        mockClient.On("DescribeClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineVersions", func(t *testing.T) {
        input := &memorydb.DescribeEngineVersionsInput{}
        output := &memorydb.DescribeEngineVersionsOutput{}

        mockClient.On("DescribeEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &memorydb.DescribeEventsInput{}
        output := &memorydb.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeParameterGroups", func(t *testing.T) {
        input := &memorydb.DescribeParameterGroupsInput{}
        output := &memorydb.DescribeParameterGroupsOutput{}

        mockClient.On("DescribeParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeParameters", func(t *testing.T) {
        input := &memorydb.DescribeParametersInput{}
        output := &memorydb.DescribeParametersOutput{}

        mockClient.On("DescribeParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedNodes", func(t *testing.T) {
        input := &memorydb.DescribeReservedNodesInput{}
        output := &memorydb.DescribeReservedNodesOutput{}

        mockClient.On("DescribeReservedNodes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedNodesOfferings", func(t *testing.T) {
        input := &memorydb.DescribeReservedNodesOfferingsInput{}
        output := &memorydb.DescribeReservedNodesOfferingsOutput{}

        mockClient.On("DescribeReservedNodesOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedNodesOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceUpdates", func(t *testing.T) {
        input := &memorydb.DescribeServiceUpdatesInput{}
        output := &memorydb.DescribeServiceUpdatesOutput{}

        mockClient.On("DescribeServiceUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshots", func(t *testing.T) {
        input := &memorydb.DescribeSnapshotsInput{}
        output := &memorydb.DescribeSnapshotsOutput{}

        mockClient.On("DescribeSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubnetGroups", func(t *testing.T) {
        input := &memorydb.DescribeSubnetGroupsInput{}
        output := &memorydb.DescribeSubnetGroupsOutput{}

        mockClient.On("DescribeSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsers", func(t *testing.T) {
        input := &memorydb.DescribeUsersInput{}
        output := &memorydb.DescribeUsersOutput{}

        mockClient.On("DescribeUsers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverShard", func(t *testing.T) {
        input := &memorydb.FailoverShardInput{}
        output := &memorydb.FailoverShardOutput{}

        mockClient.On("FailoverShard", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverShard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAllowedNodeTypeUpdates", func(t *testing.T) {
        input := &memorydb.ListAllowedNodeTypeUpdatesInput{}
        output := &memorydb.ListAllowedNodeTypeUpdatesOutput{}

        mockClient.On("ListAllowedNodeTypeUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.ListAllowedNodeTypeUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &memorydb.ListTagsInput{}
        output := &memorydb.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedNodesOffering", func(t *testing.T) {
        input := &memorydb.PurchaseReservedNodesOfferingInput{}
        output := &memorydb.PurchaseReservedNodesOfferingOutput{}

        mockClient.On("PurchaseReservedNodesOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedNodesOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetParameterGroup", func(t *testing.T) {
        input := &memorydb.ResetParameterGroupInput{}
        output := &memorydb.ResetParameterGroupOutput{}

        mockClient.On("ResetParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &memorydb.TagResourceInput{}
        output := &memorydb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &memorydb.UntagResourceInput{}
        output := &memorydb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateACL", func(t *testing.T) {
        input := &memorydb.UpdateACLInput{}
        output := &memorydb.UpdateACLOutput{}

        mockClient.On("UpdateACL", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &memorydb.UpdateClusterInput{}
        output := &memorydb.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateParameterGroup", func(t *testing.T) {
        input := &memorydb.UpdateParameterGroupInput{}
        output := &memorydb.UpdateParameterGroupOutput{}

        mockClient.On("UpdateParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubnetGroup", func(t *testing.T) {
        input := &memorydb.UpdateSubnetGroupInput{}
        output := &memorydb.UpdateSubnetGroupOutput{}

        mockClient.On("UpdateSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &memorydb.UpdateUserInput{}
        output := &memorydb.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
