
package ses

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/ses"
)

// IClient defines the interface for ses
type IClient interface {
 Options() Options 
 CloneReceiptRuleSet(ctx context.Context, params *CloneReceiptRuleSetInput, optFns ...func(*Options)) (*CloneReceiptRuleSetOutput, error) 
 CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) 
 CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) 
 CreateConfigurationSetTrackingOptions(ctx context.Context, params *CreateConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*CreateConfigurationSetTrackingOptionsOutput, error) 
 CreateCustomVerificationEmailTemplate(ctx context.Context, params *CreateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*CreateCustomVerificationEmailTemplateOutput, error) 
 CreateReceiptFilter(ctx context.Context, params *CreateReceiptFilterInput, optFns ...func(*Options)) (*CreateReceiptFilterOutput, error) 
 CreateReceiptRule(ctx context.Context, params *CreateReceiptRuleInput, optFns ...func(*Options)) (*CreateReceiptRuleOutput, error) 
 CreateReceiptRuleSet(ctx context.Context, params *CreateReceiptRuleSetInput, optFns ...func(*Options)) (*CreateReceiptRuleSetOutput, error) 
 CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) 
 DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) 
 DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) 
 DeleteConfigurationSetTrackingOptions(ctx context.Context, params *DeleteConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*DeleteConfigurationSetTrackingOptionsOutput, error) 
 DeleteCustomVerificationEmailTemplate(ctx context.Context, params *DeleteCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*DeleteCustomVerificationEmailTemplateOutput, error) 
 DeleteIdentity(ctx context.Context, params *DeleteIdentityInput, optFns ...func(*Options)) (*DeleteIdentityOutput, error) 
 DeleteIdentityPolicy(ctx context.Context, params *DeleteIdentityPolicyInput, optFns ...func(*Options)) (*DeleteIdentityPolicyOutput, error) 
 DeleteReceiptFilter(ctx context.Context, params *DeleteReceiptFilterInput, optFns ...func(*Options)) (*DeleteReceiptFilterOutput, error) 
 DeleteReceiptRule(ctx context.Context, params *DeleteReceiptRuleInput, optFns ...func(*Options)) (*DeleteReceiptRuleOutput, error) 
 DeleteReceiptRuleSet(ctx context.Context, params *DeleteReceiptRuleSetInput, optFns ...func(*Options)) (*DeleteReceiptRuleSetOutput, error) 
 DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) 
 DeleteVerifiedEmailAddress(ctx context.Context, params *DeleteVerifiedEmailAddressInput, optFns ...func(*Options)) (*DeleteVerifiedEmailAddressOutput, error) 
 DescribeActiveReceiptRuleSet(ctx context.Context, params *DescribeActiveReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeActiveReceiptRuleSetOutput, error) 
 DescribeConfigurationSet(ctx context.Context, params *DescribeConfigurationSetInput, optFns ...func(*Options)) (*DescribeConfigurationSetOutput, error) 
 DescribeReceiptRule(ctx context.Context, params *DescribeReceiptRuleInput, optFns ...func(*Options)) (*DescribeReceiptRuleOutput, error) 
 DescribeReceiptRuleSet(ctx context.Context, params *DescribeReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeReceiptRuleSetOutput, error) 
 GetAccountSendingEnabled(ctx context.Context, params *GetAccountSendingEnabledInput, optFns ...func(*Options)) (*GetAccountSendingEnabledOutput, error) 
 GetCustomVerificationEmailTemplate(ctx context.Context, params *GetCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*GetCustomVerificationEmailTemplateOutput, error) 
 GetIdentityDkimAttributes(ctx context.Context, params *GetIdentityDkimAttributesInput, optFns ...func(*Options)) (*GetIdentityDkimAttributesOutput, error) 
 GetIdentityMailFromDomainAttributes(ctx context.Context, params *GetIdentityMailFromDomainAttributesInput, optFns ...func(*Options)) (*GetIdentityMailFromDomainAttributesOutput, error) 
 GetIdentityNotificationAttributes(ctx context.Context, params *GetIdentityNotificationAttributesInput, optFns ...func(*Options)) (*GetIdentityNotificationAttributesOutput, error) 
 GetIdentityPolicies(ctx context.Context, params *GetIdentityPoliciesInput, optFns ...func(*Options)) (*GetIdentityPoliciesOutput, error) 
 GetIdentityVerificationAttributes(ctx context.Context, params *GetIdentityVerificationAttributesInput, optFns ...func(*Options)) (*GetIdentityVerificationAttributesOutput, error) 
 GetSendQuota(ctx context.Context, params *GetSendQuotaInput, optFns ...func(*Options)) (*GetSendQuotaOutput, error) 
 GetSendStatistics(ctx context.Context, params *GetSendStatisticsInput, optFns ...func(*Options)) (*GetSendStatisticsOutput, error) 
 GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) 
 ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) 
 ListCustomVerificationEmailTemplates(ctx context.Context, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) 
 ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) 
 ListIdentityPolicies(ctx context.Context, params *ListIdentityPoliciesInput, optFns ...func(*Options)) (*ListIdentityPoliciesOutput, error) 
 ListReceiptFilters(ctx context.Context, params *ListReceiptFiltersInput, optFns ...func(*Options)) (*ListReceiptFiltersOutput, error) 
 ListReceiptRuleSets(ctx context.Context, params *ListReceiptRuleSetsInput, optFns ...func(*Options)) (*ListReceiptRuleSetsOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 ListVerifiedEmailAddresses(ctx context.Context, params *ListVerifiedEmailAddressesInput, optFns ...func(*Options)) (*ListVerifiedEmailAddressesOutput, error) 
 PutConfigurationSetDeliveryOptions(ctx context.Context, params *PutConfigurationSetDeliveryOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetDeliveryOptionsOutput, error) 
 PutIdentityPolicy(ctx context.Context, params *PutIdentityPolicyInput, optFns ...func(*Options)) (*PutIdentityPolicyOutput, error) 
 ReorderReceiptRuleSet(ctx context.Context, params *ReorderReceiptRuleSetInput, optFns ...func(*Options)) (*ReorderReceiptRuleSetOutput, error) 
 SendBounce(ctx context.Context, params *SendBounceInput, optFns ...func(*Options)) (*SendBounceOutput, error) 
 SendBulkTemplatedEmail(ctx context.Context, params *SendBulkTemplatedEmailInput, optFns ...func(*Options)) (*SendBulkTemplatedEmailOutput, error) 
 SendCustomVerificationEmail(ctx context.Context, params *SendCustomVerificationEmailInput, optFns ...func(*Options)) (*SendCustomVerificationEmailOutput, error) 
 SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) 
 SendRawEmail(ctx context.Context, params *SendRawEmailInput, optFns ...func(*Options)) (*SendRawEmailOutput, error) 
 SendTemplatedEmail(ctx context.Context, params *SendTemplatedEmailInput, optFns ...func(*Options)) (*SendTemplatedEmailOutput, error) 
 SetActiveReceiptRuleSet(ctx context.Context, params *SetActiveReceiptRuleSetInput, optFns ...func(*Options)) (*SetActiveReceiptRuleSetOutput, error) 
 SetIdentityDkimEnabled(ctx context.Context, params *SetIdentityDkimEnabledInput, optFns ...func(*Options)) (*SetIdentityDkimEnabledOutput, error) 
 SetIdentityFeedbackForwardingEnabled(ctx context.Context, params *SetIdentityFeedbackForwardingEnabledInput, optFns ...func(*Options)) (*SetIdentityFeedbackForwardingEnabledOutput, error) 
 SetIdentityHeadersInNotificationsEnabled(ctx context.Context, params *SetIdentityHeadersInNotificationsEnabledInput, optFns ...func(*Options)) (*SetIdentityHeadersInNotificationsEnabledOutput, error) 
 SetIdentityMailFromDomain(ctx context.Context, params *SetIdentityMailFromDomainInput, optFns ...func(*Options)) (*SetIdentityMailFromDomainOutput, error) 
 SetIdentityNotificationTopic(ctx context.Context, params *SetIdentityNotificationTopicInput, optFns ...func(*Options)) (*SetIdentityNotificationTopicOutput, error) 
 SetReceiptRulePosition(ctx context.Context, params *SetReceiptRulePositionInput, optFns ...func(*Options)) (*SetReceiptRulePositionOutput, error) 
 TestRenderTemplate(ctx context.Context, params *TestRenderTemplateInput, optFns ...func(*Options)) (*TestRenderTemplateOutput, error) 
 UpdateAccountSendingEnabled(ctx context.Context, params *UpdateAccountSendingEnabledInput, optFns ...func(*Options)) (*UpdateAccountSendingEnabledOutput, error) 
 UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) 
 UpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, params *UpdateConfigurationSetReputationMetricsEnabledInput, optFns ...func(*Options)) (*UpdateConfigurationSetReputationMetricsEnabledOutput, error) 
 UpdateConfigurationSetSendingEnabled(ctx context.Context, params *UpdateConfigurationSetSendingEnabledInput, optFns ...func(*Options)) (*UpdateConfigurationSetSendingEnabledOutput, error) 
 UpdateConfigurationSetTrackingOptions(ctx context.Context, params *UpdateConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*UpdateConfigurationSetTrackingOptionsOutput, error) 
 UpdateCustomVerificationEmailTemplate(ctx context.Context, params *UpdateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*UpdateCustomVerificationEmailTemplateOutput, error) 
 UpdateReceiptRule(ctx context.Context, params *UpdateReceiptRuleInput, optFns ...func(*Options)) (*UpdateReceiptRuleOutput, error) 
 UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) 
 VerifyDomainDkim(ctx context.Context, params *VerifyDomainDkimInput, optFns ...func(*Options)) (*VerifyDomainDkimOutput, error) 
 VerifyDomainIdentity(ctx context.Context, params *VerifyDomainIdentityInput, optFns ...func(*Options)) (*VerifyDomainIdentityOutput, error) 
 VerifyEmailAddress(ctx context.Context, params *VerifyEmailAddressInput, optFns ...func(*Options)) (*VerifyEmailAddressOutput, error) 
 VerifyEmailIdentity(ctx context.Context, params *VerifyEmailIdentityInput, optFns ...func(*Options)) (*VerifyEmailIdentityOutput, error) 
}
