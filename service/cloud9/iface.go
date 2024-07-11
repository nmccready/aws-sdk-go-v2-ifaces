
package cloud9

import (
    "github.com/aws/aws-sdk-go-v2/service/cloud9"
)

// IClient defines the interface for cloud9
type IClient interface {
 Options() Options 
 CreateEnvironmentEC2(ctx context.Context, params *CreateEnvironmentEC2Input, optFns ...func(*Options)) (*CreateEnvironmentEC2Output, error) 
 CreateEnvironmentMembership(ctx context.Context, params *CreateEnvironmentMembershipInput, optFns ...func(*Options)) (*CreateEnvironmentMembershipOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 DeleteEnvironmentMembership(ctx context.Context, params *DeleteEnvironmentMembershipInput, optFns ...func(*Options)) (*DeleteEnvironmentMembershipOutput, error) 
 DescribeEnvironmentMemberships(ctx context.Context, params *DescribeEnvironmentMembershipsInput, optFns ...func(*Options)) (*DescribeEnvironmentMembershipsOutput, error) 
 DescribeEnvironmentStatus(ctx context.Context, params *DescribeEnvironmentStatusInput, optFns ...func(*Options)) (*DescribeEnvironmentStatusOutput, error) 
 DescribeEnvironments(ctx context.Context, params *DescribeEnvironmentsInput, optFns ...func(*Options)) (*DescribeEnvironmentsOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
 UpdateEnvironmentMembership(ctx context.Context, params *UpdateEnvironmentMembershipInput, optFns ...func(*Options)) (*UpdateEnvironmentMembershipOutput, error) 
}
