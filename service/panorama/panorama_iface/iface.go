
package panorama_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/panorama"
)

// IClient defines the interface for panorama
type IClient interface {
 Options() Options 
 CreateApplicationInstance(ctx context.Context, params *CreateApplicationInstanceInput, optFns ...func(*Options)) (*CreateApplicationInstanceOutput, error) 
 CreateJobForDevices(ctx context.Context, params *CreateJobForDevicesInput, optFns ...func(*Options)) (*CreateJobForDevicesOutput, error) 
 CreateNodeFromTemplateJob(ctx context.Context, params *CreateNodeFromTemplateJobInput, optFns ...func(*Options)) (*CreateNodeFromTemplateJobOutput, error) 
 CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) 
 CreatePackageImportJob(ctx context.Context, params *CreatePackageImportJobInput, optFns ...func(*Options)) (*CreatePackageImportJobOutput, error) 
 DeleteDevice(ctx context.Context, params *DeleteDeviceInput, optFns ...func(*Options)) (*DeleteDeviceOutput, error) 
 DeletePackage(ctx context.Context, params *DeletePackageInput, optFns ...func(*Options)) (*DeletePackageOutput, error) 
 DeregisterPackageVersion(ctx context.Context, params *DeregisterPackageVersionInput, optFns ...func(*Options)) (*DeregisterPackageVersionOutput, error) 
 DescribeApplicationInstance(ctx context.Context, params *DescribeApplicationInstanceInput, optFns ...func(*Options)) (*DescribeApplicationInstanceOutput, error) 
 DescribeApplicationInstanceDetails(ctx context.Context, params *DescribeApplicationInstanceDetailsInput, optFns ...func(*Options)) (*DescribeApplicationInstanceDetailsOutput, error) 
 DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) 
 DescribeDeviceJob(ctx context.Context, params *DescribeDeviceJobInput, optFns ...func(*Options)) (*DescribeDeviceJobOutput, error) 
 DescribeNode(ctx context.Context, params *DescribeNodeInput, optFns ...func(*Options)) (*DescribeNodeOutput, error) 
 DescribeNodeFromTemplateJob(ctx context.Context, params *DescribeNodeFromTemplateJobInput, optFns ...func(*Options)) (*DescribeNodeFromTemplateJobOutput, error) 
 DescribePackage(ctx context.Context, params *DescribePackageInput, optFns ...func(*Options)) (*DescribePackageOutput, error) 
 DescribePackageImportJob(ctx context.Context, params *DescribePackageImportJobInput, optFns ...func(*Options)) (*DescribePackageImportJobOutput, error) 
 DescribePackageVersion(ctx context.Context, params *DescribePackageVersionInput, optFns ...func(*Options)) (*DescribePackageVersionOutput, error) 
 ListApplicationInstanceDependencies(ctx context.Context, params *ListApplicationInstanceDependenciesInput, optFns ...func(*Options)) (*ListApplicationInstanceDependenciesOutput, error) 
 ListApplicationInstanceNodeInstances(ctx context.Context, params *ListApplicationInstanceNodeInstancesInput, optFns ...func(*Options)) (*ListApplicationInstanceNodeInstancesOutput, error) 
 ListApplicationInstances(ctx context.Context, params *ListApplicationInstancesInput, optFns ...func(*Options)) (*ListApplicationInstancesOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListDevicesJobs(ctx context.Context, params *ListDevicesJobsInput, optFns ...func(*Options)) (*ListDevicesJobsOutput, error) 
 ListNodeFromTemplateJobs(ctx context.Context, params *ListNodeFromTemplateJobsInput, optFns ...func(*Options)) (*ListNodeFromTemplateJobsOutput, error) 
 ListNodes(ctx context.Context, params *ListNodesInput, optFns ...func(*Options)) (*ListNodesOutput, error) 
 ListPackageImportJobs(ctx context.Context, params *ListPackageImportJobsInput, optFns ...func(*Options)) (*ListPackageImportJobsOutput, error) 
 ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ProvisionDevice(ctx context.Context, params *ProvisionDeviceInput, optFns ...func(*Options)) (*ProvisionDeviceOutput, error) 
 RegisterPackageVersion(ctx context.Context, params *RegisterPackageVersionInput, optFns ...func(*Options)) (*RegisterPackageVersionOutput, error) 
 RemoveApplicationInstance(ctx context.Context, params *RemoveApplicationInstanceInput, optFns ...func(*Options)) (*RemoveApplicationInstanceOutput, error) 
 SignalApplicationInstanceNodeInstances(ctx context.Context, params *SignalApplicationInstanceNodeInstancesInput, optFns ...func(*Options)) (*SignalApplicationInstanceNodeInstancesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDeviceMetadata(ctx context.Context, params *UpdateDeviceMetadataInput, optFns ...func(*Options)) (*UpdateDeviceMetadataOutput, error) 
}
