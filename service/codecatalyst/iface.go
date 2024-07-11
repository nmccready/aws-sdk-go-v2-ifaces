
package codecatalyst

import (
    "github.com/aws/aws-sdk-go-v2/service/codecatalyst"
)

// ICodecatalyst defines the interface for codecatalyst
type ICodecatalyst interface {
 Options() Options 
 CreateAccessToken(ctx context.Context, params *CreateAccessTokenInput, optFns ...func(*Options)) (*CreateAccessTokenOutput, error) 
 CreateDevEnvironment(ctx context.Context, params *CreateDevEnvironmentInput, optFns ...func(*Options)) (*CreateDevEnvironmentOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 CreateSourceRepository(ctx context.Context, params *CreateSourceRepositoryInput, optFns ...func(*Options)) (*CreateSourceRepositoryOutput, error) 
 CreateSourceRepositoryBranch(ctx context.Context, params *CreateSourceRepositoryBranchInput, optFns ...func(*Options)) (*CreateSourceRepositoryBranchOutput, error) 
 DeleteAccessToken(ctx context.Context, params *DeleteAccessTokenInput, optFns ...func(*Options)) (*DeleteAccessTokenOutput, error) 
 DeleteDevEnvironment(ctx context.Context, params *DeleteDevEnvironmentInput, optFns ...func(*Options)) (*DeleteDevEnvironmentOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteSourceRepository(ctx context.Context, params *DeleteSourceRepositoryInput, optFns ...func(*Options)) (*DeleteSourceRepositoryOutput, error) 
 DeleteSpace(ctx context.Context, params *DeleteSpaceInput, optFns ...func(*Options)) (*DeleteSpaceOutput, error) 
 GetDevEnvironment(ctx context.Context, params *GetDevEnvironmentInput, optFns ...func(*Options)) (*GetDevEnvironmentOutput, error) 
 GetProject(ctx context.Context, params *GetProjectInput, optFns ...func(*Options)) (*GetProjectOutput, error) 
 GetSourceRepository(ctx context.Context, params *GetSourceRepositoryInput, optFns ...func(*Options)) (*GetSourceRepositoryOutput, error) 
 GetSourceRepositoryCloneUrls(ctx context.Context, params *GetSourceRepositoryCloneUrlsInput, optFns ...func(*Options)) (*GetSourceRepositoryCloneUrlsOutput, error) 
 GetSpace(ctx context.Context, params *GetSpaceInput, optFns ...func(*Options)) (*GetSpaceOutput, error) 
 GetSubscription(ctx context.Context, params *GetSubscriptionInput, optFns ...func(*Options)) (*GetSubscriptionOutput, error) 
 GetUserDetails(ctx context.Context, params *GetUserDetailsInput, optFns ...func(*Options)) (*GetUserDetailsOutput, error) 
 GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) 
 GetWorkflowRun(ctx context.Context, params *GetWorkflowRunInput, optFns ...func(*Options)) (*GetWorkflowRunOutput, error) 
 ListAccessTokens(ctx context.Context, params *ListAccessTokensInput, optFns ...func(*Options)) (*ListAccessTokensOutput, error) 
 ListDevEnvironmentSessions(ctx context.Context, params *ListDevEnvironmentSessionsInput, optFns ...func(*Options)) (*ListDevEnvironmentSessionsOutput, error) 
 ListDevEnvironments(ctx context.Context, params *ListDevEnvironmentsInput, optFns ...func(*Options)) (*ListDevEnvironmentsOutput, error) 
 ListEventLogs(ctx context.Context, params *ListEventLogsInput, optFns ...func(*Options)) (*ListEventLogsOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListSourceRepositories(ctx context.Context, params *ListSourceRepositoriesInput, optFns ...func(*Options)) (*ListSourceRepositoriesOutput, error) 
 ListSourceRepositoryBranches(ctx context.Context, params *ListSourceRepositoryBranchesInput, optFns ...func(*Options)) (*ListSourceRepositoryBranchesOutput, error) 
 ListSpaces(ctx context.Context, params *ListSpacesInput, optFns ...func(*Options)) (*ListSpacesOutput, error) 
 ListWorkflowRuns(ctx context.Context, params *ListWorkflowRunsInput, optFns ...func(*Options)) (*ListWorkflowRunsOutput, error) 
 ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) 
 StartDevEnvironment(ctx context.Context, params *StartDevEnvironmentInput, optFns ...func(*Options)) (*StartDevEnvironmentOutput, error) 
 StartDevEnvironmentSession(ctx context.Context, params *StartDevEnvironmentSessionInput, optFns ...func(*Options)) (*StartDevEnvironmentSessionOutput, error) 
 StartWorkflowRun(ctx context.Context, params *StartWorkflowRunInput, optFns ...func(*Options)) (*StartWorkflowRunOutput, error) 
 StopDevEnvironment(ctx context.Context, params *StopDevEnvironmentInput, optFns ...func(*Options)) (*StopDevEnvironmentOutput, error) 
 StopDevEnvironmentSession(ctx context.Context, params *StopDevEnvironmentSessionInput, optFns ...func(*Options)) (*StopDevEnvironmentSessionOutput, error) 
 UpdateDevEnvironment(ctx context.Context, params *UpdateDevEnvironmentInput, optFns ...func(*Options)) (*UpdateDevEnvironmentOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
 UpdateSpace(ctx context.Context, params *UpdateSpaceInput, optFns ...func(*Options)) (*UpdateSpaceOutput, error) 
 VerifySession(ctx context.Context, params *VerifySessionInput, optFns ...func(*Options)) (*VerifySessionOutput, error) 
}
