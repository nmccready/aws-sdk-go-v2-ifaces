
package wisdom

import (
    "github.com/aws/aws-sdk-go-v2/service/wisdom"
)

// IClient defines the interface for wisdom
type IClient interface {
 Options() Options 
 CreateAssistant(ctx context.Context, params *CreateAssistantInput, optFns ...func(*Options)) (*CreateAssistantOutput, error) 
 CreateAssistantAssociation(ctx context.Context, params *CreateAssistantAssociationInput, optFns ...func(*Options)) (*CreateAssistantAssociationOutput, error) 
 CreateContent(ctx context.Context, params *CreateContentInput, optFns ...func(*Options)) (*CreateContentOutput, error) 
 CreateKnowledgeBase(ctx context.Context, params *CreateKnowledgeBaseInput, optFns ...func(*Options)) (*CreateKnowledgeBaseOutput, error) 
 CreateQuickResponse(ctx context.Context, params *CreateQuickResponseInput, optFns ...func(*Options)) (*CreateQuickResponseOutput, error) 
 CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) 
 DeleteAssistant(ctx context.Context, params *DeleteAssistantInput, optFns ...func(*Options)) (*DeleteAssistantOutput, error) 
 DeleteAssistantAssociation(ctx context.Context, params *DeleteAssistantAssociationInput, optFns ...func(*Options)) (*DeleteAssistantAssociationOutput, error) 
 DeleteContent(ctx context.Context, params *DeleteContentInput, optFns ...func(*Options)) (*DeleteContentOutput, error) 
 DeleteImportJob(ctx context.Context, params *DeleteImportJobInput, optFns ...func(*Options)) (*DeleteImportJobOutput, error) 
 DeleteKnowledgeBase(ctx context.Context, params *DeleteKnowledgeBaseInput, optFns ...func(*Options)) (*DeleteKnowledgeBaseOutput, error) 
 DeleteQuickResponse(ctx context.Context, params *DeleteQuickResponseInput, optFns ...func(*Options)) (*DeleteQuickResponseOutput, error) 
 GetAssistant(ctx context.Context, params *GetAssistantInput, optFns ...func(*Options)) (*GetAssistantOutput, error) 
 GetAssistantAssociation(ctx context.Context, params *GetAssistantAssociationInput, optFns ...func(*Options)) (*GetAssistantAssociationOutput, error) 
 GetContent(ctx context.Context, params *GetContentInput, optFns ...func(*Options)) (*GetContentOutput, error) 
 GetContentSummary(ctx context.Context, params *GetContentSummaryInput, optFns ...func(*Options)) (*GetContentSummaryOutput, error) 
 GetImportJob(ctx context.Context, params *GetImportJobInput, optFns ...func(*Options)) (*GetImportJobOutput, error) 
 GetKnowledgeBase(ctx context.Context, params *GetKnowledgeBaseInput, optFns ...func(*Options)) (*GetKnowledgeBaseOutput, error) 
 GetQuickResponse(ctx context.Context, params *GetQuickResponseInput, optFns ...func(*Options)) (*GetQuickResponseOutput, error) 
 GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 ListAssistantAssociations(ctx context.Context, params *ListAssistantAssociationsInput, optFns ...func(*Options)) (*ListAssistantAssociationsOutput, error) 
 ListAssistants(ctx context.Context, params *ListAssistantsInput, optFns ...func(*Options)) (*ListAssistantsOutput, error) 
 ListContents(ctx context.Context, params *ListContentsInput, optFns ...func(*Options)) (*ListContentsOutput, error) 
 ListImportJobs(ctx context.Context, params *ListImportJobsInput, optFns ...func(*Options)) (*ListImportJobsOutput, error) 
 ListKnowledgeBases(ctx context.Context, params *ListKnowledgeBasesInput, optFns ...func(*Options)) (*ListKnowledgeBasesOutput, error) 
 ListQuickResponses(ctx context.Context, params *ListQuickResponsesInput, optFns ...func(*Options)) (*ListQuickResponsesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 NotifyRecommendationsReceived(ctx context.Context, params *NotifyRecommendationsReceivedInput, optFns ...func(*Options)) (*NotifyRecommendationsReceivedOutput, error) 
 QueryAssistant(ctx context.Context, params *QueryAssistantInput, optFns ...func(*Options)) (*QueryAssistantOutput, error) 
 RemoveKnowledgeBaseTemplateUri(ctx context.Context, params *RemoveKnowledgeBaseTemplateUriInput, optFns ...func(*Options)) (*RemoveKnowledgeBaseTemplateUriOutput, error) 
 SearchContent(ctx context.Context, params *SearchContentInput, optFns ...func(*Options)) (*SearchContentOutput, error) 
 SearchQuickResponses(ctx context.Context, params *SearchQuickResponsesInput, optFns ...func(*Options)) (*SearchQuickResponsesOutput, error) 
 SearchSessions(ctx context.Context, params *SearchSessionsInput, optFns ...func(*Options)) (*SearchSessionsOutput, error) 
 StartContentUpload(ctx context.Context, params *StartContentUploadInput, optFns ...func(*Options)) (*StartContentUploadOutput, error) 
 StartImportJob(ctx context.Context, params *StartImportJobInput, optFns ...func(*Options)) (*StartImportJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateContent(ctx context.Context, params *UpdateContentInput, optFns ...func(*Options)) (*UpdateContentOutput, error) 
 UpdateKnowledgeBaseTemplateUri(ctx context.Context, params *UpdateKnowledgeBaseTemplateUriInput, optFns ...func(*Options)) (*UpdateKnowledgeBaseTemplateUriOutput, error) 
 UpdateQuickResponse(ctx context.Context, params *UpdateQuickResponseInput, optFns ...func(*Options)) (*UpdateQuickResponseOutput, error) 
}
