
package cloudwatch

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// IClient defines the interface for cloudwatch
type IClient interface {
 Options() Options 
 DeleteAlarms(ctx context.Context, params *DeleteAlarmsInput, optFns ...func(*Options)) (*DeleteAlarmsOutput, error) 
 DeleteAnomalyDetector(ctx context.Context, params *DeleteAnomalyDetectorInput, optFns ...func(*Options)) (*DeleteAnomalyDetectorOutput, error) 
 DeleteDashboards(ctx context.Context, params *DeleteDashboardsInput, optFns ...func(*Options)) (*DeleteDashboardsOutput, error) 
 DeleteInsightRules(ctx context.Context, params *DeleteInsightRulesInput, optFns ...func(*Options)) (*DeleteInsightRulesOutput, error) 
 DeleteMetricStream(ctx context.Context, params *DeleteMetricStreamInput, optFns ...func(*Options)) (*DeleteMetricStreamOutput, error) 
 DescribeAlarmHistory(ctx context.Context, params *DescribeAlarmHistoryInput, optFns ...func(*Options)) (*DescribeAlarmHistoryOutput, error) 
 DescribeAlarms(ctx context.Context, params *DescribeAlarmsInput, optFns ...func(*Options)) (*DescribeAlarmsOutput, error) 
 DescribeAlarmsForMetric(ctx context.Context, params *DescribeAlarmsForMetricInput, optFns ...func(*Options)) (*DescribeAlarmsForMetricOutput, error) 
 DescribeAnomalyDetectors(ctx context.Context, params *DescribeAnomalyDetectorsInput, optFns ...func(*Options)) (*DescribeAnomalyDetectorsOutput, error) 
 DescribeInsightRules(ctx context.Context, params *DescribeInsightRulesInput, optFns ...func(*Options)) (*DescribeInsightRulesOutput, error) 
 DisableAlarmActions(ctx context.Context, params *DisableAlarmActionsInput, optFns ...func(*Options)) (*DisableAlarmActionsOutput, error) 
 DisableInsightRules(ctx context.Context, params *DisableInsightRulesInput, optFns ...func(*Options)) (*DisableInsightRulesOutput, error) 
 EnableAlarmActions(ctx context.Context, params *EnableAlarmActionsInput, optFns ...func(*Options)) (*EnableAlarmActionsOutput, error) 
 EnableInsightRules(ctx context.Context, params *EnableInsightRulesInput, optFns ...func(*Options)) (*EnableInsightRulesOutput, error) 
 GetDashboard(ctx context.Context, params *GetDashboardInput, optFns ...func(*Options)) (*GetDashboardOutput, error) 
 GetInsightRuleReport(ctx context.Context, params *GetInsightRuleReportInput, optFns ...func(*Options)) (*GetInsightRuleReportOutput, error) 
 GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) 
 GetMetricStatistics(ctx context.Context, params *GetMetricStatisticsInput, optFns ...func(*Options)) (*GetMetricStatisticsOutput, error) 
 GetMetricStream(ctx context.Context, params *GetMetricStreamInput, optFns ...func(*Options)) (*GetMetricStreamOutput, error) 
 GetMetricWidgetImage(ctx context.Context, params *GetMetricWidgetImageInput, optFns ...func(*Options)) (*GetMetricWidgetImageOutput, error) 
 ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) 
 ListManagedInsightRules(ctx context.Context, params *ListManagedInsightRulesInput, optFns ...func(*Options)) (*ListManagedInsightRulesOutput, error) 
 ListMetricStreams(ctx context.Context, params *ListMetricStreamsInput, optFns ...func(*Options)) (*ListMetricStreamsOutput, error) 
 ListMetrics(ctx context.Context, params *ListMetricsInput, optFns ...func(*Options)) (*ListMetricsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutAnomalyDetector(ctx context.Context, params *PutAnomalyDetectorInput, optFns ...func(*Options)) (*PutAnomalyDetectorOutput, error) 
 PutCompositeAlarm(ctx context.Context, params *PutCompositeAlarmInput, optFns ...func(*Options)) (*PutCompositeAlarmOutput, error) 
 PutDashboard(ctx context.Context, params *PutDashboardInput, optFns ...func(*Options)) (*PutDashboardOutput, error) 
 PutInsightRule(ctx context.Context, params *PutInsightRuleInput, optFns ...func(*Options)) (*PutInsightRuleOutput, error) 
 PutManagedInsightRules(ctx context.Context, params *PutManagedInsightRulesInput, optFns ...func(*Options)) (*PutManagedInsightRulesOutput, error) 
 PutMetricAlarm(ctx context.Context, params *PutMetricAlarmInput, optFns ...func(*Options)) (*PutMetricAlarmOutput, error) 
 PutMetricData(ctx context.Context, params *PutMetricDataInput, optFns ...func(*Options)) (*PutMetricDataOutput, error) 
 PutMetricStream(ctx context.Context, params *PutMetricStreamInput, optFns ...func(*Options)) (*PutMetricStreamOutput, error) 
 SetAlarmState(ctx context.Context, params *SetAlarmStateInput, optFns ...func(*Options)) (*SetAlarmStateOutput, error) 
 StartMetricStreams(ctx context.Context, params *StartMetricStreamsInput, optFns ...func(*Options)) (*StartMetricStreamsOutput, error) 
 StopMetricStreams(ctx context.Context, params *StopMetricStreamsInput, optFns ...func(*Options)) (*StopMetricStreamsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
