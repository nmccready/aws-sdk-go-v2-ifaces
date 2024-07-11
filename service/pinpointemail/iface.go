
package pinpointemail

import (
    "github.com/aws/aws-sdk-go-v2/service/pinpointemail"
)

// IClient defines the interface for pinpointemail
type IClient interface {
 Options() Options 
 CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) 
 CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) 
 CreateDedicatedIpPool(ctx context.Context, params *CreateDedicatedIpPoolInput, optFns ...func(*Options)) (*CreateDedicatedIpPoolOutput, error) 
 CreateDeliverabilityTestReport(ctx context.Context, params *CreateDeliverabilityTestReportInput, optFns ...func(*Options)) (*CreateDeliverabilityTestReportOutput, error) 
 CreateEmailIdentity(ctx context.Context, params *CreateEmailIdentityInput, optFns ...func(*Options)) (*CreateEmailIdentityOutput, error) 
 DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) 
 DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) 
 DeleteDedicatedIpPool(ctx context.Context, params *DeleteDedicatedIpPoolInput, optFns ...func(*Options)) (*DeleteDedicatedIpPoolOutput, error) 
 DeleteEmailIdentity(ctx context.Context, params *DeleteEmailIdentityInput, optFns ...func(*Options)) (*DeleteEmailIdentityOutput, error) 
 GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) 
 GetBlacklistReports(ctx context.Context, params *GetBlacklistReportsInput, optFns ...func(*Options)) (*GetBlacklistReportsOutput, error) 
 GetConfigurationSet(ctx context.Context, params *GetConfigurationSetInput, optFns ...func(*Options)) (*GetConfigurationSetOutput, error) 
 GetConfigurationSetEventDestinations(ctx context.Context, params *GetConfigurationSetEventDestinationsInput, optFns ...func(*Options)) (*GetConfigurationSetEventDestinationsOutput, error) 
 GetDedicatedIp(ctx context.Context, params *GetDedicatedIpInput, optFns ...func(*Options)) (*GetDedicatedIpOutput, error) 
 GetDedicatedIps(ctx context.Context, params *GetDedicatedIpsInput, optFns ...func(*Options)) (*GetDedicatedIpsOutput, error) 
 GetDeliverabilityDashboardOptions(ctx context.Context, params *GetDeliverabilityDashboardOptionsInput, optFns ...func(*Options)) (*GetDeliverabilityDashboardOptionsOutput, error) 
 GetDeliverabilityTestReport(ctx context.Context, params *GetDeliverabilityTestReportInput, optFns ...func(*Options)) (*GetDeliverabilityTestReportOutput, error) 
 GetDomainDeliverabilityCampaign(ctx context.Context, params *GetDomainDeliverabilityCampaignInput, optFns ...func(*Options)) (*GetDomainDeliverabilityCampaignOutput, error) 
 GetDomainStatisticsReport(ctx context.Context, params *GetDomainStatisticsReportInput, optFns ...func(*Options)) (*GetDomainStatisticsReportOutput, error) 
 GetEmailIdentity(ctx context.Context, params *GetEmailIdentityInput, optFns ...func(*Options)) (*GetEmailIdentityOutput, error) 
 ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) 
 ListDedicatedIpPools(ctx context.Context, params *ListDedicatedIpPoolsInput, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) 
 ListDeliverabilityTestReports(ctx context.Context, params *ListDeliverabilityTestReportsInput, optFns ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error) 
 ListDomainDeliverabilityCampaigns(ctx context.Context, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) 
 ListEmailIdentities(ctx context.Context, params *ListEmailIdentitiesInput, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAccountDedicatedIpWarmupAttributes(ctx context.Context, params *PutAccountDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutAccountDedicatedIpWarmupAttributesOutput, error) 
 PutAccountSendingAttributes(ctx context.Context, params *PutAccountSendingAttributesInput, optFns ...func(*Options)) (*PutAccountSendingAttributesOutput, error) 
 PutConfigurationSetDeliveryOptions(ctx context.Context, params *PutConfigurationSetDeliveryOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetDeliveryOptionsOutput, error) 
 PutConfigurationSetReputationOptions(ctx context.Context, params *PutConfigurationSetReputationOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetReputationOptionsOutput, error) 
 PutConfigurationSetSendingOptions(ctx context.Context, params *PutConfigurationSetSendingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetSendingOptionsOutput, error) 
 PutConfigurationSetTrackingOptions(ctx context.Context, params *PutConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetTrackingOptionsOutput, error) 
 PutDedicatedIpInPool(ctx context.Context, params *PutDedicatedIpInPoolInput, optFns ...func(*Options)) (*PutDedicatedIpInPoolOutput, error) 
 PutDedicatedIpWarmupAttributes(ctx context.Context, params *PutDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutDedicatedIpWarmupAttributesOutput, error) 
 PutDeliverabilityDashboardOption(ctx context.Context, params *PutDeliverabilityDashboardOptionInput, optFns ...func(*Options)) (*PutDeliverabilityDashboardOptionOutput, error) 
 PutEmailIdentityDkimAttributes(ctx context.Context, params *PutEmailIdentityDkimAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimAttributesOutput, error) 
 PutEmailIdentityFeedbackAttributes(ctx context.Context, params *PutEmailIdentityFeedbackAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityFeedbackAttributesOutput, error) 
 PutEmailIdentityMailFromAttributes(ctx context.Context, params *PutEmailIdentityMailFromAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityMailFromAttributesOutput, error) 
 SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) 
}
