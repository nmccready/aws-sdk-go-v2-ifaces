
package kinesisvideo

import (
    "github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
)

// IKinesisvideo defines the interface for kinesisvideo
type IKinesisvideo interface {
 Options() Options 
 CreateSignalingChannel(ctx context.Context, params *CreateSignalingChannelInput, optFns ...func(*Options)) (*CreateSignalingChannelOutput, error) 
 CreateStream(ctx context.Context, params *CreateStreamInput, optFns ...func(*Options)) (*CreateStreamOutput, error) 
 DeleteEdgeConfiguration(ctx context.Context, params *DeleteEdgeConfigurationInput, optFns ...func(*Options)) (*DeleteEdgeConfigurationOutput, error) 
 DeleteSignalingChannel(ctx context.Context, params *DeleteSignalingChannelInput, optFns ...func(*Options)) (*DeleteSignalingChannelOutput, error) 
 DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) 
 DescribeEdgeConfiguration(ctx context.Context, params *DescribeEdgeConfigurationInput, optFns ...func(*Options)) (*DescribeEdgeConfigurationOutput, error) 
 DescribeImageGenerationConfiguration(ctx context.Context, params *DescribeImageGenerationConfigurationInput, optFns ...func(*Options)) (*DescribeImageGenerationConfigurationOutput, error) 
 DescribeMappedResourceConfiguration(ctx context.Context, params *DescribeMappedResourceConfigurationInput, optFns ...func(*Options)) (*DescribeMappedResourceConfigurationOutput, error) 
 DescribeMediaStorageConfiguration(ctx context.Context, params *DescribeMediaStorageConfigurationInput, optFns ...func(*Options)) (*DescribeMediaStorageConfigurationOutput, error) 
 DescribeNotificationConfiguration(ctx context.Context, params *DescribeNotificationConfigurationInput, optFns ...func(*Options)) (*DescribeNotificationConfigurationOutput, error) 
 DescribeSignalingChannel(ctx context.Context, params *DescribeSignalingChannelInput, optFns ...func(*Options)) (*DescribeSignalingChannelOutput, error) 
 DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
 GetDataEndpoint(ctx context.Context, params *GetDataEndpointInput, optFns ...func(*Options)) (*GetDataEndpointOutput, error) 
 GetSignalingChannelEndpoint(ctx context.Context, params *GetSignalingChannelEndpointInput, optFns ...func(*Options)) (*GetSignalingChannelEndpointOutput, error) 
 ListEdgeAgentConfigurations(ctx context.Context, params *ListEdgeAgentConfigurationsInput, optFns ...func(*Options)) (*ListEdgeAgentConfigurationsOutput, error) 
 ListSignalingChannels(ctx context.Context, params *ListSignalingChannelsInput, optFns ...func(*Options)) (*ListSignalingChannelsOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) 
 StartEdgeConfigurationUpdate(ctx context.Context, params *StartEdgeConfigurationUpdateInput, optFns ...func(*Options)) (*StartEdgeConfigurationUpdateOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TagStream(ctx context.Context, params *TagStreamInput, optFns ...func(*Options)) (*TagStreamOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UntagStream(ctx context.Context, params *UntagStreamInput, optFns ...func(*Options)) (*UntagStreamOutput, error) 
 UpdateDataRetention(ctx context.Context, params *UpdateDataRetentionInput, optFns ...func(*Options)) (*UpdateDataRetentionOutput, error) 
 UpdateImageGenerationConfiguration(ctx context.Context, params *UpdateImageGenerationConfigurationInput, optFns ...func(*Options)) (*UpdateImageGenerationConfigurationOutput, error) 
 UpdateMediaStorageConfiguration(ctx context.Context, params *UpdateMediaStorageConfigurationInput, optFns ...func(*Options)) (*UpdateMediaStorageConfigurationOutput, error) 
 UpdateNotificationConfiguration(ctx context.Context, params *UpdateNotificationConfigurationInput, optFns ...func(*Options)) (*UpdateNotificationConfigurationOutput, error) 
 UpdateSignalingChannel(ctx context.Context, params *UpdateSignalingChannelInput, optFns ...func(*Options)) (*UpdateSignalingChannelOutput, error) 
 UpdateStream(ctx context.Context, params *UpdateStreamInput, optFns ...func(*Options)) (*UpdateStreamOutput, error) 
}
