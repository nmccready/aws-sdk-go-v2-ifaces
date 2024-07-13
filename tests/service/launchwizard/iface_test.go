package launchwizard_test

// tests for the launchwizard service interface for this ../../../service/launchwizard/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/launchwizard"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/launchwizard/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/launchwizard/launchwizard_iface"
	"github.com/stretchr/testify/assert"
)

func TestLaunchwizardServiceCanBeMocked(t *testing.T) {
	var iface launchwizard_iface.IClient
	iface = &launchwizard.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := launchwizard.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &launchwizard.CreateDeploymentInput{}
        output := &launchwizard.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeployment", func(t *testing.T) {
        input := &launchwizard.DeleteDeploymentInput{}
        output := &launchwizard.DeleteDeploymentOutput{}

        mockClient.On("DeleteDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &launchwizard.GetDeploymentInput{}
        output := &launchwizard.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkload", func(t *testing.T) {
        input := &launchwizard.GetWorkloadInput{}
        output := &launchwizard.GetWorkloadOutput{}

        mockClient.On("GetWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkloadDeploymentPattern", func(t *testing.T) {
        input := &launchwizard.GetWorkloadDeploymentPatternInput{}
        output := &launchwizard.GetWorkloadDeploymentPatternOutput{}

        mockClient.On("GetWorkloadDeploymentPattern", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkloadDeploymentPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentEvents", func(t *testing.T) {
        input := &launchwizard.ListDeploymentEventsInput{}
        output := &launchwizard.ListDeploymentEventsOutput{}

        mockClient.On("ListDeploymentEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &launchwizard.ListDeploymentsInput{}
        output := &launchwizard.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &launchwizard.ListTagsForResourceInput{}
        output := &launchwizard.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloadDeploymentPatterns", func(t *testing.T) {
        input := &launchwizard.ListWorkloadDeploymentPatternsInput{}
        output := &launchwizard.ListWorkloadDeploymentPatternsOutput{}

        mockClient.On("ListWorkloadDeploymentPatterns", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloadDeploymentPatterns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloads", func(t *testing.T) {
        input := &launchwizard.ListWorkloadsInput{}
        output := &launchwizard.ListWorkloadsOutput{}

        mockClient.On("ListWorkloads", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &launchwizard.TagResourceInput{}
        output := &launchwizard.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &launchwizard.UntagResourceInput{}
        output := &launchwizard.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
