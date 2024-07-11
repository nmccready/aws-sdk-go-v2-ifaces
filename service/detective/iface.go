
package detective

import (
    "github.com/aws/aws-sdk-go-v2/service/detective"
)

// IClient defines the interface for detective
type IClient interface {
 Options() Options 
 AcceptInvitation(ctx context.Context, params *AcceptInvitationInput, optFns ...func(*Options)) (*AcceptInvitationOutput, error) 
 BatchGetGraphMemberDatasources(ctx context.Context, params *BatchGetGraphMemberDatasourcesInput, optFns ...func(*Options)) (*BatchGetGraphMemberDatasourcesOutput, error) 
 BatchGetMembershipDatasources(ctx context.Context, params *BatchGetMembershipDatasourcesInput, optFns ...func(*Options)) (*BatchGetMembershipDatasourcesOutput, error) 
 CreateGraph(ctx context.Context, params *CreateGraphInput, optFns ...func(*Options)) (*CreateGraphOutput, error) 
 CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) 
 DeleteGraph(ctx context.Context, params *DeleteGraphInput, optFns ...func(*Options)) (*DeleteGraphOutput, error) 
 DeleteMembers(ctx context.Context, params *DeleteMembersInput, optFns ...func(*Options)) (*DeleteMembersOutput, error) 
 DescribeOrganizationConfiguration(ctx context.Context, params *DescribeOrganizationConfigurationInput, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) 
 DisableOrganizationAdminAccount(ctx context.Context, params *DisableOrganizationAdminAccountInput, optFns ...func(*Options)) (*DisableOrganizationAdminAccountOutput, error) 
 DisassociateMembership(ctx context.Context, params *DisassociateMembershipInput, optFns ...func(*Options)) (*DisassociateMembershipOutput, error) 
 EnableOrganizationAdminAccount(ctx context.Context, params *EnableOrganizationAdminAccountInput, optFns ...func(*Options)) (*EnableOrganizationAdminAccountOutput, error) 
 GetInvestigation(ctx context.Context, params *GetInvestigationInput, optFns ...func(*Options)) (*GetInvestigationOutput, error) 
 GetMembers(ctx context.Context, params *GetMembersInput, optFns ...func(*Options)) (*GetMembersOutput, error) 
 ListDatasourcePackages(ctx context.Context, params *ListDatasourcePackagesInput, optFns ...func(*Options)) (*ListDatasourcePackagesOutput, error) 
 ListGraphs(ctx context.Context, params *ListGraphsInput, optFns ...func(*Options)) (*ListGraphsOutput, error) 
 ListIndicators(ctx context.Context, params *ListIndicatorsInput, optFns ...func(*Options)) (*ListIndicatorsOutput, error) 
 ListInvestigations(ctx context.Context, params *ListInvestigationsInput, optFns ...func(*Options)) (*ListInvestigationsOutput, error) 
 ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) 
 ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) 
 ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RejectInvitation(ctx context.Context, params *RejectInvitationInput, optFns ...func(*Options)) (*RejectInvitationOutput, error) 
 StartInvestigation(ctx context.Context, params *StartInvestigationInput, optFns ...func(*Options)) (*StartInvestigationOutput, error) 
 StartMonitoringMember(ctx context.Context, params *StartMonitoringMemberInput, optFns ...func(*Options)) (*StartMonitoringMemberOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDatasourcePackages(ctx context.Context, params *UpdateDatasourcePackagesInput, optFns ...func(*Options)) (*UpdateDatasourcePackagesOutput, error) 
 UpdateInvestigationState(ctx context.Context, params *UpdateInvestigationStateInput, optFns ...func(*Options)) (*UpdateInvestigationStateOutput, error) 
 UpdateOrganizationConfiguration(ctx context.Context, params *UpdateOrganizationConfigurationInput, optFns ...func(*Options)) (*UpdateOrganizationConfigurationOutput, error) 
}
