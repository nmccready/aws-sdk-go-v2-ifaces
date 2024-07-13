package elasticsearchservice_test

// tests for the elasticsearchservice service interface for this ../../../service/elasticsearchservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticsearchservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticsearchservice/elasticsearchservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticsearchserviceServiceCanBeMocked(t *testing.T) {
	var iface elasticsearchservice_iface.IClient
	iface = &elasticsearchservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticsearchservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInboundCrossClusterSearchConnection", func(t *testing.T) {
        input := &elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput{}
        output := &elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput{}

        mockClient.On("AcceptInboundCrossClusterSearchConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInboundCrossClusterSearchConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &elasticsearchservice.AddTagsInput{}
        output := &elasticsearchservice.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePackage", func(t *testing.T) {
        input := &elasticsearchservice.AssociatePackageInput{}
        output := &elasticsearchservice.AssociatePackageOutput{}

        mockClient.On("AssociatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeVpcEndpointAccess", func(t *testing.T) {
        input := &elasticsearchservice.AuthorizeVpcEndpointAccessInput{}
        output := &elasticsearchservice.AuthorizeVpcEndpointAccessOutput{}

        mockClient.On("AuthorizeVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDomainConfigChange", func(t *testing.T) {
        input := &elasticsearchservice.CancelDomainConfigChangeInput{}
        output := &elasticsearchservice.CancelDomainConfigChangeOutput{}

        mockClient.On("CancelDomainConfigChange", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDomainConfigChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelElasticsearchServiceSoftwareUpdate", func(t *testing.T) {
        input := &elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput{}
        output := &elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput{}

        mockClient.On("CancelElasticsearchServiceSoftwareUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.CancelElasticsearchServiceSoftwareUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateElasticsearchDomain", func(t *testing.T) {
        input := &elasticsearchservice.CreateElasticsearchDomainInput{}
        output := &elasticsearchservice.CreateElasticsearchDomainOutput{}

        mockClient.On("CreateElasticsearchDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateElasticsearchDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOutboundCrossClusterSearchConnection", func(t *testing.T) {
        input := &elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput{}
        output := &elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput{}

        mockClient.On("CreateOutboundCrossClusterSearchConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOutboundCrossClusterSearchConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackage", func(t *testing.T) {
        input := &elasticsearchservice.CreatePackageInput{}
        output := &elasticsearchservice.CreatePackageOutput{}

        mockClient.On("CreatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpoint", func(t *testing.T) {
        input := &elasticsearchservice.CreateVpcEndpointInput{}
        output := &elasticsearchservice.CreateVpcEndpointOutput{}

        mockClient.On("CreateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteElasticsearchDomain", func(t *testing.T) {
        input := &elasticsearchservice.DeleteElasticsearchDomainInput{}
        output := &elasticsearchservice.DeleteElasticsearchDomainOutput{}

        mockClient.On("DeleteElasticsearchDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteElasticsearchDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteElasticsearchServiceRole", func(t *testing.T) {
        input := &elasticsearchservice.DeleteElasticsearchServiceRoleInput{}
        output := &elasticsearchservice.DeleteElasticsearchServiceRoleOutput{}

        mockClient.On("DeleteElasticsearchServiceRole", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteElasticsearchServiceRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInboundCrossClusterSearchConnection", func(t *testing.T) {
        input := &elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput{}
        output := &elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput{}

        mockClient.On("DeleteInboundCrossClusterSearchConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInboundCrossClusterSearchConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOutboundCrossClusterSearchConnection", func(t *testing.T) {
        input := &elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput{}
        output := &elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput{}

        mockClient.On("DeleteOutboundCrossClusterSearchConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOutboundCrossClusterSearchConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackage", func(t *testing.T) {
        input := &elasticsearchservice.DeletePackageInput{}
        output := &elasticsearchservice.DeletePackageOutput{}

        mockClient.On("DeletePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpoint", func(t *testing.T) {
        input := &elasticsearchservice.DeleteVpcEndpointInput{}
        output := &elasticsearchservice.DeleteVpcEndpointOutput{}

        mockClient.On("DeleteVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainAutoTunes", func(t *testing.T) {
        input := &elasticsearchservice.DescribeDomainAutoTunesInput{}
        output := &elasticsearchservice.DescribeDomainAutoTunesOutput{}

        mockClient.On("DescribeDomainAutoTunes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainAutoTunes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainChangeProgress", func(t *testing.T) {
        input := &elasticsearchservice.DescribeDomainChangeProgressInput{}
        output := &elasticsearchservice.DescribeDomainChangeProgressOutput{}

        mockClient.On("DescribeDomainChangeProgress", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainChangeProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticsearchDomain", func(t *testing.T) {
        input := &elasticsearchservice.DescribeElasticsearchDomainInput{}
        output := &elasticsearchservice.DescribeElasticsearchDomainOutput{}

        mockClient.On("DescribeElasticsearchDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticsearchDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticsearchDomainConfig", func(t *testing.T) {
        input := &elasticsearchservice.DescribeElasticsearchDomainConfigInput{}
        output := &elasticsearchservice.DescribeElasticsearchDomainConfigOutput{}

        mockClient.On("DescribeElasticsearchDomainConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticsearchDomainConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticsearchDomains", func(t *testing.T) {
        input := &elasticsearchservice.DescribeElasticsearchDomainsInput{}
        output := &elasticsearchservice.DescribeElasticsearchDomainsOutput{}

        mockClient.On("DescribeElasticsearchDomains", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticsearchDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeElasticsearchInstanceTypeLimits", func(t *testing.T) {
        input := &elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput{}
        output := &elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput{}

        mockClient.On("DescribeElasticsearchInstanceTypeLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeElasticsearchInstanceTypeLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInboundCrossClusterSearchConnections", func(t *testing.T) {
        input := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput{}
        output := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput{}

        mockClient.On("DescribeInboundCrossClusterSearchConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInboundCrossClusterSearchConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOutboundCrossClusterSearchConnections", func(t *testing.T) {
        input := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput{}
        output := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput{}

        mockClient.On("DescribeOutboundCrossClusterSearchConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOutboundCrossClusterSearchConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackages", func(t *testing.T) {
        input := &elasticsearchservice.DescribePackagesInput{}
        output := &elasticsearchservice.DescribePackagesOutput{}

        mockClient.On("DescribePackages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedElasticsearchInstanceOfferings", func(t *testing.T) {
        input := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput{}
        output := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput{}

        mockClient.On("DescribeReservedElasticsearchInstanceOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedElasticsearchInstanceOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedElasticsearchInstances", func(t *testing.T) {
        input := &elasticsearchservice.DescribeReservedElasticsearchInstancesInput{}
        output := &elasticsearchservice.DescribeReservedElasticsearchInstancesOutput{}

        mockClient.On("DescribeReservedElasticsearchInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedElasticsearchInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpoints", func(t *testing.T) {
        input := &elasticsearchservice.DescribeVpcEndpointsInput{}
        output := &elasticsearchservice.DescribeVpcEndpointsOutput{}

        mockClient.On("DescribeVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDissociatePackage", func(t *testing.T) {
        input := &elasticsearchservice.DissociatePackageInput{}
        output := &elasticsearchservice.DissociatePackageOutput{}

        mockClient.On("DissociatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DissociatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCompatibleElasticsearchVersions", func(t *testing.T) {
        input := &elasticsearchservice.GetCompatibleElasticsearchVersionsInput{}
        output := &elasticsearchservice.GetCompatibleElasticsearchVersionsOutput{}

        mockClient.On("GetCompatibleElasticsearchVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetCompatibleElasticsearchVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageVersionHistory", func(t *testing.T) {
        input := &elasticsearchservice.GetPackageVersionHistoryInput{}
        output := &elasticsearchservice.GetPackageVersionHistoryOutput{}

        mockClient.On("GetPackageVersionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageVersionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUpgradeHistory", func(t *testing.T) {
        input := &elasticsearchservice.GetUpgradeHistoryInput{}
        output := &elasticsearchservice.GetUpgradeHistoryOutput{}

        mockClient.On("GetUpgradeHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetUpgradeHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUpgradeStatus", func(t *testing.T) {
        input := &elasticsearchservice.GetUpgradeStatusInput{}
        output := &elasticsearchservice.GetUpgradeStatusOutput{}

        mockClient.On("GetUpgradeStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetUpgradeStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainNames", func(t *testing.T) {
        input := &elasticsearchservice.ListDomainNamesInput{}
        output := &elasticsearchservice.ListDomainNamesOutput{}

        mockClient.On("ListDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainsForPackage", func(t *testing.T) {
        input := &elasticsearchservice.ListDomainsForPackageInput{}
        output := &elasticsearchservice.ListDomainsForPackageOutput{}

        mockClient.On("ListDomainsForPackage", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainsForPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListElasticsearchInstanceTypes", func(t *testing.T) {
        input := &elasticsearchservice.ListElasticsearchInstanceTypesInput{}
        output := &elasticsearchservice.ListElasticsearchInstanceTypesOutput{}

        mockClient.On("ListElasticsearchInstanceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListElasticsearchInstanceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListElasticsearchVersions", func(t *testing.T) {
        input := &elasticsearchservice.ListElasticsearchVersionsInput{}
        output := &elasticsearchservice.ListElasticsearchVersionsOutput{}

        mockClient.On("ListElasticsearchVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListElasticsearchVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackagesForDomain", func(t *testing.T) {
        input := &elasticsearchservice.ListPackagesForDomainInput{}
        output := &elasticsearchservice.ListPackagesForDomainOutput{}

        mockClient.On("ListPackagesForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackagesForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &elasticsearchservice.ListTagsInput{}
        output := &elasticsearchservice.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpointAccess", func(t *testing.T) {
        input := &elasticsearchservice.ListVpcEndpointAccessInput{}
        output := &elasticsearchservice.ListVpcEndpointAccessOutput{}

        mockClient.On("ListVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpoints", func(t *testing.T) {
        input := &elasticsearchservice.ListVpcEndpointsInput{}
        output := &elasticsearchservice.ListVpcEndpointsOutput{}

        mockClient.On("ListVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpointsForDomain", func(t *testing.T) {
        input := &elasticsearchservice.ListVpcEndpointsForDomainInput{}
        output := &elasticsearchservice.ListVpcEndpointsForDomainOutput{}

        mockClient.On("ListVpcEndpointsForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpointsForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedElasticsearchInstanceOffering", func(t *testing.T) {
        input := &elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput{}
        output := &elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput{}

        mockClient.On("PurchaseReservedElasticsearchInstanceOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedElasticsearchInstanceOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectInboundCrossClusterSearchConnection", func(t *testing.T) {
        input := &elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput{}
        output := &elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput{}

        mockClient.On("RejectInboundCrossClusterSearchConnection", ctx, input).Return(output, nil)

        result, err := mockClient.RejectInboundCrossClusterSearchConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &elasticsearchservice.RemoveTagsInput{}
        output := &elasticsearchservice.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeVpcEndpointAccess", func(t *testing.T) {
        input := &elasticsearchservice.RevokeVpcEndpointAccessInput{}
        output := &elasticsearchservice.RevokeVpcEndpointAccessOutput{}

        mockClient.On("RevokeVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartElasticsearchServiceSoftwareUpdate", func(t *testing.T) {
        input := &elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput{}
        output := &elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput{}

        mockClient.On("StartElasticsearchServiceSoftwareUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartElasticsearchServiceSoftwareUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateElasticsearchDomainConfig", func(t *testing.T) {
        input := &elasticsearchservice.UpdateElasticsearchDomainConfigInput{}
        output := &elasticsearchservice.UpdateElasticsearchDomainConfigOutput{}

        mockClient.On("UpdateElasticsearchDomainConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateElasticsearchDomainConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackage", func(t *testing.T) {
        input := &elasticsearchservice.UpdatePackageInput{}
        output := &elasticsearchservice.UpdatePackageOutput{}

        mockClient.On("UpdatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcEndpoint", func(t *testing.T) {
        input := &elasticsearchservice.UpdateVpcEndpointInput{}
        output := &elasticsearchservice.UpdateVpcEndpointOutput{}

        mockClient.On("UpdateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeElasticsearchDomain", func(t *testing.T) {
        input := &elasticsearchservice.UpgradeElasticsearchDomainInput{}
        output := &elasticsearchservice.UpgradeElasticsearchDomainOutput{}

        mockClient.On("UpgradeElasticsearchDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeElasticsearchDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
