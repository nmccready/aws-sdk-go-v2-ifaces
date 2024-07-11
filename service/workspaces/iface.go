
package workspaces

import (
    "github.com/aws/aws-sdk-go-v2/service/workspaces"
)

// IClient defines the interface for workspaces
type IClient interface {
 Options() Options 
 AcceptAccountLinkInvitation(ctx context.Context, params *AcceptAccountLinkInvitationInput, optFns ...func(*Options)) (*AcceptAccountLinkInvitationOutput, error) 
 AssociateConnectionAlias(ctx context.Context, params *AssociateConnectionAliasInput, optFns ...func(*Options)) (*AssociateConnectionAliasOutput, error) 
 AssociateIpGroups(ctx context.Context, params *AssociateIpGroupsInput, optFns ...func(*Options)) (*AssociateIpGroupsOutput, error) 
 AssociateWorkspaceApplication(ctx context.Context, params *AssociateWorkspaceApplicationInput, optFns ...func(*Options)) (*AssociateWorkspaceApplicationOutput, error) 
 AuthorizeIpRules(ctx context.Context, params *AuthorizeIpRulesInput, optFns ...func(*Options)) (*AuthorizeIpRulesOutput, error) 
 CopyWorkspaceImage(ctx context.Context, params *CopyWorkspaceImageInput, optFns ...func(*Options)) (*CopyWorkspaceImageOutput, error) 
 CreateAccountLinkInvitation(ctx context.Context, params *CreateAccountLinkInvitationInput, optFns ...func(*Options)) (*CreateAccountLinkInvitationOutput, error) 
 CreateConnectClientAddIn(ctx context.Context, params *CreateConnectClientAddInInput, optFns ...func(*Options)) (*CreateConnectClientAddInOutput, error) 
 CreateConnectionAlias(ctx context.Context, params *CreateConnectionAliasInput, optFns ...func(*Options)) (*CreateConnectionAliasOutput, error) 
 CreateIpGroup(ctx context.Context, params *CreateIpGroupInput, optFns ...func(*Options)) (*CreateIpGroupOutput, error) 
 CreateStandbyWorkspaces(ctx context.Context, params *CreateStandbyWorkspacesInput, optFns ...func(*Options)) (*CreateStandbyWorkspacesOutput, error) 
 CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) 
 CreateUpdatedWorkspaceImage(ctx context.Context, params *CreateUpdatedWorkspaceImageInput, optFns ...func(*Options)) (*CreateUpdatedWorkspaceImageOutput, error) 
 CreateWorkspaceBundle(ctx context.Context, params *CreateWorkspaceBundleInput, optFns ...func(*Options)) (*CreateWorkspaceBundleOutput, error) 
 CreateWorkspaceImage(ctx context.Context, params *CreateWorkspaceImageInput, optFns ...func(*Options)) (*CreateWorkspaceImageOutput, error) 
 CreateWorkspaces(ctx context.Context, params *CreateWorkspacesInput, optFns ...func(*Options)) (*CreateWorkspacesOutput, error) 
 CreateWorkspacesPool(ctx context.Context, params *CreateWorkspacesPoolInput, optFns ...func(*Options)) (*CreateWorkspacesPoolOutput, error) 
 DeleteAccountLinkInvitation(ctx context.Context, params *DeleteAccountLinkInvitationInput, optFns ...func(*Options)) (*DeleteAccountLinkInvitationOutput, error) 
 DeleteClientBranding(ctx context.Context, params *DeleteClientBrandingInput, optFns ...func(*Options)) (*DeleteClientBrandingOutput, error) 
 DeleteConnectClientAddIn(ctx context.Context, params *DeleteConnectClientAddInInput, optFns ...func(*Options)) (*DeleteConnectClientAddInOutput, error) 
 DeleteConnectionAlias(ctx context.Context, params *DeleteConnectionAliasInput, optFns ...func(*Options)) (*DeleteConnectionAliasOutput, error) 
 DeleteIpGroup(ctx context.Context, params *DeleteIpGroupInput, optFns ...func(*Options)) (*DeleteIpGroupOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DeleteWorkspaceBundle(ctx context.Context, params *DeleteWorkspaceBundleInput, optFns ...func(*Options)) (*DeleteWorkspaceBundleOutput, error) 
 DeleteWorkspaceImage(ctx context.Context, params *DeleteWorkspaceImageInput, optFns ...func(*Options)) (*DeleteWorkspaceImageOutput, error) 
 DeployWorkspaceApplications(ctx context.Context, params *DeployWorkspaceApplicationsInput, optFns ...func(*Options)) (*DeployWorkspaceApplicationsOutput, error) 
 DeregisterWorkspaceDirectory(ctx context.Context, params *DeregisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*DeregisterWorkspaceDirectoryOutput, error) 
 DescribeAccount(ctx context.Context, params *DescribeAccountInput, optFns ...func(*Options)) (*DescribeAccountOutput, error) 
 DescribeAccountModifications(ctx context.Context, params *DescribeAccountModificationsInput, optFns ...func(*Options)) (*DescribeAccountModificationsOutput, error) 
 DescribeApplicationAssociations(ctx context.Context, params *DescribeApplicationAssociationsInput, optFns ...func(*Options)) (*DescribeApplicationAssociationsOutput, error) 
 DescribeApplications(ctx context.Context, params *DescribeApplicationsInput, optFns ...func(*Options)) (*DescribeApplicationsOutput, error) 
 DescribeBundleAssociations(ctx context.Context, params *DescribeBundleAssociationsInput, optFns ...func(*Options)) (*DescribeBundleAssociationsOutput, error) 
 DescribeClientBranding(ctx context.Context, params *DescribeClientBrandingInput, optFns ...func(*Options)) (*DescribeClientBrandingOutput, error) 
 DescribeClientProperties(ctx context.Context, params *DescribeClientPropertiesInput, optFns ...func(*Options)) (*DescribeClientPropertiesOutput, error) 
 DescribeConnectClientAddIns(ctx context.Context, params *DescribeConnectClientAddInsInput, optFns ...func(*Options)) (*DescribeConnectClientAddInsOutput, error) 
 DescribeConnectionAliasPermissions(ctx context.Context, params *DescribeConnectionAliasPermissionsInput, optFns ...func(*Options)) (*DescribeConnectionAliasPermissionsOutput, error) 
 DescribeConnectionAliases(ctx context.Context, params *DescribeConnectionAliasesInput, optFns ...func(*Options)) (*DescribeConnectionAliasesOutput, error) 
 DescribeImageAssociations(ctx context.Context, params *DescribeImageAssociationsInput, optFns ...func(*Options)) (*DescribeImageAssociationsOutput, error) 
 DescribeIpGroups(ctx context.Context, params *DescribeIpGroupsInput, optFns ...func(*Options)) (*DescribeIpGroupsOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 DescribeWorkspaceAssociations(ctx context.Context, params *DescribeWorkspaceAssociationsInput, optFns ...func(*Options)) (*DescribeWorkspaceAssociationsOutput, error) 
 DescribeWorkspaceBundles(ctx context.Context, params *DescribeWorkspaceBundlesInput, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) 
 DescribeWorkspaceDirectories(ctx context.Context, params *DescribeWorkspaceDirectoriesInput, optFns ...func(*Options)) (*DescribeWorkspaceDirectoriesOutput, error) 
 DescribeWorkspaceImagePermissions(ctx context.Context, params *DescribeWorkspaceImagePermissionsInput, optFns ...func(*Options)) (*DescribeWorkspaceImagePermissionsOutput, error) 
 DescribeWorkspaceImages(ctx context.Context, params *DescribeWorkspaceImagesInput, optFns ...func(*Options)) (*DescribeWorkspaceImagesOutput, error) 
 DescribeWorkspaceSnapshots(ctx context.Context, params *DescribeWorkspaceSnapshotsInput, optFns ...func(*Options)) (*DescribeWorkspaceSnapshotsOutput, error) 
 DescribeWorkspaces(ctx context.Context, params *DescribeWorkspacesInput, optFns ...func(*Options)) (*DescribeWorkspacesOutput, error) 
 DescribeWorkspacesConnectionStatus(ctx context.Context, params *DescribeWorkspacesConnectionStatusInput, optFns ...func(*Options)) (*DescribeWorkspacesConnectionStatusOutput, error) 
 DescribeWorkspacesPoolSessions(ctx context.Context, params *DescribeWorkspacesPoolSessionsInput, optFns ...func(*Options)) (*DescribeWorkspacesPoolSessionsOutput, error) 
 DescribeWorkspacesPools(ctx context.Context, params *DescribeWorkspacesPoolsInput, optFns ...func(*Options)) (*DescribeWorkspacesPoolsOutput, error) 
 DisassociateConnectionAlias(ctx context.Context, params *DisassociateConnectionAliasInput, optFns ...func(*Options)) (*DisassociateConnectionAliasOutput, error) 
 DisassociateIpGroups(ctx context.Context, params *DisassociateIpGroupsInput, optFns ...func(*Options)) (*DisassociateIpGroupsOutput, error) 
 DisassociateWorkspaceApplication(ctx context.Context, params *DisassociateWorkspaceApplicationInput, optFns ...func(*Options)) (*DisassociateWorkspaceApplicationOutput, error) 
 GetAccountLink(ctx context.Context, params *GetAccountLinkInput, optFns ...func(*Options)) (*GetAccountLinkOutput, error) 
 ImportClientBranding(ctx context.Context, params *ImportClientBrandingInput, optFns ...func(*Options)) (*ImportClientBrandingOutput, error) 
 ImportWorkspaceImage(ctx context.Context, params *ImportWorkspaceImageInput, optFns ...func(*Options)) (*ImportWorkspaceImageOutput, error) 
 ListAccountLinks(ctx context.Context, params *ListAccountLinksInput, optFns ...func(*Options)) (*ListAccountLinksOutput, error) 
 ListAvailableManagementCidrRanges(ctx context.Context, params *ListAvailableManagementCidrRangesInput, optFns ...func(*Options)) (*ListAvailableManagementCidrRangesOutput, error) 
 MigrateWorkspace(ctx context.Context, params *MigrateWorkspaceInput, optFns ...func(*Options)) (*MigrateWorkspaceOutput, error) 
 ModifyAccount(ctx context.Context, params *ModifyAccountInput, optFns ...func(*Options)) (*ModifyAccountOutput, error) 
 ModifyCertificateBasedAuthProperties(ctx context.Context, params *ModifyCertificateBasedAuthPropertiesInput, optFns ...func(*Options)) (*ModifyCertificateBasedAuthPropertiesOutput, error) 
 ModifyClientProperties(ctx context.Context, params *ModifyClientPropertiesInput, optFns ...func(*Options)) (*ModifyClientPropertiesOutput, error) 
 ModifySamlProperties(ctx context.Context, params *ModifySamlPropertiesInput, optFns ...func(*Options)) (*ModifySamlPropertiesOutput, error) 
 ModifySelfservicePermissions(ctx context.Context, params *ModifySelfservicePermissionsInput, optFns ...func(*Options)) (*ModifySelfservicePermissionsOutput, error) 
 ModifyStreamingProperties(ctx context.Context, params *ModifyStreamingPropertiesInput, optFns ...func(*Options)) (*ModifyStreamingPropertiesOutput, error) 
 ModifyWorkspaceAccessProperties(ctx context.Context, params *ModifyWorkspaceAccessPropertiesInput, optFns ...func(*Options)) (*ModifyWorkspaceAccessPropertiesOutput, error) 
 ModifyWorkspaceCreationProperties(ctx context.Context, params *ModifyWorkspaceCreationPropertiesInput, optFns ...func(*Options)) (*ModifyWorkspaceCreationPropertiesOutput, error) 
 ModifyWorkspaceProperties(ctx context.Context, params *ModifyWorkspacePropertiesInput, optFns ...func(*Options)) (*ModifyWorkspacePropertiesOutput, error) 
 ModifyWorkspaceState(ctx context.Context, params *ModifyWorkspaceStateInput, optFns ...func(*Options)) (*ModifyWorkspaceStateOutput, error) 
 RebootWorkspaces(ctx context.Context, params *RebootWorkspacesInput, optFns ...func(*Options)) (*RebootWorkspacesOutput, error) 
 RebuildWorkspaces(ctx context.Context, params *RebuildWorkspacesInput, optFns ...func(*Options)) (*RebuildWorkspacesOutput, error) 
 RegisterWorkspaceDirectory(ctx context.Context, params *RegisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*RegisterWorkspaceDirectoryOutput, error) 
 RejectAccountLinkInvitation(ctx context.Context, params *RejectAccountLinkInvitationInput, optFns ...func(*Options)) (*RejectAccountLinkInvitationOutput, error) 
 RestoreWorkspace(ctx context.Context, params *RestoreWorkspaceInput, optFns ...func(*Options)) (*RestoreWorkspaceOutput, error) 
 RevokeIpRules(ctx context.Context, params *RevokeIpRulesInput, optFns ...func(*Options)) (*RevokeIpRulesOutput, error) 
 StartWorkspaces(ctx context.Context, params *StartWorkspacesInput, optFns ...func(*Options)) (*StartWorkspacesOutput, error) 
 StartWorkspacesPool(ctx context.Context, params *StartWorkspacesPoolInput, optFns ...func(*Options)) (*StartWorkspacesPoolOutput, error) 
 StopWorkspaces(ctx context.Context, params *StopWorkspacesInput, optFns ...func(*Options)) (*StopWorkspacesOutput, error) 
 StopWorkspacesPool(ctx context.Context, params *StopWorkspacesPoolInput, optFns ...func(*Options)) (*StopWorkspacesPoolOutput, error) 
 TerminateWorkspaces(ctx context.Context, params *TerminateWorkspacesInput, optFns ...func(*Options)) (*TerminateWorkspacesOutput, error) 
 TerminateWorkspacesPool(ctx context.Context, params *TerminateWorkspacesPoolInput, optFns ...func(*Options)) (*TerminateWorkspacesPoolOutput, error) 
 TerminateWorkspacesPoolSession(ctx context.Context, params *TerminateWorkspacesPoolSessionInput, optFns ...func(*Options)) (*TerminateWorkspacesPoolSessionOutput, error) 
 UpdateConnectClientAddIn(ctx context.Context, params *UpdateConnectClientAddInInput, optFns ...func(*Options)) (*UpdateConnectClientAddInOutput, error) 
 UpdateConnectionAliasPermission(ctx context.Context, params *UpdateConnectionAliasPermissionInput, optFns ...func(*Options)) (*UpdateConnectionAliasPermissionOutput, error) 
 UpdateRulesOfIpGroup(ctx context.Context, params *UpdateRulesOfIpGroupInput, optFns ...func(*Options)) (*UpdateRulesOfIpGroupOutput, error) 
 UpdateWorkspaceBundle(ctx context.Context, params *UpdateWorkspaceBundleInput, optFns ...func(*Options)) (*UpdateWorkspaceBundleOutput, error) 
 UpdateWorkspaceImagePermission(ctx context.Context, params *UpdateWorkspaceImagePermissionInput, optFns ...func(*Options)) (*UpdateWorkspaceImagePermissionOutput, error) 
 UpdateWorkspacesPool(ctx context.Context, params *UpdateWorkspacesPoolInput, optFns ...func(*Options)) (*UpdateWorkspacesPoolOutput, error) 
}
