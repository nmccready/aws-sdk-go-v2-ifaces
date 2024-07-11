
package codegurusecurity

import (
    "github.com/aws/aws-sdk-go-v2/service/codegurusecurity"
)

// IClient defines the interface for codegurusecurity
type IClient interface {
 Options() Options 
 BatchGetFindings(ctx context.Context, params *BatchGetFindingsInput, optFns ...func(*Options)) (*BatchGetFindingsOutput, error) 
 CreateScan(ctx context.Context, params *CreateScanInput, optFns ...func(*Options)) (*CreateScanOutput, error) 
 CreateUploadUrl(ctx context.Context, params *CreateUploadUrlInput, optFns ...func(*Options)) (*CreateUploadUrlOutput, error) 
 GetAccountConfiguration(ctx context.Context, params *GetAccountConfigurationInput, optFns ...func(*Options)) (*GetAccountConfigurationOutput, error) 
 GetFindings(ctx context.Context, params *GetFindingsInput, optFns ...func(*Options)) (*GetFindingsOutput, error) 
 GetMetricsSummary(ctx context.Context, params *GetMetricsSummaryInput, optFns ...func(*Options)) (*GetMetricsSummaryOutput, error) 
 GetScan(ctx context.Context, params *GetScanInput, optFns ...func(*Options)) (*GetScanOutput, error) 
 ListFindingsMetrics(ctx context.Context, params *ListFindingsMetricsInput, optFns ...func(*Options)) (*ListFindingsMetricsOutput, error) 
 ListScans(ctx context.Context, params *ListScansInput, optFns ...func(*Options)) (*ListScansOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccountConfiguration(ctx context.Context, params *UpdateAccountConfigurationInput, optFns ...func(*Options)) (*UpdateAccountConfigurationOutput, error) 
}
