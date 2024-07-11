
package kendraranking

import (
    "github.com/aws/aws-sdk-go-v2/service/kendraranking"
)

// IClient defines the interface for kendraranking
type IClient interface {
 Options() Options 
 CreateRescoreExecutionPlan(ctx context.Context, params *CreateRescoreExecutionPlanInput, optFns ...func(*Options)) (*CreateRescoreExecutionPlanOutput, error) 
 DeleteRescoreExecutionPlan(ctx context.Context, params *DeleteRescoreExecutionPlanInput, optFns ...func(*Options)) (*DeleteRescoreExecutionPlanOutput, error) 
 DescribeRescoreExecutionPlan(ctx context.Context, params *DescribeRescoreExecutionPlanInput, optFns ...func(*Options)) (*DescribeRescoreExecutionPlanOutput, error) 
 ListRescoreExecutionPlans(ctx context.Context, params *ListRescoreExecutionPlansInput, optFns ...func(*Options)) (*ListRescoreExecutionPlansOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 Rescore(ctx context.Context, params *RescoreInput, optFns ...func(*Options)) (*RescoreOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateRescoreExecutionPlan(ctx context.Context, params *UpdateRescoreExecutionPlanInput, optFns ...func(*Options)) (*UpdateRescoreExecutionPlanOutput, error) 
}
