// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package sqs_test

// tests for the sqs service interface for 
// this ../../../service/sqs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sqs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sqs/sqs_iface"
	"github.com/stretchr/testify/assert"
)

func TestSqsServiceCanBeMocked(t *testing.T) {
	var iface sqs_iface.IClient
	iface = &sqs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sqs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddPermission", func(t *testing.T) {
        input := &sqs.AddPermissionInput{}
        output := &sqs.AddPermissionOutput{}

        mockClient.On("AddPermission", ctx, input).Return(output, nil)

        result, err := mockClient.AddPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMessageMoveTask", func(t *testing.T) {
        input := &sqs.CancelMessageMoveTaskInput{}
        output := &sqs.CancelMessageMoveTaskOutput{}

        mockClient.On("CancelMessageMoveTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMessageMoveTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeMessageVisibility", func(t *testing.T) {
        input := &sqs.ChangeMessageVisibilityInput{}
        output := &sqs.ChangeMessageVisibilityOutput{}

        mockClient.On("ChangeMessageVisibility", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeMessageVisibility(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeMessageVisibilityBatch", func(t *testing.T) {
        input := &sqs.ChangeMessageVisibilityBatchInput{}
        output := &sqs.ChangeMessageVisibilityBatchOutput{}

        mockClient.On("ChangeMessageVisibilityBatch", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeMessageVisibilityBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueue", func(t *testing.T) {
        input := &sqs.CreateQueueInput{}
        output := &sqs.CreateQueueOutput{}

        mockClient.On("CreateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMessage", func(t *testing.T) {
        input := &sqs.DeleteMessageInput{}
        output := &sqs.DeleteMessageOutput{}

        mockClient.On("DeleteMessage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMessageBatch", func(t *testing.T) {
        input := &sqs.DeleteMessageBatchInput{}
        output := &sqs.DeleteMessageBatchOutput{}

        mockClient.On("DeleteMessageBatch", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMessageBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueue", func(t *testing.T) {
        input := &sqs.DeleteQueueInput{}
        output := &sqs.DeleteQueueOutput{}

        mockClient.On("DeleteQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueueAttributes", func(t *testing.T) {
        input := &sqs.GetQueueAttributesInput{}
        output := &sqs.GetQueueAttributesOutput{}

        mockClient.On("GetQueueAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueueAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueueUrl", func(t *testing.T) {
        input := &sqs.GetQueueUrlInput{}
        output := &sqs.GetQueueUrlOutput{}

        mockClient.On("GetQueueUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueueUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeadLetterSourceQueues", func(t *testing.T) {
        input := &sqs.ListDeadLetterSourceQueuesInput{}
        output := &sqs.ListDeadLetterSourceQueuesOutput{}

        mockClient.On("ListDeadLetterSourceQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeadLetterSourceQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMessageMoveTasks", func(t *testing.T) {
        input := &sqs.ListMessageMoveTasksInput{}
        output := &sqs.ListMessageMoveTasksOutput{}

        mockClient.On("ListMessageMoveTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListMessageMoveTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueueTags", func(t *testing.T) {
        input := &sqs.ListQueueTagsInput{}
        output := &sqs.ListQueueTagsOutput{}

        mockClient.On("ListQueueTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueueTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueues", func(t *testing.T) {
        input := &sqs.ListQueuesInput{}
        output := &sqs.ListQueuesOutput{}

        mockClient.On("ListQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurgeQueue", func(t *testing.T) {
        input := &sqs.PurgeQueueInput{}
        output := &sqs.PurgeQueueOutput{}

        mockClient.On("PurgeQueue", ctx, input).Return(output, nil)

        result, err := mockClient.PurgeQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReceiveMessage", func(t *testing.T) {
        input := &sqs.ReceiveMessageInput{}
        output := &sqs.ReceiveMessageOutput{}

        mockClient.On("ReceiveMessage", ctx, input).Return(output, nil)

        result, err := mockClient.ReceiveMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &sqs.RemovePermissionInput{}
        output := &sqs.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendMessage", func(t *testing.T) {
        input := &sqs.SendMessageInput{}
        output := &sqs.SendMessageOutput{}

        mockClient.On("SendMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendMessageBatch", func(t *testing.T) {
        input := &sqs.SendMessageBatchInput{}
        output := &sqs.SendMessageBatchOutput{}

        mockClient.On("SendMessageBatch", ctx, input).Return(output, nil)

        result, err := mockClient.SendMessageBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetQueueAttributes", func(t *testing.T) {
        input := &sqs.SetQueueAttributesInput{}
        output := &sqs.SetQueueAttributesOutput{}

        mockClient.On("SetQueueAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetQueueAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMessageMoveTask", func(t *testing.T) {
        input := &sqs.StartMessageMoveTaskInput{}
        output := &sqs.StartMessageMoveTaskOutput{}

        mockClient.On("StartMessageMoveTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartMessageMoveTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagQueue", func(t *testing.T) {
        input := &sqs.TagQueueInput{}
        output := &sqs.TagQueueOutput{}

        mockClient.On("TagQueue", ctx, input).Return(output, nil)

        result, err := mockClient.TagQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagQueue", func(t *testing.T) {
        input := &sqs.UntagQueueInput{}
        output := &sqs.UntagQueueOutput{}

        mockClient.On("UntagQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UntagQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}