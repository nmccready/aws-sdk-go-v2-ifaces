package sesv2_test

// tests for the sesv2 service interface for this ../../../service/sesv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sesv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sesv2/sesv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestSesv2ServiceCanBeMocked(t *testing.T) {
	var iface sesv2_iface.IClient
	iface = &sesv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sesv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetMetricData", func(t *testing.T) {
        input := &sesv2.BatchGetMetricDataInput{}
        output := &sesv2.BatchGetMetricDataOutput{}

        mockClient.On("BatchGetMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelExportJob", func(t *testing.T) {
        input := &sesv2.CancelExportJobInput{}
        output := &sesv2.CancelExportJobOutput{}

        mockClient.On("CancelExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSet", func(t *testing.T) {
        input := &sesv2.CreateConfigurationSetInput{}
        output := &sesv2.CreateConfigurationSetOutput{}

        mockClient.On("CreateConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationSetEventDestination", func(t *testing.T) {
        input := &sesv2.CreateConfigurationSetEventDestinationInput{}
        output := &sesv2.CreateConfigurationSetEventDestinationOutput{}

        mockClient.On("CreateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContact", func(t *testing.T) {
        input := &sesv2.CreateContactInput{}
        output := &sesv2.CreateContactOutput{}

        mockClient.On("CreateContact", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContactList", func(t *testing.T) {
        input := &sesv2.CreateContactListInput{}
        output := &sesv2.CreateContactListOutput{}

        mockClient.On("CreateContactList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContactList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &sesv2.CreateCustomVerificationEmailTemplateInput{}
        output := &sesv2.CreateCustomVerificationEmailTemplateOutput{}

        mockClient.On("CreateCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDedicatedIpPool", func(t *testing.T) {
        input := &sesv2.CreateDedicatedIpPoolInput{}
        output := &sesv2.CreateDedicatedIpPoolOutput{}

        mockClient.On("CreateDedicatedIpPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDedicatedIpPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeliverabilityTestReport", func(t *testing.T) {
        input := &sesv2.CreateDeliverabilityTestReportInput{}
        output := &sesv2.CreateDeliverabilityTestReportOutput{}

        mockClient.On("CreateDeliverabilityTestReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeliverabilityTestReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEmailIdentity", func(t *testing.T) {
        input := &sesv2.CreateEmailIdentityInput{}
        output := &sesv2.CreateEmailIdentityOutput{}

        mockClient.On("CreateEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEmailIdentityPolicy", func(t *testing.T) {
        input := &sesv2.CreateEmailIdentityPolicyInput{}
        output := &sesv2.CreateEmailIdentityPolicyOutput{}

        mockClient.On("CreateEmailIdentityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEmailIdentityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEmailTemplate", func(t *testing.T) {
        input := &sesv2.CreateEmailTemplateInput{}
        output := &sesv2.CreateEmailTemplateOutput{}

        mockClient.On("CreateEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExportJob", func(t *testing.T) {
        input := &sesv2.CreateExportJobInput{}
        output := &sesv2.CreateExportJobOutput{}

        mockClient.On("CreateExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImportJob", func(t *testing.T) {
        input := &sesv2.CreateImportJobInput{}
        output := &sesv2.CreateImportJobOutput{}

        mockClient.On("CreateImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSet", func(t *testing.T) {
        input := &sesv2.DeleteConfigurationSetInput{}
        output := &sesv2.DeleteConfigurationSetOutput{}

        mockClient.On("DeleteConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationSetEventDestination", func(t *testing.T) {
        input := &sesv2.DeleteConfigurationSetEventDestinationInput{}
        output := &sesv2.DeleteConfigurationSetEventDestinationOutput{}

        mockClient.On("DeleteConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContact", func(t *testing.T) {
        input := &sesv2.DeleteContactInput{}
        output := &sesv2.DeleteContactOutput{}

        mockClient.On("DeleteContact", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactList", func(t *testing.T) {
        input := &sesv2.DeleteContactListInput{}
        output := &sesv2.DeleteContactListOutput{}

        mockClient.On("DeleteContactList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &sesv2.DeleteCustomVerificationEmailTemplateInput{}
        output := &sesv2.DeleteCustomVerificationEmailTemplateOutput{}

        mockClient.On("DeleteCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDedicatedIpPool", func(t *testing.T) {
        input := &sesv2.DeleteDedicatedIpPoolInput{}
        output := &sesv2.DeleteDedicatedIpPoolOutput{}

        mockClient.On("DeleteDedicatedIpPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDedicatedIpPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailIdentity", func(t *testing.T) {
        input := &sesv2.DeleteEmailIdentityInput{}
        output := &sesv2.DeleteEmailIdentityOutput{}

        mockClient.On("DeleteEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailIdentityPolicy", func(t *testing.T) {
        input := &sesv2.DeleteEmailIdentityPolicyInput{}
        output := &sesv2.DeleteEmailIdentityPolicyOutput{}

        mockClient.On("DeleteEmailIdentityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailIdentityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailTemplate", func(t *testing.T) {
        input := &sesv2.DeleteEmailTemplateInput{}
        output := &sesv2.DeleteEmailTemplateOutput{}

        mockClient.On("DeleteEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSuppressedDestination", func(t *testing.T) {
        input := &sesv2.DeleteSuppressedDestinationInput{}
        output := &sesv2.DeleteSuppressedDestinationOutput{}

        mockClient.On("DeleteSuppressedDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSuppressedDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccount", func(t *testing.T) {
        input := &sesv2.GetAccountInput{}
        output := &sesv2.GetAccountOutput{}

        mockClient.On("GetAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlacklistReports", func(t *testing.T) {
        input := &sesv2.GetBlacklistReportsInput{}
        output := &sesv2.GetBlacklistReportsOutput{}

        mockClient.On("GetBlacklistReports", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlacklistReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationSet", func(t *testing.T) {
        input := &sesv2.GetConfigurationSetInput{}
        output := &sesv2.GetConfigurationSetOutput{}

        mockClient.On("GetConfigurationSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationSetEventDestinations", func(t *testing.T) {
        input := &sesv2.GetConfigurationSetEventDestinationsInput{}
        output := &sesv2.GetConfigurationSetEventDestinationsOutput{}

        mockClient.On("GetConfigurationSetEventDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationSetEventDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContact", func(t *testing.T) {
        input := &sesv2.GetContactInput{}
        output := &sesv2.GetContactOutput{}

        mockClient.On("GetContact", ctx, input).Return(output, nil)

        result, err := mockClient.GetContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactList", func(t *testing.T) {
        input := &sesv2.GetContactListInput{}
        output := &sesv2.GetContactListOutput{}

        mockClient.On("GetContactList", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &sesv2.GetCustomVerificationEmailTemplateInput{}
        output := &sesv2.GetCustomVerificationEmailTemplateOutput{}

        mockClient.On("GetCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDedicatedIp", func(t *testing.T) {
        input := &sesv2.GetDedicatedIpInput{}
        output := &sesv2.GetDedicatedIpOutput{}

        mockClient.On("GetDedicatedIp", ctx, input).Return(output, nil)

        result, err := mockClient.GetDedicatedIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDedicatedIpPool", func(t *testing.T) {
        input := &sesv2.GetDedicatedIpPoolInput{}
        output := &sesv2.GetDedicatedIpPoolOutput{}

        mockClient.On("GetDedicatedIpPool", ctx, input).Return(output, nil)

        result, err := mockClient.GetDedicatedIpPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDedicatedIps", func(t *testing.T) {
        input := &sesv2.GetDedicatedIpsInput{}
        output := &sesv2.GetDedicatedIpsOutput{}

        mockClient.On("GetDedicatedIps", ctx, input).Return(output, nil)

        result, err := mockClient.GetDedicatedIps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliverabilityDashboardOptions", func(t *testing.T) {
        input := &sesv2.GetDeliverabilityDashboardOptionsInput{}
        output := &sesv2.GetDeliverabilityDashboardOptionsOutput{}

        mockClient.On("GetDeliverabilityDashboardOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliverabilityDashboardOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliverabilityTestReport", func(t *testing.T) {
        input := &sesv2.GetDeliverabilityTestReportInput{}
        output := &sesv2.GetDeliverabilityTestReportOutput{}

        mockClient.On("GetDeliverabilityTestReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliverabilityTestReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainDeliverabilityCampaign", func(t *testing.T) {
        input := &sesv2.GetDomainDeliverabilityCampaignInput{}
        output := &sesv2.GetDomainDeliverabilityCampaignOutput{}

        mockClient.On("GetDomainDeliverabilityCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainDeliverabilityCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainStatisticsReport", func(t *testing.T) {
        input := &sesv2.GetDomainStatisticsReportInput{}
        output := &sesv2.GetDomainStatisticsReportOutput{}

        mockClient.On("GetDomainStatisticsReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainStatisticsReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailIdentity", func(t *testing.T) {
        input := &sesv2.GetEmailIdentityInput{}
        output := &sesv2.GetEmailIdentityOutput{}

        mockClient.On("GetEmailIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailIdentityPolicies", func(t *testing.T) {
        input := &sesv2.GetEmailIdentityPoliciesInput{}
        output := &sesv2.GetEmailIdentityPoliciesOutput{}

        mockClient.On("GetEmailIdentityPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailIdentityPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailTemplate", func(t *testing.T) {
        input := &sesv2.GetEmailTemplateInput{}
        output := &sesv2.GetEmailTemplateOutput{}

        mockClient.On("GetEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExportJob", func(t *testing.T) {
        input := &sesv2.GetExportJobInput{}
        output := &sesv2.GetExportJobOutput{}

        mockClient.On("GetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportJob", func(t *testing.T) {
        input := &sesv2.GetImportJobInput{}
        output := &sesv2.GetImportJobOutput{}

        mockClient.On("GetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMessageInsights", func(t *testing.T) {
        input := &sesv2.GetMessageInsightsInput{}
        output := &sesv2.GetMessageInsightsOutput{}

        mockClient.On("GetMessageInsights", ctx, input).Return(output, nil)

        result, err := mockClient.GetMessageInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSuppressedDestination", func(t *testing.T) {
        input := &sesv2.GetSuppressedDestinationInput{}
        output := &sesv2.GetSuppressedDestinationOutput{}

        mockClient.On("GetSuppressedDestination", ctx, input).Return(output, nil)

        result, err := mockClient.GetSuppressedDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationSets", func(t *testing.T) {
        input := &sesv2.ListConfigurationSetsInput{}
        output := &sesv2.ListConfigurationSetsOutput{}

        mockClient.On("ListConfigurationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContactLists", func(t *testing.T) {
        input := &sesv2.ListContactListsInput{}
        output := &sesv2.ListContactListsOutput{}

        mockClient.On("ListContactLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListContactLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContacts", func(t *testing.T) {
        input := &sesv2.ListContactsInput{}
        output := &sesv2.ListContactsOutput{}

        mockClient.On("ListContacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListContacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomVerificationEmailTemplates", func(t *testing.T) {
        input := &sesv2.ListCustomVerificationEmailTemplatesInput{}
        output := &sesv2.ListCustomVerificationEmailTemplatesOutput{}

        mockClient.On("ListCustomVerificationEmailTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomVerificationEmailTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDedicatedIpPools", func(t *testing.T) {
        input := &sesv2.ListDedicatedIpPoolsInput{}
        output := &sesv2.ListDedicatedIpPoolsOutput{}

        mockClient.On("ListDedicatedIpPools", ctx, input).Return(output, nil)

        result, err := mockClient.ListDedicatedIpPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeliverabilityTestReports", func(t *testing.T) {
        input := &sesv2.ListDeliverabilityTestReportsInput{}
        output := &sesv2.ListDeliverabilityTestReportsOutput{}

        mockClient.On("ListDeliverabilityTestReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeliverabilityTestReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainDeliverabilityCampaigns", func(t *testing.T) {
        input := &sesv2.ListDomainDeliverabilityCampaignsInput{}
        output := &sesv2.ListDomainDeliverabilityCampaignsOutput{}

        mockClient.On("ListDomainDeliverabilityCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainDeliverabilityCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEmailIdentities", func(t *testing.T) {
        input := &sesv2.ListEmailIdentitiesInput{}
        output := &sesv2.ListEmailIdentitiesOutput{}

        mockClient.On("ListEmailIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListEmailIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEmailTemplates", func(t *testing.T) {
        input := &sesv2.ListEmailTemplatesInput{}
        output := &sesv2.ListEmailTemplatesOutput{}

        mockClient.On("ListEmailTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListEmailTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExportJobs", func(t *testing.T) {
        input := &sesv2.ListExportJobsInput{}
        output := &sesv2.ListExportJobsOutput{}

        mockClient.On("ListExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportJobs", func(t *testing.T) {
        input := &sesv2.ListImportJobsInput{}
        output := &sesv2.ListImportJobsOutput{}

        mockClient.On("ListImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &sesv2.ListRecommendationsInput{}
        output := &sesv2.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSuppressedDestinations", func(t *testing.T) {
        input := &sesv2.ListSuppressedDestinationsInput{}
        output := &sesv2.ListSuppressedDestinationsOutput{}

        mockClient.On("ListSuppressedDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSuppressedDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &sesv2.ListTagsForResourceInput{}
        output := &sesv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountDedicatedIpWarmupAttributes", func(t *testing.T) {
        input := &sesv2.PutAccountDedicatedIpWarmupAttributesInput{}
        output := &sesv2.PutAccountDedicatedIpWarmupAttributesOutput{}

        mockClient.On("PutAccountDedicatedIpWarmupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountDedicatedIpWarmupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountDetails", func(t *testing.T) {
        input := &sesv2.PutAccountDetailsInput{}
        output := &sesv2.PutAccountDetailsOutput{}

        mockClient.On("PutAccountDetails", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSendingAttributes", func(t *testing.T) {
        input := &sesv2.PutAccountSendingAttributesInput{}
        output := &sesv2.PutAccountSendingAttributesOutput{}

        mockClient.On("PutAccountSendingAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSendingAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSuppressionAttributes", func(t *testing.T) {
        input := &sesv2.PutAccountSuppressionAttributesInput{}
        output := &sesv2.PutAccountSuppressionAttributesOutput{}

        mockClient.On("PutAccountSuppressionAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSuppressionAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountVdmAttributes", func(t *testing.T) {
        input := &sesv2.PutAccountVdmAttributesInput{}
        output := &sesv2.PutAccountVdmAttributesOutput{}

        mockClient.On("PutAccountVdmAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountVdmAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetDeliveryOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetDeliveryOptionsInput{}
        output := &sesv2.PutConfigurationSetDeliveryOptionsOutput{}

        mockClient.On("PutConfigurationSetDeliveryOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetDeliveryOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetReputationOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetReputationOptionsInput{}
        output := &sesv2.PutConfigurationSetReputationOptionsOutput{}

        mockClient.On("PutConfigurationSetReputationOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetReputationOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetSendingOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetSendingOptionsInput{}
        output := &sesv2.PutConfigurationSetSendingOptionsOutput{}

        mockClient.On("PutConfigurationSetSendingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetSendingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetSuppressionOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetSuppressionOptionsInput{}
        output := &sesv2.PutConfigurationSetSuppressionOptionsOutput{}

        mockClient.On("PutConfigurationSetSuppressionOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetSuppressionOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetTrackingOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetTrackingOptionsInput{}
        output := &sesv2.PutConfigurationSetTrackingOptionsOutput{}

        mockClient.On("PutConfigurationSetTrackingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetTrackingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConfigurationSetVdmOptions", func(t *testing.T) {
        input := &sesv2.PutConfigurationSetVdmOptionsInput{}
        output := &sesv2.PutConfigurationSetVdmOptionsOutput{}

        mockClient.On("PutConfigurationSetVdmOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutConfigurationSetVdmOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDedicatedIpInPool", func(t *testing.T) {
        input := &sesv2.PutDedicatedIpInPoolInput{}
        output := &sesv2.PutDedicatedIpInPoolOutput{}

        mockClient.On("PutDedicatedIpInPool", ctx, input).Return(output, nil)

        result, err := mockClient.PutDedicatedIpInPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDedicatedIpPoolScalingAttributes", func(t *testing.T) {
        input := &sesv2.PutDedicatedIpPoolScalingAttributesInput{}
        output := &sesv2.PutDedicatedIpPoolScalingAttributesOutput{}

        mockClient.On("PutDedicatedIpPoolScalingAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutDedicatedIpPoolScalingAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDedicatedIpWarmupAttributes", func(t *testing.T) {
        input := &sesv2.PutDedicatedIpWarmupAttributesInput{}
        output := &sesv2.PutDedicatedIpWarmupAttributesOutput{}

        mockClient.On("PutDedicatedIpWarmupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutDedicatedIpWarmupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliverabilityDashboardOption", func(t *testing.T) {
        input := &sesv2.PutDeliverabilityDashboardOptionInput{}
        output := &sesv2.PutDeliverabilityDashboardOptionOutput{}

        mockClient.On("PutDeliverabilityDashboardOption", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliverabilityDashboardOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityConfigurationSetAttributes", func(t *testing.T) {
        input := &sesv2.PutEmailIdentityConfigurationSetAttributesInput{}
        output := &sesv2.PutEmailIdentityConfigurationSetAttributesOutput{}

        mockClient.On("PutEmailIdentityConfigurationSetAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityConfigurationSetAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityDkimAttributes", func(t *testing.T) {
        input := &sesv2.PutEmailIdentityDkimAttributesInput{}
        output := &sesv2.PutEmailIdentityDkimAttributesOutput{}

        mockClient.On("PutEmailIdentityDkimAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityDkimAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityDkimSigningAttributes", func(t *testing.T) {
        input := &sesv2.PutEmailIdentityDkimSigningAttributesInput{}
        output := &sesv2.PutEmailIdentityDkimSigningAttributesOutput{}

        mockClient.On("PutEmailIdentityDkimSigningAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityDkimSigningAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityFeedbackAttributes", func(t *testing.T) {
        input := &sesv2.PutEmailIdentityFeedbackAttributesInput{}
        output := &sesv2.PutEmailIdentityFeedbackAttributesOutput{}

        mockClient.On("PutEmailIdentityFeedbackAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityFeedbackAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailIdentityMailFromAttributes", func(t *testing.T) {
        input := &sesv2.PutEmailIdentityMailFromAttributesInput{}
        output := &sesv2.PutEmailIdentityMailFromAttributesOutput{}

        mockClient.On("PutEmailIdentityMailFromAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailIdentityMailFromAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSuppressedDestination", func(t *testing.T) {
        input := &sesv2.PutSuppressedDestinationInput{}
        output := &sesv2.PutSuppressedDestinationOutput{}

        mockClient.On("PutSuppressedDestination", ctx, input).Return(output, nil)

        result, err := mockClient.PutSuppressedDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendBulkEmail", func(t *testing.T) {
        input := &sesv2.SendBulkEmailInput{}
        output := &sesv2.SendBulkEmailOutput{}

        mockClient.On("SendBulkEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendBulkEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCustomVerificationEmail", func(t *testing.T) {
        input := &sesv2.SendCustomVerificationEmailInput{}
        output := &sesv2.SendCustomVerificationEmailOutput{}

        mockClient.On("SendCustomVerificationEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendCustomVerificationEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEmail", func(t *testing.T) {
        input := &sesv2.SendEmailInput{}
        output := &sesv2.SendEmailOutput{}

        mockClient.On("SendEmail", ctx, input).Return(output, nil)

        result, err := mockClient.SendEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &sesv2.TagResourceInput{}
        output := &sesv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestRenderEmailTemplate", func(t *testing.T) {
        input := &sesv2.TestRenderEmailTemplateInput{}
        output := &sesv2.TestRenderEmailTemplateOutput{}

        mockClient.On("TestRenderEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.TestRenderEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &sesv2.UntagResourceInput{}
        output := &sesv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationSetEventDestination", func(t *testing.T) {
        input := &sesv2.UpdateConfigurationSetEventDestinationInput{}
        output := &sesv2.UpdateConfigurationSetEventDestinationOutput{}

        mockClient.On("UpdateConfigurationSetEventDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationSetEventDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContact", func(t *testing.T) {
        input := &sesv2.UpdateContactInput{}
        output := &sesv2.UpdateContactOutput{}

        mockClient.On("UpdateContact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContactList", func(t *testing.T) {
        input := &sesv2.UpdateContactListInput{}
        output := &sesv2.UpdateContactListOutput{}

        mockClient.On("UpdateContactList", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContactList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomVerificationEmailTemplate", func(t *testing.T) {
        input := &sesv2.UpdateCustomVerificationEmailTemplateInput{}
        output := &sesv2.UpdateCustomVerificationEmailTemplateOutput{}

        mockClient.On("UpdateCustomVerificationEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomVerificationEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEmailIdentityPolicy", func(t *testing.T) {
        input := &sesv2.UpdateEmailIdentityPolicyInput{}
        output := &sesv2.UpdateEmailIdentityPolicyOutput{}

        mockClient.On("UpdateEmailIdentityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEmailIdentityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEmailTemplate", func(t *testing.T) {
        input := &sesv2.UpdateEmailTemplateInput{}
        output := &sesv2.UpdateEmailTemplateOutput{}

        mockClient.On("UpdateEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
