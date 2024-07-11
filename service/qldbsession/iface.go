
package qldbsession

import (
    "github.com/aws/aws-sdk-go-v2/service/qldbsession"
)

// IQldbsession defines the interface for qldbsession
type IQldbsession interface {
 Options() Options 
 SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(*Options)) (*SendCommandOutput, error) 
}
