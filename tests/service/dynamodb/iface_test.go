package dynamodb_test

// tests for the dynamodb service interface for this ../../../service/dynamodb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dynamodb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dynamodb/dynamodb_iface"
	"github.com/stretchr/testify/assert"
)

func TestDynamodbServiceCanBeMocked(t *testing.T) {
	var iface dynamodb_iface.IClient
	iface = &dynamodb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := dynamodb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchExecuteStatement", func(t *testing.T) {
        input := &dynamodb.BatchExecuteStatementInput{}
        output := &dynamodb.BatchExecuteStatementOutput{}

        mockClient.On("BatchExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.BatchExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetItem", func(t *testing.T) {
        input := &dynamodb.BatchGetItemInput{}
        output := &dynamodb.BatchGetItemOutput{}

        mockClient.On("BatchGetItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchWriteItem", func(t *testing.T) {
        input := &dynamodb.BatchWriteItemInput{}
        output := &dynamodb.BatchWriteItemOutput{}

        mockClient.On("BatchWriteItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchWriteItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackup", func(t *testing.T) {
        input := &dynamodb.CreateBackupInput{}
        output := &dynamodb.CreateBackupOutput{}

        mockClient.On("CreateBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlobalTable", func(t *testing.T) {
        input := &dynamodb.CreateGlobalTableInput{}
        output := &dynamodb.CreateGlobalTableOutput{}

        mockClient.On("CreateGlobalTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlobalTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTable", func(t *testing.T) {
        input := &dynamodb.CreateTableInput{}
        output := &dynamodb.CreateTableOutput{}

        mockClient.On("CreateTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackup", func(t *testing.T) {
        input := &dynamodb.DeleteBackupInput{}
        output := &dynamodb.DeleteBackupOutput{}

        mockClient.On("DeleteBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteItem", func(t *testing.T) {
        input := &dynamodb.DeleteItemInput{}
        output := &dynamodb.DeleteItemOutput{}

        mockClient.On("DeleteItem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &dynamodb.DeleteResourcePolicyInput{}
        output := &dynamodb.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTable", func(t *testing.T) {
        input := &dynamodb.DeleteTableInput{}
        output := &dynamodb.DeleteTableOutput{}

        mockClient.On("DeleteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackup", func(t *testing.T) {
        input := &dynamodb.DescribeBackupInput{}
        output := &dynamodb.DescribeBackupOutput{}

        mockClient.On("DescribeBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContinuousBackups", func(t *testing.T) {
        input := &dynamodb.DescribeContinuousBackupsInput{}
        output := &dynamodb.DescribeContinuousBackupsOutput{}

        mockClient.On("DescribeContinuousBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContinuousBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContributorInsights", func(t *testing.T) {
        input := &dynamodb.DescribeContributorInsightsInput{}
        output := &dynamodb.DescribeContributorInsightsOutput{}

        mockClient.On("DescribeContributorInsights", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContributorInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoints", func(t *testing.T) {
        input := &dynamodb.DescribeEndpointsInput{}
        output := &dynamodb.DescribeEndpointsOutput{}

        mockClient.On("DescribeEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExport", func(t *testing.T) {
        input := &dynamodb.DescribeExportInput{}
        output := &dynamodb.DescribeExportOutput{}

        mockClient.On("DescribeExport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalTable", func(t *testing.T) {
        input := &dynamodb.DescribeGlobalTableInput{}
        output := &dynamodb.DescribeGlobalTableOutput{}

        mockClient.On("DescribeGlobalTable", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalTableSettings", func(t *testing.T) {
        input := &dynamodb.DescribeGlobalTableSettingsInput{}
        output := &dynamodb.DescribeGlobalTableSettingsOutput{}

        mockClient.On("DescribeGlobalTableSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalTableSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImport", func(t *testing.T) {
        input := &dynamodb.DescribeImportInput{}
        output := &dynamodb.DescribeImportOutput{}

        mockClient.On("DescribeImport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKinesisStreamingDestination", func(t *testing.T) {
        input := &dynamodb.DescribeKinesisStreamingDestinationInput{}
        output := &dynamodb.DescribeKinesisStreamingDestinationOutput{}

        mockClient.On("DescribeKinesisStreamingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKinesisStreamingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLimits", func(t *testing.T) {
        input := &dynamodb.DescribeLimitsInput{}
        output := &dynamodb.DescribeLimitsOutput{}

        mockClient.On("DescribeLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTable", func(t *testing.T) {
        input := &dynamodb.DescribeTableInput{}
        output := &dynamodb.DescribeTableOutput{}

        mockClient.On("DescribeTable", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTableReplicaAutoScaling", func(t *testing.T) {
        input := &dynamodb.DescribeTableReplicaAutoScalingInput{}
        output := &dynamodb.DescribeTableReplicaAutoScalingOutput{}

        mockClient.On("DescribeTableReplicaAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTableReplicaAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTimeToLive", func(t *testing.T) {
        input := &dynamodb.DescribeTimeToLiveInput{}
        output := &dynamodb.DescribeTimeToLiveOutput{}

        mockClient.On("DescribeTimeToLive", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTimeToLive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableKinesisStreamingDestination", func(t *testing.T) {
        input := &dynamodb.DisableKinesisStreamingDestinationInput{}
        output := &dynamodb.DisableKinesisStreamingDestinationOutput{}

        mockClient.On("DisableKinesisStreamingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DisableKinesisStreamingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableKinesisStreamingDestination", func(t *testing.T) {
        input := &dynamodb.EnableKinesisStreamingDestinationInput{}
        output := &dynamodb.EnableKinesisStreamingDestinationOutput{}

        mockClient.On("EnableKinesisStreamingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.EnableKinesisStreamingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteStatement", func(t *testing.T) {
        input := &dynamodb.ExecuteStatementInput{}
        output := &dynamodb.ExecuteStatementOutput{}

        mockClient.On("ExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteTransaction", func(t *testing.T) {
        input := &dynamodb.ExecuteTransactionInput{}
        output := &dynamodb.ExecuteTransactionOutput{}

        mockClient.On("ExecuteTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportTableToPointInTime", func(t *testing.T) {
        input := &dynamodb.ExportTableToPointInTimeInput{}
        output := &dynamodb.ExportTableToPointInTimeOutput{}

        mockClient.On("ExportTableToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.ExportTableToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetItem", func(t *testing.T) {
        input := &dynamodb.GetItemInput{}
        output := &dynamodb.GetItemOutput{}

        mockClient.On("GetItem", ctx, input).Return(output, nil)

        result, err := mockClient.GetItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &dynamodb.GetResourcePolicyInput{}
        output := &dynamodb.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportTable", func(t *testing.T) {
        input := &dynamodb.ImportTableInput{}
        output := &dynamodb.ImportTableOutput{}

        mockClient.On("ImportTable", ctx, input).Return(output, nil)

        result, err := mockClient.ImportTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackups", func(t *testing.T) {
        input := &dynamodb.ListBackupsInput{}
        output := &dynamodb.ListBackupsOutput{}

        mockClient.On("ListBackups", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContributorInsights", func(t *testing.T) {
        input := &dynamodb.ListContributorInsightsInput{}
        output := &dynamodb.ListContributorInsightsOutput{}

        mockClient.On("ListContributorInsights", ctx, input).Return(output, nil)

        result, err := mockClient.ListContributorInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExports", func(t *testing.T) {
        input := &dynamodb.ListExportsInput{}
        output := &dynamodb.ListExportsOutput{}

        mockClient.On("ListExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGlobalTables", func(t *testing.T) {
        input := &dynamodb.ListGlobalTablesInput{}
        output := &dynamodb.ListGlobalTablesOutput{}

        mockClient.On("ListGlobalTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListGlobalTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImports", func(t *testing.T) {
        input := &dynamodb.ListImportsInput{}
        output := &dynamodb.ListImportsOutput{}

        mockClient.On("ListImports", ctx, input).Return(output, nil)

        result, err := mockClient.ListImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTables", func(t *testing.T) {
        input := &dynamodb.ListTablesInput{}
        output := &dynamodb.ListTablesOutput{}

        mockClient.On("ListTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsOfResource", func(t *testing.T) {
        input := &dynamodb.ListTagsOfResourceInput{}
        output := &dynamodb.ListTagsOfResourceOutput{}

        mockClient.On("ListTagsOfResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsOfResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutItem", func(t *testing.T) {
        input := &dynamodb.PutItemInput{}
        output := &dynamodb.PutItemOutput{}

        mockClient.On("PutItem", ctx, input).Return(output, nil)

        result, err := mockClient.PutItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &dynamodb.PutResourcePolicyInput{}
        output := &dynamodb.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQuery", func(t *testing.T) {
        input := &dynamodb.QueryInput{}
        output := &dynamodb.QueryOutput{}

        mockClient.On("Query", ctx, input).Return(output, nil)

        result, err := mockClient.Query(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTableFromBackup", func(t *testing.T) {
        input := &dynamodb.RestoreTableFromBackupInput{}
        output := &dynamodb.RestoreTableFromBackupOutput{}

        mockClient.On("RestoreTableFromBackup", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTableFromBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTableToPointInTime", func(t *testing.T) {
        input := &dynamodb.RestoreTableToPointInTimeInput{}
        output := &dynamodb.RestoreTableToPointInTimeOutput{}

        mockClient.On("RestoreTableToPointInTime", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTableToPointInTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestScan", func(t *testing.T) {
        input := &dynamodb.ScanInput{}
        output := &dynamodb.ScanOutput{}

        mockClient.On("Scan", ctx, input).Return(output, nil)

        result, err := mockClient.Scan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &dynamodb.TagResourceInput{}
        output := &dynamodb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransactGetItems", func(t *testing.T) {
        input := &dynamodb.TransactGetItemsInput{}
        output := &dynamodb.TransactGetItemsOutput{}

        mockClient.On("TransactGetItems", ctx, input).Return(output, nil)

        result, err := mockClient.TransactGetItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransactWriteItems", func(t *testing.T) {
        input := &dynamodb.TransactWriteItemsInput{}
        output := &dynamodb.TransactWriteItemsOutput{}

        mockClient.On("TransactWriteItems", ctx, input).Return(output, nil)

        result, err := mockClient.TransactWriteItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &dynamodb.UntagResourceInput{}
        output := &dynamodb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContinuousBackups", func(t *testing.T) {
        input := &dynamodb.UpdateContinuousBackupsInput{}
        output := &dynamodb.UpdateContinuousBackupsOutput{}

        mockClient.On("UpdateContinuousBackups", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContinuousBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContributorInsights", func(t *testing.T) {
        input := &dynamodb.UpdateContributorInsightsInput{}
        output := &dynamodb.UpdateContributorInsightsOutput{}

        mockClient.On("UpdateContributorInsights", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContributorInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalTable", func(t *testing.T) {
        input := &dynamodb.UpdateGlobalTableInput{}
        output := &dynamodb.UpdateGlobalTableOutput{}

        mockClient.On("UpdateGlobalTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalTableSettings", func(t *testing.T) {
        input := &dynamodb.UpdateGlobalTableSettingsInput{}
        output := &dynamodb.UpdateGlobalTableSettingsOutput{}

        mockClient.On("UpdateGlobalTableSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalTableSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateItem", func(t *testing.T) {
        input := &dynamodb.UpdateItemInput{}
        output := &dynamodb.UpdateItemOutput{}

        mockClient.On("UpdateItem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKinesisStreamingDestination", func(t *testing.T) {
        input := &dynamodb.UpdateKinesisStreamingDestinationInput{}
        output := &dynamodb.UpdateKinesisStreamingDestinationOutput{}

        mockClient.On("UpdateKinesisStreamingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKinesisStreamingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTable", func(t *testing.T) {
        input := &dynamodb.UpdateTableInput{}
        output := &dynamodb.UpdateTableOutput{}

        mockClient.On("UpdateTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTableReplicaAutoScaling", func(t *testing.T) {
        input := &dynamodb.UpdateTableReplicaAutoScalingInput{}
        output := &dynamodb.UpdateTableReplicaAutoScalingOutput{}

        mockClient.On("UpdateTableReplicaAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTableReplicaAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTimeToLive", func(t *testing.T) {
        input := &dynamodb.UpdateTimeToLiveInput{}
        output := &dynamodb.UpdateTimeToLiveOutput{}

        mockClient.On("UpdateTimeToLive", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTimeToLive(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
