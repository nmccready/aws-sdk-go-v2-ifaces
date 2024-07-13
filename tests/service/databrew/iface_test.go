package databrew_test

// tests for the databrew service interface for this ../../../service/databrew/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/databrew"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/databrew/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/databrew/databrew_iface"
	"github.com/stretchr/testify/assert"
)

func TestDatabrewServiceCanBeMocked(t *testing.T) {
	var iface databrew_iface.IClient
	iface = &databrew.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := databrew.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteRecipeVersion", func(t *testing.T) {
        input := &databrew.BatchDeleteRecipeVersionInput{}
        output := &databrew.BatchDeleteRecipeVersionOutput{}

        mockClient.On("BatchDeleteRecipeVersion", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteRecipeVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &databrew.CreateDatasetInput{}
        output := &databrew.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfileJob", func(t *testing.T) {
        input := &databrew.CreateProfileJobInput{}
        output := &databrew.CreateProfileJobOutput{}

        mockClient.On("CreateProfileJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfileJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &databrew.CreateProjectInput{}
        output := &databrew.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecipe", func(t *testing.T) {
        input := &databrew.CreateRecipeInput{}
        output := &databrew.CreateRecipeOutput{}

        mockClient.On("CreateRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecipeJob", func(t *testing.T) {
        input := &databrew.CreateRecipeJobInput{}
        output := &databrew.CreateRecipeJobOutput{}

        mockClient.On("CreateRecipeJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecipeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleset", func(t *testing.T) {
        input := &databrew.CreateRulesetInput{}
        output := &databrew.CreateRulesetOutput{}

        mockClient.On("CreateRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchedule", func(t *testing.T) {
        input := &databrew.CreateScheduleInput{}
        output := &databrew.CreateScheduleOutput{}

        mockClient.On("CreateSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &databrew.DeleteDatasetInput{}
        output := &databrew.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &databrew.DeleteJobInput{}
        output := &databrew.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &databrew.DeleteProjectInput{}
        output := &databrew.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecipeVersion", func(t *testing.T) {
        input := &databrew.DeleteRecipeVersionInput{}
        output := &databrew.DeleteRecipeVersionOutput{}

        mockClient.On("DeleteRecipeVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecipeVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleset", func(t *testing.T) {
        input := &databrew.DeleteRulesetInput{}
        output := &databrew.DeleteRulesetOutput{}

        mockClient.On("DeleteRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchedule", func(t *testing.T) {
        input := &databrew.DeleteScheduleInput{}
        output := &databrew.DeleteScheduleOutput{}

        mockClient.On("DeleteSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &databrew.DescribeDatasetInput{}
        output := &databrew.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJob", func(t *testing.T) {
        input := &databrew.DescribeJobInput{}
        output := &databrew.DescribeJobOutput{}

        mockClient.On("DescribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobRun", func(t *testing.T) {
        input := &databrew.DescribeJobRunInput{}
        output := &databrew.DescribeJobRunOutput{}

        mockClient.On("DescribeJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &databrew.DescribeProjectInput{}
        output := &databrew.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecipe", func(t *testing.T) {
        input := &databrew.DescribeRecipeInput{}
        output := &databrew.DescribeRecipeOutput{}

        mockClient.On("DescribeRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuleset", func(t *testing.T) {
        input := &databrew.DescribeRulesetInput{}
        output := &databrew.DescribeRulesetOutput{}

        mockClient.On("DescribeRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchedule", func(t *testing.T) {
        input := &databrew.DescribeScheduleInput{}
        output := &databrew.DescribeScheduleOutput{}

        mockClient.On("DescribeSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &databrew.ListDatasetsInput{}
        output := &databrew.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobRuns", func(t *testing.T) {
        input := &databrew.ListJobRunsInput{}
        output := &databrew.ListJobRunsOutput{}

        mockClient.On("ListJobRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &databrew.ListJobsInput{}
        output := &databrew.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &databrew.ListProjectsInput{}
        output := &databrew.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecipeVersions", func(t *testing.T) {
        input := &databrew.ListRecipeVersionsInput{}
        output := &databrew.ListRecipeVersionsOutput{}

        mockClient.On("ListRecipeVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecipeVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecipes", func(t *testing.T) {
        input := &databrew.ListRecipesInput{}
        output := &databrew.ListRecipesOutput{}

        mockClient.On("ListRecipes", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecipes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRulesets", func(t *testing.T) {
        input := &databrew.ListRulesetsInput{}
        output := &databrew.ListRulesetsOutput{}

        mockClient.On("ListRulesets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRulesets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchedules", func(t *testing.T) {
        input := &databrew.ListSchedulesInput{}
        output := &databrew.ListSchedulesOutput{}

        mockClient.On("ListSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &databrew.ListTagsForResourceInput{}
        output := &databrew.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishRecipe", func(t *testing.T) {
        input := &databrew.PublishRecipeInput{}
        output := &databrew.PublishRecipeOutput{}

        mockClient.On("PublishRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.PublishRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendProjectSessionAction", func(t *testing.T) {
        input := &databrew.SendProjectSessionActionInput{}
        output := &databrew.SendProjectSessionActionOutput{}

        mockClient.On("SendProjectSessionAction", ctx, input).Return(output, nil)

        result, err := mockClient.SendProjectSessionAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJobRun", func(t *testing.T) {
        input := &databrew.StartJobRunInput{}
        output := &databrew.StartJobRunOutput{}

        mockClient.On("StartJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartProjectSession", func(t *testing.T) {
        input := &databrew.StartProjectSessionInput{}
        output := &databrew.StartProjectSessionOutput{}

        mockClient.On("StartProjectSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartProjectSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopJobRun", func(t *testing.T) {
        input := &databrew.StopJobRunInput{}
        output := &databrew.StopJobRunOutput{}

        mockClient.On("StopJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &databrew.TagResourceInput{}
        output := &databrew.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &databrew.UntagResourceInput{}
        output := &databrew.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataset", func(t *testing.T) {
        input := &databrew.UpdateDatasetInput{}
        output := &databrew.UpdateDatasetOutput{}

        mockClient.On("UpdateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfileJob", func(t *testing.T) {
        input := &databrew.UpdateProfileJobInput{}
        output := &databrew.UpdateProfileJobOutput{}

        mockClient.On("UpdateProfileJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfileJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &databrew.UpdateProjectInput{}
        output := &databrew.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecipe", func(t *testing.T) {
        input := &databrew.UpdateRecipeInput{}
        output := &databrew.UpdateRecipeOutput{}

        mockClient.On("UpdateRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecipeJob", func(t *testing.T) {
        input := &databrew.UpdateRecipeJobInput{}
        output := &databrew.UpdateRecipeJobOutput{}

        mockClient.On("UpdateRecipeJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecipeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleset", func(t *testing.T) {
        input := &databrew.UpdateRulesetInput{}
        output := &databrew.UpdateRulesetOutput{}

        mockClient.On("UpdateRuleset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchedule", func(t *testing.T) {
        input := &databrew.UpdateScheduleInput{}
        output := &databrew.UpdateScheduleOutput{}

        mockClient.On("UpdateSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
