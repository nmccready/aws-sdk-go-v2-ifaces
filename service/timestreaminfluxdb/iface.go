
package timestreaminfluxdb

import (
    "github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb"
)

// ITimestreaminfluxdb defines the interface for timestreaminfluxdb
type ITimestreaminfluxdb interface {
 Options() Options 
 CreateDbInstance(ctx context.Context, params *CreateDbInstanceInput, optFns ...func(*Options)) (*CreateDbInstanceOutput, error) 
 CreateDbParameterGroup(ctx context.Context, params *CreateDbParameterGroupInput, optFns ...func(*Options)) (*CreateDbParameterGroupOutput, error) 
 DeleteDbInstance(ctx context.Context, params *DeleteDbInstanceInput, optFns ...func(*Options)) (*DeleteDbInstanceOutput, error) 
 GetDbInstance(ctx context.Context, params *GetDbInstanceInput, optFns ...func(*Options)) (*GetDbInstanceOutput, error) 
 GetDbParameterGroup(ctx context.Context, params *GetDbParameterGroupInput, optFns ...func(*Options)) (*GetDbParameterGroupOutput, error) 
 ListDbInstances(ctx context.Context, params *ListDbInstancesInput, optFns ...func(*Options)) (*ListDbInstancesOutput, error) 
 ListDbParameterGroups(ctx context.Context, params *ListDbParameterGroupsInput, optFns ...func(*Options)) (*ListDbParameterGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDbInstance(ctx context.Context, params *UpdateDbInstanceInput, optFns ...func(*Options)) (*UpdateDbInstanceOutput, error) 
}
