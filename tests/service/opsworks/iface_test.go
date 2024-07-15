// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package opsworks_test

// tests for the opsworks service interface for 
// this ../../../service/opsworks/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opsworks/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opsworks/opsworks_iface"
	"github.com/stretchr/testify/assert"
)

func TestOpsworksServiceCanBeMocked(t *testing.T) {
	var iface opsworks_iface.IClient
	iface = &opsworks.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := opsworks.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignInstance", func(t *testing.T) {
        input := &opsworks.AssignInstanceInput{}
        output := &opsworks.AssignInstanceOutput{}

        mockClient.On("AssignInstance", ctx, input).Return(output, nil)

        result, err := mockClient.AssignInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssignVolume", func(t *testing.T) {
        input := &opsworks.AssignVolumeInput{}
        output := &opsworks.AssignVolumeOutput{}

        mockClient.On("AssignVolume", ctx, input).Return(output, nil)

        result, err := mockClient.AssignVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateElasticIp", func(t *testing.T) {
        input := &opsworks.AssociateElasticIpInput{}
        output := &opsworks.AssociateElasticIpOutput{}

        mockClient.On("AssociateElasticIp", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateElasticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachElasticLoadBalancer", func(t *testing.T) {
        input := &opsworks.AttachElasticLoadBalancerInput{}
        output := &opsworks.AttachElasticLoadBalancerOutput{}

        mockClient.On("AttachElasticLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.AttachElasticLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloneStack", func(t *testing.T) {
        input := &opsworks.CloneStackInput{}
        output := &opsworks.CloneStackOutput{}

        mockClient.On("CloneStack", ctx, input).Return(output, nil)

        result, err := mockClient.CloneStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &opsworks.CreateAppInput{}
        output := &opsworks.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &opsworks.CreateDeploymentInput{}
        output := &opsworks.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstance", func(t *testing.T) {
        input := &opsworks.CreateInstanceInput{}
        output := &opsworks.CreateInstanceOutput{}

        mockClient.On("CreateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLayer", func(t *testing.T) {
        input := &opsworks.CreateLayerInput{}
        output := &opsworks.CreateLayerOutput{}

        mockClient.On("CreateLayer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLayer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStack", func(t *testing.T) {
        input := &opsworks.CreateStackInput{}
        output := &opsworks.CreateStackOutput{}

        mockClient.On("CreateStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserProfile", func(t *testing.T) {
        input := &opsworks.CreateUserProfileInput{}
        output := &opsworks.CreateUserProfileOutput{}

        mockClient.On("CreateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &opsworks.DeleteAppInput{}
        output := &opsworks.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstance", func(t *testing.T) {
        input := &opsworks.DeleteInstanceInput{}
        output := &opsworks.DeleteInstanceOutput{}

        mockClient.On("DeleteInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLayer", func(t *testing.T) {
        input := &opsworks.DeleteLayerInput{}
        output := &opsworks.DeleteLayerOutput{}

        mockClient.On("DeleteLayer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLayer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStack", func(t *testing.T) {
        input := &opsworks.DeleteStackInput{}
        output := &opsworks.DeleteStackOutput{}

        mockClient.On("DeleteStack", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserProfile", func(t *testing.T) {
        input := &opsworks.DeleteUserProfileInput{}
        output := &opsworks.DeleteUserProfileOutput{}

        mockClient.On("DeleteUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterEcsCluster", func(t *testing.T) {
        input := &opsworks.DeregisterEcsClusterInput{}
        output := &opsworks.DeregisterEcsClusterOutput{}

        mockClient.On("DeregisterEcsCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterEcsCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterElasticIp", func(t *testing.T) {
        input := &opsworks.DeregisterElasticIpInput{}
        output := &opsworks.DeregisterElasticIpOutput{}

        mockClient.On("DeregisterElasticIp", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterElasticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterInstance", func(t *testing.T) {
        input := &opsworks.DeregisterInstanceInput{}
        output := &opsworks.DeregisterInstanceOutput{}

        mockClient.On("DeregisterInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterRdsDbInstance", func(t *testing.T) {
        input := &opsworks.DeregisterRdsDbInstanceInput{}
        output := &opsworks.DeregisterRdsDbInstanceOutput{}

        mockClient.On("DeregisterRdsDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterRdsDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterVolume", func(t *testing.T) {
        input := &opsworks.DeregisterVolumeInput{}
        output := &opsworks.DeregisterVolumeOutput{}

        mockClient.On("DeregisterVolume", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgentVersions", func(t *testing.T) {
        input := &opsworks.DescribeAgentVersionsInput{}
        output := &opsworks.DescribeAgentVersionsOutput{}

        mockClient.On("DescribeAgentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApps", func(t *testing.T) {
        input := &opsworks.DescribeAppsInput{}
        output := &opsworks.DescribeAppsOutput{}

        mockClient.On("DescribeApps", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCommands", func(t *testing.T) {
        input := &opsworks.DescribeCommandsInput{}
        output := &opsworks.DescribeCommandsOutput{}

        mockClient.On("DescribeCommands", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCommands(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeployments", func(t *testing.T) {
        input := &opsworks.DescribeDeploymentsInput{}
        output := &opsworks.DescribeDeploymentsOutput{}

        mockClient.On("DescribeDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEcsClusters", func(t *testing.T) {
        input := &opsworks.DescribeEcsClustersInput{}
        output := &opsworks.DescribeEcsClustersOutput{}

        mockClient.On("DescribeEcsClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEcsClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticIps", func(t *testing.T) {
        input := &opsworks.DescribeElasticIpsInput{}
        output := &opsworks.DescribeElasticIpsOutput{}

        mockClient.On("DescribeElasticIps", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticIps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticLoadBalancers", func(t *testing.T) {
        input := &opsworks.DescribeElasticLoadBalancersInput{}
        output := &opsworks.DescribeElasticLoadBalancersOutput{}

        mockClient.On("DescribeElasticLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstances", func(t *testing.T) {
        input := &opsworks.DescribeInstancesInput{}
        output := &opsworks.DescribeInstancesOutput{}

        mockClient.On("DescribeInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLayers", func(t *testing.T) {
        input := &opsworks.DescribeLayersInput{}
        output := &opsworks.DescribeLayersOutput{}

        mockClient.On("DescribeLayers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLayers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBasedAutoScaling", func(t *testing.T) {
        input := &opsworks.DescribeLoadBasedAutoScalingInput{}
        output := &opsworks.DescribeLoadBasedAutoScalingOutput{}

        mockClient.On("DescribeLoadBasedAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBasedAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMyUserProfile", func(t *testing.T) {
        input := &opsworks.DescribeMyUserProfileInput{}
        output := &opsworks.DescribeMyUserProfileOutput{}

        mockClient.On("DescribeMyUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMyUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOperatingSystems", func(t *testing.T) {
        input := &opsworks.DescribeOperatingSystemsInput{}
        output := &opsworks.DescribeOperatingSystemsOutput{}

        mockClient.On("DescribeOperatingSystems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOperatingSystems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePermissions", func(t *testing.T) {
        input := &opsworks.DescribePermissionsInput{}
        output := &opsworks.DescribePermissionsOutput{}

        mockClient.On("DescribePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRaidArrays", func(t *testing.T) {
        input := &opsworks.DescribeRaidArraysInput{}
        output := &opsworks.DescribeRaidArraysOutput{}

        mockClient.On("DescribeRaidArrays", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRaidArrays(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRdsDbInstances", func(t *testing.T) {
        input := &opsworks.DescribeRdsDbInstancesInput{}
        output := &opsworks.DescribeRdsDbInstancesOutput{}

        mockClient.On("DescribeRdsDbInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRdsDbInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceErrors", func(t *testing.T) {
        input := &opsworks.DescribeServiceErrorsInput{}
        output := &opsworks.DescribeServiceErrorsOutput{}

        mockClient.On("DescribeServiceErrors", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceErrors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackProvisioningParameters", func(t *testing.T) {
        input := &opsworks.DescribeStackProvisioningParametersInput{}
        output := &opsworks.DescribeStackProvisioningParametersOutput{}

        mockClient.On("DescribeStackProvisioningParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackProvisioningParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStackSummary", func(t *testing.T) {
        input := &opsworks.DescribeStackSummaryInput{}
        output := &opsworks.DescribeStackSummaryOutput{}

        mockClient.On("DescribeStackSummary", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStackSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStacks", func(t *testing.T) {
        input := &opsworks.DescribeStacksInput{}
        output := &opsworks.DescribeStacksOutput{}

        mockClient.On("DescribeStacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTimeBasedAutoScaling", func(t *testing.T) {
        input := &opsworks.DescribeTimeBasedAutoScalingInput{}
        output := &opsworks.DescribeTimeBasedAutoScalingOutput{}

        mockClient.On("DescribeTimeBasedAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTimeBasedAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserProfiles", func(t *testing.T) {
        input := &opsworks.DescribeUserProfilesInput{}
        output := &opsworks.DescribeUserProfilesOutput{}

        mockClient.On("DescribeUserProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVolumes", func(t *testing.T) {
        input := &opsworks.DescribeVolumesInput{}
        output := &opsworks.DescribeVolumesOutput{}

        mockClient.On("DescribeVolumes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVolumes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachElasticLoadBalancer", func(t *testing.T) {
        input := &opsworks.DetachElasticLoadBalancerInput{}
        output := &opsworks.DetachElasticLoadBalancerOutput{}

        mockClient.On("DetachElasticLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DetachElasticLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateElasticIp", func(t *testing.T) {
        input := &opsworks.DisassociateElasticIpInput{}
        output := &opsworks.DisassociateElasticIpOutput{}

        mockClient.On("DisassociateElasticIp", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateElasticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostnameSuggestion", func(t *testing.T) {
        input := &opsworks.GetHostnameSuggestionInput{}
        output := &opsworks.GetHostnameSuggestionOutput{}

        mockClient.On("GetHostnameSuggestion", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostnameSuggestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGrantAccess", func(t *testing.T) {
        input := &opsworks.GrantAccessInput{}
        output := &opsworks.GrantAccessOutput{}

        mockClient.On("GrantAccess", ctx, input).Return(output, nil)

        result, err := mockClient.GrantAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &opsworks.ListTagsInput{}
        output := &opsworks.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootInstance", func(t *testing.T) {
        input := &opsworks.RebootInstanceInput{}
        output := &opsworks.RebootInstanceOutput{}

        mockClient.On("RebootInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterEcsCluster", func(t *testing.T) {
        input := &opsworks.RegisterEcsClusterInput{}
        output := &opsworks.RegisterEcsClusterOutput{}

        mockClient.On("RegisterEcsCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterEcsCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterElasticIp", func(t *testing.T) {
        input := &opsworks.RegisterElasticIpInput{}
        output := &opsworks.RegisterElasticIpOutput{}

        mockClient.On("RegisterElasticIp", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterElasticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterInstance", func(t *testing.T) {
        input := &opsworks.RegisterInstanceInput{}
        output := &opsworks.RegisterInstanceOutput{}

        mockClient.On("RegisterInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterRdsDbInstance", func(t *testing.T) {
        input := &opsworks.RegisterRdsDbInstanceInput{}
        output := &opsworks.RegisterRdsDbInstanceOutput{}

        mockClient.On("RegisterRdsDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterRdsDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterVolume", func(t *testing.T) {
        input := &opsworks.RegisterVolumeInput{}
        output := &opsworks.RegisterVolumeOutput{}

        mockClient.On("RegisterVolume", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLoadBasedAutoScaling", func(t *testing.T) {
        input := &opsworks.SetLoadBasedAutoScalingInput{}
        output := &opsworks.SetLoadBasedAutoScalingOutput{}

        mockClient.On("SetLoadBasedAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.SetLoadBasedAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetPermission", func(t *testing.T) {
        input := &opsworks.SetPermissionInput{}
        output := &opsworks.SetPermissionOutput{}

        mockClient.On("SetPermission", ctx, input).Return(output, nil)

        result, err := mockClient.SetPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTimeBasedAutoScaling", func(t *testing.T) {
        input := &opsworks.SetTimeBasedAutoScalingInput{}
        output := &opsworks.SetTimeBasedAutoScalingOutput{}

        mockClient.On("SetTimeBasedAutoScaling", ctx, input).Return(output, nil)

        result, err := mockClient.SetTimeBasedAutoScaling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstance", func(t *testing.T) {
        input := &opsworks.StartInstanceInput{}
        output := &opsworks.StartInstanceOutput{}

        mockClient.On("StartInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStack", func(t *testing.T) {
        input := &opsworks.StartStackInput{}
        output := &opsworks.StartStackOutput{}

        mockClient.On("StartStack", ctx, input).Return(output, nil)

        result, err := mockClient.StartStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInstance", func(t *testing.T) {
        input := &opsworks.StopInstanceInput{}
        output := &opsworks.StopInstanceOutput{}

        mockClient.On("StopInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StopInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStack", func(t *testing.T) {
        input := &opsworks.StopStackInput{}
        output := &opsworks.StopStackOutput{}

        mockClient.On("StopStack", ctx, input).Return(output, nil)

        result, err := mockClient.StopStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &opsworks.TagResourceInput{}
        output := &opsworks.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnassignInstance", func(t *testing.T) {
        input := &opsworks.UnassignInstanceInput{}
        output := &opsworks.UnassignInstanceOutput{}

        mockClient.On("UnassignInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UnassignInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnassignVolume", func(t *testing.T) {
        input := &opsworks.UnassignVolumeInput{}
        output := &opsworks.UnassignVolumeOutput{}

        mockClient.On("UnassignVolume", ctx, input).Return(output, nil)

        result, err := mockClient.UnassignVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &opsworks.UntagResourceInput{}
        output := &opsworks.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApp", func(t *testing.T) {
        input := &opsworks.UpdateAppInput{}
        output := &opsworks.UpdateAppOutput{}

        mockClient.On("UpdateApp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateElasticIp", func(t *testing.T) {
        input := &opsworks.UpdateElasticIpInput{}
        output := &opsworks.UpdateElasticIpOutput{}

        mockClient.On("UpdateElasticIp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateElasticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstance", func(t *testing.T) {
        input := &opsworks.UpdateInstanceInput{}
        output := &opsworks.UpdateInstanceOutput{}

        mockClient.On("UpdateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLayer", func(t *testing.T) {
        input := &opsworks.UpdateLayerInput{}
        output := &opsworks.UpdateLayerOutput{}

        mockClient.On("UpdateLayer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLayer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMyUserProfile", func(t *testing.T) {
        input := &opsworks.UpdateMyUserProfileInput{}
        output := &opsworks.UpdateMyUserProfileOutput{}

        mockClient.On("UpdateMyUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMyUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRdsDbInstance", func(t *testing.T) {
        input := &opsworks.UpdateRdsDbInstanceInput{}
        output := &opsworks.UpdateRdsDbInstanceOutput{}

        mockClient.On("UpdateRdsDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRdsDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStack", func(t *testing.T) {
        input := &opsworks.UpdateStackInput{}
        output := &opsworks.UpdateStackOutput{}

        mockClient.On("UpdateStack", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserProfile", func(t *testing.T) {
        input := &opsworks.UpdateUserProfileInput{}
        output := &opsworks.UpdateUserProfileOutput{}

        mockClient.On("UpdateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVolume", func(t *testing.T) {
        input := &opsworks.UpdateVolumeInput{}
        output := &opsworks.UpdateVolumeOutput{}

        mockClient.On("UpdateVolume", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVolume(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
