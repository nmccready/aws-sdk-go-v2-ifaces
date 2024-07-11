
package migrationhubrefactorspaces

import (
    "github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"
)

// IClient defines the interface for migrationhubrefactorspaces
type IClient interface {
 Options() Options 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) 
 CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteRoute(ctx context.Context, params *DeleteRouteInput, optFns ...func(*Options)) (*DeleteRouteOutput, error) 
 DeleteService(ctx context.Context, params *DeleteServiceInput, optFns ...func(*Options)) (*DeleteServiceOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 GetRoute(ctx context.Context, params *GetRouteInput, optFns ...func(*Options)) (*GetRouteOutput, error) 
 GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListEnvironmentVpcs(ctx context.Context, params *ListEnvironmentVpcsInput, optFns ...func(*Options)) (*ListEnvironmentVpcsOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListRoutes(ctx context.Context, params *ListRoutesInput, optFns ...func(*Options)) (*ListRoutesOutput, error) 
 ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateRoute(ctx context.Context, params *UpdateRouteInput, optFns ...func(*Options)) (*UpdateRouteOutput, error) 
}
