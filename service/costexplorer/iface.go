
package costexplorer

import (
    "github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

// IClient defines the interface for costexplorer
type IClient interface {
 Options() Options 
 CreateAnomalyMonitor(ctx context.Context, params *CreateAnomalyMonitorInput, optFns ...func(*Options)) (*CreateAnomalyMonitorOutput, error) 
 CreateAnomalySubscription(ctx context.Context, params *CreateAnomalySubscriptionInput, optFns ...func(*Options)) (*CreateAnomalySubscriptionOutput, error) 
 CreateCostCategoryDefinition(ctx context.Context, params *CreateCostCategoryDefinitionInput, optFns ...func(*Options)) (*CreateCostCategoryDefinitionOutput, error) 
 DeleteAnomalyMonitor(ctx context.Context, params *DeleteAnomalyMonitorInput, optFns ...func(*Options)) (*DeleteAnomalyMonitorOutput, error) 
 DeleteAnomalySubscription(ctx context.Context, params *DeleteAnomalySubscriptionInput, optFns ...func(*Options)) (*DeleteAnomalySubscriptionOutput, error) 
 DeleteCostCategoryDefinition(ctx context.Context, params *DeleteCostCategoryDefinitionInput, optFns ...func(*Options)) (*DeleteCostCategoryDefinitionOutput, error) 
 DescribeCostCategoryDefinition(ctx context.Context, params *DescribeCostCategoryDefinitionInput, optFns ...func(*Options)) (*DescribeCostCategoryDefinitionOutput, error) 
 GetAnomalies(ctx context.Context, params *GetAnomaliesInput, optFns ...func(*Options)) (*GetAnomaliesOutput, error) 
 GetAnomalyMonitors(ctx context.Context, params *GetAnomalyMonitorsInput, optFns ...func(*Options)) (*GetAnomalyMonitorsOutput, error) 
 GetAnomalySubscriptions(ctx context.Context, params *GetAnomalySubscriptionsInput, optFns ...func(*Options)) (*GetAnomalySubscriptionsOutput, error) 
 GetApproximateUsageRecords(ctx context.Context, params *GetApproximateUsageRecordsInput, optFns ...func(*Options)) (*GetApproximateUsageRecordsOutput, error) 
 GetCostAndUsage(ctx context.Context, params *GetCostAndUsageInput, optFns ...func(*Options)) (*GetCostAndUsageOutput, error) 
 GetCostAndUsageWithResources(ctx context.Context, params *GetCostAndUsageWithResourcesInput, optFns ...func(*Options)) (*GetCostAndUsageWithResourcesOutput, error) 
 GetCostCategories(ctx context.Context, params *GetCostCategoriesInput, optFns ...func(*Options)) (*GetCostCategoriesOutput, error) 
 GetCostForecast(ctx context.Context, params *GetCostForecastInput, optFns ...func(*Options)) (*GetCostForecastOutput, error) 
 GetDimensionValues(ctx context.Context, params *GetDimensionValuesInput, optFns ...func(*Options)) (*GetDimensionValuesOutput, error) 
 GetReservationCoverage(ctx context.Context, params *GetReservationCoverageInput, optFns ...func(*Options)) (*GetReservationCoverageOutput, error) 
 GetReservationPurchaseRecommendation(ctx context.Context, params *GetReservationPurchaseRecommendationInput, optFns ...func(*Options)) (*GetReservationPurchaseRecommendationOutput, error) 
 GetReservationUtilization(ctx context.Context, params *GetReservationUtilizationInput, optFns ...func(*Options)) (*GetReservationUtilizationOutput, error) 
 GetRightsizingRecommendation(ctx context.Context, params *GetRightsizingRecommendationInput, optFns ...func(*Options)) (*GetRightsizingRecommendationOutput, error) 
 GetSavingsPlanPurchaseRecommendationDetails(ctx context.Context, params *GetSavingsPlanPurchaseRecommendationDetailsInput, optFns ...func(*Options)) (*GetSavingsPlanPurchaseRecommendationDetailsOutput, error) 
 GetSavingsPlansCoverage(ctx context.Context, params *GetSavingsPlansCoverageInput, optFns ...func(*Options)) (*GetSavingsPlansCoverageOutput, error) 
 GetSavingsPlansPurchaseRecommendation(ctx context.Context, params *GetSavingsPlansPurchaseRecommendationInput, optFns ...func(*Options)) (*GetSavingsPlansPurchaseRecommendationOutput, error) 
 GetSavingsPlansUtilization(ctx context.Context, params *GetSavingsPlansUtilizationInput, optFns ...func(*Options)) (*GetSavingsPlansUtilizationOutput, error) 
 GetSavingsPlansUtilizationDetails(ctx context.Context, params *GetSavingsPlansUtilizationDetailsInput, optFns ...func(*Options)) (*GetSavingsPlansUtilizationDetailsOutput, error) 
 GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) 
 GetUsageForecast(ctx context.Context, params *GetUsageForecastInput, optFns ...func(*Options)) (*GetUsageForecastOutput, error) 
 ListCostAllocationTagBackfillHistory(ctx context.Context, params *ListCostAllocationTagBackfillHistoryInput, optFns ...func(*Options)) (*ListCostAllocationTagBackfillHistoryOutput, error) 
 ListCostAllocationTags(ctx context.Context, params *ListCostAllocationTagsInput, optFns ...func(*Options)) (*ListCostAllocationTagsOutput, error) 
 ListCostCategoryDefinitions(ctx context.Context, params *ListCostCategoryDefinitionsInput, optFns ...func(*Options)) (*ListCostCategoryDefinitionsOutput, error) 
 ListSavingsPlansPurchaseRecommendationGeneration(ctx context.Context, params *ListSavingsPlansPurchaseRecommendationGenerationInput, optFns ...func(*Options)) (*ListSavingsPlansPurchaseRecommendationGenerationOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ProvideAnomalyFeedback(ctx context.Context, params *ProvideAnomalyFeedbackInput, optFns ...func(*Options)) (*ProvideAnomalyFeedbackOutput, error) 
 StartCostAllocationTagBackfill(ctx context.Context, params *StartCostAllocationTagBackfillInput, optFns ...func(*Options)) (*StartCostAllocationTagBackfillOutput, error) 
 StartSavingsPlansPurchaseRecommendationGeneration(ctx context.Context, params *StartSavingsPlansPurchaseRecommendationGenerationInput, optFns ...func(*Options)) (*StartSavingsPlansPurchaseRecommendationGenerationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAnomalyMonitor(ctx context.Context, params *UpdateAnomalyMonitorInput, optFns ...func(*Options)) (*UpdateAnomalyMonitorOutput, error) 
 UpdateAnomalySubscription(ctx context.Context, params *UpdateAnomalySubscriptionInput, optFns ...func(*Options)) (*UpdateAnomalySubscriptionOutput, error) 
 UpdateCostAllocationTagsStatus(ctx context.Context, params *UpdateCostAllocationTagsStatusInput, optFns ...func(*Options)) (*UpdateCostAllocationTagsStatusOutput, error) 
 UpdateCostCategoryDefinition(ctx context.Context, params *UpdateCostCategoryDefinitionInput, optFns ...func(*Options)) (*UpdateCostCategoryDefinitionOutput, error) 
}
