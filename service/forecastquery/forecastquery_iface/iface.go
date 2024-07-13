
package forecastquery_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/forecastquery"
)

// IClient defines the interface for forecastquery
type IClient interface {
 Options() Options 
 QueryForecast(ctx context.Context, params *QueryForecastInput, optFns ...func(*Options)) (*QueryForecastOutput, error) 
 QueryWhatIfForecast(ctx context.Context, params *QueryWhatIfForecastInput, optFns ...func(*Options)) (*QueryWhatIfForecastOutput, error) 
}