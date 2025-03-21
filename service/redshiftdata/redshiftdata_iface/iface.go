
package redshiftdata_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/redshiftdata"
)

// IClient defines the interface for redshiftdata
type IClient interface {
 Options() Options 
 BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error) 
 CancelStatement(ctx context.Context, params *CancelStatementInput, optFns ...func(*Options)) (*CancelStatementOutput, error) 
 DescribeStatement(ctx context.Context, params *DescribeStatementInput, optFns ...func(*Options)) (*DescribeStatementOutput, error) 
 DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error) 
 ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error) 
 GetStatementResult(ctx context.Context, params *GetStatementResultInput, optFns ...func(*Options)) (*GetStatementResultOutput, error) 
 GetStatementResultV2(ctx context.Context, params *GetStatementResultV2Input, optFns ...func(*Options)) (*GetStatementResultV2Output, error) 
 ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) 
 ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) 
 ListStatements(ctx context.Context, params *ListStatementsInput, optFns ...func(*Options)) (*ListStatementsOutput, error) 
 ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) 
}
