package emr_test

// tests for the emr service interface for this ../../../service/emr/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emr/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emr/emr_iface"
	"github.com/stretchr/testify/assert"
)

func TestEmrServiceCanBeMocked(t *testing.T) {
	var iface emr_iface.IClient
	iface = &emr.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := emr.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddInstanceFleet", func(t *testing.T) {
        input := &emr.AddInstanceFleetInput{}
        output := &emr.AddInstanceFleetOutput{}

        mockClient.On("AddInstanceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.AddInstanceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddInstanceGroups", func(t *testing.T) {
        input := &emr.AddInstanceGroupsInput{}
        output := &emr.AddInstanceGroupsOutput{}

        mockClient.On("AddInstanceGroups", ctx, input).Return(output, nil)

        result, err := mockClient.AddInstanceGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddJobFlowSteps", func(t *testing.T) {
        input := &emr.AddJobFlowStepsInput{}
        output := &emr.AddJobFlowStepsOutput{}

        mockClient.On("AddJobFlowSteps", ctx, input).Return(output, nil)

        result, err := mockClient.AddJobFlowSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &emr.AddTagsInput{}
        output := &emr.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSteps", func(t *testing.T) {
        input := &emr.CancelStepsInput{}
        output := &emr.CancelStepsOutput{}

        mockClient.On("CancelSteps", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityConfiguration", func(t *testing.T) {
        input := &emr.CreateSecurityConfigurationInput{}
        output := &emr.CreateSecurityConfigurationOutput{}

        mockClient.On("CreateSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStudio", func(t *testing.T) {
        input := &emr.CreateStudioInput{}
        output := &emr.CreateStudioOutput{}

        mockClient.On("CreateStudio", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStudioSessionMapping", func(t *testing.T) {
        input := &emr.CreateStudioSessionMappingInput{}
        output := &emr.CreateStudioSessionMappingOutput{}

        mockClient.On("CreateStudioSessionMapping", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStudioSessionMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityConfiguration", func(t *testing.T) {
        input := &emr.DeleteSecurityConfigurationInput{}
        output := &emr.DeleteSecurityConfigurationOutput{}

        mockClient.On("DeleteSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudio", func(t *testing.T) {
        input := &emr.DeleteStudioInput{}
        output := &emr.DeleteStudioOutput{}

        mockClient.On("DeleteStudio", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudioSessionMapping", func(t *testing.T) {
        input := &emr.DeleteStudioSessionMappingInput{}
        output := &emr.DeleteStudioSessionMappingOutput{}

        mockClient.On("DeleteStudioSessionMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudioSessionMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &emr.DescribeClusterInput{}
        output := &emr.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobFlows", func(t *testing.T) {
        input := &emr.DescribeJobFlowsInput{}
        output := &emr.DescribeJobFlowsOutput{}

        mockClient.On("DescribeJobFlows", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotebookExecution", func(t *testing.T) {
        input := &emr.DescribeNotebookExecutionInput{}
        output := &emr.DescribeNotebookExecutionOutput{}

        mockClient.On("DescribeNotebookExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotebookExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReleaseLabel", func(t *testing.T) {
        input := &emr.DescribeReleaseLabelInput{}
        output := &emr.DescribeReleaseLabelOutput{}

        mockClient.On("DescribeReleaseLabel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReleaseLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityConfiguration", func(t *testing.T) {
        input := &emr.DescribeSecurityConfigurationInput{}
        output := &emr.DescribeSecurityConfigurationOutput{}

        mockClient.On("DescribeSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStep", func(t *testing.T) {
        input := &emr.DescribeStepInput{}
        output := &emr.DescribeStepOutput{}

        mockClient.On("DescribeStep", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStudio", func(t *testing.T) {
        input := &emr.DescribeStudioInput{}
        output := &emr.DescribeStudioOutput{}

        mockClient.On("DescribeStudio", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutoTerminationPolicy", func(t *testing.T) {
        input := &emr.GetAutoTerminationPolicyInput{}
        output := &emr.GetAutoTerminationPolicyOutput{}

        mockClient.On("GetAutoTerminationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutoTerminationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlockPublicAccessConfiguration", func(t *testing.T) {
        input := &emr.GetBlockPublicAccessConfigurationInput{}
        output := &emr.GetBlockPublicAccessConfigurationOutput{}

        mockClient.On("GetBlockPublicAccessConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlockPublicAccessConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClusterSessionCredentials", func(t *testing.T) {
        input := &emr.GetClusterSessionCredentialsInput{}
        output := &emr.GetClusterSessionCredentialsOutput{}

        mockClient.On("GetClusterSessionCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetClusterSessionCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedScalingPolicy", func(t *testing.T) {
        input := &emr.GetManagedScalingPolicyInput{}
        output := &emr.GetManagedScalingPolicyOutput{}

        mockClient.On("GetManagedScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStudioSessionMapping", func(t *testing.T) {
        input := &emr.GetStudioSessionMappingInput{}
        output := &emr.GetStudioSessionMappingOutput{}

        mockClient.On("GetStudioSessionMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetStudioSessionMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBootstrapActions", func(t *testing.T) {
        input := &emr.ListBootstrapActionsInput{}
        output := &emr.ListBootstrapActionsOutput{}

        mockClient.On("ListBootstrapActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListBootstrapActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &emr.ListClustersInput{}
        output := &emr.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceFleets", func(t *testing.T) {
        input := &emr.ListInstanceFleetsInput{}
        output := &emr.ListInstanceFleetsOutput{}

        mockClient.On("ListInstanceFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceGroups", func(t *testing.T) {
        input := &emr.ListInstanceGroupsInput{}
        output := &emr.ListInstanceGroupsOutput{}

        mockClient.On("ListInstanceGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstances", func(t *testing.T) {
        input := &emr.ListInstancesInput{}
        output := &emr.ListInstancesOutput{}

        mockClient.On("ListInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotebookExecutions", func(t *testing.T) {
        input := &emr.ListNotebookExecutionsInput{}
        output := &emr.ListNotebookExecutionsOutput{}

        mockClient.On("ListNotebookExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotebookExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReleaseLabels", func(t *testing.T) {
        input := &emr.ListReleaseLabelsInput{}
        output := &emr.ListReleaseLabelsOutput{}

        mockClient.On("ListReleaseLabels", ctx, input).Return(output, nil)

        result, err := mockClient.ListReleaseLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityConfigurations", func(t *testing.T) {
        input := &emr.ListSecurityConfigurationsInput{}
        output := &emr.ListSecurityConfigurationsOutput{}

        mockClient.On("ListSecurityConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSteps", func(t *testing.T) {
        input := &emr.ListStepsInput{}
        output := &emr.ListStepsOutput{}

        mockClient.On("ListSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudioSessionMappings", func(t *testing.T) {
        input := &emr.ListStudioSessionMappingsInput{}
        output := &emr.ListStudioSessionMappingsOutput{}

        mockClient.On("ListStudioSessionMappings", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudioSessionMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudios", func(t *testing.T) {
        input := &emr.ListStudiosInput{}
        output := &emr.ListStudiosOutput{}

        mockClient.On("ListStudios", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudios(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSupportedInstanceTypes", func(t *testing.T) {
        input := &emr.ListSupportedInstanceTypesInput{}
        output := &emr.ListSupportedInstanceTypesOutput{}

        mockClient.On("ListSupportedInstanceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListSupportedInstanceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCluster", func(t *testing.T) {
        input := &emr.ModifyClusterInput{}
        output := &emr.ModifyClusterOutput{}

        mockClient.On("ModifyCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceFleet", func(t *testing.T) {
        input := &emr.ModifyInstanceFleetInput{}
        output := &emr.ModifyInstanceFleetOutput{}

        mockClient.On("ModifyInstanceFleet", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyInstanceGroups", func(t *testing.T) {
        input := &emr.ModifyInstanceGroupsInput{}
        output := &emr.ModifyInstanceGroupsOutput{}

        mockClient.On("ModifyInstanceGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyInstanceGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAutoScalingPolicy", func(t *testing.T) {
        input := &emr.PutAutoScalingPolicyInput{}
        output := &emr.PutAutoScalingPolicyOutput{}

        mockClient.On("PutAutoScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAutoScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAutoTerminationPolicy", func(t *testing.T) {
        input := &emr.PutAutoTerminationPolicyInput{}
        output := &emr.PutAutoTerminationPolicyOutput{}

        mockClient.On("PutAutoTerminationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAutoTerminationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBlockPublicAccessConfiguration", func(t *testing.T) {
        input := &emr.PutBlockPublicAccessConfigurationInput{}
        output := &emr.PutBlockPublicAccessConfigurationOutput{}

        mockClient.On("PutBlockPublicAccessConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBlockPublicAccessConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutManagedScalingPolicy", func(t *testing.T) {
        input := &emr.PutManagedScalingPolicyInput{}
        output := &emr.PutManagedScalingPolicyOutput{}

        mockClient.On("PutManagedScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutManagedScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAutoScalingPolicy", func(t *testing.T) {
        input := &emr.RemoveAutoScalingPolicyInput{}
        output := &emr.RemoveAutoScalingPolicyOutput{}

        mockClient.On("RemoveAutoScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAutoScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAutoTerminationPolicy", func(t *testing.T) {
        input := &emr.RemoveAutoTerminationPolicyInput{}
        output := &emr.RemoveAutoTerminationPolicyOutput{}

        mockClient.On("RemoveAutoTerminationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAutoTerminationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveManagedScalingPolicy", func(t *testing.T) {
        input := &emr.RemoveManagedScalingPolicyInput{}
        output := &emr.RemoveManagedScalingPolicyOutput{}

        mockClient.On("RemoveManagedScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveManagedScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &emr.RemoveTagsInput{}
        output := &emr.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunJobFlow", func(t *testing.T) {
        input := &emr.RunJobFlowInput{}
        output := &emr.RunJobFlowOutput{}

        mockClient.On("RunJobFlow", ctx, input).Return(output, nil)

        result, err := mockClient.RunJobFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetKeepJobFlowAliveWhenNoSteps", func(t *testing.T) {
        input := &emr.SetKeepJobFlowAliveWhenNoStepsInput{}
        output := &emr.SetKeepJobFlowAliveWhenNoStepsOutput{}

        mockClient.On("SetKeepJobFlowAliveWhenNoSteps", ctx, input).Return(output, nil)

        result, err := mockClient.SetKeepJobFlowAliveWhenNoSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTerminationProtection", func(t *testing.T) {
        input := &emr.SetTerminationProtectionInput{}
        output := &emr.SetTerminationProtectionOutput{}

        mockClient.On("SetTerminationProtection", ctx, input).Return(output, nil)

        result, err := mockClient.SetTerminationProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetUnhealthyNodeReplacement", func(t *testing.T) {
        input := &emr.SetUnhealthyNodeReplacementInput{}
        output := &emr.SetUnhealthyNodeReplacementOutput{}

        mockClient.On("SetUnhealthyNodeReplacement", ctx, input).Return(output, nil)

        result, err := mockClient.SetUnhealthyNodeReplacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetVisibleToAllUsers", func(t *testing.T) {
        input := &emr.SetVisibleToAllUsersInput{}
        output := &emr.SetVisibleToAllUsersOutput{}

        mockClient.On("SetVisibleToAllUsers", ctx, input).Return(output, nil)

        result, err := mockClient.SetVisibleToAllUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNotebookExecution", func(t *testing.T) {
        input := &emr.StartNotebookExecutionInput{}
        output := &emr.StartNotebookExecutionOutput{}

        mockClient.On("StartNotebookExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartNotebookExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopNotebookExecution", func(t *testing.T) {
        input := &emr.StopNotebookExecutionInput{}
        output := &emr.StopNotebookExecutionOutput{}

        mockClient.On("StopNotebookExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopNotebookExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateJobFlows", func(t *testing.T) {
        input := &emr.TerminateJobFlowsInput{}
        output := &emr.TerminateJobFlowsOutput{}

        mockClient.On("TerminateJobFlows", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateJobFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStudio", func(t *testing.T) {
        input := &emr.UpdateStudioInput{}
        output := &emr.UpdateStudioOutput{}

        mockClient.On("UpdateStudio", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStudioSessionMapping", func(t *testing.T) {
        input := &emr.UpdateStudioSessionMappingInput{}
        output := &emr.UpdateStudioSessionMappingOutput{}

        mockClient.On("UpdateStudioSessionMapping", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStudioSessionMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
