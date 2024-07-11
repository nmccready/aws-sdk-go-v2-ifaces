
package opensearchserverless

import (
    "github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
)

// IClient defines the interface for opensearchserverless
type IClient interface {
 Options() Options 
 BatchGetCollection(ctx context.Context, params *BatchGetCollectionInput, optFns ...func(*Options)) (*BatchGetCollectionOutput, error) 
 BatchGetEffectiveLifecyclePolicy(ctx context.Context, params *BatchGetEffectiveLifecyclePolicyInput, optFns ...func(*Options)) (*BatchGetEffectiveLifecyclePolicyOutput, error) 
 BatchGetLifecyclePolicy(ctx context.Context, params *BatchGetLifecyclePolicyInput, optFns ...func(*Options)) (*BatchGetLifecyclePolicyOutput, error) 
 BatchGetVpcEndpoint(ctx context.Context, params *BatchGetVpcEndpointInput, optFns ...func(*Options)) (*BatchGetVpcEndpointOutput, error) 
 CreateAccessPolicy(ctx context.Context, params *CreateAccessPolicyInput, optFns ...func(*Options)) (*CreateAccessPolicyOutput, error) 
 CreateCollection(ctx context.Context, params *CreateCollectionInput, optFns ...func(*Options)) (*CreateCollectionOutput, error) 
 CreateLifecyclePolicy(ctx context.Context, params *CreateLifecyclePolicyInput, optFns ...func(*Options)) (*CreateLifecyclePolicyOutput, error) 
 CreateSecurityConfig(ctx context.Context, params *CreateSecurityConfigInput, optFns ...func(*Options)) (*CreateSecurityConfigOutput, error) 
 CreateSecurityPolicy(ctx context.Context, params *CreateSecurityPolicyInput, optFns ...func(*Options)) (*CreateSecurityPolicyOutput, error) 
 CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) 
 DeleteAccessPolicy(ctx context.Context, params *DeleteAccessPolicyInput, optFns ...func(*Options)) (*DeleteAccessPolicyOutput, error) 
 DeleteCollection(ctx context.Context, params *DeleteCollectionInput, optFns ...func(*Options)) (*DeleteCollectionOutput, error) 
 DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) 
 DeleteSecurityConfig(ctx context.Context, params *DeleteSecurityConfigInput, optFns ...func(*Options)) (*DeleteSecurityConfigOutput, error) 
 DeleteSecurityPolicy(ctx context.Context, params *DeleteSecurityPolicyInput, optFns ...func(*Options)) (*DeleteSecurityPolicyOutput, error) 
 DeleteVpcEndpoint(ctx context.Context, params *DeleteVpcEndpointInput, optFns ...func(*Options)) (*DeleteVpcEndpointOutput, error) 
 GetAccessPolicy(ctx context.Context, params *GetAccessPolicyInput, optFns ...func(*Options)) (*GetAccessPolicyOutput, error) 
 GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) 
 GetPoliciesStats(ctx context.Context, params *GetPoliciesStatsInput, optFns ...func(*Options)) (*GetPoliciesStatsOutput, error) 
 GetSecurityConfig(ctx context.Context, params *GetSecurityConfigInput, optFns ...func(*Options)) (*GetSecurityConfigOutput, error) 
 GetSecurityPolicy(ctx context.Context, params *GetSecurityPolicyInput, optFns ...func(*Options)) (*GetSecurityPolicyOutput, error) 
 ListAccessPolicies(ctx context.Context, params *ListAccessPoliciesInput, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) 
 ListCollections(ctx context.Context, params *ListCollectionsInput, optFns ...func(*Options)) (*ListCollectionsOutput, error) 
 ListLifecyclePolicies(ctx context.Context, params *ListLifecyclePoliciesInput, optFns ...func(*Options)) (*ListLifecyclePoliciesOutput, error) 
 ListSecurityConfigs(ctx context.Context, params *ListSecurityConfigsInput, optFns ...func(*Options)) (*ListSecurityConfigsOutput, error) 
 ListSecurityPolicies(ctx context.Context, params *ListSecurityPoliciesInput, optFns ...func(*Options)) (*ListSecurityPoliciesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVpcEndpoints(ctx context.Context, params *ListVpcEndpointsInput, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccessPolicy(ctx context.Context, params *UpdateAccessPolicyInput, optFns ...func(*Options)) (*UpdateAccessPolicyOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateCollection(ctx context.Context, params *UpdateCollectionInput, optFns ...func(*Options)) (*UpdateCollectionOutput, error) 
 UpdateLifecyclePolicy(ctx context.Context, params *UpdateLifecyclePolicyInput, optFns ...func(*Options)) (*UpdateLifecyclePolicyOutput, error) 
 UpdateSecurityConfig(ctx context.Context, params *UpdateSecurityConfigInput, optFns ...func(*Options)) (*UpdateSecurityConfigOutput, error) 
 UpdateSecurityPolicy(ctx context.Context, params *UpdateSecurityPolicyInput, optFns ...func(*Options)) (*UpdateSecurityPolicyOutput, error) 
 UpdateVpcEndpoint(ctx context.Context, params *UpdateVpcEndpointInput, optFns ...func(*Options)) (*UpdateVpcEndpointOutput, error) 
}
