package workspacesthinclient_test

// tests for the workspacesthinclient service interface for this ../../../service/workspacesthinclient/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspacesthinclient"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspacesthinclient/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workspacesthinclient/workspacesthinclient_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkspacesthinclientServiceCanBeMocked(t *testing.T) {
	var iface workspacesthinclient_iface.IClient
	iface = &workspacesthinclient.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workspacesthinclient.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &workspacesthinclient.CreateEnvironmentInput{}
        output := &workspacesthinclient.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevice", func(t *testing.T) {
        input := &workspacesthinclient.DeleteDeviceInput{}
        output := &workspacesthinclient.DeleteDeviceOutput{}

        mockClient.On("DeleteDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &workspacesthinclient.DeleteEnvironmentInput{}
        output := &workspacesthinclient.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterDevice", func(t *testing.T) {
        input := &workspacesthinclient.DeregisterDeviceInput{}
        output := &workspacesthinclient.DeregisterDeviceOutput{}

        mockClient.On("DeregisterDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevice", func(t *testing.T) {
        input := &workspacesthinclient.GetDeviceInput{}
        output := &workspacesthinclient.GetDeviceOutput{}

        mockClient.On("GetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &workspacesthinclient.GetEnvironmentInput{}
        output := &workspacesthinclient.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSoftwareSet", func(t *testing.T) {
        input := &workspacesthinclient.GetSoftwareSetInput{}
        output := &workspacesthinclient.GetSoftwareSetOutput{}

        mockClient.On("GetSoftwareSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetSoftwareSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &workspacesthinclient.ListDevicesInput{}
        output := &workspacesthinclient.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &workspacesthinclient.ListEnvironmentsInput{}
        output := &workspacesthinclient.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSoftwareSets", func(t *testing.T) {
        input := &workspacesthinclient.ListSoftwareSetsInput{}
        output := &workspacesthinclient.ListSoftwareSetsOutput{}

        mockClient.On("ListSoftwareSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSoftwareSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &workspacesthinclient.ListTagsForResourceInput{}
        output := &workspacesthinclient.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &workspacesthinclient.TagResourceInput{}
        output := &workspacesthinclient.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &workspacesthinclient.UntagResourceInput{}
        output := &workspacesthinclient.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevice", func(t *testing.T) {
        input := &workspacesthinclient.UpdateDeviceInput{}
        output := &workspacesthinclient.UpdateDeviceOutput{}

        mockClient.On("UpdateDevice", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &workspacesthinclient.UpdateEnvironmentInput{}
        output := &workspacesthinclient.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSoftwareSet", func(t *testing.T) {
        input := &workspacesthinclient.UpdateSoftwareSetInput{}
        output := &workspacesthinclient.UpdateSoftwareSetOutput{}

        mockClient.On("UpdateSoftwareSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSoftwareSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
