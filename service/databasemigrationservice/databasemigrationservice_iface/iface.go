
package databasemigrationservice_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

// IClient defines the interface for databasemigrationservice
type IClient interface {
 Options() Options 
 AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) 
 ApplyPendingMaintenanceAction(ctx context.Context, params *ApplyPendingMaintenanceActionInput, optFns ...func(*Options)) (*ApplyPendingMaintenanceActionOutput, error) 
 BatchStartRecommendations(ctx context.Context, params *BatchStartRecommendationsInput, optFns ...func(*Options)) (*BatchStartRecommendationsOutput, error) 
 CancelReplicationTaskAssessmentRun(ctx context.Context, params *CancelReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*CancelReplicationTaskAssessmentRunOutput, error) 
 CreateDataMigration(ctx context.Context, params *CreateDataMigrationInput, optFns ...func(*Options)) (*CreateDataMigrationOutput, error) 
 CreateDataProvider(ctx context.Context, params *CreateDataProviderInput, optFns ...func(*Options)) (*CreateDataProviderOutput, error) 
 CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) 
 CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) 
 CreateFleetAdvisorCollector(ctx context.Context, params *CreateFleetAdvisorCollectorInput, optFns ...func(*Options)) (*CreateFleetAdvisorCollectorOutput, error) 
 CreateInstanceProfile(ctx context.Context, params *CreateInstanceProfileInput, optFns ...func(*Options)) (*CreateInstanceProfileOutput, error) 
 CreateMigrationProject(ctx context.Context, params *CreateMigrationProjectInput, optFns ...func(*Options)) (*CreateMigrationProjectOutput, error) 
 CreateReplicationConfig(ctx context.Context, params *CreateReplicationConfigInput, optFns ...func(*Options)) (*CreateReplicationConfigOutput, error) 
 CreateReplicationInstance(ctx context.Context, params *CreateReplicationInstanceInput, optFns ...func(*Options)) (*CreateReplicationInstanceOutput, error) 
 CreateReplicationSubnetGroup(ctx context.Context, params *CreateReplicationSubnetGroupInput, optFns ...func(*Options)) (*CreateReplicationSubnetGroupOutput, error) 
 CreateReplicationTask(ctx context.Context, params *CreateReplicationTaskInput, optFns ...func(*Options)) (*CreateReplicationTaskOutput, error) 
 DeleteCertificate(ctx context.Context, params *DeleteCertificateInput, optFns ...func(*Options)) (*DeleteCertificateOutput, error) 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 DeleteDataMigration(ctx context.Context, params *DeleteDataMigrationInput, optFns ...func(*Options)) (*DeleteDataMigrationOutput, error) 
 DeleteDataProvider(ctx context.Context, params *DeleteDataProviderInput, optFns ...func(*Options)) (*DeleteDataProviderOutput, error) 
 DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) 
 DeleteEventSubscription(ctx context.Context, params *DeleteEventSubscriptionInput, optFns ...func(*Options)) (*DeleteEventSubscriptionOutput, error) 
 DeleteFleetAdvisorCollector(ctx context.Context, params *DeleteFleetAdvisorCollectorInput, optFns ...func(*Options)) (*DeleteFleetAdvisorCollectorOutput, error) 
 DeleteFleetAdvisorDatabases(ctx context.Context, params *DeleteFleetAdvisorDatabasesInput, optFns ...func(*Options)) (*DeleteFleetAdvisorDatabasesOutput, error) 
 DeleteInstanceProfile(ctx context.Context, params *DeleteInstanceProfileInput, optFns ...func(*Options)) (*DeleteInstanceProfileOutput, error) 
 DeleteMigrationProject(ctx context.Context, params *DeleteMigrationProjectInput, optFns ...func(*Options)) (*DeleteMigrationProjectOutput, error) 
 DeleteReplicationConfig(ctx context.Context, params *DeleteReplicationConfigInput, optFns ...func(*Options)) (*DeleteReplicationConfigOutput, error) 
 DeleteReplicationInstance(ctx context.Context, params *DeleteReplicationInstanceInput, optFns ...func(*Options)) (*DeleteReplicationInstanceOutput, error) 
 DeleteReplicationSubnetGroup(ctx context.Context, params *DeleteReplicationSubnetGroupInput, optFns ...func(*Options)) (*DeleteReplicationSubnetGroupOutput, error) 
 DeleteReplicationTask(ctx context.Context, params *DeleteReplicationTaskInput, optFns ...func(*Options)) (*DeleteReplicationTaskOutput, error) 
 DeleteReplicationTaskAssessmentRun(ctx context.Context, params *DeleteReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*DeleteReplicationTaskAssessmentRunOutput, error) 
 DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) 
 DescribeApplicableIndividualAssessments(ctx context.Context, params *DescribeApplicableIndividualAssessmentsInput, optFns ...func(*Options)) (*DescribeApplicableIndividualAssessmentsOutput, error) 
 DescribeCertificates(ctx context.Context, params *DescribeCertificatesInput, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) 
 DescribeConnections(ctx context.Context, params *DescribeConnectionsInput, optFns ...func(*Options)) (*DescribeConnectionsOutput, error) 
 DescribeConversionConfiguration(ctx context.Context, params *DescribeConversionConfigurationInput, optFns ...func(*Options)) (*DescribeConversionConfigurationOutput, error) 
 DescribeDataMigrations(ctx context.Context, params *DescribeDataMigrationsInput, optFns ...func(*Options)) (*DescribeDataMigrationsOutput, error) 
 DescribeDataProviders(ctx context.Context, params *DescribeDataProvidersInput, optFns ...func(*Options)) (*DescribeDataProvidersOutput, error) 
 DescribeEndpointSettings(ctx context.Context, params *DescribeEndpointSettingsInput, optFns ...func(*Options)) (*DescribeEndpointSettingsOutput, error) 
 DescribeEndpointTypes(ctx context.Context, params *DescribeEndpointTypesInput, optFns ...func(*Options)) (*DescribeEndpointTypesOutput, error) 
 DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) 
 DescribeEngineVersions(ctx context.Context, params *DescribeEngineVersionsInput, optFns ...func(*Options)) (*DescribeEngineVersionsOutput, error) 
 DescribeEventCategories(ctx context.Context, params *DescribeEventCategoriesInput, optFns ...func(*Options)) (*DescribeEventCategoriesOutput, error) 
 DescribeEventSubscriptions(ctx context.Context, params *DescribeEventSubscriptionsInput, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeExtensionPackAssociations(ctx context.Context, params *DescribeExtensionPackAssociationsInput, optFns ...func(*Options)) (*DescribeExtensionPackAssociationsOutput, error) 
 DescribeFleetAdvisorCollectors(ctx context.Context, params *DescribeFleetAdvisorCollectorsInput, optFns ...func(*Options)) (*DescribeFleetAdvisorCollectorsOutput, error) 
 DescribeFleetAdvisorDatabases(ctx context.Context, params *DescribeFleetAdvisorDatabasesInput, optFns ...func(*Options)) (*DescribeFleetAdvisorDatabasesOutput, error) 
 DescribeFleetAdvisorLsaAnalysis(ctx context.Context, params *DescribeFleetAdvisorLsaAnalysisInput, optFns ...func(*Options)) (*DescribeFleetAdvisorLsaAnalysisOutput, error) 
 DescribeFleetAdvisorSchemaObjectSummary(ctx context.Context, params *DescribeFleetAdvisorSchemaObjectSummaryInput, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemaObjectSummaryOutput, error) 
 DescribeFleetAdvisorSchemas(ctx context.Context, params *DescribeFleetAdvisorSchemasInput, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemasOutput, error) 
 DescribeInstanceProfiles(ctx context.Context, params *DescribeInstanceProfilesInput, optFns ...func(*Options)) (*DescribeInstanceProfilesOutput, error) 
 DescribeMetadataModelAssessments(ctx context.Context, params *DescribeMetadataModelAssessmentsInput, optFns ...func(*Options)) (*DescribeMetadataModelAssessmentsOutput, error) 
 DescribeMetadataModelConversions(ctx context.Context, params *DescribeMetadataModelConversionsInput, optFns ...func(*Options)) (*DescribeMetadataModelConversionsOutput, error) 
 DescribeMetadataModelExportsAsScript(ctx context.Context, params *DescribeMetadataModelExportsAsScriptInput, optFns ...func(*Options)) (*DescribeMetadataModelExportsAsScriptOutput, error) 
 DescribeMetadataModelExportsToTarget(ctx context.Context, params *DescribeMetadataModelExportsToTargetInput, optFns ...func(*Options)) (*DescribeMetadataModelExportsToTargetOutput, error) 
 DescribeMetadataModelImports(ctx context.Context, params *DescribeMetadataModelImportsInput, optFns ...func(*Options)) (*DescribeMetadataModelImportsOutput, error) 
 DescribeMigrationProjects(ctx context.Context, params *DescribeMigrationProjectsInput, optFns ...func(*Options)) (*DescribeMigrationProjectsOutput, error) 
 DescribeOrderableReplicationInstances(ctx context.Context, params *DescribeOrderableReplicationInstancesInput, optFns ...func(*Options)) (*DescribeOrderableReplicationInstancesOutput, error) 
 DescribePendingMaintenanceActions(ctx context.Context, params *DescribePendingMaintenanceActionsInput, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) 
 DescribeRecommendationLimitations(ctx context.Context, params *DescribeRecommendationLimitationsInput, optFns ...func(*Options)) (*DescribeRecommendationLimitationsOutput, error) 
 DescribeRecommendations(ctx context.Context, params *DescribeRecommendationsInput, optFns ...func(*Options)) (*DescribeRecommendationsOutput, error) 
 DescribeRefreshSchemasStatus(ctx context.Context, params *DescribeRefreshSchemasStatusInput, optFns ...func(*Options)) (*DescribeRefreshSchemasStatusOutput, error) 
 DescribeReplicationConfigs(ctx context.Context, params *DescribeReplicationConfigsInput, optFns ...func(*Options)) (*DescribeReplicationConfigsOutput, error) 
 DescribeReplicationInstanceTaskLogs(ctx context.Context, params *DescribeReplicationInstanceTaskLogsInput, optFns ...func(*Options)) (*DescribeReplicationInstanceTaskLogsOutput, error) 
 DescribeReplicationInstances(ctx context.Context, params *DescribeReplicationInstancesInput, optFns ...func(*Options)) (*DescribeReplicationInstancesOutput, error) 
 DescribeReplicationSubnetGroups(ctx context.Context, params *DescribeReplicationSubnetGroupsInput, optFns ...func(*Options)) (*DescribeReplicationSubnetGroupsOutput, error) 
 DescribeReplicationTableStatistics(ctx context.Context, params *DescribeReplicationTableStatisticsInput, optFns ...func(*Options)) (*DescribeReplicationTableStatisticsOutput, error) 
 DescribeReplicationTaskAssessmentResults(ctx context.Context, params *DescribeReplicationTaskAssessmentResultsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentResultsOutput, error) 
 DescribeReplicationTaskAssessmentRuns(ctx context.Context, params *DescribeReplicationTaskAssessmentRunsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentRunsOutput, error) 
 DescribeReplicationTaskIndividualAssessments(ctx context.Context, params *DescribeReplicationTaskIndividualAssessmentsInput, optFns ...func(*Options)) (*DescribeReplicationTaskIndividualAssessmentsOutput, error) 
 DescribeReplicationTasks(ctx context.Context, params *DescribeReplicationTasksInput, optFns ...func(*Options)) (*DescribeReplicationTasksOutput, error) 
 DescribeReplications(ctx context.Context, params *DescribeReplicationsInput, optFns ...func(*Options)) (*DescribeReplicationsOutput, error) 
 DescribeSchemas(ctx context.Context, params *DescribeSchemasInput, optFns ...func(*Options)) (*DescribeSchemasOutput, error) 
 DescribeTableStatistics(ctx context.Context, params *DescribeTableStatisticsInput, optFns ...func(*Options)) (*DescribeTableStatisticsOutput, error) 
 ExportMetadataModelAssessment(ctx context.Context, params *ExportMetadataModelAssessmentInput, optFns ...func(*Options)) (*ExportMetadataModelAssessmentOutput, error) 
 ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyConversionConfiguration(ctx context.Context, params *ModifyConversionConfigurationInput, optFns ...func(*Options)) (*ModifyConversionConfigurationOutput, error) 
 ModifyDataMigration(ctx context.Context, params *ModifyDataMigrationInput, optFns ...func(*Options)) (*ModifyDataMigrationOutput, error) 
 ModifyDataProvider(ctx context.Context, params *ModifyDataProviderInput, optFns ...func(*Options)) (*ModifyDataProviderOutput, error) 
 ModifyEndpoint(ctx context.Context, params *ModifyEndpointInput, optFns ...func(*Options)) (*ModifyEndpointOutput, error) 
 ModifyEventSubscription(ctx context.Context, params *ModifyEventSubscriptionInput, optFns ...func(*Options)) (*ModifyEventSubscriptionOutput, error) 
 ModifyInstanceProfile(ctx context.Context, params *ModifyInstanceProfileInput, optFns ...func(*Options)) (*ModifyInstanceProfileOutput, error) 
 ModifyMigrationProject(ctx context.Context, params *ModifyMigrationProjectInput, optFns ...func(*Options)) (*ModifyMigrationProjectOutput, error) 
 ModifyReplicationConfig(ctx context.Context, params *ModifyReplicationConfigInput, optFns ...func(*Options)) (*ModifyReplicationConfigOutput, error) 
 ModifyReplicationInstance(ctx context.Context, params *ModifyReplicationInstanceInput, optFns ...func(*Options)) (*ModifyReplicationInstanceOutput, error) 
 ModifyReplicationSubnetGroup(ctx context.Context, params *ModifyReplicationSubnetGroupInput, optFns ...func(*Options)) (*ModifyReplicationSubnetGroupOutput, error) 
 ModifyReplicationTask(ctx context.Context, params *ModifyReplicationTaskInput, optFns ...func(*Options)) (*ModifyReplicationTaskOutput, error) 
 MoveReplicationTask(ctx context.Context, params *MoveReplicationTaskInput, optFns ...func(*Options)) (*MoveReplicationTaskOutput, error) 
 RebootReplicationInstance(ctx context.Context, params *RebootReplicationInstanceInput, optFns ...func(*Options)) (*RebootReplicationInstanceOutput, error) 
 RefreshSchemas(ctx context.Context, params *RefreshSchemasInput, optFns ...func(*Options)) (*RefreshSchemasOutput, error) 
 ReloadReplicationTables(ctx context.Context, params *ReloadReplicationTablesInput, optFns ...func(*Options)) (*ReloadReplicationTablesOutput, error) 
 ReloadTables(ctx context.Context, params *ReloadTablesInput, optFns ...func(*Options)) (*ReloadTablesOutput, error) 
 RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) 
 RunFleetAdvisorLsaAnalysis(ctx context.Context, params *RunFleetAdvisorLsaAnalysisInput, optFns ...func(*Options)) (*RunFleetAdvisorLsaAnalysisOutput, error) 
 StartDataMigration(ctx context.Context, params *StartDataMigrationInput, optFns ...func(*Options)) (*StartDataMigrationOutput, error) 
 StartExtensionPackAssociation(ctx context.Context, params *StartExtensionPackAssociationInput, optFns ...func(*Options)) (*StartExtensionPackAssociationOutput, error) 
 StartMetadataModelAssessment(ctx context.Context, params *StartMetadataModelAssessmentInput, optFns ...func(*Options)) (*StartMetadataModelAssessmentOutput, error) 
 StartMetadataModelConversion(ctx context.Context, params *StartMetadataModelConversionInput, optFns ...func(*Options)) (*StartMetadataModelConversionOutput, error) 
 StartMetadataModelExportAsScript(ctx context.Context, params *StartMetadataModelExportAsScriptInput, optFns ...func(*Options)) (*StartMetadataModelExportAsScriptOutput, error) 
 StartMetadataModelExportToTarget(ctx context.Context, params *StartMetadataModelExportToTargetInput, optFns ...func(*Options)) (*StartMetadataModelExportToTargetOutput, error) 
 StartMetadataModelImport(ctx context.Context, params *StartMetadataModelImportInput, optFns ...func(*Options)) (*StartMetadataModelImportOutput, error) 
 StartRecommendations(ctx context.Context, params *StartRecommendationsInput, optFns ...func(*Options)) (*StartRecommendationsOutput, error) 
 StartReplication(ctx context.Context, params *StartReplicationInput, optFns ...func(*Options)) (*StartReplicationOutput, error) 
 StartReplicationTask(ctx context.Context, params *StartReplicationTaskInput, optFns ...func(*Options)) (*StartReplicationTaskOutput, error) 
 StartReplicationTaskAssessment(ctx context.Context, params *StartReplicationTaskAssessmentInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentOutput, error) 
 StartReplicationTaskAssessmentRun(ctx context.Context, params *StartReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentRunOutput, error) 
 StopDataMigration(ctx context.Context, params *StopDataMigrationInput, optFns ...func(*Options)) (*StopDataMigrationOutput, error) 
 StopReplication(ctx context.Context, params *StopReplicationInput, optFns ...func(*Options)) (*StopReplicationOutput, error) 
 StopReplicationTask(ctx context.Context, params *StopReplicationTaskInput, optFns ...func(*Options)) (*StopReplicationTaskOutput, error) 
 TestConnection(ctx context.Context, params *TestConnectionInput, optFns ...func(*Options)) (*TestConnectionOutput, error) 
 UpdateSubscriptionsToEventBridge(ctx context.Context, params *UpdateSubscriptionsToEventBridgeInput, optFns ...func(*Options)) (*UpdateSubscriptionsToEventBridgeOutput, error) 
}
