
package lexruntimev2

import (
    "github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
)

// ILexruntimev2 defines the interface for lexruntimev2
type ILexruntimev2 interface {
 Options() Options 
 DeleteSession(ctx context.Context, params *DeleteSessionInput, optFns ...func(*Options)) (*DeleteSessionOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 PutSession(ctx context.Context, params *PutSessionInput, optFns ...func(*Options)) (*PutSessionOutput, error) 
 RecognizeText(ctx context.Context, params *RecognizeTextInput, optFns ...func(*Options)) (*RecognizeTextOutput, error) 
 RecognizeUtterance(ctx context.Context, params *RecognizeUtteranceInput, optFns ...func(*Options)) (*RecognizeUtteranceOutput, error) 
 StartConversation(ctx context.Context, params *StartConversationInput, optFns ...func(*Options)) (*StartConversationOutput, error) 
}
