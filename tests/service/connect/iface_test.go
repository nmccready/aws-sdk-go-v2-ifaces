package connect_test

// tests for the connect service interface for this ../../../service/connect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connect/connect_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectServiceCanBeMocked(t *testing.T) {
	var iface connect_iface.IClient
	iface = &connect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateEvaluationForm", func(t *testing.T) {
        input := &connect.ActivateEvaluationFormInput{}
        output := &connect.ActivateEvaluationFormOutput{}

        mockClient.On("ActivateEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAnalyticsDataSet", func(t *testing.T) {
        input := &connect.AssociateAnalyticsDataSetInput{}
        output := &connect.AssociateAnalyticsDataSetOutput{}

        mockClient.On("AssociateAnalyticsDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAnalyticsDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApprovedOrigin", func(t *testing.T) {
        input := &connect.AssociateApprovedOriginInput{}
        output := &connect.AssociateApprovedOriginOutput{}

        mockClient.On("AssociateApprovedOrigin", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApprovedOrigin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateBot", func(t *testing.T) {
        input := &connect.AssociateBotInput{}
        output := &connect.AssociateBotOutput{}

        mockClient.On("AssociateBot", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDefaultVocabulary", func(t *testing.T) {
        input := &connect.AssociateDefaultVocabularyInput{}
        output := &connect.AssociateDefaultVocabularyOutput{}

        mockClient.On("AssociateDefaultVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDefaultVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFlow", func(t *testing.T) {
        input := &connect.AssociateFlowInput{}
        output := &connect.AssociateFlowOutput{}

        mockClient.On("AssociateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateInstanceStorageConfig", func(t *testing.T) {
        input := &connect.AssociateInstanceStorageConfigInput{}
        output := &connect.AssociateInstanceStorageConfigOutput{}

        mockClient.On("AssociateInstanceStorageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateInstanceStorageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLambdaFunction", func(t *testing.T) {
        input := &connect.AssociateLambdaFunctionInput{}
        output := &connect.AssociateLambdaFunctionOutput{}

        mockClient.On("AssociateLambdaFunction", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLambdaFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLexBot", func(t *testing.T) {
        input := &connect.AssociateLexBotInput{}
        output := &connect.AssociateLexBotOutput{}

        mockClient.On("AssociateLexBot", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLexBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumberContactFlow", func(t *testing.T) {
        input := &connect.AssociatePhoneNumberContactFlowInput{}
        output := &connect.AssociatePhoneNumberContactFlowOutput{}

        mockClient.On("AssociatePhoneNumberContactFlow", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumberContactFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateQueueQuickConnects", func(t *testing.T) {
        input := &connect.AssociateQueueQuickConnectsInput{}
        output := &connect.AssociateQueueQuickConnectsOutput{}

        mockClient.On("AssociateQueueQuickConnects", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateQueueQuickConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateRoutingProfileQueues", func(t *testing.T) {
        input := &connect.AssociateRoutingProfileQueuesInput{}
        output := &connect.AssociateRoutingProfileQueuesOutput{}

        mockClient.On("AssociateRoutingProfileQueues", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateRoutingProfileQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSecurityKey", func(t *testing.T) {
        input := &connect.AssociateSecurityKeyInput{}
        output := &connect.AssociateSecurityKeyOutput{}

        mockClient.On("AssociateSecurityKey", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSecurityKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTrafficDistributionGroupUser", func(t *testing.T) {
        input := &connect.AssociateTrafficDistributionGroupUserInput{}
        output := &connect.AssociateTrafficDistributionGroupUserOutput{}

        mockClient.On("AssociateTrafficDistributionGroupUser", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTrafficDistributionGroupUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateUserProficiencies", func(t *testing.T) {
        input := &connect.AssociateUserProficienciesInput{}
        output := &connect.AssociateUserProficienciesOutput{}

        mockClient.On("AssociateUserProficiencies", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateUserProficiencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateAnalyticsDataSet", func(t *testing.T) {
        input := &connect.BatchAssociateAnalyticsDataSetInput{}
        output := &connect.BatchAssociateAnalyticsDataSetOutput{}

        mockClient.On("BatchAssociateAnalyticsDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateAnalyticsDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateAnalyticsDataSet", func(t *testing.T) {
        input := &connect.BatchDisassociateAnalyticsDataSetInput{}
        output := &connect.BatchDisassociateAnalyticsDataSetOutput{}

        mockClient.On("BatchDisassociateAnalyticsDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateAnalyticsDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAttachedFileMetadata", func(t *testing.T) {
        input := &connect.BatchGetAttachedFileMetadataInput{}
        output := &connect.BatchGetAttachedFileMetadataOutput{}

        mockClient.On("BatchGetAttachedFileMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAttachedFileMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFlowAssociation", func(t *testing.T) {
        input := &connect.BatchGetFlowAssociationInput{}
        output := &connect.BatchGetFlowAssociationOutput{}

        mockClient.On("BatchGetFlowAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFlowAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutContact", func(t *testing.T) {
        input := &connect.BatchPutContactInput{}
        output := &connect.BatchPutContactOutput{}

        mockClient.On("BatchPutContact", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClaimPhoneNumber", func(t *testing.T) {
        input := &connect.ClaimPhoneNumberInput{}
        output := &connect.ClaimPhoneNumberOutput{}

        mockClient.On("ClaimPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.ClaimPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteAttachedFileUpload", func(t *testing.T) {
        input := &connect.CompleteAttachedFileUploadInput{}
        output := &connect.CompleteAttachedFileUploadOutput{}

        mockClient.On("CompleteAttachedFileUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteAttachedFileUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgentStatus", func(t *testing.T) {
        input := &connect.CreateAgentStatusInput{}
        output := &connect.CreateAgentStatusOutput{}

        mockClient.On("CreateAgentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContactFlow", func(t *testing.T) {
        input := &connect.CreateContactFlowInput{}
        output := &connect.CreateContactFlowOutput{}

        mockClient.On("CreateContactFlow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContactFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContactFlowModule", func(t *testing.T) {
        input := &connect.CreateContactFlowModuleInput{}
        output := &connect.CreateContactFlowModuleOutput{}

        mockClient.On("CreateContactFlowModule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContactFlowModule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEvaluationForm", func(t *testing.T) {
        input := &connect.CreateEvaluationFormInput{}
        output := &connect.CreateEvaluationFormOutput{}

        mockClient.On("CreateEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHoursOfOperation", func(t *testing.T) {
        input := &connect.CreateHoursOfOperationInput{}
        output := &connect.CreateHoursOfOperationOutput{}

        mockClient.On("CreateHoursOfOperation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHoursOfOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstance", func(t *testing.T) {
        input := &connect.CreateInstanceInput{}
        output := &connect.CreateInstanceOutput{}

        mockClient.On("CreateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntegrationAssociation", func(t *testing.T) {
        input := &connect.CreateIntegrationAssociationInput{}
        output := &connect.CreateIntegrationAssociationOutput{}

        mockClient.On("CreateIntegrationAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntegrationAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateParticipant", func(t *testing.T) {
        input := &connect.CreateParticipantInput{}
        output := &connect.CreateParticipantOutput{}

        mockClient.On("CreateParticipant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateParticipant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePersistentContactAssociation", func(t *testing.T) {
        input := &connect.CreatePersistentContactAssociationInput{}
        output := &connect.CreatePersistentContactAssociationOutput{}

        mockClient.On("CreatePersistentContactAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePersistentContactAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePredefinedAttribute", func(t *testing.T) {
        input := &connect.CreatePredefinedAttributeInput{}
        output := &connect.CreatePredefinedAttributeOutput{}

        mockClient.On("CreatePredefinedAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePredefinedAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrompt", func(t *testing.T) {
        input := &connect.CreatePromptInput{}
        output := &connect.CreatePromptOutput{}

        mockClient.On("CreatePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueue", func(t *testing.T) {
        input := &connect.CreateQueueInput{}
        output := &connect.CreateQueueOutput{}

        mockClient.On("CreateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQuickConnect", func(t *testing.T) {
        input := &connect.CreateQuickConnectInput{}
        output := &connect.CreateQuickConnectOutput{}

        mockClient.On("CreateQuickConnect", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQuickConnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoutingProfile", func(t *testing.T) {
        input := &connect.CreateRoutingProfileInput{}
        output := &connect.CreateRoutingProfileOutput{}

        mockClient.On("CreateRoutingProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoutingProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &connect.CreateRuleInput{}
        output := &connect.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityProfile", func(t *testing.T) {
        input := &connect.CreateSecurityProfileInput{}
        output := &connect.CreateSecurityProfileOutput{}

        mockClient.On("CreateSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTaskTemplate", func(t *testing.T) {
        input := &connect.CreateTaskTemplateInput{}
        output := &connect.CreateTaskTemplateOutput{}

        mockClient.On("CreateTaskTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTaskTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficDistributionGroup", func(t *testing.T) {
        input := &connect.CreateTrafficDistributionGroupInput{}
        output := &connect.CreateTrafficDistributionGroupOutput{}

        mockClient.On("CreateTrafficDistributionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficDistributionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUseCase", func(t *testing.T) {
        input := &connect.CreateUseCaseInput{}
        output := &connect.CreateUseCaseOutput{}

        mockClient.On("CreateUseCase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUseCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &connect.CreateUserInput{}
        output := &connect.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserHierarchyGroup", func(t *testing.T) {
        input := &connect.CreateUserHierarchyGroupInput{}
        output := &connect.CreateUserHierarchyGroupOutput{}

        mockClient.On("CreateUserHierarchyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserHierarchyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateView", func(t *testing.T) {
        input := &connect.CreateViewInput{}
        output := &connect.CreateViewOutput{}

        mockClient.On("CreateView", ctx, input).Return(output, nil)

        result, err := mockClient.CreateView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateViewVersion", func(t *testing.T) {
        input := &connect.CreateViewVersionInput{}
        output := &connect.CreateViewVersionOutput{}

        mockClient.On("CreateViewVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateViewVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVocabulary", func(t *testing.T) {
        input := &connect.CreateVocabularyInput{}
        output := &connect.CreateVocabularyOutput{}

        mockClient.On("CreateVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateEvaluationForm", func(t *testing.T) {
        input := &connect.DeactivateEvaluationFormInput{}
        output := &connect.DeactivateEvaluationFormOutput{}

        mockClient.On("DeactivateEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttachedFile", func(t *testing.T) {
        input := &connect.DeleteAttachedFileInput{}
        output := &connect.DeleteAttachedFileOutput{}

        mockClient.On("DeleteAttachedFile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttachedFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactEvaluation", func(t *testing.T) {
        input := &connect.DeleteContactEvaluationInput{}
        output := &connect.DeleteContactEvaluationOutput{}

        mockClient.On("DeleteContactEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactFlow", func(t *testing.T) {
        input := &connect.DeleteContactFlowInput{}
        output := &connect.DeleteContactFlowOutput{}

        mockClient.On("DeleteContactFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactFlowModule", func(t *testing.T) {
        input := &connect.DeleteContactFlowModuleInput{}
        output := &connect.DeleteContactFlowModuleOutput{}

        mockClient.On("DeleteContactFlowModule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactFlowModule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEvaluationForm", func(t *testing.T) {
        input := &connect.DeleteEvaluationFormInput{}
        output := &connect.DeleteEvaluationFormOutput{}

        mockClient.On("DeleteEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHoursOfOperation", func(t *testing.T) {
        input := &connect.DeleteHoursOfOperationInput{}
        output := &connect.DeleteHoursOfOperationOutput{}

        mockClient.On("DeleteHoursOfOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHoursOfOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstance", func(t *testing.T) {
        input := &connect.DeleteInstanceInput{}
        output := &connect.DeleteInstanceOutput{}

        mockClient.On("DeleteInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegrationAssociation", func(t *testing.T) {
        input := &connect.DeleteIntegrationAssociationInput{}
        output := &connect.DeleteIntegrationAssociationOutput{}

        mockClient.On("DeleteIntegrationAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegrationAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePredefinedAttribute", func(t *testing.T) {
        input := &connect.DeletePredefinedAttributeInput{}
        output := &connect.DeletePredefinedAttributeOutput{}

        mockClient.On("DeletePredefinedAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePredefinedAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrompt", func(t *testing.T) {
        input := &connect.DeletePromptInput{}
        output := &connect.DeletePromptOutput{}

        mockClient.On("DeletePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueue", func(t *testing.T) {
        input := &connect.DeleteQueueInput{}
        output := &connect.DeleteQueueOutput{}

        mockClient.On("DeleteQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQuickConnect", func(t *testing.T) {
        input := &connect.DeleteQuickConnectInput{}
        output := &connect.DeleteQuickConnectOutput{}

        mockClient.On("DeleteQuickConnect", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQuickConnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoutingProfile", func(t *testing.T) {
        input := &connect.DeleteRoutingProfileInput{}
        output := &connect.DeleteRoutingProfileOutput{}

        mockClient.On("DeleteRoutingProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoutingProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &connect.DeleteRuleInput{}
        output := &connect.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityProfile", func(t *testing.T) {
        input := &connect.DeleteSecurityProfileInput{}
        output := &connect.DeleteSecurityProfileOutput{}

        mockClient.On("DeleteSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTaskTemplate", func(t *testing.T) {
        input := &connect.DeleteTaskTemplateInput{}
        output := &connect.DeleteTaskTemplateOutput{}

        mockClient.On("DeleteTaskTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTaskTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficDistributionGroup", func(t *testing.T) {
        input := &connect.DeleteTrafficDistributionGroupInput{}
        output := &connect.DeleteTrafficDistributionGroupOutput{}

        mockClient.On("DeleteTrafficDistributionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficDistributionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUseCase", func(t *testing.T) {
        input := &connect.DeleteUseCaseInput{}
        output := &connect.DeleteUseCaseOutput{}

        mockClient.On("DeleteUseCase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUseCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &connect.DeleteUserInput{}
        output := &connect.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserHierarchyGroup", func(t *testing.T) {
        input := &connect.DeleteUserHierarchyGroupInput{}
        output := &connect.DeleteUserHierarchyGroupOutput{}

        mockClient.On("DeleteUserHierarchyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserHierarchyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteView", func(t *testing.T) {
        input := &connect.DeleteViewInput{}
        output := &connect.DeleteViewOutput{}

        mockClient.On("DeleteView", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteViewVersion", func(t *testing.T) {
        input := &connect.DeleteViewVersionInput{}
        output := &connect.DeleteViewVersionOutput{}

        mockClient.On("DeleteViewVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteViewVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVocabulary", func(t *testing.T) {
        input := &connect.DeleteVocabularyInput{}
        output := &connect.DeleteVocabularyOutput{}

        mockClient.On("DeleteVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgentStatus", func(t *testing.T) {
        input := &connect.DescribeAgentStatusInput{}
        output := &connect.DescribeAgentStatusOutput{}

        mockClient.On("DescribeAgentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuthenticationProfile", func(t *testing.T) {
        input := &connect.DescribeAuthenticationProfileInput{}
        output := &connect.DescribeAuthenticationProfileOutput{}

        mockClient.On("DescribeAuthenticationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuthenticationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContact", func(t *testing.T) {
        input := &connect.DescribeContactInput{}
        output := &connect.DescribeContactOutput{}

        mockClient.On("DescribeContact", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContactEvaluation", func(t *testing.T) {
        input := &connect.DescribeContactEvaluationInput{}
        output := &connect.DescribeContactEvaluationOutput{}

        mockClient.On("DescribeContactEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContactEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContactFlow", func(t *testing.T) {
        input := &connect.DescribeContactFlowInput{}
        output := &connect.DescribeContactFlowOutput{}

        mockClient.On("DescribeContactFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContactFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContactFlowModule", func(t *testing.T) {
        input := &connect.DescribeContactFlowModuleInput{}
        output := &connect.DescribeContactFlowModuleOutput{}

        mockClient.On("DescribeContactFlowModule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContactFlowModule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvaluationForm", func(t *testing.T) {
        input := &connect.DescribeEvaluationFormInput{}
        output := &connect.DescribeEvaluationFormOutput{}

        mockClient.On("DescribeEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHoursOfOperation", func(t *testing.T) {
        input := &connect.DescribeHoursOfOperationInput{}
        output := &connect.DescribeHoursOfOperationOutput{}

        mockClient.On("DescribeHoursOfOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHoursOfOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstance", func(t *testing.T) {
        input := &connect.DescribeInstanceInput{}
        output := &connect.DescribeInstanceOutput{}

        mockClient.On("DescribeInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceAttribute", func(t *testing.T) {
        input := &connect.DescribeInstanceAttributeInput{}
        output := &connect.DescribeInstanceAttributeOutput{}

        mockClient.On("DescribeInstanceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceStorageConfig", func(t *testing.T) {
        input := &connect.DescribeInstanceStorageConfigInput{}
        output := &connect.DescribeInstanceStorageConfigOutput{}

        mockClient.On("DescribeInstanceStorageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceStorageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePhoneNumber", func(t *testing.T) {
        input := &connect.DescribePhoneNumberInput{}
        output := &connect.DescribePhoneNumberOutput{}

        mockClient.On("DescribePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePredefinedAttribute", func(t *testing.T) {
        input := &connect.DescribePredefinedAttributeInput{}
        output := &connect.DescribePredefinedAttributeOutput{}

        mockClient.On("DescribePredefinedAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePredefinedAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePrompt", func(t *testing.T) {
        input := &connect.DescribePromptInput{}
        output := &connect.DescribePromptOutput{}

        mockClient.On("DescribePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQueue", func(t *testing.T) {
        input := &connect.DescribeQueueInput{}
        output := &connect.DescribeQueueOutput{}

        mockClient.On("DescribeQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQuickConnect", func(t *testing.T) {
        input := &connect.DescribeQuickConnectInput{}
        output := &connect.DescribeQuickConnectOutput{}

        mockClient.On("DescribeQuickConnect", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQuickConnect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRoutingProfile", func(t *testing.T) {
        input := &connect.DescribeRoutingProfileInput{}
        output := &connect.DescribeRoutingProfileOutput{}

        mockClient.On("DescribeRoutingProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRoutingProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRule", func(t *testing.T) {
        input := &connect.DescribeRuleInput{}
        output := &connect.DescribeRuleOutput{}

        mockClient.On("DescribeRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityProfile", func(t *testing.T) {
        input := &connect.DescribeSecurityProfileInput{}
        output := &connect.DescribeSecurityProfileOutput{}

        mockClient.On("DescribeSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrafficDistributionGroup", func(t *testing.T) {
        input := &connect.DescribeTrafficDistributionGroupInput{}
        output := &connect.DescribeTrafficDistributionGroupOutput{}

        mockClient.On("DescribeTrafficDistributionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrafficDistributionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &connect.DescribeUserInput{}
        output := &connect.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserHierarchyGroup", func(t *testing.T) {
        input := &connect.DescribeUserHierarchyGroupInput{}
        output := &connect.DescribeUserHierarchyGroupOutput{}

        mockClient.On("DescribeUserHierarchyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserHierarchyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserHierarchyStructure", func(t *testing.T) {
        input := &connect.DescribeUserHierarchyStructureInput{}
        output := &connect.DescribeUserHierarchyStructureOutput{}

        mockClient.On("DescribeUserHierarchyStructure", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserHierarchyStructure(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeView", func(t *testing.T) {
        input := &connect.DescribeViewInput{}
        output := &connect.DescribeViewOutput{}

        mockClient.On("DescribeView", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVocabulary", func(t *testing.T) {
        input := &connect.DescribeVocabularyInput{}
        output := &connect.DescribeVocabularyOutput{}

        mockClient.On("DescribeVocabulary", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVocabulary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAnalyticsDataSet", func(t *testing.T) {
        input := &connect.DisassociateAnalyticsDataSetInput{}
        output := &connect.DisassociateAnalyticsDataSetOutput{}

        mockClient.On("DisassociateAnalyticsDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAnalyticsDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApprovedOrigin", func(t *testing.T) {
        input := &connect.DisassociateApprovedOriginInput{}
        output := &connect.DisassociateApprovedOriginOutput{}

        mockClient.On("DisassociateApprovedOrigin", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApprovedOrigin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateBot", func(t *testing.T) {
        input := &connect.DisassociateBotInput{}
        output := &connect.DisassociateBotOutput{}

        mockClient.On("DisassociateBot", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFlow", func(t *testing.T) {
        input := &connect.DisassociateFlowInput{}
        output := &connect.DisassociateFlowOutput{}

        mockClient.On("DisassociateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateInstanceStorageConfig", func(t *testing.T) {
        input := &connect.DisassociateInstanceStorageConfigInput{}
        output := &connect.DisassociateInstanceStorageConfigOutput{}

        mockClient.On("DisassociateInstanceStorageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateInstanceStorageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLambdaFunction", func(t *testing.T) {
        input := &connect.DisassociateLambdaFunctionInput{}
        output := &connect.DisassociateLambdaFunctionOutput{}

        mockClient.On("DisassociateLambdaFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLambdaFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLexBot", func(t *testing.T) {
        input := &connect.DisassociateLexBotInput{}
        output := &connect.DisassociateLexBotOutput{}

        mockClient.On("DisassociateLexBot", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLexBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumberContactFlow", func(t *testing.T) {
        input := &connect.DisassociatePhoneNumberContactFlowInput{}
        output := &connect.DisassociatePhoneNumberContactFlowOutput{}

        mockClient.On("DisassociatePhoneNumberContactFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumberContactFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateQueueQuickConnects", func(t *testing.T) {
        input := &connect.DisassociateQueueQuickConnectsInput{}
        output := &connect.DisassociateQueueQuickConnectsOutput{}

        mockClient.On("DisassociateQueueQuickConnects", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateQueueQuickConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRoutingProfileQueues", func(t *testing.T) {
        input := &connect.DisassociateRoutingProfileQueuesInput{}
        output := &connect.DisassociateRoutingProfileQueuesOutput{}

        mockClient.On("DisassociateRoutingProfileQueues", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRoutingProfileQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSecurityKey", func(t *testing.T) {
        input := &connect.DisassociateSecurityKeyInput{}
        output := &connect.DisassociateSecurityKeyOutput{}

        mockClient.On("DisassociateSecurityKey", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSecurityKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTrafficDistributionGroupUser", func(t *testing.T) {
        input := &connect.DisassociateTrafficDistributionGroupUserInput{}
        output := &connect.DisassociateTrafficDistributionGroupUserOutput{}

        mockClient.On("DisassociateTrafficDistributionGroupUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTrafficDistributionGroupUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateUserProficiencies", func(t *testing.T) {
        input := &connect.DisassociateUserProficienciesInput{}
        output := &connect.DisassociateUserProficienciesOutput{}

        mockClient.On("DisassociateUserProficiencies", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateUserProficiencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDismissUserContact", func(t *testing.T) {
        input := &connect.DismissUserContactInput{}
        output := &connect.DismissUserContactOutput{}

        mockClient.On("DismissUserContact", ctx, input).Return(output, nil)

        result, err := mockClient.DismissUserContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttachedFile", func(t *testing.T) {
        input := &connect.GetAttachedFileInput{}
        output := &connect.GetAttachedFileOutput{}

        mockClient.On("GetAttachedFile", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttachedFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactAttributes", func(t *testing.T) {
        input := &connect.GetContactAttributesInput{}
        output := &connect.GetContactAttributesOutput{}

        mockClient.On("GetContactAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCurrentMetricData", func(t *testing.T) {
        input := &connect.GetCurrentMetricDataInput{}
        output := &connect.GetCurrentMetricDataOutput{}

        mockClient.On("GetCurrentMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetCurrentMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCurrentUserData", func(t *testing.T) {
        input := &connect.GetCurrentUserDataInput{}
        output := &connect.GetCurrentUserDataOutput{}

        mockClient.On("GetCurrentUserData", ctx, input).Return(output, nil)

        result, err := mockClient.GetCurrentUserData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFederationToken", func(t *testing.T) {
        input := &connect.GetFederationTokenInput{}
        output := &connect.GetFederationTokenOutput{}

        mockClient.On("GetFederationToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetFederationToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFlowAssociation", func(t *testing.T) {
        input := &connect.GetFlowAssociationInput{}
        output := &connect.GetFlowAssociationOutput{}

        mockClient.On("GetFlowAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetFlowAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricData", func(t *testing.T) {
        input := &connect.GetMetricDataInput{}
        output := &connect.GetMetricDataOutput{}

        mockClient.On("GetMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricDataV2", func(t *testing.T) {
        input := &connect.GetMetricDataV2Input{}
        output := &connect.GetMetricDataV2Output{}

        mockClient.On("GetMetricDataV2", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricDataV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPromptFile", func(t *testing.T) {
        input := &connect.GetPromptFileInput{}
        output := &connect.GetPromptFileOutput{}

        mockClient.On("GetPromptFile", ctx, input).Return(output, nil)

        result, err := mockClient.GetPromptFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaskTemplate", func(t *testing.T) {
        input := &connect.GetTaskTemplateInput{}
        output := &connect.GetTaskTemplateOutput{}

        mockClient.On("GetTaskTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaskTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrafficDistribution", func(t *testing.T) {
        input := &connect.GetTrafficDistributionInput{}
        output := &connect.GetTrafficDistributionOutput{}

        mockClient.On("GetTrafficDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrafficDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportPhoneNumber", func(t *testing.T) {
        input := &connect.ImportPhoneNumberInput{}
        output := &connect.ImportPhoneNumberOutput{}

        mockClient.On("ImportPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.ImportPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgentStatuses", func(t *testing.T) {
        input := &connect.ListAgentStatusesInput{}
        output := &connect.ListAgentStatusesOutput{}

        mockClient.On("ListAgentStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgentStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalyticsDataAssociations", func(t *testing.T) {
        input := &connect.ListAnalyticsDataAssociationsInput{}
        output := &connect.ListAnalyticsDataAssociationsOutput{}

        mockClient.On("ListAnalyticsDataAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalyticsDataAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApprovedOrigins", func(t *testing.T) {
        input := &connect.ListApprovedOriginsInput{}
        output := &connect.ListApprovedOriginsOutput{}

        mockClient.On("ListApprovedOrigins", ctx, input).Return(output, nil)

        result, err := mockClient.ListApprovedOrigins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuthenticationProfiles", func(t *testing.T) {
        input := &connect.ListAuthenticationProfilesInput{}
        output := &connect.ListAuthenticationProfilesOutput{}

        mockClient.On("ListAuthenticationProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuthenticationProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBots", func(t *testing.T) {
        input := &connect.ListBotsInput{}
        output := &connect.ListBotsOutput{}

        mockClient.On("ListBots", ctx, input).Return(output, nil)

        result, err := mockClient.ListBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactEvaluations", func(t *testing.T) {
        input := &connect.ListContactEvaluationsInput{}
        output := &connect.ListContactEvaluationsOutput{}

        mockClient.On("ListContactEvaluations", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactEvaluations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactFlowModules", func(t *testing.T) {
        input := &connect.ListContactFlowModulesInput{}
        output := &connect.ListContactFlowModulesOutput{}

        mockClient.On("ListContactFlowModules", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactFlowModules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactFlows", func(t *testing.T) {
        input := &connect.ListContactFlowsInput{}
        output := &connect.ListContactFlowsOutput{}

        mockClient.On("ListContactFlows", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactReferences", func(t *testing.T) {
        input := &connect.ListContactReferencesInput{}
        output := &connect.ListContactReferencesOutput{}

        mockClient.On("ListContactReferences", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactReferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDefaultVocabularies", func(t *testing.T) {
        input := &connect.ListDefaultVocabulariesInput{}
        output := &connect.ListDefaultVocabulariesOutput{}

        mockClient.On("ListDefaultVocabularies", ctx, input).Return(output, nil)

        result, err := mockClient.ListDefaultVocabularies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEvaluationFormVersions", func(t *testing.T) {
        input := &connect.ListEvaluationFormVersionsInput{}
        output := &connect.ListEvaluationFormVersionsOutput{}

        mockClient.On("ListEvaluationFormVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEvaluationFormVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEvaluationForms", func(t *testing.T) {
        input := &connect.ListEvaluationFormsInput{}
        output := &connect.ListEvaluationFormsOutput{}

        mockClient.On("ListEvaluationForms", ctx, input).Return(output, nil)

        result, err := mockClient.ListEvaluationForms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowAssociations", func(t *testing.T) {
        input := &connect.ListFlowAssociationsInput{}
        output := &connect.ListFlowAssociationsOutput{}

        mockClient.On("ListFlowAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHoursOfOperations", func(t *testing.T) {
        input := &connect.ListHoursOfOperationsInput{}
        output := &connect.ListHoursOfOperationsOutput{}

        mockClient.On("ListHoursOfOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListHoursOfOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceAttributes", func(t *testing.T) {
        input := &connect.ListInstanceAttributesInput{}
        output := &connect.ListInstanceAttributesOutput{}

        mockClient.On("ListInstanceAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceStorageConfigs", func(t *testing.T) {
        input := &connect.ListInstanceStorageConfigsInput{}
        output := &connect.ListInstanceStorageConfigsOutput{}

        mockClient.On("ListInstanceStorageConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceStorageConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstances", func(t *testing.T) {
        input := &connect.ListInstancesInput{}
        output := &connect.ListInstancesOutput{}

        mockClient.On("ListInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntegrationAssociations", func(t *testing.T) {
        input := &connect.ListIntegrationAssociationsInput{}
        output := &connect.ListIntegrationAssociationsOutput{}

        mockClient.On("ListIntegrationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntegrationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLambdaFunctions", func(t *testing.T) {
        input := &connect.ListLambdaFunctionsInput{}
        output := &connect.ListLambdaFunctionsOutput{}

        mockClient.On("ListLambdaFunctions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLambdaFunctions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLexBots", func(t *testing.T) {
        input := &connect.ListLexBotsInput{}
        output := &connect.ListLexBotsOutput{}

        mockClient.On("ListLexBots", ctx, input).Return(output, nil)

        result, err := mockClient.ListLexBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumbers", func(t *testing.T) {
        input := &connect.ListPhoneNumbersInput{}
        output := &connect.ListPhoneNumbersOutput{}

        mockClient.On("ListPhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumbersV2", func(t *testing.T) {
        input := &connect.ListPhoneNumbersV2Input{}
        output := &connect.ListPhoneNumbersV2Output{}

        mockClient.On("ListPhoneNumbersV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumbersV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPredefinedAttributes", func(t *testing.T) {
        input := &connect.ListPredefinedAttributesInput{}
        output := &connect.ListPredefinedAttributesOutput{}

        mockClient.On("ListPredefinedAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ListPredefinedAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrompts", func(t *testing.T) {
        input := &connect.ListPromptsInput{}
        output := &connect.ListPromptsOutput{}

        mockClient.On("ListPrompts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrompts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueueQuickConnects", func(t *testing.T) {
        input := &connect.ListQueueQuickConnectsInput{}
        output := &connect.ListQueueQuickConnectsOutput{}

        mockClient.On("ListQueueQuickConnects", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueueQuickConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueues", func(t *testing.T) {
        input := &connect.ListQueuesInput{}
        output := &connect.ListQueuesOutput{}

        mockClient.On("ListQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQuickConnects", func(t *testing.T) {
        input := &connect.ListQuickConnectsInput{}
        output := &connect.ListQuickConnectsOutput{}

        mockClient.On("ListQuickConnects", ctx, input).Return(output, nil)

        result, err := mockClient.ListQuickConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRealtimeContactAnalysisSegmentsV2", func(t *testing.T) {
        input := &connect.ListRealtimeContactAnalysisSegmentsV2Input{}
        output := &connect.ListRealtimeContactAnalysisSegmentsV2Output{}

        mockClient.On("ListRealtimeContactAnalysisSegmentsV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListRealtimeContactAnalysisSegmentsV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutingProfileQueues", func(t *testing.T) {
        input := &connect.ListRoutingProfileQueuesInput{}
        output := &connect.ListRoutingProfileQueuesOutput{}

        mockClient.On("ListRoutingProfileQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutingProfileQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutingProfiles", func(t *testing.T) {
        input := &connect.ListRoutingProfilesInput{}
        output := &connect.ListRoutingProfilesOutput{}

        mockClient.On("ListRoutingProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutingProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &connect.ListRulesInput{}
        output := &connect.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityKeys", func(t *testing.T) {
        input := &connect.ListSecurityKeysInput{}
        output := &connect.ListSecurityKeysOutput{}

        mockClient.On("ListSecurityKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityProfileApplications", func(t *testing.T) {
        input := &connect.ListSecurityProfileApplicationsInput{}
        output := &connect.ListSecurityProfileApplicationsOutput{}

        mockClient.On("ListSecurityProfileApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityProfileApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityProfilePermissions", func(t *testing.T) {
        input := &connect.ListSecurityProfilePermissionsInput{}
        output := &connect.ListSecurityProfilePermissionsOutput{}

        mockClient.On("ListSecurityProfilePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityProfilePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityProfiles", func(t *testing.T) {
        input := &connect.ListSecurityProfilesInput{}
        output := &connect.ListSecurityProfilesOutput{}

        mockClient.On("ListSecurityProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &connect.ListTagsForResourceInput{}
        output := &connect.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaskTemplates", func(t *testing.T) {
        input := &connect.ListTaskTemplatesInput{}
        output := &connect.ListTaskTemplatesOutput{}

        mockClient.On("ListTaskTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaskTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficDistributionGroupUsers", func(t *testing.T) {
        input := &connect.ListTrafficDistributionGroupUsersInput{}
        output := &connect.ListTrafficDistributionGroupUsersOutput{}

        mockClient.On("ListTrafficDistributionGroupUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficDistributionGroupUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficDistributionGroups", func(t *testing.T) {
        input := &connect.ListTrafficDistributionGroupsInput{}
        output := &connect.ListTrafficDistributionGroupsOutput{}

        mockClient.On("ListTrafficDistributionGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficDistributionGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUseCases", func(t *testing.T) {
        input := &connect.ListUseCasesInput{}
        output := &connect.ListUseCasesOutput{}

        mockClient.On("ListUseCases", ctx, input).Return(output, nil)

        result, err := mockClient.ListUseCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserHierarchyGroups", func(t *testing.T) {
        input := &connect.ListUserHierarchyGroupsInput{}
        output := &connect.ListUserHierarchyGroupsOutput{}

        mockClient.On("ListUserHierarchyGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserHierarchyGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserProficiencies", func(t *testing.T) {
        input := &connect.ListUserProficienciesInput{}
        output := &connect.ListUserProficienciesOutput{}

        mockClient.On("ListUserProficiencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserProficiencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &connect.ListUsersInput{}
        output := &connect.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListViewVersions", func(t *testing.T) {
        input := &connect.ListViewVersionsInput{}
        output := &connect.ListViewVersionsOutput{}

        mockClient.On("ListViewVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListViewVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListViews", func(t *testing.T) {
        input := &connect.ListViewsInput{}
        output := &connect.ListViewsOutput{}

        mockClient.On("ListViews", ctx, input).Return(output, nil)

        result, err := mockClient.ListViews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMonitorContact", func(t *testing.T) {
        input := &connect.MonitorContactInput{}
        output := &connect.MonitorContactOutput{}

        mockClient.On("MonitorContact", ctx, input).Return(output, nil)

        result, err := mockClient.MonitorContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseContact", func(t *testing.T) {
        input := &connect.PauseContactInput{}
        output := &connect.PauseContactOutput{}

        mockClient.On("PauseContact", ctx, input).Return(output, nil)

        result, err := mockClient.PauseContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutUserStatus", func(t *testing.T) {
        input := &connect.PutUserStatusInput{}
        output := &connect.PutUserStatusOutput{}

        mockClient.On("PutUserStatus", ctx, input).Return(output, nil)

        result, err := mockClient.PutUserStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleasePhoneNumber", func(t *testing.T) {
        input := &connect.ReleasePhoneNumberInput{}
        output := &connect.ReleasePhoneNumberOutput{}

        mockClient.On("ReleasePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.ReleasePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplicateInstance", func(t *testing.T) {
        input := &connect.ReplicateInstanceInput{}
        output := &connect.ReplicateInstanceOutput{}

        mockClient.On("ReplicateInstance", ctx, input).Return(output, nil)

        result, err := mockClient.ReplicateInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeContact", func(t *testing.T) {
        input := &connect.ResumeContactInput{}
        output := &connect.ResumeContactOutput{}

        mockClient.On("ResumeContact", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeContactRecording", func(t *testing.T) {
        input := &connect.ResumeContactRecordingInput{}
        output := &connect.ResumeContactRecordingOutput{}

        mockClient.On("ResumeContactRecording", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeContactRecording(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAvailablePhoneNumbers", func(t *testing.T) {
        input := &connect.SearchAvailablePhoneNumbersInput{}
        output := &connect.SearchAvailablePhoneNumbersOutput{}

        mockClient.On("SearchAvailablePhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAvailablePhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchContactFlowModules", func(t *testing.T) {
        input := &connect.SearchContactFlowModulesInput{}
        output := &connect.SearchContactFlowModulesOutput{}

        mockClient.On("SearchContactFlowModules", ctx, input).Return(output, nil)

        result, err := mockClient.SearchContactFlowModules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchContactFlows", func(t *testing.T) {
        input := &connect.SearchContactFlowsInput{}
        output := &connect.SearchContactFlowsOutput{}

        mockClient.On("SearchContactFlows", ctx, input).Return(output, nil)

        result, err := mockClient.SearchContactFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchContacts", func(t *testing.T) {
        input := &connect.SearchContactsInput{}
        output := &connect.SearchContactsOutput{}

        mockClient.On("SearchContacts", ctx, input).Return(output, nil)

        result, err := mockClient.SearchContacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchHoursOfOperations", func(t *testing.T) {
        input := &connect.SearchHoursOfOperationsInput{}
        output := &connect.SearchHoursOfOperationsOutput{}

        mockClient.On("SearchHoursOfOperations", ctx, input).Return(output, nil)

        result, err := mockClient.SearchHoursOfOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchPredefinedAttributes", func(t *testing.T) {
        input := &connect.SearchPredefinedAttributesInput{}
        output := &connect.SearchPredefinedAttributesOutput{}

        mockClient.On("SearchPredefinedAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SearchPredefinedAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchPrompts", func(t *testing.T) {
        input := &connect.SearchPromptsInput{}
        output := &connect.SearchPromptsOutput{}

        mockClient.On("SearchPrompts", ctx, input).Return(output, nil)

        result, err := mockClient.SearchPrompts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchQueues", func(t *testing.T) {
        input := &connect.SearchQueuesInput{}
        output := &connect.SearchQueuesOutput{}

        mockClient.On("SearchQueues", ctx, input).Return(output, nil)

        result, err := mockClient.SearchQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchQuickConnects", func(t *testing.T) {
        input := &connect.SearchQuickConnectsInput{}
        output := &connect.SearchQuickConnectsOutput{}

        mockClient.On("SearchQuickConnects", ctx, input).Return(output, nil)

        result, err := mockClient.SearchQuickConnects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchResourceTags", func(t *testing.T) {
        input := &connect.SearchResourceTagsInput{}
        output := &connect.SearchResourceTagsOutput{}

        mockClient.On("SearchResourceTags", ctx, input).Return(output, nil)

        result, err := mockClient.SearchResourceTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchRoutingProfiles", func(t *testing.T) {
        input := &connect.SearchRoutingProfilesInput{}
        output := &connect.SearchRoutingProfilesOutput{}

        mockClient.On("SearchRoutingProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.SearchRoutingProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSecurityProfiles", func(t *testing.T) {
        input := &connect.SearchSecurityProfilesInput{}
        output := &connect.SearchSecurityProfilesOutput{}

        mockClient.On("SearchSecurityProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSecurityProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchUsers", func(t *testing.T) {
        input := &connect.SearchUsersInput{}
        output := &connect.SearchUsersOutput{}

        mockClient.On("SearchUsers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchVocabularies", func(t *testing.T) {
        input := &connect.SearchVocabulariesInput{}
        output := &connect.SearchVocabulariesOutput{}

        mockClient.On("SearchVocabularies", ctx, input).Return(output, nil)

        result, err := mockClient.SearchVocabularies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendChatIntegrationEvent", func(t *testing.T) {
        input := &connect.SendChatIntegrationEventInput{}
        output := &connect.SendChatIntegrationEventOutput{}

        mockClient.On("SendChatIntegrationEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SendChatIntegrationEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAttachedFileUpload", func(t *testing.T) {
        input := &connect.StartAttachedFileUploadInput{}
        output := &connect.StartAttachedFileUploadOutput{}

        mockClient.On("StartAttachedFileUpload", ctx, input).Return(output, nil)

        result, err := mockClient.StartAttachedFileUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartChatContact", func(t *testing.T) {
        input := &connect.StartChatContactInput{}
        output := &connect.StartChatContactOutput{}

        mockClient.On("StartChatContact", ctx, input).Return(output, nil)

        result, err := mockClient.StartChatContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContactEvaluation", func(t *testing.T) {
        input := &connect.StartContactEvaluationInput{}
        output := &connect.StartContactEvaluationOutput{}

        mockClient.On("StartContactEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.StartContactEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContactRecording", func(t *testing.T) {
        input := &connect.StartContactRecordingInput{}
        output := &connect.StartContactRecordingOutput{}

        mockClient.On("StartContactRecording", ctx, input).Return(output, nil)

        result, err := mockClient.StartContactRecording(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContactStreaming", func(t *testing.T) {
        input := &connect.StartContactStreamingInput{}
        output := &connect.StartContactStreamingOutput{}

        mockClient.On("StartContactStreaming", ctx, input).Return(output, nil)

        result, err := mockClient.StartContactStreaming(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartOutboundVoiceContact", func(t *testing.T) {
        input := &connect.StartOutboundVoiceContactInput{}
        output := &connect.StartOutboundVoiceContactOutput{}

        mockClient.On("StartOutboundVoiceContact", ctx, input).Return(output, nil)

        result, err := mockClient.StartOutboundVoiceContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTaskContact", func(t *testing.T) {
        input := &connect.StartTaskContactInput{}
        output := &connect.StartTaskContactOutput{}

        mockClient.On("StartTaskContact", ctx, input).Return(output, nil)

        result, err := mockClient.StartTaskContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWebRTCContact", func(t *testing.T) {
        input := &connect.StartWebRTCContactInput{}
        output := &connect.StartWebRTCContactOutput{}

        mockClient.On("StartWebRTCContact", ctx, input).Return(output, nil)

        result, err := mockClient.StartWebRTCContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopContact", func(t *testing.T) {
        input := &connect.StopContactInput{}
        output := &connect.StopContactOutput{}

        mockClient.On("StopContact", ctx, input).Return(output, nil)

        result, err := mockClient.StopContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopContactRecording", func(t *testing.T) {
        input := &connect.StopContactRecordingInput{}
        output := &connect.StopContactRecordingOutput{}

        mockClient.On("StopContactRecording", ctx, input).Return(output, nil)

        result, err := mockClient.StopContactRecording(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopContactStreaming", func(t *testing.T) {
        input := &connect.StopContactStreamingInput{}
        output := &connect.StopContactStreamingOutput{}

        mockClient.On("StopContactStreaming", ctx, input).Return(output, nil)

        result, err := mockClient.StopContactStreaming(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitContactEvaluation", func(t *testing.T) {
        input := &connect.SubmitContactEvaluationInput{}
        output := &connect.SubmitContactEvaluationOutput{}

        mockClient.On("SubmitContactEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitContactEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSuspendContactRecording", func(t *testing.T) {
        input := &connect.SuspendContactRecordingInput{}
        output := &connect.SuspendContactRecordingOutput{}

        mockClient.On("SuspendContactRecording", ctx, input).Return(output, nil)

        result, err := mockClient.SuspendContactRecording(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagContact", func(t *testing.T) {
        input := &connect.TagContactInput{}
        output := &connect.TagContactOutput{}

        mockClient.On("TagContact", ctx, input).Return(output, nil)

        result, err := mockClient.TagContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &connect.TagResourceInput{}
        output := &connect.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransferContact", func(t *testing.T) {
        input := &connect.TransferContactInput{}
        output := &connect.TransferContactOutput{}

        mockClient.On("TransferContact", ctx, input).Return(output, nil)

        result, err := mockClient.TransferContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagContact", func(t *testing.T) {
        input := &connect.UntagContactInput{}
        output := &connect.UntagContactOutput{}

        mockClient.On("UntagContact", ctx, input).Return(output, nil)

        result, err := mockClient.UntagContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &connect.UntagResourceInput{}
        output := &connect.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentStatus", func(t *testing.T) {
        input := &connect.UpdateAgentStatusInput{}
        output := &connect.UpdateAgentStatusOutput{}

        mockClient.On("UpdateAgentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuthenticationProfile", func(t *testing.T) {
        input := &connect.UpdateAuthenticationProfileInput{}
        output := &connect.UpdateAuthenticationProfileOutput{}

        mockClient.On("UpdateAuthenticationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuthenticationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContact", func(t *testing.T) {
        input := &connect.UpdateContactInput{}
        output := &connect.UpdateContactOutput{}

        mockClient.On("UpdateContact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactAttributes", func(t *testing.T) {
        input := &connect.UpdateContactAttributesInput{}
        output := &connect.UpdateContactAttributesOutput{}

        mockClient.On("UpdateContactAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactEvaluation", func(t *testing.T) {
        input := &connect.UpdateContactEvaluationInput{}
        output := &connect.UpdateContactEvaluationOutput{}

        mockClient.On("UpdateContactEvaluation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactEvaluation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactFlowContent", func(t *testing.T) {
        input := &connect.UpdateContactFlowContentInput{}
        output := &connect.UpdateContactFlowContentOutput{}

        mockClient.On("UpdateContactFlowContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactFlowContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactFlowMetadata", func(t *testing.T) {
        input := &connect.UpdateContactFlowMetadataInput{}
        output := &connect.UpdateContactFlowMetadataOutput{}

        mockClient.On("UpdateContactFlowMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactFlowMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactFlowModuleContent", func(t *testing.T) {
        input := &connect.UpdateContactFlowModuleContentInput{}
        output := &connect.UpdateContactFlowModuleContentOutput{}

        mockClient.On("UpdateContactFlowModuleContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactFlowModuleContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactFlowModuleMetadata", func(t *testing.T) {
        input := &connect.UpdateContactFlowModuleMetadataInput{}
        output := &connect.UpdateContactFlowModuleMetadataOutput{}

        mockClient.On("UpdateContactFlowModuleMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactFlowModuleMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactFlowName", func(t *testing.T) {
        input := &connect.UpdateContactFlowNameInput{}
        output := &connect.UpdateContactFlowNameOutput{}

        mockClient.On("UpdateContactFlowName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactFlowName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactRoutingData", func(t *testing.T) {
        input := &connect.UpdateContactRoutingDataInput{}
        output := &connect.UpdateContactRoutingDataOutput{}

        mockClient.On("UpdateContactRoutingData", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactRoutingData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactSchedule", func(t *testing.T) {
        input := &connect.UpdateContactScheduleInput{}
        output := &connect.UpdateContactScheduleOutput{}

        mockClient.On("UpdateContactSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEvaluationForm", func(t *testing.T) {
        input := &connect.UpdateEvaluationFormInput{}
        output := &connect.UpdateEvaluationFormOutput{}

        mockClient.On("UpdateEvaluationForm", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEvaluationForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHoursOfOperation", func(t *testing.T) {
        input := &connect.UpdateHoursOfOperationInput{}
        output := &connect.UpdateHoursOfOperationOutput{}

        mockClient.On("UpdateHoursOfOperation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHoursOfOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceAttribute", func(t *testing.T) {
        input := &connect.UpdateInstanceAttributeInput{}
        output := &connect.UpdateInstanceAttributeOutput{}

        mockClient.On("UpdateInstanceAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceStorageConfig", func(t *testing.T) {
        input := &connect.UpdateInstanceStorageConfigInput{}
        output := &connect.UpdateInstanceStorageConfigOutput{}

        mockClient.On("UpdateInstanceStorageConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceStorageConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateParticipantRoleConfig", func(t *testing.T) {
        input := &connect.UpdateParticipantRoleConfigInput{}
        output := &connect.UpdateParticipantRoleConfigOutput{}

        mockClient.On("UpdateParticipantRoleConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateParticipantRoleConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumber", func(t *testing.T) {
        input := &connect.UpdatePhoneNumberInput{}
        output := &connect.UpdatePhoneNumberOutput{}

        mockClient.On("UpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumberMetadata", func(t *testing.T) {
        input := &connect.UpdatePhoneNumberMetadataInput{}
        output := &connect.UpdatePhoneNumberMetadataOutput{}

        mockClient.On("UpdatePhoneNumberMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumberMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePredefinedAttribute", func(t *testing.T) {
        input := &connect.UpdatePredefinedAttributeInput{}
        output := &connect.UpdatePredefinedAttributeOutput{}

        mockClient.On("UpdatePredefinedAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePredefinedAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrompt", func(t *testing.T) {
        input := &connect.UpdatePromptInput{}
        output := &connect.UpdatePromptOutput{}

        mockClient.On("UpdatePrompt", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrompt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueHoursOfOperation", func(t *testing.T) {
        input := &connect.UpdateQueueHoursOfOperationInput{}
        output := &connect.UpdateQueueHoursOfOperationOutput{}

        mockClient.On("UpdateQueueHoursOfOperation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueHoursOfOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueMaxContacts", func(t *testing.T) {
        input := &connect.UpdateQueueMaxContactsInput{}
        output := &connect.UpdateQueueMaxContactsOutput{}

        mockClient.On("UpdateQueueMaxContacts", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueMaxContacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueName", func(t *testing.T) {
        input := &connect.UpdateQueueNameInput{}
        output := &connect.UpdateQueueNameOutput{}

        mockClient.On("UpdateQueueName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueOutboundCallerConfig", func(t *testing.T) {
        input := &connect.UpdateQueueOutboundCallerConfigInput{}
        output := &connect.UpdateQueueOutboundCallerConfigOutput{}

        mockClient.On("UpdateQueueOutboundCallerConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueOutboundCallerConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueStatus", func(t *testing.T) {
        input := &connect.UpdateQueueStatusInput{}
        output := &connect.UpdateQueueStatusOutput{}

        mockClient.On("UpdateQueueStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuickConnectConfig", func(t *testing.T) {
        input := &connect.UpdateQuickConnectConfigInput{}
        output := &connect.UpdateQuickConnectConfigOutput{}

        mockClient.On("UpdateQuickConnectConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuickConnectConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuickConnectName", func(t *testing.T) {
        input := &connect.UpdateQuickConnectNameInput{}
        output := &connect.UpdateQuickConnectNameOutput{}

        mockClient.On("UpdateQuickConnectName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuickConnectName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingProfileAgentAvailabilityTimer", func(t *testing.T) {
        input := &connect.UpdateRoutingProfileAgentAvailabilityTimerInput{}
        output := &connect.UpdateRoutingProfileAgentAvailabilityTimerOutput{}

        mockClient.On("UpdateRoutingProfileAgentAvailabilityTimer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingProfileAgentAvailabilityTimer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingProfileConcurrency", func(t *testing.T) {
        input := &connect.UpdateRoutingProfileConcurrencyInput{}
        output := &connect.UpdateRoutingProfileConcurrencyOutput{}

        mockClient.On("UpdateRoutingProfileConcurrency", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingProfileConcurrency(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingProfileDefaultOutboundQueue", func(t *testing.T) {
        input := &connect.UpdateRoutingProfileDefaultOutboundQueueInput{}
        output := &connect.UpdateRoutingProfileDefaultOutboundQueueOutput{}

        mockClient.On("UpdateRoutingProfileDefaultOutboundQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingProfileDefaultOutboundQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingProfileName", func(t *testing.T) {
        input := &connect.UpdateRoutingProfileNameInput{}
        output := &connect.UpdateRoutingProfileNameOutput{}

        mockClient.On("UpdateRoutingProfileName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingProfileName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingProfileQueues", func(t *testing.T) {
        input := &connect.UpdateRoutingProfileQueuesInput{}
        output := &connect.UpdateRoutingProfileQueuesOutput{}

        mockClient.On("UpdateRoutingProfileQueues", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingProfileQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRule", func(t *testing.T) {
        input := &connect.UpdateRuleInput{}
        output := &connect.UpdateRuleOutput{}

        mockClient.On("UpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityProfile", func(t *testing.T) {
        input := &connect.UpdateSecurityProfileInput{}
        output := &connect.UpdateSecurityProfileOutput{}

        mockClient.On("UpdateSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTaskTemplate", func(t *testing.T) {
        input := &connect.UpdateTaskTemplateInput{}
        output := &connect.UpdateTaskTemplateOutput{}

        mockClient.On("UpdateTaskTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTaskTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrafficDistribution", func(t *testing.T) {
        input := &connect.UpdateTrafficDistributionInput{}
        output := &connect.UpdateTrafficDistributionOutput{}

        mockClient.On("UpdateTrafficDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrafficDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserHierarchy", func(t *testing.T) {
        input := &connect.UpdateUserHierarchyInput{}
        output := &connect.UpdateUserHierarchyOutput{}

        mockClient.On("UpdateUserHierarchy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserHierarchy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserHierarchyGroupName", func(t *testing.T) {
        input := &connect.UpdateUserHierarchyGroupNameInput{}
        output := &connect.UpdateUserHierarchyGroupNameOutput{}

        mockClient.On("UpdateUserHierarchyGroupName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserHierarchyGroupName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserHierarchyStructure", func(t *testing.T) {
        input := &connect.UpdateUserHierarchyStructureInput{}
        output := &connect.UpdateUserHierarchyStructureOutput{}

        mockClient.On("UpdateUserHierarchyStructure", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserHierarchyStructure(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserIdentityInfo", func(t *testing.T) {
        input := &connect.UpdateUserIdentityInfoInput{}
        output := &connect.UpdateUserIdentityInfoOutput{}

        mockClient.On("UpdateUserIdentityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserIdentityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserPhoneConfig", func(t *testing.T) {
        input := &connect.UpdateUserPhoneConfigInput{}
        output := &connect.UpdateUserPhoneConfigOutput{}

        mockClient.On("UpdateUserPhoneConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserPhoneConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserProficiencies", func(t *testing.T) {
        input := &connect.UpdateUserProficienciesInput{}
        output := &connect.UpdateUserProficienciesOutput{}

        mockClient.On("UpdateUserProficiencies", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserProficiencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserRoutingProfile", func(t *testing.T) {
        input := &connect.UpdateUserRoutingProfileInput{}
        output := &connect.UpdateUserRoutingProfileOutput{}

        mockClient.On("UpdateUserRoutingProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserRoutingProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserSecurityProfiles", func(t *testing.T) {
        input := &connect.UpdateUserSecurityProfilesInput{}
        output := &connect.UpdateUserSecurityProfilesOutput{}

        mockClient.On("UpdateUserSecurityProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserSecurityProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateViewContent", func(t *testing.T) {
        input := &connect.UpdateViewContentInput{}
        output := &connect.UpdateViewContentOutput{}

        mockClient.On("UpdateViewContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateViewContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateViewMetadata", func(t *testing.T) {
        input := &connect.UpdateViewMetadataInput{}
        output := &connect.UpdateViewMetadataOutput{}

        mockClient.On("UpdateViewMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateViewMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
