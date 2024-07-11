
package fms

import (
    "github.com/aws/aws-sdk-go-v2/service/fms"
)

// IFms defines the interface for fms
type IFms interface {
 Options() Options 
 AssociateAdminAccount(ctx context.Context, params *AssociateAdminAccountInput, optFns ...func(*Options)) (*AssociateAdminAccountOutput, error) 
 AssociateThirdPartyFirewall(ctx context.Context, params *AssociateThirdPartyFirewallInput, optFns ...func(*Options)) (*AssociateThirdPartyFirewallOutput, error) 
 BatchAssociateResource(ctx context.Context, params *BatchAssociateResourceInput, optFns ...func(*Options)) (*BatchAssociateResourceOutput, error) 
 BatchDisassociateResource(ctx context.Context, params *BatchDisassociateResourceInput, optFns ...func(*Options)) (*BatchDisassociateResourceOutput, error) 
 DeleteAppsList(ctx context.Context, params *DeleteAppsListInput, optFns ...func(*Options)) (*DeleteAppsListOutput, error) 
 DeleteNotificationChannel(ctx context.Context, params *DeleteNotificationChannelInput, optFns ...func(*Options)) (*DeleteNotificationChannelOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DeleteProtocolsList(ctx context.Context, params *DeleteProtocolsListInput, optFns ...func(*Options)) (*DeleteProtocolsListOutput, error) 
 DeleteResourceSet(ctx context.Context, params *DeleteResourceSetInput, optFns ...func(*Options)) (*DeleteResourceSetOutput, error) 
 DisassociateAdminAccount(ctx context.Context, params *DisassociateAdminAccountInput, optFns ...func(*Options)) (*DisassociateAdminAccountOutput, error) 
 DisassociateThirdPartyFirewall(ctx context.Context, params *DisassociateThirdPartyFirewallInput, optFns ...func(*Options)) (*DisassociateThirdPartyFirewallOutput, error) 
 GetAdminAccount(ctx context.Context, params *GetAdminAccountInput, optFns ...func(*Options)) (*GetAdminAccountOutput, error) 
 GetAdminScope(ctx context.Context, params *GetAdminScopeInput, optFns ...func(*Options)) (*GetAdminScopeOutput, error) 
 GetAppsList(ctx context.Context, params *GetAppsListInput, optFns ...func(*Options)) (*GetAppsListOutput, error) 
 GetComplianceDetail(ctx context.Context, params *GetComplianceDetailInput, optFns ...func(*Options)) (*GetComplianceDetailOutput, error) 
 GetNotificationChannel(ctx context.Context, params *GetNotificationChannelInput, optFns ...func(*Options)) (*GetNotificationChannelOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetProtectionStatus(ctx context.Context, params *GetProtectionStatusInput, optFns ...func(*Options)) (*GetProtectionStatusOutput, error) 
 GetProtocolsList(ctx context.Context, params *GetProtocolsListInput, optFns ...func(*Options)) (*GetProtocolsListOutput, error) 
 GetResourceSet(ctx context.Context, params *GetResourceSetInput, optFns ...func(*Options)) (*GetResourceSetOutput, error) 
 GetThirdPartyFirewallAssociationStatus(ctx context.Context, params *GetThirdPartyFirewallAssociationStatusInput, optFns ...func(*Options)) (*GetThirdPartyFirewallAssociationStatusOutput, error) 
 GetViolationDetails(ctx context.Context, params *GetViolationDetailsInput, optFns ...func(*Options)) (*GetViolationDetailsOutput, error) 
 ListAdminAccountsForOrganization(ctx context.Context, params *ListAdminAccountsForOrganizationInput, optFns ...func(*Options)) (*ListAdminAccountsForOrganizationOutput, error) 
 ListAdminsManagingAccount(ctx context.Context, params *ListAdminsManagingAccountInput, optFns ...func(*Options)) (*ListAdminsManagingAccountOutput, error) 
 ListAppsLists(ctx context.Context, params *ListAppsListsInput, optFns ...func(*Options)) (*ListAppsListsOutput, error) 
 ListComplianceStatus(ctx context.Context, params *ListComplianceStatusInput, optFns ...func(*Options)) (*ListComplianceStatusOutput, error) 
 ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) 
 ListMemberAccounts(ctx context.Context, params *ListMemberAccountsInput, optFns ...func(*Options)) (*ListMemberAccountsOutput, error) 
 ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) 
 ListProtocolsLists(ctx context.Context, params *ListProtocolsListsInput, optFns ...func(*Options)) (*ListProtocolsListsOutput, error) 
 ListResourceSetResources(ctx context.Context, params *ListResourceSetResourcesInput, optFns ...func(*Options)) (*ListResourceSetResourcesOutput, error) 
 ListResourceSets(ctx context.Context, params *ListResourceSetsInput, optFns ...func(*Options)) (*ListResourceSetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListThirdPartyFirewallFirewallPolicies(ctx context.Context, params *ListThirdPartyFirewallFirewallPoliciesInput, optFns ...func(*Options)) (*ListThirdPartyFirewallFirewallPoliciesOutput, error) 
 PutAdminAccount(ctx context.Context, params *PutAdminAccountInput, optFns ...func(*Options)) (*PutAdminAccountOutput, error) 
 PutAppsList(ctx context.Context, params *PutAppsListInput, optFns ...func(*Options)) (*PutAppsListOutput, error) 
 PutNotificationChannel(ctx context.Context, params *PutNotificationChannelInput, optFns ...func(*Options)) (*PutNotificationChannelOutput, error) 
 PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) 
 PutProtocolsList(ctx context.Context, params *PutProtocolsListInput, optFns ...func(*Options)) (*PutProtocolsListOutput, error) 
 PutResourceSet(ctx context.Context, params *PutResourceSetInput, optFns ...func(*Options)) (*PutResourceSetOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
