
package marketplacedeployment

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplacedeployment"
)

// IMarketplacedeployment defines the interface for marketplacedeployment
type IMarketplacedeployment interface {
 Options() Options 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutDeploymentParameter(ctx context.Context, params *PutDeploymentParameterInput, optFns ...func(*Options)) (*PutDeploymentParameterOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
