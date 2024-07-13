
package sagemakermetrics_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
)

// IClient defines the interface for sagemakermetrics
type IClient interface {
 Options() Options 
 BatchPutMetrics(ctx context.Context, params *BatchPutMetricsInput, optFns ...func(*Options)) (*BatchPutMetricsOutput, error) 
}