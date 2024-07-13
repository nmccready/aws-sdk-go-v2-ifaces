package emrserverless_test

// tests for the emrserverless service interface for this ../../../service/emrserverless/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emrserverless"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emrserverless/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/emrserverless/emrserverless_iface"
	"github.com/stretchr/testify/assert"
)

func TestEmrserverlessServiceCanBeMocked(t *testing.T) {
	var iface emrserverless_iface.IClient
	iface = &emrserverless.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := emrserverless.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJobRun", func(t *testing.T) {
        input := &emrserverless.CancelJobRunInput{}
        output := &emrserverless.CancelJobRunOutput{}

        mockClient.On("CancelJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &emrserverless.CreateApplicationInput{}
        output := &emrserverless.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &emrserverless.DeleteApplicationInput{}
        output := &emrserverless.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &emrserverless.GetApplicationInput{}
        output := &emrserverless.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDashboardForJobRun", func(t *testing.T) {
        input := &emrserverless.GetDashboardForJobRunInput{}
        output := &emrserverless.GetDashboardForJobRunOutput{}

        mockClient.On("GetDashboardForJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetDashboardForJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobRun", func(t *testing.T) {
        input := &emrserverless.GetJobRunInput{}
        output := &emrserverless.GetJobRunOutput{}

        mockClient.On("GetJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &emrserverless.ListApplicationsInput{}
        output := &emrserverless.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobRunAttempts", func(t *testing.T) {
        input := &emrserverless.ListJobRunAttemptsInput{}
        output := &emrserverless.ListJobRunAttemptsOutput{}

        mockClient.On("ListJobRunAttempts", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobRunAttempts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobRuns", func(t *testing.T) {
        input := &emrserverless.ListJobRunsInput{}
        output := &emrserverless.ListJobRunsOutput{}

        mockClient.On("ListJobRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &emrserverless.ListTagsForResourceInput{}
        output := &emrserverless.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplication", func(t *testing.T) {
        input := &emrserverless.StartApplicationInput{}
        output := &emrserverless.StartApplicationOutput{}

        mockClient.On("StartApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartJobRun", func(t *testing.T) {
        input := &emrserverless.StartJobRunInput{}
        output := &emrserverless.StartJobRunOutput{}

        mockClient.On("StartJobRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartJobRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApplication", func(t *testing.T) {
        input := &emrserverless.StopApplicationInput{}
        output := &emrserverless.StopApplicationOutput{}

        mockClient.On("StopApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &emrserverless.TagResourceInput{}
        output := &emrserverless.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &emrserverless.UntagResourceInput{}
        output := &emrserverless.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &emrserverless.UpdateApplicationInput{}
        output := &emrserverless.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
