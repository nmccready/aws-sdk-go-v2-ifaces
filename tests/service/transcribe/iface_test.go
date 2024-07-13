package transcribe_test

// tests for the transcribe service interface for this ../../../service/transcribe/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transcribe/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transcribe/transcribe_iface"
	"github.com/stretchr/testify/assert"
)

func TestTranscribeServiceCanBeMocked(t *testing.T) {
	var iface transcribe_iface.IClient
	iface = &transcribe.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := transcribe.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCallAnalyticsCategory", func(t *testing.T) {
        input := &transcribe.CreateCallAnalyticsCategoryInput{}
        output := &transcribe.CreateCallAnalyticsCategoryOutput{}

        mockClient.On("CreateCallAnalyticsCategory", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCallAnalyticsCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLanguageModel", func(t *testing.T) {
        input := &transcribe.CreateLanguageModelInput{}
        output := &transcribe.CreateLanguageModelOutput{}

        mockClient.On("CreateLanguageModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLanguageModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMedicalVocabulary", func(t *testing.T) {
        input := &transcribe.CreateMedicalVocabularyInput{}
        output := &transcribe.CreateMedicalVocabularyOutput{}

        mockClient.On("CreateMedicalVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMedicalVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVocabulary", func(t *testing.T) {
        input := &transcribe.CreateVocabularyInput{}
        output := &transcribe.CreateVocabularyOutput{}

        mockClient.On("CreateVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVocabularyFilter", func(t *testing.T) {
        input := &transcribe.CreateVocabularyFilterInput{}
        output := &transcribe.CreateVocabularyFilterOutput{}

        mockClient.On("CreateVocabularyFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVocabularyFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCallAnalyticsCategory", func(t *testing.T) {
        input := &transcribe.DeleteCallAnalyticsCategoryInput{}
        output := &transcribe.DeleteCallAnalyticsCategoryOutput{}

        mockClient.On("DeleteCallAnalyticsCategory", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCallAnalyticsCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCallAnalyticsJob", func(t *testing.T) {
        input := &transcribe.DeleteCallAnalyticsJobInput{}
        output := &transcribe.DeleteCallAnalyticsJobOutput{}

        mockClient.On("DeleteCallAnalyticsJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCallAnalyticsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLanguageModel", func(t *testing.T) {
        input := &transcribe.DeleteLanguageModelInput{}
        output := &transcribe.DeleteLanguageModelOutput{}

        mockClient.On("DeleteLanguageModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLanguageModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMedicalScribeJob", func(t *testing.T) {
        input := &transcribe.DeleteMedicalScribeJobInput{}
        output := &transcribe.DeleteMedicalScribeJobOutput{}

        mockClient.On("DeleteMedicalScribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMedicalScribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMedicalTranscriptionJob", func(t *testing.T) {
        input := &transcribe.DeleteMedicalTranscriptionJobInput{}
        output := &transcribe.DeleteMedicalTranscriptionJobOutput{}

        mockClient.On("DeleteMedicalTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMedicalTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMedicalVocabulary", func(t *testing.T) {
        input := &transcribe.DeleteMedicalVocabularyInput{}
        output := &transcribe.DeleteMedicalVocabularyOutput{}

        mockClient.On("DeleteMedicalVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMedicalVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTranscriptionJob", func(t *testing.T) {
        input := &transcribe.DeleteTranscriptionJobInput{}
        output := &transcribe.DeleteTranscriptionJobOutput{}

        mockClient.On("DeleteTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVocabulary", func(t *testing.T) {
        input := &transcribe.DeleteVocabularyInput{}
        output := &transcribe.DeleteVocabularyOutput{}

        mockClient.On("DeleteVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVocabularyFilter", func(t *testing.T) {
        input := &transcribe.DeleteVocabularyFilterInput{}
        output := &transcribe.DeleteVocabularyFilterOutput{}

        mockClient.On("DeleteVocabularyFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVocabularyFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLanguageModel", func(t *testing.T) {
        input := &transcribe.DescribeLanguageModelInput{}
        output := &transcribe.DescribeLanguageModelOutput{}

        mockClient.On("DescribeLanguageModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLanguageModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCallAnalyticsCategory", func(t *testing.T) {
        input := &transcribe.GetCallAnalyticsCategoryInput{}
        output := &transcribe.GetCallAnalyticsCategoryOutput{}

        mockClient.On("GetCallAnalyticsCategory", ctx, input).Return(output, nil)

        result, err := mockClient.GetCallAnalyticsCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCallAnalyticsJob", func(t *testing.T) {
        input := &transcribe.GetCallAnalyticsJobInput{}
        output := &transcribe.GetCallAnalyticsJobOutput{}

        mockClient.On("GetCallAnalyticsJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetCallAnalyticsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMedicalScribeJob", func(t *testing.T) {
        input := &transcribe.GetMedicalScribeJobInput{}
        output := &transcribe.GetMedicalScribeJobOutput{}

        mockClient.On("GetMedicalScribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMedicalScribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMedicalTranscriptionJob", func(t *testing.T) {
        input := &transcribe.GetMedicalTranscriptionJobInput{}
        output := &transcribe.GetMedicalTranscriptionJobOutput{}

        mockClient.On("GetMedicalTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMedicalTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMedicalVocabulary", func(t *testing.T) {
        input := &transcribe.GetMedicalVocabularyInput{}
        output := &transcribe.GetMedicalVocabularyOutput{}

        mockClient.On("GetMedicalVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.GetMedicalVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTranscriptionJob", func(t *testing.T) {
        input := &transcribe.GetTranscriptionJobInput{}
        output := &transcribe.GetTranscriptionJobOutput{}

        mockClient.On("GetTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVocabulary", func(t *testing.T) {
        input := &transcribe.GetVocabularyInput{}
        output := &transcribe.GetVocabularyOutput{}

        mockClient.On("GetVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.GetVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVocabularyFilter", func(t *testing.T) {
        input := &transcribe.GetVocabularyFilterInput{}
        output := &transcribe.GetVocabularyFilterOutput{}

        mockClient.On("GetVocabularyFilter", ctx, input).Return(output, nil)

        result, err := mockClient.GetVocabularyFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCallAnalyticsCategories", func(t *testing.T) {
        input := &transcribe.ListCallAnalyticsCategoriesInput{}
        output := &transcribe.ListCallAnalyticsCategoriesOutput{}

        mockClient.On("ListCallAnalyticsCategories", ctx, input).Return(output, nil)

        result, err := mockClient.ListCallAnalyticsCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCallAnalyticsJobs", func(t *testing.T) {
        input := &transcribe.ListCallAnalyticsJobsInput{}
        output := &transcribe.ListCallAnalyticsJobsOutput{}

        mockClient.On("ListCallAnalyticsJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListCallAnalyticsJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLanguageModels", func(t *testing.T) {
        input := &transcribe.ListLanguageModelsInput{}
        output := &transcribe.ListLanguageModelsOutput{}

        mockClient.On("ListLanguageModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListLanguageModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMedicalScribeJobs", func(t *testing.T) {
        input := &transcribe.ListMedicalScribeJobsInput{}
        output := &transcribe.ListMedicalScribeJobsOutput{}

        mockClient.On("ListMedicalScribeJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMedicalScribeJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMedicalTranscriptionJobs", func(t *testing.T) {
        input := &transcribe.ListMedicalTranscriptionJobsInput{}
        output := &transcribe.ListMedicalTranscriptionJobsOutput{}

        mockClient.On("ListMedicalTranscriptionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMedicalTranscriptionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMedicalVocabularies", func(t *testing.T) {
        input := &transcribe.ListMedicalVocabulariesInput{}
        output := &transcribe.ListMedicalVocabulariesOutput{}

        mockClient.On("ListMedicalVocabularies", ctx, input).Return(output, nil)

        result, err := mockClient.ListMedicalVocabularies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &transcribe.ListTagsForResourceInput{}
        output := &transcribe.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTranscriptionJobs", func(t *testing.T) {
        input := &transcribe.ListTranscriptionJobsInput{}
        output := &transcribe.ListTranscriptionJobsOutput{}

        mockClient.On("ListTranscriptionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTranscriptionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVocabularies", func(t *testing.T) {
        input := &transcribe.ListVocabulariesInput{}
        output := &transcribe.ListVocabulariesOutput{}

        mockClient.On("ListVocabularies", ctx, input).Return(output, nil)

        result, err := mockClient.ListVocabularies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVocabularyFilters", func(t *testing.T) {
        input := &transcribe.ListVocabularyFiltersInput{}
        output := &transcribe.ListVocabularyFiltersOutput{}

        mockClient.On("ListVocabularyFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListVocabularyFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCallAnalyticsJob", func(t *testing.T) {
        input := &transcribe.StartCallAnalyticsJobInput{}
        output := &transcribe.StartCallAnalyticsJobOutput{}

        mockClient.On("StartCallAnalyticsJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartCallAnalyticsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMedicalScribeJob", func(t *testing.T) {
        input := &transcribe.StartMedicalScribeJobInput{}
        output := &transcribe.StartMedicalScribeJobOutput{}

        mockClient.On("StartMedicalScribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMedicalScribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMedicalTranscriptionJob", func(t *testing.T) {
        input := &transcribe.StartMedicalTranscriptionJobInput{}
        output := &transcribe.StartMedicalTranscriptionJobOutput{}

        mockClient.On("StartMedicalTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMedicalTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTranscriptionJob", func(t *testing.T) {
        input := &transcribe.StartTranscriptionJobInput{}
        output := &transcribe.StartTranscriptionJobOutput{}

        mockClient.On("StartTranscriptionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartTranscriptionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &transcribe.TagResourceInput{}
        output := &transcribe.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &transcribe.UntagResourceInput{}
        output := &transcribe.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCallAnalyticsCategory", func(t *testing.T) {
        input := &transcribe.UpdateCallAnalyticsCategoryInput{}
        output := &transcribe.UpdateCallAnalyticsCategoryOutput{}

        mockClient.On("UpdateCallAnalyticsCategory", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCallAnalyticsCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMedicalVocabulary", func(t *testing.T) {
        input := &transcribe.UpdateMedicalVocabularyInput{}
        output := &transcribe.UpdateMedicalVocabularyOutput{}

        mockClient.On("UpdateMedicalVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMedicalVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVocabulary", func(t *testing.T) {
        input := &transcribe.UpdateVocabularyInput{}
        output := &transcribe.UpdateVocabularyOutput{}

        mockClient.On("UpdateVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVocabularyFilter", func(t *testing.T) {
        input := &transcribe.UpdateVocabularyFilterInput{}
        output := &transcribe.UpdateVocabularyFilterOutput{}

        mockClient.On("UpdateVocabularyFilter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVocabularyFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
