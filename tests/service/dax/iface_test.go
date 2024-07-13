package dax_test

// tests for the dax service interface for this ../../../service/dax/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dax/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dax/dax_iface"
	"github.com/stretchr/testify/assert"
)

func TestDaxServiceCanBeMocked(t *testing.T) {
	var iface dax_iface.IClient
	iface = &dax.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := dax.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &dax.CreateClusterInput{}
        output := &dax.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParameterGroup", func(t *testing.T) {
        input := &dax.CreateParameterGroupInput{}
        output := &dax.CreateParameterGroupOutput{}

        mockClient.On("CreateParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubnetGroup", func(t *testing.T) {
        input := &dax.CreateSubnetGroupInput{}
        output := &dax.CreateSubnetGroupOutput{}

        mockClient.On("CreateSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecreaseReplicationFactor", func(t *testing.T) {
        input := &dax.DecreaseReplicationFactorInput{}
        output := &dax.DecreaseReplicationFactorOutput{}

        mockClient.On("DecreaseReplicationFactor", ctx, input).Return(output, nil)

        result, err := mockClient.DecreaseReplicationFactor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &dax.DeleteClusterInput{}
        output := &dax.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteParameterGroup", func(t *testing.T) {
        input := &dax.DeleteParameterGroupInput{}
        output := &dax.DeleteParameterGroupOutput{}

        mockClient.On("DeleteParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubnetGroup", func(t *testing.T) {
        input := &dax.DeleteSubnetGroupInput{}
        output := &dax.DeleteSubnetGroupOutput{}

        mockClient.On("DeleteSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusters", func(t *testing.T) {
        input := &dax.DescribeClustersInput{}
        output := &dax.DescribeClustersOutput{}

        mockClient.On("DescribeClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDefaultParameters", func(t *testing.T) {
        input := &dax.DescribeDefaultParametersInput{}
        output := &dax.DescribeDefaultParametersOutput{}

        mockClient.On("DescribeDefaultParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDefaultParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &dax.DescribeEventsInput{}
        output := &dax.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeParameterGroups", func(t *testing.T) {
        input := &dax.DescribeParameterGroupsInput{}
        output := &dax.DescribeParameterGroupsOutput{}

        mockClient.On("DescribeParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeParameters", func(t *testing.T) {
        input := &dax.DescribeParametersInput{}
        output := &dax.DescribeParametersOutput{}

        mockClient.On("DescribeParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubnetGroups", func(t *testing.T) {
        input := &dax.DescribeSubnetGroupsInput{}
        output := &dax.DescribeSubnetGroupsOutput{}

        mockClient.On("DescribeSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIncreaseReplicationFactor", func(t *testing.T) {
        input := &dax.IncreaseReplicationFactorInput{}
        output := &dax.IncreaseReplicationFactorOutput{}

        mockClient.On("IncreaseReplicationFactor", ctx, input).Return(output, nil)

        result, err := mockClient.IncreaseReplicationFactor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &dax.ListTagsInput{}
        output := &dax.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootNode", func(t *testing.T) {
        input := &dax.RebootNodeInput{}
        output := &dax.RebootNodeOutput{}

        mockClient.On("RebootNode", ctx, input).Return(output, nil)

        result, err := mockClient.RebootNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &dax.TagResourceInput{}
        output := &dax.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &dax.UntagResourceInput{}
        output := &dax.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &dax.UpdateClusterInput{}
        output := &dax.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateParameterGroup", func(t *testing.T) {
        input := &dax.UpdateParameterGroupInput{}
        output := &dax.UpdateParameterGroupOutput{}

        mockClient.On("UpdateParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubnetGroup", func(t *testing.T) {
        input := &dax.UpdateSubnetGroupInput{}
        output := &dax.UpdateSubnetGroupOutput{}

        mockClient.On("UpdateSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
