package repostspace_test

// tests for the repostspace service interface for this ../../../service/repostspace/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/repostspace"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/repostspace/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/repostspace/repostspace_iface"
	"github.com/stretchr/testify/assert"
)

func TestRepostspaceServiceCanBeMocked(t *testing.T) {
	var iface repostspace_iface.IClient
	iface = &repostspace.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := repostspace.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSpace", func(t *testing.T) {
        input := &repostspace.CreateSpaceInput{}
        output := &repostspace.CreateSpaceOutput{}

        mockClient.On("CreateSpace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSpace", func(t *testing.T) {
        input := &repostspace.DeleteSpaceInput{}
        output := &repostspace.DeleteSpaceOutput{}

        mockClient.On("DeleteSpace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterAdmin", func(t *testing.T) {
        input := &repostspace.DeregisterAdminInput{}
        output := &repostspace.DeregisterAdminOutput{}

        mockClient.On("DeregisterAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpace", func(t *testing.T) {
        input := &repostspace.GetSpaceInput{}
        output := &repostspace.GetSpaceOutput{}

        mockClient.On("GetSpace", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpaces", func(t *testing.T) {
        input := &repostspace.ListSpacesInput{}
        output := &repostspace.ListSpacesOutput{}

        mockClient.On("ListSpaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &repostspace.ListTagsForResourceInput{}
        output := &repostspace.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterAdmin", func(t *testing.T) {
        input := &repostspace.RegisterAdminInput{}
        output := &repostspace.RegisterAdminOutput{}

        mockClient.On("RegisterAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendInvites", func(t *testing.T) {
        input := &repostspace.SendInvitesInput{}
        output := &repostspace.SendInvitesOutput{}

        mockClient.On("SendInvites", ctx, input).Return(output, nil)

        result, err := mockClient.SendInvites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &repostspace.TagResourceInput{}
        output := &repostspace.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &repostspace.UntagResourceInput{}
        output := &repostspace.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSpace", func(t *testing.T) {
        input := &repostspace.UpdateSpaceInput{}
        output := &repostspace.UpdateSpaceOutput{}

        mockClient.On("UpdateSpace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSpace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
