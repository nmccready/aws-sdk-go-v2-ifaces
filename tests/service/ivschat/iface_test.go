package ivschat_test

// tests for the ivschat service interface for this ../../../service/ivschat/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ivschat"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivschat/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivschat/ivschat_iface"
	"github.com/stretchr/testify/assert"
)

func TestIvschatServiceCanBeMocked(t *testing.T) {
	var iface ivschat_iface.IClient
	iface = &ivschat.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ivschat.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChatToken", func(t *testing.T) {
        input := &ivschat.CreateChatTokenInput{}
        output := &ivschat.CreateChatTokenOutput{}

        mockClient.On("CreateChatToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChatToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoggingConfiguration", func(t *testing.T) {
        input := &ivschat.CreateLoggingConfigurationInput{}
        output := &ivschat.CreateLoggingConfigurationOutput{}

        mockClient.On("CreateLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoom", func(t *testing.T) {
        input := &ivschat.CreateRoomInput{}
        output := &ivschat.CreateRoomOutput{}

        mockClient.On("CreateRoom", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggingConfiguration", func(t *testing.T) {
        input := &ivschat.DeleteLoggingConfigurationInput{}
        output := &ivschat.DeleteLoggingConfigurationOutput{}

        mockClient.On("DeleteLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMessage", func(t *testing.T) {
        input := &ivschat.DeleteMessageInput{}
        output := &ivschat.DeleteMessageOutput{}

        mockClient.On("DeleteMessage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoom", func(t *testing.T) {
        input := &ivschat.DeleteRoomInput{}
        output := &ivschat.DeleteRoomOutput{}

        mockClient.On("DeleteRoom", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectUser", func(t *testing.T) {
        input := &ivschat.DisconnectUserInput{}
        output := &ivschat.DisconnectUserOutput{}

        mockClient.On("DisconnectUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingConfiguration", func(t *testing.T) {
        input := &ivschat.GetLoggingConfigurationInput{}
        output := &ivschat.GetLoggingConfigurationOutput{}

        mockClient.On("GetLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoom", func(t *testing.T) {
        input := &ivschat.GetRoomInput{}
        output := &ivschat.GetRoomOutput{}

        mockClient.On("GetRoom", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggingConfigurations", func(t *testing.T) {
        input := &ivschat.ListLoggingConfigurationsInput{}
        output := &ivschat.ListLoggingConfigurationsOutput{}

        mockClient.On("ListLoggingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRooms", func(t *testing.T) {
        input := &ivschat.ListRoomsInput{}
        output := &ivschat.ListRoomsOutput{}

        mockClient.On("ListRooms", ctx, input).Return(output, nil)

        result, err := mockClient.ListRooms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ivschat.ListTagsForResourceInput{}
        output := &ivschat.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEvent", func(t *testing.T) {
        input := &ivschat.SendEventInput{}
        output := &ivschat.SendEventOutput{}

        mockClient.On("SendEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SendEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ivschat.TagResourceInput{}
        output := &ivschat.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ivschat.UntagResourceInput{}
        output := &ivschat.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoggingConfiguration", func(t *testing.T) {
        input := &ivschat.UpdateLoggingConfigurationInput{}
        output := &ivschat.UpdateLoggingConfigurationOutput{}

        mockClient.On("UpdateLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoom", func(t *testing.T) {
        input := &ivschat.UpdateRoomInput{}
        output := &ivschat.UpdateRoomOutput{}

        mockClient.On("UpdateRoom", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
