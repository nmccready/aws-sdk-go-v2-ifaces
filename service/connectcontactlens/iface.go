
package connectcontactlens

import (
    "github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
)

// IClient defines the interface for connectcontactlens
type IClient interface {
 Options() Options 
 ListRealtimeContactAnalysisSegments(ctx context.Context, params *ListRealtimeContactAnalysisSegmentsInput, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsOutput, error) 
}
