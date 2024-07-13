package configservice_test

// tests for the configservice service interface for this ../../../service/configservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/configservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/configservice/configservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestConfigserviceServiceCanBeMocked(t *testing.T) {
	var iface configservice_iface.IClient
	iface = &configservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := configservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAggregateResourceConfig", func(t *testing.T) {
        input := &configservice.BatchGetAggregateResourceConfigInput{}
        output := &configservice.BatchGetAggregateResourceConfigOutput{}

        mockClient.On("BatchGetAggregateResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAggregateResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetResourceConfig", func(t *testing.T) {
        input := &configservice.BatchGetResourceConfigInput{}
        output := &configservice.BatchGetResourceConfigOutput{}

        mockClient.On("BatchGetResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAggregationAuthorization", func(t *testing.T) {
        input := &configservice.DeleteAggregationAuthorizationInput{}
        output := &configservice.DeleteAggregationAuthorizationOutput{}

        mockClient.On("DeleteAggregationAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAggregationAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigRule", func(t *testing.T) {
        input := &configservice.DeleteConfigRuleInput{}
        output := &configservice.DeleteConfigRuleOutput{}

        mockClient.On("DeleteConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationAggregator", func(t *testing.T) {
        input := &configservice.DeleteConfigurationAggregatorInput{}
        output := &configservice.DeleteConfigurationAggregatorOutput{}

        mockClient.On("DeleteConfigurationAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationRecorder", func(t *testing.T) {
        input := &configservice.DeleteConfigurationRecorderInput{}
        output := &configservice.DeleteConfigurationRecorderOutput{}

        mockClient.On("DeleteConfigurationRecorder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationRecorder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConformancePack", func(t *testing.T) {
        input := &configservice.DeleteConformancePackInput{}
        output := &configservice.DeleteConformancePackOutput{}

        mockClient.On("DeleteConformancePack", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConformancePack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeliveryChannel", func(t *testing.T) {
        input := &configservice.DeleteDeliveryChannelInput{}
        output := &configservice.DeleteDeliveryChannelOutput{}

        mockClient.On("DeleteDeliveryChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeliveryChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEvaluationResults", func(t *testing.T) {
        input := &configservice.DeleteEvaluationResultsInput{}
        output := &configservice.DeleteEvaluationResultsOutput{}

        mockClient.On("DeleteEvaluationResults", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEvaluationResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOrganizationConfigRule", func(t *testing.T) {
        input := &configservice.DeleteOrganizationConfigRuleInput{}
        output := &configservice.DeleteOrganizationConfigRuleOutput{}

        mockClient.On("DeleteOrganizationConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOrganizationConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOrganizationConformancePack", func(t *testing.T) {
        input := &configservice.DeleteOrganizationConformancePackInput{}
        output := &configservice.DeleteOrganizationConformancePackOutput{}

        mockClient.On("DeleteOrganizationConformancePack", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOrganizationConformancePack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePendingAggregationRequest", func(t *testing.T) {
        input := &configservice.DeletePendingAggregationRequestInput{}
        output := &configservice.DeletePendingAggregationRequestOutput{}

        mockClient.On("DeletePendingAggregationRequest", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePendingAggregationRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRemediationConfiguration", func(t *testing.T) {
        input := &configservice.DeleteRemediationConfigurationInput{}
        output := &configservice.DeleteRemediationConfigurationOutput{}

        mockClient.On("DeleteRemediationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRemediationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRemediationExceptions", func(t *testing.T) {
        input := &configservice.DeleteRemediationExceptionsInput{}
        output := &configservice.DeleteRemediationExceptionsOutput{}

        mockClient.On("DeleteRemediationExceptions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRemediationExceptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceConfig", func(t *testing.T) {
        input := &configservice.DeleteResourceConfigInput{}
        output := &configservice.DeleteResourceConfigOutput{}

        mockClient.On("DeleteResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRetentionConfiguration", func(t *testing.T) {
        input := &configservice.DeleteRetentionConfigurationInput{}
        output := &configservice.DeleteRetentionConfigurationOutput{}

        mockClient.On("DeleteRetentionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRetentionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStoredQuery", func(t *testing.T) {
        input := &configservice.DeleteStoredQueryInput{}
        output := &configservice.DeleteStoredQueryOutput{}

        mockClient.On("DeleteStoredQuery", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStoredQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeliverConfigSnapshot", func(t *testing.T) {
        input := &configservice.DeliverConfigSnapshotInput{}
        output := &configservice.DeliverConfigSnapshotOutput{}

        mockClient.On("DeliverConfigSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeliverConfigSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAggregateComplianceByConfigRules", func(t *testing.T) {
        input := &configservice.DescribeAggregateComplianceByConfigRulesInput{}
        output := &configservice.DescribeAggregateComplianceByConfigRulesOutput{}

        mockClient.On("DescribeAggregateComplianceByConfigRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAggregateComplianceByConfigRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAggregateComplianceByConformancePacks", func(t *testing.T) {
        input := &configservice.DescribeAggregateComplianceByConformancePacksInput{}
        output := &configservice.DescribeAggregateComplianceByConformancePacksOutput{}

        mockClient.On("DescribeAggregateComplianceByConformancePacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAggregateComplianceByConformancePacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAggregationAuthorizations", func(t *testing.T) {
        input := &configservice.DescribeAggregationAuthorizationsInput{}
        output := &configservice.DescribeAggregationAuthorizationsOutput{}

        mockClient.On("DescribeAggregationAuthorizations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAggregationAuthorizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComplianceByConfigRule", func(t *testing.T) {
        input := &configservice.DescribeComplianceByConfigRuleInput{}
        output := &configservice.DescribeComplianceByConfigRuleOutput{}

        mockClient.On("DescribeComplianceByConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComplianceByConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComplianceByResource", func(t *testing.T) {
        input := &configservice.DescribeComplianceByResourceInput{}
        output := &configservice.DescribeComplianceByResourceOutput{}

        mockClient.On("DescribeComplianceByResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComplianceByResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigRuleEvaluationStatus", func(t *testing.T) {
        input := &configservice.DescribeConfigRuleEvaluationStatusInput{}
        output := &configservice.DescribeConfigRuleEvaluationStatusOutput{}

        mockClient.On("DescribeConfigRuleEvaluationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigRuleEvaluationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigRules", func(t *testing.T) {
        input := &configservice.DescribeConfigRulesInput{}
        output := &configservice.DescribeConfigRulesOutput{}

        mockClient.On("DescribeConfigRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationAggregatorSourcesStatus", func(t *testing.T) {
        input := &configservice.DescribeConfigurationAggregatorSourcesStatusInput{}
        output := &configservice.DescribeConfigurationAggregatorSourcesStatusOutput{}

        mockClient.On("DescribeConfigurationAggregatorSourcesStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationAggregatorSourcesStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationAggregators", func(t *testing.T) {
        input := &configservice.DescribeConfigurationAggregatorsInput{}
        output := &configservice.DescribeConfigurationAggregatorsOutput{}

        mockClient.On("DescribeConfigurationAggregators", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationAggregators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationRecorderStatus", func(t *testing.T) {
        input := &configservice.DescribeConfigurationRecorderStatusInput{}
        output := &configservice.DescribeConfigurationRecorderStatusOutput{}

        mockClient.On("DescribeConfigurationRecorderStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationRecorderStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationRecorders", func(t *testing.T) {
        input := &configservice.DescribeConfigurationRecordersInput{}
        output := &configservice.DescribeConfigurationRecordersOutput{}

        mockClient.On("DescribeConfigurationRecorders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationRecorders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConformancePackCompliance", func(t *testing.T) {
        input := &configservice.DescribeConformancePackComplianceInput{}
        output := &configservice.DescribeConformancePackComplianceOutput{}

        mockClient.On("DescribeConformancePackCompliance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConformancePackCompliance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConformancePackStatus", func(t *testing.T) {
        input := &configservice.DescribeConformancePackStatusInput{}
        output := &configservice.DescribeConformancePackStatusOutput{}

        mockClient.On("DescribeConformancePackStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConformancePackStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConformancePacks", func(t *testing.T) {
        input := &configservice.DescribeConformancePacksInput{}
        output := &configservice.DescribeConformancePacksOutput{}

        mockClient.On("DescribeConformancePacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConformancePacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliveryChannelStatus", func(t *testing.T) {
        input := &configservice.DescribeDeliveryChannelStatusInput{}
        output := &configservice.DescribeDeliveryChannelStatusOutput{}

        mockClient.On("DescribeDeliveryChannelStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliveryChannelStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliveryChannels", func(t *testing.T) {
        input := &configservice.DescribeDeliveryChannelsInput{}
        output := &configservice.DescribeDeliveryChannelsOutput{}

        mockClient.On("DescribeDeliveryChannels", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliveryChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfigRuleStatuses", func(t *testing.T) {
        input := &configservice.DescribeOrganizationConfigRuleStatusesInput{}
        output := &configservice.DescribeOrganizationConfigRuleStatusesOutput{}

        mockClient.On("DescribeOrganizationConfigRuleStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfigRuleStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfigRules", func(t *testing.T) {
        input := &configservice.DescribeOrganizationConfigRulesInput{}
        output := &configservice.DescribeOrganizationConfigRulesOutput{}

        mockClient.On("DescribeOrganizationConfigRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfigRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConformancePackStatuses", func(t *testing.T) {
        input := &configservice.DescribeOrganizationConformancePackStatusesInput{}
        output := &configservice.DescribeOrganizationConformancePackStatusesOutput{}

        mockClient.On("DescribeOrganizationConformancePackStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConformancePackStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConformancePacks", func(t *testing.T) {
        input := &configservice.DescribeOrganizationConformancePacksInput{}
        output := &configservice.DescribeOrganizationConformancePacksOutput{}

        mockClient.On("DescribeOrganizationConformancePacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConformancePacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePendingAggregationRequests", func(t *testing.T) {
        input := &configservice.DescribePendingAggregationRequestsInput{}
        output := &configservice.DescribePendingAggregationRequestsOutput{}

        mockClient.On("DescribePendingAggregationRequests", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePendingAggregationRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRemediationConfigurations", func(t *testing.T) {
        input := &configservice.DescribeRemediationConfigurationsInput{}
        output := &configservice.DescribeRemediationConfigurationsOutput{}

        mockClient.On("DescribeRemediationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRemediationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRemediationExceptions", func(t *testing.T) {
        input := &configservice.DescribeRemediationExceptionsInput{}
        output := &configservice.DescribeRemediationExceptionsOutput{}

        mockClient.On("DescribeRemediationExceptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRemediationExceptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRemediationExecutionStatus", func(t *testing.T) {
        input := &configservice.DescribeRemediationExecutionStatusInput{}
        output := &configservice.DescribeRemediationExecutionStatusOutput{}

        mockClient.On("DescribeRemediationExecutionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRemediationExecutionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRetentionConfigurations", func(t *testing.T) {
        input := &configservice.DescribeRetentionConfigurationsInput{}
        output := &configservice.DescribeRetentionConfigurationsOutput{}

        mockClient.On("DescribeRetentionConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRetentionConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAggregateComplianceDetailsByConfigRule", func(t *testing.T) {
        input := &configservice.GetAggregateComplianceDetailsByConfigRuleInput{}
        output := &configservice.GetAggregateComplianceDetailsByConfigRuleOutput{}

        mockClient.On("GetAggregateComplianceDetailsByConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetAggregateComplianceDetailsByConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAggregateConfigRuleComplianceSummary", func(t *testing.T) {
        input := &configservice.GetAggregateConfigRuleComplianceSummaryInput{}
        output := &configservice.GetAggregateConfigRuleComplianceSummaryOutput{}

        mockClient.On("GetAggregateConfigRuleComplianceSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetAggregateConfigRuleComplianceSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAggregateConformancePackComplianceSummary", func(t *testing.T) {
        input := &configservice.GetAggregateConformancePackComplianceSummaryInput{}
        output := &configservice.GetAggregateConformancePackComplianceSummaryOutput{}

        mockClient.On("GetAggregateConformancePackComplianceSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetAggregateConformancePackComplianceSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAggregateDiscoveredResourceCounts", func(t *testing.T) {
        input := &configservice.GetAggregateDiscoveredResourceCountsInput{}
        output := &configservice.GetAggregateDiscoveredResourceCountsOutput{}

        mockClient.On("GetAggregateDiscoveredResourceCounts", ctx, input).Return(output, nil)

        result, err := mockClient.GetAggregateDiscoveredResourceCounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAggregateResourceConfig", func(t *testing.T) {
        input := &configservice.GetAggregateResourceConfigInput{}
        output := &configservice.GetAggregateResourceConfigOutput{}

        mockClient.On("GetAggregateResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetAggregateResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceDetailsByConfigRule", func(t *testing.T) {
        input := &configservice.GetComplianceDetailsByConfigRuleInput{}
        output := &configservice.GetComplianceDetailsByConfigRuleOutput{}

        mockClient.On("GetComplianceDetailsByConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceDetailsByConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceDetailsByResource", func(t *testing.T) {
        input := &configservice.GetComplianceDetailsByResourceInput{}
        output := &configservice.GetComplianceDetailsByResourceOutput{}

        mockClient.On("GetComplianceDetailsByResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceDetailsByResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceSummaryByConfigRule", func(t *testing.T) {
        input := &configservice.GetComplianceSummaryByConfigRuleInput{}
        output := &configservice.GetComplianceSummaryByConfigRuleOutput{}

        mockClient.On("GetComplianceSummaryByConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceSummaryByConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceSummaryByResourceType", func(t *testing.T) {
        input := &configservice.GetComplianceSummaryByResourceTypeInput{}
        output := &configservice.GetComplianceSummaryByResourceTypeOutput{}

        mockClient.On("GetComplianceSummaryByResourceType", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceSummaryByResourceType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConformancePackComplianceDetails", func(t *testing.T) {
        input := &configservice.GetConformancePackComplianceDetailsInput{}
        output := &configservice.GetConformancePackComplianceDetailsOutput{}

        mockClient.On("GetConformancePackComplianceDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetConformancePackComplianceDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConformancePackComplianceSummary", func(t *testing.T) {
        input := &configservice.GetConformancePackComplianceSummaryInput{}
        output := &configservice.GetConformancePackComplianceSummaryOutput{}

        mockClient.On("GetConformancePackComplianceSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetConformancePackComplianceSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomRulePolicy", func(t *testing.T) {
        input := &configservice.GetCustomRulePolicyInput{}
        output := &configservice.GetCustomRulePolicyOutput{}

        mockClient.On("GetCustomRulePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomRulePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDiscoveredResourceCounts", func(t *testing.T) {
        input := &configservice.GetDiscoveredResourceCountsInput{}
        output := &configservice.GetDiscoveredResourceCountsOutput{}

        mockClient.On("GetDiscoveredResourceCounts", ctx, input).Return(output, nil)

        result, err := mockClient.GetDiscoveredResourceCounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationConfigRuleDetailedStatus", func(t *testing.T) {
        input := &configservice.GetOrganizationConfigRuleDetailedStatusInput{}
        output := &configservice.GetOrganizationConfigRuleDetailedStatusOutput{}

        mockClient.On("GetOrganizationConfigRuleDetailedStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationConfigRuleDetailedStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationConformancePackDetailedStatus", func(t *testing.T) {
        input := &configservice.GetOrganizationConformancePackDetailedStatusInput{}
        output := &configservice.GetOrganizationConformancePackDetailedStatusOutput{}

        mockClient.On("GetOrganizationConformancePackDetailedStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationConformancePackDetailedStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationCustomRulePolicy", func(t *testing.T) {
        input := &configservice.GetOrganizationCustomRulePolicyInput{}
        output := &configservice.GetOrganizationCustomRulePolicyOutput{}

        mockClient.On("GetOrganizationCustomRulePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationCustomRulePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceConfigHistory", func(t *testing.T) {
        input := &configservice.GetResourceConfigHistoryInput{}
        output := &configservice.GetResourceConfigHistoryOutput{}

        mockClient.On("GetResourceConfigHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceConfigHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceEvaluationSummary", func(t *testing.T) {
        input := &configservice.GetResourceEvaluationSummaryInput{}
        output := &configservice.GetResourceEvaluationSummaryOutput{}

        mockClient.On("GetResourceEvaluationSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceEvaluationSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStoredQuery", func(t *testing.T) {
        input := &configservice.GetStoredQueryInput{}
        output := &configservice.GetStoredQueryOutput{}

        mockClient.On("GetStoredQuery", ctx, input).Return(output, nil)

        result, err := mockClient.GetStoredQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAggregateDiscoveredResources", func(t *testing.T) {
        input := &configservice.ListAggregateDiscoveredResourcesInput{}
        output := &configservice.ListAggregateDiscoveredResourcesOutput{}

        mockClient.On("ListAggregateDiscoveredResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListAggregateDiscoveredResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConformancePackComplianceScores", func(t *testing.T) {
        input := &configservice.ListConformancePackComplianceScoresInput{}
        output := &configservice.ListConformancePackComplianceScoresOutput{}

        mockClient.On("ListConformancePackComplianceScores", ctx, input).Return(output, nil)

        result, err := mockClient.ListConformancePackComplianceScores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDiscoveredResources", func(t *testing.T) {
        input := &configservice.ListDiscoveredResourcesInput{}
        output := &configservice.ListDiscoveredResourcesOutput{}

        mockClient.On("ListDiscoveredResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDiscoveredResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceEvaluations", func(t *testing.T) {
        input := &configservice.ListResourceEvaluationsInput{}
        output := &configservice.ListResourceEvaluationsOutput{}

        mockClient.On("ListResourceEvaluations", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceEvaluations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStoredQueries", func(t *testing.T) {
        input := &configservice.ListStoredQueriesInput{}
        output := &configservice.ListStoredQueriesOutput{}

        mockClient.On("ListStoredQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListStoredQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &configservice.ListTagsForResourceInput{}
        output := &configservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAggregationAuthorization", func(t *testing.T) {
        input := &configservice.PutAggregationAuthorizationInput{}
        output := &configservice.PutAggregationAuthorizationOutput{}

        mockClient.On("PutAggregationAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.PutAggregationAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigRule", func(t *testing.T) {
        input := &configservice.PutConfigRuleInput{}
        output := &configservice.PutConfigRuleOutput{}

        mockClient.On("PutConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationAggregator", func(t *testing.T) {
        input := &configservice.PutConfigurationAggregatorInput{}
        output := &configservice.PutConfigurationAggregatorOutput{}

        mockClient.On("PutConfigurationAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationRecorder", func(t *testing.T) {
        input := &configservice.PutConfigurationRecorderInput{}
        output := &configservice.PutConfigurationRecorderOutput{}

        mockClient.On("PutConfigurationRecorder", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationRecorder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConformancePack", func(t *testing.T) {
        input := &configservice.PutConformancePackInput{}
        output := &configservice.PutConformancePackOutput{}

        mockClient.On("PutConformancePack", ctx, input).Return(output, nil)

        result, err := mockClient.PutConformancePack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliveryChannel", func(t *testing.T) {
        input := &configservice.PutDeliveryChannelInput{}
        output := &configservice.PutDeliveryChannelOutput{}

        mockClient.On("PutDeliveryChannel", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliveryChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEvaluations", func(t *testing.T) {
        input := &configservice.PutEvaluationsInput{}
        output := &configservice.PutEvaluationsOutput{}

        mockClient.On("PutEvaluations", ctx, input).Return(output, nil)

        result, err := mockClient.PutEvaluations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutExternalEvaluation", func(t *testing.T) {
        input := &configservice.PutExternalEvaluationInput{}
        output := &configservice.PutExternalEvaluationOutput{}

        mockClient.On("PutExternalEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.PutExternalEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOrganizationConfigRule", func(t *testing.T) {
        input := &configservice.PutOrganizationConfigRuleInput{}
        output := &configservice.PutOrganizationConfigRuleOutput{}

        mockClient.On("PutOrganizationConfigRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutOrganizationConfigRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOrganizationConformancePack", func(t *testing.T) {
        input := &configservice.PutOrganizationConformancePackInput{}
        output := &configservice.PutOrganizationConformancePackOutput{}

        mockClient.On("PutOrganizationConformancePack", ctx, input).Return(output, nil)

        result, err := mockClient.PutOrganizationConformancePack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRemediationConfigurations", func(t *testing.T) {
        input := &configservice.PutRemediationConfigurationsInput{}
        output := &configservice.PutRemediationConfigurationsOutput{}

        mockClient.On("PutRemediationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.PutRemediationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRemediationExceptions", func(t *testing.T) {
        input := &configservice.PutRemediationExceptionsInput{}
        output := &configservice.PutRemediationExceptionsOutput{}

        mockClient.On("PutRemediationExceptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutRemediationExceptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourceConfig", func(t *testing.T) {
        input := &configservice.PutResourceConfigInput{}
        output := &configservice.PutResourceConfigOutput{}

        mockClient.On("PutResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRetentionConfiguration", func(t *testing.T) {
        input := &configservice.PutRetentionConfigurationInput{}
        output := &configservice.PutRetentionConfigurationOutput{}

        mockClient.On("PutRetentionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutRetentionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutStoredQuery", func(t *testing.T) {
        input := &configservice.PutStoredQueryInput{}
        output := &configservice.PutStoredQueryOutput{}

        mockClient.On("PutStoredQuery", ctx, input).Return(output, nil)

        result, err := mockClient.PutStoredQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSelectAggregateResourceConfig", func(t *testing.T) {
        input := &configservice.SelectAggregateResourceConfigInput{}
        output := &configservice.SelectAggregateResourceConfigOutput{}

        mockClient.On("SelectAggregateResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.SelectAggregateResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSelectResourceConfig", func(t *testing.T) {
        input := &configservice.SelectResourceConfigInput{}
        output := &configservice.SelectResourceConfigOutput{}

        mockClient.On("SelectResourceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.SelectResourceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConfigRulesEvaluation", func(t *testing.T) {
        input := &configservice.StartConfigRulesEvaluationInput{}
        output := &configservice.StartConfigRulesEvaluationOutput{}

        mockClient.On("StartConfigRulesEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.StartConfigRulesEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConfigurationRecorder", func(t *testing.T) {
        input := &configservice.StartConfigurationRecorderInput{}
        output := &configservice.StartConfigurationRecorderOutput{}

        mockClient.On("StartConfigurationRecorder", ctx, input).Return(output, nil)

        result, err := mockClient.StartConfigurationRecorder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRemediationExecution", func(t *testing.T) {
        input := &configservice.StartRemediationExecutionInput{}
        output := &configservice.StartRemediationExecutionOutput{}

        mockClient.On("StartRemediationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartRemediationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartResourceEvaluation", func(t *testing.T) {
        input := &configservice.StartResourceEvaluationInput{}
        output := &configservice.StartResourceEvaluationOutput{}

        mockClient.On("StartResourceEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.StartResourceEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopConfigurationRecorder", func(t *testing.T) {
        input := &configservice.StopConfigurationRecorderInput{}
        output := &configservice.StopConfigurationRecorderOutput{}

        mockClient.On("StopConfigurationRecorder", ctx, input).Return(output, nil)

        result, err := mockClient.StopConfigurationRecorder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &configservice.TagResourceInput{}
        output := &configservice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &configservice.UntagResourceInput{}
        output := &configservice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
