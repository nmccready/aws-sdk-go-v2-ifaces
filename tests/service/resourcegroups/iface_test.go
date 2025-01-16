// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package resourcegroups_test

// tests for the resourcegroups service interface for 
// this ../../../service/resourcegroups/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourcegroups/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourcegroups/resourcegroups_iface"
	"github.com/stretchr/testify/assert"
)

func TestResourcegroupsServiceCanBeMocked(t *testing.T) {
	var iface resourcegroups_iface.IClient
	iface = &resourcegroups.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := resourcegroups.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelTagSyncTask", func(t *testing.T) {
        input := &resourcegroups.CancelTagSyncTaskInput{}
        output := &resourcegroups.CancelTagSyncTaskOutput{}

        mockClient.On("CancelTagSyncTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelTagSyncTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &resourcegroups.CreateGroupInput{}
        output := &resourcegroups.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &resourcegroups.DeleteGroupInput{}
        output := &resourcegroups.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &resourcegroups.GetAccountSettingsInput{}
        output := &resourcegroups.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &resourcegroups.GetGroupInput{}
        output := &resourcegroups.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupConfiguration", func(t *testing.T) {
        input := &resourcegroups.GetGroupConfigurationInput{}
        output := &resourcegroups.GetGroupConfigurationOutput{}

        mockClient.On("GetGroupConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupQuery", func(t *testing.T) {
        input := &resourcegroups.GetGroupQueryInput{}
        output := &resourcegroups.GetGroupQueryOutput{}

        mockClient.On("GetGroupQuery", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTagSyncTask", func(t *testing.T) {
        input := &resourcegroups.GetTagSyncTaskInput{}
        output := &resourcegroups.GetTagSyncTaskOutput{}

        mockClient.On("GetTagSyncTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetTagSyncTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTags", func(t *testing.T) {
        input := &resourcegroups.GetTagsInput{}
        output := &resourcegroups.GetTagsOutput{}

        mockClient.On("GetTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGroupResources", func(t *testing.T) {
        input := &resourcegroups.GroupResourcesInput{}
        output := &resourcegroups.GroupResourcesOutput{}

        mockClient.On("GroupResources", ctx, input).Return(output, nil)

        result, err := mockClient.GroupResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupResources", func(t *testing.T) {
        input := &resourcegroups.ListGroupResourcesInput{}
        output := &resourcegroups.ListGroupResourcesOutput{}

        mockClient.On("ListGroupResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupingStatuses", func(t *testing.T) {
        input := &resourcegroups.ListGroupingStatusesInput{}
        output := &resourcegroups.ListGroupingStatusesOutput{}

        mockClient.On("ListGroupingStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupingStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &resourcegroups.ListGroupsInput{}
        output := &resourcegroups.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagSyncTasks", func(t *testing.T) {
        input := &resourcegroups.ListTagSyncTasksInput{}
        output := &resourcegroups.ListTagSyncTasksOutput{}

        mockClient.On("ListTagSyncTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagSyncTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGroupConfiguration", func(t *testing.T) {
        input := &resourcegroups.PutGroupConfigurationInput{}
        output := &resourcegroups.PutGroupConfigurationOutput{}

        mockClient.On("PutGroupConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutGroupConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchResources", func(t *testing.T) {
        input := &resourcegroups.SearchResourcesInput{}
        output := &resourcegroups.SearchResourcesOutput{}

        mockClient.On("SearchResources", ctx, input).Return(output, nil)

        result, err := mockClient.SearchResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTagSyncTask", func(t *testing.T) {
        input := &resourcegroups.StartTagSyncTaskInput{}
        output := &resourcegroups.StartTagSyncTaskOutput{}

        mockClient.On("StartTagSyncTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartTagSyncTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTag", func(t *testing.T) {
        input := &resourcegroups.TagInput{}
        output := &resourcegroups.TagOutput{}

        mockClient.On("Tag", ctx, input).Return(output, nil)

        result, err := mockClient.Tag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUngroupResources", func(t *testing.T) {
        input := &resourcegroups.UngroupResourcesInput{}
        output := &resourcegroups.UngroupResourcesOutput{}

        mockClient.On("UngroupResources", ctx, input).Return(output, nil)

        result, err := mockClient.UngroupResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntag", func(t *testing.T) {
        input := &resourcegroups.UntagInput{}
        output := &resourcegroups.UntagOutput{}

        mockClient.On("Untag", ctx, input).Return(output, nil)

        result, err := mockClient.Untag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &resourcegroups.UpdateAccountSettingsInput{}
        output := &resourcegroups.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &resourcegroups.UpdateGroupInput{}
        output := &resourcegroups.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroupQuery", func(t *testing.T) {
        input := &resourcegroups.UpdateGroupQueryInput{}
        output := &resourcegroups.UpdateGroupQueryOutput{}

        mockClient.On("UpdateGroupQuery", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroupQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
