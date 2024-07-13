package inspectorscan_test

// tests for the inspectorscan service interface for this ../../../service/inspectorscan/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspectorscan"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspectorscan/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspectorscan/inspectorscan_iface"
	"github.com/stretchr/testify/assert"
)

func TestInspectorscanServiceCanBeMocked(t *testing.T) {
	var iface inspectorscan_iface.IClient
	iface = &inspectorscan.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := inspectorscan.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestScanSbom", func(t *testing.T) {
        input := &inspectorscan.ScanSbomInput{}
        output := &inspectorscan.ScanSbomOutput{}

        mockClient.On("ScanSbom", ctx, input).Return(output, nil)

        result, err := mockClient.ScanSbom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
