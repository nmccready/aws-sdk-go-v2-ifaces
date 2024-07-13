package batch_test

// tests for the batch service interface for this ../../../service/batch/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/batch/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/batch/batch_iface"
	"github.com/stretchr/testify/assert"
)

func TestBatchServiceCanBeMocked(t *testing.T) {
	var iface batch_iface.IClient
	iface = &batch.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := batch.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &batch.CancelJobInput{}
        output := &batch.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComputeEnvironment", func(t *testing.T) {
        input := &batch.CreateComputeEnvironmentInput{}
        output := &batch.CreateComputeEnvironmentOutput{}

        mockClient.On("CreateComputeEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComputeEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJobQueue", func(t *testing.T) {
        input := &batch.CreateJobQueueInput{}
        output := &batch.CreateJobQueueOutput{}

        mockClient.On("CreateJobQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJobQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchedulingPolicy", func(t *testing.T) {
        input := &batch.CreateSchedulingPolicyInput{}
        output := &batch.CreateSchedulingPolicyOutput{}

        mockClient.On("CreateSchedulingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchedulingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComputeEnvironment", func(t *testing.T) {
        input := &batch.DeleteComputeEnvironmentInput{}
        output := &batch.DeleteComputeEnvironmentOutput{}

        mockClient.On("DeleteComputeEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComputeEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobQueue", func(t *testing.T) {
        input := &batch.DeleteJobQueueInput{}
        output := &batch.DeleteJobQueueOutput{}

        mockClient.On("DeleteJobQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchedulingPolicy", func(t *testing.T) {
        input := &batch.DeleteSchedulingPolicyInput{}
        output := &batch.DeleteSchedulingPolicyOutput{}

        mockClient.On("DeleteSchedulingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchedulingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterJobDefinition", func(t *testing.T) {
        input := &batch.DeregisterJobDefinitionInput{}
        output := &batch.DeregisterJobDefinitionOutput{}

        mockClient.On("DeregisterJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComputeEnvironments", func(t *testing.T) {
        input := &batch.DescribeComputeEnvironmentsInput{}
        output := &batch.DescribeComputeEnvironmentsOutput{}

        mockClient.On("DescribeComputeEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComputeEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobDefinitions", func(t *testing.T) {
        input := &batch.DescribeJobDefinitionsInput{}
        output := &batch.DescribeJobDefinitionsOutput{}

        mockClient.On("DescribeJobDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobQueues", func(t *testing.T) {
        input := &batch.DescribeJobQueuesInput{}
        output := &batch.DescribeJobQueuesOutput{}

        mockClient.On("DescribeJobQueues", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobs", func(t *testing.T) {
        input := &batch.DescribeJobsInput{}
        output := &batch.DescribeJobsOutput{}

        mockClient.On("DescribeJobs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchedulingPolicies", func(t *testing.T) {
        input := &batch.DescribeSchedulingPoliciesInput{}
        output := &batch.DescribeSchedulingPoliciesOutput{}

        mockClient.On("DescribeSchedulingPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchedulingPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobQueueSnapshot", func(t *testing.T) {
        input := &batch.GetJobQueueSnapshotInput{}
        output := &batch.GetJobQueueSnapshotOutput{}

        mockClient.On("GetJobQueueSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobQueueSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &batch.ListJobsInput{}
        output := &batch.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchedulingPolicies", func(t *testing.T) {
        input := &batch.ListSchedulingPoliciesInput{}
        output := &batch.ListSchedulingPoliciesOutput{}

        mockClient.On("ListSchedulingPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchedulingPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &batch.ListTagsForResourceInput{}
        output := &batch.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterJobDefinition", func(t *testing.T) {
        input := &batch.RegisterJobDefinitionInput{}
        output := &batch.RegisterJobDefinitionOutput{}

        mockClient.On("RegisterJobDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterJobDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitJob", func(t *testing.T) {
        input := &batch.SubmitJobInput{}
        output := &batch.SubmitJobOutput{}

        mockClient.On("SubmitJob", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &batch.TagResourceInput{}
        output := &batch.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateJob", func(t *testing.T) {
        input := &batch.TerminateJobInput{}
        output := &batch.TerminateJobOutput{}

        mockClient.On("TerminateJob", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &batch.UntagResourceInput{}
        output := &batch.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComputeEnvironment", func(t *testing.T) {
        input := &batch.UpdateComputeEnvironmentInput{}
        output := &batch.UpdateComputeEnvironmentOutput{}

        mockClient.On("UpdateComputeEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComputeEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobQueue", func(t *testing.T) {
        input := &batch.UpdateJobQueueInput{}
        output := &batch.UpdateJobQueueOutput{}

        mockClient.On("UpdateJobQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchedulingPolicy", func(t *testing.T) {
        input := &batch.UpdateSchedulingPolicyInput{}
        output := &batch.UpdateSchedulingPolicyOutput{}

        mockClient.On("UpdateSchedulingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchedulingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
