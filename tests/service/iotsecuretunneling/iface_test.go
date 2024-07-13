package iotsecuretunneling_test

// tests for the iotsecuretunneling service interface for this ../../../service/iotsecuretunneling/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotsecuretunneling/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotsecuretunneling/iotsecuretunneling_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotsecuretunnelingServiceCanBeMocked(t *testing.T) {
	var iface iotsecuretunneling_iface.IClient
	iface = &iotsecuretunneling.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotsecuretunneling.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloseTunnel", func(t *testing.T) {
        input := &iotsecuretunneling.CloseTunnelInput{}
        output := &iotsecuretunneling.CloseTunnelOutput{}

        mockClient.On("CloseTunnel", ctx, input).Return(output, nil)

        result, err := mockClient.CloseTunnel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTunnel", func(t *testing.T) {
        input := &iotsecuretunneling.DescribeTunnelInput{}
        output := &iotsecuretunneling.DescribeTunnelOutput{}

        mockClient.On("DescribeTunnel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTunnel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotsecuretunneling.ListTagsForResourceInput{}
        output := &iotsecuretunneling.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTunnels", func(t *testing.T) {
        input := &iotsecuretunneling.ListTunnelsInput{}
        output := &iotsecuretunneling.ListTunnelsOutput{}

        mockClient.On("ListTunnels", ctx, input).Return(output, nil)

        result, err := mockClient.ListTunnels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestOpenTunnel", func(t *testing.T) {
        input := &iotsecuretunneling.OpenTunnelInput{}
        output := &iotsecuretunneling.OpenTunnelOutput{}

        mockClient.On("OpenTunnel", ctx, input).Return(output, nil)

        result, err := mockClient.OpenTunnel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateTunnelAccessToken", func(t *testing.T) {
        input := &iotsecuretunneling.RotateTunnelAccessTokenInput{}
        output := &iotsecuretunneling.RotateTunnelAccessTokenOutput{}

        mockClient.On("RotateTunnelAccessToken", ctx, input).Return(output, nil)

        result, err := mockClient.RotateTunnelAccessToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotsecuretunneling.TagResourceInput{}
        output := &iotsecuretunneling.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotsecuretunneling.UntagResourceInput{}
        output := &iotsecuretunneling.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
