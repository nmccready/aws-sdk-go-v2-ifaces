
package datapipeline

import (
    "github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

// IDatapipeline defines the interface for datapipeline
type IDatapipeline interface {
 Options() Options 
 ActivatePipeline(ctx context.Context, params *ActivatePipelineInput, optFns ...func(*Options)) (*ActivatePipelineOutput, error) 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) 
 DeactivatePipeline(ctx context.Context, params *DeactivatePipelineInput, optFns ...func(*Options)) (*DeactivatePipelineOutput, error) 
 DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) 
 DescribeObjects(ctx context.Context, params *DescribeObjectsInput, optFns ...func(*Options)) (*DescribeObjectsOutput, error) 
 DescribePipelines(ctx context.Context, params *DescribePipelinesInput, optFns ...func(*Options)) (*DescribePipelinesOutput, error) 
 EvaluateExpression(ctx context.Context, params *EvaluateExpressionInput, optFns ...func(*Options)) (*EvaluateExpressionOutput, error) 
 GetPipelineDefinition(ctx context.Context, params *GetPipelineDefinitionInput, optFns ...func(*Options)) (*GetPipelineDefinitionOutput, error) 
 ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) 
 PollForTask(ctx context.Context, params *PollForTaskInput, optFns ...func(*Options)) (*PollForTaskOutput, error) 
 PutPipelineDefinition(ctx context.Context, params *PutPipelineDefinitionInput, optFns ...func(*Options)) (*PutPipelineDefinitionOutput, error) 
 QueryObjects(ctx context.Context, params *QueryObjectsInput, optFns ...func(*Options)) (*QueryObjectsOutput, error) 
 RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) 
 ReportTaskProgress(ctx context.Context, params *ReportTaskProgressInput, optFns ...func(*Options)) (*ReportTaskProgressOutput, error) 
 ReportTaskRunnerHeartbeat(ctx context.Context, params *ReportTaskRunnerHeartbeatInput, optFns ...func(*Options)) (*ReportTaskRunnerHeartbeatOutput, error) 
 SetStatus(ctx context.Context, params *SetStatusInput, optFns ...func(*Options)) (*SetStatusOutput, error) 
 SetTaskStatus(ctx context.Context, params *SetTaskStatusInput, optFns ...func(*Options)) (*SetTaskStatusOutput, error) 
 ValidatePipelineDefinition(ctx context.Context, params *ValidatePipelineDefinitionInput, optFns ...func(*Options)) (*ValidatePipelineDefinitionOutput, error) 
}
