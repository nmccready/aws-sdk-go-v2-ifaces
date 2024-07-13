package ses_test

// tests for the ses service interface for this ../../../service/ses/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ses/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ses/ses_iface"
	"github.com/stretchr/testify/assert"
)

func TestSesServiceCanBeMocked(t *testing.T) {
	var iface ses_iface.IClient
	iface = &ses.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ses.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloneReceiptRuleSet", func(t *testing.T) {
        input := &ses.CloneReceiptRuleSetInput{}
        output := &ses.CloneReceiptRuleSetOutput{}

        mockClient.On("CloneReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.CloneReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSet", func(t *testing.T) {
        input := &ses.CreateConfigurationSetInput{}
        output := &ses.CreateConfigurationSetOutput{}

        mockClient.On("CreateConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSetEventDestination", func(t *testing.T) {
        input := &ses.CreateConfigurationSetEventDestinationInput{}
        output := &ses.CreateConfigurationSetEventDestinationOutput{}

        mockClient.On("CreateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSetTrackingOptions", func(t *testing.T) {
        input := &ses.CreateConfigurationSetTrackingOptionsInput{}
        output := &ses.CreateConfigurationSetTrackingOptionsOutput{}

        mockClient.On("CreateConfigurationSetTrackingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSetTrackingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &ses.CreateCustomVerificationEmailTemplateInput{}
        output := &ses.CreateCustomVerificationEmailTemplateOutput{}

        mockClient.On("CreateCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReceiptFilter", func(t *testing.T) {
        input := &ses.CreateReceiptFilterInput{}
        output := &ses.CreateReceiptFilterOutput{}

        mockClient.On("CreateReceiptFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReceiptFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReceiptRule", func(t *testing.T) {
        input := &ses.CreateReceiptRuleInput{}
        output := &ses.CreateReceiptRuleOutput{}

        mockClient.On("CreateReceiptRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReceiptRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReceiptRuleSet", func(t *testing.T) {
        input := &ses.CreateReceiptRuleSetInput{}
        output := &ses.CreateReceiptRuleSetOutput{}

        mockClient.On("CreateReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplate", func(t *testing.T) {
        input := &ses.CreateTemplateInput{}
        output := &ses.CreateTemplateOutput{}

        mockClient.On("CreateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSet", func(t *testing.T) {
        input := &ses.DeleteConfigurationSetInput{}
        output := &ses.DeleteConfigurationSetOutput{}

        mockClient.On("DeleteConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSetEventDestination", func(t *testing.T) {
        input := &ses.DeleteConfigurationSetEventDestinationInput{}
        output := &ses.DeleteConfigurationSetEventDestinationOutput{}

        mockClient.On("DeleteConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSetTrackingOptions", func(t *testing.T) {
        input := &ses.DeleteConfigurationSetTrackingOptionsInput{}
        output := &ses.DeleteConfigurationSetTrackingOptionsOutput{}

        mockClient.On("DeleteConfigurationSetTrackingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSetTrackingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &ses.DeleteCustomVerificationEmailTemplateInput{}
        output := &ses.DeleteCustomVerificationEmailTemplateOutput{}

        mockClient.On("DeleteCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentity", func(t *testing.T) {
        input := &ses.DeleteIdentityInput{}
        output := &ses.DeleteIdentityOutput{}

        mockClient.On("DeleteIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentityPolicy", func(t *testing.T) {
        input := &ses.DeleteIdentityPolicyInput{}
        output := &ses.DeleteIdentityPolicyOutput{}

        mockClient.On("DeleteIdentityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReceiptFilter", func(t *testing.T) {
        input := &ses.DeleteReceiptFilterInput{}
        output := &ses.DeleteReceiptFilterOutput{}

        mockClient.On("DeleteReceiptFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReceiptFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReceiptRule", func(t *testing.T) {
        input := &ses.DeleteReceiptRuleInput{}
        output := &ses.DeleteReceiptRuleOutput{}

        mockClient.On("DeleteReceiptRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReceiptRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReceiptRuleSet", func(t *testing.T) {
        input := &ses.DeleteReceiptRuleSetInput{}
        output := &ses.DeleteReceiptRuleSetOutput{}

        mockClient.On("DeleteReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplate", func(t *testing.T) {
        input := &ses.DeleteTemplateInput{}
        output := &ses.DeleteTemplateOutput{}

        mockClient.On("DeleteTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVerifiedEmailAddress", func(t *testing.T) {
        input := &ses.DeleteVerifiedEmailAddressInput{}
        output := &ses.DeleteVerifiedEmailAddressOutput{}

        mockClient.On("DeleteVerifiedEmailAddress", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVerifiedEmailAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActiveReceiptRuleSet", func(t *testing.T) {
        input := &ses.DescribeActiveReceiptRuleSetInput{}
        output := &ses.DescribeActiveReceiptRuleSetOutput{}

        mockClient.On("DescribeActiveReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActiveReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationSet", func(t *testing.T) {
        input := &ses.DescribeConfigurationSetInput{}
        output := &ses.DescribeConfigurationSetOutput{}

        mockClient.On("DescribeConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReceiptRule", func(t *testing.T) {
        input := &ses.DescribeReceiptRuleInput{}
        output := &ses.DescribeReceiptRuleOutput{}

        mockClient.On("DescribeReceiptRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReceiptRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReceiptRuleSet", func(t *testing.T) {
        input := &ses.DescribeReceiptRuleSetInput{}
        output := &ses.DescribeReceiptRuleSetOutput{}

        mockClient.On("DescribeReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSendingEnabled", func(t *testing.T) {
        input := &ses.GetAccountSendingEnabledInput{}
        output := &ses.GetAccountSendingEnabledOutput{}

        mockClient.On("GetAccountSendingEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSendingEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &ses.GetCustomVerificationEmailTemplateInput{}
        output := &ses.GetCustomVerificationEmailTemplateOutput{}

        mockClient.On("GetCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityDkimAttributes", func(t *testing.T) {
        input := &ses.GetIdentityDkimAttributesInput{}
        output := &ses.GetIdentityDkimAttributesOutput{}

        mockClient.On("GetIdentityDkimAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityDkimAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityMailFromDomainAttributes", func(t *testing.T) {
        input := &ses.GetIdentityMailFromDomainAttributesInput{}
        output := &ses.GetIdentityMailFromDomainAttributesOutput{}

        mockClient.On("GetIdentityMailFromDomainAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityMailFromDomainAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityNotificationAttributes", func(t *testing.T) {
        input := &ses.GetIdentityNotificationAttributesInput{}
        output := &ses.GetIdentityNotificationAttributesOutput{}

        mockClient.On("GetIdentityNotificationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityNotificationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityPolicies", func(t *testing.T) {
        input := &ses.GetIdentityPoliciesInput{}
        output := &ses.GetIdentityPoliciesOutput{}

        mockClient.On("GetIdentityPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityVerificationAttributes", func(t *testing.T) {
        input := &ses.GetIdentityVerificationAttributesInput{}
        output := &ses.GetIdentityVerificationAttributesOutput{}

        mockClient.On("GetIdentityVerificationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityVerificationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSendQuota", func(t *testing.T) {
        input := &ses.GetSendQuotaInput{}
        output := &ses.GetSendQuotaOutput{}

        mockClient.On("GetSendQuota", ctx, input).Return(output, nil)

        result, err := mockClient.GetSendQuota(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSendStatistics", func(t *testing.T) {
        input := &ses.GetSendStatisticsInput{}
        output := &ses.GetSendStatisticsOutput{}

        mockClient.On("GetSendStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetSendStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplate", func(t *testing.T) {
        input := &ses.GetTemplateInput{}
        output := &ses.GetTemplateOutput{}

        mockClient.On("GetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationSets", func(t *testing.T) {
        input := &ses.ListConfigurationSetsInput{}
        output := &ses.ListConfigurationSetsOutput{}

        mockClient.On("ListConfigurationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomVerificationEmailTemplates", func(t *testing.T) {
        input := &ses.ListCustomVerificationEmailTemplatesInput{}
        output := &ses.ListCustomVerificationEmailTemplatesOutput{}

        mockClient.On("ListCustomVerificationEmailTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomVerificationEmailTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentities", func(t *testing.T) {
        input := &ses.ListIdentitiesInput{}
        output := &ses.ListIdentitiesOutput{}

        mockClient.On("ListIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityPolicies", func(t *testing.T) {
        input := &ses.ListIdentityPoliciesInput{}
        output := &ses.ListIdentityPoliciesOutput{}

        mockClient.On("ListIdentityPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceiptFilters", func(t *testing.T) {
        input := &ses.ListReceiptFiltersInput{}
        output := &ses.ListReceiptFiltersOutput{}

        mockClient.On("ListReceiptFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceiptFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceiptRuleSets", func(t *testing.T) {
        input := &ses.ListReceiptRuleSetsInput{}
        output := &ses.ListReceiptRuleSetsOutput{}

        mockClient.On("ListReceiptRuleSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceiptRuleSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &ses.ListTemplatesInput{}
        output := &ses.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVerifiedEmailAddresses", func(t *testing.T) {
        input := &ses.ListVerifiedEmailAddressesInput{}
        output := &ses.ListVerifiedEmailAddressesOutput{}

        mockClient.On("ListVerifiedEmailAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.ListVerifiedEmailAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetDeliveryOptions", func(t *testing.T) {
        input := &ses.PutConfigurationSetDeliveryOptionsInput{}
        output := &ses.PutConfigurationSetDeliveryOptionsOutput{}

        mockClient.On("PutConfigurationSetDeliveryOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetDeliveryOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutIdentityPolicy", func(t *testing.T) {
        input := &ses.PutIdentityPolicyInput{}
        output := &ses.PutIdentityPolicyOutput{}

        mockClient.On("PutIdentityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutIdentityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReorderReceiptRuleSet", func(t *testing.T) {
        input := &ses.ReorderReceiptRuleSetInput{}
        output := &ses.ReorderReceiptRuleSetOutput{}

        mockClient.On("ReorderReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.ReorderReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendBounce", func(t *testing.T) {
        input := &ses.SendBounceInput{}
        output := &ses.SendBounceOutput{}

        mockClient.On("SendBounce", ctx, input).Return(output, nil)

        result, err := mockClient.SendBounce(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendBulkTemplatedEmail", func(t *testing.T) {
        input := &ses.SendBulkTemplatedEmailInput{}
        output := &ses.SendBulkTemplatedEmailOutput{}

        mockClient.On("SendBulkTemplatedEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendBulkTemplatedEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCustomVerificationEmail", func(t *testing.T) {
        input := &ses.SendCustomVerificationEmailInput{}
        output := &ses.SendCustomVerificationEmailOutput{}

        mockClient.On("SendCustomVerificationEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendCustomVerificationEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEmail", func(t *testing.T) {
        input := &ses.SendEmailInput{}
        output := &ses.SendEmailOutput{}

        mockClient.On("SendEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendRawEmail", func(t *testing.T) {
        input := &ses.SendRawEmailInput{}
        output := &ses.SendRawEmailOutput{}

        mockClient.On("SendRawEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendRawEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTemplatedEmail", func(t *testing.T) {
        input := &ses.SendTemplatedEmailInput{}
        output := &ses.SendTemplatedEmailOutput{}

        mockClient.On("SendTemplatedEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendTemplatedEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetActiveReceiptRuleSet", func(t *testing.T) {
        input := &ses.SetActiveReceiptRuleSetInput{}
        output := &ses.SetActiveReceiptRuleSetOutput{}

        mockClient.On("SetActiveReceiptRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.SetActiveReceiptRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityDkimEnabled", func(t *testing.T) {
        input := &ses.SetIdentityDkimEnabledInput{}
        output := &ses.SetIdentityDkimEnabledOutput{}

        mockClient.On("SetIdentityDkimEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityDkimEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityFeedbackForwardingEnabled", func(t *testing.T) {
        input := &ses.SetIdentityFeedbackForwardingEnabledInput{}
        output := &ses.SetIdentityFeedbackForwardingEnabledOutput{}

        mockClient.On("SetIdentityFeedbackForwardingEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityFeedbackForwardingEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityHeadersInNotificationsEnabled", func(t *testing.T) {
        input := &ses.SetIdentityHeadersInNotificationsEnabledInput{}
        output := &ses.SetIdentityHeadersInNotificationsEnabledOutput{}

        mockClient.On("SetIdentityHeadersInNotificationsEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityHeadersInNotificationsEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityMailFromDomain", func(t *testing.T) {
        input := &ses.SetIdentityMailFromDomainInput{}
        output := &ses.SetIdentityMailFromDomainOutput{}

        mockClient.On("SetIdentityMailFromDomain", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityMailFromDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityNotificationTopic", func(t *testing.T) {
        input := &ses.SetIdentityNotificationTopicInput{}
        output := &ses.SetIdentityNotificationTopicOutput{}

        mockClient.On("SetIdentityNotificationTopic", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityNotificationTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetReceiptRulePosition", func(t *testing.T) {
        input := &ses.SetReceiptRulePositionInput{}
        output := &ses.SetReceiptRulePositionOutput{}

        mockClient.On("SetReceiptRulePosition", ctx, input).Return(output, nil)

        result, err := mockClient.SetReceiptRulePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestRenderTemplate", func(t *testing.T) {
        input := &ses.TestRenderTemplateInput{}
        output := &ses.TestRenderTemplateOutput{}

        mockClient.On("TestRenderTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.TestRenderTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSendingEnabled", func(t *testing.T) {
        input := &ses.UpdateAccountSendingEnabledInput{}
        output := &ses.UpdateAccountSendingEnabledOutput{}

        mockClient.On("UpdateAccountSendingEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSendingEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetEventDestination", func(t *testing.T) {
        input := &ses.UpdateConfigurationSetEventDestinationInput{}
        output := &ses.UpdateConfigurationSetEventDestinationOutput{}

        mockClient.On("UpdateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetReputationMetricsEnabled", func(t *testing.T) {
        input := &ses.UpdateConfigurationSetReputationMetricsEnabledInput{}
        output := &ses.UpdateConfigurationSetReputationMetricsEnabledOutput{}

        mockClient.On("UpdateConfigurationSetReputationMetricsEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetReputationMetricsEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetSendingEnabled", func(t *testing.T) {
        input := &ses.UpdateConfigurationSetSendingEnabledInput{}
        output := &ses.UpdateConfigurationSetSendingEnabledOutput{}

        mockClient.On("UpdateConfigurationSetSendingEnabled", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetSendingEnabled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetTrackingOptions", func(t *testing.T) {
        input := &ses.UpdateConfigurationSetTrackingOptionsInput{}
        output := &ses.UpdateConfigurationSetTrackingOptionsOutput{}

        mockClient.On("UpdateConfigurationSetTrackingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetTrackingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &ses.UpdateCustomVerificationEmailTemplateInput{}
        output := &ses.UpdateCustomVerificationEmailTemplateOutput{}

        mockClient.On("UpdateCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReceiptRule", func(t *testing.T) {
        input := &ses.UpdateReceiptRuleInput{}
        output := &ses.UpdateReceiptRuleOutput{}

        mockClient.On("UpdateReceiptRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReceiptRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplate", func(t *testing.T) {
        input := &ses.UpdateTemplateInput{}
        output := &ses.UpdateTemplateOutput{}

        mockClient.On("UpdateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyDomainDkim", func(t *testing.T) {
        input := &ses.VerifyDomainDkimInput{}
        output := &ses.VerifyDomainDkimOutput{}

        mockClient.On("VerifyDomainDkim", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyDomainDkim(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyDomainIdentity", func(t *testing.T) {
        input := &ses.VerifyDomainIdentityInput{}
        output := &ses.VerifyDomainIdentityOutput{}

        mockClient.On("VerifyDomainIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyDomainIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyEmailAddress", func(t *testing.T) {
        input := &ses.VerifyEmailAddressInput{}
        output := &ses.VerifyEmailAddressOutput{}

        mockClient.On("VerifyEmailAddress", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyEmailAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyEmailIdentity", func(t *testing.T) {
        input := &ses.VerifyEmailIdentityInput{}
        output := &ses.VerifyEmailIdentityOutput{}

        mockClient.On("VerifyEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
