// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package clouddirectory_test

// tests for the clouddirectory service interface for 
// this ../../../service/clouddirectory/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/clouddirectory/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/clouddirectory/clouddirectory_iface"
	"github.com/stretchr/testify/assert"
)

func TestClouddirectoryServiceCanBeMocked(t *testing.T) {
	var iface clouddirectory_iface.IClient
	iface = &clouddirectory.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := clouddirectory.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddFacetToObject", func(t *testing.T) {
        input := &clouddirectory.AddFacetToObjectInput{}
        output := &clouddirectory.AddFacetToObjectOutput{}

        mockClient.On("AddFacetToObject", ctx, input).Return(output, nil)

        result, err := mockClient.AddFacetToObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplySchema", func(t *testing.T) {
        input := &clouddirectory.ApplySchemaInput{}
        output := &clouddirectory.ApplySchemaOutput{}

        mockClient.On("ApplySchema", ctx, input).Return(output, nil)

        result, err := mockClient.ApplySchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachObject", func(t *testing.T) {
        input := &clouddirectory.AttachObjectInput{}
        output := &clouddirectory.AttachObjectOutput{}

        mockClient.On("AttachObject", ctx, input).Return(output, nil)

        result, err := mockClient.AttachObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachPolicy", func(t *testing.T) {
        input := &clouddirectory.AttachPolicyInput{}
        output := &clouddirectory.AttachPolicyOutput{}

        mockClient.On("AttachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachToIndex", func(t *testing.T) {
        input := &clouddirectory.AttachToIndexInput{}
        output := &clouddirectory.AttachToIndexOutput{}

        mockClient.On("AttachToIndex", ctx, input).Return(output, nil)

        result, err := mockClient.AttachToIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachTypedLink", func(t *testing.T) {
        input := &clouddirectory.AttachTypedLinkInput{}
        output := &clouddirectory.AttachTypedLinkOutput{}

        mockClient.On("AttachTypedLink", ctx, input).Return(output, nil)

        result, err := mockClient.AttachTypedLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchRead", func(t *testing.T) {
        input := &clouddirectory.BatchReadInput{}
        output := &clouddirectory.BatchReadOutput{}

        mockClient.On("BatchRead", ctx, input).Return(output, nil)

        result, err := mockClient.BatchRead(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchWrite", func(t *testing.T) {
        input := &clouddirectory.BatchWriteInput{}
        output := &clouddirectory.BatchWriteOutput{}

        mockClient.On("BatchWrite", ctx, input).Return(output, nil)

        result, err := mockClient.BatchWrite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectory", func(t *testing.T) {
        input := &clouddirectory.CreateDirectoryInput{}
        output := &clouddirectory.CreateDirectoryOutput{}

        mockClient.On("CreateDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFacet", func(t *testing.T) {
        input := &clouddirectory.CreateFacetInput{}
        output := &clouddirectory.CreateFacetOutput{}

        mockClient.On("CreateFacet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIndex", func(t *testing.T) {
        input := &clouddirectory.CreateIndexInput{}
        output := &clouddirectory.CreateIndexOutput{}

        mockClient.On("CreateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateObject", func(t *testing.T) {
        input := &clouddirectory.CreateObjectInput{}
        output := &clouddirectory.CreateObjectOutput{}

        mockClient.On("CreateObject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchema", func(t *testing.T) {
        input := &clouddirectory.CreateSchemaInput{}
        output := &clouddirectory.CreateSchemaOutput{}

        mockClient.On("CreateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTypedLinkFacet", func(t *testing.T) {
        input := &clouddirectory.CreateTypedLinkFacetInput{}
        output := &clouddirectory.CreateTypedLinkFacetOutput{}

        mockClient.On("CreateTypedLinkFacet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTypedLinkFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectory", func(t *testing.T) {
        input := &clouddirectory.DeleteDirectoryInput{}
        output := &clouddirectory.DeleteDirectoryOutput{}

        mockClient.On("DeleteDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFacet", func(t *testing.T) {
        input := &clouddirectory.DeleteFacetInput{}
        output := &clouddirectory.DeleteFacetOutput{}

        mockClient.On("DeleteFacet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObject", func(t *testing.T) {
        input := &clouddirectory.DeleteObjectInput{}
        output := &clouddirectory.DeleteObjectOutput{}

        mockClient.On("DeleteObject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchema", func(t *testing.T) {
        input := &clouddirectory.DeleteSchemaInput{}
        output := &clouddirectory.DeleteSchemaOutput{}

        mockClient.On("DeleteSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTypedLinkFacet", func(t *testing.T) {
        input := &clouddirectory.DeleteTypedLinkFacetInput{}
        output := &clouddirectory.DeleteTypedLinkFacetOutput{}

        mockClient.On("DeleteTypedLinkFacet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTypedLinkFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachFromIndex", func(t *testing.T) {
        input := &clouddirectory.DetachFromIndexInput{}
        output := &clouddirectory.DetachFromIndexOutput{}

        mockClient.On("DetachFromIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DetachFromIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachObject", func(t *testing.T) {
        input := &clouddirectory.DetachObjectInput{}
        output := &clouddirectory.DetachObjectOutput{}

        mockClient.On("DetachObject", ctx, input).Return(output, nil)

        result, err := mockClient.DetachObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachPolicy", func(t *testing.T) {
        input := &clouddirectory.DetachPolicyInput{}
        output := &clouddirectory.DetachPolicyOutput{}

        mockClient.On("DetachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachTypedLink", func(t *testing.T) {
        input := &clouddirectory.DetachTypedLinkInput{}
        output := &clouddirectory.DetachTypedLinkOutput{}

        mockClient.On("DetachTypedLink", ctx, input).Return(output, nil)

        result, err := mockClient.DetachTypedLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableDirectory", func(t *testing.T) {
        input := &clouddirectory.DisableDirectoryInput{}
        output := &clouddirectory.DisableDirectoryOutput{}

        mockClient.On("DisableDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.DisableDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableDirectory", func(t *testing.T) {
        input := &clouddirectory.EnableDirectoryInput{}
        output := &clouddirectory.EnableDirectoryOutput{}

        mockClient.On("EnableDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.EnableDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppliedSchemaVersion", func(t *testing.T) {
        input := &clouddirectory.GetAppliedSchemaVersionInput{}
        output := &clouddirectory.GetAppliedSchemaVersionOutput{}

        mockClient.On("GetAppliedSchemaVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppliedSchemaVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDirectory", func(t *testing.T) {
        input := &clouddirectory.GetDirectoryInput{}
        output := &clouddirectory.GetDirectoryOutput{}

        mockClient.On("GetDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.GetDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFacet", func(t *testing.T) {
        input := &clouddirectory.GetFacetInput{}
        output := &clouddirectory.GetFacetOutput{}

        mockClient.On("GetFacet", ctx, input).Return(output, nil)

        result, err := mockClient.GetFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLinkAttributes", func(t *testing.T) {
        input := &clouddirectory.GetLinkAttributesInput{}
        output := &clouddirectory.GetLinkAttributesOutput{}

        mockClient.On("GetLinkAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetLinkAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectAttributes", func(t *testing.T) {
        input := &clouddirectory.GetObjectAttributesInput{}
        output := &clouddirectory.GetObjectAttributesOutput{}

        mockClient.On("GetObjectAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObjectInformation", func(t *testing.T) {
        input := &clouddirectory.GetObjectInformationInput{}
        output := &clouddirectory.GetObjectInformationOutput{}

        mockClient.On("GetObjectInformation", ctx, input).Return(output, nil)

        result, err := mockClient.GetObjectInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaAsJson", func(t *testing.T) {
        input := &clouddirectory.GetSchemaAsJsonInput{}
        output := &clouddirectory.GetSchemaAsJsonOutput{}

        mockClient.On("GetSchemaAsJson", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaAsJson(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTypedLinkFacetInformation", func(t *testing.T) {
        input := &clouddirectory.GetTypedLinkFacetInformationInput{}
        output := &clouddirectory.GetTypedLinkFacetInformationOutput{}

        mockClient.On("GetTypedLinkFacetInformation", ctx, input).Return(output, nil)

        result, err := mockClient.GetTypedLinkFacetInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppliedSchemaArns", func(t *testing.T) {
        input := &clouddirectory.ListAppliedSchemaArnsInput{}
        output := &clouddirectory.ListAppliedSchemaArnsOutput{}

        mockClient.On("ListAppliedSchemaArns", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppliedSchemaArns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedIndices", func(t *testing.T) {
        input := &clouddirectory.ListAttachedIndicesInput{}
        output := &clouddirectory.ListAttachedIndicesOutput{}

        mockClient.On("ListAttachedIndices", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedIndices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevelopmentSchemaArns", func(t *testing.T) {
        input := &clouddirectory.ListDevelopmentSchemaArnsInput{}
        output := &clouddirectory.ListDevelopmentSchemaArnsOutput{}

        mockClient.On("ListDevelopmentSchemaArns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevelopmentSchemaArns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDirectories", func(t *testing.T) {
        input := &clouddirectory.ListDirectoriesInput{}
        output := &clouddirectory.ListDirectoriesOutput{}

        mockClient.On("ListDirectories", ctx, input).Return(output, nil)

        result, err := mockClient.ListDirectories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFacetAttributes", func(t *testing.T) {
        input := &clouddirectory.ListFacetAttributesInput{}
        output := &clouddirectory.ListFacetAttributesOutput{}

        mockClient.On("ListFacetAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListFacetAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFacetNames", func(t *testing.T) {
        input := &clouddirectory.ListFacetNamesInput{}
        output := &clouddirectory.ListFacetNamesOutput{}

        mockClient.On("ListFacetNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListFacetNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIncomingTypedLinks", func(t *testing.T) {
        input := &clouddirectory.ListIncomingTypedLinksInput{}
        output := &clouddirectory.ListIncomingTypedLinksOutput{}

        mockClient.On("ListIncomingTypedLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListIncomingTypedLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndex", func(t *testing.T) {
        input := &clouddirectory.ListIndexInput{}
        output := &clouddirectory.ListIndexOutput{}

        mockClient.On("ListIndex", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedSchemaArns", func(t *testing.T) {
        input := &clouddirectory.ListManagedSchemaArnsInput{}
        output := &clouddirectory.ListManagedSchemaArnsOutput{}

        mockClient.On("ListManagedSchemaArns", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedSchemaArns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectAttributes", func(t *testing.T) {
        input := &clouddirectory.ListObjectAttributesInput{}
        output := &clouddirectory.ListObjectAttributesOutput{}

        mockClient.On("ListObjectAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectChildren", func(t *testing.T) {
        input := &clouddirectory.ListObjectChildrenInput{}
        output := &clouddirectory.ListObjectChildrenOutput{}

        mockClient.On("ListObjectChildren", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectChildren(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectParentPaths", func(t *testing.T) {
        input := &clouddirectory.ListObjectParentPathsInput{}
        output := &clouddirectory.ListObjectParentPathsOutput{}

        mockClient.On("ListObjectParentPaths", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectParentPaths(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectParents", func(t *testing.T) {
        input := &clouddirectory.ListObjectParentsInput{}
        output := &clouddirectory.ListObjectParentsOutput{}

        mockClient.On("ListObjectParents", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectParents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectPolicies", func(t *testing.T) {
        input := &clouddirectory.ListObjectPoliciesInput{}
        output := &clouddirectory.ListObjectPoliciesOutput{}

        mockClient.On("ListObjectPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOutgoingTypedLinks", func(t *testing.T) {
        input := &clouddirectory.ListOutgoingTypedLinksInput{}
        output := &clouddirectory.ListOutgoingTypedLinksOutput{}

        mockClient.On("ListOutgoingTypedLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListOutgoingTypedLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyAttachments", func(t *testing.T) {
        input := &clouddirectory.ListPolicyAttachmentsInput{}
        output := &clouddirectory.ListPolicyAttachmentsOutput{}

        mockClient.On("ListPolicyAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPublishedSchemaArns", func(t *testing.T) {
        input := &clouddirectory.ListPublishedSchemaArnsInput{}
        output := &clouddirectory.ListPublishedSchemaArnsOutput{}

        mockClient.On("ListPublishedSchemaArns", ctx, input).Return(output, nil)

        result, err := mockClient.ListPublishedSchemaArns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &clouddirectory.ListTagsForResourceInput{}
        output := &clouddirectory.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypedLinkFacetAttributes", func(t *testing.T) {
        input := &clouddirectory.ListTypedLinkFacetAttributesInput{}
        output := &clouddirectory.ListTypedLinkFacetAttributesOutput{}

        mockClient.On("ListTypedLinkFacetAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypedLinkFacetAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypedLinkFacetNames", func(t *testing.T) {
        input := &clouddirectory.ListTypedLinkFacetNamesInput{}
        output := &clouddirectory.ListTypedLinkFacetNamesOutput{}

        mockClient.On("ListTypedLinkFacetNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypedLinkFacetNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLookupPolicy", func(t *testing.T) {
        input := &clouddirectory.LookupPolicyInput{}
        output := &clouddirectory.LookupPolicyOutput{}

        mockClient.On("LookupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.LookupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishSchema", func(t *testing.T) {
        input := &clouddirectory.PublishSchemaInput{}
        output := &clouddirectory.PublishSchemaOutput{}

        mockClient.On("PublishSchema", ctx, input).Return(output, nil)

        result, err := mockClient.PublishSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSchemaFromJson", func(t *testing.T) {
        input := &clouddirectory.PutSchemaFromJsonInput{}
        output := &clouddirectory.PutSchemaFromJsonOutput{}

        mockClient.On("PutSchemaFromJson", ctx, input).Return(output, nil)

        result, err := mockClient.PutSchemaFromJson(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveFacetFromObject", func(t *testing.T) {
        input := &clouddirectory.RemoveFacetFromObjectInput{}
        output := &clouddirectory.RemoveFacetFromObjectOutput{}

        mockClient.On("RemoveFacetFromObject", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveFacetFromObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &clouddirectory.TagResourceInput{}
        output := &clouddirectory.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &clouddirectory.UntagResourceInput{}
        output := &clouddirectory.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFacet", func(t *testing.T) {
        input := &clouddirectory.UpdateFacetInput{}
        output := &clouddirectory.UpdateFacetOutput{}

        mockClient.On("UpdateFacet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLinkAttributes", func(t *testing.T) {
        input := &clouddirectory.UpdateLinkAttributesInput{}
        output := &clouddirectory.UpdateLinkAttributesOutput{}

        mockClient.On("UpdateLinkAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLinkAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateObjectAttributes", func(t *testing.T) {
        input := &clouddirectory.UpdateObjectAttributesInput{}
        output := &clouddirectory.UpdateObjectAttributesOutput{}

        mockClient.On("UpdateObjectAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateObjectAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchema", func(t *testing.T) {
        input := &clouddirectory.UpdateSchemaInput{}
        output := &clouddirectory.UpdateSchemaOutput{}

        mockClient.On("UpdateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTypedLinkFacet", func(t *testing.T) {
        input := &clouddirectory.UpdateTypedLinkFacetInput{}
        output := &clouddirectory.UpdateTypedLinkFacetOutput{}

        mockClient.On("UpdateTypedLinkFacet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTypedLinkFacet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeAppliedSchema", func(t *testing.T) {
        input := &clouddirectory.UpgradeAppliedSchemaInput{}
        output := &clouddirectory.UpgradeAppliedSchemaOutput{}

        mockClient.On("UpgradeAppliedSchema", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeAppliedSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradePublishedSchema", func(t *testing.T) {
        input := &clouddirectory.UpgradePublishedSchemaInput{}
        output := &clouddirectory.UpgradePublishedSchemaOutput{}

        mockClient.On("UpgradePublishedSchema", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradePublishedSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
