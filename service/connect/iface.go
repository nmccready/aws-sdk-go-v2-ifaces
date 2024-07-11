
package connect

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/connect"
)

// IClient defines the interface for connect
type IClient interface {
 Options() Options 
 ActivateEvaluationForm(ctx context.Context, params *ActivateEvaluationFormInput, optFns ...func(*Options)) (*ActivateEvaluationFormOutput, error) 
 AssociateAnalyticsDataSet(ctx context.Context, params *AssociateAnalyticsDataSetInput, optFns ...func(*Options)) (*AssociateAnalyticsDataSetOutput, error) 
 AssociateApprovedOrigin(ctx context.Context, params *AssociateApprovedOriginInput, optFns ...func(*Options)) (*AssociateApprovedOriginOutput, error) 
 AssociateBot(ctx context.Context, params *AssociateBotInput, optFns ...func(*Options)) (*AssociateBotOutput, error) 
 AssociateDefaultVocabulary(ctx context.Context, params *AssociateDefaultVocabularyInput, optFns ...func(*Options)) (*AssociateDefaultVocabularyOutput, error) 
 AssociateFlow(ctx context.Context, params *AssociateFlowInput, optFns ...func(*Options)) (*AssociateFlowOutput, error) 
 AssociateInstanceStorageConfig(ctx context.Context, params *AssociateInstanceStorageConfigInput, optFns ...func(*Options)) (*AssociateInstanceStorageConfigOutput, error) 
 AssociateLambdaFunction(ctx context.Context, params *AssociateLambdaFunctionInput, optFns ...func(*Options)) (*AssociateLambdaFunctionOutput, error) 
 AssociateLexBot(ctx context.Context, params *AssociateLexBotInput, optFns ...func(*Options)) (*AssociateLexBotOutput, error) 
 AssociatePhoneNumberContactFlow(ctx context.Context, params *AssociatePhoneNumberContactFlowInput, optFns ...func(*Options)) (*AssociatePhoneNumberContactFlowOutput, error) 
 AssociateQueueQuickConnects(ctx context.Context, params *AssociateQueueQuickConnectsInput, optFns ...func(*Options)) (*AssociateQueueQuickConnectsOutput, error) 
 AssociateRoutingProfileQueues(ctx context.Context, params *AssociateRoutingProfileQueuesInput, optFns ...func(*Options)) (*AssociateRoutingProfileQueuesOutput, error) 
 AssociateSecurityKey(ctx context.Context, params *AssociateSecurityKeyInput, optFns ...func(*Options)) (*AssociateSecurityKeyOutput, error) 
 AssociateTrafficDistributionGroupUser(ctx context.Context, params *AssociateTrafficDistributionGroupUserInput, optFns ...func(*Options)) (*AssociateTrafficDistributionGroupUserOutput, error) 
 AssociateUserProficiencies(ctx context.Context, params *AssociateUserProficienciesInput, optFns ...func(*Options)) (*AssociateUserProficienciesOutput, error) 
 BatchAssociateAnalyticsDataSet(ctx context.Context, params *BatchAssociateAnalyticsDataSetInput, optFns ...func(*Options)) (*BatchAssociateAnalyticsDataSetOutput, error) 
 BatchDisassociateAnalyticsDataSet(ctx context.Context, params *BatchDisassociateAnalyticsDataSetInput, optFns ...func(*Options)) (*BatchDisassociateAnalyticsDataSetOutput, error) 
 BatchGetAttachedFileMetadata(ctx context.Context, params *BatchGetAttachedFileMetadataInput, optFns ...func(*Options)) (*BatchGetAttachedFileMetadataOutput, error) 
 BatchGetFlowAssociation(ctx context.Context, params *BatchGetFlowAssociationInput, optFns ...func(*Options)) (*BatchGetFlowAssociationOutput, error) 
 BatchPutContact(ctx context.Context, params *BatchPutContactInput, optFns ...func(*Options)) (*BatchPutContactOutput, error) 
 ClaimPhoneNumber(ctx context.Context, params *ClaimPhoneNumberInput, optFns ...func(*Options)) (*ClaimPhoneNumberOutput, error) 
 CompleteAttachedFileUpload(ctx context.Context, params *CompleteAttachedFileUploadInput, optFns ...func(*Options)) (*CompleteAttachedFileUploadOutput, error) 
 CreateAgentStatus(ctx context.Context, params *CreateAgentStatusInput, optFns ...func(*Options)) (*CreateAgentStatusOutput, error) 
 CreateContactFlow(ctx context.Context, params *CreateContactFlowInput, optFns ...func(*Options)) (*CreateContactFlowOutput, error) 
 CreateContactFlowModule(ctx context.Context, params *CreateContactFlowModuleInput, optFns ...func(*Options)) (*CreateContactFlowModuleOutput, error) 
 CreateEvaluationForm(ctx context.Context, params *CreateEvaluationFormInput, optFns ...func(*Options)) (*CreateEvaluationFormOutput, error) 
 CreateHoursOfOperation(ctx context.Context, params *CreateHoursOfOperationInput, optFns ...func(*Options)) (*CreateHoursOfOperationOutput, error) 
 CreateInstance(ctx context.Context, params *CreateInstanceInput, optFns ...func(*Options)) (*CreateInstanceOutput, error) 
 CreateIntegrationAssociation(ctx context.Context, params *CreateIntegrationAssociationInput, optFns ...func(*Options)) (*CreateIntegrationAssociationOutput, error) 
 CreateParticipant(ctx context.Context, params *CreateParticipantInput, optFns ...func(*Options)) (*CreateParticipantOutput, error) 
 CreatePersistentContactAssociation(ctx context.Context, params *CreatePersistentContactAssociationInput, optFns ...func(*Options)) (*CreatePersistentContactAssociationOutput, error) 
 CreatePredefinedAttribute(ctx context.Context, params *CreatePredefinedAttributeInput, optFns ...func(*Options)) (*CreatePredefinedAttributeOutput, error) 
 CreatePrompt(ctx context.Context, params *CreatePromptInput, optFns ...func(*Options)) (*CreatePromptOutput, error) 
 CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) 
 CreateQuickConnect(ctx context.Context, params *CreateQuickConnectInput, optFns ...func(*Options)) (*CreateQuickConnectOutput, error) 
 CreateRoutingProfile(ctx context.Context, params *CreateRoutingProfileInput, optFns ...func(*Options)) (*CreateRoutingProfileOutput, error) 
 CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) 
 CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) 
 CreateTaskTemplate(ctx context.Context, params *CreateTaskTemplateInput, optFns ...func(*Options)) (*CreateTaskTemplateOutput, error) 
 CreateTrafficDistributionGroup(ctx context.Context, params *CreateTrafficDistributionGroupInput, optFns ...func(*Options)) (*CreateTrafficDistributionGroupOutput, error) 
 CreateUseCase(ctx context.Context, params *CreateUseCaseInput, optFns ...func(*Options)) (*CreateUseCaseOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 CreateUserHierarchyGroup(ctx context.Context, params *CreateUserHierarchyGroupInput, optFns ...func(*Options)) (*CreateUserHierarchyGroupOutput, error) 
 CreateView(ctx context.Context, params *CreateViewInput, optFns ...func(*Options)) (*CreateViewOutput, error) 
 CreateViewVersion(ctx context.Context, params *CreateViewVersionInput, optFns ...func(*Options)) (*CreateViewVersionOutput, error) 
 CreateVocabulary(ctx context.Context, params *CreateVocabularyInput, optFns ...func(*Options)) (*CreateVocabularyOutput, error) 
 DeactivateEvaluationForm(ctx context.Context, params *DeactivateEvaluationFormInput, optFns ...func(*Options)) (*DeactivateEvaluationFormOutput, error) 
 DeleteAttachedFile(ctx context.Context, params *DeleteAttachedFileInput, optFns ...func(*Options)) (*DeleteAttachedFileOutput, error) 
 DeleteContactEvaluation(ctx context.Context, params *DeleteContactEvaluationInput, optFns ...func(*Options)) (*DeleteContactEvaluationOutput, error) 
 DeleteContactFlow(ctx context.Context, params *DeleteContactFlowInput, optFns ...func(*Options)) (*DeleteContactFlowOutput, error) 
 DeleteContactFlowModule(ctx context.Context, params *DeleteContactFlowModuleInput, optFns ...func(*Options)) (*DeleteContactFlowModuleOutput, error) 
 DeleteEvaluationForm(ctx context.Context, params *DeleteEvaluationFormInput, optFns ...func(*Options)) (*DeleteEvaluationFormOutput, error) 
 DeleteHoursOfOperation(ctx context.Context, params *DeleteHoursOfOperationInput, optFns ...func(*Options)) (*DeleteHoursOfOperationOutput, error) 
 DeleteInstance(ctx context.Context, params *DeleteInstanceInput, optFns ...func(*Options)) (*DeleteInstanceOutput, error) 
 DeleteIntegrationAssociation(ctx context.Context, params *DeleteIntegrationAssociationInput, optFns ...func(*Options)) (*DeleteIntegrationAssociationOutput, error) 
 DeletePredefinedAttribute(ctx context.Context, params *DeletePredefinedAttributeInput, optFns ...func(*Options)) (*DeletePredefinedAttributeOutput, error) 
 DeletePrompt(ctx context.Context, params *DeletePromptInput, optFns ...func(*Options)) (*DeletePromptOutput, error) 
 DeleteQueue(ctx context.Context, params *DeleteQueueInput, optFns ...func(*Options)) (*DeleteQueueOutput, error) 
 DeleteQuickConnect(ctx context.Context, params *DeleteQuickConnectInput, optFns ...func(*Options)) (*DeleteQuickConnectOutput, error) 
 DeleteRoutingProfile(ctx context.Context, params *DeleteRoutingProfileInput, optFns ...func(*Options)) (*DeleteRoutingProfileOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 DeleteSecurityProfile(ctx context.Context, params *DeleteSecurityProfileInput, optFns ...func(*Options)) (*DeleteSecurityProfileOutput, error) 
 DeleteTaskTemplate(ctx context.Context, params *DeleteTaskTemplateInput, optFns ...func(*Options)) (*DeleteTaskTemplateOutput, error) 
 DeleteTrafficDistributionGroup(ctx context.Context, params *DeleteTrafficDistributionGroupInput, optFns ...func(*Options)) (*DeleteTrafficDistributionGroupOutput, error) 
 DeleteUseCase(ctx context.Context, params *DeleteUseCaseInput, optFns ...func(*Options)) (*DeleteUseCaseOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DeleteUserHierarchyGroup(ctx context.Context, params *DeleteUserHierarchyGroupInput, optFns ...func(*Options)) (*DeleteUserHierarchyGroupOutput, error) 
 DeleteView(ctx context.Context, params *DeleteViewInput, optFns ...func(*Options)) (*DeleteViewOutput, error) 
 DeleteViewVersion(ctx context.Context, params *DeleteViewVersionInput, optFns ...func(*Options)) (*DeleteViewVersionOutput, error) 
 DeleteVocabulary(ctx context.Context, params *DeleteVocabularyInput, optFns ...func(*Options)) (*DeleteVocabularyOutput, error) 
 DescribeAgentStatus(ctx context.Context, params *DescribeAgentStatusInput, optFns ...func(*Options)) (*DescribeAgentStatusOutput, error) 
 DescribeAuthenticationProfile(ctx context.Context, params *DescribeAuthenticationProfileInput, optFns ...func(*Options)) (*DescribeAuthenticationProfileOutput, error) 
 DescribeContact(ctx context.Context, params *DescribeContactInput, optFns ...func(*Options)) (*DescribeContactOutput, error) 
 DescribeContactEvaluation(ctx context.Context, params *DescribeContactEvaluationInput, optFns ...func(*Options)) (*DescribeContactEvaluationOutput, error) 
 DescribeContactFlow(ctx context.Context, params *DescribeContactFlowInput, optFns ...func(*Options)) (*DescribeContactFlowOutput, error) 
 DescribeContactFlowModule(ctx context.Context, params *DescribeContactFlowModuleInput, optFns ...func(*Options)) (*DescribeContactFlowModuleOutput, error) 
 DescribeEvaluationForm(ctx context.Context, params *DescribeEvaluationFormInput, optFns ...func(*Options)) (*DescribeEvaluationFormOutput, error) 
 DescribeHoursOfOperation(ctx context.Context, params *DescribeHoursOfOperationInput, optFns ...func(*Options)) (*DescribeHoursOfOperationOutput, error) 
 DescribeInstance(ctx context.Context, params *DescribeInstanceInput, optFns ...func(*Options)) (*DescribeInstanceOutput, error) 
 DescribeInstanceAttribute(ctx context.Context, params *DescribeInstanceAttributeInput, optFns ...func(*Options)) (*DescribeInstanceAttributeOutput, error) 
 DescribeInstanceStorageConfig(ctx context.Context, params *DescribeInstanceStorageConfigInput, optFns ...func(*Options)) (*DescribeInstanceStorageConfigOutput, error) 
 DescribePhoneNumber(ctx context.Context, params *DescribePhoneNumberInput, optFns ...func(*Options)) (*DescribePhoneNumberOutput, error) 
 DescribePredefinedAttribute(ctx context.Context, params *DescribePredefinedAttributeInput, optFns ...func(*Options)) (*DescribePredefinedAttributeOutput, error) 
 DescribePrompt(ctx context.Context, params *DescribePromptInput, optFns ...func(*Options)) (*DescribePromptOutput, error) 
 DescribeQueue(ctx context.Context, params *DescribeQueueInput, optFns ...func(*Options)) (*DescribeQueueOutput, error) 
 DescribeQuickConnect(ctx context.Context, params *DescribeQuickConnectInput, optFns ...func(*Options)) (*DescribeQuickConnectOutput, error) 
 DescribeRoutingProfile(ctx context.Context, params *DescribeRoutingProfileInput, optFns ...func(*Options)) (*DescribeRoutingProfileOutput, error) 
 DescribeRule(ctx context.Context, params *DescribeRuleInput, optFns ...func(*Options)) (*DescribeRuleOutput, error) 
 DescribeSecurityProfile(ctx context.Context, params *DescribeSecurityProfileInput, optFns ...func(*Options)) (*DescribeSecurityProfileOutput, error) 
 DescribeTrafficDistributionGroup(ctx context.Context, params *DescribeTrafficDistributionGroupInput, optFns ...func(*Options)) (*DescribeTrafficDistributionGroupOutput, error) 
 DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) 
 DescribeUserHierarchyGroup(ctx context.Context, params *DescribeUserHierarchyGroupInput, optFns ...func(*Options)) (*DescribeUserHierarchyGroupOutput, error) 
 DescribeUserHierarchyStructure(ctx context.Context, params *DescribeUserHierarchyStructureInput, optFns ...func(*Options)) (*DescribeUserHierarchyStructureOutput, error) 
 DescribeView(ctx context.Context, params *DescribeViewInput, optFns ...func(*Options)) (*DescribeViewOutput, error) 
 DescribeVocabulary(ctx context.Context, params *DescribeVocabularyInput, optFns ...func(*Options)) (*DescribeVocabularyOutput, error) 
 DisassociateAnalyticsDataSet(ctx context.Context, params *DisassociateAnalyticsDataSetInput, optFns ...func(*Options)) (*DisassociateAnalyticsDataSetOutput, error) 
 DisassociateApprovedOrigin(ctx context.Context, params *DisassociateApprovedOriginInput, optFns ...func(*Options)) (*DisassociateApprovedOriginOutput, error) 
 DisassociateBot(ctx context.Context, params *DisassociateBotInput, optFns ...func(*Options)) (*DisassociateBotOutput, error) 
 DisassociateFlow(ctx context.Context, params *DisassociateFlowInput, optFns ...func(*Options)) (*DisassociateFlowOutput, error) 
 DisassociateInstanceStorageConfig(ctx context.Context, params *DisassociateInstanceStorageConfigInput, optFns ...func(*Options)) (*DisassociateInstanceStorageConfigOutput, error) 
 DisassociateLambdaFunction(ctx context.Context, params *DisassociateLambdaFunctionInput, optFns ...func(*Options)) (*DisassociateLambdaFunctionOutput, error) 
 DisassociateLexBot(ctx context.Context, params *DisassociateLexBotInput, optFns ...func(*Options)) (*DisassociateLexBotOutput, error) 
 DisassociatePhoneNumberContactFlow(ctx context.Context, params *DisassociatePhoneNumberContactFlowInput, optFns ...func(*Options)) (*DisassociatePhoneNumberContactFlowOutput, error) 
 DisassociateQueueQuickConnects(ctx context.Context, params *DisassociateQueueQuickConnectsInput, optFns ...func(*Options)) (*DisassociateQueueQuickConnectsOutput, error) 
 DisassociateRoutingProfileQueues(ctx context.Context, params *DisassociateRoutingProfileQueuesInput, optFns ...func(*Options)) (*DisassociateRoutingProfileQueuesOutput, error) 
 DisassociateSecurityKey(ctx context.Context, params *DisassociateSecurityKeyInput, optFns ...func(*Options)) (*DisassociateSecurityKeyOutput, error) 
 DisassociateTrafficDistributionGroupUser(ctx context.Context, params *DisassociateTrafficDistributionGroupUserInput, optFns ...func(*Options)) (*DisassociateTrafficDistributionGroupUserOutput, error) 
 DisassociateUserProficiencies(ctx context.Context, params *DisassociateUserProficienciesInput, optFns ...func(*Options)) (*DisassociateUserProficienciesOutput, error) 
 DismissUserContact(ctx context.Context, params *DismissUserContactInput, optFns ...func(*Options)) (*DismissUserContactOutput, error) 
 GetAttachedFile(ctx context.Context, params *GetAttachedFileInput, optFns ...func(*Options)) (*GetAttachedFileOutput, error) 
 GetContactAttributes(ctx context.Context, params *GetContactAttributesInput, optFns ...func(*Options)) (*GetContactAttributesOutput, error) 
 GetCurrentMetricData(ctx context.Context, params *GetCurrentMetricDataInput, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) 
 GetCurrentUserData(ctx context.Context, params *GetCurrentUserDataInput, optFns ...func(*Options)) (*GetCurrentUserDataOutput, error) 
 GetFederationToken(ctx context.Context, params *GetFederationTokenInput, optFns ...func(*Options)) (*GetFederationTokenOutput, error) 
 GetFlowAssociation(ctx context.Context, params *GetFlowAssociationInput, optFns ...func(*Options)) (*GetFlowAssociationOutput, error) 
 GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) 
 GetMetricDataV2(ctx context.Context, params *GetMetricDataV2Input, optFns ...func(*Options)) (*GetMetricDataV2Output, error) 
 GetPromptFile(ctx context.Context, params *GetPromptFileInput, optFns ...func(*Options)) (*GetPromptFileOutput, error) 
 GetTaskTemplate(ctx context.Context, params *GetTaskTemplateInput, optFns ...func(*Options)) (*GetTaskTemplateOutput, error) 
 GetTrafficDistribution(ctx context.Context, params *GetTrafficDistributionInput, optFns ...func(*Options)) (*GetTrafficDistributionOutput, error) 
 ImportPhoneNumber(ctx context.Context, params *ImportPhoneNumberInput, optFns ...func(*Options)) (*ImportPhoneNumberOutput, error) 
 ListAgentStatuses(ctx context.Context, params *ListAgentStatusesInput, optFns ...func(*Options)) (*ListAgentStatusesOutput, error) 
 ListAnalyticsDataAssociations(ctx context.Context, params *ListAnalyticsDataAssociationsInput, optFns ...func(*Options)) (*ListAnalyticsDataAssociationsOutput, error) 
 ListApprovedOrigins(ctx context.Context, params *ListApprovedOriginsInput, optFns ...func(*Options)) (*ListApprovedOriginsOutput, error) 
 ListAuthenticationProfiles(ctx context.Context, params *ListAuthenticationProfilesInput, optFns ...func(*Options)) (*ListAuthenticationProfilesOutput, error) 
 ListBots(ctx context.Context, params *ListBotsInput, optFns ...func(*Options)) (*ListBotsOutput, error) 
 ListContactEvaluations(ctx context.Context, params *ListContactEvaluationsInput, optFns ...func(*Options)) (*ListContactEvaluationsOutput, error) 
 ListContactFlowModules(ctx context.Context, params *ListContactFlowModulesInput, optFns ...func(*Options)) (*ListContactFlowModulesOutput, error) 
 ListContactFlows(ctx context.Context, params *ListContactFlowsInput, optFns ...func(*Options)) (*ListContactFlowsOutput, error) 
 ListContactReferences(ctx context.Context, params *ListContactReferencesInput, optFns ...func(*Options)) (*ListContactReferencesOutput, error) 
 ListDefaultVocabularies(ctx context.Context, params *ListDefaultVocabulariesInput, optFns ...func(*Options)) (*ListDefaultVocabulariesOutput, error) 
 ListEvaluationFormVersions(ctx context.Context, params *ListEvaluationFormVersionsInput, optFns ...func(*Options)) (*ListEvaluationFormVersionsOutput, error) 
 ListEvaluationForms(ctx context.Context, params *ListEvaluationFormsInput, optFns ...func(*Options)) (*ListEvaluationFormsOutput, error) 
 ListFlowAssociations(ctx context.Context, params *ListFlowAssociationsInput, optFns ...func(*Options)) (*ListFlowAssociationsOutput, error) 
 ListHoursOfOperations(ctx context.Context, params *ListHoursOfOperationsInput, optFns ...func(*Options)) (*ListHoursOfOperationsOutput, error) 
 ListInstanceAttributes(ctx context.Context, params *ListInstanceAttributesInput, optFns ...func(*Options)) (*ListInstanceAttributesOutput, error) 
 ListInstanceStorageConfigs(ctx context.Context, params *ListInstanceStorageConfigsInput, optFns ...func(*Options)) (*ListInstanceStorageConfigsOutput, error) 
 ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) 
 ListIntegrationAssociations(ctx context.Context, params *ListIntegrationAssociationsInput, optFns ...func(*Options)) (*ListIntegrationAssociationsOutput, error) 
 ListLambdaFunctions(ctx context.Context, params *ListLambdaFunctionsInput, optFns ...func(*Options)) (*ListLambdaFunctionsOutput, error) 
 ListLexBots(ctx context.Context, params *ListLexBotsInput, optFns ...func(*Options)) (*ListLexBotsOutput, error) 
 ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) 
 ListPhoneNumbersV2(ctx context.Context, params *ListPhoneNumbersV2Input, optFns ...func(*Options)) (*ListPhoneNumbersV2Output, error) 
 ListPredefinedAttributes(ctx context.Context, params *ListPredefinedAttributesInput, optFns ...func(*Options)) (*ListPredefinedAttributesOutput, error) 
 ListPrompts(ctx context.Context, params *ListPromptsInput, optFns ...func(*Options)) (*ListPromptsOutput, error) 
 ListQueueQuickConnects(ctx context.Context, params *ListQueueQuickConnectsInput, optFns ...func(*Options)) (*ListQueueQuickConnectsOutput, error) 
 ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) 
 ListQuickConnects(ctx context.Context, params *ListQuickConnectsInput, optFns ...func(*Options)) (*ListQuickConnectsOutput, error) 
 ListRealtimeContactAnalysisSegmentsV2(ctx context.Context, params *ListRealtimeContactAnalysisSegmentsV2Input, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsV2Output, error) 
 ListRoutingProfileQueues(ctx context.Context, params *ListRoutingProfileQueuesInput, optFns ...func(*Options)) (*ListRoutingProfileQueuesOutput, error) 
 ListRoutingProfiles(ctx context.Context, params *ListRoutingProfilesInput, optFns ...func(*Options)) (*ListRoutingProfilesOutput, error) 
 ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) 
 ListSecurityKeys(ctx context.Context, params *ListSecurityKeysInput, optFns ...func(*Options)) (*ListSecurityKeysOutput, error) 
 ListSecurityProfileApplications(ctx context.Context, params *ListSecurityProfileApplicationsInput, optFns ...func(*Options)) (*ListSecurityProfileApplicationsOutput, error) 
 ListSecurityProfilePermissions(ctx context.Context, params *ListSecurityProfilePermissionsInput, optFns ...func(*Options)) (*ListSecurityProfilePermissionsOutput, error) 
 ListSecurityProfiles(ctx context.Context, params *ListSecurityProfilesInput, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTaskTemplates(ctx context.Context, params *ListTaskTemplatesInput, optFns ...func(*Options)) (*ListTaskTemplatesOutput, error) 
 ListTrafficDistributionGroupUsers(ctx context.Context, params *ListTrafficDistributionGroupUsersInput, optFns ...func(*Options)) (*ListTrafficDistributionGroupUsersOutput, error) 
 ListTrafficDistributionGroups(ctx context.Context, params *ListTrafficDistributionGroupsInput, optFns ...func(*Options)) (*ListTrafficDistributionGroupsOutput, error) 
 ListUseCases(ctx context.Context, params *ListUseCasesInput, optFns ...func(*Options)) (*ListUseCasesOutput, error) 
 ListUserHierarchyGroups(ctx context.Context, params *ListUserHierarchyGroupsInput, optFns ...func(*Options)) (*ListUserHierarchyGroupsOutput, error) 
 ListUserProficiencies(ctx context.Context, params *ListUserProficienciesInput, optFns ...func(*Options)) (*ListUserProficienciesOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListViewVersions(ctx context.Context, params *ListViewVersionsInput, optFns ...func(*Options)) (*ListViewVersionsOutput, error) 
 ListViews(ctx context.Context, params *ListViewsInput, optFns ...func(*Options)) (*ListViewsOutput, error) 
 MonitorContact(ctx context.Context, params *MonitorContactInput, optFns ...func(*Options)) (*MonitorContactOutput, error) 
 PauseContact(ctx context.Context, params *PauseContactInput, optFns ...func(*Options)) (*PauseContactOutput, error) 
 PutUserStatus(ctx context.Context, params *PutUserStatusInput, optFns ...func(*Options)) (*PutUserStatusOutput, error) 
 ReleasePhoneNumber(ctx context.Context, params *ReleasePhoneNumberInput, optFns ...func(*Options)) (*ReleasePhoneNumberOutput, error) 
 ReplicateInstance(ctx context.Context, params *ReplicateInstanceInput, optFns ...func(*Options)) (*ReplicateInstanceOutput, error) 
 ResumeContact(ctx context.Context, params *ResumeContactInput, optFns ...func(*Options)) (*ResumeContactOutput, error) 
 ResumeContactRecording(ctx context.Context, params *ResumeContactRecordingInput, optFns ...func(*Options)) (*ResumeContactRecordingOutput, error) 
 SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) 
 SearchContactFlowModules(ctx context.Context, params *SearchContactFlowModulesInput, optFns ...func(*Options)) (*SearchContactFlowModulesOutput, error) 
 SearchContactFlows(ctx context.Context, params *SearchContactFlowsInput, optFns ...func(*Options)) (*SearchContactFlowsOutput, error) 
 SearchContacts(ctx context.Context, params *SearchContactsInput, optFns ...func(*Options)) (*SearchContactsOutput, error) 
 SearchHoursOfOperations(ctx context.Context, params *SearchHoursOfOperationsInput, optFns ...func(*Options)) (*SearchHoursOfOperationsOutput, error) 
 SearchPredefinedAttributes(ctx context.Context, params *SearchPredefinedAttributesInput, optFns ...func(*Options)) (*SearchPredefinedAttributesOutput, error) 
 SearchPrompts(ctx context.Context, params *SearchPromptsInput, optFns ...func(*Options)) (*SearchPromptsOutput, error) 
 SearchQueues(ctx context.Context, params *SearchQueuesInput, optFns ...func(*Options)) (*SearchQueuesOutput, error) 
 SearchQuickConnects(ctx context.Context, params *SearchQuickConnectsInput, optFns ...func(*Options)) (*SearchQuickConnectsOutput, error) 
 SearchResourceTags(ctx context.Context, params *SearchResourceTagsInput, optFns ...func(*Options)) (*SearchResourceTagsOutput, error) 
 SearchRoutingProfiles(ctx context.Context, params *SearchRoutingProfilesInput, optFns ...func(*Options)) (*SearchRoutingProfilesOutput, error) 
 SearchSecurityProfiles(ctx context.Context, params *SearchSecurityProfilesInput, optFns ...func(*Options)) (*SearchSecurityProfilesOutput, error) 
 SearchUsers(ctx context.Context, params *SearchUsersInput, optFns ...func(*Options)) (*SearchUsersOutput, error) 
 SearchVocabularies(ctx context.Context, params *SearchVocabulariesInput, optFns ...func(*Options)) (*SearchVocabulariesOutput, error) 
 SendChatIntegrationEvent(ctx context.Context, params *SendChatIntegrationEventInput, optFns ...func(*Options)) (*SendChatIntegrationEventOutput, error) 
 StartAttachedFileUpload(ctx context.Context, params *StartAttachedFileUploadInput, optFns ...func(*Options)) (*StartAttachedFileUploadOutput, error) 
 StartChatContact(ctx context.Context, params *StartChatContactInput, optFns ...func(*Options)) (*StartChatContactOutput, error) 
 StartContactEvaluation(ctx context.Context, params *StartContactEvaluationInput, optFns ...func(*Options)) (*StartContactEvaluationOutput, error) 
 StartContactRecording(ctx context.Context, params *StartContactRecordingInput, optFns ...func(*Options)) (*StartContactRecordingOutput, error) 
 StartContactStreaming(ctx context.Context, params *StartContactStreamingInput, optFns ...func(*Options)) (*StartContactStreamingOutput, error) 
 StartOutboundVoiceContact(ctx context.Context, params *StartOutboundVoiceContactInput, optFns ...func(*Options)) (*StartOutboundVoiceContactOutput, error) 
 StartTaskContact(ctx context.Context, params *StartTaskContactInput, optFns ...func(*Options)) (*StartTaskContactOutput, error) 
 StartWebRTCContact(ctx context.Context, params *StartWebRTCContactInput, optFns ...func(*Options)) (*StartWebRTCContactOutput, error) 
 StopContact(ctx context.Context, params *StopContactInput, optFns ...func(*Options)) (*StopContactOutput, error) 
 StopContactRecording(ctx context.Context, params *StopContactRecordingInput, optFns ...func(*Options)) (*StopContactRecordingOutput, error) 
 StopContactStreaming(ctx context.Context, params *StopContactStreamingInput, optFns ...func(*Options)) (*StopContactStreamingOutput, error) 
 SubmitContactEvaluation(ctx context.Context, params *SubmitContactEvaluationInput, optFns ...func(*Options)) (*SubmitContactEvaluationOutput, error) 
 SuspendContactRecording(ctx context.Context, params *SuspendContactRecordingInput, optFns ...func(*Options)) (*SuspendContactRecordingOutput, error) 
 TagContact(ctx context.Context, params *TagContactInput, optFns ...func(*Options)) (*TagContactOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TransferContact(ctx context.Context, params *TransferContactInput, optFns ...func(*Options)) (*TransferContactOutput, error) 
 UntagContact(ctx context.Context, params *UntagContactInput, optFns ...func(*Options)) (*UntagContactOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAgentStatus(ctx context.Context, params *UpdateAgentStatusInput, optFns ...func(*Options)) (*UpdateAgentStatusOutput, error) 
 UpdateAuthenticationProfile(ctx context.Context, params *UpdateAuthenticationProfileInput, optFns ...func(*Options)) (*UpdateAuthenticationProfileOutput, error) 
 UpdateContact(ctx context.Context, params *UpdateContactInput, optFns ...func(*Options)) (*UpdateContactOutput, error) 
 UpdateContactAttributes(ctx context.Context, params *UpdateContactAttributesInput, optFns ...func(*Options)) (*UpdateContactAttributesOutput, error) 
 UpdateContactEvaluation(ctx context.Context, params *UpdateContactEvaluationInput, optFns ...func(*Options)) (*UpdateContactEvaluationOutput, error) 
 UpdateContactFlowContent(ctx context.Context, params *UpdateContactFlowContentInput, optFns ...func(*Options)) (*UpdateContactFlowContentOutput, error) 
 UpdateContactFlowMetadata(ctx context.Context, params *UpdateContactFlowMetadataInput, optFns ...func(*Options)) (*UpdateContactFlowMetadataOutput, error) 
 UpdateContactFlowModuleContent(ctx context.Context, params *UpdateContactFlowModuleContentInput, optFns ...func(*Options)) (*UpdateContactFlowModuleContentOutput, error) 
 UpdateContactFlowModuleMetadata(ctx context.Context, params *UpdateContactFlowModuleMetadataInput, optFns ...func(*Options)) (*UpdateContactFlowModuleMetadataOutput, error) 
 UpdateContactFlowName(ctx context.Context, params *UpdateContactFlowNameInput, optFns ...func(*Options)) (*UpdateContactFlowNameOutput, error) 
 UpdateContactRoutingData(ctx context.Context, params *UpdateContactRoutingDataInput, optFns ...func(*Options)) (*UpdateContactRoutingDataOutput, error) 
 UpdateContactSchedule(ctx context.Context, params *UpdateContactScheduleInput, optFns ...func(*Options)) (*UpdateContactScheduleOutput, error) 
 UpdateEvaluationForm(ctx context.Context, params *UpdateEvaluationFormInput, optFns ...func(*Options)) (*UpdateEvaluationFormOutput, error) 
 UpdateHoursOfOperation(ctx context.Context, params *UpdateHoursOfOperationInput, optFns ...func(*Options)) (*UpdateHoursOfOperationOutput, error) 
 UpdateInstanceAttribute(ctx context.Context, params *UpdateInstanceAttributeInput, optFns ...func(*Options)) (*UpdateInstanceAttributeOutput, error) 
 UpdateInstanceStorageConfig(ctx context.Context, params *UpdateInstanceStorageConfigInput, optFns ...func(*Options)) (*UpdateInstanceStorageConfigOutput, error) 
 UpdateParticipantRoleConfig(ctx context.Context, params *UpdateParticipantRoleConfigInput, optFns ...func(*Options)) (*UpdateParticipantRoleConfigOutput, error) 
 UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) 
 UpdatePhoneNumberMetadata(ctx context.Context, params *UpdatePhoneNumberMetadataInput, optFns ...func(*Options)) (*UpdatePhoneNumberMetadataOutput, error) 
 UpdatePredefinedAttribute(ctx context.Context, params *UpdatePredefinedAttributeInput, optFns ...func(*Options)) (*UpdatePredefinedAttributeOutput, error) 
 UpdatePrompt(ctx context.Context, params *UpdatePromptInput, optFns ...func(*Options)) (*UpdatePromptOutput, error) 
 UpdateQueueHoursOfOperation(ctx context.Context, params *UpdateQueueHoursOfOperationInput, optFns ...func(*Options)) (*UpdateQueueHoursOfOperationOutput, error) 
 UpdateQueueMaxContacts(ctx context.Context, params *UpdateQueueMaxContactsInput, optFns ...func(*Options)) (*UpdateQueueMaxContactsOutput, error) 
 UpdateQueueName(ctx context.Context, params *UpdateQueueNameInput, optFns ...func(*Options)) (*UpdateQueueNameOutput, error) 
 UpdateQueueOutboundCallerConfig(ctx context.Context, params *UpdateQueueOutboundCallerConfigInput, optFns ...func(*Options)) (*UpdateQueueOutboundCallerConfigOutput, error) 
 UpdateQueueStatus(ctx context.Context, params *UpdateQueueStatusInput, optFns ...func(*Options)) (*UpdateQueueStatusOutput, error) 
 UpdateQuickConnectConfig(ctx context.Context, params *UpdateQuickConnectConfigInput, optFns ...func(*Options)) (*UpdateQuickConnectConfigOutput, error) 
 UpdateQuickConnectName(ctx context.Context, params *UpdateQuickConnectNameInput, optFns ...func(*Options)) (*UpdateQuickConnectNameOutput, error) 
 UpdateRoutingProfileAgentAvailabilityTimer(ctx context.Context, params *UpdateRoutingProfileAgentAvailabilityTimerInput, optFns ...func(*Options)) (*UpdateRoutingProfileAgentAvailabilityTimerOutput, error) 
 UpdateRoutingProfileConcurrency(ctx context.Context, params *UpdateRoutingProfileConcurrencyInput, optFns ...func(*Options)) (*UpdateRoutingProfileConcurrencyOutput, error) 
 UpdateRoutingProfileDefaultOutboundQueue(ctx context.Context, params *UpdateRoutingProfileDefaultOutboundQueueInput, optFns ...func(*Options)) (*UpdateRoutingProfileDefaultOutboundQueueOutput, error) 
 UpdateRoutingProfileName(ctx context.Context, params *UpdateRoutingProfileNameInput, optFns ...func(*Options)) (*UpdateRoutingProfileNameOutput, error) 
 UpdateRoutingProfileQueues(ctx context.Context, params *UpdateRoutingProfileQueuesInput, optFns ...func(*Options)) (*UpdateRoutingProfileQueuesOutput, error) 
 UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) 
 UpdateSecurityProfile(ctx context.Context, params *UpdateSecurityProfileInput, optFns ...func(*Options)) (*UpdateSecurityProfileOutput, error) 
 UpdateTaskTemplate(ctx context.Context, params *UpdateTaskTemplateInput, optFns ...func(*Options)) (*UpdateTaskTemplateOutput, error) 
 UpdateTrafficDistribution(ctx context.Context, params *UpdateTrafficDistributionInput, optFns ...func(*Options)) (*UpdateTrafficDistributionOutput, error) 
 UpdateUserHierarchy(ctx context.Context, params *UpdateUserHierarchyInput, optFns ...func(*Options)) (*UpdateUserHierarchyOutput, error) 
 UpdateUserHierarchyGroupName(ctx context.Context, params *UpdateUserHierarchyGroupNameInput, optFns ...func(*Options)) (*UpdateUserHierarchyGroupNameOutput, error) 
 UpdateUserHierarchyStructure(ctx context.Context, params *UpdateUserHierarchyStructureInput, optFns ...func(*Options)) (*UpdateUserHierarchyStructureOutput, error) 
 UpdateUserIdentityInfo(ctx context.Context, params *UpdateUserIdentityInfoInput, optFns ...func(*Options)) (*UpdateUserIdentityInfoOutput, error) 
 UpdateUserPhoneConfig(ctx context.Context, params *UpdateUserPhoneConfigInput, optFns ...func(*Options)) (*UpdateUserPhoneConfigOutput, error) 
 UpdateUserProficiencies(ctx context.Context, params *UpdateUserProficienciesInput, optFns ...func(*Options)) (*UpdateUserProficienciesOutput, error) 
 UpdateUserRoutingProfile(ctx context.Context, params *UpdateUserRoutingProfileInput, optFns ...func(*Options)) (*UpdateUserRoutingProfileOutput, error) 
 UpdateUserSecurityProfiles(ctx context.Context, params *UpdateUserSecurityProfilesInput, optFns ...func(*Options)) (*UpdateUserSecurityProfilesOutput, error) 
 UpdateViewContent(ctx context.Context, params *UpdateViewContentInput, optFns ...func(*Options)) (*UpdateViewContentOutput, error) 
 UpdateViewMetadata(ctx context.Context, params *UpdateViewMetadataInput, optFns ...func(*Options)) (*UpdateViewMetadataOutput, error) 
}
