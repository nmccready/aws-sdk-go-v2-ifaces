
package worklink

import (
    "github.com/aws/aws-sdk-go-v2/service/worklink"
)

// IClient defines the interface for worklink
type IClient interface {
 Options() Options 
 AssociateDomain(ctx context.Context, params *AssociateDomainInput, optFns ...func(*Options)) (*AssociateDomainOutput, error) 
 AssociateWebsiteAuthorizationProvider(ctx context.Context, params *AssociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*AssociateWebsiteAuthorizationProviderOutput, error) 
 AssociateWebsiteCertificateAuthority(ctx context.Context, params *AssociateWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*AssociateWebsiteCertificateAuthorityOutput, error) 
 CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) 
 DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) 
 DescribeAuditStreamConfiguration(ctx context.Context, params *DescribeAuditStreamConfigurationInput, optFns ...func(*Options)) (*DescribeAuditStreamConfigurationOutput, error) 
 DescribeCompanyNetworkConfiguration(ctx context.Context, params *DescribeCompanyNetworkConfigurationInput, optFns ...func(*Options)) (*DescribeCompanyNetworkConfigurationOutput, error) 
 DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) 
 DescribeDevicePolicyConfiguration(ctx context.Context, params *DescribeDevicePolicyConfigurationInput, optFns ...func(*Options)) (*DescribeDevicePolicyConfigurationOutput, error) 
 DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) 
 DescribeFleetMetadata(ctx context.Context, params *DescribeFleetMetadataInput, optFns ...func(*Options)) (*DescribeFleetMetadataOutput, error) 
 DescribeIdentityProviderConfiguration(ctx context.Context, params *DescribeIdentityProviderConfigurationInput, optFns ...func(*Options)) (*DescribeIdentityProviderConfigurationOutput, error) 
 DescribeWebsiteCertificateAuthority(ctx context.Context, params *DescribeWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*DescribeWebsiteCertificateAuthorityOutput, error) 
 DisassociateDomain(ctx context.Context, params *DisassociateDomainInput, optFns ...func(*Options)) (*DisassociateDomainOutput, error) 
 DisassociateWebsiteAuthorizationProvider(ctx context.Context, params *DisassociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*DisassociateWebsiteAuthorizationProviderOutput, error) 
 DisassociateWebsiteCertificateAuthority(ctx context.Context, params *DisassociateWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*DisassociateWebsiteCertificateAuthorityOutput, error) 
 ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWebsiteAuthorizationProviders(ctx context.Context, params *ListWebsiteAuthorizationProvidersInput, optFns ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error) 
 ListWebsiteCertificateAuthorities(ctx context.Context, params *ListWebsiteCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListWebsiteCertificateAuthoritiesOutput, error) 
 RestoreDomainAccess(ctx context.Context, params *RestoreDomainAccessInput, optFns ...func(*Options)) (*RestoreDomainAccessOutput, error) 
 RevokeDomainAccess(ctx context.Context, params *RevokeDomainAccessInput, optFns ...func(*Options)) (*RevokeDomainAccessOutput, error) 
 SignOutUser(ctx context.Context, params *SignOutUserInput, optFns ...func(*Options)) (*SignOutUserOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAuditStreamConfiguration(ctx context.Context, params *UpdateAuditStreamConfigurationInput, optFns ...func(*Options)) (*UpdateAuditStreamConfigurationOutput, error) 
 UpdateCompanyNetworkConfiguration(ctx context.Context, params *UpdateCompanyNetworkConfigurationInput, optFns ...func(*Options)) (*UpdateCompanyNetworkConfigurationOutput, error) 
 UpdateDevicePolicyConfiguration(ctx context.Context, params *UpdateDevicePolicyConfigurationInput, optFns ...func(*Options)) (*UpdateDevicePolicyConfigurationOutput, error) 
 UpdateDomainMetadata(ctx context.Context, params *UpdateDomainMetadataInput, optFns ...func(*Options)) (*UpdateDomainMetadataOutput, error) 
 UpdateFleetMetadata(ctx context.Context, params *UpdateFleetMetadataInput, optFns ...func(*Options)) (*UpdateFleetMetadataOutput, error) 
 UpdateIdentityProviderConfiguration(ctx context.Context, params *UpdateIdentityProviderConfigurationInput, optFns ...func(*Options)) (*UpdateIdentityProviderConfigurationOutput, error) 
}
