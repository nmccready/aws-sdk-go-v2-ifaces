
package eksauth

import (
    "github.com/aws/aws-sdk-go-v2/service/eksauth"
)

// IEksauth defines the interface for eksauth
type IEksauth interface {
 Options() Options 
 AssumeRoleForPodIdentity(ctx context.Context, params *AssumeRoleForPodIdentityInput, optFns ...func(*Options)) (*AssumeRoleForPodIdentityOutput, error) 
}
