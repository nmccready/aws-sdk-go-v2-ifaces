
package ivschat

import (
    "github.com/aws/aws-sdk-go-v2/service/ivschat"
)

// IClient defines the interface for ivschat
type IClient interface {
 Options() Options 
 CreateChatToken(ctx context.Context, params *CreateChatTokenInput, optFns ...func(*Options)) (*CreateChatTokenOutput, error) 
 CreateLoggingConfiguration(ctx context.Context, params *CreateLoggingConfigurationInput, optFns ...func(*Options)) (*CreateLoggingConfigurationOutput, error) 
 CreateRoom(ctx context.Context, params *CreateRoomInput, optFns ...func(*Options)) (*CreateRoomOutput, error) 
 DeleteLoggingConfiguration(ctx context.Context, params *DeleteLoggingConfigurationInput, optFns ...func(*Options)) (*DeleteLoggingConfigurationOutput, error) 
 DeleteMessage(ctx context.Context, params *DeleteMessageInput, optFns ...func(*Options)) (*DeleteMessageOutput, error) 
 DeleteRoom(ctx context.Context, params *DeleteRoomInput, optFns ...func(*Options)) (*DeleteRoomOutput, error) 
 DisconnectUser(ctx context.Context, params *DisconnectUserInput, optFns ...func(*Options)) (*DisconnectUserOutput, error) 
 GetLoggingConfiguration(ctx context.Context, params *GetLoggingConfigurationInput, optFns ...func(*Options)) (*GetLoggingConfigurationOutput, error) 
 GetRoom(ctx context.Context, params *GetRoomInput, optFns ...func(*Options)) (*GetRoomOutput, error) 
 ListLoggingConfigurations(ctx context.Context, params *ListLoggingConfigurationsInput, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) 
 ListRooms(ctx context.Context, params *ListRoomsInput, optFns ...func(*Options)) (*ListRoomsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLoggingConfiguration(ctx context.Context, params *UpdateLoggingConfigurationInput, optFns ...func(*Options)) (*UpdateLoggingConfigurationOutput, error) 
 UpdateRoom(ctx context.Context, params *UpdateRoomInput, optFns ...func(*Options)) (*UpdateRoomOutput, error) 
}
