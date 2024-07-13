package codeartifact_test

// tests for the codeartifact service interface for this ../../../service/codeartifact/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeartifact/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeartifact/codeartifact_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodeartifactServiceCanBeMocked(t *testing.T) {
	var iface codeartifact_iface.IClient
	iface = &codeartifact.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codeartifact.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateExternalConnection", func(t *testing.T) {
        input := &codeartifact.AssociateExternalConnectionInput{}
        output := &codeartifact.AssociateExternalConnectionOutput{}

        mockClient.On("AssociateExternalConnection", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateExternalConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyPackageVersions", func(t *testing.T) {
        input := &codeartifact.CopyPackageVersionsInput{}
        output := &codeartifact.CopyPackageVersionsOutput{}

        mockClient.On("CopyPackageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.CopyPackageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &codeartifact.CreateDomainInput{}
        output := &codeartifact.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackageGroup", func(t *testing.T) {
        input := &codeartifact.CreatePackageGroupInput{}
        output := &codeartifact.CreatePackageGroupOutput{}

        mockClient.On("CreatePackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepository", func(t *testing.T) {
        input := &codeartifact.CreateRepositoryInput{}
        output := &codeartifact.CreateRepositoryOutput{}

        mockClient.On("CreateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &codeartifact.DeleteDomainInput{}
        output := &codeartifact.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.DeleteDomainPermissionsPolicyInput{}
        output := &codeartifact.DeleteDomainPermissionsPolicyOutput{}

        mockClient.On("DeleteDomainPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackage", func(t *testing.T) {
        input := &codeartifact.DeletePackageInput{}
        output := &codeartifact.DeletePackageOutput{}

        mockClient.On("DeletePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackageGroup", func(t *testing.T) {
        input := &codeartifact.DeletePackageGroupInput{}
        output := &codeartifact.DeletePackageGroupOutput{}

        mockClient.On("DeletePackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackageVersions", func(t *testing.T) {
        input := &codeartifact.DeletePackageVersionsInput{}
        output := &codeartifact.DeletePackageVersionsOutput{}

        mockClient.On("DeletePackageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepository", func(t *testing.T) {
        input := &codeartifact.DeleteRepositoryInput{}
        output := &codeartifact.DeleteRepositoryOutput{}

        mockClient.On("DeleteRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepositoryPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.DeleteRepositoryPermissionsPolicyInput{}
        output := &codeartifact.DeleteRepositoryPermissionsPolicyOutput{}

        mockClient.On("DeleteRepositoryPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepositoryPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &codeartifact.DescribeDomainInput{}
        output := &codeartifact.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackage", func(t *testing.T) {
        input := &codeartifact.DescribePackageInput{}
        output := &codeartifact.DescribePackageOutput{}

        mockClient.On("DescribePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackageGroup", func(t *testing.T) {
        input := &codeartifact.DescribePackageGroupInput{}
        output := &codeartifact.DescribePackageGroupOutput{}

        mockClient.On("DescribePackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackageVersion", func(t *testing.T) {
        input := &codeartifact.DescribePackageVersionInput{}
        output := &codeartifact.DescribePackageVersionOutput{}

        mockClient.On("DescribePackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRepository", func(t *testing.T) {
        input := &codeartifact.DescribeRepositoryInput{}
        output := &codeartifact.DescribeRepositoryOutput{}

        mockClient.On("DescribeRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateExternalConnection", func(t *testing.T) {
        input := &codeartifact.DisassociateExternalConnectionInput{}
        output := &codeartifact.DisassociateExternalConnectionOutput{}

        mockClient.On("DisassociateExternalConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateExternalConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisposePackageVersions", func(t *testing.T) {
        input := &codeartifact.DisposePackageVersionsInput{}
        output := &codeartifact.DisposePackageVersionsOutput{}

        mockClient.On("DisposePackageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DisposePackageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociatedPackageGroup", func(t *testing.T) {
        input := &codeartifact.GetAssociatedPackageGroupInput{}
        output := &codeartifact.GetAssociatedPackageGroupOutput{}

        mockClient.On("GetAssociatedPackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociatedPackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizationToken", func(t *testing.T) {
        input := &codeartifact.GetAuthorizationTokenInput{}
        output := &codeartifact.GetAuthorizationTokenOutput{}

        mockClient.On("GetAuthorizationToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizationToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.GetDomainPermissionsPolicyInput{}
        output := &codeartifact.GetDomainPermissionsPolicyOutput{}

        mockClient.On("GetDomainPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageVersionAsset", func(t *testing.T) {
        input := &codeartifact.GetPackageVersionAssetInput{}
        output := &codeartifact.GetPackageVersionAssetOutput{}

        mockClient.On("GetPackageVersionAsset", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageVersionAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageVersionReadme", func(t *testing.T) {
        input := &codeartifact.GetPackageVersionReadmeInput{}
        output := &codeartifact.GetPackageVersionReadmeOutput{}

        mockClient.On("GetPackageVersionReadme", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageVersionReadme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryEndpoint", func(t *testing.T) {
        input := &codeartifact.GetRepositoryEndpointInput{}
        output := &codeartifact.GetRepositoryEndpointOutput{}

        mockClient.On("GetRepositoryEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.GetRepositoryPermissionsPolicyInput{}
        output := &codeartifact.GetRepositoryPermissionsPolicyOutput{}

        mockClient.On("GetRepositoryPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAllowedRepositoriesForGroup", func(t *testing.T) {
        input := &codeartifact.ListAllowedRepositoriesForGroupInput{}
        output := &codeartifact.ListAllowedRepositoriesForGroupOutput{}

        mockClient.On("ListAllowedRepositoriesForGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListAllowedRepositoriesForGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedPackages", func(t *testing.T) {
        input := &codeartifact.ListAssociatedPackagesInput{}
        output := &codeartifact.ListAssociatedPackagesOutput{}

        mockClient.On("ListAssociatedPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &codeartifact.ListDomainsInput{}
        output := &codeartifact.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageGroups", func(t *testing.T) {
        input := &codeartifact.ListPackageGroupsInput{}
        output := &codeartifact.ListPackageGroupsOutput{}

        mockClient.On("ListPackageGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageVersionAssets", func(t *testing.T) {
        input := &codeartifact.ListPackageVersionAssetsInput{}
        output := &codeartifact.ListPackageVersionAssetsOutput{}

        mockClient.On("ListPackageVersionAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageVersionAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageVersionDependencies", func(t *testing.T) {
        input := &codeartifact.ListPackageVersionDependenciesInput{}
        output := &codeartifact.ListPackageVersionDependenciesOutput{}

        mockClient.On("ListPackageVersionDependencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageVersionDependencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageVersions", func(t *testing.T) {
        input := &codeartifact.ListPackageVersionsInput{}
        output := &codeartifact.ListPackageVersionsOutput{}

        mockClient.On("ListPackageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackages", func(t *testing.T) {
        input := &codeartifact.ListPackagesInput{}
        output := &codeartifact.ListPackagesOutput{}

        mockClient.On("ListPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositories", func(t *testing.T) {
        input := &codeartifact.ListRepositoriesInput{}
        output := &codeartifact.ListRepositoriesOutput{}

        mockClient.On("ListRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositoriesInDomain", func(t *testing.T) {
        input := &codeartifact.ListRepositoriesInDomainInput{}
        output := &codeartifact.ListRepositoriesInDomainOutput{}

        mockClient.On("ListRepositoriesInDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositoriesInDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubPackageGroups", func(t *testing.T) {
        input := &codeartifact.ListSubPackageGroupsInput{}
        output := &codeartifact.ListSubPackageGroupsOutput{}

        mockClient.On("ListSubPackageGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubPackageGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codeartifact.ListTagsForResourceInput{}
        output := &codeartifact.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishPackageVersion", func(t *testing.T) {
        input := &codeartifact.PublishPackageVersionInput{}
        output := &codeartifact.PublishPackageVersionOutput{}

        mockClient.On("PublishPackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.PublishPackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDomainPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.PutDomainPermissionsPolicyInput{}
        output := &codeartifact.PutDomainPermissionsPolicyOutput{}

        mockClient.On("PutDomainPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutDomainPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPackageOriginConfiguration", func(t *testing.T) {
        input := &codeartifact.PutPackageOriginConfigurationInput{}
        output := &codeartifact.PutPackageOriginConfigurationOutput{}

        mockClient.On("PutPackageOriginConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutPackageOriginConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRepositoryPermissionsPolicy", func(t *testing.T) {
        input := &codeartifact.PutRepositoryPermissionsPolicyInput{}
        output := &codeartifact.PutRepositoryPermissionsPolicyOutput{}

        mockClient.On("PutRepositoryPermissionsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutRepositoryPermissionsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codeartifact.TagResourceInput{}
        output := &codeartifact.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codeartifact.UntagResourceInput{}
        output := &codeartifact.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackageGroup", func(t *testing.T) {
        input := &codeartifact.UpdatePackageGroupInput{}
        output := &codeartifact.UpdatePackageGroupOutput{}

        mockClient.On("UpdatePackageGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackageGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackageGroupOriginConfiguration", func(t *testing.T) {
        input := &codeartifact.UpdatePackageGroupOriginConfigurationInput{}
        output := &codeartifact.UpdatePackageGroupOriginConfigurationOutput{}

        mockClient.On("UpdatePackageGroupOriginConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackageGroupOriginConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackageVersionsStatus", func(t *testing.T) {
        input := &codeartifact.UpdatePackageVersionsStatusInput{}
        output := &codeartifact.UpdatePackageVersionsStatusOutput{}

        mockClient.On("UpdatePackageVersionsStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackageVersionsStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepository", func(t *testing.T) {
        input := &codeartifact.UpdateRepositoryInput{}
        output := &codeartifact.UpdateRepositoryOutput{}

        mockClient.On("UpdateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
