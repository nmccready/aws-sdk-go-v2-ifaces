package codebuild_test

// tests for the codebuild service interface for this ../../../service/codebuild/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codebuild/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codebuild/codebuild_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodebuildServiceCanBeMocked(t *testing.T) {
	var iface codebuild_iface.IClient
	iface = &codebuild.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codebuild.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteBuilds", func(t *testing.T) {
        input := &codebuild.BatchDeleteBuildsInput{}
        output := &codebuild.BatchDeleteBuildsOutput{}

        mockClient.On("BatchDeleteBuilds", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteBuilds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetBuildBatches", func(t *testing.T) {
        input := &codebuild.BatchGetBuildBatchesInput{}
        output := &codebuild.BatchGetBuildBatchesOutput{}

        mockClient.On("BatchGetBuildBatches", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetBuildBatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetBuilds", func(t *testing.T) {
        input := &codebuild.BatchGetBuildsInput{}
        output := &codebuild.BatchGetBuildsOutput{}

        mockClient.On("BatchGetBuilds", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetBuilds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFleets", func(t *testing.T) {
        input := &codebuild.BatchGetFleetsInput{}
        output := &codebuild.BatchGetFleetsOutput{}

        mockClient.On("BatchGetFleets", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetProjects", func(t *testing.T) {
        input := &codebuild.BatchGetProjectsInput{}
        output := &codebuild.BatchGetProjectsOutput{}

        mockClient.On("BatchGetProjects", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetReportGroups", func(t *testing.T) {
        input := &codebuild.BatchGetReportGroupsInput{}
        output := &codebuild.BatchGetReportGroupsOutput{}

        mockClient.On("BatchGetReportGroups", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetReportGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetReports", func(t *testing.T) {
        input := &codebuild.BatchGetReportsInput{}
        output := &codebuild.BatchGetReportsOutput{}

        mockClient.On("BatchGetReports", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &codebuild.CreateFleetInput{}
        output := &codebuild.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &codebuild.CreateProjectInput{}
        output := &codebuild.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReportGroup", func(t *testing.T) {
        input := &codebuild.CreateReportGroupInput{}
        output := &codebuild.CreateReportGroupOutput{}

        mockClient.On("CreateReportGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReportGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebhook", func(t *testing.T) {
        input := &codebuild.CreateWebhookInput{}
        output := &codebuild.CreateWebhookOutput{}

        mockClient.On("CreateWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBuildBatch", func(t *testing.T) {
        input := &codebuild.DeleteBuildBatchInput{}
        output := &codebuild.DeleteBuildBatchOutput{}

        mockClient.On("DeleteBuildBatch", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBuildBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &codebuild.DeleteFleetInput{}
        output := &codebuild.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &codebuild.DeleteProjectInput{}
        output := &codebuild.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReport", func(t *testing.T) {
        input := &codebuild.DeleteReportInput{}
        output := &codebuild.DeleteReportOutput{}

        mockClient.On("DeleteReport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReportGroup", func(t *testing.T) {
        input := &codebuild.DeleteReportGroupInput{}
        output := &codebuild.DeleteReportGroupOutput{}

        mockClient.On("DeleteReportGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReportGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &codebuild.DeleteResourcePolicyInput{}
        output := &codebuild.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceCredentials", func(t *testing.T) {
        input := &codebuild.DeleteSourceCredentialsInput{}
        output := &codebuild.DeleteSourceCredentialsOutput{}

        mockClient.On("DeleteSourceCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebhook", func(t *testing.T) {
        input := &codebuild.DeleteWebhookInput{}
        output := &codebuild.DeleteWebhookOutput{}

        mockClient.On("DeleteWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCodeCoverages", func(t *testing.T) {
        input := &codebuild.DescribeCodeCoveragesInput{}
        output := &codebuild.DescribeCodeCoveragesOutput{}

        mockClient.On("DescribeCodeCoverages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCodeCoverages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTestCases", func(t *testing.T) {
        input := &codebuild.DescribeTestCasesInput{}
        output := &codebuild.DescribeTestCasesOutput{}

        mockClient.On("DescribeTestCases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTestCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReportGroupTrend", func(t *testing.T) {
        input := &codebuild.GetReportGroupTrendInput{}
        output := &codebuild.GetReportGroupTrendOutput{}

        mockClient.On("GetReportGroupTrend", ctx, input).Return(output, nil)

        result, err := mockClient.GetReportGroupTrend(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &codebuild.GetResourcePolicyInput{}
        output := &codebuild.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportSourceCredentials", func(t *testing.T) {
        input := &codebuild.ImportSourceCredentialsInput{}
        output := &codebuild.ImportSourceCredentialsOutput{}

        mockClient.On("ImportSourceCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.ImportSourceCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvalidateProjectCache", func(t *testing.T) {
        input := &codebuild.InvalidateProjectCacheInput{}
        output := &codebuild.InvalidateProjectCacheOutput{}

        mockClient.On("InvalidateProjectCache", ctx, input).Return(output, nil)

        result, err := mockClient.InvalidateProjectCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuildBatches", func(t *testing.T) {
        input := &codebuild.ListBuildBatchesInput{}
        output := &codebuild.ListBuildBatchesOutput{}

        mockClient.On("ListBuildBatches", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuildBatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuildBatchesForProject", func(t *testing.T) {
        input := &codebuild.ListBuildBatchesForProjectInput{}
        output := &codebuild.ListBuildBatchesForProjectOutput{}

        mockClient.On("ListBuildBatchesForProject", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuildBatchesForProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuilds", func(t *testing.T) {
        input := &codebuild.ListBuildsInput{}
        output := &codebuild.ListBuildsOutput{}

        mockClient.On("ListBuilds", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuilds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuildsForProject", func(t *testing.T) {
        input := &codebuild.ListBuildsForProjectInput{}
        output := &codebuild.ListBuildsForProjectOutput{}

        mockClient.On("ListBuildsForProject", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuildsForProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCuratedEnvironmentImages", func(t *testing.T) {
        input := &codebuild.ListCuratedEnvironmentImagesInput{}
        output := &codebuild.ListCuratedEnvironmentImagesOutput{}

        mockClient.On("ListCuratedEnvironmentImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListCuratedEnvironmentImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &codebuild.ListFleetsInput{}
        output := &codebuild.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &codebuild.ListProjectsInput{}
        output := &codebuild.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReportGroups", func(t *testing.T) {
        input := &codebuild.ListReportGroupsInput{}
        output := &codebuild.ListReportGroupsOutput{}

        mockClient.On("ListReportGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListReportGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReports", func(t *testing.T) {
        input := &codebuild.ListReportsInput{}
        output := &codebuild.ListReportsOutput{}

        mockClient.On("ListReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReportsForReportGroup", func(t *testing.T) {
        input := &codebuild.ListReportsForReportGroupInput{}
        output := &codebuild.ListReportsForReportGroupOutput{}

        mockClient.On("ListReportsForReportGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListReportsForReportGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSharedProjects", func(t *testing.T) {
        input := &codebuild.ListSharedProjectsInput{}
        output := &codebuild.ListSharedProjectsOutput{}

        mockClient.On("ListSharedProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListSharedProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSharedReportGroups", func(t *testing.T) {
        input := &codebuild.ListSharedReportGroupsInput{}
        output := &codebuild.ListSharedReportGroupsOutput{}

        mockClient.On("ListSharedReportGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListSharedReportGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceCredentials", func(t *testing.T) {
        input := &codebuild.ListSourceCredentialsInput{}
        output := &codebuild.ListSourceCredentialsOutput{}

        mockClient.On("ListSourceCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &codebuild.PutResourcePolicyInput{}
        output := &codebuild.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryBuild", func(t *testing.T) {
        input := &codebuild.RetryBuildInput{}
        output := &codebuild.RetryBuildOutput{}

        mockClient.On("RetryBuild", ctx, input).Return(output, nil)

        result, err := mockClient.RetryBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryBuildBatch", func(t *testing.T) {
        input := &codebuild.RetryBuildBatchInput{}
        output := &codebuild.RetryBuildBatchOutput{}

        mockClient.On("RetryBuildBatch", ctx, input).Return(output, nil)

        result, err := mockClient.RetryBuildBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBuild", func(t *testing.T) {
        input := &codebuild.StartBuildInput{}
        output := &codebuild.StartBuildOutput{}

        mockClient.On("StartBuild", ctx, input).Return(output, nil)

        result, err := mockClient.StartBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBuildBatch", func(t *testing.T) {
        input := &codebuild.StartBuildBatchInput{}
        output := &codebuild.StartBuildBatchOutput{}

        mockClient.On("StartBuildBatch", ctx, input).Return(output, nil)

        result, err := mockClient.StartBuildBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBuild", func(t *testing.T) {
        input := &codebuild.StopBuildInput{}
        output := &codebuild.StopBuildOutput{}

        mockClient.On("StopBuild", ctx, input).Return(output, nil)

        result, err := mockClient.StopBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBuildBatch", func(t *testing.T) {
        input := &codebuild.StopBuildBatchInput{}
        output := &codebuild.StopBuildBatchOutput{}

        mockClient.On("StopBuildBatch", ctx, input).Return(output, nil)

        result, err := mockClient.StopBuildBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleet", func(t *testing.T) {
        input := &codebuild.UpdateFleetInput{}
        output := &codebuild.UpdateFleetOutput{}

        mockClient.On("UpdateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &codebuild.UpdateProjectInput{}
        output := &codebuild.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProjectVisibility", func(t *testing.T) {
        input := &codebuild.UpdateProjectVisibilityInput{}
        output := &codebuild.UpdateProjectVisibilityOutput{}

        mockClient.On("UpdateProjectVisibility", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProjectVisibility(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReportGroup", func(t *testing.T) {
        input := &codebuild.UpdateReportGroupInput{}
        output := &codebuild.UpdateReportGroupOutput{}

        mockClient.On("UpdateReportGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReportGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebhook", func(t *testing.T) {
        input := &codebuild.UpdateWebhookInput{}
        output := &codebuild.UpdateWebhookOutput{}

        mockClient.On("UpdateWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
