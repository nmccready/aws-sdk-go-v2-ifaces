
package machinelearning

import (
    "github.com/aws/aws-sdk-go-v2/service/machinelearning"
)

// IClient defines the interface for machinelearning
type IClient interface {
 Options() Options 
 AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) 
 CreateBatchPrediction(ctx context.Context, params *CreateBatchPredictionInput, optFns ...func(*Options)) (*CreateBatchPredictionOutput, error) 
 CreateDataSourceFromRDS(ctx context.Context, params *CreateDataSourceFromRDSInput, optFns ...func(*Options)) (*CreateDataSourceFromRDSOutput, error) 
 CreateDataSourceFromRedshift(ctx context.Context, params *CreateDataSourceFromRedshiftInput, optFns ...func(*Options)) (*CreateDataSourceFromRedshiftOutput, error) 
 CreateDataSourceFromS3(ctx context.Context, params *CreateDataSourceFromS3Input, optFns ...func(*Options)) (*CreateDataSourceFromS3Output, error) 
 CreateEvaluation(ctx context.Context, params *CreateEvaluationInput, optFns ...func(*Options)) (*CreateEvaluationOutput, error) 
 CreateMLModel(ctx context.Context, params *CreateMLModelInput, optFns ...func(*Options)) (*CreateMLModelOutput, error) 
 CreateRealtimeEndpoint(ctx context.Context, params *CreateRealtimeEndpointInput, optFns ...func(*Options)) (*CreateRealtimeEndpointOutput, error) 
 DeleteBatchPrediction(ctx context.Context, params *DeleteBatchPredictionInput, optFns ...func(*Options)) (*DeleteBatchPredictionOutput, error) 
 DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) 
 DeleteEvaluation(ctx context.Context, params *DeleteEvaluationInput, optFns ...func(*Options)) (*DeleteEvaluationOutput, error) 
 DeleteMLModel(ctx context.Context, params *DeleteMLModelInput, optFns ...func(*Options)) (*DeleteMLModelOutput, error) 
 DeleteRealtimeEndpoint(ctx context.Context, params *DeleteRealtimeEndpointInput, optFns ...func(*Options)) (*DeleteRealtimeEndpointOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DescribeBatchPredictions(ctx context.Context, params *DescribeBatchPredictionsInput, optFns ...func(*Options)) (*DescribeBatchPredictionsOutput, error) 
 DescribeDataSources(ctx context.Context, params *DescribeDataSourcesInput, optFns ...func(*Options)) (*DescribeDataSourcesOutput, error) 
 DescribeEvaluations(ctx context.Context, params *DescribeEvaluationsInput, optFns ...func(*Options)) (*DescribeEvaluationsOutput, error) 
 DescribeMLModels(ctx context.Context, params *DescribeMLModelsInput, optFns ...func(*Options)) (*DescribeMLModelsOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 GetBatchPrediction(ctx context.Context, params *GetBatchPredictionInput, optFns ...func(*Options)) (*GetBatchPredictionOutput, error) 
 GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) 
 GetEvaluation(ctx context.Context, params *GetEvaluationInput, optFns ...func(*Options)) (*GetEvaluationOutput, error) 
 GetMLModel(ctx context.Context, params *GetMLModelInput, optFns ...func(*Options)) (*GetMLModelOutput, error) 
 Predict(ctx context.Context, params *PredictInput, optFns ...func(*Options)) (*PredictOutput, error) 
 UpdateBatchPrediction(ctx context.Context, params *UpdateBatchPredictionInput, optFns ...func(*Options)) (*UpdateBatchPredictionOutput, error) 
 UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) 
 UpdateEvaluation(ctx context.Context, params *UpdateEvaluationInput, optFns ...func(*Options)) (*UpdateEvaluationOutput, error) 
 UpdateMLModel(ctx context.Context, params *UpdateMLModelInput, optFns ...func(*Options)) (*UpdateMLModelOutput, error) 
}
