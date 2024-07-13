package dataexchange_test

// tests for the dataexchange service interface for this ../../../service/dataexchange/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dataexchange"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dataexchange/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dataexchange/dataexchange_iface"
	"github.com/stretchr/testify/assert"
)

func TestDataexchangeServiceCanBeMocked(t *testing.T) {
	var iface dataexchange_iface.IClient
	iface = &dataexchange.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := dataexchange.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &dataexchange.CancelJobInput{}
        output := &dataexchange.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSet", func(t *testing.T) {
        input := &dataexchange.CreateDataSetInput{}
        output := &dataexchange.CreateDataSetOutput{}

        mockClient.On("CreateDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventAction", func(t *testing.T) {
        input := &dataexchange.CreateEventActionInput{}
        output := &dataexchange.CreateEventActionOutput{}

        mockClient.On("CreateEventAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &dataexchange.CreateJobInput{}
        output := &dataexchange.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRevision", func(t *testing.T) {
        input := &dataexchange.CreateRevisionInput{}
        output := &dataexchange.CreateRevisionOutput{}

        mockClient.On("CreateRevision", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAsset", func(t *testing.T) {
        input := &dataexchange.DeleteAssetInput{}
        output := &dataexchange.DeleteAssetOutput{}

        mockClient.On("DeleteAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSet", func(t *testing.T) {
        input := &dataexchange.DeleteDataSetInput{}
        output := &dataexchange.DeleteDataSetOutput{}

        mockClient.On("DeleteDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventAction", func(t *testing.T) {
        input := &dataexchange.DeleteEventActionInput{}
        output := &dataexchange.DeleteEventActionOutput{}

        mockClient.On("DeleteEventAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRevision", func(t *testing.T) {
        input := &dataexchange.DeleteRevisionInput{}
        output := &dataexchange.DeleteRevisionOutput{}

        mockClient.On("DeleteRevision", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAsset", func(t *testing.T) {
        input := &dataexchange.GetAssetInput{}
        output := &dataexchange.GetAssetOutput{}

        mockClient.On("GetAsset", ctx, input).Return(output, nil)

        result, err := mockClient.GetAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSet", func(t *testing.T) {
        input := &dataexchange.GetDataSetInput{}
        output := &dataexchange.GetDataSetOutput{}

        mockClient.On("GetDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventAction", func(t *testing.T) {
        input := &dataexchange.GetEventActionInput{}
        output := &dataexchange.GetEventActionOutput{}

        mockClient.On("GetEventAction", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &dataexchange.GetJobInput{}
        output := &dataexchange.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRevision", func(t *testing.T) {
        input := &dataexchange.GetRevisionInput{}
        output := &dataexchange.GetRevisionOutput{}

        mockClient.On("GetRevision", ctx, input).Return(output, nil)

        result, err := mockClient.GetRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSetRevisions", func(t *testing.T) {
        input := &dataexchange.ListDataSetRevisionsInput{}
        output := &dataexchange.ListDataSetRevisionsOutput{}

        mockClient.On("ListDataSetRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSetRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSets", func(t *testing.T) {
        input := &dataexchange.ListDataSetsInput{}
        output := &dataexchange.ListDataSetsOutput{}

        mockClient.On("ListDataSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventActions", func(t *testing.T) {
        input := &dataexchange.ListEventActionsInput{}
        output := &dataexchange.ListEventActionsOutput{}

        mockClient.On("ListEventActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &dataexchange.ListJobsInput{}
        output := &dataexchange.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRevisionAssets", func(t *testing.T) {
        input := &dataexchange.ListRevisionAssetsInput{}
        output := &dataexchange.ListRevisionAssetsOutput{}

        mockClient.On("ListRevisionAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRevisionAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &dataexchange.ListTagsForResourceInput{}
        output := &dataexchange.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeRevision", func(t *testing.T) {
        input := &dataexchange.RevokeRevisionInput{}
        output := &dataexchange.RevokeRevisionOutput{}

        mockClient.On("RevokeRevision", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendApiAsset", func(t *testing.T) {
        input := &dataexchange.SendApiAssetInput{}
        output := &dataexchange.SendApiAssetOutput{}

        mockClient.On("SendApiAsset", ctx, input).Return(output, nil)

        result, err := mockClient.SendApiAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDataSetNotification", func(t *testing.T) {
        input := &dataexchange.SendDataSetNotificationInput{}
        output := &dataexchange.SendDataSetNotificationOutput{}

        mockClient.On("SendDataSetNotification", ctx, input).Return(output, nil)

        result, err := mockClient.SendDataSetNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJob", func(t *testing.T) {
        input := &dataexchange.StartJobInput{}
        output := &dataexchange.StartJobOutput{}

        mockClient.On("StartJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &dataexchange.TagResourceInput{}
        output := &dataexchange.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &dataexchange.UntagResourceInput{}
        output := &dataexchange.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAsset", func(t *testing.T) {
        input := &dataexchange.UpdateAssetInput{}
        output := &dataexchange.UpdateAssetOutput{}

        mockClient.On("UpdateAsset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSet", func(t *testing.T) {
        input := &dataexchange.UpdateDataSetInput{}
        output := &dataexchange.UpdateDataSetOutput{}

        mockClient.On("UpdateDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventAction", func(t *testing.T) {
        input := &dataexchange.UpdateEventActionInput{}
        output := &dataexchange.UpdateEventActionOutput{}

        mockClient.On("UpdateEventAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRevision", func(t *testing.T) {
        input := &dataexchange.UpdateRevisionInput{}
        output := &dataexchange.UpdateRevisionOutput{}

        mockClient.On("UpdateRevision", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
