
package qldbsession

import (
    "github.com/aws/aws-sdk-go-v2/service/qldbsession"
)

// IClient defines the interface for qldbsession
type IClient interface {
 Options() Options 
 SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(*Options)) (*SendCommandOutput, error) 
}
