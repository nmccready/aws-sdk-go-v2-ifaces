
package quicksight_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/quicksight"
)

// IClient defines the interface for quicksight
type IClient interface {
 Options() Options 
 CancelIngestion(ctx context.Context, params *CancelIngestionInput, optFns ...func(*Options)) (*CancelIngestionOutput, error) 
 CreateAccountCustomization(ctx context.Context, params *CreateAccountCustomizationInput, optFns ...func(*Options)) (*CreateAccountCustomizationOutput, error) 
 CreateAccountSubscription(ctx context.Context, params *CreateAccountSubscriptionInput, optFns ...func(*Options)) (*CreateAccountSubscriptionOutput, error) 
 CreateAnalysis(ctx context.Context, params *CreateAnalysisInput, optFns ...func(*Options)) (*CreateAnalysisOutput, error) 
 CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) 
 CreateDataSet(ctx context.Context, params *CreateDataSetInput, optFns ...func(*Options)) (*CreateDataSetOutput, error) 
 CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) 
 CreateFolder(ctx context.Context, params *CreateFolderInput, optFns ...func(*Options)) (*CreateFolderOutput, error) 
 CreateFolderMembership(ctx context.Context, params *CreateFolderMembershipInput, optFns ...func(*Options)) (*CreateFolderMembershipOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateGroupMembership(ctx context.Context, params *CreateGroupMembershipInput, optFns ...func(*Options)) (*CreateGroupMembershipOutput, error) 
 CreateIAMPolicyAssignment(ctx context.Context, params *CreateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*CreateIAMPolicyAssignmentOutput, error) 
 CreateIngestion(ctx context.Context, params *CreateIngestionInput, optFns ...func(*Options)) (*CreateIngestionOutput, error) 
 CreateNamespace(ctx context.Context, params *CreateNamespaceInput, optFns ...func(*Options)) (*CreateNamespaceOutput, error) 
 CreateRefreshSchedule(ctx context.Context, params *CreateRefreshScheduleInput, optFns ...func(*Options)) (*CreateRefreshScheduleOutput, error) 
 CreateRoleMembership(ctx context.Context, params *CreateRoleMembershipInput, optFns ...func(*Options)) (*CreateRoleMembershipOutput, error) 
 CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) 
 CreateTemplateAlias(ctx context.Context, params *CreateTemplateAliasInput, optFns ...func(*Options)) (*CreateTemplateAliasOutput, error) 
 CreateTheme(ctx context.Context, params *CreateThemeInput, optFns ...func(*Options)) (*CreateThemeOutput, error) 
 CreateThemeAlias(ctx context.Context, params *CreateThemeAliasInput, optFns ...func(*Options)) (*CreateThemeAliasOutput, error) 
 CreateTopic(ctx context.Context, params *CreateTopicInput, optFns ...func(*Options)) (*CreateTopicOutput, error) 
 CreateTopicRefreshSchedule(ctx context.Context, params *CreateTopicRefreshScheduleInput, optFns ...func(*Options)) (*CreateTopicRefreshScheduleOutput, error) 
 CreateVPCConnection(ctx context.Context, params *CreateVPCConnectionInput, optFns ...func(*Options)) (*CreateVPCConnectionOutput, error) 
 DeleteAccountCustomization(ctx context.Context, params *DeleteAccountCustomizationInput, optFns ...func(*Options)) (*DeleteAccountCustomizationOutput, error) 
 DeleteAccountSubscription(ctx context.Context, params *DeleteAccountSubscriptionInput, optFns ...func(*Options)) (*DeleteAccountSubscriptionOutput, error) 
 DeleteAnalysis(ctx context.Context, params *DeleteAnalysisInput, optFns ...func(*Options)) (*DeleteAnalysisOutput, error) 
 DeleteDashboard(ctx context.Context, params *DeleteDashboardInput, optFns ...func(*Options)) (*DeleteDashboardOutput, error) 
 DeleteDataSet(ctx context.Context, params *DeleteDataSetInput, optFns ...func(*Options)) (*DeleteDataSetOutput, error) 
 DeleteDataSetRefreshProperties(ctx context.Context, params *DeleteDataSetRefreshPropertiesInput, optFns ...func(*Options)) (*DeleteDataSetRefreshPropertiesOutput, error) 
 DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) 
 DeleteFolder(ctx context.Context, params *DeleteFolderInput, optFns ...func(*Options)) (*DeleteFolderOutput, error) 
 DeleteFolderMembership(ctx context.Context, params *DeleteFolderMembershipInput, optFns ...func(*Options)) (*DeleteFolderMembershipOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteGroupMembership(ctx context.Context, params *DeleteGroupMembershipInput, optFns ...func(*Options)) (*DeleteGroupMembershipOutput, error) 
 DeleteIAMPolicyAssignment(ctx context.Context, params *DeleteIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DeleteIAMPolicyAssignmentOutput, error) 
 DeleteIdentityPropagationConfig(ctx context.Context, params *DeleteIdentityPropagationConfigInput, optFns ...func(*Options)) (*DeleteIdentityPropagationConfigOutput, error) 
 DeleteNamespace(ctx context.Context, params *DeleteNamespaceInput, optFns ...func(*Options)) (*DeleteNamespaceOutput, error) 
 DeleteRefreshSchedule(ctx context.Context, params *DeleteRefreshScheduleInput, optFns ...func(*Options)) (*DeleteRefreshScheduleOutput, error) 
 DeleteRoleCustomPermission(ctx context.Context, params *DeleteRoleCustomPermissionInput, optFns ...func(*Options)) (*DeleteRoleCustomPermissionOutput, error) 
 DeleteRoleMembership(ctx context.Context, params *DeleteRoleMembershipInput, optFns ...func(*Options)) (*DeleteRoleMembershipOutput, error) 
 DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) 
 DeleteTemplateAlias(ctx context.Context, params *DeleteTemplateAliasInput, optFns ...func(*Options)) (*DeleteTemplateAliasOutput, error) 
 DeleteTheme(ctx context.Context, params *DeleteThemeInput, optFns ...func(*Options)) (*DeleteThemeOutput, error) 
 DeleteThemeAlias(ctx context.Context, params *DeleteThemeAliasInput, optFns ...func(*Options)) (*DeleteThemeAliasOutput, error) 
 DeleteTopic(ctx context.Context, params *DeleteTopicInput, optFns ...func(*Options)) (*DeleteTopicOutput, error) 
 DeleteTopicRefreshSchedule(ctx context.Context, params *DeleteTopicRefreshScheduleInput, optFns ...func(*Options)) (*DeleteTopicRefreshScheduleOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DeleteUserByPrincipalId(ctx context.Context, params *DeleteUserByPrincipalIdInput, optFns ...func(*Options)) (*DeleteUserByPrincipalIdOutput, error) 
 DeleteVPCConnection(ctx context.Context, params *DeleteVPCConnectionInput, optFns ...func(*Options)) (*DeleteVPCConnectionOutput, error) 
 DescribeAccountCustomization(ctx context.Context, params *DescribeAccountCustomizationInput, optFns ...func(*Options)) (*DescribeAccountCustomizationOutput, error) 
 DescribeAccountSettings(ctx context.Context, params *DescribeAccountSettingsInput, optFns ...func(*Options)) (*DescribeAccountSettingsOutput, error) 
 DescribeAccountSubscription(ctx context.Context, params *DescribeAccountSubscriptionInput, optFns ...func(*Options)) (*DescribeAccountSubscriptionOutput, error) 
 DescribeAnalysis(ctx context.Context, params *DescribeAnalysisInput, optFns ...func(*Options)) (*DescribeAnalysisOutput, error) 
 DescribeAnalysisDefinition(ctx context.Context, params *DescribeAnalysisDefinitionInput, optFns ...func(*Options)) (*DescribeAnalysisDefinitionOutput, error) 
 DescribeAnalysisPermissions(ctx context.Context, params *DescribeAnalysisPermissionsInput, optFns ...func(*Options)) (*DescribeAnalysisPermissionsOutput, error) 
 DescribeAssetBundleExportJob(ctx context.Context, params *DescribeAssetBundleExportJobInput, optFns ...func(*Options)) (*DescribeAssetBundleExportJobOutput, error) 
 DescribeAssetBundleImportJob(ctx context.Context, params *DescribeAssetBundleImportJobInput, optFns ...func(*Options)) (*DescribeAssetBundleImportJobOutput, error) 
 DescribeDashboard(ctx context.Context, params *DescribeDashboardInput, optFns ...func(*Options)) (*DescribeDashboardOutput, error) 
 DescribeDashboardDefinition(ctx context.Context, params *DescribeDashboardDefinitionInput, optFns ...func(*Options)) (*DescribeDashboardDefinitionOutput, error) 
 DescribeDashboardPermissions(ctx context.Context, params *DescribeDashboardPermissionsInput, optFns ...func(*Options)) (*DescribeDashboardPermissionsOutput, error) 
 DescribeDashboardSnapshotJob(ctx context.Context, params *DescribeDashboardSnapshotJobInput, optFns ...func(*Options)) (*DescribeDashboardSnapshotJobOutput, error) 
 DescribeDashboardSnapshotJobResult(ctx context.Context, params *DescribeDashboardSnapshotJobResultInput, optFns ...func(*Options)) (*DescribeDashboardSnapshotJobResultOutput, error) 
 DescribeDataSet(ctx context.Context, params *DescribeDataSetInput, optFns ...func(*Options)) (*DescribeDataSetOutput, error) 
 DescribeDataSetPermissions(ctx context.Context, params *DescribeDataSetPermissionsInput, optFns ...func(*Options)) (*DescribeDataSetPermissionsOutput, error) 
 DescribeDataSetRefreshProperties(ctx context.Context, params *DescribeDataSetRefreshPropertiesInput, optFns ...func(*Options)) (*DescribeDataSetRefreshPropertiesOutput, error) 
 DescribeDataSource(ctx context.Context, params *DescribeDataSourceInput, optFns ...func(*Options)) (*DescribeDataSourceOutput, error) 
 DescribeDataSourcePermissions(ctx context.Context, params *DescribeDataSourcePermissionsInput, optFns ...func(*Options)) (*DescribeDataSourcePermissionsOutput, error) 
 DescribeFolder(ctx context.Context, params *DescribeFolderInput, optFns ...func(*Options)) (*DescribeFolderOutput, error) 
 DescribeFolderPermissions(ctx context.Context, params *DescribeFolderPermissionsInput, optFns ...func(*Options)) (*DescribeFolderPermissionsOutput, error) 
 DescribeFolderResolvedPermissions(ctx context.Context, params *DescribeFolderResolvedPermissionsInput, optFns ...func(*Options)) (*DescribeFolderResolvedPermissionsOutput, error) 
 DescribeGroup(ctx context.Context, params *DescribeGroupInput, optFns ...func(*Options)) (*DescribeGroupOutput, error) 
 DescribeGroupMembership(ctx context.Context, params *DescribeGroupMembershipInput, optFns ...func(*Options)) (*DescribeGroupMembershipOutput, error) 
 DescribeIAMPolicyAssignment(ctx context.Context, params *DescribeIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DescribeIAMPolicyAssignmentOutput, error) 
 DescribeIngestion(ctx context.Context, params *DescribeIngestionInput, optFns ...func(*Options)) (*DescribeIngestionOutput, error) 
 DescribeIpRestriction(ctx context.Context, params *DescribeIpRestrictionInput, optFns ...func(*Options)) (*DescribeIpRestrictionOutput, error) 
 DescribeKeyRegistration(ctx context.Context, params *DescribeKeyRegistrationInput, optFns ...func(*Options)) (*DescribeKeyRegistrationOutput, error) 
 DescribeNamespace(ctx context.Context, params *DescribeNamespaceInput, optFns ...func(*Options)) (*DescribeNamespaceOutput, error) 
 DescribeRefreshSchedule(ctx context.Context, params *DescribeRefreshScheduleInput, optFns ...func(*Options)) (*DescribeRefreshScheduleOutput, error) 
 DescribeRoleCustomPermission(ctx context.Context, params *DescribeRoleCustomPermissionInput, optFns ...func(*Options)) (*DescribeRoleCustomPermissionOutput, error) 
 DescribeTemplate(ctx context.Context, params *DescribeTemplateInput, optFns ...func(*Options)) (*DescribeTemplateOutput, error) 
 DescribeTemplateAlias(ctx context.Context, params *DescribeTemplateAliasInput, optFns ...func(*Options)) (*DescribeTemplateAliasOutput, error) 
 DescribeTemplateDefinition(ctx context.Context, params *DescribeTemplateDefinitionInput, optFns ...func(*Options)) (*DescribeTemplateDefinitionOutput, error) 
 DescribeTemplatePermissions(ctx context.Context, params *DescribeTemplatePermissionsInput, optFns ...func(*Options)) (*DescribeTemplatePermissionsOutput, error) 
 DescribeTheme(ctx context.Context, params *DescribeThemeInput, optFns ...func(*Options)) (*DescribeThemeOutput, error) 
 DescribeThemeAlias(ctx context.Context, params *DescribeThemeAliasInput, optFns ...func(*Options)) (*DescribeThemeAliasOutput, error) 
 DescribeThemePermissions(ctx context.Context, params *DescribeThemePermissionsInput, optFns ...func(*Options)) (*DescribeThemePermissionsOutput, error) 
 DescribeTopic(ctx context.Context, params *DescribeTopicInput, optFns ...func(*Options)) (*DescribeTopicOutput, error) 
 DescribeTopicPermissions(ctx context.Context, params *DescribeTopicPermissionsInput, optFns ...func(*Options)) (*DescribeTopicPermissionsOutput, error) 
 DescribeTopicRefresh(ctx context.Context, params *DescribeTopicRefreshInput, optFns ...func(*Options)) (*DescribeTopicRefreshOutput, error) 
 DescribeTopicRefreshSchedule(ctx context.Context, params *DescribeTopicRefreshScheduleInput, optFns ...func(*Options)) (*DescribeTopicRefreshScheduleOutput, error) 
 DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) 
 DescribeVPCConnection(ctx context.Context, params *DescribeVPCConnectionInput, optFns ...func(*Options)) (*DescribeVPCConnectionOutput, error) 
 GenerateEmbedUrlForAnonymousUser(ctx context.Context, params *GenerateEmbedUrlForAnonymousUserInput, optFns ...func(*Options)) (*GenerateEmbedUrlForAnonymousUserOutput, error) 
 GenerateEmbedUrlForRegisteredUser(ctx context.Context, params *GenerateEmbedUrlForRegisteredUserInput, optFns ...func(*Options)) (*GenerateEmbedUrlForRegisteredUserOutput, error) 
 GetDashboardEmbedUrl(ctx context.Context, params *GetDashboardEmbedUrlInput, optFns ...func(*Options)) (*GetDashboardEmbedUrlOutput, error) 
 GetSessionEmbedUrl(ctx context.Context, params *GetSessionEmbedUrlInput, optFns ...func(*Options)) (*GetSessionEmbedUrlOutput, error) 
 ListAnalyses(ctx context.Context, params *ListAnalysesInput, optFns ...func(*Options)) (*ListAnalysesOutput, error) 
 ListAssetBundleExportJobs(ctx context.Context, params *ListAssetBundleExportJobsInput, optFns ...func(*Options)) (*ListAssetBundleExportJobsOutput, error) 
 ListAssetBundleImportJobs(ctx context.Context, params *ListAssetBundleImportJobsInput, optFns ...func(*Options)) (*ListAssetBundleImportJobsOutput, error) 
 ListDashboardVersions(ctx context.Context, params *ListDashboardVersionsInput, optFns ...func(*Options)) (*ListDashboardVersionsOutput, error) 
 ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) 
 ListDataSets(ctx context.Context, params *ListDataSetsInput, optFns ...func(*Options)) (*ListDataSetsOutput, error) 
 ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) 
 ListFolderMembers(ctx context.Context, params *ListFolderMembersInput, optFns ...func(*Options)) (*ListFolderMembersOutput, error) 
 ListFolders(ctx context.Context, params *ListFoldersInput, optFns ...func(*Options)) (*ListFoldersOutput, error) 
 ListGroupMemberships(ctx context.Context, params *ListGroupMembershipsInput, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListIAMPolicyAssignments(ctx context.Context, params *ListIAMPolicyAssignmentsInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsOutput, error) 
 ListIAMPolicyAssignmentsForUser(ctx context.Context, params *ListIAMPolicyAssignmentsForUserInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsForUserOutput, error) 
 ListIdentityPropagationConfigs(ctx context.Context, params *ListIdentityPropagationConfigsInput, optFns ...func(*Options)) (*ListIdentityPropagationConfigsOutput, error) 
 ListIngestions(ctx context.Context, params *ListIngestionsInput, optFns ...func(*Options)) (*ListIngestionsOutput, error) 
 ListNamespaces(ctx context.Context, params *ListNamespacesInput, optFns ...func(*Options)) (*ListNamespacesOutput, error) 
 ListRefreshSchedules(ctx context.Context, params *ListRefreshSchedulesInput, optFns ...func(*Options)) (*ListRefreshSchedulesOutput, error) 
 ListRoleMemberships(ctx context.Context, params *ListRoleMembershipsInput, optFns ...func(*Options)) (*ListRoleMembershipsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplateAliases(ctx context.Context, params *ListTemplateAliasesInput, optFns ...func(*Options)) (*ListTemplateAliasesOutput, error) 
 ListTemplateVersions(ctx context.Context, params *ListTemplateVersionsInput, optFns ...func(*Options)) (*ListTemplateVersionsOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 ListThemeAliases(ctx context.Context, params *ListThemeAliasesInput, optFns ...func(*Options)) (*ListThemeAliasesOutput, error) 
 ListThemeVersions(ctx context.Context, params *ListThemeVersionsInput, optFns ...func(*Options)) (*ListThemeVersionsOutput, error) 
 ListThemes(ctx context.Context, params *ListThemesInput, optFns ...func(*Options)) (*ListThemesOutput, error) 
 ListTopicRefreshSchedules(ctx context.Context, params *ListTopicRefreshSchedulesInput, optFns ...func(*Options)) (*ListTopicRefreshSchedulesOutput, error) 
 ListTopics(ctx context.Context, params *ListTopicsInput, optFns ...func(*Options)) (*ListTopicsOutput, error) 
 ListUserGroups(ctx context.Context, params *ListUserGroupsInput, optFns ...func(*Options)) (*ListUserGroupsOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListVPCConnections(ctx context.Context, params *ListVPCConnectionsInput, optFns ...func(*Options)) (*ListVPCConnectionsOutput, error) 
 PutDataSetRefreshProperties(ctx context.Context, params *PutDataSetRefreshPropertiesInput, optFns ...func(*Options)) (*PutDataSetRefreshPropertiesOutput, error) 
 RegisterUser(ctx context.Context, params *RegisterUserInput, optFns ...func(*Options)) (*RegisterUserOutput, error) 
 RestoreAnalysis(ctx context.Context, params *RestoreAnalysisInput, optFns ...func(*Options)) (*RestoreAnalysisOutput, error) 
 SearchAnalyses(ctx context.Context, params *SearchAnalysesInput, optFns ...func(*Options)) (*SearchAnalysesOutput, error) 
 SearchDashboards(ctx context.Context, params *SearchDashboardsInput, optFns ...func(*Options)) (*SearchDashboardsOutput, error) 
 SearchDataSets(ctx context.Context, params *SearchDataSetsInput, optFns ...func(*Options)) (*SearchDataSetsOutput, error) 
 SearchDataSources(ctx context.Context, params *SearchDataSourcesInput, optFns ...func(*Options)) (*SearchDataSourcesOutput, error) 
 SearchFolders(ctx context.Context, params *SearchFoldersInput, optFns ...func(*Options)) (*SearchFoldersOutput, error) 
 SearchGroups(ctx context.Context, params *SearchGroupsInput, optFns ...func(*Options)) (*SearchGroupsOutput, error) 
 StartAssetBundleExportJob(ctx context.Context, params *StartAssetBundleExportJobInput, optFns ...func(*Options)) (*StartAssetBundleExportJobOutput, error) 
 StartAssetBundleImportJob(ctx context.Context, params *StartAssetBundleImportJobInput, optFns ...func(*Options)) (*StartAssetBundleImportJobOutput, error) 
 StartDashboardSnapshotJob(ctx context.Context, params *StartDashboardSnapshotJobInput, optFns ...func(*Options)) (*StartDashboardSnapshotJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccountCustomization(ctx context.Context, params *UpdateAccountCustomizationInput, optFns ...func(*Options)) (*UpdateAccountCustomizationOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateAnalysis(ctx context.Context, params *UpdateAnalysisInput, optFns ...func(*Options)) (*UpdateAnalysisOutput, error) 
 UpdateAnalysisPermissions(ctx context.Context, params *UpdateAnalysisPermissionsInput, optFns ...func(*Options)) (*UpdateAnalysisPermissionsOutput, error) 
 UpdateDashboard(ctx context.Context, params *UpdateDashboardInput, optFns ...func(*Options)) (*UpdateDashboardOutput, error) 
 UpdateDashboardLinks(ctx context.Context, params *UpdateDashboardLinksInput, optFns ...func(*Options)) (*UpdateDashboardLinksOutput, error) 
 UpdateDashboardPermissions(ctx context.Context, params *UpdateDashboardPermissionsInput, optFns ...func(*Options)) (*UpdateDashboardPermissionsOutput, error) 
 UpdateDashboardPublishedVersion(ctx context.Context, params *UpdateDashboardPublishedVersionInput, optFns ...func(*Options)) (*UpdateDashboardPublishedVersionOutput, error) 
 UpdateDataSet(ctx context.Context, params *UpdateDataSetInput, optFns ...func(*Options)) (*UpdateDataSetOutput, error) 
 UpdateDataSetPermissions(ctx context.Context, params *UpdateDataSetPermissionsInput, optFns ...func(*Options)) (*UpdateDataSetPermissionsOutput, error) 
 UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) 
 UpdateDataSourcePermissions(ctx context.Context, params *UpdateDataSourcePermissionsInput, optFns ...func(*Options)) (*UpdateDataSourcePermissionsOutput, error) 
 UpdateFolder(ctx context.Context, params *UpdateFolderInput, optFns ...func(*Options)) (*UpdateFolderOutput, error) 
 UpdateFolderPermissions(ctx context.Context, params *UpdateFolderPermissionsInput, optFns ...func(*Options)) (*UpdateFolderPermissionsOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateIAMPolicyAssignment(ctx context.Context, params *UpdateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*UpdateIAMPolicyAssignmentOutput, error) 
 UpdateIdentityPropagationConfig(ctx context.Context, params *UpdateIdentityPropagationConfigInput, optFns ...func(*Options)) (*UpdateIdentityPropagationConfigOutput, error) 
 UpdateIpRestriction(ctx context.Context, params *UpdateIpRestrictionInput, optFns ...func(*Options)) (*UpdateIpRestrictionOutput, error) 
 UpdateKeyRegistration(ctx context.Context, params *UpdateKeyRegistrationInput, optFns ...func(*Options)) (*UpdateKeyRegistrationOutput, error) 
 UpdatePublicSharingSettings(ctx context.Context, params *UpdatePublicSharingSettingsInput, optFns ...func(*Options)) (*UpdatePublicSharingSettingsOutput, error) 
 UpdateRefreshSchedule(ctx context.Context, params *UpdateRefreshScheduleInput, optFns ...func(*Options)) (*UpdateRefreshScheduleOutput, error) 
 UpdateRoleCustomPermission(ctx context.Context, params *UpdateRoleCustomPermissionInput, optFns ...func(*Options)) (*UpdateRoleCustomPermissionOutput, error) 
 UpdateSPICECapacityConfiguration(ctx context.Context, params *UpdateSPICECapacityConfigurationInput, optFns ...func(*Options)) (*UpdateSPICECapacityConfigurationOutput, error) 
 UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) 
 UpdateTemplateAlias(ctx context.Context, params *UpdateTemplateAliasInput, optFns ...func(*Options)) (*UpdateTemplateAliasOutput, error) 
 UpdateTemplatePermissions(ctx context.Context, params *UpdateTemplatePermissionsInput, optFns ...func(*Options)) (*UpdateTemplatePermissionsOutput, error) 
 UpdateTheme(ctx context.Context, params *UpdateThemeInput, optFns ...func(*Options)) (*UpdateThemeOutput, error) 
 UpdateThemeAlias(ctx context.Context, params *UpdateThemeAliasInput, optFns ...func(*Options)) (*UpdateThemeAliasOutput, error) 
 UpdateThemePermissions(ctx context.Context, params *UpdateThemePermissionsInput, optFns ...func(*Options)) (*UpdateThemePermissionsOutput, error) 
 UpdateTopic(ctx context.Context, params *UpdateTopicInput, optFns ...func(*Options)) (*UpdateTopicOutput, error) 
 UpdateTopicPermissions(ctx context.Context, params *UpdateTopicPermissionsInput, optFns ...func(*Options)) (*UpdateTopicPermissionsOutput, error) 
 UpdateTopicRefreshSchedule(ctx context.Context, params *UpdateTopicRefreshScheduleInput, optFns ...func(*Options)) (*UpdateTopicRefreshScheduleOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
 UpdateVPCConnection(ctx context.Context, params *UpdateVPCConnectionInput, optFns ...func(*Options)) (*UpdateVPCConnectionOutput, error) 
}