package greengrass_test

// tests for the greengrass service interface for this ../../../service/greengrass/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/greengrass/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/greengrass/greengrass_iface"
	"github.com/stretchr/testify/assert"
)

func TestGreengrassServiceCanBeMocked(t *testing.T) {
	var iface greengrass_iface.IClient
	iface = &greengrass.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := greengrass.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateRoleToGroup", func(t *testing.T) {
        input := &greengrass.AssociateRoleToGroupInput{}
        output := &greengrass.AssociateRoleToGroupOutput{}

        mockClient.On("AssociateRoleToGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateRoleToGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateServiceRoleToAccount", func(t *testing.T) {
        input := &greengrass.AssociateServiceRoleToAccountInput{}
        output := &greengrass.AssociateServiceRoleToAccountOutput{}

        mockClient.On("AssociateServiceRoleToAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateServiceRoleToAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectorDefinition", func(t *testing.T) {
        input := &greengrass.CreateConnectorDefinitionInput{}
        output := &greengrass.CreateConnectorDefinitionOutput{}

        mockClient.On("CreateConnectorDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectorDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectorDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateConnectorDefinitionVersionInput{}
        output := &greengrass.CreateConnectorDefinitionVersionOutput{}

        mockClient.On("CreateConnectorDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectorDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCoreDefinition", func(t *testing.T) {
        input := &greengrass.CreateCoreDefinitionInput{}
        output := &greengrass.CreateCoreDefinitionOutput{}

        mockClient.On("CreateCoreDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCoreDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCoreDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateCoreDefinitionVersionInput{}
        output := &greengrass.CreateCoreDefinitionVersionOutput{}

        mockClient.On("CreateCoreDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCoreDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &greengrass.CreateDeploymentInput{}
        output := &greengrass.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeviceDefinition", func(t *testing.T) {
        input := &greengrass.CreateDeviceDefinitionInput{}
        output := &greengrass.CreateDeviceDefinitionOutput{}

        mockClient.On("CreateDeviceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeviceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeviceDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateDeviceDefinitionVersionInput{}
        output := &greengrass.CreateDeviceDefinitionVersionOutput{}

        mockClient.On("CreateDeviceDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeviceDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunctionDefinition", func(t *testing.T) {
        input := &greengrass.CreateFunctionDefinitionInput{}
        output := &greengrass.CreateFunctionDefinitionOutput{}

        mockClient.On("CreateFunctionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunctionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunctionDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateFunctionDefinitionVersionInput{}
        output := &greengrass.CreateFunctionDefinitionVersionOutput{}

        mockClient.On("CreateFunctionDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunctionDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &greengrass.CreateGroupInput{}
        output := &greengrass.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroupCertificateAuthority", func(t *testing.T) {
        input := &greengrass.CreateGroupCertificateAuthorityInput{}
        output := &greengrass.CreateGroupCertificateAuthorityOutput{}

        mockClient.On("CreateGroupCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroupCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroupVersion", func(t *testing.T) {
        input := &greengrass.CreateGroupVersionInput{}
        output := &greengrass.CreateGroupVersionOutput{}

        mockClient.On("CreateGroupVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroupVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoggerDefinition", func(t *testing.T) {
        input := &greengrass.CreateLoggerDefinitionInput{}
        output := &greengrass.CreateLoggerDefinitionOutput{}

        mockClient.On("CreateLoggerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoggerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoggerDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateLoggerDefinitionVersionInput{}
        output := &greengrass.CreateLoggerDefinitionVersionOutput{}

        mockClient.On("CreateLoggerDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoggerDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceDefinition", func(t *testing.T) {
        input := &greengrass.CreateResourceDefinitionInput{}
        output := &greengrass.CreateResourceDefinitionOutput{}

        mockClient.On("CreateResourceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateResourceDefinitionVersionInput{}
        output := &greengrass.CreateResourceDefinitionVersionOutput{}

        mockClient.On("CreateResourceDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSoftwareUpdateJob", func(t *testing.T) {
        input := &greengrass.CreateSoftwareUpdateJobInput{}
        output := &greengrass.CreateSoftwareUpdateJobOutput{}

        mockClient.On("CreateSoftwareUpdateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSoftwareUpdateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriptionDefinition", func(t *testing.T) {
        input := &greengrass.CreateSubscriptionDefinitionInput{}
        output := &greengrass.CreateSubscriptionDefinitionOutput{}

        mockClient.On("CreateSubscriptionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriptionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriptionDefinitionVersion", func(t *testing.T) {
        input := &greengrass.CreateSubscriptionDefinitionVersionInput{}
        output := &greengrass.CreateSubscriptionDefinitionVersionOutput{}

        mockClient.On("CreateSubscriptionDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriptionDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectorDefinition", func(t *testing.T) {
        input := &greengrass.DeleteConnectorDefinitionInput{}
        output := &greengrass.DeleteConnectorDefinitionOutput{}

        mockClient.On("DeleteConnectorDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectorDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoreDefinition", func(t *testing.T) {
        input := &greengrass.DeleteCoreDefinitionInput{}
        output := &greengrass.DeleteCoreDefinitionOutput{}

        mockClient.On("DeleteCoreDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoreDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeviceDefinition", func(t *testing.T) {
        input := &greengrass.DeleteDeviceDefinitionInput{}
        output := &greengrass.DeleteDeviceDefinitionOutput{}

        mockClient.On("DeleteDeviceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeviceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunctionDefinition", func(t *testing.T) {
        input := &greengrass.DeleteFunctionDefinitionInput{}
        output := &greengrass.DeleteFunctionDefinitionOutput{}

        mockClient.On("DeleteFunctionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunctionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &greengrass.DeleteGroupInput{}
        output := &greengrass.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggerDefinition", func(t *testing.T) {
        input := &greengrass.DeleteLoggerDefinitionInput{}
        output := &greengrass.DeleteLoggerDefinitionOutput{}

        mockClient.On("DeleteLoggerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceDefinition", func(t *testing.T) {
        input := &greengrass.DeleteResourceDefinitionInput{}
        output := &greengrass.DeleteResourceDefinitionOutput{}

        mockClient.On("DeleteResourceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriptionDefinition", func(t *testing.T) {
        input := &greengrass.DeleteSubscriptionDefinitionInput{}
        output := &greengrass.DeleteSubscriptionDefinitionOutput{}

        mockClient.On("DeleteSubscriptionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriptionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRoleFromGroup", func(t *testing.T) {
        input := &greengrass.DisassociateRoleFromGroupInput{}
        output := &greengrass.DisassociateRoleFromGroupOutput{}

        mockClient.On("DisassociateRoleFromGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRoleFromGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateServiceRoleFromAccount", func(t *testing.T) {
        input := &greengrass.DisassociateServiceRoleFromAccountInput{}
        output := &greengrass.DisassociateServiceRoleFromAccountOutput{}

        mockClient.On("DisassociateServiceRoleFromAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateServiceRoleFromAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociatedRole", func(t *testing.T) {
        input := &greengrass.GetAssociatedRoleInput{}
        output := &greengrass.GetAssociatedRoleOutput{}

        mockClient.On("GetAssociatedRole", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociatedRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBulkDeploymentStatus", func(t *testing.T) {
        input := &greengrass.GetBulkDeploymentStatusInput{}
        output := &greengrass.GetBulkDeploymentStatusOutput{}

        mockClient.On("GetBulkDeploymentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetBulkDeploymentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectivityInfo", func(t *testing.T) {
        input := &greengrass.GetConnectivityInfoInput{}
        output := &greengrass.GetConnectivityInfoOutput{}

        mockClient.On("GetConnectivityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectivityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectorDefinition", func(t *testing.T) {
        input := &greengrass.GetConnectorDefinitionInput{}
        output := &greengrass.GetConnectorDefinitionOutput{}

        mockClient.On("GetConnectorDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectorDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectorDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetConnectorDefinitionVersionInput{}
        output := &greengrass.GetConnectorDefinitionVersionOutput{}

        mockClient.On("GetConnectorDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectorDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreDefinition", func(t *testing.T) {
        input := &greengrass.GetCoreDefinitionInput{}
        output := &greengrass.GetCoreDefinitionOutput{}

        mockClient.On("GetCoreDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetCoreDefinitionVersionInput{}
        output := &greengrass.GetCoreDefinitionVersionOutput{}

        mockClient.On("GetCoreDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentStatus", func(t *testing.T) {
        input := &greengrass.GetDeploymentStatusInput{}
        output := &greengrass.GetDeploymentStatusOutput{}

        mockClient.On("GetDeploymentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceDefinition", func(t *testing.T) {
        input := &greengrass.GetDeviceDefinitionInput{}
        output := &greengrass.GetDeviceDefinitionOutput{}

        mockClient.On("GetDeviceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetDeviceDefinitionVersionInput{}
        output := &greengrass.GetDeviceDefinitionVersionOutput{}

        mockClient.On("GetDeviceDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionDefinition", func(t *testing.T) {
        input := &greengrass.GetFunctionDefinitionInput{}
        output := &greengrass.GetFunctionDefinitionOutput{}

        mockClient.On("GetFunctionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetFunctionDefinitionVersionInput{}
        output := &greengrass.GetFunctionDefinitionVersionOutput{}

        mockClient.On("GetFunctionDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &greengrass.GetGroupInput{}
        output := &greengrass.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupCertificateAuthority", func(t *testing.T) {
        input := &greengrass.GetGroupCertificateAuthorityInput{}
        output := &greengrass.GetGroupCertificateAuthorityOutput{}

        mockClient.On("GetGroupCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupCertificateConfiguration", func(t *testing.T) {
        input := &greengrass.GetGroupCertificateConfigurationInput{}
        output := &greengrass.GetGroupCertificateConfigurationOutput{}

        mockClient.On("GetGroupCertificateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupCertificateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupVersion", func(t *testing.T) {
        input := &greengrass.GetGroupVersionInput{}
        output := &greengrass.GetGroupVersionOutput{}

        mockClient.On("GetGroupVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggerDefinition", func(t *testing.T) {
        input := &greengrass.GetLoggerDefinitionInput{}
        output := &greengrass.GetLoggerDefinitionOutput{}

        mockClient.On("GetLoggerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggerDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetLoggerDefinitionVersionInput{}
        output := &greengrass.GetLoggerDefinitionVersionOutput{}

        mockClient.On("GetLoggerDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggerDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceDefinition", func(t *testing.T) {
        input := &greengrass.GetResourceDefinitionInput{}
        output := &greengrass.GetResourceDefinitionOutput{}

        mockClient.On("GetResourceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetResourceDefinitionVersionInput{}
        output := &greengrass.GetResourceDefinitionVersionOutput{}

        mockClient.On("GetResourceDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceRoleForAccount", func(t *testing.T) {
        input := &greengrass.GetServiceRoleForAccountInput{}
        output := &greengrass.GetServiceRoleForAccountOutput{}

        mockClient.On("GetServiceRoleForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceRoleForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionDefinition", func(t *testing.T) {
        input := &greengrass.GetSubscriptionDefinitionInput{}
        output := &greengrass.GetSubscriptionDefinitionOutput{}

        mockClient.On("GetSubscriptionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionDefinitionVersion", func(t *testing.T) {
        input := &greengrass.GetSubscriptionDefinitionVersionInput{}
        output := &greengrass.GetSubscriptionDefinitionVersionOutput{}

        mockClient.On("GetSubscriptionDefinitionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionDefinitionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetThingRuntimeConfiguration", func(t *testing.T) {
        input := &greengrass.GetThingRuntimeConfigurationInput{}
        output := &greengrass.GetThingRuntimeConfigurationOutput{}

        mockClient.On("GetThingRuntimeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetThingRuntimeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBulkDeploymentDetailedReports", func(t *testing.T) {
        input := &greengrass.ListBulkDeploymentDetailedReportsInput{}
        output := &greengrass.ListBulkDeploymentDetailedReportsOutput{}

        mockClient.On("ListBulkDeploymentDetailedReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListBulkDeploymentDetailedReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBulkDeployments", func(t *testing.T) {
        input := &greengrass.ListBulkDeploymentsInput{}
        output := &greengrass.ListBulkDeploymentsOutput{}

        mockClient.On("ListBulkDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListBulkDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectorDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListConnectorDefinitionVersionsInput{}
        output := &greengrass.ListConnectorDefinitionVersionsOutput{}

        mockClient.On("ListConnectorDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectorDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectorDefinitions", func(t *testing.T) {
        input := &greengrass.ListConnectorDefinitionsInput{}
        output := &greengrass.ListConnectorDefinitionsOutput{}

        mockClient.On("ListConnectorDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectorDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoreDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListCoreDefinitionVersionsInput{}
        output := &greengrass.ListCoreDefinitionVersionsOutput{}

        mockClient.On("ListCoreDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoreDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoreDefinitions", func(t *testing.T) {
        input := &greengrass.ListCoreDefinitionsInput{}
        output := &greengrass.ListCoreDefinitionsOutput{}

        mockClient.On("ListCoreDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoreDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &greengrass.ListDeploymentsInput{}
        output := &greengrass.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListDeviceDefinitionVersionsInput{}
        output := &greengrass.ListDeviceDefinitionVersionsOutput{}

        mockClient.On("ListDeviceDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceDefinitions", func(t *testing.T) {
        input := &greengrass.ListDeviceDefinitionsInput{}
        output := &greengrass.ListDeviceDefinitionsOutput{}

        mockClient.On("ListDeviceDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctionDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListFunctionDefinitionVersionsInput{}
        output := &greengrass.ListFunctionDefinitionVersionsOutput{}

        mockClient.On("ListFunctionDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctionDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctionDefinitions", func(t *testing.T) {
        input := &greengrass.ListFunctionDefinitionsInput{}
        output := &greengrass.ListFunctionDefinitionsOutput{}

        mockClient.On("ListFunctionDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctionDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupCertificateAuthorities", func(t *testing.T) {
        input := &greengrass.ListGroupCertificateAuthoritiesInput{}
        output := &greengrass.ListGroupCertificateAuthoritiesOutput{}

        mockClient.On("ListGroupCertificateAuthorities", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupCertificateAuthorities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupVersions", func(t *testing.T) {
        input := &greengrass.ListGroupVersionsInput{}
        output := &greengrass.ListGroupVersionsOutput{}

        mockClient.On("ListGroupVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &greengrass.ListGroupsInput{}
        output := &greengrass.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggerDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListLoggerDefinitionVersionsInput{}
        output := &greengrass.ListLoggerDefinitionVersionsOutput{}

        mockClient.On("ListLoggerDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggerDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggerDefinitions", func(t *testing.T) {
        input := &greengrass.ListLoggerDefinitionsInput{}
        output := &greengrass.ListLoggerDefinitionsOutput{}

        mockClient.On("ListLoggerDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggerDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListResourceDefinitionVersionsInput{}
        output := &greengrass.ListResourceDefinitionVersionsOutput{}

        mockClient.On("ListResourceDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceDefinitions", func(t *testing.T) {
        input := &greengrass.ListResourceDefinitionsInput{}
        output := &greengrass.ListResourceDefinitionsOutput{}

        mockClient.On("ListResourceDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionDefinitionVersions", func(t *testing.T) {
        input := &greengrass.ListSubscriptionDefinitionVersionsInput{}
        output := &greengrass.ListSubscriptionDefinitionVersionsOutput{}

        mockClient.On("ListSubscriptionDefinitionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionDefinitionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionDefinitions", func(t *testing.T) {
        input := &greengrass.ListSubscriptionDefinitionsInput{}
        output := &greengrass.ListSubscriptionDefinitionsOutput{}

        mockClient.On("ListSubscriptionDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &greengrass.ListTagsForResourceInput{}
        output := &greengrass.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDeployments", func(t *testing.T) {
        input := &greengrass.ResetDeploymentsInput{}
        output := &greengrass.ResetDeploymentsOutput{}

        mockClient.On("ResetDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBulkDeployment", func(t *testing.T) {
        input := &greengrass.StartBulkDeploymentInput{}
        output := &greengrass.StartBulkDeploymentOutput{}

        mockClient.On("StartBulkDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartBulkDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBulkDeployment", func(t *testing.T) {
        input := &greengrass.StopBulkDeploymentInput{}
        output := &greengrass.StopBulkDeploymentOutput{}

        mockClient.On("StopBulkDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StopBulkDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &greengrass.TagResourceInput{}
        output := &greengrass.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &greengrass.UntagResourceInput{}
        output := &greengrass.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectivityInfo", func(t *testing.T) {
        input := &greengrass.UpdateConnectivityInfoInput{}
        output := &greengrass.UpdateConnectivityInfoOutput{}

        mockClient.On("UpdateConnectivityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectivityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectorDefinition", func(t *testing.T) {
        input := &greengrass.UpdateConnectorDefinitionInput{}
        output := &greengrass.UpdateConnectorDefinitionOutput{}

        mockClient.On("UpdateConnectorDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectorDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCoreDefinition", func(t *testing.T) {
        input := &greengrass.UpdateCoreDefinitionInput{}
        output := &greengrass.UpdateCoreDefinitionOutput{}

        mockClient.On("UpdateCoreDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCoreDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceDefinition", func(t *testing.T) {
        input := &greengrass.UpdateDeviceDefinitionInput{}
        output := &greengrass.UpdateDeviceDefinitionOutput{}

        mockClient.On("UpdateDeviceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunctionDefinition", func(t *testing.T) {
        input := &greengrass.UpdateFunctionDefinitionInput{}
        output := &greengrass.UpdateFunctionDefinitionOutput{}

        mockClient.On("UpdateFunctionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunctionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &greengrass.UpdateGroupInput{}
        output := &greengrass.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroupCertificateConfiguration", func(t *testing.T) {
        input := &greengrass.UpdateGroupCertificateConfigurationInput{}
        output := &greengrass.UpdateGroupCertificateConfigurationOutput{}

        mockClient.On("UpdateGroupCertificateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroupCertificateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoggerDefinition", func(t *testing.T) {
        input := &greengrass.UpdateLoggerDefinitionInput{}
        output := &greengrass.UpdateLoggerDefinitionOutput{}

        mockClient.On("UpdateLoggerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoggerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceDefinition", func(t *testing.T) {
        input := &greengrass.UpdateResourceDefinitionInput{}
        output := &greengrass.UpdateResourceDefinitionOutput{}

        mockClient.On("UpdateResourceDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriptionDefinition", func(t *testing.T) {
        input := &greengrass.UpdateSubscriptionDefinitionInput{}
        output := &greengrass.UpdateSubscriptionDefinitionOutput{}

        mockClient.On("UpdateSubscriptionDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriptionDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThingRuntimeConfiguration", func(t *testing.T) {
        input := &greengrass.UpdateThingRuntimeConfigurationInput{}
        output := &greengrass.UpdateThingRuntimeConfigurationOutput{}

        mockClient.On("UpdateThingRuntimeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThingRuntimeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
