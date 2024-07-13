package imagebuilder_test

// tests for the imagebuilder service interface for this ../../../service/imagebuilder/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/imagebuilder/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/imagebuilder/imagebuilder_iface"
	"github.com/stretchr/testify/assert"
)

func TestImagebuilderServiceCanBeMocked(t *testing.T) {
	var iface imagebuilder_iface.IClient
	iface = &imagebuilder.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := imagebuilder.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelImageCreation", func(t *testing.T) {
        input := &imagebuilder.CancelImageCreationInput{}
        output := &imagebuilder.CancelImageCreationOutput{}

        mockClient.On("CancelImageCreation", ctx, input).Return(output, nil)

        result, err := mockClient.CancelImageCreation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelLifecycleExecution", func(t *testing.T) {
        input := &imagebuilder.CancelLifecycleExecutionInput{}
        output := &imagebuilder.CancelLifecycleExecutionOutput{}

        mockClient.On("CancelLifecycleExecution", ctx, input).Return(output, nil)

        result, err := mockClient.CancelLifecycleExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponent", func(t *testing.T) {
        input := &imagebuilder.CreateComponentInput{}
        output := &imagebuilder.CreateComponentOutput{}

        mockClient.On("CreateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainerRecipe", func(t *testing.T) {
        input := &imagebuilder.CreateContainerRecipeInput{}
        output := &imagebuilder.CreateContainerRecipeOutput{}

        mockClient.On("CreateContainerRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainerRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDistributionConfiguration", func(t *testing.T) {
        input := &imagebuilder.CreateDistributionConfigurationInput{}
        output := &imagebuilder.CreateDistributionConfigurationOutput{}

        mockClient.On("CreateDistributionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDistributionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImage", func(t *testing.T) {
        input := &imagebuilder.CreateImageInput{}
        output := &imagebuilder.CreateImageOutput{}

        mockClient.On("CreateImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImagePipeline", func(t *testing.T) {
        input := &imagebuilder.CreateImagePipelineInput{}
        output := &imagebuilder.CreateImagePipelineOutput{}

        mockClient.On("CreateImagePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImagePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImageRecipe", func(t *testing.T) {
        input := &imagebuilder.CreateImageRecipeInput{}
        output := &imagebuilder.CreateImageRecipeOutput{}

        mockClient.On("CreateImageRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImageRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInfrastructureConfiguration", func(t *testing.T) {
        input := &imagebuilder.CreateInfrastructureConfigurationInput{}
        output := &imagebuilder.CreateInfrastructureConfigurationOutput{}

        mockClient.On("CreateInfrastructureConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInfrastructureConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLifecyclePolicy", func(t *testing.T) {
        input := &imagebuilder.CreateLifecyclePolicyInput{}
        output := &imagebuilder.CreateLifecyclePolicyOutput{}

        mockClient.On("CreateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflow", func(t *testing.T) {
        input := &imagebuilder.CreateWorkflowInput{}
        output := &imagebuilder.CreateWorkflowOutput{}

        mockClient.On("CreateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponent", func(t *testing.T) {
        input := &imagebuilder.DeleteComponentInput{}
        output := &imagebuilder.DeleteComponentOutput{}

        mockClient.On("DeleteComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainerRecipe", func(t *testing.T) {
        input := &imagebuilder.DeleteContainerRecipeInput{}
        output := &imagebuilder.DeleteContainerRecipeOutput{}

        mockClient.On("DeleteContainerRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainerRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDistributionConfiguration", func(t *testing.T) {
        input := &imagebuilder.DeleteDistributionConfigurationInput{}
        output := &imagebuilder.DeleteDistributionConfigurationOutput{}

        mockClient.On("DeleteDistributionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDistributionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImage", func(t *testing.T) {
        input := &imagebuilder.DeleteImageInput{}
        output := &imagebuilder.DeleteImageOutput{}

        mockClient.On("DeleteImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImagePipeline", func(t *testing.T) {
        input := &imagebuilder.DeleteImagePipelineInput{}
        output := &imagebuilder.DeleteImagePipelineOutput{}

        mockClient.On("DeleteImagePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImagePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImageRecipe", func(t *testing.T) {
        input := &imagebuilder.DeleteImageRecipeInput{}
        output := &imagebuilder.DeleteImageRecipeOutput{}

        mockClient.On("DeleteImageRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImageRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInfrastructureConfiguration", func(t *testing.T) {
        input := &imagebuilder.DeleteInfrastructureConfigurationInput{}
        output := &imagebuilder.DeleteInfrastructureConfigurationOutput{}

        mockClient.On("DeleteInfrastructureConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInfrastructureConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecyclePolicy", func(t *testing.T) {
        input := &imagebuilder.DeleteLifecyclePolicyInput{}
        output := &imagebuilder.DeleteLifecyclePolicyOutput{}

        mockClient.On("DeleteLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &imagebuilder.DeleteWorkflowInput{}
        output := &imagebuilder.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponent", func(t *testing.T) {
        input := &imagebuilder.GetComponentInput{}
        output := &imagebuilder.GetComponentOutput{}

        mockClient.On("GetComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponentPolicy", func(t *testing.T) {
        input := &imagebuilder.GetComponentPolicyInput{}
        output := &imagebuilder.GetComponentPolicyOutput{}

        mockClient.On("GetComponentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerRecipe", func(t *testing.T) {
        input := &imagebuilder.GetContainerRecipeInput{}
        output := &imagebuilder.GetContainerRecipeOutput{}

        mockClient.On("GetContainerRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerRecipePolicy", func(t *testing.T) {
        input := &imagebuilder.GetContainerRecipePolicyInput{}
        output := &imagebuilder.GetContainerRecipePolicyOutput{}

        mockClient.On("GetContainerRecipePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerRecipePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributionConfiguration", func(t *testing.T) {
        input := &imagebuilder.GetDistributionConfigurationInput{}
        output := &imagebuilder.GetDistributionConfigurationOutput{}

        mockClient.On("GetDistributionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImage", func(t *testing.T) {
        input := &imagebuilder.GetImageInput{}
        output := &imagebuilder.GetImageOutput{}

        mockClient.On("GetImage", ctx, input).Return(output, nil)

        result, err := mockClient.GetImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImagePipeline", func(t *testing.T) {
        input := &imagebuilder.GetImagePipelineInput{}
        output := &imagebuilder.GetImagePipelineOutput{}

        mockClient.On("GetImagePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.GetImagePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImagePolicy", func(t *testing.T) {
        input := &imagebuilder.GetImagePolicyInput{}
        output := &imagebuilder.GetImagePolicyOutput{}

        mockClient.On("GetImagePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetImagePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageRecipe", func(t *testing.T) {
        input := &imagebuilder.GetImageRecipeInput{}
        output := &imagebuilder.GetImageRecipeOutput{}

        mockClient.On("GetImageRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageRecipePolicy", func(t *testing.T) {
        input := &imagebuilder.GetImageRecipePolicyInput{}
        output := &imagebuilder.GetImageRecipePolicyOutput{}

        mockClient.On("GetImageRecipePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageRecipePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInfrastructureConfiguration", func(t *testing.T) {
        input := &imagebuilder.GetInfrastructureConfigurationInput{}
        output := &imagebuilder.GetInfrastructureConfigurationOutput{}

        mockClient.On("GetInfrastructureConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetInfrastructureConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecycleExecution", func(t *testing.T) {
        input := &imagebuilder.GetLifecycleExecutionInput{}
        output := &imagebuilder.GetLifecycleExecutionOutput{}

        mockClient.On("GetLifecycleExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecycleExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicy", func(t *testing.T) {
        input := &imagebuilder.GetLifecyclePolicyInput{}
        output := &imagebuilder.GetLifecyclePolicyOutput{}

        mockClient.On("GetLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &imagebuilder.GetWorkflowInput{}
        output := &imagebuilder.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowExecution", func(t *testing.T) {
        input := &imagebuilder.GetWorkflowExecutionInput{}
        output := &imagebuilder.GetWorkflowExecutionOutput{}

        mockClient.On("GetWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowStepExecution", func(t *testing.T) {
        input := &imagebuilder.GetWorkflowStepExecutionInput{}
        output := &imagebuilder.GetWorkflowStepExecutionOutput{}

        mockClient.On("GetWorkflowStepExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowStepExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportComponent", func(t *testing.T) {
        input := &imagebuilder.ImportComponentInput{}
        output := &imagebuilder.ImportComponentOutput{}

        mockClient.On("ImportComponent", ctx, input).Return(output, nil)

        result, err := mockClient.ImportComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportVmImage", func(t *testing.T) {
        input := &imagebuilder.ImportVmImageInput{}
        output := &imagebuilder.ImportVmImageOutput{}

        mockClient.On("ImportVmImage", ctx, input).Return(output, nil)

        result, err := mockClient.ImportVmImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponentBuildVersions", func(t *testing.T) {
        input := &imagebuilder.ListComponentBuildVersionsInput{}
        output := &imagebuilder.ListComponentBuildVersionsOutput{}

        mockClient.On("ListComponentBuildVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponentBuildVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &imagebuilder.ListComponentsInput{}
        output := &imagebuilder.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContainerRecipes", func(t *testing.T) {
        input := &imagebuilder.ListContainerRecipesInput{}
        output := &imagebuilder.ListContainerRecipesOutput{}

        mockClient.On("ListContainerRecipes", ctx, input).Return(output, nil)

        result, err := mockClient.ListContainerRecipes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionConfigurations", func(t *testing.T) {
        input := &imagebuilder.ListDistributionConfigurationsInput{}
        output := &imagebuilder.ListDistributionConfigurationsOutput{}

        mockClient.On("ListDistributionConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageBuildVersions", func(t *testing.T) {
        input := &imagebuilder.ListImageBuildVersionsInput{}
        output := &imagebuilder.ListImageBuildVersionsOutput{}

        mockClient.On("ListImageBuildVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageBuildVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImagePackages", func(t *testing.T) {
        input := &imagebuilder.ListImagePackagesInput{}
        output := &imagebuilder.ListImagePackagesOutput{}

        mockClient.On("ListImagePackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListImagePackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImagePipelineImages", func(t *testing.T) {
        input := &imagebuilder.ListImagePipelineImagesInput{}
        output := &imagebuilder.ListImagePipelineImagesOutput{}

        mockClient.On("ListImagePipelineImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListImagePipelineImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImagePipelines", func(t *testing.T) {
        input := &imagebuilder.ListImagePipelinesInput{}
        output := &imagebuilder.ListImagePipelinesOutput{}

        mockClient.On("ListImagePipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListImagePipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageRecipes", func(t *testing.T) {
        input := &imagebuilder.ListImageRecipesInput{}
        output := &imagebuilder.ListImageRecipesOutput{}

        mockClient.On("ListImageRecipes", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageRecipes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageScanFindingAggregations", func(t *testing.T) {
        input := &imagebuilder.ListImageScanFindingAggregationsInput{}
        output := &imagebuilder.ListImageScanFindingAggregationsOutput{}

        mockClient.On("ListImageScanFindingAggregations", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageScanFindingAggregations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageScanFindings", func(t *testing.T) {
        input := &imagebuilder.ListImageScanFindingsInput{}
        output := &imagebuilder.ListImageScanFindingsOutput{}

        mockClient.On("ListImageScanFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageScanFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImages", func(t *testing.T) {
        input := &imagebuilder.ListImagesInput{}
        output := &imagebuilder.ListImagesOutput{}

        mockClient.On("ListImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInfrastructureConfigurations", func(t *testing.T) {
        input := &imagebuilder.ListInfrastructureConfigurationsInput{}
        output := &imagebuilder.ListInfrastructureConfigurationsOutput{}

        mockClient.On("ListInfrastructureConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInfrastructureConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLifecycleExecutionResources", func(t *testing.T) {
        input := &imagebuilder.ListLifecycleExecutionResourcesInput{}
        output := &imagebuilder.ListLifecycleExecutionResourcesOutput{}

        mockClient.On("ListLifecycleExecutionResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListLifecycleExecutionResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLifecycleExecutions", func(t *testing.T) {
        input := &imagebuilder.ListLifecycleExecutionsInput{}
        output := &imagebuilder.ListLifecycleExecutionsOutput{}

        mockClient.On("ListLifecycleExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLifecycleExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLifecyclePolicies", func(t *testing.T) {
        input := &imagebuilder.ListLifecyclePoliciesInput{}
        output := &imagebuilder.ListLifecyclePoliciesOutput{}

        mockClient.On("ListLifecyclePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListLifecyclePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &imagebuilder.ListTagsForResourceInput{}
        output := &imagebuilder.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWaitingWorkflowSteps", func(t *testing.T) {
        input := &imagebuilder.ListWaitingWorkflowStepsInput{}
        output := &imagebuilder.ListWaitingWorkflowStepsOutput{}

        mockClient.On("ListWaitingWorkflowSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListWaitingWorkflowSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowBuildVersions", func(t *testing.T) {
        input := &imagebuilder.ListWorkflowBuildVersionsInput{}
        output := &imagebuilder.ListWorkflowBuildVersionsOutput{}

        mockClient.On("ListWorkflowBuildVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowBuildVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowExecutions", func(t *testing.T) {
        input := &imagebuilder.ListWorkflowExecutionsInput{}
        output := &imagebuilder.ListWorkflowExecutionsOutput{}

        mockClient.On("ListWorkflowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowStepExecutions", func(t *testing.T) {
        input := &imagebuilder.ListWorkflowStepExecutionsInput{}
        output := &imagebuilder.ListWorkflowStepExecutionsOutput{}

        mockClient.On("ListWorkflowStepExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowStepExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &imagebuilder.ListWorkflowsInput{}
        output := &imagebuilder.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutComponentPolicy", func(t *testing.T) {
        input := &imagebuilder.PutComponentPolicyInput{}
        output := &imagebuilder.PutComponentPolicyOutput{}

        mockClient.On("PutComponentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutComponentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutContainerRecipePolicy", func(t *testing.T) {
        input := &imagebuilder.PutContainerRecipePolicyInput{}
        output := &imagebuilder.PutContainerRecipePolicyOutput{}

        mockClient.On("PutContainerRecipePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutContainerRecipePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImagePolicy", func(t *testing.T) {
        input := &imagebuilder.PutImagePolicyInput{}
        output := &imagebuilder.PutImagePolicyOutput{}

        mockClient.On("PutImagePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutImagePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutImageRecipePolicy", func(t *testing.T) {
        input := &imagebuilder.PutImageRecipePolicyInput{}
        output := &imagebuilder.PutImageRecipePolicyOutput{}

        mockClient.On("PutImageRecipePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutImageRecipePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendWorkflowStepAction", func(t *testing.T) {
        input := &imagebuilder.SendWorkflowStepActionInput{}
        output := &imagebuilder.SendWorkflowStepActionOutput{}

        mockClient.On("SendWorkflowStepAction", ctx, input).Return(output, nil)

        result, err := mockClient.SendWorkflowStepAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImagePipelineExecution", func(t *testing.T) {
        input := &imagebuilder.StartImagePipelineExecutionInput{}
        output := &imagebuilder.StartImagePipelineExecutionOutput{}

        mockClient.On("StartImagePipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartImagePipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartResourceStateUpdate", func(t *testing.T) {
        input := &imagebuilder.StartResourceStateUpdateInput{}
        output := &imagebuilder.StartResourceStateUpdateOutput{}

        mockClient.On("StartResourceStateUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartResourceStateUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &imagebuilder.TagResourceInput{}
        output := &imagebuilder.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &imagebuilder.UntagResourceInput{}
        output := &imagebuilder.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDistributionConfiguration", func(t *testing.T) {
        input := &imagebuilder.UpdateDistributionConfigurationInput{}
        output := &imagebuilder.UpdateDistributionConfigurationOutput{}

        mockClient.On("UpdateDistributionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDistributionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImagePipeline", func(t *testing.T) {
        input := &imagebuilder.UpdateImagePipelineInput{}
        output := &imagebuilder.UpdateImagePipelineOutput{}

        mockClient.On("UpdateImagePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImagePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInfrastructureConfiguration", func(t *testing.T) {
        input := &imagebuilder.UpdateInfrastructureConfigurationInput{}
        output := &imagebuilder.UpdateInfrastructureConfigurationOutput{}

        mockClient.On("UpdateInfrastructureConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInfrastructureConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLifecyclePolicy", func(t *testing.T) {
        input := &imagebuilder.UpdateLifecyclePolicyInput{}
        output := &imagebuilder.UpdateLifecyclePolicyOutput{}

        mockClient.On("UpdateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
