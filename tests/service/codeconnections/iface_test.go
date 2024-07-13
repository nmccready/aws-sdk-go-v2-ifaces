package codeconnections_test

// tests for the codeconnections service interface for this ../../../service/codeconnections/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeconnections"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeconnections/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeconnections/codeconnections_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodeconnectionsServiceCanBeMocked(t *testing.T) {
	var iface codeconnections_iface.IClient
	iface = &codeconnections.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codeconnections.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnection", func(t *testing.T) {
        input := &codeconnections.CreateConnectionInput{}
        output := &codeconnections.CreateConnectionOutput{}

        mockClient.On("CreateConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHost", func(t *testing.T) {
        input := &codeconnections.CreateHostInput{}
        output := &codeconnections.CreateHostOutput{}

        mockClient.On("CreateHost", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepositoryLink", func(t *testing.T) {
        input := &codeconnections.CreateRepositoryLinkInput{}
        output := &codeconnections.CreateRepositoryLinkOutput{}

        mockClient.On("CreateRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSyncConfiguration", func(t *testing.T) {
        input := &codeconnections.CreateSyncConfigurationInput{}
        output := &codeconnections.CreateSyncConfigurationOutput{}

        mockClient.On("CreateSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnection", func(t *testing.T) {
        input := &codeconnections.DeleteConnectionInput{}
        output := &codeconnections.DeleteConnectionOutput{}

        mockClient.On("DeleteConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHost", func(t *testing.T) {
        input := &codeconnections.DeleteHostInput{}
        output := &codeconnections.DeleteHostOutput{}

        mockClient.On("DeleteHost", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepositoryLink", func(t *testing.T) {
        input := &codeconnections.DeleteRepositoryLinkInput{}
        output := &codeconnections.DeleteRepositoryLinkOutput{}

        mockClient.On("DeleteRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSyncConfiguration", func(t *testing.T) {
        input := &codeconnections.DeleteSyncConfigurationInput{}
        output := &codeconnections.DeleteSyncConfigurationOutput{}

        mockClient.On("DeleteSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnection", func(t *testing.T) {
        input := &codeconnections.GetConnectionInput{}
        output := &codeconnections.GetConnectionOutput{}

        mockClient.On("GetConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHost", func(t *testing.T) {
        input := &codeconnections.GetHostInput{}
        output := &codeconnections.GetHostOutput{}

        mockClient.On("GetHost", ctx, input).Return(output, nil)

        result, err := mockClient.GetHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryLink", func(t *testing.T) {
        input := &codeconnections.GetRepositoryLinkInput{}
        output := &codeconnections.GetRepositoryLinkOutput{}

        mockClient.On("GetRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositorySyncStatus", func(t *testing.T) {
        input := &codeconnections.GetRepositorySyncStatusInput{}
        output := &codeconnections.GetRepositorySyncStatusOutput{}

        mockClient.On("GetRepositorySyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositorySyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceSyncStatus", func(t *testing.T) {
        input := &codeconnections.GetResourceSyncStatusInput{}
        output := &codeconnections.GetResourceSyncStatusOutput{}

        mockClient.On("GetResourceSyncStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceSyncStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSyncBlockerSummary", func(t *testing.T) {
        input := &codeconnections.GetSyncBlockerSummaryInput{}
        output := &codeconnections.GetSyncBlockerSummaryOutput{}

        mockClient.On("GetSyncBlockerSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetSyncBlockerSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSyncConfiguration", func(t *testing.T) {
        input := &codeconnections.GetSyncConfigurationInput{}
        output := &codeconnections.GetSyncConfigurationOutput{}

        mockClient.On("GetSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnections", func(t *testing.T) {
        input := &codeconnections.ListConnectionsInput{}
        output := &codeconnections.ListConnectionsOutput{}

        mockClient.On("ListConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHosts", func(t *testing.T) {
        input := &codeconnections.ListHostsInput{}
        output := &codeconnections.ListHostsOutput{}

        mockClient.On("ListHosts", ctx, input).Return(output, nil)

        result, err := mockClient.ListHosts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositoryLinks", func(t *testing.T) {
        input := &codeconnections.ListRepositoryLinksInput{}
        output := &codeconnections.ListRepositoryLinksOutput{}

        mockClient.On("ListRepositoryLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositoryLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositorySyncDefinitions", func(t *testing.T) {
        input := &codeconnections.ListRepositorySyncDefinitionsInput{}
        output := &codeconnections.ListRepositorySyncDefinitionsOutput{}

        mockClient.On("ListRepositorySyncDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositorySyncDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSyncConfigurations", func(t *testing.T) {
        input := &codeconnections.ListSyncConfigurationsInput{}
        output := &codeconnections.ListSyncConfigurationsOutput{}

        mockClient.On("ListSyncConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSyncConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codeconnections.ListTagsForResourceInput{}
        output := &codeconnections.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codeconnections.TagResourceInput{}
        output := &codeconnections.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codeconnections.UntagResourceInput{}
        output := &codeconnections.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHost", func(t *testing.T) {
        input := &codeconnections.UpdateHostInput{}
        output := &codeconnections.UpdateHostOutput{}

        mockClient.On("UpdateHost", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepositoryLink", func(t *testing.T) {
        input := &codeconnections.UpdateRepositoryLinkInput{}
        output := &codeconnections.UpdateRepositoryLinkOutput{}

        mockClient.On("UpdateRepositoryLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepositoryLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSyncBlocker", func(t *testing.T) {
        input := &codeconnections.UpdateSyncBlockerInput{}
        output := &codeconnections.UpdateSyncBlockerOutput{}

        mockClient.On("UpdateSyncBlocker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSyncBlocker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSyncConfiguration", func(t *testing.T) {
        input := &codeconnections.UpdateSyncConfigurationInput{}
        output := &codeconnections.UpdateSyncConfigurationOutput{}

        mockClient.On("UpdateSyncConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSyncConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
