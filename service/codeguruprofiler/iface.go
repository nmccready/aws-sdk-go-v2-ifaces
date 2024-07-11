
package codeguruprofiler

import (
    "github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
)

// ICodeguruprofiler defines the interface for codeguruprofiler
type ICodeguruprofiler interface {
 Options() Options 
 AddNotificationChannels(ctx context.Context, params *AddNotificationChannelsInput, optFns ...func(*Options)) (*AddNotificationChannelsOutput, error) 
 BatchGetFrameMetricData(ctx context.Context, params *BatchGetFrameMetricDataInput, optFns ...func(*Options)) (*BatchGetFrameMetricDataOutput, error) 
 ConfigureAgent(ctx context.Context, params *ConfigureAgentInput, optFns ...func(*Options)) (*ConfigureAgentOutput, error) 
 CreateProfilingGroup(ctx context.Context, params *CreateProfilingGroupInput, optFns ...func(*Options)) (*CreateProfilingGroupOutput, error) 
 DeleteProfilingGroup(ctx context.Context, params *DeleteProfilingGroupInput, optFns ...func(*Options)) (*DeleteProfilingGroupOutput, error) 
 DescribeProfilingGroup(ctx context.Context, params *DescribeProfilingGroupInput, optFns ...func(*Options)) (*DescribeProfilingGroupOutput, error) 
 GetFindingsReportAccountSummary(ctx context.Context, params *GetFindingsReportAccountSummaryInput, optFns ...func(*Options)) (*GetFindingsReportAccountSummaryOutput, error) 
 GetNotificationConfiguration(ctx context.Context, params *GetNotificationConfigurationInput, optFns ...func(*Options)) (*GetNotificationConfigurationOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) 
 GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) 
 ListFindingsReports(ctx context.Context, params *ListFindingsReportsInput, optFns ...func(*Options)) (*ListFindingsReportsOutput, error) 
 ListProfileTimes(ctx context.Context, params *ListProfileTimesInput, optFns ...func(*Options)) (*ListProfileTimesOutput, error) 
 ListProfilingGroups(ctx context.Context, params *ListProfilingGroupsInput, optFns ...func(*Options)) (*ListProfilingGroupsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PostAgentProfile(ctx context.Context, params *PostAgentProfileInput, optFns ...func(*Options)) (*PostAgentProfileOutput, error) 
 PutPermission(ctx context.Context, params *PutPermissionInput, optFns ...func(*Options)) (*PutPermissionOutput, error) 
 RemoveNotificationChannel(ctx context.Context, params *RemoveNotificationChannelInput, optFns ...func(*Options)) (*RemoveNotificationChannelOutput, error) 
 RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) 
 SubmitFeedback(ctx context.Context, params *SubmitFeedbackInput, optFns ...func(*Options)) (*SubmitFeedbackOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateProfilingGroup(ctx context.Context, params *UpdateProfilingGroupInput, optFns ...func(*Options)) (*UpdateProfilingGroupOutput, error) 
}
