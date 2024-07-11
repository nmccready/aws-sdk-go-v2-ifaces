
package eks

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/eks"
)

// IClient defines the interface for eks
type IClient interface {
 Options() Options 
 AssociateAccessPolicy(ctx context.Context, params *AssociateAccessPolicyInput, optFns ...func(*Options)) (*AssociateAccessPolicyOutput, error) 
 AssociateEncryptionConfig(ctx context.Context, params *AssociateEncryptionConfigInput, optFns ...func(*Options)) (*AssociateEncryptionConfigOutput, error) 
 AssociateIdentityProviderConfig(ctx context.Context, params *AssociateIdentityProviderConfigInput, optFns ...func(*Options)) (*AssociateIdentityProviderConfigOutput, error) 
 CreateAccessEntry(ctx context.Context, params *CreateAccessEntryInput, optFns ...func(*Options)) (*CreateAccessEntryOutput, error) 
 CreateAddon(ctx context.Context, params *CreateAddonInput, optFns ...func(*Options)) (*CreateAddonOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateEksAnywhereSubscription(ctx context.Context, params *CreateEksAnywhereSubscriptionInput, optFns ...func(*Options)) (*CreateEksAnywhereSubscriptionOutput, error) 
 CreateFargateProfile(ctx context.Context, params *CreateFargateProfileInput, optFns ...func(*Options)) (*CreateFargateProfileOutput, error) 
 CreateNodegroup(ctx context.Context, params *CreateNodegroupInput, optFns ...func(*Options)) (*CreateNodegroupOutput, error) 
 CreatePodIdentityAssociation(ctx context.Context, params *CreatePodIdentityAssociationInput, optFns ...func(*Options)) (*CreatePodIdentityAssociationOutput, error) 
 DeleteAccessEntry(ctx context.Context, params *DeleteAccessEntryInput, optFns ...func(*Options)) (*DeleteAccessEntryOutput, error) 
 DeleteAddon(ctx context.Context, params *DeleteAddonInput, optFns ...func(*Options)) (*DeleteAddonOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteEksAnywhereSubscription(ctx context.Context, params *DeleteEksAnywhereSubscriptionInput, optFns ...func(*Options)) (*DeleteEksAnywhereSubscriptionOutput, error) 
 DeleteFargateProfile(ctx context.Context, params *DeleteFargateProfileInput, optFns ...func(*Options)) (*DeleteFargateProfileOutput, error) 
 DeleteNodegroup(ctx context.Context, params *DeleteNodegroupInput, optFns ...func(*Options)) (*DeleteNodegroupOutput, error) 
 DeletePodIdentityAssociation(ctx context.Context, params *DeletePodIdentityAssociationInput, optFns ...func(*Options)) (*DeletePodIdentityAssociationOutput, error) 
 DeregisterCluster(ctx context.Context, params *DeregisterClusterInput, optFns ...func(*Options)) (*DeregisterClusterOutput, error) 
 DescribeAccessEntry(ctx context.Context, params *DescribeAccessEntryInput, optFns ...func(*Options)) (*DescribeAccessEntryOutput, error) 
 DescribeAddon(ctx context.Context, params *DescribeAddonInput, optFns ...func(*Options)) (*DescribeAddonOutput, error) 
 DescribeAddonConfiguration(ctx context.Context, params *DescribeAddonConfigurationInput, optFns ...func(*Options)) (*DescribeAddonConfigurationOutput, error) 
 DescribeAddonVersions(ctx context.Context, params *DescribeAddonVersionsInput, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) 
 DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) 
 DescribeEksAnywhereSubscription(ctx context.Context, params *DescribeEksAnywhereSubscriptionInput, optFns ...func(*Options)) (*DescribeEksAnywhereSubscriptionOutput, error) 
 DescribeFargateProfile(ctx context.Context, params *DescribeFargateProfileInput, optFns ...func(*Options)) (*DescribeFargateProfileOutput, error) 
 DescribeIdentityProviderConfig(ctx context.Context, params *DescribeIdentityProviderConfigInput, optFns ...func(*Options)) (*DescribeIdentityProviderConfigOutput, error) 
 DescribeInsight(ctx context.Context, params *DescribeInsightInput, optFns ...func(*Options)) (*DescribeInsightOutput, error) 
 DescribeNodegroup(ctx context.Context, params *DescribeNodegroupInput, optFns ...func(*Options)) (*DescribeNodegroupOutput, error) 
 DescribePodIdentityAssociation(ctx context.Context, params *DescribePodIdentityAssociationInput, optFns ...func(*Options)) (*DescribePodIdentityAssociationOutput, error) 
 DescribeUpdate(ctx context.Context, params *DescribeUpdateInput, optFns ...func(*Options)) (*DescribeUpdateOutput, error) 
 DisassociateAccessPolicy(ctx context.Context, params *DisassociateAccessPolicyInput, optFns ...func(*Options)) (*DisassociateAccessPolicyOutput, error) 
 DisassociateIdentityProviderConfig(ctx context.Context, params *DisassociateIdentityProviderConfigInput, optFns ...func(*Options)) (*DisassociateIdentityProviderConfigOutput, error) 
 ListAccessEntries(ctx context.Context, params *ListAccessEntriesInput, optFns ...func(*Options)) (*ListAccessEntriesOutput, error) 
 ListAccessPolicies(ctx context.Context, params *ListAccessPoliciesInput, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) 
 ListAddons(ctx context.Context, params *ListAddonsInput, optFns ...func(*Options)) (*ListAddonsOutput, error) 
 ListAssociatedAccessPolicies(ctx context.Context, params *ListAssociatedAccessPoliciesInput, optFns ...func(*Options)) (*ListAssociatedAccessPoliciesOutput, error) 
 ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) 
 ListEksAnywhereSubscriptions(ctx context.Context, params *ListEksAnywhereSubscriptionsInput, optFns ...func(*Options)) (*ListEksAnywhereSubscriptionsOutput, error) 
 ListFargateProfiles(ctx context.Context, params *ListFargateProfilesInput, optFns ...func(*Options)) (*ListFargateProfilesOutput, error) 
 ListIdentityProviderConfigs(ctx context.Context, params *ListIdentityProviderConfigsInput, optFns ...func(*Options)) (*ListIdentityProviderConfigsOutput, error) 
 ListInsights(ctx context.Context, params *ListInsightsInput, optFns ...func(*Options)) (*ListInsightsOutput, error) 
 ListNodegroups(ctx context.Context, params *ListNodegroupsInput, optFns ...func(*Options)) (*ListNodegroupsOutput, error) 
 ListPodIdentityAssociations(ctx context.Context, params *ListPodIdentityAssociationsInput, optFns ...func(*Options)) (*ListPodIdentityAssociationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListUpdates(ctx context.Context, params *ListUpdatesInput, optFns ...func(*Options)) (*ListUpdatesOutput, error) 
 RegisterCluster(ctx context.Context, params *RegisterClusterInput, optFns ...func(*Options)) (*RegisterClusterOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccessEntry(ctx context.Context, params *UpdateAccessEntryInput, optFns ...func(*Options)) (*UpdateAccessEntryOutput, error) 
 UpdateAddon(ctx context.Context, params *UpdateAddonInput, optFns ...func(*Options)) (*UpdateAddonOutput, error) 
 UpdateClusterConfig(ctx context.Context, params *UpdateClusterConfigInput, optFns ...func(*Options)) (*UpdateClusterConfigOutput, error) 
 UpdateClusterVersion(ctx context.Context, params *UpdateClusterVersionInput, optFns ...func(*Options)) (*UpdateClusterVersionOutput, error) 
 UpdateEksAnywhereSubscription(ctx context.Context, params *UpdateEksAnywhereSubscriptionInput, optFns ...func(*Options)) (*UpdateEksAnywhereSubscriptionOutput, error) 
 UpdateNodegroupConfig(ctx context.Context, params *UpdateNodegroupConfigInput, optFns ...func(*Options)) (*UpdateNodegroupConfigOutput, error) 
 UpdateNodegroupVersion(ctx context.Context, params *UpdateNodegroupVersionInput, optFns ...func(*Options)) (*UpdateNodegroupVersionOutput, error) 
 UpdatePodIdentityAssociation(ctx context.Context, params *UpdatePodIdentityAssociationInput, optFns ...func(*Options)) (*UpdatePodIdentityAssociationOutput, error) 
}
