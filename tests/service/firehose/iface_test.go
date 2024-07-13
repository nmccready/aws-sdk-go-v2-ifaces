package firehose_test

// tests for the firehose service interface for this ../../../service/firehose/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/firehose/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/firehose/firehose_iface"
	"github.com/stretchr/testify/assert"
)

func TestFirehoseServiceCanBeMocked(t *testing.T) {
	var iface firehose_iface.IClient
	iface = &firehose.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := firehose.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeliveryStream", func(t *testing.T) {
        input := &firehose.CreateDeliveryStreamInput{}
        output := &firehose.CreateDeliveryStreamOutput{}

        mockClient.On("CreateDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeliveryStream", func(t *testing.T) {
        input := &firehose.DeleteDeliveryStreamInput{}
        output := &firehose.DeleteDeliveryStreamOutput{}

        mockClient.On("DeleteDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliveryStream", func(t *testing.T) {
        input := &firehose.DescribeDeliveryStreamInput{}
        output := &firehose.DescribeDeliveryStreamOutput{}

        mockClient.On("DescribeDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeliveryStreams", func(t *testing.T) {
        input := &firehose.ListDeliveryStreamsInput{}
        output := &firehose.ListDeliveryStreamsOutput{}

        mockClient.On("ListDeliveryStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeliveryStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForDeliveryStream", func(t *testing.T) {
        input := &firehose.ListTagsForDeliveryStreamInput{}
        output := &firehose.ListTagsForDeliveryStreamOutput{}

        mockClient.On("ListTagsForDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecord", func(t *testing.T) {
        input := &firehose.PutRecordInput{}
        output := &firehose.PutRecordOutput{}

        mockClient.On("PutRecord", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecordBatch", func(t *testing.T) {
        input := &firehose.PutRecordBatchInput{}
        output := &firehose.PutRecordBatchOutput{}

        mockClient.On("PutRecordBatch", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecordBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeliveryStreamEncryption", func(t *testing.T) {
        input := &firehose.StartDeliveryStreamEncryptionInput{}
        output := &firehose.StartDeliveryStreamEncryptionOutput{}

        mockClient.On("StartDeliveryStreamEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeliveryStreamEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDeliveryStreamEncryption", func(t *testing.T) {
        input := &firehose.StopDeliveryStreamEncryptionInput{}
        output := &firehose.StopDeliveryStreamEncryptionOutput{}

        mockClient.On("StopDeliveryStreamEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.StopDeliveryStreamEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagDeliveryStream", func(t *testing.T) {
        input := &firehose.TagDeliveryStreamInput{}
        output := &firehose.TagDeliveryStreamOutput{}

        mockClient.On("TagDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.TagDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagDeliveryStream", func(t *testing.T) {
        input := &firehose.UntagDeliveryStreamInput{}
        output := &firehose.UntagDeliveryStreamOutput{}

        mockClient.On("UntagDeliveryStream", ctx, input).Return(output, nil)

        result, err := mockClient.UntagDeliveryStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDestination", func(t *testing.T) {
        input := &firehose.UpdateDestinationInput{}
        output := &firehose.UpdateDestinationOutput{}

        mockClient.On("UpdateDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
