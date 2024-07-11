
package fis

import (
    "github.com/aws/aws-sdk-go-v2/service/fis"
)

// IClient defines the interface for fis
type IClient interface {
 Options() Options 
 CreateExperimentTemplate(ctx context.Context, params *CreateExperimentTemplateInput, optFns ...func(*Options)) (*CreateExperimentTemplateOutput, error) 
 CreateTargetAccountConfiguration(ctx context.Context, params *CreateTargetAccountConfigurationInput, optFns ...func(*Options)) (*CreateTargetAccountConfigurationOutput, error) 
 DeleteExperimentTemplate(ctx context.Context, params *DeleteExperimentTemplateInput, optFns ...func(*Options)) (*DeleteExperimentTemplateOutput, error) 
 DeleteTargetAccountConfiguration(ctx context.Context, params *DeleteTargetAccountConfigurationInput, optFns ...func(*Options)) (*DeleteTargetAccountConfigurationOutput, error) 
 GetAction(ctx context.Context, params *GetActionInput, optFns ...func(*Options)) (*GetActionOutput, error) 
 GetExperiment(ctx context.Context, params *GetExperimentInput, optFns ...func(*Options)) (*GetExperimentOutput, error) 
 GetExperimentTargetAccountConfiguration(ctx context.Context, params *GetExperimentTargetAccountConfigurationInput, optFns ...func(*Options)) (*GetExperimentTargetAccountConfigurationOutput, error) 
 GetExperimentTemplate(ctx context.Context, params *GetExperimentTemplateInput, optFns ...func(*Options)) (*GetExperimentTemplateOutput, error) 
 GetTargetAccountConfiguration(ctx context.Context, params *GetTargetAccountConfigurationInput, optFns ...func(*Options)) (*GetTargetAccountConfigurationOutput, error) 
 GetTargetResourceType(ctx context.Context, params *GetTargetResourceTypeInput, optFns ...func(*Options)) (*GetTargetResourceTypeOutput, error) 
 ListActions(ctx context.Context, params *ListActionsInput, optFns ...func(*Options)) (*ListActionsOutput, error) 
 ListExperimentResolvedTargets(ctx context.Context, params *ListExperimentResolvedTargetsInput, optFns ...func(*Options)) (*ListExperimentResolvedTargetsOutput, error) 
 ListExperimentTargetAccountConfigurations(ctx context.Context, params *ListExperimentTargetAccountConfigurationsInput, optFns ...func(*Options)) (*ListExperimentTargetAccountConfigurationsOutput, error) 
 ListExperimentTemplates(ctx context.Context, params *ListExperimentTemplatesInput, optFns ...func(*Options)) (*ListExperimentTemplatesOutput, error) 
 ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTargetAccountConfigurations(ctx context.Context, params *ListTargetAccountConfigurationsInput, optFns ...func(*Options)) (*ListTargetAccountConfigurationsOutput, error) 
 ListTargetResourceTypes(ctx context.Context, params *ListTargetResourceTypesInput, optFns ...func(*Options)) (*ListTargetResourceTypesOutput, error) 
 StartExperiment(ctx context.Context, params *StartExperimentInput, optFns ...func(*Options)) (*StartExperimentOutput, error) 
 StopExperiment(ctx context.Context, params *StopExperimentInput, optFns ...func(*Options)) (*StopExperimentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateExperimentTemplate(ctx context.Context, params *UpdateExperimentTemplateInput, optFns ...func(*Options)) (*UpdateExperimentTemplateOutput, error) 
 UpdateTargetAccountConfiguration(ctx context.Context, params *UpdateTargetAccountConfigurationInput, optFns ...func(*Options)) (*UpdateTargetAccountConfigurationOutput, error) 
}
