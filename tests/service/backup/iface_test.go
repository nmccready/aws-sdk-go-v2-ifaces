package backup_test

// tests for the backup service interface for this ../../../service/backup/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backup/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backup/backup_iface"
	"github.com/stretchr/testify/assert"
)

func TestBackupServiceCanBeMocked(t *testing.T) {
	var iface backup_iface.IClient
	iface = &backup.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := backup.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelLegalHold", func(t *testing.T) {
        input := &backup.CancelLegalHoldInput{}
        output := &backup.CancelLegalHoldOutput{}

        mockClient.On("CancelLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.CancelLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackupPlan", func(t *testing.T) {
        input := &backup.CreateBackupPlanInput{}
        output := &backup.CreateBackupPlanOutput{}

        mockClient.On("CreateBackupPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackupPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackupSelection", func(t *testing.T) {
        input := &backup.CreateBackupSelectionInput{}
        output := &backup.CreateBackupSelectionOutput{}

        mockClient.On("CreateBackupSelection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackupSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackupVault", func(t *testing.T) {
        input := &backup.CreateBackupVaultInput{}
        output := &backup.CreateBackupVaultOutput{}

        mockClient.On("CreateBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFramework", func(t *testing.T) {
        input := &backup.CreateFrameworkInput{}
        output := &backup.CreateFrameworkOutput{}

        mockClient.On("CreateFramework", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLegalHold", func(t *testing.T) {
        input := &backup.CreateLegalHoldInput{}
        output := &backup.CreateLegalHoldOutput{}

        mockClient.On("CreateLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogicallyAirGappedBackupVault", func(t *testing.T) {
        input := &backup.CreateLogicallyAirGappedBackupVaultInput{}
        output := &backup.CreateLogicallyAirGappedBackupVaultOutput{}

        mockClient.On("CreateLogicallyAirGappedBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogicallyAirGappedBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReportPlan", func(t *testing.T) {
        input := &backup.CreateReportPlanInput{}
        output := &backup.CreateReportPlanOutput{}

        mockClient.On("CreateReportPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReportPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRestoreTestingPlan", func(t *testing.T) {
        input := &backup.CreateRestoreTestingPlanInput{}
        output := &backup.CreateRestoreTestingPlanOutput{}

        mockClient.On("CreateRestoreTestingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRestoreTestingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRestoreTestingSelection", func(t *testing.T) {
        input := &backup.CreateRestoreTestingSelectionInput{}
        output := &backup.CreateRestoreTestingSelectionOutput{}

        mockClient.On("CreateRestoreTestingSelection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRestoreTestingSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupPlan", func(t *testing.T) {
        input := &backup.DeleteBackupPlanInput{}
        output := &backup.DeleteBackupPlanOutput{}

        mockClient.On("DeleteBackupPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupSelection", func(t *testing.T) {
        input := &backup.DeleteBackupSelectionInput{}
        output := &backup.DeleteBackupSelectionOutput{}

        mockClient.On("DeleteBackupSelection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupVault", func(t *testing.T) {
        input := &backup.DeleteBackupVaultInput{}
        output := &backup.DeleteBackupVaultOutput{}

        mockClient.On("DeleteBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupVaultAccessPolicy", func(t *testing.T) {
        input := &backup.DeleteBackupVaultAccessPolicyInput{}
        output := &backup.DeleteBackupVaultAccessPolicyOutput{}

        mockClient.On("DeleteBackupVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupVaultLockConfiguration", func(t *testing.T) {
        input := &backup.DeleteBackupVaultLockConfigurationInput{}
        output := &backup.DeleteBackupVaultLockConfigurationOutput{}

        mockClient.On("DeleteBackupVaultLockConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupVaultLockConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackupVaultNotifications", func(t *testing.T) {
        input := &backup.DeleteBackupVaultNotificationsInput{}
        output := &backup.DeleteBackupVaultNotificationsOutput{}

        mockClient.On("DeleteBackupVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackupVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFramework", func(t *testing.T) {
        input := &backup.DeleteFrameworkInput{}
        output := &backup.DeleteFrameworkOutput{}

        mockClient.On("DeleteFramework", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecoveryPoint", func(t *testing.T) {
        input := &backup.DeleteRecoveryPointInput{}
        output := &backup.DeleteRecoveryPointOutput{}

        mockClient.On("DeleteRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReportPlan", func(t *testing.T) {
        input := &backup.DeleteReportPlanInput{}
        output := &backup.DeleteReportPlanOutput{}

        mockClient.On("DeleteReportPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReportPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRestoreTestingPlan", func(t *testing.T) {
        input := &backup.DeleteRestoreTestingPlanInput{}
        output := &backup.DeleteRestoreTestingPlanOutput{}

        mockClient.On("DeleteRestoreTestingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRestoreTestingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRestoreTestingSelection", func(t *testing.T) {
        input := &backup.DeleteRestoreTestingSelectionInput{}
        output := &backup.DeleteRestoreTestingSelectionOutput{}

        mockClient.On("DeleteRestoreTestingSelection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRestoreTestingSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackupJob", func(t *testing.T) {
        input := &backup.DescribeBackupJobInput{}
        output := &backup.DescribeBackupJobOutput{}

        mockClient.On("DescribeBackupJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackupJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackupVault", func(t *testing.T) {
        input := &backup.DescribeBackupVaultInput{}
        output := &backup.DescribeBackupVaultOutput{}

        mockClient.On("DescribeBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCopyJob", func(t *testing.T) {
        input := &backup.DescribeCopyJobInput{}
        output := &backup.DescribeCopyJobOutput{}

        mockClient.On("DescribeCopyJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCopyJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFramework", func(t *testing.T) {
        input := &backup.DescribeFrameworkInput{}
        output := &backup.DescribeFrameworkOutput{}

        mockClient.On("DescribeFramework", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGlobalSettings", func(t *testing.T) {
        input := &backup.DescribeGlobalSettingsInput{}
        output := &backup.DescribeGlobalSettingsOutput{}

        mockClient.On("DescribeGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProtectedResource", func(t *testing.T) {
        input := &backup.DescribeProtectedResourceInput{}
        output := &backup.DescribeProtectedResourceOutput{}

        mockClient.On("DescribeProtectedResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProtectedResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecoveryPoint", func(t *testing.T) {
        input := &backup.DescribeRecoveryPointInput{}
        output := &backup.DescribeRecoveryPointOutput{}

        mockClient.On("DescribeRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegionSettings", func(t *testing.T) {
        input := &backup.DescribeRegionSettingsInput{}
        output := &backup.DescribeRegionSettingsOutput{}

        mockClient.On("DescribeRegionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReportJob", func(t *testing.T) {
        input := &backup.DescribeReportJobInput{}
        output := &backup.DescribeReportJobOutput{}

        mockClient.On("DescribeReportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReportPlan", func(t *testing.T) {
        input := &backup.DescribeReportPlanInput{}
        output := &backup.DescribeReportPlanOutput{}

        mockClient.On("DescribeReportPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReportPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRestoreJob", func(t *testing.T) {
        input := &backup.DescribeRestoreJobInput{}
        output := &backup.DescribeRestoreJobOutput{}

        mockClient.On("DescribeRestoreJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRestoreJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRecoveryPoint", func(t *testing.T) {
        input := &backup.DisassociateRecoveryPointInput{}
        output := &backup.DisassociateRecoveryPointOutput{}

        mockClient.On("DisassociateRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateRecoveryPointFromParent", func(t *testing.T) {
        input := &backup.DisassociateRecoveryPointFromParentInput{}
        output := &backup.DisassociateRecoveryPointFromParentOutput{}

        mockClient.On("DisassociateRecoveryPointFromParent", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateRecoveryPointFromParent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportBackupPlanTemplate", func(t *testing.T) {
        input := &backup.ExportBackupPlanTemplateInput{}
        output := &backup.ExportBackupPlanTemplateOutput{}

        mockClient.On("ExportBackupPlanTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ExportBackupPlanTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupPlan", func(t *testing.T) {
        input := &backup.GetBackupPlanInput{}
        output := &backup.GetBackupPlanOutput{}

        mockClient.On("GetBackupPlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupPlanFromJSON", func(t *testing.T) {
        input := &backup.GetBackupPlanFromJSONInput{}
        output := &backup.GetBackupPlanFromJSONOutput{}

        mockClient.On("GetBackupPlanFromJSON", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupPlanFromJSON(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupPlanFromTemplate", func(t *testing.T) {
        input := &backup.GetBackupPlanFromTemplateInput{}
        output := &backup.GetBackupPlanFromTemplateOutput{}

        mockClient.On("GetBackupPlanFromTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupPlanFromTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupSelection", func(t *testing.T) {
        input := &backup.GetBackupSelectionInput{}
        output := &backup.GetBackupSelectionOutput{}

        mockClient.On("GetBackupSelection", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupVaultAccessPolicy", func(t *testing.T) {
        input := &backup.GetBackupVaultAccessPolicyInput{}
        output := &backup.GetBackupVaultAccessPolicyOutput{}

        mockClient.On("GetBackupVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackupVaultNotifications", func(t *testing.T) {
        input := &backup.GetBackupVaultNotificationsInput{}
        output := &backup.GetBackupVaultNotificationsOutput{}

        mockClient.On("GetBackupVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackupVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLegalHold", func(t *testing.T) {
        input := &backup.GetLegalHoldInput{}
        output := &backup.GetLegalHoldOutput{}

        mockClient.On("GetLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.GetLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecoveryPointRestoreMetadata", func(t *testing.T) {
        input := &backup.GetRecoveryPointRestoreMetadataInput{}
        output := &backup.GetRecoveryPointRestoreMetadataOutput{}

        mockClient.On("GetRecoveryPointRestoreMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecoveryPointRestoreMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestoreJobMetadata", func(t *testing.T) {
        input := &backup.GetRestoreJobMetadataInput{}
        output := &backup.GetRestoreJobMetadataOutput{}

        mockClient.On("GetRestoreJobMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestoreJobMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestoreTestingInferredMetadata", func(t *testing.T) {
        input := &backup.GetRestoreTestingInferredMetadataInput{}
        output := &backup.GetRestoreTestingInferredMetadataOutput{}

        mockClient.On("GetRestoreTestingInferredMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestoreTestingInferredMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestoreTestingPlan", func(t *testing.T) {
        input := &backup.GetRestoreTestingPlanInput{}
        output := &backup.GetRestoreTestingPlanOutput{}

        mockClient.On("GetRestoreTestingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestoreTestingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestoreTestingSelection", func(t *testing.T) {
        input := &backup.GetRestoreTestingSelectionInput{}
        output := &backup.GetRestoreTestingSelectionOutput{}

        mockClient.On("GetRestoreTestingSelection", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestoreTestingSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSupportedResourceTypes", func(t *testing.T) {
        input := &backup.GetSupportedResourceTypesInput{}
        output := &backup.GetSupportedResourceTypesOutput{}

        mockClient.On("GetSupportedResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetSupportedResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupJobSummaries", func(t *testing.T) {
        input := &backup.ListBackupJobSummariesInput{}
        output := &backup.ListBackupJobSummariesOutput{}

        mockClient.On("ListBackupJobSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupJobSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupJobs", func(t *testing.T) {
        input := &backup.ListBackupJobsInput{}
        output := &backup.ListBackupJobsOutput{}

        mockClient.On("ListBackupJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupPlanTemplates", func(t *testing.T) {
        input := &backup.ListBackupPlanTemplatesInput{}
        output := &backup.ListBackupPlanTemplatesOutput{}

        mockClient.On("ListBackupPlanTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupPlanTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupPlanVersions", func(t *testing.T) {
        input := &backup.ListBackupPlanVersionsInput{}
        output := &backup.ListBackupPlanVersionsOutput{}

        mockClient.On("ListBackupPlanVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupPlanVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupPlans", func(t *testing.T) {
        input := &backup.ListBackupPlansInput{}
        output := &backup.ListBackupPlansOutput{}

        mockClient.On("ListBackupPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupSelections", func(t *testing.T) {
        input := &backup.ListBackupSelectionsInput{}
        output := &backup.ListBackupSelectionsOutput{}

        mockClient.On("ListBackupSelections", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupSelections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackupVaults", func(t *testing.T) {
        input := &backup.ListBackupVaultsInput{}
        output := &backup.ListBackupVaultsOutput{}

        mockClient.On("ListBackupVaults", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackupVaults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCopyJobSummaries", func(t *testing.T) {
        input := &backup.ListCopyJobSummariesInput{}
        output := &backup.ListCopyJobSummariesOutput{}

        mockClient.On("ListCopyJobSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListCopyJobSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCopyJobs", func(t *testing.T) {
        input := &backup.ListCopyJobsInput{}
        output := &backup.ListCopyJobsOutput{}

        mockClient.On("ListCopyJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListCopyJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFrameworks", func(t *testing.T) {
        input := &backup.ListFrameworksInput{}
        output := &backup.ListFrameworksOutput{}

        mockClient.On("ListFrameworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListFrameworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLegalHolds", func(t *testing.T) {
        input := &backup.ListLegalHoldsInput{}
        output := &backup.ListLegalHoldsOutput{}

        mockClient.On("ListLegalHolds", ctx, input).Return(output, nil)

        result, err := mockClient.ListLegalHolds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtectedResources", func(t *testing.T) {
        input := &backup.ListProtectedResourcesInput{}
        output := &backup.ListProtectedResourcesOutput{}

        mockClient.On("ListProtectedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtectedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtectedResourcesByBackupVault", func(t *testing.T) {
        input := &backup.ListProtectedResourcesByBackupVaultInput{}
        output := &backup.ListProtectedResourcesByBackupVaultOutput{}

        mockClient.On("ListProtectedResourcesByBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtectedResourcesByBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecoveryPointsByBackupVault", func(t *testing.T) {
        input := &backup.ListRecoveryPointsByBackupVaultInput{}
        output := &backup.ListRecoveryPointsByBackupVaultOutput{}

        mockClient.On("ListRecoveryPointsByBackupVault", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecoveryPointsByBackupVault(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecoveryPointsByLegalHold", func(t *testing.T) {
        input := &backup.ListRecoveryPointsByLegalHoldInput{}
        output := &backup.ListRecoveryPointsByLegalHoldOutput{}

        mockClient.On("ListRecoveryPointsByLegalHold", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecoveryPointsByLegalHold(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecoveryPointsByResource", func(t *testing.T) {
        input := &backup.ListRecoveryPointsByResourceInput{}
        output := &backup.ListRecoveryPointsByResourceOutput{}

        mockClient.On("ListRecoveryPointsByResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecoveryPointsByResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReportJobs", func(t *testing.T) {
        input := &backup.ListReportJobsInput{}
        output := &backup.ListReportJobsOutput{}

        mockClient.On("ListReportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReportPlans", func(t *testing.T) {
        input := &backup.ListReportPlansInput{}
        output := &backup.ListReportPlansOutput{}

        mockClient.On("ListReportPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListReportPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRestoreJobSummaries", func(t *testing.T) {
        input := &backup.ListRestoreJobSummariesInput{}
        output := &backup.ListRestoreJobSummariesOutput{}

        mockClient.On("ListRestoreJobSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListRestoreJobSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRestoreJobs", func(t *testing.T) {
        input := &backup.ListRestoreJobsInput{}
        output := &backup.ListRestoreJobsOutput{}

        mockClient.On("ListRestoreJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListRestoreJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRestoreJobsByProtectedResource", func(t *testing.T) {
        input := &backup.ListRestoreJobsByProtectedResourceInput{}
        output := &backup.ListRestoreJobsByProtectedResourceOutput{}

        mockClient.On("ListRestoreJobsByProtectedResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListRestoreJobsByProtectedResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRestoreTestingPlans", func(t *testing.T) {
        input := &backup.ListRestoreTestingPlansInput{}
        output := &backup.ListRestoreTestingPlansOutput{}

        mockClient.On("ListRestoreTestingPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListRestoreTestingPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRestoreTestingSelections", func(t *testing.T) {
        input := &backup.ListRestoreTestingSelectionsInput{}
        output := &backup.ListRestoreTestingSelectionsOutput{}

        mockClient.On("ListRestoreTestingSelections", ctx, input).Return(output, nil)

        result, err := mockClient.ListRestoreTestingSelections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &backup.ListTagsInput{}
        output := &backup.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBackupVaultAccessPolicy", func(t *testing.T) {
        input := &backup.PutBackupVaultAccessPolicyInput{}
        output := &backup.PutBackupVaultAccessPolicyOutput{}

        mockClient.On("PutBackupVaultAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutBackupVaultAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBackupVaultLockConfiguration", func(t *testing.T) {
        input := &backup.PutBackupVaultLockConfigurationInput{}
        output := &backup.PutBackupVaultLockConfigurationOutput{}

        mockClient.On("PutBackupVaultLockConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBackupVaultLockConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBackupVaultNotifications", func(t *testing.T) {
        input := &backup.PutBackupVaultNotificationsInput{}
        output := &backup.PutBackupVaultNotificationsOutput{}

        mockClient.On("PutBackupVaultNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.PutBackupVaultNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRestoreValidationResult", func(t *testing.T) {
        input := &backup.PutRestoreValidationResultInput{}
        output := &backup.PutRestoreValidationResultOutput{}

        mockClient.On("PutRestoreValidationResult", ctx, input).Return(output, nil)

        result, err := mockClient.PutRestoreValidationResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBackupJob", func(t *testing.T) {
        input := &backup.StartBackupJobInput{}
        output := &backup.StartBackupJobOutput{}

        mockClient.On("StartBackupJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartBackupJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCopyJob", func(t *testing.T) {
        input := &backup.StartCopyJobInput{}
        output := &backup.StartCopyJobOutput{}

        mockClient.On("StartCopyJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartCopyJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReportJob", func(t *testing.T) {
        input := &backup.StartReportJobInput{}
        output := &backup.StartReportJobOutput{}

        mockClient.On("StartReportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartReportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRestoreJob", func(t *testing.T) {
        input := &backup.StartRestoreJobInput{}
        output := &backup.StartRestoreJobOutput{}

        mockClient.On("StartRestoreJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartRestoreJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopBackupJob", func(t *testing.T) {
        input := &backup.StopBackupJobInput{}
        output := &backup.StopBackupJobOutput{}

        mockClient.On("StopBackupJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopBackupJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &backup.TagResourceInput{}
        output := &backup.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &backup.UntagResourceInput{}
        output := &backup.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackupPlan", func(t *testing.T) {
        input := &backup.UpdateBackupPlanInput{}
        output := &backup.UpdateBackupPlanOutput{}

        mockClient.On("UpdateBackupPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackupPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFramework", func(t *testing.T) {
        input := &backup.UpdateFrameworkInput{}
        output := &backup.UpdateFrameworkOutput{}

        mockClient.On("UpdateFramework", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFramework(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalSettings", func(t *testing.T) {
        input := &backup.UpdateGlobalSettingsInput{}
        output := &backup.UpdateGlobalSettingsOutput{}

        mockClient.On("UpdateGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecoveryPointLifecycle", func(t *testing.T) {
        input := &backup.UpdateRecoveryPointLifecycleInput{}
        output := &backup.UpdateRecoveryPointLifecycleOutput{}

        mockClient.On("UpdateRecoveryPointLifecycle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecoveryPointLifecycle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegionSettings", func(t *testing.T) {
        input := &backup.UpdateRegionSettingsInput{}
        output := &backup.UpdateRegionSettingsOutput{}

        mockClient.On("UpdateRegionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReportPlan", func(t *testing.T) {
        input := &backup.UpdateReportPlanInput{}
        output := &backup.UpdateReportPlanOutput{}

        mockClient.On("UpdateReportPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReportPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRestoreTestingPlan", func(t *testing.T) {
        input := &backup.UpdateRestoreTestingPlanInput{}
        output := &backup.UpdateRestoreTestingPlanOutput{}

        mockClient.On("UpdateRestoreTestingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRestoreTestingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRestoreTestingSelection", func(t *testing.T) {
        input := &backup.UpdateRestoreTestingSelectionInput{}
        output := &backup.UpdateRestoreTestingSelectionOutput{}

        mockClient.On("UpdateRestoreTestingSelection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRestoreTestingSelection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
