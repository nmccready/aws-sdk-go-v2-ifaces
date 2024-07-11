
package workmailmessageflow

import (
    "github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
)

// IClient defines the interface for workmailmessageflow
type IClient interface {
 Options() Options 
 GetRawMessageContent(ctx context.Context, params *GetRawMessageContentInput, optFns ...func(*Options)) (*GetRawMessageContentOutput, error) 
 PutRawMessageContent(ctx context.Context, params *PutRawMessageContentInput, optFns ...func(*Options)) (*PutRawMessageContentOutput, error) 
}
