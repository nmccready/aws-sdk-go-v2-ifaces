package ssm_test

// tests for the ssm service interface for this ../../../service/ssm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssm/ssm_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsmServiceCanBeMocked(t *testing.T) {
	var iface ssm_iface.IClient
	iface = &ssm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &ssm.AddTagsToResourceInput{}
        output := &ssm.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateOpsItemRelatedItem", func(t *testing.T) {
        input := &ssm.AssociateOpsItemRelatedItemInput{}
        output := &ssm.AssociateOpsItemRelatedItemOutput{}

        mockClient.On("AssociateOpsItemRelatedItem", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateOpsItemRelatedItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCommand", func(t *testing.T) {
        input := &ssm.CancelCommandInput{}
        output := &ssm.CancelCommandOutput{}

        mockClient.On("CancelCommand", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCommand(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMaintenanceWindowExecution", func(t *testing.T) {
        input := &ssm.CancelMaintenanceWindowExecutionInput{}
        output := &ssm.CancelMaintenanceWindowExecutionOutput{}

        mockClient.On("CancelMaintenanceWindowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMaintenanceWindowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateActivation", func(t *testing.T) {
        input := &ssm.CreateActivationInput{}
        output := &ssm.CreateActivationOutput{}

        mockClient.On("CreateActivation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateActivation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssociation", func(t *testing.T) {
        input := &ssm.CreateAssociationInput{}
        output := &ssm.CreateAssociationOutput{}

        mockClient.On("CreateAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssociationBatch", func(t *testing.T) {
        input := &ssm.CreateAssociationBatchInput{}
        output := &ssm.CreateAssociationBatchOutput{}

        mockClient.On("CreateAssociationBatch", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssociationBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDocument", func(t *testing.T) {
        input := &ssm.CreateDocumentInput{}
        output := &ssm.CreateDocumentOutput{}

        mockClient.On("CreateDocument", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMaintenanceWindow", func(t *testing.T) {
        input := &ssm.CreateMaintenanceWindowInput{}
        output := &ssm.CreateMaintenanceWindowOutput{}

        mockClient.On("CreateMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOpsItem", func(t *testing.T) {
        input := &ssm.CreateOpsItemInput{}
        output := &ssm.CreateOpsItemOutput{}

        mockClient.On("CreateOpsItem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOpsItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOpsMetadata", func(t *testing.T) {
        input := &ssm.CreateOpsMetadataInput{}
        output := &ssm.CreateOpsMetadataOutput{}

        mockClient.On("CreateOpsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOpsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePatchBaseline", func(t *testing.T) {
        input := &ssm.CreatePatchBaselineInput{}
        output := &ssm.CreatePatchBaselineOutput{}

        mockClient.On("CreatePatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceDataSync", func(t *testing.T) {
        input := &ssm.CreateResourceDataSyncInput{}
        output := &ssm.CreateResourceDataSyncOutput{}

        mockClient.On("CreateResourceDataSync", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceDataSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteActivation", func(t *testing.T) {
        input := &ssm.DeleteActivationInput{}
        output := &ssm.DeleteActivationOutput{}

        mockClient.On("DeleteActivation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteActivation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssociation", func(t *testing.T) {
        input := &ssm.DeleteAssociationInput{}
        output := &ssm.DeleteAssociationOutput{}

        mockClient.On("DeleteAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocument", func(t *testing.T) {
        input := &ssm.DeleteDocumentInput{}
        output := &ssm.DeleteDocumentOutput{}

        mockClient.On("DeleteDocument", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInventory", func(t *testing.T) {
        input := &ssm.DeleteInventoryInput{}
        output := &ssm.DeleteInventoryOutput{}

        mockClient.On("DeleteInventory", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInventory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMaintenanceWindow", func(t *testing.T) {
        input := &ssm.DeleteMaintenanceWindowInput{}
        output := &ssm.DeleteMaintenanceWindowOutput{}

        mockClient.On("DeleteMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOpsItem", func(t *testing.T) {
        input := &ssm.DeleteOpsItemInput{}
        output := &ssm.DeleteOpsItemOutput{}

        mockClient.On("DeleteOpsItem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOpsItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOpsMetadata", func(t *testing.T) {
        input := &ssm.DeleteOpsMetadataInput{}
        output := &ssm.DeleteOpsMetadataOutput{}

        mockClient.On("DeleteOpsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOpsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteParameter", func(t *testing.T) {
        input := &ssm.DeleteParameterInput{}
        output := &ssm.DeleteParameterOutput{}

        mockClient.On("DeleteParameter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteParameter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteParameters", func(t *testing.T) {
        input := &ssm.DeleteParametersInput{}
        output := &ssm.DeleteParametersOutput{}

        mockClient.On("DeleteParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePatchBaseline", func(t *testing.T) {
        input := &ssm.DeletePatchBaselineInput{}
        output := &ssm.DeletePatchBaselineOutput{}

        mockClient.On("DeletePatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceDataSync", func(t *testing.T) {
        input := &ssm.DeleteResourceDataSyncInput{}
        output := &ssm.DeleteResourceDataSyncOutput{}

        mockClient.On("DeleteResourceDataSync", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceDataSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &ssm.DeleteResourcePolicyInput{}
        output := &ssm.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterManagedInstance", func(t *testing.T) {
        input := &ssm.DeregisterManagedInstanceInput{}
        output := &ssm.DeregisterManagedInstanceOutput{}

        mockClient.On("DeregisterManagedInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterManagedInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterPatchBaselineForPatchGroup", func(t *testing.T) {
        input := &ssm.DeregisterPatchBaselineForPatchGroupInput{}
        output := &ssm.DeregisterPatchBaselineForPatchGroupOutput{}

        mockClient.On("DeregisterPatchBaselineForPatchGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterPatchBaselineForPatchGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTargetFromMaintenanceWindow", func(t *testing.T) {
        input := &ssm.DeregisterTargetFromMaintenanceWindowInput{}
        output := &ssm.DeregisterTargetFromMaintenanceWindowOutput{}

        mockClient.On("DeregisterTargetFromMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTargetFromMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTaskFromMaintenanceWindow", func(t *testing.T) {
        input := &ssm.DeregisterTaskFromMaintenanceWindowInput{}
        output := &ssm.DeregisterTaskFromMaintenanceWindowOutput{}

        mockClient.On("DeregisterTaskFromMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTaskFromMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActivations", func(t *testing.T) {
        input := &ssm.DescribeActivationsInput{}
        output := &ssm.DescribeActivationsOutput{}

        mockClient.On("DescribeActivations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActivations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssociation", func(t *testing.T) {
        input := &ssm.DescribeAssociationInput{}
        output := &ssm.DescribeAssociationOutput{}

        mockClient.On("DescribeAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssociationExecutionTargets", func(t *testing.T) {
        input := &ssm.DescribeAssociationExecutionTargetsInput{}
        output := &ssm.DescribeAssociationExecutionTargetsOutput{}

        mockClient.On("DescribeAssociationExecutionTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssociationExecutionTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssociationExecutions", func(t *testing.T) {
        input := &ssm.DescribeAssociationExecutionsInput{}
        output := &ssm.DescribeAssociationExecutionsOutput{}

        mockClient.On("DescribeAssociationExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssociationExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutomationExecutions", func(t *testing.T) {
        input := &ssm.DescribeAutomationExecutionsInput{}
        output := &ssm.DescribeAutomationExecutionsOutput{}

        mockClient.On("DescribeAutomationExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutomationExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutomationStepExecutions", func(t *testing.T) {
        input := &ssm.DescribeAutomationStepExecutionsInput{}
        output := &ssm.DescribeAutomationStepExecutionsOutput{}

        mockClient.On("DescribeAutomationStepExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutomationStepExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAvailablePatches", func(t *testing.T) {
        input := &ssm.DescribeAvailablePatchesInput{}
        output := &ssm.DescribeAvailablePatchesOutput{}

        mockClient.On("DescribeAvailablePatches", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAvailablePatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDocument", func(t *testing.T) {
        input := &ssm.DescribeDocumentInput{}
        output := &ssm.DescribeDocumentOutput{}

        mockClient.On("DescribeDocument", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDocumentPermission", func(t *testing.T) {
        input := &ssm.DescribeDocumentPermissionInput{}
        output := &ssm.DescribeDocumentPermissionOutput{}

        mockClient.On("DescribeDocumentPermission", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDocumentPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEffectiveInstanceAssociations", func(t *testing.T) {
        input := &ssm.DescribeEffectiveInstanceAssociationsInput{}
        output := &ssm.DescribeEffectiveInstanceAssociationsOutput{}

        mockClient.On("DescribeEffectiveInstanceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEffectiveInstanceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEffectivePatchesForPatchBaseline", func(t *testing.T) {
        input := &ssm.DescribeEffectivePatchesForPatchBaselineInput{}
        output := &ssm.DescribeEffectivePatchesForPatchBaselineOutput{}

        mockClient.On("DescribeEffectivePatchesForPatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEffectivePatchesForPatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceAssociationsStatus", func(t *testing.T) {
        input := &ssm.DescribeInstanceAssociationsStatusInput{}
        output := &ssm.DescribeInstanceAssociationsStatusOutput{}

        mockClient.On("DescribeInstanceAssociationsStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceAssociationsStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceInformation", func(t *testing.T) {
        input := &ssm.DescribeInstanceInformationInput{}
        output := &ssm.DescribeInstanceInformationOutput{}

        mockClient.On("DescribeInstanceInformation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstancePatchStates", func(t *testing.T) {
        input := &ssm.DescribeInstancePatchStatesInput{}
        output := &ssm.DescribeInstancePatchStatesOutput{}

        mockClient.On("DescribeInstancePatchStates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstancePatchStates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstancePatchStatesForPatchGroup", func(t *testing.T) {
        input := &ssm.DescribeInstancePatchStatesForPatchGroupInput{}
        output := &ssm.DescribeInstancePatchStatesForPatchGroupOutput{}

        mockClient.On("DescribeInstancePatchStatesForPatchGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstancePatchStatesForPatchGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstancePatches", func(t *testing.T) {
        input := &ssm.DescribeInstancePatchesInput{}
        output := &ssm.DescribeInstancePatchesOutput{}

        mockClient.On("DescribeInstancePatches", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstancePatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceProperties", func(t *testing.T) {
        input := &ssm.DescribeInstancePropertiesInput{}
        output := &ssm.DescribeInstancePropertiesOutput{}

        mockClient.On("DescribeInstanceProperties", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInventoryDeletions", func(t *testing.T) {
        input := &ssm.DescribeInventoryDeletionsInput{}
        output := &ssm.DescribeInventoryDeletionsOutput{}

        mockClient.On("DescribeInventoryDeletions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInventoryDeletions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowExecutionTaskInvocations", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
        output := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput{}

        mockClient.On("DescribeMaintenanceWindowExecutionTaskInvocations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowExecutionTaskInvocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowExecutionTasks", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowExecutionTasksInput{}
        output := &ssm.DescribeMaintenanceWindowExecutionTasksOutput{}

        mockClient.On("DescribeMaintenanceWindowExecutionTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowExecutionTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowExecutions", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowExecutionsInput{}
        output := &ssm.DescribeMaintenanceWindowExecutionsOutput{}

        mockClient.On("DescribeMaintenanceWindowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowSchedule", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowScheduleInput{}
        output := &ssm.DescribeMaintenanceWindowScheduleOutput{}

        mockClient.On("DescribeMaintenanceWindowSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowTargets", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowTargetsInput{}
        output := &ssm.DescribeMaintenanceWindowTargetsOutput{}

        mockClient.On("DescribeMaintenanceWindowTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowTasks", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowTasksInput{}
        output := &ssm.DescribeMaintenanceWindowTasksOutput{}

        mockClient.On("DescribeMaintenanceWindowTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindows", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowsInput{}
        output := &ssm.DescribeMaintenanceWindowsOutput{}

        mockClient.On("DescribeMaintenanceWindows", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMaintenanceWindowsForTarget", func(t *testing.T) {
        input := &ssm.DescribeMaintenanceWindowsForTargetInput{}
        output := &ssm.DescribeMaintenanceWindowsForTargetOutput{}

        mockClient.On("DescribeMaintenanceWindowsForTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMaintenanceWindowsForTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOpsItems", func(t *testing.T) {
        input := &ssm.DescribeOpsItemsInput{}
        output := &ssm.DescribeOpsItemsOutput{}

        mockClient.On("DescribeOpsItems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOpsItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeParameters", func(t *testing.T) {
        input := &ssm.DescribeParametersInput{}
        output := &ssm.DescribeParametersOutput{}

        mockClient.On("DescribeParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePatchBaselines", func(t *testing.T) {
        input := &ssm.DescribePatchBaselinesInput{}
        output := &ssm.DescribePatchBaselinesOutput{}

        mockClient.On("DescribePatchBaselines", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePatchBaselines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePatchGroupState", func(t *testing.T) {
        input := &ssm.DescribePatchGroupStateInput{}
        output := &ssm.DescribePatchGroupStateOutput{}

        mockClient.On("DescribePatchGroupState", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePatchGroupState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePatchGroups", func(t *testing.T) {
        input := &ssm.DescribePatchGroupsInput{}
        output := &ssm.DescribePatchGroupsOutput{}

        mockClient.On("DescribePatchGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePatchGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePatchProperties", func(t *testing.T) {
        input := &ssm.DescribePatchPropertiesInput{}
        output := &ssm.DescribePatchPropertiesOutput{}

        mockClient.On("DescribePatchProperties", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePatchProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSessions", func(t *testing.T) {
        input := &ssm.DescribeSessionsInput{}
        output := &ssm.DescribeSessionsOutput{}

        mockClient.On("DescribeSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateOpsItemRelatedItem", func(t *testing.T) {
        input := &ssm.DisassociateOpsItemRelatedItemInput{}
        output := &ssm.DisassociateOpsItemRelatedItemOutput{}

        mockClient.On("DisassociateOpsItemRelatedItem", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateOpsItemRelatedItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutomationExecution", func(t *testing.T) {
        input := &ssm.GetAutomationExecutionInput{}
        output := &ssm.GetAutomationExecutionOutput{}

        mockClient.On("GetAutomationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutomationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalendarState", func(t *testing.T) {
        input := &ssm.GetCalendarStateInput{}
        output := &ssm.GetCalendarStateOutput{}

        mockClient.On("GetCalendarState", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalendarState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCommandInvocation", func(t *testing.T) {
        input := &ssm.GetCommandInvocationInput{}
        output := &ssm.GetCommandInvocationOutput{}

        mockClient.On("GetCommandInvocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetCommandInvocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectionStatus", func(t *testing.T) {
        input := &ssm.GetConnectionStatusInput{}
        output := &ssm.GetConnectionStatusOutput{}

        mockClient.On("GetConnectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDefaultPatchBaseline", func(t *testing.T) {
        input := &ssm.GetDefaultPatchBaselineInput{}
        output := &ssm.GetDefaultPatchBaselineOutput{}

        mockClient.On("GetDefaultPatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.GetDefaultPatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployablePatchSnapshotForInstance", func(t *testing.T) {
        input := &ssm.GetDeployablePatchSnapshotForInstanceInput{}
        output := &ssm.GetDeployablePatchSnapshotForInstanceOutput{}

        mockClient.On("GetDeployablePatchSnapshotForInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployablePatchSnapshotForInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocument", func(t *testing.T) {
        input := &ssm.GetDocumentInput{}
        output := &ssm.GetDocumentOutput{}

        mockClient.On("GetDocument", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInventory", func(t *testing.T) {
        input := &ssm.GetInventoryInput{}
        output := &ssm.GetInventoryOutput{}

        mockClient.On("GetInventory", ctx, input).Return(output, nil)

        result, err := mockClient.GetInventory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInventorySchema", func(t *testing.T) {
        input := &ssm.GetInventorySchemaInput{}
        output := &ssm.GetInventorySchemaOutput{}

        mockClient.On("GetInventorySchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetInventorySchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMaintenanceWindow", func(t *testing.T) {
        input := &ssm.GetMaintenanceWindowInput{}
        output := &ssm.GetMaintenanceWindowOutput{}

        mockClient.On("GetMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.GetMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMaintenanceWindowExecution", func(t *testing.T) {
        input := &ssm.GetMaintenanceWindowExecutionInput{}
        output := &ssm.GetMaintenanceWindowExecutionOutput{}

        mockClient.On("GetMaintenanceWindowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetMaintenanceWindowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMaintenanceWindowExecutionTask", func(t *testing.T) {
        input := &ssm.GetMaintenanceWindowExecutionTaskInput{}
        output := &ssm.GetMaintenanceWindowExecutionTaskOutput{}

        mockClient.On("GetMaintenanceWindowExecutionTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetMaintenanceWindowExecutionTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMaintenanceWindowExecutionTaskInvocation", func(t *testing.T) {
        input := &ssm.GetMaintenanceWindowExecutionTaskInvocationInput{}
        output := &ssm.GetMaintenanceWindowExecutionTaskInvocationOutput{}

        mockClient.On("GetMaintenanceWindowExecutionTaskInvocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetMaintenanceWindowExecutionTaskInvocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMaintenanceWindowTask", func(t *testing.T) {
        input := &ssm.GetMaintenanceWindowTaskInput{}
        output := &ssm.GetMaintenanceWindowTaskOutput{}

        mockClient.On("GetMaintenanceWindowTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetMaintenanceWindowTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpsItem", func(t *testing.T) {
        input := &ssm.GetOpsItemInput{}
        output := &ssm.GetOpsItemOutput{}

        mockClient.On("GetOpsItem", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpsItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpsMetadata", func(t *testing.T) {
        input := &ssm.GetOpsMetadataInput{}
        output := &ssm.GetOpsMetadataOutput{}

        mockClient.On("GetOpsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpsSummary", func(t *testing.T) {
        input := &ssm.GetOpsSummaryInput{}
        output := &ssm.GetOpsSummaryOutput{}

        mockClient.On("GetOpsSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpsSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParameter", func(t *testing.T) {
        input := &ssm.GetParameterInput{}
        output := &ssm.GetParameterOutput{}

        mockClient.On("GetParameter", ctx, input).Return(output, nil)

        result, err := mockClient.GetParameter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParameterHistory", func(t *testing.T) {
        input := &ssm.GetParameterHistoryInput{}
        output := &ssm.GetParameterHistoryOutput{}

        mockClient.On("GetParameterHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetParameterHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParameters", func(t *testing.T) {
        input := &ssm.GetParametersInput{}
        output := &ssm.GetParametersOutput{}

        mockClient.On("GetParameters", ctx, input).Return(output, nil)

        result, err := mockClient.GetParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParametersByPath", func(t *testing.T) {
        input := &ssm.GetParametersByPathInput{}
        output := &ssm.GetParametersByPathOutput{}

        mockClient.On("GetParametersByPath", ctx, input).Return(output, nil)

        result, err := mockClient.GetParametersByPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPatchBaseline", func(t *testing.T) {
        input := &ssm.GetPatchBaselineInput{}
        output := &ssm.GetPatchBaselineOutput{}

        mockClient.On("GetPatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.GetPatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPatchBaselineForPatchGroup", func(t *testing.T) {
        input := &ssm.GetPatchBaselineForPatchGroupInput{}
        output := &ssm.GetPatchBaselineForPatchGroupOutput{}

        mockClient.On("GetPatchBaselineForPatchGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetPatchBaselineForPatchGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicies", func(t *testing.T) {
        input := &ssm.GetResourcePoliciesInput{}
        output := &ssm.GetResourcePoliciesOutput{}

        mockClient.On("GetResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceSetting", func(t *testing.T) {
        input := &ssm.GetServiceSettingInput{}
        output := &ssm.GetServiceSettingOutput{}

        mockClient.On("GetServiceSetting", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceSetting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLabelParameterVersion", func(t *testing.T) {
        input := &ssm.LabelParameterVersionInput{}
        output := &ssm.LabelParameterVersionOutput{}

        mockClient.On("LabelParameterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.LabelParameterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociationVersions", func(t *testing.T) {
        input := &ssm.ListAssociationVersionsInput{}
        output := &ssm.ListAssociationVersionsOutput{}

        mockClient.On("ListAssociationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociations", func(t *testing.T) {
        input := &ssm.ListAssociationsInput{}
        output := &ssm.ListAssociationsOutput{}

        mockClient.On("ListAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCommandInvocations", func(t *testing.T) {
        input := &ssm.ListCommandInvocationsInput{}
        output := &ssm.ListCommandInvocationsOutput{}

        mockClient.On("ListCommandInvocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCommandInvocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCommands", func(t *testing.T) {
        input := &ssm.ListCommandsInput{}
        output := &ssm.ListCommandsOutput{}

        mockClient.On("ListCommands", ctx, input).Return(output, nil)

        result, err := mockClient.ListCommands(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComplianceItems", func(t *testing.T) {
        input := &ssm.ListComplianceItemsInput{}
        output := &ssm.ListComplianceItemsOutput{}

        mockClient.On("ListComplianceItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListComplianceItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComplianceSummaries", func(t *testing.T) {
        input := &ssm.ListComplianceSummariesInput{}
        output := &ssm.ListComplianceSummariesOutput{}

        mockClient.On("ListComplianceSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListComplianceSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocumentMetadataHistory", func(t *testing.T) {
        input := &ssm.ListDocumentMetadataHistoryInput{}
        output := &ssm.ListDocumentMetadataHistoryOutput{}

        mockClient.On("ListDocumentMetadataHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocumentMetadataHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocumentVersions", func(t *testing.T) {
        input := &ssm.ListDocumentVersionsInput{}
        output := &ssm.ListDocumentVersionsOutput{}

        mockClient.On("ListDocumentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocumentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocuments", func(t *testing.T) {
        input := &ssm.ListDocumentsInput{}
        output := &ssm.ListDocumentsOutput{}

        mockClient.On("ListDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInventoryEntries", func(t *testing.T) {
        input := &ssm.ListInventoryEntriesInput{}
        output := &ssm.ListInventoryEntriesOutput{}

        mockClient.On("ListInventoryEntries", ctx, input).Return(output, nil)

        result, err := mockClient.ListInventoryEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpsItemEvents", func(t *testing.T) {
        input := &ssm.ListOpsItemEventsInput{}
        output := &ssm.ListOpsItemEventsOutput{}

        mockClient.On("ListOpsItemEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpsItemEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpsItemRelatedItems", func(t *testing.T) {
        input := &ssm.ListOpsItemRelatedItemsInput{}
        output := &ssm.ListOpsItemRelatedItemsOutput{}

        mockClient.On("ListOpsItemRelatedItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpsItemRelatedItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpsMetadata", func(t *testing.T) {
        input := &ssm.ListOpsMetadataInput{}
        output := &ssm.ListOpsMetadataOutput{}

        mockClient.On("ListOpsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceComplianceSummaries", func(t *testing.T) {
        input := &ssm.ListResourceComplianceSummariesInput{}
        output := &ssm.ListResourceComplianceSummariesOutput{}

        mockClient.On("ListResourceComplianceSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceComplianceSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceDataSync", func(t *testing.T) {
        input := &ssm.ListResourceDataSyncInput{}
        output := &ssm.ListResourceDataSyncOutput{}

        mockClient.On("ListResourceDataSync", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceDataSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ssm.ListTagsForResourceInput{}
        output := &ssm.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyDocumentPermission", func(t *testing.T) {
        input := &ssm.ModifyDocumentPermissionInput{}
        output := &ssm.ModifyDocumentPermissionOutput{}

        mockClient.On("ModifyDocumentPermission", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyDocumentPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutComplianceItems", func(t *testing.T) {
        input := &ssm.PutComplianceItemsInput{}
        output := &ssm.PutComplianceItemsOutput{}

        mockClient.On("PutComplianceItems", ctx, input).Return(output, nil)

        result, err := mockClient.PutComplianceItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInventory", func(t *testing.T) {
        input := &ssm.PutInventoryInput{}
        output := &ssm.PutInventoryOutput{}

        mockClient.On("PutInventory", ctx, input).Return(output, nil)

        result, err := mockClient.PutInventory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutParameter", func(t *testing.T) {
        input := &ssm.PutParameterInput{}
        output := &ssm.PutParameterOutput{}

        mockClient.On("PutParameter", ctx, input).Return(output, nil)

        result, err := mockClient.PutParameter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &ssm.PutResourcePolicyInput{}
        output := &ssm.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDefaultPatchBaseline", func(t *testing.T) {
        input := &ssm.RegisterDefaultPatchBaselineInput{}
        output := &ssm.RegisterDefaultPatchBaselineOutput{}

        mockClient.On("RegisterDefaultPatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDefaultPatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterPatchBaselineForPatchGroup", func(t *testing.T) {
        input := &ssm.RegisterPatchBaselineForPatchGroupInput{}
        output := &ssm.RegisterPatchBaselineForPatchGroupOutput{}

        mockClient.On("RegisterPatchBaselineForPatchGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterPatchBaselineForPatchGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTargetWithMaintenanceWindow", func(t *testing.T) {
        input := &ssm.RegisterTargetWithMaintenanceWindowInput{}
        output := &ssm.RegisterTargetWithMaintenanceWindowOutput{}

        mockClient.On("RegisterTargetWithMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTargetWithMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTaskWithMaintenanceWindow", func(t *testing.T) {
        input := &ssm.RegisterTaskWithMaintenanceWindowInput{}
        output := &ssm.RegisterTaskWithMaintenanceWindowOutput{}

        mockClient.On("RegisterTaskWithMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTaskWithMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &ssm.RemoveTagsFromResourceInput{}
        output := &ssm.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetServiceSetting", func(t *testing.T) {
        input := &ssm.ResetServiceSettingInput{}
        output := &ssm.ResetServiceSettingOutput{}

        mockClient.On("ResetServiceSetting", ctx, input).Return(output, nil)

        result, err := mockClient.ResetServiceSetting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeSession", func(t *testing.T) {
        input := &ssm.ResumeSessionInput{}
        output := &ssm.ResumeSessionOutput{}

        mockClient.On("ResumeSession", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendAutomationSignal", func(t *testing.T) {
        input := &ssm.SendAutomationSignalInput{}
        output := &ssm.SendAutomationSignalOutput{}

        mockClient.On("SendAutomationSignal", ctx, input).Return(output, nil)

        result, err := mockClient.SendAutomationSignal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCommand", func(t *testing.T) {
        input := &ssm.SendCommandInput{}
        output := &ssm.SendCommandOutput{}

        mockClient.On("SendCommand", ctx, input).Return(output, nil)

        result, err := mockClient.SendCommand(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssociationsOnce", func(t *testing.T) {
        input := &ssm.StartAssociationsOnceInput{}
        output := &ssm.StartAssociationsOnceOutput{}

        mockClient.On("StartAssociationsOnce", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssociationsOnce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAutomationExecution", func(t *testing.T) {
        input := &ssm.StartAutomationExecutionInput{}
        output := &ssm.StartAutomationExecutionOutput{}

        mockClient.On("StartAutomationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartAutomationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartChangeRequestExecution", func(t *testing.T) {
        input := &ssm.StartChangeRequestExecutionInput{}
        output := &ssm.StartChangeRequestExecutionOutput{}

        mockClient.On("StartChangeRequestExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartChangeRequestExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSession", func(t *testing.T) {
        input := &ssm.StartSessionInput{}
        output := &ssm.StartSessionOutput{}

        mockClient.On("StartSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAutomationExecution", func(t *testing.T) {
        input := &ssm.StopAutomationExecutionInput{}
        output := &ssm.StopAutomationExecutionOutput{}

        mockClient.On("StopAutomationExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopAutomationExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateSession", func(t *testing.T) {
        input := &ssm.TerminateSessionInput{}
        output := &ssm.TerminateSessionOutput{}

        mockClient.On("TerminateSession", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnlabelParameterVersion", func(t *testing.T) {
        input := &ssm.UnlabelParameterVersionInput{}
        output := &ssm.UnlabelParameterVersionOutput{}

        mockClient.On("UnlabelParameterVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UnlabelParameterVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssociation", func(t *testing.T) {
        input := &ssm.UpdateAssociationInput{}
        output := &ssm.UpdateAssociationOutput{}

        mockClient.On("UpdateAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssociationStatus", func(t *testing.T) {
        input := &ssm.UpdateAssociationStatusInput{}
        output := &ssm.UpdateAssociationStatusOutput{}

        mockClient.On("UpdateAssociationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssociationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocument", func(t *testing.T) {
        input := &ssm.UpdateDocumentInput{}
        output := &ssm.UpdateDocumentOutput{}

        mockClient.On("UpdateDocument", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocumentDefaultVersion", func(t *testing.T) {
        input := &ssm.UpdateDocumentDefaultVersionInput{}
        output := &ssm.UpdateDocumentDefaultVersionOutput{}

        mockClient.On("UpdateDocumentDefaultVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocumentDefaultVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocumentMetadata", func(t *testing.T) {
        input := &ssm.UpdateDocumentMetadataInput{}
        output := &ssm.UpdateDocumentMetadataOutput{}

        mockClient.On("UpdateDocumentMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocumentMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMaintenanceWindow", func(t *testing.T) {
        input := &ssm.UpdateMaintenanceWindowInput{}
        output := &ssm.UpdateMaintenanceWindowOutput{}

        mockClient.On("UpdateMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMaintenanceWindowTarget", func(t *testing.T) {
        input := &ssm.UpdateMaintenanceWindowTargetInput{}
        output := &ssm.UpdateMaintenanceWindowTargetOutput{}

        mockClient.On("UpdateMaintenanceWindowTarget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMaintenanceWindowTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMaintenanceWindowTask", func(t *testing.T) {
        input := &ssm.UpdateMaintenanceWindowTaskInput{}
        output := &ssm.UpdateMaintenanceWindowTaskOutput{}

        mockClient.On("UpdateMaintenanceWindowTask", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMaintenanceWindowTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateManagedInstanceRole", func(t *testing.T) {
        input := &ssm.UpdateManagedInstanceRoleInput{}
        output := &ssm.UpdateManagedInstanceRoleOutput{}

        mockClient.On("UpdateManagedInstanceRole", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateManagedInstanceRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOpsItem", func(t *testing.T) {
        input := &ssm.UpdateOpsItemInput{}
        output := &ssm.UpdateOpsItemOutput{}

        mockClient.On("UpdateOpsItem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOpsItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOpsMetadata", func(t *testing.T) {
        input := &ssm.UpdateOpsMetadataInput{}
        output := &ssm.UpdateOpsMetadataOutput{}

        mockClient.On("UpdateOpsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOpsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePatchBaseline", func(t *testing.T) {
        input := &ssm.UpdatePatchBaselineInput{}
        output := &ssm.UpdatePatchBaselineOutput{}

        mockClient.On("UpdatePatchBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePatchBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceDataSync", func(t *testing.T) {
        input := &ssm.UpdateResourceDataSyncInput{}
        output := &ssm.UpdateResourceDataSyncOutput{}

        mockClient.On("UpdateResourceDataSync", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceDataSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSetting", func(t *testing.T) {
        input := &ssm.UpdateServiceSettingInput{}
        output := &ssm.UpdateServiceSettingOutput{}

        mockClient.On("UpdateServiceSetting", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSetting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
