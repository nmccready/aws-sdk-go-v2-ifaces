
package mediastore

import (
    "github.com/aws/aws-sdk-go-v2/service/mediastore"
)

// IMediastore defines the interface for mediastore
type IMediastore interface {
 Options() Options 
 CreateContainer(ctx context.Context, params *CreateContainerInput, optFns ...func(*Options)) (*CreateContainerOutput, error) 
 DeleteContainer(ctx context.Context, params *DeleteContainerInput, optFns ...func(*Options)) (*DeleteContainerOutput, error) 
 DeleteContainerPolicy(ctx context.Context, params *DeleteContainerPolicyInput, optFns ...func(*Options)) (*DeleteContainerPolicyOutput, error) 
 DeleteCorsPolicy(ctx context.Context, params *DeleteCorsPolicyInput, optFns ...func(*Options)) (*DeleteCorsPolicyOutput, error) 
 DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) 
 DeleteMetricPolicy(ctx context.Context, params *DeleteMetricPolicyInput, optFns ...func(*Options)) (*DeleteMetricPolicyOutput, error) 
 DescribeContainer(ctx context.Context, params *DescribeContainerInput, optFns ...func(*Options)) (*DescribeContainerOutput, error) 
 GetContainerPolicy(ctx context.Context, params *GetContainerPolicyInput, optFns ...func(*Options)) (*GetContainerPolicyOutput, error) 
 GetCorsPolicy(ctx context.Context, params *GetCorsPolicyInput, optFns ...func(*Options)) (*GetCorsPolicyOutput, error) 
 GetLifecyclePolicy(ctx context.Context, params *GetLifecyclePolicyInput, optFns ...func(*Options)) (*GetLifecyclePolicyOutput, error) 
 GetMetricPolicy(ctx context.Context, params *GetMetricPolicyInput, optFns ...func(*Options)) (*GetMetricPolicyOutput, error) 
 ListContainers(ctx context.Context, params *ListContainersInput, optFns ...func(*Options)) (*ListContainersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutContainerPolicy(ctx context.Context, params *PutContainerPolicyInput, optFns ...func(*Options)) (*PutContainerPolicyOutput, error) 
 PutCorsPolicy(ctx context.Context, params *PutCorsPolicyInput, optFns ...func(*Options)) (*PutCorsPolicyOutput, error) 
 PutLifecyclePolicy(ctx context.Context, params *PutLifecyclePolicyInput, optFns ...func(*Options)) (*PutLifecyclePolicyOutput, error) 
 PutMetricPolicy(ctx context.Context, params *PutMetricPolicyInput, optFns ...func(*Options)) (*PutMetricPolicyOutput, error) 
 StartAccessLogging(ctx context.Context, params *StartAccessLoggingInput, optFns ...func(*Options)) (*StartAccessLoggingOutput, error) 
 StopAccessLogging(ctx context.Context, params *StopAccessLoggingInput, optFns ...func(*Options)) (*StopAccessLoggingOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
