package kendraranking_test

// tests for the kendraranking service interface for this ../../../service/kendraranking/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kendraranking"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kendraranking/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kendraranking/kendraranking_iface"
	"github.com/stretchr/testify/assert"
)

func TestKendrarankingServiceCanBeMocked(t *testing.T) {
	var iface kendraranking_iface.IClient
	iface = &kendraranking.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kendraranking.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRescoreExecutionPlan", func(t *testing.T) {
        input := &kendraranking.CreateRescoreExecutionPlanInput{}
        output := &kendraranking.CreateRescoreExecutionPlanOutput{}

        mockClient.On("CreateRescoreExecutionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRescoreExecutionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRescoreExecutionPlan", func(t *testing.T) {
        input := &kendraranking.DeleteRescoreExecutionPlanInput{}
        output := &kendraranking.DeleteRescoreExecutionPlanOutput{}

        mockClient.On("DeleteRescoreExecutionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRescoreExecutionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRescoreExecutionPlan", func(t *testing.T) {
        input := &kendraranking.DescribeRescoreExecutionPlanInput{}
        output := &kendraranking.DescribeRescoreExecutionPlanOutput{}

        mockClient.On("DescribeRescoreExecutionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRescoreExecutionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRescoreExecutionPlans", func(t *testing.T) {
        input := &kendraranking.ListRescoreExecutionPlansInput{}
        output := &kendraranking.ListRescoreExecutionPlansOutput{}

        mockClient.On("ListRescoreExecutionPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListRescoreExecutionPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kendraranking.ListTagsForResourceInput{}
        output := &kendraranking.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRescore", func(t *testing.T) {
        input := &kendraranking.RescoreInput{}
        output := &kendraranking.RescoreOutput{}

        mockClient.On("Rescore", ctx, input).Return(output, nil)

        result, err := mockClient.Rescore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kendraranking.TagResourceInput{}
        output := &kendraranking.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kendraranking.UntagResourceInput{}
        output := &kendraranking.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRescoreExecutionPlan", func(t *testing.T) {
        input := &kendraranking.UpdateRescoreExecutionPlanInput{}
        output := &kendraranking.UpdateRescoreExecutionPlanOutput{}

        mockClient.On("UpdateRescoreExecutionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRescoreExecutionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
