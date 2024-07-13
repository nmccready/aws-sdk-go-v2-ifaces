package pcaconnectorad_test

// tests for the pcaconnectorad service interface for this ../../../service/pcaconnectorad/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pcaconnectorad/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pcaconnectorad/pcaconnectorad_iface"
	"github.com/stretchr/testify/assert"
)

func TestPcaconnectoradServiceCanBeMocked(t *testing.T) {
	var iface pcaconnectorad_iface.IClient
	iface = &pcaconnectorad.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pcaconnectorad.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnector", func(t *testing.T) {
        input := &pcaconnectorad.CreateConnectorInput{}
        output := &pcaconnectorad.CreateConnectorOutput{}

        mockClient.On("CreateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectoryRegistration", func(t *testing.T) {
        input := &pcaconnectorad.CreateDirectoryRegistrationInput{}
        output := &pcaconnectorad.CreateDirectoryRegistrationOutput{}

        mockClient.On("CreateDirectoryRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectoryRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServicePrincipalName", func(t *testing.T) {
        input := &pcaconnectorad.CreateServicePrincipalNameInput{}
        output := &pcaconnectorad.CreateServicePrincipalNameOutput{}

        mockClient.On("CreateServicePrincipalName", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServicePrincipalName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplate", func(t *testing.T) {
        input := &pcaconnectorad.CreateTemplateInput{}
        output := &pcaconnectorad.CreateTemplateOutput{}

        mockClient.On("CreateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplateGroupAccessControlEntry", func(t *testing.T) {
        input := &pcaconnectorad.CreateTemplateGroupAccessControlEntryInput{}
        output := &pcaconnectorad.CreateTemplateGroupAccessControlEntryOutput{}

        mockClient.On("CreateTemplateGroupAccessControlEntry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplateGroupAccessControlEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnector", func(t *testing.T) {
        input := &pcaconnectorad.DeleteConnectorInput{}
        output := &pcaconnectorad.DeleteConnectorOutput{}

        mockClient.On("DeleteConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectoryRegistration", func(t *testing.T) {
        input := &pcaconnectorad.DeleteDirectoryRegistrationInput{}
        output := &pcaconnectorad.DeleteDirectoryRegistrationOutput{}

        mockClient.On("DeleteDirectoryRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectoryRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServicePrincipalName", func(t *testing.T) {
        input := &pcaconnectorad.DeleteServicePrincipalNameInput{}
        output := &pcaconnectorad.DeleteServicePrincipalNameOutput{}

        mockClient.On("DeleteServicePrincipalName", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServicePrincipalName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplate", func(t *testing.T) {
        input := &pcaconnectorad.DeleteTemplateInput{}
        output := &pcaconnectorad.DeleteTemplateOutput{}

        mockClient.On("DeleteTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplateGroupAccessControlEntry", func(t *testing.T) {
        input := &pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput{}
        output := &pcaconnectorad.DeleteTemplateGroupAccessControlEntryOutput{}

        mockClient.On("DeleteTemplateGroupAccessControlEntry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplateGroupAccessControlEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnector", func(t *testing.T) {
        input := &pcaconnectorad.GetConnectorInput{}
        output := &pcaconnectorad.GetConnectorOutput{}

        mockClient.On("GetConnector", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDirectoryRegistration", func(t *testing.T) {
        input := &pcaconnectorad.GetDirectoryRegistrationInput{}
        output := &pcaconnectorad.GetDirectoryRegistrationOutput{}

        mockClient.On("GetDirectoryRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDirectoryRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServicePrincipalName", func(t *testing.T) {
        input := &pcaconnectorad.GetServicePrincipalNameInput{}
        output := &pcaconnectorad.GetServicePrincipalNameOutput{}

        mockClient.On("GetServicePrincipalName", ctx, input).Return(output, nil)

        result, err := mockClient.GetServicePrincipalName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplate", func(t *testing.T) {
        input := &pcaconnectorad.GetTemplateInput{}
        output := &pcaconnectorad.GetTemplateOutput{}

        mockClient.On("GetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateGroupAccessControlEntry", func(t *testing.T) {
        input := &pcaconnectorad.GetTemplateGroupAccessControlEntryInput{}
        output := &pcaconnectorad.GetTemplateGroupAccessControlEntryOutput{}

        mockClient.On("GetTemplateGroupAccessControlEntry", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateGroupAccessControlEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &pcaconnectorad.ListConnectorsInput{}
        output := &pcaconnectorad.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDirectoryRegistrations", func(t *testing.T) {
        input := &pcaconnectorad.ListDirectoryRegistrationsInput{}
        output := &pcaconnectorad.ListDirectoryRegistrationsOutput{}

        mockClient.On("ListDirectoryRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDirectoryRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServicePrincipalNames", func(t *testing.T) {
        input := &pcaconnectorad.ListServicePrincipalNamesInput{}
        output := &pcaconnectorad.ListServicePrincipalNamesOutput{}

        mockClient.On("ListServicePrincipalNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListServicePrincipalNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pcaconnectorad.ListTagsForResourceInput{}
        output := &pcaconnectorad.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateGroupAccessControlEntries", func(t *testing.T) {
        input := &pcaconnectorad.ListTemplateGroupAccessControlEntriesInput{}
        output := &pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput{}

        mockClient.On("ListTemplateGroupAccessControlEntries", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateGroupAccessControlEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &pcaconnectorad.ListTemplatesInput{}
        output := &pcaconnectorad.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pcaconnectorad.TagResourceInput{}
        output := &pcaconnectorad.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pcaconnectorad.UntagResourceInput{}
        output := &pcaconnectorad.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplate", func(t *testing.T) {
        input := &pcaconnectorad.UpdateTemplateInput{}
        output := &pcaconnectorad.UpdateTemplateOutput{}

        mockClient.On("UpdateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplateGroupAccessControlEntry", func(t *testing.T) {
        input := &pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput{}
        output := &pcaconnectorad.UpdateTemplateGroupAccessControlEntryOutput{}

        mockClient.On("UpdateTemplateGroupAccessControlEntry", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplateGroupAccessControlEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
