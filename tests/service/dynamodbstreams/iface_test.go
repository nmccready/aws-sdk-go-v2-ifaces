package dynamodbstreams_test

// tests for the dynamodbstreams service interface for this ../../../service/dynamodbstreams/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dynamodbstreams/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dynamodbstreams/dynamodbstreams_iface"
	"github.com/stretchr/testify/assert"
)

func TestDynamodbstreamsServiceCanBeMocked(t *testing.T) {
	var iface dynamodbstreams_iface.IClient
	iface = &dynamodbstreams.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := dynamodbstreams.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStream", func(t *testing.T) {
        input := &dynamodbstreams.DescribeStreamInput{}
        output := &dynamodbstreams.DescribeStreamOutput{}

        mockClient.On("DescribeStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecords", func(t *testing.T) {
        input := &dynamodbstreams.GetRecordsInput{}
        output := &dynamodbstreams.GetRecordsOutput{}

        mockClient.On("GetRecords", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetShardIterator", func(t *testing.T) {
        input := &dynamodbstreams.GetShardIteratorInput{}
        output := &dynamodbstreams.GetShardIteratorOutput{}

        mockClient.On("GetShardIterator", ctx, input).Return(output, nil)

        result, err := mockClient.GetShardIterator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreams", func(t *testing.T) {
        input := &dynamodbstreams.ListStreamsInput{}
        output := &dynamodbstreams.ListStreamsOutput{}

        mockClient.On("ListStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
