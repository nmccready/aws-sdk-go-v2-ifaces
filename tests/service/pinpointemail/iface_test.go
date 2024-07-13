package pinpointemail_test

// tests for the pinpointemail service interface for this ../../../service/pinpointemail/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointemail/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpointemail/pinpointemail_iface"
	"github.com/stretchr/testify/assert"
)

func TestPinpointemailServiceCanBeMocked(t *testing.T) {
	var iface pinpointemail_iface.IClient
	iface = &pinpointemail.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pinpointemail.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSet", func(t *testing.T) {
        input := &pinpointemail.CreateConfigurationSetInput{}
        output := &pinpointemail.CreateConfigurationSetOutput{}

        mockClient.On("CreateConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointemail.CreateConfigurationSetEventDestinationInput{}
        output := &pinpointemail.CreateConfigurationSetEventDestinationOutput{}

        mockClient.On("CreateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDedicatedIpPool", func(t *testing.T) {
        input := &pinpointemail.CreateDedicatedIpPoolInput{}
        output := &pinpointemail.CreateDedicatedIpPoolOutput{}

        mockClient.On("CreateDedicatedIpPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDedicatedIpPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeliverabilityTestReport", func(t *testing.T) {
        input := &pinpointemail.CreateDeliverabilityTestReportInput{}
        output := &pinpointemail.CreateDeliverabilityTestReportOutput{}

        mockClient.On("CreateDeliverabilityTestReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeliverabilityTestReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEmailIdentity", func(t *testing.T) {
        input := &pinpointemail.CreateEmailIdentityInput{}
        output := &pinpointemail.CreateEmailIdentityOutput{}

        mockClient.On("CreateEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSet", func(t *testing.T) {
        input := &pinpointemail.DeleteConfigurationSetInput{}
        output := &pinpointemail.DeleteConfigurationSetOutput{}

        mockClient.On("DeleteConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointemail.DeleteConfigurationSetEventDestinationInput{}
        output := &pinpointemail.DeleteConfigurationSetEventDestinationOutput{}

        mockClient.On("DeleteConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDedicatedIpPool", func(t *testing.T) {
        input := &pinpointemail.DeleteDedicatedIpPoolInput{}
        output := &pinpointemail.DeleteDedicatedIpPoolOutput{}

        mockClient.On("DeleteDedicatedIpPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDedicatedIpPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailIdentity", func(t *testing.T) {
        input := &pinpointemail.DeleteEmailIdentityInput{}
        output := &pinpointemail.DeleteEmailIdentityOutput{}

        mockClient.On("DeleteEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccount", func(t *testing.T) {
        input := &pinpointemail.GetAccountInput{}
        output := &pinpointemail.GetAccountOutput{}

        mockClient.On("GetAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlacklistReports", func(t *testing.T) {
        input := &pinpointemail.GetBlacklistReportsInput{}
        output := &pinpointemail.GetBlacklistReportsOutput{}

        mockClient.On("GetBlacklistReports", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlacklistReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationSet", func(t *testing.T) {
        input := &pinpointemail.GetConfigurationSetInput{}
        output := &pinpointemail.GetConfigurationSetOutput{}

        mockClient.On("GetConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationSetEventDestinations", func(t *testing.T) {
        input := &pinpointemail.GetConfigurationSetEventDestinationsInput{}
        output := &pinpointemail.GetConfigurationSetEventDestinationsOutput{}

        mockClient.On("GetConfigurationSetEventDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationSetEventDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDedicatedIp", func(t *testing.T) {
        input := &pinpointemail.GetDedicatedIpInput{}
        output := &pinpointemail.GetDedicatedIpOutput{}

        mockClient.On("GetDedicatedIp", ctx, input).Return(output, nil)

        result, err := mockClient.GetDedicatedIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDedicatedIps", func(t *testing.T) {
        input := &pinpointemail.GetDedicatedIpsInput{}
        output := &pinpointemail.GetDedicatedIpsOutput{}

        mockClient.On("GetDedicatedIps", ctx, input).Return(output, nil)

        result, err := mockClient.GetDedicatedIps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliverabilityDashboardOptions", func(t *testing.T) {
        input := &pinpointemail.GetDeliverabilityDashboardOptionsInput{}
        output := &pinpointemail.GetDeliverabilityDashboardOptionsOutput{}

        mockClient.On("GetDeliverabilityDashboardOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliverabilityDashboardOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliverabilityTestReport", func(t *testing.T) {
        input := &pinpointemail.GetDeliverabilityTestReportInput{}
        output := &pinpointemail.GetDeliverabilityTestReportOutput{}

        mockClient.On("GetDeliverabilityTestReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliverabilityTestReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainDeliverabilityCampaign", func(t *testing.T) {
        input := &pinpointemail.GetDomainDeliverabilityCampaignInput{}
        output := &pinpointemail.GetDomainDeliverabilityCampaignOutput{}

        mockClient.On("GetDomainDeliverabilityCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainDeliverabilityCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainStatisticsReport", func(t *testing.T) {
        input := &pinpointemail.GetDomainStatisticsReportInput{}
        output := &pinpointemail.GetDomainStatisticsReportOutput{}

        mockClient.On("GetDomainStatisticsReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainStatisticsReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailIdentity", func(t *testing.T) {
        input := &pinpointemail.GetEmailIdentityInput{}
        output := &pinpointemail.GetEmailIdentityOutput{}

        mockClient.On("GetEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationSets", func(t *testing.T) {
        input := &pinpointemail.ListConfigurationSetsInput{}
        output := &pinpointemail.ListConfigurationSetsOutput{}

        mockClient.On("ListConfigurationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDedicatedIpPools", func(t *testing.T) {
        input := &pinpointemail.ListDedicatedIpPoolsInput{}
        output := &pinpointemail.ListDedicatedIpPoolsOutput{}

        mockClient.On("ListDedicatedIpPools", ctx, input).Return(output, nil)

        result, err := mockClient.ListDedicatedIpPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeliverabilityTestReports", func(t *testing.T) {
        input := &pinpointemail.ListDeliverabilityTestReportsInput{}
        output := &pinpointemail.ListDeliverabilityTestReportsOutput{}

        mockClient.On("ListDeliverabilityTestReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeliverabilityTestReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainDeliverabilityCampaigns", func(t *testing.T) {
        input := &pinpointemail.ListDomainDeliverabilityCampaignsInput{}
        output := &pinpointemail.ListDomainDeliverabilityCampaignsOutput{}

        mockClient.On("ListDomainDeliverabilityCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainDeliverabilityCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEmailIdentities", func(t *testing.T) {
        input := &pinpointemail.ListEmailIdentitiesInput{}
        output := &pinpointemail.ListEmailIdentitiesOutput{}

        mockClient.On("ListEmailIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListEmailIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pinpointemail.ListTagsForResourceInput{}
        output := &pinpointemail.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountDedicatedIpWarmupAttributes", func(t *testing.T) {
        input := &pinpointemail.PutAccountDedicatedIpWarmupAttributesInput{}
        output := &pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput{}

        mockClient.On("PutAccountDedicatedIpWarmupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountDedicatedIpWarmupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSendingAttributes", func(t *testing.T) {
        input := &pinpointemail.PutAccountSendingAttributesInput{}
        output := &pinpointemail.PutAccountSendingAttributesOutput{}

        mockClient.On("PutAccountSendingAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSendingAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetDeliveryOptions", func(t *testing.T) {
        input := &pinpointemail.PutConfigurationSetDeliveryOptionsInput{}
        output := &pinpointemail.PutConfigurationSetDeliveryOptionsOutput{}

        mockClient.On("PutConfigurationSetDeliveryOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetDeliveryOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetReputationOptions", func(t *testing.T) {
        input := &pinpointemail.PutConfigurationSetReputationOptionsInput{}
        output := &pinpointemail.PutConfigurationSetReputationOptionsOutput{}

        mockClient.On("PutConfigurationSetReputationOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetReputationOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetSendingOptions", func(t *testing.T) {
        input := &pinpointemail.PutConfigurationSetSendingOptionsInput{}
        output := &pinpointemail.PutConfigurationSetSendingOptionsOutput{}

        mockClient.On("PutConfigurationSetSendingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetSendingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetTrackingOptions", func(t *testing.T) {
        input := &pinpointemail.PutConfigurationSetTrackingOptionsInput{}
        output := &pinpointemail.PutConfigurationSetTrackingOptionsOutput{}

        mockClient.On("PutConfigurationSetTrackingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetTrackingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDedicatedIpInPool", func(t *testing.T) {
        input := &pinpointemail.PutDedicatedIpInPoolInput{}
        output := &pinpointemail.PutDedicatedIpInPoolOutput{}

        mockClient.On("PutDedicatedIpInPool", ctx, input).Return(output, nil)

        result, err := mockClient.PutDedicatedIpInPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDedicatedIpWarmupAttributes", func(t *testing.T) {
        input := &pinpointemail.PutDedicatedIpWarmupAttributesInput{}
        output := &pinpointemail.PutDedicatedIpWarmupAttributesOutput{}

        mockClient.On("PutDedicatedIpWarmupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutDedicatedIpWarmupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliverabilityDashboardOption", func(t *testing.T) {
        input := &pinpointemail.PutDeliverabilityDashboardOptionInput{}
        output := &pinpointemail.PutDeliverabilityDashboardOptionOutput{}

        mockClient.On("PutDeliverabilityDashboardOption", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliverabilityDashboardOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityDkimAttributes", func(t *testing.T) {
        input := &pinpointemail.PutEmailIdentityDkimAttributesInput{}
        output := &pinpointemail.PutEmailIdentityDkimAttributesOutput{}

        mockClient.On("PutEmailIdentityDkimAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityDkimAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityFeedbackAttributes", func(t *testing.T) {
        input := &pinpointemail.PutEmailIdentityFeedbackAttributesInput{}
        output := &pinpointemail.PutEmailIdentityFeedbackAttributesOutput{}

        mockClient.On("PutEmailIdentityFeedbackAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityFeedbackAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityMailFromAttributes", func(t *testing.T) {
        input := &pinpointemail.PutEmailIdentityMailFromAttributesInput{}
        output := &pinpointemail.PutEmailIdentityMailFromAttributesOutput{}

        mockClient.On("PutEmailIdentityMailFromAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityMailFromAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEmail", func(t *testing.T) {
        input := &pinpointemail.SendEmailInput{}
        output := &pinpointemail.SendEmailOutput{}

        mockClient.On("SendEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pinpointemail.TagResourceInput{}
        output := &pinpointemail.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pinpointemail.UntagResourceInput{}
        output := &pinpointemail.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetEventDestination", func(t *testing.T) {
        input := &pinpointemail.UpdateConfigurationSetEventDestinationInput{}
        output := &pinpointemail.UpdateConfigurationSetEventDestinationOutput{}

        mockClient.On("UpdateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
