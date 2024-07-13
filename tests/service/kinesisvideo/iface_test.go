package kinesisvideo_test

// tests for the kinesisvideo service interface for this ../../../service/kinesisvideo/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideo/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideo/kinesisvideo_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisvideoServiceCanBeMocked(t *testing.T) {
	var iface kinesisvideo_iface.IClient
	iface = &kinesisvideo.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisvideo.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSignalingChannel", func(t *testing.T) {
        input := &kinesisvideo.CreateSignalingChannelInput{}
        output := &kinesisvideo.CreateSignalingChannelOutput{}

        mockClient.On("CreateSignalingChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSignalingChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStream", func(t *testing.T) {
        input := &kinesisvideo.CreateStreamInput{}
        output := &kinesisvideo.CreateStreamOutput{}

        mockClient.On("CreateStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEdgeConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DeleteEdgeConfigurationInput{}
        output := &kinesisvideo.DeleteEdgeConfigurationOutput{}

        mockClient.On("DeleteEdgeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEdgeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSignalingChannel", func(t *testing.T) {
        input := &kinesisvideo.DeleteSignalingChannelInput{}
        output := &kinesisvideo.DeleteSignalingChannelOutput{}

        mockClient.On("DeleteSignalingChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSignalingChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStream", func(t *testing.T) {
        input := &kinesisvideo.DeleteStreamInput{}
        output := &kinesisvideo.DeleteStreamOutput{}

        mockClient.On("DeleteStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEdgeConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DescribeEdgeConfigurationInput{}
        output := &kinesisvideo.DescribeEdgeConfigurationOutput{}

        mockClient.On("DescribeEdgeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEdgeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageGenerationConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DescribeImageGenerationConfigurationInput{}
        output := &kinesisvideo.DescribeImageGenerationConfigurationOutput{}

        mockClient.On("DescribeImageGenerationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageGenerationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMappedResourceConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DescribeMappedResourceConfigurationInput{}
        output := &kinesisvideo.DescribeMappedResourceConfigurationOutput{}

        mockClient.On("DescribeMappedResourceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMappedResourceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMediaStorageConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DescribeMediaStorageConfigurationInput{}
        output := &kinesisvideo.DescribeMediaStorageConfigurationOutput{}

        mockClient.On("DescribeMediaStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMediaStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotificationConfiguration", func(t *testing.T) {
        input := &kinesisvideo.DescribeNotificationConfigurationInput{}
        output := &kinesisvideo.DescribeNotificationConfigurationOutput{}

        mockClient.On("DescribeNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSignalingChannel", func(t *testing.T) {
        input := &kinesisvideo.DescribeSignalingChannelInput{}
        output := &kinesisvideo.DescribeSignalingChannelOutput{}

        mockClient.On("DescribeSignalingChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSignalingChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStream", func(t *testing.T) {
        input := &kinesisvideo.DescribeStreamInput{}
        output := &kinesisvideo.DescribeStreamOutput{}

        mockClient.On("DescribeStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataEndpoint", func(t *testing.T) {
        input := &kinesisvideo.GetDataEndpointInput{}
        output := &kinesisvideo.GetDataEndpointOutput{}

        mockClient.On("GetDataEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSignalingChannelEndpoint", func(t *testing.T) {
        input := &kinesisvideo.GetSignalingChannelEndpointInput{}
        output := &kinesisvideo.GetSignalingChannelEndpointOutput{}

        mockClient.On("GetSignalingChannelEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetSignalingChannelEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEdgeAgentConfigurations", func(t *testing.T) {
        input := &kinesisvideo.ListEdgeAgentConfigurationsInput{}
        output := &kinesisvideo.ListEdgeAgentConfigurationsOutput{}

        mockClient.On("ListEdgeAgentConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEdgeAgentConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSignalingChannels", func(t *testing.T) {
        input := &kinesisvideo.ListSignalingChannelsInput{}
        output := &kinesisvideo.ListSignalingChannelsOutput{}

        mockClient.On("ListSignalingChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListSignalingChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreams", func(t *testing.T) {
        input := &kinesisvideo.ListStreamsInput{}
        output := &kinesisvideo.ListStreamsOutput{}

        mockClient.On("ListStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kinesisvideo.ListTagsForResourceInput{}
        output := &kinesisvideo.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForStream", func(t *testing.T) {
        input := &kinesisvideo.ListTagsForStreamInput{}
        output := &kinesisvideo.ListTagsForStreamOutput{}

        mockClient.On("ListTagsForStream", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEdgeConfigurationUpdate", func(t *testing.T) {
        input := &kinesisvideo.StartEdgeConfigurationUpdateInput{}
        output := &kinesisvideo.StartEdgeConfigurationUpdateOutput{}

        mockClient.On("StartEdgeConfigurationUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartEdgeConfigurationUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kinesisvideo.TagResourceInput{}
        output := &kinesisvideo.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagStream", func(t *testing.T) {
        input := &kinesisvideo.TagStreamInput{}
        output := &kinesisvideo.TagStreamOutput{}

        mockClient.On("TagStream", ctx, input).Return(output, nil)

        result, err := mockClient.TagStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kinesisvideo.UntagResourceInput{}
        output := &kinesisvideo.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagStream", func(t *testing.T) {
        input := &kinesisvideo.UntagStreamInput{}
        output := &kinesisvideo.UntagStreamOutput{}

        mockClient.On("UntagStream", ctx, input).Return(output, nil)

        result, err := mockClient.UntagStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataRetention", func(t *testing.T) {
        input := &kinesisvideo.UpdateDataRetentionInput{}
        output := &kinesisvideo.UpdateDataRetentionOutput{}

        mockClient.On("UpdateDataRetention", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataRetention(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImageGenerationConfiguration", func(t *testing.T) {
        input := &kinesisvideo.UpdateImageGenerationConfigurationInput{}
        output := &kinesisvideo.UpdateImageGenerationConfigurationOutput{}

        mockClient.On("UpdateImageGenerationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImageGenerationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMediaStorageConfiguration", func(t *testing.T) {
        input := &kinesisvideo.UpdateMediaStorageConfigurationInput{}
        output := &kinesisvideo.UpdateMediaStorageConfigurationOutput{}

        mockClient.On("UpdateMediaStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMediaStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotificationConfiguration", func(t *testing.T) {
        input := &kinesisvideo.UpdateNotificationConfigurationInput{}
        output := &kinesisvideo.UpdateNotificationConfigurationOutput{}

        mockClient.On("UpdateNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSignalingChannel", func(t *testing.T) {
        input := &kinesisvideo.UpdateSignalingChannelInput{}
        output := &kinesisvideo.UpdateSignalingChannelOutput{}

        mockClient.On("UpdateSignalingChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSignalingChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStream", func(t *testing.T) {
        input := &kinesisvideo.UpdateStreamInput{}
        output := &kinesisvideo.UpdateStreamOutput{}

        mockClient.On("UpdateStream", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
