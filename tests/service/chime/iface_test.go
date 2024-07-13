package chime_test

// tests for the chime service interface for this ../../../service/chime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chime/chime_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimeServiceCanBeMocked(t *testing.T) {
	var iface chime_iface.IClient
	iface = &chime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumberWithUser", func(t *testing.T) {
        input := &chime.AssociatePhoneNumberWithUserInput{}
        output := &chime.AssociatePhoneNumberWithUserOutput{}

        mockClient.On("AssociatePhoneNumberWithUser", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumberWithUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumbersWithVoiceConnector", func(t *testing.T) {
        input := &chime.AssociatePhoneNumbersWithVoiceConnectorInput{}
        output := &chime.AssociatePhoneNumbersWithVoiceConnectorOutput{}

        mockClient.On("AssociatePhoneNumbersWithVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumbersWithVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumbersWithVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput{}
        output := &chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput{}

        mockClient.On("AssociatePhoneNumbersWithVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumbersWithVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSigninDelegateGroupsWithAccount", func(t *testing.T) {
        input := &chime.AssociateSigninDelegateGroupsWithAccountInput{}
        output := &chime.AssociateSigninDelegateGroupsWithAccountOutput{}

        mockClient.On("AssociateSigninDelegateGroupsWithAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSigninDelegateGroupsWithAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateAttendee", func(t *testing.T) {
        input := &chime.BatchCreateAttendeeInput{}
        output := &chime.BatchCreateAttendeeOutput{}

        mockClient.On("BatchCreateAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateChannelMembership", func(t *testing.T) {
        input := &chime.BatchCreateChannelMembershipInput{}
        output := &chime.BatchCreateChannelMembershipOutput{}

        mockClient.On("BatchCreateChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateRoomMembership", func(t *testing.T) {
        input := &chime.BatchCreateRoomMembershipInput{}
        output := &chime.BatchCreateRoomMembershipOutput{}

        mockClient.On("BatchCreateRoomMembership", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateRoomMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeletePhoneNumber", func(t *testing.T) {
        input := &chime.BatchDeletePhoneNumberInput{}
        output := &chime.BatchDeletePhoneNumberOutput{}

        mockClient.On("BatchDeletePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeletePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchSuspendUser", func(t *testing.T) {
        input := &chime.BatchSuspendUserInput{}
        output := &chime.BatchSuspendUserOutput{}

        mockClient.On("BatchSuspendUser", ctx, input).Return(output, nil)

        result, err := mockClient.BatchSuspendUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUnsuspendUser", func(t *testing.T) {
        input := &chime.BatchUnsuspendUserInput{}
        output := &chime.BatchUnsuspendUserOutput{}

        mockClient.On("BatchUnsuspendUser", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUnsuspendUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdatePhoneNumber", func(t *testing.T) {
        input := &chime.BatchUpdatePhoneNumberInput{}
        output := &chime.BatchUpdatePhoneNumberOutput{}

        mockClient.On("BatchUpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateUser", func(t *testing.T) {
        input := &chime.BatchUpdateUserInput{}
        output := &chime.BatchUpdateUserOutput{}

        mockClient.On("BatchUpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccount", func(t *testing.T) {
        input := &chime.CreateAccountInput{}
        output := &chime.CreateAccountOutput{}

        mockClient.On("CreateAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstance", func(t *testing.T) {
        input := &chime.CreateAppInstanceInput{}
        output := &chime.CreateAppInstanceOutput{}

        mockClient.On("CreateAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstanceAdmin", func(t *testing.T) {
        input := &chime.CreateAppInstanceAdminInput{}
        output := &chime.CreateAppInstanceAdminOutput{}

        mockClient.On("CreateAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstanceUser", func(t *testing.T) {
        input := &chime.CreateAppInstanceUserInput{}
        output := &chime.CreateAppInstanceUserOutput{}

        mockClient.On("CreateAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAttendee", func(t *testing.T) {
        input := &chime.CreateAttendeeInput{}
        output := &chime.CreateAttendeeOutput{}

        mockClient.On("CreateAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBot", func(t *testing.T) {
        input := &chime.CreateBotInput{}
        output := &chime.CreateBotOutput{}

        mockClient.On("CreateBot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &chime.CreateChannelInput{}
        output := &chime.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelBan", func(t *testing.T) {
        input := &chime.CreateChannelBanInput{}
        output := &chime.CreateChannelBanOutput{}

        mockClient.On("CreateChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelMembership", func(t *testing.T) {
        input := &chime.CreateChannelMembershipInput{}
        output := &chime.CreateChannelMembershipOutput{}

        mockClient.On("CreateChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelModerator", func(t *testing.T) {
        input := &chime.CreateChannelModeratorInput{}
        output := &chime.CreateChannelModeratorOutput{}

        mockClient.On("CreateChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMediaCapturePipeline", func(t *testing.T) {
        input := &chime.CreateMediaCapturePipelineInput{}
        output := &chime.CreateMediaCapturePipelineOutput{}

        mockClient.On("CreateMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMeeting", func(t *testing.T) {
        input := &chime.CreateMeetingInput{}
        output := &chime.CreateMeetingOutput{}

        mockClient.On("CreateMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMeetingDialOut", func(t *testing.T) {
        input := &chime.CreateMeetingDialOutInput{}
        output := &chime.CreateMeetingDialOutOutput{}

        mockClient.On("CreateMeetingDialOut", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMeetingDialOut(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMeetingWithAttendees", func(t *testing.T) {
        input := &chime.CreateMeetingWithAttendeesInput{}
        output := &chime.CreateMeetingWithAttendeesOutput{}

        mockClient.On("CreateMeetingWithAttendees", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMeetingWithAttendees(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePhoneNumberOrder", func(t *testing.T) {
        input := &chime.CreatePhoneNumberOrderInput{}
        output := &chime.CreatePhoneNumberOrderOutput{}

        mockClient.On("CreatePhoneNumberOrder", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePhoneNumberOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProxySession", func(t *testing.T) {
        input := &chime.CreateProxySessionInput{}
        output := &chime.CreateProxySessionOutput{}

        mockClient.On("CreateProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoom", func(t *testing.T) {
        input := &chime.CreateRoomInput{}
        output := &chime.CreateRoomOutput{}

        mockClient.On("CreateRoom", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoomMembership", func(t *testing.T) {
        input := &chime.CreateRoomMembershipInput{}
        output := &chime.CreateRoomMembershipOutput{}

        mockClient.On("CreateRoomMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoomMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipMediaApplication", func(t *testing.T) {
        input := &chime.CreateSipMediaApplicationInput{}
        output := &chime.CreateSipMediaApplicationOutput{}

        mockClient.On("CreateSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipMediaApplicationCall", func(t *testing.T) {
        input := &chime.CreateSipMediaApplicationCallInput{}
        output := &chime.CreateSipMediaApplicationCallOutput{}

        mockClient.On("CreateSipMediaApplicationCall", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipMediaApplicationCall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipRule", func(t *testing.T) {
        input := &chime.CreateSipRuleInput{}
        output := &chime.CreateSipRuleOutput{}

        mockClient.On("CreateSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &chime.CreateUserInput{}
        output := &chime.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceConnector", func(t *testing.T) {
        input := &chime.CreateVoiceConnectorInput{}
        output := &chime.CreateVoiceConnectorOutput{}

        mockClient.On("CreateVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.CreateVoiceConnectorGroupInput{}
        output := &chime.CreateVoiceConnectorGroupOutput{}

        mockClient.On("CreateVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccount", func(t *testing.T) {
        input := &chime.DeleteAccountInput{}
        output := &chime.DeleteAccountOutput{}

        mockClient.On("DeleteAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstance", func(t *testing.T) {
        input := &chime.DeleteAppInstanceInput{}
        output := &chime.DeleteAppInstanceOutput{}

        mockClient.On("DeleteAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceAdmin", func(t *testing.T) {
        input := &chime.DeleteAppInstanceAdminInput{}
        output := &chime.DeleteAppInstanceAdminOutput{}

        mockClient.On("DeleteAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceStreamingConfigurations", func(t *testing.T) {
        input := &chime.DeleteAppInstanceStreamingConfigurationsInput{}
        output := &chime.DeleteAppInstanceStreamingConfigurationsOutput{}

        mockClient.On("DeleteAppInstanceStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceUser", func(t *testing.T) {
        input := &chime.DeleteAppInstanceUserInput{}
        output := &chime.DeleteAppInstanceUserOutput{}

        mockClient.On("DeleteAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttendee", func(t *testing.T) {
        input := &chime.DeleteAttendeeInput{}
        output := &chime.DeleteAttendeeOutput{}

        mockClient.On("DeleteAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &chime.DeleteChannelInput{}
        output := &chime.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelBan", func(t *testing.T) {
        input := &chime.DeleteChannelBanInput{}
        output := &chime.DeleteChannelBanOutput{}

        mockClient.On("DeleteChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelMembership", func(t *testing.T) {
        input := &chime.DeleteChannelMembershipInput{}
        output := &chime.DeleteChannelMembershipOutput{}

        mockClient.On("DeleteChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelMessage", func(t *testing.T) {
        input := &chime.DeleteChannelMessageInput{}
        output := &chime.DeleteChannelMessageOutput{}

        mockClient.On("DeleteChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelModerator", func(t *testing.T) {
        input := &chime.DeleteChannelModeratorInput{}
        output := &chime.DeleteChannelModeratorOutput{}

        mockClient.On("DeleteChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventsConfiguration", func(t *testing.T) {
        input := &chime.DeleteEventsConfigurationInput{}
        output := &chime.DeleteEventsConfigurationOutput{}

        mockClient.On("DeleteEventsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaCapturePipeline", func(t *testing.T) {
        input := &chime.DeleteMediaCapturePipelineInput{}
        output := &chime.DeleteMediaCapturePipelineOutput{}

        mockClient.On("DeleteMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMeeting", func(t *testing.T) {
        input := &chime.DeleteMeetingInput{}
        output := &chime.DeleteMeetingOutput{}

        mockClient.On("DeleteMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePhoneNumber", func(t *testing.T) {
        input := &chime.DeletePhoneNumberInput{}
        output := &chime.DeletePhoneNumberOutput{}

        mockClient.On("DeletePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProxySession", func(t *testing.T) {
        input := &chime.DeleteProxySessionInput{}
        output := &chime.DeleteProxySessionOutput{}

        mockClient.On("DeleteProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoom", func(t *testing.T) {
        input := &chime.DeleteRoomInput{}
        output := &chime.DeleteRoomOutput{}

        mockClient.On("DeleteRoom", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoomMembership", func(t *testing.T) {
        input := &chime.DeleteRoomMembershipInput{}
        output := &chime.DeleteRoomMembershipOutput{}

        mockClient.On("DeleteRoomMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoomMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSipMediaApplication", func(t *testing.T) {
        input := &chime.DeleteSipMediaApplicationInput{}
        output := &chime.DeleteSipMediaApplicationOutput{}

        mockClient.On("DeleteSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSipRule", func(t *testing.T) {
        input := &chime.DeleteSipRuleInput{}
        output := &chime.DeleteSipRuleOutput{}

        mockClient.On("DeleteSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnector", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorInput{}
        output := &chime.DeleteVoiceConnectorOutput{}

        mockClient.On("DeleteVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("DeleteVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorGroupInput{}
        output := &chime.DeleteVoiceConnectorGroupOutput{}

        mockClient.On("DeleteVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorOrigination", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorOriginationInput{}
        output := &chime.DeleteVoiceConnectorOriginationOutput{}

        mockClient.On("DeleteVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorProxy", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorProxyInput{}
        output := &chime.DeleteVoiceConnectorProxyOutput{}

        mockClient.On("DeleteVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorStreamingConfigurationInput{}
        output := &chime.DeleteVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("DeleteVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorTermination", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorTerminationInput{}
        output := &chime.DeleteVoiceConnectorTerminationOutput{}

        mockClient.On("DeleteVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chime.DeleteVoiceConnectorTerminationCredentialsInput{}
        output := &chime.DeleteVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("DeleteVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstance", func(t *testing.T) {
        input := &chime.DescribeAppInstanceInput{}
        output := &chime.DescribeAppInstanceOutput{}

        mockClient.On("DescribeAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceAdmin", func(t *testing.T) {
        input := &chime.DescribeAppInstanceAdminInput{}
        output := &chime.DescribeAppInstanceAdminOutput{}

        mockClient.On("DescribeAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceUser", func(t *testing.T) {
        input := &chime.DescribeAppInstanceUserInput{}
        output := &chime.DescribeAppInstanceUserOutput{}

        mockClient.On("DescribeAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &chime.DescribeChannelInput{}
        output := &chime.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelBan", func(t *testing.T) {
        input := &chime.DescribeChannelBanInput{}
        output := &chime.DescribeChannelBanOutput{}

        mockClient.On("DescribeChannelBan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelBan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelMembership", func(t *testing.T) {
        input := &chime.DescribeChannelMembershipInput{}
        output := &chime.DescribeChannelMembershipOutput{}

        mockClient.On("DescribeChannelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelMembershipForAppInstanceUser", func(t *testing.T) {
        input := &chime.DescribeChannelMembershipForAppInstanceUserInput{}
        output := &chime.DescribeChannelMembershipForAppInstanceUserOutput{}

        mockClient.On("DescribeChannelMembershipForAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelMembershipForAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelModeratedByAppInstanceUser", func(t *testing.T) {
        input := &chime.DescribeChannelModeratedByAppInstanceUserInput{}
        output := &chime.DescribeChannelModeratedByAppInstanceUserOutput{}

        mockClient.On("DescribeChannelModeratedByAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelModeratedByAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannelModerator", func(t *testing.T) {
        input := &chime.DescribeChannelModeratorInput{}
        output := &chime.DescribeChannelModeratorOutput{}

        mockClient.On("DescribeChannelModerator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannelModerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumberFromUser", func(t *testing.T) {
        input := &chime.DisassociatePhoneNumberFromUserInput{}
        output := &chime.DisassociatePhoneNumberFromUserOutput{}

        mockClient.On("DisassociatePhoneNumberFromUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumberFromUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumbersFromVoiceConnector", func(t *testing.T) {
        input := &chime.DisassociatePhoneNumbersFromVoiceConnectorInput{}
        output := &chime.DisassociatePhoneNumbersFromVoiceConnectorOutput{}

        mockClient.On("DisassociatePhoneNumbersFromVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumbersFromVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumbersFromVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput{}
        output := &chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput{}

        mockClient.On("DisassociatePhoneNumbersFromVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSigninDelegateGroupsFromAccount", func(t *testing.T) {
        input := &chime.DisassociateSigninDelegateGroupsFromAccountInput{}
        output := &chime.DisassociateSigninDelegateGroupsFromAccountOutput{}

        mockClient.On("DisassociateSigninDelegateGroupsFromAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSigninDelegateGroupsFromAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccount", func(t *testing.T) {
        input := &chime.GetAccountInput{}
        output := &chime.GetAccountOutput{}

        mockClient.On("GetAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &chime.GetAccountSettingsInput{}
        output := &chime.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppInstanceRetentionSettings", func(t *testing.T) {
        input := &chime.GetAppInstanceRetentionSettingsInput{}
        output := &chime.GetAppInstanceRetentionSettingsOutput{}

        mockClient.On("GetAppInstanceRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppInstanceRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppInstanceStreamingConfigurations", func(t *testing.T) {
        input := &chime.GetAppInstanceStreamingConfigurationsInput{}
        output := &chime.GetAppInstanceStreamingConfigurationsOutput{}

        mockClient.On("GetAppInstanceStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppInstanceStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttendee", func(t *testing.T) {
        input := &chime.GetAttendeeInput{}
        output := &chime.GetAttendeeOutput{}

        mockClient.On("GetAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBot", func(t *testing.T) {
        input := &chime.GetBotInput{}
        output := &chime.GetBotOutput{}

        mockClient.On("GetBot", ctx, input).Return(output, nil)

        result, err := mockClient.GetBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelMessage", func(t *testing.T) {
        input := &chime.GetChannelMessageInput{}
        output := &chime.GetChannelMessageOutput{}

        mockClient.On("GetChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventsConfiguration", func(t *testing.T) {
        input := &chime.GetEventsConfigurationInput{}
        output := &chime.GetEventsConfigurationOutput{}

        mockClient.On("GetEventsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGlobalSettings", func(t *testing.T) {
        input := &chime.GetGlobalSettingsInput{}
        output := &chime.GetGlobalSettingsOutput{}

        mockClient.On("GetGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaCapturePipeline", func(t *testing.T) {
        input := &chime.GetMediaCapturePipelineInput{}
        output := &chime.GetMediaCapturePipelineOutput{}

        mockClient.On("GetMediaCapturePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaCapturePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMeeting", func(t *testing.T) {
        input := &chime.GetMeetingInput{}
        output := &chime.GetMeetingOutput{}

        mockClient.On("GetMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.GetMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMessagingSessionEndpoint", func(t *testing.T) {
        input := &chime.GetMessagingSessionEndpointInput{}
        output := &chime.GetMessagingSessionEndpointOutput{}

        mockClient.On("GetMessagingSessionEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetMessagingSessionEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumber", func(t *testing.T) {
        input := &chime.GetPhoneNumberInput{}
        output := &chime.GetPhoneNumberOutput{}

        mockClient.On("GetPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumberOrder", func(t *testing.T) {
        input := &chime.GetPhoneNumberOrderInput{}
        output := &chime.GetPhoneNumberOrderOutput{}

        mockClient.On("GetPhoneNumberOrder", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumberOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumberSettings", func(t *testing.T) {
        input := &chime.GetPhoneNumberSettingsInput{}
        output := &chime.GetPhoneNumberSettingsOutput{}

        mockClient.On("GetPhoneNumberSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumberSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProxySession", func(t *testing.T) {
        input := &chime.GetProxySessionInput{}
        output := &chime.GetProxySessionOutput{}

        mockClient.On("GetProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.GetProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRetentionSettings", func(t *testing.T) {
        input := &chime.GetRetentionSettingsInput{}
        output := &chime.GetRetentionSettingsOutput{}

        mockClient.On("GetRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoom", func(t *testing.T) {
        input := &chime.GetRoomInput{}
        output := &chime.GetRoomOutput{}

        mockClient.On("GetRoom", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipMediaApplication", func(t *testing.T) {
        input := &chime.GetSipMediaApplicationInput{}
        output := &chime.GetSipMediaApplicationOutput{}

        mockClient.On("GetSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipMediaApplicationLoggingConfiguration", func(t *testing.T) {
        input := &chime.GetSipMediaApplicationLoggingConfigurationInput{}
        output := &chime.GetSipMediaApplicationLoggingConfigurationOutput{}

        mockClient.On("GetSipMediaApplicationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipMediaApplicationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipRule", func(t *testing.T) {
        input := &chime.GetSipRuleInput{}
        output := &chime.GetSipRuleOutput{}

        mockClient.On("GetSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUser", func(t *testing.T) {
        input := &chime.GetUserInput{}
        output := &chime.GetUserOutput{}

        mockClient.On("GetUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserSettings", func(t *testing.T) {
        input := &chime.GetUserSettingsInput{}
        output := &chime.GetUserSettingsOutput{}

        mockClient.On("GetUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnector", func(t *testing.T) {
        input := &chime.GetVoiceConnectorInput{}
        output := &chime.GetVoiceConnectorOutput{}

        mockClient.On("GetVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chime.GetVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chime.GetVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.GetVoiceConnectorGroupInput{}
        output := &chime.GetVoiceConnectorGroupOutput{}

        mockClient.On("GetVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorLoggingConfiguration", func(t *testing.T) {
        input := &chime.GetVoiceConnectorLoggingConfigurationInput{}
        output := &chime.GetVoiceConnectorLoggingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorOrigination", func(t *testing.T) {
        input := &chime.GetVoiceConnectorOriginationInput{}
        output := &chime.GetVoiceConnectorOriginationOutput{}

        mockClient.On("GetVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorProxy", func(t *testing.T) {
        input := &chime.GetVoiceConnectorProxyInput{}
        output := &chime.GetVoiceConnectorProxyOutput{}

        mockClient.On("GetVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chime.GetVoiceConnectorStreamingConfigurationInput{}
        output := &chime.GetVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorTermination", func(t *testing.T) {
        input := &chime.GetVoiceConnectorTerminationInput{}
        output := &chime.GetVoiceConnectorTerminationOutput{}

        mockClient.On("GetVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorTerminationHealth", func(t *testing.T) {
        input := &chime.GetVoiceConnectorTerminationHealthInput{}
        output := &chime.GetVoiceConnectorTerminationHealthOutput{}

        mockClient.On("GetVoiceConnectorTerminationHealth", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorTerminationHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInviteUsers", func(t *testing.T) {
        input := &chime.InviteUsersInput{}
        output := &chime.InviteUsersOutput{}

        mockClient.On("InviteUsers", ctx, input).Return(output, nil)

        result, err := mockClient.InviteUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccounts", func(t *testing.T) {
        input := &chime.ListAccountsInput{}
        output := &chime.ListAccountsOutput{}

        mockClient.On("ListAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceAdmins", func(t *testing.T) {
        input := &chime.ListAppInstanceAdminsInput{}
        output := &chime.ListAppInstanceAdminsOutput{}

        mockClient.On("ListAppInstanceAdmins", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceAdmins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceUsers", func(t *testing.T) {
        input := &chime.ListAppInstanceUsersInput{}
        output := &chime.ListAppInstanceUsersOutput{}

        mockClient.On("ListAppInstanceUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstances", func(t *testing.T) {
        input := &chime.ListAppInstancesInput{}
        output := &chime.ListAppInstancesOutput{}

        mockClient.On("ListAppInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttendeeTags", func(t *testing.T) {
        input := &chime.ListAttendeeTagsInput{}
        output := &chime.ListAttendeeTagsOutput{}

        mockClient.On("ListAttendeeTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttendeeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttendees", func(t *testing.T) {
        input := &chime.ListAttendeesInput{}
        output := &chime.ListAttendeesOutput{}

        mockClient.On("ListAttendees", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttendees(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBots", func(t *testing.T) {
        input := &chime.ListBotsInput{}
        output := &chime.ListBotsOutput{}

        mockClient.On("ListBots", ctx, input).Return(output, nil)

        result, err := mockClient.ListBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelBans", func(t *testing.T) {
        input := &chime.ListChannelBansInput{}
        output := &chime.ListChannelBansOutput{}

        mockClient.On("ListChannelBans", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelBans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMemberships", func(t *testing.T) {
        input := &chime.ListChannelMembershipsInput{}
        output := &chime.ListChannelMembershipsOutput{}

        mockClient.On("ListChannelMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMembershipsForAppInstanceUser", func(t *testing.T) {
        input := &chime.ListChannelMembershipsForAppInstanceUserInput{}
        output := &chime.ListChannelMembershipsForAppInstanceUserOutput{}

        mockClient.On("ListChannelMembershipsForAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMembershipsForAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelMessages", func(t *testing.T) {
        input := &chime.ListChannelMessagesInput{}
        output := &chime.ListChannelMessagesOutput{}

        mockClient.On("ListChannelMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelModerators", func(t *testing.T) {
        input := &chime.ListChannelModeratorsInput{}
        output := &chime.ListChannelModeratorsOutput{}

        mockClient.On("ListChannelModerators", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelModerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &chime.ListChannelsInput{}
        output := &chime.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelsModeratedByAppInstanceUser", func(t *testing.T) {
        input := &chime.ListChannelsModeratedByAppInstanceUserInput{}
        output := &chime.ListChannelsModeratedByAppInstanceUserOutput{}

        mockClient.On("ListChannelsModeratedByAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelsModeratedByAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaCapturePipelines", func(t *testing.T) {
        input := &chime.ListMediaCapturePipelinesInput{}
        output := &chime.ListMediaCapturePipelinesOutput{}

        mockClient.On("ListMediaCapturePipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaCapturePipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMeetingTags", func(t *testing.T) {
        input := &chime.ListMeetingTagsInput{}
        output := &chime.ListMeetingTagsOutput{}

        mockClient.On("ListMeetingTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListMeetingTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMeetings", func(t *testing.T) {
        input := &chime.ListMeetingsInput{}
        output := &chime.ListMeetingsOutput{}

        mockClient.On("ListMeetings", ctx, input).Return(output, nil)

        result, err := mockClient.ListMeetings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumberOrders", func(t *testing.T) {
        input := &chime.ListPhoneNumberOrdersInput{}
        output := &chime.ListPhoneNumberOrdersOutput{}

        mockClient.On("ListPhoneNumberOrders", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumberOrders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumbers", func(t *testing.T) {
        input := &chime.ListPhoneNumbersInput{}
        output := &chime.ListPhoneNumbersOutput{}

        mockClient.On("ListPhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProxySessions", func(t *testing.T) {
        input := &chime.ListProxySessionsInput{}
        output := &chime.ListProxySessionsOutput{}

        mockClient.On("ListProxySessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListProxySessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoomMemberships", func(t *testing.T) {
        input := &chime.ListRoomMembershipsInput{}
        output := &chime.ListRoomMembershipsOutput{}

        mockClient.On("ListRoomMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoomMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRooms", func(t *testing.T) {
        input := &chime.ListRoomsInput{}
        output := &chime.ListRoomsOutput{}

        mockClient.On("ListRooms", ctx, input).Return(output, nil)

        result, err := mockClient.ListRooms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSipMediaApplications", func(t *testing.T) {
        input := &chime.ListSipMediaApplicationsInput{}
        output := &chime.ListSipMediaApplicationsOutput{}

        mockClient.On("ListSipMediaApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListSipMediaApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSipRules", func(t *testing.T) {
        input := &chime.ListSipRulesInput{}
        output := &chime.ListSipRulesOutput{}

        mockClient.On("ListSipRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListSipRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSupportedPhoneNumberCountries", func(t *testing.T) {
        input := &chime.ListSupportedPhoneNumberCountriesInput{}
        output := &chime.ListSupportedPhoneNumberCountriesOutput{}

        mockClient.On("ListSupportedPhoneNumberCountries", ctx, input).Return(output, nil)

        result, err := mockClient.ListSupportedPhoneNumberCountries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chime.ListTagsForResourceInput{}
        output := &chime.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &chime.ListUsersInput{}
        output := &chime.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectorGroups", func(t *testing.T) {
        input := &chime.ListVoiceConnectorGroupsInput{}
        output := &chime.ListVoiceConnectorGroupsOutput{}

        mockClient.On("ListVoiceConnectorGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectorGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chime.ListVoiceConnectorTerminationCredentialsInput{}
        output := &chime.ListVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("ListVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectors", func(t *testing.T) {
        input := &chime.ListVoiceConnectorsInput{}
        output := &chime.ListVoiceConnectorsOutput{}

        mockClient.On("ListVoiceConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLogoutUser", func(t *testing.T) {
        input := &chime.LogoutUserInput{}
        output := &chime.LogoutUserOutput{}

        mockClient.On("LogoutUser", ctx, input).Return(output, nil)

        result, err := mockClient.LogoutUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppInstanceRetentionSettings", func(t *testing.T) {
        input := &chime.PutAppInstanceRetentionSettingsInput{}
        output := &chime.PutAppInstanceRetentionSettingsOutput{}

        mockClient.On("PutAppInstanceRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppInstanceRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppInstanceStreamingConfigurations", func(t *testing.T) {
        input := &chime.PutAppInstanceStreamingConfigurationsInput{}
        output := &chime.PutAppInstanceStreamingConfigurationsOutput{}

        mockClient.On("PutAppInstanceStreamingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppInstanceStreamingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEventsConfiguration", func(t *testing.T) {
        input := &chime.PutEventsConfigurationInput{}
        output := &chime.PutEventsConfigurationOutput{}

        mockClient.On("PutEventsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutEventsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRetentionSettings", func(t *testing.T) {
        input := &chime.PutRetentionSettingsInput{}
        output := &chime.PutRetentionSettingsOutput{}

        mockClient.On("PutRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSipMediaApplicationLoggingConfiguration", func(t *testing.T) {
        input := &chime.PutSipMediaApplicationLoggingConfigurationInput{}
        output := &chime.PutSipMediaApplicationLoggingConfigurationOutput{}

        mockClient.On("PutSipMediaApplicationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutSipMediaApplicationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chime.PutVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chime.PutVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorLoggingConfiguration", func(t *testing.T) {
        input := &chime.PutVoiceConnectorLoggingConfigurationInput{}
        output := &chime.PutVoiceConnectorLoggingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorOrigination", func(t *testing.T) {
        input := &chime.PutVoiceConnectorOriginationInput{}
        output := &chime.PutVoiceConnectorOriginationOutput{}

        mockClient.On("PutVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorProxy", func(t *testing.T) {
        input := &chime.PutVoiceConnectorProxyInput{}
        output := &chime.PutVoiceConnectorProxyOutput{}

        mockClient.On("PutVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chime.PutVoiceConnectorStreamingConfigurationInput{}
        output := &chime.PutVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorTermination", func(t *testing.T) {
        input := &chime.PutVoiceConnectorTerminationInput{}
        output := &chime.PutVoiceConnectorTerminationOutput{}

        mockClient.On("PutVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chime.PutVoiceConnectorTerminationCredentialsInput{}
        output := &chime.PutVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("PutVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRedactChannelMessage", func(t *testing.T) {
        input := &chime.RedactChannelMessageInput{}
        output := &chime.RedactChannelMessageOutput{}

        mockClient.On("RedactChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.RedactChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRedactConversationMessage", func(t *testing.T) {
        input := &chime.RedactConversationMessageInput{}
        output := &chime.RedactConversationMessageOutput{}

        mockClient.On("RedactConversationMessage", ctx, input).Return(output, nil)

        result, err := mockClient.RedactConversationMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRedactRoomMessage", func(t *testing.T) {
        input := &chime.RedactRoomMessageInput{}
        output := &chime.RedactRoomMessageOutput{}

        mockClient.On("RedactRoomMessage", ctx, input).Return(output, nil)

        result, err := mockClient.RedactRoomMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegenerateSecurityToken", func(t *testing.T) {
        input := &chime.RegenerateSecurityTokenInput{}
        output := &chime.RegenerateSecurityTokenOutput{}

        mockClient.On("RegenerateSecurityToken", ctx, input).Return(output, nil)

        result, err := mockClient.RegenerateSecurityToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetPersonalPIN", func(t *testing.T) {
        input := &chime.ResetPersonalPINInput{}
        output := &chime.ResetPersonalPINOutput{}

        mockClient.On("ResetPersonalPIN", ctx, input).Return(output, nil)

        result, err := mockClient.ResetPersonalPIN(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestorePhoneNumber", func(t *testing.T) {
        input := &chime.RestorePhoneNumberInput{}
        output := &chime.RestorePhoneNumberOutput{}

        mockClient.On("RestorePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.RestorePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAvailablePhoneNumbers", func(t *testing.T) {
        input := &chime.SearchAvailablePhoneNumbersInput{}
        output := &chime.SearchAvailablePhoneNumbersOutput{}

        mockClient.On("SearchAvailablePhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAvailablePhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendChannelMessage", func(t *testing.T) {
        input := &chime.SendChannelMessageInput{}
        output := &chime.SendChannelMessageOutput{}

        mockClient.On("SendChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMeetingTranscription", func(t *testing.T) {
        input := &chime.StartMeetingTranscriptionInput{}
        output := &chime.StartMeetingTranscriptionOutput{}

        mockClient.On("StartMeetingTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartMeetingTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMeetingTranscription", func(t *testing.T) {
        input := &chime.StopMeetingTranscriptionInput{}
        output := &chime.StopMeetingTranscriptionOutput{}

        mockClient.On("StopMeetingTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StopMeetingTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagAttendee", func(t *testing.T) {
        input := &chime.TagAttendeeInput{}
        output := &chime.TagAttendeeOutput{}

        mockClient.On("TagAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.TagAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagMeeting", func(t *testing.T) {
        input := &chime.TagMeetingInput{}
        output := &chime.TagMeetingOutput{}

        mockClient.On("TagMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.TagMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chime.TagResourceInput{}
        output := &chime.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagAttendee", func(t *testing.T) {
        input := &chime.UntagAttendeeInput{}
        output := &chime.UntagAttendeeOutput{}

        mockClient.On("UntagAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.UntagAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagMeeting", func(t *testing.T) {
        input := &chime.UntagMeetingInput{}
        output := &chime.UntagMeetingOutput{}

        mockClient.On("UntagMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.UntagMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chime.UntagResourceInput{}
        output := &chime.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccount", func(t *testing.T) {
        input := &chime.UpdateAccountInput{}
        output := &chime.UpdateAccountOutput{}

        mockClient.On("UpdateAccount", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &chime.UpdateAccountSettingsInput{}
        output := &chime.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstance", func(t *testing.T) {
        input := &chime.UpdateAppInstanceInput{}
        output := &chime.UpdateAppInstanceOutput{}

        mockClient.On("UpdateAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstanceUser", func(t *testing.T) {
        input := &chime.UpdateAppInstanceUserInput{}
        output := &chime.UpdateAppInstanceUserOutput{}

        mockClient.On("UpdateAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBot", func(t *testing.T) {
        input := &chime.UpdateBotInput{}
        output := &chime.UpdateBotOutput{}

        mockClient.On("UpdateBot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &chime.UpdateChannelInput{}
        output := &chime.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelMessage", func(t *testing.T) {
        input := &chime.UpdateChannelMessageInput{}
        output := &chime.UpdateChannelMessageOutput{}

        mockClient.On("UpdateChannelMessage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelReadMarker", func(t *testing.T) {
        input := &chime.UpdateChannelReadMarkerInput{}
        output := &chime.UpdateChannelReadMarkerOutput{}

        mockClient.On("UpdateChannelReadMarker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelReadMarker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalSettings", func(t *testing.T) {
        input := &chime.UpdateGlobalSettingsInput{}
        output := &chime.UpdateGlobalSettingsOutput{}

        mockClient.On("UpdateGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumber", func(t *testing.T) {
        input := &chime.UpdatePhoneNumberInput{}
        output := &chime.UpdatePhoneNumberOutput{}

        mockClient.On("UpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumberSettings", func(t *testing.T) {
        input := &chime.UpdatePhoneNumberSettingsInput{}
        output := &chime.UpdatePhoneNumberSettingsOutput{}

        mockClient.On("UpdatePhoneNumberSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumberSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProxySession", func(t *testing.T) {
        input := &chime.UpdateProxySessionInput{}
        output := &chime.UpdateProxySessionOutput{}

        mockClient.On("UpdateProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoom", func(t *testing.T) {
        input := &chime.UpdateRoomInput{}
        output := &chime.UpdateRoomOutput{}

        mockClient.On("UpdateRoom", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoomMembership", func(t *testing.T) {
        input := &chime.UpdateRoomMembershipInput{}
        output := &chime.UpdateRoomMembershipOutput{}

        mockClient.On("UpdateRoomMembership", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoomMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipMediaApplication", func(t *testing.T) {
        input := &chime.UpdateSipMediaApplicationInput{}
        output := &chime.UpdateSipMediaApplicationOutput{}

        mockClient.On("UpdateSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipMediaApplicationCall", func(t *testing.T) {
        input := &chime.UpdateSipMediaApplicationCallInput{}
        output := &chime.UpdateSipMediaApplicationCallOutput{}

        mockClient.On("UpdateSipMediaApplicationCall", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipMediaApplicationCall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipRule", func(t *testing.T) {
        input := &chime.UpdateSipRuleInput{}
        output := &chime.UpdateSipRuleOutput{}

        mockClient.On("UpdateSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &chime.UpdateUserInput{}
        output := &chime.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserSettings", func(t *testing.T) {
        input := &chime.UpdateUserSettingsInput{}
        output := &chime.UpdateUserSettingsOutput{}

        mockClient.On("UpdateUserSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceConnector", func(t *testing.T) {
        input := &chime.UpdateVoiceConnectorInput{}
        output := &chime.UpdateVoiceConnectorOutput{}

        mockClient.On("UpdateVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceConnectorGroup", func(t *testing.T) {
        input := &chime.UpdateVoiceConnectorGroupInput{}
        output := &chime.UpdateVoiceConnectorGroupOutput{}

        mockClient.On("UpdateVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateE911Address", func(t *testing.T) {
        input := &chime.ValidateE911AddressInput{}
        output := &chime.ValidateE911AddressOutput{}

        mockClient.On("ValidateE911Address", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateE911Address(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
