
package neptunegraph

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/neptunegraph"
)

// IClient defines the interface for neptunegraph
type IClient interface {
 Options() Options 
 CancelImportTask(ctx context.Context, params *CancelImportTaskInput, optFns ...func(*Options)) (*CancelImportTaskOutput, error) 
 CancelQuery(ctx context.Context, params *CancelQueryInput, optFns ...func(*Options)) (*CancelQueryOutput, error) 
 CreateGraph(ctx context.Context, params *CreateGraphInput, optFns ...func(*Options)) (*CreateGraphOutput, error) 
 CreateGraphSnapshot(ctx context.Context, params *CreateGraphSnapshotInput, optFns ...func(*Options)) (*CreateGraphSnapshotOutput, error) 
 CreateGraphUsingImportTask(ctx context.Context, params *CreateGraphUsingImportTaskInput, optFns ...func(*Options)) (*CreateGraphUsingImportTaskOutput, error) 
 CreatePrivateGraphEndpoint(ctx context.Context, params *CreatePrivateGraphEndpointInput, optFns ...func(*Options)) (*CreatePrivateGraphEndpointOutput, error) 
 DeleteGraph(ctx context.Context, params *DeleteGraphInput, optFns ...func(*Options)) (*DeleteGraphOutput, error) 
 DeleteGraphSnapshot(ctx context.Context, params *DeleteGraphSnapshotInput, optFns ...func(*Options)) (*DeleteGraphSnapshotOutput, error) 
 DeletePrivateGraphEndpoint(ctx context.Context, params *DeletePrivateGraphEndpointInput, optFns ...func(*Options)) (*DeletePrivateGraphEndpointOutput, error) 
 ExecuteQuery(ctx context.Context, params *ExecuteQueryInput, optFns ...func(*Options)) (*ExecuteQueryOutput, error) 
 GetGraph(ctx context.Context, params *GetGraphInput, optFns ...func(*Options)) (*GetGraphOutput, error) 
 GetGraphSnapshot(ctx context.Context, params *GetGraphSnapshotInput, optFns ...func(*Options)) (*GetGraphSnapshotOutput, error) 
 GetGraphSummary(ctx context.Context, params *GetGraphSummaryInput, optFns ...func(*Options)) (*GetGraphSummaryOutput, error) 
 GetImportTask(ctx context.Context, params *GetImportTaskInput, optFns ...func(*Options)) (*GetImportTaskOutput, error) 
 GetPrivateGraphEndpoint(ctx context.Context, params *GetPrivateGraphEndpointInput, optFns ...func(*Options)) (*GetPrivateGraphEndpointOutput, error) 
 GetQuery(ctx context.Context, params *GetQueryInput, optFns ...func(*Options)) (*GetQueryOutput, error) 
 ListGraphSnapshots(ctx context.Context, params *ListGraphSnapshotsInput, optFns ...func(*Options)) (*ListGraphSnapshotsOutput, error) 
 ListGraphs(ctx context.Context, params *ListGraphsInput, optFns ...func(*Options)) (*ListGraphsOutput, error) 
 ListImportTasks(ctx context.Context, params *ListImportTasksInput, optFns ...func(*Options)) (*ListImportTasksOutput, error) 
 ListPrivateGraphEndpoints(ctx context.Context, params *ListPrivateGraphEndpointsInput, optFns ...func(*Options)) (*ListPrivateGraphEndpointsOutput, error) 
 ListQueries(ctx context.Context, params *ListQueriesInput, optFns ...func(*Options)) (*ListQueriesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ResetGraph(ctx context.Context, params *ResetGraphInput, optFns ...func(*Options)) (*ResetGraphOutput, error) 
 RestoreGraphFromSnapshot(ctx context.Context, params *RestoreGraphFromSnapshotInput, optFns ...func(*Options)) (*RestoreGraphFromSnapshotOutput, error) 
 StartImportTask(ctx context.Context, params *StartImportTaskInput, optFns ...func(*Options)) (*StartImportTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGraph(ctx context.Context, params *UpdateGraphInput, optFns ...func(*Options)) (*UpdateGraphOutput, error) 
}
