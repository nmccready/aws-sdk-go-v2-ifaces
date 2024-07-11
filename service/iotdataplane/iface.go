
package iotdataplane

import (
    "github.com/aws/aws-sdk-go-v2/service/iotdataplane"
)

// IClient defines the interface for iotdataplane
type IClient interface {
 Options() Options 
 DeleteThingShadow(ctx context.Context, params *DeleteThingShadowInput, optFns ...func(*Options)) (*DeleteThingShadowOutput, error) 
 GetRetainedMessage(ctx context.Context, params *GetRetainedMessageInput, optFns ...func(*Options)) (*GetRetainedMessageOutput, error) 
 GetThingShadow(ctx context.Context, params *GetThingShadowInput, optFns ...func(*Options)) (*GetThingShadowOutput, error) 
 ListNamedShadowsForThing(ctx context.Context, params *ListNamedShadowsForThingInput, optFns ...func(*Options)) (*ListNamedShadowsForThingOutput, error) 
 ListRetainedMessages(ctx context.Context, params *ListRetainedMessagesInput, optFns ...func(*Options)) (*ListRetainedMessagesOutput, error) 
 Publish(ctx context.Context, params *PublishInput, optFns ...func(*Options)) (*PublishOutput, error) 
 UpdateThingShadow(ctx context.Context, params *UpdateThingShadowInput, optFns ...func(*Options)) (*UpdateThingShadowOutput, error) 
}
