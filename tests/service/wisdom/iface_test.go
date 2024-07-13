package wisdom_test

// tests for the wisdom service interface for this ../../../service/wisdom/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wisdom"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wisdom/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wisdom/wisdom_iface"
	"github.com/stretchr/testify/assert"
)

func TestWisdomServiceCanBeMocked(t *testing.T) {
	var iface wisdom_iface.IClient
	iface = &wisdom.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := wisdom.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssistant", func(t *testing.T) {
        input := &wisdom.CreateAssistantInput{}
        output := &wisdom.CreateAssistantOutput{}

        mockClient.On("CreateAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssistantAssociation", func(t *testing.T) {
        input := &wisdom.CreateAssistantAssociationInput{}
        output := &wisdom.CreateAssistantAssociationOutput{}

        mockClient.On("CreateAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContent", func(t *testing.T) {
        input := &wisdom.CreateContentInput{}
        output := &wisdom.CreateContentOutput{}

        mockClient.On("CreateContent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKnowledgeBase", func(t *testing.T) {
        input := &wisdom.CreateKnowledgeBaseInput{}
        output := &wisdom.CreateKnowledgeBaseOutput{}

        mockClient.On("CreateKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQuickResponse", func(t *testing.T) {
        input := &wisdom.CreateQuickResponseInput{}
        output := &wisdom.CreateQuickResponseOutput{}

        mockClient.On("CreateQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSession", func(t *testing.T) {
        input := &wisdom.CreateSessionInput{}
        output := &wisdom.CreateSessionOutput{}

        mockClient.On("CreateSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssistant", func(t *testing.T) {
        input := &wisdom.DeleteAssistantInput{}
        output := &wisdom.DeleteAssistantOutput{}

        mockClient.On("DeleteAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssistantAssociation", func(t *testing.T) {
        input := &wisdom.DeleteAssistantAssociationInput{}
        output := &wisdom.DeleteAssistantAssociationOutput{}

        mockClient.On("DeleteAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContent", func(t *testing.T) {
        input := &wisdom.DeleteContentInput{}
        output := &wisdom.DeleteContentOutput{}

        mockClient.On("DeleteContent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImportJob", func(t *testing.T) {
        input := &wisdom.DeleteImportJobInput{}
        output := &wisdom.DeleteImportJobOutput{}

        mockClient.On("DeleteImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKnowledgeBase", func(t *testing.T) {
        input := &wisdom.DeleteKnowledgeBaseInput{}
        output := &wisdom.DeleteKnowledgeBaseOutput{}

        mockClient.On("DeleteKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQuickResponse", func(t *testing.T) {
        input := &wisdom.DeleteQuickResponseInput{}
        output := &wisdom.DeleteQuickResponseOutput{}

        mockClient.On("DeleteQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssistant", func(t *testing.T) {
        input := &wisdom.GetAssistantInput{}
        output := &wisdom.GetAssistantOutput{}

        mockClient.On("GetAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssistantAssociation", func(t *testing.T) {
        input := &wisdom.GetAssistantAssociationInput{}
        output := &wisdom.GetAssistantAssociationOutput{}

        mockClient.On("GetAssistantAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssistantAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContent", func(t *testing.T) {
        input := &wisdom.GetContentInput{}
        output := &wisdom.GetContentOutput{}

        mockClient.On("GetContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContentSummary", func(t *testing.T) {
        input := &wisdom.GetContentSummaryInput{}
        output := &wisdom.GetContentSummaryOutput{}

        mockClient.On("GetContentSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetContentSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportJob", func(t *testing.T) {
        input := &wisdom.GetImportJobInput{}
        output := &wisdom.GetImportJobOutput{}

        mockClient.On("GetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKnowledgeBase", func(t *testing.T) {
        input := &wisdom.GetKnowledgeBaseInput{}
        output := &wisdom.GetKnowledgeBaseOutput{}

        mockClient.On("GetKnowledgeBase", ctx, input).Return(output, nil)

        result, err := mockClient.GetKnowledgeBase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQuickResponse", func(t *testing.T) {
        input := &wisdom.GetQuickResponseInput{}
        output := &wisdom.GetQuickResponseOutput{}

        mockClient.On("GetQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendations", func(t *testing.T) {
        input := &wisdom.GetRecommendationsInput{}
        output := &wisdom.GetRecommendationsOutput{}

        mockClient.On("GetRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &wisdom.GetSessionInput{}
        output := &wisdom.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssistantAssociations", func(t *testing.T) {
        input := &wisdom.ListAssistantAssociationsInput{}
        output := &wisdom.ListAssistantAssociationsOutput{}

        mockClient.On("ListAssistantAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssistantAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssistants", func(t *testing.T) {
        input := &wisdom.ListAssistantsInput{}
        output := &wisdom.ListAssistantsOutput{}

        mockClient.On("ListAssistants", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssistants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContents", func(t *testing.T) {
        input := &wisdom.ListContentsInput{}
        output := &wisdom.ListContentsOutput{}

        mockClient.On("ListContents", ctx, input).Return(output, nil)

        result, err := mockClient.ListContents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportJobs", func(t *testing.T) {
        input := &wisdom.ListImportJobsInput{}
        output := &wisdom.ListImportJobsOutput{}

        mockClient.On("ListImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKnowledgeBases", func(t *testing.T) {
        input := &wisdom.ListKnowledgeBasesInput{}
        output := &wisdom.ListKnowledgeBasesOutput{}

        mockClient.On("ListKnowledgeBases", ctx, input).Return(output, nil)

        result, err := mockClient.ListKnowledgeBases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQuickResponses", func(t *testing.T) {
        input := &wisdom.ListQuickResponsesInput{}
        output := &wisdom.ListQuickResponsesOutput{}

        mockClient.On("ListQuickResponses", ctx, input).Return(output, nil)

        result, err := mockClient.ListQuickResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &wisdom.ListTagsForResourceInput{}
        output := &wisdom.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyRecommendationsReceived", func(t *testing.T) {
        input := &wisdom.NotifyRecommendationsReceivedInput{}
        output := &wisdom.NotifyRecommendationsReceivedOutput{}

        mockClient.On("NotifyRecommendationsReceived", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyRecommendationsReceived(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryAssistant", func(t *testing.T) {
        input := &wisdom.QueryAssistantInput{}
        output := &wisdom.QueryAssistantOutput{}

        mockClient.On("QueryAssistant", ctx, input).Return(output, nil)

        result, err := mockClient.QueryAssistant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveKnowledgeBaseTemplateUri", func(t *testing.T) {
        input := &wisdom.RemoveKnowledgeBaseTemplateUriInput{}
        output := &wisdom.RemoveKnowledgeBaseTemplateUriOutput{}

        mockClient.On("RemoveKnowledgeBaseTemplateUri", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveKnowledgeBaseTemplateUri(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchContent", func(t *testing.T) {
        input := &wisdom.SearchContentInput{}
        output := &wisdom.SearchContentOutput{}

        mockClient.On("SearchContent", ctx, input).Return(output, nil)

        result, err := mockClient.SearchContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchQuickResponses", func(t *testing.T) {
        input := &wisdom.SearchQuickResponsesInput{}
        output := &wisdom.SearchQuickResponsesOutput{}

        mockClient.On("SearchQuickResponses", ctx, input).Return(output, nil)

        result, err := mockClient.SearchQuickResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSessions", func(t *testing.T) {
        input := &wisdom.SearchSessionsInput{}
        output := &wisdom.SearchSessionsOutput{}

        mockClient.On("SearchSessions", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContentUpload", func(t *testing.T) {
        input := &wisdom.StartContentUploadInput{}
        output := &wisdom.StartContentUploadOutput{}

        mockClient.On("StartContentUpload", ctx, input).Return(output, nil)

        result, err := mockClient.StartContentUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportJob", func(t *testing.T) {
        input := &wisdom.StartImportJobInput{}
        output := &wisdom.StartImportJobOutput{}

        mockClient.On("StartImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &wisdom.TagResourceInput{}
        output := &wisdom.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &wisdom.UntagResourceInput{}
        output := &wisdom.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContent", func(t *testing.T) {
        input := &wisdom.UpdateContentInput{}
        output := &wisdom.UpdateContentOutput{}

        mockClient.On("UpdateContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKnowledgeBaseTemplateUri", func(t *testing.T) {
        input := &wisdom.UpdateKnowledgeBaseTemplateUriInput{}
        output := &wisdom.UpdateKnowledgeBaseTemplateUriOutput{}

        mockClient.On("UpdateKnowledgeBaseTemplateUri", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKnowledgeBaseTemplateUri(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuickResponse", func(t *testing.T) {
        input := &wisdom.UpdateQuickResponseInput{}
        output := &wisdom.UpdateQuickResponseOutput{}

        mockClient.On("UpdateQuickResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuickResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
