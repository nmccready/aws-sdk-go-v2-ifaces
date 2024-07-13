package cognitosync_test

// tests for the cognitosync service interface for this ../../../service/cognitosync/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitosync/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitosync/cognitosync_iface"
	"github.com/stretchr/testify/assert"
)

func TestCognitosyncServiceCanBeMocked(t *testing.T) {
	var iface cognitosync_iface.IClient
	iface = &cognitosync.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cognitosync.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBulkPublish", func(t *testing.T) {
        input := &cognitosync.BulkPublishInput{}
        output := &cognitosync.BulkPublishOutput{}

        mockClient.On("BulkPublish", ctx, input).Return(output, nil)

        result, err := mockClient.BulkPublish(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &cognitosync.DeleteDatasetInput{}
        output := &cognitosync.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &cognitosync.DescribeDatasetInput{}
        output := &cognitosync.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityPoolUsage", func(t *testing.T) {
        input := &cognitosync.DescribeIdentityPoolUsageInput{}
        output := &cognitosync.DescribeIdentityPoolUsageOutput{}

        mockClient.On("DescribeIdentityPoolUsage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityPoolUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityUsage", func(t *testing.T) {
        input := &cognitosync.DescribeIdentityUsageInput{}
        output := &cognitosync.DescribeIdentityUsageOutput{}

        mockClient.On("DescribeIdentityUsage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBulkPublishDetails", func(t *testing.T) {
        input := &cognitosync.GetBulkPublishDetailsInput{}
        output := &cognitosync.GetBulkPublishDetailsOutput{}

        mockClient.On("GetBulkPublishDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetBulkPublishDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCognitoEvents", func(t *testing.T) {
        input := &cognitosync.GetCognitoEventsInput{}
        output := &cognitosync.GetCognitoEventsOutput{}

        mockClient.On("GetCognitoEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetCognitoEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityPoolConfiguration", func(t *testing.T) {
        input := &cognitosync.GetIdentityPoolConfigurationInput{}
        output := &cognitosync.GetIdentityPoolConfigurationOutput{}

        mockClient.On("GetIdentityPoolConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityPoolConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &cognitosync.ListDatasetsInput{}
        output := &cognitosync.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityPoolUsage", func(t *testing.T) {
        input := &cognitosync.ListIdentityPoolUsageInput{}
        output := &cognitosync.ListIdentityPoolUsageOutput{}

        mockClient.On("ListIdentityPoolUsage", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityPoolUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecords", func(t *testing.T) {
        input := &cognitosync.ListRecordsInput{}
        output := &cognitosync.ListRecordsOutput{}

        mockClient.On("ListRecords", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDevice", func(t *testing.T) {
        input := &cognitosync.RegisterDeviceInput{}
        output := &cognitosync.RegisterDeviceOutput{}

        mockClient.On("RegisterDevice", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetCognitoEvents", func(t *testing.T) {
        input := &cognitosync.SetCognitoEventsInput{}
        output := &cognitosync.SetCognitoEventsOutput{}

        mockClient.On("SetCognitoEvents", ctx, input).Return(output, nil)

        result, err := mockClient.SetCognitoEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityPoolConfiguration", func(t *testing.T) {
        input := &cognitosync.SetIdentityPoolConfigurationInput{}
        output := &cognitosync.SetIdentityPoolConfigurationOutput{}

        mockClient.On("SetIdentityPoolConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityPoolConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubscribeToDataset", func(t *testing.T) {
        input := &cognitosync.SubscribeToDatasetInput{}
        output := &cognitosync.SubscribeToDatasetOutput{}

        mockClient.On("SubscribeToDataset", ctx, input).Return(output, nil)

        result, err := mockClient.SubscribeToDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnsubscribeFromDataset", func(t *testing.T) {
        input := &cognitosync.UnsubscribeFromDatasetInput{}
        output := &cognitosync.UnsubscribeFromDatasetOutput{}

        mockClient.On("UnsubscribeFromDataset", ctx, input).Return(output, nil)

        result, err := mockClient.UnsubscribeFromDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecords", func(t *testing.T) {
        input := &cognitosync.UpdateRecordsInput{}
        output := &cognitosync.UpdateRecordsOutput{}

        mockClient.On("UpdateRecords", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
