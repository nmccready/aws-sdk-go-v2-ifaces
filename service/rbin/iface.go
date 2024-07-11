
package rbin

import (
    "github.com/aws/aws-sdk-go-v2/service/rbin"
)

// IClient defines the interface for rbin
type IClient interface {
 Options() Options 
 CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 GetRule(ctx context.Context, params *GetRuleInput, optFns ...func(*Options)) (*GetRuleOutput, error) 
 ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 LockRule(ctx context.Context, params *LockRuleInput, optFns ...func(*Options)) (*LockRuleOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UnlockRule(ctx context.Context, params *UnlockRuleInput, optFns ...func(*Options)) (*UnlockRuleOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) 
}
