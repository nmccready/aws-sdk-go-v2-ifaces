package pipes_test

// tests for the pipes service interface for this ../../../service/pipes/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pipes"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pipes/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pipes/pipes_iface"
	"github.com/stretchr/testify/assert"
)

func TestPipesServiceCanBeMocked(t *testing.T) {
	var iface pipes_iface.IClient
	iface = &pipes.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pipes.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePipe", func(t *testing.T) {
        input := &pipes.CreatePipeInput{}
        output := &pipes.CreatePipeOutput{}

        mockClient.On("CreatePipe", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePipe", func(t *testing.T) {
        input := &pipes.DeletePipeInput{}
        output := &pipes.DeletePipeOutput{}

        mockClient.On("DeletePipe", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePipe", func(t *testing.T) {
        input := &pipes.DescribePipeInput{}
        output := &pipes.DescribePipeOutput{}

        mockClient.On("DescribePipe", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPipes", func(t *testing.T) {
        input := &pipes.ListPipesInput{}
        output := &pipes.ListPipesOutput{}

        mockClient.On("ListPipes", ctx, input).Return(output, nil)

        result, err := mockClient.ListPipes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pipes.ListTagsForResourceInput{}
        output := &pipes.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPipe", func(t *testing.T) {
        input := &pipes.StartPipeInput{}
        output := &pipes.StartPipeOutput{}

        mockClient.On("StartPipe", ctx, input).Return(output, nil)

        result, err := mockClient.StartPipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPipe", func(t *testing.T) {
        input := &pipes.StopPipeInput{}
        output := &pipes.StopPipeOutput{}

        mockClient.On("StopPipe", ctx, input).Return(output, nil)

        result, err := mockClient.StopPipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pipes.TagResourceInput{}
        output := &pipes.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pipes.UntagResourceInput{}
        output := &pipes.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePipe", func(t *testing.T) {
        input := &pipes.UpdatePipeInput{}
        output := &pipes.UpdatePipeOutput{}

        mockClient.On("UpdatePipe", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
