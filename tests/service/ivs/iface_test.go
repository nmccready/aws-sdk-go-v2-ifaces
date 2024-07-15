// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package ivs_test

// tests for the ivs service interface for 
// this ../../../service/ivs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ivs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ivs/ivs_iface"
	"github.com/stretchr/testify/assert"
)

func TestIvsServiceCanBeMocked(t *testing.T) {
	var iface ivs_iface.IClient
	iface = &ivs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ivs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetChannel", func(t *testing.T) {
        input := &ivs.BatchGetChannelInput{}
        output := &ivs.BatchGetChannelOutput{}

        mockClient.On("BatchGetChannel", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetStreamKey", func(t *testing.T) {
        input := &ivs.BatchGetStreamKeyInput{}
        output := &ivs.BatchGetStreamKeyOutput{}

        mockClient.On("BatchGetStreamKey", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetStreamKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStartViewerSessionRevocation", func(t *testing.T) {
        input := &ivs.BatchStartViewerSessionRevocationInput{}
        output := &ivs.BatchStartViewerSessionRevocationOutput{}

        mockClient.On("BatchStartViewerSessionRevocation", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStartViewerSessionRevocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &ivs.CreateChannelInput{}
        output := &ivs.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlaybackRestrictionPolicy", func(t *testing.T) {
        input := &ivs.CreatePlaybackRestrictionPolicyInput{}
        output := &ivs.CreatePlaybackRestrictionPolicyOutput{}

        mockClient.On("CreatePlaybackRestrictionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlaybackRestrictionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecordingConfiguration", func(t *testing.T) {
        input := &ivs.CreateRecordingConfigurationInput{}
        output := &ivs.CreateRecordingConfigurationOutput{}

        mockClient.On("CreateRecordingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecordingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamKey", func(t *testing.T) {
        input := &ivs.CreateStreamKeyInput{}
        output := &ivs.CreateStreamKeyOutput{}

        mockClient.On("CreateStreamKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &ivs.DeleteChannelInput{}
        output := &ivs.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlaybackKeyPair", func(t *testing.T) {
        input := &ivs.DeletePlaybackKeyPairInput{}
        output := &ivs.DeletePlaybackKeyPairOutput{}

        mockClient.On("DeletePlaybackKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlaybackKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlaybackRestrictionPolicy", func(t *testing.T) {
        input := &ivs.DeletePlaybackRestrictionPolicyInput{}
        output := &ivs.DeletePlaybackRestrictionPolicyOutput{}

        mockClient.On("DeletePlaybackRestrictionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlaybackRestrictionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecordingConfiguration", func(t *testing.T) {
        input := &ivs.DeleteRecordingConfigurationInput{}
        output := &ivs.DeleteRecordingConfigurationOutput{}

        mockClient.On("DeleteRecordingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecordingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStreamKey", func(t *testing.T) {
        input := &ivs.DeleteStreamKeyInput{}
        output := &ivs.DeleteStreamKeyOutput{}

        mockClient.On("DeleteStreamKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStreamKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannel", func(t *testing.T) {
        input := &ivs.GetChannelInput{}
        output := &ivs.GetChannelOutput{}

        mockClient.On("GetChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlaybackKeyPair", func(t *testing.T) {
        input := &ivs.GetPlaybackKeyPairInput{}
        output := &ivs.GetPlaybackKeyPairOutput{}

        mockClient.On("GetPlaybackKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlaybackKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlaybackRestrictionPolicy", func(t *testing.T) {
        input := &ivs.GetPlaybackRestrictionPolicyInput{}
        output := &ivs.GetPlaybackRestrictionPolicyOutput{}

        mockClient.On("GetPlaybackRestrictionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlaybackRestrictionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecordingConfiguration", func(t *testing.T) {
        input := &ivs.GetRecordingConfigurationInput{}
        output := &ivs.GetRecordingConfigurationOutput{}

        mockClient.On("GetRecordingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecordingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStream", func(t *testing.T) {
        input := &ivs.GetStreamInput{}
        output := &ivs.GetStreamOutput{}

        mockClient.On("GetStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamKey", func(t *testing.T) {
        input := &ivs.GetStreamKeyInput{}
        output := &ivs.GetStreamKeyOutput{}

        mockClient.On("GetStreamKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamSession", func(t *testing.T) {
        input := &ivs.GetStreamSessionInput{}
        output := &ivs.GetStreamSessionOutput{}

        mockClient.On("GetStreamSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportPlaybackKeyPair", func(t *testing.T) {
        input := &ivs.ImportPlaybackKeyPairInput{}
        output := &ivs.ImportPlaybackKeyPairOutput{}

        mockClient.On("ImportPlaybackKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.ImportPlaybackKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &ivs.ListChannelsInput{}
        output := &ivs.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlaybackKeyPairs", func(t *testing.T) {
        input := &ivs.ListPlaybackKeyPairsInput{}
        output := &ivs.ListPlaybackKeyPairsOutput{}

        mockClient.On("ListPlaybackKeyPairs", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlaybackKeyPairs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlaybackRestrictionPolicies", func(t *testing.T) {
        input := &ivs.ListPlaybackRestrictionPoliciesInput{}
        output := &ivs.ListPlaybackRestrictionPoliciesOutput{}

        mockClient.On("ListPlaybackRestrictionPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlaybackRestrictionPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecordingConfigurations", func(t *testing.T) {
        input := &ivs.ListRecordingConfigurationsInput{}
        output := &ivs.ListRecordingConfigurationsOutput{}

        mockClient.On("ListRecordingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecordingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamKeys", func(t *testing.T) {
        input := &ivs.ListStreamKeysInput{}
        output := &ivs.ListStreamKeysOutput{}

        mockClient.On("ListStreamKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamSessions", func(t *testing.T) {
        input := &ivs.ListStreamSessionsInput{}
        output := &ivs.ListStreamSessionsOutput{}

        mockClient.On("ListStreamSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreams", func(t *testing.T) {
        input := &ivs.ListStreamsInput{}
        output := &ivs.ListStreamsOutput{}

        mockClient.On("ListStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ivs.ListTagsForResourceInput{}
        output := &ivs.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetadata", func(t *testing.T) {
        input := &ivs.PutMetadataInput{}
        output := &ivs.PutMetadataOutput{}

        mockClient.On("PutMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartViewerSessionRevocation", func(t *testing.T) {
        input := &ivs.StartViewerSessionRevocationInput{}
        output := &ivs.StartViewerSessionRevocationOutput{}

        mockClient.On("StartViewerSessionRevocation", ctx, input).Return(output, nil)

        result, err := mockClient.StartViewerSessionRevocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStream", func(t *testing.T) {
        input := &ivs.StopStreamInput{}
        output := &ivs.StopStreamOutput{}

        mockClient.On("StopStream", ctx, input).Return(output, nil)

        result, err := mockClient.StopStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ivs.TagResourceInput{}
        output := &ivs.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ivs.UntagResourceInput{}
        output := &ivs.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &ivs.UpdateChannelInput{}
        output := &ivs.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePlaybackRestrictionPolicy", func(t *testing.T) {
        input := &ivs.UpdatePlaybackRestrictionPolicyInput{}
        output := &ivs.UpdatePlaybackRestrictionPolicyOutput{}

        mockClient.On("UpdatePlaybackRestrictionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePlaybackRestrictionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}