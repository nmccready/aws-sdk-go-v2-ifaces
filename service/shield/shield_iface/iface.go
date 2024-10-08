
package shield_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/shield"
)

// IClient defines the interface for shield
type IClient interface {
 Options() Options 
 AssociateDRTLogBucket(ctx context.Context, params *AssociateDRTLogBucketInput, optFns ...func(*Options)) (*AssociateDRTLogBucketOutput, error) 
 AssociateDRTRole(ctx context.Context, params *AssociateDRTRoleInput, optFns ...func(*Options)) (*AssociateDRTRoleOutput, error) 
 AssociateHealthCheck(ctx context.Context, params *AssociateHealthCheckInput, optFns ...func(*Options)) (*AssociateHealthCheckOutput, error) 
 AssociateProactiveEngagementDetails(ctx context.Context, params *AssociateProactiveEngagementDetailsInput, optFns ...func(*Options)) (*AssociateProactiveEngagementDetailsOutput, error) 
 CreateProtection(ctx context.Context, params *CreateProtectionInput, optFns ...func(*Options)) (*CreateProtectionOutput, error) 
 CreateProtectionGroup(ctx context.Context, params *CreateProtectionGroupInput, optFns ...func(*Options)) (*CreateProtectionGroupOutput, error) 
 CreateSubscription(ctx context.Context, params *CreateSubscriptionInput, optFns ...func(*Options)) (*CreateSubscriptionOutput, error) 
 DeleteProtection(ctx context.Context, params *DeleteProtectionInput, optFns ...func(*Options)) (*DeleteProtectionOutput, error) 
 DeleteProtectionGroup(ctx context.Context, params *DeleteProtectionGroupInput, optFns ...func(*Options)) (*DeleteProtectionGroupOutput, error) 
 DeleteSubscription(ctx context.Context, params *DeleteSubscriptionInput, optFns ...func(*Options)) (*DeleteSubscriptionOutput, error) 
 DescribeAttack(ctx context.Context, params *DescribeAttackInput, optFns ...func(*Options)) (*DescribeAttackOutput, error) 
 DescribeAttackStatistics(ctx context.Context, params *DescribeAttackStatisticsInput, optFns ...func(*Options)) (*DescribeAttackStatisticsOutput, error) 
 DescribeDRTAccess(ctx context.Context, params *DescribeDRTAccessInput, optFns ...func(*Options)) (*DescribeDRTAccessOutput, error) 
 DescribeEmergencyContactSettings(ctx context.Context, params *DescribeEmergencyContactSettingsInput, optFns ...func(*Options)) (*DescribeEmergencyContactSettingsOutput, error) 
 DescribeProtection(ctx context.Context, params *DescribeProtectionInput, optFns ...func(*Options)) (*DescribeProtectionOutput, error) 
 DescribeProtectionGroup(ctx context.Context, params *DescribeProtectionGroupInput, optFns ...func(*Options)) (*DescribeProtectionGroupOutput, error) 
 DescribeSubscription(ctx context.Context, params *DescribeSubscriptionInput, optFns ...func(*Options)) (*DescribeSubscriptionOutput, error) 
 DisableApplicationLayerAutomaticResponse(ctx context.Context, params *DisableApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*DisableApplicationLayerAutomaticResponseOutput, error) 
 DisableProactiveEngagement(ctx context.Context, params *DisableProactiveEngagementInput, optFns ...func(*Options)) (*DisableProactiveEngagementOutput, error) 
 DisassociateDRTLogBucket(ctx context.Context, params *DisassociateDRTLogBucketInput, optFns ...func(*Options)) (*DisassociateDRTLogBucketOutput, error) 
 DisassociateDRTRole(ctx context.Context, params *DisassociateDRTRoleInput, optFns ...func(*Options)) (*DisassociateDRTRoleOutput, error) 
 DisassociateHealthCheck(ctx context.Context, params *DisassociateHealthCheckInput, optFns ...func(*Options)) (*DisassociateHealthCheckOutput, error) 
 EnableApplicationLayerAutomaticResponse(ctx context.Context, params *EnableApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*EnableApplicationLayerAutomaticResponseOutput, error) 
 EnableProactiveEngagement(ctx context.Context, params *EnableProactiveEngagementInput, optFns ...func(*Options)) (*EnableProactiveEngagementOutput, error) 
 GetSubscriptionState(ctx context.Context, params *GetSubscriptionStateInput, optFns ...func(*Options)) (*GetSubscriptionStateOutput, error) 
 ListAttacks(ctx context.Context, params *ListAttacksInput, optFns ...func(*Options)) (*ListAttacksOutput, error) 
 ListProtectionGroups(ctx context.Context, params *ListProtectionGroupsInput, optFns ...func(*Options)) (*ListProtectionGroupsOutput, error) 
 ListProtections(ctx context.Context, params *ListProtectionsInput, optFns ...func(*Options)) (*ListProtectionsOutput, error) 
 ListResourcesInProtectionGroup(ctx context.Context, params *ListResourcesInProtectionGroupInput, optFns ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplicationLayerAutomaticResponse(ctx context.Context, params *UpdateApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*UpdateApplicationLayerAutomaticResponseOutput, error) 
 UpdateEmergencyContactSettings(ctx context.Context, params *UpdateEmergencyContactSettingsInput, optFns ...func(*Options)) (*UpdateEmergencyContactSettingsOutput, error) 
 UpdateProtectionGroup(ctx context.Context, params *UpdateProtectionGroupInput, optFns ...func(*Options)) (*UpdateProtectionGroupOutput, error) 
 UpdateSubscription(ctx context.Context, params *UpdateSubscriptionInput, optFns ...func(*Options)) (*UpdateSubscriptionOutput, error) 
}
