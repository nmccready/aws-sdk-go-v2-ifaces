package applicationautoscaling_test

// tests for the applicationautoscaling service interface for this ../../../service/applicationautoscaling/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationautoscaling/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationautoscaling/applicationautoscaling_iface"
	"github.com/stretchr/testify/assert"
)

func TestApplicationautoscalingServiceCanBeMocked(t *testing.T) {
	var iface applicationautoscaling_iface.IClient
	iface = &applicationautoscaling.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := applicationautoscaling.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScalingPolicy", func(t *testing.T) {
        input := &applicationautoscaling.DeleteScalingPolicyInput{}
        output := &applicationautoscaling.DeleteScalingPolicyOutput{}

        mockClient.On("DeleteScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledAction", func(t *testing.T) {
        input := &applicationautoscaling.DeleteScheduledActionInput{}
        output := &applicationautoscaling.DeleteScheduledActionOutput{}

        mockClient.On("DeleteScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterScalableTarget", func(t *testing.T) {
        input := &applicationautoscaling.DeregisterScalableTargetInput{}
        output := &applicationautoscaling.DeregisterScalableTargetOutput{}

        mockClient.On("DeregisterScalableTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterScalableTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalableTargets", func(t *testing.T) {
        input := &applicationautoscaling.DescribeScalableTargetsInput{}
        output := &applicationautoscaling.DescribeScalableTargetsOutput{}

        mockClient.On("DescribeScalableTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalableTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingActivities", func(t *testing.T) {
        input := &applicationautoscaling.DescribeScalingActivitiesInput{}
        output := &applicationautoscaling.DescribeScalingActivitiesOutput{}

        mockClient.On("DescribeScalingActivities", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingPolicies", func(t *testing.T) {
        input := &applicationautoscaling.DescribeScalingPoliciesInput{}
        output := &applicationautoscaling.DescribeScalingPoliciesOutput{}

        mockClient.On("DescribeScalingPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledActions", func(t *testing.T) {
        input := &applicationautoscaling.DescribeScheduledActionsInput{}
        output := &applicationautoscaling.DescribeScheduledActionsOutput{}

        mockClient.On("DescribeScheduledActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &applicationautoscaling.ListTagsForResourceInput{}
        output := &applicationautoscaling.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutScalingPolicy", func(t *testing.T) {
        input := &applicationautoscaling.PutScalingPolicyInput{}
        output := &applicationautoscaling.PutScalingPolicyOutput{}

        mockClient.On("PutScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutScheduledAction", func(t *testing.T) {
        input := &applicationautoscaling.PutScheduledActionInput{}
        output := &applicationautoscaling.PutScheduledActionOutput{}

        mockClient.On("PutScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.PutScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterScalableTarget", func(t *testing.T) {
        input := &applicationautoscaling.RegisterScalableTargetInput{}
        output := &applicationautoscaling.RegisterScalableTargetOutput{}

        mockClient.On("RegisterScalableTarget", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterScalableTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &applicationautoscaling.TagResourceInput{}
        output := &applicationautoscaling.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &applicationautoscaling.UntagResourceInput{}
        output := &applicationautoscaling.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
