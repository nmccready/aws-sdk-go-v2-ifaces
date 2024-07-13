package mediapackage_test

// tests for the mediapackage service interface for this ../../../service/mediapackage/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackage/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackage/mediapackage_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediapackageServiceCanBeMocked(t *testing.T) {
	var iface mediapackage_iface.IClient
	iface = &mediapackage.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediapackage.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureLogs", func(t *testing.T) {
        input := &mediapackage.ConfigureLogsInput{}
        output := &mediapackage.ConfigureLogsOutput{}

        mockClient.On("ConfigureLogs", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &mediapackage.CreateChannelInput{}
        output := &mediapackage.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHarvestJob", func(t *testing.T) {
        input := &mediapackage.CreateHarvestJobInput{}
        output := &mediapackage.CreateHarvestJobOutput{}

        mockClient.On("CreateHarvestJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHarvestJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOriginEndpoint", func(t *testing.T) {
        input := &mediapackage.CreateOriginEndpointInput{}
        output := &mediapackage.CreateOriginEndpointOutput{}

        mockClient.On("CreateOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &mediapackage.DeleteChannelInput{}
        output := &mediapackage.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOriginEndpoint", func(t *testing.T) {
        input := &mediapackage.DeleteOriginEndpointInput{}
        output := &mediapackage.DeleteOriginEndpointOutput{}

        mockClient.On("DeleteOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &mediapackage.DescribeChannelInput{}
        output := &mediapackage.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHarvestJob", func(t *testing.T) {
        input := &mediapackage.DescribeHarvestJobInput{}
        output := &mediapackage.DescribeHarvestJobOutput{}

        mockClient.On("DescribeHarvestJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHarvestJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOriginEndpoint", func(t *testing.T) {
        input := &mediapackage.DescribeOriginEndpointInput{}
        output := &mediapackage.DescribeOriginEndpointOutput{}

        mockClient.On("DescribeOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &mediapackage.ListChannelsInput{}
        output := &mediapackage.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHarvestJobs", func(t *testing.T) {
        input := &mediapackage.ListHarvestJobsInput{}
        output := &mediapackage.ListHarvestJobsOutput{}

        mockClient.On("ListHarvestJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListHarvestJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOriginEndpoints", func(t *testing.T) {
        input := &mediapackage.ListOriginEndpointsInput{}
        output := &mediapackage.ListOriginEndpointsOutput{}

        mockClient.On("ListOriginEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListOriginEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediapackage.ListTagsForResourceInput{}
        output := &mediapackage.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateChannelCredentials", func(t *testing.T) {
        input := &mediapackage.RotateChannelCredentialsInput{}
        output := &mediapackage.RotateChannelCredentialsOutput{}

        mockClient.On("RotateChannelCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.RotateChannelCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateIngestEndpointCredentials", func(t *testing.T) {
        input := &mediapackage.RotateIngestEndpointCredentialsInput{}
        output := &mediapackage.RotateIngestEndpointCredentialsOutput{}

        mockClient.On("RotateIngestEndpointCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.RotateIngestEndpointCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediapackage.TagResourceInput{}
        output := &mediapackage.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediapackage.UntagResourceInput{}
        output := &mediapackage.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &mediapackage.UpdateChannelInput{}
        output := &mediapackage.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOriginEndpoint", func(t *testing.T) {
        input := &mediapackage.UpdateOriginEndpointInput{}
        output := &mediapackage.UpdateOriginEndpointOutput{}

        mockClient.On("UpdateOriginEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOriginEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
