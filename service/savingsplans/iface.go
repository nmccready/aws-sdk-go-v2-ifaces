
package savingsplans

import (
    "github.com/aws/aws-sdk-go-v2/service/savingsplans"
)

// ISavingsplans defines the interface for savingsplans
type ISavingsplans interface {
 Options() Options 
 CreateSavingsPlan(ctx context.Context, params *CreateSavingsPlanInput, optFns ...func(*Options)) (*CreateSavingsPlanOutput, error) 
 DeleteQueuedSavingsPlan(ctx context.Context, params *DeleteQueuedSavingsPlanInput, optFns ...func(*Options)) (*DeleteQueuedSavingsPlanOutput, error) 
 DescribeSavingsPlanRates(ctx context.Context, params *DescribeSavingsPlanRatesInput, optFns ...func(*Options)) (*DescribeSavingsPlanRatesOutput, error) 
 DescribeSavingsPlans(ctx context.Context, params *DescribeSavingsPlansInput, optFns ...func(*Options)) (*DescribeSavingsPlansOutput, error) 
 DescribeSavingsPlansOfferingRates(ctx context.Context, params *DescribeSavingsPlansOfferingRatesInput, optFns ...func(*Options)) (*DescribeSavingsPlansOfferingRatesOutput, error) 
 DescribeSavingsPlansOfferings(ctx context.Context, params *DescribeSavingsPlansOfferingsInput, optFns ...func(*Options)) (*DescribeSavingsPlansOfferingsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ReturnSavingsPlan(ctx context.Context, params *ReturnSavingsPlanInput, optFns ...func(*Options)) (*ReturnSavingsPlanOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
