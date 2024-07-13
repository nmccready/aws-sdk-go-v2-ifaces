package auditmanager_test

// tests for the auditmanager service interface for this ../../../service/auditmanager/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/auditmanager/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/auditmanager/auditmanager_iface"
	"github.com/stretchr/testify/assert"
)

func TestAuditmanagerServiceCanBeMocked(t *testing.T) {
	var iface auditmanager_iface.IClient
	iface = &auditmanager.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := auditmanager.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAssessmentReportEvidenceFolder", func(t *testing.T) {
        input := &auditmanager.AssociateAssessmentReportEvidenceFolderInput{}
        output := &auditmanager.AssociateAssessmentReportEvidenceFolderOutput{}

        mockClient.On("AssociateAssessmentReportEvidenceFolder", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAssessmentReportEvidenceFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateAssessmentReportEvidence", func(t *testing.T) {
        input := &auditmanager.BatchAssociateAssessmentReportEvidenceInput{}
        output := &auditmanager.BatchAssociateAssessmentReportEvidenceOutput{}

        mockClient.On("BatchAssociateAssessmentReportEvidence", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateAssessmentReportEvidence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateDelegationByAssessment", func(t *testing.T) {
        input := &auditmanager.BatchCreateDelegationByAssessmentInput{}
        output := &auditmanager.BatchCreateDelegationByAssessmentOutput{}

        mockClient.On("BatchCreateDelegationByAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateDelegationByAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteDelegationByAssessment", func(t *testing.T) {
        input := &auditmanager.BatchDeleteDelegationByAssessmentInput{}
        output := &auditmanager.BatchDeleteDelegationByAssessmentOutput{}

        mockClient.On("BatchDeleteDelegationByAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteDelegationByAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateAssessmentReportEvidence", func(t *testing.T) {
        input := &auditmanager.BatchDisassociateAssessmentReportEvidenceInput{}
        output := &auditmanager.BatchDisassociateAssessmentReportEvidenceOutput{}

        mockClient.On("BatchDisassociateAssessmentReportEvidence", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateAssessmentReportEvidence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchImportEvidenceToAssessmentControl", func(t *testing.T) {
        input := &auditmanager.BatchImportEvidenceToAssessmentControlInput{}
        output := &auditmanager.BatchImportEvidenceToAssessmentControlOutput{}

        mockClient.On("BatchImportEvidenceToAssessmentControl", ctx, input).Return(output, nil)

        result, err := mockClient.BatchImportEvidenceToAssessmentControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssessment", func(t *testing.T) {
        input := &auditmanager.CreateAssessmentInput{}
        output := &auditmanager.CreateAssessmentOutput{}

        mockClient.On("CreateAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssessmentFramework", func(t *testing.T) {
        input := &auditmanager.CreateAssessmentFrameworkInput{}
        output := &auditmanager.CreateAssessmentFrameworkOutput{}

        mockClient.On("CreateAssessmentFramework", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssessmentFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssessmentReport", func(t *testing.T) {
        input := &auditmanager.CreateAssessmentReportInput{}
        output := &auditmanager.CreateAssessmentReportOutput{}

        mockClient.On("CreateAssessmentReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssessmentReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateControl", func(t *testing.T) {
        input := &auditmanager.CreateControlInput{}
        output := &auditmanager.CreateControlOutput{}

        mockClient.On("CreateControl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessment", func(t *testing.T) {
        input := &auditmanager.DeleteAssessmentInput{}
        output := &auditmanager.DeleteAssessmentOutput{}

        mockClient.On("DeleteAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentFramework", func(t *testing.T) {
        input := &auditmanager.DeleteAssessmentFrameworkInput{}
        output := &auditmanager.DeleteAssessmentFrameworkOutput{}

        mockClient.On("DeleteAssessmentFramework", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentFrameworkShare", func(t *testing.T) {
        input := &auditmanager.DeleteAssessmentFrameworkShareInput{}
        output := &auditmanager.DeleteAssessmentFrameworkShareOutput{}

        mockClient.On("DeleteAssessmentFrameworkShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentFrameworkShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentReport", func(t *testing.T) {
        input := &auditmanager.DeleteAssessmentReportInput{}
        output := &auditmanager.DeleteAssessmentReportOutput{}

        mockClient.On("DeleteAssessmentReport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteControl", func(t *testing.T) {
        input := &auditmanager.DeleteControlInput{}
        output := &auditmanager.DeleteControlOutput{}

        mockClient.On("DeleteControl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterAccount", func(t *testing.T) {
        input := &auditmanager.DeregisterAccountInput{}
        output := &auditmanager.DeregisterAccountOutput{}

        mockClient.On("DeregisterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterOrganizationAdminAccount", func(t *testing.T) {
        input := &auditmanager.DeregisterOrganizationAdminAccountInput{}
        output := &auditmanager.DeregisterOrganizationAdminAccountOutput{}

        mockClient.On("DeregisterOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAssessmentReportEvidenceFolder", func(t *testing.T) {
        input := &auditmanager.DisassociateAssessmentReportEvidenceFolderInput{}
        output := &auditmanager.DisassociateAssessmentReportEvidenceFolderOutput{}

        mockClient.On("DisassociateAssessmentReportEvidenceFolder", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAssessmentReportEvidenceFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountStatus", func(t *testing.T) {
        input := &auditmanager.GetAccountStatusInput{}
        output := &auditmanager.GetAccountStatusOutput{}

        mockClient.On("GetAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssessment", func(t *testing.T) {
        input := &auditmanager.GetAssessmentInput{}
        output := &auditmanager.GetAssessmentOutput{}

        mockClient.On("GetAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssessmentFramework", func(t *testing.T) {
        input := &auditmanager.GetAssessmentFrameworkInput{}
        output := &auditmanager.GetAssessmentFrameworkOutput{}

        mockClient.On("GetAssessmentFramework", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssessmentFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssessmentReportUrl", func(t *testing.T) {
        input := &auditmanager.GetAssessmentReportUrlInput{}
        output := &auditmanager.GetAssessmentReportUrlOutput{}

        mockClient.On("GetAssessmentReportUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssessmentReportUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeLogs", func(t *testing.T) {
        input := &auditmanager.GetChangeLogsInput{}
        output := &auditmanager.GetChangeLogsOutput{}

        mockClient.On("GetChangeLogs", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetControl", func(t *testing.T) {
        input := &auditmanager.GetControlInput{}
        output := &auditmanager.GetControlOutput{}

        mockClient.On("GetControl", ctx, input).Return(output, nil)

        result, err := mockClient.GetControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDelegations", func(t *testing.T) {
        input := &auditmanager.GetDelegationsInput{}
        output := &auditmanager.GetDelegationsOutput{}

        mockClient.On("GetDelegations", ctx, input).Return(output, nil)

        result, err := mockClient.GetDelegations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidence", func(t *testing.T) {
        input := &auditmanager.GetEvidenceInput{}
        output := &auditmanager.GetEvidenceOutput{}

        mockClient.On("GetEvidence", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidenceByEvidenceFolder", func(t *testing.T) {
        input := &auditmanager.GetEvidenceByEvidenceFolderInput{}
        output := &auditmanager.GetEvidenceByEvidenceFolderOutput{}

        mockClient.On("GetEvidenceByEvidenceFolder", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidenceByEvidenceFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidenceFileUploadUrl", func(t *testing.T) {
        input := &auditmanager.GetEvidenceFileUploadUrlInput{}
        output := &auditmanager.GetEvidenceFileUploadUrlOutput{}

        mockClient.On("GetEvidenceFileUploadUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidenceFileUploadUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidenceFolder", func(t *testing.T) {
        input := &auditmanager.GetEvidenceFolderInput{}
        output := &auditmanager.GetEvidenceFolderOutput{}

        mockClient.On("GetEvidenceFolder", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidenceFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidenceFoldersByAssessment", func(t *testing.T) {
        input := &auditmanager.GetEvidenceFoldersByAssessmentInput{}
        output := &auditmanager.GetEvidenceFoldersByAssessmentOutput{}

        mockClient.On("GetEvidenceFoldersByAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidenceFoldersByAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvidenceFoldersByAssessmentControl", func(t *testing.T) {
        input := &auditmanager.GetEvidenceFoldersByAssessmentControlInput{}
        output := &auditmanager.GetEvidenceFoldersByAssessmentControlOutput{}

        mockClient.On("GetEvidenceFoldersByAssessmentControl", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvidenceFoldersByAssessmentControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsights", func(t *testing.T) {
        input := &auditmanager.GetInsightsInput{}
        output := &auditmanager.GetInsightsOutput{}

        mockClient.On("GetInsights", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightsByAssessment", func(t *testing.T) {
        input := &auditmanager.GetInsightsByAssessmentInput{}
        output := &auditmanager.GetInsightsByAssessmentOutput{}

        mockClient.On("GetInsightsByAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightsByAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationAdminAccount", func(t *testing.T) {
        input := &auditmanager.GetOrganizationAdminAccountInput{}
        output := &auditmanager.GetOrganizationAdminAccountOutput{}

        mockClient.On("GetOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServicesInScope", func(t *testing.T) {
        input := &auditmanager.GetServicesInScopeInput{}
        output := &auditmanager.GetServicesInScopeOutput{}

        mockClient.On("GetServicesInScope", ctx, input).Return(output, nil)

        result, err := mockClient.GetServicesInScope(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSettings", func(t *testing.T) {
        input := &auditmanager.GetSettingsInput{}
        output := &auditmanager.GetSettingsOutput{}

        mockClient.On("GetSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentControlInsightsByControlDomain", func(t *testing.T) {
        input := &auditmanager.ListAssessmentControlInsightsByControlDomainInput{}
        output := &auditmanager.ListAssessmentControlInsightsByControlDomainOutput{}

        mockClient.On("ListAssessmentControlInsightsByControlDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentControlInsightsByControlDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentFrameworkShareRequests", func(t *testing.T) {
        input := &auditmanager.ListAssessmentFrameworkShareRequestsInput{}
        output := &auditmanager.ListAssessmentFrameworkShareRequestsOutput{}

        mockClient.On("ListAssessmentFrameworkShareRequests", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentFrameworkShareRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentFrameworks", func(t *testing.T) {
        input := &auditmanager.ListAssessmentFrameworksInput{}
        output := &auditmanager.ListAssessmentFrameworksOutput{}

        mockClient.On("ListAssessmentFrameworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentFrameworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentReports", func(t *testing.T) {
        input := &auditmanager.ListAssessmentReportsInput{}
        output := &auditmanager.ListAssessmentReportsOutput{}

        mockClient.On("ListAssessmentReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessments", func(t *testing.T) {
        input := &auditmanager.ListAssessmentsInput{}
        output := &auditmanager.ListAssessmentsOutput{}

        mockClient.On("ListAssessments", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControlDomainInsights", func(t *testing.T) {
        input := &auditmanager.ListControlDomainInsightsInput{}
        output := &auditmanager.ListControlDomainInsightsOutput{}

        mockClient.On("ListControlDomainInsights", ctx, input).Return(output, nil)

        result, err := mockClient.ListControlDomainInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControlDomainInsightsByAssessment", func(t *testing.T) {
        input := &auditmanager.ListControlDomainInsightsByAssessmentInput{}
        output := &auditmanager.ListControlDomainInsightsByAssessmentOutput{}

        mockClient.On("ListControlDomainInsightsByAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.ListControlDomainInsightsByAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControlInsightsByControlDomain", func(t *testing.T) {
        input := &auditmanager.ListControlInsightsByControlDomainInput{}
        output := &auditmanager.ListControlInsightsByControlDomainOutput{}

        mockClient.On("ListControlInsightsByControlDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListControlInsightsByControlDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControls", func(t *testing.T) {
        input := &auditmanager.ListControlsInput{}
        output := &auditmanager.ListControlsOutput{}

        mockClient.On("ListControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeywordsForDataSource", func(t *testing.T) {
        input := &auditmanager.ListKeywordsForDataSourceInput{}
        output := &auditmanager.ListKeywordsForDataSourceOutput{}

        mockClient.On("ListKeywordsForDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeywordsForDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotifications", func(t *testing.T) {
        input := &auditmanager.ListNotificationsInput{}
        output := &auditmanager.ListNotificationsOutput{}

        mockClient.On("ListNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &auditmanager.ListTagsForResourceInput{}
        output := &auditmanager.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterAccount", func(t *testing.T) {
        input := &auditmanager.RegisterAccountInput{}
        output := &auditmanager.RegisterAccountOutput{}

        mockClient.On("RegisterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterOrganizationAdminAccount", func(t *testing.T) {
        input := &auditmanager.RegisterOrganizationAdminAccountInput{}
        output := &auditmanager.RegisterOrganizationAdminAccountOutput{}

        mockClient.On("RegisterOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssessmentFrameworkShare", func(t *testing.T) {
        input := &auditmanager.StartAssessmentFrameworkShareInput{}
        output := &auditmanager.StartAssessmentFrameworkShareOutput{}

        mockClient.On("StartAssessmentFrameworkShare", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssessmentFrameworkShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &auditmanager.TagResourceInput{}
        output := &auditmanager.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &auditmanager.UntagResourceInput{}
        output := &auditmanager.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessment", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentInput{}
        output := &auditmanager.UpdateAssessmentOutput{}

        mockClient.On("UpdateAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentControl", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentControlInput{}
        output := &auditmanager.UpdateAssessmentControlOutput{}

        mockClient.On("UpdateAssessmentControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentControlSetStatus", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentControlSetStatusInput{}
        output := &auditmanager.UpdateAssessmentControlSetStatusOutput{}

        mockClient.On("UpdateAssessmentControlSetStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentControlSetStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentFramework", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentFrameworkInput{}
        output := &auditmanager.UpdateAssessmentFrameworkOutput{}

        mockClient.On("UpdateAssessmentFramework", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentFrameworkShare", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentFrameworkShareInput{}
        output := &auditmanager.UpdateAssessmentFrameworkShareOutput{}

        mockClient.On("UpdateAssessmentFrameworkShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentFrameworkShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentStatus", func(t *testing.T) {
        input := &auditmanager.UpdateAssessmentStatusInput{}
        output := &auditmanager.UpdateAssessmentStatusOutput{}

        mockClient.On("UpdateAssessmentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateControl", func(t *testing.T) {
        input := &auditmanager.UpdateControlInput{}
        output := &auditmanager.UpdateControlOutput{}

        mockClient.On("UpdateControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSettings", func(t *testing.T) {
        input := &auditmanager.UpdateSettingsInput{}
        output := &auditmanager.UpdateSettingsOutput{}

        mockClient.On("UpdateSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateAssessmentReportIntegrity", func(t *testing.T) {
        input := &auditmanager.ValidateAssessmentReportIntegrityInput{}
        output := &auditmanager.ValidateAssessmentReportIntegrityOutput{}

        mockClient.On("ValidateAssessmentReportIntegrity", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateAssessmentReportIntegrity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
