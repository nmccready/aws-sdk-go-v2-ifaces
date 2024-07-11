
package kinesis

import (
    "github.com/aws/aws-sdk-go-v2/service/kinesis"
)

// IClient defines the interface for kinesis
type IClient interface {
 Options() Options 
 AddTagsToStream(ctx context.Context, params *AddTagsToStreamInput, optFns ...func(*Options)) (*AddTagsToStreamOutput, error) 
 CreateStream(ctx context.Context, params *CreateStreamInput, optFns ...func(*Options)) (*CreateStreamOutput, error) 
 DecreaseStreamRetentionPeriod(ctx context.Context, params *DecreaseStreamRetentionPeriodInput, optFns ...func(*Options)) (*DecreaseStreamRetentionPeriodOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) 
 DeregisterStreamConsumer(ctx context.Context, params *DeregisterStreamConsumerInput, optFns ...func(*Options)) (*DeregisterStreamConsumerOutput, error) 
 DescribeLimits(ctx context.Context, params *DescribeLimitsInput, optFns ...func(*Options)) (*DescribeLimitsOutput, error) 
 DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
 DescribeStreamConsumer(ctx context.Context, params *DescribeStreamConsumerInput, optFns ...func(*Options)) (*DescribeStreamConsumerOutput, error) 
 DescribeStreamSummary(ctx context.Context, params *DescribeStreamSummaryInput, optFns ...func(*Options)) (*DescribeStreamSummaryOutput, error) 
 DisableEnhancedMonitoring(ctx context.Context, params *DisableEnhancedMonitoringInput, optFns ...func(*Options)) (*DisableEnhancedMonitoringOutput, error) 
 EnableEnhancedMonitoring(ctx context.Context, params *EnableEnhancedMonitoringInput, optFns ...func(*Options)) (*EnableEnhancedMonitoringOutput, error) 
 GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error) 
 IncreaseStreamRetentionPeriod(ctx context.Context, params *IncreaseStreamRetentionPeriodInput, optFns ...func(*Options)) (*IncreaseStreamRetentionPeriodOutput, error) 
 ListShards(ctx context.Context, params *ListShardsInput, optFns ...func(*Options)) (*ListShardsOutput, error) 
 ListStreamConsumers(ctx context.Context, params *ListStreamConsumersInput, optFns ...func(*Options)) (*ListStreamConsumersOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
 ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) 
 MergeShards(ctx context.Context, params *MergeShardsInput, optFns ...func(*Options)) (*MergeShardsOutput, error) 
 PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) 
 PutRecords(ctx context.Context, params *PutRecordsInput, optFns ...func(*Options)) (*PutRecordsOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 RegisterStreamConsumer(ctx context.Context, params *RegisterStreamConsumerInput, optFns ...func(*Options)) (*RegisterStreamConsumerOutput, error) 
 RemoveTagsFromStream(ctx context.Context, params *RemoveTagsFromStreamInput, optFns ...func(*Options)) (*RemoveTagsFromStreamOutput, error) 
 SplitShard(ctx context.Context, params *SplitShardInput, optFns ...func(*Options)) (*SplitShardOutput, error) 
 StartStreamEncryption(ctx context.Context, params *StartStreamEncryptionInput, optFns ...func(*Options)) (*StartStreamEncryptionOutput, error) 
 StopStreamEncryption(ctx context.Context, params *StopStreamEncryptionInput, optFns ...func(*Options)) (*StopStreamEncryptionOutput, error) 
 SubscribeToShard(ctx context.Context, params *SubscribeToShardInput, optFns ...func(*Options)) (*SubscribeToShardOutput, error) 
 UpdateShardCount(ctx context.Context, params *UpdateShardCountInput, optFns ...func(*Options)) (*UpdateShardCountOutput, error) 
 UpdateStreamMode(ctx context.Context, params *UpdateStreamModeInput, optFns ...func(*Options)) (*UpdateStreamModeOutput, error) 
 DescribeStream(ctx context.Context, input *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
}
