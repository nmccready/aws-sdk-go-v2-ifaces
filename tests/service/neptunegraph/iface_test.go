package neptunegraph_test

// tests for the neptunegraph service interface for this ../../../service/neptunegraph/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/neptunegraph"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptunegraph/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/neptunegraph/neptunegraph_iface"
	"github.com/stretchr/testify/assert"
)

func TestNeptunegraphServiceCanBeMocked(t *testing.T) {
	var iface neptunegraph_iface.IClient
	iface = &neptunegraph.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := neptunegraph.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelImportTask", func(t *testing.T) {
        input := &neptunegraph.CancelImportTaskInput{}
        output := &neptunegraph.CancelImportTaskOutput{}

        mockClient.On("CancelImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelQuery", func(t *testing.T) {
        input := &neptunegraph.CancelQueryInput{}
        output := &neptunegraph.CancelQueryOutput{}

        mockClient.On("CancelQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CancelQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGraph", func(t *testing.T) {
        input := &neptunegraph.CreateGraphInput{}
        output := &neptunegraph.CreateGraphOutput{}

        mockClient.On("CreateGraph", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGraphSnapshot", func(t *testing.T) {
        input := &neptunegraph.CreateGraphSnapshotInput{}
        output := &neptunegraph.CreateGraphSnapshotOutput{}

        mockClient.On("CreateGraphSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGraphSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGraphUsingImportTask", func(t *testing.T) {
        input := &neptunegraph.CreateGraphUsingImportTaskInput{}
        output := &neptunegraph.CreateGraphUsingImportTaskOutput{}

        mockClient.On("CreateGraphUsingImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGraphUsingImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrivateGraphEndpoint", func(t *testing.T) {
        input := &neptunegraph.CreatePrivateGraphEndpointInput{}
        output := &neptunegraph.CreatePrivateGraphEndpointOutput{}

        mockClient.On("CreatePrivateGraphEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrivateGraphEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGraph", func(t *testing.T) {
        input := &neptunegraph.DeleteGraphInput{}
        output := &neptunegraph.DeleteGraphOutput{}

        mockClient.On("DeleteGraph", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGraphSnapshot", func(t *testing.T) {
        input := &neptunegraph.DeleteGraphSnapshotInput{}
        output := &neptunegraph.DeleteGraphSnapshotOutput{}

        mockClient.On("DeleteGraphSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGraphSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrivateGraphEndpoint", func(t *testing.T) {
        input := &neptunegraph.DeletePrivateGraphEndpointInput{}
        output := &neptunegraph.DeletePrivateGraphEndpointOutput{}

        mockClient.On("DeletePrivateGraphEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrivateGraphEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteQuery", func(t *testing.T) {
        input := &neptunegraph.ExecuteQueryInput{}
        output := &neptunegraph.ExecuteQueryOutput{}

        mockClient.On("ExecuteQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGraph", func(t *testing.T) {
        input := &neptunegraph.GetGraphInput{}
        output := &neptunegraph.GetGraphOutput{}

        mockClient.On("GetGraph", ctx, input).Return(output, nil)

        result, err := mockClient.GetGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGraphSnapshot", func(t *testing.T) {
        input := &neptunegraph.GetGraphSnapshotInput{}
        output := &neptunegraph.GetGraphSnapshotOutput{}

        mockClient.On("GetGraphSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetGraphSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGraphSummary", func(t *testing.T) {
        input := &neptunegraph.GetGraphSummaryInput{}
        output := &neptunegraph.GetGraphSummaryOutput{}

        mockClient.On("GetGraphSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetGraphSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportTask", func(t *testing.T) {
        input := &neptunegraph.GetImportTaskInput{}
        output := &neptunegraph.GetImportTaskOutput{}

        mockClient.On("GetImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrivateGraphEndpoint", func(t *testing.T) {
        input := &neptunegraph.GetPrivateGraphEndpointInput{}
        output := &neptunegraph.GetPrivateGraphEndpointOutput{}

        mockClient.On("GetPrivateGraphEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrivateGraphEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQuery", func(t *testing.T) {
        input := &neptunegraph.GetQueryInput{}
        output := &neptunegraph.GetQueryOutput{}

        mockClient.On("GetQuery", ctx, input).Return(output, nil)

        result, err := mockClient.GetQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGraphSnapshots", func(t *testing.T) {
        input := &neptunegraph.ListGraphSnapshotsInput{}
        output := &neptunegraph.ListGraphSnapshotsOutput{}

        mockClient.On("ListGraphSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.ListGraphSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGraphs", func(t *testing.T) {
        input := &neptunegraph.ListGraphsInput{}
        output := &neptunegraph.ListGraphsOutput{}

        mockClient.On("ListGraphs", ctx, input).Return(output, nil)

        result, err := mockClient.ListGraphs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportTasks", func(t *testing.T) {
        input := &neptunegraph.ListImportTasksInput{}
        output := &neptunegraph.ListImportTasksOutput{}

        mockClient.On("ListImportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrivateGraphEndpoints", func(t *testing.T) {
        input := &neptunegraph.ListPrivateGraphEndpointsInput{}
        output := &neptunegraph.ListPrivateGraphEndpointsOutput{}

        mockClient.On("ListPrivateGraphEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrivateGraphEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueries", func(t *testing.T) {
        input := &neptunegraph.ListQueriesInput{}
        output := &neptunegraph.ListQueriesOutput{}

        mockClient.On("ListQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &neptunegraph.ListTagsForResourceInput{}
        output := &neptunegraph.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetGraph", func(t *testing.T) {
        input := &neptunegraph.ResetGraphInput{}
        output := &neptunegraph.ResetGraphOutput{}

        mockClient.On("ResetGraph", ctx, input).Return(output, nil)

        result, err := mockClient.ResetGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreGraphFromSnapshot", func(t *testing.T) {
        input := &neptunegraph.RestoreGraphFromSnapshotInput{}
        output := &neptunegraph.RestoreGraphFromSnapshotOutput{}

        mockClient.On("RestoreGraphFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreGraphFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportTask", func(t *testing.T) {
        input := &neptunegraph.StartImportTaskInput{}
        output := &neptunegraph.StartImportTaskOutput{}

        mockClient.On("StartImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &neptunegraph.TagResourceInput{}
        output := &neptunegraph.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &neptunegraph.UntagResourceInput{}
        output := &neptunegraph.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGraph", func(t *testing.T) {
        input := &neptunegraph.UpdateGraphInput{}
        output := &neptunegraph.UpdateGraphOutput{}

        mockClient.On("UpdateGraph", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
