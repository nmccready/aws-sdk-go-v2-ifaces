package kinesisvideomedia_test

// tests for the kinesisvideomedia service interface for this ../../../service/kinesisvideomedia/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideomedia/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideomedia/kinesisvideomedia_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisvideomediaServiceCanBeMocked(t *testing.T) {
	var iface kinesisvideomedia_iface.IClient
	iface = &kinesisvideomedia.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisvideomedia.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMedia", func(t *testing.T) {
        input := &kinesisvideomedia.GetMediaInput{}
        output := &kinesisvideomedia.GetMediaOutput{}

        mockClient.On("GetMedia", ctx, input).Return(output, nil)

        result, err := mockClient.GetMedia(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
