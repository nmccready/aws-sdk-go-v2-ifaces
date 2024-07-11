
package wafv2

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/wafv2"
)

// IClient defines the interface for wafv2
type IClient interface {
 Options() Options 
 AssociateWebACL(ctx context.Context, params *AssociateWebACLInput, optFns ...func(*Options)) (*AssociateWebACLOutput, error) 
 CheckCapacity(ctx context.Context, params *CheckCapacityInput, optFns ...func(*Options)) (*CheckCapacityOutput, error) 
 CreateAPIKey(ctx context.Context, params *CreateAPIKeyInput, optFns ...func(*Options)) (*CreateAPIKeyOutput, error) 
 CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) 
 CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) 
 CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) 
 CreateWebACL(ctx context.Context, params *CreateWebACLInput, optFns ...func(*Options)) (*CreateWebACLOutput, error) 
 DeleteAPIKey(ctx context.Context, params *DeleteAPIKeyInput, optFns ...func(*Options)) (*DeleteAPIKeyOutput, error) 
 DeleteFirewallManagerRuleGroups(ctx context.Context, params *DeleteFirewallManagerRuleGroupsInput, optFns ...func(*Options)) (*DeleteFirewallManagerRuleGroupsOutput, error) 
 DeleteIPSet(ctx context.Context, params *DeleteIPSetInput, optFns ...func(*Options)) (*DeleteIPSetOutput, error) 
 DeleteLoggingConfiguration(ctx context.Context, params *DeleteLoggingConfigurationInput, optFns ...func(*Options)) (*DeleteLoggingConfigurationOutput, error) 
 DeletePermissionPolicy(ctx context.Context, params *DeletePermissionPolicyInput, optFns ...func(*Options)) (*DeletePermissionPolicyOutput, error) 
 DeleteRegexPatternSet(ctx context.Context, params *DeleteRegexPatternSetInput, optFns ...func(*Options)) (*DeleteRegexPatternSetOutput, error) 
 DeleteRuleGroup(ctx context.Context, params *DeleteRuleGroupInput, optFns ...func(*Options)) (*DeleteRuleGroupOutput, error) 
 DeleteWebACL(ctx context.Context, params *DeleteWebACLInput, optFns ...func(*Options)) (*DeleteWebACLOutput, error) 
 DescribeAllManagedProducts(ctx context.Context, params *DescribeAllManagedProductsInput, optFns ...func(*Options)) (*DescribeAllManagedProductsOutput, error) 
 DescribeManagedProductsByVendor(ctx context.Context, params *DescribeManagedProductsByVendorInput, optFns ...func(*Options)) (*DescribeManagedProductsByVendorOutput, error) 
 DescribeManagedRuleGroup(ctx context.Context, params *DescribeManagedRuleGroupInput, optFns ...func(*Options)) (*DescribeManagedRuleGroupOutput, error) 
 DisassociateWebACL(ctx context.Context, params *DisassociateWebACLInput, optFns ...func(*Options)) (*DisassociateWebACLOutput, error) 
 GenerateMobileSdkReleaseUrl(ctx context.Context, params *GenerateMobileSdkReleaseUrlInput, optFns ...func(*Options)) (*GenerateMobileSdkReleaseUrlOutput, error) 
 GetDecryptedAPIKey(ctx context.Context, params *GetDecryptedAPIKeyInput, optFns ...func(*Options)) (*GetDecryptedAPIKeyOutput, error) 
 GetIPSet(ctx context.Context, params *GetIPSetInput, optFns ...func(*Options)) (*GetIPSetOutput, error) 
 GetLoggingConfiguration(ctx context.Context, params *GetLoggingConfigurationInput, optFns ...func(*Options)) (*GetLoggingConfigurationOutput, error) 
 GetManagedRuleSet(ctx context.Context, params *GetManagedRuleSetInput, optFns ...func(*Options)) (*GetManagedRuleSetOutput, error) 
 GetMobileSdkRelease(ctx context.Context, params *GetMobileSdkReleaseInput, optFns ...func(*Options)) (*GetMobileSdkReleaseOutput, error) 
 GetPermissionPolicy(ctx context.Context, params *GetPermissionPolicyInput, optFns ...func(*Options)) (*GetPermissionPolicyOutput, error) 
 GetRateBasedStatementManagedKeys(ctx context.Context, params *GetRateBasedStatementManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedStatementManagedKeysOutput, error) 
 GetRegexPatternSet(ctx context.Context, params *GetRegexPatternSetInput, optFns ...func(*Options)) (*GetRegexPatternSetOutput, error) 
 GetRuleGroup(ctx context.Context, params *GetRuleGroupInput, optFns ...func(*Options)) (*GetRuleGroupOutput, error) 
 GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) 
 GetWebACL(ctx context.Context, params *GetWebACLInput, optFns ...func(*Options)) (*GetWebACLOutput, error) 
 GetWebACLForResource(ctx context.Context, params *GetWebACLForResourceInput, optFns ...func(*Options)) (*GetWebACLForResourceOutput, error) 
 ListAPIKeys(ctx context.Context, params *ListAPIKeysInput, optFns ...func(*Options)) (*ListAPIKeysOutput, error) 
 ListAvailableManagedRuleGroupVersions(ctx context.Context, params *ListAvailableManagedRuleGroupVersionsInput, optFns ...func(*Options)) (*ListAvailableManagedRuleGroupVersionsOutput, error) 
 ListAvailableManagedRuleGroups(ctx context.Context, params *ListAvailableManagedRuleGroupsInput, optFns ...func(*Options)) (*ListAvailableManagedRuleGroupsOutput, error) 
 ListIPSets(ctx context.Context, params *ListIPSetsInput, optFns ...func(*Options)) (*ListIPSetsOutput, error) 
 ListLoggingConfigurations(ctx context.Context, params *ListLoggingConfigurationsInput, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) 
 ListManagedRuleSets(ctx context.Context, params *ListManagedRuleSetsInput, optFns ...func(*Options)) (*ListManagedRuleSetsOutput, error) 
 ListMobileSdkReleases(ctx context.Context, params *ListMobileSdkReleasesInput, optFns ...func(*Options)) (*ListMobileSdkReleasesOutput, error) 
 ListRegexPatternSets(ctx context.Context, params *ListRegexPatternSetsInput, optFns ...func(*Options)) (*ListRegexPatternSetsOutput, error) 
 ListResourcesForWebACL(ctx context.Context, params *ListResourcesForWebACLInput, optFns ...func(*Options)) (*ListResourcesForWebACLOutput, error) 
 ListRuleGroups(ctx context.Context, params *ListRuleGroupsInput, optFns ...func(*Options)) (*ListRuleGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebACLs(ctx context.Context, params *ListWebACLsInput, optFns ...func(*Options)) (*ListWebACLsOutput, error) 
 PutLoggingConfiguration(ctx context.Context, params *PutLoggingConfigurationInput, optFns ...func(*Options)) (*PutLoggingConfigurationOutput, error) 
 PutManagedRuleSetVersions(ctx context.Context, params *PutManagedRuleSetVersionsInput, optFns ...func(*Options)) (*PutManagedRuleSetVersionsOutput, error) 
 PutPermissionPolicy(ctx context.Context, params *PutPermissionPolicyInput, optFns ...func(*Options)) (*PutPermissionPolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateIPSet(ctx context.Context, params *UpdateIPSetInput, optFns ...func(*Options)) (*UpdateIPSetOutput, error) 
 UpdateManagedRuleSetVersionExpiryDate(ctx context.Context, params *UpdateManagedRuleSetVersionExpiryDateInput, optFns ...func(*Options)) (*UpdateManagedRuleSetVersionExpiryDateOutput, error) 
 UpdateRegexPatternSet(ctx context.Context, params *UpdateRegexPatternSetInput, optFns ...func(*Options)) (*UpdateRegexPatternSetOutput, error) 
 UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) 
 UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) 
}
