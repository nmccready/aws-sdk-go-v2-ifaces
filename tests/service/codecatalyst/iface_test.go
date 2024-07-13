package codecatalyst_test

// tests for the codecatalyst service interface for this ../../../service/codecatalyst/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codecatalyst"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codecatalyst/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codecatalyst/codecatalyst_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodecatalystServiceCanBeMocked(t *testing.T) {
	var iface codecatalyst_iface.IClient
	iface = &codecatalyst.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codecatalyst.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessToken", func(t *testing.T) {
        input := &codecatalyst.CreateAccessTokenInput{}
        output := &codecatalyst.CreateAccessTokenOutput{}

        mockClient.On("CreateAccessToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.CreateDevEnvironmentInput{}
        output := &codecatalyst.CreateDevEnvironmentOutput{}

        mockClient.On("CreateDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &codecatalyst.CreateProjectInput{}
        output := &codecatalyst.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSourceRepository", func(t *testing.T) {
        input := &codecatalyst.CreateSourceRepositoryInput{}
        output := &codecatalyst.CreateSourceRepositoryOutput{}

        mockClient.On("CreateSourceRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSourceRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSourceRepositoryBranch", func(t *testing.T) {
        input := &codecatalyst.CreateSourceRepositoryBranchInput{}
        output := &codecatalyst.CreateSourceRepositoryBranchOutput{}

        mockClient.On("CreateSourceRepositoryBranch", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSourceRepositoryBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessToken", func(t *testing.T) {
        input := &codecatalyst.DeleteAccessTokenInput{}
        output := &codecatalyst.DeleteAccessTokenOutput{}

        mockClient.On("DeleteAccessToken", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.DeleteDevEnvironmentInput{}
        output := &codecatalyst.DeleteDevEnvironmentOutput{}

        mockClient.On("DeleteDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &codecatalyst.DeleteProjectInput{}
        output := &codecatalyst.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceRepository", func(t *testing.T) {
        input := &codecatalyst.DeleteSourceRepositoryInput{}
        output := &codecatalyst.DeleteSourceRepositoryOutput{}

        mockClient.On("DeleteSourceRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSpace", func(t *testing.T) {
        input := &codecatalyst.DeleteSpaceInput{}
        output := &codecatalyst.DeleteSpaceOutput{}

        mockClient.On("DeleteSpace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.GetDevEnvironmentInput{}
        output := &codecatalyst.GetDevEnvironmentOutput{}

        mockClient.On("GetDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProject", func(t *testing.T) {
        input := &codecatalyst.GetProjectInput{}
        output := &codecatalyst.GetProjectOutput{}

        mockClient.On("GetProject", ctx, input).Return(output, nil)

        result, err := mockClient.GetProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSourceRepository", func(t *testing.T) {
        input := &codecatalyst.GetSourceRepositoryInput{}
        output := &codecatalyst.GetSourceRepositoryOutput{}

        mockClient.On("GetSourceRepository", ctx, input).Return(output, nil)

        result, err := mockClient.GetSourceRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSourceRepositoryCloneUrls", func(t *testing.T) {
        input := &codecatalyst.GetSourceRepositoryCloneUrlsInput{}
        output := &codecatalyst.GetSourceRepositoryCloneUrlsOutput{}

        mockClient.On("GetSourceRepositoryCloneUrls", ctx, input).Return(output, nil)

        result, err := mockClient.GetSourceRepositoryCloneUrls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpace", func(t *testing.T) {
        input := &codecatalyst.GetSpaceInput{}
        output := &codecatalyst.GetSpaceOutput{}

        mockClient.On("GetSpace", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscription", func(t *testing.T) {
        input := &codecatalyst.GetSubscriptionInput{}
        output := &codecatalyst.GetSubscriptionOutput{}

        mockClient.On("GetSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserDetails", func(t *testing.T) {
        input := &codecatalyst.GetUserDetailsInput{}
        output := &codecatalyst.GetUserDetailsOutput{}

        mockClient.On("GetUserDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &codecatalyst.GetWorkflowInput{}
        output := &codecatalyst.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowRun", func(t *testing.T) {
        input := &codecatalyst.GetWorkflowRunInput{}
        output := &codecatalyst.GetWorkflowRunOutput{}

        mockClient.On("GetWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessTokens", func(t *testing.T) {
        input := &codecatalyst.ListAccessTokensInput{}
        output := &codecatalyst.ListAccessTokensOutput{}

        mockClient.On("ListAccessTokens", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessTokens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevEnvironmentSessions", func(t *testing.T) {
        input := &codecatalyst.ListDevEnvironmentSessionsInput{}
        output := &codecatalyst.ListDevEnvironmentSessionsOutput{}

        mockClient.On("ListDevEnvironmentSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevEnvironmentSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevEnvironments", func(t *testing.T) {
        input := &codecatalyst.ListDevEnvironmentsInput{}
        output := &codecatalyst.ListDevEnvironmentsOutput{}

        mockClient.On("ListDevEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventLogs", func(t *testing.T) {
        input := &codecatalyst.ListEventLogsInput{}
        output := &codecatalyst.ListEventLogsOutput{}

        mockClient.On("ListEventLogs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &codecatalyst.ListProjectsInput{}
        output := &codecatalyst.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceRepositories", func(t *testing.T) {
        input := &codecatalyst.ListSourceRepositoriesInput{}
        output := &codecatalyst.ListSourceRepositoriesOutput{}

        mockClient.On("ListSourceRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceRepositoryBranches", func(t *testing.T) {
        input := &codecatalyst.ListSourceRepositoryBranchesInput{}
        output := &codecatalyst.ListSourceRepositoryBranchesOutput{}

        mockClient.On("ListSourceRepositoryBranches", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceRepositoryBranches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpaces", func(t *testing.T) {
        input := &codecatalyst.ListSpacesInput{}
        output := &codecatalyst.ListSpacesOutput{}

        mockClient.On("ListSpaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowRuns", func(t *testing.T) {
        input := &codecatalyst.ListWorkflowRunsInput{}
        output := &codecatalyst.ListWorkflowRunsOutput{}

        mockClient.On("ListWorkflowRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &codecatalyst.ListWorkflowsInput{}
        output := &codecatalyst.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.StartDevEnvironmentInput{}
        output := &codecatalyst.StartDevEnvironmentOutput{}

        mockClient.On("StartDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.StartDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDevEnvironmentSession", func(t *testing.T) {
        input := &codecatalyst.StartDevEnvironmentSessionInput{}
        output := &codecatalyst.StartDevEnvironmentSessionOutput{}

        mockClient.On("StartDevEnvironmentSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartDevEnvironmentSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkflowRun", func(t *testing.T) {
        input := &codecatalyst.StartWorkflowRunInput{}
        output := &codecatalyst.StartWorkflowRunOutput{}

        mockClient.On("StartWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.StopDevEnvironmentInput{}
        output := &codecatalyst.StopDevEnvironmentOutput{}

        mockClient.On("StopDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.StopDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDevEnvironmentSession", func(t *testing.T) {
        input := &codecatalyst.StopDevEnvironmentSessionInput{}
        output := &codecatalyst.StopDevEnvironmentSessionOutput{}

        mockClient.On("StopDevEnvironmentSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopDevEnvironmentSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevEnvironment", func(t *testing.T) {
        input := &codecatalyst.UpdateDevEnvironmentInput{}
        output := &codecatalyst.UpdateDevEnvironmentOutput{}

        mockClient.On("UpdateDevEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &codecatalyst.UpdateProjectInput{}
        output := &codecatalyst.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSpace", func(t *testing.T) {
        input := &codecatalyst.UpdateSpaceInput{}
        output := &codecatalyst.UpdateSpaceOutput{}

        mockClient.On("UpdateSpace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifySession", func(t *testing.T) {
        input := &codecatalyst.VerifySessionInput{}
        output := &codecatalyst.VerifySessionOutput{}

        mockClient.On("VerifySession", ctx, input).Return(output, nil)

        result, err := mockClient.VerifySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
