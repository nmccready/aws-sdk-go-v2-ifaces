package sms_test

// tests for the sms service interface for this ../../../service/sms/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sms"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sms/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sms/sms_iface"
	"github.com/stretchr/testify/assert"
)

func TestSmsServiceCanBeMocked(t *testing.T) {
	var iface sms_iface.IClient
	iface = &sms.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sms.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &sms.CreateAppInput{}
        output := &sms.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationJob", func(t *testing.T) {
        input := &sms.CreateReplicationJobInput{}
        output := &sms.CreateReplicationJobOutput{}

        mockClient.On("CreateReplicationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &sms.DeleteAppInput{}
        output := &sms.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppLaunchConfiguration", func(t *testing.T) {
        input := &sms.DeleteAppLaunchConfigurationInput{}
        output := &sms.DeleteAppLaunchConfigurationOutput{}

        mockClient.On("DeleteAppLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppReplicationConfiguration", func(t *testing.T) {
        input := &sms.DeleteAppReplicationConfigurationInput{}
        output := &sms.DeleteAppReplicationConfigurationOutput{}

        mockClient.On("DeleteAppReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppValidationConfiguration", func(t *testing.T) {
        input := &sms.DeleteAppValidationConfigurationInput{}
        output := &sms.DeleteAppValidationConfigurationOutput{}

        mockClient.On("DeleteAppValidationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppValidationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationJob", func(t *testing.T) {
        input := &sms.DeleteReplicationJobInput{}
        output := &sms.DeleteReplicationJobOutput{}

        mockClient.On("DeleteReplicationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServerCatalog", func(t *testing.T) {
        input := &sms.DeleteServerCatalogInput{}
        output := &sms.DeleteServerCatalogOutput{}

        mockClient.On("DeleteServerCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServerCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateConnector", func(t *testing.T) {
        input := &sms.DisassociateConnectorInput{}
        output := &sms.DisassociateConnectorOutput{}

        mockClient.On("DisassociateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateChangeSet", func(t *testing.T) {
        input := &sms.GenerateChangeSetInput{}
        output := &sms.GenerateChangeSetOutput{}

        mockClient.On("GenerateChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateTemplate", func(t *testing.T) {
        input := &sms.GenerateTemplateInput{}
        output := &sms.GenerateTemplateOutput{}

        mockClient.On("GenerateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApp", func(t *testing.T) {
        input := &sms.GetAppInput{}
        output := &sms.GetAppOutput{}

        mockClient.On("GetApp", ctx, input).Return(output, nil)

        result, err := mockClient.GetApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppLaunchConfiguration", func(t *testing.T) {
        input := &sms.GetAppLaunchConfigurationInput{}
        output := &sms.GetAppLaunchConfigurationOutput{}

        mockClient.On("GetAppLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppReplicationConfiguration", func(t *testing.T) {
        input := &sms.GetAppReplicationConfigurationInput{}
        output := &sms.GetAppReplicationConfigurationOutput{}

        mockClient.On("GetAppReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppValidationConfiguration", func(t *testing.T) {
        input := &sms.GetAppValidationConfigurationInput{}
        output := &sms.GetAppValidationConfigurationOutput{}

        mockClient.On("GetAppValidationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppValidationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppValidationOutput", func(t *testing.T) {
        input := &sms.GetAppValidationOutputInput{}
        output := &sms.GetAppValidationOutputOutput{}

        mockClient.On("GetAppValidationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppValidationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectors", func(t *testing.T) {
        input := &sms.GetConnectorsInput{}
        output := &sms.GetConnectorsOutput{}

        mockClient.On("GetConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReplicationJobs", func(t *testing.T) {
        input := &sms.GetReplicationJobsInput{}
        output := &sms.GetReplicationJobsOutput{}

        mockClient.On("GetReplicationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetReplicationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReplicationRuns", func(t *testing.T) {
        input := &sms.GetReplicationRunsInput{}
        output := &sms.GetReplicationRunsOutput{}

        mockClient.On("GetReplicationRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetReplicationRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServers", func(t *testing.T) {
        input := &sms.GetServersInput{}
        output := &sms.GetServersOutput{}

        mockClient.On("GetServers", ctx, input).Return(output, nil)

        result, err := mockClient.GetServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportAppCatalog", func(t *testing.T) {
        input := &sms.ImportAppCatalogInput{}
        output := &sms.ImportAppCatalogOutput{}

        mockClient.On("ImportAppCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.ImportAppCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportServerCatalog", func(t *testing.T) {
        input := &sms.ImportServerCatalogInput{}
        output := &sms.ImportServerCatalogOutput{}

        mockClient.On("ImportServerCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.ImportServerCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLaunchApp", func(t *testing.T) {
        input := &sms.LaunchAppInput{}
        output := &sms.LaunchAppOutput{}

        mockClient.On("LaunchApp", ctx, input).Return(output, nil)

        result, err := mockClient.LaunchApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApps", func(t *testing.T) {
        input := &sms.ListAppsInput{}
        output := &sms.ListAppsOutput{}

        mockClient.On("ListApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyAppValidationOutput", func(t *testing.T) {
        input := &sms.NotifyAppValidationOutputInput{}
        output := &sms.NotifyAppValidationOutputOutput{}

        mockClient.On("NotifyAppValidationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyAppValidationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppLaunchConfiguration", func(t *testing.T) {
        input := &sms.PutAppLaunchConfigurationInput{}
        output := &sms.PutAppLaunchConfigurationOutput{}

        mockClient.On("PutAppLaunchConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppLaunchConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppReplicationConfiguration", func(t *testing.T) {
        input := &sms.PutAppReplicationConfigurationInput{}
        output := &sms.PutAppReplicationConfigurationOutput{}

        mockClient.On("PutAppReplicationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppReplicationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppValidationConfiguration", func(t *testing.T) {
        input := &sms.PutAppValidationConfigurationInput{}
        output := &sms.PutAppValidationConfigurationOutput{}

        mockClient.On("PutAppValidationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppValidationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAppReplication", func(t *testing.T) {
        input := &sms.StartAppReplicationInput{}
        output := &sms.StartAppReplicationOutput{}

        mockClient.On("StartAppReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartAppReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartOnDemandAppReplication", func(t *testing.T) {
        input := &sms.StartOnDemandAppReplicationInput{}
        output := &sms.StartOnDemandAppReplicationOutput{}

        mockClient.On("StartOnDemandAppReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartOnDemandAppReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartOnDemandReplicationRun", func(t *testing.T) {
        input := &sms.StartOnDemandReplicationRunInput{}
        output := &sms.StartOnDemandReplicationRunOutput{}

        mockClient.On("StartOnDemandReplicationRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartOnDemandReplicationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAppReplication", func(t *testing.T) {
        input := &sms.StopAppReplicationInput{}
        output := &sms.StopAppReplicationOutput{}

        mockClient.On("StopAppReplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopAppReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateApp", func(t *testing.T) {
        input := &sms.TerminateAppInput{}
        output := &sms.TerminateAppOutput{}

        mockClient.On("TerminateApp", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApp", func(t *testing.T) {
        input := &sms.UpdateAppInput{}
        output := &sms.UpdateAppOutput{}

        mockClient.On("UpdateApp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationJob", func(t *testing.T) {
        input := &sms.UpdateReplicationJobInput{}
        output := &sms.UpdateReplicationJobOutput{}

        mockClient.On("UpdateReplicationJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
