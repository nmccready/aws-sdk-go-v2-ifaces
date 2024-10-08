
package appflow_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/appflow"
)

// IClient defines the interface for appflow
type IClient interface {
 Options() Options 
 CancelFlowExecutions(ctx context.Context, params *CancelFlowExecutionsInput, optFns ...func(*Options)) (*CancelFlowExecutionsOutput, error) 
 CreateConnectorProfile(ctx context.Context, params *CreateConnectorProfileInput, optFns ...func(*Options)) (*CreateConnectorProfileOutput, error) 
 CreateFlow(ctx context.Context, params *CreateFlowInput, optFns ...func(*Options)) (*CreateFlowOutput, error) 
 DeleteConnectorProfile(ctx context.Context, params *DeleteConnectorProfileInput, optFns ...func(*Options)) (*DeleteConnectorProfileOutput, error) 
 DeleteFlow(ctx context.Context, params *DeleteFlowInput, optFns ...func(*Options)) (*DeleteFlowOutput, error) 
 DescribeConnector(ctx context.Context, params *DescribeConnectorInput, optFns ...func(*Options)) (*DescribeConnectorOutput, error) 
 DescribeConnectorEntity(ctx context.Context, params *DescribeConnectorEntityInput, optFns ...func(*Options)) (*DescribeConnectorEntityOutput, error) 
 DescribeConnectorProfiles(ctx context.Context, params *DescribeConnectorProfilesInput, optFns ...func(*Options)) (*DescribeConnectorProfilesOutput, error) 
 DescribeConnectors(ctx context.Context, params *DescribeConnectorsInput, optFns ...func(*Options)) (*DescribeConnectorsOutput, error) 
 DescribeFlow(ctx context.Context, params *DescribeFlowInput, optFns ...func(*Options)) (*DescribeFlowOutput, error) 
 DescribeFlowExecutionRecords(ctx context.Context, params *DescribeFlowExecutionRecordsInput, optFns ...func(*Options)) (*DescribeFlowExecutionRecordsOutput, error) 
 ListConnectorEntities(ctx context.Context, params *ListConnectorEntitiesInput, optFns ...func(*Options)) (*ListConnectorEntitiesOutput, error) 
 ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) 
 ListFlows(ctx context.Context, params *ListFlowsInput, optFns ...func(*Options)) (*ListFlowsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterConnector(ctx context.Context, params *RegisterConnectorInput, optFns ...func(*Options)) (*RegisterConnectorOutput, error) 
 ResetConnectorMetadataCache(ctx context.Context, params *ResetConnectorMetadataCacheInput, optFns ...func(*Options)) (*ResetConnectorMetadataCacheOutput, error) 
 StartFlow(ctx context.Context, params *StartFlowInput, optFns ...func(*Options)) (*StartFlowOutput, error) 
 StopFlow(ctx context.Context, params *StopFlowInput, optFns ...func(*Options)) (*StopFlowOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UnregisterConnector(ctx context.Context, params *UnregisterConnectorInput, optFns ...func(*Options)) (*UnregisterConnectorOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConnectorProfile(ctx context.Context, params *UpdateConnectorProfileInput, optFns ...func(*Options)) (*UpdateConnectorProfileOutput, error) 
 UpdateConnectorRegistration(ctx context.Context, params *UpdateConnectorRegistrationInput, optFns ...func(*Options)) (*UpdateConnectorRegistrationOutput, error) 
 UpdateFlow(ctx context.Context, params *UpdateFlowInput, optFns ...func(*Options)) (*UpdateFlowOutput, error) 
}
