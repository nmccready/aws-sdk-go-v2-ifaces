package robomaker_test

// tests for the robomaker service interface for this ../../../service/robomaker/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/robomaker"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/robomaker/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/robomaker/robomaker_iface"
	"github.com/stretchr/testify/assert"
)

func TestRobomakerServiceCanBeMocked(t *testing.T) {
	var iface robomaker_iface.IClient
	iface = &robomaker.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := robomaker.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteWorlds", func(t *testing.T) {
        input := &robomaker.BatchDeleteWorldsInput{}
        output := &robomaker.BatchDeleteWorldsOutput{}

        mockClient.On("BatchDeleteWorlds", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteWorlds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDescribeSimulationJob", func(t *testing.T) {
        input := &robomaker.BatchDescribeSimulationJobInput{}
        output := &robomaker.BatchDescribeSimulationJobOutput{}

        mockClient.On("BatchDescribeSimulationJob", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDescribeSimulationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDeploymentJob", func(t *testing.T) {
        input := &robomaker.CancelDeploymentJobInput{}
        output := &robomaker.CancelDeploymentJobOutput{}

        mockClient.On("CancelDeploymentJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDeploymentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSimulationJob", func(t *testing.T) {
        input := &robomaker.CancelSimulationJobInput{}
        output := &robomaker.CancelSimulationJobOutput{}

        mockClient.On("CancelSimulationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSimulationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSimulationJobBatch", func(t *testing.T) {
        input := &robomaker.CancelSimulationJobBatchInput{}
        output := &robomaker.CancelSimulationJobBatchOutput{}

        mockClient.On("CancelSimulationJobBatch", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSimulationJobBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelWorldExportJob", func(t *testing.T) {
        input := &robomaker.CancelWorldExportJobInput{}
        output := &robomaker.CancelWorldExportJobOutput{}

        mockClient.On("CancelWorldExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelWorldExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelWorldGenerationJob", func(t *testing.T) {
        input := &robomaker.CancelWorldGenerationJobInput{}
        output := &robomaker.CancelWorldGenerationJobOutput{}

        mockClient.On("CancelWorldGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelWorldGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeploymentJob", func(t *testing.T) {
        input := &robomaker.CreateDeploymentJobInput{}
        output := &robomaker.CreateDeploymentJobOutput{}

        mockClient.On("CreateDeploymentJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeploymentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &robomaker.CreateFleetInput{}
        output := &robomaker.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRobot", func(t *testing.T) {
        input := &robomaker.CreateRobotInput{}
        output := &robomaker.CreateRobotOutput{}

        mockClient.On("CreateRobot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRobot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRobotApplication", func(t *testing.T) {
        input := &robomaker.CreateRobotApplicationInput{}
        output := &robomaker.CreateRobotApplicationOutput{}

        mockClient.On("CreateRobotApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRobotApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRobotApplicationVersion", func(t *testing.T) {
        input := &robomaker.CreateRobotApplicationVersionInput{}
        output := &robomaker.CreateRobotApplicationVersionOutput{}

        mockClient.On("CreateRobotApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRobotApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSimulationApplication", func(t *testing.T) {
        input := &robomaker.CreateSimulationApplicationInput{}
        output := &robomaker.CreateSimulationApplicationOutput{}

        mockClient.On("CreateSimulationApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSimulationApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSimulationApplicationVersion", func(t *testing.T) {
        input := &robomaker.CreateSimulationApplicationVersionInput{}
        output := &robomaker.CreateSimulationApplicationVersionOutput{}

        mockClient.On("CreateSimulationApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSimulationApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSimulationJob", func(t *testing.T) {
        input := &robomaker.CreateSimulationJobInput{}
        output := &robomaker.CreateSimulationJobOutput{}

        mockClient.On("CreateSimulationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSimulationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorldExportJob", func(t *testing.T) {
        input := &robomaker.CreateWorldExportJobInput{}
        output := &robomaker.CreateWorldExportJobOutput{}

        mockClient.On("CreateWorldExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorldExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorldGenerationJob", func(t *testing.T) {
        input := &robomaker.CreateWorldGenerationJobInput{}
        output := &robomaker.CreateWorldGenerationJobOutput{}

        mockClient.On("CreateWorldGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorldGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorldTemplate", func(t *testing.T) {
        input := &robomaker.CreateWorldTemplateInput{}
        output := &robomaker.CreateWorldTemplateOutput{}

        mockClient.On("CreateWorldTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorldTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &robomaker.DeleteFleetInput{}
        output := &robomaker.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRobot", func(t *testing.T) {
        input := &robomaker.DeleteRobotInput{}
        output := &robomaker.DeleteRobotOutput{}

        mockClient.On("DeleteRobot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRobot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRobotApplication", func(t *testing.T) {
        input := &robomaker.DeleteRobotApplicationInput{}
        output := &robomaker.DeleteRobotApplicationOutput{}

        mockClient.On("DeleteRobotApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRobotApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSimulationApplication", func(t *testing.T) {
        input := &robomaker.DeleteSimulationApplicationInput{}
        output := &robomaker.DeleteSimulationApplicationOutput{}

        mockClient.On("DeleteSimulationApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSimulationApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorldTemplate", func(t *testing.T) {
        input := &robomaker.DeleteWorldTemplateInput{}
        output := &robomaker.DeleteWorldTemplateOutput{}

        mockClient.On("DeleteWorldTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorldTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterRobot", func(t *testing.T) {
        input := &robomaker.DeregisterRobotInput{}
        output := &robomaker.DeregisterRobotOutput{}

        mockClient.On("DeregisterRobot", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterRobot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeploymentJob", func(t *testing.T) {
        input := &robomaker.DescribeDeploymentJobInput{}
        output := &robomaker.DescribeDeploymentJobOutput{}

        mockClient.On("DescribeDeploymentJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeploymentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleet", func(t *testing.T) {
        input := &robomaker.DescribeFleetInput{}
        output := &robomaker.DescribeFleetOutput{}

        mockClient.On("DescribeFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRobot", func(t *testing.T) {
        input := &robomaker.DescribeRobotInput{}
        output := &robomaker.DescribeRobotOutput{}

        mockClient.On("DescribeRobot", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRobot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRobotApplication", func(t *testing.T) {
        input := &robomaker.DescribeRobotApplicationInput{}
        output := &robomaker.DescribeRobotApplicationOutput{}

        mockClient.On("DescribeRobotApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRobotApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSimulationApplication", func(t *testing.T) {
        input := &robomaker.DescribeSimulationApplicationInput{}
        output := &robomaker.DescribeSimulationApplicationOutput{}

        mockClient.On("DescribeSimulationApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSimulationApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSimulationJob", func(t *testing.T) {
        input := &robomaker.DescribeSimulationJobInput{}
        output := &robomaker.DescribeSimulationJobOutput{}

        mockClient.On("DescribeSimulationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSimulationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSimulationJobBatch", func(t *testing.T) {
        input := &robomaker.DescribeSimulationJobBatchInput{}
        output := &robomaker.DescribeSimulationJobBatchOutput{}

        mockClient.On("DescribeSimulationJobBatch", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSimulationJobBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorld", func(t *testing.T) {
        input := &robomaker.DescribeWorldInput{}
        output := &robomaker.DescribeWorldOutput{}

        mockClient.On("DescribeWorld", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorld(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorldExportJob", func(t *testing.T) {
        input := &robomaker.DescribeWorldExportJobInput{}
        output := &robomaker.DescribeWorldExportJobOutput{}

        mockClient.On("DescribeWorldExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorldExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorldGenerationJob", func(t *testing.T) {
        input := &robomaker.DescribeWorldGenerationJobInput{}
        output := &robomaker.DescribeWorldGenerationJobOutput{}

        mockClient.On("DescribeWorldGenerationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorldGenerationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorldTemplate", func(t *testing.T) {
        input := &robomaker.DescribeWorldTemplateInput{}
        output := &robomaker.DescribeWorldTemplateOutput{}

        mockClient.On("DescribeWorldTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorldTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorldTemplateBody", func(t *testing.T) {
        input := &robomaker.GetWorldTemplateBodyInput{}
        output := &robomaker.GetWorldTemplateBodyOutput{}

        mockClient.On("GetWorldTemplateBody", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorldTemplateBody(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeploymentJobs", func(t *testing.T) {
        input := &robomaker.ListDeploymentJobsInput{}
        output := &robomaker.ListDeploymentJobsOutput{}

        mockClient.On("ListDeploymentJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeploymentJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &robomaker.ListFleetsInput{}
        output := &robomaker.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRobotApplications", func(t *testing.T) {
        input := &robomaker.ListRobotApplicationsInput{}
        output := &robomaker.ListRobotApplicationsOutput{}

        mockClient.On("ListRobotApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListRobotApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRobots", func(t *testing.T) {
        input := &robomaker.ListRobotsInput{}
        output := &robomaker.ListRobotsOutput{}

        mockClient.On("ListRobots", ctx, input).Return(output, nil)

        result, err := mockClient.ListRobots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSimulationApplications", func(t *testing.T) {
        input := &robomaker.ListSimulationApplicationsInput{}
        output := &robomaker.ListSimulationApplicationsOutput{}

        mockClient.On("ListSimulationApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListSimulationApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSimulationJobBatches", func(t *testing.T) {
        input := &robomaker.ListSimulationJobBatchesInput{}
        output := &robomaker.ListSimulationJobBatchesOutput{}

        mockClient.On("ListSimulationJobBatches", ctx, input).Return(output, nil)

        result, err := mockClient.ListSimulationJobBatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSimulationJobs", func(t *testing.T) {
        input := &robomaker.ListSimulationJobsInput{}
        output := &robomaker.ListSimulationJobsOutput{}

        mockClient.On("ListSimulationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSimulationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &robomaker.ListTagsForResourceInput{}
        output := &robomaker.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorldExportJobs", func(t *testing.T) {
        input := &robomaker.ListWorldExportJobsInput{}
        output := &robomaker.ListWorldExportJobsOutput{}

        mockClient.On("ListWorldExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorldExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorldGenerationJobs", func(t *testing.T) {
        input := &robomaker.ListWorldGenerationJobsInput{}
        output := &robomaker.ListWorldGenerationJobsOutput{}

        mockClient.On("ListWorldGenerationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorldGenerationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorldTemplates", func(t *testing.T) {
        input := &robomaker.ListWorldTemplatesInput{}
        output := &robomaker.ListWorldTemplatesOutput{}

        mockClient.On("ListWorldTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorldTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorlds", func(t *testing.T) {
        input := &robomaker.ListWorldsInput{}
        output := &robomaker.ListWorldsOutput{}

        mockClient.On("ListWorlds", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorlds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterRobot", func(t *testing.T) {
        input := &robomaker.RegisterRobotInput{}
        output := &robomaker.RegisterRobotOutput{}

        mockClient.On("RegisterRobot", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterRobot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestartSimulationJob", func(t *testing.T) {
        input := &robomaker.RestartSimulationJobInput{}
        output := &robomaker.RestartSimulationJobOutput{}

        mockClient.On("RestartSimulationJob", ctx, input).Return(output, nil)

        result, err := mockClient.RestartSimulationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSimulationJobBatch", func(t *testing.T) {
        input := &robomaker.StartSimulationJobBatchInput{}
        output := &robomaker.StartSimulationJobBatchOutput{}

        mockClient.On("StartSimulationJobBatch", ctx, input).Return(output, nil)

        result, err := mockClient.StartSimulationJobBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSyncDeploymentJob", func(t *testing.T) {
        input := &robomaker.SyncDeploymentJobInput{}
        output := &robomaker.SyncDeploymentJobOutput{}

        mockClient.On("SyncDeploymentJob", ctx, input).Return(output, nil)

        result, err := mockClient.SyncDeploymentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &robomaker.TagResourceInput{}
        output := &robomaker.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &robomaker.UntagResourceInput{}
        output := &robomaker.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRobotApplication", func(t *testing.T) {
        input := &robomaker.UpdateRobotApplicationInput{}
        output := &robomaker.UpdateRobotApplicationOutput{}

        mockClient.On("UpdateRobotApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRobotApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSimulationApplication", func(t *testing.T) {
        input := &robomaker.UpdateSimulationApplicationInput{}
        output := &robomaker.UpdateSimulationApplicationOutput{}

        mockClient.On("UpdateSimulationApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSimulationApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorldTemplate", func(t *testing.T) {
        input := &robomaker.UpdateWorldTemplateInput{}
        output := &robomaker.UpdateWorldTemplateOutput{}

        mockClient.On("UpdateWorldTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorldTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
