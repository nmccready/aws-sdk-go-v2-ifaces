package lookoutequipment_test

// tests for the lookoutequipment service interface for this ../../../service/lookoutequipment/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutequipment/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutequipment/lookoutequipment_iface"
	"github.com/stretchr/testify/assert"
)

func TestLookoutequipmentServiceCanBeMocked(t *testing.T) {
	var iface lookoutequipment_iface.IClient
	iface = &lookoutequipment.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lookoutequipment.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &lookoutequipment.CreateDatasetInput{}
        output := &lookoutequipment.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.CreateInferenceSchedulerInput{}
        output := &lookoutequipment.CreateInferenceSchedulerOutput{}

        mockClient.On("CreateInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLabel", func(t *testing.T) {
        input := &lookoutequipment.CreateLabelInput{}
        output := &lookoutequipment.CreateLabelOutput{}

        mockClient.On("CreateLabel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLabelGroup", func(t *testing.T) {
        input := &lookoutequipment.CreateLabelGroupInput{}
        output := &lookoutequipment.CreateLabelGroupOutput{}

        mockClient.On("CreateLabelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLabelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &lookoutequipment.CreateModelInput{}
        output := &lookoutequipment.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.CreateRetrainingSchedulerInput{}
        output := &lookoutequipment.CreateRetrainingSchedulerOutput{}

        mockClient.On("CreateRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &lookoutequipment.DeleteDatasetInput{}
        output := &lookoutequipment.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.DeleteInferenceSchedulerInput{}
        output := &lookoutequipment.DeleteInferenceSchedulerOutput{}

        mockClient.On("DeleteInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLabel", func(t *testing.T) {
        input := &lookoutequipment.DeleteLabelInput{}
        output := &lookoutequipment.DeleteLabelOutput{}

        mockClient.On("DeleteLabel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLabelGroup", func(t *testing.T) {
        input := &lookoutequipment.DeleteLabelGroupInput{}
        output := &lookoutequipment.DeleteLabelGroupOutput{}

        mockClient.On("DeleteLabelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLabelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &lookoutequipment.DeleteModelInput{}
        output := &lookoutequipment.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &lookoutequipment.DeleteResourcePolicyInput{}
        output := &lookoutequipment.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.DeleteRetrainingSchedulerInput{}
        output := &lookoutequipment.DeleteRetrainingSchedulerOutput{}

        mockClient.On("DeleteRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataIngestionJob", func(t *testing.T) {
        input := &lookoutequipment.DescribeDataIngestionJobInput{}
        output := &lookoutequipment.DescribeDataIngestionJobOutput{}

        mockClient.On("DescribeDataIngestionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataIngestionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &lookoutequipment.DescribeDatasetInput{}
        output := &lookoutequipment.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.DescribeInferenceSchedulerInput{}
        output := &lookoutequipment.DescribeInferenceSchedulerOutput{}

        mockClient.On("DescribeInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLabel", func(t *testing.T) {
        input := &lookoutequipment.DescribeLabelInput{}
        output := &lookoutequipment.DescribeLabelOutput{}

        mockClient.On("DescribeLabel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLabelGroup", func(t *testing.T) {
        input := &lookoutequipment.DescribeLabelGroupInput{}
        output := &lookoutequipment.DescribeLabelGroupOutput{}

        mockClient.On("DescribeLabelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLabelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModel", func(t *testing.T) {
        input := &lookoutequipment.DescribeModelInput{}
        output := &lookoutequipment.DescribeModelOutput{}

        mockClient.On("DescribeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelVersion", func(t *testing.T) {
        input := &lookoutequipment.DescribeModelVersionInput{}
        output := &lookoutequipment.DescribeModelVersionOutput{}

        mockClient.On("DescribeModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicy", func(t *testing.T) {
        input := &lookoutequipment.DescribeResourcePolicyInput{}
        output := &lookoutequipment.DescribeResourcePolicyOutput{}

        mockClient.On("DescribeResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.DescribeRetrainingSchedulerInput{}
        output := &lookoutequipment.DescribeRetrainingSchedulerOutput{}

        mockClient.On("DescribeRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportDataset", func(t *testing.T) {
        input := &lookoutequipment.ImportDatasetInput{}
        output := &lookoutequipment.ImportDatasetOutput{}

        mockClient.On("ImportDataset", ctx, input).Return(output, nil)

        result, err := mockClient.ImportDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportModelVersion", func(t *testing.T) {
        input := &lookoutequipment.ImportModelVersionInput{}
        output := &lookoutequipment.ImportModelVersionOutput{}

        mockClient.On("ImportModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.ImportModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataIngestionJobs", func(t *testing.T) {
        input := &lookoutequipment.ListDataIngestionJobsInput{}
        output := &lookoutequipment.ListDataIngestionJobsOutput{}

        mockClient.On("ListDataIngestionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataIngestionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &lookoutequipment.ListDatasetsInput{}
        output := &lookoutequipment.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceEvents", func(t *testing.T) {
        input := &lookoutequipment.ListInferenceEventsInput{}
        output := &lookoutequipment.ListInferenceEventsOutput{}

        mockClient.On("ListInferenceEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceExecutions", func(t *testing.T) {
        input := &lookoutequipment.ListInferenceExecutionsInput{}
        output := &lookoutequipment.ListInferenceExecutionsOutput{}

        mockClient.On("ListInferenceExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceSchedulers", func(t *testing.T) {
        input := &lookoutequipment.ListInferenceSchedulersInput{}
        output := &lookoutequipment.ListInferenceSchedulersOutput{}

        mockClient.On("ListInferenceSchedulers", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceSchedulers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLabelGroups", func(t *testing.T) {
        input := &lookoutequipment.ListLabelGroupsInput{}
        output := &lookoutequipment.ListLabelGroupsOutput{}

        mockClient.On("ListLabelGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListLabelGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLabels", func(t *testing.T) {
        input := &lookoutequipment.ListLabelsInput{}
        output := &lookoutequipment.ListLabelsOutput{}

        mockClient.On("ListLabels", ctx, input).Return(output, nil)

        result, err := mockClient.ListLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelVersions", func(t *testing.T) {
        input := &lookoutequipment.ListModelVersionsInput{}
        output := &lookoutequipment.ListModelVersionsOutput{}

        mockClient.On("ListModelVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModels", func(t *testing.T) {
        input := &lookoutequipment.ListModelsInput{}
        output := &lookoutequipment.ListModelsOutput{}

        mockClient.On("ListModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRetrainingSchedulers", func(t *testing.T) {
        input := &lookoutequipment.ListRetrainingSchedulersInput{}
        output := &lookoutequipment.ListRetrainingSchedulersOutput{}

        mockClient.On("ListRetrainingSchedulers", ctx, input).Return(output, nil)

        result, err := mockClient.ListRetrainingSchedulers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSensorStatistics", func(t *testing.T) {
        input := &lookoutequipment.ListSensorStatisticsInput{}
        output := &lookoutequipment.ListSensorStatisticsOutput{}

        mockClient.On("ListSensorStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.ListSensorStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &lookoutequipment.ListTagsForResourceInput{}
        output := &lookoutequipment.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &lookoutequipment.PutResourcePolicyInput{}
        output := &lookoutequipment.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataIngestionJob", func(t *testing.T) {
        input := &lookoutequipment.StartDataIngestionJobInput{}
        output := &lookoutequipment.StartDataIngestionJobOutput{}

        mockClient.On("StartDataIngestionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataIngestionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.StartInferenceSchedulerInput{}
        output := &lookoutequipment.StartInferenceSchedulerOutput{}

        mockClient.On("StartInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.StartInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.StartRetrainingSchedulerInput{}
        output := &lookoutequipment.StartRetrainingSchedulerOutput{}

        mockClient.On("StartRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.StartRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.StopInferenceSchedulerInput{}
        output := &lookoutequipment.StopInferenceSchedulerOutput{}

        mockClient.On("StopInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.StopInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.StopRetrainingSchedulerInput{}
        output := &lookoutequipment.StopRetrainingSchedulerOutput{}

        mockClient.On("StopRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.StopRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lookoutequipment.TagResourceInput{}
        output := &lookoutequipment.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lookoutequipment.UntagResourceInput{}
        output := &lookoutequipment.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateActiveModelVersion", func(t *testing.T) {
        input := &lookoutequipment.UpdateActiveModelVersionInput{}
        output := &lookoutequipment.UpdateActiveModelVersionOutput{}

        mockClient.On("UpdateActiveModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateActiveModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInferenceScheduler", func(t *testing.T) {
        input := &lookoutequipment.UpdateInferenceSchedulerInput{}
        output := &lookoutequipment.UpdateInferenceSchedulerOutput{}

        mockClient.On("UpdateInferenceScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInferenceScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLabelGroup", func(t *testing.T) {
        input := &lookoutequipment.UpdateLabelGroupInput{}
        output := &lookoutequipment.UpdateLabelGroupOutput{}

        mockClient.On("UpdateLabelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLabelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModel", func(t *testing.T) {
        input := &lookoutequipment.UpdateModelInput{}
        output := &lookoutequipment.UpdateModelOutput{}

        mockClient.On("UpdateModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRetrainingScheduler", func(t *testing.T) {
        input := &lookoutequipment.UpdateRetrainingSchedulerInput{}
        output := &lookoutequipment.UpdateRetrainingSchedulerOutput{}

        mockClient.On("UpdateRetrainingScheduler", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRetrainingScheduler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
