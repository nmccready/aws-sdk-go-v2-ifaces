
package elasticloadbalancing

import (
    "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

// IClient defines the interface for elasticloadbalancing
type IClient interface {
 Options() Options 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 ApplySecurityGroupsToLoadBalancer(ctx context.Context, params *ApplySecurityGroupsToLoadBalancerInput, optFns ...func(*Options)) (*ApplySecurityGroupsToLoadBalancerOutput, error) 
 AttachLoadBalancerToSubnets(ctx context.Context, params *AttachLoadBalancerToSubnetsInput, optFns ...func(*Options)) (*AttachLoadBalancerToSubnetsOutput, error) 
 ConfigureHealthCheck(ctx context.Context, params *ConfigureHealthCheckInput, optFns ...func(*Options)) (*ConfigureHealthCheckOutput, error) 
 CreateAppCookieStickinessPolicy(ctx context.Context, params *CreateAppCookieStickinessPolicyInput, optFns ...func(*Options)) (*CreateAppCookieStickinessPolicyOutput, error) 
 CreateLBCookieStickinessPolicy(ctx context.Context, params *CreateLBCookieStickinessPolicyInput, optFns ...func(*Options)) (*CreateLBCookieStickinessPolicyOutput, error) 
 CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) 
 CreateLoadBalancerListeners(ctx context.Context, params *CreateLoadBalancerListenersInput, optFns ...func(*Options)) (*CreateLoadBalancerListenersOutput, error) 
 CreateLoadBalancerPolicy(ctx context.Context, params *CreateLoadBalancerPolicyInput, optFns ...func(*Options)) (*CreateLoadBalancerPolicyOutput, error) 
 DeleteLoadBalancer(ctx context.Context, params *DeleteLoadBalancerInput, optFns ...func(*Options)) (*DeleteLoadBalancerOutput, error) 
 DeleteLoadBalancerListeners(ctx context.Context, params *DeleteLoadBalancerListenersInput, optFns ...func(*Options)) (*DeleteLoadBalancerListenersOutput, error) 
 DeleteLoadBalancerPolicy(ctx context.Context, params *DeleteLoadBalancerPolicyInput, optFns ...func(*Options)) (*DeleteLoadBalancerPolicyOutput, error) 
 DeregisterInstancesFromLoadBalancer(ctx context.Context, params *DeregisterInstancesFromLoadBalancerInput, optFns ...func(*Options)) (*DeregisterInstancesFromLoadBalancerOutput, error) 
 DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) 
 DescribeInstanceHealth(ctx context.Context, params *DescribeInstanceHealthInput, optFns ...func(*Options)) (*DescribeInstanceHealthOutput, error) 
 DescribeLoadBalancerAttributes(ctx context.Context, params *DescribeLoadBalancerAttributesInput, optFns ...func(*Options)) (*DescribeLoadBalancerAttributesOutput, error) 
 DescribeLoadBalancerPolicies(ctx context.Context, params *DescribeLoadBalancerPoliciesInput, optFns ...func(*Options)) (*DescribeLoadBalancerPoliciesOutput, error) 
 DescribeLoadBalancerPolicyTypes(ctx context.Context, params *DescribeLoadBalancerPolicyTypesInput, optFns ...func(*Options)) (*DescribeLoadBalancerPolicyTypesOutput, error) 
 DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 DetachLoadBalancerFromSubnets(ctx context.Context, params *DetachLoadBalancerFromSubnetsInput, optFns ...func(*Options)) (*DetachLoadBalancerFromSubnetsOutput, error) 
 DisableAvailabilityZonesForLoadBalancer(ctx context.Context, params *DisableAvailabilityZonesForLoadBalancerInput, optFns ...func(*Options)) (*DisableAvailabilityZonesForLoadBalancerOutput, error) 
 EnableAvailabilityZonesForLoadBalancer(ctx context.Context, params *EnableAvailabilityZonesForLoadBalancerInput, optFns ...func(*Options)) (*EnableAvailabilityZonesForLoadBalancerOutput, error) 
 ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error) 
 RegisterInstancesWithLoadBalancer(ctx context.Context, params *RegisterInstancesWithLoadBalancerInput, optFns ...func(*Options)) (*RegisterInstancesWithLoadBalancerOutput, error) 
 RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) 
 SetLoadBalancerListenerSSLCertificate(ctx context.Context, params *SetLoadBalancerListenerSSLCertificateInput, optFns ...func(*Options)) (*SetLoadBalancerListenerSSLCertificateOutput, error) 
 SetLoadBalancerPoliciesForBackendServer(ctx context.Context, params *SetLoadBalancerPoliciesForBackendServerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesForBackendServerOutput, error) 
 SetLoadBalancerPoliciesOfListener(ctx context.Context, params *SetLoadBalancerPoliciesOfListenerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesOfListenerOutput, error) 
}
