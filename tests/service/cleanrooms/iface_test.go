package cleanrooms_test

// tests for the cleanrooms service interface for this ../../../service/cleanrooms/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cleanrooms"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cleanrooms/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cleanrooms/cleanrooms_iface"
	"github.com/stretchr/testify/assert"
)

func TestCleanroomsServiceCanBeMocked(t *testing.T) {
	var iface cleanrooms_iface.IClient
	iface = &cleanrooms.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cleanrooms.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCollaborationAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.BatchGetCollaborationAnalysisTemplateInput{}
        output := &cleanrooms.BatchGetCollaborationAnalysisTemplateOutput{}

        mockClient.On("BatchGetCollaborationAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCollaborationAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetSchema", func(t *testing.T) {
        input := &cleanrooms.BatchGetSchemaInput{}
        output := &cleanrooms.BatchGetSchemaOutput{}

        mockClient.On("BatchGetSchema", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetSchemaAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.BatchGetSchemaAnalysisRuleInput{}
        output := &cleanrooms.BatchGetSchemaAnalysisRuleOutput{}

        mockClient.On("BatchGetSchemaAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetSchemaAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.CreateAnalysisTemplateInput{}
        output := &cleanrooms.CreateAnalysisTemplateOutput{}

        mockClient.On("CreateAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCollaboration", func(t *testing.T) {
        input := &cleanrooms.CreateCollaborationInput{}
        output := &cleanrooms.CreateCollaborationOutput{}

        mockClient.On("CreateCollaboration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCollaboration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguredAudienceModelAssociation", func(t *testing.T) {
        input := &cleanrooms.CreateConfiguredAudienceModelAssociationInput{}
        output := &cleanrooms.CreateConfiguredAudienceModelAssociationOutput{}

        mockClient.On("CreateConfiguredAudienceModelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguredAudienceModelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguredTable", func(t *testing.T) {
        input := &cleanrooms.CreateConfiguredTableInput{}
        output := &cleanrooms.CreateConfiguredTableOutput{}

        mockClient.On("CreateConfiguredTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguredTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguredTableAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.CreateConfiguredTableAnalysisRuleInput{}
        output := &cleanrooms.CreateConfiguredTableAnalysisRuleOutput{}

        mockClient.On("CreateConfiguredTableAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguredTableAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguredTableAssociation", func(t *testing.T) {
        input := &cleanrooms.CreateConfiguredTableAssociationInput{}
        output := &cleanrooms.CreateConfiguredTableAssociationOutput{}

        mockClient.On("CreateConfiguredTableAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguredTableAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMembership", func(t *testing.T) {
        input := &cleanrooms.CreateMembershipInput{}
        output := &cleanrooms.CreateMembershipOutput{}

        mockClient.On("CreateMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrivacyBudgetTemplate", func(t *testing.T) {
        input := &cleanrooms.CreatePrivacyBudgetTemplateInput{}
        output := &cleanrooms.CreatePrivacyBudgetTemplateOutput{}

        mockClient.On("CreatePrivacyBudgetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrivacyBudgetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.DeleteAnalysisTemplateInput{}
        output := &cleanrooms.DeleteAnalysisTemplateOutput{}

        mockClient.On("DeleteAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCollaboration", func(t *testing.T) {
        input := &cleanrooms.DeleteCollaborationInput{}
        output := &cleanrooms.DeleteCollaborationOutput{}

        mockClient.On("DeleteCollaboration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCollaboration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredAudienceModelAssociation", func(t *testing.T) {
        input := &cleanrooms.DeleteConfiguredAudienceModelAssociationInput{}
        output := &cleanrooms.DeleteConfiguredAudienceModelAssociationOutput{}

        mockClient.On("DeleteConfiguredAudienceModelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredAudienceModelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredTable", func(t *testing.T) {
        input := &cleanrooms.DeleteConfiguredTableInput{}
        output := &cleanrooms.DeleteConfiguredTableOutput{}

        mockClient.On("DeleteConfiguredTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredTableAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.DeleteConfiguredTableAnalysisRuleInput{}
        output := &cleanrooms.DeleteConfiguredTableAnalysisRuleOutput{}

        mockClient.On("DeleteConfiguredTableAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredTableAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguredTableAssociation", func(t *testing.T) {
        input := &cleanrooms.DeleteConfiguredTableAssociationInput{}
        output := &cleanrooms.DeleteConfiguredTableAssociationOutput{}

        mockClient.On("DeleteConfiguredTableAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguredTableAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMember", func(t *testing.T) {
        input := &cleanrooms.DeleteMemberInput{}
        output := &cleanrooms.DeleteMemberOutput{}

        mockClient.On("DeleteMember", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMembership", func(t *testing.T) {
        input := &cleanrooms.DeleteMembershipInput{}
        output := &cleanrooms.DeleteMembershipOutput{}

        mockClient.On("DeleteMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrivacyBudgetTemplate", func(t *testing.T) {
        input := &cleanrooms.DeletePrivacyBudgetTemplateInput{}
        output := &cleanrooms.DeletePrivacyBudgetTemplateOutput{}

        mockClient.On("DeletePrivacyBudgetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrivacyBudgetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.GetAnalysisTemplateInput{}
        output := &cleanrooms.GetAnalysisTemplateOutput{}

        mockClient.On("GetAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCollaboration", func(t *testing.T) {
        input := &cleanrooms.GetCollaborationInput{}
        output := &cleanrooms.GetCollaborationOutput{}

        mockClient.On("GetCollaboration", ctx, input).Return(output, nil)

        result, err := mockClient.GetCollaboration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCollaborationAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.GetCollaborationAnalysisTemplateInput{}
        output := &cleanrooms.GetCollaborationAnalysisTemplateOutput{}

        mockClient.On("GetCollaborationAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCollaborationAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCollaborationConfiguredAudienceModelAssociation", func(t *testing.T) {
        input := &cleanrooms.GetCollaborationConfiguredAudienceModelAssociationInput{}
        output := &cleanrooms.GetCollaborationConfiguredAudienceModelAssociationOutput{}

        mockClient.On("GetCollaborationConfiguredAudienceModelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetCollaborationConfiguredAudienceModelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCollaborationPrivacyBudgetTemplate", func(t *testing.T) {
        input := &cleanrooms.GetCollaborationPrivacyBudgetTemplateInput{}
        output := &cleanrooms.GetCollaborationPrivacyBudgetTemplateOutput{}

        mockClient.On("GetCollaborationPrivacyBudgetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCollaborationPrivacyBudgetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredAudienceModelAssociation", func(t *testing.T) {
        input := &cleanrooms.GetConfiguredAudienceModelAssociationInput{}
        output := &cleanrooms.GetConfiguredAudienceModelAssociationOutput{}

        mockClient.On("GetConfiguredAudienceModelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredAudienceModelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredTable", func(t *testing.T) {
        input := &cleanrooms.GetConfiguredTableInput{}
        output := &cleanrooms.GetConfiguredTableOutput{}

        mockClient.On("GetConfiguredTable", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredTableAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.GetConfiguredTableAnalysisRuleInput{}
        output := &cleanrooms.GetConfiguredTableAnalysisRuleOutput{}

        mockClient.On("GetConfiguredTableAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredTableAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguredTableAssociation", func(t *testing.T) {
        input := &cleanrooms.GetConfiguredTableAssociationInput{}
        output := &cleanrooms.GetConfiguredTableAssociationOutput{}

        mockClient.On("GetConfiguredTableAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguredTableAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMembership", func(t *testing.T) {
        input := &cleanrooms.GetMembershipInput{}
        output := &cleanrooms.GetMembershipOutput{}

        mockClient.On("GetMembership", ctx, input).Return(output, nil)

        result, err := mockClient.GetMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrivacyBudgetTemplate", func(t *testing.T) {
        input := &cleanrooms.GetPrivacyBudgetTemplateInput{}
        output := &cleanrooms.GetPrivacyBudgetTemplateOutput{}

        mockClient.On("GetPrivacyBudgetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrivacyBudgetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProtectedQuery", func(t *testing.T) {
        input := &cleanrooms.GetProtectedQueryInput{}
        output := &cleanrooms.GetProtectedQueryOutput{}

        mockClient.On("GetProtectedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.GetProtectedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchema", func(t *testing.T) {
        input := &cleanrooms.GetSchemaInput{}
        output := &cleanrooms.GetSchemaOutput{}

        mockClient.On("GetSchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.GetSchemaAnalysisRuleInput{}
        output := &cleanrooms.GetSchemaAnalysisRuleOutput{}

        mockClient.On("GetSchemaAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalysisTemplates", func(t *testing.T) {
        input := &cleanrooms.ListAnalysisTemplatesInput{}
        output := &cleanrooms.ListAnalysisTemplatesOutput{}

        mockClient.On("ListAnalysisTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalysisTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollaborationAnalysisTemplates", func(t *testing.T) {
        input := &cleanrooms.ListCollaborationAnalysisTemplatesInput{}
        output := &cleanrooms.ListCollaborationAnalysisTemplatesOutput{}

        mockClient.On("ListCollaborationAnalysisTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollaborationAnalysisTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollaborationConfiguredAudienceModelAssociations", func(t *testing.T) {
        input := &cleanrooms.ListCollaborationConfiguredAudienceModelAssociationsInput{}
        output := &cleanrooms.ListCollaborationConfiguredAudienceModelAssociationsOutput{}

        mockClient.On("ListCollaborationConfiguredAudienceModelAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollaborationConfiguredAudienceModelAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollaborationPrivacyBudgetTemplates", func(t *testing.T) {
        input := &cleanrooms.ListCollaborationPrivacyBudgetTemplatesInput{}
        output := &cleanrooms.ListCollaborationPrivacyBudgetTemplatesOutput{}

        mockClient.On("ListCollaborationPrivacyBudgetTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollaborationPrivacyBudgetTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollaborationPrivacyBudgets", func(t *testing.T) {
        input := &cleanrooms.ListCollaborationPrivacyBudgetsInput{}
        output := &cleanrooms.ListCollaborationPrivacyBudgetsOutput{}

        mockClient.On("ListCollaborationPrivacyBudgets", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollaborationPrivacyBudgets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollaborations", func(t *testing.T) {
        input := &cleanrooms.ListCollaborationsInput{}
        output := &cleanrooms.ListCollaborationsOutput{}

        mockClient.On("ListCollaborations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollaborations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfiguredAudienceModelAssociations", func(t *testing.T) {
        input := &cleanrooms.ListConfiguredAudienceModelAssociationsInput{}
        output := &cleanrooms.ListConfiguredAudienceModelAssociationsOutput{}

        mockClient.On("ListConfiguredAudienceModelAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfiguredAudienceModelAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfiguredTableAssociations", func(t *testing.T) {
        input := &cleanrooms.ListConfiguredTableAssociationsInput{}
        output := &cleanrooms.ListConfiguredTableAssociationsOutput{}

        mockClient.On("ListConfiguredTableAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfiguredTableAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfiguredTables", func(t *testing.T) {
        input := &cleanrooms.ListConfiguredTablesInput{}
        output := &cleanrooms.ListConfiguredTablesOutput{}

        mockClient.On("ListConfiguredTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfiguredTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &cleanrooms.ListMembersInput{}
        output := &cleanrooms.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMemberships", func(t *testing.T) {
        input := &cleanrooms.ListMembershipsInput{}
        output := &cleanrooms.ListMembershipsOutput{}

        mockClient.On("ListMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrivacyBudgetTemplates", func(t *testing.T) {
        input := &cleanrooms.ListPrivacyBudgetTemplatesInput{}
        output := &cleanrooms.ListPrivacyBudgetTemplatesOutput{}

        mockClient.On("ListPrivacyBudgetTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrivacyBudgetTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrivacyBudgets", func(t *testing.T) {
        input := &cleanrooms.ListPrivacyBudgetsInput{}
        output := &cleanrooms.ListPrivacyBudgetsOutput{}

        mockClient.On("ListPrivacyBudgets", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrivacyBudgets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtectedQueries", func(t *testing.T) {
        input := &cleanrooms.ListProtectedQueriesInput{}
        output := &cleanrooms.ListProtectedQueriesOutput{}

        mockClient.On("ListProtectedQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtectedQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemas", func(t *testing.T) {
        input := &cleanrooms.ListSchemasInput{}
        output := &cleanrooms.ListSchemasOutput{}

        mockClient.On("ListSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cleanrooms.ListTagsForResourceInput{}
        output := &cleanrooms.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPreviewPrivacyImpact", func(t *testing.T) {
        input := &cleanrooms.PreviewPrivacyImpactInput{}
        output := &cleanrooms.PreviewPrivacyImpactOutput{}

        mockClient.On("PreviewPrivacyImpact", ctx, input).Return(output, nil)

        result, err := mockClient.PreviewPrivacyImpact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartProtectedQuery", func(t *testing.T) {
        input := &cleanrooms.StartProtectedQueryInput{}
        output := &cleanrooms.StartProtectedQueryOutput{}

        mockClient.On("StartProtectedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StartProtectedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cleanrooms.TagResourceInput{}
        output := &cleanrooms.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cleanrooms.UntagResourceInput{}
        output := &cleanrooms.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnalysisTemplate", func(t *testing.T) {
        input := &cleanrooms.UpdateAnalysisTemplateInput{}
        output := &cleanrooms.UpdateAnalysisTemplateOutput{}

        mockClient.On("UpdateAnalysisTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnalysisTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCollaboration", func(t *testing.T) {
        input := &cleanrooms.UpdateCollaborationInput{}
        output := &cleanrooms.UpdateCollaborationOutput{}

        mockClient.On("UpdateCollaboration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCollaboration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguredAudienceModelAssociation", func(t *testing.T) {
        input := &cleanrooms.UpdateConfiguredAudienceModelAssociationInput{}
        output := &cleanrooms.UpdateConfiguredAudienceModelAssociationOutput{}

        mockClient.On("UpdateConfiguredAudienceModelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguredAudienceModelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguredTable", func(t *testing.T) {
        input := &cleanrooms.UpdateConfiguredTableInput{}
        output := &cleanrooms.UpdateConfiguredTableOutput{}

        mockClient.On("UpdateConfiguredTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguredTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguredTableAnalysisRule", func(t *testing.T) {
        input := &cleanrooms.UpdateConfiguredTableAnalysisRuleInput{}
        output := &cleanrooms.UpdateConfiguredTableAnalysisRuleOutput{}

        mockClient.On("UpdateConfiguredTableAnalysisRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguredTableAnalysisRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguredTableAssociation", func(t *testing.T) {
        input := &cleanrooms.UpdateConfiguredTableAssociationInput{}
        output := &cleanrooms.UpdateConfiguredTableAssociationOutput{}

        mockClient.On("UpdateConfiguredTableAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguredTableAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMembership", func(t *testing.T) {
        input := &cleanrooms.UpdateMembershipInput{}
        output := &cleanrooms.UpdateMembershipOutput{}

        mockClient.On("UpdateMembership", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrivacyBudgetTemplate", func(t *testing.T) {
        input := &cleanrooms.UpdatePrivacyBudgetTemplateInput{}
        output := &cleanrooms.UpdatePrivacyBudgetTemplateOutput{}

        mockClient.On("UpdatePrivacyBudgetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrivacyBudgetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProtectedQuery", func(t *testing.T) {
        input := &cleanrooms.UpdateProtectedQueryInput{}
        output := &cleanrooms.UpdateProtectedQueryOutput{}

        mockClient.On("UpdateProtectedQuery", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProtectedQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
