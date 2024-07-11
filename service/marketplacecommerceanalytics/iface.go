
package marketplacecommerceanalytics

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
)

// IClient defines the interface for marketplacecommerceanalytics
type IClient interface {
 Options() Options 
 GenerateDataSet(ctx context.Context, params *GenerateDataSetInput, optFns ...func(*Options)) (*GenerateDataSetOutput, error) 
 StartSupportDataExport(ctx context.Context, params *StartSupportDataExportInput, optFns ...func(*Options)) (*StartSupportDataExportOutput, error) 
}
