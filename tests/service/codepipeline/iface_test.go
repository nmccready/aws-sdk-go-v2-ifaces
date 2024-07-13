package codepipeline_test

// tests for the codepipeline service interface for this ../../../service/codepipeline/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codepipeline/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codepipeline/codepipeline_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodepipelineServiceCanBeMocked(t *testing.T) {
	var iface codepipeline_iface.IClient
	iface = &codepipeline.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codepipeline.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcknowledgeJob", func(t *testing.T) {
        input := &codepipeline.AcknowledgeJobInput{}
        output := &codepipeline.AcknowledgeJobOutput{}

        mockClient.On("AcknowledgeJob", ctx, input).Return(output, nil)

        result, err := mockClient.AcknowledgeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcknowledgeThirdPartyJob", func(t *testing.T) {
        input := &codepipeline.AcknowledgeThirdPartyJobInput{}
        output := &codepipeline.AcknowledgeThirdPartyJobOutput{}

        mockClient.On("AcknowledgeThirdPartyJob", ctx, input).Return(output, nil)

        result, err := mockClient.AcknowledgeThirdPartyJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomActionType", func(t *testing.T) {
        input := &codepipeline.CreateCustomActionTypeInput{}
        output := &codepipeline.CreateCustomActionTypeOutput{}

        mockClient.On("CreateCustomActionType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomActionType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &codepipeline.CreatePipelineInput{}
        output := &codepipeline.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomActionType", func(t *testing.T) {
        input := &codepipeline.DeleteCustomActionTypeInput{}
        output := &codepipeline.DeleteCustomActionTypeOutput{}

        mockClient.On("DeleteCustomActionType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomActionType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &codepipeline.DeletePipelineInput{}
        output := &codepipeline.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebhook", func(t *testing.T) {
        input := &codepipeline.DeleteWebhookInput{}
        output := &codepipeline.DeleteWebhookOutput{}

        mockClient.On("DeleteWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterWebhookWithThirdParty", func(t *testing.T) {
        input := &codepipeline.DeregisterWebhookWithThirdPartyInput{}
        output := &codepipeline.DeregisterWebhookWithThirdPartyOutput{}

        mockClient.On("DeregisterWebhookWithThirdParty", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterWebhookWithThirdParty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableStageTransition", func(t *testing.T) {
        input := &codepipeline.DisableStageTransitionInput{}
        output := &codepipeline.DisableStageTransitionOutput{}

        mockClient.On("DisableStageTransition", ctx, input).Return(output, nil)

        result, err := mockClient.DisableStageTransition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableStageTransition", func(t *testing.T) {
        input := &codepipeline.EnableStageTransitionInput{}
        output := &codepipeline.EnableStageTransitionOutput{}

        mockClient.On("EnableStageTransition", ctx, input).Return(output, nil)

        result, err := mockClient.EnableStageTransition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetActionType", func(t *testing.T) {
        input := &codepipeline.GetActionTypeInput{}
        output := &codepipeline.GetActionTypeOutput{}

        mockClient.On("GetActionType", ctx, input).Return(output, nil)

        result, err := mockClient.GetActionType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobDetails", func(t *testing.T) {
        input := &codepipeline.GetJobDetailsInput{}
        output := &codepipeline.GetJobDetailsOutput{}

        mockClient.On("GetJobDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipeline", func(t *testing.T) {
        input := &codepipeline.GetPipelineInput{}
        output := &codepipeline.GetPipelineOutput{}

        mockClient.On("GetPipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipelineExecution", func(t *testing.T) {
        input := &codepipeline.GetPipelineExecutionInput{}
        output := &codepipeline.GetPipelineExecutionOutput{}

        mockClient.On("GetPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPipelineState", func(t *testing.T) {
        input := &codepipeline.GetPipelineStateInput{}
        output := &codepipeline.GetPipelineStateOutput{}

        mockClient.On("GetPipelineState", ctx, input).Return(output, nil)

        result, err := mockClient.GetPipelineState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetThirdPartyJobDetails", func(t *testing.T) {
        input := &codepipeline.GetThirdPartyJobDetailsInput{}
        output := &codepipeline.GetThirdPartyJobDetailsOutput{}

        mockClient.On("GetThirdPartyJobDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetThirdPartyJobDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActionExecutions", func(t *testing.T) {
        input := &codepipeline.ListActionExecutionsInput{}
        output := &codepipeline.ListActionExecutionsOutput{}

        mockClient.On("ListActionExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListActionExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActionTypes", func(t *testing.T) {
        input := &codepipeline.ListActionTypesInput{}
        output := &codepipeline.ListActionTypesOutput{}

        mockClient.On("ListActionTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListActionTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelineExecutions", func(t *testing.T) {
        input := &codepipeline.ListPipelineExecutionsInput{}
        output := &codepipeline.ListPipelineExecutionsOutput{}

        mockClient.On("ListPipelineExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelineExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &codepipeline.ListPipelinesInput{}
        output := &codepipeline.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codepipeline.ListTagsForResourceInput{}
        output := &codepipeline.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebhooks", func(t *testing.T) {
        input := &codepipeline.ListWebhooksInput{}
        output := &codepipeline.ListWebhooksOutput{}

        mockClient.On("ListWebhooks", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebhooks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPollForJobs", func(t *testing.T) {
        input := &codepipeline.PollForJobsInput{}
        output := &codepipeline.PollForJobsOutput{}

        mockClient.On("PollForJobs", ctx, input).Return(output, nil)

        result, err := mockClient.PollForJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPollForThirdPartyJobs", func(t *testing.T) {
        input := &codepipeline.PollForThirdPartyJobsInput{}
        output := &codepipeline.PollForThirdPartyJobsOutput{}

        mockClient.On("PollForThirdPartyJobs", ctx, input).Return(output, nil)

        result, err := mockClient.PollForThirdPartyJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutActionRevision", func(t *testing.T) {
        input := &codepipeline.PutActionRevisionInput{}
        output := &codepipeline.PutActionRevisionOutput{}

        mockClient.On("PutActionRevision", ctx, input).Return(output, nil)

        result, err := mockClient.PutActionRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApprovalResult", func(t *testing.T) {
        input := &codepipeline.PutApprovalResultInput{}
        output := &codepipeline.PutApprovalResultOutput{}

        mockClient.On("PutApprovalResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutApprovalResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutJobFailureResult", func(t *testing.T) {
        input := &codepipeline.PutJobFailureResultInput{}
        output := &codepipeline.PutJobFailureResultOutput{}

        mockClient.On("PutJobFailureResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutJobFailureResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutJobSuccessResult", func(t *testing.T) {
        input := &codepipeline.PutJobSuccessResultInput{}
        output := &codepipeline.PutJobSuccessResultOutput{}

        mockClient.On("PutJobSuccessResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutJobSuccessResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutThirdPartyJobFailureResult", func(t *testing.T) {
        input := &codepipeline.PutThirdPartyJobFailureResultInput{}
        output := &codepipeline.PutThirdPartyJobFailureResultOutput{}

        mockClient.On("PutThirdPartyJobFailureResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutThirdPartyJobFailureResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutThirdPartyJobSuccessResult", func(t *testing.T) {
        input := &codepipeline.PutThirdPartyJobSuccessResultInput{}
        output := &codepipeline.PutThirdPartyJobSuccessResultOutput{}

        mockClient.On("PutThirdPartyJobSuccessResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutThirdPartyJobSuccessResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutWebhook", func(t *testing.T) {
        input := &codepipeline.PutWebhookInput{}
        output := &codepipeline.PutWebhookOutput{}

        mockClient.On("PutWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.PutWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterWebhookWithThirdParty", func(t *testing.T) {
        input := &codepipeline.RegisterWebhookWithThirdPartyInput{}
        output := &codepipeline.RegisterWebhookWithThirdPartyOutput{}

        mockClient.On("RegisterWebhookWithThirdParty", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterWebhookWithThirdParty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryStageExecution", func(t *testing.T) {
        input := &codepipeline.RetryStageExecutionInput{}
        output := &codepipeline.RetryStageExecutionOutput{}

        mockClient.On("RetryStageExecution", ctx, input).Return(output, nil)

        result, err := mockClient.RetryStageExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRollbackStage", func(t *testing.T) {
        input := &codepipeline.RollbackStageInput{}
        output := &codepipeline.RollbackStageOutput{}

        mockClient.On("RollbackStage", ctx, input).Return(output, nil)

        result, err := mockClient.RollbackStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPipelineExecution", func(t *testing.T) {
        input := &codepipeline.StartPipelineExecutionInput{}
        output := &codepipeline.StartPipelineExecutionOutput{}

        mockClient.On("StartPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPipelineExecution", func(t *testing.T) {
        input := &codepipeline.StopPipelineExecutionInput{}
        output := &codepipeline.StopPipelineExecutionOutput{}

        mockClient.On("StopPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codepipeline.TagResourceInput{}
        output := &codepipeline.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codepipeline.UntagResourceInput{}
        output := &codepipeline.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateActionType", func(t *testing.T) {
        input := &codepipeline.UpdateActionTypeInput{}
        output := &codepipeline.UpdateActionTypeOutput{}

        mockClient.On("UpdateActionType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateActionType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipeline", func(t *testing.T) {
        input := &codepipeline.UpdatePipelineInput{}
        output := &codepipeline.UpdatePipelineOutput{}

        mockClient.On("UpdatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
