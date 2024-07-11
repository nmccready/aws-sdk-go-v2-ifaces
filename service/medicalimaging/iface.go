
package medicalimaging

import (
    "github.com/aws/aws-sdk-go-v2/service/medicalimaging"
)

// IClient defines the interface for medicalimaging
type IClient interface {
 Options() Options 
 CopyImageSet(ctx context.Context, params *CopyImageSetInput, optFns ...func(*Options)) (*CopyImageSetOutput, error) 
 CreateDatastore(ctx context.Context, params *CreateDatastoreInput, optFns ...func(*Options)) (*CreateDatastoreOutput, error) 
 DeleteDatastore(ctx context.Context, params *DeleteDatastoreInput, optFns ...func(*Options)) (*DeleteDatastoreOutput, error) 
 DeleteImageSet(ctx context.Context, params *DeleteImageSetInput, optFns ...func(*Options)) (*DeleteImageSetOutput, error) 
 GetDICOMImportJob(ctx context.Context, params *GetDICOMImportJobInput, optFns ...func(*Options)) (*GetDICOMImportJobOutput, error) 
 GetDatastore(ctx context.Context, params *GetDatastoreInput, optFns ...func(*Options)) (*GetDatastoreOutput, error) 
 GetImageFrame(ctx context.Context, params *GetImageFrameInput, optFns ...func(*Options)) (*GetImageFrameOutput, error) 
 GetImageSet(ctx context.Context, params *GetImageSetInput, optFns ...func(*Options)) (*GetImageSetOutput, error) 
 GetImageSetMetadata(ctx context.Context, params *GetImageSetMetadataInput, optFns ...func(*Options)) (*GetImageSetMetadataOutput, error) 
 ListDICOMImportJobs(ctx context.Context, params *ListDICOMImportJobsInput, optFns ...func(*Options)) (*ListDICOMImportJobsOutput, error) 
 ListDatastores(ctx context.Context, params *ListDatastoresInput, optFns ...func(*Options)) (*ListDatastoresOutput, error) 
 ListImageSetVersions(ctx context.Context, params *ListImageSetVersionsInput, optFns ...func(*Options)) (*ListImageSetVersionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 SearchImageSets(ctx context.Context, params *SearchImageSetsInput, optFns ...func(*Options)) (*SearchImageSetsOutput, error) 
 StartDICOMImportJob(ctx context.Context, params *StartDICOMImportJobInput, optFns ...func(*Options)) (*StartDICOMImportJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateImageSetMetadata(ctx context.Context, params *UpdateImageSetMetadataInput, optFns ...func(*Options)) (*UpdateImageSetMetadataOutput, error) 
}
