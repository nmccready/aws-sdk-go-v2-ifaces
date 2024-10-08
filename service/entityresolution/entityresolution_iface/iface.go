
package entityresolution_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/entityresolution"
)

// IClient defines the interface for entityresolution
type IClient interface {
 Options() Options 
 AddPolicyStatement(ctx context.Context, params *AddPolicyStatementInput, optFns ...func(*Options)) (*AddPolicyStatementOutput, error) 
 BatchDeleteUniqueId(ctx context.Context, params *BatchDeleteUniqueIdInput, optFns ...func(*Options)) (*BatchDeleteUniqueIdOutput, error) 
 CreateIdMappingWorkflow(ctx context.Context, params *CreateIdMappingWorkflowInput, optFns ...func(*Options)) (*CreateIdMappingWorkflowOutput, error) 
 CreateIdNamespace(ctx context.Context, params *CreateIdNamespaceInput, optFns ...func(*Options)) (*CreateIdNamespaceOutput, error) 
 CreateMatchingWorkflow(ctx context.Context, params *CreateMatchingWorkflowInput, optFns ...func(*Options)) (*CreateMatchingWorkflowOutput, error) 
 CreateSchemaMapping(ctx context.Context, params *CreateSchemaMappingInput, optFns ...func(*Options)) (*CreateSchemaMappingOutput, error) 
 DeleteIdMappingWorkflow(ctx context.Context, params *DeleteIdMappingWorkflowInput, optFns ...func(*Options)) (*DeleteIdMappingWorkflowOutput, error) 
 DeleteIdNamespace(ctx context.Context, params *DeleteIdNamespaceInput, optFns ...func(*Options)) (*DeleteIdNamespaceOutput, error) 
 DeleteMatchingWorkflow(ctx context.Context, params *DeleteMatchingWorkflowInput, optFns ...func(*Options)) (*DeleteMatchingWorkflowOutput, error) 
 DeletePolicyStatement(ctx context.Context, params *DeletePolicyStatementInput, optFns ...func(*Options)) (*DeletePolicyStatementOutput, error) 
 DeleteSchemaMapping(ctx context.Context, params *DeleteSchemaMappingInput, optFns ...func(*Options)) (*DeleteSchemaMappingOutput, error) 
 GetIdMappingJob(ctx context.Context, params *GetIdMappingJobInput, optFns ...func(*Options)) (*GetIdMappingJobOutput, error) 
 GetIdMappingWorkflow(ctx context.Context, params *GetIdMappingWorkflowInput, optFns ...func(*Options)) (*GetIdMappingWorkflowOutput, error) 
 GetIdNamespace(ctx context.Context, params *GetIdNamespaceInput, optFns ...func(*Options)) (*GetIdNamespaceOutput, error) 
 GetMatchId(ctx context.Context, params *GetMatchIdInput, optFns ...func(*Options)) (*GetMatchIdOutput, error) 
 GetMatchingJob(ctx context.Context, params *GetMatchingJobInput, optFns ...func(*Options)) (*GetMatchingJobOutput, error) 
 GetMatchingWorkflow(ctx context.Context, params *GetMatchingWorkflowInput, optFns ...func(*Options)) (*GetMatchingWorkflowOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetProviderService(ctx context.Context, params *GetProviderServiceInput, optFns ...func(*Options)) (*GetProviderServiceOutput, error) 
 GetSchemaMapping(ctx context.Context, params *GetSchemaMappingInput, optFns ...func(*Options)) (*GetSchemaMappingOutput, error) 
 ListIdMappingJobs(ctx context.Context, params *ListIdMappingJobsInput, optFns ...func(*Options)) (*ListIdMappingJobsOutput, error) 
 ListIdMappingWorkflows(ctx context.Context, params *ListIdMappingWorkflowsInput, optFns ...func(*Options)) (*ListIdMappingWorkflowsOutput, error) 
 ListIdNamespaces(ctx context.Context, params *ListIdNamespacesInput, optFns ...func(*Options)) (*ListIdNamespacesOutput, error) 
 ListMatchingJobs(ctx context.Context, params *ListMatchingJobsInput, optFns ...func(*Options)) (*ListMatchingJobsOutput, error) 
 ListMatchingWorkflows(ctx context.Context, params *ListMatchingWorkflowsInput, optFns ...func(*Options)) (*ListMatchingWorkflowsOutput, error) 
 ListProviderServices(ctx context.Context, params *ListProviderServicesInput, optFns ...func(*Options)) (*ListProviderServicesOutput, error) 
 ListSchemaMappings(ctx context.Context, params *ListSchemaMappingsInput, optFns ...func(*Options)) (*ListSchemaMappingsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) 
 StartIdMappingJob(ctx context.Context, params *StartIdMappingJobInput, optFns ...func(*Options)) (*StartIdMappingJobOutput, error) 
 StartMatchingJob(ctx context.Context, params *StartMatchingJobInput, optFns ...func(*Options)) (*StartMatchingJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateIdMappingWorkflow(ctx context.Context, params *UpdateIdMappingWorkflowInput, optFns ...func(*Options)) (*UpdateIdMappingWorkflowOutput, error) 
 UpdateIdNamespace(ctx context.Context, params *UpdateIdNamespaceInput, optFns ...func(*Options)) (*UpdateIdNamespaceOutput, error) 
 UpdateMatchingWorkflow(ctx context.Context, params *UpdateMatchingWorkflowInput, optFns ...func(*Options)) (*UpdateMatchingWorkflowOutput, error) 
 UpdateSchemaMapping(ctx context.Context, params *UpdateSchemaMappingInput, optFns ...func(*Options)) (*UpdateSchemaMappingOutput, error) 
}
