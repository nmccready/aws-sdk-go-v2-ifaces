
package xray_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/xray"
)

// IClient defines the interface for xray
type IClient interface {
 Options() Options 
 BatchGetTraces(ctx context.Context, params *BatchGetTracesInput, optFns ...func(*Options)) (*BatchGetTracesOutput, error) 
 CancelTraceRetrieval(ctx context.Context, params *CancelTraceRetrievalInput, optFns ...func(*Options)) (*CancelTraceRetrievalOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateSamplingRule(ctx context.Context, params *CreateSamplingRuleInput, optFns ...func(*Options)) (*CreateSamplingRuleOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteSamplingRule(ctx context.Context, params *DeleteSamplingRuleInput, optFns ...func(*Options)) (*DeleteSamplingRuleOutput, error) 
 GetEncryptionConfig(ctx context.Context, params *GetEncryptionConfigInput, optFns ...func(*Options)) (*GetEncryptionConfigOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetGroups(ctx context.Context, params *GetGroupsInput, optFns ...func(*Options)) (*GetGroupsOutput, error) 
 GetIndexingRules(ctx context.Context, params *GetIndexingRulesInput, optFns ...func(*Options)) (*GetIndexingRulesOutput, error) 
 GetInsight(ctx context.Context, params *GetInsightInput, optFns ...func(*Options)) (*GetInsightOutput, error) 
 GetInsightEvents(ctx context.Context, params *GetInsightEventsInput, optFns ...func(*Options)) (*GetInsightEventsOutput, error) 
 GetInsightImpactGraph(ctx context.Context, params *GetInsightImpactGraphInput, optFns ...func(*Options)) (*GetInsightImpactGraphOutput, error) 
 GetInsightSummaries(ctx context.Context, params *GetInsightSummariesInput, optFns ...func(*Options)) (*GetInsightSummariesOutput, error) 
 GetRetrievedTracesGraph(ctx context.Context, params *GetRetrievedTracesGraphInput, optFns ...func(*Options)) (*GetRetrievedTracesGraphOutput, error) 
 GetSamplingRules(ctx context.Context, params *GetSamplingRulesInput, optFns ...func(*Options)) (*GetSamplingRulesOutput, error) 
 GetSamplingStatisticSummaries(ctx context.Context, params *GetSamplingStatisticSummariesInput, optFns ...func(*Options)) (*GetSamplingStatisticSummariesOutput, error) 
 GetSamplingTargets(ctx context.Context, params *GetSamplingTargetsInput, optFns ...func(*Options)) (*GetSamplingTargetsOutput, error) 
 GetServiceGraph(ctx context.Context, params *GetServiceGraphInput, optFns ...func(*Options)) (*GetServiceGraphOutput, error) 
 GetTimeSeriesServiceStatistics(ctx context.Context, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) 
 GetTraceGraph(ctx context.Context, params *GetTraceGraphInput, optFns ...func(*Options)) (*GetTraceGraphOutput, error) 
 GetTraceSegmentDestination(ctx context.Context, params *GetTraceSegmentDestinationInput, optFns ...func(*Options)) (*GetTraceSegmentDestinationOutput, error) 
 GetTraceSummaries(ctx context.Context, params *GetTraceSummariesInput, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) 
 ListResourcePolicies(ctx context.Context, params *ListResourcePoliciesInput, optFns ...func(*Options)) (*ListResourcePoliciesOutput, error) 
 ListRetrievedTraces(ctx context.Context, params *ListRetrievedTracesInput, optFns ...func(*Options)) (*ListRetrievedTracesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutEncryptionConfig(ctx context.Context, params *PutEncryptionConfigInput, optFns ...func(*Options)) (*PutEncryptionConfigOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 PutTelemetryRecords(ctx context.Context, params *PutTelemetryRecordsInput, optFns ...func(*Options)) (*PutTelemetryRecordsOutput, error) 
 PutTraceSegments(ctx context.Context, params *PutTraceSegmentsInput, optFns ...func(*Options)) (*PutTraceSegmentsOutput, error) 
 StartTraceRetrieval(ctx context.Context, params *StartTraceRetrievalInput, optFns ...func(*Options)) (*StartTraceRetrievalOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateIndexingRule(ctx context.Context, params *UpdateIndexingRuleInput, optFns ...func(*Options)) (*UpdateIndexingRuleOutput, error) 
 UpdateSamplingRule(ctx context.Context, params *UpdateSamplingRuleInput, optFns ...func(*Options)) (*UpdateSamplingRuleOutput, error) 
 UpdateTraceSegmentDestination(ctx context.Context, params *UpdateTraceSegmentDestinationInput, optFns ...func(*Options)) (*UpdateTraceSegmentDestinationOutput, error) 
}
