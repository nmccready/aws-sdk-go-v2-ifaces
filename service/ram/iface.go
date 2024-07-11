
package ram

import (
    "github.com/aws/aws-sdk-go-v2/service/ram"
)

// IClient defines the interface for ram
type IClient interface {
 Options() Options 
 AcceptResourceShareInvitation(ctx context.Context, params *AcceptResourceShareInvitationInput, optFns ...func(*Options)) (*AcceptResourceShareInvitationOutput, error) 
 AssociateResourceShare(ctx context.Context, params *AssociateResourceShareInput, optFns ...func(*Options)) (*AssociateResourceShareOutput, error) 
 AssociateResourceSharePermission(ctx context.Context, params *AssociateResourceSharePermissionInput, optFns ...func(*Options)) (*AssociateResourceSharePermissionOutput, error) 
 CreatePermission(ctx context.Context, params *CreatePermissionInput, optFns ...func(*Options)) (*CreatePermissionOutput, error) 
 CreatePermissionVersion(ctx context.Context, params *CreatePermissionVersionInput, optFns ...func(*Options)) (*CreatePermissionVersionOutput, error) 
 CreateResourceShare(ctx context.Context, params *CreateResourceShareInput, optFns ...func(*Options)) (*CreateResourceShareOutput, error) 
 DeletePermission(ctx context.Context, params *DeletePermissionInput, optFns ...func(*Options)) (*DeletePermissionOutput, error) 
 DeletePermissionVersion(ctx context.Context, params *DeletePermissionVersionInput, optFns ...func(*Options)) (*DeletePermissionVersionOutput, error) 
 DeleteResourceShare(ctx context.Context, params *DeleteResourceShareInput, optFns ...func(*Options)) (*DeleteResourceShareOutput, error) 
 DisassociateResourceShare(ctx context.Context, params *DisassociateResourceShareInput, optFns ...func(*Options)) (*DisassociateResourceShareOutput, error) 
 DisassociateResourceSharePermission(ctx context.Context, params *DisassociateResourceSharePermissionInput, optFns ...func(*Options)) (*DisassociateResourceSharePermissionOutput, error) 
 EnableSharingWithAwsOrganization(ctx context.Context, params *EnableSharingWithAwsOrganizationInput, optFns ...func(*Options)) (*EnableSharingWithAwsOrganizationOutput, error) 
 GetPermission(ctx context.Context, params *GetPermissionInput, optFns ...func(*Options)) (*GetPermissionOutput, error) 
 GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) 
 GetResourceShareAssociations(ctx context.Context, params *GetResourceShareAssociationsInput, optFns ...func(*Options)) (*GetResourceShareAssociationsOutput, error) 
 GetResourceShareInvitations(ctx context.Context, params *GetResourceShareInvitationsInput, optFns ...func(*Options)) (*GetResourceShareInvitationsOutput, error) 
 GetResourceShares(ctx context.Context, params *GetResourceSharesInput, optFns ...func(*Options)) (*GetResourceSharesOutput, error) 
 ListPendingInvitationResources(ctx context.Context, params *ListPendingInvitationResourcesInput, optFns ...func(*Options)) (*ListPendingInvitationResourcesOutput, error) 
 ListPermissionAssociations(ctx context.Context, params *ListPermissionAssociationsInput, optFns ...func(*Options)) (*ListPermissionAssociationsOutput, error) 
 ListPermissionVersions(ctx context.Context, params *ListPermissionVersionsInput, optFns ...func(*Options)) (*ListPermissionVersionsOutput, error) 
 ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) 
 ListPrincipals(ctx context.Context, params *ListPrincipalsInput, optFns ...func(*Options)) (*ListPrincipalsOutput, error) 
 ListReplacePermissionAssociationsWork(ctx context.Context, params *ListReplacePermissionAssociationsWorkInput, optFns ...func(*Options)) (*ListReplacePermissionAssociationsWorkOutput, error) 
 ListResourceSharePermissions(ctx context.Context, params *ListResourceSharePermissionsInput, optFns ...func(*Options)) (*ListResourceSharePermissionsOutput, error) 
 ListResourceTypes(ctx context.Context, params *ListResourceTypesInput, optFns ...func(*Options)) (*ListResourceTypesOutput, error) 
 ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) 
 PromotePermissionCreatedFromPolicy(ctx context.Context, params *PromotePermissionCreatedFromPolicyInput, optFns ...func(*Options)) (*PromotePermissionCreatedFromPolicyOutput, error) 
 PromoteResourceShareCreatedFromPolicy(ctx context.Context, params *PromoteResourceShareCreatedFromPolicyInput, optFns ...func(*Options)) (*PromoteResourceShareCreatedFromPolicyOutput, error) 
 RejectResourceShareInvitation(ctx context.Context, params *RejectResourceShareInvitationInput, optFns ...func(*Options)) (*RejectResourceShareInvitationOutput, error) 
 ReplacePermissionAssociations(ctx context.Context, params *ReplacePermissionAssociationsInput, optFns ...func(*Options)) (*ReplacePermissionAssociationsOutput, error) 
 SetDefaultPermissionVersion(ctx context.Context, params *SetDefaultPermissionVersionInput, optFns ...func(*Options)) (*SetDefaultPermissionVersionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateResourceShare(ctx context.Context, params *UpdateResourceShareInput, optFns ...func(*Options)) (*UpdateResourceShareOutput, error) 
}
