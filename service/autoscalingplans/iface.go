
package autoscalingplans

import (
    "github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

// IAutoscalingplans defines the interface for autoscalingplans
type IAutoscalingplans interface {
 Options() Options 
 CreateScalingPlan(ctx context.Context, params *CreateScalingPlanInput, optFns ...func(*Options)) (*CreateScalingPlanOutput, error) 
 DeleteScalingPlan(ctx context.Context, params *DeleteScalingPlanInput, optFns ...func(*Options)) (*DeleteScalingPlanOutput, error) 
 DescribeScalingPlanResources(ctx context.Context, params *DescribeScalingPlanResourcesInput, optFns ...func(*Options)) (*DescribeScalingPlanResourcesOutput, error) 
 DescribeScalingPlans(ctx context.Context, params *DescribeScalingPlansInput, optFns ...func(*Options)) (*DescribeScalingPlansOutput, error) 
 GetScalingPlanResourceForecastData(ctx context.Context, params *GetScalingPlanResourceForecastDataInput, optFns ...func(*Options)) (*GetScalingPlanResourceForecastDataOutput, error) 
 UpdateScalingPlan(ctx context.Context, params *UpdateScalingPlanInput, optFns ...func(*Options)) (*UpdateScalingPlanOutput, error) 
}
