
package dlm

import (
    "github.com/aws/aws-sdk-go-v2/service/dlm"
)

// IClient defines the interface for dlm
type IClient interface {
 Options() Options 
 CreateLifecyclePolicy(ctx context.Context, params *CreateLifecyclePolicyInput, optFns ...func(*Options)) (*CreateLifecyclePolicyOutput, error) 
 DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) 
 GetLifecyclePolicies(ctx context.Context, params *GetLifecyclePoliciesInput, optFns ...func(*Options)) (*GetLifecyclePoliciesOutput, error) 
 GetLifecyclePolicy(ctx context.Context, params *GetLifecyclePolicyInput, optFns ...func(*Options)) (*GetLifecyclePolicyOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLifecyclePolicy(ctx context.Context, params *UpdateLifecyclePolicyInput, optFns ...func(*Options)) (*UpdateLifecyclePolicyOutput, error) 
}
