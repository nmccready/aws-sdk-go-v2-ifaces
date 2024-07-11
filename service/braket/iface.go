
package braket

import (
    "github.com/aws/aws-sdk-go-v2/service/braket"
)

// IBraket defines the interface for braket
type IBraket interface {
 Options() Options 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CancelQuantumTask(ctx context.Context, params *CancelQuantumTaskInput, optFns ...func(*Options)) (*CancelQuantumTaskOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateQuantumTask(ctx context.Context, params *CreateQuantumTaskInput, optFns ...func(*Options)) (*CreateQuantumTaskOutput, error) 
 GetDevice(ctx context.Context, params *GetDeviceInput, optFns ...func(*Options)) (*GetDeviceOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetQuantumTask(ctx context.Context, params *GetQuantumTaskInput, optFns ...func(*Options)) (*GetQuantumTaskOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 SearchDevices(ctx context.Context, params *SearchDevicesInput, optFns ...func(*Options)) (*SearchDevicesOutput, error) 
 SearchJobs(ctx context.Context, params *SearchJobsInput, optFns ...func(*Options)) (*SearchJobsOutput, error) 
 SearchQuantumTasks(ctx context.Context, params *SearchQuantumTasksInput, optFns ...func(*Options)) (*SearchQuantumTasksOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
