
package lookoutmetrics_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/lookoutmetrics"
)

// IClient defines the interface for lookoutmetrics
type IClient interface {
 Options() Options 
 ActivateAnomalyDetector(ctx context.Context, params *ActivateAnomalyDetectorInput, optFns ...func(*Options)) (*ActivateAnomalyDetectorOutput, error) 
 BackTestAnomalyDetector(ctx context.Context, params *BackTestAnomalyDetectorInput, optFns ...func(*Options)) (*BackTestAnomalyDetectorOutput, error) 
 CreateAlert(ctx context.Context, params *CreateAlertInput, optFns ...func(*Options)) (*CreateAlertOutput, error) 
 CreateAnomalyDetector(ctx context.Context, params *CreateAnomalyDetectorInput, optFns ...func(*Options)) (*CreateAnomalyDetectorOutput, error) 
 CreateMetricSet(ctx context.Context, params *CreateMetricSetInput, optFns ...func(*Options)) (*CreateMetricSetOutput, error) 
 DeactivateAnomalyDetector(ctx context.Context, params *DeactivateAnomalyDetectorInput, optFns ...func(*Options)) (*DeactivateAnomalyDetectorOutput, error) 
 DeleteAlert(ctx context.Context, params *DeleteAlertInput, optFns ...func(*Options)) (*DeleteAlertOutput, error) 
 DeleteAnomalyDetector(ctx context.Context, params *DeleteAnomalyDetectorInput, optFns ...func(*Options)) (*DeleteAnomalyDetectorOutput, error) 
 DescribeAlert(ctx context.Context, params *DescribeAlertInput, optFns ...func(*Options)) (*DescribeAlertOutput, error) 
 DescribeAnomalyDetectionExecutions(ctx context.Context, params *DescribeAnomalyDetectionExecutionsInput, optFns ...func(*Options)) (*DescribeAnomalyDetectionExecutionsOutput, error) 
 DescribeAnomalyDetector(ctx context.Context, params *DescribeAnomalyDetectorInput, optFns ...func(*Options)) (*DescribeAnomalyDetectorOutput, error) 
 DescribeMetricSet(ctx context.Context, params *DescribeMetricSetInput, optFns ...func(*Options)) (*DescribeMetricSetOutput, error) 
 DetectMetricSetConfig(ctx context.Context, params *DetectMetricSetConfigInput, optFns ...func(*Options)) (*DetectMetricSetConfigOutput, error) 
 GetAnomalyGroup(ctx context.Context, params *GetAnomalyGroupInput, optFns ...func(*Options)) (*GetAnomalyGroupOutput, error) 
 GetDataQualityMetrics(ctx context.Context, params *GetDataQualityMetricsInput, optFns ...func(*Options)) (*GetDataQualityMetricsOutput, error) 
 GetFeedback(ctx context.Context, params *GetFeedbackInput, optFns ...func(*Options)) (*GetFeedbackOutput, error) 
 GetSampleData(ctx context.Context, params *GetSampleDataInput, optFns ...func(*Options)) (*GetSampleDataOutput, error) 
 ListAlerts(ctx context.Context, params *ListAlertsInput, optFns ...func(*Options)) (*ListAlertsOutput, error) 
 ListAnomalyDetectors(ctx context.Context, params *ListAnomalyDetectorsInput, optFns ...func(*Options)) (*ListAnomalyDetectorsOutput, error) 
 ListAnomalyGroupRelatedMetrics(ctx context.Context, params *ListAnomalyGroupRelatedMetricsInput, optFns ...func(*Options)) (*ListAnomalyGroupRelatedMetricsOutput, error) 
 ListAnomalyGroupSummaries(ctx context.Context, params *ListAnomalyGroupSummariesInput, optFns ...func(*Options)) (*ListAnomalyGroupSummariesOutput, error) 
 ListAnomalyGroupTimeSeries(ctx context.Context, params *ListAnomalyGroupTimeSeriesInput, optFns ...func(*Options)) (*ListAnomalyGroupTimeSeriesOutput, error) 
 ListMetricSets(ctx context.Context, params *ListMetricSetsInput, optFns ...func(*Options)) (*ListMetricSetsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutFeedback(ctx context.Context, params *PutFeedbackInput, optFns ...func(*Options)) (*PutFeedbackOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAlert(ctx context.Context, params *UpdateAlertInput, optFns ...func(*Options)) (*UpdateAlertOutput, error) 
 UpdateAnomalyDetector(ctx context.Context, params *UpdateAnomalyDetectorInput, optFns ...func(*Options)) (*UpdateAnomalyDetectorOutput, error) 
 UpdateMetricSet(ctx context.Context, params *UpdateMetricSetInput, optFns ...func(*Options)) (*UpdateMetricSetOutput, error) 
}
