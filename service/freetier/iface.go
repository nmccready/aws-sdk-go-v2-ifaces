
package freetier

import (
    "github.com/aws/aws-sdk-go-v2/service/freetier"
)

// IClient defines the interface for freetier
type IClient interface {
 Options() Options 
 GetFreeTierUsage(ctx context.Context, params *GetFreeTierUsageInput, optFns ...func(*Options)) (*GetFreeTierUsageOutput, error) 
}
