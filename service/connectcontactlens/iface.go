
package connectcontactlens

import (
    "github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
)

// IConnectcontactlens defines the interface for connectcontactlens
type IConnectcontactlens interface {
 Options() Options 
 ListRealtimeContactAnalysisSegments(ctx context.Context, params *ListRealtimeContactAnalysisSegmentsInput, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsOutput, error) 
}
