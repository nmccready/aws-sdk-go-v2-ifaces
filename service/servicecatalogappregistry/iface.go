
package servicecatalogappregistry

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
)

// IClient defines the interface for servicecatalogappregistry
type IClient interface {
 Options() Options 
 AssociateAttributeGroup(ctx context.Context, params *AssociateAttributeGroupInput, optFns ...func(*Options)) (*AssociateAttributeGroupOutput, error) 
 AssociateResource(ctx context.Context, params *AssociateResourceInput, optFns ...func(*Options)) (*AssociateResourceOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateAttributeGroup(ctx context.Context, params *CreateAttributeGroupInput, optFns ...func(*Options)) (*CreateAttributeGroupOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteAttributeGroup(ctx context.Context, params *DeleteAttributeGroupInput, optFns ...func(*Options)) (*DeleteAttributeGroupOutput, error) 
 DisassociateAttributeGroup(ctx context.Context, params *DisassociateAttributeGroupInput, optFns ...func(*Options)) (*DisassociateAttributeGroupOutput, error) 
 DisassociateResource(ctx context.Context, params *DisassociateResourceInput, optFns ...func(*Options)) (*DisassociateResourceOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetAssociatedResource(ctx context.Context, params *GetAssociatedResourceInput, optFns ...func(*Options)) (*GetAssociatedResourceOutput, error) 
 GetAttributeGroup(ctx context.Context, params *GetAttributeGroupInput, optFns ...func(*Options)) (*GetAttributeGroupOutput, error) 
 GetConfiguration(ctx context.Context, params *GetConfigurationInput, optFns ...func(*Options)) (*GetConfigurationOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListAssociatedAttributeGroups(ctx context.Context, params *ListAssociatedAttributeGroupsInput, optFns ...func(*Options)) (*ListAssociatedAttributeGroupsOutput, error) 
 ListAssociatedResources(ctx context.Context, params *ListAssociatedResourcesInput, optFns ...func(*Options)) (*ListAssociatedResourcesOutput, error) 
 ListAttributeGroups(ctx context.Context, params *ListAttributeGroupsInput, optFns ...func(*Options)) (*ListAttributeGroupsOutput, error) 
 ListAttributeGroupsForApplication(ctx context.Context, params *ListAttributeGroupsForApplicationInput, optFns ...func(*Options)) (*ListAttributeGroupsForApplicationOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutConfiguration(ctx context.Context, params *PutConfigurationInput, optFns ...func(*Options)) (*PutConfigurationOutput, error) 
 SyncResource(ctx context.Context, params *SyncResourceInput, optFns ...func(*Options)) (*SyncResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateAttributeGroup(ctx context.Context, params *UpdateAttributeGroupInput, optFns ...func(*Options)) (*UpdateAttributeGroupOutput, error) 
}
