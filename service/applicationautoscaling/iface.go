
package applicationautoscaling

import (
    "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

// IClient defines the interface for applicationautoscaling
type IClient interface {
 Options() Options 
 DeleteScalingPolicy(ctx context.Context, params *DeleteScalingPolicyInput, optFns ...func(*Options)) (*DeleteScalingPolicyOutput, error) 
 DeleteScheduledAction(ctx context.Context, params *DeleteScheduledActionInput, optFns ...func(*Options)) (*DeleteScheduledActionOutput, error) 
 DeregisterScalableTarget(ctx context.Context, params *DeregisterScalableTargetInput, optFns ...func(*Options)) (*DeregisterScalableTargetOutput, error) 
 DescribeScalableTargets(ctx context.Context, params *DescribeScalableTargetsInput, optFns ...func(*Options)) (*DescribeScalableTargetsOutput, error) 
 DescribeScalingActivities(ctx context.Context, params *DescribeScalingActivitiesInput, optFns ...func(*Options)) (*DescribeScalingActivitiesOutput, error) 
 DescribeScalingPolicies(ctx context.Context, params *DescribeScalingPoliciesInput, optFns ...func(*Options)) (*DescribeScalingPoliciesOutput, error) 
 DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) 
 PutScheduledAction(ctx context.Context, params *PutScheduledActionInput, optFns ...func(*Options)) (*PutScheduledActionOutput, error) 
 RegisterScalableTarget(ctx context.Context, params *RegisterScalableTargetInput, optFns ...func(*Options)) (*RegisterScalableTargetOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
