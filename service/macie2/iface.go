
package macie2

import (
    "github.com/aws/aws-sdk-go-v2/service/macie2"
)

// IMacie2 defines the interface for macie2
type IMacie2 interface {
 Options() Options 
 AcceptInvitation(ctx context.Context, params *AcceptInvitationInput, optFns ...func(*Options)) (*AcceptInvitationOutput, error) 
 BatchGetCustomDataIdentifiers(ctx context.Context, params *BatchGetCustomDataIdentifiersInput, optFns ...func(*Options)) (*BatchGetCustomDataIdentifiersOutput, error) 
 BatchUpdateAutomatedDiscoveryAccounts(ctx context.Context, params *BatchUpdateAutomatedDiscoveryAccountsInput, optFns ...func(*Options)) (*BatchUpdateAutomatedDiscoveryAccountsOutput, error) 
 CreateAllowList(ctx context.Context, params *CreateAllowListInput, optFns ...func(*Options)) (*CreateAllowListOutput, error) 
 CreateClassificationJob(ctx context.Context, params *CreateClassificationJobInput, optFns ...func(*Options)) (*CreateClassificationJobOutput, error) 
 CreateCustomDataIdentifier(ctx context.Context, params *CreateCustomDataIdentifierInput, optFns ...func(*Options)) (*CreateCustomDataIdentifierOutput, error) 
 CreateFindingsFilter(ctx context.Context, params *CreateFindingsFilterInput, optFns ...func(*Options)) (*CreateFindingsFilterOutput, error) 
 CreateInvitations(ctx context.Context, params *CreateInvitationsInput, optFns ...func(*Options)) (*CreateInvitationsOutput, error) 
 CreateMember(ctx context.Context, params *CreateMemberInput, optFns ...func(*Options)) (*CreateMemberOutput, error) 
 CreateSampleFindings(ctx context.Context, params *CreateSampleFindingsInput, optFns ...func(*Options)) (*CreateSampleFindingsOutput, error) 
 DeclineInvitations(ctx context.Context, params *DeclineInvitationsInput, optFns ...func(*Options)) (*DeclineInvitationsOutput, error) 
 DeleteAllowList(ctx context.Context, params *DeleteAllowListInput, optFns ...func(*Options)) (*DeleteAllowListOutput, error) 
 DeleteCustomDataIdentifier(ctx context.Context, params *DeleteCustomDataIdentifierInput, optFns ...func(*Options)) (*DeleteCustomDataIdentifierOutput, error) 
 DeleteFindingsFilter(ctx context.Context, params *DeleteFindingsFilterInput, optFns ...func(*Options)) (*DeleteFindingsFilterOutput, error) 
 DeleteInvitations(ctx context.Context, params *DeleteInvitationsInput, optFns ...func(*Options)) (*DeleteInvitationsOutput, error) 
 DeleteMember(ctx context.Context, params *DeleteMemberInput, optFns ...func(*Options)) (*DeleteMemberOutput, error) 
 DescribeBuckets(ctx context.Context, params *DescribeBucketsInput, optFns ...func(*Options)) (*DescribeBucketsOutput, error) 
 DescribeClassificationJob(ctx context.Context, params *DescribeClassificationJobInput, optFns ...func(*Options)) (*DescribeClassificationJobOutput, error) 
 DescribeOrganizationConfiguration(ctx context.Context, params *DescribeOrganizationConfigurationInput, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) 
 DisableMacie(ctx context.Context, params *DisableMacieInput, optFns ...func(*Options)) (*DisableMacieOutput, error) 
 DisableOrganizationAdminAccount(ctx context.Context, params *DisableOrganizationAdminAccountInput, optFns ...func(*Options)) (*DisableOrganizationAdminAccountOutput, error) 
 DisassociateFromAdministratorAccount(ctx context.Context, params *DisassociateFromAdministratorAccountInput, optFns ...func(*Options)) (*DisassociateFromAdministratorAccountOutput, error) 
 DisassociateFromMasterAccount(ctx context.Context, params *DisassociateFromMasterAccountInput, optFns ...func(*Options)) (*DisassociateFromMasterAccountOutput, error) 
 DisassociateMember(ctx context.Context, params *DisassociateMemberInput, optFns ...func(*Options)) (*DisassociateMemberOutput, error) 
 EnableMacie(ctx context.Context, params *EnableMacieInput, optFns ...func(*Options)) (*EnableMacieOutput, error) 
 EnableOrganizationAdminAccount(ctx context.Context, params *EnableOrganizationAdminAccountInput, optFns ...func(*Options)) (*EnableOrganizationAdminAccountOutput, error) 
 GetAdministratorAccount(ctx context.Context, params *GetAdministratorAccountInput, optFns ...func(*Options)) (*GetAdministratorAccountOutput, error) 
 GetAllowList(ctx context.Context, params *GetAllowListInput, optFns ...func(*Options)) (*GetAllowListOutput, error) 
 GetAutomatedDiscoveryConfiguration(ctx context.Context, params *GetAutomatedDiscoveryConfigurationInput, optFns ...func(*Options)) (*GetAutomatedDiscoveryConfigurationOutput, error) 
 GetBucketStatistics(ctx context.Context, params *GetBucketStatisticsInput, optFns ...func(*Options)) (*GetBucketStatisticsOutput, error) 
 GetClassificationExportConfiguration(ctx context.Context, params *GetClassificationExportConfigurationInput, optFns ...func(*Options)) (*GetClassificationExportConfigurationOutput, error) 
 GetClassificationScope(ctx context.Context, params *GetClassificationScopeInput, optFns ...func(*Options)) (*GetClassificationScopeOutput, error) 
 GetCustomDataIdentifier(ctx context.Context, params *GetCustomDataIdentifierInput, optFns ...func(*Options)) (*GetCustomDataIdentifierOutput, error) 
 GetFindingStatistics(ctx context.Context, params *GetFindingStatisticsInput, optFns ...func(*Options)) (*GetFindingStatisticsOutput, error) 
 GetFindings(ctx context.Context, params *GetFindingsInput, optFns ...func(*Options)) (*GetFindingsOutput, error) 
 GetFindingsFilter(ctx context.Context, params *GetFindingsFilterInput, optFns ...func(*Options)) (*GetFindingsFilterOutput, error) 
 GetFindingsPublicationConfiguration(ctx context.Context, params *GetFindingsPublicationConfigurationInput, optFns ...func(*Options)) (*GetFindingsPublicationConfigurationOutput, error) 
 GetInvitationsCount(ctx context.Context, params *GetInvitationsCountInput, optFns ...func(*Options)) (*GetInvitationsCountOutput, error) 
 GetMacieSession(ctx context.Context, params *GetMacieSessionInput, optFns ...func(*Options)) (*GetMacieSessionOutput, error) 
 GetMasterAccount(ctx context.Context, params *GetMasterAccountInput, optFns ...func(*Options)) (*GetMasterAccountOutput, error) 
 GetMember(ctx context.Context, params *GetMemberInput, optFns ...func(*Options)) (*GetMemberOutput, error) 
 GetResourceProfile(ctx context.Context, params *GetResourceProfileInput, optFns ...func(*Options)) (*GetResourceProfileOutput, error) 
 GetRevealConfiguration(ctx context.Context, params *GetRevealConfigurationInput, optFns ...func(*Options)) (*GetRevealConfigurationOutput, error) 
 GetSensitiveDataOccurrences(ctx context.Context, params *GetSensitiveDataOccurrencesInput, optFns ...func(*Options)) (*GetSensitiveDataOccurrencesOutput, error) 
 GetSensitiveDataOccurrencesAvailability(ctx context.Context, params *GetSensitiveDataOccurrencesAvailabilityInput, optFns ...func(*Options)) (*GetSensitiveDataOccurrencesAvailabilityOutput, error) 
 GetSensitivityInspectionTemplate(ctx context.Context, params *GetSensitivityInspectionTemplateInput, optFns ...func(*Options)) (*GetSensitivityInspectionTemplateOutput, error) 
 GetUsageStatistics(ctx context.Context, params *GetUsageStatisticsInput, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) 
 GetUsageTotals(ctx context.Context, params *GetUsageTotalsInput, optFns ...func(*Options)) (*GetUsageTotalsOutput, error) 
 ListAllowLists(ctx context.Context, params *ListAllowListsInput, optFns ...func(*Options)) (*ListAllowListsOutput, error) 
 ListAutomatedDiscoveryAccounts(ctx context.Context, params *ListAutomatedDiscoveryAccountsInput, optFns ...func(*Options)) (*ListAutomatedDiscoveryAccountsOutput, error) 
 ListClassificationJobs(ctx context.Context, params *ListClassificationJobsInput, optFns ...func(*Options)) (*ListClassificationJobsOutput, error) 
 ListClassificationScopes(ctx context.Context, params *ListClassificationScopesInput, optFns ...func(*Options)) (*ListClassificationScopesOutput, error) 
 ListCustomDataIdentifiers(ctx context.Context, params *ListCustomDataIdentifiersInput, optFns ...func(*Options)) (*ListCustomDataIdentifiersOutput, error) 
 ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) 
 ListFindingsFilters(ctx context.Context, params *ListFindingsFiltersInput, optFns ...func(*Options)) (*ListFindingsFiltersOutput, error) 
 ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) 
 ListManagedDataIdentifiers(ctx context.Context, params *ListManagedDataIdentifiersInput, optFns ...func(*Options)) (*ListManagedDataIdentifiersOutput, error) 
 ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) 
 ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) 
 ListResourceProfileArtifacts(ctx context.Context, params *ListResourceProfileArtifactsInput, optFns ...func(*Options)) (*ListResourceProfileArtifactsOutput, error) 
 ListResourceProfileDetections(ctx context.Context, params *ListResourceProfileDetectionsInput, optFns ...func(*Options)) (*ListResourceProfileDetectionsOutput, error) 
 ListSensitivityInspectionTemplates(ctx context.Context, params *ListSensitivityInspectionTemplatesInput, optFns ...func(*Options)) (*ListSensitivityInspectionTemplatesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutClassificationExportConfiguration(ctx context.Context, params *PutClassificationExportConfigurationInput, optFns ...func(*Options)) (*PutClassificationExportConfigurationOutput, error) 
 PutFindingsPublicationConfiguration(ctx context.Context, params *PutFindingsPublicationConfigurationInput, optFns ...func(*Options)) (*PutFindingsPublicationConfigurationOutput, error) 
 SearchResources(ctx context.Context, params *SearchResourcesInput, optFns ...func(*Options)) (*SearchResourcesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestCustomDataIdentifier(ctx context.Context, params *TestCustomDataIdentifierInput, optFns ...func(*Options)) (*TestCustomDataIdentifierOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAllowList(ctx context.Context, params *UpdateAllowListInput, optFns ...func(*Options)) (*UpdateAllowListOutput, error) 
 UpdateAutomatedDiscoveryConfiguration(ctx context.Context, params *UpdateAutomatedDiscoveryConfigurationInput, optFns ...func(*Options)) (*UpdateAutomatedDiscoveryConfigurationOutput, error) 
 UpdateClassificationJob(ctx context.Context, params *UpdateClassificationJobInput, optFns ...func(*Options)) (*UpdateClassificationJobOutput, error) 
 UpdateClassificationScope(ctx context.Context, params *UpdateClassificationScopeInput, optFns ...func(*Options)) (*UpdateClassificationScopeOutput, error) 
 UpdateFindingsFilter(ctx context.Context, params *UpdateFindingsFilterInput, optFns ...func(*Options)) (*UpdateFindingsFilterOutput, error) 
 UpdateMacieSession(ctx context.Context, params *UpdateMacieSessionInput, optFns ...func(*Options)) (*UpdateMacieSessionOutput, error) 
 UpdateMemberSession(ctx context.Context, params *UpdateMemberSessionInput, optFns ...func(*Options)) (*UpdateMemberSessionOutput, error) 
 UpdateOrganizationConfiguration(ctx context.Context, params *UpdateOrganizationConfigurationInput, optFns ...func(*Options)) (*UpdateOrganizationConfigurationOutput, error) 
 UpdateResourceProfile(ctx context.Context, params *UpdateResourceProfileInput, optFns ...func(*Options)) (*UpdateResourceProfileOutput, error) 
 UpdateResourceProfileDetections(ctx context.Context, params *UpdateResourceProfileDetectionsInput, optFns ...func(*Options)) (*UpdateResourceProfileDetectionsOutput, error) 
 UpdateRevealConfiguration(ctx context.Context, params *UpdateRevealConfigurationInput, optFns ...func(*Options)) (*UpdateRevealConfigurationOutput, error) 
 UpdateSensitivityInspectionTemplate(ctx context.Context, params *UpdateSensitivityInspectionTemplateInput, optFns ...func(*Options)) (*UpdateSensitivityInspectionTemplateOutput, error) 
}
