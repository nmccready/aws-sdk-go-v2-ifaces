package sns_test

// tests for the sns service interface for this ../../../service/sns/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sns/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sns/sns_iface"
	"github.com/stretchr/testify/assert"
)

func TestSnsServiceCanBeMocked(t *testing.T) {
	var iface sns_iface.IClient
	iface = &sns.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sns.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddPermission", func(t *testing.T) {
        input := &sns.AddPermissionInput{}
        output := &sns.AddPermissionOutput{}

        mockClient.On("AddPermission", ctx, input).Return(output, nil)

        result, err := mockClient.AddPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckIfPhoneNumberIsOptedOut", func(t *testing.T) {
        input := &sns.CheckIfPhoneNumberIsOptedOutInput{}
        output := &sns.CheckIfPhoneNumberIsOptedOutOutput{}

        mockClient.On("CheckIfPhoneNumberIsOptedOut", ctx, input).Return(output, nil)

        result, err := mockClient.CheckIfPhoneNumberIsOptedOut(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmSubscription", func(t *testing.T) {
        input := &sns.ConfirmSubscriptionInput{}
        output := &sns.ConfirmSubscriptionOutput{}

        mockClient.On("ConfirmSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlatformApplication", func(t *testing.T) {
        input := &sns.CreatePlatformApplicationInput{}
        output := &sns.CreatePlatformApplicationOutput{}

        mockClient.On("CreatePlatformApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlatformApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlatformEndpoint", func(t *testing.T) {
        input := &sns.CreatePlatformEndpointInput{}
        output := &sns.CreatePlatformEndpointOutput{}

        mockClient.On("CreatePlatformEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlatformEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSMSSandboxPhoneNumber", func(t *testing.T) {
        input := &sns.CreateSMSSandboxPhoneNumberInput{}
        output := &sns.CreateSMSSandboxPhoneNumberOutput{}

        mockClient.On("CreateSMSSandboxPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSMSSandboxPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTopic", func(t *testing.T) {
        input := &sns.CreateTopicInput{}
        output := &sns.CreateTopicOutput{}

        mockClient.On("CreateTopic", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &sns.DeleteEndpointInput{}
        output := &sns.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlatformApplication", func(t *testing.T) {
        input := &sns.DeletePlatformApplicationInput{}
        output := &sns.DeletePlatformApplicationOutput{}

        mockClient.On("DeletePlatformApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlatformApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSMSSandboxPhoneNumber", func(t *testing.T) {
        input := &sns.DeleteSMSSandboxPhoneNumberInput{}
        output := &sns.DeleteSMSSandboxPhoneNumberOutput{}

        mockClient.On("DeleteSMSSandboxPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSMSSandboxPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTopic", func(t *testing.T) {
        input := &sns.DeleteTopicInput{}
        output := &sns.DeleteTopicOutput{}

        mockClient.On("DeleteTopic", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataProtectionPolicy", func(t *testing.T) {
        input := &sns.GetDataProtectionPolicyInput{}
        output := &sns.GetDataProtectionPolicyOutput{}

        mockClient.On("GetDataProtectionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataProtectionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEndpointAttributes", func(t *testing.T) {
        input := &sns.GetEndpointAttributesInput{}
        output := &sns.GetEndpointAttributesOutput{}

        mockClient.On("GetEndpointAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetEndpointAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlatformApplicationAttributes", func(t *testing.T) {
        input := &sns.GetPlatformApplicationAttributesInput{}
        output := &sns.GetPlatformApplicationAttributesOutput{}

        mockClient.On("GetPlatformApplicationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlatformApplicationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSMSAttributes", func(t *testing.T) {
        input := &sns.GetSMSAttributesInput{}
        output := &sns.GetSMSAttributesOutput{}

        mockClient.On("GetSMSAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetSMSAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSMSSandboxAccountStatus", func(t *testing.T) {
        input := &sns.GetSMSSandboxAccountStatusInput{}
        output := &sns.GetSMSSandboxAccountStatusOutput{}

        mockClient.On("GetSMSSandboxAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetSMSSandboxAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionAttributes", func(t *testing.T) {
        input := &sns.GetSubscriptionAttributesInput{}
        output := &sns.GetSubscriptionAttributesOutput{}

        mockClient.On("GetSubscriptionAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTopicAttributes", func(t *testing.T) {
        input := &sns.GetTopicAttributesInput{}
        output := &sns.GetTopicAttributesOutput{}

        mockClient.On("GetTopicAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetTopicAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpointsByPlatformApplication", func(t *testing.T) {
        input := &sns.ListEndpointsByPlatformApplicationInput{}
        output := &sns.ListEndpointsByPlatformApplicationOutput{}

        mockClient.On("ListEndpointsByPlatformApplication", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpointsByPlatformApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOriginationNumbers", func(t *testing.T) {
        input := &sns.ListOriginationNumbersInput{}
        output := &sns.ListOriginationNumbersOutput{}

        mockClient.On("ListOriginationNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.ListOriginationNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumbersOptedOut", func(t *testing.T) {
        input := &sns.ListPhoneNumbersOptedOutInput{}
        output := &sns.ListPhoneNumbersOptedOutOutput{}

        mockClient.On("ListPhoneNumbersOptedOut", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumbersOptedOut(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlatformApplications", func(t *testing.T) {
        input := &sns.ListPlatformApplicationsInput{}
        output := &sns.ListPlatformApplicationsOutput{}

        mockClient.On("ListPlatformApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlatformApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSMSSandboxPhoneNumbers", func(t *testing.T) {
        input := &sns.ListSMSSandboxPhoneNumbersInput{}
        output := &sns.ListSMSSandboxPhoneNumbersOutput{}

        mockClient.On("ListSMSSandboxPhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.ListSMSSandboxPhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptions", func(t *testing.T) {
        input := &sns.ListSubscriptionsInput{}
        output := &sns.ListSubscriptionsOutput{}

        mockClient.On("ListSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionsByTopic", func(t *testing.T) {
        input := &sns.ListSubscriptionsByTopicInput{}
        output := &sns.ListSubscriptionsByTopicOutput{}

        mockClient.On("ListSubscriptionsByTopic", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionsByTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &sns.ListTagsForResourceInput{}
        output := &sns.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopics", func(t *testing.T) {
        input := &sns.ListTopicsInput{}
        output := &sns.ListTopicsOutput{}

        mockClient.On("ListTopics", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestOptInPhoneNumber", func(t *testing.T) {
        input := &sns.OptInPhoneNumberInput{}
        output := &sns.OptInPhoneNumberOutput{}

        mockClient.On("OptInPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.OptInPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublish", func(t *testing.T) {
        input := &sns.PublishInput{}
        output := &sns.PublishOutput{}

        mockClient.On("Publish", ctx, input).Return(output, nil)

        result, err := mockClient.Publish(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishBatch", func(t *testing.T) {
        input := &sns.PublishBatchInput{}
        output := &sns.PublishBatchOutput{}

        mockClient.On("PublishBatch", ctx, input).Return(output, nil)

        result, err := mockClient.PublishBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDataProtectionPolicy", func(t *testing.T) {
        input := &sns.PutDataProtectionPolicyInput{}
        output := &sns.PutDataProtectionPolicyOutput{}

        mockClient.On("PutDataProtectionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutDataProtectionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &sns.RemovePermissionInput{}
        output := &sns.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetEndpointAttributes", func(t *testing.T) {
        input := &sns.SetEndpointAttributesInput{}
        output := &sns.SetEndpointAttributesOutput{}

        mockClient.On("SetEndpointAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetEndpointAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetPlatformApplicationAttributes", func(t *testing.T) {
        input := &sns.SetPlatformApplicationAttributesInput{}
        output := &sns.SetPlatformApplicationAttributesOutput{}

        mockClient.On("SetPlatformApplicationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetPlatformApplicationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSMSAttributes", func(t *testing.T) {
        input := &sns.SetSMSAttributesInput{}
        output := &sns.SetSMSAttributesOutput{}

        mockClient.On("SetSMSAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetSMSAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSubscriptionAttributes", func(t *testing.T) {
        input := &sns.SetSubscriptionAttributesInput{}
        output := &sns.SetSubscriptionAttributesOutput{}

        mockClient.On("SetSubscriptionAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetSubscriptionAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTopicAttributes", func(t *testing.T) {
        input := &sns.SetTopicAttributesInput{}
        output := &sns.SetTopicAttributesOutput{}

        mockClient.On("SetTopicAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.SetTopicAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubscribe", func(t *testing.T) {
        input := &sns.SubscribeInput{}
        output := &sns.SubscribeOutput{}

        mockClient.On("Subscribe", ctx, input).Return(output, nil)

        result, err := mockClient.Subscribe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &sns.TagResourceInput{}
        output := &sns.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnsubscribe", func(t *testing.T) {
        input := &sns.UnsubscribeInput{}
        output := &sns.UnsubscribeOutput{}

        mockClient.On("Unsubscribe", ctx, input).Return(output, nil)

        result, err := mockClient.Unsubscribe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &sns.UntagResourceInput{}
        output := &sns.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifySMSSandboxPhoneNumber", func(t *testing.T) {
        input := &sns.VerifySMSSandboxPhoneNumberInput{}
        output := &sns.VerifySMSSandboxPhoneNumberOutput{}

        mockClient.On("VerifySMSSandboxPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.VerifySMSSandboxPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
