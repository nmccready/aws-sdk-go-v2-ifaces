
package wafregional

import (
    "github.com/aws/aws-sdk-go-v2/service/wafregional"
)

// IWafregional defines the interface for wafregional
type IWafregional interface {
 Options() Options 
 AssociateWebACL(ctx context.Context, params *AssociateWebACLInput, optFns ...func(*Options)) (*AssociateWebACLOutput, error) 
 CreateByteMatchSet(ctx context.Context, params *CreateByteMatchSetInput, optFns ...func(*Options)) (*CreateByteMatchSetOutput, error) 
 CreateGeoMatchSet(ctx context.Context, params *CreateGeoMatchSetInput, optFns ...func(*Options)) (*CreateGeoMatchSetOutput, error) 
 CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) 
 CreateRateBasedRule(ctx context.Context, params *CreateRateBasedRuleInput, optFns ...func(*Options)) (*CreateRateBasedRuleOutput, error) 
 CreateRegexMatchSet(ctx context.Context, params *CreateRegexMatchSetInput, optFns ...func(*Options)) (*CreateRegexMatchSetOutput, error) 
 CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) 
 CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) 
 CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) 
 CreateSizeConstraintSet(ctx context.Context, params *CreateSizeConstraintSetInput, optFns ...func(*Options)) (*CreateSizeConstraintSetOutput, error) 
 CreateSqlInjectionMatchSet(ctx context.Context, params *CreateSqlInjectionMatchSetInput, optFns ...func(*Options)) (*CreateSqlInjectionMatchSetOutput, error) 
 CreateWebACL(ctx context.Context, params *CreateWebACLInput, optFns ...func(*Options)) (*CreateWebACLOutput, error) 
 CreateWebACLMigrationStack(ctx context.Context, params *CreateWebACLMigrationStackInput, optFns ...func(*Options)) (*CreateWebACLMigrationStackOutput, error) 
 CreateXssMatchSet(ctx context.Context, params *CreateXssMatchSetInput, optFns ...func(*Options)) (*CreateXssMatchSetOutput, error) 
 DeleteByteMatchSet(ctx context.Context, params *DeleteByteMatchSetInput, optFns ...func(*Options)) (*DeleteByteMatchSetOutput, error) 
 DeleteGeoMatchSet(ctx context.Context, params *DeleteGeoMatchSetInput, optFns ...func(*Options)) (*DeleteGeoMatchSetOutput, error) 
 DeleteIPSet(ctx context.Context, params *DeleteIPSetInput, optFns ...func(*Options)) (*DeleteIPSetOutput, error) 
 DeleteLoggingConfiguration(ctx context.Context, params *DeleteLoggingConfigurationInput, optFns ...func(*Options)) (*DeleteLoggingConfigurationOutput, error) 
 DeletePermissionPolicy(ctx context.Context, params *DeletePermissionPolicyInput, optFns ...func(*Options)) (*DeletePermissionPolicyOutput, error) 
 DeleteRateBasedRule(ctx context.Context, params *DeleteRateBasedRuleInput, optFns ...func(*Options)) (*DeleteRateBasedRuleOutput, error) 
 DeleteRegexMatchSet(ctx context.Context, params *DeleteRegexMatchSetInput, optFns ...func(*Options)) (*DeleteRegexMatchSetOutput, error) 
 DeleteRegexPatternSet(ctx context.Context, params *DeleteRegexPatternSetInput, optFns ...func(*Options)) (*DeleteRegexPatternSetOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 DeleteRuleGroup(ctx context.Context, params *DeleteRuleGroupInput, optFns ...func(*Options)) (*DeleteRuleGroupOutput, error) 
 DeleteSizeConstraintSet(ctx context.Context, params *DeleteSizeConstraintSetInput, optFns ...func(*Options)) (*DeleteSizeConstraintSetOutput, error) 
 DeleteSqlInjectionMatchSet(ctx context.Context, params *DeleteSqlInjectionMatchSetInput, optFns ...func(*Options)) (*DeleteSqlInjectionMatchSetOutput, error) 
 DeleteWebACL(ctx context.Context, params *DeleteWebACLInput, optFns ...func(*Options)) (*DeleteWebACLOutput, error) 
 DeleteXssMatchSet(ctx context.Context, params *DeleteXssMatchSetInput, optFns ...func(*Options)) (*DeleteXssMatchSetOutput, error) 
 DisassociateWebACL(ctx context.Context, params *DisassociateWebACLInput, optFns ...func(*Options)) (*DisassociateWebACLOutput, error) 
 GetByteMatchSet(ctx context.Context, params *GetByteMatchSetInput, optFns ...func(*Options)) (*GetByteMatchSetOutput, error) 
 GetChangeToken(ctx context.Context, params *GetChangeTokenInput, optFns ...func(*Options)) (*GetChangeTokenOutput, error) 
 GetChangeTokenStatus(ctx context.Context, params *GetChangeTokenStatusInput, optFns ...func(*Options)) (*GetChangeTokenStatusOutput, error) 
 GetGeoMatchSet(ctx context.Context, params *GetGeoMatchSetInput, optFns ...func(*Options)) (*GetGeoMatchSetOutput, error) 
 GetIPSet(ctx context.Context, params *GetIPSetInput, optFns ...func(*Options)) (*GetIPSetOutput, error) 
 GetLoggingConfiguration(ctx context.Context, params *GetLoggingConfigurationInput, optFns ...func(*Options)) (*GetLoggingConfigurationOutput, error) 
 GetPermissionPolicy(ctx context.Context, params *GetPermissionPolicyInput, optFns ...func(*Options)) (*GetPermissionPolicyOutput, error) 
 GetRateBasedRule(ctx context.Context, params *GetRateBasedRuleInput, optFns ...func(*Options)) (*GetRateBasedRuleOutput, error) 
 GetRateBasedRuleManagedKeys(ctx context.Context, params *GetRateBasedRuleManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedRuleManagedKeysOutput, error) 
 GetRegexMatchSet(ctx context.Context, params *GetRegexMatchSetInput, optFns ...func(*Options)) (*GetRegexMatchSetOutput, error) 
 GetRegexPatternSet(ctx context.Context, params *GetRegexPatternSetInput, optFns ...func(*Options)) (*GetRegexPatternSetOutput, error) 
 GetRule(ctx context.Context, params *GetRuleInput, optFns ...func(*Options)) (*GetRuleOutput, error) 
 GetRuleGroup(ctx context.Context, params *GetRuleGroupInput, optFns ...func(*Options)) (*GetRuleGroupOutput, error) 
 GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) 
 GetSizeConstraintSet(ctx context.Context, params *GetSizeConstraintSetInput, optFns ...func(*Options)) (*GetSizeConstraintSetOutput, error) 
 GetSqlInjectionMatchSet(ctx context.Context, params *GetSqlInjectionMatchSetInput, optFns ...func(*Options)) (*GetSqlInjectionMatchSetOutput, error) 
 GetWebACL(ctx context.Context, params *GetWebACLInput, optFns ...func(*Options)) (*GetWebACLOutput, error) 
 GetWebACLForResource(ctx context.Context, params *GetWebACLForResourceInput, optFns ...func(*Options)) (*GetWebACLForResourceOutput, error) 
 GetXssMatchSet(ctx context.Context, params *GetXssMatchSetInput, optFns ...func(*Options)) (*GetXssMatchSetOutput, error) 
 ListActivatedRulesInRuleGroup(ctx context.Context, params *ListActivatedRulesInRuleGroupInput, optFns ...func(*Options)) (*ListActivatedRulesInRuleGroupOutput, error) 
 ListByteMatchSets(ctx context.Context, params *ListByteMatchSetsInput, optFns ...func(*Options)) (*ListByteMatchSetsOutput, error) 
 ListGeoMatchSets(ctx context.Context, params *ListGeoMatchSetsInput, optFns ...func(*Options)) (*ListGeoMatchSetsOutput, error) 
 ListIPSets(ctx context.Context, params *ListIPSetsInput, optFns ...func(*Options)) (*ListIPSetsOutput, error) 
 ListLoggingConfigurations(ctx context.Context, params *ListLoggingConfigurationsInput, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) 
 ListRateBasedRules(ctx context.Context, params *ListRateBasedRulesInput, optFns ...func(*Options)) (*ListRateBasedRulesOutput, error) 
 ListRegexMatchSets(ctx context.Context, params *ListRegexMatchSetsInput, optFns ...func(*Options)) (*ListRegexMatchSetsOutput, error) 
 ListRegexPatternSets(ctx context.Context, params *ListRegexPatternSetsInput, optFns ...func(*Options)) (*ListRegexPatternSetsOutput, error) 
 ListResourcesForWebACL(ctx context.Context, params *ListResourcesForWebACLInput, optFns ...func(*Options)) (*ListResourcesForWebACLOutput, error) 
 ListRuleGroups(ctx context.Context, params *ListRuleGroupsInput, optFns ...func(*Options)) (*ListRuleGroupsOutput, error) 
 ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) 
 ListSizeConstraintSets(ctx context.Context, params *ListSizeConstraintSetsInput, optFns ...func(*Options)) (*ListSizeConstraintSetsOutput, error) 
 ListSqlInjectionMatchSets(ctx context.Context, params *ListSqlInjectionMatchSetsInput, optFns ...func(*Options)) (*ListSqlInjectionMatchSetsOutput, error) 
 ListSubscribedRuleGroups(ctx context.Context, params *ListSubscribedRuleGroupsInput, optFns ...func(*Options)) (*ListSubscribedRuleGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebACLs(ctx context.Context, params *ListWebACLsInput, optFns ...func(*Options)) (*ListWebACLsOutput, error) 
 ListXssMatchSets(ctx context.Context, params *ListXssMatchSetsInput, optFns ...func(*Options)) (*ListXssMatchSetsOutput, error) 
 PutLoggingConfiguration(ctx context.Context, params *PutLoggingConfigurationInput, optFns ...func(*Options)) (*PutLoggingConfigurationOutput, error) 
 PutPermissionPolicy(ctx context.Context, params *PutPermissionPolicyInput, optFns ...func(*Options)) (*PutPermissionPolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateByteMatchSet(ctx context.Context, params *UpdateByteMatchSetInput, optFns ...func(*Options)) (*UpdateByteMatchSetOutput, error) 
 UpdateGeoMatchSet(ctx context.Context, params *UpdateGeoMatchSetInput, optFns ...func(*Options)) (*UpdateGeoMatchSetOutput, error) 
 UpdateIPSet(ctx context.Context, params *UpdateIPSetInput, optFns ...func(*Options)) (*UpdateIPSetOutput, error) 
 UpdateRateBasedRule(ctx context.Context, params *UpdateRateBasedRuleInput, optFns ...func(*Options)) (*UpdateRateBasedRuleOutput, error) 
 UpdateRegexMatchSet(ctx context.Context, params *UpdateRegexMatchSetInput, optFns ...func(*Options)) (*UpdateRegexMatchSetOutput, error) 
 UpdateRegexPatternSet(ctx context.Context, params *UpdateRegexPatternSetInput, optFns ...func(*Options)) (*UpdateRegexPatternSetOutput, error) 
 UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) 
 UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) 
 UpdateSizeConstraintSet(ctx context.Context, params *UpdateSizeConstraintSetInput, optFns ...func(*Options)) (*UpdateSizeConstraintSetOutput, error) 
 UpdateSqlInjectionMatchSet(ctx context.Context, params *UpdateSqlInjectionMatchSetInput, optFns ...func(*Options)) (*UpdateSqlInjectionMatchSetOutput, error) 
 UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) 
 UpdateXssMatchSet(ctx context.Context, params *UpdateXssMatchSetInput, optFns ...func(*Options)) (*UpdateXssMatchSetOutput, error) 
}
