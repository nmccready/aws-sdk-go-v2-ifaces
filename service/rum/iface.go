
package rum

import (
    "github.com/aws/aws-sdk-go-v2/service/rum"
)

// IClient defines the interface for rum
type IClient interface {
 Options() Options 
 BatchCreateRumMetricDefinitions(ctx context.Context, params *BatchCreateRumMetricDefinitionsInput, optFns ...func(*Options)) (*BatchCreateRumMetricDefinitionsOutput, error) 
 BatchDeleteRumMetricDefinitions(ctx context.Context, params *BatchDeleteRumMetricDefinitionsInput, optFns ...func(*Options)) (*BatchDeleteRumMetricDefinitionsOutput, error) 
 BatchGetRumMetricDefinitions(ctx context.Context, params *BatchGetRumMetricDefinitionsInput, optFns ...func(*Options)) (*BatchGetRumMetricDefinitionsOutput, error) 
 CreateAppMonitor(ctx context.Context, params *CreateAppMonitorInput, optFns ...func(*Options)) (*CreateAppMonitorOutput, error) 
 DeleteAppMonitor(ctx context.Context, params *DeleteAppMonitorInput, optFns ...func(*Options)) (*DeleteAppMonitorOutput, error) 
 DeleteRumMetricsDestination(ctx context.Context, params *DeleteRumMetricsDestinationInput, optFns ...func(*Options)) (*DeleteRumMetricsDestinationOutput, error) 
 GetAppMonitor(ctx context.Context, params *GetAppMonitorInput, optFns ...func(*Options)) (*GetAppMonitorOutput, error) 
 GetAppMonitorData(ctx context.Context, params *GetAppMonitorDataInput, optFns ...func(*Options)) (*GetAppMonitorDataOutput, error) 
 ListAppMonitors(ctx context.Context, params *ListAppMonitorsInput, optFns ...func(*Options)) (*ListAppMonitorsOutput, error) 
 ListRumMetricsDestinations(ctx context.Context, params *ListRumMetricsDestinationsInput, optFns ...func(*Options)) (*ListRumMetricsDestinationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutRumEvents(ctx context.Context, params *PutRumEventsInput, optFns ...func(*Options)) (*PutRumEventsOutput, error) 
 PutRumMetricsDestination(ctx context.Context, params *PutRumMetricsDestinationInput, optFns ...func(*Options)) (*PutRumMetricsDestinationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAppMonitor(ctx context.Context, params *UpdateAppMonitorInput, optFns ...func(*Options)) (*UpdateAppMonitorOutput, error) 
 UpdateRumMetricDefinition(ctx context.Context, params *UpdateRumMetricDefinitionInput, optFns ...func(*Options)) (*UpdateRumMetricDefinitionOutput, error) 
}
