
package appconfigdata

import (
    "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
)

// IClient defines the interface for appconfigdata
type IClient interface {
 Options() Options 
 GetLatestConfiguration(ctx context.Context, params *GetLatestConfigurationInput, optFns ...func(*Options)) (*GetLatestConfigurationOutput, error) 
 StartConfigurationSession(ctx context.Context, params *StartConfigurationSessionInput, optFns ...func(*Options)) (*StartConfigurationSessionOutput, error) 
}
