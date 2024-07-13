package cognitoidentityprovider_test

// tests for the cognitoidentityprovider service interface for this ../../../service/cognitoidentityprovider/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitoidentityprovider/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitoidentityprovider/cognitoidentityprovider_iface"
	"github.com/stretchr/testify/assert"
)

func TestCognitoidentityproviderServiceCanBeMocked(t *testing.T) {
	var iface cognitoidentityprovider_iface.IClient
	iface = &cognitoidentityprovider.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cognitoidentityprovider.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddCustomAttributes", func(t *testing.T) {
        input := &cognitoidentityprovider.AddCustomAttributesInput{}
        output := &cognitoidentityprovider.AddCustomAttributesOutput{}

        mockClient.On("AddCustomAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.AddCustomAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminAddUserToGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminAddUserToGroupInput{}
        output := &cognitoidentityprovider.AdminAddUserToGroupOutput{}

        mockClient.On("AdminAddUserToGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AdminAddUserToGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminConfirmSignUp", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminConfirmSignUpInput{}
        output := &cognitoidentityprovider.AdminConfirmSignUpOutput{}

        mockClient.On("AdminConfirmSignUp", ctx, input).Return(output, nil)

        result, err := mockClient.AdminConfirmSignUp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminCreateUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminCreateUserInput{}
        output := &cognitoidentityprovider.AdminCreateUserOutput{}

        mockClient.On("AdminCreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminCreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminDeleteUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminDeleteUserInput{}
        output := &cognitoidentityprovider.AdminDeleteUserOutput{}

        mockClient.On("AdminDeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminDeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminDeleteUserAttributes", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminDeleteUserAttributesInput{}
        output := &cognitoidentityprovider.AdminDeleteUserAttributesOutput{}

        mockClient.On("AdminDeleteUserAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.AdminDeleteUserAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminDisableProviderForUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminDisableProviderForUserInput{}
        output := &cognitoidentityprovider.AdminDisableProviderForUserOutput{}

        mockClient.On("AdminDisableProviderForUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminDisableProviderForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminDisableUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminDisableUserInput{}
        output := &cognitoidentityprovider.AdminDisableUserOutput{}

        mockClient.On("AdminDisableUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminDisableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminEnableUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminEnableUserInput{}
        output := &cognitoidentityprovider.AdminEnableUserOutput{}

        mockClient.On("AdminEnableUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminEnableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminForgetDevice", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminForgetDeviceInput{}
        output := &cognitoidentityprovider.AdminForgetDeviceOutput{}

        mockClient.On("AdminForgetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.AdminForgetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminGetDevice", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminGetDeviceInput{}
        output := &cognitoidentityprovider.AdminGetDeviceOutput{}

        mockClient.On("AdminGetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.AdminGetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminGetUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminGetUserInput{}
        output := &cognitoidentityprovider.AdminGetUserOutput{}

        mockClient.On("AdminGetUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminGetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminInitiateAuth", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminInitiateAuthInput{}
        output := &cognitoidentityprovider.AdminInitiateAuthOutput{}

        mockClient.On("AdminInitiateAuth", ctx, input).Return(output, nil)

        result, err := mockClient.AdminInitiateAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminLinkProviderForUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminLinkProviderForUserInput{}
        output := &cognitoidentityprovider.AdminLinkProviderForUserOutput{}

        mockClient.On("AdminLinkProviderForUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminLinkProviderForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminListDevices", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminListDevicesInput{}
        output := &cognitoidentityprovider.AdminListDevicesOutput{}

        mockClient.On("AdminListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.AdminListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminListGroupsForUser", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminListGroupsForUserInput{}
        output := &cognitoidentityprovider.AdminListGroupsForUserOutput{}

        mockClient.On("AdminListGroupsForUser", ctx, input).Return(output, nil)

        result, err := mockClient.AdminListGroupsForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminListUserAuthEvents", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminListUserAuthEventsInput{}
        output := &cognitoidentityprovider.AdminListUserAuthEventsOutput{}

        mockClient.On("AdminListUserAuthEvents", ctx, input).Return(output, nil)

        result, err := mockClient.AdminListUserAuthEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminRemoveUserFromGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminRemoveUserFromGroupInput{}
        output := &cognitoidentityprovider.AdminRemoveUserFromGroupOutput{}

        mockClient.On("AdminRemoveUserFromGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AdminRemoveUserFromGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminResetUserPassword", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminResetUserPasswordInput{}
        output := &cognitoidentityprovider.AdminResetUserPasswordOutput{}

        mockClient.On("AdminResetUserPassword", ctx, input).Return(output, nil)

        result, err := mockClient.AdminResetUserPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminRespondToAuthChallenge", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{}
        output := &cognitoidentityprovider.AdminRespondToAuthChallengeOutput{}

        mockClient.On("AdminRespondToAuthChallenge", ctx, input).Return(output, nil)

        result, err := mockClient.AdminRespondToAuthChallenge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminSetUserMFAPreference", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminSetUserMFAPreferenceInput{}
        output := &cognitoidentityprovider.AdminSetUserMFAPreferenceOutput{}

        mockClient.On("AdminSetUserMFAPreference", ctx, input).Return(output, nil)

        result, err := mockClient.AdminSetUserMFAPreference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminSetUserPassword", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminSetUserPasswordInput{}
        output := &cognitoidentityprovider.AdminSetUserPasswordOutput{}

        mockClient.On("AdminSetUserPassword", ctx, input).Return(output, nil)

        result, err := mockClient.AdminSetUserPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminSetUserSettings", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminSetUserSettingsInput{}
        output := &cognitoidentityprovider.AdminSetUserSettingsOutput{}

        mockClient.On("AdminSetUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.AdminSetUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminUpdateAuthEventFeedback", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput{}
        output := &cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput{}

        mockClient.On("AdminUpdateAuthEventFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.AdminUpdateAuthEventFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminUpdateDeviceStatus", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminUpdateDeviceStatusInput{}
        output := &cognitoidentityprovider.AdminUpdateDeviceStatusOutput{}

        mockClient.On("AdminUpdateDeviceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.AdminUpdateDeviceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminUpdateUserAttributes", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{}
        output := &cognitoidentityprovider.AdminUpdateUserAttributesOutput{}

        mockClient.On("AdminUpdateUserAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.AdminUpdateUserAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAdminUserGlobalSignOut", func(t *testing.T) {
        input := &cognitoidentityprovider.AdminUserGlobalSignOutInput{}
        output := &cognitoidentityprovider.AdminUserGlobalSignOutOutput{}

        mockClient.On("AdminUserGlobalSignOut", ctx, input).Return(output, nil)

        result, err := mockClient.AdminUserGlobalSignOut(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSoftwareToken", func(t *testing.T) {
        input := &cognitoidentityprovider.AssociateSoftwareTokenInput{}
        output := &cognitoidentityprovider.AssociateSoftwareTokenOutput{}

        mockClient.On("AssociateSoftwareToken", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSoftwareToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangePassword", func(t *testing.T) {
        input := &cognitoidentityprovider.ChangePasswordInput{}
        output := &cognitoidentityprovider.ChangePasswordOutput{}

        mockClient.On("ChangePassword", ctx, input).Return(output, nil)

        result, err := mockClient.ChangePassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmDevice", func(t *testing.T) {
        input := &cognitoidentityprovider.ConfirmDeviceInput{}
        output := &cognitoidentityprovider.ConfirmDeviceOutput{}

        mockClient.On("ConfirmDevice", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmForgotPassword", func(t *testing.T) {
        input := &cognitoidentityprovider.ConfirmForgotPasswordInput{}
        output := &cognitoidentityprovider.ConfirmForgotPasswordOutput{}

        mockClient.On("ConfirmForgotPassword", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmForgotPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmSignUp", func(t *testing.T) {
        input := &cognitoidentityprovider.ConfirmSignUpInput{}
        output := &cognitoidentityprovider.ConfirmSignUpOutput{}

        mockClient.On("ConfirmSignUp", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmSignUp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateGroupInput{}
        output := &cognitoidentityprovider.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdentityProvider", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateIdentityProviderInput{}
        output := &cognitoidentityprovider.CreateIdentityProviderOutput{}

        mockClient.On("CreateIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceServer", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateResourceServerInput{}
        output := &cognitoidentityprovider.CreateResourceServerOutput{}

        mockClient.On("CreateResourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserImportJob", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateUserImportJobInput{}
        output := &cognitoidentityprovider.CreateUserImportJobOutput{}

        mockClient.On("CreateUserImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserPool", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateUserPoolInput{}
        output := &cognitoidentityprovider.CreateUserPoolOutput{}

        mockClient.On("CreateUserPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserPoolClient", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateUserPoolClientInput{}
        output := &cognitoidentityprovider.CreateUserPoolClientOutput{}

        mockClient.On("CreateUserPoolClient", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserPoolClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserPoolDomain", func(t *testing.T) {
        input := &cognitoidentityprovider.CreateUserPoolDomainInput{}
        output := &cognitoidentityprovider.CreateUserPoolDomainOutput{}

        mockClient.On("CreateUserPoolDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserPoolDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteGroupInput{}
        output := &cognitoidentityprovider.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentityProvider", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteIdentityProviderInput{}
        output := &cognitoidentityprovider.DeleteIdentityProviderOutput{}

        mockClient.On("DeleteIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceServer", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteResourceServerInput{}
        output := &cognitoidentityprovider.DeleteResourceServerOutput{}

        mockClient.On("DeleteResourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteUserInput{}
        output := &cognitoidentityprovider.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserAttributes", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteUserAttributesInput{}
        output := &cognitoidentityprovider.DeleteUserAttributesOutput{}

        mockClient.On("DeleteUserAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserPool", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteUserPoolInput{}
        output := &cognitoidentityprovider.DeleteUserPoolOutput{}

        mockClient.On("DeleteUserPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserPoolClient", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteUserPoolClientInput{}
        output := &cognitoidentityprovider.DeleteUserPoolClientOutput{}

        mockClient.On("DeleteUserPoolClient", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserPoolClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserPoolDomain", func(t *testing.T) {
        input := &cognitoidentityprovider.DeleteUserPoolDomainInput{}
        output := &cognitoidentityprovider.DeleteUserPoolDomainOutput{}

        mockClient.On("DeleteUserPoolDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserPoolDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityProvider", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeIdentityProviderInput{}
        output := &cognitoidentityprovider.DescribeIdentityProviderOutput{}

        mockClient.On("DescribeIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourceServer", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeResourceServerInput{}
        output := &cognitoidentityprovider.DescribeResourceServerOutput{}

        mockClient.On("DescribeResourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRiskConfiguration", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeRiskConfigurationInput{}
        output := &cognitoidentityprovider.DescribeRiskConfigurationOutput{}

        mockClient.On("DescribeRiskConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRiskConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserImportJob", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeUserImportJobInput{}
        output := &cognitoidentityprovider.DescribeUserImportJobOutput{}

        mockClient.On("DescribeUserImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserPool", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeUserPoolInput{}
        output := &cognitoidentityprovider.DescribeUserPoolOutput{}

        mockClient.On("DescribeUserPool", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserPoolClient", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeUserPoolClientInput{}
        output := &cognitoidentityprovider.DescribeUserPoolClientOutput{}

        mockClient.On("DescribeUserPoolClient", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserPoolClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserPoolDomain", func(t *testing.T) {
        input := &cognitoidentityprovider.DescribeUserPoolDomainInput{}
        output := &cognitoidentityprovider.DescribeUserPoolDomainOutput{}

        mockClient.On("DescribeUserPoolDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserPoolDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestForgetDevice", func(t *testing.T) {
        input := &cognitoidentityprovider.ForgetDeviceInput{}
        output := &cognitoidentityprovider.ForgetDeviceOutput{}

        mockClient.On("ForgetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.ForgetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestForgotPassword", func(t *testing.T) {
        input := &cognitoidentityprovider.ForgotPasswordInput{}
        output := &cognitoidentityprovider.ForgotPasswordOutput{}

        mockClient.On("ForgotPassword", ctx, input).Return(output, nil)

        result, err := mockClient.ForgotPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCSVHeader", func(t *testing.T) {
        input := &cognitoidentityprovider.GetCSVHeaderInput{}
        output := &cognitoidentityprovider.GetCSVHeaderOutput{}

        mockClient.On("GetCSVHeader", ctx, input).Return(output, nil)

        result, err := mockClient.GetCSVHeader(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevice", func(t *testing.T) {
        input := &cognitoidentityprovider.GetDeviceInput{}
        output := &cognitoidentityprovider.GetDeviceOutput{}

        mockClient.On("GetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.GetGroupInput{}
        output := &cognitoidentityprovider.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityProviderByIdentifier", func(t *testing.T) {
        input := &cognitoidentityprovider.GetIdentityProviderByIdentifierInput{}
        output := &cognitoidentityprovider.GetIdentityProviderByIdentifierOutput{}

        mockClient.On("GetIdentityProviderByIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityProviderByIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogDeliveryConfiguration", func(t *testing.T) {
        input := &cognitoidentityprovider.GetLogDeliveryConfigurationInput{}
        output := &cognitoidentityprovider.GetLogDeliveryConfigurationOutput{}

        mockClient.On("GetLogDeliveryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogDeliveryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSigningCertificate", func(t *testing.T) {
        input := &cognitoidentityprovider.GetSigningCertificateInput{}
        output := &cognitoidentityprovider.GetSigningCertificateOutput{}

        mockClient.On("GetSigningCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetSigningCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUICustomization", func(t *testing.T) {
        input := &cognitoidentityprovider.GetUICustomizationInput{}
        output := &cognitoidentityprovider.GetUICustomizationOutput{}

        mockClient.On("GetUICustomization", ctx, input).Return(output, nil)

        result, err := mockClient.GetUICustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUser", func(t *testing.T) {
        input := &cognitoidentityprovider.GetUserInput{}
        output := &cognitoidentityprovider.GetUserOutput{}

        mockClient.On("GetUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserAttributeVerificationCode", func(t *testing.T) {
        input := &cognitoidentityprovider.GetUserAttributeVerificationCodeInput{}
        output := &cognitoidentityprovider.GetUserAttributeVerificationCodeOutput{}

        mockClient.On("GetUserAttributeVerificationCode", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserAttributeVerificationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserPoolMfaConfig", func(t *testing.T) {
        input := &cognitoidentityprovider.GetUserPoolMfaConfigInput{}
        output := &cognitoidentityprovider.GetUserPoolMfaConfigOutput{}

        mockClient.On("GetUserPoolMfaConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserPoolMfaConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGlobalSignOut", func(t *testing.T) {
        input := &cognitoidentityprovider.GlobalSignOutInput{}
        output := &cognitoidentityprovider.GlobalSignOutOutput{}

        mockClient.On("GlobalSignOut", ctx, input).Return(output, nil)

        result, err := mockClient.GlobalSignOut(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateAuth", func(t *testing.T) {
        input := &cognitoidentityprovider.InitiateAuthInput{}
        output := &cognitoidentityprovider.InitiateAuthOutput{}

        mockClient.On("InitiateAuth", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &cognitoidentityprovider.ListDevicesInput{}
        output := &cognitoidentityprovider.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &cognitoidentityprovider.ListGroupsInput{}
        output := &cognitoidentityprovider.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityProviders", func(t *testing.T) {
        input := &cognitoidentityprovider.ListIdentityProvidersInput{}
        output := &cognitoidentityprovider.ListIdentityProvidersOutput{}

        mockClient.On("ListIdentityProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceServers", func(t *testing.T) {
        input := &cognitoidentityprovider.ListResourceServersInput{}
        output := &cognitoidentityprovider.ListResourceServersOutput{}

        mockClient.On("ListResourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cognitoidentityprovider.ListTagsForResourceInput{}
        output := &cognitoidentityprovider.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserImportJobs", func(t *testing.T) {
        input := &cognitoidentityprovider.ListUserImportJobsInput{}
        output := &cognitoidentityprovider.ListUserImportJobsOutput{}

        mockClient.On("ListUserImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserPoolClients", func(t *testing.T) {
        input := &cognitoidentityprovider.ListUserPoolClientsInput{}
        output := &cognitoidentityprovider.ListUserPoolClientsOutput{}

        mockClient.On("ListUserPoolClients", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserPoolClients(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserPools", func(t *testing.T) {
        input := &cognitoidentityprovider.ListUserPoolsInput{}
        output := &cognitoidentityprovider.ListUserPoolsOutput{}

        mockClient.On("ListUserPools", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &cognitoidentityprovider.ListUsersInput{}
        output := &cognitoidentityprovider.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsersInGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.ListUsersInGroupInput{}
        output := &cognitoidentityprovider.ListUsersInGroupOutput{}

        mockClient.On("ListUsersInGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsersInGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResendConfirmationCode", func(t *testing.T) {
        input := &cognitoidentityprovider.ResendConfirmationCodeInput{}
        output := &cognitoidentityprovider.ResendConfirmationCodeOutput{}

        mockClient.On("ResendConfirmationCode", ctx, input).Return(output, nil)

        result, err := mockClient.ResendConfirmationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRespondToAuthChallenge", func(t *testing.T) {
        input := &cognitoidentityprovider.RespondToAuthChallengeInput{}
        output := &cognitoidentityprovider.RespondToAuthChallengeOutput{}

        mockClient.On("RespondToAuthChallenge", ctx, input).Return(output, nil)

        result, err := mockClient.RespondToAuthChallenge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeToken", func(t *testing.T) {
        input := &cognitoidentityprovider.RevokeTokenInput{}
        output := &cognitoidentityprovider.RevokeTokenOutput{}

        mockClient.On("RevokeToken", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLogDeliveryConfiguration", func(t *testing.T) {
        input := &cognitoidentityprovider.SetLogDeliveryConfigurationInput{}
        output := &cognitoidentityprovider.SetLogDeliveryConfigurationOutput{}

        mockClient.On("SetLogDeliveryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.SetLogDeliveryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetRiskConfiguration", func(t *testing.T) {
        input := &cognitoidentityprovider.SetRiskConfigurationInput{}
        output := &cognitoidentityprovider.SetRiskConfigurationOutput{}

        mockClient.On("SetRiskConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.SetRiskConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetUICustomization", func(t *testing.T) {
        input := &cognitoidentityprovider.SetUICustomizationInput{}
        output := &cognitoidentityprovider.SetUICustomizationOutput{}

        mockClient.On("SetUICustomization", ctx, input).Return(output, nil)

        result, err := mockClient.SetUICustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetUserMFAPreference", func(t *testing.T) {
        input := &cognitoidentityprovider.SetUserMFAPreferenceInput{}
        output := &cognitoidentityprovider.SetUserMFAPreferenceOutput{}

        mockClient.On("SetUserMFAPreference", ctx, input).Return(output, nil)

        result, err := mockClient.SetUserMFAPreference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetUserPoolMfaConfig", func(t *testing.T) {
        input := &cognitoidentityprovider.SetUserPoolMfaConfigInput{}
        output := &cognitoidentityprovider.SetUserPoolMfaConfigOutput{}

        mockClient.On("SetUserPoolMfaConfig", ctx, input).Return(output, nil)

        result, err := mockClient.SetUserPoolMfaConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetUserSettings", func(t *testing.T) {
        input := &cognitoidentityprovider.SetUserSettingsInput{}
        output := &cognitoidentityprovider.SetUserSettingsOutput{}

        mockClient.On("SetUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.SetUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignUp", func(t *testing.T) {
        input := &cognitoidentityprovider.SignUpInput{}
        output := &cognitoidentityprovider.SignUpOutput{}

        mockClient.On("SignUp", ctx, input).Return(output, nil)

        result, err := mockClient.SignUp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartUserImportJob", func(t *testing.T) {
        input := &cognitoidentityprovider.StartUserImportJobInput{}
        output := &cognitoidentityprovider.StartUserImportJobOutput{}

        mockClient.On("StartUserImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartUserImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopUserImportJob", func(t *testing.T) {
        input := &cognitoidentityprovider.StopUserImportJobInput{}
        output := &cognitoidentityprovider.StopUserImportJobOutput{}

        mockClient.On("StopUserImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopUserImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cognitoidentityprovider.TagResourceInput{}
        output := &cognitoidentityprovider.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cognitoidentityprovider.UntagResourceInput{}
        output := &cognitoidentityprovider.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuthEventFeedback", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateAuthEventFeedbackInput{}
        output := &cognitoidentityprovider.UpdateAuthEventFeedbackOutput{}

        mockClient.On("UpdateAuthEventFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuthEventFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceStatus", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateDeviceStatusInput{}
        output := &cognitoidentityprovider.UpdateDeviceStatusOutput{}

        mockClient.On("UpdateDeviceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateGroupInput{}
        output := &cognitoidentityprovider.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityProvider", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateIdentityProviderInput{}
        output := &cognitoidentityprovider.UpdateIdentityProviderOutput{}

        mockClient.On("UpdateIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceServer", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateResourceServerInput{}
        output := &cognitoidentityprovider.UpdateResourceServerOutput{}

        mockClient.On("UpdateResourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserAttributes", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateUserAttributesInput{}
        output := &cognitoidentityprovider.UpdateUserAttributesOutput{}

        mockClient.On("UpdateUserAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserPool", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateUserPoolInput{}
        output := &cognitoidentityprovider.UpdateUserPoolOutput{}

        mockClient.On("UpdateUserPool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserPoolClient", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateUserPoolClientInput{}
        output := &cognitoidentityprovider.UpdateUserPoolClientOutput{}

        mockClient.On("UpdateUserPoolClient", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserPoolClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserPoolDomain", func(t *testing.T) {
        input := &cognitoidentityprovider.UpdateUserPoolDomainInput{}
        output := &cognitoidentityprovider.UpdateUserPoolDomainOutput{}

        mockClient.On("UpdateUserPoolDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserPoolDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifySoftwareToken", func(t *testing.T) {
        input := &cognitoidentityprovider.VerifySoftwareTokenInput{}
        output := &cognitoidentityprovider.VerifySoftwareTokenOutput{}

        mockClient.On("VerifySoftwareToken", ctx, input).Return(output, nil)

        result, err := mockClient.VerifySoftwareToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyUserAttribute", func(t *testing.T) {
        input := &cognitoidentityprovider.VerifyUserAttributeInput{}
        output := &cognitoidentityprovider.VerifyUserAttributeOutput{}

        mockClient.On("VerifyUserAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyUserAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
