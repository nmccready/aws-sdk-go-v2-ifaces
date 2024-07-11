
package inspectorscan

import (
    "github.com/aws/aws-sdk-go-v2/service/inspectorscan"
)

// IInspectorscan defines the interface for inspectorscan
type IInspectorscan interface {
 Options() Options 
 ScanSbom(ctx context.Context, params *ScanSbomInput, optFns ...func(*Options)) (*ScanSbomOutput, error) 
}
