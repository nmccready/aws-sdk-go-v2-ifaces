package apprunner_test

// tests for the apprunner service interface for this ../../../service/apprunner/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apprunner/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apprunner/apprunner_iface"
	"github.com/stretchr/testify/assert"
)

func TestApprunnerServiceCanBeMocked(t *testing.T) {
	var iface apprunner_iface.IClient
	iface = &apprunner.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := apprunner.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateCustomDomain", func(t *testing.T) {
        input := &apprunner.AssociateCustomDomainInput{}
        output := &apprunner.AssociateCustomDomainOutput{}

        mockClient.On("AssociateCustomDomain", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateCustomDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutoScalingConfiguration", func(t *testing.T) {
        input := &apprunner.CreateAutoScalingConfigurationInput{}
        output := &apprunner.CreateAutoScalingConfigurationOutput{}

        mockClient.On("CreateAutoScalingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutoScalingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &apprunner.CreateConnectionInput{}
        output := &apprunner.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateObservabilityConfiguration", func(t *testing.T) {
        input := &apprunner.CreateObservabilityConfigurationInput{}
        output := &apprunner.CreateObservabilityConfigurationOutput{}

        mockClient.On("CreateObservabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateObservabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &apprunner.CreateServiceInput{}
        output := &apprunner.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcConnector", func(t *testing.T) {
        input := &apprunner.CreateVpcConnectorInput{}
        output := &apprunner.CreateVpcConnectorOutput{}

        mockClient.On("CreateVpcConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcIngressConnection", func(t *testing.T) {
        input := &apprunner.CreateVpcIngressConnectionInput{}
        output := &apprunner.CreateVpcIngressConnectionOutput{}

        mockClient.On("CreateVpcIngressConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcIngressConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAutoScalingConfiguration", func(t *testing.T) {
        input := &apprunner.DeleteAutoScalingConfigurationInput{}
        output := &apprunner.DeleteAutoScalingConfigurationOutput{}

        mockClient.On("DeleteAutoScalingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAutoScalingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &apprunner.DeleteConnectionInput{}
        output := &apprunner.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObservabilityConfiguration", func(t *testing.T) {
        input := &apprunner.DeleteObservabilityConfigurationInput{}
        output := &apprunner.DeleteObservabilityConfigurationOutput{}

        mockClient.On("DeleteObservabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObservabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &apprunner.DeleteServiceInput{}
        output := &apprunner.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcConnector", func(t *testing.T) {
        input := &apprunner.DeleteVpcConnectorInput{}
        output := &apprunner.DeleteVpcConnectorOutput{}

        mockClient.On("DeleteVpcConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcIngressConnection", func(t *testing.T) {
        input := &apprunner.DeleteVpcIngressConnectionInput{}
        output := &apprunner.DeleteVpcIngressConnectionOutput{}

        mockClient.On("DeleteVpcIngressConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcIngressConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoScalingConfiguration", func(t *testing.T) {
        input := &apprunner.DescribeAutoScalingConfigurationInput{}
        output := &apprunner.DescribeAutoScalingConfigurationOutput{}

        mockClient.On("DescribeAutoScalingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoScalingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomDomains", func(t *testing.T) {
        input := &apprunner.DescribeCustomDomainsInput{}
        output := &apprunner.DescribeCustomDomainsOutput{}

        mockClient.On("DescribeCustomDomains", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeObservabilityConfiguration", func(t *testing.T) {
        input := &apprunner.DescribeObservabilityConfigurationInput{}
        output := &apprunner.DescribeObservabilityConfigurationOutput{}

        mockClient.On("DescribeObservabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeObservabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeService", func(t *testing.T) {
        input := &apprunner.DescribeServiceInput{}
        output := &apprunner.DescribeServiceOutput{}

        mockClient.On("DescribeService", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcConnector", func(t *testing.T) {
        input := &apprunner.DescribeVpcConnectorInput{}
        output := &apprunner.DescribeVpcConnectorOutput{}

        mockClient.On("DescribeVpcConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcIngressConnection", func(t *testing.T) {
        input := &apprunner.DescribeVpcIngressConnectionInput{}
        output := &apprunner.DescribeVpcIngressConnectionOutput{}

        mockClient.On("DescribeVpcIngressConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcIngressConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateCustomDomain", func(t *testing.T) {
        input := &apprunner.DisassociateCustomDomainInput{}
        output := &apprunner.DisassociateCustomDomainOutput{}

        mockClient.On("DisassociateCustomDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateCustomDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutoScalingConfigurations", func(t *testing.T) {
        input := &apprunner.ListAutoScalingConfigurationsInput{}
        output := &apprunner.ListAutoScalingConfigurationsOutput{}

        mockClient.On("ListAutoScalingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutoScalingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnections", func(t *testing.T) {
        input := &apprunner.ListConnectionsInput{}
        output := &apprunner.ListConnectionsOutput{}

        mockClient.On("ListConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObservabilityConfigurations", func(t *testing.T) {
        input := &apprunner.ListObservabilityConfigurationsInput{}
        output := &apprunner.ListObservabilityConfigurationsOutput{}

        mockClient.On("ListObservabilityConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListObservabilityConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOperations", func(t *testing.T) {
        input := &apprunner.ListOperationsInput{}
        output := &apprunner.ListOperationsOutput{}

        mockClient.On("ListOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &apprunner.ListServicesInput{}
        output := &apprunner.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServicesForAutoScalingConfiguration", func(t *testing.T) {
        input := &apprunner.ListServicesForAutoScalingConfigurationInput{}
        output := &apprunner.ListServicesForAutoScalingConfigurationOutput{}

        mockClient.On("ListServicesForAutoScalingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ListServicesForAutoScalingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &apprunner.ListTagsForResourceInput{}
        output := &apprunner.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcConnectors", func(t *testing.T) {
        input := &apprunner.ListVpcConnectorsInput{}
        output := &apprunner.ListVpcConnectorsOutput{}

        mockClient.On("ListVpcConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcIngressConnections", func(t *testing.T) {
        input := &apprunner.ListVpcIngressConnectionsInput{}
        output := &apprunner.ListVpcIngressConnectionsOutput{}

        mockClient.On("ListVpcIngressConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcIngressConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseService", func(t *testing.T) {
        input := &apprunner.PauseServiceInput{}
        output := &apprunner.PauseServiceOutput{}

        mockClient.On("PauseService", ctx, input).Return(output, nil)

        result, err := mockClient.PauseService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeService", func(t *testing.T) {
        input := &apprunner.ResumeServiceInput{}
        output := &apprunner.ResumeServiceOutput{}

        mockClient.On("ResumeService", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeployment", func(t *testing.T) {
        input := &apprunner.StartDeploymentInput{}
        output := &apprunner.StartDeploymentOutput{}

        mockClient.On("StartDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &apprunner.TagResourceInput{}
        output := &apprunner.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &apprunner.UntagResourceInput{}
        output := &apprunner.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDefaultAutoScalingConfiguration", func(t *testing.T) {
        input := &apprunner.UpdateDefaultAutoScalingConfigurationInput{}
        output := &apprunner.UpdateDefaultAutoScalingConfigurationOutput{}

        mockClient.On("UpdateDefaultAutoScalingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDefaultAutoScalingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateService", func(t *testing.T) {
        input := &apprunner.UpdateServiceInput{}
        output := &apprunner.UpdateServiceOutput{}

        mockClient.On("UpdateService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcIngressConnection", func(t *testing.T) {
        input := &apprunner.UpdateVpcIngressConnectionInput{}
        output := &apprunner.UpdateVpcIngressConnectionOutput{}

        mockClient.On("UpdateVpcIngressConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcIngressConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
