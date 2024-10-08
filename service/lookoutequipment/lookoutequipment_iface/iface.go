
package lookoutequipment_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
)

// IClient defines the interface for lookoutequipment
type IClient interface {
 Options() Options 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateInferenceScheduler(ctx context.Context, params *CreateInferenceSchedulerInput, optFns ...func(*Options)) (*CreateInferenceSchedulerOutput, error) 
 CreateLabel(ctx context.Context, params *CreateLabelInput, optFns ...func(*Options)) (*CreateLabelOutput, error) 
 CreateLabelGroup(ctx context.Context, params *CreateLabelGroupInput, optFns ...func(*Options)) (*CreateLabelGroupOutput, error) 
 CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) 
 CreateRetrainingScheduler(ctx context.Context, params *CreateRetrainingSchedulerInput, optFns ...func(*Options)) (*CreateRetrainingSchedulerOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteInferenceScheduler(ctx context.Context, params *DeleteInferenceSchedulerInput, optFns ...func(*Options)) (*DeleteInferenceSchedulerOutput, error) 
 DeleteLabel(ctx context.Context, params *DeleteLabelInput, optFns ...func(*Options)) (*DeleteLabelOutput, error) 
 DeleteLabelGroup(ctx context.Context, params *DeleteLabelGroupInput, optFns ...func(*Options)) (*DeleteLabelGroupOutput, error) 
 DeleteModel(ctx context.Context, params *DeleteModelInput, optFns ...func(*Options)) (*DeleteModelOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteRetrainingScheduler(ctx context.Context, params *DeleteRetrainingSchedulerInput, optFns ...func(*Options)) (*DeleteRetrainingSchedulerOutput, error) 
 DescribeDataIngestionJob(ctx context.Context, params *DescribeDataIngestionJobInput, optFns ...func(*Options)) (*DescribeDataIngestionJobOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeInferenceScheduler(ctx context.Context, params *DescribeInferenceSchedulerInput, optFns ...func(*Options)) (*DescribeInferenceSchedulerOutput, error) 
 DescribeLabel(ctx context.Context, params *DescribeLabelInput, optFns ...func(*Options)) (*DescribeLabelOutput, error) 
 DescribeLabelGroup(ctx context.Context, params *DescribeLabelGroupInput, optFns ...func(*Options)) (*DescribeLabelGroupOutput, error) 
 DescribeModel(ctx context.Context, params *DescribeModelInput, optFns ...func(*Options)) (*DescribeModelOutput, error) 
 DescribeModelVersion(ctx context.Context, params *DescribeModelVersionInput, optFns ...func(*Options)) (*DescribeModelVersionOutput, error) 
 DescribeResourcePolicy(ctx context.Context, params *DescribeResourcePolicyInput, optFns ...func(*Options)) (*DescribeResourcePolicyOutput, error) 
 DescribeRetrainingScheduler(ctx context.Context, params *DescribeRetrainingSchedulerInput, optFns ...func(*Options)) (*DescribeRetrainingSchedulerOutput, error) 
 ImportDataset(ctx context.Context, params *ImportDatasetInput, optFns ...func(*Options)) (*ImportDatasetOutput, error) 
 ImportModelVersion(ctx context.Context, params *ImportModelVersionInput, optFns ...func(*Options)) (*ImportModelVersionOutput, error) 
 ListDataIngestionJobs(ctx context.Context, params *ListDataIngestionJobsInput, optFns ...func(*Options)) (*ListDataIngestionJobsOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListInferenceEvents(ctx context.Context, params *ListInferenceEventsInput, optFns ...func(*Options)) (*ListInferenceEventsOutput, error) 
 ListInferenceExecutions(ctx context.Context, params *ListInferenceExecutionsInput, optFns ...func(*Options)) (*ListInferenceExecutionsOutput, error) 
 ListInferenceSchedulers(ctx context.Context, params *ListInferenceSchedulersInput, optFns ...func(*Options)) (*ListInferenceSchedulersOutput, error) 
 ListLabelGroups(ctx context.Context, params *ListLabelGroupsInput, optFns ...func(*Options)) (*ListLabelGroupsOutput, error) 
 ListLabels(ctx context.Context, params *ListLabelsInput, optFns ...func(*Options)) (*ListLabelsOutput, error) 
 ListModelVersions(ctx context.Context, params *ListModelVersionsInput, optFns ...func(*Options)) (*ListModelVersionsOutput, error) 
 ListModels(ctx context.Context, params *ListModelsInput, optFns ...func(*Options)) (*ListModelsOutput, error) 
 ListRetrainingSchedulers(ctx context.Context, params *ListRetrainingSchedulersInput, optFns ...func(*Options)) (*ListRetrainingSchedulersOutput, error) 
 ListSensorStatistics(ctx context.Context, params *ListSensorStatisticsInput, optFns ...func(*Options)) (*ListSensorStatisticsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 StartDataIngestionJob(ctx context.Context, params *StartDataIngestionJobInput, optFns ...func(*Options)) (*StartDataIngestionJobOutput, error) 
 StartInferenceScheduler(ctx context.Context, params *StartInferenceSchedulerInput, optFns ...func(*Options)) (*StartInferenceSchedulerOutput, error) 
 StartRetrainingScheduler(ctx context.Context, params *StartRetrainingSchedulerInput, optFns ...func(*Options)) (*StartRetrainingSchedulerOutput, error) 
 StopInferenceScheduler(ctx context.Context, params *StopInferenceSchedulerInput, optFns ...func(*Options)) (*StopInferenceSchedulerOutput, error) 
 StopRetrainingScheduler(ctx context.Context, params *StopRetrainingSchedulerInput, optFns ...func(*Options)) (*StopRetrainingSchedulerOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateActiveModelVersion(ctx context.Context, params *UpdateActiveModelVersionInput, optFns ...func(*Options)) (*UpdateActiveModelVersionOutput, error) 
 UpdateInferenceScheduler(ctx context.Context, params *UpdateInferenceSchedulerInput, optFns ...func(*Options)) (*UpdateInferenceSchedulerOutput, error) 
 UpdateLabelGroup(ctx context.Context, params *UpdateLabelGroupInput, optFns ...func(*Options)) (*UpdateLabelGroupOutput, error) 
 UpdateModel(ctx context.Context, params *UpdateModelInput, optFns ...func(*Options)) (*UpdateModelOutput, error) 
 UpdateRetrainingScheduler(ctx context.Context, params *UpdateRetrainingSchedulerInput, optFns ...func(*Options)) (*UpdateRetrainingSchedulerOutput, error) 
}
