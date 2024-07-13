package mgn_test

// tests for the mgn service interface for this ../../../service/mgn/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mgn"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mgn/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mgn/mgn_iface"
	"github.com/stretchr/testify/assert"
)

func TestMgnServiceCanBeMocked(t *testing.T) {
	var iface mgn_iface.IClient
	iface = &mgn.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mgn.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestArchiveApplication", func(t *testing.T) {
        input := &mgn.ArchiveApplicationInput{}
        output := &mgn.ArchiveApplicationOutput{}

        mockClient.On("ArchiveApplication", ctx, input).Return(output, nil)

        result, err := mockClient.ArchiveApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestArchiveWave", func(t *testing.T) {
        input := &mgn.ArchiveWaveInput{}
        output := &mgn.ArchiveWaveOutput{}

        mockClient.On("ArchiveWave", ctx, input).Return(output, nil)

        result, err := mockClient.ArchiveWave(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApplications", func(t *testing.T) {
        input := &mgn.AssociateApplicationsInput{}
        output := &mgn.AssociateApplicationsOutput{}

        mockClient.On("AssociateApplications", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSourceServers", func(t *testing.T) {
        input := &mgn.AssociateSourceServersInput{}
        output := &mgn.AssociateSourceServersOutput{}

        mockClient.On("AssociateSourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeServerLifeCycleState", func(t *testing.T) {
        input := &mgn.ChangeServerLifeCycleStateInput{}
        output := &mgn.ChangeServerLifeCycleStateOutput{}

        mockClient.On("ChangeServerLifeCycleState", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeServerLifeCycleState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &mgn.CreateApplicationInput{}
        output := &mgn.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnector", func(t *testing.T) {
        input := &mgn.CreateConnectorInput{}
        output := &mgn.CreateConnectorOutput{}

        mockClient.On("CreateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchConfigurationTemplate", func(t *testing.T) {
        input := &mgn.CreateLaunchConfigurationTemplateInput{}
        output := &mgn.CreateLaunchConfigurationTemplateOutput{}

        mockClient.On("CreateLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationConfigurationTemplate", func(t *testing.T) {
        input := &mgn.CreateReplicationConfigurationTemplateInput{}
        output := &mgn.CreateReplicationConfigurationTemplateOutput{}

        mockClient.On("CreateReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWave", func(t *testing.T) {
        input := &mgn.CreateWaveInput{}
        output := &mgn.CreateWaveOutput{}

        mockClient.On("CreateWave", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWave(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &mgn.DeleteApplicationInput{}
        output := &mgn.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnector", func(t *testing.T) {
        input := &mgn.DeleteConnectorInput{}
        output := &mgn.DeleteConnectorOutput{}

        mockClient.On("DeleteConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &mgn.DeleteJobInput{}
        output := &mgn.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchConfigurationTemplate", func(t *testing.T) {
        input := &mgn.DeleteLaunchConfigurationTemplateInput{}
        output := &mgn.DeleteLaunchConfigurationTemplateOutput{}

        mockClient.On("DeleteLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationConfigurationTemplate", func(t *testing.T) {
        input := &mgn.DeleteReplicationConfigurationTemplateInput{}
        output := &mgn.DeleteReplicationConfigurationTemplateOutput{}

        mockClient.On("DeleteReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceServer", func(t *testing.T) {
        input := &mgn.DeleteSourceServerInput{}
        output := &mgn.DeleteSourceServerOutput{}

        mockClient.On("DeleteSourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVcenterClient", func(t *testing.T) {
        input := &mgn.DeleteVcenterClientInput{}
        output := &mgn.DeleteVcenterClientOutput{}

        mockClient.On("DeleteVcenterClient", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVcenterClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWave", func(t *testing.T) {
        input := &mgn.DeleteWaveInput{}
        output := &mgn.DeleteWaveOutput{}

        mockClient.On("DeleteWave", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWave(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobLogItems", func(t *testing.T) {
        input := &mgn.DescribeJobLogItemsInput{}
        output := &mgn.DescribeJobLogItemsOutput{}

        mockClient.On("DescribeJobLogItems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobLogItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobs", func(t *testing.T) {
        input := &mgn.DescribeJobsInput{}
        output := &mgn.DescribeJobsOutput{}

        mockClient.On("DescribeJobs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLaunchConfigurationTemplates", func(t *testing.T) {
        input := &mgn.DescribeLaunchConfigurationTemplatesInput{}
        output := &mgn.DescribeLaunchConfigurationTemplatesOutput{}

        mockClient.On("DescribeLaunchConfigurationTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLaunchConfigurationTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationConfigurationTemplates", func(t *testing.T) {
        input := &mgn.DescribeReplicationConfigurationTemplatesInput{}
        output := &mgn.DescribeReplicationConfigurationTemplatesOutput{}

        mockClient.On("DescribeReplicationConfigurationTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationConfigurationTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSourceServers", func(t *testing.T) {
        input := &mgn.DescribeSourceServersInput{}
        output := &mgn.DescribeSourceServersOutput{}

        mockClient.On("DescribeSourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVcenterClients", func(t *testing.T) {
        input := &mgn.DescribeVcenterClientsInput{}
        output := &mgn.DescribeVcenterClientsOutput{}

        mockClient.On("DescribeVcenterClients", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVcenterClients(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApplications", func(t *testing.T) {
        input := &mgn.DisassociateApplicationsInput{}
        output := &mgn.DisassociateApplicationsOutput{}

        mockClient.On("DisassociateApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSourceServers", func(t *testing.T) {
        input := &mgn.DisassociateSourceServersInput{}
        output := &mgn.DisassociateSourceServersOutput{}

        mockClient.On("DisassociateSourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectFromService", func(t *testing.T) {
        input := &mgn.DisconnectFromServiceInput{}
        output := &mgn.DisconnectFromServiceOutput{}

        mockClient.On("DisconnectFromService", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectFromService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFinalizeCutover", func(t *testing.T) {
        input := &mgn.FinalizeCutoverInput{}
        output := &mgn.FinalizeCutoverOutput{}

        mockClient.On("FinalizeCutover", ctx, input).Return(output, nil)

        result, err := mockClient.FinalizeCutover(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchConfiguration", func(t *testing.T) {
        input := &mgn.GetLaunchConfigurationInput{}
        output := &mgn.GetLaunchConfigurationOutput{}

        mockClient.On("GetLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReplicationConfiguration", func(t *testing.T) {
        input := &mgn.GetReplicationConfigurationInput{}
        output := &mgn.GetReplicationConfigurationOutput{}

        mockClient.On("GetReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitializeService", func(t *testing.T) {
        input := &mgn.InitializeServiceInput{}
        output := &mgn.InitializeServiceOutput{}

        mockClient.On("InitializeService", ctx, input).Return(output, nil)

        result, err := mockClient.InitializeService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &mgn.ListApplicationsInput{}
        output := &mgn.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &mgn.ListConnectorsInput{}
        output := &mgn.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExportErrors", func(t *testing.T) {
        input := &mgn.ListExportErrorsInput{}
        output := &mgn.ListExportErrorsOutput{}

        mockClient.On("ListExportErrors", ctx, input).Return(output, nil)

        result, err := mockClient.ListExportErrors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExports", func(t *testing.T) {
        input := &mgn.ListExportsInput{}
        output := &mgn.ListExportsOutput{}

        mockClient.On("ListExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportErrors", func(t *testing.T) {
        input := &mgn.ListImportErrorsInput{}
        output := &mgn.ListImportErrorsOutput{}

        mockClient.On("ListImportErrors", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportErrors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImports", func(t *testing.T) {
        input := &mgn.ListImportsInput{}
        output := &mgn.ListImportsOutput{}

        mockClient.On("ListImports", ctx, input).Return(output, nil)

        result, err := mockClient.ListImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedAccounts", func(t *testing.T) {
        input := &mgn.ListManagedAccountsInput{}
        output := &mgn.ListManagedAccountsOutput{}

        mockClient.On("ListManagedAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceServerActions", func(t *testing.T) {
        input := &mgn.ListSourceServerActionsInput{}
        output := &mgn.ListSourceServerActionsOutput{}

        mockClient.On("ListSourceServerActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceServerActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mgn.ListTagsForResourceInput{}
        output := &mgn.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateActions", func(t *testing.T) {
        input := &mgn.ListTemplateActionsInput{}
        output := &mgn.ListTemplateActionsOutput{}

        mockClient.On("ListTemplateActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWaves", func(t *testing.T) {
        input := &mgn.ListWavesInput{}
        output := &mgn.ListWavesOutput{}

        mockClient.On("ListWaves", ctx, input).Return(output, nil)

        result, err := mockClient.ListWaves(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMarkAsArchived", func(t *testing.T) {
        input := &mgn.MarkAsArchivedInput{}
        output := &mgn.MarkAsArchivedOutput{}

        mockClient.On("MarkAsArchived", ctx, input).Return(output, nil)

        result, err := mockClient.MarkAsArchived(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseReplication", func(t *testing.T) {
        input := &mgn.PauseReplicationInput{}
        output := &mgn.PauseReplicationOutput{}

        mockClient.On("PauseReplication", ctx, input).Return(output, nil)

        result, err := mockClient.PauseReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSourceServerAction", func(t *testing.T) {
        input := &mgn.PutSourceServerActionInput{}
        output := &mgn.PutSourceServerActionOutput{}

        mockClient.On("PutSourceServerAction", ctx, input).Return(output, nil)

        result, err := mockClient.PutSourceServerAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTemplateAction", func(t *testing.T) {
        input := &mgn.PutTemplateActionInput{}
        output := &mgn.PutTemplateActionOutput{}

        mockClient.On("PutTemplateAction", ctx, input).Return(output, nil)

        result, err := mockClient.PutTemplateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveSourceServerAction", func(t *testing.T) {
        input := &mgn.RemoveSourceServerActionInput{}
        output := &mgn.RemoveSourceServerActionOutput{}

        mockClient.On("RemoveSourceServerAction", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveSourceServerAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTemplateAction", func(t *testing.T) {
        input := &mgn.RemoveTemplateActionInput{}
        output := &mgn.RemoveTemplateActionOutput{}

        mockClient.On("RemoveTemplateAction", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTemplateAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeReplication", func(t *testing.T) {
        input := &mgn.ResumeReplicationInput{}
        output := &mgn.ResumeReplicationOutput{}

        mockClient.On("ResumeReplication", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryDataReplication", func(t *testing.T) {
        input := &mgn.RetryDataReplicationInput{}
        output := &mgn.RetryDataReplicationOutput{}

        mockClient.On("RetryDataReplication", ctx, input).Return(output, nil)

        result, err := mockClient.RetryDataReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCutover", func(t *testing.T) {
        input := &mgn.StartCutoverInput{}
        output := &mgn.StartCutoverOutput{}

        mockClient.On("StartCutover", ctx, input).Return(output, nil)

        result, err := mockClient.StartCutover(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExport", func(t *testing.T) {
        input := &mgn.StartExportInput{}
        output := &mgn.StartExportOutput{}

        mockClient.On("StartExport", ctx, input).Return(output, nil)

        result, err := mockClient.StartExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImport", func(t *testing.T) {
        input := &mgn.StartImportInput{}
        output := &mgn.StartImportOutput{}

        mockClient.On("StartImport", ctx, input).Return(output, nil)

        result, err := mockClient.StartImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplication", func(t *testing.T) {
        input := &mgn.StartReplicationInput{}
        output := &mgn.StartReplicationOutput{}

        mockClient.On("StartReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTest", func(t *testing.T) {
        input := &mgn.StartTestInput{}
        output := &mgn.StartTestOutput{}

        mockClient.On("StartTest", ctx, input).Return(output, nil)

        result, err := mockClient.StartTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopReplication", func(t *testing.T) {
        input := &mgn.StopReplicationInput{}
        output := &mgn.StopReplicationOutput{}

        mockClient.On("StopReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mgn.TagResourceInput{}
        output := &mgn.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateTargetInstances", func(t *testing.T) {
        input := &mgn.TerminateTargetInstancesInput{}
        output := &mgn.TerminateTargetInstancesOutput{}

        mockClient.On("TerminateTargetInstances", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateTargetInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnarchiveApplication", func(t *testing.T) {
        input := &mgn.UnarchiveApplicationInput{}
        output := &mgn.UnarchiveApplicationOutput{}

        mockClient.On("UnarchiveApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UnarchiveApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnarchiveWave", func(t *testing.T) {
        input := &mgn.UnarchiveWaveInput{}
        output := &mgn.UnarchiveWaveOutput{}

        mockClient.On("UnarchiveWave", ctx, input).Return(output, nil)

        result, err := mockClient.UnarchiveWave(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mgn.UntagResourceInput{}
        output := &mgn.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &mgn.UpdateApplicationInput{}
        output := &mgn.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnector", func(t *testing.T) {
        input := &mgn.UpdateConnectorInput{}
        output := &mgn.UpdateConnectorOutput{}

        mockClient.On("UpdateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchConfiguration", func(t *testing.T) {
        input := &mgn.UpdateLaunchConfigurationInput{}
        output := &mgn.UpdateLaunchConfigurationOutput{}

        mockClient.On("UpdateLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchConfigurationTemplate", func(t *testing.T) {
        input := &mgn.UpdateLaunchConfigurationTemplateInput{}
        output := &mgn.UpdateLaunchConfigurationTemplateOutput{}

        mockClient.On("UpdateLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationConfiguration", func(t *testing.T) {
        input := &mgn.UpdateReplicationConfigurationInput{}
        output := &mgn.UpdateReplicationConfigurationOutput{}

        mockClient.On("UpdateReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationConfigurationTemplate", func(t *testing.T) {
        input := &mgn.UpdateReplicationConfigurationTemplateInput{}
        output := &mgn.UpdateReplicationConfigurationTemplateOutput{}

        mockClient.On("UpdateReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSourceServer", func(t *testing.T) {
        input := &mgn.UpdateSourceServerInput{}
        output := &mgn.UpdateSourceServerOutput{}

        mockClient.On("UpdateSourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSourceServerReplicationType", func(t *testing.T) {
        input := &mgn.UpdateSourceServerReplicationTypeInput{}
        output := &mgn.UpdateSourceServerReplicationTypeOutput{}

        mockClient.On("UpdateSourceServerReplicationType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSourceServerReplicationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWave", func(t *testing.T) {
        input := &mgn.UpdateWaveInput{}
        output := &mgn.UpdateWaveOutput{}

        mockClient.On("UpdateWave", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWave(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
