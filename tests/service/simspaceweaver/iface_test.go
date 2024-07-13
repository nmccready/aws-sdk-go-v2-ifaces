package simspaceweaver_test

// tests for the simspaceweaver service interface for this ../../../service/simspaceweaver/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/simspaceweaver/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/simspaceweaver/simspaceweaver_iface"
	"github.com/stretchr/testify/assert"
)

func TestSimspaceweaverServiceCanBeMocked(t *testing.T) {
	var iface simspaceweaver_iface.IClient
	iface = &simspaceweaver.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := simspaceweaver.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &simspaceweaver.CreateSnapshotInput{}
        output := &simspaceweaver.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &simspaceweaver.DeleteAppInput{}
        output := &simspaceweaver.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSimulation", func(t *testing.T) {
        input := &simspaceweaver.DeleteSimulationInput{}
        output := &simspaceweaver.DeleteSimulationOutput{}

        mockClient.On("DeleteSimulation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSimulation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApp", func(t *testing.T) {
        input := &simspaceweaver.DescribeAppInput{}
        output := &simspaceweaver.DescribeAppOutput{}

        mockClient.On("DescribeApp", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSimulation", func(t *testing.T) {
        input := &simspaceweaver.DescribeSimulationInput{}
        output := &simspaceweaver.DescribeSimulationOutput{}

        mockClient.On("DescribeSimulation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSimulation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApps", func(t *testing.T) {
        input := &simspaceweaver.ListAppsInput{}
        output := &simspaceweaver.ListAppsOutput{}

        mockClient.On("ListApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSimulations", func(t *testing.T) {
        input := &simspaceweaver.ListSimulationsInput{}
        output := &simspaceweaver.ListSimulationsOutput{}

        mockClient.On("ListSimulations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSimulations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &simspaceweaver.ListTagsForResourceInput{}
        output := &simspaceweaver.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApp", func(t *testing.T) {
        input := &simspaceweaver.StartAppInput{}
        output := &simspaceweaver.StartAppOutput{}

        mockClient.On("StartApp", ctx, input).Return(output, nil)

        result, err := mockClient.StartApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartClock", func(t *testing.T) {
        input := &simspaceweaver.StartClockInput{}
        output := &simspaceweaver.StartClockOutput{}

        mockClient.On("StartClock", ctx, input).Return(output, nil)

        result, err := mockClient.StartClock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSimulation", func(t *testing.T) {
        input := &simspaceweaver.StartSimulationInput{}
        output := &simspaceweaver.StartSimulationOutput{}

        mockClient.On("StartSimulation", ctx, input).Return(output, nil)

        result, err := mockClient.StartSimulation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApp", func(t *testing.T) {
        input := &simspaceweaver.StopAppInput{}
        output := &simspaceweaver.StopAppOutput{}

        mockClient.On("StopApp", ctx, input).Return(output, nil)

        result, err := mockClient.StopApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopClock", func(t *testing.T) {
        input := &simspaceweaver.StopClockInput{}
        output := &simspaceweaver.StopClockOutput{}

        mockClient.On("StopClock", ctx, input).Return(output, nil)

        result, err := mockClient.StopClock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSimulation", func(t *testing.T) {
        input := &simspaceweaver.StopSimulationInput{}
        output := &simspaceweaver.StopSimulationOutput{}

        mockClient.On("StopSimulation", ctx, input).Return(output, nil)

        result, err := mockClient.StopSimulation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &simspaceweaver.TagResourceInput{}
        output := &simspaceweaver.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &simspaceweaver.UntagResourceInput{}
        output := &simspaceweaver.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
