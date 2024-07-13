package servicecatalogappregistry_test

// tests for the servicecatalogappregistry service interface for this ../../../service/servicecatalogappregistry/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicecatalogappregistry/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicecatalogappregistry/servicecatalogappregistry_iface"
	"github.com/stretchr/testify/assert"
)

func TestServicecatalogappregistryServiceCanBeMocked(t *testing.T) {
	var iface servicecatalogappregistry_iface.IClient
	iface = &servicecatalogappregistry.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := servicecatalogappregistry.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.AssociateAttributeGroupInput{}
        output := &servicecatalogappregistry.AssociateAttributeGroupOutput{}

        mockClient.On("AssociateAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResource", func(t *testing.T) {
        input := &servicecatalogappregistry.AssociateResourceInput{}
        output := &servicecatalogappregistry.AssociateResourceOutput{}

        mockClient.On("AssociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &servicecatalogappregistry.CreateApplicationInput{}
        output := &servicecatalogappregistry.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.CreateAttributeGroupInput{}
        output := &servicecatalogappregistry.CreateAttributeGroupOutput{}

        mockClient.On("CreateAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &servicecatalogappregistry.DeleteApplicationInput{}
        output := &servicecatalogappregistry.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.DeleteAttributeGroupInput{}
        output := &servicecatalogappregistry.DeleteAttributeGroupOutput{}

        mockClient.On("DeleteAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.DisassociateAttributeGroupInput{}
        output := &servicecatalogappregistry.DisassociateAttributeGroupOutput{}

        mockClient.On("DisassociateAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResource", func(t *testing.T) {
        input := &servicecatalogappregistry.DisassociateResourceInput{}
        output := &servicecatalogappregistry.DisassociateResourceOutput{}

        mockClient.On("DisassociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &servicecatalogappregistry.GetApplicationInput{}
        output := &servicecatalogappregistry.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociatedResource", func(t *testing.T) {
        input := &servicecatalogappregistry.GetAssociatedResourceInput{}
        output := &servicecatalogappregistry.GetAssociatedResourceOutput{}

        mockClient.On("GetAssociatedResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociatedResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.GetAttributeGroupInput{}
        output := &servicecatalogappregistry.GetAttributeGroupOutput{}

        mockClient.On("GetAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguration", func(t *testing.T) {
        input := &servicecatalogappregistry.GetConfigurationInput{}
        output := &servicecatalogappregistry.GetConfigurationOutput{}

        mockClient.On("GetConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &servicecatalogappregistry.ListApplicationsInput{}
        output := &servicecatalogappregistry.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedAttributeGroups", func(t *testing.T) {
        input := &servicecatalogappregistry.ListAssociatedAttributeGroupsInput{}
        output := &servicecatalogappregistry.ListAssociatedAttributeGroupsOutput{}

        mockClient.On("ListAssociatedAttributeGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedAttributeGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedResources", func(t *testing.T) {
        input := &servicecatalogappregistry.ListAssociatedResourcesInput{}
        output := &servicecatalogappregistry.ListAssociatedResourcesOutput{}

        mockClient.On("ListAssociatedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttributeGroups", func(t *testing.T) {
        input := &servicecatalogappregistry.ListAttributeGroupsInput{}
        output := &servicecatalogappregistry.ListAttributeGroupsOutput{}

        mockClient.On("ListAttributeGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttributeGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttributeGroupsForApplication", func(t *testing.T) {
        input := &servicecatalogappregistry.ListAttributeGroupsForApplicationInput{}
        output := &servicecatalogappregistry.ListAttributeGroupsForApplicationOutput{}

        mockClient.On("ListAttributeGroupsForApplication", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttributeGroupsForApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &servicecatalogappregistry.ListTagsForResourceInput{}
        output := &servicecatalogappregistry.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfiguration", func(t *testing.T) {
        input := &servicecatalogappregistry.PutConfigurationInput{}
        output := &servicecatalogappregistry.PutConfigurationOutput{}

        mockClient.On("PutConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSyncResource", func(t *testing.T) {
        input := &servicecatalogappregistry.SyncResourceInput{}
        output := &servicecatalogappregistry.SyncResourceOutput{}

        mockClient.On("SyncResource", ctx, input).Return(output, nil)

        result, err := mockClient.SyncResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &servicecatalogappregistry.TagResourceInput{}
        output := &servicecatalogappregistry.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &servicecatalogappregistry.UntagResourceInput{}
        output := &servicecatalogappregistry.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &servicecatalogappregistry.UpdateApplicationInput{}
        output := &servicecatalogappregistry.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAttributeGroup", func(t *testing.T) {
        input := &servicecatalogappregistry.UpdateAttributeGroupInput{}
        output := &servicecatalogappregistry.UpdateAttributeGroupOutput{}

        mockClient.On("UpdateAttributeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAttributeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
