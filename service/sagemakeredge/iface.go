
package sagemakeredge

import (
    "github.com/aws/aws-sdk-go-v2/service/sagemakeredge"
)

// IClient defines the interface for sagemakeredge
type IClient interface {
 Options() Options 
 GetDeployments(ctx context.Context, params *GetDeploymentsInput, optFns ...func(*Options)) (*GetDeploymentsOutput, error) 
 GetDeviceRegistration(ctx context.Context, params *GetDeviceRegistrationInput, optFns ...func(*Options)) (*GetDeviceRegistrationOutput, error) 
 SendHeartbeat(ctx context.Context, params *SendHeartbeatInput, optFns ...func(*Options)) (*SendHeartbeatOutput, error) 
}
