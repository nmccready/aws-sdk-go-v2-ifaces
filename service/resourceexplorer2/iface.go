
package resourceexplorer2

import (
    "github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
)

// IResourceexplorer2 defines the interface for resourceexplorer2
type IResourceexplorer2 interface {
 Options() Options 
 AssociateDefaultView(ctx context.Context, params *AssociateDefaultViewInput, optFns ...func(*Options)) (*AssociateDefaultViewOutput, error) 
 BatchGetView(ctx context.Context, params *BatchGetViewInput, optFns ...func(*Options)) (*BatchGetViewOutput, error) 
 CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) 
 CreateView(ctx context.Context, params *CreateViewInput, optFns ...func(*Options)) (*CreateViewOutput, error) 
 DeleteIndex(ctx context.Context, params *DeleteIndexInput, optFns ...func(*Options)) (*DeleteIndexOutput, error) 
 DeleteView(ctx context.Context, params *DeleteViewInput, optFns ...func(*Options)) (*DeleteViewOutput, error) 
 DisassociateDefaultView(ctx context.Context, params *DisassociateDefaultViewInput, optFns ...func(*Options)) (*DisassociateDefaultViewOutput, error) 
 GetAccountLevelServiceConfiguration(ctx context.Context, params *GetAccountLevelServiceConfigurationInput, optFns ...func(*Options)) (*GetAccountLevelServiceConfigurationOutput, error) 
 GetDefaultView(ctx context.Context, params *GetDefaultViewInput, optFns ...func(*Options)) (*GetDefaultViewOutput, error) 
 GetIndex(ctx context.Context, params *GetIndexInput, optFns ...func(*Options)) (*GetIndexOutput, error) 
 GetView(ctx context.Context, params *GetViewInput, optFns ...func(*Options)) (*GetViewOutput, error) 
 ListIndexes(ctx context.Context, params *ListIndexesInput, optFns ...func(*Options)) (*ListIndexesOutput, error) 
 ListIndexesForMembers(ctx context.Context, params *ListIndexesForMembersInput, optFns ...func(*Options)) (*ListIndexesForMembersOutput, error) 
 ListSupportedResourceTypes(ctx context.Context, params *ListSupportedResourceTypesInput, optFns ...func(*Options)) (*ListSupportedResourceTypesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListViews(ctx context.Context, params *ListViewsInput, optFns ...func(*Options)) (*ListViewsOutput, error) 
 Search(ctx context.Context, params *SearchInput, optFns ...func(*Options)) (*SearchOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateIndexType(ctx context.Context, params *UpdateIndexTypeInput, optFns ...func(*Options)) (*UpdateIndexTypeOutput, error) 
 UpdateView(ctx context.Context, params *UpdateViewInput, optFns ...func(*Options)) (*UpdateViewOutput, error) 
}
