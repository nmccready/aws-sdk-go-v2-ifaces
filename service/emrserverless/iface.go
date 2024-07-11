
package emrserverless

import (
    "github.com/aws/aws-sdk-go-v2/service/emrserverless"
)

// IClient defines the interface for emrserverless
type IClient interface {
 Options() Options 
 CancelJobRun(ctx context.Context, params *CancelJobRunInput, optFns ...func(*Options)) (*CancelJobRunOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetDashboardForJobRun(ctx context.Context, params *GetDashboardForJobRunInput, optFns ...func(*Options)) (*GetDashboardForJobRunOutput, error) 
 GetJobRun(ctx context.Context, params *GetJobRunInput, optFns ...func(*Options)) (*GetJobRunOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListJobRunAttempts(ctx context.Context, params *ListJobRunAttemptsInput, optFns ...func(*Options)) (*ListJobRunAttemptsOutput, error) 
 ListJobRuns(ctx context.Context, params *ListJobRunsInput, optFns ...func(*Options)) (*ListJobRunsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) 
 StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) 
 StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
}
