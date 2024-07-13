
package controlcatalog_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/controlcatalog"
)

// IClient defines the interface for controlcatalog
type IClient interface {
 Options() Options 
 ListCommonControls(ctx context.Context, params *ListCommonControlsInput, optFns ...func(*Options)) (*ListCommonControlsOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListObjectives(ctx context.Context, params *ListObjectivesInput, optFns ...func(*Options)) (*ListObjectivesOutput, error) 
}