
package mwaa

import (
    "github.com/aws/aws-sdk-go-v2/service/mwaa"
)

// IMwaa defines the interface for mwaa
type IMwaa interface {
 Options() Options 
 CreateCliToken(ctx context.Context, params *CreateCliTokenInput, optFns ...func(*Options)) (*CreateCliTokenOutput, error) 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 CreateWebLoginToken(ctx context.Context, params *CreateWebLoginTokenInput, optFns ...func(*Options)) (*CreateWebLoginTokenOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PublishMetrics(ctx context.Context, params *PublishMetricsInput, optFns ...func(*Options)) (*PublishMetricsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
}
