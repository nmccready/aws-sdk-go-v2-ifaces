package autoscaling_test

// tests for the autoscaling service interface for this ../../../service/autoscaling/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/autoscaling/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/autoscaling/autoscaling_iface"
	"github.com/stretchr/testify/assert"
)

func TestAutoscalingServiceCanBeMocked(t *testing.T) {
	var iface autoscaling_iface.IClient
	iface = &autoscaling.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := autoscaling.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachInstances", func(t *testing.T) {
        input := &autoscaling.AttachInstancesInput{}
        output := &autoscaling.AttachInstancesOutput{}

        mockClient.On("AttachInstances", ctx, input).Return(output, nil)

        result, err := mockClient.AttachInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachLoadBalancerTargetGroups", func(t *testing.T) {
        input := &autoscaling.AttachLoadBalancerTargetGroupsInput{}
        output := &autoscaling.AttachLoadBalancerTargetGroupsOutput{}

        mockClient.On("AttachLoadBalancerTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.AttachLoadBalancerTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachLoadBalancers", func(t *testing.T) {
        input := &autoscaling.AttachLoadBalancersInput{}
        output := &autoscaling.AttachLoadBalancersOutput{}

        mockClient.On("AttachLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.AttachLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachTrafficSources", func(t *testing.T) {
        input := &autoscaling.AttachTrafficSourcesInput{}
        output := &autoscaling.AttachTrafficSourcesOutput{}

        mockClient.On("AttachTrafficSources", ctx, input).Return(output, nil)

        result, err := mockClient.AttachTrafficSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteScheduledAction", func(t *testing.T) {
        input := &autoscaling.BatchDeleteScheduledActionInput{}
        output := &autoscaling.BatchDeleteScheduledActionOutput{}

        mockClient.On("BatchDeleteScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutScheduledUpdateGroupAction", func(t *testing.T) {
        input := &autoscaling.BatchPutScheduledUpdateGroupActionInput{}
        output := &autoscaling.BatchPutScheduledUpdateGroupActionOutput{}

        mockClient.On("BatchPutScheduledUpdateGroupAction", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutScheduledUpdateGroupAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelInstanceRefresh", func(t *testing.T) {
        input := &autoscaling.CancelInstanceRefreshInput{}
        output := &autoscaling.CancelInstanceRefreshOutput{}

        mockClient.On("CancelInstanceRefresh", ctx, input).Return(output, nil)

        result, err := mockClient.CancelInstanceRefresh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteLifecycleAction", func(t *testing.T) {
        input := &autoscaling.CompleteLifecycleActionInput{}
        output := &autoscaling.CompleteLifecycleActionOutput{}

        mockClient.On("CompleteLifecycleAction", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteLifecycleAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutoScalingGroup", func(t *testing.T) {
        input := &autoscaling.CreateAutoScalingGroupInput{}
        output := &autoscaling.CreateAutoScalingGroupOutput{}

        mockClient.On("CreateAutoScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutoScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchConfiguration", func(t *testing.T) {
        input := &autoscaling.CreateLaunchConfigurationInput{}
        output := &autoscaling.CreateLaunchConfigurationOutput{}

        mockClient.On("CreateLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOrUpdateTags", func(t *testing.T) {
        input := &autoscaling.CreateOrUpdateTagsInput{}
        output := &autoscaling.CreateOrUpdateTagsOutput{}

        mockClient.On("CreateOrUpdateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOrUpdateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAutoScalingGroup", func(t *testing.T) {
        input := &autoscaling.DeleteAutoScalingGroupInput{}
        output := &autoscaling.DeleteAutoScalingGroupOutput{}

        mockClient.On("DeleteAutoScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAutoScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchConfiguration", func(t *testing.T) {
        input := &autoscaling.DeleteLaunchConfigurationInput{}
        output := &autoscaling.DeleteLaunchConfigurationOutput{}

        mockClient.On("DeleteLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecycleHook", func(t *testing.T) {
        input := &autoscaling.DeleteLifecycleHookInput{}
        output := &autoscaling.DeleteLifecycleHookOutput{}

        mockClient.On("DeleteLifecycleHook", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecycleHook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotificationConfiguration", func(t *testing.T) {
        input := &autoscaling.DeleteNotificationConfigurationInput{}
        output := &autoscaling.DeleteNotificationConfigurationOutput{}

        mockClient.On("DeleteNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &autoscaling.DeletePolicyInput{}
        output := &autoscaling.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledAction", func(t *testing.T) {
        input := &autoscaling.DeleteScheduledActionInput{}
        output := &autoscaling.DeleteScheduledActionOutput{}

        mockClient.On("DeleteScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &autoscaling.DeleteTagsInput{}
        output := &autoscaling.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWarmPool", func(t *testing.T) {
        input := &autoscaling.DeleteWarmPoolInput{}
        output := &autoscaling.DeleteWarmPoolOutput{}

        mockClient.On("DeleteWarmPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWarmPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountLimits", func(t *testing.T) {
        input := &autoscaling.DescribeAccountLimitsInput{}
        output := &autoscaling.DescribeAccountLimitsOutput{}

        mockClient.On("DescribeAccountLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAdjustmentTypes", func(t *testing.T) {
        input := &autoscaling.DescribeAdjustmentTypesInput{}
        output := &autoscaling.DescribeAdjustmentTypesOutput{}

        mockClient.On("DescribeAdjustmentTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAdjustmentTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoScalingGroups", func(t *testing.T) {
        input := &autoscaling.DescribeAutoScalingGroupsInput{}
        output := &autoscaling.DescribeAutoScalingGroupsOutput{}

        mockClient.On("DescribeAutoScalingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoScalingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoScalingInstances", func(t *testing.T) {
        input := &autoscaling.DescribeAutoScalingInstancesInput{}
        output := &autoscaling.DescribeAutoScalingInstancesOutput{}

        mockClient.On("DescribeAutoScalingInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoScalingInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAutoScalingNotificationTypes", func(t *testing.T) {
        input := &autoscaling.DescribeAutoScalingNotificationTypesInput{}
        output := &autoscaling.DescribeAutoScalingNotificationTypesOutput{}

        mockClient.On("DescribeAutoScalingNotificationTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAutoScalingNotificationTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceRefreshes", func(t *testing.T) {
        input := &autoscaling.DescribeInstanceRefreshesInput{}
        output := &autoscaling.DescribeInstanceRefreshesOutput{}

        mockClient.On("DescribeInstanceRefreshes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceRefreshes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLaunchConfigurations", func(t *testing.T) {
        input := &autoscaling.DescribeLaunchConfigurationsInput{}
        output := &autoscaling.DescribeLaunchConfigurationsOutput{}

        mockClient.On("DescribeLaunchConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLaunchConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLifecycleHookTypes", func(t *testing.T) {
        input := &autoscaling.DescribeLifecycleHookTypesInput{}
        output := &autoscaling.DescribeLifecycleHookTypesOutput{}

        mockClient.On("DescribeLifecycleHookTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLifecycleHookTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLifecycleHooks", func(t *testing.T) {
        input := &autoscaling.DescribeLifecycleHooksInput{}
        output := &autoscaling.DescribeLifecycleHooksOutput{}

        mockClient.On("DescribeLifecycleHooks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLifecycleHooks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancerTargetGroups", func(t *testing.T) {
        input := &autoscaling.DescribeLoadBalancerTargetGroupsInput{}
        output := &autoscaling.DescribeLoadBalancerTargetGroupsOutput{}

        mockClient.On("DescribeLoadBalancerTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancerTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancers", func(t *testing.T) {
        input := &autoscaling.DescribeLoadBalancersInput{}
        output := &autoscaling.DescribeLoadBalancersOutput{}

        mockClient.On("DescribeLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetricCollectionTypes", func(t *testing.T) {
        input := &autoscaling.DescribeMetricCollectionTypesInput{}
        output := &autoscaling.DescribeMetricCollectionTypesOutput{}

        mockClient.On("DescribeMetricCollectionTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetricCollectionTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotificationConfigurations", func(t *testing.T) {
        input := &autoscaling.DescribeNotificationConfigurationsInput{}
        output := &autoscaling.DescribeNotificationConfigurationsOutput{}

        mockClient.On("DescribeNotificationConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotificationConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePolicies", func(t *testing.T) {
        input := &autoscaling.DescribePoliciesInput{}
        output := &autoscaling.DescribePoliciesOutput{}

        mockClient.On("DescribePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingActivities", func(t *testing.T) {
        input := &autoscaling.DescribeScalingActivitiesInput{}
        output := &autoscaling.DescribeScalingActivitiesOutput{}

        mockClient.On("DescribeScalingActivities", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingProcessTypes", func(t *testing.T) {
        input := &autoscaling.DescribeScalingProcessTypesInput{}
        output := &autoscaling.DescribeScalingProcessTypesOutput{}

        mockClient.On("DescribeScalingProcessTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingProcessTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledActions", func(t *testing.T) {
        input := &autoscaling.DescribeScheduledActionsInput{}
        output := &autoscaling.DescribeScheduledActionsOutput{}

        mockClient.On("DescribeScheduledActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &autoscaling.DescribeTagsInput{}
        output := &autoscaling.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTerminationPolicyTypes", func(t *testing.T) {
        input := &autoscaling.DescribeTerminationPolicyTypesInput{}
        output := &autoscaling.DescribeTerminationPolicyTypesOutput{}

        mockClient.On("DescribeTerminationPolicyTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTerminationPolicyTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficSources", func(t *testing.T) {
        input := &autoscaling.DescribeTrafficSourcesInput{}
        output := &autoscaling.DescribeTrafficSourcesOutput{}

        mockClient.On("DescribeTrafficSources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWarmPool", func(t *testing.T) {
        input := &autoscaling.DescribeWarmPoolInput{}
        output := &autoscaling.DescribeWarmPoolOutput{}

        mockClient.On("DescribeWarmPool", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWarmPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachInstances", func(t *testing.T) {
        input := &autoscaling.DetachInstancesInput{}
        output := &autoscaling.DetachInstancesOutput{}

        mockClient.On("DetachInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DetachInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachLoadBalancerTargetGroups", func(t *testing.T) {
        input := &autoscaling.DetachLoadBalancerTargetGroupsInput{}
        output := &autoscaling.DetachLoadBalancerTargetGroupsOutput{}

        mockClient.On("DetachLoadBalancerTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DetachLoadBalancerTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachLoadBalancers", func(t *testing.T) {
        input := &autoscaling.DetachLoadBalancersInput{}
        output := &autoscaling.DetachLoadBalancersOutput{}

        mockClient.On("DetachLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.DetachLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachTrafficSources", func(t *testing.T) {
        input := &autoscaling.DetachTrafficSourcesInput{}
        output := &autoscaling.DetachTrafficSourcesOutput{}

        mockClient.On("DetachTrafficSources", ctx, input).Return(output, nil)

        result, err := mockClient.DetachTrafficSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableMetricsCollection", func(t *testing.T) {
        input := &autoscaling.DisableMetricsCollectionInput{}
        output := &autoscaling.DisableMetricsCollectionOutput{}

        mockClient.On("DisableMetricsCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DisableMetricsCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableMetricsCollection", func(t *testing.T) {
        input := &autoscaling.EnableMetricsCollectionInput{}
        output := &autoscaling.EnableMetricsCollectionOutput{}

        mockClient.On("EnableMetricsCollection", ctx, input).Return(output, nil)

        result, err := mockClient.EnableMetricsCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnterStandby", func(t *testing.T) {
        input := &autoscaling.EnterStandbyInput{}
        output := &autoscaling.EnterStandbyOutput{}

        mockClient.On("EnterStandby", ctx, input).Return(output, nil)

        result, err := mockClient.EnterStandby(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecutePolicy", func(t *testing.T) {
        input := &autoscaling.ExecutePolicyInput{}
        output := &autoscaling.ExecutePolicyOutput{}

        mockClient.On("ExecutePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ExecutePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExitStandby", func(t *testing.T) {
        input := &autoscaling.ExitStandbyInput{}
        output := &autoscaling.ExitStandbyOutput{}

        mockClient.On("ExitStandby", ctx, input).Return(output, nil)

        result, err := mockClient.ExitStandby(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPredictiveScalingForecast", func(t *testing.T) {
        input := &autoscaling.GetPredictiveScalingForecastInput{}
        output := &autoscaling.GetPredictiveScalingForecastOutput{}

        mockClient.On("GetPredictiveScalingForecast", ctx, input).Return(output, nil)

        result, err := mockClient.GetPredictiveScalingForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLifecycleHook", func(t *testing.T) {
        input := &autoscaling.PutLifecycleHookInput{}
        output := &autoscaling.PutLifecycleHookOutput{}

        mockClient.On("PutLifecycleHook", ctx, input).Return(output, nil)

        result, err := mockClient.PutLifecycleHook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutNotificationConfiguration", func(t *testing.T) {
        input := &autoscaling.PutNotificationConfigurationInput{}
        output := &autoscaling.PutNotificationConfigurationOutput{}

        mockClient.On("PutNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutScalingPolicy", func(t *testing.T) {
        input := &autoscaling.PutScalingPolicyInput{}
        output := &autoscaling.PutScalingPolicyOutput{}

        mockClient.On("PutScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutScheduledUpdateGroupAction", func(t *testing.T) {
        input := &autoscaling.PutScheduledUpdateGroupActionInput{}
        output := &autoscaling.PutScheduledUpdateGroupActionOutput{}

        mockClient.On("PutScheduledUpdateGroupAction", ctx, input).Return(output, nil)

        result, err := mockClient.PutScheduledUpdateGroupAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutWarmPool", func(t *testing.T) {
        input := &autoscaling.PutWarmPoolInput{}
        output := &autoscaling.PutWarmPoolOutput{}

        mockClient.On("PutWarmPool", ctx, input).Return(output, nil)

        result, err := mockClient.PutWarmPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecordLifecycleActionHeartbeat", func(t *testing.T) {
        input := &autoscaling.RecordLifecycleActionHeartbeatInput{}
        output := &autoscaling.RecordLifecycleActionHeartbeatOutput{}

        mockClient.On("RecordLifecycleActionHeartbeat", ctx, input).Return(output, nil)

        result, err := mockClient.RecordLifecycleActionHeartbeat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeProcesses", func(t *testing.T) {
        input := &autoscaling.ResumeProcessesInput{}
        output := &autoscaling.ResumeProcessesOutput{}

        mockClient.On("ResumeProcesses", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeProcesses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRollbackInstanceRefresh", func(t *testing.T) {
        input := &autoscaling.RollbackInstanceRefreshInput{}
        output := &autoscaling.RollbackInstanceRefreshOutput{}

        mockClient.On("RollbackInstanceRefresh", ctx, input).Return(output, nil)

        result, err := mockClient.RollbackInstanceRefresh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDesiredCapacity", func(t *testing.T) {
        input := &autoscaling.SetDesiredCapacityInput{}
        output := &autoscaling.SetDesiredCapacityOutput{}

        mockClient.On("SetDesiredCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.SetDesiredCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetInstanceHealth", func(t *testing.T) {
        input := &autoscaling.SetInstanceHealthInput{}
        output := &autoscaling.SetInstanceHealthOutput{}

        mockClient.On("SetInstanceHealth", ctx, input).Return(output, nil)

        result, err := mockClient.SetInstanceHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetInstanceProtection", func(t *testing.T) {
        input := &autoscaling.SetInstanceProtectionInput{}
        output := &autoscaling.SetInstanceProtectionOutput{}

        mockClient.On("SetInstanceProtection", ctx, input).Return(output, nil)

        result, err := mockClient.SetInstanceProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstanceRefresh", func(t *testing.T) {
        input := &autoscaling.StartInstanceRefreshInput{}
        output := &autoscaling.StartInstanceRefreshOutput{}

        mockClient.On("StartInstanceRefresh", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstanceRefresh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSuspendProcesses", func(t *testing.T) {
        input := &autoscaling.SuspendProcessesInput{}
        output := &autoscaling.SuspendProcessesOutput{}

        mockClient.On("SuspendProcesses", ctx, input).Return(output, nil)

        result, err := mockClient.SuspendProcesses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateInstanceInAutoScalingGroup", func(t *testing.T) {
        input := &autoscaling.TerminateInstanceInAutoScalingGroupInput{}
        output := &autoscaling.TerminateInstanceInAutoScalingGroupOutput{}

        mockClient.On("TerminateInstanceInAutoScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateInstanceInAutoScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAutoScalingGroup", func(t *testing.T) {
        input := &autoscaling.UpdateAutoScalingGroupInput{}
        output := &autoscaling.UpdateAutoScalingGroupOutput{}

        mockClient.On("UpdateAutoScalingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAutoScalingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
