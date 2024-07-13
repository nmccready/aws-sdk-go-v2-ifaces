package ssmcontacts_test

// tests for the ssmcontacts service interface for this ../../../service/ssmcontacts/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmcontacts/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmcontacts/ssmcontacts_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsmcontactsServiceCanBeMocked(t *testing.T) {
	var iface ssmcontacts_iface.IClient
	iface = &ssmcontacts.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssmcontacts.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptPage", func(t *testing.T) {
        input := &ssmcontacts.AcceptPageInput{}
        output := &ssmcontacts.AcceptPageOutput{}

        mockClient.On("AcceptPage", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptPage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateContactChannel", func(t *testing.T) {
        input := &ssmcontacts.ActivateContactChannelInput{}
        output := &ssmcontacts.ActivateContactChannelOutput{}

        mockClient.On("ActivateContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContact", func(t *testing.T) {
        input := &ssmcontacts.CreateContactInput{}
        output := &ssmcontacts.CreateContactOutput{}

        mockClient.On("CreateContact", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContactChannel", func(t *testing.T) {
        input := &ssmcontacts.CreateContactChannelInput{}
        output := &ssmcontacts.CreateContactChannelOutput{}

        mockClient.On("CreateContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRotation", func(t *testing.T) {
        input := &ssmcontacts.CreateRotationInput{}
        output := &ssmcontacts.CreateRotationOutput{}

        mockClient.On("CreateRotation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRotationOverride", func(t *testing.T) {
        input := &ssmcontacts.CreateRotationOverrideInput{}
        output := &ssmcontacts.CreateRotationOverrideOutput{}

        mockClient.On("CreateRotationOverride", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRotationOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateContactChannel", func(t *testing.T) {
        input := &ssmcontacts.DeactivateContactChannelInput{}
        output := &ssmcontacts.DeactivateContactChannelOutput{}

        mockClient.On("DeactivateContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContact", func(t *testing.T) {
        input := &ssmcontacts.DeleteContactInput{}
        output := &ssmcontacts.DeleteContactOutput{}

        mockClient.On("DeleteContact", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactChannel", func(t *testing.T) {
        input := &ssmcontacts.DeleteContactChannelInput{}
        output := &ssmcontacts.DeleteContactChannelOutput{}

        mockClient.On("DeleteContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRotation", func(t *testing.T) {
        input := &ssmcontacts.DeleteRotationInput{}
        output := &ssmcontacts.DeleteRotationOutput{}

        mockClient.On("DeleteRotation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRotationOverride", func(t *testing.T) {
        input := &ssmcontacts.DeleteRotationOverrideInput{}
        output := &ssmcontacts.DeleteRotationOverrideOutput{}

        mockClient.On("DeleteRotationOverride", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRotationOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEngagement", func(t *testing.T) {
        input := &ssmcontacts.DescribeEngagementInput{}
        output := &ssmcontacts.DescribeEngagementOutput{}

        mockClient.On("DescribeEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePage", func(t *testing.T) {
        input := &ssmcontacts.DescribePageInput{}
        output := &ssmcontacts.DescribePageOutput{}

        mockClient.On("DescribePage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContact", func(t *testing.T) {
        input := &ssmcontacts.GetContactInput{}
        output := &ssmcontacts.GetContactOutput{}

        mockClient.On("GetContact", ctx, input).Return(output, nil)

        result, err := mockClient.GetContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactChannel", func(t *testing.T) {
        input := &ssmcontacts.GetContactChannelInput{}
        output := &ssmcontacts.GetContactChannelOutput{}

        mockClient.On("GetContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactPolicy", func(t *testing.T) {
        input := &ssmcontacts.GetContactPolicyInput{}
        output := &ssmcontacts.GetContactPolicyOutput{}

        mockClient.On("GetContactPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRotation", func(t *testing.T) {
        input := &ssmcontacts.GetRotationInput{}
        output := &ssmcontacts.GetRotationOutput{}

        mockClient.On("GetRotation", ctx, input).Return(output, nil)

        result, err := mockClient.GetRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRotationOverride", func(t *testing.T) {
        input := &ssmcontacts.GetRotationOverrideInput{}
        output := &ssmcontacts.GetRotationOverrideOutput{}

        mockClient.On("GetRotationOverride", ctx, input).Return(output, nil)

        result, err := mockClient.GetRotationOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactChannels", func(t *testing.T) {
        input := &ssmcontacts.ListContactChannelsInput{}
        output := &ssmcontacts.ListContactChannelsOutput{}

        mockClient.On("ListContactChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContacts", func(t *testing.T) {
        input := &ssmcontacts.ListContactsInput{}
        output := &ssmcontacts.ListContactsOutput{}

        mockClient.On("ListContacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListContacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEngagements", func(t *testing.T) {
        input := &ssmcontacts.ListEngagementsInput{}
        output := &ssmcontacts.ListEngagementsOutput{}

        mockClient.On("ListEngagements", ctx, input).Return(output, nil)

        result, err := mockClient.ListEngagements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPageReceipts", func(t *testing.T) {
        input := &ssmcontacts.ListPageReceiptsInput{}
        output := &ssmcontacts.ListPageReceiptsOutput{}

        mockClient.On("ListPageReceipts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPageReceipts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPageResolutions", func(t *testing.T) {
        input := &ssmcontacts.ListPageResolutionsInput{}
        output := &ssmcontacts.ListPageResolutionsOutput{}

        mockClient.On("ListPageResolutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPageResolutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPagesByContact", func(t *testing.T) {
        input := &ssmcontacts.ListPagesByContactInput{}
        output := &ssmcontacts.ListPagesByContactOutput{}

        mockClient.On("ListPagesByContact", ctx, input).Return(output, nil)

        result, err := mockClient.ListPagesByContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPagesByEngagement", func(t *testing.T) {
        input := &ssmcontacts.ListPagesByEngagementInput{}
        output := &ssmcontacts.ListPagesByEngagementOutput{}

        mockClient.On("ListPagesByEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.ListPagesByEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPreviewRotationShifts", func(t *testing.T) {
        input := &ssmcontacts.ListPreviewRotationShiftsInput{}
        output := &ssmcontacts.ListPreviewRotationShiftsOutput{}

        mockClient.On("ListPreviewRotationShifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPreviewRotationShifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRotationOverrides", func(t *testing.T) {
        input := &ssmcontacts.ListRotationOverridesInput{}
        output := &ssmcontacts.ListRotationOverridesOutput{}

        mockClient.On("ListRotationOverrides", ctx, input).Return(output, nil)

        result, err := mockClient.ListRotationOverrides(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRotationShifts", func(t *testing.T) {
        input := &ssmcontacts.ListRotationShiftsInput{}
        output := &ssmcontacts.ListRotationShiftsOutput{}

        mockClient.On("ListRotationShifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListRotationShifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRotations", func(t *testing.T) {
        input := &ssmcontacts.ListRotationsInput{}
        output := &ssmcontacts.ListRotationsOutput{}

        mockClient.On("ListRotations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRotations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ssmcontacts.ListTagsForResourceInput{}
        output := &ssmcontacts.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutContactPolicy", func(t *testing.T) {
        input := &ssmcontacts.PutContactPolicyInput{}
        output := &ssmcontacts.PutContactPolicyOutput{}

        mockClient.On("PutContactPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutContactPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendActivationCode", func(t *testing.T) {
        input := &ssmcontacts.SendActivationCodeInput{}
        output := &ssmcontacts.SendActivationCodeOutput{}

        mockClient.On("SendActivationCode", ctx, input).Return(output, nil)

        result, err := mockClient.SendActivationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEngagement", func(t *testing.T) {
        input := &ssmcontacts.StartEngagementInput{}
        output := &ssmcontacts.StartEngagementOutput{}

        mockClient.On("StartEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.StartEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEngagement", func(t *testing.T) {
        input := &ssmcontacts.StopEngagementInput{}
        output := &ssmcontacts.StopEngagementOutput{}

        mockClient.On("StopEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.StopEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ssmcontacts.TagResourceInput{}
        output := &ssmcontacts.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ssmcontacts.UntagResourceInput{}
        output := &ssmcontacts.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContact", func(t *testing.T) {
        input := &ssmcontacts.UpdateContactInput{}
        output := &ssmcontacts.UpdateContactOutput{}

        mockClient.On("UpdateContact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactChannel", func(t *testing.T) {
        input := &ssmcontacts.UpdateContactChannelInput{}
        output := &ssmcontacts.UpdateContactChannelOutput{}

        mockClient.On("UpdateContactChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRotation", func(t *testing.T) {
        input := &ssmcontacts.UpdateRotationInput{}
        output := &ssmcontacts.UpdateRotationOutput{}

        mockClient.On("UpdateRotation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
