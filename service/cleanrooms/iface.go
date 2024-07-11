
package cleanrooms

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/cleanrooms"
)

// IClient defines the interface for cleanrooms
type IClient interface {
 Options() Options 
 BatchGetCollaborationAnalysisTemplate(ctx context.Context, params *BatchGetCollaborationAnalysisTemplateInput, optFns ...func(*Options)) (*BatchGetCollaborationAnalysisTemplateOutput, error) 
 BatchGetSchema(ctx context.Context, params *BatchGetSchemaInput, optFns ...func(*Options)) (*BatchGetSchemaOutput, error) 
 BatchGetSchemaAnalysisRule(ctx context.Context, params *BatchGetSchemaAnalysisRuleInput, optFns ...func(*Options)) (*BatchGetSchemaAnalysisRuleOutput, error) 
 CreateAnalysisTemplate(ctx context.Context, params *CreateAnalysisTemplateInput, optFns ...func(*Options)) (*CreateAnalysisTemplateOutput, error) 
 CreateCollaboration(ctx context.Context, params *CreateCollaborationInput, optFns ...func(*Options)) (*CreateCollaborationOutput, error) 
 CreateConfiguredAudienceModelAssociation(ctx context.Context, params *CreateConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*CreateConfiguredAudienceModelAssociationOutput, error) 
 CreateConfiguredTable(ctx context.Context, params *CreateConfiguredTableInput, optFns ...func(*Options)) (*CreateConfiguredTableOutput, error) 
 CreateConfiguredTableAnalysisRule(ctx context.Context, params *CreateConfiguredTableAnalysisRuleInput, optFns ...func(*Options)) (*CreateConfiguredTableAnalysisRuleOutput, error) 
 CreateConfiguredTableAssociation(ctx context.Context, params *CreateConfiguredTableAssociationInput, optFns ...func(*Options)) (*CreateConfiguredTableAssociationOutput, error) 
 CreateMembership(ctx context.Context, params *CreateMembershipInput, optFns ...func(*Options)) (*CreateMembershipOutput, error) 
 CreatePrivacyBudgetTemplate(ctx context.Context, params *CreatePrivacyBudgetTemplateInput, optFns ...func(*Options)) (*CreatePrivacyBudgetTemplateOutput, error) 
 DeleteAnalysisTemplate(ctx context.Context, params *DeleteAnalysisTemplateInput, optFns ...func(*Options)) (*DeleteAnalysisTemplateOutput, error) 
 DeleteCollaboration(ctx context.Context, params *DeleteCollaborationInput, optFns ...func(*Options)) (*DeleteCollaborationOutput, error) 
 DeleteConfiguredAudienceModelAssociation(ctx context.Context, params *DeleteConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*DeleteConfiguredAudienceModelAssociationOutput, error) 
 DeleteConfiguredTable(ctx context.Context, params *DeleteConfiguredTableInput, optFns ...func(*Options)) (*DeleteConfiguredTableOutput, error) 
 DeleteConfiguredTableAnalysisRule(ctx context.Context, params *DeleteConfiguredTableAnalysisRuleInput, optFns ...func(*Options)) (*DeleteConfiguredTableAnalysisRuleOutput, error) 
 DeleteConfiguredTableAssociation(ctx context.Context, params *DeleteConfiguredTableAssociationInput, optFns ...func(*Options)) (*DeleteConfiguredTableAssociationOutput, error) 
 DeleteMember(ctx context.Context, params *DeleteMemberInput, optFns ...func(*Options)) (*DeleteMemberOutput, error) 
 DeleteMembership(ctx context.Context, params *DeleteMembershipInput, optFns ...func(*Options)) (*DeleteMembershipOutput, error) 
 DeletePrivacyBudgetTemplate(ctx context.Context, params *DeletePrivacyBudgetTemplateInput, optFns ...func(*Options)) (*DeletePrivacyBudgetTemplateOutput, error) 
 GetAnalysisTemplate(ctx context.Context, params *GetAnalysisTemplateInput, optFns ...func(*Options)) (*GetAnalysisTemplateOutput, error) 
 GetCollaboration(ctx context.Context, params *GetCollaborationInput, optFns ...func(*Options)) (*GetCollaborationOutput, error) 
 GetCollaborationAnalysisTemplate(ctx context.Context, params *GetCollaborationAnalysisTemplateInput, optFns ...func(*Options)) (*GetCollaborationAnalysisTemplateOutput, error) 
 GetCollaborationConfiguredAudienceModelAssociation(ctx context.Context, params *GetCollaborationConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*GetCollaborationConfiguredAudienceModelAssociationOutput, error) 
 GetCollaborationPrivacyBudgetTemplate(ctx context.Context, params *GetCollaborationPrivacyBudgetTemplateInput, optFns ...func(*Options)) (*GetCollaborationPrivacyBudgetTemplateOutput, error) 
 GetConfiguredAudienceModelAssociation(ctx context.Context, params *GetConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*GetConfiguredAudienceModelAssociationOutput, error) 
 GetConfiguredTable(ctx context.Context, params *GetConfiguredTableInput, optFns ...func(*Options)) (*GetConfiguredTableOutput, error) 
 GetConfiguredTableAnalysisRule(ctx context.Context, params *GetConfiguredTableAnalysisRuleInput, optFns ...func(*Options)) (*GetConfiguredTableAnalysisRuleOutput, error) 
 GetConfiguredTableAssociation(ctx context.Context, params *GetConfiguredTableAssociationInput, optFns ...func(*Options)) (*GetConfiguredTableAssociationOutput, error) 
 GetMembership(ctx context.Context, params *GetMembershipInput, optFns ...func(*Options)) (*GetMembershipOutput, error) 
 GetPrivacyBudgetTemplate(ctx context.Context, params *GetPrivacyBudgetTemplateInput, optFns ...func(*Options)) (*GetPrivacyBudgetTemplateOutput, error) 
 GetProtectedQuery(ctx context.Context, params *GetProtectedQueryInput, optFns ...func(*Options)) (*GetProtectedQueryOutput, error) 
 GetSchema(ctx context.Context, params *GetSchemaInput, optFns ...func(*Options)) (*GetSchemaOutput, error) 
 GetSchemaAnalysisRule(ctx context.Context, params *GetSchemaAnalysisRuleInput, optFns ...func(*Options)) (*GetSchemaAnalysisRuleOutput, error) 
 ListAnalysisTemplates(ctx context.Context, params *ListAnalysisTemplatesInput, optFns ...func(*Options)) (*ListAnalysisTemplatesOutput, error) 
 ListCollaborationAnalysisTemplates(ctx context.Context, params *ListCollaborationAnalysisTemplatesInput, optFns ...func(*Options)) (*ListCollaborationAnalysisTemplatesOutput, error) 
 ListCollaborationConfiguredAudienceModelAssociations(ctx context.Context, params *ListCollaborationConfiguredAudienceModelAssociationsInput, optFns ...func(*Options)) (*ListCollaborationConfiguredAudienceModelAssociationsOutput, error) 
 ListCollaborationPrivacyBudgetTemplates(ctx context.Context, params *ListCollaborationPrivacyBudgetTemplatesInput, optFns ...func(*Options)) (*ListCollaborationPrivacyBudgetTemplatesOutput, error) 
 ListCollaborationPrivacyBudgets(ctx context.Context, params *ListCollaborationPrivacyBudgetsInput, optFns ...func(*Options)) (*ListCollaborationPrivacyBudgetsOutput, error) 
 ListCollaborations(ctx context.Context, params *ListCollaborationsInput, optFns ...func(*Options)) (*ListCollaborationsOutput, error) 
 ListConfiguredAudienceModelAssociations(ctx context.Context, params *ListConfiguredAudienceModelAssociationsInput, optFns ...func(*Options)) (*ListConfiguredAudienceModelAssociationsOutput, error) 
 ListConfiguredTableAssociations(ctx context.Context, params *ListConfiguredTableAssociationsInput, optFns ...func(*Options)) (*ListConfiguredTableAssociationsOutput, error) 
 ListConfiguredTables(ctx context.Context, params *ListConfiguredTablesInput, optFns ...func(*Options)) (*ListConfiguredTablesOutput, error) 
 ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) 
 ListMemberships(ctx context.Context, params *ListMembershipsInput, optFns ...func(*Options)) (*ListMembershipsOutput, error) 
 ListPrivacyBudgetTemplates(ctx context.Context, params *ListPrivacyBudgetTemplatesInput, optFns ...func(*Options)) (*ListPrivacyBudgetTemplatesOutput, error) 
 ListPrivacyBudgets(ctx context.Context, params *ListPrivacyBudgetsInput, optFns ...func(*Options)) (*ListPrivacyBudgetsOutput, error) 
 ListProtectedQueries(ctx context.Context, params *ListProtectedQueriesInput, optFns ...func(*Options)) (*ListProtectedQueriesOutput, error) 
 ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PreviewPrivacyImpact(ctx context.Context, params *PreviewPrivacyImpactInput, optFns ...func(*Options)) (*PreviewPrivacyImpactOutput, error) 
 StartProtectedQuery(ctx context.Context, params *StartProtectedQueryInput, optFns ...func(*Options)) (*StartProtectedQueryOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAnalysisTemplate(ctx context.Context, params *UpdateAnalysisTemplateInput, optFns ...func(*Options)) (*UpdateAnalysisTemplateOutput, error) 
 UpdateCollaboration(ctx context.Context, params *UpdateCollaborationInput, optFns ...func(*Options)) (*UpdateCollaborationOutput, error) 
 UpdateConfiguredAudienceModelAssociation(ctx context.Context, params *UpdateConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*UpdateConfiguredAudienceModelAssociationOutput, error) 
 UpdateConfiguredTable(ctx context.Context, params *UpdateConfiguredTableInput, optFns ...func(*Options)) (*UpdateConfiguredTableOutput, error) 
 UpdateConfiguredTableAnalysisRule(ctx context.Context, params *UpdateConfiguredTableAnalysisRuleInput, optFns ...func(*Options)) (*UpdateConfiguredTableAnalysisRuleOutput, error) 
 UpdateConfiguredTableAssociation(ctx context.Context, params *UpdateConfiguredTableAssociationInput, optFns ...func(*Options)) (*UpdateConfiguredTableAssociationOutput, error) 
 UpdateMembership(ctx context.Context, params *UpdateMembershipInput, optFns ...func(*Options)) (*UpdateMembershipOutput, error) 
 UpdatePrivacyBudgetTemplate(ctx context.Context, params *UpdatePrivacyBudgetTemplateInput, optFns ...func(*Options)) (*UpdatePrivacyBudgetTemplateOutput, error) 
 UpdateProtectedQuery(ctx context.Context, params *UpdateProtectedQueryInput, optFns ...func(*Options)) (*UpdateProtectedQueryOutput, error) 
}
