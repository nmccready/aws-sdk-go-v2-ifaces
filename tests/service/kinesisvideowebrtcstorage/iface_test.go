package kinesisvideowebrtcstorage_test

// tests for the kinesisvideowebrtcstorage service interface for this ../../../service/kinesisvideowebrtcstorage/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideowebrtcstorage/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideowebrtcstorage/kinesisvideowebrtcstorage_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisvideowebrtcstorageServiceCanBeMocked(t *testing.T) {
	var iface kinesisvideowebrtcstorage_iface.IClient
	iface = &kinesisvideowebrtcstorage.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisvideowebrtcstorage.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestJoinStorageSession", func(t *testing.T) {
        input := &kinesisvideowebrtcstorage.JoinStorageSessionInput{}
        output := &kinesisvideowebrtcstorage.JoinStorageSessionOutput{}

        mockClient.On("JoinStorageSession", ctx, input).Return(output, nil)

        result, err := mockClient.JoinStorageSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
