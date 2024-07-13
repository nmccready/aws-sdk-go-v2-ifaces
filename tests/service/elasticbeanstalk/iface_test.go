package elasticbeanstalk_test

// tests for the elasticbeanstalk service interface for this ../../../service/elasticbeanstalk/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticbeanstalk/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticbeanstalk/elasticbeanstalk_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticbeanstalkServiceCanBeMocked(t *testing.T) {
	var iface elasticbeanstalk_iface.IClient
	iface = &elasticbeanstalk.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticbeanstalk.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortEnvironmentUpdate", func(t *testing.T) {
        input := &elasticbeanstalk.AbortEnvironmentUpdateInput{}
        output := &elasticbeanstalk.AbortEnvironmentUpdateOutput{}

        mockClient.On("AbortEnvironmentUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.AbortEnvironmentUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyEnvironmentManagedAction", func(t *testing.T) {
        input := &elasticbeanstalk.ApplyEnvironmentManagedActionInput{}
        output := &elasticbeanstalk.ApplyEnvironmentManagedActionOutput{}

        mockClient.On("ApplyEnvironmentManagedAction", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyEnvironmentManagedAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEnvironmentOperationsRole", func(t *testing.T) {
        input := &elasticbeanstalk.AssociateEnvironmentOperationsRoleInput{}
        output := &elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput{}

        mockClient.On("AssociateEnvironmentOperationsRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEnvironmentOperationsRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckDNSAvailability", func(t *testing.T) {
        input := &elasticbeanstalk.CheckDNSAvailabilityInput{}
        output := &elasticbeanstalk.CheckDNSAvailabilityOutput{}

        mockClient.On("CheckDNSAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.CheckDNSAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestComposeEnvironments", func(t *testing.T) {
        input := &elasticbeanstalk.ComposeEnvironmentsInput{}
        output := &elasticbeanstalk.ComposeEnvironmentsOutput{}

        mockClient.On("ComposeEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ComposeEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &elasticbeanstalk.CreateApplicationInput{}
        output := &elasticbeanstalk.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationVersion", func(t *testing.T) {
        input := &elasticbeanstalk.CreateApplicationVersionInput{}
        output := &elasticbeanstalk.CreateApplicationVersionOutput{}

        mockClient.On("CreateApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationTemplate", func(t *testing.T) {
        input := &elasticbeanstalk.CreateConfigurationTemplateInput{}
        output := &elasticbeanstalk.CreateConfigurationTemplateOutput{}

        mockClient.On("CreateConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &elasticbeanstalk.CreateEnvironmentInput{}
        output := &elasticbeanstalk.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlatformVersion", func(t *testing.T) {
        input := &elasticbeanstalk.CreatePlatformVersionInput{}
        output := &elasticbeanstalk.CreatePlatformVersionOutput{}

        mockClient.On("CreatePlatformVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlatformVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorageLocation", func(t *testing.T) {
        input := &elasticbeanstalk.CreateStorageLocationInput{}
        output := &elasticbeanstalk.CreateStorageLocationOutput{}

        mockClient.On("CreateStorageLocation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorageLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &elasticbeanstalk.DeleteApplicationInput{}
        output := &elasticbeanstalk.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationVersion", func(t *testing.T) {
        input := &elasticbeanstalk.DeleteApplicationVersionInput{}
        output := &elasticbeanstalk.DeleteApplicationVersionOutput{}

        mockClient.On("DeleteApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationTemplate", func(t *testing.T) {
        input := &elasticbeanstalk.DeleteConfigurationTemplateInput{}
        output := &elasticbeanstalk.DeleteConfigurationTemplateOutput{}

        mockClient.On("DeleteConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentConfiguration", func(t *testing.T) {
        input := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{}
        output := &elasticbeanstalk.DeleteEnvironmentConfigurationOutput{}

        mockClient.On("DeleteEnvironmentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlatformVersion", func(t *testing.T) {
        input := &elasticbeanstalk.DeletePlatformVersionInput{}
        output := &elasticbeanstalk.DeletePlatformVersionOutput{}

        mockClient.On("DeletePlatformVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlatformVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeAccountAttributesInput{}
        output := &elasticbeanstalk.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationVersions", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeApplicationVersionsInput{}
        output := &elasticbeanstalk.DescribeApplicationVersionsOutput{}

        mockClient.On("DescribeApplicationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplications", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeApplicationsInput{}
        output := &elasticbeanstalk.DescribeApplicationsOutput{}

        mockClient.On("DescribeApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationOptions", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeConfigurationOptionsInput{}
        output := &elasticbeanstalk.DescribeConfigurationOptionsOutput{}

        mockClient.On("DescribeConfigurationOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationSettings", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeConfigurationSettingsInput{}
        output := &elasticbeanstalk.DescribeConfigurationSettingsOutput{}

        mockClient.On("DescribeConfigurationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentHealth", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEnvironmentHealthInput{}
        output := &elasticbeanstalk.DescribeEnvironmentHealthOutput{}

        mockClient.On("DescribeEnvironmentHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentManagedActionHistory", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput{}
        output := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput{}

        mockClient.On("DescribeEnvironmentManagedActionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentManagedActionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentManagedActions", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEnvironmentManagedActionsInput{}
        output := &elasticbeanstalk.DescribeEnvironmentManagedActionsOutput{}

        mockClient.On("DescribeEnvironmentManagedActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentManagedActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentResources", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEnvironmentResourcesInput{}
        output := &elasticbeanstalk.DescribeEnvironmentResourcesOutput{}

        mockClient.On("DescribeEnvironmentResources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironments", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEnvironmentsInput{}
        output := &elasticbeanstalk.DescribeEnvironmentsOutput{}

        mockClient.On("DescribeEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeEventsInput{}
        output := &elasticbeanstalk.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstancesHealth", func(t *testing.T) {
        input := &elasticbeanstalk.DescribeInstancesHealthInput{}
        output := &elasticbeanstalk.DescribeInstancesHealthOutput{}

        mockClient.On("DescribeInstancesHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstancesHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePlatformVersion", func(t *testing.T) {
        input := &elasticbeanstalk.DescribePlatformVersionInput{}
        output := &elasticbeanstalk.DescribePlatformVersionOutput{}

        mockClient.On("DescribePlatformVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePlatformVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateEnvironmentOperationsRole", func(t *testing.T) {
        input := &elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput{}
        output := &elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput{}

        mockClient.On("DisassociateEnvironmentOperationsRole", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateEnvironmentOperationsRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableSolutionStacks", func(t *testing.T) {
        input := &elasticbeanstalk.ListAvailableSolutionStacksInput{}
        output := &elasticbeanstalk.ListAvailableSolutionStacksOutput{}

        mockClient.On("ListAvailableSolutionStacks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableSolutionStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlatformBranches", func(t *testing.T) {
        input := &elasticbeanstalk.ListPlatformBranchesInput{}
        output := &elasticbeanstalk.ListPlatformBranchesOutput{}

        mockClient.On("ListPlatformBranches", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlatformBranches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlatformVersions", func(t *testing.T) {
        input := &elasticbeanstalk.ListPlatformVersionsInput{}
        output := &elasticbeanstalk.ListPlatformVersionsOutput{}

        mockClient.On("ListPlatformVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlatformVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &elasticbeanstalk.ListTagsForResourceInput{}
        output := &elasticbeanstalk.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebuildEnvironment", func(t *testing.T) {
        input := &elasticbeanstalk.RebuildEnvironmentInput{}
        output := &elasticbeanstalk.RebuildEnvironmentOutput{}

        mockClient.On("RebuildEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.RebuildEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestEnvironmentInfo", func(t *testing.T) {
        input := &elasticbeanstalk.RequestEnvironmentInfoInput{}
        output := &elasticbeanstalk.RequestEnvironmentInfoOutput{}

        mockClient.On("RequestEnvironmentInfo", ctx, input).Return(output, nil)

        result, err := mockClient.RequestEnvironmentInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestartAppServer", func(t *testing.T) {
        input := &elasticbeanstalk.RestartAppServerInput{}
        output := &elasticbeanstalk.RestartAppServerOutput{}

        mockClient.On("RestartAppServer", ctx, input).Return(output, nil)

        result, err := mockClient.RestartAppServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieveEnvironmentInfo", func(t *testing.T) {
        input := &elasticbeanstalk.RetrieveEnvironmentInfoInput{}
        output := &elasticbeanstalk.RetrieveEnvironmentInfoOutput{}

        mockClient.On("RetrieveEnvironmentInfo", ctx, input).Return(output, nil)

        result, err := mockClient.RetrieveEnvironmentInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSwapEnvironmentCNAMEs", func(t *testing.T) {
        input := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{}
        output := &elasticbeanstalk.SwapEnvironmentCNAMEsOutput{}

        mockClient.On("SwapEnvironmentCNAMEs", ctx, input).Return(output, nil)

        result, err := mockClient.SwapEnvironmentCNAMEs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateEnvironment", func(t *testing.T) {
        input := &elasticbeanstalk.TerminateEnvironmentInput{}
        output := &elasticbeanstalk.TerminateEnvironmentOutput{}

        mockClient.On("TerminateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateApplicationInput{}
        output := &elasticbeanstalk.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationResourceLifecycle", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateApplicationResourceLifecycleInput{}
        output := &elasticbeanstalk.UpdateApplicationResourceLifecycleOutput{}

        mockClient.On("UpdateApplicationResourceLifecycle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationResourceLifecycle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationVersion", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateApplicationVersionInput{}
        output := &elasticbeanstalk.UpdateApplicationVersionOutput{}

        mockClient.On("UpdateApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationTemplate", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateConfigurationTemplateInput{}
        output := &elasticbeanstalk.UpdateConfigurationTemplateOutput{}

        mockClient.On("UpdateConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateEnvironmentInput{}
        output := &elasticbeanstalk.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTagsForResource", func(t *testing.T) {
        input := &elasticbeanstalk.UpdateTagsForResourceInput{}
        output := &elasticbeanstalk.UpdateTagsForResourceOutput{}

        mockClient.On("UpdateTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateConfigurationSettings", func(t *testing.T) {
        input := &elasticbeanstalk.ValidateConfigurationSettingsInput{}
        output := &elasticbeanstalk.ValidateConfigurationSettingsOutput{}

        mockClient.On("ValidateConfigurationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateConfigurationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
