package lookoutvision_test

// tests for the lookoutvision service interface for this ../../../service/lookoutvision/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lookoutvision"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutvision/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutvision/lookoutvision_iface"
	"github.com/stretchr/testify/assert"
)

func TestLookoutvisionServiceCanBeMocked(t *testing.T) {
	var iface lookoutvision_iface.IClient
	iface = &lookoutvision.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lookoutvision.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &lookoutvision.CreateDatasetInput{}
        output := &lookoutvision.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &lookoutvision.CreateModelInput{}
        output := &lookoutvision.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &lookoutvision.CreateProjectInput{}
        output := &lookoutvision.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &lookoutvision.DeleteDatasetInput{}
        output := &lookoutvision.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &lookoutvision.DeleteModelInput{}
        output := &lookoutvision.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &lookoutvision.DeleteProjectInput{}
        output := &lookoutvision.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &lookoutvision.DescribeDatasetInput{}
        output := &lookoutvision.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModel", func(t *testing.T) {
        input := &lookoutvision.DescribeModelInput{}
        output := &lookoutvision.DescribeModelOutput{}

        mockClient.On("DescribeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelPackagingJob", func(t *testing.T) {
        input := &lookoutvision.DescribeModelPackagingJobInput{}
        output := &lookoutvision.DescribeModelPackagingJobOutput{}

        mockClient.On("DescribeModelPackagingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelPackagingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &lookoutvision.DescribeProjectInput{}
        output := &lookoutvision.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectAnomalies", func(t *testing.T) {
        input := &lookoutvision.DetectAnomaliesInput{}
        output := &lookoutvision.DetectAnomaliesOutput{}

        mockClient.On("DetectAnomalies", ctx, input).Return(output, nil)

        result, err := mockClient.DetectAnomalies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetEntries", func(t *testing.T) {
        input := &lookoutvision.ListDatasetEntriesInput{}
        output := &lookoutvision.ListDatasetEntriesOutput{}

        mockClient.On("ListDatasetEntries", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelPackagingJobs", func(t *testing.T) {
        input := &lookoutvision.ListModelPackagingJobsInput{}
        output := &lookoutvision.ListModelPackagingJobsOutput{}

        mockClient.On("ListModelPackagingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelPackagingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModels", func(t *testing.T) {
        input := &lookoutvision.ListModelsInput{}
        output := &lookoutvision.ListModelsOutput{}

        mockClient.On("ListModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &lookoutvision.ListProjectsInput{}
        output := &lookoutvision.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &lookoutvision.ListTagsForResourceInput{}
        output := &lookoutvision.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartModel", func(t *testing.T) {
        input := &lookoutvision.StartModelInput{}
        output := &lookoutvision.StartModelOutput{}

        mockClient.On("StartModel", ctx, input).Return(output, nil)

        result, err := mockClient.StartModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartModelPackagingJob", func(t *testing.T) {
        input := &lookoutvision.StartModelPackagingJobInput{}
        output := &lookoutvision.StartModelPackagingJobOutput{}

        mockClient.On("StartModelPackagingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartModelPackagingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopModel", func(t *testing.T) {
        input := &lookoutvision.StopModelInput{}
        output := &lookoutvision.StopModelOutput{}

        mockClient.On("StopModel", ctx, input).Return(output, nil)

        result, err := mockClient.StopModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lookoutvision.TagResourceInput{}
        output := &lookoutvision.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lookoutvision.UntagResourceInput{}
        output := &lookoutvision.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatasetEntries", func(t *testing.T) {
        input := &lookoutvision.UpdateDatasetEntriesInput{}
        output := &lookoutvision.UpdateDatasetEntriesOutput{}

        mockClient.On("UpdateDatasetEntries", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatasetEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
