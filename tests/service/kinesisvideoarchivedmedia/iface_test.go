package kinesisvideoarchivedmedia_test

// tests for the kinesisvideoarchivedmedia service interface for this ../../../service/kinesisvideoarchivedmedia/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideoarchivedmedia/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmedia_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisvideoarchivedmediaServiceCanBeMocked(t *testing.T) {
	var iface kinesisvideoarchivedmedia_iface.IClient
	iface = &kinesisvideoarchivedmedia.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisvideoarchivedmedia.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClip", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.GetClipInput{}
        output := &kinesisvideoarchivedmedia.GetClipOutput{}

        mockClient.On("GetClip", ctx, input).Return(output, nil)

        result, err := mockClient.GetClip(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDASHStreamingSessionURL", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput{}
        output := &kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput{}

        mockClient.On("GetDASHStreamingSessionURL", ctx, input).Return(output, nil)

        result, err := mockClient.GetDASHStreamingSessionURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHLSStreamingSessionURL", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput{}
        output := &kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput{}

        mockClient.On("GetHLSStreamingSessionURL", ctx, input).Return(output, nil)

        result, err := mockClient.GetHLSStreamingSessionURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImages", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.GetImagesInput{}
        output := &kinesisvideoarchivedmedia.GetImagesOutput{}

        mockClient.On("GetImages", ctx, input).Return(output, nil)

        result, err := mockClient.GetImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaForFragmentList", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.GetMediaForFragmentListInput{}
        output := &kinesisvideoarchivedmedia.GetMediaForFragmentListOutput{}

        mockClient.On("GetMediaForFragmentList", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaForFragmentList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFragments", func(t *testing.T) {
        input := &kinesisvideoarchivedmedia.ListFragmentsInput{}
        output := &kinesisvideoarchivedmedia.ListFragmentsOutput{}

        mockClient.On("ListFragments", ctx, input).Return(output, nil)

        result, err := mockClient.ListFragments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
