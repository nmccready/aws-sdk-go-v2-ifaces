
package connectparticipant

import (
    "github.com/aws/aws-sdk-go-v2/service/connectparticipant"
)

// IConnectparticipant defines the interface for connectparticipant
type IConnectparticipant interface {
 Options() Options 
 CompleteAttachmentUpload(ctx context.Context, params *CompleteAttachmentUploadInput, optFns ...func(*Options)) (*CompleteAttachmentUploadOutput, error) 
 CreateParticipantConnection(ctx context.Context, params *CreateParticipantConnectionInput, optFns ...func(*Options)) (*CreateParticipantConnectionOutput, error) 
 DescribeView(ctx context.Context, params *DescribeViewInput, optFns ...func(*Options)) (*DescribeViewOutput, error) 
 DisconnectParticipant(ctx context.Context, params *DisconnectParticipantInput, optFns ...func(*Options)) (*DisconnectParticipantOutput, error) 
 GetAttachment(ctx context.Context, params *GetAttachmentInput, optFns ...func(*Options)) (*GetAttachmentOutput, error) 
 GetTranscript(ctx context.Context, params *GetTranscriptInput, optFns ...func(*Options)) (*GetTranscriptOutput, error) 
 SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) 
 SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) 
 StartAttachmentUpload(ctx context.Context, params *StartAttachmentUploadInput, optFns ...func(*Options)) (*StartAttachmentUploadOutput, error) 
}
