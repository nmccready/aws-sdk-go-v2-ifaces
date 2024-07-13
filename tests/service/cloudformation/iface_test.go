package cloudformation_test

// tests for the cloudformation service interface for this ../../../service/cloudformation/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudformation/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudformation/cloudformation_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudformationServiceCanBeMocked(t *testing.T) {
	var iface cloudformation_iface.IClient
	iface = &cloudformation.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudformation.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateOrganizationsAccess", func(t *testing.T) {
        input := &cloudformation.ActivateOrganizationsAccessInput{}
        output := &cloudformation.ActivateOrganizationsAccessOutput{}

        mockClient.On("ActivateOrganizationsAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateOrganizationsAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateType", func(t *testing.T) {
        input := &cloudformation.ActivateTypeInput{}
        output := &cloudformation.ActivateTypeOutput{}

        mockClient.On("ActivateType", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDescribeTypeConfigurations", func(t *testing.T) {
        input := &cloudformation.BatchDescribeTypeConfigurationsInput{}
        output := &cloudformation.BatchDescribeTypeConfigurationsOutput{}

        mockClient.On("BatchDescribeTypeConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDescribeTypeConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelUpdateStack", func(t *testing.T) {
        input := &cloudformation.CancelUpdateStackInput{}
        output := &cloudformation.CancelUpdateStackOutput{}

        mockClient.On("CancelUpdateStack", ctx, input).Return(output, nil)

        result, err := mockClient.CancelUpdateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestContinueUpdateRollback", func(t *testing.T) {
        input := &cloudformation.ContinueUpdateRollbackInput{}
        output := &cloudformation.ContinueUpdateRollbackOutput{}

        mockClient.On("ContinueUpdateRollback", ctx, input).Return(output, nil)

        result, err := mockClient.ContinueUpdateRollback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChangeSet", func(t *testing.T) {
        input := &cloudformation.CreateChangeSetInput{}
        output := &cloudformation.CreateChangeSetOutput{}

        mockClient.On("CreateChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGeneratedTemplate", func(t *testing.T) {
        input := &cloudformation.CreateGeneratedTemplateInput{}
        output := &cloudformation.CreateGeneratedTemplateOutput{}

        mockClient.On("CreateGeneratedTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGeneratedTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStack", func(t *testing.T) {
        input := &cloudformation.CreateStackInput{}
        output := &cloudformation.CreateStackOutput{}

        mockClient.On("CreateStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStackInstances", func(t *testing.T) {
        input := &cloudformation.CreateStackInstancesInput{}
        output := &cloudformation.CreateStackInstancesOutput{}

        mockClient.On("CreateStackInstances", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStackInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStackSet", func(t *testing.T) {
        input := &cloudformation.CreateStackSetInput{}
        output := &cloudformation.CreateStackSetOutput{}

        mockClient.On("CreateStackSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStackSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateOrganizationsAccess", func(t *testing.T) {
        input := &cloudformation.DeactivateOrganizationsAccessInput{}
        output := &cloudformation.DeactivateOrganizationsAccessOutput{}

        mockClient.On("DeactivateOrganizationsAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateOrganizationsAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateType", func(t *testing.T) {
        input := &cloudformation.DeactivateTypeInput{}
        output := &cloudformation.DeactivateTypeOutput{}

        mockClient.On("DeactivateType", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChangeSet", func(t *testing.T) {
        input := &cloudformation.DeleteChangeSetInput{}
        output := &cloudformation.DeleteChangeSetOutput{}

        mockClient.On("DeleteChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGeneratedTemplate", func(t *testing.T) {
        input := &cloudformation.DeleteGeneratedTemplateInput{}
        output := &cloudformation.DeleteGeneratedTemplateOutput{}

        mockClient.On("DeleteGeneratedTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGeneratedTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStack", func(t *testing.T) {
        input := &cloudformation.DeleteStackInput{}
        output := &cloudformation.DeleteStackOutput{}

        mockClient.On("DeleteStack", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStackInstances", func(t *testing.T) {
        input := &cloudformation.DeleteStackInstancesInput{}
        output := &cloudformation.DeleteStackInstancesOutput{}

        mockClient.On("DeleteStackInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStackInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStackSet", func(t *testing.T) {
        input := &cloudformation.DeleteStackSetInput{}
        output := &cloudformation.DeleteStackSetOutput{}

        mockClient.On("DeleteStackSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStackSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterType", func(t *testing.T) {
        input := &cloudformation.DeregisterTypeInput{}
        output := &cloudformation.DeregisterTypeOutput{}

        mockClient.On("DeregisterType", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountLimits", func(t *testing.T) {
        input := &cloudformation.DescribeAccountLimitsInput{}
        output := &cloudformation.DescribeAccountLimitsOutput{}

        mockClient.On("DescribeAccountLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChangeSet", func(t *testing.T) {
        input := &cloudformation.DescribeChangeSetInput{}
        output := &cloudformation.DescribeChangeSetOutput{}

        mockClient.On("DescribeChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChangeSetHooks", func(t *testing.T) {
        input := &cloudformation.DescribeChangeSetHooksInput{}
        output := &cloudformation.DescribeChangeSetHooksOutput{}

        mockClient.On("DescribeChangeSetHooks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChangeSetHooks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGeneratedTemplate", func(t *testing.T) {
        input := &cloudformation.DescribeGeneratedTemplateInput{}
        output := &cloudformation.DescribeGeneratedTemplateOutput{}

        mockClient.On("DescribeGeneratedTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGeneratedTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationsAccess", func(t *testing.T) {
        input := &cloudformation.DescribeOrganizationsAccessInput{}
        output := &cloudformation.DescribeOrganizationsAccessOutput{}

        mockClient.On("DescribeOrganizationsAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationsAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePublisher", func(t *testing.T) {
        input := &cloudformation.DescribePublisherInput{}
        output := &cloudformation.DescribePublisherOutput{}

        mockClient.On("DescribePublisher", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePublisher(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourceScan", func(t *testing.T) {
        input := &cloudformation.DescribeResourceScanInput{}
        output := &cloudformation.DescribeResourceScanOutput{}

        mockClient.On("DescribeResourceScan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourceScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackDriftDetectionStatus", func(t *testing.T) {
        input := &cloudformation.DescribeStackDriftDetectionStatusInput{}
        output := &cloudformation.DescribeStackDriftDetectionStatusOutput{}

        mockClient.On("DescribeStackDriftDetectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackDriftDetectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackEvents", func(t *testing.T) {
        input := &cloudformation.DescribeStackEventsInput{}
        output := &cloudformation.DescribeStackEventsOutput{}

        mockClient.On("DescribeStackEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackInstance", func(t *testing.T) {
        input := &cloudformation.DescribeStackInstanceInput{}
        output := &cloudformation.DescribeStackInstanceOutput{}

        mockClient.On("DescribeStackInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackResource", func(t *testing.T) {
        input := &cloudformation.DescribeStackResourceInput{}
        output := &cloudformation.DescribeStackResourceOutput{}

        mockClient.On("DescribeStackResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackResourceDrifts", func(t *testing.T) {
        input := &cloudformation.DescribeStackResourceDriftsInput{}
        output := &cloudformation.DescribeStackResourceDriftsOutput{}

        mockClient.On("DescribeStackResourceDrifts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackResourceDrifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackResources", func(t *testing.T) {
        input := &cloudformation.DescribeStackResourcesInput{}
        output := &cloudformation.DescribeStackResourcesOutput{}

        mockClient.On("DescribeStackResources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackSet", func(t *testing.T) {
        input := &cloudformation.DescribeStackSetInput{}
        output := &cloudformation.DescribeStackSetOutput{}

        mockClient.On("DescribeStackSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackSetOperation", func(t *testing.T) {
        input := &cloudformation.DescribeStackSetOperationInput{}
        output := &cloudformation.DescribeStackSetOperationOutput{}

        mockClient.On("DescribeStackSetOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackSetOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStacks", func(t *testing.T) {
        input := &cloudformation.DescribeStacksInput{}
        output := &cloudformation.DescribeStacksOutput{}

        mockClient.On("DescribeStacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeType", func(t *testing.T) {
        input := &cloudformation.DescribeTypeInput{}
        output := &cloudformation.DescribeTypeOutput{}

        mockClient.On("DescribeType", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTypeRegistration", func(t *testing.T) {
        input := &cloudformation.DescribeTypeRegistrationInput{}
        output := &cloudformation.DescribeTypeRegistrationOutput{}

        mockClient.On("DescribeTypeRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTypeRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectStackDrift", func(t *testing.T) {
        input := &cloudformation.DetectStackDriftInput{}
        output := &cloudformation.DetectStackDriftOutput{}

        mockClient.On("DetectStackDrift", ctx, input).Return(output, nil)

        result, err := mockClient.DetectStackDrift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectStackResourceDrift", func(t *testing.T) {
        input := &cloudformation.DetectStackResourceDriftInput{}
        output := &cloudformation.DetectStackResourceDriftOutput{}

        mockClient.On("DetectStackResourceDrift", ctx, input).Return(output, nil)

        result, err := mockClient.DetectStackResourceDrift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectStackSetDrift", func(t *testing.T) {
        input := &cloudformation.DetectStackSetDriftInput{}
        output := &cloudformation.DetectStackSetDriftOutput{}

        mockClient.On("DetectStackSetDrift", ctx, input).Return(output, nil)

        result, err := mockClient.DetectStackSetDrift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEstimateTemplateCost", func(t *testing.T) {
        input := &cloudformation.EstimateTemplateCostInput{}
        output := &cloudformation.EstimateTemplateCostOutput{}

        mockClient.On("EstimateTemplateCost", ctx, input).Return(output, nil)

        result, err := mockClient.EstimateTemplateCost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteChangeSet", func(t *testing.T) {
        input := &cloudformation.ExecuteChangeSetInput{}
        output := &cloudformation.ExecuteChangeSetOutput{}

        mockClient.On("ExecuteChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeneratedTemplate", func(t *testing.T) {
        input := &cloudformation.GetGeneratedTemplateInput{}
        output := &cloudformation.GetGeneratedTemplateOutput{}

        mockClient.On("GetGeneratedTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeneratedTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStackPolicy", func(t *testing.T) {
        input := &cloudformation.GetStackPolicyInput{}
        output := &cloudformation.GetStackPolicyOutput{}

        mockClient.On("GetStackPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetStackPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplate", func(t *testing.T) {
        input := &cloudformation.GetTemplateInput{}
        output := &cloudformation.GetTemplateOutput{}

        mockClient.On("GetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateSummary", func(t *testing.T) {
        input := &cloudformation.GetTemplateSummaryInput{}
        output := &cloudformation.GetTemplateSummaryOutput{}

        mockClient.On("GetTemplateSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportStacksToStackSet", func(t *testing.T) {
        input := &cloudformation.ImportStacksToStackSetInput{}
        output := &cloudformation.ImportStacksToStackSetOutput{}

        mockClient.On("ImportStacksToStackSet", ctx, input).Return(output, nil)

        result, err := mockClient.ImportStacksToStackSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChangeSets", func(t *testing.T) {
        input := &cloudformation.ListChangeSetsInput{}
        output := &cloudformation.ListChangeSetsOutput{}

        mockClient.On("ListChangeSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListChangeSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExports", func(t *testing.T) {
        input := &cloudformation.ListExportsInput{}
        output := &cloudformation.ListExportsOutput{}

        mockClient.On("ListExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeneratedTemplates", func(t *testing.T) {
        input := &cloudformation.ListGeneratedTemplatesInput{}
        output := &cloudformation.ListGeneratedTemplatesOutput{}

        mockClient.On("ListGeneratedTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeneratedTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImports", func(t *testing.T) {
        input := &cloudformation.ListImportsInput{}
        output := &cloudformation.ListImportsOutput{}

        mockClient.On("ListImports", ctx, input).Return(output, nil)

        result, err := mockClient.ListImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceScanRelatedResources", func(t *testing.T) {
        input := &cloudformation.ListResourceScanRelatedResourcesInput{}
        output := &cloudformation.ListResourceScanRelatedResourcesOutput{}

        mockClient.On("ListResourceScanRelatedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceScanRelatedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceScanResources", func(t *testing.T) {
        input := &cloudformation.ListResourceScanResourcesInput{}
        output := &cloudformation.ListResourceScanResourcesOutput{}

        mockClient.On("ListResourceScanResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceScanResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceScans", func(t *testing.T) {
        input := &cloudformation.ListResourceScansInput{}
        output := &cloudformation.ListResourceScansOutput{}

        mockClient.On("ListResourceScans", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceScans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackInstanceResourceDrifts", func(t *testing.T) {
        input := &cloudformation.ListStackInstanceResourceDriftsInput{}
        output := &cloudformation.ListStackInstanceResourceDriftsOutput{}

        mockClient.On("ListStackInstanceResourceDrifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackInstanceResourceDrifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackInstances", func(t *testing.T) {
        input := &cloudformation.ListStackInstancesInput{}
        output := &cloudformation.ListStackInstancesOutput{}

        mockClient.On("ListStackInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackResources", func(t *testing.T) {
        input := &cloudformation.ListStackResourcesInput{}
        output := &cloudformation.ListStackResourcesOutput{}

        mockClient.On("ListStackResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackSetAutoDeploymentTargets", func(t *testing.T) {
        input := &cloudformation.ListStackSetAutoDeploymentTargetsInput{}
        output := &cloudformation.ListStackSetAutoDeploymentTargetsOutput{}

        mockClient.On("ListStackSetAutoDeploymentTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackSetAutoDeploymentTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackSetOperationResults", func(t *testing.T) {
        input := &cloudformation.ListStackSetOperationResultsInput{}
        output := &cloudformation.ListStackSetOperationResultsOutput{}

        mockClient.On("ListStackSetOperationResults", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackSetOperationResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackSetOperations", func(t *testing.T) {
        input := &cloudformation.ListStackSetOperationsInput{}
        output := &cloudformation.ListStackSetOperationsOutput{}

        mockClient.On("ListStackSetOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackSetOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackSets", func(t *testing.T) {
        input := &cloudformation.ListStackSetsInput{}
        output := &cloudformation.ListStackSetsOutput{}

        mockClient.On("ListStackSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStacks", func(t *testing.T) {
        input := &cloudformation.ListStacksInput{}
        output := &cloudformation.ListStacksOutput{}

        mockClient.On("ListStacks", ctx, input).Return(output, nil)

        result, err := mockClient.ListStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypeRegistrations", func(t *testing.T) {
        input := &cloudformation.ListTypeRegistrationsInput{}
        output := &cloudformation.ListTypeRegistrationsOutput{}

        mockClient.On("ListTypeRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypeRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypeVersions", func(t *testing.T) {
        input := &cloudformation.ListTypeVersionsInput{}
        output := &cloudformation.ListTypeVersionsOutput{}

        mockClient.On("ListTypeVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypeVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypes", func(t *testing.T) {
        input := &cloudformation.ListTypesInput{}
        output := &cloudformation.ListTypesOutput{}

        mockClient.On("ListTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishType", func(t *testing.T) {
        input := &cloudformation.PublishTypeInput{}
        output := &cloudformation.PublishTypeOutput{}

        mockClient.On("PublishType", ctx, input).Return(output, nil)

        result, err := mockClient.PublishType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecordHandlerProgress", func(t *testing.T) {
        input := &cloudformation.RecordHandlerProgressInput{}
        output := &cloudformation.RecordHandlerProgressOutput{}

        mockClient.On("RecordHandlerProgress", ctx, input).Return(output, nil)

        result, err := mockClient.RecordHandlerProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterPublisher", func(t *testing.T) {
        input := &cloudformation.RegisterPublisherInput{}
        output := &cloudformation.RegisterPublisherOutput{}

        mockClient.On("RegisterPublisher", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterPublisher(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterType", func(t *testing.T) {
        input := &cloudformation.RegisterTypeInput{}
        output := &cloudformation.RegisterTypeOutput{}

        mockClient.On("RegisterType", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRollbackStack", func(t *testing.T) {
        input := &cloudformation.RollbackStackInput{}
        output := &cloudformation.RollbackStackOutput{}

        mockClient.On("RollbackStack", ctx, input).Return(output, nil)

        result, err := mockClient.RollbackStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetStackPolicy", func(t *testing.T) {
        input := &cloudformation.SetStackPolicyInput{}
        output := &cloudformation.SetStackPolicyOutput{}

        mockClient.On("SetStackPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.SetStackPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTypeConfiguration", func(t *testing.T) {
        input := &cloudformation.SetTypeConfigurationInput{}
        output := &cloudformation.SetTypeConfigurationOutput{}

        mockClient.On("SetTypeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.SetTypeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTypeDefaultVersion", func(t *testing.T) {
        input := &cloudformation.SetTypeDefaultVersionInput{}
        output := &cloudformation.SetTypeDefaultVersionOutput{}

        mockClient.On("SetTypeDefaultVersion", ctx, input).Return(output, nil)

        result, err := mockClient.SetTypeDefaultVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignalResource", func(t *testing.T) {
        input := &cloudformation.SignalResourceInput{}
        output := &cloudformation.SignalResourceOutput{}

        mockClient.On("SignalResource", ctx, input).Return(output, nil)

        result, err := mockClient.SignalResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartResourceScan", func(t *testing.T) {
        input := &cloudformation.StartResourceScanInput{}
        output := &cloudformation.StartResourceScanOutput{}

        mockClient.On("StartResourceScan", ctx, input).Return(output, nil)

        result, err := mockClient.StartResourceScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStackSetOperation", func(t *testing.T) {
        input := &cloudformation.StopStackSetOperationInput{}
        output := &cloudformation.StopStackSetOperationOutput{}

        mockClient.On("StopStackSetOperation", ctx, input).Return(output, nil)

        result, err := mockClient.StopStackSetOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestType", func(t *testing.T) {
        input := &cloudformation.TestTypeInput{}
        output := &cloudformation.TestTypeOutput{}

        mockClient.On("TestType", ctx, input).Return(output, nil)

        result, err := mockClient.TestType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGeneratedTemplate", func(t *testing.T) {
        input := &cloudformation.UpdateGeneratedTemplateInput{}
        output := &cloudformation.UpdateGeneratedTemplateOutput{}

        mockClient.On("UpdateGeneratedTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGeneratedTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStack", func(t *testing.T) {
        input := &cloudformation.UpdateStackInput{}
        output := &cloudformation.UpdateStackOutput{}

        mockClient.On("UpdateStack", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStackInstances", func(t *testing.T) {
        input := &cloudformation.UpdateStackInstancesInput{}
        output := &cloudformation.UpdateStackInstancesOutput{}

        mockClient.On("UpdateStackInstances", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStackInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStackSet", func(t *testing.T) {
        input := &cloudformation.UpdateStackSetInput{}
        output := &cloudformation.UpdateStackSetOutput{}

        mockClient.On("UpdateStackSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStackSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTerminationProtection", func(t *testing.T) {
        input := &cloudformation.UpdateTerminationProtectionInput{}
        output := &cloudformation.UpdateTerminationProtectionOutput{}

        mockClient.On("UpdateTerminationProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTerminationProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateTemplate", func(t *testing.T) {
        input := &cloudformation.ValidateTemplateInput{}
        output := &cloudformation.ValidateTemplateOutput{}

        mockClient.On("ValidateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
