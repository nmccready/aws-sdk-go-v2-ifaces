package marketplacecommerceanalytics_test

// tests for the marketplacecommerceanalytics service interface for this ../../../service/marketplacecommerceanalytics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacecommerceanalytics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacecommerceanalytics/marketplacecommerceanalytics_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplacecommerceanalyticsServiceCanBeMocked(t *testing.T) {
	var iface marketplacecommerceanalytics_iface.IClient
	iface = &marketplacecommerceanalytics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplacecommerceanalytics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateDataSet", func(t *testing.T) {
        input := &marketplacecommerceanalytics.GenerateDataSetInput{}
        output := &marketplacecommerceanalytics.GenerateDataSetOutput{}

        mockClient.On("GenerateDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSupportDataExport", func(t *testing.T) {
        input := &marketplacecommerceanalytics.StartSupportDataExportInput{}
        output := &marketplacecommerceanalytics.StartSupportDataExportOutput{}

        mockClient.On("StartSupportDataExport", ctx, input).Return(output, nil)

        result, err := mockClient.StartSupportDataExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
