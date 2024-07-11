
package sagemakerruntime

import (
    "github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
)

// ISagemakerruntime defines the interface for sagemakerruntime
type ISagemakerruntime interface {
 Options() Options 
 InvokeEndpoint(ctx context.Context, params *InvokeEndpointInput, optFns ...func(*Options)) (*InvokeEndpointOutput, error) 
 InvokeEndpointAsync(ctx context.Context, params *InvokeEndpointAsyncInput, optFns ...func(*Options)) (*InvokeEndpointAsyncOutput, error) 
 InvokeEndpointWithResponseStream(ctx context.Context, params *InvokeEndpointWithResponseStreamInput, optFns ...func(*Options)) (*InvokeEndpointWithResponseStreamOutput, error) 
}
