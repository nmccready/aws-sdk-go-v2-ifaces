package ssoadmin_test

// tests for the ssoadmin service interface for this ../../../service/ssoadmin/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssoadmin/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssoadmin/ssoadmin_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsoadminServiceCanBeMocked(t *testing.T) {
	var iface ssoadmin_iface.IClient
	iface = &ssoadmin.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssoadmin.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachCustomerManagedPolicyReferenceToPermissionSet", func(t *testing.T) {
        input := &ssoadmin.AttachCustomerManagedPolicyReferenceToPermissionSetInput{}
        output := &ssoadmin.AttachCustomerManagedPolicyReferenceToPermissionSetOutput{}

        mockClient.On("AttachCustomerManagedPolicyReferenceToPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.AttachCustomerManagedPolicyReferenceToPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachManagedPolicyToPermissionSet", func(t *testing.T) {
        input := &ssoadmin.AttachManagedPolicyToPermissionSetInput{}
        output := &ssoadmin.AttachManagedPolicyToPermissionSetOutput{}

        mockClient.On("AttachManagedPolicyToPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.AttachManagedPolicyToPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccountAssignment", func(t *testing.T) {
        input := &ssoadmin.CreateAccountAssignmentInput{}
        output := &ssoadmin.CreateAccountAssignmentOutput{}

        mockClient.On("CreateAccountAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccountAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &ssoadmin.CreateApplicationInput{}
        output := &ssoadmin.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationAssignment", func(t *testing.T) {
        input := &ssoadmin.CreateApplicationAssignmentInput{}
        output := &ssoadmin.CreateApplicationAssignmentOutput{}

        mockClient.On("CreateApplicationAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstance", func(t *testing.T) {
        input := &ssoadmin.CreateInstanceInput{}
        output := &ssoadmin.CreateInstanceOutput{}

        mockClient.On("CreateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceAccessControlAttributeConfiguration", func(t *testing.T) {
        input := &ssoadmin.CreateInstanceAccessControlAttributeConfigurationInput{}
        output := &ssoadmin.CreateInstanceAccessControlAttributeConfigurationOutput{}

        mockClient.On("CreateInstanceAccessControlAttributeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceAccessControlAttributeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePermissionSet", func(t *testing.T) {
        input := &ssoadmin.CreatePermissionSetInput{}
        output := &ssoadmin.CreatePermissionSetOutput{}

        mockClient.On("CreatePermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrustedTokenIssuer", func(t *testing.T) {
        input := &ssoadmin.CreateTrustedTokenIssuerInput{}
        output := &ssoadmin.CreateTrustedTokenIssuerOutput{}

        mockClient.On("CreateTrustedTokenIssuer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrustedTokenIssuer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountAssignment", func(t *testing.T) {
        input := &ssoadmin.DeleteAccountAssignmentInput{}
        output := &ssoadmin.DeleteAccountAssignmentOutput{}

        mockClient.On("DeleteAccountAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &ssoadmin.DeleteApplicationInput{}
        output := &ssoadmin.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationAccessScope", func(t *testing.T) {
        input := &ssoadmin.DeleteApplicationAccessScopeInput{}
        output := &ssoadmin.DeleteApplicationAccessScopeOutput{}

        mockClient.On("DeleteApplicationAccessScope", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationAccessScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationAssignment", func(t *testing.T) {
        input := &ssoadmin.DeleteApplicationAssignmentInput{}
        output := &ssoadmin.DeleteApplicationAssignmentOutput{}

        mockClient.On("DeleteApplicationAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationAuthenticationMethod", func(t *testing.T) {
        input := &ssoadmin.DeleteApplicationAuthenticationMethodInput{}
        output := &ssoadmin.DeleteApplicationAuthenticationMethodOutput{}

        mockClient.On("DeleteApplicationAuthenticationMethod", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationAuthenticationMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationGrant", func(t *testing.T) {
        input := &ssoadmin.DeleteApplicationGrantInput{}
        output := &ssoadmin.DeleteApplicationGrantOutput{}

        mockClient.On("DeleteApplicationGrant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInlinePolicyFromPermissionSet", func(t *testing.T) {
        input := &ssoadmin.DeleteInlinePolicyFromPermissionSetInput{}
        output := &ssoadmin.DeleteInlinePolicyFromPermissionSetOutput{}

        mockClient.On("DeleteInlinePolicyFromPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInlinePolicyFromPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstance", func(t *testing.T) {
        input := &ssoadmin.DeleteInstanceInput{}
        output := &ssoadmin.DeleteInstanceOutput{}

        mockClient.On("DeleteInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceAccessControlAttributeConfiguration", func(t *testing.T) {
        input := &ssoadmin.DeleteInstanceAccessControlAttributeConfigurationInput{}
        output := &ssoadmin.DeleteInstanceAccessControlAttributeConfigurationOutput{}

        mockClient.On("DeleteInstanceAccessControlAttributeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceAccessControlAttributeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionSet", func(t *testing.T) {
        input := &ssoadmin.DeletePermissionSetInput{}
        output := &ssoadmin.DeletePermissionSetOutput{}

        mockClient.On("DeletePermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionsBoundaryFromPermissionSet", func(t *testing.T) {
        input := &ssoadmin.DeletePermissionsBoundaryFromPermissionSetInput{}
        output := &ssoadmin.DeletePermissionsBoundaryFromPermissionSetOutput{}

        mockClient.On("DeletePermissionsBoundaryFromPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionsBoundaryFromPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrustedTokenIssuer", func(t *testing.T) {
        input := &ssoadmin.DeleteTrustedTokenIssuerInput{}
        output := &ssoadmin.DeleteTrustedTokenIssuerOutput{}

        mockClient.On("DeleteTrustedTokenIssuer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrustedTokenIssuer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAssignmentCreationStatus", func(t *testing.T) {
        input := &ssoadmin.DescribeAccountAssignmentCreationStatusInput{}
        output := &ssoadmin.DescribeAccountAssignmentCreationStatusOutput{}

        mockClient.On("DescribeAccountAssignmentCreationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAssignmentCreationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAssignmentDeletionStatus", func(t *testing.T) {
        input := &ssoadmin.DescribeAccountAssignmentDeletionStatusInput{}
        output := &ssoadmin.DescribeAccountAssignmentDeletionStatusOutput{}

        mockClient.On("DescribeAccountAssignmentDeletionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAssignmentDeletionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplication", func(t *testing.T) {
        input := &ssoadmin.DescribeApplicationInput{}
        output := &ssoadmin.DescribeApplicationOutput{}

        mockClient.On("DescribeApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationAssignment", func(t *testing.T) {
        input := &ssoadmin.DescribeApplicationAssignmentInput{}
        output := &ssoadmin.DescribeApplicationAssignmentOutput{}

        mockClient.On("DescribeApplicationAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationProvider", func(t *testing.T) {
        input := &ssoadmin.DescribeApplicationProviderInput{}
        output := &ssoadmin.DescribeApplicationProviderOutput{}

        mockClient.On("DescribeApplicationProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstance", func(t *testing.T) {
        input := &ssoadmin.DescribeInstanceInput{}
        output := &ssoadmin.DescribeInstanceOutput{}

        mockClient.On("DescribeInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceAccessControlAttributeConfiguration", func(t *testing.T) {
        input := &ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput{}
        output := &ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput{}

        mockClient.On("DescribeInstanceAccessControlAttributeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceAccessControlAttributeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePermissionSet", func(t *testing.T) {
        input := &ssoadmin.DescribePermissionSetInput{}
        output := &ssoadmin.DescribePermissionSetOutput{}

        mockClient.On("DescribePermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePermissionSetProvisioningStatus", func(t *testing.T) {
        input := &ssoadmin.DescribePermissionSetProvisioningStatusInput{}
        output := &ssoadmin.DescribePermissionSetProvisioningStatusOutput{}

        mockClient.On("DescribePermissionSetProvisioningStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePermissionSetProvisioningStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustedTokenIssuer", func(t *testing.T) {
        input := &ssoadmin.DescribeTrustedTokenIssuerInput{}
        output := &ssoadmin.DescribeTrustedTokenIssuerOutput{}

        mockClient.On("DescribeTrustedTokenIssuer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustedTokenIssuer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachCustomerManagedPolicyReferenceFromPermissionSet", func(t *testing.T) {
        input := &ssoadmin.DetachCustomerManagedPolicyReferenceFromPermissionSetInput{}
        output := &ssoadmin.DetachCustomerManagedPolicyReferenceFromPermissionSetOutput{}

        mockClient.On("DetachCustomerManagedPolicyReferenceFromPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DetachCustomerManagedPolicyReferenceFromPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachManagedPolicyFromPermissionSet", func(t *testing.T) {
        input := &ssoadmin.DetachManagedPolicyFromPermissionSetInput{}
        output := &ssoadmin.DetachManagedPolicyFromPermissionSetOutput{}

        mockClient.On("DetachManagedPolicyFromPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.DetachManagedPolicyFromPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationAccessScope", func(t *testing.T) {
        input := &ssoadmin.GetApplicationAccessScopeInput{}
        output := &ssoadmin.GetApplicationAccessScopeOutput{}

        mockClient.On("GetApplicationAccessScope", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationAccessScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationAssignmentConfiguration", func(t *testing.T) {
        input := &ssoadmin.GetApplicationAssignmentConfigurationInput{}
        output := &ssoadmin.GetApplicationAssignmentConfigurationOutput{}

        mockClient.On("GetApplicationAssignmentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationAssignmentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationAuthenticationMethod", func(t *testing.T) {
        input := &ssoadmin.GetApplicationAuthenticationMethodInput{}
        output := &ssoadmin.GetApplicationAuthenticationMethodOutput{}

        mockClient.On("GetApplicationAuthenticationMethod", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationAuthenticationMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationGrant", func(t *testing.T) {
        input := &ssoadmin.GetApplicationGrantInput{}
        output := &ssoadmin.GetApplicationGrantOutput{}

        mockClient.On("GetApplicationGrant", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInlinePolicyForPermissionSet", func(t *testing.T) {
        input := &ssoadmin.GetInlinePolicyForPermissionSetInput{}
        output := &ssoadmin.GetInlinePolicyForPermissionSetOutput{}

        mockClient.On("GetInlinePolicyForPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetInlinePolicyForPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermissionsBoundaryForPermissionSet", func(t *testing.T) {
        input := &ssoadmin.GetPermissionsBoundaryForPermissionSetInput{}
        output := &ssoadmin.GetPermissionsBoundaryForPermissionSetOutput{}

        mockClient.On("GetPermissionsBoundaryForPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermissionsBoundaryForPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAssignmentCreationStatus", func(t *testing.T) {
        input := &ssoadmin.ListAccountAssignmentCreationStatusInput{}
        output := &ssoadmin.ListAccountAssignmentCreationStatusOutput{}

        mockClient.On("ListAccountAssignmentCreationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAssignmentCreationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAssignmentDeletionStatus", func(t *testing.T) {
        input := &ssoadmin.ListAccountAssignmentDeletionStatusInput{}
        output := &ssoadmin.ListAccountAssignmentDeletionStatusOutput{}

        mockClient.On("ListAccountAssignmentDeletionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAssignmentDeletionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAssignments", func(t *testing.T) {
        input := &ssoadmin.ListAccountAssignmentsInput{}
        output := &ssoadmin.ListAccountAssignmentsOutput{}

        mockClient.On("ListAccountAssignments", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAssignments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAssignmentsForPrincipal", func(t *testing.T) {
        input := &ssoadmin.ListAccountAssignmentsForPrincipalInput{}
        output := &ssoadmin.ListAccountAssignmentsForPrincipalOutput{}

        mockClient.On("ListAccountAssignmentsForPrincipal", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAssignmentsForPrincipal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountsForProvisionedPermissionSet", func(t *testing.T) {
        input := &ssoadmin.ListAccountsForProvisionedPermissionSetInput{}
        output := &ssoadmin.ListAccountsForProvisionedPermissionSetOutput{}

        mockClient.On("ListAccountsForProvisionedPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountsForProvisionedPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationAccessScopes", func(t *testing.T) {
        input := &ssoadmin.ListApplicationAccessScopesInput{}
        output := &ssoadmin.ListApplicationAccessScopesOutput{}

        mockClient.On("ListApplicationAccessScopes", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationAccessScopes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationAssignments", func(t *testing.T) {
        input := &ssoadmin.ListApplicationAssignmentsInput{}
        output := &ssoadmin.ListApplicationAssignmentsOutput{}

        mockClient.On("ListApplicationAssignments", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationAssignments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationAssignmentsForPrincipal", func(t *testing.T) {
        input := &ssoadmin.ListApplicationAssignmentsForPrincipalInput{}
        output := &ssoadmin.ListApplicationAssignmentsForPrincipalOutput{}

        mockClient.On("ListApplicationAssignmentsForPrincipal", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationAssignmentsForPrincipal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationAuthenticationMethods", func(t *testing.T) {
        input := &ssoadmin.ListApplicationAuthenticationMethodsInput{}
        output := &ssoadmin.ListApplicationAuthenticationMethodsOutput{}

        mockClient.On("ListApplicationAuthenticationMethods", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationAuthenticationMethods(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationGrants", func(t *testing.T) {
        input := &ssoadmin.ListApplicationGrantsInput{}
        output := &ssoadmin.ListApplicationGrantsOutput{}

        mockClient.On("ListApplicationGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationProviders", func(t *testing.T) {
        input := &ssoadmin.ListApplicationProvidersInput{}
        output := &ssoadmin.ListApplicationProvidersOutput{}

        mockClient.On("ListApplicationProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &ssoadmin.ListApplicationsInput{}
        output := &ssoadmin.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomerManagedPolicyReferencesInPermissionSet", func(t *testing.T) {
        input := &ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput{}
        output := &ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput{}

        mockClient.On("ListCustomerManagedPolicyReferencesInPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomerManagedPolicyReferencesInPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstances", func(t *testing.T) {
        input := &ssoadmin.ListInstancesInput{}
        output := &ssoadmin.ListInstancesOutput{}

        mockClient.On("ListInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedPoliciesInPermissionSet", func(t *testing.T) {
        input := &ssoadmin.ListManagedPoliciesInPermissionSetInput{}
        output := &ssoadmin.ListManagedPoliciesInPermissionSetOutput{}

        mockClient.On("ListManagedPoliciesInPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedPoliciesInPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionSetProvisioningStatus", func(t *testing.T) {
        input := &ssoadmin.ListPermissionSetProvisioningStatusInput{}
        output := &ssoadmin.ListPermissionSetProvisioningStatusOutput{}

        mockClient.On("ListPermissionSetProvisioningStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionSetProvisioningStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionSets", func(t *testing.T) {
        input := &ssoadmin.ListPermissionSetsInput{}
        output := &ssoadmin.ListPermissionSetsOutput{}

        mockClient.On("ListPermissionSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionSetsProvisionedToAccount", func(t *testing.T) {
        input := &ssoadmin.ListPermissionSetsProvisionedToAccountInput{}
        output := &ssoadmin.ListPermissionSetsProvisionedToAccountOutput{}

        mockClient.On("ListPermissionSetsProvisionedToAccount", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionSetsProvisionedToAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ssoadmin.ListTagsForResourceInput{}
        output := &ssoadmin.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrustedTokenIssuers", func(t *testing.T) {
        input := &ssoadmin.ListTrustedTokenIssuersInput{}
        output := &ssoadmin.ListTrustedTokenIssuersOutput{}

        mockClient.On("ListTrustedTokenIssuers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrustedTokenIssuers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionPermissionSet", func(t *testing.T) {
        input := &ssoadmin.ProvisionPermissionSetInput{}
        output := &ssoadmin.ProvisionPermissionSetOutput{}

        mockClient.On("ProvisionPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApplicationAccessScope", func(t *testing.T) {
        input := &ssoadmin.PutApplicationAccessScopeInput{}
        output := &ssoadmin.PutApplicationAccessScopeOutput{}

        mockClient.On("PutApplicationAccessScope", ctx, input).Return(output, nil)

        result, err := mockClient.PutApplicationAccessScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApplicationAssignmentConfiguration", func(t *testing.T) {
        input := &ssoadmin.PutApplicationAssignmentConfigurationInput{}
        output := &ssoadmin.PutApplicationAssignmentConfigurationOutput{}

        mockClient.On("PutApplicationAssignmentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutApplicationAssignmentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApplicationAuthenticationMethod", func(t *testing.T) {
        input := &ssoadmin.PutApplicationAuthenticationMethodInput{}
        output := &ssoadmin.PutApplicationAuthenticationMethodOutput{}

        mockClient.On("PutApplicationAuthenticationMethod", ctx, input).Return(output, nil)

        result, err := mockClient.PutApplicationAuthenticationMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApplicationGrant", func(t *testing.T) {
        input := &ssoadmin.PutApplicationGrantInput{}
        output := &ssoadmin.PutApplicationGrantOutput{}

        mockClient.On("PutApplicationGrant", ctx, input).Return(output, nil)

        result, err := mockClient.PutApplicationGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInlinePolicyToPermissionSet", func(t *testing.T) {
        input := &ssoadmin.PutInlinePolicyToPermissionSetInput{}
        output := &ssoadmin.PutInlinePolicyToPermissionSetOutput{}

        mockClient.On("PutInlinePolicyToPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.PutInlinePolicyToPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermissionsBoundaryToPermissionSet", func(t *testing.T) {
        input := &ssoadmin.PutPermissionsBoundaryToPermissionSetInput{}
        output := &ssoadmin.PutPermissionsBoundaryToPermissionSetOutput{}

        mockClient.On("PutPermissionsBoundaryToPermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermissionsBoundaryToPermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ssoadmin.TagResourceInput{}
        output := &ssoadmin.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ssoadmin.UntagResourceInput{}
        output := &ssoadmin.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &ssoadmin.UpdateApplicationInput{}
        output := &ssoadmin.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstance", func(t *testing.T) {
        input := &ssoadmin.UpdateInstanceInput{}
        output := &ssoadmin.UpdateInstanceOutput{}

        mockClient.On("UpdateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceAccessControlAttributeConfiguration", func(t *testing.T) {
        input := &ssoadmin.UpdateInstanceAccessControlAttributeConfigurationInput{}
        output := &ssoadmin.UpdateInstanceAccessControlAttributeConfigurationOutput{}

        mockClient.On("UpdateInstanceAccessControlAttributeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceAccessControlAttributeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePermissionSet", func(t *testing.T) {
        input := &ssoadmin.UpdatePermissionSetInput{}
        output := &ssoadmin.UpdatePermissionSetOutput{}

        mockClient.On("UpdatePermissionSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePermissionSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrustedTokenIssuer", func(t *testing.T) {
        input := &ssoadmin.UpdateTrustedTokenIssuerInput{}
        output := &ssoadmin.UpdateTrustedTokenIssuerOutput{}

        mockClient.On("UpdateTrustedTokenIssuer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrustedTokenIssuer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
