
package marketplacemetering

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
)

// IMarketplacemetering defines the interface for marketplacemetering
type IMarketplacemetering interface {
 Options() Options 
 BatchMeterUsage(ctx context.Context, params *BatchMeterUsageInput, optFns ...func(*Options)) (*BatchMeterUsageOutput, error) 
 MeterUsage(ctx context.Context, params *MeterUsageInput, optFns ...func(*Options)) (*MeterUsageOutput, error) 
 RegisterUsage(ctx context.Context, params *RegisterUsageInput, optFns ...func(*Options)) (*RegisterUsageOutput, error) 
 ResolveCustomer(ctx context.Context, params *ResolveCustomerInput, optFns ...func(*Options)) (*ResolveCustomerOutput, error) 
}
