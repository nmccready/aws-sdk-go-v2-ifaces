package supplychain_test

// tests for the supplychain service interface for this ../../../service/supplychain/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/supplychain"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/supplychain/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/supplychain/supplychain_iface"
	"github.com/stretchr/testify/assert"
)

func TestSupplychainServiceCanBeMocked(t *testing.T) {
	var iface supplychain_iface.IClient
	iface = &supplychain.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := supplychain.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBillOfMaterialsImportJob", func(t *testing.T) {
        input := &supplychain.CreateBillOfMaterialsImportJobInput{}
        output := &supplychain.CreateBillOfMaterialsImportJobOutput{}

        mockClient.On("CreateBillOfMaterialsImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBillOfMaterialsImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBillOfMaterialsImportJob", func(t *testing.T) {
        input := &supplychain.GetBillOfMaterialsImportJobInput{}
        output := &supplychain.GetBillOfMaterialsImportJobOutput{}

        mockClient.On("GetBillOfMaterialsImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetBillOfMaterialsImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDataIntegrationEvent", func(t *testing.T) {
        input := &supplychain.SendDataIntegrationEventInput{}
        output := &supplychain.SendDataIntegrationEventOutput{}

        mockClient.On("SendDataIntegrationEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SendDataIntegrationEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
