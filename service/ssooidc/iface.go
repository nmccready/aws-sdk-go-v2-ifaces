
package ssooidc

import (
    "github.com/aws/aws-sdk-go-v2/service/ssooidc"
)

// ISsooidc defines the interface for ssooidc
type ISsooidc interface {
 Options() Options 
 CreateToken(ctx context.Context, params *CreateTokenInput, optFns ...func(*Options)) (*CreateTokenOutput, error) 
 CreateTokenWithIAM(ctx context.Context, params *CreateTokenWithIAMInput, optFns ...func(*Options)) (*CreateTokenWithIAMOutput, error) 
 RegisterClient(ctx context.Context, params *RegisterClientInput, optFns ...func(*Options)) (*RegisterClientOutput, error) 
 StartDeviceAuthorization(ctx context.Context, params *StartDeviceAuthorizationInput, optFns ...func(*Options)) (*StartDeviceAuthorizationOutput, error) 
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
}
