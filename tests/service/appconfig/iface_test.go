package appconfig_test

// tests for the appconfig service interface for this ../../../service/appconfig/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appconfig/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appconfig/appconfig_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppconfigServiceCanBeMocked(t *testing.T) {
	var iface appconfig_iface.IClient
	iface = &appconfig.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appconfig.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &appconfig.CreateApplicationInput{}
        output := &appconfig.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationProfile", func(t *testing.T) {
        input := &appconfig.CreateConfigurationProfileInput{}
        output := &appconfig.CreateConfigurationProfileOutput{}

        mockClient.On("CreateConfigurationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeploymentStrategy", func(t *testing.T) {
        input := &appconfig.CreateDeploymentStrategyInput{}
        output := &appconfig.CreateDeploymentStrategyOutput{}

        mockClient.On("CreateDeploymentStrategy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeploymentStrategy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &appconfig.CreateEnvironmentInput{}
        output := &appconfig.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExtension", func(t *testing.T) {
        input := &appconfig.CreateExtensionInput{}
        output := &appconfig.CreateExtensionOutput{}

        mockClient.On("CreateExtension", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExtensionAssociation", func(t *testing.T) {
        input := &appconfig.CreateExtensionAssociationInput{}
        output := &appconfig.CreateExtensionAssociationOutput{}

        mockClient.On("CreateExtensionAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExtensionAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHostedConfigurationVersion", func(t *testing.T) {
        input := &appconfig.CreateHostedConfigurationVersionInput{}
        output := &appconfig.CreateHostedConfigurationVersionOutput{}

        mockClient.On("CreateHostedConfigurationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHostedConfigurationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &appconfig.DeleteApplicationInput{}
        output := &appconfig.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationProfile", func(t *testing.T) {
        input := &appconfig.DeleteConfigurationProfileInput{}
        output := &appconfig.DeleteConfigurationProfileOutput{}

        mockClient.On("DeleteConfigurationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeploymentStrategy", func(t *testing.T) {
        input := &appconfig.DeleteDeploymentStrategyInput{}
        output := &appconfig.DeleteDeploymentStrategyOutput{}

        mockClient.On("DeleteDeploymentStrategy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeploymentStrategy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &appconfig.DeleteEnvironmentInput{}
        output := &appconfig.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExtension", func(t *testing.T) {
        input := &appconfig.DeleteExtensionInput{}
        output := &appconfig.DeleteExtensionOutput{}

        mockClient.On("DeleteExtension", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExtensionAssociation", func(t *testing.T) {
        input := &appconfig.DeleteExtensionAssociationInput{}
        output := &appconfig.DeleteExtensionAssociationOutput{}

        mockClient.On("DeleteExtensionAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExtensionAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHostedConfigurationVersion", func(t *testing.T) {
        input := &appconfig.DeleteHostedConfigurationVersionInput{}
        output := &appconfig.DeleteHostedConfigurationVersionOutput{}

        mockClient.On("DeleteHostedConfigurationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHostedConfigurationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &appconfig.GetApplicationInput{}
        output := &appconfig.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguration", func(t *testing.T) {
        input := &appconfig.GetConfigurationInput{}
        output := &appconfig.GetConfigurationOutput{}

        mockClient.On("GetConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationProfile", func(t *testing.T) {
        input := &appconfig.GetConfigurationProfileInput{}
        output := &appconfig.GetConfigurationProfileOutput{}

        mockClient.On("GetConfigurationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &appconfig.GetDeploymentInput{}
        output := &appconfig.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentStrategy", func(t *testing.T) {
        input := &appconfig.GetDeploymentStrategyInput{}
        output := &appconfig.GetDeploymentStrategyOutput{}

        mockClient.On("GetDeploymentStrategy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentStrategy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &appconfig.GetEnvironmentInput{}
        output := &appconfig.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExtension", func(t *testing.T) {
        input := &appconfig.GetExtensionInput{}
        output := &appconfig.GetExtensionOutput{}

        mockClient.On("GetExtension", ctx, input).Return(output, nil)

        result, err := mockClient.GetExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExtensionAssociation", func(t *testing.T) {
        input := &appconfig.GetExtensionAssociationInput{}
        output := &appconfig.GetExtensionAssociationOutput{}

        mockClient.On("GetExtensionAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetExtensionAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostedConfigurationVersion", func(t *testing.T) {
        input := &appconfig.GetHostedConfigurationVersionInput{}
        output := &appconfig.GetHostedConfigurationVersionOutput{}

        mockClient.On("GetHostedConfigurationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostedConfigurationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &appconfig.ListApplicationsInput{}
        output := &appconfig.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationProfiles", func(t *testing.T) {
        input := &appconfig.ListConfigurationProfilesInput{}
        output := &appconfig.ListConfigurationProfilesOutput{}

        mockClient.On("ListConfigurationProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentStrategies", func(t *testing.T) {
        input := &appconfig.ListDeploymentStrategiesInput{}
        output := &appconfig.ListDeploymentStrategiesOutput{}

        mockClient.On("ListDeploymentStrategies", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentStrategies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &appconfig.ListDeploymentsInput{}
        output := &appconfig.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &appconfig.ListEnvironmentsInput{}
        output := &appconfig.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExtensionAssociations", func(t *testing.T) {
        input := &appconfig.ListExtensionAssociationsInput{}
        output := &appconfig.ListExtensionAssociationsOutput{}

        mockClient.On("ListExtensionAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListExtensionAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExtensions", func(t *testing.T) {
        input := &appconfig.ListExtensionsInput{}
        output := &appconfig.ListExtensionsOutput{}

        mockClient.On("ListExtensions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExtensions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHostedConfigurationVersions", func(t *testing.T) {
        input := &appconfig.ListHostedConfigurationVersionsInput{}
        output := &appconfig.ListHostedConfigurationVersionsOutput{}

        mockClient.On("ListHostedConfigurationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListHostedConfigurationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appconfig.ListTagsForResourceInput{}
        output := &appconfig.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeployment", func(t *testing.T) {
        input := &appconfig.StartDeploymentInput{}
        output := &appconfig.StartDeploymentOutput{}

        mockClient.On("StartDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDeployment", func(t *testing.T) {
        input := &appconfig.StopDeploymentInput{}
        output := &appconfig.StopDeploymentOutput{}

        mockClient.On("StopDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StopDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appconfig.TagResourceInput{}
        output := &appconfig.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appconfig.UntagResourceInput{}
        output := &appconfig.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &appconfig.UpdateApplicationInput{}
        output := &appconfig.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationProfile", func(t *testing.T) {
        input := &appconfig.UpdateConfigurationProfileInput{}
        output := &appconfig.UpdateConfigurationProfileOutput{}

        mockClient.On("UpdateConfigurationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeploymentStrategy", func(t *testing.T) {
        input := &appconfig.UpdateDeploymentStrategyInput{}
        output := &appconfig.UpdateDeploymentStrategyOutput{}

        mockClient.On("UpdateDeploymentStrategy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeploymentStrategy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &appconfig.UpdateEnvironmentInput{}
        output := &appconfig.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExtension", func(t *testing.T) {
        input := &appconfig.UpdateExtensionInput{}
        output := &appconfig.UpdateExtensionOutput{}

        mockClient.On("UpdateExtension", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExtensionAssociation", func(t *testing.T) {
        input := &appconfig.UpdateExtensionAssociationInput{}
        output := &appconfig.UpdateExtensionAssociationOutput{}

        mockClient.On("UpdateExtensionAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExtensionAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateConfiguration", func(t *testing.T) {
        input := &appconfig.ValidateConfigurationInput{}
        output := &appconfig.ValidateConfigurationOutput{}

        mockClient.On("ValidateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
