
package securitylake

import (
    "github.com/aws/aws-sdk-go-v2/service/securitylake"
)

// ISecuritylake defines the interface for securitylake
type ISecuritylake interface {
 Options() Options 
 CreateAwsLogSource(ctx context.Context, params *CreateAwsLogSourceInput, optFns ...func(*Options)) (*CreateAwsLogSourceOutput, error) 
 CreateCustomLogSource(ctx context.Context, params *CreateCustomLogSourceInput, optFns ...func(*Options)) (*CreateCustomLogSourceOutput, error) 
 CreateDataLake(ctx context.Context, params *CreateDataLakeInput, optFns ...func(*Options)) (*CreateDataLakeOutput, error) 
 CreateDataLakeExceptionSubscription(ctx context.Context, params *CreateDataLakeExceptionSubscriptionInput, optFns ...func(*Options)) (*CreateDataLakeExceptionSubscriptionOutput, error) 
 CreateDataLakeOrganizationConfiguration(ctx context.Context, params *CreateDataLakeOrganizationConfigurationInput, optFns ...func(*Options)) (*CreateDataLakeOrganizationConfigurationOutput, error) 
 CreateSubscriber(ctx context.Context, params *CreateSubscriberInput, optFns ...func(*Options)) (*CreateSubscriberOutput, error) 
 CreateSubscriberNotification(ctx context.Context, params *CreateSubscriberNotificationInput, optFns ...func(*Options)) (*CreateSubscriberNotificationOutput, error) 
 DeleteAwsLogSource(ctx context.Context, params *DeleteAwsLogSourceInput, optFns ...func(*Options)) (*DeleteAwsLogSourceOutput, error) 
 DeleteCustomLogSource(ctx context.Context, params *DeleteCustomLogSourceInput, optFns ...func(*Options)) (*DeleteCustomLogSourceOutput, error) 
 DeleteDataLake(ctx context.Context, params *DeleteDataLakeInput, optFns ...func(*Options)) (*DeleteDataLakeOutput, error) 
 DeleteDataLakeExceptionSubscription(ctx context.Context, params *DeleteDataLakeExceptionSubscriptionInput, optFns ...func(*Options)) (*DeleteDataLakeExceptionSubscriptionOutput, error) 
 DeleteDataLakeOrganizationConfiguration(ctx context.Context, params *DeleteDataLakeOrganizationConfigurationInput, optFns ...func(*Options)) (*DeleteDataLakeOrganizationConfigurationOutput, error) 
 DeleteSubscriber(ctx context.Context, params *DeleteSubscriberInput, optFns ...func(*Options)) (*DeleteSubscriberOutput, error) 
 DeleteSubscriberNotification(ctx context.Context, params *DeleteSubscriberNotificationInput, optFns ...func(*Options)) (*DeleteSubscriberNotificationOutput, error) 
 DeregisterDataLakeDelegatedAdministrator(ctx context.Context, params *DeregisterDataLakeDelegatedAdministratorInput, optFns ...func(*Options)) (*DeregisterDataLakeDelegatedAdministratorOutput, error) 
 GetDataLakeExceptionSubscription(ctx context.Context, params *GetDataLakeExceptionSubscriptionInput, optFns ...func(*Options)) (*GetDataLakeExceptionSubscriptionOutput, error) 
 GetDataLakeOrganizationConfiguration(ctx context.Context, params *GetDataLakeOrganizationConfigurationInput, optFns ...func(*Options)) (*GetDataLakeOrganizationConfigurationOutput, error) 
 GetDataLakeSources(ctx context.Context, params *GetDataLakeSourcesInput, optFns ...func(*Options)) (*GetDataLakeSourcesOutput, error) 
 GetSubscriber(ctx context.Context, params *GetSubscriberInput, optFns ...func(*Options)) (*GetSubscriberOutput, error) 
 ListDataLakeExceptions(ctx context.Context, params *ListDataLakeExceptionsInput, optFns ...func(*Options)) (*ListDataLakeExceptionsOutput, error) 
 ListDataLakes(ctx context.Context, params *ListDataLakesInput, optFns ...func(*Options)) (*ListDataLakesOutput, error) 
 ListLogSources(ctx context.Context, params *ListLogSourcesInput, optFns ...func(*Options)) (*ListLogSourcesOutput, error) 
 ListSubscribers(ctx context.Context, params *ListSubscribersInput, optFns ...func(*Options)) (*ListSubscribersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterDataLakeDelegatedAdministrator(ctx context.Context, params *RegisterDataLakeDelegatedAdministratorInput, optFns ...func(*Options)) (*RegisterDataLakeDelegatedAdministratorOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDataLake(ctx context.Context, params *UpdateDataLakeInput, optFns ...func(*Options)) (*UpdateDataLakeOutput, error) 
 UpdateDataLakeExceptionSubscription(ctx context.Context, params *UpdateDataLakeExceptionSubscriptionInput, optFns ...func(*Options)) (*UpdateDataLakeExceptionSubscriptionOutput, error) 
 UpdateSubscriber(ctx context.Context, params *UpdateSubscriberInput, optFns ...func(*Options)) (*UpdateSubscriberOutput, error) 
 UpdateSubscriberNotification(ctx context.Context, params *UpdateSubscriberNotificationInput, optFns ...func(*Options)) (*UpdateSubscriberNotificationOutput, error) 
}
