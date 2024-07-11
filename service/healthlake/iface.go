
package healthlake

import (
    "github.com/aws/aws-sdk-go-v2/service/healthlake"
)

// IHealthlake defines the interface for healthlake
type IHealthlake interface {
 Options() Options 
 CreateFHIRDatastore(ctx context.Context, params *CreateFHIRDatastoreInput, optFns ...func(*Options)) (*CreateFHIRDatastoreOutput, error) 
 DeleteFHIRDatastore(ctx context.Context, params *DeleteFHIRDatastoreInput, optFns ...func(*Options)) (*DeleteFHIRDatastoreOutput, error) 
 DescribeFHIRDatastore(ctx context.Context, params *DescribeFHIRDatastoreInput, optFns ...func(*Options)) (*DescribeFHIRDatastoreOutput, error) 
 DescribeFHIRExportJob(ctx context.Context, params *DescribeFHIRExportJobInput, optFns ...func(*Options)) (*DescribeFHIRExportJobOutput, error) 
 DescribeFHIRImportJob(ctx context.Context, params *DescribeFHIRImportJobInput, optFns ...func(*Options)) (*DescribeFHIRImportJobOutput, error) 
 ListFHIRDatastores(ctx context.Context, params *ListFHIRDatastoresInput, optFns ...func(*Options)) (*ListFHIRDatastoresOutput, error) 
 ListFHIRExportJobs(ctx context.Context, params *ListFHIRExportJobsInput, optFns ...func(*Options)) (*ListFHIRExportJobsOutput, error) 
 ListFHIRImportJobs(ctx context.Context, params *ListFHIRImportJobsInput, optFns ...func(*Options)) (*ListFHIRImportJobsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartFHIRExportJob(ctx context.Context, params *StartFHIRExportJobInput, optFns ...func(*Options)) (*StartFHIRExportJobOutput, error) 
 StartFHIRImportJob(ctx context.Context, params *StartFHIRImportJobInput, optFns ...func(*Options)) (*StartFHIRImportJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
