
package kinesisvideowebrtcstorage

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
)

// IClient defines the interface for kinesisvideowebrtcstorage
type IClient interface {
 Options() Options 
 JoinStorageSession(ctx context.Context, params *JoinStorageSessionInput, optFns ...func(*Options)) (*JoinStorageSessionOutput, error) 
}
