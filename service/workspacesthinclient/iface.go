
package workspacesthinclient

import (
    "github.com/aws/aws-sdk-go-v2/service/workspacesthinclient"
)

// IClient defines the interface for workspacesthinclient
type IClient interface {
 Options() Options 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 DeleteDevice(ctx context.Context, params *DeleteDeviceInput, optFns ...func(*Options)) (*DeleteDeviceOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 DeregisterDevice(ctx context.Context, params *DeregisterDeviceInput, optFns ...func(*Options)) (*DeregisterDeviceOutput, error) 
 GetDevice(ctx context.Context, params *GetDeviceInput, optFns ...func(*Options)) (*GetDeviceOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 GetSoftwareSet(ctx context.Context, params *GetSoftwareSetInput, optFns ...func(*Options)) (*GetSoftwareSetOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListSoftwareSets(ctx context.Context, params *ListSoftwareSetsInput, optFns ...func(*Options)) (*ListSoftwareSetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDevice(ctx context.Context, params *UpdateDeviceInput, optFns ...func(*Options)) (*UpdateDeviceOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
 UpdateSoftwareSet(ctx context.Context, params *UpdateSoftwareSetInput, optFns ...func(*Options)) (*UpdateSoftwareSetOutput, error) 
}
