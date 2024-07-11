
package kinesisanalyticsv2

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

// IClient defines the interface for kinesisanalyticsv2
type IClient interface {
 Options() Options 
 AddApplicationCloudWatchLoggingOption(ctx context.Context, params *AddApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*AddApplicationCloudWatchLoggingOptionOutput, error) 
 AddApplicationInput(ctx context.Context, params *AddApplicationInputInput, optFns ...func(*Options)) (*AddApplicationInputOutput, error) 
 AddApplicationInputProcessingConfiguration(ctx context.Context, params *AddApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*AddApplicationInputProcessingConfigurationOutput, error) 
 AddApplicationOutput(ctx context.Context, params *AddApplicationOutputInput, optFns ...func(*Options)) (*AddApplicationOutputOutput, error) 
 AddApplicationReferenceDataSource(ctx context.Context, params *AddApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*AddApplicationReferenceDataSourceOutput, error) 
 AddApplicationVpcConfiguration(ctx context.Context, params *AddApplicationVpcConfigurationInput, optFns ...func(*Options)) (*AddApplicationVpcConfigurationOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateApplicationPresignedUrl(ctx context.Context, params *CreateApplicationPresignedUrlInput, optFns ...func(*Options)) (*CreateApplicationPresignedUrlOutput, error) 
 CreateApplicationSnapshot(ctx context.Context, params *CreateApplicationSnapshotInput, optFns ...func(*Options)) (*CreateApplicationSnapshotOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteApplicationCloudWatchLoggingOption(ctx context.Context, params *DeleteApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*DeleteApplicationCloudWatchLoggingOptionOutput, error) 
 DeleteApplicationInputProcessingConfiguration(ctx context.Context, params *DeleteApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationInputProcessingConfigurationOutput, error) 
 DeleteApplicationOutput(ctx context.Context, params *DeleteApplicationOutputInput, optFns ...func(*Options)) (*DeleteApplicationOutputOutput, error) 
 DeleteApplicationReferenceDataSource(ctx context.Context, params *DeleteApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*DeleteApplicationReferenceDataSourceOutput, error) 
 DeleteApplicationSnapshot(ctx context.Context, params *DeleteApplicationSnapshotInput, optFns ...func(*Options)) (*DeleteApplicationSnapshotOutput, error) 
 DeleteApplicationVpcConfiguration(ctx context.Context, params *DeleteApplicationVpcConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationVpcConfigurationOutput, error) 
 DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) 
 DescribeApplicationOperation(ctx context.Context, params *DescribeApplicationOperationInput, optFns ...func(*Options)) (*DescribeApplicationOperationOutput, error) 
 DescribeApplicationSnapshot(ctx context.Context, params *DescribeApplicationSnapshotInput, optFns ...func(*Options)) (*DescribeApplicationSnapshotOutput, error) 
 DescribeApplicationVersion(ctx context.Context, params *DescribeApplicationVersionInput, optFns ...func(*Options)) (*DescribeApplicationVersionOutput, error) 
 DiscoverInputSchema(ctx context.Context, params *DiscoverInputSchemaInput, optFns ...func(*Options)) (*DiscoverInputSchemaOutput, error) 
 ListApplicationOperations(ctx context.Context, params *ListApplicationOperationsInput, optFns ...func(*Options)) (*ListApplicationOperationsOutput, error) 
 ListApplicationSnapshots(ctx context.Context, params *ListApplicationSnapshotsInput, optFns ...func(*Options)) (*ListApplicationSnapshotsOutput, error) 
 ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RollbackApplication(ctx context.Context, params *RollbackApplicationInput, optFns ...func(*Options)) (*RollbackApplicationOutput, error) 
 StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) 
 StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateApplicationMaintenanceConfiguration(ctx context.Context, params *UpdateApplicationMaintenanceConfigurationInput, optFns ...func(*Options)) (*UpdateApplicationMaintenanceConfigurationOutput, error) 
}
