
package bcmdataexports

import (
    "github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
)

// IBcmdataexports defines the interface for bcmdataexports
type IBcmdataexports interface {
 Options() Options 
 CreateExport(ctx context.Context, params *CreateExportInput, optFns ...func(*Options)) (*CreateExportOutput, error) 
 DeleteExport(ctx context.Context, params *DeleteExportInput, optFns ...func(*Options)) (*DeleteExportOutput, error) 
 GetExecution(ctx context.Context, params *GetExecutionInput, optFns ...func(*Options)) (*GetExecutionOutput, error) 
 GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) 
 GetTable(ctx context.Context, params *GetTableInput, optFns ...func(*Options)) (*GetTableOutput, error) 
 ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) 
 ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) 
 ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateExport(ctx context.Context, params *UpdateExportInput, optFns ...func(*Options)) (*UpdateExportOutput, error) 
}
