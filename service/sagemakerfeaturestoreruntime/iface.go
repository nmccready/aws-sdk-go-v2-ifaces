
package sagemakerfeaturestoreruntime

import (
    "github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
)

// ISagemakerfeaturestoreruntime defines the interface for sagemakerfeaturestoreruntime
type ISagemakerfeaturestoreruntime interface {
 Options() Options 
 BatchGetRecord(ctx context.Context, params *BatchGetRecordInput, optFns ...func(*Options)) (*BatchGetRecordOutput, error) 
 DeleteRecord(ctx context.Context, params *DeleteRecordInput, optFns ...func(*Options)) (*DeleteRecordOutput, error) 
 GetRecord(ctx context.Context, params *GetRecordInput, optFns ...func(*Options)) (*GetRecordOutput, error) 
 PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) 
}
