
package transcribestreaming

import (
    "github.com/aws/aws-sdk-go-v2/service/transcribestreaming"
)

// ITranscribestreaming defines the interface for transcribestreaming
type ITranscribestreaming interface {
 Options() Options 
 StartCallAnalyticsStreamTranscription(ctx context.Context, params *StartCallAnalyticsStreamTranscriptionInput, optFns ...func(*Options)) (*StartCallAnalyticsStreamTranscriptionOutput, error) 
 StartMedicalStreamTranscription(ctx context.Context, params *StartMedicalStreamTranscriptionInput, optFns ...func(*Options)) (*StartMedicalStreamTranscriptionOutput, error) 
 StartStreamTranscription(ctx context.Context, params *StartStreamTranscriptionInput, optFns ...func(*Options)) (*StartStreamTranscriptionOutput, error) 
}
