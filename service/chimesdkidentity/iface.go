
package chimesdkidentity

import (
    "github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
)

// IChimesdkidentity defines the interface for chimesdkidentity
type IChimesdkidentity interface {
 Options() Options 
 CreateAppInstance(ctx context.Context, params *CreateAppInstanceInput, optFns ...func(*Options)) (*CreateAppInstanceOutput, error) 
 CreateAppInstanceAdmin(ctx context.Context, params *CreateAppInstanceAdminInput, optFns ...func(*Options)) (*CreateAppInstanceAdminOutput, error) 
 CreateAppInstanceBot(ctx context.Context, params *CreateAppInstanceBotInput, optFns ...func(*Options)) (*CreateAppInstanceBotOutput, error) 
 CreateAppInstanceUser(ctx context.Context, params *CreateAppInstanceUserInput, optFns ...func(*Options)) (*CreateAppInstanceUserOutput, error) 
 DeleteAppInstance(ctx context.Context, params *DeleteAppInstanceInput, optFns ...func(*Options)) (*DeleteAppInstanceOutput, error) 
 DeleteAppInstanceAdmin(ctx context.Context, params *DeleteAppInstanceAdminInput, optFns ...func(*Options)) (*DeleteAppInstanceAdminOutput, error) 
 DeleteAppInstanceBot(ctx context.Context, params *DeleteAppInstanceBotInput, optFns ...func(*Options)) (*DeleteAppInstanceBotOutput, error) 
 DeleteAppInstanceUser(ctx context.Context, params *DeleteAppInstanceUserInput, optFns ...func(*Options)) (*DeleteAppInstanceUserOutput, error) 
 DeregisterAppInstanceUserEndpoint(ctx context.Context, params *DeregisterAppInstanceUserEndpointInput, optFns ...func(*Options)) (*DeregisterAppInstanceUserEndpointOutput, error) 
 DescribeAppInstance(ctx context.Context, params *DescribeAppInstanceInput, optFns ...func(*Options)) (*DescribeAppInstanceOutput, error) 
 DescribeAppInstanceAdmin(ctx context.Context, params *DescribeAppInstanceAdminInput, optFns ...func(*Options)) (*DescribeAppInstanceAdminOutput, error) 
 DescribeAppInstanceBot(ctx context.Context, params *DescribeAppInstanceBotInput, optFns ...func(*Options)) (*DescribeAppInstanceBotOutput, error) 
 DescribeAppInstanceUser(ctx context.Context, params *DescribeAppInstanceUserInput, optFns ...func(*Options)) (*DescribeAppInstanceUserOutput, error) 
 DescribeAppInstanceUserEndpoint(ctx context.Context, params *DescribeAppInstanceUserEndpointInput, optFns ...func(*Options)) (*DescribeAppInstanceUserEndpointOutput, error) 
 GetAppInstanceRetentionSettings(ctx context.Context, params *GetAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*GetAppInstanceRetentionSettingsOutput, error) 
 ListAppInstanceAdmins(ctx context.Context, params *ListAppInstanceAdminsInput, optFns ...func(*Options)) (*ListAppInstanceAdminsOutput, error) 
 ListAppInstanceBots(ctx context.Context, params *ListAppInstanceBotsInput, optFns ...func(*Options)) (*ListAppInstanceBotsOutput, error) 
 ListAppInstanceUserEndpoints(ctx context.Context, params *ListAppInstanceUserEndpointsInput, optFns ...func(*Options)) (*ListAppInstanceUserEndpointsOutput, error) 
 ListAppInstanceUsers(ctx context.Context, params *ListAppInstanceUsersInput, optFns ...func(*Options)) (*ListAppInstanceUsersOutput, error) 
 ListAppInstances(ctx context.Context, params *ListAppInstancesInput, optFns ...func(*Options)) (*ListAppInstancesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAppInstanceRetentionSettings(ctx context.Context, params *PutAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*PutAppInstanceRetentionSettingsOutput, error) 
 PutAppInstanceUserExpirationSettings(ctx context.Context, params *PutAppInstanceUserExpirationSettingsInput, optFns ...func(*Options)) (*PutAppInstanceUserExpirationSettingsOutput, error) 
 RegisterAppInstanceUserEndpoint(ctx context.Context, params *RegisterAppInstanceUserEndpointInput, optFns ...func(*Options)) (*RegisterAppInstanceUserEndpointOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAppInstance(ctx context.Context, params *UpdateAppInstanceInput, optFns ...func(*Options)) (*UpdateAppInstanceOutput, error) 
 UpdateAppInstanceBot(ctx context.Context, params *UpdateAppInstanceBotInput, optFns ...func(*Options)) (*UpdateAppInstanceBotOutput, error) 
 UpdateAppInstanceUser(ctx context.Context, params *UpdateAppInstanceUserInput, optFns ...func(*Options)) (*UpdateAppInstanceUserOutput, error) 
 UpdateAppInstanceUserEndpoint(ctx context.Context, params *UpdateAppInstanceUserEndpointInput, optFns ...func(*Options)) (*UpdateAppInstanceUserEndpointOutput, error) 
}
