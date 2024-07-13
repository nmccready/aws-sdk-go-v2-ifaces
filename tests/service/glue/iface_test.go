package glue_test

// tests for the glue service interface for this ../../../service/glue/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/glue/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/glue/glue_iface"
	"github.com/stretchr/testify/assert"
)

func TestGlueServiceCanBeMocked(t *testing.T) {
	var iface glue_iface.IClient
	iface = &glue.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := glue.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreatePartition", func(t *testing.T) {
        input := &glue.BatchCreatePartitionInput{}
        output := &glue.BatchCreatePartitionOutput{}

        mockClient.On("BatchCreatePartition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreatePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteConnection", func(t *testing.T) {
        input := &glue.BatchDeleteConnectionInput{}
        output := &glue.BatchDeleteConnectionOutput{}

        mockClient.On("BatchDeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeletePartition", func(t *testing.T) {
        input := &glue.BatchDeletePartitionInput{}
        output := &glue.BatchDeletePartitionOutput{}

        mockClient.On("BatchDeletePartition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeletePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteTable", func(t *testing.T) {
        input := &glue.BatchDeleteTableInput{}
        output := &glue.BatchDeleteTableOutput{}

        mockClient.On("BatchDeleteTable", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteTableVersion", func(t *testing.T) {
        input := &glue.BatchDeleteTableVersionInput{}
        output := &glue.BatchDeleteTableVersionOutput{}

        mockClient.On("BatchDeleteTableVersion", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteTableVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetBlueprints", func(t *testing.T) {
        input := &glue.BatchGetBlueprintsInput{}
        output := &glue.BatchGetBlueprintsOutput{}

        mockClient.On("BatchGetBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCrawlers", func(t *testing.T) {
        input := &glue.BatchGetCrawlersInput{}
        output := &glue.BatchGetCrawlersOutput{}

        mockClient.On("BatchGetCrawlers", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCrawlers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCustomEntityTypes", func(t *testing.T) {
        input := &glue.BatchGetCustomEntityTypesInput{}
        output := &glue.BatchGetCustomEntityTypesOutput{}

        mockClient.On("BatchGetCustomEntityTypes", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCustomEntityTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDataQualityResult", func(t *testing.T) {
        input := &glue.BatchGetDataQualityResultInput{}
        output := &glue.BatchGetDataQualityResultOutput{}

        mockClient.On("BatchGetDataQualityResult", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDataQualityResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDevEndpoints", func(t *testing.T) {
        input := &glue.BatchGetDevEndpointsInput{}
        output := &glue.BatchGetDevEndpointsOutput{}

        mockClient.On("BatchGetDevEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDevEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetJobs", func(t *testing.T) {
        input := &glue.BatchGetJobsInput{}
        output := &glue.BatchGetJobsOutput{}

        mockClient.On("BatchGetJobs", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetPartition", func(t *testing.T) {
        input := &glue.BatchGetPartitionInput{}
        output := &glue.BatchGetPartitionOutput{}

        mockClient.On("BatchGetPartition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetPartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetTableOptimizer", func(t *testing.T) {
        input := &glue.BatchGetTableOptimizerInput{}
        output := &glue.BatchGetTableOptimizerOutput{}

        mockClient.On("BatchGetTableOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetTableOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetTriggers", func(t *testing.T) {
        input := &glue.BatchGetTriggersInput{}
        output := &glue.BatchGetTriggersOutput{}

        mockClient.On("BatchGetTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetWorkflows", func(t *testing.T) {
        input := &glue.BatchGetWorkflowsInput{}
        output := &glue.BatchGetWorkflowsOutput{}

        mockClient.On("BatchGetWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStopJobRun", func(t *testing.T) {
        input := &glue.BatchStopJobRunInput{}
        output := &glue.BatchStopJobRunOutput{}

        mockClient.On("BatchStopJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStopJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdatePartition", func(t *testing.T) {
        input := &glue.BatchUpdatePartitionInput{}
        output := &glue.BatchUpdatePartitionOutput{}

        mockClient.On("BatchUpdatePartition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdatePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDataQualityRuleRecommendationRun", func(t *testing.T) {
        input := &glue.CancelDataQualityRuleRecommendationRunInput{}
        output := &glue.CancelDataQualityRuleRecommendationRunOutput{}

        mockClient.On("CancelDataQualityRuleRecommendationRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDataQualityRuleRecommendationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDataQualityRulesetEvaluationRun", func(t *testing.T) {
        input := &glue.CancelDataQualityRulesetEvaluationRunInput{}
        output := &glue.CancelDataQualityRulesetEvaluationRunOutput{}

        mockClient.On("CancelDataQualityRulesetEvaluationRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDataQualityRulesetEvaluationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMLTaskRun", func(t *testing.T) {
        input := &glue.CancelMLTaskRunInput{}
        output := &glue.CancelMLTaskRunOutput{}

        mockClient.On("CancelMLTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMLTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelStatement", func(t *testing.T) {
        input := &glue.CancelStatementInput{}
        output := &glue.CancelStatementOutput{}

        mockClient.On("CancelStatement", ctx, input).Return(output, nil)

        result, err := mockClient.CancelStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckSchemaVersionValidity", func(t *testing.T) {
        input := &glue.CheckSchemaVersionValidityInput{}
        output := &glue.CheckSchemaVersionValidityOutput{}

        mockClient.On("CheckSchemaVersionValidity", ctx, input).Return(output, nil)

        result, err := mockClient.CheckSchemaVersionValidity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBlueprint", func(t *testing.T) {
        input := &glue.CreateBlueprintInput{}
        output := &glue.CreateBlueprintOutput{}

        mockClient.On("CreateBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClassifier", func(t *testing.T) {
        input := &glue.CreateClassifierInput{}
        output := &glue.CreateClassifierOutput{}

        mockClient.On("CreateClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &glue.CreateConnectionInput{}
        output := &glue.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCrawler", func(t *testing.T) {
        input := &glue.CreateCrawlerInput{}
        output := &glue.CreateCrawlerOutput{}

        mockClient.On("CreateCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomEntityType", func(t *testing.T) {
        input := &glue.CreateCustomEntityTypeInput{}
        output := &glue.CreateCustomEntityTypeOutput{}

        mockClient.On("CreateCustomEntityType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomEntityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataQualityRuleset", func(t *testing.T) {
        input := &glue.CreateDataQualityRulesetInput{}
        output := &glue.CreateDataQualityRulesetOutput{}

        mockClient.On("CreateDataQualityRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataQualityRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatabase", func(t *testing.T) {
        input := &glue.CreateDatabaseInput{}
        output := &glue.CreateDatabaseOutput{}

        mockClient.On("CreateDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDevEndpoint", func(t *testing.T) {
        input := &glue.CreateDevEndpointInput{}
        output := &glue.CreateDevEndpointOutput{}

        mockClient.On("CreateDevEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDevEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &glue.CreateJobInput{}
        output := &glue.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMLTransform", func(t *testing.T) {
        input := &glue.CreateMLTransformInput{}
        output := &glue.CreateMLTransformOutput{}

        mockClient.On("CreateMLTransform", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMLTransform(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartition", func(t *testing.T) {
        input := &glue.CreatePartitionInput{}
        output := &glue.CreatePartitionOutput{}

        mockClient.On("CreatePartition", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartitionIndex", func(t *testing.T) {
        input := &glue.CreatePartitionIndexInput{}
        output := &glue.CreatePartitionIndexOutput{}

        mockClient.On("CreatePartitionIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartitionIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistry", func(t *testing.T) {
        input := &glue.CreateRegistryInput{}
        output := &glue.CreateRegistryOutput{}

        mockClient.On("CreateRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchema", func(t *testing.T) {
        input := &glue.CreateSchemaInput{}
        output := &glue.CreateSchemaOutput{}

        mockClient.On("CreateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScript", func(t *testing.T) {
        input := &glue.CreateScriptInput{}
        output := &glue.CreateScriptOutput{}

        mockClient.On("CreateScript", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityConfiguration", func(t *testing.T) {
        input := &glue.CreateSecurityConfigurationInput{}
        output := &glue.CreateSecurityConfigurationOutput{}

        mockClient.On("CreateSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSession", func(t *testing.T) {
        input := &glue.CreateSessionInput{}
        output := &glue.CreateSessionOutput{}

        mockClient.On("CreateSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTable", func(t *testing.T) {
        input := &glue.CreateTableInput{}
        output := &glue.CreateTableOutput{}

        mockClient.On("CreateTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTableOptimizer", func(t *testing.T) {
        input := &glue.CreateTableOptimizerInput{}
        output := &glue.CreateTableOptimizerOutput{}

        mockClient.On("CreateTableOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTableOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrigger", func(t *testing.T) {
        input := &glue.CreateTriggerInput{}
        output := &glue.CreateTriggerOutput{}

        mockClient.On("CreateTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsageProfile", func(t *testing.T) {
        input := &glue.CreateUsageProfileInput{}
        output := &glue.CreateUsageProfileOutput{}

        mockClient.On("CreateUsageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserDefinedFunction", func(t *testing.T) {
        input := &glue.CreateUserDefinedFunctionInput{}
        output := &glue.CreateUserDefinedFunctionOutput{}

        mockClient.On("CreateUserDefinedFunction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserDefinedFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflow", func(t *testing.T) {
        input := &glue.CreateWorkflowInput{}
        output := &glue.CreateWorkflowOutput{}

        mockClient.On("CreateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBlueprint", func(t *testing.T) {
        input := &glue.DeleteBlueprintInput{}
        output := &glue.DeleteBlueprintOutput{}

        mockClient.On("DeleteBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClassifier", func(t *testing.T) {
        input := &glue.DeleteClassifierInput{}
        output := &glue.DeleteClassifierOutput{}

        mockClient.On("DeleteClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteColumnStatisticsForPartition", func(t *testing.T) {
        input := &glue.DeleteColumnStatisticsForPartitionInput{}
        output := &glue.DeleteColumnStatisticsForPartitionOutput{}

        mockClient.On("DeleteColumnStatisticsForPartition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteColumnStatisticsForPartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteColumnStatisticsForTable", func(t *testing.T) {
        input := &glue.DeleteColumnStatisticsForTableInput{}
        output := &glue.DeleteColumnStatisticsForTableOutput{}

        mockClient.On("DeleteColumnStatisticsForTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteColumnStatisticsForTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &glue.DeleteConnectionInput{}
        output := &glue.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCrawler", func(t *testing.T) {
        input := &glue.DeleteCrawlerInput{}
        output := &glue.DeleteCrawlerOutput{}

        mockClient.On("DeleteCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomEntityType", func(t *testing.T) {
        input := &glue.DeleteCustomEntityTypeInput{}
        output := &glue.DeleteCustomEntityTypeOutput{}

        mockClient.On("DeleteCustomEntityType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomEntityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataQualityRuleset", func(t *testing.T) {
        input := &glue.DeleteDataQualityRulesetInput{}
        output := &glue.DeleteDataQualityRulesetOutput{}

        mockClient.On("DeleteDataQualityRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataQualityRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatabase", func(t *testing.T) {
        input := &glue.DeleteDatabaseInput{}
        output := &glue.DeleteDatabaseOutput{}

        mockClient.On("DeleteDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevEndpoint", func(t *testing.T) {
        input := &glue.DeleteDevEndpointInput{}
        output := &glue.DeleteDevEndpointOutput{}

        mockClient.On("DeleteDevEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &glue.DeleteJobInput{}
        output := &glue.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMLTransform", func(t *testing.T) {
        input := &glue.DeleteMLTransformInput{}
        output := &glue.DeleteMLTransformOutput{}

        mockClient.On("DeleteMLTransform", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMLTransform(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartition", func(t *testing.T) {
        input := &glue.DeletePartitionInput{}
        output := &glue.DeletePartitionOutput{}

        mockClient.On("DeletePartition", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartitionIndex", func(t *testing.T) {
        input := &glue.DeletePartitionIndexInput{}
        output := &glue.DeletePartitionIndexOutput{}

        mockClient.On("DeletePartitionIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartitionIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistry", func(t *testing.T) {
        input := &glue.DeleteRegistryInput{}
        output := &glue.DeleteRegistryOutput{}

        mockClient.On("DeleteRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &glue.DeleteResourcePolicyInput{}
        output := &glue.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchema", func(t *testing.T) {
        input := &glue.DeleteSchemaInput{}
        output := &glue.DeleteSchemaOutput{}

        mockClient.On("DeleteSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchemaVersions", func(t *testing.T) {
        input := &glue.DeleteSchemaVersionsInput{}
        output := &glue.DeleteSchemaVersionsOutput{}

        mockClient.On("DeleteSchemaVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchemaVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityConfiguration", func(t *testing.T) {
        input := &glue.DeleteSecurityConfigurationInput{}
        output := &glue.DeleteSecurityConfigurationOutput{}

        mockClient.On("DeleteSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSession", func(t *testing.T) {
        input := &glue.DeleteSessionInput{}
        output := &glue.DeleteSessionOutput{}

        mockClient.On("DeleteSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTable", func(t *testing.T) {
        input := &glue.DeleteTableInput{}
        output := &glue.DeleteTableOutput{}

        mockClient.On("DeleteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTableOptimizer", func(t *testing.T) {
        input := &glue.DeleteTableOptimizerInput{}
        output := &glue.DeleteTableOptimizerOutput{}

        mockClient.On("DeleteTableOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTableOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTableVersion", func(t *testing.T) {
        input := &glue.DeleteTableVersionInput{}
        output := &glue.DeleteTableVersionOutput{}

        mockClient.On("DeleteTableVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTableVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrigger", func(t *testing.T) {
        input := &glue.DeleteTriggerInput{}
        output := &glue.DeleteTriggerOutput{}

        mockClient.On("DeleteTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsageProfile", func(t *testing.T) {
        input := &glue.DeleteUsageProfileInput{}
        output := &glue.DeleteUsageProfileOutput{}

        mockClient.On("DeleteUsageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserDefinedFunction", func(t *testing.T) {
        input := &glue.DeleteUserDefinedFunctionInput{}
        output := &glue.DeleteUserDefinedFunctionOutput{}

        mockClient.On("DeleteUserDefinedFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserDefinedFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &glue.DeleteWorkflowInput{}
        output := &glue.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlueprint", func(t *testing.T) {
        input := &glue.GetBlueprintInput{}
        output := &glue.GetBlueprintOutput{}

        mockClient.On("GetBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlueprintRun", func(t *testing.T) {
        input := &glue.GetBlueprintRunInput{}
        output := &glue.GetBlueprintRunOutput{}

        mockClient.On("GetBlueprintRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlueprintRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlueprintRuns", func(t *testing.T) {
        input := &glue.GetBlueprintRunsInput{}
        output := &glue.GetBlueprintRunsOutput{}

        mockClient.On("GetBlueprintRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlueprintRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCatalogImportStatus", func(t *testing.T) {
        input := &glue.GetCatalogImportStatusInput{}
        output := &glue.GetCatalogImportStatusOutput{}

        mockClient.On("GetCatalogImportStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetCatalogImportStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClassifier", func(t *testing.T) {
        input := &glue.GetClassifierInput{}
        output := &glue.GetClassifierOutput{}

        mockClient.On("GetClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.GetClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClassifiers", func(t *testing.T) {
        input := &glue.GetClassifiersInput{}
        output := &glue.GetClassifiersOutput{}

        mockClient.On("GetClassifiers", ctx, input).Return(output, nil)

        result, err := mockClient.GetClassifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetColumnStatisticsForPartition", func(t *testing.T) {
        input := &glue.GetColumnStatisticsForPartitionInput{}
        output := &glue.GetColumnStatisticsForPartitionOutput{}

        mockClient.On("GetColumnStatisticsForPartition", ctx, input).Return(output, nil)

        result, err := mockClient.GetColumnStatisticsForPartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetColumnStatisticsForTable", func(t *testing.T) {
        input := &glue.GetColumnStatisticsForTableInput{}
        output := &glue.GetColumnStatisticsForTableOutput{}

        mockClient.On("GetColumnStatisticsForTable", ctx, input).Return(output, nil)

        result, err := mockClient.GetColumnStatisticsForTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetColumnStatisticsTaskRun", func(t *testing.T) {
        input := &glue.GetColumnStatisticsTaskRunInput{}
        output := &glue.GetColumnStatisticsTaskRunOutput{}

        mockClient.On("GetColumnStatisticsTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetColumnStatisticsTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetColumnStatisticsTaskRuns", func(t *testing.T) {
        input := &glue.GetColumnStatisticsTaskRunsInput{}
        output := &glue.GetColumnStatisticsTaskRunsOutput{}

        mockClient.On("GetColumnStatisticsTaskRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetColumnStatisticsTaskRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnection", func(t *testing.T) {
        input := &glue.GetConnectionInput{}
        output := &glue.GetConnectionOutput{}

        mockClient.On("GetConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnections", func(t *testing.T) {
        input := &glue.GetConnectionsInput{}
        output := &glue.GetConnectionsOutput{}

        mockClient.On("GetConnections", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCrawler", func(t *testing.T) {
        input := &glue.GetCrawlerInput{}
        output := &glue.GetCrawlerOutput{}

        mockClient.On("GetCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.GetCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCrawlerMetrics", func(t *testing.T) {
        input := &glue.GetCrawlerMetricsInput{}
        output := &glue.GetCrawlerMetricsOutput{}

        mockClient.On("GetCrawlerMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetCrawlerMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCrawlers", func(t *testing.T) {
        input := &glue.GetCrawlersInput{}
        output := &glue.GetCrawlersOutput{}

        mockClient.On("GetCrawlers", ctx, input).Return(output, nil)

        result, err := mockClient.GetCrawlers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomEntityType", func(t *testing.T) {
        input := &glue.GetCustomEntityTypeInput{}
        output := &glue.GetCustomEntityTypeOutput{}

        mockClient.On("GetCustomEntityType", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomEntityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataCatalogEncryptionSettings", func(t *testing.T) {
        input := &glue.GetDataCatalogEncryptionSettingsInput{}
        output := &glue.GetDataCatalogEncryptionSettingsOutput{}

        mockClient.On("GetDataCatalogEncryptionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataCatalogEncryptionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataQualityResult", func(t *testing.T) {
        input := &glue.GetDataQualityResultInput{}
        output := &glue.GetDataQualityResultOutput{}

        mockClient.On("GetDataQualityResult", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataQualityResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataQualityRuleRecommendationRun", func(t *testing.T) {
        input := &glue.GetDataQualityRuleRecommendationRunInput{}
        output := &glue.GetDataQualityRuleRecommendationRunOutput{}

        mockClient.On("GetDataQualityRuleRecommendationRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataQualityRuleRecommendationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataQualityRuleset", func(t *testing.T) {
        input := &glue.GetDataQualityRulesetInput{}
        output := &glue.GetDataQualityRulesetOutput{}

        mockClient.On("GetDataQualityRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataQualityRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataQualityRulesetEvaluationRun", func(t *testing.T) {
        input := &glue.GetDataQualityRulesetEvaluationRunInput{}
        output := &glue.GetDataQualityRulesetEvaluationRunOutput{}

        mockClient.On("GetDataQualityRulesetEvaluationRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataQualityRulesetEvaluationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatabase", func(t *testing.T) {
        input := &glue.GetDatabaseInput{}
        output := &glue.GetDatabaseOutput{}

        mockClient.On("GetDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatabases", func(t *testing.T) {
        input := &glue.GetDatabasesInput{}
        output := &glue.GetDatabasesOutput{}

        mockClient.On("GetDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataflowGraph", func(t *testing.T) {
        input := &glue.GetDataflowGraphInput{}
        output := &glue.GetDataflowGraphOutput{}

        mockClient.On("GetDataflowGraph", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataflowGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevEndpoint", func(t *testing.T) {
        input := &glue.GetDevEndpointInput{}
        output := &glue.GetDevEndpointOutput{}

        mockClient.On("GetDevEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevEndpoints", func(t *testing.T) {
        input := &glue.GetDevEndpointsInput{}
        output := &glue.GetDevEndpointsOutput{}

        mockClient.On("GetDevEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &glue.GetJobInput{}
        output := &glue.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobBookmark", func(t *testing.T) {
        input := &glue.GetJobBookmarkInput{}
        output := &glue.GetJobBookmarkOutput{}

        mockClient.On("GetJobBookmark", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobBookmark(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobRun", func(t *testing.T) {
        input := &glue.GetJobRunInput{}
        output := &glue.GetJobRunOutput{}

        mockClient.On("GetJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobRuns", func(t *testing.T) {
        input := &glue.GetJobRunsInput{}
        output := &glue.GetJobRunsOutput{}

        mockClient.On("GetJobRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobs", func(t *testing.T) {
        input := &glue.GetJobsInput{}
        output := &glue.GetJobsOutput{}

        mockClient.On("GetJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLTaskRun", func(t *testing.T) {
        input := &glue.GetMLTaskRunInput{}
        output := &glue.GetMLTaskRunOutput{}

        mockClient.On("GetMLTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLTaskRuns", func(t *testing.T) {
        input := &glue.GetMLTaskRunsInput{}
        output := &glue.GetMLTaskRunsOutput{}

        mockClient.On("GetMLTaskRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLTaskRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLTransform", func(t *testing.T) {
        input := &glue.GetMLTransformInput{}
        output := &glue.GetMLTransformOutput{}

        mockClient.On("GetMLTransform", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLTransform(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMLTransforms", func(t *testing.T) {
        input := &glue.GetMLTransformsInput{}
        output := &glue.GetMLTransformsOutput{}

        mockClient.On("GetMLTransforms", ctx, input).Return(output, nil)

        result, err := mockClient.GetMLTransforms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMapping", func(t *testing.T) {
        input := &glue.GetMappingInput{}
        output := &glue.GetMappingOutput{}

        mockClient.On("GetMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPartition", func(t *testing.T) {
        input := &glue.GetPartitionInput{}
        output := &glue.GetPartitionOutput{}

        mockClient.On("GetPartition", ctx, input).Return(output, nil)

        result, err := mockClient.GetPartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPartitionIndexes", func(t *testing.T) {
        input := &glue.GetPartitionIndexesInput{}
        output := &glue.GetPartitionIndexesOutput{}

        mockClient.On("GetPartitionIndexes", ctx, input).Return(output, nil)

        result, err := mockClient.GetPartitionIndexes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPartitions", func(t *testing.T) {
        input := &glue.GetPartitionsInput{}
        output := &glue.GetPartitionsOutput{}

        mockClient.On("GetPartitions", ctx, input).Return(output, nil)

        result, err := mockClient.GetPartitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlan", func(t *testing.T) {
        input := &glue.GetPlanInput{}
        output := &glue.GetPlanOutput{}

        mockClient.On("GetPlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegistry", func(t *testing.T) {
        input := &glue.GetRegistryInput{}
        output := &glue.GetRegistryOutput{}

        mockClient.On("GetRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicies", func(t *testing.T) {
        input := &glue.GetResourcePoliciesInput{}
        output := &glue.GetResourcePoliciesOutput{}

        mockClient.On("GetResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &glue.GetResourcePolicyInput{}
        output := &glue.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchema", func(t *testing.T) {
        input := &glue.GetSchemaInput{}
        output := &glue.GetSchemaOutput{}

        mockClient.On("GetSchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaByDefinition", func(t *testing.T) {
        input := &glue.GetSchemaByDefinitionInput{}
        output := &glue.GetSchemaByDefinitionOutput{}

        mockClient.On("GetSchemaByDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaByDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaVersion", func(t *testing.T) {
        input := &glue.GetSchemaVersionInput{}
        output := &glue.GetSchemaVersionOutput{}

        mockClient.On("GetSchemaVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaVersionsDiff", func(t *testing.T) {
        input := &glue.GetSchemaVersionsDiffInput{}
        output := &glue.GetSchemaVersionsDiffOutput{}

        mockClient.On("GetSchemaVersionsDiff", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaVersionsDiff(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityConfiguration", func(t *testing.T) {
        input := &glue.GetSecurityConfigurationInput{}
        output := &glue.GetSecurityConfigurationOutput{}

        mockClient.On("GetSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityConfigurations", func(t *testing.T) {
        input := &glue.GetSecurityConfigurationsInput{}
        output := &glue.GetSecurityConfigurationsOutput{}

        mockClient.On("GetSecurityConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &glue.GetSessionInput{}
        output := &glue.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStatement", func(t *testing.T) {
        input := &glue.GetStatementInput{}
        output := &glue.GetStatementOutput{}

        mockClient.On("GetStatement", ctx, input).Return(output, nil)

        result, err := mockClient.GetStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTable", func(t *testing.T) {
        input := &glue.GetTableInput{}
        output := &glue.GetTableOutput{}

        mockClient.On("GetTable", ctx, input).Return(output, nil)

        result, err := mockClient.GetTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableOptimizer", func(t *testing.T) {
        input := &glue.GetTableOptimizerInput{}
        output := &glue.GetTableOptimizerOutput{}

        mockClient.On("GetTableOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableVersion", func(t *testing.T) {
        input := &glue.GetTableVersionInput{}
        output := &glue.GetTableVersionOutput{}

        mockClient.On("GetTableVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableVersions", func(t *testing.T) {
        input := &glue.GetTableVersionsInput{}
        output := &glue.GetTableVersionsOutput{}

        mockClient.On("GetTableVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTables", func(t *testing.T) {
        input := &glue.GetTablesInput{}
        output := &glue.GetTablesOutput{}

        mockClient.On("GetTables", ctx, input).Return(output, nil)

        result, err := mockClient.GetTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTags", func(t *testing.T) {
        input := &glue.GetTagsInput{}
        output := &glue.GetTagsOutput{}

        mockClient.On("GetTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrigger", func(t *testing.T) {
        input := &glue.GetTriggerInput{}
        output := &glue.GetTriggerOutput{}

        mockClient.On("GetTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTriggers", func(t *testing.T) {
        input := &glue.GetTriggersInput{}
        output := &glue.GetTriggersOutput{}

        mockClient.On("GetTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.GetTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUnfilteredPartitionMetadata", func(t *testing.T) {
        input := &glue.GetUnfilteredPartitionMetadataInput{}
        output := &glue.GetUnfilteredPartitionMetadataOutput{}

        mockClient.On("GetUnfilteredPartitionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetUnfilteredPartitionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUnfilteredPartitionsMetadata", func(t *testing.T) {
        input := &glue.GetUnfilteredPartitionsMetadataInput{}
        output := &glue.GetUnfilteredPartitionsMetadataOutput{}

        mockClient.On("GetUnfilteredPartitionsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetUnfilteredPartitionsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUnfilteredTableMetadata", func(t *testing.T) {
        input := &glue.GetUnfilteredTableMetadataInput{}
        output := &glue.GetUnfilteredTableMetadataOutput{}

        mockClient.On("GetUnfilteredTableMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetUnfilteredTableMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageProfile", func(t *testing.T) {
        input := &glue.GetUsageProfileInput{}
        output := &glue.GetUsageProfileOutput{}

        mockClient.On("GetUsageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserDefinedFunction", func(t *testing.T) {
        input := &glue.GetUserDefinedFunctionInput{}
        output := &glue.GetUserDefinedFunctionOutput{}

        mockClient.On("GetUserDefinedFunction", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserDefinedFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserDefinedFunctions", func(t *testing.T) {
        input := &glue.GetUserDefinedFunctionsInput{}
        output := &glue.GetUserDefinedFunctionsOutput{}

        mockClient.On("GetUserDefinedFunctions", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserDefinedFunctions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &glue.GetWorkflowInput{}
        output := &glue.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowRun", func(t *testing.T) {
        input := &glue.GetWorkflowRunInput{}
        output := &glue.GetWorkflowRunOutput{}

        mockClient.On("GetWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowRunProperties", func(t *testing.T) {
        input := &glue.GetWorkflowRunPropertiesInput{}
        output := &glue.GetWorkflowRunPropertiesOutput{}

        mockClient.On("GetWorkflowRunProperties", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowRunProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowRuns", func(t *testing.T) {
        input := &glue.GetWorkflowRunsInput{}
        output := &glue.GetWorkflowRunsOutput{}

        mockClient.On("GetWorkflowRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCatalogToGlue", func(t *testing.T) {
        input := &glue.ImportCatalogToGlueInput{}
        output := &glue.ImportCatalogToGlueOutput{}

        mockClient.On("ImportCatalogToGlue", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCatalogToGlue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBlueprints", func(t *testing.T) {
        input := &glue.ListBlueprintsInput{}
        output := &glue.ListBlueprintsOutput{}

        mockClient.On("ListBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.ListBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListColumnStatisticsTaskRuns", func(t *testing.T) {
        input := &glue.ListColumnStatisticsTaskRunsInput{}
        output := &glue.ListColumnStatisticsTaskRunsOutput{}

        mockClient.On("ListColumnStatisticsTaskRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListColumnStatisticsTaskRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrawlers", func(t *testing.T) {
        input := &glue.ListCrawlersInput{}
        output := &glue.ListCrawlersOutput{}

        mockClient.On("ListCrawlers", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrawlers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrawls", func(t *testing.T) {
        input := &glue.ListCrawlsInput{}
        output := &glue.ListCrawlsOutput{}

        mockClient.On("ListCrawls", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrawls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomEntityTypes", func(t *testing.T) {
        input := &glue.ListCustomEntityTypesInput{}
        output := &glue.ListCustomEntityTypesOutput{}

        mockClient.On("ListCustomEntityTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomEntityTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataQualityResults", func(t *testing.T) {
        input := &glue.ListDataQualityResultsInput{}
        output := &glue.ListDataQualityResultsOutput{}

        mockClient.On("ListDataQualityResults", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataQualityResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataQualityRuleRecommendationRuns", func(t *testing.T) {
        input := &glue.ListDataQualityRuleRecommendationRunsInput{}
        output := &glue.ListDataQualityRuleRecommendationRunsOutput{}

        mockClient.On("ListDataQualityRuleRecommendationRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataQualityRuleRecommendationRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataQualityRulesetEvaluationRuns", func(t *testing.T) {
        input := &glue.ListDataQualityRulesetEvaluationRunsInput{}
        output := &glue.ListDataQualityRulesetEvaluationRunsOutput{}

        mockClient.On("ListDataQualityRulesetEvaluationRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataQualityRulesetEvaluationRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataQualityRulesets", func(t *testing.T) {
        input := &glue.ListDataQualityRulesetsInput{}
        output := &glue.ListDataQualityRulesetsOutput{}

        mockClient.On("ListDataQualityRulesets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataQualityRulesets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevEndpoints", func(t *testing.T) {
        input := &glue.ListDevEndpointsInput{}
        output := &glue.ListDevEndpointsOutput{}

        mockClient.On("ListDevEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &glue.ListJobsInput{}
        output := &glue.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMLTransforms", func(t *testing.T) {
        input := &glue.ListMLTransformsInput{}
        output := &glue.ListMLTransformsOutput{}

        mockClient.On("ListMLTransforms", ctx, input).Return(output, nil)

        result, err := mockClient.ListMLTransforms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegistries", func(t *testing.T) {
        input := &glue.ListRegistriesInput{}
        output := &glue.ListRegistriesOutput{}

        mockClient.On("ListRegistries", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegistries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemaVersions", func(t *testing.T) {
        input := &glue.ListSchemaVersionsInput{}
        output := &glue.ListSchemaVersionsOutput{}

        mockClient.On("ListSchemaVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemaVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemas", func(t *testing.T) {
        input := &glue.ListSchemasInput{}
        output := &glue.ListSchemasOutput{}

        mockClient.On("ListSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessions", func(t *testing.T) {
        input := &glue.ListSessionsInput{}
        output := &glue.ListSessionsOutput{}

        mockClient.On("ListSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStatements", func(t *testing.T) {
        input := &glue.ListStatementsInput{}
        output := &glue.ListStatementsOutput{}

        mockClient.On("ListStatements", ctx, input).Return(output, nil)

        result, err := mockClient.ListStatements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTableOptimizerRuns", func(t *testing.T) {
        input := &glue.ListTableOptimizerRunsInput{}
        output := &glue.ListTableOptimizerRunsOutput{}

        mockClient.On("ListTableOptimizerRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListTableOptimizerRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTriggers", func(t *testing.T) {
        input := &glue.ListTriggersInput{}
        output := &glue.ListTriggersOutput{}

        mockClient.On("ListTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsageProfiles", func(t *testing.T) {
        input := &glue.ListUsageProfilesInput{}
        output := &glue.ListUsageProfilesOutput{}

        mockClient.On("ListUsageProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsageProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &glue.ListWorkflowsInput{}
        output := &glue.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDataCatalogEncryptionSettings", func(t *testing.T) {
        input := &glue.PutDataCatalogEncryptionSettingsInput{}
        output := &glue.PutDataCatalogEncryptionSettingsOutput{}

        mockClient.On("PutDataCatalogEncryptionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutDataCatalogEncryptionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &glue.PutResourcePolicyInput{}
        output := &glue.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSchemaVersionMetadata", func(t *testing.T) {
        input := &glue.PutSchemaVersionMetadataInput{}
        output := &glue.PutSchemaVersionMetadataOutput{}

        mockClient.On("PutSchemaVersionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.PutSchemaVersionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutWorkflowRunProperties", func(t *testing.T) {
        input := &glue.PutWorkflowRunPropertiesInput{}
        output := &glue.PutWorkflowRunPropertiesOutput{}

        mockClient.On("PutWorkflowRunProperties", ctx, input).Return(output, nil)

        result, err := mockClient.PutWorkflowRunProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQuerySchemaVersionMetadata", func(t *testing.T) {
        input := &glue.QuerySchemaVersionMetadataInput{}
        output := &glue.QuerySchemaVersionMetadataOutput{}

        mockClient.On("QuerySchemaVersionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.QuerySchemaVersionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterSchemaVersion", func(t *testing.T) {
        input := &glue.RegisterSchemaVersionInput{}
        output := &glue.RegisterSchemaVersionOutput{}

        mockClient.On("RegisterSchemaVersion", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterSchemaVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveSchemaVersionMetadata", func(t *testing.T) {
        input := &glue.RemoveSchemaVersionMetadataInput{}
        output := &glue.RemoveSchemaVersionMetadataOutput{}

        mockClient.On("RemoveSchemaVersionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveSchemaVersionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetJobBookmark", func(t *testing.T) {
        input := &glue.ResetJobBookmarkInput{}
        output := &glue.ResetJobBookmarkOutput{}

        mockClient.On("ResetJobBookmark", ctx, input).Return(output, nil)

        result, err := mockClient.ResetJobBookmark(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeWorkflowRun", func(t *testing.T) {
        input := &glue.ResumeWorkflowRunInput{}
        output := &glue.ResumeWorkflowRunOutput{}

        mockClient.On("ResumeWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunStatement", func(t *testing.T) {
        input := &glue.RunStatementInput{}
        output := &glue.RunStatementOutput{}

        mockClient.On("RunStatement", ctx, input).Return(output, nil)

        result, err := mockClient.RunStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTables", func(t *testing.T) {
        input := &glue.SearchTablesInput{}
        output := &glue.SearchTablesOutput{}

        mockClient.On("SearchTables", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBlueprintRun", func(t *testing.T) {
        input := &glue.StartBlueprintRunInput{}
        output := &glue.StartBlueprintRunOutput{}

        mockClient.On("StartBlueprintRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartBlueprintRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartColumnStatisticsTaskRun", func(t *testing.T) {
        input := &glue.StartColumnStatisticsTaskRunInput{}
        output := &glue.StartColumnStatisticsTaskRunOutput{}

        mockClient.On("StartColumnStatisticsTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartColumnStatisticsTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCrawler", func(t *testing.T) {
        input := &glue.StartCrawlerInput{}
        output := &glue.StartCrawlerOutput{}

        mockClient.On("StartCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.StartCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCrawlerSchedule", func(t *testing.T) {
        input := &glue.StartCrawlerScheduleInput{}
        output := &glue.StartCrawlerScheduleOutput{}

        mockClient.On("StartCrawlerSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.StartCrawlerSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataQualityRuleRecommendationRun", func(t *testing.T) {
        input := &glue.StartDataQualityRuleRecommendationRunInput{}
        output := &glue.StartDataQualityRuleRecommendationRunOutput{}

        mockClient.On("StartDataQualityRuleRecommendationRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataQualityRuleRecommendationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataQualityRulesetEvaluationRun", func(t *testing.T) {
        input := &glue.StartDataQualityRulesetEvaluationRunInput{}
        output := &glue.StartDataQualityRulesetEvaluationRunOutput{}

        mockClient.On("StartDataQualityRulesetEvaluationRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataQualityRulesetEvaluationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExportLabelsTaskRun", func(t *testing.T) {
        input := &glue.StartExportLabelsTaskRunInput{}
        output := &glue.StartExportLabelsTaskRunOutput{}

        mockClient.On("StartExportLabelsTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartExportLabelsTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportLabelsTaskRun", func(t *testing.T) {
        input := &glue.StartImportLabelsTaskRunInput{}
        output := &glue.StartImportLabelsTaskRunOutput{}

        mockClient.On("StartImportLabelsTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportLabelsTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJobRun", func(t *testing.T) {
        input := &glue.StartJobRunInput{}
        output := &glue.StartJobRunOutput{}

        mockClient.On("StartJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMLEvaluationTaskRun", func(t *testing.T) {
        input := &glue.StartMLEvaluationTaskRunInput{}
        output := &glue.StartMLEvaluationTaskRunOutput{}

        mockClient.On("StartMLEvaluationTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartMLEvaluationTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMLLabelingSetGenerationTaskRun", func(t *testing.T) {
        input := &glue.StartMLLabelingSetGenerationTaskRunInput{}
        output := &glue.StartMLLabelingSetGenerationTaskRunOutput{}

        mockClient.On("StartMLLabelingSetGenerationTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartMLLabelingSetGenerationTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTrigger", func(t *testing.T) {
        input := &glue.StartTriggerInput{}
        output := &glue.StartTriggerOutput{}

        mockClient.On("StartTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.StartTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkflowRun", func(t *testing.T) {
        input := &glue.StartWorkflowRunInput{}
        output := &glue.StartWorkflowRunOutput{}

        mockClient.On("StartWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopColumnStatisticsTaskRun", func(t *testing.T) {
        input := &glue.StopColumnStatisticsTaskRunInput{}
        output := &glue.StopColumnStatisticsTaskRunOutput{}

        mockClient.On("StopColumnStatisticsTaskRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopColumnStatisticsTaskRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCrawler", func(t *testing.T) {
        input := &glue.StopCrawlerInput{}
        output := &glue.StopCrawlerOutput{}

        mockClient.On("StopCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.StopCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCrawlerSchedule", func(t *testing.T) {
        input := &glue.StopCrawlerScheduleInput{}
        output := &glue.StopCrawlerScheduleOutput{}

        mockClient.On("StopCrawlerSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.StopCrawlerSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSession", func(t *testing.T) {
        input := &glue.StopSessionInput{}
        output := &glue.StopSessionOutput{}

        mockClient.On("StopSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTrigger", func(t *testing.T) {
        input := &glue.StopTriggerInput{}
        output := &glue.StopTriggerOutput{}

        mockClient.On("StopTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.StopTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopWorkflowRun", func(t *testing.T) {
        input := &glue.StopWorkflowRunInput{}
        output := &glue.StopWorkflowRunOutput{}

        mockClient.On("StopWorkflowRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopWorkflowRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &glue.TagResourceInput{}
        output := &glue.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &glue.UntagResourceInput{}
        output := &glue.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBlueprint", func(t *testing.T) {
        input := &glue.UpdateBlueprintInput{}
        output := &glue.UpdateBlueprintOutput{}

        mockClient.On("UpdateBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClassifier", func(t *testing.T) {
        input := &glue.UpdateClassifierInput{}
        output := &glue.UpdateClassifierOutput{}

        mockClient.On("UpdateClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateColumnStatisticsForPartition", func(t *testing.T) {
        input := &glue.UpdateColumnStatisticsForPartitionInput{}
        output := &glue.UpdateColumnStatisticsForPartitionOutput{}

        mockClient.On("UpdateColumnStatisticsForPartition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateColumnStatisticsForPartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateColumnStatisticsForTable", func(t *testing.T) {
        input := &glue.UpdateColumnStatisticsForTableInput{}
        output := &glue.UpdateColumnStatisticsForTableOutput{}

        mockClient.On("UpdateColumnStatisticsForTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateColumnStatisticsForTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnection", func(t *testing.T) {
        input := &glue.UpdateConnectionInput{}
        output := &glue.UpdateConnectionOutput{}

        mockClient.On("UpdateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCrawler", func(t *testing.T) {
        input := &glue.UpdateCrawlerInput{}
        output := &glue.UpdateCrawlerOutput{}

        mockClient.On("UpdateCrawler", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCrawler(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCrawlerSchedule", func(t *testing.T) {
        input := &glue.UpdateCrawlerScheduleInput{}
        output := &glue.UpdateCrawlerScheduleOutput{}

        mockClient.On("UpdateCrawlerSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCrawlerSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataQualityRuleset", func(t *testing.T) {
        input := &glue.UpdateDataQualityRulesetInput{}
        output := &glue.UpdateDataQualityRulesetOutput{}

        mockClient.On("UpdateDataQualityRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataQualityRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatabase", func(t *testing.T) {
        input := &glue.UpdateDatabaseInput{}
        output := &glue.UpdateDatabaseOutput{}

        mockClient.On("UpdateDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevEndpoint", func(t *testing.T) {
        input := &glue.UpdateDevEndpointInput{}
        output := &glue.UpdateDevEndpointOutput{}

        mockClient.On("UpdateDevEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJob", func(t *testing.T) {
        input := &glue.UpdateJobInput{}
        output := &glue.UpdateJobOutput{}

        mockClient.On("UpdateJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobFromSourceControl", func(t *testing.T) {
        input := &glue.UpdateJobFromSourceControlInput{}
        output := &glue.UpdateJobFromSourceControlOutput{}

        mockClient.On("UpdateJobFromSourceControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobFromSourceControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMLTransform", func(t *testing.T) {
        input := &glue.UpdateMLTransformInput{}
        output := &glue.UpdateMLTransformOutput{}

        mockClient.On("UpdateMLTransform", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMLTransform(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePartition", func(t *testing.T) {
        input := &glue.UpdatePartitionInput{}
        output := &glue.UpdatePartitionOutput{}

        mockClient.On("UpdatePartition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePartition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegistry", func(t *testing.T) {
        input := &glue.UpdateRegistryInput{}
        output := &glue.UpdateRegistryOutput{}

        mockClient.On("UpdateRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchema", func(t *testing.T) {
        input := &glue.UpdateSchemaInput{}
        output := &glue.UpdateSchemaOutput{}

        mockClient.On("UpdateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSourceControlFromJob", func(t *testing.T) {
        input := &glue.UpdateSourceControlFromJobInput{}
        output := &glue.UpdateSourceControlFromJobOutput{}

        mockClient.On("UpdateSourceControlFromJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSourceControlFromJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTable", func(t *testing.T) {
        input := &glue.UpdateTableInput{}
        output := &glue.UpdateTableOutput{}

        mockClient.On("UpdateTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTableOptimizer", func(t *testing.T) {
        input := &glue.UpdateTableOptimizerInput{}
        output := &glue.UpdateTableOptimizerOutput{}

        mockClient.On("UpdateTableOptimizer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTableOptimizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrigger", func(t *testing.T) {
        input := &glue.UpdateTriggerInput{}
        output := &glue.UpdateTriggerOutput{}

        mockClient.On("UpdateTrigger", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrigger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUsageProfile", func(t *testing.T) {
        input := &glue.UpdateUsageProfileInput{}
        output := &glue.UpdateUsageProfileOutput{}

        mockClient.On("UpdateUsageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUsageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserDefinedFunction", func(t *testing.T) {
        input := &glue.UpdateUserDefinedFunctionInput{}
        output := &glue.UpdateUserDefinedFunctionOutput{}

        mockClient.On("UpdateUserDefinedFunction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserDefinedFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkflow", func(t *testing.T) {
        input := &glue.UpdateWorkflowInput{}
        output := &glue.UpdateWorkflowOutput{}

        mockClient.On("UpdateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
