
package inspectorscan

import (
    "github.com/aws/aws-sdk-go-v2/service/inspectorscan"
)

// IClient defines the interface for inspectorscan
type IClient interface {
 Options() Options 
 ScanSbom(ctx context.Context, params *ScanSbomInput, optFns ...func(*Options)) (*ScanSbomOutput, error) 
}
