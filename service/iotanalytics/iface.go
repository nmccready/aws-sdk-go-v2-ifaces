
package iotanalytics

import (
    "github.com/aws/aws-sdk-go-v2/service/iotanalytics"
)

// IIotanalytics defines the interface for iotanalytics
type IIotanalytics interface {
 Options() Options 
 BatchPutMessage(ctx context.Context, params *BatchPutMessageInput, optFns ...func(*Options)) (*BatchPutMessageOutput, error) 
 CancelPipelineReprocessing(ctx context.Context, params *CancelPipelineReprocessingInput, optFns ...func(*Options)) (*CancelPipelineReprocessingOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateDatasetContent(ctx context.Context, params *CreateDatasetContentInput, optFns ...func(*Options)) (*CreateDatasetContentOutput, error) 
 CreateDatastore(ctx context.Context, params *CreateDatastoreInput, optFns ...func(*Options)) (*CreateDatastoreOutput, error) 
 CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteDatasetContent(ctx context.Context, params *DeleteDatasetContentInput, optFns ...func(*Options)) (*DeleteDatasetContentOutput, error) 
 DeleteDatastore(ctx context.Context, params *DeleteDatastoreInput, optFns ...func(*Options)) (*DeleteDatastoreOutput, error) 
 DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) 
 DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeDatastore(ctx context.Context, params *DescribeDatastoreInput, optFns ...func(*Options)) (*DescribeDatastoreOutput, error) 
 DescribeLoggingOptions(ctx context.Context, params *DescribeLoggingOptionsInput, optFns ...func(*Options)) (*DescribeLoggingOptionsOutput, error) 
 DescribePipeline(ctx context.Context, params *DescribePipelineInput, optFns ...func(*Options)) (*DescribePipelineOutput, error) 
 GetDatasetContent(ctx context.Context, params *GetDatasetContentInput, optFns ...func(*Options)) (*GetDatasetContentOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListDatasetContents(ctx context.Context, params *ListDatasetContentsInput, optFns ...func(*Options)) (*ListDatasetContentsOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListDatastores(ctx context.Context, params *ListDatastoresInput, optFns ...func(*Options)) (*ListDatastoresOutput, error) 
 ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutLoggingOptions(ctx context.Context, params *PutLoggingOptionsInput, optFns ...func(*Options)) (*PutLoggingOptionsOutput, error) 
 RunPipelineActivity(ctx context.Context, params *RunPipelineActivityInput, optFns ...func(*Options)) (*RunPipelineActivityOutput, error) 
 SampleChannelData(ctx context.Context, params *SampleChannelDataInput, optFns ...func(*Options)) (*SampleChannelDataOutput, error) 
 StartPipelineReprocessing(ctx context.Context, params *StartPipelineReprocessingInput, optFns ...func(*Options)) (*StartPipelineReprocessingOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) 
 UpdateDatastore(ctx context.Context, params *UpdateDatastoreInput, optFns ...func(*Options)) (*UpdateDatastoreOutput, error) 
 UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) 
}
