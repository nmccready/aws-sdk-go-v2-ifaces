package translate_test

// tests for the translate service interface for this ../../../service/translate/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/translate/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/translate/translate_iface"
	"github.com/stretchr/testify/assert"
)

func TestTranslateServiceCanBeMocked(t *testing.T) {
	var iface translate_iface.IClient
	iface = &translate.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := translate.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParallelData", func(t *testing.T) {
        input := &translate.CreateParallelDataInput{}
        output := &translate.CreateParallelDataOutput{}

        mockClient.On("CreateParallelData", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParallelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteParallelData", func(t *testing.T) {
        input := &translate.DeleteParallelDataInput{}
        output := &translate.DeleteParallelDataOutput{}

        mockClient.On("DeleteParallelData", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteParallelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTerminology", func(t *testing.T) {
        input := &translate.DeleteTerminologyInput{}
        output := &translate.DeleteTerminologyOutput{}

        mockClient.On("DeleteTerminology", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTerminology(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTextTranslationJob", func(t *testing.T) {
        input := &translate.DescribeTextTranslationJobInput{}
        output := &translate.DescribeTextTranslationJobOutput{}

        mockClient.On("DescribeTextTranslationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTextTranslationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParallelData", func(t *testing.T) {
        input := &translate.GetParallelDataInput{}
        output := &translate.GetParallelDataOutput{}

        mockClient.On("GetParallelData", ctx, input).Return(output, nil)

        result, err := mockClient.GetParallelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTerminology", func(t *testing.T) {
        input := &translate.GetTerminologyInput{}
        output := &translate.GetTerminologyOutput{}

        mockClient.On("GetTerminology", ctx, input).Return(output, nil)

        result, err := mockClient.GetTerminology(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportTerminology", func(t *testing.T) {
        input := &translate.ImportTerminologyInput{}
        output := &translate.ImportTerminologyOutput{}

        mockClient.On("ImportTerminology", ctx, input).Return(output, nil)

        result, err := mockClient.ImportTerminology(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLanguages", func(t *testing.T) {
        input := &translate.ListLanguagesInput{}
        output := &translate.ListLanguagesOutput{}

        mockClient.On("ListLanguages", ctx, input).Return(output, nil)

        result, err := mockClient.ListLanguages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParallelData", func(t *testing.T) {
        input := &translate.ListParallelDataInput{}
        output := &translate.ListParallelDataOutput{}

        mockClient.On("ListParallelData", ctx, input).Return(output, nil)

        result, err := mockClient.ListParallelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &translate.ListTagsForResourceInput{}
        output := &translate.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTerminologies", func(t *testing.T) {
        input := &translate.ListTerminologiesInput{}
        output := &translate.ListTerminologiesOutput{}

        mockClient.On("ListTerminologies", ctx, input).Return(output, nil)

        result, err := mockClient.ListTerminologies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTextTranslationJobs", func(t *testing.T) {
        input := &translate.ListTextTranslationJobsInput{}
        output := &translate.ListTextTranslationJobsOutput{}

        mockClient.On("ListTextTranslationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTextTranslationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTextTranslationJob", func(t *testing.T) {
        input := &translate.StartTextTranslationJobInput{}
        output := &translate.StartTextTranslationJobOutput{}

        mockClient.On("StartTextTranslationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartTextTranslationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTextTranslationJob", func(t *testing.T) {
        input := &translate.StopTextTranslationJobInput{}
        output := &translate.StopTextTranslationJobOutput{}

        mockClient.On("StopTextTranslationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopTextTranslationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &translate.TagResourceInput{}
        output := &translate.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTranslateDocument", func(t *testing.T) {
        input := &translate.TranslateDocumentInput{}
        output := &translate.TranslateDocumentOutput{}

        mockClient.On("TranslateDocument", ctx, input).Return(output, nil)

        result, err := mockClient.TranslateDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTranslateText", func(t *testing.T) {
        input := &translate.TranslateTextInput{}
        output := &translate.TranslateTextOutput{}

        mockClient.On("TranslateText", ctx, input).Return(output, nil)

        result, err := mockClient.TranslateText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &translate.UntagResourceInput{}
        output := &translate.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateParallelData", func(t *testing.T) {
        input := &translate.UpdateParallelDataInput{}
        output := &translate.UpdateParallelDataOutput{}

        mockClient.On("UpdateParallelData", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateParallelData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
