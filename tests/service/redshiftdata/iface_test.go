package redshiftdata_test

// tests for the redshiftdata service interface for this ../../../service/redshiftdata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshiftdata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshiftdata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshiftdata/redshiftdata_iface"
	"github.com/stretchr/testify/assert"
)

func TestRedshiftdataServiceCanBeMocked(t *testing.T) {
	var iface redshiftdata_iface.IClient
	iface = &redshiftdata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := redshiftdata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchExecuteStatement", func(t *testing.T) {
        input := &redshiftdata.BatchExecuteStatementInput{}
        output := &redshiftdata.BatchExecuteStatementOutput{}

        mockClient.On("BatchExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.BatchExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelStatement", func(t *testing.T) {
        input := &redshiftdata.CancelStatementInput{}
        output := &redshiftdata.CancelStatementOutput{}

        mockClient.On("CancelStatement", ctx, input).Return(output, nil)

        result, err := mockClient.CancelStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStatement", func(t *testing.T) {
        input := &redshiftdata.DescribeStatementInput{}
        output := &redshiftdata.DescribeStatementOutput{}

        mockClient.On("DescribeStatement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTable", func(t *testing.T) {
        input := &redshiftdata.DescribeTableInput{}
        output := &redshiftdata.DescribeTableOutput{}

        mockClient.On("DescribeTable", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteStatement", func(t *testing.T) {
        input := &redshiftdata.ExecuteStatementInput{}
        output := &redshiftdata.ExecuteStatementOutput{}

        mockClient.On("ExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStatementResult", func(t *testing.T) {
        input := &redshiftdata.GetStatementResultInput{}
        output := &redshiftdata.GetStatementResultOutput{}

        mockClient.On("GetStatementResult", ctx, input).Return(output, nil)

        result, err := mockClient.GetStatementResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatabases", func(t *testing.T) {
        input := &redshiftdata.ListDatabasesInput{}
        output := &redshiftdata.ListDatabasesOutput{}

        mockClient.On("ListDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemas", func(t *testing.T) {
        input := &redshiftdata.ListSchemasInput{}
        output := &redshiftdata.ListSchemasOutput{}

        mockClient.On("ListSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStatements", func(t *testing.T) {
        input := &redshiftdata.ListStatementsInput{}
        output := &redshiftdata.ListStatementsOutput{}

        mockClient.On("ListStatements", ctx, input).Return(output, nil)

        result, err := mockClient.ListStatements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTables", func(t *testing.T) {
        input := &redshiftdata.ListTablesInput{}
        output := &redshiftdata.ListTablesOutput{}

        mockClient.On("ListTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
