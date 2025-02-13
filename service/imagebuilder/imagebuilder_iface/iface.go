
package imagebuilder_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/imagebuilder"
)

// IClient defines the interface for imagebuilder
type IClient interface {
 Options() Options 
 CancelImageCreation(ctx context.Context, params *CancelImageCreationInput, optFns ...func(*Options)) (*CancelImageCreationOutput, error) 
 CancelLifecycleExecution(ctx context.Context, params *CancelLifecycleExecutionInput, optFns ...func(*Options)) (*CancelLifecycleExecutionOutput, error) 
 CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) 
 CreateContainerRecipe(ctx context.Context, params *CreateContainerRecipeInput, optFns ...func(*Options)) (*CreateContainerRecipeOutput, error) 
 CreateDistributionConfiguration(ctx context.Context, params *CreateDistributionConfigurationInput, optFns ...func(*Options)) (*CreateDistributionConfigurationOutput, error) 
 CreateImage(ctx context.Context, params *CreateImageInput, optFns ...func(*Options)) (*CreateImageOutput, error) 
 CreateImagePipeline(ctx context.Context, params *CreateImagePipelineInput, optFns ...func(*Options)) (*CreateImagePipelineOutput, error) 
 CreateImageRecipe(ctx context.Context, params *CreateImageRecipeInput, optFns ...func(*Options)) (*CreateImageRecipeOutput, error) 
 CreateInfrastructureConfiguration(ctx context.Context, params *CreateInfrastructureConfigurationInput, optFns ...func(*Options)) (*CreateInfrastructureConfigurationOutput, error) 
 CreateLifecyclePolicy(ctx context.Context, params *CreateLifecyclePolicyInput, optFns ...func(*Options)) (*CreateLifecyclePolicyOutput, error) 
 CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) 
 DeleteComponent(ctx context.Context, params *DeleteComponentInput, optFns ...func(*Options)) (*DeleteComponentOutput, error) 
 DeleteContainerRecipe(ctx context.Context, params *DeleteContainerRecipeInput, optFns ...func(*Options)) (*DeleteContainerRecipeOutput, error) 
 DeleteDistributionConfiguration(ctx context.Context, params *DeleteDistributionConfigurationInput, optFns ...func(*Options)) (*DeleteDistributionConfigurationOutput, error) 
 DeleteImage(ctx context.Context, params *DeleteImageInput, optFns ...func(*Options)) (*DeleteImageOutput, error) 
 DeleteImagePipeline(ctx context.Context, params *DeleteImagePipelineInput, optFns ...func(*Options)) (*DeleteImagePipelineOutput, error) 
 DeleteImageRecipe(ctx context.Context, params *DeleteImageRecipeInput, optFns ...func(*Options)) (*DeleteImageRecipeOutput, error) 
 DeleteInfrastructureConfiguration(ctx context.Context, params *DeleteInfrastructureConfigurationInput, optFns ...func(*Options)) (*DeleteInfrastructureConfigurationOutput, error) 
 DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) 
 DeleteWorkflow(ctx context.Context, params *DeleteWorkflowInput, optFns ...func(*Options)) (*DeleteWorkflowOutput, error) 
 GetComponent(ctx context.Context, params *GetComponentInput, optFns ...func(*Options)) (*GetComponentOutput, error) 
 GetComponentPolicy(ctx context.Context, params *GetComponentPolicyInput, optFns ...func(*Options)) (*GetComponentPolicyOutput, error) 
 GetContainerRecipe(ctx context.Context, params *GetContainerRecipeInput, optFns ...func(*Options)) (*GetContainerRecipeOutput, error) 
 GetContainerRecipePolicy(ctx context.Context, params *GetContainerRecipePolicyInput, optFns ...func(*Options)) (*GetContainerRecipePolicyOutput, error) 
 GetDistributionConfiguration(ctx context.Context, params *GetDistributionConfigurationInput, optFns ...func(*Options)) (*GetDistributionConfigurationOutput, error) 
 GetImage(ctx context.Context, params *GetImageInput, optFns ...func(*Options)) (*GetImageOutput, error) 
 GetImagePipeline(ctx context.Context, params *GetImagePipelineInput, optFns ...func(*Options)) (*GetImagePipelineOutput, error) 
 GetImagePolicy(ctx context.Context, params *GetImagePolicyInput, optFns ...func(*Options)) (*GetImagePolicyOutput, error) 
 GetImageRecipe(ctx context.Context, params *GetImageRecipeInput, optFns ...func(*Options)) (*GetImageRecipeOutput, error) 
 GetImageRecipePolicy(ctx context.Context, params *GetImageRecipePolicyInput, optFns ...func(*Options)) (*GetImageRecipePolicyOutput, error) 
 GetInfrastructureConfiguration(ctx context.Context, params *GetInfrastructureConfigurationInput, optFns ...func(*Options)) (*GetInfrastructureConfigurationOutput, error) 
 GetLifecycleExecution(ctx context.Context, params *GetLifecycleExecutionInput, optFns ...func(*Options)) (*GetLifecycleExecutionOutput, error) 
 GetLifecyclePolicy(ctx context.Context, params *GetLifecyclePolicyInput, optFns ...func(*Options)) (*GetLifecyclePolicyOutput, error) 
 GetMarketplaceResource(ctx context.Context, params *GetMarketplaceResourceInput, optFns ...func(*Options)) (*GetMarketplaceResourceOutput, error) 
 GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) 
 GetWorkflowExecution(ctx context.Context, params *GetWorkflowExecutionInput, optFns ...func(*Options)) (*GetWorkflowExecutionOutput, error) 
 GetWorkflowStepExecution(ctx context.Context, params *GetWorkflowStepExecutionInput, optFns ...func(*Options)) (*GetWorkflowStepExecutionOutput, error) 
 ImportComponent(ctx context.Context, params *ImportComponentInput, optFns ...func(*Options)) (*ImportComponentOutput, error) 
 ImportDiskImage(ctx context.Context, params *ImportDiskImageInput, optFns ...func(*Options)) (*ImportDiskImageOutput, error) 
 ImportVmImage(ctx context.Context, params *ImportVmImageInput, optFns ...func(*Options)) (*ImportVmImageOutput, error) 
 ListComponentBuildVersions(ctx context.Context, params *ListComponentBuildVersionsInput, optFns ...func(*Options)) (*ListComponentBuildVersionsOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListContainerRecipes(ctx context.Context, params *ListContainerRecipesInput, optFns ...func(*Options)) (*ListContainerRecipesOutput, error) 
 ListDistributionConfigurations(ctx context.Context, params *ListDistributionConfigurationsInput, optFns ...func(*Options)) (*ListDistributionConfigurationsOutput, error) 
 ListImageBuildVersions(ctx context.Context, params *ListImageBuildVersionsInput, optFns ...func(*Options)) (*ListImageBuildVersionsOutput, error) 
 ListImagePackages(ctx context.Context, params *ListImagePackagesInput, optFns ...func(*Options)) (*ListImagePackagesOutput, error) 
 ListImagePipelineImages(ctx context.Context, params *ListImagePipelineImagesInput, optFns ...func(*Options)) (*ListImagePipelineImagesOutput, error) 
 ListImagePipelines(ctx context.Context, params *ListImagePipelinesInput, optFns ...func(*Options)) (*ListImagePipelinesOutput, error) 
 ListImageRecipes(ctx context.Context, params *ListImageRecipesInput, optFns ...func(*Options)) (*ListImageRecipesOutput, error) 
 ListImageScanFindingAggregations(ctx context.Context, params *ListImageScanFindingAggregationsInput, optFns ...func(*Options)) (*ListImageScanFindingAggregationsOutput, error) 
 ListImageScanFindings(ctx context.Context, params *ListImageScanFindingsInput, optFns ...func(*Options)) (*ListImageScanFindingsOutput, error) 
 ListImages(ctx context.Context, params *ListImagesInput, optFns ...func(*Options)) (*ListImagesOutput, error) 
 ListInfrastructureConfigurations(ctx context.Context, params *ListInfrastructureConfigurationsInput, optFns ...func(*Options)) (*ListInfrastructureConfigurationsOutput, error) 
 ListLifecycleExecutionResources(ctx context.Context, params *ListLifecycleExecutionResourcesInput, optFns ...func(*Options)) (*ListLifecycleExecutionResourcesOutput, error) 
 ListLifecycleExecutions(ctx context.Context, params *ListLifecycleExecutionsInput, optFns ...func(*Options)) (*ListLifecycleExecutionsOutput, error) 
 ListLifecyclePolicies(ctx context.Context, params *ListLifecyclePoliciesInput, optFns ...func(*Options)) (*ListLifecyclePoliciesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWaitingWorkflowSteps(ctx context.Context, params *ListWaitingWorkflowStepsInput, optFns ...func(*Options)) (*ListWaitingWorkflowStepsOutput, error) 
 ListWorkflowBuildVersions(ctx context.Context, params *ListWorkflowBuildVersionsInput, optFns ...func(*Options)) (*ListWorkflowBuildVersionsOutput, error) 
 ListWorkflowExecutions(ctx context.Context, params *ListWorkflowExecutionsInput, optFns ...func(*Options)) (*ListWorkflowExecutionsOutput, error) 
 ListWorkflowStepExecutions(ctx context.Context, params *ListWorkflowStepExecutionsInput, optFns ...func(*Options)) (*ListWorkflowStepExecutionsOutput, error) 
 ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) 
 PutComponentPolicy(ctx context.Context, params *PutComponentPolicyInput, optFns ...func(*Options)) (*PutComponentPolicyOutput, error) 
 PutContainerRecipePolicy(ctx context.Context, params *PutContainerRecipePolicyInput, optFns ...func(*Options)) (*PutContainerRecipePolicyOutput, error) 
 PutImagePolicy(ctx context.Context, params *PutImagePolicyInput, optFns ...func(*Options)) (*PutImagePolicyOutput, error) 
 PutImageRecipePolicy(ctx context.Context, params *PutImageRecipePolicyInput, optFns ...func(*Options)) (*PutImageRecipePolicyOutput, error) 
 SendWorkflowStepAction(ctx context.Context, params *SendWorkflowStepActionInput, optFns ...func(*Options)) (*SendWorkflowStepActionOutput, error) 
 StartImagePipelineExecution(ctx context.Context, params *StartImagePipelineExecutionInput, optFns ...func(*Options)) (*StartImagePipelineExecutionOutput, error) 
 StartResourceStateUpdate(ctx context.Context, params *StartResourceStateUpdateInput, optFns ...func(*Options)) (*StartResourceStateUpdateOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDistributionConfiguration(ctx context.Context, params *UpdateDistributionConfigurationInput, optFns ...func(*Options)) (*UpdateDistributionConfigurationOutput, error) 
 UpdateImagePipeline(ctx context.Context, params *UpdateImagePipelineInput, optFns ...func(*Options)) (*UpdateImagePipelineOutput, error) 
 UpdateInfrastructureConfiguration(ctx context.Context, params *UpdateInfrastructureConfigurationInput, optFns ...func(*Options)) (*UpdateInfrastructureConfigurationOutput, error) 
 UpdateLifecyclePolicy(ctx context.Context, params *UpdateLifecyclePolicyInput, optFns ...func(*Options)) (*UpdateLifecyclePolicyOutput, error) 
}
