
package mailmanager

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/mailmanager"
)

// IClient defines the interface for mailmanager
type IClient interface {
 Options() Options 
 CreateAddonInstance(ctx context.Context, params *CreateAddonInstanceInput, optFns ...func(*Options)) (*CreateAddonInstanceOutput, error) 
 CreateAddonSubscription(ctx context.Context, params *CreateAddonSubscriptionInput, optFns ...func(*Options)) (*CreateAddonSubscriptionOutput, error) 
 CreateArchive(ctx context.Context, params *CreateArchiveInput, optFns ...func(*Options)) (*CreateArchiveOutput, error) 
 CreateIngressPoint(ctx context.Context, params *CreateIngressPointInput, optFns ...func(*Options)) (*CreateIngressPointOutput, error) 
 CreateRelay(ctx context.Context, params *CreateRelayInput, optFns ...func(*Options)) (*CreateRelayOutput, error) 
 CreateRuleSet(ctx context.Context, params *CreateRuleSetInput, optFns ...func(*Options)) (*CreateRuleSetOutput, error) 
 CreateTrafficPolicy(ctx context.Context, params *CreateTrafficPolicyInput, optFns ...func(*Options)) (*CreateTrafficPolicyOutput, error) 
 DeleteAddonInstance(ctx context.Context, params *DeleteAddonInstanceInput, optFns ...func(*Options)) (*DeleteAddonInstanceOutput, error) 
 DeleteAddonSubscription(ctx context.Context, params *DeleteAddonSubscriptionInput, optFns ...func(*Options)) (*DeleteAddonSubscriptionOutput, error) 
 DeleteArchive(ctx context.Context, params *DeleteArchiveInput, optFns ...func(*Options)) (*DeleteArchiveOutput, error) 
 DeleteIngressPoint(ctx context.Context, params *DeleteIngressPointInput, optFns ...func(*Options)) (*DeleteIngressPointOutput, error) 
 DeleteRelay(ctx context.Context, params *DeleteRelayInput, optFns ...func(*Options)) (*DeleteRelayOutput, error) 
 DeleteRuleSet(ctx context.Context, params *DeleteRuleSetInput, optFns ...func(*Options)) (*DeleteRuleSetOutput, error) 
 DeleteTrafficPolicy(ctx context.Context, params *DeleteTrafficPolicyInput, optFns ...func(*Options)) (*DeleteTrafficPolicyOutput, error) 
 GetAddonInstance(ctx context.Context, params *GetAddonInstanceInput, optFns ...func(*Options)) (*GetAddonInstanceOutput, error) 
 GetAddonSubscription(ctx context.Context, params *GetAddonSubscriptionInput, optFns ...func(*Options)) (*GetAddonSubscriptionOutput, error) 
 GetArchive(ctx context.Context, params *GetArchiveInput, optFns ...func(*Options)) (*GetArchiveOutput, error) 
 GetArchiveExport(ctx context.Context, params *GetArchiveExportInput, optFns ...func(*Options)) (*GetArchiveExportOutput, error) 
 GetArchiveMessage(ctx context.Context, params *GetArchiveMessageInput, optFns ...func(*Options)) (*GetArchiveMessageOutput, error) 
 GetArchiveMessageContent(ctx context.Context, params *GetArchiveMessageContentInput, optFns ...func(*Options)) (*GetArchiveMessageContentOutput, error) 
 GetArchiveSearch(ctx context.Context, params *GetArchiveSearchInput, optFns ...func(*Options)) (*GetArchiveSearchOutput, error) 
 GetArchiveSearchResults(ctx context.Context, params *GetArchiveSearchResultsInput, optFns ...func(*Options)) (*GetArchiveSearchResultsOutput, error) 
 GetIngressPoint(ctx context.Context, params *GetIngressPointInput, optFns ...func(*Options)) (*GetIngressPointOutput, error) 
 GetRelay(ctx context.Context, params *GetRelayInput, optFns ...func(*Options)) (*GetRelayOutput, error) 
 GetRuleSet(ctx context.Context, params *GetRuleSetInput, optFns ...func(*Options)) (*GetRuleSetOutput, error) 
 GetTrafficPolicy(ctx context.Context, params *GetTrafficPolicyInput, optFns ...func(*Options)) (*GetTrafficPolicyOutput, error) 
 ListAddonInstances(ctx context.Context, params *ListAddonInstancesInput, optFns ...func(*Options)) (*ListAddonInstancesOutput, error) 
 ListAddonSubscriptions(ctx context.Context, params *ListAddonSubscriptionsInput, optFns ...func(*Options)) (*ListAddonSubscriptionsOutput, error) 
 ListArchiveExports(ctx context.Context, params *ListArchiveExportsInput, optFns ...func(*Options)) (*ListArchiveExportsOutput, error) 
 ListArchiveSearches(ctx context.Context, params *ListArchiveSearchesInput, optFns ...func(*Options)) (*ListArchiveSearchesOutput, error) 
 ListArchives(ctx context.Context, params *ListArchivesInput, optFns ...func(*Options)) (*ListArchivesOutput, error) 
 ListIngressPoints(ctx context.Context, params *ListIngressPointsInput, optFns ...func(*Options)) (*ListIngressPointsOutput, error) 
 ListRelays(ctx context.Context, params *ListRelaysInput, optFns ...func(*Options)) (*ListRelaysOutput, error) 
 ListRuleSets(ctx context.Context, params *ListRuleSetsInput, optFns ...func(*Options)) (*ListRuleSetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrafficPolicies(ctx context.Context, params *ListTrafficPoliciesInput, optFns ...func(*Options)) (*ListTrafficPoliciesOutput, error) 
 StartArchiveExport(ctx context.Context, params *StartArchiveExportInput, optFns ...func(*Options)) (*StartArchiveExportOutput, error) 
 StartArchiveSearch(ctx context.Context, params *StartArchiveSearchInput, optFns ...func(*Options)) (*StartArchiveSearchOutput, error) 
 StopArchiveExport(ctx context.Context, params *StopArchiveExportInput, optFns ...func(*Options)) (*StopArchiveExportOutput, error) 
 StopArchiveSearch(ctx context.Context, params *StopArchiveSearchInput, optFns ...func(*Options)) (*StopArchiveSearchOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateArchive(ctx context.Context, params *UpdateArchiveInput, optFns ...func(*Options)) (*UpdateArchiveOutput, error) 
 UpdateIngressPoint(ctx context.Context, params *UpdateIngressPointInput, optFns ...func(*Options)) (*UpdateIngressPointOutput, error) 
 UpdateRelay(ctx context.Context, params *UpdateRelayInput, optFns ...func(*Options)) (*UpdateRelayOutput, error) 
 UpdateRuleSet(ctx context.Context, params *UpdateRuleSetInput, optFns ...func(*Options)) (*UpdateRuleSetOutput, error) 
 UpdateTrafficPolicy(ctx context.Context, params *UpdateTrafficPolicyInput, optFns ...func(*Options)) (*UpdateTrafficPolicyOutput, error) 
}
