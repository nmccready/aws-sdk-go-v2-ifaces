package bedrock_test

// tests for the bedrock service interface for this ../../../service/bedrock/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrock/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrock/bedrock_iface"
	"github.com/stretchr/testify/assert"
)

func TestBedrockServiceCanBeMocked(t *testing.T) {
	var iface bedrock_iface.IClient
	iface = &bedrock.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bedrock.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEvaluationJob", func(t *testing.T) {
        input := &bedrock.CreateEvaluationJobInput{}
        output := &bedrock.CreateEvaluationJobOutput{}

        mockClient.On("CreateEvaluationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEvaluationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGuardrail", func(t *testing.T) {
        input := &bedrock.CreateGuardrailInput{}
        output := &bedrock.CreateGuardrailOutput{}

        mockClient.On("CreateGuardrail", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGuardrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGuardrailVersion", func(t *testing.T) {
        input := &bedrock.CreateGuardrailVersionInput{}
        output := &bedrock.CreateGuardrailVersionOutput{}

        mockClient.On("CreateGuardrailVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGuardrailVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelCustomizationJob", func(t *testing.T) {
        input := &bedrock.CreateModelCustomizationJobInput{}
        output := &bedrock.CreateModelCustomizationJobOutput{}

        mockClient.On("CreateModelCustomizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelCustomizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisionedModelThroughput", func(t *testing.T) {
        input := &bedrock.CreateProvisionedModelThroughputInput{}
        output := &bedrock.CreateProvisionedModelThroughputOutput{}

        mockClient.On("CreateProvisionedModelThroughput", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisionedModelThroughput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomModel", func(t *testing.T) {
        input := &bedrock.DeleteCustomModelInput{}
        output := &bedrock.DeleteCustomModelOutput{}

        mockClient.On("DeleteCustomModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGuardrail", func(t *testing.T) {
        input := &bedrock.DeleteGuardrailInput{}
        output := &bedrock.DeleteGuardrailOutput{}

        mockClient.On("DeleteGuardrail", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGuardrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelInvocationLoggingConfiguration", func(t *testing.T) {
        input := &bedrock.DeleteModelInvocationLoggingConfigurationInput{}
        output := &bedrock.DeleteModelInvocationLoggingConfigurationOutput{}

        mockClient.On("DeleteModelInvocationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelInvocationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisionedModelThroughput", func(t *testing.T) {
        input := &bedrock.DeleteProvisionedModelThroughputInput{}
        output := &bedrock.DeleteProvisionedModelThroughputOutput{}

        mockClient.On("DeleteProvisionedModelThroughput", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisionedModelThroughput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomModel", func(t *testing.T) {
        input := &bedrock.GetCustomModelInput{}
        output := &bedrock.GetCustomModelOutput{}

        mockClient.On("GetCustomModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvaluationJob", func(t *testing.T) {
        input := &bedrock.GetEvaluationJobInput{}
        output := &bedrock.GetEvaluationJobOutput{}

        mockClient.On("GetEvaluationJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvaluationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFoundationModel", func(t *testing.T) {
        input := &bedrock.GetFoundationModelInput{}
        output := &bedrock.GetFoundationModelOutput{}

        mockClient.On("GetFoundationModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetFoundationModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGuardrail", func(t *testing.T) {
        input := &bedrock.GetGuardrailInput{}
        output := &bedrock.GetGuardrailOutput{}

        mockClient.On("GetGuardrail", ctx, input).Return(output, nil)

        result, err := mockClient.GetGuardrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelCustomizationJob", func(t *testing.T) {
        input := &bedrock.GetModelCustomizationJobInput{}
        output := &bedrock.GetModelCustomizationJobOutput{}

        mockClient.On("GetModelCustomizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelCustomizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelInvocationLoggingConfiguration", func(t *testing.T) {
        input := &bedrock.GetModelInvocationLoggingConfigurationInput{}
        output := &bedrock.GetModelInvocationLoggingConfigurationOutput{}

        mockClient.On("GetModelInvocationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelInvocationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProvisionedModelThroughput", func(t *testing.T) {
        input := &bedrock.GetProvisionedModelThroughputInput{}
        output := &bedrock.GetProvisionedModelThroughputOutput{}

        mockClient.On("GetProvisionedModelThroughput", ctx, input).Return(output, nil)

        result, err := mockClient.GetProvisionedModelThroughput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomModels", func(t *testing.T) {
        input := &bedrock.ListCustomModelsInput{}
        output := &bedrock.ListCustomModelsOutput{}

        mockClient.On("ListCustomModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEvaluationJobs", func(t *testing.T) {
        input := &bedrock.ListEvaluationJobsInput{}
        output := &bedrock.ListEvaluationJobsOutput{}

        mockClient.On("ListEvaluationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEvaluationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFoundationModels", func(t *testing.T) {
        input := &bedrock.ListFoundationModelsInput{}
        output := &bedrock.ListFoundationModelsOutput{}

        mockClient.On("ListFoundationModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListFoundationModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGuardrails", func(t *testing.T) {
        input := &bedrock.ListGuardrailsInput{}
        output := &bedrock.ListGuardrailsOutput{}

        mockClient.On("ListGuardrails", ctx, input).Return(output, nil)

        result, err := mockClient.ListGuardrails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelCustomizationJobs", func(t *testing.T) {
        input := &bedrock.ListModelCustomizationJobsInput{}
        output := &bedrock.ListModelCustomizationJobsOutput{}

        mockClient.On("ListModelCustomizationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelCustomizationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisionedModelThroughputs", func(t *testing.T) {
        input := &bedrock.ListProvisionedModelThroughputsInput{}
        output := &bedrock.ListProvisionedModelThroughputsOutput{}

        mockClient.On("ListProvisionedModelThroughputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisionedModelThroughputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &bedrock.ListTagsForResourceInput{}
        output := &bedrock.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutModelInvocationLoggingConfiguration", func(t *testing.T) {
        input := &bedrock.PutModelInvocationLoggingConfigurationInput{}
        output := &bedrock.PutModelInvocationLoggingConfigurationOutput{}

        mockClient.On("PutModelInvocationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutModelInvocationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEvaluationJob", func(t *testing.T) {
        input := &bedrock.StopEvaluationJobInput{}
        output := &bedrock.StopEvaluationJobOutput{}

        mockClient.On("StopEvaluationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopEvaluationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopModelCustomizationJob", func(t *testing.T) {
        input := &bedrock.StopModelCustomizationJobInput{}
        output := &bedrock.StopModelCustomizationJobOutput{}

        mockClient.On("StopModelCustomizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopModelCustomizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &bedrock.TagResourceInput{}
        output := &bedrock.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &bedrock.UntagResourceInput{}
        output := &bedrock.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGuardrail", func(t *testing.T) {
        input := &bedrock.UpdateGuardrailInput{}
        output := &bedrock.UpdateGuardrailOutput{}

        mockClient.On("UpdateGuardrail", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGuardrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProvisionedModelThroughput", func(t *testing.T) {
        input := &bedrock.UpdateProvisionedModelThroughputInput{}
        output := &bedrock.UpdateProvisionedModelThroughputOutput{}

        mockClient.On("UpdateProvisionedModelThroughput", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProvisionedModelThroughput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
