
package managedblockchainquery_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"
)

// IClient defines the interface for managedblockchainquery
type IClient interface {
 Options() Options 
 BatchGetTokenBalance(ctx context.Context, params *BatchGetTokenBalanceInput, optFns ...func(*Options)) (*BatchGetTokenBalanceOutput, error) 
 GetAssetContract(ctx context.Context, params *GetAssetContractInput, optFns ...func(*Options)) (*GetAssetContractOutput, error) 
 GetTokenBalance(ctx context.Context, params *GetTokenBalanceInput, optFns ...func(*Options)) (*GetTokenBalanceOutput, error) 
 GetTransaction(ctx context.Context, params *GetTransactionInput, optFns ...func(*Options)) (*GetTransactionOutput, error) 
 ListAssetContracts(ctx context.Context, params *ListAssetContractsInput, optFns ...func(*Options)) (*ListAssetContractsOutput, error) 
 ListFilteredTransactionEvents(ctx context.Context, params *ListFilteredTransactionEventsInput, optFns ...func(*Options)) (*ListFilteredTransactionEventsOutput, error) 
 ListTokenBalances(ctx context.Context, params *ListTokenBalancesInput, optFns ...func(*Options)) (*ListTokenBalancesOutput, error) 
 ListTransactionEvents(ctx context.Context, params *ListTransactionEventsInput, optFns ...func(*Options)) (*ListTransactionEventsOutput, error) 
 ListTransactions(ctx context.Context, params *ListTransactionsInput, optFns ...func(*Options)) (*ListTransactionsOutput, error) 
}