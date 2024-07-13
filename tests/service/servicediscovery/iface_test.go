package servicediscovery_test

// tests for the servicediscovery service interface for this ../../../service/servicediscovery/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicediscovery/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicediscovery/servicediscovery_iface"
	"github.com/stretchr/testify/assert"
)

func TestServicediscoveryServiceCanBeMocked(t *testing.T) {
	var iface servicediscovery_iface.IClient
	iface = &servicediscovery.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := servicediscovery.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHttpNamespace", func(t *testing.T) {
        input := &servicediscovery.CreateHttpNamespaceInput{}
        output := &servicediscovery.CreateHttpNamespaceOutput{}

        mockClient.On("CreateHttpNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHttpNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrivateDnsNamespace", func(t *testing.T) {
        input := &servicediscovery.CreatePrivateDnsNamespaceInput{}
        output := &servicediscovery.CreatePrivateDnsNamespaceOutput{}

        mockClient.On("CreatePrivateDnsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrivateDnsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePublicDnsNamespace", func(t *testing.T) {
        input := &servicediscovery.CreatePublicDnsNamespaceInput{}
        output := &servicediscovery.CreatePublicDnsNamespaceOutput{}

        mockClient.On("CreatePublicDnsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePublicDnsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &servicediscovery.CreateServiceInput{}
        output := &servicediscovery.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNamespace", func(t *testing.T) {
        input := &servicediscovery.DeleteNamespaceInput{}
        output := &servicediscovery.DeleteNamespaceOutput{}

        mockClient.On("DeleteNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &servicediscovery.DeleteServiceInput{}
        output := &servicediscovery.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterInstance", func(t *testing.T) {
        input := &servicediscovery.DeregisterInstanceInput{}
        output := &servicediscovery.DeregisterInstanceOutput{}

        mockClient.On("DeregisterInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscoverInstances", func(t *testing.T) {
        input := &servicediscovery.DiscoverInstancesInput{}
        output := &servicediscovery.DiscoverInstancesOutput{}

        mockClient.On("DiscoverInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DiscoverInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscoverInstancesRevision", func(t *testing.T) {
        input := &servicediscovery.DiscoverInstancesRevisionInput{}
        output := &servicediscovery.DiscoverInstancesRevisionOutput{}

        mockClient.On("DiscoverInstancesRevision", ctx, input).Return(output, nil)

        result, err := mockClient.DiscoverInstancesRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstance", func(t *testing.T) {
        input := &servicediscovery.GetInstanceInput{}
        output := &servicediscovery.GetInstanceOutput{}

        mockClient.On("GetInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstancesHealthStatus", func(t *testing.T) {
        input := &servicediscovery.GetInstancesHealthStatusInput{}
        output := &servicediscovery.GetInstancesHealthStatusOutput{}

        mockClient.On("GetInstancesHealthStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstancesHealthStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNamespace", func(t *testing.T) {
        input := &servicediscovery.GetNamespaceInput{}
        output := &servicediscovery.GetNamespaceOutput{}

        mockClient.On("GetNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.GetNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperation", func(t *testing.T) {
        input := &servicediscovery.GetOperationInput{}
        output := &servicediscovery.GetOperationOutput{}

        mockClient.On("GetOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetService", func(t *testing.T) {
        input := &servicediscovery.GetServiceInput{}
        output := &servicediscovery.GetServiceOutput{}

        mockClient.On("GetService", ctx, input).Return(output, nil)

        result, err := mockClient.GetService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstances", func(t *testing.T) {
        input := &servicediscovery.ListInstancesInput{}
        output := &servicediscovery.ListInstancesOutput{}

        mockClient.On("ListInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNamespaces", func(t *testing.T) {
        input := &servicediscovery.ListNamespacesInput{}
        output := &servicediscovery.ListNamespacesOutput{}

        mockClient.On("ListNamespaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListNamespaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOperations", func(t *testing.T) {
        input := &servicediscovery.ListOperationsInput{}
        output := &servicediscovery.ListOperationsOutput{}

        mockClient.On("ListOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &servicediscovery.ListServicesInput{}
        output := &servicediscovery.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &servicediscovery.ListTagsForResourceInput{}
        output := &servicediscovery.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterInstance", func(t *testing.T) {
        input := &servicediscovery.RegisterInstanceInput{}
        output := &servicediscovery.RegisterInstanceOutput{}

        mockClient.On("RegisterInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &servicediscovery.TagResourceInput{}
        output := &servicediscovery.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &servicediscovery.UntagResourceInput{}
        output := &servicediscovery.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHttpNamespace", func(t *testing.T) {
        input := &servicediscovery.UpdateHttpNamespaceInput{}
        output := &servicediscovery.UpdateHttpNamespaceOutput{}

        mockClient.On("UpdateHttpNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHttpNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceCustomHealthStatus", func(t *testing.T) {
        input := &servicediscovery.UpdateInstanceCustomHealthStatusInput{}
        output := &servicediscovery.UpdateInstanceCustomHealthStatusOutput{}

        mockClient.On("UpdateInstanceCustomHealthStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceCustomHealthStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrivateDnsNamespace", func(t *testing.T) {
        input := &servicediscovery.UpdatePrivateDnsNamespaceInput{}
        output := &servicediscovery.UpdatePrivateDnsNamespaceOutput{}

        mockClient.On("UpdatePrivateDnsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrivateDnsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePublicDnsNamespace", func(t *testing.T) {
        input := &servicediscovery.UpdatePublicDnsNamespaceInput{}
        output := &servicediscovery.UpdatePublicDnsNamespaceOutput{}

        mockClient.On("UpdatePublicDnsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePublicDnsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateService", func(t *testing.T) {
        input := &servicediscovery.UpdateServiceInput{}
        output := &servicediscovery.UpdateServiceOutput{}

        mockClient.On("UpdateService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
