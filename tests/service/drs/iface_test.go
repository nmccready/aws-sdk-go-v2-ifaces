// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package drs_test

// tests for the drs service interface for 
// this ../../../service/drs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/drs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/drs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/drs/drs_iface"
	"github.com/stretchr/testify/assert"
)

func TestDrsServiceCanBeMocked(t *testing.T) {
	var iface drs_iface.IClient
	iface = &drs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := drs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSourceNetworkStack", func(t *testing.T) {
        input := &drs.AssociateSourceNetworkStackInput{}
        output := &drs.AssociateSourceNetworkStackOutput{}

        mockClient.On("AssociateSourceNetworkStack", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSourceNetworkStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExtendedSourceServer", func(t *testing.T) {
        input := &drs.CreateExtendedSourceServerInput{}
        output := &drs.CreateExtendedSourceServerOutput{}

        mockClient.On("CreateExtendedSourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExtendedSourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchConfigurationTemplate", func(t *testing.T) {
        input := &drs.CreateLaunchConfigurationTemplateInput{}
        output := &drs.CreateLaunchConfigurationTemplateOutput{}

        mockClient.On("CreateLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationConfigurationTemplate", func(t *testing.T) {
        input := &drs.CreateReplicationConfigurationTemplateInput{}
        output := &drs.CreateReplicationConfigurationTemplateOutput{}

        mockClient.On("CreateReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSourceNetwork", func(t *testing.T) {
        input := &drs.CreateSourceNetworkInput{}
        output := &drs.CreateSourceNetworkOutput{}

        mockClient.On("CreateSourceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSourceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &drs.DeleteJobInput{}
        output := &drs.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchAction", func(t *testing.T) {
        input := &drs.DeleteLaunchActionInput{}
        output := &drs.DeleteLaunchActionOutput{}

        mockClient.On("DeleteLaunchAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchConfigurationTemplate", func(t *testing.T) {
        input := &drs.DeleteLaunchConfigurationTemplateInput{}
        output := &drs.DeleteLaunchConfigurationTemplateOutput{}

        mockClient.On("DeleteLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecoveryInstance", func(t *testing.T) {
        input := &drs.DeleteRecoveryInstanceInput{}
        output := &drs.DeleteRecoveryInstanceOutput{}

        mockClient.On("DeleteRecoveryInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecoveryInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationConfigurationTemplate", func(t *testing.T) {
        input := &drs.DeleteReplicationConfigurationTemplateInput{}
        output := &drs.DeleteReplicationConfigurationTemplateOutput{}

        mockClient.On("DeleteReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceNetwork", func(t *testing.T) {
        input := &drs.DeleteSourceNetworkInput{}
        output := &drs.DeleteSourceNetworkOutput{}

        mockClient.On("DeleteSourceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceServer", func(t *testing.T) {
        input := &drs.DeleteSourceServerInput{}
        output := &drs.DeleteSourceServerOutput{}

        mockClient.On("DeleteSourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobLogItems", func(t *testing.T) {
        input := &drs.DescribeJobLogItemsInput{}
        output := &drs.DescribeJobLogItemsOutput{}

        mockClient.On("DescribeJobLogItems", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobLogItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobs", func(t *testing.T) {
        input := &drs.DescribeJobsInput{}
        output := &drs.DescribeJobsOutput{}

        mockClient.On("DescribeJobs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLaunchConfigurationTemplates", func(t *testing.T) {
        input := &drs.DescribeLaunchConfigurationTemplatesInput{}
        output := &drs.DescribeLaunchConfigurationTemplatesOutput{}

        mockClient.On("DescribeLaunchConfigurationTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLaunchConfigurationTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecoveryInstances", func(t *testing.T) {
        input := &drs.DescribeRecoveryInstancesInput{}
        output := &drs.DescribeRecoveryInstancesOutput{}

        mockClient.On("DescribeRecoveryInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecoveryInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecoverySnapshots", func(t *testing.T) {
        input := &drs.DescribeRecoverySnapshotsInput{}
        output := &drs.DescribeRecoverySnapshotsOutput{}

        mockClient.On("DescribeRecoverySnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecoverySnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicationConfigurationTemplates", func(t *testing.T) {
        input := &drs.DescribeReplicationConfigurationTemplatesInput{}
        output := &drs.DescribeReplicationConfigurationTemplatesOutput{}

        mockClient.On("DescribeReplicationConfigurationTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicationConfigurationTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSourceNetworks", func(t *testing.T) {
        input := &drs.DescribeSourceNetworksInput{}
        output := &drs.DescribeSourceNetworksOutput{}

        mockClient.On("DescribeSourceNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSourceNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSourceServers", func(t *testing.T) {
        input := &drs.DescribeSourceServersInput{}
        output := &drs.DescribeSourceServersOutput{}

        mockClient.On("DescribeSourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectRecoveryInstance", func(t *testing.T) {
        input := &drs.DisconnectRecoveryInstanceInput{}
        output := &drs.DisconnectRecoveryInstanceOutput{}

        mockClient.On("DisconnectRecoveryInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectRecoveryInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectSourceServer", func(t *testing.T) {
        input := &drs.DisconnectSourceServerInput{}
        output := &drs.DisconnectSourceServerOutput{}

        mockClient.On("DisconnectSourceServer", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectSourceServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportSourceNetworkCfnTemplate", func(t *testing.T) {
        input := &drs.ExportSourceNetworkCfnTemplateInput{}
        output := &drs.ExportSourceNetworkCfnTemplateOutput{}

        mockClient.On("ExportSourceNetworkCfnTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ExportSourceNetworkCfnTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFailbackReplicationConfiguration", func(t *testing.T) {
        input := &drs.GetFailbackReplicationConfigurationInput{}
        output := &drs.GetFailbackReplicationConfigurationOutput{}

        mockClient.On("GetFailbackReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetFailbackReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchConfiguration", func(t *testing.T) {
        input := &drs.GetLaunchConfigurationInput{}
        output := &drs.GetLaunchConfigurationOutput{}

        mockClient.On("GetLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReplicationConfiguration", func(t *testing.T) {
        input := &drs.GetReplicationConfigurationInput{}
        output := &drs.GetReplicationConfigurationOutput{}

        mockClient.On("GetReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInitializeService", func(t *testing.T) {
        input := &drs.InitializeServiceInput{}
        output := &drs.InitializeServiceOutput{}

        mockClient.On("InitializeService", ctx, input).Return(output, nil)

        result, err := mockClient.InitializeService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExtensibleSourceServers", func(t *testing.T) {
        input := &drs.ListExtensibleSourceServersInput{}
        output := &drs.ListExtensibleSourceServersOutput{}

        mockClient.On("ListExtensibleSourceServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListExtensibleSourceServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLaunchActions", func(t *testing.T) {
        input := &drs.ListLaunchActionsInput{}
        output := &drs.ListLaunchActionsOutput{}

        mockClient.On("ListLaunchActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLaunchActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStagingAccounts", func(t *testing.T) {
        input := &drs.ListStagingAccountsInput{}
        output := &drs.ListStagingAccountsOutput{}

        mockClient.On("ListStagingAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListStagingAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &drs.ListTagsForResourceInput{}
        output := &drs.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLaunchAction", func(t *testing.T) {
        input := &drs.PutLaunchActionInput{}
        output := &drs.PutLaunchActionOutput{}

        mockClient.On("PutLaunchAction", ctx, input).Return(output, nil)

        result, err := mockClient.PutLaunchAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetryDataReplication", func(t *testing.T) {
        input := &drs.RetryDataReplicationInput{}
        output := &drs.RetryDataReplicationOutput{}

        mockClient.On("RetryDataReplication", ctx, input).Return(output, nil)

        result, err := mockClient.RetryDataReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReverseReplication", func(t *testing.T) {
        input := &drs.ReverseReplicationInput{}
        output := &drs.ReverseReplicationOutput{}

        mockClient.On("ReverseReplication", ctx, input).Return(output, nil)

        result, err := mockClient.ReverseReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFailbackLaunch", func(t *testing.T) {
        input := &drs.StartFailbackLaunchInput{}
        output := &drs.StartFailbackLaunchOutput{}

        mockClient.On("StartFailbackLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.StartFailbackLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRecovery", func(t *testing.T) {
        input := &drs.StartRecoveryInput{}
        output := &drs.StartRecoveryOutput{}

        mockClient.On("StartRecovery", ctx, input).Return(output, nil)

        result, err := mockClient.StartRecovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReplication", func(t *testing.T) {
        input := &drs.StartReplicationInput{}
        output := &drs.StartReplicationOutput{}

        mockClient.On("StartReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSourceNetworkRecovery", func(t *testing.T) {
        input := &drs.StartSourceNetworkRecoveryInput{}
        output := &drs.StartSourceNetworkRecoveryOutput{}

        mockClient.On("StartSourceNetworkRecovery", ctx, input).Return(output, nil)

        result, err := mockClient.StartSourceNetworkRecovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSourceNetworkReplication", func(t *testing.T) {
        input := &drs.StartSourceNetworkReplicationInput{}
        output := &drs.StartSourceNetworkReplicationOutput{}

        mockClient.On("StartSourceNetworkReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartSourceNetworkReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopFailback", func(t *testing.T) {
        input := &drs.StopFailbackInput{}
        output := &drs.StopFailbackOutput{}

        mockClient.On("StopFailback", ctx, input).Return(output, nil)

        result, err := mockClient.StopFailback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopReplication", func(t *testing.T) {
        input := &drs.StopReplicationInput{}
        output := &drs.StopReplicationOutput{}

        mockClient.On("StopReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSourceNetworkReplication", func(t *testing.T) {
        input := &drs.StopSourceNetworkReplicationInput{}
        output := &drs.StopSourceNetworkReplicationOutput{}

        mockClient.On("StopSourceNetworkReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopSourceNetworkReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &drs.TagResourceInput{}
        output := &drs.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateRecoveryInstances", func(t *testing.T) {
        input := &drs.TerminateRecoveryInstancesInput{}
        output := &drs.TerminateRecoveryInstancesOutput{}

        mockClient.On("TerminateRecoveryInstances", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateRecoveryInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &drs.UntagResourceInput{}
        output := &drs.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFailbackReplicationConfiguration", func(t *testing.T) {
        input := &drs.UpdateFailbackReplicationConfigurationInput{}
        output := &drs.UpdateFailbackReplicationConfigurationOutput{}

        mockClient.On("UpdateFailbackReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFailbackReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchConfiguration", func(t *testing.T) {
        input := &drs.UpdateLaunchConfigurationInput{}
        output := &drs.UpdateLaunchConfigurationOutput{}

        mockClient.On("UpdateLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchConfigurationTemplate", func(t *testing.T) {
        input := &drs.UpdateLaunchConfigurationTemplateInput{}
        output := &drs.UpdateLaunchConfigurationTemplateOutput{}

        mockClient.On("UpdateLaunchConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationConfiguration", func(t *testing.T) {
        input := &drs.UpdateReplicationConfigurationInput{}
        output := &drs.UpdateReplicationConfigurationOutput{}

        mockClient.On("UpdateReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationConfigurationTemplate", func(t *testing.T) {
        input := &drs.UpdateReplicationConfigurationTemplateInput{}
        output := &drs.UpdateReplicationConfigurationTemplateOutput{}

        mockClient.On("UpdateReplicationConfigurationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationConfigurationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
