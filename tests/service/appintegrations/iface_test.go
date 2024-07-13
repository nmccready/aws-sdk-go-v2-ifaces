package appintegrations_test

// tests for the appintegrations service interface for this ../../../service/appintegrations/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appintegrations"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appintegrations/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appintegrations/appintegrations_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppintegrationsServiceCanBeMocked(t *testing.T) {
	var iface appintegrations_iface.IClient
	iface = &appintegrations.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appintegrations.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &appintegrations.CreateApplicationInput{}
        output := &appintegrations.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataIntegration", func(t *testing.T) {
        input := &appintegrations.CreateDataIntegrationInput{}
        output := &appintegrations.CreateDataIntegrationOutput{}

        mockClient.On("CreateDataIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventIntegration", func(t *testing.T) {
        input := &appintegrations.CreateEventIntegrationInput{}
        output := &appintegrations.CreateEventIntegrationOutput{}

        mockClient.On("CreateEventIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &appintegrations.DeleteApplicationInput{}
        output := &appintegrations.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataIntegration", func(t *testing.T) {
        input := &appintegrations.DeleteDataIntegrationInput{}
        output := &appintegrations.DeleteDataIntegrationOutput{}

        mockClient.On("DeleteDataIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventIntegration", func(t *testing.T) {
        input := &appintegrations.DeleteEventIntegrationInput{}
        output := &appintegrations.DeleteEventIntegrationOutput{}

        mockClient.On("DeleteEventIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &appintegrations.GetApplicationInput{}
        output := &appintegrations.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataIntegration", func(t *testing.T) {
        input := &appintegrations.GetDataIntegrationInput{}
        output := &appintegrations.GetDataIntegrationOutput{}

        mockClient.On("GetDataIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventIntegration", func(t *testing.T) {
        input := &appintegrations.GetEventIntegrationInput{}
        output := &appintegrations.GetEventIntegrationOutput{}

        mockClient.On("GetEventIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationAssociations", func(t *testing.T) {
        input := &appintegrations.ListApplicationAssociationsInput{}
        output := &appintegrations.ListApplicationAssociationsOutput{}

        mockClient.On("ListApplicationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &appintegrations.ListApplicationsInput{}
        output := &appintegrations.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataIntegrationAssociations", func(t *testing.T) {
        input := &appintegrations.ListDataIntegrationAssociationsInput{}
        output := &appintegrations.ListDataIntegrationAssociationsOutput{}

        mockClient.On("ListDataIntegrationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataIntegrationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataIntegrations", func(t *testing.T) {
        input := &appintegrations.ListDataIntegrationsInput{}
        output := &appintegrations.ListDataIntegrationsOutput{}

        mockClient.On("ListDataIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventIntegrationAssociations", func(t *testing.T) {
        input := &appintegrations.ListEventIntegrationAssociationsInput{}
        output := &appintegrations.ListEventIntegrationAssociationsOutput{}

        mockClient.On("ListEventIntegrationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventIntegrationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventIntegrations", func(t *testing.T) {
        input := &appintegrations.ListEventIntegrationsInput{}
        output := &appintegrations.ListEventIntegrationsOutput{}

        mockClient.On("ListEventIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appintegrations.ListTagsForResourceInput{}
        output := &appintegrations.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appintegrations.TagResourceInput{}
        output := &appintegrations.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appintegrations.UntagResourceInput{}
        output := &appintegrations.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &appintegrations.UpdateApplicationInput{}
        output := &appintegrations.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataIntegration", func(t *testing.T) {
        input := &appintegrations.UpdateDataIntegrationInput{}
        output := &appintegrations.UpdateDataIntegrationOutput{}

        mockClient.On("UpdateDataIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventIntegration", func(t *testing.T) {
        input := &appintegrations.UpdateEventIntegrationInput{}
        output := &appintegrations.UpdateEventIntegrationOutput{}

        mockClient.On("UpdateEventIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
