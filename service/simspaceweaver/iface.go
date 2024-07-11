
package simspaceweaver

import (
    "github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
)

// ISimspaceweaver defines the interface for simspaceweaver
type ISimspaceweaver interface {
 Options() Options 
 CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) 
 DeleteApp(ctx context.Context, params *DeleteAppInput, optFns ...func(*Options)) (*DeleteAppOutput, error) 
 DeleteSimulation(ctx context.Context, params *DeleteSimulationInput, optFns ...func(*Options)) (*DeleteSimulationOutput, error) 
 DescribeApp(ctx context.Context, params *DescribeAppInput, optFns ...func(*Options)) (*DescribeAppOutput, error) 
 DescribeSimulation(ctx context.Context, params *DescribeSimulationInput, optFns ...func(*Options)) (*DescribeSimulationOutput, error) 
 ListApps(ctx context.Context, params *ListAppsInput, optFns ...func(*Options)) (*ListAppsOutput, error) 
 ListSimulations(ctx context.Context, params *ListSimulationsInput, optFns ...func(*Options)) (*ListSimulationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartApp(ctx context.Context, params *StartAppInput, optFns ...func(*Options)) (*StartAppOutput, error) 
 StartClock(ctx context.Context, params *StartClockInput, optFns ...func(*Options)) (*StartClockOutput, error) 
 StartSimulation(ctx context.Context, params *StartSimulationInput, optFns ...func(*Options)) (*StartSimulationOutput, error) 
 StopApp(ctx context.Context, params *StopAppInput, optFns ...func(*Options)) (*StopAppOutput, error) 
 StopClock(ctx context.Context, params *StopClockInput, optFns ...func(*Options)) (*StopClockOutput, error) 
 StopSimulation(ctx context.Context, params *StopSimulationInput, optFns ...func(*Options)) (*StopSimulationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
