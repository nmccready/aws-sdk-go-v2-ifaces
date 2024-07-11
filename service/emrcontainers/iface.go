
package emrcontainers

import (
    "github.com/aws/aws-sdk-go-v2/service/emrcontainers"
)

// IEmrcontainers defines the interface for emrcontainers
type IEmrcontainers interface {
 Options() Options 
 CancelJobRun(ctx context.Context, params *CancelJobRunInput, optFns ...func(*Options)) (*CancelJobRunOutput, error) 
 CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) 
 CreateManagedEndpoint(ctx context.Context, params *CreateManagedEndpointInput, optFns ...func(*Options)) (*CreateManagedEndpointOutput, error) 
 CreateSecurityConfiguration(ctx context.Context, params *CreateSecurityConfigurationInput, optFns ...func(*Options)) (*CreateSecurityConfigurationOutput, error) 
 CreateVirtualCluster(ctx context.Context, params *CreateVirtualClusterInput, optFns ...func(*Options)) (*CreateVirtualClusterOutput, error) 
 DeleteJobTemplate(ctx context.Context, params *DeleteJobTemplateInput, optFns ...func(*Options)) (*DeleteJobTemplateOutput, error) 
 DeleteManagedEndpoint(ctx context.Context, params *DeleteManagedEndpointInput, optFns ...func(*Options)) (*DeleteManagedEndpointOutput, error) 
 DeleteVirtualCluster(ctx context.Context, params *DeleteVirtualClusterInput, optFns ...func(*Options)) (*DeleteVirtualClusterOutput, error) 
 DescribeJobRun(ctx context.Context, params *DescribeJobRunInput, optFns ...func(*Options)) (*DescribeJobRunOutput, error) 
 DescribeJobTemplate(ctx context.Context, params *DescribeJobTemplateInput, optFns ...func(*Options)) (*DescribeJobTemplateOutput, error) 
 DescribeManagedEndpoint(ctx context.Context, params *DescribeManagedEndpointInput, optFns ...func(*Options)) (*DescribeManagedEndpointOutput, error) 
 DescribeSecurityConfiguration(ctx context.Context, params *DescribeSecurityConfigurationInput, optFns ...func(*Options)) (*DescribeSecurityConfigurationOutput, error) 
 DescribeVirtualCluster(ctx context.Context, params *DescribeVirtualClusterInput, optFns ...func(*Options)) (*DescribeVirtualClusterOutput, error) 
 GetManagedEndpointSessionCredentials(ctx context.Context, params *GetManagedEndpointSessionCredentialsInput, optFns ...func(*Options)) (*GetManagedEndpointSessionCredentialsOutput, error) 
 ListJobRuns(ctx context.Context, params *ListJobRunsInput, optFns ...func(*Options)) (*ListJobRunsOutput, error) 
 ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) 
 ListManagedEndpoints(ctx context.Context, params *ListManagedEndpointsInput, optFns ...func(*Options)) (*ListManagedEndpointsOutput, error) 
 ListSecurityConfigurations(ctx context.Context, params *ListSecurityConfigurationsInput, optFns ...func(*Options)) (*ListSecurityConfigurationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVirtualClusters(ctx context.Context, params *ListVirtualClustersInput, optFns ...func(*Options)) (*ListVirtualClustersOutput, error) 
 StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
