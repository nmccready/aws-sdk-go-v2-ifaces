
package qbusiness_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/qbusiness"
)

// IClient defines the interface for qbusiness
type IClient interface {
 Options() Options 
 AssociatePermission(ctx context.Context, params *AssociatePermissionInput, optFns ...func(*Options)) (*AssociatePermissionOutput, error) 
 BatchDeleteDocument(ctx context.Context, params *BatchDeleteDocumentInput, optFns ...func(*Options)) (*BatchDeleteDocumentOutput, error) 
 BatchPutDocument(ctx context.Context, params *BatchPutDocumentInput, optFns ...func(*Options)) (*BatchPutDocumentOutput, error) 
 CancelSubscription(ctx context.Context, params *CancelSubscriptionInput, optFns ...func(*Options)) (*CancelSubscriptionOutput, error) 
 Chat(ctx context.Context, params *ChatInput, optFns ...func(*Options)) (*ChatOutput, error) 
 ChatSync(ctx context.Context, params *ChatSyncInput, optFns ...func(*Options)) (*ChatSyncOutput, error) 
 CheckDocumentAccess(ctx context.Context, params *CheckDocumentAccessInput, optFns ...func(*Options)) (*CheckDocumentAccessOutput, error) 
 CreateAnonymousWebExperienceUrl(ctx context.Context, params *CreateAnonymousWebExperienceUrlInput, optFns ...func(*Options)) (*CreateAnonymousWebExperienceUrlOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateDataAccessor(ctx context.Context, params *CreateDataAccessorInput, optFns ...func(*Options)) (*CreateDataAccessorOutput, error) 
 CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) 
 CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) 
 CreatePlugin(ctx context.Context, params *CreatePluginInput, optFns ...func(*Options)) (*CreatePluginOutput, error) 
 CreateRetriever(ctx context.Context, params *CreateRetrieverInput, optFns ...func(*Options)) (*CreateRetrieverOutput, error) 
 CreateSubscription(ctx context.Context, params *CreateSubscriptionInput, optFns ...func(*Options)) (*CreateSubscriptionOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 CreateWebExperience(ctx context.Context, params *CreateWebExperienceInput, optFns ...func(*Options)) (*CreateWebExperienceOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteAttachment(ctx context.Context, params *DeleteAttachmentInput, optFns ...func(*Options)) (*DeleteAttachmentOutput, error) 
 DeleteChatControlsConfiguration(ctx context.Context, params *DeleteChatControlsConfigurationInput, optFns ...func(*Options)) (*DeleteChatControlsConfigurationOutput, error) 
 DeleteConversation(ctx context.Context, params *DeleteConversationInput, optFns ...func(*Options)) (*DeleteConversationOutput, error) 
 DeleteDataAccessor(ctx context.Context, params *DeleteDataAccessorInput, optFns ...func(*Options)) (*DeleteDataAccessorOutput, error) 
 DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteIndex(ctx context.Context, params *DeleteIndexInput, optFns ...func(*Options)) (*DeleteIndexOutput, error) 
 DeletePlugin(ctx context.Context, params *DeletePluginInput, optFns ...func(*Options)) (*DeletePluginOutput, error) 
 DeleteRetriever(ctx context.Context, params *DeleteRetrieverInput, optFns ...func(*Options)) (*DeleteRetrieverOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DeleteWebExperience(ctx context.Context, params *DeleteWebExperienceInput, optFns ...func(*Options)) (*DeleteWebExperienceOutput, error) 
 DisassociatePermission(ctx context.Context, params *DisassociatePermissionInput, optFns ...func(*Options)) (*DisassociatePermissionOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetChatControlsConfiguration(ctx context.Context, params *GetChatControlsConfigurationInput, optFns ...func(*Options)) (*GetChatControlsConfigurationOutput, error) 
 GetDataAccessor(ctx context.Context, params *GetDataAccessorInput, optFns ...func(*Options)) (*GetDataAccessorOutput, error) 
 GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetIndex(ctx context.Context, params *GetIndexInput, optFns ...func(*Options)) (*GetIndexOutput, error) 
 GetMedia(ctx context.Context, params *GetMediaInput, optFns ...func(*Options)) (*GetMediaOutput, error) 
 GetPlugin(ctx context.Context, params *GetPluginInput, optFns ...func(*Options)) (*GetPluginOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetRetriever(ctx context.Context, params *GetRetrieverInput, optFns ...func(*Options)) (*GetRetrieverOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetWebExperience(ctx context.Context, params *GetWebExperienceInput, optFns ...func(*Options)) (*GetWebExperienceOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListAttachments(ctx context.Context, params *ListAttachmentsInput, optFns ...func(*Options)) (*ListAttachmentsOutput, error) 
 ListConversations(ctx context.Context, params *ListConversationsInput, optFns ...func(*Options)) (*ListConversationsOutput, error) 
 ListDataAccessors(ctx context.Context, params *ListDataAccessorsInput, optFns ...func(*Options)) (*ListDataAccessorsOutput, error) 
 ListDataSourceSyncJobs(ctx context.Context, params *ListDataSourceSyncJobsInput, optFns ...func(*Options)) (*ListDataSourceSyncJobsOutput, error) 
 ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) 
 ListDocuments(ctx context.Context, params *ListDocumentsInput, optFns ...func(*Options)) (*ListDocumentsOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListIndices(ctx context.Context, params *ListIndicesInput, optFns ...func(*Options)) (*ListIndicesOutput, error) 
 ListMessages(ctx context.Context, params *ListMessagesInput, optFns ...func(*Options)) (*ListMessagesOutput, error) 
 ListPluginActions(ctx context.Context, params *ListPluginActionsInput, optFns ...func(*Options)) (*ListPluginActionsOutput, error) 
 ListPluginTypeActions(ctx context.Context, params *ListPluginTypeActionsInput, optFns ...func(*Options)) (*ListPluginTypeActionsOutput, error) 
 ListPluginTypeMetadata(ctx context.Context, params *ListPluginTypeMetadataInput, optFns ...func(*Options)) (*ListPluginTypeMetadataOutput, error) 
 ListPlugins(ctx context.Context, params *ListPluginsInput, optFns ...func(*Options)) (*ListPluginsOutput, error) 
 ListRetrievers(ctx context.Context, params *ListRetrieversInput, optFns ...func(*Options)) (*ListRetrieversOutput, error) 
 ListSubscriptions(ctx context.Context, params *ListSubscriptionsInput, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebExperiences(ctx context.Context, params *ListWebExperiencesInput, optFns ...func(*Options)) (*ListWebExperiencesOutput, error) 
 PutFeedback(ctx context.Context, params *PutFeedbackInput, optFns ...func(*Options)) (*PutFeedbackOutput, error) 
 PutGroup(ctx context.Context, params *PutGroupInput, optFns ...func(*Options)) (*PutGroupOutput, error) 
 SearchRelevantContent(ctx context.Context, params *SearchRelevantContentInput, optFns ...func(*Options)) (*SearchRelevantContentOutput, error) 
 StartDataSourceSyncJob(ctx context.Context, params *StartDataSourceSyncJobInput, optFns ...func(*Options)) (*StartDataSourceSyncJobOutput, error) 
 StopDataSourceSyncJob(ctx context.Context, params *StopDataSourceSyncJobInput, optFns ...func(*Options)) (*StopDataSourceSyncJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateChatControlsConfiguration(ctx context.Context, params *UpdateChatControlsConfigurationInput, optFns ...func(*Options)) (*UpdateChatControlsConfigurationOutput, error) 
 UpdateDataAccessor(ctx context.Context, params *UpdateDataAccessorInput, optFns ...func(*Options)) (*UpdateDataAccessorOutput, error) 
 UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) 
 UpdateIndex(ctx context.Context, params *UpdateIndexInput, optFns ...func(*Options)) (*UpdateIndexOutput, error) 
 UpdatePlugin(ctx context.Context, params *UpdatePluginInput, optFns ...func(*Options)) (*UpdatePluginOutput, error) 
 UpdateRetriever(ctx context.Context, params *UpdateRetrieverInput, optFns ...func(*Options)) (*UpdateRetrieverOutput, error) 
 UpdateSubscription(ctx context.Context, params *UpdateSubscriptionInput, optFns ...func(*Options)) (*UpdateSubscriptionOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
 UpdateWebExperience(ctx context.Context, params *UpdateWebExperienceInput, optFns ...func(*Options)) (*UpdateWebExperienceOutput, error) 
}
