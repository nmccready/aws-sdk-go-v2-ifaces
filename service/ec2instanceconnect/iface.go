
package ec2instanceconnect

import (
    "github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
)

// IClient defines the interface for ec2instanceconnect
type IClient interface {
 Options() Options 
 SendSSHPublicKey(ctx context.Context, params *SendSSHPublicKeyInput, optFns ...func(*Options)) (*SendSSHPublicKeyOutput, error) 
 SendSerialConsoleSSHPublicKey(ctx context.Context, params *SendSerialConsoleSSHPublicKeyInput, optFns ...func(*Options)) (*SendSerialConsoleSSHPublicKeyOutput, error) 
}
