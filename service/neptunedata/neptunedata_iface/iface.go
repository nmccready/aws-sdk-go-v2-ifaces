
package neptunedata_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/neptunedata"
)

// IClient defines the interface for neptunedata
type IClient interface {
 Options() Options 
 CancelGremlinQuery(ctx context.Context, params *CancelGremlinQueryInput, optFns ...func(*Options)) (*CancelGremlinQueryOutput, error) 
 CancelLoaderJob(ctx context.Context, params *CancelLoaderJobInput, optFns ...func(*Options)) (*CancelLoaderJobOutput, error) 
 CancelMLDataProcessingJob(ctx context.Context, params *CancelMLDataProcessingJobInput, optFns ...func(*Options)) (*CancelMLDataProcessingJobOutput, error) 
 CancelMLModelTrainingJob(ctx context.Context, params *CancelMLModelTrainingJobInput, optFns ...func(*Options)) (*CancelMLModelTrainingJobOutput, error) 
 CancelMLModelTransformJob(ctx context.Context, params *CancelMLModelTransformJobInput, optFns ...func(*Options)) (*CancelMLModelTransformJobOutput, error) 
 CancelOpenCypherQuery(ctx context.Context, params *CancelOpenCypherQueryInput, optFns ...func(*Options)) (*CancelOpenCypherQueryOutput, error) 
 CreateMLEndpoint(ctx context.Context, params *CreateMLEndpointInput, optFns ...func(*Options)) (*CreateMLEndpointOutput, error) 
 DeleteMLEndpoint(ctx context.Context, params *DeleteMLEndpointInput, optFns ...func(*Options)) (*DeleteMLEndpointOutput, error) 
 DeletePropertygraphStatistics(ctx context.Context, params *DeletePropertygraphStatisticsInput, optFns ...func(*Options)) (*DeletePropertygraphStatisticsOutput, error) 
 DeleteSparqlStatistics(ctx context.Context, params *DeleteSparqlStatisticsInput, optFns ...func(*Options)) (*DeleteSparqlStatisticsOutput, error) 
 ExecuteFastReset(ctx context.Context, params *ExecuteFastResetInput, optFns ...func(*Options)) (*ExecuteFastResetOutput, error) 
 ExecuteGremlinExplainQuery(ctx context.Context, params *ExecuteGremlinExplainQueryInput, optFns ...func(*Options)) (*ExecuteGremlinExplainQueryOutput, error) 
 ExecuteGremlinProfileQuery(ctx context.Context, params *ExecuteGremlinProfileQueryInput, optFns ...func(*Options)) (*ExecuteGremlinProfileQueryOutput, error) 
 ExecuteGremlinQuery(ctx context.Context, params *ExecuteGremlinQueryInput, optFns ...func(*Options)) (*ExecuteGremlinQueryOutput, error) 
 ExecuteOpenCypherExplainQuery(ctx context.Context, params *ExecuteOpenCypherExplainQueryInput, optFns ...func(*Options)) (*ExecuteOpenCypherExplainQueryOutput, error) 
 ExecuteOpenCypherQuery(ctx context.Context, params *ExecuteOpenCypherQueryInput, optFns ...func(*Options)) (*ExecuteOpenCypherQueryOutput, error) 
 GetEngineStatus(ctx context.Context, params *GetEngineStatusInput, optFns ...func(*Options)) (*GetEngineStatusOutput, error) 
 GetGremlinQueryStatus(ctx context.Context, params *GetGremlinQueryStatusInput, optFns ...func(*Options)) (*GetGremlinQueryStatusOutput, error) 
 GetLoaderJobStatus(ctx context.Context, params *GetLoaderJobStatusInput, optFns ...func(*Options)) (*GetLoaderJobStatusOutput, error) 
 GetMLDataProcessingJob(ctx context.Context, params *GetMLDataProcessingJobInput, optFns ...func(*Options)) (*GetMLDataProcessingJobOutput, error) 
 GetMLEndpoint(ctx context.Context, params *GetMLEndpointInput, optFns ...func(*Options)) (*GetMLEndpointOutput, error) 
 GetMLModelTrainingJob(ctx context.Context, params *GetMLModelTrainingJobInput, optFns ...func(*Options)) (*GetMLModelTrainingJobOutput, error) 
 GetMLModelTransformJob(ctx context.Context, params *GetMLModelTransformJobInput, optFns ...func(*Options)) (*GetMLModelTransformJobOutput, error) 
 GetOpenCypherQueryStatus(ctx context.Context, params *GetOpenCypherQueryStatusInput, optFns ...func(*Options)) (*GetOpenCypherQueryStatusOutput, error) 
 GetPropertygraphStatistics(ctx context.Context, params *GetPropertygraphStatisticsInput, optFns ...func(*Options)) (*GetPropertygraphStatisticsOutput, error) 
 GetPropertygraphStream(ctx context.Context, params *GetPropertygraphStreamInput, optFns ...func(*Options)) (*GetPropertygraphStreamOutput, error) 
 GetPropertygraphSummary(ctx context.Context, params *GetPropertygraphSummaryInput, optFns ...func(*Options)) (*GetPropertygraphSummaryOutput, error) 
 GetRDFGraphSummary(ctx context.Context, params *GetRDFGraphSummaryInput, optFns ...func(*Options)) (*GetRDFGraphSummaryOutput, error) 
 GetSparqlStatistics(ctx context.Context, params *GetSparqlStatisticsInput, optFns ...func(*Options)) (*GetSparqlStatisticsOutput, error) 
 GetSparqlStream(ctx context.Context, params *GetSparqlStreamInput, optFns ...func(*Options)) (*GetSparqlStreamOutput, error) 
 ListGremlinQueries(ctx context.Context, params *ListGremlinQueriesInput, optFns ...func(*Options)) (*ListGremlinQueriesOutput, error) 
 ListLoaderJobs(ctx context.Context, params *ListLoaderJobsInput, optFns ...func(*Options)) (*ListLoaderJobsOutput, error) 
 ListMLDataProcessingJobs(ctx context.Context, params *ListMLDataProcessingJobsInput, optFns ...func(*Options)) (*ListMLDataProcessingJobsOutput, error) 
 ListMLEndpoints(ctx context.Context, params *ListMLEndpointsInput, optFns ...func(*Options)) (*ListMLEndpointsOutput, error) 
 ListMLModelTrainingJobs(ctx context.Context, params *ListMLModelTrainingJobsInput, optFns ...func(*Options)) (*ListMLModelTrainingJobsOutput, error) 
 ListMLModelTransformJobs(ctx context.Context, params *ListMLModelTransformJobsInput, optFns ...func(*Options)) (*ListMLModelTransformJobsOutput, error) 
 ListOpenCypherQueries(ctx context.Context, params *ListOpenCypherQueriesInput, optFns ...func(*Options)) (*ListOpenCypherQueriesOutput, error) 
 ManagePropertygraphStatistics(ctx context.Context, params *ManagePropertygraphStatisticsInput, optFns ...func(*Options)) (*ManagePropertygraphStatisticsOutput, error) 
 ManageSparqlStatistics(ctx context.Context, params *ManageSparqlStatisticsInput, optFns ...func(*Options)) (*ManageSparqlStatisticsOutput, error) 
 StartLoaderJob(ctx context.Context, params *StartLoaderJobInput, optFns ...func(*Options)) (*StartLoaderJobOutput, error) 
 StartMLDataProcessingJob(ctx context.Context, params *StartMLDataProcessingJobInput, optFns ...func(*Options)) (*StartMLDataProcessingJobOutput, error) 
 StartMLModelTrainingJob(ctx context.Context, params *StartMLModelTrainingJobInput, optFns ...func(*Options)) (*StartMLModelTrainingJobOutput, error) 
 StartMLModelTransformJob(ctx context.Context, params *StartMLModelTransformJobInput, optFns ...func(*Options)) (*StartMLModelTransformJobOutput, error) 
}
