package codedeploy_test

// tests for the codedeploy service interface for this ../../../service/codedeploy/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codedeploy/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codedeploy/codedeploy_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodedeployServiceCanBeMocked(t *testing.T) {
	var iface codedeploy_iface.IClient
	iface = &codedeploy.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codedeploy.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToOnPremisesInstances", func(t *testing.T) {
        input := &codedeploy.AddTagsToOnPremisesInstancesInput{}
        output := &codedeploy.AddTagsToOnPremisesInstancesOutput{}

        mockClient.On("AddTagsToOnPremisesInstances", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToOnPremisesInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetApplicationRevisions", func(t *testing.T) {
        input := &codedeploy.BatchGetApplicationRevisionsInput{}
        output := &codedeploy.BatchGetApplicationRevisionsOutput{}

        mockClient.On("BatchGetApplicationRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetApplicationRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetApplications", func(t *testing.T) {
        input := &codedeploy.BatchGetApplicationsInput{}
        output := &codedeploy.BatchGetApplicationsOutput{}

        mockClient.On("BatchGetApplications", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDeploymentGroups", func(t *testing.T) {
        input := &codedeploy.BatchGetDeploymentGroupsInput{}
        output := &codedeploy.BatchGetDeploymentGroupsOutput{}

        mockClient.On("BatchGetDeploymentGroups", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDeploymentGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDeploymentInstances", func(t *testing.T) {
        input := &codedeploy.BatchGetDeploymentInstancesInput{}
        output := &codedeploy.BatchGetDeploymentInstancesOutput{}

        mockClient.On("BatchGetDeploymentInstances", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDeploymentInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDeploymentTargets", func(t *testing.T) {
        input := &codedeploy.BatchGetDeploymentTargetsInput{}
        output := &codedeploy.BatchGetDeploymentTargetsOutput{}

        mockClient.On("BatchGetDeploymentTargets", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDeploymentTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDeployments", func(t *testing.T) {
        input := &codedeploy.BatchGetDeploymentsInput{}
        output := &codedeploy.BatchGetDeploymentsOutput{}

        mockClient.On("BatchGetDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetOnPremisesInstances", func(t *testing.T) {
        input := &codedeploy.BatchGetOnPremisesInstancesInput{}
        output := &codedeploy.BatchGetOnPremisesInstancesOutput{}

        mockClient.On("BatchGetOnPremisesInstances", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetOnPremisesInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestContinueDeployment", func(t *testing.T) {
        input := &codedeploy.ContinueDeploymentInput{}
        output := &codedeploy.ContinueDeploymentOutput{}

        mockClient.On("ContinueDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.ContinueDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &codedeploy.CreateApplicationInput{}
        output := &codedeploy.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &codedeploy.CreateDeploymentInput{}
        output := &codedeploy.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeploymentConfig", func(t *testing.T) {
        input := &codedeploy.CreateDeploymentConfigInput{}
        output := &codedeploy.CreateDeploymentConfigOutput{}

        mockClient.On("CreateDeploymentConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeploymentConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeploymentGroup", func(t *testing.T) {
        input := &codedeploy.CreateDeploymentGroupInput{}
        output := &codedeploy.CreateDeploymentGroupOutput{}

        mockClient.On("CreateDeploymentGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeploymentGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &codedeploy.DeleteApplicationInput{}
        output := &codedeploy.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeploymentConfig", func(t *testing.T) {
        input := &codedeploy.DeleteDeploymentConfigInput{}
        output := &codedeploy.DeleteDeploymentConfigOutput{}

        mockClient.On("DeleteDeploymentConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeploymentConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeploymentGroup", func(t *testing.T) {
        input := &codedeploy.DeleteDeploymentGroupInput{}
        output := &codedeploy.DeleteDeploymentGroupOutput{}

        mockClient.On("DeleteDeploymentGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeploymentGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGitHubAccountToken", func(t *testing.T) {
        input := &codedeploy.DeleteGitHubAccountTokenInput{}
        output := &codedeploy.DeleteGitHubAccountTokenOutput{}

        mockClient.On("DeleteGitHubAccountToken", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGitHubAccountToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcesByExternalId", func(t *testing.T) {
        input := &codedeploy.DeleteResourcesByExternalIdInput{}
        output := &codedeploy.DeleteResourcesByExternalIdOutput{}

        mockClient.On("DeleteResourcesByExternalId", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcesByExternalId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterOnPremisesInstance", func(t *testing.T) {
        input := &codedeploy.DeregisterOnPremisesInstanceInput{}
        output := &codedeploy.DeregisterOnPremisesInstanceOutput{}

        mockClient.On("DeregisterOnPremisesInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterOnPremisesInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &codedeploy.GetApplicationInput{}
        output := &codedeploy.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationRevision", func(t *testing.T) {
        input := &codedeploy.GetApplicationRevisionInput{}
        output := &codedeploy.GetApplicationRevisionOutput{}

        mockClient.On("GetApplicationRevision", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &codedeploy.GetDeploymentInput{}
        output := &codedeploy.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentConfig", func(t *testing.T) {
        input := &codedeploy.GetDeploymentConfigInput{}
        output := &codedeploy.GetDeploymentConfigOutput{}

        mockClient.On("GetDeploymentConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentGroup", func(t *testing.T) {
        input := &codedeploy.GetDeploymentGroupInput{}
        output := &codedeploy.GetDeploymentGroupOutput{}

        mockClient.On("GetDeploymentGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentInstance", func(t *testing.T) {
        input := &codedeploy.GetDeploymentInstanceInput{}
        output := &codedeploy.GetDeploymentInstanceOutput{}

        mockClient.On("GetDeploymentInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeploymentTarget", func(t *testing.T) {
        input := &codedeploy.GetDeploymentTargetInput{}
        output := &codedeploy.GetDeploymentTargetOutput{}

        mockClient.On("GetDeploymentTarget", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeploymentTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOnPremisesInstance", func(t *testing.T) {
        input := &codedeploy.GetOnPremisesInstanceInput{}
        output := &codedeploy.GetOnPremisesInstanceOutput{}

        mockClient.On("GetOnPremisesInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetOnPremisesInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationRevisions", func(t *testing.T) {
        input := &codedeploy.ListApplicationRevisionsInput{}
        output := &codedeploy.ListApplicationRevisionsOutput{}

        mockClient.On("ListApplicationRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &codedeploy.ListApplicationsInput{}
        output := &codedeploy.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentConfigs", func(t *testing.T) {
        input := &codedeploy.ListDeploymentConfigsInput{}
        output := &codedeploy.ListDeploymentConfigsOutput{}

        mockClient.On("ListDeploymentConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentGroups", func(t *testing.T) {
        input := &codedeploy.ListDeploymentGroupsInput{}
        output := &codedeploy.ListDeploymentGroupsOutput{}

        mockClient.On("ListDeploymentGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentInstances", func(t *testing.T) {
        input := &codedeploy.ListDeploymentInstancesInput{}
        output := &codedeploy.ListDeploymentInstancesOutput{}

        mockClient.On("ListDeploymentInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentTargets", func(t *testing.T) {
        input := &codedeploy.ListDeploymentTargetsInput{}
        output := &codedeploy.ListDeploymentTargetsOutput{}

        mockClient.On("ListDeploymentTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &codedeploy.ListDeploymentsInput{}
        output := &codedeploy.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGitHubAccountTokenNames", func(t *testing.T) {
        input := &codedeploy.ListGitHubAccountTokenNamesInput{}
        output := &codedeploy.ListGitHubAccountTokenNamesOutput{}

        mockClient.On("ListGitHubAccountTokenNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListGitHubAccountTokenNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOnPremisesInstances", func(t *testing.T) {
        input := &codedeploy.ListOnPremisesInstancesInput{}
        output := &codedeploy.ListOnPremisesInstancesOutput{}

        mockClient.On("ListOnPremisesInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListOnPremisesInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codedeploy.ListTagsForResourceInput{}
        output := &codedeploy.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLifecycleEventHookExecutionStatus", func(t *testing.T) {
        input := &codedeploy.PutLifecycleEventHookExecutionStatusInput{}
        output := &codedeploy.PutLifecycleEventHookExecutionStatusOutput{}

        mockClient.On("PutLifecycleEventHookExecutionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.PutLifecycleEventHookExecutionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterApplicationRevision", func(t *testing.T) {
        input := &codedeploy.RegisterApplicationRevisionInput{}
        output := &codedeploy.RegisterApplicationRevisionOutput{}

        mockClient.On("RegisterApplicationRevision", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterApplicationRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterOnPremisesInstance", func(t *testing.T) {
        input := &codedeploy.RegisterOnPremisesInstanceInput{}
        output := &codedeploy.RegisterOnPremisesInstanceOutput{}

        mockClient.On("RegisterOnPremisesInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterOnPremisesInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromOnPremisesInstances", func(t *testing.T) {
        input := &codedeploy.RemoveTagsFromOnPremisesInstancesInput{}
        output := &codedeploy.RemoveTagsFromOnPremisesInstancesOutput{}

        mockClient.On("RemoveTagsFromOnPremisesInstances", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromOnPremisesInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSkipWaitTimeForInstanceTermination", func(t *testing.T) {
        input := &codedeploy.SkipWaitTimeForInstanceTerminationInput{}
        output := &codedeploy.SkipWaitTimeForInstanceTerminationOutput{}

        mockClient.On("SkipWaitTimeForInstanceTermination", ctx, input).Return(output, nil)

        result, err := mockClient.SkipWaitTimeForInstanceTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDeployment", func(t *testing.T) {
        input := &codedeploy.StopDeploymentInput{}
        output := &codedeploy.StopDeploymentOutput{}

        mockClient.On("StopDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StopDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codedeploy.TagResourceInput{}
        output := &codedeploy.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codedeploy.UntagResourceInput{}
        output := &codedeploy.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &codedeploy.UpdateApplicationInput{}
        output := &codedeploy.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeploymentGroup", func(t *testing.T) {
        input := &codedeploy.UpdateDeploymentGroupInput{}
        output := &codedeploy.UpdateDeploymentGroupOutput{}

        mockClient.On("UpdateDeploymentGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeploymentGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
