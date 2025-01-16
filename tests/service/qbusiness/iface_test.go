// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package qbusiness_test

// tests for the qbusiness service interface for 
// this ../../../service/qbusiness/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qbusiness"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qbusiness/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qbusiness/qbusiness_iface"
	"github.com/stretchr/testify/assert"
)

func TestQbusinessServiceCanBeMocked(t *testing.T) {
	var iface qbusiness_iface.IClient
	iface = &qbusiness.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := qbusiness.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePermission", func(t *testing.T) {
        input := &qbusiness.AssociatePermissionInput{}
        output := &qbusiness.AssociatePermissionOutput{}

        mockClient.On("AssociatePermission", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteDocument", func(t *testing.T) {
        input := &qbusiness.BatchDeleteDocumentInput{}
        output := &qbusiness.BatchDeleteDocumentOutput{}

        mockClient.On("BatchDeleteDocument", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutDocument", func(t *testing.T) {
        input := &qbusiness.BatchPutDocumentInput{}
        output := &qbusiness.BatchPutDocumentOutput{}

        mockClient.On("BatchPutDocument", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChat", func(t *testing.T) {
        input := &qbusiness.ChatInput{}
        output := &qbusiness.ChatOutput{}

        mockClient.On("Chat", ctx, input).Return(output, nil)

        result, err := mockClient.Chat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChatSync", func(t *testing.T) {
        input := &qbusiness.ChatSyncInput{}
        output := &qbusiness.ChatSyncOutput{}

        mockClient.On("ChatSync", ctx, input).Return(output, nil)

        result, err := mockClient.ChatSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &qbusiness.CreateApplicationInput{}
        output := &qbusiness.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataAccessor", func(t *testing.T) {
        input := &qbusiness.CreateDataAccessorInput{}
        output := &qbusiness.CreateDataAccessorOutput{}

        mockClient.On("CreateDataAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &qbusiness.CreateDataSourceInput{}
        output := &qbusiness.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIndex", func(t *testing.T) {
        input := &qbusiness.CreateIndexInput{}
        output := &qbusiness.CreateIndexOutput{}

        mockClient.On("CreateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlugin", func(t *testing.T) {
        input := &qbusiness.CreatePluginInput{}
        output := &qbusiness.CreatePluginOutput{}

        mockClient.On("CreatePlugin", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRetriever", func(t *testing.T) {
        input := &qbusiness.CreateRetrieverInput{}
        output := &qbusiness.CreateRetrieverOutput{}

        mockClient.On("CreateRetriever", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRetriever(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &qbusiness.CreateUserInput{}
        output := &qbusiness.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebExperience", func(t *testing.T) {
        input := &qbusiness.CreateWebExperienceInput{}
        output := &qbusiness.CreateWebExperienceOutput{}

        mockClient.On("CreateWebExperience", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &qbusiness.DeleteApplicationInput{}
        output := &qbusiness.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChatControlsConfiguration", func(t *testing.T) {
        input := &qbusiness.DeleteChatControlsConfigurationInput{}
        output := &qbusiness.DeleteChatControlsConfigurationOutput{}

        mockClient.On("DeleteChatControlsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChatControlsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConversation", func(t *testing.T) {
        input := &qbusiness.DeleteConversationInput{}
        output := &qbusiness.DeleteConversationOutput{}

        mockClient.On("DeleteConversation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConversation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataAccessor", func(t *testing.T) {
        input := &qbusiness.DeleteDataAccessorInput{}
        output := &qbusiness.DeleteDataAccessorOutput{}

        mockClient.On("DeleteDataAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &qbusiness.DeleteDataSourceInput{}
        output := &qbusiness.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &qbusiness.DeleteGroupInput{}
        output := &qbusiness.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIndex", func(t *testing.T) {
        input := &qbusiness.DeleteIndexInput{}
        output := &qbusiness.DeleteIndexOutput{}

        mockClient.On("DeleteIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlugin", func(t *testing.T) {
        input := &qbusiness.DeletePluginInput{}
        output := &qbusiness.DeletePluginOutput{}

        mockClient.On("DeletePlugin", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRetriever", func(t *testing.T) {
        input := &qbusiness.DeleteRetrieverInput{}
        output := &qbusiness.DeleteRetrieverOutput{}

        mockClient.On("DeleteRetriever", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRetriever(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &qbusiness.DeleteUserInput{}
        output := &qbusiness.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebExperience", func(t *testing.T) {
        input := &qbusiness.DeleteWebExperienceInput{}
        output := &qbusiness.DeleteWebExperienceOutput{}

        mockClient.On("DeleteWebExperience", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePermission", func(t *testing.T) {
        input := &qbusiness.DisassociatePermissionInput{}
        output := &qbusiness.DisassociatePermissionOutput{}

        mockClient.On("DisassociatePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &qbusiness.GetApplicationInput{}
        output := &qbusiness.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChatControlsConfiguration", func(t *testing.T) {
        input := &qbusiness.GetChatControlsConfigurationInput{}
        output := &qbusiness.GetChatControlsConfigurationOutput{}

        mockClient.On("GetChatControlsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetChatControlsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataAccessor", func(t *testing.T) {
        input := &qbusiness.GetDataAccessorInput{}
        output := &qbusiness.GetDataAccessorOutput{}

        mockClient.On("GetDataAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &qbusiness.GetDataSourceInput{}
        output := &qbusiness.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &qbusiness.GetGroupInput{}
        output := &qbusiness.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIndex", func(t *testing.T) {
        input := &qbusiness.GetIndexInput{}
        output := &qbusiness.GetIndexOutput{}

        mockClient.On("GetIndex", ctx, input).Return(output, nil)

        result, err := mockClient.GetIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMedia", func(t *testing.T) {
        input := &qbusiness.GetMediaInput{}
        output := &qbusiness.GetMediaOutput{}

        mockClient.On("GetMedia", ctx, input).Return(output, nil)

        result, err := mockClient.GetMedia(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlugin", func(t *testing.T) {
        input := &qbusiness.GetPluginInput{}
        output := &qbusiness.GetPluginOutput{}

        mockClient.On("GetPlugin", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &qbusiness.GetPolicyInput{}
        output := &qbusiness.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRetriever", func(t *testing.T) {
        input := &qbusiness.GetRetrieverInput{}
        output := &qbusiness.GetRetrieverOutput{}

        mockClient.On("GetRetriever", ctx, input).Return(output, nil)

        result, err := mockClient.GetRetriever(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUser", func(t *testing.T) {
        input := &qbusiness.GetUserInput{}
        output := &qbusiness.GetUserOutput{}

        mockClient.On("GetUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebExperience", func(t *testing.T) {
        input := &qbusiness.GetWebExperienceInput{}
        output := &qbusiness.GetWebExperienceOutput{}

        mockClient.On("GetWebExperience", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &qbusiness.ListApplicationsInput{}
        output := &qbusiness.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachments", func(t *testing.T) {
        input := &qbusiness.ListAttachmentsInput{}
        output := &qbusiness.ListAttachmentsOutput{}

        mockClient.On("ListAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConversations", func(t *testing.T) {
        input := &qbusiness.ListConversationsInput{}
        output := &qbusiness.ListConversationsOutput{}

        mockClient.On("ListConversations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConversations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataAccessors", func(t *testing.T) {
        input := &qbusiness.ListDataAccessorsInput{}
        output := &qbusiness.ListDataAccessorsOutput{}

        mockClient.On("ListDataAccessors", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataAccessors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSourceSyncJobs", func(t *testing.T) {
        input := &qbusiness.ListDataSourceSyncJobsInput{}
        output := &qbusiness.ListDataSourceSyncJobsOutput{}

        mockClient.On("ListDataSourceSyncJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSourceSyncJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &qbusiness.ListDataSourcesInput{}
        output := &qbusiness.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocuments", func(t *testing.T) {
        input := &qbusiness.ListDocumentsInput{}
        output := &qbusiness.ListDocumentsOutput{}

        mockClient.On("ListDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &qbusiness.ListGroupsInput{}
        output := &qbusiness.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndices", func(t *testing.T) {
        input := &qbusiness.ListIndicesInput{}
        output := &qbusiness.ListIndicesOutput{}

        mockClient.On("ListIndices", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMessages", func(t *testing.T) {
        input := &qbusiness.ListMessagesInput{}
        output := &qbusiness.ListMessagesOutput{}

        mockClient.On("ListMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPluginActions", func(t *testing.T) {
        input := &qbusiness.ListPluginActionsInput{}
        output := &qbusiness.ListPluginActionsOutput{}

        mockClient.On("ListPluginActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPluginActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPluginTypeActions", func(t *testing.T) {
        input := &qbusiness.ListPluginTypeActionsInput{}
        output := &qbusiness.ListPluginTypeActionsOutput{}

        mockClient.On("ListPluginTypeActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPluginTypeActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPluginTypeMetadata", func(t *testing.T) {
        input := &qbusiness.ListPluginTypeMetadataInput{}
        output := &qbusiness.ListPluginTypeMetadataOutput{}

        mockClient.On("ListPluginTypeMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListPluginTypeMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlugins", func(t *testing.T) {
        input := &qbusiness.ListPluginsInput{}
        output := &qbusiness.ListPluginsOutput{}

        mockClient.On("ListPlugins", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlugins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRetrievers", func(t *testing.T) {
        input := &qbusiness.ListRetrieversInput{}
        output := &qbusiness.ListRetrieversOutput{}

        mockClient.On("ListRetrievers", ctx, input).Return(output, nil)

        result, err := mockClient.ListRetrievers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &qbusiness.ListTagsForResourceInput{}
        output := &qbusiness.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebExperiences", func(t *testing.T) {
        input := &qbusiness.ListWebExperiencesInput{}
        output := &qbusiness.ListWebExperiencesOutput{}

        mockClient.On("ListWebExperiences", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebExperiences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFeedback", func(t *testing.T) {
        input := &qbusiness.PutFeedbackInput{}
        output := &qbusiness.PutFeedbackOutput{}

        mockClient.On("PutFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.PutFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGroup", func(t *testing.T) {
        input := &qbusiness.PutGroupInput{}
        output := &qbusiness.PutGroupOutput{}

        mockClient.On("PutGroup", ctx, input).Return(output, nil)

        result, err := mockClient.PutGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchRelevantContent", func(t *testing.T) {
        input := &qbusiness.SearchRelevantContentInput{}
        output := &qbusiness.SearchRelevantContentOutput{}

        mockClient.On("SearchRelevantContent", ctx, input).Return(output, nil)

        result, err := mockClient.SearchRelevantContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataSourceSyncJob", func(t *testing.T) {
        input := &qbusiness.StartDataSourceSyncJobInput{}
        output := &qbusiness.StartDataSourceSyncJobOutput{}

        mockClient.On("StartDataSourceSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataSourceSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDataSourceSyncJob", func(t *testing.T) {
        input := &qbusiness.StopDataSourceSyncJobInput{}
        output := &qbusiness.StopDataSourceSyncJobOutput{}

        mockClient.On("StopDataSourceSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopDataSourceSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &qbusiness.TagResourceInput{}
        output := &qbusiness.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &qbusiness.UntagResourceInput{}
        output := &qbusiness.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &qbusiness.UpdateApplicationInput{}
        output := &qbusiness.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChatControlsConfiguration", func(t *testing.T) {
        input := &qbusiness.UpdateChatControlsConfigurationInput{}
        output := &qbusiness.UpdateChatControlsConfigurationOutput{}

        mockClient.On("UpdateChatControlsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChatControlsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataAccessor", func(t *testing.T) {
        input := &qbusiness.UpdateDataAccessorInput{}
        output := &qbusiness.UpdateDataAccessorOutput{}

        mockClient.On("UpdateDataAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &qbusiness.UpdateDataSourceInput{}
        output := &qbusiness.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIndex", func(t *testing.T) {
        input := &qbusiness.UpdateIndexInput{}
        output := &qbusiness.UpdateIndexOutput{}

        mockClient.On("UpdateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePlugin", func(t *testing.T) {
        input := &qbusiness.UpdatePluginInput{}
        output := &qbusiness.UpdatePluginOutput{}

        mockClient.On("UpdatePlugin", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRetriever", func(t *testing.T) {
        input := &qbusiness.UpdateRetrieverInput{}
        output := &qbusiness.UpdateRetrieverOutput{}

        mockClient.On("UpdateRetriever", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRetriever(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &qbusiness.UpdateUserInput{}
        output := &qbusiness.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebExperience", func(t *testing.T) {
        input := &qbusiness.UpdateWebExperienceInput{}
        output := &qbusiness.UpdateWebExperienceOutput{}

        mockClient.On("UpdateWebExperience", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
