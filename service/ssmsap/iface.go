
package ssmsap

import (
    "github.com/aws/aws-sdk-go-v2/service/ssmsap"
)

// ISsmsap defines the interface for ssmsap
type ISsmsap interface {
 Options() Options 
 DeleteResourcePermission(ctx context.Context, params *DeleteResourcePermissionInput, optFns ...func(*Options)) (*DeleteResourcePermissionOutput, error) 
 DeregisterApplication(ctx context.Context, params *DeregisterApplicationInput, optFns ...func(*Options)) (*DeregisterApplicationOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetComponent(ctx context.Context, params *GetComponentInput, optFns ...func(*Options)) (*GetComponentOutput, error) 
 GetDatabase(ctx context.Context, params *GetDatabaseInput, optFns ...func(*Options)) (*GetDatabaseOutput, error) 
 GetOperation(ctx context.Context, params *GetOperationInput, optFns ...func(*Options)) (*GetOperationOutput, error) 
 GetResourcePermission(ctx context.Context, params *GetResourcePermissionInput, optFns ...func(*Options)) (*GetResourcePermissionOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) 
 ListOperationEvents(ctx context.Context, params *ListOperationEventsInput, optFns ...func(*Options)) (*ListOperationEventsOutput, error) 
 ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutResourcePermission(ctx context.Context, params *PutResourcePermissionInput, optFns ...func(*Options)) (*PutResourcePermissionOutput, error) 
 RegisterApplication(ctx context.Context, params *RegisterApplicationInput, optFns ...func(*Options)) (*RegisterApplicationOutput, error) 
 StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) 
 StartApplicationRefresh(ctx context.Context, params *StartApplicationRefreshInput, optFns ...func(*Options)) (*StartApplicationRefreshOutput, error) 
 StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplicationSettings(ctx context.Context, params *UpdateApplicationSettingsInput, optFns ...func(*Options)) (*UpdateApplicationSettingsOutput, error) 
}
