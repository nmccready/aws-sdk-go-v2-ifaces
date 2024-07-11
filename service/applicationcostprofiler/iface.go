
package applicationcostprofiler

import (
    "github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
)

// IClient defines the interface for applicationcostprofiler
type IClient interface {
 Options() Options 
 DeleteReportDefinition(ctx context.Context, params *DeleteReportDefinitionInput, optFns ...func(*Options)) (*DeleteReportDefinitionOutput, error) 
 GetReportDefinition(ctx context.Context, params *GetReportDefinitionInput, optFns ...func(*Options)) (*GetReportDefinitionOutput, error) 
 ImportApplicationUsage(ctx context.Context, params *ImportApplicationUsageInput, optFns ...func(*Options)) (*ImportApplicationUsageOutput, error) 
 ListReportDefinitions(ctx context.Context, params *ListReportDefinitionsInput, optFns ...func(*Options)) (*ListReportDefinitionsOutput, error) 
 PutReportDefinition(ctx context.Context, params *PutReportDefinitionInput, optFns ...func(*Options)) (*PutReportDefinitionOutput, error) 
 UpdateReportDefinition(ctx context.Context, params *UpdateReportDefinitionInput, optFns ...func(*Options)) (*UpdateReportDefinitionOutput, error) 
}
