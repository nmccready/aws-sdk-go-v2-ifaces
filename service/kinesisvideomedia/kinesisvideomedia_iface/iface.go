
package kinesisvideomedia_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
)

// IClient defines the interface for kinesisvideomedia
type IClient interface {
 Options() Options 
 GetMedia(ctx context.Context, params *GetMediaInput, optFns ...func(*Options)) (*GetMediaOutput, error) 
}
