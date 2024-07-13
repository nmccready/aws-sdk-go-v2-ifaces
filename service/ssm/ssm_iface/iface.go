
package ssm_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/ssm"
)

// IClient defines the interface for ssm
type IClient interface {
 Options() Options 
 AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) 
 AssociateOpsItemRelatedItem(ctx context.Context, params *AssociateOpsItemRelatedItemInput, optFns ...func(*Options)) (*AssociateOpsItemRelatedItemOutput, error) 
 CancelCommand(ctx context.Context, params *CancelCommandInput, optFns ...func(*Options)) (*CancelCommandOutput, error) 
 CancelMaintenanceWindowExecution(ctx context.Context, params *CancelMaintenanceWindowExecutionInput, optFns ...func(*Options)) (*CancelMaintenanceWindowExecutionOutput, error) 
 CreateActivation(ctx context.Context, params *CreateActivationInput, optFns ...func(*Options)) (*CreateActivationOutput, error) 
 CreateAssociation(ctx context.Context, params *CreateAssociationInput, optFns ...func(*Options)) (*CreateAssociationOutput, error) 
 CreateAssociationBatch(ctx context.Context, params *CreateAssociationBatchInput, optFns ...func(*Options)) (*CreateAssociationBatchOutput, error) 
 CreateDocument(ctx context.Context, params *CreateDocumentInput, optFns ...func(*Options)) (*CreateDocumentOutput, error) 
 CreateMaintenanceWindow(ctx context.Context, params *CreateMaintenanceWindowInput, optFns ...func(*Options)) (*CreateMaintenanceWindowOutput, error) 
 CreateOpsItem(ctx context.Context, params *CreateOpsItemInput, optFns ...func(*Options)) (*CreateOpsItemOutput, error) 
 CreateOpsMetadata(ctx context.Context, params *CreateOpsMetadataInput, optFns ...func(*Options)) (*CreateOpsMetadataOutput, error) 
 CreatePatchBaseline(ctx context.Context, params *CreatePatchBaselineInput, optFns ...func(*Options)) (*CreatePatchBaselineOutput, error) 
 CreateResourceDataSync(ctx context.Context, params *CreateResourceDataSyncInput, optFns ...func(*Options)) (*CreateResourceDataSyncOutput, error) 
 DeleteActivation(ctx context.Context, params *DeleteActivationInput, optFns ...func(*Options)) (*DeleteActivationOutput, error) 
 DeleteAssociation(ctx context.Context, params *DeleteAssociationInput, optFns ...func(*Options)) (*DeleteAssociationOutput, error) 
 DeleteDocument(ctx context.Context, params *DeleteDocumentInput, optFns ...func(*Options)) (*DeleteDocumentOutput, error) 
 DeleteInventory(ctx context.Context, params *DeleteInventoryInput, optFns ...func(*Options)) (*DeleteInventoryOutput, error) 
 DeleteMaintenanceWindow(ctx context.Context, params *DeleteMaintenanceWindowInput, optFns ...func(*Options)) (*DeleteMaintenanceWindowOutput, error) 
 DeleteOpsItem(ctx context.Context, params *DeleteOpsItemInput, optFns ...func(*Options)) (*DeleteOpsItemOutput, error) 
 DeleteOpsMetadata(ctx context.Context, params *DeleteOpsMetadataInput, optFns ...func(*Options)) (*DeleteOpsMetadataOutput, error) 
 DeleteParameter(ctx context.Context, params *DeleteParameterInput, optFns ...func(*Options)) (*DeleteParameterOutput, error) 
 DeleteParameters(ctx context.Context, params *DeleteParametersInput, optFns ...func(*Options)) (*DeleteParametersOutput, error) 
 DeletePatchBaseline(ctx context.Context, params *DeletePatchBaselineInput, optFns ...func(*Options)) (*DeletePatchBaselineOutput, error) 
 DeleteResourceDataSync(ctx context.Context, params *DeleteResourceDataSyncInput, optFns ...func(*Options)) (*DeleteResourceDataSyncOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeregisterManagedInstance(ctx context.Context, params *DeregisterManagedInstanceInput, optFns ...func(*Options)) (*DeregisterManagedInstanceOutput, error) 
 DeregisterPatchBaselineForPatchGroup(ctx context.Context, params *DeregisterPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*DeregisterPatchBaselineForPatchGroupOutput, error) 
 DeregisterTargetFromMaintenanceWindow(ctx context.Context, params *DeregisterTargetFromMaintenanceWindowInput, optFns ...func(*Options)) (*DeregisterTargetFromMaintenanceWindowOutput, error) 
 DeregisterTaskFromMaintenanceWindow(ctx context.Context, params *DeregisterTaskFromMaintenanceWindowInput, optFns ...func(*Options)) (*DeregisterTaskFromMaintenanceWindowOutput, error) 
 DescribeActivations(ctx context.Context, params *DescribeActivationsInput, optFns ...func(*Options)) (*DescribeActivationsOutput, error) 
 DescribeAssociation(ctx context.Context, params *DescribeAssociationInput, optFns ...func(*Options)) (*DescribeAssociationOutput, error) 
 DescribeAssociationExecutionTargets(ctx context.Context, params *DescribeAssociationExecutionTargetsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error) 
 DescribeAssociationExecutions(ctx context.Context, params *DescribeAssociationExecutionsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionsOutput, error) 
 DescribeAutomationExecutions(ctx context.Context, params *DescribeAutomationExecutionsInput, optFns ...func(*Options)) (*DescribeAutomationExecutionsOutput, error) 
 DescribeAutomationStepExecutions(ctx context.Context, params *DescribeAutomationStepExecutionsInput, optFns ...func(*Options)) (*DescribeAutomationStepExecutionsOutput, error) 
 DescribeAvailablePatches(ctx context.Context, params *DescribeAvailablePatchesInput, optFns ...func(*Options)) (*DescribeAvailablePatchesOutput, error) 
 DescribeDocument(ctx context.Context, params *DescribeDocumentInput, optFns ...func(*Options)) (*DescribeDocumentOutput, error) 
 DescribeDocumentPermission(ctx context.Context, params *DescribeDocumentPermissionInput, optFns ...func(*Options)) (*DescribeDocumentPermissionOutput, error) 
 DescribeEffectiveInstanceAssociations(ctx context.Context, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) 
 DescribeEffectivePatchesForPatchBaseline(ctx context.Context, params *DescribeEffectivePatchesForPatchBaselineInput, optFns ...func(*Options)) (*DescribeEffectivePatchesForPatchBaselineOutput, error) 
 DescribeInstanceAssociationsStatus(ctx context.Context, params *DescribeInstanceAssociationsStatusInput, optFns ...func(*Options)) (*DescribeInstanceAssociationsStatusOutput, error) 
 DescribeInstanceInformation(ctx context.Context, params *DescribeInstanceInformationInput, optFns ...func(*Options)) (*DescribeInstanceInformationOutput, error) 
 DescribeInstancePatchStates(ctx context.Context, params *DescribeInstancePatchStatesInput, optFns ...func(*Options)) (*DescribeInstancePatchStatesOutput, error) 
 DescribeInstancePatchStatesForPatchGroup(ctx context.Context, params *DescribeInstancePatchStatesForPatchGroupInput, optFns ...func(*Options)) (*DescribeInstancePatchStatesForPatchGroupOutput, error) 
 DescribeInstancePatches(ctx context.Context, params *DescribeInstancePatchesInput, optFns ...func(*Options)) (*DescribeInstancePatchesOutput, error) 
 DescribeInstanceProperties(ctx context.Context, params *DescribeInstancePropertiesInput, optFns ...func(*Options)) (*DescribeInstancePropertiesOutput, error) 
 DescribeInventoryDeletions(ctx context.Context, params *DescribeInventoryDeletionsInput, optFns ...func(*Options)) (*DescribeInventoryDeletionsOutput, error) 
 DescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, params *DescribeMaintenanceWindowExecutionTaskInvocationsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) 
 DescribeMaintenanceWindowExecutionTasks(ctx context.Context, params *DescribeMaintenanceWindowExecutionTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTasksOutput, error) 
 DescribeMaintenanceWindowExecutions(ctx context.Context, params *DescribeMaintenanceWindowExecutionsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionsOutput, error) 
 DescribeMaintenanceWindowSchedule(ctx context.Context, params *DescribeMaintenanceWindowScheduleInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowScheduleOutput, error) 
 DescribeMaintenanceWindowTargets(ctx context.Context, params *DescribeMaintenanceWindowTargetsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTargetsOutput, error) 
 DescribeMaintenanceWindowTasks(ctx context.Context, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) 
 DescribeMaintenanceWindows(ctx context.Context, params *DescribeMaintenanceWindowsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowsOutput, error) 
 DescribeMaintenanceWindowsForTarget(ctx context.Context, params *DescribeMaintenanceWindowsForTargetInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowsForTargetOutput, error) 
 DescribeOpsItems(ctx context.Context, params *DescribeOpsItemsInput, optFns ...func(*Options)) (*DescribeOpsItemsOutput, error) 
 DescribeParameters(ctx context.Context, params *DescribeParametersInput, optFns ...func(*Options)) (*DescribeParametersOutput, error) 
 DescribePatchBaselines(ctx context.Context, params *DescribePatchBaselinesInput, optFns ...func(*Options)) (*DescribePatchBaselinesOutput, error) 
 DescribePatchGroupState(ctx context.Context, params *DescribePatchGroupStateInput, optFns ...func(*Options)) (*DescribePatchGroupStateOutput, error) 
 DescribePatchGroups(ctx context.Context, params *DescribePatchGroupsInput, optFns ...func(*Options)) (*DescribePatchGroupsOutput, error) 
 DescribePatchProperties(ctx context.Context, params *DescribePatchPropertiesInput, optFns ...func(*Options)) (*DescribePatchPropertiesOutput, error) 
 DescribeSessions(ctx context.Context, params *DescribeSessionsInput, optFns ...func(*Options)) (*DescribeSessionsOutput, error) 
 DisassociateOpsItemRelatedItem(ctx context.Context, params *DisassociateOpsItemRelatedItemInput, optFns ...func(*Options)) (*DisassociateOpsItemRelatedItemOutput, error) 
 GetAutomationExecution(ctx context.Context, params *GetAutomationExecutionInput, optFns ...func(*Options)) (*GetAutomationExecutionOutput, error) 
 GetCalendarState(ctx context.Context, params *GetCalendarStateInput, optFns ...func(*Options)) (*GetCalendarStateOutput, error) 
 GetCommandInvocation(ctx context.Context, params *GetCommandInvocationInput, optFns ...func(*Options)) (*GetCommandInvocationOutput, error) 
 GetConnectionStatus(ctx context.Context, params *GetConnectionStatusInput, optFns ...func(*Options)) (*GetConnectionStatusOutput, error) 
 GetDefaultPatchBaseline(ctx context.Context, params *GetDefaultPatchBaselineInput, optFns ...func(*Options)) (*GetDefaultPatchBaselineOutput, error) 
 GetDeployablePatchSnapshotForInstance(ctx context.Context, params *GetDeployablePatchSnapshotForInstanceInput, optFns ...func(*Options)) (*GetDeployablePatchSnapshotForInstanceOutput, error) 
 GetDocument(ctx context.Context, params *GetDocumentInput, optFns ...func(*Options)) (*GetDocumentOutput, error) 
 GetInventory(ctx context.Context, params *GetInventoryInput, optFns ...func(*Options)) (*GetInventoryOutput, error) 
 GetInventorySchema(ctx context.Context, params *GetInventorySchemaInput, optFns ...func(*Options)) (*GetInventorySchemaOutput, error) 
 GetMaintenanceWindow(ctx context.Context, params *GetMaintenanceWindowInput, optFns ...func(*Options)) (*GetMaintenanceWindowOutput, error) 
 GetMaintenanceWindowExecution(ctx context.Context, params *GetMaintenanceWindowExecutionInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionOutput, error) 
 GetMaintenanceWindowExecutionTask(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskOutput, error) 
 GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInvocationInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskInvocationOutput, error) 
 GetMaintenanceWindowTask(ctx context.Context, params *GetMaintenanceWindowTaskInput, optFns ...func(*Options)) (*GetMaintenanceWindowTaskOutput, error) 
 GetOpsItem(ctx context.Context, params *GetOpsItemInput, optFns ...func(*Options)) (*GetOpsItemOutput, error) 
 GetOpsMetadata(ctx context.Context, params *GetOpsMetadataInput, optFns ...func(*Options)) (*GetOpsMetadataOutput, error) 
 GetOpsSummary(ctx context.Context, params *GetOpsSummaryInput, optFns ...func(*Options)) (*GetOpsSummaryOutput, error) 
 GetParameter(ctx context.Context, params *GetParameterInput, optFns ...func(*Options)) (*GetParameterOutput, error) 
 GetParameterHistory(ctx context.Context, params *GetParameterHistoryInput, optFns ...func(*Options)) (*GetParameterHistoryOutput, error) 
 GetParameters(ctx context.Context, params *GetParametersInput, optFns ...func(*Options)) (*GetParametersOutput, error) 
 GetParametersByPath(ctx context.Context, params *GetParametersByPathInput, optFns ...func(*Options)) (*GetParametersByPathOutput, error) 
 GetPatchBaseline(ctx context.Context, params *GetPatchBaselineInput, optFns ...func(*Options)) (*GetPatchBaselineOutput, error) 
 GetPatchBaselineForPatchGroup(ctx context.Context, params *GetPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*GetPatchBaselineForPatchGroupOutput, error) 
 GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) 
 GetServiceSetting(ctx context.Context, params *GetServiceSettingInput, optFns ...func(*Options)) (*GetServiceSettingOutput, error) 
 LabelParameterVersion(ctx context.Context, params *LabelParameterVersionInput, optFns ...func(*Options)) (*LabelParameterVersionOutput, error) 
 ListAssociationVersions(ctx context.Context, params *ListAssociationVersionsInput, optFns ...func(*Options)) (*ListAssociationVersionsOutput, error) 
 ListAssociations(ctx context.Context, params *ListAssociationsInput, optFns ...func(*Options)) (*ListAssociationsOutput, error) 
 ListCommandInvocations(ctx context.Context, params *ListCommandInvocationsInput, optFns ...func(*Options)) (*ListCommandInvocationsOutput, error) 
 ListCommands(ctx context.Context, params *ListCommandsInput, optFns ...func(*Options)) (*ListCommandsOutput, error) 
 ListComplianceItems(ctx context.Context, params *ListComplianceItemsInput, optFns ...func(*Options)) (*ListComplianceItemsOutput, error) 
 ListComplianceSummaries(ctx context.Context, params *ListComplianceSummariesInput, optFns ...func(*Options)) (*ListComplianceSummariesOutput, error) 
 ListDocumentMetadataHistory(ctx context.Context, params *ListDocumentMetadataHistoryInput, optFns ...func(*Options)) (*ListDocumentMetadataHistoryOutput, error) 
 ListDocumentVersions(ctx context.Context, params *ListDocumentVersionsInput, optFns ...func(*Options)) (*ListDocumentVersionsOutput, error) 
 ListDocuments(ctx context.Context, params *ListDocumentsInput, optFns ...func(*Options)) (*ListDocumentsOutput, error) 
 ListInventoryEntries(ctx context.Context, params *ListInventoryEntriesInput, optFns ...func(*Options)) (*ListInventoryEntriesOutput, error) 
 ListOpsItemEvents(ctx context.Context, params *ListOpsItemEventsInput, optFns ...func(*Options)) (*ListOpsItemEventsOutput, error) 
 ListOpsItemRelatedItems(ctx context.Context, params *ListOpsItemRelatedItemsInput, optFns ...func(*Options)) (*ListOpsItemRelatedItemsOutput, error) 
 ListOpsMetadata(ctx context.Context, params *ListOpsMetadataInput, optFns ...func(*Options)) (*ListOpsMetadataOutput, error) 
 ListResourceComplianceSummaries(ctx context.Context, params *ListResourceComplianceSummariesInput, optFns ...func(*Options)) (*ListResourceComplianceSummariesOutput, error) 
 ListResourceDataSync(ctx context.Context, params *ListResourceDataSyncInput, optFns ...func(*Options)) (*ListResourceDataSyncOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyDocumentPermission(ctx context.Context, params *ModifyDocumentPermissionInput, optFns ...func(*Options)) (*ModifyDocumentPermissionOutput, error) 
 PutComplianceItems(ctx context.Context, params *PutComplianceItemsInput, optFns ...func(*Options)) (*PutComplianceItemsOutput, error) 
 PutInventory(ctx context.Context, params *PutInventoryInput, optFns ...func(*Options)) (*PutInventoryOutput, error) 
 PutParameter(ctx context.Context, params *PutParameterInput, optFns ...func(*Options)) (*PutParameterOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 RegisterDefaultPatchBaseline(ctx context.Context, params *RegisterDefaultPatchBaselineInput, optFns ...func(*Options)) (*RegisterDefaultPatchBaselineOutput, error) 
 RegisterPatchBaselineForPatchGroup(ctx context.Context, params *RegisterPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*RegisterPatchBaselineForPatchGroupOutput, error) 
 RegisterTargetWithMaintenanceWindow(ctx context.Context, params *RegisterTargetWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTargetWithMaintenanceWindowOutput, error) 
 RegisterTaskWithMaintenanceWindow(ctx context.Context, params *RegisterTaskWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTaskWithMaintenanceWindowOutput, error) 
 RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) 
 ResetServiceSetting(ctx context.Context, params *ResetServiceSettingInput, optFns ...func(*Options)) (*ResetServiceSettingOutput, error) 
 ResumeSession(ctx context.Context, params *ResumeSessionInput, optFns ...func(*Options)) (*ResumeSessionOutput, error) 
 SendAutomationSignal(ctx context.Context, params *SendAutomationSignalInput, optFns ...func(*Options)) (*SendAutomationSignalOutput, error) 
 SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(*Options)) (*SendCommandOutput, error) 
 StartAssociationsOnce(ctx context.Context, params *StartAssociationsOnceInput, optFns ...func(*Options)) (*StartAssociationsOnceOutput, error) 
 StartAutomationExecution(ctx context.Context, params *StartAutomationExecutionInput, optFns ...func(*Options)) (*StartAutomationExecutionOutput, error) 
 StartChangeRequestExecution(ctx context.Context, params *StartChangeRequestExecutionInput, optFns ...func(*Options)) (*StartChangeRequestExecutionOutput, error) 
 StartSession(ctx context.Context, params *StartSessionInput, optFns ...func(*Options)) (*StartSessionOutput, error) 
 StopAutomationExecution(ctx context.Context, params *StopAutomationExecutionInput, optFns ...func(*Options)) (*StopAutomationExecutionOutput, error) 
 TerminateSession(ctx context.Context, params *TerminateSessionInput, optFns ...func(*Options)) (*TerminateSessionOutput, error) 
 UnlabelParameterVersion(ctx context.Context, params *UnlabelParameterVersionInput, optFns ...func(*Options)) (*UnlabelParameterVersionOutput, error) 
 UpdateAssociation(ctx context.Context, params *UpdateAssociationInput, optFns ...func(*Options)) (*UpdateAssociationOutput, error) 
 UpdateAssociationStatus(ctx context.Context, params *UpdateAssociationStatusInput, optFns ...func(*Options)) (*UpdateAssociationStatusOutput, error) 
 UpdateDocument(ctx context.Context, params *UpdateDocumentInput, optFns ...func(*Options)) (*UpdateDocumentOutput, error) 
 UpdateDocumentDefaultVersion(ctx context.Context, params *UpdateDocumentDefaultVersionInput, optFns ...func(*Options)) (*UpdateDocumentDefaultVersionOutput, error) 
 UpdateDocumentMetadata(ctx context.Context, params *UpdateDocumentMetadataInput, optFns ...func(*Options)) (*UpdateDocumentMetadataOutput, error) 
 UpdateMaintenanceWindow(ctx context.Context, params *UpdateMaintenanceWindowInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowOutput, error) 
 UpdateMaintenanceWindowTarget(ctx context.Context, params *UpdateMaintenanceWindowTargetInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowTargetOutput, error) 
 UpdateMaintenanceWindowTask(ctx context.Context, params *UpdateMaintenanceWindowTaskInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowTaskOutput, error) 
 UpdateManagedInstanceRole(ctx context.Context, params *UpdateManagedInstanceRoleInput, optFns ...func(*Options)) (*UpdateManagedInstanceRoleOutput, error) 
 UpdateOpsItem(ctx context.Context, params *UpdateOpsItemInput, optFns ...func(*Options)) (*UpdateOpsItemOutput, error) 
 UpdateOpsMetadata(ctx context.Context, params *UpdateOpsMetadataInput, optFns ...func(*Options)) (*UpdateOpsMetadataOutput, error) 
 UpdatePatchBaseline(ctx context.Context, params *UpdatePatchBaselineInput, optFns ...func(*Options)) (*UpdatePatchBaselineOutput, error) 
 UpdateResourceDataSync(ctx context.Context, params *UpdateResourceDataSyncInput, optFns ...func(*Options)) (*UpdateResourceDataSyncOutput, error) 
 UpdateServiceSetting(ctx context.Context, params *UpdateServiceSettingInput, optFns ...func(*Options)) (*UpdateServiceSettingOutput, error) 
}