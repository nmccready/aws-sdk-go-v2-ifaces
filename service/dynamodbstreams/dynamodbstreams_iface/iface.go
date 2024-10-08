
package dynamodbstreams_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

// IClient defines the interface for dynamodbstreams
type IClient interface {
 Options() Options 
 DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
 GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) 
 GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
}
