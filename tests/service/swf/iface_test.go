package swf_test

// tests for the swf service interface for this ../../../service/swf/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/swf/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/swf/swf_iface"
	"github.com/stretchr/testify/assert"
)

func TestSwfServiceCanBeMocked(t *testing.T) {
	var iface swf_iface.IClient
	iface = &swf.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := swf.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCountClosedWorkflowExecutions", func(t *testing.T) {
        input := &swf.CountClosedWorkflowExecutionsInput{}
        output := &swf.CountClosedWorkflowExecutionsOutput{}

        mockClient.On("CountClosedWorkflowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.CountClosedWorkflowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCountOpenWorkflowExecutions", func(t *testing.T) {
        input := &swf.CountOpenWorkflowExecutionsInput{}
        output := &swf.CountOpenWorkflowExecutionsOutput{}

        mockClient.On("CountOpenWorkflowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.CountOpenWorkflowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCountPendingActivityTasks", func(t *testing.T) {
        input := &swf.CountPendingActivityTasksInput{}
        output := &swf.CountPendingActivityTasksOutput{}

        mockClient.On("CountPendingActivityTasks", ctx, input).Return(output, nil)

        result, err := mockClient.CountPendingActivityTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCountPendingDecisionTasks", func(t *testing.T) {
        input := &swf.CountPendingDecisionTasksInput{}
        output := &swf.CountPendingDecisionTasksOutput{}

        mockClient.On("CountPendingDecisionTasks", ctx, input).Return(output, nil)

        result, err := mockClient.CountPendingDecisionTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteActivityType", func(t *testing.T) {
        input := &swf.DeleteActivityTypeInput{}
        output := &swf.DeleteActivityTypeOutput{}

        mockClient.On("DeleteActivityType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteActivityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflowType", func(t *testing.T) {
        input := &swf.DeleteWorkflowTypeInput{}
        output := &swf.DeleteWorkflowTypeOutput{}

        mockClient.On("DeleteWorkflowType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflowType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateActivityType", func(t *testing.T) {
        input := &swf.DeprecateActivityTypeInput{}
        output := &swf.DeprecateActivityTypeOutput{}

        mockClient.On("DeprecateActivityType", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateActivityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateDomain", func(t *testing.T) {
        input := &swf.DeprecateDomainInput{}
        output := &swf.DeprecateDomainOutput{}

        mockClient.On("DeprecateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateWorkflowType", func(t *testing.T) {
        input := &swf.DeprecateWorkflowTypeInput{}
        output := &swf.DeprecateWorkflowTypeOutput{}

        mockClient.On("DeprecateWorkflowType", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateWorkflowType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActivityType", func(t *testing.T) {
        input := &swf.DescribeActivityTypeInput{}
        output := &swf.DescribeActivityTypeOutput{}

        mockClient.On("DescribeActivityType", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActivityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &swf.DescribeDomainInput{}
        output := &swf.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkflowExecution", func(t *testing.T) {
        input := &swf.DescribeWorkflowExecutionInput{}
        output := &swf.DescribeWorkflowExecutionOutput{}

        mockClient.On("DescribeWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkflowType", func(t *testing.T) {
        input := &swf.DescribeWorkflowTypeInput{}
        output := &swf.DescribeWorkflowTypeOutput{}

        mockClient.On("DescribeWorkflowType", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkflowType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowExecutionHistory", func(t *testing.T) {
        input := &swf.GetWorkflowExecutionHistoryInput{}
        output := &swf.GetWorkflowExecutionHistoryOutput{}

        mockClient.On("GetWorkflowExecutionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowExecutionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActivityTypes", func(t *testing.T) {
        input := &swf.ListActivityTypesInput{}
        output := &swf.ListActivityTypesOutput{}

        mockClient.On("ListActivityTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListActivityTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClosedWorkflowExecutions", func(t *testing.T) {
        input := &swf.ListClosedWorkflowExecutionsInput{}
        output := &swf.ListClosedWorkflowExecutionsOutput{}

        mockClient.On("ListClosedWorkflowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListClosedWorkflowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &swf.ListDomainsInput{}
        output := &swf.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOpenWorkflowExecutions", func(t *testing.T) {
        input := &swf.ListOpenWorkflowExecutionsInput{}
        output := &swf.ListOpenWorkflowExecutionsOutput{}

        mockClient.On("ListOpenWorkflowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListOpenWorkflowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &swf.ListTagsForResourceInput{}
        output := &swf.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflowTypes", func(t *testing.T) {
        input := &swf.ListWorkflowTypesInput{}
        output := &swf.ListWorkflowTypesOutput{}

        mockClient.On("ListWorkflowTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflowTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPollForActivityTask", func(t *testing.T) {
        input := &swf.PollForActivityTaskInput{}
        output := &swf.PollForActivityTaskOutput{}

        mockClient.On("PollForActivityTask", ctx, input).Return(output, nil)

        result, err := mockClient.PollForActivityTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPollForDecisionTask", func(t *testing.T) {
        input := &swf.PollForDecisionTaskInput{}
        output := &swf.PollForDecisionTaskOutput{}

        mockClient.On("PollForDecisionTask", ctx, input).Return(output, nil)

        result, err := mockClient.PollForDecisionTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecordActivityTaskHeartbeat", func(t *testing.T) {
        input := &swf.RecordActivityTaskHeartbeatInput{}
        output := &swf.RecordActivityTaskHeartbeatOutput{}

        mockClient.On("RecordActivityTaskHeartbeat", ctx, input).Return(output, nil)

        result, err := mockClient.RecordActivityTaskHeartbeat(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterActivityType", func(t *testing.T) {
        input := &swf.RegisterActivityTypeInput{}
        output := &swf.RegisterActivityTypeOutput{}

        mockClient.On("RegisterActivityType", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterActivityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDomain", func(t *testing.T) {
        input := &swf.RegisterDomainInput{}
        output := &swf.RegisterDomainOutput{}

        mockClient.On("RegisterDomain", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterWorkflowType", func(t *testing.T) {
        input := &swf.RegisterWorkflowTypeInput{}
        output := &swf.RegisterWorkflowTypeOutput{}

        mockClient.On("RegisterWorkflowType", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterWorkflowType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestCancelWorkflowExecution", func(t *testing.T) {
        input := &swf.RequestCancelWorkflowExecutionInput{}
        output := &swf.RequestCancelWorkflowExecutionOutput{}

        mockClient.On("RequestCancelWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.RequestCancelWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRespondActivityTaskCanceled", func(t *testing.T) {
        input := &swf.RespondActivityTaskCanceledInput{}
        output := &swf.RespondActivityTaskCanceledOutput{}

        mockClient.On("RespondActivityTaskCanceled", ctx, input).Return(output, nil)

        result, err := mockClient.RespondActivityTaskCanceled(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRespondActivityTaskCompleted", func(t *testing.T) {
        input := &swf.RespondActivityTaskCompletedInput{}
        output := &swf.RespondActivityTaskCompletedOutput{}

        mockClient.On("RespondActivityTaskCompleted", ctx, input).Return(output, nil)

        result, err := mockClient.RespondActivityTaskCompleted(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRespondActivityTaskFailed", func(t *testing.T) {
        input := &swf.RespondActivityTaskFailedInput{}
        output := &swf.RespondActivityTaskFailedOutput{}

        mockClient.On("RespondActivityTaskFailed", ctx, input).Return(output, nil)

        result, err := mockClient.RespondActivityTaskFailed(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRespondDecisionTaskCompleted", func(t *testing.T) {
        input := &swf.RespondDecisionTaskCompletedInput{}
        output := &swf.RespondDecisionTaskCompletedOutput{}

        mockClient.On("RespondDecisionTaskCompleted", ctx, input).Return(output, nil)

        result, err := mockClient.RespondDecisionTaskCompleted(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignalWorkflowExecution", func(t *testing.T) {
        input := &swf.SignalWorkflowExecutionInput{}
        output := &swf.SignalWorkflowExecutionOutput{}

        mockClient.On("SignalWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.SignalWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWorkflowExecution", func(t *testing.T) {
        input := &swf.StartWorkflowExecutionInput{}
        output := &swf.StartWorkflowExecutionOutput{}

        mockClient.On("StartWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &swf.TagResourceInput{}
        output := &swf.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateWorkflowExecution", func(t *testing.T) {
        input := &swf.TerminateWorkflowExecutionInput{}
        output := &swf.TerminateWorkflowExecutionOutput{}

        mockClient.On("TerminateWorkflowExecution", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateWorkflowExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUndeprecateActivityType", func(t *testing.T) {
        input := &swf.UndeprecateActivityTypeInput{}
        output := &swf.UndeprecateActivityTypeOutput{}

        mockClient.On("UndeprecateActivityType", ctx, input).Return(output, nil)

        result, err := mockClient.UndeprecateActivityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUndeprecateDomain", func(t *testing.T) {
        input := &swf.UndeprecateDomainInput{}
        output := &swf.UndeprecateDomainOutput{}

        mockClient.On("UndeprecateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UndeprecateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUndeprecateWorkflowType", func(t *testing.T) {
        input := &swf.UndeprecateWorkflowTypeInput{}
        output := &swf.UndeprecateWorkflowTypeOutput{}

        mockClient.On("UndeprecateWorkflowType", ctx, input).Return(output, nil)

        result, err := mockClient.UndeprecateWorkflowType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &swf.UntagResourceInput{}
        output := &swf.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
