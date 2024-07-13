package inspector2_test

// tests for the inspector2 service interface for this ../../../service/inspector2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspector2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspector2/inspector2_iface"
	"github.com/stretchr/testify/assert"
)

func TestInspector2ServiceCanBeMocked(t *testing.T) {
	var iface inspector2_iface.IClient
	iface = &inspector2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := inspector2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMember", func(t *testing.T) {
        input := &inspector2.AssociateMemberInput{}
        output := &inspector2.AssociateMemberOutput{}

        mockClient.On("AssociateMember", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAccountStatus", func(t *testing.T) {
        input := &inspector2.BatchGetAccountStatusInput{}
        output := &inspector2.BatchGetAccountStatusOutput{}

        mockClient.On("BatchGetAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCodeSnippet", func(t *testing.T) {
        input := &inspector2.BatchGetCodeSnippetInput{}
        output := &inspector2.BatchGetCodeSnippetOutput{}

        mockClient.On("BatchGetCodeSnippet", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCodeSnippet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFindingDetails", func(t *testing.T) {
        input := &inspector2.BatchGetFindingDetailsInput{}
        output := &inspector2.BatchGetFindingDetailsOutput{}

        mockClient.On("BatchGetFindingDetails", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFindingDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFreeTrialInfo", func(t *testing.T) {
        input := &inspector2.BatchGetFreeTrialInfoInput{}
        output := &inspector2.BatchGetFreeTrialInfoOutput{}

        mockClient.On("BatchGetFreeTrialInfo", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFreeTrialInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetMemberEc2DeepInspectionStatus", func(t *testing.T) {
        input := &inspector2.BatchGetMemberEc2DeepInspectionStatusInput{}
        output := &inspector2.BatchGetMemberEc2DeepInspectionStatusOutput{}

        mockClient.On("BatchGetMemberEc2DeepInspectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetMemberEc2DeepInspectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateMemberEc2DeepInspectionStatus", func(t *testing.T) {
        input := &inspector2.BatchUpdateMemberEc2DeepInspectionStatusInput{}
        output := &inspector2.BatchUpdateMemberEc2DeepInspectionStatusOutput{}

        mockClient.On("BatchUpdateMemberEc2DeepInspectionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateMemberEc2DeepInspectionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelFindingsReport", func(t *testing.T) {
        input := &inspector2.CancelFindingsReportInput{}
        output := &inspector2.CancelFindingsReportOutput{}

        mockClient.On("CancelFindingsReport", ctx, input).Return(output, nil)

        result, err := mockClient.CancelFindingsReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSbomExport", func(t *testing.T) {
        input := &inspector2.CancelSbomExportInput{}
        output := &inspector2.CancelSbomExportOutput{}

        mockClient.On("CancelSbomExport", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSbomExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCisScanConfiguration", func(t *testing.T) {
        input := &inspector2.CreateCisScanConfigurationInput{}
        output := &inspector2.CreateCisScanConfigurationOutput{}

        mockClient.On("CreateCisScanConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCisScanConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFilter", func(t *testing.T) {
        input := &inspector2.CreateFilterInput{}
        output := &inspector2.CreateFilterOutput{}

        mockClient.On("CreateFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFindingsReport", func(t *testing.T) {
        input := &inspector2.CreateFindingsReportInput{}
        output := &inspector2.CreateFindingsReportOutput{}

        mockClient.On("CreateFindingsReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFindingsReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSbomExport", func(t *testing.T) {
        input := &inspector2.CreateSbomExportInput{}
        output := &inspector2.CreateSbomExportOutput{}

        mockClient.On("CreateSbomExport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSbomExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCisScanConfiguration", func(t *testing.T) {
        input := &inspector2.DeleteCisScanConfigurationInput{}
        output := &inspector2.DeleteCisScanConfigurationOutput{}

        mockClient.On("DeleteCisScanConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCisScanConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFilter", func(t *testing.T) {
        input := &inspector2.DeleteFilterInput{}
        output := &inspector2.DeleteFilterOutput{}

        mockClient.On("DeleteFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfiguration", func(t *testing.T) {
        input := &inspector2.DescribeOrganizationConfigurationInput{}
        output := &inspector2.DescribeOrganizationConfigurationOutput{}

        mockClient.On("DescribeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisable", func(t *testing.T) {
        input := &inspector2.DisableInput{}
        output := &inspector2.DisableOutput{}

        mockClient.On("Disable", ctx, input).Return(output, nil)

        result, err := mockClient.Disable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableDelegatedAdminAccount", func(t *testing.T) {
        input := &inspector2.DisableDelegatedAdminAccountInput{}
        output := &inspector2.DisableDelegatedAdminAccountOutput{}

        mockClient.On("DisableDelegatedAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableDelegatedAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMember", func(t *testing.T) {
        input := &inspector2.DisassociateMemberInput{}
        output := &inspector2.DisassociateMemberOutput{}

        mockClient.On("DisassociateMember", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnable", func(t *testing.T) {
        input := &inspector2.EnableInput{}
        output := &inspector2.EnableOutput{}

        mockClient.On("Enable", ctx, input).Return(output, nil)

        result, err := mockClient.Enable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableDelegatedAdminAccount", func(t *testing.T) {
        input := &inspector2.EnableDelegatedAdminAccountInput{}
        output := &inspector2.EnableDelegatedAdminAccountOutput{}

        mockClient.On("EnableDelegatedAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableDelegatedAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCisScanReport", func(t *testing.T) {
        input := &inspector2.GetCisScanReportInput{}
        output := &inspector2.GetCisScanReportOutput{}

        mockClient.On("GetCisScanReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetCisScanReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCisScanResultDetails", func(t *testing.T) {
        input := &inspector2.GetCisScanResultDetailsInput{}
        output := &inspector2.GetCisScanResultDetailsOutput{}

        mockClient.On("GetCisScanResultDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetCisScanResultDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfiguration", func(t *testing.T) {
        input := &inspector2.GetConfigurationInput{}
        output := &inspector2.GetConfigurationOutput{}

        mockClient.On("GetConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDelegatedAdminAccount", func(t *testing.T) {
        input := &inspector2.GetDelegatedAdminAccountInput{}
        output := &inspector2.GetDelegatedAdminAccountOutput{}

        mockClient.On("GetDelegatedAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetDelegatedAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEc2DeepInspectionConfiguration", func(t *testing.T) {
        input := &inspector2.GetEc2DeepInspectionConfigurationInput{}
        output := &inspector2.GetEc2DeepInspectionConfigurationOutput{}

        mockClient.On("GetEc2DeepInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEc2DeepInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEncryptionKey", func(t *testing.T) {
        input := &inspector2.GetEncryptionKeyInput{}
        output := &inspector2.GetEncryptionKeyOutput{}

        mockClient.On("GetEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingsReportStatus", func(t *testing.T) {
        input := &inspector2.GetFindingsReportStatusInput{}
        output := &inspector2.GetFindingsReportStatusOutput{}

        mockClient.On("GetFindingsReportStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingsReportStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMember", func(t *testing.T) {
        input := &inspector2.GetMemberInput{}
        output := &inspector2.GetMemberOutput{}

        mockClient.On("GetMember", ctx, input).Return(output, nil)

        result, err := mockClient.GetMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSbomExport", func(t *testing.T) {
        input := &inspector2.GetSbomExportInput{}
        output := &inspector2.GetSbomExportOutput{}

        mockClient.On("GetSbomExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetSbomExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountPermissions", func(t *testing.T) {
        input := &inspector2.ListAccountPermissionsInput{}
        output := &inspector2.ListAccountPermissionsOutput{}

        mockClient.On("ListAccountPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCisScanConfigurations", func(t *testing.T) {
        input := &inspector2.ListCisScanConfigurationsInput{}
        output := &inspector2.ListCisScanConfigurationsOutput{}

        mockClient.On("ListCisScanConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCisScanConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCisScanResultsAggregatedByChecks", func(t *testing.T) {
        input := &inspector2.ListCisScanResultsAggregatedByChecksInput{}
        output := &inspector2.ListCisScanResultsAggregatedByChecksOutput{}

        mockClient.On("ListCisScanResultsAggregatedByChecks", ctx, input).Return(output, nil)

        result, err := mockClient.ListCisScanResultsAggregatedByChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCisScanResultsAggregatedByTargetResource", func(t *testing.T) {
        input := &inspector2.ListCisScanResultsAggregatedByTargetResourceInput{}
        output := &inspector2.ListCisScanResultsAggregatedByTargetResourceOutput{}

        mockClient.On("ListCisScanResultsAggregatedByTargetResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListCisScanResultsAggregatedByTargetResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCisScans", func(t *testing.T) {
        input := &inspector2.ListCisScansInput{}
        output := &inspector2.ListCisScansOutput{}

        mockClient.On("ListCisScans", ctx, input).Return(output, nil)

        result, err := mockClient.ListCisScans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoverage", func(t *testing.T) {
        input := &inspector2.ListCoverageInput{}
        output := &inspector2.ListCoverageOutput{}

        mockClient.On("ListCoverage", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoverage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoverageStatistics", func(t *testing.T) {
        input := &inspector2.ListCoverageStatisticsInput{}
        output := &inspector2.ListCoverageStatisticsOutput{}

        mockClient.On("ListCoverageStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoverageStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDelegatedAdminAccounts", func(t *testing.T) {
        input := &inspector2.ListDelegatedAdminAccountsInput{}
        output := &inspector2.ListDelegatedAdminAccountsOutput{}

        mockClient.On("ListDelegatedAdminAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListDelegatedAdminAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFilters", func(t *testing.T) {
        input := &inspector2.ListFiltersInput{}
        output := &inspector2.ListFiltersOutput{}

        mockClient.On("ListFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingAggregations", func(t *testing.T) {
        input := &inspector2.ListFindingAggregationsInput{}
        output := &inspector2.ListFindingAggregationsOutput{}

        mockClient.On("ListFindingAggregations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingAggregations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindings", func(t *testing.T) {
        input := &inspector2.ListFindingsInput{}
        output := &inspector2.ListFindingsOutput{}

        mockClient.On("ListFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &inspector2.ListMembersInput{}
        output := &inspector2.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &inspector2.ListTagsForResourceInput{}
        output := &inspector2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsageTotals", func(t *testing.T) {
        input := &inspector2.ListUsageTotalsInput{}
        output := &inspector2.ListUsageTotalsOutput{}

        mockClient.On("ListUsageTotals", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsageTotals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetEncryptionKey", func(t *testing.T) {
        input := &inspector2.ResetEncryptionKeyInput{}
        output := &inspector2.ResetEncryptionKeyOutput{}

        mockClient.On("ResetEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.ResetEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchVulnerabilities", func(t *testing.T) {
        input := &inspector2.SearchVulnerabilitiesInput{}
        output := &inspector2.SearchVulnerabilitiesOutput{}

        mockClient.On("SearchVulnerabilities", ctx, input).Return(output, nil)

        result, err := mockClient.SearchVulnerabilities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCisSessionHealth", func(t *testing.T) {
        input := &inspector2.SendCisSessionHealthInput{}
        output := &inspector2.SendCisSessionHealthOutput{}

        mockClient.On("SendCisSessionHealth", ctx, input).Return(output, nil)

        result, err := mockClient.SendCisSessionHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendCisSessionTelemetry", func(t *testing.T) {
        input := &inspector2.SendCisSessionTelemetryInput{}
        output := &inspector2.SendCisSessionTelemetryOutput{}

        mockClient.On("SendCisSessionTelemetry", ctx, input).Return(output, nil)

        result, err := mockClient.SendCisSessionTelemetry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCisSession", func(t *testing.T) {
        input := &inspector2.StartCisSessionInput{}
        output := &inspector2.StartCisSessionOutput{}

        mockClient.On("StartCisSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartCisSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCisSession", func(t *testing.T) {
        input := &inspector2.StopCisSessionInput{}
        output := &inspector2.StopCisSessionOutput{}

        mockClient.On("StopCisSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopCisSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &inspector2.TagResourceInput{}
        output := &inspector2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &inspector2.UntagResourceInput{}
        output := &inspector2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCisScanConfiguration", func(t *testing.T) {
        input := &inspector2.UpdateCisScanConfigurationInput{}
        output := &inspector2.UpdateCisScanConfigurationOutput{}

        mockClient.On("UpdateCisScanConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCisScanConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguration", func(t *testing.T) {
        input := &inspector2.UpdateConfigurationInput{}
        output := &inspector2.UpdateConfigurationOutput{}

        mockClient.On("UpdateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEc2DeepInspectionConfiguration", func(t *testing.T) {
        input := &inspector2.UpdateEc2DeepInspectionConfigurationInput{}
        output := &inspector2.UpdateEc2DeepInspectionConfigurationOutput{}

        mockClient.On("UpdateEc2DeepInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEc2DeepInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEncryptionKey", func(t *testing.T) {
        input := &inspector2.UpdateEncryptionKeyInput{}
        output := &inspector2.UpdateEncryptionKeyOutput{}

        mockClient.On("UpdateEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFilter", func(t *testing.T) {
        input := &inspector2.UpdateFilterInput{}
        output := &inspector2.UpdateFilterOutput{}

        mockClient.On("UpdateFilter", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrgEc2DeepInspectionConfiguration", func(t *testing.T) {
        input := &inspector2.UpdateOrgEc2DeepInspectionConfigurationInput{}
        output := &inspector2.UpdateOrgEc2DeepInspectionConfigurationOutput{}

        mockClient.On("UpdateOrgEc2DeepInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrgEc2DeepInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationConfiguration", func(t *testing.T) {
        input := &inspector2.UpdateOrganizationConfigurationInput{}
        output := &inspector2.UpdateOrganizationConfigurationOutput{}

        mockClient.On("UpdateOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
