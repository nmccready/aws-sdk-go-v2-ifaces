package datapipeline_test

// tests for the datapipeline service interface for this ../../../service/datapipeline/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datapipeline/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datapipeline/datapipeline_iface"
	"github.com/stretchr/testify/assert"
)

func TestDatapipelineServiceCanBeMocked(t *testing.T) {
	var iface datapipeline_iface.IClient
	iface = &datapipeline.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := datapipeline.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivatePipeline", func(t *testing.T) {
        input := &datapipeline.ActivatePipelineInput{}
        output := &datapipeline.ActivatePipelineOutput{}

        mockClient.On("ActivatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.ActivatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &datapipeline.AddTagsInput{}
        output := &datapipeline.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &datapipeline.CreatePipelineInput{}
        output := &datapipeline.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivatePipeline", func(t *testing.T) {
        input := &datapipeline.DeactivatePipelineInput{}
        output := &datapipeline.DeactivatePipelineOutput{}

        mockClient.On("DeactivatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &datapipeline.DeletePipelineInput{}
        output := &datapipeline.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeObjects", func(t *testing.T) {
        input := &datapipeline.DescribeObjectsInput{}
        output := &datapipeline.DescribeObjectsOutput{}

        mockClient.On("DescribeObjects", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipelines", func(t *testing.T) {
        input := &datapipeline.DescribePipelinesInput{}
        output := &datapipeline.DescribePipelinesOutput{}

        mockClient.On("DescribePipelines", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluateExpression", func(t *testing.T) {
        input := &datapipeline.EvaluateExpressionInput{}
        output := &datapipeline.EvaluateExpressionOutput{}

        mockClient.On("EvaluateExpression", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluateExpression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipelineDefinition", func(t *testing.T) {
        input := &datapipeline.GetPipelineDefinitionInput{}
        output := &datapipeline.GetPipelineDefinitionOutput{}

        mockClient.On("GetPipelineDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipelineDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &datapipeline.ListPipelinesInput{}
        output := &datapipeline.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPollForTask", func(t *testing.T) {
        input := &datapipeline.PollForTaskInput{}
        output := &datapipeline.PollForTaskOutput{}

        mockClient.On("PollForTask", ctx, input).Return(output, nil)

        result, err := mockClient.PollForTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPipelineDefinition", func(t *testing.T) {
        input := &datapipeline.PutPipelineDefinitionInput{}
        output := &datapipeline.PutPipelineDefinitionOutput{}

        mockClient.On("PutPipelineDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.PutPipelineDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryObjects", func(t *testing.T) {
        input := &datapipeline.QueryObjectsInput{}
        output := &datapipeline.QueryObjectsOutput{}

        mockClient.On("QueryObjects", ctx, input).Return(output, nil)

        result, err := mockClient.QueryObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &datapipeline.RemoveTagsInput{}
        output := &datapipeline.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReportTaskProgress", func(t *testing.T) {
        input := &datapipeline.ReportTaskProgressInput{}
        output := &datapipeline.ReportTaskProgressOutput{}

        mockClient.On("ReportTaskProgress", ctx, input).Return(output, nil)

        result, err := mockClient.ReportTaskProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReportTaskRunnerHeartbeat", func(t *testing.T) {
        input := &datapipeline.ReportTaskRunnerHeartbeatInput{}
        output := &datapipeline.ReportTaskRunnerHeartbeatOutput{}

        mockClient.On("ReportTaskRunnerHeartbeat", ctx, input).Return(output, nil)

        result, err := mockClient.ReportTaskRunnerHeartbeat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetStatus", func(t *testing.T) {
        input := &datapipeline.SetStatusInput{}
        output := &datapipeline.SetStatusOutput{}

        mockClient.On("SetStatus", ctx, input).Return(output, nil)

        result, err := mockClient.SetStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTaskStatus", func(t *testing.T) {
        input := &datapipeline.SetTaskStatusInput{}
        output := &datapipeline.SetTaskStatusOutput{}

        mockClient.On("SetTaskStatus", ctx, input).Return(output, nil)

        result, err := mockClient.SetTaskStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidatePipelineDefinition", func(t *testing.T) {
        input := &datapipeline.ValidatePipelineDefinitionInput{}
        output := &datapipeline.ValidatePipelineDefinitionOutput{}

        mockClient.On("ValidatePipelineDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.ValidatePipelineDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
