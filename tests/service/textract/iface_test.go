package textract_test

// tests for the textract service interface for this ../../../service/textract/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/textract/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/textract/textract_iface"
	"github.com/stretchr/testify/assert"
)

func TestTextractServiceCanBeMocked(t *testing.T) {
	var iface textract_iface.IClient
	iface = &textract.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := textract.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAnalyzeDocument", func(t *testing.T) {
        input := &textract.AnalyzeDocumentInput{}
        output := &textract.AnalyzeDocumentOutput{}

        mockClient.On("AnalyzeDocument", ctx, input).Return(output, nil)

        result, err := mockClient.AnalyzeDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAnalyzeExpense", func(t *testing.T) {
        input := &textract.AnalyzeExpenseInput{}
        output := &textract.AnalyzeExpenseOutput{}

        mockClient.On("AnalyzeExpense", ctx, input).Return(output, nil)

        result, err := mockClient.AnalyzeExpense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAnalyzeID", func(t *testing.T) {
        input := &textract.AnalyzeIDInput{}
        output := &textract.AnalyzeIDOutput{}

        mockClient.On("AnalyzeID", ctx, input).Return(output, nil)

        result, err := mockClient.AnalyzeID(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAdapter", func(t *testing.T) {
        input := &textract.CreateAdapterInput{}
        output := &textract.CreateAdapterOutput{}

        mockClient.On("CreateAdapter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAdapter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAdapterVersion", func(t *testing.T) {
        input := &textract.CreateAdapterVersionInput{}
        output := &textract.CreateAdapterVersionOutput{}

        mockClient.On("CreateAdapterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAdapterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAdapter", func(t *testing.T) {
        input := &textract.DeleteAdapterInput{}
        output := &textract.DeleteAdapterOutput{}

        mockClient.On("DeleteAdapter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAdapter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAdapterVersion", func(t *testing.T) {
        input := &textract.DeleteAdapterVersionInput{}
        output := &textract.DeleteAdapterVersionOutput{}

        mockClient.On("DeleteAdapterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAdapterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectDocumentText", func(t *testing.T) {
        input := &textract.DetectDocumentTextInput{}
        output := &textract.DetectDocumentTextOutput{}

        mockClient.On("DetectDocumentText", ctx, input).Return(output, nil)

        result, err := mockClient.DetectDocumentText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdapter", func(t *testing.T) {
        input := &textract.GetAdapterInput{}
        output := &textract.GetAdapterOutput{}

        mockClient.On("GetAdapter", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdapter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdapterVersion", func(t *testing.T) {
        input := &textract.GetAdapterVersionInput{}
        output := &textract.GetAdapterVersionOutput{}

        mockClient.On("GetAdapterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdapterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentAnalysis", func(t *testing.T) {
        input := &textract.GetDocumentAnalysisInput{}
        output := &textract.GetDocumentAnalysisOutput{}

        mockClient.On("GetDocumentAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentTextDetection", func(t *testing.T) {
        input := &textract.GetDocumentTextDetectionInput{}
        output := &textract.GetDocumentTextDetectionOutput{}

        mockClient.On("GetDocumentTextDetection", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentTextDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExpenseAnalysis", func(t *testing.T) {
        input := &textract.GetExpenseAnalysisInput{}
        output := &textract.GetExpenseAnalysisOutput{}

        mockClient.On("GetExpenseAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.GetExpenseAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLendingAnalysis", func(t *testing.T) {
        input := &textract.GetLendingAnalysisInput{}
        output := &textract.GetLendingAnalysisOutput{}

        mockClient.On("GetLendingAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.GetLendingAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLendingAnalysisSummary", func(t *testing.T) {
        input := &textract.GetLendingAnalysisSummaryInput{}
        output := &textract.GetLendingAnalysisSummaryOutput{}

        mockClient.On("GetLendingAnalysisSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetLendingAnalysisSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAdapterVersions", func(t *testing.T) {
        input := &textract.ListAdapterVersionsInput{}
        output := &textract.ListAdapterVersionsOutput{}

        mockClient.On("ListAdapterVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAdapterVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAdapters", func(t *testing.T) {
        input := &textract.ListAdaptersInput{}
        output := &textract.ListAdaptersOutput{}

        mockClient.On("ListAdapters", ctx, input).Return(output, nil)

        result, err := mockClient.ListAdapters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &textract.ListTagsForResourceInput{}
        output := &textract.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDocumentAnalysis", func(t *testing.T) {
        input := &textract.StartDocumentAnalysisInput{}
        output := &textract.StartDocumentAnalysisOutput{}

        mockClient.On("StartDocumentAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartDocumentAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDocumentTextDetection", func(t *testing.T) {
        input := &textract.StartDocumentTextDetectionInput{}
        output := &textract.StartDocumentTextDetectionOutput{}

        mockClient.On("StartDocumentTextDetection", ctx, input).Return(output, nil)

        result, err := mockClient.StartDocumentTextDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExpenseAnalysis", func(t *testing.T) {
        input := &textract.StartExpenseAnalysisInput{}
        output := &textract.StartExpenseAnalysisOutput{}

        mockClient.On("StartExpenseAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartExpenseAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLendingAnalysis", func(t *testing.T) {
        input := &textract.StartLendingAnalysisInput{}
        output := &textract.StartLendingAnalysisOutput{}

        mockClient.On("StartLendingAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartLendingAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &textract.TagResourceInput{}
        output := &textract.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &textract.UntagResourceInput{}
        output := &textract.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAdapter", func(t *testing.T) {
        input := &textract.UpdateAdapterInput{}
        output := &textract.UpdateAdapterOutput{}

        mockClient.On("UpdateAdapter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAdapter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
