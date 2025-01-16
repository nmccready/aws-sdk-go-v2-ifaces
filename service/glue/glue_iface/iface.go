
package glue_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/glue"
)

// IClient defines the interface for glue
type IClient interface {
 Options() Options 
 BatchCreatePartition(ctx context.Context, params *BatchCreatePartitionInput, optFns ...func(*Options)) (*BatchCreatePartitionOutput, error) 
 BatchDeleteConnection(ctx context.Context, params *BatchDeleteConnectionInput, optFns ...func(*Options)) (*BatchDeleteConnectionOutput, error) 
 BatchDeletePartition(ctx context.Context, params *BatchDeletePartitionInput, optFns ...func(*Options)) (*BatchDeletePartitionOutput, error) 
 BatchDeleteTable(ctx context.Context, params *BatchDeleteTableInput, optFns ...func(*Options)) (*BatchDeleteTableOutput, error) 
 BatchDeleteTableVersion(ctx context.Context, params *BatchDeleteTableVersionInput, optFns ...func(*Options)) (*BatchDeleteTableVersionOutput, error) 
 BatchGetBlueprints(ctx context.Context, params *BatchGetBlueprintsInput, optFns ...func(*Options)) (*BatchGetBlueprintsOutput, error) 
 BatchGetCrawlers(ctx context.Context, params *BatchGetCrawlersInput, optFns ...func(*Options)) (*BatchGetCrawlersOutput, error) 
 BatchGetCustomEntityTypes(ctx context.Context, params *BatchGetCustomEntityTypesInput, optFns ...func(*Options)) (*BatchGetCustomEntityTypesOutput, error) 
 BatchGetDataQualityResult(ctx context.Context, params *BatchGetDataQualityResultInput, optFns ...func(*Options)) (*BatchGetDataQualityResultOutput, error) 
 BatchGetDevEndpoints(ctx context.Context, params *BatchGetDevEndpointsInput, optFns ...func(*Options)) (*BatchGetDevEndpointsOutput, error) 
 BatchGetJobs(ctx context.Context, params *BatchGetJobsInput, optFns ...func(*Options)) (*BatchGetJobsOutput, error) 
 BatchGetPartition(ctx context.Context, params *BatchGetPartitionInput, optFns ...func(*Options)) (*BatchGetPartitionOutput, error) 
 BatchGetTableOptimizer(ctx context.Context, params *BatchGetTableOptimizerInput, optFns ...func(*Options)) (*BatchGetTableOptimizerOutput, error) 
 BatchGetTriggers(ctx context.Context, params *BatchGetTriggersInput, optFns ...func(*Options)) (*BatchGetTriggersOutput, error) 
 BatchGetWorkflows(ctx context.Context, params *BatchGetWorkflowsInput, optFns ...func(*Options)) (*BatchGetWorkflowsOutput, error) 
 BatchPutDataQualityStatisticAnnotation(ctx context.Context, params *BatchPutDataQualityStatisticAnnotationInput, optFns ...func(*Options)) (*BatchPutDataQualityStatisticAnnotationOutput, error) 
 BatchStopJobRun(ctx context.Context, params *BatchStopJobRunInput, optFns ...func(*Options)) (*BatchStopJobRunOutput, error) 
 BatchUpdatePartition(ctx context.Context, params *BatchUpdatePartitionInput, optFns ...func(*Options)) (*BatchUpdatePartitionOutput, error) 
 CancelDataQualityRuleRecommendationRun(ctx context.Context, params *CancelDataQualityRuleRecommendationRunInput, optFns ...func(*Options)) (*CancelDataQualityRuleRecommendationRunOutput, error) 
 CancelDataQualityRulesetEvaluationRun(ctx context.Context, params *CancelDataQualityRulesetEvaluationRunInput, optFns ...func(*Options)) (*CancelDataQualityRulesetEvaluationRunOutput, error) 
 CancelMLTaskRun(ctx context.Context, params *CancelMLTaskRunInput, optFns ...func(*Options)) (*CancelMLTaskRunOutput, error) 
 CancelStatement(ctx context.Context, params *CancelStatementInput, optFns ...func(*Options)) (*CancelStatementOutput, error) 
 CheckSchemaVersionValidity(ctx context.Context, params *CheckSchemaVersionValidityInput, optFns ...func(*Options)) (*CheckSchemaVersionValidityOutput, error) 
 CreateBlueprint(ctx context.Context, params *CreateBlueprintInput, optFns ...func(*Options)) (*CreateBlueprintOutput, error) 
 CreateCatalog(ctx context.Context, params *CreateCatalogInput, optFns ...func(*Options)) (*CreateCatalogOutput, error) 
 CreateClassifier(ctx context.Context, params *CreateClassifierInput, optFns ...func(*Options)) (*CreateClassifierOutput, error) 
 CreateColumnStatisticsTaskSettings(ctx context.Context, params *CreateColumnStatisticsTaskSettingsInput, optFns ...func(*Options)) (*CreateColumnStatisticsTaskSettingsOutput, error) 
 CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) 
 CreateCrawler(ctx context.Context, params *CreateCrawlerInput, optFns ...func(*Options)) (*CreateCrawlerOutput, error) 
 CreateCustomEntityType(ctx context.Context, params *CreateCustomEntityTypeInput, optFns ...func(*Options)) (*CreateCustomEntityTypeOutput, error) 
 CreateDataQualityRuleset(ctx context.Context, params *CreateDataQualityRulesetInput, optFns ...func(*Options)) (*CreateDataQualityRulesetOutput, error) 
 CreateDatabase(ctx context.Context, params *CreateDatabaseInput, optFns ...func(*Options)) (*CreateDatabaseOutput, error) 
 CreateDevEndpoint(ctx context.Context, params *CreateDevEndpointInput, optFns ...func(*Options)) (*CreateDevEndpointOutput, error) 
 CreateIntegration(ctx context.Context, params *CreateIntegrationInput, optFns ...func(*Options)) (*CreateIntegrationOutput, error) 
 CreateIntegrationResourceProperty(ctx context.Context, params *CreateIntegrationResourcePropertyInput, optFns ...func(*Options)) (*CreateIntegrationResourcePropertyOutput, error) 
 CreateIntegrationTableProperties(ctx context.Context, params *CreateIntegrationTablePropertiesInput, optFns ...func(*Options)) (*CreateIntegrationTablePropertiesOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateMLTransform(ctx context.Context, params *CreateMLTransformInput, optFns ...func(*Options)) (*CreateMLTransformOutput, error) 
 CreatePartition(ctx context.Context, params *CreatePartitionInput, optFns ...func(*Options)) (*CreatePartitionOutput, error) 
 CreatePartitionIndex(ctx context.Context, params *CreatePartitionIndexInput, optFns ...func(*Options)) (*CreatePartitionIndexOutput, error) 
 CreateRegistry(ctx context.Context, params *CreateRegistryInput, optFns ...func(*Options)) (*CreateRegistryOutput, error) 
 CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) 
 CreateScript(ctx context.Context, params *CreateScriptInput, optFns ...func(*Options)) (*CreateScriptOutput, error) 
 CreateSecurityConfiguration(ctx context.Context, params *CreateSecurityConfigurationInput, optFns ...func(*Options)) (*CreateSecurityConfigurationOutput, error) 
 CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) 
 CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) 
 CreateTableOptimizer(ctx context.Context, params *CreateTableOptimizerInput, optFns ...func(*Options)) (*CreateTableOptimizerOutput, error) 
 CreateTrigger(ctx context.Context, params *CreateTriggerInput, optFns ...func(*Options)) (*CreateTriggerOutput, error) 
 CreateUsageProfile(ctx context.Context, params *CreateUsageProfileInput, optFns ...func(*Options)) (*CreateUsageProfileOutput, error) 
 CreateUserDefinedFunction(ctx context.Context, params *CreateUserDefinedFunctionInput, optFns ...func(*Options)) (*CreateUserDefinedFunctionOutput, error) 
 CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) 
 DeleteBlueprint(ctx context.Context, params *DeleteBlueprintInput, optFns ...func(*Options)) (*DeleteBlueprintOutput, error) 
 DeleteCatalog(ctx context.Context, params *DeleteCatalogInput, optFns ...func(*Options)) (*DeleteCatalogOutput, error) 
 DeleteClassifier(ctx context.Context, params *DeleteClassifierInput, optFns ...func(*Options)) (*DeleteClassifierOutput, error) 
 DeleteColumnStatisticsForPartition(ctx context.Context, params *DeleteColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForPartitionOutput, error) 
 DeleteColumnStatisticsForTable(ctx context.Context, params *DeleteColumnStatisticsForTableInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForTableOutput, error) 
 DeleteColumnStatisticsTaskSettings(ctx context.Context, params *DeleteColumnStatisticsTaskSettingsInput, optFns ...func(*Options)) (*DeleteColumnStatisticsTaskSettingsOutput, error) 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 DeleteCrawler(ctx context.Context, params *DeleteCrawlerInput, optFns ...func(*Options)) (*DeleteCrawlerOutput, error) 
 DeleteCustomEntityType(ctx context.Context, params *DeleteCustomEntityTypeInput, optFns ...func(*Options)) (*DeleteCustomEntityTypeOutput, error) 
 DeleteDataQualityRuleset(ctx context.Context, params *DeleteDataQualityRulesetInput, optFns ...func(*Options)) (*DeleteDataQualityRulesetOutput, error) 
 DeleteDatabase(ctx context.Context, params *DeleteDatabaseInput, optFns ...func(*Options)) (*DeleteDatabaseOutput, error) 
 DeleteDevEndpoint(ctx context.Context, params *DeleteDevEndpointInput, optFns ...func(*Options)) (*DeleteDevEndpointOutput, error) 
 DeleteIntegration(ctx context.Context, params *DeleteIntegrationInput, optFns ...func(*Options)) (*DeleteIntegrationOutput, error) 
 DeleteIntegrationTableProperties(ctx context.Context, params *DeleteIntegrationTablePropertiesInput, optFns ...func(*Options)) (*DeleteIntegrationTablePropertiesOutput, error) 
 DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) 
 DeleteMLTransform(ctx context.Context, params *DeleteMLTransformInput, optFns ...func(*Options)) (*DeleteMLTransformOutput, error) 
 DeletePartition(ctx context.Context, params *DeletePartitionInput, optFns ...func(*Options)) (*DeletePartitionOutput, error) 
 DeletePartitionIndex(ctx context.Context, params *DeletePartitionIndexInput, optFns ...func(*Options)) (*DeletePartitionIndexOutput, error) 
 DeleteRegistry(ctx context.Context, params *DeleteRegistryInput, optFns ...func(*Options)) (*DeleteRegistryOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteSchema(ctx context.Context, params *DeleteSchemaInput, optFns ...func(*Options)) (*DeleteSchemaOutput, error) 
 DeleteSchemaVersions(ctx context.Context, params *DeleteSchemaVersionsInput, optFns ...func(*Options)) (*DeleteSchemaVersionsOutput, error) 
 DeleteSecurityConfiguration(ctx context.Context, params *DeleteSecurityConfigurationInput, optFns ...func(*Options)) (*DeleteSecurityConfigurationOutput, error) 
 DeleteSession(ctx context.Context, params *DeleteSessionInput, optFns ...func(*Options)) (*DeleteSessionOutput, error) 
 DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) 
 DeleteTableOptimizer(ctx context.Context, params *DeleteTableOptimizerInput, optFns ...func(*Options)) (*DeleteTableOptimizerOutput, error) 
 DeleteTableVersion(ctx context.Context, params *DeleteTableVersionInput, optFns ...func(*Options)) (*DeleteTableVersionOutput, error) 
 DeleteTrigger(ctx context.Context, params *DeleteTriggerInput, optFns ...func(*Options)) (*DeleteTriggerOutput, error) 
 DeleteUsageProfile(ctx context.Context, params *DeleteUsageProfileInput, optFns ...func(*Options)) (*DeleteUsageProfileOutput, error) 
 DeleteUserDefinedFunction(ctx context.Context, params *DeleteUserDefinedFunctionInput, optFns ...func(*Options)) (*DeleteUserDefinedFunctionOutput, error) 
 DeleteWorkflow(ctx context.Context, params *DeleteWorkflowInput, optFns ...func(*Options)) (*DeleteWorkflowOutput, error) 
 DescribeConnectionType(ctx context.Context, params *DescribeConnectionTypeInput, optFns ...func(*Options)) (*DescribeConnectionTypeOutput, error) 
 DescribeEntity(ctx context.Context, params *DescribeEntityInput, optFns ...func(*Options)) (*DescribeEntityOutput, error) 
 DescribeInboundIntegrations(ctx context.Context, params *DescribeInboundIntegrationsInput, optFns ...func(*Options)) (*DescribeInboundIntegrationsOutput, error) 
 DescribeIntegrations(ctx context.Context, params *DescribeIntegrationsInput, optFns ...func(*Options)) (*DescribeIntegrationsOutput, error) 
 GetBlueprint(ctx context.Context, params *GetBlueprintInput, optFns ...func(*Options)) (*GetBlueprintOutput, error) 
 GetBlueprintRun(ctx context.Context, params *GetBlueprintRunInput, optFns ...func(*Options)) (*GetBlueprintRunOutput, error) 
 GetBlueprintRuns(ctx context.Context, params *GetBlueprintRunsInput, optFns ...func(*Options)) (*GetBlueprintRunsOutput, error) 
 GetCatalog(ctx context.Context, params *GetCatalogInput, optFns ...func(*Options)) (*GetCatalogOutput, error) 
 GetCatalogImportStatus(ctx context.Context, params *GetCatalogImportStatusInput, optFns ...func(*Options)) (*GetCatalogImportStatusOutput, error) 
 GetCatalogs(ctx context.Context, params *GetCatalogsInput, optFns ...func(*Options)) (*GetCatalogsOutput, error) 
 GetClassifier(ctx context.Context, params *GetClassifierInput, optFns ...func(*Options)) (*GetClassifierOutput, error) 
 GetClassifiers(ctx context.Context, params *GetClassifiersInput, optFns ...func(*Options)) (*GetClassifiersOutput, error) 
 GetColumnStatisticsForPartition(ctx context.Context, params *GetColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*GetColumnStatisticsForPartitionOutput, error) 
 GetColumnStatisticsForTable(ctx context.Context, params *GetColumnStatisticsForTableInput, optFns ...func(*Options)) (*GetColumnStatisticsForTableOutput, error) 
 GetColumnStatisticsTaskRun(ctx context.Context, params *GetColumnStatisticsTaskRunInput, optFns ...func(*Options)) (*GetColumnStatisticsTaskRunOutput, error) 
 GetColumnStatisticsTaskRuns(ctx context.Context, params *GetColumnStatisticsTaskRunsInput, optFns ...func(*Options)) (*GetColumnStatisticsTaskRunsOutput, error) 
 GetColumnStatisticsTaskSettings(ctx context.Context, params *GetColumnStatisticsTaskSettingsInput, optFns ...func(*Options)) (*GetColumnStatisticsTaskSettingsOutput, error) 
 GetConnection(ctx context.Context, params *GetConnectionInput, optFns ...func(*Options)) (*GetConnectionOutput, error) 
 GetConnections(ctx context.Context, params *GetConnectionsInput, optFns ...func(*Options)) (*GetConnectionsOutput, error) 
 GetCrawler(ctx context.Context, params *GetCrawlerInput, optFns ...func(*Options)) (*GetCrawlerOutput, error) 
 GetCrawlerMetrics(ctx context.Context, params *GetCrawlerMetricsInput, optFns ...func(*Options)) (*GetCrawlerMetricsOutput, error) 
 GetCrawlers(ctx context.Context, params *GetCrawlersInput, optFns ...func(*Options)) (*GetCrawlersOutput, error) 
 GetCustomEntityType(ctx context.Context, params *GetCustomEntityTypeInput, optFns ...func(*Options)) (*GetCustomEntityTypeOutput, error) 
 GetDataCatalogEncryptionSettings(ctx context.Context, params *GetDataCatalogEncryptionSettingsInput, optFns ...func(*Options)) (*GetDataCatalogEncryptionSettingsOutput, error) 
 GetDataQualityModel(ctx context.Context, params *GetDataQualityModelInput, optFns ...func(*Options)) (*GetDataQualityModelOutput, error) 
 GetDataQualityModelResult(ctx context.Context, params *GetDataQualityModelResultInput, optFns ...func(*Options)) (*GetDataQualityModelResultOutput, error) 
 GetDataQualityResult(ctx context.Context, params *GetDataQualityResultInput, optFns ...func(*Options)) (*GetDataQualityResultOutput, error) 
 GetDataQualityRuleRecommendationRun(ctx context.Context, params *GetDataQualityRuleRecommendationRunInput, optFns ...func(*Options)) (*GetDataQualityRuleRecommendationRunOutput, error) 
 GetDataQualityRuleset(ctx context.Context, params *GetDataQualityRulesetInput, optFns ...func(*Options)) (*GetDataQualityRulesetOutput, error) 
 GetDataQualityRulesetEvaluationRun(ctx context.Context, params *GetDataQualityRulesetEvaluationRunInput, optFns ...func(*Options)) (*GetDataQualityRulesetEvaluationRunOutput, error) 
 GetDatabase(ctx context.Context, params *GetDatabaseInput, optFns ...func(*Options)) (*GetDatabaseOutput, error) 
 GetDatabases(ctx context.Context, params *GetDatabasesInput, optFns ...func(*Options)) (*GetDatabasesOutput, error) 
 GetDataflowGraph(ctx context.Context, params *GetDataflowGraphInput, optFns ...func(*Options)) (*GetDataflowGraphOutput, error) 
 GetDevEndpoint(ctx context.Context, params *GetDevEndpointInput, optFns ...func(*Options)) (*GetDevEndpointOutput, error) 
 GetDevEndpoints(ctx context.Context, params *GetDevEndpointsInput, optFns ...func(*Options)) (*GetDevEndpointsOutput, error) 
 GetEntityRecords(ctx context.Context, params *GetEntityRecordsInput, optFns ...func(*Options)) (*GetEntityRecordsOutput, error) 
 GetIntegrationResourceProperty(ctx context.Context, params *GetIntegrationResourcePropertyInput, optFns ...func(*Options)) (*GetIntegrationResourcePropertyOutput, error) 
 GetIntegrationTableProperties(ctx context.Context, params *GetIntegrationTablePropertiesInput, optFns ...func(*Options)) (*GetIntegrationTablePropertiesOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetJobBookmark(ctx context.Context, params *GetJobBookmarkInput, optFns ...func(*Options)) (*GetJobBookmarkOutput, error) 
 GetJobRun(ctx context.Context, params *GetJobRunInput, optFns ...func(*Options)) (*GetJobRunOutput, error) 
 GetJobRuns(ctx context.Context, params *GetJobRunsInput, optFns ...func(*Options)) (*GetJobRunsOutput, error) 
 GetJobs(ctx context.Context, params *GetJobsInput, optFns ...func(*Options)) (*GetJobsOutput, error) 
 GetMLTaskRun(ctx context.Context, params *GetMLTaskRunInput, optFns ...func(*Options)) (*GetMLTaskRunOutput, error) 
 GetMLTaskRuns(ctx context.Context, params *GetMLTaskRunsInput, optFns ...func(*Options)) (*GetMLTaskRunsOutput, error) 
 GetMLTransform(ctx context.Context, params *GetMLTransformInput, optFns ...func(*Options)) (*GetMLTransformOutput, error) 
 GetMLTransforms(ctx context.Context, params *GetMLTransformsInput, optFns ...func(*Options)) (*GetMLTransformsOutput, error) 
 GetMapping(ctx context.Context, params *GetMappingInput, optFns ...func(*Options)) (*GetMappingOutput, error) 
 GetPartition(ctx context.Context, params *GetPartitionInput, optFns ...func(*Options)) (*GetPartitionOutput, error) 
 GetPartitionIndexes(ctx context.Context, params *GetPartitionIndexesInput, optFns ...func(*Options)) (*GetPartitionIndexesOutput, error) 
 GetPartitions(ctx context.Context, params *GetPartitionsInput, optFns ...func(*Options)) (*GetPartitionsOutput, error) 
 GetPlan(ctx context.Context, params *GetPlanInput, optFns ...func(*Options)) (*GetPlanOutput, error) 
 GetRegistry(ctx context.Context, params *GetRegistryInput, optFns ...func(*Options)) (*GetRegistryOutput, error) 
 GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 GetSchema(ctx context.Context, params *GetSchemaInput, optFns ...func(*Options)) (*GetSchemaOutput, error) 
 GetSchemaByDefinition(ctx context.Context, params *GetSchemaByDefinitionInput, optFns ...func(*Options)) (*GetSchemaByDefinitionOutput, error) 
 GetSchemaVersion(ctx context.Context, params *GetSchemaVersionInput, optFns ...func(*Options)) (*GetSchemaVersionOutput, error) 
 GetSchemaVersionsDiff(ctx context.Context, params *GetSchemaVersionsDiffInput, optFns ...func(*Options)) (*GetSchemaVersionsDiffOutput, error) 
 GetSecurityConfiguration(ctx context.Context, params *GetSecurityConfigurationInput, optFns ...func(*Options)) (*GetSecurityConfigurationOutput, error) 
 GetSecurityConfigurations(ctx context.Context, params *GetSecurityConfigurationsInput, optFns ...func(*Options)) (*GetSecurityConfigurationsOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 GetStatement(ctx context.Context, params *GetStatementInput, optFns ...func(*Options)) (*GetStatementOutput, error) 
 GetTable(ctx context.Context, params *GetTableInput, optFns ...func(*Options)) (*GetTableOutput, error) 
 GetTableOptimizer(ctx context.Context, params *GetTableOptimizerInput, optFns ...func(*Options)) (*GetTableOptimizerOutput, error) 
 GetTableVersion(ctx context.Context, params *GetTableVersionInput, optFns ...func(*Options)) (*GetTableVersionOutput, error) 
 GetTableVersions(ctx context.Context, params *GetTableVersionsInput, optFns ...func(*Options)) (*GetTableVersionsOutput, error) 
 GetTables(ctx context.Context, params *GetTablesInput, optFns ...func(*Options)) (*GetTablesOutput, error) 
 GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) 
 GetTrigger(ctx context.Context, params *GetTriggerInput, optFns ...func(*Options)) (*GetTriggerOutput, error) 
 GetTriggers(ctx context.Context, params *GetTriggersInput, optFns ...func(*Options)) (*GetTriggersOutput, error) 
 GetUnfilteredPartitionMetadata(ctx context.Context, params *GetUnfilteredPartitionMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionMetadataOutput, error) 
 GetUnfilteredPartitionsMetadata(ctx context.Context, params *GetUnfilteredPartitionsMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error) 
 GetUnfilteredTableMetadata(ctx context.Context, params *GetUnfilteredTableMetadataInput, optFns ...func(*Options)) (*GetUnfilteredTableMetadataOutput, error) 
 GetUsageProfile(ctx context.Context, params *GetUsageProfileInput, optFns ...func(*Options)) (*GetUsageProfileOutput, error) 
 GetUserDefinedFunction(ctx context.Context, params *GetUserDefinedFunctionInput, optFns ...func(*Options)) (*GetUserDefinedFunctionOutput, error) 
 GetUserDefinedFunctions(ctx context.Context, params *GetUserDefinedFunctionsInput, optFns ...func(*Options)) (*GetUserDefinedFunctionsOutput, error) 
 GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) 
 GetWorkflowRun(ctx context.Context, params *GetWorkflowRunInput, optFns ...func(*Options)) (*GetWorkflowRunOutput, error) 
 GetWorkflowRunProperties(ctx context.Context, params *GetWorkflowRunPropertiesInput, optFns ...func(*Options)) (*GetWorkflowRunPropertiesOutput, error) 
 GetWorkflowRuns(ctx context.Context, params *GetWorkflowRunsInput, optFns ...func(*Options)) (*GetWorkflowRunsOutput, error) 
 ImportCatalogToGlue(ctx context.Context, params *ImportCatalogToGlueInput, optFns ...func(*Options)) (*ImportCatalogToGlueOutput, error) 
 ListBlueprints(ctx context.Context, params *ListBlueprintsInput, optFns ...func(*Options)) (*ListBlueprintsOutput, error) 
 ListColumnStatisticsTaskRuns(ctx context.Context, params *ListColumnStatisticsTaskRunsInput, optFns ...func(*Options)) (*ListColumnStatisticsTaskRunsOutput, error) 
 ListConnectionTypes(ctx context.Context, params *ListConnectionTypesInput, optFns ...func(*Options)) (*ListConnectionTypesOutput, error) 
 ListCrawlers(ctx context.Context, params *ListCrawlersInput, optFns ...func(*Options)) (*ListCrawlersOutput, error) 
 ListCrawls(ctx context.Context, params *ListCrawlsInput, optFns ...func(*Options)) (*ListCrawlsOutput, error) 
 ListCustomEntityTypes(ctx context.Context, params *ListCustomEntityTypesInput, optFns ...func(*Options)) (*ListCustomEntityTypesOutput, error) 
 ListDataQualityResults(ctx context.Context, params *ListDataQualityResultsInput, optFns ...func(*Options)) (*ListDataQualityResultsOutput, error) 
 ListDataQualityRuleRecommendationRuns(ctx context.Context, params *ListDataQualityRuleRecommendationRunsInput, optFns ...func(*Options)) (*ListDataQualityRuleRecommendationRunsOutput, error) 
 ListDataQualityRulesetEvaluationRuns(ctx context.Context, params *ListDataQualityRulesetEvaluationRunsInput, optFns ...func(*Options)) (*ListDataQualityRulesetEvaluationRunsOutput, error) 
 ListDataQualityRulesets(ctx context.Context, params *ListDataQualityRulesetsInput, optFns ...func(*Options)) (*ListDataQualityRulesetsOutput, error) 
 ListDataQualityStatisticAnnotations(ctx context.Context, params *ListDataQualityStatisticAnnotationsInput, optFns ...func(*Options)) (*ListDataQualityStatisticAnnotationsOutput, error) 
 ListDataQualityStatistics(ctx context.Context, params *ListDataQualityStatisticsInput, optFns ...func(*Options)) (*ListDataQualityStatisticsOutput, error) 
 ListDevEndpoints(ctx context.Context, params *ListDevEndpointsInput, optFns ...func(*Options)) (*ListDevEndpointsOutput, error) 
 ListEntities(ctx context.Context, params *ListEntitiesInput, optFns ...func(*Options)) (*ListEntitiesOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListMLTransforms(ctx context.Context, params *ListMLTransformsInput, optFns ...func(*Options)) (*ListMLTransformsOutput, error) 
 ListRegistries(ctx context.Context, params *ListRegistriesInput, optFns ...func(*Options)) (*ListRegistriesOutput, error) 
 ListSchemaVersions(ctx context.Context, params *ListSchemaVersionsInput, optFns ...func(*Options)) (*ListSchemaVersionsOutput, error) 
 ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) 
 ListSessions(ctx context.Context, params *ListSessionsInput, optFns ...func(*Options)) (*ListSessionsOutput, error) 
 ListStatements(ctx context.Context, params *ListStatementsInput, optFns ...func(*Options)) (*ListStatementsOutput, error) 
 ListTableOptimizerRuns(ctx context.Context, params *ListTableOptimizerRunsInput, optFns ...func(*Options)) (*ListTableOptimizerRunsOutput, error) 
 ListTriggers(ctx context.Context, params *ListTriggersInput, optFns ...func(*Options)) (*ListTriggersOutput, error) 
 ListUsageProfiles(ctx context.Context, params *ListUsageProfilesInput, optFns ...func(*Options)) (*ListUsageProfilesOutput, error) 
 ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) 
 ModifyIntegration(ctx context.Context, params *ModifyIntegrationInput, optFns ...func(*Options)) (*ModifyIntegrationOutput, error) 
 PutDataCatalogEncryptionSettings(ctx context.Context, params *PutDataCatalogEncryptionSettingsInput, optFns ...func(*Options)) (*PutDataCatalogEncryptionSettingsOutput, error) 
 PutDataQualityProfileAnnotation(ctx context.Context, params *PutDataQualityProfileAnnotationInput, optFns ...func(*Options)) (*PutDataQualityProfileAnnotationOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 PutSchemaVersionMetadata(ctx context.Context, params *PutSchemaVersionMetadataInput, optFns ...func(*Options)) (*PutSchemaVersionMetadataOutput, error) 
 PutWorkflowRunProperties(ctx context.Context, params *PutWorkflowRunPropertiesInput, optFns ...func(*Options)) (*PutWorkflowRunPropertiesOutput, error) 
 QuerySchemaVersionMetadata(ctx context.Context, params *QuerySchemaVersionMetadataInput, optFns ...func(*Options)) (*QuerySchemaVersionMetadataOutput, error) 
 RegisterSchemaVersion(ctx context.Context, params *RegisterSchemaVersionInput, optFns ...func(*Options)) (*RegisterSchemaVersionOutput, error) 
 RemoveSchemaVersionMetadata(ctx context.Context, params *RemoveSchemaVersionMetadataInput, optFns ...func(*Options)) (*RemoveSchemaVersionMetadataOutput, error) 
 ResetJobBookmark(ctx context.Context, params *ResetJobBookmarkInput, optFns ...func(*Options)) (*ResetJobBookmarkOutput, error) 
 ResumeWorkflowRun(ctx context.Context, params *ResumeWorkflowRunInput, optFns ...func(*Options)) (*ResumeWorkflowRunOutput, error) 
 RunStatement(ctx context.Context, params *RunStatementInput, optFns ...func(*Options)) (*RunStatementOutput, error) 
 SearchTables(ctx context.Context, params *SearchTablesInput, optFns ...func(*Options)) (*SearchTablesOutput, error) 
 StartBlueprintRun(ctx context.Context, params *StartBlueprintRunInput, optFns ...func(*Options)) (*StartBlueprintRunOutput, error) 
 StartColumnStatisticsTaskRun(ctx context.Context, params *StartColumnStatisticsTaskRunInput, optFns ...func(*Options)) (*StartColumnStatisticsTaskRunOutput, error) 
 StartColumnStatisticsTaskRunSchedule(ctx context.Context, params *StartColumnStatisticsTaskRunScheduleInput, optFns ...func(*Options)) (*StartColumnStatisticsTaskRunScheduleOutput, error) 
 StartCrawler(ctx context.Context, params *StartCrawlerInput, optFns ...func(*Options)) (*StartCrawlerOutput, error) 
 StartCrawlerSchedule(ctx context.Context, params *StartCrawlerScheduleInput, optFns ...func(*Options)) (*StartCrawlerScheduleOutput, error) 
 StartDataQualityRuleRecommendationRun(ctx context.Context, params *StartDataQualityRuleRecommendationRunInput, optFns ...func(*Options)) (*StartDataQualityRuleRecommendationRunOutput, error) 
 StartDataQualityRulesetEvaluationRun(ctx context.Context, params *StartDataQualityRulesetEvaluationRunInput, optFns ...func(*Options)) (*StartDataQualityRulesetEvaluationRunOutput, error) 
 StartExportLabelsTaskRun(ctx context.Context, params *StartExportLabelsTaskRunInput, optFns ...func(*Options)) (*StartExportLabelsTaskRunOutput, error) 
 StartImportLabelsTaskRun(ctx context.Context, params *StartImportLabelsTaskRunInput, optFns ...func(*Options)) (*StartImportLabelsTaskRunOutput, error) 
 StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) 
 StartMLEvaluationTaskRun(ctx context.Context, params *StartMLEvaluationTaskRunInput, optFns ...func(*Options)) (*StartMLEvaluationTaskRunOutput, error) 
 StartMLLabelingSetGenerationTaskRun(ctx context.Context, params *StartMLLabelingSetGenerationTaskRunInput, optFns ...func(*Options)) (*StartMLLabelingSetGenerationTaskRunOutput, error) 
 StartTrigger(ctx context.Context, params *StartTriggerInput, optFns ...func(*Options)) (*StartTriggerOutput, error) 
 StartWorkflowRun(ctx context.Context, params *StartWorkflowRunInput, optFns ...func(*Options)) (*StartWorkflowRunOutput, error) 
 StopColumnStatisticsTaskRun(ctx context.Context, params *StopColumnStatisticsTaskRunInput, optFns ...func(*Options)) (*StopColumnStatisticsTaskRunOutput, error) 
 StopColumnStatisticsTaskRunSchedule(ctx context.Context, params *StopColumnStatisticsTaskRunScheduleInput, optFns ...func(*Options)) (*StopColumnStatisticsTaskRunScheduleOutput, error) 
 StopCrawler(ctx context.Context, params *StopCrawlerInput, optFns ...func(*Options)) (*StopCrawlerOutput, error) 
 StopCrawlerSchedule(ctx context.Context, params *StopCrawlerScheduleInput, optFns ...func(*Options)) (*StopCrawlerScheduleOutput, error) 
 StopSession(ctx context.Context, params *StopSessionInput, optFns ...func(*Options)) (*StopSessionOutput, error) 
 StopTrigger(ctx context.Context, params *StopTriggerInput, optFns ...func(*Options)) (*StopTriggerOutput, error) 
 StopWorkflowRun(ctx context.Context, params *StopWorkflowRunInput, optFns ...func(*Options)) (*StopWorkflowRunOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestConnection(ctx context.Context, params *TestConnectionInput, optFns ...func(*Options)) (*TestConnectionOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBlueprint(ctx context.Context, params *UpdateBlueprintInput, optFns ...func(*Options)) (*UpdateBlueprintOutput, error) 
 UpdateCatalog(ctx context.Context, params *UpdateCatalogInput, optFns ...func(*Options)) (*UpdateCatalogOutput, error) 
 UpdateClassifier(ctx context.Context, params *UpdateClassifierInput, optFns ...func(*Options)) (*UpdateClassifierOutput, error) 
 UpdateColumnStatisticsForPartition(ctx context.Context, params *UpdateColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*UpdateColumnStatisticsForPartitionOutput, error) 
 UpdateColumnStatisticsForTable(ctx context.Context, params *UpdateColumnStatisticsForTableInput, optFns ...func(*Options)) (*UpdateColumnStatisticsForTableOutput, error) 
 UpdateColumnStatisticsTaskSettings(ctx context.Context, params *UpdateColumnStatisticsTaskSettingsInput, optFns ...func(*Options)) (*UpdateColumnStatisticsTaskSettingsOutput, error) 
 UpdateConnection(ctx context.Context, params *UpdateConnectionInput, optFns ...func(*Options)) (*UpdateConnectionOutput, error) 
 UpdateCrawler(ctx context.Context, params *UpdateCrawlerInput, optFns ...func(*Options)) (*UpdateCrawlerOutput, error) 
 UpdateCrawlerSchedule(ctx context.Context, params *UpdateCrawlerScheduleInput, optFns ...func(*Options)) (*UpdateCrawlerScheduleOutput, error) 
 UpdateDataQualityRuleset(ctx context.Context, params *UpdateDataQualityRulesetInput, optFns ...func(*Options)) (*UpdateDataQualityRulesetOutput, error) 
 UpdateDatabase(ctx context.Context, params *UpdateDatabaseInput, optFns ...func(*Options)) (*UpdateDatabaseOutput, error) 
 UpdateDevEndpoint(ctx context.Context, params *UpdateDevEndpointInput, optFns ...func(*Options)) (*UpdateDevEndpointOutput, error) 
 UpdateIntegrationResourceProperty(ctx context.Context, params *UpdateIntegrationResourcePropertyInput, optFns ...func(*Options)) (*UpdateIntegrationResourcePropertyOutput, error) 
 UpdateIntegrationTableProperties(ctx context.Context, params *UpdateIntegrationTablePropertiesInput, optFns ...func(*Options)) (*UpdateIntegrationTablePropertiesOutput, error) 
 UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) 
 UpdateJobFromSourceControl(ctx context.Context, params *UpdateJobFromSourceControlInput, optFns ...func(*Options)) (*UpdateJobFromSourceControlOutput, error) 
 UpdateMLTransform(ctx context.Context, params *UpdateMLTransformInput, optFns ...func(*Options)) (*UpdateMLTransformOutput, error) 
 UpdatePartition(ctx context.Context, params *UpdatePartitionInput, optFns ...func(*Options)) (*UpdatePartitionOutput, error) 
 UpdateRegistry(ctx context.Context, params *UpdateRegistryInput, optFns ...func(*Options)) (*UpdateRegistryOutput, error) 
 UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) 
 UpdateSourceControlFromJob(ctx context.Context, params *UpdateSourceControlFromJobInput, optFns ...func(*Options)) (*UpdateSourceControlFromJobOutput, error) 
 UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) 
 UpdateTableOptimizer(ctx context.Context, params *UpdateTableOptimizerInput, optFns ...func(*Options)) (*UpdateTableOptimizerOutput, error) 
 UpdateTrigger(ctx context.Context, params *UpdateTriggerInput, optFns ...func(*Options)) (*UpdateTriggerOutput, error) 
 UpdateUsageProfile(ctx context.Context, params *UpdateUsageProfileInput, optFns ...func(*Options)) (*UpdateUsageProfileOutput, error) 
 UpdateUserDefinedFunction(ctx context.Context, params *UpdateUserDefinedFunctionInput, optFns ...func(*Options)) (*UpdateUserDefinedFunctionOutput, error) 
 UpdateWorkflow(ctx context.Context, params *UpdateWorkflowInput, optFns ...func(*Options)) (*UpdateWorkflowOutput, error) 
}
