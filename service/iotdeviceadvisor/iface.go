
package iotdeviceadvisor

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
)

// IClient defines the interface for iotdeviceadvisor
type IClient interface {
 Options() Options 
 CreateSuiteDefinition(ctx context.Context, params *CreateSuiteDefinitionInput, optFns ...func(*Options)) (*CreateSuiteDefinitionOutput, error) 
 DeleteSuiteDefinition(ctx context.Context, params *DeleteSuiteDefinitionInput, optFns ...func(*Options)) (*DeleteSuiteDefinitionOutput, error) 
 GetEndpoint(ctx context.Context, params *GetEndpointInput, optFns ...func(*Options)) (*GetEndpointOutput, error) 
 GetSuiteDefinition(ctx context.Context, params *GetSuiteDefinitionInput, optFns ...func(*Options)) (*GetSuiteDefinitionOutput, error) 
 GetSuiteRun(ctx context.Context, params *GetSuiteRunInput, optFns ...func(*Options)) (*GetSuiteRunOutput, error) 
 GetSuiteRunReport(ctx context.Context, params *GetSuiteRunReportInput, optFns ...func(*Options)) (*GetSuiteRunReportOutput, error) 
 ListSuiteDefinitions(ctx context.Context, params *ListSuiteDefinitionsInput, optFns ...func(*Options)) (*ListSuiteDefinitionsOutput, error) 
 ListSuiteRuns(ctx context.Context, params *ListSuiteRunsInput, optFns ...func(*Options)) (*ListSuiteRunsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartSuiteRun(ctx context.Context, params *StartSuiteRunInput, optFns ...func(*Options)) (*StartSuiteRunOutput, error) 
 StopSuiteRun(ctx context.Context, params *StopSuiteRunInput, optFns ...func(*Options)) (*StopSuiteRunOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateSuiteDefinition(ctx context.Context, params *UpdateSuiteDefinitionInput, optFns ...func(*Options)) (*UpdateSuiteDefinitionOutput, error) 
}
