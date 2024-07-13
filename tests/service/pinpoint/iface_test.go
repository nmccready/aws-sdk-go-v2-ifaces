package pinpoint_test

// tests for the pinpoint service interface for this ../../../service/pinpoint/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpoint/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pinpoint/pinpoint_iface"
	"github.com/stretchr/testify/assert"
)

func TestPinpointServiceCanBeMocked(t *testing.T) {
	var iface pinpoint_iface.IClient
	iface = &pinpoint.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pinpoint.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &pinpoint.CreateAppInput{}
        output := &pinpoint.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCampaign", func(t *testing.T) {
        input := &pinpoint.CreateCampaignInput{}
        output := &pinpoint.CreateCampaignOutput{}

        mockClient.On("CreateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEmailTemplate", func(t *testing.T) {
        input := &pinpoint.CreateEmailTemplateInput{}
        output := &pinpoint.CreateEmailTemplateOutput{}

        mockClient.On("CreateEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExportJob", func(t *testing.T) {
        input := &pinpoint.CreateExportJobInput{}
        output := &pinpoint.CreateExportJobOutput{}

        mockClient.On("CreateExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImportJob", func(t *testing.T) {
        input := &pinpoint.CreateImportJobInput{}
        output := &pinpoint.CreateImportJobOutput{}

        mockClient.On("CreateImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInAppTemplate", func(t *testing.T) {
        input := &pinpoint.CreateInAppTemplateInput{}
        output := &pinpoint.CreateInAppTemplateOutput{}

        mockClient.On("CreateInAppTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInAppTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJourney", func(t *testing.T) {
        input := &pinpoint.CreateJourneyInput{}
        output := &pinpoint.CreateJourneyOutput{}

        mockClient.On("CreateJourney", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJourney(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePushTemplate", func(t *testing.T) {
        input := &pinpoint.CreatePushTemplateInput{}
        output := &pinpoint.CreatePushTemplateOutput{}

        mockClient.On("CreatePushTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePushTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecommenderConfiguration", func(t *testing.T) {
        input := &pinpoint.CreateRecommenderConfigurationInput{}
        output := &pinpoint.CreateRecommenderConfigurationOutput{}

        mockClient.On("CreateRecommenderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecommenderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSegment", func(t *testing.T) {
        input := &pinpoint.CreateSegmentInput{}
        output := &pinpoint.CreateSegmentOutput{}

        mockClient.On("CreateSegment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSmsTemplate", func(t *testing.T) {
        input := &pinpoint.CreateSmsTemplateInput{}
        output := &pinpoint.CreateSmsTemplateOutput{}

        mockClient.On("CreateSmsTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSmsTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceTemplate", func(t *testing.T) {
        input := &pinpoint.CreateVoiceTemplateInput{}
        output := &pinpoint.CreateVoiceTemplateOutput{}

        mockClient.On("CreateVoiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAdmChannel", func(t *testing.T) {
        input := &pinpoint.DeleteAdmChannelInput{}
        output := &pinpoint.DeleteAdmChannelOutput{}

        mockClient.On("DeleteAdmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAdmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApnsChannel", func(t *testing.T) {
        input := &pinpoint.DeleteApnsChannelInput{}
        output := &pinpoint.DeleteApnsChannelOutput{}

        mockClient.On("DeleteApnsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApnsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApnsSandboxChannel", func(t *testing.T) {
        input := &pinpoint.DeleteApnsSandboxChannelInput{}
        output := &pinpoint.DeleteApnsSandboxChannelOutput{}

        mockClient.On("DeleteApnsSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApnsSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApnsVoipChannel", func(t *testing.T) {
        input := &pinpoint.DeleteApnsVoipChannelInput{}
        output := &pinpoint.DeleteApnsVoipChannelOutput{}

        mockClient.On("DeleteApnsVoipChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApnsVoipChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApnsVoipSandboxChannel", func(t *testing.T) {
        input := &pinpoint.DeleteApnsVoipSandboxChannelInput{}
        output := &pinpoint.DeleteApnsVoipSandboxChannelOutput{}

        mockClient.On("DeleteApnsVoipSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApnsVoipSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &pinpoint.DeleteAppInput{}
        output := &pinpoint.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBaiduChannel", func(t *testing.T) {
        input := &pinpoint.DeleteBaiduChannelInput{}
        output := &pinpoint.DeleteBaiduChannelOutput{}

        mockClient.On("DeleteBaiduChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBaiduChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaign", func(t *testing.T) {
        input := &pinpoint.DeleteCampaignInput{}
        output := &pinpoint.DeleteCampaignOutput{}

        mockClient.On("DeleteCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailChannel", func(t *testing.T) {
        input := &pinpoint.DeleteEmailChannelInput{}
        output := &pinpoint.DeleteEmailChannelOutput{}

        mockClient.On("DeleteEmailChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailTemplate", func(t *testing.T) {
        input := &pinpoint.DeleteEmailTemplateInput{}
        output := &pinpoint.DeleteEmailTemplateOutput{}

        mockClient.On("DeleteEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &pinpoint.DeleteEndpointInput{}
        output := &pinpoint.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventStream", func(t *testing.T) {
        input := &pinpoint.DeleteEventStreamInput{}
        output := &pinpoint.DeleteEventStreamOutput{}

        mockClient.On("DeleteEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGcmChannel", func(t *testing.T) {
        input := &pinpoint.DeleteGcmChannelInput{}
        output := &pinpoint.DeleteGcmChannelOutput{}

        mockClient.On("DeleteGcmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGcmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInAppTemplate", func(t *testing.T) {
        input := &pinpoint.DeleteInAppTemplateInput{}
        output := &pinpoint.DeleteInAppTemplateOutput{}

        mockClient.On("DeleteInAppTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInAppTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJourney", func(t *testing.T) {
        input := &pinpoint.DeleteJourneyInput{}
        output := &pinpoint.DeleteJourneyOutput{}

        mockClient.On("DeleteJourney", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJourney(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePushTemplate", func(t *testing.T) {
        input := &pinpoint.DeletePushTemplateInput{}
        output := &pinpoint.DeletePushTemplateOutput{}

        mockClient.On("DeletePushTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePushTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecommenderConfiguration", func(t *testing.T) {
        input := &pinpoint.DeleteRecommenderConfigurationInput{}
        output := &pinpoint.DeleteRecommenderConfigurationOutput{}

        mockClient.On("DeleteRecommenderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecommenderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSegment", func(t *testing.T) {
        input := &pinpoint.DeleteSegmentInput{}
        output := &pinpoint.DeleteSegmentOutput{}

        mockClient.On("DeleteSegment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSmsChannel", func(t *testing.T) {
        input := &pinpoint.DeleteSmsChannelInput{}
        output := &pinpoint.DeleteSmsChannelOutput{}

        mockClient.On("DeleteSmsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSmsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSmsTemplate", func(t *testing.T) {
        input := &pinpoint.DeleteSmsTemplateInput{}
        output := &pinpoint.DeleteSmsTemplateOutput{}

        mockClient.On("DeleteSmsTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSmsTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserEndpoints", func(t *testing.T) {
        input := &pinpoint.DeleteUserEndpointsInput{}
        output := &pinpoint.DeleteUserEndpointsOutput{}

        mockClient.On("DeleteUserEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceChannel", func(t *testing.T) {
        input := &pinpoint.DeleteVoiceChannelInput{}
        output := &pinpoint.DeleteVoiceChannelOutput{}

        mockClient.On("DeleteVoiceChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceTemplate", func(t *testing.T) {
        input := &pinpoint.DeleteVoiceTemplateInput{}
        output := &pinpoint.DeleteVoiceTemplateOutput{}

        mockClient.On("DeleteVoiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdmChannel", func(t *testing.T) {
        input := &pinpoint.GetAdmChannelInput{}
        output := &pinpoint.GetAdmChannelOutput{}

        mockClient.On("GetAdmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApnsChannel", func(t *testing.T) {
        input := &pinpoint.GetApnsChannelInput{}
        output := &pinpoint.GetApnsChannelOutput{}

        mockClient.On("GetApnsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetApnsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApnsSandboxChannel", func(t *testing.T) {
        input := &pinpoint.GetApnsSandboxChannelInput{}
        output := &pinpoint.GetApnsSandboxChannelOutput{}

        mockClient.On("GetApnsSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetApnsSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApnsVoipChannel", func(t *testing.T) {
        input := &pinpoint.GetApnsVoipChannelInput{}
        output := &pinpoint.GetApnsVoipChannelOutput{}

        mockClient.On("GetApnsVoipChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetApnsVoipChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApnsVoipSandboxChannel", func(t *testing.T) {
        input := &pinpoint.GetApnsVoipSandboxChannelInput{}
        output := &pinpoint.GetApnsVoipSandboxChannelOutput{}

        mockClient.On("GetApnsVoipSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetApnsVoipSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApp", func(t *testing.T) {
        input := &pinpoint.GetAppInput{}
        output := &pinpoint.GetAppOutput{}

        mockClient.On("GetApp", ctx, input).Return(output, nil)

        result, err := mockClient.GetApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationDateRangeKpi", func(t *testing.T) {
        input := &pinpoint.GetApplicationDateRangeKpiInput{}
        output := &pinpoint.GetApplicationDateRangeKpiOutput{}

        mockClient.On("GetApplicationDateRangeKpi", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationDateRangeKpi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationSettings", func(t *testing.T) {
        input := &pinpoint.GetApplicationSettingsInput{}
        output := &pinpoint.GetApplicationSettingsOutput{}

        mockClient.On("GetApplicationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApps", func(t *testing.T) {
        input := &pinpoint.GetAppsInput{}
        output := &pinpoint.GetAppsOutput{}

        mockClient.On("GetApps", ctx, input).Return(output, nil)

        result, err := mockClient.GetApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBaiduChannel", func(t *testing.T) {
        input := &pinpoint.GetBaiduChannelInput{}
        output := &pinpoint.GetBaiduChannelOutput{}

        mockClient.On("GetBaiduChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetBaiduChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaign", func(t *testing.T) {
        input := &pinpoint.GetCampaignInput{}
        output := &pinpoint.GetCampaignOutput{}

        mockClient.On("GetCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignActivities", func(t *testing.T) {
        input := &pinpoint.GetCampaignActivitiesInput{}
        output := &pinpoint.GetCampaignActivitiesOutput{}

        mockClient.On("GetCampaignActivities", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignDateRangeKpi", func(t *testing.T) {
        input := &pinpoint.GetCampaignDateRangeKpiInput{}
        output := &pinpoint.GetCampaignDateRangeKpiOutput{}

        mockClient.On("GetCampaignDateRangeKpi", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignDateRangeKpi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignVersion", func(t *testing.T) {
        input := &pinpoint.GetCampaignVersionInput{}
        output := &pinpoint.GetCampaignVersionOutput{}

        mockClient.On("GetCampaignVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignVersions", func(t *testing.T) {
        input := &pinpoint.GetCampaignVersionsInput{}
        output := &pinpoint.GetCampaignVersionsOutput{}

        mockClient.On("GetCampaignVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaigns", func(t *testing.T) {
        input := &pinpoint.GetCampaignsInput{}
        output := &pinpoint.GetCampaignsOutput{}

        mockClient.On("GetCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannels", func(t *testing.T) {
        input := &pinpoint.GetChannelsInput{}
        output := &pinpoint.GetChannelsOutput{}

        mockClient.On("GetChannels", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailChannel", func(t *testing.T) {
        input := &pinpoint.GetEmailChannelInput{}
        output := &pinpoint.GetEmailChannelOutput{}

        mockClient.On("GetEmailChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEmailTemplate", func(t *testing.T) {
        input := &pinpoint.GetEmailTemplateInput{}
        output := &pinpoint.GetEmailTemplateOutput{}

        mockClient.On("GetEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEndpoint", func(t *testing.T) {
        input := &pinpoint.GetEndpointInput{}
        output := &pinpoint.GetEndpointOutput{}

        mockClient.On("GetEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventStream", func(t *testing.T) {
        input := &pinpoint.GetEventStreamInput{}
        output := &pinpoint.GetEventStreamOutput{}

        mockClient.On("GetEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExportJob", func(t *testing.T) {
        input := &pinpoint.GetExportJobInput{}
        output := &pinpoint.GetExportJobOutput{}

        mockClient.On("GetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExportJobs", func(t *testing.T) {
        input := &pinpoint.GetExportJobsInput{}
        output := &pinpoint.GetExportJobsOutput{}

        mockClient.On("GetExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGcmChannel", func(t *testing.T) {
        input := &pinpoint.GetGcmChannelInput{}
        output := &pinpoint.GetGcmChannelOutput{}

        mockClient.On("GetGcmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetGcmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportJob", func(t *testing.T) {
        input := &pinpoint.GetImportJobInput{}
        output := &pinpoint.GetImportJobOutput{}

        mockClient.On("GetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImportJobs", func(t *testing.T) {
        input := &pinpoint.GetImportJobsInput{}
        output := &pinpoint.GetImportJobsOutput{}

        mockClient.On("GetImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInAppMessages", func(t *testing.T) {
        input := &pinpoint.GetInAppMessagesInput{}
        output := &pinpoint.GetInAppMessagesOutput{}

        mockClient.On("GetInAppMessages", ctx, input).Return(output, nil)

        result, err := mockClient.GetInAppMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInAppTemplate", func(t *testing.T) {
        input := &pinpoint.GetInAppTemplateInput{}
        output := &pinpoint.GetInAppTemplateOutput{}

        mockClient.On("GetInAppTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetInAppTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourney", func(t *testing.T) {
        input := &pinpoint.GetJourneyInput{}
        output := &pinpoint.GetJourneyOutput{}

        mockClient.On("GetJourney", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourney(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyDateRangeKpi", func(t *testing.T) {
        input := &pinpoint.GetJourneyDateRangeKpiInput{}
        output := &pinpoint.GetJourneyDateRangeKpiOutput{}

        mockClient.On("GetJourneyDateRangeKpi", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyDateRangeKpi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyExecutionActivityMetrics", func(t *testing.T) {
        input := &pinpoint.GetJourneyExecutionActivityMetricsInput{}
        output := &pinpoint.GetJourneyExecutionActivityMetricsOutput{}

        mockClient.On("GetJourneyExecutionActivityMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyExecutionActivityMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyExecutionMetrics", func(t *testing.T) {
        input := &pinpoint.GetJourneyExecutionMetricsInput{}
        output := &pinpoint.GetJourneyExecutionMetricsOutput{}

        mockClient.On("GetJourneyExecutionMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyExecutionMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyRunExecutionActivityMetrics", func(t *testing.T) {
        input := &pinpoint.GetJourneyRunExecutionActivityMetricsInput{}
        output := &pinpoint.GetJourneyRunExecutionActivityMetricsOutput{}

        mockClient.On("GetJourneyRunExecutionActivityMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyRunExecutionActivityMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyRunExecutionMetrics", func(t *testing.T) {
        input := &pinpoint.GetJourneyRunExecutionMetricsInput{}
        output := &pinpoint.GetJourneyRunExecutionMetricsOutput{}

        mockClient.On("GetJourneyRunExecutionMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyRunExecutionMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJourneyRuns", func(t *testing.T) {
        input := &pinpoint.GetJourneyRunsInput{}
        output := &pinpoint.GetJourneyRunsOutput{}

        mockClient.On("GetJourneyRuns", ctx, input).Return(output, nil)

        result, err := mockClient.GetJourneyRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPushTemplate", func(t *testing.T) {
        input := &pinpoint.GetPushTemplateInput{}
        output := &pinpoint.GetPushTemplateOutput{}

        mockClient.On("GetPushTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetPushTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommenderConfiguration", func(t *testing.T) {
        input := &pinpoint.GetRecommenderConfigurationInput{}
        output := &pinpoint.GetRecommenderConfigurationOutput{}

        mockClient.On("GetRecommenderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommenderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommenderConfigurations", func(t *testing.T) {
        input := &pinpoint.GetRecommenderConfigurationsInput{}
        output := &pinpoint.GetRecommenderConfigurationsOutput{}

        mockClient.On("GetRecommenderConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommenderConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegment", func(t *testing.T) {
        input := &pinpoint.GetSegmentInput{}
        output := &pinpoint.GetSegmentOutput{}

        mockClient.On("GetSegment", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegmentExportJobs", func(t *testing.T) {
        input := &pinpoint.GetSegmentExportJobsInput{}
        output := &pinpoint.GetSegmentExportJobsOutput{}

        mockClient.On("GetSegmentExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegmentExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegmentImportJobs", func(t *testing.T) {
        input := &pinpoint.GetSegmentImportJobsInput{}
        output := &pinpoint.GetSegmentImportJobsOutput{}

        mockClient.On("GetSegmentImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegmentImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegmentVersion", func(t *testing.T) {
        input := &pinpoint.GetSegmentVersionInput{}
        output := &pinpoint.GetSegmentVersionOutput{}

        mockClient.On("GetSegmentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegmentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegmentVersions", func(t *testing.T) {
        input := &pinpoint.GetSegmentVersionsInput{}
        output := &pinpoint.GetSegmentVersionsOutput{}

        mockClient.On("GetSegmentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegmentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegments", func(t *testing.T) {
        input := &pinpoint.GetSegmentsInput{}
        output := &pinpoint.GetSegmentsOutput{}

        mockClient.On("GetSegments", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSmsChannel", func(t *testing.T) {
        input := &pinpoint.GetSmsChannelInput{}
        output := &pinpoint.GetSmsChannelOutput{}

        mockClient.On("GetSmsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetSmsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSmsTemplate", func(t *testing.T) {
        input := &pinpoint.GetSmsTemplateInput{}
        output := &pinpoint.GetSmsTemplateOutput{}

        mockClient.On("GetSmsTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetSmsTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserEndpoints", func(t *testing.T) {
        input := &pinpoint.GetUserEndpointsInput{}
        output := &pinpoint.GetUserEndpointsOutput{}

        mockClient.On("GetUserEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceChannel", func(t *testing.T) {
        input := &pinpoint.GetVoiceChannelInput{}
        output := &pinpoint.GetVoiceChannelOutput{}

        mockClient.On("GetVoiceChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceTemplate", func(t *testing.T) {
        input := &pinpoint.GetVoiceTemplateInput{}
        output := &pinpoint.GetVoiceTemplateOutput{}

        mockClient.On("GetVoiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJourneys", func(t *testing.T) {
        input := &pinpoint.ListJourneysInput{}
        output := &pinpoint.ListJourneysOutput{}

        mockClient.On("ListJourneys", ctx, input).Return(output, nil)

        result, err := mockClient.ListJourneys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pinpoint.ListTagsForResourceInput{}
        output := &pinpoint.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateVersions", func(t *testing.T) {
        input := &pinpoint.ListTemplateVersionsInput{}
        output := &pinpoint.ListTemplateVersionsOutput{}

        mockClient.On("ListTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &pinpoint.ListTemplatesInput{}
        output := &pinpoint.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPhoneNumberValidate", func(t *testing.T) {
        input := &pinpoint.PhoneNumberValidateInput{}
        output := &pinpoint.PhoneNumberValidateOutput{}

        mockClient.On("PhoneNumberValidate", ctx, input).Return(output, nil)

        result, err := mockClient.PhoneNumberValidate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEventStream", func(t *testing.T) {
        input := &pinpoint.PutEventStreamInput{}
        output := &pinpoint.PutEventStreamOutput{}

        mockClient.On("PutEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.PutEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEvents", func(t *testing.T) {
        input := &pinpoint.PutEventsInput{}
        output := &pinpoint.PutEventsOutput{}

        mockClient.On("PutEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAttributes", func(t *testing.T) {
        input := &pinpoint.RemoveAttributesInput{}
        output := &pinpoint.RemoveAttributesOutput{}

        mockClient.On("RemoveAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendMessages", func(t *testing.T) {
        input := &pinpoint.SendMessagesInput{}
        output := &pinpoint.SendMessagesOutput{}

        mockClient.On("SendMessages", ctx, input).Return(output, nil)

        result, err := mockClient.SendMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendOTPMessage", func(t *testing.T) {
        input := &pinpoint.SendOTPMessageInput{}
        output := &pinpoint.SendOTPMessageOutput{}

        mockClient.On("SendOTPMessage", ctx, input).Return(output, nil)

        result, err := mockClient.SendOTPMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendUsersMessages", func(t *testing.T) {
        input := &pinpoint.SendUsersMessagesInput{}
        output := &pinpoint.SendUsersMessagesOutput{}

        mockClient.On("SendUsersMessages", ctx, input).Return(output, nil)

        result, err := mockClient.SendUsersMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pinpoint.TagResourceInput{}
        output := &pinpoint.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pinpoint.UntagResourceInput{}
        output := &pinpoint.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAdmChannel", func(t *testing.T) {
        input := &pinpoint.UpdateAdmChannelInput{}
        output := &pinpoint.UpdateAdmChannelOutput{}

        mockClient.On("UpdateAdmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAdmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApnsChannel", func(t *testing.T) {
        input := &pinpoint.UpdateApnsChannelInput{}
        output := &pinpoint.UpdateApnsChannelOutput{}

        mockClient.On("UpdateApnsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApnsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApnsSandboxChannel", func(t *testing.T) {
        input := &pinpoint.UpdateApnsSandboxChannelInput{}
        output := &pinpoint.UpdateApnsSandboxChannelOutput{}

        mockClient.On("UpdateApnsSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApnsSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApnsVoipChannel", func(t *testing.T) {
        input := &pinpoint.UpdateApnsVoipChannelInput{}
        output := &pinpoint.UpdateApnsVoipChannelOutput{}

        mockClient.On("UpdateApnsVoipChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApnsVoipChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApnsVoipSandboxChannel", func(t *testing.T) {
        input := &pinpoint.UpdateApnsVoipSandboxChannelInput{}
        output := &pinpoint.UpdateApnsVoipSandboxChannelOutput{}

        mockClient.On("UpdateApnsVoipSandboxChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApnsVoipSandboxChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationSettings", func(t *testing.T) {
        input := &pinpoint.UpdateApplicationSettingsInput{}
        output := &pinpoint.UpdateApplicationSettingsOutput{}

        mockClient.On("UpdateApplicationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBaiduChannel", func(t *testing.T) {
        input := &pinpoint.UpdateBaiduChannelInput{}
        output := &pinpoint.UpdateBaiduChannelOutput{}

        mockClient.On("UpdateBaiduChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBaiduChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaign", func(t *testing.T) {
        input := &pinpoint.UpdateCampaignInput{}
        output := &pinpoint.UpdateCampaignOutput{}

        mockClient.On("UpdateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEmailChannel", func(t *testing.T) {
        input := &pinpoint.UpdateEmailChannelInput{}
        output := &pinpoint.UpdateEmailChannelOutput{}

        mockClient.On("UpdateEmailChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEmailChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEmailTemplate", func(t *testing.T) {
        input := &pinpoint.UpdateEmailTemplateInput{}
        output := &pinpoint.UpdateEmailTemplateOutput{}

        mockClient.On("UpdateEmailTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEmailTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpoint", func(t *testing.T) {
        input := &pinpoint.UpdateEndpointInput{}
        output := &pinpoint.UpdateEndpointOutput{}

        mockClient.On("UpdateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpointsBatch", func(t *testing.T) {
        input := &pinpoint.UpdateEndpointsBatchInput{}
        output := &pinpoint.UpdateEndpointsBatchOutput{}

        mockClient.On("UpdateEndpointsBatch", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpointsBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGcmChannel", func(t *testing.T) {
        input := &pinpoint.UpdateGcmChannelInput{}
        output := &pinpoint.UpdateGcmChannelOutput{}

        mockClient.On("UpdateGcmChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGcmChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInAppTemplate", func(t *testing.T) {
        input := &pinpoint.UpdateInAppTemplateInput{}
        output := &pinpoint.UpdateInAppTemplateOutput{}

        mockClient.On("UpdateInAppTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInAppTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJourney", func(t *testing.T) {
        input := &pinpoint.UpdateJourneyInput{}
        output := &pinpoint.UpdateJourneyOutput{}

        mockClient.On("UpdateJourney", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJourney(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJourneyState", func(t *testing.T) {
        input := &pinpoint.UpdateJourneyStateInput{}
        output := &pinpoint.UpdateJourneyStateOutput{}

        mockClient.On("UpdateJourneyState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJourneyState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePushTemplate", func(t *testing.T) {
        input := &pinpoint.UpdatePushTemplateInput{}
        output := &pinpoint.UpdatePushTemplateOutput{}

        mockClient.On("UpdatePushTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePushTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecommenderConfiguration", func(t *testing.T) {
        input := &pinpoint.UpdateRecommenderConfigurationInput{}
        output := &pinpoint.UpdateRecommenderConfigurationOutput{}

        mockClient.On("UpdateRecommenderConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecommenderConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSegment", func(t *testing.T) {
        input := &pinpoint.UpdateSegmentInput{}
        output := &pinpoint.UpdateSegmentOutput{}

        mockClient.On("UpdateSegment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSmsChannel", func(t *testing.T) {
        input := &pinpoint.UpdateSmsChannelInput{}
        output := &pinpoint.UpdateSmsChannelOutput{}

        mockClient.On("UpdateSmsChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSmsChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSmsTemplate", func(t *testing.T) {
        input := &pinpoint.UpdateSmsTemplateInput{}
        output := &pinpoint.UpdateSmsTemplateOutput{}

        mockClient.On("UpdateSmsTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSmsTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplateActiveVersion", func(t *testing.T) {
        input := &pinpoint.UpdateTemplateActiveVersionInput{}
        output := &pinpoint.UpdateTemplateActiveVersionOutput{}

        mockClient.On("UpdateTemplateActiveVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplateActiveVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceChannel", func(t *testing.T) {
        input := &pinpoint.UpdateVoiceChannelInput{}
        output := &pinpoint.UpdateVoiceChannelOutput{}

        mockClient.On("UpdateVoiceChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceTemplate", func(t *testing.T) {
        input := &pinpoint.UpdateVoiceTemplateInput{}
        output := &pinpoint.UpdateVoiceTemplateOutput{}

        mockClient.On("UpdateVoiceTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyOTPMessage", func(t *testing.T) {
        input := &pinpoint.VerifyOTPMessageInput{}
        output := &pinpoint.VerifyOTPMessageOutput{}

        mockClient.On("VerifyOTPMessage", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyOTPMessage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
