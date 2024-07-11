
package chime

import (
    "github.com/aws/aws-sdk-go-v2/service/chime"
)

// IChime defines the interface for chime
type IChime interface {
 Options() Options 
 AssociatePhoneNumberWithUser(ctx context.Context, params *AssociatePhoneNumberWithUserInput, optFns ...func(*Options)) (*AssociatePhoneNumberWithUserOutput, error) 
 AssociatePhoneNumbersWithVoiceConnector(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorOutput, error) 
 AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorGroupInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) 
 AssociateSigninDelegateGroupsWithAccount(ctx context.Context, params *AssociateSigninDelegateGroupsWithAccountInput, optFns ...func(*Options)) (*AssociateSigninDelegateGroupsWithAccountOutput, error) 
 BatchCreateAttendee(ctx context.Context, params *BatchCreateAttendeeInput, optFns ...func(*Options)) (*BatchCreateAttendeeOutput, error) 
 BatchCreateChannelMembership(ctx context.Context, params *BatchCreateChannelMembershipInput, optFns ...func(*Options)) (*BatchCreateChannelMembershipOutput, error) 
 BatchCreateRoomMembership(ctx context.Context, params *BatchCreateRoomMembershipInput, optFns ...func(*Options)) (*BatchCreateRoomMembershipOutput, error) 
 BatchDeletePhoneNumber(ctx context.Context, params *BatchDeletePhoneNumberInput, optFns ...func(*Options)) (*BatchDeletePhoneNumberOutput, error) 
 BatchSuspendUser(ctx context.Context, params *BatchSuspendUserInput, optFns ...func(*Options)) (*BatchSuspendUserOutput, error) 
 BatchUnsuspendUser(ctx context.Context, params *BatchUnsuspendUserInput, optFns ...func(*Options)) (*BatchUnsuspendUserOutput, error) 
 BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) 
 BatchUpdateUser(ctx context.Context, params *BatchUpdateUserInput, optFns ...func(*Options)) (*BatchUpdateUserOutput, error) 
 CreateAccount(ctx context.Context, params *CreateAccountInput, optFns ...func(*Options)) (*CreateAccountOutput, error) 
 CreateAppInstance(ctx context.Context, params *CreateAppInstanceInput, optFns ...func(*Options)) (*CreateAppInstanceOutput, error) 
 CreateAppInstanceAdmin(ctx context.Context, params *CreateAppInstanceAdminInput, optFns ...func(*Options)) (*CreateAppInstanceAdminOutput, error) 
 CreateAppInstanceUser(ctx context.Context, params *CreateAppInstanceUserInput, optFns ...func(*Options)) (*CreateAppInstanceUserOutput, error) 
 CreateAttendee(ctx context.Context, params *CreateAttendeeInput, optFns ...func(*Options)) (*CreateAttendeeOutput, error) 
 CreateBot(ctx context.Context, params *CreateBotInput, optFns ...func(*Options)) (*CreateBotOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateChannelBan(ctx context.Context, params *CreateChannelBanInput, optFns ...func(*Options)) (*CreateChannelBanOutput, error) 
 CreateChannelMembership(ctx context.Context, params *CreateChannelMembershipInput, optFns ...func(*Options)) (*CreateChannelMembershipOutput, error) 
 CreateChannelModerator(ctx context.Context, params *CreateChannelModeratorInput, optFns ...func(*Options)) (*CreateChannelModeratorOutput, error) 
 CreateMediaCapturePipeline(ctx context.Context, params *CreateMediaCapturePipelineInput, optFns ...func(*Options)) (*CreateMediaCapturePipelineOutput, error) 
 CreateMeeting(ctx context.Context, params *CreateMeetingInput, optFns ...func(*Options)) (*CreateMeetingOutput, error) 
 CreateMeetingDialOut(ctx context.Context, params *CreateMeetingDialOutInput, optFns ...func(*Options)) (*CreateMeetingDialOutOutput, error) 
 CreateMeetingWithAttendees(ctx context.Context, params *CreateMeetingWithAttendeesInput, optFns ...func(*Options)) (*CreateMeetingWithAttendeesOutput, error) 
 CreatePhoneNumberOrder(ctx context.Context, params *CreatePhoneNumberOrderInput, optFns ...func(*Options)) (*CreatePhoneNumberOrderOutput, error) 
 CreateProxySession(ctx context.Context, params *CreateProxySessionInput, optFns ...func(*Options)) (*CreateProxySessionOutput, error) 
 CreateRoom(ctx context.Context, params *CreateRoomInput, optFns ...func(*Options)) (*CreateRoomOutput, error) 
 CreateRoomMembership(ctx context.Context, params *CreateRoomMembershipInput, optFns ...func(*Options)) (*CreateRoomMembershipOutput, error) 
 CreateSipMediaApplication(ctx context.Context, params *CreateSipMediaApplicationInput, optFns ...func(*Options)) (*CreateSipMediaApplicationOutput, error) 
 CreateSipMediaApplicationCall(ctx context.Context, params *CreateSipMediaApplicationCallInput, optFns ...func(*Options)) (*CreateSipMediaApplicationCallOutput, error) 
 CreateSipRule(ctx context.Context, params *CreateSipRuleInput, optFns ...func(*Options)) (*CreateSipRuleOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 CreateVoiceConnector(ctx context.Context, params *CreateVoiceConnectorInput, optFns ...func(*Options)) (*CreateVoiceConnectorOutput, error) 
 CreateVoiceConnectorGroup(ctx context.Context, params *CreateVoiceConnectorGroupInput, optFns ...func(*Options)) (*CreateVoiceConnectorGroupOutput, error) 
 DeleteAccount(ctx context.Context, params *DeleteAccountInput, optFns ...func(*Options)) (*DeleteAccountOutput, error) 
 DeleteAppInstance(ctx context.Context, params *DeleteAppInstanceInput, optFns ...func(*Options)) (*DeleteAppInstanceOutput, error) 
 DeleteAppInstanceAdmin(ctx context.Context, params *DeleteAppInstanceAdminInput, optFns ...func(*Options)) (*DeleteAppInstanceAdminOutput, error) 
 DeleteAppInstanceStreamingConfigurations(ctx context.Context, params *DeleteAppInstanceStreamingConfigurationsInput, optFns ...func(*Options)) (*DeleteAppInstanceStreamingConfigurationsOutput, error) 
 DeleteAppInstanceUser(ctx context.Context, params *DeleteAppInstanceUserInput, optFns ...func(*Options)) (*DeleteAppInstanceUserOutput, error) 
 DeleteAttendee(ctx context.Context, params *DeleteAttendeeInput, optFns ...func(*Options)) (*DeleteAttendeeOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteChannelBan(ctx context.Context, params *DeleteChannelBanInput, optFns ...func(*Options)) (*DeleteChannelBanOutput, error) 
 DeleteChannelMembership(ctx context.Context, params *DeleteChannelMembershipInput, optFns ...func(*Options)) (*DeleteChannelMembershipOutput, error) 
 DeleteChannelMessage(ctx context.Context, params *DeleteChannelMessageInput, optFns ...func(*Options)) (*DeleteChannelMessageOutput, error) 
 DeleteChannelModerator(ctx context.Context, params *DeleteChannelModeratorInput, optFns ...func(*Options)) (*DeleteChannelModeratorOutput, error) 
 DeleteEventsConfiguration(ctx context.Context, params *DeleteEventsConfigurationInput, optFns ...func(*Options)) (*DeleteEventsConfigurationOutput, error) 
 DeleteMediaCapturePipeline(ctx context.Context, params *DeleteMediaCapturePipelineInput, optFns ...func(*Options)) (*DeleteMediaCapturePipelineOutput, error) 
 DeleteMeeting(ctx context.Context, params *DeleteMeetingInput, optFns ...func(*Options)) (*DeleteMeetingOutput, error) 
 DeletePhoneNumber(ctx context.Context, params *DeletePhoneNumberInput, optFns ...func(*Options)) (*DeletePhoneNumberOutput, error) 
 DeleteProxySession(ctx context.Context, params *DeleteProxySessionInput, optFns ...func(*Options)) (*DeleteProxySessionOutput, error) 
 DeleteRoom(ctx context.Context, params *DeleteRoomInput, optFns ...func(*Options)) (*DeleteRoomOutput, error) 
 DeleteRoomMembership(ctx context.Context, params *DeleteRoomMembershipInput, optFns ...func(*Options)) (*DeleteRoomMembershipOutput, error) 
 DeleteSipMediaApplication(ctx context.Context, params *DeleteSipMediaApplicationInput, optFns ...func(*Options)) (*DeleteSipMediaApplicationOutput, error) 
 DeleteSipRule(ctx context.Context, params *DeleteSipRuleInput, optFns ...func(*Options)) (*DeleteSipRuleOutput, error) 
 DeleteVoiceConnector(ctx context.Context, params *DeleteVoiceConnectorInput, optFns ...func(*Options)) (*DeleteVoiceConnectorOutput, error) 
 DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *DeleteVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 DeleteVoiceConnectorGroup(ctx context.Context, params *DeleteVoiceConnectorGroupInput, optFns ...func(*Options)) (*DeleteVoiceConnectorGroupOutput, error) 
 DeleteVoiceConnectorOrigination(ctx context.Context, params *DeleteVoiceConnectorOriginationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorOriginationOutput, error) 
 DeleteVoiceConnectorProxy(ctx context.Context, params *DeleteVoiceConnectorProxyInput, optFns ...func(*Options)) (*DeleteVoiceConnectorProxyOutput, error) 
 DeleteVoiceConnectorStreamingConfiguration(ctx context.Context, params *DeleteVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorStreamingConfigurationOutput, error) 
 DeleteVoiceConnectorTermination(ctx context.Context, params *DeleteVoiceConnectorTerminationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationOutput, error) 
 DeleteVoiceConnectorTerminationCredentials(ctx context.Context, params *DeleteVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationCredentialsOutput, error) 
 DescribeAppInstance(ctx context.Context, params *DescribeAppInstanceInput, optFns ...func(*Options)) (*DescribeAppInstanceOutput, error) 
 DescribeAppInstanceAdmin(ctx context.Context, params *DescribeAppInstanceAdminInput, optFns ...func(*Options)) (*DescribeAppInstanceAdminOutput, error) 
 DescribeAppInstanceUser(ctx context.Context, params *DescribeAppInstanceUserInput, optFns ...func(*Options)) (*DescribeAppInstanceUserOutput, error) 
 DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) 
 DescribeChannelBan(ctx context.Context, params *DescribeChannelBanInput, optFns ...func(*Options)) (*DescribeChannelBanOutput, error) 
 DescribeChannelMembership(ctx context.Context, params *DescribeChannelMembershipInput, optFns ...func(*Options)) (*DescribeChannelMembershipOutput, error) 
 DescribeChannelMembershipForAppInstanceUser(ctx context.Context, params *DescribeChannelMembershipForAppInstanceUserInput, optFns ...func(*Options)) (*DescribeChannelMembershipForAppInstanceUserOutput, error) 
 DescribeChannelModeratedByAppInstanceUser(ctx context.Context, params *DescribeChannelModeratedByAppInstanceUserInput, optFns ...func(*Options)) (*DescribeChannelModeratedByAppInstanceUserOutput, error) 
 DescribeChannelModerator(ctx context.Context, params *DescribeChannelModeratorInput, optFns ...func(*Options)) (*DescribeChannelModeratorOutput, error) 
 DisassociatePhoneNumberFromUser(ctx context.Context, params *DisassociatePhoneNumberFromUserInput, optFns ...func(*Options)) (*DisassociatePhoneNumberFromUserOutput, error) 
 DisassociatePhoneNumbersFromVoiceConnector(ctx context.Context, params *DisassociatePhoneNumbersFromVoiceConnectorInput, optFns ...func(*Options)) (*DisassociatePhoneNumbersFromVoiceConnectorOutput, error) 
 DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx context.Context, params *DisassociatePhoneNumbersFromVoiceConnectorGroupInput, optFns ...func(*Options)) (*DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) 
 DisassociateSigninDelegateGroupsFromAccount(ctx context.Context, params *DisassociateSigninDelegateGroupsFromAccountInput, optFns ...func(*Options)) (*DisassociateSigninDelegateGroupsFromAccountOutput, error) 
 GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) 
 GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) 
 GetAppInstanceRetentionSettings(ctx context.Context, params *GetAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*GetAppInstanceRetentionSettingsOutput, error) 
 GetAppInstanceStreamingConfigurations(ctx context.Context, params *GetAppInstanceStreamingConfigurationsInput, optFns ...func(*Options)) (*GetAppInstanceStreamingConfigurationsOutput, error) 
 GetAttendee(ctx context.Context, params *GetAttendeeInput, optFns ...func(*Options)) (*GetAttendeeOutput, error) 
 GetBot(ctx context.Context, params *GetBotInput, optFns ...func(*Options)) (*GetBotOutput, error) 
 GetChannelMessage(ctx context.Context, params *GetChannelMessageInput, optFns ...func(*Options)) (*GetChannelMessageOutput, error) 
 GetEventsConfiguration(ctx context.Context, params *GetEventsConfigurationInput, optFns ...func(*Options)) (*GetEventsConfigurationOutput, error) 
 GetGlobalSettings(ctx context.Context, params *GetGlobalSettingsInput, optFns ...func(*Options)) (*GetGlobalSettingsOutput, error) 
 GetMediaCapturePipeline(ctx context.Context, params *GetMediaCapturePipelineInput, optFns ...func(*Options)) (*GetMediaCapturePipelineOutput, error) 
 GetMeeting(ctx context.Context, params *GetMeetingInput, optFns ...func(*Options)) (*GetMeetingOutput, error) 
 GetMessagingSessionEndpoint(ctx context.Context, params *GetMessagingSessionEndpointInput, optFns ...func(*Options)) (*GetMessagingSessionEndpointOutput, error) 
 GetPhoneNumber(ctx context.Context, params *GetPhoneNumberInput, optFns ...func(*Options)) (*GetPhoneNumberOutput, error) 
 GetPhoneNumberOrder(ctx context.Context, params *GetPhoneNumberOrderInput, optFns ...func(*Options)) (*GetPhoneNumberOrderOutput, error) 
 GetPhoneNumberSettings(ctx context.Context, params *GetPhoneNumberSettingsInput, optFns ...func(*Options)) (*GetPhoneNumberSettingsOutput, error) 
 GetProxySession(ctx context.Context, params *GetProxySessionInput, optFns ...func(*Options)) (*GetProxySessionOutput, error) 
 GetRetentionSettings(ctx context.Context, params *GetRetentionSettingsInput, optFns ...func(*Options)) (*GetRetentionSettingsOutput, error) 
 GetRoom(ctx context.Context, params *GetRoomInput, optFns ...func(*Options)) (*GetRoomOutput, error) 
 GetSipMediaApplication(ctx context.Context, params *GetSipMediaApplicationInput, optFns ...func(*Options)) (*GetSipMediaApplicationOutput, error) 
 GetSipMediaApplicationLoggingConfiguration(ctx context.Context, params *GetSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*GetSipMediaApplicationLoggingConfigurationOutput, error) 
 GetSipRule(ctx context.Context, params *GetSipRuleInput, optFns ...func(*Options)) (*GetSipRuleOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetUserSettings(ctx context.Context, params *GetUserSettingsInput, optFns ...func(*Options)) (*GetUserSettingsOutput, error) 
 GetVoiceConnector(ctx context.Context, params *GetVoiceConnectorInput, optFns ...func(*Options)) (*GetVoiceConnectorOutput, error) 
 GetVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *GetVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 GetVoiceConnectorGroup(ctx context.Context, params *GetVoiceConnectorGroupInput, optFns ...func(*Options)) (*GetVoiceConnectorGroupOutput, error) 
 GetVoiceConnectorLoggingConfiguration(ctx context.Context, params *GetVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorLoggingConfigurationOutput, error) 
 GetVoiceConnectorOrigination(ctx context.Context, params *GetVoiceConnectorOriginationInput, optFns ...func(*Options)) (*GetVoiceConnectorOriginationOutput, error) 
 GetVoiceConnectorProxy(ctx context.Context, params *GetVoiceConnectorProxyInput, optFns ...func(*Options)) (*GetVoiceConnectorProxyOutput, error) 
 GetVoiceConnectorStreamingConfiguration(ctx context.Context, params *GetVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorStreamingConfigurationOutput, error) 
 GetVoiceConnectorTermination(ctx context.Context, params *GetVoiceConnectorTerminationInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationOutput, error) 
 GetVoiceConnectorTerminationHealth(ctx context.Context, params *GetVoiceConnectorTerminationHealthInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationHealthOutput, error) 
 InviteUsers(ctx context.Context, params *InviteUsersInput, optFns ...func(*Options)) (*InviteUsersOutput, error) 
 ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) 
 ListAppInstanceAdmins(ctx context.Context, params *ListAppInstanceAdminsInput, optFns ...func(*Options)) (*ListAppInstanceAdminsOutput, error) 
 ListAppInstanceUsers(ctx context.Context, params *ListAppInstanceUsersInput, optFns ...func(*Options)) (*ListAppInstanceUsersOutput, error) 
 ListAppInstances(ctx context.Context, params *ListAppInstancesInput, optFns ...func(*Options)) (*ListAppInstancesOutput, error) 
 ListAttendeeTags(ctx context.Context, params *ListAttendeeTagsInput, optFns ...func(*Options)) (*ListAttendeeTagsOutput, error) 
 ListAttendees(ctx context.Context, params *ListAttendeesInput, optFns ...func(*Options)) (*ListAttendeesOutput, error) 
 ListBots(ctx context.Context, params *ListBotsInput, optFns ...func(*Options)) (*ListBotsOutput, error) 
 ListChannelBans(ctx context.Context, params *ListChannelBansInput, optFns ...func(*Options)) (*ListChannelBansOutput, error) 
 ListChannelMemberships(ctx context.Context, params *ListChannelMembershipsInput, optFns ...func(*Options)) (*ListChannelMembershipsOutput, error) 
 ListChannelMembershipsForAppInstanceUser(ctx context.Context, params *ListChannelMembershipsForAppInstanceUserInput, optFns ...func(*Options)) (*ListChannelMembershipsForAppInstanceUserOutput, error) 
 ListChannelMessages(ctx context.Context, params *ListChannelMessagesInput, optFns ...func(*Options)) (*ListChannelMessagesOutput, error) 
 ListChannelModerators(ctx context.Context, params *ListChannelModeratorsInput, optFns ...func(*Options)) (*ListChannelModeratorsOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListChannelsModeratedByAppInstanceUser(ctx context.Context, params *ListChannelsModeratedByAppInstanceUserInput, optFns ...func(*Options)) (*ListChannelsModeratedByAppInstanceUserOutput, error) 
 ListMediaCapturePipelines(ctx context.Context, params *ListMediaCapturePipelinesInput, optFns ...func(*Options)) (*ListMediaCapturePipelinesOutput, error) 
 ListMeetingTags(ctx context.Context, params *ListMeetingTagsInput, optFns ...func(*Options)) (*ListMeetingTagsOutput, error) 
 ListMeetings(ctx context.Context, params *ListMeetingsInput, optFns ...func(*Options)) (*ListMeetingsOutput, error) 
 ListPhoneNumberOrders(ctx context.Context, params *ListPhoneNumberOrdersInput, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) 
 ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) 
 ListProxySessions(ctx context.Context, params *ListProxySessionsInput, optFns ...func(*Options)) (*ListProxySessionsOutput, error) 
 ListRoomMemberships(ctx context.Context, params *ListRoomMembershipsInput, optFns ...func(*Options)) (*ListRoomMembershipsOutput, error) 
 ListRooms(ctx context.Context, params *ListRoomsInput, optFns ...func(*Options)) (*ListRoomsOutput, error) 
 ListSipMediaApplications(ctx context.Context, params *ListSipMediaApplicationsInput, optFns ...func(*Options)) (*ListSipMediaApplicationsOutput, error) 
 ListSipRules(ctx context.Context, params *ListSipRulesInput, optFns ...func(*Options)) (*ListSipRulesOutput, error) 
 ListSupportedPhoneNumberCountries(ctx context.Context, params *ListSupportedPhoneNumberCountriesInput, optFns ...func(*Options)) (*ListSupportedPhoneNumberCountriesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListVoiceConnectorGroups(ctx context.Context, params *ListVoiceConnectorGroupsInput, optFns ...func(*Options)) (*ListVoiceConnectorGroupsOutput, error) 
 ListVoiceConnectorTerminationCredentials(ctx context.Context, params *ListVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*ListVoiceConnectorTerminationCredentialsOutput, error) 
 ListVoiceConnectors(ctx context.Context, params *ListVoiceConnectorsInput, optFns ...func(*Options)) (*ListVoiceConnectorsOutput, error) 
 LogoutUser(ctx context.Context, params *LogoutUserInput, optFns ...func(*Options)) (*LogoutUserOutput, error) 
 PutAppInstanceRetentionSettings(ctx context.Context, params *PutAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*PutAppInstanceRetentionSettingsOutput, error) 
 PutAppInstanceStreamingConfigurations(ctx context.Context, params *PutAppInstanceStreamingConfigurationsInput, optFns ...func(*Options)) (*PutAppInstanceStreamingConfigurationsOutput, error) 
 PutEventsConfiguration(ctx context.Context, params *PutEventsConfigurationInput, optFns ...func(*Options)) (*PutEventsConfigurationOutput, error) 
 PutRetentionSettings(ctx context.Context, params *PutRetentionSettingsInput, optFns ...func(*Options)) (*PutRetentionSettingsOutput, error) 
 PutSipMediaApplicationLoggingConfiguration(ctx context.Context, params *PutSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*PutSipMediaApplicationLoggingConfigurationOutput, error) 
 PutVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *PutVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 PutVoiceConnectorLoggingConfiguration(ctx context.Context, params *PutVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorLoggingConfigurationOutput, error) 
 PutVoiceConnectorOrigination(ctx context.Context, params *PutVoiceConnectorOriginationInput, optFns ...func(*Options)) (*PutVoiceConnectorOriginationOutput, error) 
 PutVoiceConnectorProxy(ctx context.Context, params *PutVoiceConnectorProxyInput, optFns ...func(*Options)) (*PutVoiceConnectorProxyOutput, error) 
 PutVoiceConnectorStreamingConfiguration(ctx context.Context, params *PutVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorStreamingConfigurationOutput, error) 
 PutVoiceConnectorTermination(ctx context.Context, params *PutVoiceConnectorTerminationInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationOutput, error) 
 PutVoiceConnectorTerminationCredentials(ctx context.Context, params *PutVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationCredentialsOutput, error) 
 RedactChannelMessage(ctx context.Context, params *RedactChannelMessageInput, optFns ...func(*Options)) (*RedactChannelMessageOutput, error) 
 RedactConversationMessage(ctx context.Context, params *RedactConversationMessageInput, optFns ...func(*Options)) (*RedactConversationMessageOutput, error) 
 RedactRoomMessage(ctx context.Context, params *RedactRoomMessageInput, optFns ...func(*Options)) (*RedactRoomMessageOutput, error) 
 RegenerateSecurityToken(ctx context.Context, params *RegenerateSecurityTokenInput, optFns ...func(*Options)) (*RegenerateSecurityTokenOutput, error) 
 ResetPersonalPIN(ctx context.Context, params *ResetPersonalPINInput, optFns ...func(*Options)) (*ResetPersonalPINOutput, error) 
 RestorePhoneNumber(ctx context.Context, params *RestorePhoneNumberInput, optFns ...func(*Options)) (*RestorePhoneNumberOutput, error) 
 SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) 
 SendChannelMessage(ctx context.Context, params *SendChannelMessageInput, optFns ...func(*Options)) (*SendChannelMessageOutput, error) 
 StartMeetingTranscription(ctx context.Context, params *StartMeetingTranscriptionInput, optFns ...func(*Options)) (*StartMeetingTranscriptionOutput, error) 
 StopMeetingTranscription(ctx context.Context, params *StopMeetingTranscriptionInput, optFns ...func(*Options)) (*StopMeetingTranscriptionOutput, error) 
 TagAttendee(ctx context.Context, params *TagAttendeeInput, optFns ...func(*Options)) (*TagAttendeeOutput, error) 
 TagMeeting(ctx context.Context, params *TagMeetingInput, optFns ...func(*Options)) (*TagMeetingOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagAttendee(ctx context.Context, params *UntagAttendeeInput, optFns ...func(*Options)) (*UntagAttendeeOutput, error) 
 UntagMeeting(ctx context.Context, params *UntagMeetingInput, optFns ...func(*Options)) (*UntagMeetingOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccount(ctx context.Context, params *UpdateAccountInput, optFns ...func(*Options)) (*UpdateAccountOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateAppInstance(ctx context.Context, params *UpdateAppInstanceInput, optFns ...func(*Options)) (*UpdateAppInstanceOutput, error) 
 UpdateAppInstanceUser(ctx context.Context, params *UpdateAppInstanceUserInput, optFns ...func(*Options)) (*UpdateAppInstanceUserOutput, error) 
 UpdateBot(ctx context.Context, params *UpdateBotInput, optFns ...func(*Options)) (*UpdateBotOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateChannelMessage(ctx context.Context, params *UpdateChannelMessageInput, optFns ...func(*Options)) (*UpdateChannelMessageOutput, error) 
 UpdateChannelReadMarker(ctx context.Context, params *UpdateChannelReadMarkerInput, optFns ...func(*Options)) (*UpdateChannelReadMarkerOutput, error) 
 UpdateGlobalSettings(ctx context.Context, params *UpdateGlobalSettingsInput, optFns ...func(*Options)) (*UpdateGlobalSettingsOutput, error) 
 UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) 
 UpdatePhoneNumberSettings(ctx context.Context, params *UpdatePhoneNumberSettingsInput, optFns ...func(*Options)) (*UpdatePhoneNumberSettingsOutput, error) 
 UpdateProxySession(ctx context.Context, params *UpdateProxySessionInput, optFns ...func(*Options)) (*UpdateProxySessionOutput, error) 
 UpdateRoom(ctx context.Context, params *UpdateRoomInput, optFns ...func(*Options)) (*UpdateRoomOutput, error) 
 UpdateRoomMembership(ctx context.Context, params *UpdateRoomMembershipInput, optFns ...func(*Options)) (*UpdateRoomMembershipOutput, error) 
 UpdateSipMediaApplication(ctx context.Context, params *UpdateSipMediaApplicationInput, optFns ...func(*Options)) (*UpdateSipMediaApplicationOutput, error) 
 UpdateSipMediaApplicationCall(ctx context.Context, params *UpdateSipMediaApplicationCallInput, optFns ...func(*Options)) (*UpdateSipMediaApplicationCallOutput, error) 
 UpdateSipRule(ctx context.Context, params *UpdateSipRuleInput, optFns ...func(*Options)) (*UpdateSipRuleOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
 UpdateUserSettings(ctx context.Context, params *UpdateUserSettingsInput, optFns ...func(*Options)) (*UpdateUserSettingsOutput, error) 
 UpdateVoiceConnector(ctx context.Context, params *UpdateVoiceConnectorInput, optFns ...func(*Options)) (*UpdateVoiceConnectorOutput, error) 
 UpdateVoiceConnectorGroup(ctx context.Context, params *UpdateVoiceConnectorGroupInput, optFns ...func(*Options)) (*UpdateVoiceConnectorGroupOutput, error) 
 ValidateE911Address(ctx context.Context, params *ValidateE911AddressInput, optFns ...func(*Options)) (*ValidateE911AddressOutput, error) 
}
