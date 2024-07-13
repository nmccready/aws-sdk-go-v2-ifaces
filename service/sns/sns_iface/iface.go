
package sns_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/sns"
)

// IClient defines the interface for sns
type IClient interface {
 Options() Options 
 AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) 
 CheckIfPhoneNumberIsOptedOut(ctx context.Context, params *CheckIfPhoneNumberIsOptedOutInput, optFns ...func(*Options)) (*CheckIfPhoneNumberIsOptedOutOutput, error) 
 ConfirmSubscription(ctx context.Context, params *ConfirmSubscriptionInput, optFns ...func(*Options)) (*ConfirmSubscriptionOutput, error) 
 CreatePlatformApplication(ctx context.Context, params *CreatePlatformApplicationInput, optFns ...func(*Options)) (*CreatePlatformApplicationOutput, error) 
 CreatePlatformEndpoint(ctx context.Context, params *CreatePlatformEndpointInput, optFns ...func(*Options)) (*CreatePlatformEndpointOutput, error) 
 CreateSMSSandboxPhoneNumber(ctx context.Context, params *CreateSMSSandboxPhoneNumberInput, optFns ...func(*Options)) (*CreateSMSSandboxPhoneNumberOutput, error) 
 CreateTopic(ctx context.Context, params *CreateTopicInput, optFns ...func(*Options)) (*CreateTopicOutput, error) 
 DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) 
 DeletePlatformApplication(ctx context.Context, params *DeletePlatformApplicationInput, optFns ...func(*Options)) (*DeletePlatformApplicationOutput, error) 
 DeleteSMSSandboxPhoneNumber(ctx context.Context, params *DeleteSMSSandboxPhoneNumberInput, optFns ...func(*Options)) (*DeleteSMSSandboxPhoneNumberOutput, error) 
 DeleteTopic(ctx context.Context, params *DeleteTopicInput, optFns ...func(*Options)) (*DeleteTopicOutput, error) 
 GetDataProtectionPolicy(ctx context.Context, params *GetDataProtectionPolicyInput, optFns ...func(*Options)) (*GetDataProtectionPolicyOutput, error) 
 GetEndpointAttributes(ctx context.Context, params *GetEndpointAttributesInput, optFns ...func(*Options)) (*GetEndpointAttributesOutput, error) 
 GetPlatformApplicationAttributes(ctx context.Context, params *GetPlatformApplicationAttributesInput, optFns ...func(*Options)) (*GetPlatformApplicationAttributesOutput, error) 
 GetSMSAttributes(ctx context.Context, params *GetSMSAttributesInput, optFns ...func(*Options)) (*GetSMSAttributesOutput, error) 
 GetSMSSandboxAccountStatus(ctx context.Context, params *GetSMSSandboxAccountStatusInput, optFns ...func(*Options)) (*GetSMSSandboxAccountStatusOutput, error) 
 GetSubscriptionAttributes(ctx context.Context, params *GetSubscriptionAttributesInput, optFns ...func(*Options)) (*GetSubscriptionAttributesOutput, error) 
 GetTopicAttributes(ctx context.Context, params *GetTopicAttributesInput, optFns ...func(*Options)) (*GetTopicAttributesOutput, error) 
 ListEndpointsByPlatformApplication(ctx context.Context, params *ListEndpointsByPlatformApplicationInput, optFns ...func(*Options)) (*ListEndpointsByPlatformApplicationOutput, error) 
 ListOriginationNumbers(ctx context.Context, params *ListOriginationNumbersInput, optFns ...func(*Options)) (*ListOriginationNumbersOutput, error) 
 ListPhoneNumbersOptedOut(ctx context.Context, params *ListPhoneNumbersOptedOutInput, optFns ...func(*Options)) (*ListPhoneNumbersOptedOutOutput, error) 
 ListPlatformApplications(ctx context.Context, params *ListPlatformApplicationsInput, optFns ...func(*Options)) (*ListPlatformApplicationsOutput, error) 
 ListSMSSandboxPhoneNumbers(ctx context.Context, params *ListSMSSandboxPhoneNumbersInput, optFns ...func(*Options)) (*ListSMSSandboxPhoneNumbersOutput, error) 
 ListSubscriptions(ctx context.Context, params *ListSubscriptionsInput, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) 
 ListSubscriptionsByTopic(ctx context.Context, params *ListSubscriptionsByTopicInput, optFns ...func(*Options)) (*ListSubscriptionsByTopicOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTopics(ctx context.Context, params *ListTopicsInput, optFns ...func(*Options)) (*ListTopicsOutput, error) 
 OptInPhoneNumber(ctx context.Context, params *OptInPhoneNumberInput, optFns ...func(*Options)) (*OptInPhoneNumberOutput, error) 
 Publish(ctx context.Context, params *PublishInput, optFns ...func(*Options)) (*PublishOutput, error) 
 PublishBatch(ctx context.Context, params *PublishBatchInput, optFns ...func(*Options)) (*PublishBatchOutput, error) 
 PutDataProtectionPolicy(ctx context.Context, params *PutDataProtectionPolicyInput, optFns ...func(*Options)) (*PutDataProtectionPolicyOutput, error) 
 RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) 
 SetEndpointAttributes(ctx context.Context, params *SetEndpointAttributesInput, optFns ...func(*Options)) (*SetEndpointAttributesOutput, error) 
 SetPlatformApplicationAttributes(ctx context.Context, params *SetPlatformApplicationAttributesInput, optFns ...func(*Options)) (*SetPlatformApplicationAttributesOutput, error) 
 SetSMSAttributes(ctx context.Context, params *SetSMSAttributesInput, optFns ...func(*Options)) (*SetSMSAttributesOutput, error) 
 SetSubscriptionAttributes(ctx context.Context, params *SetSubscriptionAttributesInput, optFns ...func(*Options)) (*SetSubscriptionAttributesOutput, error) 
 SetTopicAttributes(ctx context.Context, params *SetTopicAttributesInput, optFns ...func(*Options)) (*SetTopicAttributesOutput, error) 
 Subscribe(ctx context.Context, params *SubscribeInput, optFns ...func(*Options)) (*SubscribeOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 Unsubscribe(ctx context.Context, params *UnsubscribeInput, optFns ...func(*Options)) (*UnsubscribeOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 VerifySMSSandboxPhoneNumber(ctx context.Context, params *VerifySMSSandboxPhoneNumberInput, optFns ...func(*Options)) (*VerifySMSSandboxPhoneNumberOutput, error) 
}