
package kinesisvideosignaling_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
)

// IClient defines the interface for kinesisvideosignaling
type IClient interface {
 Options() Options 
 GetIceServerConfig(ctx context.Context, params *GetIceServerConfigInput, optFns ...func(*Options)) (*GetIceServerConfigOutput, error) 
 SendAlexaOfferToMaster(ctx context.Context, params *SendAlexaOfferToMasterInput, optFns ...func(*Options)) (*SendAlexaOfferToMasterOutput, error) 
}