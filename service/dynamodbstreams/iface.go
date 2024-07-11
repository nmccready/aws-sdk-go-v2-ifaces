
package dynamodbstreams

import (
    "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

// IDynamodbstreams defines the interface for dynamodbstreams
type IDynamodbstreams interface {
 Options() Options 
 DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
 GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) 
 GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
}
