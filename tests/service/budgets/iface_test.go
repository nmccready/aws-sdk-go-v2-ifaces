package budgets_test

// tests for the budgets service interface for this ../../../service/budgets/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/budgets/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/budgets/budgets_iface"
	"github.com/stretchr/testify/assert"
)

func TestBudgetsServiceCanBeMocked(t *testing.T) {
	var iface budgets_iface.IClient
	iface = &budgets.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := budgets.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBudget", func(t *testing.T) {
        input := &budgets.CreateBudgetInput{}
        output := &budgets.CreateBudgetOutput{}

        mockClient.On("CreateBudget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBudgetAction", func(t *testing.T) {
        input := &budgets.CreateBudgetActionInput{}
        output := &budgets.CreateBudgetActionOutput{}

        mockClient.On("CreateBudgetAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBudgetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotification", func(t *testing.T) {
        input := &budgets.CreateNotificationInput{}
        output := &budgets.CreateNotificationOutput{}

        mockClient.On("CreateNotification", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriber", func(t *testing.T) {
        input := &budgets.CreateSubscriberInput{}
        output := &budgets.CreateSubscriberOutput{}

        mockClient.On("CreateSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBudget", func(t *testing.T) {
        input := &budgets.DeleteBudgetInput{}
        output := &budgets.DeleteBudgetOutput{}

        mockClient.On("DeleteBudget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBudgetAction", func(t *testing.T) {
        input := &budgets.DeleteBudgetActionInput{}
        output := &budgets.DeleteBudgetActionOutput{}

        mockClient.On("DeleteBudgetAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBudgetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotification", func(t *testing.T) {
        input := &budgets.DeleteNotificationInput{}
        output := &budgets.DeleteNotificationOutput{}

        mockClient.On("DeleteNotification", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriber", func(t *testing.T) {
        input := &budgets.DeleteSubscriberInput{}
        output := &budgets.DeleteSubscriberOutput{}

        mockClient.On("DeleteSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudget", func(t *testing.T) {
        input := &budgets.DescribeBudgetInput{}
        output := &budgets.DescribeBudgetOutput{}

        mockClient.On("DescribeBudget", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetAction", func(t *testing.T) {
        input := &budgets.DescribeBudgetActionInput{}
        output := &budgets.DescribeBudgetActionOutput{}

        mockClient.On("DescribeBudgetAction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetActionHistories", func(t *testing.T) {
        input := &budgets.DescribeBudgetActionHistoriesInput{}
        output := &budgets.DescribeBudgetActionHistoriesOutput{}

        mockClient.On("DescribeBudgetActionHistories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetActionHistories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetActionsForAccount", func(t *testing.T) {
        input := &budgets.DescribeBudgetActionsForAccountInput{}
        output := &budgets.DescribeBudgetActionsForAccountOutput{}

        mockClient.On("DescribeBudgetActionsForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetActionsForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetActionsForBudget", func(t *testing.T) {
        input := &budgets.DescribeBudgetActionsForBudgetInput{}
        output := &budgets.DescribeBudgetActionsForBudgetOutput{}

        mockClient.On("DescribeBudgetActionsForBudget", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetActionsForBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetNotificationsForAccount", func(t *testing.T) {
        input := &budgets.DescribeBudgetNotificationsForAccountInput{}
        output := &budgets.DescribeBudgetNotificationsForAccountOutput{}

        mockClient.On("DescribeBudgetNotificationsForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetNotificationsForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgetPerformanceHistory", func(t *testing.T) {
        input := &budgets.DescribeBudgetPerformanceHistoryInput{}
        output := &budgets.DescribeBudgetPerformanceHistoryOutput{}

        mockClient.On("DescribeBudgetPerformanceHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgetPerformanceHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBudgets", func(t *testing.T) {
        input := &budgets.DescribeBudgetsInput{}
        output := &budgets.DescribeBudgetsOutput{}

        mockClient.On("DescribeBudgets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBudgets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotificationsForBudget", func(t *testing.T) {
        input := &budgets.DescribeNotificationsForBudgetInput{}
        output := &budgets.DescribeNotificationsForBudgetOutput{}

        mockClient.On("DescribeNotificationsForBudget", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotificationsForBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubscribersForNotification", func(t *testing.T) {
        input := &budgets.DescribeSubscribersForNotificationInput{}
        output := &budgets.DescribeSubscribersForNotificationOutput{}

        mockClient.On("DescribeSubscribersForNotification", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubscribersForNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteBudgetAction", func(t *testing.T) {
        input := &budgets.ExecuteBudgetActionInput{}
        output := &budgets.ExecuteBudgetActionOutput{}

        mockClient.On("ExecuteBudgetAction", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteBudgetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &budgets.ListTagsForResourceInput{}
        output := &budgets.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &budgets.TagResourceInput{}
        output := &budgets.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &budgets.UntagResourceInput{}
        output := &budgets.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBudget", func(t *testing.T) {
        input := &budgets.UpdateBudgetInput{}
        output := &budgets.UpdateBudgetOutput{}

        mockClient.On("UpdateBudget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBudgetAction", func(t *testing.T) {
        input := &budgets.UpdateBudgetActionInput{}
        output := &budgets.UpdateBudgetActionOutput{}

        mockClient.On("UpdateBudgetAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBudgetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotification", func(t *testing.T) {
        input := &budgets.UpdateNotificationInput{}
        output := &budgets.UpdateNotificationOutput{}

        mockClient.On("UpdateNotification", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriber", func(t *testing.T) {
        input := &budgets.UpdateSubscriberInput{}
        output := &budgets.UpdateSubscriberOutput{}

        mockClient.On("UpdateSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
