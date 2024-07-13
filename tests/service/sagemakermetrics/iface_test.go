package sagemakermetrics_test

// tests for the sagemakermetrics service interface for this ../../../service/sagemakermetrics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakermetrics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakermetrics/sagemakermetrics_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakermetricsServiceCanBeMocked(t *testing.T) {
	var iface sagemakermetrics_iface.IClient
	iface = &sagemakermetrics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakermetrics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutMetrics", func(t *testing.T) {
        input := &sagemakermetrics.BatchPutMetricsInput{}
        output := &sagemakermetrics.BatchPutMetricsOutput{}

        mockClient.On("BatchPutMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
