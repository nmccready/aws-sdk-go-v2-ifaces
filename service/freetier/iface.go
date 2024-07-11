
package freetier

import (
    "github.com/aws/aws-sdk-go-v2/service/freetier"
)

// IFreetier defines the interface for freetier
type IFreetier interface {
 Options() Options 
 GetFreeTierUsage(ctx context.Context, params *GetFreeTierUsageInput, optFns ...func(*Options)) (*GetFreeTierUsageOutput, error) 
}
