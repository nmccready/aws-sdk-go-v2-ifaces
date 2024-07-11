
package budgets

import (
    "github.com/aws/aws-sdk-go-v2/service/budgets"
)

// IBudgets defines the interface for budgets
type IBudgets interface {
 Options() Options 
 CreateBudget(ctx context.Context, params *CreateBudgetInput, optFns ...func(*Options)) (*CreateBudgetOutput, error) 
 CreateBudgetAction(ctx context.Context, params *CreateBudgetActionInput, optFns ...func(*Options)) (*CreateBudgetActionOutput, error) 
 CreateNotification(ctx context.Context, params *CreateNotificationInput, optFns ...func(*Options)) (*CreateNotificationOutput, error) 
 CreateSubscriber(ctx context.Context, params *CreateSubscriberInput, optFns ...func(*Options)) (*CreateSubscriberOutput, error) 
 DeleteBudget(ctx context.Context, params *DeleteBudgetInput, optFns ...func(*Options)) (*DeleteBudgetOutput, error) 
 DeleteBudgetAction(ctx context.Context, params *DeleteBudgetActionInput, optFns ...func(*Options)) (*DeleteBudgetActionOutput, error) 
 DeleteNotification(ctx context.Context, params *DeleteNotificationInput, optFns ...func(*Options)) (*DeleteNotificationOutput, error) 
 DeleteSubscriber(ctx context.Context, params *DeleteSubscriberInput, optFns ...func(*Options)) (*DeleteSubscriberOutput, error) 
 DescribeBudget(ctx context.Context, params *DescribeBudgetInput, optFns ...func(*Options)) (*DescribeBudgetOutput, error) 
 DescribeBudgetAction(ctx context.Context, params *DescribeBudgetActionInput, optFns ...func(*Options)) (*DescribeBudgetActionOutput, error) 
 DescribeBudgetActionHistories(ctx context.Context, params *DescribeBudgetActionHistoriesInput, optFns ...func(*Options)) (*DescribeBudgetActionHistoriesOutput, error) 
 DescribeBudgetActionsForAccount(ctx context.Context, params *DescribeBudgetActionsForAccountInput, optFns ...func(*Options)) (*DescribeBudgetActionsForAccountOutput, error) 
 DescribeBudgetActionsForBudget(ctx context.Context, params *DescribeBudgetActionsForBudgetInput, optFns ...func(*Options)) (*DescribeBudgetActionsForBudgetOutput, error) 
 DescribeBudgetNotificationsForAccount(ctx context.Context, params *DescribeBudgetNotificationsForAccountInput, optFns ...func(*Options)) (*DescribeBudgetNotificationsForAccountOutput, error) 
 DescribeBudgetPerformanceHistory(ctx context.Context, params *DescribeBudgetPerformanceHistoryInput, optFns ...func(*Options)) (*DescribeBudgetPerformanceHistoryOutput, error) 
 DescribeBudgets(ctx context.Context, params *DescribeBudgetsInput, optFns ...func(*Options)) (*DescribeBudgetsOutput, error) 
 DescribeNotificationsForBudget(ctx context.Context, params *DescribeNotificationsForBudgetInput, optFns ...func(*Options)) (*DescribeNotificationsForBudgetOutput, error) 
 DescribeSubscribersForNotification(ctx context.Context, params *DescribeSubscribersForNotificationInput, optFns ...func(*Options)) (*DescribeSubscribersForNotificationOutput, error) 
 ExecuteBudgetAction(ctx context.Context, params *ExecuteBudgetActionInput, optFns ...func(*Options)) (*ExecuteBudgetActionOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBudget(ctx context.Context, params *UpdateBudgetInput, optFns ...func(*Options)) (*UpdateBudgetOutput, error) 
 UpdateBudgetAction(ctx context.Context, params *UpdateBudgetActionInput, optFns ...func(*Options)) (*UpdateBudgetActionOutput, error) 
 UpdateNotification(ctx context.Context, params *UpdateNotificationInput, optFns ...func(*Options)) (*UpdateNotificationOutput, error) 
 UpdateSubscriber(ctx context.Context, params *UpdateSubscriberInput, optFns ...func(*Options)) (*UpdateSubscriberOutput, error) 
}
