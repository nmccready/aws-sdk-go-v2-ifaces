package pinpointsmsvoicev2_test

// tests for the pinpointsmsvoicev2 service interface for this ../../../service/pinpointsmsvoicev2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointsmsvoicev2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointsmsvoicev2/pinpointsmsvoicev2_iface"
	"github.com/stretchr/testify/assert"
)

func TestPinpointsmsvoicev2ServiceCanBeMocked(t *testing.T) {
	var iface pinpointsmsvoicev2_iface.IClient
	iface = &pinpointsmsvoicev2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pinpointsmsvoicev2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateOriginationIdentity", func(t *testing.T) {
        input := &pinpointsmsvoicev2.AssociateOriginationIdentityInput{}
        output := &pinpointsmsvoicev2.AssociateOriginationIdentityOutput{}

        mockClient.On("AssociateOriginationIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateOriginationIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.AssociateProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.AssociateProtectConfigurationOutput{}

        mockClient.On("AssociateProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSet", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateConfigurationSetInput{}
        output := &pinpointsmsvoicev2.CreateConfigurationSetOutput{}

        mockClient.On("CreateConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateEventDestinationInput{}
        output := &pinpointsmsvoicev2.CreateEventDestinationOutput{}

        mockClient.On("CreateEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOptOutList", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateOptOutListInput{}
        output := &pinpointsmsvoicev2.CreateOptOutListOutput{}

        mockClient.On("CreateOptOutList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOptOutList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePool", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreatePoolInput{}
        output := &pinpointsmsvoicev2.CreatePoolOutput{}

        mockClient.On("CreatePool", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.CreateProtectConfigurationOutput{}

        mockClient.On("CreateProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateRegistrationInput{}
        output := &pinpointsmsvoicev2.CreateRegistrationOutput{}

        mockClient.On("CreateRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistrationAssociation", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateRegistrationAssociationInput{}
        output := &pinpointsmsvoicev2.CreateRegistrationAssociationOutput{}

        mockClient.On("CreateRegistrationAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistrationAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistrationAttachment", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateRegistrationAttachmentInput{}
        output := &pinpointsmsvoicev2.CreateRegistrationAttachmentOutput{}

        mockClient.On("CreateRegistrationAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistrationAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistrationVersion", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateRegistrationVersionInput{}
        output := &pinpointsmsvoicev2.CreateRegistrationVersionOutput{}

        mockClient.On("CreateRegistrationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistrationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVerifiedDestinationNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.CreateVerifiedDestinationNumberInput{}
        output := &pinpointsmsvoicev2.CreateVerifiedDestinationNumberOutput{}

        mockClient.On("CreateVerifiedDestinationNumber", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVerifiedDestinationNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountDefaultProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteAccountDefaultProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.DeleteAccountDefaultProtectConfigurationOutput{}

        mockClient.On("DeleteAccountDefaultProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountDefaultProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSet", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteConfigurationSetInput{}
        output := &pinpointsmsvoicev2.DeleteConfigurationSetOutput{}

        mockClient.On("DeleteConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDefaultMessageType", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteDefaultMessageTypeInput{}
        output := &pinpointsmsvoicev2.DeleteDefaultMessageTypeOutput{}

        mockClient.On("DeleteDefaultMessageType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDefaultMessageType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDefaultSenderId", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteDefaultSenderIdInput{}
        output := &pinpointsmsvoicev2.DeleteDefaultSenderIdOutput{}

        mockClient.On("DeleteDefaultSenderId", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDefaultSenderId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteEventDestinationInput{}
        output := &pinpointsmsvoicev2.DeleteEventDestinationOutput{}

        mockClient.On("DeleteEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyword", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteKeywordInput{}
        output := &pinpointsmsvoicev2.DeleteKeywordOutput{}

        mockClient.On("DeleteKeyword", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMediaMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteMediaMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.DeleteMediaMessageSpendLimitOverrideOutput{}

        mockClient.On("DeleteMediaMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMediaMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOptOutList", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteOptOutListInput{}
        output := &pinpointsmsvoicev2.DeleteOptOutListOutput{}

        mockClient.On("DeleteOptOutList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOptOutList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOptedOutNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteOptedOutNumberInput{}
        output := &pinpointsmsvoicev2.DeleteOptedOutNumberOutput{}

        mockClient.On("DeleteOptedOutNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOptedOutNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePool", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeletePoolInput{}
        output := &pinpointsmsvoicev2.DeletePoolOutput{}

        mockClient.On("DeletePool", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.DeleteProtectConfigurationOutput{}

        mockClient.On("DeleteProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteRegistrationInput{}
        output := &pinpointsmsvoicev2.DeleteRegistrationOutput{}

        mockClient.On("DeleteRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistrationAttachment", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteRegistrationAttachmentInput{}
        output := &pinpointsmsvoicev2.DeleteRegistrationAttachmentOutput{}

        mockClient.On("DeleteRegistrationAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistrationAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistrationFieldValue", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteRegistrationFieldValueInput{}
        output := &pinpointsmsvoicev2.DeleteRegistrationFieldValueOutput{}

        mockClient.On("DeleteRegistrationFieldValue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistrationFieldValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTextMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideOutput{}

        mockClient.On("DeleteTextMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTextMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedDestinationNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteVerifiedDestinationNumberInput{}
        output := &pinpointsmsvoicev2.DeleteVerifiedDestinationNumberOutput{}

        mockClient.On("DeleteVerifiedDestinationNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedDestinationNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideOutput{}

        mockClient.On("DeleteVoiceMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeAccountAttributesInput{}
        output := &pinpointsmsvoicev2.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountLimits", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeAccountLimitsInput{}
        output := &pinpointsmsvoicev2.DescribeAccountLimitsOutput{}

        mockClient.On("DescribeAccountLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationSets", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeConfigurationSetsInput{}
        output := &pinpointsmsvoicev2.DescribeConfigurationSetsOutput{}

        mockClient.On("DescribeConfigurationSets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeywords", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeKeywordsInput{}
        output := &pinpointsmsvoicev2.DescribeKeywordsOutput{}

        mockClient.On("DescribeKeywords", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeywords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOptOutLists", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeOptOutListsInput{}
        output := &pinpointsmsvoicev2.DescribeOptOutListsOutput{}

        mockClient.On("DescribeOptOutLists", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOptOutLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOptedOutNumbers", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeOptedOutNumbersInput{}
        output := &pinpointsmsvoicev2.DescribeOptedOutNumbersOutput{}

        mockClient.On("DescribeOptedOutNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOptedOutNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePhoneNumbers", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribePhoneNumbersInput{}
        output := &pinpointsmsvoicev2.DescribePhoneNumbersOutput{}

        mockClient.On("DescribePhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePools", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribePoolsInput{}
        output := &pinpointsmsvoicev2.DescribePoolsOutput{}

        mockClient.On("DescribePools", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProtectConfigurations", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeProtectConfigurationsInput{}
        output := &pinpointsmsvoicev2.DescribeProtectConfigurationsOutput{}

        mockClient.On("DescribeProtectConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProtectConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationAttachments", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationAttachmentsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationAttachmentsOutput{}

        mockClient.On("DescribeRegistrationAttachments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationAttachments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationFieldDefinitions", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsOutput{}

        mockClient.On("DescribeRegistrationFieldDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationFieldDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationFieldValues", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationFieldValuesInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationFieldValuesOutput{}

        mockClient.On("DescribeRegistrationFieldValues", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationFieldValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationSectionDefinitions", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsOutput{}

        mockClient.On("DescribeRegistrationSectionDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationSectionDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationTypeDefinitions", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsOutput{}

        mockClient.On("DescribeRegistrationTypeDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationTypeDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrationVersions", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationVersionsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationVersionsOutput{}

        mockClient.On("DescribeRegistrationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistrations", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeRegistrationsInput{}
        output := &pinpointsmsvoicev2.DescribeRegistrationsOutput{}

        mockClient.On("DescribeRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSenderIds", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeSenderIdsInput{}
        output := &pinpointsmsvoicev2.DescribeSenderIdsOutput{}

        mockClient.On("DescribeSenderIds", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSenderIds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpendLimits", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeSpendLimitsInput{}
        output := &pinpointsmsvoicev2.DescribeSpendLimitsOutput{}

        mockClient.On("DescribeSpendLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpendLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVerifiedDestinationNumbers", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersInput{}
        output := &pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersOutput{}

        mockClient.On("DescribeVerifiedDestinationNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVerifiedDestinationNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateOriginationIdentity", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DisassociateOriginationIdentityInput{}
        output := &pinpointsmsvoicev2.DisassociateOriginationIdentityOutput{}

        mockClient.On("DisassociateOriginationIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateOriginationIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DisassociateProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.DisassociateProtectConfigurationOutput{}

        mockClient.On("DisassociateProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscardRegistrationVersion", func(t *testing.T) {
        input := &pinpointsmsvoicev2.DiscardRegistrationVersionInput{}
        output := &pinpointsmsvoicev2.DiscardRegistrationVersionOutput{}

        mockClient.On("DiscardRegistrationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DiscardRegistrationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProtectConfigurationCountryRuleSet", func(t *testing.T) {
        input := &pinpointsmsvoicev2.GetProtectConfigurationCountryRuleSetInput{}
        output := &pinpointsmsvoicev2.GetProtectConfigurationCountryRuleSetOutput{}

        mockClient.On("GetProtectConfigurationCountryRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetProtectConfigurationCountryRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPoolOriginationIdentities", func(t *testing.T) {
        input := &pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput{}
        output := &pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput{}

        mockClient.On("ListPoolOriginationIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListPoolOriginationIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegistrationAssociations", func(t *testing.T) {
        input := &pinpointsmsvoicev2.ListRegistrationAssociationsInput{}
        output := &pinpointsmsvoicev2.ListRegistrationAssociationsOutput{}

        mockClient.On("ListRegistrationAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegistrationAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pinpointsmsvoicev2.ListTagsForResourceInput{}
        output := &pinpointsmsvoicev2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutKeyword", func(t *testing.T) {
        input := &pinpointsmsvoicev2.PutKeywordInput{}
        output := &pinpointsmsvoicev2.PutKeywordOutput{}

        mockClient.On("PutKeyword", ctx, input).Return(output, nil)

        result, err := mockClient.PutKeyword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOptedOutNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.PutOptedOutNumberInput{}
        output := &pinpointsmsvoicev2.PutOptedOutNumberOutput{}

        mockClient.On("PutOptedOutNumber", ctx, input).Return(output, nil)

        result, err := mockClient.PutOptedOutNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRegistrationFieldValue", func(t *testing.T) {
        input := &pinpointsmsvoicev2.PutRegistrationFieldValueInput{}
        output := &pinpointsmsvoicev2.PutRegistrationFieldValueOutput{}

        mockClient.On("PutRegistrationFieldValue", ctx, input).Return(output, nil)

        result, err := mockClient.PutRegistrationFieldValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleasePhoneNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.ReleasePhoneNumberInput{}
        output := &pinpointsmsvoicev2.ReleasePhoneNumberOutput{}

        mockClient.On("ReleasePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.ReleasePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseSenderId", func(t *testing.T) {
        input := &pinpointsmsvoicev2.ReleaseSenderIdInput{}
        output := &pinpointsmsvoicev2.ReleaseSenderIdOutput{}

        mockClient.On("ReleaseSenderId", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseSenderId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestPhoneNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.RequestPhoneNumberInput{}
        output := &pinpointsmsvoicev2.RequestPhoneNumberOutput{}

        mockClient.On("RequestPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.RequestPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestSenderId", func(t *testing.T) {
        input := &pinpointsmsvoicev2.RequestSenderIdInput{}
        output := &pinpointsmsvoicev2.RequestSenderIdOutput{}

        mockClient.On("RequestSenderId", ctx, input).Return(output, nil)

        result, err := mockClient.RequestSenderId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDestinationNumberVerificationCode", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SendDestinationNumberVerificationCodeInput{}
        output := &pinpointsmsvoicev2.SendDestinationNumberVerificationCodeOutput{}

        mockClient.On("SendDestinationNumberVerificationCode", ctx, input).Return(output, nil)

        result, err := mockClient.SendDestinationNumberVerificationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendMediaMessage", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SendMediaMessageInput{}
        output := &pinpointsmsvoicev2.SendMediaMessageOutput{}

        mockClient.On("SendMediaMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendMediaMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTextMessage", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SendTextMessageInput{}
        output := &pinpointsmsvoicev2.SendTextMessageOutput{}

        mockClient.On("SendTextMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendTextMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendVoiceMessage", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SendVoiceMessageInput{}
        output := &pinpointsmsvoicev2.SendVoiceMessageOutput{}

        mockClient.On("SendVoiceMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendVoiceMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetAccountDefaultProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetAccountDefaultProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.SetAccountDefaultProtectConfigurationOutput{}

        mockClient.On("SetAccountDefaultProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.SetAccountDefaultProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultMessageType", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetDefaultMessageTypeInput{}
        output := &pinpointsmsvoicev2.SetDefaultMessageTypeOutput{}

        mockClient.On("SetDefaultMessageType", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultMessageType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultSenderId", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetDefaultSenderIdInput{}
        output := &pinpointsmsvoicev2.SetDefaultSenderIdOutput{}

        mockClient.On("SetDefaultSenderId", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultSenderId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetMediaMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetMediaMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.SetMediaMessageSpendLimitOverrideOutput{}

        mockClient.On("SetMediaMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.SetMediaMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTextMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideOutput{}

        mockClient.On("SetTextMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.SetTextMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetVoiceMessageSpendLimitOverride", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput{}
        output := &pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideOutput{}

        mockClient.On("SetVoiceMessageSpendLimitOverride", ctx, input).Return(output, nil)

        result, err := mockClient.SetVoiceMessageSpendLimitOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitRegistrationVersion", func(t *testing.T) {
        input := &pinpointsmsvoicev2.SubmitRegistrationVersionInput{}
        output := &pinpointsmsvoicev2.SubmitRegistrationVersionOutput{}

        mockClient.On("SubmitRegistrationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitRegistrationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pinpointsmsvoicev2.TagResourceInput{}
        output := &pinpointsmsvoicev2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UntagResourceInput{}
        output := &pinpointsmsvoicev2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventDestination", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdateEventDestinationInput{}
        output := &pinpointsmsvoicev2.UpdateEventDestinationOutput{}

        mockClient.On("UpdateEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdatePhoneNumberInput{}
        output := &pinpointsmsvoicev2.UpdatePhoneNumberOutput{}

        mockClient.On("UpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePool", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdatePoolInput{}
        output := &pinpointsmsvoicev2.UpdatePoolOutput{}

        mockClient.On("UpdatePool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProtectConfiguration", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdateProtectConfigurationInput{}
        output := &pinpointsmsvoicev2.UpdateProtectConfigurationOutput{}

        mockClient.On("UpdateProtectConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProtectConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProtectConfigurationCountryRuleSet", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdateProtectConfigurationCountryRuleSetInput{}
        output := &pinpointsmsvoicev2.UpdateProtectConfigurationCountryRuleSetOutput{}

        mockClient.On("UpdateProtectConfigurationCountryRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProtectConfigurationCountryRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSenderId", func(t *testing.T) {
        input := &pinpointsmsvoicev2.UpdateSenderIdInput{}
        output := &pinpointsmsvoicev2.UpdateSenderIdOutput{}

        mockClient.On("UpdateSenderId", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSenderId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyDestinationNumber", func(t *testing.T) {
        input := &pinpointsmsvoicev2.VerifyDestinationNumberInput{}
        output := &pinpointsmsvoicev2.VerifyDestinationNumberOutput{}

        mockClient.On("VerifyDestinationNumber", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyDestinationNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
