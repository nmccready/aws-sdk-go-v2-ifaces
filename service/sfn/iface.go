
package sfn

import (
    "github.com/aws/aws-sdk-go-v2/service/sfn"
)

// IClient defines the interface for sfn
type IClient interface {
 Options() Options 
 CreateActivity(ctx context.Context, params *CreateActivityInput, optFns ...func(*Options)) (*CreateActivityOutput, error) 
 CreateStateMachine(ctx context.Context, params *CreateStateMachineInput, optFns ...func(*Options)) (*CreateStateMachineOutput, error) 
 CreateStateMachineAlias(ctx context.Context, params *CreateStateMachineAliasInput, optFns ...func(*Options)) (*CreateStateMachineAliasOutput, error) 
 DeleteActivity(ctx context.Context, params *DeleteActivityInput, optFns ...func(*Options)) (*DeleteActivityOutput, error) 
 DeleteStateMachine(ctx context.Context, params *DeleteStateMachineInput, optFns ...func(*Options)) (*DeleteStateMachineOutput, error) 
 DeleteStateMachineAlias(ctx context.Context, params *DeleteStateMachineAliasInput, optFns ...func(*Options)) (*DeleteStateMachineAliasOutput, error) 
 DeleteStateMachineVersion(ctx context.Context, params *DeleteStateMachineVersionInput, optFns ...func(*Options)) (*DeleteStateMachineVersionOutput, error) 
 DescribeActivity(ctx context.Context, params *DescribeActivityInput, optFns ...func(*Options)) (*DescribeActivityOutput, error) 
 DescribeExecution(ctx context.Context, params *DescribeExecutionInput, optFns ...func(*Options)) (*DescribeExecutionOutput, error) 
 DescribeMapRun(ctx context.Context, params *DescribeMapRunInput, optFns ...func(*Options)) (*DescribeMapRunOutput, error) 
 DescribeStateMachine(ctx context.Context, params *DescribeStateMachineInput, optFns ...func(*Options)) (*DescribeStateMachineOutput, error) 
 DescribeStateMachineAlias(ctx context.Context, params *DescribeStateMachineAliasInput, optFns ...func(*Options)) (*DescribeStateMachineAliasOutput, error) 
 DescribeStateMachineForExecution(ctx context.Context, params *DescribeStateMachineForExecutionInput, optFns ...func(*Options)) (*DescribeStateMachineForExecutionOutput, error) 
 GetActivityTask(ctx context.Context, params *GetActivityTaskInput, optFns ...func(*Options)) (*GetActivityTaskOutput, error) 
 GetExecutionHistory(ctx context.Context, params *GetExecutionHistoryInput, optFns ...func(*Options)) (*GetExecutionHistoryOutput, error) 
 ListActivities(ctx context.Context, params *ListActivitiesInput, optFns ...func(*Options)) (*ListActivitiesOutput, error) 
 ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) 
 ListMapRuns(ctx context.Context, params *ListMapRunsInput, optFns ...func(*Options)) (*ListMapRunsOutput, error) 
 ListStateMachineAliases(ctx context.Context, params *ListStateMachineAliasesInput, optFns ...func(*Options)) (*ListStateMachineAliasesOutput, error) 
 ListStateMachineVersions(ctx context.Context, params *ListStateMachineVersionsInput, optFns ...func(*Options)) (*ListStateMachineVersionsOutput, error) 
 ListStateMachines(ctx context.Context, params *ListStateMachinesInput, optFns ...func(*Options)) (*ListStateMachinesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PublishStateMachineVersion(ctx context.Context, params *PublishStateMachineVersionInput, optFns ...func(*Options)) (*PublishStateMachineVersionOutput, error) 
 RedriveExecution(ctx context.Context, params *RedriveExecutionInput, optFns ...func(*Options)) (*RedriveExecutionOutput, error) 
 SendTaskFailure(ctx context.Context, params *SendTaskFailureInput, optFns ...func(*Options)) (*SendTaskFailureOutput, error) 
 SendTaskHeartbeat(ctx context.Context, params *SendTaskHeartbeatInput, optFns ...func(*Options)) (*SendTaskHeartbeatOutput, error) 
 SendTaskSuccess(ctx context.Context, params *SendTaskSuccessInput, optFns ...func(*Options)) (*SendTaskSuccessOutput, error) 
 StartExecution(ctx context.Context, params *StartExecutionInput, optFns ...func(*Options)) (*StartExecutionOutput, error) 
 StartSyncExecution(ctx context.Context, params *StartSyncExecutionInput, optFns ...func(*Options)) (*StartSyncExecutionOutput, error) 
 StopExecution(ctx context.Context, params *StopExecutionInput, optFns ...func(*Options)) (*StopExecutionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestState(ctx context.Context, params *TestStateInput, optFns ...func(*Options)) (*TestStateOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateMapRun(ctx context.Context, params *UpdateMapRunInput, optFns ...func(*Options)) (*UpdateMapRunOutput, error) 
 UpdateStateMachine(ctx context.Context, params *UpdateStateMachineInput, optFns ...func(*Options)) (*UpdateStateMachineOutput, error) 
 UpdateStateMachineAlias(ctx context.Context, params *UpdateStateMachineAliasInput, optFns ...func(*Options)) (*UpdateStateMachineAliasOutput, error) 
 ValidateStateMachineDefinition(ctx context.Context, params *ValidateStateMachineDefinitionInput, optFns ...func(*Options)) (*ValidateStateMachineDefinitionOutput, error) 
}
