package emrcontainers_test

// tests for the emrcontainers service interface for this ../../../service/emrcontainers/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emrcontainers/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emrcontainers/emrcontainers_iface"
	"github.com/stretchr/testify/assert"
)

func TestEmrcontainersServiceCanBeMocked(t *testing.T) {
	var iface emrcontainers_iface.IClient
	iface = &emrcontainers.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := emrcontainers.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJobRun", func(t *testing.T) {
        input := &emrcontainers.CancelJobRunInput{}
        output := &emrcontainers.CancelJobRunOutput{}

        mockClient.On("CancelJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJobTemplate", func(t *testing.T) {
        input := &emrcontainers.CreateJobTemplateInput{}
        output := &emrcontainers.CreateJobTemplateOutput{}

        mockClient.On("CreateJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateManagedEndpoint", func(t *testing.T) {
        input := &emrcontainers.CreateManagedEndpointInput{}
        output := &emrcontainers.CreateManagedEndpointOutput{}

        mockClient.On("CreateManagedEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateManagedEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityConfiguration", func(t *testing.T) {
        input := &emrcontainers.CreateSecurityConfigurationInput{}
        output := &emrcontainers.CreateSecurityConfigurationOutput{}

        mockClient.On("CreateSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVirtualCluster", func(t *testing.T) {
        input := &emrcontainers.CreateVirtualClusterInput{}
        output := &emrcontainers.CreateVirtualClusterOutput{}

        mockClient.On("CreateVirtualCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVirtualCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobTemplate", func(t *testing.T) {
        input := &emrcontainers.DeleteJobTemplateInput{}
        output := &emrcontainers.DeleteJobTemplateOutput{}

        mockClient.On("DeleteJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteManagedEndpoint", func(t *testing.T) {
        input := &emrcontainers.DeleteManagedEndpointInput{}
        output := &emrcontainers.DeleteManagedEndpointOutput{}

        mockClient.On("DeleteManagedEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteManagedEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVirtualCluster", func(t *testing.T) {
        input := &emrcontainers.DeleteVirtualClusterInput{}
        output := &emrcontainers.DeleteVirtualClusterOutput{}

        mockClient.On("DeleteVirtualCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVirtualCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobRun", func(t *testing.T) {
        input := &emrcontainers.DescribeJobRunInput{}
        output := &emrcontainers.DescribeJobRunOutput{}

        mockClient.On("DescribeJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobTemplate", func(t *testing.T) {
        input := &emrcontainers.DescribeJobTemplateInput{}
        output := &emrcontainers.DescribeJobTemplateOutput{}

        mockClient.On("DescribeJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeManagedEndpoint", func(t *testing.T) {
        input := &emrcontainers.DescribeManagedEndpointInput{}
        output := &emrcontainers.DescribeManagedEndpointOutput{}

        mockClient.On("DescribeManagedEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeManagedEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityConfiguration", func(t *testing.T) {
        input := &emrcontainers.DescribeSecurityConfigurationInput{}
        output := &emrcontainers.DescribeSecurityConfigurationOutput{}

        mockClient.On("DescribeSecurityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVirtualCluster", func(t *testing.T) {
        input := &emrcontainers.DescribeVirtualClusterInput{}
        output := &emrcontainers.DescribeVirtualClusterOutput{}

        mockClient.On("DescribeVirtualCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVirtualCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedEndpointSessionCredentials", func(t *testing.T) {
        input := &emrcontainers.GetManagedEndpointSessionCredentialsInput{}
        output := &emrcontainers.GetManagedEndpointSessionCredentialsOutput{}

        mockClient.On("GetManagedEndpointSessionCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedEndpointSessionCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobRuns", func(t *testing.T) {
        input := &emrcontainers.ListJobRunsInput{}
        output := &emrcontainers.ListJobRunsOutput{}

        mockClient.On("ListJobRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobTemplates", func(t *testing.T) {
        input := &emrcontainers.ListJobTemplatesInput{}
        output := &emrcontainers.ListJobTemplatesOutput{}

        mockClient.On("ListJobTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedEndpoints", func(t *testing.T) {
        input := &emrcontainers.ListManagedEndpointsInput{}
        output := &emrcontainers.ListManagedEndpointsOutput{}

        mockClient.On("ListManagedEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityConfigurations", func(t *testing.T) {
        input := &emrcontainers.ListSecurityConfigurationsInput{}
        output := &emrcontainers.ListSecurityConfigurationsOutput{}

        mockClient.On("ListSecurityConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &emrcontainers.ListTagsForResourceInput{}
        output := &emrcontainers.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualClusters", func(t *testing.T) {
        input := &emrcontainers.ListVirtualClustersInput{}
        output := &emrcontainers.ListVirtualClustersOutput{}

        mockClient.On("ListVirtualClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJobRun", func(t *testing.T) {
        input := &emrcontainers.StartJobRunInput{}
        output := &emrcontainers.StartJobRunOutput{}

        mockClient.On("StartJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &emrcontainers.TagResourceInput{}
        output := &emrcontainers.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &emrcontainers.UntagResourceInput{}
        output := &emrcontainers.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
