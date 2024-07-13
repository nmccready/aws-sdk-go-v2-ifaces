package freetier_test

// tests for the freetier service interface for this ../../../service/freetier/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/freetier"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/freetier/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/freetier/freetier_iface"
	"github.com/stretchr/testify/assert"
)

func TestFreetierServiceCanBeMocked(t *testing.T) {
	var iface freetier_iface.IClient
	iface = &freetier.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := freetier.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFreeTierUsage", func(t *testing.T) {
        input := &freetier.GetFreeTierUsageInput{}
        output := &freetier.GetFreeTierUsageOutput{}

        mockClient.On("GetFreeTierUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetFreeTierUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
