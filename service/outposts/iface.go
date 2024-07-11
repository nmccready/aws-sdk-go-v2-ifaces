
package outposts

import (
    "github.com/aws/aws-sdk-go-v2/service/outposts"
)

// IOutposts defines the interface for outposts
type IOutposts interface {
 Options() Options 
 CancelCapacityTask(ctx context.Context, params *CancelCapacityTaskInput, optFns ...func(*Options)) (*CancelCapacityTaskOutput, error) 
 CancelOrder(ctx context.Context, params *CancelOrderInput, optFns ...func(*Options)) (*CancelOrderOutput, error) 
 CreateOrder(ctx context.Context, params *CreateOrderInput, optFns ...func(*Options)) (*CreateOrderOutput, error) 
 CreateOutpost(ctx context.Context, params *CreateOutpostInput, optFns ...func(*Options)) (*CreateOutpostOutput, error) 
 CreateSite(ctx context.Context, params *CreateSiteInput, optFns ...func(*Options)) (*CreateSiteOutput, error) 
 DeleteOutpost(ctx context.Context, params *DeleteOutpostInput, optFns ...func(*Options)) (*DeleteOutpostOutput, error) 
 DeleteSite(ctx context.Context, params *DeleteSiteInput, optFns ...func(*Options)) (*DeleteSiteOutput, error) 
 GetCapacityTask(ctx context.Context, params *GetCapacityTaskInput, optFns ...func(*Options)) (*GetCapacityTaskOutput, error) 
 GetCatalogItem(ctx context.Context, params *GetCatalogItemInput, optFns ...func(*Options)) (*GetCatalogItemOutput, error) 
 GetConnection(ctx context.Context, params *GetConnectionInput, optFns ...func(*Options)) (*GetConnectionOutput, error) 
 GetOrder(ctx context.Context, params *GetOrderInput, optFns ...func(*Options)) (*GetOrderOutput, error) 
 GetOutpost(ctx context.Context, params *GetOutpostInput, optFns ...func(*Options)) (*GetOutpostOutput, error) 
 GetOutpostInstanceTypes(ctx context.Context, params *GetOutpostInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error) 
 GetOutpostSupportedInstanceTypes(ctx context.Context, params *GetOutpostSupportedInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostSupportedInstanceTypesOutput, error) 
 GetSite(ctx context.Context, params *GetSiteInput, optFns ...func(*Options)) (*GetSiteOutput, error) 
 GetSiteAddress(ctx context.Context, params *GetSiteAddressInput, optFns ...func(*Options)) (*GetSiteAddressOutput, error) 
 ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error) 
 ListCapacityTasks(ctx context.Context, params *ListCapacityTasksInput, optFns ...func(*Options)) (*ListCapacityTasksOutput, error) 
 ListCatalogItems(ctx context.Context, params *ListCatalogItemsInput, optFns ...func(*Options)) (*ListCatalogItemsOutput, error) 
 ListOrders(ctx context.Context, params *ListOrdersInput, optFns ...func(*Options)) (*ListOrdersOutput, error) 
 ListOutposts(ctx context.Context, params *ListOutpostsInput, optFns ...func(*Options)) (*ListOutpostsOutput, error) 
 ListSites(ctx context.Context, params *ListSitesInput, optFns ...func(*Options)) (*ListSitesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartCapacityTask(ctx context.Context, params *StartCapacityTaskInput, optFns ...func(*Options)) (*StartCapacityTaskOutput, error) 
 StartConnection(ctx context.Context, params *StartConnectionInput, optFns ...func(*Options)) (*StartConnectionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateOutpost(ctx context.Context, params *UpdateOutpostInput, optFns ...func(*Options)) (*UpdateOutpostOutput, error) 
 UpdateSite(ctx context.Context, params *UpdateSiteInput, optFns ...func(*Options)) (*UpdateSiteOutput, error) 
 UpdateSiteAddress(ctx context.Context, params *UpdateSiteAddressInput, optFns ...func(*Options)) (*UpdateSiteAddressOutput, error) 
 UpdateSiteRackPhysicalProperties(ctx context.Context, params *UpdateSiteRackPhysicalPropertiesInput, optFns ...func(*Options)) (*UpdateSiteRackPhysicalPropertiesOutput, error) 
}
