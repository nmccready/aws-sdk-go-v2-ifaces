package opensearch_test

// tests for the opensearch service interface for this ../../../service/opensearch/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/opensearch"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opensearch/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opensearch/opensearch_iface"
	"github.com/stretchr/testify/assert"
)

func TestOpensearchServiceCanBeMocked(t *testing.T) {
	var iface opensearch_iface.IClient
	iface = &opensearch.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := opensearch.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInboundConnection", func(t *testing.T) {
        input := &opensearch.AcceptInboundConnectionInput{}
        output := &opensearch.AcceptInboundConnectionOutput{}

        mockClient.On("AcceptInboundConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInboundConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddDataSource", func(t *testing.T) {
        input := &opensearch.AddDataSourceInput{}
        output := &opensearch.AddDataSourceOutput{}

        mockClient.On("AddDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.AddDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &opensearch.AddTagsInput{}
        output := &opensearch.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePackage", func(t *testing.T) {
        input := &opensearch.AssociatePackageInput{}
        output := &opensearch.AssociatePackageOutput{}

        mockClient.On("AssociatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeVpcEndpointAccess", func(t *testing.T) {
        input := &opensearch.AuthorizeVpcEndpointAccessInput{}
        output := &opensearch.AuthorizeVpcEndpointAccessOutput{}

        mockClient.On("AuthorizeVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDomainConfigChange", func(t *testing.T) {
        input := &opensearch.CancelDomainConfigChangeInput{}
        output := &opensearch.CancelDomainConfigChangeOutput{}

        mockClient.On("CancelDomainConfigChange", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDomainConfigChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelServiceSoftwareUpdate", func(t *testing.T) {
        input := &opensearch.CancelServiceSoftwareUpdateInput{}
        output := &opensearch.CancelServiceSoftwareUpdateOutput{}

        mockClient.On("CancelServiceSoftwareUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.CancelServiceSoftwareUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &opensearch.CreateDomainInput{}
        output := &opensearch.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOutboundConnection", func(t *testing.T) {
        input := &opensearch.CreateOutboundConnectionInput{}
        output := &opensearch.CreateOutboundConnectionOutput{}

        mockClient.On("CreateOutboundConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOutboundConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackage", func(t *testing.T) {
        input := &opensearch.CreatePackageInput{}
        output := &opensearch.CreatePackageOutput{}

        mockClient.On("CreatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpoint", func(t *testing.T) {
        input := &opensearch.CreateVpcEndpointInput{}
        output := &opensearch.CreateVpcEndpointOutput{}

        mockClient.On("CreateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &opensearch.DeleteDataSourceInput{}
        output := &opensearch.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &opensearch.DeleteDomainInput{}
        output := &opensearch.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInboundConnection", func(t *testing.T) {
        input := &opensearch.DeleteInboundConnectionInput{}
        output := &opensearch.DeleteInboundConnectionOutput{}

        mockClient.On("DeleteInboundConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInboundConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOutboundConnection", func(t *testing.T) {
        input := &opensearch.DeleteOutboundConnectionInput{}
        output := &opensearch.DeleteOutboundConnectionOutput{}

        mockClient.On("DeleteOutboundConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOutboundConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackage", func(t *testing.T) {
        input := &opensearch.DeletePackageInput{}
        output := &opensearch.DeletePackageOutput{}

        mockClient.On("DeletePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpoint", func(t *testing.T) {
        input := &opensearch.DeleteVpcEndpointInput{}
        output := &opensearch.DeleteVpcEndpointOutput{}

        mockClient.On("DeleteVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &opensearch.DescribeDomainInput{}
        output := &opensearch.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainAutoTunes", func(t *testing.T) {
        input := &opensearch.DescribeDomainAutoTunesInput{}
        output := &opensearch.DescribeDomainAutoTunesOutput{}

        mockClient.On("DescribeDomainAutoTunes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainAutoTunes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainChangeProgress", func(t *testing.T) {
        input := &opensearch.DescribeDomainChangeProgressInput{}
        output := &opensearch.DescribeDomainChangeProgressOutput{}

        mockClient.On("DescribeDomainChangeProgress", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainChangeProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainConfig", func(t *testing.T) {
        input := &opensearch.DescribeDomainConfigInput{}
        output := &opensearch.DescribeDomainConfigOutput{}

        mockClient.On("DescribeDomainConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainHealth", func(t *testing.T) {
        input := &opensearch.DescribeDomainHealthInput{}
        output := &opensearch.DescribeDomainHealthOutput{}

        mockClient.On("DescribeDomainHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainNodes", func(t *testing.T) {
        input := &opensearch.DescribeDomainNodesInput{}
        output := &opensearch.DescribeDomainNodesOutput{}

        mockClient.On("DescribeDomainNodes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomains", func(t *testing.T) {
        input := &opensearch.DescribeDomainsInput{}
        output := &opensearch.DescribeDomainsOutput{}

        mockClient.On("DescribeDomains", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDryRunProgress", func(t *testing.T) {
        input := &opensearch.DescribeDryRunProgressInput{}
        output := &opensearch.DescribeDryRunProgressOutput{}

        mockClient.On("DescribeDryRunProgress", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDryRunProgress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInboundConnections", func(t *testing.T) {
        input := &opensearch.DescribeInboundConnectionsInput{}
        output := &opensearch.DescribeInboundConnectionsOutput{}

        mockClient.On("DescribeInboundConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInboundConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceTypeLimits", func(t *testing.T) {
        input := &opensearch.DescribeInstanceTypeLimitsInput{}
        output := &opensearch.DescribeInstanceTypeLimitsOutput{}

        mockClient.On("DescribeInstanceTypeLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceTypeLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOutboundConnections", func(t *testing.T) {
        input := &opensearch.DescribeOutboundConnectionsInput{}
        output := &opensearch.DescribeOutboundConnectionsOutput{}

        mockClient.On("DescribeOutboundConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOutboundConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackages", func(t *testing.T) {
        input := &opensearch.DescribePackagesInput{}
        output := &opensearch.DescribePackagesOutput{}

        mockClient.On("DescribePackages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstanceOfferings", func(t *testing.T) {
        input := &opensearch.DescribeReservedInstanceOfferingsInput{}
        output := &opensearch.DescribeReservedInstanceOfferingsOutput{}

        mockClient.On("DescribeReservedInstanceOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstanceOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedInstances", func(t *testing.T) {
        input := &opensearch.DescribeReservedInstancesInput{}
        output := &opensearch.DescribeReservedInstancesOutput{}

        mockClient.On("DescribeReservedInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcEndpoints", func(t *testing.T) {
        input := &opensearch.DescribeVpcEndpointsInput{}
        output := &opensearch.DescribeVpcEndpointsOutput{}

        mockClient.On("DescribeVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDissociatePackage", func(t *testing.T) {
        input := &opensearch.DissociatePackageInput{}
        output := &opensearch.DissociatePackageOutput{}

        mockClient.On("DissociatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DissociatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCompatibleVersions", func(t *testing.T) {
        input := &opensearch.GetCompatibleVersionsInput{}
        output := &opensearch.GetCompatibleVersionsOutput{}

        mockClient.On("GetCompatibleVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetCompatibleVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &opensearch.GetDataSourceInput{}
        output := &opensearch.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainMaintenanceStatus", func(t *testing.T) {
        input := &opensearch.GetDomainMaintenanceStatusInput{}
        output := &opensearch.GetDomainMaintenanceStatusOutput{}

        mockClient.On("GetDomainMaintenanceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainMaintenanceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageVersionHistory", func(t *testing.T) {
        input := &opensearch.GetPackageVersionHistoryInput{}
        output := &opensearch.GetPackageVersionHistoryOutput{}

        mockClient.On("GetPackageVersionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageVersionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUpgradeHistory", func(t *testing.T) {
        input := &opensearch.GetUpgradeHistoryInput{}
        output := &opensearch.GetUpgradeHistoryOutput{}

        mockClient.On("GetUpgradeHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetUpgradeHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUpgradeStatus", func(t *testing.T) {
        input := &opensearch.GetUpgradeStatusInput{}
        output := &opensearch.GetUpgradeStatusOutput{}

        mockClient.On("GetUpgradeStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetUpgradeStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &opensearch.ListDataSourcesInput{}
        output := &opensearch.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainMaintenances", func(t *testing.T) {
        input := &opensearch.ListDomainMaintenancesInput{}
        output := &opensearch.ListDomainMaintenancesOutput{}

        mockClient.On("ListDomainMaintenances", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainMaintenances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainNames", func(t *testing.T) {
        input := &opensearch.ListDomainNamesInput{}
        output := &opensearch.ListDomainNamesOutput{}

        mockClient.On("ListDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainsForPackage", func(t *testing.T) {
        input := &opensearch.ListDomainsForPackageInput{}
        output := &opensearch.ListDomainsForPackageOutput{}

        mockClient.On("ListDomainsForPackage", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainsForPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceTypeDetails", func(t *testing.T) {
        input := &opensearch.ListInstanceTypeDetailsInput{}
        output := &opensearch.ListInstanceTypeDetailsOutput{}

        mockClient.On("ListInstanceTypeDetails", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceTypeDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackagesForDomain", func(t *testing.T) {
        input := &opensearch.ListPackagesForDomainInput{}
        output := &opensearch.ListPackagesForDomainOutput{}

        mockClient.On("ListPackagesForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackagesForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScheduledActions", func(t *testing.T) {
        input := &opensearch.ListScheduledActionsInput{}
        output := &opensearch.ListScheduledActionsOutput{}

        mockClient.On("ListScheduledActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListScheduledActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &opensearch.ListTagsInput{}
        output := &opensearch.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVersions", func(t *testing.T) {
        input := &opensearch.ListVersionsInput{}
        output := &opensearch.ListVersionsOutput{}

        mockClient.On("ListVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpointAccess", func(t *testing.T) {
        input := &opensearch.ListVpcEndpointAccessInput{}
        output := &opensearch.ListVpcEndpointAccessOutput{}

        mockClient.On("ListVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpoints", func(t *testing.T) {
        input := &opensearch.ListVpcEndpointsInput{}
        output := &opensearch.ListVpcEndpointsOutput{}

        mockClient.On("ListVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpointsForDomain", func(t *testing.T) {
        input := &opensearch.ListVpcEndpointsForDomainInput{}
        output := &opensearch.ListVpcEndpointsForDomainOutput{}

        mockClient.On("ListVpcEndpointsForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpointsForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedInstanceOffering", func(t *testing.T) {
        input := &opensearch.PurchaseReservedInstanceOfferingInput{}
        output := &opensearch.PurchaseReservedInstanceOfferingOutput{}

        mockClient.On("PurchaseReservedInstanceOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedInstanceOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectInboundConnection", func(t *testing.T) {
        input := &opensearch.RejectInboundConnectionInput{}
        output := &opensearch.RejectInboundConnectionOutput{}

        mockClient.On("RejectInboundConnection", ctx, input).Return(output, nil)

        result, err := mockClient.RejectInboundConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &opensearch.RemoveTagsInput{}
        output := &opensearch.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeVpcEndpointAccess", func(t *testing.T) {
        input := &opensearch.RevokeVpcEndpointAccessInput{}
        output := &opensearch.RevokeVpcEndpointAccessOutput{}

        mockClient.On("RevokeVpcEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeVpcEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDomainMaintenance", func(t *testing.T) {
        input := &opensearch.StartDomainMaintenanceInput{}
        output := &opensearch.StartDomainMaintenanceOutput{}

        mockClient.On("StartDomainMaintenance", ctx, input).Return(output, nil)

        result, err := mockClient.StartDomainMaintenance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartServiceSoftwareUpdate", func(t *testing.T) {
        input := &opensearch.StartServiceSoftwareUpdateInput{}
        output := &opensearch.StartServiceSoftwareUpdateOutput{}

        mockClient.On("StartServiceSoftwareUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.StartServiceSoftwareUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &opensearch.UpdateDataSourceInput{}
        output := &opensearch.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainConfig", func(t *testing.T) {
        input := &opensearch.UpdateDomainConfigInput{}
        output := &opensearch.UpdateDomainConfigOutput{}

        mockClient.On("UpdateDomainConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackage", func(t *testing.T) {
        input := &opensearch.UpdatePackageInput{}
        output := &opensearch.UpdatePackageOutput{}

        mockClient.On("UpdatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScheduledAction", func(t *testing.T) {
        input := &opensearch.UpdateScheduledActionInput{}
        output := &opensearch.UpdateScheduledActionOutput{}

        mockClient.On("UpdateScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcEndpoint", func(t *testing.T) {
        input := &opensearch.UpdateVpcEndpointInput{}
        output := &opensearch.UpdateVpcEndpointOutput{}

        mockClient.On("UpdateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeDomain", func(t *testing.T) {
        input := &opensearch.UpgradeDomainInput{}
        output := &opensearch.UpgradeDomainOutput{}

        mockClient.On("UpgradeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
