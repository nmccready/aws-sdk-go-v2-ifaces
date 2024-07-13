package ec2instanceconnect_test

// tests for the ec2instanceconnect service interface for this ../../../service/ec2instanceconnect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ec2instanceconnect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ec2instanceconnect/ec2instanceconnect_iface"
	"github.com/stretchr/testify/assert"
)

func TestEc2instanceconnectServiceCanBeMocked(t *testing.T) {
	var iface ec2instanceconnect_iface.IClient
	iface = &ec2instanceconnect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ec2instanceconnect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendSSHPublicKey", func(t *testing.T) {
        input := &ec2instanceconnect.SendSSHPublicKeyInput{}
        output := &ec2instanceconnect.SendSSHPublicKeyOutput{}

        mockClient.On("SendSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.SendSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendSerialConsoleSSHPublicKey", func(t *testing.T) {
        input := &ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput{}
        output := &ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput{}

        mockClient.On("SendSerialConsoleSSHPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.SendSerialConsoleSSHPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
