
package mediatailor_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/mediatailor"
)

// IClient defines the interface for mediatailor
type IClient interface {
 Options() Options 
 ConfigureLogsForChannel(ctx context.Context, params *ConfigureLogsForChannelInput, optFns ...func(*Options)) (*ConfigureLogsForChannelOutput, error) 
 ConfigureLogsForPlaybackConfiguration(ctx context.Context, params *ConfigureLogsForPlaybackConfigurationInput, optFns ...func(*Options)) (*ConfigureLogsForPlaybackConfigurationOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateLiveSource(ctx context.Context, params *CreateLiveSourceInput, optFns ...func(*Options)) (*CreateLiveSourceOutput, error) 
 CreatePrefetchSchedule(ctx context.Context, params *CreatePrefetchScheduleInput, optFns ...func(*Options)) (*CreatePrefetchScheduleOutput, error) 
 CreateProgram(ctx context.Context, params *CreateProgramInput, optFns ...func(*Options)) (*CreateProgramOutput, error) 
 CreateSourceLocation(ctx context.Context, params *CreateSourceLocationInput, optFns ...func(*Options)) (*CreateSourceLocationOutput, error) 
 CreateVodSource(ctx context.Context, params *CreateVodSourceInput, optFns ...func(*Options)) (*CreateVodSourceOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteChannelPolicy(ctx context.Context, params *DeleteChannelPolicyInput, optFns ...func(*Options)) (*DeleteChannelPolicyOutput, error) 
 DeleteLiveSource(ctx context.Context, params *DeleteLiveSourceInput, optFns ...func(*Options)) (*DeleteLiveSourceOutput, error) 
 DeletePlaybackConfiguration(ctx context.Context, params *DeletePlaybackConfigurationInput, optFns ...func(*Options)) (*DeletePlaybackConfigurationOutput, error) 
 DeletePrefetchSchedule(ctx context.Context, params *DeletePrefetchScheduleInput, optFns ...func(*Options)) (*DeletePrefetchScheduleOutput, error) 
 DeleteProgram(ctx context.Context, params *DeleteProgramInput, optFns ...func(*Options)) (*DeleteProgramOutput, error) 
 DeleteSourceLocation(ctx context.Context, params *DeleteSourceLocationInput, optFns ...func(*Options)) (*DeleteSourceLocationOutput, error) 
 DeleteVodSource(ctx context.Context, params *DeleteVodSourceInput, optFns ...func(*Options)) (*DeleteVodSourceOutput, error) 
 DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) 
 DescribeLiveSource(ctx context.Context, params *DescribeLiveSourceInput, optFns ...func(*Options)) (*DescribeLiveSourceOutput, error) 
 DescribeProgram(ctx context.Context, params *DescribeProgramInput, optFns ...func(*Options)) (*DescribeProgramOutput, error) 
 DescribeSourceLocation(ctx context.Context, params *DescribeSourceLocationInput, optFns ...func(*Options)) (*DescribeSourceLocationOutput, error) 
 DescribeVodSource(ctx context.Context, params *DescribeVodSourceInput, optFns ...func(*Options)) (*DescribeVodSourceOutput, error) 
 GetChannelPolicy(ctx context.Context, params *GetChannelPolicyInput, optFns ...func(*Options)) (*GetChannelPolicyOutput, error) 
 GetChannelSchedule(ctx context.Context, params *GetChannelScheduleInput, optFns ...func(*Options)) (*GetChannelScheduleOutput, error) 
 GetPlaybackConfiguration(ctx context.Context, params *GetPlaybackConfigurationInput, optFns ...func(*Options)) (*GetPlaybackConfigurationOutput, error) 
 GetPrefetchSchedule(ctx context.Context, params *GetPrefetchScheduleInput, optFns ...func(*Options)) (*GetPrefetchScheduleOutput, error) 
 ListAlerts(ctx context.Context, params *ListAlertsInput, optFns ...func(*Options)) (*ListAlertsOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListLiveSources(ctx context.Context, params *ListLiveSourcesInput, optFns ...func(*Options)) (*ListLiveSourcesOutput, error) 
 ListPlaybackConfigurations(ctx context.Context, params *ListPlaybackConfigurationsInput, optFns ...func(*Options)) (*ListPlaybackConfigurationsOutput, error) 
 ListPrefetchSchedules(ctx context.Context, params *ListPrefetchSchedulesInput, optFns ...func(*Options)) (*ListPrefetchSchedulesOutput, error) 
 ListSourceLocations(ctx context.Context, params *ListSourceLocationsInput, optFns ...func(*Options)) (*ListSourceLocationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVodSources(ctx context.Context, params *ListVodSourcesInput, optFns ...func(*Options)) (*ListVodSourcesOutput, error) 
 PutChannelPolicy(ctx context.Context, params *PutChannelPolicyInput, optFns ...func(*Options)) (*PutChannelPolicyOutput, error) 
 PutPlaybackConfiguration(ctx context.Context, params *PutPlaybackConfigurationInput, optFns ...func(*Options)) (*PutPlaybackConfigurationOutput, error) 
 StartChannel(ctx context.Context, params *StartChannelInput, optFns ...func(*Options)) (*StartChannelOutput, error) 
 StopChannel(ctx context.Context, params *StopChannelInput, optFns ...func(*Options)) (*StopChannelOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateLiveSource(ctx context.Context, params *UpdateLiveSourceInput, optFns ...func(*Options)) (*UpdateLiveSourceOutput, error) 
 UpdateProgram(ctx context.Context, params *UpdateProgramInput, optFns ...func(*Options)) (*UpdateProgramOutput, error) 
 UpdateSourceLocation(ctx context.Context, params *UpdateSourceLocationInput, optFns ...func(*Options)) (*UpdateSourceLocationOutput, error) 
 UpdateVodSource(ctx context.Context, params *UpdateVodSourceInput, optFns ...func(*Options)) (*UpdateVodSourceOutput, error) 
}
