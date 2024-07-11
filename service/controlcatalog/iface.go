
package controlcatalog

import (
    "github.com/aws/aws-sdk-go-v2/service/controlcatalog"
)

// IControlcatalog defines the interface for controlcatalog
type IControlcatalog interface {
 Options() Options 
 ListCommonControls(ctx context.Context, params *ListCommonControlsInput, optFns ...func(*Options)) (*ListCommonControlsOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListObjectives(ctx context.Context, params *ListObjectivesInput, optFns ...func(*Options)) (*ListObjectivesOutput, error) 
}
