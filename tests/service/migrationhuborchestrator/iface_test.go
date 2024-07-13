package migrationhuborchestrator_test

// tests for the migrationhuborchestrator service interface for this ../../../service/migrationhuborchestrator/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhuborchestrator/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhuborchestrator/migrationhuborchestrator_iface"
	"github.com/stretchr/testify/assert"
)

func TestMigrationhuborchestratorServiceCanBeMocked(t *testing.T) {
	var iface migrationhuborchestrator_iface.IClient
	iface = &migrationhuborchestrator.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := migrationhuborchestrator.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplate", func(t *testing.T) {
        input := &migrationhuborchestrator.CreateTemplateInput{}
        output := &migrationhuborchestrator.CreateTemplateOutput{}

        mockClient.On("CreateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.CreateWorkflowInput{}
        output := &migrationhuborchestrator.CreateWorkflowOutput{}

        mockClient.On("CreateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflowStep", func(t *testing.T) {
        input := &migrationhuborchestrator.CreateWorkflowStepInput{}
        output := &migrationhuborchestrator.CreateWorkflowStepOutput{}

        mockClient.On("CreateWorkflowStep", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflowStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflowStepGroup", func(t *testing.T) {
        input := &migrationhuborchestrator.CreateWorkflowStepGroupInput{}
        output := &migrationhuborchestrator.CreateWorkflowStepGroupOutput{}

        mockClient.On("CreateWorkflowStepGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflowStepGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplate", func(t *testing.T) {
        input := &migrationhuborchestrator.DeleteTemplateInput{}
        output := &migrationhuborchestrator.DeleteTemplateOutput{}

        mockClient.On("DeleteTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.DeleteWorkflowInput{}
        output := &migrationhuborchestrator.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflowStep", func(t *testing.T) {
        input := &migrationhuborchestrator.DeleteWorkflowStepInput{}
        output := &migrationhuborchestrator.DeleteWorkflowStepOutput{}

        mockClient.On("DeleteWorkflowStep", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflowStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflowStepGroup", func(t *testing.T) {
        input := &migrationhuborchestrator.DeleteWorkflowStepGroupInput{}
        output := &migrationhuborchestrator.DeleteWorkflowStepGroupOutput{}

        mockClient.On("DeleteWorkflowStepGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflowStepGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplate", func(t *testing.T) {
        input := &migrationhuborchestrator.GetTemplateInput{}
        output := &migrationhuborchestrator.GetTemplateOutput{}

        mockClient.On("GetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateStep", func(t *testing.T) {
        input := &migrationhuborchestrator.GetTemplateStepInput{}
        output := &migrationhuborchestrator.GetTemplateStepOutput{}

        mockClient.On("GetTemplateStep", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateStepGroup", func(t *testing.T) {
        input := &migrationhuborchestrator.GetTemplateStepGroupInput{}
        output := &migrationhuborchestrator.GetTemplateStepGroupOutput{}

        mockClient.On("GetTemplateStepGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateStepGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.GetWorkflowInput{}
        output := &migrationhuborchestrator.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowStep", func(t *testing.T) {
        input := &migrationhuborchestrator.GetWorkflowStepInput{}
        output := &migrationhuborchestrator.GetWorkflowStepOutput{}

        mockClient.On("GetWorkflowStep", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowStepGroup", func(t *testing.T) {
        input := &migrationhuborchestrator.GetWorkflowStepGroupInput{}
        output := &migrationhuborchestrator.GetWorkflowStepGroupOutput{}

        mockClient.On("GetWorkflowStepGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowStepGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlugins", func(t *testing.T) {
        input := &migrationhuborchestrator.ListPluginsInput{}
        output := &migrationhuborchestrator.ListPluginsOutput{}

        mockClient.On("ListPlugins", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlugins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &migrationhuborchestrator.ListTagsForResourceInput{}
        output := &migrationhuborchestrator.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateStepGroups", func(t *testing.T) {
        input := &migrationhuborchestrator.ListTemplateStepGroupsInput{}
        output := &migrationhuborchestrator.ListTemplateStepGroupsOutput{}

        mockClient.On("ListTemplateStepGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateStepGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateSteps", func(t *testing.T) {
        input := &migrationhuborchestrator.ListTemplateStepsInput{}
        output := &migrationhuborchestrator.ListTemplateStepsOutput{}

        mockClient.On("ListTemplateSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &migrationhuborchestrator.ListTemplatesInput{}
        output := &migrationhuborchestrator.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowStepGroups", func(t *testing.T) {
        input := &migrationhuborchestrator.ListWorkflowStepGroupsInput{}
        output := &migrationhuborchestrator.ListWorkflowStepGroupsOutput{}

        mockClient.On("ListWorkflowStepGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowStepGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowSteps", func(t *testing.T) {
        input := &migrationhuborchestrator.ListWorkflowStepsInput{}
        output := &migrationhuborchestrator.ListWorkflowStepsOutput{}

        mockClient.On("ListWorkflowSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &migrationhuborchestrator.ListWorkflowsInput{}
        output := &migrationhuborchestrator.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryWorkflowStep", func(t *testing.T) {
        input := &migrationhuborchestrator.RetryWorkflowStepInput{}
        output := &migrationhuborchestrator.RetryWorkflowStepOutput{}

        mockClient.On("RetryWorkflowStep", ctx, input).Return(output, nil)

        result, err := mockClient.RetryWorkflowStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.StartWorkflowInput{}
        output := &migrationhuborchestrator.StartWorkflowOutput{}

        mockClient.On("StartWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.StopWorkflowInput{}
        output := &migrationhuborchestrator.StopWorkflowOutput{}

        mockClient.On("StopWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.StopWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &migrationhuborchestrator.TagResourceInput{}
        output := &migrationhuborchestrator.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &migrationhuborchestrator.UntagResourceInput{}
        output := &migrationhuborchestrator.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplate", func(t *testing.T) {
        input := &migrationhuborchestrator.UpdateTemplateInput{}
        output := &migrationhuborchestrator.UpdateTemplateOutput{}

        mockClient.On("UpdateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkflow", func(t *testing.T) {
        input := &migrationhuborchestrator.UpdateWorkflowInput{}
        output := &migrationhuborchestrator.UpdateWorkflowOutput{}

        mockClient.On("UpdateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkflowStep", func(t *testing.T) {
        input := &migrationhuborchestrator.UpdateWorkflowStepInput{}
        output := &migrationhuborchestrator.UpdateWorkflowStepOutput{}

        mockClient.On("UpdateWorkflowStep", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkflowStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkflowStepGroup", func(t *testing.T) {
        input := &migrationhuborchestrator.UpdateWorkflowStepGroupInput{}
        output := &migrationhuborchestrator.UpdateWorkflowStepGroupOutput{}

        mockClient.On("UpdateWorkflowStepGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkflowStepGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
