package marketplacemetering_test

// tests for the marketplacemetering service interface for this ../../../service/marketplacemetering/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacemetering/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacemetering/marketplacemetering_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplacemeteringServiceCanBeMocked(t *testing.T) {
	var iface marketplacemetering_iface.IClient
	iface = &marketplacemetering.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplacemetering.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchMeterUsage", func(t *testing.T) {
        input := &marketplacemetering.BatchMeterUsageInput{}
        output := &marketplacemetering.BatchMeterUsageOutput{}

        mockClient.On("BatchMeterUsage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchMeterUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMeterUsage", func(t *testing.T) {
        input := &marketplacemetering.MeterUsageInput{}
        output := &marketplacemetering.MeterUsageOutput{}

        mockClient.On("MeterUsage", ctx, input).Return(output, nil)

        result, err := mockClient.MeterUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterUsage", func(t *testing.T) {
        input := &marketplacemetering.RegisterUsageInput{}
        output := &marketplacemetering.RegisterUsageOutput{}

        mockClient.On("RegisterUsage", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResolveCustomer", func(t *testing.T) {
        input := &marketplacemetering.ResolveCustomerInput{}
        output := &marketplacemetering.ResolveCustomerOutput{}

        mockClient.On("ResolveCustomer", ctx, input).Return(output, nil)

        result, err := mockClient.ResolveCustomer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
