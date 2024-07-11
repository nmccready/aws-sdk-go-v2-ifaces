
package codepipeline

import (
    "github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

// IClient defines the interface for codepipeline
type IClient interface {
 Options() Options 
 AcknowledgeJob(ctx context.Context, params *AcknowledgeJobInput, optFns ...func(*Options)) (*AcknowledgeJobOutput, error) 
 AcknowledgeThirdPartyJob(ctx context.Context, params *AcknowledgeThirdPartyJobInput, optFns ...func(*Options)) (*AcknowledgeThirdPartyJobOutput, error) 
 CreateCustomActionType(ctx context.Context, params *CreateCustomActionTypeInput, optFns ...func(*Options)) (*CreateCustomActionTypeOutput, error) 
 CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) 
 DeleteCustomActionType(ctx context.Context, params *DeleteCustomActionTypeInput, optFns ...func(*Options)) (*DeleteCustomActionTypeOutput, error) 
 DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) 
 DeleteWebhook(ctx context.Context, params *DeleteWebhookInput, optFns ...func(*Options)) (*DeleteWebhookOutput, error) 
 DeregisterWebhookWithThirdParty(ctx context.Context, params *DeregisterWebhookWithThirdPartyInput, optFns ...func(*Options)) (*DeregisterWebhookWithThirdPartyOutput, error) 
 DisableStageTransition(ctx context.Context, params *DisableStageTransitionInput, optFns ...func(*Options)) (*DisableStageTransitionOutput, error) 
 EnableStageTransition(ctx context.Context, params *EnableStageTransitionInput, optFns ...func(*Options)) (*EnableStageTransitionOutput, error) 
 GetActionType(ctx context.Context, params *GetActionTypeInput, optFns ...func(*Options)) (*GetActionTypeOutput, error) 
 GetJobDetails(ctx context.Context, params *GetJobDetailsInput, optFns ...func(*Options)) (*GetJobDetailsOutput, error) 
 GetPipeline(ctx context.Context, params *GetPipelineInput, optFns ...func(*Options)) (*GetPipelineOutput, error) 
 GetPipelineExecution(ctx context.Context, params *GetPipelineExecutionInput, optFns ...func(*Options)) (*GetPipelineExecutionOutput, error) 
 GetPipelineState(ctx context.Context, params *GetPipelineStateInput, optFns ...func(*Options)) (*GetPipelineStateOutput, error) 
 GetThirdPartyJobDetails(ctx context.Context, params *GetThirdPartyJobDetailsInput, optFns ...func(*Options)) (*GetThirdPartyJobDetailsOutput, error) 
 ListActionExecutions(ctx context.Context, params *ListActionExecutionsInput, optFns ...func(*Options)) (*ListActionExecutionsOutput, error) 
 ListActionTypes(ctx context.Context, params *ListActionTypesInput, optFns ...func(*Options)) (*ListActionTypesOutput, error) 
 ListPipelineExecutions(ctx context.Context, params *ListPipelineExecutionsInput, optFns ...func(*Options)) (*ListPipelineExecutionsOutput, error) 
 ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebhooks(ctx context.Context, params *ListWebhooksInput, optFns ...func(*Options)) (*ListWebhooksOutput, error) 
 PollForJobs(ctx context.Context, params *PollForJobsInput, optFns ...func(*Options)) (*PollForJobsOutput, error) 
 PollForThirdPartyJobs(ctx context.Context, params *PollForThirdPartyJobsInput, optFns ...func(*Options)) (*PollForThirdPartyJobsOutput, error) 
 PutActionRevision(ctx context.Context, params *PutActionRevisionInput, optFns ...func(*Options)) (*PutActionRevisionOutput, error) 
 PutApprovalResult(ctx context.Context, params *PutApprovalResultInput, optFns ...func(*Options)) (*PutApprovalResultOutput, error) 
 PutJobFailureResult(ctx context.Context, params *PutJobFailureResultInput, optFns ...func(*Options)) (*PutJobFailureResultOutput, error) 
 PutJobSuccessResult(ctx context.Context, params *PutJobSuccessResultInput, optFns ...func(*Options)) (*PutJobSuccessResultOutput, error) 
 PutThirdPartyJobFailureResult(ctx context.Context, params *PutThirdPartyJobFailureResultInput, optFns ...func(*Options)) (*PutThirdPartyJobFailureResultOutput, error) 
 PutThirdPartyJobSuccessResult(ctx context.Context, params *PutThirdPartyJobSuccessResultInput, optFns ...func(*Options)) (*PutThirdPartyJobSuccessResultOutput, error) 
 PutWebhook(ctx context.Context, params *PutWebhookInput, optFns ...func(*Options)) (*PutWebhookOutput, error) 
 RegisterWebhookWithThirdParty(ctx context.Context, params *RegisterWebhookWithThirdPartyInput, optFns ...func(*Options)) (*RegisterWebhookWithThirdPartyOutput, error) 
 RetryStageExecution(ctx context.Context, params *RetryStageExecutionInput, optFns ...func(*Options)) (*RetryStageExecutionOutput, error) 
 RollbackStage(ctx context.Context, params *RollbackStageInput, optFns ...func(*Options)) (*RollbackStageOutput, error) 
 StartPipelineExecution(ctx context.Context, params *StartPipelineExecutionInput, optFns ...func(*Options)) (*StartPipelineExecutionOutput, error) 
 StopPipelineExecution(ctx context.Context, params *StopPipelineExecutionInput, optFns ...func(*Options)) (*StopPipelineExecutionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateActionType(ctx context.Context, params *UpdateActionTypeInput, optFns ...func(*Options)) (*UpdateActionTypeOutput, error) 
 UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) 
}
