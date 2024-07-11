
package migrationhuborchestrator

import (
    "github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator"
)

// IMigrationhuborchestrator defines the interface for migrationhuborchestrator
type IMigrationhuborchestrator interface {
 Options() Options 
 CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) 
 CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) 
 CreateWorkflowStep(ctx context.Context, params *CreateWorkflowStepInput, optFns ...func(*Options)) (*CreateWorkflowStepOutput, error) 
 CreateWorkflowStepGroup(ctx context.Context, params *CreateWorkflowStepGroupInput, optFns ...func(*Options)) (*CreateWorkflowStepGroupOutput, error) 
 DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) 
 DeleteWorkflow(ctx context.Context, params *DeleteWorkflowInput, optFns ...func(*Options)) (*DeleteWorkflowOutput, error) 
 DeleteWorkflowStep(ctx context.Context, params *DeleteWorkflowStepInput, optFns ...func(*Options)) (*DeleteWorkflowStepOutput, error) 
 DeleteWorkflowStepGroup(ctx context.Context, params *DeleteWorkflowStepGroupInput, optFns ...func(*Options)) (*DeleteWorkflowStepGroupOutput, error) 
 GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) 
 GetTemplateStep(ctx context.Context, params *GetTemplateStepInput, optFns ...func(*Options)) (*GetTemplateStepOutput, error) 
 GetTemplateStepGroup(ctx context.Context, params *GetTemplateStepGroupInput, optFns ...func(*Options)) (*GetTemplateStepGroupOutput, error) 
 GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) 
 GetWorkflowStep(ctx context.Context, params *GetWorkflowStepInput, optFns ...func(*Options)) (*GetWorkflowStepOutput, error) 
 GetWorkflowStepGroup(ctx context.Context, params *GetWorkflowStepGroupInput, optFns ...func(*Options)) (*GetWorkflowStepGroupOutput, error) 
 ListPlugins(ctx context.Context, params *ListPluginsInput, optFns ...func(*Options)) (*ListPluginsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplateStepGroups(ctx context.Context, params *ListTemplateStepGroupsInput, optFns ...func(*Options)) (*ListTemplateStepGroupsOutput, error) 
 ListTemplateSteps(ctx context.Context, params *ListTemplateStepsInput, optFns ...func(*Options)) (*ListTemplateStepsOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 ListWorkflowStepGroups(ctx context.Context, params *ListWorkflowStepGroupsInput, optFns ...func(*Options)) (*ListWorkflowStepGroupsOutput, error) 
 ListWorkflowSteps(ctx context.Context, params *ListWorkflowStepsInput, optFns ...func(*Options)) (*ListWorkflowStepsOutput, error) 
 ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) 
 RetryWorkflowStep(ctx context.Context, params *RetryWorkflowStepInput, optFns ...func(*Options)) (*RetryWorkflowStepOutput, error) 
 StartWorkflow(ctx context.Context, params *StartWorkflowInput, optFns ...func(*Options)) (*StartWorkflowOutput, error) 
 StopWorkflow(ctx context.Context, params *StopWorkflowInput, optFns ...func(*Options)) (*StopWorkflowOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) 
 UpdateWorkflow(ctx context.Context, params *UpdateWorkflowInput, optFns ...func(*Options)) (*UpdateWorkflowOutput, error) 
 UpdateWorkflowStep(ctx context.Context, params *UpdateWorkflowStepInput, optFns ...func(*Options)) (*UpdateWorkflowStepOutput, error) 
 UpdateWorkflowStepGroup(ctx context.Context, params *UpdateWorkflowStepGroupInput, optFns ...func(*Options)) (*UpdateWorkflowStepGroupOutput, error) 
}
