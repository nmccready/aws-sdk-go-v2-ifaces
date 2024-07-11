
package iot1clickdevicesservice

import (
    "github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

// IClient defines the interface for iot1clickdevicesservice
type IClient interface {
 Options() Options 
 ClaimDevicesByClaimCode(ctx context.Context, params *ClaimDevicesByClaimCodeInput, optFns ...func(*Options)) (*ClaimDevicesByClaimCodeOutput, error) 
 DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) 
 FinalizeDeviceClaim(ctx context.Context, params *FinalizeDeviceClaimInput, optFns ...func(*Options)) (*FinalizeDeviceClaimOutput, error) 
 GetDeviceMethods(ctx context.Context, params *GetDeviceMethodsInput, optFns ...func(*Options)) (*GetDeviceMethodsOutput, error) 
 InitiateDeviceClaim(ctx context.Context, params *InitiateDeviceClaimInput, optFns ...func(*Options)) (*InitiateDeviceClaimOutput, error) 
 InvokeDeviceMethod(ctx context.Context, params *InvokeDeviceMethodInput, optFns ...func(*Options)) (*InvokeDeviceMethodOutput, error) 
 ListDeviceEvents(ctx context.Context, params *ListDeviceEventsInput, optFns ...func(*Options)) (*ListDeviceEventsOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UnclaimDevice(ctx context.Context, params *UnclaimDeviceInput, optFns ...func(*Options)) (*UnclaimDeviceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDeviceState(ctx context.Context, params *UpdateDeviceStateInput, optFns ...func(*Options)) (*UpdateDeviceStateOutput, error) 
}
