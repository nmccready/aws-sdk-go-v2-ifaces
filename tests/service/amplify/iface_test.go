package amplify_test

// tests for the amplify service interface for this ../../../service/amplify/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplify/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplify/amplify_iface"
	"github.com/stretchr/testify/assert"
)

func TestAmplifyServiceCanBeMocked(t *testing.T) {
	var iface amplify_iface.IClient
	iface = &amplify.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := amplify.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &amplify.CreateAppInput{}
        output := &amplify.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackendEnvironment", func(t *testing.T) {
        input := &amplify.CreateBackendEnvironmentInput{}
        output := &amplify.CreateBackendEnvironmentOutput{}

        mockClient.On("CreateBackendEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackendEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBranch", func(t *testing.T) {
        input := &amplify.CreateBranchInput{}
        output := &amplify.CreateBranchOutput{}

        mockClient.On("CreateBranch", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &amplify.CreateDeploymentInput{}
        output := &amplify.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainAssociation", func(t *testing.T) {
        input := &amplify.CreateDomainAssociationInput{}
        output := &amplify.CreateDomainAssociationOutput{}

        mockClient.On("CreateDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebhook", func(t *testing.T) {
        input := &amplify.CreateWebhookInput{}
        output := &amplify.CreateWebhookOutput{}

        mockClient.On("CreateWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &amplify.DeleteAppInput{}
        output := &amplify.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackendEnvironment", func(t *testing.T) {
        input := &amplify.DeleteBackendEnvironmentInput{}
        output := &amplify.DeleteBackendEnvironmentOutput{}

        mockClient.On("DeleteBackendEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackendEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBranch", func(t *testing.T) {
        input := &amplify.DeleteBranchInput{}
        output := &amplify.DeleteBranchOutput{}

        mockClient.On("DeleteBranch", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainAssociation", func(t *testing.T) {
        input := &amplify.DeleteDomainAssociationInput{}
        output := &amplify.DeleteDomainAssociationOutput{}

        mockClient.On("DeleteDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &amplify.DeleteJobInput{}
        output := &amplify.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebhook", func(t *testing.T) {
        input := &amplify.DeleteWebhookInput{}
        output := &amplify.DeleteWebhookOutput{}

        mockClient.On("DeleteWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateAccessLogs", func(t *testing.T) {
        input := &amplify.GenerateAccessLogsInput{}
        output := &amplify.GenerateAccessLogsOutput{}

        mockClient.On("GenerateAccessLogs", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateAccessLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApp", func(t *testing.T) {
        input := &amplify.GetAppInput{}
        output := &amplify.GetAppOutput{}

        mockClient.On("GetApp", ctx, input).Return(output, nil)

        result, err := mockClient.GetApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArtifactUrl", func(t *testing.T) {
        input := &amplify.GetArtifactUrlInput{}
        output := &amplify.GetArtifactUrlOutput{}

        mockClient.On("GetArtifactUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetArtifactUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendEnvironment", func(t *testing.T) {
        input := &amplify.GetBackendEnvironmentInput{}
        output := &amplify.GetBackendEnvironmentOutput{}

        mockClient.On("GetBackendEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBranch", func(t *testing.T) {
        input := &amplify.GetBranchInput{}
        output := &amplify.GetBranchOutput{}

        mockClient.On("GetBranch", ctx, input).Return(output, nil)

        result, err := mockClient.GetBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainAssociation", func(t *testing.T) {
        input := &amplify.GetDomainAssociationInput{}
        output := &amplify.GetDomainAssociationOutput{}

        mockClient.On("GetDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &amplify.GetJobInput{}
        output := &amplify.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebhook", func(t *testing.T) {
        input := &amplify.GetWebhookInput{}
        output := &amplify.GetWebhookOutput{}

        mockClient.On("GetWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApps", func(t *testing.T) {
        input := &amplify.ListAppsInput{}
        output := &amplify.ListAppsOutput{}

        mockClient.On("ListApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArtifacts", func(t *testing.T) {
        input := &amplify.ListArtifactsInput{}
        output := &amplify.ListArtifactsOutput{}

        mockClient.On("ListArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackendEnvironments", func(t *testing.T) {
        input := &amplify.ListBackendEnvironmentsInput{}
        output := &amplify.ListBackendEnvironmentsOutput{}

        mockClient.On("ListBackendEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackendEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBranches", func(t *testing.T) {
        input := &amplify.ListBranchesInput{}
        output := &amplify.ListBranchesOutput{}

        mockClient.On("ListBranches", ctx, input).Return(output, nil)

        result, err := mockClient.ListBranches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainAssociations", func(t *testing.T) {
        input := &amplify.ListDomainAssociationsInput{}
        output := &amplify.ListDomainAssociationsOutput{}

        mockClient.On("ListDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &amplify.ListJobsInput{}
        output := &amplify.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &amplify.ListTagsForResourceInput{}
        output := &amplify.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebhooks", func(t *testing.T) {
        input := &amplify.ListWebhooksInput{}
        output := &amplify.ListWebhooksOutput{}

        mockClient.On("ListWebhooks", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebhooks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeployment", func(t *testing.T) {
        input := &amplify.StartDeploymentInput{}
        output := &amplify.StartDeploymentOutput{}

        mockClient.On("StartDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJob", func(t *testing.T) {
        input := &amplify.StartJobInput{}
        output := &amplify.StartJobOutput{}

        mockClient.On("StartJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopJob", func(t *testing.T) {
        input := &amplify.StopJobInput{}
        output := &amplify.StopJobOutput{}

        mockClient.On("StopJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &amplify.TagResourceInput{}
        output := &amplify.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &amplify.UntagResourceInput{}
        output := &amplify.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApp", func(t *testing.T) {
        input := &amplify.UpdateAppInput{}
        output := &amplify.UpdateAppOutput{}

        mockClient.On("UpdateApp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBranch", func(t *testing.T) {
        input := &amplify.UpdateBranchInput{}
        output := &amplify.UpdateBranchOutput{}

        mockClient.On("UpdateBranch", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainAssociation", func(t *testing.T) {
        input := &amplify.UpdateDomainAssociationInput{}
        output := &amplify.UpdateDomainAssociationOutput{}

        mockClient.On("UpdateDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebhook", func(t *testing.T) {
        input := &amplify.UpdateWebhookInput{}
        output := &amplify.UpdateWebhookOutput{}

        mockClient.On("UpdateWebhook", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebhook(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
