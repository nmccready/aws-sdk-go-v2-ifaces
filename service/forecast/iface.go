
package forecast

import (
    "github.com/aws/aws-sdk-go-v2/service/forecast"
)

// IForecast defines the interface for forecast
type IForecast interface {
 Options() Options 
 CreateAutoPredictor(ctx context.Context, params *CreateAutoPredictorInput, optFns ...func(*Options)) (*CreateAutoPredictorOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateDatasetGroup(ctx context.Context, params *CreateDatasetGroupInput, optFns ...func(*Options)) (*CreateDatasetGroupOutput, error) 
 CreateDatasetImportJob(ctx context.Context, params *CreateDatasetImportJobInput, optFns ...func(*Options)) (*CreateDatasetImportJobOutput, error) 
 CreateExplainability(ctx context.Context, params *CreateExplainabilityInput, optFns ...func(*Options)) (*CreateExplainabilityOutput, error) 
 CreateExplainabilityExport(ctx context.Context, params *CreateExplainabilityExportInput, optFns ...func(*Options)) (*CreateExplainabilityExportOutput, error) 
 CreateForecast(ctx context.Context, params *CreateForecastInput, optFns ...func(*Options)) (*CreateForecastOutput, error) 
 CreateForecastExportJob(ctx context.Context, params *CreateForecastExportJobInput, optFns ...func(*Options)) (*CreateForecastExportJobOutput, error) 
 CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) 
 CreatePredictor(ctx context.Context, params *CreatePredictorInput, optFns ...func(*Options)) (*CreatePredictorOutput, error) 
 CreatePredictorBacktestExportJob(ctx context.Context, params *CreatePredictorBacktestExportJobInput, optFns ...func(*Options)) (*CreatePredictorBacktestExportJobOutput, error) 
 CreateWhatIfAnalysis(ctx context.Context, params *CreateWhatIfAnalysisInput, optFns ...func(*Options)) (*CreateWhatIfAnalysisOutput, error) 
 CreateWhatIfForecast(ctx context.Context, params *CreateWhatIfForecastInput, optFns ...func(*Options)) (*CreateWhatIfForecastOutput, error) 
 CreateWhatIfForecastExport(ctx context.Context, params *CreateWhatIfForecastExportInput, optFns ...func(*Options)) (*CreateWhatIfForecastExportOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteDatasetGroup(ctx context.Context, params *DeleteDatasetGroupInput, optFns ...func(*Options)) (*DeleteDatasetGroupOutput, error) 
 DeleteDatasetImportJob(ctx context.Context, params *DeleteDatasetImportJobInput, optFns ...func(*Options)) (*DeleteDatasetImportJobOutput, error) 
 DeleteExplainability(ctx context.Context, params *DeleteExplainabilityInput, optFns ...func(*Options)) (*DeleteExplainabilityOutput, error) 
 DeleteExplainabilityExport(ctx context.Context, params *DeleteExplainabilityExportInput, optFns ...func(*Options)) (*DeleteExplainabilityExportOutput, error) 
 DeleteForecast(ctx context.Context, params *DeleteForecastInput, optFns ...func(*Options)) (*DeleteForecastOutput, error) 
 DeleteForecastExportJob(ctx context.Context, params *DeleteForecastExportJobInput, optFns ...func(*Options)) (*DeleteForecastExportJobOutput, error) 
 DeleteMonitor(ctx context.Context, params *DeleteMonitorInput, optFns ...func(*Options)) (*DeleteMonitorOutput, error) 
 DeletePredictor(ctx context.Context, params *DeletePredictorInput, optFns ...func(*Options)) (*DeletePredictorOutput, error) 
 DeletePredictorBacktestExportJob(ctx context.Context, params *DeletePredictorBacktestExportJobInput, optFns ...func(*Options)) (*DeletePredictorBacktestExportJobOutput, error) 
 DeleteResourceTree(ctx context.Context, params *DeleteResourceTreeInput, optFns ...func(*Options)) (*DeleteResourceTreeOutput, error) 
 DeleteWhatIfAnalysis(ctx context.Context, params *DeleteWhatIfAnalysisInput, optFns ...func(*Options)) (*DeleteWhatIfAnalysisOutput, error) 
 DeleteWhatIfForecast(ctx context.Context, params *DeleteWhatIfForecastInput, optFns ...func(*Options)) (*DeleteWhatIfForecastOutput, error) 
 DeleteWhatIfForecastExport(ctx context.Context, params *DeleteWhatIfForecastExportInput, optFns ...func(*Options)) (*DeleteWhatIfForecastExportOutput, error) 
 DescribeAutoPredictor(ctx context.Context, params *DescribeAutoPredictorInput, optFns ...func(*Options)) (*DescribeAutoPredictorOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeDatasetGroup(ctx context.Context, params *DescribeDatasetGroupInput, optFns ...func(*Options)) (*DescribeDatasetGroupOutput, error) 
 DescribeDatasetImportJob(ctx context.Context, params *DescribeDatasetImportJobInput, optFns ...func(*Options)) (*DescribeDatasetImportJobOutput, error) 
 DescribeExplainability(ctx context.Context, params *DescribeExplainabilityInput, optFns ...func(*Options)) (*DescribeExplainabilityOutput, error) 
 DescribeExplainabilityExport(ctx context.Context, params *DescribeExplainabilityExportInput, optFns ...func(*Options)) (*DescribeExplainabilityExportOutput, error) 
 DescribeForecast(ctx context.Context, params *DescribeForecastInput, optFns ...func(*Options)) (*DescribeForecastOutput, error) 
 DescribeForecastExportJob(ctx context.Context, params *DescribeForecastExportJobInput, optFns ...func(*Options)) (*DescribeForecastExportJobOutput, error) 
 DescribeMonitor(ctx context.Context, params *DescribeMonitorInput, optFns ...func(*Options)) (*DescribeMonitorOutput, error) 
 DescribePredictor(ctx context.Context, params *DescribePredictorInput, optFns ...func(*Options)) (*DescribePredictorOutput, error) 
 DescribePredictorBacktestExportJob(ctx context.Context, params *DescribePredictorBacktestExportJobInput, optFns ...func(*Options)) (*DescribePredictorBacktestExportJobOutput, error) 
 DescribeWhatIfAnalysis(ctx context.Context, params *DescribeWhatIfAnalysisInput, optFns ...func(*Options)) (*DescribeWhatIfAnalysisOutput, error) 
 DescribeWhatIfForecast(ctx context.Context, params *DescribeWhatIfForecastInput, optFns ...func(*Options)) (*DescribeWhatIfForecastOutput, error) 
 DescribeWhatIfForecastExport(ctx context.Context, params *DescribeWhatIfForecastExportInput, optFns ...func(*Options)) (*DescribeWhatIfForecastExportOutput, error) 
 GetAccuracyMetrics(ctx context.Context, params *GetAccuracyMetricsInput, optFns ...func(*Options)) (*GetAccuracyMetricsOutput, error) 
 ListDatasetGroups(ctx context.Context, params *ListDatasetGroupsInput, optFns ...func(*Options)) (*ListDatasetGroupsOutput, error) 
 ListDatasetImportJobs(ctx context.Context, params *ListDatasetImportJobsInput, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListExplainabilities(ctx context.Context, params *ListExplainabilitiesInput, optFns ...func(*Options)) (*ListExplainabilitiesOutput, error) 
 ListExplainabilityExports(ctx context.Context, params *ListExplainabilityExportsInput, optFns ...func(*Options)) (*ListExplainabilityExportsOutput, error) 
 ListForecastExportJobs(ctx context.Context, params *ListForecastExportJobsInput, optFns ...func(*Options)) (*ListForecastExportJobsOutput, error) 
 ListForecasts(ctx context.Context, params *ListForecastsInput, optFns ...func(*Options)) (*ListForecastsOutput, error) 
 ListMonitorEvaluations(ctx context.Context, params *ListMonitorEvaluationsInput, optFns ...func(*Options)) (*ListMonitorEvaluationsOutput, error) 
 ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) 
 ListPredictorBacktestExportJobs(ctx context.Context, params *ListPredictorBacktestExportJobsInput, optFns ...func(*Options)) (*ListPredictorBacktestExportJobsOutput, error) 
 ListPredictors(ctx context.Context, params *ListPredictorsInput, optFns ...func(*Options)) (*ListPredictorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWhatIfAnalyses(ctx context.Context, params *ListWhatIfAnalysesInput, optFns ...func(*Options)) (*ListWhatIfAnalysesOutput, error) 
 ListWhatIfForecastExports(ctx context.Context, params *ListWhatIfForecastExportsInput, optFns ...func(*Options)) (*ListWhatIfForecastExportsOutput, error) 
 ListWhatIfForecasts(ctx context.Context, params *ListWhatIfForecastsInput, optFns ...func(*Options)) (*ListWhatIfForecastsOutput, error) 
 ResumeResource(ctx context.Context, params *ResumeResourceInput, optFns ...func(*Options)) (*ResumeResourceOutput, error) 
 StopResource(ctx context.Context, params *StopResourceInput, optFns ...func(*Options)) (*StopResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDatasetGroup(ctx context.Context, params *UpdateDatasetGroupInput, optFns ...func(*Options)) (*UpdateDatasetGroupOutput, error) 
}
