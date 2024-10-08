// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package neptunedata_test

// tests for the neptunedata service interface for 
// this ../../../service/neptunedata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/neptunedata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptunedata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptunedata/neptunedata_iface"
	"github.com/stretchr/testify/assert"
)

func TestNeptunedataServiceCanBeMocked(t *testing.T) {
	var iface neptunedata_iface.IClient
	iface = &neptunedata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := neptunedata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelGremlinQuery", func(t *testing.T) {
        input := &neptunedata.CancelGremlinQueryInput{}
        output := &neptunedata.CancelGremlinQueryOutput{}

        mockClient.On("CancelGremlinQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CancelGremlinQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelLoaderJob", func(t *testing.T) {
        input := &neptunedata.CancelLoaderJobInput{}
        output := &neptunedata.CancelLoaderJobOutput{}

        mockClient.On("CancelLoaderJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelLoaderJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMLDataProcessingJob", func(t *testing.T) {
        input := &neptunedata.CancelMLDataProcessingJobInput{}
        output := &neptunedata.CancelMLDataProcessingJobOutput{}

        mockClient.On("CancelMLDataProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMLDataProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMLModelTrainingJob", func(t *testing.T) {
        input := &neptunedata.CancelMLModelTrainingJobInput{}
        output := &neptunedata.CancelMLModelTrainingJobOutput{}

        mockClient.On("CancelMLModelTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMLModelTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMLModelTransformJob", func(t *testing.T) {
        input := &neptunedata.CancelMLModelTransformJobInput{}
        output := &neptunedata.CancelMLModelTransformJobOutput{}

        mockClient.On("CancelMLModelTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMLModelTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelOpenCypherQuery", func(t *testing.T) {
        input := &neptunedata.CancelOpenCypherQueryInput{}
        output := &neptunedata.CancelOpenCypherQueryOutput{}

        mockClient.On("CancelOpenCypherQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CancelOpenCypherQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMLEndpoint", func(t *testing.T) {
        input := &neptunedata.CreateMLEndpointInput{}
        output := &neptunedata.CreateMLEndpointOutput{}

        mockClient.On("CreateMLEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMLEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMLEndpoint", func(t *testing.T) {
        input := &neptunedata.DeleteMLEndpointInput{}
        output := &neptunedata.DeleteMLEndpointOutput{}

        mockClient.On("DeleteMLEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMLEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePropertygraphStatistics", func(t *testing.T) {
        input := &neptunedata.DeletePropertygraphStatisticsInput{}
        output := &neptunedata.DeletePropertygraphStatisticsOutput{}

        mockClient.On("DeletePropertygraphStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePropertygraphStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSparqlStatistics", func(t *testing.T) {
        input := &neptunedata.DeleteSparqlStatisticsInput{}
        output := &neptunedata.DeleteSparqlStatisticsOutput{}

        mockClient.On("DeleteSparqlStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSparqlStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteFastReset", func(t *testing.T) {
        input := &neptunedata.ExecuteFastResetInput{}
        output := &neptunedata.ExecuteFastResetOutput{}

        mockClient.On("ExecuteFastReset", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteFastReset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteGremlinExplainQuery", func(t *testing.T) {
        input := &neptunedata.ExecuteGremlinExplainQueryInput{}
        output := &neptunedata.ExecuteGremlinExplainQueryOutput{}

        mockClient.On("ExecuteGremlinExplainQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteGremlinExplainQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteGremlinProfileQuery", func(t *testing.T) {
        input := &neptunedata.ExecuteGremlinProfileQueryInput{}
        output := &neptunedata.ExecuteGremlinProfileQueryOutput{}

        mockClient.On("ExecuteGremlinProfileQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteGremlinProfileQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteGremlinQuery", func(t *testing.T) {
        input := &neptunedata.ExecuteGremlinQueryInput{}
        output := &neptunedata.ExecuteGremlinQueryOutput{}

        mockClient.On("ExecuteGremlinQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteGremlinQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteOpenCypherExplainQuery", func(t *testing.T) {
        input := &neptunedata.ExecuteOpenCypherExplainQueryInput{}
        output := &neptunedata.ExecuteOpenCypherExplainQueryOutput{}

        mockClient.On("ExecuteOpenCypherExplainQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteOpenCypherExplainQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteOpenCypherQuery", func(t *testing.T) {
        input := &neptunedata.ExecuteOpenCypherQueryInput{}
        output := &neptunedata.ExecuteOpenCypherQueryOutput{}

        mockClient.On("ExecuteOpenCypherQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteOpenCypherQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEngineStatus", func(t *testing.T) {
        input := &neptunedata.GetEngineStatusInput{}
        output := &neptunedata.GetEngineStatusOutput{}

        mockClient.On("GetEngineStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetEngineStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGremlinQueryStatus", func(t *testing.T) {
        input := &neptunedata.GetGremlinQueryStatusInput{}
        output := &neptunedata.GetGremlinQueryStatusOutput{}

        mockClient.On("GetGremlinQueryStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetGremlinQueryStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoaderJobStatus", func(t *testing.T) {
        input := &neptunedata.GetLoaderJobStatusInput{}
        output := &neptunedata.GetLoaderJobStatusOutput{}

        mockClient.On("GetLoaderJobStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoaderJobStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLDataProcessingJob", func(t *testing.T) {
        input := &neptunedata.GetMLDataProcessingJobInput{}
        output := &neptunedata.GetMLDataProcessingJobOutput{}

        mockClient.On("GetMLDataProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLDataProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLEndpoint", func(t *testing.T) {
        input := &neptunedata.GetMLEndpointInput{}
        output := &neptunedata.GetMLEndpointOutput{}

        mockClient.On("GetMLEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLModelTrainingJob", func(t *testing.T) {
        input := &neptunedata.GetMLModelTrainingJobInput{}
        output := &neptunedata.GetMLModelTrainingJobOutput{}

        mockClient.On("GetMLModelTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLModelTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLModelTransformJob", func(t *testing.T) {
        input := &neptunedata.GetMLModelTransformJobInput{}
        output := &neptunedata.GetMLModelTransformJobOutput{}

        mockClient.On("GetMLModelTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLModelTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpenCypherQueryStatus", func(t *testing.T) {
        input := &neptunedata.GetOpenCypherQueryStatusInput{}
        output := &neptunedata.GetOpenCypherQueryStatusOutput{}

        mockClient.On("GetOpenCypherQueryStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpenCypherQueryStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPropertygraphStatistics", func(t *testing.T) {
        input := &neptunedata.GetPropertygraphStatisticsInput{}
        output := &neptunedata.GetPropertygraphStatisticsOutput{}

        mockClient.On("GetPropertygraphStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetPropertygraphStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPropertygraphStream", func(t *testing.T) {
        input := &neptunedata.GetPropertygraphStreamInput{}
        output := &neptunedata.GetPropertygraphStreamOutput{}

        mockClient.On("GetPropertygraphStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetPropertygraphStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPropertygraphSummary", func(t *testing.T) {
        input := &neptunedata.GetPropertygraphSummaryInput{}
        output := &neptunedata.GetPropertygraphSummaryOutput{}

        mockClient.On("GetPropertygraphSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetPropertygraphSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRDFGraphSummary", func(t *testing.T) {
        input := &neptunedata.GetRDFGraphSummaryInput{}
        output := &neptunedata.GetRDFGraphSummaryOutput{}

        mockClient.On("GetRDFGraphSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetRDFGraphSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSparqlStatistics", func(t *testing.T) {
        input := &neptunedata.GetSparqlStatisticsInput{}
        output := &neptunedata.GetSparqlStatisticsOutput{}

        mockClient.On("GetSparqlStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetSparqlStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSparqlStream", func(t *testing.T) {
        input := &neptunedata.GetSparqlStreamInput{}
        output := &neptunedata.GetSparqlStreamOutput{}

        mockClient.On("GetSparqlStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetSparqlStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGremlinQueries", func(t *testing.T) {
        input := &neptunedata.ListGremlinQueriesInput{}
        output := &neptunedata.ListGremlinQueriesOutput{}

        mockClient.On("ListGremlinQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListGremlinQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoaderJobs", func(t *testing.T) {
        input := &neptunedata.ListLoaderJobsInput{}
        output := &neptunedata.ListLoaderJobsOutput{}

        mockClient.On("ListLoaderJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoaderJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMLDataProcessingJobs", func(t *testing.T) {
        input := &neptunedata.ListMLDataProcessingJobsInput{}
        output := &neptunedata.ListMLDataProcessingJobsOutput{}

        mockClient.On("ListMLDataProcessingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMLDataProcessingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMLEndpoints", func(t *testing.T) {
        input := &neptunedata.ListMLEndpointsInput{}
        output := &neptunedata.ListMLEndpointsOutput{}

        mockClient.On("ListMLEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListMLEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMLModelTrainingJobs", func(t *testing.T) {
        input := &neptunedata.ListMLModelTrainingJobsInput{}
        output := &neptunedata.ListMLModelTrainingJobsOutput{}

        mockClient.On("ListMLModelTrainingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMLModelTrainingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMLModelTransformJobs", func(t *testing.T) {
        input := &neptunedata.ListMLModelTransformJobsInput{}
        output := &neptunedata.ListMLModelTransformJobsOutput{}

        mockClient.On("ListMLModelTransformJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMLModelTransformJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpenCypherQueries", func(t *testing.T) {
        input := &neptunedata.ListOpenCypherQueriesInput{}
        output := &neptunedata.ListOpenCypherQueriesOutput{}

        mockClient.On("ListOpenCypherQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpenCypherQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestManagePropertygraphStatistics", func(t *testing.T) {
        input := &neptunedata.ManagePropertygraphStatisticsInput{}
        output := &neptunedata.ManagePropertygraphStatisticsOutput{}

        mockClient.On("ManagePropertygraphStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.ManagePropertygraphStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestManageSparqlStatistics", func(t *testing.T) {
        input := &neptunedata.ManageSparqlStatisticsInput{}
        output := &neptunedata.ManageSparqlStatisticsOutput{}

        mockClient.On("ManageSparqlStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.ManageSparqlStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLoaderJob", func(t *testing.T) {
        input := &neptunedata.StartLoaderJobInput{}
        output := &neptunedata.StartLoaderJobOutput{}

        mockClient.On("StartLoaderJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartLoaderJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMLDataProcessingJob", func(t *testing.T) {
        input := &neptunedata.StartMLDataProcessingJobInput{}
        output := &neptunedata.StartMLDataProcessingJobOutput{}

        mockClient.On("StartMLDataProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMLDataProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMLModelTrainingJob", func(t *testing.T) {
        input := &neptunedata.StartMLModelTrainingJobInput{}
        output := &neptunedata.StartMLModelTrainingJobOutput{}

        mockClient.On("StartMLModelTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMLModelTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMLModelTransformJob", func(t *testing.T) {
        input := &neptunedata.StartMLModelTransformJobInput{}
        output := &neptunedata.StartMLModelTransformJobOutput{}

        mockClient.On("StartMLModelTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMLModelTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
