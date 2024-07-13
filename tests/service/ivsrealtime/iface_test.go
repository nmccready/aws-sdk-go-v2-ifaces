package ivsrealtime_test

// tests for the ivsrealtime service interface for this ../../../service/ivsrealtime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivsrealtime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivsrealtime/ivsrealtime_iface"
	"github.com/stretchr/testify/assert"
)

func TestIvsrealtimeServiceCanBeMocked(t *testing.T) {
	var iface ivsrealtime_iface.IClient
	iface = &ivsrealtime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ivsrealtime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEncoderConfiguration", func(t *testing.T) {
        input := &ivsrealtime.CreateEncoderConfigurationInput{}
        output := &ivsrealtime.CreateEncoderConfigurationOutput{}

        mockClient.On("CreateEncoderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEncoderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParticipantToken", func(t *testing.T) {
        input := &ivsrealtime.CreateParticipantTokenInput{}
        output := &ivsrealtime.CreateParticipantTokenOutput{}

        mockClient.On("CreateParticipantToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParticipantToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStage", func(t *testing.T) {
        input := &ivsrealtime.CreateStageInput{}
        output := &ivsrealtime.CreateStageOutput{}

        mockClient.On("CreateStage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorageConfiguration", func(t *testing.T) {
        input := &ivsrealtime.CreateStorageConfigurationInput{}
        output := &ivsrealtime.CreateStorageConfigurationOutput{}

        mockClient.On("CreateStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEncoderConfiguration", func(t *testing.T) {
        input := &ivsrealtime.DeleteEncoderConfigurationInput{}
        output := &ivsrealtime.DeleteEncoderConfigurationOutput{}

        mockClient.On("DeleteEncoderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEncoderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublicKey", func(t *testing.T) {
        input := &ivsrealtime.DeletePublicKeyInput{}
        output := &ivsrealtime.DeletePublicKeyOutput{}

        mockClient.On("DeletePublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStage", func(t *testing.T) {
        input := &ivsrealtime.DeleteStageInput{}
        output := &ivsrealtime.DeleteStageOutput{}

        mockClient.On("DeleteStage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageConfiguration", func(t *testing.T) {
        input := &ivsrealtime.DeleteStorageConfigurationInput{}
        output := &ivsrealtime.DeleteStorageConfigurationOutput{}

        mockClient.On("DeleteStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectParticipant", func(t *testing.T) {
        input := &ivsrealtime.DisconnectParticipantInput{}
        output := &ivsrealtime.DisconnectParticipantOutput{}

        mockClient.On("DisconnectParticipant", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectParticipant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComposition", func(t *testing.T) {
        input := &ivsrealtime.GetCompositionInput{}
        output := &ivsrealtime.GetCompositionOutput{}

        mockClient.On("GetComposition", ctx, input).Return(output, nil)

        result, err := mockClient.GetComposition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEncoderConfiguration", func(t *testing.T) {
        input := &ivsrealtime.GetEncoderConfigurationInput{}
        output := &ivsrealtime.GetEncoderConfigurationOutput{}

        mockClient.On("GetEncoderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEncoderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParticipant", func(t *testing.T) {
        input := &ivsrealtime.GetParticipantInput{}
        output := &ivsrealtime.GetParticipantOutput{}

        mockClient.On("GetParticipant", ctx, input).Return(output, nil)

        result, err := mockClient.GetParticipant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicKey", func(t *testing.T) {
        input := &ivsrealtime.GetPublicKeyInput{}
        output := &ivsrealtime.GetPublicKeyOutput{}

        mockClient.On("GetPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStage", func(t *testing.T) {
        input := &ivsrealtime.GetStageInput{}
        output := &ivsrealtime.GetStageOutput{}

        mockClient.On("GetStage", ctx, input).Return(output, nil)

        result, err := mockClient.GetStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStageSession", func(t *testing.T) {
        input := &ivsrealtime.GetStageSessionInput{}
        output := &ivsrealtime.GetStageSessionOutput{}

        mockClient.On("GetStageSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetStageSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageConfiguration", func(t *testing.T) {
        input := &ivsrealtime.GetStorageConfigurationInput{}
        output := &ivsrealtime.GetStorageConfigurationOutput{}

        mockClient.On("GetStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportPublicKey", func(t *testing.T) {
        input := &ivsrealtime.ImportPublicKeyInput{}
        output := &ivsrealtime.ImportPublicKeyOutput{}

        mockClient.On("ImportPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.ImportPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCompositions", func(t *testing.T) {
        input := &ivsrealtime.ListCompositionsInput{}
        output := &ivsrealtime.ListCompositionsOutput{}

        mockClient.On("ListCompositions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCompositions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEncoderConfigurations", func(t *testing.T) {
        input := &ivsrealtime.ListEncoderConfigurationsInput{}
        output := &ivsrealtime.ListEncoderConfigurationsOutput{}

        mockClient.On("ListEncoderConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEncoderConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParticipantEvents", func(t *testing.T) {
        input := &ivsrealtime.ListParticipantEventsInput{}
        output := &ivsrealtime.ListParticipantEventsOutput{}

        mockClient.On("ListParticipantEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListParticipantEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParticipants", func(t *testing.T) {
        input := &ivsrealtime.ListParticipantsInput{}
        output := &ivsrealtime.ListParticipantsOutput{}

        mockClient.On("ListParticipants", ctx, input).Return(output, nil)

        result, err := mockClient.ListParticipants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPublicKeys", func(t *testing.T) {
        input := &ivsrealtime.ListPublicKeysInput{}
        output := &ivsrealtime.ListPublicKeysOutput{}

        mockClient.On("ListPublicKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListPublicKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStageSessions", func(t *testing.T) {
        input := &ivsrealtime.ListStageSessionsInput{}
        output := &ivsrealtime.ListStageSessionsOutput{}

        mockClient.On("ListStageSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListStageSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStages", func(t *testing.T) {
        input := &ivsrealtime.ListStagesInput{}
        output := &ivsrealtime.ListStagesOutput{}

        mockClient.On("ListStages", ctx, input).Return(output, nil)

        result, err := mockClient.ListStages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageConfigurations", func(t *testing.T) {
        input := &ivsrealtime.ListStorageConfigurationsInput{}
        output := &ivsrealtime.ListStorageConfigurationsOutput{}

        mockClient.On("ListStorageConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ivsrealtime.ListTagsForResourceInput{}
        output := &ivsrealtime.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartComposition", func(t *testing.T) {
        input := &ivsrealtime.StartCompositionInput{}
        output := &ivsrealtime.StartCompositionOutput{}

        mockClient.On("StartComposition", ctx, input).Return(output, nil)

        result, err := mockClient.StartComposition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopComposition", func(t *testing.T) {
        input := &ivsrealtime.StopCompositionInput{}
        output := &ivsrealtime.StopCompositionOutput{}

        mockClient.On("StopComposition", ctx, input).Return(output, nil)

        result, err := mockClient.StopComposition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ivsrealtime.TagResourceInput{}
        output := &ivsrealtime.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ivsrealtime.UntagResourceInput{}
        output := &ivsrealtime.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStage", func(t *testing.T) {
        input := &ivsrealtime.UpdateStageInput{}
        output := &ivsrealtime.UpdateStageOutput{}

        mockClient.On("UpdateStage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
