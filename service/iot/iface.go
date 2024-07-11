
package iot

import (
    "github.com/aws/aws-sdk-go-v2/service/iot"
)

// IIot defines the interface for iot
type IIot interface {
 Options() Options 
 AcceptCertificateTransfer(ctx context.Context, params *AcceptCertificateTransferInput, optFns ...func(*Options)) (*AcceptCertificateTransferOutput, error) 
 AddThingToBillingGroup(ctx context.Context, params *AddThingToBillingGroupInput, optFns ...func(*Options)) (*AddThingToBillingGroupOutput, error) 
 AddThingToThingGroup(ctx context.Context, params *AddThingToThingGroupInput, optFns ...func(*Options)) (*AddThingToThingGroupOutput, error) 
 AssociateTargetsWithJob(ctx context.Context, params *AssociateTargetsWithJobInput, optFns ...func(*Options)) (*AssociateTargetsWithJobOutput, error) 
 AttachPolicy(ctx context.Context, params *AttachPolicyInput, optFns ...func(*Options)) (*AttachPolicyOutput, error) 
 AttachPrincipalPolicy(ctx context.Context, params *AttachPrincipalPolicyInput, optFns ...func(*Options)) (*AttachPrincipalPolicyOutput, error) 
 AttachSecurityProfile(ctx context.Context, params *AttachSecurityProfileInput, optFns ...func(*Options)) (*AttachSecurityProfileOutput, error) 
 AttachThingPrincipal(ctx context.Context, params *AttachThingPrincipalInput, optFns ...func(*Options)) (*AttachThingPrincipalOutput, error) 
 CancelAuditMitigationActionsTask(ctx context.Context, params *CancelAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*CancelAuditMitigationActionsTaskOutput, error) 
 CancelAuditTask(ctx context.Context, params *CancelAuditTaskInput, optFns ...func(*Options)) (*CancelAuditTaskOutput, error) 
 CancelCertificateTransfer(ctx context.Context, params *CancelCertificateTransferInput, optFns ...func(*Options)) (*CancelCertificateTransferOutput, error) 
 CancelDetectMitigationActionsTask(ctx context.Context, params *CancelDetectMitigationActionsTaskInput, optFns ...func(*Options)) (*CancelDetectMitigationActionsTaskOutput, error) 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CancelJobExecution(ctx context.Context, params *CancelJobExecutionInput, optFns ...func(*Options)) (*CancelJobExecutionOutput, error) 
 ClearDefaultAuthorizer(ctx context.Context, params *ClearDefaultAuthorizerInput, optFns ...func(*Options)) (*ClearDefaultAuthorizerOutput, error) 
 ConfirmTopicRuleDestination(ctx context.Context, params *ConfirmTopicRuleDestinationInput, optFns ...func(*Options)) (*ConfirmTopicRuleDestinationOutput, error) 
 CreateAuditSuppression(ctx context.Context, params *CreateAuditSuppressionInput, optFns ...func(*Options)) (*CreateAuditSuppressionOutput, error) 
 CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) 
 CreateBillingGroup(ctx context.Context, params *CreateBillingGroupInput, optFns ...func(*Options)) (*CreateBillingGroupOutput, error) 
 CreateCertificateFromCsr(ctx context.Context, params *CreateCertificateFromCsrInput, optFns ...func(*Options)) (*CreateCertificateFromCsrOutput, error) 
 CreateCertificateProvider(ctx context.Context, params *CreateCertificateProviderInput, optFns ...func(*Options)) (*CreateCertificateProviderOutput, error) 
 CreateCustomMetric(ctx context.Context, params *CreateCustomMetricInput, optFns ...func(*Options)) (*CreateCustomMetricOutput, error) 
 CreateDimension(ctx context.Context, params *CreateDimensionInput, optFns ...func(*Options)) (*CreateDimensionOutput, error) 
 CreateDomainConfiguration(ctx context.Context, params *CreateDomainConfigurationInput, optFns ...func(*Options)) (*CreateDomainConfigurationOutput, error) 
 CreateDynamicThingGroup(ctx context.Context, params *CreateDynamicThingGroupInput, optFns ...func(*Options)) (*CreateDynamicThingGroupOutput, error) 
 CreateFleetMetric(ctx context.Context, params *CreateFleetMetricInput, optFns ...func(*Options)) (*CreateFleetMetricOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) 
 CreateKeysAndCertificate(ctx context.Context, params *CreateKeysAndCertificateInput, optFns ...func(*Options)) (*CreateKeysAndCertificateOutput, error) 
 CreateMitigationAction(ctx context.Context, params *CreateMitigationActionInput, optFns ...func(*Options)) (*CreateMitigationActionOutput, error) 
 CreateOTAUpdate(ctx context.Context, params *CreateOTAUpdateInput, optFns ...func(*Options)) (*CreateOTAUpdateOutput, error) 
 CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) 
 CreatePackageVersion(ctx context.Context, params *CreatePackageVersionInput, optFns ...func(*Options)) (*CreatePackageVersionOutput, error) 
 CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) 
 CreatePolicyVersion(ctx context.Context, params *CreatePolicyVersionInput, optFns ...func(*Options)) (*CreatePolicyVersionOutput, error) 
 CreateProvisioningClaim(ctx context.Context, params *CreateProvisioningClaimInput, optFns ...func(*Options)) (*CreateProvisioningClaimOutput, error) 
 CreateProvisioningTemplate(ctx context.Context, params *CreateProvisioningTemplateInput, optFns ...func(*Options)) (*CreateProvisioningTemplateOutput, error) 
 CreateProvisioningTemplateVersion(ctx context.Context, params *CreateProvisioningTemplateVersionInput, optFns ...func(*Options)) (*CreateProvisioningTemplateVersionOutput, error) 
 CreateRoleAlias(ctx context.Context, params *CreateRoleAliasInput, optFns ...func(*Options)) (*CreateRoleAliasOutput, error) 
 CreateScheduledAudit(ctx context.Context, params *CreateScheduledAuditInput, optFns ...func(*Options)) (*CreateScheduledAuditOutput, error) 
 CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) 
 CreateStream(ctx context.Context, params *CreateStreamInput, optFns ...func(*Options)) (*CreateStreamOutput, error) 
 CreateThing(ctx context.Context, params *CreateThingInput, optFns ...func(*Options)) (*CreateThingOutput, error) 
 CreateThingGroup(ctx context.Context, params *CreateThingGroupInput, optFns ...func(*Options)) (*CreateThingGroupOutput, error) 
 CreateThingType(ctx context.Context, params *CreateThingTypeInput, optFns ...func(*Options)) (*CreateThingTypeOutput, error) 
 CreateTopicRule(ctx context.Context, params *CreateTopicRuleInput, optFns ...func(*Options)) (*CreateTopicRuleOutput, error) 
 CreateTopicRuleDestination(ctx context.Context, params *CreateTopicRuleDestinationInput, optFns ...func(*Options)) (*CreateTopicRuleDestinationOutput, error) 
 DeleteAccountAuditConfiguration(ctx context.Context, params *DeleteAccountAuditConfigurationInput, optFns ...func(*Options)) (*DeleteAccountAuditConfigurationOutput, error) 
 DeleteAuditSuppression(ctx context.Context, params *DeleteAuditSuppressionInput, optFns ...func(*Options)) (*DeleteAuditSuppressionOutput, error) 
 DeleteAuthorizer(ctx context.Context, params *DeleteAuthorizerInput, optFns ...func(*Options)) (*DeleteAuthorizerOutput, error) 
 DeleteBillingGroup(ctx context.Context, params *DeleteBillingGroupInput, optFns ...func(*Options)) (*DeleteBillingGroupOutput, error) 
 DeleteCACertificate(ctx context.Context, params *DeleteCACertificateInput, optFns ...func(*Options)) (*DeleteCACertificateOutput, error) 
 DeleteCertificate(ctx context.Context, params *DeleteCertificateInput, optFns ...func(*Options)) (*DeleteCertificateOutput, error) 
 DeleteCertificateProvider(ctx context.Context, params *DeleteCertificateProviderInput, optFns ...func(*Options)) (*DeleteCertificateProviderOutput, error) 
 DeleteCustomMetric(ctx context.Context, params *DeleteCustomMetricInput, optFns ...func(*Options)) (*DeleteCustomMetricOutput, error) 
 DeleteDimension(ctx context.Context, params *DeleteDimensionInput, optFns ...func(*Options)) (*DeleteDimensionOutput, error) 
 DeleteDomainConfiguration(ctx context.Context, params *DeleteDomainConfigurationInput, optFns ...func(*Options)) (*DeleteDomainConfigurationOutput, error) 
 DeleteDynamicThingGroup(ctx context.Context, params *DeleteDynamicThingGroupInput, optFns ...func(*Options)) (*DeleteDynamicThingGroupOutput, error) 
 DeleteFleetMetric(ctx context.Context, params *DeleteFleetMetricInput, optFns ...func(*Options)) (*DeleteFleetMetricOutput, error) 
 DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) 
 DeleteJobExecution(ctx context.Context, params *DeleteJobExecutionInput, optFns ...func(*Options)) (*DeleteJobExecutionOutput, error) 
 DeleteJobTemplate(ctx context.Context, params *DeleteJobTemplateInput, optFns ...func(*Options)) (*DeleteJobTemplateOutput, error) 
 DeleteMitigationAction(ctx context.Context, params *DeleteMitigationActionInput, optFns ...func(*Options)) (*DeleteMitigationActionOutput, error) 
 DeleteOTAUpdate(ctx context.Context, params *DeleteOTAUpdateInput, optFns ...func(*Options)) (*DeleteOTAUpdateOutput, error) 
 DeletePackage(ctx context.Context, params *DeletePackageInput, optFns ...func(*Options)) (*DeletePackageOutput, error) 
 DeletePackageVersion(ctx context.Context, params *DeletePackageVersionInput, optFns ...func(*Options)) (*DeletePackageVersionOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DeletePolicyVersion(ctx context.Context, params *DeletePolicyVersionInput, optFns ...func(*Options)) (*DeletePolicyVersionOutput, error) 
 DeleteProvisioningTemplate(ctx context.Context, params *DeleteProvisioningTemplateInput, optFns ...func(*Options)) (*DeleteProvisioningTemplateOutput, error) 
 DeleteProvisioningTemplateVersion(ctx context.Context, params *DeleteProvisioningTemplateVersionInput, optFns ...func(*Options)) (*DeleteProvisioningTemplateVersionOutput, error) 
 DeleteRegistrationCode(ctx context.Context, params *DeleteRegistrationCodeInput, optFns ...func(*Options)) (*DeleteRegistrationCodeOutput, error) 
 DeleteRoleAlias(ctx context.Context, params *DeleteRoleAliasInput, optFns ...func(*Options)) (*DeleteRoleAliasOutput, error) 
 DeleteScheduledAudit(ctx context.Context, params *DeleteScheduledAuditInput, optFns ...func(*Options)) (*DeleteScheduledAuditOutput, error) 
 DeleteSecurityProfile(ctx context.Context, params *DeleteSecurityProfileInput, optFns ...func(*Options)) (*DeleteSecurityProfileOutput, error) 
 DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) 
 DeleteThing(ctx context.Context, params *DeleteThingInput, optFns ...func(*Options)) (*DeleteThingOutput, error) 
 DeleteThingGroup(ctx context.Context, params *DeleteThingGroupInput, optFns ...func(*Options)) (*DeleteThingGroupOutput, error) 
 DeleteThingType(ctx context.Context, params *DeleteThingTypeInput, optFns ...func(*Options)) (*DeleteThingTypeOutput, error) 
 DeleteTopicRule(ctx context.Context, params *DeleteTopicRuleInput, optFns ...func(*Options)) (*DeleteTopicRuleOutput, error) 
 DeleteTopicRuleDestination(ctx context.Context, params *DeleteTopicRuleDestinationInput, optFns ...func(*Options)) (*DeleteTopicRuleDestinationOutput, error) 
 DeleteV2LoggingLevel(ctx context.Context, params *DeleteV2LoggingLevelInput, optFns ...func(*Options)) (*DeleteV2LoggingLevelOutput, error) 
 DeprecateThingType(ctx context.Context, params *DeprecateThingTypeInput, optFns ...func(*Options)) (*DeprecateThingTypeOutput, error) 
 DescribeAccountAuditConfiguration(ctx context.Context, params *DescribeAccountAuditConfigurationInput, optFns ...func(*Options)) (*DescribeAccountAuditConfigurationOutput, error) 
 DescribeAuditFinding(ctx context.Context, params *DescribeAuditFindingInput, optFns ...func(*Options)) (*DescribeAuditFindingOutput, error) 
 DescribeAuditMitigationActionsTask(ctx context.Context, params *DescribeAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*DescribeAuditMitigationActionsTaskOutput, error) 
 DescribeAuditSuppression(ctx context.Context, params *DescribeAuditSuppressionInput, optFns ...func(*Options)) (*DescribeAuditSuppressionOutput, error) 
 DescribeAuditTask(ctx context.Context, params *DescribeAuditTaskInput, optFns ...func(*Options)) (*DescribeAuditTaskOutput, error) 
 DescribeAuthorizer(ctx context.Context, params *DescribeAuthorizerInput, optFns ...func(*Options)) (*DescribeAuthorizerOutput, error) 
 DescribeBillingGroup(ctx context.Context, params *DescribeBillingGroupInput, optFns ...func(*Options)) (*DescribeBillingGroupOutput, error) 
 DescribeCACertificate(ctx context.Context, params *DescribeCACertificateInput, optFns ...func(*Options)) (*DescribeCACertificateOutput, error) 
 DescribeCertificate(ctx context.Context, params *DescribeCertificateInput, optFns ...func(*Options)) (*DescribeCertificateOutput, error) 
 DescribeCertificateProvider(ctx context.Context, params *DescribeCertificateProviderInput, optFns ...func(*Options)) (*DescribeCertificateProviderOutput, error) 
 DescribeCustomMetric(ctx context.Context, params *DescribeCustomMetricInput, optFns ...func(*Options)) (*DescribeCustomMetricOutput, error) 
 DescribeDefaultAuthorizer(ctx context.Context, params *DescribeDefaultAuthorizerInput, optFns ...func(*Options)) (*DescribeDefaultAuthorizerOutput, error) 
 DescribeDetectMitigationActionsTask(ctx context.Context, params *DescribeDetectMitigationActionsTaskInput, optFns ...func(*Options)) (*DescribeDetectMitigationActionsTaskOutput, error) 
 DescribeDimension(ctx context.Context, params *DescribeDimensionInput, optFns ...func(*Options)) (*DescribeDimensionOutput, error) 
 DescribeDomainConfiguration(ctx context.Context, params *DescribeDomainConfigurationInput, optFns ...func(*Options)) (*DescribeDomainConfigurationOutput, error) 
 DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) 
 DescribeEventConfigurations(ctx context.Context, params *DescribeEventConfigurationsInput, optFns ...func(*Options)) (*DescribeEventConfigurationsOutput, error) 
 DescribeFleetMetric(ctx context.Context, params *DescribeFleetMetricInput, optFns ...func(*Options)) (*DescribeFleetMetricOutput, error) 
 DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) 
 DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) 
 DescribeJobExecution(ctx context.Context, params *DescribeJobExecutionInput, optFns ...func(*Options)) (*DescribeJobExecutionOutput, error) 
 DescribeJobTemplate(ctx context.Context, params *DescribeJobTemplateInput, optFns ...func(*Options)) (*DescribeJobTemplateOutput, error) 
 DescribeManagedJobTemplate(ctx context.Context, params *DescribeManagedJobTemplateInput, optFns ...func(*Options)) (*DescribeManagedJobTemplateOutput, error) 
 DescribeMitigationAction(ctx context.Context, params *DescribeMitigationActionInput, optFns ...func(*Options)) (*DescribeMitigationActionOutput, error) 
 DescribeProvisioningTemplate(ctx context.Context, params *DescribeProvisioningTemplateInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateOutput, error) 
 DescribeProvisioningTemplateVersion(ctx context.Context, params *DescribeProvisioningTemplateVersionInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateVersionOutput, error) 
 DescribeRoleAlias(ctx context.Context, params *DescribeRoleAliasInput, optFns ...func(*Options)) (*DescribeRoleAliasOutput, error) 
 DescribeScheduledAudit(ctx context.Context, params *DescribeScheduledAuditInput, optFns ...func(*Options)) (*DescribeScheduledAuditOutput, error) 
 DescribeSecurityProfile(ctx context.Context, params *DescribeSecurityProfileInput, optFns ...func(*Options)) (*DescribeSecurityProfileOutput, error) 
 DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) 
 DescribeThing(ctx context.Context, params *DescribeThingInput, optFns ...func(*Options)) (*DescribeThingOutput, error) 
 DescribeThingGroup(ctx context.Context, params *DescribeThingGroupInput, optFns ...func(*Options)) (*DescribeThingGroupOutput, error) 
 DescribeThingRegistrationTask(ctx context.Context, params *DescribeThingRegistrationTaskInput, optFns ...func(*Options)) (*DescribeThingRegistrationTaskOutput, error) 
 DescribeThingType(ctx context.Context, params *DescribeThingTypeInput, optFns ...func(*Options)) (*DescribeThingTypeOutput, error) 
 DetachPolicy(ctx context.Context, params *DetachPolicyInput, optFns ...func(*Options)) (*DetachPolicyOutput, error) 
 DetachPrincipalPolicy(ctx context.Context, params *DetachPrincipalPolicyInput, optFns ...func(*Options)) (*DetachPrincipalPolicyOutput, error) 
 DetachSecurityProfile(ctx context.Context, params *DetachSecurityProfileInput, optFns ...func(*Options)) (*DetachSecurityProfileOutput, error) 
 DetachThingPrincipal(ctx context.Context, params *DetachThingPrincipalInput, optFns ...func(*Options)) (*DetachThingPrincipalOutput, error) 
 DisableTopicRule(ctx context.Context, params *DisableTopicRuleInput, optFns ...func(*Options)) (*DisableTopicRuleOutput, error) 
 EnableTopicRule(ctx context.Context, params *EnableTopicRuleInput, optFns ...func(*Options)) (*EnableTopicRuleOutput, error) 
 GetBehaviorModelTrainingSummaries(ctx context.Context, params *GetBehaviorModelTrainingSummariesInput, optFns ...func(*Options)) (*GetBehaviorModelTrainingSummariesOutput, error) 
 GetBucketsAggregation(ctx context.Context, params *GetBucketsAggregationInput, optFns ...func(*Options)) (*GetBucketsAggregationOutput, error) 
 GetCardinality(ctx context.Context, params *GetCardinalityInput, optFns ...func(*Options)) (*GetCardinalityOutput, error) 
 GetEffectivePolicies(ctx context.Context, params *GetEffectivePoliciesInput, optFns ...func(*Options)) (*GetEffectivePoliciesOutput, error) 
 GetIndexingConfiguration(ctx context.Context, params *GetIndexingConfigurationInput, optFns ...func(*Options)) (*GetIndexingConfigurationOutput, error) 
 GetJobDocument(ctx context.Context, params *GetJobDocumentInput, optFns ...func(*Options)) (*GetJobDocumentOutput, error) 
 GetLoggingOptions(ctx context.Context, params *GetLoggingOptionsInput, optFns ...func(*Options)) (*GetLoggingOptionsOutput, error) 
 GetOTAUpdate(ctx context.Context, params *GetOTAUpdateInput, optFns ...func(*Options)) (*GetOTAUpdateOutput, error) 
 GetPackage(ctx context.Context, params *GetPackageInput, optFns ...func(*Options)) (*GetPackageOutput, error) 
 GetPackageConfiguration(ctx context.Context, params *GetPackageConfigurationInput, optFns ...func(*Options)) (*GetPackageConfigurationOutput, error) 
 GetPackageVersion(ctx context.Context, params *GetPackageVersionInput, optFns ...func(*Options)) (*GetPackageVersionOutput, error) 
 GetPercentiles(ctx context.Context, params *GetPercentilesInput, optFns ...func(*Options)) (*GetPercentilesOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetPolicyVersion(ctx context.Context, params *GetPolicyVersionInput, optFns ...func(*Options)) (*GetPolicyVersionOutput, error) 
 GetRegistrationCode(ctx context.Context, params *GetRegistrationCodeInput, optFns ...func(*Options)) (*GetRegistrationCodeOutput, error) 
 GetStatistics(ctx context.Context, params *GetStatisticsInput, optFns ...func(*Options)) (*GetStatisticsOutput, error) 
 GetTopicRule(ctx context.Context, params *GetTopicRuleInput, optFns ...func(*Options)) (*GetTopicRuleOutput, error) 
 GetTopicRuleDestination(ctx context.Context, params *GetTopicRuleDestinationInput, optFns ...func(*Options)) (*GetTopicRuleDestinationOutput, error) 
 GetV2LoggingOptions(ctx context.Context, params *GetV2LoggingOptionsInput, optFns ...func(*Options)) (*GetV2LoggingOptionsOutput, error) 
 ListActiveViolations(ctx context.Context, params *ListActiveViolationsInput, optFns ...func(*Options)) (*ListActiveViolationsOutput, error) 
 ListAttachedPolicies(ctx context.Context, params *ListAttachedPoliciesInput, optFns ...func(*Options)) (*ListAttachedPoliciesOutput, error) 
 ListAuditFindings(ctx context.Context, params *ListAuditFindingsInput, optFns ...func(*Options)) (*ListAuditFindingsOutput, error) 
 ListAuditMitigationActionsExecutions(ctx context.Context, params *ListAuditMitigationActionsExecutionsInput, optFns ...func(*Options)) (*ListAuditMitigationActionsExecutionsOutput, error) 
 ListAuditMitigationActionsTasks(ctx context.Context, params *ListAuditMitigationActionsTasksInput, optFns ...func(*Options)) (*ListAuditMitigationActionsTasksOutput, error) 
 ListAuditSuppressions(ctx context.Context, params *ListAuditSuppressionsInput, optFns ...func(*Options)) (*ListAuditSuppressionsOutput, error) 
 ListAuditTasks(ctx context.Context, params *ListAuditTasksInput, optFns ...func(*Options)) (*ListAuditTasksOutput, error) 
 ListAuthorizers(ctx context.Context, params *ListAuthorizersInput, optFns ...func(*Options)) (*ListAuthorizersOutput, error) 
 ListBillingGroups(ctx context.Context, params *ListBillingGroupsInput, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) 
 ListCACertificates(ctx context.Context, params *ListCACertificatesInput, optFns ...func(*Options)) (*ListCACertificatesOutput, error) 
 ListCertificateProviders(ctx context.Context, params *ListCertificateProvidersInput, optFns ...func(*Options)) (*ListCertificateProvidersOutput, error) 
 ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) 
 ListCertificatesByCA(ctx context.Context, params *ListCertificatesByCAInput, optFns ...func(*Options)) (*ListCertificatesByCAOutput, error) 
 ListCustomMetrics(ctx context.Context, params *ListCustomMetricsInput, optFns ...func(*Options)) (*ListCustomMetricsOutput, error) 
 ListDetectMitigationActionsExecutions(ctx context.Context, params *ListDetectMitigationActionsExecutionsInput, optFns ...func(*Options)) (*ListDetectMitigationActionsExecutionsOutput, error) 
 ListDetectMitigationActionsTasks(ctx context.Context, params *ListDetectMitigationActionsTasksInput, optFns ...func(*Options)) (*ListDetectMitigationActionsTasksOutput, error) 
 ListDimensions(ctx context.Context, params *ListDimensionsInput, optFns ...func(*Options)) (*ListDimensionsOutput, error) 
 ListDomainConfigurations(ctx context.Context, params *ListDomainConfigurationsInput, optFns ...func(*Options)) (*ListDomainConfigurationsOutput, error) 
 ListFleetMetrics(ctx context.Context, params *ListFleetMetricsInput, optFns ...func(*Options)) (*ListFleetMetricsOutput, error) 
 ListIndices(ctx context.Context, params *ListIndicesInput, optFns ...func(*Options)) (*ListIndicesOutput, error) 
 ListJobExecutionsForJob(ctx context.Context, params *ListJobExecutionsForJobInput, optFns ...func(*Options)) (*ListJobExecutionsForJobOutput, error) 
 ListJobExecutionsForThing(ctx context.Context, params *ListJobExecutionsForThingInput, optFns ...func(*Options)) (*ListJobExecutionsForThingOutput, error) 
 ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListManagedJobTemplates(ctx context.Context, params *ListManagedJobTemplatesInput, optFns ...func(*Options)) (*ListManagedJobTemplatesOutput, error) 
 ListMetricValues(ctx context.Context, params *ListMetricValuesInput, optFns ...func(*Options)) (*ListMetricValuesOutput, error) 
 ListMitigationActions(ctx context.Context, params *ListMitigationActionsInput, optFns ...func(*Options)) (*ListMitigationActionsOutput, error) 
 ListOTAUpdates(ctx context.Context, params *ListOTAUpdatesInput, optFns ...func(*Options)) (*ListOTAUpdatesOutput, error) 
 ListOutgoingCertificates(ctx context.Context, params *ListOutgoingCertificatesInput, optFns ...func(*Options)) (*ListOutgoingCertificatesOutput, error) 
 ListPackageVersions(ctx context.Context, params *ListPackageVersionsInput, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) 
 ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) 
 ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) 
 ListPolicyPrincipals(ctx context.Context, params *ListPolicyPrincipalsInput, optFns ...func(*Options)) (*ListPolicyPrincipalsOutput, error) 
 ListPolicyVersions(ctx context.Context, params *ListPolicyVersionsInput, optFns ...func(*Options)) (*ListPolicyVersionsOutput, error) 
 ListPrincipalPolicies(ctx context.Context, params *ListPrincipalPoliciesInput, optFns ...func(*Options)) (*ListPrincipalPoliciesOutput, error) 
 ListPrincipalThings(ctx context.Context, params *ListPrincipalThingsInput, optFns ...func(*Options)) (*ListPrincipalThingsOutput, error) 
 ListProvisioningTemplateVersions(ctx context.Context, params *ListProvisioningTemplateVersionsInput, optFns ...func(*Options)) (*ListProvisioningTemplateVersionsOutput, error) 
 ListProvisioningTemplates(ctx context.Context, params *ListProvisioningTemplatesInput, optFns ...func(*Options)) (*ListProvisioningTemplatesOutput, error) 
 ListRelatedResourcesForAuditFinding(ctx context.Context, params *ListRelatedResourcesForAuditFindingInput, optFns ...func(*Options)) (*ListRelatedResourcesForAuditFindingOutput, error) 
 ListRoleAliases(ctx context.Context, params *ListRoleAliasesInput, optFns ...func(*Options)) (*ListRoleAliasesOutput, error) 
 ListScheduledAudits(ctx context.Context, params *ListScheduledAuditsInput, optFns ...func(*Options)) (*ListScheduledAuditsOutput, error) 
 ListSecurityProfiles(ctx context.Context, params *ListSecurityProfilesInput, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) 
 ListSecurityProfilesForTarget(ctx context.Context, params *ListSecurityProfilesForTargetInput, optFns ...func(*Options)) (*ListSecurityProfilesForTargetOutput, error) 
 ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTargetsForPolicy(ctx context.Context, params *ListTargetsForPolicyInput, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) 
 ListTargetsForSecurityProfile(ctx context.Context, params *ListTargetsForSecurityProfileInput, optFns ...func(*Options)) (*ListTargetsForSecurityProfileOutput, error) 
 ListThingGroups(ctx context.Context, params *ListThingGroupsInput, optFns ...func(*Options)) (*ListThingGroupsOutput, error) 
 ListThingGroupsForThing(ctx context.Context, params *ListThingGroupsForThingInput, optFns ...func(*Options)) (*ListThingGroupsForThingOutput, error) 
 ListThingPrincipals(ctx context.Context, params *ListThingPrincipalsInput, optFns ...func(*Options)) (*ListThingPrincipalsOutput, error) 
 ListThingRegistrationTaskReports(ctx context.Context, params *ListThingRegistrationTaskReportsInput, optFns ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error) 
 ListThingRegistrationTasks(ctx context.Context, params *ListThingRegistrationTasksInput, optFns ...func(*Options)) (*ListThingRegistrationTasksOutput, error) 
 ListThingTypes(ctx context.Context, params *ListThingTypesInput, optFns ...func(*Options)) (*ListThingTypesOutput, error) 
 ListThings(ctx context.Context, params *ListThingsInput, optFns ...func(*Options)) (*ListThingsOutput, error) 
 ListThingsInBillingGroup(ctx context.Context, params *ListThingsInBillingGroupInput, optFns ...func(*Options)) (*ListThingsInBillingGroupOutput, error) 
 ListThingsInThingGroup(ctx context.Context, params *ListThingsInThingGroupInput, optFns ...func(*Options)) (*ListThingsInThingGroupOutput, error) 
 ListTopicRuleDestinations(ctx context.Context, params *ListTopicRuleDestinationsInput, optFns ...func(*Options)) (*ListTopicRuleDestinationsOutput, error) 
 ListTopicRules(ctx context.Context, params *ListTopicRulesInput, optFns ...func(*Options)) (*ListTopicRulesOutput, error) 
 ListV2LoggingLevels(ctx context.Context, params *ListV2LoggingLevelsInput, optFns ...func(*Options)) (*ListV2LoggingLevelsOutput, error) 
 ListViolationEvents(ctx context.Context, params *ListViolationEventsInput, optFns ...func(*Options)) (*ListViolationEventsOutput, error) 
 PutVerificationStateOnViolation(ctx context.Context, params *PutVerificationStateOnViolationInput, optFns ...func(*Options)) (*PutVerificationStateOnViolationOutput, error) 
 RegisterCACertificate(ctx context.Context, params *RegisterCACertificateInput, optFns ...func(*Options)) (*RegisterCACertificateOutput, error) 
 RegisterCertificate(ctx context.Context, params *RegisterCertificateInput, optFns ...func(*Options)) (*RegisterCertificateOutput, error) 
 RegisterCertificateWithoutCA(ctx context.Context, params *RegisterCertificateWithoutCAInput, optFns ...func(*Options)) (*RegisterCertificateWithoutCAOutput, error) 
 RegisterThing(ctx context.Context, params *RegisterThingInput, optFns ...func(*Options)) (*RegisterThingOutput, error) 
 RejectCertificateTransfer(ctx context.Context, params *RejectCertificateTransferInput, optFns ...func(*Options)) (*RejectCertificateTransferOutput, error) 
 RemoveThingFromBillingGroup(ctx context.Context, params *RemoveThingFromBillingGroupInput, optFns ...func(*Options)) (*RemoveThingFromBillingGroupOutput, error) 
 RemoveThingFromThingGroup(ctx context.Context, params *RemoveThingFromThingGroupInput, optFns ...func(*Options)) (*RemoveThingFromThingGroupOutput, error) 
 ReplaceTopicRule(ctx context.Context, params *ReplaceTopicRuleInput, optFns ...func(*Options)) (*ReplaceTopicRuleOutput, error) 
 SearchIndex(ctx context.Context, params *SearchIndexInput, optFns ...func(*Options)) (*SearchIndexOutput, error) 
 SetDefaultAuthorizer(ctx context.Context, params *SetDefaultAuthorizerInput, optFns ...func(*Options)) (*SetDefaultAuthorizerOutput, error) 
 SetDefaultPolicyVersion(ctx context.Context, params *SetDefaultPolicyVersionInput, optFns ...func(*Options)) (*SetDefaultPolicyVersionOutput, error) 
 SetLoggingOptions(ctx context.Context, params *SetLoggingOptionsInput, optFns ...func(*Options)) (*SetLoggingOptionsOutput, error) 
 SetV2LoggingLevel(ctx context.Context, params *SetV2LoggingLevelInput, optFns ...func(*Options)) (*SetV2LoggingLevelOutput, error) 
 SetV2LoggingOptions(ctx context.Context, params *SetV2LoggingOptionsInput, optFns ...func(*Options)) (*SetV2LoggingOptionsOutput, error) 
 StartAuditMitigationActionsTask(ctx context.Context, params *StartAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*StartAuditMitigationActionsTaskOutput, error) 
 StartDetectMitigationActionsTask(ctx context.Context, params *StartDetectMitigationActionsTaskInput, optFns ...func(*Options)) (*StartDetectMitigationActionsTaskOutput, error) 
 StartOnDemandAuditTask(ctx context.Context, params *StartOnDemandAuditTaskInput, optFns ...func(*Options)) (*StartOnDemandAuditTaskOutput, error) 
 StartThingRegistrationTask(ctx context.Context, params *StartThingRegistrationTaskInput, optFns ...func(*Options)) (*StartThingRegistrationTaskOutput, error) 
 StopThingRegistrationTask(ctx context.Context, params *StopThingRegistrationTaskInput, optFns ...func(*Options)) (*StopThingRegistrationTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestAuthorization(ctx context.Context, params *TestAuthorizationInput, optFns ...func(*Options)) (*TestAuthorizationOutput, error) 
 TestInvokeAuthorizer(ctx context.Context, params *TestInvokeAuthorizerInput, optFns ...func(*Options)) (*TestInvokeAuthorizerOutput, error) 
 TransferCertificate(ctx context.Context, params *TransferCertificateInput, optFns ...func(*Options)) (*TransferCertificateOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccountAuditConfiguration(ctx context.Context, params *UpdateAccountAuditConfigurationInput, optFns ...func(*Options)) (*UpdateAccountAuditConfigurationOutput, error) 
 UpdateAuditSuppression(ctx context.Context, params *UpdateAuditSuppressionInput, optFns ...func(*Options)) (*UpdateAuditSuppressionOutput, error) 
 UpdateAuthorizer(ctx context.Context, params *UpdateAuthorizerInput, optFns ...func(*Options)) (*UpdateAuthorizerOutput, error) 
 UpdateBillingGroup(ctx context.Context, params *UpdateBillingGroupInput, optFns ...func(*Options)) (*UpdateBillingGroupOutput, error) 
 UpdateCACertificate(ctx context.Context, params *UpdateCACertificateInput, optFns ...func(*Options)) (*UpdateCACertificateOutput, error) 
 UpdateCertificate(ctx context.Context, params *UpdateCertificateInput, optFns ...func(*Options)) (*UpdateCertificateOutput, error) 
 UpdateCertificateProvider(ctx context.Context, params *UpdateCertificateProviderInput, optFns ...func(*Options)) (*UpdateCertificateProviderOutput, error) 
 UpdateCustomMetric(ctx context.Context, params *UpdateCustomMetricInput, optFns ...func(*Options)) (*UpdateCustomMetricOutput, error) 
 UpdateDimension(ctx context.Context, params *UpdateDimensionInput, optFns ...func(*Options)) (*UpdateDimensionOutput, error) 
 UpdateDomainConfiguration(ctx context.Context, params *UpdateDomainConfigurationInput, optFns ...func(*Options)) (*UpdateDomainConfigurationOutput, error) 
 UpdateDynamicThingGroup(ctx context.Context, params *UpdateDynamicThingGroupInput, optFns ...func(*Options)) (*UpdateDynamicThingGroupOutput, error) 
 UpdateEventConfigurations(ctx context.Context, params *UpdateEventConfigurationsInput, optFns ...func(*Options)) (*UpdateEventConfigurationsOutput, error) 
 UpdateFleetMetric(ctx context.Context, params *UpdateFleetMetricInput, optFns ...func(*Options)) (*UpdateFleetMetricOutput, error) 
 UpdateIndexingConfiguration(ctx context.Context, params *UpdateIndexingConfigurationInput, optFns ...func(*Options)) (*UpdateIndexingConfigurationOutput, error) 
 UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) 
 UpdateMitigationAction(ctx context.Context, params *UpdateMitigationActionInput, optFns ...func(*Options)) (*UpdateMitigationActionOutput, error) 
 UpdatePackage(ctx context.Context, params *UpdatePackageInput, optFns ...func(*Options)) (*UpdatePackageOutput, error) 
 UpdatePackageConfiguration(ctx context.Context, params *UpdatePackageConfigurationInput, optFns ...func(*Options)) (*UpdatePackageConfigurationOutput, error) 
 UpdatePackageVersion(ctx context.Context, params *UpdatePackageVersionInput, optFns ...func(*Options)) (*UpdatePackageVersionOutput, error) 
 UpdateProvisioningTemplate(ctx context.Context, params *UpdateProvisioningTemplateInput, optFns ...func(*Options)) (*UpdateProvisioningTemplateOutput, error) 
 UpdateRoleAlias(ctx context.Context, params *UpdateRoleAliasInput, optFns ...func(*Options)) (*UpdateRoleAliasOutput, error) 
 UpdateScheduledAudit(ctx context.Context, params *UpdateScheduledAuditInput, optFns ...func(*Options)) (*UpdateScheduledAuditOutput, error) 
 UpdateSecurityProfile(ctx context.Context, params *UpdateSecurityProfileInput, optFns ...func(*Options)) (*UpdateSecurityProfileOutput, error) 
 UpdateStream(ctx context.Context, params *UpdateStreamInput, optFns ...func(*Options)) (*UpdateStreamOutput, error) 
 UpdateThing(ctx context.Context, params *UpdateThingInput, optFns ...func(*Options)) (*UpdateThingOutput, error) 
 UpdateThingGroup(ctx context.Context, params *UpdateThingGroupInput, optFns ...func(*Options)) (*UpdateThingGroupOutput, error) 
 UpdateThingGroupsForThing(ctx context.Context, params *UpdateThingGroupsForThingInput, optFns ...func(*Options)) (*UpdateThingGroupsForThingOutput, error) 
 UpdateTopicRuleDestination(ctx context.Context, params *UpdateTopicRuleDestinationInput, optFns ...func(*Options)) (*UpdateTopicRuleDestinationOutput, error) 
 ValidateSecurityProfileBehaviors(ctx context.Context, params *ValidateSecurityProfileBehaviorsInput, optFns ...func(*Options)) (*ValidateSecurityProfileBehaviorsOutput, error) 
}
