
package dataexchange

import (
    "github.com/aws/aws-sdk-go-v2/service/dataexchange"
)

// IClient defines the interface for dataexchange
type IClient interface {
 Options() Options 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CreateDataSet(ctx context.Context, params *CreateDataSetInput, optFns ...func(*Options)) (*CreateDataSetOutput, error) 
 CreateEventAction(ctx context.Context, params *CreateEventActionInput, optFns ...func(*Options)) (*CreateEventActionOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateRevision(ctx context.Context, params *CreateRevisionInput, optFns ...func(*Options)) (*CreateRevisionOutput, error) 
 DeleteAsset(ctx context.Context, params *DeleteAssetInput, optFns ...func(*Options)) (*DeleteAssetOutput, error) 
 DeleteDataSet(ctx context.Context, params *DeleteDataSetInput, optFns ...func(*Options)) (*DeleteDataSetOutput, error) 
 DeleteEventAction(ctx context.Context, params *DeleteEventActionInput, optFns ...func(*Options)) (*DeleteEventActionOutput, error) 
 DeleteRevision(ctx context.Context, params *DeleteRevisionInput, optFns ...func(*Options)) (*DeleteRevisionOutput, error) 
 GetAsset(ctx context.Context, params *GetAssetInput, optFns ...func(*Options)) (*GetAssetOutput, error) 
 GetDataSet(ctx context.Context, params *GetDataSetInput, optFns ...func(*Options)) (*GetDataSetOutput, error) 
 GetEventAction(ctx context.Context, params *GetEventActionInput, optFns ...func(*Options)) (*GetEventActionOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetRevision(ctx context.Context, params *GetRevisionInput, optFns ...func(*Options)) (*GetRevisionOutput, error) 
 ListDataSetRevisions(ctx context.Context, params *ListDataSetRevisionsInput, optFns ...func(*Options)) (*ListDataSetRevisionsOutput, error) 
 ListDataSets(ctx context.Context, params *ListDataSetsInput, optFns ...func(*Options)) (*ListDataSetsOutput, error) 
 ListEventActions(ctx context.Context, params *ListEventActionsInput, optFns ...func(*Options)) (*ListEventActionsOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListRevisionAssets(ctx context.Context, params *ListRevisionAssetsInput, optFns ...func(*Options)) (*ListRevisionAssetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RevokeRevision(ctx context.Context, params *RevokeRevisionInput, optFns ...func(*Options)) (*RevokeRevisionOutput, error) 
 SendApiAsset(ctx context.Context, params *SendApiAssetInput, optFns ...func(*Options)) (*SendApiAssetOutput, error) 
 SendDataSetNotification(ctx context.Context, params *SendDataSetNotificationInput, optFns ...func(*Options)) (*SendDataSetNotificationOutput, error) 
 StartJob(ctx context.Context, params *StartJobInput, optFns ...func(*Options)) (*StartJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAsset(ctx context.Context, params *UpdateAssetInput, optFns ...func(*Options)) (*UpdateAssetOutput, error) 
 UpdateDataSet(ctx context.Context, params *UpdateDataSetInput, optFns ...func(*Options)) (*UpdateDataSetOutput, error) 
 UpdateEventAction(ctx context.Context, params *UpdateEventActionInput, optFns ...func(*Options)) (*UpdateEventActionOutput, error) 
 UpdateRevision(ctx context.Context, params *UpdateRevisionInput, optFns ...func(*Options)) (*UpdateRevisionOutput, error) 
}
