package databasemigrationservice_test

// tests for the databasemigrationservice service interface for this ../../../service/databasemigrationservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/databasemigrationservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/databasemigrationservice/databasemigrationservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestDatabasemigrationserviceServiceCanBeMocked(t *testing.T) {
	var iface databasemigrationservice_iface.IClient
	iface = &databasemigrationservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := databasemigrationservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &databasemigrationservice.AddTagsToResourceInput{}
        output := &databasemigrationservice.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyPendingMaintenanceAction", func(t *testing.T) {
        input := &databasemigrationservice.ApplyPendingMaintenanceActionInput{}
        output := &databasemigrationservice.ApplyPendingMaintenanceActionOutput{}

        mockClient.On("ApplyPendingMaintenanceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyPendingMaintenanceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStartRecommendations", func(t *testing.T) {
        input := &databasemigrationservice.BatchStartRecommendationsInput{}
        output := &databasemigrationservice.BatchStartRecommendationsOutput{}

        mockClient.On("BatchStartRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStartRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelReplicationTaskAssessmentRun", func(t *testing.T) {
        input := &databasemigrationservice.CancelReplicationTaskAssessmentRunInput{}
        output := &databasemigrationservice.CancelReplicationTaskAssessmentRunOutput{}

        mockClient.On("CancelReplicationTaskAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelReplicationTaskAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataProvider", func(t *testing.T) {
        input := &databasemigrationservice.CreateDataProviderInput{}
        output := &databasemigrationservice.CreateDataProviderOutput{}

        mockClient.On("CreateDataProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpoint", func(t *testing.T) {
        input := &databasemigrationservice.CreateEndpointInput{}
        output := &databasemigrationservice.CreateEndpointOutput{}

        mockClient.On("CreateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSubscription", func(t *testing.T) {
        input := &databasemigrationservice.CreateEventSubscriptionInput{}
        output := &databasemigrationservice.CreateEventSubscriptionOutput{}

        mockClient.On("CreateEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleetAdvisorCollector", func(t *testing.T) {
        input := &databasemigrationservice.CreateFleetAdvisorCollectorInput{}
        output := &databasemigrationservice.CreateFleetAdvisorCollectorOutput{}

        mockClient.On("CreateFleetAdvisorCollector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleetAdvisorCollector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceProfile", func(t *testing.T) {
        input := &databasemigrationservice.CreateInstanceProfileInput{}
        output := &databasemigrationservice.CreateInstanceProfileOutput{}

        mockClient.On("CreateInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMigrationProject", func(t *testing.T) {
        input := &databasemigrationservice.CreateMigrationProjectInput{}
        output := &databasemigrationservice.CreateMigrationProjectOutput{}

        mockClient.On("CreateMigrationProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMigrationProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationConfig", func(t *testing.T) {
        input := &databasemigrationservice.CreateReplicationConfigInput{}
        output := &databasemigrationservice.CreateReplicationConfigOutput{}

        mockClient.On("CreateReplicationConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationInstance", func(t *testing.T) {
        input := &databasemigrationservice.CreateReplicationInstanceInput{}
        output := &databasemigrationservice.CreateReplicationInstanceOutput{}

        mockClient.On("CreateReplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationSubnetGroup", func(t *testing.T) {
        input := &databasemigrationservice.CreateReplicationSubnetGroupInput{}
        output := &databasemigrationservice.CreateReplicationSubnetGroupOutput{}

        mockClient.On("CreateReplicationSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.CreateReplicationTaskInput{}
        output := &databasemigrationservice.CreateReplicationTaskOutput{}

        mockClient.On("CreateReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificate", func(t *testing.T) {
        input := &databasemigrationservice.DeleteCertificateInput{}
        output := &databasemigrationservice.DeleteCertificateOutput{}

        mockClient.On("DeleteCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &databasemigrationservice.DeleteConnectionInput{}
        output := &databasemigrationservice.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataProvider", func(t *testing.T) {
        input := &databasemigrationservice.DeleteDataProviderInput{}
        output := &databasemigrationservice.DeleteDataProviderOutput{}

        mockClient.On("DeleteDataProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &databasemigrationservice.DeleteEndpointInput{}
        output := &databasemigrationservice.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSubscription", func(t *testing.T) {
        input := &databasemigrationservice.DeleteEventSubscriptionInput{}
        output := &databasemigrationservice.DeleteEventSubscriptionOutput{}

        mockClient.On("DeleteEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleetAdvisorCollector", func(t *testing.T) {
        input := &databasemigrationservice.DeleteFleetAdvisorCollectorInput{}
        output := &databasemigrationservice.DeleteFleetAdvisorCollectorOutput{}

        mockClient.On("DeleteFleetAdvisorCollector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleetAdvisorCollector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleetAdvisorDatabases", func(t *testing.T) {
        input := &databasemigrationservice.DeleteFleetAdvisorDatabasesInput{}
        output := &databasemigrationservice.DeleteFleetAdvisorDatabasesOutput{}

        mockClient.On("DeleteFleetAdvisorDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleetAdvisorDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceProfile", func(t *testing.T) {
        input := &databasemigrationservice.DeleteInstanceProfileInput{}
        output := &databasemigrationservice.DeleteInstanceProfileOutput{}

        mockClient.On("DeleteInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMigrationProject", func(t *testing.T) {
        input := &databasemigrationservice.DeleteMigrationProjectInput{}
        output := &databasemigrationservice.DeleteMigrationProjectOutput{}

        mockClient.On("DeleteMigrationProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMigrationProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationConfig", func(t *testing.T) {
        input := &databasemigrationservice.DeleteReplicationConfigInput{}
        output := &databasemigrationservice.DeleteReplicationConfigOutput{}

        mockClient.On("DeleteReplicationConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationInstance", func(t *testing.T) {
        input := &databasemigrationservice.DeleteReplicationInstanceInput{}
        output := &databasemigrationservice.DeleteReplicationInstanceOutput{}

        mockClient.On("DeleteReplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationSubnetGroup", func(t *testing.T) {
        input := &databasemigrationservice.DeleteReplicationSubnetGroupInput{}
        output := &databasemigrationservice.DeleteReplicationSubnetGroupOutput{}

        mockClient.On("DeleteReplicationSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.DeleteReplicationTaskInput{}
        output := &databasemigrationservice.DeleteReplicationTaskOutput{}

        mockClient.On("DeleteReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationTaskAssessmentRun", func(t *testing.T) {
        input := &databasemigrationservice.DeleteReplicationTaskAssessmentRunInput{}
        output := &databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput{}

        mockClient.On("DeleteReplicationTaskAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationTaskAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &databasemigrationservice.DescribeAccountAttributesInput{}
        output := &databasemigrationservice.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicableIndividualAssessments", func(t *testing.T) {
        input := &databasemigrationservice.DescribeApplicableIndividualAssessmentsInput{}
        output := &databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput{}

        mockClient.On("DescribeApplicableIndividualAssessments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicableIndividualAssessments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificates", func(t *testing.T) {
        input := &databasemigrationservice.DescribeCertificatesInput{}
        output := &databasemigrationservice.DescribeCertificatesOutput{}

        mockClient.On("DescribeCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnections", func(t *testing.T) {
        input := &databasemigrationservice.DescribeConnectionsInput{}
        output := &databasemigrationservice.DescribeConnectionsOutput{}

        mockClient.On("DescribeConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConversionConfiguration", func(t *testing.T) {
        input := &databasemigrationservice.DescribeConversionConfigurationInput{}
        output := &databasemigrationservice.DescribeConversionConfigurationOutput{}

        mockClient.On("DescribeConversionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConversionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataProviders", func(t *testing.T) {
        input := &databasemigrationservice.DescribeDataProvidersInput{}
        output := &databasemigrationservice.DescribeDataProvidersOutput{}

        mockClient.On("DescribeDataProviders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointSettings", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEndpointSettingsInput{}
        output := &databasemigrationservice.DescribeEndpointSettingsOutput{}

        mockClient.On("DescribeEndpointSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointTypes", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEndpointTypesInput{}
        output := &databasemigrationservice.DescribeEndpointTypesOutput{}

        mockClient.On("DescribeEndpointTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoints", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEndpointsInput{}
        output := &databasemigrationservice.DescribeEndpointsOutput{}

        mockClient.On("DescribeEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngineVersions", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEngineVersionsInput{}
        output := &databasemigrationservice.DescribeEngineVersionsOutput{}

        mockClient.On("DescribeEngineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventCategories", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEventCategoriesInput{}
        output := &databasemigrationservice.DescribeEventCategoriesOutput{}

        mockClient.On("DescribeEventCategories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSubscriptions", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEventSubscriptionsInput{}
        output := &databasemigrationservice.DescribeEventSubscriptionsOutput{}

        mockClient.On("DescribeEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &databasemigrationservice.DescribeEventsInput{}
        output := &databasemigrationservice.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExtensionPackAssociations", func(t *testing.T) {
        input := &databasemigrationservice.DescribeExtensionPackAssociationsInput{}
        output := &databasemigrationservice.DescribeExtensionPackAssociationsOutput{}

        mockClient.On("DescribeExtensionPackAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExtensionPackAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAdvisorCollectors", func(t *testing.T) {
        input := &databasemigrationservice.DescribeFleetAdvisorCollectorsInput{}
        output := &databasemigrationservice.DescribeFleetAdvisorCollectorsOutput{}

        mockClient.On("DescribeFleetAdvisorCollectors", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAdvisorCollectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAdvisorDatabases", func(t *testing.T) {
        input := &databasemigrationservice.DescribeFleetAdvisorDatabasesInput{}
        output := &databasemigrationservice.DescribeFleetAdvisorDatabasesOutput{}

        mockClient.On("DescribeFleetAdvisorDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAdvisorDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAdvisorLsaAnalysis", func(t *testing.T) {
        input := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput{}
        output := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput{}

        mockClient.On("DescribeFleetAdvisorLsaAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAdvisorLsaAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAdvisorSchemaObjectSummary", func(t *testing.T) {
        input := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput{}
        output := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput{}

        mockClient.On("DescribeFleetAdvisorSchemaObjectSummary", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAdvisorSchemaObjectSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAdvisorSchemas", func(t *testing.T) {
        input := &databasemigrationservice.DescribeFleetAdvisorSchemasInput{}
        output := &databasemigrationservice.DescribeFleetAdvisorSchemasOutput{}

        mockClient.On("DescribeFleetAdvisorSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAdvisorSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceProfiles", func(t *testing.T) {
        input := &databasemigrationservice.DescribeInstanceProfilesInput{}
        output := &databasemigrationservice.DescribeInstanceProfilesOutput{}

        mockClient.On("DescribeInstanceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetadataModelAssessments", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMetadataModelAssessmentsInput{}
        output := &databasemigrationservice.DescribeMetadataModelAssessmentsOutput{}

        mockClient.On("DescribeMetadataModelAssessments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetadataModelAssessments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetadataModelConversions", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMetadataModelConversionsInput{}
        output := &databasemigrationservice.DescribeMetadataModelConversionsOutput{}

        mockClient.On("DescribeMetadataModelConversions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetadataModelConversions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetadataModelExportsAsScript", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMetadataModelExportsAsScriptInput{}
        output := &databasemigrationservice.DescribeMetadataModelExportsAsScriptOutput{}

        mockClient.On("DescribeMetadataModelExportsAsScript", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetadataModelExportsAsScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetadataModelExportsToTarget", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMetadataModelExportsToTargetInput{}
        output := &databasemigrationservice.DescribeMetadataModelExportsToTargetOutput{}

        mockClient.On("DescribeMetadataModelExportsToTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetadataModelExportsToTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetadataModelImports", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMetadataModelImportsInput{}
        output := &databasemigrationservice.DescribeMetadataModelImportsOutput{}

        mockClient.On("DescribeMetadataModelImports", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetadataModelImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMigrationProjects", func(t *testing.T) {
        input := &databasemigrationservice.DescribeMigrationProjectsInput{}
        output := &databasemigrationservice.DescribeMigrationProjectsOutput{}

        mockClient.On("DescribeMigrationProjects", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMigrationProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrderableReplicationInstances", func(t *testing.T) {
        input := &databasemigrationservice.DescribeOrderableReplicationInstancesInput{}
        output := &databasemigrationservice.DescribeOrderableReplicationInstancesOutput{}

        mockClient.On("DescribeOrderableReplicationInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrderableReplicationInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePendingMaintenanceActions", func(t *testing.T) {
        input := &databasemigrationservice.DescribePendingMaintenanceActionsInput{}
        output := &databasemigrationservice.DescribePendingMaintenanceActionsOutput{}

        mockClient.On("DescribePendingMaintenanceActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePendingMaintenanceActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecommendationLimitations", func(t *testing.T) {
        input := &databasemigrationservice.DescribeRecommendationLimitationsInput{}
        output := &databasemigrationservice.DescribeRecommendationLimitationsOutput{}

        mockClient.On("DescribeRecommendationLimitations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecommendationLimitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecommendations", func(t *testing.T) {
        input := &databasemigrationservice.DescribeRecommendationsInput{}
        output := &databasemigrationservice.DescribeRecommendationsOutput{}

        mockClient.On("DescribeRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRefreshSchemasStatus", func(t *testing.T) {
        input := &databasemigrationservice.DescribeRefreshSchemasStatusInput{}
        output := &databasemigrationservice.DescribeRefreshSchemasStatusOutput{}

        mockClient.On("DescribeRefreshSchemasStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRefreshSchemasStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationConfigs", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationConfigsInput{}
        output := &databasemigrationservice.DescribeReplicationConfigsOutput{}

        mockClient.On("DescribeReplicationConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationInstanceTaskLogs", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationInstanceTaskLogsInput{}
        output := &databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput{}

        mockClient.On("DescribeReplicationInstanceTaskLogs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationInstanceTaskLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationInstances", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationInstancesInput{}
        output := &databasemigrationservice.DescribeReplicationInstancesOutput{}

        mockClient.On("DescribeReplicationInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationSubnetGroups", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationSubnetGroupsInput{}
        output := &databasemigrationservice.DescribeReplicationSubnetGroupsOutput{}

        mockClient.On("DescribeReplicationSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationTableStatistics", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationTableStatisticsInput{}
        output := &databasemigrationservice.DescribeReplicationTableStatisticsOutput{}

        mockClient.On("DescribeReplicationTableStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationTableStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationTaskAssessmentResults", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput{}
        output := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput{}

        mockClient.On("DescribeReplicationTaskAssessmentResults", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationTaskAssessmentResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationTaskAssessmentRuns", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput{}
        output := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput{}

        mockClient.On("DescribeReplicationTaskAssessmentRuns", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationTaskAssessmentRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationTaskIndividualAssessments", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput{}
        output := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput{}

        mockClient.On("DescribeReplicationTaskIndividualAssessments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationTaskIndividualAssessments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationTasks", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationTasksInput{}
        output := &databasemigrationservice.DescribeReplicationTasksOutput{}

        mockClient.On("DescribeReplicationTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplications", func(t *testing.T) {
        input := &databasemigrationservice.DescribeReplicationsInput{}
        output := &databasemigrationservice.DescribeReplicationsOutput{}

        mockClient.On("DescribeReplications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchemas", func(t *testing.T) {
        input := &databasemigrationservice.DescribeSchemasInput{}
        output := &databasemigrationservice.DescribeSchemasOutput{}

        mockClient.On("DescribeSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTableStatistics", func(t *testing.T) {
        input := &databasemigrationservice.DescribeTableStatisticsInput{}
        output := &databasemigrationservice.DescribeTableStatisticsOutput{}

        mockClient.On("DescribeTableStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTableStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportMetadataModelAssessment", func(t *testing.T) {
        input := &databasemigrationservice.ExportMetadataModelAssessmentInput{}
        output := &databasemigrationservice.ExportMetadataModelAssessmentOutput{}

        mockClient.On("ExportMetadataModelAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.ExportMetadataModelAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCertificate", func(t *testing.T) {
        input := &databasemigrationservice.ImportCertificateInput{}
        output := &databasemigrationservice.ImportCertificateOutput{}

        mockClient.On("ImportCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &databasemigrationservice.ListTagsForResourceInput{}
        output := &databasemigrationservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyConversionConfiguration", func(t *testing.T) {
        input := &databasemigrationservice.ModifyConversionConfigurationInput{}
        output := &databasemigrationservice.ModifyConversionConfigurationOutput{}

        mockClient.On("ModifyConversionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyConversionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDataProvider", func(t *testing.T) {
        input := &databasemigrationservice.ModifyDataProviderInput{}
        output := &databasemigrationservice.ModifyDataProviderOutput{}

        mockClient.On("ModifyDataProvider", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDataProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEndpoint", func(t *testing.T) {
        input := &databasemigrationservice.ModifyEndpointInput{}
        output := &databasemigrationservice.ModifyEndpointOutput{}

        mockClient.On("ModifyEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEventSubscription", func(t *testing.T) {
        input := &databasemigrationservice.ModifyEventSubscriptionInput{}
        output := &databasemigrationservice.ModifyEventSubscriptionOutput{}

        mockClient.On("ModifyEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceProfile", func(t *testing.T) {
        input := &databasemigrationservice.ModifyInstanceProfileInput{}
        output := &databasemigrationservice.ModifyInstanceProfileOutput{}

        mockClient.On("ModifyInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyMigrationProject", func(t *testing.T) {
        input := &databasemigrationservice.ModifyMigrationProjectInput{}
        output := &databasemigrationservice.ModifyMigrationProjectOutput{}

        mockClient.On("ModifyMigrationProject", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyMigrationProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationConfig", func(t *testing.T) {
        input := &databasemigrationservice.ModifyReplicationConfigInput{}
        output := &databasemigrationservice.ModifyReplicationConfigOutput{}

        mockClient.On("ModifyReplicationConfig", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationInstance", func(t *testing.T) {
        input := &databasemigrationservice.ModifyReplicationInstanceInput{}
        output := &databasemigrationservice.ModifyReplicationInstanceOutput{}

        mockClient.On("ModifyReplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationSubnetGroup", func(t *testing.T) {
        input := &databasemigrationservice.ModifyReplicationSubnetGroupInput{}
        output := &databasemigrationservice.ModifyReplicationSubnetGroupOutput{}

        mockClient.On("ModifyReplicationSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.ModifyReplicationTaskInput{}
        output := &databasemigrationservice.ModifyReplicationTaskOutput{}

        mockClient.On("ModifyReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMoveReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.MoveReplicationTaskInput{}
        output := &databasemigrationservice.MoveReplicationTaskOutput{}

        mockClient.On("MoveReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.MoveReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootReplicationInstance", func(t *testing.T) {
        input := &databasemigrationservice.RebootReplicationInstanceInput{}
        output := &databasemigrationservice.RebootReplicationInstanceOutput{}

        mockClient.On("RebootReplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootReplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRefreshSchemas", func(t *testing.T) {
        input := &databasemigrationservice.RefreshSchemasInput{}
        output := &databasemigrationservice.RefreshSchemasOutput{}

        mockClient.On("RefreshSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.RefreshSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReloadReplicationTables", func(t *testing.T) {
        input := &databasemigrationservice.ReloadReplicationTablesInput{}
        output := &databasemigrationservice.ReloadReplicationTablesOutput{}

        mockClient.On("ReloadReplicationTables", ctx, input).Return(output, nil)

        result, err := mockClient.ReloadReplicationTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReloadTables", func(t *testing.T) {
        input := &databasemigrationservice.ReloadTablesInput{}
        output := &databasemigrationservice.ReloadTablesOutput{}

        mockClient.On("ReloadTables", ctx, input).Return(output, nil)

        result, err := mockClient.ReloadTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &databasemigrationservice.RemoveTagsFromResourceInput{}
        output := &databasemigrationservice.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunFleetAdvisorLsaAnalysis", func(t *testing.T) {
        input := &databasemigrationservice.RunFleetAdvisorLsaAnalysisInput{}
        output := &databasemigrationservice.RunFleetAdvisorLsaAnalysisOutput{}

        mockClient.On("RunFleetAdvisorLsaAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.RunFleetAdvisorLsaAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExtensionPackAssociation", func(t *testing.T) {
        input := &databasemigrationservice.StartExtensionPackAssociationInput{}
        output := &databasemigrationservice.StartExtensionPackAssociationOutput{}

        mockClient.On("StartExtensionPackAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.StartExtensionPackAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataModelAssessment", func(t *testing.T) {
        input := &databasemigrationservice.StartMetadataModelAssessmentInput{}
        output := &databasemigrationservice.StartMetadataModelAssessmentOutput{}

        mockClient.On("StartMetadataModelAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataModelAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataModelConversion", func(t *testing.T) {
        input := &databasemigrationservice.StartMetadataModelConversionInput{}
        output := &databasemigrationservice.StartMetadataModelConversionOutput{}

        mockClient.On("StartMetadataModelConversion", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataModelConversion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataModelExportAsScript", func(t *testing.T) {
        input := &databasemigrationservice.StartMetadataModelExportAsScriptInput{}
        output := &databasemigrationservice.StartMetadataModelExportAsScriptOutput{}

        mockClient.On("StartMetadataModelExportAsScript", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataModelExportAsScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataModelExportToTarget", func(t *testing.T) {
        input := &databasemigrationservice.StartMetadataModelExportToTargetInput{}
        output := &databasemigrationservice.StartMetadataModelExportToTargetOutput{}

        mockClient.On("StartMetadataModelExportToTarget", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataModelExportToTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataModelImport", func(t *testing.T) {
        input := &databasemigrationservice.StartMetadataModelImportInput{}
        output := &databasemigrationservice.StartMetadataModelImportOutput{}

        mockClient.On("StartMetadataModelImport", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataModelImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRecommendations", func(t *testing.T) {
        input := &databasemigrationservice.StartRecommendationsInput{}
        output := &databasemigrationservice.StartRecommendationsOutput{}

        mockClient.On("StartRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.StartRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplication", func(t *testing.T) {
        input := &databasemigrationservice.StartReplicationInput{}
        output := &databasemigrationservice.StartReplicationOutput{}

        mockClient.On("StartReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.StartReplicationTaskInput{}
        output := &databasemigrationservice.StartReplicationTaskOutput{}

        mockClient.On("StartReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplicationTaskAssessment", func(t *testing.T) {
        input := &databasemigrationservice.StartReplicationTaskAssessmentInput{}
        output := &databasemigrationservice.StartReplicationTaskAssessmentOutput{}

        mockClient.On("StartReplicationTaskAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplicationTaskAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplicationTaskAssessmentRun", func(t *testing.T) {
        input := &databasemigrationservice.StartReplicationTaskAssessmentRunInput{}
        output := &databasemigrationservice.StartReplicationTaskAssessmentRunOutput{}

        mockClient.On("StartReplicationTaskAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplicationTaskAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopReplication", func(t *testing.T) {
        input := &databasemigrationservice.StopReplicationInput{}
        output := &databasemigrationservice.StopReplicationOutput{}

        mockClient.On("StopReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopReplicationTask", func(t *testing.T) {
        input := &databasemigrationservice.StopReplicationTaskInput{}
        output := &databasemigrationservice.StopReplicationTaskOutput{}

        mockClient.On("StopReplicationTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopReplicationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestConnection", func(t *testing.T) {
        input := &databasemigrationservice.TestConnectionInput{}
        output := &databasemigrationservice.TestConnectionOutput{}

        mockClient.On("TestConnection", ctx, input).Return(output, nil)

        result, err := mockClient.TestConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriptionsToEventBridge", func(t *testing.T) {
        input := &databasemigrationservice.UpdateSubscriptionsToEventBridgeInput{}
        output := &databasemigrationservice.UpdateSubscriptionsToEventBridgeOutput{}

        mockClient.On("UpdateSubscriptionsToEventBridge", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriptionsToEventBridge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
