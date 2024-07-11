
package codestarnotifications

import (
    "github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

// IClient defines the interface for codestarnotifications
type IClient interface {
 Options() Options 
 CreateNotificationRule(ctx context.Context, params *CreateNotificationRuleInput, optFns ...func(*Options)) (*CreateNotificationRuleOutput, error) 
 DeleteNotificationRule(ctx context.Context, params *DeleteNotificationRuleInput, optFns ...func(*Options)) (*DeleteNotificationRuleOutput, error) 
 DeleteTarget(ctx context.Context, params *DeleteTargetInput, optFns ...func(*Options)) (*DeleteTargetOutput, error) 
 DescribeNotificationRule(ctx context.Context, params *DescribeNotificationRuleInput, optFns ...func(*Options)) (*DescribeNotificationRuleOutput, error) 
 ListEventTypes(ctx context.Context, params *ListEventTypesInput, optFns ...func(*Options)) (*ListEventTypesOutput, error) 
 ListNotificationRules(ctx context.Context, params *ListNotificationRulesInput, optFns ...func(*Options)) (*ListNotificationRulesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTargets(ctx context.Context, params *ListTargetsInput, optFns ...func(*Options)) (*ListTargetsOutput, error) 
 Subscribe(ctx context.Context, params *SubscribeInput, optFns ...func(*Options)) (*SubscribeOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 Unsubscribe(ctx context.Context, params *UnsubscribeInput, optFns ...func(*Options)) (*UnsubscribeOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateNotificationRule(ctx context.Context, params *UpdateNotificationRuleInput, optFns ...func(*Options)) (*UpdateNotificationRuleOutput, error) 
}
