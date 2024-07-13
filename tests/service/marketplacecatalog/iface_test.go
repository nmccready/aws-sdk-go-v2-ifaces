package marketplacecatalog_test

// tests for the marketplacecatalog service interface for this ../../../service/marketplacecatalog/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacecatalog/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacecatalog/marketplacecatalog_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplacecatalogServiceCanBeMocked(t *testing.T) {
	var iface marketplacecatalog_iface.IClient
	iface = &marketplacecatalog.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplacecatalog.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDescribeEntities", func(t *testing.T) {
        input := &marketplacecatalog.BatchDescribeEntitiesInput{}
        output := &marketplacecatalog.BatchDescribeEntitiesOutput{}

        mockClient.On("BatchDescribeEntities", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDescribeEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelChangeSet", func(t *testing.T) {
        input := &marketplacecatalog.CancelChangeSetInput{}
        output := &marketplacecatalog.CancelChangeSetOutput{}

        mockClient.On("CancelChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.CancelChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &marketplacecatalog.DeleteResourcePolicyInput{}
        output := &marketplacecatalog.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChangeSet", func(t *testing.T) {
        input := &marketplacecatalog.DescribeChangeSetInput{}
        output := &marketplacecatalog.DescribeChangeSetOutput{}

        mockClient.On("DescribeChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntity", func(t *testing.T) {
        input := &marketplacecatalog.DescribeEntityInput{}
        output := &marketplacecatalog.DescribeEntityOutput{}

        mockClient.On("DescribeEntity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &marketplacecatalog.GetResourcePolicyInput{}
        output := &marketplacecatalog.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChangeSets", func(t *testing.T) {
        input := &marketplacecatalog.ListChangeSetsInput{}
        output := &marketplacecatalog.ListChangeSetsOutput{}

        mockClient.On("ListChangeSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListChangeSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntities", func(t *testing.T) {
        input := &marketplacecatalog.ListEntitiesInput{}
        output := &marketplacecatalog.ListEntitiesOutput{}

        mockClient.On("ListEntities", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &marketplacecatalog.ListTagsForResourceInput{}
        output := &marketplacecatalog.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &marketplacecatalog.PutResourcePolicyInput{}
        output := &marketplacecatalog.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartChangeSet", func(t *testing.T) {
        input := &marketplacecatalog.StartChangeSetInput{}
        output := &marketplacecatalog.StartChangeSetOutput{}

        mockClient.On("StartChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.StartChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &marketplacecatalog.TagResourceInput{}
        output := &marketplacecatalog.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &marketplacecatalog.UntagResourceInput{}
        output := &marketplacecatalog.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
