package sagemakerfeaturestoreruntime_test

// tests for the sagemakerfeaturestoreruntime service interface for this ../../../service/sagemakerfeaturestoreruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakerfeaturestoreruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakerfeaturestoreruntime/sagemakerfeaturestoreruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakerfeaturestoreruntimeServiceCanBeMocked(t *testing.T) {
	var iface sagemakerfeaturestoreruntime_iface.IClient
	iface = &sagemakerfeaturestoreruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakerfeaturestoreruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetRecord", func(t *testing.T) {
        input := &sagemakerfeaturestoreruntime.BatchGetRecordInput{}
        output := &sagemakerfeaturestoreruntime.BatchGetRecordOutput{}

        mockClient.On("BatchGetRecord", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecord", func(t *testing.T) {
        input := &sagemakerfeaturestoreruntime.DeleteRecordInput{}
        output := &sagemakerfeaturestoreruntime.DeleteRecordOutput{}

        mockClient.On("DeleteRecord", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecord", func(t *testing.T) {
        input := &sagemakerfeaturestoreruntime.GetRecordInput{}
        output := &sagemakerfeaturestoreruntime.GetRecordOutput{}

        mockClient.On("GetRecord", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecord", func(t *testing.T) {
        input := &sagemakerfeaturestoreruntime.PutRecordInput{}
        output := &sagemakerfeaturestoreruntime.PutRecordOutput{}

        mockClient.On("PutRecord", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
