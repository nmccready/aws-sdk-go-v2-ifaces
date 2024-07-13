package docdb_test

// tests for the docdb service interface for this ../../../service/docdb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/docdb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/docdb/docdb_iface"
	"github.com/stretchr/testify/assert"
)

func TestDocdbServiceCanBeMocked(t *testing.T) {
	var iface docdb_iface.IClient
	iface = &docdb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := docdb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddSourceIdentifierToSubscription", func(t *testing.T) {
        input := &docdb.AddSourceIdentifierToSubscriptionInput{}
        output := &docdb.AddSourceIdentifierToSubscriptionOutput{}

        mockClient.On("AddSourceIdentifierToSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.AddSourceIdentifierToSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &docdb.AddTagsToResourceInput{}
        output := &docdb.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyPendingMaintenanceAction", func(t *testing.T) {
        input := &docdb.ApplyPendingMaintenanceActionInput{}
        output := &docdb.ApplyPendingMaintenanceActionOutput{}

        mockClient.On("ApplyPendingMaintenanceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyPendingMaintenanceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterParameterGroup", func(t *testing.T) {
        input := &docdb.CopyDBClusterParameterGroupInput{}
        output := &docdb.CopyDBClusterParameterGroupOutput{}

        mockClient.On("CopyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterSnapshot", func(t *testing.T) {
        input := &docdb.CopyDBClusterSnapshotInput{}
        output := &docdb.CopyDBClusterSnapshotOutput{}

        mockClient.On("CopyDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBCluster", func(t *testing.T) {
        input := &docdb.CreateDBClusterInput{}
        output := &docdb.CreateDBClusterOutput{}

        mockClient.On("CreateDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterParameterGroup", func(t *testing.T) {
        input := &docdb.CreateDBClusterParameterGroupInput{}
        output := &docdb.CreateDBClusterParameterGroupOutput{}

        mockClient.On("CreateDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterSnapshot", func(t *testing.T) {
        input := &docdb.CreateDBClusterSnapshotInput{}
        output := &docdb.CreateDBClusterSnapshotOutput{}

        mockClient.On("CreateDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBInstance", func(t *testing.T) {
        input := &docdb.CreateDBInstanceInput{}
        output := &docdb.CreateDBInstanceOutput{}

        mockClient.On("CreateDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBSubnetGroup", func(t *testing.T) {
        input := &docdb.CreateDBSubnetGroupInput{}
        output := &docdb.CreateDBSubnetGroupOutput{}

        mockClient.On("CreateDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSubscription", func(t *testing.T) {
        input := &docdb.CreateEventSubscriptionInput{}
        output := &docdb.CreateEventSubscriptionOutput{}

        mockClient.On("CreateEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalCluster", func(t *testing.T) {
        input := &docdb.CreateGlobalClusterInput{}
        output := &docdb.CreateGlobalClusterOutput{}

        mockClient.On("CreateGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBCluster", func(t *testing.T) {
        input := &docdb.DeleteDBClusterInput{}
        output := &docdb.DeleteDBClusterOutput{}

        mockClient.On("DeleteDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterParameterGroup", func(t *testing.T) {
        input := &docdb.DeleteDBClusterParameterGroupInput{}
        output := &docdb.DeleteDBClusterParameterGroupOutput{}

        mockClient.On("DeleteDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterSnapshot", func(t *testing.T) {
        input := &docdb.DeleteDBClusterSnapshotInput{}
        output := &docdb.DeleteDBClusterSnapshotOutput{}

        mockClient.On("DeleteDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBInstance", func(t *testing.T) {
        input := &docdb.DeleteDBInstanceInput{}
        output := &docdb.DeleteDBInstanceOutput{}

        mockClient.On("DeleteDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBSubnetGroup", func(t *testing.T) {
        input := &docdb.DeleteDBSubnetGroupInput{}
        output := &docdb.DeleteDBSubnetGroupOutput{}

        mockClient.On("DeleteDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSubscription", func(t *testing.T) {
        input := &docdb.DeleteEventSubscriptionInput{}
        output := &docdb.DeleteEventSubscriptionOutput{}

        mockClient.On("DeleteEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlobalCluster", func(t *testing.T) {
        input := &docdb.DeleteGlobalClusterInput{}
        output := &docdb.DeleteGlobalClusterOutput{}

        mockClient.On("DeleteGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificates", func(t *testing.T) {
        input := &docdb.DescribeCertificatesInput{}
        output := &docdb.DescribeCertificatesOutput{}

        mockClient.On("DescribeCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameterGroups", func(t *testing.T) {
        input := &docdb.DescribeDBClusterParameterGroupsInput{}
        output := &docdb.DescribeDBClusterParameterGroupsOutput{}

        mockClient.On("DescribeDBClusterParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameters", func(t *testing.T) {
        input := &docdb.DescribeDBClusterParametersInput{}
        output := &docdb.DescribeDBClusterParametersOutput{}

        mockClient.On("DescribeDBClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshotAttributes", func(t *testing.T) {
        input := &docdb.DescribeDBClusterSnapshotAttributesInput{}
        output := &docdb.DescribeDBClusterSnapshotAttributesOutput{}

        mockClient.On("DescribeDBClusterSnapshotAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshotAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshots", func(t *testing.T) {
        input := &docdb.DescribeDBClusterSnapshotsInput{}
        output := &docdb.DescribeDBClusterSnapshotsOutput{}

        mockClient.On("DescribeDBClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusters", func(t *testing.T) {
        input := &docdb.DescribeDBClustersInput{}
        output := &docdb.DescribeDBClustersOutput{}

        mockClient.On("DescribeDBClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBEngineVersions", func(t *testing.T) {
        input := &docdb.DescribeDBEngineVersionsInput{}
        output := &docdb.DescribeDBEngineVersionsOutput{}

        mockClient.On("DescribeDBEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBInstances", func(t *testing.T) {
        input := &docdb.DescribeDBInstancesInput{}
        output := &docdb.DescribeDBInstancesOutput{}

        mockClient.On("DescribeDBInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSubnetGroups", func(t *testing.T) {
        input := &docdb.DescribeDBSubnetGroupsInput{}
        output := &docdb.DescribeDBSubnetGroupsOutput{}

        mockClient.On("DescribeDBSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultClusterParameters", func(t *testing.T) {
        input := &docdb.DescribeEngineDefaultClusterParametersInput{}
        output := &docdb.DescribeEngineDefaultClusterParametersOutput{}

        mockClient.On("DescribeEngineDefaultClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventCategories", func(t *testing.T) {
        input := &docdb.DescribeEventCategoriesInput{}
        output := &docdb.DescribeEventCategoriesOutput{}

        mockClient.On("DescribeEventCategories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSubscriptions", func(t *testing.T) {
        input := &docdb.DescribeEventSubscriptionsInput{}
        output := &docdb.DescribeEventSubscriptionsOutput{}

        mockClient.On("DescribeEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &docdb.DescribeEventsInput{}
        output := &docdb.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalClusters", func(t *testing.T) {
        input := &docdb.DescribeGlobalClustersInput{}
        output := &docdb.DescribeGlobalClustersOutput{}

        mockClient.On("DescribeGlobalClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrderableDBInstanceOptions", func(t *testing.T) {
        input := &docdb.DescribeOrderableDBInstanceOptionsInput{}
        output := &docdb.DescribeOrderableDBInstanceOptionsOutput{}

        mockClient.On("DescribeOrderableDBInstanceOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrderableDBInstanceOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePendingMaintenanceActions", func(t *testing.T) {
        input := &docdb.DescribePendingMaintenanceActionsInput{}
        output := &docdb.DescribePendingMaintenanceActionsOutput{}

        mockClient.On("DescribePendingMaintenanceActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePendingMaintenanceActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverDBCluster", func(t *testing.T) {
        input := &docdb.FailoverDBClusterInput{}
        output := &docdb.FailoverDBClusterOutput{}

        mockClient.On("FailoverDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &docdb.ListTagsForResourceInput{}
        output := &docdb.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBCluster", func(t *testing.T) {
        input := &docdb.ModifyDBClusterInput{}
        output := &docdb.ModifyDBClusterOutput{}

        mockClient.On("ModifyDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterParameterGroup", func(t *testing.T) {
        input := &docdb.ModifyDBClusterParameterGroupInput{}
        output := &docdb.ModifyDBClusterParameterGroupOutput{}

        mockClient.On("ModifyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterSnapshotAttribute", func(t *testing.T) {
        input := &docdb.ModifyDBClusterSnapshotAttributeInput{}
        output := &docdb.ModifyDBClusterSnapshotAttributeOutput{}

        mockClient.On("ModifyDBClusterSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBInstance", func(t *testing.T) {
        input := &docdb.ModifyDBInstanceInput{}
        output := &docdb.ModifyDBInstanceOutput{}

        mockClient.On("ModifyDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBSubnetGroup", func(t *testing.T) {
        input := &docdb.ModifyDBSubnetGroupInput{}
        output := &docdb.ModifyDBSubnetGroupOutput{}

        mockClient.On("ModifyDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEventSubscription", func(t *testing.T) {
        input := &docdb.ModifyEventSubscriptionInput{}
        output := &docdb.ModifyEventSubscriptionOutput{}

        mockClient.On("ModifyEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyGlobalCluster", func(t *testing.T) {
        input := &docdb.ModifyGlobalClusterInput{}
        output := &docdb.ModifyGlobalClusterOutput{}

        mockClient.On("ModifyGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootDBInstance", func(t *testing.T) {
        input := &docdb.RebootDBInstanceInput{}
        output := &docdb.RebootDBInstanceOutput{}

        mockClient.On("RebootDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFromGlobalCluster", func(t *testing.T) {
        input := &docdb.RemoveFromGlobalClusterInput{}
        output := &docdb.RemoveFromGlobalClusterOutput{}

        mockClient.On("RemoveFromGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFromGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveSourceIdentifierFromSubscription", func(t *testing.T) {
        input := &docdb.RemoveSourceIdentifierFromSubscriptionInput{}
        output := &docdb.RemoveSourceIdentifierFromSubscriptionOutput{}

        mockClient.On("RemoveSourceIdentifierFromSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveSourceIdentifierFromSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &docdb.RemoveTagsFromResourceInput{}
        output := &docdb.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDBClusterParameterGroup", func(t *testing.T) {
        input := &docdb.ResetDBClusterParameterGroupInput{}
        output := &docdb.ResetDBClusterParameterGroupOutput{}

        mockClient.On("ResetDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterFromSnapshot", func(t *testing.T) {
        input := &docdb.RestoreDBClusterFromSnapshotInput{}
        output := &docdb.RestoreDBClusterFromSnapshotOutput{}

        mockClient.On("RestoreDBClusterFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterToPointInTime", func(t *testing.T) {
        input := &docdb.RestoreDBClusterToPointInTimeInput{}
        output := &docdb.RestoreDBClusterToPointInTimeOutput{}

        mockClient.On("RestoreDBClusterToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDBCluster", func(t *testing.T) {
        input := &docdb.StartDBClusterInput{}
        output := &docdb.StartDBClusterOutput{}

        mockClient.On("StartDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StartDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDBCluster", func(t *testing.T) {
        input := &docdb.StopDBClusterInput{}
        output := &docdb.StopDBClusterOutput{}

        mockClient.On("StopDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StopDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSwitchoverGlobalCluster", func(t *testing.T) {
        input := &docdb.SwitchoverGlobalClusterInput{}
        output := &docdb.SwitchoverGlobalClusterOutput{}

        mockClient.On("SwitchoverGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.SwitchoverGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
