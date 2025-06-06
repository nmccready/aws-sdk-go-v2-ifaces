
package route53recoverycontrolconfig_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
)

// IClient defines the interface for route53recoverycontrolconfig
type IClient interface {
 Options() Options 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateControlPanel(ctx context.Context, params *CreateControlPanelInput, optFns ...func(*Options)) (*CreateControlPanelOutput, error) 
 CreateRoutingControl(ctx context.Context, params *CreateRoutingControlInput, optFns ...func(*Options)) (*CreateRoutingControlOutput, error) 
 CreateSafetyRule(ctx context.Context, params *CreateSafetyRuleInput, optFns ...func(*Options)) (*CreateSafetyRuleOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteControlPanel(ctx context.Context, params *DeleteControlPanelInput, optFns ...func(*Options)) (*DeleteControlPanelOutput, error) 
 DeleteRoutingControl(ctx context.Context, params *DeleteRoutingControlInput, optFns ...func(*Options)) (*DeleteRoutingControlOutput, error) 
 DeleteSafetyRule(ctx context.Context, params *DeleteSafetyRuleInput, optFns ...func(*Options)) (*DeleteSafetyRuleOutput, error) 
 DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) 
 DescribeControlPanel(ctx context.Context, params *DescribeControlPanelInput, optFns ...func(*Options)) (*DescribeControlPanelOutput, error) 
 DescribeRoutingControl(ctx context.Context, params *DescribeRoutingControlInput, optFns ...func(*Options)) (*DescribeRoutingControlOutput, error) 
 DescribeSafetyRule(ctx context.Context, params *DescribeSafetyRuleInput, optFns ...func(*Options)) (*DescribeSafetyRuleOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 ListAssociatedRoute53HealthChecks(ctx context.Context, params *ListAssociatedRoute53HealthChecksInput, optFns ...func(*Options)) (*ListAssociatedRoute53HealthChecksOutput, error) 
 ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) 
 ListControlPanels(ctx context.Context, params *ListControlPanelsInput, optFns ...func(*Options)) (*ListControlPanelsOutput, error) 
 ListRoutingControls(ctx context.Context, params *ListRoutingControlsInput, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) 
 ListSafetyRules(ctx context.Context, params *ListSafetyRulesInput, optFns ...func(*Options)) (*ListSafetyRulesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) 
 UpdateControlPanel(ctx context.Context, params *UpdateControlPanelInput, optFns ...func(*Options)) (*UpdateControlPanelOutput, error) 
 UpdateRoutingControl(ctx context.Context, params *UpdateRoutingControlInput, optFns ...func(*Options)) (*UpdateRoutingControlOutput, error) 
 UpdateSafetyRule(ctx context.Context, params *UpdateSafetyRuleInput, optFns ...func(*Options)) (*UpdateSafetyRuleOutput, error) 
}
