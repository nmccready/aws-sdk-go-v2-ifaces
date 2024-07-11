
package bedrock

import (
    "github.com/aws/aws-sdk-go-v2/service/bedrock"
)

// IBedrock defines the interface for bedrock
type IBedrock interface {
 Options() Options 
 CreateEvaluationJob(ctx context.Context, params *CreateEvaluationJobInput, optFns ...func(*Options)) (*CreateEvaluationJobOutput, error) 
 CreateGuardrail(ctx context.Context, params *CreateGuardrailInput, optFns ...func(*Options)) (*CreateGuardrailOutput, error) 
 CreateGuardrailVersion(ctx context.Context, params *CreateGuardrailVersionInput, optFns ...func(*Options)) (*CreateGuardrailVersionOutput, error) 
 CreateModelCustomizationJob(ctx context.Context, params *CreateModelCustomizationJobInput, optFns ...func(*Options)) (*CreateModelCustomizationJobOutput, error) 
 CreateProvisionedModelThroughput(ctx context.Context, params *CreateProvisionedModelThroughputInput, optFns ...func(*Options)) (*CreateProvisionedModelThroughputOutput, error) 
 DeleteCustomModel(ctx context.Context, params *DeleteCustomModelInput, optFns ...func(*Options)) (*DeleteCustomModelOutput, error) 
 DeleteGuardrail(ctx context.Context, params *DeleteGuardrailInput, optFns ...func(*Options)) (*DeleteGuardrailOutput, error) 
 DeleteModelInvocationLoggingConfiguration(ctx context.Context, params *DeleteModelInvocationLoggingConfigurationInput, optFns ...func(*Options)) (*DeleteModelInvocationLoggingConfigurationOutput, error) 
 DeleteProvisionedModelThroughput(ctx context.Context, params *DeleteProvisionedModelThroughputInput, optFns ...func(*Options)) (*DeleteProvisionedModelThroughputOutput, error) 
 GetCustomModel(ctx context.Context, params *GetCustomModelInput, optFns ...func(*Options)) (*GetCustomModelOutput, error) 
 GetEvaluationJob(ctx context.Context, params *GetEvaluationJobInput, optFns ...func(*Options)) (*GetEvaluationJobOutput, error) 
 GetFoundationModel(ctx context.Context, params *GetFoundationModelInput, optFns ...func(*Options)) (*GetFoundationModelOutput, error) 
 GetGuardrail(ctx context.Context, params *GetGuardrailInput, optFns ...func(*Options)) (*GetGuardrailOutput, error) 
 GetModelCustomizationJob(ctx context.Context, params *GetModelCustomizationJobInput, optFns ...func(*Options)) (*GetModelCustomizationJobOutput, error) 
 GetModelInvocationLoggingConfiguration(ctx context.Context, params *GetModelInvocationLoggingConfigurationInput, optFns ...func(*Options)) (*GetModelInvocationLoggingConfigurationOutput, error) 
 GetProvisionedModelThroughput(ctx context.Context, params *GetProvisionedModelThroughputInput, optFns ...func(*Options)) (*GetProvisionedModelThroughputOutput, error) 
 ListCustomModels(ctx context.Context, params *ListCustomModelsInput, optFns ...func(*Options)) (*ListCustomModelsOutput, error) 
 ListEvaluationJobs(ctx context.Context, params *ListEvaluationJobsInput, optFns ...func(*Options)) (*ListEvaluationJobsOutput, error) 
 ListFoundationModels(ctx context.Context, params *ListFoundationModelsInput, optFns ...func(*Options)) (*ListFoundationModelsOutput, error) 
 ListGuardrails(ctx context.Context, params *ListGuardrailsInput, optFns ...func(*Options)) (*ListGuardrailsOutput, error) 
 ListModelCustomizationJobs(ctx context.Context, params *ListModelCustomizationJobsInput, optFns ...func(*Options)) (*ListModelCustomizationJobsOutput, error) 
 ListProvisionedModelThroughputs(ctx context.Context, params *ListProvisionedModelThroughputsInput, optFns ...func(*Options)) (*ListProvisionedModelThroughputsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutModelInvocationLoggingConfiguration(ctx context.Context, params *PutModelInvocationLoggingConfigurationInput, optFns ...func(*Options)) (*PutModelInvocationLoggingConfigurationOutput, error) 
 StopEvaluationJob(ctx context.Context, params *StopEvaluationJobInput, optFns ...func(*Options)) (*StopEvaluationJobOutput, error) 
 StopModelCustomizationJob(ctx context.Context, params *StopModelCustomizationJobInput, optFns ...func(*Options)) (*StopModelCustomizationJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGuardrail(ctx context.Context, params *UpdateGuardrailInput, optFns ...func(*Options)) (*UpdateGuardrailOutput, error) 
 UpdateProvisionedModelThroughput(ctx context.Context, params *UpdateProvisionedModelThroughputInput, optFns ...func(*Options)) (*UpdateProvisionedModelThroughputOutput, error) 
}
