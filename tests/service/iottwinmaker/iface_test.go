package iottwinmaker_test

// tests for the iottwinmaker service interface for this ../../../service/iottwinmaker/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iottwinmaker/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iottwinmaker/iottwinmaker_iface"
	"github.com/stretchr/testify/assert"
)

func TestIottwinmakerServiceCanBeMocked(t *testing.T) {
	var iface iottwinmaker_iface.IClient
	iface = &iottwinmaker.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iottwinmaker.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutPropertyValues", func(t *testing.T) {
        input := &iottwinmaker.BatchPutPropertyValuesInput{}
        output := &iottwinmaker.BatchPutPropertyValuesOutput{}

        mockClient.On("BatchPutPropertyValues", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutPropertyValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMetadataTransferJob", func(t *testing.T) {
        input := &iottwinmaker.CancelMetadataTransferJobInput{}
        output := &iottwinmaker.CancelMetadataTransferJobOutput{}

        mockClient.On("CancelMetadataTransferJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMetadataTransferJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponentType", func(t *testing.T) {
        input := &iottwinmaker.CreateComponentTypeInput{}
        output := &iottwinmaker.CreateComponentTypeOutput{}

        mockClient.On("CreateComponentType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponentType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEntity", func(t *testing.T) {
        input := &iottwinmaker.CreateEntityInput{}
        output := &iottwinmaker.CreateEntityOutput{}

        mockClient.On("CreateEntity", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMetadataTransferJob", func(t *testing.T) {
        input := &iottwinmaker.CreateMetadataTransferJobInput{}
        output := &iottwinmaker.CreateMetadataTransferJobOutput{}

        mockClient.On("CreateMetadataTransferJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMetadataTransferJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScene", func(t *testing.T) {
        input := &iottwinmaker.CreateSceneInput{}
        output := &iottwinmaker.CreateSceneOutput{}

        mockClient.On("CreateScene", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScene(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSyncJob", func(t *testing.T) {
        input := &iottwinmaker.CreateSyncJobInput{}
        output := &iottwinmaker.CreateSyncJobOutput{}

        mockClient.On("CreateSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspace", func(t *testing.T) {
        input := &iottwinmaker.CreateWorkspaceInput{}
        output := &iottwinmaker.CreateWorkspaceOutput{}

        mockClient.On("CreateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponentType", func(t *testing.T) {
        input := &iottwinmaker.DeleteComponentTypeInput{}
        output := &iottwinmaker.DeleteComponentTypeOutput{}

        mockClient.On("DeleteComponentType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponentType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEntity", func(t *testing.T) {
        input := &iottwinmaker.DeleteEntityInput{}
        output := &iottwinmaker.DeleteEntityOutput{}

        mockClient.On("DeleteEntity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScene", func(t *testing.T) {
        input := &iottwinmaker.DeleteSceneInput{}
        output := &iottwinmaker.DeleteSceneOutput{}

        mockClient.On("DeleteScene", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScene(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSyncJob", func(t *testing.T) {
        input := &iottwinmaker.DeleteSyncJobInput{}
        output := &iottwinmaker.DeleteSyncJobOutput{}

        mockClient.On("DeleteSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspace", func(t *testing.T) {
        input := &iottwinmaker.DeleteWorkspaceInput{}
        output := &iottwinmaker.DeleteWorkspaceOutput{}

        mockClient.On("DeleteWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteQuery", func(t *testing.T) {
        input := &iottwinmaker.ExecuteQueryInput{}
        output := &iottwinmaker.ExecuteQueryOutput{}

        mockClient.On("ExecuteQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponentType", func(t *testing.T) {
        input := &iottwinmaker.GetComponentTypeInput{}
        output := &iottwinmaker.GetComponentTypeOutput{}

        mockClient.On("GetComponentType", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponentType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEntity", func(t *testing.T) {
        input := &iottwinmaker.GetEntityInput{}
        output := &iottwinmaker.GetEntityOutput{}

        mockClient.On("GetEntity", ctx, input).Return(output, nil)

        result, err := mockClient.GetEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetadataTransferJob", func(t *testing.T) {
        input := &iottwinmaker.GetMetadataTransferJobInput{}
        output := &iottwinmaker.GetMetadataTransferJobOutput{}

        mockClient.On("GetMetadataTransferJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetadataTransferJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPricingPlan", func(t *testing.T) {
        input := &iottwinmaker.GetPricingPlanInput{}
        output := &iottwinmaker.GetPricingPlanOutput{}

        mockClient.On("GetPricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetPricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPropertyValue", func(t *testing.T) {
        input := &iottwinmaker.GetPropertyValueInput{}
        output := &iottwinmaker.GetPropertyValueOutput{}

        mockClient.On("GetPropertyValue", ctx, input).Return(output, nil)

        result, err := mockClient.GetPropertyValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPropertyValueHistory", func(t *testing.T) {
        input := &iottwinmaker.GetPropertyValueHistoryInput{}
        output := &iottwinmaker.GetPropertyValueHistoryOutput{}

        mockClient.On("GetPropertyValueHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetPropertyValueHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScene", func(t *testing.T) {
        input := &iottwinmaker.GetSceneInput{}
        output := &iottwinmaker.GetSceneOutput{}

        mockClient.On("GetScene", ctx, input).Return(output, nil)

        result, err := mockClient.GetScene(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSyncJob", func(t *testing.T) {
        input := &iottwinmaker.GetSyncJobInput{}
        output := &iottwinmaker.GetSyncJobOutput{}

        mockClient.On("GetSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkspace", func(t *testing.T) {
        input := &iottwinmaker.GetWorkspaceInput{}
        output := &iottwinmaker.GetWorkspaceOutput{}

        mockClient.On("GetWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponentTypes", func(t *testing.T) {
        input := &iottwinmaker.ListComponentTypesInput{}
        output := &iottwinmaker.ListComponentTypesOutput{}

        mockClient.On("ListComponentTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponentTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &iottwinmaker.ListComponentsInput{}
        output := &iottwinmaker.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntities", func(t *testing.T) {
        input := &iottwinmaker.ListEntitiesInput{}
        output := &iottwinmaker.ListEntitiesOutput{}

        mockClient.On("ListEntities", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetadataTransferJobs", func(t *testing.T) {
        input := &iottwinmaker.ListMetadataTransferJobsInput{}
        output := &iottwinmaker.ListMetadataTransferJobsOutput{}

        mockClient.On("ListMetadataTransferJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetadataTransferJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProperties", func(t *testing.T) {
        input := &iottwinmaker.ListPropertiesInput{}
        output := &iottwinmaker.ListPropertiesOutput{}

        mockClient.On("ListProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ListProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScenes", func(t *testing.T) {
        input := &iottwinmaker.ListScenesInput{}
        output := &iottwinmaker.ListScenesOutput{}

        mockClient.On("ListScenes", ctx, input).Return(output, nil)

        result, err := mockClient.ListScenes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSyncJobs", func(t *testing.T) {
        input := &iottwinmaker.ListSyncJobsInput{}
        output := &iottwinmaker.ListSyncJobsOutput{}

        mockClient.On("ListSyncJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSyncJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSyncResources", func(t *testing.T) {
        input := &iottwinmaker.ListSyncResourcesInput{}
        output := &iottwinmaker.ListSyncResourcesOutput{}

        mockClient.On("ListSyncResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListSyncResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iottwinmaker.ListTagsForResourceInput{}
        output := &iottwinmaker.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkspaces", func(t *testing.T) {
        input := &iottwinmaker.ListWorkspacesInput{}
        output := &iottwinmaker.ListWorkspacesOutput{}

        mockClient.On("ListWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iottwinmaker.TagResourceInput{}
        output := &iottwinmaker.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iottwinmaker.UntagResourceInput{}
        output := &iottwinmaker.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComponentType", func(t *testing.T) {
        input := &iottwinmaker.UpdateComponentTypeInput{}
        output := &iottwinmaker.UpdateComponentTypeOutput{}

        mockClient.On("UpdateComponentType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComponentType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEntity", func(t *testing.T) {
        input := &iottwinmaker.UpdateEntityInput{}
        output := &iottwinmaker.UpdateEntityOutput{}

        mockClient.On("UpdateEntity", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePricingPlan", func(t *testing.T) {
        input := &iottwinmaker.UpdatePricingPlanInput{}
        output := &iottwinmaker.UpdatePricingPlanOutput{}

        mockClient.On("UpdatePricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScene", func(t *testing.T) {
        input := &iottwinmaker.UpdateSceneInput{}
        output := &iottwinmaker.UpdateSceneOutput{}

        mockClient.On("UpdateScene", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScene(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspace", func(t *testing.T) {
        input := &iottwinmaker.UpdateWorkspaceInput{}
        output := &iottwinmaker.UpdateWorkspaceOutput{}

        mockClient.On("UpdateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
