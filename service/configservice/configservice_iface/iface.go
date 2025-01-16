
package configservice_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/configservice"
)

// IClient defines the interface for configservice
type IClient interface {
 Options() Options 
 AssociateResourceTypes(ctx context.Context, params *AssociateResourceTypesInput, optFns ...func(*Options)) (*AssociateResourceTypesOutput, error) 
 BatchGetAggregateResourceConfig(ctx context.Context, params *BatchGetAggregateResourceConfigInput, optFns ...func(*Options)) (*BatchGetAggregateResourceConfigOutput, error) 
 BatchGetResourceConfig(ctx context.Context, params *BatchGetResourceConfigInput, optFns ...func(*Options)) (*BatchGetResourceConfigOutput, error) 
 DeleteAggregationAuthorization(ctx context.Context, params *DeleteAggregationAuthorizationInput, optFns ...func(*Options)) (*DeleteAggregationAuthorizationOutput, error) 
 DeleteConfigRule(ctx context.Context, params *DeleteConfigRuleInput, optFns ...func(*Options)) (*DeleteConfigRuleOutput, error) 
 DeleteConfigurationAggregator(ctx context.Context, params *DeleteConfigurationAggregatorInput, optFns ...func(*Options)) (*DeleteConfigurationAggregatorOutput, error) 
 DeleteConfigurationRecorder(ctx context.Context, params *DeleteConfigurationRecorderInput, optFns ...func(*Options)) (*DeleteConfigurationRecorderOutput, error) 
 DeleteConformancePack(ctx context.Context, params *DeleteConformancePackInput, optFns ...func(*Options)) (*DeleteConformancePackOutput, error) 
 DeleteDeliveryChannel(ctx context.Context, params *DeleteDeliveryChannelInput, optFns ...func(*Options)) (*DeleteDeliveryChannelOutput, error) 
 DeleteEvaluationResults(ctx context.Context, params *DeleteEvaluationResultsInput, optFns ...func(*Options)) (*DeleteEvaluationResultsOutput, error) 
 DeleteOrganizationConfigRule(ctx context.Context, params *DeleteOrganizationConfigRuleInput, optFns ...func(*Options)) (*DeleteOrganizationConfigRuleOutput, error) 
 DeleteOrganizationConformancePack(ctx context.Context, params *DeleteOrganizationConformancePackInput, optFns ...func(*Options)) (*DeleteOrganizationConformancePackOutput, error) 
 DeletePendingAggregationRequest(ctx context.Context, params *DeletePendingAggregationRequestInput, optFns ...func(*Options)) (*DeletePendingAggregationRequestOutput, error) 
 DeleteRemediationConfiguration(ctx context.Context, params *DeleteRemediationConfigurationInput, optFns ...func(*Options)) (*DeleteRemediationConfigurationOutput, error) 
 DeleteRemediationExceptions(ctx context.Context, params *DeleteRemediationExceptionsInput, optFns ...func(*Options)) (*DeleteRemediationExceptionsOutput, error) 
 DeleteResourceConfig(ctx context.Context, params *DeleteResourceConfigInput, optFns ...func(*Options)) (*DeleteResourceConfigOutput, error) 
 DeleteRetentionConfiguration(ctx context.Context, params *DeleteRetentionConfigurationInput, optFns ...func(*Options)) (*DeleteRetentionConfigurationOutput, error) 
 DeleteServiceLinkedConfigurationRecorder(ctx context.Context, params *DeleteServiceLinkedConfigurationRecorderInput, optFns ...func(*Options)) (*DeleteServiceLinkedConfigurationRecorderOutput, error) 
 DeleteStoredQuery(ctx context.Context, params *DeleteStoredQueryInput, optFns ...func(*Options)) (*DeleteStoredQueryOutput, error) 
 DeliverConfigSnapshot(ctx context.Context, params *DeliverConfigSnapshotInput, optFns ...func(*Options)) (*DeliverConfigSnapshotOutput, error) 
 DescribeAggregateComplianceByConfigRules(ctx context.Context, params *DescribeAggregateComplianceByConfigRulesInput, optFns ...func(*Options)) (*DescribeAggregateComplianceByConfigRulesOutput, error) 
 DescribeAggregateComplianceByConformancePacks(ctx context.Context, params *DescribeAggregateComplianceByConformancePacksInput, optFns ...func(*Options)) (*DescribeAggregateComplianceByConformancePacksOutput, error) 
 DescribeAggregationAuthorizations(ctx context.Context, params *DescribeAggregationAuthorizationsInput, optFns ...func(*Options)) (*DescribeAggregationAuthorizationsOutput, error) 
 DescribeComplianceByConfigRule(ctx context.Context, params *DescribeComplianceByConfigRuleInput, optFns ...func(*Options)) (*DescribeComplianceByConfigRuleOutput, error) 
 DescribeComplianceByResource(ctx context.Context, params *DescribeComplianceByResourceInput, optFns ...func(*Options)) (*DescribeComplianceByResourceOutput, error) 
 DescribeConfigRuleEvaluationStatus(ctx context.Context, params *DescribeConfigRuleEvaluationStatusInput, optFns ...func(*Options)) (*DescribeConfigRuleEvaluationStatusOutput, error) 
 DescribeConfigRules(ctx context.Context, params *DescribeConfigRulesInput, optFns ...func(*Options)) (*DescribeConfigRulesOutput, error) 
 DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) 
 DescribeConfigurationAggregators(ctx context.Context, params *DescribeConfigurationAggregatorsInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorsOutput, error) 
 DescribeConfigurationRecorderStatus(ctx context.Context, params *DescribeConfigurationRecorderStatusInput, optFns ...func(*Options)) (*DescribeConfigurationRecorderStatusOutput, error) 
 DescribeConfigurationRecorders(ctx context.Context, params *DescribeConfigurationRecordersInput, optFns ...func(*Options)) (*DescribeConfigurationRecordersOutput, error) 
 DescribeConformancePackCompliance(ctx context.Context, params *DescribeConformancePackComplianceInput, optFns ...func(*Options)) (*DescribeConformancePackComplianceOutput, error) 
 DescribeConformancePackStatus(ctx context.Context, params *DescribeConformancePackStatusInput, optFns ...func(*Options)) (*DescribeConformancePackStatusOutput, error) 
 DescribeConformancePacks(ctx context.Context, params *DescribeConformancePacksInput, optFns ...func(*Options)) (*DescribeConformancePacksOutput, error) 
 DescribeDeliveryChannelStatus(ctx context.Context, params *DescribeDeliveryChannelStatusInput, optFns ...func(*Options)) (*DescribeDeliveryChannelStatusOutput, error) 
 DescribeDeliveryChannels(ctx context.Context, params *DescribeDeliveryChannelsInput, optFns ...func(*Options)) (*DescribeDeliveryChannelsOutput, error) 
 DescribeOrganizationConfigRuleStatuses(ctx context.Context, params *DescribeOrganizationConfigRuleStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConfigRuleStatusesOutput, error) 
 DescribeOrganizationConfigRules(ctx context.Context, params *DescribeOrganizationConfigRulesInput, optFns ...func(*Options)) (*DescribeOrganizationConfigRulesOutput, error) 
 DescribeOrganizationConformancePackStatuses(ctx context.Context, params *DescribeOrganizationConformancePackStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error) 
 DescribeOrganizationConformancePacks(ctx context.Context, params *DescribeOrganizationConformancePacksInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePacksOutput, error) 
 DescribePendingAggregationRequests(ctx context.Context, params *DescribePendingAggregationRequestsInput, optFns ...func(*Options)) (*DescribePendingAggregationRequestsOutput, error) 
 DescribeRemediationConfigurations(ctx context.Context, params *DescribeRemediationConfigurationsInput, optFns ...func(*Options)) (*DescribeRemediationConfigurationsOutput, error) 
 DescribeRemediationExceptions(ctx context.Context, params *DescribeRemediationExceptionsInput, optFns ...func(*Options)) (*DescribeRemediationExceptionsOutput, error) 
 DescribeRemediationExecutionStatus(ctx context.Context, params *DescribeRemediationExecutionStatusInput, optFns ...func(*Options)) (*DescribeRemediationExecutionStatusOutput, error) 
 DescribeRetentionConfigurations(ctx context.Context, params *DescribeRetentionConfigurationsInput, optFns ...func(*Options)) (*DescribeRetentionConfigurationsOutput, error) 
 DisassociateResourceTypes(ctx context.Context, params *DisassociateResourceTypesInput, optFns ...func(*Options)) (*DisassociateResourceTypesOutput, error) 
 GetAggregateComplianceDetailsByConfigRule(ctx context.Context, params *GetAggregateComplianceDetailsByConfigRuleInput, optFns ...func(*Options)) (*GetAggregateComplianceDetailsByConfigRuleOutput, error) 
 GetAggregateConfigRuleComplianceSummary(ctx context.Context, params *GetAggregateConfigRuleComplianceSummaryInput, optFns ...func(*Options)) (*GetAggregateConfigRuleComplianceSummaryOutput, error) 
 GetAggregateConformancePackComplianceSummary(ctx context.Context, params *GetAggregateConformancePackComplianceSummaryInput, optFns ...func(*Options)) (*GetAggregateConformancePackComplianceSummaryOutput, error) 
 GetAggregateDiscoveredResourceCounts(ctx context.Context, params *GetAggregateDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetAggregateDiscoveredResourceCountsOutput, error) 
 GetAggregateResourceConfig(ctx context.Context, params *GetAggregateResourceConfigInput, optFns ...func(*Options)) (*GetAggregateResourceConfigOutput, error) 
 GetComplianceDetailsByConfigRule(ctx context.Context, params *GetComplianceDetailsByConfigRuleInput, optFns ...func(*Options)) (*GetComplianceDetailsByConfigRuleOutput, error) 
 GetComplianceDetailsByResource(ctx context.Context, params *GetComplianceDetailsByResourceInput, optFns ...func(*Options)) (*GetComplianceDetailsByResourceOutput, error) 
 GetComplianceSummaryByConfigRule(ctx context.Context, params *GetComplianceSummaryByConfigRuleInput, optFns ...func(*Options)) (*GetComplianceSummaryByConfigRuleOutput, error) 
 GetComplianceSummaryByResourceType(ctx context.Context, params *GetComplianceSummaryByResourceTypeInput, optFns ...func(*Options)) (*GetComplianceSummaryByResourceTypeOutput, error) 
 GetConformancePackComplianceDetails(ctx context.Context, params *GetConformancePackComplianceDetailsInput, optFns ...func(*Options)) (*GetConformancePackComplianceDetailsOutput, error) 
 GetConformancePackComplianceSummary(ctx context.Context, params *GetConformancePackComplianceSummaryInput, optFns ...func(*Options)) (*GetConformancePackComplianceSummaryOutput, error) 
 GetCustomRulePolicy(ctx context.Context, params *GetCustomRulePolicyInput, optFns ...func(*Options)) (*GetCustomRulePolicyOutput, error) 
 GetDiscoveredResourceCounts(ctx context.Context, params *GetDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error) 
 GetOrganizationConfigRuleDetailedStatus(ctx context.Context, params *GetOrganizationConfigRuleDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConfigRuleDetailedStatusOutput, error) 
 GetOrganizationConformancePackDetailedStatus(ctx context.Context, params *GetOrganizationConformancePackDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error) 
 GetOrganizationCustomRulePolicy(ctx context.Context, params *GetOrganizationCustomRulePolicyInput, optFns ...func(*Options)) (*GetOrganizationCustomRulePolicyOutput, error) 
 GetResourceConfigHistory(ctx context.Context, params *GetResourceConfigHistoryInput, optFns ...func(*Options)) (*GetResourceConfigHistoryOutput, error) 
 GetResourceEvaluationSummary(ctx context.Context, params *GetResourceEvaluationSummaryInput, optFns ...func(*Options)) (*GetResourceEvaluationSummaryOutput, error) 
 GetStoredQuery(ctx context.Context, params *GetStoredQueryInput, optFns ...func(*Options)) (*GetStoredQueryOutput, error) 
 ListAggregateDiscoveredResources(ctx context.Context, params *ListAggregateDiscoveredResourcesInput, optFns ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error) 
 ListConfigurationRecorders(ctx context.Context, params *ListConfigurationRecordersInput, optFns ...func(*Options)) (*ListConfigurationRecordersOutput, error) 
 ListConformancePackComplianceScores(ctx context.Context, params *ListConformancePackComplianceScoresInput, optFns ...func(*Options)) (*ListConformancePackComplianceScoresOutput, error) 
 ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) 
 ListResourceEvaluations(ctx context.Context, params *ListResourceEvaluationsInput, optFns ...func(*Options)) (*ListResourceEvaluationsOutput, error) 
 ListStoredQueries(ctx context.Context, params *ListStoredQueriesInput, optFns ...func(*Options)) (*ListStoredQueriesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAggregationAuthorization(ctx context.Context, params *PutAggregationAuthorizationInput, optFns ...func(*Options)) (*PutAggregationAuthorizationOutput, error) 
 PutConfigRule(ctx context.Context, params *PutConfigRuleInput, optFns ...func(*Options)) (*PutConfigRuleOutput, error) 
 PutConfigurationAggregator(ctx context.Context, params *PutConfigurationAggregatorInput, optFns ...func(*Options)) (*PutConfigurationAggregatorOutput, error) 
 PutConfigurationRecorder(ctx context.Context, params *PutConfigurationRecorderInput, optFns ...func(*Options)) (*PutConfigurationRecorderOutput, error) 
 PutConformancePack(ctx context.Context, params *PutConformancePackInput, optFns ...func(*Options)) (*PutConformancePackOutput, error) 
 PutDeliveryChannel(ctx context.Context, params *PutDeliveryChannelInput, optFns ...func(*Options)) (*PutDeliveryChannelOutput, error) 
 PutEvaluations(ctx context.Context, params *PutEvaluationsInput, optFns ...func(*Options)) (*PutEvaluationsOutput, error) 
 PutExternalEvaluation(ctx context.Context, params *PutExternalEvaluationInput, optFns ...func(*Options)) (*PutExternalEvaluationOutput, error) 
 PutOrganizationConfigRule(ctx context.Context, params *PutOrganizationConfigRuleInput, optFns ...func(*Options)) (*PutOrganizationConfigRuleOutput, error) 
 PutOrganizationConformancePack(ctx context.Context, params *PutOrganizationConformancePackInput, optFns ...func(*Options)) (*PutOrganizationConformancePackOutput, error) 
 PutRemediationConfigurations(ctx context.Context, params *PutRemediationConfigurationsInput, optFns ...func(*Options)) (*PutRemediationConfigurationsOutput, error) 
 PutRemediationExceptions(ctx context.Context, params *PutRemediationExceptionsInput, optFns ...func(*Options)) (*PutRemediationExceptionsOutput, error) 
 PutResourceConfig(ctx context.Context, params *PutResourceConfigInput, optFns ...func(*Options)) (*PutResourceConfigOutput, error) 
 PutRetentionConfiguration(ctx context.Context, params *PutRetentionConfigurationInput, optFns ...func(*Options)) (*PutRetentionConfigurationOutput, error) 
 PutServiceLinkedConfigurationRecorder(ctx context.Context, params *PutServiceLinkedConfigurationRecorderInput, optFns ...func(*Options)) (*PutServiceLinkedConfigurationRecorderOutput, error) 
 PutStoredQuery(ctx context.Context, params *PutStoredQueryInput, optFns ...func(*Options)) (*PutStoredQueryOutput, error) 
 SelectAggregateResourceConfig(ctx context.Context, params *SelectAggregateResourceConfigInput, optFns ...func(*Options)) (*SelectAggregateResourceConfigOutput, error) 
 SelectResourceConfig(ctx context.Context, params *SelectResourceConfigInput, optFns ...func(*Options)) (*SelectResourceConfigOutput, error) 
 StartConfigRulesEvaluation(ctx context.Context, params *StartConfigRulesEvaluationInput, optFns ...func(*Options)) (*StartConfigRulesEvaluationOutput, error) 
 StartConfigurationRecorder(ctx context.Context, params *StartConfigurationRecorderInput, optFns ...func(*Options)) (*StartConfigurationRecorderOutput, error) 
 StartRemediationExecution(ctx context.Context, params *StartRemediationExecutionInput, optFns ...func(*Options)) (*StartRemediationExecutionOutput, error) 
 StartResourceEvaluation(ctx context.Context, params *StartResourceEvaluationInput, optFns ...func(*Options)) (*StartResourceEvaluationOutput, error) 
 StopConfigurationRecorder(ctx context.Context, params *StopConfigurationRecorderInput, optFns ...func(*Options)) (*StopConfigurationRecorderOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
