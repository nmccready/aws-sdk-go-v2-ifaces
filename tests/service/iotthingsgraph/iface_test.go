package iotthingsgraph_test

// tests for the iotthingsgraph service interface for this ../../../service/iotthingsgraph/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotthingsgraph/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotthingsgraph/iotthingsgraph_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotthingsgraphServiceCanBeMocked(t *testing.T) {
	var iface iotthingsgraph_iface.IClient
	iface = &iotthingsgraph.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotthingsgraph.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEntityToThing", func(t *testing.T) {
        input := &iotthingsgraph.AssociateEntityToThingInput{}
        output := &iotthingsgraph.AssociateEntityToThingOutput{}

        mockClient.On("AssociateEntityToThing", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEntityToThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlowTemplate", func(t *testing.T) {
        input := &iotthingsgraph.CreateFlowTemplateInput{}
        output := &iotthingsgraph.CreateFlowTemplateOutput{}

        mockClient.On("CreateFlowTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlowTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSystemInstance", func(t *testing.T) {
        input := &iotthingsgraph.CreateSystemInstanceInput{}
        output := &iotthingsgraph.CreateSystemInstanceOutput{}

        mockClient.On("CreateSystemInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSystemInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSystemTemplate", func(t *testing.T) {
        input := &iotthingsgraph.CreateSystemTemplateInput{}
        output := &iotthingsgraph.CreateSystemTemplateOutput{}

        mockClient.On("CreateSystemTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSystemTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlowTemplate", func(t *testing.T) {
        input := &iotthingsgraph.DeleteFlowTemplateInput{}
        output := &iotthingsgraph.DeleteFlowTemplateOutput{}

        mockClient.On("DeleteFlowTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlowTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNamespace", func(t *testing.T) {
        input := &iotthingsgraph.DeleteNamespaceInput{}
        output := &iotthingsgraph.DeleteNamespaceOutput{}

        mockClient.On("DeleteNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSystemInstance", func(t *testing.T) {
        input := &iotthingsgraph.DeleteSystemInstanceInput{}
        output := &iotthingsgraph.DeleteSystemInstanceOutput{}

        mockClient.On("DeleteSystemInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSystemInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSystemTemplate", func(t *testing.T) {
        input := &iotthingsgraph.DeleteSystemTemplateInput{}
        output := &iotthingsgraph.DeleteSystemTemplateOutput{}

        mockClient.On("DeleteSystemTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSystemTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeploySystemInstance", func(t *testing.T) {
        input := &iotthingsgraph.DeploySystemInstanceInput{}
        output := &iotthingsgraph.DeploySystemInstanceOutput{}

        mockClient.On("DeploySystemInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeploySystemInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateFlowTemplate", func(t *testing.T) {
        input := &iotthingsgraph.DeprecateFlowTemplateInput{}
        output := &iotthingsgraph.DeprecateFlowTemplateOutput{}

        mockClient.On("DeprecateFlowTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateFlowTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateSystemTemplate", func(t *testing.T) {
        input := &iotthingsgraph.DeprecateSystemTemplateInput{}
        output := &iotthingsgraph.DeprecateSystemTemplateOutput{}

        mockClient.On("DeprecateSystemTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateSystemTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNamespace", func(t *testing.T) {
        input := &iotthingsgraph.DescribeNamespaceInput{}
        output := &iotthingsgraph.DescribeNamespaceOutput{}

        mockClient.On("DescribeNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDissociateEntityFromThing", func(t *testing.T) {
        input := &iotthingsgraph.DissociateEntityFromThingInput{}
        output := &iotthingsgraph.DissociateEntityFromThingOutput{}

        mockClient.On("DissociateEntityFromThing", ctx, input).Return(output, nil)

        result, err := mockClient.DissociateEntityFromThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEntities", func(t *testing.T) {
        input := &iotthingsgraph.GetEntitiesInput{}
        output := &iotthingsgraph.GetEntitiesOutput{}

        mockClient.On("GetEntities", ctx, input).Return(output, nil)

        result, err := mockClient.GetEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowTemplate", func(t *testing.T) {
        input := &iotthingsgraph.GetFlowTemplateInput{}
        output := &iotthingsgraph.GetFlowTemplateOutput{}

        mockClient.On("GetFlowTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowTemplateRevisions", func(t *testing.T) {
        input := &iotthingsgraph.GetFlowTemplateRevisionsInput{}
        output := &iotthingsgraph.GetFlowTemplateRevisionsOutput{}

        mockClient.On("GetFlowTemplateRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowTemplateRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNamespaceDeletionStatus", func(t *testing.T) {
        input := &iotthingsgraph.GetNamespaceDeletionStatusInput{}
        output := &iotthingsgraph.GetNamespaceDeletionStatusOutput{}

        mockClient.On("GetNamespaceDeletionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetNamespaceDeletionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSystemInstance", func(t *testing.T) {
        input := &iotthingsgraph.GetSystemInstanceInput{}
        output := &iotthingsgraph.GetSystemInstanceOutput{}

        mockClient.On("GetSystemInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetSystemInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSystemTemplate", func(t *testing.T) {
        input := &iotthingsgraph.GetSystemTemplateInput{}
        output := &iotthingsgraph.GetSystemTemplateOutput{}

        mockClient.On("GetSystemTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetSystemTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSystemTemplateRevisions", func(t *testing.T) {
        input := &iotthingsgraph.GetSystemTemplateRevisionsInput{}
        output := &iotthingsgraph.GetSystemTemplateRevisionsOutput{}

        mockClient.On("GetSystemTemplateRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.GetSystemTemplateRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUploadStatus", func(t *testing.T) {
        input := &iotthingsgraph.GetUploadStatusInput{}
        output := &iotthingsgraph.GetUploadStatusOutput{}

        mockClient.On("GetUploadStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetUploadStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowExecutionMessages", func(t *testing.T) {
        input := &iotthingsgraph.ListFlowExecutionMessagesInput{}
        output := &iotthingsgraph.ListFlowExecutionMessagesOutput{}

        mockClient.On("ListFlowExecutionMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowExecutionMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotthingsgraph.ListTagsForResourceInput{}
        output := &iotthingsgraph.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchEntities", func(t *testing.T) {
        input := &iotthingsgraph.SearchEntitiesInput{}
        output := &iotthingsgraph.SearchEntitiesOutput{}

        mockClient.On("SearchEntities", ctx, input).Return(output, nil)

        result, err := mockClient.SearchEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchFlowExecutions", func(t *testing.T) {
        input := &iotthingsgraph.SearchFlowExecutionsInput{}
        output := &iotthingsgraph.SearchFlowExecutionsOutput{}

        mockClient.On("SearchFlowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.SearchFlowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchFlowTemplates", func(t *testing.T) {
        input := &iotthingsgraph.SearchFlowTemplatesInput{}
        output := &iotthingsgraph.SearchFlowTemplatesOutput{}

        mockClient.On("SearchFlowTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.SearchFlowTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSystemInstances", func(t *testing.T) {
        input := &iotthingsgraph.SearchSystemInstancesInput{}
        output := &iotthingsgraph.SearchSystemInstancesOutput{}

        mockClient.On("SearchSystemInstances", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSystemInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSystemTemplates", func(t *testing.T) {
        input := &iotthingsgraph.SearchSystemTemplatesInput{}
        output := &iotthingsgraph.SearchSystemTemplatesOutput{}

        mockClient.On("SearchSystemTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSystemTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchThings", func(t *testing.T) {
        input := &iotthingsgraph.SearchThingsInput{}
        output := &iotthingsgraph.SearchThingsOutput{}

        mockClient.On("SearchThings", ctx, input).Return(output, nil)

        result, err := mockClient.SearchThings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotthingsgraph.TagResourceInput{}
        output := &iotthingsgraph.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUndeploySystemInstance", func(t *testing.T) {
        input := &iotthingsgraph.UndeploySystemInstanceInput{}
        output := &iotthingsgraph.UndeploySystemInstanceOutput{}

        mockClient.On("UndeploySystemInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UndeploySystemInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotthingsgraph.UntagResourceInput{}
        output := &iotthingsgraph.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowTemplate", func(t *testing.T) {
        input := &iotthingsgraph.UpdateFlowTemplateInput{}
        output := &iotthingsgraph.UpdateFlowTemplateOutput{}

        mockClient.On("UpdateFlowTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSystemTemplate", func(t *testing.T) {
        input := &iotthingsgraph.UpdateSystemTemplateInput{}
        output := &iotthingsgraph.UpdateSystemTemplateOutput{}

        mockClient.On("UpdateSystemTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSystemTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadEntityDefinitions", func(t *testing.T) {
        input := &iotthingsgraph.UploadEntityDefinitionsInput{}
        output := &iotthingsgraph.UploadEntityDefinitionsOutput{}

        mockClient.On("UploadEntityDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.UploadEntityDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
