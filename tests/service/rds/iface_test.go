package rds_test

// tests for the rds service interface for this ../../../service/rds/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rds/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rds/rds_iface"
	"github.com/stretchr/testify/assert"
)

func TestRdsServiceCanBeMocked(t *testing.T) {
	var iface rds_iface.IClient
	iface = &rds.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rds.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddRoleToDBCluster", func(t *testing.T) {
        input := &rds.AddRoleToDBClusterInput{}
        output := &rds.AddRoleToDBClusterOutput{}

        mockClient.On("AddRoleToDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.AddRoleToDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddRoleToDBInstance", func(t *testing.T) {
        input := &rds.AddRoleToDBInstanceInput{}
        output := &rds.AddRoleToDBInstanceOutput{}

        mockClient.On("AddRoleToDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.AddRoleToDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddSourceIdentifierToSubscription", func(t *testing.T) {
        input := &rds.AddSourceIdentifierToSubscriptionInput{}
        output := &rds.AddSourceIdentifierToSubscriptionOutput{}

        mockClient.On("AddSourceIdentifierToSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.AddSourceIdentifierToSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &rds.AddTagsToResourceInput{}
        output := &rds.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyPendingMaintenanceAction", func(t *testing.T) {
        input := &rds.ApplyPendingMaintenanceActionInput{}
        output := &rds.ApplyPendingMaintenanceActionOutput{}

        mockClient.On("ApplyPendingMaintenanceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyPendingMaintenanceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeDBSecurityGroupIngress", func(t *testing.T) {
        input := &rds.AuthorizeDBSecurityGroupIngressInput{}
        output := &rds.AuthorizeDBSecurityGroupIngressOutput{}

        mockClient.On("AuthorizeDBSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeDBSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBacktrackDBCluster", func(t *testing.T) {
        input := &rds.BacktrackDBClusterInput{}
        output := &rds.BacktrackDBClusterOutput{}

        mockClient.On("BacktrackDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.BacktrackDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelExportTask", func(t *testing.T) {
        input := &rds.CancelExportTaskInput{}
        output := &rds.CancelExportTaskOutput{}

        mockClient.On("CancelExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterParameterGroup", func(t *testing.T) {
        input := &rds.CopyDBClusterParameterGroupInput{}
        output := &rds.CopyDBClusterParameterGroupOutput{}

        mockClient.On("CopyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBClusterSnapshot", func(t *testing.T) {
        input := &rds.CopyDBClusterSnapshotInput{}
        output := &rds.CopyDBClusterSnapshotOutput{}

        mockClient.On("CopyDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBParameterGroup", func(t *testing.T) {
        input := &rds.CopyDBParameterGroupInput{}
        output := &rds.CopyDBParameterGroupOutput{}

        mockClient.On("CopyDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDBSnapshot", func(t *testing.T) {
        input := &rds.CopyDBSnapshotInput{}
        output := &rds.CopyDBSnapshotOutput{}

        mockClient.On("CopyDBSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDBSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyOptionGroup", func(t *testing.T) {
        input := &rds.CopyOptionGroupInput{}
        output := &rds.CopyOptionGroupOutput{}

        mockClient.On("CopyOptionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CopyOptionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBlueGreenDeployment", func(t *testing.T) {
        input := &rds.CreateBlueGreenDeploymentInput{}
        output := &rds.CreateBlueGreenDeploymentOutput{}

        mockClient.On("CreateBlueGreenDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBlueGreenDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomDBEngineVersion", func(t *testing.T) {
        input := &rds.CreateCustomDBEngineVersionInput{}
        output := &rds.CreateCustomDBEngineVersionOutput{}

        mockClient.On("CreateCustomDBEngineVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomDBEngineVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBCluster", func(t *testing.T) {
        input := &rds.CreateDBClusterInput{}
        output := &rds.CreateDBClusterOutput{}

        mockClient.On("CreateDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterEndpoint", func(t *testing.T) {
        input := &rds.CreateDBClusterEndpointInput{}
        output := &rds.CreateDBClusterEndpointOutput{}

        mockClient.On("CreateDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterParameterGroup", func(t *testing.T) {
        input := &rds.CreateDBClusterParameterGroupInput{}
        output := &rds.CreateDBClusterParameterGroupOutput{}

        mockClient.On("CreateDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBClusterSnapshot", func(t *testing.T) {
        input := &rds.CreateDBClusterSnapshotInput{}
        output := &rds.CreateDBClusterSnapshotOutput{}

        mockClient.On("CreateDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBInstance", func(t *testing.T) {
        input := &rds.CreateDBInstanceInput{}
        output := &rds.CreateDBInstanceOutput{}

        mockClient.On("CreateDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBInstanceReadReplica", func(t *testing.T) {
        input := &rds.CreateDBInstanceReadReplicaInput{}
        output := &rds.CreateDBInstanceReadReplicaOutput{}

        mockClient.On("CreateDBInstanceReadReplica", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBInstanceReadReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBParameterGroup", func(t *testing.T) {
        input := &rds.CreateDBParameterGroupInput{}
        output := &rds.CreateDBParameterGroupOutput{}

        mockClient.On("CreateDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBProxy", func(t *testing.T) {
        input := &rds.CreateDBProxyInput{}
        output := &rds.CreateDBProxyOutput{}

        mockClient.On("CreateDBProxy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBProxyEndpoint", func(t *testing.T) {
        input := &rds.CreateDBProxyEndpointInput{}
        output := &rds.CreateDBProxyEndpointOutput{}

        mockClient.On("CreateDBProxyEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBProxyEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBSecurityGroup", func(t *testing.T) {
        input := &rds.CreateDBSecurityGroupInput{}
        output := &rds.CreateDBSecurityGroupOutput{}

        mockClient.On("CreateDBSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBShardGroup", func(t *testing.T) {
        input := &rds.CreateDBShardGroupInput{}
        output := &rds.CreateDBShardGroupOutput{}

        mockClient.On("CreateDBShardGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBShardGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBSnapshot", func(t *testing.T) {
        input := &rds.CreateDBSnapshotInput{}
        output := &rds.CreateDBSnapshotOutput{}

        mockClient.On("CreateDBSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDBSubnetGroup", func(t *testing.T) {
        input := &rds.CreateDBSubnetGroupInput{}
        output := &rds.CreateDBSubnetGroupOutput{}

        mockClient.On("CreateDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSubscription", func(t *testing.T) {
        input := &rds.CreateEventSubscriptionInput{}
        output := &rds.CreateEventSubscriptionOutput{}

        mockClient.On("CreateEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalCluster", func(t *testing.T) {
        input := &rds.CreateGlobalClusterInput{}
        output := &rds.CreateGlobalClusterOutput{}

        mockClient.On("CreateGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntegration", func(t *testing.T) {
        input := &rds.CreateIntegrationInput{}
        output := &rds.CreateIntegrationOutput{}

        mockClient.On("CreateIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOptionGroup", func(t *testing.T) {
        input := &rds.CreateOptionGroupInput{}
        output := &rds.CreateOptionGroupOutput{}

        mockClient.On("CreateOptionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOptionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTenantDatabase", func(t *testing.T) {
        input := &rds.CreateTenantDatabaseInput{}
        output := &rds.CreateTenantDatabaseOutput{}

        mockClient.On("CreateTenantDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTenantDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBlueGreenDeployment", func(t *testing.T) {
        input := &rds.DeleteBlueGreenDeploymentInput{}
        output := &rds.DeleteBlueGreenDeploymentOutput{}

        mockClient.On("DeleteBlueGreenDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBlueGreenDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomDBEngineVersion", func(t *testing.T) {
        input := &rds.DeleteCustomDBEngineVersionInput{}
        output := &rds.DeleteCustomDBEngineVersionOutput{}

        mockClient.On("DeleteCustomDBEngineVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomDBEngineVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBCluster", func(t *testing.T) {
        input := &rds.DeleteDBClusterInput{}
        output := &rds.DeleteDBClusterOutput{}

        mockClient.On("DeleteDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterAutomatedBackup", func(t *testing.T) {
        input := &rds.DeleteDBClusterAutomatedBackupInput{}
        output := &rds.DeleteDBClusterAutomatedBackupOutput{}

        mockClient.On("DeleteDBClusterAutomatedBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterAutomatedBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterEndpoint", func(t *testing.T) {
        input := &rds.DeleteDBClusterEndpointInput{}
        output := &rds.DeleteDBClusterEndpointOutput{}

        mockClient.On("DeleteDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterParameterGroup", func(t *testing.T) {
        input := &rds.DeleteDBClusterParameterGroupInput{}
        output := &rds.DeleteDBClusterParameterGroupOutput{}

        mockClient.On("DeleteDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBClusterSnapshot", func(t *testing.T) {
        input := &rds.DeleteDBClusterSnapshotInput{}
        output := &rds.DeleteDBClusterSnapshotOutput{}

        mockClient.On("DeleteDBClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBInstance", func(t *testing.T) {
        input := &rds.DeleteDBInstanceInput{}
        output := &rds.DeleteDBInstanceOutput{}

        mockClient.On("DeleteDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBInstanceAutomatedBackup", func(t *testing.T) {
        input := &rds.DeleteDBInstanceAutomatedBackupInput{}
        output := &rds.DeleteDBInstanceAutomatedBackupOutput{}

        mockClient.On("DeleteDBInstanceAutomatedBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBInstanceAutomatedBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBParameterGroup", func(t *testing.T) {
        input := &rds.DeleteDBParameterGroupInput{}
        output := &rds.DeleteDBParameterGroupOutput{}

        mockClient.On("DeleteDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBProxy", func(t *testing.T) {
        input := &rds.DeleteDBProxyInput{}
        output := &rds.DeleteDBProxyOutput{}

        mockClient.On("DeleteDBProxy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBProxyEndpoint", func(t *testing.T) {
        input := &rds.DeleteDBProxyEndpointInput{}
        output := &rds.DeleteDBProxyEndpointOutput{}

        mockClient.On("DeleteDBProxyEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBProxyEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBSecurityGroup", func(t *testing.T) {
        input := &rds.DeleteDBSecurityGroupInput{}
        output := &rds.DeleteDBSecurityGroupOutput{}

        mockClient.On("DeleteDBSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBShardGroup", func(t *testing.T) {
        input := &rds.DeleteDBShardGroupInput{}
        output := &rds.DeleteDBShardGroupOutput{}

        mockClient.On("DeleteDBShardGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBShardGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBSnapshot", func(t *testing.T) {
        input := &rds.DeleteDBSnapshotInput{}
        output := &rds.DeleteDBSnapshotOutput{}

        mockClient.On("DeleteDBSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDBSubnetGroup", func(t *testing.T) {
        input := &rds.DeleteDBSubnetGroupInput{}
        output := &rds.DeleteDBSubnetGroupOutput{}

        mockClient.On("DeleteDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSubscription", func(t *testing.T) {
        input := &rds.DeleteEventSubscriptionInput{}
        output := &rds.DeleteEventSubscriptionOutput{}

        mockClient.On("DeleteEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlobalCluster", func(t *testing.T) {
        input := &rds.DeleteGlobalClusterInput{}
        output := &rds.DeleteGlobalClusterOutput{}

        mockClient.On("DeleteGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegration", func(t *testing.T) {
        input := &rds.DeleteIntegrationInput{}
        output := &rds.DeleteIntegrationOutput{}

        mockClient.On("DeleteIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOptionGroup", func(t *testing.T) {
        input := &rds.DeleteOptionGroupInput{}
        output := &rds.DeleteOptionGroupOutput{}

        mockClient.On("DeleteOptionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOptionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTenantDatabase", func(t *testing.T) {
        input := &rds.DeleteTenantDatabaseInput{}
        output := &rds.DeleteTenantDatabaseOutput{}

        mockClient.On("DeleteTenantDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTenantDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterDBProxyTargets", func(t *testing.T) {
        input := &rds.DeregisterDBProxyTargetsInput{}
        output := &rds.DeregisterDBProxyTargetsOutput{}

        mockClient.On("DeregisterDBProxyTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterDBProxyTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &rds.DescribeAccountAttributesInput{}
        output := &rds.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBlueGreenDeployments", func(t *testing.T) {
        input := &rds.DescribeBlueGreenDeploymentsInput{}
        output := &rds.DescribeBlueGreenDeploymentsOutput{}

        mockClient.On("DescribeBlueGreenDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBlueGreenDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificates", func(t *testing.T) {
        input := &rds.DescribeCertificatesInput{}
        output := &rds.DescribeCertificatesOutput{}

        mockClient.On("DescribeCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterAutomatedBackups", func(t *testing.T) {
        input := &rds.DescribeDBClusterAutomatedBackupsInput{}
        output := &rds.DescribeDBClusterAutomatedBackupsOutput{}

        mockClient.On("DescribeDBClusterAutomatedBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterAutomatedBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterBacktracks", func(t *testing.T) {
        input := &rds.DescribeDBClusterBacktracksInput{}
        output := &rds.DescribeDBClusterBacktracksOutput{}

        mockClient.On("DescribeDBClusterBacktracks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterBacktracks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterEndpoints", func(t *testing.T) {
        input := &rds.DescribeDBClusterEndpointsInput{}
        output := &rds.DescribeDBClusterEndpointsOutput{}

        mockClient.On("DescribeDBClusterEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameterGroups", func(t *testing.T) {
        input := &rds.DescribeDBClusterParameterGroupsInput{}
        output := &rds.DescribeDBClusterParameterGroupsOutput{}

        mockClient.On("DescribeDBClusterParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterParameters", func(t *testing.T) {
        input := &rds.DescribeDBClusterParametersInput{}
        output := &rds.DescribeDBClusterParametersOutput{}

        mockClient.On("DescribeDBClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshotAttributes", func(t *testing.T) {
        input := &rds.DescribeDBClusterSnapshotAttributesInput{}
        output := &rds.DescribeDBClusterSnapshotAttributesOutput{}

        mockClient.On("DescribeDBClusterSnapshotAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshotAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusterSnapshots", func(t *testing.T) {
        input := &rds.DescribeDBClusterSnapshotsInput{}
        output := &rds.DescribeDBClusterSnapshotsOutput{}

        mockClient.On("DescribeDBClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBClusters", func(t *testing.T) {
        input := &rds.DescribeDBClustersInput{}
        output := &rds.DescribeDBClustersOutput{}

        mockClient.On("DescribeDBClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBEngineVersions", func(t *testing.T) {
        input := &rds.DescribeDBEngineVersionsInput{}
        output := &rds.DescribeDBEngineVersionsOutput{}

        mockClient.On("DescribeDBEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBInstanceAutomatedBackups", func(t *testing.T) {
        input := &rds.DescribeDBInstanceAutomatedBackupsInput{}
        output := &rds.DescribeDBInstanceAutomatedBackupsOutput{}

        mockClient.On("DescribeDBInstanceAutomatedBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBInstanceAutomatedBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBInstances", func(t *testing.T) {
        input := &rds.DescribeDBInstancesInput{}
        output := &rds.DescribeDBInstancesOutput{}

        mockClient.On("DescribeDBInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBLogFiles", func(t *testing.T) {
        input := &rds.DescribeDBLogFilesInput{}
        output := &rds.DescribeDBLogFilesOutput{}

        mockClient.On("DescribeDBLogFiles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBLogFiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBParameterGroups", func(t *testing.T) {
        input := &rds.DescribeDBParameterGroupsInput{}
        output := &rds.DescribeDBParameterGroupsOutput{}

        mockClient.On("DescribeDBParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBParameters", func(t *testing.T) {
        input := &rds.DescribeDBParametersInput{}
        output := &rds.DescribeDBParametersOutput{}

        mockClient.On("DescribeDBParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBProxies", func(t *testing.T) {
        input := &rds.DescribeDBProxiesInput{}
        output := &rds.DescribeDBProxiesOutput{}

        mockClient.On("DescribeDBProxies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBProxies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBProxyEndpoints", func(t *testing.T) {
        input := &rds.DescribeDBProxyEndpointsInput{}
        output := &rds.DescribeDBProxyEndpointsOutput{}

        mockClient.On("DescribeDBProxyEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBProxyEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBProxyTargetGroups", func(t *testing.T) {
        input := &rds.DescribeDBProxyTargetGroupsInput{}
        output := &rds.DescribeDBProxyTargetGroupsOutput{}

        mockClient.On("DescribeDBProxyTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBProxyTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBProxyTargets", func(t *testing.T) {
        input := &rds.DescribeDBProxyTargetsInput{}
        output := &rds.DescribeDBProxyTargetsOutput{}

        mockClient.On("DescribeDBProxyTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBProxyTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBRecommendations", func(t *testing.T) {
        input := &rds.DescribeDBRecommendationsInput{}
        output := &rds.DescribeDBRecommendationsOutput{}

        mockClient.On("DescribeDBRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSecurityGroups", func(t *testing.T) {
        input := &rds.DescribeDBSecurityGroupsInput{}
        output := &rds.DescribeDBSecurityGroupsOutput{}

        mockClient.On("DescribeDBSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBShardGroups", func(t *testing.T) {
        input := &rds.DescribeDBShardGroupsInput{}
        output := &rds.DescribeDBShardGroupsOutput{}

        mockClient.On("DescribeDBShardGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBShardGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSnapshotAttributes", func(t *testing.T) {
        input := &rds.DescribeDBSnapshotAttributesInput{}
        output := &rds.DescribeDBSnapshotAttributesOutput{}

        mockClient.On("DescribeDBSnapshotAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSnapshotAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSnapshotTenantDatabases", func(t *testing.T) {
        input := &rds.DescribeDBSnapshotTenantDatabasesInput{}
        output := &rds.DescribeDBSnapshotTenantDatabasesOutput{}

        mockClient.On("DescribeDBSnapshotTenantDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSnapshotTenantDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSnapshots", func(t *testing.T) {
        input := &rds.DescribeDBSnapshotsInput{}
        output := &rds.DescribeDBSnapshotsOutput{}

        mockClient.On("DescribeDBSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDBSubnetGroups", func(t *testing.T) {
        input := &rds.DescribeDBSubnetGroupsInput{}
        output := &rds.DescribeDBSubnetGroupsOutput{}

        mockClient.On("DescribeDBSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDBSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultClusterParameters", func(t *testing.T) {
        input := &rds.DescribeEngineDefaultClusterParametersInput{}
        output := &rds.DescribeEngineDefaultClusterParametersOutput{}

        mockClient.On("DescribeEngineDefaultClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineDefaultParameters", func(t *testing.T) {
        input := &rds.DescribeEngineDefaultParametersInput{}
        output := &rds.DescribeEngineDefaultParametersOutput{}

        mockClient.On("DescribeEngineDefaultParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineDefaultParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventCategories", func(t *testing.T) {
        input := &rds.DescribeEventCategoriesInput{}
        output := &rds.DescribeEventCategoriesOutput{}

        mockClient.On("DescribeEventCategories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSubscriptions", func(t *testing.T) {
        input := &rds.DescribeEventSubscriptionsInput{}
        output := &rds.DescribeEventSubscriptionsOutput{}

        mockClient.On("DescribeEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &rds.DescribeEventsInput{}
        output := &rds.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportTasks", func(t *testing.T) {
        input := &rds.DescribeExportTasksInput{}
        output := &rds.DescribeExportTasksOutput{}

        mockClient.On("DescribeExportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalClusters", func(t *testing.T) {
        input := &rds.DescribeGlobalClustersInput{}
        output := &rds.DescribeGlobalClustersOutput{}

        mockClient.On("DescribeGlobalClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIntegrations", func(t *testing.T) {
        input := &rds.DescribeIntegrationsInput{}
        output := &rds.DescribeIntegrationsOutput{}

        mockClient.On("DescribeIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOptionGroupOptions", func(t *testing.T) {
        input := &rds.DescribeOptionGroupOptionsInput{}
        output := &rds.DescribeOptionGroupOptionsOutput{}

        mockClient.On("DescribeOptionGroupOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOptionGroupOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOptionGroups", func(t *testing.T) {
        input := &rds.DescribeOptionGroupsInput{}
        output := &rds.DescribeOptionGroupsOutput{}

        mockClient.On("DescribeOptionGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOptionGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrderableDBInstanceOptions", func(t *testing.T) {
        input := &rds.DescribeOrderableDBInstanceOptionsInput{}
        output := &rds.DescribeOrderableDBInstanceOptionsOutput{}

        mockClient.On("DescribeOrderableDBInstanceOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrderableDBInstanceOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePendingMaintenanceActions", func(t *testing.T) {
        input := &rds.DescribePendingMaintenanceActionsInput{}
        output := &rds.DescribePendingMaintenanceActionsOutput{}

        mockClient.On("DescribePendingMaintenanceActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePendingMaintenanceActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedDBInstances", func(t *testing.T) {
        input := &rds.DescribeReservedDBInstancesInput{}
        output := &rds.DescribeReservedDBInstancesOutput{}

        mockClient.On("DescribeReservedDBInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedDBInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedDBInstancesOfferings", func(t *testing.T) {
        input := &rds.DescribeReservedDBInstancesOfferingsInput{}
        output := &rds.DescribeReservedDBInstancesOfferingsOutput{}

        mockClient.On("DescribeReservedDBInstancesOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedDBInstancesOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSourceRegions", func(t *testing.T) {
        input := &rds.DescribeSourceRegionsInput{}
        output := &rds.DescribeSourceRegionsOutput{}

        mockClient.On("DescribeSourceRegions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSourceRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTenantDatabases", func(t *testing.T) {
        input := &rds.DescribeTenantDatabasesInput{}
        output := &rds.DescribeTenantDatabasesOutput{}

        mockClient.On("DescribeTenantDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTenantDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeValidDBInstanceModifications", func(t *testing.T) {
        input := &rds.DescribeValidDBInstanceModificationsInput{}
        output := &rds.DescribeValidDBInstanceModificationsOutput{}

        mockClient.On("DescribeValidDBInstanceModifications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeValidDBInstanceModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableHttpEndpoint", func(t *testing.T) {
        input := &rds.DisableHttpEndpointInput{}
        output := &rds.DisableHttpEndpointOutput{}

        mockClient.On("DisableHttpEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DisableHttpEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDownloadDBLogFilePortion", func(t *testing.T) {
        input := &rds.DownloadDBLogFilePortionInput{}
        output := &rds.DownloadDBLogFilePortionOutput{}

        mockClient.On("DownloadDBLogFilePortion", ctx, input).Return(output, nil)

        result, err := mockClient.DownloadDBLogFilePortion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableHttpEndpoint", func(t *testing.T) {
        input := &rds.EnableHttpEndpointInput{}
        output := &rds.EnableHttpEndpointOutput{}

        mockClient.On("EnableHttpEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.EnableHttpEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverDBCluster", func(t *testing.T) {
        input := &rds.FailoverDBClusterInput{}
        output := &rds.FailoverDBClusterOutput{}

        mockClient.On("FailoverDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverGlobalCluster", func(t *testing.T) {
        input := &rds.FailoverGlobalClusterInput{}
        output := &rds.FailoverGlobalClusterOutput{}

        mockClient.On("FailoverGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &rds.ListTagsForResourceInput{}
        output := &rds.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyActivityStream", func(t *testing.T) {
        input := &rds.ModifyActivityStreamInput{}
        output := &rds.ModifyActivityStreamOutput{}

        mockClient.On("ModifyActivityStream", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyActivityStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCertificates", func(t *testing.T) {
        input := &rds.ModifyCertificatesInput{}
        output := &rds.ModifyCertificatesOutput{}

        mockClient.On("ModifyCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCurrentDBClusterCapacity", func(t *testing.T) {
        input := &rds.ModifyCurrentDBClusterCapacityInput{}
        output := &rds.ModifyCurrentDBClusterCapacityOutput{}

        mockClient.On("ModifyCurrentDBClusterCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCurrentDBClusterCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCustomDBEngineVersion", func(t *testing.T) {
        input := &rds.ModifyCustomDBEngineVersionInput{}
        output := &rds.ModifyCustomDBEngineVersionOutput{}

        mockClient.On("ModifyCustomDBEngineVersion", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCustomDBEngineVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBCluster", func(t *testing.T) {
        input := &rds.ModifyDBClusterInput{}
        output := &rds.ModifyDBClusterOutput{}

        mockClient.On("ModifyDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterEndpoint", func(t *testing.T) {
        input := &rds.ModifyDBClusterEndpointInput{}
        output := &rds.ModifyDBClusterEndpointOutput{}

        mockClient.On("ModifyDBClusterEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterParameterGroup", func(t *testing.T) {
        input := &rds.ModifyDBClusterParameterGroupInput{}
        output := &rds.ModifyDBClusterParameterGroupOutput{}

        mockClient.On("ModifyDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBClusterSnapshotAttribute", func(t *testing.T) {
        input := &rds.ModifyDBClusterSnapshotAttributeInput{}
        output := &rds.ModifyDBClusterSnapshotAttributeOutput{}

        mockClient.On("ModifyDBClusterSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBClusterSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBInstance", func(t *testing.T) {
        input := &rds.ModifyDBInstanceInput{}
        output := &rds.ModifyDBInstanceOutput{}

        mockClient.On("ModifyDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBParameterGroup", func(t *testing.T) {
        input := &rds.ModifyDBParameterGroupInput{}
        output := &rds.ModifyDBParameterGroupOutput{}

        mockClient.On("ModifyDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBProxy", func(t *testing.T) {
        input := &rds.ModifyDBProxyInput{}
        output := &rds.ModifyDBProxyOutput{}

        mockClient.On("ModifyDBProxy", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBProxyEndpoint", func(t *testing.T) {
        input := &rds.ModifyDBProxyEndpointInput{}
        output := &rds.ModifyDBProxyEndpointOutput{}

        mockClient.On("ModifyDBProxyEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBProxyEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBProxyTargetGroup", func(t *testing.T) {
        input := &rds.ModifyDBProxyTargetGroupInput{}
        output := &rds.ModifyDBProxyTargetGroupOutput{}

        mockClient.On("ModifyDBProxyTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBProxyTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBRecommendation", func(t *testing.T) {
        input := &rds.ModifyDBRecommendationInput{}
        output := &rds.ModifyDBRecommendationOutput{}

        mockClient.On("ModifyDBRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBShardGroup", func(t *testing.T) {
        input := &rds.ModifyDBShardGroupInput{}
        output := &rds.ModifyDBShardGroupOutput{}

        mockClient.On("ModifyDBShardGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBShardGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBSnapshot", func(t *testing.T) {
        input := &rds.ModifyDBSnapshotInput{}
        output := &rds.ModifyDBSnapshotOutput{}

        mockClient.On("ModifyDBSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBSnapshotAttribute", func(t *testing.T) {
        input := &rds.ModifyDBSnapshotAttributeInput{}
        output := &rds.ModifyDBSnapshotAttributeOutput{}

        mockClient.On("ModifyDBSnapshotAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBSnapshotAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDBSubnetGroup", func(t *testing.T) {
        input := &rds.ModifyDBSubnetGroupInput{}
        output := &rds.ModifyDBSubnetGroupOutput{}

        mockClient.On("ModifyDBSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDBSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEventSubscription", func(t *testing.T) {
        input := &rds.ModifyEventSubscriptionInput{}
        output := &rds.ModifyEventSubscriptionOutput{}

        mockClient.On("ModifyEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyGlobalCluster", func(t *testing.T) {
        input := &rds.ModifyGlobalClusterInput{}
        output := &rds.ModifyGlobalClusterOutput{}

        mockClient.On("ModifyGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyIntegration", func(t *testing.T) {
        input := &rds.ModifyIntegrationInput{}
        output := &rds.ModifyIntegrationOutput{}

        mockClient.On("ModifyIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyOptionGroup", func(t *testing.T) {
        input := &rds.ModifyOptionGroupInput{}
        output := &rds.ModifyOptionGroupOutput{}

        mockClient.On("ModifyOptionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyOptionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTenantDatabase", func(t *testing.T) {
        input := &rds.ModifyTenantDatabaseInput{}
        output := &rds.ModifyTenantDatabaseOutput{}

        mockClient.On("ModifyTenantDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTenantDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromoteReadReplica", func(t *testing.T) {
        input := &rds.PromoteReadReplicaInput{}
        output := &rds.PromoteReadReplicaOutput{}

        mockClient.On("PromoteReadReplica", ctx, input).Return(output, nil)

        result, err := mockClient.PromoteReadReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromoteReadReplicaDBCluster", func(t *testing.T) {
        input := &rds.PromoteReadReplicaDBClusterInput{}
        output := &rds.PromoteReadReplicaDBClusterOutput{}

        mockClient.On("PromoteReadReplicaDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.PromoteReadReplicaDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedDBInstancesOffering", func(t *testing.T) {
        input := &rds.PurchaseReservedDBInstancesOfferingInput{}
        output := &rds.PurchaseReservedDBInstancesOfferingOutput{}

        mockClient.On("PurchaseReservedDBInstancesOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedDBInstancesOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootDBCluster", func(t *testing.T) {
        input := &rds.RebootDBClusterInput{}
        output := &rds.RebootDBClusterOutput{}

        mockClient.On("RebootDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RebootDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootDBInstance", func(t *testing.T) {
        input := &rds.RebootDBInstanceInput{}
        output := &rds.RebootDBInstanceOutput{}

        mockClient.On("RebootDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootDBShardGroup", func(t *testing.T) {
        input := &rds.RebootDBShardGroupInput{}
        output := &rds.RebootDBShardGroupOutput{}

        mockClient.On("RebootDBShardGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RebootDBShardGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDBProxyTargets", func(t *testing.T) {
        input := &rds.RegisterDBProxyTargetsInput{}
        output := &rds.RegisterDBProxyTargetsOutput{}

        mockClient.On("RegisterDBProxyTargets", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDBProxyTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFromGlobalCluster", func(t *testing.T) {
        input := &rds.RemoveFromGlobalClusterInput{}
        output := &rds.RemoveFromGlobalClusterOutput{}

        mockClient.On("RemoveFromGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFromGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRoleFromDBCluster", func(t *testing.T) {
        input := &rds.RemoveRoleFromDBClusterInput{}
        output := &rds.RemoveRoleFromDBClusterOutput{}

        mockClient.On("RemoveRoleFromDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRoleFromDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRoleFromDBInstance", func(t *testing.T) {
        input := &rds.RemoveRoleFromDBInstanceInput{}
        output := &rds.RemoveRoleFromDBInstanceOutput{}

        mockClient.On("RemoveRoleFromDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRoleFromDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveSourceIdentifierFromSubscription", func(t *testing.T) {
        input := &rds.RemoveSourceIdentifierFromSubscriptionInput{}
        output := &rds.RemoveSourceIdentifierFromSubscriptionOutput{}

        mockClient.On("RemoveSourceIdentifierFromSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveSourceIdentifierFromSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &rds.RemoveTagsFromResourceInput{}
        output := &rds.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDBClusterParameterGroup", func(t *testing.T) {
        input := &rds.ResetDBClusterParameterGroupInput{}
        output := &rds.ResetDBClusterParameterGroupOutput{}

        mockClient.On("ResetDBClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDBClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDBParameterGroup", func(t *testing.T) {
        input := &rds.ResetDBParameterGroupInput{}
        output := &rds.ResetDBParameterGroupOutput{}

        mockClient.On("ResetDBParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDBParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterFromS3", func(t *testing.T) {
        input := &rds.RestoreDBClusterFromS3Input{}
        output := &rds.RestoreDBClusterFromS3Output{}

        mockClient.On("RestoreDBClusterFromS3", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterFromS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterFromSnapshot", func(t *testing.T) {
        input := &rds.RestoreDBClusterFromSnapshotInput{}
        output := &rds.RestoreDBClusterFromSnapshotOutput{}

        mockClient.On("RestoreDBClusterFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBClusterToPointInTime", func(t *testing.T) {
        input := &rds.RestoreDBClusterToPointInTimeInput{}
        output := &rds.RestoreDBClusterToPointInTimeOutput{}

        mockClient.On("RestoreDBClusterToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBClusterToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBInstanceFromDBSnapshot", func(t *testing.T) {
        input := &rds.RestoreDBInstanceFromDBSnapshotInput{}
        output := &rds.RestoreDBInstanceFromDBSnapshotOutput{}

        mockClient.On("RestoreDBInstanceFromDBSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBInstanceFromDBSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBInstanceFromS3", func(t *testing.T) {
        input := &rds.RestoreDBInstanceFromS3Input{}
        output := &rds.RestoreDBInstanceFromS3Output{}

        mockClient.On("RestoreDBInstanceFromS3", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBInstanceFromS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDBInstanceToPointInTime", func(t *testing.T) {
        input := &rds.RestoreDBInstanceToPointInTimeInput{}
        output := &rds.RestoreDBInstanceToPointInTimeOutput{}

        mockClient.On("RestoreDBInstanceToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDBInstanceToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeDBSecurityGroupIngress", func(t *testing.T) {
        input := &rds.RevokeDBSecurityGroupIngressInput{}
        output := &rds.RevokeDBSecurityGroupIngressOutput{}

        mockClient.On("RevokeDBSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeDBSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartActivityStream", func(t *testing.T) {
        input := &rds.StartActivityStreamInput{}
        output := &rds.StartActivityStreamOutput{}

        mockClient.On("StartActivityStream", ctx, input).Return(output, nil)

        result, err := mockClient.StartActivityStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDBCluster", func(t *testing.T) {
        input := &rds.StartDBClusterInput{}
        output := &rds.StartDBClusterOutput{}

        mockClient.On("StartDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StartDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDBInstance", func(t *testing.T) {
        input := &rds.StartDBInstanceInput{}
        output := &rds.StartDBInstanceOutput{}

        mockClient.On("StartDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StartDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDBInstanceAutomatedBackupsReplication", func(t *testing.T) {
        input := &rds.StartDBInstanceAutomatedBackupsReplicationInput{}
        output := &rds.StartDBInstanceAutomatedBackupsReplicationOutput{}

        mockClient.On("StartDBInstanceAutomatedBackupsReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartDBInstanceAutomatedBackupsReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExportTask", func(t *testing.T) {
        input := &rds.StartExportTaskInput{}
        output := &rds.StartExportTaskOutput{}

        mockClient.On("StartExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopActivityStream", func(t *testing.T) {
        input := &rds.StopActivityStreamInput{}
        output := &rds.StopActivityStreamOutput{}

        mockClient.On("StopActivityStream", ctx, input).Return(output, nil)

        result, err := mockClient.StopActivityStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDBCluster", func(t *testing.T) {
        input := &rds.StopDBClusterInput{}
        output := &rds.StopDBClusterOutput{}

        mockClient.On("StopDBCluster", ctx, input).Return(output, nil)

        result, err := mockClient.StopDBCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDBInstance", func(t *testing.T) {
        input := &rds.StopDBInstanceInput{}
        output := &rds.StopDBInstanceOutput{}

        mockClient.On("StopDBInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StopDBInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDBInstanceAutomatedBackupsReplication", func(t *testing.T) {
        input := &rds.StopDBInstanceAutomatedBackupsReplicationInput{}
        output := &rds.StopDBInstanceAutomatedBackupsReplicationOutput{}

        mockClient.On("StopDBInstanceAutomatedBackupsReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopDBInstanceAutomatedBackupsReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSwitchoverBlueGreenDeployment", func(t *testing.T) {
        input := &rds.SwitchoverBlueGreenDeploymentInput{}
        output := &rds.SwitchoverBlueGreenDeploymentOutput{}

        mockClient.On("SwitchoverBlueGreenDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.SwitchoverBlueGreenDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSwitchoverGlobalCluster", func(t *testing.T) {
        input := &rds.SwitchoverGlobalClusterInput{}
        output := &rds.SwitchoverGlobalClusterOutput{}

        mockClient.On("SwitchoverGlobalCluster", ctx, input).Return(output, nil)

        result, err := mockClient.SwitchoverGlobalCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSwitchoverReadReplica", func(t *testing.T) {
        input := &rds.SwitchoverReadReplicaInput{}
        output := &rds.SwitchoverReadReplicaOutput{}

        mockClient.On("SwitchoverReadReplica", ctx, input).Return(output, nil)

        result, err := mockClient.SwitchoverReadReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
