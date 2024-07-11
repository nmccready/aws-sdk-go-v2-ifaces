
package networkmonitor

import (
    "github.com/aws/aws-sdk-go-v2/service/networkmonitor"
)

// INetworkmonitor defines the interface for networkmonitor
type INetworkmonitor interface {
 Options() Options 
 CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) 
 CreateProbe(ctx context.Context, params *CreateProbeInput, optFns ...func(*Options)) (*CreateProbeOutput, error) 
 DeleteMonitor(ctx context.Context, params *DeleteMonitorInput, optFns ...func(*Options)) (*DeleteMonitorOutput, error) 
 DeleteProbe(ctx context.Context, params *DeleteProbeInput, optFns ...func(*Options)) (*DeleteProbeOutput, error) 
 GetMonitor(ctx context.Context, params *GetMonitorInput, optFns ...func(*Options)) (*GetMonitorOutput, error) 
 GetProbe(ctx context.Context, params *GetProbeInput, optFns ...func(*Options)) (*GetProbeOutput, error) 
 ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateMonitor(ctx context.Context, params *UpdateMonitorInput, optFns ...func(*Options)) (*UpdateMonitorOutput, error) 
 UpdateProbe(ctx context.Context, params *UpdateProbeInput, optFns ...func(*Options)) (*UpdateProbeOutput, error) 
}
