package codestarconnections_test

// tests for the codestarconnections service interface for this ../../../service/codestarconnections/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestarconnections/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestarconnections/codestarconnections_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodestarconnectionsServiceCanBeMocked(t *testing.T) {
	var iface codestarconnections_iface.IClient
	iface = &codestarconnections.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codestarconnections.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &codestarconnections.CreateConnectionInput{}
        output := &codestarconnections.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHost", func(t *testing.T) {
        input := &codestarconnections.CreateHostInput{}
        output := &codestarconnections.CreateHostOutput{}

        mockClient.On("CreateHost", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepositoryLink", func(t *testing.T) {
        input := &codestarconnections.CreateRepositoryLinkInput{}
        output := &codestarconnections.CreateRepositoryLinkOutput{}

        mockClient.On("CreateRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSyncConfiguration", func(t *testing.T) {
        input := &codestarconnections.CreateSyncConfigurationInput{}
        output := &codestarconnections.CreateSyncConfigurationOutput{}

        mockClient.On("CreateSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &codestarconnections.DeleteConnectionInput{}
        output := &codestarconnections.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHost", func(t *testing.T) {
        input := &codestarconnections.DeleteHostInput{}
        output := &codestarconnections.DeleteHostOutput{}

        mockClient.On("DeleteHost", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepositoryLink", func(t *testing.T) {
        input := &codestarconnections.DeleteRepositoryLinkInput{}
        output := &codestarconnections.DeleteRepositoryLinkOutput{}

        mockClient.On("DeleteRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSyncConfiguration", func(t *testing.T) {
        input := &codestarconnections.DeleteSyncConfigurationInput{}
        output := &codestarconnections.DeleteSyncConfigurationOutput{}

        mockClient.On("DeleteSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnection", func(t *testing.T) {
        input := &codestarconnections.GetConnectionInput{}
        output := &codestarconnections.GetConnectionOutput{}

        mockClient.On("GetConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHost", func(t *testing.T) {
        input := &codestarconnections.GetHostInput{}
        output := &codestarconnections.GetHostOutput{}

        mockClient.On("GetHost", ctx, input).Return(output, nil)

        result, err := mockClient.GetHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryLink", func(t *testing.T) {
        input := &codestarconnections.GetRepositoryLinkInput{}
        output := &codestarconnections.GetRepositoryLinkOutput{}

        mockClient.On("GetRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositorySyncStatus", func(t *testing.T) {
        input := &codestarconnections.GetRepositorySyncStatusInput{}
        output := &codestarconnections.GetRepositorySyncStatusOutput{}

        mockClient.On("GetRepositorySyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositorySyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceSyncStatus", func(t *testing.T) {
        input := &codestarconnections.GetResourceSyncStatusInput{}
        output := &codestarconnections.GetResourceSyncStatusOutput{}

        mockClient.On("GetResourceSyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceSyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSyncBlockerSummary", func(t *testing.T) {
        input := &codestarconnections.GetSyncBlockerSummaryInput{}
        output := &codestarconnections.GetSyncBlockerSummaryOutput{}

        mockClient.On("GetSyncBlockerSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetSyncBlockerSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSyncConfiguration", func(t *testing.T) {
        input := &codestarconnections.GetSyncConfigurationInput{}
        output := &codestarconnections.GetSyncConfigurationOutput{}

        mockClient.On("GetSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnections", func(t *testing.T) {
        input := &codestarconnections.ListConnectionsInput{}
        output := &codestarconnections.ListConnectionsOutput{}

        mockClient.On("ListConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHosts", func(t *testing.T) {
        input := &codestarconnections.ListHostsInput{}
        output := &codestarconnections.ListHostsOutput{}

        mockClient.On("ListHosts", ctx, input).Return(output, nil)

        result, err := mockClient.ListHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositoryLinks", func(t *testing.T) {
        input := &codestarconnections.ListRepositoryLinksInput{}
        output := &codestarconnections.ListRepositoryLinksOutput{}

        mockClient.On("ListRepositoryLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositoryLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositorySyncDefinitions", func(t *testing.T) {
        input := &codestarconnections.ListRepositorySyncDefinitionsInput{}
        output := &codestarconnections.ListRepositorySyncDefinitionsOutput{}

        mockClient.On("ListRepositorySyncDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositorySyncDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSyncConfigurations", func(t *testing.T) {
        input := &codestarconnections.ListSyncConfigurationsInput{}
        output := &codestarconnections.ListSyncConfigurationsOutput{}

        mockClient.On("ListSyncConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSyncConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codestarconnections.ListTagsForResourceInput{}
        output := &codestarconnections.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codestarconnections.TagResourceInput{}
        output := &codestarconnections.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codestarconnections.UntagResourceInput{}
        output := &codestarconnections.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHost", func(t *testing.T) {
        input := &codestarconnections.UpdateHostInput{}
        output := &codestarconnections.UpdateHostOutput{}

        mockClient.On("UpdateHost", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepositoryLink", func(t *testing.T) {
        input := &codestarconnections.UpdateRepositoryLinkInput{}
        output := &codestarconnections.UpdateRepositoryLinkOutput{}

        mockClient.On("UpdateRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSyncBlocker", func(t *testing.T) {
        input := &codestarconnections.UpdateSyncBlockerInput{}
        output := &codestarconnections.UpdateSyncBlockerOutput{}

        mockClient.On("UpdateSyncBlocker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSyncBlocker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSyncConfiguration", func(t *testing.T) {
        input := &codestarconnections.UpdateSyncConfigurationInput{}
        output := &codestarconnections.UpdateSyncConfigurationOutput{}

        mockClient.On("UpdateSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
