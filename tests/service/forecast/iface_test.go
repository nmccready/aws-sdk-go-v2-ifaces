package forecast_test

// tests for the forecast service interface for this ../../../service/forecast/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/forecast"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/forecast/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/forecast/forecast_iface"
	"github.com/stretchr/testify/assert"
)

func TestForecastServiceCanBeMocked(t *testing.T) {
	var iface forecast_iface.IClient
	iface = &forecast.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := forecast.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutoPredictor", func(t *testing.T) {
        input := &forecast.CreateAutoPredictorInput{}
        output := &forecast.CreateAutoPredictorOutput{}

        mockClient.On("CreateAutoPredictor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutoPredictor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &forecast.CreateDatasetInput{}
        output := &forecast.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetGroup", func(t *testing.T) {
        input := &forecast.CreateDatasetGroupInput{}
        output := &forecast.CreateDatasetGroupOutput{}

        mockClient.On("CreateDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetImportJob", func(t *testing.T) {
        input := &forecast.CreateDatasetImportJobInput{}
        output := &forecast.CreateDatasetImportJobOutput{}

        mockClient.On("CreateDatasetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExplainability", func(t *testing.T) {
        input := &forecast.CreateExplainabilityInput{}
        output := &forecast.CreateExplainabilityOutput{}

        mockClient.On("CreateExplainability", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExplainability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExplainabilityExport", func(t *testing.T) {
        input := &forecast.CreateExplainabilityExportInput{}
        output := &forecast.CreateExplainabilityExportOutput{}

        mockClient.On("CreateExplainabilityExport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExplainabilityExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateForecast", func(t *testing.T) {
        input := &forecast.CreateForecastInput{}
        output := &forecast.CreateForecastOutput{}

        mockClient.On("CreateForecast", ctx, input).Return(output, nil)

        result, err := mockClient.CreateForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateForecastExportJob", func(t *testing.T) {
        input := &forecast.CreateForecastExportJobInput{}
        output := &forecast.CreateForecastExportJobOutput{}

        mockClient.On("CreateForecastExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateForecastExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitor", func(t *testing.T) {
        input := &forecast.CreateMonitorInput{}
        output := &forecast.CreateMonitorOutput{}

        mockClient.On("CreateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePredictor", func(t *testing.T) {
        input := &forecast.CreatePredictorInput{}
        output := &forecast.CreatePredictorOutput{}

        mockClient.On("CreatePredictor", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePredictor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePredictorBacktestExportJob", func(t *testing.T) {
        input := &forecast.CreatePredictorBacktestExportJobInput{}
        output := &forecast.CreatePredictorBacktestExportJobOutput{}

        mockClient.On("CreatePredictorBacktestExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePredictorBacktestExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWhatIfAnalysis", func(t *testing.T) {
        input := &forecast.CreateWhatIfAnalysisInput{}
        output := &forecast.CreateWhatIfAnalysisOutput{}

        mockClient.On("CreateWhatIfAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWhatIfAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWhatIfForecast", func(t *testing.T) {
        input := &forecast.CreateWhatIfForecastInput{}
        output := &forecast.CreateWhatIfForecastOutput{}

        mockClient.On("CreateWhatIfForecast", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWhatIfForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWhatIfForecastExport", func(t *testing.T) {
        input := &forecast.CreateWhatIfForecastExportInput{}
        output := &forecast.CreateWhatIfForecastExportOutput{}

        mockClient.On("CreateWhatIfForecastExport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWhatIfForecastExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &forecast.DeleteDatasetInput{}
        output := &forecast.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatasetGroup", func(t *testing.T) {
        input := &forecast.DeleteDatasetGroupInput{}
        output := &forecast.DeleteDatasetGroupOutput{}

        mockClient.On("DeleteDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatasetImportJob", func(t *testing.T) {
        input := &forecast.DeleteDatasetImportJobInput{}
        output := &forecast.DeleteDatasetImportJobOutput{}

        mockClient.On("DeleteDatasetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatasetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExplainability", func(t *testing.T) {
        input := &forecast.DeleteExplainabilityInput{}
        output := &forecast.DeleteExplainabilityOutput{}

        mockClient.On("DeleteExplainability", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExplainability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExplainabilityExport", func(t *testing.T) {
        input := &forecast.DeleteExplainabilityExportInput{}
        output := &forecast.DeleteExplainabilityExportOutput{}

        mockClient.On("DeleteExplainabilityExport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExplainabilityExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteForecast", func(t *testing.T) {
        input := &forecast.DeleteForecastInput{}
        output := &forecast.DeleteForecastOutput{}

        mockClient.On("DeleteForecast", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteForecastExportJob", func(t *testing.T) {
        input := &forecast.DeleteForecastExportJobInput{}
        output := &forecast.DeleteForecastExportJobOutput{}

        mockClient.On("DeleteForecastExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteForecastExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitor", func(t *testing.T) {
        input := &forecast.DeleteMonitorInput{}
        output := &forecast.DeleteMonitorOutput{}

        mockClient.On("DeleteMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePredictor", func(t *testing.T) {
        input := &forecast.DeletePredictorInput{}
        output := &forecast.DeletePredictorOutput{}

        mockClient.On("DeletePredictor", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePredictor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePredictorBacktestExportJob", func(t *testing.T) {
        input := &forecast.DeletePredictorBacktestExportJobInput{}
        output := &forecast.DeletePredictorBacktestExportJobOutput{}

        mockClient.On("DeletePredictorBacktestExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePredictorBacktestExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceTree", func(t *testing.T) {
        input := &forecast.DeleteResourceTreeInput{}
        output := &forecast.DeleteResourceTreeOutput{}

        mockClient.On("DeleteResourceTree", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceTree(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWhatIfAnalysis", func(t *testing.T) {
        input := &forecast.DeleteWhatIfAnalysisInput{}
        output := &forecast.DeleteWhatIfAnalysisOutput{}

        mockClient.On("DeleteWhatIfAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWhatIfAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWhatIfForecast", func(t *testing.T) {
        input := &forecast.DeleteWhatIfForecastInput{}
        output := &forecast.DeleteWhatIfForecastOutput{}

        mockClient.On("DeleteWhatIfForecast", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWhatIfForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWhatIfForecastExport", func(t *testing.T) {
        input := &forecast.DeleteWhatIfForecastExportInput{}
        output := &forecast.DeleteWhatIfForecastExportOutput{}

        mockClient.On("DeleteWhatIfForecastExport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWhatIfForecastExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoPredictor", func(t *testing.T) {
        input := &forecast.DescribeAutoPredictorInput{}
        output := &forecast.DescribeAutoPredictorOutput{}

        mockClient.On("DescribeAutoPredictor", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoPredictor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &forecast.DescribeDatasetInput{}
        output := &forecast.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatasetGroup", func(t *testing.T) {
        input := &forecast.DescribeDatasetGroupInput{}
        output := &forecast.DescribeDatasetGroupOutput{}

        mockClient.On("DescribeDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatasetImportJob", func(t *testing.T) {
        input := &forecast.DescribeDatasetImportJobInput{}
        output := &forecast.DescribeDatasetImportJobOutput{}

        mockClient.On("DescribeDatasetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatasetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExplainability", func(t *testing.T) {
        input := &forecast.DescribeExplainabilityInput{}
        output := &forecast.DescribeExplainabilityOutput{}

        mockClient.On("DescribeExplainability", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExplainability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExplainabilityExport", func(t *testing.T) {
        input := &forecast.DescribeExplainabilityExportInput{}
        output := &forecast.DescribeExplainabilityExportOutput{}

        mockClient.On("DescribeExplainabilityExport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExplainabilityExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeForecast", func(t *testing.T) {
        input := &forecast.DescribeForecastInput{}
        output := &forecast.DescribeForecastOutput{}

        mockClient.On("DescribeForecast", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeForecastExportJob", func(t *testing.T) {
        input := &forecast.DescribeForecastExportJobInput{}
        output := &forecast.DescribeForecastExportJobOutput{}

        mockClient.On("DescribeForecastExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeForecastExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMonitor", func(t *testing.T) {
        input := &forecast.DescribeMonitorInput{}
        output := &forecast.DescribeMonitorOutput{}

        mockClient.On("DescribeMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePredictor", func(t *testing.T) {
        input := &forecast.DescribePredictorInput{}
        output := &forecast.DescribePredictorOutput{}

        mockClient.On("DescribePredictor", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePredictor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePredictorBacktestExportJob", func(t *testing.T) {
        input := &forecast.DescribePredictorBacktestExportJobInput{}
        output := &forecast.DescribePredictorBacktestExportJobOutput{}

        mockClient.On("DescribePredictorBacktestExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePredictorBacktestExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWhatIfAnalysis", func(t *testing.T) {
        input := &forecast.DescribeWhatIfAnalysisInput{}
        output := &forecast.DescribeWhatIfAnalysisOutput{}

        mockClient.On("DescribeWhatIfAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWhatIfAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWhatIfForecast", func(t *testing.T) {
        input := &forecast.DescribeWhatIfForecastInput{}
        output := &forecast.DescribeWhatIfForecastOutput{}

        mockClient.On("DescribeWhatIfForecast", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWhatIfForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWhatIfForecastExport", func(t *testing.T) {
        input := &forecast.DescribeWhatIfForecastExportInput{}
        output := &forecast.DescribeWhatIfForecastExportOutput{}

        mockClient.On("DescribeWhatIfForecastExport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWhatIfForecastExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccuracyMetrics", func(t *testing.T) {
        input := &forecast.GetAccuracyMetricsInput{}
        output := &forecast.GetAccuracyMetricsOutput{}

        mockClient.On("GetAccuracyMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccuracyMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetGroups", func(t *testing.T) {
        input := &forecast.ListDatasetGroupsInput{}
        output := &forecast.ListDatasetGroupsOutput{}

        mockClient.On("ListDatasetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetImportJobs", func(t *testing.T) {
        input := &forecast.ListDatasetImportJobsInput{}
        output := &forecast.ListDatasetImportJobsOutput{}

        mockClient.On("ListDatasetImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &forecast.ListDatasetsInput{}
        output := &forecast.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExplainabilities", func(t *testing.T) {
        input := &forecast.ListExplainabilitiesInput{}
        output := &forecast.ListExplainabilitiesOutput{}

        mockClient.On("ListExplainabilities", ctx, input).Return(output, nil)

        result, err := mockClient.ListExplainabilities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExplainabilityExports", func(t *testing.T) {
        input := &forecast.ListExplainabilityExportsInput{}
        output := &forecast.ListExplainabilityExportsOutput{}

        mockClient.On("ListExplainabilityExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExplainabilityExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListForecastExportJobs", func(t *testing.T) {
        input := &forecast.ListForecastExportJobsInput{}
        output := &forecast.ListForecastExportJobsOutput{}

        mockClient.On("ListForecastExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListForecastExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListForecasts", func(t *testing.T) {
        input := &forecast.ListForecastsInput{}
        output := &forecast.ListForecastsOutput{}

        mockClient.On("ListForecasts", ctx, input).Return(output, nil)

        result, err := mockClient.ListForecasts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitorEvaluations", func(t *testing.T) {
        input := &forecast.ListMonitorEvaluationsInput{}
        output := &forecast.ListMonitorEvaluationsOutput{}

        mockClient.On("ListMonitorEvaluations", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitorEvaluations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitors", func(t *testing.T) {
        input := &forecast.ListMonitorsInput{}
        output := &forecast.ListMonitorsOutput{}

        mockClient.On("ListMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPredictorBacktestExportJobs", func(t *testing.T) {
        input := &forecast.ListPredictorBacktestExportJobsInput{}
        output := &forecast.ListPredictorBacktestExportJobsOutput{}

        mockClient.On("ListPredictorBacktestExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListPredictorBacktestExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPredictors", func(t *testing.T) {
        input := &forecast.ListPredictorsInput{}
        output := &forecast.ListPredictorsOutput{}

        mockClient.On("ListPredictors", ctx, input).Return(output, nil)

        result, err := mockClient.ListPredictors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &forecast.ListTagsForResourceInput{}
        output := &forecast.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWhatIfAnalyses", func(t *testing.T) {
        input := &forecast.ListWhatIfAnalysesInput{}
        output := &forecast.ListWhatIfAnalysesOutput{}

        mockClient.On("ListWhatIfAnalyses", ctx, input).Return(output, nil)

        result, err := mockClient.ListWhatIfAnalyses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWhatIfForecastExports", func(t *testing.T) {
        input := &forecast.ListWhatIfForecastExportsInput{}
        output := &forecast.ListWhatIfForecastExportsOutput{}

        mockClient.On("ListWhatIfForecastExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListWhatIfForecastExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWhatIfForecasts", func(t *testing.T) {
        input := &forecast.ListWhatIfForecastsInput{}
        output := &forecast.ListWhatIfForecastsOutput{}

        mockClient.On("ListWhatIfForecasts", ctx, input).Return(output, nil)

        result, err := mockClient.ListWhatIfForecasts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeResource", func(t *testing.T) {
        input := &forecast.ResumeResourceInput{}
        output := &forecast.ResumeResourceOutput{}

        mockClient.On("ResumeResource", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopResource", func(t *testing.T) {
        input := &forecast.StopResourceInput{}
        output := &forecast.StopResourceOutput{}

        mockClient.On("StopResource", ctx, input).Return(output, nil)

        result, err := mockClient.StopResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &forecast.TagResourceInput{}
        output := &forecast.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &forecast.UntagResourceInput{}
        output := &forecast.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatasetGroup", func(t *testing.T) {
        input := &forecast.UpdateDatasetGroupInput{}
        output := &forecast.UpdateDatasetGroupOutput{}

        mockClient.On("UpdateDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
