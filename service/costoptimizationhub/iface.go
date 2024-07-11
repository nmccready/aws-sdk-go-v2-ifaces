
package costoptimizationhub

import (
    "github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
)

// ICostoptimizationhub defines the interface for costoptimizationhub
type ICostoptimizationhub interface {
 Options() Options 
 GetPreferences(ctx context.Context, params *GetPreferencesInput, optFns ...func(*Options)) (*GetPreferencesOutput, error) 
 GetRecommendation(ctx context.Context, params *GetRecommendationInput, optFns ...func(*Options)) (*GetRecommendationOutput, error) 
 ListEnrollmentStatuses(ctx context.Context, params *ListEnrollmentStatusesInput, optFns ...func(*Options)) (*ListEnrollmentStatusesOutput, error) 
 ListRecommendationSummaries(ctx context.Context, params *ListRecommendationSummariesInput, optFns ...func(*Options)) (*ListRecommendationSummariesOutput, error) 
 ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) 
 UpdateEnrollmentStatus(ctx context.Context, params *UpdateEnrollmentStatusInput, optFns ...func(*Options)) (*UpdateEnrollmentStatusOutput, error) 
 UpdatePreferences(ctx context.Context, params *UpdatePreferencesInput, optFns ...func(*Options)) (*UpdatePreferencesOutput, error) 
}
