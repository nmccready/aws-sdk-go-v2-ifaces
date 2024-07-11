
package internetmonitor

import (
    "github.com/aws/aws-sdk-go-v2/service/internetmonitor"
)

// IClient defines the interface for internetmonitor
type IClient interface {
 Options() Options 
 CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) 
 DeleteMonitor(ctx context.Context, params *DeleteMonitorInput, optFns ...func(*Options)) (*DeleteMonitorOutput, error) 
 GetHealthEvent(ctx context.Context, params *GetHealthEventInput, optFns ...func(*Options)) (*GetHealthEventOutput, error) 
 GetInternetEvent(ctx context.Context, params *GetInternetEventInput, optFns ...func(*Options)) (*GetInternetEventOutput, error) 
 GetMonitor(ctx context.Context, params *GetMonitorInput, optFns ...func(*Options)) (*GetMonitorOutput, error) 
 GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error) 
 GetQueryStatus(ctx context.Context, params *GetQueryStatusInput, optFns ...func(*Options)) (*GetQueryStatusOutput, error) 
 ListHealthEvents(ctx context.Context, params *ListHealthEventsInput, optFns ...func(*Options)) (*ListHealthEventsOutput, error) 
 ListInternetEvents(ctx context.Context, params *ListInternetEventsInput, optFns ...func(*Options)) (*ListInternetEventsOutput, error) 
 ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartQuery(ctx context.Context, params *StartQueryInput, optFns ...func(*Options)) (*StartQueryOutput, error) 
 StopQuery(ctx context.Context, params *StopQueryInput, optFns ...func(*Options)) (*StopQueryOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateMonitor(ctx context.Context, params *UpdateMonitorInput, optFns ...func(*Options)) (*UpdateMonitorOutput, error) 
}
