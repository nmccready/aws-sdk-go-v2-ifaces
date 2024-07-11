
package mediapackagevod

import (
    "github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
)

// IMediapackagevod defines the interface for mediapackagevod
type IMediapackagevod interface {
 Options() Options 
 ConfigureLogs(ctx context.Context, params *ConfigureLogsInput, optFns ...func(*Options)) (*ConfigureLogsOutput, error) 
 CreateAsset(ctx context.Context, params *CreateAssetInput, optFns ...func(*Options)) (*CreateAssetOutput, error) 
 CreatePackagingConfiguration(ctx context.Context, params *CreatePackagingConfigurationInput, optFns ...func(*Options)) (*CreatePackagingConfigurationOutput, error) 
 CreatePackagingGroup(ctx context.Context, params *CreatePackagingGroupInput, optFns ...func(*Options)) (*CreatePackagingGroupOutput, error) 
 DeleteAsset(ctx context.Context, params *DeleteAssetInput, optFns ...func(*Options)) (*DeleteAssetOutput, error) 
 DeletePackagingConfiguration(ctx context.Context, params *DeletePackagingConfigurationInput, optFns ...func(*Options)) (*DeletePackagingConfigurationOutput, error) 
 DeletePackagingGroup(ctx context.Context, params *DeletePackagingGroupInput, optFns ...func(*Options)) (*DeletePackagingGroupOutput, error) 
 DescribeAsset(ctx context.Context, params *DescribeAssetInput, optFns ...func(*Options)) (*DescribeAssetOutput, error) 
 DescribePackagingConfiguration(ctx context.Context, params *DescribePackagingConfigurationInput, optFns ...func(*Options)) (*DescribePackagingConfigurationOutput, error) 
 DescribePackagingGroup(ctx context.Context, params *DescribePackagingGroupInput, optFns ...func(*Options)) (*DescribePackagingGroupOutput, error) 
 ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error) 
 ListPackagingConfigurations(ctx context.Context, params *ListPackagingConfigurationsInput, optFns ...func(*Options)) (*ListPackagingConfigurationsOutput, error) 
 ListPackagingGroups(ctx context.Context, params *ListPackagingGroupsInput, optFns ...func(*Options)) (*ListPackagingGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePackagingGroup(ctx context.Context, params *UpdatePackagingGroupInput, optFns ...func(*Options)) (*UpdatePackagingGroupOutput, error) 
}
