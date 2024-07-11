
package iotsecuretunneling

import (
    "github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
)

// IIotsecuretunneling defines the interface for iotsecuretunneling
type IIotsecuretunneling interface {
 Options() Options 
 CloseTunnel(ctx context.Context, params *CloseTunnelInput, optFns ...func(*Options)) (*CloseTunnelOutput, error) 
 DescribeTunnel(ctx context.Context, params *DescribeTunnelInput, optFns ...func(*Options)) (*DescribeTunnelOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTunnels(ctx context.Context, params *ListTunnelsInput, optFns ...func(*Options)) (*ListTunnelsOutput, error) 
 OpenTunnel(ctx context.Context, params *OpenTunnelInput, optFns ...func(*Options)) (*OpenTunnelOutput, error) 
 RotateTunnelAccessToken(ctx context.Context, params *RotateTunnelAccessTokenInput, optFns ...func(*Options)) (*RotateTunnelAccessTokenOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
