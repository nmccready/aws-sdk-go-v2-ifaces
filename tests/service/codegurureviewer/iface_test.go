package codegurureviewer_test

// tests for the codegurureviewer service interface for this ../../../service/codegurureviewer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codegurureviewer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codegurureviewer/codegurureviewer_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodegurureviewerServiceCanBeMocked(t *testing.T) {
	var iface codegurureviewer_iface.IClient
	iface = &codegurureviewer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codegurureviewer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateRepository", func(t *testing.T) {
        input := &codegurureviewer.AssociateRepositoryInput{}
        output := &codegurureviewer.AssociateRepositoryOutput{}

        mockClient.On("AssociateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCodeReview", func(t *testing.T) {
        input := &codegurureviewer.CreateCodeReviewInput{}
        output := &codegurureviewer.CreateCodeReviewOutput{}

        mockClient.On("CreateCodeReview", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCodeReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCodeReview", func(t *testing.T) {
        input := &codegurureviewer.DescribeCodeReviewInput{}
        output := &codegurureviewer.DescribeCodeReviewOutput{}

        mockClient.On("DescribeCodeReview", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCodeReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecommendationFeedback", func(t *testing.T) {
        input := &codegurureviewer.DescribeRecommendationFeedbackInput{}
        output := &codegurureviewer.DescribeRecommendationFeedbackOutput{}

        mockClient.On("DescribeRecommendationFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecommendationFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRepositoryAssociation", func(t *testing.T) {
        input := &codegurureviewer.DescribeRepositoryAssociationInput{}
        output := &codegurureviewer.DescribeRepositoryAssociationOutput{}

        mockClient.On("DescribeRepositoryAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRepositoryAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRepository", func(t *testing.T) {
        input := &codegurureviewer.DisassociateRepositoryInput{}
        output := &codegurureviewer.DisassociateRepositoryOutput{}

        mockClient.On("DisassociateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCodeReviews", func(t *testing.T) {
        input := &codegurureviewer.ListCodeReviewsInput{}
        output := &codegurureviewer.ListCodeReviewsOutput{}

        mockClient.On("ListCodeReviews", ctx, input).Return(output, nil)

        result, err := mockClient.ListCodeReviews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendationFeedback", func(t *testing.T) {
        input := &codegurureviewer.ListRecommendationFeedbackInput{}
        output := &codegurureviewer.ListRecommendationFeedbackOutput{}

        mockClient.On("ListRecommendationFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendationFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &codegurureviewer.ListRecommendationsInput{}
        output := &codegurureviewer.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositoryAssociations", func(t *testing.T) {
        input := &codegurureviewer.ListRepositoryAssociationsInput{}
        output := &codegurureviewer.ListRepositoryAssociationsOutput{}

        mockClient.On("ListRepositoryAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositoryAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codegurureviewer.ListTagsForResourceInput{}
        output := &codegurureviewer.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecommendationFeedback", func(t *testing.T) {
        input := &codegurureviewer.PutRecommendationFeedbackInput{}
        output := &codegurureviewer.PutRecommendationFeedbackOutput{}

        mockClient.On("PutRecommendationFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecommendationFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codegurureviewer.TagResourceInput{}
        output := &codegurureviewer.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codegurureviewer.UntagResourceInput{}
        output := &codegurureviewer.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
