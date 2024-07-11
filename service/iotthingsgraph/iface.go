
package iotthingsgraph

import (
    "github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
)

// IIotthingsgraph defines the interface for iotthingsgraph
type IIotthingsgraph interface {
 Options() Options 
 AssociateEntityToThing(ctx context.Context, params *AssociateEntityToThingInput, optFns ...func(*Options)) (*AssociateEntityToThingOutput, error) 
 CreateFlowTemplate(ctx context.Context, params *CreateFlowTemplateInput, optFns ...func(*Options)) (*CreateFlowTemplateOutput, error) 
 CreateSystemInstance(ctx context.Context, params *CreateSystemInstanceInput, optFns ...func(*Options)) (*CreateSystemInstanceOutput, error) 
 CreateSystemTemplate(ctx context.Context, params *CreateSystemTemplateInput, optFns ...func(*Options)) (*CreateSystemTemplateOutput, error) 
 DeleteFlowTemplate(ctx context.Context, params *DeleteFlowTemplateInput, optFns ...func(*Options)) (*DeleteFlowTemplateOutput, error) 
 DeleteNamespace(ctx context.Context, params *DeleteNamespaceInput, optFns ...func(*Options)) (*DeleteNamespaceOutput, error) 
 DeleteSystemInstance(ctx context.Context, params *DeleteSystemInstanceInput, optFns ...func(*Options)) (*DeleteSystemInstanceOutput, error) 
 DeleteSystemTemplate(ctx context.Context, params *DeleteSystemTemplateInput, optFns ...func(*Options)) (*DeleteSystemTemplateOutput, error) 
 DeploySystemInstance(ctx context.Context, params *DeploySystemInstanceInput, optFns ...func(*Options)) (*DeploySystemInstanceOutput, error) 
 DeprecateFlowTemplate(ctx context.Context, params *DeprecateFlowTemplateInput, optFns ...func(*Options)) (*DeprecateFlowTemplateOutput, error) 
 DeprecateSystemTemplate(ctx context.Context, params *DeprecateSystemTemplateInput, optFns ...func(*Options)) (*DeprecateSystemTemplateOutput, error) 
 DescribeNamespace(ctx context.Context, params *DescribeNamespaceInput, optFns ...func(*Options)) (*DescribeNamespaceOutput, error) 
 DissociateEntityFromThing(ctx context.Context, params *DissociateEntityFromThingInput, optFns ...func(*Options)) (*DissociateEntityFromThingOutput, error) 
 GetEntities(ctx context.Context, params *GetEntitiesInput, optFns ...func(*Options)) (*GetEntitiesOutput, error) 
 GetFlowTemplate(ctx context.Context, params *GetFlowTemplateInput, optFns ...func(*Options)) (*GetFlowTemplateOutput, error) 
 GetFlowTemplateRevisions(ctx context.Context, params *GetFlowTemplateRevisionsInput, optFns ...func(*Options)) (*GetFlowTemplateRevisionsOutput, error) 
 GetNamespaceDeletionStatus(ctx context.Context, params *GetNamespaceDeletionStatusInput, optFns ...func(*Options)) (*GetNamespaceDeletionStatusOutput, error) 
 GetSystemInstance(ctx context.Context, params *GetSystemInstanceInput, optFns ...func(*Options)) (*GetSystemInstanceOutput, error) 
 GetSystemTemplate(ctx context.Context, params *GetSystemTemplateInput, optFns ...func(*Options)) (*GetSystemTemplateOutput, error) 
 GetSystemTemplateRevisions(ctx context.Context, params *GetSystemTemplateRevisionsInput, optFns ...func(*Options)) (*GetSystemTemplateRevisionsOutput, error) 
 GetUploadStatus(ctx context.Context, params *GetUploadStatusInput, optFns ...func(*Options)) (*GetUploadStatusOutput, error) 
 ListFlowExecutionMessages(ctx context.Context, params *ListFlowExecutionMessagesInput, optFns ...func(*Options)) (*ListFlowExecutionMessagesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 SearchEntities(ctx context.Context, params *SearchEntitiesInput, optFns ...func(*Options)) (*SearchEntitiesOutput, error) 
 SearchFlowExecutions(ctx context.Context, params *SearchFlowExecutionsInput, optFns ...func(*Options)) (*SearchFlowExecutionsOutput, error) 
 SearchFlowTemplates(ctx context.Context, params *SearchFlowTemplatesInput, optFns ...func(*Options)) (*SearchFlowTemplatesOutput, error) 
 SearchSystemInstances(ctx context.Context, params *SearchSystemInstancesInput, optFns ...func(*Options)) (*SearchSystemInstancesOutput, error) 
 SearchSystemTemplates(ctx context.Context, params *SearchSystemTemplatesInput, optFns ...func(*Options)) (*SearchSystemTemplatesOutput, error) 
 SearchThings(ctx context.Context, params *SearchThingsInput, optFns ...func(*Options)) (*SearchThingsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UndeploySystemInstance(ctx context.Context, params *UndeploySystemInstanceInput, optFns ...func(*Options)) (*UndeploySystemInstanceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateFlowTemplate(ctx context.Context, params *UpdateFlowTemplateInput, optFns ...func(*Options)) (*UpdateFlowTemplateOutput, error) 
 UpdateSystemTemplate(ctx context.Context, params *UpdateSystemTemplateInput, optFns ...func(*Options)) (*UpdateSystemTemplateOutput, error) 
 UploadEntityDefinitions(ctx context.Context, params *UploadEntityDefinitionsInput, optFns ...func(*Options)) (*UploadEntityDefinitionsOutput, error) 
}
