
package iottwinmaker

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/iottwinmaker"
)

// IClient defines the interface for iottwinmaker
type IClient interface {
 Options() Options 
 BatchPutPropertyValues(ctx context.Context, params *BatchPutPropertyValuesInput, optFns ...func(*Options)) (*BatchPutPropertyValuesOutput, error) 
 CancelMetadataTransferJob(ctx context.Context, params *CancelMetadataTransferJobInput, optFns ...func(*Options)) (*CancelMetadataTransferJobOutput, error) 
 CreateComponentType(ctx context.Context, params *CreateComponentTypeInput, optFns ...func(*Options)) (*CreateComponentTypeOutput, error) 
 CreateEntity(ctx context.Context, params *CreateEntityInput, optFns ...func(*Options)) (*CreateEntityOutput, error) 
 CreateMetadataTransferJob(ctx context.Context, params *CreateMetadataTransferJobInput, optFns ...func(*Options)) (*CreateMetadataTransferJobOutput, error) 
 CreateScene(ctx context.Context, params *CreateSceneInput, optFns ...func(*Options)) (*CreateSceneOutput, error) 
 CreateSyncJob(ctx context.Context, params *CreateSyncJobInput, optFns ...func(*Options)) (*CreateSyncJobOutput, error) 
 CreateWorkspace(ctx context.Context, params *CreateWorkspaceInput, optFns ...func(*Options)) (*CreateWorkspaceOutput, error) 
 DeleteComponentType(ctx context.Context, params *DeleteComponentTypeInput, optFns ...func(*Options)) (*DeleteComponentTypeOutput, error) 
 DeleteEntity(ctx context.Context, params *DeleteEntityInput, optFns ...func(*Options)) (*DeleteEntityOutput, error) 
 DeleteScene(ctx context.Context, params *DeleteSceneInput, optFns ...func(*Options)) (*DeleteSceneOutput, error) 
 DeleteSyncJob(ctx context.Context, params *DeleteSyncJobInput, optFns ...func(*Options)) (*DeleteSyncJobOutput, error) 
 DeleteWorkspace(ctx context.Context, params *DeleteWorkspaceInput, optFns ...func(*Options)) (*DeleteWorkspaceOutput, error) 
 ExecuteQuery(ctx context.Context, params *ExecuteQueryInput, optFns ...func(*Options)) (*ExecuteQueryOutput, error) 
 GetComponentType(ctx context.Context, params *GetComponentTypeInput, optFns ...func(*Options)) (*GetComponentTypeOutput, error) 
 GetEntity(ctx context.Context, params *GetEntityInput, optFns ...func(*Options)) (*GetEntityOutput, error) 
 GetMetadataTransferJob(ctx context.Context, params *GetMetadataTransferJobInput, optFns ...func(*Options)) (*GetMetadataTransferJobOutput, error) 
 GetPricingPlan(ctx context.Context, params *GetPricingPlanInput, optFns ...func(*Options)) (*GetPricingPlanOutput, error) 
 GetPropertyValue(ctx context.Context, params *GetPropertyValueInput, optFns ...func(*Options)) (*GetPropertyValueOutput, error) 
 GetPropertyValueHistory(ctx context.Context, params *GetPropertyValueHistoryInput, optFns ...func(*Options)) (*GetPropertyValueHistoryOutput, error) 
 GetScene(ctx context.Context, params *GetSceneInput, optFns ...func(*Options)) (*GetSceneOutput, error) 
 GetSyncJob(ctx context.Context, params *GetSyncJobInput, optFns ...func(*Options)) (*GetSyncJobOutput, error) 
 GetWorkspace(ctx context.Context, params *GetWorkspaceInput, optFns ...func(*Options)) (*GetWorkspaceOutput, error) 
 ListComponentTypes(ctx context.Context, params *ListComponentTypesInput, optFns ...func(*Options)) (*ListComponentTypesOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListEntities(ctx context.Context, params *ListEntitiesInput, optFns ...func(*Options)) (*ListEntitiesOutput, error) 
 ListMetadataTransferJobs(ctx context.Context, params *ListMetadataTransferJobsInput, optFns ...func(*Options)) (*ListMetadataTransferJobsOutput, error) 
 ListProperties(ctx context.Context, params *ListPropertiesInput, optFns ...func(*Options)) (*ListPropertiesOutput, error) 
 ListScenes(ctx context.Context, params *ListScenesInput, optFns ...func(*Options)) (*ListScenesOutput, error) 
 ListSyncJobs(ctx context.Context, params *ListSyncJobsInput, optFns ...func(*Options)) (*ListSyncJobsOutput, error) 
 ListSyncResources(ctx context.Context, params *ListSyncResourcesInput, optFns ...func(*Options)) (*ListSyncResourcesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkspaces(ctx context.Context, params *ListWorkspacesInput, optFns ...func(*Options)) (*ListWorkspacesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateComponentType(ctx context.Context, params *UpdateComponentTypeInput, optFns ...func(*Options)) (*UpdateComponentTypeOutput, error) 
 UpdateEntity(ctx context.Context, params *UpdateEntityInput, optFns ...func(*Options)) (*UpdateEntityOutput, error) 
 UpdatePricingPlan(ctx context.Context, params *UpdatePricingPlanInput, optFns ...func(*Options)) (*UpdatePricingPlanOutput, error) 
 UpdateScene(ctx context.Context, params *UpdateSceneInput, optFns ...func(*Options)) (*UpdateSceneOutput, error) 
 UpdateWorkspace(ctx context.Context, params *UpdateWorkspaceInput, optFns ...func(*Options)) (*UpdateWorkspaceOutput, error) 
}
