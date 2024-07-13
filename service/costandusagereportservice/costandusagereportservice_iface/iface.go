
package costandusagereportservice_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

// IClient defines the interface for costandusagereportservice
type IClient interface {
 Options() Options 
 DeleteReportDefinition(ctx context.Context, params *DeleteReportDefinitionInput, optFns ...func(*Options)) (*DeleteReportDefinitionOutput, error) 
 DescribeReportDefinitions(ctx context.Context, params *DescribeReportDefinitionsInput, optFns ...func(*Options)) (*DescribeReportDefinitionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyReportDefinition(ctx context.Context, params *ModifyReportDefinitionInput, optFns ...func(*Options)) (*ModifyReportDefinitionOutput, error) 
 PutReportDefinition(ctx context.Context, params *PutReportDefinitionInput, optFns ...func(*Options)) (*PutReportDefinitionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}