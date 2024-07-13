package synthetics_test

// tests for the synthetics service interface for this ../../../service/synthetics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/synthetics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/synthetics/synthetics_iface"
	"github.com/stretchr/testify/assert"
)

func TestSyntheticsServiceCanBeMocked(t *testing.T) {
	var iface synthetics_iface.IClient
	iface = &synthetics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := synthetics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResource", func(t *testing.T) {
        input := &synthetics.AssociateResourceInput{}
        output := &synthetics.AssociateResourceOutput{}

        mockClient.On("AssociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCanary", func(t *testing.T) {
        input := &synthetics.CreateCanaryInput{}
        output := &synthetics.CreateCanaryOutput{}

        mockClient.On("CreateCanary", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &synthetics.CreateGroupInput{}
        output := &synthetics.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCanary", func(t *testing.T) {
        input := &synthetics.DeleteCanaryInput{}
        output := &synthetics.DeleteCanaryOutput{}

        mockClient.On("DeleteCanary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &synthetics.DeleteGroupInput{}
        output := &synthetics.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCanaries", func(t *testing.T) {
        input := &synthetics.DescribeCanariesInput{}
        output := &synthetics.DescribeCanariesOutput{}

        mockClient.On("DescribeCanaries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCanaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCanariesLastRun", func(t *testing.T) {
        input := &synthetics.DescribeCanariesLastRunInput{}
        output := &synthetics.DescribeCanariesLastRunOutput{}

        mockClient.On("DescribeCanariesLastRun", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCanariesLastRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuntimeVersions", func(t *testing.T) {
        input := &synthetics.DescribeRuntimeVersionsInput{}
        output := &synthetics.DescribeRuntimeVersionsOutput{}

        mockClient.On("DescribeRuntimeVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuntimeVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResource", func(t *testing.T) {
        input := &synthetics.DisassociateResourceInput{}
        output := &synthetics.DisassociateResourceOutput{}

        mockClient.On("DisassociateResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCanary", func(t *testing.T) {
        input := &synthetics.GetCanaryInput{}
        output := &synthetics.GetCanaryOutput{}

        mockClient.On("GetCanary", ctx, input).Return(output, nil)

        result, err := mockClient.GetCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCanaryRuns", func(t *testing.T) {
        input := &synthetics.GetCanaryRunsInput{}
        output := &synthetics.GetCanaryRunsOutput{}

        mockClient.On("GetCanaryRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetCanaryRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &synthetics.GetGroupInput{}
        output := &synthetics.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedGroups", func(t *testing.T) {
        input := &synthetics.ListAssociatedGroupsInput{}
        output := &synthetics.ListAssociatedGroupsOutput{}

        mockClient.On("ListAssociatedGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupResources", func(t *testing.T) {
        input := &synthetics.ListGroupResourcesInput{}
        output := &synthetics.ListGroupResourcesOutput{}

        mockClient.On("ListGroupResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &synthetics.ListGroupsInput{}
        output := &synthetics.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &synthetics.ListTagsForResourceInput{}
        output := &synthetics.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCanary", func(t *testing.T) {
        input := &synthetics.StartCanaryInput{}
        output := &synthetics.StartCanaryOutput{}

        mockClient.On("StartCanary", ctx, input).Return(output, nil)

        result, err := mockClient.StartCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCanary", func(t *testing.T) {
        input := &synthetics.StopCanaryInput{}
        output := &synthetics.StopCanaryOutput{}

        mockClient.On("StopCanary", ctx, input).Return(output, nil)

        result, err := mockClient.StopCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &synthetics.TagResourceInput{}
        output := &synthetics.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &synthetics.UntagResourceInput{}
        output := &synthetics.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCanary", func(t *testing.T) {
        input := &synthetics.UpdateCanaryInput{}
        output := &synthetics.UpdateCanaryOutput{}

        mockClient.On("UpdateCanary", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCanary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
