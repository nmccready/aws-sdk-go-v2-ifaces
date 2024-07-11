
package kinesisvideowebrtcstorage

import (
    "github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
)

// IKinesisvideowebrtcstorage defines the interface for kinesisvideowebrtcstorage
type IKinesisvideowebrtcstorage interface {
 Options() Options 
 JoinStorageSession(ctx context.Context, params *JoinStorageSessionInput, optFns ...func(*Options)) (*JoinStorageSessionOutput, error) 
}
