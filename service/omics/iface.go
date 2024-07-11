
package omics

import (
    "github.com/aws/aws-sdk-go-v2/service/omics"
)

// IClient defines the interface for omics
type IClient interface {
 Options() Options 
 AbortMultipartReadSetUpload(ctx context.Context, params *AbortMultipartReadSetUploadInput, optFns ...func(*Options)) (*AbortMultipartReadSetUploadOutput, error) 
 AcceptShare(ctx context.Context, params *AcceptShareInput, optFns ...func(*Options)) (*AcceptShareOutput, error) 
 BatchDeleteReadSet(ctx context.Context, params *BatchDeleteReadSetInput, optFns ...func(*Options)) (*BatchDeleteReadSetOutput, error) 
 CancelAnnotationImportJob(ctx context.Context, params *CancelAnnotationImportJobInput, optFns ...func(*Options)) (*CancelAnnotationImportJobOutput, error) 
 CancelRun(ctx context.Context, params *CancelRunInput, optFns ...func(*Options)) (*CancelRunOutput, error) 
 CancelVariantImportJob(ctx context.Context, params *CancelVariantImportJobInput, optFns ...func(*Options)) (*CancelVariantImportJobOutput, error) 
 CompleteMultipartReadSetUpload(ctx context.Context, params *CompleteMultipartReadSetUploadInput, optFns ...func(*Options)) (*CompleteMultipartReadSetUploadOutput, error) 
 CreateAnnotationStore(ctx context.Context, params *CreateAnnotationStoreInput, optFns ...func(*Options)) (*CreateAnnotationStoreOutput, error) 
 CreateAnnotationStoreVersion(ctx context.Context, params *CreateAnnotationStoreVersionInput, optFns ...func(*Options)) (*CreateAnnotationStoreVersionOutput, error) 
 CreateMultipartReadSetUpload(ctx context.Context, params *CreateMultipartReadSetUploadInput, optFns ...func(*Options)) (*CreateMultipartReadSetUploadOutput, error) 
 CreateReferenceStore(ctx context.Context, params *CreateReferenceStoreInput, optFns ...func(*Options)) (*CreateReferenceStoreOutput, error) 
 CreateRunGroup(ctx context.Context, params *CreateRunGroupInput, optFns ...func(*Options)) (*CreateRunGroupOutput, error) 
 CreateSequenceStore(ctx context.Context, params *CreateSequenceStoreInput, optFns ...func(*Options)) (*CreateSequenceStoreOutput, error) 
 CreateShare(ctx context.Context, params *CreateShareInput, optFns ...func(*Options)) (*CreateShareOutput, error) 
 CreateVariantStore(ctx context.Context, params *CreateVariantStoreInput, optFns ...func(*Options)) (*CreateVariantStoreOutput, error) 
 CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) 
 DeleteAnnotationStore(ctx context.Context, params *DeleteAnnotationStoreInput, optFns ...func(*Options)) (*DeleteAnnotationStoreOutput, error) 
 DeleteAnnotationStoreVersions(ctx context.Context, params *DeleteAnnotationStoreVersionsInput, optFns ...func(*Options)) (*DeleteAnnotationStoreVersionsOutput, error) 
 DeleteReference(ctx context.Context, params *DeleteReferenceInput, optFns ...func(*Options)) (*DeleteReferenceOutput, error) 
 DeleteReferenceStore(ctx context.Context, params *DeleteReferenceStoreInput, optFns ...func(*Options)) (*DeleteReferenceStoreOutput, error) 
 DeleteRun(ctx context.Context, params *DeleteRunInput, optFns ...func(*Options)) (*DeleteRunOutput, error) 
 DeleteRunGroup(ctx context.Context, params *DeleteRunGroupInput, optFns ...func(*Options)) (*DeleteRunGroupOutput, error) 
 DeleteSequenceStore(ctx context.Context, params *DeleteSequenceStoreInput, optFns ...func(*Options)) (*DeleteSequenceStoreOutput, error) 
 DeleteShare(ctx context.Context, params *DeleteShareInput, optFns ...func(*Options)) (*DeleteShareOutput, error) 
 DeleteVariantStore(ctx context.Context, params *DeleteVariantStoreInput, optFns ...func(*Options)) (*DeleteVariantStoreOutput, error) 
 DeleteWorkflow(ctx context.Context, params *DeleteWorkflowInput, optFns ...func(*Options)) (*DeleteWorkflowOutput, error) 
 GetAnnotationImportJob(ctx context.Context, params *GetAnnotationImportJobInput, optFns ...func(*Options)) (*GetAnnotationImportJobOutput, error) 
 GetAnnotationStore(ctx context.Context, params *GetAnnotationStoreInput, optFns ...func(*Options)) (*GetAnnotationStoreOutput, error) 
 GetAnnotationStoreVersion(ctx context.Context, params *GetAnnotationStoreVersionInput, optFns ...func(*Options)) (*GetAnnotationStoreVersionOutput, error) 
 GetReadSet(ctx context.Context, params *GetReadSetInput, optFns ...func(*Options)) (*GetReadSetOutput, error) 
 GetReadSetActivationJob(ctx context.Context, params *GetReadSetActivationJobInput, optFns ...func(*Options)) (*GetReadSetActivationJobOutput, error) 
 GetReadSetExportJob(ctx context.Context, params *GetReadSetExportJobInput, optFns ...func(*Options)) (*GetReadSetExportJobOutput, error) 
 GetReadSetImportJob(ctx context.Context, params *GetReadSetImportJobInput, optFns ...func(*Options)) (*GetReadSetImportJobOutput, error) 
 GetReadSetMetadata(ctx context.Context, params *GetReadSetMetadataInput, optFns ...func(*Options)) (*GetReadSetMetadataOutput, error) 
 GetReference(ctx context.Context, params *GetReferenceInput, optFns ...func(*Options)) (*GetReferenceOutput, error) 
 GetReferenceImportJob(ctx context.Context, params *GetReferenceImportJobInput, optFns ...func(*Options)) (*GetReferenceImportJobOutput, error) 
 GetReferenceMetadata(ctx context.Context, params *GetReferenceMetadataInput, optFns ...func(*Options)) (*GetReferenceMetadataOutput, error) 
 GetReferenceStore(ctx context.Context, params *GetReferenceStoreInput, optFns ...func(*Options)) (*GetReferenceStoreOutput, error) 
 GetRun(ctx context.Context, params *GetRunInput, optFns ...func(*Options)) (*GetRunOutput, error) 
 GetRunGroup(ctx context.Context, params *GetRunGroupInput, optFns ...func(*Options)) (*GetRunGroupOutput, error) 
 GetRunTask(ctx context.Context, params *GetRunTaskInput, optFns ...func(*Options)) (*GetRunTaskOutput, error) 
 GetSequenceStore(ctx context.Context, params *GetSequenceStoreInput, optFns ...func(*Options)) (*GetSequenceStoreOutput, error) 
 GetShare(ctx context.Context, params *GetShareInput, optFns ...func(*Options)) (*GetShareOutput, error) 
 GetVariantImportJob(ctx context.Context, params *GetVariantImportJobInput, optFns ...func(*Options)) (*GetVariantImportJobOutput, error) 
 GetVariantStore(ctx context.Context, params *GetVariantStoreInput, optFns ...func(*Options)) (*GetVariantStoreOutput, error) 
 GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) 
 ListAnnotationImportJobs(ctx context.Context, params *ListAnnotationImportJobsInput, optFns ...func(*Options)) (*ListAnnotationImportJobsOutput, error) 
 ListAnnotationStoreVersions(ctx context.Context, params *ListAnnotationStoreVersionsInput, optFns ...func(*Options)) (*ListAnnotationStoreVersionsOutput, error) 
 ListAnnotationStores(ctx context.Context, params *ListAnnotationStoresInput, optFns ...func(*Options)) (*ListAnnotationStoresOutput, error) 
 ListMultipartReadSetUploads(ctx context.Context, params *ListMultipartReadSetUploadsInput, optFns ...func(*Options)) (*ListMultipartReadSetUploadsOutput, error) 
 ListReadSetActivationJobs(ctx context.Context, params *ListReadSetActivationJobsInput, optFns ...func(*Options)) (*ListReadSetActivationJobsOutput, error) 
 ListReadSetExportJobs(ctx context.Context, params *ListReadSetExportJobsInput, optFns ...func(*Options)) (*ListReadSetExportJobsOutput, error) 
 ListReadSetImportJobs(ctx context.Context, params *ListReadSetImportJobsInput, optFns ...func(*Options)) (*ListReadSetImportJobsOutput, error) 
 ListReadSetUploadParts(ctx context.Context, params *ListReadSetUploadPartsInput, optFns ...func(*Options)) (*ListReadSetUploadPartsOutput, error) 
 ListReadSets(ctx context.Context, params *ListReadSetsInput, optFns ...func(*Options)) (*ListReadSetsOutput, error) 
 ListReferenceImportJobs(ctx context.Context, params *ListReferenceImportJobsInput, optFns ...func(*Options)) (*ListReferenceImportJobsOutput, error) 
 ListReferenceStores(ctx context.Context, params *ListReferenceStoresInput, optFns ...func(*Options)) (*ListReferenceStoresOutput, error) 
 ListReferences(ctx context.Context, params *ListReferencesInput, optFns ...func(*Options)) (*ListReferencesOutput, error) 
 ListRunGroups(ctx context.Context, params *ListRunGroupsInput, optFns ...func(*Options)) (*ListRunGroupsOutput, error) 
 ListRunTasks(ctx context.Context, params *ListRunTasksInput, optFns ...func(*Options)) (*ListRunTasksOutput, error) 
 ListRuns(ctx context.Context, params *ListRunsInput, optFns ...func(*Options)) (*ListRunsOutput, error) 
 ListSequenceStores(ctx context.Context, params *ListSequenceStoresInput, optFns ...func(*Options)) (*ListSequenceStoresOutput, error) 
 ListShares(ctx context.Context, params *ListSharesInput, optFns ...func(*Options)) (*ListSharesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVariantImportJobs(ctx context.Context, params *ListVariantImportJobsInput, optFns ...func(*Options)) (*ListVariantImportJobsOutput, error) 
 ListVariantStores(ctx context.Context, params *ListVariantStoresInput, optFns ...func(*Options)) (*ListVariantStoresOutput, error) 
 ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) 
 StartAnnotationImportJob(ctx context.Context, params *StartAnnotationImportJobInput, optFns ...func(*Options)) (*StartAnnotationImportJobOutput, error) 
 StartReadSetActivationJob(ctx context.Context, params *StartReadSetActivationJobInput, optFns ...func(*Options)) (*StartReadSetActivationJobOutput, error) 
 StartReadSetExportJob(ctx context.Context, params *StartReadSetExportJobInput, optFns ...func(*Options)) (*StartReadSetExportJobOutput, error) 
 StartReadSetImportJob(ctx context.Context, params *StartReadSetImportJobInput, optFns ...func(*Options)) (*StartReadSetImportJobOutput, error) 
 StartReferenceImportJob(ctx context.Context, params *StartReferenceImportJobInput, optFns ...func(*Options)) (*StartReferenceImportJobOutput, error) 
 StartRun(ctx context.Context, params *StartRunInput, optFns ...func(*Options)) (*StartRunOutput, error) 
 StartVariantImportJob(ctx context.Context, params *StartVariantImportJobInput, optFns ...func(*Options)) (*StartVariantImportJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAnnotationStore(ctx context.Context, params *UpdateAnnotationStoreInput, optFns ...func(*Options)) (*UpdateAnnotationStoreOutput, error) 
 UpdateAnnotationStoreVersion(ctx context.Context, params *UpdateAnnotationStoreVersionInput, optFns ...func(*Options)) (*UpdateAnnotationStoreVersionOutput, error) 
 UpdateRunGroup(ctx context.Context, params *UpdateRunGroupInput, optFns ...func(*Options)) (*UpdateRunGroupOutput, error) 
 UpdateVariantStore(ctx context.Context, params *UpdateVariantStoreInput, optFns ...func(*Options)) (*UpdateVariantStoreOutput, error) 
 UpdateWorkflow(ctx context.Context, params *UpdateWorkflowInput, optFns ...func(*Options)) (*UpdateWorkflowOutput, error) 
 UploadReadSetPart(ctx context.Context, params *UploadReadSetPartInput, optFns ...func(*Options)) (*UploadReadSetPartOutput, error) 
}
