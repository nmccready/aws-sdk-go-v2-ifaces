
package kafkaconnect

import (
    "github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
)

// IClient defines the interface for kafkaconnect
type IClient interface {
 Options() Options 
 CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) 
 CreateCustomPlugin(ctx context.Context, params *CreateCustomPluginInput, optFns ...func(*Options)) (*CreateCustomPluginOutput, error) 
 CreateWorkerConfiguration(ctx context.Context, params *CreateWorkerConfigurationInput, optFns ...func(*Options)) (*CreateWorkerConfigurationOutput, error) 
 DeleteConnector(ctx context.Context, params *DeleteConnectorInput, optFns ...func(*Options)) (*DeleteConnectorOutput, error) 
 DeleteCustomPlugin(ctx context.Context, params *DeleteCustomPluginInput, optFns ...func(*Options)) (*DeleteCustomPluginOutput, error) 
 DeleteWorkerConfiguration(ctx context.Context, params *DeleteWorkerConfigurationInput, optFns ...func(*Options)) (*DeleteWorkerConfigurationOutput, error) 
 DescribeConnector(ctx context.Context, params *DescribeConnectorInput, optFns ...func(*Options)) (*DescribeConnectorOutput, error) 
 DescribeCustomPlugin(ctx context.Context, params *DescribeCustomPluginInput, optFns ...func(*Options)) (*DescribeCustomPluginOutput, error) 
 DescribeWorkerConfiguration(ctx context.Context, params *DescribeWorkerConfigurationInput, optFns ...func(*Options)) (*DescribeWorkerConfigurationOutput, error) 
 ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) 
 ListCustomPlugins(ctx context.Context, params *ListCustomPluginsInput, optFns ...func(*Options)) (*ListCustomPluginsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkerConfigurations(ctx context.Context, params *ListWorkerConfigurationsInput, optFns ...func(*Options)) (*ListWorkerConfigurationsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConnector(ctx context.Context, params *UpdateConnectorInput, optFns ...func(*Options)) (*UpdateConnectorOutput, error) 
}
