
package timestreamwrite

import (
    "github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

// ITimestreamwrite defines the interface for timestreamwrite
type ITimestreamwrite interface {
 Options() Options 
 CreateBatchLoadTask(ctx context.Context, params *CreateBatchLoadTaskInput, optFns ...func(*Options)) (*CreateBatchLoadTaskOutput, error) 
 CreateDatabase(ctx context.Context, params *CreateDatabaseInput, optFns ...func(*Options)) (*CreateDatabaseOutput, error) 
 CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) 
 DeleteDatabase(ctx context.Context, params *DeleteDatabaseInput, optFns ...func(*Options)) (*DeleteDatabaseOutput, error) 
 DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) 
 DescribeBatchLoadTask(ctx context.Context, params *DescribeBatchLoadTaskInput, optFns ...func(*Options)) (*DescribeBatchLoadTaskOutput, error) 
 DescribeDatabase(ctx context.Context, params *DescribeDatabaseInput, optFns ...func(*Options)) (*DescribeDatabaseOutput, error) 
 DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) 
 DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error) 
 ListBatchLoadTasks(ctx context.Context, params *ListBatchLoadTasksInput, optFns ...func(*Options)) (*ListBatchLoadTasksOutput, error) 
 ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) 
 ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ResumeBatchLoadTask(ctx context.Context, params *ResumeBatchLoadTaskInput, optFns ...func(*Options)) (*ResumeBatchLoadTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDatabase(ctx context.Context, params *UpdateDatabaseInput, optFns ...func(*Options)) (*UpdateDatabaseOutput, error) 
 UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) 
 WriteRecords(ctx context.Context, params *WriteRecordsInput, optFns ...func(*Options)) (*WriteRecordsOutput, error) 
}
