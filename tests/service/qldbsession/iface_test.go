package qldbsession_test

// tests for the qldbsession service interface for this ../../../service/qldbsession/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qldbsession"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qldbsession/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qldbsession/qldbsession_iface"
	"github.com/stretchr/testify/assert"
)

func TestQldbsessionServiceCanBeMocked(t *testing.T) {
	var iface qldbsession_iface.IClient
	iface = &qldbsession.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := qldbsession.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCommand", func(t *testing.T) {
        input := &qldbsession.SendCommandInput{}
        output := &qldbsession.SendCommandOutput{}

        mockClient.On("SendCommand", ctx, input).Return(output, nil)

        result, err := mockClient.SendCommand(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
