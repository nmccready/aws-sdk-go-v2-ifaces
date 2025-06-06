
package inspector2_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/inspector2"
)

// IClient defines the interface for inspector2
type IClient interface {
 Options() Options 
 AssociateMember(ctx context.Context, params *AssociateMemberInput, optFns ...func(*Options)) (*AssociateMemberOutput, error) 
 BatchGetAccountStatus(ctx context.Context, params *BatchGetAccountStatusInput, optFns ...func(*Options)) (*BatchGetAccountStatusOutput, error) 
 BatchGetCodeSnippet(ctx context.Context, params *BatchGetCodeSnippetInput, optFns ...func(*Options)) (*BatchGetCodeSnippetOutput, error) 
 BatchGetFindingDetails(ctx context.Context, params *BatchGetFindingDetailsInput, optFns ...func(*Options)) (*BatchGetFindingDetailsOutput, error) 
 BatchGetFreeTrialInfo(ctx context.Context, params *BatchGetFreeTrialInfoInput, optFns ...func(*Options)) (*BatchGetFreeTrialInfoOutput, error) 
 BatchGetMemberEc2DeepInspectionStatus(ctx context.Context, params *BatchGetMemberEc2DeepInspectionStatusInput, optFns ...func(*Options)) (*BatchGetMemberEc2DeepInspectionStatusOutput, error) 
 BatchUpdateMemberEc2DeepInspectionStatus(ctx context.Context, params *BatchUpdateMemberEc2DeepInspectionStatusInput, optFns ...func(*Options)) (*BatchUpdateMemberEc2DeepInspectionStatusOutput, error) 
 CancelFindingsReport(ctx context.Context, params *CancelFindingsReportInput, optFns ...func(*Options)) (*CancelFindingsReportOutput, error) 
 CancelSbomExport(ctx context.Context, params *CancelSbomExportInput, optFns ...func(*Options)) (*CancelSbomExportOutput, error) 
 CreateCisScanConfiguration(ctx context.Context, params *CreateCisScanConfigurationInput, optFns ...func(*Options)) (*CreateCisScanConfigurationOutput, error) 
 CreateFilter(ctx context.Context, params *CreateFilterInput, optFns ...func(*Options)) (*CreateFilterOutput, error) 
 CreateFindingsReport(ctx context.Context, params *CreateFindingsReportInput, optFns ...func(*Options)) (*CreateFindingsReportOutput, error) 
 CreateSbomExport(ctx context.Context, params *CreateSbomExportInput, optFns ...func(*Options)) (*CreateSbomExportOutput, error) 
 DeleteCisScanConfiguration(ctx context.Context, params *DeleteCisScanConfigurationInput, optFns ...func(*Options)) (*DeleteCisScanConfigurationOutput, error) 
 DeleteFilter(ctx context.Context, params *DeleteFilterInput, optFns ...func(*Options)) (*DeleteFilterOutput, error) 
 DescribeOrganizationConfiguration(ctx context.Context, params *DescribeOrganizationConfigurationInput, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) 
 Disable(ctx context.Context, params *DisableInput, optFns ...func(*Options)) (*DisableOutput, error) 
 DisableDelegatedAdminAccount(ctx context.Context, params *DisableDelegatedAdminAccountInput, optFns ...func(*Options)) (*DisableDelegatedAdminAccountOutput, error) 
 DisassociateMember(ctx context.Context, params *DisassociateMemberInput, optFns ...func(*Options)) (*DisassociateMemberOutput, error) 
 Enable(ctx context.Context, params *EnableInput, optFns ...func(*Options)) (*EnableOutput, error) 
 EnableDelegatedAdminAccount(ctx context.Context, params *EnableDelegatedAdminAccountInput, optFns ...func(*Options)) (*EnableDelegatedAdminAccountOutput, error) 
 GetCisScanReport(ctx context.Context, params *GetCisScanReportInput, optFns ...func(*Options)) (*GetCisScanReportOutput, error) 
 GetCisScanResultDetails(ctx context.Context, params *GetCisScanResultDetailsInput, optFns ...func(*Options)) (*GetCisScanResultDetailsOutput, error) 
 GetClustersForImage(ctx context.Context, params *GetClustersForImageInput, optFns ...func(*Options)) (*GetClustersForImageOutput, error) 
 GetConfiguration(ctx context.Context, params *GetConfigurationInput, optFns ...func(*Options)) (*GetConfigurationOutput, error) 
 GetDelegatedAdminAccount(ctx context.Context, params *GetDelegatedAdminAccountInput, optFns ...func(*Options)) (*GetDelegatedAdminAccountOutput, error) 
 GetEc2DeepInspectionConfiguration(ctx context.Context, params *GetEc2DeepInspectionConfigurationInput, optFns ...func(*Options)) (*GetEc2DeepInspectionConfigurationOutput, error) 
 GetEncryptionKey(ctx context.Context, params *GetEncryptionKeyInput, optFns ...func(*Options)) (*GetEncryptionKeyOutput, error) 
 GetFindingsReportStatus(ctx context.Context, params *GetFindingsReportStatusInput, optFns ...func(*Options)) (*GetFindingsReportStatusOutput, error) 
 GetMember(ctx context.Context, params *GetMemberInput, optFns ...func(*Options)) (*GetMemberOutput, error) 
 GetSbomExport(ctx context.Context, params *GetSbomExportInput, optFns ...func(*Options)) (*GetSbomExportOutput, error) 
 ListAccountPermissions(ctx context.Context, params *ListAccountPermissionsInput, optFns ...func(*Options)) (*ListAccountPermissionsOutput, error) 
 ListCisScanConfigurations(ctx context.Context, params *ListCisScanConfigurationsInput, optFns ...func(*Options)) (*ListCisScanConfigurationsOutput, error) 
 ListCisScanResultsAggregatedByChecks(ctx context.Context, params *ListCisScanResultsAggregatedByChecksInput, optFns ...func(*Options)) (*ListCisScanResultsAggregatedByChecksOutput, error) 
 ListCisScanResultsAggregatedByTargetResource(ctx context.Context, params *ListCisScanResultsAggregatedByTargetResourceInput, optFns ...func(*Options)) (*ListCisScanResultsAggregatedByTargetResourceOutput, error) 
 ListCisScans(ctx context.Context, params *ListCisScansInput, optFns ...func(*Options)) (*ListCisScansOutput, error) 
 ListCoverage(ctx context.Context, params *ListCoverageInput, optFns ...func(*Options)) (*ListCoverageOutput, error) 
 ListCoverageStatistics(ctx context.Context, params *ListCoverageStatisticsInput, optFns ...func(*Options)) (*ListCoverageStatisticsOutput, error) 
 ListDelegatedAdminAccounts(ctx context.Context, params *ListDelegatedAdminAccountsInput, optFns ...func(*Options)) (*ListDelegatedAdminAccountsOutput, error) 
 ListFilters(ctx context.Context, params *ListFiltersInput, optFns ...func(*Options)) (*ListFiltersOutput, error) 
 ListFindingAggregations(ctx context.Context, params *ListFindingAggregationsInput, optFns ...func(*Options)) (*ListFindingAggregationsOutput, error) 
 ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) 
 ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListUsageTotals(ctx context.Context, params *ListUsageTotalsInput, optFns ...func(*Options)) (*ListUsageTotalsOutput, error) 
 ResetEncryptionKey(ctx context.Context, params *ResetEncryptionKeyInput, optFns ...func(*Options)) (*ResetEncryptionKeyOutput, error) 
 SearchVulnerabilities(ctx context.Context, params *SearchVulnerabilitiesInput, optFns ...func(*Options)) (*SearchVulnerabilitiesOutput, error) 
 SendCisSessionHealth(ctx context.Context, params *SendCisSessionHealthInput, optFns ...func(*Options)) (*SendCisSessionHealthOutput, error) 
 SendCisSessionTelemetry(ctx context.Context, params *SendCisSessionTelemetryInput, optFns ...func(*Options)) (*SendCisSessionTelemetryOutput, error) 
 StartCisSession(ctx context.Context, params *StartCisSessionInput, optFns ...func(*Options)) (*StartCisSessionOutput, error) 
 StopCisSession(ctx context.Context, params *StopCisSessionInput, optFns ...func(*Options)) (*StopCisSessionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCisScanConfiguration(ctx context.Context, params *UpdateCisScanConfigurationInput, optFns ...func(*Options)) (*UpdateCisScanConfigurationOutput, error) 
 UpdateConfiguration(ctx context.Context, params *UpdateConfigurationInput, optFns ...func(*Options)) (*UpdateConfigurationOutput, error) 
 UpdateEc2DeepInspectionConfiguration(ctx context.Context, params *UpdateEc2DeepInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateEc2DeepInspectionConfigurationOutput, error) 
 UpdateEncryptionKey(ctx context.Context, params *UpdateEncryptionKeyInput, optFns ...func(*Options)) (*UpdateEncryptionKeyOutput, error) 
 UpdateFilter(ctx context.Context, params *UpdateFilterInput, optFns ...func(*Options)) (*UpdateFilterOutput, error) 
 UpdateOrgEc2DeepInspectionConfiguration(ctx context.Context, params *UpdateOrgEc2DeepInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateOrgEc2DeepInspectionConfigurationOutput, error) 
 UpdateOrganizationConfiguration(ctx context.Context, params *UpdateOrganizationConfigurationInput, optFns ...func(*Options)) (*UpdateOrganizationConfigurationOutput, error) 
}
