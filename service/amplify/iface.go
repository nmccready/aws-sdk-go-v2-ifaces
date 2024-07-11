
package amplify

import (
    "github.com/aws/aws-sdk-go-v2/service/amplify"
)

// IClient defines the interface for amplify
type IClient interface {
 Options() Options 
 CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) 
 CreateBackendEnvironment(ctx context.Context, params *CreateBackendEnvironmentInput, optFns ...func(*Options)) (*CreateBackendEnvironmentOutput, error) 
 CreateBranch(ctx context.Context, params *CreateBranchInput, optFns ...func(*Options)) (*CreateBranchOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 CreateDomainAssociation(ctx context.Context, params *CreateDomainAssociationInput, optFns ...func(*Options)) (*CreateDomainAssociationOutput, error) 
 CreateWebhook(ctx context.Context, params *CreateWebhookInput, optFns ...func(*Options)) (*CreateWebhookOutput, error) 
 DeleteApp(ctx context.Context, params *DeleteAppInput, optFns ...func(*Options)) (*DeleteAppOutput, error) 
 DeleteBackendEnvironment(ctx context.Context, params *DeleteBackendEnvironmentInput, optFns ...func(*Options)) (*DeleteBackendEnvironmentOutput, error) 
 DeleteBranch(ctx context.Context, params *DeleteBranchInput, optFns ...func(*Options)) (*DeleteBranchOutput, error) 
 DeleteDomainAssociation(ctx context.Context, params *DeleteDomainAssociationInput, optFns ...func(*Options)) (*DeleteDomainAssociationOutput, error) 
 DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) 
 DeleteWebhook(ctx context.Context, params *DeleteWebhookInput, optFns ...func(*Options)) (*DeleteWebhookOutput, error) 
 GenerateAccessLogs(ctx context.Context, params *GenerateAccessLogsInput, optFns ...func(*Options)) (*GenerateAccessLogsOutput, error) 
 GetApp(ctx context.Context, params *GetAppInput, optFns ...func(*Options)) (*GetAppOutput, error) 
 GetArtifactUrl(ctx context.Context, params *GetArtifactUrlInput, optFns ...func(*Options)) (*GetArtifactUrlOutput, error) 
 GetBackendEnvironment(ctx context.Context, params *GetBackendEnvironmentInput, optFns ...func(*Options)) (*GetBackendEnvironmentOutput, error) 
 GetBranch(ctx context.Context, params *GetBranchInput, optFns ...func(*Options)) (*GetBranchOutput, error) 
 GetDomainAssociation(ctx context.Context, params *GetDomainAssociationInput, optFns ...func(*Options)) (*GetDomainAssociationOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetWebhook(ctx context.Context, params *GetWebhookInput, optFns ...func(*Options)) (*GetWebhookOutput, error) 
 ListApps(ctx context.Context, params *ListAppsInput, optFns ...func(*Options)) (*ListAppsOutput, error) 
 ListArtifacts(ctx context.Context, params *ListArtifactsInput, optFns ...func(*Options)) (*ListArtifactsOutput, error) 
 ListBackendEnvironments(ctx context.Context, params *ListBackendEnvironmentsInput, optFns ...func(*Options)) (*ListBackendEnvironmentsOutput, error) 
 ListBranches(ctx context.Context, params *ListBranchesInput, optFns ...func(*Options)) (*ListBranchesOutput, error) 
 ListDomainAssociations(ctx context.Context, params *ListDomainAssociationsInput, optFns ...func(*Options)) (*ListDomainAssociationsOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebhooks(ctx context.Context, params *ListWebhooksInput, optFns ...func(*Options)) (*ListWebhooksOutput, error) 
 StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) 
 StartJob(ctx context.Context, params *StartJobInput, optFns ...func(*Options)) (*StartJobOutput, error) 
 StopJob(ctx context.Context, params *StopJobInput, optFns ...func(*Options)) (*StopJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) 
 UpdateBranch(ctx context.Context, params *UpdateBranchInput, optFns ...func(*Options)) (*UpdateBranchOutput, error) 
 UpdateDomainAssociation(ctx context.Context, params *UpdateDomainAssociationInput, optFns ...func(*Options)) (*UpdateDomainAssociationOutput, error) 
 UpdateWebhook(ctx context.Context, params *UpdateWebhookInput, optFns ...func(*Options)) (*UpdateWebhookOutput, error) 
}
