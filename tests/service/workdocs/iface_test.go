package workdocs_test

// tests for the workdocs service interface for this ../../../service/workdocs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workdocs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workdocs/workdocs_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkdocsServiceCanBeMocked(t *testing.T) {
	var iface workdocs_iface.IClient
	iface = &workdocs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workdocs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortDocumentVersionUpload", func(t *testing.T) {
        input := &workdocs.AbortDocumentVersionUploadInput{}
        output := &workdocs.AbortDocumentVersionUploadOutput{}

        mockClient.On("AbortDocumentVersionUpload", ctx, input).Return(output, nil)

        result, err := mockClient.AbortDocumentVersionUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateUser", func(t *testing.T) {
        input := &workdocs.ActivateUserInput{}
        output := &workdocs.ActivateUserOutput{}

        mockClient.On("ActivateUser", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddResourcePermissions", func(t *testing.T) {
        input := &workdocs.AddResourcePermissionsInput{}
        output := &workdocs.AddResourcePermissionsOutput{}

        mockClient.On("AddResourcePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.AddResourcePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComment", func(t *testing.T) {
        input := &workdocs.CreateCommentInput{}
        output := &workdocs.CreateCommentOutput{}

        mockClient.On("CreateComment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomMetadata", func(t *testing.T) {
        input := &workdocs.CreateCustomMetadataInput{}
        output := &workdocs.CreateCustomMetadataOutput{}

        mockClient.On("CreateCustomMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFolder", func(t *testing.T) {
        input := &workdocs.CreateFolderInput{}
        output := &workdocs.CreateFolderOutput{}

        mockClient.On("CreateFolder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLabels", func(t *testing.T) {
        input := &workdocs.CreateLabelsInput{}
        output := &workdocs.CreateLabelsOutput{}

        mockClient.On("CreateLabels", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotificationSubscription", func(t *testing.T) {
        input := &workdocs.CreateNotificationSubscriptionInput{}
        output := &workdocs.CreateNotificationSubscriptionOutput{}

        mockClient.On("CreateNotificationSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotificationSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &workdocs.CreateUserInput{}
        output := &workdocs.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateUser", func(t *testing.T) {
        input := &workdocs.DeactivateUserInput{}
        output := &workdocs.DeactivateUserOutput{}

        mockClient.On("DeactivateUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComment", func(t *testing.T) {
        input := &workdocs.DeleteCommentInput{}
        output := &workdocs.DeleteCommentOutput{}

        mockClient.On("DeleteComment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomMetadata", func(t *testing.T) {
        input := &workdocs.DeleteCustomMetadataInput{}
        output := &workdocs.DeleteCustomMetadataOutput{}

        mockClient.On("DeleteCustomMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocument", func(t *testing.T) {
        input := &workdocs.DeleteDocumentInput{}
        output := &workdocs.DeleteDocumentOutput{}

        mockClient.On("DeleteDocument", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocumentVersion", func(t *testing.T) {
        input := &workdocs.DeleteDocumentVersionInput{}
        output := &workdocs.DeleteDocumentVersionOutput{}

        mockClient.On("DeleteDocumentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocumentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFolder", func(t *testing.T) {
        input := &workdocs.DeleteFolderInput{}
        output := &workdocs.DeleteFolderOutput{}

        mockClient.On("DeleteFolder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFolderContents", func(t *testing.T) {
        input := &workdocs.DeleteFolderContentsInput{}
        output := &workdocs.DeleteFolderContentsOutput{}

        mockClient.On("DeleteFolderContents", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFolderContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLabels", func(t *testing.T) {
        input := &workdocs.DeleteLabelsInput{}
        output := &workdocs.DeleteLabelsOutput{}

        mockClient.On("DeleteLabels", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotificationSubscription", func(t *testing.T) {
        input := &workdocs.DeleteNotificationSubscriptionInput{}
        output := &workdocs.DeleteNotificationSubscriptionOutput{}

        mockClient.On("DeleteNotificationSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotificationSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &workdocs.DeleteUserInput{}
        output := &workdocs.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActivities", func(t *testing.T) {
        input := &workdocs.DescribeActivitiesInput{}
        output := &workdocs.DescribeActivitiesOutput{}

        mockClient.On("DescribeActivities", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComments", func(t *testing.T) {
        input := &workdocs.DescribeCommentsInput{}
        output := &workdocs.DescribeCommentsOutput{}

        mockClient.On("DescribeComments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDocumentVersions", func(t *testing.T) {
        input := &workdocs.DescribeDocumentVersionsInput{}
        output := &workdocs.DescribeDocumentVersionsOutput{}

        mockClient.On("DescribeDocumentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDocumentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFolderContents", func(t *testing.T) {
        input := &workdocs.DescribeFolderContentsInput{}
        output := &workdocs.DescribeFolderContentsOutput{}

        mockClient.On("DescribeFolderContents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFolderContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroups", func(t *testing.T) {
        input := &workdocs.DescribeGroupsInput{}
        output := &workdocs.DescribeGroupsOutput{}

        mockClient.On("DescribeGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotificationSubscriptions", func(t *testing.T) {
        input := &workdocs.DescribeNotificationSubscriptionsInput{}
        output := &workdocs.DescribeNotificationSubscriptionsOutput{}

        mockClient.On("DescribeNotificationSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotificationSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePermissions", func(t *testing.T) {
        input := &workdocs.DescribeResourcePermissionsInput{}
        output := &workdocs.DescribeResourcePermissionsOutput{}

        mockClient.On("DescribeResourcePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRootFolders", func(t *testing.T) {
        input := &workdocs.DescribeRootFoldersInput{}
        output := &workdocs.DescribeRootFoldersOutput{}

        mockClient.On("DescribeRootFolders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRootFolders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsers", func(t *testing.T) {
        input := &workdocs.DescribeUsersInput{}
        output := &workdocs.DescribeUsersOutput{}

        mockClient.On("DescribeUsers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCurrentUser", func(t *testing.T) {
        input := &workdocs.GetCurrentUserInput{}
        output := &workdocs.GetCurrentUserOutput{}

        mockClient.On("GetCurrentUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetCurrentUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocument", func(t *testing.T) {
        input := &workdocs.GetDocumentInput{}
        output := &workdocs.GetDocumentOutput{}

        mockClient.On("GetDocument", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentPath", func(t *testing.T) {
        input := &workdocs.GetDocumentPathInput{}
        output := &workdocs.GetDocumentPathOutput{}

        mockClient.On("GetDocumentPath", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentVersion", func(t *testing.T) {
        input := &workdocs.GetDocumentVersionInput{}
        output := &workdocs.GetDocumentVersionOutput{}

        mockClient.On("GetDocumentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFolder", func(t *testing.T) {
        input := &workdocs.GetFolderInput{}
        output := &workdocs.GetFolderOutput{}

        mockClient.On("GetFolder", ctx, input).Return(output, nil)

        result, err := mockClient.GetFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFolderPath", func(t *testing.T) {
        input := &workdocs.GetFolderPathInput{}
        output := &workdocs.GetFolderPathOutput{}

        mockClient.On("GetFolderPath", ctx, input).Return(output, nil)

        result, err := mockClient.GetFolderPath(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResources", func(t *testing.T) {
        input := &workdocs.GetResourcesInput{}
        output := &workdocs.GetResourcesOutput{}

        mockClient.On("GetResources", ctx, input).Return(output, nil)

        result, err := mockClient.GetResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitiateDocumentVersionUpload", func(t *testing.T) {
        input := &workdocs.InitiateDocumentVersionUploadInput{}
        output := &workdocs.InitiateDocumentVersionUploadOutput{}

        mockClient.On("InitiateDocumentVersionUpload", ctx, input).Return(output, nil)

        result, err := mockClient.InitiateDocumentVersionUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAllResourcePermissions", func(t *testing.T) {
        input := &workdocs.RemoveAllResourcePermissionsInput{}
        output := &workdocs.RemoveAllResourcePermissionsOutput{}

        mockClient.On("RemoveAllResourcePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAllResourcePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveResourcePermission", func(t *testing.T) {
        input := &workdocs.RemoveResourcePermissionInput{}
        output := &workdocs.RemoveResourcePermissionOutput{}

        mockClient.On("RemoveResourcePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveResourcePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDocumentVersions", func(t *testing.T) {
        input := &workdocs.RestoreDocumentVersionsInput{}
        output := &workdocs.RestoreDocumentVersionsOutput{}

        mockClient.On("RestoreDocumentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDocumentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchResources", func(t *testing.T) {
        input := &workdocs.SearchResourcesInput{}
        output := &workdocs.SearchResourcesOutput{}

        mockClient.On("SearchResources", ctx, input).Return(output, nil)

        result, err := mockClient.SearchResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocument", func(t *testing.T) {
        input := &workdocs.UpdateDocumentInput{}
        output := &workdocs.UpdateDocumentOutput{}

        mockClient.On("UpdateDocument", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocumentVersion", func(t *testing.T) {
        input := &workdocs.UpdateDocumentVersionInput{}
        output := &workdocs.UpdateDocumentVersionOutput{}

        mockClient.On("UpdateDocumentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocumentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFolder", func(t *testing.T) {
        input := &workdocs.UpdateFolderInput{}
        output := &workdocs.UpdateFolderOutput{}

        mockClient.On("UpdateFolder", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &workdocs.UpdateUserInput{}
        output := &workdocs.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
