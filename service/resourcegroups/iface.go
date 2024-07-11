
package resourcegroups

import (
    "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
)

// IClient defines the interface for resourcegroups
type IClient interface {
 Options() Options 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetGroupConfiguration(ctx context.Context, params *GetGroupConfigurationInput, optFns ...func(*Options)) (*GetGroupConfigurationOutput, error) 
 GetGroupQuery(ctx context.Context, params *GetGroupQueryInput, optFns ...func(*Options)) (*GetGroupQueryOutput, error) 
 GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) 
 GroupResources(ctx context.Context, params *GroupResourcesInput, optFns ...func(*Options)) (*GroupResourcesOutput, error) 
 ListGroupResources(ctx context.Context, params *ListGroupResourcesInput, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 PutGroupConfiguration(ctx context.Context, params *PutGroupConfigurationInput, optFns ...func(*Options)) (*PutGroupConfigurationOutput, error) 
 SearchResources(ctx context.Context, params *SearchResourcesInput, optFns ...func(*Options)) (*SearchResourcesOutput, error) 
 Tag(ctx context.Context, params *TagInput, optFns ...func(*Options)) (*TagOutput, error) 
 UngroupResources(ctx context.Context, params *UngroupResourcesInput, optFns ...func(*Options)) (*UngroupResourcesOutput, error) 
 Untag(ctx context.Context, params *UntagInput, optFns ...func(*Options)) (*UntagOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateGroupQuery(ctx context.Context, params *UpdateGroupQueryInput, optFns ...func(*Options)) (*UpdateGroupQueryOutput, error) 
}
