
package sts

import (
    "github.com/aws/aws-sdk-go-v2/service/sts"
)

// IClient defines the interface for sts
type IClient interface {
 Options() Options 
 AssumeRole(ctx context.Context, params *AssumeRoleInput, optFns ...func(*Options)) (*AssumeRoleOutput, error) 
 PresignAssumeRole(ctx context.Context, params *AssumeRoleInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 AssumeRoleWithSAML(ctx context.Context, params *AssumeRoleWithSAMLInput, optFns ...func(*Options)) (*AssumeRoleWithSAMLOutput, error) 
 AssumeRoleWithWebIdentity(ctx context.Context, params *AssumeRoleWithWebIdentityInput, optFns ...func(*Options)) (*AssumeRoleWithWebIdentityOutput, error) 
 DecodeAuthorizationMessage(ctx context.Context, params *DecodeAuthorizationMessageInput, optFns ...func(*Options)) (*DecodeAuthorizationMessageOutput, error) 
 GetAccessKeyInfo(ctx context.Context, params *GetAccessKeyInfoInput, optFns ...func(*Options)) (*GetAccessKeyInfoOutput, error) 
 GetCallerIdentity(ctx context.Context, params *GetCallerIdentityInput, optFns ...func(*Options)) (*GetCallerIdentityOutput, error) 
 PresignGetCallerIdentity(ctx context.Context, params *GetCallerIdentityInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
 GetFederationToken(ctx context.Context, params *GetFederationTokenInput, optFns ...func(*Options)) (*GetFederationTokenOutput, error) 
 GetSessionToken(ctx context.Context, params *GetSessionTokenInput, optFns ...func(*Options)) (*GetSessionTokenOutput, error) 
}
