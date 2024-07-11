
package ssoadmin

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

// IClient defines the interface for ssoadmin
type IClient interface {
 Options() Options 
 AttachCustomerManagedPolicyReferenceToPermissionSet(ctx context.Context, params *AttachCustomerManagedPolicyReferenceToPermissionSetInput, optFns ...func(*Options)) (*AttachCustomerManagedPolicyReferenceToPermissionSetOutput, error) 
 AttachManagedPolicyToPermissionSet(ctx context.Context, params *AttachManagedPolicyToPermissionSetInput, optFns ...func(*Options)) (*AttachManagedPolicyToPermissionSetOutput, error) 
 CreateAccountAssignment(ctx context.Context, params *CreateAccountAssignmentInput, optFns ...func(*Options)) (*CreateAccountAssignmentOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateApplicationAssignment(ctx context.Context, params *CreateApplicationAssignmentInput, optFns ...func(*Options)) (*CreateApplicationAssignmentOutput, error) 
 CreateInstance(ctx context.Context, params *CreateInstanceInput, optFns ...func(*Options)) (*CreateInstanceOutput, error) 
 CreateInstanceAccessControlAttributeConfiguration(ctx context.Context, params *CreateInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*CreateInstanceAccessControlAttributeConfigurationOutput, error) 
 CreatePermissionSet(ctx context.Context, params *CreatePermissionSetInput, optFns ...func(*Options)) (*CreatePermissionSetOutput, error) 
 CreateTrustedTokenIssuer(ctx context.Context, params *CreateTrustedTokenIssuerInput, optFns ...func(*Options)) (*CreateTrustedTokenIssuerOutput, error) 
 DeleteAccountAssignment(ctx context.Context, params *DeleteAccountAssignmentInput, optFns ...func(*Options)) (*DeleteAccountAssignmentOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteApplicationAccessScope(ctx context.Context, params *DeleteApplicationAccessScopeInput, optFns ...func(*Options)) (*DeleteApplicationAccessScopeOutput, error) 
 DeleteApplicationAssignment(ctx context.Context, params *DeleteApplicationAssignmentInput, optFns ...func(*Options)) (*DeleteApplicationAssignmentOutput, error) 
 DeleteApplicationAuthenticationMethod(ctx context.Context, params *DeleteApplicationAuthenticationMethodInput, optFns ...func(*Options)) (*DeleteApplicationAuthenticationMethodOutput, error) 
 DeleteApplicationGrant(ctx context.Context, params *DeleteApplicationGrantInput, optFns ...func(*Options)) (*DeleteApplicationGrantOutput, error) 
 DeleteInlinePolicyFromPermissionSet(ctx context.Context, params *DeleteInlinePolicyFromPermissionSetInput, optFns ...func(*Options)) (*DeleteInlinePolicyFromPermissionSetOutput, error) 
 DeleteInstance(ctx context.Context, params *DeleteInstanceInput, optFns ...func(*Options)) (*DeleteInstanceOutput, error) 
 DeleteInstanceAccessControlAttributeConfiguration(ctx context.Context, params *DeleteInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*DeleteInstanceAccessControlAttributeConfigurationOutput, error) 
 DeletePermissionSet(ctx context.Context, params *DeletePermissionSetInput, optFns ...func(*Options)) (*DeletePermissionSetOutput, error) 
 DeletePermissionsBoundaryFromPermissionSet(ctx context.Context, params *DeletePermissionsBoundaryFromPermissionSetInput, optFns ...func(*Options)) (*DeletePermissionsBoundaryFromPermissionSetOutput, error) 
 DeleteTrustedTokenIssuer(ctx context.Context, params *DeleteTrustedTokenIssuerInput, optFns ...func(*Options)) (*DeleteTrustedTokenIssuerOutput, error) 
 DescribeAccountAssignmentCreationStatus(ctx context.Context, params *DescribeAccountAssignmentCreationStatusInput, optFns ...func(*Options)) (*DescribeAccountAssignmentCreationStatusOutput, error) 
 DescribeAccountAssignmentDeletionStatus(ctx context.Context, params *DescribeAccountAssignmentDeletionStatusInput, optFns ...func(*Options)) (*DescribeAccountAssignmentDeletionStatusOutput, error) 
 DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) 
 DescribeApplicationAssignment(ctx context.Context, params *DescribeApplicationAssignmentInput, optFns ...func(*Options)) (*DescribeApplicationAssignmentOutput, error) 
 DescribeApplicationProvider(ctx context.Context, params *DescribeApplicationProviderInput, optFns ...func(*Options)) (*DescribeApplicationProviderOutput, error) 
 DescribeInstance(ctx context.Context, params *DescribeInstanceInput, optFns ...func(*Options)) (*DescribeInstanceOutput, error) 
 DescribeInstanceAccessControlAttributeConfiguration(ctx context.Context, params *DescribeInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*DescribeInstanceAccessControlAttributeConfigurationOutput, error) 
 DescribePermissionSet(ctx context.Context, params *DescribePermissionSetInput, optFns ...func(*Options)) (*DescribePermissionSetOutput, error) 
 DescribePermissionSetProvisioningStatus(ctx context.Context, params *DescribePermissionSetProvisioningStatusInput, optFns ...func(*Options)) (*DescribePermissionSetProvisioningStatusOutput, error) 
 DescribeTrustedTokenIssuer(ctx context.Context, params *DescribeTrustedTokenIssuerInput, optFns ...func(*Options)) (*DescribeTrustedTokenIssuerOutput, error) 
 DetachCustomerManagedPolicyReferenceFromPermissionSet(ctx context.Context, params *DetachCustomerManagedPolicyReferenceFromPermissionSetInput, optFns ...func(*Options)) (*DetachCustomerManagedPolicyReferenceFromPermissionSetOutput, error) 
 DetachManagedPolicyFromPermissionSet(ctx context.Context, params *DetachManagedPolicyFromPermissionSetInput, optFns ...func(*Options)) (*DetachManagedPolicyFromPermissionSetOutput, error) 
 GetApplicationAccessScope(ctx context.Context, params *GetApplicationAccessScopeInput, optFns ...func(*Options)) (*GetApplicationAccessScopeOutput, error) 
 GetApplicationAssignmentConfiguration(ctx context.Context, params *GetApplicationAssignmentConfigurationInput, optFns ...func(*Options)) (*GetApplicationAssignmentConfigurationOutput, error) 
 GetApplicationAuthenticationMethod(ctx context.Context, params *GetApplicationAuthenticationMethodInput, optFns ...func(*Options)) (*GetApplicationAuthenticationMethodOutput, error) 
 GetApplicationGrant(ctx context.Context, params *GetApplicationGrantInput, optFns ...func(*Options)) (*GetApplicationGrantOutput, error) 
 GetInlinePolicyForPermissionSet(ctx context.Context, params *GetInlinePolicyForPermissionSetInput, optFns ...func(*Options)) (*GetInlinePolicyForPermissionSetOutput, error) 
 GetPermissionsBoundaryForPermissionSet(ctx context.Context, params *GetPermissionsBoundaryForPermissionSetInput, optFns ...func(*Options)) (*GetPermissionsBoundaryForPermissionSetOutput, error) 
 ListAccountAssignmentCreationStatus(ctx context.Context, params *ListAccountAssignmentCreationStatusInput, optFns ...func(*Options)) (*ListAccountAssignmentCreationStatusOutput, error) 
 ListAccountAssignmentDeletionStatus(ctx context.Context, params *ListAccountAssignmentDeletionStatusInput, optFns ...func(*Options)) (*ListAccountAssignmentDeletionStatusOutput, error) 
 ListAccountAssignments(ctx context.Context, params *ListAccountAssignmentsInput, optFns ...func(*Options)) (*ListAccountAssignmentsOutput, error) 
 ListAccountAssignmentsForPrincipal(ctx context.Context, params *ListAccountAssignmentsForPrincipalInput, optFns ...func(*Options)) (*ListAccountAssignmentsForPrincipalOutput, error) 
 ListAccountsForProvisionedPermissionSet(ctx context.Context, params *ListAccountsForProvisionedPermissionSetInput, optFns ...func(*Options)) (*ListAccountsForProvisionedPermissionSetOutput, error) 
 ListApplicationAccessScopes(ctx context.Context, params *ListApplicationAccessScopesInput, optFns ...func(*Options)) (*ListApplicationAccessScopesOutput, error) 
 ListApplicationAssignments(ctx context.Context, params *ListApplicationAssignmentsInput, optFns ...func(*Options)) (*ListApplicationAssignmentsOutput, error) 
 ListApplicationAssignmentsForPrincipal(ctx context.Context, params *ListApplicationAssignmentsForPrincipalInput, optFns ...func(*Options)) (*ListApplicationAssignmentsForPrincipalOutput, error) 
 ListApplicationAuthenticationMethods(ctx context.Context, params *ListApplicationAuthenticationMethodsInput, optFns ...func(*Options)) (*ListApplicationAuthenticationMethodsOutput, error) 
 ListApplicationGrants(ctx context.Context, params *ListApplicationGrantsInput, optFns ...func(*Options)) (*ListApplicationGrantsOutput, error) 
 ListApplicationProviders(ctx context.Context, params *ListApplicationProvidersInput, optFns ...func(*Options)) (*ListApplicationProvidersOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListCustomerManagedPolicyReferencesInPermissionSet(ctx context.Context, params *ListCustomerManagedPolicyReferencesInPermissionSetInput, optFns ...func(*Options)) (*ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) 
 ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) 
 ListManagedPoliciesInPermissionSet(ctx context.Context, params *ListManagedPoliciesInPermissionSetInput, optFns ...func(*Options)) (*ListManagedPoliciesInPermissionSetOutput, error) 
 ListPermissionSetProvisioningStatus(ctx context.Context, params *ListPermissionSetProvisioningStatusInput, optFns ...func(*Options)) (*ListPermissionSetProvisioningStatusOutput, error) 
 ListPermissionSets(ctx context.Context, params *ListPermissionSetsInput, optFns ...func(*Options)) (*ListPermissionSetsOutput, error) 
 ListPermissionSetsProvisionedToAccount(ctx context.Context, params *ListPermissionSetsProvisionedToAccountInput, optFns ...func(*Options)) (*ListPermissionSetsProvisionedToAccountOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrustedTokenIssuers(ctx context.Context, params *ListTrustedTokenIssuersInput, optFns ...func(*Options)) (*ListTrustedTokenIssuersOutput, error) 
 ProvisionPermissionSet(ctx context.Context, params *ProvisionPermissionSetInput, optFns ...func(*Options)) (*ProvisionPermissionSetOutput, error) 
 PutApplicationAccessScope(ctx context.Context, params *PutApplicationAccessScopeInput, optFns ...func(*Options)) (*PutApplicationAccessScopeOutput, error) 
 PutApplicationAssignmentConfiguration(ctx context.Context, params *PutApplicationAssignmentConfigurationInput, optFns ...func(*Options)) (*PutApplicationAssignmentConfigurationOutput, error) 
 PutApplicationAuthenticationMethod(ctx context.Context, params *PutApplicationAuthenticationMethodInput, optFns ...func(*Options)) (*PutApplicationAuthenticationMethodOutput, error) 
 PutApplicationGrant(ctx context.Context, params *PutApplicationGrantInput, optFns ...func(*Options)) (*PutApplicationGrantOutput, error) 
 PutInlinePolicyToPermissionSet(ctx context.Context, params *PutInlinePolicyToPermissionSetInput, optFns ...func(*Options)) (*PutInlinePolicyToPermissionSetOutput, error) 
 PutPermissionsBoundaryToPermissionSet(ctx context.Context, params *PutPermissionsBoundaryToPermissionSetInput, optFns ...func(*Options)) (*PutPermissionsBoundaryToPermissionSetOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateInstance(ctx context.Context, params *UpdateInstanceInput, optFns ...func(*Options)) (*UpdateInstanceOutput, error) 
 UpdateInstanceAccessControlAttributeConfiguration(ctx context.Context, params *UpdateInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*UpdateInstanceAccessControlAttributeConfigurationOutput, error) 
 UpdatePermissionSet(ctx context.Context, params *UpdatePermissionSetInput, optFns ...func(*Options)) (*UpdatePermissionSetOutput, error) 
 UpdateTrustedTokenIssuer(ctx context.Context, params *UpdateTrustedTokenIssuerInput, optFns ...func(*Options)) (*UpdateTrustedTokenIssuerOutput, error) 
}
