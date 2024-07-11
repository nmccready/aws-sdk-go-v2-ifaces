
package iam

import (
    "github.com/aws/aws-sdk-go-v2/service/iam"
)

// IIam defines the interface for iam
type IIam interface {
 Options() Options 
 AddClientIDToOpenIDConnectProvider(ctx context.Context, params *AddClientIDToOpenIDConnectProviderInput, optFns ...func(*Options)) (*AddClientIDToOpenIDConnectProviderOutput, error) 
 AddRoleToInstanceProfile(ctx context.Context, params *AddRoleToInstanceProfileInput, optFns ...func(*Options)) (*AddRoleToInstanceProfileOutput, error) 
 AddUserToGroup(ctx context.Context, params *AddUserToGroupInput, optFns ...func(*Options)) (*AddUserToGroupOutput, error) 
 AttachGroupPolicy(ctx context.Context, params *AttachGroupPolicyInput, optFns ...func(*Options)) (*AttachGroupPolicyOutput, error) 
 AttachRolePolicy(ctx context.Context, params *AttachRolePolicyInput, optFns ...func(*Options)) (*AttachRolePolicyOutput, error) 
 AttachUserPolicy(ctx context.Context, params *AttachUserPolicyInput, optFns ...func(*Options)) (*AttachUserPolicyOutput, error) 
 ChangePassword(ctx context.Context, params *ChangePasswordInput, optFns ...func(*Options)) (*ChangePasswordOutput, error) 
 CreateAccessKey(ctx context.Context, params *CreateAccessKeyInput, optFns ...func(*Options)) (*CreateAccessKeyOutput, error) 
 CreateAccountAlias(ctx context.Context, params *CreateAccountAliasInput, optFns ...func(*Options)) (*CreateAccountAliasOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateInstanceProfile(ctx context.Context, params *CreateInstanceProfileInput, optFns ...func(*Options)) (*CreateInstanceProfileOutput, error) 
 CreateLoginProfile(ctx context.Context, params *CreateLoginProfileInput, optFns ...func(*Options)) (*CreateLoginProfileOutput, error) 
 CreateOpenIDConnectProvider(ctx context.Context, params *CreateOpenIDConnectProviderInput, optFns ...func(*Options)) (*CreateOpenIDConnectProviderOutput, error) 
 CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) 
 CreatePolicyVersion(ctx context.Context, params *CreatePolicyVersionInput, optFns ...func(*Options)) (*CreatePolicyVersionOutput, error) 
 CreateRole(ctx context.Context, params *CreateRoleInput, optFns ...func(*Options)) (*CreateRoleOutput, error) 
 CreateSAMLProvider(ctx context.Context, params *CreateSAMLProviderInput, optFns ...func(*Options)) (*CreateSAMLProviderOutput, error) 
 CreateServiceLinkedRole(ctx context.Context, params *CreateServiceLinkedRoleInput, optFns ...func(*Options)) (*CreateServiceLinkedRoleOutput, error) 
 CreateServiceSpecificCredential(ctx context.Context, params *CreateServiceSpecificCredentialInput, optFns ...func(*Options)) (*CreateServiceSpecificCredentialOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 CreateVirtualMFADevice(ctx context.Context, params *CreateVirtualMFADeviceInput, optFns ...func(*Options)) (*CreateVirtualMFADeviceOutput, error) 
 DeactivateMFADevice(ctx context.Context, params *DeactivateMFADeviceInput, optFns ...func(*Options)) (*DeactivateMFADeviceOutput, error) 
 DeleteAccessKey(ctx context.Context, params *DeleteAccessKeyInput, optFns ...func(*Options)) (*DeleteAccessKeyOutput, error) 
 DeleteAccountAlias(ctx context.Context, params *DeleteAccountAliasInput, optFns ...func(*Options)) (*DeleteAccountAliasOutput, error) 
 DeleteAccountPasswordPolicy(ctx context.Context, params *DeleteAccountPasswordPolicyInput, optFns ...func(*Options)) (*DeleteAccountPasswordPolicyOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteGroupPolicy(ctx context.Context, params *DeleteGroupPolicyInput, optFns ...func(*Options)) (*DeleteGroupPolicyOutput, error) 
 DeleteInstanceProfile(ctx context.Context, params *DeleteInstanceProfileInput, optFns ...func(*Options)) (*DeleteInstanceProfileOutput, error) 
 DeleteLoginProfile(ctx context.Context, params *DeleteLoginProfileInput, optFns ...func(*Options)) (*DeleteLoginProfileOutput, error) 
 DeleteOpenIDConnectProvider(ctx context.Context, params *DeleteOpenIDConnectProviderInput, optFns ...func(*Options)) (*DeleteOpenIDConnectProviderOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DeletePolicyVersion(ctx context.Context, params *DeletePolicyVersionInput, optFns ...func(*Options)) (*DeletePolicyVersionOutput, error) 
 DeleteRole(ctx context.Context, params *DeleteRoleInput, optFns ...func(*Options)) (*DeleteRoleOutput, error) 
 DeleteRolePermissionsBoundary(ctx context.Context, params *DeleteRolePermissionsBoundaryInput, optFns ...func(*Options)) (*DeleteRolePermissionsBoundaryOutput, error) 
 DeleteRolePolicy(ctx context.Context, params *DeleteRolePolicyInput, optFns ...func(*Options)) (*DeleteRolePolicyOutput, error) 
 DeleteSAMLProvider(ctx context.Context, params *DeleteSAMLProviderInput, optFns ...func(*Options)) (*DeleteSAMLProviderOutput, error) 
 DeleteSSHPublicKey(ctx context.Context, params *DeleteSSHPublicKeyInput, optFns ...func(*Options)) (*DeleteSSHPublicKeyOutput, error) 
 DeleteServerCertificate(ctx context.Context, params *DeleteServerCertificateInput, optFns ...func(*Options)) (*DeleteServerCertificateOutput, error) 
 DeleteServiceLinkedRole(ctx context.Context, params *DeleteServiceLinkedRoleInput, optFns ...func(*Options)) (*DeleteServiceLinkedRoleOutput, error) 
 DeleteServiceSpecificCredential(ctx context.Context, params *DeleteServiceSpecificCredentialInput, optFns ...func(*Options)) (*DeleteServiceSpecificCredentialOutput, error) 
 DeleteSigningCertificate(ctx context.Context, params *DeleteSigningCertificateInput, optFns ...func(*Options)) (*DeleteSigningCertificateOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DeleteUserPermissionsBoundary(ctx context.Context, params *DeleteUserPermissionsBoundaryInput, optFns ...func(*Options)) (*DeleteUserPermissionsBoundaryOutput, error) 
 DeleteUserPolicy(ctx context.Context, params *DeleteUserPolicyInput, optFns ...func(*Options)) (*DeleteUserPolicyOutput, error) 
 DeleteVirtualMFADevice(ctx context.Context, params *DeleteVirtualMFADeviceInput, optFns ...func(*Options)) (*DeleteVirtualMFADeviceOutput, error) 
 DetachGroupPolicy(ctx context.Context, params *DetachGroupPolicyInput, optFns ...func(*Options)) (*DetachGroupPolicyOutput, error) 
 DetachRolePolicy(ctx context.Context, params *DetachRolePolicyInput, optFns ...func(*Options)) (*DetachRolePolicyOutput, error) 
 DetachUserPolicy(ctx context.Context, params *DetachUserPolicyInput, optFns ...func(*Options)) (*DetachUserPolicyOutput, error) 
 EnableMFADevice(ctx context.Context, params *EnableMFADeviceInput, optFns ...func(*Options)) (*EnableMFADeviceOutput, error) 
 GenerateCredentialReport(ctx context.Context, params *GenerateCredentialReportInput, optFns ...func(*Options)) (*GenerateCredentialReportOutput, error) 
 GenerateOrganizationsAccessReport(ctx context.Context, params *GenerateOrganizationsAccessReportInput, optFns ...func(*Options)) (*GenerateOrganizationsAccessReportOutput, error) 
 GenerateServiceLastAccessedDetails(ctx context.Context, params *GenerateServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GenerateServiceLastAccessedDetailsOutput, error) 
 GetAccessKeyLastUsed(ctx context.Context, params *GetAccessKeyLastUsedInput, optFns ...func(*Options)) (*GetAccessKeyLastUsedOutput, error) 
 GetAccountAuthorizationDetails(ctx context.Context, params *GetAccountAuthorizationDetailsInput, optFns ...func(*Options)) (*GetAccountAuthorizationDetailsOutput, error) 
 GetAccountPasswordPolicy(ctx context.Context, params *GetAccountPasswordPolicyInput, optFns ...func(*Options)) (*GetAccountPasswordPolicyOutput, error) 
 GetAccountSummary(ctx context.Context, params *GetAccountSummaryInput, optFns ...func(*Options)) (*GetAccountSummaryOutput, error) 
 GetContextKeysForCustomPolicy(ctx context.Context, params *GetContextKeysForCustomPolicyInput, optFns ...func(*Options)) (*GetContextKeysForCustomPolicyOutput, error) 
 GetContextKeysForPrincipalPolicy(ctx context.Context, params *GetContextKeysForPrincipalPolicyInput, optFns ...func(*Options)) (*GetContextKeysForPrincipalPolicyOutput, error) 
 GetCredentialReport(ctx context.Context, params *GetCredentialReportInput, optFns ...func(*Options)) (*GetCredentialReportOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetGroupPolicy(ctx context.Context, params *GetGroupPolicyInput, optFns ...func(*Options)) (*GetGroupPolicyOutput, error) 
 GetInstanceProfile(ctx context.Context, params *GetInstanceProfileInput, optFns ...func(*Options)) (*GetInstanceProfileOutput, error) 
 GetLoginProfile(ctx context.Context, params *GetLoginProfileInput, optFns ...func(*Options)) (*GetLoginProfileOutput, error) 
 GetMFADevice(ctx context.Context, params *GetMFADeviceInput, optFns ...func(*Options)) (*GetMFADeviceOutput, error) 
 GetOpenIDConnectProvider(ctx context.Context, params *GetOpenIDConnectProviderInput, optFns ...func(*Options)) (*GetOpenIDConnectProviderOutput, error) 
 GetOrganizationsAccessReport(ctx context.Context, params *GetOrganizationsAccessReportInput, optFns ...func(*Options)) (*GetOrganizationsAccessReportOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetPolicyVersion(ctx context.Context, params *GetPolicyVersionInput, optFns ...func(*Options)) (*GetPolicyVersionOutput, error) 
 GetRole(ctx context.Context, params *GetRoleInput, optFns ...func(*Options)) (*GetRoleOutput, error) 
 GetRolePolicy(ctx context.Context, params *GetRolePolicyInput, optFns ...func(*Options)) (*GetRolePolicyOutput, error) 
 GetSAMLProvider(ctx context.Context, params *GetSAMLProviderInput, optFns ...func(*Options)) (*GetSAMLProviderOutput, error) 
 GetSSHPublicKey(ctx context.Context, params *GetSSHPublicKeyInput, optFns ...func(*Options)) (*GetSSHPublicKeyOutput, error) 
 GetServerCertificate(ctx context.Context, params *GetServerCertificateInput, optFns ...func(*Options)) (*GetServerCertificateOutput, error) 
 GetServiceLastAccessedDetails(ctx context.Context, params *GetServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GetServiceLastAccessedDetailsOutput, error) 
 GetServiceLastAccessedDetailsWithEntities(ctx context.Context, params *GetServiceLastAccessedDetailsWithEntitiesInput, optFns ...func(*Options)) (*GetServiceLastAccessedDetailsWithEntitiesOutput, error) 
 GetServiceLinkedRoleDeletionStatus(ctx context.Context, params *GetServiceLinkedRoleDeletionStatusInput, optFns ...func(*Options)) (*GetServiceLinkedRoleDeletionStatusOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetUserPolicy(ctx context.Context, params *GetUserPolicyInput, optFns ...func(*Options)) (*GetUserPolicyOutput, error) 
 ListAccessKeys(ctx context.Context, params *ListAccessKeysInput, optFns ...func(*Options)) (*ListAccessKeysOutput, error) 
 ListAccountAliases(ctx context.Context, params *ListAccountAliasesInput, optFns ...func(*Options)) (*ListAccountAliasesOutput, error) 
 ListAttachedGroupPolicies(ctx context.Context, params *ListAttachedGroupPoliciesInput, optFns ...func(*Options)) (*ListAttachedGroupPoliciesOutput, error) 
 ListAttachedRolePolicies(ctx context.Context, params *ListAttachedRolePoliciesInput, optFns ...func(*Options)) (*ListAttachedRolePoliciesOutput, error) 
 ListAttachedUserPolicies(ctx context.Context, params *ListAttachedUserPoliciesInput, optFns ...func(*Options)) (*ListAttachedUserPoliciesOutput, error) 
 ListEntitiesForPolicy(ctx context.Context, params *ListEntitiesForPolicyInput, optFns ...func(*Options)) (*ListEntitiesForPolicyOutput, error) 
 ListGroupPolicies(ctx context.Context, params *ListGroupPoliciesInput, optFns ...func(*Options)) (*ListGroupPoliciesOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListGroupsForUser(ctx context.Context, params *ListGroupsForUserInput, optFns ...func(*Options)) (*ListGroupsForUserOutput, error) 
 ListInstanceProfileTags(ctx context.Context, params *ListInstanceProfileTagsInput, optFns ...func(*Options)) (*ListInstanceProfileTagsOutput, error) 
 ListInstanceProfiles(ctx context.Context, params *ListInstanceProfilesInput, optFns ...func(*Options)) (*ListInstanceProfilesOutput, error) 
 ListInstanceProfilesForRole(ctx context.Context, params *ListInstanceProfilesForRoleInput, optFns ...func(*Options)) (*ListInstanceProfilesForRoleOutput, error) 
 ListMFADeviceTags(ctx context.Context, params *ListMFADeviceTagsInput, optFns ...func(*Options)) (*ListMFADeviceTagsOutput, error) 
 ListMFADevices(ctx context.Context, params *ListMFADevicesInput, optFns ...func(*Options)) (*ListMFADevicesOutput, error) 
 ListOpenIDConnectProviderTags(ctx context.Context, params *ListOpenIDConnectProviderTagsInput, optFns ...func(*Options)) (*ListOpenIDConnectProviderTagsOutput, error) 
 ListOpenIDConnectProviders(ctx context.Context, params *ListOpenIDConnectProvidersInput, optFns ...func(*Options)) (*ListOpenIDConnectProvidersOutput, error) 
 ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) 
 ListPoliciesGrantingServiceAccess(ctx context.Context, params *ListPoliciesGrantingServiceAccessInput, optFns ...func(*Options)) (*ListPoliciesGrantingServiceAccessOutput, error) 
 ListPolicyTags(ctx context.Context, params *ListPolicyTagsInput, optFns ...func(*Options)) (*ListPolicyTagsOutput, error) 
 ListPolicyVersions(ctx context.Context, params *ListPolicyVersionsInput, optFns ...func(*Options)) (*ListPolicyVersionsOutput, error) 
 ListRolePolicies(ctx context.Context, params *ListRolePoliciesInput, optFns ...func(*Options)) (*ListRolePoliciesOutput, error) 
 ListRoleTags(ctx context.Context, params *ListRoleTagsInput, optFns ...func(*Options)) (*ListRoleTagsOutput, error) 
 ListRoles(ctx context.Context, params *ListRolesInput, optFns ...func(*Options)) (*ListRolesOutput, error) 
 ListSAMLProviderTags(ctx context.Context, params *ListSAMLProviderTagsInput, optFns ...func(*Options)) (*ListSAMLProviderTagsOutput, error) 
 ListSAMLProviders(ctx context.Context, params *ListSAMLProvidersInput, optFns ...func(*Options)) (*ListSAMLProvidersOutput, error) 
 ListSSHPublicKeys(ctx context.Context, params *ListSSHPublicKeysInput, optFns ...func(*Options)) (*ListSSHPublicKeysOutput, error) 
 ListServerCertificateTags(ctx context.Context, params *ListServerCertificateTagsInput, optFns ...func(*Options)) (*ListServerCertificateTagsOutput, error) 
 ListServerCertificates(ctx context.Context, params *ListServerCertificatesInput, optFns ...func(*Options)) (*ListServerCertificatesOutput, error) 
 ListServiceSpecificCredentials(ctx context.Context, params *ListServiceSpecificCredentialsInput, optFns ...func(*Options)) (*ListServiceSpecificCredentialsOutput, error) 
 ListSigningCertificates(ctx context.Context, params *ListSigningCertificatesInput, optFns ...func(*Options)) (*ListSigningCertificatesOutput, error) 
 ListUserPolicies(ctx context.Context, params *ListUserPoliciesInput, optFns ...func(*Options)) (*ListUserPoliciesOutput, error) 
 ListUserTags(ctx context.Context, params *ListUserTagsInput, optFns ...func(*Options)) (*ListUserTagsOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListVirtualMFADevices(ctx context.Context, params *ListVirtualMFADevicesInput, optFns ...func(*Options)) (*ListVirtualMFADevicesOutput, error) 
 PutGroupPolicy(ctx context.Context, params *PutGroupPolicyInput, optFns ...func(*Options)) (*PutGroupPolicyOutput, error) 
 PutRolePermissionsBoundary(ctx context.Context, params *PutRolePermissionsBoundaryInput, optFns ...func(*Options)) (*PutRolePermissionsBoundaryOutput, error) 
 PutRolePolicy(ctx context.Context, params *PutRolePolicyInput, optFns ...func(*Options)) (*PutRolePolicyOutput, error) 
 PutUserPermissionsBoundary(ctx context.Context, params *PutUserPermissionsBoundaryInput, optFns ...func(*Options)) (*PutUserPermissionsBoundaryOutput, error) 
 PutUserPolicy(ctx context.Context, params *PutUserPolicyInput, optFns ...func(*Options)) (*PutUserPolicyOutput, error) 
 RemoveClientIDFromOpenIDConnectProvider(ctx context.Context, params *RemoveClientIDFromOpenIDConnectProviderInput, optFns ...func(*Options)) (*RemoveClientIDFromOpenIDConnectProviderOutput, error) 
 RemoveRoleFromInstanceProfile(ctx context.Context, params *RemoveRoleFromInstanceProfileInput, optFns ...func(*Options)) (*RemoveRoleFromInstanceProfileOutput, error) 
 RemoveUserFromGroup(ctx context.Context, params *RemoveUserFromGroupInput, optFns ...func(*Options)) (*RemoveUserFromGroupOutput, error) 
 ResetServiceSpecificCredential(ctx context.Context, params *ResetServiceSpecificCredentialInput, optFns ...func(*Options)) (*ResetServiceSpecificCredentialOutput, error) 
 ResyncMFADevice(ctx context.Context, params *ResyncMFADeviceInput, optFns ...func(*Options)) (*ResyncMFADeviceOutput, error) 
 SetDefaultPolicyVersion(ctx context.Context, params *SetDefaultPolicyVersionInput, optFns ...func(*Options)) (*SetDefaultPolicyVersionOutput, error) 
 SetSecurityTokenServicePreferences(ctx context.Context, params *SetSecurityTokenServicePreferencesInput, optFns ...func(*Options)) (*SetSecurityTokenServicePreferencesOutput, error) 
 SimulateCustomPolicy(ctx context.Context, params *SimulateCustomPolicyInput, optFns ...func(*Options)) (*SimulateCustomPolicyOutput, error) 
 SimulatePrincipalPolicy(ctx context.Context, params *SimulatePrincipalPolicyInput, optFns ...func(*Options)) (*SimulatePrincipalPolicyOutput, error) 
 TagInstanceProfile(ctx context.Context, params *TagInstanceProfileInput, optFns ...func(*Options)) (*TagInstanceProfileOutput, error) 
 TagMFADevice(ctx context.Context, params *TagMFADeviceInput, optFns ...func(*Options)) (*TagMFADeviceOutput, error) 
 TagOpenIDConnectProvider(ctx context.Context, params *TagOpenIDConnectProviderInput, optFns ...func(*Options)) (*TagOpenIDConnectProviderOutput, error) 
 TagPolicy(ctx context.Context, params *TagPolicyInput, optFns ...func(*Options)) (*TagPolicyOutput, error) 
 TagRole(ctx context.Context, params *TagRoleInput, optFns ...func(*Options)) (*TagRoleOutput, error) 
 TagSAMLProvider(ctx context.Context, params *TagSAMLProviderInput, optFns ...func(*Options)) (*TagSAMLProviderOutput, error) 
 TagServerCertificate(ctx context.Context, params *TagServerCertificateInput, optFns ...func(*Options)) (*TagServerCertificateOutput, error) 
 TagUser(ctx context.Context, params *TagUserInput, optFns ...func(*Options)) (*TagUserOutput, error) 
 UntagInstanceProfile(ctx context.Context, params *UntagInstanceProfileInput, optFns ...func(*Options)) (*UntagInstanceProfileOutput, error) 
 UntagMFADevice(ctx context.Context, params *UntagMFADeviceInput, optFns ...func(*Options)) (*UntagMFADeviceOutput, error) 
 UntagOpenIDConnectProvider(ctx context.Context, params *UntagOpenIDConnectProviderInput, optFns ...func(*Options)) (*UntagOpenIDConnectProviderOutput, error) 
 UntagPolicy(ctx context.Context, params *UntagPolicyInput, optFns ...func(*Options)) (*UntagPolicyOutput, error) 
 UntagRole(ctx context.Context, params *UntagRoleInput, optFns ...func(*Options)) (*UntagRoleOutput, error) 
 UntagSAMLProvider(ctx context.Context, params *UntagSAMLProviderInput, optFns ...func(*Options)) (*UntagSAMLProviderOutput, error) 
 UntagServerCertificate(ctx context.Context, params *UntagServerCertificateInput, optFns ...func(*Options)) (*UntagServerCertificateOutput, error) 
 UntagUser(ctx context.Context, params *UntagUserInput, optFns ...func(*Options)) (*UntagUserOutput, error) 
 UpdateAccessKey(ctx context.Context, params *UpdateAccessKeyInput, optFns ...func(*Options)) (*UpdateAccessKeyOutput, error) 
 UpdateAccountPasswordPolicy(ctx context.Context, params *UpdateAccountPasswordPolicyInput, optFns ...func(*Options)) (*UpdateAccountPasswordPolicyOutput, error) 
 UpdateAssumeRolePolicy(ctx context.Context, params *UpdateAssumeRolePolicyInput, optFns ...func(*Options)) (*UpdateAssumeRolePolicyOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateLoginProfile(ctx context.Context, params *UpdateLoginProfileInput, optFns ...func(*Options)) (*UpdateLoginProfileOutput, error) 
 UpdateOpenIDConnectProviderThumbprint(ctx context.Context, params *UpdateOpenIDConnectProviderThumbprintInput, optFns ...func(*Options)) (*UpdateOpenIDConnectProviderThumbprintOutput, error) 
 UpdateRole(ctx context.Context, params *UpdateRoleInput, optFns ...func(*Options)) (*UpdateRoleOutput, error) 
 UpdateRoleDescription(ctx context.Context, params *UpdateRoleDescriptionInput, optFns ...func(*Options)) (*UpdateRoleDescriptionOutput, error) 
 UpdateSAMLProvider(ctx context.Context, params *UpdateSAMLProviderInput, optFns ...func(*Options)) (*UpdateSAMLProviderOutput, error) 
 UpdateSSHPublicKey(ctx context.Context, params *UpdateSSHPublicKeyInput, optFns ...func(*Options)) (*UpdateSSHPublicKeyOutput, error) 
 UpdateServerCertificate(ctx context.Context, params *UpdateServerCertificateInput, optFns ...func(*Options)) (*UpdateServerCertificateOutput, error) 
 UpdateServiceSpecificCredential(ctx context.Context, params *UpdateServiceSpecificCredentialInput, optFns ...func(*Options)) (*UpdateServiceSpecificCredentialOutput, error) 
 UpdateSigningCertificate(ctx context.Context, params *UpdateSigningCertificateInput, optFns ...func(*Options)) (*UpdateSigningCertificateOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
 UploadSSHPublicKey(ctx context.Context, params *UploadSSHPublicKeyInput, optFns ...func(*Options)) (*UploadSSHPublicKeyOutput, error) 
 UploadServerCertificate(ctx context.Context, params *UploadServerCertificateInput, optFns ...func(*Options)) (*UploadServerCertificateOutput, error) 
 UploadSigningCertificate(ctx context.Context, params *UploadSigningCertificateInput, optFns ...func(*Options)) (*UploadSigningCertificateOutput, error) 
}
