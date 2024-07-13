package rdsdata_test

// tests for the rdsdata service interface for this ../../../service/rdsdata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rdsdata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rdsdata/rdsdata_iface"
	"github.com/stretchr/testify/assert"
)

func TestRdsdataServiceCanBeMocked(t *testing.T) {
	var iface rdsdata_iface.IClient
	iface = &rdsdata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rdsdata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchExecuteStatement", func(t *testing.T) {
        input := &rdsdata.BatchExecuteStatementInput{}
        output := &rdsdata.BatchExecuteStatementOutput{}

        mockClient.On("BatchExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.BatchExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBeginTransaction", func(t *testing.T) {
        input := &rdsdata.BeginTransactionInput{}
        output := &rdsdata.BeginTransactionOutput{}

        mockClient.On("BeginTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.BeginTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCommitTransaction", func(t *testing.T) {
        input := &rdsdata.CommitTransactionInput{}
        output := &rdsdata.CommitTransactionOutput{}

        mockClient.On("CommitTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.CommitTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteSql", func(t *testing.T) {
        input := &rdsdata.ExecuteSqlInput{}
        output := &rdsdata.ExecuteSqlOutput{}

        mockClient.On("ExecuteSql", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteSql(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteStatement", func(t *testing.T) {
        input := &rdsdata.ExecuteStatementInput{}
        output := &rdsdata.ExecuteStatementOutput{}

        mockClient.On("ExecuteStatement", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRollbackTransaction", func(t *testing.T) {
        input := &rdsdata.RollbackTransactionInput{}
        output := &rdsdata.RollbackTransactionOutput{}

        mockClient.On("RollbackTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.RollbackTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
