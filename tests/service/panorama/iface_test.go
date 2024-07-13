package panorama_test

// tests for the panorama service interface for this ../../../service/panorama/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/panorama"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/panorama/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/panorama/panorama_iface"
	"github.com/stretchr/testify/assert"
)

func TestPanoramaServiceCanBeMocked(t *testing.T) {
	var iface panorama_iface.IClient
	iface = &panorama.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := panorama.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationInstance", func(t *testing.T) {
        input := &panorama.CreateApplicationInstanceInput{}
        output := &panorama.CreateApplicationInstanceOutput{}

        mockClient.On("CreateApplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJobForDevices", func(t *testing.T) {
        input := &panorama.CreateJobForDevicesInput{}
        output := &panorama.CreateJobForDevicesOutput{}

        mockClient.On("CreateJobForDevices", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJobForDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNodeFromTemplateJob", func(t *testing.T) {
        input := &panorama.CreateNodeFromTemplateJobInput{}
        output := &panorama.CreateNodeFromTemplateJobOutput{}

        mockClient.On("CreateNodeFromTemplateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNodeFromTemplateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackage", func(t *testing.T) {
        input := &panorama.CreatePackageInput{}
        output := &panorama.CreatePackageOutput{}

        mockClient.On("CreatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackageImportJob", func(t *testing.T) {
        input := &panorama.CreatePackageImportJobInput{}
        output := &panorama.CreatePackageImportJobOutput{}

        mockClient.On("CreatePackageImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackageImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevice", func(t *testing.T) {
        input := &panorama.DeleteDeviceInput{}
        output := &panorama.DeleteDeviceOutput{}

        mockClient.On("DeleteDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackage", func(t *testing.T) {
        input := &panorama.DeletePackageInput{}
        output := &panorama.DeletePackageOutput{}

        mockClient.On("DeletePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterPackageVersion", func(t *testing.T) {
        input := &panorama.DeregisterPackageVersionInput{}
        output := &panorama.DeregisterPackageVersionOutput{}

        mockClient.On("DeregisterPackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterPackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationInstance", func(t *testing.T) {
        input := &panorama.DescribeApplicationInstanceInput{}
        output := &panorama.DescribeApplicationInstanceOutput{}

        mockClient.On("DescribeApplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationInstanceDetails", func(t *testing.T) {
        input := &panorama.DescribeApplicationInstanceDetailsInput{}
        output := &panorama.DescribeApplicationInstanceDetailsOutput{}

        mockClient.On("DescribeApplicationInstanceDetails", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationInstanceDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevice", func(t *testing.T) {
        input := &panorama.DescribeDeviceInput{}
        output := &panorama.DescribeDeviceOutput{}

        mockClient.On("DescribeDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeviceJob", func(t *testing.T) {
        input := &panorama.DescribeDeviceJobInput{}
        output := &panorama.DescribeDeviceJobOutput{}

        mockClient.On("DescribeDeviceJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeviceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNode", func(t *testing.T) {
        input := &panorama.DescribeNodeInput{}
        output := &panorama.DescribeNodeOutput{}

        mockClient.On("DescribeNode", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNodeFromTemplateJob", func(t *testing.T) {
        input := &panorama.DescribeNodeFromTemplateJobInput{}
        output := &panorama.DescribeNodeFromTemplateJobOutput{}

        mockClient.On("DescribeNodeFromTemplateJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNodeFromTemplateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackage", func(t *testing.T) {
        input := &panorama.DescribePackageInput{}
        output := &panorama.DescribePackageOutput{}

        mockClient.On("DescribePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackageImportJob", func(t *testing.T) {
        input := &panorama.DescribePackageImportJobInput{}
        output := &panorama.DescribePackageImportJobOutput{}

        mockClient.On("DescribePackageImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackageImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackageVersion", func(t *testing.T) {
        input := &panorama.DescribePackageVersionInput{}
        output := &panorama.DescribePackageVersionOutput{}

        mockClient.On("DescribePackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationInstanceDependencies", func(t *testing.T) {
        input := &panorama.ListApplicationInstanceDependenciesInput{}
        output := &panorama.ListApplicationInstanceDependenciesOutput{}

        mockClient.On("ListApplicationInstanceDependencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationInstanceDependencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationInstanceNodeInstances", func(t *testing.T) {
        input := &panorama.ListApplicationInstanceNodeInstancesInput{}
        output := &panorama.ListApplicationInstanceNodeInstancesOutput{}

        mockClient.On("ListApplicationInstanceNodeInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationInstanceNodeInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationInstances", func(t *testing.T) {
        input := &panorama.ListApplicationInstancesInput{}
        output := &panorama.ListApplicationInstancesOutput{}

        mockClient.On("ListApplicationInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &panorama.ListDevicesInput{}
        output := &panorama.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevicesJobs", func(t *testing.T) {
        input := &panorama.ListDevicesJobsInput{}
        output := &panorama.ListDevicesJobsOutput{}

        mockClient.On("ListDevicesJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevicesJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNodeFromTemplateJobs", func(t *testing.T) {
        input := &panorama.ListNodeFromTemplateJobsInput{}
        output := &panorama.ListNodeFromTemplateJobsOutput{}

        mockClient.On("ListNodeFromTemplateJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListNodeFromTemplateJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNodes", func(t *testing.T) {
        input := &panorama.ListNodesInput{}
        output := &panorama.ListNodesOutput{}

        mockClient.On("ListNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageImportJobs", func(t *testing.T) {
        input := &panorama.ListPackageImportJobsInput{}
        output := &panorama.ListPackageImportJobsOutput{}

        mockClient.On("ListPackageImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackages", func(t *testing.T) {
        input := &panorama.ListPackagesInput{}
        output := &panorama.ListPackagesOutput{}

        mockClient.On("ListPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &panorama.ListTagsForResourceInput{}
        output := &panorama.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionDevice", func(t *testing.T) {
        input := &panorama.ProvisionDeviceInput{}
        output := &panorama.ProvisionDeviceOutput{}

        mockClient.On("ProvisionDevice", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterPackageVersion", func(t *testing.T) {
        input := &panorama.RegisterPackageVersionInput{}
        output := &panorama.RegisterPackageVersionOutput{}

        mockClient.On("RegisterPackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterPackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveApplicationInstance", func(t *testing.T) {
        input := &panorama.RemoveApplicationInstanceInput{}
        output := &panorama.RemoveApplicationInstanceOutput{}

        mockClient.On("RemoveApplicationInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveApplicationInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignalApplicationInstanceNodeInstances", func(t *testing.T) {
        input := &panorama.SignalApplicationInstanceNodeInstancesInput{}
        output := &panorama.SignalApplicationInstanceNodeInstancesOutput{}

        mockClient.On("SignalApplicationInstanceNodeInstances", ctx, input).Return(output, nil)

        result, err := mockClient.SignalApplicationInstanceNodeInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &panorama.TagResourceInput{}
        output := &panorama.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &panorama.UntagResourceInput{}
        output := &panorama.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceMetadata", func(t *testing.T) {
        input := &panorama.UpdateDeviceMetadataInput{}
        output := &panorama.UpdateDeviceMetadataOutput{}

        mockClient.On("UpdateDeviceMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
