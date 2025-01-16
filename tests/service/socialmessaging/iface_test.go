// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package socialmessaging_test

// tests for the socialmessaging service interface for 
// this ../../../service/socialmessaging/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/socialmessaging"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/socialmessaging/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/socialmessaging/socialmessaging_iface"
	"github.com/stretchr/testify/assert"
)

func TestSocialmessagingServiceCanBeMocked(t *testing.T) {
	var iface socialmessaging_iface.IClient
	iface = &socialmessaging.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := socialmessaging.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWhatsAppBusinessAccount", func(t *testing.T) {
        input := &socialmessaging.AssociateWhatsAppBusinessAccountInput{}
        output := &socialmessaging.AssociateWhatsAppBusinessAccountOutput{}

        mockClient.On("AssociateWhatsAppBusinessAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWhatsAppBusinessAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWhatsAppMessageMedia", func(t *testing.T) {
        input := &socialmessaging.DeleteWhatsAppMessageMediaInput{}
        output := &socialmessaging.DeleteWhatsAppMessageMediaOutput{}

        mockClient.On("DeleteWhatsAppMessageMedia", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWhatsAppMessageMedia(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWhatsAppBusinessAccount", func(t *testing.T) {
        input := &socialmessaging.DisassociateWhatsAppBusinessAccountInput{}
        output := &socialmessaging.DisassociateWhatsAppBusinessAccountOutput{}

        mockClient.On("DisassociateWhatsAppBusinessAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWhatsAppBusinessAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLinkedWhatsAppBusinessAccount", func(t *testing.T) {
        input := &socialmessaging.GetLinkedWhatsAppBusinessAccountInput{}
        output := &socialmessaging.GetLinkedWhatsAppBusinessAccountOutput{}

        mockClient.On("GetLinkedWhatsAppBusinessAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetLinkedWhatsAppBusinessAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLinkedWhatsAppBusinessAccountPhoneNumber", func(t *testing.T) {
        input := &socialmessaging.GetLinkedWhatsAppBusinessAccountPhoneNumberInput{}
        output := &socialmessaging.GetLinkedWhatsAppBusinessAccountPhoneNumberOutput{}

        mockClient.On("GetLinkedWhatsAppBusinessAccountPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.GetLinkedWhatsAppBusinessAccountPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWhatsAppMessageMedia", func(t *testing.T) {
        input := &socialmessaging.GetWhatsAppMessageMediaInput{}
        output := &socialmessaging.GetWhatsAppMessageMediaOutput{}

        mockClient.On("GetWhatsAppMessageMedia", ctx, input).Return(output, nil)

        result, err := mockClient.GetWhatsAppMessageMedia(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLinkedWhatsAppBusinessAccounts", func(t *testing.T) {
        input := &socialmessaging.ListLinkedWhatsAppBusinessAccountsInput{}
        output := &socialmessaging.ListLinkedWhatsAppBusinessAccountsOutput{}

        mockClient.On("ListLinkedWhatsAppBusinessAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListLinkedWhatsAppBusinessAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &socialmessaging.ListTagsForResourceInput{}
        output := &socialmessaging.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostWhatsAppMessageMedia", func(t *testing.T) {
        input := &socialmessaging.PostWhatsAppMessageMediaInput{}
        output := &socialmessaging.PostWhatsAppMessageMediaOutput{}

        mockClient.On("PostWhatsAppMessageMedia", ctx, input).Return(output, nil)

        result, err := mockClient.PostWhatsAppMessageMedia(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutWhatsAppBusinessAccountEventDestinations", func(t *testing.T) {
        input := &socialmessaging.PutWhatsAppBusinessAccountEventDestinationsInput{}
        output := &socialmessaging.PutWhatsAppBusinessAccountEventDestinationsOutput{}

        mockClient.On("PutWhatsAppBusinessAccountEventDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.PutWhatsAppBusinessAccountEventDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendWhatsAppMessage", func(t *testing.T) {
        input := &socialmessaging.SendWhatsAppMessageInput{}
        output := &socialmessaging.SendWhatsAppMessageOutput{}

        mockClient.On("SendWhatsAppMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendWhatsAppMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &socialmessaging.TagResourceInput{}
        output := &socialmessaging.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &socialmessaging.UntagResourceInput{}
        output := &socialmessaging.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
