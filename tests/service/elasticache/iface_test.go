package elasticache_test

// tests for the elasticache service interface for this ../../../service/elasticache/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticache/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticache/elasticache_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticacheServiceCanBeMocked(t *testing.T) {
	var iface elasticache_iface.IClient
	iface = &elasticache.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticache.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &elasticache.AddTagsToResourceInput{}
        output := &elasticache.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeCacheSecurityGroupIngress", func(t *testing.T) {
        input := &elasticache.AuthorizeCacheSecurityGroupIngressInput{}
        output := &elasticache.AuthorizeCacheSecurityGroupIngressOutput{}

        mockClient.On("AuthorizeCacheSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeCacheSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchApplyUpdateAction", func(t *testing.T) {
        input := &elasticache.BatchApplyUpdateActionInput{}
        output := &elasticache.BatchApplyUpdateActionOutput{}

        mockClient.On("BatchApplyUpdateAction", ctx, input).Return(output, nil)

        result, err := mockClient.BatchApplyUpdateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStopUpdateAction", func(t *testing.T) {
        input := &elasticache.BatchStopUpdateActionInput{}
        output := &elasticache.BatchStopUpdateActionOutput{}

        mockClient.On("BatchStopUpdateAction", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStopUpdateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteMigration", func(t *testing.T) {
        input := &elasticache.CompleteMigrationInput{}
        output := &elasticache.CompleteMigrationOutput{}

        mockClient.On("CompleteMigration", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteMigration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyServerlessCacheSnapshot", func(t *testing.T) {
        input := &elasticache.CopyServerlessCacheSnapshotInput{}
        output := &elasticache.CopyServerlessCacheSnapshotOutput{}

        mockClient.On("CopyServerlessCacheSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyServerlessCacheSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopySnapshot", func(t *testing.T) {
        input := &elasticache.CopySnapshotInput{}
        output := &elasticache.CopySnapshotOutput{}

        mockClient.On("CopySnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopySnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCacheCluster", func(t *testing.T) {
        input := &elasticache.CreateCacheClusterInput{}
        output := &elasticache.CreateCacheClusterOutput{}

        mockClient.On("CreateCacheCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCacheCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCacheParameterGroup", func(t *testing.T) {
        input := &elasticache.CreateCacheParameterGroupInput{}
        output := &elasticache.CreateCacheParameterGroupOutput{}

        mockClient.On("CreateCacheParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCacheParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCacheSecurityGroup", func(t *testing.T) {
        input := &elasticache.CreateCacheSecurityGroupInput{}
        output := &elasticache.CreateCacheSecurityGroupOutput{}

        mockClient.On("CreateCacheSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCacheSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCacheSubnetGroup", func(t *testing.T) {
        input := &elasticache.CreateCacheSubnetGroupInput{}
        output := &elasticache.CreateCacheSubnetGroupOutput{}

        mockClient.On("CreateCacheSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCacheSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.CreateGlobalReplicationGroupInput{}
        output := &elasticache.CreateGlobalReplicationGroupOutput{}

        mockClient.On("CreateGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationGroup", func(t *testing.T) {
        input := &elasticache.CreateReplicationGroupInput{}
        output := &elasticache.CreateReplicationGroupOutput{}

        mockClient.On("CreateReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServerlessCache", func(t *testing.T) {
        input := &elasticache.CreateServerlessCacheInput{}
        output := &elasticache.CreateServerlessCacheOutput{}

        mockClient.On("CreateServerlessCache", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServerlessCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServerlessCacheSnapshot", func(t *testing.T) {
        input := &elasticache.CreateServerlessCacheSnapshotInput{}
        output := &elasticache.CreateServerlessCacheSnapshotOutput{}

        mockClient.On("CreateServerlessCacheSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServerlessCacheSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &elasticache.CreateSnapshotInput{}
        output := &elasticache.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &elasticache.CreateUserInput{}
        output := &elasticache.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserGroup", func(t *testing.T) {
        input := &elasticache.CreateUserGroupInput{}
        output := &elasticache.CreateUserGroupOutput{}

        mockClient.On("CreateUserGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecreaseNodeGroupsInGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput{}
        output := &elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput{}

        mockClient.On("DecreaseNodeGroupsInGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DecreaseNodeGroupsInGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecreaseReplicaCount", func(t *testing.T) {
        input := &elasticache.DecreaseReplicaCountInput{}
        output := &elasticache.DecreaseReplicaCountOutput{}

        mockClient.On("DecreaseReplicaCount", ctx, input).Return(output, nil)

        result, err := mockClient.DecreaseReplicaCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCacheCluster", func(t *testing.T) {
        input := &elasticache.DeleteCacheClusterInput{}
        output := &elasticache.DeleteCacheClusterOutput{}

        mockClient.On("DeleteCacheCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCacheCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCacheParameterGroup", func(t *testing.T) {
        input := &elasticache.DeleteCacheParameterGroupInput{}
        output := &elasticache.DeleteCacheParameterGroupOutput{}

        mockClient.On("DeleteCacheParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCacheParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCacheSecurityGroup", func(t *testing.T) {
        input := &elasticache.DeleteCacheSecurityGroupInput{}
        output := &elasticache.DeleteCacheSecurityGroupOutput{}

        mockClient.On("DeleteCacheSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCacheSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCacheSubnetGroup", func(t *testing.T) {
        input := &elasticache.DeleteCacheSubnetGroupInput{}
        output := &elasticache.DeleteCacheSubnetGroupOutput{}

        mockClient.On("DeleteCacheSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCacheSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.DeleteGlobalReplicationGroupInput{}
        output := &elasticache.DeleteGlobalReplicationGroupOutput{}

        mockClient.On("DeleteGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationGroup", func(t *testing.T) {
        input := &elasticache.DeleteReplicationGroupInput{}
        output := &elasticache.DeleteReplicationGroupOutput{}

        mockClient.On("DeleteReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServerlessCache", func(t *testing.T) {
        input := &elasticache.DeleteServerlessCacheInput{}
        output := &elasticache.DeleteServerlessCacheOutput{}

        mockClient.On("DeleteServerlessCache", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServerlessCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServerlessCacheSnapshot", func(t *testing.T) {
        input := &elasticache.DeleteServerlessCacheSnapshotInput{}
        output := &elasticache.DeleteServerlessCacheSnapshotOutput{}

        mockClient.On("DeleteServerlessCacheSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServerlessCacheSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &elasticache.DeleteSnapshotInput{}
        output := &elasticache.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &elasticache.DeleteUserInput{}
        output := &elasticache.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserGroup", func(t *testing.T) {
        input := &elasticache.DeleteUserGroupInput{}
        output := &elasticache.DeleteUserGroupOutput{}

        mockClient.On("DeleteUserGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheClusters", func(t *testing.T) {
        input := &elasticache.DescribeCacheClustersInput{}
        output := &elasticache.DescribeCacheClustersOutput{}

        mockClient.On("DescribeCacheClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheEngineVersions", func(t *testing.T) {
        input := &elasticache.DescribeCacheEngineVersionsInput{}
        output := &elasticache.DescribeCacheEngineVersionsOutput{}

        mockClient.On("DescribeCacheEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheParameterGroups", func(t *testing.T) {
        input := &elasticache.DescribeCacheParameterGroupsInput{}
        output := &elasticache.DescribeCacheParameterGroupsOutput{}

        mockClient.On("DescribeCacheParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheParameters", func(t *testing.T) {
        input := &elasticache.DescribeCacheParametersInput{}
        output := &elasticache.DescribeCacheParametersOutput{}

        mockClient.On("DescribeCacheParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheSecurityGroups", func(t *testing.T) {
        input := &elasticache.DescribeCacheSecurityGroupsInput{}
        output := &elasticache.DescribeCacheSecurityGroupsOutput{}

        mockClient.On("DescribeCacheSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCacheSubnetGroups", func(t *testing.T) {
        input := &elasticache.DescribeCacheSubnetGroupsInput{}
        output := &elasticache.DescribeCacheSubnetGroupsOutput{}

        mockClient.On("DescribeCacheSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCacheSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultParameters", func(t *testing.T) {
        input := &elasticache.DescribeEngineDefaultParametersInput{}
        output := &elasticache.DescribeEngineDefaultParametersOutput{}

        mockClient.On("DescribeEngineDefaultParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &elasticache.DescribeEventsInput{}
        output := &elasticache.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalReplicationGroups", func(t *testing.T) {
        input := &elasticache.DescribeGlobalReplicationGroupsInput{}
        output := &elasticache.DescribeGlobalReplicationGroupsOutput{}

        mockClient.On("DescribeGlobalReplicationGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalReplicationGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationGroups", func(t *testing.T) {
        input := &elasticache.DescribeReplicationGroupsInput{}
        output := &elasticache.DescribeReplicationGroupsOutput{}

        mockClient.On("DescribeReplicationGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedCacheNodes", func(t *testing.T) {
        input := &elasticache.DescribeReservedCacheNodesInput{}
        output := &elasticache.DescribeReservedCacheNodesOutput{}

        mockClient.On("DescribeReservedCacheNodes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedCacheNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedCacheNodesOfferings", func(t *testing.T) {
        input := &elasticache.DescribeReservedCacheNodesOfferingsInput{}
        output := &elasticache.DescribeReservedCacheNodesOfferingsOutput{}

        mockClient.On("DescribeReservedCacheNodesOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedCacheNodesOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServerlessCacheSnapshots", func(t *testing.T) {
        input := &elasticache.DescribeServerlessCacheSnapshotsInput{}
        output := &elasticache.DescribeServerlessCacheSnapshotsOutput{}

        mockClient.On("DescribeServerlessCacheSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServerlessCacheSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServerlessCaches", func(t *testing.T) {
        input := &elasticache.DescribeServerlessCachesInput{}
        output := &elasticache.DescribeServerlessCachesOutput{}

        mockClient.On("DescribeServerlessCaches", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServerlessCaches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceUpdates", func(t *testing.T) {
        input := &elasticache.DescribeServiceUpdatesInput{}
        output := &elasticache.DescribeServiceUpdatesOutput{}

        mockClient.On("DescribeServiceUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshots", func(t *testing.T) {
        input := &elasticache.DescribeSnapshotsInput{}
        output := &elasticache.DescribeSnapshotsOutput{}

        mockClient.On("DescribeSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUpdateActions", func(t *testing.T) {
        input := &elasticache.DescribeUpdateActionsInput{}
        output := &elasticache.DescribeUpdateActionsOutput{}

        mockClient.On("DescribeUpdateActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUpdateActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserGroups", func(t *testing.T) {
        input := &elasticache.DescribeUserGroupsInput{}
        output := &elasticache.DescribeUserGroupsOutput{}

        mockClient.On("DescribeUserGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsers", func(t *testing.T) {
        input := &elasticache.DescribeUsersInput{}
        output := &elasticache.DescribeUsersOutput{}

        mockClient.On("DescribeUsers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.DisassociateGlobalReplicationGroupInput{}
        output := &elasticache.DisassociateGlobalReplicationGroupOutput{}

        mockClient.On("DisassociateGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportServerlessCacheSnapshot", func(t *testing.T) {
        input := &elasticache.ExportServerlessCacheSnapshotInput{}
        output := &elasticache.ExportServerlessCacheSnapshotOutput{}

        mockClient.On("ExportServerlessCacheSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ExportServerlessCacheSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.FailoverGlobalReplicationGroupInput{}
        output := &elasticache.FailoverGlobalReplicationGroupOutput{}

        mockClient.On("FailoverGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIncreaseNodeGroupsInGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput{}
        output := &elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput{}

        mockClient.On("IncreaseNodeGroupsInGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.IncreaseNodeGroupsInGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIncreaseReplicaCount", func(t *testing.T) {
        input := &elasticache.IncreaseReplicaCountInput{}
        output := &elasticache.IncreaseReplicaCountOutput{}

        mockClient.On("IncreaseReplicaCount", ctx, input).Return(output, nil)

        result, err := mockClient.IncreaseReplicaCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAllowedNodeTypeModifications", func(t *testing.T) {
        input := &elasticache.ListAllowedNodeTypeModificationsInput{}
        output := &elasticache.ListAllowedNodeTypeModificationsOutput{}

        mockClient.On("ListAllowedNodeTypeModifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListAllowedNodeTypeModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &elasticache.ListTagsForResourceInput{}
        output := &elasticache.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCacheCluster", func(t *testing.T) {
        input := &elasticache.ModifyCacheClusterInput{}
        output := &elasticache.ModifyCacheClusterOutput{}

        mockClient.On("ModifyCacheCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCacheCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCacheParameterGroup", func(t *testing.T) {
        input := &elasticache.ModifyCacheParameterGroupInput{}
        output := &elasticache.ModifyCacheParameterGroupOutput{}

        mockClient.On("ModifyCacheParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCacheParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCacheSubnetGroup", func(t *testing.T) {
        input := &elasticache.ModifyCacheSubnetGroupInput{}
        output := &elasticache.ModifyCacheSubnetGroupOutput{}

        mockClient.On("ModifyCacheSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCacheSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.ModifyGlobalReplicationGroupInput{}
        output := &elasticache.ModifyGlobalReplicationGroupOutput{}

        mockClient.On("ModifyGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationGroup", func(t *testing.T) {
        input := &elasticache.ModifyReplicationGroupInput{}
        output := &elasticache.ModifyReplicationGroupOutput{}

        mockClient.On("ModifyReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationGroupShardConfiguration", func(t *testing.T) {
        input := &elasticache.ModifyReplicationGroupShardConfigurationInput{}
        output := &elasticache.ModifyReplicationGroupShardConfigurationOutput{}

        mockClient.On("ModifyReplicationGroupShardConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationGroupShardConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyServerlessCache", func(t *testing.T) {
        input := &elasticache.ModifyServerlessCacheInput{}
        output := &elasticache.ModifyServerlessCacheOutput{}

        mockClient.On("ModifyServerlessCache", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyServerlessCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyUser", func(t *testing.T) {
        input := &elasticache.ModifyUserInput{}
        output := &elasticache.ModifyUserOutput{}

        mockClient.On("ModifyUser", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyUserGroup", func(t *testing.T) {
        input := &elasticache.ModifyUserGroupInput{}
        output := &elasticache.ModifyUserGroupOutput{}

        mockClient.On("ModifyUserGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyUserGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedCacheNodesOffering", func(t *testing.T) {
        input := &elasticache.PurchaseReservedCacheNodesOfferingInput{}
        output := &elasticache.PurchaseReservedCacheNodesOfferingOutput{}

        mockClient.On("PurchaseReservedCacheNodesOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedCacheNodesOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebalanceSlotsInGlobalReplicationGroup", func(t *testing.T) {
        input := &elasticache.RebalanceSlotsInGlobalReplicationGroupInput{}
        output := &elasticache.RebalanceSlotsInGlobalReplicationGroupOutput{}

        mockClient.On("RebalanceSlotsInGlobalReplicationGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RebalanceSlotsInGlobalReplicationGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootCacheCluster", func(t *testing.T) {
        input := &elasticache.RebootCacheClusterInput{}
        output := &elasticache.RebootCacheClusterOutput{}

        mockClient.On("RebootCacheCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RebootCacheCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &elasticache.RemoveTagsFromResourceInput{}
        output := &elasticache.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetCacheParameterGroup", func(t *testing.T) {
        input := &elasticache.ResetCacheParameterGroupInput{}
        output := &elasticache.ResetCacheParameterGroupOutput{}

        mockClient.On("ResetCacheParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetCacheParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeCacheSecurityGroupIngress", func(t *testing.T) {
        input := &elasticache.RevokeCacheSecurityGroupIngressInput{}
        output := &elasticache.RevokeCacheSecurityGroupIngressOutput{}

        mockClient.On("RevokeCacheSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeCacheSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMigration", func(t *testing.T) {
        input := &elasticache.StartMigrationInput{}
        output := &elasticache.StartMigrationOutput{}

        mockClient.On("StartMigration", ctx, input).Return(output, nil)

        result, err := mockClient.StartMigration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestFailover", func(t *testing.T) {
        input := &elasticache.TestFailoverInput{}
        output := &elasticache.TestFailoverOutput{}

        mockClient.On("TestFailover", ctx, input).Return(output, nil)

        result, err := mockClient.TestFailover(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestMigration", func(t *testing.T) {
        input := &elasticache.TestMigrationInput{}
        output := &elasticache.TestMigrationOutput{}

        mockClient.On("TestMigration", ctx, input).Return(output, nil)

        result, err := mockClient.TestMigration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
