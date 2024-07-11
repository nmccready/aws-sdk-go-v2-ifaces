
package sqs

import (
    "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// ISqs defines the interface for sqs
type ISqs interface {
 Options() Options 
 AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) 
 CancelMessageMoveTask(ctx context.Context, params *CancelMessageMoveTaskInput, optFns ...func(*Options)) (*CancelMessageMoveTaskOutput, error) 
 ChangeMessageVisibility(ctx context.Context, params *ChangeMessageVisibilityInput, optFns ...func(*Options)) (*ChangeMessageVisibilityOutput, error) 
 ChangeMessageVisibilityBatch(ctx context.Context, params *ChangeMessageVisibilityBatchInput, optFns ...func(*Options)) (*ChangeMessageVisibilityBatchOutput, error) 
 CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) 
 DeleteMessage(ctx context.Context, params *DeleteMessageInput, optFns ...func(*Options)) (*DeleteMessageOutput, error) 
 DeleteMessageBatch(ctx context.Context, params *DeleteMessageBatchInput, optFns ...func(*Options)) (*DeleteMessageBatchOutput, error) 
 DeleteQueue(ctx context.Context, params *DeleteQueueInput, optFns ...func(*Options)) (*DeleteQueueOutput, error) 
 GetQueueAttributes(ctx context.Context, params *GetQueueAttributesInput, optFns ...func(*Options)) (*GetQueueAttributesOutput, error) 
 GetQueueUrl(ctx context.Context, params *GetQueueUrlInput, optFns ...func(*Options)) (*GetQueueUrlOutput, error) 
 ListDeadLetterSourceQueues(ctx context.Context, params *ListDeadLetterSourceQueuesInput, optFns ...func(*Options)) (*ListDeadLetterSourceQueuesOutput, error) 
 ListMessageMoveTasks(ctx context.Context, params *ListMessageMoveTasksInput, optFns ...func(*Options)) (*ListMessageMoveTasksOutput, error) 
 ListQueueTags(ctx context.Context, params *ListQueueTagsInput, optFns ...func(*Options)) (*ListQueueTagsOutput, error) 
 ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) 
 PurgeQueue(ctx context.Context, params *PurgeQueueInput, optFns ...func(*Options)) (*PurgeQueueOutput, error) 
 ReceiveMessage(ctx context.Context, params *ReceiveMessageInput, optFns ...func(*Options)) (*ReceiveMessageOutput, error) 
 RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) 
 SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) 
 SendMessageBatch(ctx context.Context, params *SendMessageBatchInput, optFns ...func(*Options)) (*SendMessageBatchOutput, error) 
 SetQueueAttributes(ctx context.Context, params *SetQueueAttributesInput, optFns ...func(*Options)) (*SetQueueAttributesOutput, error) 
 StartMessageMoveTask(ctx context.Context, params *StartMessageMoveTaskInput, optFns ...func(*Options)) (*StartMessageMoveTaskOutput, error) 
 TagQueue(ctx context.Context, params *TagQueueInput, optFns ...func(*Options)) (*TagQueueOutput, error) 
 UntagQueue(ctx context.Context, params *UntagQueueInput, optFns ...func(*Options)) (*UntagQueueOutput, error) 
}
