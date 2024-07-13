package applicationinsights_test

// tests for the applicationinsights service interface for this ../../../service/applicationinsights/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationinsights/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationinsights/applicationinsights_iface"
	"github.com/stretchr/testify/assert"
)

func TestApplicationinsightsServiceCanBeMocked(t *testing.T) {
	var iface applicationinsights_iface.IClient
	iface = &applicationinsights.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := applicationinsights.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddWorkload", func(t *testing.T) {
        input := &applicationinsights.AddWorkloadInput{}
        output := &applicationinsights.AddWorkloadOutput{}

        mockClient.On("AddWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.AddWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &applicationinsights.CreateApplicationInput{}
        output := &applicationinsights.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponent", func(t *testing.T) {
        input := &applicationinsights.CreateComponentInput{}
        output := &applicationinsights.CreateComponentOutput{}

        mockClient.On("CreateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogPattern", func(t *testing.T) {
        input := &applicationinsights.CreateLogPatternInput{}
        output := &applicationinsights.CreateLogPatternOutput{}

        mockClient.On("CreateLogPattern", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &applicationinsights.DeleteApplicationInput{}
        output := &applicationinsights.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponent", func(t *testing.T) {
        input := &applicationinsights.DeleteComponentInput{}
        output := &applicationinsights.DeleteComponentOutput{}

        mockClient.On("DeleteComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLogPattern", func(t *testing.T) {
        input := &applicationinsights.DeleteLogPatternInput{}
        output := &applicationinsights.DeleteLogPatternOutput{}

        mockClient.On("DeleteLogPattern", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLogPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplication", func(t *testing.T) {
        input := &applicationinsights.DescribeApplicationInput{}
        output := &applicationinsights.DescribeApplicationOutput{}

        mockClient.On("DescribeApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComponent", func(t *testing.T) {
        input := &applicationinsights.DescribeComponentInput{}
        output := &applicationinsights.DescribeComponentOutput{}

        mockClient.On("DescribeComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComponentConfiguration", func(t *testing.T) {
        input := &applicationinsights.DescribeComponentConfigurationInput{}
        output := &applicationinsights.DescribeComponentConfigurationOutput{}

        mockClient.On("DescribeComponentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComponentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComponentConfigurationRecommendation", func(t *testing.T) {
        input := &applicationinsights.DescribeComponentConfigurationRecommendationInput{}
        output := &applicationinsights.DescribeComponentConfigurationRecommendationOutput{}

        mockClient.On("DescribeComponentConfigurationRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComponentConfigurationRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLogPattern", func(t *testing.T) {
        input := &applicationinsights.DescribeLogPatternInput{}
        output := &applicationinsights.DescribeLogPatternOutput{}

        mockClient.On("DescribeLogPattern", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLogPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeObservation", func(t *testing.T) {
        input := &applicationinsights.DescribeObservationInput{}
        output := &applicationinsights.DescribeObservationOutput{}

        mockClient.On("DescribeObservation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeObservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProblem", func(t *testing.T) {
        input := &applicationinsights.DescribeProblemInput{}
        output := &applicationinsights.DescribeProblemOutput{}

        mockClient.On("DescribeProblem", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProblem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProblemObservations", func(t *testing.T) {
        input := &applicationinsights.DescribeProblemObservationsInput{}
        output := &applicationinsights.DescribeProblemObservationsOutput{}

        mockClient.On("DescribeProblemObservations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProblemObservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkload", func(t *testing.T) {
        input := &applicationinsights.DescribeWorkloadInput{}
        output := &applicationinsights.DescribeWorkloadOutput{}

        mockClient.On("DescribeWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &applicationinsights.ListApplicationsInput{}
        output := &applicationinsights.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &applicationinsights.ListComponentsInput{}
        output := &applicationinsights.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationHistory", func(t *testing.T) {
        input := &applicationinsights.ListConfigurationHistoryInput{}
        output := &applicationinsights.ListConfigurationHistoryOutput{}

        mockClient.On("ListConfigurationHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLogPatternSets", func(t *testing.T) {
        input := &applicationinsights.ListLogPatternSetsInput{}
        output := &applicationinsights.ListLogPatternSetsOutput{}

        mockClient.On("ListLogPatternSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListLogPatternSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLogPatterns", func(t *testing.T) {
        input := &applicationinsights.ListLogPatternsInput{}
        output := &applicationinsights.ListLogPatternsOutput{}

        mockClient.On("ListLogPatterns", ctx, input).Return(output, nil)

        result, err := mockClient.ListLogPatterns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProblems", func(t *testing.T) {
        input := &applicationinsights.ListProblemsInput{}
        output := &applicationinsights.ListProblemsOutput{}

        mockClient.On("ListProblems", ctx, input).Return(output, nil)

        result, err := mockClient.ListProblems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &applicationinsights.ListTagsForResourceInput{}
        output := &applicationinsights.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloads", func(t *testing.T) {
        input := &applicationinsights.ListWorkloadsInput{}
        output := &applicationinsights.ListWorkloadsOutput{}

        mockClient.On("ListWorkloads", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveWorkload", func(t *testing.T) {
        input := &applicationinsights.RemoveWorkloadInput{}
        output := &applicationinsights.RemoveWorkloadOutput{}

        mockClient.On("RemoveWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &applicationinsights.TagResourceInput{}
        output := &applicationinsights.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &applicationinsights.UntagResourceInput{}
        output := &applicationinsights.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &applicationinsights.UpdateApplicationInput{}
        output := &applicationinsights.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComponent", func(t *testing.T) {
        input := &applicationinsights.UpdateComponentInput{}
        output := &applicationinsights.UpdateComponentOutput{}

        mockClient.On("UpdateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComponentConfiguration", func(t *testing.T) {
        input := &applicationinsights.UpdateComponentConfigurationInput{}
        output := &applicationinsights.UpdateComponentConfigurationOutput{}

        mockClient.On("UpdateComponentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComponentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLogPattern", func(t *testing.T) {
        input := &applicationinsights.UpdateLogPatternInput{}
        output := &applicationinsights.UpdateLogPatternOutput{}

        mockClient.On("UpdateLogPattern", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLogPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProblem", func(t *testing.T) {
        input := &applicationinsights.UpdateProblemInput{}
        output := &applicationinsights.UpdateProblemOutput{}

        mockClient.On("UpdateProblem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProblem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkload", func(t *testing.T) {
        input := &applicationinsights.UpdateWorkloadInput{}
        output := &applicationinsights.UpdateWorkloadOutput{}

        mockClient.On("UpdateWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
