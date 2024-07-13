package oam_test

// tests for the oam service interface for this ../../../service/oam/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/oam"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/oam/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/oam/oam_iface"
	"github.com/stretchr/testify/assert"
)

func TestOamServiceCanBeMocked(t *testing.T) {
	var iface oam_iface.IClient
	iface = &oam.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := oam.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLink", func(t *testing.T) {
        input := &oam.CreateLinkInput{}
        output := &oam.CreateLinkOutput{}

        mockClient.On("CreateLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSink", func(t *testing.T) {
        input := &oam.CreateSinkInput{}
        output := &oam.CreateSinkOutput{}

        mockClient.On("CreateSink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLink", func(t *testing.T) {
        input := &oam.DeleteLinkInput{}
        output := &oam.DeleteLinkOutput{}

        mockClient.On("DeleteLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSink", func(t *testing.T) {
        input := &oam.DeleteSinkInput{}
        output := &oam.DeleteSinkOutput{}

        mockClient.On("DeleteSink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLink", func(t *testing.T) {
        input := &oam.GetLinkInput{}
        output := &oam.GetLinkOutput{}

        mockClient.On("GetLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSink", func(t *testing.T) {
        input := &oam.GetSinkInput{}
        output := &oam.GetSinkOutput{}

        mockClient.On("GetSink", ctx, input).Return(output, nil)

        result, err := mockClient.GetSink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSinkPolicy", func(t *testing.T) {
        input := &oam.GetSinkPolicyInput{}
        output := &oam.GetSinkPolicyOutput{}

        mockClient.On("GetSinkPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetSinkPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedLinks", func(t *testing.T) {
        input := &oam.ListAttachedLinksInput{}
        output := &oam.ListAttachedLinksOutput{}

        mockClient.On("ListAttachedLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLinks", func(t *testing.T) {
        input := &oam.ListLinksInput{}
        output := &oam.ListLinksOutput{}

        mockClient.On("ListLinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSinks", func(t *testing.T) {
        input := &oam.ListSinksInput{}
        output := &oam.ListSinksOutput{}

        mockClient.On("ListSinks", ctx, input).Return(output, nil)

        result, err := mockClient.ListSinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &oam.ListTagsForResourceInput{}
        output := &oam.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSinkPolicy", func(t *testing.T) {
        input := &oam.PutSinkPolicyInput{}
        output := &oam.PutSinkPolicyOutput{}

        mockClient.On("PutSinkPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutSinkPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &oam.TagResourceInput{}
        output := &oam.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &oam.UntagResourceInput{}
        output := &oam.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLink", func(t *testing.T) {
        input := &oam.UpdateLinkInput{}
        output := &oam.UpdateLinkOutput{}

        mockClient.On("UpdateLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
