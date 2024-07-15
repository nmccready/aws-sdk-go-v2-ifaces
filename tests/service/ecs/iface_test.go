// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package ecs_test

// tests for the ecs service interface for 
// this ../../../service/ecs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ecs/ecs_iface"
	"github.com/stretchr/testify/assert"
)

func TestEcsServiceCanBeMocked(t *testing.T) {
	var iface ecs_iface.IClient
	iface = &ecs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ecs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCapacityProvider", func(t *testing.T) {
        input := &ecs.CreateCapacityProviderInput{}
        output := &ecs.CreateCapacityProviderOutput{}

        mockClient.On("CreateCapacityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCapacityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &ecs.CreateClusterInput{}
        output := &ecs.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &ecs.CreateServiceInput{}
        output := &ecs.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTaskSet", func(t *testing.T) {
        input := &ecs.CreateTaskSetInput{}
        output := &ecs.CreateTaskSetOutput{}

        mockClient.On("CreateTaskSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTaskSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountSetting", func(t *testing.T) {
        input := &ecs.DeleteAccountSettingInput{}
        output := &ecs.DeleteAccountSettingOutput{}

        mockClient.On("DeleteAccountSetting", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountSetting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttributes", func(t *testing.T) {
        input := &ecs.DeleteAttributesInput{}
        output := &ecs.DeleteAttributesOutput{}

        mockClient.On("DeleteAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCapacityProvider", func(t *testing.T) {
        input := &ecs.DeleteCapacityProviderInput{}
        output := &ecs.DeleteCapacityProviderOutput{}

        mockClient.On("DeleteCapacityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCapacityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &ecs.DeleteClusterInput{}
        output := &ecs.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &ecs.DeleteServiceInput{}
        output := &ecs.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTaskDefinitions", func(t *testing.T) {
        input := &ecs.DeleteTaskDefinitionsInput{}
        output := &ecs.DeleteTaskDefinitionsOutput{}

        mockClient.On("DeleteTaskDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTaskDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTaskSet", func(t *testing.T) {
        input := &ecs.DeleteTaskSetInput{}
        output := &ecs.DeleteTaskSetOutput{}

        mockClient.On("DeleteTaskSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTaskSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterContainerInstance", func(t *testing.T) {
        input := &ecs.DeregisterContainerInstanceInput{}
        output := &ecs.DeregisterContainerInstanceOutput{}

        mockClient.On("DeregisterContainerInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterContainerInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTaskDefinition", func(t *testing.T) {
        input := &ecs.DeregisterTaskDefinitionInput{}
        output := &ecs.DeregisterTaskDefinitionOutput{}

        mockClient.On("DeregisterTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCapacityProviders", func(t *testing.T) {
        input := &ecs.DescribeCapacityProvidersInput{}
        output := &ecs.DescribeCapacityProvidersOutput{}

        mockClient.On("DescribeCapacityProviders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCapacityProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusters", func(t *testing.T) {
        input := &ecs.DescribeClustersInput{}
        output := &ecs.DescribeClustersOutput{}

        mockClient.On("DescribeClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContainerInstances", func(t *testing.T) {
        input := &ecs.DescribeContainerInstancesInput{}
        output := &ecs.DescribeContainerInstancesOutput{}

        mockClient.On("DescribeContainerInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContainerInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServices", func(t *testing.T) {
        input := &ecs.DescribeServicesInput{}
        output := &ecs.DescribeServicesOutput{}

        mockClient.On("DescribeServices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTaskDefinition", func(t *testing.T) {
        input := &ecs.DescribeTaskDefinitionInput{}
        output := &ecs.DescribeTaskDefinitionOutput{}

        mockClient.On("DescribeTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTaskSets", func(t *testing.T) {
        input := &ecs.DescribeTaskSetsInput{}
        output := &ecs.DescribeTaskSetsOutput{}

        mockClient.On("DescribeTaskSets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTaskSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTasks", func(t *testing.T) {
        input := &ecs.DescribeTasksInput{}
        output := &ecs.DescribeTasksOutput{}

        mockClient.On("DescribeTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscoverPollEndpoint", func(t *testing.T) {
        input := &ecs.DiscoverPollEndpointInput{}
        output := &ecs.DiscoverPollEndpointOutput{}

        mockClient.On("DiscoverPollEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DiscoverPollEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteCommand", func(t *testing.T) {
        input := &ecs.ExecuteCommandInput{}
        output := &ecs.ExecuteCommandOutput{}

        mockClient.On("ExecuteCommand", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteCommand(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaskProtection", func(t *testing.T) {
        input := &ecs.GetTaskProtectionInput{}
        output := &ecs.GetTaskProtectionOutput{}

        mockClient.On("GetTaskProtection", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaskProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountSettings", func(t *testing.T) {
        input := &ecs.ListAccountSettingsInput{}
        output := &ecs.ListAccountSettingsOutput{}

        mockClient.On("ListAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttributes", func(t *testing.T) {
        input := &ecs.ListAttributesInput{}
        output := &ecs.ListAttributesOutput{}

        mockClient.On("ListAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &ecs.ListClustersInput{}
        output := &ecs.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContainerInstances", func(t *testing.T) {
        input := &ecs.ListContainerInstancesInput{}
        output := &ecs.ListContainerInstancesOutput{}

        mockClient.On("ListContainerInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListContainerInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &ecs.ListServicesInput{}
        output := &ecs.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServicesByNamespace", func(t *testing.T) {
        input := &ecs.ListServicesByNamespaceInput{}
        output := &ecs.ListServicesByNamespaceOutput{}

        mockClient.On("ListServicesByNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.ListServicesByNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ecs.ListTagsForResourceInput{}
        output := &ecs.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaskDefinitionFamilies", func(t *testing.T) {
        input := &ecs.ListTaskDefinitionFamiliesInput{}
        output := &ecs.ListTaskDefinitionFamiliesOutput{}

        mockClient.On("ListTaskDefinitionFamilies", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaskDefinitionFamilies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaskDefinitions", func(t *testing.T) {
        input := &ecs.ListTaskDefinitionsInput{}
        output := &ecs.ListTaskDefinitionsOutput{}

        mockClient.On("ListTaskDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaskDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTasks", func(t *testing.T) {
        input := &ecs.ListTasksInput{}
        output := &ecs.ListTasksOutput{}

        mockClient.On("ListTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSetting", func(t *testing.T) {
        input := &ecs.PutAccountSettingInput{}
        output := &ecs.PutAccountSettingOutput{}

        mockClient.On("PutAccountSetting", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSetting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSettingDefault", func(t *testing.T) {
        input := &ecs.PutAccountSettingDefaultInput{}
        output := &ecs.PutAccountSettingDefaultOutput{}

        mockClient.On("PutAccountSettingDefault", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSettingDefault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAttributes", func(t *testing.T) {
        input := &ecs.PutAttributesInput{}
        output := &ecs.PutAttributesOutput{}

        mockClient.On("PutAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutClusterCapacityProviders", func(t *testing.T) {
        input := &ecs.PutClusterCapacityProvidersInput{}
        output := &ecs.PutClusterCapacityProvidersOutput{}

        mockClient.On("PutClusterCapacityProviders", ctx, input).Return(output, nil)

        result, err := mockClient.PutClusterCapacityProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterContainerInstance", func(t *testing.T) {
        input := &ecs.RegisterContainerInstanceInput{}
        output := &ecs.RegisterContainerInstanceOutput{}

        mockClient.On("RegisterContainerInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterContainerInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTaskDefinition", func(t *testing.T) {
        input := &ecs.RegisterTaskDefinitionInput{}
        output := &ecs.RegisterTaskDefinitionOutput{}

        mockClient.On("RegisterTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRunTask", func(t *testing.T) {
        input := &ecs.RunTaskInput{}
        output := &ecs.RunTaskOutput{}

        mockClient.On("RunTask", ctx, input).Return(output, nil)

        result, err := mockClient.RunTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTask", func(t *testing.T) {
        input := &ecs.StartTaskInput{}
        output := &ecs.StartTaskOutput{}

        mockClient.On("StartTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTask", func(t *testing.T) {
        input := &ecs.StopTaskInput{}
        output := &ecs.StopTaskOutput{}

        mockClient.On("StopTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitAttachmentStateChanges", func(t *testing.T) {
        input := &ecs.SubmitAttachmentStateChangesInput{}
        output := &ecs.SubmitAttachmentStateChangesOutput{}

        mockClient.On("SubmitAttachmentStateChanges", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitAttachmentStateChanges(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitContainerStateChange", func(t *testing.T) {
        input := &ecs.SubmitContainerStateChangeInput{}
        output := &ecs.SubmitContainerStateChangeOutput{}

        mockClient.On("SubmitContainerStateChange", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitContainerStateChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitTaskStateChange", func(t *testing.T) {
        input := &ecs.SubmitTaskStateChangeInput{}
        output := &ecs.SubmitTaskStateChangeOutput{}

        mockClient.On("SubmitTaskStateChange", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitTaskStateChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ecs.TagResourceInput{}
        output := &ecs.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ecs.UntagResourceInput{}
        output := &ecs.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCapacityProvider", func(t *testing.T) {
        input := &ecs.UpdateCapacityProviderInput{}
        output := &ecs.UpdateCapacityProviderOutput{}

        mockClient.On("UpdateCapacityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCapacityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &ecs.UpdateClusterInput{}
        output := &ecs.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterSettings", func(t *testing.T) {
        input := &ecs.UpdateClusterSettingsInput{}
        output := &ecs.UpdateClusterSettingsOutput{}

        mockClient.On("UpdateClusterSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContainerAgent", func(t *testing.T) {
        input := &ecs.UpdateContainerAgentInput{}
        output := &ecs.UpdateContainerAgentOutput{}

        mockClient.On("UpdateContainerAgent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContainerAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContainerInstancesState", func(t *testing.T) {
        input := &ecs.UpdateContainerInstancesStateInput{}
        output := &ecs.UpdateContainerInstancesStateOutput{}

        mockClient.On("UpdateContainerInstancesState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContainerInstancesState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateService", func(t *testing.T) {
        input := &ecs.UpdateServiceInput{}
        output := &ecs.UpdateServiceOutput{}

        mockClient.On("UpdateService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServicePrimaryTaskSet", func(t *testing.T) {
        input := &ecs.UpdateServicePrimaryTaskSetInput{}
        output := &ecs.UpdateServicePrimaryTaskSetOutput{}

        mockClient.On("UpdateServicePrimaryTaskSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServicePrimaryTaskSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTaskProtection", func(t *testing.T) {
        input := &ecs.UpdateTaskProtectionInput{}
        output := &ecs.UpdateTaskProtectionOutput{}

        mockClient.On("UpdateTaskProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTaskProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTaskSet", func(t *testing.T) {
        input := &ecs.UpdateTaskSetInput{}
        output := &ecs.UpdateTaskSetOutput{}

        mockClient.On("UpdateTaskSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTaskSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}