
package pi

import (
    "github.com/aws/aws-sdk-go-v2/service/pi"
)

// IClient defines the interface for pi
type IClient interface {
 Options() Options 
 CreatePerformanceAnalysisReport(ctx context.Context, params *CreatePerformanceAnalysisReportInput, optFns ...func(*Options)) (*CreatePerformanceAnalysisReportOutput, error) 
 DeletePerformanceAnalysisReport(ctx context.Context, params *DeletePerformanceAnalysisReportInput, optFns ...func(*Options)) (*DeletePerformanceAnalysisReportOutput, error) 
 DescribeDimensionKeys(ctx context.Context, params *DescribeDimensionKeysInput, optFns ...func(*Options)) (*DescribeDimensionKeysOutput, error) 
 GetDimensionKeyDetails(ctx context.Context, params *GetDimensionKeyDetailsInput, optFns ...func(*Options)) (*GetDimensionKeyDetailsOutput, error) 
 GetPerformanceAnalysisReport(ctx context.Context, params *GetPerformanceAnalysisReportInput, optFns ...func(*Options)) (*GetPerformanceAnalysisReportOutput, error) 
 GetResourceMetadata(ctx context.Context, params *GetResourceMetadataInput, optFns ...func(*Options)) (*GetResourceMetadataOutput, error) 
 GetResourceMetrics(ctx context.Context, params *GetResourceMetricsInput, optFns ...func(*Options)) (*GetResourceMetricsOutput, error) 
 ListAvailableResourceDimensions(ctx context.Context, params *ListAvailableResourceDimensionsInput, optFns ...func(*Options)) (*ListAvailableResourceDimensionsOutput, error) 
 ListAvailableResourceMetrics(ctx context.Context, params *ListAvailableResourceMetricsInput, optFns ...func(*Options)) (*ListAvailableResourceMetricsOutput, error) 
 ListPerformanceAnalysisReports(ctx context.Context, params *ListPerformanceAnalysisReportsInput, optFns ...func(*Options)) (*ListPerformanceAnalysisReportsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
