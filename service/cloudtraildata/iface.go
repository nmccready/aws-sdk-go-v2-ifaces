
package cloudtraildata

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
)

// IClient defines the interface for cloudtraildata
type IClient interface {
 Options() Options 
 PutAuditEvents(ctx context.Context, params *PutAuditEventsInput, optFns ...func(*Options)) (*PutAuditEventsOutput, error) 
}
