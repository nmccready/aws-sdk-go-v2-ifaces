
package elasticloadbalancingv2_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

// IClient defines the interface for elasticloadbalancingv2
type IClient interface {
 Options() Options 
 AddListenerCertificates(ctx context.Context, params *AddListenerCertificatesInput, optFns ...func(*Options)) (*AddListenerCertificatesOutput, error) 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 AddTrustStoreRevocations(ctx context.Context, params *AddTrustStoreRevocationsInput, optFns ...func(*Options)) (*AddTrustStoreRevocationsOutput, error) 
 CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) 
 CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) 
 CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) 
 CreateTargetGroup(ctx context.Context, params *CreateTargetGroupInput, optFns ...func(*Options)) (*CreateTargetGroupOutput, error) 
 CreateTrustStore(ctx context.Context, params *CreateTrustStoreInput, optFns ...func(*Options)) (*CreateTrustStoreOutput, error) 
 DeleteListener(ctx context.Context, params *DeleteListenerInput, optFns ...func(*Options)) (*DeleteListenerOutput, error) 
 DeleteLoadBalancer(ctx context.Context, params *DeleteLoadBalancerInput, optFns ...func(*Options)) (*DeleteLoadBalancerOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 DeleteSharedTrustStoreAssociation(ctx context.Context, params *DeleteSharedTrustStoreAssociationInput, optFns ...func(*Options)) (*DeleteSharedTrustStoreAssociationOutput, error) 
 DeleteTargetGroup(ctx context.Context, params *DeleteTargetGroupInput, optFns ...func(*Options)) (*DeleteTargetGroupOutput, error) 
 DeleteTrustStore(ctx context.Context, params *DeleteTrustStoreInput, optFns ...func(*Options)) (*DeleteTrustStoreOutput, error) 
 DeregisterTargets(ctx context.Context, params *DeregisterTargetsInput, optFns ...func(*Options)) (*DeregisterTargetsOutput, error) 
 DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) 
 DescribeCapacityReservation(ctx context.Context, params *DescribeCapacityReservationInput, optFns ...func(*Options)) (*DescribeCapacityReservationOutput, error) 
 DescribeListenerAttributes(ctx context.Context, params *DescribeListenerAttributesInput, optFns ...func(*Options)) (*DescribeListenerAttributesOutput, error) 
 DescribeListenerCertificates(ctx context.Context, params *DescribeListenerCertificatesInput, optFns ...func(*Options)) (*DescribeListenerCertificatesOutput, error) 
 DescribeListeners(ctx context.Context, params *DescribeListenersInput, optFns ...func(*Options)) (*DescribeListenersOutput, error) 
 DescribeLoadBalancerAttributes(ctx context.Context, params *DescribeLoadBalancerAttributesInput, optFns ...func(*Options)) (*DescribeLoadBalancerAttributesOutput, error) 
 DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) 
 DescribeRules(ctx context.Context, params *DescribeRulesInput, optFns ...func(*Options)) (*DescribeRulesOutput, error) 
 DescribeSSLPolicies(ctx context.Context, params *DescribeSSLPoliciesInput, optFns ...func(*Options)) (*DescribeSSLPoliciesOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 DescribeTargetGroupAttributes(ctx context.Context, params *DescribeTargetGroupAttributesInput, optFns ...func(*Options)) (*DescribeTargetGroupAttributesOutput, error) 
 DescribeTargetGroups(ctx context.Context, params *DescribeTargetGroupsInput, optFns ...func(*Options)) (*DescribeTargetGroupsOutput, error) 
 DescribeTargetHealth(ctx context.Context, params *DescribeTargetHealthInput, optFns ...func(*Options)) (*DescribeTargetHealthOutput, error) 
 DescribeTrustStoreAssociations(ctx context.Context, params *DescribeTrustStoreAssociationsInput, optFns ...func(*Options)) (*DescribeTrustStoreAssociationsOutput, error) 
 DescribeTrustStoreRevocations(ctx context.Context, params *DescribeTrustStoreRevocationsInput, optFns ...func(*Options)) (*DescribeTrustStoreRevocationsOutput, error) 
 DescribeTrustStores(ctx context.Context, params *DescribeTrustStoresInput, optFns ...func(*Options)) (*DescribeTrustStoresOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 GetTrustStoreCaCertificatesBundle(ctx context.Context, params *GetTrustStoreCaCertificatesBundleInput, optFns ...func(*Options)) (*GetTrustStoreCaCertificatesBundleOutput, error) 
 GetTrustStoreRevocationContent(ctx context.Context, params *GetTrustStoreRevocationContentInput, optFns ...func(*Options)) (*GetTrustStoreRevocationContentOutput, error) 
 ModifyCapacityReservation(ctx context.Context, params *ModifyCapacityReservationInput, optFns ...func(*Options)) (*ModifyCapacityReservationOutput, error) 
 ModifyListener(ctx context.Context, params *ModifyListenerInput, optFns ...func(*Options)) (*ModifyListenerOutput, error) 
 ModifyListenerAttributes(ctx context.Context, params *ModifyListenerAttributesInput, optFns ...func(*Options)) (*ModifyListenerAttributesOutput, error) 
 ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error) 
 ModifyRule(ctx context.Context, params *ModifyRuleInput, optFns ...func(*Options)) (*ModifyRuleOutput, error) 
 ModifyTargetGroup(ctx context.Context, params *ModifyTargetGroupInput, optFns ...func(*Options)) (*ModifyTargetGroupOutput, error) 
 ModifyTargetGroupAttributes(ctx context.Context, params *ModifyTargetGroupAttributesInput, optFns ...func(*Options)) (*ModifyTargetGroupAttributesOutput, error) 
 ModifyTrustStore(ctx context.Context, params *ModifyTrustStoreInput, optFns ...func(*Options)) (*ModifyTrustStoreOutput, error) 
 RegisterTargets(ctx context.Context, params *RegisterTargetsInput, optFns ...func(*Options)) (*RegisterTargetsOutput, error) 
 RemoveListenerCertificates(ctx context.Context, params *RemoveListenerCertificatesInput, optFns ...func(*Options)) (*RemoveListenerCertificatesOutput, error) 
 RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) 
 RemoveTrustStoreRevocations(ctx context.Context, params *RemoveTrustStoreRevocationsInput, optFns ...func(*Options)) (*RemoveTrustStoreRevocationsOutput, error) 
 SetIpAddressType(ctx context.Context, params *SetIpAddressTypeInput, optFns ...func(*Options)) (*SetIpAddressTypeOutput, error) 
 SetRulePriorities(ctx context.Context, params *SetRulePrioritiesInput, optFns ...func(*Options)) (*SetRulePrioritiesOutput, error) 
 SetSecurityGroups(ctx context.Context, params *SetSecurityGroupsInput, optFns ...func(*Options)) (*SetSecurityGroupsOutput, error) 
 SetSubnets(ctx context.Context, params *SetSubnetsInput, optFns ...func(*Options)) (*SetSubnetsOutput, error) 
}
