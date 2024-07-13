package iam_test

// tests for the iam service interface for this ../../../service/iam/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iam/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iam/iam_iface"
	"github.com/stretchr/testify/assert"
)

func TestIamServiceCanBeMocked(t *testing.T) {
	var iface iam_iface.IClient
	iface = &iam.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iam.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddClientIDToOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.AddClientIDToOpenIDConnectProviderInput{}
        output := &iam.AddClientIDToOpenIDConnectProviderOutput{}

        mockClient.On("AddClientIDToOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.AddClientIDToOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddRoleToInstanceProfile", func(t *testing.T) {
        input := &iam.AddRoleToInstanceProfileInput{}
        output := &iam.AddRoleToInstanceProfileOutput{}

        mockClient.On("AddRoleToInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.AddRoleToInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddUserToGroup", func(t *testing.T) {
        input := &iam.AddUserToGroupInput{}
        output := &iam.AddUserToGroupOutput{}

        mockClient.On("AddUserToGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AddUserToGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachGroupPolicy", func(t *testing.T) {
        input := &iam.AttachGroupPolicyInput{}
        output := &iam.AttachGroupPolicyOutput{}

        mockClient.On("AttachGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachRolePolicy", func(t *testing.T) {
        input := &iam.AttachRolePolicyInput{}
        output := &iam.AttachRolePolicyOutput{}

        mockClient.On("AttachRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachUserPolicy", func(t *testing.T) {
        input := &iam.AttachUserPolicyInput{}
        output := &iam.AttachUserPolicyOutput{}

        mockClient.On("AttachUserPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachUserPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangePassword", func(t *testing.T) {
        input := &iam.ChangePasswordInput{}
        output := &iam.ChangePasswordOutput{}

        mockClient.On("ChangePassword", ctx, input).Return(output, nil)

        result, err := mockClient.ChangePassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessKey", func(t *testing.T) {
        input := &iam.CreateAccessKeyInput{}
        output := &iam.CreateAccessKeyOutput{}

        mockClient.On("CreateAccessKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccountAlias", func(t *testing.T) {
        input := &iam.CreateAccountAliasInput{}
        output := &iam.CreateAccountAliasOutput{}

        mockClient.On("CreateAccountAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccountAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &iam.CreateGroupInput{}
        output := &iam.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceProfile", func(t *testing.T) {
        input := &iam.CreateInstanceProfileInput{}
        output := &iam.CreateInstanceProfileOutput{}

        mockClient.On("CreateInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoginProfile", func(t *testing.T) {
        input := &iam.CreateLoginProfileInput{}
        output := &iam.CreateLoginProfileOutput{}

        mockClient.On("CreateLoginProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoginProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.CreateOpenIDConnectProviderInput{}
        output := &iam.CreateOpenIDConnectProviderOutput{}

        mockClient.On("CreateOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicy", func(t *testing.T) {
        input := &iam.CreatePolicyInput{}
        output := &iam.CreatePolicyOutput{}

        mockClient.On("CreatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicyVersion", func(t *testing.T) {
        input := &iam.CreatePolicyVersionInput{}
        output := &iam.CreatePolicyVersionOutput{}

        mockClient.On("CreatePolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRole", func(t *testing.T) {
        input := &iam.CreateRoleInput{}
        output := &iam.CreateRoleOutput{}

        mockClient.On("CreateRole", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSAMLProvider", func(t *testing.T) {
        input := &iam.CreateSAMLProviderInput{}
        output := &iam.CreateSAMLProviderOutput{}

        mockClient.On("CreateSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceLinkedRole", func(t *testing.T) {
        input := &iam.CreateServiceLinkedRoleInput{}
        output := &iam.CreateServiceLinkedRoleOutput{}

        mockClient.On("CreateServiceLinkedRole", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceLinkedRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceSpecificCredential", func(t *testing.T) {
        input := &iam.CreateServiceSpecificCredentialInput{}
        output := &iam.CreateServiceSpecificCredentialOutput{}

        mockClient.On("CreateServiceSpecificCredential", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceSpecificCredential(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &iam.CreateUserInput{}
        output := &iam.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualMFADevice", func(t *testing.T) {
        input := &iam.CreateVirtualMFADeviceInput{}
        output := &iam.CreateVirtualMFADeviceOutput{}

        mockClient.On("CreateVirtualMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateMFADevice", func(t *testing.T) {
        input := &iam.DeactivateMFADeviceInput{}
        output := &iam.DeactivateMFADeviceOutput{}

        mockClient.On("DeactivateMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessKey", func(t *testing.T) {
        input := &iam.DeleteAccessKeyInput{}
        output := &iam.DeleteAccessKeyOutput{}

        mockClient.On("DeleteAccessKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountAlias", func(t *testing.T) {
        input := &iam.DeleteAccountAliasInput{}
        output := &iam.DeleteAccountAliasOutput{}

        mockClient.On("DeleteAccountAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountPasswordPolicy", func(t *testing.T) {
        input := &iam.DeleteAccountPasswordPolicyInput{}
        output := &iam.DeleteAccountPasswordPolicyOutput{}

        mockClient.On("DeleteAccountPasswordPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountPasswordPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &iam.DeleteGroupInput{}
        output := &iam.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroupPolicy", func(t *testing.T) {
        input := &iam.DeleteGroupPolicyInput{}
        output := &iam.DeleteGroupPolicyOutput{}

        mockClient.On("DeleteGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceProfile", func(t *testing.T) {
        input := &iam.DeleteInstanceProfileInput{}
        output := &iam.DeleteInstanceProfileOutput{}

        mockClient.On("DeleteInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoginProfile", func(t *testing.T) {
        input := &iam.DeleteLoginProfileInput{}
        output := &iam.DeleteLoginProfileOutput{}

        mockClient.On("DeleteLoginProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoginProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.DeleteOpenIDConnectProviderInput{}
        output := &iam.DeleteOpenIDConnectProviderOutput{}

        mockClient.On("DeleteOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &iam.DeletePolicyInput{}
        output := &iam.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicyVersion", func(t *testing.T) {
        input := &iam.DeletePolicyVersionInput{}
        output := &iam.DeletePolicyVersionOutput{}

        mockClient.On("DeletePolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRole", func(t *testing.T) {
        input := &iam.DeleteRoleInput{}
        output := &iam.DeleteRoleOutput{}

        mockClient.On("DeleteRole", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRolePermissionsBoundary", func(t *testing.T) {
        input := &iam.DeleteRolePermissionsBoundaryInput{}
        output := &iam.DeleteRolePermissionsBoundaryOutput{}

        mockClient.On("DeleteRolePermissionsBoundary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRolePermissionsBoundary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRolePolicy", func(t *testing.T) {
        input := &iam.DeleteRolePolicyInput{}
        output := &iam.DeleteRolePolicyOutput{}

        mockClient.On("DeleteRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSAMLProvider", func(t *testing.T) {
        input := &iam.DeleteSAMLProviderInput{}
        output := &iam.DeleteSAMLProviderOutput{}

        mockClient.On("DeleteSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSSHPublicKey", func(t *testing.T) {
        input := &iam.DeleteSSHPublicKeyInput{}
        output := &iam.DeleteSSHPublicKeyOutput{}

        mockClient.On("DeleteSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServerCertificate", func(t *testing.T) {
        input := &iam.DeleteServerCertificateInput{}
        output := &iam.DeleteServerCertificateOutput{}

        mockClient.On("DeleteServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceLinkedRole", func(t *testing.T) {
        input := &iam.DeleteServiceLinkedRoleInput{}
        output := &iam.DeleteServiceLinkedRoleOutput{}

        mockClient.On("DeleteServiceLinkedRole", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceLinkedRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceSpecificCredential", func(t *testing.T) {
        input := &iam.DeleteServiceSpecificCredentialInput{}
        output := &iam.DeleteServiceSpecificCredentialOutput{}

        mockClient.On("DeleteServiceSpecificCredential", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceSpecificCredential(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSigningCertificate", func(t *testing.T) {
        input := &iam.DeleteSigningCertificateInput{}
        output := &iam.DeleteSigningCertificateOutput{}

        mockClient.On("DeleteSigningCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSigningCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &iam.DeleteUserInput{}
        output := &iam.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserPermissionsBoundary", func(t *testing.T) {
        input := &iam.DeleteUserPermissionsBoundaryInput{}
        output := &iam.DeleteUserPermissionsBoundaryOutput{}

        mockClient.On("DeleteUserPermissionsBoundary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserPermissionsBoundary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserPolicy", func(t *testing.T) {
        input := &iam.DeleteUserPolicyInput{}
        output := &iam.DeleteUserPolicyOutput{}

        mockClient.On("DeleteUserPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualMFADevice", func(t *testing.T) {
        input := &iam.DeleteVirtualMFADeviceInput{}
        output := &iam.DeleteVirtualMFADeviceOutput{}

        mockClient.On("DeleteVirtualMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachGroupPolicy", func(t *testing.T) {
        input := &iam.DetachGroupPolicyInput{}
        output := &iam.DetachGroupPolicyOutput{}

        mockClient.On("DetachGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachRolePolicy", func(t *testing.T) {
        input := &iam.DetachRolePolicyInput{}
        output := &iam.DetachRolePolicyOutput{}

        mockClient.On("DetachRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachUserPolicy", func(t *testing.T) {
        input := &iam.DetachUserPolicyInput{}
        output := &iam.DetachUserPolicyOutput{}

        mockClient.On("DetachUserPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachUserPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableMFADevice", func(t *testing.T) {
        input := &iam.EnableMFADeviceInput{}
        output := &iam.EnableMFADeviceOutput{}

        mockClient.On("EnableMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.EnableMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateCredentialReport", func(t *testing.T) {
        input := &iam.GenerateCredentialReportInput{}
        output := &iam.GenerateCredentialReportOutput{}

        mockClient.On("GenerateCredentialReport", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateCredentialReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateOrganizationsAccessReport", func(t *testing.T) {
        input := &iam.GenerateOrganizationsAccessReportInput{}
        output := &iam.GenerateOrganizationsAccessReportOutput{}

        mockClient.On("GenerateOrganizationsAccessReport", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateOrganizationsAccessReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateServiceLastAccessedDetails", func(t *testing.T) {
        input := &iam.GenerateServiceLastAccessedDetailsInput{}
        output := &iam.GenerateServiceLastAccessedDetailsOutput{}

        mockClient.On("GenerateServiceLastAccessedDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateServiceLastAccessedDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessKeyLastUsed", func(t *testing.T) {
        input := &iam.GetAccessKeyLastUsedInput{}
        output := &iam.GetAccessKeyLastUsedOutput{}

        mockClient.On("GetAccessKeyLastUsed", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessKeyLastUsed(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountAuthorizationDetails", func(t *testing.T) {
        input := &iam.GetAccountAuthorizationDetailsInput{}
        output := &iam.GetAccountAuthorizationDetailsOutput{}

        mockClient.On("GetAccountAuthorizationDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountAuthorizationDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountPasswordPolicy", func(t *testing.T) {
        input := &iam.GetAccountPasswordPolicyInput{}
        output := &iam.GetAccountPasswordPolicyOutput{}

        mockClient.On("GetAccountPasswordPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountPasswordPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSummary", func(t *testing.T) {
        input := &iam.GetAccountSummaryInput{}
        output := &iam.GetAccountSummaryOutput{}

        mockClient.On("GetAccountSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContextKeysForCustomPolicy", func(t *testing.T) {
        input := &iam.GetContextKeysForCustomPolicyInput{}
        output := &iam.GetContextKeysForCustomPolicyOutput{}

        mockClient.On("GetContextKeysForCustomPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContextKeysForCustomPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContextKeysForPrincipalPolicy", func(t *testing.T) {
        input := &iam.GetContextKeysForPrincipalPolicyInput{}
        output := &iam.GetContextKeysForPrincipalPolicyOutput{}

        mockClient.On("GetContextKeysForPrincipalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContextKeysForPrincipalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCredentialReport", func(t *testing.T) {
        input := &iam.GetCredentialReportInput{}
        output := &iam.GetCredentialReportOutput{}

        mockClient.On("GetCredentialReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetCredentialReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &iam.GetGroupInput{}
        output := &iam.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupPolicy", func(t *testing.T) {
        input := &iam.GetGroupPolicyInput{}
        output := &iam.GetGroupPolicyOutput{}

        mockClient.On("GetGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceProfile", func(t *testing.T) {
        input := &iam.GetInstanceProfileInput{}
        output := &iam.GetInstanceProfileOutput{}

        mockClient.On("GetInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoginProfile", func(t *testing.T) {
        input := &iam.GetLoginProfileInput{}
        output := &iam.GetLoginProfileOutput{}

        mockClient.On("GetLoginProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoginProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMFADevice", func(t *testing.T) {
        input := &iam.GetMFADeviceInput{}
        output := &iam.GetMFADeviceOutput{}

        mockClient.On("GetMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.GetOpenIDConnectProviderInput{}
        output := &iam.GetOpenIDConnectProviderOutput{}

        mockClient.On("GetOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationsAccessReport", func(t *testing.T) {
        input := &iam.GetOrganizationsAccessReportInput{}
        output := &iam.GetOrganizationsAccessReportOutput{}

        mockClient.On("GetOrganizationsAccessReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationsAccessReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &iam.GetPolicyInput{}
        output := &iam.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicyVersion", func(t *testing.T) {
        input := &iam.GetPolicyVersionInput{}
        output := &iam.GetPolicyVersionOutput{}

        mockClient.On("GetPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRole", func(t *testing.T) {
        input := &iam.GetRoleInput{}
        output := &iam.GetRoleOutput{}

        mockClient.On("GetRole", ctx, input).Return(output, nil)

        result, err := mockClient.GetRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRolePolicy", func(t *testing.T) {
        input := &iam.GetRolePolicyInput{}
        output := &iam.GetRolePolicyOutput{}

        mockClient.On("GetRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSAMLProvider", func(t *testing.T) {
        input := &iam.GetSAMLProviderInput{}
        output := &iam.GetSAMLProviderOutput{}

        mockClient.On("GetSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.GetSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSSHPublicKey", func(t *testing.T) {
        input := &iam.GetSSHPublicKeyInput{}
        output := &iam.GetSSHPublicKeyOutput{}

        mockClient.On("GetSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServerCertificate", func(t *testing.T) {
        input := &iam.GetServerCertificateInput{}
        output := &iam.GetServerCertificateOutput{}

        mockClient.On("GetServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceLastAccessedDetails", func(t *testing.T) {
        input := &iam.GetServiceLastAccessedDetailsInput{}
        output := &iam.GetServiceLastAccessedDetailsOutput{}

        mockClient.On("GetServiceLastAccessedDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceLastAccessedDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceLastAccessedDetailsWithEntities", func(t *testing.T) {
        input := &iam.GetServiceLastAccessedDetailsWithEntitiesInput{}
        output := &iam.GetServiceLastAccessedDetailsWithEntitiesOutput{}

        mockClient.On("GetServiceLastAccessedDetailsWithEntities", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceLastAccessedDetailsWithEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceLinkedRoleDeletionStatus", func(t *testing.T) {
        input := &iam.GetServiceLinkedRoleDeletionStatusInput{}
        output := &iam.GetServiceLinkedRoleDeletionStatusOutput{}

        mockClient.On("GetServiceLinkedRoleDeletionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceLinkedRoleDeletionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUser", func(t *testing.T) {
        input := &iam.GetUserInput{}
        output := &iam.GetUserOutput{}

        mockClient.On("GetUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserPolicy", func(t *testing.T) {
        input := &iam.GetUserPolicyInput{}
        output := &iam.GetUserPolicyOutput{}

        mockClient.On("GetUserPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessKeys", func(t *testing.T) {
        input := &iam.ListAccessKeysInput{}
        output := &iam.ListAccessKeysOutput{}

        mockClient.On("ListAccessKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAliases", func(t *testing.T) {
        input := &iam.ListAccountAliasesInput{}
        output := &iam.ListAccountAliasesOutput{}

        mockClient.On("ListAccountAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedGroupPolicies", func(t *testing.T) {
        input := &iam.ListAttachedGroupPoliciesInput{}
        output := &iam.ListAttachedGroupPoliciesOutput{}

        mockClient.On("ListAttachedGroupPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedGroupPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedRolePolicies", func(t *testing.T) {
        input := &iam.ListAttachedRolePoliciesInput{}
        output := &iam.ListAttachedRolePoliciesOutput{}

        mockClient.On("ListAttachedRolePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedRolePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedUserPolicies", func(t *testing.T) {
        input := &iam.ListAttachedUserPoliciesInput{}
        output := &iam.ListAttachedUserPoliciesOutput{}

        mockClient.On("ListAttachedUserPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedUserPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntitiesForPolicy", func(t *testing.T) {
        input := &iam.ListEntitiesForPolicyInput{}
        output := &iam.ListEntitiesForPolicyOutput{}

        mockClient.On("ListEntitiesForPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntitiesForPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupPolicies", func(t *testing.T) {
        input := &iam.ListGroupPoliciesInput{}
        output := &iam.ListGroupPoliciesOutput{}

        mockClient.On("ListGroupPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &iam.ListGroupsInput{}
        output := &iam.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupsForUser", func(t *testing.T) {
        input := &iam.ListGroupsForUserInput{}
        output := &iam.ListGroupsForUserOutput{}

        mockClient.On("ListGroupsForUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupsForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceProfileTags", func(t *testing.T) {
        input := &iam.ListInstanceProfileTagsInput{}
        output := &iam.ListInstanceProfileTagsOutput{}

        mockClient.On("ListInstanceProfileTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceProfileTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceProfiles", func(t *testing.T) {
        input := &iam.ListInstanceProfilesInput{}
        output := &iam.ListInstanceProfilesOutput{}

        mockClient.On("ListInstanceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceProfilesForRole", func(t *testing.T) {
        input := &iam.ListInstanceProfilesForRoleInput{}
        output := &iam.ListInstanceProfilesForRoleOutput{}

        mockClient.On("ListInstanceProfilesForRole", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceProfilesForRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMFADeviceTags", func(t *testing.T) {
        input := &iam.ListMFADeviceTagsInput{}
        output := &iam.ListMFADeviceTagsOutput{}

        mockClient.On("ListMFADeviceTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListMFADeviceTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMFADevices", func(t *testing.T) {
        input := &iam.ListMFADevicesInput{}
        output := &iam.ListMFADevicesOutput{}

        mockClient.On("ListMFADevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListMFADevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpenIDConnectProviderTags", func(t *testing.T) {
        input := &iam.ListOpenIDConnectProviderTagsInput{}
        output := &iam.ListOpenIDConnectProviderTagsOutput{}

        mockClient.On("ListOpenIDConnectProviderTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpenIDConnectProviderTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpenIDConnectProviders", func(t *testing.T) {
        input := &iam.ListOpenIDConnectProvidersInput{}
        output := &iam.ListOpenIDConnectProvidersOutput{}

        mockClient.On("ListOpenIDConnectProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpenIDConnectProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicies", func(t *testing.T) {
        input := &iam.ListPoliciesInput{}
        output := &iam.ListPoliciesOutput{}

        mockClient.On("ListPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPoliciesGrantingServiceAccess", func(t *testing.T) {
        input := &iam.ListPoliciesGrantingServiceAccessInput{}
        output := &iam.ListPoliciesGrantingServiceAccessOutput{}

        mockClient.On("ListPoliciesGrantingServiceAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListPoliciesGrantingServiceAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyTags", func(t *testing.T) {
        input := &iam.ListPolicyTagsInput{}
        output := &iam.ListPolicyTagsOutput{}

        mockClient.On("ListPolicyTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyVersions", func(t *testing.T) {
        input := &iam.ListPolicyVersionsInput{}
        output := &iam.ListPolicyVersionsOutput{}

        mockClient.On("ListPolicyVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRolePolicies", func(t *testing.T) {
        input := &iam.ListRolePoliciesInput{}
        output := &iam.ListRolePoliciesOutput{}

        mockClient.On("ListRolePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListRolePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoleTags", func(t *testing.T) {
        input := &iam.ListRoleTagsInput{}
        output := &iam.ListRoleTagsOutput{}

        mockClient.On("ListRoleTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoleTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoles", func(t *testing.T) {
        input := &iam.ListRolesInput{}
        output := &iam.ListRolesOutput{}

        mockClient.On("ListRoles", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSAMLProviderTags", func(t *testing.T) {
        input := &iam.ListSAMLProviderTagsInput{}
        output := &iam.ListSAMLProviderTagsOutput{}

        mockClient.On("ListSAMLProviderTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListSAMLProviderTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSAMLProviders", func(t *testing.T) {
        input := &iam.ListSAMLProvidersInput{}
        output := &iam.ListSAMLProvidersOutput{}

        mockClient.On("ListSAMLProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListSAMLProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSSHPublicKeys", func(t *testing.T) {
        input := &iam.ListSSHPublicKeysInput{}
        output := &iam.ListSSHPublicKeysOutput{}

        mockClient.On("ListSSHPublicKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListSSHPublicKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServerCertificateTags", func(t *testing.T) {
        input := &iam.ListServerCertificateTagsInput{}
        output := &iam.ListServerCertificateTagsOutput{}

        mockClient.On("ListServerCertificateTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListServerCertificateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServerCertificates", func(t *testing.T) {
        input := &iam.ListServerCertificatesInput{}
        output := &iam.ListServerCertificatesOutput{}

        mockClient.On("ListServerCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListServerCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceSpecificCredentials", func(t *testing.T) {
        input := &iam.ListServiceSpecificCredentialsInput{}
        output := &iam.ListServiceSpecificCredentialsOutput{}

        mockClient.On("ListServiceSpecificCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceSpecificCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSigningCertificates", func(t *testing.T) {
        input := &iam.ListSigningCertificatesInput{}
        output := &iam.ListSigningCertificatesOutput{}

        mockClient.On("ListSigningCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListSigningCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserPolicies", func(t *testing.T) {
        input := &iam.ListUserPoliciesInput{}
        output := &iam.ListUserPoliciesOutput{}

        mockClient.On("ListUserPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserTags", func(t *testing.T) {
        input := &iam.ListUserTagsInput{}
        output := &iam.ListUserTagsOutput{}

        mockClient.On("ListUserTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &iam.ListUsersInput{}
        output := &iam.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualMFADevices", func(t *testing.T) {
        input := &iam.ListVirtualMFADevicesInput{}
        output := &iam.ListVirtualMFADevicesOutput{}

        mockClient.On("ListVirtualMFADevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualMFADevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGroupPolicy", func(t *testing.T) {
        input := &iam.PutGroupPolicyInput{}
        output := &iam.PutGroupPolicyOutput{}

        mockClient.On("PutGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRolePermissionsBoundary", func(t *testing.T) {
        input := &iam.PutRolePermissionsBoundaryInput{}
        output := &iam.PutRolePermissionsBoundaryOutput{}

        mockClient.On("PutRolePermissionsBoundary", ctx, input).Return(output, nil)

        result, err := mockClient.PutRolePermissionsBoundary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRolePolicy", func(t *testing.T) {
        input := &iam.PutRolePolicyInput{}
        output := &iam.PutRolePolicyOutput{}

        mockClient.On("PutRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutUserPermissionsBoundary", func(t *testing.T) {
        input := &iam.PutUserPermissionsBoundaryInput{}
        output := &iam.PutUserPermissionsBoundaryOutput{}

        mockClient.On("PutUserPermissionsBoundary", ctx, input).Return(output, nil)

        result, err := mockClient.PutUserPermissionsBoundary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutUserPolicy", func(t *testing.T) {
        input := &iam.PutUserPolicyInput{}
        output := &iam.PutUserPolicyOutput{}

        mockClient.On("PutUserPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutUserPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveClientIDFromOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.RemoveClientIDFromOpenIDConnectProviderInput{}
        output := &iam.RemoveClientIDFromOpenIDConnectProviderOutput{}

        mockClient.On("RemoveClientIDFromOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveClientIDFromOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRoleFromInstanceProfile", func(t *testing.T) {
        input := &iam.RemoveRoleFromInstanceProfileInput{}
        output := &iam.RemoveRoleFromInstanceProfileOutput{}

        mockClient.On("RemoveRoleFromInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRoleFromInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveUserFromGroup", func(t *testing.T) {
        input := &iam.RemoveUserFromGroupInput{}
        output := &iam.RemoveUserFromGroupOutput{}

        mockClient.On("RemoveUserFromGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveUserFromGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetServiceSpecificCredential", func(t *testing.T) {
        input := &iam.ResetServiceSpecificCredentialInput{}
        output := &iam.ResetServiceSpecificCredentialOutput{}

        mockClient.On("ResetServiceSpecificCredential", ctx, input).Return(output, nil)

        result, err := mockClient.ResetServiceSpecificCredential(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResyncMFADevice", func(t *testing.T) {
        input := &iam.ResyncMFADeviceInput{}
        output := &iam.ResyncMFADeviceOutput{}

        mockClient.On("ResyncMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.ResyncMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultPolicyVersion", func(t *testing.T) {
        input := &iam.SetDefaultPolicyVersionInput{}
        output := &iam.SetDefaultPolicyVersionOutput{}

        mockClient.On("SetDefaultPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSecurityTokenServicePreferences", func(t *testing.T) {
        input := &iam.SetSecurityTokenServicePreferencesInput{}
        output := &iam.SetSecurityTokenServicePreferencesOutput{}

        mockClient.On("SetSecurityTokenServicePreferences", ctx, input).Return(output, nil)

        result, err := mockClient.SetSecurityTokenServicePreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSimulateCustomPolicy", func(t *testing.T) {
        input := &iam.SimulateCustomPolicyInput{}
        output := &iam.SimulateCustomPolicyOutput{}

        mockClient.On("SimulateCustomPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SimulateCustomPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSimulatePrincipalPolicy", func(t *testing.T) {
        input := &iam.SimulatePrincipalPolicyInput{}
        output := &iam.SimulatePrincipalPolicyOutput{}

        mockClient.On("SimulatePrincipalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SimulatePrincipalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagInstanceProfile", func(t *testing.T) {
        input := &iam.TagInstanceProfileInput{}
        output := &iam.TagInstanceProfileOutput{}

        mockClient.On("TagInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.TagInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagMFADevice", func(t *testing.T) {
        input := &iam.TagMFADeviceInput{}
        output := &iam.TagMFADeviceOutput{}

        mockClient.On("TagMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.TagMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.TagOpenIDConnectProviderInput{}
        output := &iam.TagOpenIDConnectProviderOutput{}

        mockClient.On("TagOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.TagOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagPolicy", func(t *testing.T) {
        input := &iam.TagPolicyInput{}
        output := &iam.TagPolicyOutput{}

        mockClient.On("TagPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.TagPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagRole", func(t *testing.T) {
        input := &iam.TagRoleInput{}
        output := &iam.TagRoleOutput{}

        mockClient.On("TagRole", ctx, input).Return(output, nil)

        result, err := mockClient.TagRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagSAMLProvider", func(t *testing.T) {
        input := &iam.TagSAMLProviderInput{}
        output := &iam.TagSAMLProviderOutput{}

        mockClient.On("TagSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.TagSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagServerCertificate", func(t *testing.T) {
        input := &iam.TagServerCertificateInput{}
        output := &iam.TagServerCertificateOutput{}

        mockClient.On("TagServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.TagServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagUser", func(t *testing.T) {
        input := &iam.TagUserInput{}
        output := &iam.TagUserOutput{}

        mockClient.On("TagUser", ctx, input).Return(output, nil)

        result, err := mockClient.TagUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagInstanceProfile", func(t *testing.T) {
        input := &iam.UntagInstanceProfileInput{}
        output := &iam.UntagInstanceProfileOutput{}

        mockClient.On("UntagInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UntagInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagMFADevice", func(t *testing.T) {
        input := &iam.UntagMFADeviceInput{}
        output := &iam.UntagMFADeviceOutput{}

        mockClient.On("UntagMFADevice", ctx, input).Return(output, nil)

        result, err := mockClient.UntagMFADevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagOpenIDConnectProvider", func(t *testing.T) {
        input := &iam.UntagOpenIDConnectProviderInput{}
        output := &iam.UntagOpenIDConnectProviderOutput{}

        mockClient.On("UntagOpenIDConnectProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UntagOpenIDConnectProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagPolicy", func(t *testing.T) {
        input := &iam.UntagPolicyInput{}
        output := &iam.UntagPolicyOutput{}

        mockClient.On("UntagPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UntagPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagRole", func(t *testing.T) {
        input := &iam.UntagRoleInput{}
        output := &iam.UntagRoleOutput{}

        mockClient.On("UntagRole", ctx, input).Return(output, nil)

        result, err := mockClient.UntagRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagSAMLProvider", func(t *testing.T) {
        input := &iam.UntagSAMLProviderInput{}
        output := &iam.UntagSAMLProviderOutput{}

        mockClient.On("UntagSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UntagSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagServerCertificate", func(t *testing.T) {
        input := &iam.UntagServerCertificateInput{}
        output := &iam.UntagServerCertificateOutput{}

        mockClient.On("UntagServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UntagServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagUser", func(t *testing.T) {
        input := &iam.UntagUserInput{}
        output := &iam.UntagUserOutput{}

        mockClient.On("UntagUser", ctx, input).Return(output, nil)

        result, err := mockClient.UntagUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessKey", func(t *testing.T) {
        input := &iam.UpdateAccessKeyInput{}
        output := &iam.UpdateAccessKeyOutput{}

        mockClient.On("UpdateAccessKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountPasswordPolicy", func(t *testing.T) {
        input := &iam.UpdateAccountPasswordPolicyInput{}
        output := &iam.UpdateAccountPasswordPolicyOutput{}

        mockClient.On("UpdateAccountPasswordPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountPasswordPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssumeRolePolicy", func(t *testing.T) {
        input := &iam.UpdateAssumeRolePolicyInput{}
        output := &iam.UpdateAssumeRolePolicyOutput{}

        mockClient.On("UpdateAssumeRolePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssumeRolePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &iam.UpdateGroupInput{}
        output := &iam.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoginProfile", func(t *testing.T) {
        input := &iam.UpdateLoginProfileInput{}
        output := &iam.UpdateLoginProfileOutput{}

        mockClient.On("UpdateLoginProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoginProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOpenIDConnectProviderThumbprint", func(t *testing.T) {
        input := &iam.UpdateOpenIDConnectProviderThumbprintInput{}
        output := &iam.UpdateOpenIDConnectProviderThumbprintOutput{}

        mockClient.On("UpdateOpenIDConnectProviderThumbprint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOpenIDConnectProviderThumbprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRole", func(t *testing.T) {
        input := &iam.UpdateRoleInput{}
        output := &iam.UpdateRoleOutput{}

        mockClient.On("UpdateRole", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoleDescription", func(t *testing.T) {
        input := &iam.UpdateRoleDescriptionInput{}
        output := &iam.UpdateRoleDescriptionOutput{}

        mockClient.On("UpdateRoleDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoleDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSAMLProvider", func(t *testing.T) {
        input := &iam.UpdateSAMLProviderInput{}
        output := &iam.UpdateSAMLProviderOutput{}

        mockClient.On("UpdateSAMLProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSAMLProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSSHPublicKey", func(t *testing.T) {
        input := &iam.UpdateSSHPublicKeyInput{}
        output := &iam.UpdateSSHPublicKeyOutput{}

        mockClient.On("UpdateSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServerCertificate", func(t *testing.T) {
        input := &iam.UpdateServerCertificateInput{}
        output := &iam.UpdateServerCertificateOutput{}

        mockClient.On("UpdateServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSpecificCredential", func(t *testing.T) {
        input := &iam.UpdateServiceSpecificCredentialInput{}
        output := &iam.UpdateServiceSpecificCredentialOutput{}

        mockClient.On("UpdateServiceSpecificCredential", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSpecificCredential(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSigningCertificate", func(t *testing.T) {
        input := &iam.UpdateSigningCertificateInput{}
        output := &iam.UpdateSigningCertificateOutput{}

        mockClient.On("UpdateSigningCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSigningCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &iam.UpdateUserInput{}
        output := &iam.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadSSHPublicKey", func(t *testing.T) {
        input := &iam.UploadSSHPublicKeyInput{}
        output := &iam.UploadSSHPublicKeyOutput{}

        mockClient.On("UploadSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.UploadSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadServerCertificate", func(t *testing.T) {
        input := &iam.UploadServerCertificateInput{}
        output := &iam.UploadServerCertificateOutput{}

        mockClient.On("UploadServerCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UploadServerCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadSigningCertificate", func(t *testing.T) {
        input := &iam.UploadSigningCertificateInput{}
        output := &iam.UploadSigningCertificateOutput{}

        mockClient.On("UploadSigningCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UploadSigningCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
