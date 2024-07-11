
package launchwizard

import (
    "github.com/aws/aws-sdk-go-v2/service/launchwizard"
)

// IClient defines the interface for launchwizard
type IClient interface {
 Options() Options 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 DeleteDeployment(ctx context.Context, params *DeleteDeploymentInput, optFns ...func(*Options)) (*DeleteDeploymentOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetWorkload(ctx context.Context, params *GetWorkloadInput, optFns ...func(*Options)) (*GetWorkloadOutput, error) 
 GetWorkloadDeploymentPattern(ctx context.Context, params *GetWorkloadDeploymentPatternInput, optFns ...func(*Options)) (*GetWorkloadDeploymentPatternOutput, error) 
 ListDeploymentEvents(ctx context.Context, params *ListDeploymentEventsInput, optFns ...func(*Options)) (*ListDeploymentEventsOutput, error) 
 ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkloadDeploymentPatterns(ctx context.Context, params *ListWorkloadDeploymentPatternsInput, optFns ...func(*Options)) (*ListWorkloadDeploymentPatternsOutput, error) 
 ListWorkloads(ctx context.Context, params *ListWorkloadsInput, optFns ...func(*Options)) (*ListWorkloadsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
