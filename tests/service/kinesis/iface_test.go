package kinesis_test

// tests for the kinesis service interface for this ../../../service/kinesis/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesis/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesis/kinesis_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisServiceCanBeMocked(t *testing.T) {
	var iface kinesis_iface.IClient
	iface = &kinesis.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesis.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToStream", func(t *testing.T) {
        input := &kinesis.AddTagsToStreamInput{}
        output := &kinesis.AddTagsToStreamOutput{}

        mockClient.On("AddTagsToStream", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStream", func(t *testing.T) {
        input := &kinesis.CreateStreamInput{}
        output := &kinesis.CreateStreamOutput{}

        mockClient.On("CreateStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecreaseStreamRetentionPeriod", func(t *testing.T) {
        input := &kinesis.DecreaseStreamRetentionPeriodInput{}
        output := &kinesis.DecreaseStreamRetentionPeriodOutput{}

        mockClient.On("DecreaseStreamRetentionPeriod", ctx, input).Return(output, nil)

        result, err := mockClient.DecreaseStreamRetentionPeriod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &kinesis.DeleteResourcePolicyInput{}
        output := &kinesis.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStream", func(t *testing.T) {
        input := &kinesis.DeleteStreamInput{}
        output := &kinesis.DeleteStreamOutput{}

        mockClient.On("DeleteStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterStreamConsumer", func(t *testing.T) {
        input := &kinesis.DeregisterStreamConsumerInput{}
        output := &kinesis.DeregisterStreamConsumerOutput{}

        mockClient.On("DeregisterStreamConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterStreamConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLimits", func(t *testing.T) {
        input := &kinesis.DescribeLimitsInput{}
        output := &kinesis.DescribeLimitsOutput{}

        mockClient.On("DescribeLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStream", func(t *testing.T) {
        input := &kinesis.DescribeStreamInput{}
        output := &kinesis.DescribeStreamOutput{}

        mockClient.On("DescribeStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStreamConsumer", func(t *testing.T) {
        input := &kinesis.DescribeStreamConsumerInput{}
        output := &kinesis.DescribeStreamConsumerOutput{}

        mockClient.On("DescribeStreamConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStreamConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStreamSummary", func(t *testing.T) {
        input := &kinesis.DescribeStreamSummaryInput{}
        output := &kinesis.DescribeStreamSummaryOutput{}

        mockClient.On("DescribeStreamSummary", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStreamSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableEnhancedMonitoring", func(t *testing.T) {
        input := &kinesis.DisableEnhancedMonitoringInput{}
        output := &kinesis.DisableEnhancedMonitoringOutput{}

        mockClient.On("DisableEnhancedMonitoring", ctx, input).Return(output, nil)

        result, err := mockClient.DisableEnhancedMonitoring(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableEnhancedMonitoring", func(t *testing.T) {
        input := &kinesis.EnableEnhancedMonitoringInput{}
        output := &kinesis.EnableEnhancedMonitoringOutput{}

        mockClient.On("EnableEnhancedMonitoring", ctx, input).Return(output, nil)

        result, err := mockClient.EnableEnhancedMonitoring(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecords", func(t *testing.T) {
        input := &kinesis.GetRecordsInput{}
        output := &kinesis.GetRecordsOutput{}

        mockClient.On("GetRecords", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &kinesis.GetResourcePolicyInput{}
        output := &kinesis.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetShardIterator", func(t *testing.T) {
        input := &kinesis.GetShardIteratorInput{}
        output := &kinesis.GetShardIteratorOutput{}

        mockClient.On("GetShardIterator", ctx, input).Return(output, nil)

        result, err := mockClient.GetShardIterator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIncreaseStreamRetentionPeriod", func(t *testing.T) {
        input := &kinesis.IncreaseStreamRetentionPeriodInput{}
        output := &kinesis.IncreaseStreamRetentionPeriodOutput{}

        mockClient.On("IncreaseStreamRetentionPeriod", ctx, input).Return(output, nil)

        result, err := mockClient.IncreaseStreamRetentionPeriod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListShards", func(t *testing.T) {
        input := &kinesis.ListShardsInput{}
        output := &kinesis.ListShardsOutput{}

        mockClient.On("ListShards", ctx, input).Return(output, nil)

        result, err := mockClient.ListShards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamConsumers", func(t *testing.T) {
        input := &kinesis.ListStreamConsumersInput{}
        output := &kinesis.ListStreamConsumersOutput{}

        mockClient.On("ListStreamConsumers", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamConsumers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreams", func(t *testing.T) {
        input := &kinesis.ListStreamsInput{}
        output := &kinesis.ListStreamsOutput{}

        mockClient.On("ListStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForStream", func(t *testing.T) {
        input := &kinesis.ListTagsForStreamInput{}
        output := &kinesis.ListTagsForStreamOutput{}

        mockClient.On("ListTagsForStream", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeShards", func(t *testing.T) {
        input := &kinesis.MergeShardsInput{}
        output := &kinesis.MergeShardsOutput{}

        mockClient.On("MergeShards", ctx, input).Return(output, nil)

        result, err := mockClient.MergeShards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecord", func(t *testing.T) {
        input := &kinesis.PutRecordInput{}
        output := &kinesis.PutRecordOutput{}

        mockClient.On("PutRecord", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecords", func(t *testing.T) {
        input := &kinesis.PutRecordsInput{}
        output := &kinesis.PutRecordsOutput{}

        mockClient.On("PutRecords", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &kinesis.PutResourcePolicyInput{}
        output := &kinesis.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterStreamConsumer", func(t *testing.T) {
        input := &kinesis.RegisterStreamConsumerInput{}
        output := &kinesis.RegisterStreamConsumerOutput{}

        mockClient.On("RegisterStreamConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterStreamConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromStream", func(t *testing.T) {
        input := &kinesis.RemoveTagsFromStreamInput{}
        output := &kinesis.RemoveTagsFromStreamOutput{}

        mockClient.On("RemoveTagsFromStream", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSplitShard", func(t *testing.T) {
        input := &kinesis.SplitShardInput{}
        output := &kinesis.SplitShardOutput{}

        mockClient.On("SplitShard", ctx, input).Return(output, nil)

        result, err := mockClient.SplitShard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStreamEncryption", func(t *testing.T) {
        input := &kinesis.StartStreamEncryptionInput{}
        output := &kinesis.StartStreamEncryptionOutput{}

        mockClient.On("StartStreamEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.StartStreamEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStreamEncryption", func(t *testing.T) {
        input := &kinesis.StopStreamEncryptionInput{}
        output := &kinesis.StopStreamEncryptionOutput{}

        mockClient.On("StopStreamEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.StopStreamEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubscribeToShard", func(t *testing.T) {
        input := &kinesis.SubscribeToShardInput{}
        output := &kinesis.SubscribeToShardOutput{}

        mockClient.On("SubscribeToShard", ctx, input).Return(output, nil)

        result, err := mockClient.SubscribeToShard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateShardCount", func(t *testing.T) {
        input := &kinesis.UpdateShardCountInput{}
        output := &kinesis.UpdateShardCountOutput{}

        mockClient.On("UpdateShardCount", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateShardCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStreamMode", func(t *testing.T) {
        input := &kinesis.UpdateStreamModeInput{}
        output := &kinesis.UpdateStreamModeOutput{}

        mockClient.On("UpdateStreamMode", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStreamMode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
