package m2_test

// tests for the m2 service interface for this ../../../service/m2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/m2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/m2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/m2/m2_iface"
	"github.com/stretchr/testify/assert"
)

func TestM2ServiceCanBeMocked(t *testing.T) {
	var iface m2_iface.IClient
	iface = &m2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := m2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelBatchJobExecution", func(t *testing.T) {
        input := &m2.CancelBatchJobExecutionInput{}
        output := &m2.CancelBatchJobExecutionOutput{}

        mockClient.On("CancelBatchJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.CancelBatchJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &m2.CreateApplicationInput{}
        output := &m2.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSetImportTask", func(t *testing.T) {
        input := &m2.CreateDataSetImportTaskInput{}
        output := &m2.CreateDataSetImportTaskOutput{}

        mockClient.On("CreateDataSetImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSetImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &m2.CreateDeploymentInput{}
        output := &m2.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &m2.CreateEnvironmentInput{}
        output := &m2.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &m2.DeleteApplicationInput{}
        output := &m2.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationFromEnvironment", func(t *testing.T) {
        input := &m2.DeleteApplicationFromEnvironmentInput{}
        output := &m2.DeleteApplicationFromEnvironmentOutput{}

        mockClient.On("DeleteApplicationFromEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationFromEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &m2.DeleteEnvironmentInput{}
        output := &m2.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &m2.GetApplicationInput{}
        output := &m2.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationVersion", func(t *testing.T) {
        input := &m2.GetApplicationVersionInput{}
        output := &m2.GetApplicationVersionOutput{}

        mockClient.On("GetApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBatchJobExecution", func(t *testing.T) {
        input := &m2.GetBatchJobExecutionInput{}
        output := &m2.GetBatchJobExecutionOutput{}

        mockClient.On("GetBatchJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetBatchJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSetDetails", func(t *testing.T) {
        input := &m2.GetDataSetDetailsInput{}
        output := &m2.GetDataSetDetailsOutput{}

        mockClient.On("GetDataSetDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSetDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSetImportTask", func(t *testing.T) {
        input := &m2.GetDataSetImportTaskInput{}
        output := &m2.GetDataSetImportTaskOutput{}

        mockClient.On("GetDataSetImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSetImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &m2.GetDeploymentInput{}
        output := &m2.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &m2.GetEnvironmentInput{}
        output := &m2.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSignedBluinsightsUrl", func(t *testing.T) {
        input := &m2.GetSignedBluinsightsUrlInput{}
        output := &m2.GetSignedBluinsightsUrlOutput{}

        mockClient.On("GetSignedBluinsightsUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetSignedBluinsightsUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationVersions", func(t *testing.T) {
        input := &m2.ListApplicationVersionsInput{}
        output := &m2.ListApplicationVersionsOutput{}

        mockClient.On("ListApplicationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &m2.ListApplicationsInput{}
        output := &m2.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchJobDefinitions", func(t *testing.T) {
        input := &m2.ListBatchJobDefinitionsInput{}
        output := &m2.ListBatchJobDefinitionsOutput{}

        mockClient.On("ListBatchJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchJobExecutions", func(t *testing.T) {
        input := &m2.ListBatchJobExecutionsInput{}
        output := &m2.ListBatchJobExecutionsOutput{}

        mockClient.On("ListBatchJobExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchJobExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchJobRestartPoints", func(t *testing.T) {
        input := &m2.ListBatchJobRestartPointsInput{}
        output := &m2.ListBatchJobRestartPointsOutput{}

        mockClient.On("ListBatchJobRestartPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchJobRestartPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSetImportHistory", func(t *testing.T) {
        input := &m2.ListDataSetImportHistoryInput{}
        output := &m2.ListDataSetImportHistoryOutput{}

        mockClient.On("ListDataSetImportHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSetImportHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSets", func(t *testing.T) {
        input := &m2.ListDataSetsInput{}
        output := &m2.ListDataSetsOutput{}

        mockClient.On("ListDataSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &m2.ListDeploymentsInput{}
        output := &m2.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEngineVersions", func(t *testing.T) {
        input := &m2.ListEngineVersionsInput{}
        output := &m2.ListEngineVersionsOutput{}

        mockClient.On("ListEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &m2.ListEnvironmentsInput{}
        output := &m2.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &m2.ListTagsForResourceInput{}
        output := &m2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplication", func(t *testing.T) {
        input := &m2.StartApplicationInput{}
        output := &m2.StartApplicationOutput{}

        mockClient.On("StartApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBatchJob", func(t *testing.T) {
        input := &m2.StartBatchJobInput{}
        output := &m2.StartBatchJobOutput{}

        mockClient.On("StartBatchJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartBatchJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApplication", func(t *testing.T) {
        input := &m2.StopApplicationInput{}
        output := &m2.StopApplicationOutput{}

        mockClient.On("StopApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &m2.TagResourceInput{}
        output := &m2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &m2.UntagResourceInput{}
        output := &m2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &m2.UpdateApplicationInput{}
        output := &m2.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &m2.UpdateEnvironmentInput{}
        output := &m2.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
