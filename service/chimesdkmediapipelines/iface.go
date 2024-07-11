
package chimesdkmediapipelines

import (
    "github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines"
)

// IClient defines the interface for chimesdkmediapipelines
type IClient interface {
 Options() Options 
 CreateMediaCapturePipeline(ctx context.Context, params *CreateMediaCapturePipelineInput, optFns ...func(*Options)) (*CreateMediaCapturePipelineOutput, error) 
 CreateMediaConcatenationPipeline(ctx context.Context, params *CreateMediaConcatenationPipelineInput, optFns ...func(*Options)) (*CreateMediaConcatenationPipelineOutput, error) 
 CreateMediaInsightsPipeline(ctx context.Context, params *CreateMediaInsightsPipelineInput, optFns ...func(*Options)) (*CreateMediaInsightsPipelineOutput, error) 
 CreateMediaInsightsPipelineConfiguration(ctx context.Context, params *CreateMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*CreateMediaInsightsPipelineConfigurationOutput, error) 
 CreateMediaLiveConnectorPipeline(ctx context.Context, params *CreateMediaLiveConnectorPipelineInput, optFns ...func(*Options)) (*CreateMediaLiveConnectorPipelineOutput, error) 
 CreateMediaPipelineKinesisVideoStreamPool(ctx context.Context, params *CreateMediaPipelineKinesisVideoStreamPoolInput, optFns ...func(*Options)) (*CreateMediaPipelineKinesisVideoStreamPoolOutput, error) 
 CreateMediaStreamPipeline(ctx context.Context, params *CreateMediaStreamPipelineInput, optFns ...func(*Options)) (*CreateMediaStreamPipelineOutput, error) 
 DeleteMediaCapturePipeline(ctx context.Context, params *DeleteMediaCapturePipelineInput, optFns ...func(*Options)) (*DeleteMediaCapturePipelineOutput, error) 
 DeleteMediaInsightsPipelineConfiguration(ctx context.Context, params *DeleteMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*DeleteMediaInsightsPipelineConfigurationOutput, error) 
 DeleteMediaPipeline(ctx context.Context, params *DeleteMediaPipelineInput, optFns ...func(*Options)) (*DeleteMediaPipelineOutput, error) 
 DeleteMediaPipelineKinesisVideoStreamPool(ctx context.Context, params *DeleteMediaPipelineKinesisVideoStreamPoolInput, optFns ...func(*Options)) (*DeleteMediaPipelineKinesisVideoStreamPoolOutput, error) 
 GetMediaCapturePipeline(ctx context.Context, params *GetMediaCapturePipelineInput, optFns ...func(*Options)) (*GetMediaCapturePipelineOutput, error) 
 GetMediaInsightsPipelineConfiguration(ctx context.Context, params *GetMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*GetMediaInsightsPipelineConfigurationOutput, error) 
 GetMediaPipeline(ctx context.Context, params *GetMediaPipelineInput, optFns ...func(*Options)) (*GetMediaPipelineOutput, error) 
 GetMediaPipelineKinesisVideoStreamPool(ctx context.Context, params *GetMediaPipelineKinesisVideoStreamPoolInput, optFns ...func(*Options)) (*GetMediaPipelineKinesisVideoStreamPoolOutput, error) 
 GetSpeakerSearchTask(ctx context.Context, params *GetSpeakerSearchTaskInput, optFns ...func(*Options)) (*GetSpeakerSearchTaskOutput, error) 
 GetVoiceToneAnalysisTask(ctx context.Context, params *GetVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*GetVoiceToneAnalysisTaskOutput, error) 
 ListMediaCapturePipelines(ctx context.Context, params *ListMediaCapturePipelinesInput, optFns ...func(*Options)) (*ListMediaCapturePipelinesOutput, error) 
 ListMediaInsightsPipelineConfigurations(ctx context.Context, params *ListMediaInsightsPipelineConfigurationsInput, optFns ...func(*Options)) (*ListMediaInsightsPipelineConfigurationsOutput, error) 
 ListMediaPipelineKinesisVideoStreamPools(ctx context.Context, params *ListMediaPipelineKinesisVideoStreamPoolsInput, optFns ...func(*Options)) (*ListMediaPipelineKinesisVideoStreamPoolsOutput, error) 
 ListMediaPipelines(ctx context.Context, params *ListMediaPipelinesInput, optFns ...func(*Options)) (*ListMediaPipelinesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartSpeakerSearchTask(ctx context.Context, params *StartSpeakerSearchTaskInput, optFns ...func(*Options)) (*StartSpeakerSearchTaskOutput, error) 
 StartVoiceToneAnalysisTask(ctx context.Context, params *StartVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*StartVoiceToneAnalysisTaskOutput, error) 
 StopSpeakerSearchTask(ctx context.Context, params *StopSpeakerSearchTaskInput, optFns ...func(*Options)) (*StopSpeakerSearchTaskOutput, error) 
 StopVoiceToneAnalysisTask(ctx context.Context, params *StopVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*StopVoiceToneAnalysisTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateMediaInsightsPipelineConfiguration(ctx context.Context, params *UpdateMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*UpdateMediaInsightsPipelineConfigurationOutput, error) 
 UpdateMediaInsightsPipelineStatus(ctx context.Context, params *UpdateMediaInsightsPipelineStatusInput, optFns ...func(*Options)) (*UpdateMediaInsightsPipelineStatusOutput, error) 
 UpdateMediaPipelineKinesisVideoStreamPool(ctx context.Context, params *UpdateMediaPipelineKinesisVideoStreamPoolInput, optFns ...func(*Options)) (*UpdateMediaPipelineKinesisVideoStreamPoolOutput, error) 
}
