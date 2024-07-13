package managedblockchainquery_test

// tests for the managedblockchainquery service interface for this ../../../service/managedblockchainquery/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/managedblockchainquery/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/managedblockchainquery/managedblockchainquery_iface"
	"github.com/stretchr/testify/assert"
)

func TestManagedblockchainqueryServiceCanBeMocked(t *testing.T) {
	var iface managedblockchainquery_iface.IClient
	iface = &managedblockchainquery.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := managedblockchainquery.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetTokenBalance", func(t *testing.T) {
        input := &managedblockchainquery.BatchGetTokenBalanceInput{}
        output := &managedblockchainquery.BatchGetTokenBalanceOutput{}

        mockClient.On("BatchGetTokenBalance", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetTokenBalance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssetContract", func(t *testing.T) {
        input := &managedblockchainquery.GetAssetContractInput{}
        output := &managedblockchainquery.GetAssetContractOutput{}

        mockClient.On("GetAssetContract", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssetContract(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTokenBalance", func(t *testing.T) {
        input := &managedblockchainquery.GetTokenBalanceInput{}
        output := &managedblockchainquery.GetTokenBalanceOutput{}

        mockClient.On("GetTokenBalance", ctx, input).Return(output, nil)

        result, err := mockClient.GetTokenBalance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransaction", func(t *testing.T) {
        input := &managedblockchainquery.GetTransactionInput{}
        output := &managedblockchainquery.GetTransactionOutput{}

        mockClient.On("GetTransaction", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetContracts", func(t *testing.T) {
        input := &managedblockchainquery.ListAssetContractsInput{}
        output := &managedblockchainquery.ListAssetContractsOutput{}

        mockClient.On("ListAssetContracts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetContracts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFilteredTransactionEvents", func(t *testing.T) {
        input := &managedblockchainquery.ListFilteredTransactionEventsInput{}
        output := &managedblockchainquery.ListFilteredTransactionEventsOutput{}

        mockClient.On("ListFilteredTransactionEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListFilteredTransactionEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTokenBalances", func(t *testing.T) {
        input := &managedblockchainquery.ListTokenBalancesInput{}
        output := &managedblockchainquery.ListTokenBalancesOutput{}

        mockClient.On("ListTokenBalances", ctx, input).Return(output, nil)

        result, err := mockClient.ListTokenBalances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTransactionEvents", func(t *testing.T) {
        input := &managedblockchainquery.ListTransactionEventsInput{}
        output := &managedblockchainquery.ListTransactionEventsOutput{}

        mockClient.On("ListTransactionEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListTransactionEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTransactions", func(t *testing.T) {
        input := &managedblockchainquery.ListTransactionsInput{}
        output := &managedblockchainquery.ListTransactionsOutput{}

        mockClient.On("ListTransactions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTransactions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
