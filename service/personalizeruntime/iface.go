
package personalizeruntime

import (
    "github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
)

// IClient defines the interface for personalizeruntime
type IClient interface {
 Options() Options 
 GetActionRecommendations(ctx context.Context, params *GetActionRecommendationsInput, optFns ...func(*Options)) (*GetActionRecommendationsOutput, error) 
 GetPersonalizedRanking(ctx context.Context, params *GetPersonalizedRankingInput, optFns ...func(*Options)) (*GetPersonalizedRankingOutput, error) 
 GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) 
}
