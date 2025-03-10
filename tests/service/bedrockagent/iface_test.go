// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package bedrockagent_test

// tests for the bedrockagent service interface for 
// this ../../../service/bedrockagent/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockagent/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bedrockagent/bedrockagent_iface"
	"github.com/stretchr/testify/assert"
)

func TestBedrockagentServiceCanBeMocked(t *testing.T) {
	var iface bedrockagent_iface.IClient
	iface = &bedrockagent.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bedrockagent.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAgentCollaborator", func(t *testing.T) {
        input := &bedrockagent.AssociateAgentCollaboratorInput{}
        output := &bedrockagent.AssociateAgentCollaboratorOutput{}

        mockClient.On("AssociateAgentCollaborator", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAgentCollaborator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAgentKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.AssociateAgentKnowledgeBaseInput{}
        output := &bedrockagent.AssociateAgentKnowledgeBaseOutput{}

        mockClient.On("AssociateAgentKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAgentKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgent", func(t *testing.T) {
        input := &bedrockagent.CreateAgentInput{}
        output := &bedrockagent.CreateAgentOutput{}

        mockClient.On("CreateAgent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgentActionGroup", func(t *testing.T) {
        input := &bedrockagent.CreateAgentActionGroupInput{}
        output := &bedrockagent.CreateAgentActionGroupOutput{}

        mockClient.On("CreateAgentActionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgentActionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgentAlias", func(t *testing.T) {
        input := &bedrockagent.CreateAgentAliasInput{}
        output := &bedrockagent.CreateAgentAliasOutput{}

        mockClient.On("CreateAgentAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgentAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &bedrockagent.CreateDataSourceInput{}
        output := &bedrockagent.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlow", func(t *testing.T) {
        input := &bedrockagent.CreateFlowInput{}
        output := &bedrockagent.CreateFlowOutput{}

        mockClient.On("CreateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlowAlias", func(t *testing.T) {
        input := &bedrockagent.CreateFlowAliasInput{}
        output := &bedrockagent.CreateFlowAliasOutput{}

        mockClient.On("CreateFlowAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlowAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlowVersion", func(t *testing.T) {
        input := &bedrockagent.CreateFlowVersionInput{}
        output := &bedrockagent.CreateFlowVersionOutput{}

        mockClient.On("CreateFlowVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlowVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.CreateKnowledgeBaseInput{}
        output := &bedrockagent.CreateKnowledgeBaseOutput{}

        mockClient.On("CreateKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrompt", func(t *testing.T) {
        input := &bedrockagent.CreatePromptInput{}
        output := &bedrockagent.CreatePromptOutput{}

        mockClient.On("CreatePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePromptVersion", func(t *testing.T) {
        input := &bedrockagent.CreatePromptVersionInput{}
        output := &bedrockagent.CreatePromptVersionOutput{}

        mockClient.On("CreatePromptVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePromptVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgent", func(t *testing.T) {
        input := &bedrockagent.DeleteAgentInput{}
        output := &bedrockagent.DeleteAgentOutput{}

        mockClient.On("DeleteAgent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgentActionGroup", func(t *testing.T) {
        input := &bedrockagent.DeleteAgentActionGroupInput{}
        output := &bedrockagent.DeleteAgentActionGroupOutput{}

        mockClient.On("DeleteAgentActionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgentActionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgentAlias", func(t *testing.T) {
        input := &bedrockagent.DeleteAgentAliasInput{}
        output := &bedrockagent.DeleteAgentAliasOutput{}

        mockClient.On("DeleteAgentAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgentAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgentVersion", func(t *testing.T) {
        input := &bedrockagent.DeleteAgentVersionInput{}
        output := &bedrockagent.DeleteAgentVersionOutput{}

        mockClient.On("DeleteAgentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &bedrockagent.DeleteDataSourceInput{}
        output := &bedrockagent.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlow", func(t *testing.T) {
        input := &bedrockagent.DeleteFlowInput{}
        output := &bedrockagent.DeleteFlowOutput{}

        mockClient.On("DeleteFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlowAlias", func(t *testing.T) {
        input := &bedrockagent.DeleteFlowAliasInput{}
        output := &bedrockagent.DeleteFlowAliasOutput{}

        mockClient.On("DeleteFlowAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlowAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlowVersion", func(t *testing.T) {
        input := &bedrockagent.DeleteFlowVersionInput{}
        output := &bedrockagent.DeleteFlowVersionOutput{}

        mockClient.On("DeleteFlowVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlowVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.DeleteKnowledgeBaseInput{}
        output := &bedrockagent.DeleteKnowledgeBaseOutput{}

        mockClient.On("DeleteKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKnowledgeBaseDocuments", func(t *testing.T) {
        input := &bedrockagent.DeleteKnowledgeBaseDocumentsInput{}
        output := &bedrockagent.DeleteKnowledgeBaseDocumentsOutput{}

        mockClient.On("DeleteKnowledgeBaseDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKnowledgeBaseDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrompt", func(t *testing.T) {
        input := &bedrockagent.DeletePromptInput{}
        output := &bedrockagent.DeletePromptOutput{}

        mockClient.On("DeletePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAgentCollaborator", func(t *testing.T) {
        input := &bedrockagent.DisassociateAgentCollaboratorInput{}
        output := &bedrockagent.DisassociateAgentCollaboratorOutput{}

        mockClient.On("DisassociateAgentCollaborator", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAgentCollaborator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAgentKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.DisassociateAgentKnowledgeBaseInput{}
        output := &bedrockagent.DisassociateAgentKnowledgeBaseOutput{}

        mockClient.On("DisassociateAgentKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAgentKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgent", func(t *testing.T) {
        input := &bedrockagent.GetAgentInput{}
        output := &bedrockagent.GetAgentOutput{}

        mockClient.On("GetAgent", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentActionGroup", func(t *testing.T) {
        input := &bedrockagent.GetAgentActionGroupInput{}
        output := &bedrockagent.GetAgentActionGroupOutput{}

        mockClient.On("GetAgentActionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentActionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentAlias", func(t *testing.T) {
        input := &bedrockagent.GetAgentAliasInput{}
        output := &bedrockagent.GetAgentAliasOutput{}

        mockClient.On("GetAgentAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentCollaborator", func(t *testing.T) {
        input := &bedrockagent.GetAgentCollaboratorInput{}
        output := &bedrockagent.GetAgentCollaboratorOutput{}

        mockClient.On("GetAgentCollaborator", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentCollaborator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.GetAgentKnowledgeBaseInput{}
        output := &bedrockagent.GetAgentKnowledgeBaseOutput{}

        mockClient.On("GetAgentKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentVersion", func(t *testing.T) {
        input := &bedrockagent.GetAgentVersionInput{}
        output := &bedrockagent.GetAgentVersionOutput{}

        mockClient.On("GetAgentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &bedrockagent.GetDataSourceInput{}
        output := &bedrockagent.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlow", func(t *testing.T) {
        input := &bedrockagent.GetFlowInput{}
        output := &bedrockagent.GetFlowOutput{}

        mockClient.On("GetFlow", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowAlias", func(t *testing.T) {
        input := &bedrockagent.GetFlowAliasInput{}
        output := &bedrockagent.GetFlowAliasOutput{}

        mockClient.On("GetFlowAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowVersion", func(t *testing.T) {
        input := &bedrockagent.GetFlowVersionInput{}
        output := &bedrockagent.GetFlowVersionOutput{}

        mockClient.On("GetFlowVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIngestionJob", func(t *testing.T) {
        input := &bedrockagent.GetIngestionJobInput{}
        output := &bedrockagent.GetIngestionJobOutput{}

        mockClient.On("GetIngestionJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetIngestionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.GetKnowledgeBaseInput{}
        output := &bedrockagent.GetKnowledgeBaseOutput{}

        mockClient.On("GetKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.GetKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKnowledgeBaseDocuments", func(t *testing.T) {
        input := &bedrockagent.GetKnowledgeBaseDocumentsInput{}
        output := &bedrockagent.GetKnowledgeBaseDocumentsOutput{}

        mockClient.On("GetKnowledgeBaseDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.GetKnowledgeBaseDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrompt", func(t *testing.T) {
        input := &bedrockagent.GetPromptInput{}
        output := &bedrockagent.GetPromptOutput{}

        mockClient.On("GetPrompt", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIngestKnowledgeBaseDocuments", func(t *testing.T) {
        input := &bedrockagent.IngestKnowledgeBaseDocumentsInput{}
        output := &bedrockagent.IngestKnowledgeBaseDocumentsOutput{}

        mockClient.On("IngestKnowledgeBaseDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.IngestKnowledgeBaseDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentActionGroups", func(t *testing.T) {
        input := &bedrockagent.ListAgentActionGroupsInput{}
        output := &bedrockagent.ListAgentActionGroupsOutput{}

        mockClient.On("ListAgentActionGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentActionGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentAliases", func(t *testing.T) {
        input := &bedrockagent.ListAgentAliasesInput{}
        output := &bedrockagent.ListAgentAliasesOutput{}

        mockClient.On("ListAgentAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentCollaborators", func(t *testing.T) {
        input := &bedrockagent.ListAgentCollaboratorsInput{}
        output := &bedrockagent.ListAgentCollaboratorsOutput{}

        mockClient.On("ListAgentCollaborators", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentCollaborators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentKnowledgeBases", func(t *testing.T) {
        input := &bedrockagent.ListAgentKnowledgeBasesInput{}
        output := &bedrockagent.ListAgentKnowledgeBasesOutput{}

        mockClient.On("ListAgentKnowledgeBases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentKnowledgeBases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentVersions", func(t *testing.T) {
        input := &bedrockagent.ListAgentVersionsInput{}
        output := &bedrockagent.ListAgentVersionsOutput{}

        mockClient.On("ListAgentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgents", func(t *testing.T) {
        input := &bedrockagent.ListAgentsInput{}
        output := &bedrockagent.ListAgentsOutput{}

        mockClient.On("ListAgents", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &bedrockagent.ListDataSourcesInput{}
        output := &bedrockagent.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowAliases", func(t *testing.T) {
        input := &bedrockagent.ListFlowAliasesInput{}
        output := &bedrockagent.ListFlowAliasesOutput{}

        mockClient.On("ListFlowAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowVersions", func(t *testing.T) {
        input := &bedrockagent.ListFlowVersionsInput{}
        output := &bedrockagent.ListFlowVersionsOutput{}

        mockClient.On("ListFlowVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlows", func(t *testing.T) {
        input := &bedrockagent.ListFlowsInput{}
        output := &bedrockagent.ListFlowsOutput{}

        mockClient.On("ListFlows", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIngestionJobs", func(t *testing.T) {
        input := &bedrockagent.ListIngestionJobsInput{}
        output := &bedrockagent.ListIngestionJobsOutput{}

        mockClient.On("ListIngestionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListIngestionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKnowledgeBaseDocuments", func(t *testing.T) {
        input := &bedrockagent.ListKnowledgeBaseDocumentsInput{}
        output := &bedrockagent.ListKnowledgeBaseDocumentsOutput{}

        mockClient.On("ListKnowledgeBaseDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.ListKnowledgeBaseDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKnowledgeBases", func(t *testing.T) {
        input := &bedrockagent.ListKnowledgeBasesInput{}
        output := &bedrockagent.ListKnowledgeBasesOutput{}

        mockClient.On("ListKnowledgeBases", ctx, input).Return(output, nil)

        result, err := mockClient.ListKnowledgeBases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrompts", func(t *testing.T) {
        input := &bedrockagent.ListPromptsInput{}
        output := &bedrockagent.ListPromptsOutput{}

        mockClient.On("ListPrompts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrompts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &bedrockagent.ListTagsForResourceInput{}
        output := &bedrockagent.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPrepareAgent", func(t *testing.T) {
        input := &bedrockagent.PrepareAgentInput{}
        output := &bedrockagent.PrepareAgentOutput{}

        mockClient.On("PrepareAgent", ctx, input).Return(output, nil)

        result, err := mockClient.PrepareAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPrepareFlow", func(t *testing.T) {
        input := &bedrockagent.PrepareFlowInput{}
        output := &bedrockagent.PrepareFlowOutput{}

        mockClient.On("PrepareFlow", ctx, input).Return(output, nil)

        result, err := mockClient.PrepareFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartIngestionJob", func(t *testing.T) {
        input := &bedrockagent.StartIngestionJobInput{}
        output := &bedrockagent.StartIngestionJobOutput{}

        mockClient.On("StartIngestionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartIngestionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopIngestionJob", func(t *testing.T) {
        input := &bedrockagent.StopIngestionJobInput{}
        output := &bedrockagent.StopIngestionJobOutput{}

        mockClient.On("StopIngestionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopIngestionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &bedrockagent.TagResourceInput{}
        output := &bedrockagent.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &bedrockagent.UntagResourceInput{}
        output := &bedrockagent.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgent", func(t *testing.T) {
        input := &bedrockagent.UpdateAgentInput{}
        output := &bedrockagent.UpdateAgentOutput{}

        mockClient.On("UpdateAgent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentActionGroup", func(t *testing.T) {
        input := &bedrockagent.UpdateAgentActionGroupInput{}
        output := &bedrockagent.UpdateAgentActionGroupOutput{}

        mockClient.On("UpdateAgentActionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentActionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentAlias", func(t *testing.T) {
        input := &bedrockagent.UpdateAgentAliasInput{}
        output := &bedrockagent.UpdateAgentAliasOutput{}

        mockClient.On("UpdateAgentAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentCollaborator", func(t *testing.T) {
        input := &bedrockagent.UpdateAgentCollaboratorInput{}
        output := &bedrockagent.UpdateAgentCollaboratorOutput{}

        mockClient.On("UpdateAgentCollaborator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentCollaborator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.UpdateAgentKnowledgeBaseInput{}
        output := &bedrockagent.UpdateAgentKnowledgeBaseOutput{}

        mockClient.On("UpdateAgentKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &bedrockagent.UpdateDataSourceInput{}
        output := &bedrockagent.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlow", func(t *testing.T) {
        input := &bedrockagent.UpdateFlowInput{}
        output := &bedrockagent.UpdateFlowOutput{}

        mockClient.On("UpdateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlowAlias", func(t *testing.T) {
        input := &bedrockagent.UpdateFlowAliasInput{}
        output := &bedrockagent.UpdateFlowAliasOutput{}

        mockClient.On("UpdateFlowAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlowAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKnowledgeBase", func(t *testing.T) {
        input := &bedrockagent.UpdateKnowledgeBaseInput{}
        output := &bedrockagent.UpdateKnowledgeBaseOutput{}

        mockClient.On("UpdateKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrompt", func(t *testing.T) {
        input := &bedrockagent.UpdatePromptInput{}
        output := &bedrockagent.UpdatePromptOutput{}

        mockClient.On("UpdatePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateFlowDefinition", func(t *testing.T) {
        input := &bedrockagent.ValidateFlowDefinitionInput{}
        output := &bedrockagent.ValidateFlowDefinitionOutput{}

        mockClient.On("ValidateFlowDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateFlowDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
