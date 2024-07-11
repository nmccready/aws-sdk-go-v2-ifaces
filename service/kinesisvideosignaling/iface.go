
package kinesisvideosignaling

import (
    "github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
)

// IKinesisvideosignaling defines the interface for kinesisvideosignaling
type IKinesisvideosignaling interface {
 Options() Options 
 GetIceServerConfig(ctx context.Context, params *GetIceServerConfigInput, optFns ...func(*Options)) (*GetIceServerConfigOutput, error) 
 SendAlexaOfferToMaster(ctx context.Context, params *SendAlexaOfferToMasterInput, optFns ...func(*Options)) (*SendAlexaOfferToMasterOutput, error) 
}
