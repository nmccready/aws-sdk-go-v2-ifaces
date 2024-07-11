
package marketplaceentitlementservice

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
)

// IClient defines the interface for marketplaceentitlementservice
type IClient interface {
 Options() Options 
 GetEntitlements(ctx context.Context, params *GetEntitlementsInput, optFns ...func(*Options)) (*GetEntitlementsOutput, error) 
}
