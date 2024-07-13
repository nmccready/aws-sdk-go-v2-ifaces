package ebs_test

// tests for the ebs service interface for this ../../../service/ebs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ebs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ebs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ebs/ebs_iface"
	"github.com/stretchr/testify/assert"
)

func TestEbsServiceCanBeMocked(t *testing.T) {
	var iface ebs_iface.IClient
	iface = &ebs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ebs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteSnapshot", func(t *testing.T) {
        input := &ebs.CompleteSnapshotInput{}
        output := &ebs.CompleteSnapshotOutput{}

        mockClient.On("CompleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnapshotBlock", func(t *testing.T) {
        input := &ebs.GetSnapshotBlockInput{}
        output := &ebs.GetSnapshotBlockOutput{}

        mockClient.On("GetSnapshotBlock", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnapshotBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChangedBlocks", func(t *testing.T) {
        input := &ebs.ListChangedBlocksInput{}
        output := &ebs.ListChangedBlocksOutput{}

        mockClient.On("ListChangedBlocks", ctx, input).Return(output, nil)

        result, err := mockClient.ListChangedBlocks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSnapshotBlocks", func(t *testing.T) {
        input := &ebs.ListSnapshotBlocksInput{}
        output := &ebs.ListSnapshotBlocksOutput{}

        mockClient.On("ListSnapshotBlocks", ctx, input).Return(output, nil)

        result, err := mockClient.ListSnapshotBlocks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSnapshotBlock", func(t *testing.T) {
        input := &ebs.PutSnapshotBlockInput{}
        output := &ebs.PutSnapshotBlockOutput{}

        mockClient.On("PutSnapshotBlock", ctx, input).Return(output, nil)

        result, err := mockClient.PutSnapshotBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSnapshot", func(t *testing.T) {
        input := &ebs.StartSnapshotInput{}
        output := &ebs.StartSnapshotOutput{}

        mockClient.On("StartSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.StartSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
