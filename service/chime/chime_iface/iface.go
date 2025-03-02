
package chime_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/chime"
)

// IClient defines the interface for chime
type IClient interface {
 Options() Options 
 AssociatePhoneNumberWithUser(ctx context.Context, params *AssociatePhoneNumberWithUserInput, optFns ...func(*Options)) (*AssociatePhoneNumberWithUserOutput, error) 
 AssociateSigninDelegateGroupsWithAccount(ctx context.Context, params *AssociateSigninDelegateGroupsWithAccountInput, optFns ...func(*Options)) (*AssociateSigninDelegateGroupsWithAccountOutput, error) 
 BatchCreateRoomMembership(ctx context.Context, params *BatchCreateRoomMembershipInput, optFns ...func(*Options)) (*BatchCreateRoomMembershipOutput, error) 
 BatchDeletePhoneNumber(ctx context.Context, params *BatchDeletePhoneNumberInput, optFns ...func(*Options)) (*BatchDeletePhoneNumberOutput, error) 
 BatchSuspendUser(ctx context.Context, params *BatchSuspendUserInput, optFns ...func(*Options)) (*BatchSuspendUserOutput, error) 
 BatchUnsuspendUser(ctx context.Context, params *BatchUnsuspendUserInput, optFns ...func(*Options)) (*BatchUnsuspendUserOutput, error) 
 BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) 
 BatchUpdateUser(ctx context.Context, params *BatchUpdateUserInput, optFns ...func(*Options)) (*BatchUpdateUserOutput, error) 
 CreateAccount(ctx context.Context, params *CreateAccountInput, optFns ...func(*Options)) (*CreateAccountOutput, error) 
 CreateBot(ctx context.Context, params *CreateBotInput, optFns ...func(*Options)) (*CreateBotOutput, error) 
 CreateMeetingDialOut(ctx context.Context, params *CreateMeetingDialOutInput, optFns ...func(*Options)) (*CreateMeetingDialOutOutput, error) 
 CreatePhoneNumberOrder(ctx context.Context, params *CreatePhoneNumberOrderInput, optFns ...func(*Options)) (*CreatePhoneNumberOrderOutput, error) 
 CreateRoom(ctx context.Context, params *CreateRoomInput, optFns ...func(*Options)) (*CreateRoomOutput, error) 
 CreateRoomMembership(ctx context.Context, params *CreateRoomMembershipInput, optFns ...func(*Options)) (*CreateRoomMembershipOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteAccount(ctx context.Context, params *DeleteAccountInput, optFns ...func(*Options)) (*DeleteAccountOutput, error) 
 DeleteEventsConfiguration(ctx context.Context, params *DeleteEventsConfigurationInput, optFns ...func(*Options)) (*DeleteEventsConfigurationOutput, error) 
 DeletePhoneNumber(ctx context.Context, params *DeletePhoneNumberInput, optFns ...func(*Options)) (*DeletePhoneNumberOutput, error) 
 DeleteRoom(ctx context.Context, params *DeleteRoomInput, optFns ...func(*Options)) (*DeleteRoomOutput, error) 
 DeleteRoomMembership(ctx context.Context, params *DeleteRoomMembershipInput, optFns ...func(*Options)) (*DeleteRoomMembershipOutput, error) 
 DisassociatePhoneNumberFromUser(ctx context.Context, params *DisassociatePhoneNumberFromUserInput, optFns ...func(*Options)) (*DisassociatePhoneNumberFromUserOutput, error) 
 DisassociateSigninDelegateGroupsFromAccount(ctx context.Context, params *DisassociateSigninDelegateGroupsFromAccountInput, optFns ...func(*Options)) (*DisassociateSigninDelegateGroupsFromAccountOutput, error) 
 GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) 
 GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) 
 GetBot(ctx context.Context, params *GetBotInput, optFns ...func(*Options)) (*GetBotOutput, error) 
 GetEventsConfiguration(ctx context.Context, params *GetEventsConfigurationInput, optFns ...func(*Options)) (*GetEventsConfigurationOutput, error) 
 GetGlobalSettings(ctx context.Context, params *GetGlobalSettingsInput, optFns ...func(*Options)) (*GetGlobalSettingsOutput, error) 
 GetPhoneNumber(ctx context.Context, params *GetPhoneNumberInput, optFns ...func(*Options)) (*GetPhoneNumberOutput, error) 
 GetPhoneNumberOrder(ctx context.Context, params *GetPhoneNumberOrderInput, optFns ...func(*Options)) (*GetPhoneNumberOrderOutput, error) 
 GetPhoneNumberSettings(ctx context.Context, params *GetPhoneNumberSettingsInput, optFns ...func(*Options)) (*GetPhoneNumberSettingsOutput, error) 
 GetRetentionSettings(ctx context.Context, params *GetRetentionSettingsInput, optFns ...func(*Options)) (*GetRetentionSettingsOutput, error) 
 GetRoom(ctx context.Context, params *GetRoomInput, optFns ...func(*Options)) (*GetRoomOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetUserSettings(ctx context.Context, params *GetUserSettingsInput, optFns ...func(*Options)) (*GetUserSettingsOutput, error) 
 InviteUsers(ctx context.Context, params *InviteUsersInput, optFns ...func(*Options)) (*InviteUsersOutput, error) 
 ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) 
 ListBots(ctx context.Context, params *ListBotsInput, optFns ...func(*Options)) (*ListBotsOutput, error) 
 ListPhoneNumberOrders(ctx context.Context, params *ListPhoneNumberOrdersInput, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) 
 ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) 
 ListRoomMemberships(ctx context.Context, params *ListRoomMembershipsInput, optFns ...func(*Options)) (*ListRoomMembershipsOutput, error) 
 ListRooms(ctx context.Context, params *ListRoomsInput, optFns ...func(*Options)) (*ListRoomsOutput, error) 
 ListSupportedPhoneNumberCountries(ctx context.Context, params *ListSupportedPhoneNumberCountriesInput, optFns ...func(*Options)) (*ListSupportedPhoneNumberCountriesOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 LogoutUser(ctx context.Context, params *LogoutUserInput, optFns ...func(*Options)) (*LogoutUserOutput, error) 
 PutEventsConfiguration(ctx context.Context, params *PutEventsConfigurationInput, optFns ...func(*Options)) (*PutEventsConfigurationOutput, error) 
 PutRetentionSettings(ctx context.Context, params *PutRetentionSettingsInput, optFns ...func(*Options)) (*PutRetentionSettingsOutput, error) 
 RedactConversationMessage(ctx context.Context, params *RedactConversationMessageInput, optFns ...func(*Options)) (*RedactConversationMessageOutput, error) 
 RedactRoomMessage(ctx context.Context, params *RedactRoomMessageInput, optFns ...func(*Options)) (*RedactRoomMessageOutput, error) 
 RegenerateSecurityToken(ctx context.Context, params *RegenerateSecurityTokenInput, optFns ...func(*Options)) (*RegenerateSecurityTokenOutput, error) 
 ResetPersonalPIN(ctx context.Context, params *ResetPersonalPINInput, optFns ...func(*Options)) (*ResetPersonalPINOutput, error) 
 RestorePhoneNumber(ctx context.Context, params *RestorePhoneNumberInput, optFns ...func(*Options)) (*RestorePhoneNumberOutput, error) 
 SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) 
 UpdateAccount(ctx context.Context, params *UpdateAccountInput, optFns ...func(*Options)) (*UpdateAccountOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateBot(ctx context.Context, params *UpdateBotInput, optFns ...func(*Options)) (*UpdateBotOutput, error) 
 UpdateGlobalSettings(ctx context.Context, params *UpdateGlobalSettingsInput, optFns ...func(*Options)) (*UpdateGlobalSettingsOutput, error) 
 UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) 
 UpdatePhoneNumberSettings(ctx context.Context, params *UpdatePhoneNumberSettingsInput, optFns ...func(*Options)) (*UpdatePhoneNumberSettingsOutput, error) 
 UpdateRoom(ctx context.Context, params *UpdateRoomInput, optFns ...func(*Options)) (*UpdateRoomOutput, error) 
 UpdateRoomMembership(ctx context.Context, params *UpdateRoomMembershipInput, optFns ...func(*Options)) (*UpdateRoomMembershipOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
 UpdateUserSettings(ctx context.Context, params *UpdateUserSettingsInput, optFns ...func(*Options)) (*UpdateUserSettingsOutput, error) 
}
