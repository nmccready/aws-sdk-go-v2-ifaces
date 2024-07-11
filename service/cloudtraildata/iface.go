
package cloudtraildata

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
)

// ICloudtraildata defines the interface for cloudtraildata
type ICloudtraildata interface {
 Options() Options 
 PutAuditEvents(ctx context.Context, params *PutAuditEventsInput, optFns ...func(*Options)) (*PutAuditEventsOutput, error) 
}
