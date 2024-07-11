
package batch

import (
    "github.com/aws/aws-sdk-go-v2/service/batch"
)

// IBatch defines the interface for batch
type IBatch interface {
 Options() Options 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CreateComputeEnvironment(ctx context.Context, params *CreateComputeEnvironmentInput, optFns ...func(*Options)) (*CreateComputeEnvironmentOutput, error) 
 CreateJobQueue(ctx context.Context, params *CreateJobQueueInput, optFns ...func(*Options)) (*CreateJobQueueOutput, error) 
 CreateSchedulingPolicy(ctx context.Context, params *CreateSchedulingPolicyInput, optFns ...func(*Options)) (*CreateSchedulingPolicyOutput, error) 
 DeleteComputeEnvironment(ctx context.Context, params *DeleteComputeEnvironmentInput, optFns ...func(*Options)) (*DeleteComputeEnvironmentOutput, error) 
 DeleteJobQueue(ctx context.Context, params *DeleteJobQueueInput, optFns ...func(*Options)) (*DeleteJobQueueOutput, error) 
 DeleteSchedulingPolicy(ctx context.Context, params *DeleteSchedulingPolicyInput, optFns ...func(*Options)) (*DeleteSchedulingPolicyOutput, error) 
 DeregisterJobDefinition(ctx context.Context, params *DeregisterJobDefinitionInput, optFns ...func(*Options)) (*DeregisterJobDefinitionOutput, error) 
 DescribeComputeEnvironments(ctx context.Context, params *DescribeComputeEnvironmentsInput, optFns ...func(*Options)) (*DescribeComputeEnvironmentsOutput, error) 
 DescribeJobDefinitions(ctx context.Context, params *DescribeJobDefinitionsInput, optFns ...func(*Options)) (*DescribeJobDefinitionsOutput, error) 
 DescribeJobQueues(ctx context.Context, params *DescribeJobQueuesInput, optFns ...func(*Options)) (*DescribeJobQueuesOutput, error) 
 DescribeJobs(ctx context.Context, params *DescribeJobsInput, optFns ...func(*Options)) (*DescribeJobsOutput, error) 
 DescribeSchedulingPolicies(ctx context.Context, params *DescribeSchedulingPoliciesInput, optFns ...func(*Options)) (*DescribeSchedulingPoliciesOutput, error) 
 GetJobQueueSnapshot(ctx context.Context, params *GetJobQueueSnapshotInput, optFns ...func(*Options)) (*GetJobQueueSnapshotOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListSchedulingPolicies(ctx context.Context, params *ListSchedulingPoliciesInput, optFns ...func(*Options)) (*ListSchedulingPoliciesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterJobDefinition(ctx context.Context, params *RegisterJobDefinitionInput, optFns ...func(*Options)) (*RegisterJobDefinitionOutput, error) 
 SubmitJob(ctx context.Context, params *SubmitJobInput, optFns ...func(*Options)) (*SubmitJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TerminateJob(ctx context.Context, params *TerminateJobInput, optFns ...func(*Options)) (*TerminateJobOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateComputeEnvironment(ctx context.Context, params *UpdateComputeEnvironmentInput, optFns ...func(*Options)) (*UpdateComputeEnvironmentOutput, error) 
 UpdateJobQueue(ctx context.Context, params *UpdateJobQueueInput, optFns ...func(*Options)) (*UpdateJobQueueOutput, error) 
 UpdateSchedulingPolicy(ctx context.Context, params *UpdateSchedulingPolicyInput, optFns ...func(*Options)) (*UpdateSchedulingPolicyOutput, error) 
}
