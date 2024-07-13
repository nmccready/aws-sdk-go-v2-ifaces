
package networkfirewall_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
)

// IClient defines the interface for networkfirewall
type IClient interface {
 Options() Options 
 AssociateFirewallPolicy(ctx context.Context, params *AssociateFirewallPolicyInput, optFns ...func(*Options)) (*AssociateFirewallPolicyOutput, error) 
 AssociateSubnets(ctx context.Context, params *AssociateSubnetsInput, optFns ...func(*Options)) (*AssociateSubnetsOutput, error) 
 CreateFirewall(ctx context.Context, params *CreateFirewallInput, optFns ...func(*Options)) (*CreateFirewallOutput, error) 
 CreateFirewallPolicy(ctx context.Context, params *CreateFirewallPolicyInput, optFns ...func(*Options)) (*CreateFirewallPolicyOutput, error) 
 CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) 
 CreateTLSInspectionConfiguration(ctx context.Context, params *CreateTLSInspectionConfigurationInput, optFns ...func(*Options)) (*CreateTLSInspectionConfigurationOutput, error) 
 DeleteFirewall(ctx context.Context, params *DeleteFirewallInput, optFns ...func(*Options)) (*DeleteFirewallOutput, error) 
 DeleteFirewallPolicy(ctx context.Context, params *DeleteFirewallPolicyInput, optFns ...func(*Options)) (*DeleteFirewallPolicyOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteRuleGroup(ctx context.Context, params *DeleteRuleGroupInput, optFns ...func(*Options)) (*DeleteRuleGroupOutput, error) 
 DeleteTLSInspectionConfiguration(ctx context.Context, params *DeleteTLSInspectionConfigurationInput, optFns ...func(*Options)) (*DeleteTLSInspectionConfigurationOutput, error) 
 DescribeFirewall(ctx context.Context, params *DescribeFirewallInput, optFns ...func(*Options)) (*DescribeFirewallOutput, error) 
 DescribeFirewallPolicy(ctx context.Context, params *DescribeFirewallPolicyInput, optFns ...func(*Options)) (*DescribeFirewallPolicyOutput, error) 
 DescribeLoggingConfiguration(ctx context.Context, params *DescribeLoggingConfigurationInput, optFns ...func(*Options)) (*DescribeLoggingConfigurationOutput, error) 
 DescribeResourcePolicy(ctx context.Context, params *DescribeResourcePolicyInput, optFns ...func(*Options)) (*DescribeResourcePolicyOutput, error) 
 DescribeRuleGroup(ctx context.Context, params *DescribeRuleGroupInput, optFns ...func(*Options)) (*DescribeRuleGroupOutput, error) 
 DescribeRuleGroupMetadata(ctx context.Context, params *DescribeRuleGroupMetadataInput, optFns ...func(*Options)) (*DescribeRuleGroupMetadataOutput, error) 
 DescribeTLSInspectionConfiguration(ctx context.Context, params *DescribeTLSInspectionConfigurationInput, optFns ...func(*Options)) (*DescribeTLSInspectionConfigurationOutput, error) 
 DisassociateSubnets(ctx context.Context, params *DisassociateSubnetsInput, optFns ...func(*Options)) (*DisassociateSubnetsOutput, error) 
 ListFirewallPolicies(ctx context.Context, params *ListFirewallPoliciesInput, optFns ...func(*Options)) (*ListFirewallPoliciesOutput, error) 
 ListFirewalls(ctx context.Context, params *ListFirewallsInput, optFns ...func(*Options)) (*ListFirewallsOutput, error) 
 ListRuleGroups(ctx context.Context, params *ListRuleGroupsInput, optFns ...func(*Options)) (*ListRuleGroupsOutput, error) 
 ListTLSInspectionConfigurations(ctx context.Context, params *ListTLSInspectionConfigurationsInput, optFns ...func(*Options)) (*ListTLSInspectionConfigurationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateFirewallDeleteProtection(ctx context.Context, params *UpdateFirewallDeleteProtectionInput, optFns ...func(*Options)) (*UpdateFirewallDeleteProtectionOutput, error) 
 UpdateFirewallDescription(ctx context.Context, params *UpdateFirewallDescriptionInput, optFns ...func(*Options)) (*UpdateFirewallDescriptionOutput, error) 
 UpdateFirewallEncryptionConfiguration(ctx context.Context, params *UpdateFirewallEncryptionConfigurationInput, optFns ...func(*Options)) (*UpdateFirewallEncryptionConfigurationOutput, error) 
 UpdateFirewallPolicy(ctx context.Context, params *UpdateFirewallPolicyInput, optFns ...func(*Options)) (*UpdateFirewallPolicyOutput, error) 
 UpdateFirewallPolicyChangeProtection(ctx context.Context, params *UpdateFirewallPolicyChangeProtectionInput, optFns ...func(*Options)) (*UpdateFirewallPolicyChangeProtectionOutput, error) 
 UpdateLoggingConfiguration(ctx context.Context, params *UpdateLoggingConfigurationInput, optFns ...func(*Options)) (*UpdateLoggingConfigurationOutput, error) 
 UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) 
 UpdateSubnetChangeProtection(ctx context.Context, params *UpdateSubnetChangeProtectionInput, optFns ...func(*Options)) (*UpdateSubnetChangeProtectionOutput, error) 
 UpdateTLSInspectionConfiguration(ctx context.Context, params *UpdateTLSInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateTLSInspectionConfigurationOutput, error) 
}