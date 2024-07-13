package qconnect_test

// tests for the qconnect service interface for this ../../../service/qconnect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qconnect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qconnect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qconnect/qconnect_iface"
	"github.com/stretchr/testify/assert"
)

func TestQconnectServiceCanBeMocked(t *testing.T) {
	var iface qconnect_iface.IClient
	iface = &qconnect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := qconnect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssistant", func(t *testing.T) {
        input := &qconnect.CreateAssistantInput{}
        output := &qconnect.CreateAssistantOutput{}

        mockClient.On("CreateAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssistantAssociation", func(t *testing.T) {
        input := &qconnect.CreateAssistantAssociationInput{}
        output := &qconnect.CreateAssistantAssociationOutput{}

        mockClient.On("CreateAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContent", func(t *testing.T) {
        input := &qconnect.CreateContentInput{}
        output := &qconnect.CreateContentOutput{}

        mockClient.On("CreateContent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContentAssociation", func(t *testing.T) {
        input := &qconnect.CreateContentAssociationInput{}
        output := &qconnect.CreateContentAssociationOutput{}

        mockClient.On("CreateContentAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContentAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKnowledgeBase", func(t *testing.T) {
        input := &qconnect.CreateKnowledgeBaseInput{}
        output := &qconnect.CreateKnowledgeBaseOutput{}

        mockClient.On("CreateKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQuickResponse", func(t *testing.T) {
        input := &qconnect.CreateQuickResponseInput{}
        output := &qconnect.CreateQuickResponseOutput{}

        mockClient.On("CreateQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSession", func(t *testing.T) {
        input := &qconnect.CreateSessionInput{}
        output := &qconnect.CreateSessionOutput{}

        mockClient.On("CreateSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssistant", func(t *testing.T) {
        input := &qconnect.DeleteAssistantInput{}
        output := &qconnect.DeleteAssistantOutput{}

        mockClient.On("DeleteAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssistantAssociation", func(t *testing.T) {
        input := &qconnect.DeleteAssistantAssociationInput{}
        output := &qconnect.DeleteAssistantAssociationOutput{}

        mockClient.On("DeleteAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContent", func(t *testing.T) {
        input := &qconnect.DeleteContentInput{}
        output := &qconnect.DeleteContentOutput{}

        mockClient.On("DeleteContent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContentAssociation", func(t *testing.T) {
        input := &qconnect.DeleteContentAssociationInput{}
        output := &qconnect.DeleteContentAssociationOutput{}

        mockClient.On("DeleteContentAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContentAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImportJob", func(t *testing.T) {
        input := &qconnect.DeleteImportJobInput{}
        output := &qconnect.DeleteImportJobOutput{}

        mockClient.On("DeleteImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKnowledgeBase", func(t *testing.T) {
        input := &qconnect.DeleteKnowledgeBaseInput{}
        output := &qconnect.DeleteKnowledgeBaseOutput{}

        mockClient.On("DeleteKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQuickResponse", func(t *testing.T) {
        input := &qconnect.DeleteQuickResponseInput{}
        output := &qconnect.DeleteQuickResponseOutput{}

        mockClient.On("DeleteQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssistant", func(t *testing.T) {
        input := &qconnect.GetAssistantInput{}
        output := &qconnect.GetAssistantOutput{}

        mockClient.On("GetAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssistantAssociation", func(t *testing.T) {
        input := &qconnect.GetAssistantAssociationInput{}
        output := &qconnect.GetAssistantAssociationOutput{}

        mockClient.On("GetAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContent", func(t *testing.T) {
        input := &qconnect.GetContentInput{}
        output := &qconnect.GetContentOutput{}

        mockClient.On("GetContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContentAssociation", func(t *testing.T) {
        input := &qconnect.GetContentAssociationInput{}
        output := &qconnect.GetContentAssociationOutput{}

        mockClient.On("GetContentAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetContentAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContentSummary", func(t *testing.T) {
        input := &qconnect.GetContentSummaryInput{}
        output := &qconnect.GetContentSummaryOutput{}

        mockClient.On("GetContentSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetContentSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportJob", func(t *testing.T) {
        input := &qconnect.GetImportJobInput{}
        output := &qconnect.GetImportJobOutput{}

        mockClient.On("GetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKnowledgeBase", func(t *testing.T) {
        input := &qconnect.GetKnowledgeBaseInput{}
        output := &qconnect.GetKnowledgeBaseOutput{}

        mockClient.On("GetKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.GetKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQuickResponse", func(t *testing.T) {
        input := &qconnect.GetQuickResponseInput{}
        output := &qconnect.GetQuickResponseOutput{}

        mockClient.On("GetQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendations", func(t *testing.T) {
        input := &qconnect.GetRecommendationsInput{}
        output := &qconnect.GetRecommendationsOutput{}

        mockClient.On("GetRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &qconnect.GetSessionInput{}
        output := &qconnect.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssistantAssociations", func(t *testing.T) {
        input := &qconnect.ListAssistantAssociationsInput{}
        output := &qconnect.ListAssistantAssociationsOutput{}

        mockClient.On("ListAssistantAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssistantAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssistants", func(t *testing.T) {
        input := &qconnect.ListAssistantsInput{}
        output := &qconnect.ListAssistantsOutput{}

        mockClient.On("ListAssistants", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssistants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContentAssociations", func(t *testing.T) {
        input := &qconnect.ListContentAssociationsInput{}
        output := &qconnect.ListContentAssociationsOutput{}

        mockClient.On("ListContentAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListContentAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContents", func(t *testing.T) {
        input := &qconnect.ListContentsInput{}
        output := &qconnect.ListContentsOutput{}

        mockClient.On("ListContents", ctx, input).Return(output, nil)

        result, err := mockClient.ListContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportJobs", func(t *testing.T) {
        input := &qconnect.ListImportJobsInput{}
        output := &qconnect.ListImportJobsOutput{}

        mockClient.On("ListImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKnowledgeBases", func(t *testing.T) {
        input := &qconnect.ListKnowledgeBasesInput{}
        output := &qconnect.ListKnowledgeBasesOutput{}

        mockClient.On("ListKnowledgeBases", ctx, input).Return(output, nil)

        result, err := mockClient.ListKnowledgeBases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQuickResponses", func(t *testing.T) {
        input := &qconnect.ListQuickResponsesInput{}
        output := &qconnect.ListQuickResponsesOutput{}

        mockClient.On("ListQuickResponses", ctx, input).Return(output, nil)

        result, err := mockClient.ListQuickResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &qconnect.ListTagsForResourceInput{}
        output := &qconnect.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyRecommendationsReceived", func(t *testing.T) {
        input := &qconnect.NotifyRecommendationsReceivedInput{}
        output := &qconnect.NotifyRecommendationsReceivedOutput{}

        mockClient.On("NotifyRecommendationsReceived", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyRecommendationsReceived(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFeedback", func(t *testing.T) {
        input := &qconnect.PutFeedbackInput{}
        output := &qconnect.PutFeedbackOutput{}

        mockClient.On("PutFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.PutFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryAssistant", func(t *testing.T) {
        input := &qconnect.QueryAssistantInput{}
        output := &qconnect.QueryAssistantOutput{}

        mockClient.On("QueryAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.QueryAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveKnowledgeBaseTemplateUri", func(t *testing.T) {
        input := &qconnect.RemoveKnowledgeBaseTemplateUriInput{}
        output := &qconnect.RemoveKnowledgeBaseTemplateUriOutput{}

        mockClient.On("RemoveKnowledgeBaseTemplateUri", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveKnowledgeBaseTemplateUri(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchContent", func(t *testing.T) {
        input := &qconnect.SearchContentInput{}
        output := &qconnect.SearchContentOutput{}

        mockClient.On("SearchContent", ctx, input).Return(output, nil)

        result, err := mockClient.SearchContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchQuickResponses", func(t *testing.T) {
        input := &qconnect.SearchQuickResponsesInput{}
        output := &qconnect.SearchQuickResponsesOutput{}

        mockClient.On("SearchQuickResponses", ctx, input).Return(output, nil)

        result, err := mockClient.SearchQuickResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSessions", func(t *testing.T) {
        input := &qconnect.SearchSessionsInput{}
        output := &qconnect.SearchSessionsOutput{}

        mockClient.On("SearchSessions", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContentUpload", func(t *testing.T) {
        input := &qconnect.StartContentUploadInput{}
        output := &qconnect.StartContentUploadOutput{}

        mockClient.On("StartContentUpload", ctx, input).Return(output, nil)

        result, err := mockClient.StartContentUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportJob", func(t *testing.T) {
        input := &qconnect.StartImportJobInput{}
        output := &qconnect.StartImportJobOutput{}

        mockClient.On("StartImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &qconnect.TagResourceInput{}
        output := &qconnect.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &qconnect.UntagResourceInput{}
        output := &qconnect.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContent", func(t *testing.T) {
        input := &qconnect.UpdateContentInput{}
        output := &qconnect.UpdateContentOutput{}

        mockClient.On("UpdateContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKnowledgeBaseTemplateUri", func(t *testing.T) {
        input := &qconnect.UpdateKnowledgeBaseTemplateUriInput{}
        output := &qconnect.UpdateKnowledgeBaseTemplateUriOutput{}

        mockClient.On("UpdateKnowledgeBaseTemplateUri", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKnowledgeBaseTemplateUri(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuickResponse", func(t *testing.T) {
        input := &qconnect.UpdateQuickResponseInput{}
        output := &qconnect.UpdateQuickResponseOutput{}

        mockClient.On("UpdateQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSession", func(t *testing.T) {
        input := &qconnect.UpdateSessionInput{}
        output := &qconnect.UpdateSessionOutput{}

        mockClient.On("UpdateSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
