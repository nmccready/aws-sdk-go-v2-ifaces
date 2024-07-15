package guardduty_test

// tests for the guardduty service interface for this ../../../service/guardduty/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/guardduty/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/guardduty/guardduty_iface"
	"github.com/stretchr/testify/assert"
)

func TestGuarddutyServiceCanBeMocked(t *testing.T) {
	var iface guardduty_iface.IClient
	iface = &guardduty.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := guardduty.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptAdministratorInvitation", func(t *testing.T) {
        input := &guardduty.AcceptAdministratorInvitationInput{}
        output := &guardduty.AcceptAdministratorInvitationOutput{}

        mockClient.On("AcceptAdministratorInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptAdministratorInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInvitation", func(t *testing.T) {
        input := &guardduty.AcceptInvitationInput{}
        output := &guardduty.AcceptInvitationOutput{}

        mockClient.On("AcceptInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestArchiveFindings", func(t *testing.T) {
        input := &guardduty.ArchiveFindingsInput{}
        output := &guardduty.ArchiveFindingsOutput{}

        mockClient.On("ArchiveFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ArchiveFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDetector", func(t *testing.T) {
        input := &guardduty.CreateDetectorInput{}
        output := &guardduty.CreateDetectorOutput{}

        mockClient.On("CreateDetector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFilter", func(t *testing.T) {
        input := &guardduty.CreateFilterInput{}
        output := &guardduty.CreateFilterOutput{}

        mockClient.On("CreateFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIPSet", func(t *testing.T) {
        input := &guardduty.CreateIPSetInput{}
        output := &guardduty.CreateIPSetOutput{}

        mockClient.On("CreateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMalwareProtectionPlan", func(t *testing.T) {
        input := &guardduty.CreateMalwareProtectionPlanInput{}
        output := &guardduty.CreateMalwareProtectionPlanOutput{}

        mockClient.On("CreateMalwareProtectionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMalwareProtectionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMembers", func(t *testing.T) {
        input := &guardduty.CreateMembersInput{}
        output := &guardduty.CreateMembersOutput{}

        mockClient.On("CreateMembers", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePublishingDestination", func(t *testing.T) {
        input := &guardduty.CreatePublishingDestinationInput{}
        output := &guardduty.CreatePublishingDestinationOutput{}

        mockClient.On("CreatePublishingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePublishingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSampleFindings", func(t *testing.T) {
        input := &guardduty.CreateSampleFindingsInput{}
        output := &guardduty.CreateSampleFindingsOutput{}

        mockClient.On("CreateSampleFindings", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSampleFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThreatIntelSet", func(t *testing.T) {
        input := &guardduty.CreateThreatIntelSetInput{}
        output := &guardduty.CreateThreatIntelSetOutput{}

        mockClient.On("CreateThreatIntelSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThreatIntelSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeclineInvitations", func(t *testing.T) {
        input := &guardduty.DeclineInvitationsInput{}
        output := &guardduty.DeclineInvitationsOutput{}

        mockClient.On("DeclineInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeclineInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDetector", func(t *testing.T) {
        input := &guardduty.DeleteDetectorInput{}
        output := &guardduty.DeleteDetectorOutput{}

        mockClient.On("DeleteDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFilter", func(t *testing.T) {
        input := &guardduty.DeleteFilterInput{}
        output := &guardduty.DeleteFilterOutput{}

        mockClient.On("DeleteFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIPSet", func(t *testing.T) {
        input := &guardduty.DeleteIPSetInput{}
        output := &guardduty.DeleteIPSetOutput{}

        mockClient.On("DeleteIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInvitations", func(t *testing.T) {
        input := &guardduty.DeleteInvitationsInput{}
        output := &guardduty.DeleteInvitationsOutput{}

        mockClient.On("DeleteInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMalwareProtectionPlan", func(t *testing.T) {
        input := &guardduty.DeleteMalwareProtectionPlanInput{}
        output := &guardduty.DeleteMalwareProtectionPlanOutput{}

        mockClient.On("DeleteMalwareProtectionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMalwareProtectionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMembers", func(t *testing.T) {
        input := &guardduty.DeleteMembersInput{}
        output := &guardduty.DeleteMembersOutput{}

        mockClient.On("DeleteMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublishingDestination", func(t *testing.T) {
        input := &guardduty.DeletePublishingDestinationInput{}
        output := &guardduty.DeletePublishingDestinationOutput{}

        mockClient.On("DeletePublishingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublishingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThreatIntelSet", func(t *testing.T) {
        input := &guardduty.DeleteThreatIntelSetInput{}
        output := &guardduty.DeleteThreatIntelSetOutput{}

        mockClient.On("DeleteThreatIntelSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThreatIntelSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMalwareScans", func(t *testing.T) {
        input := &guardduty.DescribeMalwareScansInput{}
        output := &guardduty.DescribeMalwareScansOutput{}

        mockClient.On("DescribeMalwareScans", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMalwareScans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfiguration", func(t *testing.T) {
        input := &guardduty.DescribeOrganizationConfigurationInput{}
        output := &guardduty.DescribeOrganizationConfigurationOutput{}

        mockClient.On("DescribeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePublishingDestination", func(t *testing.T) {
        input := &guardduty.DescribePublishingDestinationInput{}
        output := &guardduty.DescribePublishingDestinationOutput{}

        mockClient.On("DescribePublishingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePublishingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableOrganizationAdminAccount", func(t *testing.T) {
        input := &guardduty.DisableOrganizationAdminAccountInput{}
        output := &guardduty.DisableOrganizationAdminAccountOutput{}

        mockClient.On("DisableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromAdministratorAccount", func(t *testing.T) {
        input := &guardduty.DisassociateFromAdministratorAccountInput{}
        output := &guardduty.DisassociateFromAdministratorAccountOutput{}

        mockClient.On("DisassociateFromAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromMasterAccount", func(t *testing.T) {
        input := &guardduty.DisassociateFromMasterAccountInput{}
        output := &guardduty.DisassociateFromMasterAccountOutput{}

        mockClient.On("DisassociateFromMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMembers", func(t *testing.T) {
        input := &guardduty.DisassociateMembersInput{}
        output := &guardduty.DisassociateMembersOutput{}

        mockClient.On("DisassociateMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableOrganizationAdminAccount", func(t *testing.T) {
        input := &guardduty.EnableOrganizationAdminAccountInput{}
        output := &guardduty.EnableOrganizationAdminAccountOutput{}

        mockClient.On("EnableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdministratorAccount", func(t *testing.T) {
        input := &guardduty.GetAdministratorAccountInput{}
        output := &guardduty.GetAdministratorAccountOutput{}

        mockClient.On("GetAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoverageStatistics", func(t *testing.T) {
        input := &guardduty.GetCoverageStatisticsInput{}
        output := &guardduty.GetCoverageStatisticsOutput{}

        mockClient.On("GetCoverageStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoverageStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDetector", func(t *testing.T) {
        input := &guardduty.GetDetectorInput{}
        output := &guardduty.GetDetectorOutput{}

        mockClient.On("GetDetector", ctx, input).Return(output, nil)

        result, err := mockClient.GetDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFilter", func(t *testing.T) {
        input := &guardduty.GetFilterInput{}
        output := &guardduty.GetFilterOutput{}

        mockClient.On("GetFilter", ctx, input).Return(output, nil)

        result, err := mockClient.GetFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindings", func(t *testing.T) {
        input := &guardduty.GetFindingsInput{}
        output := &guardduty.GetFindingsOutput{}

        mockClient.On("GetFindings", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingsStatistics", func(t *testing.T) {
        input := &guardduty.GetFindingsStatisticsInput{}
        output := &guardduty.GetFindingsStatisticsOutput{}

        mockClient.On("GetFindingsStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingsStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIPSet", func(t *testing.T) {
        input := &guardduty.GetIPSetInput{}
        output := &guardduty.GetIPSetOutput{}

        mockClient.On("GetIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInvitationsCount", func(t *testing.T) {
        input := &guardduty.GetInvitationsCountInput{}
        output := &guardduty.GetInvitationsCountOutput{}

        mockClient.On("GetInvitationsCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetInvitationsCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMalwareProtectionPlan", func(t *testing.T) {
        input := &guardduty.GetMalwareProtectionPlanInput{}
        output := &guardduty.GetMalwareProtectionPlanOutput{}

        mockClient.On("GetMalwareProtectionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetMalwareProtectionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMalwareScanSettings", func(t *testing.T) {
        input := &guardduty.GetMalwareScanSettingsInput{}
        output := &guardduty.GetMalwareScanSettingsOutput{}

        mockClient.On("GetMalwareScanSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetMalwareScanSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMasterAccount", func(t *testing.T) {
        input := &guardduty.GetMasterAccountInput{}
        output := &guardduty.GetMasterAccountOutput{}

        mockClient.On("GetMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMemberDetectors", func(t *testing.T) {
        input := &guardduty.GetMemberDetectorsInput{}
        output := &guardduty.GetMemberDetectorsOutput{}

        mockClient.On("GetMemberDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.GetMemberDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMembers", func(t *testing.T) {
        input := &guardduty.GetMembersInput{}
        output := &guardduty.GetMembersOutput{}

        mockClient.On("GetMembers", ctx, input).Return(output, nil)

        result, err := mockClient.GetMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationStatistics", func(t *testing.T) {
        input := &guardduty.GetOrganizationStatisticsInput{}
        output := &guardduty.GetOrganizationStatisticsOutput{}

        mockClient.On("GetOrganizationStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRemainingFreeTrialDays", func(t *testing.T) {
        input := &guardduty.GetRemainingFreeTrialDaysInput{}
        output := &guardduty.GetRemainingFreeTrialDaysOutput{}

        mockClient.On("GetRemainingFreeTrialDays", ctx, input).Return(output, nil)

        result, err := mockClient.GetRemainingFreeTrialDays(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetThreatIntelSet", func(t *testing.T) {
        input := &guardduty.GetThreatIntelSetInput{}
        output := &guardduty.GetThreatIntelSetOutput{}

        mockClient.On("GetThreatIntelSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetThreatIntelSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageStatistics", func(t *testing.T) {
        input := &guardduty.GetUsageStatisticsInput{}
        output := &guardduty.GetUsageStatisticsOutput{}

        mockClient.On("GetUsageStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInviteMembers", func(t *testing.T) {
        input := &guardduty.InviteMembersInput{}
        output := &guardduty.InviteMembersOutput{}

        mockClient.On("InviteMembers", ctx, input).Return(output, nil)

        result, err := mockClient.InviteMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoverage", func(t *testing.T) {
        input := &guardduty.ListCoverageInput{}
        output := &guardduty.ListCoverageOutput{}

        mockClient.On("ListCoverage", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoverage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectors", func(t *testing.T) {
        input := &guardduty.ListDetectorsInput{}
        output := &guardduty.ListDetectorsOutput{}

        mockClient.On("ListDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFilters", func(t *testing.T) {
        input := &guardduty.ListFiltersInput{}
        output := &guardduty.ListFiltersOutput{}

        mockClient.On("ListFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindings", func(t *testing.T) {
        input := &guardduty.ListFindingsInput{}
        output := &guardduty.ListFindingsOutput{}

        mockClient.On("ListFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIPSets", func(t *testing.T) {
        input := &guardduty.ListIPSetsInput{}
        output := &guardduty.ListIPSetsOutput{}

        mockClient.On("ListIPSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListIPSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvitations", func(t *testing.T) {
        input := &guardduty.ListInvitationsInput{}
        output := &guardduty.ListInvitationsOutput{}

        mockClient.On("ListInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMalwareProtectionPlans", func(t *testing.T) {
        input := &guardduty.ListMalwareProtectionPlansInput{}
        output := &guardduty.ListMalwareProtectionPlansOutput{}

        mockClient.On("ListMalwareProtectionPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListMalwareProtectionPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &guardduty.ListMembersInput{}
        output := &guardduty.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationAdminAccounts", func(t *testing.T) {
        input := &guardduty.ListOrganizationAdminAccountsInput{}
        output := &guardduty.ListOrganizationAdminAccountsOutput{}

        mockClient.On("ListOrganizationAdminAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationAdminAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPublishingDestinations", func(t *testing.T) {
        input := &guardduty.ListPublishingDestinationsInput{}
        output := &guardduty.ListPublishingDestinationsOutput{}

        mockClient.On("ListPublishingDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPublishingDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &guardduty.ListTagsForResourceInput{}
        output := &guardduty.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThreatIntelSets", func(t *testing.T) {
        input := &guardduty.ListThreatIntelSetsInput{}
        output := &guardduty.ListThreatIntelSetsOutput{}

        mockClient.On("ListThreatIntelSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListThreatIntelSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMalwareScan", func(t *testing.T) {
        input := &guardduty.StartMalwareScanInput{}
        output := &guardduty.StartMalwareScanOutput{}

        mockClient.On("StartMalwareScan", ctx, input).Return(output, nil)

        result, err := mockClient.StartMalwareScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMonitoringMembers", func(t *testing.T) {
        input := &guardduty.StartMonitoringMembersInput{}
        output := &guardduty.StartMonitoringMembersOutput{}

        mockClient.On("StartMonitoringMembers", ctx, input).Return(output, nil)

        result, err := mockClient.StartMonitoringMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMonitoringMembers", func(t *testing.T) {
        input := &guardduty.StopMonitoringMembersInput{}
        output := &guardduty.StopMonitoringMembersOutput{}

        mockClient.On("StopMonitoringMembers", ctx, input).Return(output, nil)

        result, err := mockClient.StopMonitoringMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &guardduty.TagResourceInput{}
        output := &guardduty.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnarchiveFindings", func(t *testing.T) {
        input := &guardduty.UnarchiveFindingsInput{}
        output := &guardduty.UnarchiveFindingsOutput{}

        mockClient.On("UnarchiveFindings", ctx, input).Return(output, nil)

        result, err := mockClient.UnarchiveFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &guardduty.UntagResourceInput{}
        output := &guardduty.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDetector", func(t *testing.T) {
        input := &guardduty.UpdateDetectorInput{}
        output := &guardduty.UpdateDetectorOutput{}

        mockClient.On("UpdateDetector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFilter", func(t *testing.T) {
        input := &guardduty.UpdateFilterInput{}
        output := &guardduty.UpdateFilterOutput{}

        mockClient.On("UpdateFilter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFindingsFeedback", func(t *testing.T) {
        input := &guardduty.UpdateFindingsFeedbackInput{}
        output := &guardduty.UpdateFindingsFeedbackOutput{}

        mockClient.On("UpdateFindingsFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFindingsFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIPSet", func(t *testing.T) {
        input := &guardduty.UpdateIPSetInput{}
        output := &guardduty.UpdateIPSetOutput{}

        mockClient.On("UpdateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMalwareProtectionPlan", func(t *testing.T) {
        input := &guardduty.UpdateMalwareProtectionPlanInput{}
        output := &guardduty.UpdateMalwareProtectionPlanOutput{}

        mockClient.On("UpdateMalwareProtectionPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMalwareProtectionPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMalwareScanSettings", func(t *testing.T) {
        input := &guardduty.UpdateMalwareScanSettingsInput{}
        output := &guardduty.UpdateMalwareScanSettingsOutput{}

        mockClient.On("UpdateMalwareScanSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMalwareScanSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMemberDetectors", func(t *testing.T) {
        input := &guardduty.UpdateMemberDetectorsInput{}
        output := &guardduty.UpdateMemberDetectorsOutput{}

        mockClient.On("UpdateMemberDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMemberDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationConfiguration", func(t *testing.T) {
        input := &guardduty.UpdateOrganizationConfigurationInput{}
        output := &guardduty.UpdateOrganizationConfigurationOutput{}

        mockClient.On("UpdateOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePublishingDestination", func(t *testing.T) {
        input := &guardduty.UpdatePublishingDestinationInput{}
        output := &guardduty.UpdatePublishingDestinationOutput{}

        mockClient.On("UpdatePublishingDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePublishingDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThreatIntelSet", func(t *testing.T) {
        input := &guardduty.UpdateThreatIntelSetInput{}
        output := &guardduty.UpdateThreatIntelSetOutput{}

        mockClient.On("UpdateThreatIntelSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThreatIntelSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
