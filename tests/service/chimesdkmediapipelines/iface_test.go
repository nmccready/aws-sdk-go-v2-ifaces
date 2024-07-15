// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package chimesdkmediapipelines_test

// tests for the chimesdkmediapipelines service interface for 
// this ../../../service/chimesdkmediapipelines/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmediapipelines/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmediapipelines/chimesdkmediapipelines_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimesdkmediapipelinesServiceCanBeMocked(t *testing.T) {
	var iface chimesdkmediapipelines_iface.IClient
	iface = &chimesdkmediapipelines.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chimesdkmediapipelines.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaCapturePipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaCapturePipelineInput{}
        output := &chimesdkmediapipelines.CreateMediaCapturePipelineOutput{}

        mockClient.On("CreateMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaConcatenationPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaConcatenationPipelineInput{}
        output := &chimesdkmediapipelines.CreateMediaConcatenationPipelineOutput{}

        mockClient.On("CreateMediaConcatenationPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaConcatenationPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaInsightsPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaInsightsPipelineInput{}
        output := &chimesdkmediapipelines.CreateMediaInsightsPipelineOutput{}

        mockClient.On("CreateMediaInsightsPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaInsightsPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaInsightsPipelineConfiguration", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationInput{}
        output := &chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationOutput{}

        mockClient.On("CreateMediaInsightsPipelineConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaInsightsPipelineConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaLiveConnectorPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaLiveConnectorPipelineInput{}
        output := &chimesdkmediapipelines.CreateMediaLiveConnectorPipelineOutput{}

        mockClient.On("CreateMediaLiveConnectorPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaLiveConnectorPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaPipelineKinesisVideoStreamPool", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaPipelineKinesisVideoStreamPoolInput{}
        output := &chimesdkmediapipelines.CreateMediaPipelineKinesisVideoStreamPoolOutput{}

        mockClient.On("CreateMediaPipelineKinesisVideoStreamPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaPipelineKinesisVideoStreamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaStreamPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.CreateMediaStreamPipelineInput{}
        output := &chimesdkmediapipelines.CreateMediaStreamPipelineOutput{}

        mockClient.On("CreateMediaStreamPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaStreamPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaCapturePipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.DeleteMediaCapturePipelineInput{}
        output := &chimesdkmediapipelines.DeleteMediaCapturePipelineOutput{}

        mockClient.On("DeleteMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaInsightsPipelineConfiguration", func(t *testing.T) {
        input := &chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationInput{}
        output := &chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationOutput{}

        mockClient.On("DeleteMediaInsightsPipelineConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaInsightsPipelineConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.DeleteMediaPipelineInput{}
        output := &chimesdkmediapipelines.DeleteMediaPipelineOutput{}

        mockClient.On("DeleteMediaPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaPipelineKinesisVideoStreamPool", func(t *testing.T) {
        input := &chimesdkmediapipelines.DeleteMediaPipelineKinesisVideoStreamPoolInput{}
        output := &chimesdkmediapipelines.DeleteMediaPipelineKinesisVideoStreamPoolOutput{}

        mockClient.On("DeleteMediaPipelineKinesisVideoStreamPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaPipelineKinesisVideoStreamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaCapturePipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetMediaCapturePipelineInput{}
        output := &chimesdkmediapipelines.GetMediaCapturePipelineOutput{}

        mockClient.On("GetMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaInsightsPipelineConfiguration", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationInput{}
        output := &chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationOutput{}

        mockClient.On("GetMediaInsightsPipelineConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaInsightsPipelineConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaPipeline", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetMediaPipelineInput{}
        output := &chimesdkmediapipelines.GetMediaPipelineOutput{}

        mockClient.On("GetMediaPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaPipelineKinesisVideoStreamPool", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetMediaPipelineKinesisVideoStreamPoolInput{}
        output := &chimesdkmediapipelines.GetMediaPipelineKinesisVideoStreamPoolOutput{}

        mockClient.On("GetMediaPipelineKinesisVideoStreamPool", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaPipelineKinesisVideoStreamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetSpeakerSearchTaskInput{}
        output := &chimesdkmediapipelines.GetSpeakerSearchTaskOutput{}

        mockClient.On("GetSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.GetVoiceToneAnalysisTaskInput{}
        output := &chimesdkmediapipelines.GetVoiceToneAnalysisTaskOutput{}

        mockClient.On("GetVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaCapturePipelines", func(t *testing.T) {
        input := &chimesdkmediapipelines.ListMediaCapturePipelinesInput{}
        output := &chimesdkmediapipelines.ListMediaCapturePipelinesOutput{}

        mockClient.On("ListMediaCapturePipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaCapturePipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaInsightsPipelineConfigurations", func(t *testing.T) {
        input := &chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput{}
        output := &chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput{}

        mockClient.On("ListMediaInsightsPipelineConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaInsightsPipelineConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaPipelineKinesisVideoStreamPools", func(t *testing.T) {
        input := &chimesdkmediapipelines.ListMediaPipelineKinesisVideoStreamPoolsInput{}
        output := &chimesdkmediapipelines.ListMediaPipelineKinesisVideoStreamPoolsOutput{}

        mockClient.On("ListMediaPipelineKinesisVideoStreamPools", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaPipelineKinesisVideoStreamPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaPipelines", func(t *testing.T) {
        input := &chimesdkmediapipelines.ListMediaPipelinesInput{}
        output := &chimesdkmediapipelines.ListMediaPipelinesOutput{}

        mockClient.On("ListMediaPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chimesdkmediapipelines.ListTagsForResourceInput{}
        output := &chimesdkmediapipelines.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.StartSpeakerSearchTaskInput{}
        output := &chimesdkmediapipelines.StartSpeakerSearchTaskOutput{}

        mockClient.On("StartSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.StartVoiceToneAnalysisTaskInput{}
        output := &chimesdkmediapipelines.StartVoiceToneAnalysisTaskOutput{}

        mockClient.On("StartVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.StopSpeakerSearchTaskInput{}
        output := &chimesdkmediapipelines.StopSpeakerSearchTaskOutput{}

        mockClient.On("StopSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkmediapipelines.StopVoiceToneAnalysisTaskInput{}
        output := &chimesdkmediapipelines.StopVoiceToneAnalysisTaskOutput{}

        mockClient.On("StopVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chimesdkmediapipelines.TagResourceInput{}
        output := &chimesdkmediapipelines.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chimesdkmediapipelines.UntagResourceInput{}
        output := &chimesdkmediapipelines.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMediaInsightsPipelineConfiguration", func(t *testing.T) {
        input := &chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationInput{}
        output := &chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationOutput{}

        mockClient.On("UpdateMediaInsightsPipelineConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMediaInsightsPipelineConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMediaInsightsPipelineStatus", func(t *testing.T) {
        input := &chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusInput{}
        output := &chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusOutput{}

        mockClient.On("UpdateMediaInsightsPipelineStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMediaInsightsPipelineStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMediaPipelineKinesisVideoStreamPool", func(t *testing.T) {
        input := &chimesdkmediapipelines.UpdateMediaPipelineKinesisVideoStreamPoolInput{}
        output := &chimesdkmediapipelines.UpdateMediaPipelineKinesisVideoStreamPoolOutput{}

        mockClient.On("UpdateMediaPipelineKinesisVideoStreamPool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMediaPipelineKinesisVideoStreamPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
