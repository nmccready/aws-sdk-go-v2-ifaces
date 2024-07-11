
package codegurureviewer

import (
    "github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
)

// IClient defines the interface for codegurureviewer
type IClient interface {
 Options() Options 
 AssociateRepository(ctx context.Context, params *AssociateRepositoryInput, optFns ...func(*Options)) (*AssociateRepositoryOutput, error) 
 CreateCodeReview(ctx context.Context, params *CreateCodeReviewInput, optFns ...func(*Options)) (*CreateCodeReviewOutput, error) 
 DescribeCodeReview(ctx context.Context, params *DescribeCodeReviewInput, optFns ...func(*Options)) (*DescribeCodeReviewOutput, error) 
 DescribeRecommendationFeedback(ctx context.Context, params *DescribeRecommendationFeedbackInput, optFns ...func(*Options)) (*DescribeRecommendationFeedbackOutput, error) 
 DescribeRepositoryAssociation(ctx context.Context, params *DescribeRepositoryAssociationInput, optFns ...func(*Options)) (*DescribeRepositoryAssociationOutput, error) 
 DisassociateRepository(ctx context.Context, params *DisassociateRepositoryInput, optFns ...func(*Options)) (*DisassociateRepositoryOutput, error) 
 ListCodeReviews(ctx context.Context, params *ListCodeReviewsInput, optFns ...func(*Options)) (*ListCodeReviewsOutput, error) 
 ListRecommendationFeedback(ctx context.Context, params *ListRecommendationFeedbackInput, optFns ...func(*Options)) (*ListRecommendationFeedbackOutput, error) 
 ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) 
 ListRepositoryAssociations(ctx context.Context, params *ListRepositoryAssociationsInput, optFns ...func(*Options)) (*ListRepositoryAssociationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutRecommendationFeedback(ctx context.Context, params *PutRecommendationFeedbackInput, optFns ...func(*Options)) (*PutRecommendationFeedbackOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
