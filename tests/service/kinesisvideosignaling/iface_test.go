// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package kinesisvideosignaling_test

// tests for the kinesisvideosignaling service interface for 
// this ../../../service/kinesisvideosignaling/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideosignaling/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisvideosignaling/kinesisvideosignaling_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisvideosignalingServiceCanBeMocked(t *testing.T) {
	var iface kinesisvideosignaling_iface.IClient
	iface = &kinesisvideosignaling.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisvideosignaling.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIceServerConfig", func(t *testing.T) {
        input := &kinesisvideosignaling.GetIceServerConfigInput{}
        output := &kinesisvideosignaling.GetIceServerConfigOutput{}

        mockClient.On("GetIceServerConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetIceServerConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendAlexaOfferToMaster", func(t *testing.T) {
        input := &kinesisvideosignaling.SendAlexaOfferToMasterInput{}
        output := &kinesisvideosignaling.SendAlexaOfferToMasterOutput{}

        mockClient.On("SendAlexaOfferToMaster", ctx, input).Return(output, nil)

        result, err := mockClient.SendAlexaOfferToMaster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
