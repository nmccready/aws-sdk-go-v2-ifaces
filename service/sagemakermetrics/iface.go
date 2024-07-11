
package sagemakermetrics

import (
    "github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
)

// ISagemakermetrics defines the interface for sagemakermetrics
type ISagemakermetrics interface {
 Options() Options 
 BatchPutMetrics(ctx context.Context, params *BatchPutMetricsInput, optFns ...func(*Options)) (*BatchPutMetricsOutput, error) 
}
