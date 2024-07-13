package osis_test

// tests for the osis service interface for this ../../../service/osis/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/osis"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/osis/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/osis/osis_iface"
	"github.com/stretchr/testify/assert"
)

func TestOsisServiceCanBeMocked(t *testing.T) {
	var iface osis_iface.IClient
	iface = &osis.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := osis.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &osis.CreatePipelineInput{}
        output := &osis.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &osis.DeletePipelineInput{}
        output := &osis.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipeline", func(t *testing.T) {
        input := &osis.GetPipelineInput{}
        output := &osis.GetPipelineOutput{}

        mockClient.On("GetPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipelineBlueprint", func(t *testing.T) {
        input := &osis.GetPipelineBlueprintInput{}
        output := &osis.GetPipelineBlueprintOutput{}

        mockClient.On("GetPipelineBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipelineBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipelineChangeProgress", func(t *testing.T) {
        input := &osis.GetPipelineChangeProgressInput{}
        output := &osis.GetPipelineChangeProgressOutput{}

        mockClient.On("GetPipelineChangeProgress", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipelineChangeProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelineBlueprints", func(t *testing.T) {
        input := &osis.ListPipelineBlueprintsInput{}
        output := &osis.ListPipelineBlueprintsOutput{}

        mockClient.On("ListPipelineBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelineBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &osis.ListPipelinesInput{}
        output := &osis.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &osis.ListTagsForResourceInput{}
        output := &osis.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPipeline", func(t *testing.T) {
        input := &osis.StartPipelineInput{}
        output := &osis.StartPipelineOutput{}

        mockClient.On("StartPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.StartPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPipeline", func(t *testing.T) {
        input := &osis.StopPipelineInput{}
        output := &osis.StopPipelineOutput{}

        mockClient.On("StopPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.StopPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &osis.TagResourceInput{}
        output := &osis.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &osis.UntagResourceInput{}
        output := &osis.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipeline", func(t *testing.T) {
        input := &osis.UpdatePipelineInput{}
        output := &osis.UpdatePipelineOutput{}

        mockClient.On("UpdatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidatePipeline", func(t *testing.T) {
        input := &osis.ValidatePipelineInput{}
        output := &osis.ValidatePipelineOutput{}

        mockClient.On("ValidatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.ValidatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
