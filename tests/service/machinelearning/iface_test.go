package machinelearning_test

// tests for the machinelearning service interface for this ../../../service/machinelearning/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/machinelearning/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/machinelearning/machinelearning_iface"
	"github.com/stretchr/testify/assert"
)

func TestMachinelearningServiceCanBeMocked(t *testing.T) {
	var iface machinelearning_iface.IClient
	iface = &machinelearning.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := machinelearning.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &machinelearning.AddTagsInput{}
        output := &machinelearning.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchPrediction", func(t *testing.T) {
        input := &machinelearning.CreateBatchPredictionInput{}
        output := &machinelearning.CreateBatchPredictionOutput{}

        mockClient.On("CreateBatchPrediction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchPrediction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSourceFromRDS", func(t *testing.T) {
        input := &machinelearning.CreateDataSourceFromRDSInput{}
        output := &machinelearning.CreateDataSourceFromRDSOutput{}

        mockClient.On("CreateDataSourceFromRDS", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSourceFromRDS(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSourceFromRedshift", func(t *testing.T) {
        input := &machinelearning.CreateDataSourceFromRedshiftInput{}
        output := &machinelearning.CreateDataSourceFromRedshiftOutput{}

        mockClient.On("CreateDataSourceFromRedshift", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSourceFromRedshift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSourceFromS3", func(t *testing.T) {
        input := &machinelearning.CreateDataSourceFromS3Input{}
        output := &machinelearning.CreateDataSourceFromS3Output{}

        mockClient.On("CreateDataSourceFromS3", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSourceFromS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEvaluation", func(t *testing.T) {
        input := &machinelearning.CreateEvaluationInput{}
        output := &machinelearning.CreateEvaluationOutput{}

        mockClient.On("CreateEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMLModel", func(t *testing.T) {
        input := &machinelearning.CreateMLModelInput{}
        output := &machinelearning.CreateMLModelOutput{}

        mockClient.On("CreateMLModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMLModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRealtimeEndpoint", func(t *testing.T) {
        input := &machinelearning.CreateRealtimeEndpointInput{}
        output := &machinelearning.CreateRealtimeEndpointOutput{}

        mockClient.On("CreateRealtimeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRealtimeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBatchPrediction", func(t *testing.T) {
        input := &machinelearning.DeleteBatchPredictionInput{}
        output := &machinelearning.DeleteBatchPredictionOutput{}

        mockClient.On("DeleteBatchPrediction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBatchPrediction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &machinelearning.DeleteDataSourceInput{}
        output := &machinelearning.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEvaluation", func(t *testing.T) {
        input := &machinelearning.DeleteEvaluationInput{}
        output := &machinelearning.DeleteEvaluationOutput{}

        mockClient.On("DeleteEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMLModel", func(t *testing.T) {
        input := &machinelearning.DeleteMLModelInput{}
        output := &machinelearning.DeleteMLModelOutput{}

        mockClient.On("DeleteMLModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMLModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRealtimeEndpoint", func(t *testing.T) {
        input := &machinelearning.DeleteRealtimeEndpointInput{}
        output := &machinelearning.DeleteRealtimeEndpointOutput{}

        mockClient.On("DeleteRealtimeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRealtimeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &machinelearning.DeleteTagsInput{}
        output := &machinelearning.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBatchPredictions", func(t *testing.T) {
        input := &machinelearning.DescribeBatchPredictionsInput{}
        output := &machinelearning.DescribeBatchPredictionsOutput{}

        mockClient.On("DescribeBatchPredictions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBatchPredictions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSources", func(t *testing.T) {
        input := &machinelearning.DescribeDataSourcesInput{}
        output := &machinelearning.DescribeDataSourcesOutput{}

        mockClient.On("DescribeDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvaluations", func(t *testing.T) {
        input := &machinelearning.DescribeEvaluationsInput{}
        output := &machinelearning.DescribeEvaluationsOutput{}

        mockClient.On("DescribeEvaluations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvaluations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMLModels", func(t *testing.T) {
        input := &machinelearning.DescribeMLModelsInput{}
        output := &machinelearning.DescribeMLModelsOutput{}

        mockClient.On("DescribeMLModels", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMLModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &machinelearning.DescribeTagsInput{}
        output := &machinelearning.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBatchPrediction", func(t *testing.T) {
        input := &machinelearning.GetBatchPredictionInput{}
        output := &machinelearning.GetBatchPredictionOutput{}

        mockClient.On("GetBatchPrediction", ctx, input).Return(output, nil)

        result, err := mockClient.GetBatchPrediction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &machinelearning.GetDataSourceInput{}
        output := &machinelearning.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvaluation", func(t *testing.T) {
        input := &machinelearning.GetEvaluationInput{}
        output := &machinelearning.GetEvaluationOutput{}

        mockClient.On("GetEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLModel", func(t *testing.T) {
        input := &machinelearning.GetMLModelInput{}
        output := &machinelearning.GetMLModelOutput{}

        mockClient.On("GetMLModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPredict", func(t *testing.T) {
        input := &machinelearning.PredictInput{}
        output := &machinelearning.PredictOutput{}

        mockClient.On("Predict", ctx, input).Return(output, nil)

        result, err := mockClient.Predict(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBatchPrediction", func(t *testing.T) {
        input := &machinelearning.UpdateBatchPredictionInput{}
        output := &machinelearning.UpdateBatchPredictionOutput{}

        mockClient.On("UpdateBatchPrediction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBatchPrediction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &machinelearning.UpdateDataSourceInput{}
        output := &machinelearning.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEvaluation", func(t *testing.T) {
        input := &machinelearning.UpdateEvaluationInput{}
        output := &machinelearning.UpdateEvaluationOutput{}

        mockClient.On("UpdateEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMLModel", func(t *testing.T) {
        input := &machinelearning.UpdateMLModelInput{}
        output := &machinelearning.UpdateMLModelOutput{}

        mockClient.On("UpdateMLModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMLModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
