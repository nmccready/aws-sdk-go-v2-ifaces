package iotjobsdataplane_test

// tests for the iotjobsdataplane service interface for this ../../../service/iotjobsdataplane/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotjobsdataplane/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotjobsdataplane/iotjobsdataplane_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotjobsdataplaneServiceCanBeMocked(t *testing.T) {
	var iface iotjobsdataplane_iface.IClient
	iface = &iotjobsdataplane.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotjobsdataplane.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobExecution", func(t *testing.T) {
        input := &iotjobsdataplane.DescribeJobExecutionInput{}
        output := &iotjobsdataplane.DescribeJobExecutionOutput{}

        mockClient.On("DescribeJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPendingJobExecutions", func(t *testing.T) {
        input := &iotjobsdataplane.GetPendingJobExecutionsInput{}
        output := &iotjobsdataplane.GetPendingJobExecutionsOutput{}

        mockClient.On("GetPendingJobExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.GetPendingJobExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartNextPendingJobExecution", func(t *testing.T) {
        input := &iotjobsdataplane.StartNextPendingJobExecutionInput{}
        output := &iotjobsdataplane.StartNextPendingJobExecutionOutput{}

        mockClient.On("StartNextPendingJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.StartNextPendingJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobExecution", func(t *testing.T) {
        input := &iotjobsdataplane.UpdateJobExecutionInput{}
        output := &iotjobsdataplane.UpdateJobExecutionOutput{}

        mockClient.On("UpdateJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
