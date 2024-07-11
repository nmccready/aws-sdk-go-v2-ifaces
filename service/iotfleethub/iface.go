
package iotfleethub

import (
    "github.com/aws/aws-sdk-go-v2/service/iotfleethub"
)

// IClient defines the interface for iotfleethub
type IClient interface {
 Options() Options 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
}
