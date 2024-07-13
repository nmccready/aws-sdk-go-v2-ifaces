package autoscalingplans_test

// tests for the autoscalingplans service interface for this ../../../service/autoscalingplans/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/autoscalingplans/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/autoscalingplans/autoscalingplans_iface"
	"github.com/stretchr/testify/assert"
)

func TestAutoscalingplansServiceCanBeMocked(t *testing.T) {
	var iface autoscalingplans_iface.IClient
	iface = &autoscalingplans.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := autoscalingplans.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScalingPlan", func(t *testing.T) {
        input := &autoscalingplans.CreateScalingPlanInput{}
        output := &autoscalingplans.CreateScalingPlanOutput{}

        mockClient.On("CreateScalingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScalingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScalingPlan", func(t *testing.T) {
        input := &autoscalingplans.DeleteScalingPlanInput{}
        output := &autoscalingplans.DeleteScalingPlanOutput{}

        mockClient.On("DeleteScalingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScalingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingPlanResources", func(t *testing.T) {
        input := &autoscalingplans.DescribeScalingPlanResourcesInput{}
        output := &autoscalingplans.DescribeScalingPlanResourcesOutput{}

        mockClient.On("DescribeScalingPlanResources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingPlanResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingPlans", func(t *testing.T) {
        input := &autoscalingplans.DescribeScalingPlansInput{}
        output := &autoscalingplans.DescribeScalingPlansOutput{}

        mockClient.On("DescribeScalingPlans", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScalingPlanResourceForecastData", func(t *testing.T) {
        input := &autoscalingplans.GetScalingPlanResourceForecastDataInput{}
        output := &autoscalingplans.GetScalingPlanResourceForecastDataOutput{}

        mockClient.On("GetScalingPlanResourceForecastData", ctx, input).Return(output, nil)

        result, err := mockClient.GetScalingPlanResourceForecastData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScalingPlan", func(t *testing.T) {
        input := &autoscalingplans.UpdateScalingPlanInput{}
        output := &autoscalingplans.UpdateScalingPlanOutput{}

        mockClient.On("UpdateScalingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScalingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
