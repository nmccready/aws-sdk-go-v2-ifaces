package pricing_test

// tests for the pricing service interface for this ../../../service/pricing/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pricing/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pricing/pricing_iface"
	"github.com/stretchr/testify/assert"
)

func TestPricingServiceCanBeMocked(t *testing.T) {
	var iface pricing_iface.IClient
	iface = &pricing.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pricing.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServices", func(t *testing.T) {
        input := &pricing.DescribeServicesInput{}
        output := &pricing.DescribeServicesOutput{}

        mockClient.On("DescribeServices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttributeValues", func(t *testing.T) {
        input := &pricing.GetAttributeValuesInput{}
        output := &pricing.GetAttributeValuesOutput{}

        mockClient.On("GetAttributeValues", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttributeValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPriceListFileUrl", func(t *testing.T) {
        input := &pricing.GetPriceListFileUrlInput{}
        output := &pricing.GetPriceListFileUrlOutput{}

        mockClient.On("GetPriceListFileUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetPriceListFileUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProducts", func(t *testing.T) {
        input := &pricing.GetProductsInput{}
        output := &pricing.GetProductsOutput{}

        mockClient.On("GetProducts", ctx, input).Return(output, nil)

        result, err := mockClient.GetProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPriceLists", func(t *testing.T) {
        input := &pricing.ListPriceListsInput{}
        output := &pricing.ListPriceListsOutput{}

        mockClient.On("ListPriceLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListPriceLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
