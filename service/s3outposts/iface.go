
package s3outposts

import (
    "github.com/aws/aws-sdk-go-v2/service/s3outposts"
)

// IClient defines the interface for s3outposts
type IClient interface {
 Options() Options 
 CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) 
 DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) 
 ListEndpoints(ctx context.Context, params *ListEndpointsInput, optFns ...func(*Options)) (*ListEndpointsOutput, error) 
 ListOutpostsWithS3(ctx context.Context, params *ListOutpostsWithS3Input, optFns ...func(*Options)) (*ListOutpostsWithS3Output, error) 
 ListSharedEndpoints(ctx context.Context, params *ListSharedEndpointsInput, optFns ...func(*Options)) (*ListSharedEndpointsOutput, error) 
}
