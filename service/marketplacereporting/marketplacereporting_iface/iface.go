
package marketplacereporting_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/marketplacereporting"
)

// IClient defines the interface for marketplacereporting
type IClient interface {
 Options() Options 
 GetBuyerDashboard(ctx context.Context, params *GetBuyerDashboardInput, optFns ...func(*Options)) (*GetBuyerDashboardOutput, error) 
}