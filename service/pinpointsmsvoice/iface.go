
package pinpointsmsvoice

import (
    "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
)

// IPinpointsmsvoice defines the interface for pinpointsmsvoice
type IPinpointsmsvoice interface {
 Options() Options 
 CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) 
 CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) 
 DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) 
 DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) 
 GetConfigurationSetEventDestinations(ctx context.Context, params *GetConfigurationSetEventDestinationsInput, optFns ...func(*Options)) (*GetConfigurationSetEventDestinationsOutput, error) 
 ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) 
 SendVoiceMessage(ctx context.Context, params *SendVoiceMessageInput, optFns ...func(*Options)) (*SendVoiceMessageOutput, error) 
 UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) 
}
