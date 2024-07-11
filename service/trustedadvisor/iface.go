
package trustedadvisor

import (
    "github.com/aws/aws-sdk-go-v2/service/trustedadvisor"
)

// ITrustedadvisor defines the interface for trustedadvisor
type ITrustedadvisor interface {
 Options() Options 
 BatchUpdateRecommendationResourceExclusion(ctx context.Context, params *BatchUpdateRecommendationResourceExclusionInput, optFns ...func(*Options)) (*BatchUpdateRecommendationResourceExclusionOutput, error) 
 GetOrganizationRecommendation(ctx context.Context, params *GetOrganizationRecommendationInput, optFns ...func(*Options)) (*GetOrganizationRecommendationOutput, error) 
 GetRecommendation(ctx context.Context, params *GetRecommendationInput, optFns ...func(*Options)) (*GetRecommendationOutput, error) 
 ListChecks(ctx context.Context, params *ListChecksInput, optFns ...func(*Options)) (*ListChecksOutput, error) 
 ListOrganizationRecommendationAccounts(ctx context.Context, params *ListOrganizationRecommendationAccountsInput, optFns ...func(*Options)) (*ListOrganizationRecommendationAccountsOutput, error) 
 ListOrganizationRecommendationResources(ctx context.Context, params *ListOrganizationRecommendationResourcesInput, optFns ...func(*Options)) (*ListOrganizationRecommendationResourcesOutput, error) 
 ListOrganizationRecommendations(ctx context.Context, params *ListOrganizationRecommendationsInput, optFns ...func(*Options)) (*ListOrganizationRecommendationsOutput, error) 
 ListRecommendationResources(ctx context.Context, params *ListRecommendationResourcesInput, optFns ...func(*Options)) (*ListRecommendationResourcesOutput, error) 
 ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) 
 UpdateOrganizationRecommendationLifecycle(ctx context.Context, params *UpdateOrganizationRecommendationLifecycleInput, optFns ...func(*Options)) (*UpdateOrganizationRecommendationLifecycleOutput, error) 
 UpdateRecommendationLifecycle(ctx context.Context, params *UpdateRecommendationLifecycleInput, optFns ...func(*Options)) (*UpdateRecommendationLifecycleOutput, error) 
}
