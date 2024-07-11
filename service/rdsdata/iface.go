
package rdsdata

import (
    "github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

// IRdsdata defines the interface for rdsdata
type IRdsdata interface {
 Options() Options 
 BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error) 
 BeginTransaction(ctx context.Context, params *BeginTransactionInput, optFns ...func(*Options)) (*BeginTransactionOutput, error) 
 CommitTransaction(ctx context.Context, params *CommitTransactionInput, optFns ...func(*Options)) (*CommitTransactionOutput, error) 
 ExecuteSql(ctx context.Context, params *ExecuteSqlInput, optFns ...func(*Options)) (*ExecuteSqlOutput, error) 
 ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error) 
 RollbackTransaction(ctx context.Context, params *RollbackTransactionInput, optFns ...func(*Options)) (*RollbackTransactionOutput, error) 
}
