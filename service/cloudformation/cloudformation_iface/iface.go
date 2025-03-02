
package cloudformation_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// IClient defines the interface for cloudformation
type IClient interface {
 Options() Options 
 ActivateOrganizationsAccess(ctx context.Context, params *ActivateOrganizationsAccessInput, optFns ...func(*Options)) (*ActivateOrganizationsAccessOutput, error) 
 ActivateType(ctx context.Context, params *ActivateTypeInput, optFns ...func(*Options)) (*ActivateTypeOutput, error) 
 BatchDescribeTypeConfigurations(ctx context.Context, params *BatchDescribeTypeConfigurationsInput, optFns ...func(*Options)) (*BatchDescribeTypeConfigurationsOutput, error) 
 CancelUpdateStack(ctx context.Context, params *CancelUpdateStackInput, optFns ...func(*Options)) (*CancelUpdateStackOutput, error) 
 ContinueUpdateRollback(ctx context.Context, params *ContinueUpdateRollbackInput, optFns ...func(*Options)) (*ContinueUpdateRollbackOutput, error) 
 CreateChangeSet(ctx context.Context, params *CreateChangeSetInput, optFns ...func(*Options)) (*CreateChangeSetOutput, error) 
 CreateGeneratedTemplate(ctx context.Context, params *CreateGeneratedTemplateInput, optFns ...func(*Options)) (*CreateGeneratedTemplateOutput, error) 
 CreateStack(ctx context.Context, params *CreateStackInput, optFns ...func(*Options)) (*CreateStackOutput, error) 
 CreateStackInstances(ctx context.Context, params *CreateStackInstancesInput, optFns ...func(*Options)) (*CreateStackInstancesOutput, error) 
 CreateStackRefactor(ctx context.Context, params *CreateStackRefactorInput, optFns ...func(*Options)) (*CreateStackRefactorOutput, error) 
 CreateStackSet(ctx context.Context, params *CreateStackSetInput, optFns ...func(*Options)) (*CreateStackSetOutput, error) 
 DeactivateOrganizationsAccess(ctx context.Context, params *DeactivateOrganizationsAccessInput, optFns ...func(*Options)) (*DeactivateOrganizationsAccessOutput, error) 
 DeactivateType(ctx context.Context, params *DeactivateTypeInput, optFns ...func(*Options)) (*DeactivateTypeOutput, error) 
 DeleteChangeSet(ctx context.Context, params *DeleteChangeSetInput, optFns ...func(*Options)) (*DeleteChangeSetOutput, error) 
 DeleteGeneratedTemplate(ctx context.Context, params *DeleteGeneratedTemplateInput, optFns ...func(*Options)) (*DeleteGeneratedTemplateOutput, error) 
 DeleteStack(ctx context.Context, params *DeleteStackInput, optFns ...func(*Options)) (*DeleteStackOutput, error) 
 DeleteStackInstances(ctx context.Context, params *DeleteStackInstancesInput, optFns ...func(*Options)) (*DeleteStackInstancesOutput, error) 
 DeleteStackSet(ctx context.Context, params *DeleteStackSetInput, optFns ...func(*Options)) (*DeleteStackSetOutput, error) 
 DeregisterType(ctx context.Context, params *DeregisterTypeInput, optFns ...func(*Options)) (*DeregisterTypeOutput, error) 
 DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) 
 DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) 
 DescribeChangeSetHooks(ctx context.Context, params *DescribeChangeSetHooksInput, optFns ...func(*Options)) (*DescribeChangeSetHooksOutput, error) 
 DescribeGeneratedTemplate(ctx context.Context, params *DescribeGeneratedTemplateInput, optFns ...func(*Options)) (*DescribeGeneratedTemplateOutput, error) 
 DescribeOrganizationsAccess(ctx context.Context, params *DescribeOrganizationsAccessInput, optFns ...func(*Options)) (*DescribeOrganizationsAccessOutput, error) 
 DescribePublisher(ctx context.Context, params *DescribePublisherInput, optFns ...func(*Options)) (*DescribePublisherOutput, error) 
 DescribeResourceScan(ctx context.Context, params *DescribeResourceScanInput, optFns ...func(*Options)) (*DescribeResourceScanOutput, error) 
 DescribeStackDriftDetectionStatus(ctx context.Context, params *DescribeStackDriftDetectionStatusInput, optFns ...func(*Options)) (*DescribeStackDriftDetectionStatusOutput, error) 
 DescribeStackEvents(ctx context.Context, params *DescribeStackEventsInput, optFns ...func(*Options)) (*DescribeStackEventsOutput, error) 
 DescribeStackInstance(ctx context.Context, params *DescribeStackInstanceInput, optFns ...func(*Options)) (*DescribeStackInstanceOutput, error) 
 DescribeStackRefactor(ctx context.Context, params *DescribeStackRefactorInput, optFns ...func(*Options)) (*DescribeStackRefactorOutput, error) 
 DescribeStackResource(ctx context.Context, params *DescribeStackResourceInput, optFns ...func(*Options)) (*DescribeStackResourceOutput, error) 
 DescribeStackResourceDrifts(ctx context.Context, params *DescribeStackResourceDriftsInput, optFns ...func(*Options)) (*DescribeStackResourceDriftsOutput, error) 
 DescribeStackResources(ctx context.Context, params *DescribeStackResourcesInput, optFns ...func(*Options)) (*DescribeStackResourcesOutput, error) 
 DescribeStackSet(ctx context.Context, params *DescribeStackSetInput, optFns ...func(*Options)) (*DescribeStackSetOutput, error) 
 DescribeStackSetOperation(ctx context.Context, params *DescribeStackSetOperationInput, optFns ...func(*Options)) (*DescribeStackSetOperationOutput, error) 
 DescribeStacks(ctx context.Context, params *DescribeStacksInput, optFns ...func(*Options)) (*DescribeStacksOutput, error) 
 DescribeType(ctx context.Context, params *DescribeTypeInput, optFns ...func(*Options)) (*DescribeTypeOutput, error) 
 DescribeTypeRegistration(ctx context.Context, params *DescribeTypeRegistrationInput, optFns ...func(*Options)) (*DescribeTypeRegistrationOutput, error) 
 DetectStackDrift(ctx context.Context, params *DetectStackDriftInput, optFns ...func(*Options)) (*DetectStackDriftOutput, error) 
 DetectStackResourceDrift(ctx context.Context, params *DetectStackResourceDriftInput, optFns ...func(*Options)) (*DetectStackResourceDriftOutput, error) 
 DetectStackSetDrift(ctx context.Context, params *DetectStackSetDriftInput, optFns ...func(*Options)) (*DetectStackSetDriftOutput, error) 
 EstimateTemplateCost(ctx context.Context, params *EstimateTemplateCostInput, optFns ...func(*Options)) (*EstimateTemplateCostOutput, error) 
 ExecuteChangeSet(ctx context.Context, params *ExecuteChangeSetInput, optFns ...func(*Options)) (*ExecuteChangeSetOutput, error) 
 ExecuteStackRefactor(ctx context.Context, params *ExecuteStackRefactorInput, optFns ...func(*Options)) (*ExecuteStackRefactorOutput, error) 
 GetGeneratedTemplate(ctx context.Context, params *GetGeneratedTemplateInput, optFns ...func(*Options)) (*GetGeneratedTemplateOutput, error) 
 GetStackPolicy(ctx context.Context, params *GetStackPolicyInput, optFns ...func(*Options)) (*GetStackPolicyOutput, error) 
 GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) 
 GetTemplateSummary(ctx context.Context, params *GetTemplateSummaryInput, optFns ...func(*Options)) (*GetTemplateSummaryOutput, error) 
 ImportStacksToStackSet(ctx context.Context, params *ImportStacksToStackSetInput, optFns ...func(*Options)) (*ImportStacksToStackSetOutput, error) 
 ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) 
 ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) 
 ListGeneratedTemplates(ctx context.Context, params *ListGeneratedTemplatesInput, optFns ...func(*Options)) (*ListGeneratedTemplatesOutput, error) 
 ListHookResults(ctx context.Context, params *ListHookResultsInput, optFns ...func(*Options)) (*ListHookResultsOutput, error) 
 ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error) 
 ListResourceScanRelatedResources(ctx context.Context, params *ListResourceScanRelatedResourcesInput, optFns ...func(*Options)) (*ListResourceScanRelatedResourcesOutput, error) 
 ListResourceScanResources(ctx context.Context, params *ListResourceScanResourcesInput, optFns ...func(*Options)) (*ListResourceScanResourcesOutput, error) 
 ListResourceScans(ctx context.Context, params *ListResourceScansInput, optFns ...func(*Options)) (*ListResourceScansOutput, error) 
 ListStackInstanceResourceDrifts(ctx context.Context, params *ListStackInstanceResourceDriftsInput, optFns ...func(*Options)) (*ListStackInstanceResourceDriftsOutput, error) 
 ListStackInstances(ctx context.Context, params *ListStackInstancesInput, optFns ...func(*Options)) (*ListStackInstancesOutput, error) 
 ListStackRefactorActions(ctx context.Context, params *ListStackRefactorActionsInput, optFns ...func(*Options)) (*ListStackRefactorActionsOutput, error) 
 ListStackRefactors(ctx context.Context, params *ListStackRefactorsInput, optFns ...func(*Options)) (*ListStackRefactorsOutput, error) 
 ListStackResources(ctx context.Context, params *ListStackResourcesInput, optFns ...func(*Options)) (*ListStackResourcesOutput, error) 
 ListStackSetAutoDeploymentTargets(ctx context.Context, params *ListStackSetAutoDeploymentTargetsInput, optFns ...func(*Options)) (*ListStackSetAutoDeploymentTargetsOutput, error) 
 ListStackSetOperationResults(ctx context.Context, params *ListStackSetOperationResultsInput, optFns ...func(*Options)) (*ListStackSetOperationResultsOutput, error) 
 ListStackSetOperations(ctx context.Context, params *ListStackSetOperationsInput, optFns ...func(*Options)) (*ListStackSetOperationsOutput, error) 
 ListStackSets(ctx context.Context, params *ListStackSetsInput, optFns ...func(*Options)) (*ListStackSetsOutput, error) 
 ListStacks(ctx context.Context, params *ListStacksInput, optFns ...func(*Options)) (*ListStacksOutput, error) 
 ListTypeRegistrations(ctx context.Context, params *ListTypeRegistrationsInput, optFns ...func(*Options)) (*ListTypeRegistrationsOutput, error) 
 ListTypeVersions(ctx context.Context, params *ListTypeVersionsInput, optFns ...func(*Options)) (*ListTypeVersionsOutput, error) 
 ListTypes(ctx context.Context, params *ListTypesInput, optFns ...func(*Options)) (*ListTypesOutput, error) 
 PublishType(ctx context.Context, params *PublishTypeInput, optFns ...func(*Options)) (*PublishTypeOutput, error) 
 RecordHandlerProgress(ctx context.Context, params *RecordHandlerProgressInput, optFns ...func(*Options)) (*RecordHandlerProgressOutput, error) 
 RegisterPublisher(ctx context.Context, params *RegisterPublisherInput, optFns ...func(*Options)) (*RegisterPublisherOutput, error) 
 RegisterType(ctx context.Context, params *RegisterTypeInput, optFns ...func(*Options)) (*RegisterTypeOutput, error) 
 RollbackStack(ctx context.Context, params *RollbackStackInput, optFns ...func(*Options)) (*RollbackStackOutput, error) 
 SetStackPolicy(ctx context.Context, params *SetStackPolicyInput, optFns ...func(*Options)) (*SetStackPolicyOutput, error) 
 SetTypeConfiguration(ctx context.Context, params *SetTypeConfigurationInput, optFns ...func(*Options)) (*SetTypeConfigurationOutput, error) 
 SetTypeDefaultVersion(ctx context.Context, params *SetTypeDefaultVersionInput, optFns ...func(*Options)) (*SetTypeDefaultVersionOutput, error) 
 SignalResource(ctx context.Context, params *SignalResourceInput, optFns ...func(*Options)) (*SignalResourceOutput, error) 
 StartResourceScan(ctx context.Context, params *StartResourceScanInput, optFns ...func(*Options)) (*StartResourceScanOutput, error) 
 StopStackSetOperation(ctx context.Context, params *StopStackSetOperationInput, optFns ...func(*Options)) (*StopStackSetOperationOutput, error) 
 TestType(ctx context.Context, params *TestTypeInput, optFns ...func(*Options)) (*TestTypeOutput, error) 
 UpdateGeneratedTemplate(ctx context.Context, params *UpdateGeneratedTemplateInput, optFns ...func(*Options)) (*UpdateGeneratedTemplateOutput, error) 
 UpdateStack(ctx context.Context, params *UpdateStackInput, optFns ...func(*Options)) (*UpdateStackOutput, error) 
 UpdateStackInstances(ctx context.Context, params *UpdateStackInstancesInput, optFns ...func(*Options)) (*UpdateStackInstancesOutput, error) 
 UpdateStackSet(ctx context.Context, params *UpdateStackSetInput, optFns ...func(*Options)) (*UpdateStackSetOutput, error) 
 UpdateTerminationProtection(ctx context.Context, params *UpdateTerminationProtectionInput, optFns ...func(*Options)) (*UpdateTerminationProtectionOutput, error) 
 ValidateTemplate(ctx context.Context, params *ValidateTemplateInput, optFns ...func(*Options)) (*ValidateTemplateOutput, error) 
}
