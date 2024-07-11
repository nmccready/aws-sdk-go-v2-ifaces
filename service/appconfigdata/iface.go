
package appconfigdata

import (
    "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
)

// IAppconfigdata defines the interface for appconfigdata
type IAppconfigdata interface {
 Options() Options 
 GetLatestConfiguration(ctx context.Context, params *GetLatestConfigurationInput, optFns ...func(*Options)) (*GetLatestConfigurationOutput, error) 
 StartConfigurationSession(ctx context.Context, params *StartConfigurationSessionInput, optFns ...func(*Options)) (*StartConfigurationSessionOutput, error) 
}
