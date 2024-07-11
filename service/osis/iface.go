
package osis

import (
    "github.com/aws/aws-sdk-go-v2/service/osis"
)

// IClient defines the interface for osis
type IClient interface {
 Options() Options 
 CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) 
 DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) 
 GetPipeline(ctx context.Context, params *GetPipelineInput, optFns ...func(*Options)) (*GetPipelineOutput, error) 
 GetPipelineBlueprint(ctx context.Context, params *GetPipelineBlueprintInput, optFns ...func(*Options)) (*GetPipelineBlueprintOutput, error) 
 GetPipelineChangeProgress(ctx context.Context, params *GetPipelineChangeProgressInput, optFns ...func(*Options)) (*GetPipelineChangeProgressOutput, error) 
 ListPipelineBlueprints(ctx context.Context, params *ListPipelineBlueprintsInput, optFns ...func(*Options)) (*ListPipelineBlueprintsOutput, error) 
 ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartPipeline(ctx context.Context, params *StartPipelineInput, optFns ...func(*Options)) (*StartPipelineOutput, error) 
 StopPipeline(ctx context.Context, params *StopPipelineInput, optFns ...func(*Options)) (*StopPipelineOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) 
 ValidatePipeline(ctx context.Context, params *ValidatePipelineInput, optFns ...func(*Options)) (*ValidatePipelineOutput, error) 
}
