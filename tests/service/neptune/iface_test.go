package neptune_test

// tests for the neptune service interface for this ../../../service/neptune/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptune/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptune/neptune_iface"
	"github.com/stretchr/testify/assert"
)

func TestNeptuneServiceCanBeMocked(t *testing.T) {
	var iface neptune_iface.IClient
	iface = &neptune.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := neptune.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddRoleToDBCluster", func(t *testing.T) {
        input := &neptune.AddRoleToDBClusterInput{}
        output := &neptune.AddRoleToDBClusterOutput{}

        mockClient.On("AddRoleToDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.AddRoleToDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddSourceIdentifierToSubscription", func(t *testing.T) {
        input := &neptune.AddSourceIdentifierToSubscriptionInput{}
        output := &neptune.AddSourceIdentifierToSubscriptionOutput{}

        mockClient.On("AddSourceIdentifierToSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.AddSourceIdentifierToSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &neptune.AddTagsToResourceInput{}
        output := &neptune.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyPendingMaintenanceAction", func(t *testing.T) {
        input := &neptune.ApplyPendingMaintenanceActionInput{}
        output := &neptune.ApplyPendingMaintenanceActionOutput{}

        mockClient.On("ApplyPendingMaintenanceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyPendingMaintenanceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterParameterGroup", func(t *testing.T) {
        input := &neptune.CopyDBClusterParameterGroupInput{}
        output := &neptune.CopyDBClusterParameterGroupOutput{}

        mockClient.On("CopyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterSnapshot", func(t *testing.T) {
        input := &neptune.CopyDBClusterSnapshotInput{}
        output := &neptune.CopyDBClusterSnapshotOutput{}

        mockClient.On("CopyDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBParameterGroup", func(t *testing.T) {
        input := &neptune.CopyDBParameterGroupInput{}
        output := &neptune.CopyDBParameterGroupOutput{}

        mockClient.On("CopyDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBCluster", func(t *testing.T) {
        input := &neptune.CreateDBClusterInput{}
        output := &neptune.CreateDBClusterOutput{}

        mockClient.On("CreateDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterEndpoint", func(t *testing.T) {
        input := &neptune.CreateDBClusterEndpointInput{}
        output := &neptune.CreateDBClusterEndpointOutput{}

        mockClient.On("CreateDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterParameterGroup", func(t *testing.T) {
        input := &neptune.CreateDBClusterParameterGroupInput{}
        output := &neptune.CreateDBClusterParameterGroupOutput{}

        mockClient.On("CreateDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterSnapshot", func(t *testing.T) {
        input := &neptune.CreateDBClusterSnapshotInput{}
        output := &neptune.CreateDBClusterSnapshotOutput{}

        mockClient.On("CreateDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBInstance", func(t *testing.T) {
        input := &neptune.CreateDBInstanceInput{}
        output := &neptune.CreateDBInstanceOutput{}

        mockClient.On("CreateDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBParameterGroup", func(t *testing.T) {
        input := &neptune.CreateDBParameterGroupInput{}
        output := &neptune.CreateDBParameterGroupOutput{}

        mockClient.On("CreateDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBSubnetGroup", func(t *testing.T) {
        input := &neptune.CreateDBSubnetGroupInput{}
        output := &neptune.CreateDBSubnetGroupOutput{}

        mockClient.On("CreateDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSubscription", func(t *testing.T) {
        input := &neptune.CreateEventSubscriptionInput{}
        output := &neptune.CreateEventSubscriptionOutput{}

        mockClient.On("CreateEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalCluster", func(t *testing.T) {
        input := &neptune.CreateGlobalClusterInput{}
        output := &neptune.CreateGlobalClusterOutput{}

        mockClient.On("CreateGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBCluster", func(t *testing.T) {
        input := &neptune.DeleteDBClusterInput{}
        output := &neptune.DeleteDBClusterOutput{}

        mockClient.On("DeleteDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterEndpoint", func(t *testing.T) {
        input := &neptune.DeleteDBClusterEndpointInput{}
        output := &neptune.DeleteDBClusterEndpointOutput{}

        mockClient.On("DeleteDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterParameterGroup", func(t *testing.T) {
        input := &neptune.DeleteDBClusterParameterGroupInput{}
        output := &neptune.DeleteDBClusterParameterGroupOutput{}

        mockClient.On("DeleteDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterSnapshot", func(t *testing.T) {
        input := &neptune.DeleteDBClusterSnapshotInput{}
        output := &neptune.DeleteDBClusterSnapshotOutput{}

        mockClient.On("DeleteDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBInstance", func(t *testing.T) {
        input := &neptune.DeleteDBInstanceInput{}
        output := &neptune.DeleteDBInstanceOutput{}

        mockClient.On("DeleteDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBParameterGroup", func(t *testing.T) {
        input := &neptune.DeleteDBParameterGroupInput{}
        output := &neptune.DeleteDBParameterGroupOutput{}

        mockClient.On("DeleteDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBSubnetGroup", func(t *testing.T) {
        input := &neptune.DeleteDBSubnetGroupInput{}
        output := &neptune.DeleteDBSubnetGroupOutput{}

        mockClient.On("DeleteDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSubscription", func(t *testing.T) {
        input := &neptune.DeleteEventSubscriptionInput{}
        output := &neptune.DeleteEventSubscriptionOutput{}

        mockClient.On("DeleteEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlobalCluster", func(t *testing.T) {
        input := &neptune.DeleteGlobalClusterInput{}
        output := &neptune.DeleteGlobalClusterOutput{}

        mockClient.On("DeleteGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterEndpoints", func(t *testing.T) {
        input := &neptune.DescribeDBClusterEndpointsInput{}
        output := &neptune.DescribeDBClusterEndpointsOutput{}

        mockClient.On("DescribeDBClusterEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameterGroups", func(t *testing.T) {
        input := &neptune.DescribeDBClusterParameterGroupsInput{}
        output := &neptune.DescribeDBClusterParameterGroupsOutput{}

        mockClient.On("DescribeDBClusterParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameters", func(t *testing.T) {
        input := &neptune.DescribeDBClusterParametersInput{}
        output := &neptune.DescribeDBClusterParametersOutput{}

        mockClient.On("DescribeDBClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshotAttributes", func(t *testing.T) {
        input := &neptune.DescribeDBClusterSnapshotAttributesInput{}
        output := &neptune.DescribeDBClusterSnapshotAttributesOutput{}

        mockClient.On("DescribeDBClusterSnapshotAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshotAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshots", func(t *testing.T) {
        input := &neptune.DescribeDBClusterSnapshotsInput{}
        output := &neptune.DescribeDBClusterSnapshotsOutput{}

        mockClient.On("DescribeDBClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusters", func(t *testing.T) {
        input := &neptune.DescribeDBClustersInput{}
        output := &neptune.DescribeDBClustersOutput{}

        mockClient.On("DescribeDBClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBEngineVersions", func(t *testing.T) {
        input := &neptune.DescribeDBEngineVersionsInput{}
        output := &neptune.DescribeDBEngineVersionsOutput{}

        mockClient.On("DescribeDBEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBInstances", func(t *testing.T) {
        input := &neptune.DescribeDBInstancesInput{}
        output := &neptune.DescribeDBInstancesOutput{}

        mockClient.On("DescribeDBInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBParameterGroups", func(t *testing.T) {
        input := &neptune.DescribeDBParameterGroupsInput{}
        output := &neptune.DescribeDBParameterGroupsOutput{}

        mockClient.On("DescribeDBParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBParameters", func(t *testing.T) {
        input := &neptune.DescribeDBParametersInput{}
        output := &neptune.DescribeDBParametersOutput{}

        mockClient.On("DescribeDBParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSubnetGroups", func(t *testing.T) {
        input := &neptune.DescribeDBSubnetGroupsInput{}
        output := &neptune.DescribeDBSubnetGroupsOutput{}

        mockClient.On("DescribeDBSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultClusterParameters", func(t *testing.T) {
        input := &neptune.DescribeEngineDefaultClusterParametersInput{}
        output := &neptune.DescribeEngineDefaultClusterParametersOutput{}

        mockClient.On("DescribeEngineDefaultClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultParameters", func(t *testing.T) {
        input := &neptune.DescribeEngineDefaultParametersInput{}
        output := &neptune.DescribeEngineDefaultParametersOutput{}

        mockClient.On("DescribeEngineDefaultParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventCategories", func(t *testing.T) {
        input := &neptune.DescribeEventCategoriesInput{}
        output := &neptune.DescribeEventCategoriesOutput{}

        mockClient.On("DescribeEventCategories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSubscriptions", func(t *testing.T) {
        input := &neptune.DescribeEventSubscriptionsInput{}
        output := &neptune.DescribeEventSubscriptionsOutput{}

        mockClient.On("DescribeEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &neptune.DescribeEventsInput{}
        output := &neptune.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalClusters", func(t *testing.T) {
        input := &neptune.DescribeGlobalClustersInput{}
        output := &neptune.DescribeGlobalClustersOutput{}

        mockClient.On("DescribeGlobalClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrderableDBInstanceOptions", func(t *testing.T) {
        input := &neptune.DescribeOrderableDBInstanceOptionsInput{}
        output := &neptune.DescribeOrderableDBInstanceOptionsOutput{}

        mockClient.On("DescribeOrderableDBInstanceOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrderableDBInstanceOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePendingMaintenanceActions", func(t *testing.T) {
        input := &neptune.DescribePendingMaintenanceActionsInput{}
        output := &neptune.DescribePendingMaintenanceActionsOutput{}

        mockClient.On("DescribePendingMaintenanceActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePendingMaintenanceActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeValidDBInstanceModifications", func(t *testing.T) {
        input := &neptune.DescribeValidDBInstanceModificationsInput{}
        output := &neptune.DescribeValidDBInstanceModificationsOutput{}

        mockClient.On("DescribeValidDBInstanceModifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeValidDBInstanceModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverDBCluster", func(t *testing.T) {
        input := &neptune.FailoverDBClusterInput{}
        output := &neptune.FailoverDBClusterOutput{}

        mockClient.On("FailoverDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverGlobalCluster", func(t *testing.T) {
        input := &neptune.FailoverGlobalClusterInput{}
        output := &neptune.FailoverGlobalClusterOutput{}

        mockClient.On("FailoverGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &neptune.ListTagsForResourceInput{}
        output := &neptune.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBCluster", func(t *testing.T) {
        input := &neptune.ModifyDBClusterInput{}
        output := &neptune.ModifyDBClusterOutput{}

        mockClient.On("ModifyDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterEndpoint", func(t *testing.T) {
        input := &neptune.ModifyDBClusterEndpointInput{}
        output := &neptune.ModifyDBClusterEndpointOutput{}

        mockClient.On("ModifyDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterParameterGroup", func(t *testing.T) {
        input := &neptune.ModifyDBClusterParameterGroupInput{}
        output := &neptune.ModifyDBClusterParameterGroupOutput{}

        mockClient.On("ModifyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterSnapshotAttribute", func(t *testing.T) {
        input := &neptune.ModifyDBClusterSnapshotAttributeInput{}
        output := &neptune.ModifyDBClusterSnapshotAttributeOutput{}

        mockClient.On("ModifyDBClusterSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBInstance", func(t *testing.T) {
        input := &neptune.ModifyDBInstanceInput{}
        output := &neptune.ModifyDBInstanceOutput{}

        mockClient.On("ModifyDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBParameterGroup", func(t *testing.T) {
        input := &neptune.ModifyDBParameterGroupInput{}
        output := &neptune.ModifyDBParameterGroupOutput{}

        mockClient.On("ModifyDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBSubnetGroup", func(t *testing.T) {
        input := &neptune.ModifyDBSubnetGroupInput{}
        output := &neptune.ModifyDBSubnetGroupOutput{}

        mockClient.On("ModifyDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEventSubscription", func(t *testing.T) {
        input := &neptune.ModifyEventSubscriptionInput{}
        output := &neptune.ModifyEventSubscriptionOutput{}

        mockClient.On("ModifyEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyGlobalCluster", func(t *testing.T) {
        input := &neptune.ModifyGlobalClusterInput{}
        output := &neptune.ModifyGlobalClusterOutput{}

        mockClient.On("ModifyGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromoteReadReplicaDBCluster", func(t *testing.T) {
        input := &neptune.PromoteReadReplicaDBClusterInput{}
        output := &neptune.PromoteReadReplicaDBClusterOutput{}

        mockClient.On("PromoteReadReplicaDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.PromoteReadReplicaDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootDBInstance", func(t *testing.T) {
        input := &neptune.RebootDBInstanceInput{}
        output := &neptune.RebootDBInstanceOutput{}

        mockClient.On("RebootDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFromGlobalCluster", func(t *testing.T) {
        input := &neptune.RemoveFromGlobalClusterInput{}
        output := &neptune.RemoveFromGlobalClusterOutput{}

        mockClient.On("RemoveFromGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFromGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRoleFromDBCluster", func(t *testing.T) {
        input := &neptune.RemoveRoleFromDBClusterInput{}
        output := &neptune.RemoveRoleFromDBClusterOutput{}

        mockClient.On("RemoveRoleFromDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRoleFromDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveSourceIdentifierFromSubscription", func(t *testing.T) {
        input := &neptune.RemoveSourceIdentifierFromSubscriptionInput{}
        output := &neptune.RemoveSourceIdentifierFromSubscriptionOutput{}

        mockClient.On("RemoveSourceIdentifierFromSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveSourceIdentifierFromSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &neptune.RemoveTagsFromResourceInput{}
        output := &neptune.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDBClusterParameterGroup", func(t *testing.T) {
        input := &neptune.ResetDBClusterParameterGroupInput{}
        output := &neptune.ResetDBClusterParameterGroupOutput{}

        mockClient.On("ResetDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDBParameterGroup", func(t *testing.T) {
        input := &neptune.ResetDBParameterGroupInput{}
        output := &neptune.ResetDBParameterGroupOutput{}

        mockClient.On("ResetDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterFromSnapshot", func(t *testing.T) {
        input := &neptune.RestoreDBClusterFromSnapshotInput{}
        output := &neptune.RestoreDBClusterFromSnapshotOutput{}

        mockClient.On("RestoreDBClusterFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterToPointInTime", func(t *testing.T) {
        input := &neptune.RestoreDBClusterToPointInTimeInput{}
        output := &neptune.RestoreDBClusterToPointInTimeOutput{}

        mockClient.On("RestoreDBClusterToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDBCluster", func(t *testing.T) {
        input := &neptune.StartDBClusterInput{}
        output := &neptune.StartDBClusterOutput{}

        mockClient.On("StartDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StartDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDBCluster", func(t *testing.T) {
        input := &neptune.StopDBClusterInput{}
        output := &neptune.StopDBClusterOutput{}

        mockClient.On("StopDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StopDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
