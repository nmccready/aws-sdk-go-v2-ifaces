package marketplaceentitlementservice_test

// tests for the marketplaceentitlementservice service interface for this ../../../service/marketplaceentitlementservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplaceentitlementservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplaceentitlementservice/marketplaceentitlementservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplaceentitlementserviceServiceCanBeMocked(t *testing.T) {
	var iface marketplaceentitlementservice_iface.IClient
	iface = &marketplaceentitlementservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplaceentitlementservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEntitlements", func(t *testing.T) {
        input := &marketplaceentitlementservice.GetEntitlementsInput{}
        output := &marketplaceentitlementservice.GetEntitlementsOutput{}

        mockClient.On("GetEntitlements", ctx, input).Return(output, nil)

        result, err := mockClient.GetEntitlements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
