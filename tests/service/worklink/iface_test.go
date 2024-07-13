package worklink_test

// tests for the worklink service interface for this ../../../service/worklink/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/worklink/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/worklink/worklink_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorklinkServiceCanBeMocked(t *testing.T) {
	var iface worklink_iface.IClient
	iface = &worklink.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := worklink.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDomain", func(t *testing.T) {
        input := &worklink.AssociateDomainInput{}
        output := &worklink.AssociateDomainOutput{}

        mockClient.On("AssociateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWebsiteAuthorizationProvider", func(t *testing.T) {
        input := &worklink.AssociateWebsiteAuthorizationProviderInput{}
        output := &worklink.AssociateWebsiteAuthorizationProviderOutput{}

        mockClient.On("AssociateWebsiteAuthorizationProvider", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWebsiteAuthorizationProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWebsiteCertificateAuthority", func(t *testing.T) {
        input := &worklink.AssociateWebsiteCertificateAuthorityInput{}
        output := &worklink.AssociateWebsiteCertificateAuthorityOutput{}

        mockClient.On("AssociateWebsiteCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWebsiteCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &worklink.CreateFleetInput{}
        output := &worklink.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &worklink.DeleteFleetInput{}
        output := &worklink.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuditStreamConfiguration", func(t *testing.T) {
        input := &worklink.DescribeAuditStreamConfigurationInput{}
        output := &worklink.DescribeAuditStreamConfigurationOutput{}

        mockClient.On("DescribeAuditStreamConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuditStreamConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCompanyNetworkConfiguration", func(t *testing.T) {
        input := &worklink.DescribeCompanyNetworkConfigurationInput{}
        output := &worklink.DescribeCompanyNetworkConfigurationOutput{}

        mockClient.On("DescribeCompanyNetworkConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCompanyNetworkConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevice", func(t *testing.T) {
        input := &worklink.DescribeDeviceInput{}
        output := &worklink.DescribeDeviceOutput{}

        mockClient.On("DescribeDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevicePolicyConfiguration", func(t *testing.T) {
        input := &worklink.DescribeDevicePolicyConfigurationInput{}
        output := &worklink.DescribeDevicePolicyConfigurationOutput{}

        mockClient.On("DescribeDevicePolicyConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevicePolicyConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &worklink.DescribeDomainInput{}
        output := &worklink.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetMetadata", func(t *testing.T) {
        input := &worklink.DescribeFleetMetadataInput{}
        output := &worklink.DescribeFleetMetadataOutput{}

        mockClient.On("DescribeFleetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityProviderConfiguration", func(t *testing.T) {
        input := &worklink.DescribeIdentityProviderConfigurationInput{}
        output := &worklink.DescribeIdentityProviderConfigurationOutput{}

        mockClient.On("DescribeIdentityProviderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityProviderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWebsiteCertificateAuthority", func(t *testing.T) {
        input := &worklink.DescribeWebsiteCertificateAuthorityInput{}
        output := &worklink.DescribeWebsiteCertificateAuthorityOutput{}

        mockClient.On("DescribeWebsiteCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWebsiteCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDomain", func(t *testing.T) {
        input := &worklink.DisassociateDomainInput{}
        output := &worklink.DisassociateDomainOutput{}

        mockClient.On("DisassociateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWebsiteAuthorizationProvider", func(t *testing.T) {
        input := &worklink.DisassociateWebsiteAuthorizationProviderInput{}
        output := &worklink.DisassociateWebsiteAuthorizationProviderOutput{}

        mockClient.On("DisassociateWebsiteAuthorizationProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWebsiteAuthorizationProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWebsiteCertificateAuthority", func(t *testing.T) {
        input := &worklink.DisassociateWebsiteCertificateAuthorityInput{}
        output := &worklink.DisassociateWebsiteCertificateAuthorityOutput{}

        mockClient.On("DisassociateWebsiteCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWebsiteCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &worklink.ListDevicesInput{}
        output := &worklink.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &worklink.ListDomainsInput{}
        output := &worklink.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &worklink.ListFleetsInput{}
        output := &worklink.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &worklink.ListTagsForResourceInput{}
        output := &worklink.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebsiteAuthorizationProviders", func(t *testing.T) {
        input := &worklink.ListWebsiteAuthorizationProvidersInput{}
        output := &worklink.ListWebsiteAuthorizationProvidersOutput{}

        mockClient.On("ListWebsiteAuthorizationProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebsiteAuthorizationProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebsiteCertificateAuthorities", func(t *testing.T) {
        input := &worklink.ListWebsiteCertificateAuthoritiesInput{}
        output := &worklink.ListWebsiteCertificateAuthoritiesOutput{}

        mockClient.On("ListWebsiteCertificateAuthorities", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebsiteCertificateAuthorities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreDomainAccess", func(t *testing.T) {
        input := &worklink.RestoreDomainAccessInput{}
        output := &worklink.RestoreDomainAccessOutput{}

        mockClient.On("RestoreDomainAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreDomainAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeDomainAccess", func(t *testing.T) {
        input := &worklink.RevokeDomainAccessInput{}
        output := &worklink.RevokeDomainAccessOutput{}

        mockClient.On("RevokeDomainAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeDomainAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignOutUser", func(t *testing.T) {
        input := &worklink.SignOutUserInput{}
        output := &worklink.SignOutUserOutput{}

        mockClient.On("SignOutUser", ctx, input).Return(output, nil)

        result, err := mockClient.SignOutUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &worklink.TagResourceInput{}
        output := &worklink.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &worklink.UntagResourceInput{}
        output := &worklink.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuditStreamConfiguration", func(t *testing.T) {
        input := &worklink.UpdateAuditStreamConfigurationInput{}
        output := &worklink.UpdateAuditStreamConfigurationOutput{}

        mockClient.On("UpdateAuditStreamConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuditStreamConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCompanyNetworkConfiguration", func(t *testing.T) {
        input := &worklink.UpdateCompanyNetworkConfigurationInput{}
        output := &worklink.UpdateCompanyNetworkConfigurationOutput{}

        mockClient.On("UpdateCompanyNetworkConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCompanyNetworkConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevicePolicyConfiguration", func(t *testing.T) {
        input := &worklink.UpdateDevicePolicyConfigurationInput{}
        output := &worklink.UpdateDevicePolicyConfigurationOutput{}

        mockClient.On("UpdateDevicePolicyConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevicePolicyConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainMetadata", func(t *testing.T) {
        input := &worklink.UpdateDomainMetadataInput{}
        output := &worklink.UpdateDomainMetadataOutput{}

        mockClient.On("UpdateDomainMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleetMetadata", func(t *testing.T) {
        input := &worklink.UpdateFleetMetadataInput{}
        output := &worklink.UpdateFleetMetadataOutput{}

        mockClient.On("UpdateFleetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityProviderConfiguration", func(t *testing.T) {
        input := &worklink.UpdateIdentityProviderConfigurationInput{}
        output := &worklink.UpdateIdentityProviderConfigurationOutput{}

        mockClient.On("UpdateIdentityProviderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityProviderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
