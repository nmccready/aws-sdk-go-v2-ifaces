package timestreamwrite_test

// tests for the timestreamwrite service interface for this ../../../service/timestreamwrite/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreamwrite/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreamwrite/timestreamwrite_iface"
	"github.com/stretchr/testify/assert"
)

func TestTimestreamwriteServiceCanBeMocked(t *testing.T) {
	var iface timestreamwrite_iface.IClient
	iface = &timestreamwrite.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := timestreamwrite.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchLoadTask", func(t *testing.T) {
        input := &timestreamwrite.CreateBatchLoadTaskInput{}
        output := &timestreamwrite.CreateBatchLoadTaskOutput{}

        mockClient.On("CreateBatchLoadTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchLoadTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatabase", func(t *testing.T) {
        input := &timestreamwrite.CreateDatabaseInput{}
        output := &timestreamwrite.CreateDatabaseOutput{}

        mockClient.On("CreateDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTable", func(t *testing.T) {
        input := &timestreamwrite.CreateTableInput{}
        output := &timestreamwrite.CreateTableOutput{}

        mockClient.On("CreateTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatabase", func(t *testing.T) {
        input := &timestreamwrite.DeleteDatabaseInput{}
        output := &timestreamwrite.DeleteDatabaseOutput{}

        mockClient.On("DeleteDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTable", func(t *testing.T) {
        input := &timestreamwrite.DeleteTableInput{}
        output := &timestreamwrite.DeleteTableOutput{}

        mockClient.On("DeleteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBatchLoadTask", func(t *testing.T) {
        input := &timestreamwrite.DescribeBatchLoadTaskInput{}
        output := &timestreamwrite.DescribeBatchLoadTaskOutput{}

        mockClient.On("DescribeBatchLoadTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBatchLoadTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatabase", func(t *testing.T) {
        input := &timestreamwrite.DescribeDatabaseInput{}
        output := &timestreamwrite.DescribeDatabaseOutput{}

        mockClient.On("DescribeDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoints", func(t *testing.T) {
        input := &timestreamwrite.DescribeEndpointsInput{}
        output := &timestreamwrite.DescribeEndpointsOutput{}

        mockClient.On("DescribeEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTable", func(t *testing.T) {
        input := &timestreamwrite.DescribeTableInput{}
        output := &timestreamwrite.DescribeTableOutput{}

        mockClient.On("DescribeTable", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchLoadTasks", func(t *testing.T) {
        input := &timestreamwrite.ListBatchLoadTasksInput{}
        output := &timestreamwrite.ListBatchLoadTasksOutput{}

        mockClient.On("ListBatchLoadTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchLoadTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatabases", func(t *testing.T) {
        input := &timestreamwrite.ListDatabasesInput{}
        output := &timestreamwrite.ListDatabasesOutput{}

        mockClient.On("ListDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTables", func(t *testing.T) {
        input := &timestreamwrite.ListTablesInput{}
        output := &timestreamwrite.ListTablesOutput{}

        mockClient.On("ListTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &timestreamwrite.ListTagsForResourceInput{}
        output := &timestreamwrite.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeBatchLoadTask", func(t *testing.T) {
        input := &timestreamwrite.ResumeBatchLoadTaskInput{}
        output := &timestreamwrite.ResumeBatchLoadTaskOutput{}

        mockClient.On("ResumeBatchLoadTask", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeBatchLoadTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &timestreamwrite.TagResourceInput{}
        output := &timestreamwrite.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &timestreamwrite.UntagResourceInput{}
        output := &timestreamwrite.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatabase", func(t *testing.T) {
        input := &timestreamwrite.UpdateDatabaseInput{}
        output := &timestreamwrite.UpdateDatabaseOutput{}

        mockClient.On("UpdateDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTable", func(t *testing.T) {
        input := &timestreamwrite.UpdateTableInput{}
        output := &timestreamwrite.UpdateTableOutput{}

        mockClient.On("UpdateTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestWriteRecords", func(t *testing.T) {
        input := &timestreamwrite.WriteRecordsInput{}
        output := &timestreamwrite.WriteRecordsOutput{}

        mockClient.On("WriteRecords", ctx, input).Return(output, nil)

        result, err := mockClient.WriteRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
