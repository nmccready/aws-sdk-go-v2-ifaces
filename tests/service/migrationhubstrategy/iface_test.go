package migrationhubstrategy_test

// tests for the migrationhubstrategy service interface for this ../../../service/migrationhubstrategy/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubstrategy/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubstrategy/migrationhubstrategy_iface"
	"github.com/stretchr/testify/assert"
)

func TestMigrationhubstrategyServiceCanBeMocked(t *testing.T) {
	var iface migrationhubstrategy_iface.IClient
	iface = &migrationhubstrategy.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := migrationhubstrategy.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationComponentDetails", func(t *testing.T) {
        input := &migrationhubstrategy.GetApplicationComponentDetailsInput{}
        output := &migrationhubstrategy.GetApplicationComponentDetailsOutput{}

        mockClient.On("GetApplicationComponentDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationComponentDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationComponentStrategies", func(t *testing.T) {
        input := &migrationhubstrategy.GetApplicationComponentStrategiesInput{}
        output := &migrationhubstrategy.GetApplicationComponentStrategiesOutput{}

        mockClient.On("GetApplicationComponentStrategies", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationComponentStrategies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssessment", func(t *testing.T) {
        input := &migrationhubstrategy.GetAssessmentInput{}
        output := &migrationhubstrategy.GetAssessmentOutput{}

        mockClient.On("GetAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportFileTask", func(t *testing.T) {
        input := &migrationhubstrategy.GetImportFileTaskInput{}
        output := &migrationhubstrategy.GetImportFileTaskOutput{}

        mockClient.On("GetImportFileTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportFileTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLatestAssessmentId", func(t *testing.T) {
        input := &migrationhubstrategy.GetLatestAssessmentIdInput{}
        output := &migrationhubstrategy.GetLatestAssessmentIdOutput{}

        mockClient.On("GetLatestAssessmentId", ctx, input).Return(output, nil)

        result, err := mockClient.GetLatestAssessmentId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPortfolioPreferences", func(t *testing.T) {
        input := &migrationhubstrategy.GetPortfolioPreferencesInput{}
        output := &migrationhubstrategy.GetPortfolioPreferencesOutput{}

        mockClient.On("GetPortfolioPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetPortfolioPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPortfolioSummary", func(t *testing.T) {
        input := &migrationhubstrategy.GetPortfolioSummaryInput{}
        output := &migrationhubstrategy.GetPortfolioSummaryOutput{}

        mockClient.On("GetPortfolioSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetPortfolioSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendationReportDetails", func(t *testing.T) {
        input := &migrationhubstrategy.GetRecommendationReportDetailsInput{}
        output := &migrationhubstrategy.GetRecommendationReportDetailsOutput{}

        mockClient.On("GetRecommendationReportDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendationReportDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServerDetails", func(t *testing.T) {
        input := &migrationhubstrategy.GetServerDetailsInput{}
        output := &migrationhubstrategy.GetServerDetailsOutput{}

        mockClient.On("GetServerDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetServerDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServerStrategies", func(t *testing.T) {
        input := &migrationhubstrategy.GetServerStrategiesInput{}
        output := &migrationhubstrategy.GetServerStrategiesOutput{}

        mockClient.On("GetServerStrategies", ctx, input).Return(output, nil)

        result, err := mockClient.GetServerStrategies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalyzableServers", func(t *testing.T) {
        input := &migrationhubstrategy.ListAnalyzableServersInput{}
        output := &migrationhubstrategy.ListAnalyzableServersOutput{}

        mockClient.On("ListAnalyzableServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalyzableServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationComponents", func(t *testing.T) {
        input := &migrationhubstrategy.ListApplicationComponentsInput{}
        output := &migrationhubstrategy.ListApplicationComponentsOutput{}

        mockClient.On("ListApplicationComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollectors", func(t *testing.T) {
        input := &migrationhubstrategy.ListCollectorsInput{}
        output := &migrationhubstrategy.ListCollectorsOutput{}

        mockClient.On("ListCollectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportFileTask", func(t *testing.T) {
        input := &migrationhubstrategy.ListImportFileTaskInput{}
        output := &migrationhubstrategy.ListImportFileTaskOutput{}

        mockClient.On("ListImportFileTask", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportFileTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServers", func(t *testing.T) {
        input := &migrationhubstrategy.ListServersInput{}
        output := &migrationhubstrategy.ListServersOutput{}

        mockClient.On("ListServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPortfolioPreferences", func(t *testing.T) {
        input := &migrationhubstrategy.PutPortfolioPreferencesInput{}
        output := &migrationhubstrategy.PutPortfolioPreferencesOutput{}

        mockClient.On("PutPortfolioPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.PutPortfolioPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssessment", func(t *testing.T) {
        input := &migrationhubstrategy.StartAssessmentInput{}
        output := &migrationhubstrategy.StartAssessmentOutput{}

        mockClient.On("StartAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportFileTask", func(t *testing.T) {
        input := &migrationhubstrategy.StartImportFileTaskInput{}
        output := &migrationhubstrategy.StartImportFileTaskOutput{}

        mockClient.On("StartImportFileTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportFileTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRecommendationReportGeneration", func(t *testing.T) {
        input := &migrationhubstrategy.StartRecommendationReportGenerationInput{}
        output := &migrationhubstrategy.StartRecommendationReportGenerationOutput{}

        mockClient.On("StartRecommendationReportGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.StartRecommendationReportGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAssessment", func(t *testing.T) {
        input := &migrationhubstrategy.StopAssessmentInput{}
        output := &migrationhubstrategy.StopAssessmentOutput{}

        mockClient.On("StopAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.StopAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationComponentConfig", func(t *testing.T) {
        input := &migrationhubstrategy.UpdateApplicationComponentConfigInput{}
        output := &migrationhubstrategy.UpdateApplicationComponentConfigOutput{}

        mockClient.On("UpdateApplicationComponentConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationComponentConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServerConfig", func(t *testing.T) {
        input := &migrationhubstrategy.UpdateServerConfigInput{}
        output := &migrationhubstrategy.UpdateServerConfigOutput{}

        mockClient.On("UpdateServerConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServerConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
