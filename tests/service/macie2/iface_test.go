package macie2_test

// tests for the macie2 service interface for this ../../../service/macie2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/macie2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/macie2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/macie2/macie2_iface"
	"github.com/stretchr/testify/assert"
)

func TestMacie2ServiceCanBeMocked(t *testing.T) {
	var iface macie2_iface.IClient
	iface = &macie2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := macie2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInvitation", func(t *testing.T) {
        input := &macie2.AcceptInvitationInput{}
        output := &macie2.AcceptInvitationOutput{}

        mockClient.On("AcceptInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCustomDataIdentifiers", func(t *testing.T) {
        input := &macie2.BatchGetCustomDataIdentifiersInput{}
        output := &macie2.BatchGetCustomDataIdentifiersOutput{}

        mockClient.On("BatchGetCustomDataIdentifiers", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCustomDataIdentifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateAutomatedDiscoveryAccounts", func(t *testing.T) {
        input := &macie2.BatchUpdateAutomatedDiscoveryAccountsInput{}
        output := &macie2.BatchUpdateAutomatedDiscoveryAccountsOutput{}

        mockClient.On("BatchUpdateAutomatedDiscoveryAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateAutomatedDiscoveryAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAllowList", func(t *testing.T) {
        input := &macie2.CreateAllowListInput{}
        output := &macie2.CreateAllowListOutput{}

        mockClient.On("CreateAllowList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAllowList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClassificationJob", func(t *testing.T) {
        input := &macie2.CreateClassificationJobInput{}
        output := &macie2.CreateClassificationJobOutput{}

        mockClient.On("CreateClassificationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClassificationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomDataIdentifier", func(t *testing.T) {
        input := &macie2.CreateCustomDataIdentifierInput{}
        output := &macie2.CreateCustomDataIdentifierOutput{}

        mockClient.On("CreateCustomDataIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomDataIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFindingsFilter", func(t *testing.T) {
        input := &macie2.CreateFindingsFilterInput{}
        output := &macie2.CreateFindingsFilterOutput{}

        mockClient.On("CreateFindingsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFindingsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInvitations", func(t *testing.T) {
        input := &macie2.CreateInvitationsInput{}
        output := &macie2.CreateInvitationsOutput{}

        mockClient.On("CreateInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMember", func(t *testing.T) {
        input := &macie2.CreateMemberInput{}
        output := &macie2.CreateMemberOutput{}

        mockClient.On("CreateMember", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSampleFindings", func(t *testing.T) {
        input := &macie2.CreateSampleFindingsInput{}
        output := &macie2.CreateSampleFindingsOutput{}

        mockClient.On("CreateSampleFindings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSampleFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeclineInvitations", func(t *testing.T) {
        input := &macie2.DeclineInvitationsInput{}
        output := &macie2.DeclineInvitationsOutput{}

        mockClient.On("DeclineInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeclineInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAllowList", func(t *testing.T) {
        input := &macie2.DeleteAllowListInput{}
        output := &macie2.DeleteAllowListOutput{}

        mockClient.On("DeleteAllowList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAllowList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomDataIdentifier", func(t *testing.T) {
        input := &macie2.DeleteCustomDataIdentifierInput{}
        output := &macie2.DeleteCustomDataIdentifierOutput{}

        mockClient.On("DeleteCustomDataIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomDataIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFindingsFilter", func(t *testing.T) {
        input := &macie2.DeleteFindingsFilterInput{}
        output := &macie2.DeleteFindingsFilterOutput{}

        mockClient.On("DeleteFindingsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFindingsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInvitations", func(t *testing.T) {
        input := &macie2.DeleteInvitationsInput{}
        output := &macie2.DeleteInvitationsOutput{}

        mockClient.On("DeleteInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMember", func(t *testing.T) {
        input := &macie2.DeleteMemberInput{}
        output := &macie2.DeleteMemberOutput{}

        mockClient.On("DeleteMember", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBuckets", func(t *testing.T) {
        input := &macie2.DescribeBucketsInput{}
        output := &macie2.DescribeBucketsOutput{}

        mockClient.On("DescribeBuckets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBuckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClassificationJob", func(t *testing.T) {
        input := &macie2.DescribeClassificationJobInput{}
        output := &macie2.DescribeClassificationJobOutput{}

        mockClient.On("DescribeClassificationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClassificationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfiguration", func(t *testing.T) {
        input := &macie2.DescribeOrganizationConfigurationInput{}
        output := &macie2.DescribeOrganizationConfigurationOutput{}

        mockClient.On("DescribeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableMacie", func(t *testing.T) {
        input := &macie2.DisableMacieInput{}
        output := &macie2.DisableMacieOutput{}

        mockClient.On("DisableMacie", ctx, input).Return(output, nil)

        result, err := mockClient.DisableMacie(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableOrganizationAdminAccount", func(t *testing.T) {
        input := &macie2.DisableOrganizationAdminAccountInput{}
        output := &macie2.DisableOrganizationAdminAccountOutput{}

        mockClient.On("DisableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromAdministratorAccount", func(t *testing.T) {
        input := &macie2.DisassociateFromAdministratorAccountInput{}
        output := &macie2.DisassociateFromAdministratorAccountOutput{}

        mockClient.On("DisassociateFromAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromMasterAccount", func(t *testing.T) {
        input := &macie2.DisassociateFromMasterAccountInput{}
        output := &macie2.DisassociateFromMasterAccountOutput{}

        mockClient.On("DisassociateFromMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMember", func(t *testing.T) {
        input := &macie2.DisassociateMemberInput{}
        output := &macie2.DisassociateMemberOutput{}

        mockClient.On("DisassociateMember", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableMacie", func(t *testing.T) {
        input := &macie2.EnableMacieInput{}
        output := &macie2.EnableMacieOutput{}

        mockClient.On("EnableMacie", ctx, input).Return(output, nil)

        result, err := mockClient.EnableMacie(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableOrganizationAdminAccount", func(t *testing.T) {
        input := &macie2.EnableOrganizationAdminAccountInput{}
        output := &macie2.EnableOrganizationAdminAccountOutput{}

        mockClient.On("EnableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdministratorAccount", func(t *testing.T) {
        input := &macie2.GetAdministratorAccountInput{}
        output := &macie2.GetAdministratorAccountOutput{}

        mockClient.On("GetAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAllowList", func(t *testing.T) {
        input := &macie2.GetAllowListInput{}
        output := &macie2.GetAllowListOutput{}

        mockClient.On("GetAllowList", ctx, input).Return(output, nil)

        result, err := mockClient.GetAllowList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutomatedDiscoveryConfiguration", func(t *testing.T) {
        input := &macie2.GetAutomatedDiscoveryConfigurationInput{}
        output := &macie2.GetAutomatedDiscoveryConfigurationOutput{}

        mockClient.On("GetAutomatedDiscoveryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutomatedDiscoveryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketStatistics", func(t *testing.T) {
        input := &macie2.GetBucketStatisticsInput{}
        output := &macie2.GetBucketStatisticsOutput{}

        mockClient.On("GetBucketStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClassificationExportConfiguration", func(t *testing.T) {
        input := &macie2.GetClassificationExportConfigurationInput{}
        output := &macie2.GetClassificationExportConfigurationOutput{}

        mockClient.On("GetClassificationExportConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetClassificationExportConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClassificationScope", func(t *testing.T) {
        input := &macie2.GetClassificationScopeInput{}
        output := &macie2.GetClassificationScopeOutput{}

        mockClient.On("GetClassificationScope", ctx, input).Return(output, nil)

        result, err := mockClient.GetClassificationScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomDataIdentifier", func(t *testing.T) {
        input := &macie2.GetCustomDataIdentifierInput{}
        output := &macie2.GetCustomDataIdentifierOutput{}

        mockClient.On("GetCustomDataIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomDataIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingStatistics", func(t *testing.T) {
        input := &macie2.GetFindingStatisticsInput{}
        output := &macie2.GetFindingStatisticsOutput{}

        mockClient.On("GetFindingStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindings", func(t *testing.T) {
        input := &macie2.GetFindingsInput{}
        output := &macie2.GetFindingsOutput{}

        mockClient.On("GetFindings", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingsFilter", func(t *testing.T) {
        input := &macie2.GetFindingsFilterInput{}
        output := &macie2.GetFindingsFilterOutput{}

        mockClient.On("GetFindingsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingsPublicationConfiguration", func(t *testing.T) {
        input := &macie2.GetFindingsPublicationConfigurationInput{}
        output := &macie2.GetFindingsPublicationConfigurationOutput{}

        mockClient.On("GetFindingsPublicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingsPublicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInvitationsCount", func(t *testing.T) {
        input := &macie2.GetInvitationsCountInput{}
        output := &macie2.GetInvitationsCountOutput{}

        mockClient.On("GetInvitationsCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetInvitationsCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMacieSession", func(t *testing.T) {
        input := &macie2.GetMacieSessionInput{}
        output := &macie2.GetMacieSessionOutput{}

        mockClient.On("GetMacieSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetMacieSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMasterAccount", func(t *testing.T) {
        input := &macie2.GetMasterAccountInput{}
        output := &macie2.GetMasterAccountOutput{}

        mockClient.On("GetMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMember", func(t *testing.T) {
        input := &macie2.GetMemberInput{}
        output := &macie2.GetMemberOutput{}

        mockClient.On("GetMember", ctx, input).Return(output, nil)

        result, err := mockClient.GetMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceProfile", func(t *testing.T) {
        input := &macie2.GetResourceProfileInput{}
        output := &macie2.GetResourceProfileOutput{}

        mockClient.On("GetResourceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRevealConfiguration", func(t *testing.T) {
        input := &macie2.GetRevealConfigurationInput{}
        output := &macie2.GetRevealConfigurationOutput{}

        mockClient.On("GetRevealConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetRevealConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSensitiveDataOccurrences", func(t *testing.T) {
        input := &macie2.GetSensitiveDataOccurrencesInput{}
        output := &macie2.GetSensitiveDataOccurrencesOutput{}

        mockClient.On("GetSensitiveDataOccurrences", ctx, input).Return(output, nil)

        result, err := mockClient.GetSensitiveDataOccurrences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSensitiveDataOccurrencesAvailability", func(t *testing.T) {
        input := &macie2.GetSensitiveDataOccurrencesAvailabilityInput{}
        output := &macie2.GetSensitiveDataOccurrencesAvailabilityOutput{}

        mockClient.On("GetSensitiveDataOccurrencesAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.GetSensitiveDataOccurrencesAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSensitivityInspectionTemplate", func(t *testing.T) {
        input := &macie2.GetSensitivityInspectionTemplateInput{}
        output := &macie2.GetSensitivityInspectionTemplateOutput{}

        mockClient.On("GetSensitivityInspectionTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetSensitivityInspectionTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageStatistics", func(t *testing.T) {
        input := &macie2.GetUsageStatisticsInput{}
        output := &macie2.GetUsageStatisticsOutput{}

        mockClient.On("GetUsageStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageTotals", func(t *testing.T) {
        input := &macie2.GetUsageTotalsInput{}
        output := &macie2.GetUsageTotalsOutput{}

        mockClient.On("GetUsageTotals", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageTotals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAllowLists", func(t *testing.T) {
        input := &macie2.ListAllowListsInput{}
        output := &macie2.ListAllowListsOutput{}

        mockClient.On("ListAllowLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListAllowLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutomatedDiscoveryAccounts", func(t *testing.T) {
        input := &macie2.ListAutomatedDiscoveryAccountsInput{}
        output := &macie2.ListAutomatedDiscoveryAccountsOutput{}

        mockClient.On("ListAutomatedDiscoveryAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutomatedDiscoveryAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClassificationJobs", func(t *testing.T) {
        input := &macie2.ListClassificationJobsInput{}
        output := &macie2.ListClassificationJobsOutput{}

        mockClient.On("ListClassificationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListClassificationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClassificationScopes", func(t *testing.T) {
        input := &macie2.ListClassificationScopesInput{}
        output := &macie2.ListClassificationScopesOutput{}

        mockClient.On("ListClassificationScopes", ctx, input).Return(output, nil)

        result, err := mockClient.ListClassificationScopes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomDataIdentifiers", func(t *testing.T) {
        input := &macie2.ListCustomDataIdentifiersInput{}
        output := &macie2.ListCustomDataIdentifiersOutput{}

        mockClient.On("ListCustomDataIdentifiers", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomDataIdentifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindings", func(t *testing.T) {
        input := &macie2.ListFindingsInput{}
        output := &macie2.ListFindingsOutput{}

        mockClient.On("ListFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingsFilters", func(t *testing.T) {
        input := &macie2.ListFindingsFiltersInput{}
        output := &macie2.ListFindingsFiltersOutput{}

        mockClient.On("ListFindingsFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingsFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvitations", func(t *testing.T) {
        input := &macie2.ListInvitationsInput{}
        output := &macie2.ListInvitationsOutput{}

        mockClient.On("ListInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedDataIdentifiers", func(t *testing.T) {
        input := &macie2.ListManagedDataIdentifiersInput{}
        output := &macie2.ListManagedDataIdentifiersOutput{}

        mockClient.On("ListManagedDataIdentifiers", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedDataIdentifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &macie2.ListMembersInput{}
        output := &macie2.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationAdminAccounts", func(t *testing.T) {
        input := &macie2.ListOrganizationAdminAccountsInput{}
        output := &macie2.ListOrganizationAdminAccountsOutput{}

        mockClient.On("ListOrganizationAdminAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationAdminAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceProfileArtifacts", func(t *testing.T) {
        input := &macie2.ListResourceProfileArtifactsInput{}
        output := &macie2.ListResourceProfileArtifactsOutput{}

        mockClient.On("ListResourceProfileArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceProfileArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceProfileDetections", func(t *testing.T) {
        input := &macie2.ListResourceProfileDetectionsInput{}
        output := &macie2.ListResourceProfileDetectionsOutput{}

        mockClient.On("ListResourceProfileDetections", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceProfileDetections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSensitivityInspectionTemplates", func(t *testing.T) {
        input := &macie2.ListSensitivityInspectionTemplatesInput{}
        output := &macie2.ListSensitivityInspectionTemplatesOutput{}

        mockClient.On("ListSensitivityInspectionTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListSensitivityInspectionTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &macie2.ListTagsForResourceInput{}
        output := &macie2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutClassificationExportConfiguration", func(t *testing.T) {
        input := &macie2.PutClassificationExportConfigurationInput{}
        output := &macie2.PutClassificationExportConfigurationOutput{}

        mockClient.On("PutClassificationExportConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutClassificationExportConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFindingsPublicationConfiguration", func(t *testing.T) {
        input := &macie2.PutFindingsPublicationConfigurationInput{}
        output := &macie2.PutFindingsPublicationConfigurationOutput{}

        mockClient.On("PutFindingsPublicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutFindingsPublicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchResources", func(t *testing.T) {
        input := &macie2.SearchResourcesInput{}
        output := &macie2.SearchResourcesOutput{}

        mockClient.On("SearchResources", ctx, input).Return(output, nil)

        result, err := mockClient.SearchResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &macie2.TagResourceInput{}
        output := &macie2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestCustomDataIdentifier", func(t *testing.T) {
        input := &macie2.TestCustomDataIdentifierInput{}
        output := &macie2.TestCustomDataIdentifierOutput{}

        mockClient.On("TestCustomDataIdentifier", ctx, input).Return(output, nil)

        result, err := mockClient.TestCustomDataIdentifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &macie2.UntagResourceInput{}
        output := &macie2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAllowList", func(t *testing.T) {
        input := &macie2.UpdateAllowListInput{}
        output := &macie2.UpdateAllowListOutput{}

        mockClient.On("UpdateAllowList", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAllowList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAutomatedDiscoveryConfiguration", func(t *testing.T) {
        input := &macie2.UpdateAutomatedDiscoveryConfigurationInput{}
        output := &macie2.UpdateAutomatedDiscoveryConfigurationOutput{}

        mockClient.On("UpdateAutomatedDiscoveryConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAutomatedDiscoveryConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClassificationJob", func(t *testing.T) {
        input := &macie2.UpdateClassificationJobInput{}
        output := &macie2.UpdateClassificationJobOutput{}

        mockClient.On("UpdateClassificationJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClassificationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClassificationScope", func(t *testing.T) {
        input := &macie2.UpdateClassificationScopeInput{}
        output := &macie2.UpdateClassificationScopeOutput{}

        mockClient.On("UpdateClassificationScope", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClassificationScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFindingsFilter", func(t *testing.T) {
        input := &macie2.UpdateFindingsFilterInput{}
        output := &macie2.UpdateFindingsFilterOutput{}

        mockClient.On("UpdateFindingsFilter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFindingsFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMacieSession", func(t *testing.T) {
        input := &macie2.UpdateMacieSessionInput{}
        output := &macie2.UpdateMacieSessionOutput{}

        mockClient.On("UpdateMacieSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMacieSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMemberSession", func(t *testing.T) {
        input := &macie2.UpdateMemberSessionInput{}
        output := &macie2.UpdateMemberSessionOutput{}

        mockClient.On("UpdateMemberSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMemberSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationConfiguration", func(t *testing.T) {
        input := &macie2.UpdateOrganizationConfigurationInput{}
        output := &macie2.UpdateOrganizationConfigurationOutput{}

        mockClient.On("UpdateOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceProfile", func(t *testing.T) {
        input := &macie2.UpdateResourceProfileInput{}
        output := &macie2.UpdateResourceProfileOutput{}

        mockClient.On("UpdateResourceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceProfileDetections", func(t *testing.T) {
        input := &macie2.UpdateResourceProfileDetectionsInput{}
        output := &macie2.UpdateResourceProfileDetectionsOutput{}

        mockClient.On("UpdateResourceProfileDetections", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceProfileDetections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRevealConfiguration", func(t *testing.T) {
        input := &macie2.UpdateRevealConfigurationInput{}
        output := &macie2.UpdateRevealConfigurationOutput{}

        mockClient.On("UpdateRevealConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRevealConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSensitivityInspectionTemplate", func(t *testing.T) {
        input := &macie2.UpdateSensitivityInspectionTemplateInput{}
        output := &macie2.UpdateSensitivityInspectionTemplateOutput{}

        mockClient.On("UpdateSensitivityInspectionTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSensitivityInspectionTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
