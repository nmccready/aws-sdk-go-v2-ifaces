
package mediastoredata

import (
    "github.com/aws/aws-sdk-go-v2/service/mediastoredata"
)

// IClient defines the interface for mediastoredata
type IClient interface {
 Options() Options 
 DeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*Options)) (*DeleteObjectOutput, error) 
 DescribeObject(ctx context.Context, params *DescribeObjectInput, optFns ...func(*Options)) (*DescribeObjectOutput, error) 
 GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error) 
 ListItems(ctx context.Context, params *ListItemsInput, optFns ...func(*Options)) (*ListItemsOutput, error) 
 PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error) 
}
