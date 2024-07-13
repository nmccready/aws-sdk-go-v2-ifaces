package sfn_test

// tests for the sfn service interface for this ../../../service/sfn/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sfn/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sfn/sfn_iface"
	"github.com/stretchr/testify/assert"
)

func TestSfnServiceCanBeMocked(t *testing.T) {
	var iface sfn_iface.IClient
	iface = &sfn.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sfn.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateActivity", func(t *testing.T) {
        input := &sfn.CreateActivityInput{}
        output := &sfn.CreateActivityOutput{}

        mockClient.On("CreateActivity", ctx, input).Return(output, nil)

        result, err := mockClient.CreateActivity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStateMachine", func(t *testing.T) {
        input := &sfn.CreateStateMachineInput{}
        output := &sfn.CreateStateMachineOutput{}

        mockClient.On("CreateStateMachine", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStateMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStateMachineAlias", func(t *testing.T) {
        input := &sfn.CreateStateMachineAliasInput{}
        output := &sfn.CreateStateMachineAliasOutput{}

        mockClient.On("CreateStateMachineAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStateMachineAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteActivity", func(t *testing.T) {
        input := &sfn.DeleteActivityInput{}
        output := &sfn.DeleteActivityOutput{}

        mockClient.On("DeleteActivity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteActivity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStateMachine", func(t *testing.T) {
        input := &sfn.DeleteStateMachineInput{}
        output := &sfn.DeleteStateMachineOutput{}

        mockClient.On("DeleteStateMachine", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStateMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStateMachineAlias", func(t *testing.T) {
        input := &sfn.DeleteStateMachineAliasInput{}
        output := &sfn.DeleteStateMachineAliasOutput{}

        mockClient.On("DeleteStateMachineAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStateMachineAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStateMachineVersion", func(t *testing.T) {
        input := &sfn.DeleteStateMachineVersionInput{}
        output := &sfn.DeleteStateMachineVersionOutput{}

        mockClient.On("DeleteStateMachineVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStateMachineVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActivity", func(t *testing.T) {
        input := &sfn.DescribeActivityInput{}
        output := &sfn.DescribeActivityOutput{}

        mockClient.On("DescribeActivity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActivity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExecution", func(t *testing.T) {
        input := &sfn.DescribeExecutionInput{}
        output := &sfn.DescribeExecutionOutput{}

        mockClient.On("DescribeExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMapRun", func(t *testing.T) {
        input := &sfn.DescribeMapRunInput{}
        output := &sfn.DescribeMapRunOutput{}

        mockClient.On("DescribeMapRun", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMapRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStateMachine", func(t *testing.T) {
        input := &sfn.DescribeStateMachineInput{}
        output := &sfn.DescribeStateMachineOutput{}

        mockClient.On("DescribeStateMachine", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStateMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStateMachineAlias", func(t *testing.T) {
        input := &sfn.DescribeStateMachineAliasInput{}
        output := &sfn.DescribeStateMachineAliasOutput{}

        mockClient.On("DescribeStateMachineAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStateMachineAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStateMachineForExecution", func(t *testing.T) {
        input := &sfn.DescribeStateMachineForExecutionInput{}
        output := &sfn.DescribeStateMachineForExecutionOutput{}

        mockClient.On("DescribeStateMachineForExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStateMachineForExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetActivityTask", func(t *testing.T) {
        input := &sfn.GetActivityTaskInput{}
        output := &sfn.GetActivityTaskOutput{}

        mockClient.On("GetActivityTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetActivityTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExecutionHistory", func(t *testing.T) {
        input := &sfn.GetExecutionHistoryInput{}
        output := &sfn.GetExecutionHistoryOutput{}

        mockClient.On("GetExecutionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetExecutionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActivities", func(t *testing.T) {
        input := &sfn.ListActivitiesInput{}
        output := &sfn.ListActivitiesOutput{}

        mockClient.On("ListActivities", ctx, input).Return(output, nil)

        result, err := mockClient.ListActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExecutions", func(t *testing.T) {
        input := &sfn.ListExecutionsInput{}
        output := &sfn.ListExecutionsOutput{}

        mockClient.On("ListExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMapRuns", func(t *testing.T) {
        input := &sfn.ListMapRunsInput{}
        output := &sfn.ListMapRunsOutput{}

        mockClient.On("ListMapRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListMapRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStateMachineAliases", func(t *testing.T) {
        input := &sfn.ListStateMachineAliasesInput{}
        output := &sfn.ListStateMachineAliasesOutput{}

        mockClient.On("ListStateMachineAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListStateMachineAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStateMachineVersions", func(t *testing.T) {
        input := &sfn.ListStateMachineVersionsInput{}
        output := &sfn.ListStateMachineVersionsOutput{}

        mockClient.On("ListStateMachineVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListStateMachineVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStateMachines", func(t *testing.T) {
        input := &sfn.ListStateMachinesInput{}
        output := &sfn.ListStateMachinesOutput{}

        mockClient.On("ListStateMachines", ctx, input).Return(output, nil)

        result, err := mockClient.ListStateMachines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &sfn.ListTagsForResourceInput{}
        output := &sfn.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishStateMachineVersion", func(t *testing.T) {
        input := &sfn.PublishStateMachineVersionInput{}
        output := &sfn.PublishStateMachineVersionOutput{}

        mockClient.On("PublishStateMachineVersion", ctx, input).Return(output, nil)

        result, err := mockClient.PublishStateMachineVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRedriveExecution", func(t *testing.T) {
        input := &sfn.RedriveExecutionInput{}
        output := &sfn.RedriveExecutionOutput{}

        mockClient.On("RedriveExecution", ctx, input).Return(output, nil)

        result, err := mockClient.RedriveExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTaskFailure", func(t *testing.T) {
        input := &sfn.SendTaskFailureInput{}
        output := &sfn.SendTaskFailureOutput{}

        mockClient.On("SendTaskFailure", ctx, input).Return(output, nil)

        result, err := mockClient.SendTaskFailure(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTaskHeartbeat", func(t *testing.T) {
        input := &sfn.SendTaskHeartbeatInput{}
        output := &sfn.SendTaskHeartbeatOutput{}

        mockClient.On("SendTaskHeartbeat", ctx, input).Return(output, nil)

        result, err := mockClient.SendTaskHeartbeat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTaskSuccess", func(t *testing.T) {
        input := &sfn.SendTaskSuccessInput{}
        output := &sfn.SendTaskSuccessOutput{}

        mockClient.On("SendTaskSuccess", ctx, input).Return(output, nil)

        result, err := mockClient.SendTaskSuccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExecution", func(t *testing.T) {
        input := &sfn.StartExecutionInput{}
        output := &sfn.StartExecutionOutput{}

        mockClient.On("StartExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSyncExecution", func(t *testing.T) {
        input := &sfn.StartSyncExecutionInput{}
        output := &sfn.StartSyncExecutionOutput{}

        mockClient.On("StartSyncExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartSyncExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopExecution", func(t *testing.T) {
        input := &sfn.StopExecutionInput{}
        output := &sfn.StopExecutionOutput{}

        mockClient.On("StopExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StopExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &sfn.TagResourceInput{}
        output := &sfn.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestState", func(t *testing.T) {
        input := &sfn.TestStateInput{}
        output := &sfn.TestStateOutput{}

        mockClient.On("TestState", ctx, input).Return(output, nil)

        result, err := mockClient.TestState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &sfn.UntagResourceInput{}
        output := &sfn.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMapRun", func(t *testing.T) {
        input := &sfn.UpdateMapRunInput{}
        output := &sfn.UpdateMapRunOutput{}

        mockClient.On("UpdateMapRun", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMapRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStateMachine", func(t *testing.T) {
        input := &sfn.UpdateStateMachineInput{}
        output := &sfn.UpdateStateMachineOutput{}

        mockClient.On("UpdateStateMachine", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStateMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStateMachineAlias", func(t *testing.T) {
        input := &sfn.UpdateStateMachineAliasInput{}
        output := &sfn.UpdateStateMachineAliasOutput{}

        mockClient.On("UpdateStateMachineAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStateMachineAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateStateMachineDefinition", func(t *testing.T) {
        input := &sfn.ValidateStateMachineDefinitionInput{}
        output := &sfn.ValidateStateMachineDefinitionOutput{}

        mockClient.On("ValidateStateMachineDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateStateMachineDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
