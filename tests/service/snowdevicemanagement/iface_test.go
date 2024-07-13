package snowdevicemanagement_test

// tests for the snowdevicemanagement service interface for this ../../../service/snowdevicemanagement/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/snowdevicemanagement/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/snowdevicemanagement/snowdevicemanagement_iface"
	"github.com/stretchr/testify/assert"
)

func TestSnowdevicemanagementServiceCanBeMocked(t *testing.T) {
	var iface snowdevicemanagement_iface.IClient
	iface = &snowdevicemanagement.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := snowdevicemanagement.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelTask", func(t *testing.T) {
        input := &snowdevicemanagement.CancelTaskInput{}
        output := &snowdevicemanagement.CancelTaskOutput{}

        mockClient.On("CancelTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTask", func(t *testing.T) {
        input := &snowdevicemanagement.CreateTaskInput{}
        output := &snowdevicemanagement.CreateTaskOutput{}

        mockClient.On("CreateTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDevice", func(t *testing.T) {
        input := &snowdevicemanagement.DescribeDeviceInput{}
        output := &snowdevicemanagement.DescribeDeviceOutput{}

        mockClient.On("DescribeDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeviceEc2Instances", func(t *testing.T) {
        input := &snowdevicemanagement.DescribeDeviceEc2InstancesInput{}
        output := &snowdevicemanagement.DescribeDeviceEc2InstancesOutput{}

        mockClient.On("DescribeDeviceEc2Instances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeviceEc2Instances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExecution", func(t *testing.T) {
        input := &snowdevicemanagement.DescribeExecutionInput{}
        output := &snowdevicemanagement.DescribeExecutionOutput{}

        mockClient.On("DescribeExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTask", func(t *testing.T) {
        input := &snowdevicemanagement.DescribeTaskInput{}
        output := &snowdevicemanagement.DescribeTaskOutput{}

        mockClient.On("DescribeTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceResources", func(t *testing.T) {
        input := &snowdevicemanagement.ListDeviceResourcesInput{}
        output := &snowdevicemanagement.ListDeviceResourcesOutput{}

        mockClient.On("ListDeviceResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &snowdevicemanagement.ListDevicesInput{}
        output := &snowdevicemanagement.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExecutions", func(t *testing.T) {
        input := &snowdevicemanagement.ListExecutionsInput{}
        output := &snowdevicemanagement.ListExecutionsOutput{}

        mockClient.On("ListExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &snowdevicemanagement.ListTagsForResourceInput{}
        output := &snowdevicemanagement.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTasks", func(t *testing.T) {
        input := &snowdevicemanagement.ListTasksInput{}
        output := &snowdevicemanagement.ListTasksOutput{}

        mockClient.On("ListTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &snowdevicemanagement.TagResourceInput{}
        output := &snowdevicemanagement.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &snowdevicemanagement.UntagResourceInput{}
        output := &snowdevicemanagement.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
