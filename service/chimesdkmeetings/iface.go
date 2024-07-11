
package chimesdkmeetings

import (
    "github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
)

// IClient defines the interface for chimesdkmeetings
type IClient interface {
 Options() Options 
 BatchCreateAttendee(ctx context.Context, params *BatchCreateAttendeeInput, optFns ...func(*Options)) (*BatchCreateAttendeeOutput, error) 
 BatchUpdateAttendeeCapabilitiesExcept(ctx context.Context, params *BatchUpdateAttendeeCapabilitiesExceptInput, optFns ...func(*Options)) (*BatchUpdateAttendeeCapabilitiesExceptOutput, error) 
 CreateAttendee(ctx context.Context, params *CreateAttendeeInput, optFns ...func(*Options)) (*CreateAttendeeOutput, error) 
 CreateMeeting(ctx context.Context, params *CreateMeetingInput, optFns ...func(*Options)) (*CreateMeetingOutput, error) 
 CreateMeetingWithAttendees(ctx context.Context, params *CreateMeetingWithAttendeesInput, optFns ...func(*Options)) (*CreateMeetingWithAttendeesOutput, error) 
 DeleteAttendee(ctx context.Context, params *DeleteAttendeeInput, optFns ...func(*Options)) (*DeleteAttendeeOutput, error) 
 DeleteMeeting(ctx context.Context, params *DeleteMeetingInput, optFns ...func(*Options)) (*DeleteMeetingOutput, error) 
 GetAttendee(ctx context.Context, params *GetAttendeeInput, optFns ...func(*Options)) (*GetAttendeeOutput, error) 
 GetMeeting(ctx context.Context, params *GetMeetingInput, optFns ...func(*Options)) (*GetMeetingOutput, error) 
 ListAttendees(ctx context.Context, params *ListAttendeesInput, optFns ...func(*Options)) (*ListAttendeesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartMeetingTranscription(ctx context.Context, params *StartMeetingTranscriptionInput, optFns ...func(*Options)) (*StartMeetingTranscriptionOutput, error) 
 StopMeetingTranscription(ctx context.Context, params *StopMeetingTranscriptionInput, optFns ...func(*Options)) (*StopMeetingTranscriptionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAttendeeCapabilities(ctx context.Context, params *UpdateAttendeeCapabilitiesInput, optFns ...func(*Options)) (*UpdateAttendeeCapabilitiesOutput, error) 
}
