package sagemakera2iruntime_test

// tests for the sagemakera2iruntime service interface for this ../../../service/sagemakera2iruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakera2iruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakera2iruntime/sagemakera2iruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakera2iruntimeServiceCanBeMocked(t *testing.T) {
	var iface sagemakera2iruntime_iface.IClient
	iface = &sagemakera2iruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakera2iruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHumanLoop", func(t *testing.T) {
        input := &sagemakera2iruntime.DeleteHumanLoopInput{}
        output := &sagemakera2iruntime.DeleteHumanLoopOutput{}

        mockClient.On("DeleteHumanLoop", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHumanLoop(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHumanLoop", func(t *testing.T) {
        input := &sagemakera2iruntime.DescribeHumanLoopInput{}
        output := &sagemakera2iruntime.DescribeHumanLoopOutput{}

        mockClient.On("DescribeHumanLoop", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHumanLoop(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHumanLoops", func(t *testing.T) {
        input := &sagemakera2iruntime.ListHumanLoopsInput{}
        output := &sagemakera2iruntime.ListHumanLoopsOutput{}

        mockClient.On("ListHumanLoops", ctx, input).Return(output, nil)

        result, err := mockClient.ListHumanLoops(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartHumanLoop", func(t *testing.T) {
        input := &sagemakera2iruntime.StartHumanLoopInput{}
        output := &sagemakera2iruntime.StartHumanLoopOutput{}

        mockClient.On("StartHumanLoop", ctx, input).Return(output, nil)

        result, err := mockClient.StartHumanLoop(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopHumanLoop", func(t *testing.T) {
        input := &sagemakera2iruntime.StopHumanLoopInput{}
        output := &sagemakera2iruntime.StopHumanLoopOutput{}

        mockClient.On("StopHumanLoop", ctx, input).Return(output, nil)

        result, err := mockClient.StopHumanLoop(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
