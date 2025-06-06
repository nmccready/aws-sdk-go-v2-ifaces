// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package qapps_test

// tests for the qapps service interface for 
// this ../../../service/qapps/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qapps"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qapps/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qapps/qapps_iface"
	"github.com/stretchr/testify/assert"
)

func TestQappsServiceCanBeMocked(t *testing.T) {
	var iface qapps_iface.IClient
	iface = &qapps.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := qapps.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLibraryItemReview", func(t *testing.T) {
        input := &qapps.AssociateLibraryItemReviewInput{}
        output := &qapps.AssociateLibraryItemReviewOutput{}

        mockClient.On("AssociateLibraryItemReview", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLibraryItemReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateQAppWithUser", func(t *testing.T) {
        input := &qapps.AssociateQAppWithUserInput{}
        output := &qapps.AssociateQAppWithUserOutput{}

        mockClient.On("AssociateQAppWithUser", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateQAppWithUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateCategory", func(t *testing.T) {
        input := &qapps.BatchCreateCategoryInput{}
        output := &qapps.BatchCreateCategoryOutput{}

        mockClient.On("BatchCreateCategory", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteCategory", func(t *testing.T) {
        input := &qapps.BatchDeleteCategoryInput{}
        output := &qapps.BatchDeleteCategoryOutput{}

        mockClient.On("BatchDeleteCategory", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateCategory", func(t *testing.T) {
        input := &qapps.BatchUpdateCategoryInput{}
        output := &qapps.BatchUpdateCategoryOutput{}

        mockClient.On("BatchUpdateCategory", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateCategory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLibraryItem", func(t *testing.T) {
        input := &qapps.CreateLibraryItemInput{}
        output := &qapps.CreateLibraryItemOutput{}

        mockClient.On("CreateLibraryItem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLibraryItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePresignedUrl", func(t *testing.T) {
        input := &qapps.CreatePresignedUrlInput{}
        output := &qapps.CreatePresignedUrlOutput{}

        mockClient.On("CreatePresignedUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePresignedUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQApp", func(t *testing.T) {
        input := &qapps.CreateQAppInput{}
        output := &qapps.CreateQAppOutput{}

        mockClient.On("CreateQApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLibraryItem", func(t *testing.T) {
        input := &qapps.DeleteLibraryItemInput{}
        output := &qapps.DeleteLibraryItemOutput{}

        mockClient.On("DeleteLibraryItem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLibraryItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQApp", func(t *testing.T) {
        input := &qapps.DeleteQAppInput{}
        output := &qapps.DeleteQAppOutput{}

        mockClient.On("DeleteQApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQAppPermissions", func(t *testing.T) {
        input := &qapps.DescribeQAppPermissionsInput{}
        output := &qapps.DescribeQAppPermissionsOutput{}

        mockClient.On("DescribeQAppPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQAppPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLibraryItemReview", func(t *testing.T) {
        input := &qapps.DisassociateLibraryItemReviewInput{}
        output := &qapps.DisassociateLibraryItemReviewOutput{}

        mockClient.On("DisassociateLibraryItemReview", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLibraryItemReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateQAppFromUser", func(t *testing.T) {
        input := &qapps.DisassociateQAppFromUserInput{}
        output := &qapps.DisassociateQAppFromUserOutput{}

        mockClient.On("DisassociateQAppFromUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateQAppFromUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportQAppSessionData", func(t *testing.T) {
        input := &qapps.ExportQAppSessionDataInput{}
        output := &qapps.ExportQAppSessionDataOutput{}

        mockClient.On("ExportQAppSessionData", ctx, input).Return(output, nil)

        result, err := mockClient.ExportQAppSessionData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLibraryItem", func(t *testing.T) {
        input := &qapps.GetLibraryItemInput{}
        output := &qapps.GetLibraryItemOutput{}

        mockClient.On("GetLibraryItem", ctx, input).Return(output, nil)

        result, err := mockClient.GetLibraryItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQApp", func(t *testing.T) {
        input := &qapps.GetQAppInput{}
        output := &qapps.GetQAppOutput{}

        mockClient.On("GetQApp", ctx, input).Return(output, nil)

        result, err := mockClient.GetQApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQAppSession", func(t *testing.T) {
        input := &qapps.GetQAppSessionInput{}
        output := &qapps.GetQAppSessionOutput{}

        mockClient.On("GetQAppSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetQAppSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQAppSessionMetadata", func(t *testing.T) {
        input := &qapps.GetQAppSessionMetadataInput{}
        output := &qapps.GetQAppSessionMetadataOutput{}

        mockClient.On("GetQAppSessionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetQAppSessionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportDocument", func(t *testing.T) {
        input := &qapps.ImportDocumentInput{}
        output := &qapps.ImportDocumentOutput{}

        mockClient.On("ImportDocument", ctx, input).Return(output, nil)

        result, err := mockClient.ImportDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCategories", func(t *testing.T) {
        input := &qapps.ListCategoriesInput{}
        output := &qapps.ListCategoriesOutput{}

        mockClient.On("ListCategories", ctx, input).Return(output, nil)

        result, err := mockClient.ListCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLibraryItems", func(t *testing.T) {
        input := &qapps.ListLibraryItemsInput{}
        output := &qapps.ListLibraryItemsOutput{}

        mockClient.On("ListLibraryItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListLibraryItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQAppSessionData", func(t *testing.T) {
        input := &qapps.ListQAppSessionDataInput{}
        output := &qapps.ListQAppSessionDataOutput{}

        mockClient.On("ListQAppSessionData", ctx, input).Return(output, nil)

        result, err := mockClient.ListQAppSessionData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQApps", func(t *testing.T) {
        input := &qapps.ListQAppsInput{}
        output := &qapps.ListQAppsOutput{}

        mockClient.On("ListQApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListQApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &qapps.ListTagsForResourceInput{}
        output := &qapps.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPredictQApp", func(t *testing.T) {
        input := &qapps.PredictQAppInput{}
        output := &qapps.PredictQAppOutput{}

        mockClient.On("PredictQApp", ctx, input).Return(output, nil)

        result, err := mockClient.PredictQApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQAppSession", func(t *testing.T) {
        input := &qapps.StartQAppSessionInput{}
        output := &qapps.StartQAppSessionOutput{}

        mockClient.On("StartQAppSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartQAppSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQAppSession", func(t *testing.T) {
        input := &qapps.StopQAppSessionInput{}
        output := &qapps.StopQAppSessionOutput{}

        mockClient.On("StopQAppSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopQAppSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &qapps.TagResourceInput{}
        output := &qapps.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &qapps.UntagResourceInput{}
        output := &qapps.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLibraryItem", func(t *testing.T) {
        input := &qapps.UpdateLibraryItemInput{}
        output := &qapps.UpdateLibraryItemOutput{}

        mockClient.On("UpdateLibraryItem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLibraryItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLibraryItemMetadata", func(t *testing.T) {
        input := &qapps.UpdateLibraryItemMetadataInput{}
        output := &qapps.UpdateLibraryItemMetadataOutput{}

        mockClient.On("UpdateLibraryItemMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLibraryItemMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQApp", func(t *testing.T) {
        input := &qapps.UpdateQAppInput{}
        output := &qapps.UpdateQAppOutput{}

        mockClient.On("UpdateQApp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQAppPermissions", func(t *testing.T) {
        input := &qapps.UpdateQAppPermissionsInput{}
        output := &qapps.UpdateQAppPermissionsOutput{}

        mockClient.On("UpdateQAppPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQAppPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQAppSession", func(t *testing.T) {
        input := &qapps.UpdateQAppSessionInput{}
        output := &qapps.UpdateQAppSessionOutput{}

        mockClient.On("UpdateQAppSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQAppSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQAppSessionMetadata", func(t *testing.T) {
        input := &qapps.UpdateQAppSessionMetadataInput{}
        output := &qapps.UpdateQAppSessionMetadataOutput{}

        mockClient.On("UpdateQAppSessionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQAppSessionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
