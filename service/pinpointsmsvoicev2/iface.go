
package pinpointsmsvoicev2

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
)

// IClient defines the interface for pinpointsmsvoicev2
type IClient interface {
 Options() Options 
 AssociateOriginationIdentity(ctx context.Context, params *AssociateOriginationIdentityInput, optFns ...func(*Options)) (*AssociateOriginationIdentityOutput, error) 
 AssociateProtectConfiguration(ctx context.Context, params *AssociateProtectConfigurationInput, optFns ...func(*Options)) (*AssociateProtectConfigurationOutput, error) 
 CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) 
 CreateEventDestination(ctx context.Context, params *CreateEventDestinationInput, optFns ...func(*Options)) (*CreateEventDestinationOutput, error) 
 CreateOptOutList(ctx context.Context, params *CreateOptOutListInput, optFns ...func(*Options)) (*CreateOptOutListOutput, error) 
 CreatePool(ctx context.Context, params *CreatePoolInput, optFns ...func(*Options)) (*CreatePoolOutput, error) 
 CreateProtectConfiguration(ctx context.Context, params *CreateProtectConfigurationInput, optFns ...func(*Options)) (*CreateProtectConfigurationOutput, error) 
 CreateRegistration(ctx context.Context, params *CreateRegistrationInput, optFns ...func(*Options)) (*CreateRegistrationOutput, error) 
 CreateRegistrationAssociation(ctx context.Context, params *CreateRegistrationAssociationInput, optFns ...func(*Options)) (*CreateRegistrationAssociationOutput, error) 
 CreateRegistrationAttachment(ctx context.Context, params *CreateRegistrationAttachmentInput, optFns ...func(*Options)) (*CreateRegistrationAttachmentOutput, error) 
 CreateRegistrationVersion(ctx context.Context, params *CreateRegistrationVersionInput, optFns ...func(*Options)) (*CreateRegistrationVersionOutput, error) 
 CreateVerifiedDestinationNumber(ctx context.Context, params *CreateVerifiedDestinationNumberInput, optFns ...func(*Options)) (*CreateVerifiedDestinationNumberOutput, error) 
 DeleteAccountDefaultProtectConfiguration(ctx context.Context, params *DeleteAccountDefaultProtectConfigurationInput, optFns ...func(*Options)) (*DeleteAccountDefaultProtectConfigurationOutput, error) 
 DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) 
 DeleteDefaultMessageType(ctx context.Context, params *DeleteDefaultMessageTypeInput, optFns ...func(*Options)) (*DeleteDefaultMessageTypeOutput, error) 
 DeleteDefaultSenderId(ctx context.Context, params *DeleteDefaultSenderIdInput, optFns ...func(*Options)) (*DeleteDefaultSenderIdOutput, error) 
 DeleteEventDestination(ctx context.Context, params *DeleteEventDestinationInput, optFns ...func(*Options)) (*DeleteEventDestinationOutput, error) 
 DeleteKeyword(ctx context.Context, params *DeleteKeywordInput, optFns ...func(*Options)) (*DeleteKeywordOutput, error) 
 DeleteMediaMessageSpendLimitOverride(ctx context.Context, params *DeleteMediaMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*DeleteMediaMessageSpendLimitOverrideOutput, error) 
 DeleteOptOutList(ctx context.Context, params *DeleteOptOutListInput, optFns ...func(*Options)) (*DeleteOptOutListOutput, error) 
 DeleteOptedOutNumber(ctx context.Context, params *DeleteOptedOutNumberInput, optFns ...func(*Options)) (*DeleteOptedOutNumberOutput, error) 
 DeletePool(ctx context.Context, params *DeletePoolInput, optFns ...func(*Options)) (*DeletePoolOutput, error) 
 DeleteProtectConfiguration(ctx context.Context, params *DeleteProtectConfigurationInput, optFns ...func(*Options)) (*DeleteProtectConfigurationOutput, error) 
 DeleteRegistration(ctx context.Context, params *DeleteRegistrationInput, optFns ...func(*Options)) (*DeleteRegistrationOutput, error) 
 DeleteRegistrationAttachment(ctx context.Context, params *DeleteRegistrationAttachmentInput, optFns ...func(*Options)) (*DeleteRegistrationAttachmentOutput, error) 
 DeleteRegistrationFieldValue(ctx context.Context, params *DeleteRegistrationFieldValueInput, optFns ...func(*Options)) (*DeleteRegistrationFieldValueOutput, error) 
 DeleteTextMessageSpendLimitOverride(ctx context.Context, params *DeleteTextMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*DeleteTextMessageSpendLimitOverrideOutput, error) 
 DeleteVerifiedDestinationNumber(ctx context.Context, params *DeleteVerifiedDestinationNumberInput, optFns ...func(*Options)) (*DeleteVerifiedDestinationNumberOutput, error) 
 DeleteVoiceMessageSpendLimitOverride(ctx context.Context, params *DeleteVoiceMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*DeleteVoiceMessageSpendLimitOverrideOutput, error) 
 DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) 
 DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) 
 DescribeConfigurationSets(ctx context.Context, params *DescribeConfigurationSetsInput, optFns ...func(*Options)) (*DescribeConfigurationSetsOutput, error) 
 DescribeKeywords(ctx context.Context, params *DescribeKeywordsInput, optFns ...func(*Options)) (*DescribeKeywordsOutput, error) 
 DescribeOptOutLists(ctx context.Context, params *DescribeOptOutListsInput, optFns ...func(*Options)) (*DescribeOptOutListsOutput, error) 
 DescribeOptedOutNumbers(ctx context.Context, params *DescribeOptedOutNumbersInput, optFns ...func(*Options)) (*DescribeOptedOutNumbersOutput, error) 
 DescribePhoneNumbers(ctx context.Context, params *DescribePhoneNumbersInput, optFns ...func(*Options)) (*DescribePhoneNumbersOutput, error) 
 DescribePools(ctx context.Context, params *DescribePoolsInput, optFns ...func(*Options)) (*DescribePoolsOutput, error) 
 DescribeProtectConfigurations(ctx context.Context, params *DescribeProtectConfigurationsInput, optFns ...func(*Options)) (*DescribeProtectConfigurationsOutput, error) 
 DescribeRegistrationAttachments(ctx context.Context, params *DescribeRegistrationAttachmentsInput, optFns ...func(*Options)) (*DescribeRegistrationAttachmentsOutput, error) 
 DescribeRegistrationFieldDefinitions(ctx context.Context, params *DescribeRegistrationFieldDefinitionsInput, optFns ...func(*Options)) (*DescribeRegistrationFieldDefinitionsOutput, error) 
 DescribeRegistrationFieldValues(ctx context.Context, params *DescribeRegistrationFieldValuesInput, optFns ...func(*Options)) (*DescribeRegistrationFieldValuesOutput, error) 
 DescribeRegistrationSectionDefinitions(ctx context.Context, params *DescribeRegistrationSectionDefinitionsInput, optFns ...func(*Options)) (*DescribeRegistrationSectionDefinitionsOutput, error) 
 DescribeRegistrationTypeDefinitions(ctx context.Context, params *DescribeRegistrationTypeDefinitionsInput, optFns ...func(*Options)) (*DescribeRegistrationTypeDefinitionsOutput, error) 
 DescribeRegistrationVersions(ctx context.Context, params *DescribeRegistrationVersionsInput, optFns ...func(*Options)) (*DescribeRegistrationVersionsOutput, error) 
 DescribeRegistrations(ctx context.Context, params *DescribeRegistrationsInput, optFns ...func(*Options)) (*DescribeRegistrationsOutput, error) 
 DescribeSenderIds(ctx context.Context, params *DescribeSenderIdsInput, optFns ...func(*Options)) (*DescribeSenderIdsOutput, error) 
 DescribeSpendLimits(ctx context.Context, params *DescribeSpendLimitsInput, optFns ...func(*Options)) (*DescribeSpendLimitsOutput, error) 
 DescribeVerifiedDestinationNumbers(ctx context.Context, params *DescribeVerifiedDestinationNumbersInput, optFns ...func(*Options)) (*DescribeVerifiedDestinationNumbersOutput, error) 
 DisassociateOriginationIdentity(ctx context.Context, params *DisassociateOriginationIdentityInput, optFns ...func(*Options)) (*DisassociateOriginationIdentityOutput, error) 
 DisassociateProtectConfiguration(ctx context.Context, params *DisassociateProtectConfigurationInput, optFns ...func(*Options)) (*DisassociateProtectConfigurationOutput, error) 
 DiscardRegistrationVersion(ctx context.Context, params *DiscardRegistrationVersionInput, optFns ...func(*Options)) (*DiscardRegistrationVersionOutput, error) 
 GetProtectConfigurationCountryRuleSet(ctx context.Context, params *GetProtectConfigurationCountryRuleSetInput, optFns ...func(*Options)) (*GetProtectConfigurationCountryRuleSetOutput, error) 
 ListPoolOriginationIdentities(ctx context.Context, params *ListPoolOriginationIdentitiesInput, optFns ...func(*Options)) (*ListPoolOriginationIdentitiesOutput, error) 
 ListRegistrationAssociations(ctx context.Context, params *ListRegistrationAssociationsInput, optFns ...func(*Options)) (*ListRegistrationAssociationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutKeyword(ctx context.Context, params *PutKeywordInput, optFns ...func(*Options)) (*PutKeywordOutput, error) 
 PutOptedOutNumber(ctx context.Context, params *PutOptedOutNumberInput, optFns ...func(*Options)) (*PutOptedOutNumberOutput, error) 
 PutRegistrationFieldValue(ctx context.Context, params *PutRegistrationFieldValueInput, optFns ...func(*Options)) (*PutRegistrationFieldValueOutput, error) 
 ReleasePhoneNumber(ctx context.Context, params *ReleasePhoneNumberInput, optFns ...func(*Options)) (*ReleasePhoneNumberOutput, error) 
 ReleaseSenderId(ctx context.Context, params *ReleaseSenderIdInput, optFns ...func(*Options)) (*ReleaseSenderIdOutput, error) 
 RequestPhoneNumber(ctx context.Context, params *RequestPhoneNumberInput, optFns ...func(*Options)) (*RequestPhoneNumberOutput, error) 
 RequestSenderId(ctx context.Context, params *RequestSenderIdInput, optFns ...func(*Options)) (*RequestSenderIdOutput, error) 
 SendDestinationNumberVerificationCode(ctx context.Context, params *SendDestinationNumberVerificationCodeInput, optFns ...func(*Options)) (*SendDestinationNumberVerificationCodeOutput, error) 
 SendMediaMessage(ctx context.Context, params *SendMediaMessageInput, optFns ...func(*Options)) (*SendMediaMessageOutput, error) 
 SendTextMessage(ctx context.Context, params *SendTextMessageInput, optFns ...func(*Options)) (*SendTextMessageOutput, error) 
 SendVoiceMessage(ctx context.Context, params *SendVoiceMessageInput, optFns ...func(*Options)) (*SendVoiceMessageOutput, error) 
 SetAccountDefaultProtectConfiguration(ctx context.Context, params *SetAccountDefaultProtectConfigurationInput, optFns ...func(*Options)) (*SetAccountDefaultProtectConfigurationOutput, error) 
 SetDefaultMessageType(ctx context.Context, params *SetDefaultMessageTypeInput, optFns ...func(*Options)) (*SetDefaultMessageTypeOutput, error) 
 SetDefaultSenderId(ctx context.Context, params *SetDefaultSenderIdInput, optFns ...func(*Options)) (*SetDefaultSenderIdOutput, error) 
 SetMediaMessageSpendLimitOverride(ctx context.Context, params *SetMediaMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*SetMediaMessageSpendLimitOverrideOutput, error) 
 SetTextMessageSpendLimitOverride(ctx context.Context, params *SetTextMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*SetTextMessageSpendLimitOverrideOutput, error) 
 SetVoiceMessageSpendLimitOverride(ctx context.Context, params *SetVoiceMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*SetVoiceMessageSpendLimitOverrideOutput, error) 
 SubmitRegistrationVersion(ctx context.Context, params *SubmitRegistrationVersionInput, optFns ...func(*Options)) (*SubmitRegistrationVersionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEventDestination(ctx context.Context, params *UpdateEventDestinationInput, optFns ...func(*Options)) (*UpdateEventDestinationOutput, error) 
 UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) 
 UpdatePool(ctx context.Context, params *UpdatePoolInput, optFns ...func(*Options)) (*UpdatePoolOutput, error) 
 UpdateProtectConfiguration(ctx context.Context, params *UpdateProtectConfigurationInput, optFns ...func(*Options)) (*UpdateProtectConfigurationOutput, error) 
 UpdateProtectConfigurationCountryRuleSet(ctx context.Context, params *UpdateProtectConfigurationCountryRuleSetInput, optFns ...func(*Options)) (*UpdateProtectConfigurationCountryRuleSetOutput, error) 
 UpdateSenderId(ctx context.Context, params *UpdateSenderIdInput, optFns ...func(*Options)) (*UpdateSenderIdOutput, error) 
 VerifyDestinationNumber(ctx context.Context, params *VerifyDestinationNumberInput, optFns ...func(*Options)) (*VerifyDestinationNumberOutput, error) 
}
