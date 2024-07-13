package cloudcontrol_test

// tests for the cloudcontrol service interface for this ../../../service/cloudcontrol/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudcontrol/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudcontrol/cloudcontrol_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudcontrolServiceCanBeMocked(t *testing.T) {
	var iface cloudcontrol_iface.IClient
	iface = &cloudcontrol.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudcontrol.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelResourceRequest", func(t *testing.T) {
        input := &cloudcontrol.CancelResourceRequestInput{}
        output := &cloudcontrol.CancelResourceRequestOutput{}

        mockClient.On("CancelResourceRequest", ctx, input).Return(output, nil)

        result, err := mockClient.CancelResourceRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResource", func(t *testing.T) {
        input := &cloudcontrol.CreateResourceInput{}
        output := &cloudcontrol.CreateResourceOutput{}

        mockClient.On("CreateResource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResource", func(t *testing.T) {
        input := &cloudcontrol.DeleteResourceInput{}
        output := &cloudcontrol.DeleteResourceOutput{}

        mockClient.On("DeleteResource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResource", func(t *testing.T) {
        input := &cloudcontrol.GetResourceInput{}
        output := &cloudcontrol.GetResourceOutput{}

        mockClient.On("GetResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceRequestStatus", func(t *testing.T) {
        input := &cloudcontrol.GetResourceRequestStatusInput{}
        output := &cloudcontrol.GetResourceRequestStatusOutput{}

        mockClient.On("GetResourceRequestStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceRequestStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceRequests", func(t *testing.T) {
        input := &cloudcontrol.ListResourceRequestsInput{}
        output := &cloudcontrol.ListResourceRequestsOutput{}

        mockClient.On("ListResourceRequests", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResources", func(t *testing.T) {
        input := &cloudcontrol.ListResourcesInput{}
        output := &cloudcontrol.ListResourcesOutput{}

        mockClient.On("ListResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResource", func(t *testing.T) {
        input := &cloudcontrol.UpdateResourceInput{}
        output := &cloudcontrol.UpdateResourceOutput{}

        mockClient.On("UpdateResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
