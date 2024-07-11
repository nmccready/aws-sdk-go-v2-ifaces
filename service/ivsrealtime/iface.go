
package ivsrealtime

import (
    "github.com/aws/aws-sdk-go-v2/service/ivsrealtime"
)

// IClient defines the interface for ivsrealtime
type IClient interface {
 Options() Options 
 CreateEncoderConfiguration(ctx context.Context, params *CreateEncoderConfigurationInput, optFns ...func(*Options)) (*CreateEncoderConfigurationOutput, error) 
 CreateParticipantToken(ctx context.Context, params *CreateParticipantTokenInput, optFns ...func(*Options)) (*CreateParticipantTokenOutput, error) 
 CreateStage(ctx context.Context, params *CreateStageInput, optFns ...func(*Options)) (*CreateStageOutput, error) 
 CreateStorageConfiguration(ctx context.Context, params *CreateStorageConfigurationInput, optFns ...func(*Options)) (*CreateStorageConfigurationOutput, error) 
 DeleteEncoderConfiguration(ctx context.Context, params *DeleteEncoderConfigurationInput, optFns ...func(*Options)) (*DeleteEncoderConfigurationOutput, error) 
 DeletePublicKey(ctx context.Context, params *DeletePublicKeyInput, optFns ...func(*Options)) (*DeletePublicKeyOutput, error) 
 DeleteStage(ctx context.Context, params *DeleteStageInput, optFns ...func(*Options)) (*DeleteStageOutput, error) 
 DeleteStorageConfiguration(ctx context.Context, params *DeleteStorageConfigurationInput, optFns ...func(*Options)) (*DeleteStorageConfigurationOutput, error) 
 DisconnectParticipant(ctx context.Context, params *DisconnectParticipantInput, optFns ...func(*Options)) (*DisconnectParticipantOutput, error) 
 GetComposition(ctx context.Context, params *GetCompositionInput, optFns ...func(*Options)) (*GetCompositionOutput, error) 
 GetEncoderConfiguration(ctx context.Context, params *GetEncoderConfigurationInput, optFns ...func(*Options)) (*GetEncoderConfigurationOutput, error) 
 GetParticipant(ctx context.Context, params *GetParticipantInput, optFns ...func(*Options)) (*GetParticipantOutput, error) 
 GetPublicKey(ctx context.Context, params *GetPublicKeyInput, optFns ...func(*Options)) (*GetPublicKeyOutput, error) 
 GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) 
 GetStageSession(ctx context.Context, params *GetStageSessionInput, optFns ...func(*Options)) (*GetStageSessionOutput, error) 
 GetStorageConfiguration(ctx context.Context, params *GetStorageConfigurationInput, optFns ...func(*Options)) (*GetStorageConfigurationOutput, error) 
 ImportPublicKey(ctx context.Context, params *ImportPublicKeyInput, optFns ...func(*Options)) (*ImportPublicKeyOutput, error) 
 ListCompositions(ctx context.Context, params *ListCompositionsInput, optFns ...func(*Options)) (*ListCompositionsOutput, error) 
 ListEncoderConfigurations(ctx context.Context, params *ListEncoderConfigurationsInput, optFns ...func(*Options)) (*ListEncoderConfigurationsOutput, error) 
 ListParticipantEvents(ctx context.Context, params *ListParticipantEventsInput, optFns ...func(*Options)) (*ListParticipantEventsOutput, error) 
 ListParticipants(ctx context.Context, params *ListParticipantsInput, optFns ...func(*Options)) (*ListParticipantsOutput, error) 
 ListPublicKeys(ctx context.Context, params *ListPublicKeysInput, optFns ...func(*Options)) (*ListPublicKeysOutput, error) 
 ListStageSessions(ctx context.Context, params *ListStageSessionsInput, optFns ...func(*Options)) (*ListStageSessionsOutput, error) 
 ListStages(ctx context.Context, params *ListStagesInput, optFns ...func(*Options)) (*ListStagesOutput, error) 
 ListStorageConfigurations(ctx context.Context, params *ListStorageConfigurationsInput, optFns ...func(*Options)) (*ListStorageConfigurationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartComposition(ctx context.Context, params *StartCompositionInput, optFns ...func(*Options)) (*StartCompositionOutput, error) 
 StopComposition(ctx context.Context, params *StopCompositionInput, optFns ...func(*Options)) (*StopCompositionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateStage(ctx context.Context, params *UpdateStageInput, optFns ...func(*Options)) (*UpdateStageOutput, error) 
}
