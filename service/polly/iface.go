
package polly

import (
    "github.com/aws/aws-sdk-go-v2/service/polly"
)

// IClient defines the interface for polly
type IClient interface {
 Options() Options 
 DeleteLexicon(ctx context.Context, params *DeleteLexiconInput, optFns ...func(*Options)) (*DeleteLexiconOutput, error) 
 DescribeVoices(ctx context.Context, params *DescribeVoicesInput, optFns ...func(*Options)) (*DescribeVoicesOutput, error) 
 GetLexicon(ctx context.Context, params *GetLexiconInput, optFns ...func(*Options)) (*GetLexiconOutput, error) 
 GetSpeechSynthesisTask(ctx context.Context, params *GetSpeechSynthesisTaskInput, optFns ...func(*Options)) (*GetSpeechSynthesisTaskOutput, error) 
 ListLexicons(ctx context.Context, params *ListLexiconsInput, optFns ...func(*Options)) (*ListLexiconsOutput, error) 
 ListSpeechSynthesisTasks(ctx context.Context, params *ListSpeechSynthesisTasksInput, optFns ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error) 
 PutLexicon(ctx context.Context, params *PutLexiconInput, optFns ...func(*Options)) (*PutLexiconOutput, error) 
 StartSpeechSynthesisTask(ctx context.Context, params *StartSpeechSynthesisTaskInput, optFns ...func(*Options)) (*StartSpeechSynthesisTaskOutput, error) 
 SynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*Options)) (*SynthesizeSpeechOutput, error) 
 PresignSynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) 
}
