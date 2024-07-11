
package codebuild

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/codebuild"
)

// IClient defines the interface for codebuild
type IClient interface {
 Options() Options 
 BatchDeleteBuilds(ctx context.Context, params *BatchDeleteBuildsInput, optFns ...func(*Options)) (*BatchDeleteBuildsOutput, error) 
 BatchGetBuildBatches(ctx context.Context, params *BatchGetBuildBatchesInput, optFns ...func(*Options)) (*BatchGetBuildBatchesOutput, error) 
 BatchGetBuilds(ctx context.Context, params *BatchGetBuildsInput, optFns ...func(*Options)) (*BatchGetBuildsOutput, error) 
 BatchGetFleets(ctx context.Context, params *BatchGetFleetsInput, optFns ...func(*Options)) (*BatchGetFleetsOutput, error) 
 BatchGetProjects(ctx context.Context, params *BatchGetProjectsInput, optFns ...func(*Options)) (*BatchGetProjectsOutput, error) 
 BatchGetReportGroups(ctx context.Context, params *BatchGetReportGroupsInput, optFns ...func(*Options)) (*BatchGetReportGroupsOutput, error) 
 BatchGetReports(ctx context.Context, params *BatchGetReportsInput, optFns ...func(*Options)) (*BatchGetReportsOutput, error) 
 CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 CreateReportGroup(ctx context.Context, params *CreateReportGroupInput, optFns ...func(*Options)) (*CreateReportGroupOutput, error) 
 CreateWebhook(ctx context.Context, params *CreateWebhookInput, optFns ...func(*Options)) (*CreateWebhookOutput, error) 
 DeleteBuildBatch(ctx context.Context, params *DeleteBuildBatchInput, optFns ...func(*Options)) (*DeleteBuildBatchOutput, error) 
 DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteReport(ctx context.Context, params *DeleteReportInput, optFns ...func(*Options)) (*DeleteReportOutput, error) 
 DeleteReportGroup(ctx context.Context, params *DeleteReportGroupInput, optFns ...func(*Options)) (*DeleteReportGroupOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteSourceCredentials(ctx context.Context, params *DeleteSourceCredentialsInput, optFns ...func(*Options)) (*DeleteSourceCredentialsOutput, error) 
 DeleteWebhook(ctx context.Context, params *DeleteWebhookInput, optFns ...func(*Options)) (*DeleteWebhookOutput, error) 
 DescribeCodeCoverages(ctx context.Context, params *DescribeCodeCoveragesInput, optFns ...func(*Options)) (*DescribeCodeCoveragesOutput, error) 
 DescribeTestCases(ctx context.Context, params *DescribeTestCasesInput, optFns ...func(*Options)) (*DescribeTestCasesOutput, error) 
 GetReportGroupTrend(ctx context.Context, params *GetReportGroupTrendInput, optFns ...func(*Options)) (*GetReportGroupTrendOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 ImportSourceCredentials(ctx context.Context, params *ImportSourceCredentialsInput, optFns ...func(*Options)) (*ImportSourceCredentialsOutput, error) 
 InvalidateProjectCache(ctx context.Context, params *InvalidateProjectCacheInput, optFns ...func(*Options)) (*InvalidateProjectCacheOutput, error) 
 ListBuildBatches(ctx context.Context, params *ListBuildBatchesInput, optFns ...func(*Options)) (*ListBuildBatchesOutput, error) 
 ListBuildBatchesForProject(ctx context.Context, params *ListBuildBatchesForProjectInput, optFns ...func(*Options)) (*ListBuildBatchesForProjectOutput, error) 
 ListBuilds(ctx context.Context, params *ListBuildsInput, optFns ...func(*Options)) (*ListBuildsOutput, error) 
 ListBuildsForProject(ctx context.Context, params *ListBuildsForProjectInput, optFns ...func(*Options)) (*ListBuildsForProjectOutput, error) 
 ListCuratedEnvironmentImages(ctx context.Context, params *ListCuratedEnvironmentImagesInput, optFns ...func(*Options)) (*ListCuratedEnvironmentImagesOutput, error) 
 ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListReportGroups(ctx context.Context, params *ListReportGroupsInput, optFns ...func(*Options)) (*ListReportGroupsOutput, error) 
 ListReports(ctx context.Context, params *ListReportsInput, optFns ...func(*Options)) (*ListReportsOutput, error) 
 ListReportsForReportGroup(ctx context.Context, params *ListReportsForReportGroupInput, optFns ...func(*Options)) (*ListReportsForReportGroupOutput, error) 
 ListSharedProjects(ctx context.Context, params *ListSharedProjectsInput, optFns ...func(*Options)) (*ListSharedProjectsOutput, error) 
 ListSharedReportGroups(ctx context.Context, params *ListSharedReportGroupsInput, optFns ...func(*Options)) (*ListSharedReportGroupsOutput, error) 
 ListSourceCredentials(ctx context.Context, params *ListSourceCredentialsInput, optFns ...func(*Options)) (*ListSourceCredentialsOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 RetryBuild(ctx context.Context, params *RetryBuildInput, optFns ...func(*Options)) (*RetryBuildOutput, error) 
 RetryBuildBatch(ctx context.Context, params *RetryBuildBatchInput, optFns ...func(*Options)) (*RetryBuildBatchOutput, error) 
 StartBuild(ctx context.Context, params *StartBuildInput, optFns ...func(*Options)) (*StartBuildOutput, error) 
 StartBuildBatch(ctx context.Context, params *StartBuildBatchInput, optFns ...func(*Options)) (*StartBuildBatchOutput, error) 
 StopBuild(ctx context.Context, params *StopBuildInput, optFns ...func(*Options)) (*StopBuildOutput, error) 
 StopBuildBatch(ctx context.Context, params *StopBuildBatchInput, optFns ...func(*Options)) (*StopBuildBatchOutput, error) 
 UpdateFleet(ctx context.Context, params *UpdateFleetInput, optFns ...func(*Options)) (*UpdateFleetOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
 UpdateProjectVisibility(ctx context.Context, params *UpdateProjectVisibilityInput, optFns ...func(*Options)) (*UpdateProjectVisibilityOutput, error) 
 UpdateReportGroup(ctx context.Context, params *UpdateReportGroupInput, optFns ...func(*Options)) (*UpdateReportGroupOutput, error) 
 UpdateWebhook(ctx context.Context, params *UpdateWebhookInput, optFns ...func(*Options)) (*UpdateWebhookOutput, error) 
}
