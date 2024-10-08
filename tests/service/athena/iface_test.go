// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package athena_test

// tests for the athena service interface for 
// this ../../../service/athena/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/athena/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/athena/athena_iface"
	"github.com/stretchr/testify/assert"
)

func TestAthenaServiceCanBeMocked(t *testing.T) {
	var iface athena_iface.IClient
	iface = &athena.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := athena.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetNamedQuery", func(t *testing.T) {
        input := &athena.BatchGetNamedQueryInput{}
        output := &athena.BatchGetNamedQueryOutput{}

        mockClient.On("BatchGetNamedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetNamedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetPreparedStatement", func(t *testing.T) {
        input := &athena.BatchGetPreparedStatementInput{}
        output := &athena.BatchGetPreparedStatementOutput{}

        mockClient.On("BatchGetPreparedStatement", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetPreparedStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetQueryExecution", func(t *testing.T) {
        input := &athena.BatchGetQueryExecutionInput{}
        output := &athena.BatchGetQueryExecutionOutput{}

        mockClient.On("BatchGetQueryExecution", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetQueryExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCapacityReservation", func(t *testing.T) {
        input := &athena.CancelCapacityReservationInput{}
        output := &athena.CancelCapacityReservationOutput{}

        mockClient.On("CancelCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCapacityReservation", func(t *testing.T) {
        input := &athena.CreateCapacityReservationInput{}
        output := &athena.CreateCapacityReservationOutput{}

        mockClient.On("CreateCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataCatalog", func(t *testing.T) {
        input := &athena.CreateDataCatalogInput{}
        output := &athena.CreateDataCatalogOutput{}

        mockClient.On("CreateDataCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNamedQuery", func(t *testing.T) {
        input := &athena.CreateNamedQueryInput{}
        output := &athena.CreateNamedQueryOutput{}

        mockClient.On("CreateNamedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNamedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotebook", func(t *testing.T) {
        input := &athena.CreateNotebookInput{}
        output := &athena.CreateNotebookOutput{}

        mockClient.On("CreateNotebook", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotebook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePreparedStatement", func(t *testing.T) {
        input := &athena.CreatePreparedStatementInput{}
        output := &athena.CreatePreparedStatementOutput{}

        mockClient.On("CreatePreparedStatement", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePreparedStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePresignedNotebookUrl", func(t *testing.T) {
        input := &athena.CreatePresignedNotebookUrlInput{}
        output := &athena.CreatePresignedNotebookUrlOutput{}

        mockClient.On("CreatePresignedNotebookUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePresignedNotebookUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkGroup", func(t *testing.T) {
        input := &athena.CreateWorkGroupInput{}
        output := &athena.CreateWorkGroupOutput{}

        mockClient.On("CreateWorkGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCapacityReservation", func(t *testing.T) {
        input := &athena.DeleteCapacityReservationInput{}
        output := &athena.DeleteCapacityReservationOutput{}

        mockClient.On("DeleteCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataCatalog", func(t *testing.T) {
        input := &athena.DeleteDataCatalogInput{}
        output := &athena.DeleteDataCatalogOutput{}

        mockClient.On("DeleteDataCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNamedQuery", func(t *testing.T) {
        input := &athena.DeleteNamedQueryInput{}
        output := &athena.DeleteNamedQueryOutput{}

        mockClient.On("DeleteNamedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNamedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotebook", func(t *testing.T) {
        input := &athena.DeleteNotebookInput{}
        output := &athena.DeleteNotebookOutput{}

        mockClient.On("DeleteNotebook", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotebook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePreparedStatement", func(t *testing.T) {
        input := &athena.DeletePreparedStatementInput{}
        output := &athena.DeletePreparedStatementOutput{}

        mockClient.On("DeletePreparedStatement", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePreparedStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkGroup", func(t *testing.T) {
        input := &athena.DeleteWorkGroupInput{}
        output := &athena.DeleteWorkGroupOutput{}

        mockClient.On("DeleteWorkGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportNotebook", func(t *testing.T) {
        input := &athena.ExportNotebookInput{}
        output := &athena.ExportNotebookOutput{}

        mockClient.On("ExportNotebook", ctx, input).Return(output, nil)

        result, err := mockClient.ExportNotebook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalculationExecution", func(t *testing.T) {
        input := &athena.GetCalculationExecutionInput{}
        output := &athena.GetCalculationExecutionOutput{}

        mockClient.On("GetCalculationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalculationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalculationExecutionCode", func(t *testing.T) {
        input := &athena.GetCalculationExecutionCodeInput{}
        output := &athena.GetCalculationExecutionCodeOutput{}

        mockClient.On("GetCalculationExecutionCode", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalculationExecutionCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalculationExecutionStatus", func(t *testing.T) {
        input := &athena.GetCalculationExecutionStatusInput{}
        output := &athena.GetCalculationExecutionStatusOutput{}

        mockClient.On("GetCalculationExecutionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalculationExecutionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCapacityAssignmentConfiguration", func(t *testing.T) {
        input := &athena.GetCapacityAssignmentConfigurationInput{}
        output := &athena.GetCapacityAssignmentConfigurationOutput{}

        mockClient.On("GetCapacityAssignmentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetCapacityAssignmentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCapacityReservation", func(t *testing.T) {
        input := &athena.GetCapacityReservationInput{}
        output := &athena.GetCapacityReservationOutput{}

        mockClient.On("GetCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.GetCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataCatalog", func(t *testing.T) {
        input := &athena.GetDataCatalogInput{}
        output := &athena.GetDataCatalogOutput{}

        mockClient.On("GetDataCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatabase", func(t *testing.T) {
        input := &athena.GetDatabaseInput{}
        output := &athena.GetDatabaseOutput{}

        mockClient.On("GetDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNamedQuery", func(t *testing.T) {
        input := &athena.GetNamedQueryInput{}
        output := &athena.GetNamedQueryOutput{}

        mockClient.On("GetNamedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.GetNamedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotebookMetadata", func(t *testing.T) {
        input := &athena.GetNotebookMetadataInput{}
        output := &athena.GetNotebookMetadataOutput{}

        mockClient.On("GetNotebookMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotebookMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPreparedStatement", func(t *testing.T) {
        input := &athena.GetPreparedStatementInput{}
        output := &athena.GetPreparedStatementOutput{}

        mockClient.On("GetPreparedStatement", ctx, input).Return(output, nil)

        result, err := mockClient.GetPreparedStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryExecution", func(t *testing.T) {
        input := &athena.GetQueryExecutionInput{}
        output := &athena.GetQueryExecutionOutput{}

        mockClient.On("GetQueryExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResults", func(t *testing.T) {
        input := &athena.GetQueryResultsInput{}
        output := &athena.GetQueryResultsOutput{}

        mockClient.On("GetQueryResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryRuntimeStatistics", func(t *testing.T) {
        input := &athena.GetQueryRuntimeStatisticsInput{}
        output := &athena.GetQueryRuntimeStatisticsOutput{}

        mockClient.On("GetQueryRuntimeStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryRuntimeStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &athena.GetSessionInput{}
        output := &athena.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSessionStatus", func(t *testing.T) {
        input := &athena.GetSessionStatusInput{}
        output := &athena.GetSessionStatusOutput{}

        mockClient.On("GetSessionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetSessionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableMetadata", func(t *testing.T) {
        input := &athena.GetTableMetadataInput{}
        output := &athena.GetTableMetadataOutput{}

        mockClient.On("GetTableMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkGroup", func(t *testing.T) {
        input := &athena.GetWorkGroupInput{}
        output := &athena.GetWorkGroupOutput{}

        mockClient.On("GetWorkGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportNotebook", func(t *testing.T) {
        input := &athena.ImportNotebookInput{}
        output := &athena.ImportNotebookOutput{}

        mockClient.On("ImportNotebook", ctx, input).Return(output, nil)

        result, err := mockClient.ImportNotebook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationDPUSizes", func(t *testing.T) {
        input := &athena.ListApplicationDPUSizesInput{}
        output := &athena.ListApplicationDPUSizesOutput{}

        mockClient.On("ListApplicationDPUSizes", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationDPUSizes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCalculationExecutions", func(t *testing.T) {
        input := &athena.ListCalculationExecutionsInput{}
        output := &athena.ListCalculationExecutionsOutput{}

        mockClient.On("ListCalculationExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCalculationExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCapacityReservations", func(t *testing.T) {
        input := &athena.ListCapacityReservationsInput{}
        output := &athena.ListCapacityReservationsOutput{}

        mockClient.On("ListCapacityReservations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCapacityReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataCatalogs", func(t *testing.T) {
        input := &athena.ListDataCatalogsInput{}
        output := &athena.ListDataCatalogsOutput{}

        mockClient.On("ListDataCatalogs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataCatalogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatabases", func(t *testing.T) {
        input := &athena.ListDatabasesInput{}
        output := &athena.ListDatabasesOutput{}

        mockClient.On("ListDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEngineVersions", func(t *testing.T) {
        input := &athena.ListEngineVersionsInput{}
        output := &athena.ListEngineVersionsOutput{}

        mockClient.On("ListEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExecutors", func(t *testing.T) {
        input := &athena.ListExecutorsInput{}
        output := &athena.ListExecutorsOutput{}

        mockClient.On("ListExecutors", ctx, input).Return(output, nil)

        result, err := mockClient.ListExecutors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNamedQueries", func(t *testing.T) {
        input := &athena.ListNamedQueriesInput{}
        output := &athena.ListNamedQueriesOutput{}

        mockClient.On("ListNamedQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListNamedQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotebookMetadata", func(t *testing.T) {
        input := &athena.ListNotebookMetadataInput{}
        output := &athena.ListNotebookMetadataOutput{}

        mockClient.On("ListNotebookMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotebookMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotebookSessions", func(t *testing.T) {
        input := &athena.ListNotebookSessionsInput{}
        output := &athena.ListNotebookSessionsOutput{}

        mockClient.On("ListNotebookSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotebookSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPreparedStatements", func(t *testing.T) {
        input := &athena.ListPreparedStatementsInput{}
        output := &athena.ListPreparedStatementsOutput{}

        mockClient.On("ListPreparedStatements", ctx, input).Return(output, nil)

        result, err := mockClient.ListPreparedStatements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueryExecutions", func(t *testing.T) {
        input := &athena.ListQueryExecutionsInput{}
        output := &athena.ListQueryExecutionsOutput{}

        mockClient.On("ListQueryExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueryExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessions", func(t *testing.T) {
        input := &athena.ListSessionsInput{}
        output := &athena.ListSessionsOutput{}

        mockClient.On("ListSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTableMetadata", func(t *testing.T) {
        input := &athena.ListTableMetadataInput{}
        output := &athena.ListTableMetadataOutput{}

        mockClient.On("ListTableMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListTableMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &athena.ListTagsForResourceInput{}
        output := &athena.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkGroups", func(t *testing.T) {
        input := &athena.ListWorkGroupsInput{}
        output := &athena.ListWorkGroupsOutput{}

        mockClient.On("ListWorkGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCapacityAssignmentConfiguration", func(t *testing.T) {
        input := &athena.PutCapacityAssignmentConfigurationInput{}
        output := &athena.PutCapacityAssignmentConfigurationOutput{}

        mockClient.On("PutCapacityAssignmentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutCapacityAssignmentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCalculationExecution", func(t *testing.T) {
        input := &athena.StartCalculationExecutionInput{}
        output := &athena.StartCalculationExecutionOutput{}

        mockClient.On("StartCalculationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartCalculationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQueryExecution", func(t *testing.T) {
        input := &athena.StartQueryExecutionInput{}
        output := &athena.StartQueryExecutionOutput{}

        mockClient.On("StartQueryExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartQueryExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSession", func(t *testing.T) {
        input := &athena.StartSessionInput{}
        output := &athena.StartSessionOutput{}

        mockClient.On("StartSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCalculationExecution", func(t *testing.T) {
        input := &athena.StopCalculationExecutionInput{}
        output := &athena.StopCalculationExecutionOutput{}

        mockClient.On("StopCalculationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopCalculationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQueryExecution", func(t *testing.T) {
        input := &athena.StopQueryExecutionInput{}
        output := &athena.StopQueryExecutionOutput{}

        mockClient.On("StopQueryExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopQueryExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &athena.TagResourceInput{}
        output := &athena.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateSession", func(t *testing.T) {
        input := &athena.TerminateSessionInput{}
        output := &athena.TerminateSessionOutput{}

        mockClient.On("TerminateSession", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &athena.UntagResourceInput{}
        output := &athena.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCapacityReservation", func(t *testing.T) {
        input := &athena.UpdateCapacityReservationInput{}
        output := &athena.UpdateCapacityReservationOutput{}

        mockClient.On("UpdateCapacityReservation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCapacityReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataCatalog", func(t *testing.T) {
        input := &athena.UpdateDataCatalogInput{}
        output := &athena.UpdateDataCatalogOutput{}

        mockClient.On("UpdateDataCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNamedQuery", func(t *testing.T) {
        input := &athena.UpdateNamedQueryInput{}
        output := &athena.UpdateNamedQueryOutput{}

        mockClient.On("UpdateNamedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNamedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotebook", func(t *testing.T) {
        input := &athena.UpdateNotebookInput{}
        output := &athena.UpdateNotebookOutput{}

        mockClient.On("UpdateNotebook", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotebook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotebookMetadata", func(t *testing.T) {
        input := &athena.UpdateNotebookMetadataInput{}
        output := &athena.UpdateNotebookMetadataOutput{}

        mockClient.On("UpdateNotebookMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotebookMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePreparedStatement", func(t *testing.T) {
        input := &athena.UpdatePreparedStatementInput{}
        output := &athena.UpdatePreparedStatementOutput{}

        mockClient.On("UpdatePreparedStatement", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePreparedStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkGroup", func(t *testing.T) {
        input := &athena.UpdateWorkGroupInput{}
        output := &athena.UpdateWorkGroupOutput{}

        mockClient.On("UpdateWorkGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
