
package backup

import (
    "github.com/aws/aws-sdk-go-v2/service/backup"
)

// IBackup defines the interface for backup
type IBackup interface {
 Options() Options 
 CancelLegalHold(ctx context.Context, params *CancelLegalHoldInput, optFns ...func(*Options)) (*CancelLegalHoldOutput, error) 
 CreateBackupPlan(ctx context.Context, params *CreateBackupPlanInput, optFns ...func(*Options)) (*CreateBackupPlanOutput, error) 
 CreateBackupSelection(ctx context.Context, params *CreateBackupSelectionInput, optFns ...func(*Options)) (*CreateBackupSelectionOutput, error) 
 CreateBackupVault(ctx context.Context, params *CreateBackupVaultInput, optFns ...func(*Options)) (*CreateBackupVaultOutput, error) 
 CreateFramework(ctx context.Context, params *CreateFrameworkInput, optFns ...func(*Options)) (*CreateFrameworkOutput, error) 
 CreateLegalHold(ctx context.Context, params *CreateLegalHoldInput, optFns ...func(*Options)) (*CreateLegalHoldOutput, error) 
 CreateLogicallyAirGappedBackupVault(ctx context.Context, params *CreateLogicallyAirGappedBackupVaultInput, optFns ...func(*Options)) (*CreateLogicallyAirGappedBackupVaultOutput, error) 
 CreateReportPlan(ctx context.Context, params *CreateReportPlanInput, optFns ...func(*Options)) (*CreateReportPlanOutput, error) 
 CreateRestoreTestingPlan(ctx context.Context, params *CreateRestoreTestingPlanInput, optFns ...func(*Options)) (*CreateRestoreTestingPlanOutput, error) 
 CreateRestoreTestingSelection(ctx context.Context, params *CreateRestoreTestingSelectionInput, optFns ...func(*Options)) (*CreateRestoreTestingSelectionOutput, error) 
 DeleteBackupPlan(ctx context.Context, params *DeleteBackupPlanInput, optFns ...func(*Options)) (*DeleteBackupPlanOutput, error) 
 DeleteBackupSelection(ctx context.Context, params *DeleteBackupSelectionInput, optFns ...func(*Options)) (*DeleteBackupSelectionOutput, error) 
 DeleteBackupVault(ctx context.Context, params *DeleteBackupVaultInput, optFns ...func(*Options)) (*DeleteBackupVaultOutput, error) 
 DeleteBackupVaultAccessPolicy(ctx context.Context, params *DeleteBackupVaultAccessPolicyInput, optFns ...func(*Options)) (*DeleteBackupVaultAccessPolicyOutput, error) 
 DeleteBackupVaultLockConfiguration(ctx context.Context, params *DeleteBackupVaultLockConfigurationInput, optFns ...func(*Options)) (*DeleteBackupVaultLockConfigurationOutput, error) 
 DeleteBackupVaultNotifications(ctx context.Context, params *DeleteBackupVaultNotificationsInput, optFns ...func(*Options)) (*DeleteBackupVaultNotificationsOutput, error) 
 DeleteFramework(ctx context.Context, params *DeleteFrameworkInput, optFns ...func(*Options)) (*DeleteFrameworkOutput, error) 
 DeleteRecoveryPoint(ctx context.Context, params *DeleteRecoveryPointInput, optFns ...func(*Options)) (*DeleteRecoveryPointOutput, error) 
 DeleteReportPlan(ctx context.Context, params *DeleteReportPlanInput, optFns ...func(*Options)) (*DeleteReportPlanOutput, error) 
 DeleteRestoreTestingPlan(ctx context.Context, params *DeleteRestoreTestingPlanInput, optFns ...func(*Options)) (*DeleteRestoreTestingPlanOutput, error) 
 DeleteRestoreTestingSelection(ctx context.Context, params *DeleteRestoreTestingSelectionInput, optFns ...func(*Options)) (*DeleteRestoreTestingSelectionOutput, error) 
 DescribeBackupJob(ctx context.Context, params *DescribeBackupJobInput, optFns ...func(*Options)) (*DescribeBackupJobOutput, error) 
 DescribeBackupVault(ctx context.Context, params *DescribeBackupVaultInput, optFns ...func(*Options)) (*DescribeBackupVaultOutput, error) 
 DescribeCopyJob(ctx context.Context, params *DescribeCopyJobInput, optFns ...func(*Options)) (*DescribeCopyJobOutput, error) 
 DescribeFramework(ctx context.Context, params *DescribeFrameworkInput, optFns ...func(*Options)) (*DescribeFrameworkOutput, error) 
 DescribeGlobalSettings(ctx context.Context, params *DescribeGlobalSettingsInput, optFns ...func(*Options)) (*DescribeGlobalSettingsOutput, error) 
 DescribeProtectedResource(ctx context.Context, params *DescribeProtectedResourceInput, optFns ...func(*Options)) (*DescribeProtectedResourceOutput, error) 
 DescribeRecoveryPoint(ctx context.Context, params *DescribeRecoveryPointInput, optFns ...func(*Options)) (*DescribeRecoveryPointOutput, error) 
 DescribeRegionSettings(ctx context.Context, params *DescribeRegionSettingsInput, optFns ...func(*Options)) (*DescribeRegionSettingsOutput, error) 
 DescribeReportJob(ctx context.Context, params *DescribeReportJobInput, optFns ...func(*Options)) (*DescribeReportJobOutput, error) 
 DescribeReportPlan(ctx context.Context, params *DescribeReportPlanInput, optFns ...func(*Options)) (*DescribeReportPlanOutput, error) 
 DescribeRestoreJob(ctx context.Context, params *DescribeRestoreJobInput, optFns ...func(*Options)) (*DescribeRestoreJobOutput, error) 
 DisassociateRecoveryPoint(ctx context.Context, params *DisassociateRecoveryPointInput, optFns ...func(*Options)) (*DisassociateRecoveryPointOutput, error) 
 DisassociateRecoveryPointFromParent(ctx context.Context, params *DisassociateRecoveryPointFromParentInput, optFns ...func(*Options)) (*DisassociateRecoveryPointFromParentOutput, error) 
 ExportBackupPlanTemplate(ctx context.Context, params *ExportBackupPlanTemplateInput, optFns ...func(*Options)) (*ExportBackupPlanTemplateOutput, error) 
 GetBackupPlan(ctx context.Context, params *GetBackupPlanInput, optFns ...func(*Options)) (*GetBackupPlanOutput, error) 
 GetBackupPlanFromJSON(ctx context.Context, params *GetBackupPlanFromJSONInput, optFns ...func(*Options)) (*GetBackupPlanFromJSONOutput, error) 
 GetBackupPlanFromTemplate(ctx context.Context, params *GetBackupPlanFromTemplateInput, optFns ...func(*Options)) (*GetBackupPlanFromTemplateOutput, error) 
 GetBackupSelection(ctx context.Context, params *GetBackupSelectionInput, optFns ...func(*Options)) (*GetBackupSelectionOutput, error) 
 GetBackupVaultAccessPolicy(ctx context.Context, params *GetBackupVaultAccessPolicyInput, optFns ...func(*Options)) (*GetBackupVaultAccessPolicyOutput, error) 
 GetBackupVaultNotifications(ctx context.Context, params *GetBackupVaultNotificationsInput, optFns ...func(*Options)) (*GetBackupVaultNotificationsOutput, error) 
 GetLegalHold(ctx context.Context, params *GetLegalHoldInput, optFns ...func(*Options)) (*GetLegalHoldOutput, error) 
 GetRecoveryPointRestoreMetadata(ctx context.Context, params *GetRecoveryPointRestoreMetadataInput, optFns ...func(*Options)) (*GetRecoveryPointRestoreMetadataOutput, error) 
 GetRestoreJobMetadata(ctx context.Context, params *GetRestoreJobMetadataInput, optFns ...func(*Options)) (*GetRestoreJobMetadataOutput, error) 
 GetRestoreTestingInferredMetadata(ctx context.Context, params *GetRestoreTestingInferredMetadataInput, optFns ...func(*Options)) (*GetRestoreTestingInferredMetadataOutput, error) 
 GetRestoreTestingPlan(ctx context.Context, params *GetRestoreTestingPlanInput, optFns ...func(*Options)) (*GetRestoreTestingPlanOutput, error) 
 GetRestoreTestingSelection(ctx context.Context, params *GetRestoreTestingSelectionInput, optFns ...func(*Options)) (*GetRestoreTestingSelectionOutput, error) 
 GetSupportedResourceTypes(ctx context.Context, params *GetSupportedResourceTypesInput, optFns ...func(*Options)) (*GetSupportedResourceTypesOutput, error) 
 ListBackupJobSummaries(ctx context.Context, params *ListBackupJobSummariesInput, optFns ...func(*Options)) (*ListBackupJobSummariesOutput, error) 
 ListBackupJobs(ctx context.Context, params *ListBackupJobsInput, optFns ...func(*Options)) (*ListBackupJobsOutput, error) 
 ListBackupPlanTemplates(ctx context.Context, params *ListBackupPlanTemplatesInput, optFns ...func(*Options)) (*ListBackupPlanTemplatesOutput, error) 
 ListBackupPlanVersions(ctx context.Context, params *ListBackupPlanVersionsInput, optFns ...func(*Options)) (*ListBackupPlanVersionsOutput, error) 
 ListBackupPlans(ctx context.Context, params *ListBackupPlansInput, optFns ...func(*Options)) (*ListBackupPlansOutput, error) 
 ListBackupSelections(ctx context.Context, params *ListBackupSelectionsInput, optFns ...func(*Options)) (*ListBackupSelectionsOutput, error) 
 ListBackupVaults(ctx context.Context, params *ListBackupVaultsInput, optFns ...func(*Options)) (*ListBackupVaultsOutput, error) 
 ListCopyJobSummaries(ctx context.Context, params *ListCopyJobSummariesInput, optFns ...func(*Options)) (*ListCopyJobSummariesOutput, error) 
 ListCopyJobs(ctx context.Context, params *ListCopyJobsInput, optFns ...func(*Options)) (*ListCopyJobsOutput, error) 
 ListFrameworks(ctx context.Context, params *ListFrameworksInput, optFns ...func(*Options)) (*ListFrameworksOutput, error) 
 ListLegalHolds(ctx context.Context, params *ListLegalHoldsInput, optFns ...func(*Options)) (*ListLegalHoldsOutput, error) 
 ListProtectedResources(ctx context.Context, params *ListProtectedResourcesInput, optFns ...func(*Options)) (*ListProtectedResourcesOutput, error) 
 ListProtectedResourcesByBackupVault(ctx context.Context, params *ListProtectedResourcesByBackupVaultInput, optFns ...func(*Options)) (*ListProtectedResourcesByBackupVaultOutput, error) 
 ListRecoveryPointsByBackupVault(ctx context.Context, params *ListRecoveryPointsByBackupVaultInput, optFns ...func(*Options)) (*ListRecoveryPointsByBackupVaultOutput, error) 
 ListRecoveryPointsByLegalHold(ctx context.Context, params *ListRecoveryPointsByLegalHoldInput, optFns ...func(*Options)) (*ListRecoveryPointsByLegalHoldOutput, error) 
 ListRecoveryPointsByResource(ctx context.Context, params *ListRecoveryPointsByResourceInput, optFns ...func(*Options)) (*ListRecoveryPointsByResourceOutput, error) 
 ListReportJobs(ctx context.Context, params *ListReportJobsInput, optFns ...func(*Options)) (*ListReportJobsOutput, error) 
 ListReportPlans(ctx context.Context, params *ListReportPlansInput, optFns ...func(*Options)) (*ListReportPlansOutput, error) 
 ListRestoreJobSummaries(ctx context.Context, params *ListRestoreJobSummariesInput, optFns ...func(*Options)) (*ListRestoreJobSummariesOutput, error) 
 ListRestoreJobs(ctx context.Context, params *ListRestoreJobsInput, optFns ...func(*Options)) (*ListRestoreJobsOutput, error) 
 ListRestoreJobsByProtectedResource(ctx context.Context, params *ListRestoreJobsByProtectedResourceInput, optFns ...func(*Options)) (*ListRestoreJobsByProtectedResourceOutput, error) 
 ListRestoreTestingPlans(ctx context.Context, params *ListRestoreTestingPlansInput, optFns ...func(*Options)) (*ListRestoreTestingPlansOutput, error) 
 ListRestoreTestingSelections(ctx context.Context, params *ListRestoreTestingSelectionsInput, optFns ...func(*Options)) (*ListRestoreTestingSelectionsOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 PutBackupVaultAccessPolicy(ctx context.Context, params *PutBackupVaultAccessPolicyInput, optFns ...func(*Options)) (*PutBackupVaultAccessPolicyOutput, error) 
 PutBackupVaultLockConfiguration(ctx context.Context, params *PutBackupVaultLockConfigurationInput, optFns ...func(*Options)) (*PutBackupVaultLockConfigurationOutput, error) 
 PutBackupVaultNotifications(ctx context.Context, params *PutBackupVaultNotificationsInput, optFns ...func(*Options)) (*PutBackupVaultNotificationsOutput, error) 
 PutRestoreValidationResult(ctx context.Context, params *PutRestoreValidationResultInput, optFns ...func(*Options)) (*PutRestoreValidationResultOutput, error) 
 StartBackupJob(ctx context.Context, params *StartBackupJobInput, optFns ...func(*Options)) (*StartBackupJobOutput, error) 
 StartCopyJob(ctx context.Context, params *StartCopyJobInput, optFns ...func(*Options)) (*StartCopyJobOutput, error) 
 StartReportJob(ctx context.Context, params *StartReportJobInput, optFns ...func(*Options)) (*StartReportJobOutput, error) 
 StartRestoreJob(ctx context.Context, params *StartRestoreJobInput, optFns ...func(*Options)) (*StartRestoreJobOutput, error) 
 StopBackupJob(ctx context.Context, params *StopBackupJobInput, optFns ...func(*Options)) (*StopBackupJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBackupPlan(ctx context.Context, params *UpdateBackupPlanInput, optFns ...func(*Options)) (*UpdateBackupPlanOutput, error) 
 UpdateFramework(ctx context.Context, params *UpdateFrameworkInput, optFns ...func(*Options)) (*UpdateFrameworkOutput, error) 
 UpdateGlobalSettings(ctx context.Context, params *UpdateGlobalSettingsInput, optFns ...func(*Options)) (*UpdateGlobalSettingsOutput, error) 
 UpdateRecoveryPointLifecycle(ctx context.Context, params *UpdateRecoveryPointLifecycleInput, optFns ...func(*Options)) (*UpdateRecoveryPointLifecycleOutput, error) 
 UpdateRegionSettings(ctx context.Context, params *UpdateRegionSettingsInput, optFns ...func(*Options)) (*UpdateRegionSettingsOutput, error) 
 UpdateReportPlan(ctx context.Context, params *UpdateReportPlanInput, optFns ...func(*Options)) (*UpdateReportPlanOutput, error) 
 UpdateRestoreTestingPlan(ctx context.Context, params *UpdateRestoreTestingPlanInput, optFns ...func(*Options)) (*UpdateRestoreTestingPlanOutput, error) 
 UpdateRestoreTestingSelection(ctx context.Context, params *UpdateRestoreTestingSelectionInput, optFns ...func(*Options)) (*UpdateRestoreTestingSelectionOutput, error) 
}
