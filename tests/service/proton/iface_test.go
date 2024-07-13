package proton_test

// tests for the proton service interface for this ../../../service/proton/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/proton"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/proton/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/proton/proton_iface"
	"github.com/stretchr/testify/assert"
)

func TestProtonServiceCanBeMocked(t *testing.T) {
	var iface proton_iface.IClient
	iface = &proton.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := proton.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.AcceptEnvironmentAccountConnectionInput{}
        output := &proton.AcceptEnvironmentAccountConnectionOutput{}

        mockClient.On("AcceptEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelComponentDeployment", func(t *testing.T) {
        input := &proton.CancelComponentDeploymentInput{}
        output := &proton.CancelComponentDeploymentOutput{}

        mockClient.On("CancelComponentDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CancelComponentDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelEnvironmentDeployment", func(t *testing.T) {
        input := &proton.CancelEnvironmentDeploymentInput{}
        output := &proton.CancelEnvironmentDeploymentOutput{}

        mockClient.On("CancelEnvironmentDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CancelEnvironmentDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelServiceInstanceDeployment", func(t *testing.T) {
        input := &proton.CancelServiceInstanceDeploymentInput{}
        output := &proton.CancelServiceInstanceDeploymentOutput{}

        mockClient.On("CancelServiceInstanceDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CancelServiceInstanceDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelServicePipelineDeployment", func(t *testing.T) {
        input := &proton.CancelServicePipelineDeploymentInput{}
        output := &proton.CancelServicePipelineDeploymentOutput{}

        mockClient.On("CancelServicePipelineDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CancelServicePipelineDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponent", func(t *testing.T) {
        input := &proton.CreateComponentInput{}
        output := &proton.CreateComponentOutput{}

        mockClient.On("CreateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &proton.CreateEnvironmentInput{}
        output := &proton.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.CreateEnvironmentAccountConnectionInput{}
        output := &proton.CreateEnvironmentAccountConnectionOutput{}

        mockClient.On("CreateEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentTemplate", func(t *testing.T) {
        input := &proton.CreateEnvironmentTemplateInput{}
        output := &proton.CreateEnvironmentTemplateOutput{}

        mockClient.On("CreateEnvironmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentTemplateVersion", func(t *testing.T) {
        input := &proton.CreateEnvironmentTemplateVersionInput{}
        output := &proton.CreateEnvironmentTemplateVersionOutput{}

        mockClient.On("CreateEnvironmentTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepository", func(t *testing.T) {
        input := &proton.CreateRepositoryInput{}
        output := &proton.CreateRepositoryOutput{}

        mockClient.On("CreateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &proton.CreateServiceInput{}
        output := &proton.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceInstance", func(t *testing.T) {
        input := &proton.CreateServiceInstanceInput{}
        output := &proton.CreateServiceInstanceOutput{}

        mockClient.On("CreateServiceInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceSyncConfig", func(t *testing.T) {
        input := &proton.CreateServiceSyncConfigInput{}
        output := &proton.CreateServiceSyncConfigOutput{}

        mockClient.On("CreateServiceSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceTemplate", func(t *testing.T) {
        input := &proton.CreateServiceTemplateInput{}
        output := &proton.CreateServiceTemplateOutput{}

        mockClient.On("CreateServiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceTemplateVersion", func(t *testing.T) {
        input := &proton.CreateServiceTemplateVersionInput{}
        output := &proton.CreateServiceTemplateVersionOutput{}

        mockClient.On("CreateServiceTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplateSyncConfig", func(t *testing.T) {
        input := &proton.CreateTemplateSyncConfigInput{}
        output := &proton.CreateTemplateSyncConfigOutput{}

        mockClient.On("CreateTemplateSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplateSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponent", func(t *testing.T) {
        input := &proton.DeleteComponentInput{}
        output := &proton.DeleteComponentOutput{}

        mockClient.On("DeleteComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeployment", func(t *testing.T) {
        input := &proton.DeleteDeploymentInput{}
        output := &proton.DeleteDeploymentOutput{}

        mockClient.On("DeleteDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &proton.DeleteEnvironmentInput{}
        output := &proton.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.DeleteEnvironmentAccountConnectionInput{}
        output := &proton.DeleteEnvironmentAccountConnectionOutput{}

        mockClient.On("DeleteEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentTemplate", func(t *testing.T) {
        input := &proton.DeleteEnvironmentTemplateInput{}
        output := &proton.DeleteEnvironmentTemplateOutput{}

        mockClient.On("DeleteEnvironmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentTemplateVersion", func(t *testing.T) {
        input := &proton.DeleteEnvironmentTemplateVersionInput{}
        output := &proton.DeleteEnvironmentTemplateVersionOutput{}

        mockClient.On("DeleteEnvironmentTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepository", func(t *testing.T) {
        input := &proton.DeleteRepositoryInput{}
        output := &proton.DeleteRepositoryOutput{}

        mockClient.On("DeleteRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &proton.DeleteServiceInput{}
        output := &proton.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceSyncConfig", func(t *testing.T) {
        input := &proton.DeleteServiceSyncConfigInput{}
        output := &proton.DeleteServiceSyncConfigOutput{}

        mockClient.On("DeleteServiceSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceTemplate", func(t *testing.T) {
        input := &proton.DeleteServiceTemplateInput{}
        output := &proton.DeleteServiceTemplateOutput{}

        mockClient.On("DeleteServiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceTemplateVersion", func(t *testing.T) {
        input := &proton.DeleteServiceTemplateVersionInput{}
        output := &proton.DeleteServiceTemplateVersionOutput{}

        mockClient.On("DeleteServiceTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplateSyncConfig", func(t *testing.T) {
        input := &proton.DeleteTemplateSyncConfigInput{}
        output := &proton.DeleteTemplateSyncConfigOutput{}

        mockClient.On("DeleteTemplateSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplateSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &proton.GetAccountSettingsInput{}
        output := &proton.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponent", func(t *testing.T) {
        input := &proton.GetComponentInput{}
        output := &proton.GetComponentOutput{}

        mockClient.On("GetComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &proton.GetDeploymentInput{}
        output := &proton.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &proton.GetEnvironmentInput{}
        output := &proton.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.GetEnvironmentAccountConnectionInput{}
        output := &proton.GetEnvironmentAccountConnectionOutput{}

        mockClient.On("GetEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentTemplate", func(t *testing.T) {
        input := &proton.GetEnvironmentTemplateInput{}
        output := &proton.GetEnvironmentTemplateOutput{}

        mockClient.On("GetEnvironmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentTemplateVersion", func(t *testing.T) {
        input := &proton.GetEnvironmentTemplateVersionInput{}
        output := &proton.GetEnvironmentTemplateVersionOutput{}

        mockClient.On("GetEnvironmentTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepository", func(t *testing.T) {
        input := &proton.GetRepositoryInput{}
        output := &proton.GetRepositoryOutput{}

        mockClient.On("GetRepository", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositorySyncStatus", func(t *testing.T) {
        input := &proton.GetRepositorySyncStatusInput{}
        output := &proton.GetRepositorySyncStatusOutput{}

        mockClient.On("GetRepositorySyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositorySyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcesSummary", func(t *testing.T) {
        input := &proton.GetResourcesSummaryInput{}
        output := &proton.GetResourcesSummaryOutput{}

        mockClient.On("GetResourcesSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcesSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetService", func(t *testing.T) {
        input := &proton.GetServiceInput{}
        output := &proton.GetServiceOutput{}

        mockClient.On("GetService", ctx, input).Return(output, nil)

        result, err := mockClient.GetService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceInstance", func(t *testing.T) {
        input := &proton.GetServiceInstanceInput{}
        output := &proton.GetServiceInstanceOutput{}

        mockClient.On("GetServiceInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceInstanceSyncStatus", func(t *testing.T) {
        input := &proton.GetServiceInstanceSyncStatusInput{}
        output := &proton.GetServiceInstanceSyncStatusOutput{}

        mockClient.On("GetServiceInstanceSyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceInstanceSyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceSyncBlockerSummary", func(t *testing.T) {
        input := &proton.GetServiceSyncBlockerSummaryInput{}
        output := &proton.GetServiceSyncBlockerSummaryOutput{}

        mockClient.On("GetServiceSyncBlockerSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceSyncBlockerSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceSyncConfig", func(t *testing.T) {
        input := &proton.GetServiceSyncConfigInput{}
        output := &proton.GetServiceSyncConfigOutput{}

        mockClient.On("GetServiceSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceTemplate", func(t *testing.T) {
        input := &proton.GetServiceTemplateInput{}
        output := &proton.GetServiceTemplateOutput{}

        mockClient.On("GetServiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceTemplateVersion", func(t *testing.T) {
        input := &proton.GetServiceTemplateVersionInput{}
        output := &proton.GetServiceTemplateVersionOutput{}

        mockClient.On("GetServiceTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateSyncConfig", func(t *testing.T) {
        input := &proton.GetTemplateSyncConfigInput{}
        output := &proton.GetTemplateSyncConfigOutput{}

        mockClient.On("GetTemplateSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplateSyncStatus", func(t *testing.T) {
        input := &proton.GetTemplateSyncStatusInput{}
        output := &proton.GetTemplateSyncStatusOutput{}

        mockClient.On("GetTemplateSyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplateSyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponentOutputs", func(t *testing.T) {
        input := &proton.ListComponentOutputsInput{}
        output := &proton.ListComponentOutputsOutput{}

        mockClient.On("ListComponentOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponentOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponentProvisionedResources", func(t *testing.T) {
        input := &proton.ListComponentProvisionedResourcesInput{}
        output := &proton.ListComponentProvisionedResourcesOutput{}

        mockClient.On("ListComponentProvisionedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponentProvisionedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &proton.ListComponentsInput{}
        output := &proton.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &proton.ListDeploymentsInput{}
        output := &proton.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentAccountConnections", func(t *testing.T) {
        input := &proton.ListEnvironmentAccountConnectionsInput{}
        output := &proton.ListEnvironmentAccountConnectionsOutput{}

        mockClient.On("ListEnvironmentAccountConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentAccountConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentOutputs", func(t *testing.T) {
        input := &proton.ListEnvironmentOutputsInput{}
        output := &proton.ListEnvironmentOutputsOutput{}

        mockClient.On("ListEnvironmentOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentProvisionedResources", func(t *testing.T) {
        input := &proton.ListEnvironmentProvisionedResourcesInput{}
        output := &proton.ListEnvironmentProvisionedResourcesOutput{}

        mockClient.On("ListEnvironmentProvisionedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentProvisionedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentTemplateVersions", func(t *testing.T) {
        input := &proton.ListEnvironmentTemplateVersionsInput{}
        output := &proton.ListEnvironmentTemplateVersionsOutput{}

        mockClient.On("ListEnvironmentTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentTemplates", func(t *testing.T) {
        input := &proton.ListEnvironmentTemplatesInput{}
        output := &proton.ListEnvironmentTemplatesOutput{}

        mockClient.On("ListEnvironmentTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &proton.ListEnvironmentsInput{}
        output := &proton.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositories", func(t *testing.T) {
        input := &proton.ListRepositoriesInput{}
        output := &proton.ListRepositoriesOutput{}

        mockClient.On("ListRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositorySyncDefinitions", func(t *testing.T) {
        input := &proton.ListRepositorySyncDefinitionsInput{}
        output := &proton.ListRepositorySyncDefinitionsOutput{}

        mockClient.On("ListRepositorySyncDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositorySyncDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceInstanceOutputs", func(t *testing.T) {
        input := &proton.ListServiceInstanceOutputsInput{}
        output := &proton.ListServiceInstanceOutputsOutput{}

        mockClient.On("ListServiceInstanceOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceInstanceOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceInstanceProvisionedResources", func(t *testing.T) {
        input := &proton.ListServiceInstanceProvisionedResourcesInput{}
        output := &proton.ListServiceInstanceProvisionedResourcesOutput{}

        mockClient.On("ListServiceInstanceProvisionedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceInstanceProvisionedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceInstances", func(t *testing.T) {
        input := &proton.ListServiceInstancesInput{}
        output := &proton.ListServiceInstancesOutput{}

        mockClient.On("ListServiceInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServicePipelineOutputs", func(t *testing.T) {
        input := &proton.ListServicePipelineOutputsInput{}
        output := &proton.ListServicePipelineOutputsOutput{}

        mockClient.On("ListServicePipelineOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListServicePipelineOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServicePipelineProvisionedResources", func(t *testing.T) {
        input := &proton.ListServicePipelineProvisionedResourcesInput{}
        output := &proton.ListServicePipelineProvisionedResourcesOutput{}

        mockClient.On("ListServicePipelineProvisionedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListServicePipelineProvisionedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceTemplateVersions", func(t *testing.T) {
        input := &proton.ListServiceTemplateVersionsInput{}
        output := &proton.ListServiceTemplateVersionsOutput{}

        mockClient.On("ListServiceTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceTemplates", func(t *testing.T) {
        input := &proton.ListServiceTemplatesInput{}
        output := &proton.ListServiceTemplatesOutput{}

        mockClient.On("ListServiceTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &proton.ListServicesInput{}
        output := &proton.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &proton.ListTagsForResourceInput{}
        output := &proton.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyResourceDeploymentStatusChange", func(t *testing.T) {
        input := &proton.NotifyResourceDeploymentStatusChangeInput{}
        output := &proton.NotifyResourceDeploymentStatusChangeOutput{}

        mockClient.On("NotifyResourceDeploymentStatusChange", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyResourceDeploymentStatusChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.RejectEnvironmentAccountConnectionInput{}
        output := &proton.RejectEnvironmentAccountConnectionOutput{}

        mockClient.On("RejectEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.RejectEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &proton.TagResourceInput{}
        output := &proton.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &proton.UntagResourceInput{}
        output := &proton.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &proton.UpdateAccountSettingsInput{}
        output := &proton.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComponent", func(t *testing.T) {
        input := &proton.UpdateComponentInput{}
        output := &proton.UpdateComponentOutput{}

        mockClient.On("UpdateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &proton.UpdateEnvironmentInput{}
        output := &proton.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentAccountConnection", func(t *testing.T) {
        input := &proton.UpdateEnvironmentAccountConnectionInput{}
        output := &proton.UpdateEnvironmentAccountConnectionOutput{}

        mockClient.On("UpdateEnvironmentAccountConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentAccountConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentTemplate", func(t *testing.T) {
        input := &proton.UpdateEnvironmentTemplateInput{}
        output := &proton.UpdateEnvironmentTemplateOutput{}

        mockClient.On("UpdateEnvironmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentTemplateVersion", func(t *testing.T) {
        input := &proton.UpdateEnvironmentTemplateVersionInput{}
        output := &proton.UpdateEnvironmentTemplateVersionOutput{}

        mockClient.On("UpdateEnvironmentTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateService", func(t *testing.T) {
        input := &proton.UpdateServiceInput{}
        output := &proton.UpdateServiceOutput{}

        mockClient.On("UpdateService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceInstance", func(t *testing.T) {
        input := &proton.UpdateServiceInstanceInput{}
        output := &proton.UpdateServiceInstanceOutput{}

        mockClient.On("UpdateServiceInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServicePipeline", func(t *testing.T) {
        input := &proton.UpdateServicePipelineInput{}
        output := &proton.UpdateServicePipelineOutput{}

        mockClient.On("UpdateServicePipeline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServicePipeline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSyncBlocker", func(t *testing.T) {
        input := &proton.UpdateServiceSyncBlockerInput{}
        output := &proton.UpdateServiceSyncBlockerOutput{}

        mockClient.On("UpdateServiceSyncBlocker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSyncBlocker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSyncConfig", func(t *testing.T) {
        input := &proton.UpdateServiceSyncConfigInput{}
        output := &proton.UpdateServiceSyncConfigOutput{}

        mockClient.On("UpdateServiceSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceTemplate", func(t *testing.T) {
        input := &proton.UpdateServiceTemplateInput{}
        output := &proton.UpdateServiceTemplateOutput{}

        mockClient.On("UpdateServiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceTemplateVersion", func(t *testing.T) {
        input := &proton.UpdateServiceTemplateVersionInput{}
        output := &proton.UpdateServiceTemplateVersionOutput{}

        mockClient.On("UpdateServiceTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplateSyncConfig", func(t *testing.T) {
        input := &proton.UpdateTemplateSyncConfigInput{}
        output := &proton.UpdateTemplateSyncConfigOutput{}

        mockClient.On("UpdateTemplateSyncConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplateSyncConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
