
package mediapackage

import (
    "github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

// IClient defines the interface for mediapackage
type IClient interface {
 Options() Options 
 ConfigureLogs(ctx context.Context, params *ConfigureLogsInput, optFns ...func(*Options)) (*ConfigureLogsOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateHarvestJob(ctx context.Context, params *CreateHarvestJobInput, optFns ...func(*Options)) (*CreateHarvestJobOutput, error) 
 CreateOriginEndpoint(ctx context.Context, params *CreateOriginEndpointInput, optFns ...func(*Options)) (*CreateOriginEndpointOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteOriginEndpoint(ctx context.Context, params *DeleteOriginEndpointInput, optFns ...func(*Options)) (*DeleteOriginEndpointOutput, error) 
 DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) 
 DescribeHarvestJob(ctx context.Context, params *DescribeHarvestJobInput, optFns ...func(*Options)) (*DescribeHarvestJobOutput, error) 
 DescribeOriginEndpoint(ctx context.Context, params *DescribeOriginEndpointInput, optFns ...func(*Options)) (*DescribeOriginEndpointOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListHarvestJobs(ctx context.Context, params *ListHarvestJobsInput, optFns ...func(*Options)) (*ListHarvestJobsOutput, error) 
 ListOriginEndpoints(ctx context.Context, params *ListOriginEndpointsInput, optFns ...func(*Options)) (*ListOriginEndpointsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RotateChannelCredentials(ctx context.Context, params *RotateChannelCredentialsInput, optFns ...func(*Options)) (*RotateChannelCredentialsOutput, error) 
 RotateIngestEndpointCredentials(ctx context.Context, params *RotateIngestEndpointCredentialsInput, optFns ...func(*Options)) (*RotateIngestEndpointCredentialsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateOriginEndpoint(ctx context.Context, params *UpdateOriginEndpointInput, optFns ...func(*Options)) (*UpdateOriginEndpointOutput, error) 
}
