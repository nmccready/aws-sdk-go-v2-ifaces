package lexmodelsv2_test

// tests for the lexmodelsv2 service interface for this ../../../service/lexmodelsv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexmodelsv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexmodelsv2/lexmodelsv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestLexmodelsv2ServiceCanBeMocked(t *testing.T) {
	var iface lexmodelsv2_iface.IClient
	iface = &lexmodelsv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lexmodelsv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateCustomVocabularyItem", func(t *testing.T) {
        input := &lexmodelsv2.BatchCreateCustomVocabularyItemInput{}
        output := &lexmodelsv2.BatchCreateCustomVocabularyItemOutput{}

        mockClient.On("BatchCreateCustomVocabularyItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateCustomVocabularyItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteCustomVocabularyItem", func(t *testing.T) {
        input := &lexmodelsv2.BatchDeleteCustomVocabularyItemInput{}
        output := &lexmodelsv2.BatchDeleteCustomVocabularyItemOutput{}

        mockClient.On("BatchDeleteCustomVocabularyItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteCustomVocabularyItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateCustomVocabularyItem", func(t *testing.T) {
        input := &lexmodelsv2.BatchUpdateCustomVocabularyItemInput{}
        output := &lexmodelsv2.BatchUpdateCustomVocabularyItemOutput{}

        mockClient.On("BatchUpdateCustomVocabularyItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateCustomVocabularyItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBuildBotLocale", func(t *testing.T) {
        input := &lexmodelsv2.BuildBotLocaleInput{}
        output := &lexmodelsv2.BuildBotLocaleOutput{}

        mockClient.On("BuildBotLocale", ctx, input).Return(output, nil)

        result, err := mockClient.BuildBotLocale(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBot", func(t *testing.T) {
        input := &lexmodelsv2.CreateBotInput{}
        output := &lexmodelsv2.CreateBotOutput{}

        mockClient.On("CreateBot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBotAlias", func(t *testing.T) {
        input := &lexmodelsv2.CreateBotAliasInput{}
        output := &lexmodelsv2.CreateBotAliasOutput{}

        mockClient.On("CreateBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBotLocale", func(t *testing.T) {
        input := &lexmodelsv2.CreateBotLocaleInput{}
        output := &lexmodelsv2.CreateBotLocaleOutput{}

        mockClient.On("CreateBotLocale", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBotLocale(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBotReplica", func(t *testing.T) {
        input := &lexmodelsv2.CreateBotReplicaInput{}
        output := &lexmodelsv2.CreateBotReplicaOutput{}

        mockClient.On("CreateBotReplica", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBotReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBotVersion", func(t *testing.T) {
        input := &lexmodelsv2.CreateBotVersionInput{}
        output := &lexmodelsv2.CreateBotVersionOutput{}

        mockClient.On("CreateBotVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBotVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExport", func(t *testing.T) {
        input := &lexmodelsv2.CreateExportInput{}
        output := &lexmodelsv2.CreateExportOutput{}

        mockClient.On("CreateExport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntent", func(t *testing.T) {
        input := &lexmodelsv2.CreateIntentInput{}
        output := &lexmodelsv2.CreateIntentOutput{}

        mockClient.On("CreateIntent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourcePolicy", func(t *testing.T) {
        input := &lexmodelsv2.CreateResourcePolicyInput{}
        output := &lexmodelsv2.CreateResourcePolicyOutput{}

        mockClient.On("CreateResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourcePolicyStatement", func(t *testing.T) {
        input := &lexmodelsv2.CreateResourcePolicyStatementInput{}
        output := &lexmodelsv2.CreateResourcePolicyStatementOutput{}

        mockClient.On("CreateResourcePolicyStatement", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourcePolicyStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSlot", func(t *testing.T) {
        input := &lexmodelsv2.CreateSlotInput{}
        output := &lexmodelsv2.CreateSlotOutput{}

        mockClient.On("CreateSlot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSlot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSlotType", func(t *testing.T) {
        input := &lexmodelsv2.CreateSlotTypeInput{}
        output := &lexmodelsv2.CreateSlotTypeOutput{}

        mockClient.On("CreateSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestSetDiscrepancyReport", func(t *testing.T) {
        input := &lexmodelsv2.CreateTestSetDiscrepancyReportInput{}
        output := &lexmodelsv2.CreateTestSetDiscrepancyReportOutput{}

        mockClient.On("CreateTestSetDiscrepancyReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestSetDiscrepancyReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUploadUrl", func(t *testing.T) {
        input := &lexmodelsv2.CreateUploadUrlInput{}
        output := &lexmodelsv2.CreateUploadUrlOutput{}

        mockClient.On("CreateUploadUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUploadUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBot", func(t *testing.T) {
        input := &lexmodelsv2.DeleteBotInput{}
        output := &lexmodelsv2.DeleteBotOutput{}

        mockClient.On("DeleteBot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotAlias", func(t *testing.T) {
        input := &lexmodelsv2.DeleteBotAliasInput{}
        output := &lexmodelsv2.DeleteBotAliasOutput{}

        mockClient.On("DeleteBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotLocale", func(t *testing.T) {
        input := &lexmodelsv2.DeleteBotLocaleInput{}
        output := &lexmodelsv2.DeleteBotLocaleOutput{}

        mockClient.On("DeleteBotLocale", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotLocale(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotReplica", func(t *testing.T) {
        input := &lexmodelsv2.DeleteBotReplicaInput{}
        output := &lexmodelsv2.DeleteBotReplicaOutput{}

        mockClient.On("DeleteBotReplica", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotVersion", func(t *testing.T) {
        input := &lexmodelsv2.DeleteBotVersionInput{}
        output := &lexmodelsv2.DeleteBotVersionOutput{}

        mockClient.On("DeleteBotVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomVocabulary", func(t *testing.T) {
        input := &lexmodelsv2.DeleteCustomVocabularyInput{}
        output := &lexmodelsv2.DeleteCustomVocabularyOutput{}

        mockClient.On("DeleteCustomVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExport", func(t *testing.T) {
        input := &lexmodelsv2.DeleteExportInput{}
        output := &lexmodelsv2.DeleteExportOutput{}

        mockClient.On("DeleteExport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImport", func(t *testing.T) {
        input := &lexmodelsv2.DeleteImportInput{}
        output := &lexmodelsv2.DeleteImportOutput{}

        mockClient.On("DeleteImport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntent", func(t *testing.T) {
        input := &lexmodelsv2.DeleteIntentInput{}
        output := &lexmodelsv2.DeleteIntentOutput{}

        mockClient.On("DeleteIntent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &lexmodelsv2.DeleteResourcePolicyInput{}
        output := &lexmodelsv2.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicyStatement", func(t *testing.T) {
        input := &lexmodelsv2.DeleteResourcePolicyStatementInput{}
        output := &lexmodelsv2.DeleteResourcePolicyStatementOutput{}

        mockClient.On("DeleteResourcePolicyStatement", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicyStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlot", func(t *testing.T) {
        input := &lexmodelsv2.DeleteSlotInput{}
        output := &lexmodelsv2.DeleteSlotOutput{}

        mockClient.On("DeleteSlot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlotType", func(t *testing.T) {
        input := &lexmodelsv2.DeleteSlotTypeInput{}
        output := &lexmodelsv2.DeleteSlotTypeOutput{}

        mockClient.On("DeleteSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestSet", func(t *testing.T) {
        input := &lexmodelsv2.DeleteTestSetInput{}
        output := &lexmodelsv2.DeleteTestSetOutput{}

        mockClient.On("DeleteTestSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUtterances", func(t *testing.T) {
        input := &lexmodelsv2.DeleteUtterancesInput{}
        output := &lexmodelsv2.DeleteUtterancesOutput{}

        mockClient.On("DeleteUtterances", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUtterances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBot", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotInput{}
        output := &lexmodelsv2.DescribeBotOutput{}

        mockClient.On("DescribeBot", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotAlias", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotAliasInput{}
        output := &lexmodelsv2.DescribeBotAliasOutput{}

        mockClient.On("DescribeBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotLocale", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotLocaleInput{}
        output := &lexmodelsv2.DescribeBotLocaleOutput{}

        mockClient.On("DescribeBotLocale", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotLocale(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotRecommendation", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotRecommendationInput{}
        output := &lexmodelsv2.DescribeBotRecommendationOutput{}

        mockClient.On("DescribeBotRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotReplica", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotReplicaInput{}
        output := &lexmodelsv2.DescribeBotReplicaOutput{}

        mockClient.On("DescribeBotReplica", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotResourceGeneration", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotResourceGenerationInput{}
        output := &lexmodelsv2.DescribeBotResourceGenerationOutput{}

        mockClient.On("DescribeBotResourceGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotResourceGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBotVersion", func(t *testing.T) {
        input := &lexmodelsv2.DescribeBotVersionInput{}
        output := &lexmodelsv2.DescribeBotVersionOutput{}

        mockClient.On("DescribeBotVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBotVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomVocabularyMetadata", func(t *testing.T) {
        input := &lexmodelsv2.DescribeCustomVocabularyMetadataInput{}
        output := &lexmodelsv2.DescribeCustomVocabularyMetadataOutput{}

        mockClient.On("DescribeCustomVocabularyMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomVocabularyMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExport", func(t *testing.T) {
        input := &lexmodelsv2.DescribeExportInput{}
        output := &lexmodelsv2.DescribeExportOutput{}

        mockClient.On("DescribeExport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImport", func(t *testing.T) {
        input := &lexmodelsv2.DescribeImportInput{}
        output := &lexmodelsv2.DescribeImportOutput{}

        mockClient.On("DescribeImport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIntent", func(t *testing.T) {
        input := &lexmodelsv2.DescribeIntentInput{}
        output := &lexmodelsv2.DescribeIntentOutput{}

        mockClient.On("DescribeIntent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicy", func(t *testing.T) {
        input := &lexmodelsv2.DescribeResourcePolicyInput{}
        output := &lexmodelsv2.DescribeResourcePolicyOutput{}

        mockClient.On("DescribeResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSlot", func(t *testing.T) {
        input := &lexmodelsv2.DescribeSlotInput{}
        output := &lexmodelsv2.DescribeSlotOutput{}

        mockClient.On("DescribeSlot", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSlot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSlotType", func(t *testing.T) {
        input := &lexmodelsv2.DescribeSlotTypeInput{}
        output := &lexmodelsv2.DescribeSlotTypeOutput{}

        mockClient.On("DescribeSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTestExecution", func(t *testing.T) {
        input := &lexmodelsv2.DescribeTestExecutionInput{}
        output := &lexmodelsv2.DescribeTestExecutionOutput{}

        mockClient.On("DescribeTestExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTestExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTestSet", func(t *testing.T) {
        input := &lexmodelsv2.DescribeTestSetInput{}
        output := &lexmodelsv2.DescribeTestSetOutput{}

        mockClient.On("DescribeTestSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTestSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTestSetDiscrepancyReport", func(t *testing.T) {
        input := &lexmodelsv2.DescribeTestSetDiscrepancyReportInput{}
        output := &lexmodelsv2.DescribeTestSetDiscrepancyReportOutput{}

        mockClient.On("DescribeTestSetDiscrepancyReport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTestSetDiscrepancyReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTestSetGeneration", func(t *testing.T) {
        input := &lexmodelsv2.DescribeTestSetGenerationInput{}
        output := &lexmodelsv2.DescribeTestSetGenerationOutput{}

        mockClient.On("DescribeTestSetGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTestSetGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateBotElement", func(t *testing.T) {
        input := &lexmodelsv2.GenerateBotElementInput{}
        output := &lexmodelsv2.GenerateBotElementOutput{}

        mockClient.On("GenerateBotElement", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateBotElement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestExecutionArtifactsUrl", func(t *testing.T) {
        input := &lexmodelsv2.GetTestExecutionArtifactsUrlInput{}
        output := &lexmodelsv2.GetTestExecutionArtifactsUrlOutput{}

        mockClient.On("GetTestExecutionArtifactsUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestExecutionArtifactsUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAggregatedUtterances", func(t *testing.T) {
        input := &lexmodelsv2.ListAggregatedUtterancesInput{}
        output := &lexmodelsv2.ListAggregatedUtterancesOutput{}

        mockClient.On("ListAggregatedUtterances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAggregatedUtterances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotAliasReplicas", func(t *testing.T) {
        input := &lexmodelsv2.ListBotAliasReplicasInput{}
        output := &lexmodelsv2.ListBotAliasReplicasOutput{}

        mockClient.On("ListBotAliasReplicas", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotAliasReplicas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotAliases", func(t *testing.T) {
        input := &lexmodelsv2.ListBotAliasesInput{}
        output := &lexmodelsv2.ListBotAliasesOutput{}

        mockClient.On("ListBotAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotLocales", func(t *testing.T) {
        input := &lexmodelsv2.ListBotLocalesInput{}
        output := &lexmodelsv2.ListBotLocalesOutput{}

        mockClient.On("ListBotLocales", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotLocales(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotRecommendations", func(t *testing.T) {
        input := &lexmodelsv2.ListBotRecommendationsInput{}
        output := &lexmodelsv2.ListBotRecommendationsOutput{}

        mockClient.On("ListBotRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotReplicas", func(t *testing.T) {
        input := &lexmodelsv2.ListBotReplicasInput{}
        output := &lexmodelsv2.ListBotReplicasOutput{}

        mockClient.On("ListBotReplicas", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotReplicas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotResourceGenerations", func(t *testing.T) {
        input := &lexmodelsv2.ListBotResourceGenerationsInput{}
        output := &lexmodelsv2.ListBotResourceGenerationsOutput{}

        mockClient.On("ListBotResourceGenerations", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotResourceGenerations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotVersionReplicas", func(t *testing.T) {
        input := &lexmodelsv2.ListBotVersionReplicasInput{}
        output := &lexmodelsv2.ListBotVersionReplicasOutput{}

        mockClient.On("ListBotVersionReplicas", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotVersionReplicas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBotVersions", func(t *testing.T) {
        input := &lexmodelsv2.ListBotVersionsInput{}
        output := &lexmodelsv2.ListBotVersionsOutput{}

        mockClient.On("ListBotVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListBotVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBots", func(t *testing.T) {
        input := &lexmodelsv2.ListBotsInput{}
        output := &lexmodelsv2.ListBotsOutput{}

        mockClient.On("ListBots", ctx, input).Return(output, nil)

        result, err := mockClient.ListBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuiltInIntents", func(t *testing.T) {
        input := &lexmodelsv2.ListBuiltInIntentsInput{}
        output := &lexmodelsv2.ListBuiltInIntentsOutput{}

        mockClient.On("ListBuiltInIntents", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuiltInIntents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuiltInSlotTypes", func(t *testing.T) {
        input := &lexmodelsv2.ListBuiltInSlotTypesInput{}
        output := &lexmodelsv2.ListBuiltInSlotTypesOutput{}

        mockClient.On("ListBuiltInSlotTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuiltInSlotTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomVocabularyItems", func(t *testing.T) {
        input := &lexmodelsv2.ListCustomVocabularyItemsInput{}
        output := &lexmodelsv2.ListCustomVocabularyItemsOutput{}

        mockClient.On("ListCustomVocabularyItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomVocabularyItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExports", func(t *testing.T) {
        input := &lexmodelsv2.ListExportsInput{}
        output := &lexmodelsv2.ListExportsOutput{}

        mockClient.On("ListExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImports", func(t *testing.T) {
        input := &lexmodelsv2.ListImportsInput{}
        output := &lexmodelsv2.ListImportsOutput{}

        mockClient.On("ListImports", ctx, input).Return(output, nil)

        result, err := mockClient.ListImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntentMetrics", func(t *testing.T) {
        input := &lexmodelsv2.ListIntentMetricsInput{}
        output := &lexmodelsv2.ListIntentMetricsOutput{}

        mockClient.On("ListIntentMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntentMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntentPaths", func(t *testing.T) {
        input := &lexmodelsv2.ListIntentPathsInput{}
        output := &lexmodelsv2.ListIntentPathsOutput{}

        mockClient.On("ListIntentPaths", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntentPaths(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntentStageMetrics", func(t *testing.T) {
        input := &lexmodelsv2.ListIntentStageMetricsInput{}
        output := &lexmodelsv2.ListIntentStageMetricsOutput{}

        mockClient.On("ListIntentStageMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntentStageMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntents", func(t *testing.T) {
        input := &lexmodelsv2.ListIntentsInput{}
        output := &lexmodelsv2.ListIntentsOutput{}

        mockClient.On("ListIntents", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendedIntents", func(t *testing.T) {
        input := &lexmodelsv2.ListRecommendedIntentsInput{}
        output := &lexmodelsv2.ListRecommendedIntentsOutput{}

        mockClient.On("ListRecommendedIntents", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendedIntents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessionAnalyticsData", func(t *testing.T) {
        input := &lexmodelsv2.ListSessionAnalyticsDataInput{}
        output := &lexmodelsv2.ListSessionAnalyticsDataOutput{}

        mockClient.On("ListSessionAnalyticsData", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessionAnalyticsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessionMetrics", func(t *testing.T) {
        input := &lexmodelsv2.ListSessionMetricsInput{}
        output := &lexmodelsv2.ListSessionMetricsOutput{}

        mockClient.On("ListSessionMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessionMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSlotTypes", func(t *testing.T) {
        input := &lexmodelsv2.ListSlotTypesInput{}
        output := &lexmodelsv2.ListSlotTypesOutput{}

        mockClient.On("ListSlotTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListSlotTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSlots", func(t *testing.T) {
        input := &lexmodelsv2.ListSlotsInput{}
        output := &lexmodelsv2.ListSlotsOutput{}

        mockClient.On("ListSlots", ctx, input).Return(output, nil)

        result, err := mockClient.ListSlots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &lexmodelsv2.ListTagsForResourceInput{}
        output := &lexmodelsv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestExecutionResultItems", func(t *testing.T) {
        input := &lexmodelsv2.ListTestExecutionResultItemsInput{}
        output := &lexmodelsv2.ListTestExecutionResultItemsOutput{}

        mockClient.On("ListTestExecutionResultItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestExecutionResultItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestExecutions", func(t *testing.T) {
        input := &lexmodelsv2.ListTestExecutionsInput{}
        output := &lexmodelsv2.ListTestExecutionsOutput{}

        mockClient.On("ListTestExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestSetRecords", func(t *testing.T) {
        input := &lexmodelsv2.ListTestSetRecordsInput{}
        output := &lexmodelsv2.ListTestSetRecordsOutput{}

        mockClient.On("ListTestSetRecords", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestSetRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestSets", func(t *testing.T) {
        input := &lexmodelsv2.ListTestSetsInput{}
        output := &lexmodelsv2.ListTestSetsOutput{}

        mockClient.On("ListTestSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUtteranceAnalyticsData", func(t *testing.T) {
        input := &lexmodelsv2.ListUtteranceAnalyticsDataInput{}
        output := &lexmodelsv2.ListUtteranceAnalyticsDataOutput{}

        mockClient.On("ListUtteranceAnalyticsData", ctx, input).Return(output, nil)

        result, err := mockClient.ListUtteranceAnalyticsData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUtteranceMetrics", func(t *testing.T) {
        input := &lexmodelsv2.ListUtteranceMetricsInput{}
        output := &lexmodelsv2.ListUtteranceMetricsOutput{}

        mockClient.On("ListUtteranceMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListUtteranceMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAssociatedTranscripts", func(t *testing.T) {
        input := &lexmodelsv2.SearchAssociatedTranscriptsInput{}
        output := &lexmodelsv2.SearchAssociatedTranscriptsOutput{}

        mockClient.On("SearchAssociatedTranscripts", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAssociatedTranscripts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBotRecommendation", func(t *testing.T) {
        input := &lexmodelsv2.StartBotRecommendationInput{}
        output := &lexmodelsv2.StartBotRecommendationOutput{}

        mockClient.On("StartBotRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.StartBotRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBotResourceGeneration", func(t *testing.T) {
        input := &lexmodelsv2.StartBotResourceGenerationInput{}
        output := &lexmodelsv2.StartBotResourceGenerationOutput{}

        mockClient.On("StartBotResourceGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.StartBotResourceGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImport", func(t *testing.T) {
        input := &lexmodelsv2.StartImportInput{}
        output := &lexmodelsv2.StartImportOutput{}

        mockClient.On("StartImport", ctx, input).Return(output, nil)

        result, err := mockClient.StartImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTestExecution", func(t *testing.T) {
        input := &lexmodelsv2.StartTestExecutionInput{}
        output := &lexmodelsv2.StartTestExecutionOutput{}

        mockClient.On("StartTestExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartTestExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTestSetGeneration", func(t *testing.T) {
        input := &lexmodelsv2.StartTestSetGenerationInput{}
        output := &lexmodelsv2.StartTestSetGenerationOutput{}

        mockClient.On("StartTestSetGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.StartTestSetGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBotRecommendation", func(t *testing.T) {
        input := &lexmodelsv2.StopBotRecommendationInput{}
        output := &lexmodelsv2.StopBotRecommendationOutput{}

        mockClient.On("StopBotRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.StopBotRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lexmodelsv2.TagResourceInput{}
        output := &lexmodelsv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lexmodelsv2.UntagResourceInput{}
        output := &lexmodelsv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBot", func(t *testing.T) {
        input := &lexmodelsv2.UpdateBotInput{}
        output := &lexmodelsv2.UpdateBotOutput{}

        mockClient.On("UpdateBot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBotAlias", func(t *testing.T) {
        input := &lexmodelsv2.UpdateBotAliasInput{}
        output := &lexmodelsv2.UpdateBotAliasOutput{}

        mockClient.On("UpdateBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBotLocale", func(t *testing.T) {
        input := &lexmodelsv2.UpdateBotLocaleInput{}
        output := &lexmodelsv2.UpdateBotLocaleOutput{}

        mockClient.On("UpdateBotLocale", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBotLocale(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBotRecommendation", func(t *testing.T) {
        input := &lexmodelsv2.UpdateBotRecommendationInput{}
        output := &lexmodelsv2.UpdateBotRecommendationOutput{}

        mockClient.On("UpdateBotRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBotRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExport", func(t *testing.T) {
        input := &lexmodelsv2.UpdateExportInput{}
        output := &lexmodelsv2.UpdateExportOutput{}

        mockClient.On("UpdateExport", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntent", func(t *testing.T) {
        input := &lexmodelsv2.UpdateIntentInput{}
        output := &lexmodelsv2.UpdateIntentOutput{}

        mockClient.On("UpdateIntent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourcePolicy", func(t *testing.T) {
        input := &lexmodelsv2.UpdateResourcePolicyInput{}
        output := &lexmodelsv2.UpdateResourcePolicyOutput{}

        mockClient.On("UpdateResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSlot", func(t *testing.T) {
        input := &lexmodelsv2.UpdateSlotInput{}
        output := &lexmodelsv2.UpdateSlotOutput{}

        mockClient.On("UpdateSlot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSlot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSlotType", func(t *testing.T) {
        input := &lexmodelsv2.UpdateSlotTypeInput{}
        output := &lexmodelsv2.UpdateSlotTypeOutput{}

        mockClient.On("UpdateSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTestSet", func(t *testing.T) {
        input := &lexmodelsv2.UpdateTestSetInput{}
        output := &lexmodelsv2.UpdateTestSetOutput{}

        mockClient.On("UpdateTestSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTestSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
