package resourceexplorer2_test

// tests for the resourceexplorer2 service interface for this ../../../service/resourceexplorer2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourceexplorer2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourceexplorer2/resourceexplorer2_iface"
	"github.com/stretchr/testify/assert"
)

func TestResourceexplorer2ServiceCanBeMocked(t *testing.T) {
	var iface resourceexplorer2_iface.IClient
	iface = &resourceexplorer2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := resourceexplorer2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDefaultView", func(t *testing.T) {
        input := &resourceexplorer2.AssociateDefaultViewInput{}
        output := &resourceexplorer2.AssociateDefaultViewOutput{}

        mockClient.On("AssociateDefaultView", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDefaultView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetView", func(t *testing.T) {
        input := &resourceexplorer2.BatchGetViewInput{}
        output := &resourceexplorer2.BatchGetViewOutput{}

        mockClient.On("BatchGetView", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIndex", func(t *testing.T) {
        input := &resourceexplorer2.CreateIndexInput{}
        output := &resourceexplorer2.CreateIndexOutput{}

        mockClient.On("CreateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateView", func(t *testing.T) {
        input := &resourceexplorer2.CreateViewInput{}
        output := &resourceexplorer2.CreateViewOutput{}

        mockClient.On("CreateView", ctx, input).Return(output, nil)

        result, err := mockClient.CreateView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIndex", func(t *testing.T) {
        input := &resourceexplorer2.DeleteIndexInput{}
        output := &resourceexplorer2.DeleteIndexOutput{}

        mockClient.On("DeleteIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteView", func(t *testing.T) {
        input := &resourceexplorer2.DeleteViewInput{}
        output := &resourceexplorer2.DeleteViewOutput{}

        mockClient.On("DeleteView", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDefaultView", func(t *testing.T) {
        input := &resourceexplorer2.DisassociateDefaultViewInput{}
        output := &resourceexplorer2.DisassociateDefaultViewOutput{}

        mockClient.On("DisassociateDefaultView", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDefaultView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountLevelServiceConfiguration", func(t *testing.T) {
        input := &resourceexplorer2.GetAccountLevelServiceConfigurationInput{}
        output := &resourceexplorer2.GetAccountLevelServiceConfigurationOutput{}

        mockClient.On("GetAccountLevelServiceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountLevelServiceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDefaultView", func(t *testing.T) {
        input := &resourceexplorer2.GetDefaultViewInput{}
        output := &resourceexplorer2.GetDefaultViewOutput{}

        mockClient.On("GetDefaultView", ctx, input).Return(output, nil)

        result, err := mockClient.GetDefaultView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIndex", func(t *testing.T) {
        input := &resourceexplorer2.GetIndexInput{}
        output := &resourceexplorer2.GetIndexOutput{}

        mockClient.On("GetIndex", ctx, input).Return(output, nil)

        result, err := mockClient.GetIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetView", func(t *testing.T) {
        input := &resourceexplorer2.GetViewInput{}
        output := &resourceexplorer2.GetViewOutput{}

        mockClient.On("GetView", ctx, input).Return(output, nil)

        result, err := mockClient.GetView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndexes", func(t *testing.T) {
        input := &resourceexplorer2.ListIndexesInput{}
        output := &resourceexplorer2.ListIndexesOutput{}

        mockClient.On("ListIndexes", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndexes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndexesForMembers", func(t *testing.T) {
        input := &resourceexplorer2.ListIndexesForMembersInput{}
        output := &resourceexplorer2.ListIndexesForMembersOutput{}

        mockClient.On("ListIndexesForMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndexesForMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSupportedResourceTypes", func(t *testing.T) {
        input := &resourceexplorer2.ListSupportedResourceTypesInput{}
        output := &resourceexplorer2.ListSupportedResourceTypesOutput{}

        mockClient.On("ListSupportedResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListSupportedResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &resourceexplorer2.ListTagsForResourceInput{}
        output := &resourceexplorer2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListViews", func(t *testing.T) {
        input := &resourceexplorer2.ListViewsInput{}
        output := &resourceexplorer2.ListViewsOutput{}

        mockClient.On("ListViews", ctx, input).Return(output, nil)

        result, err := mockClient.ListViews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearch", func(t *testing.T) {
        input := &resourceexplorer2.SearchInput{}
        output := &resourceexplorer2.SearchOutput{}

        mockClient.On("Search", ctx, input).Return(output, nil)

        result, err := mockClient.Search(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &resourceexplorer2.TagResourceInput{}
        output := &resourceexplorer2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &resourceexplorer2.UntagResourceInput{}
        output := &resourceexplorer2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIndexType", func(t *testing.T) {
        input := &resourceexplorer2.UpdateIndexTypeInput{}
        output := &resourceexplorer2.UpdateIndexTypeOutput{}

        mockClient.On("UpdateIndexType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIndexType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateView", func(t *testing.T) {
        input := &resourceexplorer2.UpdateViewInput{}
        output := &resourceexplorer2.UpdateViewOutput{}

        mockClient.On("UpdateView", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
