
package auditmanager_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/auditmanager"
)

// IClient defines the interface for auditmanager
type IClient interface {
 Options() Options 
 AssociateAssessmentReportEvidenceFolder(ctx context.Context, params *AssociateAssessmentReportEvidenceFolderInput, optFns ...func(*Options)) (*AssociateAssessmentReportEvidenceFolderOutput, error) 
 BatchAssociateAssessmentReportEvidence(ctx context.Context, params *BatchAssociateAssessmentReportEvidenceInput, optFns ...func(*Options)) (*BatchAssociateAssessmentReportEvidenceOutput, error) 
 BatchCreateDelegationByAssessment(ctx context.Context, params *BatchCreateDelegationByAssessmentInput, optFns ...func(*Options)) (*BatchCreateDelegationByAssessmentOutput, error) 
 BatchDeleteDelegationByAssessment(ctx context.Context, params *BatchDeleteDelegationByAssessmentInput, optFns ...func(*Options)) (*BatchDeleteDelegationByAssessmentOutput, error) 
 BatchDisassociateAssessmentReportEvidence(ctx context.Context, params *BatchDisassociateAssessmentReportEvidenceInput, optFns ...func(*Options)) (*BatchDisassociateAssessmentReportEvidenceOutput, error) 
 BatchImportEvidenceToAssessmentControl(ctx context.Context, params *BatchImportEvidenceToAssessmentControlInput, optFns ...func(*Options)) (*BatchImportEvidenceToAssessmentControlOutput, error) 
 CreateAssessment(ctx context.Context, params *CreateAssessmentInput, optFns ...func(*Options)) (*CreateAssessmentOutput, error) 
 CreateAssessmentFramework(ctx context.Context, params *CreateAssessmentFrameworkInput, optFns ...func(*Options)) (*CreateAssessmentFrameworkOutput, error) 
 CreateAssessmentReport(ctx context.Context, params *CreateAssessmentReportInput, optFns ...func(*Options)) (*CreateAssessmentReportOutput, error) 
 CreateControl(ctx context.Context, params *CreateControlInput, optFns ...func(*Options)) (*CreateControlOutput, error) 
 DeleteAssessment(ctx context.Context, params *DeleteAssessmentInput, optFns ...func(*Options)) (*DeleteAssessmentOutput, error) 
 DeleteAssessmentFramework(ctx context.Context, params *DeleteAssessmentFrameworkInput, optFns ...func(*Options)) (*DeleteAssessmentFrameworkOutput, error) 
 DeleteAssessmentFrameworkShare(ctx context.Context, params *DeleteAssessmentFrameworkShareInput, optFns ...func(*Options)) (*DeleteAssessmentFrameworkShareOutput, error) 
 DeleteAssessmentReport(ctx context.Context, params *DeleteAssessmentReportInput, optFns ...func(*Options)) (*DeleteAssessmentReportOutput, error) 
 DeleteControl(ctx context.Context, params *DeleteControlInput, optFns ...func(*Options)) (*DeleteControlOutput, error) 
 DeregisterAccount(ctx context.Context, params *DeregisterAccountInput, optFns ...func(*Options)) (*DeregisterAccountOutput, error) 
 DeregisterOrganizationAdminAccount(ctx context.Context, params *DeregisterOrganizationAdminAccountInput, optFns ...func(*Options)) (*DeregisterOrganizationAdminAccountOutput, error) 
 DisassociateAssessmentReportEvidenceFolder(ctx context.Context, params *DisassociateAssessmentReportEvidenceFolderInput, optFns ...func(*Options)) (*DisassociateAssessmentReportEvidenceFolderOutput, error) 
 GetAccountStatus(ctx context.Context, params *GetAccountStatusInput, optFns ...func(*Options)) (*GetAccountStatusOutput, error) 
 GetAssessment(ctx context.Context, params *GetAssessmentInput, optFns ...func(*Options)) (*GetAssessmentOutput, error) 
 GetAssessmentFramework(ctx context.Context, params *GetAssessmentFrameworkInput, optFns ...func(*Options)) (*GetAssessmentFrameworkOutput, error) 
 GetAssessmentReportUrl(ctx context.Context, params *GetAssessmentReportUrlInput, optFns ...func(*Options)) (*GetAssessmentReportUrlOutput, error) 
 GetChangeLogs(ctx context.Context, params *GetChangeLogsInput, optFns ...func(*Options)) (*GetChangeLogsOutput, error) 
 GetControl(ctx context.Context, params *GetControlInput, optFns ...func(*Options)) (*GetControlOutput, error) 
 GetDelegations(ctx context.Context, params *GetDelegationsInput, optFns ...func(*Options)) (*GetDelegationsOutput, error) 
 GetEvidence(ctx context.Context, params *GetEvidenceInput, optFns ...func(*Options)) (*GetEvidenceOutput, error) 
 GetEvidenceByEvidenceFolder(ctx context.Context, params *GetEvidenceByEvidenceFolderInput, optFns ...func(*Options)) (*GetEvidenceByEvidenceFolderOutput, error) 
 GetEvidenceFileUploadUrl(ctx context.Context, params *GetEvidenceFileUploadUrlInput, optFns ...func(*Options)) (*GetEvidenceFileUploadUrlOutput, error) 
 GetEvidenceFolder(ctx context.Context, params *GetEvidenceFolderInput, optFns ...func(*Options)) (*GetEvidenceFolderOutput, error) 
 GetEvidenceFoldersByAssessment(ctx context.Context, params *GetEvidenceFoldersByAssessmentInput, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentOutput, error) 
 GetEvidenceFoldersByAssessmentControl(ctx context.Context, params *GetEvidenceFoldersByAssessmentControlInput, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentControlOutput, error) 
 GetInsights(ctx context.Context, params *GetInsightsInput, optFns ...func(*Options)) (*GetInsightsOutput, error) 
 GetInsightsByAssessment(ctx context.Context, params *GetInsightsByAssessmentInput, optFns ...func(*Options)) (*GetInsightsByAssessmentOutput, error) 
 GetOrganizationAdminAccount(ctx context.Context, params *GetOrganizationAdminAccountInput, optFns ...func(*Options)) (*GetOrganizationAdminAccountOutput, error) 
 GetServicesInScope(ctx context.Context, params *GetServicesInScopeInput, optFns ...func(*Options)) (*GetServicesInScopeOutput, error) 
 GetSettings(ctx context.Context, params *GetSettingsInput, optFns ...func(*Options)) (*GetSettingsOutput, error) 
 ListAssessmentControlInsightsByControlDomain(ctx context.Context, params *ListAssessmentControlInsightsByControlDomainInput, optFns ...func(*Options)) (*ListAssessmentControlInsightsByControlDomainOutput, error) 
 ListAssessmentFrameworkShareRequests(ctx context.Context, params *ListAssessmentFrameworkShareRequestsInput, optFns ...func(*Options)) (*ListAssessmentFrameworkShareRequestsOutput, error) 
 ListAssessmentFrameworks(ctx context.Context, params *ListAssessmentFrameworksInput, optFns ...func(*Options)) (*ListAssessmentFrameworksOutput, error) 
 ListAssessmentReports(ctx context.Context, params *ListAssessmentReportsInput, optFns ...func(*Options)) (*ListAssessmentReportsOutput, error) 
 ListAssessments(ctx context.Context, params *ListAssessmentsInput, optFns ...func(*Options)) (*ListAssessmentsOutput, error) 
 ListControlDomainInsights(ctx context.Context, params *ListControlDomainInsightsInput, optFns ...func(*Options)) (*ListControlDomainInsightsOutput, error) 
 ListControlDomainInsightsByAssessment(ctx context.Context, params *ListControlDomainInsightsByAssessmentInput, optFns ...func(*Options)) (*ListControlDomainInsightsByAssessmentOutput, error) 
 ListControlInsightsByControlDomain(ctx context.Context, params *ListControlInsightsByControlDomainInput, optFns ...func(*Options)) (*ListControlInsightsByControlDomainOutput, error) 
 ListControls(ctx context.Context, params *ListControlsInput, optFns ...func(*Options)) (*ListControlsOutput, error) 
 ListKeywordsForDataSource(ctx context.Context, params *ListKeywordsForDataSourceInput, optFns ...func(*Options)) (*ListKeywordsForDataSourceOutput, error) 
 ListNotifications(ctx context.Context, params *ListNotificationsInput, optFns ...func(*Options)) (*ListNotificationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterAccount(ctx context.Context, params *RegisterAccountInput, optFns ...func(*Options)) (*RegisterAccountOutput, error) 
 RegisterOrganizationAdminAccount(ctx context.Context, params *RegisterOrganizationAdminAccountInput, optFns ...func(*Options)) (*RegisterOrganizationAdminAccountOutput, error) 
 StartAssessmentFrameworkShare(ctx context.Context, params *StartAssessmentFrameworkShareInput, optFns ...func(*Options)) (*StartAssessmentFrameworkShareOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAssessment(ctx context.Context, params *UpdateAssessmentInput, optFns ...func(*Options)) (*UpdateAssessmentOutput, error) 
 UpdateAssessmentControl(ctx context.Context, params *UpdateAssessmentControlInput, optFns ...func(*Options)) (*UpdateAssessmentControlOutput, error) 
 UpdateAssessmentControlSetStatus(ctx context.Context, params *UpdateAssessmentControlSetStatusInput, optFns ...func(*Options)) (*UpdateAssessmentControlSetStatusOutput, error) 
 UpdateAssessmentFramework(ctx context.Context, params *UpdateAssessmentFrameworkInput, optFns ...func(*Options)) (*UpdateAssessmentFrameworkOutput, error) 
 UpdateAssessmentFrameworkShare(ctx context.Context, params *UpdateAssessmentFrameworkShareInput, optFns ...func(*Options)) (*UpdateAssessmentFrameworkShareOutput, error) 
 UpdateAssessmentStatus(ctx context.Context, params *UpdateAssessmentStatusInput, optFns ...func(*Options)) (*UpdateAssessmentStatusOutput, error) 
 UpdateControl(ctx context.Context, params *UpdateControlInput, optFns ...func(*Options)) (*UpdateControlOutput, error) 
 UpdateSettings(ctx context.Context, params *UpdateSettingsInput, optFns ...func(*Options)) (*UpdateSettingsOutput, error) 
 ValidateAssessmentReportIntegrity(ctx context.Context, params *ValidateAssessmentReportIntegrityInput, optFns ...func(*Options)) (*ValidateAssessmentReportIntegrityOutput, error) 
}
