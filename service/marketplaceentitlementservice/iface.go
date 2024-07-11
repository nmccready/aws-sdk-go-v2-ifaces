
package marketplaceentitlementservice

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
)

// IMarketplaceentitlementservice defines the interface for marketplaceentitlementservice
type IMarketplaceentitlementservice interface {
 Options() Options 
 GetEntitlements(ctx context.Context, params *GetEntitlementsInput, optFns ...func(*Options)) (*GetEntitlementsOutput, error) 
}
