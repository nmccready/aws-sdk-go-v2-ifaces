
package eksauth

import (
    "github.com/aws/aws-sdk-go-v2/service/eksauth"
)

// IClient defines the interface for eksauth
type IClient interface {
 Options() Options 
 AssumeRoleForPodIdentity(ctx context.Context, params *AssumeRoleForPodIdentityInput, optFns ...func(*Options)) (*AssumeRoleForPodIdentityOutput, error) 
}
