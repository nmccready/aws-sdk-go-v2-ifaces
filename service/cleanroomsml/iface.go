
package cleanroomsml

import (
    "github.com/aws/aws-sdk-go-v2/service/cleanroomsml"
)

// IClient defines the interface for cleanroomsml
type IClient interface {
 Options() Options 
 CreateAudienceModel(ctx context.Context, params *CreateAudienceModelInput, optFns ...func(*Options)) (*CreateAudienceModelOutput, error) 
 CreateConfiguredAudienceModel(ctx context.Context, params *CreateConfiguredAudienceModelInput, optFns ...func(*Options)) (*CreateConfiguredAudienceModelOutput, error) 
 CreateTrainingDataset(ctx context.Context, params *CreateTrainingDatasetInput, optFns ...func(*Options)) (*CreateTrainingDatasetOutput, error) 
 DeleteAudienceGenerationJob(ctx context.Context, params *DeleteAudienceGenerationJobInput, optFns ...func(*Options)) (*DeleteAudienceGenerationJobOutput, error) 
 DeleteAudienceModel(ctx context.Context, params *DeleteAudienceModelInput, optFns ...func(*Options)) (*DeleteAudienceModelOutput, error) 
 DeleteConfiguredAudienceModel(ctx context.Context, params *DeleteConfiguredAudienceModelInput, optFns ...func(*Options)) (*DeleteConfiguredAudienceModelOutput, error) 
 DeleteConfiguredAudienceModelPolicy(ctx context.Context, params *DeleteConfiguredAudienceModelPolicyInput, optFns ...func(*Options)) (*DeleteConfiguredAudienceModelPolicyOutput, error) 
 DeleteTrainingDataset(ctx context.Context, params *DeleteTrainingDatasetInput, optFns ...func(*Options)) (*DeleteTrainingDatasetOutput, error) 
 GetAudienceGenerationJob(ctx context.Context, params *GetAudienceGenerationJobInput, optFns ...func(*Options)) (*GetAudienceGenerationJobOutput, error) 
 GetAudienceModel(ctx context.Context, params *GetAudienceModelInput, optFns ...func(*Options)) (*GetAudienceModelOutput, error) 
 GetConfiguredAudienceModel(ctx context.Context, params *GetConfiguredAudienceModelInput, optFns ...func(*Options)) (*GetConfiguredAudienceModelOutput, error) 
 GetConfiguredAudienceModelPolicy(ctx context.Context, params *GetConfiguredAudienceModelPolicyInput, optFns ...func(*Options)) (*GetConfiguredAudienceModelPolicyOutput, error) 
 GetTrainingDataset(ctx context.Context, params *GetTrainingDatasetInput, optFns ...func(*Options)) (*GetTrainingDatasetOutput, error) 
 ListAudienceExportJobs(ctx context.Context, params *ListAudienceExportJobsInput, optFns ...func(*Options)) (*ListAudienceExportJobsOutput, error) 
 ListAudienceGenerationJobs(ctx context.Context, params *ListAudienceGenerationJobsInput, optFns ...func(*Options)) (*ListAudienceGenerationJobsOutput, error) 
 ListAudienceModels(ctx context.Context, params *ListAudienceModelsInput, optFns ...func(*Options)) (*ListAudienceModelsOutput, error) 
 ListConfiguredAudienceModels(ctx context.Context, params *ListConfiguredAudienceModelsInput, optFns ...func(*Options)) (*ListConfiguredAudienceModelsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrainingDatasets(ctx context.Context, params *ListTrainingDatasetsInput, optFns ...func(*Options)) (*ListTrainingDatasetsOutput, error) 
 PutConfiguredAudienceModelPolicy(ctx context.Context, params *PutConfiguredAudienceModelPolicyInput, optFns ...func(*Options)) (*PutConfiguredAudienceModelPolicyOutput, error) 
 StartAudienceExportJob(ctx context.Context, params *StartAudienceExportJobInput, optFns ...func(*Options)) (*StartAudienceExportJobOutput, error) 
 StartAudienceGenerationJob(ctx context.Context, params *StartAudienceGenerationJobInput, optFns ...func(*Options)) (*StartAudienceGenerationJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConfiguredAudienceModel(ctx context.Context, params *UpdateConfiguredAudienceModelInput, optFns ...func(*Options)) (*UpdateConfiguredAudienceModelOutput, error) 
}
