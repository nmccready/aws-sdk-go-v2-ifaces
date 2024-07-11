
package synthetics

import (
    "github.com/aws/aws-sdk-go-v2/service/synthetics"
)

// IClient defines the interface for synthetics
type IClient interface {
 Options() Options 
 AssociateResource(ctx context.Context, params *AssociateResourceInput, optFns ...func(*Options)) (*AssociateResourceOutput, error) 
 CreateCanary(ctx context.Context, params *CreateCanaryInput, optFns ...func(*Options)) (*CreateCanaryOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 DeleteCanary(ctx context.Context, params *DeleteCanaryInput, optFns ...func(*Options)) (*DeleteCanaryOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DescribeCanaries(ctx context.Context, params *DescribeCanariesInput, optFns ...func(*Options)) (*DescribeCanariesOutput, error) 
 DescribeCanariesLastRun(ctx context.Context, params *DescribeCanariesLastRunInput, optFns ...func(*Options)) (*DescribeCanariesLastRunOutput, error) 
 DescribeRuntimeVersions(ctx context.Context, params *DescribeRuntimeVersionsInput, optFns ...func(*Options)) (*DescribeRuntimeVersionsOutput, error) 
 DisassociateResource(ctx context.Context, params *DisassociateResourceInput, optFns ...func(*Options)) (*DisassociateResourceOutput, error) 
 GetCanary(ctx context.Context, params *GetCanaryInput, optFns ...func(*Options)) (*GetCanaryOutput, error) 
 GetCanaryRuns(ctx context.Context, params *GetCanaryRunsInput, optFns ...func(*Options)) (*GetCanaryRunsOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 ListAssociatedGroups(ctx context.Context, params *ListAssociatedGroupsInput, optFns ...func(*Options)) (*ListAssociatedGroupsOutput, error) 
 ListGroupResources(ctx context.Context, params *ListGroupResourcesInput, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartCanary(ctx context.Context, params *StartCanaryInput, optFns ...func(*Options)) (*StartCanaryOutput, error) 
 StopCanary(ctx context.Context, params *StopCanaryInput, optFns ...func(*Options)) (*StopCanaryOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCanary(ctx context.Context, params *UpdateCanaryInput, optFns ...func(*Options)) (*UpdateCanaryOutput, error) 
}
