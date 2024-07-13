package mediapackagev2_test

// tests for the mediapackagev2 service interface for this ../../../service/mediapackagev2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackagev2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackagev2/mediapackagev2_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediapackagev2ServiceCanBeMocked(t *testing.T) {
	var iface mediapackagev2_iface.IClient
	iface = &mediapackagev2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediapackagev2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &mediapackagev2.CreateChannelInput{}
        output := &mediapackagev2.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannelGroup", func(t *testing.T) {
        input := &mediapackagev2.CreateChannelGroupInput{}
        output := &mediapackagev2.CreateChannelGroupOutput{}

        mockClient.On("CreateChannelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOriginEndpoint", func(t *testing.T) {
        input := &mediapackagev2.CreateOriginEndpointInput{}
        output := &mediapackagev2.CreateOriginEndpointOutput{}

        mockClient.On("CreateOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &mediapackagev2.DeleteChannelInput{}
        output := &mediapackagev2.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelGroup", func(t *testing.T) {
        input := &mediapackagev2.DeleteChannelGroupInput{}
        output := &mediapackagev2.DeleteChannelGroupOutput{}

        mockClient.On("DeleteChannelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelPolicy", func(t *testing.T) {
        input := &mediapackagev2.DeleteChannelPolicyInput{}
        output := &mediapackagev2.DeleteChannelPolicyOutput{}

        mockClient.On("DeleteChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOriginEndpoint", func(t *testing.T) {
        input := &mediapackagev2.DeleteOriginEndpointInput{}
        output := &mediapackagev2.DeleteOriginEndpointOutput{}

        mockClient.On("DeleteOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOriginEndpointPolicy", func(t *testing.T) {
        input := &mediapackagev2.DeleteOriginEndpointPolicyInput{}
        output := &mediapackagev2.DeleteOriginEndpointPolicyOutput{}

        mockClient.On("DeleteOriginEndpointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOriginEndpointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannel", func(t *testing.T) {
        input := &mediapackagev2.GetChannelInput{}
        output := &mediapackagev2.GetChannelOutput{}

        mockClient.On("GetChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelGroup", func(t *testing.T) {
        input := &mediapackagev2.GetChannelGroupInput{}
        output := &mediapackagev2.GetChannelGroupOutput{}

        mockClient.On("GetChannelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelPolicy", func(t *testing.T) {
        input := &mediapackagev2.GetChannelPolicyInput{}
        output := &mediapackagev2.GetChannelPolicyOutput{}

        mockClient.On("GetChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginEndpoint", func(t *testing.T) {
        input := &mediapackagev2.GetOriginEndpointInput{}
        output := &mediapackagev2.GetOriginEndpointOutput{}

        mockClient.On("GetOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginEndpointPolicy", func(t *testing.T) {
        input := &mediapackagev2.GetOriginEndpointPolicyInput{}
        output := &mediapackagev2.GetOriginEndpointPolicyOutput{}

        mockClient.On("GetOriginEndpointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginEndpointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannelGroups", func(t *testing.T) {
        input := &mediapackagev2.ListChannelGroupsInput{}
        output := &mediapackagev2.ListChannelGroupsOutput{}

        mockClient.On("ListChannelGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannelGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &mediapackagev2.ListChannelsInput{}
        output := &mediapackagev2.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOriginEndpoints", func(t *testing.T) {
        input := &mediapackagev2.ListOriginEndpointsInput{}
        output := &mediapackagev2.ListOriginEndpointsOutput{}

        mockClient.On("ListOriginEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListOriginEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediapackagev2.ListTagsForResourceInput{}
        output := &mediapackagev2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutChannelPolicy", func(t *testing.T) {
        input := &mediapackagev2.PutChannelPolicyInput{}
        output := &mediapackagev2.PutChannelPolicyOutput{}

        mockClient.On("PutChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOriginEndpointPolicy", func(t *testing.T) {
        input := &mediapackagev2.PutOriginEndpointPolicyInput{}
        output := &mediapackagev2.PutOriginEndpointPolicyOutput{}

        mockClient.On("PutOriginEndpointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutOriginEndpointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediapackagev2.TagResourceInput{}
        output := &mediapackagev2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediapackagev2.UntagResourceInput{}
        output := &mediapackagev2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &mediapackagev2.UpdateChannelInput{}
        output := &mediapackagev2.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelGroup", func(t *testing.T) {
        input := &mediapackagev2.UpdateChannelGroupInput{}
        output := &mediapackagev2.UpdateChannelGroupOutput{}

        mockClient.On("UpdateChannelGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOriginEndpoint", func(t *testing.T) {
        input := &mediapackagev2.UpdateOriginEndpointInput{}
        output := &mediapackagev2.UpdateOriginEndpointOutput{}

        mockClient.On("UpdateOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
