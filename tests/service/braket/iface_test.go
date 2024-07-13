package braket_test

// tests for the braket service interface for this ../../../service/braket/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/braket"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/braket/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/braket/braket_iface"
	"github.com/stretchr/testify/assert"
)

func TestBraketServiceCanBeMocked(t *testing.T) {
	var iface braket_iface.IClient
	iface = &braket.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := braket.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &braket.CancelJobInput{}
        output := &braket.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelQuantumTask", func(t *testing.T) {
        input := &braket.CancelQuantumTaskInput{}
        output := &braket.CancelQuantumTaskOutput{}

        mockClient.On("CancelQuantumTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelQuantumTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &braket.CreateJobInput{}
        output := &braket.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQuantumTask", func(t *testing.T) {
        input := &braket.CreateQuantumTaskInput{}
        output := &braket.CreateQuantumTaskOutput{}

        mockClient.On("CreateQuantumTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQuantumTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevice", func(t *testing.T) {
        input := &braket.GetDeviceInput{}
        output := &braket.GetDeviceOutput{}

        mockClient.On("GetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &braket.GetJobInput{}
        output := &braket.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQuantumTask", func(t *testing.T) {
        input := &braket.GetQuantumTaskInput{}
        output := &braket.GetQuantumTaskOutput{}

        mockClient.On("GetQuantumTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetQuantumTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &braket.ListTagsForResourceInput{}
        output := &braket.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchDevices", func(t *testing.T) {
        input := &braket.SearchDevicesInput{}
        output := &braket.SearchDevicesOutput{}

        mockClient.On("SearchDevices", ctx, input).Return(output, nil)

        result, err := mockClient.SearchDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchJobs", func(t *testing.T) {
        input := &braket.SearchJobsInput{}
        output := &braket.SearchJobsOutput{}

        mockClient.On("SearchJobs", ctx, input).Return(output, nil)

        result, err := mockClient.SearchJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchQuantumTasks", func(t *testing.T) {
        input := &braket.SearchQuantumTasksInput{}
        output := &braket.SearchQuantumTasksOutput{}

        mockClient.On("SearchQuantumTasks", ctx, input).Return(output, nil)

        result, err := mockClient.SearchQuantumTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &braket.TagResourceInput{}
        output := &braket.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &braket.UntagResourceInput{}
        output := &braket.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
