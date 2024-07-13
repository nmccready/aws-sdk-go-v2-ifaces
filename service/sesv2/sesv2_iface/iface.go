
package sesv2_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// IClient defines the interface for sesv2
type IClient interface {
 Options() Options 
 BatchGetMetricData(ctx context.Context, params *BatchGetMetricDataInput, optFns ...func(*Options)) (*BatchGetMetricDataOutput, error) 
 CancelExportJob(ctx context.Context, params *CancelExportJobInput, optFns ...func(*Options)) (*CancelExportJobOutput, error) 
 CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) 
 CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) 
 CreateContact(ctx context.Context, params *CreateContactInput, optFns ...func(*Options)) (*CreateContactOutput, error) 
 CreateContactList(ctx context.Context, params *CreateContactListInput, optFns ...func(*Options)) (*CreateContactListOutput, error) 
 CreateCustomVerificationEmailTemplate(ctx context.Context, params *CreateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*CreateCustomVerificationEmailTemplateOutput, error) 
 CreateDedicatedIpPool(ctx context.Context, params *CreateDedicatedIpPoolInput, optFns ...func(*Options)) (*CreateDedicatedIpPoolOutput, error) 
 CreateDeliverabilityTestReport(ctx context.Context, params *CreateDeliverabilityTestReportInput, optFns ...func(*Options)) (*CreateDeliverabilityTestReportOutput, error) 
 CreateEmailIdentity(ctx context.Context, params *CreateEmailIdentityInput, optFns ...func(*Options)) (*CreateEmailIdentityOutput, error) 
 CreateEmailIdentityPolicy(ctx context.Context, params *CreateEmailIdentityPolicyInput, optFns ...func(*Options)) (*CreateEmailIdentityPolicyOutput, error) 
 CreateEmailTemplate(ctx context.Context, params *CreateEmailTemplateInput, optFns ...func(*Options)) (*CreateEmailTemplateOutput, error) 
 CreateExportJob(ctx context.Context, params *CreateExportJobInput, optFns ...func(*Options)) (*CreateExportJobOutput, error) 
 CreateImportJob(ctx context.Context, params *CreateImportJobInput, optFns ...func(*Options)) (*CreateImportJobOutput, error) 
 DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) 
 DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) 
 DeleteContact(ctx context.Context, params *DeleteContactInput, optFns ...func(*Options)) (*DeleteContactOutput, error) 
 DeleteContactList(ctx context.Context, params *DeleteContactListInput, optFns ...func(*Options)) (*DeleteContactListOutput, error) 
 DeleteCustomVerificationEmailTemplate(ctx context.Context, params *DeleteCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*DeleteCustomVerificationEmailTemplateOutput, error) 
 DeleteDedicatedIpPool(ctx context.Context, params *DeleteDedicatedIpPoolInput, optFns ...func(*Options)) (*DeleteDedicatedIpPoolOutput, error) 
 DeleteEmailIdentity(ctx context.Context, params *DeleteEmailIdentityInput, optFns ...func(*Options)) (*DeleteEmailIdentityOutput, error) 
 DeleteEmailIdentityPolicy(ctx context.Context, params *DeleteEmailIdentityPolicyInput, optFns ...func(*Options)) (*DeleteEmailIdentityPolicyOutput, error) 
 DeleteEmailTemplate(ctx context.Context, params *DeleteEmailTemplateInput, optFns ...func(*Options)) (*DeleteEmailTemplateOutput, error) 
 DeleteSuppressedDestination(ctx context.Context, params *DeleteSuppressedDestinationInput, optFns ...func(*Options)) (*DeleteSuppressedDestinationOutput, error) 
 GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) 
 GetBlacklistReports(ctx context.Context, params *GetBlacklistReportsInput, optFns ...func(*Options)) (*GetBlacklistReportsOutput, error) 
 GetConfigurationSet(ctx context.Context, params *GetConfigurationSetInput, optFns ...func(*Options)) (*GetConfigurationSetOutput, error) 
 GetConfigurationSetEventDestinations(ctx context.Context, params *GetConfigurationSetEventDestinationsInput, optFns ...func(*Options)) (*GetConfigurationSetEventDestinationsOutput, error) 
 GetContact(ctx context.Context, params *GetContactInput, optFns ...func(*Options)) (*GetContactOutput, error) 
 GetContactList(ctx context.Context, params *GetContactListInput, optFns ...func(*Options)) (*GetContactListOutput, error) 
 GetCustomVerificationEmailTemplate(ctx context.Context, params *GetCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*GetCustomVerificationEmailTemplateOutput, error) 
 GetDedicatedIp(ctx context.Context, params *GetDedicatedIpInput, optFns ...func(*Options)) (*GetDedicatedIpOutput, error) 
 GetDedicatedIpPool(ctx context.Context, params *GetDedicatedIpPoolInput, optFns ...func(*Options)) (*GetDedicatedIpPoolOutput, error) 
 GetDedicatedIps(ctx context.Context, params *GetDedicatedIpsInput, optFns ...func(*Options)) (*GetDedicatedIpsOutput, error) 
 GetDeliverabilityDashboardOptions(ctx context.Context, params *GetDeliverabilityDashboardOptionsInput, optFns ...func(*Options)) (*GetDeliverabilityDashboardOptionsOutput, error) 
 GetDeliverabilityTestReport(ctx context.Context, params *GetDeliverabilityTestReportInput, optFns ...func(*Options)) (*GetDeliverabilityTestReportOutput, error) 
 GetDomainDeliverabilityCampaign(ctx context.Context, params *GetDomainDeliverabilityCampaignInput, optFns ...func(*Options)) (*GetDomainDeliverabilityCampaignOutput, error) 
 GetDomainStatisticsReport(ctx context.Context, params *GetDomainStatisticsReportInput, optFns ...func(*Options)) (*GetDomainStatisticsReportOutput, error) 
 GetEmailIdentity(ctx context.Context, params *GetEmailIdentityInput, optFns ...func(*Options)) (*GetEmailIdentityOutput, error) 
 GetEmailIdentityPolicies(ctx context.Context, params *GetEmailIdentityPoliciesInput, optFns ...func(*Options)) (*GetEmailIdentityPoliciesOutput, error) 
 GetEmailTemplate(ctx context.Context, params *GetEmailTemplateInput, optFns ...func(*Options)) (*GetEmailTemplateOutput, error) 
 GetExportJob(ctx context.Context, params *GetExportJobInput, optFns ...func(*Options)) (*GetExportJobOutput, error) 
 GetImportJob(ctx context.Context, params *GetImportJobInput, optFns ...func(*Options)) (*GetImportJobOutput, error) 
 GetMessageInsights(ctx context.Context, params *GetMessageInsightsInput, optFns ...func(*Options)) (*GetMessageInsightsOutput, error) 
 GetSuppressedDestination(ctx context.Context, params *GetSuppressedDestinationInput, optFns ...func(*Options)) (*GetSuppressedDestinationOutput, error) 
 ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) 
 ListContactLists(ctx context.Context, params *ListContactListsInput, optFns ...func(*Options)) (*ListContactListsOutput, error) 
 ListContacts(ctx context.Context, params *ListContactsInput, optFns ...func(*Options)) (*ListContactsOutput, error) 
 ListCustomVerificationEmailTemplates(ctx context.Context, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) 
 ListDedicatedIpPools(ctx context.Context, params *ListDedicatedIpPoolsInput, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) 
 ListDeliverabilityTestReports(ctx context.Context, params *ListDeliverabilityTestReportsInput, optFns ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error) 
 ListDomainDeliverabilityCampaigns(ctx context.Context, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) 
 ListEmailIdentities(ctx context.Context, params *ListEmailIdentitiesInput, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) 
 ListEmailTemplates(ctx context.Context, params *ListEmailTemplatesInput, optFns ...func(*Options)) (*ListEmailTemplatesOutput, error) 
 ListExportJobs(ctx context.Context, params *ListExportJobsInput, optFns ...func(*Options)) (*ListExportJobsOutput, error) 
 ListImportJobs(ctx context.Context, params *ListImportJobsInput, optFns ...func(*Options)) (*ListImportJobsOutput, error) 
 ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) 
 ListSuppressedDestinations(ctx context.Context, params *ListSuppressedDestinationsInput, optFns ...func(*Options)) (*ListSuppressedDestinationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAccountDedicatedIpWarmupAttributes(ctx context.Context, params *PutAccountDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutAccountDedicatedIpWarmupAttributesOutput, error) 
 PutAccountDetails(ctx context.Context, params *PutAccountDetailsInput, optFns ...func(*Options)) (*PutAccountDetailsOutput, error) 
 PutAccountSendingAttributes(ctx context.Context, params *PutAccountSendingAttributesInput, optFns ...func(*Options)) (*PutAccountSendingAttributesOutput, error) 
 PutAccountSuppressionAttributes(ctx context.Context, params *PutAccountSuppressionAttributesInput, optFns ...func(*Options)) (*PutAccountSuppressionAttributesOutput, error) 
 PutAccountVdmAttributes(ctx context.Context, params *PutAccountVdmAttributesInput, optFns ...func(*Options)) (*PutAccountVdmAttributesOutput, error) 
 PutConfigurationSetDeliveryOptions(ctx context.Context, params *PutConfigurationSetDeliveryOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetDeliveryOptionsOutput, error) 
 PutConfigurationSetReputationOptions(ctx context.Context, params *PutConfigurationSetReputationOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetReputationOptionsOutput, error) 
 PutConfigurationSetSendingOptions(ctx context.Context, params *PutConfigurationSetSendingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetSendingOptionsOutput, error) 
 PutConfigurationSetSuppressionOptions(ctx context.Context, params *PutConfigurationSetSuppressionOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetSuppressionOptionsOutput, error) 
 PutConfigurationSetTrackingOptions(ctx context.Context, params *PutConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetTrackingOptionsOutput, error) 
 PutConfigurationSetVdmOptions(ctx context.Context, params *PutConfigurationSetVdmOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetVdmOptionsOutput, error) 
 PutDedicatedIpInPool(ctx context.Context, params *PutDedicatedIpInPoolInput, optFns ...func(*Options)) (*PutDedicatedIpInPoolOutput, error) 
 PutDedicatedIpPoolScalingAttributes(ctx context.Context, params *PutDedicatedIpPoolScalingAttributesInput, optFns ...func(*Options)) (*PutDedicatedIpPoolScalingAttributesOutput, error) 
 PutDedicatedIpWarmupAttributes(ctx context.Context, params *PutDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutDedicatedIpWarmupAttributesOutput, error) 
 PutDeliverabilityDashboardOption(ctx context.Context, params *PutDeliverabilityDashboardOptionInput, optFns ...func(*Options)) (*PutDeliverabilityDashboardOptionOutput, error) 
 PutEmailIdentityConfigurationSetAttributes(ctx context.Context, params *PutEmailIdentityConfigurationSetAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityConfigurationSetAttributesOutput, error) 
 PutEmailIdentityDkimAttributes(ctx context.Context, params *PutEmailIdentityDkimAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimAttributesOutput, error) 
 PutEmailIdentityDkimSigningAttributes(ctx context.Context, params *PutEmailIdentityDkimSigningAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimSigningAttributesOutput, error) 
 PutEmailIdentityFeedbackAttributes(ctx context.Context, params *PutEmailIdentityFeedbackAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityFeedbackAttributesOutput, error) 
 PutEmailIdentityMailFromAttributes(ctx context.Context, params *PutEmailIdentityMailFromAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityMailFromAttributesOutput, error) 
 PutSuppressedDestination(ctx context.Context, params *PutSuppressedDestinationInput, optFns ...func(*Options)) (*PutSuppressedDestinationOutput, error) 
 SendBulkEmail(ctx context.Context, params *SendBulkEmailInput, optFns ...func(*Options)) (*SendBulkEmailOutput, error) 
 SendCustomVerificationEmail(ctx context.Context, params *SendCustomVerificationEmailInput, optFns ...func(*Options)) (*SendCustomVerificationEmailOutput, error) 
 SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestRenderEmailTemplate(ctx context.Context, params *TestRenderEmailTemplateInput, optFns ...func(*Options)) (*TestRenderEmailTemplateOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) 
 UpdateContact(ctx context.Context, params *UpdateContactInput, optFns ...func(*Options)) (*UpdateContactOutput, error) 
 UpdateContactList(ctx context.Context, params *UpdateContactListInput, optFns ...func(*Options)) (*UpdateContactListOutput, error) 
 UpdateCustomVerificationEmailTemplate(ctx context.Context, params *UpdateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*UpdateCustomVerificationEmailTemplateOutput, error) 
 UpdateEmailIdentityPolicy(ctx context.Context, params *UpdateEmailIdentityPolicyInput, optFns ...func(*Options)) (*UpdateEmailIdentityPolicyOutput, error) 
 UpdateEmailTemplate(ctx context.Context, params *UpdateEmailTemplateInput, optFns ...func(*Options)) (*UpdateEmailTemplateOutput, error) 
}