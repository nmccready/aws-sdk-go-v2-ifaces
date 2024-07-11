
package bedrockruntime

import (
    "github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// IBedrockruntime defines the interface for bedrockruntime
type IBedrockruntime interface {
 Options() Options 
 ApplyGuardrail(ctx context.Context, params *ApplyGuardrailInput, optFns ...func(*Options)) (*ApplyGuardrailOutput, error) 
 Converse(ctx context.Context, params *ConverseInput, optFns ...func(*Options)) (*ConverseOutput, error) 
 ConverseStream(ctx context.Context, params *ConverseStreamInput, optFns ...func(*Options)) (*ConverseStreamOutput, error) 
 InvokeModel(ctx context.Context, params *InvokeModelInput, optFns ...func(*Options)) (*InvokeModelOutput, error) 
 InvokeModelWithResponseStream(ctx context.Context, params *InvokeModelWithResponseStreamInput, optFns ...func(*Options)) (*InvokeModelWithResponseStreamOutput, error) 
}
