
package lexruntimeservice

import (
    "github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
)

// ILexruntimeservice defines the interface for lexruntimeservice
type ILexruntimeservice interface {
 Options() Options 
 DeleteSession(ctx context.Context, params *DeleteSessionInput, optFns ...func(*Options)) (*DeleteSessionOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 PostContent(ctx context.Context, params *PostContentInput, optFns ...func(*Options)) (*PostContentOutput, error) 
 PostText(ctx context.Context, params *PostTextInput, optFns ...func(*Options)) (*PostTextOutput, error) 
 PutSession(ctx context.Context, params *PutSessionInput, optFns ...func(*Options)) (*PutSessionOutput, error) 
}
