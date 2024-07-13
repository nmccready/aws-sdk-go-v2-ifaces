package elastictranscoder_test

// tests for the elastictranscoder service interface for this ../../../service/elastictranscoder/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elastictranscoder/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elastictranscoder/elastictranscoder_iface"
	"github.com/stretchr/testify/assert"
)

func TestElastictranscoderServiceCanBeMocked(t *testing.T) {
	var iface elastictranscoder_iface.IClient
	iface = &elastictranscoder.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elastictranscoder.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &elastictranscoder.CancelJobInput{}
        output := &elastictranscoder.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &elastictranscoder.CreateJobInput{}
        output := &elastictranscoder.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &elastictranscoder.CreatePipelineInput{}
        output := &elastictranscoder.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePreset", func(t *testing.T) {
        input := &elastictranscoder.CreatePresetInput{}
        output := &elastictranscoder.CreatePresetOutput{}

        mockClient.On("CreatePreset", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &elastictranscoder.DeletePipelineInput{}
        output := &elastictranscoder.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePreset", func(t *testing.T) {
        input := &elastictranscoder.DeletePresetInput{}
        output := &elastictranscoder.DeletePresetOutput{}

        mockClient.On("DeletePreset", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobsByPipeline", func(t *testing.T) {
        input := &elastictranscoder.ListJobsByPipelineInput{}
        output := &elastictranscoder.ListJobsByPipelineOutput{}

        mockClient.On("ListJobsByPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobsByPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobsByStatus", func(t *testing.T) {
        input := &elastictranscoder.ListJobsByStatusInput{}
        output := &elastictranscoder.ListJobsByStatusOutput{}

        mockClient.On("ListJobsByStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobsByStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &elastictranscoder.ListPipelinesInput{}
        output := &elastictranscoder.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPresets", func(t *testing.T) {
        input := &elastictranscoder.ListPresetsInput{}
        output := &elastictranscoder.ListPresetsOutput{}

        mockClient.On("ListPresets", ctx, input).Return(output, nil)

        result, err := mockClient.ListPresets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReadJob", func(t *testing.T) {
        input := &elastictranscoder.ReadJobInput{}
        output := &elastictranscoder.ReadJobOutput{}

        mockClient.On("ReadJob", ctx, input).Return(output, nil)

        result, err := mockClient.ReadJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReadPipeline", func(t *testing.T) {
        input := &elastictranscoder.ReadPipelineInput{}
        output := &elastictranscoder.ReadPipelineOutput{}

        mockClient.On("ReadPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.ReadPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReadPreset", func(t *testing.T) {
        input := &elastictranscoder.ReadPresetInput{}
        output := &elastictranscoder.ReadPresetOutput{}

        mockClient.On("ReadPreset", ctx, input).Return(output, nil)

        result, err := mockClient.ReadPreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestRole", func(t *testing.T) {
        input := &elastictranscoder.TestRoleInput{}
        output := &elastictranscoder.TestRoleOutput{}

        mockClient.On("TestRole", ctx, input).Return(output, nil)

        result, err := mockClient.TestRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipeline", func(t *testing.T) {
        input := &elastictranscoder.UpdatePipelineInput{}
        output := &elastictranscoder.UpdatePipelineOutput{}

        mockClient.On("UpdatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipelineNotifications", func(t *testing.T) {
        input := &elastictranscoder.UpdatePipelineNotificationsInput{}
        output := &elastictranscoder.UpdatePipelineNotificationsOutput{}

        mockClient.On("UpdatePipelineNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipelineNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipelineStatus", func(t *testing.T) {
        input := &elastictranscoder.UpdatePipelineStatusInput{}
        output := &elastictranscoder.UpdatePipelineStatusOutput{}

        mockClient.On("UpdatePipelineStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipelineStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
