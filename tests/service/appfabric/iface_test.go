package appfabric_test

// tests for the appfabric service interface for this ../../../service/appfabric/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appfabric"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appfabric/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appfabric/appfabric_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppfabricServiceCanBeMocked(t *testing.T) {
	var iface appfabric_iface.IClient
	iface = &appfabric.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appfabric.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetUserAccessTasks", func(t *testing.T) {
        input := &appfabric.BatchGetUserAccessTasksInput{}
        output := &appfabric.BatchGetUserAccessTasksOutput{}

        mockClient.On("BatchGetUserAccessTasks", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetUserAccessTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConnectAppAuthorization", func(t *testing.T) {
        input := &appfabric.ConnectAppAuthorizationInput{}
        output := &appfabric.ConnectAppAuthorizationOutput{}

        mockClient.On("ConnectAppAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.ConnectAppAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppAuthorization", func(t *testing.T) {
        input := &appfabric.CreateAppAuthorizationInput{}
        output := &appfabric.CreateAppAuthorizationOutput{}

        mockClient.On("CreateAppAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppBundle", func(t *testing.T) {
        input := &appfabric.CreateAppBundleInput{}
        output := &appfabric.CreateAppBundleOutput{}

        mockClient.On("CreateAppBundle", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIngestion", func(t *testing.T) {
        input := &appfabric.CreateIngestionInput{}
        output := &appfabric.CreateIngestionOutput{}

        mockClient.On("CreateIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIngestionDestination", func(t *testing.T) {
        input := &appfabric.CreateIngestionDestinationInput{}
        output := &appfabric.CreateIngestionDestinationOutput{}

        mockClient.On("CreateIngestionDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIngestionDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppAuthorization", func(t *testing.T) {
        input := &appfabric.DeleteAppAuthorizationInput{}
        output := &appfabric.DeleteAppAuthorizationOutput{}

        mockClient.On("DeleteAppAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppBundle", func(t *testing.T) {
        input := &appfabric.DeleteAppBundleInput{}
        output := &appfabric.DeleteAppBundleOutput{}

        mockClient.On("DeleteAppBundle", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIngestion", func(t *testing.T) {
        input := &appfabric.DeleteIngestionInput{}
        output := &appfabric.DeleteIngestionOutput{}

        mockClient.On("DeleteIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIngestionDestination", func(t *testing.T) {
        input := &appfabric.DeleteIngestionDestinationInput{}
        output := &appfabric.DeleteIngestionDestinationOutput{}

        mockClient.On("DeleteIngestionDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIngestionDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppAuthorization", func(t *testing.T) {
        input := &appfabric.GetAppAuthorizationInput{}
        output := &appfabric.GetAppAuthorizationOutput{}

        mockClient.On("GetAppAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppBundle", func(t *testing.T) {
        input := &appfabric.GetAppBundleInput{}
        output := &appfabric.GetAppBundleOutput{}

        mockClient.On("GetAppBundle", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIngestion", func(t *testing.T) {
        input := &appfabric.GetIngestionInput{}
        output := &appfabric.GetIngestionOutput{}

        mockClient.On("GetIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.GetIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIngestionDestination", func(t *testing.T) {
        input := &appfabric.GetIngestionDestinationInput{}
        output := &appfabric.GetIngestionDestinationOutput{}

        mockClient.On("GetIngestionDestination", ctx, input).Return(output, nil)

        result, err := mockClient.GetIngestionDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppAuthorizations", func(t *testing.T) {
        input := &appfabric.ListAppAuthorizationsInput{}
        output := &appfabric.ListAppAuthorizationsOutput{}

        mockClient.On("ListAppAuthorizations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppAuthorizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppBundles", func(t *testing.T) {
        input := &appfabric.ListAppBundlesInput{}
        output := &appfabric.ListAppBundlesOutput{}

        mockClient.On("ListAppBundles", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIngestionDestinations", func(t *testing.T) {
        input := &appfabric.ListIngestionDestinationsInput{}
        output := &appfabric.ListIngestionDestinationsOutput{}

        mockClient.On("ListIngestionDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListIngestionDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIngestions", func(t *testing.T) {
        input := &appfabric.ListIngestionsInput{}
        output := &appfabric.ListIngestionsOutput{}

        mockClient.On("ListIngestions", ctx, input).Return(output, nil)

        result, err := mockClient.ListIngestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appfabric.ListTagsForResourceInput{}
        output := &appfabric.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartIngestion", func(t *testing.T) {
        input := &appfabric.StartIngestionInput{}
        output := &appfabric.StartIngestionOutput{}

        mockClient.On("StartIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.StartIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartUserAccessTasks", func(t *testing.T) {
        input := &appfabric.StartUserAccessTasksInput{}
        output := &appfabric.StartUserAccessTasksOutput{}

        mockClient.On("StartUserAccessTasks", ctx, input).Return(output, nil)

        result, err := mockClient.StartUserAccessTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopIngestion", func(t *testing.T) {
        input := &appfabric.StopIngestionInput{}
        output := &appfabric.StopIngestionOutput{}

        mockClient.On("StopIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.StopIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appfabric.TagResourceInput{}
        output := &appfabric.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appfabric.UntagResourceInput{}
        output := &appfabric.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppAuthorization", func(t *testing.T) {
        input := &appfabric.UpdateAppAuthorizationInput{}
        output := &appfabric.UpdateAppAuthorizationOutput{}

        mockClient.On("UpdateAppAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIngestionDestination", func(t *testing.T) {
        input := &appfabric.UpdateIngestionDestinationInput{}
        output := &appfabric.UpdateIngestionDestinationOutput{}

        mockClient.On("UpdateIngestionDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIngestionDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
