package cleanroomsml_test

// tests for the cleanroomsml service interface for this ../../../service/cleanroomsml/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cleanroomsml/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cleanroomsml/cleanroomsml_iface"
	"github.com/stretchr/testify/assert"
)

func TestCleanroomsmlServiceCanBeMocked(t *testing.T) {
	var iface cleanroomsml_iface.IClient
	iface = &cleanroomsml.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cleanroomsml.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.CreateAudienceModelInput{}
        output := &cleanroomsml.CreateAudienceModelOutput{}

        mockClient.On("CreateAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguredAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.CreateConfiguredAudienceModelInput{}
        output := &cleanroomsml.CreateConfiguredAudienceModelOutput{}

        mockClient.On("CreateConfiguredAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguredAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrainingDataset", func(t *testing.T) {
        input := &cleanroomsml.CreateTrainingDatasetInput{}
        output := &cleanroomsml.CreateTrainingDatasetOutput{}

        mockClient.On("CreateTrainingDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrainingDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAudienceGenerationJob", func(t *testing.T) {
        input := &cleanroomsml.DeleteAudienceGenerationJobInput{}
        output := &cleanroomsml.DeleteAudienceGenerationJobOutput{}

        mockClient.On("DeleteAudienceGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAudienceGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.DeleteAudienceModelInput{}
        output := &cleanroomsml.DeleteAudienceModelOutput{}

        mockClient.On("DeleteAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.DeleteConfiguredAudienceModelInput{}
        output := &cleanroomsml.DeleteConfiguredAudienceModelOutput{}

        mockClient.On("DeleteConfiguredAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredAudienceModelPolicy", func(t *testing.T) {
        input := &cleanroomsml.DeleteConfiguredAudienceModelPolicyInput{}
        output := &cleanroomsml.DeleteConfiguredAudienceModelPolicyOutput{}

        mockClient.On("DeleteConfiguredAudienceModelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredAudienceModelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrainingDataset", func(t *testing.T) {
        input := &cleanroomsml.DeleteTrainingDatasetInput{}
        output := &cleanroomsml.DeleteTrainingDatasetOutput{}

        mockClient.On("DeleteTrainingDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrainingDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAudienceGenerationJob", func(t *testing.T) {
        input := &cleanroomsml.GetAudienceGenerationJobInput{}
        output := &cleanroomsml.GetAudienceGenerationJobOutput{}

        mockClient.On("GetAudienceGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetAudienceGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.GetAudienceModelInput{}
        output := &cleanroomsml.GetAudienceModelOutput{}

        mockClient.On("GetAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.GetConfiguredAudienceModelInput{}
        output := &cleanroomsml.GetConfiguredAudienceModelOutput{}

        mockClient.On("GetConfiguredAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredAudienceModelPolicy", func(t *testing.T) {
        input := &cleanroomsml.GetConfiguredAudienceModelPolicyInput{}
        output := &cleanroomsml.GetConfiguredAudienceModelPolicyOutput{}

        mockClient.On("GetConfiguredAudienceModelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredAudienceModelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrainingDataset", func(t *testing.T) {
        input := &cleanroomsml.GetTrainingDatasetInput{}
        output := &cleanroomsml.GetTrainingDatasetOutput{}

        mockClient.On("GetTrainingDataset", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrainingDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAudienceExportJobs", func(t *testing.T) {
        input := &cleanroomsml.ListAudienceExportJobsInput{}
        output := &cleanroomsml.ListAudienceExportJobsOutput{}

        mockClient.On("ListAudienceExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAudienceExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAudienceGenerationJobs", func(t *testing.T) {
        input := &cleanroomsml.ListAudienceGenerationJobsInput{}
        output := &cleanroomsml.ListAudienceGenerationJobsOutput{}

        mockClient.On("ListAudienceGenerationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAudienceGenerationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAudienceModels", func(t *testing.T) {
        input := &cleanroomsml.ListAudienceModelsInput{}
        output := &cleanroomsml.ListAudienceModelsOutput{}

        mockClient.On("ListAudienceModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListAudienceModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfiguredAudienceModels", func(t *testing.T) {
        input := &cleanroomsml.ListConfiguredAudienceModelsInput{}
        output := &cleanroomsml.ListConfiguredAudienceModelsOutput{}

        mockClient.On("ListConfiguredAudienceModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfiguredAudienceModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cleanroomsml.ListTagsForResourceInput{}
        output := &cleanroomsml.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrainingDatasets", func(t *testing.T) {
        input := &cleanroomsml.ListTrainingDatasetsInput{}
        output := &cleanroomsml.ListTrainingDatasetsOutput{}

        mockClient.On("ListTrainingDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrainingDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfiguredAudienceModelPolicy", func(t *testing.T) {
        input := &cleanroomsml.PutConfiguredAudienceModelPolicyInput{}
        output := &cleanroomsml.PutConfiguredAudienceModelPolicyOutput{}

        mockClient.On("PutConfiguredAudienceModelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfiguredAudienceModelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAudienceExportJob", func(t *testing.T) {
        input := &cleanroomsml.StartAudienceExportJobInput{}
        output := &cleanroomsml.StartAudienceExportJobOutput{}

        mockClient.On("StartAudienceExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartAudienceExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAudienceGenerationJob", func(t *testing.T) {
        input := &cleanroomsml.StartAudienceGenerationJobInput{}
        output := &cleanroomsml.StartAudienceGenerationJobOutput{}

        mockClient.On("StartAudienceGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartAudienceGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cleanroomsml.TagResourceInput{}
        output := &cleanroomsml.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cleanroomsml.UntagResourceInput{}
        output := &cleanroomsml.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguredAudienceModel", func(t *testing.T) {
        input := &cleanroomsml.UpdateConfiguredAudienceModelInput{}
        output := &cleanroomsml.UpdateConfiguredAudienceModelOutput{}

        mockClient.On("UpdateConfiguredAudienceModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguredAudienceModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
