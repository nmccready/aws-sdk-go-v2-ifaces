package ssmsap_test

// tests for the ssmsap service interface for this ../../../service/ssmsap/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssmsap"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmsap/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmsap/ssmsap_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsmsapServiceCanBeMocked(t *testing.T) {
	var iface ssmsap_iface.IClient
	iface = &ssmsap.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssmsap.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePermission", func(t *testing.T) {
        input := &ssmsap.DeleteResourcePermissionInput{}
        output := &ssmsap.DeleteResourcePermissionOutput{}

        mockClient.On("DeleteResourcePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterApplication", func(t *testing.T) {
        input := &ssmsap.DeregisterApplicationInput{}
        output := &ssmsap.DeregisterApplicationOutput{}

        mockClient.On("DeregisterApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &ssmsap.GetApplicationInput{}
        output := &ssmsap.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponent", func(t *testing.T) {
        input := &ssmsap.GetComponentInput{}
        output := &ssmsap.GetComponentOutput{}

        mockClient.On("GetComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatabase", func(t *testing.T) {
        input := &ssmsap.GetDatabaseInput{}
        output := &ssmsap.GetDatabaseOutput{}

        mockClient.On("GetDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperation", func(t *testing.T) {
        input := &ssmsap.GetOperationInput{}
        output := &ssmsap.GetOperationOutput{}

        mockClient.On("GetOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePermission", func(t *testing.T) {
        input := &ssmsap.GetResourcePermissionInput{}
        output := &ssmsap.GetResourcePermissionOutput{}

        mockClient.On("GetResourcePermission", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &ssmsap.ListApplicationsInput{}
        output := &ssmsap.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &ssmsap.ListComponentsInput{}
        output := &ssmsap.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatabases", func(t *testing.T) {
        input := &ssmsap.ListDatabasesInput{}
        output := &ssmsap.ListDatabasesOutput{}

        mockClient.On("ListDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOperationEvents", func(t *testing.T) {
        input := &ssmsap.ListOperationEventsInput{}
        output := &ssmsap.ListOperationEventsOutput{}

        mockClient.On("ListOperationEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListOperationEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOperations", func(t *testing.T) {
        input := &ssmsap.ListOperationsInput{}
        output := &ssmsap.ListOperationsOutput{}

        mockClient.On("ListOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ssmsap.ListTagsForResourceInput{}
        output := &ssmsap.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePermission", func(t *testing.T) {
        input := &ssmsap.PutResourcePermissionInput{}
        output := &ssmsap.PutResourcePermissionOutput{}

        mockClient.On("PutResourcePermission", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterApplication", func(t *testing.T) {
        input := &ssmsap.RegisterApplicationInput{}
        output := &ssmsap.RegisterApplicationOutput{}

        mockClient.On("RegisterApplication", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplication", func(t *testing.T) {
        input := &ssmsap.StartApplicationInput{}
        output := &ssmsap.StartApplicationOutput{}

        mockClient.On("StartApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplicationRefresh", func(t *testing.T) {
        input := &ssmsap.StartApplicationRefreshInput{}
        output := &ssmsap.StartApplicationRefreshOutput{}

        mockClient.On("StartApplicationRefresh", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplicationRefresh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApplication", func(t *testing.T) {
        input := &ssmsap.StopApplicationInput{}
        output := &ssmsap.StopApplicationOutput{}

        mockClient.On("StopApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ssmsap.TagResourceInput{}
        output := &ssmsap.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ssmsap.UntagResourceInput{}
        output := &ssmsap.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationSettings", func(t *testing.T) {
        input := &ssmsap.UpdateApplicationSettingsInput{}
        output := &ssmsap.UpdateApplicationSettingsOutput{}

        mockClient.On("UpdateApplicationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
