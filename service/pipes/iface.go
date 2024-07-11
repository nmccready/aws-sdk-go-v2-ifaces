
package pipes

import (
    "github.com/aws/aws-sdk-go-v2/service/pipes"
)

// IClient defines the interface for pipes
type IClient interface {
 Options() Options 
 CreatePipe(ctx context.Context, params *CreatePipeInput, optFns ...func(*Options)) (*CreatePipeOutput, error) 
 DeletePipe(ctx context.Context, params *DeletePipeInput, optFns ...func(*Options)) (*DeletePipeOutput, error) 
 DescribePipe(ctx context.Context, params *DescribePipeInput, optFns ...func(*Options)) (*DescribePipeOutput, error) 
 ListPipes(ctx context.Context, params *ListPipesInput, optFns ...func(*Options)) (*ListPipesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartPipe(ctx context.Context, params *StartPipeInput, optFns ...func(*Options)) (*StartPipeOutput, error) 
 StopPipe(ctx context.Context, params *StopPipeInput, optFns ...func(*Options)) (*StopPipeOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePipe(ctx context.Context, params *UpdatePipeInput, optFns ...func(*Options)) (*UpdatePipeOutput, error) 
}
