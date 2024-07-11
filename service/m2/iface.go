
package m2

import (
    "github.com/aws/aws-sdk-go-v2/service/m2"
)

// IClient defines the interface for m2
type IClient interface {
 Options() Options 
 CancelBatchJobExecution(ctx context.Context, params *CancelBatchJobExecutionInput, optFns ...func(*Options)) (*CancelBatchJobExecutionOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateDataSetImportTask(ctx context.Context, params *CreateDataSetImportTaskInput, optFns ...func(*Options)) (*CreateDataSetImportTaskOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteApplicationFromEnvironment(ctx context.Context, params *DeleteApplicationFromEnvironmentInput, optFns ...func(*Options)) (*DeleteApplicationFromEnvironmentOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetApplicationVersion(ctx context.Context, params *GetApplicationVersionInput, optFns ...func(*Options)) (*GetApplicationVersionOutput, error) 
 GetBatchJobExecution(ctx context.Context, params *GetBatchJobExecutionInput, optFns ...func(*Options)) (*GetBatchJobExecutionOutput, error) 
 GetDataSetDetails(ctx context.Context, params *GetDataSetDetailsInput, optFns ...func(*Options)) (*GetDataSetDetailsOutput, error) 
 GetDataSetImportTask(ctx context.Context, params *GetDataSetImportTaskInput, optFns ...func(*Options)) (*GetDataSetImportTaskOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 GetSignedBluinsightsUrl(ctx context.Context, params *GetSignedBluinsightsUrlInput, optFns ...func(*Options)) (*GetSignedBluinsightsUrlOutput, error) 
 ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListBatchJobDefinitions(ctx context.Context, params *ListBatchJobDefinitionsInput, optFns ...func(*Options)) (*ListBatchJobDefinitionsOutput, error) 
 ListBatchJobExecutions(ctx context.Context, params *ListBatchJobExecutionsInput, optFns ...func(*Options)) (*ListBatchJobExecutionsOutput, error) 
 ListBatchJobRestartPoints(ctx context.Context, params *ListBatchJobRestartPointsInput, optFns ...func(*Options)) (*ListBatchJobRestartPointsOutput, error) 
 ListDataSetImportHistory(ctx context.Context, params *ListDataSetImportHistoryInput, optFns ...func(*Options)) (*ListDataSetImportHistoryOutput, error) 
 ListDataSets(ctx context.Context, params *ListDataSetsInput, optFns ...func(*Options)) (*ListDataSetsOutput, error) 
 ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) 
 ListEngineVersions(ctx context.Context, params *ListEngineVersionsInput, optFns ...func(*Options)) (*ListEngineVersionsOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) 
 StartBatchJob(ctx context.Context, params *StartBatchJobInput, optFns ...func(*Options)) (*StartBatchJobOutput, error) 
 StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
}
