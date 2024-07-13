package lakeformation_test

// tests for the lakeformation service interface for this ../../../service/lakeformation/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lakeformation/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lakeformation/lakeformation_iface"
	"github.com/stretchr/testify/assert"
)

func TestLakeformationServiceCanBeMocked(t *testing.T) {
	var iface lakeformation_iface.IClient
	iface = &lakeformation.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lakeformation.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddLFTagsToResource", func(t *testing.T) {
        input := &lakeformation.AddLFTagsToResourceInput{}
        output := &lakeformation.AddLFTagsToResourceOutput{}

        mockClient.On("AddLFTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddLFTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeDecoratedRoleWithSAML", func(t *testing.T) {
        input := &lakeformation.AssumeDecoratedRoleWithSAMLInput{}
        output := &lakeformation.AssumeDecoratedRoleWithSAMLOutput{}

        mockClient.On("AssumeDecoratedRoleWithSAML", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeDecoratedRoleWithSAML(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGrantPermissions", func(t *testing.T) {
        input := &lakeformation.BatchGrantPermissionsInput{}
        output := &lakeformation.BatchGrantPermissionsOutput{}

        mockClient.On("BatchGrantPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGrantPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchRevokePermissions", func(t *testing.T) {
        input := &lakeformation.BatchRevokePermissionsInput{}
        output := &lakeformation.BatchRevokePermissionsOutput{}

        mockClient.On("BatchRevokePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchRevokePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelTransaction", func(t *testing.T) {
        input := &lakeformation.CancelTransactionInput{}
        output := &lakeformation.CancelTransactionOutput{}

        mockClient.On("CancelTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.CancelTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCommitTransaction", func(t *testing.T) {
        input := &lakeformation.CommitTransactionInput{}
        output := &lakeformation.CommitTransactionOutput{}

        mockClient.On("CommitTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.CommitTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataCellsFilter", func(t *testing.T) {
        input := &lakeformation.CreateDataCellsFilterInput{}
        output := &lakeformation.CreateDataCellsFilterOutput{}

        mockClient.On("CreateDataCellsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataCellsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLFTag", func(t *testing.T) {
        input := &lakeformation.CreateLFTagInput{}
        output := &lakeformation.CreateLFTagOutput{}

        mockClient.On("CreateLFTag", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLFTag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLakeFormationIdentityCenterConfiguration", func(t *testing.T) {
        input := &lakeformation.CreateLakeFormationIdentityCenterConfigurationInput{}
        output := &lakeformation.CreateLakeFormationIdentityCenterConfigurationOutput{}

        mockClient.On("CreateLakeFormationIdentityCenterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLakeFormationIdentityCenterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLakeFormationOptIn", func(t *testing.T) {
        input := &lakeformation.CreateLakeFormationOptInInput{}
        output := &lakeformation.CreateLakeFormationOptInOutput{}

        mockClient.On("CreateLakeFormationOptIn", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLakeFormationOptIn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataCellsFilter", func(t *testing.T) {
        input := &lakeformation.DeleteDataCellsFilterInput{}
        output := &lakeformation.DeleteDataCellsFilterOutput{}

        mockClient.On("DeleteDataCellsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataCellsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLFTag", func(t *testing.T) {
        input := &lakeformation.DeleteLFTagInput{}
        output := &lakeformation.DeleteLFTagOutput{}

        mockClient.On("DeleteLFTag", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLFTag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLakeFormationIdentityCenterConfiguration", func(t *testing.T) {
        input := &lakeformation.DeleteLakeFormationIdentityCenterConfigurationInput{}
        output := &lakeformation.DeleteLakeFormationIdentityCenterConfigurationOutput{}

        mockClient.On("DeleteLakeFormationIdentityCenterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLakeFormationIdentityCenterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLakeFormationOptIn", func(t *testing.T) {
        input := &lakeformation.DeleteLakeFormationOptInInput{}
        output := &lakeformation.DeleteLakeFormationOptInOutput{}

        mockClient.On("DeleteLakeFormationOptIn", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLakeFormationOptIn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObjectsOnCancel", func(t *testing.T) {
        input := &lakeformation.DeleteObjectsOnCancelInput{}
        output := &lakeformation.DeleteObjectsOnCancelOutput{}

        mockClient.On("DeleteObjectsOnCancel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObjectsOnCancel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterResource", func(t *testing.T) {
        input := &lakeformation.DeregisterResourceInput{}
        output := &lakeformation.DeregisterResourceOutput{}

        mockClient.On("DeregisterResource", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLakeFormationIdentityCenterConfiguration", func(t *testing.T) {
        input := &lakeformation.DescribeLakeFormationIdentityCenterConfigurationInput{}
        output := &lakeformation.DescribeLakeFormationIdentityCenterConfigurationOutput{}

        mockClient.On("DescribeLakeFormationIdentityCenterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLakeFormationIdentityCenterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResource", func(t *testing.T) {
        input := &lakeformation.DescribeResourceInput{}
        output := &lakeformation.DescribeResourceOutput{}

        mockClient.On("DescribeResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransaction", func(t *testing.T) {
        input := &lakeformation.DescribeTransactionInput{}
        output := &lakeformation.DescribeTransactionOutput{}

        mockClient.On("DescribeTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExtendTransaction", func(t *testing.T) {
        input := &lakeformation.ExtendTransactionInput{}
        output := &lakeformation.ExtendTransactionOutput{}

        mockClient.On("ExtendTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.ExtendTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataCellsFilter", func(t *testing.T) {
        input := &lakeformation.GetDataCellsFilterInput{}
        output := &lakeformation.GetDataCellsFilterOutput{}

        mockClient.On("GetDataCellsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataCellsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataLakePrincipal", func(t *testing.T) {
        input := &lakeformation.GetDataLakePrincipalInput{}
        output := &lakeformation.GetDataLakePrincipalOutput{}

        mockClient.On("GetDataLakePrincipal", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataLakePrincipal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataLakeSettings", func(t *testing.T) {
        input := &lakeformation.GetDataLakeSettingsInput{}
        output := &lakeformation.GetDataLakeSettingsOutput{}

        mockClient.On("GetDataLakeSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataLakeSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEffectivePermissionsForPath", func(t *testing.T) {
        input := &lakeformation.GetEffectivePermissionsForPathInput{}
        output := &lakeformation.GetEffectivePermissionsForPathOutput{}

        mockClient.On("GetEffectivePermissionsForPath", ctx, input).Return(output, nil)

        result, err := mockClient.GetEffectivePermissionsForPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLFTag", func(t *testing.T) {
        input := &lakeformation.GetLFTagInput{}
        output := &lakeformation.GetLFTagOutput{}

        mockClient.On("GetLFTag", ctx, input).Return(output, nil)

        result, err := mockClient.GetLFTag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryState", func(t *testing.T) {
        input := &lakeformation.GetQueryStateInput{}
        output := &lakeformation.GetQueryStateOutput{}

        mockClient.On("GetQueryState", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryStatistics", func(t *testing.T) {
        input := &lakeformation.GetQueryStatisticsInput{}
        output := &lakeformation.GetQueryStatisticsOutput{}

        mockClient.On("GetQueryStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceLFTags", func(t *testing.T) {
        input := &lakeformation.GetResourceLFTagsInput{}
        output := &lakeformation.GetResourceLFTagsOutput{}

        mockClient.On("GetResourceLFTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceLFTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableObjects", func(t *testing.T) {
        input := &lakeformation.GetTableObjectsInput{}
        output := &lakeformation.GetTableObjectsOutput{}

        mockClient.On("GetTableObjects", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemporaryGluePartitionCredentials", func(t *testing.T) {
        input := &lakeformation.GetTemporaryGluePartitionCredentialsInput{}
        output := &lakeformation.GetTemporaryGluePartitionCredentialsOutput{}

        mockClient.On("GetTemporaryGluePartitionCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemporaryGluePartitionCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemporaryGlueTableCredentials", func(t *testing.T) {
        input := &lakeformation.GetTemporaryGlueTableCredentialsInput{}
        output := &lakeformation.GetTemporaryGlueTableCredentialsOutput{}

        mockClient.On("GetTemporaryGlueTableCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemporaryGlueTableCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkUnitResults", func(t *testing.T) {
        input := &lakeformation.GetWorkUnitResultsInput{}
        output := &lakeformation.GetWorkUnitResultsOutput{}

        mockClient.On("GetWorkUnitResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkUnitResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkUnits", func(t *testing.T) {
        input := &lakeformation.GetWorkUnitsInput{}
        output := &lakeformation.GetWorkUnitsOutput{}

        mockClient.On("GetWorkUnits", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkUnits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGrantPermissions", func(t *testing.T) {
        input := &lakeformation.GrantPermissionsInput{}
        output := &lakeformation.GrantPermissionsOutput{}

        mockClient.On("GrantPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.GrantPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataCellsFilter", func(t *testing.T) {
        input := &lakeformation.ListDataCellsFilterInput{}
        output := &lakeformation.ListDataCellsFilterOutput{}

        mockClient.On("ListDataCellsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataCellsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLFTags", func(t *testing.T) {
        input := &lakeformation.ListLFTagsInput{}
        output := &lakeformation.ListLFTagsOutput{}

        mockClient.On("ListLFTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListLFTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLakeFormationOptIns", func(t *testing.T) {
        input := &lakeformation.ListLakeFormationOptInsInput{}
        output := &lakeformation.ListLakeFormationOptInsOutput{}

        mockClient.On("ListLakeFormationOptIns", ctx, input).Return(output, nil)

        result, err := mockClient.ListLakeFormationOptIns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissions", func(t *testing.T) {
        input := &lakeformation.ListPermissionsInput{}
        output := &lakeformation.ListPermissionsOutput{}

        mockClient.On("ListPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResources", func(t *testing.T) {
        input := &lakeformation.ListResourcesInput{}
        output := &lakeformation.ListResourcesOutput{}

        mockClient.On("ListResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTableStorageOptimizers", func(t *testing.T) {
        input := &lakeformation.ListTableStorageOptimizersInput{}
        output := &lakeformation.ListTableStorageOptimizersOutput{}

        mockClient.On("ListTableStorageOptimizers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTableStorageOptimizers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTransactions", func(t *testing.T) {
        input := &lakeformation.ListTransactionsInput{}
        output := &lakeformation.ListTransactionsOutput{}

        mockClient.On("ListTransactions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTransactions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDataLakeSettings", func(t *testing.T) {
        input := &lakeformation.PutDataLakeSettingsInput{}
        output := &lakeformation.PutDataLakeSettingsOutput{}

        mockClient.On("PutDataLakeSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutDataLakeSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterResource", func(t *testing.T) {
        input := &lakeformation.RegisterResourceInput{}
        output := &lakeformation.RegisterResourceOutput{}

        mockClient.On("RegisterResource", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveLFTagsFromResource", func(t *testing.T) {
        input := &lakeformation.RemoveLFTagsFromResourceInput{}
        output := &lakeformation.RemoveLFTagsFromResourceOutput{}

        mockClient.On("RemoveLFTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveLFTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokePermissions", func(t *testing.T) {
        input := &lakeformation.RevokePermissionsInput{}
        output := &lakeformation.RevokePermissionsOutput{}

        mockClient.On("RevokePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.RevokePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchDatabasesByLFTags", func(t *testing.T) {
        input := &lakeformation.SearchDatabasesByLFTagsInput{}
        output := &lakeformation.SearchDatabasesByLFTagsOutput{}

        mockClient.On("SearchDatabasesByLFTags", ctx, input).Return(output, nil)

        result, err := mockClient.SearchDatabasesByLFTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTablesByLFTags", func(t *testing.T) {
        input := &lakeformation.SearchTablesByLFTagsInput{}
        output := &lakeformation.SearchTablesByLFTagsOutput{}

        mockClient.On("SearchTablesByLFTags", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTablesByLFTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQueryPlanning", func(t *testing.T) {
        input := &lakeformation.StartQueryPlanningInput{}
        output := &lakeformation.StartQueryPlanningOutput{}

        mockClient.On("StartQueryPlanning", ctx, input).Return(output, nil)

        result, err := mockClient.StartQueryPlanning(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTransaction", func(t *testing.T) {
        input := &lakeformation.StartTransactionInput{}
        output := &lakeformation.StartTransactionOutput{}

        mockClient.On("StartTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.StartTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataCellsFilter", func(t *testing.T) {
        input := &lakeformation.UpdateDataCellsFilterInput{}
        output := &lakeformation.UpdateDataCellsFilterOutput{}

        mockClient.On("UpdateDataCellsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataCellsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLFTag", func(t *testing.T) {
        input := &lakeformation.UpdateLFTagInput{}
        output := &lakeformation.UpdateLFTagOutput{}

        mockClient.On("UpdateLFTag", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLFTag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLakeFormationIdentityCenterConfiguration", func(t *testing.T) {
        input := &lakeformation.UpdateLakeFormationIdentityCenterConfigurationInput{}
        output := &lakeformation.UpdateLakeFormationIdentityCenterConfigurationOutput{}

        mockClient.On("UpdateLakeFormationIdentityCenterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLakeFormationIdentityCenterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResource", func(t *testing.T) {
        input := &lakeformation.UpdateResourceInput{}
        output := &lakeformation.UpdateResourceOutput{}

        mockClient.On("UpdateResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTableObjects", func(t *testing.T) {
        input := &lakeformation.UpdateTableObjectsInput{}
        output := &lakeformation.UpdateTableObjectsOutput{}

        mockClient.On("UpdateTableObjects", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTableObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTableStorageOptimizer", func(t *testing.T) {
        input := &lakeformation.UpdateTableStorageOptimizerInput{}
        output := &lakeformation.UpdateTableStorageOptimizerOutput{}

        mockClient.On("UpdateTableStorageOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTableStorageOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
