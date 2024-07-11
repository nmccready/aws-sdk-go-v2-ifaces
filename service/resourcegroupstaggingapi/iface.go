
package resourcegroupstaggingapi

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

// IClient defines the interface for resourcegroupstaggingapi
type IClient interface {
 Options() Options 
 DescribeReportCreation(ctx context.Context, params *DescribeReportCreationInput, optFns ...func(*Options)) (*DescribeReportCreationOutput, error) 
 GetComplianceSummary(ctx context.Context, params *GetComplianceSummaryInput, optFns ...func(*Options)) (*GetComplianceSummaryOutput, error) 
 GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) 
 GetTagKeys(ctx context.Context, params *GetTagKeysInput, optFns ...func(*Options)) (*GetTagKeysOutput, error) 
 GetTagValues(ctx context.Context, params *GetTagValuesInput, optFns ...func(*Options)) (*GetTagValuesOutput, error) 
 StartReportCreation(ctx context.Context, params *StartReportCreationInput, optFns ...func(*Options)) (*StartReportCreationOutput, error) 
 TagResources(ctx context.Context, params *TagResourcesInput, optFns ...func(*Options)) (*TagResourcesOutput, error) 
 UntagResources(ctx context.Context, params *UntagResourcesInput, optFns ...func(*Options)) (*UntagResourcesOutput, error) 
}
