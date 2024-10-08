
package elastictranscoder_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

// IClient defines the interface for elastictranscoder
type IClient interface {
 Options() Options 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) 
 CreatePreset(ctx context.Context, params *CreatePresetInput, optFns ...func(*Options)) (*CreatePresetOutput, error) 
 DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) 
 DeletePreset(ctx context.Context, params *DeletePresetInput, optFns ...func(*Options)) (*DeletePresetOutput, error) 
 ListJobsByPipeline(ctx context.Context, params *ListJobsByPipelineInput, optFns ...func(*Options)) (*ListJobsByPipelineOutput, error) 
 ListJobsByStatus(ctx context.Context, params *ListJobsByStatusInput, optFns ...func(*Options)) (*ListJobsByStatusOutput, error) 
 ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) 
 ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) 
 ReadJob(ctx context.Context, params *ReadJobInput, optFns ...func(*Options)) (*ReadJobOutput, error) 
 ReadPipeline(ctx context.Context, params *ReadPipelineInput, optFns ...func(*Options)) (*ReadPipelineOutput, error) 
 ReadPreset(ctx context.Context, params *ReadPresetInput, optFns ...func(*Options)) (*ReadPresetOutput, error) 
 TestRole(ctx context.Context, params *TestRoleInput, optFns ...func(*Options)) (*TestRoleOutput, error) 
 UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) 
 UpdatePipelineNotifications(ctx context.Context, params *UpdatePipelineNotificationsInput, optFns ...func(*Options)) (*UpdatePipelineNotificationsOutput, error) 
 UpdatePipelineStatus(ctx context.Context, params *UpdatePipelineStatusInput, optFns ...func(*Options)) (*UpdatePipelineStatusOutput, error) 
}
