
package clouddirectory_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/clouddirectory"
)

// IClient defines the interface for clouddirectory
type IClient interface {
 Options() Options 
 AddFacetToObject(ctx context.Context, params *AddFacetToObjectInput, optFns ...func(*Options)) (*AddFacetToObjectOutput, error) 
 ApplySchema(ctx context.Context, params *ApplySchemaInput, optFns ...func(*Options)) (*ApplySchemaOutput, error) 
 AttachObject(ctx context.Context, params *AttachObjectInput, optFns ...func(*Options)) (*AttachObjectOutput, error) 
 AttachPolicy(ctx context.Context, params *AttachPolicyInput, optFns ...func(*Options)) (*AttachPolicyOutput, error) 
 AttachToIndex(ctx context.Context, params *AttachToIndexInput, optFns ...func(*Options)) (*AttachToIndexOutput, error) 
 AttachTypedLink(ctx context.Context, params *AttachTypedLinkInput, optFns ...func(*Options)) (*AttachTypedLinkOutput, error) 
 BatchRead(ctx context.Context, params *BatchReadInput, optFns ...func(*Options)) (*BatchReadOutput, error) 
 BatchWrite(ctx context.Context, params *BatchWriteInput, optFns ...func(*Options)) (*BatchWriteOutput, error) 
 CreateDirectory(ctx context.Context, params *CreateDirectoryInput, optFns ...func(*Options)) (*CreateDirectoryOutput, error) 
 CreateFacet(ctx context.Context, params *CreateFacetInput, optFns ...func(*Options)) (*CreateFacetOutput, error) 
 CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) 
 CreateObject(ctx context.Context, params *CreateObjectInput, optFns ...func(*Options)) (*CreateObjectOutput, error) 
 CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) 
 CreateTypedLinkFacet(ctx context.Context, params *CreateTypedLinkFacetInput, optFns ...func(*Options)) (*CreateTypedLinkFacetOutput, error) 
 DeleteDirectory(ctx context.Context, params *DeleteDirectoryInput, optFns ...func(*Options)) (*DeleteDirectoryOutput, error) 
 DeleteFacet(ctx context.Context, params *DeleteFacetInput, optFns ...func(*Options)) (*DeleteFacetOutput, error) 
 DeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*Options)) (*DeleteObjectOutput, error) 
 DeleteSchema(ctx context.Context, params *DeleteSchemaInput, optFns ...func(*Options)) (*DeleteSchemaOutput, error) 
 DeleteTypedLinkFacet(ctx context.Context, params *DeleteTypedLinkFacetInput, optFns ...func(*Options)) (*DeleteTypedLinkFacetOutput, error) 
 DetachFromIndex(ctx context.Context, params *DetachFromIndexInput, optFns ...func(*Options)) (*DetachFromIndexOutput, error) 
 DetachObject(ctx context.Context, params *DetachObjectInput, optFns ...func(*Options)) (*DetachObjectOutput, error) 
 DetachPolicy(ctx context.Context, params *DetachPolicyInput, optFns ...func(*Options)) (*DetachPolicyOutput, error) 
 DetachTypedLink(ctx context.Context, params *DetachTypedLinkInput, optFns ...func(*Options)) (*DetachTypedLinkOutput, error) 
 DisableDirectory(ctx context.Context, params *DisableDirectoryInput, optFns ...func(*Options)) (*DisableDirectoryOutput, error) 
 EnableDirectory(ctx context.Context, params *EnableDirectoryInput, optFns ...func(*Options)) (*EnableDirectoryOutput, error) 
 GetAppliedSchemaVersion(ctx context.Context, params *GetAppliedSchemaVersionInput, optFns ...func(*Options)) (*GetAppliedSchemaVersionOutput, error) 
 GetDirectory(ctx context.Context, params *GetDirectoryInput, optFns ...func(*Options)) (*GetDirectoryOutput, error) 
 GetFacet(ctx context.Context, params *GetFacetInput, optFns ...func(*Options)) (*GetFacetOutput, error) 
 GetLinkAttributes(ctx context.Context, params *GetLinkAttributesInput, optFns ...func(*Options)) (*GetLinkAttributesOutput, error) 
 GetObjectAttributes(ctx context.Context, params *GetObjectAttributesInput, optFns ...func(*Options)) (*GetObjectAttributesOutput, error) 
 GetObjectInformation(ctx context.Context, params *GetObjectInformationInput, optFns ...func(*Options)) (*GetObjectInformationOutput, error) 
 GetSchemaAsJson(ctx context.Context, params *GetSchemaAsJsonInput, optFns ...func(*Options)) (*GetSchemaAsJsonOutput, error) 
 GetTypedLinkFacetInformation(ctx context.Context, params *GetTypedLinkFacetInformationInput, optFns ...func(*Options)) (*GetTypedLinkFacetInformationOutput, error) 
 ListAppliedSchemaArns(ctx context.Context, params *ListAppliedSchemaArnsInput, optFns ...func(*Options)) (*ListAppliedSchemaArnsOutput, error) 
 ListAttachedIndices(ctx context.Context, params *ListAttachedIndicesInput, optFns ...func(*Options)) (*ListAttachedIndicesOutput, error) 
 ListDevelopmentSchemaArns(ctx context.Context, params *ListDevelopmentSchemaArnsInput, optFns ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error) 
 ListDirectories(ctx context.Context, params *ListDirectoriesInput, optFns ...func(*Options)) (*ListDirectoriesOutput, error) 
 ListFacetAttributes(ctx context.Context, params *ListFacetAttributesInput, optFns ...func(*Options)) (*ListFacetAttributesOutput, error) 
 ListFacetNames(ctx context.Context, params *ListFacetNamesInput, optFns ...func(*Options)) (*ListFacetNamesOutput, error) 
 ListIncomingTypedLinks(ctx context.Context, params *ListIncomingTypedLinksInput, optFns ...func(*Options)) (*ListIncomingTypedLinksOutput, error) 
 ListIndex(ctx context.Context, params *ListIndexInput, optFns ...func(*Options)) (*ListIndexOutput, error) 
 ListManagedSchemaArns(ctx context.Context, params *ListManagedSchemaArnsInput, optFns ...func(*Options)) (*ListManagedSchemaArnsOutput, error) 
 ListObjectAttributes(ctx context.Context, params *ListObjectAttributesInput, optFns ...func(*Options)) (*ListObjectAttributesOutput, error) 
 ListObjectChildren(ctx context.Context, params *ListObjectChildrenInput, optFns ...func(*Options)) (*ListObjectChildrenOutput, error) 
 ListObjectParentPaths(ctx context.Context, params *ListObjectParentPathsInput, optFns ...func(*Options)) (*ListObjectParentPathsOutput, error) 
 ListObjectParents(ctx context.Context, params *ListObjectParentsInput, optFns ...func(*Options)) (*ListObjectParentsOutput, error) 
 ListObjectPolicies(ctx context.Context, params *ListObjectPoliciesInput, optFns ...func(*Options)) (*ListObjectPoliciesOutput, error) 
 ListOutgoingTypedLinks(ctx context.Context, params *ListOutgoingTypedLinksInput, optFns ...func(*Options)) (*ListOutgoingTypedLinksOutput, error) 
 ListPolicyAttachments(ctx context.Context, params *ListPolicyAttachmentsInput, optFns ...func(*Options)) (*ListPolicyAttachmentsOutput, error) 
 ListPublishedSchemaArns(ctx context.Context, params *ListPublishedSchemaArnsInput, optFns ...func(*Options)) (*ListPublishedSchemaArnsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTypedLinkFacetAttributes(ctx context.Context, params *ListTypedLinkFacetAttributesInput, optFns ...func(*Options)) (*ListTypedLinkFacetAttributesOutput, error) 
 ListTypedLinkFacetNames(ctx context.Context, params *ListTypedLinkFacetNamesInput, optFns ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error) 
 LookupPolicy(ctx context.Context, params *LookupPolicyInput, optFns ...func(*Options)) (*LookupPolicyOutput, error) 
 PublishSchema(ctx context.Context, params *PublishSchemaInput, optFns ...func(*Options)) (*PublishSchemaOutput, error) 
 PutSchemaFromJson(ctx context.Context, params *PutSchemaFromJsonInput, optFns ...func(*Options)) (*PutSchemaFromJsonOutput, error) 
 RemoveFacetFromObject(ctx context.Context, params *RemoveFacetFromObjectInput, optFns ...func(*Options)) (*RemoveFacetFromObjectOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateFacet(ctx context.Context, params *UpdateFacetInput, optFns ...func(*Options)) (*UpdateFacetOutput, error) 
 UpdateLinkAttributes(ctx context.Context, params *UpdateLinkAttributesInput, optFns ...func(*Options)) (*UpdateLinkAttributesOutput, error) 
 UpdateObjectAttributes(ctx context.Context, params *UpdateObjectAttributesInput, optFns ...func(*Options)) (*UpdateObjectAttributesOutput, error) 
 UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) 
 UpdateTypedLinkFacet(ctx context.Context, params *UpdateTypedLinkFacetInput, optFns ...func(*Options)) (*UpdateTypedLinkFacetOutput, error) 
 UpgradeAppliedSchema(ctx context.Context, params *UpgradeAppliedSchemaInput, optFns ...func(*Options)) (*UpgradeAppliedSchemaOutput, error) 
 UpgradePublishedSchema(ctx context.Context, params *UpgradePublishedSchemaInput, optFns ...func(*Options)) (*UpgradePublishedSchemaOutput, error) 
}