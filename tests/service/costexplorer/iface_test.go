package costexplorer_test

// tests for the costexplorer service interface for this ../../../service/costexplorer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costexplorer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costexplorer/costexplorer_iface"
	"github.com/stretchr/testify/assert"
)

func TestCostexplorerServiceCanBeMocked(t *testing.T) {
	var iface costexplorer_iface.IClient
	iface = &costexplorer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := costexplorer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnomalyMonitor", func(t *testing.T) {
        input := &costexplorer.CreateAnomalyMonitorInput{}
        output := &costexplorer.CreateAnomalyMonitorOutput{}

        mockClient.On("CreateAnomalyMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnomalyMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnomalySubscription", func(t *testing.T) {
        input := &costexplorer.CreateAnomalySubscriptionInput{}
        output := &costexplorer.CreateAnomalySubscriptionOutput{}

        mockClient.On("CreateAnomalySubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnomalySubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCostCategoryDefinition", func(t *testing.T) {
        input := &costexplorer.CreateCostCategoryDefinitionInput{}
        output := &costexplorer.CreateCostCategoryDefinitionOutput{}

        mockClient.On("CreateCostCategoryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCostCategoryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnomalyMonitor", func(t *testing.T) {
        input := &costexplorer.DeleteAnomalyMonitorInput{}
        output := &costexplorer.DeleteAnomalyMonitorOutput{}

        mockClient.On("DeleteAnomalyMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnomalyMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnomalySubscription", func(t *testing.T) {
        input := &costexplorer.DeleteAnomalySubscriptionInput{}
        output := &costexplorer.DeleteAnomalySubscriptionOutput{}

        mockClient.On("DeleteAnomalySubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnomalySubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCostCategoryDefinition", func(t *testing.T) {
        input := &costexplorer.DeleteCostCategoryDefinitionInput{}
        output := &costexplorer.DeleteCostCategoryDefinitionOutput{}

        mockClient.On("DeleteCostCategoryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCostCategoryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCostCategoryDefinition", func(t *testing.T) {
        input := &costexplorer.DescribeCostCategoryDefinitionInput{}
        output := &costexplorer.DescribeCostCategoryDefinitionOutput{}

        mockClient.On("DescribeCostCategoryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCostCategoryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnomalies", func(t *testing.T) {
        input := &costexplorer.GetAnomaliesInput{}
        output := &costexplorer.GetAnomaliesOutput{}

        mockClient.On("GetAnomalies", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnomalies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnomalyMonitors", func(t *testing.T) {
        input := &costexplorer.GetAnomalyMonitorsInput{}
        output := &costexplorer.GetAnomalyMonitorsOutput{}

        mockClient.On("GetAnomalyMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnomalyMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnomalySubscriptions", func(t *testing.T) {
        input := &costexplorer.GetAnomalySubscriptionsInput{}
        output := &costexplorer.GetAnomalySubscriptionsOutput{}

        mockClient.On("GetAnomalySubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnomalySubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApproximateUsageRecords", func(t *testing.T) {
        input := &costexplorer.GetApproximateUsageRecordsInput{}
        output := &costexplorer.GetApproximateUsageRecordsOutput{}

        mockClient.On("GetApproximateUsageRecords", ctx, input).Return(output, nil)

        result, err := mockClient.GetApproximateUsageRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostAndUsage", func(t *testing.T) {
        input := &costexplorer.GetCostAndUsageInput{}
        output := &costexplorer.GetCostAndUsageOutput{}

        mockClient.On("GetCostAndUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostAndUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostAndUsageWithResources", func(t *testing.T) {
        input := &costexplorer.GetCostAndUsageWithResourcesInput{}
        output := &costexplorer.GetCostAndUsageWithResourcesOutput{}

        mockClient.On("GetCostAndUsageWithResources", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostAndUsageWithResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostCategories", func(t *testing.T) {
        input := &costexplorer.GetCostCategoriesInput{}
        output := &costexplorer.GetCostCategoriesOutput{}

        mockClient.On("GetCostCategories", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostForecast", func(t *testing.T) {
        input := &costexplorer.GetCostForecastInput{}
        output := &costexplorer.GetCostForecastOutput{}

        mockClient.On("GetCostForecast", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDimensionValues", func(t *testing.T) {
        input := &costexplorer.GetDimensionValuesInput{}
        output := &costexplorer.GetDimensionValuesOutput{}

        mockClient.On("GetDimensionValues", ctx, input).Return(output, nil)

        result, err := mockClient.GetDimensionValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservationCoverage", func(t *testing.T) {
        input := &costexplorer.GetReservationCoverageInput{}
        output := &costexplorer.GetReservationCoverageOutput{}

        mockClient.On("GetReservationCoverage", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservationCoverage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservationPurchaseRecommendation", func(t *testing.T) {
        input := &costexplorer.GetReservationPurchaseRecommendationInput{}
        output := &costexplorer.GetReservationPurchaseRecommendationOutput{}

        mockClient.On("GetReservationPurchaseRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservationPurchaseRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservationUtilization", func(t *testing.T) {
        input := &costexplorer.GetReservationUtilizationInput{}
        output := &costexplorer.GetReservationUtilizationOutput{}

        mockClient.On("GetReservationUtilization", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservationUtilization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRightsizingRecommendation", func(t *testing.T) {
        input := &costexplorer.GetRightsizingRecommendationInput{}
        output := &costexplorer.GetRightsizingRecommendationOutput{}

        mockClient.On("GetRightsizingRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetRightsizingRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSavingsPlanPurchaseRecommendationDetails", func(t *testing.T) {
        input := &costexplorer.GetSavingsPlanPurchaseRecommendationDetailsInput{}
        output := &costexplorer.GetSavingsPlanPurchaseRecommendationDetailsOutput{}

        mockClient.On("GetSavingsPlanPurchaseRecommendationDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetSavingsPlanPurchaseRecommendationDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSavingsPlansCoverage", func(t *testing.T) {
        input := &costexplorer.GetSavingsPlansCoverageInput{}
        output := &costexplorer.GetSavingsPlansCoverageOutput{}

        mockClient.On("GetSavingsPlansCoverage", ctx, input).Return(output, nil)

        result, err := mockClient.GetSavingsPlansCoverage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSavingsPlansPurchaseRecommendation", func(t *testing.T) {
        input := &costexplorer.GetSavingsPlansPurchaseRecommendationInput{}
        output := &costexplorer.GetSavingsPlansPurchaseRecommendationOutput{}

        mockClient.On("GetSavingsPlansPurchaseRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetSavingsPlansPurchaseRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSavingsPlansUtilization", func(t *testing.T) {
        input := &costexplorer.GetSavingsPlansUtilizationInput{}
        output := &costexplorer.GetSavingsPlansUtilizationOutput{}

        mockClient.On("GetSavingsPlansUtilization", ctx, input).Return(output, nil)

        result, err := mockClient.GetSavingsPlansUtilization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSavingsPlansUtilizationDetails", func(t *testing.T) {
        input := &costexplorer.GetSavingsPlansUtilizationDetailsInput{}
        output := &costexplorer.GetSavingsPlansUtilizationDetailsOutput{}

        mockClient.On("GetSavingsPlansUtilizationDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetSavingsPlansUtilizationDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTags", func(t *testing.T) {
        input := &costexplorer.GetTagsInput{}
        output := &costexplorer.GetTagsOutput{}

        mockClient.On("GetTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageForecast", func(t *testing.T) {
        input := &costexplorer.GetUsageForecastInput{}
        output := &costexplorer.GetUsageForecastOutput{}

        mockClient.On("GetUsageForecast", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCostAllocationTagBackfillHistory", func(t *testing.T) {
        input := &costexplorer.ListCostAllocationTagBackfillHistoryInput{}
        output := &costexplorer.ListCostAllocationTagBackfillHistoryOutput{}

        mockClient.On("ListCostAllocationTagBackfillHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListCostAllocationTagBackfillHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCostAllocationTags", func(t *testing.T) {
        input := &costexplorer.ListCostAllocationTagsInput{}
        output := &costexplorer.ListCostAllocationTagsOutput{}

        mockClient.On("ListCostAllocationTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListCostAllocationTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCostCategoryDefinitions", func(t *testing.T) {
        input := &costexplorer.ListCostCategoryDefinitionsInput{}
        output := &costexplorer.ListCostCategoryDefinitionsOutput{}

        mockClient.On("ListCostCategoryDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCostCategoryDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSavingsPlansPurchaseRecommendationGeneration", func(t *testing.T) {
        input := &costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput{}
        output := &costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput{}

        mockClient.On("ListSavingsPlansPurchaseRecommendationGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.ListSavingsPlansPurchaseRecommendationGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &costexplorer.ListTagsForResourceInput{}
        output := &costexplorer.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvideAnomalyFeedback", func(t *testing.T) {
        input := &costexplorer.ProvideAnomalyFeedbackInput{}
        output := &costexplorer.ProvideAnomalyFeedbackOutput{}

        mockClient.On("ProvideAnomalyFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.ProvideAnomalyFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCostAllocationTagBackfill", func(t *testing.T) {
        input := &costexplorer.StartCostAllocationTagBackfillInput{}
        output := &costexplorer.StartCostAllocationTagBackfillOutput{}

        mockClient.On("StartCostAllocationTagBackfill", ctx, input).Return(output, nil)

        result, err := mockClient.StartCostAllocationTagBackfill(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSavingsPlansPurchaseRecommendationGeneration", func(t *testing.T) {
        input := &costexplorer.StartSavingsPlansPurchaseRecommendationGenerationInput{}
        output := &costexplorer.StartSavingsPlansPurchaseRecommendationGenerationOutput{}

        mockClient.On("StartSavingsPlansPurchaseRecommendationGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.StartSavingsPlansPurchaseRecommendationGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &costexplorer.TagResourceInput{}
        output := &costexplorer.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &costexplorer.UntagResourceInput{}
        output := &costexplorer.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnomalyMonitor", func(t *testing.T) {
        input := &costexplorer.UpdateAnomalyMonitorInput{}
        output := &costexplorer.UpdateAnomalyMonitorOutput{}

        mockClient.On("UpdateAnomalyMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnomalyMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnomalySubscription", func(t *testing.T) {
        input := &costexplorer.UpdateAnomalySubscriptionInput{}
        output := &costexplorer.UpdateAnomalySubscriptionOutput{}

        mockClient.On("UpdateAnomalySubscription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnomalySubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCostAllocationTagsStatus", func(t *testing.T) {
        input := &costexplorer.UpdateCostAllocationTagsStatusInput{}
        output := &costexplorer.UpdateCostAllocationTagsStatusOutput{}

        mockClient.On("UpdateCostAllocationTagsStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCostAllocationTagsStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCostCategoryDefinition", func(t *testing.T) {
        input := &costexplorer.UpdateCostCategoryDefinitionInput{}
        output := &costexplorer.UpdateCostCategoryDefinitionOutput{}

        mockClient.On("UpdateCostCategoryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCostCategoryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
