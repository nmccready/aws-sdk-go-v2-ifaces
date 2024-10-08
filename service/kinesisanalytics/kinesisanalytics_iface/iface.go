
package kinesisanalytics_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

// IClient defines the interface for kinesisanalytics
type IClient interface {
 Options() Options 
 AddApplicationCloudWatchLoggingOption(ctx context.Context, params *AddApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*AddApplicationCloudWatchLoggingOptionOutput, error) 
 AddApplicationInput(ctx context.Context, params *AddApplicationInputInput, optFns ...func(*Options)) (*AddApplicationInputOutput, error) 
 AddApplicationInputProcessingConfiguration(ctx context.Context, params *AddApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*AddApplicationInputProcessingConfigurationOutput, error) 
 AddApplicationOutput(ctx context.Context, params *AddApplicationOutputInput, optFns ...func(*Options)) (*AddApplicationOutputOutput, error) 
 AddApplicationReferenceDataSource(ctx context.Context, params *AddApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*AddApplicationReferenceDataSourceOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteApplicationCloudWatchLoggingOption(ctx context.Context, params *DeleteApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*DeleteApplicationCloudWatchLoggingOptionOutput, error) 
 DeleteApplicationInputProcessingConfiguration(ctx context.Context, params *DeleteApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationInputProcessingConfigurationOutput, error) 
 DeleteApplicationOutput(ctx context.Context, params *DeleteApplicationOutputInput, optFns ...func(*Options)) (*DeleteApplicationOutputOutput, error) 
 DeleteApplicationReferenceDataSource(ctx context.Context, params *DeleteApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*DeleteApplicationReferenceDataSourceOutput, error) 
 DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) 
 DiscoverInputSchema(ctx context.Context, params *DiscoverInputSchemaInput, optFns ...func(*Options)) (*DiscoverInputSchemaOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) 
 StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
}
