package cloud9_test

// tests for the cloud9 service interface for this ../../../service/cloud9/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloud9/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloud9/cloud9_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloud9ServiceCanBeMocked(t *testing.T) {
	var iface cloud9_iface.IClient
	iface = &cloud9.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloud9.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentEC2", func(t *testing.T) {
        input := &cloud9.CreateEnvironmentEC2Input{}
        output := &cloud9.CreateEnvironmentEC2Output{}

        mockClient.On("CreateEnvironmentEC2", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentEC2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentMembership", func(t *testing.T) {
        input := &cloud9.CreateEnvironmentMembershipInput{}
        output := &cloud9.CreateEnvironmentMembershipOutput{}

        mockClient.On("CreateEnvironmentMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &cloud9.DeleteEnvironmentInput{}
        output := &cloud9.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentMembership", func(t *testing.T) {
        input := &cloud9.DeleteEnvironmentMembershipInput{}
        output := &cloud9.DeleteEnvironmentMembershipOutput{}

        mockClient.On("DeleteEnvironmentMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentMemberships", func(t *testing.T) {
        input := &cloud9.DescribeEnvironmentMembershipsInput{}
        output := &cloud9.DescribeEnvironmentMembershipsOutput{}

        mockClient.On("DescribeEnvironmentMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironmentStatus", func(t *testing.T) {
        input := &cloud9.DescribeEnvironmentStatusInput{}
        output := &cloud9.DescribeEnvironmentStatusOutput{}

        mockClient.On("DescribeEnvironmentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironmentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEnvironments", func(t *testing.T) {
        input := &cloud9.DescribeEnvironmentsInput{}
        output := &cloud9.DescribeEnvironmentsOutput{}

        mockClient.On("DescribeEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &cloud9.ListEnvironmentsInput{}
        output := &cloud9.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloud9.ListTagsForResourceInput{}
        output := &cloud9.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloud9.TagResourceInput{}
        output := &cloud9.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloud9.UntagResourceInput{}
        output := &cloud9.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &cloud9.UpdateEnvironmentInput{}
        output := &cloud9.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentMembership", func(t *testing.T) {
        input := &cloud9.UpdateEnvironmentMembershipInput{}
        output := &cloud9.UpdateEnvironmentMembershipOutput{}

        mockClient.On("UpdateEnvironmentMembership", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
