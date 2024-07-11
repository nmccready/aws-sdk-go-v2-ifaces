
package swf

import (
    "github.com/aws/aws-sdk-go-v2/service/swf"
)

// ISwf defines the interface for swf
type ISwf interface {
 Options() Options 
 CountClosedWorkflowExecutions(ctx context.Context, params *CountClosedWorkflowExecutionsInput, optFns ...func(*Options)) (*CountClosedWorkflowExecutionsOutput, error) 
 CountOpenWorkflowExecutions(ctx context.Context, params *CountOpenWorkflowExecutionsInput, optFns ...func(*Options)) (*CountOpenWorkflowExecutionsOutput, error) 
 CountPendingActivityTasks(ctx context.Context, params *CountPendingActivityTasksInput, optFns ...func(*Options)) (*CountPendingActivityTasksOutput, error) 
 CountPendingDecisionTasks(ctx context.Context, params *CountPendingDecisionTasksInput, optFns ...func(*Options)) (*CountPendingDecisionTasksOutput, error) 
 DeleteActivityType(ctx context.Context, params *DeleteActivityTypeInput, optFns ...func(*Options)) (*DeleteActivityTypeOutput, error) 
 DeleteWorkflowType(ctx context.Context, params *DeleteWorkflowTypeInput, optFns ...func(*Options)) (*DeleteWorkflowTypeOutput, error) 
 DeprecateActivityType(ctx context.Context, params *DeprecateActivityTypeInput, optFns ...func(*Options)) (*DeprecateActivityTypeOutput, error) 
 DeprecateDomain(ctx context.Context, params *DeprecateDomainInput, optFns ...func(*Options)) (*DeprecateDomainOutput, error) 
 DeprecateWorkflowType(ctx context.Context, params *DeprecateWorkflowTypeInput, optFns ...func(*Options)) (*DeprecateWorkflowTypeOutput, error) 
 DescribeActivityType(ctx context.Context, params *DescribeActivityTypeInput, optFns ...func(*Options)) (*DescribeActivityTypeOutput, error) 
 DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) 
 DescribeWorkflowExecution(ctx context.Context, params *DescribeWorkflowExecutionInput, optFns ...func(*Options)) (*DescribeWorkflowExecutionOutput, error) 
 DescribeWorkflowType(ctx context.Context, params *DescribeWorkflowTypeInput, optFns ...func(*Options)) (*DescribeWorkflowTypeOutput, error) 
 GetWorkflowExecutionHistory(ctx context.Context, params *GetWorkflowExecutionHistoryInput, optFns ...func(*Options)) (*GetWorkflowExecutionHistoryOutput, error) 
 ListActivityTypes(ctx context.Context, params *ListActivityTypesInput, optFns ...func(*Options)) (*ListActivityTypesOutput, error) 
 ListClosedWorkflowExecutions(ctx context.Context, params *ListClosedWorkflowExecutionsInput, optFns ...func(*Options)) (*ListClosedWorkflowExecutionsOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListOpenWorkflowExecutions(ctx context.Context, params *ListOpenWorkflowExecutionsInput, optFns ...func(*Options)) (*ListOpenWorkflowExecutionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkflowTypes(ctx context.Context, params *ListWorkflowTypesInput, optFns ...func(*Options)) (*ListWorkflowTypesOutput, error) 
 PollForActivityTask(ctx context.Context, params *PollForActivityTaskInput, optFns ...func(*Options)) (*PollForActivityTaskOutput, error) 
 PollForDecisionTask(ctx context.Context, params *PollForDecisionTaskInput, optFns ...func(*Options)) (*PollForDecisionTaskOutput, error) 
 RecordActivityTaskHeartbeat(ctx context.Context, params *RecordActivityTaskHeartbeatInput, optFns ...func(*Options)) (*RecordActivityTaskHeartbeatOutput, error) 
 RegisterActivityType(ctx context.Context, params *RegisterActivityTypeInput, optFns ...func(*Options)) (*RegisterActivityTypeOutput, error) 
 RegisterDomain(ctx context.Context, params *RegisterDomainInput, optFns ...func(*Options)) (*RegisterDomainOutput, error) 
 RegisterWorkflowType(ctx context.Context, params *RegisterWorkflowTypeInput, optFns ...func(*Options)) (*RegisterWorkflowTypeOutput, error) 
 RequestCancelWorkflowExecution(ctx context.Context, params *RequestCancelWorkflowExecutionInput, optFns ...func(*Options)) (*RequestCancelWorkflowExecutionOutput, error) 
 RespondActivityTaskCanceled(ctx context.Context, params *RespondActivityTaskCanceledInput, optFns ...func(*Options)) (*RespondActivityTaskCanceledOutput, error) 
 RespondActivityTaskCompleted(ctx context.Context, params *RespondActivityTaskCompletedInput, optFns ...func(*Options)) (*RespondActivityTaskCompletedOutput, error) 
 RespondActivityTaskFailed(ctx context.Context, params *RespondActivityTaskFailedInput, optFns ...func(*Options)) (*RespondActivityTaskFailedOutput, error) 
 RespondDecisionTaskCompleted(ctx context.Context, params *RespondDecisionTaskCompletedInput, optFns ...func(*Options)) (*RespondDecisionTaskCompletedOutput, error) 
 SignalWorkflowExecution(ctx context.Context, params *SignalWorkflowExecutionInput, optFns ...func(*Options)) (*SignalWorkflowExecutionOutput, error) 
 StartWorkflowExecution(ctx context.Context, params *StartWorkflowExecutionInput, optFns ...func(*Options)) (*StartWorkflowExecutionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TerminateWorkflowExecution(ctx context.Context, params *TerminateWorkflowExecutionInput, optFns ...func(*Options)) (*TerminateWorkflowExecutionOutput, error) 
 UndeprecateActivityType(ctx context.Context, params *UndeprecateActivityTypeInput, optFns ...func(*Options)) (*UndeprecateActivityTypeOutput, error) 
 UndeprecateDomain(ctx context.Context, params *UndeprecateDomainInput, optFns ...func(*Options)) (*UndeprecateDomainOutput, error) 
 UndeprecateWorkflowType(ctx context.Context, params *UndeprecateWorkflowTypeInput, optFns ...func(*Options)) (*UndeprecateWorkflowTypeOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
