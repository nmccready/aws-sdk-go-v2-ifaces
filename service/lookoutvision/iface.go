
package lookoutvision

import (
    "github.com/aws/aws-sdk-go-v2/service/lookoutvision"
)

// IClient defines the interface for lookoutvision
type IClient interface {
 Options() Options 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteModel(ctx context.Context, params *DeleteModelInput, optFns ...func(*Options)) (*DeleteModelOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeModel(ctx context.Context, params *DescribeModelInput, optFns ...func(*Options)) (*DescribeModelOutput, error) 
 DescribeModelPackagingJob(ctx context.Context, params *DescribeModelPackagingJobInput, optFns ...func(*Options)) (*DescribeModelPackagingJobOutput, error) 
 DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) 
 DetectAnomalies(ctx context.Context, params *DetectAnomaliesInput, optFns ...func(*Options)) (*DetectAnomaliesOutput, error) 
 ListDatasetEntries(ctx context.Context, params *ListDatasetEntriesInput, optFns ...func(*Options)) (*ListDatasetEntriesOutput, error) 
 ListModelPackagingJobs(ctx context.Context, params *ListModelPackagingJobsInput, optFns ...func(*Options)) (*ListModelPackagingJobsOutput, error) 
 ListModels(ctx context.Context, params *ListModelsInput, optFns ...func(*Options)) (*ListModelsOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartModel(ctx context.Context, params *StartModelInput, optFns ...func(*Options)) (*StartModelOutput, error) 
 StartModelPackagingJob(ctx context.Context, params *StartModelPackagingJobInput, optFns ...func(*Options)) (*StartModelPackagingJobOutput, error) 
 StopModel(ctx context.Context, params *StopModelInput, optFns ...func(*Options)) (*StopModelOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDatasetEntries(ctx context.Context, params *UpdateDatasetEntriesInput, optFns ...func(*Options)) (*UpdateDatasetEntriesOutput, error) 
}
