
package dax

import (
    "github.com/aws/aws-sdk-go-v2/service/dax"
)

// IDax defines the interface for dax
type IDax interface {
 Options() Options 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateParameterGroup(ctx context.Context, params *CreateParameterGroupInput, optFns ...func(*Options)) (*CreateParameterGroupOutput, error) 
 CreateSubnetGroup(ctx context.Context, params *CreateSubnetGroupInput, optFns ...func(*Options)) (*CreateSubnetGroupOutput, error) 
 DecreaseReplicationFactor(ctx context.Context, params *DecreaseReplicationFactorInput, optFns ...func(*Options)) (*DecreaseReplicationFactorOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteParameterGroup(ctx context.Context, params *DeleteParameterGroupInput, optFns ...func(*Options)) (*DeleteParameterGroupOutput, error) 
 DeleteSubnetGroup(ctx context.Context, params *DeleteSubnetGroupInput, optFns ...func(*Options)) (*DeleteSubnetGroupOutput, error) 
 DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) 
 DescribeDefaultParameters(ctx context.Context, params *DescribeDefaultParametersInput, optFns ...func(*Options)) (*DescribeDefaultParametersOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeParameterGroups(ctx context.Context, params *DescribeParameterGroupsInput, optFns ...func(*Options)) (*DescribeParameterGroupsOutput, error) 
 DescribeParameters(ctx context.Context, params *DescribeParametersInput, optFns ...func(*Options)) (*DescribeParametersOutput, error) 
 DescribeSubnetGroups(ctx context.Context, params *DescribeSubnetGroupsInput, optFns ...func(*Options)) (*DescribeSubnetGroupsOutput, error) 
 IncreaseReplicationFactor(ctx context.Context, params *IncreaseReplicationFactorInput, optFns ...func(*Options)) (*IncreaseReplicationFactorOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 RebootNode(ctx context.Context, params *RebootNodeInput, optFns ...func(*Options)) (*RebootNodeOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) 
 UpdateParameterGroup(ctx context.Context, params *UpdateParameterGroupInput, optFns ...func(*Options)) (*UpdateParameterGroupOutput, error) 
 UpdateSubnetGroup(ctx context.Context, params *UpdateSubnetGroupInput, optFns ...func(*Options)) (*UpdateSubnetGroupOutput, error) 
}
