package datasync_test

// tests for the datasync service interface for this ../../../service/datasync/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datasync/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datasync/datasync_iface"
	"github.com/stretchr/testify/assert"
)

func TestDatasyncServiceCanBeMocked(t *testing.T) {
	var iface datasync_iface.IClient
	iface = &datasync.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := datasync.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddStorageSystem", func(t *testing.T) {
        input := &datasync.AddStorageSystemInput{}
        output := &datasync.AddStorageSystemOutput{}

        mockClient.On("AddStorageSystem", ctx, input).Return(output, nil)

        result, err := mockClient.AddStorageSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelTaskExecution", func(t *testing.T) {
        input := &datasync.CancelTaskExecutionInput{}
        output := &datasync.CancelTaskExecutionOutput{}

        mockClient.On("CancelTaskExecution", ctx, input).Return(output, nil)

        result, err := mockClient.CancelTaskExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgent", func(t *testing.T) {
        input := &datasync.CreateAgentInput{}
        output := &datasync.CreateAgentOutput{}

        mockClient.On("CreateAgent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationAzureBlob", func(t *testing.T) {
        input := &datasync.CreateLocationAzureBlobInput{}
        output := &datasync.CreateLocationAzureBlobOutput{}

        mockClient.On("CreateLocationAzureBlob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationAzureBlob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationEfs", func(t *testing.T) {
        input := &datasync.CreateLocationEfsInput{}
        output := &datasync.CreateLocationEfsOutput{}

        mockClient.On("CreateLocationEfs", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationEfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationFsxLustre", func(t *testing.T) {
        input := &datasync.CreateLocationFsxLustreInput{}
        output := &datasync.CreateLocationFsxLustreOutput{}

        mockClient.On("CreateLocationFsxLustre", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationFsxLustre(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationFsxOntap", func(t *testing.T) {
        input := &datasync.CreateLocationFsxOntapInput{}
        output := &datasync.CreateLocationFsxOntapOutput{}

        mockClient.On("CreateLocationFsxOntap", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationFsxOntap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationFsxOpenZfs", func(t *testing.T) {
        input := &datasync.CreateLocationFsxOpenZfsInput{}
        output := &datasync.CreateLocationFsxOpenZfsOutput{}

        mockClient.On("CreateLocationFsxOpenZfs", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationFsxOpenZfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationFsxWindows", func(t *testing.T) {
        input := &datasync.CreateLocationFsxWindowsInput{}
        output := &datasync.CreateLocationFsxWindowsOutput{}

        mockClient.On("CreateLocationFsxWindows", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationFsxWindows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationHdfs", func(t *testing.T) {
        input := &datasync.CreateLocationHdfsInput{}
        output := &datasync.CreateLocationHdfsOutput{}

        mockClient.On("CreateLocationHdfs", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationHdfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationNfs", func(t *testing.T) {
        input := &datasync.CreateLocationNfsInput{}
        output := &datasync.CreateLocationNfsOutput{}

        mockClient.On("CreateLocationNfs", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationNfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationObjectStorage", func(t *testing.T) {
        input := &datasync.CreateLocationObjectStorageInput{}
        output := &datasync.CreateLocationObjectStorageOutput{}

        mockClient.On("CreateLocationObjectStorage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationObjectStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationS3", func(t *testing.T) {
        input := &datasync.CreateLocationS3Input{}
        output := &datasync.CreateLocationS3Output{}

        mockClient.On("CreateLocationS3", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocationSmb", func(t *testing.T) {
        input := &datasync.CreateLocationSmbInput{}
        output := &datasync.CreateLocationSmbOutput{}

        mockClient.On("CreateLocationSmb", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocationSmb(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTask", func(t *testing.T) {
        input := &datasync.CreateTaskInput{}
        output := &datasync.CreateTaskOutput{}

        mockClient.On("CreateTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgent", func(t *testing.T) {
        input := &datasync.DeleteAgentInput{}
        output := &datasync.DeleteAgentOutput{}

        mockClient.On("DeleteAgent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocation", func(t *testing.T) {
        input := &datasync.DeleteLocationInput{}
        output := &datasync.DeleteLocationOutput{}

        mockClient.On("DeleteLocation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTask", func(t *testing.T) {
        input := &datasync.DeleteTaskInput{}
        output := &datasync.DeleteTaskOutput{}

        mockClient.On("DeleteTask", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgent", func(t *testing.T) {
        input := &datasync.DescribeAgentInput{}
        output := &datasync.DescribeAgentOutput{}

        mockClient.On("DescribeAgent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDiscoveryJob", func(t *testing.T) {
        input := &datasync.DescribeDiscoveryJobInput{}
        output := &datasync.DescribeDiscoveryJobOutput{}

        mockClient.On("DescribeDiscoveryJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDiscoveryJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationAzureBlob", func(t *testing.T) {
        input := &datasync.DescribeLocationAzureBlobInput{}
        output := &datasync.DescribeLocationAzureBlobOutput{}

        mockClient.On("DescribeLocationAzureBlob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationAzureBlob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationEfs", func(t *testing.T) {
        input := &datasync.DescribeLocationEfsInput{}
        output := &datasync.DescribeLocationEfsOutput{}

        mockClient.On("DescribeLocationEfs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationEfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationFsxLustre", func(t *testing.T) {
        input := &datasync.DescribeLocationFsxLustreInput{}
        output := &datasync.DescribeLocationFsxLustreOutput{}

        mockClient.On("DescribeLocationFsxLustre", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationFsxLustre(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationFsxOntap", func(t *testing.T) {
        input := &datasync.DescribeLocationFsxOntapInput{}
        output := &datasync.DescribeLocationFsxOntapOutput{}

        mockClient.On("DescribeLocationFsxOntap", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationFsxOntap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationFsxOpenZfs", func(t *testing.T) {
        input := &datasync.DescribeLocationFsxOpenZfsInput{}
        output := &datasync.DescribeLocationFsxOpenZfsOutput{}

        mockClient.On("DescribeLocationFsxOpenZfs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationFsxOpenZfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationFsxWindows", func(t *testing.T) {
        input := &datasync.DescribeLocationFsxWindowsInput{}
        output := &datasync.DescribeLocationFsxWindowsOutput{}

        mockClient.On("DescribeLocationFsxWindows", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationFsxWindows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationHdfs", func(t *testing.T) {
        input := &datasync.DescribeLocationHdfsInput{}
        output := &datasync.DescribeLocationHdfsOutput{}

        mockClient.On("DescribeLocationHdfs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationHdfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationNfs", func(t *testing.T) {
        input := &datasync.DescribeLocationNfsInput{}
        output := &datasync.DescribeLocationNfsOutput{}

        mockClient.On("DescribeLocationNfs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationNfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationObjectStorage", func(t *testing.T) {
        input := &datasync.DescribeLocationObjectStorageInput{}
        output := &datasync.DescribeLocationObjectStorageOutput{}

        mockClient.On("DescribeLocationObjectStorage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationObjectStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationS3", func(t *testing.T) {
        input := &datasync.DescribeLocationS3Input{}
        output := &datasync.DescribeLocationS3Output{}

        mockClient.On("DescribeLocationS3", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLocationSmb", func(t *testing.T) {
        input := &datasync.DescribeLocationSmbInput{}
        output := &datasync.DescribeLocationSmbOutput{}

        mockClient.On("DescribeLocationSmb", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLocationSmb(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorageSystem", func(t *testing.T) {
        input := &datasync.DescribeStorageSystemInput{}
        output := &datasync.DescribeStorageSystemOutput{}

        mockClient.On("DescribeStorageSystem", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorageSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorageSystemResourceMetrics", func(t *testing.T) {
        input := &datasync.DescribeStorageSystemResourceMetricsInput{}
        output := &datasync.DescribeStorageSystemResourceMetricsOutput{}

        mockClient.On("DescribeStorageSystemResourceMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorageSystemResourceMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorageSystemResources", func(t *testing.T) {
        input := &datasync.DescribeStorageSystemResourcesInput{}
        output := &datasync.DescribeStorageSystemResourcesOutput{}

        mockClient.On("DescribeStorageSystemResources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorageSystemResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTask", func(t *testing.T) {
        input := &datasync.DescribeTaskInput{}
        output := &datasync.DescribeTaskOutput{}

        mockClient.On("DescribeTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTaskExecution", func(t *testing.T) {
        input := &datasync.DescribeTaskExecutionInput{}
        output := &datasync.DescribeTaskExecutionOutput{}

        mockClient.On("DescribeTaskExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTaskExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateRecommendations", func(t *testing.T) {
        input := &datasync.GenerateRecommendationsInput{}
        output := &datasync.GenerateRecommendationsOutput{}

        mockClient.On("GenerateRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgents", func(t *testing.T) {
        input := &datasync.ListAgentsInput{}
        output := &datasync.ListAgentsOutput{}

        mockClient.On("ListAgents", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDiscoveryJobs", func(t *testing.T) {
        input := &datasync.ListDiscoveryJobsInput{}
        output := &datasync.ListDiscoveryJobsOutput{}

        mockClient.On("ListDiscoveryJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDiscoveryJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLocations", func(t *testing.T) {
        input := &datasync.ListLocationsInput{}
        output := &datasync.ListLocationsOutput{}

        mockClient.On("ListLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageSystems", func(t *testing.T) {
        input := &datasync.ListStorageSystemsInput{}
        output := &datasync.ListStorageSystemsOutput{}

        mockClient.On("ListStorageSystems", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageSystems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &datasync.ListTagsForResourceInput{}
        output := &datasync.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaskExecutions", func(t *testing.T) {
        input := &datasync.ListTaskExecutionsInput{}
        output := &datasync.ListTaskExecutionsOutput{}

        mockClient.On("ListTaskExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaskExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTasks", func(t *testing.T) {
        input := &datasync.ListTasksInput{}
        output := &datasync.ListTasksOutput{}

        mockClient.On("ListTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveStorageSystem", func(t *testing.T) {
        input := &datasync.RemoveStorageSystemInput{}
        output := &datasync.RemoveStorageSystemOutput{}

        mockClient.On("RemoveStorageSystem", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveStorageSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDiscoveryJob", func(t *testing.T) {
        input := &datasync.StartDiscoveryJobInput{}
        output := &datasync.StartDiscoveryJobOutput{}

        mockClient.On("StartDiscoveryJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDiscoveryJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTaskExecution", func(t *testing.T) {
        input := &datasync.StartTaskExecutionInput{}
        output := &datasync.StartTaskExecutionOutput{}

        mockClient.On("StartTaskExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartTaskExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDiscoveryJob", func(t *testing.T) {
        input := &datasync.StopDiscoveryJobInput{}
        output := &datasync.StopDiscoveryJobOutput{}

        mockClient.On("StopDiscoveryJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopDiscoveryJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &datasync.TagResourceInput{}
        output := &datasync.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &datasync.UntagResourceInput{}
        output := &datasync.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgent", func(t *testing.T) {
        input := &datasync.UpdateAgentInput{}
        output := &datasync.UpdateAgentOutput{}

        mockClient.On("UpdateAgent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDiscoveryJob", func(t *testing.T) {
        input := &datasync.UpdateDiscoveryJobInput{}
        output := &datasync.UpdateDiscoveryJobOutput{}

        mockClient.On("UpdateDiscoveryJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDiscoveryJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLocationAzureBlob", func(t *testing.T) {
        input := &datasync.UpdateLocationAzureBlobInput{}
        output := &datasync.UpdateLocationAzureBlobOutput{}

        mockClient.On("UpdateLocationAzureBlob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLocationAzureBlob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLocationHdfs", func(t *testing.T) {
        input := &datasync.UpdateLocationHdfsInput{}
        output := &datasync.UpdateLocationHdfsOutput{}

        mockClient.On("UpdateLocationHdfs", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLocationHdfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLocationNfs", func(t *testing.T) {
        input := &datasync.UpdateLocationNfsInput{}
        output := &datasync.UpdateLocationNfsOutput{}

        mockClient.On("UpdateLocationNfs", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLocationNfs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLocationObjectStorage", func(t *testing.T) {
        input := &datasync.UpdateLocationObjectStorageInput{}
        output := &datasync.UpdateLocationObjectStorageOutput{}

        mockClient.On("UpdateLocationObjectStorage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLocationObjectStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLocationSmb", func(t *testing.T) {
        input := &datasync.UpdateLocationSmbInput{}
        output := &datasync.UpdateLocationSmbOutput{}

        mockClient.On("UpdateLocationSmb", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLocationSmb(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStorageSystem", func(t *testing.T) {
        input := &datasync.UpdateStorageSystemInput{}
        output := &datasync.UpdateStorageSystemOutput{}

        mockClient.On("UpdateStorageSystem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStorageSystem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTask", func(t *testing.T) {
        input := &datasync.UpdateTaskInput{}
        output := &datasync.UpdateTaskOutput{}

        mockClient.On("UpdateTask", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTaskExecution", func(t *testing.T) {
        input := &datasync.UpdateTaskExecutionInput{}
        output := &datasync.UpdateTaskExecutionOutput{}

        mockClient.On("UpdateTaskExecution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTaskExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
