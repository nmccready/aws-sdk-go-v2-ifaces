// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package connectcampaignsv2_test

// tests for the connectcampaignsv2 service interface for 
// this ../../../service/connectcampaignsv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcampaignsv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcampaignsv2/connectcampaignsv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectcampaignsv2ServiceCanBeMocked(t *testing.T) {
	var iface connectcampaignsv2_iface.IClient
	iface = &connectcampaignsv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connectcampaignsv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.CreateCampaignInput{}
        output := &connectcampaignsv2.CreateCampaignOutput{}

        mockClient.On("CreateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteCampaignInput{}
        output := &connectcampaignsv2.DeleteCampaignOutput{}

        mockClient.On("DeleteCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaignChannelSubtypeConfig", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteCampaignChannelSubtypeConfigInput{}
        output := &connectcampaignsv2.DeleteCampaignChannelSubtypeConfigOutput{}

        mockClient.On("DeleteCampaignChannelSubtypeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaignChannelSubtypeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaignCommunicationLimits", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteCampaignCommunicationLimitsInput{}
        output := &connectcampaignsv2.DeleteCampaignCommunicationLimitsOutput{}

        mockClient.On("DeleteCampaignCommunicationLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaignCommunicationLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaignCommunicationTime", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteCampaignCommunicationTimeInput{}
        output := &connectcampaignsv2.DeleteCampaignCommunicationTimeOutput{}

        mockClient.On("DeleteCampaignCommunicationTime", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaignCommunicationTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectInstanceConfig", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteConnectInstanceConfigInput{}
        output := &connectcampaignsv2.DeleteConnectInstanceConfigOutput{}

        mockClient.On("DeleteConnectInstanceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectInstanceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectInstanceIntegration", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteConnectInstanceIntegrationInput{}
        output := &connectcampaignsv2.DeleteConnectInstanceIntegrationOutput{}

        mockClient.On("DeleteConnectInstanceIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectInstanceIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceOnboardingJob", func(t *testing.T) {
        input := &connectcampaignsv2.DeleteInstanceOnboardingJobInput{}
        output := &connectcampaignsv2.DeleteInstanceOnboardingJobOutput{}

        mockClient.On("DeleteInstanceOnboardingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceOnboardingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.DescribeCampaignInput{}
        output := &connectcampaignsv2.DescribeCampaignOutput{}

        mockClient.On("DescribeCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignState", func(t *testing.T) {
        input := &connectcampaignsv2.GetCampaignStateInput{}
        output := &connectcampaignsv2.GetCampaignStateOutput{}

        mockClient.On("GetCampaignState", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignStateBatch", func(t *testing.T) {
        input := &connectcampaignsv2.GetCampaignStateBatchInput{}
        output := &connectcampaignsv2.GetCampaignStateBatchOutput{}

        mockClient.On("GetCampaignStateBatch", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignStateBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectInstanceConfig", func(t *testing.T) {
        input := &connectcampaignsv2.GetConnectInstanceConfigInput{}
        output := &connectcampaignsv2.GetConnectInstanceConfigOutput{}

        mockClient.On("GetConnectInstanceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectInstanceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceOnboardingJobStatus", func(t *testing.T) {
        input := &connectcampaignsv2.GetInstanceOnboardingJobStatusInput{}
        output := &connectcampaignsv2.GetInstanceOnboardingJobStatusOutput{}

        mockClient.On("GetInstanceOnboardingJobStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceOnboardingJobStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCampaigns", func(t *testing.T) {
        input := &connectcampaignsv2.ListCampaignsInput{}
        output := &connectcampaignsv2.ListCampaignsOutput{}

        mockClient.On("ListCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectInstanceIntegrations", func(t *testing.T) {
        input := &connectcampaignsv2.ListConnectInstanceIntegrationsInput{}
        output := &connectcampaignsv2.ListConnectInstanceIntegrationsOutput{}

        mockClient.On("ListConnectInstanceIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectInstanceIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &connectcampaignsv2.ListTagsForResourceInput{}
        output := &connectcampaignsv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.PauseCampaignInput{}
        output := &connectcampaignsv2.PauseCampaignOutput{}

        mockClient.On("PauseCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.PauseCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutConnectInstanceIntegration", func(t *testing.T) {
        input := &connectcampaignsv2.PutConnectInstanceIntegrationInput{}
        output := &connectcampaignsv2.PutConnectInstanceIntegrationOutput{}

        mockClient.On("PutConnectInstanceIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.PutConnectInstanceIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOutboundRequestBatch", func(t *testing.T) {
        input := &connectcampaignsv2.PutOutboundRequestBatchInput{}
        output := &connectcampaignsv2.PutOutboundRequestBatchOutput{}

        mockClient.On("PutOutboundRequestBatch", ctx, input).Return(output, nil)

        result, err := mockClient.PutOutboundRequestBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProfileOutboundRequestBatch", func(t *testing.T) {
        input := &connectcampaignsv2.PutProfileOutboundRequestBatchInput{}
        output := &connectcampaignsv2.PutProfileOutboundRequestBatchOutput{}

        mockClient.On("PutProfileOutboundRequestBatch", ctx, input).Return(output, nil)

        result, err := mockClient.PutProfileOutboundRequestBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.ResumeCampaignInput{}
        output := &connectcampaignsv2.ResumeCampaignOutput{}

        mockClient.On("ResumeCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.StartCampaignInput{}
        output := &connectcampaignsv2.StartCampaignOutput{}

        mockClient.On("StartCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.StartCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstanceOnboardingJob", func(t *testing.T) {
        input := &connectcampaignsv2.StartInstanceOnboardingJobInput{}
        output := &connectcampaignsv2.StartInstanceOnboardingJobOutput{}

        mockClient.On("StartInstanceOnboardingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstanceOnboardingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCampaign", func(t *testing.T) {
        input := &connectcampaignsv2.StopCampaignInput{}
        output := &connectcampaignsv2.StopCampaignOutput{}

        mockClient.On("StopCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.StopCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &connectcampaignsv2.TagResourceInput{}
        output := &connectcampaignsv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &connectcampaignsv2.UntagResourceInput{}
        output := &connectcampaignsv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignChannelSubtypeConfig", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignChannelSubtypeConfigInput{}
        output := &connectcampaignsv2.UpdateCampaignChannelSubtypeConfigOutput{}

        mockClient.On("UpdateCampaignChannelSubtypeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignChannelSubtypeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignCommunicationLimits", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignCommunicationLimitsInput{}
        output := &connectcampaignsv2.UpdateCampaignCommunicationLimitsOutput{}

        mockClient.On("UpdateCampaignCommunicationLimits", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignCommunicationLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignCommunicationTime", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignCommunicationTimeInput{}
        output := &connectcampaignsv2.UpdateCampaignCommunicationTimeOutput{}

        mockClient.On("UpdateCampaignCommunicationTime", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignCommunicationTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignFlowAssociation", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignFlowAssociationInput{}
        output := &connectcampaignsv2.UpdateCampaignFlowAssociationOutput{}

        mockClient.On("UpdateCampaignFlowAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignFlowAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignName", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignNameInput{}
        output := &connectcampaignsv2.UpdateCampaignNameOutput{}

        mockClient.On("UpdateCampaignName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignSchedule", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignScheduleInput{}
        output := &connectcampaignsv2.UpdateCampaignScheduleOutput{}

        mockClient.On("UpdateCampaignSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignSource", func(t *testing.T) {
        input := &connectcampaignsv2.UpdateCampaignSourceInput{}
        output := &connectcampaignsv2.UpdateCampaignSourceOutput{}

        mockClient.On("UpdateCampaignSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
