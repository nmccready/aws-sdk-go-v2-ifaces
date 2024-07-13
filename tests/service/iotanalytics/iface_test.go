package iotanalytics_test

// tests for the iotanalytics service interface for this ../../../service/iotanalytics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotanalytics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotanalytics/iotanalytics_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotanalyticsServiceCanBeMocked(t *testing.T) {
	var iface iotanalytics_iface.IClient
	iface = &iotanalytics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotanalytics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutMessage", func(t *testing.T) {
        input := &iotanalytics.BatchPutMessageInput{}
        output := &iotanalytics.BatchPutMessageOutput{}

        mockClient.On("BatchPutMessage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelPipelineReprocessing", func(t *testing.T) {
        input := &iotanalytics.CancelPipelineReprocessingInput{}
        output := &iotanalytics.CancelPipelineReprocessingOutput{}

        mockClient.On("CancelPipelineReprocessing", ctx, input).Return(output, nil)

        result, err := mockClient.CancelPipelineReprocessing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &iotanalytics.CreateChannelInput{}
        output := &iotanalytics.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &iotanalytics.CreateDatasetInput{}
        output := &iotanalytics.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetContent", func(t *testing.T) {
        input := &iotanalytics.CreateDatasetContentInput{}
        output := &iotanalytics.CreateDatasetContentOutput{}

        mockClient.On("CreateDatasetContent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatastore", func(t *testing.T) {
        input := &iotanalytics.CreateDatastoreInput{}
        output := &iotanalytics.CreateDatastoreOutput{}

        mockClient.On("CreateDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &iotanalytics.CreatePipelineInput{}
        output := &iotanalytics.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &iotanalytics.DeleteChannelInput{}
        output := &iotanalytics.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &iotanalytics.DeleteDatasetInput{}
        output := &iotanalytics.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatasetContent", func(t *testing.T) {
        input := &iotanalytics.DeleteDatasetContentInput{}
        output := &iotanalytics.DeleteDatasetContentOutput{}

        mockClient.On("DeleteDatasetContent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatasetContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatastore", func(t *testing.T) {
        input := &iotanalytics.DeleteDatastoreInput{}
        output := &iotanalytics.DeleteDatastoreOutput{}

        mockClient.On("DeleteDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &iotanalytics.DeletePipelineInput{}
        output := &iotanalytics.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &iotanalytics.DescribeChannelInput{}
        output := &iotanalytics.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &iotanalytics.DescribeDatasetInput{}
        output := &iotanalytics.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatastore", func(t *testing.T) {
        input := &iotanalytics.DescribeDatastoreInput{}
        output := &iotanalytics.DescribeDatastoreOutput{}

        mockClient.On("DescribeDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingOptions", func(t *testing.T) {
        input := &iotanalytics.DescribeLoggingOptionsInput{}
        output := &iotanalytics.DescribeLoggingOptionsOutput{}

        mockClient.On("DescribeLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipeline", func(t *testing.T) {
        input := &iotanalytics.DescribePipelineInput{}
        output := &iotanalytics.DescribePipelineOutput{}

        mockClient.On("DescribePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatasetContent", func(t *testing.T) {
        input := &iotanalytics.GetDatasetContentInput{}
        output := &iotanalytics.GetDatasetContentOutput{}

        mockClient.On("GetDatasetContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatasetContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &iotanalytics.ListChannelsInput{}
        output := &iotanalytics.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetContents", func(t *testing.T) {
        input := &iotanalytics.ListDatasetContentsInput{}
        output := &iotanalytics.ListDatasetContentsOutput{}

        mockClient.On("ListDatasetContents", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &iotanalytics.ListDatasetsInput{}
        output := &iotanalytics.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatastores", func(t *testing.T) {
        input := &iotanalytics.ListDatastoresInput{}
        output := &iotanalytics.ListDatastoresOutput{}

        mockClient.On("ListDatastores", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatastores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &iotanalytics.ListPipelinesInput{}
        output := &iotanalytics.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotanalytics.ListTagsForResourceInput{}
        output := &iotanalytics.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingOptions", func(t *testing.T) {
        input := &iotanalytics.PutLoggingOptionsInput{}
        output := &iotanalytics.PutLoggingOptionsOutput{}

        mockClient.On("PutLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunPipelineActivity", func(t *testing.T) {
        input := &iotanalytics.RunPipelineActivityInput{}
        output := &iotanalytics.RunPipelineActivityOutput{}

        mockClient.On("RunPipelineActivity", ctx, input).Return(output, nil)

        result, err := mockClient.RunPipelineActivity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSampleChannelData", func(t *testing.T) {
        input := &iotanalytics.SampleChannelDataInput{}
        output := &iotanalytics.SampleChannelDataOutput{}

        mockClient.On("SampleChannelData", ctx, input).Return(output, nil)

        result, err := mockClient.SampleChannelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPipelineReprocessing", func(t *testing.T) {
        input := &iotanalytics.StartPipelineReprocessingInput{}
        output := &iotanalytics.StartPipelineReprocessingOutput{}

        mockClient.On("StartPipelineReprocessing", ctx, input).Return(output, nil)

        result, err := mockClient.StartPipelineReprocessing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotanalytics.TagResourceInput{}
        output := &iotanalytics.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotanalytics.UntagResourceInput{}
        output := &iotanalytics.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &iotanalytics.UpdateChannelInput{}
        output := &iotanalytics.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataset", func(t *testing.T) {
        input := &iotanalytics.UpdateDatasetInput{}
        output := &iotanalytics.UpdateDatasetOutput{}

        mockClient.On("UpdateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatastore", func(t *testing.T) {
        input := &iotanalytics.UpdateDatastoreInput{}
        output := &iotanalytics.UpdateDatastoreOutput{}

        mockClient.On("UpdateDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipeline", func(t *testing.T) {
        input := &iotanalytics.UpdatePipelineInput{}
        output := &iotanalytics.UpdatePipelineOutput{}

        mockClient.On("UpdatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
