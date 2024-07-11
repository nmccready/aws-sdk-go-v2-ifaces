
package devopsguru

import (
    "github.com/aws/aws-sdk-go-v2/service/devopsguru"
)

// IDevopsguru defines the interface for devopsguru
type IDevopsguru interface {
 Options() Options 
 AddNotificationChannel(ctx context.Context, params *AddNotificationChannelInput, optFns ...func(*Options)) (*AddNotificationChannelOutput, error) 
 DeleteInsight(ctx context.Context, params *DeleteInsightInput, optFns ...func(*Options)) (*DeleteInsightOutput, error) 
 DescribeAccountHealth(ctx context.Context, params *DescribeAccountHealthInput, optFns ...func(*Options)) (*DescribeAccountHealthOutput, error) 
 DescribeAccountOverview(ctx context.Context, params *DescribeAccountOverviewInput, optFns ...func(*Options)) (*DescribeAccountOverviewOutput, error) 
 DescribeAnomaly(ctx context.Context, params *DescribeAnomalyInput, optFns ...func(*Options)) (*DescribeAnomalyOutput, error) 
 DescribeEventSourcesConfig(ctx context.Context, params *DescribeEventSourcesConfigInput, optFns ...func(*Options)) (*DescribeEventSourcesConfigOutput, error) 
 DescribeFeedback(ctx context.Context, params *DescribeFeedbackInput, optFns ...func(*Options)) (*DescribeFeedbackOutput, error) 
 DescribeInsight(ctx context.Context, params *DescribeInsightInput, optFns ...func(*Options)) (*DescribeInsightOutput, error) 
 DescribeOrganizationHealth(ctx context.Context, params *DescribeOrganizationHealthInput, optFns ...func(*Options)) (*DescribeOrganizationHealthOutput, error) 
 DescribeOrganizationOverview(ctx context.Context, params *DescribeOrganizationOverviewInput, optFns ...func(*Options)) (*DescribeOrganizationOverviewOutput, error) 
 DescribeOrganizationResourceCollectionHealth(ctx context.Context, params *DescribeOrganizationResourceCollectionHealthInput, optFns ...func(*Options)) (*DescribeOrganizationResourceCollectionHealthOutput, error) 
 DescribeResourceCollectionHealth(ctx context.Context, params *DescribeResourceCollectionHealthInput, optFns ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error) 
 DescribeServiceIntegration(ctx context.Context, params *DescribeServiceIntegrationInput, optFns ...func(*Options)) (*DescribeServiceIntegrationOutput, error) 
 GetCostEstimation(ctx context.Context, params *GetCostEstimationInput, optFns ...func(*Options)) (*GetCostEstimationOutput, error) 
 GetResourceCollection(ctx context.Context, params *GetResourceCollectionInput, optFns ...func(*Options)) (*GetResourceCollectionOutput, error) 
 ListAnomaliesForInsight(ctx context.Context, params *ListAnomaliesForInsightInput, optFns ...func(*Options)) (*ListAnomaliesForInsightOutput, error) 
 ListAnomalousLogGroups(ctx context.Context, params *ListAnomalousLogGroupsInput, optFns ...func(*Options)) (*ListAnomalousLogGroupsOutput, error) 
 ListEvents(ctx context.Context, params *ListEventsInput, optFns ...func(*Options)) (*ListEventsOutput, error) 
 ListInsights(ctx context.Context, params *ListInsightsInput, optFns ...func(*Options)) (*ListInsightsOutput, error) 
 ListMonitoredResources(ctx context.Context, params *ListMonitoredResourcesInput, optFns ...func(*Options)) (*ListMonitoredResourcesOutput, error) 
 ListNotificationChannels(ctx context.Context, params *ListNotificationChannelsInput, optFns ...func(*Options)) (*ListNotificationChannelsOutput, error) 
 ListOrganizationInsights(ctx context.Context, params *ListOrganizationInsightsInput, optFns ...func(*Options)) (*ListOrganizationInsightsOutput, error) 
 ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) 
 PutFeedback(ctx context.Context, params *PutFeedbackInput, optFns ...func(*Options)) (*PutFeedbackOutput, error) 
 RemoveNotificationChannel(ctx context.Context, params *RemoveNotificationChannelInput, optFns ...func(*Options)) (*RemoveNotificationChannelOutput, error) 
 SearchInsights(ctx context.Context, params *SearchInsightsInput, optFns ...func(*Options)) (*SearchInsightsOutput, error) 
 SearchOrganizationInsights(ctx context.Context, params *SearchOrganizationInsightsInput, optFns ...func(*Options)) (*SearchOrganizationInsightsOutput, error) 
 StartCostEstimation(ctx context.Context, params *StartCostEstimationInput, optFns ...func(*Options)) (*StartCostEstimationOutput, error) 
 UpdateEventSourcesConfig(ctx context.Context, params *UpdateEventSourcesConfigInput, optFns ...func(*Options)) (*UpdateEventSourcesConfigOutput, error) 
 UpdateResourceCollection(ctx context.Context, params *UpdateResourceCollectionInput, optFns ...func(*Options)) (*UpdateResourceCollectionOutput, error) 
 UpdateServiceIntegration(ctx context.Context, params *UpdateServiceIntegrationInput, optFns ...func(*Options)) (*UpdateServiceIntegrationOutput, error) 
}
