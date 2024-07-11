
package route53recoverycluster

import (
    "github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
)

// IRoute53recoverycluster defines the interface for route53recoverycluster
type IRoute53recoverycluster interface {
 Options() Options 
 GetRoutingControlState(ctx context.Context, params *GetRoutingControlStateInput, optFns ...func(*Options)) (*GetRoutingControlStateOutput, error) 
 ListRoutingControls(ctx context.Context, params *ListRoutingControlsInput, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) 
 UpdateRoutingControlState(ctx context.Context, params *UpdateRoutingControlStateInput, optFns ...func(*Options)) (*UpdateRoutingControlStateOutput, error) 
 UpdateRoutingControlStates(ctx context.Context, params *UpdateRoutingControlStatesInput, optFns ...func(*Options)) (*UpdateRoutingControlStatesOutput, error) 
}
