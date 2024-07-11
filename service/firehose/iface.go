
package firehose

import (
    "github.com/aws/aws-sdk-go-v2/service/firehose"
)

// IClient defines the interface for firehose
type IClient interface {
 Options() Options 
 CreateDeliveryStream(ctx context.Context, params *CreateDeliveryStreamInput, optFns ...func(*Options)) (*CreateDeliveryStreamOutput, error) 
 DeleteDeliveryStream(ctx context.Context, params *DeleteDeliveryStreamInput, optFns ...func(*Options)) (*DeleteDeliveryStreamOutput, error) 
 DescribeDeliveryStream(ctx context.Context, params *DescribeDeliveryStreamInput, optFns ...func(*Options)) (*DescribeDeliveryStreamOutput, error) 
 ListDeliveryStreams(ctx context.Context, params *ListDeliveryStreamsInput, optFns ...func(*Options)) (*ListDeliveryStreamsOutput, error) 
 ListTagsForDeliveryStream(ctx context.Context, params *ListTagsForDeliveryStreamInput, optFns ...func(*Options)) (*ListTagsForDeliveryStreamOutput, error) 
 PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) 
 PutRecordBatch(ctx context.Context, params *PutRecordBatchInput, optFns ...func(*Options)) (*PutRecordBatchOutput, error) 
 StartDeliveryStreamEncryption(ctx context.Context, params *StartDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StartDeliveryStreamEncryptionOutput, error) 
 StopDeliveryStreamEncryption(ctx context.Context, params *StopDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StopDeliveryStreamEncryptionOutput, error) 
 TagDeliveryStream(ctx context.Context, params *TagDeliveryStreamInput, optFns ...func(*Options)) (*TagDeliveryStreamOutput, error) 
 UntagDeliveryStream(ctx context.Context, params *UntagDeliveryStreamInput, optFns ...func(*Options)) (*UntagDeliveryStreamOutput, error) 
 UpdateDestination(ctx context.Context, params *UpdateDestinationInput, optFns ...func(*Options)) (*UpdateDestinationOutput, error) 
}
