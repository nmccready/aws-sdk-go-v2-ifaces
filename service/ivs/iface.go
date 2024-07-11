
package ivs

import (
    "github.com/aws/aws-sdk-go-v2/service/ivs"
)

// IClient defines the interface for ivs
type IClient interface {
 Options() Options 
 BatchGetChannel(ctx context.Context, params *BatchGetChannelInput, optFns ...func(*Options)) (*BatchGetChannelOutput, error) 
 BatchGetStreamKey(ctx context.Context, params *BatchGetStreamKeyInput, optFns ...func(*Options)) (*BatchGetStreamKeyOutput, error) 
 BatchStartViewerSessionRevocation(ctx context.Context, params *BatchStartViewerSessionRevocationInput, optFns ...func(*Options)) (*BatchStartViewerSessionRevocationOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreatePlaybackRestrictionPolicy(ctx context.Context, params *CreatePlaybackRestrictionPolicyInput, optFns ...func(*Options)) (*CreatePlaybackRestrictionPolicyOutput, error) 
 CreateRecordingConfiguration(ctx context.Context, params *CreateRecordingConfigurationInput, optFns ...func(*Options)) (*CreateRecordingConfigurationOutput, error) 
 CreateStreamKey(ctx context.Context, params *CreateStreamKeyInput, optFns ...func(*Options)) (*CreateStreamKeyOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeletePlaybackKeyPair(ctx context.Context, params *DeletePlaybackKeyPairInput, optFns ...func(*Options)) (*DeletePlaybackKeyPairOutput, error) 
 DeletePlaybackRestrictionPolicy(ctx context.Context, params *DeletePlaybackRestrictionPolicyInput, optFns ...func(*Options)) (*DeletePlaybackRestrictionPolicyOutput, error) 
 DeleteRecordingConfiguration(ctx context.Context, params *DeleteRecordingConfigurationInput, optFns ...func(*Options)) (*DeleteRecordingConfigurationOutput, error) 
 DeleteStreamKey(ctx context.Context, params *DeleteStreamKeyInput, optFns ...func(*Options)) (*DeleteStreamKeyOutput, error) 
 GetChannel(ctx context.Context, params *GetChannelInput, optFns ...func(*Options)) (*GetChannelOutput, error) 
 GetPlaybackKeyPair(ctx context.Context, params *GetPlaybackKeyPairInput, optFns ...func(*Options)) (*GetPlaybackKeyPairOutput, error) 
 GetPlaybackRestrictionPolicy(ctx context.Context, params *GetPlaybackRestrictionPolicyInput, optFns ...func(*Options)) (*GetPlaybackRestrictionPolicyOutput, error) 
 GetRecordingConfiguration(ctx context.Context, params *GetRecordingConfigurationInput, optFns ...func(*Options)) (*GetRecordingConfigurationOutput, error) 
 GetStream(ctx context.Context, params *GetStreamInput, optFns ...func(*Options)) (*GetStreamOutput, error) 
 GetStreamKey(ctx context.Context, params *GetStreamKeyInput, optFns ...func(*Options)) (*GetStreamKeyOutput, error) 
 GetStreamSession(ctx context.Context, params *GetStreamSessionInput, optFns ...func(*Options)) (*GetStreamSessionOutput, error) 
 ImportPlaybackKeyPair(ctx context.Context, params *ImportPlaybackKeyPairInput, optFns ...func(*Options)) (*ImportPlaybackKeyPairOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListPlaybackKeyPairs(ctx context.Context, params *ListPlaybackKeyPairsInput, optFns ...func(*Options)) (*ListPlaybackKeyPairsOutput, error) 
 ListPlaybackRestrictionPolicies(ctx context.Context, params *ListPlaybackRestrictionPoliciesInput, optFns ...func(*Options)) (*ListPlaybackRestrictionPoliciesOutput, error) 
 ListRecordingConfigurations(ctx context.Context, params *ListRecordingConfigurationsInput, optFns ...func(*Options)) (*ListRecordingConfigurationsOutput, error) 
 ListStreamKeys(ctx context.Context, params *ListStreamKeysInput, optFns ...func(*Options)) (*ListStreamKeysOutput, error) 
 ListStreamSessions(ctx context.Context, params *ListStreamSessionsInput, optFns ...func(*Options)) (*ListStreamSessionsOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutMetadata(ctx context.Context, params *PutMetadataInput, optFns ...func(*Options)) (*PutMetadataOutput, error) 
 StartViewerSessionRevocation(ctx context.Context, params *StartViewerSessionRevocationInput, optFns ...func(*Options)) (*StartViewerSessionRevocationOutput, error) 
 StopStream(ctx context.Context, params *StopStreamInput, optFns ...func(*Options)) (*StopStreamOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdatePlaybackRestrictionPolicy(ctx context.Context, params *UpdatePlaybackRestrictionPolicyInput, optFns ...func(*Options)) (*UpdatePlaybackRestrictionPolicyOutput, error) 
}
