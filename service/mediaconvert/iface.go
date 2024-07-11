
package mediaconvert

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

// IClient defines the interface for mediaconvert
type IClient interface {
 Options() Options 
 AssociateCertificate(ctx context.Context, params *AssociateCertificateInput, optFns ...func(*Options)) (*AssociateCertificateOutput, error) 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) 
 CreatePreset(ctx context.Context, params *CreatePresetInput, optFns ...func(*Options)) (*CreatePresetOutput, error) 
 CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) 
 DeleteJobTemplate(ctx context.Context, params *DeleteJobTemplateInput, optFns ...func(*Options)) (*DeleteJobTemplateOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DeletePreset(ctx context.Context, params *DeletePresetInput, optFns ...func(*Options)) (*DeletePresetOutput, error) 
 DeleteQueue(ctx context.Context, params *DeleteQueueInput, optFns ...func(*Options)) (*DeleteQueueOutput, error) 
 DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) 
 DisassociateCertificate(ctx context.Context, params *DisassociateCertificateInput, optFns ...func(*Options)) (*DisassociateCertificateOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetJobTemplate(ctx context.Context, params *GetJobTemplateInput, optFns ...func(*Options)) (*GetJobTemplateOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetPreset(ctx context.Context, params *GetPresetInput, optFns ...func(*Options)) (*GetPresetOutput, error) 
 GetQueue(ctx context.Context, params *GetQueueInput, optFns ...func(*Options)) (*GetQueueOutput, error) 
 ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) 
 ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) 
 SearchJobs(ctx context.Context, params *SearchJobsInput, optFns ...func(*Options)) (*SearchJobsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateJobTemplate(ctx context.Context, params *UpdateJobTemplateInput, optFns ...func(*Options)) (*UpdateJobTemplateOutput, error) 
 UpdatePreset(ctx context.Context, params *UpdatePresetInput, optFns ...func(*Options)) (*UpdatePresetOutput, error) 
 UpdateQueue(ctx context.Context, params *UpdateQueueInput, optFns ...func(*Options)) (*UpdateQueueOutput, error) 
}
