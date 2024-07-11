
package connectcampaigns

import (
    "github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
)

// IConnectcampaigns defines the interface for connectcampaigns
type IConnectcampaigns interface {
 Options() Options 
 CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) 
 DeleteCampaign(ctx context.Context, params *DeleteCampaignInput, optFns ...func(*Options)) (*DeleteCampaignOutput, error) 
 DeleteConnectInstanceConfig(ctx context.Context, params *DeleteConnectInstanceConfigInput, optFns ...func(*Options)) (*DeleteConnectInstanceConfigOutput, error) 
 DeleteInstanceOnboardingJob(ctx context.Context, params *DeleteInstanceOnboardingJobInput, optFns ...func(*Options)) (*DeleteInstanceOnboardingJobOutput, error) 
 DescribeCampaign(ctx context.Context, params *DescribeCampaignInput, optFns ...func(*Options)) (*DescribeCampaignOutput, error) 
 GetCampaignState(ctx context.Context, params *GetCampaignStateInput, optFns ...func(*Options)) (*GetCampaignStateOutput, error) 
 GetCampaignStateBatch(ctx context.Context, params *GetCampaignStateBatchInput, optFns ...func(*Options)) (*GetCampaignStateBatchOutput, error) 
 GetConnectInstanceConfig(ctx context.Context, params *GetConnectInstanceConfigInput, optFns ...func(*Options)) (*GetConnectInstanceConfigOutput, error) 
 GetInstanceOnboardingJobStatus(ctx context.Context, params *GetInstanceOnboardingJobStatusInput, optFns ...func(*Options)) (*GetInstanceOnboardingJobStatusOutput, error) 
 ListCampaigns(ctx context.Context, params *ListCampaignsInput, optFns ...func(*Options)) (*ListCampaignsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PauseCampaign(ctx context.Context, params *PauseCampaignInput, optFns ...func(*Options)) (*PauseCampaignOutput, error) 
 PutDialRequestBatch(ctx context.Context, params *PutDialRequestBatchInput, optFns ...func(*Options)) (*PutDialRequestBatchOutput, error) 
 ResumeCampaign(ctx context.Context, params *ResumeCampaignInput, optFns ...func(*Options)) (*ResumeCampaignOutput, error) 
 StartCampaign(ctx context.Context, params *StartCampaignInput, optFns ...func(*Options)) (*StartCampaignOutput, error) 
 StartInstanceOnboardingJob(ctx context.Context, params *StartInstanceOnboardingJobInput, optFns ...func(*Options)) (*StartInstanceOnboardingJobOutput, error) 
 StopCampaign(ctx context.Context, params *StopCampaignInput, optFns ...func(*Options)) (*StopCampaignOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCampaignDialerConfig(ctx context.Context, params *UpdateCampaignDialerConfigInput, optFns ...func(*Options)) (*UpdateCampaignDialerConfigOutput, error) 
 UpdateCampaignName(ctx context.Context, params *UpdateCampaignNameInput, optFns ...func(*Options)) (*UpdateCampaignNameOutput, error) 
 UpdateCampaignOutboundCallConfig(ctx context.Context, params *UpdateCampaignOutboundCallConfigInput, optFns ...func(*Options)) (*UpdateCampaignOutboundCallConfigOutput, error) 
}
