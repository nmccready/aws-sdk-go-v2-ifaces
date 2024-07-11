
package snowdevicemanagement

import (
    "github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
)

// ISnowdevicemanagement defines the interface for snowdevicemanagement
type ISnowdevicemanagement interface {
 Options() Options 
 CancelTask(ctx context.Context, params *CancelTaskInput, optFns ...func(*Options)) (*CancelTaskOutput, error) 
 CreateTask(ctx context.Context, params *CreateTaskInput, optFns ...func(*Options)) (*CreateTaskOutput, error) 
 DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) 
 DescribeDeviceEc2Instances(ctx context.Context, params *DescribeDeviceEc2InstancesInput, optFns ...func(*Options)) (*DescribeDeviceEc2InstancesOutput, error) 
 DescribeExecution(ctx context.Context, params *DescribeExecutionInput, optFns ...func(*Options)) (*DescribeExecutionOutput, error) 
 DescribeTask(ctx context.Context, params *DescribeTaskInput, optFns ...func(*Options)) (*DescribeTaskOutput, error) 
 ListDeviceResources(ctx context.Context, params *ListDeviceResourcesInput, optFns ...func(*Options)) (*ListDeviceResourcesOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
