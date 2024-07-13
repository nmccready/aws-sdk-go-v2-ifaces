package sagemaker_test

// tests for the sagemaker service interface for this ../../../service/sagemaker/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemaker/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemaker/sagemaker_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakerServiceCanBeMocked(t *testing.T) {
	var iface sagemaker_iface.IClient
	iface = &sagemaker.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemaker.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddAssociation", func(t *testing.T) {
        input := &sagemaker.AddAssociationInput{}
        output := &sagemaker.AddAssociationOutput{}

        mockClient.On("AddAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.AddAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &sagemaker.AddTagsInput{}
        output := &sagemaker.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTrialComponent", func(t *testing.T) {
        input := &sagemaker.AssociateTrialComponentInput{}
        output := &sagemaker.AssociateTrialComponentOutput{}

        mockClient.On("AssociateTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDescribeModelPackage", func(t *testing.T) {
        input := &sagemaker.BatchDescribeModelPackageInput{}
        output := &sagemaker.BatchDescribeModelPackageOutput{}

        mockClient.On("BatchDescribeModelPackage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDescribeModelPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAction", func(t *testing.T) {
        input := &sagemaker.CreateActionInput{}
        output := &sagemaker.CreateActionOutput{}

        mockClient.On("CreateAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlgorithm", func(t *testing.T) {
        input := &sagemaker.CreateAlgorithmInput{}
        output := &sagemaker.CreateAlgorithmOutput{}

        mockClient.On("CreateAlgorithm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlgorithm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &sagemaker.CreateAppInput{}
        output := &sagemaker.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppImageConfig", func(t *testing.T) {
        input := &sagemaker.CreateAppImageConfigInput{}
        output := &sagemaker.CreateAppImageConfigOutput{}

        mockClient.On("CreateAppImageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppImageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateArtifact", func(t *testing.T) {
        input := &sagemaker.CreateArtifactInput{}
        output := &sagemaker.CreateArtifactOutput{}

        mockClient.On("CreateArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.CreateArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutoMLJob", func(t *testing.T) {
        input := &sagemaker.CreateAutoMLJobInput{}
        output := &sagemaker.CreateAutoMLJobOutput{}

        mockClient.On("CreateAutoMLJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutoMLJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutoMLJobV2", func(t *testing.T) {
        input := &sagemaker.CreateAutoMLJobV2Input{}
        output := &sagemaker.CreateAutoMLJobV2Output{}

        mockClient.On("CreateAutoMLJobV2", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutoMLJobV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &sagemaker.CreateClusterInput{}
        output := &sagemaker.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCodeRepository", func(t *testing.T) {
        input := &sagemaker.CreateCodeRepositoryInput{}
        output := &sagemaker.CreateCodeRepositoryOutput{}

        mockClient.On("CreateCodeRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCodeRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCompilationJob", func(t *testing.T) {
        input := &sagemaker.CreateCompilationJobInput{}
        output := &sagemaker.CreateCompilationJobOutput{}

        mockClient.On("CreateCompilationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCompilationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContext", func(t *testing.T) {
        input := &sagemaker.CreateContextInput{}
        output := &sagemaker.CreateContextOutput{}

        mockClient.On("CreateContext", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.CreateDataQualityJobDefinitionInput{}
        output := &sagemaker.CreateDataQualityJobDefinitionOutput{}

        mockClient.On("CreateDataQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeviceFleet", func(t *testing.T) {
        input := &sagemaker.CreateDeviceFleetInput{}
        output := &sagemaker.CreateDeviceFleetOutput{}

        mockClient.On("CreateDeviceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeviceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &sagemaker.CreateDomainInput{}
        output := &sagemaker.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEdgeDeploymentPlan", func(t *testing.T) {
        input := &sagemaker.CreateEdgeDeploymentPlanInput{}
        output := &sagemaker.CreateEdgeDeploymentPlanOutput{}

        mockClient.On("CreateEdgeDeploymentPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEdgeDeploymentPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEdgeDeploymentStage", func(t *testing.T) {
        input := &sagemaker.CreateEdgeDeploymentStageInput{}
        output := &sagemaker.CreateEdgeDeploymentStageOutput{}

        mockClient.On("CreateEdgeDeploymentStage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEdgeDeploymentStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEdgePackagingJob", func(t *testing.T) {
        input := &sagemaker.CreateEdgePackagingJobInput{}
        output := &sagemaker.CreateEdgePackagingJobOutput{}

        mockClient.On("CreateEdgePackagingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEdgePackagingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpoint", func(t *testing.T) {
        input := &sagemaker.CreateEndpointInput{}
        output := &sagemaker.CreateEndpointOutput{}

        mockClient.On("CreateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpointConfig", func(t *testing.T) {
        input := &sagemaker.CreateEndpointConfigInput{}
        output := &sagemaker.CreateEndpointConfigOutput{}

        mockClient.On("CreateEndpointConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpointConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExperiment", func(t *testing.T) {
        input := &sagemaker.CreateExperimentInput{}
        output := &sagemaker.CreateExperimentOutput{}

        mockClient.On("CreateExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFeatureGroup", func(t *testing.T) {
        input := &sagemaker.CreateFeatureGroupInput{}
        output := &sagemaker.CreateFeatureGroupOutput{}

        mockClient.On("CreateFeatureGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFeatureGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlowDefinition", func(t *testing.T) {
        input := &sagemaker.CreateFlowDefinitionInput{}
        output := &sagemaker.CreateFlowDefinitionOutput{}

        mockClient.On("CreateFlowDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlowDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHub", func(t *testing.T) {
        input := &sagemaker.CreateHubInput{}
        output := &sagemaker.CreateHubOutput{}

        mockClient.On("CreateHub", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHubContentReference", func(t *testing.T) {
        input := &sagemaker.CreateHubContentReferenceInput{}
        output := &sagemaker.CreateHubContentReferenceOutput{}

        mockClient.On("CreateHubContentReference", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHubContentReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHumanTaskUi", func(t *testing.T) {
        input := &sagemaker.CreateHumanTaskUiInput{}
        output := &sagemaker.CreateHumanTaskUiOutput{}

        mockClient.On("CreateHumanTaskUi", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHumanTaskUi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHyperParameterTuningJob", func(t *testing.T) {
        input := &sagemaker.CreateHyperParameterTuningJobInput{}
        output := &sagemaker.CreateHyperParameterTuningJobOutput{}

        mockClient.On("CreateHyperParameterTuningJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHyperParameterTuningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImage", func(t *testing.T) {
        input := &sagemaker.CreateImageInput{}
        output := &sagemaker.CreateImageOutput{}

        mockClient.On("CreateImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImageVersion", func(t *testing.T) {
        input := &sagemaker.CreateImageVersionInput{}
        output := &sagemaker.CreateImageVersionOutput{}

        mockClient.On("CreateImageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInferenceComponent", func(t *testing.T) {
        input := &sagemaker.CreateInferenceComponentInput{}
        output := &sagemaker.CreateInferenceComponentOutput{}

        mockClient.On("CreateInferenceComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInferenceComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.CreateInferenceExperimentInput{}
        output := &sagemaker.CreateInferenceExperimentOutput{}

        mockClient.On("CreateInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInferenceRecommendationsJob", func(t *testing.T) {
        input := &sagemaker.CreateInferenceRecommendationsJobInput{}
        output := &sagemaker.CreateInferenceRecommendationsJobOutput{}

        mockClient.On("CreateInferenceRecommendationsJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInferenceRecommendationsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLabelingJob", func(t *testing.T) {
        input := &sagemaker.CreateLabelingJobInput{}
        output := &sagemaker.CreateLabelingJobOutput{}

        mockClient.On("CreateLabelingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLabelingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.CreateMlflowTrackingServerInput{}
        output := &sagemaker.CreateMlflowTrackingServerOutput{}

        mockClient.On("CreateMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &sagemaker.CreateModelInput{}
        output := &sagemaker.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelBiasJobDefinition", func(t *testing.T) {
        input := &sagemaker.CreateModelBiasJobDefinitionInput{}
        output := &sagemaker.CreateModelBiasJobDefinitionOutput{}

        mockClient.On("CreateModelBiasJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelBiasJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelCard", func(t *testing.T) {
        input := &sagemaker.CreateModelCardInput{}
        output := &sagemaker.CreateModelCardOutput{}

        mockClient.On("CreateModelCard", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelCard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelCardExportJob", func(t *testing.T) {
        input := &sagemaker.CreateModelCardExportJobInput{}
        output := &sagemaker.CreateModelCardExportJobOutput{}

        mockClient.On("CreateModelCardExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelCardExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelExplainabilityJobDefinition", func(t *testing.T) {
        input := &sagemaker.CreateModelExplainabilityJobDefinitionInput{}
        output := &sagemaker.CreateModelExplainabilityJobDefinitionOutput{}

        mockClient.On("CreateModelExplainabilityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelExplainabilityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelPackage", func(t *testing.T) {
        input := &sagemaker.CreateModelPackageInput{}
        output := &sagemaker.CreateModelPackageOutput{}

        mockClient.On("CreateModelPackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelPackageGroup", func(t *testing.T) {
        input := &sagemaker.CreateModelPackageGroupInput{}
        output := &sagemaker.CreateModelPackageGroupOutput{}

        mockClient.On("CreateModelPackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelPackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.CreateModelQualityJobDefinitionInput{}
        output := &sagemaker.CreateModelQualityJobDefinitionOutput{}

        mockClient.On("CreateModelQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.CreateMonitoringScheduleInput{}
        output := &sagemaker.CreateMonitoringScheduleOutput{}

        mockClient.On("CreateMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotebookInstance", func(t *testing.T) {
        input := &sagemaker.CreateNotebookInstanceInput{}
        output := &sagemaker.CreateNotebookInstanceOutput{}

        mockClient.On("CreateNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotebookInstanceLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.CreateNotebookInstanceLifecycleConfigInput{}
        output := &sagemaker.CreateNotebookInstanceLifecycleConfigOutput{}

        mockClient.On("CreateNotebookInstanceLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotebookInstanceLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOptimizationJob", func(t *testing.T) {
        input := &sagemaker.CreateOptimizationJobInput{}
        output := &sagemaker.CreateOptimizationJobOutput{}

        mockClient.On("CreateOptimizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOptimizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipeline", func(t *testing.T) {
        input := &sagemaker.CreatePipelineInput{}
        output := &sagemaker.CreatePipelineOutput{}

        mockClient.On("CreatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePresignedDomainUrl", func(t *testing.T) {
        input := &sagemaker.CreatePresignedDomainUrlInput{}
        output := &sagemaker.CreatePresignedDomainUrlOutput{}

        mockClient.On("CreatePresignedDomainUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePresignedDomainUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePresignedMlflowTrackingServerUrl", func(t *testing.T) {
        input := &sagemaker.CreatePresignedMlflowTrackingServerUrlInput{}
        output := &sagemaker.CreatePresignedMlflowTrackingServerUrlOutput{}

        mockClient.On("CreatePresignedMlflowTrackingServerUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePresignedMlflowTrackingServerUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePresignedNotebookInstanceUrl", func(t *testing.T) {
        input := &sagemaker.CreatePresignedNotebookInstanceUrlInput{}
        output := &sagemaker.CreatePresignedNotebookInstanceUrlOutput{}

        mockClient.On("CreatePresignedNotebookInstanceUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePresignedNotebookInstanceUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProcessingJob", func(t *testing.T) {
        input := &sagemaker.CreateProcessingJobInput{}
        output := &sagemaker.CreateProcessingJobOutput{}

        mockClient.On("CreateProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &sagemaker.CreateProjectInput{}
        output := &sagemaker.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSpace", func(t *testing.T) {
        input := &sagemaker.CreateSpaceInput{}
        output := &sagemaker.CreateSpaceOutput{}

        mockClient.On("CreateSpace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStudioLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.CreateStudioLifecycleConfigInput{}
        output := &sagemaker.CreateStudioLifecycleConfigOutput{}

        mockClient.On("CreateStudioLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStudioLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrainingJob", func(t *testing.T) {
        input := &sagemaker.CreateTrainingJobInput{}
        output := &sagemaker.CreateTrainingJobOutput{}

        mockClient.On("CreateTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransformJob", func(t *testing.T) {
        input := &sagemaker.CreateTransformJobInput{}
        output := &sagemaker.CreateTransformJobOutput{}

        mockClient.On("CreateTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrial", func(t *testing.T) {
        input := &sagemaker.CreateTrialInput{}
        output := &sagemaker.CreateTrialOutput{}

        mockClient.On("CreateTrial", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrialComponent", func(t *testing.T) {
        input := &sagemaker.CreateTrialComponentInput{}
        output := &sagemaker.CreateTrialComponentOutput{}

        mockClient.On("CreateTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserProfile", func(t *testing.T) {
        input := &sagemaker.CreateUserProfileInput{}
        output := &sagemaker.CreateUserProfileOutput{}

        mockClient.On("CreateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkforce", func(t *testing.T) {
        input := &sagemaker.CreateWorkforceInput{}
        output := &sagemaker.CreateWorkforceOutput{}

        mockClient.On("CreateWorkforce", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkforce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkteam", func(t *testing.T) {
        input := &sagemaker.CreateWorkteamInput{}
        output := &sagemaker.CreateWorkteamOutput{}

        mockClient.On("CreateWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAction", func(t *testing.T) {
        input := &sagemaker.DeleteActionInput{}
        output := &sagemaker.DeleteActionOutput{}

        mockClient.On("DeleteAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlgorithm", func(t *testing.T) {
        input := &sagemaker.DeleteAlgorithmInput{}
        output := &sagemaker.DeleteAlgorithmOutput{}

        mockClient.On("DeleteAlgorithm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlgorithm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &sagemaker.DeleteAppInput{}
        output := &sagemaker.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppImageConfig", func(t *testing.T) {
        input := &sagemaker.DeleteAppImageConfigInput{}
        output := &sagemaker.DeleteAppImageConfigOutput{}

        mockClient.On("DeleteAppImageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppImageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArtifact", func(t *testing.T) {
        input := &sagemaker.DeleteArtifactInput{}
        output := &sagemaker.DeleteArtifactOutput{}

        mockClient.On("DeleteArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssociation", func(t *testing.T) {
        input := &sagemaker.DeleteAssociationInput{}
        output := &sagemaker.DeleteAssociationOutput{}

        mockClient.On("DeleteAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &sagemaker.DeleteClusterInput{}
        output := &sagemaker.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCodeRepository", func(t *testing.T) {
        input := &sagemaker.DeleteCodeRepositoryInput{}
        output := &sagemaker.DeleteCodeRepositoryOutput{}

        mockClient.On("DeleteCodeRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCodeRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCompilationJob", func(t *testing.T) {
        input := &sagemaker.DeleteCompilationJobInput{}
        output := &sagemaker.DeleteCompilationJobOutput{}

        mockClient.On("DeleteCompilationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCompilationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContext", func(t *testing.T) {
        input := &sagemaker.DeleteContextInput{}
        output := &sagemaker.DeleteContextOutput{}

        mockClient.On("DeleteContext", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DeleteDataQualityJobDefinitionInput{}
        output := &sagemaker.DeleteDataQualityJobDefinitionOutput{}

        mockClient.On("DeleteDataQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeviceFleet", func(t *testing.T) {
        input := &sagemaker.DeleteDeviceFleetInput{}
        output := &sagemaker.DeleteDeviceFleetOutput{}

        mockClient.On("DeleteDeviceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeviceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &sagemaker.DeleteDomainInput{}
        output := &sagemaker.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEdgeDeploymentPlan", func(t *testing.T) {
        input := &sagemaker.DeleteEdgeDeploymentPlanInput{}
        output := &sagemaker.DeleteEdgeDeploymentPlanOutput{}

        mockClient.On("DeleteEdgeDeploymentPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEdgeDeploymentPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEdgeDeploymentStage", func(t *testing.T) {
        input := &sagemaker.DeleteEdgeDeploymentStageInput{}
        output := &sagemaker.DeleteEdgeDeploymentStageOutput{}

        mockClient.On("DeleteEdgeDeploymentStage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEdgeDeploymentStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &sagemaker.DeleteEndpointInput{}
        output := &sagemaker.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpointConfig", func(t *testing.T) {
        input := &sagemaker.DeleteEndpointConfigInput{}
        output := &sagemaker.DeleteEndpointConfigOutput{}

        mockClient.On("DeleteEndpointConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpointConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExperiment", func(t *testing.T) {
        input := &sagemaker.DeleteExperimentInput{}
        output := &sagemaker.DeleteExperimentOutput{}

        mockClient.On("DeleteExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFeatureGroup", func(t *testing.T) {
        input := &sagemaker.DeleteFeatureGroupInput{}
        output := &sagemaker.DeleteFeatureGroupOutput{}

        mockClient.On("DeleteFeatureGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFeatureGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlowDefinition", func(t *testing.T) {
        input := &sagemaker.DeleteFlowDefinitionInput{}
        output := &sagemaker.DeleteFlowDefinitionOutput{}

        mockClient.On("DeleteFlowDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlowDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHub", func(t *testing.T) {
        input := &sagemaker.DeleteHubInput{}
        output := &sagemaker.DeleteHubOutput{}

        mockClient.On("DeleteHub", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHubContent", func(t *testing.T) {
        input := &sagemaker.DeleteHubContentInput{}
        output := &sagemaker.DeleteHubContentOutput{}

        mockClient.On("DeleteHubContent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHubContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHubContentReference", func(t *testing.T) {
        input := &sagemaker.DeleteHubContentReferenceInput{}
        output := &sagemaker.DeleteHubContentReferenceOutput{}

        mockClient.On("DeleteHubContentReference", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHubContentReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHumanTaskUi", func(t *testing.T) {
        input := &sagemaker.DeleteHumanTaskUiInput{}
        output := &sagemaker.DeleteHumanTaskUiOutput{}

        mockClient.On("DeleteHumanTaskUi", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHumanTaskUi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHyperParameterTuningJob", func(t *testing.T) {
        input := &sagemaker.DeleteHyperParameterTuningJobInput{}
        output := &sagemaker.DeleteHyperParameterTuningJobOutput{}

        mockClient.On("DeleteHyperParameterTuningJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHyperParameterTuningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImage", func(t *testing.T) {
        input := &sagemaker.DeleteImageInput{}
        output := &sagemaker.DeleteImageOutput{}

        mockClient.On("DeleteImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImageVersion", func(t *testing.T) {
        input := &sagemaker.DeleteImageVersionInput{}
        output := &sagemaker.DeleteImageVersionOutput{}

        mockClient.On("DeleteImageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInferenceComponent", func(t *testing.T) {
        input := &sagemaker.DeleteInferenceComponentInput{}
        output := &sagemaker.DeleteInferenceComponentOutput{}

        mockClient.On("DeleteInferenceComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInferenceComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.DeleteInferenceExperimentInput{}
        output := &sagemaker.DeleteInferenceExperimentOutput{}

        mockClient.On("DeleteInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.DeleteMlflowTrackingServerInput{}
        output := &sagemaker.DeleteMlflowTrackingServerOutput{}

        mockClient.On("DeleteMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &sagemaker.DeleteModelInput{}
        output := &sagemaker.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelBiasJobDefinition", func(t *testing.T) {
        input := &sagemaker.DeleteModelBiasJobDefinitionInput{}
        output := &sagemaker.DeleteModelBiasJobDefinitionOutput{}

        mockClient.On("DeleteModelBiasJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelBiasJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelCard", func(t *testing.T) {
        input := &sagemaker.DeleteModelCardInput{}
        output := &sagemaker.DeleteModelCardOutput{}

        mockClient.On("DeleteModelCard", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelCard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelExplainabilityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DeleteModelExplainabilityJobDefinitionInput{}
        output := &sagemaker.DeleteModelExplainabilityJobDefinitionOutput{}

        mockClient.On("DeleteModelExplainabilityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelExplainabilityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelPackage", func(t *testing.T) {
        input := &sagemaker.DeleteModelPackageInput{}
        output := &sagemaker.DeleteModelPackageOutput{}

        mockClient.On("DeleteModelPackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelPackageGroup", func(t *testing.T) {
        input := &sagemaker.DeleteModelPackageGroupInput{}
        output := &sagemaker.DeleteModelPackageGroupOutput{}

        mockClient.On("DeleteModelPackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelPackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelPackageGroupPolicy", func(t *testing.T) {
        input := &sagemaker.DeleteModelPackageGroupPolicyInput{}
        output := &sagemaker.DeleteModelPackageGroupPolicyOutput{}

        mockClient.On("DeleteModelPackageGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelPackageGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DeleteModelQualityJobDefinitionInput{}
        output := &sagemaker.DeleteModelQualityJobDefinitionOutput{}

        mockClient.On("DeleteModelQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.DeleteMonitoringScheduleInput{}
        output := &sagemaker.DeleteMonitoringScheduleOutput{}

        mockClient.On("DeleteMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotebookInstance", func(t *testing.T) {
        input := &sagemaker.DeleteNotebookInstanceInput{}
        output := &sagemaker.DeleteNotebookInstanceOutput{}

        mockClient.On("DeleteNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotebookInstanceLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.DeleteNotebookInstanceLifecycleConfigInput{}
        output := &sagemaker.DeleteNotebookInstanceLifecycleConfigOutput{}

        mockClient.On("DeleteNotebookInstanceLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotebookInstanceLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOptimizationJob", func(t *testing.T) {
        input := &sagemaker.DeleteOptimizationJobInput{}
        output := &sagemaker.DeleteOptimizationJobOutput{}

        mockClient.On("DeleteOptimizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOptimizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipeline", func(t *testing.T) {
        input := &sagemaker.DeletePipelineInput{}
        output := &sagemaker.DeletePipelineOutput{}

        mockClient.On("DeletePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &sagemaker.DeleteProjectInput{}
        output := &sagemaker.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSpace", func(t *testing.T) {
        input := &sagemaker.DeleteSpaceInput{}
        output := &sagemaker.DeleteSpaceOutput{}

        mockClient.On("DeleteSpace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudioLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.DeleteStudioLifecycleConfigInput{}
        output := &sagemaker.DeleteStudioLifecycleConfigOutput{}

        mockClient.On("DeleteStudioLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudioLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &sagemaker.DeleteTagsInput{}
        output := &sagemaker.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrial", func(t *testing.T) {
        input := &sagemaker.DeleteTrialInput{}
        output := &sagemaker.DeleteTrialOutput{}

        mockClient.On("DeleteTrial", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrialComponent", func(t *testing.T) {
        input := &sagemaker.DeleteTrialComponentInput{}
        output := &sagemaker.DeleteTrialComponentOutput{}

        mockClient.On("DeleteTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserProfile", func(t *testing.T) {
        input := &sagemaker.DeleteUserProfileInput{}
        output := &sagemaker.DeleteUserProfileOutput{}

        mockClient.On("DeleteUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkforce", func(t *testing.T) {
        input := &sagemaker.DeleteWorkforceInput{}
        output := &sagemaker.DeleteWorkforceOutput{}

        mockClient.On("DeleteWorkforce", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkforce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkteam", func(t *testing.T) {
        input := &sagemaker.DeleteWorkteamInput{}
        output := &sagemaker.DeleteWorkteamOutput{}

        mockClient.On("DeleteWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterDevices", func(t *testing.T) {
        input := &sagemaker.DeregisterDevicesInput{}
        output := &sagemaker.DeregisterDevicesOutput{}

        mockClient.On("DeregisterDevices", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAction", func(t *testing.T) {
        input := &sagemaker.DescribeActionInput{}
        output := &sagemaker.DescribeActionOutput{}

        mockClient.On("DescribeAction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlgorithm", func(t *testing.T) {
        input := &sagemaker.DescribeAlgorithmInput{}
        output := &sagemaker.DescribeAlgorithmOutput{}

        mockClient.On("DescribeAlgorithm", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlgorithm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApp", func(t *testing.T) {
        input := &sagemaker.DescribeAppInput{}
        output := &sagemaker.DescribeAppOutput{}

        mockClient.On("DescribeApp", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppImageConfig", func(t *testing.T) {
        input := &sagemaker.DescribeAppImageConfigInput{}
        output := &sagemaker.DescribeAppImageConfigOutput{}

        mockClient.On("DescribeAppImageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppImageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeArtifact", func(t *testing.T) {
        input := &sagemaker.DescribeArtifactInput{}
        output := &sagemaker.DescribeArtifactOutput{}

        mockClient.On("DescribeArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoMLJob", func(t *testing.T) {
        input := &sagemaker.DescribeAutoMLJobInput{}
        output := &sagemaker.DescribeAutoMLJobOutput{}

        mockClient.On("DescribeAutoMLJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoMLJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoMLJobV2", func(t *testing.T) {
        input := &sagemaker.DescribeAutoMLJobV2Input{}
        output := &sagemaker.DescribeAutoMLJobV2Output{}

        mockClient.On("DescribeAutoMLJobV2", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoMLJobV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &sagemaker.DescribeClusterInput{}
        output := &sagemaker.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterNode", func(t *testing.T) {
        input := &sagemaker.DescribeClusterNodeInput{}
        output := &sagemaker.DescribeClusterNodeOutput{}

        mockClient.On("DescribeClusterNode", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCodeRepository", func(t *testing.T) {
        input := &sagemaker.DescribeCodeRepositoryInput{}
        output := &sagemaker.DescribeCodeRepositoryOutput{}

        mockClient.On("DescribeCodeRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCodeRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCompilationJob", func(t *testing.T) {
        input := &sagemaker.DescribeCompilationJobInput{}
        output := &sagemaker.DescribeCompilationJobOutput{}

        mockClient.On("DescribeCompilationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCompilationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContext", func(t *testing.T) {
        input := &sagemaker.DescribeContextInput{}
        output := &sagemaker.DescribeContextOutput{}

        mockClient.On("DescribeContext", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DescribeDataQualityJobDefinitionInput{}
        output := &sagemaker.DescribeDataQualityJobDefinitionOutput{}

        mockClient.On("DescribeDataQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevice", func(t *testing.T) {
        input := &sagemaker.DescribeDeviceInput{}
        output := &sagemaker.DescribeDeviceOutput{}

        mockClient.On("DescribeDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeviceFleet", func(t *testing.T) {
        input := &sagemaker.DescribeDeviceFleetInput{}
        output := &sagemaker.DescribeDeviceFleetOutput{}

        mockClient.On("DescribeDeviceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeviceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &sagemaker.DescribeDomainInput{}
        output := &sagemaker.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEdgeDeploymentPlan", func(t *testing.T) {
        input := &sagemaker.DescribeEdgeDeploymentPlanInput{}
        output := &sagemaker.DescribeEdgeDeploymentPlanOutput{}

        mockClient.On("DescribeEdgeDeploymentPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEdgeDeploymentPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEdgePackagingJob", func(t *testing.T) {
        input := &sagemaker.DescribeEdgePackagingJobInput{}
        output := &sagemaker.DescribeEdgePackagingJobOutput{}

        mockClient.On("DescribeEdgePackagingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEdgePackagingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoint", func(t *testing.T) {
        input := &sagemaker.DescribeEndpointInput{}
        output := &sagemaker.DescribeEndpointOutput{}

        mockClient.On("DescribeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointConfig", func(t *testing.T) {
        input := &sagemaker.DescribeEndpointConfigInput{}
        output := &sagemaker.DescribeEndpointConfigOutput{}

        mockClient.On("DescribeEndpointConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExperiment", func(t *testing.T) {
        input := &sagemaker.DescribeExperimentInput{}
        output := &sagemaker.DescribeExperimentOutput{}

        mockClient.On("DescribeExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFeatureGroup", func(t *testing.T) {
        input := &sagemaker.DescribeFeatureGroupInput{}
        output := &sagemaker.DescribeFeatureGroupOutput{}

        mockClient.On("DescribeFeatureGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFeatureGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFeatureMetadata", func(t *testing.T) {
        input := &sagemaker.DescribeFeatureMetadataInput{}
        output := &sagemaker.DescribeFeatureMetadataOutput{}

        mockClient.On("DescribeFeatureMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFeatureMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlowDefinition", func(t *testing.T) {
        input := &sagemaker.DescribeFlowDefinitionInput{}
        output := &sagemaker.DescribeFlowDefinitionOutput{}

        mockClient.On("DescribeFlowDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlowDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHub", func(t *testing.T) {
        input := &sagemaker.DescribeHubInput{}
        output := &sagemaker.DescribeHubOutput{}

        mockClient.On("DescribeHub", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHubContent", func(t *testing.T) {
        input := &sagemaker.DescribeHubContentInput{}
        output := &sagemaker.DescribeHubContentOutput{}

        mockClient.On("DescribeHubContent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHubContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHumanTaskUi", func(t *testing.T) {
        input := &sagemaker.DescribeHumanTaskUiInput{}
        output := &sagemaker.DescribeHumanTaskUiOutput{}

        mockClient.On("DescribeHumanTaskUi", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHumanTaskUi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHyperParameterTuningJob", func(t *testing.T) {
        input := &sagemaker.DescribeHyperParameterTuningJobInput{}
        output := &sagemaker.DescribeHyperParameterTuningJobOutput{}

        mockClient.On("DescribeHyperParameterTuningJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHyperParameterTuningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImage", func(t *testing.T) {
        input := &sagemaker.DescribeImageInput{}
        output := &sagemaker.DescribeImageOutput{}

        mockClient.On("DescribeImage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageVersion", func(t *testing.T) {
        input := &sagemaker.DescribeImageVersionInput{}
        output := &sagemaker.DescribeImageVersionOutput{}

        mockClient.On("DescribeImageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInferenceComponent", func(t *testing.T) {
        input := &sagemaker.DescribeInferenceComponentInput{}
        output := &sagemaker.DescribeInferenceComponentOutput{}

        mockClient.On("DescribeInferenceComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInferenceComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.DescribeInferenceExperimentInput{}
        output := &sagemaker.DescribeInferenceExperimentOutput{}

        mockClient.On("DescribeInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInferenceRecommendationsJob", func(t *testing.T) {
        input := &sagemaker.DescribeInferenceRecommendationsJobInput{}
        output := &sagemaker.DescribeInferenceRecommendationsJobOutput{}

        mockClient.On("DescribeInferenceRecommendationsJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInferenceRecommendationsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLabelingJob", func(t *testing.T) {
        input := &sagemaker.DescribeLabelingJobInput{}
        output := &sagemaker.DescribeLabelingJobOutput{}

        mockClient.On("DescribeLabelingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLabelingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLineageGroup", func(t *testing.T) {
        input := &sagemaker.DescribeLineageGroupInput{}
        output := &sagemaker.DescribeLineageGroupOutput{}

        mockClient.On("DescribeLineageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLineageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.DescribeMlflowTrackingServerInput{}
        output := &sagemaker.DescribeMlflowTrackingServerOutput{}

        mockClient.On("DescribeMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModel", func(t *testing.T) {
        input := &sagemaker.DescribeModelInput{}
        output := &sagemaker.DescribeModelOutput{}

        mockClient.On("DescribeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelBiasJobDefinition", func(t *testing.T) {
        input := &sagemaker.DescribeModelBiasJobDefinitionInput{}
        output := &sagemaker.DescribeModelBiasJobDefinitionOutput{}

        mockClient.On("DescribeModelBiasJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelBiasJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelCard", func(t *testing.T) {
        input := &sagemaker.DescribeModelCardInput{}
        output := &sagemaker.DescribeModelCardOutput{}

        mockClient.On("DescribeModelCard", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelCard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelCardExportJob", func(t *testing.T) {
        input := &sagemaker.DescribeModelCardExportJobInput{}
        output := &sagemaker.DescribeModelCardExportJobOutput{}

        mockClient.On("DescribeModelCardExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelCardExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelExplainabilityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DescribeModelExplainabilityJobDefinitionInput{}
        output := &sagemaker.DescribeModelExplainabilityJobDefinitionOutput{}

        mockClient.On("DescribeModelExplainabilityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelExplainabilityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelPackage", func(t *testing.T) {
        input := &sagemaker.DescribeModelPackageInput{}
        output := &sagemaker.DescribeModelPackageOutput{}

        mockClient.On("DescribeModelPackage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelPackageGroup", func(t *testing.T) {
        input := &sagemaker.DescribeModelPackageGroupInput{}
        output := &sagemaker.DescribeModelPackageGroupOutput{}

        mockClient.On("DescribeModelPackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelPackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelQualityJobDefinition", func(t *testing.T) {
        input := &sagemaker.DescribeModelQualityJobDefinitionInput{}
        output := &sagemaker.DescribeModelQualityJobDefinitionOutput{}

        mockClient.On("DescribeModelQualityJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelQualityJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.DescribeMonitoringScheduleInput{}
        output := &sagemaker.DescribeMonitoringScheduleOutput{}

        mockClient.On("DescribeMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotebookInstance", func(t *testing.T) {
        input := &sagemaker.DescribeNotebookInstanceInput{}
        output := &sagemaker.DescribeNotebookInstanceOutput{}

        mockClient.On("DescribeNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotebookInstanceLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{}
        output := &sagemaker.DescribeNotebookInstanceLifecycleConfigOutput{}

        mockClient.On("DescribeNotebookInstanceLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotebookInstanceLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOptimizationJob", func(t *testing.T) {
        input := &sagemaker.DescribeOptimizationJobInput{}
        output := &sagemaker.DescribeOptimizationJobOutput{}

        mockClient.On("DescribeOptimizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOptimizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipeline", func(t *testing.T) {
        input := &sagemaker.DescribePipelineInput{}
        output := &sagemaker.DescribePipelineOutput{}

        mockClient.On("DescribePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipelineDefinitionForExecution", func(t *testing.T) {
        input := &sagemaker.DescribePipelineDefinitionForExecutionInput{}
        output := &sagemaker.DescribePipelineDefinitionForExecutionOutput{}

        mockClient.On("DescribePipelineDefinitionForExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipelineDefinitionForExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipelineExecution", func(t *testing.T) {
        input := &sagemaker.DescribePipelineExecutionInput{}
        output := &sagemaker.DescribePipelineExecutionOutput{}

        mockClient.On("DescribePipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProcessingJob", func(t *testing.T) {
        input := &sagemaker.DescribeProcessingJobInput{}
        output := &sagemaker.DescribeProcessingJobOutput{}

        mockClient.On("DescribeProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &sagemaker.DescribeProjectInput{}
        output := &sagemaker.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpace", func(t *testing.T) {
        input := &sagemaker.DescribeSpaceInput{}
        output := &sagemaker.DescribeSpaceOutput{}

        mockClient.On("DescribeSpace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStudioLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.DescribeStudioLifecycleConfigInput{}
        output := &sagemaker.DescribeStudioLifecycleConfigOutput{}

        mockClient.On("DescribeStudioLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStudioLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubscribedWorkteam", func(t *testing.T) {
        input := &sagemaker.DescribeSubscribedWorkteamInput{}
        output := &sagemaker.DescribeSubscribedWorkteamOutput{}

        mockClient.On("DescribeSubscribedWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubscribedWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrainingJob", func(t *testing.T) {
        input := &sagemaker.DescribeTrainingJobInput{}
        output := &sagemaker.DescribeTrainingJobOutput{}

        mockClient.On("DescribeTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTransformJob", func(t *testing.T) {
        input := &sagemaker.DescribeTransformJobInput{}
        output := &sagemaker.DescribeTransformJobOutput{}

        mockClient.On("DescribeTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrial", func(t *testing.T) {
        input := &sagemaker.DescribeTrialInput{}
        output := &sagemaker.DescribeTrialOutput{}

        mockClient.On("DescribeTrial", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrialComponent", func(t *testing.T) {
        input := &sagemaker.DescribeTrialComponentInput{}
        output := &sagemaker.DescribeTrialComponentOutput{}

        mockClient.On("DescribeTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserProfile", func(t *testing.T) {
        input := &sagemaker.DescribeUserProfileInput{}
        output := &sagemaker.DescribeUserProfileOutput{}

        mockClient.On("DescribeUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkforce", func(t *testing.T) {
        input := &sagemaker.DescribeWorkforceInput{}
        output := &sagemaker.DescribeWorkforceOutput{}

        mockClient.On("DescribeWorkforce", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkforce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkteam", func(t *testing.T) {
        input := &sagemaker.DescribeWorkteamInput{}
        output := &sagemaker.DescribeWorkteamOutput{}

        mockClient.On("DescribeWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSagemakerServicecatalogPortfolio", func(t *testing.T) {
        input := &sagemaker.DisableSagemakerServicecatalogPortfolioInput{}
        output := &sagemaker.DisableSagemakerServicecatalogPortfolioOutput{}

        mockClient.On("DisableSagemakerServicecatalogPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSagemakerServicecatalogPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTrialComponent", func(t *testing.T) {
        input := &sagemaker.DisassociateTrialComponentInput{}
        output := &sagemaker.DisassociateTrialComponentOutput{}

        mockClient.On("DisassociateTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSagemakerServicecatalogPortfolio", func(t *testing.T) {
        input := &sagemaker.EnableSagemakerServicecatalogPortfolioInput{}
        output := &sagemaker.EnableSagemakerServicecatalogPortfolioOutput{}

        mockClient.On("EnableSagemakerServicecatalogPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSagemakerServicecatalogPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceFleetReport", func(t *testing.T) {
        input := &sagemaker.GetDeviceFleetReportInput{}
        output := &sagemaker.GetDeviceFleetReportOutput{}

        mockClient.On("GetDeviceFleetReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceFleetReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLineageGroupPolicy", func(t *testing.T) {
        input := &sagemaker.GetLineageGroupPolicyInput{}
        output := &sagemaker.GetLineageGroupPolicyOutput{}

        mockClient.On("GetLineageGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLineageGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelPackageGroupPolicy", func(t *testing.T) {
        input := &sagemaker.GetModelPackageGroupPolicyInput{}
        output := &sagemaker.GetModelPackageGroupPolicyOutput{}

        mockClient.On("GetModelPackageGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelPackageGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSagemakerServicecatalogPortfolioStatus", func(t *testing.T) {
        input := &sagemaker.GetSagemakerServicecatalogPortfolioStatusInput{}
        output := &sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput{}

        mockClient.On("GetSagemakerServicecatalogPortfolioStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetSagemakerServicecatalogPortfolioStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScalingConfigurationRecommendation", func(t *testing.T) {
        input := &sagemaker.GetScalingConfigurationRecommendationInput{}
        output := &sagemaker.GetScalingConfigurationRecommendationOutput{}

        mockClient.On("GetScalingConfigurationRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetScalingConfigurationRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSearchSuggestions", func(t *testing.T) {
        input := &sagemaker.GetSearchSuggestionsInput{}
        output := &sagemaker.GetSearchSuggestionsOutput{}

        mockClient.On("GetSearchSuggestions", ctx, input).Return(output, nil)

        result, err := mockClient.GetSearchSuggestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportHubContent", func(t *testing.T) {
        input := &sagemaker.ImportHubContentInput{}
        output := &sagemaker.ImportHubContentOutput{}

        mockClient.On("ImportHubContent", ctx, input).Return(output, nil)

        result, err := mockClient.ImportHubContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActions", func(t *testing.T) {
        input := &sagemaker.ListActionsInput{}
        output := &sagemaker.ListActionsOutput{}

        mockClient.On("ListActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlgorithms", func(t *testing.T) {
        input := &sagemaker.ListAlgorithmsInput{}
        output := &sagemaker.ListAlgorithmsOutput{}

        mockClient.On("ListAlgorithms", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlgorithms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &sagemaker.ListAliasesInput{}
        output := &sagemaker.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppImageConfigs", func(t *testing.T) {
        input := &sagemaker.ListAppImageConfigsInput{}
        output := &sagemaker.ListAppImageConfigsOutput{}

        mockClient.On("ListAppImageConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppImageConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApps", func(t *testing.T) {
        input := &sagemaker.ListAppsInput{}
        output := &sagemaker.ListAppsOutput{}

        mockClient.On("ListApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArtifacts", func(t *testing.T) {
        input := &sagemaker.ListArtifactsInput{}
        output := &sagemaker.ListArtifactsOutput{}

        mockClient.On("ListArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociations", func(t *testing.T) {
        input := &sagemaker.ListAssociationsInput{}
        output := &sagemaker.ListAssociationsOutput{}

        mockClient.On("ListAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutoMLJobs", func(t *testing.T) {
        input := &sagemaker.ListAutoMLJobsInput{}
        output := &sagemaker.ListAutoMLJobsOutput{}

        mockClient.On("ListAutoMLJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutoMLJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCandidatesForAutoMLJob", func(t *testing.T) {
        input := &sagemaker.ListCandidatesForAutoMLJobInput{}
        output := &sagemaker.ListCandidatesForAutoMLJobOutput{}

        mockClient.On("ListCandidatesForAutoMLJob", ctx, input).Return(output, nil)

        result, err := mockClient.ListCandidatesForAutoMLJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusterNodes", func(t *testing.T) {
        input := &sagemaker.ListClusterNodesInput{}
        output := &sagemaker.ListClusterNodesOutput{}

        mockClient.On("ListClusterNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusterNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &sagemaker.ListClustersInput{}
        output := &sagemaker.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCodeRepositories", func(t *testing.T) {
        input := &sagemaker.ListCodeRepositoriesInput{}
        output := &sagemaker.ListCodeRepositoriesOutput{}

        mockClient.On("ListCodeRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.ListCodeRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCompilationJobs", func(t *testing.T) {
        input := &sagemaker.ListCompilationJobsInput{}
        output := &sagemaker.ListCompilationJobsOutput{}

        mockClient.On("ListCompilationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListCompilationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContexts", func(t *testing.T) {
        input := &sagemaker.ListContextsInput{}
        output := &sagemaker.ListContextsOutput{}

        mockClient.On("ListContexts", ctx, input).Return(output, nil)

        result, err := mockClient.ListContexts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataQualityJobDefinitions", func(t *testing.T) {
        input := &sagemaker.ListDataQualityJobDefinitionsInput{}
        output := &sagemaker.ListDataQualityJobDefinitionsOutput{}

        mockClient.On("ListDataQualityJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataQualityJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceFleets", func(t *testing.T) {
        input := &sagemaker.ListDeviceFleetsInput{}
        output := &sagemaker.ListDeviceFleetsOutput{}

        mockClient.On("ListDeviceFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &sagemaker.ListDevicesInput{}
        output := &sagemaker.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &sagemaker.ListDomainsInput{}
        output := &sagemaker.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEdgeDeploymentPlans", func(t *testing.T) {
        input := &sagemaker.ListEdgeDeploymentPlansInput{}
        output := &sagemaker.ListEdgeDeploymentPlansOutput{}

        mockClient.On("ListEdgeDeploymentPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListEdgeDeploymentPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEdgePackagingJobs", func(t *testing.T) {
        input := &sagemaker.ListEdgePackagingJobsInput{}
        output := &sagemaker.ListEdgePackagingJobsOutput{}

        mockClient.On("ListEdgePackagingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEdgePackagingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpointConfigs", func(t *testing.T) {
        input := &sagemaker.ListEndpointConfigsInput{}
        output := &sagemaker.ListEndpointConfigsOutput{}

        mockClient.On("ListEndpointConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpointConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpoints", func(t *testing.T) {
        input := &sagemaker.ListEndpointsInput{}
        output := &sagemaker.ListEndpointsOutput{}

        mockClient.On("ListEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperiments", func(t *testing.T) {
        input := &sagemaker.ListExperimentsInput{}
        output := &sagemaker.ListExperimentsOutput{}

        mockClient.On("ListExperiments", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperiments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFeatureGroups", func(t *testing.T) {
        input := &sagemaker.ListFeatureGroupsInput{}
        output := &sagemaker.ListFeatureGroupsOutput{}

        mockClient.On("ListFeatureGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListFeatureGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowDefinitions", func(t *testing.T) {
        input := &sagemaker.ListFlowDefinitionsInput{}
        output := &sagemaker.ListFlowDefinitionsOutput{}

        mockClient.On("ListFlowDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHubContentVersions", func(t *testing.T) {
        input := &sagemaker.ListHubContentVersionsInput{}
        output := &sagemaker.ListHubContentVersionsOutput{}

        mockClient.On("ListHubContentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListHubContentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHubContents", func(t *testing.T) {
        input := &sagemaker.ListHubContentsInput{}
        output := &sagemaker.ListHubContentsOutput{}

        mockClient.On("ListHubContents", ctx, input).Return(output, nil)

        result, err := mockClient.ListHubContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHubs", func(t *testing.T) {
        input := &sagemaker.ListHubsInput{}
        output := &sagemaker.ListHubsOutput{}

        mockClient.On("ListHubs", ctx, input).Return(output, nil)

        result, err := mockClient.ListHubs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHumanTaskUis", func(t *testing.T) {
        input := &sagemaker.ListHumanTaskUisInput{}
        output := &sagemaker.ListHumanTaskUisOutput{}

        mockClient.On("ListHumanTaskUis", ctx, input).Return(output, nil)

        result, err := mockClient.ListHumanTaskUis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHyperParameterTuningJobs", func(t *testing.T) {
        input := &sagemaker.ListHyperParameterTuningJobsInput{}
        output := &sagemaker.ListHyperParameterTuningJobsOutput{}

        mockClient.On("ListHyperParameterTuningJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListHyperParameterTuningJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageVersions", func(t *testing.T) {
        input := &sagemaker.ListImageVersionsInput{}
        output := &sagemaker.ListImageVersionsOutput{}

        mockClient.On("ListImageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImages", func(t *testing.T) {
        input := &sagemaker.ListImagesInput{}
        output := &sagemaker.ListImagesOutput{}

        mockClient.On("ListImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceComponents", func(t *testing.T) {
        input := &sagemaker.ListInferenceComponentsInput{}
        output := &sagemaker.ListInferenceComponentsOutput{}

        mockClient.On("ListInferenceComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceExperiments", func(t *testing.T) {
        input := &sagemaker.ListInferenceExperimentsInput{}
        output := &sagemaker.ListInferenceExperimentsOutput{}

        mockClient.On("ListInferenceExperiments", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceExperiments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceRecommendationsJobSteps", func(t *testing.T) {
        input := &sagemaker.ListInferenceRecommendationsJobStepsInput{}
        output := &sagemaker.ListInferenceRecommendationsJobStepsOutput{}

        mockClient.On("ListInferenceRecommendationsJobSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceRecommendationsJobSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInferenceRecommendationsJobs", func(t *testing.T) {
        input := &sagemaker.ListInferenceRecommendationsJobsInput{}
        output := &sagemaker.ListInferenceRecommendationsJobsOutput{}

        mockClient.On("ListInferenceRecommendationsJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListInferenceRecommendationsJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLabelingJobs", func(t *testing.T) {
        input := &sagemaker.ListLabelingJobsInput{}
        output := &sagemaker.ListLabelingJobsOutput{}

        mockClient.On("ListLabelingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListLabelingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLabelingJobsForWorkteam", func(t *testing.T) {
        input := &sagemaker.ListLabelingJobsForWorkteamInput{}
        output := &sagemaker.ListLabelingJobsForWorkteamOutput{}

        mockClient.On("ListLabelingJobsForWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.ListLabelingJobsForWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLineageGroups", func(t *testing.T) {
        input := &sagemaker.ListLineageGroupsInput{}
        output := &sagemaker.ListLineageGroupsOutput{}

        mockClient.On("ListLineageGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListLineageGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMlflowTrackingServers", func(t *testing.T) {
        input := &sagemaker.ListMlflowTrackingServersInput{}
        output := &sagemaker.ListMlflowTrackingServersOutput{}

        mockClient.On("ListMlflowTrackingServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMlflowTrackingServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelBiasJobDefinitions", func(t *testing.T) {
        input := &sagemaker.ListModelBiasJobDefinitionsInput{}
        output := &sagemaker.ListModelBiasJobDefinitionsOutput{}

        mockClient.On("ListModelBiasJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelBiasJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelCardExportJobs", func(t *testing.T) {
        input := &sagemaker.ListModelCardExportJobsInput{}
        output := &sagemaker.ListModelCardExportJobsOutput{}

        mockClient.On("ListModelCardExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelCardExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelCardVersions", func(t *testing.T) {
        input := &sagemaker.ListModelCardVersionsInput{}
        output := &sagemaker.ListModelCardVersionsOutput{}

        mockClient.On("ListModelCardVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelCardVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelCards", func(t *testing.T) {
        input := &sagemaker.ListModelCardsInput{}
        output := &sagemaker.ListModelCardsOutput{}

        mockClient.On("ListModelCards", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelCards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelExplainabilityJobDefinitions", func(t *testing.T) {
        input := &sagemaker.ListModelExplainabilityJobDefinitionsInput{}
        output := &sagemaker.ListModelExplainabilityJobDefinitionsOutput{}

        mockClient.On("ListModelExplainabilityJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelExplainabilityJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelMetadata", func(t *testing.T) {
        input := &sagemaker.ListModelMetadataInput{}
        output := &sagemaker.ListModelMetadataOutput{}

        mockClient.On("ListModelMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelPackageGroups", func(t *testing.T) {
        input := &sagemaker.ListModelPackageGroupsInput{}
        output := &sagemaker.ListModelPackageGroupsOutput{}

        mockClient.On("ListModelPackageGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelPackageGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelPackages", func(t *testing.T) {
        input := &sagemaker.ListModelPackagesInput{}
        output := &sagemaker.ListModelPackagesOutput{}

        mockClient.On("ListModelPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelQualityJobDefinitions", func(t *testing.T) {
        input := &sagemaker.ListModelQualityJobDefinitionsInput{}
        output := &sagemaker.ListModelQualityJobDefinitionsOutput{}

        mockClient.On("ListModelQualityJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelQualityJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModels", func(t *testing.T) {
        input := &sagemaker.ListModelsInput{}
        output := &sagemaker.ListModelsOutput{}

        mockClient.On("ListModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitoringAlertHistory", func(t *testing.T) {
        input := &sagemaker.ListMonitoringAlertHistoryInput{}
        output := &sagemaker.ListMonitoringAlertHistoryOutput{}

        mockClient.On("ListMonitoringAlertHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitoringAlertHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitoringAlerts", func(t *testing.T) {
        input := &sagemaker.ListMonitoringAlertsInput{}
        output := &sagemaker.ListMonitoringAlertsOutput{}

        mockClient.On("ListMonitoringAlerts", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitoringAlerts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitoringExecutions", func(t *testing.T) {
        input := &sagemaker.ListMonitoringExecutionsInput{}
        output := &sagemaker.ListMonitoringExecutionsOutput{}

        mockClient.On("ListMonitoringExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitoringExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitoringSchedules", func(t *testing.T) {
        input := &sagemaker.ListMonitoringSchedulesInput{}
        output := &sagemaker.ListMonitoringSchedulesOutput{}

        mockClient.On("ListMonitoringSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitoringSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotebookInstanceLifecycleConfigs", func(t *testing.T) {
        input := &sagemaker.ListNotebookInstanceLifecycleConfigsInput{}
        output := &sagemaker.ListNotebookInstanceLifecycleConfigsOutput{}

        mockClient.On("ListNotebookInstanceLifecycleConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotebookInstanceLifecycleConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotebookInstances", func(t *testing.T) {
        input := &sagemaker.ListNotebookInstancesInput{}
        output := &sagemaker.ListNotebookInstancesOutput{}

        mockClient.On("ListNotebookInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotebookInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOptimizationJobs", func(t *testing.T) {
        input := &sagemaker.ListOptimizationJobsInput{}
        output := &sagemaker.ListOptimizationJobsOutput{}

        mockClient.On("ListOptimizationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListOptimizationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelineExecutionSteps", func(t *testing.T) {
        input := &sagemaker.ListPipelineExecutionStepsInput{}
        output := &sagemaker.ListPipelineExecutionStepsOutput{}

        mockClient.On("ListPipelineExecutionSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelineExecutionSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelineExecutions", func(t *testing.T) {
        input := &sagemaker.ListPipelineExecutionsInput{}
        output := &sagemaker.ListPipelineExecutionsOutput{}

        mockClient.On("ListPipelineExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelineExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelineParametersForExecution", func(t *testing.T) {
        input := &sagemaker.ListPipelineParametersForExecutionInput{}
        output := &sagemaker.ListPipelineParametersForExecutionOutput{}

        mockClient.On("ListPipelineParametersForExecution", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelineParametersForExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipelines", func(t *testing.T) {
        input := &sagemaker.ListPipelinesInput{}
        output := &sagemaker.ListPipelinesOutput{}

        mockClient.On("ListPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProcessingJobs", func(t *testing.T) {
        input := &sagemaker.ListProcessingJobsInput{}
        output := &sagemaker.ListProcessingJobsOutput{}

        mockClient.On("ListProcessingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListProcessingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &sagemaker.ListProjectsInput{}
        output := &sagemaker.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceCatalogs", func(t *testing.T) {
        input := &sagemaker.ListResourceCatalogsInput{}
        output := &sagemaker.ListResourceCatalogsOutput{}

        mockClient.On("ListResourceCatalogs", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceCatalogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpaces", func(t *testing.T) {
        input := &sagemaker.ListSpacesInput{}
        output := &sagemaker.ListSpacesOutput{}

        mockClient.On("ListSpaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStageDevices", func(t *testing.T) {
        input := &sagemaker.ListStageDevicesInput{}
        output := &sagemaker.ListStageDevicesOutput{}

        mockClient.On("ListStageDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListStageDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudioLifecycleConfigs", func(t *testing.T) {
        input := &sagemaker.ListStudioLifecycleConfigsInput{}
        output := &sagemaker.ListStudioLifecycleConfigsOutput{}

        mockClient.On("ListStudioLifecycleConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudioLifecycleConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscribedWorkteams", func(t *testing.T) {
        input := &sagemaker.ListSubscribedWorkteamsInput{}
        output := &sagemaker.ListSubscribedWorkteamsOutput{}

        mockClient.On("ListSubscribedWorkteams", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscribedWorkteams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &sagemaker.ListTagsInput{}
        output := &sagemaker.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrainingJobs", func(t *testing.T) {
        input := &sagemaker.ListTrainingJobsInput{}
        output := &sagemaker.ListTrainingJobsOutput{}

        mockClient.On("ListTrainingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrainingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrainingJobsForHyperParameterTuningJob", func(t *testing.T) {
        input := &sagemaker.ListTrainingJobsForHyperParameterTuningJobInput{}
        output := &sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput{}

        mockClient.On("ListTrainingJobsForHyperParameterTuningJob", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrainingJobsForHyperParameterTuningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTransformJobs", func(t *testing.T) {
        input := &sagemaker.ListTransformJobsInput{}
        output := &sagemaker.ListTransformJobsOutput{}

        mockClient.On("ListTransformJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTransformJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrialComponents", func(t *testing.T) {
        input := &sagemaker.ListTrialComponentsInput{}
        output := &sagemaker.ListTrialComponentsOutput{}

        mockClient.On("ListTrialComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrialComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrials", func(t *testing.T) {
        input := &sagemaker.ListTrialsInput{}
        output := &sagemaker.ListTrialsOutput{}

        mockClient.On("ListTrials", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserProfiles", func(t *testing.T) {
        input := &sagemaker.ListUserProfilesInput{}
        output := &sagemaker.ListUserProfilesOutput{}

        mockClient.On("ListUserProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkforces", func(t *testing.T) {
        input := &sagemaker.ListWorkforcesInput{}
        output := &sagemaker.ListWorkforcesOutput{}

        mockClient.On("ListWorkforces", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkforces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkteams", func(t *testing.T) {
        input := &sagemaker.ListWorkteamsInput{}
        output := &sagemaker.ListWorkteamsOutput{}

        mockClient.On("ListWorkteams", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkteams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutModelPackageGroupPolicy", func(t *testing.T) {
        input := &sagemaker.PutModelPackageGroupPolicyInput{}
        output := &sagemaker.PutModelPackageGroupPolicyOutput{}

        mockClient.On("PutModelPackageGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutModelPackageGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryLineage", func(t *testing.T) {
        input := &sagemaker.QueryLineageInput{}
        output := &sagemaker.QueryLineageOutput{}

        mockClient.On("QueryLineage", ctx, input).Return(output, nil)

        result, err := mockClient.QueryLineage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDevices", func(t *testing.T) {
        input := &sagemaker.RegisterDevicesInput{}
        output := &sagemaker.RegisterDevicesOutput{}

        mockClient.On("RegisterDevices", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRenderUiTemplate", func(t *testing.T) {
        input := &sagemaker.RenderUiTemplateInput{}
        output := &sagemaker.RenderUiTemplateOutput{}

        mockClient.On("RenderUiTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.RenderUiTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryPipelineExecution", func(t *testing.T) {
        input := &sagemaker.RetryPipelineExecutionInput{}
        output := &sagemaker.RetryPipelineExecutionOutput{}

        mockClient.On("RetryPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.RetryPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearch", func(t *testing.T) {
        input := &sagemaker.SearchInput{}
        output := &sagemaker.SearchOutput{}

        mockClient.On("Search", ctx, input).Return(output, nil)

        result, err := mockClient.Search(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendPipelineExecutionStepFailure", func(t *testing.T) {
        input := &sagemaker.SendPipelineExecutionStepFailureInput{}
        output := &sagemaker.SendPipelineExecutionStepFailureOutput{}

        mockClient.On("SendPipelineExecutionStepFailure", ctx, input).Return(output, nil)

        result, err := mockClient.SendPipelineExecutionStepFailure(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendPipelineExecutionStepSuccess", func(t *testing.T) {
        input := &sagemaker.SendPipelineExecutionStepSuccessInput{}
        output := &sagemaker.SendPipelineExecutionStepSuccessOutput{}

        mockClient.On("SendPipelineExecutionStepSuccess", ctx, input).Return(output, nil)

        result, err := mockClient.SendPipelineExecutionStepSuccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEdgeDeploymentStage", func(t *testing.T) {
        input := &sagemaker.StartEdgeDeploymentStageInput{}
        output := &sagemaker.StartEdgeDeploymentStageOutput{}

        mockClient.On("StartEdgeDeploymentStage", ctx, input).Return(output, nil)

        result, err := mockClient.StartEdgeDeploymentStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.StartInferenceExperimentInput{}
        output := &sagemaker.StartInferenceExperimentOutput{}

        mockClient.On("StartInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StartInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.StartMlflowTrackingServerInput{}
        output := &sagemaker.StartMlflowTrackingServerOutput{}

        mockClient.On("StartMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.StartMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.StartMonitoringScheduleInput{}
        output := &sagemaker.StartMonitoringScheduleOutput{}

        mockClient.On("StartMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.StartMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNotebookInstance", func(t *testing.T) {
        input := &sagemaker.StartNotebookInstanceInput{}
        output := &sagemaker.StartNotebookInstanceOutput{}

        mockClient.On("StartNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StartNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPipelineExecution", func(t *testing.T) {
        input := &sagemaker.StartPipelineExecutionInput{}
        output := &sagemaker.StartPipelineExecutionOutput{}

        mockClient.On("StartPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAutoMLJob", func(t *testing.T) {
        input := &sagemaker.StopAutoMLJobInput{}
        output := &sagemaker.StopAutoMLJobOutput{}

        mockClient.On("StopAutoMLJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopAutoMLJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCompilationJob", func(t *testing.T) {
        input := &sagemaker.StopCompilationJobInput{}
        output := &sagemaker.StopCompilationJobOutput{}

        mockClient.On("StopCompilationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopCompilationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEdgeDeploymentStage", func(t *testing.T) {
        input := &sagemaker.StopEdgeDeploymentStageInput{}
        output := &sagemaker.StopEdgeDeploymentStageOutput{}

        mockClient.On("StopEdgeDeploymentStage", ctx, input).Return(output, nil)

        result, err := mockClient.StopEdgeDeploymentStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEdgePackagingJob", func(t *testing.T) {
        input := &sagemaker.StopEdgePackagingJobInput{}
        output := &sagemaker.StopEdgePackagingJobOutput{}

        mockClient.On("StopEdgePackagingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopEdgePackagingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopHyperParameterTuningJob", func(t *testing.T) {
        input := &sagemaker.StopHyperParameterTuningJobInput{}
        output := &sagemaker.StopHyperParameterTuningJobOutput{}

        mockClient.On("StopHyperParameterTuningJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopHyperParameterTuningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.StopInferenceExperimentInput{}
        output := &sagemaker.StopInferenceExperimentOutput{}

        mockClient.On("StopInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StopInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInferenceRecommendationsJob", func(t *testing.T) {
        input := &sagemaker.StopInferenceRecommendationsJobInput{}
        output := &sagemaker.StopInferenceRecommendationsJobOutput{}

        mockClient.On("StopInferenceRecommendationsJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopInferenceRecommendationsJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopLabelingJob", func(t *testing.T) {
        input := &sagemaker.StopLabelingJobInput{}
        output := &sagemaker.StopLabelingJobOutput{}

        mockClient.On("StopLabelingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopLabelingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.StopMlflowTrackingServerInput{}
        output := &sagemaker.StopMlflowTrackingServerOutput{}

        mockClient.On("StopMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.StopMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.StopMonitoringScheduleInput{}
        output := &sagemaker.StopMonitoringScheduleOutput{}

        mockClient.On("StopMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.StopMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopNotebookInstance", func(t *testing.T) {
        input := &sagemaker.StopNotebookInstanceInput{}
        output := &sagemaker.StopNotebookInstanceOutput{}

        mockClient.On("StopNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StopNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopOptimizationJob", func(t *testing.T) {
        input := &sagemaker.StopOptimizationJobInput{}
        output := &sagemaker.StopOptimizationJobOutput{}

        mockClient.On("StopOptimizationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopOptimizationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPipelineExecution", func(t *testing.T) {
        input := &sagemaker.StopPipelineExecutionInput{}
        output := &sagemaker.StopPipelineExecutionOutput{}

        mockClient.On("StopPipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopPipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopProcessingJob", func(t *testing.T) {
        input := &sagemaker.StopProcessingJobInput{}
        output := &sagemaker.StopProcessingJobOutput{}

        mockClient.On("StopProcessingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopProcessingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTrainingJob", func(t *testing.T) {
        input := &sagemaker.StopTrainingJobInput{}
        output := &sagemaker.StopTrainingJobOutput{}

        mockClient.On("StopTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTransformJob", func(t *testing.T) {
        input := &sagemaker.StopTransformJobInput{}
        output := &sagemaker.StopTransformJobOutput{}

        mockClient.On("StopTransformJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopTransformJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAction", func(t *testing.T) {
        input := &sagemaker.UpdateActionInput{}
        output := &sagemaker.UpdateActionOutput{}

        mockClient.On("UpdateAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppImageConfig", func(t *testing.T) {
        input := &sagemaker.UpdateAppImageConfigInput{}
        output := &sagemaker.UpdateAppImageConfigOutput{}

        mockClient.On("UpdateAppImageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppImageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateArtifact", func(t *testing.T) {
        input := &sagemaker.UpdateArtifactInput{}
        output := &sagemaker.UpdateArtifactOutput{}

        mockClient.On("UpdateArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &sagemaker.UpdateClusterInput{}
        output := &sagemaker.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterSoftware", func(t *testing.T) {
        input := &sagemaker.UpdateClusterSoftwareInput{}
        output := &sagemaker.UpdateClusterSoftwareOutput{}

        mockClient.On("UpdateClusterSoftware", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterSoftware(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCodeRepository", func(t *testing.T) {
        input := &sagemaker.UpdateCodeRepositoryInput{}
        output := &sagemaker.UpdateCodeRepositoryOutput{}

        mockClient.On("UpdateCodeRepository", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCodeRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContext", func(t *testing.T) {
        input := &sagemaker.UpdateContextInput{}
        output := &sagemaker.UpdateContextOutput{}

        mockClient.On("UpdateContext", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceFleet", func(t *testing.T) {
        input := &sagemaker.UpdateDeviceFleetInput{}
        output := &sagemaker.UpdateDeviceFleetOutput{}

        mockClient.On("UpdateDeviceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevices", func(t *testing.T) {
        input := &sagemaker.UpdateDevicesInput{}
        output := &sagemaker.UpdateDevicesOutput{}

        mockClient.On("UpdateDevices", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomain", func(t *testing.T) {
        input := &sagemaker.UpdateDomainInput{}
        output := &sagemaker.UpdateDomainOutput{}

        mockClient.On("UpdateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpoint", func(t *testing.T) {
        input := &sagemaker.UpdateEndpointInput{}
        output := &sagemaker.UpdateEndpointOutput{}

        mockClient.On("UpdateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpointWeightsAndCapacities", func(t *testing.T) {
        input := &sagemaker.UpdateEndpointWeightsAndCapacitiesInput{}
        output := &sagemaker.UpdateEndpointWeightsAndCapacitiesOutput{}

        mockClient.On("UpdateEndpointWeightsAndCapacities", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpointWeightsAndCapacities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExperiment", func(t *testing.T) {
        input := &sagemaker.UpdateExperimentInput{}
        output := &sagemaker.UpdateExperimentOutput{}

        mockClient.On("UpdateExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFeatureGroup", func(t *testing.T) {
        input := &sagemaker.UpdateFeatureGroupInput{}
        output := &sagemaker.UpdateFeatureGroupOutput{}

        mockClient.On("UpdateFeatureGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFeatureGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFeatureMetadata", func(t *testing.T) {
        input := &sagemaker.UpdateFeatureMetadataInput{}
        output := &sagemaker.UpdateFeatureMetadataOutput{}

        mockClient.On("UpdateFeatureMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFeatureMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHub", func(t *testing.T) {
        input := &sagemaker.UpdateHubInput{}
        output := &sagemaker.UpdateHubOutput{}

        mockClient.On("UpdateHub", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImage", func(t *testing.T) {
        input := &sagemaker.UpdateImageInput{}
        output := &sagemaker.UpdateImageOutput{}

        mockClient.On("UpdateImage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImageVersion", func(t *testing.T) {
        input := &sagemaker.UpdateImageVersionInput{}
        output := &sagemaker.UpdateImageVersionOutput{}

        mockClient.On("UpdateImageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInferenceComponent", func(t *testing.T) {
        input := &sagemaker.UpdateInferenceComponentInput{}
        output := &sagemaker.UpdateInferenceComponentOutput{}

        mockClient.On("UpdateInferenceComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInferenceComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInferenceComponentRuntimeConfig", func(t *testing.T) {
        input := &sagemaker.UpdateInferenceComponentRuntimeConfigInput{}
        output := &sagemaker.UpdateInferenceComponentRuntimeConfigOutput{}

        mockClient.On("UpdateInferenceComponentRuntimeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInferenceComponentRuntimeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInferenceExperiment", func(t *testing.T) {
        input := &sagemaker.UpdateInferenceExperimentInput{}
        output := &sagemaker.UpdateInferenceExperimentOutput{}

        mockClient.On("UpdateInferenceExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInferenceExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMlflowTrackingServer", func(t *testing.T) {
        input := &sagemaker.UpdateMlflowTrackingServerInput{}
        output := &sagemaker.UpdateMlflowTrackingServerOutput{}

        mockClient.On("UpdateMlflowTrackingServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMlflowTrackingServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModelCard", func(t *testing.T) {
        input := &sagemaker.UpdateModelCardInput{}
        output := &sagemaker.UpdateModelCardOutput{}

        mockClient.On("UpdateModelCard", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModelCard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModelPackage", func(t *testing.T) {
        input := &sagemaker.UpdateModelPackageInput{}
        output := &sagemaker.UpdateModelPackageOutput{}

        mockClient.On("UpdateModelPackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModelPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitoringAlert", func(t *testing.T) {
        input := &sagemaker.UpdateMonitoringAlertInput{}
        output := &sagemaker.UpdateMonitoringAlertOutput{}

        mockClient.On("UpdateMonitoringAlert", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitoringAlert(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitoringSchedule", func(t *testing.T) {
        input := &sagemaker.UpdateMonitoringScheduleInput{}
        output := &sagemaker.UpdateMonitoringScheduleOutput{}

        mockClient.On("UpdateMonitoringSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitoringSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotebookInstance", func(t *testing.T) {
        input := &sagemaker.UpdateNotebookInstanceInput{}
        output := &sagemaker.UpdateNotebookInstanceOutput{}

        mockClient.On("UpdateNotebookInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotebookInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotebookInstanceLifecycleConfig", func(t *testing.T) {
        input := &sagemaker.UpdateNotebookInstanceLifecycleConfigInput{}
        output := &sagemaker.UpdateNotebookInstanceLifecycleConfigOutput{}

        mockClient.On("UpdateNotebookInstanceLifecycleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotebookInstanceLifecycleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipeline", func(t *testing.T) {
        input := &sagemaker.UpdatePipelineInput{}
        output := &sagemaker.UpdatePipelineOutput{}

        mockClient.On("UpdatePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipelineExecution", func(t *testing.T) {
        input := &sagemaker.UpdatePipelineExecutionInput{}
        output := &sagemaker.UpdatePipelineExecutionOutput{}

        mockClient.On("UpdatePipelineExecution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipelineExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &sagemaker.UpdateProjectInput{}
        output := &sagemaker.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSpace", func(t *testing.T) {
        input := &sagemaker.UpdateSpaceInput{}
        output := &sagemaker.UpdateSpaceOutput{}

        mockClient.On("UpdateSpace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrainingJob", func(t *testing.T) {
        input := &sagemaker.UpdateTrainingJobInput{}
        output := &sagemaker.UpdateTrainingJobOutput{}

        mockClient.On("UpdateTrainingJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrainingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrial", func(t *testing.T) {
        input := &sagemaker.UpdateTrialInput{}
        output := &sagemaker.UpdateTrialOutput{}

        mockClient.On("UpdateTrial", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrialComponent", func(t *testing.T) {
        input := &sagemaker.UpdateTrialComponentInput{}
        output := &sagemaker.UpdateTrialComponentOutput{}

        mockClient.On("UpdateTrialComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrialComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserProfile", func(t *testing.T) {
        input := &sagemaker.UpdateUserProfileInput{}
        output := &sagemaker.UpdateUserProfileOutput{}

        mockClient.On("UpdateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkforce", func(t *testing.T) {
        input := &sagemaker.UpdateWorkforceInput{}
        output := &sagemaker.UpdateWorkforceOutput{}

        mockClient.On("UpdateWorkforce", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkforce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkteam", func(t *testing.T) {
        input := &sagemaker.UpdateWorkteamInput{}
        output := &sagemaker.UpdateWorkteamOutput{}

        mockClient.On("UpdateWorkteam", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkteam(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
