
package cognitoidentityprovider

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// IClient defines the interface for cognitoidentityprovider
type IClient interface {
 Options() Options 
 AddCustomAttributes(ctx context.Context, params *AddCustomAttributesInput, optFns ...func(*Options)) (*AddCustomAttributesOutput, error) 
 AdminAddUserToGroup(ctx context.Context, params *AdminAddUserToGroupInput, optFns ...func(*Options)) (*AdminAddUserToGroupOutput, error) 
 AdminConfirmSignUp(ctx context.Context, params *AdminConfirmSignUpInput, optFns ...func(*Options)) (*AdminConfirmSignUpOutput, error) 
 AdminCreateUser(ctx context.Context, params *AdminCreateUserInput, optFns ...func(*Options)) (*AdminCreateUserOutput, error) 
 AdminDeleteUser(ctx context.Context, params *AdminDeleteUserInput, optFns ...func(*Options)) (*AdminDeleteUserOutput, error) 
 AdminDeleteUserAttributes(ctx context.Context, params *AdminDeleteUserAttributesInput, optFns ...func(*Options)) (*AdminDeleteUserAttributesOutput, error) 
 AdminDisableProviderForUser(ctx context.Context, params *AdminDisableProviderForUserInput, optFns ...func(*Options)) (*AdminDisableProviderForUserOutput, error) 
 AdminDisableUser(ctx context.Context, params *AdminDisableUserInput, optFns ...func(*Options)) (*AdminDisableUserOutput, error) 
 AdminEnableUser(ctx context.Context, params *AdminEnableUserInput, optFns ...func(*Options)) (*AdminEnableUserOutput, error) 
 AdminForgetDevice(ctx context.Context, params *AdminForgetDeviceInput, optFns ...func(*Options)) (*AdminForgetDeviceOutput, error) 
 AdminGetDevice(ctx context.Context, params *AdminGetDeviceInput, optFns ...func(*Options)) (*AdminGetDeviceOutput, error) 
 AdminGetUser(ctx context.Context, params *AdminGetUserInput, optFns ...func(*Options)) (*AdminGetUserOutput, error) 
 AdminInitiateAuth(ctx context.Context, params *AdminInitiateAuthInput, optFns ...func(*Options)) (*AdminInitiateAuthOutput, error) 
 AdminLinkProviderForUser(ctx context.Context, params *AdminLinkProviderForUserInput, optFns ...func(*Options)) (*AdminLinkProviderForUserOutput, error) 
 AdminListDevices(ctx context.Context, params *AdminListDevicesInput, optFns ...func(*Options)) (*AdminListDevicesOutput, error) 
 AdminListGroupsForUser(ctx context.Context, params *AdminListGroupsForUserInput, optFns ...func(*Options)) (*AdminListGroupsForUserOutput, error) 
 AdminListUserAuthEvents(ctx context.Context, params *AdminListUserAuthEventsInput, optFns ...func(*Options)) (*AdminListUserAuthEventsOutput, error) 
 AdminRemoveUserFromGroup(ctx context.Context, params *AdminRemoveUserFromGroupInput, optFns ...func(*Options)) (*AdminRemoveUserFromGroupOutput, error) 
 AdminResetUserPassword(ctx context.Context, params *AdminResetUserPasswordInput, optFns ...func(*Options)) (*AdminResetUserPasswordOutput, error) 
 AdminRespondToAuthChallenge(ctx context.Context, params *AdminRespondToAuthChallengeInput, optFns ...func(*Options)) (*AdminRespondToAuthChallengeOutput, error) 
 AdminSetUserMFAPreference(ctx context.Context, params *AdminSetUserMFAPreferenceInput, optFns ...func(*Options)) (*AdminSetUserMFAPreferenceOutput, error) 
 AdminSetUserPassword(ctx context.Context, params *AdminSetUserPasswordInput, optFns ...func(*Options)) (*AdminSetUserPasswordOutput, error) 
 AdminSetUserSettings(ctx context.Context, params *AdminSetUserSettingsInput, optFns ...func(*Options)) (*AdminSetUserSettingsOutput, error) 
 AdminUpdateAuthEventFeedback(ctx context.Context, params *AdminUpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*AdminUpdateAuthEventFeedbackOutput, error) 
 AdminUpdateDeviceStatus(ctx context.Context, params *AdminUpdateDeviceStatusInput, optFns ...func(*Options)) (*AdminUpdateDeviceStatusOutput, error) 
 AdminUpdateUserAttributes(ctx context.Context, params *AdminUpdateUserAttributesInput, optFns ...func(*Options)) (*AdminUpdateUserAttributesOutput, error) 
 AdminUserGlobalSignOut(ctx context.Context, params *AdminUserGlobalSignOutInput, optFns ...func(*Options)) (*AdminUserGlobalSignOutOutput, error) 
 AssociateSoftwareToken(ctx context.Context, params *AssociateSoftwareTokenInput, optFns ...func(*Options)) (*AssociateSoftwareTokenOutput, error) 
 ChangePassword(ctx context.Context, params *ChangePasswordInput, optFns ...func(*Options)) (*ChangePasswordOutput, error) 
 ConfirmDevice(ctx context.Context, params *ConfirmDeviceInput, optFns ...func(*Options)) (*ConfirmDeviceOutput, error) 
 ConfirmForgotPassword(ctx context.Context, params *ConfirmForgotPasswordInput, optFns ...func(*Options)) (*ConfirmForgotPasswordOutput, error) 
 ConfirmSignUp(ctx context.Context, params *ConfirmSignUpInput, optFns ...func(*Options)) (*ConfirmSignUpOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateIdentityProvider(ctx context.Context, params *CreateIdentityProviderInput, optFns ...func(*Options)) (*CreateIdentityProviderOutput, error) 
 CreateResourceServer(ctx context.Context, params *CreateResourceServerInput, optFns ...func(*Options)) (*CreateResourceServerOutput, error) 
 CreateUserImportJob(ctx context.Context, params *CreateUserImportJobInput, optFns ...func(*Options)) (*CreateUserImportJobOutput, error) 
 CreateUserPool(ctx context.Context, params *CreateUserPoolInput, optFns ...func(*Options)) (*CreateUserPoolOutput, error) 
 CreateUserPoolClient(ctx context.Context, params *CreateUserPoolClientInput, optFns ...func(*Options)) (*CreateUserPoolClientOutput, error) 
 CreateUserPoolDomain(ctx context.Context, params *CreateUserPoolDomainInput, optFns ...func(*Options)) (*CreateUserPoolDomainOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteIdentityProvider(ctx context.Context, params *DeleteIdentityProviderInput, optFns ...func(*Options)) (*DeleteIdentityProviderOutput, error) 
 DeleteResourceServer(ctx context.Context, params *DeleteResourceServerInput, optFns ...func(*Options)) (*DeleteResourceServerOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DeleteUserAttributes(ctx context.Context, params *DeleteUserAttributesInput, optFns ...func(*Options)) (*DeleteUserAttributesOutput, error) 
 DeleteUserPool(ctx context.Context, params *DeleteUserPoolInput, optFns ...func(*Options)) (*DeleteUserPoolOutput, error) 
 DeleteUserPoolClient(ctx context.Context, params *DeleteUserPoolClientInput, optFns ...func(*Options)) (*DeleteUserPoolClientOutput, error) 
 DeleteUserPoolDomain(ctx context.Context, params *DeleteUserPoolDomainInput, optFns ...func(*Options)) (*DeleteUserPoolDomainOutput, error) 
 DescribeIdentityProvider(ctx context.Context, params *DescribeIdentityProviderInput, optFns ...func(*Options)) (*DescribeIdentityProviderOutput, error) 
 DescribeResourceServer(ctx context.Context, params *DescribeResourceServerInput, optFns ...func(*Options)) (*DescribeResourceServerOutput, error) 
 DescribeRiskConfiguration(ctx context.Context, params *DescribeRiskConfigurationInput, optFns ...func(*Options)) (*DescribeRiskConfigurationOutput, error) 
 DescribeUserImportJob(ctx context.Context, params *DescribeUserImportJobInput, optFns ...func(*Options)) (*DescribeUserImportJobOutput, error) 
 DescribeUserPool(ctx context.Context, params *DescribeUserPoolInput, optFns ...func(*Options)) (*DescribeUserPoolOutput, error) 
 DescribeUserPoolClient(ctx context.Context, params *DescribeUserPoolClientInput, optFns ...func(*Options)) (*DescribeUserPoolClientOutput, error) 
 DescribeUserPoolDomain(ctx context.Context, params *DescribeUserPoolDomainInput, optFns ...func(*Options)) (*DescribeUserPoolDomainOutput, error) 
 ForgetDevice(ctx context.Context, params *ForgetDeviceInput, optFns ...func(*Options)) (*ForgetDeviceOutput, error) 
 ForgotPassword(ctx context.Context, params *ForgotPasswordInput, optFns ...func(*Options)) (*ForgotPasswordOutput, error) 
 GetCSVHeader(ctx context.Context, params *GetCSVHeaderInput, optFns ...func(*Options)) (*GetCSVHeaderOutput, error) 
 GetDevice(ctx context.Context, params *GetDeviceInput, optFns ...func(*Options)) (*GetDeviceOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetIdentityProviderByIdentifier(ctx context.Context, params *GetIdentityProviderByIdentifierInput, optFns ...func(*Options)) (*GetIdentityProviderByIdentifierOutput, error) 
 GetLogDeliveryConfiguration(ctx context.Context, params *GetLogDeliveryConfigurationInput, optFns ...func(*Options)) (*GetLogDeliveryConfigurationOutput, error) 
 GetSigningCertificate(ctx context.Context, params *GetSigningCertificateInput, optFns ...func(*Options)) (*GetSigningCertificateOutput, error) 
 GetUICustomization(ctx context.Context, params *GetUICustomizationInput, optFns ...func(*Options)) (*GetUICustomizationOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetUserAttributeVerificationCode(ctx context.Context, params *GetUserAttributeVerificationCodeInput, optFns ...func(*Options)) (*GetUserAttributeVerificationCodeOutput, error) 
 GetUserPoolMfaConfig(ctx context.Context, params *GetUserPoolMfaConfigInput, optFns ...func(*Options)) (*GetUserPoolMfaConfigOutput, error) 
 GlobalSignOut(ctx context.Context, params *GlobalSignOutInput, optFns ...func(*Options)) (*GlobalSignOutOutput, error) 
 InitiateAuth(ctx context.Context, params *InitiateAuthInput, optFns ...func(*Options)) (*InitiateAuthOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListIdentityProviders(ctx context.Context, params *ListIdentityProvidersInput, optFns ...func(*Options)) (*ListIdentityProvidersOutput, error) 
 ListResourceServers(ctx context.Context, params *ListResourceServersInput, optFns ...func(*Options)) (*ListResourceServersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListUserImportJobs(ctx context.Context, params *ListUserImportJobsInput, optFns ...func(*Options)) (*ListUserImportJobsOutput, error) 
 ListUserPoolClients(ctx context.Context, params *ListUserPoolClientsInput, optFns ...func(*Options)) (*ListUserPoolClientsOutput, error) 
 ListUserPools(ctx context.Context, params *ListUserPoolsInput, optFns ...func(*Options)) (*ListUserPoolsOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListUsersInGroup(ctx context.Context, params *ListUsersInGroupInput, optFns ...func(*Options)) (*ListUsersInGroupOutput, error) 
 ResendConfirmationCode(ctx context.Context, params *ResendConfirmationCodeInput, optFns ...func(*Options)) (*ResendConfirmationCodeOutput, error) 
 RespondToAuthChallenge(ctx context.Context, params *RespondToAuthChallengeInput, optFns ...func(*Options)) (*RespondToAuthChallengeOutput, error) 
 RevokeToken(ctx context.Context, params *RevokeTokenInput, optFns ...func(*Options)) (*RevokeTokenOutput, error) 
 SetLogDeliveryConfiguration(ctx context.Context, params *SetLogDeliveryConfigurationInput, optFns ...func(*Options)) (*SetLogDeliveryConfigurationOutput, error) 
 SetRiskConfiguration(ctx context.Context, params *SetRiskConfigurationInput, optFns ...func(*Options)) (*SetRiskConfigurationOutput, error) 
 SetUICustomization(ctx context.Context, params *SetUICustomizationInput, optFns ...func(*Options)) (*SetUICustomizationOutput, error) 
 SetUserMFAPreference(ctx context.Context, params *SetUserMFAPreferenceInput, optFns ...func(*Options)) (*SetUserMFAPreferenceOutput, error) 
 SetUserPoolMfaConfig(ctx context.Context, params *SetUserPoolMfaConfigInput, optFns ...func(*Options)) (*SetUserPoolMfaConfigOutput, error) 
 SetUserSettings(ctx context.Context, params *SetUserSettingsInput, optFns ...func(*Options)) (*SetUserSettingsOutput, error) 
 SignUp(ctx context.Context, params *SignUpInput, optFns ...func(*Options)) (*SignUpOutput, error) 
 StartUserImportJob(ctx context.Context, params *StartUserImportJobInput, optFns ...func(*Options)) (*StartUserImportJobOutput, error) 
 StopUserImportJob(ctx context.Context, params *StopUserImportJobInput, optFns ...func(*Options)) (*StopUserImportJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAuthEventFeedback(ctx context.Context, params *UpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*UpdateAuthEventFeedbackOutput, error) 
 UpdateDeviceStatus(ctx context.Context, params *UpdateDeviceStatusInput, optFns ...func(*Options)) (*UpdateDeviceStatusOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateIdentityProvider(ctx context.Context, params *UpdateIdentityProviderInput, optFns ...func(*Options)) (*UpdateIdentityProviderOutput, error) 
 UpdateResourceServer(ctx context.Context, params *UpdateResourceServerInput, optFns ...func(*Options)) (*UpdateResourceServerOutput, error) 
 UpdateUserAttributes(ctx context.Context, params *UpdateUserAttributesInput, optFns ...func(*Options)) (*UpdateUserAttributesOutput, error) 
 UpdateUserPool(ctx context.Context, params *UpdateUserPoolInput, optFns ...func(*Options)) (*UpdateUserPoolOutput, error) 
 UpdateUserPoolClient(ctx context.Context, params *UpdateUserPoolClientInput, optFns ...func(*Options)) (*UpdateUserPoolClientOutput, error) 
 UpdateUserPoolDomain(ctx context.Context, params *UpdateUserPoolDomainInput, optFns ...func(*Options)) (*UpdateUserPoolDomainOutput, error) 
 VerifySoftwareToken(ctx context.Context, params *VerifySoftwareTokenInput, optFns ...func(*Options)) (*VerifySoftwareTokenOutput, error) 
 VerifyUserAttribute(ctx context.Context, params *VerifyUserAttributeInput, optFns ...func(*Options)) (*VerifyUserAttributeOutput, error) 
}
