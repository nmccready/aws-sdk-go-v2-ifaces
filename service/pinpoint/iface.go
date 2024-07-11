
package pinpoint

import (
    "github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

// IClient defines the interface for pinpoint
type IClient interface {
 Options() Options 
 CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) 
 CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) 
 CreateEmailTemplate(ctx context.Context, params *CreateEmailTemplateInput, optFns ...func(*Options)) (*CreateEmailTemplateOutput, error) 
 CreateExportJob(ctx context.Context, params *CreateExportJobInput, optFns ...func(*Options)) (*CreateExportJobOutput, error) 
 CreateImportJob(ctx context.Context, params *CreateImportJobInput, optFns ...func(*Options)) (*CreateImportJobOutput, error) 
 CreateInAppTemplate(ctx context.Context, params *CreateInAppTemplateInput, optFns ...func(*Options)) (*CreateInAppTemplateOutput, error) 
 CreateJourney(ctx context.Context, params *CreateJourneyInput, optFns ...func(*Options)) (*CreateJourneyOutput, error) 
 CreatePushTemplate(ctx context.Context, params *CreatePushTemplateInput, optFns ...func(*Options)) (*CreatePushTemplateOutput, error) 
 CreateRecommenderConfiguration(ctx context.Context, params *CreateRecommenderConfigurationInput, optFns ...func(*Options)) (*CreateRecommenderConfigurationOutput, error) 
 CreateSegment(ctx context.Context, params *CreateSegmentInput, optFns ...func(*Options)) (*CreateSegmentOutput, error) 
 CreateSmsTemplate(ctx context.Context, params *CreateSmsTemplateInput, optFns ...func(*Options)) (*CreateSmsTemplateOutput, error) 
 CreateVoiceTemplate(ctx context.Context, params *CreateVoiceTemplateInput, optFns ...func(*Options)) (*CreateVoiceTemplateOutput, error) 
 DeleteAdmChannel(ctx context.Context, params *DeleteAdmChannelInput, optFns ...func(*Options)) (*DeleteAdmChannelOutput, error) 
 DeleteApnsChannel(ctx context.Context, params *DeleteApnsChannelInput, optFns ...func(*Options)) (*DeleteApnsChannelOutput, error) 
 DeleteApnsSandboxChannel(ctx context.Context, params *DeleteApnsSandboxChannelInput, optFns ...func(*Options)) (*DeleteApnsSandboxChannelOutput, error) 
 DeleteApnsVoipChannel(ctx context.Context, params *DeleteApnsVoipChannelInput, optFns ...func(*Options)) (*DeleteApnsVoipChannelOutput, error) 
 DeleteApnsVoipSandboxChannel(ctx context.Context, params *DeleteApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*DeleteApnsVoipSandboxChannelOutput, error) 
 DeleteApp(ctx context.Context, params *DeleteAppInput, optFns ...func(*Options)) (*DeleteAppOutput, error) 
 DeleteBaiduChannel(ctx context.Context, params *DeleteBaiduChannelInput, optFns ...func(*Options)) (*DeleteBaiduChannelOutput, error) 
 DeleteCampaign(ctx context.Context, params *DeleteCampaignInput, optFns ...func(*Options)) (*DeleteCampaignOutput, error) 
 DeleteEmailChannel(ctx context.Context, params *DeleteEmailChannelInput, optFns ...func(*Options)) (*DeleteEmailChannelOutput, error) 
 DeleteEmailTemplate(ctx context.Context, params *DeleteEmailTemplateInput, optFns ...func(*Options)) (*DeleteEmailTemplateOutput, error) 
 DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) 
 DeleteEventStream(ctx context.Context, params *DeleteEventStreamInput, optFns ...func(*Options)) (*DeleteEventStreamOutput, error) 
 DeleteGcmChannel(ctx context.Context, params *DeleteGcmChannelInput, optFns ...func(*Options)) (*DeleteGcmChannelOutput, error) 
 DeleteInAppTemplate(ctx context.Context, params *DeleteInAppTemplateInput, optFns ...func(*Options)) (*DeleteInAppTemplateOutput, error) 
 DeleteJourney(ctx context.Context, params *DeleteJourneyInput, optFns ...func(*Options)) (*DeleteJourneyOutput, error) 
 DeletePushTemplate(ctx context.Context, params *DeletePushTemplateInput, optFns ...func(*Options)) (*DeletePushTemplateOutput, error) 
 DeleteRecommenderConfiguration(ctx context.Context, params *DeleteRecommenderConfigurationInput, optFns ...func(*Options)) (*DeleteRecommenderConfigurationOutput, error) 
 DeleteSegment(ctx context.Context, params *DeleteSegmentInput, optFns ...func(*Options)) (*DeleteSegmentOutput, error) 
 DeleteSmsChannel(ctx context.Context, params *DeleteSmsChannelInput, optFns ...func(*Options)) (*DeleteSmsChannelOutput, error) 
 DeleteSmsTemplate(ctx context.Context, params *DeleteSmsTemplateInput, optFns ...func(*Options)) (*DeleteSmsTemplateOutput, error) 
 DeleteUserEndpoints(ctx context.Context, params *DeleteUserEndpointsInput, optFns ...func(*Options)) (*DeleteUserEndpointsOutput, error) 
 DeleteVoiceChannel(ctx context.Context, params *DeleteVoiceChannelInput, optFns ...func(*Options)) (*DeleteVoiceChannelOutput, error) 
 DeleteVoiceTemplate(ctx context.Context, params *DeleteVoiceTemplateInput, optFns ...func(*Options)) (*DeleteVoiceTemplateOutput, error) 
 GetAdmChannel(ctx context.Context, params *GetAdmChannelInput, optFns ...func(*Options)) (*GetAdmChannelOutput, error) 
 GetApnsChannel(ctx context.Context, params *GetApnsChannelInput, optFns ...func(*Options)) (*GetApnsChannelOutput, error) 
 GetApnsSandboxChannel(ctx context.Context, params *GetApnsSandboxChannelInput, optFns ...func(*Options)) (*GetApnsSandboxChannelOutput, error) 
 GetApnsVoipChannel(ctx context.Context, params *GetApnsVoipChannelInput, optFns ...func(*Options)) (*GetApnsVoipChannelOutput, error) 
 GetApnsVoipSandboxChannel(ctx context.Context, params *GetApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*GetApnsVoipSandboxChannelOutput, error) 
 GetApp(ctx context.Context, params *GetAppInput, optFns ...func(*Options)) (*GetAppOutput, error) 
 GetApplicationDateRangeKpi(ctx context.Context, params *GetApplicationDateRangeKpiInput, optFns ...func(*Options)) (*GetApplicationDateRangeKpiOutput, error) 
 GetApplicationSettings(ctx context.Context, params *GetApplicationSettingsInput, optFns ...func(*Options)) (*GetApplicationSettingsOutput, error) 
 GetApps(ctx context.Context, params *GetAppsInput, optFns ...func(*Options)) (*GetAppsOutput, error) 
 GetBaiduChannel(ctx context.Context, params *GetBaiduChannelInput, optFns ...func(*Options)) (*GetBaiduChannelOutput, error) 
 GetCampaign(ctx context.Context, params *GetCampaignInput, optFns ...func(*Options)) (*GetCampaignOutput, error) 
 GetCampaignActivities(ctx context.Context, params *GetCampaignActivitiesInput, optFns ...func(*Options)) (*GetCampaignActivitiesOutput, error) 
 GetCampaignDateRangeKpi(ctx context.Context, params *GetCampaignDateRangeKpiInput, optFns ...func(*Options)) (*GetCampaignDateRangeKpiOutput, error) 
 GetCampaignVersion(ctx context.Context, params *GetCampaignVersionInput, optFns ...func(*Options)) (*GetCampaignVersionOutput, error) 
 GetCampaignVersions(ctx context.Context, params *GetCampaignVersionsInput, optFns ...func(*Options)) (*GetCampaignVersionsOutput, error) 
 GetCampaigns(ctx context.Context, params *GetCampaignsInput, optFns ...func(*Options)) (*GetCampaignsOutput, error) 
 GetChannels(ctx context.Context, params *GetChannelsInput, optFns ...func(*Options)) (*GetChannelsOutput, error) 
 GetEmailChannel(ctx context.Context, params *GetEmailChannelInput, optFns ...func(*Options)) (*GetEmailChannelOutput, error) 
 GetEmailTemplate(ctx context.Context, params *GetEmailTemplateInput, optFns ...func(*Options)) (*GetEmailTemplateOutput, error) 
 GetEndpoint(ctx context.Context, params *GetEndpointInput, optFns ...func(*Options)) (*GetEndpointOutput, error) 
 GetEventStream(ctx context.Context, params *GetEventStreamInput, optFns ...func(*Options)) (*GetEventStreamOutput, error) 
 GetExportJob(ctx context.Context, params *GetExportJobInput, optFns ...func(*Options)) (*GetExportJobOutput, error) 
 GetExportJobs(ctx context.Context, params *GetExportJobsInput, optFns ...func(*Options)) (*GetExportJobsOutput, error) 
 GetGcmChannel(ctx context.Context, params *GetGcmChannelInput, optFns ...func(*Options)) (*GetGcmChannelOutput, error) 
 GetImportJob(ctx context.Context, params *GetImportJobInput, optFns ...func(*Options)) (*GetImportJobOutput, error) 
 GetImportJobs(ctx context.Context, params *GetImportJobsInput, optFns ...func(*Options)) (*GetImportJobsOutput, error) 
 GetInAppMessages(ctx context.Context, params *GetInAppMessagesInput, optFns ...func(*Options)) (*GetInAppMessagesOutput, error) 
 GetInAppTemplate(ctx context.Context, params *GetInAppTemplateInput, optFns ...func(*Options)) (*GetInAppTemplateOutput, error) 
 GetJourney(ctx context.Context, params *GetJourneyInput, optFns ...func(*Options)) (*GetJourneyOutput, error) 
 GetJourneyDateRangeKpi(ctx context.Context, params *GetJourneyDateRangeKpiInput, optFns ...func(*Options)) (*GetJourneyDateRangeKpiOutput, error) 
 GetJourneyExecutionActivityMetrics(ctx context.Context, params *GetJourneyExecutionActivityMetricsInput, optFns ...func(*Options)) (*GetJourneyExecutionActivityMetricsOutput, error) 
 GetJourneyExecutionMetrics(ctx context.Context, params *GetJourneyExecutionMetricsInput, optFns ...func(*Options)) (*GetJourneyExecutionMetricsOutput, error) 
 GetJourneyRunExecutionActivityMetrics(ctx context.Context, params *GetJourneyRunExecutionActivityMetricsInput, optFns ...func(*Options)) (*GetJourneyRunExecutionActivityMetricsOutput, error) 
 GetJourneyRunExecutionMetrics(ctx context.Context, params *GetJourneyRunExecutionMetricsInput, optFns ...func(*Options)) (*GetJourneyRunExecutionMetricsOutput, error) 
 GetJourneyRuns(ctx context.Context, params *GetJourneyRunsInput, optFns ...func(*Options)) (*GetJourneyRunsOutput, error) 
 GetPushTemplate(ctx context.Context, params *GetPushTemplateInput, optFns ...func(*Options)) (*GetPushTemplateOutput, error) 
 GetRecommenderConfiguration(ctx context.Context, params *GetRecommenderConfigurationInput, optFns ...func(*Options)) (*GetRecommenderConfigurationOutput, error) 
 GetRecommenderConfigurations(ctx context.Context, params *GetRecommenderConfigurationsInput, optFns ...func(*Options)) (*GetRecommenderConfigurationsOutput, error) 
 GetSegment(ctx context.Context, params *GetSegmentInput, optFns ...func(*Options)) (*GetSegmentOutput, error) 
 GetSegmentExportJobs(ctx context.Context, params *GetSegmentExportJobsInput, optFns ...func(*Options)) (*GetSegmentExportJobsOutput, error) 
 GetSegmentImportJobs(ctx context.Context, params *GetSegmentImportJobsInput, optFns ...func(*Options)) (*GetSegmentImportJobsOutput, error) 
 GetSegmentVersion(ctx context.Context, params *GetSegmentVersionInput, optFns ...func(*Options)) (*GetSegmentVersionOutput, error) 
 GetSegmentVersions(ctx context.Context, params *GetSegmentVersionsInput, optFns ...func(*Options)) (*GetSegmentVersionsOutput, error) 
 GetSegments(ctx context.Context, params *GetSegmentsInput, optFns ...func(*Options)) (*GetSegmentsOutput, error) 
 GetSmsChannel(ctx context.Context, params *GetSmsChannelInput, optFns ...func(*Options)) (*GetSmsChannelOutput, error) 
 GetSmsTemplate(ctx context.Context, params *GetSmsTemplateInput, optFns ...func(*Options)) (*GetSmsTemplateOutput, error) 
 GetUserEndpoints(ctx context.Context, params *GetUserEndpointsInput, optFns ...func(*Options)) (*GetUserEndpointsOutput, error) 
 GetVoiceChannel(ctx context.Context, params *GetVoiceChannelInput, optFns ...func(*Options)) (*GetVoiceChannelOutput, error) 
 GetVoiceTemplate(ctx context.Context, params *GetVoiceTemplateInput, optFns ...func(*Options)) (*GetVoiceTemplateOutput, error) 
 ListJourneys(ctx context.Context, params *ListJourneysInput, optFns ...func(*Options)) (*ListJourneysOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplateVersions(ctx context.Context, params *ListTemplateVersionsInput, optFns ...func(*Options)) (*ListTemplateVersionsOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 PhoneNumberValidate(ctx context.Context, params *PhoneNumberValidateInput, optFns ...func(*Options)) (*PhoneNumberValidateOutput, error) 
 PutEventStream(ctx context.Context, params *PutEventStreamInput, optFns ...func(*Options)) (*PutEventStreamOutput, error) 
 PutEvents(ctx context.Context, params *PutEventsInput, optFns ...func(*Options)) (*PutEventsOutput, error) 
 RemoveAttributes(ctx context.Context, params *RemoveAttributesInput, optFns ...func(*Options)) (*RemoveAttributesOutput, error) 
 SendMessages(ctx context.Context, params *SendMessagesInput, optFns ...func(*Options)) (*SendMessagesOutput, error) 
 SendOTPMessage(ctx context.Context, params *SendOTPMessageInput, optFns ...func(*Options)) (*SendOTPMessageOutput, error) 
 SendUsersMessages(ctx context.Context, params *SendUsersMessagesInput, optFns ...func(*Options)) (*SendUsersMessagesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAdmChannel(ctx context.Context, params *UpdateAdmChannelInput, optFns ...func(*Options)) (*UpdateAdmChannelOutput, error) 
 UpdateApnsChannel(ctx context.Context, params *UpdateApnsChannelInput, optFns ...func(*Options)) (*UpdateApnsChannelOutput, error) 
 UpdateApnsSandboxChannel(ctx context.Context, params *UpdateApnsSandboxChannelInput, optFns ...func(*Options)) (*UpdateApnsSandboxChannelOutput, error) 
 UpdateApnsVoipChannel(ctx context.Context, params *UpdateApnsVoipChannelInput, optFns ...func(*Options)) (*UpdateApnsVoipChannelOutput, error) 
 UpdateApnsVoipSandboxChannel(ctx context.Context, params *UpdateApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*UpdateApnsVoipSandboxChannelOutput, error) 
 UpdateApplicationSettings(ctx context.Context, params *UpdateApplicationSettingsInput, optFns ...func(*Options)) (*UpdateApplicationSettingsOutput, error) 
 UpdateBaiduChannel(ctx context.Context, params *UpdateBaiduChannelInput, optFns ...func(*Options)) (*UpdateBaiduChannelOutput, error) 
 UpdateCampaign(ctx context.Context, params *UpdateCampaignInput, optFns ...func(*Options)) (*UpdateCampaignOutput, error) 
 UpdateEmailChannel(ctx context.Context, params *UpdateEmailChannelInput, optFns ...func(*Options)) (*UpdateEmailChannelOutput, error) 
 UpdateEmailTemplate(ctx context.Context, params *UpdateEmailTemplateInput, optFns ...func(*Options)) (*UpdateEmailTemplateOutput, error) 
 UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) 
 UpdateEndpointsBatch(ctx context.Context, params *UpdateEndpointsBatchInput, optFns ...func(*Options)) (*UpdateEndpointsBatchOutput, error) 
 UpdateGcmChannel(ctx context.Context, params *UpdateGcmChannelInput, optFns ...func(*Options)) (*UpdateGcmChannelOutput, error) 
 UpdateInAppTemplate(ctx context.Context, params *UpdateInAppTemplateInput, optFns ...func(*Options)) (*UpdateInAppTemplateOutput, error) 
 UpdateJourney(ctx context.Context, params *UpdateJourneyInput, optFns ...func(*Options)) (*UpdateJourneyOutput, error) 
 UpdateJourneyState(ctx context.Context, params *UpdateJourneyStateInput, optFns ...func(*Options)) (*UpdateJourneyStateOutput, error) 
 UpdatePushTemplate(ctx context.Context, params *UpdatePushTemplateInput, optFns ...func(*Options)) (*UpdatePushTemplateOutput, error) 
 UpdateRecommenderConfiguration(ctx context.Context, params *UpdateRecommenderConfigurationInput, optFns ...func(*Options)) (*UpdateRecommenderConfigurationOutput, error) 
 UpdateSegment(ctx context.Context, params *UpdateSegmentInput, optFns ...func(*Options)) (*UpdateSegmentOutput, error) 
 UpdateSmsChannel(ctx context.Context, params *UpdateSmsChannelInput, optFns ...func(*Options)) (*UpdateSmsChannelOutput, error) 
 UpdateSmsTemplate(ctx context.Context, params *UpdateSmsTemplateInput, optFns ...func(*Options)) (*UpdateSmsTemplateOutput, error) 
 UpdateTemplateActiveVersion(ctx context.Context, params *UpdateTemplateActiveVersionInput, optFns ...func(*Options)) (*UpdateTemplateActiveVersionOutput, error) 
 UpdateVoiceChannel(ctx context.Context, params *UpdateVoiceChannelInput, optFns ...func(*Options)) (*UpdateVoiceChannelOutput, error) 
 UpdateVoiceTemplate(ctx context.Context, params *UpdateVoiceTemplateInput, optFns ...func(*Options)) (*UpdateVoiceTemplateOutput, error) 
 VerifyOTPMessage(ctx context.Context, params *VerifyOTPMessageInput, optFns ...func(*Options)) (*VerifyOTPMessageOutput, error) 
}
