
package ssmcontacts

import (
    "github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
)

// IClient defines the interface for ssmcontacts
type IClient interface {
 Options() Options 
 AcceptPage(ctx context.Context, params *AcceptPageInput, optFns ...func(*Options)) (*AcceptPageOutput, error) 
 ActivateContactChannel(ctx context.Context, params *ActivateContactChannelInput, optFns ...func(*Options)) (*ActivateContactChannelOutput, error) 
 CreateContact(ctx context.Context, params *CreateContactInput, optFns ...func(*Options)) (*CreateContactOutput, error) 
 CreateContactChannel(ctx context.Context, params *CreateContactChannelInput, optFns ...func(*Options)) (*CreateContactChannelOutput, error) 
 CreateRotation(ctx context.Context, params *CreateRotationInput, optFns ...func(*Options)) (*CreateRotationOutput, error) 
 CreateRotationOverride(ctx context.Context, params *CreateRotationOverrideInput, optFns ...func(*Options)) (*CreateRotationOverrideOutput, error) 
 DeactivateContactChannel(ctx context.Context, params *DeactivateContactChannelInput, optFns ...func(*Options)) (*DeactivateContactChannelOutput, error) 
 DeleteContact(ctx context.Context, params *DeleteContactInput, optFns ...func(*Options)) (*DeleteContactOutput, error) 
 DeleteContactChannel(ctx context.Context, params *DeleteContactChannelInput, optFns ...func(*Options)) (*DeleteContactChannelOutput, error) 
 DeleteRotation(ctx context.Context, params *DeleteRotationInput, optFns ...func(*Options)) (*DeleteRotationOutput, error) 
 DeleteRotationOverride(ctx context.Context, params *DeleteRotationOverrideInput, optFns ...func(*Options)) (*DeleteRotationOverrideOutput, error) 
 DescribeEngagement(ctx context.Context, params *DescribeEngagementInput, optFns ...func(*Options)) (*DescribeEngagementOutput, error) 
 DescribePage(ctx context.Context, params *DescribePageInput, optFns ...func(*Options)) (*DescribePageOutput, error) 
 GetContact(ctx context.Context, params *GetContactInput, optFns ...func(*Options)) (*GetContactOutput, error) 
 GetContactChannel(ctx context.Context, params *GetContactChannelInput, optFns ...func(*Options)) (*GetContactChannelOutput, error) 
 GetContactPolicy(ctx context.Context, params *GetContactPolicyInput, optFns ...func(*Options)) (*GetContactPolicyOutput, error) 
 GetRotation(ctx context.Context, params *GetRotationInput, optFns ...func(*Options)) (*GetRotationOutput, error) 
 GetRotationOverride(ctx context.Context, params *GetRotationOverrideInput, optFns ...func(*Options)) (*GetRotationOverrideOutput, error) 
 ListContactChannels(ctx context.Context, params *ListContactChannelsInput, optFns ...func(*Options)) (*ListContactChannelsOutput, error) 
 ListContacts(ctx context.Context, params *ListContactsInput, optFns ...func(*Options)) (*ListContactsOutput, error) 
 ListEngagements(ctx context.Context, params *ListEngagementsInput, optFns ...func(*Options)) (*ListEngagementsOutput, error) 
 ListPageReceipts(ctx context.Context, params *ListPageReceiptsInput, optFns ...func(*Options)) (*ListPageReceiptsOutput, error) 
 ListPageResolutions(ctx context.Context, params *ListPageResolutionsInput, optFns ...func(*Options)) (*ListPageResolutionsOutput, error) 
 ListPagesByContact(ctx context.Context, params *ListPagesByContactInput, optFns ...func(*Options)) (*ListPagesByContactOutput, error) 
 ListPagesByEngagement(ctx context.Context, params *ListPagesByEngagementInput, optFns ...func(*Options)) (*ListPagesByEngagementOutput, error) 
 ListPreviewRotationShifts(ctx context.Context, params *ListPreviewRotationShiftsInput, optFns ...func(*Options)) (*ListPreviewRotationShiftsOutput, error) 
 ListRotationOverrides(ctx context.Context, params *ListRotationOverridesInput, optFns ...func(*Options)) (*ListRotationOverridesOutput, error) 
 ListRotationShifts(ctx context.Context, params *ListRotationShiftsInput, optFns ...func(*Options)) (*ListRotationShiftsOutput, error) 
 ListRotations(ctx context.Context, params *ListRotationsInput, optFns ...func(*Options)) (*ListRotationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutContactPolicy(ctx context.Context, params *PutContactPolicyInput, optFns ...func(*Options)) (*PutContactPolicyOutput, error) 
 SendActivationCode(ctx context.Context, params *SendActivationCodeInput, optFns ...func(*Options)) (*SendActivationCodeOutput, error) 
 StartEngagement(ctx context.Context, params *StartEngagementInput, optFns ...func(*Options)) (*StartEngagementOutput, error) 
 StopEngagement(ctx context.Context, params *StopEngagementInput, optFns ...func(*Options)) (*StopEngagementOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateContact(ctx context.Context, params *UpdateContactInput, optFns ...func(*Options)) (*UpdateContactOutput, error) 
 UpdateContactChannel(ctx context.Context, params *UpdateContactChannelInput, optFns ...func(*Options)) (*UpdateContactChannelOutput, error) 
 UpdateRotation(ctx context.Context, params *UpdateRotationInput, optFns ...func(*Options)) (*UpdateRotationOutput, error) 
}
