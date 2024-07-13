package ioteventsdata_test

// tests for the ioteventsdata service interface for this ../../../service/ioteventsdata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ioteventsdata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ioteventsdata/ioteventsdata_iface"
	"github.com/stretchr/testify/assert"
)

func TestIoteventsdataServiceCanBeMocked(t *testing.T) {
	var iface ioteventsdata_iface.IClient
	iface = &ioteventsdata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ioteventsdata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAcknowledgeAlarm", func(t *testing.T) {
        input := &ioteventsdata.BatchAcknowledgeAlarmInput{}
        output := &ioteventsdata.BatchAcknowledgeAlarmOutput{}

        mockClient.On("BatchAcknowledgeAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAcknowledgeAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteDetector", func(t *testing.T) {
        input := &ioteventsdata.BatchDeleteDetectorInput{}
        output := &ioteventsdata.BatchDeleteDetectorOutput{}

        mockClient.On("BatchDeleteDetector", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisableAlarm", func(t *testing.T) {
        input := &ioteventsdata.BatchDisableAlarmInput{}
        output := &ioteventsdata.BatchDisableAlarmOutput{}

        mockClient.On("BatchDisableAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisableAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchEnableAlarm", func(t *testing.T) {
        input := &ioteventsdata.BatchEnableAlarmInput{}
        output := &ioteventsdata.BatchEnableAlarmOutput{}

        mockClient.On("BatchEnableAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.BatchEnableAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutMessage", func(t *testing.T) {
        input := &ioteventsdata.BatchPutMessageInput{}
        output := &ioteventsdata.BatchPutMessageOutput{}

        mockClient.On("BatchPutMessage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchResetAlarm", func(t *testing.T) {
        input := &ioteventsdata.BatchResetAlarmInput{}
        output := &ioteventsdata.BatchResetAlarmOutput{}

        mockClient.On("BatchResetAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.BatchResetAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchSnoozeAlarm", func(t *testing.T) {
        input := &ioteventsdata.BatchSnoozeAlarmInput{}
        output := &ioteventsdata.BatchSnoozeAlarmOutput{}

        mockClient.On("BatchSnoozeAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.BatchSnoozeAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateDetector", func(t *testing.T) {
        input := &ioteventsdata.BatchUpdateDetectorInput{}
        output := &ioteventsdata.BatchUpdateDetectorOutput{}

        mockClient.On("BatchUpdateDetector", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlarm", func(t *testing.T) {
        input := &ioteventsdata.DescribeAlarmInput{}
        output := &ioteventsdata.DescribeAlarmOutput{}

        mockClient.On("DescribeAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDetector", func(t *testing.T) {
        input := &ioteventsdata.DescribeDetectorInput{}
        output := &ioteventsdata.DescribeDetectorOutput{}

        mockClient.On("DescribeDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlarms", func(t *testing.T) {
        input := &ioteventsdata.ListAlarmsInput{}
        output := &ioteventsdata.ListAlarmsOutput{}

        mockClient.On("ListAlarms", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlarms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectors", func(t *testing.T) {
        input := &ioteventsdata.ListDetectorsInput{}
        output := &ioteventsdata.ListDetectorsOutput{}

        mockClient.On("ListDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
