
package sagemakera2iruntime_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
)

// IClient defines the interface for sagemakera2iruntime
type IClient interface {
 Options() Options 
 DeleteHumanLoop(ctx context.Context, params *DeleteHumanLoopInput, optFns ...func(*Options)) (*DeleteHumanLoopOutput, error) 
 DescribeHumanLoop(ctx context.Context, params *DescribeHumanLoopInput, optFns ...func(*Options)) (*DescribeHumanLoopOutput, error) 
 ListHumanLoops(ctx context.Context, params *ListHumanLoopsInput, optFns ...func(*Options)) (*ListHumanLoopsOutput, error) 
 StartHumanLoop(ctx context.Context, params *StartHumanLoopInput, optFns ...func(*Options)) (*StartHumanLoopOutput, error) 
 StopHumanLoop(ctx context.Context, params *StopHumanLoopInput, optFns ...func(*Options)) (*StopHumanLoopOutput, error) 
}