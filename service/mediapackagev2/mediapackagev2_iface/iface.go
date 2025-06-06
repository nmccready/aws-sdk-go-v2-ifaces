
package mediapackagev2_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

// IClient defines the interface for mediapackagev2
type IClient interface {
 Options() Options 
 CancelHarvestJob(ctx context.Context, params *CancelHarvestJobInput, optFns ...func(*Options)) (*CancelHarvestJobOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateChannelGroup(ctx context.Context, params *CreateChannelGroupInput, optFns ...func(*Options)) (*CreateChannelGroupOutput, error) 
 CreateHarvestJob(ctx context.Context, params *CreateHarvestJobInput, optFns ...func(*Options)) (*CreateHarvestJobOutput, error) 
 CreateOriginEndpoint(ctx context.Context, params *CreateOriginEndpointInput, optFns ...func(*Options)) (*CreateOriginEndpointOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteChannelGroup(ctx context.Context, params *DeleteChannelGroupInput, optFns ...func(*Options)) (*DeleteChannelGroupOutput, error) 
 DeleteChannelPolicy(ctx context.Context, params *DeleteChannelPolicyInput, optFns ...func(*Options)) (*DeleteChannelPolicyOutput, error) 
 DeleteOriginEndpoint(ctx context.Context, params *DeleteOriginEndpointInput, optFns ...func(*Options)) (*DeleteOriginEndpointOutput, error) 
 DeleteOriginEndpointPolicy(ctx context.Context, params *DeleteOriginEndpointPolicyInput, optFns ...func(*Options)) (*DeleteOriginEndpointPolicyOutput, error) 
 GetChannel(ctx context.Context, params *GetChannelInput, optFns ...func(*Options)) (*GetChannelOutput, error) 
 GetChannelGroup(ctx context.Context, params *GetChannelGroupInput, optFns ...func(*Options)) (*GetChannelGroupOutput, error) 
 GetChannelPolicy(ctx context.Context, params *GetChannelPolicyInput, optFns ...func(*Options)) (*GetChannelPolicyOutput, error) 
 GetHarvestJob(ctx context.Context, params *GetHarvestJobInput, optFns ...func(*Options)) (*GetHarvestJobOutput, error) 
 GetOriginEndpoint(ctx context.Context, params *GetOriginEndpointInput, optFns ...func(*Options)) (*GetOriginEndpointOutput, error) 
 GetOriginEndpointPolicy(ctx context.Context, params *GetOriginEndpointPolicyInput, optFns ...func(*Options)) (*GetOriginEndpointPolicyOutput, error) 
 ListChannelGroups(ctx context.Context, params *ListChannelGroupsInput, optFns ...func(*Options)) (*ListChannelGroupsOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListHarvestJobs(ctx context.Context, params *ListHarvestJobsInput, optFns ...func(*Options)) (*ListHarvestJobsOutput, error) 
 ListOriginEndpoints(ctx context.Context, params *ListOriginEndpointsInput, optFns ...func(*Options)) (*ListOriginEndpointsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutChannelPolicy(ctx context.Context, params *PutChannelPolicyInput, optFns ...func(*Options)) (*PutChannelPolicyOutput, error) 
 PutOriginEndpointPolicy(ctx context.Context, params *PutOriginEndpointPolicyInput, optFns ...func(*Options)) (*PutOriginEndpointPolicyOutput, error) 
 ResetChannelState(ctx context.Context, params *ResetChannelStateInput, optFns ...func(*Options)) (*ResetChannelStateOutput, error) 
 ResetOriginEndpointState(ctx context.Context, params *ResetOriginEndpointStateInput, optFns ...func(*Options)) (*ResetOriginEndpointStateOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateChannelGroup(ctx context.Context, params *UpdateChannelGroupInput, optFns ...func(*Options)) (*UpdateChannelGroupOutput, error) 
 UpdateOriginEndpoint(ctx context.Context, params *UpdateOriginEndpointInput, optFns ...func(*Options)) (*UpdateOriginEndpointOutput, error) 
}
